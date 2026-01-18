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
	"fmt"

	"github.com/apache/incubator-devlake/core/models/common"
	"github.com/apache/incubator-devlake/core/plugin"
)

type TestrailProject struct {
	common.Scope     `mapstructure:",squash"`
	Id               uint64              `gorm:"primaryKey;type:BIGINT NOT NULL;autoIncrement:false" json:"id"`
	Name             string              `gorm:"type:varchar(255)" json:"name"`
	Announcement     string              `json:"announcement"`
	ShowAnnouncement bool                `json:"show_announcement"`
	IsCompleted      bool                `json:"is_completed"`
	CompletedOn      *common.Iso8601Time `json:"completed_on"`
	SuiteMode        int                 `json:"suite_mode"`
	Url              string              `json:"url"`
}

func (TestrailProject) TableName() string {
	return "_tool_testrail_projects"
}

func (p TestrailProject) ScopeId() string {
	return fmt.Sprintf("%d", p.Id)
}

func (p TestrailProject) ScopeName() string {
	return p.Name
}

func (p TestrailProject) ScopeFullName() string {
	return p.Name
}

func (p TestrailProject) ScopeParams() interface{} {
	return &TestrailApiParams{
		ConnectionId: p.ConnectionId,
		ProjectId:    p.Id,
	}
}

type TestrailApiParams struct {
	ConnectionId uint64 `json:"connectionId"`
	ProjectId    uint64 `json:"projectId"`
}

var _ plugin.ToolLayerScope = (*TestrailProject)(nil)
