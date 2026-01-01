# DrawingML Tables Specification

## Requirements

### Requirement: Table Element Structure
The system SHALL provide DrawingML table element types following the ECMA-376 specification.

#### Scenario: Create table element
- WHEN NewTable() is called
- THEN a Table element is created
- AND Table has TableProperties, TableGrid, and Rows fields
- AND XML namespace is "http://schemas.openxmlformats.org/drawingml/2006/main"
- AND XML element name is "tbl"

#### Scenario: Table XML serialization
- GIVEN a Table with 2 rows and 3 columns
- WHEN table is serialized to XML
- THEN XML contains <a:tbl> root element
- AND <a:tblPr>, <a:tblGrid>, and <a:tr> child elements are present

#### Scenario: Table XML deserialization
- GIVEN valid table XML from PowerPoint
- WHEN XML is deserialized
- THEN Table object is created with correct structure
- AND all properties and cells are populated

### Requirement: Table Properties
The system SHALL provide table-level properties for styling and formatting.

#### Scenario: Create table properties
- WHEN NewTableProperties() is called
- THEN TableProperties element is created with default values

#### Scenario: Set first row formatting
- GIVEN TableProperties
- WHEN FirstRow attribute is set to true
- THEN first row will receive special header formatting

#### Scenario: Set banded rows
- GIVEN TableProperties
- WHEN BandRow attribute is set to true
- THEN alternating row background colors are enabled

#### Scenario: Set table style ID
- GIVEN TableProperties
- WHEN TableStyleId is set to "{5C22544A-7EE6-4342-B048-85BDC9FD1C3A}"
- THEN table references the Medium 2 style

#### Scenario: Table properties XML attributes
- GIVEN TableProperties with firstRow=true, bandRow=true
- WHEN serialized to XML
- THEN <a:tblPr firstRow="1" bandRow="1"> is generated

### Requirement: Table Grid and Columns
The system SHALL provide table grid for defining column structure.

#### Scenario: Create table grid
- WHEN NewTableGrid() is called
- THEN TableGrid element is created with empty columns slice

#### Scenario: Add grid column
- GIVEN a TableGrid
- WHEN a GridColumn with width 914400 EMUs is added
- THEN column is added to grid
- AND column width is stored in EMUs

#### Scenario: Grid column XML
- GIVEN a GridColumn with width 914400 (1 inch in EMUs)
- WHEN serialized to XML
- THEN <a:gridCol w="914400"/> is generated

#### Scenario: Multiple columns in grid
- GIVEN a TableGrid with 4 columns (widths: 100pt, 150pt, 200pt, 100pt)
- WHEN serialized
- THEN 4 <a:gridCol> elements are generated in order
- AND each has correct EMU width

### Requirement: Table Rows
The system SHALL provide table row elements with height and cells.

#### Scenario: Create table row
- WHEN NewTableRow() is called
- THEN TableRow element is created
- AND Cells slice is empty
- AND Height is nil (auto-height)

#### Scenario: Set row height
- GIVEN a TableRow
- WHEN Height is set to 457200 EMUs (0.5 inch)
- THEN row height is stored in EMUs

#### Scenario: Add cells to row
- GIVEN a TableRow
- WHEN 3 TableCells are added to Cells slice
- THEN row contains 3 cells

#### Scenario: Row XML serialization
- GIVEN a TableRow with height 457200 and 2 cells
- WHEN serialized
- THEN <a:tr h="457200"> contains 2 <a:tc> child elements

### Requirement: Table Cells
The system SHALL provide table cell elements with content and properties.

#### Scenario: Create table cell
- WHEN NewTableCell() is called
- THEN TableCell element is created
- AND TextBody is nil (empty cell)
- AND CellProperties is nil (default formatting)

#### Scenario: Set cell text content
- GIVEN a TableCell
- WHEN TextBody is set with DrawingML TextBody element
- THEN cell contains text content

#### Scenario: Set cell properties
- GIVEN a TableCell
- WHEN CellProperties is set
- THEN cell has custom formatting (margins, fill, borders)

#### Scenario: Cell XML serialization
- GIVEN a TableCell with TextBody and CellProperties
- WHEN serialized
- THEN <a:tc> contains <a:txBody> and <a:tcPr> child elements

### Requirement: Cell Properties
The system SHALL provide cell-level formatting properties.

#### Scenario: Create cell properties
- WHEN NewCellProperties() is called
- THEN CellProperties element is created with default values

#### Scenario: Set cell margins
- GIVEN CellProperties
- WHEN margins are set (left:45720, right:45720, top:45720, bottom:45720) in EMUs
- THEN margin attributes are set

#### Scenario: Set cell fill
- GIVEN CellProperties
- WHEN SolidFill is set to RGB(255, 200, 100)
- THEN cell background color is defined

#### Scenario: Set cell vertical alignment
- GIVEN CellProperties
- WHEN Anchor is set to TextAnchorCenter
- THEN cell content is vertically centered

#### Scenario: Cell properties XML
- GIVEN CellProperties with margins and fill
- WHEN serialized
- THEN <a:tcPr marL="45720" marR="45720" marT="45720" marB="45720" anchor="ctr"> is generated
- AND <a:solidFill> child element is present

### Requirement: Cell Borders
The system SHALL provide cell border formatting.

#### Scenario: Set cell left border
- GIVEN CellProperties
- WHEN BorderLeft is set with width 12700 (1pt) and color black
- THEN left border is defined

#### Scenario: Border XML serialization
- GIVEN a CellBorder with width and color
- WHEN serialized
- THEN <a:lnL w="12700"><a:solidFill>...</a:solidFill></a:lnL> is generated

#### Scenario: Multiple borders per cell
- GIVEN CellProperties with all 4 borders set
- WHEN serialized
- THEN <a:lnL>, <a:lnR>, <a:lnT>, <a:lnB> elements are all present

### Requirement: Table Validation
The system SHALL validate table structure to ensure correctness.

#### Scenario: Validate table with rows
- GIVEN a Table with 3 rows
- WHEN Validate() is called
- THEN no error is returned

#### Scenario: Validate table grid matches columns
- GIVEN a Table with TableGrid defining 4 columns
- AND all rows have 4 cells
- WHEN Validate() is called
- THEN validation passes

#### Scenario: Validate inconsistent column count
- GIVEN a Table where row 0 has 3 cells and row 1 has 4 cells
- WHEN Validate() is called
- THEN an error is returned indicating inconsistent column count

#### Scenario: Validate empty table
- GIVEN a Table with no rows
- WHEN Validate() is called
- THEN no error is returned (empty table is valid)

### Requirement: Table Element Cloning
The system SHALL support deep cloning of table elements.

#### Scenario: Clone table
- GIVEN a Table with data and formatting
- WHEN Clone() is called
- THEN a new Table instance is created
- AND all child elements are deep copied
- AND modifying clone does not affect original

#### Scenario: Clone row
- GIVEN a TableRow with cells
- WHEN Clone() is called
- THEN new TableRow is created with cloned cells

#### Scenario: Clone cell
- GIVEN a TableCell with TextBody and properties
- WHEN Clone() is called
- THEN new TableCell is created with cloned content

### Requirement: Text Anchoring Types
The system SHALL provide text vertical alignment enums.

#### Scenario: Text anchor top
- WHEN TextAnchorTop constant is used
- THEN value is "t"
- AND cell content aligns to top

#### Scenario: Text anchor center
- WHEN TextAnchorCenter constant is used
- THEN value is "ctr"
- AND cell content is vertically centered

#### Scenario: Text anchor bottom
- WHEN TextAnchorBottom constant is used
- THEN value is "b"
- AND cell content aligns to bottom

### Requirement: Text Overflow Types
The system SHALL provide text horizontal overflow behavior enums.

#### Scenario: Text overflow clip
- WHEN TextOverflowClip constant is used
- THEN value is "clip"
- AND text extending beyond cell bounds is clipped

#### Scenario: Text overflow overflow
- WHEN TextOverflowOverflow constant is used
- THEN value is "overflow"
- AND text can extend beyond cell bounds

### Requirement: Table Style References
The system SHALL support table style GUID references.

#### Scenario: Create table style ID
- WHEN NewTableStyleId() is created with GUID
- THEN TableStyleId element is created
- AND GUID value is stored

#### Scenario: Style ID XML
- GIVEN TableStyleId with value "{5C22544A-7EE6-4342-B048-85BDC9FD1C3A}"
- WHEN serialized
- THEN <a:tableStyleId>{5C22544A-7EE6-4342-B048-85BDC9FD1C3A}</a:tableStyleId> is generated

### Requirement: Table Namespace Handling
The system SHALL use correct XML namespaces for table elements.

#### Scenario: Table namespace
- GIVEN any table element
- WHEN namespace is queried
- THEN "http://schemas.openxmlformats.org/drawingml/2006/main" is returned
- AND prefix is "a"

#### Scenario: Nested elements use same namespace
- GIVEN a Table with rows and cells
- WHEN serialized
- THEN all elements (<a:tbl>, <a:tr>, <a:tc>) use "a:" prefix

### Requirement: Measurement Unit Consistency
The system SHALL use EMUs (English Metric Units) for all dimensional values in table elements.

#### Scenario: Column width in EMUs
- GIVEN a GridColumn
- WHEN width is set
- THEN value is stored in EMUs (914400 EMUs = 1 inch)

#### Scenario: Row height in EMUs
- GIVEN a TableRow
- WHEN height is set
- THEN value is stored in EMUs

#### Scenario: Cell margins in EMUs
- GIVEN CellProperties margins
- WHEN margins are set
- THEN values are stored in EMUs
