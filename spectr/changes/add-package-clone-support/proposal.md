# Add Package Clone Support

## Overview
Implement deep cloning of OpenXML packages and parts with proper relationship preservation. Enables efficient document duplication, template instantiation, and safe concurrent access.

## Motivation
Open-XML-SDK provides cloning for creating independent copies. Essential for mail merge, template processing, and concurrent document generation. Currently missing in goffice.

## Goals
- Implement Clone() methods for OpenXmlPackage and all part types
- Deep clone all XML content with proper type preservation
- Clone and rewrite relationships (internal IDs must be unique)
- Clone content types and package metadata
- Support selective cloning (clone specific parts only)
- Handle binary parts correctly (images, OLE objects, etc.)
- Support cloning to memory or new file

## Non-Goals
- Incremental cloning or copy-on-write
- Cross-document part sharing
- Optimization for very large documents (initial)

## Dependencies
- Depends on: packaging core, all element types
- Blocks: add-mail-merge-implementation
- Related: add-flatopc-support

## Technical Approach

### Clone API
```go
// Clone entire document
newDoc := doc.Clone()

// Clone to new file
doc.CloneAs("copy.docx")

// Clone specific parts
newDoc := doc.CloneWithParts(MainDocumentPart, StylesPart)
```

### Relationship Rewriting
```go
func CloneRelationships(src, dst Part) {
    idMap := make(map[string]string)
    for _, rel := range src.Relationships() {
        newID := GenerateUniqueID()
        idMap[rel.ID] = newID
        // Clone relationship with new ID
    }
    // Update all references in cloned content
    RewriteReferences(dst, idMap)
}
```

## Estimated Effort
3 weeks (1 engineer)

## Priority  
P0 - Essential for template processing
