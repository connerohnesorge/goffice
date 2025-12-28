//revive:disable:file-length-limit comprehensive chart series implementation

// Package drawingml provides shared DrawingML types for shapes, images,
// and effects.
//
// This file implements chart series types.
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Constants for chart series elements.
const (
	// seriesBase10 is the base used for integer to string conversion.
	seriesBase10 = 10
	// seriesBitSize32 is the bit size for 32-bit integer parsing.
	seriesBitSize32 = 32
)

// Chart series element name constants.
const (
	elemOrder = "order"
	elemIdx   = "idx"
	elemSer   = "ser"
	elemCat   = "cat"
)

// BarChartSeries represents a c:ser element for bar charts.
type BarChartSeries struct {
	*openxml.CompositeElementBase
}

// NewBarChartSeries creates a new bar chart series.
func NewBarChartSeries(
	index, order uint32,
) *BarChartSeries {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		elemSer,
		PrefixChart,
	)
	ser := &BarChartSeries{
		CompositeElementBase: elem,
	}

	// Add index
	idxElem := openxml.NewLeafElement(
		NamespaceChart,
		elemIdx,
		PrefixChart,
	)
	idxElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(index),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(idxElem)

	// Add order
	orderElem := openxml.NewLeafElement(
		NamespaceChart,
		elemOrder,
		PrefixChart,
	)
	orderElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(order),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(orderElem)

	return ser
}

// Index returns the series index.
func (s *BarChartSeries) Index() uint32 {
	return s.getUintVal(elemIdx)
}

// Order returns the series order.
func (s *BarChartSeries) Order() uint32 {
	return s.getUintVal(elemOrder)
}

// SetSeriesText sets the series text (name).
func (s *BarChartSeries) SetSeriesText(
	text *SeriesText,
) {
	if existing := s.GetElement("tx", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if text != nil {
		s.insertAfterOrder(text)
	}
}

// SetShapeProperties sets the shape properties for the series.
func (s *BarChartSeries) SetShapeProperties(
	props *ChartShapeProperties,
) {
	if existing := s.GetElement("spPr", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if props != nil {
		s.AppendChild(props)
	}
}

// SetCategoryAxisData sets the category data reference.
func (s *BarChartSeries) SetCategoryAxisData(
	data *CategoryAxisData,
) {
	if existing := s.GetElement(elemCat, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if data != nil {
		s.AppendChild(data)
	}
}

// SetValues sets the values data reference.
func (s *BarChartSeries) SetValues(
	values *Values,
) {
	if existing := s.GetElement(attrVal, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if values != nil {
		s.AppendChild(values)
	}
}

// getUintVal gets an unsigned integer attrValue from a child element.
func (s *BarChartSeries) getUintVal(
	localName string,
) uint32 {
	elem := s.GetElement(
		localName,
		NamespaceChart,
	)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return 0
	}
	parsedVal, _ := strconv.ParseUint(
		attr.Value(),
		seriesBase10,
		seriesBitSize32,
	)

	return uint32(parsedVal)
}

// insertAfterOrder inserts an element after the order element.
func (s *BarChartSeries) insertAfterOrder(
	elem openxml.Element,
) {
	order := s.GetElement(
		elemOrder,
		NamespaceChart,
	)
	if order != nil {
		s.InsertAfter(elem, order)
	} else {
		s.AppendChild(elem)
	}
}

// Clone creates a deep copy of this BarChartSeries.
func (s *BarChartSeries) Clone() openxml.Element {
	clonedBase := s.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &BarChartSeries{
		CompositeElementBase: cloned,
	}
}

// LineChartSeries represents a c:ser element for line charts.
type LineChartSeries struct {
	*openxml.CompositeElementBase
}

// NewLineChartSeries creates a new line chart series.
func NewLineChartSeries(
	index, order uint32,
) *LineChartSeries {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		elemSer,
		PrefixChart,
	)
	ser := &LineChartSeries{
		CompositeElementBase: elem,
	}

	// Add index
	idxElem := openxml.NewLeafElement(
		NamespaceChart,
		elemIdx,
		PrefixChart,
	)
	idxElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(index),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(idxElem)

	// Add order
	orderElem := openxml.NewLeafElement(
		NamespaceChart,
		elemOrder,
		PrefixChart,
	)
	orderElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(order),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(orderElem)

	return ser
}

// Index returns the series index.
func (s *LineChartSeries) Index() uint32 {
	elem := s.GetElement(elemIdx, NamespaceChart)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return 0
	}
	parsedVal, _ := strconv.ParseUint(
		attr.Value(),
		seriesBase10,
		seriesBitSize32,
	)

	return uint32(parsedVal)
}

// SetSeriesText sets the series text (name).
func (s *LineChartSeries) SetSeriesText(
	text *SeriesText,
) {
	existing := s.GetElement("tx", NamespaceChart)
	if existing != nil {
		s.RemoveChild(existing)
	}
	if text == nil {
		return
	}
	order := s.GetElement(
		elemOrder,
		NamespaceChart,
	)
	if order != nil {
		s.InsertAfter(text, order)
	} else {
		s.AppendChild(text)
	}
}

// SetShapeProperties sets the shape properties for the series.
func (s *LineChartSeries) SetShapeProperties(
	props *ChartShapeProperties,
) {
	if existing := s.GetElement("spPr", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if props != nil {
		s.AppendChild(props)
	}
}

// SetMarker sets the marker for the series.
func (s *LineChartSeries) SetMarker(
	marker *Marker,
) {
	if existing := s.GetElement("marker", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if marker != nil {
		s.AppendChild(marker)
	}
}

// SetCategoryAxisData sets the category data reference.
func (s *LineChartSeries) SetCategoryAxisData(
	data *CategoryAxisData,
) {
	if existing := s.GetElement(elemCat, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if data != nil {
		s.AppendChild(data)
	}
}

// SetValues sets the values data reference.
func (s *LineChartSeries) SetValues(
	values *Values,
) {
	if existing := s.GetElement(attrVal, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if values != nil {
		s.AppendChild(values)
	}
}

// SetSmooth sets whether the line is smoothed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (s *LineChartSeries) SetSmooth(smooth bool) {
	if existing := s.GetElement("smooth", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"smooth",
		PrefixChart,
	)
	if smooth {
		elem.SetAttribute(
			openxml.NewSimpleAttribute(
				attrVal,
				"1",
			),
		)
	} else {
		elem.SetAttribute(openxml.NewSimpleAttribute(attrVal, "0"))
	}
	s.AppendChild(elem)
}

// Clone creates a deep copy of this LineChartSeries.
func (s *LineChartSeries) Clone() openxml.Element {
	clonedBase := s.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &LineChartSeries{
		CompositeElementBase: cloned,
	}
}

// PieChartSeries represents a c:ser element for pie/doughnut charts.
type PieChartSeries struct {
	*openxml.CompositeElementBase
}

// NewPieChartSeries creates a new pie chart series.
func NewPieChartSeries(
	index, order uint32,
) *PieChartSeries {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		elemSer,
		PrefixChart,
	)
	ser := &PieChartSeries{
		CompositeElementBase: elem,
	}

	// Add index
	idxElem := openxml.NewLeafElement(
		NamespaceChart,
		elemIdx,
		PrefixChart,
	)
	idxElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(index),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(idxElem)

	// Add order
	orderElem := openxml.NewLeafElement(
		NamespaceChart,
		elemOrder,
		PrefixChart,
	)
	orderElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(order),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(orderElem)

	return ser
}

// SetCategoryAxisData sets the category data reference.
func (s *PieChartSeries) SetCategoryAxisData(
	data *CategoryAxisData,
) {
	if existing := s.GetElement(elemCat, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if data != nil {
		s.AppendChild(data)
	}
}

// SetValues sets the values data reference.
func (s *PieChartSeries) SetValues(
	values *Values,
) {
	if existing := s.GetElement(attrVal, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if values != nil {
		s.AppendChild(values)
	}
}

// SetExplosion sets the explosion value for pie slices.
func (s *PieChartSeries) SetExplosion(
	explosionVal uint32,
) {
	if existing := s.GetElement("explosion", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"explosion",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(explosionVal),
				seriesBase10,
			),
		),
	)
	s.AppendChild(elem)
}

// Clone creates a deep copy of this PieChartSeries.
func (s *PieChartSeries) Clone() openxml.Element {
	clonedBase := s.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &PieChartSeries{
		CompositeElementBase: cloned,
	}
}

// AreaChartSeries represents a c:ser element for area charts.
type AreaChartSeries struct {
	*openxml.CompositeElementBase
}

// NewAreaChartSeries creates a new area chart series.
func NewAreaChartSeries(
	index, order uint32,
) *AreaChartSeries {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		elemSer,
		PrefixChart,
	)
	ser := &AreaChartSeries{
		CompositeElementBase: elem,
	}

	// Add index
	idxElem := openxml.NewLeafElement(
		NamespaceChart,
		elemIdx,
		PrefixChart,
	)
	idxElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(index),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(idxElem)

	// Add order
	orderElem := openxml.NewLeafElement(
		NamespaceChart,
		elemOrder,
		PrefixChart,
	)
	orderElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(order),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(orderElem)

	return ser
}

// SetCategoryAxisData sets the category data reference.
func (s *AreaChartSeries) SetCategoryAxisData(
	data *CategoryAxisData,
) {
	if existing := s.GetElement(elemCat, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if data != nil {
		s.AppendChild(data)
	}
}

// SetValues sets the values data reference.
func (s *AreaChartSeries) SetValues(
	values *Values,
) {
	if existing := s.GetElement(attrVal, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if values != nil {
		s.AppendChild(values)
	}
}

// Clone creates a deep copy of this AreaChartSeries.
func (s *AreaChartSeries) Clone() openxml.Element {
	clonedBase := s.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &AreaChartSeries{
		CompositeElementBase: cloned,
	}
}

// ScatterChartSeries represents a c:ser element for scatter charts.
type ScatterChartSeries struct {
	*openxml.CompositeElementBase
}

// NewScatterChartSeries creates a new scatter chart series.
func NewScatterChartSeries(
	index, order uint32,
) *ScatterChartSeries {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		elemSer,
		PrefixChart,
	)
	ser := &ScatterChartSeries{
		CompositeElementBase: elem,
	}

	// Add index
	idxElem := openxml.NewLeafElement(
		NamespaceChart,
		elemIdx,
		PrefixChart,
	)
	idxElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(index),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(idxElem)

	// Add order
	orderElem := openxml.NewLeafElement(
		NamespaceChart,
		elemOrder,
		PrefixChart,
	)
	orderElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(order),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(orderElem)

	return ser
}

// SetXValues sets the X values data reference.
func (s *ScatterChartSeries) SetXValues(
	values *XValues,
) {
	if existing := s.GetElement("xVal", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if values != nil {
		s.AppendChild(values)
	}
}

// SetYValues sets the Y values data reference.
func (s *ScatterChartSeries) SetYValues(
	values *YValues,
) {
	if existing := s.GetElement("yVal", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if values != nil {
		s.AppendChild(values)
	}
}

// SetMarker sets the marker for the series.
func (s *ScatterChartSeries) SetMarker(
	marker *Marker,
) {
	if existing := s.GetElement("marker", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if marker != nil {
		s.AppendChild(marker)
	}
}

// Clone creates a deep copy of this ScatterChartSeries.
func (s *ScatterChartSeries) Clone() openxml.Element {
	clonedBase := s.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &ScatterChartSeries{
		CompositeElementBase: cloned,
	}
}

// BubbleChartSeries represents a c:ser element for bubble charts.
type BubbleChartSeries struct {
	*openxml.CompositeElementBase
}

// NewBubbleChartSeries creates a new bubble chart series.
func NewBubbleChartSeries(
	index, order uint32,
) *BubbleChartSeries {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		elemSer,
		PrefixChart,
	)
	ser := &BubbleChartSeries{
		CompositeElementBase: elem,
	}

	// Add index
	idxElem := openxml.NewLeafElement(
		NamespaceChart,
		elemIdx,
		PrefixChart,
	)
	idxElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(index),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(idxElem)

	// Add order
	orderElem := openxml.NewLeafElement(
		NamespaceChart,
		elemOrder,
		PrefixChart,
	)
	orderElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(order),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(orderElem)

	return ser
}

// SetXValues sets the X values data reference.
func (s *BubbleChartSeries) SetXValues(
	values *XValues,
) {
	if existing := s.GetElement("xVal", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if values != nil {
		s.AppendChild(values)
	}
}

// SetYValues sets the Y values data reference.
func (s *BubbleChartSeries) SetYValues(
	values *YValues,
) {
	if existing := s.GetElement("yVal", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if values != nil {
		s.AppendChild(values)
	}
}

// SetBubbleSize sets the bubble size data reference.
func (s *BubbleChartSeries) SetBubbleSize(
	values *BubbleSize,
) {
	if existing := s.GetElement("bubbleSize", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if values != nil {
		s.AppendChild(values)
	}
}

// Clone creates a deep copy of this BubbleChartSeries.
func (s *BubbleChartSeries) Clone() openxml.Element {
	clonedBase := s.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &BubbleChartSeries{
		CompositeElementBase: cloned,
	}
}

// SurfaceChartSeries represents a c:ser element for surface charts.
type SurfaceChartSeries struct {
	*openxml.CompositeElementBase
}

// NewSurfaceChartSeries creates a new surface chart series.
func NewSurfaceChartSeries(
	index, order uint32,
) *SurfaceChartSeries {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		elemSer,
		PrefixChart,
	)
	ser := &SurfaceChartSeries{
		CompositeElementBase: elem,
	}

	// Add index
	idxElem := openxml.NewLeafElement(
		NamespaceChart,
		elemIdx,
		PrefixChart,
	)
	idxElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(index),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(idxElem)

	// Add order
	orderElem := openxml.NewLeafElement(
		NamespaceChart,
		elemOrder,
		PrefixChart,
	)
	orderElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(order),
				seriesBase10,
			),
		),
	)
	ser.AppendChild(orderElem)

	return ser
}

// SetCategoryAxisData sets the category data reference.
func (s *SurfaceChartSeries) SetCategoryAxisData(
	data *CategoryAxisData,
) {
	if existing := s.GetElement(elemCat, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if data != nil {
		s.AppendChild(data)
	}
}

// SetValues sets the values data reference.
func (s *SurfaceChartSeries) SetValues(
	values *Values,
) {
	if existing := s.GetElement(attrVal, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	if values != nil {
		s.AppendChild(values)
	}
}

// Clone creates a deep copy of this SurfaceChartSeries.
func (s *SurfaceChartSeries) Clone() openxml.Element {
	clonedBase := s.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &SurfaceChartSeries{
		CompositeElementBase: cloned,
	}
}
