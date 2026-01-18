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

package main

import (
	"github.com/apache/incubator-devlake/core/runner"
	"github.com/apache/incubator-devlake/plugins/testrail/impl"
	"github.com/spf13/cobra"
)

var PluginEntry impl.Testrail //nolint:gochecknoglobals

func main() {
	testrailCmd := &cobra.Command{Use: "testrail"}
	connectionId := testrailCmd.Flags().Uint64P("connection", "c", 1, "testrail connection id")
	projectId := testrailCmd.Flags().Uint64P("project", "p", 0, "testrail project id")
	timeAfter := testrailCmd.Flags().StringP("timeAfter", "a", "", "collect data that are created after specified time, ie 2006-01-02T15:04:05Z")
	_ = testrailCmd.MarkFlagRequired("connection")
	_ = testrailCmd.MarkFlagRequired("project")

	testrailCmd.Run = func(cmd *cobra.Command, args []string) {
		runner.DirectRun(cmd, args, PluginEntry, map[string]interface{}{
			"connectionId": *connectionId,
			"projectId":    *projectId,
		}, *timeAfter)
	}
	runner.RunCmd(testrailCmd)
}
