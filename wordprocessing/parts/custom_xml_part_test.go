package parts_test

import (
	"encoding/xml"
	"testing"

	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

type TestData struct {
	XMLName xml.Name `xml:"root"`
	Value   string   `xml:"value"`
}

func TestCustomXml(t *testing.T) {
	doc, err := wordprocessing.New("test.docx", wordprocessing.DocTypeDocument)
	if err != nil {
		t.Fatalf("failed to create document: %s", err)
	}
	defer func() { _ = doc.Close() }()

	// Add custom XML part
	cxp, err := doc.MainPart().AddCustomXmlPart()
	if err != nil {
		t.Fatalf("failed to add custom xml part: %s", err)
	}

	// Set data
	data := TestData{Value: "Hello Custom XML"}
	if err := cxp.Marshal(&data); err != nil {
		t.Fatalf("failed to marshal data: %s", err)
	}

	// Verify data
	var readData TestData
	if err := cxp.Unmarshal(&readData); err != nil {
		t.Fatalf("failed to unmarshal data: %s", err)
	}

	if readData.Value != "Hello Custom XML" {
		t.Errorf("expected 'Hello Custom XML', got '%s'", readData.Value)
	}

	// Add properties part
	propsPart, err := cxp.AddPropertiesPart()
	if err != nil {
		t.Fatalf("failed to add properties part: %s", err)
	}

	// Set properties
	props := elements.NewDatastoreItem()
	props.ItemID = "{12345678-1234-1234-1234-123456789012}"
	props.AddSchemaRef("http://example.com/schema")

	if err := propsPart.SetProperties(props); err != nil {
		t.Fatalf("failed to set properties: %s", err)
	}

	// Verify properties
	readProps := propsPart.Properties()
	if readProps == nil {
		t.Fatal("failed to read properties")
	}

	if readProps.ItemID != props.ItemID {
		t.Errorf("expected itemID '%s', got '%s'", props.ItemID, readProps.ItemID)
	}

	if len(readProps.SchemaRefs.SchemaRef) != 1 {
		t.Errorf("expected 1 schema ref, got %d", len(readProps.SchemaRefs.SchemaRef))
	}

	if readProps.SchemaRefs.SchemaRef[0].URI != "http://example.com/schema" {
		t.Errorf("expected uri 'http://example.com/schema', got '%s'", readProps.SchemaRefs.SchemaRef[0].URI)
	}
}
