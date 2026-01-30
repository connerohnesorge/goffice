//revive:disable:file-length-limit chart elements need to be together
//revive:disable:max-public-structs chart elements have many public types

// and effects.
//
// This file implements chart element types (Title, Legend, DataLabels, etc.).
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Constants for chart elements.
const (
	// elemLegendPos is the "legendPos" element name.
	elemLegendPos = "legendPos"
	// elemTrendlineType is the "trendlineType" element name.
	elemTrendlineType = "trendlineType"
	// elementsBase10 is the base used for integer to string conversion.
	elementsBase10 = 10
	// elementsBitSize64 is the bit size for 64-bit float formatting.
	elementsBitSize64 = 64
)

// Title represents the c:title element for chart/axis titles.
type Title struct {
	*openxml.CompositeElementBase
}

// NewTitle creates a new title element.
func NewTitle() *Title {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"title",
		PrefixChart,
	)

	return &Title{CompositeElementBase: elem}
}

// NewTitleWithText creates a title with the specified text.
func NewTitleWithText(text string) *Title {
	t := NewTitle()

	// Create tx element with rich text
	tx := openxml.NewCompositeElement(
		NamespaceChart,
		"tx",
		PrefixChart,
	)
	rich := openxml.NewCompositeElement(
		NamespaceChart,
		"rich",
		PrefixChart,
	)

	// Body properties
	bodyPr := openxml.NewCompositeElement(
		NamespaceMain,
		"bodyPr",
		PrefixMain,
	)
	rich.AppendChild(bodyPr)

	// List style
	lstStyle := openxml.NewCompositeElement(
		NamespaceMain,
		"lstStyle",
		PrefixMain,
	)
	rich.AppendChild(lstStyle)

	// Paragraph with run
	p := openxml.NewCompositeElement(
		NamespaceMain,
		"p",
		PrefixMain,
	)
	r := openxml.NewCompositeElement(
		NamespaceMain,
		"r",
		PrefixMain,
	)
	tElem := openxml.NewLeafElementWithText(
		NamespaceMain,
		"t",
		PrefixMain,
		text,
	)
	r.AppendChild(tElem)
	p.AppendChild(r)
	rich.AppendChild(p)

	tx.AppendChild(rich)
	t.AppendChild(tx)

	return t
}

// SetOverlay sets whether the title overlays the chart.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Title) SetOverlay(overlay bool) {
	if existing := t.GetElement("overlay", NamespaceChart); existing != nil {
		t.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"overlay",
		PrefixChart,
	)
	if overlay {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				"val",
				"1",
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, "0"))
	}
	t.AppendChild(elem)
}

// SetLayout sets the layout for the title.
func (t *Title) SetLayout(layout *Layout) {
	if existing := t.GetElement("layout", NamespaceChart); existing != nil {
		t.RemoveChild(existing)
	}
	if layout != nil {
		t.AppendChild(layout)
	}
}

// Clone creates a deep copy of this Title.
func (t *Title) Clone() openxml.Element {
	clonedBase := t.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Title{CompositeElementBase: cloned}
}

// Legend represents the c:legend element.
type Legend struct {
	*openxml.CompositeElementBase
}

// NewLegend creates a new legend element.
func NewLegend() *Legend {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"legend",
		PrefixChart,
	)

	return &Legend{CompositeElementBase: elem}
}

// NewLegendWithPosition creates a legend with the specified position.
func NewLegendWithPosition(
	pos LegendPositionValue,
) *Legend {
	l := NewLegend()

	legendPosElem := openxml.NewLeafElement(
		NamespaceChart,
		elemLegendPos,
		PrefixChart,
	)
	legendPosElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(pos),
		),
	)
	l.AppendChild(legendPosElem)

	return l
}

// Position returns the legend position.
func (l *Legend) Position() LegendPositionValue {
	elem := l.GetElement(
		elemLegendPos,
		NamespaceChart,
	)
	if elem == nil {
		return LegendPositionRight
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return LegendPositionRight
	}

	return LegendPositionValue(attr.Value())
}

// SetPosition sets the legend position.
func (l *Legend) SetPosition(
	pos LegendPositionValue,
) {
	elem := l.GetElement(
		elemLegendPos,
		NamespaceChart,
	)
	if elem == nil {
		legendPosElem := openxml.NewLeafElement(
			NamespaceChart,
			elemLegendPos,
			PrefixChart,
		)
		legendPosElem.SetAttribute(
			openxml.NewSimpleAttribute(
				attrVal,
				string(pos),
			),
		)
		l.PrependChild(legendPosElem)

		return
	}
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(pos),
		),
	)
}

// SetOverlay sets whether the legend overlays the chart.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (l *Legend) SetOverlay(overlay bool) {
	if existing := l.GetElement("overlay", NamespaceChart); existing != nil {
		l.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"overlay",
		PrefixChart,
	)
	if overlay {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				"val",
				"1",
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, "0"))
	}
	l.AppendChild(elem)
}

// SetLayout sets the layout for the legend.
func (l *Legend) SetLayout(layout *Layout) {
	if existing := l.GetElement("layout", NamespaceChart); existing != nil {
		l.RemoveChild(existing)
	}
	if layout != nil {
		l.AppendChild(layout)
	}
}

// SetShapeProperties sets the shape properties for the legend.
func (l *Legend) SetShapeProperties(props *ChartShapeProperties) {
	if existing := l.GetElement(elemSpPr, NamespaceChart); existing != nil {
		l.RemoveChild(existing)
	}
	if props != nil {
		l.AppendChild(props)
	}
}

// SetTextProperties sets the text properties for the legend.
func (l *Legend) SetTextProperties(props *ChartText) {
	if existing := l.GetElement("txPr", NamespaceChart); existing != nil {
		l.RemoveChild(existing)
	}
	if props != nil {
		l.AppendChild(props)
	}
}

// Clone creates a deep copy of this Legend.
func (l *Legend) Clone() openxml.Element {
	clonedBase := l.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Legend{CompositeElementBase: cloned}
}

// DataLabels represents the c:dLbls element for data labels.
type DataLabels struct {
	*openxml.CompositeElementBase
}

// NewDataLabels creates a new data labels element.
func NewDataLabels() *DataLabels {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"dLbls",
		PrefixChart,
	)

	return &DataLabels{CompositeElementBase: elem}
}

// SetPosition sets the position of the data labels.
func (d *DataLabels) SetPosition(pos DataLabelPositionValue) {
	if existing := d.GetElement("dLblPos", NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"dLblPos",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(pos),
		),
	)
	d.AppendChild(elem)
}

// SetNumberFormat sets the number format for the data labels.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataLabels) SetNumberFormat(formatCode string, sourceLinked bool) {
	if existing := d.GetElement("numFmt", NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	numFmt := openxml.NewLeafElement(
		NamespaceChart,
		"numFmt",
		PrefixChart,
	)
	numFmt.SetAttribute(
		openxml.NewSimpleAttribute(
			"formatCode",
			formatCode,
		),
	)
	if sourceLinked {
		numFmt.SetAttribute(
			openxml.NewSimpleAttribute(
				"sourceLinked",
				"1",
			),
		)
	} else {
		numFmt.SetAttribute(openxml.NewSimpleAttribute("sourceLinked", "0"))
	}
	d.AppendChild(numFmt)
}

// SetShapeProperties sets the shape properties for the data labels.
func (d *DataLabels) SetShapeProperties(props *ChartShapeProperties) {
	if existing := d.GetElement(elemSpPr, NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	if props != nil {
		d.AppendChild(props)
	}
}

// SetTextProperties sets the text properties for the data labels.
func (d *DataLabels) SetTextProperties(props *ChartText) {
	if existing := d.GetElement("txPr", NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	if props != nil {
		d.AppendChild(props)
	}
}

// SetShowValue sets whether to show the value.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataLabels) SetShowValue(show bool) {
	if existing := d.GetElement("showVal", NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"showVal",
		PrefixChart,
	)
	if show {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				"val",
				"1",
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, "0"))
	}
	d.AppendChild(elem)
}

// SetShowCategoryName sets whether to show the category name.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataLabels) SetShowCategoryName(
	show bool,
) {
	existing := d.GetElement(
		"showCatName",
		NamespaceChart,
	)
	if existing != nil {
		d.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"showCatName",
		PrefixChart,
	)
	if show {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				"val",
				"1",
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, "0"))
	}
	d.AppendChild(elem)
}

// SetShowSeriesName sets whether to show the series name.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataLabels) SetShowSeriesName(
	show bool,
) {
	existing := d.GetElement(
		"showSerName",
		NamespaceChart,
	)
	if existing != nil {
		d.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"showSerName",
		PrefixChart,
	)
	if show {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				"val",
				"1",
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, "0"))
	}
	d.AppendChild(elem)
}

// SetShowPercent sets whether to show the percentage.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataLabels) SetShowPercent(show bool) {
	if existing := d.GetElement("showPercent", NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"showPercent",
		PrefixChart,
	)
	if show {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				"val",
				"1",
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, "0"))
	}
	d.AppendChild(elem)
}

// SetShowLegendKey sets whether to show the legend key.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataLabels) SetShowLegendKey(show bool) {
	existing := d.GetElement(
		"showLegendKey",
		NamespaceChart,
	)
	if existing != nil {
		d.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"showLegendKey",
		PrefixChart,
	)
	if show {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				"val",
				"1",
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, "0"))
	}
	d.AppendChild(elem)
}

// Clone creates a deep copy of this DataLabels.
func (d *DataLabels) Clone() openxml.Element {
	clonedBase := d.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &DataLabels{
		CompositeElementBase: cloned,
	}
}

// DataTable represents the c:dTable element for the data table.
type DataTable struct {
	*openxml.CompositeElementBase
}

// NewDataTable creates a new data table element.
func NewDataTable() *DataTable {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"dTable",
		PrefixChart,
	)

	return &DataTable{CompositeElementBase: elem}
}

// SetShowHorizontalBorder sets whether to show horizontal borders.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataTable) SetShowHorizontalBorder(show bool) {
	if existing := d.GetElement("showHorzBorder", NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"showHorzBorder",
		PrefixChart,
	)
	if show {
		elem.SetAttribute(openxml.NewSimpleAttribute("val", "1"))
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute("val", "0"))
	}
	d.AppendChild(elem)
}

// SetShowVerticalBorder sets whether to show vertical borders.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataTable) SetShowVerticalBorder(show bool) {
	if existing := d.GetElement("showVertBorder", NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"showVertBorder",
		PrefixChart,
	)
	if show {
		elem.SetAttribute(openxml.NewSimpleAttribute("val", "1"))
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute("val", "0"))
	}
	d.AppendChild(elem)
}

// SetShowOutline sets whether to show the outline border.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataTable) SetShowOutline(show bool) {
	if existing := d.GetElement("showOutline", NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"showOutline",
		PrefixChart,
	)
	if show {
		elem.SetAttribute(openxml.NewSimpleAttribute("val", "1"))
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute("val", "0"))
	}
	d.AppendChild(elem)
}

// SetShowKeys sets whether to show legend keys.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataTable) SetShowKeys(show bool) {
	if existing := d.GetElement("showKeys", NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"showKeys",
		PrefixChart,
	)
	if show {
		elem.SetAttribute(openxml.NewSimpleAttribute("val", "1"))
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute("val", "0"))
	}
	d.AppendChild(elem)
}

// SetShapeProperties sets the shape properties for the data table.
func (d *DataTable) SetShapeProperties(props *ChartShapeProperties) {
	if existing := d.GetElement(elemSpPr, NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	if props != nil {
		d.AppendChild(props)
	}
}

// SetTextProperties sets the text properties for the data table.
func (d *DataTable) SetTextProperties(props *ChartText) {
	if existing := d.GetElement("txPr", NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	if props != nil {
		d.AppendChild(props)
	}
}

// Clone creates a deep copy of this DataTable.
func (d *DataTable) Clone() openxml.Element {
	clonedBase := d.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &DataTable{CompositeElementBase: cloned}
}

// Marker represents the c:marker element for data point markers.
type Marker struct {
	*openxml.CompositeElementBase
}

// NewMarker creates a new marker element.
func NewMarker() *Marker {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"marker",
		PrefixChart,
	)

	return &Marker{CompositeElementBase: elem}
}

// NewMarkerWithStyle creates a marker with the specified style.
func NewMarkerWithStyle(
	style MarkerStyleValue,
) *Marker {
	m := NewMarker()
	symbol := openxml.NewLeafElement(
		NamespaceChart,
		"symbol",
		PrefixChart,
	)
	symbol.SetAttribute(
		openxml.NewSimpleAttribute(
			"val",
			string(style),
		),
	)
	m.AppendChild(symbol)

	return m
}

// SetStyle sets the marker style.
func (m *Marker) SetStyle(
	style MarkerStyleValue,
) {
	if existing := m.GetElement("symbol", NamespaceChart); existing != nil {
		m.RemoveChild(existing)
	}
	symbol := openxml.NewLeafElement(
		NamespaceChart,
		"symbol",
		PrefixChart,
	)
	symbol.SetAttribute(
		openxml.NewSimpleAttribute(
			"val",
			string(style),
		),
	)
	m.PrependChild(symbol)
}

// SetSize sets the marker size (2-72).
func (m *Marker) SetSize(size uint32) {
	if existing := m.GetElement("size", NamespaceChart); existing != nil {
		m.RemoveChild(existing)
	}
	sizeElem := openxml.NewLeafElement(
		NamespaceChart,
		"size",
		PrefixChart,
	)
	sizeElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(size),
				elementsBase10,
			),
		),
	)
	m.AppendChild(sizeElem)
}

// Clone creates a deep copy of this Marker.
func (m *Marker) Clone() openxml.Element {
	clonedBase := m.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Marker{CompositeElementBase: cloned}
}

// Layout represents the c:layout element.
type Layout struct {
	*openxml.CompositeElementBase
}

// NewLayout creates a new layout element.
func NewLayout() *Layout {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"layout",
		PrefixChart,
	)

	return &Layout{CompositeElementBase: elem}
}

// Clone creates a deep copy of this Layout.
func (l *Layout) Clone() openxml.Element {
	clonedBase := l.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Layout{CompositeElementBase: cloned}
}

// ChartShapeProperties represents the c:spPr element for chart shape
// properties. This wraps the standard DrawingML shape properties for use
// in charts.
type ChartShapeProperties struct {
	*openxml.CompositeElementBase
}

// NewChartShapeProperties creates a new chart shape properties element.
func NewChartShapeProperties() *ChartShapeProperties {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		elemSpPr,
		PrefixChart,
	)

	return &ChartShapeProperties{
		CompositeElementBase: elem,
	}
}

// SetSolidFill sets a solid fill color for the shape.
func (c *ChartShapeProperties) SetSolidFill(
	hexColor string,
) {
	// Remove existing fill
	c.removeFill()

	solidFill := openxml.NewCompositeElement(
		NamespaceMain,
		"solidFill",
		PrefixMain,
	)
	srgbClr := openxml.NewLeafElement(
		NamespaceMain,
		"srgbClr",
		PrefixMain,
	)
	srgbClr.SetAttribute(
		openxml.NewSimpleAttribute(
			"val",
			hexColor,
		),
	)
	solidFill.AppendChild(srgbClr)
	c.AppendChild(solidFill)
}

// SetNoFill sets no fill for the shape.
func (c *ChartShapeProperties) SetNoFill() {
	c.removeFill()
	noFill := openxml.NewLeafElement(
		NamespaceMain,
		"noFill",
		PrefixMain,
	)
	c.AppendChild(noFill)
}

// removeFill removes any existing fill.
func (c *ChartShapeProperties) removeFill() {
	fillTypes := []string{
		"noFill",
		"solidFill",
		"gradFill",
		"pattFill",
	}
	for _, name := range fillTypes {
		if elem := c.GetElement(name, NamespaceMain); elem != nil {
			c.RemoveChild(elem)
		}
	}
}

// SetOutline sets an outline for the shape.
func (c *ChartShapeProperties) SetOutline(
	width EMU,
	hexColor string,
) {
	if existing := c.GetElement("ln", NamespaceMain); existing != nil {
		c.RemoveChild(existing)
	}

	ln := openxml.NewCompositeElement(
		NamespaceMain,
		"ln",
		PrefixMain,
	)
	ln.SetAttribute(
		openxml.NewSimpleAttribute(
			"w",
			strconv.FormatInt(
				int64(width),
				elementsBase10,
			),
		),
	)

	solidFill := openxml.NewCompositeElement(
		NamespaceMain,
		"solidFill",
		PrefixMain,
	)
	srgbClr := openxml.NewLeafElement(
		NamespaceMain,
		"srgbClr",
		PrefixMain,
	)
	srgbClr.SetAttribute(
		openxml.NewSimpleAttribute(
			"val",
			hexColor,
		),
	)
	solidFill.AppendChild(srgbClr)
	ln.AppendChild(solidFill)

	c.AppendChild(ln)
}

// Clone creates a deep copy of this ChartShapeProperties.
func (c *ChartShapeProperties) Clone() openxml.Element {
	clonedBase := c.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &ChartShapeProperties{
		CompositeElementBase: cloned,
	}
}

// ErrorBars represents the c:errBars element for error bars.
type ErrorBars struct {
	*openxml.CompositeElementBase
}

// NewErrorBars creates a new error bars element.
func NewErrorBars() *ErrorBars {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"errBars",
		PrefixChart,
	)

	return &ErrorBars{CompositeElementBase: elem}
}

// Clone creates a deep copy of this ErrorBars.
func (e *ErrorBars) Clone() openxml.Element {
	clonedBase := e.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &ErrorBars{
		CompositeElementBase: cloned,
	}
}

// DataPoint represents the c:dPt element for individual data point styling.
type DataPoint struct {
	*openxml.CompositeElementBase
}

// NewDataPoint creates a new data point element.
func NewDataPoint(index uint32) *DataPoint {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"dPt",
		PrefixChart,
	)
	dp := &DataPoint{CompositeElementBase: elem}

	// Add index
	idx := openxml.NewLeafElement(
		NamespaceChart,
		"idx",
		PrefixChart,
	)
	idx.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(index),
				elementsBase10,
			),
		),
	)
	dp.AppendChild(idx)

	return dp
}

// SetShapeProperties sets the shape properties for the data point.
func (d *DataPoint) SetShapeProperties(
	props *ChartShapeProperties,
) {
	if existing := d.GetElement(elemSpPr, NamespaceChart); existing != nil {
		d.RemoveChild(existing)
	}
	if props != nil {
		d.AppendChild(props)
	}
}

// Clone creates a deep copy of this DataPoint.
func (d *DataPoint) Clone() openxml.Element {
	clonedBase := d.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &DataPoint{
		CompositeElementBase: cloned,
	}
}

// TrendlineTypeValue represents the type of trendline.
type TrendlineTypeValue string

// Trendline type values.
const (
	TrendlineTypeExp       TrendlineTypeValue = "exp"
	TrendlineTypeLinear    TrendlineTypeValue = "linear"
	TrendlineTypeLog       TrendlineTypeValue = "log"
	TrendlineTypeMovingAvg TrendlineTypeValue = "movingAvg"
	TrendlineTypePoly      TrendlineTypeValue = "poly"
	TrendlineTypePower     TrendlineTypeValue = "power"
)

// Trendline represents the c:trendline element for trendlines on chart series.
type Trendline struct {
	*openxml.CompositeElementBase
}

// NewTrendline creates a new trendline element.
func NewTrendline() *Trendline {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"trendline",
		PrefixChart,
	)

	return &Trendline{CompositeElementBase: elem}
}

// NewTrendlineWithType creates a trendline with the specified type.
func NewTrendlineWithType(
	trendlineType TrendlineTypeValue,
) *Trendline {
	t := NewTrendline()

	typeElem := openxml.NewLeafElement(
		NamespaceChart,
		elemTrendlineType,
		PrefixChart,
	)
	typeElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(trendlineType),
		),
	)
	t.AppendChild(typeElem)

	return t
}

// TrendlineType returns the type of trendline.
func (t *Trendline) TrendlineType() TrendlineTypeValue {
	elem := t.GetElement(
		elemTrendlineType,
		NamespaceChart,
	)
	if elem == nil {
		return TrendlineTypeLinear
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return TrendlineTypeLinear
	}

	return TrendlineTypeValue(attr.Value())
}

// SetTrendlineType sets the type of trendline.
func (t *Trendline) SetTrendlineType(
	trendlineType TrendlineTypeValue,
) {
	existing := t.GetElement(
		elemTrendlineType,
		NamespaceChart,
	)
	if existing != nil {
		t.RemoveChild(existing)
	}
	typeElem := openxml.NewLeafElement(
		NamespaceChart,
		elemTrendlineType,
		PrefixChart,
	)
	typeElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(trendlineType),
		),
	)
	t.AppendChild(typeElem)
}

// SetName sets the name of the trendline.
func (t *Trendline) SetName(name string) {
	if existing := t.GetElement("name", NamespaceChart); existing != nil {
		t.RemoveChild(existing)
	}
	nameElem := openxml.NewLeafElementWithText(
		NamespaceChart,
		"name",
		PrefixChart,
		name,
	)
	t.PrependChild(nameElem)
}

// SetOrder sets the polynomial order (for polynomial trendlines).
func (t *Trendline) SetOrder(order uint32) {
	if existing := t.GetElement("order", NamespaceChart); existing != nil {
		t.RemoveChild(existing)
	}
	orderElem := openxml.NewLeafElement(
		NamespaceChart,
		"order",
		PrefixChart,
	)
	orderElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(order),
				elementsBase10,
			),
		),
	)
	t.AppendChild(orderElem)
}

// SetPeriod sets the moving average period.
func (t *Trendline) SetPeriod(period uint32) {
	if existing := t.GetElement("period", NamespaceChart); existing != nil {
		t.RemoveChild(existing)
	}
	periodElem := openxml.NewLeafElement(
		NamespaceChart,
		"period",
		PrefixChart,
	)
	periodElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(period),
				elementsBase10,
			),
		),
	)
	t.AppendChild(periodElem)
}

// SetForward sets the forward projection value.
func (t *Trendline) SetForward(forward float64) {
	if existing := t.GetElement("forward", NamespaceChart); existing != nil {
		t.RemoveChild(existing)
	}
	forwardElem := openxml.NewLeafElement(
		NamespaceChart,
		"forward",
		PrefixChart,
	)
	forwardElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatFloat(
				forward,
				'f',
				-1,
				elementsBitSize64,
			),
		),
	)
	t.AppendChild(forwardElem)
}

// SetBackward sets the backward projection value.
func (t *Trendline) SetBackward(
	backward float64,
) {
	if existing := t.GetElement("backward", NamespaceChart); existing != nil {
		t.RemoveChild(existing)
	}
	backwardElem := openxml.NewLeafElement(
		NamespaceChart,
		"backward",
		PrefixChart,
	)
	backwardElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatFloat(
				backward,
				'f',
				-1,
				elementsBitSize64,
			),
		),
	)
	t.AppendChild(backwardElem)
}

// SetIntercept sets the intercept value.
func (t *Trendline) SetIntercept(
	intercept float64,
) {
	if existing := t.GetElement("intercept", NamespaceChart); existing != nil {
		t.RemoveChild(existing)
	}
	interceptElem := openxml.NewLeafElement(
		NamespaceChart,
		"intercept",
		PrefixChart,
	)
	interceptElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatFloat(
				intercept,
				'f',
				-1,
				elementsBitSize64,
			),
		),
	)
	t.AppendChild(interceptElem)
}

// SetDisplayRSquared sets whether to display R-squared value.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Trendline) SetDisplayRSquared(
	display bool,
) {
	if existing := t.GetElement("dispRSqr", NamespaceChart); existing != nil {
		t.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"dispRSqr",
		PrefixChart,
	)
	if display {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				attrVal,
				"1",
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, "0"))
	}
	t.AppendChild(elem)
}

// SetDisplayEquation sets whether to display the equation.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Trendline) SetDisplayEquation(
	display bool,
) {
	if existing := t.GetElement("dispEq", NamespaceChart); existing != nil {
		t.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"dispEq",
		PrefixChart,
	)
	if display {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				attrVal,
				"1",
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, "0"))
	}
	t.AppendChild(elem)
}

// Clone creates a deep copy of this Trendline.
func (t *Trendline) Clone() openxml.Element {
	clonedBase := t.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Trendline{
		CompositeElementBase: cloned,
	}
}

// ChartText represents a text body in chart context (c:rich or c:txPr).
type ChartText struct {
	*TextBody
}

// NewChartText creates a new ChartText element with the given local name.
// name should be "rich" (for title) or "txPr" (for legend/datalabels).
func NewChartText(name string) *ChartText {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		name,
		PrefixChart,
	)
	tb := &TextBody{CompositeElementBase: elem}
	// Add default body properties
	tb.AppendChild(NewTextBodyProperties())
	// Add default list style
	lstStyle := openxml.NewCompositeElement(
		NamespaceMain,
		"lstStyle",
		PrefixMain,
	)
	tb.AppendChild(lstStyle)

	return &ChartText{TextBody: tb}
}

// Clone creates a deep copy of this ChartText.
func (ct *ChartText) Clone() openxml.Element {
	clonedBase := ct.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}
	// Re-wrap in TextBody
	tb := &TextBody{CompositeElementBase: cloned}

	return &ChartText{TextBody: tb}
}
