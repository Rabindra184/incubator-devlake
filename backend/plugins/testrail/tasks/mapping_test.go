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

package tasks

import (
	"testing"

	"github.com/apache/incubator-devlake/plugins/testrail/models"
	"github.com/stretchr/testify/assert"
)

func TestMapStatusWithConfig(t *testing.T) {
	// Test without config
	assert.Equal(t, "SUCCESS", mapStatusWithConfig(1, nil))
	assert.Equal(t, "FAILED", mapStatusWithConfig(5, nil))
	assert.Equal(t, "BLOCKED", mapStatusWithConfig(2, nil))
	assert.Equal(t, "PENDING", mapStatusWithConfig(3, nil))
	assert.Equal(t, "PENDING", mapStatusWithConfig(99, nil))

	// Test with config
	config := &models.TestrailScopeConfig{
		StatusMappings: map[string]models.StatusMapping{
			"1":  {StandardStatus: "PASSED_WITH_RESERVATIONS"},
			"5":  {StandardStatus: "ABORTED"},
			"99": {StandardStatus: "CUSTOM_SUCCESS"},
		},
	}
	assert.Equal(t, "PASSED_WITH_RESERVATIONS", mapStatusWithConfig(1, config))
	assert.Equal(t, "ABORTED", mapStatusWithConfig(5, config))
	assert.Equal(t, "CUSTOM_SUCCESS", mapStatusWithConfig(99, config))

	// Fallback for missing key in config
	assert.Equal(t, "BLOCKED", mapStatusWithConfig(2, config))
}

func TestMapCaseType(t *testing.T) {
	// Test without config
	assert.Equal(t, "acceptance", mapCaseType(1, nil))
	assert.Equal(t, "regression", mapCaseType(9, nil))
	assert.Equal(t, "functional", mapCaseType(99, nil))

	// Test with config
	config := &models.TestrailScopeConfig{
		TypeMappings: map[string]models.TypeMapping{
			"1":  {StandardType: "USER_ACCEPTANCE"},
			"99": {StandardType: "EXPLORATORY"},
		},
	}
	assert.Equal(t, "USER_ACCEPTANCE", mapCaseType(1, config))
	assert.Equal(t, "EXPLORATORY", mapCaseType(99, config))

	// Fallback for missing key in config
	assert.Equal(t, "regression", mapCaseType(9, config))
}
