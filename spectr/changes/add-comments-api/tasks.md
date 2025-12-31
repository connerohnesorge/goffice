# Tasks: Word Comments API Completion

**Baseline**: ~674 LOC exists in `wordprocessing/elements/comments.go` and `wordprocessing/parts/comments_part.go`. This proposal **enhances** existing code.

## 1. Phase 1: Core API Completion (40-50 hours total)

### 1.1 Comment Range Marking API (15-20h)

- [ ] 1.1.1 Add `Paragraph.Runs()` method if not exists (returns iterator over runs)
- [ ] 1.1.2 Add `Paragraph.InsertBefore(newChild, refChild)` method if not exists
- [ ] 1.1.3 Add `Paragraph.InsertAfter(newChild, refChild)` method if not exists
- [ ] 1.1.4 Implement `Paragraph.MarkCommentRange(startRunIdx, endRunIdx, commentID int) error`
- [ ] 1.1.5 Add validation for run indices in MarkCommentRange
- [ ] 1.1.6 Implement `Document.ApplyComment(para, startIdx, endIdx, author, text) (*Comment, error)`
- [ ] 1.1.7 Add rollback logic in ApplyComment if marking fails
- [ ] 1.1.8 Write 6 unit tests for range marking

### 1.2 Comment Status Support (8-10h)

- [ ] 1.2.1 Add `NamespaceW15` constant if not exists (word 2012 namespace)
- [ ] 1.2.2 Implement `Comment.Done() bool` - reads w:done attribute
- [ ] 1.2.3 Implement `Comment.SetDone(done bool)` - sets/removes w:done attribute
- [ ] 1.2.4 Add namespace declaration for w15 in comments root if using Done
- [ ] 1.2.5 Write 4 unit tests for status support

### 1.3 Comment Deletion (10-12h)

- [ ] 1.3.1 Implement `Comments.RemoveComment(id int) bool`
- [ ] 1.3.2 Create `comment_helpers.go` with cascade deletion helper
- [ ] 1.3.3 Implement `removeCommentMarkers(parent, commentID)` recursive helper
- [ ] 1.3.4 Handle CommentRangeStart removal
- [ ] 1.3.5 Handle CommentRangeEnd removal
- [ ] 1.3.6 Handle CommentReference removal (within runs)
- [ ] 1.3.7 Handle nested elements (tables, SDTs)
- [ ] 1.3.8 Implement `Document.RemoveComment(commentID int) error`
- [ ] 1.3.9 Add `CommentsPart.RemoveComment(id int) bool` delegation
- [ ] 1.3.10 Write 6 unit tests for deletion

### 1.4 Comment Filtering (8-10h)

- [x] 1.4.1 Implement `Comments.ByAuthor(author string) []*Comment`
- [x] 1.4.2 Implement `Comments.ByDateRange(from, to time.Time) []*Comment`
- [x] 1.4.3 Implement `Comments.Count() int`
- [x] 1.4.4 Write 4 unit tests for filtering

## 2. Phase 2: Advanced Features (25-35 hours total)

### 2.1 Comment Threading/Replies (15-20h)

- [ ] 2.1.1 Research w:paraIdParent attribute format
- [ ] 2.1.2 Add `Comment.ParaIdParent() string` method
- [ ] 2.1.3 Add `Comment.SetParaIdParent(id string)` method
- [ ] 2.1.4 Implement `Comments.Replies(parentID int) []*Comment`
- [ ] 2.1.5 Implement `Comments.AddReply(parentID int, author, text string) *Comment`
- [ ] 2.1.6 Handle paraId generation for new comments
- [ ] 2.1.7 Write 6 unit tests for threading

### 2.2 Extended Comment Metadata (10-15h)

- [ ] 2.2.1 Research comments extended part format (word/commentsExtended.xml)
- [ ] 2.2.2 Create `CommentsExtendedPart` type
- [ ] 2.2.3 Define content type and relationship constants
- [ ] 2.2.4 Implement part factory and registration
- [ ] 2.2.5 Add `MainPart.CommentsExtendedPart()` accessor
- [ ] 2.2.6 Write 4 unit tests for extended part

## 3. Phase 3: Polish (15-20 hours total)

### 3.1 Examples (6-8h)

- [ ] 3.1.1 Create `examples/comments/main.go`
- [ ] 3.1.2 Demonstrate ApplyComment to text range
- [ ] 3.1.3 Demonstrate RemoveComment with cascade
- [ ] 3.1.4 Demonstrate comment filtering
- [ ] 3.1.5 Verify example compiles and runs

### 3.2 Integration Tests (6-8h)

- [ ] 3.2.1 Create test fixture with existing comments
- [ ] 3.2.2 Write roundtrip test: create → save → open → verify
- [ ] 3.2.3 Write test for applying comment → removing → verifying clean
- [ ] 3.2.4 Write test reading real Word document with comments
- [ ] 3.2.5 Verify all existing tests still pass

### 3.3 Documentation (4-6h)

- [ ] 3.3.1 Add godoc comments to all new exported types
- [ ] 3.3.2 Add godoc comments to all new exported methods
- [ ] 3.3.3 Add usage examples in doc comments
- [ ] 3.3.4 Document limitations (single-paragraph ranges)

## Summary

| Phase | Hours | Key Deliverables |
|-------|-------|------------------|
| Phase 1 | 40-50h | Range marking, status, deletion, filtering |
| Phase 2 | 25-35h | Threading/replies, extended metadata |
| Phase 3 | 15-20h | Examples, integration tests, documentation |
| **Total** | **80-105h** | **Complete Comments API** |

## Dependencies

- **No external dependencies**: Uses existing generated elements
- **Internal dependencies**:
  - 1.2 can run parallel with 1.1
  - 1.3 depends on 1.1 partially (uses same paragraph manipulation)
  - 1.4 can run parallel with 1.1-1.3
  - Phase 2 depends on Phase 1 completion
  - Phase 3 depends on Phase 1 (Phase 2 not required)

## Parallelization

- Tasks 1.1 and 1.2 can run in parallel
- Task 1.4 (filtering) can run in parallel with 1.1-1.3
- Phase 2 tasks (2.1, 2.2) can run in parallel
- Phase 3 examples can start after Phase 1 without Phase 2
