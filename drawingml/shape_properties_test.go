package drawingml

import (
	"strings"
	"testing"
)

//nolint:revive // cyclomatic: comprehensive test coverage requires many subtests
func TestShapeProperties(t *testing.T) {
	t.Run(
		"NewShapeProperties",
		func(t *testing.T) {
			sp := NewShapeProperties()
			xml := sp.OuterXml()
			if !strings.Contains(xml, "spPr") {
				t.Error("Expected spPr element")
			}
		},
	)

	t.Run("BlackWhiteMode", func(t *testing.T) {
		sp := NewShapeProperties()
		sp.SetBlackWhiteMode(BWModeGray)
		if sp.BlackWhiteMode() != BWModeGray {
			t.Errorf(
				"Expected gray, got %s",
				sp.BlackWhiteMode(),
			)
		}
	})

	t.Run("Transform", func(t *testing.T) {
		sp := NewShapeProperties()
		sp.SetTransformValues(
			InchesToEmu(1),
			InchesToEmu(2),
			InchesToEmu(3),
			InchesToEmu(4),
		)

		xfrm := sp.Transform()
		if xfrm == nil {
			t.Fatal("Expected transform")
		}
		off := xfrm.Offset()
		if off.X != InchesToEmu(1) {
			t.Errorf(
				"Expected %d, got %d",
				InchesToEmu(1),
				off.X,
			)
		}
	})

	t.Run("PresetGeometry", func(t *testing.T) {
		sp := NewShapeProperties()
		sp.SetPresetShape(ShapeTypeRectangle)

		geom := sp.PresetGeometry()
		if geom == nil {
			t.Fatal("Expected geometry")
		}
		if geom.Preset() != ShapeTypeRectangle {
			t.Errorf(
				"Expected rect, got %s",
				geom.Preset(),
			)
		}
	})

	t.Run("CustomGeometry", func(t *testing.T) {
		sp := NewShapeProperties()
		cg := NewCustomGeometry()
		sp.SetCustomGeometry(cg)

		if sp.CustomGeometry() == nil {
			t.Error("Expected custom geometry")
		}
		if sp.PresetGeometry() != nil {
			t.Error(
				"Should not have preset geometry",
			)
		}
	})

	t.Run("NoFill", func(t *testing.T) {
		sp := NewShapeProperties()
		sp.SetNoFill()

		if !sp.NoFill() {
			t.Error("Expected no fill")
		}
	})

	t.Run("SolidFill", func(t *testing.T) {
		sp := NewShapeProperties()
		sp.SetSolidFillColor("FF0000")

		sf := sp.SolidFill()
		if sf == nil {
			t.Fatal("Expected solid fill")
		}
		rgb := sf.RgbColor()
		if rgb == nil || rgb.Value() != "FF0000" {
			t.Error("Expected red color")
		}
	})

	t.Run(
		"SolidFillSchemeColor",
		func(t *testing.T) {
			sp := NewShapeProperties()
			sp.SetSolidFillSchemeColor(
				SchemeColorAccent2,
			)

			sf := sp.SolidFill()
			if sf == nil {
				t.Fatal("Expected solid fill")
			}
			sc := sf.SchemeColor()
			if sc == nil ||
				sc.Value() != SchemeColorAccent2 {
				t.Error("Expected accent2 color")
			}
		},
	)

	t.Run("GradientFill", func(t *testing.T) {
		sp := NewShapeProperties()
		gf := NewLinearGradientFill(5400000)
		gf.AddRgbStop(0, "FFFFFF")
		gf.AddRgbStop(100000, "000000")
		sp.SetGradientFill(gf)

		if sp.GradientFill() == nil {
			t.Error("Expected gradient fill")
		}
	})

	t.Run("PatternFill", func(t *testing.T) {
		sp := NewShapeProperties()
		pf := NewPatternFill(PatternDiagonalCross)
		sp.SetPatternFill(pf)

		if sp.PatternFill() == nil {
			t.Error("Expected pattern fill")
		}
	})

	t.Run("BlipFill", func(t *testing.T) {
		sp := NewShapeProperties()
		bf := NewBlipFillWithEmbed("rId1")
		bf.SetStretch()
		sp.SetBlipFill(bf)

		if sp.BlipFill() == nil {
			t.Error("Expected blip fill")
		}
	})

	t.Run("GroupFill", func(t *testing.T) {
		sp := NewShapeProperties()
		sp.SetGroupFill()

		if !sp.GroupFill() {
			t.Error("Expected group fill")
		}
	})

	t.Run("Outline", func(t *testing.T) {
		sp := NewShapeProperties()
		sp.SetOutlineColor(
			"000000",
			PointsToEmu(1.0),
		)

		ln := sp.Outline()
		if ln == nil {
			t.Fatal("Expected outline")
		}
		if ln.WidthPoints() != 1.0 {
			t.Errorf(
				"Expected 1pt, got %f",
				ln.WidthPoints(),
			)
		}
	})

	t.Run(
		"OutlineSchemeColor",
		func(t *testing.T) {
			sp := NewShapeProperties()
			sp.SetOutlineSchemeColor(
				SchemeColorDark1,
				PointsToEmu(2.0),
			)

			ln := sp.Outline()
			if ln == nil {
				t.Fatal("Expected outline")
			}
			sf := ln.SolidFill()
			if sf == nil {
				t.Fatal("Expected solid fill")
			}
		},
	)

	t.Run("NoOutline", func(t *testing.T) {
		sp := NewShapeProperties()
		sp.SetNoOutline()

		ln := sp.Outline()
		if ln == nil {
			t.Fatal("Expected outline element")
		}
		xml := ln.OuterXml()
		if !strings.Contains(xml, "noFill") {
			t.Error("Expected noFill in outline")
		}
	})

	t.Run("CompleteShape", func(t *testing.T) {
		sp := NewShapeProperties()
		sp.SetTransformValues(
			0,
			0,
			InchesToEmu(2),
			InchesToEmu(1),
		)
		sp.SetPresetShape(ShapeTypeRoundRectangle)
		sp.SetSolidFillColor("4472C4")
		sp.SetOutlineColor(
			"2F5496",
			PointsToEmu(1.0),
		)

		xml := sp.OuterXml()
		if !strings.Contains(xml, "xfrm") {
			t.Error("Expected transform")
		}
		if !strings.Contains(xml, "prstGeom") {
			t.Error("Expected geometry")
		}
		if !strings.Contains(xml, "solidFill") {
			t.Error("Expected fill")
		}
		if !strings.Contains(xml, "ln") {
			t.Error("Expected outline")
		}
	})

	t.Run("Clone", func(t *testing.T) {
		sp := NewShapeProperties()
		sp.SetPresetShape(ShapeTypeEllipse)
		sp.SetSolidFillColor("00FF00")

		cloned := sp.Clone()
		clone, ok := cloned.(*ShapeProperties)
		if !ok {
			t.Fatalf(
				"Clone() returned %T, expected *ShapeProperties",
				cloned,
			)
		}
		if clone.PresetGeometry() == nil {
			t.Error("Clone should have geometry")
		}
		if clone.SolidFill() == nil {
			t.Error("Clone should have fill")
		}
	})
}

func TestShapeStyle(t *testing.T) {
	t.Run("NewShapeStyle", func(t *testing.T) {
		style := NewShapeStyle()
		xml := style.OuterXml()
		if !strings.Contains(xml, "style") {
			t.Error("Expected style element")
		}
	})

	t.Run("SetReferences", func(t *testing.T) {
		style := NewShapeStyle()
		style.SetLineReference(
			1,
			SchemeColorAccent1,
		)
		style.SetFillReference(
			2,
			SchemeColorAccent2,
		)
		style.SetEffectReference(
			0,
			SchemeColorAccent1,
		)
		style.SetFontReference(
			"minor",
			SchemeColorDark1,
		)

		xml := style.OuterXml()
		if !strings.Contains(xml, "lnRef") {
			t.Error("Expected line reference")
		}
		if !strings.Contains(xml, "fillRef") {
			t.Error("Expected fill reference")
		}
		if !strings.Contains(xml, "effectRef") {
			t.Error("Expected effect reference")
		}
		if !strings.Contains(xml, "fontRef") {
			t.Error("Expected font reference")
		}
	})
}

func TestStyleMatrixReference(t *testing.T) {
	t.Run(
		"NewStyleMatrixReference",
		func(t *testing.T) {
			ref := NewStyleMatrixReference(
				"lnRef",
				2,
			)
			if ref.Index() != 2 {
				t.Errorf(
					"Expected 2, got %d",
					ref.Index(),
				)
			}
		},
	)

	t.Run("SetSchemeColor", func(t *testing.T) {
		ref := NewStyleMatrixReference(
			"fillRef",
			1,
		)
		ref.SetSchemeColor(SchemeColorAccent3)

		xml := ref.OuterXml()
		if !strings.Contains(xml, "schemeClr") {
			t.Error("Expected scheme color")
		}
		if !strings.Contains(xml, "accent3") {
			t.Error("Expected accent3 value")
		}
	})
}
