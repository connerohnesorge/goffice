//revive:disable:file-length-limit chart data types need to be together
//revive:disable:max-public-structs chart data has many public types

// Package drawingml implements the DrawingML (Drawing Markup Language) types
// for Office Open XML documents, including shapes, pictures, charts, and effects.
//
// This file implements chart data reference types.
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Element name constants for chart data.
const (
	// elemF is the "f" (formula) element name.
	elemF = "f"
	// elemPtCount is the "ptCount" element name.
	elemPtCount = "ptCount"
	// dataBase10 is the base used for integer to string conversion.
	dataBase10 = 10
	// floatBitSize64 is the bit size for 64-bit float formatting.
	floatBitSize64 = 64
)

// CategoryAxisData represents the c:cat element for category data.
type CategoryAxisData struct {
	*openxml.CompositeElementBase
}

// NewCategoryAxisData creates a new category axis data element.
func NewCategoryAxisData() *CategoryAxisData {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"cat",
		PrefixChart,
	)

	return &CategoryAxisData{
		CompositeElementBase: elem,
	}
}

// SetStringReference sets a string reference for the category data.
func (c *CategoryAxisData) SetStringReference(
	formula string,
) {
	c.RemoveAllChildren()
	strRef := NewStringReference(formula)
	c.AppendChild(strRef)
}

// SetNumberReference sets a number reference for the category data.
func (c *CategoryAxisData) SetNumberReference(
	formula string,
) {
	c.RemoveAllChildren()
	numRef := NewNumberReference(formula)
	c.AppendChild(numRef)
}

// Clone creates a deep copy of this CategoryAxisData.
func (c *CategoryAxisData) Clone() openxml.Element {
	cloneBase := c.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &CategoryAxisData{
		CompositeElementBase: clone,
	}
}

// Values represents the c:val element for numeric values.
type Values struct {
	*openxml.CompositeElementBase
}

// NewValues creates a new values element.
func NewValues() *Values {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"val",
		PrefixChart,
	)

	return &Values{CompositeElementBase: elem}
}

// SetNumberReference sets a number reference for the values.
func (v *Values) SetNumberReference(
	formula string,
) {
	v.RemoveAllChildren()
	numRef := NewNumberReference(formula)
	v.AppendChild(numRef)
}

// Clone creates a deep copy of this Values.
func (v *Values) Clone() openxml.Element {
	cloneBase := v.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &Values{CompositeElementBase: clone}
}

// XValues represents the c:xVal element for X values in scatter/bubble
// charts.
type XValues struct {
	*openxml.CompositeElementBase
}

// NewXValues creates a new X values element.
func NewXValues() *XValues {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"xVal",
		PrefixChart,
	)

	return &XValues{CompositeElementBase: elem}
}

// SetNumberReference sets a number reference for the X values.
func (x *XValues) SetNumberReference(
	formula string,
) {
	x.RemoveAllChildren()
	numRef := NewNumberReference(formula)
	x.AppendChild(numRef)
}

// Clone creates a deep copy of this XValues.
func (x *XValues) Clone() openxml.Element {
	cloneBase := x.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &XValues{CompositeElementBase: clone}
}

// YValues represents the c:yVal element for Y values in scatter/bubble
// charts.
type YValues struct {
	*openxml.CompositeElementBase
}

// NewYValues creates a new Y values element.
func NewYValues() *YValues {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"yVal",
		PrefixChart,
	)

	return &YValues{CompositeElementBase: elem}
}

// SetNumberReference sets a number reference for the Y values.
func (y *YValues) SetNumberReference(
	formula string,
) {
	y.RemoveAllChildren()
	numRef := NewNumberReference(formula)
	y.AppendChild(numRef)
}

// Clone creates a deep copy of this YValues.
func (y *YValues) Clone() openxml.Element {
	cloneBase := y.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &YValues{CompositeElementBase: clone}
}

// BubbleSize represents the c:bubbleSize element for bubble chart sizes.
type BubbleSize struct {
	*openxml.CompositeElementBase
}

// NewBubbleSize creates a new bubble size element.
func NewBubbleSize() *BubbleSize {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"bubbleSize",
		PrefixChart,
	)

	return &BubbleSize{CompositeElementBase: elem}
}

// SetNumberReference sets a number reference for the bubble sizes.
func (b *BubbleSize) SetNumberReference(
	formula string,
) {
	b.RemoveAllChildren()
	numRef := NewNumberReference(formula)
	b.AppendChild(numRef)
}

// Clone creates a deep copy of this BubbleSize.
func (b *BubbleSize) Clone() openxml.Element {
	cloneBase := b.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &BubbleSize{
		CompositeElementBase: clone,
	}
}

// StringReference represents the c:strRef element for string data
// references.
type StringReference struct {
	*openxml.CompositeElementBase
}

// NewStringReference creates a new string reference with the given formula.
func NewStringReference(
	formula string,
) *StringReference {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"strRef",
		PrefixChart,
	)
	sr := &StringReference{
		CompositeElementBase: elem,
	}

	// Add formula
	f := openxml.NewLeafElementWithText(
		NamespaceChart,
		elemF,
		PrefixChart,
		formula,
	)
	sr.AppendChild(f)

	return sr
}

// NewStringReferenceWithCache creates a new string reference with the given
// formula and cached values.
func NewStringReferenceWithCache(
	formula string,
	values []string,
) *StringReference {
	sr := NewStringReference(formula)
	sr.SetCache(values)

	return sr
}

// Formula returns the formula reference.
func (s *StringReference) Formula() string {
	elem := s.GetElement(elemF, NamespaceChart)
	if elem == nil {
		return ""
	}
	leaf, ok := elem.(*openxml.LeafElementBase)
	if !ok {
		return ""
	}

	return leaf.InnerText()
}

// SetFormula sets the formula reference.
func (s *StringReference) SetFormula(
	formula string,
) {
	existing := s.GetElement(
		elemF,
		NamespaceChart,
	)
	if existing != nil {
		s.RemoveChild(existing)
	}
	f := openxml.NewLeafElementWithText(
		NamespaceChart,
		elemF,
		PrefixChart,
		formula,
	)
	s.PrependChild(f)
}

// SetCache sets the cached string values for this reference.
// This creates a c:strCache element with the provided values.
func (s *StringReference) SetCache(
	values []string,
) {
	// Remove existing cache if present
	if existing := s.GetElement("strCache", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}

	// Create new cache
	cache := NewStringCache()
	cache.SetPointCount(uint32(len(values)))

	// Add all points
	for i, val := range values {
		cache.AddPoint(uint32(i), val)
	}

	s.AppendChild(cache)
}

// Clone creates a deep copy of this StringReference.
func (s *StringReference) Clone() openxml.Element {
	cloneBase := s.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &StringReference{
		CompositeElementBase: clone,
	}
}

// NumberReference represents the c:numRef element for numeric data
// references.
type NumberReference struct {
	*openxml.CompositeElementBase
}

// NewNumberReference creates a new number reference with the given formula.
func NewNumberReference(
	formula string,
) *NumberReference {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"numRef",
		PrefixChart,
	)
	nr := &NumberReference{
		CompositeElementBase: elem,
	}

	// Add formula
	f := openxml.NewLeafElementWithText(
		NamespaceChart,
		elemF,
		PrefixChart,
		formula,
	)
	nr.AppendChild(f)

	return nr
}

// NewNumberReferenceWithCache creates a new number reference with the given
// formula and cached values using "General" format.
func NewNumberReferenceWithCache(
	formula string,
	values []float64,
) *NumberReference {
	return NewNumberReferenceWithCacheAndFormat(
		formula,
		values,
		"General",
	)
}

// NewNumberReferenceWithCacheAndFormat creates a new number reference with
// the given formula, cached values, and format code.
func NewNumberReferenceWithCacheAndFormat(
	formula string,
	values []float64,
	formatCode string,
) *NumberReference {
	nr := NewNumberReference(formula)
	nr.SetCacheWithFormat(values, formatCode)

	return nr
}

// Formula returns the formula reference.
func (n *NumberReference) Formula() string {
	elem := n.GetElement(elemF, NamespaceChart)
	if elem == nil {
		return ""
	}
	leaf, ok := elem.(*openxml.LeafElementBase)
	if !ok {
		return ""
	}

	return leaf.InnerText()
}

// SetFormula sets the formula reference.
func (n *NumberReference) SetFormula(
	formula string,
) {
	existing := n.GetElement(
		elemF,
		NamespaceChart,
	)
	if existing != nil {
		n.RemoveChild(existing)
	}
	f := openxml.NewLeafElementWithText(
		NamespaceChart,
		elemF,
		PrefixChart,
		formula,
	)
	n.PrependChild(f)
}

// SetCache sets the cached numeric values for this reference using "General"
// format code.
func (n *NumberReference) SetCache(
	values []float64,
) {
	n.SetCacheWithFormat(values, "General")
}

// SetCacheWithFormat sets the cached numeric values for this reference with
// the specified format code.
func (n *NumberReference) SetCacheWithFormat(
	values []float64,
	formatCode string,
) {
	// Remove existing cache if present
	if existing := n.GetElement("numCache", NamespaceChart); existing != nil {
		n.RemoveChild(existing)
	}

	// Create new cache
	cache := NewNumberCache()
	cache.SetFormatCode(formatCode)
	cache.SetPointCount(uint32(len(values)))

	// Add all points
	for i, val := range values {
		cache.AddPoint(
			uint32(i),
			strconv.FormatFloat(
				val,
				'g',
				-1,
				floatBitSize64,
			),
		)
	}

	n.AppendChild(cache)
}

// Clone creates a deep copy of this NumberReference.
func (n *NumberReference) Clone() openxml.Element {
	cloneBase := n.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &NumberReference{
		CompositeElementBase: clone,
	}
}

// SeriesText represents the c:tx element for series text (name).
type SeriesText struct {
	*openxml.CompositeElementBase
}

// NewSeriesText creates a new series text element.
func NewSeriesText() *SeriesText {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"tx",
		PrefixChart,
	)

	return &SeriesText{CompositeElementBase: elem}
}

// NewSeriesTextWithValue creates a series text with a literal value.
func NewSeriesTextWithValue(
	text string,
) *SeriesText {
	st := NewSeriesText()
	v := openxml.NewLeafElementWithText(
		NamespaceChart,
		"v",
		PrefixChart,
		text,
	)
	st.AppendChild(v)

	return st
}

// NewSeriesTextWithReference creates a series text with a cell reference.
func NewSeriesTextWithReference(
	formula string,
) *SeriesText {
	st := NewSeriesText()
	strRef := NewStringReference(formula)
	st.AppendChild(strRef)

	return st
}

// Clone creates a deep copy of this SeriesText.
func (s *SeriesText) Clone() openxml.Element {
	cloneBase := s.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &SeriesText{
		CompositeElementBase: clone,
	}
}

// MultiLevelStringReference represents the c:multiLvlStrRef element
// for multi-level category axis data references.
type MultiLevelStringReference struct {
	*openxml.CompositeElementBase
}

// NewMultiLevelStringReference creates a new multi-level string reference
// with the given formula.
func NewMultiLevelStringReference(
	formula string,
) *MultiLevelStringReference {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"multiLvlStrRef",
		PrefixChart,
	)
	mlsr := &MultiLevelStringReference{
		CompositeElementBase: elem,
	}

	// Add formula
	f := openxml.NewLeafElementWithText(
		NamespaceChart,
		elemF,
		PrefixChart,
		formula,
	)
	mlsr.AppendChild(f)

	return mlsr
}

// Formula returns the formula reference.
func (m *MultiLevelStringReference) Formula() string {
	elem := m.GetElement(elemF, NamespaceChart)
	if elem == nil {
		return ""
	}
	leaf, ok := elem.(*openxml.LeafElementBase)
	if !ok {
		return ""
	}

	return leaf.InnerText()
}

// SetFormula sets the formula reference.
func (m *MultiLevelStringReference) SetFormula(
	formula string,
) {
	existing := m.GetElement(
		elemF,
		NamespaceChart,
	)
	if existing != nil {
		m.RemoveChild(existing)
	}
	f := openxml.NewLeafElementWithText(
		NamespaceChart,
		elemF,
		PrefixChart,
		formula,
	)
	m.PrependChild(f)
}

// Clone creates a deep copy of this MultiLevelStringReference.
func (m *MultiLevelStringReference) Clone() openxml.Element {
	cloneBase := m.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &MultiLevelStringReference{
		CompositeElementBase: clone,
	}
}

// NumberCache represents the c:numCache element for cached numeric data.
type NumberCache struct {
	*openxml.CompositeElementBase
}

// NewNumberCache creates a new number cache element.
func NewNumberCache() *NumberCache {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"numCache",
		PrefixChart,
	)

	return &NumberCache{
		CompositeElementBase: elem,
	}
}

// SetFormatCode sets the number format code.
func (n *NumberCache) SetFormatCode(code string) {
	if existing := n.GetElement("formatCode", NamespaceChart); existing != nil {
		n.RemoveChild(existing)
	}
	formatCode := openxml.NewLeafElementWithText(
		NamespaceChart,
		"formatCode",
		PrefixChart,
		code,
	)
	n.PrependChild(formatCode)
}

// SetPointCount sets the number of data points.
func (n *NumberCache) SetPointCount(
	count uint32,
) {
	if existing := n.GetElement(elemPtCount, NamespaceChart); existing != nil {
		n.RemoveChild(existing)
	}
	ptCount := openxml.NewLeafElement(
		NamespaceChart,
		elemPtCount,
		PrefixChart,
	)
	ptCount.SetAttribute(
		openxml.NewSimpleAttribute(
			"val",
			formatUint32(count),
		),
	)
	n.AppendChild(ptCount)
}

// AddPoint adds a point to the cache.
func (n *NumberCache) AddPoint(
	index uint32,
	value string,
) {
	pt := openxml.NewCompositeElement(
		NamespaceChart,
		"pt",
		PrefixChart,
	)
	pt.SetAttribute(
		openxml.NewSimpleAttribute(
			"idx",
			formatUint32(index),
		),
	)
	v := openxml.NewLeafElementWithText(
		NamespaceChart,
		"v",
		PrefixChart,
		value,
	)
	pt.AppendChild(v)
	n.AppendChild(pt)
}

// Clone creates a deep copy of this NumberCache.
func (n *NumberCache) Clone() openxml.Element {
	cloneBase := n.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &NumberCache{
		CompositeElementBase: clone,
	}
}

// StringCache represents the c:strCache element for cached string data.
type StringCache struct {
	*openxml.CompositeElementBase
}

// NewStringCache creates a new string cache element.
func NewStringCache() *StringCache {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"strCache",
		PrefixChart,
	)

	return &StringCache{
		CompositeElementBase: elem,
	}
}

// SetPointCount sets the number of data points.
func (s *StringCache) SetPointCount(
	count uint32,
) {
	if existing := s.GetElement(elemPtCount, NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	ptCount := openxml.NewLeafElement(
		NamespaceChart,
		elemPtCount,
		PrefixChart,
	)
	ptCount.SetAttribute(
		openxml.NewSimpleAttribute(
			"val",
			formatUint32(count),
		),
	)
	s.AppendChild(ptCount)
}

// AddPoint adds a point to the cache.
func (s *StringCache) AddPoint(
	index uint32,
	value string,
) {
	pt := openxml.NewCompositeElement(
		NamespaceChart,
		"pt",
		PrefixChart,
	)
	pt.SetAttribute(
		openxml.NewSimpleAttribute(
			"idx",
			formatUint32(index),
		),
	)
	v := openxml.NewLeafElementWithText(
		NamespaceChart,
		"v",
		PrefixChart,
		value,
	)
	pt.AppendChild(v)
	s.AppendChild(pt)
}

// Clone creates a deep copy of this StringCache.
func (s *StringCache) Clone() openxml.Element {
	cloneBase := s.CompositeElementBase.Clone()
	clone, _ := cloneBase.(*openxml.CompositeElementBase)

	return &StringCache{
		CompositeElementBase: clone,
	}
}

// formatUint32 formats a uint32 value as a string.
func formatUint32(val uint32) string {
	return strconv.FormatUint(
		uint64(val),
		dataBase10,
	)
}
