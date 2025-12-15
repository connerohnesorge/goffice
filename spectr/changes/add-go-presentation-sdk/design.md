# Design: Go SDK for Office Open XML Presentations

## Context

The Microsoft Open-XML-SDK is a mature, well-designed SDK that provides comprehensive support for Office Open XML documents. This design document outlines how we will translate the C# SDK's architecture and patterns into idiomatic Go while maintaining full feature parity.

### Key Architectural Patterns from Open-XML-SDK (from codebase exploration)

1. **Code Generation Infrastructure**
   - Roslyn-based incremental generator (`IIncrementalGenerator`)
   - JSON schema metadata in `/data/schemas/` (90+ files, 56K+ lines for WordprocessingML alone)
   - Generator models: `SchemaType`, `SchemaAttribute`, `Particle`, `Validator`
   - Particle system: Element, Sequence, Choice, All, Group, Any

2. **Element Metadata System**
   - `ElementMetadata` contains: Type, Attributes, Children, Validators, Constraints, Particle
   - `AttributeMetadata` with property name, QName, value type, validators
   - `CompiledParticle` for efficient O(1) element ordering
   - Factory pattern via `IElementMetadataFactoryFeature`

3. **Validation Architecture**
   - `OpenXmlValidator` → `DocumentValidator` → `SchemaTypeValidator` + `SemanticConstraint`
   - `ParticleValidator` hierarchy: Sequence, Choice, All, Group validators
   - 21 semantic constraint types (AttributeValueSet, AttributeRange, MutualExclusive, etc.)
   - `ValidationContext` with stack, errors, version, MC context

4. **Feature Collection Pattern**
   - `IFeatureCollection` with type-based Get/Set
   - Three-level hierarchy: Element/Part → Package → Global defaults
   - Lazy initialization via `GetKnown()` pattern
   - 40+ feature interfaces (IPackageFeature, IPartRelationshipsFeature, etc.)
   - Event system: `IPackageEventsFeature`, `IPartEventsFeature` with 15 event types

5. **Packaging/Part System**
   - `IPackage`/`IPackagePart` abstractions over System.IO.Packaging
   - `PartRelationshipsFeature` with lazy relationship loading
   - Content type registration via `PartExtensionProvider`
   - Builder pattern: `DelegatingPackageFeature`, `SaveablePackage`

6. **XML Infrastructure**
   - `XmlConvertingReader` for strict↔transitional namespace translation
   - `OpenXmlPartReader` for streaming (900+ lines, state machine)
   - `OpenXmlPartWriter` for streaming writes
   - Lazy parsing with `RawOuterXml` caching

## Goals / Non-Goals

### Goals
- Full feature parity with Open-XML-SDK presentation capabilities
- Idiomatic Go API design (not a direct C# port)
- Thread-safe operations where applicable
- Comprehensive test coverage
- Clear documentation with examples
- Minimal external dependencies
- Code generator for schema types (like C# SDK)

### Non-Goals
- Word processing document support (separate proposal: add-go-word-sdk)
- Spreadsheet document support (future phase)
- GUI/visual editing tools
- Template engines (can be built on top)
- LINQ-style query support (removed from scope)

## Unified Design Decisions (Cross-SDK)

These decisions apply to both Presentation and Word SDKs for consistency:

| Decision | Choice |
|----------|--------|
| **Module Path** | `github.com/connerohnesorge/goffice` |
| **Code Generation** | Generate from JSON schemas (pre-generated, committed) |
| **Office Version** | 2016+ only (ECMA-376 5th edition+) |
| **Package Structure** | `pkg/{framework,types,package,relationships,validation,features,presentation,word}/` |
| **API Naming** | Go-idiomatic short names (`Slide`, `Para`, not verbose C# names) |
| **Element Metadata** | Embedded struct fields (no global registry) |
| **Construction Pattern** | Functional options: `NewSlide(WithTitle("Hello"))` |
| **Child Access** | Generic methods only: `First[T]()`, `All[T]()`, `OfType[T]()` |

## Decisions

### Decision 1: Element Type System

**What**: Use struct embedding and interfaces instead of inheritance.

**C# Pattern**:
```csharp
public class Slide : OpenXmlCompositeElement { }
```

**Go Pattern**:
```go
// Interface for polymorphism
type Element interface {
    LocalName() string
    NamespaceURI() string
    Parent() Element
    // ... navigation methods
}

// Embedded struct for shared behavior
type Slide struct {
    CompositeElement
    // Slide-specific fields generated from schema
}
```

**Why**: Go doesn't have inheritance. Embedding provides composition with method promotion.

---

### Decision 2: Code Generation Strategy

**What**: Create a code generator that reads the same JSON schema files as the C# SDK.

**Input**: JSON schema metadata from `/data/schemas/` and `/data/typed/`
- `SchemaType` definitions with attributes and particles
- `Validator` definitions for attribute constraints
- `Particle` definitions for child element ordering

**Output**:
```go
// Generated element struct
type Slide struct {
    *CompositeElement
    cSld    *CommonSlideData // p:cSld - required
    clrMapOvr *ColorMapOverride // p:clrMapOvr - optional
    timing  *Timing          // p:timing - optional
    // ...
}

// Generated metadata registration
func init() {
    RegisterElement(ElementMetadata{
        QName:     QName{NS: PresentationML, Local: "sld"},
        Type:      reflect.TypeOf((*Slide)(nil)),
        Factory:   func() Element { return &Slide{} },
        Particle:  SequenceParticle{/* ... */},
        // ...
    })
}
```

**Why**:
- ECMA-376 defines 1000+ element types
- JSON metadata already exists and is maintained
- Same approach as C# SDK ensures parity

---

### Decision 3: Particle System for Child Element Ordering

**What**: Port the particle constraint system for validating child element order.

**Go Pattern**:
```go
type ParticleType int
const (
    ParticleElement ParticleType = iota
    ParticleSequence
    ParticleChoice
    ParticleAll
    ParticleGroup
    ParticleAny
)

type Particle interface {
    Type() ParticleType
    MinOccurs() int
    MaxOccurs() int // -1 = unbounded
    Version() FileFormatVersions
}

type CompositeParticle struct {
    particleType ParticleType
    children     []Particle
    minOccurs    int
    maxOccurs    int
}
```

**Why**: Child element ordering is strictly defined in ECMA-376 and must be validated.

---

### Decision 4: Feature Collection Pattern

**What**: Port the feature collection pattern with type-safe generics.

**Go Pattern**:
```go
type FeatureCollection interface {
    Get(key reflect.Type) any
    Set(key reflect.Type, value any)
    IsReadOnly() bool
    Revision() int
}

// Type-safe helpers using generics
func GetFeature[T any](fc FeatureCollection) (T, bool) {
    v := fc.Get(reflect.TypeOf((*T)(nil)).Elem())
    if v == nil {
        var zero T
        return zero, false
    }
    return v.(T), true
}

func GetRequiredFeature[T any](fc FeatureCollection) T {
    v, ok := GetFeature[T](fc)
    if !ok {
        panic(fmt.Sprintf("feature not registered: %T", (*T)(nil)))
    }
    return v
}
```

**Three-level hierarchy**:
1. Part features → inherits from Package
2. Package features → inherits from Global
3. Global defaults (namespace resolver, metadata factory)

**Why**: This pattern enables extensibility without modifying core types.

---

### Decision 5: Validation System

**What**: Port the multi-level validation system.

**Architecture**:
```go
type Validator struct {
    fileFormat   FileFormatVersions
    maxErrors    int
}

func (v *Validator) Validate(pkg *Package) []ValidationError
func (v *Validator) ValidatePart(part Part) []ValidationError
func (v *Validator) ValidateElement(elem Element) []ValidationError

// Validation types
type ValidationErrorType int
const (
    SchemaError ValidationErrorType = iota
    SemanticError
    PackageError
    MarkupCompatibilityError
)

// Schema validation via particles
type ParticleValidator interface {
    Validate(ctx *ValidationContext) error
    TryMatch(elem Element) (bool, error)
    GetExpectedElements() []ElementType
}

// Semantic constraints
type SemanticConstraint interface {
    Level() SemanticValidationLevel
    Version() FileFormatVersions
    Validate(ctx *ValidationContext) *ValidationError
}
```

**Semantic constraint types to implement** (21 from C# SDK):
- AttributeValueSetConstraint
- AttributeValueRangeConstraint
- AttributeMutualExclusive
- UniqueAttributeValueConstraint
- RelationshipExistConstraint
- ReferenceExistConstraint
- etc.

---

### Decision 6: Simple Type System

**What**: Create value wrapper types with XML serialization.

**Go Pattern**:
```go
type StringValue struct {
    value    string
    hasValue bool
}

func (v StringValue) Value() string       { return v.value }
func (v StringValue) HasValue() bool      { return v.hasValue }
func (v *StringValue) SetValue(s string)  { v.value = s; v.hasValue = true }
func (v StringValue) InnerText() string   { return v.value }

// Generic enum value
type EnumValue[T ~int] struct {
    value    T
    hasValue bool
}

// Must implement xml.Marshaler/Unmarshaler for each type
func (v StringValue) MarshalXMLAttr(name xml.Name) (xml.Attr, error)
func (v *StringValue) UnmarshalXMLAttr(attr xml.Attr) error
```

---

### Decision 7: XML Reading/Writing Infrastructure

**What**: Support both streaming and DOM modes like C# SDK.

**Streaming Reader** (like `OpenXmlPartReader`):
```go
type PartReader struct {
    xmlReader   *xml.Decoder
    elementStack []elementState
    state        readerState
}

func (r *PartReader) Read() bool
func (r *PartReader) ReadFirstChild() bool
func (r *PartReader) ReadNextSibling() bool
func (r *PartReader) LoadCurrentElement() (Element, error)
func (r *PartReader) Attributes() []Attribute
```

**DOM Mode** (via `OuterXml`/`InnerXml`):
```go
type Element interface {
    OuterXml() string      // Full XML including element
    InnerXml() string      // Just children XML
    SetOuterXml(string)    // Parse and replace
    WriteTo(io.Writer) error
}
```

**Lazy Parsing**:
- Store `rawOuterXml` until first child access
- Parse on demand via `ensureParsed()`

---

### Decision 8: Namespace Management

**What**: Pre-register all OpenXML namespaces with translation support.

**Go Pattern**:
```go
var (
    NSPresentationML = Namespace{
        URI:           "http://schemas.openxmlformats.org/presentationml/2006/main",
        Prefix:        "p",
        StrictURI:     "http://purl.oclc.org/ooxml/presentationml/main",
    }
    // 100+ namespace registrations
)

type NamespaceResolver interface {
    LookupNamespace(prefix string) string
    LookupPrefix(uri string) string
    TryGetTransitionalNamespace(strict string) (string, bool)
    TryGetExtendedNamespace(obsolete string) (string, bool)
}
```

**Why**: Strict↔Transitional translation is required for ISO 29500 compliance.

---

### Decision 9: Packaging Layer

**What**: Wrap `archive/zip` with OPC abstractions.

**Go Pattern**:
```go
// Low-level OPC package
type opcPackage struct {
    zipReader    *zip.Reader
    zipWriter    *zip.Writer
    parts        map[string]*opcPart
    contentTypes *contentTypes
    relationships map[string][]*relationship
}

// High-level OpenXML package
type Package struct {
    opc      *opcPackage
    features FeatureCollection
    settings OpenSettings
}

// Part abstraction
type Part interface {
    URI() string
    ContentType() string
    Package() *Package
    GetStream() (io.ReadCloser, error)
    RootElement() (Element, error)
    Features() FeatureCollection
}
```

---

### Decision 10: Event System

**What**: Port the event system for lifecycle hooks.

**Go Pattern**:
```go
type EventType int
const (
    EventClosing EventType = iota
    EventClosed
    EventCreating
    EventCreated
    EventSaving
    EventSaved
    // etc.
)

type PackageEventHandler func(pkg *Package, event EventType)

type PackageEventsFeature interface {
    OnChange(handler PackageEventHandler)
    Raise(pkg *Package, event EventType)
}
```

---

## Risks / Trade-offs

### Risk 1: JSON Schema Parsing Complexity
- **Risk**: The JSON schema format is undocumented
- **Mitigation**: Reverse-engineer from existing data, add comprehensive tests

### Risk 2: Performance with Large Documents
- **Risk**: Go's GC may struggle with very large DOM trees
- **Mitigation**: Lazy loading, streaming APIs, optional memory pooling

### Risk 3: XML Namespace Complexity
- **Risk**: Go's encoding/xml has limited namespace control
- **Mitigation**: Custom marshaling, consider etree library for complex cases

### Risk 4: Generated Code Maintenance
- **Risk**: Schema changes require regeneration
- **Mitigation**: Versioned packages, automated CI regeneration

## Resolved Design Decisions

Based on user input, the following decisions have been finalized:

### Module & Compatibility

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Go Module Path** | `github.com/connerohnesorge/goffice` | Personal namespace, easy to start |
| **Minimum Go Version** | Go 1.22+ | Latest features, loop variable semantics, mature generics |
| **External Dependencies** | Pure stdlib only | Maximum compatibility, no supply chain risk |
| **Generated Code** | Pre-generated, committed to repo | No generator tool required by users |

### API Design

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Property Access** | Direct struct fields | Go idiomatic: `slide.ShapeTree.Shapes[0]` |
| **Error Handling** | Return errors everywhere | Every operation returns `(result, error)` |
| **Validation Default** | Strict by default | Reject invalid content during load |
| **Thread Safety** | Package-level sync.Mutex | One lock per document, like C# SDK |

### Implementation Implications

**Direct Struct Fields**:
```go
type Slide struct {
    CommonSlideData    *CommonSlideData    // p:cSld
    ColorMapOverride   *ColorMapOverride   // p:clrMapOvr
    Timing             *Timing             // p:timing
    Transition         *Transition         // p:transition
    // Fields are directly accessible
}

// Usage
title := slide.CommonSlideData.ShapeTree.Shapes[0]
```

**Error Returns Everywhere**:
```go
// All operations return error
doc, err := presentation.Open("file.pptx")
if err != nil {
    return err
}

slide, err := doc.GetSlide(0)
if err != nil {
    return err
}
```

**Strict Validation**:
```go
// During load, invalid content causes error
doc, err := presentation.Open("invalid.pptx")
// err != nil if document violates ECMA-376 schema

// Explicit permissive mode available via settings
doc, err := presentation.Open("file.pptx", WithStrictValidation(false))
```

**Package-Level Mutex**:
```go
type PresentationDocument struct {
    mu       sync.Mutex
    // All methods acquire mu before operation
}

func (d *PresentationDocument) Save() error {
    d.mu.Lock()
    defer d.mu.Unlock()
    // ...
}
```

### Implementation Strategy

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **MVP Scope** | Full read/write from v0.1 | Complete Create/Open/Modify/Save API from first release |
| **Code Generation** | Pre-generated, committed | Parse JSON schemas once, commit Go code. Users need no generator |
| **MC Handling** | Process automatically | Select Choice/Fallback based on FileFormatVersions |
| **Testing Strategy** | Golden file tests | Expected XML output for operations, fast and deterministic |

### XML & Namespace Handling

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **XML Prefixes** | Fixed canonical prefixes | Always `p:` for PresentationML, `a:` for DrawingML |
| **Child Element Access** | Direct struct fields | `slide.ShapeTree`, `shape.TextBody` - Go idiomatic |
| **Package Structure** | Flat by namespace | `pkg/pml/`, `pkg/dml/`, `pkg/rml/` |
| **Version Support** | Office 2016+ | FileFormatVersions.Office2016 minimum, ECMA-376 5th edition |

### Package Organization

```
goffice/
├── pkg/
│   ├── framework/      # Core element types, attributes, base interfaces
│   ├── types/          # OpenXML value types (StringValue, Int32Value, etc.)
│   ├── package/        # OPC package abstraction and document base
│   ├── relationships/  # Part relationships and content types
│   ├── validation/     # Schema and semantic validation
│   ├── features/       # Feature collection pattern implementation
│   ├── presentation/   # PresentationDocument, document types
│   ├── pml/            # PresentationML schema elements (p: namespace)
│   ├── dml/            # DrawingML schema elements (a: namespace)
│   └── rml/            # RelationshipML elements (r: namespace)
├── internal/
│   ├── opc/            # Low-level OPC/ZIP handling
│   └── xml/            # XML utilities, MC processing
└── testdata/
    └── golden/         # Expected XML output for golden file tests
```

### MC Processing Behavior

```go
// Default: Process MC content based on target version
doc, err := presentation.Open("file.pptx")
// Automatically selects appropriate Choice based on Office 2016+

// Explicit version targeting
doc, err := presentation.Open("file.pptx",
    WithTargetVersion(FileFormatVersions.Microsoft365))
```

### Golden File Testing Pattern

```go
func TestCreateSlide(t *testing.T) {
    doc := createTestPresentation()
    doc.AddSlide()

    got := doc.ToXML()
    golden := testdata.Load("create_slide.xml")

    if diff := xmldiff.Compare(got, golden); diff != "" {
        t.Errorf("XML mismatch:\n%s", diff)
    }
}
```
