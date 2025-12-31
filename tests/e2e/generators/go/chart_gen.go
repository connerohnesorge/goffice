// Chart generator for E2E visual testing
//
// This file implements chart generation using the goffice drawingml API.
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/presentation/parts"
	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

// ChartGenerator handles chart creation for test cases
type ChartGenerator struct {
	slidePart *parts.SlidePart
	chartSpec *framework.ChartSpec
	verbose   bool
}

// NewChartGenerator creates a new chart generator
func NewChartGenerator(
	slidePart *parts.SlidePart,
	chartSpec *framework.ChartSpec,
	verbose bool,
) *ChartGenerator {
	return &ChartGenerator{
		slidePart: slidePart,
		chartSpec: chartSpec,
		verbose:   verbose,
	}
}

// Generate creates the chart based on the spec
func (g *ChartGenerator) Generate() (*parts.ChartPart, error) {
	// Create chart part
	chartPart, err := g.slidePart.AddChartPart()
	if err != nil {
		return nil, fmt.Errorf(
			"add chart part: %w",
			err,
		)
	}

	chartSpace := chartPart.ChartSpace()
	if chartSpace == nil {
		return nil, fmt.Errorf(
			"chart space is nil",
		)
	}

	chart := chartSpace.Chart()
	if chart == nil {
		return nil, fmt.Errorf("chart is nil")
	}

	// Set chart title if provided
	if g.chartSpec.Title != "" {
		if g.verbose {
			log.Printf(
				"      Setting chart title: %s",
				g.chartSpec.Title,
			)
		}
		title := drawingml.NewTitleWithText(
			g.chartSpec.Title,
		)
		chart.SetTitle(title)
	}

	// Set legend if specified
	if g.chartSpec.Legend != nil &&
		g.chartSpec.Legend.ShowLegend {
		if g.verbose {
			log.Printf(
				"      Setting legend position: %s",
				g.chartSpec.Legend.Position,
			)
		}
		legendPos := mapLegendPosition(
			g.chartSpec.Legend.Position,
		)
		legend := drawingml.NewLegendWithPosition(
			legendPos,
		)
		chart.SetLegend(legend)
	}

	// Get plot area
	plotArea := chart.PlotArea()
	if plotArea == nil {
		return nil, fmt.Errorf("plot area is nil")
	}

	// Generate chart type-specific content
	if err := g.generateChartType(plotArea); err != nil {
		return nil, fmt.Errorf(
			"generate chart type: %w",
			err,
		)
	}

	return chartPart, nil
}

// generateChartType creates the specific chart type
func (g *ChartGenerator) generateChartType(
	plotArea *drawingml.PlotArea,
) error {
	switch g.chartSpec.Type {
	case framework.ChartTypeBarClustered,
		framework.ChartTypeBarStacked,
		framework.ChartTypeBarPercentStacked,
		framework.ChartTypeColumn:
		return g.generateBarChart(plotArea)
	case framework.ChartTypeLine,
		framework.ChartTypeLineSmooth,
		framework.ChartTypeLineMarkers:
		return g.generateLineChart(plotArea)
	case framework.ChartTypePie:
		return g.generatePieChart(plotArea)
	case framework.ChartTypeDoughnut:
		return g.generateDoughnutChart(plotArea)
	case framework.ChartTypeArea,
		framework.ChartTypeAreaStacked:
		return g.generateAreaChart(plotArea)
	case framework.ChartTypeScatter,
		framework.ChartTypeScatterSmooth,
		framework.ChartTypeScatterLine:
		return g.generateScatterChart(plotArea)
	default:
		return fmt.Errorf(
			"unsupported chart type: %s",
			g.chartSpec.Type,
		)
	}
}

// generateBarChart creates a bar or column chart
func (g *ChartGenerator) generateBarChart(
	plotArea *drawingml.PlotArea,
) error {
	if g.verbose {
		log.Printf(
			"      Generating bar/column chart",
		)
	}

	// Determine direction and grouping
	var direction drawingml.BarDirectionValue
	var grouping drawingml.BarGroupingValue

	// Column charts are vertical, bar charts are horizontal
	if g.chartSpec.Type == framework.ChartTypeColumn {
		direction = drawingml.BarDirectionCol
	} else {
		direction = drawingml.BarDirectionBar
	}

	// Map grouping
	switch g.chartSpec.Type {
	case framework.ChartTypeBarStacked:
		grouping = drawingml.BarGroupingStacked
	case framework.ChartTypeBarPercentStacked:
		grouping = drawingml.BarGroupingPercentStacked
	default:
		grouping = drawingml.BarGroupingClustered
	}

	// Create bar chart
	barChart := drawingml.NewBarChart(
		direction,
		grouping,
	)

	// Add series
	for i, seriesData := range g.chartSpec.Data.Series {
		if g.verbose {
			log.Printf(
				"        Adding series %d: %s",
				i,
				seriesData.Name,
			)
		}

		series := barChart.AddSeries(
			uint32(i),
			uint32(i),
		)

		// Set series name
		if seriesData.Name != "" {
			seriesText := drawingml.NewSeriesTextWithValue(
				seriesData.Name,
			)
			series.SetSeriesText(seriesText)
		}

		// Set series color if specified
		if seriesData.Color != "" {
			shapeProps := drawingml.NewChartShapeProperties()
			shapeProps.SetSolidFill(
				seriesData.Color,
			)
			series.SetShapeProperties(shapeProps)
		}

		// Set category axis data (x-axis labels)
		if len(g.chartSpec.Data.Categories) > 0 {
			catData := g.createCategoryData(
				g.chartSpec.Data.Categories,
			)
			series.SetCategoryAxisData(catData)
		}

		// Set values (y-axis data)
		if len(seriesData.Values) > 0 {
			values := g.createNumericValues(
				seriesData.Values,
			)
			series.SetValues(values)
		}
	}

	// Add axes if specified
	if g.chartSpec.Axes != nil {
		catAxisID := uint32(100000001)
		valAxisID := uint32(100000002)

		// Add axis IDs to chart
		barChart.AddAxisID(catAxisID)
		barChart.AddAxisID(valAxisID)

		// Create category axis
		catAxis := drawingml.NewCategoryAxis(
			catAxisID,
			valAxisID,
		)
		if g.chartSpec.Axes.Category.Title != "" {
			catAxisTitle := drawingml.NewTitleWithText(
				g.chartSpec.Axes.Category.Title,
			)
			catAxis.SetTitle(catAxisTitle)
		}
		catAxis.SetAxisPosition(
			drawingml.AxisPositionBottom,
		)
		plotArea.AppendChild(catAxis)

		// Create value axis
		valAxis := drawingml.NewValueAxis(
			valAxisID,
			catAxisID,
		)
		if g.chartSpec.Axes.Value.Title != "" {
			valAxisTitle := drawingml.NewTitleWithText(
				g.chartSpec.Axes.Value.Title,
			)
			valAxis.SetTitle(valAxisTitle)
		}
		valAxis.SetAxisPosition(
			drawingml.AxisPositionLeft,
		)
		plotArea.AppendChild(valAxis)
	}

	plotArea.AppendChild(barChart)

	return nil
}

// generateLineChart creates a line chart
func (g *ChartGenerator) generateLineChart(
	plotArea *drawingml.PlotArea,
) error {
	if g.verbose {
		log.Printf("      Generating line chart")
	}

	// Determine grouping
	var grouping drawingml.GroupingValue
	if g.chartSpec.Type == framework.ChartTypeLineSmooth {
		grouping = drawingml.GroupingStandard
	} else {
		grouping = drawingml.GroupingStandard
	}

	// Create line chart
	lineChart := drawingml.NewLineChart(grouping)

	// Add series
	for i, seriesData := range g.chartSpec.Data.Series {
		if g.verbose {
			log.Printf(
				"        Adding series %d: %s",
				i,
				seriesData.Name,
			)
		}

		series := lineChart.AddSeries(
			uint32(i),
			uint32(i),
		)

		// Set series name
		if seriesData.Name != "" {
			seriesText := drawingml.NewSeriesTextWithValue(
				seriesData.Name,
			)
			series.SetSeriesText(seriesText)
		}

		// Set series color if specified
		if seriesData.Color != "" {
			shapeProps := drawingml.NewChartShapeProperties()
			shapeProps.SetSolidFill(
				seriesData.Color,
			)
			series.SetShapeProperties(shapeProps)
		}

		// Set smooth if specified
		if seriesData.Smooth ||
			g.chartSpec.Type == framework.ChartTypeLineSmooth {
			series.SetSmooth(true)
		}

		// Set category axis data
		if len(g.chartSpec.Data.Categories) > 0 {
			catData := g.createCategoryData(
				g.chartSpec.Data.Categories,
			)
			series.SetCategoryAxisData(catData)
		}

		// Set values
		if len(seriesData.Values) > 0 {
			values := g.createNumericValues(
				seriesData.Values,
			)
			series.SetValues(values)
		}
	}

	// Add axes if specified
	if g.chartSpec.Axes != nil {
		catAxisID := uint32(100000001)
		valAxisID := uint32(100000002)

		lineChart.AddAxisID(catAxisID)
		lineChart.AddAxisID(valAxisID)

		// Create category axis
		catAxis := drawingml.NewCategoryAxis(
			catAxisID,
			valAxisID,
		)
		if g.chartSpec.Axes.Category.Title != "" {
			catAxisTitle := drawingml.NewTitleWithText(
				g.chartSpec.Axes.Category.Title,
			)
			catAxis.SetTitle(catAxisTitle)
		}
		catAxis.SetAxisPosition(
			drawingml.AxisPositionBottom,
		)
		plotArea.AppendChild(catAxis)

		// Create value axis
		valAxis := drawingml.NewValueAxis(
			valAxisID,
			catAxisID,
		)
		if g.chartSpec.Axes.Value.Title != "" {
			valAxisTitle := drawingml.NewTitleWithText(
				g.chartSpec.Axes.Value.Title,
			)
			valAxis.SetTitle(valAxisTitle)
		}
		valAxis.SetAxisPosition(
			drawingml.AxisPositionLeft,
		)
		plotArea.AppendChild(valAxis)
	}

	plotArea.AppendChild(lineChart)

	return nil
}

// generatePieChart creates a pie chart
func (g *ChartGenerator) generatePieChart(
	plotArea *drawingml.PlotArea,
) error {
	if g.verbose {
		log.Printf("      Generating pie chart")
	}

	pieChart := drawingml.NewPieChart()

	// Pie charts typically have one series
	if len(g.chartSpec.Data.Series) > 0 {
		seriesData := g.chartSpec.Data.Series[0]

		series := pieChart.AddSeries(0, 0)

		// Note: Pie charts don't have SetSeriesText on the series
		// The series name is set through category labels

		// Set category axis data (slice labels)
		if len(g.chartSpec.Data.Categories) > 0 {
			catData := g.createCategoryData(
				g.chartSpec.Data.Categories,
			)
			series.SetCategoryAxisData(catData)
		}

		// Set values (slice sizes)
		if len(seriesData.Values) > 0 {
			values := g.createNumericValues(
				seriesData.Values,
			)
			series.SetValues(values)
		}
	}

	plotArea.AppendChild(pieChart)

	return nil
}

// generateDoughnutChart creates a doughnut chart
func (g *ChartGenerator) generateDoughnutChart(
	plotArea *drawingml.PlotArea,
) error {
	if g.verbose {
		log.Printf(
			"      Generating doughnut chart",
		)
	}

	doughnutChart := drawingml.NewDoughnutChart()

	// Doughnut charts are similar to pie charts
	if len(g.chartSpec.Data.Series) > 0 {
		seriesData := g.chartSpec.Data.Series[0]

		series := doughnutChart.AddSeries(0, 0)

		// Note: Doughnut charts don't have SetSeriesText on the series
		// The series name is set through category labels

		// Set category axis data
		if len(g.chartSpec.Data.Categories) > 0 {
			catData := g.createCategoryData(
				g.chartSpec.Data.Categories,
			)
			series.SetCategoryAxisData(catData)
		}

		// Set values
		if len(seriesData.Values) > 0 {
			values := g.createNumericValues(
				seriesData.Values,
			)
			series.SetValues(values)
		}
	}

	plotArea.AppendChild(doughnutChart)

	return nil
}

// generateAreaChart creates an area chart
func (g *ChartGenerator) generateAreaChart(
	plotArea *drawingml.PlotArea,
) error {
	if g.verbose {
		log.Printf("      Generating area chart")
	}

	// Determine grouping
	var grouping drawingml.GroupingValue
	if g.chartSpec.Type == framework.ChartTypeAreaStacked {
		grouping = drawingml.GroupingStacked
	} else {
		grouping = drawingml.GroupingStandard
	}

	areaChart := drawingml.NewAreaChart(grouping)

	// Add series
	for i, seriesData := range g.chartSpec.Data.Series {
		series := areaChart.AddSeries(
			uint32(i),
			uint32(i),
		)

		// Note: AreaChartSeries doesn't have SetSeriesText/SetShapeProperties methods
		// These need to be added to the drawingml package if needed

		// Set category axis data
		if len(g.chartSpec.Data.Categories) > 0 {
			catData := g.createCategoryData(
				g.chartSpec.Data.Categories,
			)
			series.SetCategoryAxisData(catData)
		}

		// Set values
		if len(seriesData.Values) > 0 {
			values := g.createNumericValues(
				seriesData.Values,
			)
			series.SetValues(values)
		}
	}

	// Add axes if specified
	if g.chartSpec.Axes != nil {
		catAxisID := uint32(100000001)
		valAxisID := uint32(100000002)

		areaChart.AddAxisID(catAxisID)
		areaChart.AddAxisID(valAxisID)

		catAxis := drawingml.NewCategoryAxis(
			catAxisID,
			valAxisID,
		)
		if g.chartSpec.Axes.Category.Title != "" {
			catAxisTitle := drawingml.NewTitleWithText(
				g.chartSpec.Axes.Category.Title,
			)
			catAxis.SetTitle(catAxisTitle)
		}
		catAxis.SetAxisPosition(
			drawingml.AxisPositionBottom,
		)
		plotArea.AppendChild(catAxis)

		valAxis := drawingml.NewValueAxis(
			valAxisID,
			catAxisID,
		)
		if g.chartSpec.Axes.Value.Title != "" {
			valAxisTitle := drawingml.NewTitleWithText(
				g.chartSpec.Axes.Value.Title,
			)
			valAxis.SetTitle(valAxisTitle)
		}
		valAxis.SetAxisPosition(
			drawingml.AxisPositionLeft,
		)
		plotArea.AppendChild(valAxis)
	}

	plotArea.AppendChild(areaChart)

	return nil
}

// generateScatterChart creates a scatter chart
func (g *ChartGenerator) generateScatterChart(
	plotArea *drawingml.PlotArea,
) error {
	if g.verbose {
		log.Printf(
			"      Generating scatter chart",
		)
	}

	// Determine scatter style
	var style drawingml.ScatterStyleValue
	switch g.chartSpec.Type {
	case framework.ChartTypeScatterSmooth:
		style = drawingml.ScatterStyleSmooth
	case framework.ChartTypeScatterLine:
		style = drawingml.ScatterStyleLine
	default:
		style = drawingml.ScatterStyleMarker
	}

	scatterChart := drawingml.NewScatterChart(
		style,
	)

	// Add series
	for i, seriesData := range g.chartSpec.Data.Series {
		series := scatterChart.AddSeries(
			uint32(i),
			uint32(i),
		)

		// Note: ScatterChartSeries doesn't have SetSeriesText/SetShapeProperties methods
		// These need to be added to the drawingml package if needed

		// Set X values
		if len(seriesData.XValues) > 0 {
			xValues := g.createXValues(
				seriesData.XValues,
			)
			series.SetXValues(xValues)
		}

		// Set Y values
		if len(seriesData.Values) > 0 {
			yValues := g.createYValues(
				seriesData.Values,
			)
			series.SetYValues(yValues)
		}
	}

	// Add axes if specified
	if g.chartSpec.Axes != nil {
		xAxisID := uint32(100000001)
		yAxisID := uint32(100000002)

		scatterChart.AddAxisID(xAxisID)
		scatterChart.AddAxisID(yAxisID)

		// For scatter charts, both axes are value axes
		xAxis := drawingml.NewValueAxis(
			xAxisID,
			yAxisID,
		)
		if g.chartSpec.Axes.Category.Title != "" {
			xAxisTitle := drawingml.NewTitleWithText(
				g.chartSpec.Axes.Category.Title,
			)
			xAxis.SetTitle(xAxisTitle)
		}
		xAxis.SetAxisPosition(
			drawingml.AxisPositionBottom,
		)
		plotArea.AppendChild(xAxis)

		yAxis := drawingml.NewValueAxis(
			yAxisID,
			xAxisID,
		)
		if g.chartSpec.Axes.Value.Title != "" {
			yAxisTitle := drawingml.NewTitleWithText(
				g.chartSpec.Axes.Value.Title,
			)
			yAxis.SetTitle(yAxisTitle)
		}
		yAxis.SetAxisPosition(
			drawingml.AxisPositionLeft,
		)
		plotArea.AppendChild(yAxis)
	}

	plotArea.AppendChild(scatterChart)

	return nil
}

// Helper methods for creating chart data structures

// createCategoryData creates category axis data
func (g *ChartGenerator) createCategoryData(
	categories []string,
) *drawingml.CategoryAxisData {
	catData := drawingml.NewCategoryAxisData()

	// Create a formula reference (placeholder)
	formula := "Sheet1!$A$2:$A$" + strconv.Itoa(
		len(categories)+1,
	)

	// Set string reference with cached values
	strRef := drawingml.NewStringReferenceWithCache(
		formula,
		categories,
	)
	catData.RemoveAllChildren()
	catData.AppendChild(strRef)

	return catData
}

// createNumericValues creates numeric values
func (g *ChartGenerator) createNumericValues(
	values []float64,
) *drawingml.Values {
	vals := drawingml.NewValues()

	// Create a formula reference (placeholder)
	formula := "Sheet1!$B$2:$B$" + strconv.Itoa(
		len(values)+1,
	)

	// Set number reference with cached values
	numRef := drawingml.NewNumberReferenceWithCache(
		formula,
		values,
	)
	vals.RemoveAllChildren()
	vals.AppendChild(numRef)

	return vals
}

// createXValues creates X values for scatter charts
func (g *ChartGenerator) createXValues(
	values []float64,
) *drawingml.XValues {
	xVals := drawingml.NewXValues()

	formula := "Sheet1!$B$2:$B$" + strconv.Itoa(
		len(values)+1,
	)
	numRef := drawingml.NewNumberReferenceWithCache(
		formula,
		values,
	)
	xVals.RemoveAllChildren()
	xVals.AppendChild(numRef)

	return xVals
}

// createYValues creates Y values for scatter charts
func (g *ChartGenerator) createYValues(
	values []float64,
) *drawingml.YValues {
	yVals := drawingml.NewYValues()

	formula := "Sheet1!$C$2:$C$" + strconv.Itoa(
		len(values)+1,
	)
	numRef := drawingml.NewNumberReferenceWithCache(
		formula,
		values,
	)
	yVals.RemoveAllChildren()
	yVals.AppendChild(numRef)

	return yVals
}

// mapLegendPosition maps framework legend position to drawingml legend position
func mapLegendPosition(
	pos framework.LegendPosition,
) drawingml.LegendPositionValue {
	switch pos {
	case framework.LegendPositionTop:
		return drawingml.LegendPositionTop
	case framework.LegendPositionBottom:
		return drawingml.LegendPositionBottom
	case framework.LegendPositionLeft:
		return drawingml.LegendPositionLeft
	case framework.LegendPositionRight:
		return drawingml.LegendPositionRight
	case framework.LegendPositionTopRight:
		return drawingml.LegendPositionTopRight
	default:
		return drawingml.LegendPositionRight
	}
}
