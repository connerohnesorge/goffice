//nolint:revive // This file contains text body wrapper for presentation elements.
package elements

import (
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// Re-export DrawingML text types for convenience.
// The actual text body implementation is in the drawingml package.

// TextBody is an alias for drawingml.TextBody.
// Use this for text content within shapes.
type TextBody = drawingml.TextBody

// TextParagraph is an alias for drawingml.TextParagraph.
type TextParagraph = drawingml.TextParagraph

// TextRun is an alias for drawingml.TextRun.
type TextRun = drawingml.TextRun

// TextBodyProperties is an alias for drawingml.TextBodyProperties.
type TextBodyProperties = drawingml.TextBodyProperties

// TextParagraphProperties is an alias for drawingml.TextParagraphProperties.
type TextParagraphProperties = drawingml.TextParagraphProperties

// TextCharacterProperties is an alias for drawingml.TextCharacterProperties.
type TextCharacterProperties = drawingml.TextCharacterProperties

// NewTextBody creates a new text body element for presentations.
// Unlike drawingml.NewTextBody() which creates <a:txBody>, this creates
// <p:txBody> with the PresentationML namespace as required by PowerPoint.
func NewTextBody() *drawingml.TextBody {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"txBody",
		PrefixP,
	)
	tb := &drawingml.TextBody{
		CompositeElementBase: elem,
	}
	// Add default body properties (a:bodyPr)
	tb.AppendChild(
		drawingml.NewTextBodyProperties(),
	)

	return tb
}

// NewTextParagraph creates a new text paragraph element.
// This is a convenience wrapper for drawingml.NewTextParagraph.
func NewTextParagraph() *drawingml.TextParagraph {
	return drawingml.NewTextParagraph()
}

// NewTextRun creates a new text run element with the given text.
// This is a convenience wrapper for drawingml.NewTextRun.
func NewTextRun(text string) *drawingml.TextRun {
	return drawingml.NewTextRun(text)
}

// Text alignment values from drawingml.
const (
	TextAlignLeft        = drawingml.TextAlignLeft
	TextAlignCenter      = drawingml.TextAlignCenter
	TextAlignRight       = drawingml.TextAlignRight
	TextAlignJustify     = drawingml.TextAlignJustify
	TextAlignDistributed = drawingml.TextAlignDistributed
)

// Text vertical anchor values from drawingml.
const (
	TextAnchorTop    = drawingml.TextAnchorTop
	TextAnchorCenter = drawingml.TextAnchorCenter
	TextAnchorBottom = drawingml.TextAnchorBottom
)

// Underline values from drawingml.
const (
	UnderlineNone   = drawingml.UnderlineNone
	UnderlineSingle = drawingml.UnderlineSingle
	UnderlineDouble = drawingml.UnderlineDouble
	UnderlineHeavy  = drawingml.UnderlineHeavy
)
