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
	"net/http"
	"net/url"
	"strconv"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	dsmodels "github.com/apache/incubator-devlake/helpers/pluginhelper/api/models"
	"github.com/apache/incubator-devlake/plugins/testrail/models"
)

type TestrailRemotePagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ProjectResponse struct {
	Offset   int                      `json:"offset"`
	Limit    int                      `json:"limit"`
	Size     int                      `json:"size"`
	Projects []models.TestrailProject `json:"projects"`
	Links    struct {
		Next string `json:"next"`
		Prev string `json:"prev"`
	} `json:"_links"`
}

func queryTestrailRemoteScopes(
	apiClient plugin.ApiClient,
	page TestrailRemotePagination,
) (
	children []dsmodels.DsRemoteApiScopeListEntry[models.TestrailProject],
	nextPage *TestrailRemotePagination,
	err errors.Error,
) {
	if page.Limit == 0 {
		page.Limit = 100
	}
	var res *http.Response
	queryParams := url.Values{
		"offset": {strconv.Itoa(page.Offset)},
		"limit":  {strconv.Itoa(page.Limit)},
	}

	res, err = apiClient.Get("index.php?/api/v2/get_projects", queryParams, nil)
	if err != nil {
		return
	}

	response := &ProjectResponse{}
	err = api.UnmarshalResponse(res, response)
	if err != nil {
		return
	}

	for i := range response.Projects {
		project := response.Projects[i]
		children = append(children, dsmodels.DsRemoteApiScopeListEntry[models.TestrailProject]{
			Type:     api.RAS_ENTRY_TYPE_SCOPE,
			Id:       strconv.FormatUint(project.Id, 10),
			Name:     project.Name,
			FullName: project.Name,
			Data:     &project,
		})
	}

	if response.Links.Next != "" {
		nextPage = &TestrailRemotePagination{
			Offset: page.Offset + page.Limit,
			Limit:  page.Limit,
		}
	}

	return
}

func listTestrailRemoteScopes(
	connection *models.TestrailConnection,
	apiClient plugin.ApiClient,
	groupId string,
	page TestrailRemotePagination,
) (
	[]dsmodels.DsRemoteApiScopeListEntry[models.TestrailProject],
	*TestrailRemotePagination,
	errors.Error,
) {
	return queryTestrailRemoteScopes(apiClient, page)
}

func searchTestrailRemoteScopes(
	apiClient plugin.ApiClient,
	params *dsmodels.DsRemoteApiScopeSearchParams,
) (
	children []dsmodels.DsRemoteApiScopeListEntry[models.TestrailProject],
	err errors.Error,
) {
	// TestRail doesn't have a search API for projects, so we list and filter locally or just return list
	// For simplicity, we just list first page
	children, _, err = queryTestrailRemoteScopes(apiClient, TestrailRemotePagination{
		Offset: 0,
		Limit:  params.PageSize,
	})
	return
}

func RemoteScopes(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return raScopeList.Get(input)
}

func SearchRemoteScopes(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return raScopeSearch.Get(input)
}
