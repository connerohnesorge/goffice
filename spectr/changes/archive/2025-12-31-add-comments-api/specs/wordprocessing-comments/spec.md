# Delta Spec: Word Comments API Completion

**Status**: ENHANCEMENT to existing comments capability
**Base Spec**: `spectr/specs/wordprocessing-elements/spec.md` (provides Comment Range Elements)
**Delta**: Adds high-level API for comment range marking, status, deletion, and filtering

## Baseline Acknowledgment

goffice has **existing Comment infrastructure** (~674 LOC in `wordprocessing/elements/comments.go` and `wordprocessing/parts/comments_part.go`) that this spec **enhances**:
- `Comments` with GetComment(id), AddComment(author, text), NextID()
- `Comment` with Id, Author, Date, Initials, Paragraphs(), AppendParagraph()
- `CommentRangeStart`, `CommentRangeEnd`, `CommentReference` elements
- `CommentsPart` with AddComment(), GetComment()
- `CommentStatus` enum (active, resolved, closed) - defined but not exposed

## ADDED Requirements

### Requirement: Comment Range Marking

The system SHALL provide methods to mark text ranges with comments.

#### Scenario: Mark run range in paragraph
- GIVEN a paragraph with 3 runs
- WHEN `para.MarkCommentRange(0, 1, 42)` is called
- THEN CommentRangeStart with id=42 is inserted before run 0
- AND CommentRangeEnd with id=42 is inserted after run 1
- AND a Run with CommentReference id=42 is appended after CommentRangeEnd

#### Scenario: Mark single run
- GIVEN a paragraph with runs
- WHEN `para.MarkCommentRange(1, 1, 5)` is called
- THEN CommentRangeStart and CommentRangeEnd wrap only run at index 1
- AND CommentReference follows

#### Scenario: Invalid start index
- GIVEN a paragraph with 3 runs
- WHEN `para.MarkCommentRange(-1, 2, 1)` is called
- THEN error is returned indicating invalid start index

#### Scenario: Invalid end index
- GIVEN a paragraph with 3 runs
- WHEN `para.MarkCommentRange(0, 5, 1)` is called
- THEN error is returned indicating invalid end index

#### Scenario: Start greater than end
- GIVEN a paragraph with runs
- WHEN `para.MarkCommentRange(2, 0, 1)` is called
- THEN error is returned indicating start > end

### Requirement: Apply Comment to Range

The system SHALL provide a high-level method to create and apply comments.

#### Scenario: Apply comment creates all elements
- GIVEN a document with a paragraph containing runs
- WHEN `doc.ApplyComment(para, 0, 2, "John", "Review this")` is called
- THEN a Comment is created in the comments part with author="John"
- AND the paragraph is marked with the comment's ID
- AND the Comment is returned

#### Scenario: Apply comment creates comments part if needed
- GIVEN a document without a comments part
- WHEN `doc.ApplyComment(para, 0, 0, "Jane", "Note")` is called
- THEN a comments part is created
- AND the comment is added to it

#### Scenario: Apply comment rollback on failure
- GIVEN a document and an invalid paragraph
- WHEN `doc.ApplyComment(para, -1, 0, "John", "text")` fails
- THEN no comment is left in the comments part
- AND error is returned

### Requirement: Comment Status

The system SHALL support marking comments as done/resolved.

#### Scenario: Get done status default
- GIVEN a Comment without w:done attribute
- WHEN `comment.Done()` is called
- THEN false is returned

#### Scenario: Set done to true
- GIVEN a Comment
- WHEN `comment.SetDone(true)` is called
- THEN `comment.Done()` returns true
- AND the w:done attribute is set in w15 namespace

#### Scenario: Set done to false
- GIVEN a Comment with done=true
- WHEN `comment.SetDone(false)` is called
- THEN `comment.Done()` returns false
- AND the w:done attribute is removed

#### Scenario: Done status persists
- GIVEN a document with a comment marked done
- WHEN saved and reopened
- THEN `comment.Done()` returns true

### Requirement: Comment Deletion

The system SHALL support removing comments with cascade cleanup.

#### Scenario: Remove comment from collection
- GIVEN a Comments collection with comment id=5
- WHEN `comments.RemoveComment(5)` is called
- THEN true is returned
- AND `comments.GetComment(5)` returns nil

#### Scenario: Remove non-existent comment
- GIVEN a Comments collection without comment id=99
- WHEN `comments.RemoveComment(99)` is called
- THEN false is returned

#### Scenario: Document RemoveComment cascade
- GIVEN a document with comment id=3 and corresponding range markers
- WHEN `doc.RemoveComment(3)` is called
- THEN the comment is removed from comments part
- AND all CommentRangeStart elements with id=3 are removed from body
- AND all CommentRangeEnd elements with id=3 are removed from body
- AND all CommentReference elements with id=3 are removed from runs

#### Scenario: Cascade handles nested elements
- GIVEN a document with comment markers inside a table cell
- WHEN `doc.RemoveComment(id)` is called
- THEN markers inside the table are also removed

#### Scenario: Remove non-existent comment returns error
- GIVEN a document without comment id=99
- WHEN `doc.RemoveComment(99)` is called
- THEN error is returned indicating comment not found

### Requirement: Comment Filtering

The system SHALL provide methods to filter comments.

#### Scenario: Filter by author
- GIVEN comments from authors "John", "Jane", "John"
- WHEN `comments.ByAuthor("John")` is called
- THEN slice of 2 comments is returned
- AND all returned comments have author "John"

#### Scenario: Filter by author no matches
- GIVEN comments without author "Bob"
- WHEN `comments.ByAuthor("Bob")` is called
- THEN empty slice is returned

#### Scenario: Filter by date range
- GIVEN comments dated Jan 1, Jan 15, Feb 1
- WHEN `comments.ByDateRange(jan1, jan31)` is called
- THEN slice of 2 comments is returned
- AND all are within January

#### Scenario: Count comments
- GIVEN a Comments element with 5 comment children
- WHEN `comments.Count()` is called
- THEN 5 is returned

#### Scenario: Count empty
- GIVEN a Comments element with no children
- WHEN `comments.Count()` is called
- THEN 0 is returned

### Requirement: CommentsPart Delegation

The system SHALL provide comment removal through CommentsPart.

#### Scenario: Remove via part
- GIVEN a CommentsPart with comment id=7
- WHEN `part.RemoveComment(7)` is called
- THEN the comment is removed from the underlying Comments element

## Implementation Notes

### File Structure (Enhancement to existing)
```
wordprocessing/elements/
└── comments.go       // ENHANCE: Add Done(), filtering, RemoveComment()

wordprocessing/
├── paragraph.go      // ENHANCE: Add MarkCommentRange()
├── document.go       // ENHANCE: Add ApplyComment(), RemoveComment()
└── comment_helpers.go // NEW: Cascade deletion helper (~50 LOC)

wordprocessing/parts/
└── comments_part.go  // ENHANCE: Add RemoveComment() delegation
```

### API Surface Summary (Additions Only)

```go
// ADDED to Comment
func (c *Comment) Done() bool
func (c *Comment) SetDone(done bool)

// ADDED to Comments
func (c *Comments) ByAuthor(author string) []*Comment
func (c *Comments) ByDateRange(from, to time.Time) []*Comment
func (c *Comments) Count() int
func (c *Comments) RemoveComment(id int) bool

// ADDED to Paragraph
func (p *Paragraph) MarkCommentRange(startRunIdx, endRunIdx, commentID int) error

// ADDED to Document
func (d *Document) ApplyComment(para *Paragraph, startRunIdx, endRunIdx int, author, text string) (*elements.Comment, error)
func (d *Document) RemoveComment(commentID int) error

// ADDED to CommentsPart
func (cp *CommentsPart) RemoveComment(id int) bool
```

## Backward Compatibility

**Unchanged APIs** (fully backward compatible):
- All existing Comment methods unchanged
- All existing Comments methods unchanged
- All existing CommentRangeStart/End/Reference methods unchanged
- All CommentsPart methods unchanged

**New APIs** (additive only):
- Done()/SetDone() on Comment
- ByAuthor(), ByDateRange(), Count(), RemoveComment() on Comments
- MarkCommentRange() on Paragraph
- ApplyComment(), RemoveComment() on Document
- RemoveComment() on CommentsPart

**No Breaking Changes**: Existing code continues to work unchanged.

## Phase 2+ Enhancements (NOT IN THIS CHANGE)

Deferred to future proposals:
- Comment threading/replies (paraIdParent)
- Comments extended part (word/commentsExtended.xml)
- Cross-paragraph comment ranges
- Comment author metadata

## Testing Requirements

**Unit Tests (20+ cases)**: Per design.md section
- Comment range marking: 6 tests
- Comment status: 4 tests
- Comment deletion: 6 tests
- Comment filtering: 4 tests

**Integration Tests (4 scenarios)**: Per design.md section
- Roundtrip with comment creation
- Remove and verify clean document
- Filter and verify results
- Read existing Word document with comments

See `tasks.md` for detailed test plan.
