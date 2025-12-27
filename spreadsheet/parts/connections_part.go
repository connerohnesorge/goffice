//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// ConnectionsPart represents the connections part (xl/connections.xml).
// This part defines external data connections for the workbook.
type ConnectionsPart struct {
	*openxml.OpenXmlPartData
}

// newConnectionsPart creates a new connections part.
func newConnectionsPart(
	workbookPart *WorkbookPart,
) (*ConnectionsPart, error) {
	uri := "/xl/connections.xml"

	packPart, relID, err := workbookPart.addChildPart(
		uri,
		ContentTypeConnections,
		RelationshipTypeConnections,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeConnections,
		packPart,
		workbookPart,
	)
	partData.SetRelationshipID(relID)

	cp := &ConnectionsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal connections content
	cp.initializeContent()

	// Add to workbook part's child parts
	if err := workbookPart.AddPart(cp, relID); err != nil {
		return nil, err
	}

	return cp, nil
}

// initializeContent sets up minimal connections content.
func (cp *ConnectionsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<connections xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main">
</connections>`
	cp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*ConnectionsPart) FixedContentType() string {
	return ContentTypeConnections
}

// Connections returns the root Connections element.
func (cp *ConnectionsPart) Connections() *elements.Connections {
	if root := cp.RootElement(); root != nil {
		if c, ok := root.(*elements.Connections); ok {
			return c
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (cp *ConnectionsPart) GetStream() io.Reader {
	return cp.OpenXmlPartData.GetStream()
}

// Ensure ConnectionsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*ConnectionsPart)(
	nil,
)

// ConnectionsPartFactory creates a ConnectionsPart from a URI and container.
func ConnectionsPartFactory(
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
		ContentTypeConnections,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewConnections()
		},
	)

	return &ConnectionsPart{
		OpenXmlPartData: partData,
	}
}

// Register the ConnectionsPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeConnections,
			RelationshipType:   RelationshipTypeConnections,
			Factory:            ConnectionsPartFactory,
			DefaultURI:         "/xl/connections.xml",
			IsFixedContentType: true,
		},
	)
}
