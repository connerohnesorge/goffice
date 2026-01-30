//revive:disable:file-length-limit chart type definitions need to be together
//revive:disable:max-public-structs chart types have many public structs

// Package drawingml provides shared DrawingML types for shapes, images, charts, diagrams,
// and effects used across Office Open XML documents.
//
// This file implements chart type structs (BarChart, LineChart, etc.).
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Chart type element name constants.
const (
	// elemGrouping is the "grouping" element name.
	elemGrouping = "grouping"
	// elemVaryColors is the "varyColors" element name.
	elemVaryColors = "varyColors"
	// elemAxID is the "axId" element name.
	elemAxID = "axId"
	// chartTypesBase10 is the base used for integer to string conversion.
	chartTypesBase10 = 10
)

// BarChart represents a c:barChart element.
type BarChart struct {
	*openxml.CompositeElementBase
}

// NewBarChart creates a new bar chart with the specified direction
// and grouping.
func NewBarChart(
	direction BarDirectionValue,
	grouping BarGroupingValue,
) *BarChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"barChart",
		PrefixChart,
	)
	bc := &BarChart{CompositeElementBase: elem}

	// Add direction
	dirElem := openxml.NewLeafElement(
		NamespaceChart,
		"barDir",
		PrefixChart,
	)
	dirElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(direction),
		),
	)
	bc.AppendChild(dirElem)

	// Add grouping
	groupElem := openxml.NewLeafElement(
		NamespaceChart,
		elemGrouping,
		PrefixChart,
	)
	groupElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(grouping),
		),
	)
	bc.AppendChild(groupElem)

	// Add varyColors
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrZero,
		),
	)
	bc.AppendChild(varyColors)

	return bc
}

// Direction returns the bar direction.
func (bc *BarChart) Direction() BarDirectionValue {
	elem := bc.GetElement(
		"barDir",
		NamespaceChart,
	)
	if elem == nil {
		return BarDirectionCol
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return BarDirectionCol
	}

	return BarDirectionValue(attr.Value())
}

// Grouping returns the bar grouping.
func (bc *BarChart) Grouping() BarGroupingValue {
	elem := bc.GetElement(
		elemGrouping,
		NamespaceChart,
	)
	if elem == nil {
		return BarGroupingClustered
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return BarGroupingClustered
	}

	return BarGroupingValue(attr.Value())
}

// AddSeries adds a series to the bar chart.
func (bc *BarChart) AddSeries(
	index, order uint32,
) *BarChartSeries {
	ser := NewBarChartSeries(index, order)
	bc.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (bc *BarChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	bc.AppendChild(axID)
}

// Clone creates a deep copy of this BarChart.
func (bc *BarChart) Clone() openxml.Element {
	clonedBase := bc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &BarChart{CompositeElementBase: cloned}
}

// LineChart represents a c:lineChart element.
type LineChart struct {
	*openxml.CompositeElementBase
}

// NewLineChart creates a new line chart with the specified grouping.
func NewLineChart(
	grouping GroupingValue,
) *LineChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"lineChart",
		PrefixChart,
	)
	lc := &LineChart{CompositeElementBase: elem}

	// Add grouping
	groupElem := openxml.NewLeafElement(
		NamespaceChart,
		elemGrouping,
		PrefixChart,
	)
	groupElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(grouping),
		),
	)
	lc.AppendChild(groupElem)

	// Add varyColors
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrZero,
		),
	)
	lc.AppendChild(varyColors)

	return lc
}

// Grouping returns the chart grouping.
func (lc *LineChart) Grouping() GroupingValue {
	elem := lc.GetElement(
		elemGrouping,
		NamespaceChart,
	)
	if elem == nil {
		return GroupingStandard
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return GroupingStandard
	}

	return GroupingValue(attr.Value())
}

// AddSeries adds a series to the line chart.
func (lc *LineChart) AddSeries(
	index, order uint32,
) *LineChartSeries {
	ser := NewLineChartSeries(index, order)
	lc.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (lc *LineChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	lc.AppendChild(axID)
}

// Clone creates a deep copy of this LineChart.
func (lc *LineChart) Clone() openxml.Element {
	clonedBase := lc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &LineChart{
		CompositeElementBase: cloned,
	}
}

// PieChart represents a c:pieChart element.
type PieChart struct {
	*openxml.CompositeElementBase
}

// NewPieChart creates a new pie chart.
func NewPieChart() *PieChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"pieChart",
		PrefixChart,
	)
	pc := &PieChart{CompositeElementBase: elem}

	// Add varyColors (typically true for pie charts)
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrOne,
		),
	)
	pc.AppendChild(varyColors)

	return pc
}

// AddSeries adds a series to the pie chart.
func (pc *PieChart) AddSeries(
	index, order uint32,
) *PieChartSeries {
	ser := NewPieChartSeries(index, order)
	pc.AppendChild(ser)

	return ser
}

// Clone creates a deep copy of this PieChart.
func (pc *PieChart) Clone() openxml.Element {
	clonedBase := pc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &PieChart{CompositeElementBase: cloned}
}

// AreaChart represents a c:areaChart element.
type AreaChart struct {
	*openxml.CompositeElementBase
}

// NewAreaChart creates a new area chart with the specified grouping.
func NewAreaChart(
	grouping GroupingValue,
) *AreaChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"areaChart",
		PrefixChart,
	)
	ac := &AreaChart{CompositeElementBase: elem}

	// Add grouping
	groupElem := openxml.NewLeafElement(
		NamespaceChart,
		elemGrouping,
		PrefixChart,
	)
	groupElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(grouping),
		),
	)
	ac.AppendChild(groupElem)

	// Add varyColors
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrZero,
		),
	)
	ac.AppendChild(varyColors)

	return ac
}

// AddSeries adds a series to the area chart.
func (ac *AreaChart) AddSeries(
	index, order uint32,
) *AreaChartSeries {
	ser := NewAreaChartSeries(index, order)
	ac.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (ac *AreaChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	ac.AppendChild(axID)
}

// Clone creates a deep copy of this AreaChart.
func (ac *AreaChart) Clone() openxml.Element {
	clonedBase := ac.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &AreaChart{
		CompositeElementBase: cloned,
	}
}

// ScatterChart represents a c:scatterChart element.
type ScatterChart struct {
	*openxml.CompositeElementBase
}

// NewScatterChart creates a new scatter chart with the specified style.
func NewScatterChart(
	style ScatterStyleValue,
) *ScatterChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"scatterChart",
		PrefixChart,
	)
	sc := &ScatterChart{
		CompositeElementBase: elem,
	}

	// Add scatter style
	styleElem := openxml.NewLeafElement(
		NamespaceChart,
		"scatterStyle",
		PrefixChart,
	)
	styleElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(style),
		),
	)
	sc.AppendChild(styleElem)

	// Add varyColors
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrZero,
		),
	)
	sc.AppendChild(varyColors)

	return sc
}

// AddSeries adds a series to the scatter chart.
func (sc *ScatterChart) AddSeries(
	index, order uint32,
) *ScatterChartSeries {
	ser := NewScatterChartSeries(index, order)
	sc.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (sc *ScatterChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	sc.AppendChild(axID)
}

// Clone creates a deep copy of this ScatterChart.
func (sc *ScatterChart) Clone() openxml.Element {
	clonedBase := sc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &ScatterChart{
		CompositeElementBase: cloned,
	}
}

// DoughnutChart represents a c:doughnutChart element.
type DoughnutChart struct {
	*openxml.CompositeElementBase
}

// NewDoughnutChart creates a new doughnut chart.
func NewDoughnutChart() *DoughnutChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"doughnutChart",
		PrefixChart,
	)
	dc := &DoughnutChart{
		CompositeElementBase: elem,
	}

	// Add varyColors (typically true for doughnut charts)
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrOne,
		),
	)
	dc.AppendChild(varyColors)

	return dc
}

// AddSeries adds a series to the doughnut chart.
func (dc *DoughnutChart) AddSeries(
	index, order uint32,
) *PieChartSeries {
	ser := NewPieChartSeries(index, order)
	dc.AppendChild(ser)

	return ser
}

// SetHoleSize sets the hole size (0-90).
func (dc *DoughnutChart) SetHoleSize(
	size uint32,
) {
	existing := dc.GetElement(
		"holeSize",
		NamespaceChart,
	)
	if existing != nil {
		dc.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"holeSize",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(size),
				chartTypesBase10,
			),
		),
	)
	dc.AppendChild(elem)
}

// Clone creates a deep copy of this DoughnutChart.
func (dc *DoughnutChart) Clone() openxml.Element {
	clonedBase := dc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &DoughnutChart{
		CompositeElementBase: cloned,
	}
}

// RadarChart represents a c:radarChart element.
type RadarChart struct {
	*openxml.CompositeElementBase
}

// NewRadarChart creates a new radar chart with the specified style.
func NewRadarChart(
	style RadarStyleValue,
) *RadarChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"radarChart",
		PrefixChart,
	)
	rc := &RadarChart{CompositeElementBase: elem}

	// Add radar style
	styleElem := openxml.NewLeafElement(
		NamespaceChart,
		"radarStyle",
		PrefixChart,
	)
	styleElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(style),
		),
	)
	rc.AppendChild(styleElem)

	// Add varyColors
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrZero,
		),
	)
	rc.AppendChild(varyColors)

	return rc
}

// AddSeries adds a series to the radar chart.
func (rc *RadarChart) AddSeries(
	index, order uint32,
) *LineChartSeries {
	ser := NewLineChartSeries(index, order)
	rc.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (rc *RadarChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	rc.AppendChild(axID)
}

// Clone creates a deep copy of this RadarChart.
func (rc *RadarChart) Clone() openxml.Element {
	clonedBase := rc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &RadarChart{
		CompositeElementBase: cloned,
	}
}

// BubbleChart represents a c:bubbleChart element.
type BubbleChart struct {
	*openxml.CompositeElementBase
}

// NewBubbleChart creates a new bubble chart.
func NewBubbleChart() *BubbleChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"bubbleChart",
		PrefixChart,
	)
	bc := &BubbleChart{CompositeElementBase: elem}

	// Add varyColors
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrZero,
		),
	)
	bc.AppendChild(varyColors)

	return bc
}

// AddSeries adds a series to the bubble chart.
func (bc *BubbleChart) AddSeries(
	index, order uint32,
) *BubbleChartSeries {
	ser := NewBubbleChartSeries(index, order)
	bc.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (bc *BubbleChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	bc.AppendChild(axID)
}

// Clone creates a deep copy of this BubbleChart.
func (bc *BubbleChart) Clone() openxml.Element {
	clonedBase := bc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &BubbleChart{
		CompositeElementBase: cloned,
	}
}

// StockChart represents a c:stockChart element.
type StockChart struct {
	*openxml.CompositeElementBase
}

// NewStockChart creates a new stock chart.
func NewStockChart() *StockChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"stockChart",
		PrefixChart,
	)

	return &StockChart{CompositeElementBase: elem}
}

// AddSeries adds a series to the stock chart.
func (sc *StockChart) AddSeries(
	index, order uint32,
) *LineChartSeries {
	ser := NewLineChartSeries(index, order)
	sc.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (sc *StockChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	sc.AppendChild(axID)
}

// Clone creates a deep copy of this StockChart.
func (sc *StockChart) Clone() openxml.Element {
	clonedBase := sc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &StockChart{
		CompositeElementBase: cloned,
	}
}

// SurfaceChart represents a c:surfaceChart element.
type SurfaceChart struct {
	*openxml.CompositeElementBase
}

// NewSurfaceChart creates a new surface chart.
func NewSurfaceChart() *SurfaceChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"surfaceChart",
		PrefixChart,
	)

	return &SurfaceChart{
		CompositeElementBase: elem,
	}
}

// AddSeries adds a series to the surface chart.
func (sc *SurfaceChart) AddSeries(
	index, order uint32,
) *SurfaceChartSeries {
	ser := NewSurfaceChartSeries(index, order)
	sc.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (sc *SurfaceChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	sc.AppendChild(axID)
}

// Clone creates a deep copy of this SurfaceChart.
func (sc *SurfaceChart) Clone() openxml.Element {
	clonedBase := sc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &SurfaceChart{
		CompositeElementBase: cloned,
	}
}

// OfPieChart represents a c:ofPieChart element (pie-of-pie or bar-of-pie).
type OfPieChart struct {
	*openxml.CompositeElementBase
}

// NewOfPieChart creates a new of-pie chart with the specified type.
func NewOfPieChart(
	pieType OfPieTypeValue,
) *OfPieChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"ofPieChart",
		PrefixChart,
	)
	opc := &OfPieChart{CompositeElementBase: elem}

	// Add of-pie type
	typeElem := openxml.NewLeafElement(
		NamespaceChart,
		"ofPieType",
		PrefixChart,
	)
	typeElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(pieType),
		),
	)
	opc.AppendChild(typeElem)

	// Add varyColors (typically true for pie charts)
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrOne,
		),
	)
	opc.AppendChild(varyColors)

	return opc
}

// OfPieType returns the type of of-pie chart (pie or bar).
func (o *OfPieChart) OfPieType() OfPieTypeValue {
	elem := o.GetElement(
		"ofPieType",
		NamespaceChart,
	)
	if elem == nil {
		return OfPieTypePie
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return OfPieTypePie
	}

	return OfPieTypeValue(attr.Value())
}

// AddSeries adds a series to the of-pie chart.
func (o *OfPieChart) AddSeries(
	index, order uint32,
) *PieChartSeries {
	ser := NewPieChartSeries(index, order)
	o.AppendChild(ser)

	return ser
}

// SetGapWidth sets the gap width between the pie and secondary chart.
func (o *OfPieChart) SetGapWidth(width uint32) {
	existing := o.GetElement(
		"gapWidth",
		NamespaceChart,
	)
	if existing != nil {
		o.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"gapWidth",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(width),
				chartTypesBase10,
			),
		),
	)
	o.AppendChild(elem)
}

// SetSecondPieSize sets the size of the secondary pie/bar as percentage.
func (o *OfPieChart) SetSecondPieSize(
	size uint32,
) {
	existing := o.GetElement(
		"secondPieSize",
		NamespaceChart,
	)
	if existing != nil {
		o.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"secondPieSize",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(size),
				chartTypesBase10,
			),
		),
	)
	o.AppendChild(elem)
}

// Clone creates a deep copy of this OfPieChart.
func (o *OfPieChart) Clone() openxml.Element {
	clonedBase := o.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &OfPieChart{
		CompositeElementBase: cloned,
	}
}

// Bar3DChart represents a c:bar3DChart element.
type Bar3DChart struct {
	*openxml.CompositeElementBase
}

// NewBar3DChart creates a new 3D bar chart with the specified direction
// and grouping.
func NewBar3DChart(
	direction BarDirectionValue,
	grouping BarGroupingValue,
) *Bar3DChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"bar3DChart",
		PrefixChart,
	)
	bc := &Bar3DChart{CompositeElementBase: elem}

	// Add direction
	dirElem := openxml.NewLeafElement(
		NamespaceChart,
		"barDir",
		PrefixChart,
	)
	dirElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(direction),
		),
	)
	bc.AppendChild(dirElem)

	// Add grouping
	groupElem := openxml.NewLeafElement(
		NamespaceChart,
		elemGrouping,
		PrefixChart,
	)
	groupElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(grouping),
		),
	)
	bc.AppendChild(groupElem)

	// Add varyColors
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrZero,
		),
	)
	bc.AppendChild(varyColors)

	return bc
}

// AddSeries adds a series to the 3D bar chart.
func (bc *Bar3DChart) AddSeries(
	index, order uint32,
) *BarChartSeries {
	ser := NewBarChartSeries(index, order)
	bc.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (bc *Bar3DChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	bc.AppendChild(axID)
}

// Clone creates a deep copy of this Bar3DChart.
func (bc *Bar3DChart) Clone() openxml.Element {
	clonedBase := bc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Bar3DChart{
		CompositeElementBase: cloned,
	}
}

// Line3DChart represents a c:line3DChart element.
type Line3DChart struct {
	*openxml.CompositeElementBase
}

// NewLine3DChart creates a new 3D line chart with the specified grouping.
func NewLine3DChart(
	grouping GroupingValue,
) *Line3DChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"line3DChart",
		PrefixChart,
	)
	lc := &Line3DChart{CompositeElementBase: elem}

	// Add grouping
	groupElem := openxml.NewLeafElement(
		NamespaceChart,
		elemGrouping,
		PrefixChart,
	)
	groupElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(grouping),
		),
	)
	lc.AppendChild(groupElem)

	// Add varyColors
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrZero,
		),
	)
	lc.AppendChild(varyColors)

	return lc
}

// AddSeries adds a series to the 3D line chart.
func (lc *Line3DChart) AddSeries(
	index, order uint32,
) *LineChartSeries {
	ser := NewLineChartSeries(index, order)
	lc.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (lc *Line3DChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	lc.AppendChild(axID)
}

// Clone creates a deep copy of this Line3DChart.
func (lc *Line3DChart) Clone() openxml.Element {
	clonedBase := lc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Line3DChart{
		CompositeElementBase: cloned,
	}
}

// Pie3DChart represents a c:pie3DChart element.
type Pie3DChart struct {
	*openxml.CompositeElementBase
}

// NewPie3DChart creates a new 3D pie chart.
func NewPie3DChart() *Pie3DChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"pie3DChart",
		PrefixChart,
	)
	pc := &Pie3DChart{CompositeElementBase: elem}

	// Add varyColors (typically true for pie charts)
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrOne,
		),
	)
	pc.AppendChild(varyColors)

	return pc
}

// AddSeries adds a series to the 3D pie chart.
func (pc *Pie3DChart) AddSeries(
	index, order uint32,
) *PieChartSeries {
	ser := NewPieChartSeries(index, order)
	pc.AppendChild(ser)

	return ser
}

// Clone creates a deep copy of this Pie3DChart.
func (pc *Pie3DChart) Clone() openxml.Element {
	clonedBase := pc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Pie3DChart{
		CompositeElementBase: cloned,
	}
}

// Area3DChart represents a c:area3DChart element.
type Area3DChart struct {
	*openxml.CompositeElementBase
}

// NewArea3DChart creates a new 3D area chart with the specified grouping.
func NewArea3DChart(
	grouping GroupingValue,
) *Area3DChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"area3DChart",
		PrefixChart,
	)
	ac := &Area3DChart{CompositeElementBase: elem}

	// Add grouping
	groupElem := openxml.NewLeafElement(
		NamespaceChart,
		elemGrouping,
		PrefixChart,
	)
	groupElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(grouping),
		),
	)
	ac.AppendChild(groupElem)

	// Add varyColors
	varyColors := openxml.NewLeafElement(
		NamespaceChart,
		elemVaryColors,
		PrefixChart,
	)
	varyColors.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			attrZero,
		),
	)
	ac.AppendChild(varyColors)

	return ac
}

// AddSeries adds a series to the 3D area chart.
func (ac *Area3DChart) AddSeries(
	index, order uint32,
) *AreaChartSeries {
	ser := NewAreaChartSeries(index, order)
	ac.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (ac *Area3DChart) AddAxisID(axisID uint32) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	ac.AppendChild(axID)
}

// Clone creates a deep copy of this Area3DChart.
func (ac *Area3DChart) Clone() openxml.Element {
	clonedBase := ac.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Area3DChart{
		CompositeElementBase: cloned,
	}
}

// Surface3DChart represents a c:surface3DChart element.
type Surface3DChart struct {
	*openxml.CompositeElementBase
}

// NewSurface3DChart creates a new 3D surface chart.
func NewSurface3DChart() *Surface3DChart {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"surface3DChart",
		PrefixChart,
	)

	return &Surface3DChart{
		CompositeElementBase: elem,
	}
}

// AddSeries adds a series to the 3D surface chart.
func (sc *Surface3DChart) AddSeries(
	index, order uint32,
) *SurfaceChartSeries {
	ser := NewSurfaceChartSeries(index, order)
	sc.AppendChild(ser)

	return ser
}

// AddAxisID adds an axis ID reference.
func (sc *Surface3DChart) AddAxisID(
	axisID uint32,
) {
	axID := openxml.NewLeafElement(
		NamespaceChart,
		elemAxID,
		PrefixChart,
	)
	axID.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				chartTypesBase10,
			),
		),
	)
	sc.AppendChild(axID)
}

// Clone creates a deep copy of this Surface3DChart.
func (sc *Surface3DChart) Clone() openxml.Element {
	clonedBase := sc.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Surface3DChart{
		CompositeElementBase: cloned,
	}
}
