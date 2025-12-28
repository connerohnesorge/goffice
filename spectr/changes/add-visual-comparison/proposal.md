# Change: Add Automated Visual Comparison for PDF Fidelity Tests

## Why

The current fidelity test suite (pdf/fidelity_test.go:285) generates PDF output but requires manual visual comparison with Microsoft Office output. This is labor-intensive, error-prone, and blocks automated regression detection. An automated visual comparison system would enable continuous validation of rendering fidelity and catch visual regressions before they propagate.

## What Changes

- Add `pdf/comparison` package with image-based PDF comparison utilities
- Implement PDF-to-image conversion (PNG) using standard library + system tools
- Implement pixel-level difference detection and reporting
- Integrate comparison into existing fidelity test helpers
- Generate visual diff reports (annotated images showing differences)
- Support configurable tolerance levels (ignore minor rendering variations <2pt)

Breaking changes: None. This is purely additive and optional for existing tests.

## Impact

- Affected specs: `pdf-testing` (new capability)
- Affected code: 
  - `pdf/fidelity_test.go` - integrate comparison into test helpers
  - `pdf/comparison/*.go` - new comparison implementation
  - `testdata/fidelity/*/` - test fixtures and baseline images

## Key Design Decisions

1. **PDF to Image**: Use Ghostscript (gs) as the system dependency for PDF→PNG conversion. It's widely available, production-proven, and produces consistent output.

2. **Pixel Comparison**: Implement custom pixel-diff logic (Euclidean color distance) rather than external libraries to keep dependencies minimal (Go std lib + Ghostscript).

3. **Tolerance Strategy**: 
   - Per-pixel tolerance: 2.0 delta E (CIE Lab color space approximation)
   - Per-image threshold: Allow up to 0.5% of pixels to exceed tolerance (fidelity goal is visual equivalence, not pixel-perfect)

4. **Baseline Management**: Store baseline images alongside test files (e.g., `testdata/fidelity/word/basic_text.docx.baseline.png`)

5. **Error Handling**: Generate annotated diff images showing pixel-level differences for manual review when comparison fails.

## Implementation Scope

### Phase 1 (MVP)
- PDF→PNG conversion wrapper
- Pixel-level difference detection
- Integration into Word fidelity tests only (3-5 test cases)
- Basic comparison reporting (pass/fail + diff image)

### Phase 2 (Future)
- Extend to Excel and PowerPoint fidelity tests
- Perceptual hashing for layout-level differences
- Interactive baseline review/approval workflow
- Integration with CI/CD pipeline
