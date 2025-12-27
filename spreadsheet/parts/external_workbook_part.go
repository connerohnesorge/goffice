//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
)

// ExternalWorkbookPart represents an external workbook reference part.
// This part contains references to external workbooks used in formulas.
type ExternalWorkbookPart struct {
	*openxml.OpenXmlPartData
}

// Counter for generating unique external link filenames.
var externalLinkCounter uint64

// newExternalWorkbookPart creates a new external workbook part.
func newExternalWorkbookPart(
	workbookPart *WorkbookPart,
) (*ExternalWorkbookPart, error) {
	num := atomic.AddUint64(
		&externalLinkCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/xl/externalLinks/externalLink%d.xml",
		num,
	)

	packPart, relID, err := workbookPart.addChildPart(
		uri,
		ContentTypeExternalLink,
		RelationshipTypeExternalLink,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeExternalLink,
		packPart,
		workbookPart,
	)
	partData.SetRelationshipID(relID)

	ep := &ExternalWorkbookPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal external link content
	ep.initializeContent()

	// Add to workbook part's child parts
	if err := workbookPart.AddPart(ep, relID); err != nil {
		return nil, err
	}

	return ep, nil
}

// initializeContent sets up minimal external link content.
func (ep *ExternalWorkbookPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<externalLink xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
  <externalBook r:id="rId1">
    <sheetNames>
      <sheetName val="Sheet1"/>
    </sheetNames>
  </externalBook>
</externalLink>`
	ep.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*ExternalWorkbookPart) FixedContentType() string {
	return ContentTypeExternalLink
}

// ExternalLink returns the root ExternalLink element.
// TODO: Return a proper ExternalLink element type when elements are implemented.
func (ep *ExternalWorkbookPart) ExternalLink() openxml.PartRootElement {
	return ep.RootElement()
}

// GetStream returns a reader for the part content.
func (ep *ExternalWorkbookPart) GetStream() io.Reader {
	return ep.OpenXmlPartData.GetStream()
}

// Ensure ExternalWorkbookPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*ExternalWorkbookPart)(
	nil,
)

// ExternalWorkbookPartFactory creates an ExternalWorkbookPart from a URI and container.
func ExternalWorkbookPartFactory(
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
		ContentTypeExternalLink,
		packPart,
		container,
	)

	return &ExternalWorkbookPart{
		OpenXmlPartData: partData,
	}
}

// Register the ExternalWorkbookPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeExternalLink,
			RelationshipType:   RelationshipTypeExternalLink,
			Factory:            ExternalWorkbookPartFactory,
			DefaultURI:         "/xl/externalLinks/externalLink1.xml",
			IsFixedContentType: true,
		},
	)
}
