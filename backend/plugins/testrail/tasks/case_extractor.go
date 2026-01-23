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

var ExtractCasesMeta = plugin.SubTaskMeta{
	Name:             "extractCases",
	EntryPoint:       ExtractCases,
	EnabledByDefault: true,
	Description:      "Extract raw cases data into tool layer table testrail_cases",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func ExtractCases(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_CASE_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, errors.Error) {
			var apiCase models.TestrailCase
			err := json.Unmarshal(row.Data, &apiCase)
			if err != nil {
				return nil, errors.Default.Wrap(err, "error unmarshaling case")
			}

			// Extract custom fields from raw data
			var rawData map[string]interface{}
			if err := json.Unmarshal(row.Data, &rawData); err == nil {
				customFields := make(map[string]interface{})
				for key, value := range rawData {
					if len(key) > 7 && key[:7] == "custom_" {
						customFields[key] = value
					}
				}
				if len(customFields) > 0 {
					if customFieldsJSON, err := json.Marshal(customFields); err == nil {
						apiCase.CustomFields = string(customFieldsJSON)
					}
				}
			}

			apiCase.ConnectionId = data.Options.ConnectionId
			return []interface{}{&apiCase}, nil
		},
	})

	if err != nil {
		return err
	}

	return extractor.Execute()
}
