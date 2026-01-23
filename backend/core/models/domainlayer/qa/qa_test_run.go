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

package qa

import (
	"time"

	"github.com/apache/incubator-devlake/core/models/domainlayer"
)

// QaTestRun represents a test run cycle in the domain layer
type QaTestRun struct {
	domainlayer.DomainEntityExtended
	QaProjectId  string     `gorm:"type:varchar(255);index;comment:Project ID"`
	Name         string     `gorm:"type:varchar(255);comment:Run name"`
	Description  string     `gorm:"type:text;comment:Description"`
	StartTime    *time.Time `gorm:"comment:Run start time"`
	FinishTime   *time.Time `gorm:"comment:Run finish time"`
	Status       string     `gorm:"type:varchar(255);comment:Run status | COMPLETED | IN_PROGRESS"`
	PassedCount  int        `gorm:"comment:Number of passed tests"`
	FailedCount  int        `gorm:"comment:Number of failed tests"`
	SkippedCount int        `gorm:"comment:Number of skipped tests"`
	TotalCount   int        `gorm:"comment:Total number of tests"`
}

func (QaTestRun) TableName() string {
	return "qa_test_runs"
}
