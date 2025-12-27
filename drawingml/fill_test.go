package drawingml

import (
	"strings"
	"testing"
)

func TestNoFill(t *testing.T) {
	noFill := NewNoFill()
	xml := noFill.OuterXml()
	if !strings.Contains(xml, "noFill") {
		t.Error("Expected noFill element")
	}
}

func TestSolidFill(t *testing.T) {
	t.Run(
		"NewSolidFillWithRgb",
		func(t *testing.T) {
			fill := NewSolidFillWithRgb("FF0000")
			rgb := fill.RgbColor()
			if rgb == nil {
				t.Fatal(
					"Expected RGB color to be set",
				)
			}
			if rgb.Value() != "FF0000" {
				t.Errorf(
					"Expected FF0000, got %s",
					rgb.Value(),
				)
			}
		},
	)

	t.Run(
		"NewSolidFillWithSchemeColor",
		func(t *testing.T) {
			fill := NewSolidFillWithSchemeColor(
				SchemeColorAccent1,
			)
			sc := fill.SchemeColor()
			if sc == nil {
				t.Fatal(
					"Expected scheme color to be set",
				)
			}
			if sc.Value() != SchemeColorAccent1 {
				t.Errorf(
					"Expected accent1, got %s",
					sc.Value(),
				)
			}
		},
	)

	t.Run("SetPresetColor", func(t *testing.T) {
		fill := NewSolidFill()
		fill.SetPresetColor(PresetColorGreen)
		xml := fill.OuterXml()
		if !strings.Contains(xml, "prstClr") {
			t.Error(
				"Expected preset color element",
			)
		}
		if !strings.Contains(xml, "green") {
			t.Error("Expected green value")
		}
	})

	t.Run("XML output", func(t *testing.T) {
		fill := NewSolidFillWithRgb("0000FF")
		xml := fill.OuterXml()
		if !strings.Contains(xml, "solidFill") {
			t.Error("Expected solidFill element")
		}
		if !strings.Contains(xml, "srgbClr") {
			t.Error("Expected srgbClr element")
		}
		if !strings.Contains(xml, "0000FF") {
			t.Error("Expected color value")
		}
	})
}

func TestGradientFill(t *testing.T) {
	t.Run(
		"NewLinearGradientFill",
		func(t *testing.T) {
			fill := NewLinearGradientFill(
				5400000,
			) // 90 degrees
			xml := fill.OuterXml()
			if !strings.Contains(
				xml,
				"gradFill",
			) {
				t.Error(
					"Expected gradFill element",
				)
			}
			if !strings.Contains(xml, "lin") {
				t.Error("Expected lin element")
			}
			if !strings.Contains(xml, "5400000") {
				t.Error("Expected angle value")
			}
		},
	)

	t.Run("AddStops", func(t *testing.T) {
		fill := NewGradientFill()
		fill.AddRgbStop(0, "FFFFFF")
		fill.AddRgbStop(100000, "000000")
		fill.SetLinear(5400000, false)

		xml := fill.OuterXml()
		if !strings.Contains(xml, "gsLst") {
			t.Error("Expected gsLst element")
		}
		if !strings.Contains(xml, "gs") {
			t.Error(
				"Expected gs (gradient stop) element",
			)
		}
	})

	t.Run(
		"AddSchemeColorStop",
		func(t *testing.T) {
			fill := NewGradientFill()
			fill.AddSchemeColorStop(
				0,
				SchemeColorLight1,
			)
			fill.AddSchemeColorStop(
				100000,
				SchemeColorDark1,
			)

			xml := fill.OuterXml()
			if !strings.Contains(
				xml,
				"schemeClr",
			) {
				t.Error(
					"Expected scheme color elements",
				)
			}
		},
	)

	t.Run("SetPath", func(t *testing.T) {
		fill := NewGradientFill()
		fill.SetPath(PathShapeCircle)

		xml := fill.OuterXml()
		if !strings.Contains(xml, "path") {
			t.Error("Expected path element")
		}
		if !strings.Contains(xml, "circle") {
			t.Error("Expected circle path type")
		}
	})

	t.Run("RotateWithShape", func(t *testing.T) {
		fill := NewGradientFill()
		if !fill.RotateWithShape() {
			t.Error("Default should be true")
		}
		fill.SetRotateWithShape(false)
		if fill.RotateWithShape() {
			t.Error(
				"Should be false after setting",
			)
		}
	})
}

func TestPatternFill(t *testing.T) {
	t.Run("NewPatternFill", func(t *testing.T) {
		fill := NewPatternFill(PatternHorizontal)
		if fill.Preset() != PatternHorizontal {
			t.Errorf(
				"Expected horz, got %s",
				fill.Preset(),
			)
		}
	})

	t.Run("SetColors", func(t *testing.T) {
		fill := NewPatternFill(
			PatternDiagonalCross,
		)
		fill.SetForegroundColor("FF0000")
		fill.SetBackgroundColor("FFFFFF")

		xml := fill.OuterXml()
		if !strings.Contains(xml, "fgClr") {
			t.Error("Expected foreground color")
		}
		if !strings.Contains(xml, "bgClr") {
			t.Error("Expected background color")
		}
	})
}

func TestBlipFill(t *testing.T) {
	t.Run(
		"NewBlipFillWithEmbed",
		func(t *testing.T) {
			fill := NewBlipFillWithEmbed("rId1")
			if fill.Embed() != "rId1" {
				t.Errorf(
					"Expected rId1, got %s",
					fill.Embed(),
				)
			}
		},
	)

	t.Run("SetStretch", func(t *testing.T) {
		fill := NewBlipFill()
		fill.SetEmbed("rId2")
		fill.SetStretch()

		xml := fill.OuterXml()
		if !strings.Contains(xml, "stretch") {
			t.Error("Expected stretch element")
		}
		if !strings.Contains(xml, "fillRect") {
			t.Error("Expected fillRect element")
		}
	})

	t.Run("SetTile", func(t *testing.T) {
		fill := NewBlipFill()
		fill.SetEmbed("rId3")
		fill.SetTile(RectAlignCenter, TileFlipXY)

		xml := fill.OuterXml()
		if !strings.Contains(xml, "tile") {
			t.Error("Expected tile element")
		}
	})
}

func TestGroupFill(t *testing.T) {
	grpFill := NewGroupFill()
	xml := grpFill.OuterXml()
	if !strings.Contains(xml, "grpFill") {
		t.Error("Expected grpFill element")
	}
}

func TestGradientStop(t *testing.T) {
	t.Run("Position", func(t *testing.T) {
		stop := NewGradientStop(50000)
		if stop.Position() != 50000 {
			t.Errorf(
				"Expected 50000, got %d",
				stop.Position(),
			)
		}
		if stop.PositionPercent() != 50.0 {
			t.Errorf(
				"Expected 50.0, got %f",
				stop.PositionPercent(),
			)
		}
	})

	t.Run("WithRgb", func(t *testing.T) {
		stop := NewGradientStopWithRgb(
			25000,
			"AABBCC",
		)
		xml := stop.OuterXml()
		if !strings.Contains(xml, "AABBCC") {
			t.Error("Expected color value")
		}
	})
}
