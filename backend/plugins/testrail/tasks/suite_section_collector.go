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
	"strconv"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

var CollectSuitesMeta = plugin.SubTaskMeta{
	Name:             "collectSuites",
	EntryPoint:       CollectSuites,
	EnabledByDefault: true,
	Description:      "Collect suites data from Testrail api",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func CollectSuites(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	logger := taskCtx.GetLogger()
	logger.Info("collecting suites for project %d", data.Options.ProjectId)

	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_SUITE_TABLE,
		},
		ApiClient:   data.ApiClient,
		PageSize:    250,
		Incremental: false,
		// Suites API doesn't use pagination, returns all suites for a project
		UrlTemplate: "index.php?/api/v2/get_suites/{{ .Params.ProjectId }}",
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			var suites []json.RawMessage
			err := helper.UnmarshalResponse(res, &suites)
			if err != nil {
				return nil, err
			}
			return suites, nil
		},
	})

	if err != nil {
		return err
	}

	return collector.Execute()
}

var CollectSectionsMeta = plugin.SubTaskMeta{
	Name:             "collectSections",
	EntryPoint:       CollectSections,
	EnabledByDefault: true,
	Description:      "Collect sections data from Testrail api",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func CollectSections(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	logger := taskCtx.GetLogger()
	logger.Info("collecting sections for project %d", data.Options.ProjectId)

	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_SECTION_TABLE,
		},
		ApiClient:   data.ApiClient,
		PageSize:    250,
		Incremental: false,
		UrlTemplate: "index.php?/api/v2/get_sections/{{ .Params.ProjectId }}",
		Query: func(reqData *helper.RequestData) (url.Values, errors.Error) {
			query := url.Values{}
			query.Set("limit", strconv.Itoa(reqData.Pager.Size))
			query.Set("offset", strconv.Itoa(reqData.Pager.Skip))
			return query, nil
		},
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			var result struct {
				Sections []json.RawMessage `json:"sections"`
			}
			err := helper.UnmarshalResponse(res, &result)
			if err != nil {
				return nil, err
			}
			return result.Sections, nil
		},
	})

	if err != nil {
		return err
	}

	return collector.Execute()
}
