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

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

var CollectUsersMeta = plugin.SubTaskMeta{
	Name:             "collectUsers",
	EntryPoint:       CollectUsers,
	EnabledByDefault: true,
	Description:      "Collect users data from Testrail api",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func CollectUsers(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	logger := taskCtx.GetLogger()
	logger.Info("collecting users")

	// Users API returns all users for the TestRail instance (not project-specific)
	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_USER_TABLE,
		},
		ApiClient:   data.ApiClient,
		PageSize:    1, // Users endpoint returns all users in one call
		Incremental: false,
		UrlTemplate: "index.php?/api/v2/get_users",
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			var users []json.RawMessage
			err := helper.UnmarshalResponse(res, &users)
			if err != nil {
				return nil, err
			}
			return users, nil
		},
	})

	if err != nil {
		return err
	}

	return collector.Execute()
}

var CollectCaseFieldsMeta = plugin.SubTaskMeta{
	Name:             "collectCaseFields",
	EntryPoint:       CollectCaseFields,
	EnabledByDefault: true,
	Description:      "Collect case field definitions from Testrail api",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func CollectCaseFields(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	logger := taskCtx.GetLogger()
	logger.Info("collecting case fields")

	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_CASE_FIELD_TABLE,
		},
		ApiClient:   data.ApiClient,
		PageSize:    1,
		Incremental: false,
		UrlTemplate: "index.php?/api/v2/get_case_fields",
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			var fields []json.RawMessage
			err := helper.UnmarshalResponse(res, &fields)
			if err != nil {
				return nil, err
			}
			return fields, nil
		},
	})

	if err != nil {
		return err
	}

	return collector.Execute()
}

var CollectCaseTypesMeta = plugin.SubTaskMeta{
	Name:             "collectCaseTypes",
	EntryPoint:       CollectCaseTypes,
	EnabledByDefault: true,
	Description:      "Collect case types from Testrail api",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func CollectCaseTypes(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	logger := taskCtx.GetLogger()
	logger.Info("collecting case types")

	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_CASE_TYPE_TABLE,
		},
		ApiClient:   data.ApiClient,
		PageSize:    1,
		Incremental: false,
		UrlTemplate: "index.php?/api/v2/get_case_types",
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			var types []json.RawMessage
			err := helper.UnmarshalResponse(res, &types)
			if err != nil {
				return nil, err
			}
			return types, nil
		},
	})

	if err != nil {
		return err
	}

	return collector.Execute()
}

var CollectPrioritiesMeta = plugin.SubTaskMeta{
	Name:             "collectPriorities",
	EntryPoint:       CollectPriorities,
	EnabledByDefault: true,
	Description:      "Collect priorities from Testrail api",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func CollectPriorities(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	logger := taskCtx.GetLogger()
	logger.Info("collecting priorities")

	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_PRIORITY_TABLE,
		},
		ApiClient:   data.ApiClient,
		PageSize:    1,
		Incremental: false,
		UrlTemplate: "index.php?/api/v2/get_priorities",
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			var priorities []json.RawMessage
			err := helper.UnmarshalResponse(res, &priorities)
			if err != nil {
				return nil, err
			}
			return priorities, nil
		},
	})

	if err != nil {
		return err
	}

	return collector.Execute()
}

var CollectStatusesMeta = plugin.SubTaskMeta{
	Name:             "collectStatuses",
	EntryPoint:       CollectStatuses,
	EnabledByDefault: true,
	Description:      "Collect result statuses from Testrail api",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func CollectStatuses(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	logger := taskCtx.GetLogger()
	logger.Info("collecting statuses")

	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_STATUS_TABLE,
		},
		ApiClient:   data.ApiClient,
		PageSize:    1,
		Incremental: false,
		UrlTemplate: "index.php?/api/v2/get_statuses",
		ResponseParser: func(res *http.Response) ([]json.RawMessage, errors.Error) {
			var statuses []json.RawMessage
			err := helper.UnmarshalResponse(res, &statuses)
			if err != nil {
				return nil, err
			}
			return statuses, nil
		},
	})

	if err != nil {
		return err
	}

	return collector.Execute()
}
