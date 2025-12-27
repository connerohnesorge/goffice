// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements the high-level Chart API.
package spreadsheet

import (
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

	// TODO: Create the chart elements based on chartType and dataRange
	// This would involve creating the appropriate chart type element
	// (BarChart, LineChart, etc.) and populating it with series data

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
	// TODO: Update the underlying chart element title
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
	// TODO: Update the underlying chart element

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

// SetLegendPosition sets the legend position.
func (*Chart) SetLegendPosition(
	position LegendPosition,
) {
	// TODO: Update the underlying chart element
	_ = position
}

// SetShowLegend sets whether to show the legend.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (*Chart) SetShowLegend(show bool) {
	// TODO: Update the underlying chart element
	_ = show
}

// LegendPosition represents the position of the chart legend.
type LegendPosition string

const (
	// LegendPositionBottom places the legend at the bottom.
	LegendPositionBottom LegendPosition = "b"
	// LegendPositionTop places the legend at the top.
	LegendPositionTop LegendPosition = "t"
	// LegendPositionLeft places the legend at the left.
	LegendPositionLeft LegendPosition = "l"
	// LegendPositionRight places the legend at the right.
	LegendPositionRight LegendPosition = "r"
	// LegendPositionTopRight places the legend at the top right.
	LegendPositionTopRight LegendPosition = "tr"
)

// ChartSeries represents a data series in a chart.
type ChartSeries struct {
	chart       *Chart
	name        string
	valuesRange string
	catRange    string
}

// newChartSeries creates a new ChartSeries.
func newChartSeries(
	chart *Chart,
	name, valuesRange string,
) *ChartSeries {
	return &ChartSeries{
		chart:       chart,
		name:        name,
		valuesRange: valuesRange,
	}
}

// Name returns the series name.
func (s *ChartSeries) Name() string {
	return s.name
}

// SetName sets the series name.
func (s *ChartSeries) SetName(name string) {
	s.name = name
	// TODO: Update the underlying series element
}

// ValuesRange returns the values range reference.
func (s *ChartSeries) ValuesRange() string {
	return s.valuesRange
}

// SetValuesRange sets the values range reference.
func (s *ChartSeries) SetValuesRange(
	rangeRef string,
) {
	s.valuesRange = rangeRef
	// TODO: Update the underlying series element
}

// CategoriesRange returns the categories range reference.
func (s *ChartSeries) CategoriesRange() string {
	return s.catRange
}

// SetCategoriesRange sets the categories range reference.
func (s *ChartSeries) SetCategoriesRange(
	rangeRef string,
) {
	s.catRange = rangeRef
	// TODO: Update the underlying series element
}

// ChartAxis represents an axis in a chart.
type ChartAxis struct {
	chart    *Chart
	axisType string
	title    string
	minVal   *float64
	maxVal   *float64
}

// newChartAxis creates a new ChartAxis.
func newChartAxis(
	chart *Chart,
	axisType string,
) *ChartAxis {
	return &ChartAxis{
		chart:    chart,
		axisType: axisType,
	}
}

// Title returns the axis title.
func (a *ChartAxis) Title() string {
	return a.title
}

// SetTitle sets the axis title.
func (a *ChartAxis) SetTitle(title string) {
	a.title = title
	// TODO: Update the underlying axis element
}

// SetMinimum sets the minimum value for the axis.
func (a *ChartAxis) SetMinimum(val float64) {
	a.minVal = &val
	// TODO: Update the underlying axis element
}

// SetMaximum sets the maximum value for the axis.
func (a *ChartAxis) SetMaximum(val float64) {
	a.maxVal = &val
	// TODO: Update the underlying axis element
}

// SetMajorGridlines sets whether to show major gridlines.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (a *ChartAxis) SetMajorGridlines(show bool) {
	// TODO: Update the underlying axis element
	_ = show
}

// SetMinorGridlines sets whether to show minor gridlines.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (a *ChartAxis) SetMinorGridlines(show bool) {
	// TODO: Update the underlying axis element
	_ = show
}

// SetNumberFormat sets the number format for axis labels.
func (*ChartAxis) SetNumberFormat(
	format string,
) {
	// TODO: Update the underlying axis element
	_ = format
}
