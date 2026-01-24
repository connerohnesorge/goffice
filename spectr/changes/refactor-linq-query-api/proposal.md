# Change: LINQ-Style Query API for Element Traversal

## Why
Open-XML-SDK provides LINQ (Language Integrated Query) support for element traversal and filtering. While goffice has basic traversal, a Go equivalent using iterators and functional patterns would improve discoverability and usability.

## What Changes
- Iterator-based element queries (descendants, ancestors, children)
- Filter combinators (where, select, first, last, count)
- Sibling and relative navigation
- Element collection queries with lazy evaluation
- Axis-based navigation (parent, child, attribute)
- Query composition and chaining
- Element indexing and position queries
- Distinct and unique element queries
- Query performance optimization (lazy vs. eager)

## Impact
- Affected specs: linq-integration, framework
- Affected code: openxml/element.go, openxml/query/*.go (new)
- Breaking changes: None (additive, parallel API)
