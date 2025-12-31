// Package drawing provides chart rendering tests.
package drawing

import (
	"testing"

	"github.com/connerohnesorge/goffice-pdf/core"
)

func TestChartRenderer(t *testing.T) {
	mockPage := core.NewMockPage()
	ctx := core.NewRenderingContext(
		core.PageSizeLetter.Width,
		core.PageSizeLetter.Height,
	).WithPage(mockPage)

	renderer := NewChartRenderer(ctx)

	// Create test data
	data := ChartData{
		Categories: []string{
			"Q1",
			"Q2",
			"Q3",
			"Q4",
		},
		Series: []ChartSeries{
			{
				Name: "Revenue",
				Values: []float64{
					100,
					120,
					90,
					150,
				},
				Color: NewRGB(0.2, 0.4, 0.8),
			},
			{
				Name: "Expenses",
				Values: []float64{
					80,
					85,
					95,
					100,
				},
				Color: NewRGB(0.8, 0.3, 0.3),
			},
		},
	}

	t.Run("BarChart", func(t *testing.T) {
		err := renderer.RenderBarChart(
			50,
			50,
			400,
			300,
			data,
			false,
		)
		if err != nil {
			t.Errorf(
				"Failed to render bar chart: %v",
				err,
			)
		}
	})

	t.Run("LineChart", func(t *testing.T) {
		err := renderer.RenderLineChart(
			50,
			400,
			400,
			300,
			data,
		)
		if err != nil {
			t.Errorf(
				"Failed to render line chart: %v",
				err,
			)
		}
	})

	t.Run("PieChart", func(t *testing.T) {
		pieData := ChartData{
			Categories: []string{
				"A",
				"B",
				"C",
				"D",
			},
			Series: []ChartSeries{
				{
					Name: "Distribution",
					Values: []float64{
						25,
						30,
						20,
						25,
					},
					Color: NewRGB(0.2, 0.4, 0.8),
				},
			},
		}

		err := renderer.RenderPieChart(
			500,
			50,
			300,
			300,
			pieData,
		)
		if err != nil {
			t.Errorf(
				"Failed to render pie chart: %v",
				err,
			)
		}
	})
}

func TestGenerateChartColors(t *testing.T) {
	colors := generateChartColors(10)
	if len(colors) != 10 {
		t.Errorf(
			"Expected 10 colors, got %d",
			len(colors),
		)
	}

	// Test that colors are valid
	for i, color := range colors {
		if color.R < 0 || color.R > 1 {
			t.Errorf(
				"Color %d has invalid R value: %f",
				i,
				color.R,
			)
		}
		if color.G < 0 || color.G > 1 {
			t.Errorf(
				"Color %d has invalid G value: %f",
				i,
				color.G,
			)
		}
		if color.B < 0 || color.B > 1 {
			t.Errorf(
				"Color %d has invalid B value: %f",
				i,
				color.B,
			)
		}
	}
}
