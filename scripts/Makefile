# Makefile for goffice - Go SDK for Office Open XML documents
#
# NOTE: This Makefile is in the scripts/ directory.
# Run from project root: make -f scripts/Makefile <target>
# Or create a symlink: ln -s scripts/Makefile Makefile
#
# Usage:
#   make -f scripts/Makefile          - Run all targets (build, test, lint)
#   make -f scripts/Makefile build    - Build the project
#   make -f scripts/Makefile test     - Run all tests
#   make -f scripts/Makefile lint     - Run golangci-lint
#   make -f scripts/Makefile clean    - Clean build artifacts
#   make -f scripts/Makefile testdata - Generate test data files

.PHONY: all build test lint clean testdata help

# Default target
all: build test lint

# Build the project
build:
	@echo "Building goffice..."
	go build ./...

# Run all tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run golangci-lint
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	fi

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	go clean ./...
	rm -rf bin/
	rm -f testdata/fixtures/minimal.docx

# Generate test data files
testdata:
	@echo "Generating test data files..."
	go run ./cmd/generate-testdata

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Run tests with race detection
test-race:
	@echo "Running tests with race detection..."
	go test -race ./...

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Check if code needs formatting
fmt-check:
	@echo "Checking code format..."
	@test -z "$$(gofmt -l .)" || (echo "Code is not formatted. Run 'make fmt'" && exit 1)

# Tidy go.mod
tidy:
	@echo "Tidying go.mod..."
	go mod tidy

# Help
help:
	@echo "goffice Makefile targets:"
	@echo "  all       - Run build, test, and lint (default)"
	@echo "  build     - Build the project"
	@echo "  test      - Run all tests"
	@echo "  lint      - Run golangci-lint"
	@echo "  clean     - Clean build artifacts"
	@echo "  testdata  - Generate test data files"
	@echo "  coverage  - Run tests with coverage report"
	@echo "  test-race - Run tests with race detection"
	@echo "  fmt       - Format code"
	@echo "  fmt-check - Check if code needs formatting"
	@echo "  tidy      - Tidy go.mod"
	@echo "  help      - Show this help message"
