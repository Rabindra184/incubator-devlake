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

// TestrailSuite represents a test suite in TestRail
// In multi-suite mode, a project can have multiple suites, each containing sections and cases
type TestrailSuite struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	ProjectId        uint64 `gorm:"index;type:BIGINT NOT NULL" json:"project_id"`
	Name             string `gorm:"type:varchar(255)" json:"name"`
	Description      string `gorm:"type:text" json:"description"`
	Url              string `gorm:"type:varchar(255)" json:"url"`
	IsBaseline       bool   `json:"is_baseline"`
	IsMaster         bool   `json:"is_master"`
	IsCompleted      bool   `json:"is_completed"`
	CompletedOn      int64  `json:"completed_on"`
}

func (TestrailSuite) TableName() string {
	return "_tool_testrail_suites"
}
