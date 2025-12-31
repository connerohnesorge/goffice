# Tasks: Enable Chart Rendering Infrastructure

**Baseline**: ~3,543 LOC exists in `pdf/drawing/*.wip` files. This proposal **enables** existing disabled code by implementing the missing Page API.

## 1. Phase 1: Core Page API (40-60 hours total)

### 1.1 Page Interface Definition (8-12h)

- [ ] 1.1.1 Create `pdf/core/page.go`
- [ ] 1.1.2 Define `Page` interface with graphics state methods (SaveGraphicsState, RestoreGraphicsState)
- [ ] 1.1.3 Add color methods to Page interface (SetFillColor, SetFillColorRGBA, SetStrokeColor, SetStrokeColorRGBA)
- [ ] 1.1.4 Add line property methods (SetLineWidth, SetLineCap, SetLineJoin, SetLineDashPattern)
- [ ] 1.1.5 Add shape methods (DrawRectangle with fill/stroke bools, DrawRoundedRectangle, DrawCircle with fill/stroke bools, DrawEllipse, DrawLine, DrawPath)
- [ ] 1.1.6 Add fill/stroke methods (Fill, Stroke, FillStroke)
- [ ] 1.1.7 Add text methods (SetFont with name string, DrawText, DrawTextRotated)
- [ ] 1.1.8 Add image method (DrawImage)
- [ ] 1.1.9 Add transformation methods (Translate, Scale, Rotate)
- [ ] 1.1.10 Add clipping methods (Clip, SetClipRect, ClearClip)
- [ ] 1.1.11 Add content stream method (WriteContent)
- [ ] 1.1.12 Document Path, Font, Image type references (Path = *PathBuilder from pdf/drawing)

### 1.2 PDF Page Implementation (15-20h)

- [ ] 1.2.1 Create `pdf/core/pdf_page.go`
- [ ] 1.2.2 Define `PDFPage` struct with doc, pageObj, content, fonts, images fields
- [ ] 1.2.3 Implement `NewPDFPage(doc *Document, width, height float64) *PDFPage`
- [ ] 1.2.4 Implement `SaveGraphicsState()` - write "q" operator
- [ ] 1.2.5 Implement `RestoreGraphicsState()` - write "Q" operator
- [ ] 1.2.6 Implement `SetFillColor(r, g, b float64)` - write "rg" operator
- [ ] 1.2.7 Implement `SetFillColorRGBA(r, g, b, a float64)` - handle alpha via ExtGState
- [ ] 1.2.8 Implement `SetStrokeColor(r, g, b float64)` - write "RG" operator
- [ ] 1.2.9 Implement `SetStrokeColorRGBA(r, g, b, a float64)`
- [ ] 1.2.10 Implement `SetLineWidth(width float64)` - write "w" operator
- [ ] 1.2.11 Implement `SetLineCap(cap int)` - write "J" operator (0=butt, 1=round, 2=square)
- [ ] 1.2.12 Implement `SetLineJoin(join int)` - write "j" operator (0=miter, 1=round, 2=bevel)
- [ ] 1.2.13 Implement `SetLineDashPattern(pattern []float64, phase float64)` - write "d" operator
- [ ] 1.2.14 Implement `DrawRectangle(x, y, w, h float64, fill, stroke bool)` - write "re" operator + paint ops
- [ ] 1.2.15 Implement `DrawRoundedRectangle(x, y, w, h, radius float64)` - path construction
- [ ] 1.2.16 Implement `DrawCircle(cx, cy, r float64, fill, stroke bool)` - Bezier approximation + paint ops
- [ ] 1.2.17 Implement `DrawEllipse(cx, cy, rx, ry float64)` - Bezier approximation
- [ ] 1.2.18 Implement `DrawLine(x1, y1, x2, y2 float64)` - "m" and "l" operators
- [ ] 1.2.19 Implement `DrawPath(path *Path)` - complex path rendering
- [ ] 1.2.20 Implement `Fill()` - write "f" operator
- [ ] 1.2.21 Implement `Stroke()` - write "S" operator
- [ ] 1.2.22 Implement `FillStroke()` - write "B" operator
- [ ] 1.2.23 Implement `Clip()` - write "W n" operator
- [ ] 1.2.24 Implement `SetFont(name string, size float64)` - write "Tf" operator
- [ ] 1.2.25 Implement `DrawText(x, y float64, text string)` - "BT/ET" text block
- [ ] 1.2.26 Implement `DrawTextRotated(x, y, angle float64, text string)` - with transform matrix
- [ ] 1.2.27 Implement `DrawImage(img *Image, x, y, w, h float64)` - XObject reference
- [ ] 1.2.28 Implement `Translate(dx, dy float64)` - "cm" operator
- [ ] 1.2.29 Implement `Scale(sx, sy float64)` - "cm" operator
- [ ] 1.2.30 Implement `Rotate(angle float64)` - "cm" operator
- [ ] 1.2.31 Implement `SetClipRect(x, y, w, h float64)` - "W" operator
- [ ] 1.2.32 Implement `ClearClip()` - restore graphics state
- [ ] 1.2.33 Implement `WriteContent(content string)` - raw content stream
- [ ] 1.2.34 Write 15 unit tests for PDFPage implementation

### 1.3 Context Integration (4-6h)

- [ ] 1.3.1 Add `Page Page` field to `RenderingContext` struct in `pdf/core/context.go`
- [ ] 1.3.2 Add `SetPage(page Page)` method to `RenderingContext`
- [ ] 1.3.3 Update `NewRenderingContext` to initialize Page as nil
- [ ] 1.3.4 Add page creation helper to context or document
- [ ] 1.3.5 Write 4 unit tests for context integration

### 1.4 Enable Renderers (8-12h)

- [ ] 1.4.1 Rename `pdf/drawing/chart_renderer.go.wip` to `chart_renderer.go`
- [ ] 1.4.2 Rename `pdf/drawing/shape_renderer.go.wip` to `shape_renderer.go`
- [ ] 1.4.3 Rename `pdf/drawing/fill_renderer.go.wip` to `fill_renderer.go`
- [ ] 1.4.4 Rename `pdf/drawing/stroke_renderer.go.wip` to `stroke_renderer.go`
- [ ] 1.4.5 Rename `pdf/drawing/effects_renderer.go.wip` to `effects_renderer.go`
- [ ] 1.4.6 Rename `pdf/drawing/transform_renderer.go.wip` to `transform_renderer.go`
- [ ] 1.4.7 Rename `pdf/drawing/text_in_shape.go.wip` to `text_in_shape.go`
- [ ] 1.4.8 Rename `pdf/drawing/image_renderer.go.wip` to `image_renderer.go`
- [ ] 1.4.9 Fix any `ctx.Page` references to match new API
- [ ] 1.4.10 Fix any type mismatches between .wip code and new Page interface
- [ ] 1.4.11 Resolve any import issues
- [ ] 1.4.12 Run `go build ./pdf/drawing/...` and fix compilation errors

### 1.5 Enable Tests (4-6h)

- [ ] 1.5.1 Rename `pdf/drawing/chart_renderer_test.go.wip` to `chart_renderer_test.go`
- [ ] 1.5.2 Rename `pdf/drawing/shape_renderer_test.go.wip` to `shape_renderer_test.go`
- [ ] 1.5.3 Rename `pdf/drawing/fill_renderer_test.go.wip` to `fill_renderer_test.go`
- [ ] 1.5.4 Rename `pdf/drawing/drawingml_integration_test.go.wip` to `drawingml_integration_test.go`
- [ ] 1.5.5 Update test imports if needed
- [ ] 1.5.6 Create mock Page implementation for unit tests
- [ ] 1.5.7 Run `go test ./pdf/drawing/...` and verify all tests pass

## 2. Phase 2: Integration Testing (12-18 hours total)

### 2.1 Chart Integration Tests (6-8h)

- [ ] 2.1.1 Create `pdf/drawing/chart_integration_test.go`
- [ ] 2.1.2 Test bar chart rendering with sample data
- [ ] 2.1.3 Test line chart rendering with markers
- [ ] 2.1.4 Test pie chart rendering with slices
- [ ] 2.1.5 Test chart title rendering
- [ ] 2.1.6 Test chart legend rendering
- [ ] 2.1.7 Test chart axis labels

### 2.2 Shape Integration Tests (4-6h)

- [ ] 2.2.1 Create `pdf/drawing/shape_integration_test.go`
- [ ] 2.2.2 Test rectangle with solid fill
- [ ] 2.2.3 Test rectangle with gradient fill
- [ ] 2.2.4 Test circle/ellipse rendering
- [ ] 2.2.5 Test shape with stroke
- [ ] 2.2.6 Test shape with shadow effect
- [ ] 2.2.7 Test shape with rotation transform

### 2.3 Presentation Integration (4-6h)

- [ ] 2.3.1 Test pptx-charts example renders charts to PDF
- [ ] 2.3.2 Verify bar chart visible in output PDF
- [ ] 2.3.3 Verify pie chart visible in output PDF
- [ ] 2.3.4 Verify line chart visible in output PDF
- [ ] 2.3.5 Compare output to LibreOffice reference visually

## 3. Phase 3: Fidelity Testing (15-25 hours total)

### 3.1 Reference Image Generation (8-12h)

- [ ] 3.1.1 Create reference images from Excel charts
- [ ] 3.1.2 Create reference images from PowerPoint charts
- [ ] 3.1.3 Save references in `testdata/pdf_drawing/`
- [ ] 3.1.4 Document reference creation process

### 3.2 Image Comparison Tests (8-12h)

- [ ] 3.2.1 Implement image diff utility
- [ ] 3.2.2 Add tolerance threshold for acceptable differences
- [ ] 3.2.3 Test bar chart against reference
- [ ] 3.2.4 Test line chart against reference
- [ ] 3.2.5 Test pie chart against reference
- [ ] 3.2.6 Document known differences

### 3.3 Regression Prevention (4-6h)

- [ ] 3.3.1 Add golden file tests for chart output
- [ ] 3.3.2 Add CI job for fidelity tests
- [ ] 3.3.3 Verify Word PDF rendering not regressed
- [ ] 3.3.4 Performance benchmark: charts <100ms each

## Summary

| Phase | Hours | Key Deliverables |
|-------|-------|------------------|
| Phase 1 | 40-60h | Page interface, PDFPage impl, enable 3,543 LOC |
| Phase 2 | 12-18h | Integration tests for charts, shapes, presentations |
| Phase 3 | 15-25h | Fidelity testing, reference images, regression prevention |
| **Total** | **67-103h** | **Full chart rendering capability** |

## Dependencies

- **No external dependencies**: Uses existing pdf/core infrastructure
- **Internal dependencies**:
  - 1.2 depends on 1.1 (implementation needs interface)
  - 1.3 depends on 1.1 (context needs Page type)
  - 1.4 depends on 1.2 and 1.3 (renderers need working Page)
  - 1.5 depends on 1.4 (tests need compiled renderers)
  - Phase 2 depends on Phase 1 completion
  - Phase 3 depends on Phase 2 completion

## Parallelization

- Tasks 1.1 and initial 1.2 can run in parallel
- Tasks 1.4.1-1.4.8 (renames) can run in parallel
- Phase 2 tasks (2.1-2.3) can run in parallel
- Phase 3.1 and 3.2 can run in parallel after Phase 2
