package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Common PresentationML constants.
const (
	NamespacePresentationML   = openxml.NamespacePresentationML
	NamespaceDrawingML        = openxml.NamespaceDrawingML
	NamespaceDrawingMLPicture = openxml.NamespaceDrawingMLPicture
	NamespaceRelationships    = openxml.NamespaceRelationships
	PrefixP                   = "p"
	PrefixA                   = "a"
	PrefixPic                 = "pic"
	PrefixR                   = "r"
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
