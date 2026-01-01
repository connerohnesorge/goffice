package routing

import "github.com/connerohnesorge/goffice/drawingml"

// Router calculates a path between two points.
type Router interface {
	// Route calculates a path from start to end.
	// Returns the path as a sequence of drawing commands, or error if routing fails.
	Route(start, end Point, ctx *RoutingContext) (*Path, error)
}

// Point represents a 2D coordinate in EMU.
type Point struct {
	X drawingml.EMU
	Y drawingml.EMU
}

// RoutingContext provides context for routing operations.
type RoutingContext struct {
	// StartAngle is the exit angle from start connection site (in 1/60000 degrees)
	StartAngle int32

	// EndAngle is the entry angle to end connection site (in 1/60000 degrees)
	EndAngle int32

	// SlideWidth and SlideHeight define the routing area bounds
	SlideWidth  drawingml.EMU
	SlideHeight drawingml.EMU
}

// Path represents a connector path as a sequence of drawing commands.
type Path struct {
	Segments []PathSegment
}

// PathSegment is a marker interface for path command types.
type PathSegment interface {
	isPathSegment()
}

// MoveTo represents a move command (start of path).
type MoveTo struct {
	X drawingml.EMU
	Y drawingml.EMU
}

func (MoveTo) isPathSegment() {}

// LineTo represents a line segment.
type LineTo struct {
	X drawingml.EMU
	Y drawingml.EMU
}

func (LineTo) isPathSegment() {}

// CubicBezierTo represents a cubic Bezier curve.
type CubicBezierTo struct {
	CP1X drawingml.EMU // First control point X
	CP1Y drawingml.EMU // First control point Y
	CP2X drawingml.EMU // Second control point X
	CP2Y drawingml.EMU // Second control point Y
	X    drawingml.EMU // End point X
	Y    drawingml.EMU // End point Y
}

func (CubicBezierTo) isPathSegment() {}

// RouterType represents the type of routing algorithm to use.
type RouterType int

const (
	// RouterTypeStraight creates a direct line from start to end
	RouterTypeStraight RouterType = iota

	// RouterTypeElbow creates orthogonal paths with right-angle bends
	RouterTypeElbow

	// RouterTypeCurved creates smooth Bezier curves
	RouterTypeCurved
)

// String returns the string representation of RouterType.
func (rt RouterType) String() string {
	switch rt {
	case RouterTypeStraight:
		return "straight"
	case RouterTypeElbow:
		return "elbow"
	case RouterTypeCurved:
		return "curved"
	default:
		return "straight"
	}
}
