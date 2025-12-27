package drawingml

import (
	"strings"
	"testing"
)

func TestPresetGeometry(t *testing.T) {
	t.Run(
		"NewPresetGeometry",
		func(t *testing.T) {
			geom := NewPresetGeometry(
				ShapeTypeRectangle,
			)
			if geom.Preset() != ShapeTypeRectangle {
				t.Errorf(
					"Expected rect, got %s",
					geom.Preset(),
				)
			}
		},
	)

	t.Run("SetPreset", func(t *testing.T) {
		geom := NewPresetGeometry(
			ShapeTypeRectangle,
		)
		geom.SetPreset(ShapeTypeEllipse)
		if geom.Preset() != ShapeTypeEllipse {
			t.Errorf(
				"Expected ellipse, got %s",
				geom.Preset(),
			)
		}
	})

	t.Run("AddAdjustValue", func(t *testing.T) {
		geom := NewPresetGeometry(
			ShapeTypeRoundRectangle,
		)
		geom.AddAdjustValue("adj", "val 16667")

		xml := geom.OuterXml()
		if !strings.Contains(xml, "avLst") {
			t.Error("Expected avLst element")
		}
		if !strings.Contains(xml, "gd") {
			t.Error("Expected gd element")
		}
	})

	t.Run("SetAdjustValue", func(t *testing.T) {
		geom := NewPresetGeometry(
			ShapeTypeRoundRectangle,
		)
		geom.SetAdjustValue("adj", 16667)

		xml := geom.OuterXml()
		if !strings.Contains(xml, "val 16667") {
			t.Error("Expected value formula")
		}
	})

	t.Run(
		"ClearAdjustValues",
		func(t *testing.T) {
			geom := NewPresetGeometry(
				ShapeTypeRoundRectangle,
			)
			geom.SetAdjustValue("adj", 16667)
			geom.ClearAdjustValues()

			xml := geom.OuterXml()
			if strings.Contains(xml, "avLst") {
				t.Error(
					"Expected avLst to be removed",
				)
			}
		},
	)

	t.Run("XML output", func(t *testing.T) {
		geom := NewPresetGeometry(ShapeTypeStar5)
		xml := geom.OuterXml()
		if !strings.Contains(xml, "prstGeom") {
			t.Error("Expected prstGeom element")
		}
		if !strings.Contains(xml, "star5") {
			t.Error("Expected star5 preset")
		}
	})

	t.Run("Clone", func(t *testing.T) {
		geom := NewPresetGeometry(
			ShapeTypeTriangle,
		)
		cloned := geom.Clone()
		clone, ok := cloned.(*PresetGeometry)
		if !ok {
			t.Fatal(
				"Clone did not return *PresetGeometry",
			)
		}
		if clone.Preset() != geom.Preset() {
			t.Error(
				"Clone should have same preset",
			)
		}
	})
}

func TestShapeGuide(t *testing.T) {
	t.Run("NewShapeGuide", func(t *testing.T) {
		guide := NewShapeGuide("adj", "val 50000")
		if guide.Name() != "adj" {
			t.Errorf(
				"Expected adj, got %s",
				guide.Name(),
			)
		}
		if guide.Formula() != "val 50000" {
			t.Errorf(
				"Expected val 50000, got %s",
				guide.Formula(),
			)
		}
	})

	t.Run(
		"SetName and SetFormula",
		func(t *testing.T) {
			guide := NewShapeGuide("x", "0")
			guide.SetName("y")
			guide.SetFormula("+/ adj 100000 2")
			if guide.Name() != "y" {
				t.Error("Expected y")
			}
			if guide.Formula() != "+/ adj 100000 2" {
				t.Error("Expected formula")
			}
		},
	)
}

func TestPath2D(t *testing.T) {
	t.Run("NewPath2D", func(t *testing.T) {
		path := NewPath2D()
		xml := path.OuterXml()
		if !strings.Contains(xml, "path") {
			t.Error("Expected path element")
		}
	})

	t.Run(
		"NewPath2DWithSize",
		func(t *testing.T) {
			path := NewPath2DWithSize(
				100000,
				200000,
			)
			if path.Width() != 100000 {
				t.Errorf(
					"Expected 100000, got %d",
					path.Width(),
				)
			}
			if path.Height() != 200000 {
				t.Errorf(
					"Expected 200000, got %d",
					path.Height(),
				)
			}
		},
	)

	t.Run("Fill and Stroke", func(t *testing.T) {
		path := NewPath2D()
		if path.Fill() != PathFillNorm {
			t.Error("Default fill should be norm")
		}
		if !path.Stroke() {
			t.Error(
				"Default stroke should be true",
			)
		}

		path.SetFill(PathFillNone)
		path.SetStroke(false)
		if path.Fill() != PathFillNone {
			t.Error("Expected none fill")
		}
		if path.Stroke() {
			t.Error("Expected false stroke")
		}
	})

	t.Run("AddMoveTo", func(t *testing.T) {
		path := NewPath2D()
		path.AddMoveTo(0, 0)
		xml := path.OuterXml()
		if !strings.Contains(xml, "moveTo") {
			t.Error("Expected moveTo element")
		}
		if !strings.Contains(xml, "pt") {
			t.Error("Expected pt element")
		}
	})

	t.Run("AddLineTo", func(t *testing.T) {
		path := NewPath2D()
		path.AddMoveTo(0, 0)
		path.AddLineTo(100000, 100000)
		xml := path.OuterXml()
		if !strings.Contains(xml, "lnTo") {
			t.Error("Expected lnTo element")
		}
	})

	t.Run("AddArcTo", func(t *testing.T) {
		path := NewPath2D()
		path.AddArcTo(50000, 50000, 0, 5400000)
		xml := path.OuterXml()
		if !strings.Contains(xml, "arcTo") {
			t.Error("Expected arcTo element")
		}
	})

	t.Run("AddQuadBezierTo", func(t *testing.T) {
		path := NewPath2D()
		path.AddQuadBezierTo(
			50000,
			0,
			100000,
			100000,
		)
		xml := path.OuterXml()
		if !strings.Contains(xml, "quadBezTo") {
			t.Error("Expected quadBezTo element")
		}
	})

	t.Run("AddCubicBezierTo", func(t *testing.T) {
		path := NewPath2D()
		path.AddCubicBezierTo(
			25000,
			0,
			75000,
			100000,
			100000,
			50000,
		)
		xml := path.OuterXml()
		if !strings.Contains(xml, "cubicBezTo") {
			t.Error("Expected cubicBezTo element")
		}
	})

	t.Run("AddClose", func(t *testing.T) {
		path := NewPath2D()
		path.AddMoveTo(0, 0)
		path.AddLineTo(100000, 0)
		path.AddLineTo(50000, 100000)
		path.AddClose()
		xml := path.OuterXml()
		if !strings.Contains(xml, "close") {
			t.Error("Expected close element")
		}
	})
}

func TestPathList(t *testing.T) {
	t.Run("NewPathList", func(t *testing.T) {
		pl := NewPathList()
		xml := pl.OuterXml()
		if !strings.Contains(xml, "pathLst") {
			t.Error("Expected pathLst element")
		}
	})

	t.Run("AddPath", func(t *testing.T) {
		pl := NewPathList()
		path := NewPath2D()
		path.AddMoveTo(0, 0)
		path.AddLineTo(100000, 100000)
		pl.AddPath(path)

		xml := pl.OuterXml()
		if !strings.Contains(xml, "moveTo") {
			t.Error("Expected path content")
		}
	})
}

func TestCustomGeometry(t *testing.T) {
	t.Run(
		"NewCustomGeometry",
		func(t *testing.T) {
			geom := NewCustomGeometry()
			xml := geom.OuterXml()
			if !strings.Contains(
				xml,
				"custGeom",
			) {
				t.Error(
					"Expected custGeom element",
				)
			}
			if !strings.Contains(xml, "pathLst") {
				t.Error(
					"Expected pathLst element",
				)
			}
		},
	)

	t.Run("AddGuide", func(t *testing.T) {
		geom := NewCustomGeometry()
		geom.AddGuide("x1", "val 50000")

		xml := geom.OuterXml()
		if !strings.Contains(xml, "gdLst") {
			t.Error("Expected gdLst element")
		}
	})

	t.Run("AddAdjustValue", func(t *testing.T) {
		geom := NewCustomGeometry()
		geom.AddAdjustValue("adj", "val 25000")

		xml := geom.OuterXml()
		if !strings.Contains(xml, "avLst") {
			t.Error("Expected avLst element")
		}
	})

	t.Run("AddPath", func(t *testing.T) {
		geom := NewCustomGeometry()
		path := NewPath2DWithSize(100000, 100000)
		path.AddMoveTo(0, 0)
		path.AddLineTo(100000, 0)
		path.AddLineTo(50000, 100000)
		path.AddClose()
		geom.AddPath(path)

		xml := geom.OuterXml()
		if !strings.Contains(xml, "moveTo") {
			t.Error("Expected path commands")
		}
	})

	t.Run("SetRectangle", func(t *testing.T) {
		geom := NewCustomGeometry()
		geom.SetRectangle("l", "t", "r", "b")

		xml := geom.OuterXml()
		if !strings.Contains(xml, "rect") {
			t.Error("Expected rect element")
		}
	})

	t.Run("Triangle shape", func(t *testing.T) {
		geom := NewCustomGeometry()
		path := NewPath2DWithSize(100000, 100000)
		path.AddMoveTo(
			50000,
			0,
		) // top center
		path.AddLineTo(
			100000,
			100000,
		) // bottom right
		path.AddLineTo(
			0,
			100000,
		) // bottom left
		path.AddClose()
		geom.AddPath(path)

		xml := geom.OuterXml()
		// Should have the triangle path
		if !strings.Contains(xml, "close") {
			t.Error("Expected closed path")
		}
	})
}
