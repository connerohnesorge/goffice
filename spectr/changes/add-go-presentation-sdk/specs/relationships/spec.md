# Spec: Relationships

## Overview

Defines the relationship management system for OPC (Open Packaging Conventions) packages. Relationships connect parts within a package and link to external resources, stored in `_rels/*.rels` files.

## API

### Types

```go
// Relationship represents a relationship between parts or to external resources
type Relationship struct {
    ID         string
    Type       string
    Target     string
    TargetMode TargetMode // Internal or External
}

// TargetMode indicates whether a relationship target is internal or external
type TargetMode int

const (
    TargetModeInternal TargetMode = iota
    TargetModeExternal
)

// RelationshipCollection manages relationships for a part or package
type RelationshipCollection interface {
    // Get returns a relationship by ID
    Get(id string) (*Relationship, error)

    // GetByType returns all relationships of a given type
    GetByType(relType string) []*Relationship

    // Add creates a new relationship
    Add(relType, target string, targetMode TargetMode) *Relationship

    // AddWithID creates a new relationship with a specific ID
    AddWithID(id, relType, target string, targetMode TargetMode) (*Relationship, error)

    // Remove deletes a relationship by ID
    Remove(id string) error

    // All returns all relationships
    All() []*Relationship
}
```

### Common Relationship Types

```go
const (
    RelTypeOfficeDocument     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
    RelTypeSlide              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slide"
    RelTypeSlideMaster        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster"
    RelTypeSlideLayout        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideLayout"
    RelTypeTheme              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme"
    RelTypeImage              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"
    RelTypeHyperlink          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"
    RelTypeNotesMaster        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesMaster"
    RelTypeNotesSlide         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesSlide"
    RelTypeHandoutMaster      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/handoutMaster"
)
```

## Implementation Notes

- Relationships are stored in XML files under `_rels/` directories
- Package-level relationships are in `_rels/.rels`
- Part-level relationships are in `_rels/<partname>.rels`
- Relationship IDs must be unique within their containing `.rels` file
- Internal relationships use relative URIs from the source part
- External relationships use absolute URIs
- Hyperlinks are a special case of external relationships
