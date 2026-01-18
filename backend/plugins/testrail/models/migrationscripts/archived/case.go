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

package archived

import (
	"github.com/apache/incubator-devlake/core/models/common"
)

type TestrailCase struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"id"`
	ProjectId        uint64 `gorm:"index;type:BIGINT NOT NULL" json:"project_id"`
	SuiteId          uint64 `json:"suite_id"`
	SectionId        uint64 `json:"section_id"`
	Title            string `json:"title"`
	TypeId           int    `json:"type_id"`
	PriorityId       int    `json:"priority_id"`
	Estimate         string `json:"estimate"`
	CreatedOn        int64  `json:"created_on"`
	UpdatedOn        int64  `json:"updated_on"`
}

func (TestrailCase) TableName() string {
	return "_tool_testrail_cases"
}
