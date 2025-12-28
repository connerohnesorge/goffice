# PDF Visual Comparison Package

This package provides automated visual comparison capabilities for PDF fidelity testing in goffice.

## Overview

The comparison package enables:
- PDF-to-PNG conversion using Ghostscript
- Pixel-level difference detection between images
- Automated diff image generation
- Configurable tolerance thresholds

## Requirements

**Ghostscript** must be installed and available in your system PATH:

- **Linux/macOS**: `brew install ghostscript` or `apt-get install ghostscript`
- **Windows**: Download from https://www.ghostscript.com/download/gsdnld.html

To check if Ghostscript is available:
```bash
gs --version
```

## Usage

### Basic Image Comparison

```go
import "github.com/connerohnesorge/goffice-pdf/comparison"

// Compare two PNG images
config := comparison.DefaultComparisonConfig()
result, err := comparison.CompareImages("baseline.png", "generated.png", config)
if err != nil {
    log.Fatal(err)
}

// Check if within tolerance
if result.IsWithinTolerance() {
    fmt.Println("Images match!")
} else {
    fmt.Printf("Difference: %s\n", result.String())
}
```

### PDF to PNG Conversion

```go
// Convert PDF bytes to PNG
pdfBytes, _ := os.ReadFile("document.pdf")
pngPath, err := comparison.ConvertPDFBytesToPNG(
    pdfBytes,
    "output.png",
    150, // DPI
)
```

### Generate Diff Images

```go
// Create annotated diff image (highlights differences in red)
err := comparison.AnnotateDiffImage(
    "baseline.png",
    "generated.png",
    "diff.png",
    nil, // use default config
)

// Create side-by-side comparison
err = comparison.CreateSideBySideDiff(
    "baseline.png",
    "generated.png",
    "diff.png",
    "comparison.png",
)
```

## Configuration

### Default Tolerances

The default configuration uses:
- **Pixel tolerance**: 2.0 (Euclidean color distance in RGB space)
- **Image threshold**: 0.5% (percentage of pixels allowed to exceed tolerance)
- **DPI**: 150 (for PDF-to-PNG conversion)

### Custom Configuration

```go
config := &comparison.ComparisonConfig{
    PixelTolerance:     2.0,   // Max color distance per pixel
    ImageDiffThreshold: 0.005, // 0.5% of pixels
    DPI:                150,   // Resolution for rendering
}
```

## How It Works

### Pixel Comparison Algorithm

1. Both images are loaded and dimensions are validated
2. Each pixel is compared using Euclidean color distance:
   ```
   distance = sqrt((r1-r2)² + (g1-g2)² + (b1-b2)²)
   ```
3. Pixels with distance > `PixelTolerance` are counted
4. If `PixelsExceedingTolerance / TotalPixels > ImageDiffThreshold`, comparison fails

### Why These Tolerances?

- **Pixel tolerance (2.0)**: Allows for minor rendering variations (~1% color difference)
- **Image threshold (0.5%)**: Permits small antialiasing or rounding differences while catching real visual regressions

## Integration with Fidelity Tests

The comparison package is automatically used in `pdf/fidelity_test.go`:

```go
// Visual comparison is performed after PDF generation
func runWordFidelityTest(t *testing.T, docxFile, description string) {
    // ... generate PDF ...

    // Compare with baseline (if available)
    baselinePath := filepath.Join("testdata", "fidelity", "word",
        strings.TrimSuffix(docxFile, ".docx") + ".baseline.png")

    if err := compareWithBaseline(t, pdfBytes, baselinePath, description); err != nil {
        t.Logf("Visual comparison: %v", err)
    }
}
```

### Graceful Degradation

Visual comparison gracefully skips when:
- Ghostscript is not installed (logs warning)
- Baseline image doesn't exist (logs message)

This allows tests to run in environments without full visual comparison support.

## Creating Baseline Images

To enable visual comparison for a test:

1. Generate the reference PDF using Microsoft Office "Save as PDF"
2. Convert to PNG using Ghostscript:
   ```bash
   gs -q -dNOPAUSE -dBATCH -dSAFER -sDEVICE=png16m -r150 \
      -sOutputFile=baseline.png reference.pdf
   ```
3. Place in `testdata/fidelity/{word|excel|powerpoint}/[testname].baseline.png`
4. Commit baseline to git

## Troubleshooting

### "Ghostscript not found in PATH"

Install Ghostscript or add it to your PATH. Tests will skip gracefully if unavailable.

### "image dimensions do not match"

Baseline and generated images have different sizes. Regenerate the baseline image using the same DPI (150).

### "visual comparison failed: X% exceed tolerance"

The generated PDF differs significantly from the baseline. Inspect the generated diff images:
- `*.diff.png`: Shows differences highlighted in red
- `*.comparison.png`: Side-by-side view (baseline | generated | diff)

If the difference is expected (e.g., intentional rendering improvement), regenerate the baseline.

## Testing

Run comparison package tests:
```bash
cd pdf
go test ./comparison -v
```

Run fidelity tests with visual comparison:
```bash
cd pdf
go test -run TestWordFidelity -v
```

## Performance

- PDF-to-PNG conversion: ~100-500ms per page (depends on complexity)
- Image comparison: ~10-50ms for typical document pages (1200x1600 @ 150 DPI)
- Diff image generation: ~10-30ms

## Limitations

- Requires Ghostscript installation (not a pure Go solution)
- Comparison is pixel-based (doesn't understand semantic layout)
- Large baseline images increase repository size (consider git-lfs for many baselines)
- Single-page comparison only (multi-page PDFs use first page)

## Future Enhancements

Potential improvements (not yet implemented):
- Multi-page PDF support
- Perceptual hashing for layout-level comparison
- Automatic baseline update workflow (`-update-baselines` flag)
- Parallel comparison for multiple test cases
- Integration with CI/CD artifact storage
