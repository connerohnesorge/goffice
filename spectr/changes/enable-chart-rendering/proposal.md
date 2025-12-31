# Change: Enable Chart Rendering Infrastructure

## Why

goffice has **3,543 lines of chart/shape rendering code DISABLED** in 12 `.wip` files:
- `chart_renderer.go.wip` (499 LOC) - Bar, Line, Pie chart rendering with legends/axes
- `shape_renderer.go.wip` (485 LOC) - Shape geometry rendering
- `fill_renderer.go.wip` (307 LOC) - Solid, gradient, pattern fills
- `stroke_renderer.go.wip` (221 LOC) - Line/border rendering
- `effects_renderer.go.wip` (272 LOC) - Shadows, glows, reflections
- `transform_renderer.go.wip` (233 LOC) - Rotation, flip, scale
- `text_in_shape.go.wip` (392 LOC) - Text within shapes
- `image_renderer.go.wip` (168 LOC) - Image/picture rendering
- Plus 4 test files (966 LOC total)

**The code is production-quality but blocked by ONE missing API:** `core.RenderingContext.Page`

The renderers need a `Page` interface with methods like:
- `SetFillColor()`, `SetStrokeColor()`, `SetLineWidth()`, `SetLineDashPattern()`
- `DrawRectangle(x, y, w, h, fill, stroke bool)`, `DrawCircle(cx, cy, r, fill, stroke bool)`
- `SaveGraphicsState()`, `RestoreGraphicsState()`, `Clip()`
- `SetFont(name string, size float64)`, `DrawText()`

**Impact**: 8/10 - Charts are essential for PowerPoint and Excel output. Without rendering, PDF export produces blank areas where charts should appear.

## What Changes

### Phase 1: Core Page API (40-60 hours)
Create the missing PDF page abstraction:

1. **Page Interface** (NEW - 20-25h)
   - Define `Page` interface with drawing primitives
   - Implement for `pdf/core` package
   - Add to `RenderingContext`

2. **Re-enable Renderers** (ENABLE - 8-12h)
   - Rename `.wip` files to `.go`
   - Fix any API mismatches
   - Run existing tests

3. **Integration Testing** (NEW - 12-18h)
   - Chart rendering tests
   - Shape rendering tests
   - Fill/stroke tests

### Phase 2: Chart Type Completion (30-40 hours)
4. **Advanced Charts** (15-20h) - Waterfall, Funnel, Stock/OHLC
5. **Combination Charts** (10-15h) - Multi-type charts
6. **3D Chart Support** (10-15h) - 3D perspective rendering

### Phase 3: Polish (15-25 hours)
7. **Fidelity Testing** (10-15h) - Compare to Excel/PowerPoint output
8. **Documentation** (5-10h) - Renderer API docs

**Breaking changes**: None. This enables disabled code.

## Impact

- **Affected specs**: `pdf-drawing` (new delta spec)
- **New capabilities**: Chart/shape rendering to PDF
- **Affected code**:
  - `pdf/core/page.go` - NEW: Page interface (distinct from existing Page struct in document.go)
  - `pdf/core/pdf_page.go` - NEW: PDFPage implementation wrapping existing Page struct
  - `pdf/core/render_context.go` - ENHANCE: Add Page field to RenderingContext
  - `pdf/drawing/*.go` - ENABLE: Rename from .wip
  - `pdf/presentation/` - ENHANCE: Use chart renderer
  - `pdf/spreadsheet/` - ENHANCE: Use chart renderer

**Note**: The existing `Page` struct in `document.go` is a concrete type for PDF page objects.
The new `Page` interface abstracts drawing operations for testability and future backends.

## Key Design Decisions

1. **Page Interface Pattern**: Abstract page operations behind interface
   - Rationale: Allows mock testing, future backends (SVG, PNG)
   - Pattern: Matches existing `pdf/core` abstractions

2. **Minimal API Surface**: Only methods actually used by .wip files
   - Rationale: Avoid over-engineering
   - Existing code defines requirements

3. **Enable Before Extend**: Get existing code working first
   - Rationale: 3,543 LOC already written and tested conceptually
   - Phase 2 adds new chart types after Phase 1 proven

## Success Criteria

1. ✅ All 12 .wip files renamed to .go and compiling
2. ✅ All existing .wip tests passing
3. ✅ Bar/Line/Pie charts render correctly to PDF
4. ✅ Shapes render with fills, strokes, effects
5. ✅ Text renders within shapes
6. ✅ Images render in DrawingML contexts
7. ✅ pptx-charts example produces visible charts in PDF
8. ✅ No regressions in Word PDF rendering

## Risk Assessment

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Page API design misses edge cases | Medium | Medium | Start with existing .wip requirements |
| Coordinate system issues (EMU vs points) | High | Medium | Use existing transformation code |
| Font rendering in shapes | Medium | High | Leverage existing pdf/font infrastructure |
| Performance on complex charts | Low | Medium | Profile after basic functionality works |

## Specs Changed

- `spectr/changes/enable-chart-rendering/specs/pdf-drawing/spec.md` (new delta)
