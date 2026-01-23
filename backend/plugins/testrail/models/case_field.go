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
	"encoding/json"

	"github.com/apache/incubator-devlake/core/models/common"
)

// TestrailCaseField represents a custom field definition in TestRail
// TestRail allows defining custom fields (custom_*) for test cases
type TestrailCaseField struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64          `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64          `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	SystemName       string          `gorm:"type:varchar(255);index" json:"system_name"` // e.g., "custom_automation_type"
	Label            string          `gorm:"type:varchar(255)" json:"label"`
	Name             string          `gorm:"type:varchar(255)" json:"name"`
	Description      string          `gorm:"type:text" json:"description"`
	TypeId           int             `json:"type_id"` // 1=String, 2=Integer, 3=Text, 4=URL, 5=Checkbox, 6=Dropdown, 7=User, 8=Date, 9=Milestone, 10=Steps, 11=Step Results, 12=Multi-select
	IsGlobal         bool            `json:"is_global"`
	IsActive         bool            `json:"is_active"`
	DisplayOrder     int             `json:"display_order"`
	IncludeAll       bool            `json:"include_all"`                   // Applies to all projects
	TemplateIds      json.RawMessage `gorm:"type:json" json:"template_ids"` // Project template restrictions
	Configs          json.RawMessage `gorm:"type:json" json:"configs"`      // Field configuration (options for dropdowns, etc.)
}

func (TestrailCaseField) TableName() string {
	return "_tool_testrail_case_fields"
}

// TestrailCaseType represents a test case type in TestRail
type TestrailCaseType struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	Name             string `gorm:"type:varchar(255)" json:"name"`
	IsDefault        bool   `json:"is_default"`
}

func (TestrailCaseType) TableName() string {
	return "_tool_testrail_case_types"
}

// TestrailPriority represents a priority level in TestRail
type TestrailPriority struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	Name             string `gorm:"type:varchar(255)" json:"name"`
	ShortName        string `gorm:"type:varchar(50)" json:"short_name"`
	Priority         int    `json:"priority"` // Sort order/weight
	IsDefault        bool   `json:"is_default"`
}

func (TestrailPriority) TableName() string {
	return "_tool_testrail_priorities"
}

// TestrailStatus represents a test result status in TestRail
type TestrailStatus struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	Name             string `gorm:"type:varchar(255)" json:"name"`
	Label            string `gorm:"type:varchar(255)" json:"label"`
	ColorDark        int    `json:"color_dark"`
	ColorMedium      int    `json:"color_medium"`
	ColorBright      int    `json:"color_bright"`
	IsFinal          bool   `json:"is_final"`
	IsSystem         bool   `json:"is_system"`
	IsUntested       bool   `json:"is_untested"`
}

func (TestrailStatus) TableName() string {
	return "_tool_testrail_statuses"
}
