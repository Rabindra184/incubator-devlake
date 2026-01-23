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

// TestrailUser represents a user in TestRail
// Users can be assignees, creators, or testers for various entities
type TestrailUser struct {
	common.NoPKModel `mapstructure:",squash"`
	ConnectionId     uint64 `gorm:"primaryKey;type:BIGINT NOT NULL" json:"connectionId"`
	Id               uint64 `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	Name             string `gorm:"type:varchar(255)" json:"name"`
	Email            string `gorm:"type:varchar(255);index" json:"email"`
	IsActive         bool   `json:"is_active"`
	RoleId           int    `json:"role_id"`
	Role             string `gorm:"type:varchar(100)" json:"role"`
}

func (TestrailUser) TableName() string {
	return "_tool_testrail_users"
}
