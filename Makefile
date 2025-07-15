.PHONY: all deps run test mock

all: deps run

deps:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/vektra/mockery/v2@latest

run:
	go run main.go

test:
	go test -v -cover ./... 

swag:
	swag init -o .internal/docs

mock:
	mockery