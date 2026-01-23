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
	"fmt"
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

var ConvertResultsMeta = plugin.SubTaskMeta{
	Name:             "convertResults",
	EntryPoint:       ConvertResults,
	EnabledByDefault: true,
	Description:      "Convert tool layer table testrail_results into domain layer table qa_test_case_executions",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func ConvertResults(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	db := taskCtx.GetDal()

	cursor, err := db.Cursor(
		dal.From(&models.TestrailResult{}),
		dal.Where("connection_id = ?", data.Options.ConnectionId),
		// We might want to filter by project but Results don't have project_id directly in the table I defined.
		// However, we collect results for project, so maybe we should add project_id to the tool layer model.
	)
	if err != nil {
		return err
	}
	defer cursor.Close()

	resultIdGen := didgen.NewDomainIdGenerator(&models.TestrailResult{})
	caseIdGen := didgen.NewDomainIdGenerator(&models.TestrailCase{})
	projectIdGen := didgen.NewDomainIdGenerator(&models.TestrailProject{})

	converter, err := helper.NewDataConverter(helper.DataConverterArgs{
		InputRowType: reflect.TypeOf(models.TestrailResult{}),
		Input:        cursor,
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_RESULT_TABLE,
		},
		Convert: func(inputRow interface{}) ([]interface{}, errors.Error) {
			result := inputRow.(*models.TestrailResult)
			// Generate executor ID if available
			var executorId string
			if result.CreatedBy > 0 {
				userIdGen := didgen.NewDomainIdGenerator(&models.TestrailUser{})
				executorId = userIdGen.Generate(result.ConnectionId, result.CreatedBy)
			}

			domainExecution := &qa.QaTestCaseExecution{
				DomainEntityExtended: domainlayer.DomainEntityExtended{
					Id: resultIdGen.Generate(result.ConnectionId, result.Id),
				},
				QaProjectId:  projectIdGen.Generate(result.ConnectionId, data.Options.ProjectId),
				QaTestCaseId: caseIdGen.Generate(result.ConnectionId, result.CaseId),
				CreateTime:   time.Unix(result.CreatedOn, 0),
				StartTime:    time.Unix(result.CreatedOn, 0),
				FinishTime:   time.Unix(result.CreatedOn, 0),
				CreatorId:    executorId,
				Status:       mapStatusWithConfig(result.StatusId, data.Options.ScopeConfig),
			}
			return []interface{}{domainExecution}, nil
		},
	})
	if err != nil {
		return err
	}

	return converter.Execute()
}

func mapStatus(statusId int) string {
	// Default status mapping based on TestRail's standard statuses
	// Users can override these via ScopeConfig.StatusMappings
	switch statusId {
	case 1: // Passed
		return "SUCCESS"
	case 2: // Blocked
		return "BLOCKED"
	case 3: // Untested
		return "PENDING"
	case 4: // Retest
		return "PENDING"
	case 5: // Failed
		return "FAILED"
	case 6: // Custom status 1 (often used for "Skipped")
		return "SKIPPED"
	case 7: // Custom status 2
		return "PENDING"
	case 8: // Custom status 3
		return "PENDING"
	case 9: // Custom status 4
		return "PENDING"
	case 10: // Custom status 5
		return "PENDING"
	default:
		return "PENDING"
	}
}

// mapStatusWithConfig uses the scope configuration to map statuses if available
func mapStatusWithConfig(statusId int, scopeConfig *models.TestrailScopeConfig) string {
	if scopeConfig != nil && scopeConfig.StatusMappings != nil {
		statusKey := fmt.Sprintf("%d", statusId)
		if mapping, ok := scopeConfig.StatusMappings[statusKey]; ok {
			return mapping.StandardStatus
		}
	}
	return mapStatus(statusId)
}
