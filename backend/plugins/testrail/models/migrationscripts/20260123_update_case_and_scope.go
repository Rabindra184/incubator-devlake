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

type updateCaseAndScopeConfig struct{}

// TestrailCase20260123v2 - Updated case model with additional fields
type TestrailCase20260123v2 struct {
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false"`
	ProjectId        uint64 `gorm:"index;type:BIGINT NOT NULL"`
	SuiteId          uint64 `gorm:"index"`
	SectionId        uint64 `gorm:"index"`
	Title            string `gorm:"type:varchar(255)"`
	TypeId           int
	PriorityId       int
	TemplateId       int
	MilestoneId      uint64 `gorm:"index"`
	Refs             string `gorm:"type:varchar(255)"`
	Estimate         string `gorm:"type:varchar(50)"`
	EstimateForecast string `gorm:"type:varchar(50)"`
	Preconditions    string `gorm:"type:text"`
	CreatedBy        uint64 `gorm:"index"`
	CreatedOn        int64
	UpdatedBy        uint64
	UpdatedOn        int64
	CustomFields     string `gorm:"type:json"`
}

func (TestrailCase20260123v2) TableName() string {
	return "_tool_testrail_cases"
}

// TestrailScopeConfig20260123 - Updated scope config with mappings
type TestrailScopeConfig20260123 struct {
	StatusMappings        string `gorm:"type:json"`
	TypeMappings          string `gorm:"type:json"`
	PriorityMappings      string `gorm:"type:json"`
	CustomFieldMappings   string `gorm:"type:json"`
	AutomationStatusField string `gorm:"type:varchar(255)"`
	AutomatedValues       string `gorm:"type:json"`
	IncludeSuiteIds       string `gorm:"type:json"`
	ExcludeSuiteIds       string `gorm:"type:json"`
	IncludeMilestoneIds   string `gorm:"type:json"`
}

func (TestrailScopeConfig20260123) TableName() string {
	return "_tool_testrail_scope_configs"
}

type TestrailResult20260123 struct {
	CreatedBy uint64 `gorm:"index"`
}

func (TestrailResult20260123) TableName() string {
	return "_tool_testrail_results"
}

func (*updateCaseAndScopeConfig) Up(basicRes context.BasicRes) errors.Error {
	db := basicRes.GetDal()

	// Add new columns to cases table
	if err := migrationhelper.AutoMigrateTables(basicRes, &TestrailCase20260123v2{}); err != nil {
		return err
	}

	// Add new columns to scope_configs table
	if err := db.AutoMigrate(&TestrailScopeConfig20260123{}); err != nil {
		return errors.Convert(err)
	}

	// Add new columns to results table
	if err := migrationhelper.AutoMigrateTables(basicRes, &TestrailResult20260123{}); err != nil {
		return err
	}

	return nil
}

func (*updateCaseAndScopeConfig) Version() uint64 {
	return 20260123000002
}

func (*updateCaseAndScopeConfig) Name() string {
	return "Testrail update case with custom fields and enhance scope config"
}

var _ plugin.MigrationScript = (*updateCaseAndScopeConfig)(nil)
