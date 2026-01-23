# TestRail Metrics & Configuration Guide

This document provides a comprehensive guide to the metrics supported by the TestRail plugin and how to configure them for optimal results.

## Metrics Catalog

The TestRail plugin enables a wide range of quality and process metrics by collecting data across projects, suites, sections, cases, runs, and results.

### 1. Velocity Metrics
*   **Test Case Creation Rate**: Measure how many new test cases are being defined per day/week/month. This helps track the progress of test design during new feature development.
*   **Execution Frequency**: Number of test runs initiated over time.

### 2. Quality Metrics
*   **Success Rate (Pass/Fail Ratio)**: The percentage of tests that passed in a given run or period.
*   **Flaky Test Analysis**: Identify tests that frequently flip between Pass and Fail without code changes (requires longitudinal analysis of results).
*   **Defect Density**: Correlation between failing test cases and specific modules/sections.

### 3. Coverage Metrics
*   **Functional Coverage**: Breakdown of test cases by Suite and Section to ensure all functional areas have adequate test depth.
*   **Automation Coverage**: Percentage of test cases that are automated vs. manual (requires custom field mapping).
*   **Requirement Coverage**: Mapping test cases to requirements/user stories via the `refs` field.

### 4. Resource & Productivity Metrics
*   **Tester Workload**: Distribution of test executions and result submissions among team members.
*   **Time to Execute**: Average duration of test runs and individual execution times.

---

## Configuration Guide

To unlock the full potential of these metrics, proper mapping is required in the **Scope Configuration**.

### Status Mappings
TestRail allows for custom result statuses. You must map these to standard DevLake statuses (`SUCCESS`, `FAILED`, `PENDING`, `BLOCKED`, `SKIPPED`) to ensure accurate success rate calculations.

**Example Configuration:**
```json
{
  "statusMappings": {
    "1": { "standardStatus": "SUCCESS" },
    "5": { "standardStatus": "FAILED" },
    "2": { "standardStatus": "BLOCKED" },
    "6": { "standardStatus": "SKIPPED" }
  }
}
```

### Type & Priority Mappings
Standardize your test categorization by mapping TestRail's internal IDs to standard labels.

**Example Configuration:**
```json
{
  "typeMappings": {
    "1": { "standardType": "acceptance" },
    "9": { "standardType": "regression" }
  },
  "priorityMappings": {
    "1": { "standardPriority": "low" },
    "4": { "standardPriority": "critical" }
  }
}
```

### Custom Field Mapping for Automation
If you track automation status in a custom field (e.g., `custom_automation_type`), configure it so the plugin can identify automated tests.

**Example Configuration:**
```json
{
  "automationStatusField": "custom_automation_type",
  "automatedValues": ["automated", "yes", "true"]
}
```

---

## Grafana Dashboard Highlights

The included TestRail dashboard provides several pre-built visualizations:

1.  **Overall Summary Stats**: Quick view of total projects, cases, and runs.
2.  **Success Rate Trend**: Real-time tracking of test pass rates.
3.  **Creation Velocity**: Monitor your test design pipeline.
4.  **Suite Distribution**: Visualize test coverage across different functional areas.
5.  **User Attribution**: Track who is executing tests and reporting results.

---

## Entity Coverage Matrix

| TestRail Entity | Tool Layer Support | Domain Layer Support | Notes |
|-----------------|-------------------|----------------------|-------|
| Project | ✅ Full | ✅ Full | Mapped to `qa_projects` |
| Suite | ✅ Full | ❌ Partial | Organization only |
| Section | ✅ Full | ❌ Partial | Organization only |
| Case | ✅ Full | ✅ Full | Mapped to `qa_test_cases` |
| Run | ✅ Full | ✅ Full | Mapped to `qa_test_runs` |
| Result | ✅ Full | ✅ Full | Mapped to `qa_test_case_executions`|
| Milestone | ✅ Full | ❌ None | Collected for context |
| User | ✅ Full | ✅ Full | Used for attribution |
| Custom Fields| ✅ Full | ❌ Metadata | Stored as JSON |

---

*Last Updated: 2026-01-23*
