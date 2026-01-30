package core

import (
	"fmt"
	"math"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

// ShadingType represents the type of PDF shading.
type ShadingType int

const (
	// ShadingTypeAxial represents an axial (linear) shading.
	ShadingTypeAxial ShadingType = 2
	// ShadingTypeRadial represents a radial shading.
	ShadingTypeRadial ShadingType = 3
)

// GradientStop represents a color stop in a gradient.
type GradientStop struct {
	// Position is the position of this stop along the gradient (0-1).
	Position float64
	// R, G, B are the color components (0-1).
	R, G, B float64
}

// LinearGradient represents a linear gradient definition.
type LinearGradient struct {
	// Angle is the angle of the gradient in degrees (0 = left to right, 90 = bottom to top).
	Angle float64
	// Stops is the list of color stops.
	Stops []GradientStop
}

// GradientResource represents a gradient that can be registered as a PDF shading resource.
type GradientResource struct {
	gradient *LinearGradient
	bounds   Rectangle
}

// NewGradientResource creates a new gradient resource.
func NewGradientResource(gradient *LinearGradient, bounds Rectangle) *GradientResource {
	return &GradientResource{
		gradient: gradient,
		bounds:   bounds,
	}
}

// RegisterGradient registers a gradient as a shading pattern on the given page and returns
// the pattern name used for drawing (e.g., "P1").
func (d *Document) RegisterGradient(
	page *Page,
	gradient *LinearGradient,
	bounds Rectangle,
) (string, error) {
	if page == nil {
		return "", fmt.Errorf("page is nil")
	}
	if gradient == nil {
		return "", fmt.Errorf("gradient is nil")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	page.mu.Lock()
	defer page.mu.Unlock()

	// Create and register the shading
	shadingRef, err := d.registerAxialShadingLocked(gradient, bounds)
	if err != nil {
		return "", err
	}

	// Create and register the pattern that uses the shading
	patternRef, err := d.registerShadingPatternLocked(shadingRef)
	if err != nil {
		return "", err
	}

	name := nextPatternNameLocked(page)

	if page.resources == nil {
		page.resources = types.NewDict()
	}

	// Add to Pattern dictionary
	patternDict, ok := page.resources["Pattern"].(types.Dict)
	if !ok {
		patternDict = types.NewDict()
	}
	patternDict[name] = *patternRef
	page.resources["Pattern"] = patternDict

	// Set up Pattern colorspace if not already present
	csDict, ok := page.resources["ColorSpace"].(types.Dict)
	if !ok {
		csDict = types.NewDict()
	}
	// Pattern colorspace allows using patterns as colors
	csDict["Pattern"] = types.Name("Pattern")
	page.resources["ColorSpace"] = csDict

	return name, nil
}

// registerAxialShadingLocked creates an axial shading object for a linear gradient.
func (d *Document) registerAxialShadingLocked(
	gradient *LinearGradient,
	bounds Rectangle,
) (*types.IndirectRef, error) {
	// Calculate gradient vector based on angle
	angleRad := gradient.Angle * math.Pi / 180

	// Calculate the diagonal of the bounding box to ensure gradient covers entire area
	diagonal := math.Sqrt(bounds.Width()*bounds.Width() + bounds.Height()*bounds.Height())

	// Center point of bounds
	centerX := bounds.LLX + bounds.Width()/2
	centerY := bounds.LLY + bounds.Height()/2

	// Calculate start and end points based on angle
	// The gradient line passes through the center and extends beyond bounds
	halfDiag := diagonal / 2
	cosAngle := math.Cos(angleRad)
	sinAngle := math.Sin(angleRad)

	// Start point (opposite direction of angle)
	x0 := centerX - halfDiag*cosAngle
	y0 := centerY - halfDiag*sinAngle

	// End point (direction of angle)
	x1 := centerX + halfDiag*cosAngle
	y1 := centerY + halfDiag*sinAngle

	// Create the shading dictionary
	shadingDict := types.NewDict()
	shadingDict.Insert("ShadingType", types.Integer(2)) // Axial shading
	shadingDict.Insert("ColorSpace", types.Name("DeviceRGB"))
	shadingDict.Insert("Coords", types.NewNumberArray(x0, y0, x1, y1))

	// Create the function for color interpolation
	functionRef, err := d.createGradientFunctionLocked(gradient.Stops)
	if err != nil {
		return nil, fmt.Errorf("create gradient function: %w", err)
	}

	shadingDict.Insert("Function", *functionRef)
	shadingDict.Insert("Extend", types.Array{types.Boolean(true), types.Boolean(true)})

	return d.ctx.XRefTable.IndRefForNewObject(shadingDict)
}

// createGradientFunctionLocked creates a PDF function for gradient color interpolation.
// For multiple stops, this creates a stitching function (Type 3).
func (d *Document) createGradientFunctionLocked(stops []GradientStop) (*types.IndirectRef, error) {
	if len(stops) < 2 {
		return nil, fmt.Errorf("gradient must have at least 2 stops")
	}

	if len(stops) == 2 {
		// Simple linear interpolation between two colors
		return d.createExponentialFunctionLocked(stops[0], stops[1])
	}

	// Multiple stops - create stitching function
	return d.createStitchingFunctionLocked(stops)
}

// createExponentialFunctionLocked creates a Type 2 (exponential) function for linear interpolation.
func (d *Document) createExponentialFunctionLocked(
	start, end GradientStop,
) (*types.IndirectRef, error) {
	funcDict := types.NewDict()
	funcDict.Insert("FunctionType", types.Integer(2))
	funcDict.Insert("Domain", types.NewNumberArray(0, 1))
	funcDict.Insert("C0", types.NewNumberArray(start.R, start.G, start.B))
	funcDict.Insert("C1", types.NewNumberArray(end.R, end.G, end.B))
	funcDict.Insert("N", types.Integer(1)) // Linear interpolation

	return d.ctx.XRefTable.IndRefForNewObject(funcDict)
}

// createStitchingFunctionLocked creates a Type 3 (stitching) function for multiple stops.
func (d *Document) createStitchingFunctionLocked(stops []GradientStop) (*types.IndirectRef, error) {
	// Sort stops by position
	sortedStops := make([]GradientStop, len(stops))
	copy(sortedStops, stops)
	for i := 0; i < len(sortedStops)-1; i++ {
		for j := i + 1; j < len(sortedStops); j++ {
			if sortedStops[j].Position < sortedStops[i].Position {
				sortedStops[i], sortedStops[j] = sortedStops[j], sortedStops[i]
			}
		}
	}

	// Create sub-functions for each segment
	var functions types.Array
	var bounds types.Array
	var encode types.Array

	for i := 0; i < len(sortedStops)-1; i++ {
		start := sortedStops[i]
		end := sortedStops[i+1]

		// Create function for this segment
		funcRef, err := d.createExponentialFunctionLocked(start, end)
		if err != nil {
			return nil, err
		}
		functions = append(functions, *funcRef)

		// Add bound (except for last segment)
		if i < len(sortedStops)-2 {
			bounds = append(bounds, types.Float(end.Position))
		}

		// Each sub-function maps from 0 to 1
		encode = append(encode, types.Integer(0), types.Integer(1))
	}

	stitchDict := types.NewDict()
	stitchDict.Insert("FunctionType", types.Integer(3))
	stitchDict.Insert("Domain", types.NewNumberArray(sortedStops[0].Position, sortedStops[len(sortedStops)-1].Position))
	stitchDict.Insert("Range", types.NewNumberArray(0, 1, 0, 1, 0, 1))
	stitchDict.Insert("Functions", functions)

	if len(bounds) > 0 {
		stitchDict.Insert("Bounds", bounds)
	}
	stitchDict.Insert("Encode", encode)

	return d.ctx.XRefTable.IndRefForNewObject(stitchDict)
}

// registerShadingPatternLocked creates a pattern object that references a shading.
func (d *Document) registerShadingPatternLocked(shadingRef *types.IndirectRef) (*types.IndirectRef, error) {
	patternDict := types.NewDict()
	patternDict.Insert("Type", types.Name("Pattern"))
	patternDict.Insert("PatternType", types.Integer(2)) // Shading pattern
	patternDict.Insert("Shading", *shadingRef)

	return d.ctx.XRefTable.IndRefForNewObject(patternDict)
}

func nextPatternNameLocked(page *Page) string {
	if page.patterns == nil {
		page.patterns = make(map[string]bool)
	}

	index := len(page.patterns) + 1
	name := fmt.Sprintf("P%d", index)
	for page.patterns[name] {
		index++
		name = fmt.Sprintf("P%d", index)
	}

	page.patterns[name] = true

	return name
}

// Page interface extension methods

// SetGradientFill sets the fill to use a gradient pattern.
// This writes the PDF operators to use a pattern colorspace.
func (p *PageImpl) SetGradientFill(patternName string) {
	// Set pattern colorspace and select pattern
	// /Pattern cs - set colorspace to Pattern
	// /PatternName scn - set color to pattern (for fill)
	p.content.WriteString(fmt.Sprintf("/Pattern cs /%s scn\n", patternName))
	p.fillColorSet = false // Gradient is not a solid color
}

// GradientFillType identifies what type of gradient fill is being used.
type GradientFillType int

const (
	// GradientFillNone means no gradient fill is active.
	GradientFillNone GradientFillType = iota
	// GradientFillLinear means a linear gradient fill is active.
	GradientFillLinear
)
