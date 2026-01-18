# TestRail Plugin

## Summary
The TestRail plugin allows Apache DevLake to collect and enrich data from TestRail, a comprehensive test management tool. By integrating TestRail data into DevLake, users can gain insights into their testing processes, track test execution trends, and monitor quality metrics alongside other development data.

## Supported Versions
- TestRail API v2

## Data Entities Collected
The plugin currently collects the following entities:
- **Projects**: The top-level containers in TestRail.
- **Test Cases**: Individual test definitions including titles and metadata.
- **Runs**: Specific test execution cycles.
- **Results**: The outcome of each test case within a run (Passed, Failed, Blocked, etc.).

## Domain Layer Mapping
TestRail entities are mapped to DevLake's standard QA domain models:
- `Projects` -> `qa_projects`
- `Test Cases` -> `qa_test_cases`
- `Results` -> `qa_test_case_executions`

## Configuration
To use the TestRail plugin, you need to configure a connection with the following details:
- **Endpoint**: Your TestRail instance URL (e.g., `https://your-domain.testrail.io`).
- **Username**: Your TestRail account email.
- **Password / API Key**: Your TestRail password or ideally an API Key generated from your User Settings.

## API Endpoints
The plugin exposes standard DevLake API endpoints for:
- Connection management (`/plugins/testrail/connections`)
- Remote scope discovery (`/plugins/testrail/connections/{connectionId}/remote-scopes`)
- Scope management (`/plugins/testrail/connections/{connectionId}/scopes`)
- Pipeline planning (`MakeDataSourcePipelinePlanV200`)

## Grafana Dashboard
A sample dashboard is included in `grafana/dashboards/Testrail.json`. It provides:
- High-level stats (Total Projects, Cases, Runs).
- Success Rate monitoring.
- Daily Execution Trends.
- Project-level distribution of tests.

## Implementation Details
The plugin follows the standard DevLake collector-extractor-converter pattern:
1. **Collectors**: Request data from TestRail API endpoints using pagination.
2. **Extractors**: Store raw JSON responses into tool-layer tables.
3. **Converters**: Transform tool-layer data into the standardized domain layer for cross-plugin analysis.

---

For technical details on how to build, test, and contribute to this plugin, please see the [Development Guide](./DEVELOPMENT.md).
