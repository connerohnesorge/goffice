// Package drawingml provides shared DrawingML types.
// This file implements extended chart types (Office 2016+).
package drawingml

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Constants for extended charts.
const (
	// NamespaceChartEx is the namespace for extended charts (cx:).
	NamespaceChartEx = "http://schemas.microsoft.com/office/drawing/2014/chartex"
	// PrefixChartEx is the prefix for extended charts.
	PrefixChartEx = "cx"
)

// ChartSpaceEx represents the cx:chartSpace root element for extended chart parts.
type ChartSpaceEx struct {
	*openxml.PartRootElementBase
}

// NewChartSpaceEx creates a new extended chart space element.
func NewChartSpaceEx() *ChartSpaceEx {
	elem := openxml.NewPartRootElement(
		NamespaceChartEx,
		"chartSpace",
		PrefixChartEx,
	)
	cs := &ChartSpaceEx{PartRootElementBase: elem}

	// Add xmlns declaration for chart extended namespace
	elem.SetAttribute(openxml.NewAttribute(
		"http://www.w3.org/2000/xmlns/",
		PrefixChartEx,
		"xmlns",
		NamespaceChartEx,
	))

	// Add chart data container
	cs.AppendChild(NewChartData())
	// Add chart container
	cs.AppendChild(NewChartEx())

	return cs
}

// ChartData represents the cx:chartData element.
type ChartData struct {
	*openxml.CompositeElementBase
}

// NewChartData creates a new chart data element.
func NewChartData() *ChartData {
	elem := openxml.NewCompositeElement(
		NamespaceChartEx,
		"chartData",
		PrefixChartEx,
	)
	return &ChartData{CompositeElementBase: elem}
}

// ChartEx represents the cx:chart element.
type ChartEx struct {
	*openxml.CompositeElementBase
}

// NewChartEx creates a new extended chart element.
func NewChartEx() *ChartEx {
	elem := openxml.NewCompositeElement(
		NamespaceChartEx,
		"chart",
		PrefixChartEx,
	)
	return &ChartEx{CompositeElementBase: elem}
}

// PlotAreaEx represents the cx:plotArea element.
type PlotAreaEx struct {
	*openxml.CompositeElementBase
}

// NewPlotAreaEx creates a new extended plot area.
func NewPlotAreaEx() *PlotAreaEx {
	elem := openxml.NewCompositeElement(
		NamespaceChartEx,
		"plotArea",
		PrefixChartEx,
	)
	return &PlotAreaEx{CompositeElementBase: elem}
}

// SetPlotArea sets the plot area for the chart.
func (c *ChartEx) SetPlotArea(pa *PlotAreaEx) {
	if existing := c.GetElement("plotArea", NamespaceChartEx); existing != nil {
		c.RemoveChild(existing)
	}
	if pa != nil {
		c.AppendChild(pa)
	}
}

// PlotAreaRegion represents a region in the plot area (cx:plotAreaRegion).
type PlotAreaRegion struct {
	*openxml.CompositeElementBase
}

// NewPlotAreaRegion creates a new plot area region.
func NewPlotAreaRegion() *PlotAreaRegion {
	elem := openxml.NewCompositeElement(
		NamespaceChartEx,
		"plotAreaRegion",
		PrefixChartEx,
	)
	return &PlotAreaRegion{CompositeElementBase: elem}
}

// AddPlotAreaRegion adds a region to the plot area.
func (pa *PlotAreaEx) AddPlotAreaRegion(region *PlotAreaRegion) {
	pa.AppendChild(region)
}

// SeriesEx represents a series in an extended chart (cx:series).
type SeriesEx struct {
	*openxml.CompositeElementBase
}

// NewSeriesEx creates a new extended series.
func NewSeriesEx() *SeriesEx {
	elem := openxml.NewCompositeElement(
		NamespaceChartEx,
		"series",
		PrefixChartEx,
	)
	return &SeriesEx{CompositeElementBase: elem}
}

// AddSeries adds a series to the plot area region.
func (par *PlotAreaRegion) AddSeries(series *SeriesEx) {
	par.AppendChild(series)
}

// SetLayout sets the layout ID for the series (e.g., "waterfall", "sunburst").
func (s *SeriesEx) SetLayoutID(layoutID string) {
	s.SetAttribute(openxml.NewSimpleAttribute("layoutId", layoutID))
}

// SeriesLayout constants for extended charts.
const (
	SeriesLayoutWaterfall  = "waterfall"
	SeriesLayoutSunburst   = "sunburst"
	SeriesLayoutTreemap    = "treemap"
	SeriesLayoutBoxWhisker = "boxWhisker"
)
