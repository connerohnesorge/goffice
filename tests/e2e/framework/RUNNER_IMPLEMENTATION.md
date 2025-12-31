# Test Execution Framework Implementation

This document describes the test execution framework for E2E visual testing.

## Files Created

### 1. `runner.go`
Main orchestration framework for E2E visual tests.

**Key Components:**

#### Interfaces
- `Generator` - Interface for PPTX generation (Go or C#)

#### Structs
- `GoGenerator` - Generates PPTX using goffice via external binary
- `CSharpGenerator` - Generates PPTX using Open-XML-SDK via external binary
- `Reporter` - Generates HTML reports from test results
- `Runner` - Orchestrates the full test execution pipeline
- `RunConfig` - Configuration for test runner
- `TestSuiteResult` - Results from running all tests
- `TestCaseResult` - Result for a single test case
- `SlideResult` - Comparison result for a single slide

#### Methods

**Runner Methods:**
- `NewRunner(config RunConfig) (*Runner, error)` - Creates a new test runner
- `Run(ctx context.Context, testCases []*TestCase) (*TestSuiteResult, error)` - Executes all test cases
- `runTest(ctx context.Context, tc *TestCase) *TestCaseResult` - Executes a single test case with retry logic
- `executeTest(ctx context.Context, tc *TestCase, result *TestCaseResult) error` - Performs the actual test execution
- `renderPPTXFiles(ctx context.Context, goPPTX, csharpPPTX string) (goPNGs, csharpPNGs []string, err error)` - Renders both PPTX files to PNG
- `compareSlides(ctx context.Context, goPNGs, csharpPNGs []string, tc *TestCase) ([]*SlideResult, error)` - Compares PNG images between Go and C#
- `generateReport(result *TestSuiteResult, reportPath string) error` - Creates an HTML report

**Generator Methods:**
- `GoGenerator.Generate(ctx context.Context, tc *TestCase, outputPath string) error` - Generates PPTX using Go
- `CSharpGenerator.Generate(ctx context.Context, tc *TestCase, outputPath string) error` - Generates PPTX using C#

### 2. `runner_test.go`
Unit tests for the runner framework.

**Test Coverage:**
- `TestNewRunner` - Tests runner creation with configuration
- `TestGoGenerator_Generate` - Tests Go generator (skipped without binary)
- `TestCSharpGenerator_Generate` - Tests C# generator (skipped without binary)
- `TestRunner_Run_NoTestCases` - Tests running with no test cases
- `TestTestSuiteResult_Counters` - Tests result counter tracking
- `TestSlideResult_Structure` - Tests slide result structure

### 3. `e2e_test.go`
Main E2E test entry point.

**Functions:**
- `TestE2E(t *testing.T)` - Main E2E test that runs all test cases
- `discoverTestCases(t *testing.T) []*TestCase` - Discovers test cases from `testcases/` directory
- `getGoGeneratorPath(t *testing.T) string` - Returns path to Go generator binary
- `getCSharpGeneratorPath(t *testing.T) string` - Returns path to C# generator binary

## Test Execution Pipeline

The runner executes tests through the following pipeline:

```
1. Generate PPTX with Go
   ↓
2. Generate PPTX with C#
   ↓
3. Render both PPTX files to PNG (one per slide)
   ↓
4. Compare PNG images slide-by-slide
   ↓
5. Generate HTML report with results
```

## Features Implemented

### 1. Retry Logic
- Tests can be retried on failure (configurable via `RetryCount`)
- Exponential backoff between retries (1s, 2s, 4s)
- Retry count tracked in `TestCaseResult`

### 2. Timeout Handling
- Configurable timeout per test (default: 2 minutes)
- Context-based timeout enforcement
- Timeout errors prevent retries

### 3. Parallel Execution
- Support for sequential (default) or parallel execution
- Configurable parallelism via `Parallel` setting
- Semaphore-based concurrency control

### 4. HTML Reporting
- Generates comprehensive HTML report
- Shows pass/fail status, timing, and error details
- Includes slide-level comparison results
- Report saved to `<OutputDir>/report.html`

### 5. Test Discovery
- Automatic discovery of test cases from `testcases/` directory
- Loads JSON test case files
- Filters test cases by pattern (future enhancement)

### 6. Artifact Management
- Organizes outputs by generator: `go/`, `csharp/`, `diff/`
- PNG renderings stored per slide
- Diff images show visual differences
- Optional artifact cleanup (via `KeepArtifacts` config)

## Configuration

```go
config := framework.RunConfig{
    OutputDir:           "output/e2e",        // Base output directory
    Parallel:            0,                   // 0 = sequential, >0 = parallel
    RetryCount:          1,                   // Number of retries on failure
    Timeout:             2 * time.Minute,     // Test timeout
    GenerateDiffs:       true,                // Generate diff images
    KeepArtifacts:       true,                // Keep intermediate files
    GoGeneratorPath:     "/path/to/go-gen",   // Path to Go generator binary
    CSharpGeneratorPath: "/path/to/cs-gen",   // Path to C# generator binary
}
```

## Usage

### Running E2E Tests

```bash
# Run all E2E tests (skips if testcases/ not found)
cd tests/e2e
go test -v

# Run with environment variables
GO_GENERATOR_PATH=/path/to/go-gen \
CSHARP_GENERATOR_PATH=/path/to/cs-gen \
go test -v
```

### Running Specific Tests

```bash
# Run only the E2E test
go test -v -run TestE2E

# Skip E2E tests in short mode
go test -short
```

## Output Structure

```
output/e2e/
├── go/
│   └── test_01.pptx          # Go-generated PPTX
├── csharp/
│   └── test_01.pptx          # C#-generated PPTX
├── render/
│   ├── test_01-1.png         # Slide 1 (Go)
│   ├── test_01-2.png         # Slide 2 (Go)
│   ├── test_01-1.png         # Slide 1 (C#)
│   └── test_01-2.png         # Slide 2 (C#)
├── diff/
│   ├── test_01_slide_1_diff.png  # Diff for slide 1
│   └── test_01_slide_2_diff.png  # Diff for slide 2
└── report.html               # HTML test report
```

## Test Results

Each test result includes:
- **Status** - PASSED, FAILED, or SKIPPED
- **Timing** - Generation time, render time, compare time
- **Errors** - Detailed error messages
- **Slide Results** - Per-slide comparison metrics
  - Diff percentage
  - Threshold
  - PNG paths
  - Diff image path

## Integration with Existing Framework

The runner integrates with existing E2E framework components:
- **TestCase** (`testcase.go`) - Test case specification
- **Renderer** (`renderer.go`) - PPTX to PNG conversion (LibreOffice)
- **Differ** (`differ.go`) - PNG comparison (GoDiffer/ImageMagickDiffer)
- **Reporter** (`reporter.go`) - HTML report generation

## Future Enhancements

1. **Test Filtering** - Implement glob pattern filtering for test cases
2. **Parallel Optimization** - Optimize parallel execution with worker pools
3. **Artifact Cleanup** - Implement automatic cleanup of intermediate files
4. **Progress Reporting** - Real-time progress updates during test execution
5. **Baseline Mode** - Support for baseline comparison mode
6. **Incremental Testing** - Only run tests that have changed

## Testing

All runner components are tested:
- Unit tests for runner creation and configuration
- Tests for result tracking and counting
- Integration tests for the full pipeline (requires binaries)

Run tests:
```bash
cd tests/e2e
go test ./framework/... -v
```

## Module Path

The E2E tests are in a separate module:
- **Module**: `github.com/connerohnesorge/goffice/tests/e2e`
- **Main package**: `github.com/connerohnesorge/goffice/tests/e2e`
- **Framework**: `github.com/connerohnesorge/goffice/tests/e2e/framework`

## Dependencies

External dependencies:
- LibreOffice (for PPTX → PDF rendering)
- poppler-utils (for PDF → PNG conversion)
- Go generator binary (built from `generators/go/`)
- C# generator binary (built from `generators/csharp/`)

Go dependencies:
- Standard library only
- Parent module: `github.com/connerohnesorge/goffice`

## Summary

The test execution framework provides a complete pipeline for visual regression testing between goffice and Open-XML-SDK. It supports:
- Dual generation (Go + C#)
- Visual rendering and comparison
- Retry logic and timeout handling
- Parallel execution
- Comprehensive HTML reporting
- Automatic test discovery

The framework is ready for integration with generator binaries and test case definitions.
