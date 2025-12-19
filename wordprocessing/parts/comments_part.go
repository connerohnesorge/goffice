package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
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
func (cp *CommentsPart) FixedContentType() string {
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
	pkg := container.Package()
	if pkg == nil {
		return nil
	}

	packPart, err := pkg.Part(uri)
	if err != nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeComments,
		packPart,
		container,
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
