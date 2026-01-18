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
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/testrail/models"
)

var CollectResultsMeta = plugin.SubTaskMeta{
	Name:             "collectResults",
	EntryPoint:       CollectResults,
	EnabledByDefault: true,
	Description:      "Collect results data from Testrail api",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func CollectResults(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	db := taskCtx.GetDal()
	logger := taskCtx.GetLogger()
	logger.Info("collecting results")

	cursor, err := db.Cursor(
		dal.From(models.TestrailRun{}),
		dal.Where("connection_id = ? AND project_id = ?", data.Options.ConnectionId, data.Options.ProjectId),
	)
	if err != nil {
		return err
	}
	iterator, err := helper.NewDalCursorIterator(db, cursor, reflect.TypeOf(models.TestrailRun{}))
	if err != nil {
		return err
	}

	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_RESULT_TABLE,
		},
		ApiClient:   data.ApiClient,
		PageSize:    250,
		Incremental: false,
		Input:       iterator,
		UrlTemplate: "index.php?/api/v2/get_results_for_run/{{ .Input.Id }}",
		Query: func(reqData *helper.RequestData) (url.Values, errors.Error) {
			query := url.Values{}
			query.Set("limit", strconv.Itoa(reqData.Pager.Size))
			query.Set("offset", strconv.Itoa(reqData.Pager.Skip))
			return query, nil
		},
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			var result struct {
				Results []json.RawMessage `json:"results"`
			}
			err := helper.UnmarshalResponse(res, &result)
			if err != nil {
				return nil, err
			}
			return result.Results, nil
		},
	})

	if err != nil {
		return err
	}

	return collector.Execute()
}
