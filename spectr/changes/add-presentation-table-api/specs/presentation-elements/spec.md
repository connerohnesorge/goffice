# Presentation Elements Spec Delta

## ADDED Requirements

### Requirement: PowerPoint Table Creation
The system SHALL provide methods to create tables in PowerPoint slides programmatically.

#### Scenario: Create table with rows and columns
- GIVEN a slide in a presentation
- WHEN AddTable(3, 4) is called
- THEN a table with 3 rows and 4 columns is created
- AND table is positioned at default location on slide
- AND table has default size

#### Scenario: Create table with minimum dimensions
- GIVEN a slide
- WHEN AddTable(1, 1) is called
- THEN a single-cell table is created
- AND no error occurs

#### Scenario: Create table with invalid dimensions
- GIVEN a slide
- WHEN AddTable(0, 5) is called with zero rows
- THEN an error is returned
- AND no table is created

#### Scenario: Multiple tables on single slide
- GIVEN a slide
- WHEN AddTable(2, 2) is called twice
- THEN two separate tables are created on the slide
- AND tables do not overlap (unless manually positioned to overlap)

### Requirement: Table Row Manipulation
The system SHALL provide methods to add, insert, and delete table rows.

#### Scenario: Append row to table
- GIVEN a table with 3 rows
- WHEN AppendRow() is called
- THEN row count increases to 4
- AND new row has same number of columns as existing rows
- AND new row cells are empty

#### Scenario: Insert row at specific position
- GIVEN a table with 3 rows
- WHEN InsertRow(1) is called
- THEN row count increases to 4
- AND new row is inserted at index 1 (second position)
- AND existing rows shift down

#### Scenario: Delete row from table
- GIVEN a table with 4 rows
- WHEN DeleteRow(1) is called
- THEN row count decreases to 3
- AND row at index 1 is removed
- AND remaining rows maintain their content

#### Scenario: Delete invalid row index
- GIVEN a table with 3 rows
- WHEN DeleteRow(5) is called with index out of bounds
- THEN an error is returned
- AND row count remains 3

### Requirement: Table Column Width Management
The system SHALL provide methods to set and retrieve column widths.

#### Scenario: Set column width in points
- GIVEN a table with 3 columns
- WHEN SetColumnWidth(1, 150) is called
- THEN column 1 width is set to 150 points
- AND width is converted to EMUs internally

#### Scenario: Get column width in points
- GIVEN a table with column 0 set to 100 points wide
- WHEN GetColumnWidth(0) is called
- THEN 100.0 is returned (in points)

#### Scenario: Set equal column widths
- GIVEN a table with 4 columns and total width 400 points
- WHEN SetEqualColumnWidths() is called
- THEN each column is 100 points wide

#### Scenario: Set column width with invalid index
- GIVEN a table with 3 columns
- WHEN SetColumnWidth(5, 100) is called with invalid index
- THEN an error is returned
- AND column widths remain unchanged

### Requirement: Table Cell Content Manipulation
The system SHALL provide methods to get and set cell content.

#### Scenario: Set cell text
- GIVEN a table
- WHEN SetCellText(0, 0, "Header 1") is called
- THEN cell at row 0, column 0 contains text "Header 1"

#### Scenario: Get cell text
- GIVEN a table with cell (1, 2) containing "Data"
- WHEN GetCellText(1, 2) is called
- THEN "Data" is returned

#### Scenario: Set cell text with invalid coordinates
- GIVEN a 3x3 table
- WHEN SetCellText(5, 5, "Invalid") is called
- THEN an error is returned
- AND no cell content is modified

#### Scenario: Get empty cell text
- GIVEN a table with empty cell (0, 0)
- WHEN GetCellText(0, 0) is called
- THEN empty string "" is returned
- AND no error occurs

### Requirement: Table Cell Access
The system SHALL provide methods to access individual table cells.

#### Scenario: Get cell by coordinates
- GIVEN a table
- WHEN GetCell(1, 2) is called
- THEN the cell at row 1, column 2 is returned

#### Scenario: Get cell with invalid coordinates
- GIVEN a 3x3 table
- WHEN GetCell(10, 10) is called
- THEN an error is returned

### Requirement: Table Position and Size
The system SHALL provide methods to position and size tables on slides.

#### Scenario: Set table position in points
- WHEN SetPosition(100, 50) is called
- THEN table is positioned at (100pt, 50pt) from slide origin
- AND position is converted to EMUs for storage

#### Scenario: Get table position in points
- GIVEN a table positioned at (150, 75) points
- WHEN GetPosition() is called
- THEN (150.0, 75.0) is returned

#### Scenario: Set table size in points
- WHEN SetSize(500, 300) is called
- THEN table dimensions are set to 500pt width, 300pt height
- AND column widths are recalculated to fit new width

#### Scenario: Get table size in points
- GIVEN a table sized to 400x200 points
- WHEN GetSize() is called
- THEN (400.0, 200.0) is returned

### Requirement: Table Row and Column Counts
The system SHALL provide methods to query table dimensions.

#### Scenario: Get row count
- GIVEN a table with 5 rows
- WHEN RowCount() is called
- THEN 5 is returned

#### Scenario: Get column count
- GIVEN a table with 4 columns
- WHEN ColumnCount() is called
- THEN 4 is returned

### Requirement: Table Style Application
The system SHALL provide methods to apply table styles using GUID references.

#### Scenario: Apply built-in table style
- GIVEN a table
- WHEN SetStyle("{5C22544A-7EE6-4342-B048-85BDC9FD1C3A}") is called
- THEN table style is set to Medium 2 style
- AND style GUID is stored in table properties

#### Scenario: Enable banded rows
- GIVEN a table
- WHEN SetBandedRows(true) is called
- THEN banded row formatting is enabled
- AND alternate rows will render with different backgrounds

#### Scenario: Disable banded rows
- GIVEN a table with banded rows enabled
- WHEN SetBandedRows(false) is called
- THEN banded row formatting is disabled

#### Scenario: Enable first row bold
- GIVEN a table
- WHEN SetFirstRowBold(true) is called
- THEN first row is marked for special header formatting

#### Scenario: Enable first column bold
- GIVEN a table
- WHEN SetFirstColumnBold(true) is called
- THEN first column is marked for special formatting

#### Scenario: Enable last row bold
- GIVEN a table
- WHEN SetLastRowBold(true) is called
- THEN last row is marked for special total row formatting

### Requirement: Table Cell Formatting
The system SHALL provide methods to format individual table cells.

#### Scenario: Set cell background color
- GIVEN a table
- WHEN SetCellBackgroundColor(0, 0, 255, 200, 100) is called
- THEN cell (0, 0) background is set to RGB(255, 200, 100)

#### Scenario: Set cell vertical alignment
- GIVEN a table
- WHEN SetCellAlignment(1, 1, TextAnchorCenter) is called
- THEN cell (1, 1) text is vertically centered

#### Scenario: Set cell margins
- GIVEN a table
- WHEN SetCellMargins(0, 0, 5, 5, 3, 3) is called
- THEN cell (0, 0) has 5pt left/right margins and 3pt top/bottom margins

### Requirement: Table Integration with GraphicFrame
The system SHALL integrate tables with the existing GraphicFrame infrastructure.

#### Scenario: Table uses GraphicFrame for positioning
- GIVEN a table created with AddTable()
- WHEN the table is examined internally
- THEN table is wrapped in a GraphicFrame element
- AND GraphicFrame manages position and size

#### Scenario: Table GraphicFrame has correct URI
- GIVEN a table created with AddTable()
- WHEN the GraphicFrame's graphic data is examined
- THEN URI is "http://schemas.openxmlformats.org/drawingml/2006/table"

#### Scenario: Multiple GraphicFrames on slide with different content
- GIVEN a slide with a chart and a table
- WHEN GraphicFrames are enumerated
- THEN 2 GraphicFrames are found
- AND chart GraphicFrame has chart URI
- AND table GraphicFrame has table URI
