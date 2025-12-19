# Change: Add Go SDK for Word Document Processing

## Why

The Office Open XML (OOXML) format is the standard for Microsoft Word documents (.docx, .dotx, .docm, .dotm), yet Go lacks a comprehensive, production-ready SDK for creating, reading, and manipulating these documents. This proposal creates a full-featured Go SDK modeled after Microsoft's Open-XML-SDK for C#, enabling Go developers to work with Word documents programmatically with the same capabilities available in .NET.

**Problem Statement:**
- Existing Go libraries for OOXML are incomplete, unmaintained, or lack proper validation
- No Go library provides the full feature set of the C# Open-XML-SDK
- Developers need strongly-typed, idiomatic Go APIs for Word document manipulation
- Enterprise applications require reliable document generation without Office dependencies

**Opportunity:**
- Provide feature parity with Microsoft's Open-XML-SDK in idiomatic Go
- Enable server-side document generation in Go microservices
- Support all Word document types: docx, dotx, docm, dotm
- Deliver proper validation against ECMA-376 and ISO/IEC 29500 standards

## What Changes

### Core Framework (New)
- **ADDED** `framework`: DOM-style element tree with XML serialization
- **ADDED** `types`: Typed attribute values (StringValue, Int32Value, EnumValue, etc.)
- **ADDED** `package`: OPC (Open Packaging Conventions) implementation over ZIP, abstract part system with content types and URIs
- **ADDED** `relationships`: Part relationships and external references
- **ADDED** `validation`: Schema and semantic validation framework
- **ADDED** `features`: Feature collection system for extensibility

### WordProcessing Domain (New)
- **ADDED** `wordprocessing-document`: WordprocessingDocument type with Create/Open/Save lifecycle
- **ADDED** `wordprocessing-parts`: All Word-specific parts (MainDocumentPart, StylesPart, etc.)
- **ADDED** `wordprocessing-elements`: Document structure elements (Document, Body, Paragraph, Run, Text)
- **ADDED** `wordprocessing-properties`: Formatting properties (ParagraphProperties, RunProperties)
- **ADDED** `wordprocessing-tables`: Table elements (Table, TableRow, TableCell, TableProperties)
- **ADDED** `wordprocessing-styles`: Style definitions and application
- **ADDED** `wordprocessing-numbering`: Numbering definitions for lists
- **ADDED** `wordprocessing-settings`: Document settings and preferences

### API Design Principles
- Idiomatic Go: Use interfaces, composition, and Go conventions
- Feature parity: Match Open-XML-SDK capabilities
- Performance: Lazy loading, streaming support, minimal allocations
- Type safety: Strongly-typed elements and attributes
- Extensibility: Feature collection pattern for configuration

## Impact

### Affected Specs
- `framework` (NEW)
- `types` (NEW)
- `package` (NEW)
- `relationships` (NEW)
- `validation` (NEW)
- `features` (NEW)
- `wordprocessing-document` (NEW)
- `wordprocessing-parts` (NEW)
- `wordprocessing-elements` (NEW)
- `wordprocessing-properties` (NEW)
- `wordprocessing-tables` (NEW)
- `wordprocessing-styles` (NEW)
- `wordprocessing-numbering` (NEW)
- `wordprocessing-settings` (NEW)

### Affected Code
- New package: `github.com/connerohnesorge/goffice/packaging` - OPC implementation
- New package: `github.com/connerohnesorge/goffice/openxml` - Core framework (element types, features, validation)
- New package: `github.com/connerohnesorge/goffice/openxml/types` - Simple types
- New package: `github.com/connerohnesorge/goffice/openxml/validation` - Validation framework
- New package: `github.com/connerohnesorge/goffice/openxml/features` - Feature collection
- New package: `github.com/connerohnesorge/goffice/wordprocessing` - Word document support
- New package: `github.com/connerohnesorge/goffice/wordprocessing/elements` - WordprocessingML schema types (generated)
- New package: `github.com/connerohnesorge/goffice/wordprocessing/parts` - Word document parts
- New package: `github.com/connerohnesorge/goffice/drawingml` - Shared DrawingML types

### Dependencies
- `archive/zip` - Standard library ZIP support
- `encoding/xml` - Standard library XML support
- No external dependencies for core functionality

### Risks
- **Complexity**: OOXML is a large specification; phased implementation recommended
- **Compatibility**: Must handle documents from various Office versions
- **Performance**: Large documents require careful memory management
- **Validation**: Full schema validation is computationally expensive

### Migration
- N/A - This is a new capability with no existing implementation to migrate

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
| **Element Metadata** | Embedded struct fields | XMLName, Namespace, LocalName as struct fields, self-contained, works with encoding/xml |
| **Construction Pattern** | Functional options | `NewPara(WithText("Hello"), WithBold(true))` - idiomatic Go, extensible, self-documenting |
| **Child Access** | Generic functions | `First[T](el)`, `All[T](el)`, `OfType[T](el)` using Go 1.18+ generics with range-over-func iterators |
| **API Naming** | Go-idiomatic short | `Para`, `ParaProps`, `Run` (not verbose `ParagraphProperties`) |

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
| **DrawingML** | Shared `drawingml/` package | Used by all three SDKs (Word, Presentation, Spreadsheet) |

### XML Processing

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **XML Prefixes** | Fixed canonical | Always use canonical prefixes: w: for WordML, p: for PresentationML, a: for DrawingML |
| **MC Handling** | Parse-time processing | Process AlternateContent/Choice/Fallback during XML parsing based on target FileFormatVersion |
| **Namespace Management** | Embedded in types | Each generated type knows its namespace URI and local name |

### Implementation Scope

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Phase 1 Scope** | Full feature parity with Open-XML-SDK | Complete implementation, not an MVP |
| **Package Structure** | Top-level packages | `drawingml/`, `packaging/`, `openxml/`, `wordprocessing/` |

## Package Structure

```
goffice/
├── drawingml/           # SHARED - shapes, images, effects (used by Word, Presentation, Spreadsheet)
├── packaging/           # OPC layer (Open Packaging Conventions)
├── openxml/             # Core framework (element types, features, validation)
└── wordprocessing/      # WordprocessingML
    ├── elements/        # Word document elements
    └── parts/           # Word document parts
```

## Key Technical Findings

### Document Element Hierarchy
Word documents follow a strict element hierarchy:
```
Document → Body → Paragraph → Run → Text
```

### Part Types
WordprocessingML documents contain 30+ part types:
- **MainDocumentPart** - Primary document content
- **StylesPart** - Style definitions
- **NumberingPart** - Numbering/list definitions
- **HeaderPart** - Header content (multiple per document: default, first, even)
- **FooterPart** - Footer content (multiple per document: default, first, even)
- **SettingsPart** - Document settings
- **FontTablePart** - Font definitions
- **ThemePart** - Theme definitions
- **CommentsPart** - Document comments
- **FootnotesPart** / **EndnotesPart** - Notes
- **ImagePart** - Embedded images
- **CustomXmlPart** - Custom XML data
- Plus 20+ additional specialized parts

### Properties Complexity
- **40+ paragraph properties** (justification, indentation, spacing, borders, numbering, etc.)
- **50+ run properties** (bold, italic, underline, font, color, size, etc.)

### Tables with Recursive Structure
Tables contain nested cells which can contain paragraphs, which can contain tables:
```
Table → TableRow → TableCell → Paragraph → Run → Text
                             → Table (nested)
```

### Style Inheritance Chain
Formatting resolves through inheritance:
```
Document defaults → Styles → Direct formatting
```
Later sources override earlier ones.

### Track Changes Markers
Revision tracking uses wrapper elements:
- **InsertedRun** (`w:ins`) - Inserted content
- **DeletedRun** (`w:del`) - Deleted content
- **MoveFrom** / **MoveTo** - Moved content markers

### Section Properties Placement
Section properties (`w:sectPr`) appear at the end of the body but affect preceding content. This is a key architectural consideration for document processing.

### Headers/Footers Architecture
Headers and footers are stored as separate parts, linked via relationships:
- Each section can have different headers/footers
- Three types per section: default, first page, even pages
- Referenced via relationship IDs in section properties
