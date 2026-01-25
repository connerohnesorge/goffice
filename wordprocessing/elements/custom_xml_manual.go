package elements

import (
	"encoding/xml"
)

// DatastoreItem represents the properties of a custom XML part.
type DatastoreItem struct {
	XMLName    xml.Name             `xml:"http://schemas.openxmlformats.org/officeDocument/2006/customXml datastoreItem"`
	ItemID     string               `xml:"itemID,attr"`
	SchemaRefs *DatastoreSchemaRefs `xml:"schemaRefs,omitempty"`
}

// NewDatastoreItem creates a new DatastoreItem.
func NewDatastoreItem() *DatastoreItem {
	return &DatastoreItem{}
}

// DatastoreSchemaRefs represents a collection of schema references.
type DatastoreSchemaRefs struct {
	SchemaRef []*DatastoreSchemaRef `xml:"schemaRef"`
}

// DatastoreSchemaRef represents a reference to an XML schema.
type DatastoreSchemaRef struct {
	URI string `xml:"uri,attr"`
}

// AddSchemaRef adds a schema reference.
func (d *DatastoreItem) AddSchemaRef(uri string) {
	if d.SchemaRefs == nil {
		d.SchemaRefs = &DatastoreSchemaRefs{}
	}
	d.SchemaRefs.SchemaRef = append(d.SchemaRefs.SchemaRef, &DatastoreSchemaRef{URI: uri})
}
