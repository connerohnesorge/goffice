# Design: Word Comments API Completion

## Context

goffice has **existing Comment infrastructure** (~674 LOC) in `wordprocessing/elements/comments.go` and `wordprocessing/parts/comments_part.go` that provides:

**Existing (DO NOT REBUILD):**
- `Comments` type (root element for comments part)
- `Comment` type with Id, Author, Date, Initials, Paragraphs(), AppendParagraph()
- `CommentRangeStart`, `CommentRangeEnd`, `CommentReference` elements
- `CommentsPart` with AddComment(), GetComment()
- `CommentStatus` enum in enums.go (active, resolved, closed) - NOT exposed on Comment

**Missing (BUILD THESE):**
1. High-level comment range marking API
2. Comment status/done support
3. Comment deletion with cascade
4. Comment filtering
5. Comment threading/replies

The solution must **enhance existing code**, not replace it.

## Goals

1. **Enable comment range marking** - Apply comments to text spans
2. **Support comment status** - Expose done/resolved state
3. **Enable comment deletion** - Clean removal with cascade
4. **Support filtering** - By author, date range
5. **Maintain compatibility** - All existing tests pass unchanged

## Non-Goals

- Replace existing Comment/Comments types
- Full comments extended XML support (Phase 2)
- Cross-paragraph comment ranges (complex edge case)
- Real-time collaboration features

## Architecture Overview

### Existing Code (DO NOT MODIFY core structure)

```
wordprocessing/elements/
├── comments.go          // Comments, Comment, CommentRangeStart, etc. (500 LOC)
└── enums.go             // CommentStatus enum

wordprocessing/parts/
└── comments_part.go     // CommentsPart (174 LOC)
```

### New/Enhanced Code

```
wordprocessing/elements/
└── comments.go          // ENHANCE: Add filtering, Done() methods

wordprocessing/
├── paragraph.go         // ENHANCE: Add MarkCommentRange()
├── document.go          // ENHANCE: Add ApplyComment(), RemoveComment()
└── comment_helpers.go   // NEW: Helper functions for cascade deletion (~150 LOC)

wordprocessing/parts/
├── comments_part.go     // ENHANCE: Add RemoveComment()
└── comments_extended_part.go // NEW: Phase 2 (threading support)
```

## Detailed Design Decisions

### Decision 1: Comment Range Marking API

**Design**: Add methods at Paragraph and Document levels:

```go
// wordprocessing/paragraph.go

// MarkCommentRange inserts comment range markers around runs.
// startRunIdx and endRunIdx are 0-based indices into the paragraph's runs.
// Returns error if indices are invalid.
func (p *Paragraph) MarkCommentRange(startRunIdx, endRunIdx, commentID int) error {
    runs := p.Runs() // Get all runs

    // Validate indices
    runSlice := slices.Collect(runs)
    if startRunIdx < 0 || endRunIdx >= len(runSlice) || startRunIdx > endRunIdx {
        return fmt.Errorf("invalid run indices: start=%d, end=%d, total=%d",
            startRunIdx, endRunIdx, len(runSlice))
    }

    // Insert CommentRangeStart before startRunIdx
    startMarker := elements.NewCommentRangeStart(commentID)
    p.InsertBefore(startMarker, runSlice[startRunIdx])

    // Insert CommentRangeEnd after endRunIdx
    endMarker := elements.NewCommentRangeEnd(commentID)
    p.InsertAfter(endMarker, runSlice[endRunIdx])

    // Append CommentReference run after end marker
    refRun := elements.NewRun("")
    refRun.AppendChild(elements.NewCommentReference(commentID))
    p.InsertAfter(refRun, endMarker)

    return nil
}
```

```go
// wordprocessing/document.go

// ApplyComment creates a comment and marks the specified run range.
// Returns the created comment or error if operation fails.
func (d *Document) ApplyComment(
    para *Paragraph,
    startRunIdx, endRunIdx int,
    author, text string,
) (*elements.Comment, error) {
    // Get or create comments part
    cp, err := d.CommentsPart()
    if err != nil {
        mp := d.MainDocumentPart()
        cp, err = mp.AddCommentsPart()
        if err != nil {
            return nil, fmt.Errorf("failed to create comments part: %w", err)
        }
    }

    // Create the comment
    comment := cp.AddComment(author, text)
    commentID := comment.Id()

    // Mark the range in the paragraph
    if err := para.MarkCommentRange(startRunIdx, endRunIdx, commentID); err != nil {
        // Rollback: remove the comment we just created
        cp.GetOrCreateComments().RemoveComment(commentID)
        return nil, fmt.Errorf("failed to mark comment range: %w", err)
    }

    return comment, nil
}
```

**Rationale**:
- Paragraph-level marking matches Word's model (comments span runs within paragraphs)
- Document-level ApplyComment is the convenient high-level API
- Automatic rollback on failure prevents orphaned comments

### Decision 2: Comment Status Support

**Design**: Use w:done attribute (Office standard) and expose on Comment:

```go
// wordprocessing/elements/comments.go - ADD to Comment type

// Done returns whether the comment is marked as done/resolved.
// This corresponds to the w:done attribute.
func (c *Comment) Done() bool {
    attr, found := c.GetAttribute("done", NamespaceW15)
    if !found {
        return false
    }
    return attr.Value() == "1" || strings.ToLower(attr.Value()) == "true"
}

// SetDone marks the comment as done/resolved.
func (c *Comment) SetDone(done bool) {
    if done {
        c.SetAttribute(openxml.NewAttribute(
            NamespaceW15, "done", PrefixW15, "1",
        ))
    } else {
        c.RemoveAttribute("done", NamespaceW15)
    }
}
```

**Rationale**:
- Office uses w:done in w15 namespace for resolved comments
- Boolean is simpler than status enum for most use cases
- The existing CommentStatus enum can be used for display purposes

### Decision 3: Comment Deletion with Cascade

**Design**: Remove comment AND all related elements from document:

```go
// wordprocessing/document.go

// RemoveComment removes a comment and all its range markers from the document.
func (d *Document) RemoveComment(commentID int) error {
    // 1. Remove from comments part
    cp, err := d.CommentsPart()
    if err != nil {
        return fmt.Errorf("no comments part: %w", err)
    }

    comments := cp.GetOrCreateComments()
    if comments.GetComment(commentID) == nil {
        return fmt.Errorf("comment %d not found", commentID)
    }
    comments.removeCommentByID(commentID)

    // 2. Remove range markers from document body
    body := d.MainDocumentPart().Document().Body()
    removeCommentMarkers(body, commentID)

    return nil
}

// removeCommentMarkers recursively removes all markers for a comment ID.
func removeCommentMarkers(parent openxml.Element, commentID int) {
    var toRemove []openxml.Element

    for child := range parent.Children() {
        // Check if this is a comment marker with matching ID
        switch elem := child.(type) {
        case *elements.CommentRangeStart:
            if elem.Id() == commentID {
                toRemove = append(toRemove, elem)
            }
        case *elements.CommentRangeEnd:
            if elem.Id() == commentID {
                toRemove = append(toRemove, elem)
            }
        case *elements.CommentReference:
            if elem.Id() == commentID {
                toRemove = append(toRemove, elem)
            }
        }

        // Recurse into children
        if composite, ok := child.(openxml.CompositeElement); ok {
            removeCommentMarkers(composite, commentID)
        }
    }

    // Remove collected elements
    for _, elem := range toRemove {
        parent.RemoveChild(elem)
    }
}
```

**Rationale**:
- Cascade deletion prevents orphaned markers
- Recursive traversal handles nested structures (tables, SDTs)
- Collect-then-remove pattern avoids mutation during iteration

### Decision 4: Comment Filtering

**Design**: Add filter methods to Comments type:

```go
// wordprocessing/elements/comments.go - ADD to Comments type

// ByAuthor returns all comments by the specified author.
func (c *Comments) ByAuthor(author string) []*Comment {
    var result []*Comment
    for comment := range c.Comments() {
        if comment.Author() == author {
            result = append(result, comment)
        }
    }
    return result
}

// ByDateRange returns comments within the specified date range (inclusive).
func (c *Comments) ByDateRange(from, to time.Time) []*Comment {
    var result []*Comment
    for comment := range c.Comments() {
        date := comment.Date()
        if !date.Before(from) && !date.After(to) {
            result = append(result, comment)
        }
    }
    return result
}

// Count returns the total number of comments.
func (c *Comments) Count() int {
    count := 0
    for range c.Comments() {
        count++
    }
    return count
}
```

**Rationale**:
- Simple slice return for maximum flexibility
- Users can chain with Go code: `c.ByAuthor("John")[0]`
- Matches patterns in track changes proposal

### Decision 5: Comment Removal from Comments Collection

**Design**: Add removal method to Comments type:

```go
// wordprocessing/elements/comments.go - ADD to Comments type

// RemoveComment removes the comment with the specified ID.
// Returns true if a comment was removed, false if not found.
func (c *Comments) RemoveComment(id int) bool {
    for child := range c.Children() {
        if child.LocalName() == "comment" && child.NamespaceURI() == NamespaceWML {
            var comment *Comment
            switch v := child.(type) {
            case *Comment:
                comment = v
            case *openxml.CompositeElementBase:
                comment = &Comment{CompositeElementBase: v}
            }
            if comment != nil && comment.Id() == id {
                c.RemoveChild(child)
                return true
            }
        }
    }
    return false
}
```

**Rationale**:
- Needed for cascade deletion
- Returns boolean for caller to check success
- Matches pattern of other removal methods

## Module Organization

```
wordprocessing/
├── elements/
│   └── comments.go       // ENHANCE: Add filtering, Done(), RemoveComment()
├── parts/
│   └── comments_part.go  // ENHANCE: Add RemoveComment() delegation
├── paragraph.go          // ENHANCE: Add MarkCommentRange()
├── document.go           // ENHANCE: Add ApplyComment(), RemoveComment()
└── comment_helpers.go    // NEW: removeCommentMarkers() helper (~50 LOC)
```

**Total new code**: ~250 LOC
**Total enhanced code**: ~100 LOC modifications

## Testing Strategy

### Unit Tests (Per-Feature)

**Comment Range Marking** (6 tests):
- [ ] MarkCommentRange with valid indices
- [ ] MarkCommentRange with invalid start index returns error
- [ ] MarkCommentRange with invalid end index returns error
- [ ] MarkCommentRange with start > end returns error
- [ ] ApplyComment creates comment and marks range
- [ ] ApplyComment rollback on marking failure

**Comment Status** (4 tests):
- [ ] Done() returns false by default
- [ ] SetDone(true) sets w:done attribute
- [ ] SetDone(false) removes attribute
- [ ] Done() persists through save/load cycle

**Comment Deletion** (6 tests):
- [ ] RemoveComment removes from comments part
- [ ] RemoveComment removes CommentRangeStart from body
- [ ] RemoveComment removes CommentRangeEnd from body
- [ ] RemoveComment removes CommentReference from runs
- [ ] RemoveComment handles nested elements (tables)
- [ ] RemoveComment returns error for non-existent comment

**Comment Filtering** (4 tests):
- [ ] ByAuthor returns matching comments
- [ ] ByAuthor returns empty for no matches
- [ ] ByDateRange returns comments in range
- [ ] Count returns correct total

### Integration Tests (4 scenarios)

- [ ] Create comment → save → open → verify markers present
- [ ] Apply comment → remove comment → save → verify clean
- [ ] Multiple comments by different authors → filter → verify
- [ ] Read existing Word doc with comments → enumerate → verify

## API Surface Summary

### New Methods on Comment

```go
func (c *Comment) Done() bool
func (c *Comment) SetDone(done bool)
```

### New Methods on Comments

```go
func (c *Comments) ByAuthor(author string) []*Comment
func (c *Comments) ByDateRange(from, to time.Time) []*Comment
func (c *Comments) Count() int
func (c *Comments) RemoveComment(id int) bool
```

### New Methods on Paragraph

```go
func (p *Paragraph) MarkCommentRange(startRunIdx, endRunIdx, commentID int) error
```

### New Methods on Document

```go
func (d *Document) ApplyComment(para *Paragraph, startRunIdx, endRunIdx int, author, text string) (*elements.Comment, error)
func (d *Document) RemoveComment(commentID int) error
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

**No Breaking Changes**: Existing code continues to work unchanged.

## Success Metrics

1. **Backward compatible**: All existing tests pass unchanged
2. **Range marking works**: ApplyComment creates valid comment structure
3. **Status works**: Done() persists through save/load cycle
4. **Deletion works**: RemoveComment cleans up all references
5. **Filtering works**: ByAuthor/ByDateRange return correct results
6. **Performance**: 1000 comments filtered <100ms

## Phase 2 Extensions (NOT IN THIS CHANGE)

- Comment threading/replies (paraIdParent)
- Comments extended part (word/commentsExtended.xml)
- Cross-paragraph comment ranges
- Comment author metadata
