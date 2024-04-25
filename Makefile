# Run the Go application
run:
	@go run cmd/main.go

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
