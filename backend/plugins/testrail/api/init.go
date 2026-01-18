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
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/testrail/models"
	"github.com/go-playground/validator/v10"
)

var vld *validator.Validate
var basicRes context.BasicRes

var dsHelper *api.DsHelper[models.TestrailConnection, models.TestrailProject, models.TestrailScopeConfig]
var raProxy *api.DsRemoteApiProxyHelper[models.TestrailConnection]
var raScopeList *api.DsRemoteApiScopeListHelper[models.TestrailConnection, models.TestrailProject, TestrailRemotePagination]
var raScopeSearch *api.DsRemoteApiScopeSearchHelper[models.TestrailConnection, models.TestrailProject]

func Init(br context.BasicRes, p plugin.PluginMeta) {
	vld = validator.New()
	basicRes = br
	dsHelper = api.NewDataSourceHelper[
		models.TestrailConnection, models.TestrailProject, models.TestrailScopeConfig,
	](
		br,
		p.Name(),
		[]string{"name"},
		func(c models.TestrailConnection) models.TestrailConnection {
			return c.Sanitize()
		},
		nil,
		nil,
	)
	raProxy = api.NewDsRemoteApiProxyHelper[models.TestrailConnection](dsHelper.ConnApi.ModelApiHelper)
	raScopeList = api.NewDsRemoteApiScopeListHelper[models.TestrailConnection, models.TestrailProject, TestrailRemotePagination](raProxy, listTestrailRemoteScopes)
	raScopeSearch = api.NewDsRemoteApiScopeSearchHelper[models.TestrailConnection, models.TestrailProject](raProxy, searchTestrailRemoteScopes)
}
