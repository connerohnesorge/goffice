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
		chartSeriesConfig{
			"North",
			barCategoryRange,
			barNorthRange,
			0,
		},
	)
	addBarChartSeries(
		barChart,
		chartSeriesConfig{
			"South",
			barCategoryRange,
			barSouthRange,
			1,
		},
	)
	addBarChartSeries(
		barChart,
		chartSeriesConfig{
			"East",
			barCategoryRange,
			barEastRange,
			2,
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
		chartSeriesConfig{
			"New York",
			lineCategoryRange,
			lineNYRange,
			0,
		},
	)
	addLineChartSeries(
		lineChart,
		chartSeriesConfig{
			"Los Angeles",
			lineCategoryRange,
			lineLARange,
			1,
		},
	)
	addLineChartSeries(
		lineChart,
		chartSeriesConfig{
			"Chicago",
			lineCategoryRange,
			lineChicagoRange,
			2,
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

// createPieChartSlide creates a slide with a pie chart showing market share.
// The chart displays market share distribution across five companies.
func createPieChartSlide(
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
		"Market Share Distribution",
	)

	// Create chart with title
	chartSpace, chart := createChartWithTitle(
		"2024 Market Share by Company",
	)
	plotArea := chart.PlotArea()

	// Add pie chart
	pieChart := plotArea.AddPieChart()

	// Add single series with category labels and values
	series := pieChart.AddSeries(0, 0)
	cat := drawingml.NewCategoryAxisData()
	cat.SetStringReference(pieCategoryRange)
	series.SetCategoryAxisData(cat)

	vals := drawingml.NewValues()
	vals.SetNumberReference(pieValuesRange)
	series.SetValues(vals)

	// Add legend on the right side
	legend := drawingml.NewLegendWithPosition(
		drawingml.LegendPositionRight,
	)
	chart.SetLegend(legend)

	return addChartToSlide(slidePart, chartSpace)
}

// createAreaChartSlide creates a slide with an area chart showing cumulative revenue.
// The chart displays stacked revenue growth for three product lines.
func createAreaChartSlide(
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
		"Cumulative Revenue Growth",
	)

	// Create chart with title
	chartSpace, chart := createChartWithTitle(
		"Revenue Growth by Product Line",
	)
	plotArea := chart.PlotArea()

	// Add stacked area chart
	areaChart := plotArea.AddAreaChart(
		drawingml.GroupingStacked,
	)

	// Add series for each product (note: seriesName not used for area charts)
	addAreaChartSeries(
		areaChart,
		chartSeriesConfig{
			"",
			areaCategoryRange,
			areaProductARange,
			0,
		},
	)
	addAreaChartSeries(
		areaChart,
		chartSeriesConfig{
			"",
			areaCategoryRange,
			areaProductBRange,
			1,
		},
	)
	addAreaChartSeries(
		areaChart,
		chartSeriesConfig{
			"",
			areaCategoryRange,
			areaProductCRange,
			2,
		},
	)

	// Configure axes
	areaChart.AddAxisID(categoryAxisID)
	areaChart.AddAxisID(valueAxisID)
	addChartAxes(plotArea)

	// Add legend at the top
	legend := drawingml.NewLegendWithPosition(
		drawingml.LegendPositionTop,
	)
	chart.SetLegend(legend)

	return addChartToSlide(slidePart, chartSpace)
}

// createScatterChartSlide creates a slide with a scatter chart showing correlation.
// The chart displays the relationship between price and demand for two product lines.
func createScatterChartSlide(
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
		"Price vs Demand Analysis",
	)

	// Create chart with title
	chartSpace, chart := createChartWithTitle(
		"Product Demand vs Price Point",
	)
	plotArea := chart.PlotArea()

	// Add scatter chart with lines and markers
	scatterChart := plotArea.AddScatterChart(
		drawingml.ScatterStyleLineMarker,
	)

	// Add series for each product line
	addScatterChartSeries(
		scatterChart,
		scatterSeriesConfig{
			scatterXRange,
			scatterProduct1YRange,
			0,
		},
	)
	addScatterChartSeries(
		scatterChart,
		scatterSeriesConfig{
			scatterXRange,
			scatterProduct2YRange,
			1,
		},
	)

	// Configure axes (both value axes for scatter charts)
	scatterChart.AddAxisID(categoryAxisID)
	scatterChart.AddAxisID(valueAxisID)
	addValueAxes(plotArea)

	// Add legend at the bottom
	legend := drawingml.NewLegendWithPosition(
		drawingml.LegendPositionBottom,
	)
	chart.SetLegend(legend)

	return addChartToSlide(slidePart, chartSpace)
}

// createDoughnutChartSlide creates a slide with a doughnut chart showing budget allocation.
// The chart displays annual budget distribution across six departments.
func createDoughnutChartSlide(
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
		"Annual Budget Allocation",
	)

	// Create chart with title
	chartSpace, chart := createChartWithTitle(
		"2024 Department Budget Distribution",
	)
	plotArea := chart.PlotArea()

	// Add doughnut chart with 50% hole size
	doughnutChart := plotArea.AddDoughnutChart()
	doughnutChart.SetHoleSize(doughnutHoleSize)

	// Add single series with category labels and values
	series := doughnutChart.AddSeries(0, 0)
	cat := drawingml.NewCategoryAxisData()
	cat.SetStringReference(doughnutCategoryRange)
	series.SetCategoryAxisData(cat)

	vals := drawingml.NewValues()
	vals.SetNumberReference(doughnutValuesRange)
	series.SetValues(vals)

	// Add legend on the right side
	legend := drawingml.NewLegendWithPosition(
		drawingml.LegendPositionRight,
	)
	chart.SetLegend(legend)

	return addChartToSlide(slidePart, chartSpace)
}
