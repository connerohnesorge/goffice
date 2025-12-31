// Text generator for E2E visual testing
//
// This file implements text generation using the goffice drawingml API.
package main

import (
	"fmt"
	"log"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/presentation/elements"
	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

// TextGenerator handles text creation for test cases
type TextGenerator struct {
	textSpec *framework.TextSpec
	verbose  bool
}

// NewTextGenerator creates a new text generator
func NewTextGenerator(
	textSpec *framework.TextSpec,
	verbose bool,
) *TextGenerator {
	return &TextGenerator{
		textSpec: textSpec,
		verbose:  verbose,
	}
}

// Generate creates a text shape and adds it to the slide
func (g *TextGenerator) Generate(
	slide *elements.Slide,
	pos framework.Position,
	size framework.Size,
) (*elements.Shape, error) {
	if g.verbose {
		log.Printf(
			"      Generating text element",
		)
	}

	// Create a shape to hold the text
	shape := slide.AddShape()
	if shape == nil {
		return nil, fmt.Errorf(
			"failed to create shape for text",
		)
	}

	// Get shape properties and set position/size
	spPr := shape.GetOrCreateShapeProperties()
	if spPr == nil {
		return nil, fmt.Errorf(
			"failed to get shape properties",
		)
	}

	spPr.SetOffset(int(pos.X), int(pos.Y))
	spPr.SetExtents(
		int(size.Width),
		int(size.Height),
	)

	// Use rectangle shape for text box
	spPr.SetPresetGeometry(
		string(drawingml.ShapeTypeRectangle),
	)

	// Make the shape transparent (no fill, no line) so only text is visible
	spPr.SetNoFill()
	ln := drawingml.NewLineProperties()
	ln.SetNoFill()
	spPr.AppendChild(ln)

	// Apply text to the shape
	if err := g.ApplyToShape(shape); err != nil {
		return nil, fmt.Errorf(
			"apply text: %w",
			err,
		)
	}

	return shape, nil
}

// ApplyToShape applies text content to an existing shape
func (g *TextGenerator) ApplyToShape(
	shape *elements.Shape,
) error {
	// Get or create text body
	textBody := shape.GetOrCreateTextBody()
	if textBody == nil {
		return fmt.Errorf(
			"failed to get text body",
		)
	}

	// Clear existing paragraphs
	textBody.ClearParagraphs()

	// Check if we have structured paragraphs or simple content
	if len(g.textSpec.Paragraphs) > 0 {
		// Generate from paragraph specs
		for _, paraSpec := range g.textSpec.Paragraphs {
			if err := g.addParagraph(textBody, &paraSpec); err != nil {
				return fmt.Errorf(
					"add paragraph: %w",
					err,
				)
			}
		}
	} else if g.textSpec.Content != "" {
		// Simple content - create a single paragraph with default formatting
		para := textBody.AddEmptyParagraph()
		run := drawingml.NewTextRun(g.textSpec.Content)

		// Apply default run formatting
		if err := g.applyRunFormatting(run, &g.textSpec.DefaultRun); err != nil {
			return fmt.Errorf("apply default formatting: %w", err)
		}

		para.AppendChild(run)
	}

	return nil
}

// addParagraph adds a paragraph with formatting to the text body
func (g *TextGenerator) addParagraph(
	textBody *drawingml.TextBody,
	paraSpec *framework.ParagraphSpec,
) error {
	para := textBody.AddEmptyParagraph()

	// Set paragraph properties
	paraPr := drawingml.NewTextParagraphProperties()

	// Set alignment
	if paraSpec.Alignment != "" {
		alignment := mapTextAlignment(
			paraSpec.Alignment,
		)
		paraPr.SetAlignment(alignment)
	}

	// Set line spacing
	if paraSpec.LineSpacing > 0 {
		// Line spacing in DrawingML is percentage (100% = 100000)
		paraPr.SetLineSpacingPercent(
			int(paraSpec.LineSpacing * 100000),
		)
	}

	// Set space before/after (convert EMUs to points)
	if paraSpec.SpaceBefore > 0 {
		paraPr.SetSpaceBeforePoints(
			int(paraSpec.SpaceBefore / 12700),
		) // EMUs to points
	}
	if paraSpec.SpaceAfter > 0 {
		paraPr.SetSpaceAfterPoints(
			int(paraSpec.SpaceAfter / 12700),
		) // EMUs to points
	}

	// Set indent (EMUs)
	if paraSpec.Indent > 0 {
		paraPr.SetIndent(int(paraSpec.Indent))
	}

	// Set bullet if specified
	if paraSpec.BulletChar != "" {
		paraPr.SetCharacterBullet(
			paraSpec.BulletChar,
		)
	}

	// Insert paragraph properties
	para.PrependChild(paraPr)

	// Add text runs
	for _, runSpec := range paraSpec.Runs {
		run := drawingml.NewTextRun(runSpec.Text)

		// Apply run formatting
		if err := g.applyRunFormatting(run, &runSpec); err != nil {
			return fmt.Errorf(
				"apply run formatting: %w",
				err,
			)
		}

		para.AppendChild(run)
	}

	return nil
}

// applyRunFormatting applies formatting to a text run
func (g *TextGenerator) applyRunFormatting(
	run *drawingml.TextRun,
	runSpec *framework.RunSpec,
) error {
	// Get or create run properties
	runPr := run.EnsureProperties()
	if runPr == nil {
		return fmt.Errorf(
			"failed to get run properties",
		)
	}

	// Set font size (points to hundredths of a point)
	if runSpec.FontSize > 0 {
		runPr.SetFontSize(runSpec.FontSize * 100)
	}

	// Set font family
	if runSpec.FontFamily != "" {
		runPr.SetLatinFont(runSpec.FontFamily)
	}

	// Set color
	if runSpec.Color != "" {
		runPr.SetSolidFill(runSpec.Color)
	}

	// Set bold
	if runSpec.Bold {
		runPr.SetBold(true)
	}

	// Set italic
	if runSpec.Italic {
		runPr.SetItalic(true)
	}

	// Set underline
	if runSpec.Underline != framework.UnderlineNone &&
		runSpec.Underline != "" {
		underline := mapUnderlineType(
			runSpec.Underline,
		)
		runPr.SetUnderline(underline)
	}

	// Set strikethrough
	if runSpec.Strikethrough {
		runPr.SetStrike(
			drawingml.StrikeSingleStrike,
		)
	}

	return nil
}

// mapTextAlignment maps framework text alignment to drawingml text alignment
func mapTextAlignment(
	align framework.TextAlignment,
) drawingml.TextAlignValue {
	switch align {
	case framework.TextAlignLeft:
		return drawingml.TextAlignLeft
	case framework.TextAlignCenter:
		return drawingml.TextAlignCenter
	case framework.TextAlignRight:
		return drawingml.TextAlignRight
	case framework.TextAlignJustify:
		return drawingml.TextAlignJustify
	default:
		return drawingml.TextAlignLeft
	}
}

// mapUnderlineType maps framework underline type to drawingml underline value
func mapUnderlineType(
	underline framework.UnderlineType,
) drawingml.UnderlineValue {
	switch underline {
	case framework.UnderlineSingle:
		return drawingml.UnderlineSingle
	case framework.UnderlineDouble:
		return drawingml.UnderlineDouble
	case framework.UnderlineDotted:
		return drawingml.UnderlineDotted
	case framework.UnderlineDash:
		return drawingml.UnderlineDash
	default:
		return drawingml.UnderlineNone
	}
}
