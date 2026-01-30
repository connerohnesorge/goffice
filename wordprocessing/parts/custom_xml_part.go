//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"encoding/xml"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
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

// Marshal serializes the given value to XML and stores it in the part.
func (cp *CustomXmlPart) Marshal(v interface{}) error {
	data, err := xml.Marshal(v)
	if err != nil {
		return err
	}
	// Add XML header
	content := []byte(xml.Header)
	content = append(content, data...)
	cp.SetData(content)

	return nil
}

// Unmarshal deserializes the part data into the given value.
func (cp *CustomXmlPart) Unmarshal(v interface{}) error {
	data := cp.GetData()

	return xml.Unmarshal(data, v)
}

// AddPropertiesPart adds a custom XML properties part to this custom XML part.
func (cp *CustomXmlPart) AddPropertiesPart() (*CustomXmlPropertiesPart, error) {
	// Generate URI for the properties part
	uri := cp.URI() + "Props.xml"

	// Create the child part
	packPart, relID, err := cp.addChildPart(
		uri,
		ContentTypeCustomXmlProperties,
		RelationshipTypeCustomXmlProperties,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeCustomXmlProperties,
		packPart,
		cp,
	)
	partData.SetRelationshipID(relID)

	propsPart := &CustomXmlPropertiesPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal properties content
	propsPart.initializeContent()

	// Add to this part's child parts
	if err := cp.AddPart(propsPart, relID); err != nil {
		return nil, err
	}

	return propsPart, nil
}

// addChildPart is a helper to add a child part.
func (cp *CustomXmlPart) addChildPart(
	uri string,
	contentType string,
	relationshipType string,
) (*packaging.Part, string, error) {
	pkg := cp.Package()
	if pkg == nil {
		return nil, "", fmt.Errorf("package is nil")
	}

	packPart, err := pkg.CreatePart(uri, contentType)
	if err != nil {
		return nil, "", err
	}

	rel, err := pkg.CreatePartRelationship(
		cp.URI(),
		uri,
		relationshipType,
		"",
	)
	if err != nil {
		return nil, "", err
	}

	return packPart, rel.ID(), nil
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
