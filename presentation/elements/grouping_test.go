package elements

import (
	"testing"

	"github.com/connerohnesorge/goffice/drawingml"
)

func TestComputeBoundingBox(t *testing.T) {
	t.Run("EmptyShapes", func(t *testing.T) {
		x, y, cx, cy := computeBoundingBox(nil)
		if x != 0 || y != 0 || cx != 0 || cy != 0 {
			t.Errorf("Expected zero bounding box, got (%d, %d, %d, %d)", x, y, cx, cy)
		}
	})

	t.Run("SingleShape", func(t *testing.T) {
		shape := NewShape()
		spPr := shape.GetOrCreateShapeProperties()

		// Create transform with offset and extent
		xfrm := spPr.getOrCreateTransform()
		xfrmT2D := &drawingml.Transform2D{CompositeElementBase: xfrm}
		xfrmT2D.SetOffset(drawingml.Offset{X: 100, Y: 200})
		xfrmT2D.SetExtent(drawingml.Extent{Cx: 300, Cy: 400})

		x, y, cx, cy := computeBoundingBox([]any{shape})
		if x != 100 || y != 200 || cx != 300 || cy != 400 {
			t.Errorf("Expected (100, 200, 300, 400), got (%d, %d, %d, %d)", x, y, cx, cy)
		}
	})

	t.Run("MultipleShapes", func(t *testing.T) {
		// Shape 1 at (100, 200) with size (300, 400)
		shape1 := NewShape()
		spPr1 := shape1.GetOrCreateShapeProperties()
		xfrm1 := spPr1.getOrCreateTransform()
		xfrmT2D1 := &drawingml.Transform2D{CompositeElementBase: xfrm1}
		xfrmT2D1.SetOffset(drawingml.Offset{X: 100, Y: 200})
		xfrmT2D1.SetExtent(drawingml.Extent{Cx: 300, Cy: 400})

		// Shape 2 at (500, 100) with size (200, 300)
		shape2 := NewShape()
		spPr2 := shape2.GetOrCreateShapeProperties()
		xfrm2 := spPr2.getOrCreateTransform()
		xfrmT2D2 := &drawingml.Transform2D{CompositeElementBase: xfrm2}
		xfrmT2D2.SetOffset(drawingml.Offset{X: 500, Y: 100})
		xfrmT2D2.SetExtent(drawingml.Extent{Cx: 200, Cy: 300})

		// Bounding box should be:
		// minX = 100, minY = 100
		// maxX = 500 + 200 = 700, maxY = 200 + 400 = 600
		// So: (100, 100) with size (600, 500)
		x, y, cx, cy := computeBoundingBox([]any{shape1, shape2})
		expectedX, expectedY := drawingml.EMU(100), drawingml.EMU(100)
		expectedCx, expectedCy := drawingml.EMU(600), drawingml.EMU(500)

		if x != expectedX || y != expectedY || cx != expectedCx || cy != expectedCy {
			t.Errorf("Expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
				expectedX, expectedY, expectedCx, expectedCy, x, y, cx, cy)
		}
	})

	t.Run("MixedShapeTypes", func(t *testing.T) {
		// Create a shape
		shape := NewShape()
		shape.SetPosition(0, 0)
		shape.SetSize(100, 100)

		// Create a picture
		pic := NewPicture("rId1")
		pic.SetPosition(200, 200)
		pic.SetSize(150, 150)

		// Bounding box: (0, 0) to (350, 350)
		x, y, cx, cy := computeBoundingBox([]any{shape, pic})
		if x != 0 || y != 0 || cx != 350 || cy != 350 {
			t.Errorf("Expected (0, 0, 350, 350), got (%d, %d, %d, %d)", x, y, cx, cy)
		}
	})
}

func TestAdjustCoordinatesRelativeToGroup(t *testing.T) {
	t.Run("AdjustShapeCoordinates", func(t *testing.T) {
		shape := NewShape()
		spPr := shape.GetOrCreateShapeProperties()
		xfrm := spPr.getOrCreateTransform()
		xfrmT2D := &drawingml.Transform2D{CompositeElementBase: xfrm}

		// Shape at absolute position (500, 600)
		xfrmT2D.SetOffset(drawingml.Offset{X: 500, Y: 600})
		xfrmT2D.SetExtent(drawingml.Extent{Cx: 100, Cy: 100})

		// Adjust relative to group at (100, 200)
		adjustCoordinatesRelativeToGroup(shape, 100, 200)

		// Should now be at (400, 400) relative to group
		off := xfrmT2D.Offset()
		if off.X != 400 || off.Y != 400 {
			t.Errorf("Expected offset (400, 400), got (%d, %d)", off.X, off.Y)
		}
	})

	t.Run("AdjustPictureCoordinates", func(t *testing.T) {
		pic := NewPicture("rId1")
		spPr := pic.GetOrCreateShapeProperties()
		xfrm := spPr.getOrCreateTransform()
		xfrmT2D := &drawingml.Transform2D{CompositeElementBase: xfrm}

		// Picture at absolute position (300, 400)
		xfrmT2D.SetOffset(drawingml.Offset{X: 300, Y: 400})
		xfrmT2D.SetExtent(drawingml.Extent{Cx: 200, Cy: 200})

		// Adjust relative to group at (50, 50)
		adjustCoordinatesRelativeToGroup(pic, 50, 50)

		// Should now be at (250, 350) relative to group
		off := xfrmT2D.Offset()
		if off.X != 250 || off.Y != 350 {
			t.Errorf("Expected offset (250, 350), got (%d, %d)", off.X, off.Y)
		}
	})
}

func TestCreateGroupUngroupRoundtrip(t *testing.T) {
	// Create a slide
	slide := NewSlide()

	// Create shapes at known positions
	shape1 := slide.AddShape()
	spPr1 := shape1.GetOrCreateShapeProperties()
	xfrm1 := spPr1.getOrCreateTransform()
	xfrmT2D1 := &drawingml.Transform2D{CompositeElementBase: xfrm1}
	xfrmT2D1.SetOffset(drawingml.Offset{X: 100, Y: 200})
	xfrmT2D1.SetExtent(drawingml.Extent{Cx: 300, Cy: 400})

	shape2 := slide.AddShape()
	spPr2 := shape2.GetOrCreateShapeProperties()
	xfrm2 := spPr2.getOrCreateTransform()
	xfrmT2D2 := &drawingml.Transform2D{CompositeElementBase: xfrm2}
	xfrmT2D2.SetOffset(drawingml.Offset{X: 500, Y: 600})
	xfrmT2D2.SetExtent(drawingml.Extent{Cx: 200, Cy: 300})

	// Store original positions
	origX1, origY1 := xfrmT2D1.Offset().X, xfrmT2D1.Offset().Y
	origX2, origY2 := xfrmT2D2.Offset().X, xfrmT2D2.Offset().Y

	// Create group
	group := slide.CreateGroup(shape1, shape2)
	if group == nil {
		t.Fatal("CreateGroup returned nil")
	}

	// Verify group is created with correct transform
	gspPr := group.GroupShapeProperties()
	if gspPr == nil {
		t.Fatal("GroupShapeProperties is nil")
	}

	xfrm := gspPr.Transform()
	if xfrm == nil {
		t.Fatal("Group transform is nil")
	}

	xfrmT2D := &drawingml.Transform2D{CompositeElementBase: xfrm}
	groupOff := xfrmT2D.Offset()
	groupExt := xfrmT2D.Extent()

	// Group should be at (100, 200) with size (600, 700)
	// Because bounding box is from (100, 200) to (700, 900)
	if groupOff.X != 100 || groupOff.Y != 200 {
		t.Errorf("Expected group offset (100, 200), got (%d, %d)", groupOff.X, groupOff.Y)
	}
	if groupExt.Cx != 600 || groupExt.Cy != 700 {
		t.Errorf("Expected group extent (600, 700), got (%d, %d)", groupExt.Cx, groupExt.Cy)
	}

	// Verify shapes are in group
	groupShapes := group.Shapes()
	if len(groupShapes) != 2 {
		t.Fatalf("Expected 2 shapes in group, got %d", len(groupShapes))
	}

	// Verify shapes have relative coordinates
	relOff1 := xfrmT2D1.Offset()
	relOff2 := xfrmT2D2.Offset()

	// shape1 should be at (0, 0) relative to group
	if relOff1.X != 0 || relOff1.Y != 0 {
		t.Errorf("Expected shape1 relative offset (0, 0), got (%d, %d)", relOff1.X, relOff1.Y)
	}

	// shape2 should be at (400, 400) relative to group
	if relOff2.X != 400 || relOff2.Y != 400 {
		t.Errorf("Expected shape2 relative offset (400, 400), got (%d, %d)", relOff2.X, relOff2.Y)
	}

	// Ungroup
	ungroupedShapes := slide.UngroupShape(group)
	if len(ungroupedShapes) != 2 {
		t.Fatalf("Expected 2 ungrouped shapes, got %d", len(ungroupedShapes))
	}

	// Verify shapes are back at original absolute positions
	absOff1 := xfrmT2D1.Offset()
	absOff2 := xfrmT2D2.Offset()

	if absOff1.X != origX1 || absOff1.Y != origY1 {
		t.Errorf("Expected shape1 back at (%d, %d), got (%d, %d)",
			origX1, origY1, absOff1.X, absOff1.Y)
	}

	if absOff2.X != origX2 || absOff2.Y != origY2 {
		t.Errorf("Expected shape2 back at (%d, %d), got (%d, %d)",
			origX2, origY2, absOff2.X, absOff2.Y)
	}

	// Verify shapes are back in slide's shape tree
	st := slide.ShapeTree()
	slideshapes := st.Shapes()
	if len(slideshapes) != 2 {
		t.Errorf("Expected 2 shapes in slide after ungroup, got %d", len(slideshapes))
	}
}

func TestCreateGroupEmptyShapes(t *testing.T) {
	slide := NewSlide()

	group := slide.CreateGroup()
	if group != nil {
		t.Error("Expected nil for empty CreateGroup, got non-nil")
	}
}

func TestChildOffsetExtent(t *testing.T) {
	t.Run("SetAndGetChildOffset", func(t *testing.T) {
		xfrm := drawingml.NewTransform2D(0, 0, 1000, 1000)

		// Initially should not be present
		_, _, present := xfrm.ChildOffset()
		if present {
			t.Error("Expected ChildOffset not present initially")
		}

		// Set child offset
		xfrm.SetChildOffset(100, 200)

		x, y, present := xfrm.ChildOffset()
		if !present {
			t.Error("Expected ChildOffset to be present after SetChildOffset")
		}
		if x != 100 || y != 200 {
			t.Errorf("Expected ChildOffset (100, 200), got (%d, %d)", x, y)
		}
	})

	t.Run("SetAndGetChildExtent", func(t *testing.T) {
		xfrm := drawingml.NewTransform2D(0, 0, 1000, 1000)

		// Initially should not be present
		_, _, present := xfrm.ChildExtent()
		if present {
			t.Error("Expected ChildExtent not present initially")
		}

		// Set child extent
		xfrm.SetChildExtent(800, 900)

		cx, cy, present := xfrm.ChildExtent()
		if !present {
			t.Error("Expected ChildExtent to be present after SetChildExtent")
		}
		if cx != 800 || cy != 900 {
			t.Errorf("Expected ChildExtent (800, 900), got (%d, %d)", cx, cy)
		}
	})

	t.Run("BothChildOffsetAndExtent", func(t *testing.T) {
		xfrm := drawingml.NewTransform2D(0, 0, 1000, 1000)

		xfrm.SetChildOffset(50, 60)
		xfrm.SetChildExtent(700, 800)

		x, y, presentOff := xfrm.ChildOffset()
		cx, cy, presentExt := xfrm.ChildExtent()

		if !presentOff || !presentExt {
			t.Error("Expected both ChildOffset and ChildExtent to be present")
		}

		if x != 50 || y != 60 {
			t.Errorf("Expected ChildOffset (50, 60), got (%d, %d)", x, y)
		}

		if cx != 700 || cy != 800 {
			t.Errorf("Expected ChildExtent (700, 800), got (%d, %d)", cx, cy)
		}
	})
}
