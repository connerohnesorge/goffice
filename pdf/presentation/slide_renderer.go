package presentation

import (
	"fmt"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/presentation/elements"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// SlideRenderer renders a single slide to a PDF page.
// It handles master slide inheritance, backgrounds, and shape rendering.
type SlideRenderer struct {
	doc       *presentation.Document
	slidePart *parts.SlidePart
	page      *core.Page
	pageSize  core.PageSize
}

// NewSlideRenderer creates a new slide renderer.
func NewSlideRenderer(
	doc *presentation.Document,
	slidePart *parts.SlidePart,
	page *core.Page,
	pageSize core.PageSize,
) *SlideRenderer {
	return &SlideRenderer{
		doc:       doc,
		slidePart: slidePart,
		page:      page,
		pageSize:  pageSize,
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
		layoutSlideElem := layout.SlideLayout()
		if layoutSlide, ok := layoutSlideElem.(*elements.SlideLayout); ok {
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
		masterSlideElem := master.SlideMaster()
		if masterSlide, ok := masterSlideElem.(*elements.SlideMaster); ok {
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
	// For now, render white background
	// TODO: Parse and render actual background properties
	return sr.renderWhiteBackground()
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
	masterSlideElem := master.SlideMaster()
	masterSlide, ok := masterSlideElem.(*elements.SlideMaster)
	if !ok || masterSlide == nil {
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
	layoutSlideElem := layout.SlideLayout()
	layoutSlide, ok := layoutSlideElem.(*elements.SlideLayout)
	if !ok || layoutSlide == nil {
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

	// Get transform (position and size)
	xfrm := findChild(
		spPr,
		"xfrm",
		drawingml.NamespaceMain,
	)
	if xfrm == nil {
		return nil // No transform, can't position shape
	}

	// Extract position and size
	// TODO: Parse actual transform element
	// For now, just skip rendering the shape itself

	// Render text body if present
	textBody := shape.TextBody()
	if textBody != nil {
		// TODO: Render text
		_ = textBody
	}

	return nil
}

// renderPicture renders a picture shape.
func (sr *SlideRenderer) renderPicture(
	picture *elements.Picture,
) error {
	// TODO: Implement picture rendering
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
