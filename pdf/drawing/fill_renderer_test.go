// Package drawing provides fill rendering tests.
package drawing

import (
	"testing"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
)

func TestSolidFillRenderer(t *testing.T) {
	mockPage := core.NewMockPage()
	ctx := core.NewRenderingContext(
		core.PageSizeLetter.Width,
		core.PageSizeLetter.Height,
	).WithPage(mockPage)

	path := NewPathBuilder()
	path.Rectangle(0, 0, 100, 100)

	t.Run("RGB Color", func(t *testing.T) {
		fill := drawingml.NewSolidFillWithRgb(
			"FF0000",
		)
		renderer := NewSolidFillRenderer(fill)

		err := renderer.Apply(ctx, path)
		if err != nil {
			t.Errorf(
				"Failed to apply solid fill: %v",
				err,
			)
		}
	})

	t.Run("Scheme Color", func(t *testing.T) {
		fill := drawingml.NewSolidFillWithSchemeColor(
			drawingml.SchemeColorAccent1,
		)
		renderer := NewSolidFillRenderer(fill)

		err := renderer.Apply(ctx, path)
		if err != nil {
			t.Errorf(
				"Failed to apply scheme color fill: %v",
				err,
			)
		}
	})
}

func TestGradientFillRenderer(t *testing.T) {
	mockPage := core.NewMockPage()
	ctx := core.NewRenderingContext(
		core.PageSizeLetter.Width,
		core.PageSizeLetter.Height,
	).WithPage(mockPage)

	path := NewPathBuilder()
	path.Rectangle(0, 0, 100, 100)

	t.Run("Linear Gradient", func(t *testing.T) {
		fill := drawingml.NewLinearGradientFill(0)
		fill.AddRgbStop(0, "0000FF")
		fill.AddRgbStop(100000, "FF0000")

		renderer := NewGradientFillRenderer(fill)
		err := renderer.Apply(ctx, path)
		if err != nil {
			t.Errorf(
				"Failed to apply gradient fill: %v",
				err,
			)
		}
	})
}

func TestPatternFillRenderer(t *testing.T) {
	mockPage := core.NewMockPage()
	ctx := core.NewRenderingContext(
		core.PageSizeLetter.Width,
		core.PageSizeLetter.Height,
	).WithPage(mockPage)

	path := NewPathBuilder()
	path.Rectangle(0, 0, 100, 100)

	t.Run("Pattern Fill", func(t *testing.T) {
		fill := drawingml.NewPatternFill(
			drawingml.PatternPct20,
		)
		fill.SetForegroundColor("000000")
		fill.SetBackgroundColor("FFFFFF")

		renderer := NewPatternFillRenderer(fill)
		err := renderer.Apply(ctx, path)
		if err != nil {
			t.Errorf(
				"Failed to apply pattern fill: %v",
				err,
			)
		}
	})
}

func TestPictureFillRenderer(t *testing.T) {
	mockPage := core.NewMockPage()
	ctx := core.NewRenderingContext(
		core.PageSizeLetter.Width,
		core.PageSizeLetter.Height,
	).WithPage(mockPage)

	path := NewPathBuilder()
	path.Rectangle(0, 0, 100, 100)

	t.Run("Picture Fill", func(t *testing.T) {
		fill := drawingml.NewBlipFillWithEmbed(
			"rId1",
		)
		fill.SetStretch()

		renderer := NewPictureFillRenderer(fill)
		err := renderer.Apply(ctx, path)
		if err != nil {
			t.Errorf(
				"Failed to apply picture fill: %v",
				err,
			)
		}
	})
}

func TestNoFillRenderer(t *testing.T) {
	mockPage := core.NewMockPage()
	ctx := core.NewRenderingContext(
		core.PageSizeLetter.Width,
		core.PageSizeLetter.Height,
	).WithPage(mockPage)

	path := NewPathBuilder()
	path.Rectangle(0, 0, 100, 100)

	renderer := NewNoFillRenderer()
	err := renderer.Apply(ctx, path)
	if err != nil {
		t.Errorf(
			"No fill renderer should not error: %v",
			err,
		)
	}
}

func TestCreateFillRenderer(t *testing.T) {
	sp := drawingml.NewShapeProperties()

	t.Run(
		"No fill by default",
		func(t *testing.T) {
			renderer := CreateFillRenderer(sp)
			if renderer == nil {
				t.Error(
					"Should return a renderer even for no fill",
				)
			}
		},
	)

	t.Run(
		"Nil shape properties",
		func(t *testing.T) {
			renderer := CreateFillRenderer(nil)
			if renderer != nil {
				t.Error(
					"Should return nil for nil shape properties",
				)
			}
		},
	)
}
