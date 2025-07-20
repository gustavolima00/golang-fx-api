# Go Sample API

This project is an example of a Go API using the `fx` framework. It demonstrates how to set up a simple API with health check endpoints, Swagger documentation, and mock services for testing.

## Tools Used

- **Go**: The main programming language used for this project.
- **Fx**: A dependency injection framework for Go.
- **Echo**: A high-performance, extensible, minimalist web framework for Go.
- **Swagger**: Used for API documentation.
- **Mockery**: A mock code autogenerator for Golang.
- **PostgreSQL**: The database used for this project.
- **Docker**: Used to containerize the application and run the database.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Running the API](#running-the-api)
- [Database Migrations](#database-migrations)
- [Running Tests](#running-tests)
- [Updating Swagger Docs](#updating-swagger-docs)
- [Updating Mock Files](#updating-mock-files)
- [Makefile Commands](#makefile-commands)

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [make](https://www.gnu.org/software/make/)
- [migrate](https://github.com/golang-migrate/migrate)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gustavolima00/sample-go-fx-api.git go-api
    cd go-api
    ```

2. Install dependencies:
    ```sh
    make deps
    ```

## Running the API

To run the API, use the following command:
```sh
make run
```

The API will be available at `http://localhost:8080`.

## Database Migrations

To run database migrations, use the following commands:

- `make migrate-up`: Apply all available migrations.
- `make migrate-down`: Roll back the last migration.
- `make migrate-create name=migration_name`: Create a new migration file.
- `make migrate-force version=version_number`: Force a specific migration version.

## Running Tests

To run the tests, use the following command:
```sh
make test
```

## Updating Swagger Docs

To update the Swagger documentation, run the following command:

```bash
make swag
```

This will generate a new `docs` directory with the updated Swagger files.

## Updating Mock Files

To update the mock files, run the following command:

```bash
make mock
```

## Makefile Commands

This project uses a `Makefile` to automate common tasks. Here are some of the most useful commands:

- `make all`: Install dependencies and run the application.
- `make deps`: Install all dependencies.
- `make run`: Run the application.
- `make test`: Run all tests.
- `make swag`: Generate Swagger documentation.
- `make mock`: Generate mock files.
- `make postgres`: Start the PostgreSQL container.
- `make postgres-stop`: Stop the PostgreSQL container.
- `make migrate-up`: Apply all available migrations.
- `make migrate-down`: Roll back the last migration.
- `make migrate-create name=migration_name`: Create a new migration file.
- `make migrate-force version=version_number`: Force a specific migration version.
- `make teardown`: Stop all services.