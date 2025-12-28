# Design: Automated Visual Comparison for PDF Fidelity Testing

## Context

The goffice PDF rendering implementation needs validation that generated PDFs match Microsoft Office output visually. Manual comparison is impractical for regression testing. The solution must be:
- Automated and reproducible
- Minimal external dependencies (aligned with project philosophy)
- Accurate for catching real regressions
- Tolerant of intentional rendering differences (<2pt variations)

## Goals

- Enable automated comparison of generated PDFs against baseline images
- Support pixel-level and perceptual difference detection
- Integrate seamlessly into existing test framework
- Minimize external dependencies (only Ghostscript)
- Generate actionable diff reports for manual review

## Non-Goals

- Exact byte-for-byte PDF comparison (not visual)
- Complex ML-based image analysis
- Web-based diff viewer (initially)
- Automatic baseline approval (requires manual review)

## Technical Decisions

### PDF-to-PNG Conversion
**Decision**: Use system Ghostscript (gs) for PDF→PNG conversion.

**Alternatives considered**:
1. Go native PDF parsing library (pdfium-go) - adds CGO dependency, contraints project philosophy
2. ImageMagick/convert - simpler interface but less control over rendering parameters
3. Cloud API (e.g., CloudConvert) - requires external service, impractical for local testing

**Rationale**: Ghostscript is:
- Industry standard for PDF rendering (used by many PDF viewers/printers)
- Available on all major platforms (Linux, macOS, Windows)
- Produces consistent, high-quality PNG output
- Controlled via command-line (no CGO needed)
- Open-source and free

**Implementation**: Wrapper in `pdf/comparison/convert.go` that shells out to `gs` with hardcoded parameters:
```
gs -q -dNOPAUSE -dBATCH -dSAFER -sDEVICE=png16m -r150 -sOutputFile=output.png input.pdf
```

### Pixel Comparison Strategy
**Decision**: Implement custom pixel-diff using Euclidean color distance in RGB space, with per-pixel and per-image tolerances.

**Alternatives considered**:
1. Exact pixel matching - too strict, fails on minor rendering variations
2. SSIM (Structural Similarity) - better perceptual, requires math library
3. Delta E (CIE Lab) - more perceptually accurate, requires color conversion

**Rationale**: 
- RGB Euclidean distance is simple, fast, and sufficient for detecting real differences
- Tolerance of 2.0/255 per color channel catches visual differences at ~1% color variation
- Per-pixel and per-image thresholds allow tuning sensitivity
- Easy to debug (pixel differences are visualizable)

**Implementation in `pdf/comparison/diff.go`**:
```go
const (
    // Per-pixel color distance tolerance (0-255 per channel, Euclidean)
    pixelTolerance = 2.0
    
    // Per-image threshold: allow up to 0.5% of pixels to exceed tolerance
    imageDiffThreshold = 0.005
)
```

### Baseline Image Management
**Decision**: Store baseline images as `.baseline.png` files in the same directory as test fixtures.

**File structure**:
```
testdata/fidelity/word/
├── basic_text.docx
├── basic_text.docx.baseline.png
├── complex_formatting.docx
├── complex_formatting.docx.baseline.png
```

**Rationale**:
- Collocated with test data for easy management
- Clear naming convention (`[fixture].baseline.png`)
- Easy to update baselines (regenerate and commit)
- Version controlled alongside test data

**Baseline Update Workflow** (future):
```bash
# Regenerate baselines after verification
go test -run TestWordFidelity -update-baselines
```

### Difference Visualization
**Decision**: Generate annotated PNG diff images showing pixel-level differences.

**Implementation in `pdf/comparison/annotate.go`**:
- Highlight differing pixels in red
- Show pixel coordinates and color delta
- Include metadata (total diff %, threshold)

**Example output**: `testdata/fidelity/word/basic_text.docx.diff.png`

## Architecture

```
pdf/comparison/
├── convert.go        # PDF→PNG using gs
├── diff.go           # Pixel-level comparison
├── annotate.go       # Diff image generation
├── types.go          # DiffResult, ComparisonConfig types
└── comparison_test.go
```

**Integration point**: Modify `fidelity_test.go` helpers:
```go
func compareWithBaseline(t *testing.T, generatedPDF []byte, baselineImagePath string) error
```

## Risks & Mitigations

| Risk | Mitigation |
|------|-----------|
| Ghostscript not installed | Graceful skip with helpful error message (don't break CI) |
| PDF rendering variations across OS | Lock gs version, document in README |
| Baseline images diverge from reality | Annual manual audit, diff thresholds |
| Large baseline images in repo | Compress PNG (lossless), store in git-lfs if needed |
| False positives on legit rendering changes | Configurable tolerance, manual review of diffs |

## Migration Plan

### Phase 1 (MVP)
1. Implement `pdf/comparison` package (convert + diff + annotate)
2. Add baseline images for 3 Word test cases (basic_text, complex_formatting, tables)
3. Update `runWordFidelityTest` to call comparison
4. Document usage and baseline management

### Phase 2 (Future)
1. Extend to Excel and PowerPoint fidelity tests
2. Add perceptual hashing for layout-level differences
3. Build baseline approval workflow
4. Integrate into CI/CD with artifact storage

### Backward Compatibility
- Comparison is optional (gracefully skip if baseline missing or gs unavailable)
- Existing tests continue to work (just generate PDFs, no comparison)
- No API changes to public test functions

## Open Questions

1. **Baseline versioning**: Should baselines differ by Office version tested? (e.g., `baseline-2019.png`, `baseline-365.png`)
2. **Tolerance tuning**: Are 2.0 per-pixel delta and 0.5% image threshold appropriate? (Need empirical data)
3. **Git-LFS**: Should large baseline images use git-lfs to keep repo size reasonable?
4. **CI integration**: Should CI auto-fail on visual diffs, or just report?

## Testing Strategy

**Unit tests**:
- `comparison_test.go`: Test diff detection on synthetic images
- Test tolerance thresholds (exact match, slight diff, large diff)
- Test PNG generation and reading

**Integration tests**:
- Real Word→PDF→PNG→Comparison pipeline
- Verify baseline matching works end-to-end
- Verify diff detection on intentional changes

**Acceptance criteria**:
- Word fidelity tests compare output and pass with correct baselines
- Test fails when PDF differs from baseline (e.g., rendering bug injected)
- Diff images are generated and saved for manual review
- Graceful handling when gs unavailable or baseline missing
