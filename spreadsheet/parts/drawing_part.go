//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
)

// DrawingsPart represents a drawings part (xl/drawings/drawing1.xml, etc.).
// This part contains the positioning and reference information for drawings,
// charts, and images in a worksheet.
type DrawingsPart struct {
	*openxml.OpenXmlPartData
}

// newDrawingsPart creates a new drawings part for a worksheet.
func newDrawingsPart(
	worksheetPart *WorksheetPart,
	uri string,
) (*DrawingsPart, error) {
	packPart, relID, err := worksheetPart.addChildPart(
		uri,
		ContentTypeDrawing,
		RelationshipTypeDrawing,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeDrawing,
		packPart,
		worksheetPart,
	)
	partData.SetRelationshipID(relID)

	dp := &DrawingsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal drawing content
	dp.initializeContent()

	// Add to worksheet part's child parts
	if err := worksheetPart.AddPart(dp, relID); err != nil {
		return nil, err
	}

	return dp, nil
}

// newDrawingsPartWithContainer creates a new drawings part for a generic container.
// This is used by ChartsheetPart and other containers that need drawings.
func newDrawingsPartWithContainer(
	container interface {
		addChildPart(uri, contentType, relType string) (*packaging.Part, string, error)
		openxml.OpenXmlPartContainer
	},
	uri string,
) (*DrawingsPart, error) {
	packPart, relID, err := container.addChildPart(
		uri,
		ContentTypeDrawing,
		RelationshipTypeDrawing,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeDrawing,
		packPart,
		container,
	)
	partData.SetRelationshipID(relID)

	dp := &DrawingsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal drawing content
	dp.initializeContent()

	// Add to container's child parts
	if err := container.AddPart(dp, relID); err != nil {
		return nil, err
	}

	return dp, nil
}

// initializeContent sets up minimal drawing content.
func (dp *DrawingsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<xdr:wsDr xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
</xdr:wsDr>`
	dp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*DrawingsPart) FixedContentType() string {
	return ContentTypeDrawing
}

// Drawing returns the root WorksheetDrawing element.
// TODO: Return a proper WorksheetDrawing element type when elements are implemented.
func (dp *DrawingsPart) Drawing() openxml.PartRootElement {
	return dp.RootElement()
}

// Counter for generating unique chart filenames.
var chartCounter uint64

// AddChartPart adds a chart part to this drawings part.
func (dp *DrawingsPart) AddChartPart() (*ChartPart, error) {
	num := atomic.AddUint64(&chartCounter, 1)
	uri := fmt.Sprintf(
		"/xl/charts/chart%d.xml",
		num,
	)

	return newChartPart(dp, uri)
}

// ChartParts returns all chart parts.
func (dp *DrawingsPart) ChartParts() []*ChartPart {
	var charts []*ChartPart
	for part := range dp.Parts() {
		if cp, ok := part.(*ChartPart); ok {
			charts = append(charts, cp)
		}
	}

	return charts
}

// Counter for generating unique image filenames.
var imageCounter uint64

// AddImagePart adds an image part to this drawings part.
func (dp *DrawingsPart) AddImagePart(
	imageType ImageType,
) (*ImagePart, error) {
	num := atomic.AddUint64(&imageCounter, 1)
	uri := fmt.Sprintf(
		"/xl/media/image%d%s",
		num,
		imageType.Extension(),
	)

	return newImagePart(dp, uri, imageType)
}

// ImageParts returns all image parts.
func (dp *DrawingsPart) ImageParts() []*ImagePart {
	var images []*ImagePart
	for part := range dp.Parts() {
		if ip, ok := part.(*ImagePart); ok {
			images = append(images, ip)
		}
	}

	return images
}

// GetStream returns a reader for the part content.
func (dp *DrawingsPart) GetStream() io.Reader {
	return dp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (dp *DrawingsPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		dp,
		uri,
		contentType,
		relType,
	)
}

// Ensure DrawingsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*DrawingsPart)(nil)

// DrawingsPartFactory creates a DrawingsPart from a URI and container.
func DrawingsPartFactory(
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
		ContentTypeDrawing,
		packPart,
		container,
	)

	return &DrawingsPart{
		OpenXmlPartData: partData,
	}
}

// Register the DrawingsPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeDrawing,
			RelationshipType:   RelationshipTypeDrawing,
			Factory:            DrawingsPartFactory,
			DefaultURI:         "/xl/drawings/drawing1.xml",
			IsFixedContentType: true,
		},
	)
}
