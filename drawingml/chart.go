//revive:disable:file-length-limit chart types need to be together

// Package drawingml implements chart types and effects.
//
// This file implements chart types for the DrawingML Chart namespace (c:).
package drawingml

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Element name constants for charts.
const (
	// attrTrue is the string "true" for boolean attributes.
	attrTrue = "true"
	// elemChart is the "chart" element name.
	elemChart = "chart"
	// elemSpPr is the "spPr" shape properties element name.
	elemSpPr = "spPr"
)

// =============================================================================
// Chart Space - Root element for chart parts
// =============================================================================

// ChartSpace represents the c:chartSpace root element for chart parts.
// It is the top-level container for all chart content.
type ChartSpace struct {
	*openxml.PartRootElementBase
}

// NewChartSpace creates a new chart space element.
func NewChartSpace() *ChartSpace {
	elem := openxml.NewPartRootElement(
		NamespaceChart,
		"chartSpace",
		PrefixChart,
	)
	cs := &ChartSpace{PartRootElementBase: elem}

	// Add the main drawing namespace declaration
	elem.SetAttribute(openxml.NewAttribute(
		"http://www.w3.org/2000/xmlns/",
		PrefixMain,
		"xmlns",
		NamespaceMain,
	))

	// Add default chart
	cs.AppendChild(NewChart())

	return cs
}

// Date1904 returns whether the chart uses 1904 date system.
func (cs *ChartSpace) Date1904() bool {
	elem := cs.GetElement(
		"date1904",
		NamespaceChart,
	)
	if elem == nil {
		return false
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return true // default is true if element exists
	}

	return attr.Value() == attrOne ||
		attr.Value() == attrTrue
}

// SetDate1904 sets whether the chart uses 1904 date system.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cs *ChartSpace) SetDate1904(val bool) {
	cs.removeElement("date1904")
	if !val {
		return
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"date1904",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrOne,
		),
	)
	cs.prependElement(elem)
}

// RoundedCorners returns whether the chart has rounded corners.
func (cs *ChartSpace) RoundedCorners() bool {
	elem := cs.GetElement(
		"roundedCorners",
		NamespaceChart,
	)
	if elem == nil {
		return false
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return true
	}

	return attr.Value() == attrOne ||
		attr.Value() == attrTrue
}

// SetRoundedCorners sets whether the chart has rounded corners.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cs *ChartSpace) SetRoundedCorners(
	val bool,
) {
	cs.removeElement("roundedCorners")
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"roundedCorners",
		PrefixChart,
	)
	if val {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				attrVal,
				attrOne,
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, attrZero))
	}
	cs.insertBeforeChart(elem)
}

// Chart returns the chart element.
func (cs *ChartSpace) Chart() *Chart {
	elem := cs.GetElement(
		elemChart,
		NamespaceChart,
	)
	if elem == nil {
		return nil
	}
	if c, ok := elem.(*Chart); ok {
		return c
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Chart{CompositeElementBase: comp}
	}

	return nil
}

// SetChart sets or replaces the chart element.
func (cs *ChartSpace) SetChart(chart *Chart) {
	cs.removeElement(elemChart)
	if chart != nil {
		cs.AppendChild(chart)
	}
}

// ShapeProperties returns the shape properties for the chart space.
func (cs *ChartSpace) ShapeProperties() *ChartShapeProperties {
	elem := cs.GetElement(
		elemSpPr,
		NamespaceChart,
	)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*ChartShapeProperties); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ChartShapeProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetShapeProperties sets the shape properties for the chart space.
func (cs *ChartSpace) SetShapeProperties(
	props *ChartShapeProperties,
) {
	if existing := cs.GetElement(elemSpPr, NamespaceChart); existing != nil {
		cs.RemoveChild(existing)
	}
	if props != nil {
		cs.AppendChild(props)
	}
}

// removeElement removes an element by local name.
func (cs *ChartSpace) removeElement(
	localName string,
) {
	if elem := cs.GetElement(localName, NamespaceChart); elem != nil {
		cs.RemoveChild(elem)
	}
}

// prependElement adds an element at the beginning.
func (cs *ChartSpace) prependElement(
	elem openxml.Element,
) {
	cs.PrependChild(elem)
}

// insertBeforeChart inserts an element before the chart element.
func (cs *ChartSpace) insertBeforeChart(
	elem openxml.Element,
) {
	chart := cs.GetElement(
		elemChart,
		NamespaceChart,
	)
	if chart != nil {
		cs.InsertBefore(elem, chart)
	} else {
		cs.AppendChild(elem)
	}
}

// Clone creates a deep copy of this ChartSpace.
func (cs *ChartSpace) Clone() openxml.Element {
	cloned := cs.PartRootElementBase.Clone()
	clone, ok := cloned.(*openxml.PartRootElementBase)
	if !ok {
		return nil
	}

	return &ChartSpace{PartRootElementBase: clone}
}

// =============================================================================
// Chart - Main chart element
// =============================================================================

// Chart represents the c:chart element containing chart data and formatting.
type Chart struct {
	*openxml.CompositeElementBase
}

// NewChart creates a new chart element.
func NewChart() *Chart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"chart",
		PrefixChart,
	)
	c := &Chart{CompositeElementBase: elem}

	// Add default plot area
	c.AppendChild(NewPlotArea())

	return c
}

// Title returns the chart title element.
func (c *Chart) Title() *Title {
	elem := c.GetElement("title", NamespaceChart)
	if elem == nil {
		return nil
	}
	if t, ok := elem.(*Title); ok {
		return t
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Title{CompositeElementBase: comp}
	}

	return nil
}

// SetTitle sets the chart title.
func (c *Chart) SetTitle(title *Title) {
	if existing := c.GetElement("title", NamespaceChart); existing != nil {
		c.RemoveChild(existing)
	}
	if title != nil {
		c.PrependChild(title)
	}
}

// AutoTitleDeleted returns whether the auto title is deleted.
func (c *Chart) AutoTitleDeleted() bool {
	elem := c.GetElement(
		"autoTitleDeleted",
		NamespaceChart,
	)
	if elem == nil {
		return false
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return true
	}

	return attr.Value() == attrOne ||
		attr.Value() == attrTrue
}

// SetAutoTitleDeleted sets whether the auto title is deleted.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Chart) SetAutoTitleDeleted(val bool) {
	existing := c.GetElement(
		"autoTitleDeleted",
		NamespaceChart,
	)
	if existing != nil {
		c.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"autoTitleDeleted",
		PrefixChart,
	)
	if val {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				attrVal,
				attrOne,
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, attrZero))
	}
	c.insertAfterTitle(elem)
}

// PlotArea returns the plot area element.
func (c *Chart) PlotArea() *PlotArea {
	elem := c.GetElement(
		"plotArea",
		NamespaceChart,
	)
	if elem == nil {
		return nil
	}
	if pa, ok := elem.(*PlotArea); ok {
		return pa
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PlotArea{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetPlotArea sets the plot area.
func (c *Chart) SetPlotArea(plotArea *PlotArea) {
	if existing := c.GetElement("plotArea", NamespaceChart); existing != nil {
		c.RemoveChild(existing)
	}
	if plotArea != nil {
		c.AppendChild(plotArea)
	}
}

// Legend returns the legend element.
func (c *Chart) Legend() *Legend {
	elem := c.GetElement("legend", NamespaceChart)
	if elem == nil {
		return nil
	}
	if l, ok := elem.(*Legend); ok {
		return l
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Legend{CompositeElementBase: comp}
	}

	return nil
}

// SetLegend sets the chart legend.
func (c *Chart) SetLegend(legend *Legend) {
	if existing := c.GetElement("legend", NamespaceChart); existing != nil {
		c.RemoveChild(existing)
	}
	if legend != nil {
		c.AppendChild(legend)
	}
}

// PlotVisibleOnly returns whether only visible cells are plotted.
func (c *Chart) PlotVisibleOnly() bool {
	elem := c.GetElement(
		"plotVisOnly",
		NamespaceChart,
	)
	if elem == nil {
		return true // default
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return true
	}

	return attr.Value() == attrOne ||
		attr.Value() == attrTrue
}

// SetPlotVisibleOnly sets whether only visible cells are plotted.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Chart) SetPlotVisibleOnly(val bool) {
	existing := c.GetElement(
		"plotVisOnly",
		NamespaceChart,
	)
	if existing != nil {
		c.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"plotVisOnly",
		PrefixChart,
	)
	if val {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				attrVal,
				attrOne,
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, attrZero))
	}
	c.AppendChild(elem)
}

// DisplayBlanksAs returns how blank cells are displayed.
func (c *Chart) DisplayBlanksAs() DisplayBlanksAsValue {
	elem := c.GetElement(
		"dispBlanksAs",
		NamespaceChart,
	)
	if elem == nil {
		return DisplayBlanksAsGap
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return DisplayBlanksAsGap
	}

	return DisplayBlanksAsValue(attr.Value())
}

// SetDisplayBlanksAs sets how blank cells are displayed.
func (c *Chart) SetDisplayBlanksAs(
	val DisplayBlanksAsValue,
) {
	existing := c.GetElement(
		"dispBlanksAs",
		NamespaceChart,
	)
	if existing != nil {
		c.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"dispBlanksAs",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(val),
		),
	)
	c.AppendChild(elem)
}

// insertAfterTitle inserts an element after the title element.
func (c *Chart) insertAfterTitle(
	elem openxml.Element,
) {
	title := c.GetElement("title", NamespaceChart)
	if title != nil {
		c.InsertAfter(elem, title)
	} else {
		c.PrependChild(elem)
	}
}

// Clone creates a deep copy of this Chart.
func (c *Chart) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()
	clone, ok := cloned.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Chart{CompositeElementBase: clone}
}

// =============================================================================
// Plot Area
// =============================================================================

// PlotArea represents the c:plotArea element containing chart data.
type PlotArea struct {
	*openxml.CompositeElementBase
}

// NewPlotArea creates a new plot area element.
func NewPlotArea() *PlotArea {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"plotArea",
		PrefixChart,
	)

	return &PlotArea{CompositeElementBase: elem}
}

// Layout returns the layout element.
func (pa *PlotArea) Layout() *Layout {
	elem := pa.GetElement(
		"layout",
		NamespaceChart,
	)
	if elem == nil {
		return nil
	}
	if l, ok := elem.(*Layout); ok {
		return l
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Layout{CompositeElementBase: comp}
	}

	return nil
}

// SetLayout sets the layout element.
func (pa *PlotArea) SetLayout(layout *Layout) {
	if existing := pa.GetElement("layout", NamespaceChart); existing != nil {
		pa.RemoveChild(existing)
	}
	if layout != nil {
		pa.PrependChild(layout)
	}
}

// AddBarChart adds a bar chart to the plot area.
func (pa *PlotArea) AddBarChart(
	direction BarDirectionValue,
	grouping BarGroupingValue,
) *BarChart {
	bc := NewBarChart(direction, grouping)
	pa.AppendChild(bc)

	return bc
}

// AddLineChart adds a line chart to the plot area.
func (pa *PlotArea) AddLineChart(
	grouping GroupingValue,
) *LineChart {
	lc := NewLineChart(grouping)
	pa.AppendChild(lc)

	return lc
}

// AddPieChart adds a pie chart to the plot area.
func (pa *PlotArea) AddPieChart() *PieChart {
	pc := NewPieChart()
	pa.AppendChild(pc)

	return pc
}

// AddAreaChart adds an area chart to the plot area.
func (pa *PlotArea) AddAreaChart(
	grouping GroupingValue,
) *AreaChart {
	ac := NewAreaChart(grouping)
	pa.AppendChild(ac)

	return ac
}

// AddScatterChart adds a scatter chart to the plot area.
func (pa *PlotArea) AddScatterChart(
	style ScatterStyleValue,
) *ScatterChart {
	sc := NewScatterChart(style)
	pa.AppendChild(sc)

	return sc
}

// AddDoughnutChart adds a doughnut chart to the plot area.
func (pa *PlotArea) AddDoughnutChart() *DoughnutChart {
	dc := NewDoughnutChart()
	pa.AppendChild(dc)

	return dc
}

// AddRadarChart adds a radar chart to the plot area.
func (pa *PlotArea) AddRadarChart(
	style RadarStyleValue,
) *RadarChart {
	rc := NewRadarChart(style)
	pa.AppendChild(rc)

	return rc
}

// AddBubbleChart adds a bubble chart to the plot area.
func (pa *PlotArea) AddBubbleChart() *BubbleChart {
	bc := NewBubbleChart()
	pa.AppendChild(bc)

	return bc
}

// AddStockChart adds a stock chart to the plot area.
func (pa *PlotArea) AddStockChart() *StockChart {
	sc := NewStockChart()
	pa.AppendChild(sc)

	return sc
}

// AddSurfaceChart adds a surface chart to the plot area.
func (pa *PlotArea) AddSurfaceChart() *SurfaceChart {
	sc := NewSurfaceChart()
	pa.AppendChild(sc)

	return sc
}

// AddOfPieChart adds an of-pie chart (pie-of-pie or bar-of-pie) to the
// plot area.
func (pa *PlotArea) AddOfPieChart(
	pieType OfPieTypeValue,
) *OfPieChart {
	opc := NewOfPieChart(pieType)
	pa.AppendChild(opc)

	return opc
}

// AddBar3DChart adds a 3D bar chart to the plot area.
func (pa *PlotArea) AddBar3DChart(
	direction BarDirectionValue,
	grouping BarGroupingValue,
) *Bar3DChart {
	bc := NewBar3DChart(direction, grouping)
	pa.AppendChild(bc)

	return bc
}

// AddLine3DChart adds a 3D line chart to the plot area.
func (pa *PlotArea) AddLine3DChart(
	grouping GroupingValue,
) *Line3DChart {
	lc := NewLine3DChart(grouping)
	pa.AppendChild(lc)

	return lc
}

// AddPie3DChart adds a 3D pie chart to the plot area.
func (pa *PlotArea) AddPie3DChart() *Pie3DChart {
	pc := NewPie3DChart()
	pa.AppendChild(pc)

	return pc
}

// AddArea3DChart adds a 3D area chart to the plot area.
func (pa *PlotArea) AddArea3DChart(
	grouping GroupingValue,
) *Area3DChart {
	ac := NewArea3DChart(grouping)
	pa.AppendChild(ac)

	return ac
}

// AddSurface3DChart adds a 3D surface chart to the plot area.
func (pa *PlotArea) AddSurface3DChart() *Surface3DChart {
	sc := NewSurface3DChart()
	pa.AppendChild(sc)

	return sc
}

// AddCategoryAxis adds a category axis to the plot area.
func (pa *PlotArea) AddCategoryAxis(
	axisID, crossAxisID uint32,
) *CategoryAxis {
	ax := NewCategoryAxis(axisID, crossAxisID)
	pa.AppendChild(ax)

	return ax
}

// AddValueAxis adds a value axis to the plot area.
func (pa *PlotArea) AddValueAxis(
	axisID, crossAxisID uint32,
) *ValueAxis {
	ax := NewValueAxis(axisID, crossAxisID)
	pa.AppendChild(ax)

	return ax
}

// AddDateAxis adds a date axis to the plot area.
func (pa *PlotArea) AddDateAxis(
	axisID, crossAxisID uint32,
) *DateAxis {
	ax := NewDateAxis(axisID, crossAxisID)
	pa.AppendChild(ax)

	return ax
}

// AddSeriesAxis adds a series axis to the plot area.
func (pa *PlotArea) AddSeriesAxis(
	axisID, crossAxisID uint32,
) *SeriesAxis {
	ax := NewSeriesAxis(axisID, crossAxisID)
	pa.AppendChild(ax)

	return ax
}

// ShapeProperties returns the shape properties for the plot area.
func (pa *PlotArea) ShapeProperties() *ChartShapeProperties {
	elem := pa.GetElement(
		elemSpPr,
		NamespaceChart,
	)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*ChartShapeProperties); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ChartShapeProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetShapeProperties sets the shape properties for the plot area.
func (pa *PlotArea) SetShapeProperties(
	props *ChartShapeProperties,
) {
	if existing := pa.GetElement(elemSpPr, NamespaceChart); existing != nil {
		pa.RemoveChild(existing)
	}
	if props != nil {
		pa.AppendChild(props)
	}
}

// SetDataTable sets the data table for the plot area.
func (pa *PlotArea) SetDataTable(dt *DataTable) {
	if existing := pa.GetElement("dTable", NamespaceChart); existing != nil {
		pa.RemoveChild(existing)
	}
	if dt != nil {
		pa.AppendChild(dt)
	}
}

// Clone creates a deep copy of this PlotArea.
func (pa *PlotArea) Clone() openxml.Element {
	cloned := pa.CompositeElementBase.Clone()
	clone, ok := cloned.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &PlotArea{CompositeElementBase: clone}
}
