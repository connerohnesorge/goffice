# Visual Comparison Implementation Summary

## Overview

This document summarizes the implementation of automated visual comparison for PDF fidelity testing in the goffice project.

## What Was Implemented

### Core Components

1. **pdf/comparison/types.go**
   - `DiffResult` struct with comprehensive statistics
   - `ComparisonConfig` for configurable tolerances
   - Helper methods (`IsWithinTolerance`, `DiffPercentage`, etc.)
   - Default configuration constants

2. **pdf/comparison/convert.go**
   - `ConvertPDFToPNG()` - Convert PDF files to PNG using Ghostscript
   - `ConvertPDFBytesToPNG()` - Convenience wrapper for in-memory PDFs
   - `IsGhostscriptAvailable()` - Check for Ghostscript availability
   - Graceful error handling when Ghostscript is unavailable

3. **pdf/comparison/diff.go**
   - `CompareImages()` - Pixel-level image comparison
   - Euclidean color distance calculation in RGB space
   - Per-pixel and per-image tolerance thresholds
   - Detailed statistics (total pixels, diff pixels, max/avg deltas)

4. **pdf/comparison/annotate.go**
   - `AnnotateDiffImage()` - Generate diff images with red highlights
   - `CreateSideBySideDiff()` - Side-by-side comparison view
   - PNG saving utilities

5. **pdf/comparison/comparison_test.go**
   - Unit tests for image comparison (identical, slight diff, large diff)
   - Dimension mismatch handling tests
   - Diff image generation tests
   - Synthetic test image utilities

### Integration

6. **pdf/fidelity_test.go** (Updated)
   - Added `compareWithBaseline()` helper function
   - Integrated visual comparison into `runWordFidelityTest()`
   - Graceful skip when baseline missing or Ghostscript unavailable
   - Automatic diff image generation on comparison failure
   - Side-by-side comparison images for manual review

### Documentation

7. **pdf/comparison/README.md**
   - Comprehensive usage guide
   - Ghostscript installation instructions
   - Configuration examples
   - Troubleshooting section
   - Performance characteristics
   - Future enhancement ideas

8. **pdf/comparison/IMPLEMENTATION.md** (This file)
   - Implementation summary
   - File listing
   - Design decisions
   - Testing results

## Design Decisions

### 1. Ghostscript for PDF Rendering
- **Why**: Industry-standard, widely available, consistent output
- **Alternative**: Native Go PDF libraries (requires CGO, against project philosophy)
- **Trade-off**: External dependency, but graceful degradation when unavailable

### 2. RGB Euclidean Distance
- **Why**: Simple, fast, sufficient for detecting visual differences
- **Alternative**: SSIM, Delta E (more perceptually accurate but complex)
- **Trade-off**: May be less accurate for subtle color differences, but good enough for regression testing

### 3. Tolerance Thresholds
- **Pixel tolerance**: 2.0 (allows ~1% color variation)
- **Image threshold**: 0.5% of pixels (allows minor antialiasing differences)
- **Why**: Balance between catching real regressions and allowing minor rendering variations
- **Configurable**: Can be adjusted via `ComparisonConfig`

### 4. Graceful Degradation
- **Why**: Tests should run in environments without Ghostscript
- **How**: Check for Ghostscript availability, skip comparison with warning
- **Benefit**: No CI/CD breakage, optional enhancement

## File Structure

```
pdf/comparison/
├── annotate.go           # Diff image generation
├── comparison_test.go    # Unit tests
├── convert.go            # PDF-to-PNG conversion
├── diff.go               # Pixel-level comparison
├── types.go              # Core types and constants
├── README.md             # User documentation
└── IMPLEMENTATION.md     # This file
```

## Testing Results

All tests pass successfully:

```
=== RUN   TestCompareImages_Identical
--- PASS: TestCompareImages_Identical (0.00s)
=== RUN   TestCompareImages_SlightDifference
--- PASS: TestCompareImages_SlightDifference (0.00s)
=== RUN   TestCompareImages_LargeDifference
--- PASS: TestCompareImages_LargeDifference (0.00s)
=== RUN   TestCompareImages_DimensionMismatch
--- PASS: TestCompareImages_DimensionMismatch (0.00s)
=== RUN   TestAnnotateDiffImage
--- PASS: TestAnnotateDiffImage (0.00s)
=== RUN   TestIsGhostscriptAvailable
--- PASS: TestIsGhostscriptAvailable (0.00s)
PASS
ok  	github.com/connerohnesorge/goffice-pdf/comparison
```

Fidelity tests compile and run successfully (skip because no test files exist yet):

```
=== RUN   TestWordFidelity
--- PASS: TestWordFidelity (0.00s)
    --- SKIP: TestWordFidelity/BasicText (0.00s)
    --- SKIP: TestWordFidelity/ComplexFormatting (0.00s)
    ...
```

## What's NOT Implemented (Future Work)

1. **Baseline Image Generation**: No actual baseline images created yet
   - Need test DOCX/XLSX/PPTX files
   - Need Microsoft Office "Save as PDF" reference outputs
   - Need to convert reference PDFs to baseline PNGs

2. **Multi-page PDF Support**: Currently only processes first page
   - Could extend to compare all pages
   - Would need page-by-page comparison logic

3. **Baseline Update Workflow**: No `-update-baselines` flag yet
   - Would regenerate all baselines automatically
   - Useful for intentional rendering changes

4. **Perceptual Hashing**: Layout-level comparison not implemented
   - Could detect structural changes (e.g., page breaks)
   - More robust than pixel-level for major refactors

5. **CI/CD Integration**: No automated reporting or artifact storage
   - Could upload diff images to CI artifacts
   - Could fail builds on visual regressions

## How to Use

### Basic Usage

```go
// In a test
baselinePath := "testdata/fidelity/word/basic_text.baseline.png"
pdfBytes := generatePDF()

if err := compareWithBaseline(t, pdfBytes, baselinePath, "Basic text test"); err != nil {
    t.Errorf("Visual comparison failed: %v", err)
}
```

### Creating Baseline Images

```bash
# 1. Generate reference PDF with Microsoft Office
# 2. Convert to PNG with Ghostscript
gs -q -dNOPAUSE -dBATCH -dSAFER -sDEVICE=png16m -r150 \
   -sOutputFile=basic_text.baseline.png reference.pdf
# 3. Place in testdata/fidelity/word/
# 4. Run tests
go test -run TestWordFidelity/BasicText -v
```

### Inspecting Failures

When comparison fails, check generated files:
- `*.diff.png` - Differences highlighted in red
- `*.comparison.png` - Side-by-side view (baseline | generated | diff)

## Success Metrics

✅ All comparison package unit tests pass
✅ Fidelity tests compile and integrate comparison
✅ Graceful skip when Ghostscript unavailable
✅ Graceful skip when baseline missing
✅ Comprehensive documentation
✅ Clean, well-documented code
✅ No new Go dependencies (only system Ghostscript)
✅ Configurable tolerances

## Notes

- The TODO at `pdf/fidelity_test.go:285` has been **REMOVED** and replaced with working implementation
- Visual comparison is **optional** and degrades gracefully
- All code follows Go best practices with comprehensive godoc comments
- Integration is non-breaking (existing tests continue to work)

## Next Steps

To fully enable visual comparison:

1. Create test documents (DOCX/XLSX/PPTX) for fidelity tests
2. Generate reference PDFs using Microsoft Office
3. Create baseline PNGs using the provided Ghostscript command
4. Commit baseline images to repository
5. Run fidelity tests with visual comparison enabled

Once baseline images are in place, the comparison system will automatically:
- Convert generated PDFs to PNG
- Compare with baselines
- Report differences
- Generate diff images on failure
- Catch visual regressions in CI/CD
