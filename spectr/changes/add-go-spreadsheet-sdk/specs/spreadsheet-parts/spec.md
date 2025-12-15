## ADDED Requirements

### Requirement: WorkbookPart

The system SHALL provide a `WorkbookPart` type that contains the workbook root element and manages sheet references.

#### Scenario: Access workbook part from document
- GIVEN an open SpreadsheetDocument
- WHEN `document.WorkbookPart()` is accessed
- THEN the WorkbookPart is returned
- AND it contains the workbook root element

#### Scenario: Add workbook part to document
- GIVEN a new SpreadsheetDocument without a workbook part
- WHEN `document.AddWorkbookPart()` is called
- THEN a new WorkbookPart is created and attached
- AND the workbook root element is initialized

#### Scenario: Access worksheets through workbook
- GIVEN a WorkbookPart with multiple worksheets
- WHEN `workbook.WorksheetParts()` is called
- THEN all WorksheetParts are returned in sheet order

### Requirement: WorksheetPart

The system SHALL provide a `WorksheetPart` type that contains individual worksheet content.

#### Scenario: Add worksheet to workbook
- GIVEN a WorkbookPart
- WHEN `workbook.AddNewPart[WorksheetPart]()` is called
- THEN a new WorksheetPart is created
- AND the worksheet root element is initialized
- AND a relationship is created to the workbook

#### Scenario: Access worksheet root element
- GIVEN a WorksheetPart
- WHEN `part.Worksheet()` is accessed
- THEN the Worksheet root element is returned
- AND it is lazily loaded on first access

#### Scenario: Access sheet data
- GIVEN a WorksheetPart
- WHEN `part.Worksheet().SheetData` is accessed
- THEN the SheetData element containing rows and cells is returned

### Requirement: SharedStringTablePart

The system SHALL provide a `SharedStringTablePart` type for string deduplication.

#### Scenario: Add shared string table to workbook
- GIVEN a WorkbookPart without a shared string table
- WHEN `workbook.AddSharedStringTablePart()` is called
- THEN a new SharedStringTablePart is created
- AND the SharedStringTable root element is initialized

#### Scenario: Add string to shared string table
- GIVEN a SharedStringTablePart
- WHEN a new unique string is added via cell
- THEN the string is added to the SharedStringTable
- AND the index is returned for cell reference

#### Scenario: Reuse existing shared string
- GIVEN a SharedStringTablePart with existing strings
- WHEN a duplicate string is set on a cell
- THEN the existing index is returned
- AND no duplicate entry is created

### Requirement: WorkbookStylesPart

The system SHALL provide a `WorkbookStylesPart` type for style definitions.

#### Scenario: Add styles part to workbook
- GIVEN a WorkbookPart without a styles part
- WHEN `workbook.AddWorkbookStylesPart()` is called
- THEN a new WorkbookStylesPart is created
- AND default styles are initialized

#### Scenario: Access style collections
- GIVEN a WorkbookStylesPart
- WHEN `styles.Stylesheet()` is accessed
- THEN the Stylesheet root element is returned
- AND NumberFormats, Fonts, Fills, Borders, CellXfs are accessible

### Requirement: CalculationChainPart

The system SHALL provide a `CalculationChainPart` type for formula calculation order.

#### Scenario: Add calculation chain to workbook
- GIVEN a WorkbookPart with formulas
- WHEN `workbook.AddCalculationChainPart()` is called
- THEN a new CalculationChainPart is created
- AND it tracks formula calculation dependencies

#### Scenario: Access calculation chain
- GIVEN a CalculationChainPart
- WHEN `part.CalculationChain()` is accessed
- THEN the ordered list of formula cells is returned

### Requirement: ThemePart

The system SHALL provide a `ThemePart` type for document theming.

#### Scenario: Add theme to workbook
- GIVEN a WorkbookPart
- WHEN `workbook.AddThemePart()` is called
- THEN a new ThemePart is created
- AND default Office theme is initialized

#### Scenario: Access theme colors
- GIVEN a ThemePart
- WHEN `theme.Theme().ThemeElements.ColorScheme` is accessed
- THEN theme colors (dk1, lt1, accent1-6, etc.) are returned

### Requirement: ChartsheetPart

The system SHALL provide a `ChartsheetPart` type for chart-only sheets.

#### Scenario: Add chartsheet to workbook
- GIVEN a WorkbookPart
- WHEN `workbook.AddNewPart[ChartsheetPart]()` is called
- THEN a new ChartsheetPart is created
- AND the chartsheet root element is initialized

### Requirement: PivotTablePart

The system SHALL provide a `PivotTablePart` type for pivot table definitions.

#### Scenario: Add pivot table to worksheet
- GIVEN a WorksheetPart
- WHEN `worksheet.AddNewPart[PivotTablePart]()` is called
- THEN a new PivotTablePart is created
- AND it references a PivotTableCacheDefinitionPart

### Requirement: PivotTableCacheDefinitionPart

The system SHALL provide a `PivotTableCacheDefinitionPart` type for pivot cache data.

#### Scenario: Add pivot cache to workbook
- GIVEN a WorkbookPart
- WHEN `workbook.AddNewPart[PivotTableCacheDefinitionPart]()` is called
- THEN a new PivotTableCacheDefinitionPart is created
- AND it can be referenced by multiple pivot tables

### Requirement: TableDefinitionPart

The system SHALL provide a `TableDefinitionPart` type for Excel table definitions.

#### Scenario: Add table to worksheet
- GIVEN a WorksheetPart
- WHEN `worksheet.AddNewPart[TableDefinitionPart]()` is called
- THEN a new TableDefinitionPart is created
- AND the table root element is initialized

### Requirement: DrawingsPart

The system SHALL provide a `DrawingsPart` type for embedded drawings.

#### Scenario: Add drawings to worksheet
- GIVEN a WorksheetPart
- WHEN `worksheet.AddDrawingsPart()` is called
- THEN a new DrawingsPart is created
- AND the WorksheetDrawing root element is initialized

### Requirement: ChartPart

The system SHALL provide a `ChartPart` type for chart definitions.

#### Scenario: Add chart to drawings
- GIVEN a DrawingsPart
- WHEN `drawings.AddNewPart[ChartPart]()` is called
- THEN a new ChartPart is created
- AND the ChartSpace root element is initialized

### Requirement: WorksheetCommentsPart

The system SHALL provide a `WorksheetCommentsPart` type for cell comments.

#### Scenario: Add comments to worksheet
- GIVEN a WorksheetPart
- WHEN `worksheet.AddWorksheetCommentsPart()` is called
- THEN a new WorksheetCommentsPart is created
- AND the Comments root element is initialized

### Requirement: VmlDrawingPart

The system SHALL provide a `VmlDrawingPart` type for legacy VML drawings.

#### Scenario: Access VML drawings for comments
- GIVEN a WorksheetPart with comments
- WHEN `worksheet.VmlDrawingParts()` is accessed
- THEN VML drawing parts containing comment shapes are returned

### Requirement: QueryTablePart

The system SHALL provide a `QueryTablePart` type for external data queries.

#### Scenario: Add query table to worksheet
- GIVEN a WorksheetPart
- WHEN `worksheet.AddNewPart[QueryTablePart]()` is called
- THEN a new QueryTablePart is created
- AND it defines an external data query

### Requirement: SlicerPart

The system SHALL provide a `SlicerPart` type for slicer controls.

#### Scenario: Add slicer to worksheet
- GIVEN a WorksheetPart
- WHEN `worksheet.AddNewPart[SlicerPart]()` is called
- THEN a new SlicerPart is created
- AND the Slicers root element is initialized

### Requirement: TimeLinePart

The system SHALL provide a `TimeLinePart` type for timeline controls.

#### Scenario: Add timeline to worksheet
- GIVEN a WorksheetPart
- WHEN `worksheet.AddNewPart[TimeLinePart]()` is called
- THEN a new TimeLinePart is created
- AND the Timelines root element is initialized

### Requirement: ExternalWorkbookPart

The system SHALL provide an `ExternalWorkbookPart` type for external references.

#### Scenario: Add external workbook reference
- GIVEN a WorkbookPart
- WHEN `workbook.AddNewPart[ExternalWorkbookPart]()` is called
- THEN a new ExternalWorkbookPart is created
- AND it can reference an external workbook for formula links
