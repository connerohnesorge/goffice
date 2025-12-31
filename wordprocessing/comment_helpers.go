package wordprocessing

import (
	"fmt"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
	"github.com/connerohnesorge/goffice/wordprocessing/parts"
)

// ApplyComment creates a new comment and marks the specified range of runs
// in the paragraph with comment markers.
//
// Parameters:
//   - para: The paragraph to apply the comment to
//   - startIdx: 0-based index of the first run in the comment range
//   - endIdx: 0-based index of the last run in the comment range
//   - author: The author of the comment
//   - text: The comment text content
//
// Returns the created Comment element and an error if the operation fails.
// If marking the range fails, the comment is rolled back (removed from the part).
//
//nolint:revive // argument-limit: public API requires all parameters
func (d *Document) ApplyComment(
	para *elements.Paragraph,
	startIdx, endIdx int,
	author, text string,
) (*elements.Comment, error) {
	// Get or create the comments part
	commentsPart, err := d.CommentsPart()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get/create comments part: %w",
			err,
		)
	}

	// Get or create the Comments element
	comments := commentsPart.GetOrCreateComments()

	// Create the comment
	comment := comments.AddComment(author, text)
	commentID := comment.Id()

	// Mark the range in the paragraph
	if err := para.MarkCommentRange(
		startIdx,
		endIdx,
		commentID,
	); err != nil {
		// Rollback: remove the comment from the comments part
		comments.RemoveChild(comment)

		return nil, fmt.Errorf(
			"failed to mark comment range: %w",
			err,
		)
	}

	return comment, nil
}

// GetOrCreateCommentsPart returns the comments part, creating it if it doesn't exist.
// This is a convenience method that wraps CommentsPart() to always return a valid part.
func (d *Document) GetOrCreateCommentsPart() (*parts.CommentsPart, error) {
	return d.CommentsPart()
}

// RemoveComment removes a comment by ID from the document.
// This includes removing:
// - The comment from the CommentsPart
// - All CommentRangeStart markers with matching ID
// - All CommentRangeEnd markers with matching ID
// - All CommentReference elements with matching ID
//
// Returns an error if the comment is not found.
func (d *Document) RemoveComment(
	commentID int,
) error {
	// Get the comments part
	commentsPart, err := d.CommentsPart()
	if err != nil {
		return fmt.Errorf(
			"failed to get comments part: %w",
			err,
		)
	}
	if commentsPart == nil {
		return fmt.Errorf(
			"comment %d not found: no comments part exists",
			commentID,
		)
	}

	// Remove the comment from the comments part
	if !commentsPart.RemoveComment(commentID) {
		return fmt.Errorf(
			"comment %d not found in comments part",
			commentID,
		)
	}

	// Remove all comment markers from the document body
	mainPart := d.MainPart()
	doc := mainPart.Document()
	body := doc.Body()
	if body != nil {
		removeCommentMarkers(body, commentID)
	}

	return nil
}

// removeCommentMarkers recursively removes all comment markers (CommentRangeStart,
// CommentRangeEnd, and CommentReference) with the specified comment ID from the
// given element and its descendants.
func removeCommentMarkers(
	parent openxml.Element,
	commentID int,
) {
	// Only process composite elements
	composite, ok := parent.(openxml.CompositeElement)
	if !ok {
		return
	}

	toRemove := make([]openxml.Element, 0)

	for child := range composite.Children() {
		// Check for comment markers that match the ID
		switch elem := child.(type) {
		case *elements.CommentRangeStart:
			if elem.Id() == commentID {
				toRemove = append(toRemove, child)
			}
		case *elements.CommentRangeEnd:
			if elem.Id() == commentID {
				toRemove = append(toRemove, child)
			}
		case *elements.Run:
			// Check inside runs for CommentReference
			removeCommentReferenceFromRun(elem, commentID)
		}

		// Recurse into composite elements (paragraphs, tables, SDTs, etc.)
		removeCommentMarkers(child, commentID)
	}

	// Remove all collected elements
	for _, elem := range toRemove {
		composite.RemoveChild(elem)
	}
}

// removeCommentReferenceFromRun removes CommentReference elements with the
// specified comment ID from a Run.
func removeCommentReferenceFromRun(
	run *elements.Run,
	commentID int,
) {
	toRemove := make([]openxml.Element, 0)

	for child := range run.Children() {
		ref, ok := child.(*elements.CommentReference)
		if !ok {
			continue
		}
		if ref.Id() == commentID {
			toRemove = append(toRemove, child)
		}
	}

	// Remove all collected comment references
	for _, elem := range toRemove {
		run.RemoveChild(elem)
	}
}
