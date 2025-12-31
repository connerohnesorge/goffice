# Design: Enable Chart Rendering Infrastructure

## Context

goffice has **3,543 lines of disabled rendering code** in `pdf/drawing/*.wip` files that is blocked by a missing `Page` abstraction in `core.RenderingContext`.

**Existing (DISABLED - In .wip Files):**
- `chart_renderer.go.wip` - RenderBarChart, RenderLineChart, RenderPieChart
- `shape_renderer.go.wip` - RenderShape with geometry handling
- `fill_renderer.go.wip` - RenderSolidFill, RenderGradientFill
- `stroke_renderer.go.wip` - RenderStroke with line patterns
- `effects_renderer.go.wip` - RenderShadow, RenderGlow
- `transform_renderer.go.wip` - RenderTransform
- `text_in_shape.go.wip` - RenderTextInShape
- `image_renderer.go.wip` - RenderImage

**Missing (BUILD THIS):**
- `pdf/core/page.go` - Page interface with drawing primitives
- `core.RenderingContext.Page` field

The solution must **enable existing code**, not rewrite it.

## Goals

1. **Define Page interface** - Abstract drawing operations
2. **Implement Page for PDF** - Connect to pdf/core
3. **Enable .wip files** - Rename and fix API mismatches
4. **Run existing tests** - Validate the 966 LOC of test code

## Non-Goals

- Rewrite the chart renderer (it works)
- Add new chart types before existing ones work
- SVG or PNG backends (future phase)
- Perfect Excel/PowerPoint fidelity (good enough for v1)

## Architecture Overview

### Existing Code (IN .wip FILES)

```
pdf/drawing/
├── chart_renderer.go.wip    (499 LOC) - Bar/Line/Pie charts
├── shape_renderer.go.wip    (485 LOC) - Shape geometry
├── fill_renderer.go.wip     (307 LOC) - Fill rendering
├── stroke_renderer.go.wip   (221 LOC) - Stroke rendering
├── effects_renderer.go.wip  (272 LOC) - Effects
├── transform_renderer.go.wip(233 LOC) - Transforms
├── text_in_shape.go.wip     (392 LOC) - Text in shapes
├── image_renderer.go.wip    (168 LOC) - Images
├── *_test.go.wip            (966 LOC) - Tests
└── WIP_RENDERERS.md         - Documents the blocking issue
```

### New Code (ADD THIS)

```
pdf/core/
└── page.go          // NEW: Page interface (~200 LOC)

pdf/drawing/
└── (rename all .wip → .go)
```

### Enhancement to Existing Code

```go
// In pdf/core/render_context.go - ADD field:
type RenderingContext struct {
    // existing fields (pageWidth, pageHeight, originStack, transformStack, etc.)...
    Page Page  // NEW: Page interface for drawing
}
```

## Detailed Design Decisions

### Decision 1: Page Interface Definition

**Design**: Define interface based on what .wip files actually use:

```go
// pdf/core/page.go

// Page provides drawing primitives for PDF page rendering
type Page interface {
    // Graphics state
    SaveGraphicsState()
    RestoreGraphicsState()

    // Colors
    SetFillColor(r, g, b float64)
    SetFillColorRGBA(r, g, b, a float64)
    SetStrokeColor(r, g, b float64)
    SetStrokeColorRGBA(r, g, b, a float64)

    // Line properties
    SetLineWidth(width float64)
    SetLineCap(cap int)        // cap values: 0=butt, 1=round, 2=square
    SetLineJoin(join int)      // join values: 0=miter, 1=round, 2=bevel
    SetLineDashPattern(pattern []float64, phase float64)

    // Shapes
    DrawRectangle(x, y, w, h float64, fill, stroke bool)
    DrawRoundedRectangle(x, y, w, h, radius float64)
    DrawCircle(cx, cy, r float64, fill, stroke bool)
    DrawEllipse(cx, cy, rx, ry float64)
    DrawLine(x1, y1, x2, y2 float64)
    DrawPath(path *Path)

    // Fill/Stroke
    Fill()
    Stroke()
    FillStroke()

    // Text
    SetFont(name string, size float64)
    DrawText(x, y float64, text string)
    DrawTextRotated(x, y, angle float64, text string)

    // Images
    DrawImage(img *Image, x, y, w, h float64)

    // Transformations
    Translate(dx, dy float64)
    Scale(sx, sy float64)
    Rotate(angle float64)

    // Clipping
    Clip()                     // Apply current path as clipping path
    SetClipRect(x, y, w, h float64)
    ClearClip()

    // Content stream
    WriteContent(content string)
}

// Supporting types used by the Page interface
type Path = *PathBuilder    // Path is an alias for PathBuilder from pdf/drawing
type Font struct { ... }    // Font resource (existing in pdf/font)
type Image struct { ... }   // Image resource (existing in pdf/core)

// Helper functions already exist in pdf/drawing:
// - NewRGB(r, g, b float64) Color
// - ParseColor(hex string) Color
// - NewPathBuilder() *PathBuilder
```

**Rationale**:
- Interface methods derived from actual .wip file calls
- Standard PDF graphics model operations
- Matches pdf/core existing patterns

### Decision 2: PDF Page Implementation

**Design**: Implement Page for PDF content streams:

```go
// pdf/core/pdf_page.go

type PDFPage struct {
    doc      *Document
    pageObj  *Object
    content  *ContentStream
    fonts    map[string]*Font
    images   map[string]*Image
}

func NewPDFPage(doc *Document, width, height float64) *PDFPage {
    // Create page object
    // Initialize content stream
    // Set up resources
}

func (p *PDFPage) SetFillColor(r, g, b float64) {
    p.content.WriteOperator("rg", r, g, b)
}

func (p *PDFPage) DrawRectangle(x, y, w, h float64, fill, stroke bool) {
    p.content.WriteOperator("re", x, y, w, h)
    if fill && stroke {
        p.content.WriteOperator("B")
    } else if fill {
        p.content.WriteOperator("f")
    } else if stroke {
        p.content.WriteOperator("S")
    }
}

// ... implement all interface methods
```

**Rationale**:
- Wraps existing pdf/core primitives
- Manages content stream writing
- Handles resource dictionary

### Decision 3: Context Integration

**Design**: Add Page to RenderingContext:

```go
// pdf/core/render_context.go

// RenderingContext already exists with these fields:
// - pageWidth, pageHeight float64
// - originStack []Point
// - transformStack []PDFMatrix
// - contentBoundsStack []Rectangle
// - currentMargins Margins
// - coordTransformer *CoordTransformer

// ADD this field:
type RenderingContext struct {
    // ... existing fields ...
    Page Page  // NEW: Page interface for drawing operations
}

// SetPage sets the current page for rendering
func (ctx *RenderingContext) SetPage(page Page) {
    ctx.Page = page
}
```

**Relationship to Existing Page Struct**:
- `pdf/core/document.go` has a concrete `Page` struct (page object, content buffer, resources)
- The NEW `Page` interface abstracts drawing operations (SetFillColor, DrawRectangle, etc.)
- `PDFPage` implementation wraps the concrete `Page` struct and translates interface calls to PDF operators

**Rationale**:
- Minimal change to existing context
- Page is set per-page during rendering
- Interface enables mock testing and future backends (SVG, PNG)

### Decision 4: Enable Strategy

**Design**: Rename files, fix API calls, run tests:

```bash
# Step 1: Rename all .wip files
for f in pdf/drawing/*.go.wip; do
    mv "$f" "${f%.wip}"
done

# Step 2: Fix any API mismatches (compile errors)
# - Update ctx.Page references
# - Fix any type mismatches

# Step 3: Run tests
go test ./pdf/drawing/...
```

**Rationale**:
- Existing code is complete and conceptually tested
- Just needs API connection
- Tests validate functionality

## Coordinate System

**PDF uses points (72 points = 1 inch)**
**DrawingML uses EMUs (914400 EMUs = 1 inch)**

The existing .wip code includes transformation:
```go
func emuToPoints(emu int64) float64 {
    return float64(emu) / 914400.0 * 72.0
}
```

This is already implemented in the disabled code.

## Critical API Details from WIP Files

The .wip files use these EXACT signatures (confirmed by reading actual code):

1. **DrawRectangle**: `DrawRectangle(x, y, w, h float64, fill, stroke bool)` - 6 params
   - Used in chart_renderer.go.wip lines 130-137, 423-430, 441-448
   - The fill/stroke bools control whether to fill and/or stroke the shape

2. **DrawCircle**: `DrawCircle(cx, cy, r float64, fill, stroke bool)` - 5 params
   - Used in chart_renderer.go.wip lines 225-231, 240-246
   - The fill/stroke bools control whether to fill and/or stroke the shape

3. **SetLineDashPattern**: NOT `SetDashPattern`
   - Used in stroke_renderer.go.wip line 70
   - Signature: `SetLineDashPattern(pattern []float64, phase float64)`

4. **Clip**: Method exists (commented out in image_renderer.go.wip line 91)
   - Signature: `Clip()`
   - Applies current path as clipping path

5. **SetFont**: Takes string name, NOT *Font pointer
   - Used in chart_renderer.go.wip lines 452, 491
   - Signature: `SetFont(name string, size float64)`
   - Example: `ctx.Page.SetFont("Helvetica", 10)`

6. **Color helpers**: Use NewRGB and ParseColor
   - NewRGB(r, g, b float64) Color - lines 467, 470, etc in chart_renderer.go.wip
   - ParseColor(hex string) Color - used in fill_renderer.go.wip line 49

## Module Organization

```
pdf/core/
├── document.go       // EXISTING: PDF document + Page struct (concrete)
├── page.go           // NEW: Page interface (~100 LOC)
├── pdf_page.go       // NEW: PDFPage implementation (~300 LOC)
├── render_context.go // ENHANCE: Add Page field to RenderingContext

pdf/drawing/
├── chart_renderer.go    // ENABLE: Rename from .wip
├── shape_renderer.go    // ENABLE: Rename from .wip
├── fill_renderer.go     // ENABLE: Rename from .wip
├── stroke_renderer.go   // ENABLE: Rename from .wip
├── effects_renderer.go  // ENABLE: Rename from .wip
├── transform_renderer.go // ENABLE: Rename from .wip
├── text_in_shape.go     // ENABLE: Rename from .wip
├── image_renderer.go    // ENABLE: Rename from .wip
└── *_test.go            // ENABLE: Rename from .wip
```

**Total new code**: ~400 LOC (Page interface + implementation)
**Enabled code**: 3,543 LOC (existing .wip files)

## Testing Strategy

### Phase 1: Enable Existing Tests

The .wip test files already contain:
- chart_renderer_test.go.wip (174 LOC)
- shape_renderer_test.go.wip (204 LOC)
- fill_renderer_test.go.wip (184 LOC)
- drawingml_integration_test.go.wip (404 LOC)

These tests should pass once the Page API is implemented.

### Phase 2: Integration Tests

- [ ] Render pptx-charts example to PDF
- [ ] Verify bar chart renders with legend and axes
- [ ] Verify pie chart renders with slices and labels
- [ ] Verify line chart renders with markers
- [ ] Compare to LibreOffice output visually

### Phase 3: Fidelity Tests

- [ ] Create reference images from Excel/PowerPoint
- [ ] Compare rendered output using image diff
- [ ] Document known differences

## API Surface Summary

### New Interface: Page

```go
type Page interface {
    // See complete list in Decision 1
    SaveGraphicsState()
    RestoreGraphicsState()
    SetFillColor(r, g, b float64)
    SetStrokeColor(r, g, b float64)
    DrawRectangle(x, y, w, h float64)
    DrawCircle(cx, cy, r float64)
    DrawLine(x1, y1, x2, y2 float64)
    DrawPath(path *Path)
    Fill()
    Stroke()
    SetFont(font *Font, size float64)
    DrawText(x, y float64, text string)
    DrawImage(img *Image, x, y, w, h float64)
    // ... etc
}
```

### New Method on RenderingContext

```go
func (ctx *RenderingContext) SetPage(page Page)
```

### Enabled Functions (from .wip files)

```go
// chart_renderer.go
func RenderBarChart(ctx *RenderingContext, chart *drawingml.BarChart, bounds Rect)
func RenderLineChart(ctx *RenderingContext, chart *drawingml.LineChart, bounds Rect)
func RenderPieChart(ctx *RenderingContext, chart *drawingml.PieChart, bounds Rect)
func RenderChartTitle(ctx *RenderingContext, title string, bounds Rect)

// shape_renderer.go
func RenderShape(ctx *RenderingContext, shape *drawingml.Shape, bounds Rect)

// fill_renderer.go
func RenderSolidFill(ctx *RenderingContext, fill *drawingml.SolidFill)
func RenderGradientFill(ctx *RenderingContext, fill *drawingml.GradientFill, bounds Rect)

// stroke_renderer.go
func RenderStroke(ctx *RenderingContext, outline *drawingml.Outline)

// effects_renderer.go
func RenderShadow(ctx *RenderingContext, shadow *drawingml.Shadow)
func RenderGlow(ctx *RenderingContext, glow *drawingml.Glow)

// transform_renderer.go
func RenderTransform(ctx *RenderingContext, transform *drawingml.Transform)

// text_in_shape.go
func RenderTextInShape(ctx *RenderingContext, text *drawingml.TextBody, bounds Rect)

// image_renderer.go
func RenderImage(ctx *RenderingContext, image *drawingml.Blip, bounds Rect)
```

## Success Metrics

1. **Compilation**: All .wip files compile after rename
2. **Tests pass**: All 966 LOC of tests pass
3. **Charts visible**: pptx-charts example produces visible charts
4. **No regressions**: Word PDF rendering still works
5. **Performance**: Chart rendering <100ms per chart

## Phase 2 Extensions (NOT IN THIS CHANGE)

- Waterfall, Funnel, Stock chart types
- 3D chart perspective rendering
- Combination charts
- SVG backend
- Higher fidelity compared to Office output
