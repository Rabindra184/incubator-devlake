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

package models

import (
	"github.com/apache/incubator-devlake/core/models/common"
)

type TestrailRun struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	ProjectId        uint64 `gorm:"index;type:BIGINT NOT NULL" json:"project_id"`
	SuiteId          uint64 `json:"suite_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	MilestoneId      uint64 `json:"milestone_id"`
	AssignedtoId     uint64 `json:"assignedto_id"`
	IsCompleted      bool   `json:"is_completed"`
	CompletedOn      int64  `json:"completed_on"`
	PassedCount      int    `json:"passed_count"`
	BlockedCount     int    `json:"blocked_count"`
	UntestedCount    int    `json:"untested_count"`
	RetestCount      int    `json:"retest_count"`
	FailedCount      int    `json:"failed_count"`
	CreatedOn        int64  `json:"created_on"`
	Url              string `json:"url"`
}

func (TestrailRun) TableName() string {
	return "_tool_testrail_runs"
}

type TestrailResult struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	RunId            uint64 `gorm:"index;type:BIGINT NOT NULL" json:"run_id"`
	CaseId           uint64 `gorm:"index;type:BIGINT NOT NULL" json:"case_id"`
	StatusId         int    `json:"status_id"`
	CreatedBy        uint64 `json:"created_by"`
	CreatedOn        int64  `json:"created_on"`
	Elapsed          string `json:"elapsed"`
	Comment          string `json:"comment"`
}

func (TestrailResult) TableName() string {
	return "_tool_testrail_results"
}
