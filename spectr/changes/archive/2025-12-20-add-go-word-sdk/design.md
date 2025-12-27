# Design: Go SDK for Word Document Processing

## Context

This design document captures architectural decisions for implementing a Go SDK for Office Open XML (OOXML) Word documents, modeled after Microsoft's Open-XML-SDK for C#. The SDK must handle the full complexity of the OOXML specification while providing idiomatic Go APIs.

**Stakeholders:**
- Go developers needing Word document generation/manipulation
- Enterprise applications requiring server-side document processing
- Migration from .NET Open-XML-SDK to Go

**Constraints:**
- Must support ECMA-376 and ISO/IEC 29500 standards
- Target Office 2016+ documents (ECMA-376 5th edition and later)
- Must work without Microsoft Office installed
- Standard library only for core functionality (no CGO)

**Decided Design Choices (Unified with Presentation SDK):**
- Code generation from JSON schemas (same source as C# SDK) - pre-generated, committed to repo
- Idiomatic Go naming conventions (not mirroring C# names)
- Full validation: schema + all 19+ semantic constraint types - strict by default
- Office 2016+ only (reduces version complexity)
- Functional options pattern for element construction
- Generic methods only for child access: `First[T]()`, `All[T]()`, `OfType[T]()`
- DrawingML as shared package used by Word, Presentation, and Spreadsheet SDKs
- Phase 1 scope: Full feature parity with Open-XML-SDK (not MVP)

## Goals / Non-Goals

### Goals
- Feature parity with Open-XML-SDK for WordprocessingML
- Idiomatic Go API design (interfaces, composition, error handling)
- Strong typing for all elements and attributes
- Comprehensive validation (schema + semantic)
- Support all Word document types (.docx, .dotx, .docm, .dotm)
- Lazy loading for performance with large documents
- Thread-safe read operations
- Extensible architecture via feature collections

### Non-Goals
- SpreadsheetML (Excel) support - future proposal
- PresentationML (PowerPoint) support - separate proposal: add-go-presentation-sdk
- Visual Basic for Applications (VBA) macro execution
- Document rendering/preview generation
- PDF conversion
- Real-time collaborative editing support
- LINQ-style query support (removed from scope)

## Confirmed Design Decisions (ULTRATHINK Approved)

These decisions apply consistently across all goffice SDKs (Word, Presentation, Spreadsheet, DrawingML):

### Core Architecture

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Module Path** | `github.com/connerohnesorge/goffice` | Single unified module, simple imports, atomic versioning |
| **Go Version** | 1.25+ | Range-over-func for iterators, modern generics features |
| **Module Layout** | Single unified go.mod | Shared types, atomic versioning, simple dependency graph |
| **Code Generation** | JSON Schema → Go | Parse Open-XML-SDK's JSON schema files, generate Go structs with XML tags |
| **Office Version** | 2016+ only (ECMA-376 5th edition+) | Modern documents, reduced complexity, covers 95%+ of real-world files |
| **Thread Safety** | Per-Document RWMutex | Single sync.RWMutex per OpenXmlPackage, simple and sufficient |
| **Dependencies** | Pure stdlib only | archive/zip, encoding/xml, sync - no external dependencies |

### Element System

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Package Structure** | Top-level packages | `drawingml/`, `packaging/`, `openxml/`, `wordprocessing/` |
| **Element Metadata** | Embedded struct fields | XMLName, Namespace, LocalName as struct fields, self-contained, works with encoding/xml |
| **Construction Pattern** | Functional options | `NewPara(WithText("Hello"), WithBold(true))` - idiomatic Go, extensible, self-documenting |
| **Child Access** | Generic functions | `First[T](el)`, `All[T](el)`, `OfType[T](el)` using Go 1.18+ generics with range-over-func iterators |
| **API Naming** | Go-idiomatic short names | `Para`, `ParaProps`, `Run` (not verbose `ParagraphProperties`) |

### Validation System

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Validation Strategy** | Code-generated Validate() methods | Generate type-specific Validate() during code gen, zero reflection, compile-time type safety |
| **Error Handling** | Structured errors | Custom error types (ValidationError, ParseError) with path, element, constraint info, errors.Is/As compatible |
| **Validation Mode** | Strict by default | Reject invalid content, return errors for malformed documents |

### Extensibility

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Feature Collection** | Interface Registry pattern | Define Feature interface, register by interface type, supports hierarchy (Element→Part→Package→Global) with fallback chain |
| **DrawingML** | Shared `drawingml/` package | Used by all three SDKs (Word, Presentation, Spreadsheet) for shapes, images, effects |

### XML Processing

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **XML Prefixes** | Fixed canonical | Always use canonical prefixes: w: for WordML, a: for DrawingML, r: for relationships |
| **MC Handling** | Parse-time processing | Process AlternateContent/Choice/Fallback during XML parsing based on target FileFormatVersion |
| **Namespace Management** | Embedded in types | Each generated type knows its namespace URI and local name |

### Implementation Scope

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Phase 1 Scope** | Full feature parity with Open-XML-SDK | Complete implementation, not an MVP |

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                        User Application                          │
└─────────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────────┐
│                        word package                              │
│  ┌─────────────────┐  ┌──────────────┐  ┌───────────────────┐  │
│  │ Document        │  │ Parts        │  │ Elements          │  │
│  │ - New()         │  │ - Main       │  │ - Doc (root)      │  │
│  │ - Open()        │  │ - Styles     │  │ - Body            │  │
│  │ - Save()        │  │ - Numbering  │  │ - Para, Run       │  │
│  │ - Close()       │  │ - Settings   │  │ - Table, Cell     │  │
│  └─────────────────┘  └──────────────┘  └───────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────────┐
│                      framework package                           │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐  │
│  │ Element      │  │ Composite    │  │ Leaf                 │  │
│  │ - Parent     │  │ - Children   │  │ - InnerText          │  │
│  │ - LocalName  │  │ - Append     │  │ - SetInnerText       │  │
│  │ - Attributes │  │ - Remove     │  │                      │  │
│  │ - Clone()    │  │ - First/Last │  │                      │  │
│  └──────────────┘  └──────────────┘  └──────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────────┐
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐  │
│  │ types        │  │ relationships│  │ validation           │  │
│  │ - StringValue│  │ - Internal   │  │ - SchemaValidator    │  │
│  │ - Int32Value │  │ - External   │  │ - SemanticValidator  │  │
│  │ - EnumValue  │  │ - Hyperlinks │  │ - Constraints        │  │
│  └──────────────┘  └──────────────┘  └──────────────────────┘  │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐  │
│  │ package      │  │ features     │  │                      │  │
│  │ - Open()     │  │ - Collection │  │                      │  │
│  │ - Create()   │  │ - Cascade    │  │                      │  │
│  │ - Save()     │  │ - Configure  │  │                      │  │
│  │ - Parts      │  │              │  │                      │  │
│  └──────────────┘  └──────────────┘  └──────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────────┐
│                     archive/zip (stdlib)                         │
└─────────────────────────────────────────────────────────────────┘
```

## Decisions

### D1: Package Structure

**Decision:** Organize into top-level packages (consistent with Presentation SDK).

```
goffice/
├── drawingml/               # SHARED - DrawingML types (shapes, images, effects)
│   ├── types.go             # Core DrawingML types
│   ├── shapes.go            # Shape definitions
│   ├── effects.go           # Visual effects
│   └── ...                  # Used by Word, Presentation, and Spreadsheet
│
├── packaging/               # OPC layer (Open Packaging Conventions)
│   ├── package.go           # OpenXmlPackage
│   ├── part.go              # OpenXmlPart base
│   ├── container.go         # OpenXmlPartContainer
│   ├── relationships.go     # Relationship management
│   └── contenttypes.go      # Content type handling
│
├── openxml/                 # Core framework
│   ├── element.go           # Element interface & BaseElement
│   ├── composite.go         # CompositeElement (has children)
│   ├── leaf.go              # LeafElement, LeafTextElement
│   ├── attribute.go         # Attribute handling
│   ├── types/               # OpenXML value types
│   │   ├── string.go        # StringValue
│   │   ├── bool.go          # BooleanValue, OnOffValue
│   │   ├── numeric.go       # Int32Value, Int64Value, etc.
│   │   └── enum.go          # EnumValue[T] generic
│   ├── features/            # Feature collection pattern
│   │   ├── collection.go    # FeatureCollection
│   │   └── interfaces.go    # Feature interfaces
│   └── validation/          # Document validation
│       ├── validator.go     # OpenXmlValidator
│       ├── schema.go        # Schema validation
│       └── semantic.go      # Semantic constraints (19+ types)
│
├── wordprocessing/          # WordprocessingML - Word documents
│   ├── document.go          # Document type (Create/Open/Save)
│   ├── elements/            # Word document elements
│   │   ├── paragraph.go     # Para, Run, Text
│   │   ├── table.go         # Table, Row, Cell
│   │   ├── properties.go    # ParaProps, RunProps
│   │   └── ...              # 685+ element types (generated)
│   └── parts/               # Word document parts
│       ├── main.go          # MainDocumentPart
│       ├── styles.go        # StylesPart
│       ├── numbering.go     # NumberingPart
│       ├── header.go        # HeaderPart
│       ├── footer.go        # FooterPart
│       └── ...              # 30+ part types
│
├── internal/
│   ├── opc/                 # Low-level OPC/ZIP implementation
│   └── xml/                 # XML utilities
│
└── testdata/
    └── golden/              # Expected XML for golden file tests
```

**Rationale:** Top-level packages provide cleaner imports (`import "goffice/wordprocessing"`). DrawingML is shared across all three document types (Word, Presentation, Spreadsheet).

### D2: Element Implementation Strategy

**Decision:** Use interface + struct composition instead of class inheritance.

```go
// Core interfaces
type Element interface {
    LocalName() string
    NamespaceURI() string
    Parent() Element
    SetParent(Element)
    OuterXML() string
    InnerXML() string
    Clone() Element
    WriteTo(w io.Writer) error
}

type CompositeElement interface {
    Element
    Children() []Element
    AppendChild(Element) error
    RemoveChild(Element) error
    FirstChild() Element
    LastChild() Element
}

type LeafElement interface {
    Element
    InnerText() string
    SetInnerText(string)
}

// Base implementation via embedding
type elementBase struct {
    localName    string
    namespaceURI string
    parent       Element
    attributes   []Attribute
    features     *FeatureCollection
}

type compositeBase struct {
    elementBase
    children []Element
}

// Concrete type embeds base
type Para struct {
    compositeBase
    Props *ParaProps
}
```

**Rationale:** Go doesn't have inheritance; composition with interfaces is idiomatic and provides flexibility.

### D3: Attribute Type System

**Decision:** Implement typed value wrappers matching C# SDK's OpenXmlSimpleType hierarchy.

```go
// Base interface for all simple types
type SimpleValue interface {
    HasValue() bool
    InnerText() string
    SetInnerText(string) error
}

// Typed implementations
type StringValue struct {
    value *string
}

type Int32Value struct {
    value *int32
}

type BooleanValue struct {
    value *bool
}

type EnumValue[T ~string] struct {
    value *T
}

// Usage in elements
type RunProps struct {
    Bold      *BooleanValue   `xml:"b,attr"`
    Italic    *BooleanValue   `xml:"i,attr"`
    FontSize  *HalfPointValue `xml:"sz,attr"`
    Color     *HexColorValue  `xml:"color,attr"`
}
```

**Rationale:** Type safety prevents invalid attribute values; matches C# SDK pattern.

### D4: Lazy Loading Strategy

**Decision:** Parts load XML on first access; raw bytes preserved until needed.

```go
type OpenXmlPart struct {
    uri         string
    contentType string
    stream      io.ReadSeeker  // Underlying stream
    rootElement Element        // Lazily loaded
    loaded      bool
    mu          sync.RWMutex
}

func (p *OpenXmlPart) RootElement() (Element, error) {
    p.mu.RLock()
    if p.loaded {
        defer p.mu.RUnlock()
        return p.rootElement, nil
    }
    p.mu.RUnlock()

    p.mu.Lock()
    defer p.mu.Unlock()
    if p.loaded {
        return p.rootElement, nil
    }

    // Parse XML into element tree
    root, err := parseElement(p.stream)
    if err != nil {
        return nil, err
    }
    p.rootElement = root
    p.loaded = true
    return root, nil
}
```

**Rationale:** Performance optimization for large documents; only parse what's accessed.

### D5: Relationship Management

**Decision:** Separate internal (part-to-part) and external (hyperlink) relationships.

```go
type RelationshipType string

const (
    RelTypeMainDocument  RelationshipType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
    RelTypeStyles        RelationshipType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"
    RelTypeHyperlink     RelationshipType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"
    // ... etc
)

type Relationship struct {
    ID         string
    Type       RelationshipType
    TargetURI  string
    TargetMode TargetMode // Internal or External
}

type PartRelationship struct {
    Relationship
    TargetPart Part
}

type ExternalRelationship struct {
    Relationship
    TargetURL *url.URL
}
```

**Rationale:** Clear separation matches OOXML spec and C# SDK design.

### D6: Feature Collection Pattern

**Decision:** Implement feature collections for configuration and extensibility.

```go
type FeatureCollection struct {
    parent   *FeatureCollection
    features map[reflect.Type]interface{}
    mu       sync.RWMutex
}

func (fc *FeatureCollection) Get(featureType interface{}) interface{} {
    fc.mu.RLock()
    defer fc.mu.RUnlock()

    t := reflect.TypeOf(featureType).Elem()
    if f, ok := fc.features[t]; ok {
        return f
    }
    if fc.parent != nil {
        return fc.parent.Get(featureType)
    }
    return nil
}

func (fc *FeatureCollection) Set(feature interface{}) {
    fc.mu.Lock()
    defer fc.mu.Unlock()

    t := reflect.TypeOf(feature)
    fc.features[t] = feature
}

// Feature interfaces
type PackageFeature interface {
    Package() *Package
    Capabilities() PackageCapabilities
}

type NamespaceFeature interface {
    ResolveNamespace(prefix string) string
    ResolvePrefix(namespace string) string
}
```

**Rationale:** Matches C# SDK's IFeatureCollection; enables configuration injection without global state.

### D7: Validation Architecture

**Decision:** Two-tier validation: schema (structural) and semantic (business rules).

```go
type Validator interface {
    Validate(element Element, context *ValidationContext) []ValidationError
}

type SchemaValidator struct {
    particle CompiledParticle
}

type SemanticValidator struct {
    constraints []Constraint
}

type Constraint interface {
    Check(element Element) *ValidationError
}

// Built-in constraints
type AttributeValueRangeConstraint struct {
    Attribute string
    Min, Max  int64
}

type RequiredAttributeConstraint struct {
    Attribute string
}

type ParentTypeConstraint struct {
    AllowedParents []reflect.Type
}

// Validation levels (Office 2016+ only)
type FileFormatVersion int

const (
    Office2016 FileFormatVersion = iota
    Office2019
    Office2021
    Microsoft365
)

type ValidationContext struct {
    Version     FileFormatVersion
    StrictMode  bool
    Errors      []ValidationError
}
```

**Rationale:** Separates concerns; matches C# SDK's validation architecture.

### D8: XML Serialization

**Decision:** Custom XML marshaling for precise control over output format.

```go
type XMLWriter struct {
    w           io.Writer
    encoder     *xml.Encoder
    namespaces  map[string]string  // prefix -> uri
    defaultNS   string
}

func (e Element) WriteTo(w *XMLWriter) error {
    // Start element with proper namespace handling
    w.WriteStartElement(e.LocalName(), e.NamespaceURI())

    // Write attributes
    for _, attr := range e.Attributes() {
        w.WriteAttribute(attr)
    }

    // Write children or text
    if composite, ok := e.(CompositeElement); ok {
        for _, child := range composite.Children() {
            child.WriteTo(w)
        }
    } else if leaf, ok := e.(LeafElement); ok {
        w.WriteText(leaf.InnerText())
    }

    w.WriteEndElement()
    return nil
}
```

**Rationale:** Standard library xml.Marshal doesn't handle OOXML's namespace requirements properly.

### D9: Error Handling

**Decision:** Use Go's error patterns with typed errors for specific conditions.

```go
// Sentinel errors
var (
    ErrPackageNotOpen     = errors.New("package is not open")
    ErrReadOnlyPackage    = errors.New("package is read-only")
    ErrPartNotFound       = errors.New("part not found")
    ErrInvalidContentType = errors.New("invalid content type")
)

// Typed errors for detailed information
type ValidationError struct {
    Element     Element
    Path        string
    Description string
    Severity    ValidationSeverity
}

type PartError struct {
    URI string
    Op  string
    Err error
}

func (e *PartError) Error() string {
    return fmt.Sprintf("part %s: %s: %v", e.URI, e.Op, e.Err)
}

func (e *PartError) Unwrap() error {
    return e.Err
}
```

**Rationale:** Standard Go error handling; typed errors enable programmatic handling.

### D10: Concurrency Model

**Decision:** Thread-safe reads; exclusive writes; document-level locking.

```go
type Document struct {
    pkg      *Package
    mu       sync.RWMutex
    modified bool
}

func (d *Document) MainPart() (*MainPart, error) {
    d.mu.RLock()
    defer d.mu.RUnlock()
    // Safe to read concurrently
    return d.mainPart, nil
}

func (d *Document) Save() error {
    d.mu.Lock()
    defer d.mu.Unlock()
    // Exclusive access for write
    return d.pkg.Save()
}
```

**Rationale:** Common pattern; documents typically modified by single goroutine.

## Key Technical Findings: WordprocessingML Architecture

### Document Element Hierarchy
Word documents follow a strict element hierarchy:
```
Document (w:document)
└── Body (w:body)
    ├── Paragraph (w:p)
    │   ├── ParagraphProperties (w:pPr)
    │   └── Run (w:r)
    │       ├── RunProperties (w:rPr)
    │       └── Text (w:t)
    ├── Table (w:tbl)
    │   ├── TableProperties (w:tblPr)
    │   ├── TableGrid (w:tblGrid)
    │   └── TableRow (w:tr)
    │       └── TableCell (w:tc)
    │           └── Paragraph (recursive)
    └── SectionProperties (w:sectPr) - affects preceding content
```

### Part Types (30+)
WordprocessingML documents contain multiple part types linked via relationships:

| Part Type | Content Type | Purpose |
|-----------|--------------|---------|
| MainDocumentPart | application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml | Primary document content |
| StylesPart | .../styles+xml | Style definitions |
| NumberingPart | .../numbering+xml | List/numbering definitions |
| HeaderPart | .../header+xml | Header content (default/first/even) |
| FooterPart | .../footer+xml | Footer content (default/first/even) |
| SettingsPart | .../settings+xml | Document settings |
| FontTablePart | .../fontTable+xml | Font definitions |
| ThemePart | .../theme+xml | Theme definitions |
| CommentsPart | .../comments+xml | Document comments |
| FootnotesPart | .../footnotes+xml | Footnotes |
| EndnotesPart | .../endnotes+xml | Endnotes |
| ImagePart | image/png, image/jpeg, etc. | Embedded images |
| CustomXmlPart | application/xml | Custom XML data |

### Properties Complexity
- **40+ paragraph properties**: justification, indentation, spacing, borders, shading, numbering, keep-with-next, widow/orphan control, tabs, outline level, etc.
- **50+ run properties**: bold, italic, underline (15+ styles), strikethrough, font family (4 slots: ASCII, high ANSI, complex script, East Asian), font size, color, highlight, caps, small caps, subscript/superscript, spacing, kerning, language, etc.

### Style Inheritance Chain
Formatting resolves through a three-level inheritance chain:
```
Document Defaults (w:docDefaults)
    ↓ (base layer)
Style Definitions (w:styles → w:style)
    ↓ (style layer)
Direct Formatting (w:pPr, w:rPr on elements)
    ↓ (override layer)
Final Rendered Formatting
```

### Track Changes Architecture
Revision tracking uses wrapper elements with author/date metadata:
- **InsertedRun** (`w:ins`) - Wraps inserted content
- **DeletedRun** (`w:del`) - Wraps deleted content with `w:delText` for deleted text
- **MoveFrom** / **MoveTo** (`w:moveFrom`, `w:moveTo`) - Content relocation markers
- **rsid attributes** - Revision Save IDs track editing sessions (rsidR, rsidRDefault, rsidDel, rsidP, rsidRPr)

### Section Properties Placement
Section properties (`w:sectPr`) appear at the end of the body but affect **preceding** content. The last section's properties are in `w:body/w:sectPr`, while intermediate section breaks use `w:p/w:pPr/w:sectPr`.

### Headers/Footers Architecture
Headers and footers are stored as separate parts, linked via relationship IDs:
- Each section can reference different headers/footers
- Three types per section: `default`, `first` (title page), `even` (even pages)
- Referenced in section properties: `<w:headerReference w:type="default" r:id="rId4"/>`

### Validation Behavior
**Decision: Strict by default**
- Invalid content is rejected with descriptive errors
- Malformed documents return errors on open
- All 19+ semantic constraint types are enforced
- No silent data loss or auto-repair (explicit user action required)

## Risks / Trade-offs

| Risk | Impact | Mitigation |
|------|--------|------------|
| OOXML complexity | High - spec is 6000+ pages | Phased implementation; focus on common use cases first |
| Performance with large docs | Medium - memory pressure | Lazy loading; streaming API for bulk operations |
| Office version compatibility | Medium - subtle differences | Comprehensive test suite with real documents |
| Schema changes | Low - spec is stable | Version-aware validation; feature flags |
| Go generics limitations | Low - generics are mature | Use interface{} where generics insufficient |

## Migration Plan

N/A - New capability with no existing implementation.

## Resolved Design Questions

### Q1: Code Generation Strategy ✓ DECIDED (UPDATED)
**Decision:** Pre-generate from JSON schemas, commit to repository (unified with Presentation SDK).

The C# SDK generates 67,000+ lines (685 classes) for WordprocessingML alone. Using the same JSON schema approach:
- Reuses the C# SDK's JSON schema metadata in `/data/schemas/`
- **Pre-generated Go code committed to repo** - users do not need to run the generator
- Code generation is a development-time tool, not a runtime dependency
- Consistent with Presentation SDK approach
- Maintains Go idioms in the generator output
- Elements are pre-generated but can have hand-written extensions

**Impact:** Code generator reads JSON schemas and produces ~685 element types as Go structs. These are committed to the repository so consumers can use the SDK without needing to run any code generation.

### Q2: API Naming Style ✓ DECIDED
**Decision:** Idiomatic Go naming conventions.

| C# SDK Name | Go Name |
|-------------|---------|
| `WordprocessingDocument` | `Document` |
| `MainDocumentPart` | `MainPart` |
| `Paragraph` | `Para` |
| `ParagraphProperties` | `ParaProps` |
| `RunProperties` | `RunProps` |
| `TableProperties` | `TableProps` |
| `TableCellProperties` | `CellProps` |
| `SectionProperties` | `SectProps` |

**Rationale:** Go convention favors short, clear names. Package context provides disambiguation (e.g., `word.Document` vs `openxml.Document`).

### Q3: Validation Depth ✓ DECIDED
**Decision:** Full validation (schema + all 19+ semantic constraints).

All semantic constraint types will be implemented:
1. AttributeValueRange - min/max bounds
2. AttributeValuePattern - regex validation
3. AttributeValueSet - enumeration values
4. ParentType - allowed parent elements
5. ChildElement - required/allowed children
6. UniqueValue - uniqueness within scope
7. RelationshipExist - relationship must exist
8. RelationshipType - relationship type validation
9. ReferenceExist - ID reference validation
10. IndexRange - valid index bounds
11. RootAttribute - required root attributes
12. AttributeAbsent - mutually exclusive attributes
13. AttributeCannotOmit - conditionally required
14. AttributeValueLength - string length limits
15. UniqueAttributeValue - unique across document
16. PartContainer - part must be in container
17. DataPart - data part validation
18. PartType - part content type validation
19. Custom constraints - extensible constraint system

### Q4: Office Version Support ✓ DECIDED
**Decision:** Office 2016+ only (ECMA-376 5th edition and later).

This simplifies implementation by not supporting legacy Office 2007/2010/2013 quirks. The validation system will still have version infrastructure for potential future expansion, but initial implementation targets modern documents.

### Q5: Element Metadata Registration ✓ DECIDED
**Decision:** Embedded struct fields.

Each element carries its own metadata in embedded struct fields rather than using a global registry. This provides:
- Self-contained elements with no global state
- No registry lookup overhead at runtime
- Easier testing in isolation
- Higher memory per element, but cleaner architecture

```go
type Para struct {
    compositeBase
    // Embedded metadata
    localName    string // "p"
    namespaceURI string // WordprocessingML namespace
    allowedChildren []reflect.Type
    particle    CompositeParticle
}
```

### Q6: Namespace Handling ✓ DECIDED
**Decision:** Per-element declarations.

Each element is fully self-contained with its own namespace information. This means:
- Elements know their own namespace without external lookup
- No shared namespace registry state
- Elements can be moved between documents without context issues
- More verbose but completely self-describing

### Q7: Typed Child Element Access ✓ DECIDED
**Decision:** Generic methods only.

Use only generic methods for child access: `First[T]()`, `All[T]()`, `OfType[T]()`. No type-specific convenience methods.
- Consistent, minimal API surface across all elements
- Single pattern to learn
- Slightly more verbose but highly predictable

```go
// Usage
props := para.First[ParaProps]()
runs := para.All[Run]()
bookmarks := para.OfType[BookmarkStart]()
```

### Q8: Office Math Support ✓ DECIDED
**Decision:** Full v1 implementation (all 152 OMML types).

Implement complete Office Math (OMML) support in v1 for full feature parity with C# SDK. The `m:` namespace includes:
- Mathematical expressions (OfficeMath)
- Fractions, radicals, matrices
- Limits, integrals, summations
- Mathematical functions
- All 152 element types

### Q9: Optional/Nullable Values ✓ DECIDED
**Decision:** *T pointers (standard Go pattern).

Use pointer types for optional attributes and child elements:
- `nil` = not set / omitted from XML
- Non-nil pointer = value is set
- Familiar to Go developers
- Works well with `omitempty` XML tags

```go
type ParaProps struct {
    Jc      *JustificationValue  // nil = inherit from style
    Spacing *Spacing             // nil = default spacing
    Ind     *Indentation         // nil = no indentation override
}
```

### Q10: Drawing Package Organization ✓ DECIDED
**Decision:** Single `word/drawing` subpackage.

All drawing-related types consolidated in one package despite spanning multiple XML namespaces (wp:, a:, pic:):
- Simpler import: `import "goffice/word/drawing"`
- Practical grouping by functionality
- Package-level comments clarify namespace origins

### Q11: Element Construction Pattern ✓ DECIDED
**Decision:** Functional options pattern.

Use functional options for constructing complex nested structures:

```go
// Functional options for flexible construction
para := word.NewPara(
    word.WithJustification(wml.JustificationCenter),
    word.WithSpacing(word.LineSpacing(240)),
    word.WithRun(
        word.WithText("Hello, World!"),
        word.WithBold(true),
    ),
)
```

Benefits:
- Self-documenting API
- Optional parameters without overloads
- Idiomatic Go pattern
- Easy to extend without breaking changes

### Q12: Implementation Priority Order ✓ DECIDED
**Decision:** Core → Formatting → Tables → Advanced

Phase order:
1. **Core:** Para, Run, Text, Body, Document structure
2. **Formatting:** Styles, Fonts, Run/Paragraph properties
3. **Tables:** Table, Row, Cell, properties, merging
4. **Advanced:** Headers/Footers, Numbering, Track changes, SDT, Math

## Deferred to v2

1. **Memory-Mapped Files:** Large document mmap optimization
2. **Streaming API:** SAX-style streaming for bulk operations
3. **SpreadsheetML:** Excel document support
4. **PresentationML:** PowerPoint document support

## Other Decisions

- **Context Support:** Yes for I/O operations; enables cancellation
- **Enum Implementation:** Type-safe string aliases with validation methods in `wml` package

## Key Findings from SDK Analysis

### Simple Types Implementation Pattern
The C# SDK uses a dual-storage pattern for simple types that enables lazy parsing:
```go
// Go equivalent pattern
type StringValue struct {
    innerValue *string  // Parsed typed value
    textValue  string   // Raw XML text
}

func (v *StringValue) Value() string {
    if v.innerValue != nil {
        return *v.innerValue
    }
    return v.textValue
}
```

**EnumValue<T> Pattern:** The C# SDK requires enum types to implement both `IEnumValue` (provides IsValid, Value properties) AND `IEnumValueFactory<T>` (creates new instances from strings). In Go:
```go
type JustificationValue string

const (
    JustificationLeft    JustificationValue = "left"
    JustificationCenter  JustificationValue = "center"
    JustificationRight   JustificationValue = "right"
    JustificationBoth    JustificationValue = "both"
)

func (v JustificationValue) IsValid() bool {
    switch v {
    case JustificationLeft, JustificationCenter, JustificationRight, JustificationBoth:
        return true
    }
    return false
}
```

### Simple Types Hierarchy (29 types)
```
OpenXmlSimpleType
├── OpenXmlComparableSimpleReference<T> (string-based)
│   └── StringValue, HexBinaryValue, Base64BinaryValue
└── OpenXmlSimpleValue<T> (struct-based)
    ├── OpenXmlComparableSimpleValue<T>
    │   ├── BooleanValue, OnOffValue, TrueFalseValue
    │   ├── Int32Value, Int64Value, UInt32Value, etc.
    │   ├── DoubleValue, DecimalValue, SingleValue
    │   └── DateTimeValue
    ├── EnumValue<T> (with IEnumValue, IEnumValueFactory<T>)
    └── ListValue<T> (space-separated collections)
```

### Element Metadata ConfigureMetadata() Pattern
Elements use a fluent builder pattern to declare their structure:
```csharp
// C# SDK pattern
internal override void ConfigureMetadata(ElementMetadata.Builder builder) {
    builder.SetSchema("w:p");
    builder.AddChild<ParagraphProperties>();
    builder.AddChild<Run>();
    builder.AddChild<BookmarkStart>();
    // ... more children
    builder.Particle = new CompositeParticle.Builder(ParticleType.Sequence, 0, 1) {
        new ElementParticle(typeof(ParagraphProperties), 0, 1),
        new CompositeParticle.Builder(ParticleType.Choice, 0, 0) {
            new ElementParticle(typeof(Run), 0, 1),
            new ElementParticle(typeof(BookmarkStart), 0, 1),
            // ...
        }
    };
}
```

**Go equivalent:**
```go
func (p *Para) ConfigureMetadata() *ElementMetadata {
    return NewElementMetadata().
        SetSchema("w", "p").
        AddChild(reflect.TypeOf((*ParaProps)(nil))).
        AddChild(reflect.TypeOf((*Run)(nil))).
        SetParticle(Sequence(0, 1,
            Element(ParaPropsType, 0, 1),
            Choice(0, Unbounded,
                Element(RunType, 0, 1),
                Element(BookmarkStartType, 0, 1),
            ),
        ))
}
```

### Part Relationship rId Generation
Parts use `rId{N}` format for relationship IDs:
```go
// Pattern for generating unique relationship IDs
func (p *PartContainer) NextRelationshipId() string {
    max := 0
    for _, rel := range p.relationships {
        if strings.HasPrefix(rel.ID, "rId") {
            n, _ := strconv.Atoi(rel.ID[3:])
            if n > max {
                max = n
            }
        }
    }
    return fmt.Sprintf("rId%d", max+1)
}
```

Relationship files (.rels) are stored alongside parts:
- `/_rels/.rels` - Package-level relationships
- `/word/_rels/document.xml.rels` - Main document relationships

### Multi-Namespace Complexity
WordprocessingML documents use 6+ XML namespaces that must be handled:

| Prefix | URI | Purpose |
|--------|-----|---------|
| `w` | `http://schemas.openxmlformats.org/.../wordprocessingml/2006/main` | WordprocessingML |
| `wp` | `http://schemas.openxmlformats.org/.../wordprocessingDrawing/2006/main` | Word drawings |
| `a` | `http://schemas.openxmlformats.org/drawingml/2006/main` | DrawingML |
| `pic` | `http://schemas.openxmlformats.org/.../picture/2006/main` | Pictures |
| `r` | `http://schemas.openxmlformats.org/.../relationships` | Relationships |
| `m` | `http://schemas.openxmlformats.org/.../math/2006/main` | Office Math |
| `mc` | `http://schemas.openxmlformats.org/markup-compatibility/2006` | Compatibility |

**Challenge:** A single Run element containing an inline image requires elements from 5 namespaces:
```xml
<w:r>
  <w:drawing>
    <wp:inline>
      <a:graphic>
        <a:graphicData>
          <pic:pic>
            <pic:blipFill>
              <a:blip r:embed="rId4"/>
```

### Drawing/Image Element Complexity
Images in Word documents have complex nesting across multiple namespaces:
- **Inline images:** `<wp:inline>` - positioned in text flow
- **Anchored images:** `<wp:anchor>` - floating with positioning

The embedding pattern uses relationship IDs:
```go
type Blip struct {
    Embed string `xml:"r:embed,attr"` // rId reference to image part
    Link  string `xml:"r:link,attr"`  // External link (optional)
}
```

### Track Changes / Revision System
Tracked changes use wrapper elements and rsid attributes:
```xml
<w:p w:rsidR="00A1B2C3" w:rsidRDefault="00D4E5F6">
  <w:ins w:id="0" w:author="User" w:date="2024-01-15T10:30:00Z">
    <w:r w:rsidR="00D4E5F6">
      <w:t>inserted text</w:t>
    </w:r>
  </w:ins>
  <w:del w:id="1" w:author="User" w:date="2024-01-15T10:31:00Z">
    <w:r w:rsidDel="00D4E5F6">
      <w:delText>deleted text</w:delText>
    </w:r>
  </w:del>
</w:p>
```

**rsid attributes:** Revision Save IDs track which editing session made changes:
- `rsidR` - Run revision ID
- `rsidRDefault` - Default run revision ID
- `rsidDel` - Deletion revision ID
- `rsidP` - Paragraph revision ID
- `rsidRPr` - Run properties revision ID

### Content Controls (SDT) Pattern
Structured Document Tags provide form fields and data binding:
```go
type SdtBlock struct {
    Props   *SdtProps   // Properties and bindings
    Content *SdtContent // Actual content (paragraphs, etc.)
}

type SdtProps struct {
    Alias       *StringValue    // Display name
    Tag         *StringValue    // Programmatic identifier
    DataBinding *DataBinding    // XPath binding to custom XML
    // Content type declarations
    Text        *SdtContentText
    DropDown    *SdtContentDropDownList
    Date        *SdtContentDate
    ComboBox    *SdtContentComboBox
    Checkbox    *SdtContentCheckbox
}

type DataBinding struct {
    XPath       string // XPath to data
    PrefixMaps  string // Namespace prefix mappings
    StoreItemID string // Custom XML part GUID
}
```

SDT types: `SdtBlock`, `SdtRun`, `SdtCell`, `SdtRow` (context-specific)

### Office Math (OMML) Complexity
The math namespace (`m:`) contains 152 element types:
```go
// Mathematical expression structure
type OfficeMath struct {
    // Content can be any math expression
    Children []MathElement // Fraction, Radical, Matrix, etc.
}

type Fraction struct {
    Type      FractionType   // Bar, NoBar, Skewed, Linear
    Numerator *Numerator
    Denominator *Denominator
}

type Matrix struct {
    Rows []MatrixRow
    ColumnProperties *MatrixColumnProperties
}
```

**Complexity:** Math expressions can deeply nest (e.g., fraction within matrix within radical).

### Headers/Footers/Sections Pattern
Each header/footer type is stored in a separate part:
```go
type HeaderFooterType string

const (
    HeaderFooterDefault HeaderFooterType = "default"
    HeaderFooterFirst   HeaderFooterType = "first"
    HeaderFooterEven    HeaderFooterType = "even"
)

// Section properties reference header/footer parts
type SectProps struct {
    HeaderRefs []HeaderRef // Multiple headers per section
    FooterRefs []FooterRef // Multiple footers per section
}

type HeaderRef struct {
    Type HeaderFooterType
    ID   string // rId to header part
}
```

Documents can have multiple sections with different headers/footers.

### Generated Code Statistics
- WordprocessingML: 685 classes, 67,119 lines
- Uses Roslyn incremental source generator
- Partial classes allow hand-written extensions
- ConfigureMetadata() fluent builder pattern

### Validation System Architecture
- Schema validation: ParticleValidator hierarchy (Sequence, Choice, All)
- Semantic validation: 19+ constraint types (AttributeValueRange, ParentType, UniqueValue, RelationshipExist, etc.)
- ValidationContext with StateManager for caching
- FileFormatVersions flags for version-aware validation

**19 Semantic Constraint Types:**
1. `AttributeValueRangeConstraint` - min/max bounds
2. `AttributeValuePatternConstraint` - regex validation
3. `AttributeValueSetConstraint` - enumeration values
4. `ParentTypeConstraint` - allowed parent elements
5. `ChildElementConstraint` - required/allowed children
6. `UniqueAttributeValueConstraint` - uniqueness in scope
7. `RelationshipExistConstraint` - relationship must exist
8. `RelationshipTypeConstraint` - relationship type check
9. `ReferenceExistConstraint` - ID reference validation
10. `IndexRangeConstraint` - valid index bounds
11. `RootAttributeConstraint` - required root attributes
12. `AttributeAbsentConstraint` - mutually exclusive attrs
13. `AttributeCannotOmitConstraint` - conditionally required
14. `AttributeValueLengthConstraint` - string length limits
15. `PartContainerConstraint` - part must be in container
16. `DataPartConstraint` - data part validation
17. `PartTypeConstraint` - part content type check
18. `UniqueParticleConstraint` - unique child elements
19. `CompatibilityRuleConstraint` - markup compatibility

### Part Hierarchy
- 60+ part types for WordprocessingML
- MainDocumentPart has 25+ possible child parts
- Multiple-instance parts: Header, Footer, Image, CustomXml
- IFixedContentTypePart vs variable content types

## Appendix: C# to Go Type Mapping

| C# Type | Go Type |
|---------|---------|
| `class` with inheritance | `struct` with embedded interface |
| `abstract class` | `interface` + base `struct` |
| `partial class` | N/A - use composition |
| `IEnumerable<T>` | `[]T` or `iter.Seq[T]` |
| `event` | Callback functions |
| `async Task<T>` | Goroutine + channel or context |
| `IDisposable` | `io.Closer` |
| `Nullable<T>` | `*T` (pointer) |
| `Dictionary<K,V>` | `map[K]V` |
| `List<T>` | `[]T` |

## Appendix: Idiomatic Go Naming Conventions

### Package Names
| C# Namespace | Go Package |
|--------------|------------|
| `System.IO.Packaging` | `opc` |
| `DocumentFormat.OpenXml.Framework` | `openxml` |
| `DocumentFormat.OpenXml.Wordprocessing` | `word` |
| `DocumentFormat.OpenXml.Validation` | `openxml/validation` |

### Type Names
| C# Type | Go Type | Rationale |
|---------|---------|-----------|
| `WordprocessingDocument` | `Document` | Package context (word.Document) |
| `MainDocumentPart` | `MainPart` | Shorter, clear |
| `StyleDefinitionsPart` | `StylesPart` | Consistent with C# alternative |
| `NumberingDefinitionsPart` | `NumberingPart` | Consistent pattern |
| `Paragraph` | `Para` | Common abbreviation |
| `ParagraphProperties` | `ParaProps` | Go-style abbreviation |
| `RunProperties` | `RunProps` | Consistent pattern |
| `TableProperties` | `TableProps` | Consistent pattern |
| `TableCellProperties` | `CellProps` | Shorter, context clear |
| `SectionProperties` | `SectProps` | Go-style |
| `DocumentSettings` | `Settings` | Package context (word.Settings) |

### Method Names
| C# Method | Go Method |
|-----------|-----------|
| `AppendChild<T>()` | `Append()` |
| `PrependChild<T>()` | `Prepend()` |
| `GetFirstChild<T>()` | `First()` or `FirstOfType[T]()` |
| `GetPartsOfType<T>()` | `PartsOfType[T]()` |
| `AddMainDocumentPart()` | `AddMainPart()` |

### Constructor Patterns
| C# Pattern | Go Pattern |
|------------|------------|
| `new Paragraph(...)` | `word.NewPara(...)` |
| `WordprocessingDocument.Create(...)` | `word.New(...)` or `word.Create(...)` |
| `WordprocessingDocument.Open(...)` | `word.Open(...)` |
