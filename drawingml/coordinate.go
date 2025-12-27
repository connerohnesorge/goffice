package drawingml

// Point2D represents a point in 2D space with X and Y coordinates in EMUs.
// This is used for positioning elements within a drawing.
type Point2D struct {
	// X is the horizontal coordinate in EMUs.
	X EMU
	// Y is the vertical coordinate in EMUs.
	Y EMU
}

// NewPoint2D creates a new Point2D with the specified coordinates in EMUs.
func NewPoint2D(x, y EMU) Point2D {
	return Point2D{X: x, Y: y}
}

// NewPoint2DFromInches creates a new Point2D from coordinates in inches.
func NewPoint2DFromInches(x, y float64) Point2D {
	return Point2D{
		X: InchesToEmu(x),
		Y: InchesToEmu(y),
	}
}

// NewPoint2DFromCm creates a new Point2D from coordinates in centimeters.
func NewPoint2DFromCm(x, y float64) Point2D {
	return Point2D{
		X: CmToEmu(x),
		Y: CmToEmu(y),
	}
}

// NewPoint2DFromPixels creates a new Point2D from coordinates in pixels
// (at 96 DPI).
func NewPoint2DFromPixels(x, y float64) Point2D {
	return Point2D{
		X: PixelsToEmu(x),
		Y: PixelsToEmu(y),
	}
}

// IsZero returns true if both coordinates are zero.
func (p Point2D) IsZero() bool {
	return p.X == 0 && p.Y == 0
}

// Add returns a new Point2D with the coordinates of both points added.
func (p Point2D) Add(other Point2D) Point2D {
	return Point2D{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

// Sub returns a new Point2D with the coordinates of other subtracted from p.
func (p Point2D) Sub(other Point2D) Point2D {
	return Point2D{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

// Clone creates a copy of this Point2D.
func (p Point2D) Clone() Point2D {
	return Point2D{X: p.X, Y: p.Y}
}

// Offset represents a positional offset in 2D space.
// This is semantically similar to Point2D but represents a displacement
// rather than an absolute position.
type Offset struct {
	// X is the horizontal offset in EMUs.
	X EMU
	// Y is the vertical offset in EMUs.
	Y EMU
}

// NewOffset creates a new Offset with the specified values in EMUs.
func NewOffset(x, y EMU) Offset {
	return Offset{X: x, Y: y}
}

// NewOffsetFromInches creates a new Offset from values in inches.
func NewOffsetFromInches(x, y float64) Offset {
	return Offset{
		X: InchesToEmu(x),
		Y: InchesToEmu(y),
	}
}

// NewOffsetFromCm creates a new Offset from values in centimeters.
func NewOffsetFromCm(x, y float64) Offset {
	return Offset{
		X: CmToEmu(x),
		Y: CmToEmu(y),
	}
}

// NewOffsetFromPixels creates a new Offset from values in pixels (at 96 DPI).
func NewOffsetFromPixels(x, y float64) Offset {
	return Offset{
		X: PixelsToEmu(x),
		Y: PixelsToEmu(y),
	}
}

// IsZero returns true if both offset values are zero.
func (o Offset) IsZero() bool {
	return o.X == 0 && o.Y == 0
}

// Clone creates a copy of this Offset.
func (o Offset) Clone() Offset {
	return Offset{X: o.X, Y: o.Y}
}

// ToPoint2D converts this Offset to a Point2D.
func (o Offset) ToPoint2D() Point2D {
	return Point2D(o)
}

// Extent represents dimensions (width and height) in 2D space.
// This is used to specify the size of drawing elements.
type Extent struct {
	// Cx is the width (extent in x direction) in EMUs.
	Cx EMU
	// Cy is the height (extent in y direction) in EMUs.
	Cy EMU
}

// NewExtent creates a new Extent with the specified dimensions in EMUs.
func NewExtent(cx, cy EMU) Extent {
	return Extent{Cx: cx, Cy: cy}
}

// NewExtentFromInches creates a new Extent from dimensions in inches.
func NewExtentFromInches(
	width, height float64,
) Extent {
	return Extent{
		Cx: InchesToEmu(width),
		Cy: InchesToEmu(height),
	}
}

// NewExtentFromCm creates a new Extent from dimensions in centimeters.
func NewExtentFromCm(
	width, height float64,
) Extent {
	return Extent{
		Cx: CmToEmu(width),
		Cy: CmToEmu(height),
	}
}

// NewExtentFromPixels creates a new Extent from dimensions in pixels
// (at 96 DPI).
func NewExtentFromPixels(
	width, height float64,
) Extent {
	return Extent{
		Cx: PixelsToEmu(width),
		Cy: PixelsToEmu(height),
	}
}

// Width returns the width (Cx) of the extent.
func (e Extent) Width() EMU {
	return e.Cx
}

// Height returns the height (Cy) of the extent.
func (e Extent) Height() EMU {
	return e.Cy
}

// IsZero returns true if both dimensions are zero.
func (e Extent) IsZero() bool {
	return e.Cx == 0 && e.Cy == 0
}

// Clone creates a copy of this Extent.
func (e Extent) Clone() Extent {
	return Extent{Cx: e.Cx, Cy: e.Cy}
}

// AspectRatio returns the aspect ratio (width/height) of the extent.
// Returns 0 if height is zero to avoid division by zero.
func (e Extent) AspectRatio() float64 {
	if e.Cy == 0 {
		return 0
	}

	return float64(e.Cx) / float64(e.Cy)
}

// ScaleToWidth returns a new Extent scaled to the specified width
// while maintaining the aspect ratio.
func (e Extent) ScaleToWidth(width EMU) Extent {
	if e.Cx == 0 {
		return Extent{Cx: width, Cy: 0}
	}
	ratio := float64(width) / float64(e.Cx)

	return Extent{
		Cx: width,
		Cy: EMU(float64(e.Cy) * ratio),
	}
}

// ScaleToHeight returns a new Extent scaled to the specified height
// while maintaining the aspect ratio.
func (e Extent) ScaleToHeight(height EMU) Extent {
	if e.Cy == 0 {
		return Extent{Cx: 0, Cy: height}
	}
	ratio := float64(height) / float64(e.Cy)

	return Extent{
		Cx: EMU(float64(e.Cx) * ratio),
		Cy: height,
	}
}

// Scale returns a new Extent with both dimensions multiplied by the given
// factor.
func (e Extent) Scale(factor float64) Extent {
	return Extent{
		Cx: EMU(float64(e.Cx) * factor),
		Cy: EMU(float64(e.Cy) * factor),
	}
}

// EffectExtent represents the additional space needed for effects
// like shadows and reflections around a drawing element.
type EffectExtent struct {
	// Left is the additional space on the left in EMUs.
	Left EMU
	// Top is the additional space on the top in EMUs.
	Top EMU
	// Right is the additional space on the right in EMUs.
	Right EMU
	// Bottom is the additional space on the bottom in EMUs.
	Bottom EMU
}

// NewEffectExtent creates a new EffectExtent with the specified values in EMUs.
func NewEffectExtent(
	left, top, right, bottom EMU,
) EffectExtent {
	return EffectExtent{
		Left:   left,
		Top:    top,
		Right:  right,
		Bottom: bottom,
	}
}

// NewEffectExtentUniform creates a new EffectExtent with the same value
// on all sides.
func NewEffectExtentUniform(
	value EMU,
) EffectExtent {
	return EffectExtent{
		Left:   value,
		Top:    value,
		Right:  value,
		Bottom: value,
	}
}

// IsZero returns true if all effect extents are zero.
func (e EffectExtent) IsZero() bool {
	return e.Left == 0 && e.Top == 0 &&
		e.Right == 0 &&
		e.Bottom == 0
}

// Clone creates a copy of this EffectExtent.
func (e EffectExtent) Clone() EffectExtent {
	return EffectExtent{
		Left:   e.Left,
		Top:    e.Top,
		Right:  e.Right,
		Bottom: e.Bottom,
	}
}

// TotalWidth returns the total horizontal effect extent (left + right).
func (e EffectExtent) TotalWidth() EMU {
	return e.Left + e.Right
}

// TotalHeight returns the total vertical effect extent (top + bottom).
func (e EffectExtent) TotalHeight() EMU {
	return e.Top + e.Bottom
}
