// Shape generator for E2E visual testing
//
// This file implements shape generation using the goffice drawingml API.
package main

import (
	"fmt"
	"log"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/presentation/elements"
	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

// ShapeGenerator handles shape creation for test cases
type ShapeGenerator struct {
	shapeSpec *framework.ShapeSpec
	verbose   bool
}

// NewShapeGenerator creates a new shape generator
func NewShapeGenerator(
	shapeSpec *framework.ShapeSpec,
	verbose bool,
) *ShapeGenerator {
	return &ShapeGenerator{
		shapeSpec: shapeSpec,
		verbose:   verbose,
	}
}

// Generate creates the shape and adds it to the slide
func (g *ShapeGenerator) Generate(
	slide *elements.Slide,
	pos framework.Position,
	size framework.Size,
) (*elements.Shape, error) {
	if g.verbose {
		log.Printf(
			"      Generating shape: %s",
			g.shapeSpec.Type,
		)
	}

	// Create shape
	shape := slide.AddShape()
	if shape == nil {
		return nil, fmt.Errorf(
			"failed to create shape",
		)
	}

	// Get shape properties
	spPr := shape.GetOrCreateShapeProperties()
	if spPr == nil {
		return nil, fmt.Errorf(
			"failed to get shape properties",
		)
	}

	// Set position and size
	spPr.SetOffset(int(pos.X), int(pos.Y))
	spPr.SetExtents(
		int(size.Width),
		int(size.Height),
	)

	// Set shape geometry (preset shape type)
	shapeType := mapShapeType(g.shapeSpec.Type)
	spPr.SetPresetGeometry(string(shapeType))

	// Apply fill
	if err := g.applyFill(spPr); err != nil {
		return nil, fmt.Errorf(
			"apply fill: %w",
			err,
		)
	}

	// Apply stroke
	if err := g.applyStroke(spPr); err != nil {
		return nil, fmt.Errorf(
			"apply stroke: %w",
			err,
		)
	}

	// Apply effects
	if len(g.shapeSpec.Effects) > 0 {
		if err := g.applyEffects(spPr); err != nil {
			return nil, fmt.Errorf(
				"apply effects: %w",
				err,
			)
		}
	}

	// Add text inside shape if specified
	if g.shapeSpec.Text != nil {
		if g.verbose {
			log.Printf(
				"        Adding text to shape",
			)
		}
		textGen := NewTextGenerator(
			g.shapeSpec.Text,
			g.verbose,
		)
		if err := textGen.ApplyToShape(shape); err != nil {
			return nil, fmt.Errorf(
				"apply text to shape: %w",
				err,
			)
		}
	}

	return shape, nil
}

// applyFill applies fill styling to the shape
func (g *ShapeGenerator) applyFill(
	spPr *elements.ShapeProperties,
) error {
	// Remove existing fills
	g.removeFill(spPr)

	switch g.shapeSpec.Fill.Type {
	case framework.FillTypeNone:
		// No fill
		spPr.SetNoFill()

	case framework.FillTypeSolid:
		// Solid fill
		if g.shapeSpec.Fill.Color == "" {
			return fmt.Errorf(
				"solid fill requires a color",
			)
		}
		spPr.SetSolidFill(g.shapeSpec.Fill.Color)

	case framework.FillTypeGradient:
		// Gradient fill
		if g.shapeSpec.Fill.Gradient == nil {
			return fmt.Errorf(
				"gradient fill requires gradient spec",
			)
		}
		gradFill := g.createGradientFill(
			g.shapeSpec.Fill.Gradient,
		)
		spPr.AppendChild(gradFill)

	case framework.FillTypePattern:
		// Pattern fill
		if g.shapeSpec.Fill.Pattern == nil {
			return fmt.Errorf(
				"pattern fill requires pattern spec",
			)
		}
		pattFill := g.createPatternFill(
			g.shapeSpec.Fill.Pattern,
		)
		spPr.AppendChild(pattFill)

	default:
		return fmt.Errorf(
			"unsupported fill type: %s",
			g.shapeSpec.Fill.Type,
		)
	}

	return nil
}

// removeFill removes any existing fill elements
func (g *ShapeGenerator) removeFill(
	spPr *elements.ShapeProperties,
) {
	if nf := spPr.GetElement("noFill", elements.NamespaceDrawingML); nf != nil {
		spPr.RemoveChild(nf)
	}
	if sf := spPr.GetElement("solidFill", elements.NamespaceDrawingML); sf != nil {
		spPr.RemoveChild(sf)
	}
	if gf := spPr.GetElement("gradFill", elements.NamespaceDrawingML); gf != nil {
		spPr.RemoveChild(gf)
	}
	if pf := spPr.GetElement("pattFill", elements.NamespaceDrawingML); pf != nil {
		spPr.RemoveChild(pf)
	}
}

// createGradientFill creates a gradient fill element
func (g *ShapeGenerator) createGradientFill(
	spec *framework.GradientSpec,
) *drawingml.GradientFill {
	gradFill := drawingml.NewGradientFill()

	// Add gradient stops first
	for _, stop := range spec.Stops {
		// Convert 0-1 to 0-100000 (percentage in 1000ths)
		position := int(stop.Position * 100000)
		gradFill.AddRgbStop(position, stop.Color)
	}

	// Set gradient type
	switch spec.Type {
	case framework.GradientTypeLinear:
		// Linear gradient - angle in 60000ths of a degree
		gradFill.SetLinear(
			spec.Angle*60000,
			false,
		)

	case framework.GradientTypeRadial:
		// Radial gradient (path gradient)
		gradFill.SetPath(
			drawingml.PathShapeCircle,
		)
	}

	return gradFill
}

// createPatternFill creates a pattern fill element
func (g *ShapeGenerator) createPatternFill(
	spec *framework.PatternSpec,
) *drawingml.PatternFill {
	patternType := mapPatternType(spec.Type)
	pattFill := drawingml.NewPatternFill(
		patternType,
	)

	// Set foreground color
	if spec.Foreground != "" {
		pattFill.SetForegroundColor(
			spec.Foreground,
		)
	}

	// Set background color
	if spec.Background != "" {
		pattFill.SetBackgroundColor(
			spec.Background,
		)
	}

	return pattFill
}

// applyStroke applies stroke styling to the shape
func (g *ShapeGenerator) applyStroke(
	spPr *elements.ShapeProperties,
) error {
	// Create line properties
	ln := drawingml.NewLineProperties()

	// Set width (EMUs)
	if g.shapeSpec.Stroke.Width > 0 {
		ln.SetWidth(
			drawingml.EMU(
				g.shapeSpec.Stroke.Width,
			),
		)
	} else {
		// Default stroke width: 1pt = 12700 EMUs
		ln.SetWidth(12700)
	}

	// Set color (solid fill for the line)
	if g.shapeSpec.Stroke.Color != "" {
		ln.SetSolidFill(g.shapeSpec.Stroke.Color)
	}

	// Set dash style
	if g.shapeSpec.Stroke.DashStyle != "" {
		dashStyle := mapDashStyle(
			g.shapeSpec.Stroke.DashStyle,
		)
		ln.SetPresetDash(dashStyle)
	}

	// Set cap style
	if g.shapeSpec.Stroke.Cap != "" {
		capStyle := mapCapStyle(
			g.shapeSpec.Stroke.Cap,
		)
		ln.SetCapType(capStyle)
	}

	// Append line to shape properties
	spPr.AppendChild(ln)

	return nil
}

// applyEffects applies visual effects to the shape
func (g *ShapeGenerator) applyEffects(
	spPr *elements.ShapeProperties,
) error {
	effectList := drawingml.NewEffectList()

	for _, effectSpec := range g.shapeSpec.Effects {
		switch effectSpec.Type {
		case framework.EffectTypeShadow:
			if effectSpec.Shadow == nil {
				continue
			}
			shadow := g.createShadow(
				effectSpec.Shadow,
			)
			effectList.SetOuterShadow(shadow)

		case framework.EffectTypeGlow:
			if effectSpec.Glow == nil {
				continue
			}
			glow := g.createGlow(effectSpec.Glow)
			effectList.SetGlow(glow)

		case framework.EffectTypeReflection:
			if effectSpec.Reflection == nil {
				continue
			}
			reflection := g.createReflection(
				effectSpec.Reflection,
			)
			effectList.SetReflection(reflection)

		case framework.EffectTypeSoftEdge:
			if effectSpec.SoftEdge == nil {
				continue
			}
			softEdge := g.createSoftEdge(
				effectSpec.SoftEdge,
			)
			effectList.SetSoftEdge(softEdge)
		}
	}

	// Append effect list to shape properties
	spPr.AppendChild(effectList)

	return nil
}

// createShadow creates a shadow effect
func (g *ShapeGenerator) createShadow(
	spec *framework.ShadowSpec,
) *drawingml.OuterShadow {
	shadow := drawingml.NewOuterShadow()

	// Set direction (degrees to angle units: deg * 60000)
	shadow.SetDirection(spec.Angle * 60000)

	// Set distance (EMUs)
	shadow.SetDistance(
		drawingml.EMU(spec.Distance),
	)

	// Set blur radius (EMUs)
	shadow.SetBlurRadius(
		drawingml.EMU(spec.BlurRadius),
	)

	// Set color
	if spec.Color != "" {
		shadow.SetRgbColor(spec.Color)
	}

	// Note: SetAlpha is not available in the current API
	// Transparency would need to be set on the color itself

	return shadow
}

// createGlow creates a glow effect
func (g *ShapeGenerator) createGlow(
	spec *framework.GlowSpec,
) *drawingml.Glow {
	glow := drawingml.NewGlow()

	// Set radius (EMUs)
	glow.SetRadius(drawingml.EMU(spec.Radius))

	// Set color
	if spec.Color != "" {
		glow.SetRgbColor(spec.Color)
	}

	return glow
}

// createReflection creates a reflection effect
func (g *ShapeGenerator) createReflection(
	spec *framework.ReflectionSpec,
) *drawingml.Reflection {
	reflection := drawingml.NewReflection()

	// Set blur radius (EMUs)
	reflection.SetBlurRadius(
		drawingml.EMU(spec.BlurRadius),
	)

	// Set start/end opacity (0-1 to 0-100000)
	reflection.SetStartOpacity(
		int(spec.StartOpacity * 100000),
	)
	reflection.SetEndAlpha(
		int(spec.EndOpacity * 100000),
	)

	// Set distance (EMUs)
	reflection.SetDistance(
		drawingml.EMU(spec.Distance),
	)

	// Set direction (degrees to angle units)
	reflection.SetDirection(
		spec.Direction * 60000,
	)

	// Set fade direction (degrees to angle units)
	reflection.SetFadeDirection(
		spec.FadeDirection * 60000,
	)

	// Set start/end position (0-1 to 0-100000)
	reflection.SetStartPosition(
		int(spec.StartPosition * 100000),
	)
	reflection.SetEndPosition(
		int(spec.EndPosition * 100000),
	)

	return reflection
}

// createSoftEdge creates a soft edge effect
func (g *ShapeGenerator) createSoftEdge(
	spec *framework.SoftEdgeSpec,
) *drawingml.SoftEdge {
	// NewSoftEdge requires a radius parameter
	softEdge := drawingml.NewSoftEdge(
		drawingml.EMU(spec.Radius),
	)

	return softEdge
}

// mapShapeType maps framework shape type to drawingml shape type
func mapShapeType(
	shapeType framework.ShapeType,
) drawingml.ShapeTypeValue {
	switch shapeType {
	case framework.ShapeTypeRectangle:
		return drawingml.ShapeTypeRectangle
	case framework.ShapeTypeRoundRect:
		return drawingml.ShapeTypeRoundRectangle
	case framework.ShapeTypeEllipse:
		return drawingml.ShapeTypeEllipse
	case framework.ShapeTypeTriangle:
		return drawingml.ShapeTypeTriangle
	case framework.ShapeTypeDiamond:
		return drawingml.ShapeTypeDiamond
	case framework.ShapeTypePentagon:
		return drawingml.ShapeTypePentagon
	case framework.ShapeTypeHexagon:
		return drawingml.ShapeTypeHexagon
	case framework.ShapeTypeOctagon:
		return drawingml.ShapeTypeOctagon
	case framework.ShapeTypeStar5:
		return drawingml.ShapeTypeStar5
	case framework.ShapeTypeArrowRight:
		return drawingml.ShapeTypeRightArrow
	case framework.ShapeTypeCallout:
		return drawingml.ShapeTypeWedgeRectCallout
	default:
		// Default to rectangle
		return drawingml.ShapeTypeRectangle
	}
}

// mapPatternType maps framework pattern type to drawingml pattern type
func mapPatternType(
	patternType framework.PatternType,
) drawingml.PatternFillValue {
	switch patternType {
	case framework.PatternTypeDots:
		return drawingml.PatternPct10 // Use 10% pattern for dots
	case framework.PatternTypeGrid:
		return drawingml.PatternSmallGrid
	case framework.PatternTypeDiagonalStripe:
		return drawingml.PatternDiagBrick
	case framework.PatternTypeCheckered:
		return drawingml.PatternSmallCheck
	default:
		return drawingml.PatternSmallGrid
	}
}

// mapDashStyle maps framework dash style to drawingml preset dash
func mapDashStyle(
	dashStyle framework.DashStyle,
) drawingml.LineDashValue {
	switch dashStyle {
	case framework.DashStyleSolid:
		return drawingml.LineDashSolid
	case framework.DashStyleDot:
		return drawingml.LineDashDot
	case framework.DashStyleDash:
		return drawingml.LineDashDash
	case framework.DashStyleDashDot:
		return drawingml.LineDashDashDot
	case framework.DashStyleLongDash:
		return drawingml.LineDashLongDash
	default:
		return drawingml.LineDashSolid
	}
}

// mapCapStyle maps framework cap style to drawingml line cap
func mapCapStyle(
	capStyle framework.CapStyle,
) drawingml.LineCapValue {
	switch capStyle {
	case framework.CapStyleFlat:
		return drawingml.LineCapFlat
	case framework.CapStyleRound:
		return drawingml.LineCapRound
	case framework.CapStyleSquare:
		return drawingml.LineCapSquare
	default:
		return drawingml.LineCapFlat
	}
}
