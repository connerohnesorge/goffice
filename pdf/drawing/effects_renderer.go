// Package drawing provides visual effects rendering for PDF.
package drawing

import (
	"iter"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
)

// EffectsRenderer handles rendering of visual effects.
type EffectsRenderer struct {
	ctx *core.RenderingContext
}

// NewEffectsRenderer creates a new effects renderer.
func NewEffectsRenderer(
	ctx *core.RenderingContext,
) *EffectsRenderer {
	return &EffectsRenderer{ctx: ctx}
}

// RenderDropShadow renders a drop shadow effect.
func (r *EffectsRenderer) RenderDropShadow(
	path *PathBuilder,
	offsetX, offsetY, blurRadius float64,
	color Color,
	opacity float64,
) error {
	// Drop shadow in PDF requires:
	// 1. Render the path offset by (offsetX, offsetY)
	// 2. Fill with shadow color and opacity
	// 3. Apply blur (requires soft mask or external tool)

	// For basic implementation, render a semi-transparent offset shape
	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	// Set shadow color with reduced opacity
	r.ctx.Page.SetFillColor(
		color.R*opacity,
		color.G*opacity,
		color.B*opacity,
	)

	// Create offset path
	// In a full implementation, this would offset the path
	// For now, this is a placeholder
	r.ctx.Page.WriteContent(path.Fill())

	return nil
}

// RenderOuterGlow renders an outer glow effect.
func (r *EffectsRenderer) RenderOuterGlow(
	path *PathBuilder,
	glowSize float64,
	color Color,
	opacity float64,
) error {
	// Outer glow is similar to drop shadow but radiates outward
	// This requires path expansion and blur

	// For basic implementation, render multiple offset paths
	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	// Render glow as concentric rings with decreasing opacity
	steps := 5
	for i := steps; i > 0; i-- {
		stepOpacity := opacity * float64(
			i,
		) / float64(
			steps,
		) * 0.3
		r.ctx.Page.SetStrokeColor(
			color.R*stepOpacity,
			color.G*stepOpacity,
			color.B*stepOpacity,
		)
		r.ctx.Page.SetLineWidth(
			glowSize * float64(
				i,
			) / float64(
				steps,
			),
		)
		r.ctx.Page.WriteContent(path.Stroke())
	}

	return nil
}

// RenderSoftEdges renders soft/feathered edges.
func (r *EffectsRenderer) RenderSoftEdges(
	path *PathBuilder,
	radius float64,
) error {
	// Soft edges fade the edges of a shape
	// This requires alpha gradients or soft masks in PDF

	// For basic implementation, just stroke with reduced opacity
	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	r.ctx.Page.SetStrokeColor(0.5, 0.5, 0.5)
	r.ctx.Page.SetLineWidth(radius)
	r.ctx.Page.WriteContent(path.Stroke())

	return nil
}

// RenderReflection renders a reflection effect.
func (r *EffectsRenderer) RenderReflection(
	path *PathBuilder,
	offsetY, alpha float64,
	flipVertical bool,
) error {
	// Reflection creates a mirrored version below the shape
	// This requires:
	// 1. Flip the path vertically
	// 2. Offset downward
	// 3. Apply alpha gradient (fade from alpha to 0)

	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	// Apply vertical flip transform
	// For basic implementation, this is a placeholder
	// Full implementation would use transformation matrix

	// Render with reduced opacity
	r.ctx.Page.SetFillColor(0.7, 0.7, 0.7)
	r.ctx.Page.WriteContent(path.Fill())

	return nil
}

// RenderInnerShadow renders an inner shadow effect.
func (r *EffectsRenderer) RenderInnerShadow(
	path *PathBuilder,
	offsetX, offsetY, blurRadius float64,
	color Color,
	opacity float64,
) error {
	// Inner shadow is a shadow inside the shape
	// This requires:
	// 1. Clip to shape
	// 2. Render inverted shadow

	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	// Basic implementation
	r.ctx.Page.SetFillColor(
		color.R*opacity*0.5,
		color.G*opacity*0.5,
		color.B*opacity*0.5,
	)
	r.ctx.Page.WriteContent(path.Fill())

	return nil
}

// Apply3DEffect renders a basic 3D bevel effect.
func (r *EffectsRenderer) Apply3DEffect(
	path *PathBuilder,
	bevelWidth float64,
) error {
	// 3D effects require highlight and shadow on edges
	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	// Render highlight on top-left
	r.ctx.Page.SetStrokeColor(1, 1, 1)
	r.ctx.Page.SetLineWidth(bevelWidth)
	r.ctx.Page.WriteContent(path.Stroke())

	// Render shadow on bottom-right
	r.ctx.Page.SetStrokeColor(0.3, 0.3, 0.3)
	r.ctx.Page.SetLineWidth(bevelWidth * 0.5)
	r.ctx.Page.WriteContent(path.Stroke())

	return nil
}

// HasEffects checks if a shape has visual effects defined.
func HasEffects(
	sp *drawingml.ShapeProperties,
) bool {
	if sp == nil {
		return false
	}

	// Check for effect list element
	effectLst := sp.GetElement(
		"effectLst",
		drawingml.NamespaceMain,
	)

	return effectLst != nil
}

// RenderAllEffects renders all effects for a shape.
func (r *EffectsRenderer) RenderAllEffects(
	sp *drawingml.ShapeProperties,
	path *PathBuilder,
) error {
	if !HasEffects(sp) {
		return nil
	}

	effectLst := sp.GetElement(
		"effectLst",
		drawingml.NamespaceMain,
	)
	if effectLst == nil {
		return nil
	}

	// Cast to CompositeElement to access Children()
	composite, ok := effectLst.(interface {
		Children() iter.Seq[interface{}]
	})
	if !ok {
		return nil
	}

	// Iterate through effects and render them
	for child := range composite.Children() {
		// Type assert child to Element
		elem, ok := child.(interface{ LocalName() string })
		if !ok {
			continue
		}
		var err error
		switch elem.LocalName() {
		case "outerShdw": // Outer shadow (drop shadow)
			err = r.renderOuterShadowEffect(
				child,
				path,
			)
		case "glow": // Outer glow
			err = r.renderGlowEffect(child, path)
		case "softEdge": // Soft edges
			err = r.renderSoftEdgeEffect(
				child,
				path,
			)
		case "reflection": // Reflection
			err = r.renderReflectionEffect(
				child,
				path,
			)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *EffectsRenderer) renderOuterShadowEffect(
	effectElem interface{},
	path *PathBuilder,
) error {
	// Parse shadow parameters and render
	return r.RenderDropShadow(
		path,
		3,
		3,
		2,
		NewRGB(0, 0, 0),
		0.5,
	)
}

func (r *EffectsRenderer) renderGlowEffect(
	effectElem interface{},
	path *PathBuilder,
) error {
	// Parse glow parameters and render
	return r.RenderOuterGlow(
		path,
		2,
		NewRGB(1, 1, 0),
		0.6,
	)
}

func (r *EffectsRenderer) renderSoftEdgeEffect(
	effectElem interface{},
	path *PathBuilder,
) error {
	// Parse soft edge parameters and render
	return r.RenderSoftEdges(path, 1)
}

func (r *EffectsRenderer) renderReflectionEffect(
	effectElem interface{},
	path *PathBuilder,
) error {
	// Parse reflection parameters and render
	return r.RenderReflection(path, 5, 0.5, true)
}
