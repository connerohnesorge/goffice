#!/usr/bin/env bash
# Simple E2E Test Runner
#
# This script provides a simplified interface for running E2E tests.
# It builds both generators and runs the test suite.

set -e

echo "Building Go generator..."
cd generators/go && go build -o e2e-go-generator . && cd ../..

echo "Building C# generator..."
cd generators/csharp && dotnet build -c Release && cd ../..

echo "Running E2E tests..."
go test -v -timeout 20m ./...

echo "Done!"
