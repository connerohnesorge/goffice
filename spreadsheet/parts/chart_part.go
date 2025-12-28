//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/drawingml"
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
	chartSpace := drawingml.NewChartSpace()
	cp.SetRootElement(chartSpace)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*ChartPart) FixedContentType() string {
	return ContentTypeChart
}

// ChartSpace returns the root ChartSpace element.
func (cp *ChartPart) ChartSpace() *drawingml.ChartSpace {
	root := cp.RootElement()
	if root == nil {
		return nil
	}
	if chartSpace, ok := root.(*drawingml.ChartSpace); ok {
		return chartSpace
	}

	return nil
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
	// Use GetPackagingPart which is safe during initialization (no locks)
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
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
