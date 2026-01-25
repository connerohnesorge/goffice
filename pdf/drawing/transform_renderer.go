package drawing

import (
	"math"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
)

const (
	// degreesToRadians is the conversion factor from degrees to radians.
	degreesToRadians = math.Pi / 180.0
)

// TransformRenderer handles rendering transformations (rotation, flip, scale).
type TransformRenderer struct {
	ctx *core.RenderingContext
}

// NewTransformRenderer creates a new transform renderer.
func NewTransformRenderer(
	ctx *core.RenderingContext,
) *TransformRenderer {
	return &TransformRenderer{ctx: ctx}
}

// ApplyTransform applies all transformations for a shape.
func (r *TransformRenderer) ApplyTransform(
	transform *drawingml.Transform2D,
) error {
	if transform == nil {
		return nil
	}

	ooxml := core.NewOOXMLTransform(
		int64(transform.OffsetX()),
		int64(transform.OffsetY()),
		int64(transform.Width()),
		int64(transform.Height()),
	).WithRotation(int32(transform.Rotation())).
		WithFlipH(transform.FlipH()).
		WithFlipV(transform.FlipV())

	if !ooxml.HasTransform() {
		return nil
	}

	r.ctx.Page.SaveGraphicsState()
	r.ctx.Page.Transform(
		core.OOXMLTransformToPDFMatrix(
			ooxml,
			r.ctx.GetPageHeight(),
		),
	)

	return nil
}

// ApplyRotation applies rotation around a center point.
func (r *TransformRenderer) ApplyRotation(
	centerX, centerY, angleDegrees float64,
) {
	// Save state before transformation
	r.ctx.Page.SaveGraphicsState()

	// Translate to center
	matrix := core.IdentityMatrix().
		Translate(centerX, centerY).
		RotateDegrees(angleDegrees).
		Translate(-centerX, -centerY)

	r.ctx.Page.Transform(matrix)
}

// ApplyFlipHorizontal flips content horizontally.
func (r *TransformRenderer) ApplyFlipHorizontal(
	centerX, width float64,
) {
	r.ctx.Page.SaveGraphicsState()

	// Horizontal flip: scale X by -1
	// Matrix: [-1 0 0 1 2*centerX 0]

	matrix := core.IdentityMatrix().
		Translate(centerX+width/2, 0).
		Scale(-1, 1).
		Translate(-(centerX + width/2), 0)
	r.ctx.Page.Transform(matrix)
}

// ApplyFlipVertical flips content vertically.
func (r *TransformRenderer) ApplyFlipVertical(
	centerY, height float64,
) {
	r.ctx.Page.SaveGraphicsState()

	// Vertical flip: scale Y by -1
	// Matrix: [1 0 0 -1 0 2*centerY]

	matrix := core.IdentityMatrix().
		Translate(0, centerY+height/2).
		Scale(1, -1).
		Translate(0, -(centerY + height/2))
	r.ctx.Page.Transform(matrix)
}

// ApplyScale applies scaling transformation.
func (r *TransformRenderer) ApplyScale(
	scaleX, scaleY float64,
) {
	r.ctx.Page.SaveGraphicsState()

	// Scale transformation matrix: [scaleX 0 0 scaleY 0 0]

	matrix := core.IdentityMatrix().Scale(scaleX, scaleY)
	r.ctx.Page.Transform(matrix)
}

// ApplyGroupTransform applies transformations for grouped shapes.
func (*TransformRenderer) ApplyGroupTransform(
	groupTransform, childTransform *drawingml.Transform2D,
) error {
	// Group transformations are relative to the group's coordinate system
	// Child positions are relative to the group

	if groupTransform == nil ||
		childTransform == nil {
		return nil
	}

	// Get group bounds
	groupOffset := groupTransform.Offset()
	groupExtent := groupTransform.Extent()

	// Get child bounds (relative to group)
	childOffset := childTransform.Offset()
	childExtent := childTransform.Extent()

	// Silence unused variable warnings (placeholder implementation)
	_ = groupExtent

	// Convert to absolute coordinates
	absX := drawingml.EmuToPoints(
		groupOffset.X + childOffset.X,
	)
	absY := drawingml.EmuToPoints(
		groupOffset.Y + childOffset.Y,
	)
	absWidth := drawingml.EmuToPoints(
		childExtent.Cx,
	)
	absHeight := drawingml.EmuToPoints(
		childExtent.Cy,
	)

	// Use absolute coordinates for rendering
	_ = absX
	_ = absY
	_ = absWidth
	_ = absHeight

	return nil
}

// RestoreTransform restores the graphics state after transformations.
func (r *TransformRenderer) RestoreTransform() {
	r.ctx.Page.RestoreGraphicsState()
}

// RotatedBounds represents the bounding box after rotation.
type RotatedBounds struct {
	MinX, MinY, MaxX, MaxY float64
}

// BoundingBox represents a rectangular bounding box.
type BoundingBox struct {
	X, Y, Width, Height float64
}

// CalculateRotatedBounds calculates the bounding box after rotation.
func CalculateRotatedBounds(
	box BoundingBox,
	angleDegrees float64,
) RotatedBounds {
	// Calculate the four corners of the rectangle
	corners := [][2]float64{
		{box.X, box.Y},
		{box.X + box.Width, box.Y},
		{box.X + box.Width, box.Y + box.Height},
		{box.X, box.Y + box.Height},
	}

	// Center point
	cx := box.X + box.Width/2
	cy := box.Y + box.Height/2

	// Convert angle to radians
	angleRad := angleDegrees * degreesToRadians
	cos := math.Cos(angleRad)
	sin := math.Sin(angleRad)

	// Initialize bounds
	bounds := RotatedBounds{
		MinX: math.MaxFloat64,
		MinY: math.MaxFloat64,
		MaxX: -math.MaxFloat64,
		MaxY: -math.MaxFloat64,
	}

	// Rotate each corner and find bounding box
	for _, corner := range corners {
		// Translate to origin
		dx := corner[0] - cx
		dy := corner[1] - cy

		// Rotate
		rotX := dx*cos - dy*sin
		rotY := dx*sin + dy*cos

		// Translate back
		newX := rotX + cx
		newY := rotY + cy

		// Update bounds
		if newX < bounds.MinX {
			bounds.MinX = newX
		}
		if newX > bounds.MaxX {
			bounds.MaxX = newX
		}
		if newY < bounds.MinY {
			bounds.MinY = newY
		}
		if newY > bounds.MaxY {
			bounds.MaxY = newY
		}
	}

	return bounds
}
