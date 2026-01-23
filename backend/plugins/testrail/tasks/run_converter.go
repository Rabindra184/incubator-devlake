/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tasks

import (
	"reflect"
	"time"

	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/models/domainlayer"
	"github.com/apache/incubator-devlake/core/models/domainlayer/didgen"
	"github.com/apache/incubator-devlake/core/models/domainlayer/qa"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/testrail/models"
)

var ConvertRunsMeta = plugin.SubTaskMeta{
	Name:             "convertRuns",
	EntryPoint:       ConvertRuns,
	EnabledByDefault: true,
	Description:      "Convert tool layer table _tool_testrail_runs into domain layer table qa_test_runs",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func ConvertRuns(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	db := taskCtx.GetDal()

	cursor, err := db.Cursor(
		dal.From(&models.TestrailRun{}),
		dal.Where("connection_id = ? AND project_id = ?", data.Options.ConnectionId, data.Options.ProjectId),
	)
	if err != nil {
		return err
	}
	defer cursor.Close()

	runIdGen := didgen.NewDomainIdGenerator(&models.TestrailRun{})
	projectIdGen := didgen.NewDomainIdGenerator(&models.TestrailProject{})

	converter, err := helper.NewDataConverter(helper.DataConverterArgs{
		InputRowType: reflect.TypeOf(models.TestrailRun{}),
		Input:        cursor,
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_RUN_TABLE,
		},
		Convert: func(inputRow interface{}) ([]interface{}, errors.Error) {
			run := inputRow.(*models.TestrailRun)

			startTime := time.Unix(run.CreatedOn, 0)
			var finishTime *time.Time
			if run.CompletedOn > 0 {
				ft := time.Unix(run.CompletedOn, 0)
				finishTime = &ft
			}

			status := "IN_PROGRESS"
			if run.IsCompleted {
				status = "COMPLETED"
			}

			domainRun := &qa.QaTestRun{
				DomainEntityExtended: domainlayer.DomainEntityExtended{
					Id: runIdGen.Generate(run.ConnectionId, run.Id),
				},
				QaProjectId:  projectIdGen.Generate(run.ConnectionId, run.ProjectId),
				Name:         run.Name,
				Description:  run.Description,
				StartTime:    &startTime,
				FinishTime:   finishTime,
				Status:       status,
				PassedCount:  run.PassedCount,
				FailedCount:  run.FailedCount,
				SkippedCount: run.BlockedCount, // Tests blocked are usually considered skipped in high-level domain
				TotalCount:   run.PassedCount + run.FailedCount + run.BlockedCount + run.UntestedCount + run.RetestCount,
			}

			return []interface{}{domainRun}, nil
		},
	})
	if err != nil {
		return err
	}

	return converter.Execute()
}
