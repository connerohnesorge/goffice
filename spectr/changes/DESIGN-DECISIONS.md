# Goffice Unified Design Decisions (ULTRATHINK Approved)

This document consolidates all design decisions approved through ULTRATHINK analysis for the goffice SDK implementation. These decisions apply consistently across all four SDK proposals:

- `add-go-word-sdk` - WordprocessingML (.docx)
- `add-go-spreadsheet-sdk` - SpreadsheetML (.xlsx)
- `add-go-presentation-sdk` - PresentationML (.pptx)
- `add-go-drawingml` - Shared DrawingML graphics layer

## Core Architecture

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Module Path** | `github.com/connerohnesorge/goffice` | Single unified module, simple imports, atomic versioning |
| **Go Version** | 1.25+ | Range-over-func for iterators, modern generics features |
| **Module Layout** | Single unified go.mod | Shared types, atomic versioning, simple dependency graph |
| **Dependencies** | Pure stdlib only | `archive/zip`, `encoding/xml`, `sync` - no external dependencies |
| **Office Version** | 2016+ only (ECMA-376 5th edition+) | Modern documents, reduced complexity, covers 95%+ of real-world files |
| **Thread Safety** | Per-Document RWMutex | Single `sync.RWMutex` per `OpenXmlPackage`, simple and sufficient |

## Element System

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Element Metadata** | Embedded struct fields | `XMLName`, `Namespace`, `LocalName` as struct fields, self-contained, works with `encoding/xml` |
| **Construction Pattern** | Functional options | `NewPara(WithText("Hello"), WithBold(true))` - idiomatic Go, extensible, self-documenting |
| **Child Access** | Generic functions | `First[T](el)`, `All[T](el)`, `OfType[T](el)` using Go 1.18+ generics with range-over-func iterators |
| **API Naming** | Go-idiomatic short names | `Para`, `ParaProps`, `Run`, `Slide`, `Cell` (not verbose C# names like `ParagraphProperties`) |

## Validation System

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Validation Strategy** | Code-generated Validate() methods | Generate type-specific `Validate()` during code gen, zero reflection, compile-time type safety |
| **Error Handling** | Structured errors | Custom error types (`ValidationError`, `ParseError`) with path, element, constraint info, `errors.Is/As` compatible |
| **Validation Mode** | Strict by default | Reject invalid content, return errors for malformed documents |

## Extensibility

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Feature Collection** | Interface Registry pattern | Define `Feature` interface, register by interface type, supports hierarchy (Element→Part→Package→Global) with fallback chain |
| **DrawingML** | Shared `drawingml/` package | Used by all three SDKs (Word, Presentation, Spreadsheet) for shapes, images, effects, and charts |

## XML Processing

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **XML Prefixes** | Fixed canonical | Always use canonical prefixes per namespace (see table below) |
| **MC Handling** | Parse-time processing | Process `AlternateContent/Choice/Fallback` during XML parsing based on target `FileFormatVersion` |
| **Namespace Management** | Embedded in types | Each generated type knows its namespace URI and local name |

### Canonical XML Prefixes

| Prefix | Namespace URI | Domain |
|--------|---------------|--------|
| `w` | `http://schemas.openxmlformats.org/wordprocessingml/2006/main` | WordprocessingML |
| `x` | `http://schemas.openxmlformats.org/spreadsheetml/2006/main` | SpreadsheetML |
| `p` | `http://schemas.openxmlformats.org/presentationml/2006/main` | PresentationML |
| `a` | `http://schemas.openxmlformats.org/drawingml/2006/main` | DrawingML Core |
| `c` | `http://schemas.openxmlformats.org/drawingml/2006/chart` | Charts |
| `wp` | `http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing` | Word Drawing |
| `xdr` | `http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing` | Spreadsheet Drawing |
| `pic` | `http://schemas.openxmlformats.org/drawingml/2006/picture` | Pictures |
| `r` | `http://schemas.openxmlformats.org/officeDocument/2006/relationships` | Relationships |

## Code Generation

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Strategy** | JSON Schema → Go | Parse Open-XML-SDK's JSON schema files, generate Go structs with XML tags |
| **Output** | Pre-generated and committed | Users don't need the generator; ready-to-use types in repository |
| **Scope** | ~1000+ types per domain | Full coverage of all ECMA-376 element types |

## Implementation Scope

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Phase 1** | Full feature parity with Open-XML-SDK | Complete implementation, not an MVP |
| **Testing** | Golden file tests | Compare serialized XML against expected output |

## Package Structure

```
goffice/
├── drawingml/               # SHARED - shapes, images, charts (a: namespace)
│   ├── main/                # Core DrawingML types
│   ├── chart/               # Chart types (c: namespace)
│   ├── spreadsheet/         # Spreadsheet drawing (xdr: namespace)
│   ├── word/                # Word drawing (wp: namespace)
│   └── picture/             # Picture types (pic: namespace)
│
├── packaging/               # OPC layer (shared)
│   ├── package.go           # OpenXmlPackage
│   ├── part.go              # OpenXmlPart base
│   ├── relationship.go      # Relationship management
│   └── content_types.go     # Content types
│
├── openxml/                 # Core framework (shared)
│   ├── element.go           # Element interface and base types
│   ├── composite.go         # CompositeElement for containers
│   ├── leaf.go              # LeafElement for simple content
│   ├── validation.go        # Validation framework
│   ├── types/               # OpenXML value types
│   │   ├── string.go        # StringValue
│   │   ├── int.go           # Int32Value, Int64Value, etc.
│   │   ├── bool.go          # BooleanValue
│   │   └── enum.go          # EnumValue[T]
│   └── features/            # Feature collection pattern
│
├── wordprocessing/          # WordprocessingML (w: namespace)
│   ├── elements/            # Document, Paragraph, Run elements
│   └── parts/               # Document parts
│
├── spreadsheet/             # SpreadsheetML (x: namespace)
│   ├── elements/            # Workbook, Worksheet, Cell elements
│   └── parts/               # Workbook parts
│
├── presentation/            # PresentationML (p: namespace)
│   ├── elements/            # Presentation, Slide elements
│   └── parts/               # Presentation parts
│
└── internal/
    ├── opc/                 # Low-level OPC/ZIP implementation
    └── xml/                 # XML utilities, MC processing
```

## Design Decision Summary

### ULTRATHINK Question 1: Validation System
**Q:** How should constraints be implemented?
**A:** Code-generated Validate() methods per type. Zero reflection, compile-time type safety, matches SDK pattern.

### ULTRATHINK Question 2: Feature Collection Pattern
**Q:** How should Feature Collection translate to Go?
**A:** Interface Registry pattern. Define Feature interface, register by interface type, supports hierarchy (Element→Part→Package→Global) with fallback chain.

### ULTRATHINK Question 3: Markup Compatibility (MC) Handling
**Q:** When should MC blocks be processed?
**A:** Parse-time processing. Process MC during XML parsing based on target FileFormatVersion. Users get clean element tree without MC complexity.

### ULTRATHINK Question 4: Element Construction
**Q:** What construction pattern should be used?
**A:** Functional options. `NewParagraph(WithText("Hello"), WithStyle("Heading1"))`. Idiomatic Go, extensible, self-documenting.

### ULTRATHINK Question 5: Child Access
**Q:** How should child element access work?
**A:** Generic functions. `First[T](el)`, `All[T](el)`, `OfType[T](el)` using Go 1.18+ generics with range-over-func iterators.

### ULTRATHINK Question 6: Code Generation Strategy
**Q:** What code generation strategy should be used?
**A:** JSON Schema → Go. Parse Open-XML-SDK's JSON schema files, generate Go structs with XML tags. Matches C# SDK approach, maintainable.

### ULTRATHINK Question 7: Module Layout
**Q:** Single module or multi-module?
**A:** Single unified go.mod. Shared types, atomic versioning, simple dependency graph.

### ULTRATHINK Question 8: Thread Safety
**Q:** Fine-grained or coarse-grained locking?
**A:** Per-document RWMutex. Single sync.RWMutex per OpenXmlPackage, simple and sufficient for typical use cases.

### ULTRATHINK Question 9: Element Metadata Storage
**Q:** How should element metadata be stored?
**A:** Embedded struct fields. XMLName, Namespace, LocalName as struct fields, self-contained, works with encoding/xml.

### ULTRATHINK Question 10: XML Prefix Handling
**Q:** Preserve original prefixes or use canonical?
**A:** Fixed canonical. Always use canonical prefixes (w:, p:, a:, x:, etc.). Predictable output, simpler code.

### ULTRATHINK Question 11: Error Handling Strategy
**Q:** What error pattern should be used?
**A:** Structured errors. Custom error types (ValidationError, ParseError) with path, element, constraint info. errors.Is/As compatible.

### ULTRATHINK Question 12: Minimum Office Version
**Q:** What Office version should be the minimum target?
**A:** Office 2016+ (ECMA-376 5th edition+). Modern documents only, reduced complexity, covers 95%+ of real-world files.
