# Go API Template Walkthrough

Welcome to the Go API Template! This walkthrough will guide you through the project structure, common development tasks like migrations and testing, and how to extend the application with new features.

## 1. Project Structure

The project follows a **Clean Architecture** approach, separating concerns into distinct layers:

```
src/
├── handlers/       # HTTP Transport Layer (e.g., Echo handlers)
├── services/       # Business Logic Layer (Use Cases)
├── repositories/   # Data Access Layer (DB, External APIs)
├── models/         # Domain Models and Data Transfer Objects (DTOs)
├── config/         # Configuration (Env vars, dependencies)
├── clients/        # External clients (Postgres, Keycloak, etc.)
└── server/         # Server setup and routing
```

**Dependency Injection** is handled by `uber/fx`, which wires up the application modules in `main.go`.

## 2. Getting Started

### Prerequisites
- Go 1.23+
- Docker & Docker Compose
- Make

### Quick Start
To stand up the entire environment (Postgres, Keycloak, Migrations) and run the app:

```bash
make setup
make run
```

This will:
1.  Start Postgres and Keycloak in Docker containers.
2.  Run database migrations.
3.  Configure Keycloak (create realm, client, etc.).
4.  Start the Go API server.

## 3. Database Migrations

We use `golang-migrate` to manage database schema changes.

### Create a New Migration
To create a new pair of up/down migration files:

```bash
make migrate-create name=create_users_table
```
This generates files in the `migrations/` folder.

### Run Migrations
Apply all up migrations:
```bash
make migrate-up
```

Rollback the last migration:
```bash
make migrate-down
```

## 4. Testing

### Run All Tests
```bash
make test
```

### Mocking
We use `mockery` to generate mocks for testing interfaces. If you change an interface (e.g., in `repositories/`), regenerate the mocks:

```bash
make mock
```
*Note: Ensure you have mockery installed via `make mockery-install`.*

## 5. Adding a New Feature

Let's say you want to add a `Task` resource. Follow these steps:

### Step 1: Define the Model
Create `src/models/task/task.go`. Define your structs for the database and API responses.

### Step 2: Create the Repository
1.  Create `src/repositories/task/types.go` defining the `Repository` interface.
2.  Create `src/repositories/task/repository.go` implementing the interface using `sqlx`.
3.  Add the repository to `src/repositories/module.go` for DI.

### Step 3: Create the Service
1.  Create `src/services/task/types.go` defining the `Service` interface.
2.  Create `src/services/task/service.go` implementing the business logic.
3.  Add the service to `src/services/module.go` for DI.

### Step 4: Create the Handler
1.  Create `src/handlers/task/handler.go` with Echo handler functions.
2.  Add the handler to `src/handlers/module.go` for DI.

### Step 5: Register Routes
Update `src/server/routes.go` to register the new endpoints, applying any necessary middlewares (like Auth).

### Step 6: Add Migrations
```bash
make migrate-create name=create_tasks_table
```
Fill in the SQL to create the table.

## 6. Keycloak Integration

The template comes pre-configured with Keycloak for Authentication.
- **Local Dev**: `make setup` configures a local Keycloak instance.
- **Middleware**: `src/server/middlewares/auth.go` handles token validation.
- **Protecting Routes**: Use the auth middleware in `routes.go` to protect endpoints.
