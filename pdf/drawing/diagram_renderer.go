package drawing

import (
	"github.com/connerohnesorge/goffice-pdf/core"
)

const (
	// placeholderDiagramColorR is the red component for diagram placeholder rectangles.
	placeholderDiagramColorR = 0.95
	// placeholderDiagramColorG is the green component for diagram placeholder rectangles.
	placeholderDiagramColorG = 0.95
	// placeholderDiagramColorB is the blue component for diagram placeholder rectangles.
	placeholderDiagramColorB = 0.95
	// placeholderDiagramBorderR is the red component for diagram placeholder border.
	placeholderDiagramBorderR = 0.6
	// placeholderDiagramBorderG is the green component for diagram placeholder border.
	placeholderDiagramBorderG = 0.6
	// placeholderDiagramBorderB is the blue component for diagram placeholder border.
	placeholderDiagramBorderB = 0.7
)

// DiagramRenderer handles rendering of SmartArt diagrams to PDF.
//
// Phase 1 Implementation Note:
// This renderer currently provides placeholder rendering only. Full diagram
// rendering with layout engines and style application will be implemented in
// Phase 2 of the SmartArt support project.
type DiagramRenderer struct {
	ctx *core.RenderingContext
}

// NewDiagramRenderer creates a new diagram renderer.
// The rendering context must have its Page field set before rendering.
func NewDiagramRenderer(
	ctx *core.RenderingContext,
) *DiagramRenderer {
	return &DiagramRenderer{ctx: ctx}
}

// RenderPlaceholder renders a placeholder box for a diagram.
// This is used in Phase 1 to indicate the presence of a SmartArt diagram
// without implementing the full layout and rendering engine.
//
// The placeholder consists of:
//   - A light gray filled rectangle
//   - A dashed border in a darker gray/blue color
//   - Centered text label "SmartArt Diagram"
//
// Parameters:
//   - x, y: Bottom-left corner position in PDF points
//   - width, height: Dimensions in PDF points
//
// Full diagram rendering will be implemented in Phase 2, which will include:
//   - Data model parsing and layout algorithm execution
//   - Style and color application
//   - Text rendering within diagram shapes
//   - Connection lines and arrows
func (r *DiagramRenderer) RenderPlaceholder(
	x, y, width, height float64,
) error {
	if r.ctx.Page == nil {
		// If no page is set, silently skip rendering
		// This matches the pattern used by other renderers
		return nil
	}

	page := r.ctx.Page

	// Save graphics state to restore after rendering
	page.SaveGraphicsState()
	defer page.RestoreGraphicsState()

	// Draw filled background rectangle
	page.SetFillColor(
		placeholderDiagramColorR,
		placeholderDiagramColorG,
		placeholderDiagramColorB,
	)
	page.DrawRectangle(
		x,
		y,
		width,
		height,
		true,
		false,
	)

	// Draw dashed border
	page.SetStrokeColor(
		placeholderDiagramBorderR,
		placeholderDiagramBorderG,
		placeholderDiagramBorderB,
	)
	page.SetLineWidth(1.5)
	// Dashed pattern: 5 points on, 3 points off
	page.SetLineDashPattern([]float64{5, 3}, 0)
	page.DrawRectangle(
		x,
		y,
		width,
		height,
		false,
		true,
	)

	// Draw centered label text
	label := "SmartArt Diagram"
	fontSize := 12.0

	// Calculate approximate text width (rough estimate)
	// Helvetica is roughly 0.5 * fontSize for each character on average
	approxCharWidth := fontSize * 0.5
	textWidth := float64(
		len(label),
	) * approxCharWidth

	// Center text horizontally and vertically
	textX := x + (width-textWidth)/2
	textY := y + (height-fontSize)/2

	page.SetFont("Helvetica", fontSize)
	page.SetFillColor(
		placeholderDiagramBorderR,
		placeholderDiagramBorderG,
		placeholderDiagramBorderB,
	)
	page.DrawText(textX, textY, label)

	return nil
}

// RenderDiagramBounds is a helper to render a diagram with bounds from a RenderBounds struct.
// This provides compatibility with the rendering interface used by other DrawingML renderers.
func (r *DiagramRenderer) RenderDiagramBounds(
	bounds RenderBounds,
) error {
	return r.RenderPlaceholder(
		bounds.X,
		bounds.Y,
		bounds.Width,
		bounds.Height,
	)
}
