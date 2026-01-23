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

var ExtractSuitesMeta = plugin.SubTaskMeta{
	Name:             "extractSuites",
	EntryPoint:       ExtractSuites,
	EnabledByDefault: true,
	Description:      "Extract raw suites data into tool layer table testrail_suites",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func ExtractSuites(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_SUITE_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, errors.Error) {
			var apiSuite models.TestrailSuite
			err := json.Unmarshal(row.Data, &apiSuite)
			if err != nil {
				return nil, errors.Default.Wrap(err, "error unmarshaling suite")
			}

			apiSuite.ConnectionId = data.Options.ConnectionId
			apiSuite.ProjectId = data.Options.ProjectId
			return []interface{}{&apiSuite}, nil
		},
	})

	if err != nil {
		return err
	}

	return extractor.Execute()
}

var ExtractSectionsMeta = plugin.SubTaskMeta{
	Name:             "extractSections",
	EntryPoint:       ExtractSections,
	EnabledByDefault: true,
	Description:      "Extract raw sections data into tool layer table testrail_sections",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CODE_QUALITY},
}

func ExtractSections(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TestrailTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TestrailApiParams{
				ConnectionId: data.Options.ConnectionId,
				ProjectId:    data.Options.ProjectId,
			},
			Table: RAW_SECTION_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, errors.Error) {
			var apiSection models.TestrailSection
			err := json.Unmarshal(row.Data, &apiSection)
			if err != nil {
				return nil, errors.Default.Wrap(err, "error unmarshaling section")
			}

			apiSection.ConnectionId = data.Options.ConnectionId
			apiSection.ProjectId = data.Options.ProjectId
			return []interface{}{&apiSection}, nil
		},
	})

	if err != nil {
		return err
	}

	return extractor.Execute()
}
