package drawing

import (
	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
)

// TextInShapeRenderer handles rendering text within shapes.
type TextInShapeRenderer struct {
	ctx *core.RenderingContext
}

// NewTextInShapeRenderer creates a new text-in-shape renderer.
func NewTextInShapeRenderer(
	ctx *core.RenderingContext,
) *TextInShapeRenderer {
	return &TextInShapeRenderer{ctx: ctx}
}

// RenderTextInShape renders text within a shape's bounds.
func (r *TextInShapeRenderer) RenderTextInShape(
	txBody *drawingml.TextBody,
	x, y, width, height float64,
) error {
	if txBody == nil {
		return nil
	}

	// Get body properties
	bodyPr := txBody.BodyProperties()
	if bodyPr == nil {
		bodyPr = drawingml.NewTextBodyProperties()
	}

	// Get insets (margins)
	leftInset := drawingml.EmuToPoints(
		drawingml.EMU(bodyPr.LeftInset()),
	)
	topInset := drawingml.EmuToPoints(
		drawingml.EMU(bodyPr.TopInset()),
	)
	rightInset := drawingml.EmuToPoints(
		drawingml.EMU(bodyPr.RightInset()),
	)
	bottomInset := drawingml.EmuToPoints(
		drawingml.EMU(bodyPr.BottomInset()),
	)

	// Calculate text area
	textX := x + leftInset
	textY := y + topInset
	textWidth := width - leftInset - rightInset
	textHeight := height - topInset - bottomInset

	// Get vertical anchor
	anchor := bodyPr.Anchor()

	// Get paragraphs
	paragraphs := txBody.Paragraphs()
	if len(paragraphs) == 0 {
		return nil
	}

	// Calculate total text height for vertical positioning
	totalTextHeight := r.calculateTextHeight(
		paragraphs,
		textWidth,
	)

	// Adjust Y position based on anchor
	var startY float64
	switch anchor {
	case drawingml.TextAnchorTop:
		startY = textY
	case drawingml.TextAnchorCenter:
		startY = textY + (textHeight-totalTextHeight)/2
	case drawingml.TextAnchorBottom:
		startY = textY + textHeight - totalTextHeight
	default:
		startY = textY
	}

	// Render each paragraph
	currentY := startY
	for _, para := range paragraphs {
		currentY = r.renderParagraph(
			para,
			textX,
			currentY,
			textWidth,
		)
	}

	return nil
}

// calculateTextHeight calculates the total height of text.
func (r *TextInShapeRenderer) calculateTextHeight(
	paragraphs []*drawingml.TextParagraph,
	width float64,
) float64 {
	totalHeight := 0.0

	for _, para := range paragraphs {
		// Get text runs
		runs := para.Runs()
		if len(runs) == 0 {
			totalHeight += 12.0 // Default line height for empty paragraph

			continue
		}

		// Calculate paragraph height based on font size
		maxFontSize := 12.0 // Default
		for _, run := range runs {
			if props := run.Properties(); props != nil {
				fontSize := props.FontSizePoints()
				if fontSize > maxFontSize {
					maxFontSize = fontSize
				}
			}
		}

		// Line height is typically 1.2x font size
		lineHeight := maxFontSize * 1.2
		totalHeight += lineHeight
	}

	return totalHeight
}

// renderParagraph renders a single paragraph.
func (r *TextInShapeRenderer) renderParagraph(
	para *drawingml.TextParagraph,
	x, y, width float64,
) float64 {
	runs := para.Runs()
	if len(runs) == 0 {
		return y + 12.0 // Empty paragraph
	}

	// Get paragraph properties
	props := para.Properties()
	var alignment drawingml.TextAlignValue
	if props != nil {
		alignment = props.Alignment()
	}

	// Calculate line metrics
	maxFontSize := 12.0
	text := ""
	for _, run := range runs {
		text += run.Text()
		if rProps := run.Properties(); rProps != nil {
			fontSize := rProps.FontSizePoints()
			if fontSize > 0 &&
				fontSize > maxFontSize {
				maxFontSize = fontSize
			}
		}
	}

	lineHeight := maxFontSize * 1.2

	// Calculate text width for alignment
	textWidth := r.estimateTextWidth(
		text,
		maxFontSize,
	)

	// Adjust X based on alignment
	renderX := x
	switch alignment {
	case drawingml.TextAlignCenter:
		if textWidth < width {
			renderX = x + (width-textWidth)/2
		}
	case drawingml.TextAlignRight:
		if textWidth < width {
			renderX = x + width - textWidth
		}
	case drawingml.TextAlignLeft,
		drawingml.TextAlignJustify:
		// Left aligned or justified - stay at x
	}

	// Render each run
	currentX := renderX
	for _, run := range runs {
		currentX = r.renderRun(
			run,
			currentX,
			y,
			maxFontSize,
		)
	}

	return y + lineHeight
}

// renderRun renders a single text run.
func (r *TextInShapeRenderer) renderRun(
	run *drawingml.TextRun,
	x, y, defaultFontSize float64,
) float64 {
	text := run.Text()
	if text == "" {
		return x
	}

	// Get run properties
	props := run.Properties()

	fontSize := defaultFontSize
	fontName := "Helvetica"
	bold := false
	italic := false
	color := Black

	if props != nil {
		if props.FontSize() > 0 {
			fontSize = props.FontSizePoints()
		}
		bold = props.Bold()
		italic = props.Italic()

		// Get font
		if latinFont := props.LatinFont(); latinFont != nil {
			if typeface := latinFont.Typeface(); typeface != "" {
				fontName = typeface
			}
		}

		// Get color
		if solidFill := props.SolidFill(); solidFill != nil {
			if rgb := solidFill.RgbColor(); rgb != nil {
				hexAttr, found := rgb.GetAttribute(
					"val",
					"",
				)
				if found {
					color = ParseColor(
						hexAttr.Value(),
					)
				}
			}
		}
	}

	// Select font variant based on bold/italic
	if bold && italic {
		fontName += "-BoldItalic"
	} else if bold {
		fontName += "-Bold"
	} else if italic {
		fontName += "-Italic"
	}

	// Render the text
	r.ctx.Page.SetFont(fontName, fontSize)
	r.ctx.Page.SetFillColor(
		color.R,
		color.G,
		color.B,
	)
	r.ctx.Page.DrawText(x, y, text)

	// Estimate text width to advance X
	textWidth := r.estimateTextWidth(
		text,
		fontSize,
	)

	return x + textWidth
}

// estimateTextWidth estimates the width of text.
func (r *TextInShapeRenderer) estimateTextWidth(
	text string,
	fontSize float64,
) float64 {
	// Rough estimation: average character width is about 0.5 * fontSize
	charCount := float64(len([]rune(text)))

	return charCount * fontSize * 0.5
}

// AutofitText adjusts text size to fit within shape bounds.
func (r *TextInShapeRenderer) AutofitText(
	txBody *drawingml.TextBody,
	width, height float64,
) {
	if txBody == nil {
		return
	}

	bodyPr := txBody.BodyProperties()
	if bodyPr == nil {
		return
	}

	// Check if autofit is enabled
	// This would check for autofit elements like <a:normAutofit> or <a:spAutoFit>
	// For now, this is a placeholder for autofit logic

	paragraphs := txBody.Paragraphs()
	totalHeight := r.calculateTextHeight(
		paragraphs,
		width,
	)

	// If text exceeds height, shrink font sizes proportionally
	if totalHeight > height && height > 0 {
		scaleFactor := height / totalHeight
		if scaleFactor < 0.5 {
			scaleFactor = 0.5 // Minimum 50% scaling
		}

		// Apply scaling to all runs
		for _, para := range paragraphs {
			for _, run := range para.Runs() {
				if props := run.Properties(); props != nil {
					currentSize := props.FontSize()
					if currentSize > 0 {
						newSize := int(
							float64(
								currentSize,
							) * scaleFactor,
						)
						props.SetFontSize(newSize)
					}
				}
			}
		}
	}
}

// RenderVerticalText renders vertical text (rotated 90 or 270 degrees).
func (r *TextInShapeRenderer) RenderVerticalText(
	txBody *drawingml.TextBody,
	x, y, width, height float64,
) error {
	if txBody == nil {
		return nil
	}

	bodyPr := txBody.BodyProperties()
	if bodyPr == nil {
		return nil
	}

	vertical := bodyPr.Vertical()

	switch vertical {
	case drawingml.TextVerticalRotate90:
		// Rotate 90 degrees clockwise
		// In PDF, this requires a rotation transformation
		r.ctx.Page.SaveGraphicsState()
		defer r.ctx.Page.RestoreGraphicsState()

		// Apply rotation transform
		// For now, render normally (full rotation support requires transform matrix)
		return r.RenderTextInShape(
			txBody,
			x,
			y,
			width,
			height,
		)
	case drawingml.TextVerticalRotate270:
		// Rotate 270 degrees clockwise (90 counter-clockwise)
		r.ctx.Page.SaveGraphicsState()
		defer r.ctx.Page.RestoreGraphicsState()

		return r.RenderTextInShape(
			txBody,
			x,
			y,
			width,
			height,
		)
	case drawingml.TextVerticalHorizontal,
		drawingml.TextVerticalWordArt,
		drawingml.TextVerticalEAVert,
		drawingml.TextVerticalMongolian,
		drawingml.TextVerticalWordArtRtl:
		fallthrough
	default:
		// Normal horizontal text
		return r.RenderTextInShape(
			txBody,
			x,
			y,
			width,
			height,
		)
	}
}
