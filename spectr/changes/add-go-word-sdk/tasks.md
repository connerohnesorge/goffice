# Implementation Tasks: Go SDK for Word Document Processing

**Unified Design Decisions Applied (consistent with Presentation SDK):**
- Module path: `github.com/connerohnesorge/goffice`
- Code generation from JSON schemas (pre-generated, committed)
- Office 2016+ only (ECMA-376 5th edition+)
- Package structure: `pkg/{framework,types,package,relationships,validation,features,word,wml}/`
- Go-idiomatic short API names (`Document`, `Para`, `ParaProps`)
- Embedded struct fields for element metadata
- Functional options pattern for element construction
- Generic methods only for child access: `First[T]()`, `All[T]()`, `OfType[T]()`
- Full Office Math (OMML) v1 implementation (all 152 types)
- *T pointers for optional/nullable values
- Priority: Core -> Formatting -> Tables -> Advanced

## Phase 1: Project Foundation

### 1.1 Project Setup
- [ ] 1.1.1 Initialize Go module `github.com/connerohnesorge/goffice`
- [ ] 1.1.2 Create directory structure: `pkg/`, `internal/`, `cmd/`, `testdata/`
- [ ] 1.1.3 Set up testing infrastructure with `go test` and test fixtures directory
- [ ] 1.1.4 Create sample .docx test files for roundtrip testing (Office 2016+)
- [ ] 1.1.5 Add Makefile with build, test, lint targets

### 1.2 Core Simple Types (pkg/types/)
- [ ] 1.2.1 Implement `SimpleValue` interface with `HasValue()`, `InnerText()`, `SetInnerText()`
- [ ] 1.2.2 Implement `StringValue` type
- [ ] 1.2.3 Implement `Int32Value` and `UInt32Value` types
- [ ] 1.2.4 Implement `Int64Value` and `UInt64Value` types
- [ ] 1.2.5 Implement `BooleanValue` type (true/false and 1/0)
- [ ] 1.2.6 Implement `OnOffValue` type (on/off Word-style)
- [ ] 1.2.7 Implement `EnumValue[T]` generic type for enumerations
- [ ] 1.2.8 Implement `HexBinaryValue` type
- [ ] 1.2.9 Implement `Base64BinaryValue` type
- [ ] 1.2.10 Implement `DateTimeValue` type (ISO 8601)
- [ ] 1.2.11 Implement `DecimalValue` type
- [ ] 1.2.12 Implement unit types: `TwipsValue`, `HalfPointsValue`, `EmuValue`
- [ ] 1.2.13 Implement `PercentageValue` type

### 1.3 Phase 1 Testing
- [ ] 1.3.1 Write unit tests for all simple types in pkg/types/
- [ ] 1.3.2 Write unit tests for SimpleValue interface compliance
- [ ] 1.3.3 Write unit tests for type conversion edge cases
- [ ] 1.3.4 Run tests and verify 100% pass rate
- [ ] 1.3.5 Integration test: verify project setup compiles and runs

## Phase 2: OPC Packaging Layer (internal/opc/ and pkg/package/)

### 2.1 Core Package Types
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

### 2.3 Package Parts
- [ ] 2.3.1 Implement `Part` struct with URI, ContentType, Stream
- [ ] 2.3.2 Implement `Package.CreatePart(uri, contentType)`
- [ ] 2.3.3 Implement `Package.Part(uri)` accessor
- [ ] 2.3.4 Implement `Package.DeletePart(uri)`
- [ ] 2.3.5 Implement `Package.Parts()` enumeration
- [ ] 2.3.6 Implement `Part.Stream(mode)` for part content access
- [ ] 2.3.7 Implement part URI normalization and validation
- [ ] 2.3.8 Implement `ResolveURI(base, relative)` for relative URI resolution

### 2.4 Package Relationships (pkg/relationships/)
- [ ] 2.4.1 Implement `Relationship` struct
- [ ] 2.4.2 Implement `Package.CreateRel(target, relType, id)`
- [ ] 2.4.3 Implement `Package.RelsByType(relType)` accessor
- [ ] 2.4.4 Implement `Package.Rel(id)` accessor
- [ ] 2.4.5 Implement `Package.DeleteRel(id)`
- [ ] 2.4.6 Implement `_rels/.rels` serialization/deserialization
- [ ] 2.4.7 Implement part-level `.rels` file handling

### 2.5 Core Properties
- [ ] 2.5.1 Implement `CoreProperties` struct (title, author, subject, etc.)
- [ ] 2.5.2 Implement `Package.CoreProperties()` accessor
- [ ] 2.5.3 Implement serialization to `docProps/core.xml`

### 2.6 Phase 2 Testing
- [ ] 2.6.1 Write unit tests for Package struct and all methods
- [ ] 2.6.2 Write unit tests for ContentTypes management
- [ ] 2.6.3 Write unit tests for Package Parts operations
- [ ] 2.6.4 Write unit tests for Package Relationships
- [ ] 2.6.5 Write unit tests for Core Properties
- [ ] 2.6.6 Run tests and verify 100% pass rate
- [ ] 2.6.7 Integration test: create/open/save empty .docx (OPC compliance)
- [ ] 2.6.8 Integration test with Phase 1 simple types

## Phase 3: OpenXML Framework (pkg/framework/)

### 3.1 Feature Collection System (pkg/features/)
- [ ] 3.1.1 Implement `FeatureCollection` struct with parent inheritance
- [ ] 3.1.2 Implement `Get[T]()` with parent chain traversal
- [ ] 3.1.3 Implement `Set(feature)` for registration
- [ ] 3.1.4 Implement thread-safe access (sync.Mutex)
- [ ] 3.1.5 Define `IPackageFeature` interface and implementation
- [ ] 3.1.6 Define `IContentTypeFeature` interface and implementation
- [ ] 3.1.7 Define `INamespaceFeature` interface and implementation
- [ ] 3.1.8 Define `IPartRelationshipsFeature` interface
- [ ] 3.1.9 Define `IMainPartFeature` interface

### 3.2 Element Base Types
- [ ] 3.2.1 Implement `Element` interface with core methods
- [ ] 3.2.2 Implement `elementBase` struct with common fields
- [ ] 3.2.3 Implement `CompositeElement` interface and `compositeBase` struct
- [ ] 3.2.4 Implement `LeafElement` interface and `leafBase` struct
- [ ] 3.2.5 Implement `OpenXmlQualifiedName` for namespace:localname
- [ ] 3.2.6 Implement `Attribute` struct and attribute management

### 3.3 Element Tree Operations
- [ ] 3.3.1 Implement `AppendChild()`, `PrependChild()` operations
- [ ] 3.3.2 Implement `InsertBefore()`, `InsertAfter()` operations
- [ ] 3.3.3 Implement `RemoveChild()`, `RemoveAllChildren()` operations
- [ ] 3.3.4 Implement `ReplaceChild()` operation
- [ ] 3.3.5 Implement `FirstChild()`, `LastChild()` accessors
- [ ] 3.3.6 Implement `NextSibling()`, `PreviousSibling()` navigation
- [ ] 3.3.7 Implement `Parent()` accessor and parent tracking
- [ ] 3.3.8 Implement `Children()` and `Elements[T]()` enumeration

### 3.4 Element Cloning
- [ ] 3.4.1 Implement `Clone()` for deep copy
- [ ] 3.4.2 Implement `CloneNode(deep bool)` for optional shallow copy
- [ ] 3.4.3 Ensure clone independence from original

### 3.5 XML Serialization
- [ ] 3.5.1 Implement `XMLWriter` for namespace-aware serialization
- [ ] 3.5.2 Implement `Element.WriteTo(io.Writer)` serialization
- [ ] 3.5.3 Implement `Element.OuterXml()` and `InnerXml()`
- [ ] 3.5.4 Implement `ParseElement(io.Reader)` for XML parsing
- [ ] 3.5.5 Implement namespace prefix management
- [ ] 3.5.6 Handle strict vs transitional namespace mapping

### 3.6 Part System
- [ ] 3.6.1 Implement `OpenXmlPart` base struct
- [ ] 3.6.2 Implement lazy loading of part root element
- [ ] 3.6.3 Implement `OpenXmlPartContainer` for child part management
- [ ] 3.6.4 Implement `GetPartsOfType[T]()` generic accessor
- [ ] 3.6.5 Implement `AddPart()`, `DeletePart()` operations
- [ ] 3.6.6 Implement part type registry by content type
- [ ] 3.6.7 Implement `OpenXmlPartRootElement` with Save/Reload

### 3.7 Relationship Management
- [ ] 3.7.1 Implement `Relationship` base struct
- [ ] 3.7.2 Implement `PartRelationship` for internal relationships
- [ ] 3.7.3 Implement `ExternalRelationship` for external URIs
- [ ] 3.7.4 Implement `HyperlinkRelationship` for hyperlinks
- [ ] 3.7.5 Implement relationship type constants (all standard types)
- [ ] 3.7.6 Implement unique ID generation (rId1, rId2, etc.)
- [ ] 3.7.7 Implement target URI resolution

### 3.8 Package Abstraction
- [ ] 3.8.1 Implement `OpenXmlPackage` abstract struct
- [ ] 3.8.2 Implement package-level feature collection
- [ ] 3.8.3 Implement package save orchestration (dirty tracking)

### 3.9 Phase 3 Testing
- [ ] 3.9.1 Write unit tests for FeatureCollection system
- [ ] 3.9.2 Write unit tests for Element base types
- [ ] 3.9.3 Write unit tests for Element tree operations
- [ ] 3.9.4 Write unit tests for Element cloning
- [ ] 3.9.5 Write unit tests for XML serialization with roundtrip validation
- [ ] 3.9.6 Write unit tests for Part system
- [ ] 3.9.7 Write unit tests for Relationship management
- [ ] 3.9.8 Write unit tests for OpenXmlPackage
- [ ] 3.9.9 Run tests and verify 100% pass rate
- [ ] 3.9.10 Integration test with Phase 1 and Phase 2 components

## Phase 4: Validation Framework (pkg/validation/)

### 4.1 Validation Infrastructure
- [ ] 4.1.1 Implement `Validator` interface
- [ ] 4.1.2 Implement `ValidationContext` with version and settings
- [ ] 4.1.3 Implement `ValidationError` with path, description, severity
- [ ] 4.1.4 Implement `ValidationSettings` (MaxErrors, ContinueOnError)
- [ ] 4.1.5 Define `FileFormatVersions` enum (Office2016, Office2019, Office2021, Microsoft365)

### 4.2 Schema Validation
- [ ] 4.2.1 Implement `SchemaValidator` for structure validation
- [ ] 4.2.2 Implement `CompiledParticle` for schema particles
- [ ] 4.2.3 Implement sequence particle validation
- [ ] 4.2.4 Implement choice particle validation
- [ ] 4.2.5 Implement all particle validation
- [ ] 4.2.6 Implement minOccurs/maxOccurs validation

### 4.3 Attribute Validation
- [ ] 4.3.1 Implement required attribute validation
- [ ] 4.3.2 Implement attribute type validation
- [ ] 4.3.3 Implement enumeration value validation
- [ ] 4.3.4 Implement range constraint validation (min/max)

### 4.4 Semantic Validation (All 19+ Constraint Types)
- [ ] 4.4.1 Implement `SemanticValidator` base
- [ ] 4.4.2 Implement `Constraint` interface
- [ ] 4.4.3 Implement `AttributeValueRangeConstraint` (min/max bounds)
- [ ] 4.4.4 Implement `AttributeValuePatternConstraint` (regex validation)
- [ ] 4.4.5 Implement `AttributeValueSetConstraint` (enumeration values)
- [ ] 4.4.6 Implement `ParentTypeConstraint` (allowed parent elements)
- [ ] 4.4.7 Implement `ChildElementConstraint` (required/allowed children)
- [ ] 4.4.8 Implement `UniqueValueConstraint` (uniqueness within scope)
- [ ] 4.4.9 Implement `RelationshipExistConstraint` (relationship must exist)
- [ ] 4.4.10 Implement `RelationshipTypeConstraint` (relationship type validation)
- [ ] 4.4.11 Implement `ReferenceExistConstraint` (ID reference validation)
- [ ] 4.4.12 Implement `IndexRangeConstraint` (valid index bounds)
- [ ] 4.4.13 Implement `RootAttributeConstraint` (required root attributes)
- [ ] 4.4.14 Implement `AttributeAbsentConstraint` (mutually exclusive attributes)
- [ ] 4.4.15 Implement `AttributeCannotOmitConstraint` (conditionally required)
- [ ] 4.4.16 Implement `AttributeValueLengthConstraint` (string length limits)
- [ ] 4.4.17 Implement `UniqueAttributeValueConstraint` (unique across document)
- [ ] 4.4.18 Implement `PartContainerConstraint` (part must be in container)
- [ ] 4.4.19 Implement `DataPartConstraint` (data part validation)
- [ ] 4.4.20 Implement `PartTypeConstraint` (part content type validation)
- [ ] 4.4.21 Implement custom constraint registration system

### 4.5 Element Metadata
- [ ] 4.5.1 Implement `IElementMetadata` interface
- [ ] 4.5.2 Implement `AttributeMetadata` struct
- [ ] 4.5.3 Implement child element metadata
- [ ] 4.5.4 Implement version availability tracking

### 4.6 Phase 4 Testing
- [ ] 4.6.1 Write unit tests for validation infrastructure
- [ ] 4.6.2 Write unit tests for schema validation
- [ ] 4.6.3 Write unit tests for attribute validation
- [ ] 4.6.4 Write comprehensive tests for all 19+ semantic constraints
- [ ] 4.6.5 Write unit tests for element metadata system
- [ ] 4.6.6 Run tests and verify 100% pass rate
- [ ] 4.6.7 Integration test with Phase 1-3 components

## Phase 5: Word Document (pkg/word/)

### 5.1 Document Type
- [ ] 5.1.1 Implement `Document` struct (was WordprocessingDocument)
- [ ] 5.1.2 Implement `DocType` enum (Document, Template, MacroEnabled, MacroTemplate)
- [ ] 5.1.3 Implement `New(path, docType)` for new documents
- [ ] 5.1.4 Implement `NewWriter(w, docType)` for stream creation
- [ ] 5.1.5 Implement `Open(path, editable)` for existing documents
- [ ] 5.1.6 Implement `OpenReader(r, size, editable)` for streams
- [ ] 5.1.7 Implement `OpenOptions` struct for configuration
- [ ] 5.1.8 Implement `NewFromTemplate(templatePath, attachTemplate)`
- [ ] 5.1.9 Implement `Save()`, `SaveAs(path)`, `SaveTo(w)`
- [ ] 5.1.10 Implement `Close()` with io.Closer interface
- [ ] 5.1.11 Implement `ChangeType(newType)`
- [ ] 5.1.12 Implement document-level `Validate(version)`

### 5.2 Document Parts
- [ ] 5.2.1 Implement `MainPart` with content type detection
- [ ] 5.2.2 Implement `StylesPart`
- [ ] 5.2.3 Implement `NumberingPart`
- [ ] 5.2.4 Implement `SettingsPart`
- [ ] 5.2.5 Implement `WebSettingsPart`
- [ ] 5.2.6 Implement `FontsPart`
- [ ] 5.2.7 Implement `HeaderPart` and `FooterPart`
- [ ] 5.2.8 Implement `FootnotesPart` and `EndnotesPart`
- [ ] 5.2.9 Implement `CommentsPart`
- [ ] 5.2.10 Implement `GlossaryPart`
- [ ] 5.2.11 Implement `ThemePart`
- [ ] 5.2.12 Implement `ImagePart` with image type enum
- [ ] 5.2.13 Implement `CustomXmlPart`
- [ ] 5.2.14 Implement `VbaPart` for macro-enabled documents
- [ ] 5.2.15 Implement part addition methods on MainPart

### 5.3 Phase 5 Testing
- [ ] 5.3.1 Write unit tests for Document struct and all methods
- [ ] 5.3.2 Write unit tests for DocType enum
- [ ] 5.3.3 Write unit tests for all Document Parts
- [ ] 5.3.4 Write integration tests with real .docx files (Office 2016+)
- [ ] 5.3.5 Run tests and verify 100% pass rate
- [ ] 5.3.6 Integration test: create basic documents with text
- [ ] 5.3.7 Integration test with Phase 1-4 components

## Phase 6: Word Elements - Generated (pkg/wml/)

**Note:** All 685+ element types are generated from JSON schemas (unified with Presentation SDK). The generator reads C# SDK schema files and produces Go code. Organized by category for implementation tracking.

### 6.1 Document Structure Elements (~25 types)
- [ ] 6.1.1 Implement `Doc` root element (w:document)
- [ ] 6.1.2 Implement `Body` element (w:body)
- [ ] 6.1.3 Implement `SectProps` element (w:sectPr)
- [ ] 6.1.4 Implement `PageSz` element (w:pgSz)
- [ ] 6.1.5 Implement `PageMar` element (w:pgMar)
- [ ] 6.1.6 Implement `Cols` element (w:cols)
- [ ] 6.1.7 Implement `Col` element (w:col)
- [ ] 6.1.8 Implement `DocGrid` element (w:docGrid)
- [ ] 6.1.9 Implement `HeaderRef` and `FooterRef` elements
- [ ] 6.1.10 Implement `TitlePage` element (w:titlePg)
- [ ] 6.1.11 Implement `SectType` element (w:type)
- [ ] 6.1.12 Implement `PageNumType` element (w:pgNumType)
- [ ] 6.1.13 Implement `FormProt` element (w:formProt)
- [ ] 6.1.14 Implement `VAlign` element (w:vAlign)
- [ ] 6.1.15 Implement `NoEndnote` element (w:noEndnote)
- [ ] 6.1.16 Implement `PaperSrc` element (w:paperSrc)
- [ ] 6.1.17 Implement `LnNumType` element (w:lnNumType)
- [ ] 6.1.18 Implement `TextDir` element (w:textDirection)
- [ ] 6.1.19 Implement `RTLGutter` element (w:rtlGutter)
- [ ] 6.1.20 Implement `DocVar` elements (w:docVars, w:docVar)
- [ ] 6.1.21 Implement `Background` element (w:background)

### 6.1 Testing - Document Structure
- [ ] 6.1.22 Write unit tests for all document structure elements
- [ ] 6.1.23 Run tests and verify 100% pass rate
- [ ] 6.1.24 Integration test with previous phases

### 6.2 Paragraph Elements (~15 types)
- [ ] 6.2.1 Implement `Para` element (w:p) with convenience constructors
- [ ] 6.2.2 Implement `ParaProps` element (w:pPr)
- [ ] 6.2.3 Implement `Run` element (w:r) with text convenience
- [ ] 6.2.4 Implement `RunProps` element (w:rPr)
- [ ] 6.2.5 Implement `Text` element (w:t) with space preservation
- [ ] 6.2.6 Implement `Break` element (w:br) - line, page, column
- [ ] 6.2.7 Implement `Tab` element (w:tab)
- [ ] 6.2.8 Implement `CarriageReturn` element (w:cr)
- [ ] 6.2.9 Implement `SoftHyphen` element (w:softHyphen)
- [ ] 6.2.10 Implement `NoBreakHyphen` element (w:noBreakHyphen)
- [ ] 6.2.11 Implement `LastRenderedPageBreak` element
- [ ] 6.2.12 Implement `DayShort`, `MonthShort`, `YearShort` date elements
- [ ] 6.2.13 Implement `DayLong`, `MonthLong`, `YearLong` date elements
- [ ] 6.2.14 Implement `AnnotationRef` element

### 6.2 Testing - Paragraph Elements
- [ ] 6.2.15 Write unit tests for all paragraph elements
- [ ] 6.2.16 Run tests and verify 100% pass rate
- [ ] 6.2.17 Integration test with previous phases

### 6.3 Paragraph Properties (~50 types)
- [ ] 6.3.1 Implement `Jc` (justification/alignment)
- [ ] 6.3.2 Implement `Ind` (indentation - left, right, firstLine, hanging)
- [ ] 6.3.3 Implement `Spacing` (before, after, line spacing)
- [ ] 6.3.4 Implement `KeepNext`, `KeepLines`, `PageBreakBefore`
- [ ] 6.3.5 Implement `WidowControl`
- [ ] 6.3.6 Implement `NumPr` (numbering properties)
- [ ] 6.3.7 Implement `NumId` and `Ilvl` elements
- [ ] 6.3.8 Implement `Tabs` and `TabStop` elements
- [ ] 6.3.9 Implement `PBdr` (paragraph borders) element
- [ ] 6.3.10 Implement border elements: `Top`, `Bottom`, `Left`, `Right`, `Between`, `Bar`
- [ ] 6.3.11 Implement `Shd` (shading) element
- [ ] 6.3.12 Implement `OutlineLvl` element
- [ ] 6.3.13 Implement `PStyle` (paragraph style reference)
- [ ] 6.3.14 Implement `SuppressAutoHyphens`
- [ ] 6.3.15 Implement `Kinsoku`
- [ ] 6.3.16 Implement `WordWrap`
- [ ] 6.3.17 Implement `OverflowPunct`
- [ ] 6.3.18 Implement `TopLinePunct`
- [ ] 6.3.19 Implement `AutoSpaceDE`, `AutoSpaceDN`
- [ ] 6.3.20 Implement `Bidi` (bidirectional)
- [ ] 6.3.21 Implement `AdjustRightInd`
- [ ] 6.3.22 Implement `SnapToGrid`
- [ ] 6.3.23 Implement `ContextualSpacing`
- [ ] 6.3.24 Implement `MirrorIndents`
- [ ] 6.3.25 Implement `SuppressOverlap`
- [ ] 6.3.26 Implement `TextAlignment`
- [ ] 6.3.27 Implement `TextboxTightWrap`
- [ ] 6.3.28 Implement `FramePr` (frame properties)
- [ ] 6.3.29 Implement `DivId`
- [ ] 6.3.30 Implement `CnfStyle` (conditional formatting)
- [ ] 6.3.31 Implement `ParaPropsChange` (revision tracking)

### 6.3 Testing - Paragraph Properties
- [ ] 6.3.32 Write unit tests for all paragraph properties
- [ ] 6.3.33 Run tests and verify 100% pass rate
- [ ] 6.3.34 Integration test with previous phases

### 6.4 Run Properties (~60 types)
- [ ] 6.4.1 Implement `Bold` (w:b), `BoldComplexScript` (w:bCs)
- [ ] 6.4.2 Implement `Italic` (w:i), `ItalicComplexScript` (w:iCs)
- [ ] 6.4.3 Implement `Underline` (w:u) with all underline styles
- [ ] 6.4.4 Implement `Strike` (w:strike), `DoubleStrike` (w:dstrike)
- [ ] 6.4.5 Implement `Sz` (font size), `SzCs` (complex script size)
- [ ] 6.4.6 Implement `RFonts` (run fonts) with all font families
- [ ] 6.4.7 Implement `Color` (w:color) with theme colors
- [ ] 6.4.8 Implement `Highlight` (w:highlight)
- [ ] 6.4.9 Implement `Caps`, `SmallCaps`
- [ ] 6.4.10 Implement `VertAlign` (sub/superscript)
- [ ] 6.4.11 Implement `CharSpacing` (w:spacing)
- [ ] 6.4.12 Implement `Position` (raise/lower)
- [ ] 6.4.13 Implement `Width` (character scale)
- [ ] 6.4.14 Implement `Emboss`, `Imprint`, `Shadow`, `Outline`
- [ ] 6.4.15 Implement `Vanish` (hidden text)
- [ ] 6.4.16 Implement `Lang` (languages)
- [ ] 6.4.17 Implement `RStyle` (run style reference)
- [ ] 6.4.18 Implement `Kern` (kerning)
- [ ] 6.4.19 Implement `Effect` (text effects)
- [ ] 6.4.20 Implement `Speculative` (speculative run)
- [ ] 6.4.21 Implement `NoProof`
- [ ] 6.4.22 Implement `WebHidden`
- [ ] 6.4.23 Implement `SnapToGrid`
- [ ] 6.4.24 Implement `EastAsianLayout`
- [ ] 6.4.25 Implement `FitText`
- [ ] 6.4.26 Implement `RunPropsChange` (revision tracking)
- [ ] 6.4.27 Implement `RunFontsChoice`
- [ ] 6.4.28 Implement `EM` (emphasis mark)
- [ ] 6.4.29 Implement `CS` (complex script)
- [ ] 6.4.30 Implement `RTL` (right-to-left)

### 6.4 Testing - Run Properties
- [ ] 6.4.31 Write unit tests for all run properties
- [ ] 6.4.32 Run tests and verify 100% pass rate
- [ ] 6.4.33 Integration test with previous phases

### 6.5 Table Elements (~30 types)
- [ ] 6.5.1 Implement `Table` (w:tbl) with grid factory
- [ ] 6.5.2 Implement `TableProps` (w:tblPr)
- [ ] 6.5.3 Implement `TableGrid` (w:tblGrid)
- [ ] 6.5.4 Implement `GridCol` (w:gridCol)
- [ ] 6.5.5 Implement `TableRow` (w:tr)
- [ ] 6.5.6 Implement `TableRowProps` (w:trPr)
- [ ] 6.5.7 Implement `TableCell` (w:tc)
- [ ] 6.5.8 Implement `CellProps` (w:tcPr)
- [ ] 6.5.9 Implement `TableBorders` (w:tblBorders)
- [ ] 6.5.10 Implement `CellBorders` (w:tcBorders) with diagonal
- [ ] 6.5.11 Implement `TableWidth` (w:tblW)
- [ ] 6.5.12 Implement `CellWidth` (w:tcW)
- [ ] 6.5.13 Implement `VMerge` (vertical merge)
- [ ] 6.5.14 Implement `GridSpan` (w:gridSpan)
- [ ] 6.5.15 Implement `CellVAlign` (w:vAlign)
- [ ] 6.5.16 Implement `TableLook` (w:tblLook)
- [ ] 6.5.17 Implement `TableLayout` (w:tblLayout)
- [ ] 6.5.18 Implement `TableCellMargin` (w:tblCellMar)
- [ ] 6.5.19 Implement `TableInd` (w:tblInd)
- [ ] 6.5.20 Implement `TableJc` (w:jc table alignment)
- [ ] 6.5.21 Implement `TableOverlap` (w:tblOverlap)
- [ ] 6.5.22 Implement `TableStyle` (w:tblStyle)
- [ ] 6.5.23 Implement `TableStyleColBandSize`, `TableStyleRowBandSize`
- [ ] 6.5.24 Implement `CellMargins` (individual cell margins)
- [ ] 6.5.25 Implement `Shd` for table/cell shading
- [ ] 6.5.26 Implement `RowHeight` (w:trHeight)
- [ ] 6.5.27 Implement `TableHeader` (w:tblHeader)
- [ ] 6.5.28 Implement `CantSplit` (w:cantSplit)
- [ ] 6.5.29 Implement convenience `NewTable(rows, cols)` factory

### 6.5 Testing - Table Elements
- [ ] 6.5.30 Write unit tests for all table elements
- [ ] 6.5.31 Run tests and verify 100% pass rate
- [ ] 6.5.32 Integration test with previous phases

### 6.6 Bookmark, Hyperlink, and Anchor Elements (~20 types)
- [ ] 6.6.1 Implement `BookmarkStart` (w:bookmarkStart)
- [ ] 6.6.2 Implement `BookmarkEnd` (w:bookmarkEnd)
- [ ] 6.6.3 Implement `Hyperlink` (w:hyperlink)
- [ ] 6.6.4 Implement `AnchorId` attribute support
- [ ] 6.6.5 Implement `RangePermStart`, `RangePermEnd`
- [ ] 6.6.6 Implement `ProofErr` (w:proofErr)
- [ ] 6.6.7 Implement `CustomXmlDelRange` elements
- [ ] 6.6.8 Implement `MoveFromRangeStart`, `MoveFromRangeEnd`
- [ ] 6.6.9 Implement `MoveToRangeStart`, `MoveToRangeEnd`

### 6.6 Testing - Bookmark/Hyperlink Elements
- [ ] 6.6.10 Write unit tests for all bookmark/hyperlink elements
- [ ] 6.6.11 Run tests and verify 100% pass rate
- [ ] 6.6.12 Integration test with previous phases

### 6.7 Field Elements (~25 types)
- [ ] 6.7.1 Implement `SimpleField` (w:fldSimple)
- [ ] 6.7.2 Implement `FieldChar` (w:fldChar) - begin, separate, end
- [ ] 6.7.3 Implement `InstrText` (w:instrText)
- [ ] 6.7.4 Implement `FieldCode` element
- [ ] 6.7.5 Implement `FFData` (form field data)
- [ ] 6.7.6 Implement `TextInput` (w:textInput)
- [ ] 6.7.7 Implement `CheckBox` (w:checkBox)
- [ ] 6.7.8 Implement `DDList` (w:ddList - dropdown list)
- [ ] 6.7.9 Implement `FFDefault`, `FFMaxLength`, `FFFormat`
- [ ] 6.7.10 Implement `FFTextType`, `FFName`
- [ ] 6.7.11 Implement `FFHelpText`, `FFStatusText`
- [ ] 6.7.12 Implement `FFEnabled`, `FFCalcOnExit`

### 6.7 Testing - Field Elements
- [ ] 6.7.13 Write unit tests for all field elements
- [ ] 6.7.14 Run tests and verify 100% pass rate
- [ ] 6.7.15 Integration test with previous phases

### 6.8 Drawing and Image Elements (~40 types)
- [ ] 6.8.1 Implement `Drawing` (w:drawing)
- [ ] 6.8.2 Implement `Inline` positioning (wp:inline)
- [ ] 6.8.3 Implement `Anchor` positioning (wp:anchor)
- [ ] 6.8.4 Implement `DocPr` (wp:docPr)
- [ ] 6.8.5 Implement `CNvGraphicFramePr`
- [ ] 6.8.6 Implement `Extent` (wp:extent)
- [ ] 6.8.7 Implement `EffectExtent` (wp:effectExtent)
- [ ] 6.8.8 Implement `WrapNone`, `WrapSquare`, `WrapTight`, `WrapThrough`
- [ ] 6.8.9 Implement `WrapTopAndBottom`
- [ ] 6.8.10 Implement `SimplePos`, `PositionH`, `PositionV`
- [ ] 6.8.11 Implement `Graphic` (a:graphic)
- [ ] 6.8.12 Implement `GraphicData` (a:graphicData)
- [ ] 6.8.13 Implement `Pic` (pic:pic) - picture
- [ ] 6.8.14 Implement `BlipFill` (pic:blipFill)
- [ ] 6.8.15 Implement `Blip` (a:blip)
- [ ] 6.8.16 Implement `SpPr` (pic:spPr) - shape properties
- [ ] 6.8.17 Implement `Xfrm` (a:xfrm) - transform
- [ ] 6.8.18 Implement `PrstGeom` (a:prstGeom) - preset geometry
- [ ] 6.8.19 Implement `NoFill`, `SolidFill`, `GradFill`
- [ ] 6.8.20 Implement `Ln` (a:ln) - line/outline
- [ ] 6.8.21 Implement `NvPicPr`, `NvCxnSpPr`

### 6.8 Testing - Drawing/Image Elements
- [ ] 6.8.22 Write unit tests for all drawing elements
- [ ] 6.8.23 Run tests and verify 100% pass rate
- [ ] 6.8.24 Integration test with previous phases

### 6.9 Comment and Annotation Elements (~20 types)
- [ ] 6.9.1 Implement `CommentRangeStart` (w:commentRangeStart)
- [ ] 6.9.2 Implement `CommentRangeEnd` (w:commentRangeEnd)
- [ ] 6.9.3 Implement `CommentRef` (w:commentReference)
- [ ] 6.9.4 Implement `Comment` (w:comment)
- [ ] 6.9.5 Implement `Comments` root element (w:comments)
- [ ] 6.9.6 Implement `CommentEx` (w15:commentEx)
- [ ] 6.9.7 Implement `CommentsEx` (w15:commentsEx)
- [ ] 6.9.8 Implement `Annotation` base type
- [ ] 6.9.9 Implement `PermStart`, `PermEnd`

### 6.9 Testing - Comment/Annotation Elements
- [ ] 6.9.10 Write unit tests for all comment elements
- [ ] 6.9.11 Run tests and verify 100% pass rate
- [ ] 6.9.12 Integration test with previous phases

### 6.10 Revision/Track Changes Elements (~30 types)
- [ ] 6.10.1 Implement `Ins` (w:ins) - inserted content
- [ ] 6.10.2 Implement `Del` (w:del) - deleted content
- [ ] 6.10.3 Implement `DelText` (w:delText)
- [ ] 6.10.4 Implement `MoveFrom` (w:moveFrom)
- [ ] 6.10.5 Implement `MoveTo` (w:moveTo)
- [ ] 6.10.6 Implement `RunPropsChange` (w:rPrChange)
- [ ] 6.10.7 Implement `ParaPropsChange` (w:pPrChange)
- [ ] 6.10.8 Implement `SectPropsChange` (w:sectPrChange)
- [ ] 6.10.9 Implement `TablePropsChange` (w:tblPrChange)
- [ ] 6.10.10 Implement `TableRowPropsChange` (w:trPrChange)
- [ ] 6.10.11 Implement `CellPropsChange` (w:tcPrChange)
- [ ] 6.10.12 Implement `TableGridChange` (w:tblGridChange)
- [ ] 6.10.13 Implement `CustomXmlIns`, `CustomXmlDel`
- [ ] 6.10.14 Implement `CustomXmlMoveFrom`, `CustomXmlMoveTo`
- [ ] 6.10.15 Implement `NumberingChange` (w:numPrChange)

### 6.10 Testing - Revision Elements
- [ ] 6.10.16 Write unit tests for all revision elements
- [ ] 6.10.17 Run tests and verify 100% pass rate
- [ ] 6.10.18 Integration test with previous phases

### 6.11 Content Control (SDT) Elements (~25 types)
- [ ] 6.11.1 Implement `SdtBlock` (w:sdt for block-level)
- [ ] 6.11.2 Implement `SdtRun` (w:sdt for inline)
- [ ] 6.11.3 Implement `SdtCell` (w:sdt for table cells)
- [ ] 6.11.4 Implement `SdtRow` (w:sdt for table rows)
- [ ] 6.11.5 Implement `SdtPr` (w:sdtPr) - SDT properties
- [ ] 6.11.6 Implement `SdtContent` (w:sdtContent)
- [ ] 6.11.7 Implement `SdtEndPr`
- [ ] 6.11.8 Implement `Alias` (w:alias)
- [ ] 6.11.9 Implement `Tag` (w:tag)
- [ ] 6.11.10 Implement `Id` (w:id)
- [ ] 6.11.11 Implement `Lock` (w:lock)
- [ ] 6.11.12 Implement `Placeholder`
- [ ] 6.11.13 Implement `Temporary`
- [ ] 6.11.14 Implement `ShowingPlcHdr`
- [ ] 6.11.15 Implement `DataBinding`
- [ ] 6.11.16 Implement `Date`, `ComboBox`, `DropDownList`, `Picture`, `RichText`, `Text`, `Citation`, `Group`, `Bibliography`, `Equation`, `DocPartObj`, `DocPartList`

### 6.11 Testing - SDT Elements
- [ ] 6.11.17 Write unit tests for all SDT elements
- [ ] 6.11.18 Run tests and verify 100% pass rate
- [ ] 6.11.19 Integration test with previous phases

### 6.12 Symbol and Special Character Elements (~15 types)
- [ ] 6.12.1 Implement `Sym` (w:sym) - symbol
- [ ] 6.12.2 Implement `Ruby` (w:ruby) - ruby text
- [ ] 6.12.3 Implement `RubyPr` (w:rubyPr)
- [ ] 6.12.4 Implement `RubyBase`, `RubyContent`
- [ ] 6.12.5 Implement `Object` (w:object)
- [ ] 6.12.6 Implement `PTab` (w:ptab) - absolute position tab
- [ ] 6.12.7 Implement `FootnoteRef` (w:footnoteRef)
- [ ] 6.12.8 Implement `EndnoteRef` (w:endnoteRef)
- [ ] 6.12.9 Implement `Separator`, `ContinuationSeparator`
- [ ] 6.12.10 Implement `PgNum` (w:pgNum)

### 6.12 Testing - Symbol Elements
- [ ] 6.12.11 Write unit tests for all symbol elements
- [ ] 6.12.12 Run tests and verify 100% pass rate
- [ ] 6.12.13 Integration test with previous phases

### 6.13 Math Elements (Office Math - OMML, all 152 types)
**Note:** Full v1 implementation per design decision Q8.

#### 6.13.1 Core Math Containers
- [ ] 6.13.1.1 Implement `OMath` (m:oMath) - math zone
- [ ] 6.13.1.2 Implement `OMathPara` (m:oMathPara) - math paragraph
- [ ] 6.13.1.3 Implement `OMathParaPr` (m:oMathParaPr) - math paragraph properties
- [ ] 6.13.1.4 Implement `MathPr` (m:mathPr) - document math properties

#### 6.13.2 Math Run Elements
- [ ] 6.13.2.1 Implement `R` (m:r) - math run
- [ ] 6.13.2.2 Implement `T` (m:t) - math text
- [ ] 6.13.2.3 Implement `RPr` (m:rPr) - math run properties
- [ ] 6.13.2.4 Implement `Lit` (m:lit) - literal
- [ ] 6.13.2.5 Implement `Nor` (m:nor) - normal text
- [ ] 6.13.2.6 Implement `Brk` (m:brk) - break

#### 6.13.3 Accents and Bars
- [ ] 6.13.3.1 Implement `Acc` (m:acc) - accent
- [ ] 6.13.3.2 Implement `AccPr` (m:accPr) - accent properties
- [ ] 6.13.3.3 Implement `Bar` (m:bar) - bar/overline
- [ ] 6.13.3.4 Implement `BarPr` (m:barPr) - bar properties
- [ ] 6.13.3.5 Implement `Pos` (m:pos) - position

#### 6.13.4 Box Elements
- [ ] 6.13.4.1 Implement `Box` (m:box) - box
- [ ] 6.13.4.2 Implement `BoxPr` (m:boxPr) - box properties
- [ ] 6.13.4.3 Implement `BorderBox` (m:borderBox) - bordered box
- [ ] 6.13.4.4 Implement `BorderBoxPr` (m:borderBoxPr) - border box properties
- [ ] 6.13.4.5 Implement `HideBot`, `HideTop`, `HideLeft`, `HideRight` (border visibility)
- [ ] 6.13.4.6 Implement `StrikeBLTR`, `StrikeTLBR`, `StrikeH`, `StrikeV` (strikethrough)

#### 6.13.5 Delimiters
- [ ] 6.13.5.1 Implement `D` (m:d) - delimiter
- [ ] 6.13.5.2 Implement `DPr` (m:dPr) - delimiter properties
- [ ] 6.13.5.3 Implement `BegChr` (m:begChr) - beginning character
- [ ] 6.13.5.4 Implement `EndChr` (m:endChr) - ending character
- [ ] 6.13.5.5 Implement `SepChr` (m:sepChr) - separator character
- [ ] 6.13.5.6 Implement `Grow` (m:grow) - grow with content
- [ ] 6.13.5.7 Implement `Shp` (m:shp) - delimiter shape

#### 6.13.6 Fractions
- [ ] 6.13.6.1 Implement `F` (m:f) - fraction
- [ ] 6.13.6.2 Implement `FPr` (m:fPr) - fraction properties
- [ ] 6.13.6.3 Implement `Num` (m:num) - numerator
- [ ] 6.13.6.4 Implement `Den` (m:den) - denominator
- [ ] 6.13.6.5 Implement `FType` (m:type) - fraction type (bar, noBar, skewed, lin)

#### 6.13.7 Functions
- [ ] 6.13.7.1 Implement `Func` (m:func) - function application
- [ ] 6.13.7.2 Implement `FuncPr` (m:funcPr) - function properties
- [ ] 6.13.7.3 Implement `FName` (m:fName) - function name

#### 6.13.8 Grouping
- [ ] 6.13.8.1 Implement `GroupChr` (m:groupChr) - grouping character
- [ ] 6.13.8.2 Implement `GroupChrPr` (m:groupChrPr) - grouping properties
- [ ] 6.13.8.3 Implement `Chr` (m:chr) - character
- [ ] 6.13.8.4 Implement `VertJc` (m:vertJc) - vertical justification

#### 6.13.9 Limits
- [ ] 6.13.9.1 Implement `LimLow` (m:limLow) - lower limit
- [ ] 6.13.9.2 Implement `LimLowPr` (m:limLowPr) - lower limit properties
- [ ] 6.13.9.3 Implement `LimUpp` (m:limUpp) - upper limit
- [ ] 6.13.9.4 Implement `LimUppPr` (m:limUppPr) - upper limit properties
- [ ] 6.13.9.5 Implement `Lim` (m:lim) - limit argument

#### 6.13.10 Matrices
- [ ] 6.13.10.1 Implement `M` (m:m) - matrix
- [ ] 6.13.10.2 Implement `MPr` (m:mPr) - matrix properties
- [ ] 6.13.10.3 Implement `Mr` (m:mr) - matrix row
- [ ] 6.13.10.4 Implement `Mc` (m:mc) - matrix column
- [ ] 6.13.10.5 Implement `McPr` (m:mcPr) - matrix column properties
- [ ] 6.13.10.6 Implement `Mcs` (m:mcs) - matrix columns
- [ ] 6.13.10.7 Implement `Count` (m:count) - column count
- [ ] 6.13.10.8 Implement `RowSpacing`, `CGpRule`, `CGp` (spacing rules)
- [ ] 6.13.10.9 Implement `PlcHide` (m:plcHide) - placeholder hide
- [ ] 6.13.10.10 Implement `BaseJc` (m:baseJc) - matrix base justification

#### 6.13.11 N-ary Operations (Integrals, Summations)
- [ ] 6.13.11.1 Implement `Nary` (m:nary) - n-ary operator
- [ ] 6.13.11.2 Implement `NaryPr` (m:naryPr) - n-ary properties
- [ ] 6.13.11.3 Implement `Sub` (m:sub) - subscript/lower limit
- [ ] 6.13.11.4 Implement `Sup` (m:sup) - superscript/upper limit
- [ ] 6.13.11.5 Implement `SubHide` (m:subHide) - hide subscript
- [ ] 6.13.11.6 Implement `SupHide` (m:supHide) - hide superscript
- [ ] 6.13.11.7 Implement `LimLoc` (m:limLoc) - limit location

#### 6.13.12 Phantom Elements
- [ ] 6.13.12.1 Implement `Phant` (m:phant) - phantom
- [ ] 6.13.12.2 Implement `PhantPr` (m:phantPr) - phantom properties
- [ ] 6.13.12.3 Implement `Show` (m:show) - show phantom
- [ ] 6.13.12.4 Implement `ZeroWid` (m:zeroWid) - zero width
- [ ] 6.13.12.5 Implement `ZeroAsc`, `ZeroDesc` (m:zeroAsc/Desc) - zero ascent/descent
- [ ] 6.13.12.6 Implement `Transp` (m:transp) - transparent

#### 6.13.13 Radicals
- [ ] 6.13.13.1 Implement `Rad` (m:rad) - radical
- [ ] 6.13.13.2 Implement `RadPr` (m:radPr) - radical properties
- [ ] 6.13.13.3 Implement `Deg` (m:deg) - degree
- [ ] 6.13.13.4 Implement `DegHide` (m:degHide) - hide degree

#### 6.13.14 Scripts (Sub/Superscript)
- [ ] 6.13.14.1 Implement `SPre` (m:sPre) - pre-subscript/superscript
- [ ] 6.13.14.2 Implement `SPrePr` (m:sPrePr) - pre-script properties
- [ ] 6.13.14.3 Implement `SSub` (m:sSub) - subscript
- [ ] 6.13.14.4 Implement `SSubPr` (m:sSubPr) - subscript properties
- [ ] 6.13.14.5 Implement `SSup` (m:sSup) - superscript
- [ ] 6.13.14.6 Implement `SSupPr` (m:sSupPr) - superscript properties
- [ ] 6.13.14.7 Implement `SSubSup` (m:sSubSup) - sub-superscript
- [ ] 6.13.14.8 Implement `SSubSupPr` (m:sSubSupPr) - sub-superscript properties
- [ ] 6.13.14.9 Implement `AlnScr` (m:alnScr) - align scripts

#### 6.13.15 Equation Arrays
- [ ] 6.13.15.1 Implement `EqArr` (m:eqArr) - equation array
- [ ] 6.13.15.2 Implement `EqArrPr` (m:eqArrPr) - equation array properties
- [ ] 6.13.15.3 Implement `MaxDist` (m:maxDist) - maximum distribution
- [ ] 6.13.15.4 Implement `ObjDist` (m:objDist) - object distribution
- [ ] 6.13.15.5 Implement `RSpRule` (m:rSpRule) - row spacing rule
- [ ] 6.13.15.6 Implement `RSp` (m:rSp) - row spacing

#### 6.13.16 Base Elements
- [ ] 6.13.16.1 Implement `E` (m:e) - base element
- [ ] 6.13.16.2 Implement `ArgPr` (m:argPr) - argument properties
- [ ] 6.13.16.3 Implement `ArgSz` (m:argSz) - argument size

#### 6.13.17 Math Document Settings
- [ ] 6.13.17.1 Implement `MathFont` (m:mathFont) - math font
- [ ] 6.13.17.2 Implement `BrkBin` (m:brkBin) - break binary operator
- [ ] 6.13.17.3 Implement `BrkBinSub` (m:brkBinSub) - break binary subtraction
- [ ] 6.13.17.4 Implement `SmallFrac` (m:smallFrac) - small fractions
- [ ] 6.13.17.5 Implement `DispDef` (m:dispDef) - display default
- [ ] 6.13.17.6 Implement `LMargin`, `RMargin` (m:lMargin/rMargin) - margins
- [ ] 6.13.17.7 Implement `DefJc` (m:defJc) - default justification
- [ ] 6.13.17.8 Implement `PreSp`, `PostSp` (m:preSp/postSp) - spacing
- [ ] 6.13.17.9 Implement `InterSp` (m:interSp) - inter-equation spacing
- [ ] 6.13.17.10 Implement `IntraSp` (m:intraSp) - intra-equation spacing
- [ ] 6.13.17.11 Implement `WrapIndent`, `WrapRight` (m:wrapIndent/Right) - wrapping
- [ ] 6.13.17.12 Implement `IntLim`, `NaryLim` (m:intLim/naryLim) - limit locations

#### 6.13.18 Control Properties
- [ ] 6.13.18.1 Implement `CtrlPr` (m:ctrlPr) - control properties
- [ ] 6.13.18.2 Implement `Scr` (m:scr) - script
- [ ] 6.13.18.3 Implement `Sty` (m:sty) - style
- [ ] 6.13.18.4 Implement `Aln` (m:aln) - alignment

### 6.13 Testing - Math Elements
- [ ] 6.13.19.1 Write tests for simple math expressions (fractions, radicals)
- [ ] 6.13.19.2 Write tests for complex nested expressions (matrices with limits)
- [ ] 6.13.19.3 Write tests for n-ary operations (integrals, summations)
- [ ] 6.13.19.4 Write roundtrip tests with real Word documents containing math
- [ ] 6.13.19.5 Write tests for math properties and formatting
- [ ] 6.13.19.6 Run tests and verify 100% pass rate
- [ ] 6.13.19.7 Integration test with previous phases

### 6.14 Footnote and Endnote Elements (~15 types)
- [ ] 6.14.1 Implement `Footnotes` root (w:footnotes)
- [ ] 6.14.2 Implement `Footnote` (w:footnote)
- [ ] 6.14.3 Implement `Endnotes` root (w:endnotes)
- [ ] 6.14.4 Implement `Endnote` (w:endnote)
- [ ] 6.14.5 Implement `FootnoteRef`, `EndnoteRef`
- [ ] 6.14.6 Implement `FootnotePr`, `EndnotePr`
- [ ] 6.14.7 Implement `NumFmt`, `NumStart`, `NumRestart`
- [ ] 6.14.8 Implement `Pos` (w:pos) - position

### 6.14 Testing - Footnote/Endnote Elements
- [ ] 6.14.9 Write unit tests for all footnote/endnote elements
- [ ] 6.14.10 Run tests and verify 100% pass rate
- [ ] 6.14.11 Integration test with previous phases

### 6.15 Header and Footer Elements (~15 types)
- [ ] 6.15.1 Implement `Header` root (w:hdr)
- [ ] 6.15.2 Implement `Footer` root (w:ftr)
- [ ] 6.15.3 Implement `HeaderRef` (w:headerReference)
- [ ] 6.15.4 Implement `FooterRef` (w:footerReference)
- [ ] 6.15.5 Implement header/footer types: Default, First, Even
- [ ] 6.15.6 Implement `EvenAndOddHeaders` setting

### 6.15 Testing - Header/Footer Elements
- [ ] 6.15.7 Write unit tests for all header/footer elements
- [ ] 6.15.8 Run tests and verify 100% pass rate
- [ ] 6.15.9 Integration test with previous phases

### 6.16 Glossary Document Elements (~10 types)
- [ ] 6.16.1 Implement `GlossaryDocument` root (w:glossaryDocument)
- [ ] 6.16.2 Implement `DocParts` (w:docParts)
- [ ] 6.16.3 Implement `DocPart` (w:docPart)
- [ ] 6.16.4 Implement `DocPartPr` (w:docPartPr)
- [ ] 6.16.5 Implement `DocPartBody` (w:docPartBody)
- [ ] 6.16.6 Implement `Name`, `Style`, `Guid`, `Description`, `Category`, `Types`, `Behaviors`

### 6.16 Testing - Glossary Elements
- [ ] 6.16.7 Write unit tests for all glossary elements
- [ ] 6.16.8 Run tests and verify 100% pass rate
- [ ] 6.16.9 Integration test with previous phases

### 6.17 Miscellaneous Elements (~50+ types)
- [ ] 6.17.1 Implement `AltChunk` (w:altChunk) - alternative content
- [ ] 6.17.2 Implement `AltChunkPr` (w:altChunkPr)
- [ ] 6.17.3 Implement `CustomXml` (w:customXml)
- [ ] 6.17.4 Implement `CustomXmlPr` (w:customXmlPr)
- [ ] 6.17.5 Implement `SmartTag` (w:smartTag)
- [ ] 6.17.6 Implement `SmartTagPr` (w:smartTagPr)
- [ ] 6.17.7 Implement `SubDoc` (w:subDoc)
- [ ] 6.17.8 Implement `Dir` (w:dir) - bidirectional
- [ ] 6.17.9 Implement `Bdo` (w:bdo) - bidirectional override
- [ ] 6.17.10 Implement `Control` (w:control) - ActiveX
- [ ] 6.17.11 Implement all remaining WordprocessingML elements

### 6.17 Testing - Miscellaneous Elements
- [ ] 6.17.12 Write unit tests for all miscellaneous elements
- [ ] 6.17.13 Write comprehensive integration tests for element interactions
- [ ] 6.17.14 Run tests and verify 100% pass rate
- [ ] 6.17.15 Integration test with previous phases

### 6.18 Phase 6 Final Testing
- [ ] 6.18.1 Run all Phase 6 tests together
- [ ] 6.18.2 Verify 100% pass rate across all element categories
- [ ] 6.18.3 Integration test: create documents with formatting
- [ ] 6.18.4 Full integration test with Phases 1-5

## Phase 7: Styles and Numbering (pkg/word/styles/, pkg/word/numbering/)

### 7.1 Styles
- [ ] 7.1.1 Implement `Styles` root element (w:styles)
- [ ] 7.1.2 Implement `Style` element with all type support (paragraph, character, table, numbering)
- [ ] 7.1.3 Implement `DocDefaults` element (w:docDefaults)
- [ ] 7.1.4 Implement `LatentStyles` element (w:latentStyles)
- [ ] 7.1.5 Implement `TableStyleProps` for conditional formatting
- [ ] 7.1.6 Implement style inheritance resolution
- [ ] 7.1.7 Implement built-in style ID constants (Normal, Heading1-9, etc.)
- [ ] 7.1.8 Implement style factory methods (NewParaStyle, NewCharStyle, etc.)
- [ ] 7.1.9 Implement `LsdException` elements
- [ ] 7.1.10 Implement `StyleLink`, `NextStyle`, `BasedOn` relationships

### 7.2 Numbering
- [ ] 7.2.1 Implement `Numbering` root element (w:numbering)
- [ ] 7.2.2 Implement `AbstractNum` element (w:abstractNum)
- [ ] 7.2.3 Implement `Level` element (w:lvl) with all properties
- [ ] 7.2.4 Implement `NumInstance` element (w:num)
- [ ] 7.2.5 Implement `LevelOverride` element (w:lvlOverride)
- [ ] 7.2.6 Implement `NumFmt` enum (Decimal, LowerRoman, UpperRoman, LowerLetter, etc.)
- [ ] 7.2.7 Implement bullet list factory `NewBulletList()`
- [ ] 7.2.8 Implement numbered list factory `NewNumberedList()`
- [ ] 7.2.9 Implement outline numbering factory `NewOutlineList()`
- [ ] 7.2.10 Implement numbering ID generation (unique IDs)
- [ ] 7.2.11 Implement `MultiLevelType` enum

### 7.3 Phase 7 Testing
- [ ] 7.3.1 Write unit tests for all Styles elements
- [ ] 7.3.2 Write unit tests for style inheritance resolution
- [ ] 7.3.3 Write unit tests for all Numbering elements
- [ ] 7.3.4 Write unit tests for list factories
- [ ] 7.3.5 Run tests and verify 100% pass rate
- [ ] 7.3.6 Integration test: create styled documents with lists
- [ ] 7.3.7 Integration test with Phases 1-6

## Phase 8: Settings and Configuration (pkg/word/settings/)

### 8.1 Document Settings
- [ ] 8.1.1 Implement `Settings` root element (w:settings)
- [ ] 8.1.2 Implement `Zoom` settings (w:zoom)
- [ ] 8.1.3 Implement display settings (DisplayBackgroundShape, PrintFractionalWidth, etc.)
- [ ] 8.1.4 Implement proofing settings (HideSpellingErrors, HideGrammaticalErrors, ProofState)
- [ ] 8.1.5 Implement `Compat` settings (w:compat) with version-specific options
- [ ] 8.1.6 Implement track changes settings (TrackRevisions, RevisionView)
- [ ] 8.1.7 Implement `DocProtection` (w:documentProtection)
- [ ] 8.1.8 Implement `MailMerge` settings (w:mailMerge)
- [ ] 8.1.9 Implement `WriteProtection` (w:writeProtection)
- [ ] 8.1.10 Implement `DocVars` (w:docVars) - document variables
- [ ] 8.1.11 Implement `DefaultTabStop` (w:defaultTabStop)
- [ ] 8.1.12 Implement `Rsids` (w:rsids) - revision save IDs
- [ ] 8.1.13 Implement `ThemeFontLang` (w:themeFontLang)

### 8.2 Web Settings
- [ ] 8.2.1 Implement `WebSettings` root element (w:webSettings)
- [ ] 8.2.2 Implement browser optimization settings (OptimizeForBrowser, AllowPNG)
- [ ] 8.2.3 Implement `TargetScreenSz`, `Encoding`

### 8.3 Font Table
- [ ] 8.3.1 Implement `Fonts` root element (w:fonts)
- [ ] 8.3.2 Implement `Font` element (w:font)
- [ ] 8.3.3 Implement font properties (Family, Pitch, Charset, etc.)
- [ ] 8.3.4 Implement embedded font support (EmbedRegular, EmbedBold, etc.)

### 8.4 Phase 8 Testing
- [ ] 8.4.1 Write unit tests for all Document Settings
- [ ] 8.4.2 Write unit tests for Web Settings
- [ ] 8.4.3 Write unit tests for Font Table
- [ ] 8.4.4 Run tests and verify 100% pass rate
- [ ] 8.4.5 Integration test with Phases 1-7

## Phase 9: Headers, Footers, and References (pkg/word/parts/)

### 9.1 Headers and Footers (covered partially in Phase 6)
- [ ] 9.1.1 Implement `HeaderPart` with root element binding
- [ ] 9.1.2 Implement `FooterPart` with root element binding
- [ ] 9.1.3 Implement first/odd/even header/footer management
- [ ] 9.1.4 Implement section-specific header/footer assignment
- [ ] 9.1.5 Implement convenience methods `AddHeader()`, `AddFooter()` on Document

### 9.2 Footnotes and Endnotes (covered partially in Phase 6)
- [ ] 9.2.1 Implement `FootnotesPart` with root element binding
- [ ] 9.2.2 Implement `EndnotesPart` with root element binding
- [ ] 9.2.3 Implement footnote/endnote insertion convenience methods
- [ ] 9.2.4 Implement footnote/endnote numbering management

### 9.3 Comments (covered partially in Phase 6)
- [ ] 9.3.1 Implement `CommentsPart` with root element binding
- [ ] 9.3.2 Implement comment insertion convenience methods
- [ ] 9.3.3 Implement comment threading/reply support
- [ ] 9.3.4 Implement `CommentsExPart` for extended comments (Office 2016+)

### 9.4 Phase 9 Testing
- [ ] 9.4.1 Write unit tests for HeaderPart and FooterPart
- [ ] 9.4.2 Write unit tests for header/footer management
- [ ] 9.4.3 Write unit tests for FootnotesPart and EndnotesPart
- [ ] 9.4.4 Write unit tests for CommentsPart
- [ ] 9.4.5 Run tests and verify 100% pass rate
- [ ] 9.4.6 Integration test with Phases 1-8

## Phase 10: Advanced Features

### 10.1 Images and Media
- [ ] 10.1.1 Implement `ImagePart` with content type detection (JPEG, PNG, GIF, etc.)
- [ ] 10.1.2 Implement `AddImage(reader, contentType)` on MainPart
- [ ] 10.1.3 Implement inline image insertion convenience method
- [ ] 10.1.4 Implement floating image insertion (anchor positioning)
- [ ] 10.1.5 Implement image sizing and aspect ratio handling
- [ ] 10.1.6 Implement image cropping support

### 10.2 Document Builder Pattern
- [ ] 10.2.1 Implement `DocumentBuilder` for fluent document creation
- [ ] 10.2.2 Implement paragraph builder methods
- [ ] 10.2.3 Implement table builder methods
- [ ] 10.2.4 Implement `NewDefaultDocument()` factory

### 10.3 Macro-Enabled Documents
- [ ] 10.3.1 Implement `VbaPart` handling
- [ ] 10.3.2 Implement macro-enabled content types (.docm, .dotm)
- [ ] 10.3.3 Implement macro preservation during document manipulation

### 10.4 Template Support
- [ ] 10.4.1 Implement template document types (.dotx, .dotm)
- [ ] 10.4.2 Implement `NewFromTemplate()` with template attachment
- [ ] 10.4.3 Implement template detachment

### 10.5 Phase 10 Testing
- [ ] 10.5.1 Write unit tests for Images and Media
- [ ] 10.5.2 Write unit tests for Document Builder Pattern
- [ ] 10.5.3 Write unit tests for Macro-Enabled Documents
- [ ] 10.5.4 Write unit tests for Template Support
- [ ] 10.5.5 Run tests and verify 100% pass rate
- [ ] 10.5.6 Integration test with Phases 1-9

## Phase 11: Integration and Quality

### 11.1 Integration Tests
- [ ] 11.1.1 Create end-to-end tests creating documents from scratch
- [ ] 11.1.2 Create roundtrip tests (open, modify, save, reopen)
- [ ] 11.1.3 Test compatibility with Microsoft Word (Office 2016+)
- [ ] 11.1.4 Test compatibility with LibreOffice
- [ ] 11.1.5 Test with various Office 2016+ documents
- [ ] 11.1.6 Create fixture documents from multiple Office versions for testing
- [ ] 11.1.7 Test all validation constraints with invalid documents

### 11.2 Documentation
- [ ] 11.2.1 Write package-level godoc documentation
- [ ] 11.2.2 Create API examples for common tasks
- [ ] 11.2.3 Document relationship to Open-XML-SDK
- [ ] 11.2.4 Create migration guide from other Go libraries
- [ ] 11.2.5 Document idiomatic Go naming conventions vs C# SDK names

### 11.3 Performance
- [ ] 11.3.1 Benchmark document creation (target: 1MB in <100ms)
- [ ] 11.3.2 Benchmark large document handling (memory <10x file size)
- [ ] 11.3.3 Benchmark validation performance (target: <1s for typical docs)
- [ ] 11.3.4 Optimize hot paths based on benchmarks
- [ ] 11.3.5 Memory profiling and optimization
- [ ] 11.3.6 Ensure lazy loading works correctly for large documents

### 11.4 Enum Values (wml/ package)
- [ ] 11.4.1 Implement all JustificationValues (Left, Center, Right, Both, etc.)
- [ ] 11.4.2 Implement all UnderlineValues (Single, Double, Dotted, etc.)
- [ ] 11.4.3 Implement all HighlightColorValues
- [ ] 11.4.4 Implement all BorderValues
- [ ] 11.4.5 Implement all ShadingPatternValues
- [ ] 11.4.6 Implement all TabStopTypeValues
- [ ] 11.4.7 Implement all BreakValues
- [ ] 11.4.8 Implement all VerticalAlignRunValues
- [ ] 11.4.9 Implement all FontFamilyValues
- [ ] 11.4.10 Implement all remaining WordprocessingML enum types
- [ ] 11.4.11 Write enum validation functions

### 11.5 Phase 11 Final Testing
- [ ] 11.5.1 Run complete test suite across all phases
- [ ] 11.5.2 Verify 100% pass rate for all tests
- [ ] 11.5.3 Performance benchmarks meet targets
- [ ] 11.5.4 Full feature parity verification with Open-XML-SDK WordprocessingML

## Notes

**Unified Design Decisions Applied (consistent with Presentation SDK):**
- All 685+ element types are generated from JSON schemas (pre-generated, committed)
- Idiomatic Go naming: `Document` not `WordprocessingDocument`, `Para` not `Paragraph`
- Full validation with all 19+ semantic constraint types
- Office 2016+ only (ECMA-376 5th edition)

**Dependencies:**
- Phase 2 depends on Phase 1
- Phase 3 depends on Phase 2
- Phase 4 depends on Phase 3
- Phase 5 depends on Phases 3 & 4
- Phases 6-9 depend on Phase 5
- Phase 10 depends on Phases 6-9
- Phase 11 runs throughout and at end

**Parallelization:**
- Within each phase, tasks marked x.y.z can often run in parallel
- Different element type implementations (6.1-6.17) can be parallelized
- Test writing can parallel implementation
- Multiple element categories in Phase 6 can be worked on concurrently

**Validation Checkpoints:**
- After Phase 1: Project compiles, simple types work
- After Phase 2: Can create/open/save empty .docx (OPC compliance)
- After Phase 5: Can create basic documents with text
- After Phase 6 (partial): Can create documents with formatting
- After Phase 7: Can create styled documents with lists
- After Phase 11: Full feature parity with Open-XML-SDK WordprocessingML

**Estimated Scope:**
- ~685 element types to generate from JSON schemas
- ~29 simple types
- ~19 semantic constraint types
- ~60 part types
- ~100+ enum types in pkg/wml/ package
- Total estimated: 400+ Go files, 50,000+ lines of code (mostly generated)
