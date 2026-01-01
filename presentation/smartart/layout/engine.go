package layout

// LayoutType identifies the layout algorithm
type LayoutType int

const (
	LayoutTypeHierarchy LayoutType = iota
	LayoutTypeList
	LayoutTypeCycle
	LayoutTypePyramid
	LayoutTypeMatrix
)

// LayoutEngine defines the interface for layout algorithms
type LayoutEngine interface {
	// Layout positions shapes within the given bounds
	Layout(shapes []LayoutShape, bounds Rectangle) error

	// Type returns the layout type this engine handles
	Type() LayoutType
}

// LayoutShape represents a shape to be positioned
type LayoutShape struct {
	ID     string
	Text   string
	Width  int64 // EMU
	Height int64 // EMU
	X      int64 // Output X position (EMU)
	Y      int64 // Output Y position (EMU)
}
