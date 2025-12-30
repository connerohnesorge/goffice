# Change: Enable DrawingML PDF Rendering

## Why

The PDF rendering pipeline has **12 critical renderer files disabled** (`.wip` extension) in `pdf/drawing/` because they depend on a non-existent Page abstraction API in `pdf/core/`. This completely blocks rendering of visual elements in PDF output:

- **Charts** (Excel, PowerPoint, Word) - Cannot render any chart types
- **Shapes** with fills, strokes, and effects - No geometric rendering
- **Images** within shapes - Pictures cannot be placed  
- **Visual effects** (shadows, glows, reflections) - Effects not rendered
- **Gradient and pattern fills** - Only solid colors work
- **Text within shapes** - Text boxes cannot be rendered
- **Geometric transformations** - Rotations, flips, scaling broken

**Root Cause**: All 12 disabled renderers expect `ctx.Page.SetFillColor()`, `ctx.Page.DrawRectangle()`, etc., but `core.RenderingContext` has no `Page` field or these methods.

**Impact**: **Critical (10/10)** - Blocks all DrawingML rendering:
- Excel PDFs: Missing charts (major gap for reporting/dashboards)
- PowerPoint PDFs: Missing charts, shapes, images (unusable for presentations)
- Word PDFs: Missing charts and inline images (moderate impact)

Current state documented in `pdf/drawing/WIP_RENDERERS.md`.

## What Changes

**Core Infrastructure**:
- Add `Page` interface to `pdf/core/` with drawing primitives, state management, transforms
- Add `Page` field to `core.RenderingContext` 
- Implement `Page` interface using pdfcpu backend (existing dependency)
- Add coordinate conversion utilities (EMU → PDF points, top-left → bottom-left origin)

**DrawingML Renderers**:
- Re-enable all 12 `.wip` files by removing extension
- Fix any API mismatches discovered during re-enablement
- Add integration between Page API and existing DrawingML elements

**Testing & Documentation**:
- Add unit tests for Page abstraction with mock implementation
- Add integration tests for chart, shape, image rendering
- Enhance fidelity tests with DrawingML elements
- Update LIMITATIONS.md and FIDELITY.md with new capabilities

**Breaking changes**: None. This enables disabled functionality without changing existing working code.

## Impact

**Affected specs**: 
- `pdf-core` (ADDED Page abstraction requirements)
- `pdf-drawing` (MODIFIED to reflect enabled renderers)

**New capabilities**: 
- Charts render in all PDF outputs
- Shapes with complex fills/strokes render
- Images in shapes display
- Effects (shadows, glows) render
- All DrawingML visual elements supported

**Affected code**:
- `pdf/core/page.go` - NEW: Page interface and implementation
- `pdf/core/render_context.go` - MODIFIED: Add Page field
- `pdf/drawing/*.go.wip` → `*.go` - RE-ENABLED: 8 renderer files
- `pdf/drawing/*_test.go.wip` → `*_test.go` - RE-ENABLED: 4 test files
- `pdf/word/renderer.go` - MODIFIED: Use Page API for charts/images
- `pdf/spreadsheet/renderer.go` - MODIFIED: Use Page API for charts
- `pdf/presentation/renderer.go` - MODIFIED: Use Page API for shapes/charts

## Key Design Decisions

### 1. Page Interface Design
Minimal interface covering essential drawing operations:
```go
type Page interface {
    // Drawing primitives
    DrawRectangle(x, y, width, height float64)
    DrawEllipse(cx, cy, rx, ry float64)  
    DrawPath(path *PathBuilder)
    
    // State management
    SetFillColor(color Color)
    SetStrokeColor(color Color)
    SetLineWidth(width float64)
    SetLineDashPattern(pattern []float64, phase float64)
    
    // Graphics state
    PushState()
    PopState()
    Transform(matrix Matrix)
    
    // Content
    AddImage(img Image, x, y, width, height float64)
    DrawText(text string, x, y float64, font Font, size float64)
}
```

**Rationale**: Minimal API surface matching PDF capabilities, not DrawingML complexity. Higher-level renderers (chart, shape, fill, etc.) consume this interface.

### 2. pdfcpu Integration Strategy
Wrap pdfcpu's content stream API rather than reimplementing PDF generation.

**Rationale**: pdfcpu is already a dependency (used for PDF document structure). Reusing it for page-level drawing avoids duplication and leverages a proven library. Implementation becomes a thin adapter layer.

### 3. Coordinate System Handling
- **Input**: EMU coordinates (DrawingML native: 914,400 EMU = 1 inch)
- **Output**: PDF points (72 points = 1 inch)
- **Origin**: Convert top-left (OOXML) to bottom-left (PDF) at page boundary

**Rationale**: Keeps DrawingML renderers in native units, isolates coordinate conversion to Page implementation. Clear separation of concerns.

### 4. Defer Advanced Features
Phase 1 focuses on basic shapes/charts. Defer to future:
- Clipping paths
- Blend modes beyond normal
- Advanced gradients (mesh/coons)
- Soft masks for transparency

**Rationale**: 80% of use cases covered by solid fills, linear gradients, basic shapes. Advanced features can be added incrementally without API changes.

### 5. Mock Testing Strategy
Page interface enables unit testing renderers without actual PDF generation using mock implementation.

**Rationale**: Fast, isolated tests for renderer logic. Integration tests cover end-to-end with real PDFs.

## Implementation Scope

### Phase 1: Core Infrastructure (2 weeks)
- Design Page interface
- Implement pdfcpu-backed Page
- Add to RenderingContext
- Unit tests with mock Page

### Phase 2: Re-enable Renderers (2 weeks)
- Remove `.wip` extensions (12 files)
- Fix API mismatches
- Compile and basic validation
- Renderer-specific unit tests

### Phase 3: Integration & Testing (1.5 weeks)
- End-to-end PDF generation with charts
- Fidelity tests with visual inspection
- LibreOffice compatibility validation
- Regression test suite

### Phase 4: Documentation (0.5 weeks)
- API documentation for Page interface
- Update LIMITATIONS.md (remove chart/shape limitations)
- Update FIDELITY.md with new rendering quality notes
- Developer guide for DrawingML PDF rendering

**Total: 6 weeks**

### Out of Scope
- Formula calculation for charts (uses cached values)
- SmartArt rendering (requires SmartArt implementation first)
- 3D charts (2D projection acceptable)
- Video/audio embedding (not applicable to PDF)

## Dependencies

**Existing**:
- pdfcpu library (already integrated)
- DrawingML color handling (`pdf/drawing/color.go`)
- Font subsetting (`pdf/font/subset.go`)
- Text layout engine (`pdf/layout/`)

**New**: None

## Success Criteria

- [ ] All 12 `.wip` files renamed to `.go` and compiling
- [ ] Page interface fully implemented with pdfcpu backend
- [ ] Charts render in Excel→PDF, PowerPoint→PDF, Word→PDF
- [ ] Shapes with fills/strokes render correctly
- [ ] Images within shapes display
- [ ] Shadows/glows render (basic effects)
- [ ] Existing PDF tests pass (no regression)
- [ ] New DrawingML integration tests pass
- [ ] Visual fidelity validation with sample documents
