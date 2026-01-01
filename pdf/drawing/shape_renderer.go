// Package drawing provides DrawingML rendering capabilities for PDF generation.
package drawing

import (
	"fmt"
	"math"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
)

// ShapeRenderer handles rendering of complete DrawingML shapes with all properties.
type ShapeRenderer struct {
	ctx *core.RenderingContext
}

// NewShapeRenderer creates a new shape renderer.
func NewShapeRenderer(
	ctx *core.RenderingContext,
) *ShapeRenderer {
	return &ShapeRenderer{ctx: ctx}
}

// RenderShapeWithFill renders a shape with fill, stroke, effects, and text.
func (r *ShapeRenderer) RenderShapeWithFill(
	sp *drawingml.ShapeProperties,
	fill FillRenderer,
	stroke StrokeRenderer,
) error {
	if sp == nil {
		return fmt.Errorf(
			"shape properties is nil",
		)
	}

	// Save graphics state
	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	// Apply transform (rotation, flip)
	if err := r.applyTransform(sp); err != nil {
		return err
	}

	// Create the shape path
	path := NewPathBuilder()
	if err := r.createShapePath(sp, path); err != nil {
		return err
	}

	effectsRenderer := NewEffectsRenderer(r.ctx)
	if err := effectsRenderer.RenderAllEffects(
		sp,
		path,
	); err != nil {
		return fmt.Errorf(
			"failed to render effects: %w",
			err,
		)
	}

	// Apply fill
	if fill != nil {
		if err := fill.Apply(r.ctx, path); err != nil {
			return fmt.Errorf(
				"failed to apply fill: %w",
				err,
			)
		}
	}

	// Apply stroke/outline
	if stroke != nil {
		if err := stroke.Apply(r.ctx, path); err != nil {
			return fmt.Errorf(
				"failed to apply stroke: %w",
				err,
			)
		}
	}

	return nil
}

// applyTransform applies rotation and flip transformations.
func (r *ShapeRenderer) applyTransform(
	sp *drawingml.ShapeProperties,
) error {
	transform := sp.Transform()
	if transform == nil {
		return nil
	}

	ooxml := toOOXMLTransform(transform)
	if !ooxml.HasTransform() {
		return nil
	}

	matrix := core.OOXMLTransformToPDFMatrix(
		ooxml,
		r.ctx.GetPageHeight(),
	)
	r.ctx.Page.Transform(matrix)

	return nil
}

// createShapePath creates the geometric path for a shape.
func (r *ShapeRenderer) createShapePath(
	sp *drawingml.ShapeProperties,
	path *PathBuilder,
) error {
	transform := sp.Transform()
	if transform == nil {
		return fmt.Errorf(
			"shape has no transform",
		)
	}

	x, y, width, height, err := r.shapeBounds(
		transform,
	)
	if err != nil {
		return err
	}

	// Handle different geometry types
	if presetGeom := sp.PresetGeometry(); presetGeom != nil {
		return renderPresetGeometry(
			presetGeom,
			path,
			x,
			y,
			width,
			height,
		)
	}

	if customGeom := sp.CustomGeometry(); customGeom != nil {
		return renderCustomGeometry(
			customGeom,
			path,
			x,
			y,
			width,
			height,
		)
	}

	// Default to rectangle
	path.Rectangle(x, y, width, height)

	return nil
}

func (r *ShapeRenderer) shapeBounds(
	transform *drawingml.Transform2D,
) (x, y, width, height float64, err error) {
	if transform == nil {
		return 0, 0, 0, 0, fmt.Errorf(
			"shape has no transform",
		)
	}

	offset := transform.Offset()
	extent := transform.Extent()

	width = drawingml.EmuToPoints(extent.Cx)
	height = drawingml.EmuToPoints(extent.Cy)
	x = drawingml.EmuToPoints(offset.X)
	yTop := drawingml.EmuToPoints(offset.Y)
	y = r.ctx.GetPageHeight() - yTop - height

	return x, y, width, height, nil
}

func toOOXMLTransform(
	transform *drawingml.Transform2D,
) core.OOXMLTransform {
	offset := transform.Offset()
	extent := transform.Extent()

	ooxml := core.NewOOXMLTransform(
		int64(offset.X),
		int64(offset.Y),
		int64(extent.Cx),
		int64(extent.Cy),
	)
	ooxml = ooxml.WithRotation(
		int32(transform.Rotation()),
	)
	ooxml = ooxml.WithFlipH(transform.FlipH())
	ooxml = ooxml.WithFlipV(transform.FlipV())

	return ooxml
}

// Preset shape rendering functions (Section 4.1)

func renderPresetGeometry(
	geom *drawingml.PresetGeometry,
	path *PathBuilder,
	x, y, width, height float64,
) error {
	shapeType := geom.Preset()

	//nolint:exhaustive // Many shape types not yet implemented, fall back to rectangle
	switch shapeType {
	case drawingml.ShapeTypeRectangle:
		path.Rectangle(x, y, width, height)
	case drawingml.ShapeTypeRoundRectangle:
		renderRoundRectangle(
			path,
			x,
			y,
			width,
			height,
			math.Min(width, height)*0.1,
		)
	case drawingml.ShapeTypeEllipse:
		renderEllipse(path, x, y, width, height)
	case drawingml.ShapeTypeTriangle:
		renderTriangle(path, x, y, width, height)
	case drawingml.ShapeTypeDiamond:
		renderDiamond(path, x, y, width, height)
	case drawingml.ShapeTypePentagon:
		renderPolygon(
			path,
			x,
			y,
			width,
			height,
			5,
		)
	case drawingml.ShapeTypeHexagon:
		renderPolygon(
			path,
			x,
			y,
			width,
			height,
			6,
		)
	case drawingml.ShapeTypeOctagon:
		renderPolygon(
			path,
			x,
			y,
			width,
			height,
			8,
		)
	case drawingml.ShapeTypeStar4:
		renderStar(path, x, y, width, height, 4)
	case drawingml.ShapeTypeStar5:
		renderStar(path, x, y, width, height, 5)
	case drawingml.ShapeTypeStar6:
		renderStar(path, x, y, width, height, 6)
	case drawingml.ShapeTypeStar8:
		renderStar(path, x, y, width, height, 8)
	case drawingml.ShapeTypeStar10:
		renderStar(path, x, y, width, height, 10)
	case drawingml.ShapeTypeStar12:
		renderStar(path, x, y, width, height, 12)
	case drawingml.ShapeTypeRightArrow:
		renderArrow(path, x, y, width, height, 0)
	case drawingml.ShapeTypeLeftArrow:
		renderArrow(
			path,
			x,
			y,
			width,
			height,
			180,
		)
	case drawingml.ShapeTypeUpArrow:
		renderArrow(
			path,
			x,
			y,
			width,
			height,
			270,
		)
	case drawingml.ShapeTypeDownArrow:
		renderArrow(path, x, y, width, height, 90)
	case drawingml.ShapeTypeLine:
		path.MoveTo(x, y)
		path.LineTo(x+width, y+height)
	default:
		// Fallback to rectangle for unsupported shapes
		path.Rectangle(x, y, width, height)
	}

	return nil
}

func renderRoundRectangle(
	path *PathBuilder,
	x, y, width, height, radius float64,
) {
	// Use PathBuilder's built-in RoundedRect method
	path.RoundedRect(x, y, width, height, radius)
}

func renderEllipse(
	path *PathBuilder,
	x, y, width, height float64,
) {
	cx := x + width/2
	cy := y + height/2
	rx := width / 2
	ry := height / 2

	// Use PathBuilder's built-in Ellipse method
	path.Ellipse(cx, cy, rx, ry)
}

func renderTriangle(
	path *PathBuilder,
	x, y, width, height float64,
) {
	path.MoveTo(x+width/2, y+height)
	path.LineTo(x+width, y)
	path.LineTo(x, y)
	path.ClosePath()
}

func renderDiamond(
	path *PathBuilder,
	x, y, width, height float64,
) {
	cx := x + width/2
	cy := y + height/2

	path.MoveTo(cx, y)
	path.LineTo(x+width, cy)
	path.LineTo(cx, y+height)
	path.LineTo(x, cy)
	path.ClosePath()
}

func renderPolygon(
	path *PathBuilder,
	x, y, width, height float64,
	sides int,
) {
	cx := x + width/2
	cy := y + height/2
	radius := math.Min(width, height) / 2

	// Use PathBuilder's built-in RegularPolygon method
	// Start at 90 degrees to put the first vertex at the top in PDF coordinates.
	path.RegularPolygon(
		cx,
		cy,
		radius,
		sides,
		90,
	)
}

func renderStar(
	path *PathBuilder,
	x, y, width, height float64,
	points int,
) {
	cx := x + width/2
	cy := y + height/2
	outerRadius := math.Min(width, height) / 2
	innerRadius := outerRadius * 0.382 // Golden ratio approximation

	// Use PathBuilder's built-in Star method
	// Start at 90 degrees to have the top point in PDF coordinates.
	path.Star(
		cx,
		cy,
		outerRadius,
		innerRadius,
		points,
		90,
	)
}

func renderArrow(
	path *PathBuilder,
	x, y, width, height float64,
	angle float64,
) {
	// Simple right-pointing arrow
	headWidth := width * 0.3
	shaftHeight := height * 0.4

	cy := y + height/2

	path.MoveTo(x, cy-shaftHeight/2)
	path.LineTo(
		x+width-headWidth,
		cy-shaftHeight/2,
	)
	path.LineTo(x+width-headWidth, y)
	path.LineTo(x+width, cy)
	path.LineTo(x+width-headWidth, y+height)
	path.LineTo(
		x+width-headWidth,
		cy+shaftHeight/2,
	)
	path.LineTo(x, cy+shaftHeight/2)
	path.ClosePath()

	// Rotation would be applied via transform matrix
	_ = angle
}

func renderCustomGeometry(
	geom *drawingml.CustomGeometry,
	path *PathBuilder,
	x, y, width, height float64,
) error {
	pathList := geom.PathList()
	if pathList == nil {
		return fmt.Errorf(
			"custom geometry has no path list",
		)
	}

	// Parse and render custom paths
	// This would iterate through path commands and convert them to PDF paths
	// For now, this is a basic stub

	return nil
}
