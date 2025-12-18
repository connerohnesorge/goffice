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
| **Go Version** | 1.22+ |
| **Dependencies** | Pure stdlib only |
| **Generated Code** | Pre-generated from JSON schemas, committed |
| **API Style** | Direct struct field access |
| **Error Handling** | Return errors everywhere |
| **Validation** | Strict by default |
| **Thread Safety** | Package-level sync.Mutex |
| **MVP Scope** | Full read/write from v0.1 |
| **MC Handling** | Process automatically (select Choice/Fallback) |
| **Testing Strategy** | Golden file tests |
| **XML Prefixes** | Fixed canonical (x:, a:, r:, xdr:) |
| **Package Structure** | Flat by namespace (sml/, dml/, xdr/) |
| **Version Support** | Office 2016+ (FileFormatVersions.Office2016) |

## Architecture Overview

```
goffice/
├── pkg/
│   ├── framework/           # Core element types (shared)
│   ├── types/               # OpenXML value types (shared)
│   ├── package/             # OPC package handling (shared)
│   ├── validation/          # Schema/semantic validation (shared)
│   │
│   ├── spreadsheet/         # SpreadsheetDocument & types
│   │   ├── document.go      # SpreadsheetDocument
│   │   ├── types.go         # SpreadsheetDocumentType enum
│   │   └── parts/           # All spreadsheet part types
│   │       ├── workbook.go  # WorkbookPart
│   │       ├── worksheet.go # WorksheetPart
│   │       ├── styles.go    # WorkbookStylesPart
│   │       ├── strings.go   # SharedStringTablePart
│   │       ├── chart.go     # ChartPart
│   │       └── ...          # All other parts
│   │
│   ├── sml/                 # SpreadsheetML elements (x: namespace)
│   │   └── *.go             # Pre-generated from JSON schemas
│   │
│   ├── xdr/                 # SpreadsheetDrawing elements
│   │   └── *.go             # Pre-generated from JSON schemas
│   │
│   ├── dml/                 # DrawingML elements (shared with pres/word)
│   ├── word/                # WordprocessingML (existing)
│   └── presentation/        # PresentationML (existing)
│
├── internal/
│   ├── opc/                 # Low-level OPC/ZIP implementation
│   └── xml/                 # XML utilities, MC processing
│
└── testdata/
    └── golden/              # Expected XML for golden file tests
```

## Compatibility Goals

- Full API parity with Open-XML-SDK for SpreadsheetML features
- Same document fidelity (round-trip without loss)
- Compatible content types and relationship types
- Support for Office 2016, 2019, 2021, and Microsoft 365 formats (ECMA-376 5th edition+)
- Interoperability with existing Word and Presentation SDKs in goffice

## Unified Design Decisions

These decisions apply consistently across all three SDKs (Word, Presentation, Spreadsheet):

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Module Path** | `github.com/connerohnesorge/goffice` | Personal namespace, unified module |
| **Code Generation** | Generate from JSON schemas | Reuse C# SDK schema data, 1000+ types per domain |
| **Office Version** | 2016+ only | Modern documents, reduced complexity (ECMA-376 5th edition+) |
| **Package Structure** | `pkg/` style | `pkg/{framework,types,package,spreadsheet,sml}` |
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

## Risk Assessment

| Risk | Impact | Mitigation |
|------|--------|------------|
| SpreadsheetML complexity | High - largest OOXML domain | Phased implementation; focus on common use cases first |
| Formula parsing | Medium - complex syntax | Store formulas as strings; defer calculation to Excel |
| Performance with large sheets | High - millions of cells possible | Lazy loading; streaming API for bulk operations |
| Style system complexity | Medium - many interdependencies | Follow Open-XML-SDK pattern exactly |
| Chart complexity | Medium - many chart types | Implement common types first; expand iteratively |
| Pivot table complexity | High - complex data model | Implement read-only first; add write support iteratively |
