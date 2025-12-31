#!/usr/bin/env bash
# E2E Visual Test Runner Script
#
# This script orchestrates the end-to-end visual testing workflow:
# 1. Build Go and C# generators
# 2. Discover and execute test cases
# 3. Generate HTML reports
# 4. Exit with appropriate status code

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
E2E_DIR="$(dirname "$SCRIPT_DIR")"
REPO_ROOT="$(cd "$E2E_DIR/../.." && pwd)"

# Configuration
VERBOSE=${VERBOSE:-0}
CLEAN=${CLEAN:-0}
CATEGORY=${CATEGORY:-""}
TAG=${TAG:-""}

# Helper functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Parse command line arguments
usage() {
    cat <<EOF
E2E Visual Test Runner

Usage: $0 [OPTIONS]

Options:
    -h, --help          Show this help message
    -v, --verbose       Enable verbose output
    -c, --clean         Clean output directories before running
    --category CATEGORY Filter tests by category (chart, shape, text, etc.)
    --tag TAG           Filter tests by tag
    --skip-build        Skip building generators
    --skip-report       Skip generating HTML report

Examples:
    $0                          # Run all tests
    $0 --category chart         # Run only chart tests
    $0 --tag basic --verbose    # Run tests tagged with 'basic' in verbose mode
    $0 --clean                  # Clean and run all tests

EOF
    exit 0
}

SKIP_BUILD=0
SKIP_REPORT=0

while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            usage
            ;;
        -v|--verbose)
            VERBOSE=1
            shift
            ;;
        -c|--clean)
            CLEAN=1
            shift
            ;;
        --category)
            CATEGORY="$2"
            shift 2
            ;;
        --tag)
            TAG="$2"
            shift 2
            ;;
        --skip-build)
            SKIP_BUILD=1
            shift
            ;;
        --skip-report)
            SKIP_REPORT=1
            shift
            ;;
        *)
            log_error "Unknown option: $1"
            usage
            ;;
    esac
done

# Change to E2E directory
cd "$E2E_DIR"

log_info "Starting E2E Visual Testing"
log_info "E2E Directory: $E2E_DIR"
log_info "Repository Root: $REPO_ROOT"

# Clean output directories if requested
if [[ $CLEAN -eq 1 ]]; then
    log_info "Cleaning output directories..."
    rm -rf output/* reports/*
    log_success "Output directories cleaned"
fi

# Ensure output directories exist
mkdir -p output/{pptx/{go,csharp},pdf,png,diffs}
mkdir -p reports/{assets,diffs}

# Build Go generator
if [[ $SKIP_BUILD -eq 0 ]]; then
    log_info "Building Go generator..."
    cd generators/go
    if [[ $VERBOSE -eq 1 ]]; then
        go build -v -o e2e-go-generator .
    else
        go build -o e2e-go-generator . 2>&1 | grep -v "^#" || true
    fi
    log_success "Go generator built: generators/go/e2e-go-generator"
    cd "$E2E_DIR"

    # Build C# generator
    log_info "Building C# generator..."
    cd generators/csharp
    if [[ $VERBOSE -eq 1 ]]; then
        dotnet build -c Release
    else
        dotnet build -c Release --verbosity quiet
    fi
    log_success "C# generator built: generators/csharp/bin/Release/net9.0/"
    cd "$E2E_DIR"
else
    log_info "Skipping generator build (--skip-build)"
fi

# Check that generators exist
if [[ ! -f "generators/go/e2e-go-generator" ]]; then
    log_error "Go generator not found. Run without --skip-build or build manually."
    exit 1
fi

if [[ ! -d "generators/csharp/bin/Release/net9.0" ]]; then
    log_error "C# generator not found. Run without --skip-build or build manually."
    exit 1
fi

# Run Go tests
log_info "Running E2E tests..."

TEST_ARGS=()
if [[ -n "$CATEGORY" ]]; then
    TEST_ARGS+=("-run" "Test${CATEGORY^}")
    log_info "Filtering by category: $CATEGORY"
fi

if [[ $VERBOSE -eq 1 ]]; then
    TEST_ARGS+=("-v")
fi

# Run tests with gotestsum if available, otherwise use go test
if command -v gotestsum &> /dev/null; then
    if gotestsum --format short-verbose "${TEST_ARGS[@]}" ./... ; then
        TEST_RESULT=0
    else
        TEST_RESULT=$?
    fi
else
    if go test "${TEST_ARGS[@]}" ./... ; then
        TEST_RESULT=0
    else
        TEST_RESULT=$?
    fi
fi

# Generate HTML report
if [[ $SKIP_REPORT -eq 0 ]]; then
    log_info "Generating HTML report..."
    if [[ -f "reports/index.html" ]]; then
        log_success "HTML report generated: reports/index.html"
        log_info "View report: file://$(pwd)/reports/index.html"
    else
        log_warn "HTML report not found (tests may have been skipped or failed)"
    fi
else
    log_info "Skipping HTML report generation (--skip-report)"
fi

# Summary
echo ""
log_info "========================================="
if [[ $TEST_RESULT -eq 0 ]]; then
    log_success "All E2E tests passed!"
    log_info "========================================="
    exit 0
else
    log_error "Some E2E tests failed (exit code: $TEST_RESULT)"
    log_info "========================================="
    log_info "Check reports/index.html for details"
    exit $TEST_RESULT
fi
