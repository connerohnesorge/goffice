//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// TableDefinitionPart represents a table definition part (xl/tables/table1.xml, etc.).
// This part defines an Excel table with columns, filters, and styles.
type TableDefinitionPart struct {
	*openxml.OpenXmlPartData
}

// newTableDefinitionPart creates a new table definition part.
func newTableDefinitionPart(
	worksheetPart *WorksheetPart,
	uri string,
) (*TableDefinitionPart, error) {
	packPart, relID, err := worksheetPart.addChildPart(
		uri,
		ContentTypeTable,
		RelationshipTypeTable,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTable,
		packPart,
		worksheetPart,
	)
	partData.SetRelationshipID(relID)

	tp := &TableDefinitionPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal table content
	tp.initializeContent()

	// Add to worksheet part's child parts
	if err := worksheetPart.AddPart(tp, relID); err != nil {
		return nil, err
	}

	return tp, nil
}

// initializeContent sets up minimal table content.
func (tp *TableDefinitionPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<table xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" id="1" name="Table1" displayName="Table1" ref="A1:A1" totalsRowShown="0">
  <autoFilter ref="A1:A1"/>
  <tableColumns count="1">
    <tableColumn id="1" name="Column1"/>
  </tableColumns>
  <tableStyleInfo name="TableStyleMedium2" showFirstColumn="0" showLastColumn="0" showRowStripes="1" showColumnStripes="0"/>
</table>`
	tp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*TableDefinitionPart) FixedContentType() string {
	return ContentTypeTable
}

// Table returns the root Table element.
// If the root element exists but is a generic PartRootElementBase (from parsing),
// it wraps it as a Table element for convenient access.
func (tp *TableDefinitionPart) Table() *elements.Table {
	root := tp.RootElement()
	if root == nil {
		return nil
	}
	if t, ok := root.(*elements.Table); ok {
		return t
	}
	// If the root is a PartRootElementBase that was loaded from XML,
	// wrap it in our Table type
	if base, ok := root.(*openxml.PartRootElementBase); ok {
		return &elements.Table{
			PartRootElementBase: base,
		}
	}

	return nil
}

// TableId returns the table ID from the root element.
func (tp *TableDefinitionPart) TableId() uint32 {
	t := tp.Table()
	if t == nil {
		return 0
	}

	return t.Id()
}

// TableName returns the table name.
func (tp *TableDefinitionPart) TableName() string {
	t := tp.Table()
	if t == nil {
		return ""
	}

	return t.Name()
}

// TableRef returns the table cell range reference.
func (tp *TableDefinitionPart) TableRef() string {
	t := tp.Table()
	if t == nil {
		return ""
	}

	return t.Ref()
}

// GetStream returns a reader for the part content.
func (tp *TableDefinitionPart) GetStream() io.Reader {
	return tp.OpenXmlPartData.GetStream()
}

// Ensure TableDefinitionPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*TableDefinitionPart)(
	nil,
)

// TableDefinitionPartFactory creates a TableDefinitionPart from a URI and container.
func TableDefinitionPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Use GetPackagingPart which is safe during initialization (no locks)
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTable,
		packPart,
		container,
	)

	return &TableDefinitionPart{
		OpenXmlPartData: partData,
	}
}

// Register the TableDefinitionPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeTable,
			RelationshipType:   RelationshipTypeTable,
			Factory:            TableDefinitionPartFactory,
			DefaultURI:         "/xl/tables/table1.xml",
			IsFixedContentType: true,
		},
	)
}
