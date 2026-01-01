# Implementation Tasks

## Phase 1: DrawingML Table Elements (Week 1)

### 1.1 Package Structure and Core Types
- [ ] 1.1.1 Create `drawingml/table/` package directory
- [ ] 1.1.2 Create `drawingml/table/table.go` with Table struct
- [ ] 1.1.3 Implement NewTable() constructor
- [ ] 1.1.4 Add XML namespace registration
- [ ] 1.1.5 Implement Clone() method for Table
- [ ] 1.1.6 Add unit tests for Table creation

### 1.2 TableProperties
- [ ] 1.2.1 Create TableProperties struct in `table.go`
- [ ] 1.2.2 Add boolean attributes (FirstRow, LastRow, BandRow, BandCol)
- [ ] 1.2.3 Create TableStyleId struct for GUID references
- [ ] 1.2.4 Implement NewTableProperties() constructor
- [ ] 1.2.5 Add XML serialization tests
- [ ] 1.2.6 Add validation logic

### 1.3 TableGrid and GridColumn
- [ ] 1.3.1 Create `drawingml/table/grid.go`
- [ ] 1.3.2 Implement TableGrid struct with Columns slice
- [ ] 1.3.3 Implement GridColumn struct with Width attribute (EMUs)
- [ ] 1.3.4 Implement NewTableGrid() and NewGridColumn() constructors
- [ ] 1.3.5 Add grid serialization tests
- [ ] 1.3.6 Test column width calculations

### 1.4 TableRow
- [ ] 1.4.1 Create `drawingml/table/row.go`
- [ ] 1.4.2 Implement TableRow struct with Height and Cells
- [ ] 1.4.3 Implement NewTableRow() constructor
- [ ] 1.4.4 Add row height management
- [ ] 1.4.5 Add unit tests for row operations

### 1.5 TableCell and CellProperties
- [ ] 1.5.1 Create `drawingml/table/cell.go`
- [ ] 1.5.2 Implement TableCell struct (TextBody, CellProperties, merge attributes)
- [ ] 1.5.3 Implement CellProperties struct (margins, anchor, fill, borders)
- [ ] 1.5.4 Implement CellBorder struct for border lines
- [ ] 1.5.5 Implement NewTableCell() and NewCellProperties() constructors
- [ ] 1.5.6 Add cell serialization tests

### 1.6 Enums and Constants
- [ ] 1.6.1 Create `drawingml/table/enums.go`
- [ ] 1.6.2 Define TextAnchoringType enum (top, center, bottom)
- [ ] 1.6.3 Define TextHorzOverflowType enum (clip, overflow)
- [ ] 1.6.4 Add common table style GUID constants
- [ ] 1.6.5 Document enum values

### 1.7 Validation
- [ ] 1.7.1 Implement Validate() for Table
- [ ] 1.7.2 Implement Validate() for TableRow (ensure all rows have same cell count)
- [ ] 1.7.3 Implement Validate() for TableCell
- [ ] 1.7.4 Add validation tests for invalid structures
- [ ] 1.7.5 Test with malformed XML

### 1.8 XML Serialization
- [ ] 1.8.1 Test Table serialization to XML
- [ ] 1.8.2 Test Table deserialization from XML
- [ ] 1.8.3 Roundtrip test (serialize → deserialize → compare)
- [ ] 1.8.4 Compare output with real PowerPoint table XML
- [ ] 1.8.5 Verify namespace handling

## Phase 2: High-Level Table API (Week 2)

### 2.1 Table Wrapper Type
- [ ] 2.1.1 Create `presentation/table.go`
- [ ] 2.1.2 Define Table struct (graphicFrame, table, slide fields)
- [ ] 2.1.3 Implement helper to extract table from GraphicFrame
- [ ] 2.1.4 Add doc comments for Table type

### 2.2 Table Creation
- [ ] 2.2.1 Implement Slide.AddTable(rows, cols) method
- [ ] 2.2.2 Create LinkGraphicFrameToTable() helper in `elements/helpers.go`
- [ ] 2.2.3 Implement default table initialization (grid, rows, cells)
- [ ] 2.2.4 Set reasonable default dimensions and position
- [ ] 2.2.5 Add validation for minimum rows/cols (must be >= 1)
- [ ] 2.2.6 Add unit tests for table creation

### 2.3 Row Access
- [ ] 2.3.1 Implement RowCount() method
- [ ] 2.3.2 Implement ColumnCount() method
- [ ] 2.3.3 Implement GetRow(index) method with bounds checking
- [ ] 2.3.4 Add row iterator method Rows()
- [ ] 2.3.5 Add unit tests for row access

### 2.4 Row Manipulation
- [ ] 2.4.1 Implement AppendRow() method
- [ ] 2.4.2 Implement InsertRow(index) method
- [ ] 2.4.3 Implement DeleteRow(index) method
- [ ] 2.4.4 Ensure cells are created matching column count
- [ ] 2.4.5 Add unit tests for row manipulation
- [ ] 2.4.6 Test edge cases (insert at 0, delete last row, etc.)

### 2.5 Cell Access
- [ ] 2.5.1 Implement GetCell(row, col) method with bounds checking
- [ ] 2.5.2 Implement SetCellText(row, col, text) method
- [ ] 2.5.3 Implement GetCellText(row, col) method
- [ ] 2.5.4 Handle TextBody creation if needed
- [ ] 2.5.5 Add unit tests for cell access and content

### 2.6 Column Width Management
- [ ] 2.6.1 Implement SetColumnWidth(col, widthPt) method
- [ ] 2.6.2 Implement GetColumnWidth(col) method
- [ ] 2.6.3 Implement SetEqualColumnWidths() method
- [ ] 2.6.4 Add EMU ↔ points conversion
- [ ] 2.6.5 Add unit tests for column widths

### 2.7 Position and Size
- [ ] 2.7.1 Implement SetPosition(xPt, yPt) method
- [ ] 2.7.2 Implement GetPosition() method
- [ ] 2.7.3 Implement SetSize(widthPt, heightPt) method
- [ ] 2.7.4 Implement GetSize() method
- [ ] 2.7.5 Update column widths when size changes
- [ ] 2.7.6 Add unit tests for positioning

## Phase 3: Table Styling (Week 3)

### 3.1 Table-Level Styling
- [ ] 3.1.1 Implement SetStyle(styleGUID) method
- [ ] 3.1.2 Define common table style GUID constants
- [ ] 3.1.3 Implement SetBandedRows(enabled) method
- [ ] 3.1.4 Implement SetBandedColumns(enabled) method
- [ ] 3.1.5 Implement SetFirstRowBold(enabled) method
- [ ] 3.1.6 Implement SetFirstColumnBold(enabled) method
- [ ] 3.1.7 Implement SetLastRowBold(enabled) method
- [ ] 3.1.8 Implement SetLastColumnBold(enabled) method
- [ ] 3.1.9 Add unit tests for styling methods

### 3.2 Cell-Level Formatting
- [ ] 3.2.1 Implement SetCellBackgroundColor(row, col, r, g, b) method
- [ ] 3.2.2 Implement SetCellAlignment(row, col, align) method
- [ ] 3.2.3 Implement SetCellMargins(row, col, left, right, top, bottom) method
- [ ] 3.2.4 Add unit tests for cell formatting

### 3.3 Cell Borders (Basic)
- [ ] 3.3.1 Implement SetCellBorder(row, col, side, width, r, g, b) method
- [ ] 3.3.2 Support border sides: left, right, top, bottom
- [ ] 3.3.3 Add unit tests for borders

### 3.4 Integration Tests with Styles
- [ ] 3.4.1 Create table with Medium2 style
- [ ] 3.4.2 Create table with Light1 style
- [ ] 3.4.3 Test banded rows rendering
- [ ] 3.4.4 Test first row formatting
- [ ] 3.4.5 Compare with PowerPoint output

## Phase 4: PDF Rendering (Week 4)

### 4.1 Table Renderer Setup
- [ ] 4.1.1 Create `pdf/presentation/table_renderer.go`
- [ ] 4.1.2 Define TableRenderer struct
- [ ] 4.1.3 Implement RenderTable(ctx, table) entry point
- [ ] 4.1.4 Calculate cell dimensions from column widths and row heights

### 4.2 Grid Rendering
- [ ] 4.2.1 Implement renderTableGrid() for borders
- [ ] 4.2.2 Draw horizontal lines for rows
- [ ] 4.2.3 Draw vertical lines for columns
- [ ] 4.2.4 Use column widths from table grid
- [ ] 4.2.5 Add unit tests for grid rendering

### 4.3 Cell Background Rendering
- [ ] 4.3.1 Implement renderCellBackground() method
- [ ] 4.3.2 Extract fill color from CellProperties
- [ ] 4.3.3 Draw filled rectangle for cell
- [ ] 4.3.4 Handle cells without background (skip)
- [ ] 4.3.5 Test with various background colors

### 4.4 Cell Text Rendering
- [ ] 4.4.1 Implement renderCellText() method
- [ ] 4.4.2 Extract text from cell's TextBody
- [ ] 4.4.3 Apply vertical alignment (top, center, bottom)
- [ ] 4.4.4 Add cell padding to text position
- [ ] 4.4.5 Handle multi-line text (wrap if needed)
- [ ] 4.4.6 Test with long text content

### 4.5 Style-Based Rendering
- [ ] 4.5.1 Implement banded row coloring in PDF
- [ ] 4.5.2 Implement first row bold rendering
- [ ] 4.5.3 Implement last row bold rendering
- [ ] 4.5.4 Test visual appearance matches expectations

### 4.6 PDF Integration Tests
- [ ] 4.6.1 Render simple 3x3 table to PDF
- [ ] 4.6.2 Render table with styled headers to PDF
- [ ] 4.6.3 Render table with banded rows to PDF
- [ ] 4.6.4 Render table with custom backgrounds to PDF
- [ ] 4.6.5 Visual comparison with expected output

## Phase 5: Testing and Documentation (Week 4)

### 5.1 Unit Tests
- [ ] 5.1.1 Complete test coverage for DrawingML table elements
- [ ] 5.1.2 Complete test coverage for Table API methods
- [ ] 5.1.3 Edge case tests (empty tables, single cell, etc.)
- [ ] 5.1.4 Error handling tests (invalid indices, etc.)
- [ ] 5.1.5 Achieve >80% code coverage

### 5.2 Integration Tests
- [ ] 5.2.1 Create table → save → load → verify structure
- [ ] 5.2.2 Create table → save → open in PowerPoint → verify
- [ ] 5.2.3 Load real PowerPoint file → extract table → verify
- [ ] 5.2.4 Modify table → save → verify changes persist
- [ ] 5.2.5 Test with PowerPoint 2010, 2013, 2016, 2019, 365

### 5.3 Roundtrip Tests
- [ ] 5.3.1 Roundtrip: create table with data
- [ ] 5.3.2 Roundtrip: create table with styles
- [ ] 5.3.3 Roundtrip: create table with cell formatting
- [ ] 5.3.4 Roundtrip: load PowerPoint table → modify → save
- [ ] 5.3.5 Verify XML structure matches PowerPoint output

### 5.4 Examples
- [ ] 5.4.1 Create `examples/presentation/table/simple/main.go` - basic table
- [ ] 5.4.2 Create `examples/presentation/table/styled/main.go` - styled table
- [ ] 5.4.3 Create `examples/presentation/table/data/main.go` - data table with calculations
- [ ] 5.4.4 Create `examples/presentation/table/advanced/main.go` - advanced formatting
- [ ] 5.4.5 Verify all examples run without errors

### 5.5 Documentation
- [ ] 5.5.1 Write godoc for all public types and methods
- [ ] 5.5.2 Create `drawingml/table/doc.go` package documentation
- [ ] 5.5.3 Create `presentation/table_doc.go` API documentation
- [ ] 5.5.4 Document table style GUIDs and their visual appearance
- [ ] 5.5.5 Document limitations (what's NOT supported)

### 5.6 Performance Testing
- [ ] 5.6.1 Benchmark table creation (10x10, 50x50, 100x100)
- [ ] 5.6.2 Benchmark cell text manipulation (1000 operations)
- [ ] 5.6.3 Benchmark XML serialization (large tables)
- [ ] 5.6.4 Benchmark PDF rendering (various table sizes)
- [ ] 5.6.5 Ensure operations complete in reasonable time

### 5.7 Compatibility Testing
- [ ] 5.7.1 Test with PowerPoint 2010
- [ ] 5.7.2 Test with PowerPoint 2013
- [ ] 5.7.3 Test with PowerPoint 2016
- [ ] 5.7.4 Test with PowerPoint 2019
- [ ] 5.7.5 Test with Office 365
- [ ] 5.7.6 Test with LibreOffice Impress
- [ ] 5.7.7 Document known compatibility issues

### 5.8 Final Integration
- [ ] 5.8.1 Update `CHANGELOG.md` with table API addition
- [ ] 5.8.2 Run golangci-lint and fix all issues
- [ ] 5.8.3 Ensure all existing tests still pass (no regressions)
- [ ] 5.8.4 Verify no breaking changes to existing API
- [ ] 5.8.5 Update example_test.go with working table example

## Phase 6: Advanced Features (Optional - Future)

These are deferred to future enhancements but documented here for reference:

- [ ] 6.1 Cell Merge/Split Operations
  - [ ] 6.1.1 Implement MergeCells(startRow, startCol, endRow, endCol)
  - [ ] 6.1.2 Implement SplitCell(row, col)
  - [ ] 6.1.3 Handle GridSpan and RowSpan attributes
  - [ ] 6.1.4 Update rendering to handle merged cells

- [ ] 6.2 Custom Table Styles
  - [ ] 6.2.1 Create TableStyle struct for custom styles
  - [ ] 6.2.2 Implement style definition XML
  - [ ] 6.2.3 Add custom style to presentation
  - [ ] 6.2.4 Reference custom style from table

- [ ] 6.3 Data Binding
  - [ ] 6.3.1 PopulateFromSlice(data [][]string) method
  - [ ] 6.3.2 PopulateFromCSV(reader io.Reader) method
  - [ ] 6.3.3 PopulateFromStruct(data interface{}) method

- [ ] 6.4 Advanced Formatting
  - [ ] 6.4.1 Gradient fills for cells
  - [ ] 6.4.2 Images in cells
  - [ ] 6.4.3 Diagonal borders
  - [ ] 6.4.4 Custom border patterns

## Notes

- **Dependencies**: Phase 1 must complete before Phase 2 can begin
- **Parallelization**: Phase 3 (styling) and Phase 4 (PDF) can be done in parallel after Phase 2
- **Testing**: Write tests concurrently with implementation, not after
- **Documentation**: Write godoc as you implement each method
- **Validation**: Run `spectr validate add-presentation-table-api` before marking complete
