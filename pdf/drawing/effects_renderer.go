// Package drawing provides visual effects rendering for PDF.
package drawing

import (
	"iter"
	"math"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
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

	if r.ctx == nil || r.ctx.Page == nil {
		return nil
	}

	// For basic implementation, render a semi-transparent offset shape.
	shadowColor := applyOpacity(color, opacity)

	passes := 1
	if blurRadius > 0 {
		passes = 3
	}
	for i := 0; i < passes; i++ {
		shift := 0.0
		if passes > 1 {
			factor := float64(i) - float64(passes-1)/2
			shift = factor * blurRadius / float64(passes)
		}

		r.ctx.Page.SaveGraphicsState()
		r.ctx.Page.Transform(
			core.IdentityMatrix().Translate(
				offsetX+shift,
				offsetY+shift,
			),
		)
		r.ctx.Page.SetFillColor(
			shadowColor.R,
			shadowColor.G,
			shadowColor.B,
		)
		r.ctx.Page.WriteContent(path.Fill())
		r.ctx.Page.RestoreGraphicsState()
	}

	return nil
}

// RenderOuterGlow renders an outer glow effect.
func (r *EffectsRenderer) RenderOuterGlow(
	path *PathBuilder,
	glowSize float64,
	color Color,
	opacity float64,
) error {
	if r.ctx == nil || r.ctx.Page == nil {
		return nil
	}

	// Outer glow is approximated with concentric strokes.
	glowColor := applyOpacity(color, opacity)
	steps := 5
	for i := steps; i > 0; i-- {
		stepOpacity := float64(i) / float64(steps)
		r.ctx.Page.SetStrokeColor(
			glowColor.R*stepOpacity,
			glowColor.G*stepOpacity,
			glowColor.B*stepOpacity,
		)
		r.ctx.Page.SetLineWidth(
			glowSize * float64(i) / float64(steps),
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
	if r.ctx == nil || r.ctx.Page == nil {
		return nil
	}

	// Soft edges are approximated with a light stroke.
	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	r.ctx.Page.SetStrokeColor(0.7, 0.7, 0.7)
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
	if r.ctx == nil || r.ctx.Page == nil {
		return nil
	}

	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	if flipVertical {
		r.ctx.Page.Transform(
			core.IdentityMatrix().
				Translate(0, offsetY).
				Scale(1, -1),
		)
	} else {
		r.ctx.Page.Transform(
			core.IdentityMatrix().Translate(0, -offsetY),
		)
	}

	reflectionColor := applyOpacity(
		NewGray(0.7),
		alpha,
	)
	r.ctx.Page.SetFillColor(
		reflectionColor.R,
		reflectionColor.G,
		reflectionColor.B,
	)
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
	if r.ctx == nil || r.ctx.Page == nil {
		return nil
	}

	// Basic implementation without clipping.
	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	shadowColor := applyOpacity(color, opacity*0.5)
	r.ctx.Page.Transform(
		core.IdentityMatrix().Translate(offsetX, offsetY),
	)
	r.ctx.Page.SetFillColor(
		shadowColor.R,
		shadowColor.G,
		shadowColor.B,
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
	shadow := wrapOuterShadow(effectElem)
	if shadow == nil {
		return nil
	}

	distance := drawingml.EmuToPoints(shadow.Distance())
	direction := float64(shadow.Direction()) / 60000.0
	angle := direction * math.Pi / 180.0
	offsetX := math.Cos(angle) * distance
	offsetY := -math.Sin(angle) * distance
	blurRadius := drawingml.EmuToPoints(shadow.BlurRadius())

	color := resolveEffectColor(
		shadow.RgbColor(),
		shadow.SchemeColor(),
	)
	opacity := color.A
	if opacity == 0 {
		opacity = 0.5
	}

	return r.RenderDropShadow(
		path,
		offsetX,
		offsetY,
		blurRadius,
		color,
		opacity,
	)
}

func (r *EffectsRenderer) renderGlowEffect(
	effectElem interface{},
	path *PathBuilder,
) error {
	glow := wrapGlow(effectElem)
	if glow == nil {
		return nil
	}

	radius := drawingml.EmuToPoints(glow.Radius())
	color := resolveEffectColor(
		glow.RgbColor(),
		glow.SchemeColor(),
	)
	opacity := color.A
	if opacity == 0 {
		opacity = 0.6
	}

	return r.RenderOuterGlow(
		path,
		radius,
		color,
		opacity,
	)
}

func (r *EffectsRenderer) renderSoftEdgeEffect(
	effectElem interface{},
	path *PathBuilder,
) error {
	softEdge := wrapSoftEdge(effectElem)
	if softEdge == nil {
		return nil
	}

	radius := drawingml.EmuToPoints(softEdge.Radius())
	if radius == 0 {
		radius = 1
	}

	return r.RenderSoftEdges(path, radius)
}

func (r *EffectsRenderer) renderReflectionEffect(
	effectElem interface{},
	path *PathBuilder,
) error {
	reflection := wrapReflection(effectElem)
	if reflection == nil {
		return nil
	}

	offset := drawingml.EmuToPoints(reflection.Distance())
	alpha := float64(reflection.StartOpacity()) / 100000.0
	if alpha == 0 {
		alpha = 0.5
	}

	return r.RenderReflection(path, -offset, alpha, true)
}

func applyOpacity(
	color Color,
	opacity float64,
) Color {
	if opacity <= 0 {
		return Color{}
	}

	alpha := opacity
	if color.A > 0 {
		alpha *= color.A
	}
	if alpha > 1 {
		alpha = 1
	}

	return Color{
		R: color.R * alpha,
		G: color.G * alpha,
		B: color.B * alpha,
		A: alpha,
	}
}

func resolveEffectColor(
	rgb *drawingml.RgbColor,
	scheme *drawingml.SchemeColor,
) Color {
	if rgb != nil {
		if color := FromRgbColor(rgb); color != nil {
			return color.Resolve(nil)
		}
	}
	if scheme != nil {
		if color := FromSchemeColor(scheme); color != nil {
			return color.Resolve(nil)
		}
	}

	return Black
}

func wrapOuterShadow(
	effectElem interface{},
) *drawingml.OuterShadow {
	switch elem := effectElem.(type) {
	case *drawingml.OuterShadow:
		return elem
	case *openxml.CompositeElementBase:
		return &drawingml.OuterShadow{
			CompositeElementBase: elem,
		}
	default:
		return nil
	}
}

func wrapGlow(
	effectElem interface{},
) *drawingml.Glow {
	switch elem := effectElem.(type) {
	case *drawingml.Glow:
		return elem
	case *openxml.CompositeElementBase:
		return &drawingml.Glow{
			CompositeElementBase: elem,
		}
	default:
		return nil
	}
}

func wrapSoftEdge(
	effectElem interface{},
) *drawingml.SoftEdge {
	switch elem := effectElem.(type) {
	case *drawingml.SoftEdge:
		return elem
	case *openxml.LeafElementBase:
		return &drawingml.SoftEdge{
			LeafElementBase: elem,
		}
	default:
		return nil
	}
}

func wrapReflection(
	effectElem interface{},
) *drawingml.Reflection {
	switch elem := effectElem.(type) {
	case *drawingml.Reflection:
		return elem
	case *openxml.LeafElementBase:
		return &drawingml.Reflection{
			LeafElementBase: elem,
		}
	default:
		return nil
	}
}
