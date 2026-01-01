package routing

import (
	"testing"

	"github.com/connerohnesorge/goffice/drawingml"
)

func TestStraightRouter_BasicLine(t *testing.T) {
	router := &StraightRouter{}

	start := Point{X: 100 * drawingml.EMUsPerPoint, Y: 100 * drawingml.EMUsPerPoint}
	end := Point{X: 300 * drawingml.EMUsPerPoint, Y: 200 * drawingml.EMUsPerPoint}

	ctx := &RoutingContext{
		SlideWidth:  1000 * drawingml.EMUsPerPoint,
		SlideHeight: 750 * drawingml.EMUsPerPoint,
	}

	path, err := router.Route(start, end, ctx)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(path.Segments) != 2 {
		t.Fatalf("Expected 2 segments (MoveTo, LineTo), got %d", len(path.Segments))
	}

	// Check MoveTo segment
	moveTo, ok := path.Segments[0].(MoveTo)
	if !ok {
		t.Fatalf("Expected first segment to be MoveTo, got %T", path.Segments[0])
	}
	if moveTo.X != start.X || moveTo.Y != start.Y {
		t.Errorf("MoveTo: expected (%d, %d), got (%d, %d)", start.X, start.Y, moveTo.X, moveTo.Y)
	}

	// Check LineTo segment
	lineTo, ok := path.Segments[1].(LineTo)
	if !ok {
		t.Fatalf("Expected second segment to be LineTo, got %T", path.Segments[1])
	}
	if lineTo.X != end.X || lineTo.Y != end.Y {
		t.Errorf("LineTo: expected (%d, %d), got (%d, %d)", end.X, end.Y, lineTo.X, lineTo.Y)
	}
}

func TestStraightRouter_CoincidentPoints(t *testing.T) {
	router := &StraightRouter{}

	point := Point{X: 100 * drawingml.EMUsPerPoint, Y: 100 * drawingml.EMUsPerPoint}

	ctx := &RoutingContext{
		SlideWidth:  1000 * drawingml.EMUsPerPoint,
		SlideHeight: 750 * drawingml.EMUsPerPoint,
	}

	path, err := router.Route(point, point, ctx)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(path.Segments) != 1 {
		t.Fatalf("Expected 1 segment (MoveTo only), got %d", len(path.Segments))
	}

	// Check MoveTo segment
	moveTo, ok := path.Segments[0].(MoveTo)
	if !ok {
		t.Fatalf("Expected segment to be MoveTo, got %T", path.Segments[0])
	}
	if moveTo.X != point.X || moveTo.Y != point.Y {
		t.Errorf("MoveTo: expected (%d, %d), got (%d, %d)", point.X, point.Y, moveTo.X, moveTo.Y)
	}
}

func TestStraightRouter_VerticalLine(t *testing.T) {
	router := &StraightRouter{}

	start := Point{X: 100 * drawingml.EMUsPerPoint, Y: 100 * drawingml.EMUsPerPoint}
	end := Point{X: 100 * drawingml.EMUsPerPoint, Y: 300 * drawingml.EMUsPerPoint}

	ctx := &RoutingContext{}

	path, err := router.Route(start, end, ctx)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(path.Segments) != 2 {
		t.Fatalf("Expected 2 segments, got %d", len(path.Segments))
	}

	//nolint:revive // Type already verified above
	lineTo := path.Segments[1].(LineTo)
	if lineTo.X != start.X {
		t.Errorf("Expected vertical line (same X), got start.X=%d, end.X=%d", start.X, lineTo.X)
	}
}

func TestStraightRouter_HorizontalLine(t *testing.T) {
	router := &StraightRouter{}

	start := Point{X: 100 * drawingml.EMUsPerPoint, Y: 100 * drawingml.EMUsPerPoint}
	end := Point{X: 300 * drawingml.EMUsPerPoint, Y: 100 * drawingml.EMUsPerPoint}

	ctx := &RoutingContext{}

	path, err := router.Route(start, end, ctx)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(path.Segments) != 2 {
		t.Fatalf("Expected 2 segments, got %d", len(path.Segments))
	}

	//nolint:revive // Type already verified above
	lineTo := path.Segments[1].(LineTo)
	if lineTo.Y != start.Y {
		t.Errorf("Expected horizontal line (same Y), got start.Y=%d, end.Y=%d", start.Y, lineTo.Y)
	}
}

func TestRouterType_String(t *testing.T) {
	tests := []struct {
		routerType RouterType
		expected   string
	}{
		{RouterTypeStraight, "straight"},
		{RouterTypeElbow, "elbow"},
		{RouterTypeCurved, "curved"},
		{RouterType(999), "straight"}, // Unknown type defaults to straight
	}

	for _, tt := range tests {
		got := tt.routerType.String()
		if got != tt.expected {
			t.Errorf("RouterType(%d).String() = %s, want %s", tt.routerType, got, tt.expected)
		}
	}
}

func TestPathSegments_Interface(_ *testing.T) {
	// Verify all path segment types implement PathSegment interface
	var _ PathSegment = MoveTo{}
	var _ PathSegment = LineTo{}
	var _ PathSegment = CubicBezierTo{}
}
