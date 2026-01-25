// This file implements the extended chart part (Office 2016+).
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// ContentTypeChartEx is the content type for extended charts.
const ContentTypeChartEx = "application/vnd.ms-office.chartex+xml"

// RelationshipTypeChartEx is the relationship type for extended charts.
const RelationshipTypeChartEx = "http://schemas.microsoft.com/office/2014/relationships/chartEx"

// ChartExPart represents an extended chart part (xl/charts/chartEx1.xml).
// This part contains extended chart definitions (Waterfall, Sunburst, etc.).
type ChartExPart struct {
	*openxml.OpenXmlPartData
}

// newChartExPart creates a new extended chart part.
func newChartExPart(
	drawingsPart *DrawingsPart,
	uri string,
) (*ChartExPart, error) {
	packPart, relID, err := drawingsPart.addChildPart(
		uri,
		ContentTypeChartEx,
		RelationshipTypeChartEx,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeChartEx,
		packPart,
		drawingsPart,
	)
	partData.SetRelationshipID(relID)

	cp := &ChartExPart{
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

// initializeContent sets up minimal extended chart content.
func (cp *ChartExPart) initializeContent() {
	chartSpace := drawingml.NewChartSpaceEx()
	cp.SetRootElement(chartSpace)
}

// FixedContentType returns the content type for this part.
func (*ChartExPart) FixedContentType() string {
	return ContentTypeChartEx
}

// ChartSpace returns the root ChartSpaceEx element.
func (cp *ChartExPart) ChartSpace() *drawingml.ChartSpaceEx {
	root := cp.RootElement()
	if root == nil {
		return nil
	}
	if chartSpace, ok := root.(*drawingml.ChartSpaceEx); ok {
		return chartSpace
	}

	return nil
}

// GetStream returns a reader for the part content.
func (cp *ChartExPart) GetStream() io.Reader {
	return cp.OpenXmlPartData.GetStream()
}

// Ensure ChartExPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*ChartExPart)(nil)

// ChartExPartFactory creates a ChartExPart from a URI and container.
func ChartExPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeChartEx,
		packPart,
		container,
	)
	partData.SetRootFactory(func() openxml.PartRootElement {
		return drawingml.NewChartSpaceEx()
	})

	return &ChartExPart{
		OpenXmlPartData: partData,
	}
}

// Register the ChartExPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeChartEx,
			RelationshipType:   RelationshipTypeChartEx,
			Factory:            ChartExPartFactory,
			DefaultURI:         "/xl/charts/chartEx1.xml",
			IsFixedContentType: true,
		},
	)
}
