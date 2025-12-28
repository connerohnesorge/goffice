package main

import (
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/presentation/elements"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// chartSeriesConfig holds configuration for adding a chart series.
type chartSeriesConfig struct {
	seriesName    string
	categoryRange string
	valuesRange   string
	seriesIndex   uint32
}

// scatterSeriesConfig holds configuration for adding a scatter chart series.
type scatterSeriesConfig struct {
	xRange      string
	yRange      string
	seriesIndex uint32
}

// addTitleToSlide adds a title shape to a slide with the given text.
func addTitleToSlide(
	slide *elements.Slide,
	titleText string,
) {
	titleShape := slide.AddShape()
	titleShape.SetPosition(titlePosX, titlePosY)
	titleShape.SetSize(titleWidth, titleHeight)
	titleTb := titleShape.GetOrCreateTextBody()
	titleP := titleTb.AddParagraph(titleText)
	titleRun := titleP.AddRun("")
	titleRun.SetFontSize(titleFontSize)
}

// createChartWithTitle creates a chart space with a title.
func createChartWithTitle(
	title string,
) (*drawingml.ChartSpace, *drawingml.Chart) {
	chartSpace := drawingml.NewChartSpace()
	chart := chartSpace.Chart()
	chartTitle := drawingml.NewTitleWithText(
		title,
	)
	chart.SetTitle(chartTitle)

	return chartSpace, chart
}

// axisResult holds the result of adding axes to a plot area.
type axisResult struct {
	xAxis *drawingml.ValueAxis
	yAxis *drawingml.ValueAxis
}

// addChartAxes adds category and value axes to a plot area.
func addChartAxes(
	plotArea *drawingml.PlotArea,
) (*drawingml.CategoryAxis, *drawingml.ValueAxis) {
	catAxis := plotArea.AddCategoryAxis(
		categoryAxisID,
		valueAxisID,
	)
	catAxis.SetAxisPosition(
		drawingml.AxisPositionBottom,
	)
	valAxis := plotArea.AddValueAxis(
		valueAxisID,
		categoryAxisID,
	)
	valAxis.SetAxisPosition(
		drawingml.AxisPositionLeft,
	)

	return catAxis, valAxis
}

// addValueAxes adds two value axes (for scatter charts) and returns them in a struct.
func addValueAxes(
	plotArea *drawingml.PlotArea,
) axisResult {
	xAxis := plotArea.AddValueAxis(
		categoryAxisID,
		valueAxisID,
	)
	xAxis.SetAxisPosition(
		drawingml.AxisPositionBottom,
	)
	yAxis := plotArea.AddValueAxis(
		valueAxisID,
		categoryAxisID,
	)
	yAxis.SetAxisPosition(
		drawingml.AxisPositionLeft,
	)

	return axisResult{xAxis: xAxis, yAxis: yAxis}
}

// addBarChartSeries adds a series to a bar chart using the provided configuration.
func addBarChartSeries(
	barChart *drawingml.BarChart,
	cfg chartSeriesConfig,
) {
	series := barChart.AddSeries(
		cfg.seriesIndex,
		cfg.seriesIndex,
	)
	series.SetSeriesText(
		drawingml.NewSeriesTextWithValue(
			cfg.seriesName,
		),
	)

	cat := drawingml.NewCategoryAxisData()
	cat.SetStringReference(cfg.categoryRange)
	series.SetCategoryAxisData(cat)

	vals := drawingml.NewValues()
	vals.SetNumberReference(cfg.valuesRange)
	series.SetValues(vals)
}

// addLineChartSeries adds a series to a line chart using the provided configuration.
func addLineChartSeries(
	lineChart *drawingml.LineChart,
	cfg chartSeriesConfig,
) {
	series := lineChart.AddSeries(
		cfg.seriesIndex,
		cfg.seriesIndex,
	)
	series.SetSeriesText(
		drawingml.NewSeriesTextWithValue(
			cfg.seriesName,
		),
	)

	cat := drawingml.NewCategoryAxisData()
	cat.SetStringReference(cfg.categoryRange)
	series.SetCategoryAxisData(cat)

	vals := drawingml.NewValues()
	vals.SetNumberReference(cfg.valuesRange)
	series.SetValues(vals)
}

// addAreaChartSeries adds a series to an area chart using the provided configuration.
func addAreaChartSeries(
	areaChart *drawingml.AreaChart,
	cfg chartSeriesConfig,
) {
	series := areaChart.AddSeries(
		cfg.seriesIndex,
		cfg.seriesIndex,
	)

	cat := drawingml.NewCategoryAxisData()
	cat.SetStringReference(cfg.categoryRange)
	series.SetCategoryAxisData(cat)

	vals := drawingml.NewValues()
	vals.SetNumberReference(cfg.valuesRange)
	series.SetValues(vals)
}

// addScatterChartSeries adds a series to a scatter chart using the provided configuration.
func addScatterChartSeries(
	scatterChart *drawingml.ScatterChart,
	cfg scatterSeriesConfig,
) {
	series := scatterChart.AddSeries(
		cfg.seriesIndex,
		cfg.seriesIndex,
	)

	xVals := drawingml.NewXValues()
	xVals.SetNumberReference(cfg.xRange)
	series.SetXValues(xVals)

	yVals := drawingml.NewYValues()
	yVals.SetNumberReference(cfg.yRange)
	series.SetYValues(yVals)
}

// addChartToSlide adds a chart to a slide by creating a ChartPart and linking it via a GraphicFrame.
func addChartToSlide(
	slidePart *parts.SlidePart,
	chartSpace *drawingml.ChartSpace,
) error {
	// Create a chart part in the slide
	chartPart, err := slidePart.AddChartPart()
	if err != nil {
		return err
	}

	// Replace the default chartSpace with the one we created
	chartPart.SetRootElement(chartSpace)

	// Get the slide and create a graphic frame to display the chart
	slide := slidePart.Slide()
	gf := slide.GetOrCreateShapeTree().
		AddGraphicFrame()

	// Position and size the graphic frame
	gf.SetPosition(chartPosX, chartPosY)
	gf.SetSize(chartWidth, chartHeight)

	// Link the graphic frame to the chart part
	elements.LinkGraphicFrameToChart(
		gf,
		chartPart.RelationshipID(),
	)

	return nil
}
