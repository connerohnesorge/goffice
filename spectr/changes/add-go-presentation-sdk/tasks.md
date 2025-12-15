# Tasks: Go SDK for Office Open XML Presentations

**Unified Design Decisions Applied:**
- Module path: `github.com/connerohnesorge/goffice`
- Code generation from JSON schemas (pre-generated, committed)
- Office 2016+ only (ECMA-376 5th edition+)
- Package structure: `pkg/{framework,types,package,relationships,validation,features,...}/`
- Go-idiomatic short API names (`Slide`, `Para`)
- Embedded struct fields for element metadata
- Functional options pattern for element construction
- Generic methods only for child access: `First[T]()`, `All[T]()`, `OfType[T]()`

## Phase 1: Project Foundation

### 1.1 Project Setup
- [ ] 1.1.1 Initialize Go module with `go mod init github.com/connerohnesorge/goffice`
- [ ] 1.1.2 Create directory structure (`pkg/`, `internal/`, `cmd/`)
- [ ] 1.1.3 Set up CI/CD pipeline (GitHub Actions)
- [ ] 1.1.4 Configure linting (golangci-lint) and formatting
- [ ] 1.1.5 Create initial README.md with project overview

### 1.2 Namespace Management (`pkg/framework/`)
- [ ] 1.2.1 Define `OpenXmlNamespace` struct with URI and default prefix
- [ ] 1.2.2 Define `OpenXmlQualifiedName` for namespace-qualified names
- [ ] 1.2.3 Create namespace registry with all known OpenXML namespaces
- [ ] 1.2.4 Implement `IOpenXmlNamespaceResolver` interface
- [ ] 1.2.5 Implement prefix lookup and registration
- [ ] 1.2.6 Handle Strict vs Transitional namespace mapping

### Phase 1 Testing
- [ ] 1.T.1 Write unit tests for namespace resolution
- [ ] 1.T.2 Verify project builds and linting passes
- [ ] 1.T.3 Verify CI/CD pipeline runs successfully

## Phase 2: Simple Types

### 2.1 Basic Value Types (`pkg/types/`)
- [ ] 2.1.1 Implement `StringValue` with nil-safe value handling
- [ ] 2.1.2 Implement `BooleanValue` with XML bool formats (true/false, 1/0)
- [ ] 2.1.3 Implement `Int32Value`, `Int64Value`, `IntegerValue`
- [ ] 2.1.4 Implement `UInt16Value`, `UInt32Value`, `UInt64Value`
- [ ] 2.1.5 Implement `DoubleValue`, `SingleValue`, `DecimalValue`
- [ ] 2.1.6 Implement `DateTimeValue` with ISO 8601 parsing
- [ ] 2.1.7 Implement `ByteValue`, `SByteValue`

### 2.2 Binary and Enum Types (`pkg/types/`)
- [ ] 2.2.1 Implement `HexBinaryValue` with hex encoding/decoding
- [ ] 2.2.2 Implement `Base64BinaryValue`
- [ ] 2.2.3 Implement generic `EnumValue[T]` for enumerated types
- [ ] 2.2.4 Implement `ListValue[T]` for space-separated values

### 2.3 Office-Specific Boolean Types (`pkg/types/`)
- [ ] 2.3.1 Implement `OnOffValue` (Office-style boolean: on/off/1/0)
- [ ] 2.3.2 Implement `TrueFalseValue` (t/f/true/false)
- [ ] 2.3.3 Implement `TrueFalseBlankValue` (with blank support)

### Phase 2 Testing
- [ ] 2.T.1 Write comprehensive unit tests for all basic value types
- [ ] 2.T.2 Write unit tests for binary and enum types
- [ ] 2.T.3 Write unit tests for Office-specific boolean types
- [ ] 2.T.4 Verify all simple types handle nil/empty values correctly

## Phase 3: OPC Packaging

### 3.1 OPC Package Abstraction (`internal/opc/`)
- [ ] 3.1.1 Define `IPackage` interface
- [ ] 3.1.2 Define `IPackagePart` interface
- [ ] 3.1.3 Define `IPackageRelationship` interface
- [ ] 3.1.4 Implement ZIP-based package using `archive/zip`
- [ ] 3.1.5 Implement part creation with content type
- [ ] 3.1.6 Implement part deletion
- [ ] 3.1.7 Implement relationship storage (`_rels/.rels`)
- [ ] 3.1.8 Implement `[Content_Types].xml` management

### 3.2 Package Properties (`pkg/package/`)
- [ ] 3.2.1 Define `IPackageProperties` interface
- [ ] 3.2.2 Implement core properties (title, author, created, modified, etc.)
- [ ] 3.2.3 Implement extended properties (application, company, etc.)
- [ ] 3.2.4 Implement custom properties

### 3.3 Part URI Handling (`pkg/package/`)
- [ ] 3.3.1 Implement URI parsing and validation
- [ ] 3.3.2 Implement unique URI generation
- [ ] 3.3.3 Implement relative URI resolution
- [ ] 3.3.4 Implement part naming conventions

### Phase 3 Testing
- [ ] 3.T.1 Write unit tests for OPC operations
- [ ] 3.T.2 Write unit tests for package properties
- [ ] 3.T.3 Write unit tests for URI handling
- [ ] 3.T.4 Verify ZIP package creation and extraction works correctly

## Phase 4: Framework

### 4.1 Attribute System (`pkg/framework/`)
- [ ] 4.1.1 Define `OpenXmlAttribute` struct
- [ ] 4.1.2 Implement attribute collection with fixed/extended separation
- [ ] 4.1.3 Implement `GetAttribute(localName, namespaceUri)` method
- [ ] 4.1.4 Implement `SetAttribute(attribute)` method
- [ ] 4.1.5 Implement `RemoveAttribute(localName, namespaceUri)` method
- [ ] 4.1.6 Implement `GetAttributes()` to get all attributes
- [ ] 4.1.7 Implement `ClearAllAttributes()` method

### 4.2 Element Base Types (`pkg/framework/`)
- [ ] 4.2.1 Define `OpenXmlElement` interface with core methods
- [ ] 4.2.2 Implement `BaseElement` struct with common functionality
- [ ] 4.2.3 Implement `Parent`, `FirstChild`, `LastChild` navigation
- [ ] 4.2.4 Implement `NextSibling`, `PreviousSibling` navigation
- [ ] 4.2.5 Implement `Ancestors()` iterator
- [ ] 4.2.6 Implement `Descendants()` iterator
- [ ] 4.2.7 Implement `Elements()` and `Elements[T]()` methods
- [ ] 4.2.8 Implement `GetFirstChild[T]()` method
- [ ] 4.2.9 Implement `AppendChild`, `PrependChild` methods
- [ ] 4.2.10 Implement `InsertBefore`, `InsertAfter` methods
- [ ] 4.2.11 Implement `InsertAt` method
- [ ] 4.2.12 Implement `RemoveChild`, `RemoveAllChildren` methods
- [ ] 4.2.13 Implement `ReplaceChild` method
- [ ] 4.2.14 Implement `Remove` (self-removal from parent)
- [ ] 4.2.15 Implement `CloneNode(deep bool)` method
- [ ] 4.2.16 Implement `LocalName`, `NamespaceUri`, `Prefix` properties
- [ ] 4.2.17 Implement `OuterXml`, `InnerXml`, `InnerText` properties
- [ ] 4.2.18 Implement namespace declaration management

### 4.3 Composite Element (`pkg/framework/`)
- [ ] 4.3.1 Implement `CompositeElement` struct embedding `BaseElement`
- [ ] 4.3.2 Implement linked list child storage
- [ ] 4.3.3 Implement `HasChildren` property
- [ ] 4.3.4 Implement child element iteration

### 4.4 Leaf Elements (`pkg/framework/`)
- [ ] 4.4.1 Implement `LeafElement` struct (no children)
- [ ] 4.4.2 Implement `LeafTextElement` struct (text content)
- [ ] 4.4.3 Implement `Text` property for leaf text elements

### 4.5 Feature Collection (`pkg/features/`)
- [ ] 4.5.1 Define `IFeatureCollection` interface
- [ ] 4.5.2 Implement `FeatureCollection` struct
- [ ] 4.5.3 Implement `Get[T]()` and `Set[T]()` methods
- [ ] 4.5.4 Implement feature inheritance from parent
- [ ] 4.5.5 Implement read-only feature collection wrapper

### 4.6 Element Metadata (`pkg/framework/metadata/`)
- [ ] 4.6.1 Define `IElementMetadata` interface
- [ ] 4.6.2 Implement `ElementMetadata` with schema information
- [ ] 4.6.3 Implement `AttributeMetadata` for attribute definitions
- [ ] 4.6.4 Implement child element constraints
- [ ] 4.6.5 Implement element factory registration

### 4.7 XML Serialization (`pkg/framework/`)
- [ ] 4.7.1 Implement `WriteTo(xmlWriter)` method
- [ ] 4.7.2 Implement `WriteContentTo(xmlWriter)` for child serialization
- [ ] 4.7.3 Implement `WriteAttributesTo(xmlWriter)` for attribute serialization
- [ ] 4.7.4 Implement `Load(xmlReader)` for parsing
- [ ] 4.7.5 Implement lazy parsing with `RawOuterXml`
- [ ] 4.7.6 Implement `MakeSureParsed()` for on-demand parsing
- [ ] 4.7.7 Handle markup compatibility attributes

### Phase 4 Testing
- [ ] 4.T.1 Write unit tests for attribute operations
- [ ] 4.T.2 Write unit tests for element operations
- [ ] 4.T.3 Write unit tests for composite element
- [ ] 4.T.4 Write unit tests for leaf elements
- [ ] 4.T.5 Write unit tests for feature collection
- [ ] 4.T.6 Write unit tests for metadata
- [ ] 4.T.7 Write integration tests for XML round-trip

## Phase 5: Validation

### 5.1 Validation Framework (`pkg/validation/`)
- [ ] 5.1.1 Define `ValidationError` struct
- [ ] 5.1.2 Define `ValidationErrorType` enum
- [ ] 5.1.3 Implement `ValidationContext` for tracking state
- [ ] 5.1.4 Implement `ValidationSettings` configuration
- [ ] 5.1.5 Define `FileFormatVersions` enum (Office 2016-2021, Microsoft 365)

### 5.2 OpenXmlValidator (`pkg/validation/`)
- [ ] 5.2.1 Implement `OpenXmlValidator` struct
- [ ] 5.2.2 Implement `Validate(element)` method
- [ ] 5.2.3 Implement `Validate(part)` method
- [ ] 5.2.4 Implement `Validate(document)` method
- [ ] 5.2.5 Implement configurable file format version targeting
- [ ] 5.2.6 Implement max errors limit

### 5.3 Schema Validation (`pkg/validation/`)
- [ ] 5.3.1 Implement element name validation
- [ ] 5.3.2 Implement required attribute validation
- [ ] 5.3.3 Implement attribute value validation
- [ ] 5.3.4 Implement child element sequence validation
- [ ] 5.3.5 Implement child element cardinality validation
- [ ] 5.3.6 Implement content model validation

### 5.4 Semantic Validation (`pkg/validation/`)
- [ ] 5.4.1 Implement relationship target validation
- [ ] 5.4.2 Implement ID reference validation
- [ ] 5.4.3 Implement cross-part reference validation
- [ ] 5.4.4 Implement presentation-specific semantic rules

### Phase 5 Testing
- [ ] 5.T.1 Write unit tests for validation types
- [ ] 5.T.2 Write unit tests for validator
- [ ] 5.T.3 Write unit tests for schema validation
- [ ] 5.T.4 Write unit tests for semantic validation
- [ ] 5.T.5 Verify validation errors are accurate and descriptive

## Phase 6: Presentation Document

### 6.1 PresentationDocument (`pkg/presentation/`)
- [ ] 6.1.1 Implement `PresentationDocument` struct
- [ ] 6.1.2 Implement `DocumentType` property (enum)
- [ ] 6.1.3 Implement `Create(path, type)` static method
- [ ] 6.1.4 Implement `Create(stream, type)` static method
- [ ] 6.1.5 Implement `CreateFromTemplate(path)` static method
- [ ] 6.1.6 Implement `Open(path, isEditable)` static method
- [ ] 6.1.7 Implement `Open(stream, isEditable)` static method
- [ ] 6.1.8 Implement `Open(path, isEditable, settings)` overload
- [ ] 6.1.9 Implement `ChangeDocumentType(newType)` method
- [ ] 6.1.10 Implement `AddPresentationPart()` method
- [ ] 6.1.11 Implement all Add*Part methods (CoreFileProperties, Extended, etc.)
- [ ] 6.1.12 Implement `PresentationPart` property getter

### 6.2 PresentationDocumentType (`pkg/presentation/`)
- [ ] 6.2.1 Define `PresentationDocumentType` enum with all types
- [ ] 6.2.2 Implement content type mapping for each type
- [ ] 6.2.3 Implement file extension mapping

### 6.3 FlatOPC Support (`pkg/presentation/`)
- [ ] 6.3.1 Implement `ToFlatOpcDocument()` method
- [ ] 6.3.2 Implement `ToFlatOpcString()` method
- [ ] 6.3.3 Implement `FromFlatOpcDocument(xdocument)` static method
- [ ] 6.3.4 Implement `FromFlatOpcString(xml)` static method

### Phase 6 Testing
- [ ] 6.T.1 Write unit tests for PresentationDocument
- [ ] 6.T.2 Write unit tests for document types
- [ ] 6.T.3 Write unit tests for FlatOPC conversion
- [ ] 6.T.4 Verify creating and opening presentations works correctly

## Phase 7: Presentation Parts

### 7.1 Core Presentation Parts (`pkg/presentation/parts/`)
- [ ] 7.1.1 Implement `PresentationPart`
- [ ] 7.1.2 Implement `SlidePart`
- [ ] 7.1.3 Implement `SlideLayoutPart`
- [ ] 7.1.4 Implement `SlideMasterPart`
- [ ] 7.1.5 Implement `NotesSlidePart`
- [ ] 7.1.6 Implement `NotesMasterPart`
- [ ] 7.1.7 Implement `HandoutMasterPart`

### 7.2 Theme and Style Parts (`pkg/presentation/parts/`)
- [ ] 7.2.1 Implement `ThemePart`
- [ ] 7.2.2 Implement `ThemeOverridePart`
- [ ] 7.2.3 Implement `TableStylesPart`

### 7.3 Comment Parts (`pkg/presentation/parts/`)
- [ ] 7.3.1 Implement `CommentAuthorsPart`
- [ ] 7.3.2 Implement `SlideCommentsPart`
- [ ] 7.3.3 Implement modern `CommentPart` (Office 365)

### 7.4 Media Parts (`pkg/presentation/parts/`)
- [ ] 7.4.1 Implement `ImagePart` with all image types
- [ ] 7.4.2 Implement `AudioPart` (embedded audio)
- [ ] 7.4.3 Implement `VideoPart` (embedded video)
- [ ] 7.4.4 Implement `FontPart` (embedded fonts)
- [ ] 7.4.5 Implement `ThumbnailPart`

### 7.5 Embedded Object Parts (`pkg/presentation/parts/`)
- [ ] 7.5.1 Implement `EmbeddedObjectPart`
- [ ] 7.5.2 Implement `EmbeddedPackagePart`
- [ ] 7.5.3 Implement `EmbeddedControlPersistencePart`
- [ ] 7.5.4 Implement `OleObjectPart`

### 7.6 Custom XML Parts (`pkg/presentation/parts/`)
- [ ] 7.6.1 Implement `CustomXmlPart`
- [ ] 7.6.2 Implement `CustomXmlPropertiesPart`

### 7.7 VBA Parts (`pkg/presentation/parts/`)
- [ ] 7.7.1 Implement `VbaProjectPart`
- [ ] 7.7.2 Implement `VbaDataPart`

### 7.8 Additional Parts (`pkg/presentation/parts/`)
- [ ] 7.8.1 Implement `ViewPropertiesPart`
- [ ] 7.8.2 Implement `PresentationPropertiesPart`
- [ ] 7.8.3 Implement `TagsPart` (user-defined tags)
- [ ] 7.8.4 Implement `UserDefinedTagsPart`
- [ ] 7.8.5 Implement `SlideShowTimingsPart`

### Phase 7 Testing
- [ ] 7.T.1 Write unit tests for core parts
- [ ] 7.T.2 Write unit tests for theme parts
- [ ] 7.T.3 Write unit tests for comment parts
- [ ] 7.T.4 Write unit tests for media parts
- [ ] 7.T.5 Write unit tests for embedded parts
- [ ] 7.T.6 Write unit tests for custom XML parts
- [ ] 7.T.7 Write unit tests for VBA parts
- [ ] 7.T.8 Write unit tests for additional parts

## Phase 8: Code Generator

### 8.1 Schema Parser (`cmd/goffice-gen/`)
- [ ] 8.1.1 Parse ECMA-376 XSD schemas
- [ ] 8.1.2 Extract complex type definitions
- [ ] 8.1.3 Extract simple type definitions
- [ ] 8.1.4 Extract element declarations
- [ ] 8.1.5 Extract attribute declarations
- [ ] 8.1.6 Handle schema imports and includes

### 8.2 Go Code Generator (`cmd/goffice-gen/`)
- [ ] 8.2.1 Generate element struct definitions
- [ ] 8.2.2 Generate property getter/setter methods
- [ ] 8.2.3 Generate constructor functions
- [ ] 8.2.4 Generate `CloneNode` implementations
- [ ] 8.2.5 Generate XML tag constants
- [ ] 8.2.6 Generate validation rules
- [ ] 8.2.7 Generate documentation from schema annotations

### 8.3 Generated Schema Types (`pkg/presentation/schema/`)
- [ ] 8.3.1 Generate PresentationML types (`p:` namespace)
- [ ] 8.3.2 Generate DrawingML types (`a:` namespace)
- [ ] 8.3.3 Generate RelationshipML types (`r:` namespace)
- [ ] 8.3.4 Generate Office extension types

### Phase 8 Testing
- [ ] 8.T.1 Write unit tests for schema parser
- [ ] 8.T.2 Write unit tests for code generator output
- [ ] 8.T.3 Write integration tests for generated types
- [ ] 8.T.4 Verify generated code compiles and passes linting

## Phase 9: Presentation Elements

### 9.1 Relationships (`pkg/relationships/`)
- [ ] 9.1.1 Implement `Relationship` struct
- [ ] 9.1.2 Implement `ExternalRelationship` for external targets
- [ ] 9.1.3 Implement `HyperlinkRelationship`
- [ ] 9.1.4 Implement `DataPartReferenceRelationship` (audio/video/media)
- [ ] 9.1.5 Implement relationship ID generation
- [ ] 9.1.6 Implement relationship type constants

### 9.2 OpenXmlPart (`pkg/package/`)
- [ ] 9.2.1 Define `OpenXmlPart` interface
- [ ] 9.2.2 Implement `BasePart` struct with common functionality
- [ ] 9.2.3 Implement `ContentType` property
- [ ] 9.2.4 Implement `Uri` property
- [ ] 9.2.5 Implement `OpenXmlPackage` back-reference
- [ ] 9.2.6 Implement `GetStream()` for reading part content
- [ ] 9.2.7 Implement `FeedData(stream)` for writing content
- [ ] 9.2.8 Implement `RootElement` property with lazy loading
- [ ] 9.2.9 Implement part relationships management

### 9.3 OpenXmlPartContainer (`pkg/package/`)
- [ ] 9.3.1 Implement `PartContainer` interface
- [ ] 9.3.2 Implement `AddPart[T]()` method
- [ ] 9.3.3 Implement `AddNewPart[T]()` method
- [ ] 9.3.4 Implement `DeletePart()` method
- [ ] 9.3.5 Implement `GetPartById()` method
- [ ] 9.3.6 Implement `GetPartsOfType[T]()` method
- [ ] 9.3.7 Implement `GetIdOfPart()` method
- [ ] 9.3.8 Implement `ChangeIdOfPart()` method

### 9.4 OpenXmlPackage (`pkg/package/`)
- [ ] 9.4.1 Implement `OpenXmlPackage` struct
- [ ] 9.4.2 Implement `RootPart` property
- [ ] 9.4.3 Implement `FileOpenAccess` property
- [ ] 9.4.4 Implement `AutoSave` property and behavior
- [ ] 9.4.5 Implement `Save()` method
- [ ] 9.4.6 Implement `Close()` and `Dispose()` methods
- [ ] 9.4.7 Implement `DataParts` collection
- [ ] 9.4.8 Implement `CreateMediaDataPart()` methods
- [ ] 9.4.9 Implement `CompressionOption` setting

### 9.5 DataPart (`pkg/package/`)
- [ ] 9.5.1 Implement `DataPart` struct
- [ ] 9.5.2 Implement `MediaDataPart` for media content
- [ ] 9.5.3 Implement content type detection
- [ ] 9.5.4 Implement stream access

### 9.6 Chart and Diagram Parts (`pkg/presentation/parts/`)
- [ ] 9.6.1 Implement `ChartPart`
- [ ] 9.6.2 Implement `ChartColorStylePart`
- [ ] 9.6.3 Implement `ChartStylePart`
- [ ] 9.6.4 Implement `DiagramColorsPart`
- [ ] 9.6.5 Implement `DiagramDataPart`
- [ ] 9.6.6 Implement `DiagramLayoutDefinitionPart`
- [ ] 9.6.7 Implement `DiagramPersistLayoutPart`
- [ ] 9.6.8 Implement `DiagramStylePart`

### Phase 9 Testing
- [ ] 9.T.1 Write unit tests for relationships
- [ ] 9.T.2 Write unit tests for parts
- [ ] 9.T.3 Write unit tests for part container
- [ ] 9.T.4 Write unit tests for package operations
- [ ] 9.T.5 Write unit tests for data parts
- [ ] 9.T.6 Write unit tests for chart/diagram parts

## Phase 10: Advanced Features

### 10.1 API Documentation
- [ ] 10.1.1 Add godoc comments to all public types
- [ ] 10.1.2 Add godoc comments to all public methods
- [ ] 10.1.3 Add package-level documentation
- [ ] 10.1.4 Generate API reference documentation

### 10.2 Examples
- [ ] 10.2.1 Create "Hello World" presentation example
- [ ] 10.2.2 Create slide manipulation example
- [ ] 10.2.3 Create shape creation example
- [ ] 10.2.4 Create text formatting example
- [ ] 10.2.5 Create image insertion example
- [ ] 10.2.6 Create chart creation example
- [ ] 10.2.7 Create table creation example
- [ ] 10.2.8 Create animation example
- [ ] 10.2.9 Create reading/modifying existing presentation example
- [ ] 10.2.10 Create validation example

### Phase 10 Testing
- [ ] 10.T.1 Verify all examples compile and run correctly
- [ ] 10.T.2 Verify API documentation is complete and accurate
- [ ] 10.T.3 Review examples for clarity and best practices

## Phase 11: Integration and Quality

### 11.1 Integration Testing
- [ ] 11.1.1 Add integration tests with real .pptx files
- [ ] 11.1.2 Add round-trip tests (open -> modify -> save -> reopen)
- [ ] 11.1.3 Add compatibility tests against Office-generated files
- [ ] 11.1.4 Add performance benchmarks

### 11.2 Code Quality
- [ ] 11.2.1 Achieve >80% code coverage
- [ ] 11.2.2 Run full linting suite and fix all issues
- [ ] 11.2.3 Review all public APIs for consistency
- [ ] 11.2.4 Ensure all error messages are clear and actionable

### 11.3 Release Preparation
- [ ] 11.3.1 Update README with comprehensive usage guide
- [ ] 11.3.2 Create CHANGELOG.md
- [ ] 11.3.3 Set up semantic versioning
- [ ] 11.3.4 Prepare v1.0.0 release

### Phase 11 Testing
- [ ] 11.T.1 Run full test suite and verify all tests pass
- [ ] 11.T.2 Verify code coverage meets requirements
- [ ] 11.T.3 Perform final manual testing with various .pptx files
- [ ] 11.T.4 Verify package can be installed and used as a dependency
