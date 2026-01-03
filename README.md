# Go API Template

A robust, production-ready Go API template using **Echo**, **Uber Fx**, **Postgres**, and **Keycloak**. 
Designed with **Clean Architecture** principles to help you kickstart your next project.

## 🚀 Key Features

*   **Clean Architecture**: Separation of concerns (Handlers, Services, Repositories).
*   **Dependency Injection**: Powered by [uber/fx](https://github.com/uber-go/fx).
*   **Database**: Postgres driver with [sqlx](https://github.com/jmoiron/sqlx).
*   **Migrations**: Managed via `golang-migrate`.
*   **Authentication**: Keycloak integration for secure access.
*   **Documentation**: Swagger/OpenAPI auto-generation.
*   **Testing**: Ready-to-go setup with `testify` and `mockery`.

## 📚 Documentation

Detailed guides can be found in [WALKTHROUGH.md](WALKTHROUGH.md):
*   [Project Structure](WALKTHROUGH.md#1-project-structure)
*   [Database Migrations](WALKTHROUGH.md#3-database-migrations)
*   [Adding a New Feature](WALKTHROUGH.md#5-adding-a-new-feature)

## 🛠️ Quick Start

### Prerequisites
*   Go 1.23+
*   Docker & Docker Compose
*   Make

### Setup & Run
One command to rule them all (starts DBs, Keycloak, runs migrations, configured Auth):
```bash
make setup
```

Then start the application:
```bash
make run
```
The API will be available at `http://localhost:8080`.

## 📦 Using this Template

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/gustavolima00/sample-go-fx-api.git my-new-app
    cd my-new-app
    ```

2.  **Rename the Module**:
    ```bash
    go mod edit -module github.com/myuser/my-new-app
    # Find and replace imports in the codebase
    grep -r "github.com/gustavolima00/sample-go-fx-api" .
    ```

## 🧪 Testing

Run all tests:
```bash
make test
```

Generate mocks after interface changes:
```bash
make mock
```

## 📝 API Documentation

After running the app, access Swagger docs at:
```
http://localhost:8080/swagger/index.html
```




