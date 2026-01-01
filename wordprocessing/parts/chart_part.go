//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

const (
	// ContentTypeChart is the content type for chart parts.
	ContentTypeChart = "application/vnd.openxmlformats-officedocument.drawingml.chart+xml"
	// RelationshipTypeChart is the relationship type for chart parts.
	RelationshipTypeChart = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chart"
)

// ChartPart represents a chart part (word/charts/chart1.xml, etc.).
// This part contains the chart definition and data.
type ChartPart struct {
	*openxml.OpenXmlPartData
}

//nolint:unused // scaffolding for future chart creation API
var chartUniqueCounter uint64

// newChartPart creates a new chart part for a main document.
//
//nolint:unused // scaffolding for future chart creation API
func newChartPart(
	mainPart *MainPart,
) (*ChartPart, error) {
	num := atomic.AddUint64(
		&chartUniqueCounter,
		1,
	)
	uri := "/word/charts/chart" +
		itoaUint(num) +
		".xml"

	packPart, relID, err := mainPart.addChildPart(
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
		mainPart,
	)
	partData.SetRelationshipID(relID)

	cp := &ChartPart{
		OpenXmlPartData: partData,
	}
	cp.initializeContent()

	if err := mainPart.AddPart(cp, relID); err != nil {
		return nil, err
	}

	return cp, nil
}

//nolint:unused // called by newChartPart which is scaffolding for future API
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
	partData.SetRootFactory(func() openxml.PartRootElement {
		return drawingml.NewChartSpace()
	})

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
			DefaultURI:         "/word/charts/chart1.xml",
			IsFixedContentType: true,
		},
	)
}

//nolint:unused // used by newChartPart which is scaffolding for future API
func itoaUint(value uint64) string {
	if value == 0 {
		return "0"
	}

	buf := [20]byte{}
	i := len(buf)
	for value > 0 {
		i--
		buf[i] = byte('0' + value%10)
		value /= 10
	}

	return string(buf[i:])
}
