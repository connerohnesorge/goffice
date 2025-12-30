# Proposal Brief: Word Track Changes API

## Priority: TIER 1 - Critical
**Impact**: 9/10  
**Effort**: 3-4 weeks  
**Status**: Elements exist, NO high-level API

## Gap Analysis

### Current State
**What exists:**
- ✅ **Elements fully generated** (wordprocessing/elements/revision.go - 947 lines):
  - `InsertedRun` with metadata (id, author, date)
  - `DeletedRun` with DeletedText support
  - `MoveFromRun` / `MoveToRun`
  - `RunPropertiesChange` / `ParagraphPropertiesChange`
  - Iterator-based API for traversing revisions
  - InnerText extraction
- ✅ **Test evidence**: Can parse documents with track changes (revision elements roundtrip correctly)

**What's missing:**
- ❌ No `AcceptRevisions()` method
- ❌ No `RejectRevisions()` method  
- ❌ No `AcceptRevisionsByAuthor(author string)`
- ❌ No `GetRevisions() iter.Seq[*Revision]`
- ❌ No high-level revision management API
- ❌ No convenience methods on Document type

### OpenXML-SDK Comparison
**Microsoft's Open-XML-SDK provides:**
```csharp
// Document-level operations
wordDoc.MainDocumentPart.GetRevisions()
wordDoc.AcceptAllRevisions()
wordDoc.RejectAllRevisions()

// Revision-level operations  
revision.Accept()
revision.Reject()
revision.Author
revision.Date
```

**Our gap**: We have the element classes but no document-level or revision-level methods.

## Implementation Scope

### Files to Create
1. **wordprocessing/revisions.go** (NEW)
   - Revision wrapper class
   - RevisionType enum (Insert, Delete, MoveFrom, MoveTo, FormatChange)
   - Accept/Reject logic

2. **wordprocessing/document.go** (MODIFY)
   - Add `GetRevisions() iter.Seq[*Revision]`
   - Add `AcceptAllRevisions()`
   - Add `RejectAllRevisions()`
   - Add `AcceptRevisionsByAuthor(author string)`
   - Add `RejectRevisionsByAuthor(author string)`

### Core Design Decisions

**1. Accept/Reject Implementation**
- **Accept Insert**: Keep inserted content, remove revision markup
- **Reject Insert**: Remove inserted content entirely
- **Accept Delete**: Remove deleted content and markup
- **Reject Delete**: Restore deleted content, remove markup
- **Accept Move**: Finalize move operation
- **Reject Move**: Revert to original positions

**2. Revision Traversal**
Use Go iterators (iter.Seq) for memory efficiency:
```go
for revision := range doc.GetRevisions() {
    if revision.Author() == "John Doe" {
        revision.Accept()
    }
}
```

**3. Wrapper Pattern**
Wrap generated elements in high-level Revision type:
```go
type Revision struct {
    element  elements.RevisionElement  // Interface
    typ      RevisionType
    id       int64
    author   string
    date     time.Time
}

func (r *Revision) Accept() error { ... }
func (r *Revision) Reject() error { ... }
```

## References

### Existing Code
- `wordprocessing/elements/revision.go` - All revision element types
- `wordprocessing/document.go` - Document API (add methods here)
- `wordprocessing/parts/main_part.go` - MainDocumentPart

### OpenXML-SDK References
Look at @OpenXML-SDK/ for:
- `DocumentFormat.OpenXml.Wordprocessing.Revision` class
- `MainDocumentPart.GetRevisions()` method
- Revision accept/reject logic

### Test Data
- `wordprocessing/testdata/` - Check for documents with track changes
- Create test documents with various revision types

## Success Criteria

- [ ] Can call `doc.GetRevisions()` and iterate all revisions
- [ ] Can call `revision.Accept()` on any revision type
- [ ] Can call `revision.Reject()` on any revision type
- [ ] Can call `doc.AcceptAllRevisions()` and all markup is removed
- [ ] Can call `doc.RejectAllRevisions()` and all changes are reverted
- [ ] Filter by author works correctly
- [ ] Moved content is handled properly (MoveFrom/MoveTo pairs)
- [ ] Format changes are accepted/rejected correctly
- [ ] Roundtrip test: document with revisions → accept some → save → open → verify

## Spec Capabilities to Update

- **wordprocessing-document**: Add revision management requirements
- **wordprocessing-elements**: Modify to note high-level API additions

## Dependencies

- No new dependencies
- Uses existing revision elements
- Builds on existing iterator patterns in goffice

## Out of Scope

- Real-time collaboration (track changes as they happen)
- Conflict resolution for simultaneous edits
- Version comparison (diff two documents)
- Revision history/timeline UI

## Example Usage (Target API)

```go
// Open document with track changes
doc, _ := wordprocessing.Open("document-with-changes.docx")

// Accept all revisions by specific author
for revision := range doc.GetRevisions() {
    if revision.Author() == "Jane Smith" {
        revision.Accept()
    }
}

// Reject all deletions
for revision := range doc.GetRevisions() {
    if revision.Type() == RevisionTypeDelete {
        revision.Reject()
    }
}

// Accept all remaining
doc.AcceptAllRevisions()

doc.Save()
```
