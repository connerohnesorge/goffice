package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Common PresentationML constants.
const (
	NamespacePresentationML   = openxml.NamespacePresentationML
	NamespaceDrawingML        = openxml.NamespaceDrawingML
	NamespaceDrawingMLPicture = openxml.NamespaceDrawingMLPicture
	NamespaceDrawingMLChart   = openxml.NamespaceDrawingMLChart
	NamespaceRelationships    = openxml.NamespaceRelationships
	PrefixP                   = "p"
	PrefixA                   = "a"
	PrefixPic                 = "pic"
	PrefixC                   = "c"
	PrefixR                   = "r"
)

// Local names for shape tree elements.
const (
	localNamePic          = "pic"
	localNameGroupShape   = "grpSp"
	localNameGraphicFrame = "graphicFrame"
	localNameConnShape    = "cxnSp"
)

// Common PresentationML sizes.
const (
	Screen4x3Width  = 9144000
	Screen4x3Height = 6858000
)

// SlideSizeType represents the type of slide size.
type SlideSizeType string

// Slide size types.
const (
	SlideSizeScreen4x3   SlideSizeType = "screen4x3"
	SlideSizeLetter      SlideSizeType = "letter"
	SlideSizeA4          SlideSizeType = "A4"
	SlideSize35mm        SlideSizeType = "35mm"
	SlideSizeOverhead    SlideSizeType = "overhead"
	SlideSizeBanner      SlideSizeType = "banner"
	SlideSizeCustom      SlideSizeType = "custom"
	SlideSizeLedger      SlideSizeType = "ledger"
	SlideSizeA3          SlideSizeType = "A3"
	SlideSizeB4ISO       SlideSizeType = "B4ISO"
	SlideSizeB5ISO       SlideSizeType = "B5ISO"
	SlideSizeB4JIS       SlideSizeType = "B4JIS"
	SlideSizeB5JIS       SlideSizeType = "B5JIS"
	SlideSizeHagakiCard  SlideSizeType = "hagakiCard"
	SlideSizeScreen16x9  SlideSizeType = "screen16x9"
	SlideSizeScreen16x10 SlideSizeType = "screen16x10"
)

// ShapeType represents the type of a shape.
type ShapeType string

// Shape type constants.
const (
	ShapeTypeRectangle      ShapeType = "rect"
	ShapeTypeRoundRectangle ShapeType = "roundRect"
	ShapeTypeEllipse        ShapeType = "ellipse"
	ShapeTypeTriangle       ShapeType = "triangle"
	ShapeTypeDiamond        ShapeType = "diamond"
	ShapeTypeLine           ShapeType = "line"
)

// PlaceholderType represents the type of a placeholder.
type PlaceholderType string

// Placeholder type constants.
const (
	PlaceholderTitle         PlaceholderType = "title"
	PlaceholderBody          PlaceholderType = "body"
	PlaceholderCenteredTitle PlaceholderType = "ctrTitle"
	PlaceholderSubTitle      PlaceholderType = "subTitle"
	PlaceholderDateTime      PlaceholderType = "dt"
	PlaceholderSlideNumber   PlaceholderType = "sldNum"
	PlaceholderFooter        PlaceholderType = "ftr"
	PlaceholderHeader        PlaceholderType = "hdr"
	PlaceholderObject        PlaceholderType = "obj"
	PlaceholderChart         PlaceholderType = "chart"
	PlaceholderTable         PlaceholderType = "tbl"
	PlaceholderClipArt       PlaceholderType = "clipArt"
	PlaceholderDiagram       PlaceholderType = "dgm"
	PlaceholderMedia         PlaceholderType = "media"
	PlaceholderSlideImage    PlaceholderType = "sldImg"
	PlaceholderPicture       PlaceholderType = "pic"
)

// TransitionSpeed represents the speed of a transition.
type TransitionSpeed string

// Transition speed values.
const (
	TransitionSpeedSlow TransitionSpeed = "slow"
	TransitionSpeedMed  TransitionSpeed = "med"
	TransitionSpeedFast TransitionSpeed = "fast"
)

// wrapCompositeElement attempts to wrap an element as a CompositeElementBase.
func wrapCompositeElement(
	elem openxml.Element,
) *openxml.CompositeElementBase {
	if elem == nil {
		return nil
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return comp
	}
	// Many generated types embed CompositeElementBase but are not directly it.
	// We can't easily cast to the embedded field without reflection.
	return nil
}

// wrapLeafElement attempts to wrap an element as a LeafElementBase.
func wrapLeafElement(
	elem openxml.Element,
) *openxml.LeafElementBase {
	if elem == nil {
		return nil
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return leaf
	}

	return nil
}

// LinkGraphicFrameToChart sets up a GraphicFrame to reference a chart part.
// This creates the proper <a:graphic><a:graphicData> structure with a chart reference.
func LinkGraphicFrameToChart(
	gf *GraphicFrame,
	chartRelID string,
) {
	if gf == nil || chartRelID == "" {
		return
	}

	// Create the graphic element (a:graphic)
	graphic := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"graphic",
		PrefixA,
	)

	// Create the graphic data element (a:graphicData)
	// with the chart URI namespace
	graphicData := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"graphicData",
		PrefixA,
	)
	graphicData.SetAttribute(
		openxml.NewAttribute(
			"",
			"uri",
			"",
			"http://schemas.openxmlformats.org/drawingml/2006/chart",
		),
	)

	// Create the chart reference element (c:chart)
	chartRef := openxml.NewLeafElement(
		"http://schemas.openxmlformats.org/drawingml/2006/chart",
		"chart",
		"c",
	)
	chartRef.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"id",
			PrefixR,
			chartRelID,
		),
	)

	// Build the structure: graphic > graphicData > chartRef
	graphicData.AppendChild(chartRef)
	graphic.AppendChild(graphicData)
	gf.AppendChild(graphic)
}

// LinkGraphicFrameToTable sets up a GraphicFrame to contain a DrawingML table.
// This creates the proper <a:graphic><a:graphicData> structure with the table element.
func LinkGraphicFrameToTable(
	gf *GraphicFrame,
	table openxml.Element,
) {
	if gf == nil || table == nil {
		return
	}

	// Create the graphic element (a:graphic)
	graphic := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"graphic",
		PrefixA,
	)

	// Create the graphic data element (a:graphicData)
	// with the table URI namespace
	graphicData := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"graphicData",
		PrefixA,
	)
	graphicData.SetAttribute(
		openxml.NewAttribute(
			"",
			"uri",
			"",
			"http://schemas.openxmlformats.org/drawingml/2006/table",
		),
	)

	// Add the table element to graphicData
	graphicData.AppendChild(table)

	// Build the structure: graphic > graphicData > table
	graphic.AppendChild(graphicData)
	gf.AppendChild(graphic)
}
