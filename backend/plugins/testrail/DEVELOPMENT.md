# TestRail Plugin Development & Testing Guide

This document provides instructions for developers who want to test the TestRail plugin locally or contribute to its development.

## Prerequisites

- **Go**: Version 1.20 or later.
- **Make**: Standard build tool (available via Xcode tools on Mac).
- **TestRail Account**: A trial or existing TestRail instance for live verification.
- **MySQL/PostgreSQL**: A local database instance for DevLake to connect to.

## Local Environment Setup

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/apache/incubator-devlake.git
    cd incubator-devlake
    ```

2.  **Configure `.env`**:
    Copy `backend/.env.example` to `backend/.env` and configure your database connection string (`DB_URL`).

## Running with Docker Compose

If you want to run the entire DevLake stack (Backend, Config-UI, MySQL, Grafana) with the new TestRail plugin:

1.  **Build and Start**:
    ```bash
    docker-compose -f docker-compose-dev.yml up --build
    ```
    *Note: The `--build` flag is required to ensuring your local TestRail plugin changes are compiled into the container.*

2.  **Access the Services**:
    - **Config-UI**: `http://localhost:4000`
    - **DevLake API**: `http://localhost:8080`
    - **Grafana Dashboard**: `http://localhost:3002` (Login: `admin/admin`)

## Building the Plugin (Direct)

To build only the TestRail plugin:
```bash
cd backend
DEVLAKE_PLUGINS=testrail make build-plugin
```
This compiles the plugin into `backend/bin/plugins/testrail/testrail.so`.

## Running Locally

To run the DevLake server with the TestRail plugin loaded:
```bash
cd backend
make dev
```
The server will be reachable at `http://localhost:8080`.

## Testing

### 1. Automated Tests (E2E & Unit)
We use DevLake's `DataFlowTester` to verify the data flow without a live API.

Run the tests in the plugin directory:
```bash
cd backend/plugins/testrail
go test ./e2e/...
```

### 2. Manual API Verification
You can use `curl` to test the plugin endpoints:

**Create a Connection:**
```bash
curl --request POST 'http://localhost:8080/plugins/testrail/connections' \
--header 'Content-Type: application/json' \
--data-raw '{
    "endpoint": "https://your-domain.testrail.io/",
    "username": "your-email@example.com",
    "password": "your-api-key",
    "name": "Local-TestRail"
}'
```

**Discovery (Listing Projects):**
```bash
# Replace 1 with your connection ID
curl 'http://localhost:8080/plugins/testrail/connections/1/remote-scopes'
```

### 3. Verification with Config-UI
1.  Run the Config-UI (requires Node.js):
    ```bash
    cd config-ui
    npm install
    npm run dev
    ```
2.  Open `http://localhost:4000`.
3.  Go to `Data Connections` -> `TestRail` and follow the setup wizard.

## Project Structure

- `api/`: REST API handlers for connection and scope management.
- `models/`: Database models for tool layer (raw) and archived migrations.
- `tasks/`:
  - `*_collector.go`: Handles API requests and pagination.
  - `*_extractor.go`: Parses raw JSON into tool-layer tables.
  - `*_converter.go`: Transforms tool-layer data into Domain Layer models.
- `e2e/`: End-to-end data flow tests.

## Troubleshooting

-   **Plugin not loaded**: Check `backend/bin/plugins/testrail/` to ensure the `.so` file exists.
-   **Database Errors**: Ensure your `DB_URL` in `.env` is correct and the database user has migration permissions.
-   **API Errors**: Verify that your TestRail "API" is enabled in your TestRail Site Settings (Site Settings -> API).
