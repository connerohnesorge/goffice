# Add OpenXML LINQ Support

## Overview
Implement LINQ-style query and manipulation API for OpenXML elements, matching the DocumentFormat.OpenXml.Linq package functionality. This provides a more functional and composable way to work with OpenXML documents using XElement/XDocument semantics.

## Motivation
The Open-XML-SDK provides DocumentFormat.OpenXml.Linq package that enables LINQ-to-XML style operations on OpenXML documents. This is missing in goffice. Users who prefer functional programming patterns or need to perform complex queries would benefit from this abstraction.

## Goals
- Provide XElement-based representation of OpenXML parts
- Enable LINQ-style queries over document structure
- Support bidirectional conversion between DOM and LINQ representations
- Maintain compatibility with standard xml.Encoder/Decoder
- Zero external dependencies

## Non-Goals
- Full reimplementation of LINQ query operators (Go has no LINQ)
- Database-style query language
- Performance optimization over DOM API (this is a convenience layer)

## Dependencies
- Depends on: openxml core framework
- Blocks: N/A
- Related: add-openxml-equality-comparison, add-element-query-helpers

## Risks & Mitigations
- **Risk**: Go lacks LINQ operators, making API less intuitive than C# version
  - **Mitigation**: Provide common query helpers and iterator patterns
- **Risk**: Memory overhead from dual representation
  - **Mitigation**: Lazy conversion on-demand, allow GC of unused representation
