# Design: DrawingML PDF Rendering Architecture

## Problem Statement

The PDF rendering system can produce Word documents with text, tables, and basic formatting, but cannot render any DrawingML elements (charts, shapes, images, effects) because there is no page-level drawing API. The disabled `.wip` renderers in `pdf/drawing/` all expect a `Page` abstraction that doesn't exist.

## Current Architecture

```
Document Renderers (Word/Excel/PowerPoint)
    ↓
RenderingContext (coordinates, fonts, state)
    ↓
??? No Page Abstraction ???
    ↓
pdfcpu (PDF document structure)
```

**Gap**: No way to draw shapes, paths, images onto PDF pages.

## Proposed Architecture

```
Document Renderers (Word/Excel/PowerPoint)
    ↓  
DrawingML Renderers (chart, shape, fill, stroke, image)
    ↓
Page Interface (drawing primitives + state management)
    ↓
PageImpl (pdfcpu adapter)
    ↓
pdfcpu Content Streams
    ↓
PDF Output
```

## Design Alternatives Considered

### Alternative 1: Direct pdfcpu Usage in Renderers
**Approach**: Each renderer imports and uses pdfcpu directly

**Pros**:
- No abstraction layer
- Direct access to pdfcpu features

**Cons**:
- Tight coupling to pdfcpu
- Difficult to test (no mocking)
- Duplicate coordinate conversion logic
- Hard to swap PDF libraries later

**Decision**: REJECTED - Coupling and testability concerns outweigh simplicity

### Alternative 2: Cairo/Pango Integration
**Approach**: Use Cairo graphics library for rendering

**Pros**:
- Industry-standard graphics library
- Rich feature set

**Cons**:
- CGO dependency (violates goffice zero-CGO constraint)
- Large external dependency
- Platform-specific builds

**Decision**: REJECTED - Violates core goffice design principle (no CGO)

### Alternative 3: Pure Go PDF Generation
**Approach**: Implement PDF content streams from scratch

**Pros**:
- Full control
- No external dependencies

**Cons**:
- Months of development effort
- Bug-prone (PDF is complex)
- pdfcpu already exists and works

**Decision**: REJECTED - Unnecessary duplication, high risk

### Alternative 4: Page Interface with pdfcpu Backend (CHOSEN)
**Approach**: Define minimal Page interface, implement with pdfcpu adapter

**Pros**:
- Testable (mock Page for unit tests)
- Decouples renderers from PDF implementation
- Reuses existing pdfcpu integration
- Minimal API surface (only what's needed)
- Allows future backend swaps if needed

**Cons**:
- One extra indirection layer

**Decision**: ACCEPTED - Best balance of testability, maintainability, and simplicity

## Page Interface Design

### Core Principles
1. **Minimal**: Only operations actually needed by DrawingML renderers
2. **Stateful**: Maintains graphics state stack (like PDF/Canvas)
3. **Coordinate-agnostic**: Accepts coordinates in caller's system
4. **Type-safe**: Strong typing for colors, fonts, images

### Interface Methods

```go
type Page interface {
    // === Drawing Primitives ===
    // DrawRectangle draws a rectangle with optional fill and stroke
    // Parameters: x, y (top-left corner), width, height, fill (bool), stroke (bool)
    DrawRectangle(x, y, width, height float64, fill, stroke bool)

    // DrawCircle draws a circle with optional fill and stroke
    // Parameters: cx, cy (center), radius, fill (bool), stroke (bool)
    DrawCircle(cx, cy, radius float64, fill, stroke bool)

    // DrawEllipse draws an ellipse with optional fill and stroke
    // Parameters: cx, cy (center), rx (x-radius), ry (y-radius), fill (bool), stroke (bool)
    DrawEllipse(cx, cy, rx, ry float64, fill, stroke bool)

    // DrawPath renders an arbitrary path (builder generates PDF operators)
    DrawPath(path *PathBuilder)

    // === State Management ===
    // SetFillColor sets the fill color for subsequent shapes
    // Parameters: r, g, b in range [0.0, 1.0]
    SetFillColor(r, g, b float64)

    // SetStrokeColor sets the stroke color
    // Parameters: r, g, b in range [0.0, 1.0]
    SetStrokeColor(r, g, b float64)

    // SetLineWidth sets stroke width in points
    SetLineWidth(width float64)

    // SetLineDashPattern sets dash pattern (empty = solid)
    // Parameters: pattern (on/off lengths), phase (offset)
    SetLineDashPattern(pattern []float64, phase float64)

    // SetLineCap sets line cap style (butt=0, round=1, square=2)
    SetLineCap(cap LineCap)

    // SetLineJoin sets line join style (miter=0, round=1, bevel=2)
    SetLineJoin(join LineJoin)

    // === Graphics State Stack ===
    // SaveGraphicsState saves current graphics state (PDF 'q' operator)
    SaveGraphicsState()

    // RestoreGraphicsState restores previous graphics state (PDF 'Q' operator)
    RestoreGraphicsState()

    // Transform applies transformation matrix to CTM
    Transform(matrix Matrix)

    // === Content Embedding ===
    // AddImage embeds and draws an image
    AddImage(img Image, x, y, width, height float64)

    // DrawText renders text at position using current font
    // Call SetFont() before DrawText()
    DrawText(text string, x, y float64)

    // SetFont sets the font for subsequent DrawText calls
    // Parameters: name (e.g., "Helvetica", "Times-Roman"), size (in points)
    SetFont(name string, size float64)

    // === Low-Level Access ===
    // WriteContent writes raw PDF content stream operators
    // Used for complex paths that PathBuilder generates
    // Example: WriteContent(path.Stroke()) or WriteContent(path.Fill())
    WriteContent(content string)
}
```

### Design Rationale for Method Signatures

**DrawRectangle/DrawCircle/DrawEllipse with fill/stroke bools**:
- Actual usage in .wip files: `page.DrawRectangle(x, y, w, h, true, true)` for filled+stroked
- Alternative considered: separate FillRectangle/StrokeRectangle methods
- Chosen approach reduces method count, matches PDF semantics (f/S/B operators)

**SetFillColor/SetStrokeColor with RGB floats**:
- Actual usage: `page.SetFillColor(0.5, 0.5, 0.5)` not `SetFillColor(Color{...})`
- Avoids allocation overhead for color structs
- Matches PDF operator format: `0.5 0.5 0.5 rg`

**SaveGraphicsState/RestoreGraphicsState naming**:
- Matches PDF terminology (q/Q operators)
- More explicit than PushState/PopState
- Consistent with graphics programming conventions

**WriteContent() as escape hatch**:
- Used in .wip files: `page.WriteContent(path.Stroke())`
- PathBuilder generates PDF operators as strings
- High-level methods handle common cases, WriteContent handles complex paths
- This is intentional "leaky abstraction" for performance/flexibility

### Not Included (Deferred)
- Clipping paths (W/W* operators - future enhancement)
- Blend modes beyond normal (future)
- Gradients as first-class API (renderers compose from primitives or use WriteContent)
- Patterns as first-class API (renderers use repeated drawing)

**Rationale**: Start with proven needs from .wip files, expand based on real usage. Most DrawingML features can be built by composing primitives or using WriteContent().

## pdfcpu Integration Details

### Content Stream Generation
pdfcpu provides `PageContentWriter` for appending PDF operators. Page implementation wraps this.

```go
type pageImpl struct {
    writer  *pdfcpu.PageContentWriter
    context *RenderingContext
    
    // Cached state
    fillColor   Color
    strokeColor Color
    lineWidth   float64
    dashPattern []float64
    dashPhase   float64
}
```

### Coordinate Conversion
- DrawingML uses EMU (1 inch = 914,400 EMU)
- PDF uses points (1 inch = 72 points)
- Conversion: `points = emu / 12,700`

- DrawingML origin: top-left
- PDF origin: bottom-left
- Transform: `pdfY = pageHeight - drawingMLY`

**Implementation**: Coordinate conversion happens in Page implementation methods, transparent to callers.

### Graphics State Management
PDF uses state stack (q/Q operators):
- `PushState()` → `q` operator
- `PopState()` → `Q` operator

Page tracks current state to avoid redundant operators (optimization).

## Renderer Integration

### Chart Renderer Example (Actual from .wip files)
```go
func (r *ChartRenderer) RenderBarChart(x, y, width, height float64, data ChartData, horizontal bool) error {
    page := r.ctx.Page  // Access Page from context

    // Set fill color and draw bars
    for seriesIdx, series := range data.Series {
        page.SetFillColor(series.Color.R, series.Color.G, series.Color.B)

        for catIdx, val := range series.Values {
            barX, barY, barW, barH := calculateBarGeometry(...)

            // Draw filled and stroked rectangle
            page.DrawRectangle(barX, barY, barW, barH, true, true)
        }
    }

    // Draw legend background
    page.SetFillColor(1, 1, 1)
    page.SetStrokeColor(0, 0, 0)
    page.DrawRectangle(legendX, legendY, legendWidth, legendHeight, true, true)

    // Draw legend text
    page.SetFont("Helvetica", 10)
    page.SetFillColor(0, 0, 0)
    page.DrawText(series.Name, textX, textY)

    return nil
}
```

### Fill Renderer Example (Actual from .wip files)
```go
func (r *SolidFillRenderer) Apply(ctx *core.RenderingContext, path *PathBuilder) error {
    page := ctx.Page

    // Get color from DrawingML element
    color := parseColor(r.fill)

    // Set fill color (RGB floats)
    page.SetFillColor(color.R, color.G, color.B)

    // Write path fill operation
    page.WriteContent(path.Fill())  // Uses low-level content stream access

    return nil
}
```

### Shape Renderer Example (Actual from .wip files)
```go
func (r *ShapeRenderer) RenderShapeWithFill(sp *ShapeProperties, fill FillRenderer, stroke StrokeRenderer) error {
    // Save graphics state
    r.ctx.Page.SaveGraphicsState()
    defer r.ctx.Page.RestoreGraphicsState()

    // Create the shape path
    path := NewPathBuilder()
    createShapePath(sp, path)  // Builds complex bezier path

    // Apply fill
    if fill != nil {
        fill.Apply(r.ctx, path)  // Calls SetFillColor + WriteContent
    }

    // Apply stroke
    if stroke != nil {
        stroke.Apply(r.ctx, path)  // Calls SetStrokeColor + WriteContent
    }

    return nil
}
```

Note: These examples show the actual API usage from the .wip renderer files, including:
- RGB float parameters for SetFillColor/SetStrokeColor
- Boolean fill/stroke parameters for DrawRectangle
- SaveGraphicsState/RestoreGraphicsState for state isolation
- WriteContent() for complex path operations
- SetFont before DrawText

## Testing Strategy

### Unit Tests (Fast, Isolated)
Mock Page implementation records method calls:
```go
type MockPage struct {
    Calls []string
}

func (m *MockPage) DrawRectangle(x, y, w, h float64) {
    m.Calls = append(m.Calls, fmt.Sprintf("DrawRectangle(%v,%v,%v,%v)", x, y, w, h))
}

func TestChartRenderer(t *testing.T) {
    mock := &MockPage{}
    ctx := &core.RenderingContext{Page: mock}
    
    renderer.Render(ctx, chart)
    
    assert.Contains(t, mock.Calls, "DrawRectangle(10,20,100,50)")
}
```

### Integration Tests (Realistic)
Generate actual PDFs, validate structure:
```go
func TestChartPDFGeneration(t *testing.T) {
    doc := spreadsheet.Open("testdata/chart.xlsx")
    pdfBytes := pdf.Render(doc)
    
    // Validate PDF structure
    assert.Contains(pdfBytes, "/Type /XObject")  // Chart embedded
    assert.True(pdf.IsValidPDF(pdfBytes))
}
```

### Fidelity Tests (Visual Validation)
Compare with Microsoft Office output:
```go
func TestChartFidelity(t *testing.T) {
    goffice PDF := pdf.RenderExcel("sales.xlsx")
    baseline PDF := loadBaseline("sales-msoffice.pdf")
    
    // Visual comparison (pixel diff with tolerance)
    diff := pdf.CompareVisual(goffice PDF, baseline PDF)
    assert.Less(diff.Percentage, 2.0, "Visual difference should be <2%")
}
```

## Performance Considerations

### State Change Optimization
Track current state to avoid redundant PDF operators:
```go
func (p *pageImpl) SetFillColor(color Color) {
    if p.fillColor.Equals(color) {
        return  // Skip redundant operator
    }
    p.fillColor = color
    p.writer.WriteOperator("...")  // Actually write to PDF
}
```

### Path Composition
Build complex paths before rendering to minimize PDF operators:
```go
// Good: One path, one fill operation
path := NewPathBuilder()
path.AddRectangle(...)
path.AddEllipse(...)
page.SetFillColor(red)
page.DrawPath(path)

// Bad: Multiple fill operations
page.SetFillColor(red)
page.DrawRectangle(...)
page.SetFillColor(red)  // Redundant
page.DrawEllipse(...)
```

## Migration Path

### Phase 1: Core Infrastructure
1. Implement Page interface
2. Add to RenderingContext
3. Unit tests with mock

### Phase 2: Renderer Re-enablement
1. Remove `.wip` extensions
2. Fix compilation errors
3. Basic validation (compiles, simple tests pass)

### Phase 3: Integration
1. Hook up Word/Excel/PowerPoint renderers
2. End-to-end tests
3. Fidelity validation

### Phase 4: Optimization
1. Profile rendering performance
2. Optimize state changes
3. Benchmark critical paths

## Risks & Mitigation

### Risk: pdfcpu API Changes
**Impact**: Medium
**Probability**: Low (stable library)
**Mitigation**: Page abstraction insulates renderers; only Page implementation needs updates

### Risk: PDF Compliance Issues
**Impact**: High (PDFs don't render correctly)
**Probability**: Medium (PDF is complex)
**Mitigation**: 
- Extensive testing with Adobe Reader, macOS Preview, LibreOffice
- Reference PDF spec for all operators
- Fidelity test suite catches regressions

### Risk: Performance Degradation
**Impact**: Medium
**Probability**: Low
**Mitigation**:
- State change optimization
- Benchmarks for critical rendering paths
- Profile before optimizing

### Risk: Coordinate Bugs
**Impact**: High (shapes misaligned)
**Probability**: Medium (coordinate systems are tricky)
**Mitigation**:
- Comprehensive unit tests for coordinate conversion
- Visual validation tests
- Test with various page sizes and orientations

## Success Metrics

- [ ] Zero `.wip` files remaining in `pdf/drawing/`
- [ ] >90% of charts from examples render correctly
- [ ] <5% visual difference from Microsoft Office output
- [ ] All PDF validators (pdfinfo, qpdf) report valid PDFs
- [ ] Rendering performance <2x slower than Word-only (acceptable for added features)

## Future Enhancements (Not in Scope)

1. **Clipping Paths**: For cropped images, complex masks
2. **Advanced Blend Modes**: Multiply, screen, overlay, etc.
3. **Mesh Gradients**: For smooth color transitions
4. **Soft Masks**: Advanced transparency
5. **Type3 Fonts**: For special effects
6. **Optional Content Groups**: For layers/visibility

These can be added incrementally by extending Page interface without breaking existing renderers.
