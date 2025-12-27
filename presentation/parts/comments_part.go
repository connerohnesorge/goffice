//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
)

// Counter for generating unique comment filenames.
var commentUniqueCounter uint64

// CommentAuthorsPart represents the comment authors part (ppt/commentAuthors.xml).
type CommentAuthorsPart struct {
	*openxml.OpenXmlPartData
}

// newCommentAuthorsPart creates a new comment authors part.
func newCommentAuthorsPart(
	presentationPart *PresentationPart,
) (*CommentAuthorsPart, error) {
	uri := "/ppt/commentAuthors.xml"

	packPart, relID, err := presentationPart.addChildPart(
		uri,
		ContentTypeCommentAuthors,
		RelationshipTypeCommentAuthors,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeCommentAuthors,
		packPart,
		presentationPart,
	)
	partData.SetRelationshipID(relID)

	cap := &CommentAuthorsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal comment authors content
	cap.initializeContent()

	// Add to presentation part's child parts
	if err := presentationPart.AddPart(cap, relID); err != nil {
		return nil, err
	}

	return cap, nil
}

// initializeContent sets up minimal comment authors content.
func (cap *CommentAuthorsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:cmAuthorLst xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
</p:cmAuthorLst>`
	cap.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*CommentAuthorsPart) FixedContentType() string {
	return ContentTypeCommentAuthors
}

// CommentAuthors returns the root CommentAuthorList element.
// TODO: Return a proper CommentAuthorList element type when elements are implemented.
func (cap *CommentAuthorsPart) CommentAuthors() openxml.PartRootElement {
	return cap.RootElement()
}

// GetStream returns a reader for the part content.
func (cap *CommentAuthorsPart) GetStream() io.Reader {
	return cap.OpenXmlPartData.GetStream()
}

// Ensure CommentAuthorsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*CommentAuthorsPart)(
	nil,
)

// CommentAuthorsPartFactory creates a CommentAuthorsPart from a URI and container.
func CommentAuthorsPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeCommentAuthors,
		packPart,
		container,
	)

	return &CommentAuthorsPart{
		OpenXmlPartData: partData,
	}
}

// SlideCommentsPart represents a slide comments part (ppt/comments/comment1.xml, etc.).
type SlideCommentsPart struct {
	*openxml.OpenXmlPartData
}

// newSlideCommentsPart creates a new slide comments part.
func newSlideCommentsPart(
	slidePart *SlidePart,
) (*SlideCommentsPart, error) {
	num := atomic.AddUint64(
		&commentUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/comments/comment%d.xml",
		num,
	)

	packPart, relID, err := slidePart.addChildPart(
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
		slidePart,
	)
	partData.SetRelationshipID(relID)

	scp := &SlideCommentsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal slide comments content
	scp.initializeContent()

	// Add to slide part's child parts
	if err := slidePart.AddPart(scp, relID); err != nil {
		return nil, err
	}

	return scp, nil
}

// initializeContent sets up minimal slide comments content.
func (scp *SlideCommentsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:cmLst xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
</p:cmLst>`
	scp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*SlideCommentsPart) FixedContentType() string {
	return ContentTypeComments
}

// Comments returns the root CommentList element.
// TODO: Return a proper CommentList element type when elements are implemented.
func (scp *SlideCommentsPart) Comments() openxml.PartRootElement {
	return scp.RootElement()
}

// GetStream returns a reader for the part content.
func (scp *SlideCommentsPart) GetStream() io.Reader {
	return scp.OpenXmlPartData.GetStream()
}

// Ensure SlideCommentsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*SlideCommentsPart)(
	nil,
)

// SlideCommentsPartFactory creates a SlideCommentsPart from a URI and container.
func SlideCommentsPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
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

	return &SlideCommentsPart{
		OpenXmlPartData: partData,
	}
}

// Register the comment part types.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeCommentAuthors,
			RelationshipType:   RelationshipTypeCommentAuthors,
			Factory:            CommentAuthorsPartFactory,
			DefaultURI:         "/ppt/commentAuthors.xml",
			IsFixedContentType: true,
		},
	)

	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeComments,
			RelationshipType:   RelationshipTypeComments,
			Factory:            SlideCommentsPartFactory,
			DefaultURI:         "/ppt/comments/comment1.xml",
			IsFixedContentType: true,
		},
	)
}
