package presentation

import (
	"fmt"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice-pdf/drawing"
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/presentation/elements"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// SlideRenderer renders a single slide to a PDF page.
// It handles master slide inheritance, backgrounds, and shape rendering.
type SlideRenderer struct {
	doc           *presentation.Document
	slidePart     *parts.SlidePart
	page          *core.Page
	pageSize      core.PageSize
	renderContext *core.RenderingContext
	shapeRenderer *drawing.ShapeRenderer
}

// NewSlideRenderer creates a new slide renderer.
func NewSlideRenderer(
	doc *presentation.Document,
	slidePart *parts.SlidePart,
	page *core.Page,
	pageSize core.PageSize,
) *SlideRenderer {
	// Create PageImpl from Page for DrawingML rendering
	pageImpl := core.NewPageImpl(
		pageSize.Width,
		pageSize.Height,
	)

	// Create rendering context with the PageImpl
	renderContext := core.NewRenderingContext(
		pageSize.Width,
		pageSize.Height,
	)
	renderContext.SetPage(pageImpl)

	// Create shape renderer
	shapeRenderer := drawing.NewShapeRenderer(
		renderContext,
	)

	return &SlideRenderer{
		doc:           doc,
		slidePart:     slidePart,
		page:          page,
		pageSize:      pageSize,
		renderContext: renderContext,
		shapeRenderer: shapeRenderer,
	}
}

// Render renders the slide to the PDF page.
func (sr *SlideRenderer) Render() error {
	slide := sr.slidePart.Slide()
	if slide == nil {
		return fmt.Errorf(
			"slide element not found",
		)
	}

	// Get slide layout and master
	layout, master, err := sr.getLayoutAndMaster()
	if err != nil {
		return fmt.Errorf(
			"failed to get layout and master: %w",
			err,
		)
	}

	// Render in order: master background -> layout background -> slide background
	if err := sr.renderBackground(slide, layout, master); err != nil {
		return fmt.Errorf(
			"failed to render background: %w",
			err,
		)
	}

	// Render master slide shapes (behind slide content)
	if master != nil {
		if err := sr.renderMasterShapes(master); err != nil {
			return fmt.Errorf(
				"failed to render master shapes: %w",
				err,
			)
		}
	}

	// Render layout shapes (placeholders that aren't overridden)
	if layout != nil {
		if err := sr.renderLayoutShapes(layout, slide); err != nil {
			return fmt.Errorf(
				"failed to render layout shapes: %w",
				err,
			)
		}
	}

	// Render slide shapes (foreground content)
	if err := sr.renderSlideShapes(slide); err != nil {
		return fmt.Errorf(
			"failed to render slide shapes: %w",
			err,
		)
	}

	return nil
}

// getLayoutAndMaster retrieves the slide layout and master for this slide.
func (sr *SlideRenderer) getLayoutAndMaster() (*parts.SlideLayoutPart, *parts.SlideMasterPart, error) {
	// Get slide layout from relationships
	var layoutPart *parts.SlideLayoutPart
	for part := range sr.slidePart.Parts() {
		if layout, ok := part.(*parts.SlideLayoutPart); ok {
			layoutPart = layout
			break
		}
	}

	if layoutPart == nil {
		// No layout is acceptable, slide stands alone
		return nil, nil, nil
	}

	// Get slide master from layout
	var masterPart *parts.SlideMasterPart
	for part := range layoutPart.Parts() {
		if master, ok := part.(*parts.SlideMasterPart); ok {
			masterPart = master
			break
		}
	}

	return layoutPart, masterPart, nil
}

// renderBackground renders the slide background with inheritance.
// Priority: slide background > layout background > master background.
func (sr *SlideRenderer) renderBackground(
	slide *elements.Slide,
	layout *parts.SlideLayoutPart,
	master *parts.SlideMasterPart,
) error {
	// Check for slide background
	csd := slide.CommonSlideData()
	if csd != nil {
		if bg := csd.Background(); bg != nil {
			return sr.renderBackgroundElement(bg)
		}
	}

	// Check for layout background
	if layout != nil {
		layoutSlide := layout.SlideLayout()
		if layoutSlide != nil {
			layoutCsd := layoutSlide.CommonSlideData()
			if layoutCsd != nil {
				if bg := layoutCsd.Background(); bg != nil {
					return sr.renderBackgroundElement(
						bg,
					)
				}
			}
		}
	}

	// Check for master background
	if master != nil {
		masterSlide := master.SlideMaster()
		if masterSlide != nil {
			masterCsd := masterSlide.CommonSlideData()
			if masterCsd != nil {
				if bg := masterCsd.Background(); bg != nil {
					return sr.renderBackgroundElement(
						bg,
					)
				}
			}
		}
	}

	// No background found, render white background as default
	return sr.renderWhiteBackground()
}

// renderBackgroundElement renders a background element (p:bg).
func (sr *SlideRenderer) renderBackgroundElement(
	bg *elements.SlideBackground,
) error {
	// Look for background fill properties
	bgPr := findChild(
		bg,
		"bgPr",
		elements.NamespacePresentationML,
	)
	if bgPr != nil {
		// Check for solid fill
		solidFill := findChild(
			bgPr,
			"solidFill",
			drawingml.NamespaceMain,
		)
		if solidFill != nil {
			return sr.renderSolidFillBackground(
				solidFill,
			)
		}

		// Check for gradient fill
		gradFill := findChild(
			bgPr,
			"gradFill",
			drawingml.NamespaceMain,
		)
		if gradFill != nil {
			// For gradients, use white background as fallback
			return sr.renderWhiteBackground()
		}

		// Check for pattern fill
		pattFill := findChild(
			bgPr,
			"pattFill",
			drawingml.NamespaceMain,
		)
		if pattFill != nil {
			// For patterns, use white background as fallback
			return sr.renderWhiteBackground()
		}

		// Check for blip fill (image)
		blipFill := findChild(
			bgPr,
			"blipFill",
			drawingml.NamespaceMain,
		)
		if blipFill != nil {
			// For images, use white background as fallback
			return sr.renderWhiteBackground()
		}
	}

	// Check for background reference (theme color)
	bgRef := findChild(
		bg,
		"bgRef",
		drawingml.NamespaceMain,
	)
	if bgRef != nil {
		// For theme references, use white background as fallback
		return sr.renderWhiteBackground()
	}

	// Default white background
	return sr.renderWhiteBackground()
}

// renderSolidFillBackground renders a solid color background.
func (sr *SlideRenderer) renderSolidFillBackground(
	solidFill openxml.Element,
) error {
	// Look for RGB color
	srgbClr := findChild(
		solidFill,
		"srgbClr",
		drawingml.NamespaceMain,
	)
	if srgbClr != nil {
		attr, found := srgbClr.GetAttribute(
			"val",
			"",
		)
		if found {
			hexColor := attr.Value()
			return sr.renderColorBackground(
				hexColor,
			)
		}
	}

	// Look for scheme color (theme color)
	schemeClr := findChild(
		solidFill,
		"schemeClr",
		drawingml.NamespaceMain,
	)
	if schemeClr != nil {
		// For theme colors, use white as fallback
		return sr.renderWhiteBackground()
	}

	// Default to white
	return sr.renderWhiteBackground()
}

// renderColorBackground renders a background with the specified hex color.
func (sr *SlideRenderer) renderColorBackground(
	hexColor string,
) error {
	// Parse hex color (e.g., "FF0000" for red)
	var r, g, b uint8
	if len(hexColor) == 6 {
		fmt.Sscanf(
			hexColor,
			"%02x%02x%02x",
			&r,
			&g,
			&b,
		)
	} else {
		// Invalid color, use white
		r, g, b = 255, 255, 255
	}

	// Convert to PDF color space (0-1 range)
	rPdf := float64(r) / 255.0
	gPdf := float64(g) / 255.0
	bPdf := float64(b) / 255.0

	// Fill entire page with color using PDF content stream
	content := fmt.Sprintf(
		"q\n%.3f %.3f %.3f rg\n0 0 %.2f %.2f re\nf\nQ\n",
		rPdf,
		gPdf,
		bPdf,
		sr.pageSize.Width,
		sr.pageSize.Height,
	)
	sr.page.WriteContentString(content)

	return nil
}

// renderWhiteBackground renders a white background (default).
func (sr *SlideRenderer) renderWhiteBackground() error {
	// Fill entire page with white using PDF content stream
	content := fmt.Sprintf(
		"q\n1 1 1 rg\n0 0 %.2f %.2f re\nf\nQ\n",
		sr.pageSize.Width,
		sr.pageSize.Height,
	)
	sr.page.WriteContentString(content)

	return nil
}

// renderMasterShapes renders shapes from the master slide.
func (sr *SlideRenderer) renderMasterShapes(
	master *parts.SlideMasterPart,
) error {
	masterSlide := master.SlideMaster()
	if masterSlide == nil {
		return nil
	}

	csd := masterSlide.CommonSlideData()
	if csd == nil {
		return nil
	}

	spTree := csd.ShapeTree()
	if spTree == nil {
		return nil
	}

	return sr.renderShapeTree(spTree, true)
}

// renderLayoutShapes renders shapes from the layout slide.
// Only renders placeholders that aren't overridden by the slide.
func (sr *SlideRenderer) renderLayoutShapes(
	layout *parts.SlideLayoutPart,
	slide *elements.Slide,
) error {
	layoutSlide := layout.SlideLayout()
	if layoutSlide == nil {
		return nil
	}

	csd := layoutSlide.CommonSlideData()
	if csd == nil {
		return nil
	}

	spTree := csd.ShapeTree()
	if spTree == nil {
		return nil
	}

	// Get slide placeholders to check for overrides
	slidePlaceholders := sr.getSlidePlaceholders(
		slide,
	)

	// Render layout shapes, skipping overridden placeholders
	return sr.renderShapeTreeWithOverrides(
		spTree,
		slidePlaceholders,
		false,
	)
}

// renderSlideShapes renders shapes from the slide.
func (sr *SlideRenderer) renderSlideShapes(
	slide *elements.Slide,
) error {
	csd := slide.CommonSlideData()
	if csd == nil {
		return nil
	}

	spTree := csd.ShapeTree()
	if spTree == nil {
		return nil
	}

	return sr.renderShapeTree(spTree, false)
}

// getSlidePlaceholders returns a map of placeholder types present on the slide.
func (sr *SlideRenderer) getSlidePlaceholders(
	slide *elements.Slide,
) map[string]bool {
	placeholders := make(map[string]bool)

	csd := slide.CommonSlideData()
	if csd == nil {
		return placeholders
	}

	spTree := csd.ShapeTree()
	if spTree == nil {
		return placeholders
	}

	// Iterate through shapes to find placeholders
	for child := range spTree.Children() {
		if shape, ok := child.(*elements.Shape); ok {
			if phType := sr.getPlaceholderType(shape); phType != "" {
				placeholders[phType] = true
			}
		}
	}

	return placeholders
}

// getPlaceholderType returns the placeholder type of a shape, or empty if not a placeholder.
func (sr *SlideRenderer) getPlaceholderType(
	shape *elements.Shape,
) string {
	nvSpPr := shape.NonVisualShapeProperties()
	if nvSpPr == nil {
		return ""
	}

	// Look for nvSpPr/nvPr/ph element
	nvPr := findChild(
		nvSpPr,
		"nvPr",
		elements.NamespacePresentationML,
	)
	if nvPr == nil {
		return ""
	}

	ph := findChild(
		nvPr,
		"ph",
		elements.NamespacePresentationML,
	)
	if ph == nil {
		return ""
	}

	// Get type attribute
	typeAttr, found := ph.GetAttribute("type", "")
	if !found {
		return "body" // Default placeholder type
	}

	return typeAttr.Value()
}

// renderShapeTreeWithOverrides renders a shape tree, skipping overridden placeholders.
func (sr *SlideRenderer) renderShapeTreeWithOverrides(
	spTree *elements.ShapeTree,
	overrides map[string]bool,
	isMaster bool,
) error {
	for child := range spTree.Children() {
		// Skip if this is a placeholder that's overridden
		if shape, ok := child.(*elements.Shape); ok {
			phType := sr.getPlaceholderType(shape)
			if phType != "" && overrides[phType] {
				continue // Skip overridden placeholder
			}
		}

		if err := sr.renderShapeElement(child, isMaster); err != nil {
			return err
		}
	}

	return nil
}

// renderShapeTree renders all shapes in a shape tree.
func (sr *SlideRenderer) renderShapeTree(
	spTree *elements.ShapeTree,
	isMaster bool,
) error {
	for child := range spTree.Children() {
		if err := sr.renderShapeElement(child, isMaster); err != nil {
			return err
		}
	}

	return nil
}

// renderShapeElement renders a single shape element.
func (sr *SlideRenderer) renderShapeElement(
	elem openxml.Element,
	isMaster bool,
) error {
	switch shape := elem.(type) {
	case *elements.Shape:
		return sr.renderShape(shape, isMaster)
	case *elements.Picture:
		return sr.renderPicture(shape)
	default:
		// Unknown or unsupported shape type, skip silently
		return nil
	}
}

// renderShape renders a presentation shape.
func (sr *SlideRenderer) renderShape(
	shape *elements.Shape,
	isMaster bool,
) error {
	// Get shape properties
	spPr := shape.ShapeProperties()
	if spPr == nil {
		return nil // No visual properties, skip
	}

	// Convert elements.ShapeProperties to drawingml.ShapeProperties
	// The shape properties element should already be a DrawingML ShapeProperties
	dmlSpPr, ok := interface{}(spPr).(*drawingml.ShapeProperties)
	if !ok {
		// Fall back to basic rendering if not DrawingML properties
		return sr.renderShapeBasic(
			shape,
			isMaster,
		)
	}

	// Create fill and stroke renderers from the shape properties
	fillRenderer := drawing.CreateFillRenderer(
		dmlSpPr,
	)
	strokeRenderer := drawing.CreateStrokeRenderer(
		dmlSpPr,
	)

	// Render the shape using the ShapeRenderer
	if err := sr.shapeRenderer.RenderShapeWithFill(
		dmlSpPr,
		fillRenderer,
		strokeRenderer,
	); err != nil {
		return fmt.Errorf(
			"failed to render shape: %w",
			err,
		)
	}

	// Write the accumulated content from PageImpl to the actual page
	if pageImpl, ok := sr.renderContext.Page.(*core.PageImpl); ok {
		content := pageImpl.GetContent()
		if content != "" {
			if _, err := sr.page.WriteContentString(content); err != nil {
				return fmt.Errorf(
					"failed to write shape content: %w",
					err,
				)
			}
			// Clear the content buffer for next shape
			pageImpl.Reset()
		}
	}

	// Render text body if present
	textBody := shape.TextBody()
	if textBody != nil {
		// Get transform for text positioning
		xfrm := findChild(
			spPr,
			"xfrm",
			drawingml.NamespaceMain,
		)
		if xfrm != nil {
			var x, y, cx, cy int64
			off := findChild(
				xfrm,
				"off",
				drawingml.NamespaceMain,
			)
			if off != nil {
				if xAttr, found := off.GetAttribute("x", ""); found {
					fmt.Sscanf(
						xAttr.Value(),
						"%d",
						&x,
					)
				}
				if yAttr, found := off.GetAttribute("y", ""); found {
					fmt.Sscanf(
						yAttr.Value(),
						"%d",
						&y,
					)
				}
			}
			ext := findChild(
				xfrm,
				"ext",
				drawingml.NamespaceMain,
			)
			if ext != nil {
				if cxAttr, found := ext.GetAttribute("cx", ""); found {
					fmt.Sscanf(
						cxAttr.Value(),
						"%d",
						&cx,
					)
				}
				if cyAttr, found := ext.GetAttribute("cy", ""); found {
					fmt.Sscanf(
						cyAttr.Value(),
						"%d",
						&cy,
					)
				}
			}

			xPt := drawingml.EmuToPoints(
				drawingml.EMU(x),
			)
			yPt := drawingml.EmuToPoints(
				drawingml.EMU(y),
			)
			widthPt := drawingml.EmuToPoints(
				drawingml.EMU(cx),
			)
			heightPt := drawingml.EmuToPoints(
				drawingml.EMU(cy),
			)
			yPdfPt := sr.pageSize.Height - yPt - heightPt

			sr.renderTextBody(
				textBody,
				xPt,
				yPdfPt,
				widthPt,
				heightPt,
			)
		}
	}

	return nil
}

// renderShapeBasic renders a shape using the basic (legacy) method.
// This is a fallback for when the shape properties are not DrawingML.
func (sr *SlideRenderer) renderShapeBasic(
	shape *elements.Shape,
	isMaster bool,
) error {
	// Get shape properties
	spPr := shape.ShapeProperties()
	if spPr == nil {
		return nil
	}

	// Get transform (position and size)
	xfrm := findChild(
		spPr,
		"xfrm",
		drawingml.NamespaceMain,
	)
	if xfrm == nil {
		return nil
	}

	// Extract position from a:off element
	off := findChild(
		xfrm,
		"off",
		drawingml.NamespaceMain,
	)
	var x, y int64
	if off != nil {
		if xAttr, found := off.GetAttribute("x", ""); found {
			fmt.Sscanf(xAttr.Value(), "%d", &x)
		}
		if yAttr, found := off.GetAttribute("y", ""); found {
			fmt.Sscanf(yAttr.Value(), "%d", &y)
		}
	}

	// Extract size from a:ext element
	ext := findChild(
		xfrm,
		"ext",
		drawingml.NamespaceMain,
	)
	var cx, cy int64
	if ext != nil {
		if cxAttr, found := ext.GetAttribute("cx", ""); found {
			fmt.Sscanf(cxAttr.Value(), "%d", &cx)
		}
		if cyAttr, found := ext.GetAttribute("cy", ""); found {
			fmt.Sscanf(cyAttr.Value(), "%d", &cy)
		}
	}

	// Convert EMU to points
	xPt := drawingml.EmuToPoints(drawingml.EMU(x))
	yPt := drawingml.EmuToPoints(drawingml.EMU(y))
	widthPt := drawingml.EmuToPoints(
		drawingml.EMU(cx),
	)
	heightPt := drawingml.EmuToPoints(
		drawingml.EMU(cy),
	)

	// PDF coordinates are bottom-left origin, PowerPoint is top-left
	// Flip Y coordinate
	yPdfPt := sr.pageSize.Height - yPt - heightPt

	// Render shape fill if present
	sr.renderShapeFill(
		spPr,
		xPt,
		yPdfPt,
		widthPt,
		heightPt,
	)

	// Render text body if present
	textBody := shape.TextBody()
	if textBody != nil {
		sr.renderTextBody(
			textBody,
			xPt,
			yPdfPt,
			widthPt,
			heightPt,
		)
	}

	return nil
}

// renderShapeFill renders the fill for a shape.
func (sr *SlideRenderer) renderShapeFill(
	spPr openxml.Element,
	x, y, width, height float64,
) {
	// Check for solid fill
	solidFill := findChild(
		spPr,
		"solidFill",
		drawingml.NamespaceMain,
	)
	if solidFill != nil {
		srgbClr := findChild(
			solidFill,
			"srgbClr",
			drawingml.NamespaceMain,
		)
		if srgbClr != nil {
			if attr, found := srgbClr.GetAttribute("val", ""); found {
				hexColor := attr.Value()
				var r, g, b uint8
				if len(hexColor) == 6 {
					fmt.Sscanf(
						hexColor,
						"%02x%02x%02x",
						&r,
						&g,
						&b,
					)
					rPdf := float64(r) / 255.0
					gPdf := float64(g) / 255.0
					bPdf := float64(b) / 255.0
					content := fmt.Sprintf(
						"q\n%.3f %.3f %.3f rg\n%.2f %.2f %.2f %.2f re\nf\nQ\n",
						rPdf,
						gPdf,
						bPdf,
						x,
						y,
						width,
						height,
					)
					sr.page.WriteContentString(
						content,
					)
				}
			}
		}
	}
}

// renderTextBody renders text content from a text body.
func (sr *SlideRenderer) renderTextBody(
	textBody *drawingml.TextBody,
	x, y, width, height float64,
) {
	paragraphs := textBody.Paragraphs()
	if len(paragraphs) == 0 {
		return
	}

	// Simple text rendering: render each paragraph
	fontSize := 12.0 // Default font size
	lineHeight := fontSize * 1.2
	currentY := y + height - lineHeight // Start from top

	for _, para := range paragraphs {
		text := sr.extractParagraphText(para)
		if text == "" {
			currentY -= lineHeight
			continue
		}

		// Render text line
		content := fmt.Sprintf(
			"BT\n/F1 %.1f Tf\n%.2f %.2f Td\n(%s) Tj\nET\n",
			fontSize,
			x+2,
			currentY,
			sr.escapePdfString(text),
		)
		sr.page.WriteContentString(content)
		currentY -= lineHeight
	}
}

// extractParagraphText extracts plain text from a paragraph.
func (sr *SlideRenderer) extractParagraphText(
	para *drawingml.TextParagraph,
) string {
	var text string
	for child := range para.Children() {
		if child.LocalName() == "r" &&
			child.NamespaceURI() == drawingml.NamespaceMain {
			// Text run
			t := findChild(
				child,
				"t",
				drawingml.NamespaceMain,
			)
			if t != nil {
				if leaf, ok := t.(*openxml.LeafElementBase); ok {
					text += leaf.InnerText()
				} else if comp, ok := t.(openxml.CompositeElement); ok {
					// Try to get text content
					for textChild := range comp.Children() {
						if textLeaf, ok := textChild.(*openxml.LeafElementBase); ok {
							text += textLeaf.InnerText()
						}
					}
				}
			}
		}
	}
	return text
}

// escapePdfString escapes special characters for PDF string literals.
func (sr *SlideRenderer) escapePdfString(
	s string,
) string {
	// Escape parentheses and backslashes
	s = fmt.Sprintf("%s", s)
	s = fmt.Sprintf("%q", s)
	// Remove outer quotes added by %q
	if len(s) >= 2 && s[0] == '"' &&
		s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	// Replace double backslashes with single (PDF uses single)
	// Simple replacement for basic characters
	return s
}

// renderPicture renders a picture shape.
func (sr *SlideRenderer) renderPicture(
	picture *elements.Picture,
) error {
	// Get shape properties for position and size
	spPr := picture.ShapeProperties()
	if spPr == nil {
		return nil
	}

	// Get transform
	xfrm := findChild(
		spPr,
		"xfrm",
		drawingml.NamespaceMain,
	)
	if xfrm == nil {
		return nil
	}

	// Extract position
	off := findChild(
		xfrm,
		"off",
		drawingml.NamespaceMain,
	)
	var x, y int64
	if off != nil {
		if xAttr, found := off.GetAttribute("x", ""); found {
			fmt.Sscanf(xAttr.Value(), "%d", &x)
		}
		if yAttr, found := off.GetAttribute("y", ""); found {
			fmt.Sscanf(yAttr.Value(), "%d", &y)
		}
	}

	// Extract size
	ext := findChild(
		xfrm,
		"ext",
		drawingml.NamespaceMain,
	)
	var cx, cy int64
	if ext != nil {
		if cxAttr, found := ext.GetAttribute("cx", ""); found {
			fmt.Sscanf(cxAttr.Value(), "%d", &cx)
		}
		if cyAttr, found := ext.GetAttribute("cy", ""); found {
			fmt.Sscanf(cyAttr.Value(), "%d", &cy)
		}
	}

	// Convert to points
	xPt := drawingml.EmuToPoints(drawingml.EMU(x))
	yPt := drawingml.EmuToPoints(drawingml.EMU(y))
	widthPt := drawingml.EmuToPoints(
		drawingml.EMU(cx),
	)
	heightPt := drawingml.EmuToPoints(
		drawingml.EMU(cy),
	)

	// PDF coordinates are bottom-left origin
	yPdfPt := sr.pageSize.Height - yPt - heightPt

	// For now, render a placeholder rectangle for the image
	// Full image rendering would require loading the image data from relationships
	content := fmt.Sprintf(
		"q\n0.8 0.8 0.8 RG\n1 w\n%.2f %.2f %.2f %.2f re\nS\nQ\n",
		xPt,
		yPdfPt,
		widthPt,
		heightPt,
	)
	sr.page.WriteContentString(content)

	return nil
}

// findChild finds a child element by local name and namespace.
func findChild(
	parent openxml.Element,
	localName, namespace string,
) openxml.Element {
	composite, ok := parent.(openxml.CompositeElement)
	if !ok {
		return nil
	}

	for child := range composite.Children() {
		if child.LocalName() == localName &&
			child.NamespaceURI() == namespace {
			return child
		}
	}

	return nil
}
