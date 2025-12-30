# Implementation Tasks

## Phase 1: Core Page Abstraction (2 weeks)

- [ ] **Define Page interface** in `pdf/core/page.go`
  - Drawing primitives (DrawRectangle, DrawEllipse, DrawPath)
  - State management (SetFillColor, SetStrokeColor, SetLineWidth, SetLineDashPattern)
  - Graphics state stack (PushState, PopState, Transform)
  - Content embedding (AddImage, DrawText)
  - Validation: Interface compiles, godoc looks correct

- [ ] **Create pdfcpu adapter implementation** in `pdf/core/page_impl.go`
  - Implement Page interface wrapping pdfcpu.PageContentWriter
  - Add coordinate conversion (EMU→points, top-left→bottom-left)
  - Add state tracking to avoid redundant PDF operators
  - Validation: Implementation compiles, satisfies Page interface

- [ ] **Add Page field to RenderingContext** in `pdf/core/render_context.go`
  - Add `Page Page` field to struct
  - Initialize Page in NewRenderingContext
  - Update existing code that creates RenderingContext
  - Validation: All callers compile, tests pass

- [ ] **Create MockPage for testing** in `pdf/core/page_mock_test.go`
  - Implement Page interface recording method calls
  - Add assertion helpers for common patterns
  - Validation: Mock is usable in tests

- [ ] **Write Page unit tests** in `pdf/core/page_test.go`
  - Test each Page method with pdfcpu backend
  - Test coordinate conversion accuracy
  - Test state stack push/pop behavior
  - Test redundant state change optimization
  - Validation: All tests pass, >90% coverage

## Phase 2: Re-enable Renderers (2 weeks)

- [ ] **Re-enable chart_renderer.go**
  - Rename `pdf/drawing/chart_renderer.go.wip` → `chart_renderer.go`
  - Fix any Page API mismatches
  - Fix import statements
  - Validation: File compiles without errors

- [ ] **Re-enable fill_renderer.go**
  - Rename `pdf/drawing/fill_renderer.go.wip` → `fill_renderer.go`
  - Fix any Page API mismatches
  - Update gradient rendering to use Page primitives
  - Validation: File compiles

- [ ] **Re-enable shape_renderer.go**
  - Rename `pdf/drawing/shape_renderer.go.wip` → `shape_renderer.go`
  - Fix any Page API mismatches
  - Validation: File compiles

- [ ] **Re-enable stroke_renderer.go**
  - Rename `pdf/drawing/stroke_renderer.go.wip` → `stroke_renderer.go`
  - Fix any Page API mismatches
  - Validation: File compiles

- [ ] **Re-enable image_renderer.go**
  - Rename `pdf/drawing/image_renderer.go.wip` → `image_renderer.go`
  - Fix any Page API mismatches
  - Validation: File compiles

- [ ] **Re-enable text_in_shape.go**
  - Rename `pdf/drawing/text_in_shape.go.wip` → `text_in_shape.go`
  - Fix any Page API mismatches
  - Validation: File compiles

- [ ] **Re-enable transform_renderer.go**
  - Rename `pdf/drawing/transform_renderer.go.wip` → `transform_renderer.go`
  - Fix any Page API mismatches
  - Validation: File compiles

- [ ] **Re-enable effects_renderer.go**
  - Rename `pdf/drawing/effects_renderer.go.wip` → `effects_renderer.go`
  - Fix any Page API mismatches
  - Validation: File compiles

- [ ] **Re-enable test files**
  - Rename all `*_test.go.wip` → `*_test.go` (4 files)
  - Fix test setup to use MockPage or real Page
  - Validation: Tests compile and pass

- [ ] **Update WIP_RENDERERS.md**
  - Document that renderers are now enabled
  - Remove file from pdf/drawing/ (or archive)
  - Validation: Documentation reflects current state

## Phase 3: Integration & Testing (1.5 weeks)

- [ ] **Integrate chart rendering in Excel renderer**
  - Update `pdf/spreadsheet/renderer.go` to call chart_renderer
  - Pass RenderingContext with Page to chart renderer
  - Validation: Excel with charts compiles

- [ ] **Integrate shape rendering in PowerPoint renderer**
  - Update `pdf/presentation/renderer.go` to call shape_renderer
  - Pass RenderingContext with Page
  - Validation: PowerPoint with shapes compiles

- [ ] **Integrate image rendering in Word renderer**
  - Update `pdf/word/renderer.go` to call image_renderer for inline images
  - Validation: Word with images compiles

- [ ] **Create end-to-end Excel chart test**
  - Test file: `pdf/spreadsheet/chart_integration_test.go`
  - Load Excel with column chart, render to PDF, validate output
  - Validation: Test passes, PDF contains chart XObject

- [ ] **Create end-to-end PowerPoint shape test**
  - Test file: `pdf/presentation/shape_integration_test.go`
  - Load PowerPoint with shapes, render to PDF, validate
  - Validation: Test passes, shapes visible in PDF

- [ ] **Create end-to-end Word image test**
  - Test file: `pdf/word/image_integration_test.go`
  - Load Word with images, render to PDF, validate
  - Validation: Test passes, images embedded in PDF

- [ ] **Run fidelity tests with DrawingML elements**
  - Add chart/shape test cases to pdf/fidelity_test.go
  - Generate baseline PDFs from Microsoft Office
  - Visual comparison with tolerance
  - Validation: Fidelity within 5% of baseline

- [ ] **Validate with PDF tools**
  - Run pdfinfo on generated PDFs (check for errors)
  - Run qpdf --check on generated PDFs
  - Open in Adobe Reader, macOS Preview, LibreOffice
  - Validation: No errors, renders correctly in all viewers

## Phase 4: Documentation (0.5 weeks)

- [ ] **Document Page interface** in `pdf/core/page.go`
  - Add godoc comments to all Page methods
  - Add package-level documentation explaining Page abstraction
  - Include usage examples in comments
  - Validation: godoc pdf/core looks professional

- [ ] **Update LIMITATIONS.md**
  - Remove "Charts not rendered" limitation
  - Remove "Shapes not rendered" limitation
  - Add note about gradient approximation
  - Add note about deferred advanced features (clipping, blend modes)
  - Validation: Documentation accurate

- [ ] **Update FIDELITY.md**
  - Add section on DrawingML rendering quality
  - Document gradient approximation approach
  - Document known differences from Office (if any)
  - Validation: Users understand rendering quality

- [ ] **Create developer guide** in `pdf/DRAWING_RENDERING.md` (new file)
  - Explain Page abstraction architecture
  - Show how to add new DrawingML renderers
  - Include code examples
  - Validation: Guide is helpful for contributors

- [ ] **Update examples**
  - Ensure examples/ml_results includes chart rendering
  - Ensure examples/pptx-charts works with PDF output
  - Add comments explaining DrawingML rendering
  - Validation: Examples work correctly

## Validation Checkpoints

After Phase 1:
- [ ] Page interface is complete and tested
- [ ] MockPage enables renderer unit testing
- [ ] RenderingContext integrates Page cleanly

After Phase 2:
- [ ] All 12 .wip files are renamed and compiling
- [ ] Zero compile errors in pdf/drawing/
- [ ] Renderer unit tests pass

After Phase 3:
- [ ] Charts render in Excel PDFs
- [ ] Shapes render in PowerPoint PDFs
- [ ] Images render in Word PDFs
- [ ] End-to-end tests pass
- [ ] PDFs validate correctly

After Phase 4:
- [ ] Documentation is complete and accurate
- [ ] Examples demonstrate new capabilities
- [ ] Developers can extend DrawingML rendering

## Dependencies

**Blocked by**: None (all dependencies already exist)

**Blocks**:
- SmartArt rendering (needs DrawingML rendering first)
- Advanced PDF features (clipping, blend modes can build on Page API)

## Parallelizable Work

- Phase 1 tasks are sequential (interface → implementation → tests)
- Phase 2 renderer re-enablement can be parallelized (each file independent)
- Phase 3 integration tests can be parallelized by document type
- Phase 4 documentation tasks can be parallelized

## Risk Mitigation

- **Risk**: pdfcpu API insufficient for needed operations
  - **Mitigation**: Prototype Page implementation early, validate all operations work
  
- **Risk**: Coordinate conversion bugs cause misalignment
  - **Mitigation**: Comprehensive unit tests, visual validation
  
- **Risk**: PDF compliance issues
  - **Mitigation**: Validate with multiple PDF tools, test in multiple viewers

## Success Criteria

All tasks above completed AND:
- [ ] Zero .wip files in pdf/drawing/
- [ ] All tests pass (unit, integration, fidelity)
- [ ] PDFs render correctly in Adobe Reader, Preview, LibreOffice
- [ ] Visual fidelity within 5% of Microsoft Office baseline
- [ ] Documentation complete and accurate
