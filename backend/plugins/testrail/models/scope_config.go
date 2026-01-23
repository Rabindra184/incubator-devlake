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

// StatusMapping defines how TestRail statuses map to standard DevLake statuses
type StatusMapping struct {
	StandardStatus string `json:"standardStatus"` // SUCCESS, FAILED, PENDING, BLOCKED, SKIPPED
}

// TypeMapping defines how TestRail case types map to standard types
type TypeMapping struct {
	StandardType string `json:"standardType"` // functional, api, integration, unit, etc.
}

// PriorityMapping defines how TestRail priorities map to standard priorities
type PriorityMapping struct {
	StandardPriority string `json:"standardPriority"` // critical, high, medium, low
}

// CustomFieldMapping defines how to interpret a custom field
type CustomFieldMapping struct {
	TargetField string `json:"targetField"` // Where to map this field in domain layer
	FieldType   string `json:"fieldType"`   // string, boolean, number, date
}

type TestrailScopeConfig struct {
	common.ScopeConfig `mapstructure:",squash" json:",inline" gorm:"embedded"`

	// Status Mappings: Map TestRail status IDs/names to standard statuses
	// Example: {"1": {"standardStatus": "SUCCESS"}, "5": {"standardStatus": "FAILED"}}
	StatusMappings map[string]StatusMapping `mapstructure:"statusMappings" json:"statusMappings" gorm:"type:json;serializer:json"`

	// Type Mappings: Map TestRail case type IDs/names to standard types
	// Example: {"1": {"standardType": "functional"}, "2": {"standardType": "regression"}}
	TypeMappings map[string]TypeMapping `mapstructure:"typeMappings" json:"typeMappings" gorm:"type:json;serializer:json"`

	// Priority Mappings: Map TestRail priority IDs/names to standard priorities
	// Example: {"1": {"standardPriority": "low"}, "4": {"standardPriority": "critical"}}
	PriorityMappings map[string]PriorityMapping `mapstructure:"priorityMappings" json:"priorityMappings" gorm:"type:json;serializer:json"`

	// Custom Field Mappings: Define how to handle specific custom fields
	// Key is the system_name of the custom field (e.g., "custom_automation_type")
	CustomFieldMappings map[string]CustomFieldMapping `mapstructure:"customFieldMappings" json:"customFieldMappings" gorm:"type:json;serializer:json"`

	// Automation Status Field: The custom field that indicates automation status
	// Example: "custom_automation_type" or "custom_automated"
	AutomationStatusField string `mapstructure:"automationStatusField" json:"automationStatusField" gorm:"type:varchar(255)"`

	// Values that indicate a test is automated
	// Example: ["automated", "yes", "true", "1"]
	AutomatedValues []string `mapstructure:"automatedValues" json:"automatedValues" gorm:"type:json;serializer:json"`

	// Suite filter: Only sync specific suites (empty means all)
	IncludeSuiteIds []uint64 `mapstructure:"includeSuiteIds" json:"includeSuiteIds" gorm:"type:json;serializer:json"`
	ExcludeSuiteIds []uint64 `mapstructure:"excludeSuiteIds" json:"excludeSuiteIds" gorm:"type:json;serializer:json"`

	// Milestone filter: Only sync specific milestones (empty means all)
	IncludeMilestoneIds []uint64 `mapstructure:"includeMilestoneIds" json:"includeMilestoneIds" gorm:"type:json;serializer:json"`
}

func (TestrailScopeConfig) TableName() string {
	return "_tool_testrail_scope_configs"
}
