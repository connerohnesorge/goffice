# E2E Test Reporter

The E2E Test Reporter generates comprehensive HTML and JSON reports from test execution results.

## Features

### HTML Report (`index.html`)
- **Interactive Dashboard**: Search and filter test results by name or status
- **Summary Statistics**: Total tests, passed/failed/skipped counts, pass rate, total duration
- **Detailed Results**: Individual test results with:
  - Test scenario name and status
  - Execution duration
  - Output file paths (Go and .NET)
  - Comparison results (XML, Binary, Visual match status)
  - Error messages for failures
- **Embedded CSS & JavaScript**: Self-contained HTML file with no external dependencies
- **Responsive Design**: Clean, professional layout that works on all screen sizes

### JSON Report (`results.json`)
- **Machine-Readable Format**: Structured JSON for CI/CD integration
- **Complete Test Data**: All test results with timestamps and comparison details
- **Schema Version**: Versioned format for compatibility tracking
- **Summary Metrics**: Aggregated statistics for quick analysis

## Usage

### Basic Usage

```go
import "github.com/connerohnesorge/goffice/tests/e2e/framework"

// Create reporter
reporter := framework.NewReporter("/path/to/output")

// Generate both HTML and JSON reports
err := reporter.Generate(testResults)
if err != nil {
    log.Fatalf("Failed to generate reports: %v", err)
}

// Reports created at:
// - /path/to/output/index.html
// - /path/to/output/results.json
```

### Backward-Compatible Function

```go
// Generate only HTML report (maintains backward compatibility)
err := framework.GenerateHTMLReport(testResults, "/path/to/report.html")
```

## Report Structure

### HTML Report Sections

1. **Header**: Report title and generation timestamp
2. **Controls**: Search box and status filter dropdown
3. **Summary Cards**:
   - Total Tests
   - Passed
   - Failed
   - Skipped
   - Pass Rate (%)
   - Total Duration
4. **Test Results**: Expandable list of all test executions
5. **Footer**: Framework branding

### JSON Report Schema

```json
{
  "version": "1.0",
  "generated_at": "2025-12-31T12:00:00Z",
  "summary": {
    "Total": 10,
    "Passed": 8,
    "Failed": 1,
    "Skipped": 1,
    "PassRate": 80.0,
    "TotalDuration": 45000000000
  },
  "results": [
    {
      "scenario": "test_name",
      "status": "PASSED",
      "start_time": "2025-12-31T12:00:00Z",
      "end_time": "2025-12-31T12:00:05Z",
      "duration_seconds": 5.0,
      "go_output_path": "/tmp/go/test.docx",
      "dotnet_output_path": "/tmp/dotnet/test.docx",
      "comparison": {
        "xml_match": true,
        "binary_match": false,
        "visual_match": true
      }
    }
  ]
}
```

## Interactive Features

### Search
- Type in the search box to filter tests by scenario name
- Search is case-insensitive and matches partial strings
- Updates results in real-time as you type

### Status Filter
- Dropdown to filter by test status:
  - All Tests
  - Passed
  - Failed
  - Skipped
- Combines with search for precise filtering

## Integration with Test Runner

The Reporter is automatically used by the test Runner:

```go
runner, err := framework.NewRunner(config)
if err != nil {
    log.Fatal(err)
}

// Runner automatically generates reports after execution
result, err := runner.Run(ctx, testCases)
if err != nil {
    log.Fatal(err)
}

// Reports available at: <outputDir>/report.html
```

## Customization

### Custom Output Directory

```go
reporter := framework.NewReporter("/custom/path")
```

### Custom Report Processing

```go
// Access summary statistics
summary := framework.CalculateSummary(results)
fmt.Printf("Pass rate: %.1f%%\n", summary.PassRate)

// Generate individual reports
reporter.generateHTMLReport(results, "custom.html")
reporter.generateJSONReport(results, "custom.json")
```

## Error Handling

The reporter includes comprehensive error handling:

```go
err := reporter.Generate(results)
if err != nil {
    // Errors include context about which report failed
    if strings.Contains(err.Error(), "HTML") {
        log.Printf("HTML report generation failed: %v", err)
    } else if strings.Contains(err.Error(), "JSON") {
        log.Printf("JSON report generation failed: %v", err)
    }
}
```

## CI/CD Integration

### Reading JSON Reports in CI

```bash
# Extract pass rate from JSON report
jq '.summary.PassRate' results.json

# Count failures
jq '.summary.Failed' results.json

# List failed tests
jq -r '.results[] | select(.status == "FAILED") | .scenario' results.json

# Check if all tests passed
if [ $(jq '.summary.Failed' results.json) -eq 0 ]; then
    echo "All tests passed!"
fi
```

### GitHub Actions Example

```yaml
- name: Run E2E Tests
  run: go test -v ./tests/e2e/...

- name: Upload HTML Report
  if: always()
  uses: actions/upload-artifact@v3
  with:
    name: e2e-report
    path: tests/e2e/output/index.html

- name: Check Test Results
  run: |
    FAILURES=$(jq '.summary.Failed' tests/e2e/output/results.json)
    if [ $FAILURES -gt 0 ]; then
      echo "::error::$FAILURES tests failed"
      exit 1
    fi
```

## Design Philosophy

- **Self-Contained**: No external dependencies (CSS/JS embedded)
- **Accessibility**: Clean, readable design with semantic HTML
- **Performance**: Lightweight JavaScript for filtering
- **Standards**: Valid HTML5 and JSON
- **Maintainability**: Template-based generation for easy updates

## Future Enhancements

Potential additions in future versions:
- Side-by-side image comparison viewer
- Diff visualization for XML/binary comparisons
- Trend analysis across multiple test runs
- Export to PDF
- Custom themes/branding
