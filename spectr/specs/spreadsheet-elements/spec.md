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

### Requirement: Office 2010 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2010 extensions.

#### Scenario: Slicer elements
- GIVEN Office 2010 slicer elements (x14 namespace)
- WHEN parsed from XML
- THEN Slicer, SlicerCache, SlicerCacheDefinition elements are created

#### Scenario: Conditional formatting extensions
- GIVEN Office 2010 conditional formatting extensions (x14 namespace)
- WHEN parsed
- THEN ConditionalFormatting, DataBar, IconSet, ColorScale elements available

#### Scenario: Sparkline elements
- GIVEN Office 2010 sparkline elements (x14 namespace)
- WHEN parsed
- THEN Sparkline, SparklineGroup elements available

#### Scenario: PivotTable extensions
- GIVEN Office 2010 pivot table extensions (x14 namespace)
- WHEN parsed
- THEN PivotTableDefinition extensions available

### Requirement: Office 2013-2015 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2013-2015 extensions.

#### Scenario: Timeline elements
- GIVEN Office 2013 timeline elements (x15 namespace)
- WHEN parsed
- THEN Timeline, TimelineCacheDefinition, TimelineStyle elements available

#### Scenario: Chart elements
- GIVEN Office 2015 chart extensions (x15 namespace)
- WHEN parsed
- THEN extended chart elements available

#### Scenario: Data model elements
- GIVEN Office 2015 data model extensions (x15 namespace)
- WHEN parsed
- THEN DataModel, ModelTable, ModelRelationship elements available

### Requirement: Office 2016 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2016 extensions.

#### Scenario: Revision elements (x16r5)
- GIVEN Office 2016 revision elements (x16r5 namespace)
- WHEN parsed
- THEN revision tracking elements available

#### Scenario: Pivot default layout elements
- GIVEN Office 2016 pivot default layout elements (x16pdl namespace)
- WHEN parsed
- THEN PivotDefaultLayout element available

#### Scenario: SVG support elements
- GIVEN Office 2016 SVG elements (x16svg namespace)
- WHEN parsed
- THEN SVG image reference elements available

### Requirement: Office 2017-2020 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2017-2020 extensions.

#### Scenario: Rich data elements
- GIVEN Office 2017 rich data elements (xlrd namespace)
- WHEN parsed
- THEN RichValue, RichValueStructure elements available

#### Scenario: Rich data web image elements
- GIVEN Office 2020 rich data web image elements (xlrdwi namespace)
- WHEN parsed
- THEN WebImageElement available

### Requirement: Office 2021-2024 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2021-2024 extensions.

#### Scenario: Rich value relation elements
- GIVEN Office 2022 rich value relation elements (xlrvrel namespace)
- WHEN parsed
- THEN RichValueRelation elements available

#### Scenario: Pivot auto refresh elements
- GIVEN Office 2024 pivot auto refresh elements (xpar namespace)
- WHEN parsed
- THEN PivotAutoRefresh element available

### Requirement: Office 2025 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2025 extensions.

#### Scenario: Pivot data source elements
- GIVEN Office 2025 pivot data source elements (x25pds namespace)
- WHEN parsed
- THEN PivotDataSource elements available

#### Scenario: External code service elements
- GIVEN Office 2025 external code service elements (x25ecs namespace)
- WHEN parsed
- THEN ExternalCodeService elements available

### Requirement: Spreadsheet Extension Element Attributes
The system SHALL provide typed attributes for spreadsheet extension elements.

#### Scenario: Slicer name attribute
- GIVEN a Slicer element with name attribute
- WHEN Name is accessed
- THEN StringValue wrapper with slicer name is returned

#### Scenario: Timeline cache id attribute
- GIVEN a Timeline element with cache id
- WHEN CacheId is accessed
- THEN UInt32Value wrapper is returned

#### Scenario: Rich value type attribute
- GIVEN a RichValue element with type attribute
- WHEN Type is accessed
- THEN enum value from rich data schema is returned

### Requirement: Spreadsheet Extension Element Version Metadata
The system SHALL provide version information for spreadsheet extension elements.

#### Scenario: Office 2010 element version
- GIVEN a Slicer element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2010 is returned

#### Scenario: Office 2024 element version
- GIVEN a PivotAutoRefresh element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2024 is returned

#### Scenario: Office 2025 element version
- GIVEN a PivotDataSource element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2025 is returned

### Requirement: Spreadsheet Extension Element Namespace Handling
The system SHALL correctly handle namespaces for spreadsheet extension elements.

#### Scenario: Excel 2010 namespace
- GIVEN a Slicer element
- WHEN NamespaceURI() is called
- THEN "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main" is returned

#### Scenario: Namespace prefix in XML
- GIVEN a Slicer element
- WHEN written to XML
- THEN element is serialized as `<x14:slicer>`

#### Scenario: Rich data namespace
- GIVEN a RichValue element
- WHEN NamespaceURI() is called
- THEN rich data namespace URI is returned

### Requirement: Spreadsheet Extension Element Generation
The system SHALL generate spreadsheet element types from extension schemas in addition to main schemas.

#### Scenario: Generate from main schema
- GIVEN spreadsheetml main.json schema
- WHEN generator runs
- THEN core SpreadsheetML elements are generated

#### Scenario: Generate from extension schemas
- GIVEN Excel 2010-2025 extension schemas
- WHEN generator runs
- THEN extension element types are generated with correct namespaces

#### Scenario: Total element count increase
- GIVEN all schemas loaded
- WHEN generation completes
- THEN 80-100 additional element types exist beyond baseline

#### Scenario: Generated element organization
- GIVEN generated spreadsheet elements
- WHEN inspecting spreadsheet/elements/ package
- THEN extension elements are organized with main elements

### Requirement: Spreadsheet Extension Element Registration
The system SHALL register all spreadsheet extension element types.

#### Scenario: Register slicer elements
- GIVEN Slicer, SlicerCache element types
- WHEN element registry is initialized
- THEN slicer elements are registered with x14 namespace

#### Scenario: Parse registered extension element
- GIVEN XML with x14:slicer element
- WHEN parsing workbook
- THEN Slicer instance is created (not UnknownElement)

#### Scenario: Parse timeline element
- GIVEN XML with x15:timeline element
- WHEN parsing workbook
- THEN Timeline instance is created
