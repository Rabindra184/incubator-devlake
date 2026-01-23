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

package impl

import (
	"fmt"

	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	coreModels "github.com/apache/incubator-devlake/core/models"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/testrail/api"
	"github.com/apache/incubator-devlake/plugins/testrail/models"
	"github.com/apache/incubator-devlake/plugins/testrail/models/migrationscripts"
	"github.com/apache/incubator-devlake/plugins/testrail/tasks"
)

var _ interface {
	plugin.PluginMeta
	plugin.PluginInit
	plugin.PluginTask
	plugin.PluginApi
	plugin.PluginModel
	plugin.PluginMigration
	plugin.CloseablePluginTask
	plugin.PluginSource
	plugin.DataSourcePluginBlueprintV200
} = (*Testrail)(nil)

type Testrail struct{}

func (p Testrail) Init(basicRes context.BasicRes) errors.Error {
	api.Init(basicRes, p)
	return nil
}

func (p Testrail) Connection() dal.Tabler {
	return &models.TestrailConnection{}
}

func (p Testrail) Scope() plugin.ToolLayerScope {
	return &models.TestrailProject{}
}

func (p Testrail) ScopeConfig() dal.Tabler {
	return &models.TestrailScopeConfig{}
}

func (p Testrail) GetTablesInfo() []dal.Tabler {
	return []dal.Tabler{
		&models.TestrailConnection{},
		&models.TestrailProject{},
		&models.TestrailScopeConfig{},
		&models.TestrailCase{},
		&models.TestrailRun{},
		&models.TestrailResult{},
		&models.TestrailSuite{},
		&models.TestrailSection{},
		&models.TestrailPlan{},
		&models.TestrailPlanEntry{},
		&models.TestrailMilestone{},
		&models.TestrailUser{},
		&models.TestrailCaseField{},
		&models.TestrailCaseType{},
		&models.TestrailPriority{},
		&models.TestrailStatus{},
	}
}

func (p Testrail) Description() string {
	return "To collect and enrich data from Testrail"
}

func (p Testrail) Name() string {
	return "testrail"
}

func (p Testrail) SubTaskMetas() []plugin.SubTaskMeta {
	return []plugin.SubTaskMeta{
		// Metadata collectors - run first to get reference data
		tasks.CollectStatusesMeta,
		tasks.ExtractStatusesMeta,
		tasks.CollectPrioritiesMeta,
		tasks.ExtractPrioritiesMeta,
		tasks.CollectCaseTypesMeta,
		tasks.ExtractCaseTypesMeta,
		tasks.CollectCaseFieldsMeta,
		tasks.ExtractCaseFieldsMeta,
		tasks.CollectUsersMeta,
		tasks.ExtractUsersMeta,

		// Project and hierarchy collectors
		tasks.CollectProjectsMeta,
		tasks.ExtractProjectsMeta,
		tasks.ConvertProjectsMeta,
		tasks.CollectSuitesMeta,
		tasks.ExtractSuitesMeta,
		tasks.CollectSectionsMeta,
		tasks.ExtractSectionsMeta,
		tasks.CollectMilestonesMeta,
		tasks.ExtractMilestonesMeta,

		// Test case collectors
		tasks.CollectCasesMeta,
		tasks.ExtractCasesMeta,
		tasks.ConvertCasesMeta,

		// Test execution collectors
		tasks.CollectPlansMeta,
		tasks.ExtractPlansMeta,
		tasks.CollectRunsMeta,
		tasks.ExtractRunsMeta,
		tasks.ConvertRunsMeta,
		tasks.CollectResultsMeta,
		tasks.ExtractResultsMeta,
		tasks.ConvertResultsMeta,
	}
}

func (p Testrail) PrepareTaskData(taskCtx plugin.TaskContext, options map[string]interface{}) (interface{}, errors.Error) {
	op, err := tasks.DecodeTaskOptions(options)
	if err != nil {
		return nil, err
	}

	logger := taskCtx.GetLogger()
	logger.Debug("%v", options)

	connection := &models.TestrailConnection{}
	connectionHelper := helper.NewConnectionHelper(
		taskCtx,
		nil,
		p.Name(),
	)

	err = connectionHelper.FirstById(connection, op.ConnectionId)
	if err != nil {
		return nil, err
	}

	apiClient, err := tasks.CreateTestrailApiClient(taskCtx, connection)
	if err != nil {
		return nil, err
	}

	taskData := &tasks.TestrailTaskData{
		Options:   op,
		ApiClient: apiClient,
	}

	return taskData, nil
}

func (p Testrail) RootPkgPath() string {
	return "github.com/apache/incubator-devlake/plugins/testrail"
}

func (p Testrail) MigrationScripts() []plugin.MigrationScript {
	return migrationscripts.All()
}

func (p Testrail) ApiResources() map[string]map[string]plugin.ApiResourceHandler {
	return map[string]map[string]plugin.ApiResourceHandler{
		"test": {
			"POST": api.TestConnection,
		},
		"connections": {
			"POST": api.PostConnections,
			"GET":  api.ListConnections,
		},
		"connections/:connectionId": {
			"PATCH":  api.PatchConnection,
			"DELETE": api.DeleteConnection,
			"GET":    api.GetConnection,
		},
		"connections/:connectionId/test": {
			"POST": api.TestExistingConnection,
		},
		"connections/:connectionId/remote-scopes": {
			"GET": api.RemoteScopes,
		},
		"connections/:connectionId/search-remote-scopes": {
			"GET": api.SearchRemoteScopes,
		},
		"connections/:connectionId/scopes/:scopeId": {
			"GET":    api.GetScope,
			"PATCH":  api.PatchScope,
			"DELETE": api.DeleteScope,
		},
		"connections/:connectionId/scopes/:scopeId/latest-sync-state": {
			"GET": api.GetScopeLatestSyncState,
		},
		"connections/:connectionId/scopes": {
			"GET": api.GetScopeList,
			"PUT": api.PutScopes,
		},
		"connections/:connectionId/scope-configs": {
			"POST": api.CreateScopeConfig,
			"GET":  api.GetScopeConfigList,
		},
		"connections/:connectionId/scope-configs/:scopeConfigId": {
			"PATCH":  api.UpdateScopeConfig,
			"GET":    api.GetScopeConfig,
			"DELETE": api.DeleteScopeConfig,
		},
	}
}

func (p Testrail) Close(taskCtx plugin.TaskContext) errors.Error {
	data, ok := taskCtx.GetData().(*tasks.TestrailTaskData)
	if !ok {
		return errors.Default.New(fmt.Sprintf("GetData failed when try to close %+v", taskCtx))
	}
	data.ApiClient.Release()
	return nil
}

func (p Testrail) MakeDataSourcePipelinePlanV200(
	connectionId uint64,
	scopes []*coreModels.BlueprintScope,
) (coreModels.PipelinePlan, []plugin.Scope, errors.Error) {
	return api.MakeDataSourcePipelinePlanV200(p.SubTaskMetas(), connectionId, scopes)
}
