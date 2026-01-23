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

// TestrailMilestone represents a milestone in TestRail
// Milestones are used to track release/sprint goals and can contain sub-milestones
type TestrailMilestone struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	ProjectId        uint64 `gorm:"index;type:BIGINT NOT NULL" json:"project_id"`
	ParentId         uint64 `gorm:"index" json:"parent_id"` // For sub-milestones
	Name             string `gorm:"type:varchar(255)" json:"name"`
	Description      string `gorm:"type:text" json:"description"`
	Refs             string `gorm:"type:varchar(255)" json:"refs"` // Reference/version string
	Url              string `gorm:"type:varchar(255)" json:"url"`
	StartOn          int64  `json:"start_on"`   // Unix timestamp
	StartedOn        int64  `json:"started_on"` // Actual start
	DueOn            int64  `json:"due_on"`     // Deadline
	IsCompleted      bool   `json:"is_completed"`
	IsStarted        bool   `json:"is_started"`
	CompletedOn      int64  `json:"completed_on"`
}

func (TestrailMilestone) TableName() string {
	return "_tool_testrail_milestones"
}
