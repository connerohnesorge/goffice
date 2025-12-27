//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
)

// CustomXmlPart represents a custom XML part.
type CustomXmlPart struct {
	*openxml.OpenXmlPartData
}

// Content type and relationship type for custom XML.
const (
	ContentTypeCustomXml      = "application/xml"
	RelationshipTypeCustomXml = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/customXml"
)

// Counter for generating unique custom XML part filenames.
var customXmlCounter uint64

// newCustomXmlPart creates a new custom XML part.
func newCustomXmlPart(
	mainPart *MainPart,
) (*CustomXmlPart, error) {
	num := atomic.AddUint64(&customXmlCounter, 1)
	uri := fmt.Sprintf(
		"/customXml/item%d.xml",
		num,
	)

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeCustomXml,
		RelationshipTypeCustomXml,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeCustomXml,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	cp := &CustomXmlPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal custom XML content
	cp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(cp, relID); err != nil {
		return nil, err
	}

	return cp, nil
}

// initializeContent sets up minimal custom XML content.
func (cp *CustomXmlPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<root></root>`
	cp.SetData([]byte(content))
}

// SetXmlData sets the custom XML data.
func (cp *CustomXmlPart) SetXmlData(
	data []byte,
) {
	cp.SetData(data)
}

// GetXmlData returns the custom XML data.
func (cp *CustomXmlPart) GetXmlData() []byte {
	return cp.GetData()
}

// GetStream returns a reader for the part content.
func (cp *CustomXmlPart) GetStream() io.Reader {
	return cp.OpenXmlPartData.GetStream()
}

// Ensure CustomXmlPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*CustomXmlPart)(nil)

// CustomXmlPartFactory creates a CustomXmlPart from a URI and container.
func CustomXmlPartFactory(
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
		ContentTypeCustomXml,
		packPart,
		container,
	)

	return &CustomXmlPart{
		OpenXmlPartData: partData,
	}
}

// Register the CustomXmlPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeCustomXml,
			RelationshipType:   RelationshipTypeCustomXml,
			Factory:            CustomXmlPartFactory,
			DefaultURI:         "/customXml/item1.xml",
			IsFixedContentType: false, // Custom XML can have different content types
		},
	)
}
