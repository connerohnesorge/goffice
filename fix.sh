#!/usr/bin/env bash

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
MAGENTA='\033[0;35m'
NC='\033[0m' # No Color

# Configuration
DRY_RUN=false
FILTER_LINTER=""
FILTER_SEVERITY=""
FILTER_FILE=""
MAX_ISSUES=0
DELAY_BETWEEN_CALLS=1

# Function to print colored messages
log_info() {
    echo -e "${BLUE}[INFO]${NC} $*" >&2
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $*" >&2
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $*" >&2
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $*" >&2
}

log_debug() {
    echo -e "${MAGENTA}[DEBUG]${NC} $*" >&2
}

# Function to show usage
show_usage() {
    cat << EOF
Usage: $0 [OPTIONS]

A script to automatically fix golangci-lint issues using clds.

OPTIONS:
    -h, --help              Show this help message
    -d, --dry-run           Show what would be done without calling clds
    -l, --linter LINTER     Filter issues by specific linter (e.g., errcheck, revive)
    -s, --severity LEVEL    Filter issues by severity (debug, error, warning)
    -f, --file PATTERN      Filter issues by file pattern (e.g., "pkg/nar/*.go")
    -m, --max-issues N      Process at most N issues (default: all)
    -w, --wait SECONDS      Delay between clds calls in seconds (default: 1)
    --list-linters          List all linters found in the output and exit
    --list-files            List all files with issues and exit

EXAMPLES:
    # Process all issues
    $0

    # Dry run to see what would be fixed
    $0 --dry-run

    # Fix only errcheck issues
    $0 --linter errcheck

    # Fix only error-severity issues in pkg/nar/
    $0 --severity error --file "pkg/nar/*"

    # Process first 5 issues with 2 second delay
    $0 --max-issues 5 --wait 2


EOF
}

# Parse command line arguments
parse_args() {
    while [[ $# -gt 0 ]]; do
        case $1 in
            -h|--help)
                show_usage
                exit 0
                ;;
            -d|--dry-run)
                DRY_RUN=true
                shift
                ;;
            -l|--linter)
                FILTER_LINTER="$2"
                shift 2
                ;;
            -s|--severity)
                FILTER_SEVERITY="$2"
                shift 2
                ;;
            -f|--file)
                FILTER_FILE="$2"
                shift 2
                ;;
            -m|--max-issues)
                MAX_ISSUES="$2"
                shift 2
                ;;
            -w|--wait)
                DELAY_BETWEEN_CALLS="$2"
                shift 2
                ;;
            --list-linters)
                list_linters
                exit 0
                ;;
            --list-files)
                list_files
                exit 0
                ;;
            *)
                log_error "Unknown option: $1"
                show_usage
                exit 1
                ;;
        esac
    done
}

# Check dependencies
check_dependencies() {
    if ! command -v jq &> /dev/null; then
        log_error "jq is required but not installed. Please install jq first."
        exit 1
    fi
    
    if ! command -v clds &> /dev/null && [ "$DRY_RUN" = false ]; then
        log_error "clds is required but not found in PATH."
        exit 1
    fi
}

# Get golangci-lint output (FIRST LINE only)
get_lint_output() {
    log_info "Running golangci-lint..."
    
    local lint_output=""
    
    # Try standard JSON output format and get first line
    if ! lint_output=$(golangci-lint run --out-format json:stdout 2>&1 | head -n 1); then
        log_info "golangci-lint completed with issues found"
    fi
    
    # If the above doesn't produce valid JSON, try alternative format
    if [ -z "$lint_output" ] || ! echo "$lint_output" | jq empty 2>/dev/null; then
        log_info "Trying alternative command format..."
        if ! lint_output=$(golangci-lint run --output.json.path stdout 2>&1 | head -n 1); then
            log_info "golangci-lint completed"
        fi
    fi
    
    if ! echo "$lint_output" | jq empty 2>/dev/null; then
        log_error "Failed to parse golangci-lint output as JSON"
        log_error "First line of output: $lint_output"
        exit 1
    fi
    
    echo "$lint_output"
}

# List all linters in the output
list_linters() {
    local output=$(get_lint_output)
    log_info "Linters found in output:"
    while IFS= read -r linter; do
        local count=$(echo "$output" | jq "[.Issues[] | select(.FromLinter == \"$linter\")] | length")
        echo "  - $linter ($count issues)" >&2
    done < <(echo "$output" | jq -r '.Issues[].FromLinter' | sort -u)
}

# List all files with issues
list_files() {
    local output=$(get_lint_output)
    log_info "Files with issues:"
    while IFS= read -r file; do
        local count=$(echo "$output" | jq "[.Issues[] | select(.Pos.Filename == \"$file\")] | length")
        echo "  - $file ($count issues)" >&2
    done < <(echo "$output" | jq -r '.Issues[].Pos.Filename' | sort -u)
}

# Filter issues based on criteria
filter_issues() {
    local issues="$1"
    
    # Apply linter filter
    if [ -n "$FILTER_LINTER" ]; then
        log_info "Filtering by linter: $FILTER_LINTER"
        issues=$(echo "$issues" | jq --arg linter "$FILTER_LINTER" '.Issues = [.Issues[] | select(.FromLinter == $linter)]')
    fi
    
    # Apply severity filter
    if [ -n "$FILTER_SEVERITY" ]; then
        log_info "Filtering by severity: $FILTER_SEVERITY"
        issues=$(echo "$issues" | jq --arg severity "$FILTER_SEVERITY" '.Issues = [.Issues[] | select(.Severity == $severity)]')
    fi
    
    # Apply file filter
    if [ -n "$FILTER_FILE" ]; then
        log_info "Filtering by file pattern: $FILTER_FILE"
        # Convert shell glob to regex
        local pattern="${FILTER_FILE//\*/.*}"
        issues=$(echo "$issues" | jq --arg pattern "$pattern" '.Issues = [.Issues[] | select(.Pos.Filename | test($pattern))]')
    fi
    
    # Apply max issues limit
    if [ "$MAX_ISSUES" -gt 0 ]; then
        log_info "Limiting to first $MAX_ISSUES issues"
        issues=$(echo "$issues" | jq --argjson max "$MAX_ISSUES" '.Issues = .Issues[:$max]')
    fi
    
    echo "$issues"
}

# Sanitize problem text for shell command
sanitize_problem() {
    local text="$1"
    text=$(echo "$text" | tr '\n' ' ')
    text="${text//\"/\\\"}"
    text=$(echo "$text" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')
    echo "$text"
}

# Process a single issue
process_issue() {
    local issue_json="$1"
    local issue_num="$2"
    local total_issues="$3"
    
    local linter=$(echo "$issue_json" | jq -r '.FromLinter')
    local text=$(echo "$issue_json" | jq -r '.Text')
    local filename=$(echo "$issue_json" | jq -r '.Pos.Filename')
    local line=$(echo "$issue_json" | jq -r '.Pos.Line')
    local column=$(echo "$issue_json" | jq -r '.Pos.Column')
    local severity=$(echo "$issue_json" | jq -r '.Severity')
    
    log_info "Issue [$issue_num/$total_issues]"
    echo "  Linter:   $linter" >&2
    echo "  File:     $filename:$line:$column" >&2
    echo "  Severity: $severity" >&2
    echo "  Problem:  $text" >&2
    echo "" >&2
    
    local sanitized_problem=$(sanitize_problem "$text")
    local problem_context="In file $filename at line $line, column $column: $sanitized_problem (detected by $linter)"
    
    local PROMPT="fix this problem: $filename:$line:$column $text"
    log_info "Calling clds..."
    log_debug "Prompt: $PROMPT"
    
    clds --print --verbose --output-format=stream-json <(<"$PROMPT")
}

# Main script
main() {
    check_dependencies
    
    local lint_output=$(get_lint_output)
    local filtered_output=$(filter_issues "$lint_output")
    local total_issues=$(echo "$filtered_output" | jq '.Issues | length')
    
    if [ "$total_issues" -eq 0 ]; then
        log_success "No issues found matching your criteria! Code is clean."
        exit 0
    fi
    
    log_info "Found $total_issues issues to process"

    local counter=1
    while IFS= read -r issue; do
        process_issue "$issue" "$counter" "$total_issues"
        counter=$((counter + 1))
        
        if [ "$counter" -le "$total_issues" ] && [ "$DRY_RUN" = false ]; then
            sleep "$DELAY_BETWEEN_CALLS"
        fi
    done < <(echo "$filtered_output" | jq -c '.Issues[]')
    
    log_success "Finished processing all issues"
}

# Parse arguments and run
parse_args "$@"
main
