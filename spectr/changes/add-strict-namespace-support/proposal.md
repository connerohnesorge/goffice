# Add Strict Namespace Support

## Overview
Implement comprehensive support for ISO/IEC 29500 Strict namespace variant, matching Open-XML-SDK IStrictNamespaceFeature. This enables reading and writing documents in strict conformance mode with automatic namespace translation between Transitional and Strict variants.

## Motivation
The Open-XML-SDK supports both Transitional (ECMA-376) and Strict (ISO 29500) namespace variants. Current goffice primarily targets Transitional mode. Full strict support is needed for ISO 29500 compliance required by government and enterprise environments.

## Goals
- Implement IStrictNamespaceFeature for namespace mode tracking
- Support automatic namespace translation (Strict ↔ Transitional)
- Detect document namespace mode on load
- Enable explicit strict mode for new documents  
- Update all element types to support both namespace variants
- Validate strict mode constraints (stricter than transitional)
- Provide migration helpers (transitional → strict)

## Non-Goals
- Dual-mode elements within single document
- Automatic strict mode selection (require explicit opt-in)
- Strict mode as default (keep transitional as default)

## Dependencies
- Depends on: openxml core, all element implementations
- Blocks: ISO 29500 compliance certification
- Related: add-validation-semantic-constraints

## Technical Approach

### Namespace Mapping
```go
var strictToTransitional = map[string]string{
    "http://purl.oclc.org/ooxml/wordprocessingml/main": "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
    // ... all namespace mappings
}
```

### Strict Mode Detection
```go
func DetectNamespaceMode(doc *Document) NamespaceMode {
    // Check root element namespace
    // Return Strict or Transitional
}
```

### Auto-Translation
```go
func TranslateNamespaces(elem CommonWrapper, mode NamespaceMode) {
    // Walk tree, translate all namespaces
}
```

## Estimated Effort
6 weeks (1 engineer)

## Priority
P0 - Required for ISO 29500 compliance
