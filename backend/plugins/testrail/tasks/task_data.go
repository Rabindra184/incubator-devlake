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
	"time"

	"github.com/apache/incubator-devlake/core/errors"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/testrail/models"
)

const RAW_PROJECT_TABLE = "testrail_projects"
const RAW_CASE_TABLE = "testrail_cases"
const RAW_RUN_TABLE = "testrail_runs"
const RAW_RESULT_TABLE = "testrail_results"
const RAW_SUITE_TABLE = "testrail_suites"
const RAW_SECTION_TABLE = "testrail_sections"
const RAW_PLAN_TABLE = "testrail_plans"
const RAW_MILESTONE_TABLE = "testrail_milestones"
const RAW_USER_TABLE = "testrail_users"
const RAW_CASE_FIELD_TABLE = "testrail_case_fields"
const RAW_CASE_TYPE_TABLE = "testrail_case_types"
const RAW_PRIORITY_TABLE = "testrail_priorities"
const RAW_STATUS_TABLE = "testrail_statuses"

type TestrailApiParams struct {
	ConnectionId uint64 `json:"connectionId"`
	ProjectId    uint64 `json:"projectId"`
}

type TestrailOptions struct {
	ConnectionId uint64 `json:"connectionId"`
	ProjectId    uint64 `json:"projectId"`
	Name         string `json:"name"`
	ScopeConfig  *models.TestrailScopeConfig
	Title        string `json:"title"`

	// Time filter
	CreatedDateAfter *time.Time
}

type TestrailTaskData struct {
	Options   *TestrailOptions
	ApiClient *helper.ApiAsyncClient
}

func DecodeTaskOptions(options map[string]interface{}) (*TestrailOptions, errors.Error) {
	var op TestrailOptions
	err := helper.Decode(options, &op, nil)
	if err != nil {
		return nil, err
	}
	if op.ProjectId == 0 {
		return nil, errors.BadInput.New(fmt.Sprintf("invalid projectId:%d", op.ProjectId))
	}
	if op.ConnectionId == 0 {
		return nil, errors.BadInput.New(fmt.Sprintf("invalid connectionId:%d", op.ConnectionId))
	}
	return &op, nil
}
