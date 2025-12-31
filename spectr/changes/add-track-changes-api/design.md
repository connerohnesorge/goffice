# Design: Word Track Changes API Completion

## Context

goffice has **existing Track Changes infrastructure** (~946 LOC) that provides:

**Existing (DO NOT REBUILD):**
- `InsertedRun` (w:ins) with Id, Author, Date, Runs(), AppendRun(), InnerText()
- `DeletedRun` (w:del) with Id, Author, Date, DeletedTexts()
- `DeletedText` (w:delText) with SetText(), Space()
- `MoveFromRun`, `MoveToRun` with Id, Author, Date
- `RunPropertiesChange`, `ParagraphPropertiesChange` with PreviousProperties
- Settings: `TrackRevisions()`, `SetTrackRevisions(bool)`
- 15 comprehensive unit tests

**Missing (BUILD THESE):**
1. Document-level change enumeration
2. Accept/Reject operations
3. Paragraph/Body integration for creating tracked changes
4. Query/filter API
5. Bulk operations

The solution must **enhance existing code**, not replace it.

## Goals

1. **Enable change enumeration** - Query all changes across document
2. **Support accept/reject** - Apply or revert individual changes
3. **Integrate with Paragraph** - Fluent API for creating tracked changes
4. **Support filtering** - By author, date range, change type
5. **Maintain compatibility** - All existing tests pass unchanged

## Non-Goals

- Replace existing InsertedRun/DeletedRun types (they work)
- OOXML conflict resolution (too complex for Phase 1)
- Merge operations (Phase 2)
- Real-time collaboration features
- Change comments/annotations

## Architecture Overview

### Existing Code (DO NOT MODIFY)

```
wordprocessing/elements/
├── revision.go          // InsertedRun, DeletedRun, etc. (946 lines)
└── revision_test.go     // 15 tests

wordprocessing/elements/settings.go
└── TrackRevisions(), SetTrackRevisions()
```

### New Code (ADD THESE)

```
wordprocessing/
├── track_changes.go     // NEW: TrackChangeCollection, operations (~300 LOC)
├── track_change.go      // NEW: TrackChange interface, wrapper (~150 LOC)
└── track_changes_test.go // NEW: Tests for new functionality
```

### Enhancement to Existing Code

```go
// In wordprocessing/paragraph.go - ADD methods:
func (p *Paragraph) AppendInsertedRun(author, text string) *elements.InsertedRun
func (p *Paragraph) AppendDeletedRun(author, text string) *elements.DeletedRun

// In wordprocessing/body.go - ADD methods:
func (b *Body) TrackChanges() *TrackChangeCollection

// In wordprocessing/document.go - ADD methods:
func (d *Document) TrackChanges() *TrackChangeCollection
```

## Detailed Design Decisions

### Decision 1: TrackChange Interface

**Design**: Create interface over existing element types:

```go
// wordprocessing/track_change.go

// TrackChange provides unified access to all track change types
type TrackChange interface {
    // Common metadata
    Id() string
    Author() string
    Date() time.Time

    // Type identification
    Type() TrackChangeType

    // Content access
    Content() string  // InnerText equivalent

    // Operations
    Accept() error
    Reject() error

    // Source element access
    Element() openxml.Element
}

type TrackChangeType int
const (
    TrackChangeInsertion TrackChangeType = iota
    TrackChangeDeletion
    TrackChangeMoveFrom
    TrackChangeMoveTo
    TrackChangePropertyChange
)
```

**Rationale**:
- Unifies access to InsertedRun, DeletedRun, MoveFromRun, MoveToRun
- Enables polymorphic operations (Accept/Reject)
- Matches OpenXML-SDK patterns where changes are accessed uniformly
- Doesn't modify existing element types

### Decision 2: RevisionCollection with Iterator Support

**Design**: Collection type with fluent filtering AND iterator support for memory efficiency:

```go
// wordprocessing/revisions.go

type RevisionCollection struct {
    revisions   []*Revision
    moveTracker *moveTracker  // Pre-built during construction for O(1) move pair lookup
}

// moveTracker enables efficient move pair coordination
type moveTracker struct {
    moveFromById map[int]*Revision  // MoveFrom elements indexed by revision ID
    moveToById   map[int]*Revision  // MoveTo elements indexed by revision ID
}

// Factory
func NewRevisionCollection(revisions []*Revision) *RevisionCollection

// Filtering (returns new collection)
func (c *RevisionCollection) ByAuthor(author string) *RevisionCollection
func (c *RevisionCollection) ByDateRange(from, to time.Time) *RevisionCollection
func (c *RevisionCollection) ByType(typ RevisionType) *RevisionCollection
func (c *RevisionCollection) Insertions() *RevisionCollection
func (c *RevisionCollection) Deletions() *RevisionCollection
func (c *RevisionCollection) Moves() *RevisionCollection
func (c *RevisionCollection) FormatChanges() *RevisionCollection

// Access
func (c *RevisionCollection) All() []*Revision
func (c *RevisionCollection) Count() int
func (c *RevisionCollection) First() *Revision
func (c *RevisionCollection) ById(id int) *Revision

// Iterator for memory-efficient enumeration
func (c *RevisionCollection) Iterator() iter.Seq[*Revision]

// Bulk operations
func (c *RevisionCollection) AcceptAll() error
func (c *RevisionCollection) RejectAll() error
```

**Rationale**:
- Collection API enables fluent filtering: `doc.GetRevisions().ByAuthor("John").Insertions().AcceptAll()`
- Iterator method provides memory-efficient option for large documents
- Filtering returns new collection (immutable pattern, cheap slice operations)
- moveTracker provides O(1) move pair lookup for atomic processing
- Best of both worlds: convenience (collection) AND performance (iterator)
- Bulk operations for common workflows

### Decision 3: Accept/Reject Implementation with Move Pair Coordination

**Design**: In-place modification of document with atomic move processing:

```go
// For InsertedRun.Accept():
// 1. Get parent paragraph
// 2. Replace InsertedRun with its child Runs
// 3. Remove InsertedRun wrapper

func (r *Revision) acceptInsertedRun() error {
    ins := r.element.(*InsertedRun)
    parent := ins.Parent()
    if parent == nil {
        return errors.New("track change has no parent")
    }

    // Get runs from inside InsertedRun
    runs := ins.Runs()

    // Insert runs before InsertedRun in parent
    para, ok := parent.(*Paragraph)
    if !ok {
        return fmt.Errorf("unexpected parent type: %T", parent)
    }

    // Replace element with its contents
    para.ReplaceChild(ins, runs...)

    return nil
}

// For DeletedRun operations:
func (r *Revision) acceptDeletedRun() error {
    // Simply remove DeletedRun and all contents from parent
    del := r.element.(*DeletedRun)
    parent := del.Parent()
    if parent == nil {
        return errors.New("track change has no parent")
    }
    parent.RemoveChild(del)
    return nil
}

func (r *Revision) rejectDeletedRun() error {
    // Reconstruct text from DeletedTexts and restore to parent
    del := r.element.(*DeletedRun)
    parent := del.Parent()
    if parent == nil {
        return errors.New("track change has no parent")
    }

    // Convert DeletedText to regular Text elements
    runs := make([]*Run, 0)
    for deletedText := range del.DeletedTexts() {
        run := NewRun()
        text := NewText()
        text.SetText(deletedText.Text())
        run.AppendChild(text)
        runs = append(runs, run)
    }

    // Replace DeletedRun with reconstructed runs
    para := parent.(*Paragraph)
    para.ReplaceChild(del, runs...)
    return nil
}

// For Move operations with pair coordination:
func (r *Revision) acceptMovePair(pair *Revision) error {
    // Atomic processing of both move ends
    if r.Type() == RevisionTypeMoveFrom {
        moveFrom := r.element.(*MoveFromRun)
        moveTo := pair.element.(*MoveToRun)

        // Remove content from origin
        moveFrom.Parent().RemoveChild(moveFrom)

        // Promote content at destination
        runs := moveTo.Runs()
        moveTo.Parent().ReplaceChild(moveTo, runs...)
    } else {
        // Handle MoveTo first (same logic, reversed)
        moveTo := r.element.(*MoveToRun)
        moveFrom := pair.element.(*MoveFromRun)

        // Remove content from origin
        moveFrom.Parent().RemoveChild(moveFrom)

        // Promote content at destination
        runs := moveTo.Runs()
        moveTo.Parent().ReplaceChild(moveTo, runs...)
    }
    return nil
}

func (r *Revision) rejectMovePair(pair *Revision) error {
    // Atomic processing of both move ends
    if r.Type() == RevisionTypeMoveFrom {
        moveFrom := r.element.(*MoveFromRun)
        moveTo := pair.element.(*MoveToRun)

        // Restore content at origin
        runs := moveFrom.Runs()
        moveFrom.Parent().ReplaceChild(moveFrom, runs...)

        // Remove content from destination
        moveTo.Parent().RemoveChild(moveTo)
    } else {
        // Handle MoveTo first (same logic, reversed)
        moveTo := r.element.(*MoveToRun)
        moveFrom := pair.element.(*MoveFromRun)

        // Restore content at origin
        runs := moveFrom.Runs()
        moveFrom.Parent().ReplaceChild(moveFrom, runs...)

        // Remove content from destination
        moveTo.Parent().RemoveChild(moveTo)
    }
    return nil
}

// For Format changes with property extraction:
func (r *Revision) rejectRunPropertiesChange() error {
    rPrChange := r.element.(*RunPropertiesChange)

    // Get parent Run element (rPr -> Run)
    parentRun := rPrChange.Parent().Parent()

    // Extract previous properties from revision element
    previousProps := rPrChange.RunProperties()
    if previousProps == nil {
        return fmt.Errorf("missing previous properties in revision")
    }

    // Get current properties container
    currentPropsContainer := parentRun.RunProperties()

    // Remove revision wrapper
    currentPropsContainer.RemoveChild(rPrChange)

    // Replace current properties with previous properties
    currentPropsContainer.RemoveAllChildren()
    for child := range previousProps.Children() {
        currentPropsContainer.AppendChild(child.CloneNode())
    }

    return nil
}
```

**Rationale**:
- Accept insertion = promote content, remove wrapper
- Reject insertion = remove wrapper and content
- Accept deletion = remove wrapper and content (text already gone)
- Reject deletion = reconstruct text from DeletedTexts, remove wrapper
- Move pairs processed atomically using pre-built moveTracker for O(1) lookup
- Format changes extract complete property snapshots (no partial conflicts)
- All operations maintain document validity (no orphaned elements)

### Decision 4: Enumeration Strategy with Move Tracker

**Design**: Recursive tree traversal with move pair indexing:

```go
func (d *Document) GetRevisions() *RevisionCollection {
    var revisions []*Revision

    // Traverse all document parts
    body := d.MainDocumentPart().Document().Body()
    collectRevisions(body, &revisions, d)

    // Also traverse headers, footers, footnotes, endnotes, comments
    for _, headerPart := range d.MainDocumentPart().HeaderParts() {
        if header := headerPart.Header(); header != nil {
            collectRevisions(header, &revisions, d)
        }
    }
    // ... similar for footers, footnotes, endnotes, comments

    // Build move tracker for efficient pair lookup
    collection := NewRevisionCollection(revisions)
    return collection
}

func collectRevisions(node openxml.Element, revisions *[]*Revision, doc *Document) {
    // Check if this node is a revision element
    switch e := node.(type) {
    case *elements.InsertedRun:
        *revisions = append(*revisions, newRevision(e, RevisionTypeInsert, node.Parent(), doc))
    case *elements.DeletedRun:
        *revisions = append(*revisions, newRevision(e, RevisionTypeDelete, node.Parent(), doc))
    case *elements.MoveFromRun:
        *revisions = append(*revisions, newRevision(e, RevisionTypeMoveFrom, node.Parent(), doc))
    case *elements.MoveToRun:
        *revisions = append(*revisions, newRevision(e, RevisionTypeMoveTo, node.Parent(), doc))
    case *elements.RunPropertiesChange:
        *revisions = append(*revisions, newRevision(e, RevisionTypeFormatChange, node.Parent(), doc))
    case *elements.ParagraphPropertiesChange:
        *revisions = append(*revisions, newRevision(e, RevisionTypeFormatChange, node.Parent(), doc))
    }

    // Recurse into children
    for _, child := range node.Children() {
        if ce, ok := child.(openxml.Element); ok {
            collectRevisions(ce, revisions, doc)
        }
    }
}

func NewRevisionCollection(revisions []*Revision) *RevisionCollection {
    collection := &RevisionCollection{
        revisions: revisions,
    }

    // Build move tracker for O(1) pair lookup
    collection.moveTracker = buildMoveTracker(revisions)

    // Link collection back to revisions for pair lookup
    for _, rev := range revisions {
        rev.collection = collection
    }

    return collection
}

func buildMoveTracker(revisions []*Revision) *moveTracker {
    tracker := &moveTracker{
        moveFromById: make(map[int]*Revision),
        moveToById:   make(map[int]*Revision),
    }

    for _, rev := range revisions {
        if rev.Type() == RevisionTypeMoveFrom {
            tracker.moveFromById[rev.Id()] = rev
        } else if rev.Type() == RevisionTypeMoveTo {
            tracker.moveToById[rev.Id()] = rev
        }
    }

    return tracker
}
```

**Rationale**:
- Single pass over all document parts (body, headers, footers, etc.)
- Collects all revision types uniformly
- Builds move tracker during collection construction for O(1) pair lookup
- Links revisions to collection for access to move tracker
- Returns consistent RevisionCollection type with pre-built indexes

### Decision 5: Paragraph Integration

**Design**: Convenience methods on Paragraph:

```go
// wordprocessing/paragraph.go

func (p *Paragraph) AppendInsertedRun(author, text string) *elements.InsertedRun {
    // Note: Existing NewInsertedRun(id, author, date) requires 3 args.
    // This convenience method handles generation internally:
    id := p.nextChangeId()  // Implemented in task 1.5.3
    ins := elements.NewInsertedRun(id, author, time.Now())

    run := ins.AppendRun()
    run.AppendText(text)

    p.AppendChild(ins)
    return ins
}

func (p *Paragraph) AppendDeletedRun(author, deletedText string) *elements.DeletedRun {
    // Note: Existing NewDeletedRun(id, author, date) requires 3 args.
    id := p.nextChangeId()  // Document-level atomic counter
    del := elements.NewDeletedRun(id, author, time.Now())

    dt := elements.NewDeletedText()
    dt.SetText(deletedText)
    del.AppendChild(dt)

    p.AppendChild(del)
    return del
}
```

**Rationale**:
- Fluent API: `para.AppendInsertedRun("John", "new text")`
- Auto-generates id, sets current timestamp
- Returns element for further customization

## Module Organization

```
wordprocessing/
├── track_changes.go        // NEW: Collection, enumeration (~200 LOC)
├── track_change.go         // NEW: Interface, wrapper types (~150 LOC)
├── track_changes_test.go   // NEW: Tests (~300 LOC)
├── paragraph.go            // ENHANCE: Add AppendInsertedRun/DeletedRun
├── body.go                 // ENHANCE: Add TrackChanges()
└── document.go             // ENHANCE: Add TrackChanges()
```

**Total new code**: ~650 LOC
**Total enhanced code**: ~50 LOC modifications

## Testing Strategy

### Unit Tests (Per-Feature)

**TrackChange Interface** (8 tests):
- [ ] InsertedRun implements TrackChange correctly
- [ ] DeletedRun implements TrackChange correctly
- [ ] MoveFromRun implements TrackChange correctly
- [ ] MoveToRun implements TrackChange correctly
- [ ] Type() returns correct TrackChangeType
- [ ] Content() returns InnerText
- [ ] Element() returns underlying element
- [ ] Date() returns RFC3339 parsed time

**Accept/Reject Operations** (12 tests):
- [ ] Accept InsertedRun promotes content to parent
- [ ] Accept InsertedRun removes wrapper element
- [ ] Reject InsertedRun removes wrapper and content
- [ ] Accept DeletedRun removes wrapper and content
- [ ] Reject DeletedRun reconstructs text
- [ ] Accept on orphan element returns error
- [ ] Nested changes handled correctly
- [ ] Multiple accepts in sequence work
- [ ] AcceptAll processes all changes
- [ ] RejectAll processes all changes
- [ ] Partial failure doesn't corrupt document
- [ ] Accept/Reject updates document state

**Collection Filtering** (10 tests):
- [ ] ByAuthor filters correctly
- [ ] ByAuthor with no matches returns empty
- [ ] ByDateRange includes boundary dates
- [ ] Insertions() filters to InsertedRun only
- [ ] Deletions() filters to DeletedRun only
- [ ] Moves() includes MoveFrom and MoveTo
- [ ] Chained filters work correctly
- [ ] Count() returns correct count
- [ ] First() returns nil for empty
- [ ] ById() finds correct change

**Paragraph Integration** (6 tests):
- [ ] AppendInsertedRun creates valid element
- [ ] AppendInsertedRun sets all attributes
- [ ] AppendDeletedRun creates valid element
- [ ] AppendDeletedRun sets all attributes
- [ ] Change appears in para.Children()
- [ ] Change appears in doc.TrackChanges()

### Integration Tests (8 scenarios)

- [ ] Open Word doc with changes → enumerate → verify
- [ ] Create changes → save → open → enumerate → verify roundtrip
- [ ] Accept all changes → save → verify clean document
- [ ] Reject all changes → save → verify reverted document
- [ ] Mixed operations (some accept, some reject)
- [ ] Large document with 1000+ changes
- [ ] Document with nested changes
- [ ] Document with move ranges

## API Surface Summary

### New Type: Revision

```go
type Revision struct {
    element    interface{}              // Underlying element (InsertedRun, DeletedRun, etc.)
    typ        RevisionType
    parent     openxml.Element          // Parent for removal/manipulation
    document   *WordprocessingDocument  // For document-wide operations
    collection *RevisionCollection      // For move pair lookup
}

func (r *Revision) Id() int
func (r *Revision) Author() string
func (r *Revision) Date() time.Time
func (r *Revision) Type() RevisionType
func (r *Revision) Content() string
func (r *Revision) Accept() error
func (r *Revision) Reject() error
func (r *Revision) Element() openxml.Element

type RevisionType int
const (
    RevisionTypeInsert RevisionType = iota
    RevisionTypeDelete
    RevisionTypeMoveFrom
    RevisionTypeMoveTo
    RevisionTypeFormatChange
)
```

### New Type: RevisionCollection

```go
type RevisionCollection struct {
    revisions   []*Revision
    moveTracker *moveTracker  // Pre-built during construction
}

func NewRevisionCollection(revisions []*Revision) *RevisionCollection

// Filtering (returns new collection)
func (c *RevisionCollection) ByAuthor(author string) *RevisionCollection
func (c *RevisionCollection) ByDateRange(from, to time.Time) *RevisionCollection
func (c *RevisionCollection) ByType(typ RevisionType) *RevisionCollection
func (c *RevisionCollection) Insertions() *RevisionCollection
func (c *RevisionCollection) Deletions() *RevisionCollection
func (c *RevisionCollection) Moves() *RevisionCollection
func (c *RevisionCollection) FormatChanges() *RevisionCollection

// Access
func (c *RevisionCollection) All() []*Revision
func (c *RevisionCollection) Count() int
func (c *RevisionCollection) First() *Revision
func (c *RevisionCollection) ById(id int) *Revision

// Iterator for memory-efficient enumeration
func (c *RevisionCollection) Iterator() iter.Seq[*Revision]

// Bulk operations
func (c *RevisionCollection) AcceptAll() error
func (c *RevisionCollection) RejectAll() error
```

### New Methods on WordprocessingDocument

```go
func (d *WordprocessingDocument) GetRevisions() *RevisionCollection
func (d *WordprocessingDocument) AcceptAllRevisions() error
func (d *WordprocessingDocument) RejectAllRevisions() error
```

### New Methods on Paragraph

```go
func (p *Paragraph) AppendInsertedRun(author, text string) *elements.InsertedRun
func (p *Paragraph) AppendDeletedRun(author, deletedText string) *elements.DeletedRun
```

## Design Resolution: Iterator vs Collection Pattern

### Resolution of API Pattern Conflict

**Original Conflict**:
- proposal.md initially advocated `iter.Seq[*Revision]` for memory efficiency
- design.md showed `ByAuthor()`, `ByDateRange()` collection methods
- These appeared contradictory

**Resolution**: **Both patterns coexist harmoniously**

### Hybrid Approach: Collection with Iterator Method

```go
// Primary API: Collection-based (convenient, fluent)
revisions := doc.GetRevisions()                    // Returns *RevisionCollection
johnRevisions := revisions.ByAuthor("John Doe")    // Fluent filtering
johnRevisions.Insertions().AcceptAll()             // Chainable operations

// Optional API: Iterator-based (memory-efficient for large docs)
for revision := range doc.GetRevisions().Iterator() {  // iter.Seq[*Revision]
    if revision.Author() == "John Doe" {
        revision.Accept()
    }
}
```

### Why This Works

**Collection Benefits**:
- Convenient fluent API for common use cases
- Filtering returns new collections (immutable pattern)
- Pre-built indexes (moveTracker) shared across filtered collections
- Bulk operations on filtered sets

**Iterator Benefits**:
- Memory-efficient for very large documents (10,000+ revisions)
- Early termination without processing entire document
- Consistent with Go 1.23+ idioms
- Lazy evaluation when needed

**No Conflict**:
- `GetRevisions()` performs full scan once, returns collection
- Collection stores slice of revisions (already in memory)
- `Iterator()` method on collection yields from existing slice
- Filtering creates new collection (cheap - just slice filtering)
- Both use same underlying data structure

### Performance Characteristics

**GetRevisions()**: O(n) document traversal, builds collection once
**Filtering**: O(m) where m = collection size (cheap slice operations)
**Iterator()**: O(1) to start, yields from existing slice
**moveTracker**: Built once during GetRevisions(), O(1) lookups thereafter

### Usage Patterns

**Pattern 1: Fluent Collection API** (recommended for most cases)
```go
// Accept all insertions from John Doe
doc.GetRevisions().ByAuthor("John Doe").Insertions().AcceptAll()

// Get count of deletions in date range
count := doc.GetRevisions().ByDateRange(start, end).Deletions().Count()

// Find specific revision by ID
rev := doc.GetRevisions().ById(42)
```

**Pattern 2: Iterator for Large Documents** (when memory matters)
```go
// Process revisions one at a time without holding all in memory
// Note: GetRevisions() still scans document, but iterator avoids
// additional allocations during processing
for rev := range doc.GetRevisions().Iterator() {
    if shouldProcess(rev) {
        rev.Accept()
        break  // Early termination
    }
}
```

**Pattern 3: Filtered Iterator** (best of both)
```go
// Filter first (collection), then iterate efficiently
johnRevisions := doc.GetRevisions().ByAuthor("John Doe")
for rev := range johnRevisions.Iterator() {
    // Process John's revisions one at a time
    processRevision(rev)
}
```

## Backward Compatibility

**Unchanged APIs** (fully backward compatible):
- All `InsertedRun`, `DeletedRun`, etc. methods unchanged
- All `revision_test.go` tests pass without modification
- Settings `TrackRevisions()`, `SetTrackRevisions()` unchanged

**New APIs** (additive only):
- `TrackChange` interface (new)
- `TrackChangeCollection` type (new)
- Document/Body/Paragraph convenience methods (new)

**No Breaking Changes**: Existing code continues to work unchanged.

## Success Metrics

1. **Backward compatible**: All 15 revision_test.go tests pass unchanged
2. **Enumeration works**: `doc.TrackChanges().Count()` returns correct count
3. **Filtering works**: `doc.TrackChanges().ByAuthor("John")` filters correctly
4. **Accept works**: Inserted content promoted, wrapper removed
5. **Reject works**: Deleted content restored, wrapper removed
6. **Integration works**: `para.AppendInsertedRun()` creates valid change
7. **Performance**: 1000 changes enumerated <100ms

## Phase 2 Extensions (NOT IN THIS CHANGE)

- Move range validation and iteration
- Property change diff/comparison
- Conflict detection for overlapping changes
- Change history/timeline view
- Comment integration
