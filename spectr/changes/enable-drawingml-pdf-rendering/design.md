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
    // DrawRectangle draws a filled/stroked rectangle
    DrawRectangle(x, y, width, height float64)
    
    // DrawEllipse draws a filled/stroked ellipse
    DrawEllipse(cx, cy, rx, ry float64)
    
    // DrawPath renders an arbitrary path
    DrawPath(path *PathBuilder)
    
    // === State Management ===
    // SetFillColor sets the fill color for subsequent shapes
    SetFillColor(color Color)
    
    // SetStrokeColor sets the stroke color
    SetStrokeColor(color Color)
    
    // SetLineWidth sets stroke width in points
    SetLineWidth(width float64)
    
    // SetLineDashPattern sets dash pattern (empty = solid)
    SetLineDashPattern(pattern []float64, phase float64)
    
    // === Graphics State Stack ===
    // PushState saves current graphics state
    PushState()
    
    // PopState restores previous graphics state
    PopState()
    
    // Transform applies transformation matrix to CTM
    Transform(matrix Matrix)
    
    // === Content Embedding ===
    // AddImage embeds and draws an image
    AddImage(img Image, x, y, width, height float64)
    
    // DrawText renders text at position
    DrawText(text string, x, y float64, font Font, size float64)
}
```

### Not Included (Deferred)
- Clipping paths (future enhancement)
- Blend modes beyond normal (future)
- Gradients as first-class (renderers compose from primitives)
- Patterns as first-class (renderers use repeated drawing)

**Rationale**: Start minimal, expand based on real needs. Most DrawingML features can be built by composing primitives.

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

### Chart Renderer Example
```go
func (r *ChartRenderer) Render(ctx *core.RenderingContext, chart *Chart) error {
    page := ctx.Page  // Access Page from context
    
    // Render chart background
    page.SetFillColor(chart.BackgroundColor)
    page.DrawRectangle(chart.X, chart.Y, chart.Width, chart.Height)
    
    // Render series bars
    for _, series := range chart.Series {
        page.SetFillColor(series.Color)
        for i, value := range series.Values {
            barX, barY, barW, barH := r.calculateBarGeometry(i, value)
            page.DrawRectangle(barX, barY, barW, barH)
        }
    }
    
    return nil
}
```

### Fill Renderer Example
```go
func (r *FillRenderer) RenderGradient(ctx *core.RenderingContext, gradient *Gradient, bounds Rect) error {
    page := ctx.Page
    
    // PDF gradients require shading patterns - decompose into steps
    steps := r.approximateGradient(gradient, 10) // 10 color stops
    
    for i, step := range steps {
        page.SetFillColor(step.Color)
        stepRect := r.calculateStepRect(bounds, i, len(steps), gradient.Angle)
        page.DrawRectangle(stepRect.X, stepRect.Y, stepRect.W, stepRect.H)
    }
    
    return nil
}
```

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
