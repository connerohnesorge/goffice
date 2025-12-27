// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements the high-level Chart API.
package spreadsheet

import (
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// ChartType represents the type of chart.
type ChartType string

const (
	// ChartTypeBar represents a bar chart.
	ChartTypeBar ChartType = "bar"
	// ChartTypeBar3D represents a 3D bar chart.
	ChartTypeBar3D ChartType = "bar3D"
	// ChartTypeLine represents a line chart.
	ChartTypeLine ChartType = "line"
	// ChartTypeLine3D represents a 3D line chart.
	ChartTypeLine3D ChartType = "line3D"
	// ChartTypePie represents a pie chart.
	ChartTypePie ChartType = "pie"
	// ChartTypePie3D represents a 3D pie chart.
	ChartTypePie3D ChartType = "pie3D"
	// ChartTypeDoughnut represents a doughnut chart.
	ChartTypeDoughnut ChartType = "doughnut"
	// ChartTypeArea represents an area chart.
	ChartTypeArea ChartType = "area"
	// ChartTypeArea3D represents a 3D area chart.
	ChartTypeArea3D ChartType = "area3D"
	// ChartTypeScatter represents a scatter chart.
	ChartTypeScatter ChartType = "scatter"
	// ChartTypeBubble represents a bubble chart.
	ChartTypeBubble ChartType = "bubble"
	// ChartTypeStock represents a stock chart.
	ChartTypeStock ChartType = "stock"
	// ChartTypeSurface represents a surface chart.
	ChartTypeSurface ChartType = "surface"
	// ChartTypeSurface3D represents a 3D surface chart.
	ChartTypeSurface3D ChartType = "surface3D"
	// ChartTypeRadar represents a radar chart.
	ChartTypeRadar ChartType = "radar"
)

// Chart represents a high-level wrapper around an Excel chart.
type Chart struct {
	sheet        *Sheet
	drawingsPart *parts.DrawingsPart
	chartPart    *parts.ChartPart
	chartType    ChartType
	title        string
	series       []*ChartSeries
	categoryAxis *ChartAxis
	valueAxis    *ChartAxis
}

// newChart creates a new Chart wrapper and initializes it.
//
//nolint:revive // argument-limit: all parameters are needed for chart creation
func newChart(
	sheet *Sheet,
	drawingsPart *parts.DrawingsPart,
	chartType ChartType,
	_ string, // dataRange - TODO: use for data binding
	_ CellRef, // anchor - TODO: use for positioning
) (*Chart, error) {
	// Add a chart part to the drawings part
	chartPart, err := drawingsPart.AddChartPart()
	if err != nil {
		return nil, err
	}

	c := &Chart{
		sheet:        sheet,
		drawingsPart: drawingsPart,
		chartPart:    chartPart,
		chartType:    chartType,
	}

	// Initialize the chart with default settings
	c.categoryAxis = newChartAxis(c, "category")
	c.valueAxis = newChartAxis(c, "value")

	// Create the chart elements based on chartType
	chartSpace := chartPart.ChartSpace()
	if chartSpace == nil {
		return nil, nil
	}

	chart := chartSpace.Chart()
	if chart == nil {
		return nil, nil
	}

	plotArea := chart.PlotArea()
	if plotArea == nil {
		plotArea = drawingml.NewPlotArea()
		chart.SetPlotArea(plotArea)
	}

	// Create the appropriate chart type element based on chartType
	// Default axis IDs
	const (
		categoryAxisID uint32 = 100000000
		valueAxisID    uint32 = 100000001
	)

	switch chartType {
	case ChartTypeBar:
		bc := plotArea.AddBarChart(
			drawingml.BarDirectionCol,
			drawingml.BarGroupingClustered,
		)
		bc.AddAxisID(categoryAxisID)
		bc.AddAxisID(valueAxisID)
	case ChartTypeBar3D:
		bc := plotArea.AddBar3DChart(
			drawingml.BarDirectionCol,
			drawingml.BarGroupingClustered,
		)
		bc.AddAxisID(categoryAxisID)
		bc.AddAxisID(valueAxisID)
	case ChartTypeLine:
		lc := plotArea.AddLineChart(
			drawingml.GroupingStandard,
		)
		lc.AddAxisID(categoryAxisID)
		lc.AddAxisID(valueAxisID)
	case ChartTypeLine3D:
		lc := plotArea.AddLine3DChart(
			drawingml.GroupingStandard,
		)
		lc.AddAxisID(categoryAxisID)
		lc.AddAxisID(valueAxisID)
	case ChartTypePie:
		plotArea.AddPieChart()
	case ChartTypePie3D:
		plotArea.AddPie3DChart()
	case ChartTypeDoughnut:
		plotArea.AddDoughnutChart()
	case ChartTypeArea:
		ac := plotArea.AddAreaChart(
			drawingml.GroupingStandard,
		)
		ac.AddAxisID(categoryAxisID)
		ac.AddAxisID(valueAxisID)
	case ChartTypeArea3D:
		ac := plotArea.AddArea3DChart(
			drawingml.GroupingStandard,
		)
		ac.AddAxisID(categoryAxisID)
		ac.AddAxisID(valueAxisID)
	case ChartTypeScatter:
		sc := plotArea.AddScatterChart(
			drawingml.ScatterStyleMarker,
		)
		sc.AddAxisID(categoryAxisID)
		sc.AddAxisID(valueAxisID)
	case ChartTypeBubble:
		bc := plotArea.AddBubbleChart()
		bc.AddAxisID(categoryAxisID)
		bc.AddAxisID(valueAxisID)
	case ChartTypeStock:
		sc := plotArea.AddStockChart()
		sc.AddAxisID(categoryAxisID)
		sc.AddAxisID(valueAxisID)
	case ChartTypeSurface:
		sc := plotArea.AddSurfaceChart()
		sc.AddAxisID(categoryAxisID)
		sc.AddAxisID(valueAxisID)
	case ChartTypeSurface3D:
		sc := plotArea.AddSurface3DChart()
		sc.AddAxisID(categoryAxisID)
		sc.AddAxisID(valueAxisID)
	case ChartTypeRadar:
		rc := plotArea.AddRadarChart(
			drawingml.RadarStyleStandard,
		)
		rc.AddAxisID(categoryAxisID)
		rc.AddAxisID(valueAxisID)
	default:
		// Default to bar chart
		bc := plotArea.AddBarChart(
			drawingml.BarDirectionCol,
			drawingml.BarGroupingClustered,
		)
		bc.AddAxisID(categoryAxisID)
		bc.AddAxisID(valueAxisID)
	}

	// Add axes for chart types that need them (not pie/doughnut)
	if chartType != ChartTypePie &&
		chartType != ChartTypePie3D &&
		chartType != ChartTypeDoughnut {
		plotArea.AddCategoryAxis(
			categoryAxisID,
			valueAxisID,
		)
		plotArea.AddValueAxis(
			valueAxisID,
			categoryAxisID,
		)
	}

	return c, nil
}

// Type returns the chart type.
func (c *Chart) Type() ChartType {
	return c.chartType
}

// Title returns the chart title.
func (c *Chart) Title() string {
	return c.title
}

// SetTitle sets the chart title.
func (c *Chart) SetTitle(title string) {
	c.title = title

	chartSpace := c.chartPart.ChartSpace()
	if chartSpace == nil {
		return
	}

	chart := chartSpace.Chart()
	if chart == nil {
		return
	}

	if title == "" {
		// Remove title and set autoTitleDeleted
		chart.SetTitle(nil)
		chart.SetAutoTitleDeleted(true)
	} else {
		// Create and set title
		titleElem := drawingml.NewTitleWithText(title)
		chart.SetTitle(titleElem)
		chart.SetAutoTitleDeleted(false)
	}
}

// Series returns the chart series at the given index.
func (c *Chart) Series(index int) *ChartSeries {
	if index < 0 || index >= len(c.series) {
		return nil
	}

	return c.series[index]
}

// SeriesCount returns the number of series in the chart.
func (c *Chart) SeriesCount() int {
	return len(c.series)
}

// AddSeries adds a new series to the chart.
func (c *Chart) AddSeries(
	name, valuesRange string,
) *ChartSeries {
	series := newChartSeries(c, name, valuesRange)
	c.series = append(c.series, series)

	// Update the underlying chart element
	series.updateUnderlyingElement()

	return series
}

// CategoryAxis returns the category axis.
func (c *Chart) CategoryAxis() *ChartAxis {
	return c.categoryAxis
}

// ValueAxis returns the value axis.
func (c *Chart) ValueAxis() *ChartAxis {
	return c.valueAxis
}
