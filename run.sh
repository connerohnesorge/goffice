#!/bin/bash

# run.sh - Apply spectr changes to all proposals
# Runs /spectr:apply on each change folder 3 times serially

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Logging functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $(date '+%Y-%m-%d %H:%M:%S') $*"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $(date '+%Y-%m-%d %H:%M:%S') $*"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $(date '+%Y-%m-%d %H:%M:%S') $*"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $(date '+%Y-%m-%d %H:%M:%S') $*"
}

log_header() {
    echo ""
    echo -e "${CYAN}============================================================${NC}"
    echo -e "${CYAN}$*${NC}"
    echo -e "${CYAN}============================================================${NC}"
}

# Configuration
CHANGES_DIR="spectr/changes"
ITERATIONS=3
LOG_DIR="logs/spectr-apply"

# Create log directory
mkdir -p "$LOG_DIR"

# Get all change directories
CHANGE_DIRS=$(find "$CHANGES_DIR" -mindepth 1 -maxdepth 1 -type d | sort)

if [ -z "$CHANGE_DIRS" ]; then
    log_error "No change directories found in $CHANGES_DIR"
    exit 1
fi

# Count total changes
TOTAL_CHANGES=$(echo "$CHANGE_DIRS" | wc -l | tr -d ' ')
CURRENT_CHANGE=0
FAILED_CHANGES=()
SUCCESSFUL_CHANGES=()

log_header "Starting spectr:apply for $TOTAL_CHANGES changes ($ITERATIONS iterations each)"
echo ""

for change_path in $CHANGE_DIRS; do
    CURRENT_CHANGE=$((CURRENT_CHANGE + 1))
    change_id=$(basename "$change_path")

    log_header "[$CURRENT_CHANGE/$TOTAL_CHANGES] Processing: $change_id"

    # Log file for this change
    change_log="$LOG_DIR/${change_id}.log"
    > "$change_log"  # Clear/create log file

    change_failed=false

    for iteration in $(seq 1 $ITERATIONS); do
        log_info "Iteration $iteration/$ITERATIONS for $change_id"
        echo "--- Iteration $iteration ---" >> "$change_log"

        # Run the command and capture output
        set +e  # Don't exit on error
        # output=$(gemini --yolo -o text "/spectr:apply $change_id" 2>&1)
        output=$(copilot --allow-all-tools -p "/spectr:apply $change_id" 2>&1)
        exit_code=$?
        set -e

        # Log output
        echo "$output" >> "$change_log"
        echo "" >> "$change_log"

        # Show live output (truncated if too long)
        if [ -n "$output" ]; then
            # Show first 20 lines of output for live feedback
            echo "$output" | head -20
            line_count=$(echo "$output" | wc -l)
            if [ "$line_count" -gt 20 ]; then
                echo -e "${YELLOW}  ... ($((line_count - 20)) more lines, see $change_log)${NC}"
            fi
        fi

        if [ $exit_code -eq 0 ]; then
            log_success "Iteration $iteration completed successfully"
        else
            log_error "Iteration $iteration failed with exit code $exit_code"
            change_failed=true
        fi

        echo ""
    done

    if [ "$change_failed" = true ]; then
        log_error "Change $change_id had failures (see $change_log)"
        FAILED_CHANGES+=("$change_id")
    else
        log_success "Change $change_id completed all iterations successfully"
        SUCCESSFUL_CHANGES+=("$change_id")
    fi
done

# Summary
log_header "Summary"
echo ""
log_info "Total changes processed: $TOTAL_CHANGES"
log_success "Successful: ${#SUCCESSFUL_CHANGES[@]}"
log_error "Failed: ${#FAILED_CHANGES[@]}"

if [ ${#FAILED_CHANGES[@]} -gt 0 ]; then
    echo ""
    log_warn "Failed changes:"
    for failed in "${FAILED_CHANGES[@]}"; do
        echo "  - $failed"
    done
fi

echo ""
log_info "Logs saved to: $LOG_DIR/"
echo ""

# Exit with error if any changes failed
if [ ${#FAILED_CHANGES[@]} -gt 0 ]; then
    exit 1
fi

exit 0
