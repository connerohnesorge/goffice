package drawing

import (
	"fmt"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/drawingml/transform"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

// GroupRenderer handles rendering of grouped shapes with hierarchical transforms.
// Groups use the PDF graphics state stack to maintain transform hierarchy.
type GroupRenderer struct {
	ctx *core.RenderingContext
}

// NewGroupRenderer creates a new group renderer.
func NewGroupRenderer(ctx *core.RenderingContext) *GroupRenderer {
	return &GroupRenderer{ctx: ctx}
}

// RenderGroup renders a group shape and all its children.
// This method:
// 1. Saves the current graphics state
// 2. Applies the group's transformation
// 3. Recursively renders all child shapes
// 4. Restores the graphics state
func (r *GroupRenderer) RenderGroup(group *elements.GroupShape) error {
	if group == nil {
		return fmt.Errorf("group shape is nil")
	}

	// Save graphics state (pushes current transform onto stack)
	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	// Apply group transformation to PDF CTM (Current Transformation Matrix)
	if err := r.applyGroupTransform(group); err != nil {
		return fmt.Errorf("failed to apply group transform: %w", err)
	}

	// Render all children in the transformed coordinate space
	for child := range group.Children() {
		if err := r.renderChild(child); err != nil {
			return fmt.Errorf("failed to render group child: %w", err)
		}
	}

	return nil
}

// applyGroupTransform extracts and applies the group's transform to the PDF CTM.
func (r *GroupRenderer) applyGroupTransform(group *elements.GroupShape) error {
	// Get group properties
	props := group.GroupShapeProperties()
	if props == nil {
		return nil // No properties, no transform
	}

	// Get transform element
	xfrm := props.Transform()
	if xfrm == nil {
		return nil // No transform element
	}

	// Convert DrawingML transform to matrix
	matrix := transform.FromTransform2D(xfrm)

	// Convert to PDF matrix
	// Note: PDF uses different coordinate system (bottom-left origin)
	// DrawingML uses top-left origin, so we need to flip Y
	pageHeight := r.ctx.GetPageHeight()

	// Convert matrix components
	// PDF matrix: [a b c d e f] where point transform is:
	// x' = a*x + c*y + e
	// y' = b*x + d*y + f
	pdfMatrix := core.PDFMatrix{
		A: matrix.A,
		B: -matrix.B, // Flip Y component
		C: -matrix.C, // Flip Y component
		D: matrix.D,
		E: matrix.E,
		F: pageHeight - matrix.F, // Flip Y translation
	}

	// Apply transform to CTM
	r.ctx.Page.Transform(pdfMatrix)

	return nil
}

// renderChild renders a single child element within the group.
// This handles different child types (shapes, pictures, nested groups, etc.)
func (r *GroupRenderer) renderChild(child openxml.Element) error {
	switch c := child.(type) {
	case *elements.Shape:
		return r.renderShape(c)
	case *elements.Picture:
		return r.renderPicture(c)
	case *elements.GroupShape:
		// Recursive rendering of nested groups
		return r.RenderGroup(c)
	case *elements.ConnectionShape:
		return r.renderConnection(c)
	case *elements.GraphicFrame:
		return r.renderGraphicFrame(c)
	default:
		// Unknown or unsupported element type - skip silently
		// This includes non-visual elements like nvGrpSpPr, grpSpPr
		return nil
	}
}

// renderShape renders a shape within the group.
func (r *GroupRenderer) renderShape(shape *elements.Shape) error {
	if shape == nil {
		return nil
	}

	// Get shape properties
	spPr := shape.ShapeProperties()
	if spPr == nil {
		return nil // No visual properties
	}

	// Convert elements.ShapeProperties to drawingml.ShapeProperties
	// The shape properties element should already be a DrawingML ShapeProperties
	dmlSpPr, ok := interface{}(spPr).(*drawingml.ShapeProperties)
	if !ok {
		// Not DrawingML properties, skip
		return nil
	}

	// Use the existing shape renderer
	shapeRenderer := NewShapeRenderer(r.ctx)

	// Create fill and stroke renderers
	fillRenderer := CreateFillRenderer(dmlSpPr)
	strokeRenderer := CreateStrokeRenderer(dmlSpPr)

	return shapeRenderer.RenderShapeWithFill(
		dmlSpPr,
		fillRenderer,
		strokeRenderer,
	)
}

// renderPicture renders a picture within the group.
func (r *GroupRenderer) renderPicture(pic *elements.Picture) error {
	if pic == nil {
		return nil
	}

	// Get picture properties
	// Pictures use BlipFill for the image
	// TODO: Implement picture rendering (this requires image handling)
	// For now, skip silently
	return nil
}

// renderConnection renders a connection shape (connector) within the group.
func (r *GroupRenderer) renderConnection(conn *elements.ConnectionShape) error {
	if conn == nil {
		return nil
	}

	// TODO: Implement connection shape rendering
	// For now, skip silently
	return nil
}

// renderGraphicFrame renders a graphic frame (chart, table, etc.) within the group.
func (r *GroupRenderer) renderGraphicFrame(frame *elements.GraphicFrame) error {
	if frame == nil {
		return nil
	}

	// TODO: Implement graphic frame rendering (charts, tables)
	// For now, skip silently
	return nil
}
