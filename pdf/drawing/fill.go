package drawing

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// FillType represents the type of fill operation.
type FillType int

const (
	// FillTypeNone indicates no fill.
	FillTypeNone FillType = iota
	// FillTypeSolid indicates a solid color fill.
	FillTypeSolid
	// FillTypeLinearGradient indicates a linear gradient fill.
	FillTypeLinearGradient
	// FillTypeRadialGradient indicates a radial gradient fill.
	FillTypeRadialGradient
	// FillTypePattern indicates a pattern fill.
	FillTypePattern
)

// Fill represents a fill operation for shapes in PDF.
// Fills can be solid colors, gradients, or patterns.
type Fill interface {
	// Type returns the type of fill.
	Type() FillType
	// IsNone returns true if this represents no fill.
	IsNone() bool
	// ContentStream returns PDF operators to apply this fill.
	// The bounds parameter defines the area to fill.
	ContentStream(bounds Rect) string
}

// Rect represents a bounding rectangle.
type Rect struct {
	X, Y, Width, Height float64
}

// NewRect creates a new rectangle.
func NewRect(x, y, width, height float64) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

// Center returns the center point of the rectangle.
func (r Rect) Center() (x, y float64) {
	return r.X + r.Width/2, r.Y + r.Height/2
}

// Right returns the right edge x-coordinate.
func (r Rect) Right() float64 {
	return r.X + r.Width
}

// Top returns the top edge y-coordinate.
func (r Rect) Top() float64 {
	return r.Y + r.Height
}

// NoFill represents no fill.
type NoFill struct{}

// Type returns FillTypeNone.
func (f NoFill) Type() FillType { return FillTypeNone }

// IsNone returns true.
func (f NoFill) IsNone() bool { return true }

// ContentStream returns an empty string (no fill operators).
func (f NoFill) ContentStream(
	bounds Rect,
) string {
	return ""
}

// SolidFill represents a solid color fill.
type SolidFill struct {
	Color Color
}

// NewSolidFill creates a new solid fill with the specified color.
func NewSolidFill(color Color) *SolidFill {
	return &SolidFill{Color: color}
}

// Type returns FillTypeSolid.
func (f *SolidFill) Type() FillType { return FillTypeSolid }

// IsNone returns true if the color is fully transparent.
func (f *SolidFill) IsNone() bool {
	return f.Color.IsTransparent()
}

// ContentStream returns the PDF operators to set this fill color.
func (f *SolidFill) ContentStream(
	bounds Rect,
) string {
	if f.IsNone() {
		return ""
	}

	return f.Color.SetFillRGB()
}

// GradientStop represents a color stop in a gradient.
type GradientStop struct {
	// Position is the position of this stop along the gradient (0-1).
	Position float64
	// Color is the color at this position.
	Color Color
}

// NewGradientStop creates a new gradient stop.
func NewGradientStop(
	position float64,
	color Color,
) GradientStop {
	return GradientStop{
		Position: clamp01(position),
		Color:    color,
	}
}

// GradientExtend specifies how a gradient extends beyond its defined area.
type GradientExtend int

const (
	// GradientExtendPad extends the gradient by padding with the end colors.
	GradientExtendPad GradientExtend = iota
	// GradientExtendRepeat repeats the gradient.
	GradientExtendRepeat
	// GradientExtendReflect reflects the gradient.
	GradientExtendReflect
)

// LinearGradient represents a linear gradient fill.
type LinearGradient struct {
	// Angle is the angle of the gradient in degrees (0 = left to right).
	Angle float64
	// Stops is the list of color stops.
	Stops []GradientStop
	// Extend specifies how the gradient extends beyond its bounds.
	Extend GradientExtend
	// Opacity is the overall opacity of the gradient (0-1).
	Opacity float64
}

// NewLinearGradient creates a new linear gradient with two colors.
func NewLinearGradient(
	angle float64,
	start, end Color,
) *LinearGradient {
	return &LinearGradient{
		Angle: angle,
		Stops: []GradientStop{
			{Position: 0, Color: start},
			{Position: 1, Color: end},
		},
		Extend:  GradientExtendPad,
		Opacity: 1.0,
	}
}

// NewLinearGradientWithStops creates a new linear gradient with multiple stops.
func NewLinearGradientWithStops(
	angle float64,
	stops []GradientStop,
) *LinearGradient {
	// Sort stops by position
	sortedStops := make(
		[]GradientStop,
		len(stops),
	)
	copy(sortedStops, stops)
	sort.Slice(sortedStops, func(i, j int) bool {
		return sortedStops[i].Position < sortedStops[j].Position
	})

	return &LinearGradient{
		Angle:   angle,
		Stops:   sortedStops,
		Extend:  GradientExtendPad,
		Opacity: 1.0,
	}
}

// AddStop adds a color stop to the gradient.
func (g *LinearGradient) AddStop(
	position float64,
	color Color,
) *LinearGradient {
	g.Stops = append(
		g.Stops,
		NewGradientStop(position, color),
	)
	// Keep stops sorted
	sort.Slice(g.Stops, func(i, j int) bool {
		return g.Stops[i].Position < g.Stops[j].Position
	})

	return g
}

// SetExtend sets the gradient extend mode.
func (g *LinearGradient) SetExtend(
	extend GradientExtend,
) *LinearGradient {
	g.Extend = extend

	return g
}

// SetOpacity sets the overall gradient opacity.
func (g *LinearGradient) SetOpacity(
	opacity float64,
) *LinearGradient {
	g.Opacity = clamp01(opacity)

	return g
}

// Type returns FillTypeLinearGradient.
func (g *LinearGradient) Type() FillType { return FillTypeLinearGradient }

// IsNone returns false (gradients are never "none").
func (g *LinearGradient) IsNone() bool { return len(g.Stops) == 0 }

// StartPoint returns the start point of the gradient line for the given bounds.
func (g *LinearGradient) StartPoint(
	bounds Rect,
) (x, y float64) {
	cx, cy := bounds.Center()
	radians := g.Angle * math.Pi / 180

	// Calculate the diagonal length to ensure gradient covers entire rectangle
	diagonal := math.Sqrt(
		bounds.Width*bounds.Width+bounds.Height*bounds.Height,
	) / 2

	return cx - diagonal*math.Cos(
			radians,
		), cy - diagonal*math.Sin(
			radians,
		)
}

// EndPoint returns the end point of the gradient line for the given bounds.
func (g *LinearGradient) EndPoint(
	bounds Rect,
) (x, y float64) {
	cx, cy := bounds.Center()
	radians := g.Angle * math.Pi / 180
	diagonal := math.Sqrt(
		bounds.Width*bounds.Width+bounds.Height*bounds.Height,
	) / 2

	return cx + diagonal*math.Cos(
			radians,
		), cy + diagonal*math.Sin(
			radians,
		)
}

// ColorAt returns the interpolated color at the given position (0-1).
func (g *LinearGradient) ColorAt(
	t float64,
) Color {
	if len(g.Stops) == 0 {
		return Black
	}
	if len(g.Stops) == 1 {
		return g.Stops[0].Color
	}

	t = clamp01(t)

	// Find the two stops to interpolate between
	var prevStop, nextStop GradientStop
	prevStop = g.Stops[0]
	nextStop = g.Stops[len(g.Stops)-1]

	for i := range len(g.Stops) - 1 {
		if t >= g.Stops[i].Position &&
			t <= g.Stops[i+1].Position {
			prevStop = g.Stops[i]
			nextStop = g.Stops[i+1]

			break
		}
	}

	// Handle edge cases
	if t <= prevStop.Position {
		return prevStop.Color
	}
	if t >= nextStop.Position {
		return nextStop.Color
	}

	// Interpolate between stops
	range01 := nextStop.Position - prevStop.Position
	if range01 <= 0 {
		return prevStop.Color
	}
	amount := (t - prevStop.Position) / range01

	return prevStop.Color.Blend(
		nextStop.Color,
		amount,
	)
}

// ContentStream generates PDF operators for the linear gradient.
// Note: True PDF gradients require shading patterns. This provides a
// simplified approximation using multiple filled rectangles.
func (g *LinearGradient) ContentStream(
	bounds Rect,
) string {
	if g.IsNone() {
		return ""
	}

	var sb strings.Builder

	// For a true gradient, we would use PDF shading patterns.
	// Here we provide a simplified version that approximates the gradient
	// with a series of filled rectangles.
	// For production use, this should create a proper PDF shading dictionary.

	// Use the color at position 0.5 as an approximation
	// (real implementation would create shading pattern)
	midColor := g.ColorAt(0.5)
	if g.Opacity < 1.0 {
		midColor = midColor.WithAlpha(
			midColor.A * g.Opacity,
		)
	}
	sb.WriteString(midColor.SetFillRGB())

	return sb.String()
}

// ShadingDict returns the PDF shading dictionary parameters for this gradient.
// This can be used to create a proper PDF shading pattern.
func (g *LinearGradient) ShadingDict(
	bounds Rect,
) map[string]interface{} {
	x0, y0 := g.StartPoint(bounds)
	x1, y1 := g.EndPoint(bounds)

	// Build function for color interpolation
	// PDF uses a stitching function for multi-stop gradients

	return map[string]interface{}{
		"ShadingType": 2, // Axial shading
		"ColorSpace":  "/DeviceRGB",
		"Coords":      []float64{x0, y0, x1, y1},
		"Extend": []bool{
			true,
			true,
		}, // Extend at both ends
		"Function": g.colorFunction(),
	}
}

// colorFunction returns the PDF function specification for the gradient colors.
func (g *LinearGradient) colorFunction() map[string]interface{} {
	if len(g.Stops) <= 2 {
		// Simple linear interpolation
		startColor := Black
		endColor := Black
		if len(g.Stops) > 0 {
			startColor = g.Stops[0].Color
		}
		if len(g.Stops) > 1 {
			endColor = g.Stops[len(g.Stops)-1].Color
		}

		return map[string]interface{}{
			"FunctionType": 2, // Exponential interpolation
			"Domain":       []float64{0, 1},
			"C0": []float64{
				startColor.R,
				startColor.G,
				startColor.B,
			},
			"C1": []float64{
				endColor.R,
				endColor.G,
				endColor.B,
			},
			"N": 1, // Linear interpolation
		}
	}

	// Multi-stop gradient uses stitching function
	numSegments := len(g.Stops) - 1
	bounds := make([]float64, 0, numSegments-1)
	encode := make([]float64, 0, 2*numSegments)
	functions := make(
		[]map[string]interface{},
		0,
		numSegments,
	)

	for i := range numSegments {
		c0 := g.Stops[i].Color
		c1 := g.Stops[i+1].Color

		functions = append(
			functions,
			map[string]interface{}{
				"FunctionType": 2,
				"Domain":       []float64{0, 1},
				"C0": []float64{
					c0.R,
					c0.G,
					c0.B,
				},
				"C1": []float64{
					c1.R,
					c1.G,
					c1.B,
				},
				"N": 1,
			},
		)

		if i < len(g.Stops)-2 {
			bounds = append(
				bounds,
				g.Stops[i+1].Position,
			)
		}
		encode = append(encode, 0, 1)
	}

	return map[string]interface{}{
		"FunctionType": 3, // Stitching function
		"Domain":       []float64{0, 1},
		"Functions":    functions,
		"Bounds":       bounds,
		"Encode":       encode,
	}
}

// RadialGradient represents a radial gradient fill.
type RadialGradient struct {
	// CenterX, CenterY are the center of the gradient (0-1 relative to bounds).
	CenterX, CenterY float64
	// FocusX, FocusY are the focus point (0-1 relative to bounds).
	// For a simple circular gradient, focus equals center.
	FocusX, FocusY float64
	// Radius is the radius of the gradient (0-1 relative to bounds diagonal).
	Radius float64
	// Stops is the list of color stops.
	Stops []GradientStop
	// Extend specifies how the gradient extends beyond its bounds.
	Extend GradientExtend
	// Opacity is the overall opacity of the gradient (0-1).
	Opacity float64
}

// NewRadialGradient creates a new radial gradient centered in the bounds.
func NewRadialGradient(
	start, end Color,
) *RadialGradient {
	return &RadialGradient{
		CenterX: 0.5,
		CenterY: 0.5,
		FocusX:  0.5,
		FocusY:  0.5,
		Radius:  0.5,
		Stops: []GradientStop{
			{Position: 0, Color: start},
			{Position: 1, Color: end},
		},
		Extend:  GradientExtendPad,
		Opacity: 1.0,
	}
}

// NewRadialGradientWithStops creates a new radial gradient with multiple stops.
func NewRadialGradientWithStops(
	stops []GradientStop,
) *RadialGradient {
	sortedStops := make(
		[]GradientStop,
		len(stops),
	)
	copy(sortedStops, stops)
	sort.Slice(sortedStops, func(i, j int) bool {
		return sortedStops[i].Position < sortedStops[j].Position
	})

	return &RadialGradient{
		CenterX: 0.5,
		CenterY: 0.5,
		FocusX:  0.5,
		FocusY:  0.5,
		Radius:  0.5,
		Stops:   sortedStops,
		Extend:  GradientExtendPad,
		Opacity: 1.0,
	}
}

// SetCenter sets the center of the gradient (0-1 relative to bounds).
func (g *RadialGradient) SetCenter(
	x, y float64,
) *RadialGradient {
	g.CenterX = x
	g.CenterY = y

	return g
}

// SetFocus sets the focus point of the gradient (0-1 relative to bounds).
func (g *RadialGradient) SetFocus(
	x, y float64,
) *RadialGradient {
	g.FocusX = x
	g.FocusY = y

	return g
}

// SetRadius sets the radius of the gradient (0-1 relative to bounds).
func (g *RadialGradient) SetRadius(
	r float64,
) *RadialGradient {
	g.Radius = r

	return g
}

// AddStop adds a color stop to the gradient.
func (g *RadialGradient) AddStop(
	position float64,
	color Color,
) *RadialGradient {
	g.Stops = append(
		g.Stops,
		NewGradientStop(position, color),
	)
	sort.Slice(g.Stops, func(i, j int) bool {
		return g.Stops[i].Position < g.Stops[j].Position
	})

	return g
}

// SetExtend sets the gradient extend mode.
func (g *RadialGradient) SetExtend(
	extend GradientExtend,
) *RadialGradient {
	g.Extend = extend

	return g
}

// SetOpacity sets the overall gradient opacity.
func (g *RadialGradient) SetOpacity(
	opacity float64,
) *RadialGradient {
	g.Opacity = clamp01(opacity)

	return g
}

// Type returns FillTypeRadialGradient.
func (g *RadialGradient) Type() FillType { return FillTypeRadialGradient }

// IsNone returns false.
func (g *RadialGradient) IsNone() bool { return len(g.Stops) == 0 }

// CenterPoint returns the absolute center point for the given bounds.
func (g *RadialGradient) CenterPoint(
	bounds Rect,
) (x, y float64) {
	return bounds.X + bounds.Width*g.CenterX, bounds.Y + bounds.Height*g.CenterY
}

// FocusPoint returns the absolute focus point for the given bounds.
func (g *RadialGradient) FocusPoint(
	bounds Rect,
) (x, y float64) {
	return bounds.X + bounds.Width*g.FocusX, bounds.Y + bounds.Height*g.FocusY
}

// RadiusValue returns the absolute radius for the given bounds.
func (g *RadialGradient) RadiusValue(
	bounds Rect,
) float64 {
	diagonal := math.Sqrt(
		bounds.Width*bounds.Width + bounds.Height*bounds.Height,
	)

	return diagonal * g.Radius
}

// ColorAt returns the interpolated color at the given position (0-1).
func (g *RadialGradient) ColorAt(
	t float64,
) Color {
	if len(g.Stops) == 0 {
		return Black
	}
	if len(g.Stops) == 1 {
		return g.Stops[0].Color
	}

	t = clamp01(t)

	var prevStop, nextStop GradientStop
	prevStop = g.Stops[0]
	nextStop = g.Stops[len(g.Stops)-1]

	for i := range len(g.Stops) - 1 {
		if t >= g.Stops[i].Position &&
			t <= g.Stops[i+1].Position {
			prevStop = g.Stops[i]
			nextStop = g.Stops[i+1]

			break
		}
	}

	if t <= prevStop.Position {
		return prevStop.Color
	}
	if t >= nextStop.Position {
		return nextStop.Color
	}

	range01 := nextStop.Position - prevStop.Position
	if range01 <= 0 {
		return prevStop.Color
	}
	amount := (t - prevStop.Position) / range01

	return prevStop.Color.Blend(
		nextStop.Color,
		amount,
	)
}

// ContentStream generates PDF operators for the radial gradient.
func (g *RadialGradient) ContentStream(
	bounds Rect,
) string {
	if g.IsNone() {
		return ""
	}

	// Simplified approximation - use center color
	midColor := g.ColorAt(0.5)
	if g.Opacity < 1.0 {
		midColor = midColor.WithAlpha(
			midColor.A * g.Opacity,
		)
	}

	return midColor.SetFillRGB()
}

// ShadingDict returns the PDF shading dictionary parameters for this gradient.
func (g *RadialGradient) ShadingDict(
	bounds Rect,
) map[string]interface{} {
	fx, fy := g.FocusPoint(bounds)
	cx, cy := g.CenterPoint(bounds)
	r := g.RadiusValue(bounds)

	return map[string]interface{}{
		"ShadingType": 3, // Radial shading
		"ColorSpace":  "/DeviceRGB",
		"Coords": []float64{
			fx,
			fy,
			0,
			cx,
			cy,
			r,
		},
		"Extend":   []bool{true, true},
		"Function": g.colorFunction(),
	}
}

// colorFunction returns the PDF function specification for the gradient colors.
func (g *RadialGradient) colorFunction() map[string]interface{} {
	if len(g.Stops) <= 2 {
		startColor := Black
		endColor := Black
		if len(g.Stops) > 0 {
			startColor = g.Stops[0].Color
		}
		if len(g.Stops) > 1 {
			endColor = g.Stops[len(g.Stops)-1].Color
		}

		return map[string]interface{}{
			"FunctionType": 2,
			"Domain":       []float64{0, 1},
			"C0": []float64{
				startColor.R,
				startColor.G,
				startColor.B,
			},
			"C1": []float64{
				endColor.R,
				endColor.G,
				endColor.B,
			},
			"N": 1,
		}
	}

	numSegments := len(g.Stops) - 1
	var bounds []float64
	if numSegments > 1 {
		bounds = make([]float64, 0, numSegments-1)
	}
	encode := make([]float64, 0, 2*numSegments)
	functions := make(
		[]map[string]interface{},
		0,
		numSegments,
	)

	for i := range numSegments {
		c0 := g.Stops[i].Color
		c1 := g.Stops[i+1].Color

		functions = append(
			functions,
			map[string]interface{}{
				"FunctionType": 2,
				"Domain":       []float64{0, 1},
				"C0": []float64{
					c0.R,
					c0.G,
					c0.B,
				},
				"C1": []float64{
					c1.R,
					c1.G,
					c1.B,
				},
				"N": 1,
			},
		)

		if i < len(g.Stops)-2 {
			bounds = append(
				bounds,
				g.Stops[i+1].Position,
			)
		}
		encode = append(encode, 0, 1)
	}

	return map[string]interface{}{
		"FunctionType": 3,
		"Domain":       []float64{0, 1},
		"Functions":    functions,
		"Bounds":       bounds,
		"Encode":       encode,
	}
}

// PatternFill represents a pattern fill.
type PatternFill struct {
	// PatternType is the type of pattern.
	PatternType PatternType
	// Color is the foreground color for the pattern.
	Color Color
	// BackgroundColor is the background color for the pattern.
	BackgroundColor Color
	// Spacing is the pattern spacing.
	Spacing float64
	// Angle is the pattern angle in degrees.
	Angle float64
}

// PatternType represents the type of pattern.
type PatternType int

const (
	// PatternTypeSolid is a solid pattern (no actual pattern).
	PatternTypeSolid PatternType = iota
	// PatternTypeHorizontalLines is horizontal lines.
	PatternTypeHorizontalLines
	// PatternTypeVerticalLines is vertical lines.
	PatternTypeVerticalLines
	// PatternTypeDiagonalLines is diagonal lines.
	PatternTypeDiagonalLines
	// PatternTypeCrossHatch is crosshatch pattern.
	PatternTypeCrossHatch
	// PatternTypeDots is a dot pattern.
	PatternTypeDots
)

// NewPatternFill creates a new pattern fill.
func NewPatternFill(
	patternType PatternType,
	color Color,
) *PatternFill {
	return &PatternFill{
		PatternType:     patternType,
		Color:           color,
		BackgroundColor: Transparent,
		Spacing:         4.0,
		Angle:           0,
	}
}

// SetBackground sets the pattern background color.
func (p *PatternFill) SetBackground(
	color Color,
) *PatternFill {
	p.BackgroundColor = color

	return p
}

// SetSpacing sets the pattern spacing.
func (p *PatternFill) SetSpacing(
	spacing float64,
) *PatternFill {
	p.Spacing = spacing

	return p
}

// SetAngle sets the pattern angle.
func (p *PatternFill) SetAngle(
	angle float64,
) *PatternFill {
	p.Angle = angle

	return p
}

// Type returns FillTypePattern.
func (p *PatternFill) Type() FillType { return FillTypePattern }

// IsNone returns true if both colors are transparent.
func (p *PatternFill) IsNone() bool {
	return p.Color.IsTransparent() &&
		p.BackgroundColor.IsTransparent()
}

// ContentStream returns PDF operators for the pattern.
// For simplicity, this returns the foreground color.
// A full implementation would create a tiling pattern.
func (p *PatternFill) ContentStream(
	bounds Rect,
) string {
	if p.IsNone() {
		return ""
	}
	// For solid pattern or simplified output, just use the foreground color
	if p.PatternType == PatternTypeSolid {
		return p.Color.SetFillRGB()
	}
	// For other patterns, we would need to create a tiling pattern
	// For now, blend the colors as an approximation
	blended := p.BackgroundColor.Blend(
		p.Color,
		0.5,
	)

	return blended.SetFillRGB()
}

// Utility functions for fill operations

// FillRect returns PDF operators to fill a rectangle with the given fill.
func FillRect(bounds Rect, fill Fill) string {
	if fill == nil || fill.IsNone() {
		return ""
	}

	var sb strings.Builder

	// Set fill color/pattern
	sb.WriteString(fill.ContentStream(bounds))
	sb.WriteString("\n")

	// Draw and fill rectangle
	sb.WriteString(
		fmt.Sprintf(
			"%s %s %s %s re f\n",
			formatFloat(
				bounds.X,
			),
			formatFloat(bounds.Y),
			formatFloat(
				bounds.Width,
			),
			formatFloat(bounds.Height),
		),
	)

	return sb.String()
}

// FillPath returns PDF operators to fill a path with the given fill.
func FillPath(
	path *PathBuilder,
	fill Fill,
	bounds Rect,
) string {
	if fill == nil || fill.IsNone() ||
		path == nil {
		return ""
	}

	var sb strings.Builder

	// Set fill color/pattern
	sb.WriteString(fill.ContentStream(bounds))
	sb.WriteString("\n")

	// Add path and fill
	sb.WriteString(path.Fill())

	return sb.String()
}

// SetFill returns PDF operators to set the current fill without drawing.
func SetFill(fill Fill, bounds Rect) string {
	if fill == nil || fill.IsNone() {
		return ""
	}

	return fill.ContentStream(bounds)
}
