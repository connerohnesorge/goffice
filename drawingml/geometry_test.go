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
		// Verify the XML structure contains proper adjustment list
		if !strings.Contains(xml, "<a:avLst>") {
			t.Error("Expected avLst element with proper namespace")
		}
		if !strings.Contains(xml, "<a:gd") {
			t.Error("Expected gd element with proper namespace")
		}
		// Verify the gd element has both name and formula attributes
		if !strings.Contains(xml, `name="adj"`) {
			t.Error("Expected gd element to have name attribute with value 'adj'")
		}
		if !strings.Contains(xml, `fmla="val 16667"`) {
			t.Error("Expected gd element to have fmla attribute with value 'val 16667'")
		}
	})

	t.Run("SetAdjustValue", func(t *testing.T) {
		geom := NewPresetGeometry(
			ShapeTypeRoundRectangle,
		)
		geom.SetAdjustValue("adj", 16667)

		xml := geom.OuterXml()
		// Verify the XML contains the adjustment value properly formatted
		if !strings.Contains(xml, `fmla="val 16667"`) {
			t.Error("Expected gd element to have fmla attribute with value 'val 16667'")
		}
		// Also verify it's within an avLst element
		if !strings.Contains(xml, "<a:avLst>") {
			t.Error("Expected avLst element with proper namespace")
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
		// Verify the XML structure
		if !strings.HasPrefix(xml, "<a:prstGeom") {
			t.Error("Expected XML to start with <a:prstGeom")
		}
		// Verify the prst attribute contains the correct shape type
		if !strings.Contains(xml, `prst="star5"`) {
			t.Error("Expected prst attribute with value 'star5'")
		}
		// Verify it's properly closed
		if !strings.HasSuffix(xml, ">") {
			t.Error("Expected properly closed XML element")
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
		// Verify XML structure
		if !strings.HasPrefix(xml, "<a:path") {
			t.Error("Expected XML to start with <a:path")
		}
		// Path should be self-closing when empty
		if !strings.HasSuffix(xml, "/>") {
			t.Error("Expected self-closing path element")
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
		// Verify moveTo element structure
		if !strings.Contains(xml, "<a:moveTo>") {
			t.Error("Expected moveTo element with proper namespace")
		}
		if !strings.Contains(xml, "<a:pt") {
			t.Error("Expected pt element with proper namespace")
		}
		// Verify pt element has correct coordinates
		if !strings.Contains(xml, `x="0"`) {
			t.Error("Expected pt element to have x coordinate of 0")
		}
		if !strings.Contains(xml, `y="0"`) {
			t.Error("Expected pt element to have y coordinate of 0")
		}
	})

	t.Run("AddLineTo", func(t *testing.T) {
		path := NewPath2D()
		path.AddMoveTo(0, 0)
		path.AddLineTo(100000, 100000)
		xml := path.OuterXml()
		// Verify lineTo element structure
		if !strings.Contains(xml, "<a:lnTo>") {
			t.Error("Expected lnTo element with proper namespace")
		}
		// Verify it comes after moveTo (proper order)
		moveToIndex := strings.Index(xml, "<a:moveTo>")
		lnToIndex := strings.Index(xml, "<a:lnTo>")
		if moveToIndex == -1 || lnToIndex == -1 || lnToIndex < moveToIndex {
			t.Error("Expected lnTo to come after moveTo in XML")
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
