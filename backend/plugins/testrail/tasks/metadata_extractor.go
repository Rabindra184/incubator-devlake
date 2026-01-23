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

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/testrail/models"
)

var ExtractUsersMeta = plugin.SubTaskMeta{
	Name:             "extractUsers",
	EntryPoint:       ExtractUsers,
	EnabledByDefault: true,
	Description:      "Extract raw users data into tool layer table testrail_users",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func ExtractUsers(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_USER_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, errors.Error) {
			var apiUser models.TestrailUser
			err := json.Unmarshal(row.Data, &apiUser)
			if err != nil {
				return nil, errors.Default.Wrap(err, "error unmarshaling user")
			}

			apiUser.ConnectionId = data.Options.ConnectionId
			return []interface{}{&apiUser}, nil
		},
	})

	if err != nil {
		return err
	}

	return extractor.Execute()
}

var ExtractCaseFieldsMeta = plugin.SubTaskMeta{
	Name:             "extractCaseFields",
	EntryPoint:       ExtractCaseFields,
	EnabledByDefault: true,
	Description:      "Extract raw case fields data into tool layer table testrail_case_fields",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func ExtractCaseFields(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_CASE_FIELD_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, errors.Error) {
			var apiField models.TestrailCaseField
			err := json.Unmarshal(row.Data, &apiField)
			if err != nil {
				return nil, errors.Default.Wrap(err, "error unmarshaling case field")
			}

			apiField.ConnectionId = data.Options.ConnectionId
			return []interface{}{&apiField}, nil
		},
	})

	if err != nil {
		return err
	}

	return extractor.Execute()
}

var ExtractCaseTypesMeta = plugin.SubTaskMeta{
	Name:             "extractCaseTypes",
	EntryPoint:       ExtractCaseTypes,
	EnabledByDefault: true,
	Description:      "Extract raw case types data into tool layer table testrail_case_types",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func ExtractCaseTypes(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_CASE_TYPE_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, errors.Error) {
			var apiType models.TestrailCaseType
			err := json.Unmarshal(row.Data, &apiType)
			if err != nil {
				return nil, errors.Default.Wrap(err, "error unmarshaling case type")
			}

			apiType.ConnectionId = data.Options.ConnectionId
			return []interface{}{&apiType}, nil
		},
	})

	if err != nil {
		return err
	}

	return extractor.Execute()
}

var ExtractPrioritiesMeta = plugin.SubTaskMeta{
	Name:             "extractPriorities",
	EntryPoint:       ExtractPriorities,
	EnabledByDefault: true,
	Description:      "Extract raw priorities data into tool layer table testrail_priorities",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func ExtractPriorities(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_PRIORITY_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, errors.Error) {
			var apiPriority models.TestrailPriority
			err := json.Unmarshal(row.Data, &apiPriority)
			if err != nil {
				return nil, errors.Default.Wrap(err, "error unmarshaling priority")
			}

			apiPriority.ConnectionId = data.Options.ConnectionId
			return []interface{}{&apiPriority}, nil
		},
	})

	if err != nil {
		return err
	}

	return extractor.Execute()
}

var ExtractStatusesMeta = plugin.SubTaskMeta{
	Name:             "extractStatuses",
	EntryPoint:       ExtractStatuses,
	EnabledByDefault: true,
	Description:      "Extract raw statuses data into tool layer table testrail_statuses",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func ExtractStatuses(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_STATUS_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, errors.Error) {
			var apiStatus models.TestrailStatus
			err := json.Unmarshal(row.Data, &apiStatus)
			if err != nil {
				return nil, errors.Default.Wrap(err, "error unmarshaling status")
			}

			apiStatus.ConnectionId = data.Options.ConnectionId
			return []interface{}{&apiStatus}, nil
		},
	})

	if err != nil {
		return err
	}

	return extractor.Execute()
}
