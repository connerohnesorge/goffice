//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
)

// ChartsheetPart represents a chartsheet part (xl/chartsheets/sheet1.xml, etc.).
// A chartsheet is a special type of sheet that displays a chart full-screen.
type ChartsheetPart struct {
	*openxml.OpenXmlPartData
}

// newChartsheetPart creates a new chartsheet part.
func newChartsheetPart(
	workbookPart *WorkbookPart,
	uri string,
) (*ChartsheetPart, error) {
	packPart, relID, err := workbookPart.addChildPart(
		uri,
		ContentTypeChartsheet,
		RelationshipTypeChartsheet,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeChartsheet,
		packPart,
		workbookPart,
	)
	partData.SetRelationshipID(relID)

	csp := &ChartsheetPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal chartsheet content
	csp.initializeContent()

	// Add to workbook part's child parts
	if err := workbookPart.AddPart(csp, relID); err != nil {
		return nil, err
	}

	return csp, nil
}

// initializeContent sets up minimal chartsheet content.
func (csp *ChartsheetPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<chartsheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
  <sheetViews>
    <sheetView workbookViewId="0"/>
  </sheetViews>
</chartsheet>`
	csp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*ChartsheetPart) FixedContentType() string {
	return ContentTypeChartsheet
}

// Chartsheet returns the root Chartsheet element.
// TODO: Return a proper Chartsheet element type when elements are implemented.
func (csp *ChartsheetPart) Chartsheet() openxml.PartRootElement {
	return csp.RootElement()
}

// AddDrawingsPart adds a drawings part to this chartsheet.
func (csp *ChartsheetPart) AddDrawingsPart() (*DrawingsPart, error) {
	// Chartsheets use a different path for drawings
	uri := "/xl/drawings/drawing1.xml"

	return newDrawingsPartWithContainer(csp, uri)
}

// DrawingsPart returns the drawings part if present.
func (csp *ChartsheetPart) DrawingsPart() *DrawingsPart {
	for part := range csp.Parts() {
		if dp, ok := part.(*DrawingsPart); ok {
			return dp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (csp *ChartsheetPart) GetStream() io.Reader {
	return csp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (csp *ChartsheetPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		csp,
		uri,
		contentType,
		relType,
	)
}

// Ensure ChartsheetPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*ChartsheetPart)(nil)

// ChartsheetPartFactory creates a ChartsheetPart from a URI and container.
func ChartsheetPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeChartsheet,
		packPart,
		container,
	)

	return &ChartsheetPart{
		OpenXmlPartData: partData,
	}
}

// Register the ChartsheetPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeChartsheet,
			RelationshipType:   RelationshipTypeChartsheet,
			Factory:            ChartsheetPartFactory,
			DefaultURI:         "/xl/chartsheets/sheet1.xml",
			IsFixedContentType: true,
		},
	)
}
