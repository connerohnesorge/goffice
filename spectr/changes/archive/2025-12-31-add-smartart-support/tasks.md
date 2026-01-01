# Implementation Tasks

## Phase 1: Schema Generation & Code Generation (Week 1)

### 1.1 Update Generator to Load Diagram Schemas
- [ ] Modify `cmd/gen-go-drawingml/main.go` to add diagram schema loading
- [ ] Add `loadDiagramSchemas()` function to load all 5 diagram schema files
- [ ] Update schema path configuration to include diagram schemas
- [ ] Validation: Run generator, verify schemas load without errors

### 1.2 Generate Diagram Element Package
- [ ] Configure generator to output to `drawingml/diagram/` package
- [ ] Run generator with diagram schemas
- [ ] Verify generated files compile without errors
- [ ] Validation: `go build ./drawingml/diagram` succeeds

### 1.3 Verify Generated Elements
- [ ] Review generated element classes for correctness
- [ ] Verify key elements exist: DataModelRoot, Point, Connection, LayoutDefinition, etc.
- [ ] Check XML struct tags are correct
- [ ] Run golangci-lint on generated code
- [ ] Validation: `golangci-lint run ./drawingml/diagram` passes

### 1.4 Add Package Documentation
- [ ] Create `drawingml/diagram/doc.go` with package overview
- [ ] Document dgm: namespace purpose
- [ ] Add usage examples in package doc
- [ ] Validation: `go doc drawingml/diagram` shows helpful documentation

## Phase 2: Part Infrastructure (Week 2)

### 2.1 Add Diagram Content Types
- [ ] Add diagram content type constants to `packaging/content_types.go`
- [ ] Add constants for all four part types (data, layout, style, colors)
- [ ] Validation: Constants compile and follow naming conventions

### 2.2 Add Diagram Relationship Types
- [ ] Add relationship type constants to `packaging/relationship_types.go`
- [ ] Add RelationshipTypeDiagramData
- [ ] Add RelationshipTypeDiagramLayout
- [ ] Add RelationshipTypeDiagramColors
- [ ] Add RelationshipTypeDiagramQuickStyle
- [ ] Validation: Constants compile

### 2.3 Implement DiagramDataPart for Presentation
- [ ] Create `presentation/parts/diagram_data_part.go`
- [ ] Implement DiagramDataPart struct with lazy loading
- [ ] Implement NewDiagramDataPart constructor
- [ ] Implement DataModel() getter with XML parsing (returns *diagram.DataModelRoot)
- [ ] Implement Save() method
- [ ] Implement relationships to layout/colors/style parts
- [ ] Validation: Part compiles and follows existing part patterns

### 2.4 Implement DiagramLayoutDefinitionPart for Presentation
- [ ] Create `presentation/parts/diagram_layout_definition_part.go`
- [ ] Implement DiagramLayoutDefinitionPart struct
- [ ] Implement NewDiagramLayoutDefinitionPart constructor
- [ ] Implement LayoutDefinition() getter (returns *diagram.LayoutDefinition)
- [ ] Implement Save() method
- [ ] Validation: Part compiles

### 2.5 Implement DiagramStylePart for Presentation
- [ ] Create `presentation/parts/diagram_style_part.go`
- [ ] Implement DiagramStylePart struct
- [ ] Implement NewDiagramStylePart constructor
- [ ] Implement StyleData() getter (returns *diagram.StyleData)
- [ ] Implement Save() method
- [ ] Validation: Part compiles

### 2.6 Implement DiagramColorsPart for Presentation
- [ ] Create `presentation/parts/diagram_colors_part.go`
- [ ] Implement DiagramColorsPart struct
- [ ] Implement NewDiagramColorsPart constructor
- [ ] Implement ColorsDefinition() getter (returns *diagram.ColorsDefinition)
- [ ] Implement Save() method
- [ ] Validation: Part compiles

### 2.7 Add Diagram Part Discovery to PresentationDocument
- [ ] Modify `presentation/presentation_document.go`
- [ ] Add DiagramDataParts() method
- [ ] Implement part discovery via relationships
- [ ] Validation: Method compiles and follows existing patterns

### 2.8 Implement Diagram Parts for Word
- [ ] Create `wordprocessing/parts/diagram_data_part.go`
- [ ] Create `wordprocessing/parts/diagram_layout_definition_part.go`
- [ ] Create `wordprocessing/parts/diagram_style_part.go`
- [ ] Create `wordprocessing/parts/diagram_colors_part.go`
- [ ] Add DiagramDataParts() to WordprocessingDocument
- [ ] Validation: All parts compile

### 2.9 Implement Diagram Parts for Excel
- [ ] Create `spreadsheet/parts/diagram_data_part.go`
- [ ] Create `spreadsheet/parts/diagram_layout_definition_part.go`
- [ ] Create `spreadsheet/parts/diagram_style_part.go`
- [ ] Create `spreadsheet/parts/diagram_colors_part.go`
- [ ] Add DiagramDataParts() to SpreadsheetDocument
- [ ] Validation: All parts compile

## Phase 3: Testing (Week 3)

### 3.1 Create Test Documents
- [ ] Create `testdata/smartart/org-chart.pptx` with organizational chart
- [ ] Create `testdata/smartart/process-flow.pptx` with process diagram
- [ ] Create `testdata/smartart/cycle-diagram.pptx` with cycle diagram
- [ ] Create `testdata/smartart/hierarchy.docx` with hierarchy in Word
- [ ] Create `testdata/smartart/list-diagram.xlsx` with list in Excel
- [ ] Validation: Documents open in Office without errors

### 3.2 Unit Tests - Element Parsing
- [ ] Create `drawingml/diagram/data_model_test.go`
- [ ] Test DataModelRoot XML parsing
- [ ] Test Point XML parsing
- [ ] Test Connection XML parsing
- [ ] Test LayoutDefinition XML parsing
- [ ] Test StyleData XML parsing
- [ ] Test ColorsDefinition XML parsing
- [ ] Validation: All element parsing tests pass

### 3.3 Unit Tests - Part Functionality
- [ ] Create `presentation/parts/diagram_part_test.go`
- [ ] Test DiagramDataPart creation
- [ ] Test DataModel lazy loading
- [ ] Test Save/Load roundtrip
- [ ] Test relationship management
- [ ] Validation: All part tests pass

### 3.4 Integration Tests - PowerPoint Roundtrip
- [ ] Create `presentation/smartart_integration_test.go`
- [ ] Test opening PowerPoint with org chart SmartArt
- [ ] Test diagram part discovery
- [ ] Test data model parsing
- [ ] Test save and reload
- [ ] Test XML byte-for-byte comparison
- [ ] Validation: PowerPoint SmartArt roundtrips correctly

### 3.5 Integration Tests - Word Roundtrip
- [ ] Create `wordprocessing/smartart_integration_test.go`
- [ ] Test opening Word with hierarchy SmartArt
- [ ] Test diagram part discovery
- [ ] Test save and reload
- [ ] Validation: Word SmartArt roundtrips correctly

### 3.6 Integration Tests - Excel Roundtrip
- [ ] Create `spreadsheet/smartart_integration_test.go`
- [ ] Test opening Excel with list SmartArt
- [ ] Test diagram part discovery
- [ ] Test save and reload
- [ ] Validation: Excel SmartArt roundtrips correctly

### 3.7 Office Compatibility Tests
- [ ] Open goffice-saved PowerPoint in Office 2016
- [ ] Verify SmartArt displays correctly
- [ ] Verify SmartArt is editable in Office
- [ ] Test with Office 2019 and Office 365
- [ ] Validation: All Office versions can open and edit saved SmartArt

### 3.8 Regression Tests
- [ ] Run full test suite: `go test ./...`
- [ ] Verify no existing tests break
- [ ] Check for memory leaks with race detector: `go test -race ./...`
- [ ] Run benchmarks to check for performance regression
- [ ] Validation: Zero regressions

## Phase 4: PDF Rendering (Minimal - Week 3)

### 4.1 Implement Diagram Placeholder Renderer
- [ ] Create `pdf/drawing/diagram_renderer.go`
- [ ] Implement DiagramRenderer struct
- [ ] Implement Render() method (draws bounding box placeholder)
- [ ] Add dashed border rendering
- [ ] Add "SmartArt Diagram" text label
- [ ] Validation: Renderer compiles

### 4.2 Integrate Diagram Renderer
- [ ] Hook up renderer in `pdf/presentation/renderer.go`
- [ ] Hook up renderer in `pdf/word/renderer.go`
- [ ] Hook up renderer in `pdf/spreadsheet/renderer.go`
- [ ] Validation: PDFs render with diagram placeholders

### 4.3 Test PDF Placeholder Rendering
- [ ] Create `pdf/drawing/diagram_renderer_test.go`
- [ ] Test placeholder rendering with mock Page
- [ ] Verify bounding box drawn
- [ ] Verify label text rendered
- [ ] Validation: Placeholder tests pass

### 4.4 Integration Test - PDF with SmartArt
- [ ] Test PowerPoint with SmartArt → PDF
- [ ] Test Word with SmartArt → PDF
- [ ] Test Excel with SmartArt → PDF
- [ ] Verify placeholders appear in correct locations
- [ ] Validation: PDFs contain diagram placeholders

## Phase 5: Documentation & Finalization (Week 3)

### 5.1 Update Package Documentation
- [ ] Update `drawingml/diagram/doc.go` with complete overview
- [ ] Add usage examples for reading SmartArt
- [ ] Document limitations (Phase 1: read-only, no layout engine)
- [ ] Validation: `go doc drawingml/diagram` is comprehensive

### 5.2 Update Presentation Package Docs
- [ ] Document DiagramDataParts() method
- [ ] Add SmartArt examples to presentation package docs
- [ ] Validation: `go doc presentation` includes diagram info

### 5.3 Update Word Package Docs
- [ ] Document diagram part support
- [ ] Add usage examples
- [ ] Validation: Documentation complete

### 5.4 Update Excel Package Docs
- [ ] Document diagram part support
- [ ] Add usage examples
- [ ] Validation: Documentation complete

### 5.5 Create Examples
- [ ] Create `examples/smartart-roundtrip/main.go`
- [ ] Example: Open PowerPoint with SmartArt, print structure, save
- [ ] Add comments explaining diagram parts
- [ ] Validation: Example runs successfully

### 5.6 Update CHANGELOG.md
- [ ] Add entry for SmartArt roundtrip support
- [ ] List new packages and parts
- [ ] Note limitations (no layout engine in Phase 1)
- [ ] Validation: CHANGELOG is accurate

### 5.7 Update Main README (if needed)
- [ ] Add SmartArt to feature list
- [ ] Note roundtrip support
- [ ] Note future Phase 2 for creation/rendering
- [ ] Validation: README reflects new capability

### 5.8 Create Migration Guide
- [ ] Document how to work with SmartArt in goffice
- [ ] Explain four-part structure
- [ ] Provide code examples
- [ ] Note Phase 2 plans
- [ ] Validation: Guide is helpful for users

## Validation Checkpoints

### After Week 1 (Code Generation)
- [ ] Diagram schemas load without errors
- [ ] 100+ diagram element classes generated
- [ ] Generated code compiles cleanly
- [ ] golangci-lint passes on generated code
- [ ] Package documentation exists

### After Week 2 (Part Infrastructure)
- [ ] All four part types implemented for Presentation/Word/Excel
- [ ] Part discovery methods work
- [ ] Relationship types defined
- [ ] Content types defined
- [ ] Parts follow existing patterns

### After Week 3 (Testing & Documentation)
- [ ] 20+ tests passing (unit + integration)
- [ ] PowerPoint/Word/Excel roundtrip tests pass
- [ ] Office compatibility validated
- [ ] PDF placeholder rendering works
- [ ] Documentation complete
- [ ] Examples working

## Success Criteria (All Tasks Complete)

- [ ] Zero `.wip` or placeholder files
- [ ] All tests pass: `go test ./...`
- [ ] All linters pass: `golangci-lint run ./...`
- [ ] Roundtrip tests pass for all Office applications
- [ ] Office 2016/2019/365 can open saved documents
- [ ] SmartArt is preserved (no content loss)
- [ ] PDF rendering shows placeholders
- [ ] Documentation is complete and helpful
- [ ] Examples demonstrate new capabilities
- [ ] Zero regressions in existing functionality

## Dependencies

**Blocked by**: None (all infrastructure exists)

**Blocks**:
- Phase 2: SmartArt Layout Engine (future proposal)
- Phase 2: SmartArt Creation API (future proposal)
- Phase 2: Full SmartArt PDF Rendering (future proposal)

## Parallelizable Work

- Week 2: Presentation/Word/Excel part implementations can be done in parallel
- Week 3: Unit tests, integration tests, and documentation can be parallelized
- Week 3: PDF placeholder work can be done in parallel with documentation

## Risk Mitigation

### Risk: Schema Generation Issues
**Mitigation**: Test schema loading on Day 1, Week 1
**Fallback**: Manual generator adjustments if needed

### Risk: Roundtrip Failures
**Mitigation**: Test with real Office documents early (Week 2)
**Fallback**: Debug relationship/XML issues incrementally

### Risk: Office Compatibility
**Mitigation**: Validate with multiple Office versions throughout Week 3
**Fallback**: Adjust XML serialization if compatibility issues found

## Estimated Effort

- **Week 1 (Code Generation)**: 5 days
  - Schema loading: 0.5 days
  - Element generation: 1 day
  - Validation & fixes: 2 days
  - Documentation: 1.5 days

- **Week 2 (Part Infrastructure)**: 5 days
  - Content/relationship types: 0.5 days
  - Presentation parts: 2 days
  - Word/Excel parts: 1.5 days
  - Part discovery: 1 day

- **Week 3 (Testing & Documentation)**: 5-8 days
  - Test document creation: 1 day
  - Unit tests: 1.5 days
  - Integration tests: 1.5 days
  - PDF placeholder: 1 day
  - Documentation & examples: 2-3 days
  - Office validation: 1 day

**Total: 15-18 days (3 weeks)**
