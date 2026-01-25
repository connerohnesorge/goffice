//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package elements provides PresentationML element types.
//
// This package contains the core elements used in
// PresentationML documents (.pptx):
//
// # Presentation Structure
//
// The presentation hierarchy is:
//
//	Presentation
//	  -> SlideMasterIdList
//	  -> SlideIdList
//	  -> SlideSize
//	  -> NotesSize
//
// # Slide Elements
//
//   - Slide: Individual slide content
//   - CommonSlideData: Shared slide content container
//   - ShapeTree: Container for shapes and visual elements
//
// # Shape Elements
//
//   - Shape: Text shapes and placeholders
//   - Picture: Image elements
//   - GroupShape: Grouped shapes
//   - GraphicFrame: Charts, tables, and other objects
//
// # Text Elements
//
// Text elements leverage the shared DrawingML text types:
//
//   - TextBody: Container for text content
//   - TextParagraph: Individual paragraphs
//   - TextRun: Formatted text runs
//
// # Usage Example
//
//	pres := elements.NewPresentation()
//	slide := elements.NewSlide()
//	shape := elements.NewShape()
//	shape.TextBody().AddParagraph("Hello, World!")
//
// # Namespaces
//
// Standard PresentationML namespaces:
//
//	elements.NamespacePresentationML  // Main namespace
//	elements.PrefixP                  // "p:" prefix
//
//nolint:revive // max-public-structs: elements package defines many public types
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