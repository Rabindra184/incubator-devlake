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

package migrationscripts

import (
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/migrationhelper"
)

type addEnterpriseEntities struct{}

// TestrailSuite20260123 - Archived model for migration
type TestrailSuite20260123 struct {
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id           uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false"`
	ProjectId    uint64 `gorm:"index;type:BIGINT NOT NULL"`
	Name         string `gorm:"type:varchar(255)"`
	Description  string `gorm:"type:text"`
	Url          string `gorm:"type:varchar(255)"`
	IsBaseline   bool
	IsMaster     bool
	IsCompleted  bool
	CompletedOn  int64
}

func (TestrailSuite20260123) TableName() string {
	return "_tool_testrail_suites"
}

// TestrailSection20260123 - Archived model for migration
type TestrailSection20260123 struct {
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id           uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false"`
	ProjectId    uint64 `gorm:"index;type:BIGINT NOT NULL"`
	SuiteId      uint64 `gorm:"index"`
	ParentId     uint64 `gorm:"index"`
	Name         string `gorm:"type:varchar(255)"`
	Description  string `gorm:"type:text"`
	DisplayOrder int
	Depth        int
}

func (TestrailSection20260123) TableName() string {
	return "_tool_testrail_sections"
}

// TestrailPlan20260123 - Archived model for migration
type TestrailPlan20260123 struct {
	ConnectionId  uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id            uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false"`
	ProjectId     uint64 `gorm:"index;type:BIGINT NOT NULL"`
	Name          string `gorm:"type:varchar(255)"`
	Description   string `gorm:"type:text"`
	MilestoneId   uint64 `gorm:"index"`
	AssignedtoId  uint64 `gorm:"index"`
	IsCompleted   bool
	CompletedOn   int64
	CreatedOn     int64
	CreatedBy     uint64
	Url           string `gorm:"type:varchar(255)"`
	PassedCount   int
	BlockedCount  int
	UntestedCount int
	RetestCount   int
	FailedCount   int
}

func (TestrailPlan20260123) TableName() string {
	return "_tool_testrail_plans"
}

// TestrailPlanEntry20260123 - Archived model for migration
type TestrailPlanEntry20260123 struct {
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id           string `gorm:"primaryKey;type:varchar(255) NOT NULL"`
	PlanId       uint64 `gorm:"index;type:BIGINT NOT NULL"`
	SuiteId      uint64 `gorm:"index"`
	Name         string `gorm:"type:varchar(255)"`
	Description  string `gorm:"type:text"`
	AssignedtoId uint64
	IncludeAll   bool
}

func (TestrailPlanEntry20260123) TableName() string {
	return "_tool_testrail_plan_entries"
}

// TestrailMilestone20260123 - Archived model for migration
type TestrailMilestone20260123 struct {
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id           uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false"`
	ProjectId    uint64 `gorm:"index;type:BIGINT NOT NULL"`
	ParentId     uint64 `gorm:"index"`
	Name         string `gorm:"type:varchar(255)"`
	Description  string `gorm:"type:text"`
	Refs         string `gorm:"type:varchar(255)"`
	Url          string `gorm:"type:varchar(255)"`
	StartOn      int64
	StartedOn    int64
	DueOn        int64
	IsCompleted  bool
	IsStarted    bool
	CompletedOn  int64
}

func (TestrailMilestone20260123) TableName() string {
	return "_tool_testrail_milestones"
}

// TestrailUser20260123 - Archived model for migration
type TestrailUser20260123 struct {
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id           uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false"`
	Name         string `gorm:"type:varchar(255)"`
	Email        string `gorm:"type:varchar(255);index"`
	IsActive     bool
	RoleId       int
	Role         string `gorm:"type:varchar(100)"`
}

func (TestrailUser20260123) TableName() string {
	return "_tool_testrail_users"
}

// TestrailCaseField20260123 - Archived model for migration
type TestrailCaseField20260123 struct {
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id           uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false"`
	SystemName   string `gorm:"type:varchar(255);index"`
	Label        string `gorm:"type:varchar(255)"`
	Name         string `gorm:"type:varchar(255)"`
	Description  string `gorm:"type:text"`
	TypeId       int
	IsGlobal     bool
	IsActive     bool
	DisplayOrder int
	IncludeAll   bool
	TemplateIds  string `gorm:"type:json"`
	Configs      string `gorm:"type:json"`
}

func (TestrailCaseField20260123) TableName() string {
	return "_tool_testrail_case_fields"
}

// TestrailCaseType20260123 - Archived model for migration
type TestrailCaseType20260123 struct {
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id           uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false"`
	Name         string `gorm:"type:varchar(255)"`
	IsDefault    bool
}

func (TestrailCaseType20260123) TableName() string {
	return "_tool_testrail_case_types"
}

// TestrailPriority20260123 - Archived model for migration
type TestrailPriority20260123 struct {
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id           uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false"`
	Name         string `gorm:"type:varchar(255)"`
	ShortName    string `gorm:"type:varchar(50)"`
	Priority     int
	IsDefault    bool
}

func (TestrailPriority20260123) TableName() string {
	return "_tool_testrail_priorities"
}

// TestrailStatus20260123 - Archived model for migration
type TestrailStatus20260123 struct {
	ConnectionId uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id           uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false"`
	Name         string `gorm:"type:varchar(255)"`
	Label        string `gorm:"type:varchar(255)"`
	ColorDark    int
	ColorMedium  int
	ColorBright  int
	IsFinal      bool
	IsSystem     bool
	IsUntested   bool
}

func (TestrailStatus20260123) TableName() string {
	return "_tool_testrail_statuses"
}

func (*addEnterpriseEntities) Up(basicRes context.BasicRes) errors.Error {
	return migrationhelper.AutoMigrateTables(
		basicRes,
		&TestrailSuite20260123{},
		&TestrailSection20260123{},
		&TestrailPlan20260123{},
		&TestrailPlanEntry20260123{},
		&TestrailMilestone20260123{},
		&TestrailUser20260123{},
		&TestrailCaseField20260123{},
		&TestrailCaseType20260123{},
		&TestrailPriority20260123{},
		&TestrailStatus20260123{},
	)
}

func (*addEnterpriseEntities) Version() uint64 {
	return 20260123000001
}

func (*addEnterpriseEntities) Name() string {
	return "Testrail add enterprise entities (suites, sections, plans, milestones, users, case fields)"
}

var _ plugin.MigrationScript = (*addEnterpriseEntities)(nil)
