# Design: Go SDK for Office Open XML Spreadsheets

## Context

This design document captures architectural decisions for implementing SpreadsheetML support in the goffice Go SDK, modeled after Microsoft's Open-XML-SDK for C#. The spreadsheet SDK builds upon the shared framework established by the Word and Presentation SDKs, while addressing domain-specific complexities unique to Excel workbooks.

**Stakeholders:**
- Go developers needing Excel file generation/manipulation
- Enterprise applications requiring server-side spreadsheet processing
- Data pipeline applications that produce Excel reports
- Migration from .NET Open-XML-SDK to Go

**Constraints:**
- Must support ECMA-376 and ISO/IEC 29500 standards
- Target Office 2016+ documents (ECMA-376 5th edition and later)
- Must work without Microsoft Excel installed
- Standard library only for core functionality (no CGO)
- Must share common components with existing Word and Presentation SDKs

**Decided Design Choices (Unified with Word/Presentation SDKs):**
- Code generation from JSON schemas (same source as C# SDK)
- Idiomatic Go naming conventions (not mirroring C# names)
- Full validation: schema + all semantic constraint types
- Office 2016+ only (reduces version complexity)
- Functional options pattern for element construction
- Generic methods only for child access: `First[T]()`, `All[T]()`, `OfType[T]()`

## Goals / Non-Goals

### Goals
- Feature parity with Open-XML-SDK for SpreadsheetML
- Idiomatic Go API design (interfaces, composition, error handling)
- Strong typing for all elements and attributes
- Comprehensive validation (schema + semantic)
- Support all Excel document types (.xlsx, .xltx, .xlsm, .xltm, .xlam)
- Lazy loading for performance with large workbooks
- Thread-safe read operations
- Efficient handling of large worksheets (millions of cells)
- Shared string table management for string deduplication
- Formula storage (not evaluation)
- Full cell style support

### Non-Goals
- Formula evaluation engine (Excel calculates when opening)
- Spreadsheet rendering/preview generation
- PDF conversion
- VBA macro execution
- Real-time collaborative editing support
- External data source connections (refresh)

## Unified Design Decisions (Cross-SDK)

These decisions apply to all three SDKs (Word, Presentation, Spreadsheet) for consistency:

| Decision | Choice |
|----------|--------|
| **Module Path** | `github.com/connerohnesorge/goffice` |
| **Code Generation** | Generate from JSON schemas (pre-generated, committed) |
| **Office Version** | 2016+ only (ECMA-376 5th edition+) |
| **Package Structure** | `pkg/{framework,simpletypes,packaging,spreadsheet,...}/` |
| **API Naming** | Go-idiomatic short names (`Workbook`, `Sheet`, not verbose C# names) |
| **Element Metadata** | Embedded struct fields (no global registry) |
| **Construction Pattern** | Functional options: `NewSheet(WithName("Data"))` |
| **Child Access** | Generic methods only: `First[T]()`, `All[T]()`, `OfType[T]()` |

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                        User Application                          │
└─────────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────────┐
│                     spreadsheet package                          │
│  ┌─────────────────┐  ┌──────────────┐  ┌───────────────────┐  │
│  │ Document        │  │ Parts        │  │ Helpers           │  │
│  │ - New()         │  │ - Workbook   │  │ - CellRef         │  │
│  │ - Open()        │  │ - Worksheet  │  │ - RangeRef        │  │
│  │ - Save()        │  │ - Styles     │  │ - FormulaBuilder  │  │
│  │ - Close()       │  │ - Strings    │  │ - StyleBuilder    │  │
│  └─────────────────┘  └──────────────┘  └───────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────────┐
│                        sml package                               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐  │
│  │ Workbook     │  │ Worksheet    │  │ Cell/Row/Col         │  │
│  │ Sheets       │  │ SheetData    │  │ Cell, CellValue      │  │
│  │ DefinedNames │  │ SheetViews   │  │ Row, Col, Cols       │  │
│  │ CalcPr       │  │ MergeCells   │  │ CellFormula          │  │
│  └──────────────┘  └──────────────┘  └──────────────────────┘  │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐  │
│  │ Styles       │  │ SharedStrings│  │ Tables/Pivot         │  │
│  │ CellXfs      │  │ SharedString │  │ Table, PivotTable    │  │
│  │ Fonts, Fills │  │ Text, Run    │  │ PivotCache           │  │
│  │ Borders      │  │ PhoneticRun  │  │ PivotField           │  │
│  └──────────────┘  └──────────────┘  └──────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────────┐
│                     framework package (shared)                   │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐  │
│  │ Element      │  │ Part         │  │ Validation           │  │
│  │ - Parent     │  │ - URI        │  │ - SchemaValidator    │  │
│  │ - Children   │  │ - ContentType│  │ - SemanticValidator  │  │
│  │ - Attributes │  │ - Root       │  │ - Constraints        │  │
│  │ - Clone()    │  │ - Stream     │  │                      │  │
│  └──────────────┘  └──────────────┘  └──────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────────┐
│                     packaging package (shared)                   │
└─────────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────────┐
│                     archive/zip (stdlib)                         │
└─────────────────────────────────────────────────────────────────┘
```

## Spreadsheet-Specific Decisions

### D1: Cell Reference System

**Decision:** Implement both A1-style and R1C1-style cell references with utilities.

```go
// CellRef represents a cell reference
type CellRef struct {
    Col    uint32  // 1-based column (1 = A)
    Row    uint32  // 1-based row
    AbsCol bool    // $A vs A
    AbsRow bool    // $1 vs 1
}

// Parse A1-style reference
func ParseCellRef(ref string) (CellRef, error)

// Create A1-style string
func (c CellRef) String() string  // "A1", "$A$1", etc.

// Create R1C1-style string
func (c CellRef) R1C1() string    // "R1C1", "R[1]C[1]", etc.

// RangeRef represents a cell range
type RangeRef struct {
    Start CellRef
    End   CellRef
}

// Parse range reference
func ParseRangeRef(ref string) (RangeRef, error)  // "A1:C10"

// Column utilities
func ColumnName(col uint32) string      // 1 -> "A", 27 -> "AA"
func ColumnIndex(name string) uint32    // "A" -> 1, "AA" -> 27
```

**Rationale:** Cell references are fundamental to spreadsheet operations. A1-style is most common, but R1C1 is needed for relative formulas.

### D2: Shared String Table Management

**Decision:** Automatic shared string management with explicit opt-out for inline strings.

```go
// Cell automatically uses shared strings for string values
cell := sheet.Cell("A1")
cell.SetString("Hello")  // Automatically added to shared string table

// Explicit inline string (rare, for special cases)
cell.SetInlineString("Inline text")

// SharedStringTable is managed by WorkbookPart
sst := workbook.SharedStringTablePart()
index := sst.AddString("New string")  // Returns index
text := sst.String(index)              // Returns string by index
```

**Rationale:** Shared strings reduce file size significantly for repeated strings. The SDK should handle this transparently while allowing explicit control when needed.

### D3: Cell Style System

**Decision:** Style builder pattern with style registry for efficient reuse.

```go
// StyleRegistry manages styles to avoid duplication
type StyleRegistry struct {
    numberFormats map[string]uint32
    fonts         map[fontKey]uint32
    fills         map[fillKey]uint32
    borders       map[borderKey]uint32
    cellXfs       map[xfKey]uint32
}

// Style builder for fluent API
style := spreadsheet.NewStyle().
    Font(spreadsheet.FontBold()).
    Fill(spreadsheet.FillSolid("#FF0000")).
    Border(spreadsheet.BorderThin()).
    NumberFormat("$#,##0.00").
    Alignment(spreadsheet.AlignCenter())

// Apply to cell
cell.SetStyle(style)

// Styles are automatically registered and deduplicated
```

**Rationale:** Excel's style system has complex interdependencies (fonts, fills, borders are separate tables referenced by cell XFs). The builder pattern hides this complexity.

### D4: Formula Storage Strategy

**Decision:** Store formulas as strings without parsing or evaluation.

```go
// Store formula as-is
cell.SetFormula("=SUM(A1:A10)")
cell.SetFormula("=VLOOKUP(A1,Sheet2!A:B,2,FALSE)")

// Array formulas (Ctrl+Shift+Enter style)
cell.SetArrayFormula("=A1:A10*B1:B10", "C1:C10")  // formula, range

// Dynamic array formulas (Office 365+)
cell.SetFormula("=SORT(A1:A10)")  // Spills automatically

// Shared formulas (optimization for repeated formulas)
sheet.SetSharedFormula("=A1+1", "B1:B100")  // base formula, range

// Defined names
workbook.AddDefinedName("SalesData", "Sheet1!$A$1:$D$100")
```

**Rationale:** Formula evaluation is complex (1000+ functions) and Excel handles it on open. Storing formulas without evaluation avoids massive complexity while providing full functionality.

### D5: Large Worksheet Handling

**Decision:** Streaming row writer for large data sets; lazy loading for reading.

```go
// Writing large data sets - streaming API
writer := sheet.NewRowWriter()
for i := 0; i < 1_000_000; i++ {
    row := writer.AddRow()
    row.AddCell().SetNumber(float64(i))
    row.AddCell().SetString("Data")
}
writer.Close()  // Flushes to part

// Reading large worksheets - lazy loading
// Rows are not loaded until accessed
for row := sheet.FirstRow(); row != nil; row = row.NextRow() {
    // Process row
}

// Direct cell access (loads only needed portion)
value := sheet.Cell("Z1000000").Value()
```

**Rationale:** Excel worksheets can have 1,048,576 rows × 16,384 columns. Loading everything into memory is impractical for large sheets.

### D6: Chart Implementation

**Decision:** Support common chart types with DrawingML integration.

```go
// Create chart
chart := sheet.AddChart(spreadsheet.ChartTypeColumn, "A1:B10")
chart.SetTitle("Sales by Month")
chart.SetPosition("D1", "J15")  // Anchor positions

// Configure series
series := chart.Series(0)
series.SetName("Sales")
series.SetCategories("Sheet1!$A$2:$A$10")
series.SetValues("Sheet1!$B$2:$B$10")

// Add axis labels
chart.CategoryAxis().SetTitle("Month")
chart.ValueAxis().SetTitle("Amount")

// Sparklines (mini-charts in cells)
sheet.AddSparkline("C1:C10", "D1:D10", spreadsheet.SparklineTypeLine)
```

**Rationale:** Charts are a key Excel feature. The API should be simple for common cases while allowing full customization.

### D7: Pivot Table Implementation

**Decision:** Full pivot table support with separate cache management.

```go
// Create pivot table from data range
pivot := sheet.AddPivotTable("Sheet1!A1:D100", "Sheet2!A1")

// Configure fields
pivot.AddRowField("Category")
pivot.AddColumnField("Region")
pivot.AddDataField("Sales", spreadsheet.AggregateSum)
pivot.AddDataField("Units", spreadsheet.AggregateCount)

// Add filter
pivot.AddPageField("Year")

// Calculated field
pivot.AddCalculatedField("Profit", "=Sales-Costs")

// Pivot cache is managed automatically
cache := pivot.Cache()
cache.RefreshOnLoad(true)
```

**Rationale:** Pivot tables are complex but essential. Separate cache management follows OOXML architecture and enables advanced scenarios.

### D8: Table Support

**Decision:** Excel tables (ListObject) as first-class citizens.

```go
// Create table from range
table := sheet.AddTable("A1:D10", "SalesTable")
table.SetTotalsRowShown(true)
table.SetFilterButtonShown(true)

// Access columns
col := table.Column("Sales")
col.SetTotalsRowFunction(spreadsheet.TotalsSum)
col.SetCalculatedColumnFormula("=[Units]*[Price]")

// Table style
table.SetStyle(spreadsheet.TableStyleMedium2)

// Structured references in formulas
cell.SetFormula("=SUM(SalesTable[Sales])")
```

**Rationale:** Tables provide powerful features (structured references, auto-expand, totals row) that are heavily used.

### D9: Conditional Formatting

**Decision:** Support all conditional formatting types with fluent builder.

```go
// Cell value rules
sheet.AddConditionalFormatting("A1:A100",
    spreadsheet.CF().
        CellValue(spreadsheet.GreaterThan, "100").
        Style(spreadsheet.NewStyle().Fill(spreadsheet.FillSolid("#00FF00"))))

// Color scale
sheet.AddConditionalFormatting("B1:B100",
    spreadsheet.CF().
        ColorScale(
            spreadsheet.MinValue(), "#FF0000",
            spreadsheet.Percentile(50), "#FFFF00",
            spreadsheet.MaxValue(), "#00FF00"))

// Data bars
sheet.AddConditionalFormatting("C1:C100",
    spreadsheet.CF().
        DataBar("#638EC6"))

// Icon sets
sheet.AddConditionalFormatting("D1:D100",
    spreadsheet.CF().
        IconSet(spreadsheet.IconSetTrafficLights))

// Formula-based
sheet.AddConditionalFormatting("E1:E100",
    spreadsheet.CF().
        Formula("=$A1>$B1").
        Style(spreadsheet.NewStyle().Font(spreadsheet.FontBold())))
```

**Rationale:** Conditional formatting is widely used for data visualization. The builder pattern handles the complexity of different rule types.

### D10: Data Validation

**Decision:** Support all validation types with error/input messages.

```go
// List validation (dropdown)
sheet.AddDataValidation("A1:A100",
    spreadsheet.DV().
        List("Option1,Option2,Option3").
        ShowDropdown(true))

// Whole number validation
sheet.AddDataValidation("B1:B100",
    spreadsheet.DV().
        WholeNumber(spreadsheet.Between, 1, 100).
        ErrorTitle("Invalid Entry").
        ErrorMessage("Please enter a number between 1 and 100"))

// Date validation
sheet.AddDataValidation("C1:C100",
    spreadsheet.DV().
        Date(spreadsheet.GreaterThan, time.Now()))

// Custom formula validation
sheet.AddDataValidation("D1:D100",
    spreadsheet.DV().
        Formula("=LEN(D1)<=50"))
```

**Rationale:** Data validation ensures data integrity. The builder pattern provides a clean API for various validation types.

## Risks / Trade-offs

| Risk | Impact | Mitigation |
|------|--------|------------|
| SpreadsheetML size | High - largest OOXML domain (~800 element types) | Phase implementation; prioritize common features |
| Formula complexity | Medium - 500+ functions | Store only, don't evaluate; Excel recalculates |
| Large file performance | High - millions of cells possible | Streaming API; lazy loading; careful memory management |
| Style system complexity | Medium - many cross-references | Builder pattern; automatic deduplication |
| Chart/Drawing complexity | Medium - multiple namespaces | DrawingML shared with other SDKs |
| Pivot table complexity | High - cache + table + fields | Implement read-only first; add write support iteratively |

## Migration Plan

N/A - New capability with no existing implementation.

## Resolved Design Questions

### Q1: How to handle the shared string table?
**Decision:** Automatic management with transparent API.

Users set string values on cells. The SDK automatically:
1. Checks if string exists in shared string table
2. Adds new strings if needed
3. Updates cell to reference the shared string index

This matches how the C# SDK works and provides the best file size optimization.

### Q2: How to handle number formats?
**Decision:** Support both built-in and custom formats.

```go
// Built-in format (by ID)
cell.SetNumberFormatId(2)  // "0.00"

// Custom format
cell.SetNumberFormat("$#,##0.00")

// Date format
cell.SetNumberFormat("yyyy-mm-dd")
```

Built-in format IDs (1-49) are defined by ECMA-376. Custom formats are added to the styles part.

### Q3: How to handle cell references in APIs?
**Decision:** Accept string references with parsing and validation.

```go
// String references (validated)
cell := sheet.Cell("A1")        // Single cell
range := sheet.Range("A1:C10")  // Range

// Programmatic references (no parsing needed)
cell := sheet.CellAt(1, 1)      // Row 1, Col 1
range := sheet.RangeAt(1, 1, 10, 3)  // R1C1:R10C3
```

This provides both user-friendly string APIs and efficient programmatic access.

### Q4: How to handle merged cells?
**Decision:** Explicit merge API with validation.

```go
// Merge cells
sheet.MergeCells("A1:D1")

// Value goes in top-left cell
sheet.Cell("A1").SetString("Merged Header")

// Check if cell is in merge
if merge := sheet.GetMerge("B1"); merge != nil {
    fmt.Println("Part of merge:", merge.Reference())  // "A1:D1"
}

// Unmerge
sheet.UnmergeCells("A1:D1")
```

### Q5: How to handle worksheet protection?
**Decision:** Full protection support with granular permissions.

```go
// Protect sheet with options
sheet.Protect(spreadsheet.Protection().
    Password("secret").
    AllowSelectLockedCells(true).
    AllowSelectUnlockedCells(true).
    AllowFormatCells(false).
    AllowInsertRows(false))

// Protect workbook structure
workbook.Protect(spreadsheet.WorkbookProtection().
    StructureProtected(true))

// Cell-level protection (requires sheet protection)
cell.SetLocked(true)    // Locked (default)
cell.SetLocked(false)   // Unlocked (editable when protected)
cell.SetHidden(true)    // Formula hidden when protected
```

## SpreadsheetML-Specific Type Mappings

### C# to Go Type Mapping for Spreadsheet Domain

| C# Type | Go Type | Notes |
|---------|---------|-------|
| `SpreadsheetDocument` | `Document` | Main entry point |
| `WorkbookPart` | `WorkbookPart` | Contains workbook root |
| `WorksheetPart` | `WorksheetPart` | Contains worksheet root |
| `SharedStringTablePart` | `SharedStringsPart` | String deduplication |
| `WorkbookStylesPart` | `StylesPart` | All style definitions |
| `Cell` | `Cell` | Individual cell |
| `CellValue` | `CellValue` | Cell's value element |
| `CellFormula` | `Formula` | Cell's formula |
| `Row` | `Row` | Row container |
| `SheetData` | `SheetData` | All rows container |
| `DefinedName` | `DefinedName` | Named range/formula |
| `PivotTablePart` | `PivotTablePart` | Pivot table definition |

### Namespace Mapping

| XML Prefix | Namespace | Go Package |
|------------|-----------|------------|
| `x` | `http://schemas.openxmlformats.org/spreadsheetml/2006/main` | `sml` |
| `xdr` | `http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing` | `xdr` |
| `a` | `http://schemas.openxmlformats.org/drawingml/2006/main` | `dml` (shared) |
| `r` | `http://schemas.openxmlformats.org/officeDocument/2006/relationships` | `rml` (shared) |
| `c` | `http://schemas.openxmlformats.org/drawingml/2006/chart` | `dml/chart` |

## Key Findings from SDK Analysis

### Part Hierarchy (SpreadsheetML-specific)

WorkbookPart is the root with these key children:
- `WorksheetPart[]` - Individual worksheets (multiple)
- `SharedStringTablePart` - Shared strings (singleton)
- `WorkbookStylesPart` - All styles (singleton)
- `CalculationChainPart` - Calculation order (singleton)
- `ThemePart` - Document theme (singleton)
- `ConnectionsPart` - External connections (singleton)
- `ChartsheetPart[]` - Chart-only sheets (multiple)
- `PivotTableCacheDefinitionPart[]` - Pivot caches (multiple)

WorksheetPart has these children:
- `DrawingsPart` - Embedded drawings/charts (singleton)
- `TableDefinitionPart[]` - Excel tables (multiple)
- `PivotTablePart[]` - Pivot tables (multiple)
- `CommentsPart` - Cell comments (singleton)
- `VmlDrawingPart[]` - Legacy drawings (multiple)

### Cell Value Storage Patterns

1. **Numeric values**: Stored directly in `<v>` element
2. **Boolean values**: "0" or "1" in `<v>` with `t="b"`
3. **Shared strings**: Index in `<v>` with `t="s"`
4. **Inline strings**: `<is>` element instead of `<v>`
5. **Formulas**: `<f>` element with result in `<v>`
6. **Dates**: Stored as numbers (days since 1900) with date format
7. **Errors**: Error code in `<v>` with `t="e"`

### Style System Structure

```
StylesPart
├── NumberFormats[]     - Custom number formats (id >= 164)
├── Fonts[]            - Font definitions (index-based)
├── Fills[]            - Fill definitions (index-based)
├── Borders[]          - Border definitions (index-based)
├── CellStyleXfs[]     - Named style formatting records
├── CellXfs[]          - Cell formatting records (combines above)
├── CellStyles[]       - Named styles (Normal, Heading 1, etc.)
├── DxFs[]             - Differential formats (conditional formatting)
└── TableStyles[]      - Table style definitions
```

Cell references a CellXf by index. CellXf references Font, Fill, Border by their indices.

## Appendix: SpreadsheetML Element Count

Based on JSON schema analysis:
- Core spreadsheet elements (`x:` namespace): ~400 types
- Spreadsheet drawing elements (`xdr:` namespace): ~30 types
- Chart elements (`c:` namespace): ~200 types
- Shared DrawingML elements (`a:` namespace): ~300 types (shared with other SDKs)
- Office 2010+ extensions: ~100 types
- Office 2016+ extensions: ~50 types

Total unique to spreadsheet domain: ~600+ element types
