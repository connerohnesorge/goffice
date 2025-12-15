# Change: Add Go SDK for Office Open XML Presentations

## Why

The Office Open XML format is the standard for modern Microsoft Office documents (.pptx, .potx, .ppsx). Currently, Go lacks a comprehensive, well-structured SDK for creating, reading, and manipulating PowerPoint presentations that mirrors the capabilities of the official Microsoft Open-XML-SDK for .NET. This proposal introduces **goffice**, a Go SDK that replicates all features present in the referenced C# SDK, providing Go developers with first-class support for presentation document manipulation.

## What Changes

### New Capabilities

1. **Framework** (`framework`)
   - Base `OpenXmlElement` type with full DOM traversal (parent, children, siblings, ancestors, descendants)
   - `OpenXmlCompositeElement` for container elements with child management
   - `OpenXmlLeafElement` and `OpenXmlLeafTextElement` for terminal elements
   - Attribute system with fixed (schema-defined) and extended attributes
   - Namespace management and prefix resolution
   - XML serialization/deserialization with markup compatibility support
   - Element cloning and deep copy capabilities

2. **Types** (`types`)
   - All OpenXML simple value types: `StringValue`, `BooleanValue`, `Int32Value`, `Int64Value`, `UInt32Value`, `DoubleValue`, `DecimalValue`, `DateTimeValue`
   - `EnumValue[T]` with string/int conversion
   - `HexBinaryValue` and `Base64BinaryValue`
   - `ListValue[T]` for space-separated lists
   - `OnOffValue`, `TrueFalseValue`, `TrueFalseBlankValue` for Office-specific booleans
   - Implicit conversion operators (Go methods) for ergonomic usage

3. **Package** (`package`)
   - `OpenXmlPackage` base type wrapping ZIP-based OPC packages
   - `OpenXmlPart` abstract part type with content type and relationship handling
   - `OpenXmlPartContainer` with part management (add, remove, get)
   - `DataPart` and `MediaDataPart` for binary content
   - Package properties (core, extended, custom)
   - Part URI generation and uniqueness enforcement
   - Stream-based and file-based package access
   - Auto-save and manual save support

4. **Relationships** (`relationships`)
   - Relationship management (internal, external, hyperlink)
   - Part-to-part relationships
   - External relationship support
   - Hyperlink relationship handling

5. **Validation** (`validation`)
   - `OpenXmlValidator` with configurable file format versions
   - Schema validation against ECMA-376 / ISO/IEC 29500
   - Semantic validation rules
   - Validation error collection with path information
   - Document-level and part-level validation

6. **Features** (`features`)
   - Feature collection pattern for extensibility
   - Three-level hierarchy: Element/Part → Package → Global defaults
   - Lazy initialization via `GetKnown()` pattern

7. **Presentation Document** (`presentation-document`)
   - `PresentationDocument` type supporting all presentation formats:
     - Presentation (`.pptx`)
     - Template (`.potx`)
     - Slideshow (`.ppsx`)
     - MacroEnabledPresentation (`.pptm`)
     - MacroEnabledTemplate (`.potm`)
     - MacroEnabledSlideshow (`.ppsm`)
     - AddIn (`.ppam`)
   - Create from file, stream, or template
   - Open existing documents (read-only or editable)
   - Document type conversion
   - FlatOPC format support (single XML representation)

8. **Presentation Parts** (`presentation-parts`)
   - `PresentationPart` - main presentation content
   - `SlidePart`, `SlideLayoutPart`, `SlideMasterPart`
   - `NotesSlidePart`, `NotesMasterPart`, `HandoutMasterPart`
   - `ThemePart`, `TableStylesPart`
   - `CommentsPart`, `CommentAuthorsPart`
   - `ChartPart`, `DiagramPart`, `EmbeddedObjectPart`
   - `ImagePart`, `MediaPart` (audio/video)
   - `CustomXmlPart`, `VbaProjectPart`
   - `FontPart`, `ThumbnailPart`
   - Part constraint enforcement

9. **Presentation Elements** (`presentation-elements`)
   - All PresentationML schema elements from ECMA-376
   - Strongly-typed element classes with property accessors
   - Slide structure: `Presentation`, `SlideIdList`, `Slide`, `CommonSlideData`
   - Shape tree: `ShapeTree`, `Shape`, `TextBody`, `Paragraph`, `Run`
   - Graphics: `Picture`, `GraphicFrame`, `GroupShape`, `ConnectionShape`
   - Tables: `Table`, `TableRow`, `TableCell`
   - Charts: `ChartSpace`, `Chart`, `PlotArea`, `Series`
   - Animations: `Timing`, `TimeNodeList`, `ParallelTimeNode`, `AnimateEffect`
   - Transitions: `Transition`, various transition effects
   - Comments and notes
   - DrawingML integration for shapes and formatting

### BREAKING Changes
- None (new project)

### Non-Breaking Changes
- Introduces new `goffice` module with multiple packages

## Impact

- **Affected specs**: None (all new)
- **Affected code**: Creates new `/pkg/` directory structure with all SDK packages
- **Dependencies**: Pure Go stdlib only
  - `archive/zip` (standard library)
  - `encoding/xml` (standard library)
  - `sync` (standard library for Mutex)

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
| **XML Prefixes** | Fixed canonical (p:, a:, r:) |
| **Package Structure** | Flat by namespace (pml/, dml/, rml/) |
| **Version Support** | Office 2016+ (FileFormatVersions.Office2016) |

## Architecture Overview

```
goffice/
├── pkg/
│   ├── framework/           # Core element types
│   │   ├── element.go       # OpenXmlElement interface & BaseElement
│   │   ├── composite.go     # CompositeElement (has children)
│   │   ├── leaf.go          # LeafElement, LeafTextElement
│   │   └── attribute.go     # Attribute handling (fixed + extended)
│   │
│   ├── types/               # OpenXML value types
│   │   ├── string.go        # StringValue
│   │   ├── bool.go          # BooleanValue, OnOffValue, TrueFalseValue
│   │   ├── numeric.go       # Int32Value, Int64Value, DoubleValue, etc.
│   │   ├── enum.go          # EnumValue[T] generic
│   │   └── hex.go           # HexBinaryValue, Base64BinaryValue
│   │
│   ├── package/             # OPC package handling
│   │   ├── package.go       # OpenXmlPackage
│   │   ├── part.go          # OpenXmlPart base
│   │   ├── container.go     # OpenXmlPartContainer
│   │   └── properties.go    # Core, extended, custom properties
│   │
│   ├── relationships/       # Relationship management
│   │   └── relationship.go  # Relationships (internal, external, hyperlink)
│   │
│   ├── validation/          # Document validation
│   │   ├── validator.go     # OpenXmlValidator
│   │   ├── schema.go        # Schema validation (particle-based)
│   │   └── semantic.go      # Semantic constraints
│   │
│   ├── features/            # Feature collection pattern
│   │   └── features.go      # IFeatureCollection, type-safe Get/Set
│   │
│   ├── presentation/        # PresentationDocument & types
│   │   ├── document.go      # PresentationDocument
│   │   ├── types.go         # PresentationDocumentType enum
│   │   └── parts/           # All presentation part types
│   │
│   ├── pml/                 # PresentationML elements (p: namespace)
│   │   └── *.go             # Pre-generated from JSON schemas
│   │
│   ├── dml/                 # DrawingML elements (a: namespace)
│   │   └── *.go             # Pre-generated from JSON schemas
│   │
│   └── rml/                 # RelationshipML elements (r: namespace)
│       └── *.go             # Pre-generated from JSON schemas
│
├── internal/
│   ├── opc/                 # Low-level OPC/ZIP implementation
│   └── xml/                 # XML utilities, MC processing
│
└── testdata/
    └── golden/              # Expected XML for golden file tests
```

## Compatibility Goals

- Full API parity with Open-XML-SDK for presentation features
- Same document fidelity (round-trip without loss)
- Compatible content types and relationship types
- Support for Office 2016, 2019, 2021, and Microsoft 365 formats (ECMA-376 5th edition+)

## Unified Design Decisions

These decisions apply to both Presentation and Word SDKs for consistency:

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Module Path** | `github.com/connerohnesorge/goffice` | Personal namespace, easy to start |
| **Code Generation** | Generate from JSON schemas | Reuse C# SDK schema data, 1000+ types |
| **Office Version** | 2016+ only | Modern documents, reduced complexity |
| **Package Structure** | `pkg/` style | `pkg/{framework,types,package,relationships,validation,features,...}` |
| **API Naming** | Go-idiomatic short | `Document`, `Para`, `Slide` (not verbose C# names) |
| **Element Metadata** | Embedded struct fields | Self-contained elements, no global registry |
| **Construction Pattern** | Functional options | `NewSlide(WithTitle("Hello"))` |
| **Child Access** | Generic methods only | `First[T]()`, `All[T]()`, `OfType[T]()` |
