# Design: Advanced Relationship Management

## Overview
Comprehensive relationship system matching Open-XML-SDK with multiple types, external targets, and bidirectional tracking.

## Core Types & Interfaces

### Relationship Types
```go
type RelationshipType string
type Relationship interface {
    ID() string
    Type() RelationshipType
    Target() RelationshipTarget
    IsExternal() bool
    GetPart() (Part, error)
}

type RelationshipTarget interface {
    URI() string
    IsExternal() bool
}

type InternalTarget struct {
    PartURI string
}

type ExternalTarget struct {
    URL        string
    TargetMode string
}
```

### Relationship Container
```go
type RelationshipCollection interface {
    Add(RelationshipType, RelationshipTarget) (*Relationship, error)
    Remove(id string) error
    GetById(id string) (Relationship, error)
    GetByType(RelationshipType) []Relationship
    All() []Relationship
}
```

## Architectural Decisions

### 1. Relationship Storage
**Decision**: In-memory with lazy part loading

**Rationale**:
- Parsed from .rels files on demand
- Parts loaded only on access
- Reduces memory usage

### 2. Bidirectional Tracking
**Decision**: Maintain reverse index

**Rationale**:
- Support parent queries
- Track dependencies
- Enable cleanup

### 3. External vs Internal
**Decision**: Polymorphic resolution

**Rationale**:
- Different validation rules
- Different serialization

## Implementation Strategy

### Phase 1: Core (1 week)
- [ ] Relationship interfaces
- [ ] Target types
- [ ] RelationshipManager
- [ ] .rels parsing

### Phase 2: Advanced (1 week)
- [ ] Bidirectional tracking
- [ ] Type queries
- [ ] Lazy loading
- [ ] Dependency detection

### Phase 3: Integration (1 week)
- [ ] Packaging layer
- [ ] Validation hooks
- [ ] Modification triggers

## Performance Characteristics

- Lazy loading reduces initial memory
- Reverse index enables O(1) lookups
- Minimal overhead for relationship tracking
