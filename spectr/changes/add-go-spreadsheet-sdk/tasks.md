# Implementation Tasks: Go SDK for Spreadsheet Document Processing

**Unified Design Decisions Applied (consistent with Word and Presentation SDKs):**
- Module path: `github.com/connerohnesorge/goffice`
- Code generation from JSON schemas (pre-generated, committed)
- Office 2016+ only (ECMA-376 5th edition+)
- Package structure: `pkg/{framework,types,package,spreadsheet,sml,xdr}/`
- Go-idiomatic short API names (`Workbook`, `Sheet`, `Cell`)
- Embedded struct fields for element metadata
- Functional options pattern for element construction
- Generic methods only for child access: `First[T]()`, `All[T]()`, `OfType[T]()`
- *T pointers for optional/nullable values
- Priority: Core → Cells/Styles → Tables → Charts → Pivot Tables → Advanced

## Phase 1: Project Foundation (Shared with Word/Presentation)

### 1.1 Project Setup (if not already done)
- [ ] 1.1.1 Initialize Go module `github.com/connerohnesorge/goffice`
- [ ] 1.1.2 Create directory structure: `pkg/`, `internal/`, `cmd/`, `testdata/`
- [ ] 1.1.3 Set up testing infrastructure with `go test` and test fixtures directory
- [ ] 1.1.4 Create sample .xlsx test files for roundtrip testing (Office 2016+)
- [ ] 1.1.5 Add Makefile with build, test, lint targets

### 1.2 Core Simple Types (pkg/types/) - Shared
- [ ] 1.2.1 Implement `SimpleValue` interface with `HasValue()`, `InnerText()`, `SetInnerText()`
- [ ] 1.2.2 Implement `StringValue` type
- [ ] 1.2.3 Implement `Int32Value` and `UInt32Value` types
- [ ] 1.2.4 Implement `Int64Value` and `UInt64Value` types
- [ ] 1.2.5 Implement `BooleanValue` type (true/false and 1/0)
- [ ] 1.2.6 Implement `DoubleValue` type
- [ ] 1.2.7 Implement `EnumValue[T]` generic type for enumerations
- [ ] 1.2.8 Implement `HexBinaryValue` type
- [ ] 1.2.9 Implement `Base64BinaryValue` type
- [ ] 1.2.10 Implement `DateTimeValue` type (ISO 8601)
- [ ] 1.2.11 Implement `DecimalValue` type
- [ ] 1.2.12 Write comprehensive tests for all simple types

## Phase 2: OPC Packaging Layer (Shared with Word/Presentation)

### 2.1 Core Package Types (internal/opc/ and pkg/package/)
- [ ] 2.1.1 Implement `Package` struct with ZIP backing via `archive/zip`
- [ ] 2.1.2 Implement `New(path)` and `NewWriter(w)` for new packages
- [ ] 2.1.3 Implement `Open(path, readOnly)` for existing packages
- [ ] 2.1.4 Implement `OpenReader(r io.ReaderAt, size)` for stream-based open
- [ ] 2.1.5 Implement `Package.Save()` for in-place save
- [ ] 2.1.6 Implement `Package.SaveAs(path)` for save-to-new-location
- [ ] 2.1.7 Implement `Package.Close()` with resource cleanup
- [ ] 2.1.8 Implement package capabilities (Read, Write, ReadWrite)

### 2.2 Content Types Management
- [ ] 2.2.1 Implement `ContentTypes` struct for [Content_Types].xml
- [ ] 2.2.2 Implement default content type registration by extension
- [ ] 2.2.3 Implement override content type for specific URIs
- [ ] 2.2.4 Implement content type lookup by part URI
- [ ] 2.2.5 Implement serialization/deserialization of [Content_Types].xml
- [ ] 2.2.6 Write tests for content types management

### 2.3 Package Parts
- [ ] 2.3.1 Implement `Part` struct with URI, ContentType, Stream
- [ ] 2.3.2 Implement `Package.CreatePart(uri, contentType)`
- [ ] 2.3.3 Implement `Package.Part(uri)` accessor
- [ ] 2.3.4 Implement `Package.DeletePart(uri)`
- [ ] 2.3.5 Implement `Package.Parts()` enumeration
- [ ] 2.3.6 Implement `Part.Stream(mode)` for part content access
- [ ] 2.3.7 Implement part URI normalization and validation
- [ ] 2.3.8 Write tests for package parts

### 2.4 Package Relationships
- [ ] 2.4.1 Implement `Relationship` struct
- [ ] 2.4.2 Implement `Package.CreateRel(target, relType, id)`
- [ ] 2.4.3 Implement `Package.RelsByType(relType)` accessor
- [ ] 2.4.4 Implement `Package.Rel(id)` accessor
- [ ] 2.4.5 Implement `Package.DeleteRel(id)`
- [ ] 2.4.6 Implement `_rels/.rels` serialization/deserialization
- [ ] 2.4.7 Implement part-level `.rels` file handling
- [ ] 2.4.8 Write tests for package relationships

## Phase 3: OpenXML Framework (Shared with Word/Presentation)

### 3.1 Feature Collection System
- [ ] 3.1.1 Implement `FeatureCollection` struct with parent inheritance
- [ ] 3.1.2 Implement `Get[T]()` with parent chain traversal
- [ ] 3.1.3 Implement `Set(feature)` for registration
- [ ] 3.1.4 Implement thread-safe access (sync.Mutex)
- [ ] 3.1.5 Define `IPackageFeature` interface and implementation
- [ ] 3.1.6 Define `IContentTypeFeature` interface and implementation
- [ ] 3.1.7 Define `INamespaceFeature` interface and implementation
- [ ] 3.1.8 Write tests for feature collection

### 3.2 Element Base Types
- [ ] 3.2.1 Implement `Element` interface with core methods
- [ ] 3.2.2 Implement `elementBase` struct with common fields
- [ ] 3.2.3 Implement `CompositeElement` interface and `compositeBase` struct
- [ ] 3.2.4 Implement `LeafElement` interface and `leafBase` struct
- [ ] 3.2.5 Implement `OpenXmlQualifiedName` for namespace:localname
- [ ] 3.2.6 Implement `Attribute` struct and attribute management
- [ ] 3.2.7 Write tests for element base types

### 3.3 Element Tree Operations
- [ ] 3.3.1 Implement `AppendChild()`, `PrependChild()` operations
- [ ] 3.3.2 Implement `InsertBefore()`, `InsertAfter()` operations
- [ ] 3.3.3 Implement `RemoveChild()`, `RemoveAllChildren()` operations
- [ ] 3.3.4 Implement `FirstChild()`, `LastChild()` accessors
- [ ] 3.3.5 Implement `NextSibling()`, `PreviousSibling()` navigation
- [ ] 3.3.6 Implement `Parent()` accessor and parent tracking
- [ ] 3.3.7 Implement `Children()` and `Elements[T]()` enumeration
- [ ] 3.3.8 Write tests for tree operations

### 3.4 XML Serialization
- [ ] 3.4.1 Implement `XMLWriter` for namespace-aware serialization
- [ ] 3.4.2 Implement `Element.WriteTo(io.Writer)` serialization
- [ ] 3.4.3 Implement `Element.OuterXml()` and `InnerXml()`
- [ ] 3.4.4 Implement `ParseElement(io.Reader)` for XML parsing
- [ ] 3.4.5 Implement namespace prefix management (x:, a:, r:, xdr:)
- [ ] 3.4.6 Write tests for XML serialization with roundtrip validation

### 3.5 Part System
- [ ] 3.5.1 Implement `OpenXmlPart` base struct
- [ ] 3.5.2 Implement lazy loading of part root element
- [ ] 3.5.3 Implement `OpenXmlPartContainer` for child part management
- [ ] 3.5.4 Implement `GetPartsOfType[T]()` generic accessor
- [ ] 3.5.5 Implement `AddPart()`, `DeletePart()` operations
- [ ] 3.5.6 Implement part type registry by content type
- [ ] 3.5.7 Write tests for part system

## Phase 4: Validation Framework (Shared with Word/Presentation)

### 4.1 Validation Infrastructure
- [ ] 4.1.1 Implement `Validator` interface
- [ ] 4.1.2 Implement `ValidationContext` with version and settings
- [ ] 4.1.3 Implement `ValidationError` with path, description, severity
- [ ] 4.1.4 Implement `ValidationSettings` (MaxErrors, ContinueOnError)
- [ ] 4.1.5 Define `FileFormatVersions` enum (Office2016, Office2019, Office2021, Microsoft365)
- [ ] 4.1.6 Write tests for validation infrastructure

### 4.2 Schema Validation
- [ ] 4.2.1 Implement `SchemaValidator` for structure validation
- [ ] 4.2.2 Implement `CompiledParticle` for schema particles
- [ ] 4.2.3 Implement sequence particle validation
- [ ] 4.2.4 Implement choice particle validation
- [ ] 4.2.5 Implement all particle validation
- [ ] 4.2.6 Implement minOccurs/maxOccurs validation
- [ ] 4.2.7 Write tests for schema validation

### 4.3 Semantic Validation
- [ ] 4.3.1 Implement `SemanticValidator` base
- [ ] 4.3.2 Implement `Constraint` interface
- [ ] 4.3.3 Implement `AttributeValueRangeConstraint`
- [ ] 4.3.4 Implement `AttributeValuePatternConstraint`
- [ ] 4.3.5 Implement `UniqueValueConstraint`
- [ ] 4.3.6 Implement `RelationshipExistConstraint`
- [ ] 4.3.7 Implement `ReferenceExistConstraint`
- [ ] 4.3.8 Write comprehensive tests for semantic constraints

## Phase 5: Spreadsheet Document (pkg/spreadsheet/)

### 5.1 Document Type
- [ ] 5.1.1 Implement `Document` struct (SpreadsheetDocument equivalent)
- [ ] 5.1.2 Implement `DocType` enum (Workbook, Template, MacroEnabledWorkbook, MacroEnabledTemplate, AddIn)
- [ ] 5.1.3 Implement `Create(path, docType)` for new documents
- [ ] 5.1.4 Implement `CreateFromStream(w, docType)` for stream creation
- [ ] 5.1.5 Implement `Open(path, editable)` for existing documents
- [ ] 5.1.6 Implement `OpenFromStream(r, size, editable)` for streams
- [ ] 5.1.7 Implement `OpenOptions` struct for configuration
- [ ] 5.1.8 Implement `CreateFromTemplate(templatePath)`
- [ ] 5.1.9 Implement `Save()`, `SaveAs(path)`, `SaveTo(w)`
- [ ] 5.1.10 Implement `Close()` with io.Closer interface
- [ ] 5.1.11 Implement `ChangeType(newType)`
- [ ] 5.1.12 Implement document-level `Validate(version)`
- [ ] 5.1.13 Write integration tests with real .xlsx files (Office 2016+)

### 5.2 Document Parts
- [ ] 5.2.1 Implement `WorkbookPart` with content type detection
- [ ] 5.2.2 Implement `WorksheetPart`
- [ ] 5.2.3 Implement `SharedStringTablePart`
- [ ] 5.2.4 Implement `WorkbookStylesPart`
- [ ] 5.2.5 Implement `CalculationChainPart`
- [ ] 5.2.6 Implement `ConnectionsPart`
- [ ] 5.2.7 Implement `ThemePart`
- [ ] 5.2.8 Implement `ChartsheetPart`
- [ ] 5.2.9 Implement `DrawingsPart`
- [ ] 5.2.10 Implement `ChartPart`
- [ ] 5.2.11 Implement `TableDefinitionPart`
- [ ] 5.2.12 Implement `PivotTablePart`
- [ ] 5.2.13 Implement `PivotTableCacheDefinitionPart`
- [ ] 5.2.14 Implement `PivotTableCacheRecordsPart`
- [ ] 5.2.15 Implement `WorksheetCommentsPart`
- [ ] 5.2.16 Implement `VmlDrawingPart`
- [ ] 5.2.17 Implement `QueryTablePart`
- [ ] 5.2.18 Implement `SlicerPart` and `SlicerCachePart`
- [ ] 5.2.19 Implement `TimeLinePart` and `TimeLineCachePart`
- [ ] 5.2.20 Implement `ExternalWorkbookPart`
- [ ] 5.2.21 Implement `VbaProjectPart` for macro-enabled documents
- [ ] 5.2.22 Implement `ImagePart` with content type detection
- [ ] 5.2.23 Write tests for all part types

## Phase 6: SpreadsheetML Elements (pkg/sml/)

### 6.1 Workbook Elements
- [ ] 6.1.1 Implement `Workbook` root element (x:workbook)
- [ ] 6.1.2 Implement `FileVersion` element
- [ ] 6.1.3 Implement `WorkbookPr` element (workbook properties)
- [ ] 6.1.4 Implement `BookViews` and `WorkbookView` elements
- [ ] 6.1.5 Implement `Sheets` and `Sheet` elements
- [ ] 6.1.6 Implement `DefinedNames` and `DefinedName` elements
- [ ] 6.1.7 Implement `CalcPr` element (calculation properties)
- [ ] 6.1.8 Implement `PivotCaches` and `PivotCache` elements
- [ ] 6.1.9 Implement `WebPublishing` element
- [ ] 6.1.10 Implement `FileRecoveryPr` element
- [ ] 6.1.11 Implement `WorkbookProtection` element
- [ ] 6.1.12 Write tests for workbook elements

### 6.2 Worksheet Elements
- [ ] 6.2.1 Implement `Worksheet` root element (x:worksheet)
- [ ] 6.2.2 Implement `SheetPr` element (sheet properties)
- [ ] 6.2.3 Implement `Dimension` element
- [ ] 6.2.4 Implement `SheetViews` and `SheetView` elements
- [ ] 6.2.5 Implement `Pane` element (frozen panes)
- [ ] 6.2.6 Implement `Selection` element
- [ ] 6.2.7 Implement `SheetFormatPr` element
- [ ] 6.2.8 Implement `Cols` and `Col` elements
- [ ] 6.2.9 Implement `SheetProtection` element
- [ ] 6.2.10 Implement `PageMargins` element
- [ ] 6.2.11 Implement `PageSetup` element
- [ ] 6.2.12 Implement `HeaderFooter` element
- [ ] 6.2.13 Implement `PrintOptions` element
- [ ] 6.2.14 Implement `PageBreaks` elements (rowBreaks, colBreaks)
- [ ] 6.2.15 Implement `Drawing` element (reference to DrawingsPart)
- [ ] 6.2.16 Implement `LegacyDrawing` element (VML reference)
- [ ] 6.2.17 Implement `Picture` element (background image)
- [ ] 6.2.18 Write tests for worksheet elements

### 6.3 Sheet Data Elements
- [ ] 6.3.1 Implement `SheetData` element
- [ ] 6.3.2 Implement `Row` element with all attributes
- [ ] 6.3.3 Implement `Cell` element with all attributes
- [ ] 6.3.4 Implement `CellValue` element (x:v)
- [ ] 6.3.5 Implement `CellFormula` element (x:f)
- [ ] 6.3.6 Implement `InlineStr` and `Is` elements
- [ ] 6.3.7 Implement `RichTextRun` (x:r) and `Text` (x:t) elements
- [ ] 6.3.8 Implement `RunProperties` (x:rPr) for rich text
- [ ] 6.3.9 Write tests for sheet data elements

### 6.4 Merge Cells and Hyperlinks
- [ ] 6.4.1 Implement `MergeCells` and `MergeCell` elements
- [ ] 6.4.2 Implement `Hyperlinks` and `Hyperlink` elements
- [ ] 6.4.3 Write tests for merge cells and hyperlinks

### 6.5 Conditional Formatting
- [ ] 6.5.1 Implement `ConditionalFormatting` element
- [ ] 6.5.2 Implement `CfRule` element with all types
- [ ] 6.5.3 Implement `ColorScale` element
- [ ] 6.5.4 Implement `DataBar` element
- [ ] 6.5.5 Implement `IconSet` element
- [ ] 6.5.6 Implement `Formula` elements for conditional formatting
- [ ] 6.5.7 Implement `Cfvo` (conditional format value object) element
- [ ] 6.5.8 Write tests for conditional formatting

### 6.6 Data Validation
- [ ] 6.6.1 Implement `DataValidations` and `DataValidation` elements
- [ ] 6.6.2 Implement all validation types (list, whole, decimal, date, time, textLength, custom)
- [ ] 6.6.3 Implement error/input messages
- [ ] 6.6.4 Write tests for data validation

### 6.7 Auto Filter and Sort
- [ ] 6.7.1 Implement `AutoFilter` element
- [ ] 6.7.2 Implement `FilterColumn` element
- [ ] 6.7.3 Implement `Filters`, `Filter`, `DateGroupItem` elements
- [ ] 6.7.4 Implement `CustomFilters`, `CustomFilter` elements
- [ ] 6.7.5 Implement `Top10` element
- [ ] 6.7.6 Implement `DynamicFilter` element
- [ ] 6.7.7 Implement `ColorFilter` element
- [ ] 6.7.8 Implement `IconFilter` element
- [ ] 6.7.9 Implement `SortState` and `SortCondition` elements
- [ ] 6.7.10 Write tests for auto filter and sort

## Phase 7: Shared String Table (pkg/sml/)

### 7.1 Shared Strings
- [ ] 7.1.1 Implement `Sst` root element (x:sst)
- [ ] 7.1.2 Implement `Si` (string item) element
- [ ] 7.1.3 Implement `T` (text) element
- [ ] 7.1.4 Implement `R` (rich text run) element
- [ ] 7.1.5 Implement `RPr` (run properties) element
- [ ] 7.1.6 Implement `PhoneticRun` and `PhoneticPr` elements
- [ ] 7.1.7 Implement shared string table management (add, get, indexOf)
- [ ] 7.1.8 Implement automatic deduplication
- [ ] 7.1.9 Write tests for shared strings

## Phase 8: Stylesheet Elements (pkg/sml/)

### 8.1 Stylesheet Root
- [ ] 8.1.1 Implement `Stylesheet` root element (x:styleSheet)
- [ ] 8.1.2 Implement `NumFmts` and `NumFmt` elements
- [ ] 8.1.3 Implement built-in number format constants (0-49)
- [ ] 8.1.4 Write tests for number formats

### 8.2 Fonts
- [ ] 8.2.1 Implement `Fonts` and `Font` elements
- [ ] 8.2.2 Implement all font properties (b, i, u, strike, sz, color, name, family, scheme, charset, condense, extend, shadow, outline)
- [ ] 8.2.3 Write tests for fonts

### 8.3 Fills
- [ ] 8.3.1 Implement `Fills` and `Fill` elements
- [ ] 8.3.2 Implement `PatternFill` element with all pattern types
- [ ] 8.3.3 Implement `GradientFill` element with stops
- [ ] 8.3.4 Write tests for fills

### 8.4 Borders
- [ ] 8.4.1 Implement `Borders` and `Border` elements
- [ ] 8.4.2 Implement `Left`, `Right`, `Top`, `Bottom`, `Diagonal` border elements
- [ ] 8.4.3 Implement all border styles
- [ ] 8.4.4 Write tests for borders

### 8.5 Cell Formats
- [ ] 8.5.1 Implement `CellStyleXfs` and `Xf` elements (base styles)
- [ ] 8.5.2 Implement `CellXfs` and `Xf` elements (cell formats)
- [ ] 8.5.3 Implement `Alignment` element
- [ ] 8.5.4 Implement `Protection` element
- [ ] 8.5.5 Write tests for cell formats

### 8.6 Cell Styles
- [ ] 8.6.1 Implement `CellStyles` and `CellStyle` elements
- [ ] 8.6.2 Implement built-in style constants (Normal, etc.)
- [ ] 8.6.3 Write tests for cell styles

### 8.7 Differential Formats
- [ ] 8.7.1 Implement `Dxfs` and `Dxf` elements
- [ ] 8.7.2 Implement differential format components
- [ ] 8.7.3 Write tests for differential formats

### 8.8 Table Styles
- [ ] 8.8.1 Implement `TableStyles` and `TableStyle` elements
- [ ] 8.8.2 Implement `TableStyleElement` element
- [ ] 8.8.3 Implement built-in table style constants
- [ ] 8.8.4 Write tests for table styles

### 8.9 Colors
- [ ] 8.9.1 Implement `Colors` and `IndexedColors` elements
- [ ] 8.9.2 Implement `Color` element with RGB, theme, indexed, auto, tint
- [ ] 8.9.3 Implement theme color mapping (0-9)
- [ ] 8.9.4 Write tests for colors

### 8.10 Style Builder API
- [ ] 8.10.1 Implement `StyleBuilder` for fluent style creation
- [ ] 8.10.2 Implement automatic style deduplication
- [ ] 8.10.3 Write tests for style builder

## Phase 9: Cell Reference and Formula Support

### 9.1 Cell References
- [ ] 9.1.1 Implement `CellRef` struct with Col, Row, AbsCol, AbsRow
- [ ] 9.1.2 Implement `ParseCellRef(string)` function
- [ ] 9.1.3 Implement `CellRef.String()` for A1-style output
- [ ] 9.1.4 Implement `CellRef.R1C1()` for R1C1-style output
- [ ] 9.1.5 Implement `ColumnName(uint32)` function
- [ ] 9.1.6 Implement `ColumnIndex(string)` function
- [ ] 9.1.7 Write tests for cell references

### 9.2 Range References
- [ ] 9.2.1 Implement `RangeRef` struct with Start, End
- [ ] 9.2.2 Implement `ParseRangeRef(string)` function
- [ ] 9.2.3 Implement range iteration
- [ ] 9.2.4 Write tests for range references

### 9.3 Excel Date Serial Numbers
- [ ] 9.3.1 Implement `ExcelDateSerial(time.Time)` function
- [ ] 9.3.2 Implement `DateFromExcelSerial(float64)` function
- [ ] 9.3.3 Handle 1900 and 1904 date systems
- [ ] 9.3.4 Handle Excel's leap year bug (1900-02-29)
- [ ] 9.3.5 Write tests for date conversions

### 9.4 Formula Support
- [ ] 9.4.1 Implement normal formula storage
- [ ] 9.4.2 Implement shared formula storage
- [ ] 9.4.3 Implement array formula storage
- [ ] 9.4.4 Implement data table formula storage
- [ ] 9.4.5 Write tests for formula storage

## Phase 10: Table Elements (pkg/sml/)

### 10.1 Table Definition
- [ ] 10.1.1 Implement `Table` root element (x:table)
- [ ] 10.1.2 Implement `AutoFilter` within table
- [ ] 10.1.3 Implement `SortState` within table
- [ ] 10.1.4 Implement `TableColumns` and `TableColumn` elements
- [ ] 10.1.5 Implement `TableStyleInfo` element
- [ ] 10.1.6 Implement calculated column formulas
- [ ] 10.1.7 Implement totals row functions
- [ ] 10.1.8 Write tests for tables

## Phase 11: Drawing Elements (pkg/xdr/)

### 11.1 Worksheet Drawing
- [ ] 11.1.1 Implement `WsDr` root element (xdr:wsDr)
- [ ] 11.1.2 Implement `TwoCellAnchor` element
- [ ] 11.1.3 Implement `OneCellAnchor` element
- [ ] 11.1.4 Implement `AbsoluteAnchor` element
- [ ] 11.1.5 Implement `From` and `To` marker elements
- [ ] 11.1.6 Implement `Ext` (extent) element
- [ ] 11.1.7 Implement `Pos` (position) element
- [ ] 11.1.8 Write tests for anchors

### 11.2 Drawing Objects
- [ ] 11.2.1 Implement `Sp` (shape) element
- [ ] 11.2.2 Implement `Pic` (picture) element
- [ ] 11.2.3 Implement `CxnSp` (connection shape) element
- [ ] 11.2.4 Implement `GraphicFrame` element
- [ ] 11.2.5 Implement `GrpSp` (group shape) element
- [ ] 11.2.6 Implement `ClientData` element
- [ ] 11.2.7 Implement non-visual properties elements
- [ ] 11.2.8 Write tests for drawing objects

### 11.3 EMU Utilities
- [ ] 11.3.1 Implement EMU conversion constants
- [ ] 11.3.2 Implement inch/cm/point/pixel to EMU functions
- [ ] 11.3.3 Write tests for EMU conversions

## Phase 12: Chart Elements (pkg/dml/chart/)

### 12.1 Chart Root
- [ ] 12.1.1 Implement `ChartSpace` root element (c:chartSpace)
- [ ] 12.1.2 Implement `Chart` element
- [ ] 12.1.3 Implement `PlotArea` element
- [ ] 12.1.4 Implement `Legend` element
- [ ] 12.1.5 Implement `Title` element
- [ ] 12.1.6 Write tests for chart root

### 12.2 Chart Types
- [ ] 12.2.1 Implement `BarChart` and `Bar3DChart` elements
- [ ] 12.2.2 Implement `LineChart` and `Line3DChart` elements
- [ ] 12.2.3 Implement `PieChart`, `Pie3DChart`, `DoughnutChart` elements
- [ ] 12.2.4 Implement `AreaChart` and `Area3DChart` elements
- [ ] 12.2.5 Implement `ScatterChart` element
- [ ] 12.2.6 Implement `BubbleChart` element
- [ ] 12.2.7 Implement `StockChart` element
- [ ] 12.2.8 Implement `SurfaceChart` and `Surface3DChart` elements
- [ ] 12.2.9 Implement `RadarChart` element
- [ ] 12.2.10 Implement `OfPieChart` element
- [ ] 12.2.11 Write tests for chart types

### 12.3 Chart Components
- [ ] 12.3.1 Implement series elements for each chart type
- [ ] 12.3.2 Implement `CatAx` (category axis) element
- [ ] 12.3.3 Implement `ValAx` (value axis) element
- [ ] 12.3.4 Implement `DateAx` (date axis) element
- [ ] 12.3.5 Implement `SerAx` (series axis) element
- [ ] 12.3.6 Implement `DataLabels` element
- [ ] 12.3.7 Implement `Trendline` element
- [ ] 12.3.8 Implement `ErrBars` element
- [ ] 12.3.9 Write tests for chart components

### 12.4 Chart Data References
- [ ] 12.4.1 Implement `NumRef` and `NumCache` elements
- [ ] 12.4.2 Implement `StrRef` and `StrCache` elements
- [ ] 12.4.3 Implement `MultiLvlStrRef` element
- [ ] 12.4.4 Write tests for data references

### 12.5 Sparklines
- [ ] 12.5.1 Implement `SparklineGroups` and `SparklineGroup` elements
- [ ] 12.5.2 Implement `Sparklines` and `Sparkline` elements
- [ ] 12.5.3 Implement sparkline types (line, column, stacked)
- [ ] 12.5.4 Write tests for sparklines

## Phase 13: Pivot Table Elements (pkg/sml/)

### 13.1 Pivot Table Definition
- [ ] 13.1.1 Implement `PivotTableDefinition` root element
- [ ] 13.1.2 Implement `Location` element
- [ ] 13.1.3 Implement `PivotFields` and `PivotField` elements
- [ ] 13.1.4 Implement `Items` and `Item` elements
- [ ] 13.1.5 Implement `RowFields` and `ColFields` elements
- [ ] 13.1.6 Implement `PageFields` and `PageField` elements
- [ ] 13.1.7 Implement `DataFields` and `DataField` elements
- [ ] 13.1.8 Implement `PivotTableStyleInfo` element
- [ ] 13.1.9 Write tests for pivot table definition

### 13.2 Pivot Cache
- [ ] 13.2.1 Implement `PivotCacheDefinition` root element
- [ ] 13.2.2 Implement `CacheSource` element
- [ ] 13.2.3 Implement `WorksheetSource` element
- [ ] 13.2.4 Implement `CacheFields` and `CacheField` elements
- [ ] 13.2.5 Implement `SharedItems` element
- [ ] 13.2.6 Implement item elements (Boolean, DateTime, Error, Missing, Number, String)
- [ ] 13.2.7 Write tests for pivot cache definition

### 13.3 Pivot Cache Records
- [ ] 13.3.1 Implement `PivotCacheRecords` root element
- [ ] 13.3.2 Implement `R` (record) element
- [ ] 13.3.3 Implement record value elements
- [ ] 13.3.4 Write tests for pivot cache records

### 13.4 Slicers
- [ ] 13.4.1 Implement `Slicers` root element
- [ ] 13.4.2 Implement `Slicer` element
- [ ] 13.4.3 Implement `SlicerCacheDefinition` root element
- [ ] 13.4.4 Write tests for slicers

### 13.5 Timelines
- [ ] 13.5.1 Implement `Timelines` root element
- [ ] 13.5.2 Implement `Timeline` element
- [ ] 13.5.3 Implement `TimeLineCacheDefinition` root element
- [ ] 13.5.4 Write tests for timelines

## Phase 14: Comments Elements (pkg/sml/)

### 14.1 Comments
- [ ] 14.1.1 Implement `Comments` root element (x:comments)
- [ ] 14.1.2 Implement `Authors` and `Author` elements
- [ ] 14.1.3 Implement `CommentList` and `Comment` elements
- [ ] 14.1.4 Implement comment text with rich formatting
- [ ] 14.1.5 Write tests for comments

### 14.2 Threaded Comments
- [ ] 14.2.1 Implement `ThreadedComments` root element
- [ ] 14.2.2 Implement `ThreadedComment` element
- [ ] 14.2.3 Implement `PersonList` and `Person` elements
- [ ] 14.2.4 Write tests for threaded comments

## Phase 15: High-Level API

### 15.1 Worksheet API
- [ ] 15.1.1 Implement `Sheet.Cell(ref string)` accessor
- [ ] 15.1.2 Implement `Sheet.CellAt(row, col uint32)` accessor
- [ ] 15.1.3 Implement `Sheet.Range(ref string)` accessor
- [ ] 15.1.4 Implement `Sheet.Row(index uint32)` accessor
- [ ] 15.1.5 Implement `Sheet.AddRow()` method
- [ ] 15.1.6 Implement `Sheet.MergeCells(ref string)` method
- [ ] 15.1.7 Implement `Sheet.UnmergeCells(ref string)` method
- [ ] 15.1.8 Implement `Sheet.SetColumnWidth(col uint32, width float64)` method
- [ ] 15.1.9 Implement `Sheet.SetRowHeight(row uint32, height float64)` method
- [ ] 15.1.10 Write tests for worksheet API

### 15.2 Cell API
- [ ] 15.2.1 Implement `Cell.SetNumber(float64)` method
- [ ] 15.2.2 Implement `Cell.SetString(string)` method (with shared strings)
- [ ] 15.2.3 Implement `Cell.SetInlineString(string)` method
- [ ] 15.2.4 Implement `Cell.SetBoolean(bool)` method
- [ ] 15.2.5 Implement `Cell.SetDate(time.Time)` method
- [ ] 15.2.6 Implement `Cell.SetFormula(string)` method
- [ ] 15.2.7 Implement `Cell.SetStyle(Style)` method
- [ ] 15.2.8 Implement `Cell.GetNumber()`, `GetString()`, `GetBoolean()`, `GetDate()` methods
- [ ] 15.2.9 Write tests for cell API

### 15.3 Style Builder API
- [ ] 15.3.1 Implement `NewStyle()` factory
- [ ] 15.3.2 Implement `Style.Font(...)` method
- [ ] 15.3.3 Implement `Style.Fill(...)` method
- [ ] 15.3.4 Implement `Style.Border(...)` method
- [ ] 15.3.5 Implement `Style.NumberFormat(...)` method
- [ ] 15.3.6 Implement `Style.Alignment(...)` method
- [ ] 15.3.7 Implement `Style.Protection(...)` method
- [ ] 15.3.8 Write tests for style builder

### 15.4 Table API
- [ ] 15.4.1 Implement `Sheet.AddTable(ref string, name string)` method
- [ ] 15.4.2 Implement `Table.SetStyle(...)` method
- [ ] 15.4.3 Implement `Table.Column(name string)` accessor
- [ ] 15.4.4 Write tests for table API

### 15.5 Chart API
- [ ] 15.5.1 Implement `Sheet.AddChart(chartType, dataRange, position)` method
- [ ] 15.5.2 Implement `Chart.SetTitle(string)` method
- [ ] 15.5.3 Implement `Chart.Series(index int)` accessor
- [ ] 15.5.4 Implement chart axis accessors
- [ ] 15.5.5 Write tests for chart API

### 15.6 Pivot Table API
- [ ] 15.6.1 Implement `Sheet.AddPivotTable(sourceRange, destCell)` method
- [ ] 15.6.2 Implement `PivotTable.AddRowField(name)` method
- [ ] 15.6.3 Implement `PivotTable.AddColumnField(name)` method
- [ ] 15.6.4 Implement `PivotTable.AddDataField(name, aggregate)` method
- [ ] 15.6.5 Implement `PivotTable.AddPageField(name)` method
- [ ] 15.6.6 Write tests for pivot table API

### 15.7 Drawing API
- [ ] 15.7.1 Implement `Sheet.AddImage(cell, data, contentType)` method
- [ ] 15.7.2 Implement `Sheet.AddImageBetween(from, to, data, contentType)` method
- [ ] 15.7.3 Write tests for drawing API

## Phase 16: Integration and Quality

### 16.1 Integration Tests
- [ ] 16.1.1 Create end-to-end tests creating workbooks from scratch
- [ ] 16.1.2 Create roundtrip tests (open, modify, save, reopen)
- [ ] 16.1.3 Test compatibility with Microsoft Excel (Office 2016+)
- [ ] 16.1.4 Test compatibility with LibreOffice Calc
- [ ] 16.1.5 Test with various Office 2016+ documents
- [ ] 16.1.6 Create fixture documents from multiple Office versions for testing
- [ ] 16.1.7 Test all validation constraints with invalid documents
- [ ] 16.1.8 Test large worksheet handling (100K+ rows)

### 16.2 Documentation
- [ ] 16.2.1 Write package-level godoc documentation
- [ ] 16.2.2 Create API examples for common tasks
- [ ] 16.2.3 Document relationship to Open-XML-SDK
- [ ] 16.2.4 Create migration guide from other Go libraries (excelize, xlsx)
- [ ] 16.2.5 Document idiomatic Go naming conventions vs C# SDK names

### 16.3 Performance
- [ ] 16.3.1 Benchmark workbook creation (target: 10MB in <1s)
- [ ] 16.3.2 Benchmark large worksheet handling (target: 1M cells in <5s)
- [ ] 16.3.3 Benchmark validation performance (target: <2s for typical docs)
- [ ] 16.3.4 Optimize hot paths based on benchmarks
- [ ] 16.3.5 Memory profiling and optimization
- [ ] 16.3.6 Ensure lazy loading works correctly for large documents
- [ ] 16.3.7 Implement streaming row writer for large data sets

### 16.4 Enum Values (sml/ package)
- [ ] 16.4.1 Implement all CellType values
- [ ] 16.4.2 Implement all FormulaType values
- [ ] 16.4.3 Implement all BorderStyle values
- [ ] 16.4.4 Implement all PatternType values (fills)
- [ ] 16.4.5 Implement all HorizontalAlignment values
- [ ] 16.4.6 Implement all VerticalAlignment values
- [ ] 16.4.7 Implement all DataValidationType values
- [ ] 16.4.8 Implement all ConditionalFormatType values
- [ ] 16.4.9 Implement all ChartType values
- [ ] 16.4.10 Implement all remaining SpreadsheetML enum types
- [ ] 16.4.11 Write enum validation functions

## Notes

**Design Decisions Applied:**
- All element types are generated from JSON schemas (pre-generated, committed)
- Idiomatic Go naming: `Document` not `SpreadsheetDocument`, `Sheet` not `Worksheet`
- Full validation with all semantic constraint types
- Office 2016+ only (ECMA-376 5th edition)

**Dependencies:**
- Phase 2-4 are shared with Word and Presentation SDKs
- Phase 5 depends on Phases 1-4
- Phase 6-14 depend on Phase 5
- Phase 15 depends on Phases 6-14
- Phase 16 runs throughout and at end

**Parallelization:**
- Within each phase, tasks marked x.y.z can often run in parallel
- Different element type implementations can be parallelized
- Test writing can parallel implementation

**Validation Checkpoints:**
- After Phase 5: Can create/open/save empty .xlsx (OPC compliance)
- After Phase 6+7: Can create workbooks with basic sheets and data
- After Phase 8: Can create formatted workbooks with styles
- After Phase 12: Can create workbooks with charts
- After Phase 13: Can create workbooks with pivot tables
- After Phase 16: Full feature parity with Open-XML-SDK SpreadsheetML

**Estimated Scope:**
- ~600+ element types to generate/write
- ~29 simple types (shared)
- ~19 semantic constraint types (shared)
- ~25+ part types
- ~100+ enum types in sml/ package
- Total estimated: 300+ Go files, 40,000+ lines of code
