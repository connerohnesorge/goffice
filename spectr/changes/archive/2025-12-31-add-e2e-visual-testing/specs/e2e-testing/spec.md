# Capability: E2E Testing

End-to-end visual testing framework for comparing goffice PPTX generation against Microsoft's official Open-XML-SDK for .NET.

---

## ADDED Requirements

### Requirement: Test Case Definition

The system SHALL provide a declarative, platform-agnostic JSON-based format for defining visual comparison test cases.

#### Scenario: Define chart test case
- GIVEN a developer wants to test bar chart generation
- WHEN they create a test case JSON with chart spec (type, data, styling)
- THEN both Go and C# generators can read and execute the same test case
- AND the test case includes position, size, colors, data series, and axes configuration

#### Scenario: Define shape test case
- GIVEN a developer wants to test shape rendering
- WHEN they create a test case with shape specs (type, fill, stroke, effects)
- THEN both generators create identical shapes with same visual properties
- AND effects (shadows, glows, reflections) are rendered consistently

#### Scenario: Define text formatting test case
- GIVEN a developer wants to test text rendering
- WHEN they create a test case with text specs (font, size, color, bold, italic, alignment)
- THEN both generators produce visually identical text output
- AND paragraph spacing, indentation, and line spacing match exactly

###Requirement: Dual Generator Infrastructure

The system SHALL provide both Go and C# generator harnesses that execute identical test cases.

#### Scenario: Execute test in Go generator
- GIVEN a test case JSON file
- WHEN the Go generator is invoked with `-test <file> -output <path>`
- THEN it creates a .pptx file using goffice APIs
- AND the PPTX contains all elements specified in the test case
- AND element positions match EMU coordinates exactly

#### Scenario: Execute test in C# generator
- GIVEN the same test case JSON file
- WHEN the C# generator is invoked with `--test <file> --output <path>`
- THEN it creates a .pptx file using Open-XML-SDK APIs
- AND the PPTX matches the Go output structurally
- AND all styling, colors, and formatting are applied identically

#### Scenario: Handle generator timeout
- GIVEN a generator is running
- WHEN execution exceeds configured timeout (default: 60s)
- THEN the process is terminated gracefully
- AND the test result is marked as "error" with timeout message

### Requirement: PPTX to Image Rendering

The system SHALL convert generated PPTX files to high-resolution PNG images for pixel-level comparison.

#### Scenario: Render using LibreOffice headless
- GIVEN a PPTX file from either generator
- WHEN rendering is invoked with LibreOffice backend
- THEN each slide is converted to PDF first
- AND PDF pages are extracted as 300 DPI PNG images
- AND PNG files are saved with sequential naming (slide_1.png, slide_2.png, ...)

#### Scenario: Render using PowerPoint automation (Windows)
- GIVEN a PPTX file and Windows environment
- WHEN rendering is invoked with PowerPoint backend
- THEN PowerPoint COM automation exports each slide as PNG
- AND images are rendered at specified DPI (default: 300)
- AND PowerPoint instance is properly closed after rendering

#### Scenario: Handle rendering failure
- GIVEN rendering is attempted
- WHEN LibreOffice/PowerPoint fails or is unavailable
- THEN the system logs detailed error message
- AND the test result is marked as "error"
- AND intermediate files (PDF) are preserved for debugging if configured

### Requirement: Visual Difference Detection

The system SHALL compare PNG images pixel-by-pixel and calculate perceptual similarity metrics.

#### Scenario: Pixel-perfect comparison
- GIVEN two PNG images from Go and C# generators
- WHEN comparison is performed with PixelPerfect algorithm
- THEN exact pixel-by-pixel RGB comparison is performed
- AND difference count and percentage are calculated
- AND a diff image is generated highlighting changed pixels in red

#### Scenario: Perceptual similarity comparison (SSIM)
- GIVEN two PNG images
- WHEN comparison is performed with SSIM algorithm
- THEN Structural Similarity Index is calculated (range: -1 to 1)
- AND SSIM >= 0.99 indicates <1% visual difference
- AND test passes if SSIM >= (1 - threshold)

#### Scenario: MSE/PSNR comparison
- GIVEN two PNG images
- WHEN comparison is performed with MSE or PSNR algorithm
- THEN Mean Squared Error is calculated
- AND Peak Signal-to-Noise Ratio is derived (higher = more similar)
- AND PSNR > 30 dB is considered high quality match

#### Scenario: Handle dimension mismatch
- GIVEN two images with different dimensions
- WHEN comparison is attempted
- THEN the system rejects the comparison immediately
- AND the test result is marked as "error"
- AND error message indicates dimension mismatch (e.g., "1920x1080 vs 1920x1200")

### Requirement: Diff Threshold Configuration

The system SHALL allow configurable tolerance for visual differences to reduce false positives.

#### Scenario: Set global threshold
- GIVEN a test configuration
- WHEN diff_threshold is set to 0.01 (1%)
- THEN tests pass if visual difference is ≤ 1%
- AND tests fail if visual difference is > 1%

#### Scenario: Set per-test threshold
- GIVEN a specific test case with known anti-aliasing differences
- WHEN the test case overrides threshold to 0.05 (5%)
- THEN only that test uses 5% tolerance
- AND other tests use global default

#### Scenario: Ignore anti-aliasing differences
- GIVEN a test configuration with ignore_antialiasing=true
- WHEN comparison detects edge pixel differences < 10% luminance change
- THEN those pixels are excluded from diff count
- AND tests are less sensitive to font rendering variations

### Requirement: HTML Report Generation

The system SHALL generate comprehensive HTML reports with side-by-side visual comparisons.

#### Scenario: Generate summary report
- GIVEN all tests have completed
- WHEN report generation is invoked
- THEN an index.html file is created with:
  - Total, passed, failed test counts
  - Average, min, max diff percentages
  - List of all test results with status badges
  - Execution time for each test

#### Scenario: Display side-by-side comparisons
- GIVEN a test result with diff data
- WHEN viewing the HTML report
- THEN each slide shows three images side-by-side:
  - Go generator output
  - C# generator output
  - Diff overlay (red highlights)
- AND clicking an image opens a full-size zoom view

#### Scenario: Export JSON summary
- GIVEN a generated report
- WHEN results.json is accessed
- THEN it contains machine-readable test results
- AND it includes all metrics (SSIM, MSE, PSNR, diff %)
- AND it can be consumed by CI systems or dashboards

### Requirement: Nix Development Environment

The system SHALL provide a reproducible, isolated development environment via Nix flake.

#### Scenario: Enter development shell
- GIVEN the project flake.nix
- WHEN developer runs `nix develop`
- THEN shell provides:
  - Go 1.25+ toolchain
  - .NET SDK 9.0+
  - LibreOffice 7.6+ headless
  - ImageMagick 7.x
  - Poppler utils (pdftoppm)
  - All test utilities (gotestsum)

#### Scenario: Run tests with single command
- GIVEN Nix dev shell is active
- WHEN developer runs `nix run .#e2e-tests`
- THEN all E2E tests execute automatically
- AND report is generated in tests/e2e/reports/
- AND exit code reflects pass/fail status

#### Scenario: Hermetic build
- GIVEN two developers on different platforms (Linux, macOS)
- WHEN both enter Nix dev shell
- THEN they get identical tool versions
- AND test results are reproducible across machines

### Requirement: CI Integration

The system SHALL integrate with GitHub Actions for automated visual regression testing on PRs.

#### Scenario: Run E2E tests on PR
- GIVEN a pull request is opened
- WHEN CI workflow triggers
- THEN E2E visual tests execute in GitHub Actions runner
- AND test results are uploaded as workflow artifacts
- AND PR comment is posted with summary (X passed, Y failed)

#### Scenario: Cache dependencies for speed
- GIVEN CI has run previously
- WHEN tests run again on same branch
- THEN Nix store is cached across runs
- AND .NET packages are cached
- AND test execution completes in <5 minutes

#### Scenario: Store baseline images
- GIVEN a new test is added
- WHEN it runs for the first time and passes
- THEN C# output PNG is stored as baseline (optional)
- AND subsequent runs compare against stored baseline
- AND baseline update command regenerates from latest C# output

### Requirement: Test Case Library

The system SHALL provide a comprehensive library of pre-defined test cases covering all PPTX features.

#### Scenario: Chart test coverage
- GIVEN the test case library
- WHEN developer lists chart tests
- THEN test cases exist for all 11 chart types:
  - Bar (clustered, stacked, percent-stacked)
  - Line (standard, smooth, with markers)
  - Pie, Doughnut
  - Area (standard, stacked)
  - Scatter (marker, line, smooth)
  - Bubble, Radar, Stock, Surface
- AND each chart type has variants testing colors, legends, axes, data labels

#### Scenario: Shape test coverage
- GIVEN the test case library
- WHEN developer lists shape tests
- THEN test cases exist for:
  - Basic shapes (rectangle, ellipse, triangle, etc.)
  - Fill types (solid, gradient, pattern)
  - Stroke styles (solid, dashed, dotted)
  - Effects (shadows, glows, reflections)

#### Scenario: Add custom test case
- GIVEN a developer wants to test a specific feature combination
- WHEN they create a new test case Go file in testcases/
- THEN they define TestCase struct with spec
- AND running test framework auto-discovers and executes it
- AND results appear in report alongside built-in tests

### Requirement: Performance Benchmarking

The system SHALL measure and report generation and rendering performance.

#### Scenario: Measure generation time
- GIVEN a test case is executed
- WHEN Go generator runs
- THEN execution time is measured and recorded
- WHEN C# generator runs
- THEN execution time is measured and recorded
- AND report shows performance comparison (Go: 150ms, C#: 200ms)

#### Scenario: Identify slow tests
- GIVEN test suite has completed
- WHEN viewing performance dashboard
- THEN tests are sorted by execution time
- AND tests exceeding 10 seconds are flagged
- AND recommendations are provided for optimization

### Requirement: Error Handling and Debugging

The system SHALL provide detailed error messages and preserve artifacts for debugging failures.

#### Scenario: Preserve intermediate files on failure
- GIVEN a test fails
- WHEN save_intermediates config is true
- THEN PPTX, PDF, and PNG files are preserved
- AND file paths are included in error message
- AND developer can manually inspect files

#### Scenario: Retry flaky tests
- GIVEN a test fails
- WHEN retry_count is configured to 2
- THEN test is re-run up to 2 more times
- AND test passes if any retry succeeds
- AND report indicates retry count (e.g., "Passed on retry 2")

#### Scenario: Detailed error logging
- GIVEN any operation fails (generation, rendering, comparison)
- WHEN error occurs
- THEN detailed error message is logged including:
  - Component that failed (generator, renderer, differ)
  - Exact error from subprocess (LibreOffice, ImageMagick)
  - Stack trace if applicable
  - File paths for debugging

---

## Validation Criteria

1. **Test Execution**: All test cases can execute in both Go and C# generators without errors
2. **Reproducibility**: Running the same test twice produces identical results (bit-identical PNGs)
3. **Performance**: Full test suite completes in <5 minutes with caching enabled
4. **Accuracy**: False positive rate <5% with default threshold (0.01)
5. **Coverage**: Test library includes all 11 chart types, 10+ shape types, text formatting variants
6. **CI Integration**: GitHub Actions workflow runs successfully on Linux and macOS runners
7. **Reporting**: HTML report renders correctly in Chrome, Firefox, Safari
8. **Nix Environment**: `nix develop` provides all dependencies on Linux (x86_64, aarch64) and macOS (x86_64, aarch64)

---

## Dependencies

- Go 1.25+ (for Go generator and framework)
- .NET SDK 9.0+ (for C# generator)
- goffice library (Go OOXML implementation)
- Open-XML-SDK 3.x (C# OOXML implementation)
- LibreOffice 7.6+ or Microsoft PowerPoint (for rendering)
- ImageMagick 7.x (for visual comparison)
- Poppler utils (pdftoppm for PDF → PNG conversion)
- Nix package manager 2.18+ (for dev environment)

---

## Future Enhancements (Out of Scope)

- **Animated slides**: Testing PowerPoint transitions/animations (static rendering only in V1)
- **Excel/Word E2E**: Extending framework to spreadsheet and wordprocessing (presentation-only in V1)
- **Cloud storage**: Storing test artifacts in S3/GCS (local filesystem only in V1)
- **Parallel execution**: Running tests concurrently (sequential execution in V1)
- **Web UI**: Interactive dashboard for browsing test results (HTML reports only in V1)
- **Baseline management**: Git LFS integration for storing reference images (optional manual baseline storage in V1)

---

## Success Metrics

1. **Regression Detection**: Catch at least 95% of visual regressions before merge
2. **Developer Adoption**: ≥80% of PRs touching presentation code include E2E test updates
3. **Maintenance Burden**: <10% of E2E test failures are false positives requiring threshold tuning
4. **CI Reliability**: E2E test step passes ≥95% of the time (excluding legitimate failures)
5. **Coverage Growth**: Test library grows by ≥10 test cases per quarter
