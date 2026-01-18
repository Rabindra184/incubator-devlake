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
	"context"
	"net/http"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/testrail/models"
)

// @Summary test connection
// @Description Test functionality of an existing connection
// @Tags plugins/testrail
// @Accept application/json
// @Param connectionId path int true "connection ID"
// @Success 200  {object} shared.ApiBody "Success"
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections/{connectionId}/test [POST]
func TestExistingConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	connection, err := dsHelper.ConnApi.GetMergedConnection(input)
	if err != nil {
		return nil, errors.BadInput.Wrap(err, "find connection from db")
	}
	if err := helper.DecodeMapStruct(input.Body, connection, false); err != nil {
		return nil, err
	}
	testConnectionResult, testConnectionErr := testConnection(context.TODO(), connection.TestrailConn)
	if testConnectionErr != nil {
		return nil, plugin.WrapTestConnectionErrResp(basicRes, testConnectionErr)
	}
	return testConnectionResult, nil
}

// @Summary test connection
// @Description Test functionality of a new connection
// @Tags plugins/testrail
// @Accept application/json
// @Param body body models.TestrailConn true "json body"
// @Success 200  {object} shared.ApiBody "Success"
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/test [POST]
func TestConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	var connection models.TestrailConn
	err := helper.Decode(input.Body, &connection, vld)
	if err != nil {
		return nil, err
	}
	testConnectionResult, testConnectionErr := testConnection(context.TODO(), connection)
	if testConnectionErr != nil {
		return nil, plugin.WrapTestConnectionErrResp(basicRes, testConnectionErr)
	}
	return testConnectionResult, nil
}

// @Summary create connection
// @Description Create a new connection
// @Tags plugins/testrail
// @Accept application/json
// @Param body body models.TestrailConnection true "json body"
// @Success 200  {object} models.TestrailConnection
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections [POST]
func PostConnections(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return dsHelper.ConnApi.Post(input)
}

// @Summary list connections
// @Description List Testrail connections
// @Tags plugins/testrail
// @Accept application/json
// @Success 200  {object} []models.TestrailConnection
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections [GET]
func ListConnections(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return dsHelper.ConnApi.GetAll(input)
}

// @Summary get connection
// @Description Get a Testrail connection
// @Tags plugins/testrail
// @Accept application/json
// @Param connectionId path int true "connection ID"
// @Success 200  {object} models.TestrailConnection
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections/{connectionId} [GET]
func GetConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return dsHelper.ConnApi.GetDetail(input)
}

// @Summary update connection
// @Description Update a Testrail connection
// @Tags plugins/testrail
// @Accept application/json
// @Param connectionId path int true "connection ID"
// @Param body body models.TestrailConnection true "json body"
// @Success 200  {object} models.TestrailConnection
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections/{connectionId} [PATCH]
func PatchConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return dsHelper.ConnApi.Patch(input)
}

// @Summary delete connection
// @Description Delete a Testrail connection
// @Tags plugins/testrail
// @Accept application/json
// @Param connectionId path int true "connection ID"
// @Success 200  {object} models.TestrailConnection
// @Failure 400  {object} shared.ApiBody "Bad Request"
// @Failure 500  {object} shared.ApiBody "Internal Error"
// @Router /plugins/testrail/connections/{connectionId} [DELETE]
func DeleteConnection(input *plugin.ApiResourceInput) (*plugin.ApiResourceOutput, errors.Error) {
	return dsHelper.ConnApi.Delete(input)
}

func testConnection(ctx context.Context, connection models.TestrailConn) (*plugin.ApiResourceOutput, errors.Error) {
	apiClient, err := helper.NewApiClientFromConnection(ctx, basicRes, &connection)
	if err != nil {
		return nil, err
	}
	res, err := apiClient.Get("index.php?/api/v2/get_projects", nil, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.HttpStatus(res.StatusCode).New("failed to connect to testrail")
	}
	return &plugin.ApiResourceOutput{Body: nil, Status: http.StatusOK}, nil
}
