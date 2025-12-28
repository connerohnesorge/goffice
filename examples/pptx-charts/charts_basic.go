package main

import (
	"fmt"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/presentation"
)

// createBarChartSlide creates a slide with a clustered bar chart showing quarterly sales.
// The chart displays sales data for three regions (North, South, East) across four quarters.
func createBarChartSlide(
	doc *presentation.Document,
) error {
	slidePart, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf(
			errFailedToAddSlide,
			err,
		)
	}
	slide := slidePart.Slide()

	// Add slide title
	addTitleToSlide(
		slide,
		"Quarterly Sales Comparison",
	)

	// Create chart with title
	chartSpace, chart := createChartWithTitle(
		"2024 Quarterly Sales by Region",
	)
	plotArea := chart.PlotArea()

	// Add clustered bar chart
	barChart := plotArea.AddBarChart(
		drawingml.BarDirectionCol,
		drawingml.BarGroupingClustered,
	)

	// Add series for each region
	addBarChartSeries(
		barChart,
		&chartSeriesConfig{
			seriesName:     "North",
			categoryRange:  barCategoryRange,
			categoryValues: barCategories,
			valuesRange:    barNorthRange,
			values:         barNorthData,
			seriesIndex:    0,
		},
	)
	addBarChartSeries(
		barChart,
		&chartSeriesConfig{
			seriesName:     "South",
			categoryRange:  barCategoryRange,
			categoryValues: barCategories,
			valuesRange:    barSouthRange,
			values:         barSouthData,
			seriesIndex:    1,
		},
	)
	addBarChartSeries(
		barChart,
		&chartSeriesConfig{
			seriesName:     "East",
			categoryRange:  barCategoryRange,
			categoryValues: barCategories,
			valuesRange:    barEastRange,
			values:         barEastData,
			seriesIndex:    2,
		},
	)

	// Configure axes
	barChart.AddAxisID(categoryAxisID)
	barChart.AddAxisID(valueAxisID)
	addChartAxes(plotArea)

	// Add legend on the right side
	legend := drawingml.NewLegendWithPosition(
		drawingml.LegendPositionRight,
	)
	chart.SetLegend(legend)

	return addChartToSlide(slidePart, chartSpace)
}

// createLineChartSlide creates a slide with a line chart showing temperature trends.
// The chart displays average monthly temperatures for three cities.
func createLineChartSlide(
	doc *presentation.Document,
) error {
	slidePart, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf(
			errFailedToAddSlide,
			err,
		)
	}
	slide := slidePart.Slide()

	// Add slide title
	addTitleToSlide(
		slide,
		"Monthly Temperature Trends",
	)

	// Create chart with title
	chartSpace, chart := createChartWithTitle(
		"Average Temperature by City (2024)",
	)
	plotArea := chart.PlotArea()

	// Add line chart
	lineChart := plotArea.AddLineChart(
		drawingml.GroupingStandard,
	)

	// Add series for each city
	addLineChartSeries(
		lineChart,
		&chartSeriesConfig{
			seriesName:     "New York",
			categoryRange:  lineCategoryRange,
			categoryValues: lineCategories,
			valuesRange:    lineNYRange,
			values:         lineNYData,
			seriesIndex:    0,
		},
	)
	addLineChartSeries(
		lineChart,
		&chartSeriesConfig{
			seriesName:     "Los Angeles",
			categoryRange:  lineCategoryRange,
			categoryValues: lineCategories,
			valuesRange:    lineLARange,
			values:         lineLAData,
			seriesIndex:    1,
		},
	)
	addLineChartSeries(
		lineChart,
		&chartSeriesConfig{
			seriesName:     "Chicago",
			categoryRange:  lineCategoryRange,
			categoryValues: lineCategories,
			valuesRange:    lineChicagoRange,
			values:         lineChicagoData,
			seriesIndex:    2,
		},
	)

	// Configure axes
	lineChart.AddAxisID(categoryAxisID)
	lineChart.AddAxisID(valueAxisID)
	addChartAxes(plotArea)

	// Add legend at the bottom
	legend := drawingml.NewLegendWithPosition(
		drawingml.LegendPositionBottom,
	)
	chart.SetLegend(legend)

	return addChartToSlide(slidePart, chartSpace)
}
