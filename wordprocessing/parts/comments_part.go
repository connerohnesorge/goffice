//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"time"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
	"github.com/connerohnesorge/goffice/wordprocessing/validation"
)

// CommentsPart represents the comments part (word/comments.xml).
type CommentsPart struct {
	*openxml.OpenXmlPartData
}

// Content type and relationship type for comments.
const (
	ContentTypeComments      = "application/vnd.openxmlformats-officedocument.wordprocessingml.comments+xml"
	RelationshipTypeComments = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"
)

// newCommentsPart creates a new comments part.
func newCommentsPart(
	mainPart *MainPart,
) (*CommentsPart, error) {
	uri := "/word/comments.xml"

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeComments,
		RelationshipTypeComments,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeComments,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	cp := &CommentsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal comments content
	cp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(cp, relID); err != nil {
		return nil, err
	}

	return cp, nil
}

// initializeContent sets up minimal comments content.
func (cp *CommentsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:comments xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
</w:comments>`
	cp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
func (*CommentsPart) FixedContentType() string {
	return ContentTypeComments
}

// Comments returns the root Comments element.
func (cp *CommentsPart) Comments() *elements.Comments {
	root := cp.RootElement()
	if root == nil {
		return nil
	}
	if c, ok := root.(*elements.Comments); ok {
		return c
	}
	// Wrap the root element as Comments
	if pre, ok := root.(*openxml.PartRootElementBase); ok {
		return &elements.Comments{
			PartRootElementBase: pre,
		}
	}

	return nil
}

// GetOrCreateComments returns the Comments element, creating if necessary.
func (cp *CommentsPart) GetOrCreateComments() *elements.Comments {
	c := cp.Comments()
	if c != nil {
		return c
	}
	c = elements.NewComments()
	cp.SetRootElement(c)

	return c
}

// GetCommentThreads returns all comments organized into threads.
func (cp *CommentsPart) GetCommentThreads() []*elements.CommentThread {
	comments := cp.GetOrCreateComments()
	if comments == nil {
		return []*elements.CommentThread{}
	}

	// Build comment hierarchy map
	commentMap := make(map[int]*elements.Comment)
	for comment := range comments.Comments() {
		commentMap[comment.Id()] = comment
	}

	// Create thread structures
	var threads []*elements.CommentThread
	for comment := range comments.Comments() {
		if comment.ParentId() == 0 {
			// Root comment - start new thread
			threads = append(threads, elements.NewCommentThread(comment))
		} else {
			// Reply comment - add to parent's thread
			if parent, exists := commentMap[comment.ParentId()]; exists {
				// Find existing thread for parent and add reply
				for _, thread := range threads {
					if thread.Root.Id() == parent.Id() {
						thread.AddReply(comment)
						break
					}
				}
			} else {
				// Create new thread with parent as root
				threads = append(threads, elements.NewCommentThread(parent))
			}
		}
	}

	return threads
}

// GetCommentsByAuthor returns all comments by a specific author.
func (cp *CommentsPart) GetCommentsByAuthor(author string) []*elements.Comment {
	comments := cp.GetOrCreateComments()
	return comments.ByAuthor(author)
}

// GetCommentsByDateRange returns comments within a date range.
func (cp *CommentsPart) GetCommentsByDateRange(from, to time.Time) []*elements.Comment {
	comments := cp.GetOrCreateComments()
	return comments.ByDateRange(from, to)
}

// GetCommentsByStatus returns comments by their resolved status.
func (cp *CommentsPart) GetCommentsByStatus(resolved bool) []*elements.Comment {
	var result []*elements.Comment
	comments := cp.GetOrCreateComments()
	for comment := range comments.Comments() {
		if comment.Done() == resolved {
			result = append(result, comment)
		}
	}
	return result
}

// GetUnresolvedComments returns all comments that are not marked as done.
func (cp *CommentsPart) GetUnresolvedComments() []*elements.Comment {
	return cp.GetCommentsByStatus(false)
}

// GetResolvedComments returns all comments that are marked as done.
func (cp *CommentsPart) GetResolvedComments() []*elements.Comment {
	return cp.GetCommentsByStatus(true)
}

// AddCommentWithParent adds a comment with optional parent ID.
func (cp *CommentsPart) AddCommentWithParent(author, text string, parentId int) (*elements.Comment, error) {
	comment := cp.AddComment(author, text)
	if parentId > 0 {
		comment.SetParentId(parentId)
	}
	return comment, nil
}

// AddReply adds a reply to an existing comment.
func (cp *CommentsPart) AddReply(parentComment *elements.Comment, author, text string) (*elements.Comment, error) {
	return cp.AddCommentWithParent(author, text, parentComment.Id())
}

// GetCommentHierarchy returns comments organized by thread hierarchy.
func (cp *CommentsPart) GetCommentHierarchy() map[int][]*elements.Comment {
	hierarchy := make(map[int][]*elements.Comment)
	comments := cp.GetOrCreateComments()

	for comment := range comments.Comments() {
		parentId := comment.ParentId()
		if parentId > 0 {
			hierarchy[parentId] = append(hierarchy[parentId], comment)
		} else {
			hierarchy[0] = append(hierarchy[0], comment)
		}
	}

	return hierarchy
}

// GetThreadComments returns comments for a specific thread (root + all replies).
func (cp *CommentsPart) GetThreadComments(rootId int) []*elements.Comment {
	var threadComments []*elements.Comment
	comments := cp.GetOrCreateComments()

	for comment := range comments.Comments() {
		if comment.Id() == rootId || comment.ParentId() == rootId {
			threadComments = append(threadComments, comment)
		}
	}

	return threadComments
}

// ValidateCommentThread validates a complete comment thread structure.
func (cp *CommentsPart) ValidateCommentThread(thread *elements.CommentThread) error {
	// Use validation package we created
	return validation.ValidateCommentThread(thread)
}

// MarkThreadResolved marks an entire thread as resolved.
func (cp *CommentsPart) MarkThreadResolved(rootId int) error {
	comments := cp.GetOrCreateComments()

	// Find root comment
	var rootComment *elements.Comment
	for comment := range comments.Comments() {
		if comment.Id() == rootId {
			rootComment = comment
			break
		}
	}

	if rootComment == nil {
		return fmt.Errorf("root comment with ID %d not found", rootId)
	}

	// Mark thread as resolved
	thread := elements.NewCommentThread(rootComment)
	thread.MarkResolved()

	return nil
}

// AddCommentValidation adds validation results to comment metadata.
func (cp *CommentsPart) AddCommentValidation(comment *elements.Comment) error {
	if validationErr := validation.ValidateComment(comment); validationErr != nil {
		// Store validation error as extended attribute for debugging
		comment.SetAttribute(openxml.NewAttribute(
			"",
			"validation",
			"",
			validationErr.Error(),
		))
		return validationErr
	}

	return nil
}

// GetCommentWithValidation returns comment with validation status.
func (cp *CommentsPart) GetCommentWithValidation(id int) (*elements.Comment, error) {
	comment := cp.GetComment(id)
	if comment == nil {
		return nil, fmt.Errorf("comment with ID %d not found", id)
	}

	// Validate and add validation metadata
	if validationErr := cp.AddCommentValidation(comment); validationErr != nil {
		return comment, validationErr
	}

	return comment, nil
}

// AddComment adds a new comment and returns it.
func (cp *CommentsPart) AddComment(
	author, text string,
) *elements.Comment {
	c := cp.GetOrCreateComments()

	return c.AddComment(author, text)
}

// GetComment returns the comment with the specified ID.
func (cp *CommentsPart) GetComment(
	id int,
) *elements.Comment {
	c := cp.Comments()
	if c == nil {
		return nil
	}

	return c.GetComment(id)
}

// RemoveComment removes the comment with the specified ID from the collection.
// Returns true if the comment was found and removed, false if not found.
func (cp *CommentsPart) RemoveComment(
	id int,
) bool {
	c := cp.Comments()
	if c == nil {
		return false
	}

	return c.RemoveComment(id)
}

// GetStream returns a reader for the part content.
func (cp *CommentsPart) GetStream() io.Reader {
	return cp.OpenXmlPartData.GetStream()
}

// Ensure CommentsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*CommentsPart)(nil)

// CommentsPartFactory creates a CommentsPart from a URI and container.
func CommentsPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Use GetPackagingPart instead of Package() to avoid locking issues
	// during initialization when the package lock is already held
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeComments,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewComments()
		},
	)

	return &CommentsPart{
		OpenXmlPartData: partData,
	}
}

// Register the CommentsPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeComments,
			RelationshipType:   RelationshipTypeComments,
			Factory:            CommentsPartFactory,
			DefaultURI:         "/word/comments.xml",
			IsFixedContentType: true,
		},
	)
}
