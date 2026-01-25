package drawingml

import (
	"math"
	"testing"
)

// TestPoint2D tests Point2D type.
func TestPoint2D(t *testing.T) {
	t.Run("NewPoint2D", func(t *testing.T) {
		p := NewPoint2D(100, 200)
		if p.X != 100 || p.Y != 200 {
			t.Errorf(
				"NewPoint2D(100, 200) = {%d, %d}, want {100, 200}",
				p.X,
				p.Y,
			)
		}
	})

	t.Run("IsZero", func(t *testing.T) {
		zero := Point2D{}
		if !zero.IsZero() {
			t.Error(
				"Zero point should return true for IsZero()",
			)
		}
		nonZero := NewPoint2D(1, 0)
		if nonZero.IsZero() {
			t.Error(
				"Non-zero point should return false for IsZero()",
			)
		}
	})

	t.Run("Add", func(t *testing.T) {
		p1 := NewPoint2D(100, 200)
		p2 := NewPoint2D(50, 75)
		result := p1.Add(p2)
		if result.X != 150 || result.Y != 275 {
			t.Errorf(
				"Add() = {%d, %d}, want {150, 275}",
				result.X,
				result.Y,
			)
		}
	})

	t.Run("Sub", func(t *testing.T) {
		p1 := NewPoint2D(100, 200)
		p2 := NewPoint2D(50, 75)
		result := p1.Sub(p2)
		if result.X != 50 || result.Y != 125 {
			t.Errorf(
				"Sub() = {%d, %d}, want {50, 125}",
				result.X,
				result.Y,
			)
		}
	})
}

// TestOffset tests Offset type.
func TestOffset(t *testing.T) {
	t.Run("NewOffset", func(t *testing.T) {
		o := NewOffset(100, 200)
		if o.X != 100 || o.Y != 200 {
			t.Errorf(
				"NewOffset(100, 200) = {%d, %d}, want {100, 200}",
				o.X,
				o.Y,
			)
		}
	})

	t.Run("ToPoint2D", func(t *testing.T) {
		o := NewOffset(100, 200)
		p := o.ToPoint2D()
		if p.X != 100 || p.Y != 200 {
			t.Errorf(
				"ToPoint2D() = {%d, %d}, want {100, 200}",
				p.X,
				p.Y,
			)
		}
	})
}

// TestExtent tests Extent type.
func TestExtent(t *testing.T) {
	t.Run("NewExtent", func(t *testing.T) {
		e := NewExtent(100, 200)
		if e.Cx != 100 || e.Cy != 200 {
			t.Errorf(
				"NewExtent(100, 200) = {%d, %d}, want {100, 200}",
				e.Cx,
				e.Cy,
			)
		}
	})

	t.Run("Width and Height", func(t *testing.T) {
		e := NewExtent(100, 200)
		if e.Width() != 100 {
			t.Errorf(
				"Width() = %d, want 100",
				e.Width(),
			)
		}
		if e.Height() != 200 {
			t.Errorf(
				"Height() = %d, want 200",
				e.Height(),
			)
		}
	})

	t.Run("AspectRatio", func(t *testing.T) {
		e := NewExtent(200, 100)
		ratio := e.AspectRatio()
		if ratio != 2.0 {
			t.Errorf(
				"AspectRatio() = %f, want 2.0",
				ratio,
			)
		}
	})

	t.Run(
		"AspectRatio ZeroHeight",
		func(t *testing.T) {
			e := NewExtent(100, 0)
			ratio := e.AspectRatio()
			if ratio != 0 {
				t.Errorf(
					"AspectRatio() with zero height = %f, want 0",
					ratio,
				)
			}
		},
	)

	t.Run("ScaleToWidth", func(t *testing.T) {
		e := NewExtent(200, 100)
		scaled := e.ScaleToWidth(400)
		if scaled.Cx != 400 || scaled.Cy != 200 {
			t.Errorf(
				"ScaleToWidth(400) = {%d, %d}, want {400, 200}",
				scaled.Cx,
				scaled.Cy,
			)
		}
	})

	t.Run("ScaleToHeight", func(t *testing.T) {
		e := NewExtent(200, 100)
		scaled := e.ScaleToHeight(200)
		if scaled.Cx != 400 || scaled.Cy != 200 {
			t.Errorf(
				"ScaleToHeight(200) = {%d, %d}, want {400, 200}",
				scaled.Cx,
				scaled.Cy,
			)
		}
	})

	t.Run("Scale", func(t *testing.T) {
		e := NewExtent(100, 50)
		scaled := e.Scale(2.0)
		if scaled.Cx != 200 || scaled.Cy != 100 {
			t.Errorf(
				"Scale(2.0) = {%d, %d}, want {200, 100}",
				scaled.Cx,
				scaled.Cy,
			)
		}
	})
}

// TestEffectExtent tests EffectExtent type.
func TestEffectExtent(t *testing.T) {
	t.Run("NewEffectExtent", func(t *testing.T) {
		ee := NewEffectExtent(10, 20, 30, 40)
		if ee.Left != 10 || ee.Top != 20 ||
			ee.Right != 30 ||
			ee.Bottom != 40 {
			t.Error(
				"NewEffectExtent values incorrect",
			)
		}
	})

	t.Run(
		"NewEffectExtentUniform",
		func(t *testing.T) {
			ee := NewEffectExtentUniform(25)
			if ee.Left != 25 || ee.Top != 25 ||
				ee.Right != 25 ||
				ee.Bottom != 25 {
				t.Error(
					"NewEffectExtentUniform values not uniform",
				)
			}
		},
	)

	t.Run(
		"TotalWidth and TotalHeight",
		func(t *testing.T) {
			ee := NewEffectExtent(10, 20, 30, 40)
			if ee.TotalWidth() != 40 {
				t.Errorf(
					"TotalWidth() = %d, want 40",
					ee.TotalWidth(),
				)
			}
			if ee.TotalHeight() != 60 {
				t.Errorf(
					"TotalHeight() = %d, want 60",
					ee.TotalHeight(),
				)
			}
		},
	)
}

// TestTransform2D tests Transform2D element.
//
//nolint:revive // cyclomatic: test function with many subtests
func TestTransform2D(t *testing.T) {
	t.Run("NewTransform2D", func(t *testing.T) {
		xfrm := NewTransform2D(100, 200, 300, 400)
		if xfrm == nil {
			t.Fatal("NewTransform2D returned nil")
		}

		offset := xfrm.Offset()
		if offset.X != 100 || offset.Y != 200 {
			t.Errorf(
				"Offset = {%d, %d}, want {100, 200}",
				offset.X,
				offset.Y,
			)
		}

		extent := xfrm.Extent()
		if extent.Cx != 300 || extent.Cy != 400 {
			t.Errorf(
				"Extent = {%d, %d}, want {300, 400}",
				extent.Cx,
				extent.Cy,
			)
		}
	})

	t.Run(
		"NewTransform2DWithOffset",
		func(t *testing.T) {
			off := NewOffset(100, 200)
			ext := NewExtent(300, 400)
			xfrm := NewTransform2DWithOffset(
				off,
				ext,
			)

			if xfrm.Offset() != off {
				t.Errorf("Offset mismatch")
			}
			if xfrm.Extent() != ext {
				t.Errorf("Extent mismatch")
			}
		},
	)

	t.Run("Rotation", func(t *testing.T) {
		xfrm := NewTransform2D(0, 0, 100, 100)

		// Default rotation should be 0
		if xfrm.Rotation() != 0 {
			t.Errorf(
				"Default rotation = %d, want 0",
				xfrm.Rotation(),
			)
		}

		// Set rotation to 90 degrees (5400000 units)
		xfrm.SetRotation(5400000)
		if xfrm.Rotation() != 5400000 {
			t.Errorf(
				"Rotation = %d, want 5400000",
				xfrm.Rotation(),
			)
		}

		// Test degrees methods
		xfrm.SetRotationDegrees(45)
		degrees := xfrm.RotationDegrees()
		if math.Abs(degrees-45) > 0.001 {
			t.Errorf(
				"RotationDegrees = %f, want 45",
				degrees,
			)
		}
	})

	t.Run("FlipH", func(t *testing.T) {
		xfrm := NewTransform2D(0, 0, 100, 100)

		if xfrm.FlipH() {
			t.Error(
				"Default FlipH should be false",
			)
		}

		xfrm.SetFlipH(true)
		if !xfrm.FlipH() {
			t.Error(
				"FlipH should be true after setting",
			)
		}

		xfrm.SetFlipH(false)
		if xfrm.FlipH() {
			t.Error(
				"FlipH should be false after unsetting",
			)
		}
	})

	t.Run("FlipV", func(t *testing.T) {
		xfrm := NewTransform2D(0, 0, 100, 100)

		if xfrm.FlipV() {
			t.Error(
				"Default FlipV should be false",
			)
		}

		xfrm.SetFlipV(true)
		if !xfrm.FlipV() {
			t.Error(
				"FlipV should be true after setting",
			)
		}
	})

	t.Run("Width and Height", func(t *testing.T) {
		xfrm := NewTransform2D(0, 0, 300, 400)

		if xfrm.Width() != 300 {
			t.Errorf(
				"Width = %d, want 300",
				xfrm.Width(),
			)
		}
		if xfrm.Height() != 400 {
			t.Errorf(
				"Height = %d, want 400",
				xfrm.Height(),
			)
		}

		xfrm.SetWidth(500)
		xfrm.SetHeight(600)

		if xfrm.Width() != 500 {
			t.Errorf(
				"Width after set = %d, want 500",
				xfrm.Width(),
			)
		}
		if xfrm.Height() != 600 {
			t.Errorf(
				"Height after set = %d, want 600",
				xfrm.Height(),
			)
		}
	})

	t.Run(
		"OffsetX and OffsetY",
		func(t *testing.T) {
			xfrm := NewTransform2D(
				100,
				200,
				300,
				400,
			)

			if xfrm.OffsetX() != 100 {
				t.Errorf(
					"OffsetX = %d, want 100",
					xfrm.OffsetX(),
				)
			}
			if xfrm.OffsetY() != 200 {
				t.Errorf(
					"OffsetY = %d, want 200",
					xfrm.OffsetY(),
				)
			}

			xfrm.SetOffsetX(150)
			xfrm.SetOffsetY(250)

			if xfrm.OffsetX() != 150 {
				t.Errorf(
					"OffsetX after set = %d, want 150",
					xfrm.OffsetX(),
				)
			}
			if xfrm.OffsetY() != 250 {
				t.Errorf(
					"OffsetY after set = %d, want 250",
					xfrm.OffsetY(),
				)
			}
		},
	)

	t.Run("MoveTo", func(t *testing.T) {
		xfrm := NewTransform2D(0, 0, 100, 100)
		xfrm.MoveTo(500, 600)

		if xfrm.OffsetX() != 500 ||
			xfrm.OffsetY() != 600 {
			t.Errorf(
				"After MoveTo: offset = {%d, %d}, want {500, 600}",
				xfrm.OffsetX(),
				xfrm.OffsetY(),
			)
		}
	})

	t.Run("MoveBy", func(t *testing.T) {
		xfrm := NewTransform2D(100, 200, 100, 100)
		xfrm.MoveBy(50, -50)

		if xfrm.OffsetX() != 150 ||
			xfrm.OffsetY() != 150 {
			t.Errorf(
				"After MoveBy: offset = {%d, %d}, want {150, 150}",
				xfrm.OffsetX(),
				xfrm.OffsetY(),
			)
		}
	})

	t.Run("Resize", func(t *testing.T) {
		xfrm := NewTransform2D(0, 0, 100, 100)
		xfrm.Resize(500, 600)

		if xfrm.Width() != 500 ||
			xfrm.Height() != 600 {
			t.Errorf(
				"After Resize: size = {%d, %d}, want {500, 600}",
				xfrm.Width(),
				xfrm.Height(),
			)
		}
	})

	t.Run("ScaleBy", func(t *testing.T) {
		xfrm := NewTransform2D(0, 0, 100, 200)
		xfrm.ScaleBy(2.0)

		if xfrm.Width() != 200 ||
			xfrm.Height() != 400 {
			t.Errorf(
				"After ScaleBy(2.0): size = {%d, %d}, want {200, 400}",
				xfrm.Width(),
				xfrm.Height(),
			)
		}
	})

	t.Run("Clone", func(t *testing.T) {
		xfrm := NewTransform2D(100, 200, 300, 400)
		xfrm.SetRotation(5400000)
		xfrm.SetFlipH(true)

		clone := xfrm.Clone().(*Transform2D)

		if clone.OffsetX() != 100 ||
			clone.OffsetY() != 200 {
			t.Errorf("Clone offset mismatch")
		}
		if clone.Width() != 300 ||
			clone.Height() != 400 {
			t.Errorf("Clone extent mismatch")
		}
		if clone.Rotation() != 5400000 {
			t.Errorf("Clone rotation mismatch")
		}
		if !clone.FlipH() {
			t.Errorf("Clone FlipH mismatch")
		}

		// Verify it's a deep copy
		clone.SetWidth(999)
		if xfrm.Width() == 999 {
			t.Error("Clone is not a deep copy")
		}
	})
}
