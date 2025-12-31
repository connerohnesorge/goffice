# E2E Visual Testing Framework

End-to-end visual testing framework for comparing goffice PPTX generation against Microsoft's official Open-XML-SDK for .NET.

## Overview

This framework enables automated visual regression testing by:
1. Generating identical presentations using both Go (goffice) and C# (Open-XML-SDK)
2. Rendering presentations to high-resolution PNG images
3. Comparing images pixel-by-pixel with perceptual metrics
4. Generating comprehensive HTML reports with side-by-side comparisons

## Quick Start

### Prerequisites

- Go 1.25+ (provided by Nix)
- .NET SDK 9.0+ (provided by Nix)
- LibreOffice 7.6+ headless (provided by Nix)
- ImageMagick 7.x (provided by Nix)

### Running Tests

```bash
# Enter Nix development shell (includes all dependencies)
nix develop

# Run all E2E tests
go test ./... -v

# Run specific test category
go test ./... -run TestCharts -v

# Generate HTML report
./scripts/run_tests.sh
```

### Viewing Reports

After running tests, open the generated HTML report:

```bash
open reports/index.html  # macOS
xdg-open reports/index.html  # Linux
```

## Architecture

```
Test Case → Go Generator → PPTX → PNG ─┐
                                        ├─→ Visual Diff → HTML Report
Test Case → C# Generator → PPTX → PNG ─┘
```

### Components

- **framework/**: Core framework code (test definitions, rendering, comparison)
- **generators/go/**: Go generator using goffice
- **generators/csharp/**: C# generator using Open-XML-SDK
- **testcases/**: Test case definitions (charts, shapes, text)
- **output/**: Generated PPTX, PDF, PNG files (gitignored)
- **reports/**: HTML test reports (gitignored)

## Creating Test Cases

Test cases are defined in Go using declarative structs:

```go
package charts

import "github.com/unidoc/goffice/tests/e2e/framework"

func init() {
    framework.RegisterTestCase(framework.TestCase{
        ID:          "chart_bar_basic",
        Name:        "Basic Bar Chart",
        Description: "Simple clustered bar chart with 3 series",
        Category:    framework.CategoryChart,
        Tags:        []string{"chart", "bar", "basic"},
        Spec: framework.TestSpec{
            SlideCount: 1,
            SlideSize:  framework.SlideSize16x9,
            Slides: []framework.SlideSpec{
                {
                    Index:  0,
                    Layout: "Blank",
                    Elements: []framework.ElementSpec{
                        {
                            Type:     framework.ElementTypeChart,
                            Position: framework.Position{X: 914400, Y: 914400},
                            Size:     framework.Size{Width: 7315200, Height: 3657600},
                            Chart: &framework.ChartSpec{
                                Type:  framework.ChartTypeBarClustered,
                                Title: "Sales by Region",
                                Data: framework.ChartData{
                                    Categories: []string{"Q1", "Q2", "Q3", "Q4"},
                                    Series: []framework.SeriesData{
                                        {
                                            Name:   "North",
                                            Values: []float64{120, 150, 180, 200},
                                            Color:  "4472C4",
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
    })
}
```

## Configuration

Test configuration is set per-test case or globally:

```go
Config: framework.TestConfig{
    DiffThreshold:      0.01,  // 1% tolerance
    IgnoreAntialiasing: true,
    Renderer:           framework.RendererLibreOffice,
    DPI:                300,
    Timeout:            60 * time.Second,
    RetryCount:         2,
}
```

## Visual Comparison Algorithms

The framework supports multiple comparison algorithms:

- **Pixel-Perfect**: Exact pixel-by-pixel RGB comparison
- **SSIM**: Structural Similarity Index (perceptual quality metric)
- **MSE/PSNR**: Mean Squared Error and Peak Signal-to-Noise Ratio

Default: SSIM with 1% tolerance (99% similarity required)

## CI Integration

The framework integrates with GitHub Actions for automated testing on PRs:

```yaml
# .github/workflows/e2e-visual-tests.yml
- name: Run E2E Visual Tests
  run: |
    nix develop --command go test ./tests/e2e/... -v

- name: Upload Test Reports
  uses: actions/upload-artifact@v3
  with:
    name: e2e-reports
    path: tests/e2e/reports/
```

## Directory Reference

```
tests/e2e/
├── framework/          # Core framework
│   ├── testcase.go    # Test case types
│   ├── generator.go   # Generator interface
│   ├── renderer.go    # PPTX → PNG rendering
│   ├── differ.go      # Image comparison
│   ├── reporter.go    # HTML reports
│   └── config.go      # Configuration
│
├── generators/        # Platform generators
│   ├── go/           # Go (goffice)
│   └── csharp/       # C# (Open-XML-SDK)
│
├── testcases/        # Test definitions
│   ├── charts/       # Chart tests
│   ├── shapes/       # Shape tests
│   ├── text/         # Text tests
│   └── integration/  # Multi-feature tests
│
├── fixtures/         # Test data
│   ├── data/        # JSON data files
│   ├── images/      # Test images
│   └── fonts/       # Embedded fonts
│
├── golden/          # Reference images (optional)
├── output/          # Generated artifacts (gitignored)
├── reports/         # HTML reports (gitignored)
└── scripts/         # Helper scripts
```

## Troubleshooting

### Tests Failing with Rendering Errors

Ensure LibreOffice is installed and headless mode is working:

```bash
libreoffice --headless --version
```

### Visual Diff False Positives

Adjust threshold in test configuration:

```go
Config: framework.TestConfig{
    DiffThreshold:      0.05,  // Increase tolerance to 5%
    IgnoreAntialiasing: true,   // Reduce anti-aliasing sensitivity
}
```

### C# Generator Not Found

Ensure .NET SDK is installed:

```bash
dotnet --version  # Should show 9.0+
```

## Performance

- **Full test suite**: ~5 minutes (with caching)
- **Single test**: ~10-15 seconds
- **Rendering overhead**: ~2-3 seconds per slide

## Contributing

When adding new test cases:

1. Create test definition in `testcases/<category>/`
2. Register test case with `framework.RegisterTestCase()`
3. Run locally to verify both generators produce identical output
4. Ensure diff percentage is < configured threshold
5. Add to PR with updated test case count

## License

Same as goffice main license.
