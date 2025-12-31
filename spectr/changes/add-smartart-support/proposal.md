# Change: Add SmartArt/Diagram Support

## Why

goffice currently cannot read or preserve SmartArt diagrams in Office documents. When documents containing SmartArt (organizational charts, process flows, hierarchical diagrams, etc.) are opened and saved, the SmartArt content is lost. This is a critical gap for users working with presentations, reports, and documents that rely on SmartArt for visual communication.

**Current State**:
- NamespaceDiagram constant exists but is unused
- No diagram element classes generated
- No diagram parts (DiagramDataPart, DiagramLayoutDefinitionPart, etc.)
- No SmartArt API for creation or modification
- Documents with SmartArt cannot be safely roundtripped

**Impact**: Users cannot work with PowerPoint presentations, Word documents, or Excel worksheets containing SmartArt without losing this content.

## What Changes

**Phase 1: Schema Generation & Roundtrip (This Proposal)**
- Generate diagram element classes from OpenXML schemas (dgm: namespace)
- Add diagram part support (DiagramDataPart, DiagramLayoutDefinitionPart, DiagramStylePart, DiagramColorsPart)
- Enable reading and preserving existing SmartArt in documents
- Ensure roundtrip works: open document with SmartArt → save → verify preserved

**Phase 2: Layout Engine (Future Proposal)**
- Implement layout algorithms for simple diagram types (list, basic hierarchy)
- Add programmatic SmartArt creation API
- Enable SmartArt rendering to PDF

**Out of Scope (This Proposal)**:
- SmartArt layout algorithms - deferred to future proposal
- Programmatic SmartArt creation - deferred to future proposal
- SmartArt PDF rendering - deferred to future proposal
- Full 50+ layout type support - incremental future work
- Custom layout definitions
- DiagramPersistLayoutPart (Office 2010+ cached layout) - deferred to future proposal

## Impact

**Affected specs**:
- `drawingml-core` - ADDED diagram element requirements
- `drawingml-diagrams` - NEW spec for SmartArt diagram structure
- `pdf-drawing` - ADDED (minimal) diagram placeholder rendering requirements

**New capabilities**:
- Read documents with SmartArt without errors
- Preserve SmartArt content on save (roundtrip)
- Parse diagram data model, layout, styles, colors
- Access diagram structure programmatically (read-only)

**Affected code**:
- `cmd/gen-go-drawingml/` - MODIFIED: Load diagram schemas
- `drawingml/diagram/` - NEW PACKAGE: Generated diagram elements
- `presentation/parts/` - ADDED: DiagramDataPart, DiagramLayoutDefinitionPart, DiagramStylePart, DiagramColorsPart
- `wordprocessing/parts/` - ADDED: Same diagram parts
- `spreadsheet/parts/` - ADDED: Same diagram parts
- `pdf/drawing/` - ADDED: Diagram placeholder renderer (renders bounding box only)

**Breaking changes**: None. This is additive functionality.

## Key Design Decisions

### 1. Two-Phase Approach
**Decision**: Split SmartArt support into Phase 1 (roundtrip) and Phase 2 (layout engine)

**Rationale**:
- Phase 1 (2-3 weeks) provides immediate value: preserving SmartArt content
- Phase 2 (6-12 weeks) requires significant investment in layout algorithms
- Separating concerns keeps proposals focused and reviewable
- Users can benefit from roundtrip support while layout engine is developed

### 2. Minimal Schema Coverage
**Decision**: Generate elements from all diagram schemas but only implement part infrastructure

**Rationale**:
- Element classes are auto-generated, so including all schemas is low cost
- Part infrastructure enables reading/writing diagram XML
- Layout algorithms are complex and warrant separate proposal
- This provides foundation for future SmartArt work

### 3. Diagram Part Organization
**Decision**: Create four separate part types (Data, Layout, Style, Colors)

**Rationale**:
- Matches Office Open XML structure exactly
- Each part serves distinct purpose:
  - Data: Logical node/connection structure
  - Layout: Algorithm for positioning nodes
  - Style: Visual styling rules
  - Colors: Color scheme transformations
- Keeps parts focused and maintainable

### 4. Schema Generation Strategy
**Decision**: Load all diagram schemas (2006, 2008, 2010, 2016) in generator

**Rationale**:
- Microsoft extended diagram schemas across multiple Office versions
- All versions may appear in real documents
- Generator already handles multi-version schemas (see chart schemas)
- Comprehensive coverage prevents parse failures

### 5. PDF Rendering Approach (Phase 1)
**Decision**: Render placeholder bounding box only, defer full rendering to Phase 2

**Rationale**:
- Full diagram rendering requires layout engine (Phase 2)
- Placeholder prevents blank space in PDF output
- Allows users to identify diagram locations
- Consistent with approach for other complex elements (e.g., SmartArt in Microsoft's own converters often shows simplified views)

## Implementation Scope

### Phase 1: Schema Generation & Roundtrip (2-3 weeks)

**Week 1: Code Generation**
- Update `cmd/gen-go-drawingml/main.go` to load diagram schemas
- Generate diagram element classes in `drawingml/diagram/` package
- Validate generated code compiles and passes golangci-lint

**Week 2: Part Infrastructure**
- Implement DiagramDataPart (data model XML)
- Implement DiagramLayoutDefinitionPart (layout definition XML)
- Implement DiagramStylePart (style definition XML)
- Implement DiagramColorsPart (color transform XML)
- Add part discovery to WordprocessingDocument, PresentationDocument, SpreadsheetDocument
- Add relationship types for diagram parts

**Week 3: Testing & Validation**
- Create test documents with SmartArt (PowerPoint, Word, Excel)
- Roundtrip tests: open → save → compare XML
- Validate diagram parts parse correctly
- Validate relationships preserved
- Integration tests with real Office documents

### Success Criteria (Phase 1)

- [ ] Diagram schemas loaded in generator
- [ ] Diagram elements generated in `drawingml/diagram/` package
- [ ] DiagramDataPart, DiagramLayoutDefinitionPart, DiagramStylePart, DiagramColorsPart implemented
- [ ] Can open PowerPoint/Word/Excel with SmartArt without errors
- [ ] SmartArt preserved on save (XML comparison shows no loss)
- [ ] Can parse data model (nodes, connections) from existing SmartArt
- [ ] All tests pass (unit, integration, roundtrip)
- [ ] Documentation complete (godoc, examples)

### Phase 2: Layout Engine (Future Proposal - 6-12 weeks)

**Out of scope for this proposal. Future work includes:**
- Layout algorithm implementations (list, hierarchy, cycle, etc.)
- SmartArt creation API (add nodes, set layout, apply styles)
- Full PDF rendering with positioned shapes
- Layout testing with visual validation

## Dependencies

**Existing**:
- DrawingML infrastructure (`drawingml/` package) - exists
- OpenXML element framework (`openxml/` package) - exists
- Part loading system (`packaging/` package) - exists
- Code generation framework (`cmd/gen-go-*/`) - exists

**New**: None. All required infrastructure exists.

## Risks & Mitigation

### Risk: Diagram Schemas Incompatible with Generator
**Impact**: High (blocks code generation)
**Probability**: Low (schemas follow same format as chart/main DrawingML)
**Mitigation**: Validate schemas early, adjust generator if needed (similar to chart schema handling)

### Risk: Roundtrip Failures Due to Complex Relationships
**Impact**: Medium (SmartArt not preserved correctly)
**Probability**: Medium (diagram parts have multiple interdependencies)
**Mitigation**:
- Comprehensive roundtrip tests with real Office documents
- Test with SmartArt from different Office versions (2010, 2013, 2016, 365)
- Compare saved XML byte-for-byte with original

### Risk: Part Type Conflicts
**Impact**: Low (compilation errors)
**Probability**: Low (part types are well-defined in spec)
**Mitigation**: Follow existing part patterns from chart parts

## Technical Challenges

### Challenge 1: Multi-Part Relationships
SmartArt uses 4 interconnected parts (data, layout, style, colors) with relationships between them.

**Solution**:
- Each part maintains relationships to others
- Follow pattern from chart parts (ChartPart has relationships to style/color parts)
- Lazy loading: parts loaded on first access

### Challenge 2: Schema Complexity
Diagram schemas include complex nested structures for layout algorithms.

**Solution**:
- Generator handles complexity automatically
- Phase 1 only requires parsing/roundtrip, not algorithm implementation
- Layout algorithm understanding deferred to Phase 2

### Challenge 3: Office Version Variations
Diagram schemas evolved across Office 2007, 2010, 2013, 2016, 2019, 365.

**Solution**:
- Load all diagram schema versions in generator
- Support all namespace variations (2006, 2008, 2010, 2016)
- Test with documents from each Office version

## References

### OpenXML-SDK References
- `DocumentFormat.OpenXml.Drawing.Diagrams` namespace
- DiagramData, DiagramLayout, DiagramColors, DiagramStyles element structures
- Point (dgm:pt), Connection (dgm:cxn) data model elements

### Diagram Schemas
Located in Open-XML-SDK `data/schemas/`:
- `schemas_openxmlformats_org_drawingml_2006_diagram.json` (core)
- `schemas_microsoft_com_office_drawing_2008_diagram.json` (Office 2010)
- `schemas_microsoft_com_office_drawing_2010_diagram.json` (Office 2013)
- `schemas_microsoft_com_office_drawing_2016_11_diagram.json` (Office 2016)
- `schemas_microsoft_com_office_drawing_2016_12_diagram.json` (Office 2019/365)

### ECMA-376 References
- Part 1, Section 21.4: Diagrams
- Part 4: Transitional Migration Features

## Files to Create/Modify

### Code Generation
1. `cmd/gen-go-drawingml/main.go` - MODIFY: Load diagram schemas
2. `cmd/gen-go-drawingml/loaders.go` - MODIFY: Add diagram schema loader functions
3. `drawingml/diagram/` - NEW PACKAGE: Generated element files

### Part Infrastructure
4. `presentation/parts/diagram_part.go` - NEW: DiagramDataPart, DiagramLayoutDefinitionPart, DiagramStylePart, DiagramColorsPart
5. `wordprocessing/parts/diagram_part.go` - NEW: Same parts for Word
6. `spreadsheet/parts/diagram_part.go` - NEW: Same parts for Excel
7. `packaging/relationship_types.go` - MODIFY: Add diagram relationship type constants

### PDF Rendering (Minimal)
8. `pdf/drawing/diagram_renderer.go` - NEW: Placeholder bounding box renderer

### Testing
9. `drawingml/diagram/diagram_test.go` - NEW: Element parsing tests
10. `presentation/parts/diagram_part_test.go` - NEW: Part roundtrip tests
11. `testdata/smartart/` - NEW: Test documents with various SmartArt types

### Documentation
12. `drawingml/diagram/doc.go` - NEW: Package documentation
13. `CHANGELOG.md` - MODIFY: Document SmartArt roundtrip support

## Estimated Effort

**Phase 1 (This Proposal): 2-3 weeks**
- Schema loading: 3-4 days
- Element generation & validation: 4-5 days
- Part implementation: 4-5 days
- Testing & roundtrip validation: 3-4 days
- Documentation: 1-2 days

**Phase 2 (Future): 6-12 weeks**
- Simple layout implementation (list): 2 weeks
- Org chart layout: 3 weeks
- Cycle/process layouts: 3 weeks
- Testing & fidelity validation: 2-4 weeks

## Future Enhancements (Not in Scope)

These will be addressed in separate proposals:
1. SmartArt layout engine (Phase 2 proposal)
2. SmartArt creation API
3. SmartArt PDF rendering (full)
4. SmartArt style customization
5. Custom layout definitions
6. SmartArt animation (presentation-specific)

## Migration Path

**For Users**:
- After Phase 1: Documents with SmartArt can be safely opened and saved
- After Phase 2: Programmatic SmartArt creation becomes possible

**For Developers**:
- Phase 1 provides read/write infrastructure
- Phase 2 builds on Phase 1 foundation
- No breaking changes between phases
