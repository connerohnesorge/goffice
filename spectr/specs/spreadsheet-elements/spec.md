# Spreadsheet Elements Specification

## Requirements

### Requirement: Workbook Root Element

The system SHALL provide a `Workbook` element as the root of WorkbookPart.

#### Scenario: Access workbook properties
- GIVEN a Workbook element
- WHEN properties are accessed
- THEN FileVersion, WorkbookPr, BookViews, Sheets, DefinedNames, CalcPr are accessible

#### Scenario: Access sheet list
- GIVEN a Workbook element
- WHEN `workbook.Sheets` is accessed
- THEN the Sheets element containing all Sheet references is returned

### Requirement: Sheet Element

The system SHALL provide a `Sheet` element for sheet references in the workbook.

#### Scenario: Sheet properties
- GIVEN a Sheet element
- WHEN properties are accessed
- THEN Name, SheetId, State, and relationship Id are accessible

#### Scenario: Sheet states
- GIVEN a Sheet element
- WHEN State is set
- THEN values Visible, Hidden, or VeryHidden are supported

### Requirement: Worksheet Root Element

The system SHALL provide a `Worksheet` element as the root of WorksheetPart.

#### Scenario: Access worksheet components
- GIVEN a Worksheet element
- WHEN components are accessed
- THEN SheetPr, Dimension, SheetViews, SheetFormatPr, Cols, SheetData, MergeCells, ConditionalFormatting, DataValidations, Hyperlinks, PageMargins, PageSetup are accessible

### Requirement: SheetData Element

The system SHALL provide a `SheetData` element containing all row and cell data.

#### Scenario: Access rows in sheet data
- GIVEN a SheetData element
- WHEN rows are enumerated
- THEN all Row elements are returned in row order

#### Scenario: Add row to sheet data
- GIVEN a SheetData element
- WHEN a new Row is appended
- THEN the row is added at the specified row index

### Requirement: Row Element

The system SHALL provide a `Row` element for worksheet rows.

#### Scenario: Row properties
- GIVEN a Row element
- WHEN properties are accessed
- THEN RowIndex, Spans, Height, CustomHeight, Hidden, Collapsed, StyleIndex are accessible

#### Scenario: Access cells in row
- GIVEN a Row element
- WHEN cells are enumerated
- THEN all Cell elements in the row are returned

### Requirement: Cell Element

The system SHALL provide a `Cell` element for individual cells.

#### Scenario: Cell properties
- GIVEN a Cell element
- WHEN properties are accessed
- THEN CellReference, StyleIndex, DataType, CellValue, CellFormula are accessible

#### Scenario: Cell data types
- GIVEN a Cell element
- WHEN DataType is set
- THEN values Boolean, Date, Error, InlineString, Number, SharedString, String are supported

### Requirement: CellValue Element

The system SHALL provide a `CellValue` element for cell values.

#### Scenario: Numeric cell value
- GIVEN a Cell with numeric data
- WHEN CellValue is accessed
- THEN the numeric value is stored as a string representation

#### Scenario: Shared string cell value
- GIVEN a Cell with DataType SharedString
- WHEN CellValue is accessed
- THEN the value is an index into the shared string table

### Requirement: CellFormula Element

The system SHALL provide a `CellFormula` element for formulas.

#### Scenario: Normal formula
- GIVEN a Cell with a formula
- WHEN CellFormula is accessed
- THEN the formula text is returned
- AND FormulaType is Normal (default)

#### Scenario: Shared formula
- GIVEN a Cell with a shared formula
- WHEN CellFormula is accessed
- THEN SharedIndex references the formula group
- AND Reference specifies the range for the master cell

#### Scenario: Array formula
- GIVEN a Cell with an array formula
- WHEN CellFormula is accessed
- THEN FormulaType is Array
- AND Reference specifies the array range

### Requirement: SheetView Element

The system SHALL provide a `SheetView` element for view settings.

#### Scenario: Sheet view properties
- GIVEN a SheetView element
- WHEN properties are accessed
- THEN TabSelected, ZoomScale, ShowGridLines, ShowRowColHeaders, View are accessible

#### Scenario: Frozen panes
- GIVEN a SheetView with frozen panes
- WHEN Pane element is accessed
- THEN XSplit, YSplit, TopLeftCell, ActivePane, State are accessible

#### Scenario: Selection
- GIVEN a SheetView with a selection
- WHEN Selection elements are accessed
- THEN ActiveCell, SelectedCellsRange (sqref) are accessible

### Requirement: Cols Element

The system SHALL provide a `Cols` element for column formatting.

#### Scenario: Column properties
- GIVEN a Cols element with Col children
- WHEN Col elements are accessed
- THEN Min, Max, Width, CustomWidth, Hidden, Style, BestFit, Collapsed are accessible

#### Scenario: Apply column width
- GIVEN a Col element
- WHEN Width is set
- THEN the column width applies to columns Min through Max

### Requirement: MergeCells Element

The system SHALL provide a `MergeCells` element for merged cell ranges.

#### Scenario: Access merge cells
- GIVEN a Worksheet with merged cells
- WHEN MergeCells is accessed
- THEN all MergeCell elements with Reference attributes are returned

#### Scenario: Add merge cell
- GIVEN a MergeCells element
- WHEN a new MergeCell is added with Reference "A1:D1"
- THEN cells A1 through D1 are merged

### Requirement: ConditionalFormatting Element

The system SHALL provide `ConditionalFormatting` elements for conditional formatting rules.

#### Scenario: Conditional formatting rules
- GIVEN a ConditionalFormatting element
- WHEN CfRule children are accessed
- THEN Type, Priority, Operator, Formula, DxfId are accessible

#### Scenario: Conditional formatting types
- GIVEN a CfRule element
- WHEN Type is set
- THEN values CellIs, Expression, ColorScale, DataBar, IconSet, Top10, AboveAverage, etc. are supported

### Requirement: DataValidation Element

The system SHALL provide a `DataValidation` element for input validation.

#### Scenario: Data validation properties
- GIVEN a DataValidation element
- WHEN properties are accessed
- THEN Type, Operator, ShowErrorMessage, ErrorTitle, Error, ShowInputMessage, PromptTitle, Prompt, Formula1, Formula2, Sqref are accessible

#### Scenario: List validation
- GIVEN a DataValidation with Type List
- WHEN Formula1 contains a comma-separated list or range
- THEN a dropdown is shown in Excel

### Requirement: Hyperlinks Element

The system SHALL provide a `Hyperlinks` element for cell hyperlinks.

#### Scenario: Hyperlink properties
- GIVEN a Hyperlink element
- WHEN properties are accessed
- THEN Ref, relationship Id, Location, Display, Tooltip are accessible

### Requirement: PageSetup Element

The system SHALL provide a `PageSetup` element for print settings.

#### Scenario: Page setup properties
- GIVEN a PageSetup element
- WHEN properties are accessed
- THEN PaperSize, Orientation, Scale, FitToWidth, FitToHeight, FirstPageNumber, BlackAndWhite are accessible

### Requirement: HeaderFooter Element

The system SHALL provide a `HeaderFooter` element for page headers/footers.

#### Scenario: Header footer properties
- GIVEN a HeaderFooter element
- WHEN OddHeader, OddFooter, EvenHeader, EvenFooter, FirstHeader, FirstFooter are accessed
- THEN the header/footer text with formatting codes is returned

### Requirement: AutoFilter Element

The system SHALL provide an `AutoFilter` element for data filtering.

#### Scenario: Auto filter properties
- GIVEN an AutoFilter element
- WHEN properties are accessed
- THEN Ref defines the filter range
- AND FilterColumn children define column filters

### Requirement: SortState Element

The system SHALL provide a `SortState` element for sorting.

#### Scenario: Sort state properties
- GIVEN a SortState element
- WHEN properties are accessed
- THEN Ref, CaseSensitive, SortCondition children are accessible

### Requirement: DefinedName Element

The system SHALL provide a `DefinedName` element for named ranges.

#### Scenario: Defined name properties
- GIVEN a DefinedName element
- WHEN properties are accessed
- THEN Name, LocalSheetId, Hidden, and formula content are accessible

#### Scenario: Global vs local scope
- GIVEN a DefinedName element
- WHEN LocalSheetId is set
- THEN the name is scoped to that sheet
- WHEN LocalSheetId is not set
- THEN the name is global to the workbook

