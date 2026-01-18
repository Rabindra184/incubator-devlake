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

package e2e

import (
	"testing"

	"github.com/apache/incubator-devlake/core/models/domainlayer/qa"
	"github.com/apache/incubator-devlake/helpers/e2ehelper"
	"github.com/apache/incubator-devlake/plugins/testrail/impl"
	"github.com/apache/incubator-devlake/plugins/testrail/models"
	"github.com/apache/incubator-devlake/plugins/testrail/tasks"
)

func TestProjectDataFlow(t *testing.T) {
	var plugin impl.Testrail
	dataflowTester := e2ehelper.NewDataFlowTester(t, "testrail", plugin)

	taskData := &tasks.TestrailTaskData{
		Options: &tasks.TestrailOptions{
			ConnectionId: 1,
			ProjectId:    1,
		},
	}

	// import raw data
	dataflowTester.ImportCsvIntoRawTable("./testdata/raw_project.csv", "_raw_testrail_projects")

	// verify extraction
	dataflowTester.FlushTabler(&models.TestrailProject{})
	dataflowTester.Subtask(tasks.ExtractProjectsMeta, taskData)
	dataflowTester.VerifyTable(
		models.TestrailProject{},
		"./snapshot/testrail_projects.csv",
		[]string{"connection_id", "id", "name"},
	)

	// verify conversion
	dataflowTester.FlushTabler(&qa.QaProject{})
	dataflowTester.Subtask(tasks.ConvertProjectsMeta, taskData)
	dataflowTester.VerifyTable(
		qa.QaProject{},
		"./snapshot/qa_projects.csv",
		[]string{"id", "name"},
	)
}
