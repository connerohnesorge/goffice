package routing

// StraightRouter creates a direct line from start to end.
// This is the simplest routing algorithm that ignores obstacles.
type StraightRouter struct{}

// Route creates a straight line path from start to end.
func (*StraightRouter) Route(start, end Point, _ *RoutingContext) (*Path, error) {
	// Handle coincident points (zero-length line)
	if start.X == end.X && start.Y == end.Y {
		return &Path{
			Segments: []PathSegment{
				MoveTo(start),
			},
		}, nil
	}

	return &Path{
		Segments: []PathSegment{
			MoveTo(start),
			LineTo(end),
		},
	}, nil
}
