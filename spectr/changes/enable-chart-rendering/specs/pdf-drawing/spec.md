# Delta Spec: Enable Chart Rendering Infrastructure

**Status**: ENABLE existing disabled code
**Base Spec**: `spectr/specs/pdf-core/spec.md` (provides PDF document primitives)
**Delta**: Adds Page interface and enables 3,543 LOC of disabled rendering code

## Baseline Acknowledgment

goffice has **existing chart/shape rendering code DISABLED** (~3,543 LOC in `pdf/drawing/*.wip` files) that this spec **enables**:
- `chart_renderer.go.wip` (499 LOC) - Bar, Line, Pie chart rendering
- `shape_renderer.go.wip` (485 LOC) - Shape geometry rendering
- `fill_renderer.go.wip` (307 LOC) - Solid, gradient fills
- `stroke_renderer.go.wip` (221 LOC) - Stroke/line rendering
- `effects_renderer.go.wip` (272 LOC) - Shadows, glows
- `transform_renderer.go.wip` (233 LOC) - Rotation, scale, flip
- `text_in_shape.go.wip` (392 LOC) - Text within shapes
- `image_renderer.go.wip` (168 LOC) - Image rendering
- Plus test files (966 LOC)

**Blocking Issue**: Missing `core.RenderingContext.Page` abstraction

## ADDED Requirements

### Requirement: Page Interface

The system SHALL provide a Page interface for abstract drawing operations.

#### Scenario: Graphics state management
- GIVEN a Page implementation
- WHEN `page.SaveGraphicsState()` is called
- THEN the current graphics state is pushed to a stack
- AND `page.RestoreGraphicsState()` pops and restores it

#### Scenario: Set fill color RGB
- GIVEN a Page implementation
- WHEN `page.SetFillColor(1.0, 0.0, 0.0)` is called
- THEN subsequent fill operations use red color
- AND the PDF content stream contains "1 0 0 rg"

#### Scenario: Set fill color with alpha
- GIVEN a Page implementation
- WHEN `page.SetFillColorRGBA(1.0, 0.0, 0.0, 0.5)` is called
- THEN subsequent fill operations use red with 50% opacity
- AND an ExtGState resource is created for transparency

#### Scenario: Set stroke color
- GIVEN a Page implementation
- WHEN `page.SetStrokeColor(0.0, 0.0, 1.0)` is called
- THEN subsequent stroke operations use blue color
- AND the PDF content stream contains "0 0 1 RG"

#### Scenario: Set line width
- GIVEN a Page implementation
- WHEN `page.SetLineWidth(2.5)` is called
- THEN subsequent strokes use 2.5 point line width
- AND the PDF content stream contains "2.5 w"

#### Scenario: Set line dash pattern
- GIVEN a Page implementation
- WHEN `page.SetLineDashPattern([]float64{3, 2}, 0)` is called
- THEN subsequent strokes use dash pattern (3 on, 2 off)
- AND the PDF content stream contains "[3 2] 0 d"

#### Scenario: Set line cap
- GIVEN a Page implementation
- WHEN `page.SetLineCap(1)` is called (1 = round cap)
- THEN subsequent strokes use round cap style
- AND the PDF content stream contains "1 J"

#### Scenario: Set line join
- GIVEN a Page implementation
- WHEN `page.SetLineJoin(2)` is called (2 = bevel join)
- THEN subsequent strokes use bevel join style
- AND the PDF content stream contains "2 j"

### Requirement: Shape Drawing

The system SHALL provide primitive shape drawing operations.

#### Scenario: Draw rectangle with fill and stroke
- GIVEN a Page implementation
- WHEN `page.DrawRectangle(100, 200, 50, 30, true, true)` is called
- THEN a rectangle is drawn and both filled and stroked
- AND the PDF content stream contains "100 200 50 30 re" followed by "B"

#### Scenario: Draw rectangle with fill only
- GIVEN a Page implementation
- WHEN `page.DrawRectangle(100, 200, 50, 30, true, false)` is called
- THEN a rectangle is drawn and filled but not stroked
- AND the PDF content stream contains "100 200 50 30 re" followed by "f"

#### Scenario: Draw rounded rectangle
- GIVEN a Page implementation
- WHEN `page.DrawRoundedRectangle(100, 200, 50, 30, 5)` is called
- THEN a rectangle with 5-point corner radius is drawn
- AND the path uses Bezier curves for corners

#### Scenario: Draw circle with fill and stroke
- GIVEN a Page implementation
- WHEN `page.DrawCircle(150, 150, 25, true, true)` is called
- THEN a circle centered at (150, 150) with radius 25 is drawn
- AND the circle is both filled and stroked
- AND the path uses Bezier approximation (4 curves)

#### Scenario: Draw ellipse
- GIVEN a Page implementation
- WHEN `page.DrawEllipse(150, 150, 30, 20)` is called
- THEN an ellipse with horizontal radius 30 and vertical radius 20 is drawn

#### Scenario: Draw line
- GIVEN a Page implementation
- WHEN `page.DrawLine(10, 10, 100, 100)` is called
- THEN a line from (10, 10) to (100, 100) is drawn
- AND the PDF content stream contains "10 10 m 100 100 l"

#### Scenario: Draw complex path
- GIVEN a Page implementation and a Path with multiple segments
- WHEN `page.DrawPath(path)` is called
- THEN the complete path is rendered to the content stream

### Requirement: Fill and Stroke Operations

The system SHALL provide fill and stroke painting operations.

#### Scenario: Fill path
- GIVEN a Page with a rectangle path
- WHEN `page.Fill()` is called
- THEN the path is filled with current fill color
- AND the PDF content stream contains "f"

#### Scenario: Stroke path
- GIVEN a Page with a rectangle path
- WHEN `page.Stroke()` is called
- THEN the path is stroked with current stroke color and line width
- AND the PDF content stream contains "S"

#### Scenario: Fill and stroke path
- GIVEN a Page with a rectangle path
- WHEN `page.FillStroke()` is called
- THEN the path is both filled and stroked
- AND the PDF content stream contains "B"

### Requirement: Text Drawing

The system SHALL provide text drawing operations.

#### Scenario: Set font
- GIVEN a Page implementation
- WHEN `page.SetFont("Helvetica", 12.0)` is called
- THEN subsequent text uses the Helvetica font at 12 points
- AND the font is registered in the page resources

#### Scenario: Draw text
- GIVEN a Page with font set
- WHEN `page.DrawText(100, 200, "Hello")` is called
- THEN "Hello" is drawn at position (100, 200)
- AND text is wrapped in BT/ET operators

#### Scenario: Draw rotated text
- GIVEN a Page with font set
- WHEN `page.DrawTextRotated(100, 200, 45.0, "Rotated")` is called
- THEN "Rotated" is drawn at (100, 200) rotated 45 degrees
- AND a transformation matrix is applied

### Requirement: Image Drawing

The system SHALL provide image drawing operations.

#### Scenario: Draw image
- GIVEN a Page implementation and an Image resource
- WHEN `page.DrawImage(img, 100, 200, 150, 100)` is called
- THEN the image is drawn at (100, 200) scaled to 150x100 points
- AND an XObject reference is written to content stream

### Requirement: Transformations

The system SHALL provide coordinate transformation operations.

#### Scenario: Translate
- GIVEN a Page implementation
- WHEN `page.Translate(50, 100)` is called
- THEN subsequent drawing is offset by (50, 100)
- AND a "cm" operator with translation matrix is written

#### Scenario: Scale
- GIVEN a Page implementation
- WHEN `page.Scale(2.0, 0.5)` is called
- THEN subsequent drawing is scaled 2x horizontally and 0.5x vertically

#### Scenario: Rotate
- GIVEN a Page implementation
- WHEN `page.Rotate(90.0)` is called
- THEN subsequent drawing is rotated 90 degrees

### Requirement: Clipping

The system SHALL provide clipping operations.

#### Scenario: Apply clipping path
- GIVEN a Page implementation with a path drawn
- WHEN `page.Clip()` is called
- THEN the current path becomes the clipping path
- AND subsequent drawing is clipped to that path
- AND the "W n" operator is used

#### Scenario: Set clip rectangle
- GIVEN a Page implementation
- WHEN `page.SetClipRect(50, 50, 200, 200)` is called
- THEN subsequent drawing is clipped to that rectangle
- AND the "W" operator is used

#### Scenario: Clear clip
- GIVEN a Page with a clip rectangle set
- WHEN `page.ClearClip()` is called
- THEN clipping is removed (via graphics state restore)

### Requirement: RenderingContext Integration

The system SHALL integrate Page with the existing RenderingContext.

#### Scenario: Context has Page field
- GIVEN a RenderingContext
- WHEN `ctx.Page` is accessed
- THEN it returns the current Page interface (or nil)

#### Scenario: Set page on context
- GIVEN a RenderingContext and a PDFPage
- WHEN `ctx.SetPage(page)` is called
- THEN `ctx.Page` returns the provided page
- AND renderers can use `ctx.Page.DrawRectangle(...)` etc.

### Requirement: Chart Rendering (Enabled)

The system SHALL render charts via the enabled chart_renderer.go.

#### Scenario: Render bar chart
- GIVEN a RenderingContext with Page set and a BarChart with data
- WHEN `RenderBarChart(ctx, chart, bounds)` is called
- THEN bars are drawn with correct heights proportional to data values
- AND bar colors match chart series colors
- AND no error is returned

#### Scenario: Render line chart
- GIVEN a RenderingContext with Page set and a LineChart with data
- WHEN `RenderLineChart(ctx, chart, bounds)` is called
- THEN lines connect data points
- AND markers are drawn at data points if specified
- AND no error is returned

#### Scenario: Render pie chart
- GIVEN a RenderingContext with Page set and a PieChart with data
- WHEN `RenderPieChart(ctx, chart, bounds)` is called
- THEN pie slices are drawn with correct angles proportional to values
- AND slice colors match series colors
- AND no error is returned

#### Scenario: Render chart title
- GIVEN a RenderingContext and a chart with title "Sales Data"
- WHEN `RenderChartTitle(ctx, "Sales Data", bounds)` is called
- THEN the title text is drawn centered above the chart area

### Requirement: Shape Rendering (Enabled)

The system SHALL render shapes via the enabled shape_renderer.go.

#### Scenario: Render shape with geometry
- GIVEN a RenderingContext with Page set and a Shape with preset geometry
- WHEN `RenderShape(ctx, shape, bounds)` is called
- THEN the shape geometry is rendered within bounds
- AND fill and stroke are applied per shape properties

### Requirement: Fill Rendering (Enabled)

The system SHALL render fills via the enabled fill_renderer.go.

#### Scenario: Render solid fill
- GIVEN a RenderingContext with Page set and a SolidFill with RGB color
- WHEN `RenderSolidFill(ctx, fill)` is called
- THEN `ctx.Page.SetFillColor()` is called with correct RGB values

#### Scenario: Render gradient fill
- GIVEN a RenderingContext with Page set and a GradientFill
- WHEN `RenderGradientFill(ctx, fill, bounds)` is called
- THEN gradient is approximated via multiple filled shapes
- AND colors transition smoothly across bounds

### Requirement: Stroke Rendering (Enabled)

The system SHALL render strokes via the enabled stroke_renderer.go.

#### Scenario: Render stroke
- GIVEN a RenderingContext with Page set and an Outline with width and color
- WHEN `RenderStroke(ctx, outline)` is called
- THEN `ctx.Page.SetStrokeColor()` and `SetLineWidth()` are called appropriately

### Requirement: Effects Rendering (Enabled)

The system SHALL render effects via the enabled effects_renderer.go.

#### Scenario: Render shadow
- GIVEN a RenderingContext with Page set and a Shadow effect
- WHEN `RenderShadow(ctx, shadow)` is called
- THEN a shadow is drawn offset from the shape
- AND shadow uses blur and transparency from effect definition

#### Scenario: Render glow
- GIVEN a RenderingContext with Page set and a Glow effect
- WHEN `RenderGlow(ctx, glow)` is called
- THEN a glow effect is rendered around the shape

### Requirement: Transform Rendering (Enabled)

The system SHALL render transforms via the enabled transform_renderer.go.

#### Scenario: Render transform
- GIVEN a RenderingContext with Page set and a Transform with rotation
- WHEN `RenderTransform(ctx, transform)` is called
- THEN `ctx.Page.Rotate()` is called with correct angle

### Requirement: Text in Shape (Enabled)

The system SHALL render text in shapes via the enabled text_in_shape.go.

#### Scenario: Render text in shape
- GIVEN a RenderingContext with Page set and a TextBody within shape bounds
- WHEN `RenderTextInShape(ctx, text, bounds)` is called
- THEN text is rendered within the shape bounds
- AND text wrapping respects bounds

### Requirement: Image Rendering (Enabled)

The system SHALL render images via the enabled image_renderer.go.

#### Scenario: Render image
- GIVEN a RenderingContext with Page set and a Blip (image reference)
- WHEN `RenderImage(ctx, blip, bounds)` is called
- THEN the image is drawn within bounds
- AND aspect ratio may be preserved per settings

## Implementation Notes

### File Structure (Enable existing + add new)
```
pdf/core/
├── document.go       // EXISTING: PDF document + concrete Page struct
├── page.go           // NEW: Page interface (~100 LOC)
├── pdf_page.go       // NEW: PDFPage implementation (~300 LOC)
├── render_context.go // ENHANCE: Add Page field to RenderingContext

pdf/drawing/
├── chart_renderer.go    // ENABLE: Rename from .wip (499 LOC)
├── shape_renderer.go    // ENABLE: Rename from .wip (485 LOC)
├── fill_renderer.go     // ENABLE: Rename from .wip (307 LOC)
├── stroke_renderer.go   // ENABLE: Rename from .wip (221 LOC)
├── effects_renderer.go  // ENABLE: Rename from .wip (272 LOC)
├── transform_renderer.go // ENABLE: Rename from .wip (233 LOC)
├── text_in_shape.go     // ENABLE: Rename from .wip (392 LOC)
├── image_renderer.go    // ENABLE: Rename from .wip (168 LOC)
└── *_test.go            // ENABLE: Rename from .wip (966 LOC)
```

### API Surface Summary (New)

```go
// NEW Interface
type Page interface {
    SaveGraphicsState()
    RestoreGraphicsState()
    SetFillColor(r, g, b float64)
    SetFillColorRGBA(r, g, b, a float64)
    SetStrokeColor(r, g, b float64)
    SetStrokeColorRGBA(r, g, b, a float64)
    SetLineWidth(width float64)
    SetLineCap(cap int)             // 0=butt, 1=round, 2=square
    SetLineJoin(join int)            // 0=miter, 1=round, 2=bevel
    SetLineDashPattern(pattern []float64, phase float64)
    DrawRectangle(x, y, w, h float64, fill, stroke bool)
    DrawRoundedRectangle(x, y, w, h, radius float64)
    DrawCircle(cx, cy, r float64, fill, stroke bool)
    DrawEllipse(cx, cy, rx, ry float64)
    DrawLine(x1, y1, x2, y2 float64)
    DrawPath(path *Path)
    Fill()
    Stroke()
    FillStroke()
    Clip()                           // Apply current path as clipping path
    SetFont(name string, size float64)
    DrawText(x, y float64, text string)
    DrawTextRotated(x, y, angle float64, text string)
    DrawImage(img *Image, x, y, w, h float64)
    Translate(dx, dy float64)
    Scale(sx, sy float64)
    Rotate(angle float64)
    SetClipRect(x, y, w, h float64)
    ClearClip()
    WriteContent(content string)
}

// Supporting types
type Path = *PathBuilder   // From pdf/drawing/path.go
type Font struct { ... }   // From pdf/font package
type Image struct { ... }  // From pdf/core package

// Color helpers (already exist in pdf/drawing/color.go)
func NewRGB(r, g, b float64) Color
func ParseColor(hex string) Color
func NewPathBuilder() *PathBuilder

// NEW Type
type PDFPage struct { ... }
func NewPDFPage(doc *Document, width, height float64) *PDFPage

// ENHANCED Method on RenderingContext
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

## Coordinate System

**PDF uses points**: 72 points = 1 inch
**DrawingML uses EMUs**: 914400 EMUs = 1 inch

Conversion (already in .wip code):
```go
func emuToPoints(emu int64) float64 {
    return float64(emu) / 914400.0 * 72.0
}
```

## Backward Compatibility

**Unchanged APIs** (fully backward compatible):
- All existing pdf/core types and methods unchanged
- RenderingContext struct gains optional Page field (nil by default)
- All pdf/word rendering continues to work

**New APIs** (additive only):
- Page interface (new)
- PDFPage type (new)
- RenderingContext.SetPage() method (new)

**Enabled APIs** (previously .wip):
- All Render* functions in pdf/drawing/

**No Breaking Changes**: Existing code continues to work unchanged.

## Testing Requirements

**Unit Tests (from tasks.md)**:
- Page interface: 15+ tests for PDFPage implementation
- Context integration: 4 tests
- Mock Page for renderer tests

**Integration Tests**:
- Chart rendering: 7 tests (bar, line, pie, title, legend, axes)
- Shape rendering: 7 tests (rect, gradient, circle, stroke, shadow, rotation)
- Presentation: 5 tests (pptx-charts example)

**Fidelity Tests**:
- Reference image comparison
- Regression prevention

## Success Metrics

1. **Compilation**: All .wip files compile after rename
2. **Tests pass**: All 966 LOC of tests pass
3. **Charts visible**: pptx-charts example produces visible charts
4. **No regressions**: Word PDF rendering still works
5. **Performance**: Chart rendering <100ms per chart

## Phase 2+ Enhancements (NOT IN THIS CHANGE)

Deferred to future proposals:
- Waterfall, Funnel, Stock chart types
- 3D chart perspective rendering
- Combination charts (multi-type)
- SVG backend
- Higher fidelity compared to Office output
