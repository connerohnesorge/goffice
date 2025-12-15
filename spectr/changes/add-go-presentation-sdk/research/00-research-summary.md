# Go Presentation SDK Research Summary

## Overview

This directory contains detailed research findings from exploring the Microsoft Open-XML-SDK C# implementation. The research was conducted to inform the design of the `goffice` Go SDK for Office Open XML Presentations.

## Research Documents

| File | Topic | Key Findings |
|------|-------|--------------|
| [01-element-system.md](01-element-system.md) | OpenXmlElement Architecture | Circular linked list for children, dual state (lazy/parsed), MiscAttrContainer |
| [02-attribute-system.md](02-attribute-system.md) | Attribute Management | Fixed vs extended attributes, AttributeCollection, OpenXmlSimpleType hierarchy |
| [03-relationships.md](03-relationships.md) | Relationship System | Internal vs external, lazy loading, HyperlinkRelationship, DataPartReferenceRelationship |
| [04-content-types.md](04-content-types.md) | Content Types | [Content_Types].xml, PartExtensionProvider, 14 presentation part types |
| [05-json-schema-format.md](05-json-schema-format.md) | Code Generation Schema | 155 JSON schema files, particle system, validator types |
| [06-opensettings.md](06-opensettings.md) | Configuration Options | OpenSettings, FileFormatVersions, PackageCapabilities |
| [07-drawingml-integration.md](07-drawingml-integration.md) | DrawingML Integration | Transform2D, fills, text body, preset geometry |
| [08-flatopc-format.md](08-flatopc-format.md) | FlatOPC Format | Single XML representation, binary encoding, conversion methods |
| [09-error-handling.md](09-error-handling.md) | Error Handling | Custom exceptions, validation errors, SR.Format pattern |

## Key Architectural Patterns

### 1. Element System

- **Circular Linked List**: Children stored via `_lastChild` pointer with `Next` references forming a circle
- **Lazy Loading**: Elements can store raw XML and parse on demand
- **Attribute Separation**: Fixed (schema) vs extended (unknown) attributes
- **Feature Collection**: Type-based extensibility via `IFeatureCollection`

### 2. Packaging Layer

- **OPC Abstraction**: ZIP-based package wrapped in `OpenXmlPackage`
- **Part Hierarchy**: `OpenXmlPart` → `TypedPart` → Specific parts
- **Relationship Management**: Lazy-loaded with dictionary caching
- **Content Types**: Default (extension) and Override (URI) mappings

### 3. Validation System

- **Schema Validation**: Particle-based child ordering
- **Semantic Constraints**: 21 constraint types for business rules
- **Non-Fatal Errors**: ValidationErrorInfo collected, not thrown
- **Version Targeting**: FileFormatVersions for Office compatibility

### 4. Code Generation

- **JSON Schemas**: 155 files defining all element types
- **Particle System**: Sequence, Choice, All, Group for content models
- **Validators**: RequiredValidator, StringValidator, NumberValidator, etc.
- **Generated Code**: Pre-generated C# committed to repo

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

### Phase 1: Core Framework
1. Element types (OpenXmlElement, CompositeElement, LeafElement)
2. Attribute system (fixed, extended, MC)
3. Child management (circular linked list)
4. XML serialization/deserialization

### Phase 2: Simple Types
1. Value types (StringValue, Int32Value, etc.)
2. EnumValue generic
3. HexBinaryValue, Base64BinaryValue

### Phase 3: Packaging
1. OPC package handling (ZIP-based)
2. Part management
3. Relationship handling
4. Content types

### Phase 4: Presentation Parts
1. PresentationDocument
2. SlidePart, SlideLayoutPart, SlideMasterPart
3. ThemePart
4. Media parts

### Phase 5: Schema Elements
1. Presentation root elements
2. Slide structure (ShapeTree, Shape, TextBody)
3. DrawingML integration

### Phase 6: Validation
1. Particle validation
2. Attribute validation
3. Semantic constraints

### Phase 7: Advanced Features
1. FlatOPC support
2. LINQ-style queries
3. Clone operations

## Source Code Locations

Key C# SDK files explored:

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
├── data/
│   ├── schemas/*.json (155 files)
│   ├── namespaces.json
│   └── parts/*.json
└── generated/
    └── *.g.cs
```

## Next Steps

1. Create Go package structure matching architecture
2. Implement core element types with embedded structs
3. Build attribute system with generics
4. Implement OPC package handling using archive/zip
5. Generate presentation element types from JSON schemas
6. Add validation layer
7. Write comprehensive tests
