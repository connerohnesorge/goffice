package drawingml

import (
	"strings"
	"testing"
)

func TestLineEndProperties(t *testing.T) {
	t.Run("NewHeadEnd", func(t *testing.T) {
		head := NewHeadEnd()
		xml := head.OuterXml()
		if !strings.Contains(xml, "headEnd") {
			t.Error("Expected headEnd element")
		}
	})

	t.Run("NewTailEnd", func(t *testing.T) {
		tail := NewTailEnd()
		xml := tail.OuterXml()
		if !strings.Contains(xml, "tailEnd") {
			t.Error("Expected tailEnd element")
		}
	})

	t.Run(
		"NewHeadEndWithArrow",
		func(t *testing.T) {
			head := NewHeadEndWithArrow(
				LineEndTriangle,
				LineEndSizeLarge,
				LineEndSizeMedium,
			)
			if head.Type() != LineEndTriangle {
				t.Errorf(
					"Expected triangle, got %s",
					head.Type(),
				)
			}
			if head.Width() != LineEndSizeLarge {
				t.Errorf(
					"Expected lg, got %s",
					head.Width(),
				)
			}
		},
	)

	t.Run("DefaultValues", func(t *testing.T) {
		head := NewHeadEnd()
		if head.Type() != LineEndNone {
			t.Error("Default type should be none")
		}
		if head.Width() != LineEndSizeMedium {
			t.Error(
				"Default width should be medium",
			)
		}
		if head.Length() != LineEndSizeMedium {
			t.Error(
				"Default length should be medium",
			)
		}
	})
}

//nolint:revive // cyclomatic: test function with many subtests
func TestLineProperties(t *testing.T) {
	t.Run(
		"NewLineProperties",
		func(t *testing.T) {
			line := NewLineProperties()
			xml := line.OuterXml()
			if !strings.Contains(xml, "ln") {
				t.Error("Expected ln element")
			}
		},
	)

	t.Run("Width", func(t *testing.T) {
		line := NewLinePropertiesWithWidth(
			PointsToEmu(2.0),
		)
		width := line.Width()
		if width != PointsToEmu(2.0) {
			t.Errorf(
				"Expected %d, got %d",
				PointsToEmu(2.0),
				width,
			)
		}
		if line.WidthPoints() != 2.0 {
			t.Errorf(
				"Expected 2.0 points, got %f",
				line.WidthPoints(),
			)
		}
	})

	t.Run("SetWidthPoints", func(t *testing.T) {
		line := NewLineProperties()
		line.SetWidthPoints(1.5)
		if line.WidthPoints() != 1.5 {
			t.Errorf(
				"Expected 1.5, got %f",
				line.WidthPoints(),
			)
		}
	})

	t.Run("CapType", func(t *testing.T) {
		line := NewLineProperties()
		line.SetCapType(LineCapRound)
		if line.CapType() != LineCapRound {
			t.Errorf(
				"Expected rnd, got %s",
				line.CapType(),
			)
		}
	})

	t.Run("CompoundLineType", func(t *testing.T) {
		line := NewLineProperties()
		if line.CompoundLineType() != CompoundLineSingle {
			t.Error("Default should be single")
		}
		line.SetCompoundLineType(
			CompoundLineDouble,
		)
		if line.CompoundLineType() != CompoundLineDouble {
			t.Errorf(
				"Expected dbl, got %s",
				line.CompoundLineType(),
			)
		}
	})

	t.Run("SetSolidFill", func(t *testing.T) {
		line := NewLineProperties()
		line.SetSolidFill("FF0000")
		sf := line.SolidFill()
		if sf == nil {
			t.Fatal("Expected solid fill")
		}
		rgb := sf.RgbColor()
		if rgb == nil || rgb.Value() != "FF0000" {
			t.Error("Expected red color")
		}
	})

	t.Run(
		"SetSolidFillSchemeColor",
		func(t *testing.T) {
			line := NewLineProperties()
			line.SetSolidFillSchemeColor(
				SchemeColorAccent1,
			)
			xml := line.OuterXml()
			if !strings.Contains(
				xml,
				"schemeClr",
			) {
				t.Error("Expected scheme color")
			}
		},
	)

	t.Run("SetNoFill", func(t *testing.T) {
		line := NewLineProperties()
		line.SetNoFill()
		xml := line.OuterXml()
		if !strings.Contains(xml, "noFill") {
			t.Error("Expected noFill element")
		}
	})

	t.Run("PresetDash", func(t *testing.T) {
		line := NewLineProperties()
		if line.PresetDash() != LineDashSolid {
			t.Error("Default should be solid")
		}
		line.SetPresetDash(LineDashDot)
		if line.PresetDash() != LineDashDot {
			t.Errorf(
				"Expected dot, got %s",
				line.PresetDash(),
			)
		}
	})

	t.Run("LineJoin", func(t *testing.T) {
		line := NewLineProperties()
		line.SetRoundJoin()
		if line.LineJoin() != LineJoinRound {
			t.Error("Expected round join")
		}

		line.SetBevelJoin()
		if line.LineJoin() != LineJoinBevel {
			t.Error("Expected bevel join")
		}

		line.SetMiterJoin(400000)
		if line.LineJoin() != LineJoinMiter {
			t.Error("Expected miter join")
		}
	})

	t.Run(
		"HeadEnd and TailEnd",
		func(t *testing.T) {
			line := NewLineProperties()
			head := NewHeadEndWithArrow(
				LineEndArrow,
				LineEndSizeMedium,
				LineEndSizeMedium,
			)
			tail := NewTailEndWithArrow(
				LineEndTriangle,
				LineEndSizeLarge,
				LineEndSizeLarge,
			)
			line.SetHeadEnd(head)
			line.SetTailEnd(tail)

			xml := line.OuterXml()
			if !strings.Contains(xml, "headEnd") {
				t.Error(
					"Expected headEnd element",
				)
			}
			if !strings.Contains(xml, "tailEnd") {
				t.Error(
					"Expected tailEnd element",
				)
			}
		},
	)

	t.Run("SetArrow", func(t *testing.T) {
		line := NewLineProperties()
		line.SetArrow(
			LineEndNone,
			LineEndTriangle,
		)

		if line.HeadEnd() != nil {
			t.Error(
				"Head should not be set for none type",
			)
		}
		if line.TailEnd() == nil {
			t.Error("Tail should be set")
		}
	})

	t.Run("Clone", func(t *testing.T) {
		line := NewLinePropertiesWithWidth(
			PointsToEmu(3.0),
		)
		line.SetSolidFill("00FF00")
		cloned := line.Clone()
		clone, ok := cloned.(*LineProperties)
		if !ok {
			t.Fatal(
				"Clone did not return *LineProperties",
			)
		}
		if clone.Width() != line.Width() {
			t.Error(
				"Clone should have same width",
			)
		}
	})
}
