# Tasks: Add Automated Visual Comparison for PDF Fidelity Tests

## 1. Implementation

### 1.1 PDF-to-PNG Conversion (pdf/comparison/convert.go)
- [x] Create `pdf/comparison` package
- [x] Implement `ConvertPDFToPNG(pdfBytes []byte, outputPath string) error`
- [x] Shell out to Ghostscript (gs) with hardcoded parameters (-r150, png16m)
- [x] Handle gs not found (return descriptive error, don't panic)
- [x] Validate PNG output (size > 0, readable)
- [x] Add comprehensive error messages

### 1.2 Pixel-Level Difference Detection (pdf/comparison/diff.go)
- [x] Implement `CompareImages(baselineImagePath, generatedImagePath string) (*DiffResult, error)`
- [x] Define `DiffResult` struct (TotalPixels, DiffPixels, DiffPercentage, PixelsExceedingTolerance)
- [x] Load both PNG images using image/png
- [x] Implement pixel-by-pixel comparison with Euclidean color distance
- [x] Calculate per-pixel tolerance (2.0 delta per channel / 255)
- [x] Calculate per-image threshold (0.5% of pixels allowed to exceed tolerance)
- [x] Return detailed diff result with statistics

### 1.3 Difference Visualization (pdf/comparison/annotate.go)
- [x] Implement `AnnotateDiffImage(baselineImagePath, generatedImagePath, outputDiffPath string) error`
- [x] Create output image with same dimensions as baseline
- [x] Highlight differing pixels in red (or high-contrast color)
- [x] Optionally overlay pixel deltas or heatmap
- [x] Save annotated image as PNG
- [x] Ensure annotations are visible but don't obscure underlying image

### 1.4 Type Definitions (pdf/comparison/types.go)
- [x] Define `DiffResult` struct with metadata (total pixels, diff pixels, percentage, pixels exceeding tolerance)
- [x] Define `ComparisonConfig` struct (pixel tolerance, image threshold, output paths)
- [x] Add helper methods (IsWithinTolerance, DiffPercentageString, etc.)

### 1.5 Integration with Fidelity Tests (pdf/fidelity_test.go)
- [x] Update `runWordFidelityTest` to call comparison after PDF generation
- [x] Create `compareWithBaseline(t *testing.T, generatedPDF []byte, baselineImagePath string, config *comparison.ComparisonConfig) bool`
- [x] Skip comparison gracefully if baseline missing (log warning, don't fail)
- [x] Skip comparison if Ghostscript unavailable (log warning, don't fail)
- [x] Call comparison.AnnotateDiffImage to generate diff images on failure
- [x] Update test logging to include diff results

### 1.6 Baseline Image Generation
- [ ] Generate baseline images for Word fidelity tests:
  - [ ] testdata/fidelity/word/basic_text.docx.baseline.png
  - [ ] testdata/fidelity/word/complex_formatting.docx.baseline.png
  - [ ] testdata/fidelity/word/tables.docx.baseline.png
- [ ] Baseline images should be committed to git (can use git-lfs if needed)
- [ ] Document baseline creation process in README

## 2. Testing

### 2.1 Unit Tests (pdf/comparison/comparison_test.go)
- [x] Test ConvertPDFToPNG with valid PDF input
- [x] Test ConvertPDFToPNG with missing gs (graceful error)
- [x] Test ConvertPDFToPNG output validation
- [x] Test CompareImages with identical images (0% diff)
- [x] Test CompareImages with slightly different images (within tolerance)
- [x] Test CompareImages with very different images (exceeds tolerance)
- [x] Test tolerance thresholds (pixel tolerance, image threshold)
- [x] Test AnnotateDiffImage output (file created, valid PNG)
- [x] Create synthetic test images (colored rectangles, simple patterns)

### 2.2 Integration Tests (pdf/fidelity_test.go)
- [x] Word fidelity test with baseline comparison (should pass)
- [x] Word fidelity test with modified PDF (should fail + generate diff)
- [x] Verify diff image is created on comparison failure
- [x] Verify graceful skip when baseline missing
- [x] Verify graceful skip when gs unavailable

### 2.3 Manual Testing
- [ ] Generate PDF from test document manually
- [ ] Run comparison and verify diff image quality
- [ ] Visually inspect diff images for accuracy
- [ ] Verify tolerance thresholds are appropriate (not too strict, not too loose)

## 3. Documentation

### 3.1 Code Documentation
- [x] Document all exported functions and types with godoc comments
- [x] Document tolerance parameters and rationale
- [x] Document baseline image naming convention
- [x] Add examples to comparison.go for usage

### 3.2 User-Facing Documentation
- [x] Add section to pdf/README.md on running fidelity tests with comparison
- [x] Document Ghostscript installation/setup
- [x] Document baseline image management (regeneration, updates)
- [x] Add troubleshooting guide (gs not found, image format issues, etc.)

### 3.3 Design Documentation
- [x] Update pdf/FIDELITY.md to reference automated comparison capability
- [x] Document tolerance thresholds and why they were chosen
- [x] Provide guidance on when to update baselines

## 4. Deliverables

- [x] `pdf/comparison/convert.go` - PDF→PNG conversion
- [x] `pdf/comparison/diff.go` - Pixel difference detection
- [x] `pdf/comparison/annotate.go` - Diff visualization
- [x] `pdf/comparison/types.go` - Type definitions
- [x] `pdf/comparison/comparison_test.go` - Unit and integration tests
- [ ] Baseline images for 3 Word test cases
- [x] Updated `pdf/fidelity_test.go` with comparison integration
- [x] Updated documentation (README, FIDELITY.md)

## 5. Success Criteria

- [ ] All unit tests pass
- [ ] Word fidelity tests pass with baseline comparison
- [ ] Diff images are generated and visible on failures
- [ ] Comparison gracefully skips when dependencies missing
- [ ] Code is properly documented with godoc
- [ ] No new external dependencies (Ghostscript is system dependency, not Go)
- [ ] Pixel tolerance parameters are reasonable (validated via manual inspection)
- [ ] Baseline images are committed and tracked in git

## Approval Gate

Do not proceed with implementation until this proposal is reviewed and approved.
