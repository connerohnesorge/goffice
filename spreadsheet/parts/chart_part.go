//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// ChartPart represents a chart part (xl/charts/chart1.xml, etc.).
// This part contains the chart definition and data.
type ChartPart struct {
	*openxml.OpenXmlPartData
}

// newChartPart creates a new chart part.
func newChartPart(
	drawingsPart *DrawingsPart,
	uri string,
) (*ChartPart, error) {
	packPart, relID, err := drawingsPart.addChildPart(
		uri,
		ContentTypeChart,
		RelationshipTypeChart,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeChart,
		packPart,
		drawingsPart,
	)
	partData.SetRelationshipID(relID)

	cp := &ChartPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal chart content
	cp.initializeContent()

	// Add to drawings part's child parts
	if err := drawingsPart.AddPart(cp, relID); err != nil {
		return nil, err
	}

	return cp, nil
}

// initializeContent sets up minimal chart content.
func (cp *ChartPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<c:chartSpace xmlns:c="http://schemas.openxmlformats.org/drawingml/2006/chart" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
  <c:date1904 val="0"/>
  <c:lang val="en-US"/>
  <c:roundedCorners val="0"/>
  <c:chart>
    <c:autoTitleDeleted val="0"/>
    <c:plotArea>
      <c:layout/>
    </c:plotArea>
    <c:plotVisOnly val="1"/>
    <c:dispBlanksAs val="gap"/>
    <c:showDLblsOverMax val="0"/>
  </c:chart>
  <c:printSettings>
    <c:headerFooter/>
    <c:pageMargins b="0.75" l="0.7" r="0.7" t="0.75" header="0.3" footer="0.3"/>
    <c:pageSetup/>
  </c:printSettings>
</c:chartSpace>`
	cp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*ChartPart) FixedContentType() string {
	return ContentTypeChart
}

// ChartSpace returns the root ChartSpace element.
// TODO: Return a proper ChartSpace element type when elements are implemented.
func (cp *ChartPart) ChartSpace() openxml.PartRootElement {
	return cp.RootElement()
}

// GetStream returns a reader for the part content.
func (cp *ChartPart) GetStream() io.Reader {
	return cp.OpenXmlPartData.GetStream()
}

// Ensure ChartPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*ChartPart)(nil)

// ChartPartFactory creates a ChartPart from a URI and container.
func ChartPartFactory(
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
		ContentTypeChart,
		packPart,
		container,
	)

	return &ChartPart{
		OpenXmlPartData: partData,
	}
}

// Register the ChartPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeChart,
			RelationshipType:   RelationshipTypeChart,
			Factory:            ChartPartFactory,
			DefaultURI:         "/xl/charts/chart1.xml",
			IsFixedContentType: true,
		},
	)
}
