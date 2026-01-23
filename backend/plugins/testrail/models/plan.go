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

// TestrailPlan represents a test plan in TestRail
// Test plans allow you to group test runs and execute them against milestones
type TestrailPlan struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	ProjectId        uint64 `gorm:"index;type:BIGINT NOT NULL" json:"project_id"`
	Name             string `gorm:"type:varchar(255)" json:"name"`
	Description      string `gorm:"type:text" json:"description"`
	MilestoneId      uint64 `gorm:"index" json:"milestone_id"`
	AssignedtoId     uint64 `gorm:"index" json:"assignedto_id"`
	IsCompleted      bool   `json:"is_completed"`
	CompletedOn      int64  `json:"completed_on"`
	CreatedOn        int64  `json:"created_on"`
	CreatedBy        uint64 `json:"created_by"`
	Url              string `gorm:"type:varchar(255)" json:"url"`

	// Aggregate counts from plan entries
	PassedCount   int `json:"passed_count"`
	BlockedCount  int `json:"blocked_count"`
	UntestedCount int `json:"untested_count"`
	RetestCount   int `json:"retest_count"`
	FailedCount   int `json:"failed_count"`
}

func (TestrailPlan) TableName() string {
	return "_tool_testrail_plans"
}

// TestrailPlanEntry represents an entry (run configuration) within a test plan
type TestrailPlanEntry struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               string `gorm:"primaryKey;type:varchar(255) NOT NULL" json:"id"` // UUID string
	PlanId           uint64 `gorm:"index;type:BIGINT NOT NULL" json:"plan_id"`
	SuiteId          uint64 `gorm:"index" json:"suite_id"`
	Name             string `gorm:"type:varchar(255)" json:"name"`
	Description      string `gorm:"type:text" json:"description"`
	AssignedtoId     uint64 `json:"assignedto_id"`
	IncludeAll       bool   `json:"include_all"`
}

func (TestrailPlanEntry) TableName() string {
	return "_tool_testrail_plan_entries"
}
