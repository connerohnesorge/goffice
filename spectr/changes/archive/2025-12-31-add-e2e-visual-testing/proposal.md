# Change: Add E2E Visual Testing Framework for PPTX Generation Comparison

## Why

Currently, goffice lacks a systematic way to verify that our Go library's PPTX generation matches the official Microsoft Open-XML-SDK for .NET in terms of **visual output**. While we have unit tests and integration tests that verify XML structure and API behavior, we cannot guarantee that the presentations we generate will render identically to those created by the reference C# implementation.

**Problem:**
- No visual regression testing
- Manual verification required for chart rendering accuracy
- Difficult to catch subtle rendering differences (colors, positioning, fonts, effects)
- No automated comparison against the authoritative .NET SDK
- Risk of divergence from Microsoft's reference implementation

**Opportunity:**
- Build confidence that goffice produces Office-compatible presentations
- Catch visual regressions early in development
- Verify chart rendering across all 11 chart types
- Establish baseline for PDF rendering fidelity
- Create reusable infrastructure for Word/Excel visual testing

## What Changes

This change introduces a **comprehensive Nix-powered end-to-end visual testing framework** in `tests/e2e/` that:

1. **Test Definition System**
   - Declarative test case definitions in Go
   - Parameterized test scenarios (chart types, formatting, layouts)
   - Reusable test primitives for common patterns

2. **Dual Generator Infrastructure**
   - Go generator using goffice presentation API
   - C# generator using Open-XML-SDK
   - Identical test case execution on both platforms
   - Synchronized test data and parameters

3. **Rendering Pipeline**
   - PDF rendering from both PPTX outputs
   - Image extraction from PDFs (PNG @ 300 DPI)
   - Optional LibreOffice/PowerPoint screenshot automation

4. **Visual Comparison Engine**
   - Pixel-perfect image diffing
   - Perceptual diff metrics (SSIM, MSE, PSNR)
   - Configurable tolerance thresholds
   - Highlighting of visual differences

5. **Reporting & CI Integration**
   - HTML report with side-by-side comparisons
   - Visual diff overlays
   - Metrics dashboard
   - GitHub Actions integration for PR checks

6. **Nix Development Environment**
   - Isolated, reproducible test environment
   - Manages .NET SDK, LibreOffice, ImageMagick, etc.
   - One-command setup via flake.nix

### Capabilities Affected

- **NEW**: `e2e-testing` - End-to-end visual testing infrastructure

### Architectural Changes

- Adds `tests/e2e/` directory structure
- Integrates .NET SDK alongside Go toolchain
- Introduces visual diff tooling (ImageMagick, pdfimages, etc.)
- Extends CI/CD pipeline with visual regression checks

### Breaking Changes

**NONE** - This is purely additive testing infrastructure.

## Impact

### Affected Specs
- NEW: `e2e-testing` capability

### Affected Code
- **New directories:**
  - `tests/e2e/` - Test framework root
  - `tests/e2e/framework/` - Core framework code
  - `tests/e2e/generators/` - Go and C# generators
  - `tests/e2e/testcases/` - Test case definitions
  - `tests/e2e/fixtures/` - Test data and expected outputs
  - `tests/e2e/reports/` - Generated comparison reports

- **Modified files:**
  - `flake.nix` - Add .NET SDK, LibreOffice, visual diff tools
  - `.github/workflows/` - Add visual testing workflow
  - `CLAUDE.md` - Document E2E testing conventions

### Dependencies
- .NET SDK 9.0+ (for C# generator)
- Open-XML-SDK NuGet package
- LibreOffice 7.6+ or PowerPoint (for rendering)
- ImageMagick (for image comparison)
- Poppler (pdfimages for PDF → PNG conversion)
- Go 1.25+ (existing)

### Migration Path
**NONE** - Existing tests unaffected. New framework is opt-in for visual verification.

## Success Criteria

1. Framework can execute identical test cases in both Go and C#
2. Framework generates visual comparison reports automatically
3. All 11 chart types can be tested with visual diff validation
4. CI pipeline runs visual tests on PRs (with caching)
5. Developer can add new test cases with <20 lines of code
6. Test execution completes in <5 minutes for full suite
7. Visual diff false positive rate <5% (configurable thresholds)

## Risks & Mitigations

**Risk:** Rendering differences due to LibreOffice vs. PowerPoint
**Mitigation:** Support both renderers; document known differences; use perceptual metrics

**Risk:** Flaky tests due to font rendering variations
**Mitigation:** Embed fonts in presentations; use Docker for consistent environment

**Risk:** Large binary artifacts in git (PDFs, PNGs)
**Mitigation:** Store only test definitions in git; generate artifacts in CI; use Git LFS if needed

**Risk:** Slow test execution blocking PR velocity
**Mitigation:** Implement smart caching; run subset on commit, full suite nightly

## Open Questions

- [ ] Should we use LibreOffice headless or automate PowerPoint via COM?
- [ ] What perceptual diff threshold provides best signal-to-noise?
- [ ] Should visual tests be required for PR merge or informational only?
- [ ] Do we need baseline images checked into git or regenerate from C# each time?
