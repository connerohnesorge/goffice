//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// WorksheetCommentsPart represents a worksheet comments part (xl/comments1.xml, etc.).
// This part contains cell comments for a worksheet.
type WorksheetCommentsPart struct {
	*openxml.OpenXmlPartData
}

// newWorksheetCommentsPart creates a new worksheet comments part.
func newWorksheetCommentsPart(
	worksheetPart *WorksheetPart,
	uri string,
) (*WorksheetCommentsPart, error) {
	packPart, relID, err := worksheetPart.addChildPart(
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
		worksheetPart,
	)
	partData.SetRelationshipID(relID)

	cp := &WorksheetCommentsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal comments content
	cp.initializeContent()

	// Add to worksheet part's child parts
	if err := worksheetPart.AddPart(cp, relID); err != nil {
		return nil, err
	}

	return cp, nil
}

// initializeContent sets up minimal comments content.
func (cp *WorksheetCommentsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<comments xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main">
  <authors>
    <author>Author</author>
  </authors>
  <commentList>
  </commentList>
</comments>`
	cp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*WorksheetCommentsPart) FixedContentType() string {
	return ContentTypeComments
}

// Comments returns the root Comments element.
func (cp *WorksheetCommentsPart) Comments() *elements.Comments {
	if root := cp.RootElement(); root != nil {
		if c, ok := root.(*elements.Comments); ok {
			return c
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (cp *WorksheetCommentsPart) GetStream() io.Reader {
	return cp.OpenXmlPartData.GetStream()
}

// Ensure WorksheetCommentsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*WorksheetCommentsPart)(
	nil,
)

// WorksheetCommentsPartFactory creates a WorksheetCommentsPart from a URI and container.
func WorksheetCommentsPartFactory(
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
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewComments()
		},
	)

	return &WorksheetCommentsPart{
		OpenXmlPartData: partData,
	}
}

// Register the WorksheetCommentsPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeComments,
			RelationshipType:   RelationshipTypeComments,
			Factory:            WorksheetCommentsPartFactory,
			DefaultURI:         "/xl/comments1.xml",
			IsFixedContentType: true,
		},
	)
}
