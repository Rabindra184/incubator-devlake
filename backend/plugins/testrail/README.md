# TestRail Plugin

## Summary
The TestRail plugin allows Apache DevLake to collect and enrich data from TestRail, a comprehensive test management tool. By integrating TestRail data into DevLake, users can gain insights into their testing processes, track test execution trends, and monitor quality metrics alongside other development data.

## Supported Versions
- TestRail API v2

## Data Entities Collected

The plugin collects the following entities from TestRail:

### Core Entities
| Entity | Description | API Endpoint |
|--------|-------------|--------------|
| **Projects** | Top-level containers in TestRail | `get_projects` |
| **Suites** | Test suites (for multi-suite projects) | `get_suites/{project_id}` |
| **Sections** | Hierarchical organization of test cases | `get_sections/{project_id}` |
| **Cases** | Individual test case definitions | `get_cases/{project_id}` |
| **Runs** | Test execution cycles | `get_runs/{project_id}` |
| **Plans** | Test plan groupings | `get_plans/{project_id}` |
| **Results** | Test execution results | `get_results_for_run/{run_id}` |
| **Milestones** | Release/sprint milestones | `get_milestones/{project_id}` |

### Metadata Entities
| Entity | Description | API Endpoint |
|--------|-------------|--------------|
| **Users** | TestRail users (testers, assignees) | `get_users` |
| **Case Fields** | Custom field definitions | `get_case_fields` |
| **Case Types** | Test case type definitions | `get_case_types` |
| **Priorities** | Priority level definitions | `get_priorities` |
| **Statuses** | Result status definitions | `get_statuses` |

## Custom Field Support

The plugin supports TestRail custom fields (`custom_*` fields):
- Custom field definitions are collected and stored
- Custom field values are extracted from test cases and stored as JSON
- Scope configuration allows mapping custom fields to standard analyses

## Domain Layer Mapping

TestRail entities are mapped to DevLake's standard QA domain models:

| Tool Layer | Domain Layer |
|------------|--------------|
| `_tool_testrail_projects` | `qa_projects` |
| `_tool_testrail_cases` | `qa_test_cases` |
| `_tool_testrail_results` | `qa_test_case_executions` |

## Configuration

### Connection Settings

To use the TestRail plugin, configure a connection with:
- **Endpoint**: Your TestRail instance URL (e.g., `https://your-domain.testrail.io`)
- **Username**: Your TestRail account email
- **Password / API Key**: Your TestRail password or API Key (recommended)

### Scope Configuration

The plugin supports rich scope configuration:

```json
{
  "statusMappings": {
    "1": { "standardStatus": "SUCCESS" },
    "2": { "standardStatus": "BLOCKED" },
    "5": { "standardStatus": "FAILED" }
  },
  "typeMappings": {
    "1": { "standardType": "functional" },
    "3": { "standardType": "automated" }
  },
  "priorityMappings": {
    "1": { "standardPriority": "low" },
    "4": { "standardPriority": "critical" }
  },
  "automationStatusField": "custom_automation_type",
  "automatedValues": ["automated", "yes"],
  "includeSuiteIds": [],
  "excludeSuiteIds": []
}
```

## API Endpoints

The plugin exposes standard DevLake API endpoints:

| Endpoint | Methods | Description |
|----------|---------|-------------|
| `/plugins/testrail/connections` | GET, POST | Connection management |
| `/plugins/testrail/connections/:id` | GET, PATCH, DELETE | Single connection |
| `/plugins/testrail/connections/:id/test` | POST | Test connection |
| `/plugins/testrail/connections/:id/remote-scopes` | GET | Discover projects |
| `/plugins/testrail/connections/:id/scopes` | GET, PUT | Scope management |
| `/plugins/testrail/connections/:id/scope-configs` | GET, POST | Configuration |

## Grafana Dashboard

A sample dashboard is included in `grafana/dashboards/Testrail.json` providing:
- High-level stats (Total Projects, Cases, Runs)
- Success Rate monitoring
- Daily Execution Trends
- Project-level distribution of tests

## Subtasks

The plugin executes the following subtasks in order:

### Metadata Collection
1. `collectStatuses` → `extractStatuses`
2. `collectPriorities` → `extractPriorities`
3. `collectCaseTypes` → `extractCaseTypes`
4. `collectCaseFields` → `extractCaseFields`
5. `collectUsers` → `extractUsers`

### Project & Hierarchy
6. `collectProjects` → `extractProjects` → `convertProjects`
7. `collectSuites` → `extractSuites`
8. `collectSections` → `extractSections`
9. `collectMilestones` → `extractMilestones`

### Test Cases
10. `collectCases` → `extractCases` → `convertCases`

### Test Execution
11. `collectPlans` → `extractPlans`
12. `collectRuns` → `extractRuns`
13. `collectResults` → `extractResults` → `convertResults`

## Implementation Details

The plugin follows the standard DevLake collector-extractor-converter pattern:

1. **Collectors**: Request data from TestRail API endpoints with pagination support
2. **Extractors**: Parse raw JSON responses into tool-layer tables
3. **Converters**: Transform tool-layer data into the standardized domain layer

### Database Tables

| Table | Description |
|-------|-------------|
| `_tool_testrail_connections` | Connection configurations |
| `_tool_testrail_projects` | TestRail projects |
| `_tool_testrail_suites` | Test suites |
| `_tool_testrail_sections` | Test sections |
| `_tool_testrail_cases` | Test cases with custom fields |
| `_tool_testrail_runs` | Test runs |
| `_tool_testrail_plans` | Test plans |
| `_tool_testrail_results` | Test results |
| `_tool_testrail_milestones` | Milestones |
| `_tool_testrail_users` | Users |
| `_tool_testrail_case_fields` | Custom field definitions |
| `_tool_testrail_case_types` | Case type definitions |
| `_tool_testrail_priorities` | Priority definitions |
| `_tool_testrail_statuses` | Status definitions |
| `_tool_testrail_scope_configs` | Scope configurations |

---

For technical details on how to build, test, and contribute to this plugin, please see the [Development Guide](./DEVELOPMENT.md).
