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
- New package: `github.com/connerohnesorge/goffice/pkg/package` - OPC implementation
- New package: `github.com/connerohnesorge/goffice/pkg/framework` - Core element types
- New package: `github.com/connerohnesorge/goffice/pkg/types` - Simple types
- New package: `github.com/connerohnesorge/goffice/pkg/relationships` - Relationship management
- New package: `github.com/connerohnesorge/goffice/pkg/validation` - Validation framework
- New package: `github.com/connerohnesorge/goffice/pkg/features` - Feature collection
- New package: `github.com/connerohnesorge/goffice/pkg/word` - Word document support
- New package: `github.com/connerohnesorge/goffice/pkg/wml` - WordprocessingML schema types (generated)

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

## Unified Design Decisions

These decisions apply to both Presentation and Word SDKs for consistency:

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Module Path** | `github.com/connerohnesorge/goffice` | Personal namespace, easy to start |
| **Code Generation** | Generate from JSON schemas | Reuse C# SDK schema data, 1000+ types |
| **Office Version** | 2016+ only | Modern documents, reduced complexity (ECMA-376 5th edition+) |
| **Package Structure** | `pkg/` style | `pkg/{framework,types,package,relationships,validation,features,word}` |
| **API Naming** | Go-idiomatic short | `Document`, `Para`, `ParaProps` (not verbose C# names) |
| **Element Metadata** | Embedded struct fields | Self-contained elements, no global registry |
| **Construction Pattern** | Functional options | `NewPara(WithText("Hello"), WithBold(true))` |
| **Child Access** | Generic methods only | `First[T]()`, `All[T]()`, `OfType[T]()` |
