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
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/apache/incubator-devlake/core/errors"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

// TestrailConn holds the essential information to connect to the Testrail API
type TestrailConn struct {
	helper.RestConnection `mapstructure:",squash"`
	Username              string `mapstructure:"username" validate:"required" json:"username"`
	Password              string `mapstructure:"password" validate:"required" json:"password" encrypt:"yes"`
}

func (connection TestrailConn) Sanitize() TestrailConn {
	connection.Password = ""
	return connection
}

// SetupAuthentication sets up the request headers for authentication
func (connection *TestrailConn) SetupAuthentication(req *http.Request) errors.Error {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", connection.Username, connection.Password)))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", auth))
	return nil
}

type TestrailConnection struct {
	helper.BaseConnection `mapstructure:",squash"`
	TestrailConn          `mapstructure:",squash"`
}

func (TestrailConnection) TableName() string {
	return "_tool_testrail_connections"
}

func (connection TestrailConnection) Sanitize() TestrailConnection {
	connection.TestrailConn = connection.TestrailConn.Sanitize()
	return connection
}

func (connection *TestrailConnection) MergeFromRequest(target *TestrailConnection, body map[string]interface{}) error {
	password := target.Password
	if err := helper.DecodeMapStruct(body, target, true); err != nil {
		return err
	}
	modifiedPassword := target.Password
	if modifiedPassword == "" {
		target.Password = password
	}
	return nil
}
