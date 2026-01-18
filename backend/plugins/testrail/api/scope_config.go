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

package api

import (
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
)

// @Summary create scope config
// @Description Create scope config for a connection
// @Tags plugins/testrail
// @Accept application/json
// @Param connectionId path int true "connection ID"
// @Param body body models.TestrailScopeConfig true "json body"
// @Success 200  {object} models.TestrailScopeConfig
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections/{connectionId}/scope-configs [POST]
func CreateScopeConfig(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return dsHelper.ScopeConfigApi.Post(input)
}

// @Summary update scope config
// @Description Update scope config for a connection
// @Tags plugins/testrail
// @Accept application/json
// @Param connectionId path int true "connection ID"
// @Param scopeConfigId path int true "scope config ID"
// @Param body body models.TestrailScopeConfig true "json body"
// @Success 200  {object} models.TestrailScopeConfig
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections/{connectionId}/scope-configs/{scopeConfigId} [PATCH]
func UpdateScopeConfig(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return dsHelper.ScopeConfigApi.Patch(input)
}

// @Summary get scope config
// @Description Get scope config for a connection
// @Tags plugins/testrail
// @Accept application/json
// @Param connectionId path int true "connection ID"
// @Param scopeConfigId path int true "scope config ID"
// @Success 200  {object} models.TestrailScopeConfig
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections/{connectionId}/scope-configs/{scopeConfigId} [GET]
func GetScopeConfig(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return dsHelper.ScopeConfigApi.GetDetail(input)
}

// @Summary list scope configs
// @Description List scope configs for a connection
// @Tags plugins/testrail
// @Accept application/json
// @Param connectionId path int true "connection ID"
// @Success 200  {object} []models.TestrailScopeConfig
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections/{connectionId}/scope-configs [GET]
func GetScopeConfigList(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return dsHelper.ScopeConfigApi.GetAll(input)
}

// @Summary delete scope config
// @Description Delete scope config for a connection
// @Tags plugins/testrail
// @Accept application/json
// @Param connectionId path int true "connection ID"
// @Param scopeConfigId path int true "scope config ID"
// @Success 200  {object} models.TestrailScopeConfig
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections/{connectionId}/scope-configs/{scopeConfigId} [DELETE]
func DeleteScopeConfig(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return dsHelper.ScopeConfigApi.Delete(input)
}
