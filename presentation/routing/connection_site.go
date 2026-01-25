package routing

import "github.com/connerohnesorge/goffice/drawingml"

// Connection site angle constants (in 1/60000 degrees).
const (
	AngleUp    = 270 * 60000 // Up/North
	AngleRight = 0           // Right/East
	AngleDown  = 90 * 60000  // Down/South
	AngleLeft  = 180 * 60000 // Left/West
)

// Default connection site indices for rectangles.
const (
	SiteIndexTop = iota
	SiteIndexRight
	SiteIndexBottom
	SiteIndexLeft
)

// ConnectionSite represents a point where connectors can attach to a shape.
// Corresponds to ECMA-376 CT_ConnectionSite (a:cxn element).
type ConnectionSite struct{
	// Index is the 0-based identifier for this connection site.
	// Used in connector start/end connection references.
	Index int

	// X is the X position in EMU relative to shape bounds.
	// (0, 0) = top-left corner of shape.
	X drawingml.EMU

	// Y is the Y position in EMU relative to shape bounds.
	Y drawingml.EMU

	// Angle is the direction normal to the shape edge at this point.
	// Measured in 1/60000 degrees (ECMA-376 standard).
	// 0 = right, 90*60000 = down, 180*60000 = left, 270*60000 = up.
	Angle int32
}

// GenerateDefaultConnectionSites returns default connection sites for rectangles.
// Returns 4 sites: top center, right center, bottom center, left center.
func GenerateDefaultConnectionSites(shapeWidth, shapeHeight drawingml.EMU) []*ConnectionSite {
	return []*ConnectionSite{
		// Index 0: Top center (normal pointing up)
		{
			Index: SiteIndexTop,
			X:     shapeWidth / 2,
			Y:     0,
			Angle: AngleUp,
		},

		// Index 1: Right center (normal pointing right)
		{
			Index: SiteIndexRight,
			X:     shapeWidth,
			Y:     shapeHeight / 2,
			Angle: AngleRight,
		},

		// Index 2: Bottom center (normal pointing down)
		{
			Index: SiteIndexBottom,
			X:     shapeWidth / 2,
			Y:     shapeHeight,
			Angle: AngleDown,
		},

		// Index 3: Left center (normal pointing left)
		{
			Index: SiteIndexLeft,
			X:     0,
			Y:     shapeHeight / 2,
			Angle: AngleLeft,
		},
	}
}
