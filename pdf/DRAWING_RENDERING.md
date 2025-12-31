# DrawingML PDF Rendering Developer Guide

## Overview

This guide explains how to use and extend the DrawingML PDF rendering system in goffice. The system provides comprehensive rendering of Office Open XML (OOXML) DrawingML elements to PDF, including shapes, charts, images, fills, strokes, effects, and text.

## Table of Contents

1. [Architecture](#architecture)
2. [Core Components](#core-components)
3. [Usage](#usage)
4. [Unit Conversions](#unit-conversions)
5. [Extending the System](#extending-the-system)
6. [Testing](#testing)
7. [API Reference](#api-reference)
8. [Troubleshooting](#troubleshooting)

---

## Architecture

The DrawingML PDF rendering system follows a **modular, renderer-based architecture** where specialized renderers handle different aspects of shape rendering.

### Design Principles

1. **Separation of Concerns**: Each renderer handles one aspect (fill, stroke, effects, etc.)
2. **Abstraction**: The `PageDrawer` interface decouples rendering from PDF library specifics
3. **Context Management**: `RenderingContext` manages coordinate transformations and state
4. **Composability**: Renderers can be combined to create complex effects

### System Layers

```
Application Layer (Word/Excel/PowerPoint renderers)
    ↓ uses
DrawingML Rendering Layer (pdf/drawing/)
    ↓ uses
PDF Core Layer (pdf/core/)
    ↓ uses
PDF Library (pdfcpu)
```

---

## Core Components

### 1. PageDrawer Interface

**Location**: `pdf/core/page.go`

The `PageDrawer` interface defines the drawing API for PDF pages. It abstracts PDF content stream operations, enabling both actual PDF generation and testing with mock implementations.

**Key Methods**:
- `DrawRectangle(x, y, width, height, fill, stroke)` - Draw rectangles
- `DrawCircle(cx, cy, radius, fill, stroke)` - Draw circles
- `DrawEllipse(cx, cy, rx, ry, fill, stroke)` - Draw ellipses
- `SetFillColor(r, g, b)` - Set fill color (RGB 0-1)
- `SetStrokeColor(r, g, b)` - Set stroke color
- `SetLineWidth(width)` - Set line width
- `SetLineDashPattern(pattern, phase)` - Set dash pattern
- `SaveGraphicsState()` / `RestoreGraphicsState()` - State management
- `Transform(matrix)` - Apply transformation matrix
- `DrawText(x, y, text)` - Draw text
- `SetFont(name, size)` - Set font
- `WriteContent(content)` - Write raw PDF operators

**Coordinate System**:
- PDF uses **bottom-left origin** (X increases right, Y increases up)
- Units are in **points** (1/72 inch)
- Use `RenderingContext` for OOXML top-left origin conversion

### 2. PageImpl

**Location**: `pdf/core/page_impl.go`

Concrete implementation of `PageDrawer` that generates PDF content stream operators.

**Features**:
- Accumulates PDF drawing commands as strings
- Optimizes by tracking state (avoids redundant color/width changes)
- Uses PDF operators: `re` (rectangle), `c` (curve), `S` (stroke), `f` (fill), `B` (fill+stroke)
- Implements circle/ellipse using Bézier curve approximations

**Example Usage**:
```go
page := core.NewPageImpl(pageWidth, pageHeight)
page.SetFillColor(0.8, 0.2, 0.2)  // Red
page.SetStrokeColor(0, 0, 0)      // Black outline
page.DrawRectangle(100, 100, 200, 150, true, true)
content := page.GetContent()  // Get PDF operators
```

### 3. RenderingContext

**Location**: `pdf/core/render_context.go`

Manages rendering state for converting OOXML content to PDF, including:

**Transformation Management**:
- **Origin Stack**: Nested coordinate contexts (e.g., table cells, shape groups)
- **Transform Stack**: Cumulative transformation matrices
- **Bounds Stack**: Content clipping regions

**Key Features**:
- Converts OOXML top-left origin to PDF bottom-left origin
- Handles nested contexts (headers/footers, table cells, shape groups)
- Tracks page dimensions and margins
- Provides coordinate transformation methods

**Common Methods**:
- `TransformPoint(x, y)` - Convert OOXML coordinates to PDF
- `TransformPointEMU(emuX, emuY)` - Convert EMU to PDF coordinates
- `TransformRect(x, y, w, h)` - Transform rectangle
- `PushOrigin(x, y)` / `PopOrigin()` - Manage coordinate contexts
- `PushTransform(matrix)` / `PopTransform()` - Manage transformations
- `SaveState()` / `RestoreState()` - Save/restore all stacks

**Example Usage**:
```go
ctx := core.NewRenderingContext(pageWidth, pageHeight)
ctx.SetPage(page)  // Set the PageDrawer

// Transform OOXML coordinates (top-left origin, EMU) to PDF
pdfX, pdfY := ctx.TransformPointEMU(emuX, emuY)

// Push a nested context (e.g., for a table cell)
ctx.PushOrigin(cellX, cellY)
// ... render cell contents ...
ctx.PopOrigin()
```

### 4. Renderer Components

All renderers are located in `pdf/drawing/` and use the `RenderingContext` and `PageDrawer` to perform their work.

#### ChartRenderer

**Location**: `pdf/drawing/chart_renderer.go`

Renders DrawingML charts (bar, line, pie).

**Capabilities**:
- **Bar/Column Charts**: Vertical and horizontal orientations
- **Line Charts**: With data point markers
- **Pie Charts**: With automatic slicing
- **Chart Elements**: Axes, legends, titles
- **Auto-scaling**: Scales data to fit chart area

**Example**:
```go
chartRenderer := drawing.NewChartRenderer(ctx)

data := drawing.ChartData{
    Categories: []string{"Q1", "Q2", "Q3", "Q4"},
    Series: []drawing.ChartSeries{
        {
            Name:   "Sales",
            Values: []float64{100, 150, 120, 180},
            Color:  drawing.NewRGB(0.2, 0.4, 0.8),
        },
    },
}

err := chartRenderer.RenderBarChart(x, y, width, height, data, false)
```

#### ShapeRenderer

**Location**: `pdf/drawing/shape_renderer.go`

Renders complete DrawingML shapes with all properties.

**Capabilities**:
- **Preset Geometries**: Rectangle, ellipse, triangle, diamond, pentagon, hexagon, octagon, stars, arrows
- **Custom Geometries**: Path-based shapes (basic support)
- **Transformations**: Rotation, flip (basic support)
- **Composition**: Combines fill, stroke, and effects

**Supported Shapes**:
- `rectangle`, `roundRectangle`, `ellipse`
- `triangle`, `diamond`
- `pentagon`, `hexagon`, `octagon`
- `star4`, `star5`, `star6`, `star8`, `star10`, `star12`
- `rightArrow`, `leftArrow`, `upArrow`, `downArrow`
- `line`

**Example**:
```go
shapeRenderer := drawing.NewShapeRenderer(ctx)

// Get shape properties from DrawingML element
sp := shapeElement.ShapeProperties()

// Create renderers for fill and stroke
fill := drawing.CreateFillRenderer(sp)
stroke := drawing.CreateStrokeRenderer(sp)

err := shapeRenderer.RenderShapeWithFill(sp, fill, stroke)
```

#### FillRenderer

**Location**: `pdf/drawing/fill_renderer.go`

Interface for rendering shape fills with multiple implementations.

**Fill Types**:
- **SolidFill**: Single color fill
- **GradientFill**: Linear/radial gradients (approximated with solid color)
- **PatternFill**: Tiled patterns (approximated with solid color)
- **PictureFill**: Image fills (placeholder implementation)
- **NoFill**: Transparent (no fill)

**Example**:
```go
// Automatically create appropriate renderer
fill := drawing.CreateFillRenderer(shapeProperties)

// Or create specific fill type
solidFill := drawing.NewSolidFillRenderer(solidFillElement)

// Apply fill to a path
path := drawing.NewPathBuilder()
path.Rectangle(x, y, width, height)
err := fill.Apply(ctx, path)
```

#### StrokeRenderer

**Location**: `pdf/drawing/stroke_renderer.go`

Renders shape outlines/strokes.

**Capabilities**:
- **Line Styles**: Solid, dashed, dotted
- **Line Caps**: Butt, round, square
- **Line Joins**: Miter, round, bevel
- **Line Width**: Adjustable width
- **Colors**: RGB, scheme colors

**Example**:
```go
stroke := drawing.CreateStrokeRenderer(shapeProperties)

path := drawing.NewPathBuilder()
path.Circle(cx, cy, radius)
err := stroke.Apply(ctx, path)
```

#### ImageRenderer

**Location**: `pdf/drawing/image_renderer.go`

Renders embedded images (PNG, JPEG).

**Capabilities**:
- Image placement and scaling
- Aspect ratio handling
- Placeholder rendering (full implementation requires XObject support)

#### EffectsRenderer

**Location**: `pdf/drawing/effects_renderer.go`

Renders visual effects like shadows, glows, reflections.

**Current Status**: Basic placeholder implementations (full effect support is complex in PDF)

#### TextInShapeRenderer

**Location**: `pdf/drawing/text_in_shape.go`

Renders text within shape bounds.

**Capabilities**:
- **Text Layout**: Paragraph and run-level formatting
- **Alignment**: Left, center, right, justified
- **Vertical Alignment**: Top, center, bottom
- **Text Insets**: Margins within shape
- **Font Properties**: Font family, size, bold, italic, color
- **Autofit**: Text scaling to fit shape (basic)
- **Vertical Text**: Rotated text (basic support)

**Example**:
```go
textRenderer := drawing.NewTextInShapeRenderer(ctx)

// Render text from DrawingML TextBody
err := textRenderer.RenderTextInShape(
    textBody,
    shapeX, shapeY,
    shapeWidth, shapeHeight,
)
```

#### TransformRenderer

**Location**: `pdf/drawing/transform_renderer.go`

Handles geometric transformations.

**Capabilities**:
- Translation
- Rotation
- Scaling
- Flip (horizontal/vertical)
- Matrix operations

### 5. PathBuilder

**Location**: `pdf/drawing/path.go`

Builder pattern for constructing PDF paths.

**Methods**:
- `MoveTo(x, y)` - Start new subpath
- `LineTo(x, y)` - Add line segment
- `CurveTo(cp1x, cp1y, cp2x, cp2y, x, y)` - Cubic Bézier curve
- `QuadraticCurveTo(cpx, cpy, x, y)` - Quadratic Bézier curve (converted to cubic)
- `Rectangle(x, y, w, h)` - Rectangle using `re` operator
- `RoundedRect(x, y, w, h, radius)` - Rounded rectangle
- `Circle(cx, cy, radius)` - Circle (Bézier approximation)
- `Ellipse(cx, cy, rx, ry)` - Ellipse (Bézier approximation)
- `RegularPolygon(cx, cy, radius, sides, startAngle)` - N-sided polygon
- `Star(cx, cy, outer, inner, points, startAngle)` - Star shape
- `ClosePath()` - Close current subpath
- `Fill()`, `Stroke()`, `FillAndStroke()` - Generate PDF painting operators

**Example**:
```go
path := drawing.NewPathBuilder()
path.MoveTo(100, 100)
path.LineTo(200, 100)
path.LineTo(200, 200)
path.ClosePath()

// Generate PDF operators
ctx.Page.WriteContent(path.FillAndStroke())
```

---

## Usage

### Basic Rendering Workflow

1. **Create RenderingContext**:
   ```go
   ctx := core.NewRenderingContext(pageWidth, pageHeight)
   ```

2. **Create and attach PageDrawer**:
   ```go
   page := core.NewPageImpl(pageWidth, pageHeight)
   ctx.SetPage(page)
   ```

3. **Create Renderers**:
   ```go
   shapeRenderer := drawing.NewShapeRenderer(ctx)
   chartRenderer := drawing.NewChartRenderer(ctx)
   textRenderer := drawing.NewTextInShapeRenderer(ctx)
   ```

4. **Render DrawingML Elements**:
   ```go
   // Render a shape
   fill := drawing.CreateFillRenderer(shapeProps)
   stroke := drawing.CreateStrokeRenderer(shapeProps)
   err := shapeRenderer.RenderShapeWithFill(shapeProps, fill, stroke)

   // Render text in shape
   err = textRenderer.RenderTextInShape(textBody, x, y, width, height)
   ```

5. **Extract PDF Content**:
   ```go
   pdfContent := page.GetContent()
   // Write to PDF content stream
   ```

### Rendering a Complete Document

```go
import (
    "github.com/connerohnesorge/goffice-pdf/core"
    "github.com/connerohnesorge/goffice-pdf/drawing"
    "github.com/connerohnesorge/goffice/wordprocessing"
)

func renderDocumentWithDrawing(doc *wordprocessing.Document) error {
    // Set up page
    pageOpts := core.DefaultPageOptions()
    ctx := core.NewRenderingContextFromPageOptions(pageOpts)
    page := core.NewPageImpl(pageOpts.Size.Width, pageOpts.Size.Height)
    ctx.SetPage(page)

    // Get DrawingML elements from document
    drawings := extractDrawings(doc)

    // Create renderers
    shapeRenderer := drawing.NewShapeRenderer(ctx)

    // Render each drawing
    for _, drw := range drawings {
        if inline := drw.Inline(); inline != nil {
            // Handle inline drawing
            extent := inline.Extent()
            width := core.EMUToPoints(float64(extent.Cx()))
            height := core.EMUToPoints(float64(extent.Cy()))

            graphic := inline.Graphic()
            if graphicData := graphic.GraphicData(); graphicData != nil {
                // Render shape, chart, image, etc.
                renderGraphicData(ctx, shapeRenderer, graphicData, width, height)
            }
        }
    }

    // Get PDF content
    content := page.GetContent()
    // ... write to PDF document ...

    return nil
}
```

### Handling Coordinate Transformations

OOXML uses **English Metric Units (EMU)** with a **top-left origin**, while PDF uses **points** with a **bottom-left origin**.

```go
// OOXML coordinates (EMU, top-left origin)
emuX := int64(1000000)  // ~1.09 inches from left
emuY := int64(500000)   // ~0.55 inches from top

// Convert to PDF coordinates
pdfX, pdfY := ctx.TransformPointEMU(emuX, emuY)

// Now pdfX, pdfY are in points with bottom-left origin
```

### Nested Contexts

For rendering content within specific areas (table cells, shape groups):

```go
// Enter a table cell context
cellCtx := ctx.EnterTableCell(x, y, width, height, row, col)
// ... render cell contents ...
cellCtx.ExitContentArea()

// Enter a shape group context
groupCtx := ctx.EnterShapeGroup(groupTransform)
// ... render shapes within group ...
groupCtx.ExitShapeGroup()
```

---

## Unit Conversions

### Key Constants

From `pdf/core/units.go`:

```go
const (
    PointsPerInch = 72.0         // 72 points = 1 inch
    PointsPerMM   = 2.834645     // ~2.83 points = 1 mm
    PointsPerCM   = 28.346457    // ~28.35 points = 1 cm
    PointsPerEMU  = 0.000078740  // 1 point = 12700 EMU
    EMUPerInch    = 914400.0     // 914400 EMU = 1 inch
    PointsPerTwip = 0.05         // 20 twips = 1 point
)
```

### Conversion Functions

**OOXML (EMU) to PDF (Points)**:
```go
points := core.EMUToPoints(emuValue)
// Example: 914400 EMU = 72 points = 1 inch
```

**PDF (Points) to OOXML (EMU)**:
```go
emu := core.PointsToEMU(pointsValue)
```

**Other Units**:
```go
points := core.InchesToPoints(1.0)     // 72 points
points := core.MMToPoints(25.4)        // 72 points (1 inch)
points := core.CMToPoints(2.54)        // 72 points (1 inch)
points := core.TwipsToPoints(1440)     // 72 points (1 inch)
```

### Dimension Type

For unit-aware calculations:

```go
dim := core.Inches(1.5)              // 1.5 inches
points := dim.ToPoints()             // 108 points

dim2 := core.MM(10)                  // 10 mm
combined := dim.Add(dim2).ToPoints() // Both converted to points
```

---

## Extending the System

### Adding a New Shape Type

To add support for a new preset shape type:

1. **Add to ShapeType enum** (in DrawingML package)
2. **Implement rendering in `shape_renderer.go`**:

```go
func renderPresetGeometry(
    geom *drawingml.PresetGeometry,
    path *PathBuilder,
    x, y, width, height float64,
) error {
    shapeType := geom.Preset()

    switch shapeType {
    // ... existing cases ...

    case drawingml.ShapeTypeYourNewShape:
        renderYourNewShape(path, x, y, width, height)

    default:
        // Fallback to rectangle
        path.Rectangle(x, y, width, height)
    }

    return nil
}

func renderYourNewShape(
    path *PathBuilder,
    x, y, width, height float64,
) {
    // Use PathBuilder methods to construct the shape
    path.MoveTo(x, y)
    path.LineTo(x + width/2, y + height)
    path.LineTo(x + width, y)
    path.ClosePath()
}
```

3. **Add tests** in `shape_renderer_test.go`:

```go
func TestShapeRenderer_YourNewShape(t *testing.T) {
    mockPage := &MockPage{}
    ctx := core.NewRenderingContext(612, 792)
    ctx.SetPage(mockPage)

    renderer := NewShapeRenderer(ctx)

    // Create test shape properties
    sp := createTestShapeWithGeometry(drawingml.ShapeTypeYourNewShape)

    // Render
    err := renderer.RenderShapeWithFill(sp, nil, nil)
    if err != nil {
        t.Fatalf("Failed to render: %v", err)
    }

    // Verify PDF operators
    if !mockPage.HasContent("m", "l", "h") {
        t.Error("Expected path operators not found")
    }
}
```

### Adding a New Fill Type

To add a new fill renderer:

1. **Create a new renderer struct**:

```go
type YourFillRenderer struct {
    fill *drawingml.YourFillType
}

func NewYourFillRenderer(fill *drawingml.YourFillType) *YourFillRenderer {
    return &YourFillRenderer{fill: fill}
}

func (r *YourFillRenderer) Apply(
    ctx *core.RenderingContext,
    path *PathBuilder,
) error {
    if r.fill == nil {
        return nil
    }

    // Extract fill properties
    // ... your logic ...

    // Apply fill
    ctx.Page.SetFillColor(r, g, b)
    ctx.Page.WriteContent(path.Fill())

    return nil
}
```

2. **Update `CreateFillRenderer` factory**:

```go
func CreateFillRenderer(sp *drawingml.ShapeProperties) FillRenderer {
    // ... existing checks ...

    if yourFillElem := sp.GetElement("yourFill", drawingml.NamespaceMain); yourFillElem != nil {
        if yourFill, ok := yourFillElem.(*drawingml.YourFillType); ok {
            return NewYourFillRenderer(yourFill)
        }
    }

    return NewNoFillRenderer()
}
```

### Adding Custom Effects

Effects in PDF require advanced techniques (transparency groups, blending modes). Basic approach:

```go
type YourEffectRenderer struct {
    effect *drawingml.YourEffect
}

func (r *YourEffectRenderer) Apply(ctx *core.RenderingContext) error {
    // Save graphics state
    ctx.Page.SaveGraphicsState()
    defer ctx.Page.RestoreGraphicsState()

    // Apply effect parameters
    // For shadows: render offset shape with lighter color
    // For glows: render multiple copies with increasing size
    // For reflections: render flipped copy with transparency

    return nil
}
```

### Mocking for Testing

Use `MockPage` for unit tests without PDF generation:

```go
type MockPage struct {
    operations []string
    fillColor  [3]float64
    // ... other state ...
}

func (m *MockPage) DrawRectangle(x, y, w, h float64, fill, stroke bool) {
    m.operations = append(m.operations, "DrawRectangle")
}

func (m *MockPage) SetFillColor(r, g, b float64) {
    m.fillColor = [3]float64{r, g, b}
}

// Test usage
func TestMyRenderer(t *testing.T) {
    mockPage := &MockPage{}
    ctx := core.NewRenderingContext(612, 792)
    ctx.SetPage(mockPage)

    // ... perform rendering ...

    // Verify calls
    if len(mockPage.operations) != expectedCount {
        t.Errorf("Expected %d operations, got %d", expectedCount, len(mockPage.operations))
    }
}
```

---

## Testing

### Unit Testing

Each renderer has dedicated test files:

- `chart_renderer_test.go` - Chart rendering tests
- `fill_renderer_test.go` - Fill rendering tests
- `shape_renderer_test.go` - Shape rendering tests
- `path_test.go` - Path construction tests
- `integration_test.go` - End-to-end rendering tests

**Run all tests**:
```bash
cd pdf/drawing
go test ./...
```

**Run specific test**:
```bash
go test -run TestShapeRenderer_Rectangle
```

**With coverage**:
```bash
go test -cover ./...
```

### Integration Testing

Integration tests verify complete rendering pipelines:

```go
func TestIntegration_CompleteShapeRendering(t *testing.T) {
    // Create full rendering stack
    page := core.NewPageImpl(612, 792)
    ctx := core.NewRenderingContext(612, 792)
    ctx.SetPage(page)

    // Create renderers
    shapeRenderer := NewShapeRenderer(ctx)

    // Create test shape
    sp := createTestShape()
    fill := CreateFillRenderer(sp)
    stroke := CreateStrokeRenderer(sp)

    // Render
    err := shapeRenderer.RenderShapeWithFill(sp, fill, stroke)
    if err != nil {
        t.Fatalf("Rendering failed: %v", err)
    }

    // Verify PDF content
    content := page.GetContent()
    if !strings.Contains(content, "f") {  // Fill operator
        t.Error("Expected fill operator")
    }
    if !strings.Contains(content, "S") {  // Stroke operator
        t.Error("Expected stroke operator")
    }
}
```

### Visual Testing

For visual verification:

1. Generate test PDFs with known shapes
2. Open in PDF viewer (Adobe Reader, Preview, LibreOffice)
3. Compare with expected output
4. Use `pdfinfo` or `pdftotext` for automated checks

```bash
# Generate test PDF
go run examples/render_shapes.go

# Validate PDF structure
pdfinfo test_output.pdf

# Extract text for verification
pdftotext test_output.pdf
```

---

## API Reference

### Key Types

#### PageDrawer Interface
```go
type PageDrawer interface {
    DrawRectangle(x, y, width, height float64, fill, stroke bool)
    DrawCircle(cx, cy, radius float64, fill, stroke bool)
    DrawEllipse(cx, cy, rx, ry float64, fill, stroke bool)
    SetFillColor(r, g, b float64)
    SetStrokeColor(r, g, b float64)
    SetLineWidth(width float64)
    SetLineDashPattern(pattern []float64, phase float64)
    SetLineCap(lineCap LineCap)
    SetLineJoin(lineJoin LineJoin)
    SaveGraphicsState()
    RestoreGraphicsState()
    Transform(matrix PDFMatrix)
    AddImage(img PageImage, x, y, width, height float64)
    DrawText(x, y float64, text string)
    SetFont(name string, size float64)
    WriteContent(content string)
}
```

#### RenderingContext

```go
type RenderingContext struct {
    Page PageDrawer  // Must be set before rendering
    // ... internal fields ...
}

// Creation
func NewRenderingContext(pageWidth, pageHeight float64) *RenderingContext
func NewRenderingContextFromPageSize(size PageSize) *RenderingContext
func NewRenderingContextFromPageOptions(opts *PageOptions) *RenderingContext

// Coordinate transformation
func (rc *RenderingContext) TransformPoint(x, y float64) (pdfX, pdfY float64)
func (rc *RenderingContext) TransformPointEMU(emuX, emuY int64) (pdfX, pdfY float64)
func (rc *RenderingContext) TransformRect(x, y, width, height float64) Rectangle

// State management
func (rc *RenderingContext) PushOrigin(x, y float64)
func (rc *RenderingContext) PopOrigin() bool
func (rc *RenderingContext) PushTransform(matrix PDFMatrix)
func (rc *RenderingContext) PopTransform() bool
func (rc *RenderingContext) SaveState() RenderState
func (rc *RenderingContext) RestoreState(state RenderState)
```

#### PathBuilder

```go
type PathBuilder struct { /* ... */ }

func NewPathBuilder() *PathBuilder

// Path construction
func (pb *PathBuilder) MoveTo(x, y float64) *PathBuilder
func (pb *PathBuilder) LineTo(x, y float64) *PathBuilder
func (pb *PathBuilder) CurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) *PathBuilder
func (pb *PathBuilder) ClosePath() *PathBuilder

// Shapes
func (pb *PathBuilder) Rectangle(x, y, w, h float64) *PathBuilder
func (pb *PathBuilder) RoundedRect(x, y, w, h, radius float64) *PathBuilder
func (pb *PathBuilder) Circle(cx, cy, radius float64) *PathBuilder
func (pb *PathBuilder) Ellipse(cx, cy, rx, ry float64) *PathBuilder
func (pb *PathBuilder) RegularPolygon(cx, cy, radius float64, sides int, startAngle float64) *PathBuilder
func (pb *PathBuilder) Star(cx, cy, outerRadius, innerRadius float64, points int, startAngle float64) *PathBuilder

// Output
func (pb *PathBuilder) String() string
func (pb *PathBuilder) Fill() string
func (pb *PathBuilder) Stroke() string
func (pb *PathBuilder) FillAndStroke() string
```

#### Renderer Constructors

```go
// Shape rendering
func NewShapeRenderer(ctx *RenderingContext) *ShapeRenderer

// Chart rendering
func NewChartRenderer(ctx *RenderingContext) *ChartRenderer

// Fill rendering
func CreateFillRenderer(sp *drawingml.ShapeProperties) FillRenderer
func NewSolidFillRenderer(fill *drawingml.SolidFill) *SolidFillRenderer
func NewGradientFillRenderer(fill *drawingml.GradientFill) *GradientFillRenderer
func NewNoFillRenderer() *NoFillRenderer

// Stroke rendering
func CreateStrokeRenderer(sp *drawingml.ShapeProperties) StrokeRenderer

// Text rendering
func NewTextInShapeRenderer(ctx *RenderingContext) *TextInShapeRenderer
```

---

## Troubleshooting

### Common Issues

#### 1. Coordinate Mismatch

**Problem**: Shapes appear upside-down or in wrong position.

**Cause**: Mixing OOXML (top-left origin) and PDF (bottom-left origin) coordinates.

**Solution**: Always use `RenderingContext.TransformPoint()` or `TransformPointEMU()`:

```go
// Wrong
ctx.Page.DrawRectangle(ooxmlX, ooxmlY, width, height, true, false)

// Correct
pdfX, pdfY := ctx.TransformPoint(ooxmlX, ooxmlY)
ctx.Page.DrawRectangle(pdfX, pdfY, width, height, true, false)
```

#### 2. Unit Conversion Errors

**Problem**: Shapes are way too large or too small.

**Cause**: Forgetting to convert EMU to points.

**Solution**: Always convert EMU values:

```go
// Wrong
width := float64(extent.Cx())  // Still in EMU!

// Correct
width := core.EMUToPoints(float64(extent.Cx()))
```

#### 3. Missing Graphics State Save/Restore

**Problem**: Settings from one shape affect the next shape.

**Cause**: Not using graphics state stack properly.

**Solution**: Always save/restore around temporary changes:

```go
ctx.Page.SaveGraphicsState()
ctx.Page.SetFillColor(1, 0, 0)  // Temporary red
ctx.Page.DrawRectangle(x, y, w, h, true, false)
ctx.Page.RestoreGraphicsState()  // Back to previous color
```

#### 4. Text Not Appearing

**Problem**: `DrawText()` doesn't show anything.

**Cause**: Font not set, or text position outside page bounds.

**Solution**: Always set font before drawing text:

```go
ctx.Page.SetFont("Helvetica", 12)
ctx.Page.SetFillColor(0, 0, 0)  // Black text
ctx.Page.DrawText(x, y, "Hello World")
```

#### 5. Paths Not Rendering

**Problem**: Path constructed but nothing appears.

**Cause**: Forgot to call `Fill()`, `Stroke()`, or `FillAndStroke()`.

**Solution**: Always finalize paths with a painting operator:

```go
path := drawing.NewPathBuilder()
path.Rectangle(x, y, w, h)
ctx.Page.WriteContent(path.Fill())  // Don't forget this!
```

### Debugging Tips

**1. Enable Verbose Logging**:
```go
import "log"

// Log all drawing operations
ctx.Page.WriteContent = func(content string) {
    log.Printf("PDF: %s", content)
    realPage.WriteContent(content)
}
```

**2. Inspect Generated PDF Content**:
```go
content := page.GetContent()
fmt.Println("Generated PDF operators:")
fmt.Println(content)
```

**3. Use MockPage for Isolated Testing**:
```go
mockPage := &MockPage{}
ctx.SetPage(mockPage)
// ... render ...
fmt.Println("Operations called:", mockPage.operations)
```

**4. Verify Coordinate Transformations**:
```go
x, y := ctx.TransformPoint(100, 100)
fmt.Printf("OOXML (100, 100) -> PDF (%.2f, %.2f)\n", x, y)
```

**5. Check PDF with External Tools**:
```bash
# Validate PDF syntax
pdfinfo output.pdf

# Extract and examine content streams
pdftk output.pdf output uncompressed.pdf uncompress
# Then view uncompressed.pdf in text editor
```

### Performance Optimization

**1. Minimize State Changes**:
`PageImpl` already optimizes redundant state changes. Avoid unnecessary `SetFillColor()` calls.

**2. Batch Drawing Operations**:
Group similar shapes together to reduce context switches.

**3. Reuse PathBuilder**:
```go
path := drawing.NewPathBuilder()
for _, shape := range shapes {
    path.Reset()  // Reuse builder
    // ... construct shape ...
}
```

**4. Avoid Deep Nesting**:
Minimize `PushOrigin()` / `PopOrigin()` depth for cleaner code and better performance.

---

## Further Reading

- **PDF Reference 1.7**: Adobe PDF specification for low-level details
- **ECMA-376 Part 1**: DrawingML specification
- **goffice CLAUDE.md**: Project-wide development guide
- **pdf/FIDELITY.md**: Rendering fidelity guidelines
- **pdf/LIMITATIONS.md**: Known limitations and workarounds

---

## Example: Complete Shape Rendering

```go
package main

import (
    "fmt"
    "github.com/connerohnesorge/goffice-pdf/core"
    "github.com/connerohnesorge/goffice-pdf/drawing"
    "github.com/connerohnesorge/goffice/drawingml"
)

func main() {
    // 1. Set up page
    pageOpts := core.NewPageOptions(core.PageSizeA4)
    ctx := core.NewRenderingContextFromPageOptions(pageOpts)
    page := core.NewPageImpl(pageOpts.Size.Width, pageOpts.Size.Height)
    ctx.SetPage(page)

    // 2. Create a blue rectangle with black border
    shapeRenderer := drawing.NewShapeRenderer(ctx)

    // Create shape properties (normally from DrawingML)
    sp := drawingml.NewShapeProperties()
    xfrm := sp.EnsureTransform2D()
    xfrm.SetOffset(core.PointsToEMU(100), core.PointsToEMU(100))
    xfrm.SetExtent(core.PointsToEMU(200), core.PointsToEMU(150))

    // Set preset geometry
    presetGeom := drawingml.NewPresetGeometry()
    presetGeom.SetPreset(drawingml.ShapeTypeRectangle)
    sp.AppendChild(presetGeom)

    // Add solid blue fill
    solidFill := drawingml.NewSolidFill()
    rgbColor := drawingml.NewRgbColor()
    rgbColor.SetVal("4472C4")  // Blue
    solidFill.AppendChild(rgbColor)
    sp.AppendChild(solidFill)

    // Add black outline
    ln := drawingml.NewOutline()
    ln.SetWidth(drawingml.EMU(12700))  // 1 point
    lnFill := drawingml.NewSolidFill()
    lnRgb := drawingml.NewRgbColor()
    lnRgb.SetVal("000000")  // Black
    lnFill.AppendChild(lnRgb)
    ln.AppendChild(lnFill)
    sp.AppendChild(ln)

    // 3. Render the shape
    fill := drawing.CreateFillRenderer(sp)
    stroke := drawing.CreateStrokeRenderer(sp)

    err := shapeRenderer.RenderShapeWithFill(sp, fill, stroke)
    if err != nil {
        fmt.Printf("Error rendering shape: %v\n", err)
        return
    }

    // 4. Get PDF content
    pdfContent := page.GetContent()
    fmt.Println("Generated PDF operators:")
    fmt.Println(pdfContent)

    // Output would include:
    // 0.267 0.447 0.769 rg    (blue fill color)
    // 0 0 0 RG                (black stroke color)
    // 1 w                     (1 point line width)
    // 100 642 200 150 re      (rectangle)
    // B                       (fill and stroke)
}
```

---

**Version**: 1.0
**Last Updated**: 2025-12-30
**Author**: goffice-pdf DrawingML rendering team
