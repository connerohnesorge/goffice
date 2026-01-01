# Change: Word Comments API Completion

## Why

goffice has **existing Comment infrastructure** (~674 LOC) that provides basic comment creation but lacks high-level convenience APIs:

**Existing (DO NOT REBUILD):**
- `Comments` type with GetComment(id), AddComment(author, text), NextID()
- `Comment` type with Id, Author, Date, Initials, Paragraphs(), AppendParagraph()
- `CommentRangeStart`, `CommentRangeEnd`, `CommentReference` elements
- `CommentsPart` integration with Document
- `CommentStatus` enum exists (active, resolved, closed) but NOT exposed on Comment
- Basic tests and Comments.docx fixture

**Missing (BUILD THESE):**
1. High-level comment range marking API (apply comment to text range)
2. Comment status support (expose existing enum on Comment)
3. Comment deletion with cascade cleanup
4. Comment filtering by author, date, status
5. Comment threading/replies (parentID)

**Impact**: 7/10 - Comments are essential for document review workflows. Without high-level APIs, users must manually insert range markers and references.

## What Changes

### Phase 1: Core API Completion (40-50 hours)

1. **Comment Range Marking API** (15-20h)
   - Add `Paragraph.MarkCommentRange(startRunIdx, endRunIdx, commentID int)` - inserts CommentRangeStart/End
   - Add `Document.ApplyComment(para *Paragraph, startIdx, endIdx int, author, text string) (*Comment, error)`
   - Auto-insert CommentRangeStart, CommentRangeEnd, CommentReference in correct positions

2. **Comment Status Support** (8-10h)
   - Expose existing CommentStatus enum on Comment type
   - Add `Comment.Done()`, `SetDone(done bool)` methods for w:done attribute
   - Support resolved, closed, active states via extended attributes

3. **Comment Deletion** (10-12h)
   - Add `Document.RemoveComment(id int) error` - removes comment + all references
   - Add `CommentsPart.RemoveComment(id int) error`
   - Cascade delete CommentRangeStart, CommentRangeEnd, CommentReference from document body

4. **Comment Filtering** (8-10h)
   - Add `Comments.ByAuthor(author string) []*Comment`
   - Add `Comments.ByDateRange(from, to time.Time) []*Comment`
   - Add `Comments.Count() int`

### Phase 2: Advanced Features (25-35 hours)

5. **Comment Threading/Replies** (15-20h)
   - Support w:paraIdParent attribute for threaded replies
   - Add `Comment.ParentID()`, `SetParentID(id string)`
   - Add `Comment.Replies() []*Comment`
   - Add `Comments.AddReply(parentID int, author, text string) *Comment`

6. **Extended Comment Metadata** (10-15h)
   - Add CommentsExtendedPart (word/commentsExtended.xml)
   - Support for comment extended attributes (done, paraIdParent)

### Phase 3: Polish (15-20 hours)

7. **Examples & Tests** (10-12h)
   - Create examples/comments/main.go - comprehensive demo
   - Add roundtrip tests with complex comment scenarios
   - Test reading/modifying existing comments

8. **Documentation** (5-8h)
   - Godoc comments on all exported types/methods
   - Usage examples in doc comments

**Breaking changes**: None. This enhances existing code.

## Impact

- **Affected specs**: `wordprocessing-comments` (new delta spec)
- **New capabilities**: High-level comment APIs
- **Affected code**:
  - `wordprocessing/elements/comments.go` - ENHANCE: Add filtering, status methods
  - `wordprocessing/paragraph.go` - ENHANCE: Add MarkCommentRange()
  - `wordprocessing/document.go` - ENHANCE: Add ApplyComment(), RemoveComment()
  - `wordprocessing/parts/comments_part.go` - ENHANCE: Add RemoveComment()
  - `wordprocessing/parts/comments_extended_part.go` - NEW: Extended metadata

## Key Design Decisions

1. **Paragraph-level Marking**: Comment ranges are marked at paragraph level since Word comments span runs within paragraphs

2. **Cascade Deletion**: Removing a comment must cleanup all related elements (range markers, references) in document body

3. **Done Attribute for Status**: Use w:done attribute matching Office behavior, expose via Done()/SetDone()

4. **Filter Returns Slices**: Filtering methods return slices for flexibility (can be chained with Go code)

## Success Criteria

1. Existing 674 LOC unchanged, all current tests pass
2. `Document.ApplyComment(para, start, end, author, text)` works end-to-end
3. Comment done status persists correctly
4. Removing a comment cleans up all references
5. Roundtrip: create comment → save → open → read → verify

## Risk Assessment

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Range marking across paragraphs | Medium | Medium | Document single-paragraph limitation first |
| Comments extended part format | Low | Low | Test with real Word documents |
| Cascade deletion misses references | Medium | High | Exhaustive tree traversal |

## Specs Changed

- `spectr/changes/add-comments-api/specs/wordprocessing-comments/spec.md` (new delta)
