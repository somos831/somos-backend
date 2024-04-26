# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

# Run the Go application
run: migrate-up
	@go run cmd/main.go

# Run database migrations
migrate-up:
	migrate -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -path "db/migrations" up

# Build the Go application into a binary
build: lint test
	@go build -o bin/myapp cmd/main.go

# Clean up generated files
clean:
	@rm -rf bin/

# Run tests
test:
	@go test ./...

# Lint the Go application
lint: format
	@golangci-lint run

# Format the code using go fmt
format:
	@go fmt ./...

# Install dependencies
deps:
	@go mod tidy
