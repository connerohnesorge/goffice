# Word SDK Research Directory

This directory contains research notes and findings for implementing the Go Word SDK (WordprocessingML support) in the goffice project.

## Purpose

The research documents in this directory capture detailed analysis of the Microsoft Open-XML-SDK C# implementation to inform the design and development of the Go Word SDK.

## Research Documents

| File | Description |
|------|-------------|
| [00-research-summary.md](00-research-summary.md) | High-level summary of all research findings and implementation priorities |
| [01-element-system.md](01-element-system.md) | OpenXmlElement architecture, circular linked list, lazy loading |
| [02-attribute-system.md](02-attribute-system.md) | Fixed vs extended attributes, attribute collection, value types |
| [03-relationships.md](03-relationships.md) | Part relationships, internal vs external, lazy loading |
| [04-content-types.md](04-content-types.md) | Content types, Word-specific parts, PartExtensionProvider |
| [05-json-schema-format.md](05-json-schema-format.md) | Code generation schema format, particles, validators |
| [06-opensettings.md](06-opensettings.md) | Configuration options, file format versions, package capabilities |
| [07-drawingml-integration.md](07-drawingml-integration.md) | DrawingML in Word: inline/anchor drawings, shapes, text wrapping |
| [08-flatopc-format.md](08-flatopc-format.md) | Single-XML document format, conversion methods |
| [09-error-handling.md](09-error-handling.md) | Custom exceptions, validation errors, error messages |
| [10-wordprocessingml-specifics.md](10-wordprocessingml-specifics.md) | Word-specific elements: paragraphs, runs, tables, sections, styles |

## Shared vs Word-Specific Research

Many of the research documents (01-06, 08-09) cover shared functionality that applies to all Office document types (Word, Excel, PowerPoint). These topics are part of the core OpenXML framework:

- Element system
- Attribute system
- Relationships
- Content types (core)
- JSON schema format
- OpenSettings
- FlatOPC format
- Error handling

The Word-specific research documents are:

- **04-content-types.md** - Includes Word-specific part types
- **07-drawingml-integration.md** - Word-specific drawing integration (inline/anchor)
- **10-wordprocessingml-specifics.md** - Word document structure, paragraphs, tables, sections

## How to Use

1. Start with **00-research-summary.md** for an overview
2. For shared framework understanding, review documents 01-06
3. For Word-specific implementation, focus on documents 07 and 10
4. Reference individual documents as needed during implementation

## Contributing Research

When adding new research findings:

1. Update the relevant existing document, or
2. Create a new numbered document following the naming convention
3. Update this README and 00-research-summary.md with the new document

## Reference Sources

- [Microsoft Open-XML-SDK](https://github.com/dotnet/Open-XML-SDK) - C# reference implementation
- [ECMA-376 Standard](https://www.ecma-international.org/publications-and-standards/standards/ecma-376/) - Office Open XML specification
- `/Open-XML-SDK/data/schemas/*.json` - JSON schema definitions
- `/Open-XML-SDK/src/` - C# source code
