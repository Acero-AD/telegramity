# Makefile for Telegramity

# Binary name
BINARY_NAME=telegramity

# Build the application
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run:
	go run main.go

# Run with hot reload (requires air: go install github.com/cosmtrek/air@latest)
dev:
	air

# Test the application
test:
	go test ./...

# Test with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Clean build artifacts
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f coverage.out

# Install dependencies
deps:
	go mod tidy
	go mod download

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Help
help:
	@echo "Available commands:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  dev          - Run with hot reload (requires air)"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Install dependencies"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
	@echo "  install-tools- Install development tools"
	@echo "  help         - Show this help message"

.PHONY: build run dev test test-coverage clean deps fmt lint install-tools help 

.PHONY: build test clean install example

# Build the SDK
build:
	go build ./pkg/telegramity/

# Run tests
test:
	go test ./tests/unit/ -v

# Run integration tests
test-integration:
	go test ./tests/integration/ -v

# Run all tests
test-all:
	go test ./... -v

# Clean build artifacts
clean:
	go clean
	rm -rf build/ dist/

# Install dependencies
install:
	go mod download
	go mod tidy

# Run examples
example:
	go run cmd/example/global_singleton_example.go

# Run singleton example
singleton:
	go run cmd/example/singleton_example.go

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Generate documentation
docs:
	godoc -http=:6060

# Build for distribution
dist: clean
	mkdir -p dist
	go build -o dist/telegramity ./cmd/example/

# Help
help:
	@echo "Available commands:"
	@echo "  build      - Build the SDK"
	@echo "  test       - Run unit tests"
	@echo "  test-all   - Run all tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  install    - Install dependencies"
	@echo "  example    - Run basic example"
	@echo "  singleton  - Run singleton example"
	@echo "  fmt        - Format code"
	@echo "  lint       - Lint code"
	@echo "  docs       - Generate documentation"
	@echo "  dist       - Build for distribution"
	@echo "  help       - Show this help" 