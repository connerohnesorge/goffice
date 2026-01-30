package parts

import (
	"encoding/xml"
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// CustomXmlPropertiesPart represents a custom XML properties part.
type CustomXmlPropertiesPart struct {
	*openxml.OpenXmlPartData
}

// Content type and relationship type for custom XML properties.
const (
	ContentTypeCustomXmlProperties      = "application/vnd.openxmlformats-officedocument.customXmlProperties+xml"
	RelationshipTypeCustomXmlProperties = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/customXmlProps"
)

// GetStream returns a reader for the part content.
func (cp *CustomXmlPropertiesPart) GetStream() io.Reader {
	return cp.OpenXmlPartData.GetStream()
}

// initializeContent sets up minimal custom XML properties content.
func (cp *CustomXmlPropertiesPart) initializeContent() {
	props := elements.NewDatastoreItem()
	props.ItemID = "{00000000-0000-0000-0000-000000000000}"
	_ = cp.SetProperties(props)
}

// Properties returns the properties element.
func (cp *CustomXmlPropertiesPart) Properties() *elements.DatastoreItem {
	if cp.OpenXmlPartData == nil {
		return nil
	}

	var props elements.DatastoreItem
	if err := cp.unmarshal(&props); err != nil {
		return nil
	}

	return &props
}

// SetProperties sets the properties element.
func (cp *CustomXmlPropertiesPart) SetProperties(props *elements.DatastoreItem) error {
	data, err := xml.Marshal(props)
	if err != nil {
		return err
	}
	// Add XML header
	content := []byte(xml.Header)
	content = append(content, data...)
	cp.SetData(content)

	return nil
}

// unmarshal deserializes the part data into the given value.
func (cp *CustomXmlPropertiesPart) unmarshal(v interface{}) error {
	data := cp.GetData()

	return xml.Unmarshal(data, v)
}

// Ensure CustomXmlPropertiesPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*CustomXmlPropertiesPart)(nil)

// CustomXmlPropertiesPartFactory creates a CustomXmlPropertiesPart from a URI and container.
func CustomXmlPropertiesPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeCustomXmlProperties,
		packPart,
		container,
	)

	return &CustomXmlPropertiesPart{
		OpenXmlPartData: partData,
	}
}

// Register the CustomXmlPropertiesPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeCustomXmlProperties,
			RelationshipType:   RelationshipTypeCustomXmlProperties,
			Factory:            CustomXmlPropertiesPartFactory,
			DefaultURI:         "/customXml/itemProps1.xml",
			IsFixedContentType: true,
		},
	)
}
