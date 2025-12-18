# Change: Add Go SDK for Office Open XML Spreadsheets

## Why

The Office Open XML (OOXML) format is the standard for Microsoft Excel spreadsheets (.xlsx, .xltx, .xlsm, .xltm, .xlam). Currently, Go lacks a comprehensive, well-structured SDK for creating, reading, and manipulating Excel workbooks that mirrors the capabilities of the official Microsoft Open-XML-SDK for .NET. This proposal extends **goffice** with SpreadsheetML support, completing the Office document trifecta alongside the existing Word and Presentation SDKs.

**Problem Statement:**
- Existing Go libraries for Excel files (excelize, xlsx) lack full OOXML compliance and Open-XML-SDK feature parity
- No Go library provides the complete feature set including validation, pivot tables, charts, and all cell types
- Developers need strongly-typed, idiomatic Go APIs for spreadsheet manipulation
- Enterprise applications require reliable spreadsheet generation without Excel dependencies

**Opportunity:**
- Provide feature parity with Microsoft's Open-XML-SDK SpreadsheetML capabilities
- Enable server-side spreadsheet generation in Go microservices
- Support all Excel document types: xlsx, xltx, xlsm, xltm, xlam
- Deliver proper validation against ECMA-376 and ISO/IEC 29500 standards
- Leverage shared framework components from Word and Presentation SDKs

## What Changes

### Spreadsheet Domain (New)

1. **Spreadsheet Document** (`spreadsheet-document`)
   - `SpreadsheetDocument` type supporting all spreadsheet formats:
     - Workbook (`.xlsx`)
     - Template (`.xltx`)
     - MacroEnabledWorkbook (`.xlsm`)
     - MacroEnabledTemplate (`.xltm`)
     - AddIn (`.xlam`)
   - Create from file, stream, or template
   - Open existing documents (read-only or editable)
   - Document type conversion
   - FlatOPC format support (single XML representation)

2. **Spreadsheet Parts** (`spreadsheet-parts`)
   - `WorkbookPart` - main workbook content with sheet references
   - `WorksheetPart` - individual worksheet content
   - `SharedStringTablePart` - shared string storage for efficiency
   - `WorkbookStylesPart` - cell styles, formats, fonts, fills, borders
   - `CalculationChainPart` - formula calculation order
   - `ConnectionsPart` - external data connections
   - `ChartsheetPart` - chart-only sheets
   - `DialogsheetPart` - dialog sheets (legacy)
   - `MacroSheetPart` - macro sheets (legacy)
   - `PivotTablePart`, `PivotTableCacheDefinitionPart`, `PivotTableCacheRecordsPart`
   - `TableDefinitionPart` - Excel table definitions
   - `SlicerPart`, `SlicerCachePart` - slicer controls
   - `TimeLinePart`, `TimeLineCachePart` - timeline controls
   - `DrawingsPart` - embedded drawings and charts
   - `ChartPart`, `ChartColorStylePart`, `ChartStylePart` - chart components
   - `CommentsPart`, `ThreadedCommentsPart` - cell comments
   - `QueryTablePart` - query table definitions
   - `VolatileDependenciesPart` - volatile function dependencies
   - `CellMetadataPart` - cell metadata
   - `RichDataParts` - rich data types (Office 2019+)

3. **Spreadsheet Elements** (`spreadsheet-elements`)
   - All SpreadsheetML schema elements from ECMA-376
   - Strongly-typed element classes with property accessors
   - Workbook structure: `Workbook`, `Sheets`, `Sheet`, `DefinedNames`
   - Worksheet structure: `Worksheet`, `SheetData`, `SheetView`, `Cols`
   - Row and cell: `Row`, `Cell`, `CellValue`, `CellFormula`
   - Sheet formatting: `SheetFormatPr`, `SheetProtection`, `ConditionalFormatting`
   - Data validation: `DataValidations`, `DataValidation`
   - Filtering and sorting: `AutoFilter`, `SortState`, `FilterColumn`
   - Merge cells: `MergeCells`, `MergeCell`
   - Page setup: `PageSetup`, `PageMargins`, `HeaderFooter`
   - Print settings: `PrintOptions`, `PageBreaks`
   - Sheet views: `SheetViews`, `SheetView`, `Pane`, `Selection`

4. **Cell System** (`spreadsheet-cells`)
   - Cell reference system (A1, R1C1 notation)
   - Cell value types: String, Number, Boolean, Date, Error, SharedString
   - Cell formatting: NumberFormat, Font, Fill, Border, Alignment, Protection
   - Rich text in cells (multiple formats within single cell)
   - Hyperlinks in cells
   - Cell comments and threaded comments
   - Inline strings vs shared strings

5. **Formula System** (`spreadsheet-formulas`)
   - Formula storage and representation
   - Shared formulas
   - Array formulas (CSE and dynamic arrays)
   - Data table formulas
   - Calculation chain management
   - Defined names (named ranges, formulas)
   - External references (links to other workbooks)
   - Structured references (table references)

6. **Style System** (`spreadsheet-styles`)
   - Cell styles (built-in and custom)
   - Number formats (built-in and custom)
   - Fonts collection
   - Fills collection (solid, pattern, gradient)
   - Borders collection
   - Cell style XF (cell formatting records)
   - Differential formatting (for conditional formatting)
   - Table styles
   - Indexed colors and theme colors

7. **Charts** (`spreadsheet-charts`)
   - All chart types: Bar, Column, Line, Pie, Area, Scatter, etc.
   - Chart series and data points
   - Chart axes and legends
   - Chart titles and labels
   - Chart styles and formatting
   - Sparklines (mini-charts in cells)
   - Chart sheets vs embedded charts

8. **Pivot Tables** (`spreadsheet-pivottables`)
   - Pivot table definitions
   - Pivot cache and records
   - Pivot fields and items
   - Calculated fields and items
   - Pivot table styles
   - Slicers and timelines

9. **Drawing Integration** (`spreadsheet-drawing`)
   - SpreadsheetDrawing elements (xdr: namespace)
   - Cell-anchored objects (oneCellAnchor, twoCellAnchor, absoluteAnchor)
   - Shapes, pictures, charts in worksheets
   - Drawing properties and transforms

### Shared Components (Reused)

The spreadsheet SDK shares these components with Word and Presentation SDKs:
- **Core Framework** (`pkg/framework`) - Element base types, features
- **Simple Types** (`pkg/types`) - Value types like StringValue, Int32Value
- **Packaging** (`pkg/package`) - OPC package handling
- **Validation** (`pkg/validation`) - Schema and semantic validation
- **DrawingML** (`pkg/dml`) - Common drawing elements (a: namespace)

### BREAKING Changes
- None (new project)

### Non-Breaking Changes
- Introduces new `pkg/spreadsheet/` package with all spreadsheet types
- Introduces new `pkg/sml/` package with SpreadsheetML schema elements (x: namespace)
- Introduces new `pkg/xdr/` package with SpreadsheetDrawing elements

## Impact

- **Affected specs**: None (all new capabilities)
- **Affected code**: Creates new packages under `pkg/`:
  - `pkg/spreadsheet/` - SpreadsheetDocument and parts
  - `pkg/sml/` - SpreadsheetML elements (generated from schemas)
  - `pkg/xdr/` - SpreadsheetDrawing elements
- **Dependencies**:
  - Pure Go stdlib only (`archive/zip`, `encoding/xml`, `sync`)
  - Shares common components with Word and Presentation SDKs

## Resolved Configuration

| Setting | Value |
|---------|-------|
| **Go Module** | `github.com/connerohnesorge/goffice` |
| **Go Version** | Go 1.25+ |
| **Dependencies** | Pure stdlib only |
| **Generated Code** | Pre-generated from JSON schemas, committed |
| **API Style** | Direct struct field access |
| **Error Handling** | Return errors everywhere |
| **Validation** | Strict by default |
| **Thread Safety** | Document-level sync.RWMutex (concurrent reads via RLock(), exclusive writes via Lock()) |
| **MVP Scope** | Full read/write from v0.1 |
| **MC Handling** | Process automatically (select Choice/Fallback) |
| **Testing Strategy** | Golden file tests |
| **XML Prefixes** | Fixed canonical (x:, a:, r:, xdr:) |
| **Package Structure** | Domain-based: `drawingml/`, `packaging/`, `openxml/`, `spreadsheet/{elements,parts}/` |
| **Version Support** | Office 2016+ (FileFormatVersions.Office2016) |

## Architecture Overview

The package structure has been updated to use a cleaner top-level organization with shared DrawingML:

```
goffice/
├── drawingml/               # SHARED - shapes, images, charts (a: namespace)
│   └── *.go                 # Pre-generated DrawingML types
│
├── packaging/               # OPC layer (shared)
│   ├── package.go           # Package struct with ZIP backing
│   ├── part.go              # Part struct
│   ├── relationship.go      # Relationship struct
│   └── content_types.go     # Content types management
│
├── openxml/                 # Core framework (shared)
│   ├── element.go           # Element interface and base types
│   ├── composite.go         # CompositeElement for containers
│   ├── leaf.go              # LeafElement for simple content
│   ├── part.go              # OpenXmlPart base struct
│   ├── validation.go        # Validation framework
│   └── types/               # OpenXML value types
│       ├── string.go        # StringValue
│       ├── int.go           # Int32Value, Int64Value, etc.
│       ├── bool.go          # BooleanValue
│       └── ...              # Other simple types
│
├── spreadsheet/             # SpreadsheetML domain
│   ├── elements/            # Workbook, Worksheet, Cell, Formula elements
│   │   ├── workbook.go      # x:workbook, x:sheets, x:sheet, etc.
│   │   ├── worksheet.go     # x:worksheet, x:sheetData, x:sheetViews
│   │   ├── cell.go          # x:row, x:c, x:v, x:f
│   │   ├── styles.go        # x:styleSheet, x:fonts, x:fills, etc.
│   │   ├── strings.go       # x:sst, x:si shared strings
│   │   ├── table.go         # x:table elements
│   │   ├── pivot.go         # Pivot table elements
│   │   ├── chart.go         # Chart integration (uses drawingml/)
│   │   └── conditional.go   # Conditional formatting elements
│   │
│   └── parts/               # Spreadsheet parts
│       ├── workbook.go      # WorkbookPart
│       ├── worksheet.go     # WorksheetPart
│       ├── styles.go        # WorkbookStylesPart
│       ├── strings.go       # SharedStringTablePart
│       ├── chart.go         # ChartPart
│       ├── drawing.go       # DrawingsPart
│       ├── pivot.go         # PivotTablePart, PivotCacheParts
│       ├── table.go         # TableDefinitionPart
│       └── comments.go      # CommentsPart, ThreadedCommentsPart
│
├── word/                    # WordprocessingML (existing/planned)
│   ├── elements/            # Document, Paragraph, Run elements
│   └── parts/               # Document parts
│
├── presentation/            # PresentationML (existing/planned)
│   ├── elements/            # Presentation, Slide elements
│   └── parts/               # Presentation parts
│
├── internal/
│   ├── opc/                 # Low-level OPC/ZIP implementation
│   └── xml/                 # XML utilities, MC processing
│
└── testdata/
    └── golden/              # Expected XML for golden file tests
```

**Key Package Responsibilities:**
- `drawingml/` - Shared DrawingML types (a: namespace) used by all three SDKs for shapes, images, charts
- `packaging/` - OPC package layer for ZIP-based document handling
- `openxml/` - Core framework: element interfaces, part base types, validation
- `spreadsheet/elements/` - SpreadsheetML XML element types (x: namespace)
- `spreadsheet/parts/` - Spreadsheet part types (WorkbookPart, WorksheetPart, etc.)

## Compatibility Goals

- Full API parity with Open-XML-SDK for SpreadsheetML features
- Same document fidelity (round-trip without loss)
- Compatible content types and relationship types
- Support for Office 2016, 2019, 2021, and Microsoft 365 formats (ECMA-376 5th edition+)
- Interoperability with existing Word and Presentation SDKs in goffice

## Confirmed Design Decisions

These key decisions have been confirmed through analysis and discussion:

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **DrawingML** | Shared `drawingml/` package | Used by all three SDKs (Word, Presentation, Spreadsheet) for shapes, images, and charts |
| **Code Generation** | Pre-generate and commit | Go types generated from JSON schemas, committed to repository; users do not need the generator |
| **Phase 1 Scope** | Full feature parity with Open-XML-SDK | Including charts and pivot tables from the start |
| **Validation** | Strict by default | Reject invalid content, return errors for malformed documents |

## Unified Design Decisions

These decisions apply consistently across all three SDKs (Word, Presentation, Spreadsheet):

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Module Path** | `github.com/connerohnesorge/goffice` | Personal namespace, unified module |
| **Code Generation** | Pre-generate from JSON schemas, commit to repo | Reuse C# SDK schema data, 1000+ types per domain; users get ready-to-use types |
| **Office Version** | 2016+ only | Modern documents, reduced complexity (ECMA-376 5th edition+) |
| **Package Structure** | Domain-based packages | See updated architecture below |
| **API Naming** | Go-idiomatic short | `Workbook`, `Sheet`, `Cell` (not verbose C# names) |
| **Element Metadata** | Embedded struct fields | Self-contained elements, no global registry |
| **Construction Pattern** | Functional options | `NewSheet(WithName("Data"), WithIndex(0))` |
| **Child Access** | Generic methods only | `First[T]()`, `All[T]()`, `OfType[T]()` |

## Key SpreadsheetML Specifics

### Cell Reference System
```go
// A1-style references (most common)
cell := sheet.Cell("A1")
range := sheet.Range("A1:C10")

// R1C1-style references
cell := sheet.CellRC(1, 1)  // Row 1, Column 1

// Named ranges
sales := workbook.DefinedName("SalesData")
```

### Cell Value Types
```go
type CellType string
const (
    CellTypeBoolean      CellType = "b"
    CellTypeDate         CellType = "d"
    CellTypeError        CellType = "e"
    CellTypeInlineString CellType = "inlineStr"
    CellTypeNumber       CellType = "n"
    CellTypeSharedString CellType = "s"
    CellTypeString       CellType = "str"  // Formula string
)
```

### Shared String Table
Strings are deduplicated in a shared string table for efficiency:
```go
// Cell references index into shared string table
cell.DataType = CellTypeSharedString
cell.CellValue = "0"  // Index into shared string table

// SharedStringTablePart contains actual strings
sst.SharedStringItem[0].Text = "Hello, World!"
```

### Formula Types
```go
type FormulaType string
const (
    FormulaTypeNormal    FormulaType = "normal"
    FormulaTypeArray     FormulaType = "array"
    FormulaTypeDataTable FormulaType = "dataTable"
    FormulaTypeShared    FormulaType = "shared"
)
```

## Technical Findings from SpreadsheetML Exploration

These findings were gathered from analyzing the Open-XML-SDK source, ECMA-376 specification, and real Excel documents:

### Part Hierarchy

**WorkbookPart** is the root and contains:
- `WorksheetPart[]` - Individual worksheets (multiple)
- `SharedStringTablePart` - Shared string storage (singleton)
- `WorkbookStylesPart` - All style definitions (singleton)
- `CalculationChainPart` - Formula calculation order (singleton)
- `ThemePart` - Document theme (singleton)
- `ChartsheetPart[]` - Chart-only sheets (multiple)
- `PivotTableCacheDefinitionPart[]` - Pivot caches (multiple)

**WorksheetPart** contains:
- `DrawingsPart` - Embedded drawings/charts (singleton)
- `TableDefinitionPart[]` - Excel tables (multiple)
- `PivotTablePart[]` - Pivot tables (multiple)
- `CommentsPart` - Cell comments (singleton)

### Cell Reference System

- **Range**: A1 to XFD1048576 (maximum 16,384 columns x 1,048,576 rows)
- **Column naming**: A-Z, AA-AZ, ..., XFD (base-26 with offset)
- **Reference styles**: A1-style (default) and R1C1-style (for relative formulas)
- **Absolute references**: Use $ prefix ($A$1 = absolute, A$1 = mixed)

### Cell Data Types

| Type | Attribute Value | Description |
|------|-----------------|-------------|
| Number | (default/omit) | Numeric value stored in `<v>` element |
| Shared String | `s` | Index into SharedStringTable in `<v>` element |
| Formula String | `str` | Formula result is a string |
| Error | `e` | Error value (#DIV/0!, #VALUE!, #REF!, etc.) |
| Boolean | `b` | "0" or "1" in `<v>` element |
| Inline String | `inlineStr` | Rich text in `<is>` element instead of `<v>` |

### SharedStringTable

- Deduplicates strings across all cells in the workbook
- Cell stores integer index into the table
- Supports plain text (`<t>`) and rich text runs (`<r>`)
- `count` attribute = total references, `uniqueCount` = unique strings

### Stylesheet Structure

```
Stylesheet
+-- NumFmts[]           # Custom number formats (id >= 164)
+-- Fonts[]             # Font definitions (0-indexed)
+-- Fills[]             # Fill definitions (0-indexed, 0=none, 1=gray125)
+-- Borders[]           # Border definitions (0-indexed)
+-- CellStyleXfs[]      # Base formatting records for named styles
+-- CellXfs[]           # Cell formatting records (combines font/fill/border)
+-- CellStyles[]        # Named styles (Normal, Heading 1, etc.)
+-- Dxfs[]              # Differential formats (for conditional formatting)
+-- TableStyles[]       # Table style definitions
```

**Built-in Number Format IDs**: 0-163 are predefined (0 = General, 1 = 0, 2 = 0.00, etc.)
**Custom formats**: Start at ID 164

### Formula Types

| Type | Usage |
|------|-------|
| `normal` | Standard formula |
| `array` | Array formula (legacy CSE: Ctrl+Shift+Enter) |
| `dataTable` | Data table what-if analysis |
| `shared` | Shared formula (optimization for repeated patterns) |

### Conditional Formatting Rule Types

Supports 15+ rule types:
- `cellIs` - Compare cell value (equal, greaterThan, between, etc.)
- `colorScale` - 2-color or 3-color gradient based on value
- `dataBar` - In-cell data bar visualization
- `iconSet` - Icon sets (arrows, traffic lights, ratings, etc.)
- `top10` - Top/bottom N values or percentages
- `aboveAverage` - Above/below average values
- `duplicateValues` / `uniqueValues` - Duplicate/unique detection
- `containsText` / `notContainsText` / `beginsWith` / `endsWith` - Text rules
- `containsBlanks` / `notContainsBlanks` - Blank cell detection
- `containsErrors` / `notContainsErrors` - Error detection
- `expression` - Custom formula-based rules
- `timePeriod` - Date-based rules (today, thisWeek, lastMonth, etc.)

### Table Structure

Excel tables (ListObjects) have:
- `TableColumns` with optional calculated column formulas
- `TotalsRow` with aggregation functions (sum, average, count, etc.)
- `AutoFilter` for filtering
- `TableStyleInfo` referencing built-in or custom styles
- Support for structured references in formulas: `=SUM(Table1[Sales])`

### Chart Types

**2D Charts**: Bar, Line, Area, Pie, Doughnut, Radar, Scatter, Bubble, Stock
**3D Charts**: Bar3D, Line3D, Area3D, Pie3D, Surface3D
**Special**: OfPie (pie-of-pie, bar-of-pie), Stock (OHLC)

Charts use DrawingML (a: namespace) for:
- Shape properties (fills, lines, effects)
- Text formatting
- Themes and styles

### Pivot Tables (Two-Part System)

1. **PivotTableDefinition** - The pivot table structure
   - Row fields, column fields, page fields (filters)
   - Data fields with aggregation functions
   - Location and styling

2. **PivotCacheDefinition + PivotCacheRecords** - The data cache
   - Source data reference (worksheet range or external)
   - Field definitions and shared items
   - Cached record values for offline access

### Drawing Anchors

Drawings can be anchored to worksheets in three ways:

| Anchor Type | Usage |
|-------------|-------|
| `twoCellAnchor` | Anchored to two cells (from/to), resizes with cells |
| `oneCellAnchor` | Anchored to one cell with fixed extent |
| `absoluteAnchor` | Fixed position in EMUs (English Metric Units) |

### CellValue Format Notes

- **Numbers**: Stored using invariant culture format (decimal point, no thousands separator)
- **Dates**: Stored as Excel serial date numbers (days since 1899-12-30, with 1900 leap year bug)
- **Booleans**: Stored as "0" (false) or "1" (true)
- **Errors**: Stored as error codes (#DIV/0!, #VALUE!, #REF!, #NAME?, #NUM!, #N/A, #NULL!)

## Risk Assessment

| Risk | Impact | Mitigation |
|------|--------|------------|
| SpreadsheetML complexity | High - largest OOXML domain | Phased implementation; focus on common use cases first |
| Formula parsing | Medium - complex syntax | Store formulas as strings; defer calculation to Excel |
| Performance with large sheets | High - millions of cells possible | Lazy loading; streaming API for bulk operations |
| Style system complexity | Medium - many interdependencies | Follow Open-XML-SDK pattern exactly |
| Chart complexity | Medium - many chart types | Implement common types first; expand iteratively |
| Pivot table complexity | High - complex data model | Implement read-only first; add write support iteratively |
