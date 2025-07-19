.PHONY: all deps run test swag-install swag mockery-install mock \
		postgres postgres-stop \
		migrate-install migrate-create migrate-up migrate-down migrate-force

all: deps run

deps: postgres
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/vektra/mockery/v2@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

run:
	go run main.go

test:
	go test -cover ./... 

swag:
	swag init -o .internal/docs

mock:
	mockery

# PostgreSQL commands
POSTGRES_IMAGE ?= postgres:15-alpine
POSTGRES_PORT ?= 5432
POSTGRES_USER ?= myuser
POSTGRES_PASSWORD ?= mypassword
POSTGRES_DB ?= mydb
POSTGRES_HOST ?= localhost
POSTGRES_CONNECTION_STRING ?= postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

postgres:
	@if [ "$$(docker ps -q -f name=^/postgres$$)" ]; then \
		echo "Postgres container is already running."; \
	else \
		if [ "$$(docker ps -aq -f name=^/postgres$$)" ]; then \
			echo "Postgres container exists but is stopped. Starting it..."; \
			docker start postgres; \
		else \
			echo "Starting PostgreSQL with configuration:"; \
			echo "-------------------------------------"; \
			echo "Image: $(POSTGRES_IMAGE)"; \
			echo "Port: $(POSTGRES_PORT)"; \
			echo "Username: $(POSTGRES_USER)"; \
			echo "Password: $(POSTGRES_PASSWORD)"; \
			echo "Application Database: $(POSTGRES_DB)"; \
			echo "-------------------------------------"; \
			docker run -d --name postgres \
				-p $(POSTGRES_PORT):5432 \
				-e POSTGRES_USER=$(POSTGRES_USER) \
				-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
				-e POSTGRES_DB=$(POSTGRES_DB) \
				-v postgres_data:/var/lib/postgresql/data \
				-v $(PWD)/migrations:/migrations \
				$(POSTGRES_IMAGE); \
		fi \
	fi

postgres-stop:
	docker stop postgres || true
	docker rm postgres || true

migrate-create:
	migrate create -ext sql -dir migrations -seq ${name}

migrate-up:
	migrate -path migrations -database "${POSTGRES_CONNECTION_STRING}" up

migrate-down:
	migrate -path migrations -database "${POSTGRES_CONNECTION_STRING}" down

migrate-force:
	migrate -path migrations -database "${POSTGRES_CONNECTION_STRING}" force ${version}

teardown: postgres-stop
	@echo "All services are down!"