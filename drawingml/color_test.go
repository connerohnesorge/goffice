package drawingml

import (
	"strings"
	"testing"
)

const (
	testColorRed = "FF0000"
)

func TestRgbColor(t *testing.T) {
	t.Run("NewRgbColor", func(t *testing.T) {
		color := NewRgbColor(testColorRed)
		if color.Value() != testColorRed {
			t.Errorf(
				"Expected FF0000, got %s",
				color.Value(),
			)
		}
	})

	t.Run(
		"NewRgbColorFromRGB",
		func(t *testing.T) {
			color := NewRgbColorFromRGB(
				255,
				128,
				64,
			)
			// The value should be uppercase hex
			val := color.Value()
			if val != "FF8040" {
				t.Errorf(
					"Expected FF8040, got %s",
					val,
				)
			}
		},
	)

	t.Run("RGB method", func(t *testing.T) {
		color := NewRgbColor("1A2B3C")
		r, g, b := color.RGB()
		if r != 0x1A || g != 0x2B || b != 0x3C {
			t.Errorf(
				"Expected (26, 43, 60), got (%d, %d, %d)",
				r,
				g,
				b,
			)
		}
	})

	t.Run(
		"AddTransformations",
		func(t *testing.T) {
			color := NewRgbColor(testColorRed)
			color.AddTint(50000)
			color.AddAlpha(75000)

			xml := color.OuterXml()
			if !strings.Contains(xml, "tint") {
				t.Error(
					"Expected tint transformation in XML",
				)
			}
			if !strings.Contains(xml, "alpha") {
				t.Error(
					"Expected alpha transformation in XML",
				)
			}
		},
	)

	t.Run("Clone", func(t *testing.T) {
		color := NewRgbColor("00FF00")
		cloned := color.Clone()
		clone, ok := cloned.(*RgbColor)
		if !ok {
			t.Fatal(
				"Clone did not return *RgbColor",
			)
		}
		if clone.Value() != color.Value() {
			t.Error(
				"Clone should have same value",
			)
		}
	})
}

func TestHslColor(t *testing.T) {
	t.Run("NewHslColor", func(t *testing.T) {
		// 180 degrees hue, 50% sat, 75% lum
		color := NewHslColor(
			10800000,
			50000,
			75000,
		)
		if color.Hue() != 10800000 {
			t.Errorf(
				"Expected hue 10800000, got %d",
				color.Hue(),
			)
		}
		if color.Saturation() != 50000 {
			t.Errorf(
				"Expected sat 50000, got %d",
				color.Saturation(),
			)
		}
		if color.Luminance() != 75000 {
			t.Errorf(
				"Expected lum 75000, got %d",
				color.Luminance(),
			)
		}
	})

	t.Run(
		"NewHslColorFromDegrees",
		func(t *testing.T) {
			color := NewHslColorFromDegrees(
				180.0,
				50.0,
				75.0,
			)
			if color.HueDegrees() != 180.0 {
				t.Errorf(
					"Expected 180 degrees, got %f",
					color.HueDegrees(),
				)
			}
			if color.SaturationPercent() != 50.0 {
				t.Errorf(
					"Expected 50%%, got %f",
					color.SaturationPercent(),
				)
			}
			if color.LuminancePercent() != 75.0 {
				t.Errorf(
					"Expected 75%%, got %f",
					color.LuminancePercent(),
				)
			}
		},
	)
}

func TestSchemeColor(t *testing.T) {
	t.Run("NewSchemeColor", func(t *testing.T) {
		color := NewSchemeColor(
			SchemeColorAccent1,
		)
		if color.Value() != SchemeColorAccent1 {
			t.Errorf(
				"Expected accent1, got %s",
				color.Value(),
			)
		}
	})

	t.Run("XML output", func(t *testing.T) {
		color := NewSchemeColor(SchemeColorDark1)
		xml := color.OuterXml()
		if !strings.Contains(xml, "schemeClr") {
			t.Error("Expected schemeClr element")
		}
		if !strings.Contains(xml, "dk1") {
			t.Error("Expected dk1 value")
		}
	})

	t.Run(
		"AddTransformations",
		func(t *testing.T) {
			color := NewSchemeColor(
				SchemeColorAccent2,
			)
			color.AddShade(50000)
			color.AddLuminanceModulation(110000)

			xml := color.OuterXml()
			if !strings.Contains(xml, "shade") {
				t.Error(
					"Expected shade transformation",
				)
			}
			if !strings.Contains(xml, "lumMod") {
				t.Error(
					"Expected lumMod transformation",
				)
			}
		},
	)
}

func TestPresetColor(t *testing.T) {
	t.Run("NewPresetColor", func(t *testing.T) {
		color := NewPresetColor(PresetColorRed)
		if color.Value() != PresetColorRed {
			t.Errorf(
				"Expected red, got %s",
				color.Value(),
			)
		}
	})

	t.Run("XML output", func(t *testing.T) {
		color := NewPresetColor(PresetColorBlue)
		xml := color.OuterXml()
		if !strings.Contains(xml, "prstClr") {
			t.Error("Expected prstClr element")
		}
		if !strings.Contains(xml, "blue") {
			t.Error("Expected blue value")
		}
	})
}

func TestSystemColor(t *testing.T) {
	t.Run("NewSystemColor", func(t *testing.T) {
		color := NewSystemColor(SystemColorWindow)
		if color.Value() != SystemColorWindow {
			t.Errorf(
				"Expected window, got %s",
				color.Value(),
			)
		}
	})

	t.Run(
		"NewSystemColorWithLastColor",
		func(t *testing.T) {
			color := NewSystemColorWithLastColor(
				SystemColorWindowText,
				"000000",
			)
			if color.LastColor() != "000000" {
				t.Errorf(
					"Expected 000000, got %s",
					color.LastColor(),
				)
			}
		},
	)
}
