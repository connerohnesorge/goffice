package drawing

import (
	"testing"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
)

func TestShapeRenderer(t *testing.T) {
	// Create test render context
	mockPage := core.NewMockPage()
	ctx := core.NewRenderingContext(
		core.PageSizeLetter.Width,
		core.PageSizeLetter.Height,
	).WithPage(mockPage)

	renderer := NewShapeRenderer(ctx)

	// Test preset shape rendering
	t.Run("Rectangle", func(t *testing.T) {
		sp := drawingml.NewShapeProperties()
		sp.SetTransformValues(
			drawingml.PointsToEmu(100),
			drawingml.PointsToEmu(100),
			drawingml.PointsToEmu(200),
			drawingml.PointsToEmu(150),
		)
		sp.SetPresetShape(
			drawingml.ShapeTypeRectangle,
		)

		err := renderer.RenderShapeWithFill(
			sp,
			nil,
			nil,
		)
		if err != nil {
			t.Errorf(
				"Failed to render rectangle: %v",
				err,
			)
		}
	})

	t.Run("Ellipse", func(t *testing.T) {
		sp := drawingml.NewShapeProperties()
		sp.SetTransformValues(
			drawingml.PointsToEmu(100),
			drawingml.PointsToEmu(200),
			drawingml.PointsToEmu(150),
			drawingml.PointsToEmu(150),
		)
		sp.SetPresetShape(
			drawingml.ShapeTypeEllipse,
		)

		err := renderer.RenderShapeWithFill(
			sp,
			nil,
			nil,
		)
		if err != nil {
			t.Errorf(
				"Failed to render ellipse: %v",
				err,
			)
		}
	})

	t.Run("Star", func(t *testing.T) {
		sp := drawingml.NewShapeProperties()
		sp.SetTransformValues(
			drawingml.PointsToEmu(300),
			drawingml.PointsToEmu(100),
			drawingml.PointsToEmu(100),
			drawingml.PointsToEmu(100),
		)
		sp.SetPresetShape(
			drawingml.ShapeTypeStar5,
		)

		err := renderer.RenderShapeWithFill(
			sp,
			nil,
			nil,
		)
		if err != nil {
			t.Errorf(
				"Failed to render star: %v",
				err,
			)
		}
	})
}

func TestPresetGeometries(t *testing.T) {
	path := NewPathBuilder()

	tests := []struct {
		name      string
		shapeType drawingml.ShapeTypeValue
	}{
		{
			"Rectangle",
			drawingml.ShapeTypeRectangle,
		},
		{
			"RoundRectangle",
			drawingml.ShapeTypeRoundRectangle,
		},
		{"Ellipse", drawingml.ShapeTypeEllipse},
		{"Triangle", drawingml.ShapeTypeTriangle},
		{"Diamond", drawingml.ShapeTypeDiamond},
		{"Pentagon", drawingml.ShapeTypePentagon},
		{"Hexagon", drawingml.ShapeTypeHexagon},
		{"Star5", drawingml.ShapeTypeStar5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			geom := drawingml.NewPresetGeometry(
				tt.shapeType,
			)
			err := renderPresetGeometry(
				geom,
				path,
				0,
				0,
				100,
				100,
			)
			if err != nil {
				t.Errorf(
					"Failed to render %s: %v",
					tt.name,
					err,
				)
			}
		})
	}
}
