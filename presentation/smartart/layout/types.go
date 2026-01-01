package layout

// Rectangle defines a bounding area
type Rectangle struct {
	X, Y          int64 // Top-left corner (EMU)
	Width, Height int64 // Dimensions (EMU)
}

// Point represents a 2D coordinate
type Point struct {
	X, Y int64
}

// Orientation for layouts
type Orientation int

const (
	OrientationVertical Orientation = iota
	OrientationHorizontal
)

// Alignment for positioning
type Alignment int

const (
	AlignStart Alignment = iota
	AlignCenter
	AlignEnd
)
