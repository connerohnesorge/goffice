package layout

import "testing"

func TestListLayout_Vertical(t *testing.T) {
	layout := NewListLayout()
	layout.Orientation = OrientationVertical

	shapes := []LayoutShape{
		{ID: "1", Text: "Item 1"},
		{ID: "2", Text: "Item 2"},
		{ID: "3", Text: "Item 3"},
	}

	bounds := Rectangle{X: 0, Y: 0, Width: 1000000, Height: 3000000}

	err := layout.Layout(shapes, bounds)
	if err != nil {
		t.Fatal(err)
	}

	// Verify shapes are positioned vertically
	for i := 1; i < len(shapes); i++ {
		if shapes[i].Y <= shapes[i-1].Y {
			t.Errorf("Shape %d not below shape %d", i, i-1)
		}
	}

	// Verify all shapes have been assigned positions
	for i, shape := range shapes {
		if shape.Width == 0 || shape.Height == 0 {
			t.Errorf("Shape %d has zero dimensions: width=%d, height=%d", i, shape.Width, shape.Height)
		}
	}
}

func TestListLayout_Horizontal(t *testing.T) {
	layout := NewListLayout()
	layout.Orientation = OrientationHorizontal

	shapes := []LayoutShape{
		{ID: "1", Text: "Item 1"},
		{ID: "2", Text: "Item 2"},
		{ID: "3", Text: "Item 3"},
	}

	bounds := Rectangle{X: 0, Y: 0, Width: 3000000, Height: 1000000}

	err := layout.Layout(shapes, bounds)
	if err != nil {
		t.Fatal(err)
	}

	// Verify shapes are positioned horizontally
	for i := 1; i < len(shapes); i++ {
		if shapes[i].X <= shapes[i-1].X {
			t.Errorf("Shape %d not right of shape %d", i, i-1)
		}
	}

	// Verify all shapes have been assigned positions
	for i, shape := range shapes {
		if shape.Width == 0 || shape.Height == 0 {
			t.Errorf("Shape %d has zero dimensions: width=%d, height=%d", i, shape.Width, shape.Height)
		}
	}
}

func TestListLayout_EmptyShapes(t *testing.T) {
	layout := NewListLayout()
	var shapes []LayoutShape
	bounds := Rectangle{X: 0, Y: 0, Width: 1000000, Height: 1000000}

	err := layout.Layout(shapes, bounds)
	if err != nil {
		t.Errorf("Layout should handle empty shapes without error: %v", err)
	}
}

func TestListLayout_SingleShape(t *testing.T) {
	layout := NewListLayout()
	layout.Orientation = OrientationVertical

	shapes := []LayoutShape{
		{ID: "1", Text: "Single Item"},
	}

	bounds := Rectangle{X: 0, Y: 0, Width: 1000000, Height: 1000000}

	err := layout.Layout(shapes, bounds)
	if err != nil {
		t.Fatal(err)
	}

	// Verify single shape is positioned
	if shapes[0].Width == 0 || shapes[0].Height == 0 {
		t.Errorf("Single shape has zero dimensions: width=%d, height=%d", shapes[0].Width, shapes[0].Height)
	}
}

func TestRegistry(t *testing.T) {
	reg := DefaultRegistry()

	engine, ok := reg.GetEngine(LayoutTypeList)
	if !ok {
		t.Fatal("List layout not registered")
	}

	if engine.Type() != LayoutTypeList {
		t.Errorf("Wrong layout type: got %v, want %v", engine.Type(), LayoutTypeList)
	}
}

func TestRegistry_UnknownType(t *testing.T) {
	reg := DefaultRegistry()

	// Try to get an engine for a layout type that hasn't been registered yet
	_, ok := reg.GetEngine(LayoutTypeHierarchy)
	if ok {
		t.Error("Should not find unregistered layout type")
	}
}

func TestDefaultConstraints(t *testing.T) {
	c := DefaultConstraints()

	if c.Spacing == 0 {
		t.Error("Default spacing should not be zero")
	}
	if c.Padding == 0 {
		t.Error("Default padding should not be zero")
	}
	if c.MinWidth == 0 {
		t.Error("Default min width should not be zero")
	}
	if c.MinHeight == 0 {
		t.Error("Default min height should not be zero")
	}
	if c.Alignment != AlignCenter {
		t.Errorf("Default alignment should be center, got %v", c.Alignment)
	}
}

func TestListLayout_RespectsBounds(t *testing.T) {
	layout := NewListLayout()
	layout.Orientation = OrientationVertical

	shapes := []LayoutShape{
		{ID: "1", Text: "Item 1"},
		{ID: "2", Text: "Item 2"},
	}

	bounds := Rectangle{X: 100, Y: 200, Width: 1000000, Height: 2000000}

	err := layout.Layout(shapes, bounds)
	if err != nil {
		t.Fatal(err)
	}

	// Verify shapes respect the bounds
	for i, shape := range shapes {
		if shape.X < bounds.X {
			t.Errorf("Shape %d X position %d is outside left bound %d", i, shape.X, bounds.X)
		}
		if shape.Y < bounds.Y {
			t.Errorf("Shape %d Y position %d is outside top bound %d", i, shape.Y, bounds.Y)
		}
		if shape.X+shape.Width > bounds.X+bounds.Width {
			t.Errorf("Shape %d extends beyond right bound", i)
		}
	}
}

func TestListLayout_CustomConstraints(t *testing.T) {
	layout := NewListLayout()
	layout.Orientation = OrientationVertical
	layout.Constraints = Constraints{
		Spacing:   500000, // Custom spacing
		Padding:   250000, // Custom padding
		MinWidth:  500000,
		MinHeight: 250000,
		Alignment: AlignCenter,
	}

	shapes := []LayoutShape{
		{ID: "1", Text: "Item 1"},
		{ID: "2", Text: "Item 2"},
	}

	bounds := Rectangle{X: 0, Y: 0, Width: 2000000, Height: 2000000}

	err := layout.Layout(shapes, bounds)
	if err != nil {
		t.Fatal(err)
	}

	// Verify custom constraints are applied
	if shapes[0].X != layout.Constraints.Padding {
		t.Errorf("Shape 0 X should match padding constraint: got %d, want %d", shapes[0].X, layout.Constraints.Padding)
	}

	// Verify spacing between shapes
	spacing := shapes[1].Y - (shapes[0].Y + shapes[0].Height)
	if spacing != layout.Constraints.Spacing {
		t.Errorf("Spacing between shapes should be %d, got %d", layout.Constraints.Spacing, spacing)
	}
}
