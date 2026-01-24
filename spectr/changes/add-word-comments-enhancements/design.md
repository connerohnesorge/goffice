# Design: Word Comments Enhancements

## Overview
Implementation of comprehensive Word processing comments support, including threaded comments, authors, dates, and range tracking.

## Core Types & Interfaces

### Comment Structure
```go
type Comment struct {
    ID          int
    Author      string
    Initials    string
    Date        time.Time
    Content     []BlockElement // Paragraphs, Tables, etc.
    ParentID    *int          // For threaded replies
}

type CommentsPart struct {
    Comments []*Comment
}
```

### Range Tracking
```go
type CommentRangeStart struct {
    ID int
}

type CommentRangeEnd struct {
    ID int
}
```

## Architectural Decisions

### 1. Comment Storage
**Decision**: Separate part management for comments

**Rationale**:
- Comments are stored in a separate XML part
- Reduces main document complexity
- Enables lazy loading of comments

### 2. Threading Model
**Decision**: Adjacency list with ParentID

**Rationale**:
- Simple to serialize to XML
- Easy to reconstruct conversation trees
- Matches OpenXML structure

## Implementation Strategy

### Phase 1: Basic Comments
- [ ] Comment part definition
- [ ] Comment serialization
- [ ] Author/Date handling

### Phase 2: Ranges & Anchors
- [ ] CommentRangeStart/End elements
- [ ] Reference integration in document body
- [ ] Range validation

### Phase 3: Advanced Features
- [ ] Threaded replies
- [ ] Extended properties
- [ ] Comment history

## Testing Strategy
- Round-trip serialization tests
- Multi-author comment scenarios
- Threaded conversation reconstruction