# Wordprocessing Tables Specification

## Requirements

### Requirement: Table Element
The system SHALL provide a Table element for tabular content.

#### Scenario: Access table properties
- GIVEN a Table element
- WHEN TableProperties() is called
- THEN the TableProperties element is returned

#### Scenario: Access table grid
- GIVEN a Table element
- WHEN TableGrid() is called
- THEN the TableGrid element defining columns is returned

#### Scenario: Access table rows
- GIVEN a Table element
- WHEN Elements[TableRow]() is called
- THEN all TableRow children are returned

#### Scenario: Get row count
- GIVEN a Table element
- WHEN RowCount() is called
- THEN the number of rows is returned

#### Scenario: Get column count
- GIVEN a Table element with TableGrid
- WHEN ColumnCount() is called
- THEN the number of columns is returned

#### Scenario: Append row
- GIVEN a Table element
- WHEN AppendChild(NewTableRow()) is called
- THEN a row is added to the table

#### Scenario: Create simple table
- GIVEN rows and columns count
- WHEN NewTable(rows, cols) is called
- THEN a table with grid and empty cells is created

### Requirement: TableRow Element
The system SHALL provide a TableRow element for table rows.

#### Scenario: Access row properties
- GIVEN a TableRow element
- WHEN TableRowProperties() is called
- THEN the TableRowProperties element is returned

#### Scenario: Access cells
- GIVEN a TableRow element
- WHEN Elements[TableCell]() is called
- THEN all TableCell children are returned

#### Scenario: Get cell count
- GIVEN a TableRow element
- WHEN CellCount() is called
- THEN the number of cells is returned

#### Scenario: Append cell
- GIVEN a TableRow element
- WHEN AppendChild(NewTableCell()) is called
- THEN a cell is added to the row

#### Scenario: Create row with cells
- GIVEN cell count
- WHEN NewTableRow(cellCount) is called
- THEN a row with specified empty cells is created

### Requirement: TableCell Element
The system SHALL provide a TableCell element for table cells.

#### Scenario: Access cell properties
- GIVEN a TableCell element
- WHEN TableCellProperties() is called
- THEN the TableCellProperties element is returned

#### Scenario: Access cell content
- GIVEN a TableCell element
- WHEN Elements[Paragraph]() is called
- THEN paragraph content is returned

#### Scenario: Cell requires paragraph
- GIVEN a TableCell element
- WHEN validated
- THEN at least one Paragraph child is required

#### Scenario: Get cell text
- GIVEN a TableCell element
- WHEN InnerText() is called
- THEN the concatenated text of all paragraphs is returned

#### Scenario: Set cell text
- GIVEN a TableCell element
- WHEN SetText(value) is called
- THEN the cell's paragraph content is set

#### Scenario: Create cell with text
- GIVEN text content
- WHEN NewTableCell(text) is called
- THEN a cell with paragraph containing text is created

### Requirement: TableGrid Element
The system SHALL provide a TableGrid element for column definitions.

#### Scenario: Access grid columns
- GIVEN a TableGrid element
- WHEN Elements[GridColumn]() is called
- THEN all GridColumn children are returned

#### Scenario: Set column width
- GIVEN a TableGrid element
- WHEN GridColumn[i].Width.SetValue(width) is called
- THEN the column width in twips is set

#### Scenario: Create table grid
- GIVEN column widths
- WHEN NewTableGrid(widths...) is called
- THEN a grid with specified column widths is created

### Requirement: TableProperties Element
The system SHALL provide TableProperties for table formatting.

#### Scenario: Table style
- GIVEN a TableProperties element
- WHEN TableStyle() is called
- THEN the table style ID is returned

#### Scenario: Set table style
- GIVEN a TableProperties element
- WHEN SetTableStyle("TableGrid") is called
- THEN the table style is applied

#### Scenario: Table width
- GIVEN a TableProperties element
- WHEN TableWidth() is called
- THEN the table width settings are returned

#### Scenario: Set table width
- GIVEN a TableProperties element
- WHEN TableWidth().Width.SetValue(5000) and Type.SetValue("pct") is called
- THEN 50% table width is set

#### Scenario: Table justification (alignment)
- GIVEN a TableProperties element
- WHEN TableJustification() is called
- THEN Left, Center, or Right is returned

#### Scenario: Table borders
- GIVEN a TableProperties element
- WHEN TableBorders() is called
- THEN border settings for all sides are returned

#### Scenario: Table cell margins
- GIVEN a TableProperties element
- WHEN TableCellMarginDefault() is called
- THEN default cell margin settings are returned

#### Scenario: Table cell spacing
- GIVEN a TableProperties element
- WHEN TableCellSpacing() is called
- THEN cell spacing settings are returned

#### Scenario: Table indent
- GIVEN a TableProperties element
- WHEN TableIndentation() is called
- THEN left indent from margin is returned

#### Scenario: Table look
- GIVEN a TableProperties element
- WHEN TableLook() is called
- THEN table style conditional formatting flags are returned

#### Scenario: Table layout
- GIVEN a TableProperties element
- WHEN TableLayout() is called
- THEN Fixed or AutoFit layout type is returned

### Requirement: TableRowProperties Element
The system SHALL provide TableRowProperties for row formatting.

#### Scenario: Row height
- GIVEN a TableRowProperties element
- WHEN TableRowHeight() is called
- THEN row height settings are returned

#### Scenario: Set row height
- GIVEN a TableRowProperties element
- WHEN TableRowHeight().Val.SetValue(720) is called
- THEN 0.5 inch row height is set

#### Scenario: Header row
- GIVEN a TableRowProperties element
- WHEN TableHeader() is called
- THEN whether row repeats as header is returned

#### Scenario: Set header row
- GIVEN a TableRowProperties element
- WHEN SetTableHeader(true) is called
- THEN the row repeats at top of each page

#### Scenario: Cannot split
- GIVEN a TableRowProperties element
- WHEN CantSplit() is called
- THEN whether row can split across pages is returned

#### Scenario: Row alignment
- GIVEN a TableRowProperties element
- WHEN TableJustification() is called
- THEN row alignment is returned

### Requirement: TableCellProperties Element
The system SHALL provide TableCellProperties for cell formatting.

#### Scenario: Cell width
- GIVEN a TableCellProperties element
- WHEN TableCellWidth() is called
- THEN cell width settings are returned

#### Scenario: Set cell width
- GIVEN a TableCellProperties element
- WHEN TableCellWidth().Width.SetValue(2880) is called
- THEN 2 inch cell width is set

#### Scenario: Vertical merge
- GIVEN a TableCellProperties element
- WHEN VerticalMerge() is called
- THEN Restart or Continue merge status is returned

#### Scenario: Set vertical merge
- GIVEN cells to merge vertically
- WHEN first cell VerticalMerge="restart" and following cells VerticalMerge="continue"
- THEN cells are merged vertically

#### Scenario: Horizontal merge (grid span)
- GIVEN a TableCellProperties element
- WHEN GridSpan() is called
- THEN the number of columns spanned is returned

#### Scenario: Set grid span
- GIVEN a TableCellProperties element
- WHEN GridSpan().Val.SetValue(2) is called
- THEN the cell spans 2 columns

#### Scenario: Cell borders
- GIVEN a TableCellProperties element
- WHEN TableCellBorders() is called
- THEN cell-specific border settings are returned

#### Scenario: Cell margins
- GIVEN a TableCellProperties element
- WHEN TableCellMargin() is called
- THEN cell-specific margin settings are returned

#### Scenario: Cell shading
- GIVEN a TableCellProperties element
- WHEN Shading() is called
- THEN cell background settings are returned

#### Scenario: Vertical alignment
- GIVEN a TableCellProperties element
- WHEN TableCellVerticalAlignment() is called
- THEN Top, Center, or Bottom is returned

#### Scenario: Text direction
- GIVEN a TableCellProperties element
- WHEN TextDirection() is called
- THEN text flow direction (LR, TB, etc.) is returned

#### Scenario: No wrap
- GIVEN a TableCellProperties element
- WHEN NoWrap() is called
- THEN whether text wraps in cell is returned

### Requirement: TableBorders Element
The system SHALL provide TableBorders for table border settings.

#### Scenario: Individual borders
- GIVEN a TableBorders element
- WHEN Top(), Bottom(), Left(), Right(), InsideH(), InsideV() are called
- THEN individual border settings are returned

#### Scenario: Set all borders
- GIVEN need for uniform borders
- WHEN SetAllBorders(style, size, color) is called
- THEN all borders are set uniformly

### Requirement: TableCellBorders Element
The system SHALL provide TableCellBorders for cell border settings.

#### Scenario: Individual cell borders
- GIVEN a TableCellBorders element
- WHEN Top(), Bottom(), Left(), Right() are called
- THEN individual cell border settings are returned

#### Scenario: Diagonal borders
- GIVEN a TableCellBorders element
- WHEN TopLeftToBottomRightBorder() and BottomLeftToTopRightBorder() are called
- THEN diagonal border settings are returned

### Requirement: Table Conditional Formatting
The system SHALL support conditional formatting via TableLook.

#### Scenario: First row styling
- GIVEN a TableLook element
- WHEN FirstRow() is accessed
- THEN whether first row special formatting applies is returned

#### Scenario: Last row styling
- GIVEN a TableLook element
- WHEN LastRow() is accessed
- THEN whether last row special formatting applies is returned

#### Scenario: First column styling
- GIVEN a TableLook element
- WHEN FirstColumn() is accessed
- THEN whether first column special formatting applies is returned

#### Scenario: Last column styling
- GIVEN a TableLook element
- WHEN LastColumn() is accessed
- THEN whether last column special formatting applies is returned

#### Scenario: Banding
- GIVEN a TableLook element
- WHEN NoHorizontalBand() and NoVerticalBand() are accessed
- THEN row/column banding status is returned

