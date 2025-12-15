# Go Word SDK Research Summary

## Overview

This directory contains detailed research findings from exploring the Microsoft Open-XML-SDK C# implementation. The research was conducted to inform the design of the `goffice` Go SDK for Office Open XML Word Processing (WordprocessingML).

## Research Documents

| File | Topic | Key Findings |
|------|-------|--------------|
| [01-element-system.md](01-element-system.md) | OpenXmlElement Architecture | Circular linked list for children, dual state (lazy/parsed), MiscAttrContainer |
| [02-attribute-system.md](02-attribute-system.md) | Attribute Management | Fixed vs extended attributes, AttributeCollection, OpenXmlSimpleType hierarchy |
| [03-relationships.md](03-relationships.md) | Relationship System | Internal vs external, lazy loading, HyperlinkRelationship, DataPartReferenceRelationship |
| [04-content-types.md](04-content-types.md) | Content Types | [Content_Types].xml, PartExtensionProvider, Word-specific part types |
| [05-json-schema-format.md](05-json-schema-format.md) | Code Generation Schema | JSON schema files, particle system, validator types |
| [06-opensettings.md](06-opensettings.md) | Configuration Options | OpenSettings, FileFormatVersions, PackageCapabilities |
| [07-drawingml-integration.md](07-drawingml-integration.md) | DrawingML Integration | Inline/anchor drawings, shapes in Word documents |
| [08-flatopc-format.md](08-flatopc-format.md) | FlatOPC Format | Single XML representation, binary encoding, conversion methods |
| [09-error-handling.md](09-error-handling.md) | Error Handling | Custom exceptions, validation errors, SR.Format pattern |
| [10-wordprocessingml-specifics.md](10-wordprocessingml-specifics.md) | WordprocessingML-Specific | Document structure, paragraphs, runs, tables, styles |

## Key Architectural Patterns

### 1. Element System (Shared with PresentationML)

- **Circular Linked List**: Children stored via `_lastChild` pointer with `Next` references forming a circle
- **Lazy Loading**: Elements can store raw XML and parse on demand
- **Attribute Separation**: Fixed (schema) vs extended (unknown) attributes
- **Feature Collection**: Type-based extensibility via `IFeatureCollection`

### 2. Packaging Layer (Shared with PresentationML)

- **OPC Abstraction**: ZIP-based package wrapped in `OpenXmlPackage`
- **Part Hierarchy**: `OpenXmlPart` -> `TypedPart` -> Specific parts
- **Relationship Management**: Lazy-loaded with dictionary caching
- **Content Types**: Default (extension) and Override (URI) mappings

### 3. WordprocessingML-Specific Architecture

- **Document Body**: Main content container (`w:body`)
- **Block-level Elements**: Paragraphs, tables, sections
- **Inline Elements**: Runs, text, formatting
- **Style Inheritance**: Document defaults -> styles -> direct formatting

### 4. Validation System (Shared with PresentationML)

- **Schema Validation**: Particle-based child ordering
- **Semantic Constraints**: 21 constraint types for business rules
- **Non-Fatal Errors**: ValidationErrorInfo collected, not thrown
- **Version Targeting**: FileFormatVersions for Office compatibility

## Resolved Design Decisions

Based on user input and research:

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Go Module | `github.com/connerohnesorge/goffice` | Personal namespace |
| Go Version | 1.22+ | Latest features, mature generics |
| Dependencies | Pure stdlib only | No supply chain risk |
| Generated Code | Pre-generated, committed | No generator required by users |
| API Style | Direct struct fields | Go idiomatic |
| Error Handling | Return errors everywhere | Explicit error handling |
| Validation | Strict by default | Reject invalid content |
| Thread Safety | Package-level sync.Mutex | Document-level locking |

## Implementation Priority

### Phase 1: Core Framework (Shared)
1. Element types (OpenXmlElement, CompositeElement, LeafElement)
2. Attribute system (fixed, extended, MC)
3. Child management (circular linked list)
4. XML serialization/deserialization

### Phase 2: Simple Types (Shared)
1. Value types (StringValue, Int32Value, etc.)
2. EnumValue generic
3. HexBinaryValue, Base64BinaryValue

### Phase 3: Packaging (Shared)
1. OPC package handling (ZIP-based)
2. Part management
3. Relationship handling
4. Content types

### Phase 4: Word Document Parts
1. WordprocessingDocument
2. MainDocumentPart
3. StylesPart, NumberingPart
4. HeaderPart, FooterPart
5. Media parts

### Phase 5: Schema Elements (Word-Specific)
1. Document root elements
2. Body, Paragraph, Run, Text structure
3. Tables (TableRow, TableCell)
4. Sections and page setup

### Phase 6: Validation (Shared)
1. Particle validation
2. Attribute validation
3. Semantic constraints

### Phase 7: Advanced Features
1. FlatOPC support
2. LINQ-style queries
3. Clone operations

## Source Code Locations

Key C# SDK files to explore:

```
Open-XML-SDK/
├── src/DocumentFormat.OpenXml.Framework/
│   ├── OpenXmlElement.cs
│   ├── OpenXmlCompositeElement.cs
│   ├── OpenXmlLeafElement.cs
│   ├── SimpleTypes/
│   ├── Packaging/
│   │   ├── OpenXmlPackage.cs
│   │   ├── OpenXmlPart.cs
│   │   ├── PartRelationshipsFeature.cs
│   │   └── FlatOpcExtensions.cs
│   ├── Features/
│   ├── Validation/
│   └── Framework/
│       ├── Metadata/
│       └── Schema/
├── src/DocumentFormat.OpenXml/
│   └── GeneratedCode/
│       └── schemas/
│           └── wordprocessingml/ (Word-specific schemas)
├── data/
│   ├── schemas/*.json (155 files)
│   ├── namespaces.json
│   └── parts/*.json
└── generated/
    └── *.g.cs
```

## Next Steps

1. **Research WordprocessingML specifics** - document structure, paragraph model, table model
2. Create Go package structure matching architecture
3. Implement core element types with embedded structs (shared with Presentation)
4. Build attribute system with generics (shared)
5. Implement OPC package handling using archive/zip (shared)
6. Generate Word element types from JSON schemas
7. Add validation layer
8. Write comprehensive tests
