//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// QueryTablePart represents a query table part (xl/queryTables/queryTable1.xml, etc.).
// This part defines external data queries for a worksheet.
type QueryTablePart struct {
	*openxml.OpenXmlPartData
}

// newQueryTablePart creates a new query table part.
func newQueryTablePart(
	worksheetPart *WorksheetPart,
	uri string,
) (*QueryTablePart, error) {
	packPart, relID, err := worksheetPart.addChildPart(
		uri,
		ContentTypeQueryTable,
		RelationshipTypeQueryTable,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeQueryTable,
		packPart,
		worksheetPart,
	)
	partData.SetRelationshipID(relID)

	qp := &QueryTablePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal query table content
	qp.initializeContent()

	// Add to worksheet part's child parts
	if err := worksheetPart.AddPart(qp, relID); err != nil {
		return nil, err
	}

	return qp, nil
}

// initializeContent sets up minimal query table content.
func (qp *QueryTablePart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<queryTable xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" name="Query1" connectionId="1" autoFormatId="16" applyNumberFormats="0" applyBorderFormats="0" applyFontFormats="0" applyPatternFormats="0" applyAlignmentFormats="0" applyWidthHeightFormats="0">
  <queryTableRefresh nextId="1">
    <queryTableFields count="0"/>
  </queryTableRefresh>
</queryTable>`
	qp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*QueryTablePart) FixedContentType() string {
	return ContentTypeQueryTable
}

// QueryTable returns the root QueryTable element.
// TODO: Return a proper QueryTable element type when elements are implemented.
func (qp *QueryTablePart) QueryTable() openxml.PartRootElement {
	return qp.RootElement()
}

// GetStream returns a reader for the part content.
func (qp *QueryTablePart) GetStream() io.Reader {
	return qp.OpenXmlPartData.GetStream()
}

// Ensure QueryTablePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*QueryTablePart)(nil)

// QueryTablePartFactory creates a QueryTablePart from a URI and container.
func QueryTablePartFactory(
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
		ContentTypeQueryTable,
		packPart,
		container,
	)

	return &QueryTablePart{
		OpenXmlPartData: partData,
	}
}

// Register the QueryTablePart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeQueryTable,
			RelationshipType:   RelationshipTypeQueryTable,
			Factory:            QueryTablePartFactory,
			DefaultURI:         "/xl/queryTables/queryTable1.xml",
			IsFixedContentType: true,
		},
	)
}
