# Change: Add PowerPoint Table High-Level API

## Why

goffice currently lacks a high-level API for creating and manipulating tables in PowerPoint presentations. While the GraphicFrame infrastructure exists (30% complete), DrawingML table elements are not generated and there is no table manipulation API. This prevents users from programmatically creating tables in presentations.

**Current State**:
- GraphicFrame elements exist with positioning/sizing support
- Table placeholder type defined but not functional
- Example code shows stub implementation with "TODO" comment
- DrawingML table elements (a:tbl, a:tr, a:tc) are not generated
- No high-level Table API (unlike Word which has comprehensive table support)

**Impact Without This Change**:
- **PowerPoint**: Cannot create data tables, comparison matrices, or structured content
- **User Experience**: Must manually edit XML to add tables to presentations
- **Feature Parity**: Word has table API, but PowerPoint does not
- **Office Compatibility**: Cannot programmatically create tables that match PowerPoint's native tables

**Evidence from Analysis**:
- GraphicFrame infrastructure exists in `presentation/elements/shape_tree.go`
- DrawingML table elements missing from code generation
- Example_tableCreation() shows stub implementation with comment: "In a full implementation, you would set the graphic data to a table type and add rows/columns."
- Word table API at `wordprocessing/elements/table.go` (1,636 lines) provides reference pattern
- Test PPTX files demonstrate expected XML structure

## What Changes

**Scope: Complete PowerPoint Table API**

This proposal adds full table support to PowerPoint presentations, matching the feature set of Word's table API.

### 1. DrawingML Table Elements

Generate or create table-specific DrawingML elements:
- `Table` (a:tbl) - main table container
- `TableProperties` (a:tblPr) - table-level styling (banded rows, first row, etc.)
- `TableGrid` (a:tblGrid) - column definitions
- `GridColumn` (a:gridCol) - individual column with width
- `TableRow` (a:tr) - row container with height
- `TableCell` (a:tc) - cell with text body and properties
- `CellProperties` (a:tcPr) - cell-level formatting (borders, fill, margins)
- `TableStyleId` - style reference

### 2. High-Level Table API

Create presentation-layer table API following Word's pattern:

```go
// Create table
table := slide.AddTable(rows, cols)

// Manipulate structure
table.AppendRow()
table.InsertRow(index)
table.DeleteRow(index)
table.SetColumnWidth(col, width)

// Populate content
table.SetCellText(row, col, "Header 1")
cell := table.GetCell(row, col)
cell.SetText("Content")

// Apply styling
table.SetStyle("{5C22544A-7EE6-4342-B048-85BDC9FD1C3A}")
table.SetBandedRows(true)
table.SetFirstRowBold(true)

// Position and size
table.SetPosition(100, 100)  // in points
table.SetSize(400, 300)      // in points
```

### 3. Helper Functions

Add table-specific helper functions:
- `LinkGraphicFrameToTable()` - similar to existing `LinkGraphicFrameToChart()`
- `NewTable(rows, cols)` - table factory
- Cell content population helpers

**Breaking Changes**: None. This is purely additive functionality.

## Impact

**Affected specs**:
- `presentation-elements` - ADDED table creation and manipulation requirements
- `drawingml-tables` - NEW spec for DrawingML table elements
- `pdf-presentation` - ADDED table rendering to PDF requirements

**New capabilities**:
- Create tables programmatically in PowerPoint slides
- Manipulate table structure (add/remove rows/columns)
- Set cell content and formatting
- Apply table styles
- Position and size tables on slides
- Render tables in PDF output

**Affected code**:
- `drawingml/table/` - NEW PACKAGE: DrawingML table elements
- `presentation/table.go` - NEW: High-level Table API
- `presentation/elements/helpers.go` - MODIFIED: Add LinkGraphicFrameToTable()
- `presentation/slide.go` - MODIFIED: Add AddTable() method
- `pdf/presentation/table_renderer.go` - NEW: Table rendering for PDF
- `cmd/gen-go-drawingml/` - MODIFIED: Add table schema loading

## Key Design Decisions

### 1. Follow Word Table API Pattern
**Decision**: Model PowerPoint table API after existing Word table API

**Rationale**:
- Word table API is mature and well-tested (1,636 lines)
- Consistent API across document types (Word, PowerPoint)
- Users familiar with Word tables can easily use PowerPoint tables
- Proven design patterns for row/column/cell manipulation

**API Consistency**:
```go
// Word
wordTable := doc.AddTable(3, 4)
wordTable.SetCellText(0, 0, "Header")

// PowerPoint (proposed)
pptTable := slide.AddTable(3, 4)
pptTable.SetCellText(0, 0, "Header")
```

### 2. Generate vs. Manual Element Creation
**Decision**: Manually create DrawingML table element types (do not generate)

**Rationale**:
- Table elements are in DrawingML namespace (a:), not PresentationML (p:)
- Code generator focuses on PresentationML schemas currently
- Manual implementation faster for this specific use case
- Only ~15 types needed (vs. hundreds in full schema)
- Can migrate to generated code in future if needed

**Alternatives Considered**:
- Extend code generator for DrawingML tables: Too much effort for 15 types
- Use generic XML elements: Loses type safety and validation

### 3. Table Style Management
**Decision**: Support built-in Office table styles via GUID references

**Rationale**:
- PowerPoint uses GUID-based style references (e.g., {5C22544A-7EE6-4342-B048-85BDC9FD1C3A})
- Matches Office behavior exactly
- Allows users to reference standard Office table styles
- Custom styles can be added in future enhancement

### 4. Cell Content Model
**Decision**: Use TextBody for cell content (reuse existing DrawingML text infrastructure)

**Rationale**:
- Table cells in PowerPoint use a:txBody (TextBody) element
- DrawingML TextBody infrastructure already exists and is mature
- Supports rich text formatting (paragraphs, runs, fonts, colors)
- Consistent with other PowerPoint text handling

### 5. Measurement Units
**Decision**: Table API uses points for dimensions, converts to EMUs internally

**Rationale**:
- Points are user-friendly (standard typography unit)
- Consistent with other goffice APIs (PDF uses points)
- GraphicFrame already has EMU conversion in SetPosition/SetSize
- Users don't need to think about EMUs (internal Office unit)

**Example**:
```go
table.SetSize(400, 300)  // 400pt width, 300pt height
table.SetColumnWidth(0, 100)  // 100pt wide column
```

## Implementation Scope

**Timeline: 3-4 weeks**

### Week 1: DrawingML Table Elements
- Create `drawingml/table/` package
- Implement Table, TableProperties, TableGrid types
- Implement TableRow, TableCell, CellProperties types
- Add validation and XML serialization
- Unit tests for element creation and serialization

### Week 2: High-Level Table API
- Create `presentation/table.go` with Table wrapper
- Implement table creation (AddTable on Slide)
- Implement row manipulation (AppendRow, InsertRow, DeleteRow)
- Implement column manipulation (SetColumnWidth)
- Implement cell access (GetCell, SetCellText)
- Add LinkGraphicFrameToTable() helper
- Unit tests for API methods

### Week 3: Styling and Advanced Features
- Implement table style application (SetStyle)
- Implement banded rows/columns
- Implement first row/column formatting
- Cell-level formatting (borders, fill, alignment)
- Integration tests with real PPTX files

### Week 4: PDF Rendering and Documentation
- Implement table rendering for PDF
- Roundtrip tests (create → save → load → verify)
- Example programs
- Documentation and godoc
- Compatibility testing with PowerPoint

### Success Criteria

- [ ] DrawingML table elements implemented and validated
- [ ] Table API supports create, read, update, delete operations
- [ ] Can create table with N rows and M columns
- [ ] Can add/insert/delete rows
- [ ] Can set column widths
- [ ] Can get/set cell content (text)
- [ ] Can apply table styles
- [ ] Tables render correctly in PDF
- [ ] Roundtrip preserves table structure and content
- [ ] Tables open correctly in PowerPoint
- [ ] All tests pass (unit, integration, roundtrip)
- [ ] Documentation complete with examples

## Out of Scope

**Deferred to Future Enhancements**:
- Custom table style definitions (use built-in styles for now)
- Cell merge/split operations
- Advanced cell formatting (gradients, images in cells)
- Table templates
- Data binding (populate table from data source)
- Cell formulas (Excel-like calculations)
- Conditional formatting
- Table of contents auto-generation
- Advanced borders (diagonal, custom patterns)

**Not Planned**:
- Excel-style formulas in tables
- Real-time table collaboration
- Table animation sequences

## Dependencies

**Existing Infrastructure**:
- GraphicFrame support (`presentation/elements/shape_tree.go`)
- DrawingML TextBody for cell content
- Slide API for adding content
- XML serialization framework

**New Dependencies**: None. All required infrastructure exists.

## Risks & Mitigation

### Risk: Style GUID Mismatch
**Impact**: Medium (tables render with wrong styles)
**Probability**: Low (GUIDs are standardized)
**Mitigation**:
- Test with multiple Office versions (2010, 2013, 2016, 2019, 365)
- Document supported style GUIDs
- Validate against real PowerPoint output

### Risk: Cell Content Rendering Complexity
**Impact**: Medium (text doesn't fit in cells)
**Probability**: Medium (text layout is non-trivial)
**Mitigation**:
- Reuse existing TextBody rendering from PDF/presentation
- Add cell padding to prevent overflow
- Test with various text lengths

### Risk: Column Width Calculation
**Impact**: Low (columns not evenly sized)
**Probability**: Low (simple proportional calculation)
**Mitigation**:
- Default to equal column widths
- Allow explicit width setting
- Validate total width matches table width

### Risk: Office Compatibility
**Impact**: High (tables don't open in PowerPoint)
**Probability**: Low (following official XML structure)
**Mitigation**:
- Test with real PPTX files from PowerPoint
- Validate XML structure against Office schema
- Roundtrip testing with multiple Office versions

## References

### Office Open XML Spec
- ECMA-376 Part 1, Section 21.1.3: DrawingML - Table
- Table element structure (a:tbl)
- TableProperties, TableGrid, TableRow, TableCell

### Existing goffice Code
- `wordprocessing/elements/table.go` - Reference implementation for API pattern
- `presentation/elements/shape_tree.go` - GraphicFrame infrastructure
- `presentation/elements/helpers.go` - LinkGraphicFrameToChart pattern
- `drawingml/` - DrawingML element patterns

### Test Files
- `Open-XML-SDK/test/.../presentation/Table_Small.pptx` - Real PowerPoint table structure
- Shows complete XML hierarchy for tables

## Files to Create/Modify

### DrawingML Table Elements
1. `drawingml/table/table.go` - NEW: Table, TableProperties, TableStyleId types
2. `drawingml/table/grid.go` - NEW: TableGrid, GridColumn types
3. `drawingml/table/row.go` - NEW: TableRow type
4. `drawingml/table/cell.go` - NEW: TableCell, CellProperties types
5. `drawingml/table/enums.go` - NEW: Table-related enums (alignment, border styles)

### High-Level API
6. `presentation/table.go` - NEW: Table wrapper with high-level methods
7. `presentation/slide.go` - MODIFIED: Add AddTable() method
8. `presentation/elements/helpers.go` - MODIFIED: Add LinkGraphicFrameToTable()

### PDF Rendering
9. `pdf/presentation/table_renderer.go` - NEW: Table rendering for PDF

### Testing
10. `presentation/table_test.go` - NEW: Table API unit tests
11. `presentation/table_integration_test.go` - NEW: Integration tests with PPTX files
12. `drawingml/table/table_test.go` - NEW: Element tests
13. `testdata/presentation/tables/` - NEW: Test PPTX files with tables

### Documentation
14. `drawingml/table/doc.go` - NEW: Package documentation
15. `examples/presentation/table/` - NEW: Example programs
16. `CHANGELOG.md` - MODIFIED: Document table API support

## Estimated Effort

**Total: 3-4 weeks (120-160 hours)**

- DrawingML elements: 1 week (40 hours)
  - Type definitions: 2 days
  - XML serialization: 2 days
  - Validation and tests: 1 day

- High-level API: 1 week (40 hours)
  - Table wrapper: 2 days
  - Row/column manipulation: 2 days
  - Cell access and content: 1 day

- Styling and advanced features: 1 week (40 hours)
  - Style application: 2 days
  - Banded rows/columns: 1 day
  - Cell formatting: 2 days

- Testing, PDF, documentation: 1 week (40 hours)
  - PDF rendering: 2 days
  - Roundtrip and integration tests: 2 days
  - Examples and documentation: 1 day

## Migration Path

**For Users**:
- New table API is immediately available after implementation
- No migration needed (purely additive)
- Existing code continues to work unchanged

**For Developers**:
- Table API follows Word table pattern (easy to learn if familiar with Word)
- Examples demonstrate common use cases
- Documentation covers all API methods

## Future Enhancements (Not in Scope)

These will be addressed in separate proposals:
1. Cell merge/split operations
2. Custom table style definitions
3. Advanced cell formatting (gradients, images)
4. Table templates and themes
5. Data binding and auto-population
6. Conditional formatting
7. Advanced border customization
