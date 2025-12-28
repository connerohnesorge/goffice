//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// ChartPart represents a chart part (ppt/charts/chart1.xml, etc.).
// This part contains the chart definition and data.
type ChartPart struct {
	*openxml.OpenXmlPartData
}

// Counter for generating unique chart filenames.
var chartUniqueCounter uint64

// newChartPartForSlide creates a new chart part for a slide.
func newChartPartForSlide(
	slidePart *SlidePart,
) (*ChartPart, error) {
	num := atomic.AddUint64(
		&chartUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/charts/chart%d.xml",
		num,
	)

	packPart, relID, err := slidePart.addChildPart(
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
		slidePart,
	)
	partData.SetRelationshipID(relID)

	cp := &ChartPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal chart content
	cp.initializeContent()

	// Add to slide part's child parts
	if err := slidePart.AddPart(cp, relID); err != nil {
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
			DefaultURI:         "/ppt/charts/chart1.xml",
			IsFixedContentType: true,
		},
	)
}
