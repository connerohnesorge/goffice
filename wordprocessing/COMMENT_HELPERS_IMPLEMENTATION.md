# Comment Range Marking API Implementation

## Summary

This implementation adds Phase 1.1 of the Comment Range Marking API for Word Comments in the goffice project.

## Files Created/Modified

### New Files

1. **wordprocessing/comment_helpers.go**
   - `ApplyComment(para, startIdx, endIdx, author, text)` - High-level API to create and apply comments
   - `GetOrCreateCommentsPart()` - Convenience method for getting/creating comments part

2. **wordprocessing/comment_helpers_test.go**
   - Comprehensive test suite covering all functionality
   - Tests for valid and invalid scenarios
   - Tests for rollback behavior
   - Tests for multiple comments on same paragraph

3. **examples/comment_example.go**
   - Example demonstrating how to use the comment API
   - Shows creation of comments on paragraph runs

### Modified Files

1. **wordprocessing/elements/paragraph.go**
   - Added `InsertChildBefore(newChild, refChild)` method
   - Added `InsertChildAfter(newChild, refChild)` method
   - Added `MarkCommentRange(startRunIdx, endRunIdx, commentID)` method
   - All methods include proper validation and error handling

## API Overview

### Paragraph Methods

```go
// Insert a child element before a reference child
func (p *Paragraph) InsertChildBefore(newChild, refChild openxml.Element) error

// Insert a child element after a reference child
func (p *Paragraph) InsertChildAfter(newChild, refChild openxml.Element) error

// Mark a range of runs with comment markers
func (p *Paragraph) MarkCommentRange(startRunIdx, endRunIdx, commentID int) error
```

### Document Methods

```go
// Apply a comment to a range of runs in a paragraph
func (d *Document) ApplyComment(
    para *elements.Paragraph,
    startIdx, endIdx int,
    author, text string,
) (*elements.Comment, error)
```

## Implementation Details

### MarkCommentRange

This method inserts the necessary XML elements for a comment range:

1. **CommentRangeStart** - Inserted before the start run
2. **CommentRangeEnd** - Inserted after the end run
3. **CommentReference** - Inserted inside the end run (at the end)

The method validates:
- Run indices are within bounds
- Start index <= end index
- Paragraph has at least one run

### ApplyComment

This high-level method:

1. Gets or creates the CommentsPart
2. Creates a new Comment with the specified author and text
3. Calls `MarkCommentRange` to insert the markers
4. Includes rollback logic if marking fails

### Error Handling

- All methods return descriptive errors
- Invalid indices are caught and reported
- Comment creation is rolled back if range marking fails

## Testing

All tests pass successfully:

```bash
go test ./wordprocessing/... -run TestComment -v
```

Test coverage includes:

- Valid range marking (single and multiple runs)
- Invalid indices (negative, out of bounds, start > end)
- Empty paragraphs (no runs)
- Comment creation and part management
- Multiple comments on same paragraph
- Rollback on failure
- InsertChildBefore/After validation

## Example Usage

```go
// Create a document
doc, _ := wordprocessing.New("example.docx", wordprocessing.DocTypeDocument)
defer doc.Close()

// Add a paragraph with runs
para := elements.NewParagraph()
para.AppendRun("First ")
para.AppendRun("Second ")
para.AppendRun("Third")
body.AppendChild(para)

// Apply a comment to runs 0-1
comment, err := doc.ApplyComment(
    para,
    0, 1,
    "John Doe",
    "This needs revision!",
)

// Save the document
doc.Save()
```

## Compliance

- All code follows goffice coding standards
- Linter passes with 0 issues
- All tests pass
- Proper error handling and validation
- Pre-allocation for performance
- Clear documentation and comments

## Next Steps

This implementation covers Phase 1.1. Future phases would include:

- Phase 1.2: Comment querying and retrieval
- Phase 1.3: Comment modification and deletion
- Phase 2: Reply support
- Phase 3: Resolved/unresolved tracking
