//revive:disable:file-length-limit axis types need to be together

// Package drawingml provides shared DrawingML types for shapes, images,
// and effects.
//
// This file implements chart axis types.
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Numeric constants for strconv operations.
const (
	// axesBase10 is the base used for integer to string conversion.
	axesBase10 = 10
	// axesBitSize32 is the bit size for 32-bit integer parsing.
	axesBitSize32 = 32
	// axesBitSize64 is the bit size for 64-bit integer parsing.
	axesBitSize64 = 64
)

// Attribute name constants.
const (
	attrAxId    = "axId"
	attrAxPos   = "axPos"
	attrCrossAx = "crossAx"
)

// CategoryAxis represents the c:catAx element.
type CategoryAxis struct {
	*openxml.CompositeElementBase
}

// NewCategoryAxis creates a new category axis with the given IDs.
func NewCategoryAxis(
	axisID, crossAxisID uint32,
) *CategoryAxis {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"catAx",
		PrefixChart,
	)
	ax := &CategoryAxis{
		CompositeElementBase: elem,
	}

	// Add axis ID
	axId := openxml.NewLeafElement(
		NamespaceChart,
		attrAxId,
		PrefixChart,
	)
	axId.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				axesBase10,
			),
		),
	)
	ax.AppendChild(axId)

	// Add scaling
	scaling := NewScaling()
	ax.AppendChild(scaling)

	// Add axis position (default bottom)
	axPos := openxml.NewLeafElement(
		NamespaceChart,
		attrAxPos,
		PrefixChart,
	)
	axPos.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(AxisPositionBottom),
		),
	)
	ax.AppendChild(axPos)

	// Add cross axis ID
	crossAx := openxml.NewLeafElement(
		NamespaceChart,
		attrCrossAx,
		PrefixChart,
	)
	crossAx.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(crossAxisID),
				axesBase10,
			),
		),
	)
	ax.AppendChild(crossAx)

	return ax
}

// AxisID returns the axis ID.
func (ax *CategoryAxis) AxisID() uint32 {
	elem := ax.GetElement(
		attrAxId,
		NamespaceChart,
	)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		axesBase10,
		axesBitSize32,
	)

	return uint32(val)
}

// SetAxisPosition sets the axis position.
func (ax *CategoryAxis) SetAxisPosition(
	pos AxisPositionValue,
) {
	elem := ax.GetElement(
		attrAxPos,
		NamespaceChart,
	)
	if elem == nil {
		return
	}
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(pos),
		),
	)
}

// SetMajorTickMark sets the major tick mark style.
func (ax *CategoryAxis) SetMajorTickMark(
	mark TickMarkValue,
) {
	existing := ax.GetElement(
		"majorTickMark",
		NamespaceChart,
	)
	if existing != nil {
		ax.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"majorTickMark",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(mark),
		),
	)
	ax.AppendChild(elem)
}

// SetMinorTickMark sets the minor tick mark style.
func (ax *CategoryAxis) SetMinorTickMark(
	mark TickMarkValue,
) {
	existing := ax.GetElement(
		"minorTickMark",
		NamespaceChart,
	)
	if existing != nil {
		ax.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"minorTickMark",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(mark),
		),
	)
	ax.AppendChild(elem)
}

// SetTickLabelPosition sets the tick label position.
func (ax *CategoryAxis) SetTickLabelPosition(
	pos TickLabelPositionValue,
) {
	existing := ax.GetElement(
		"tickLblPos",
		NamespaceChart,
	)
	if existing != nil {
		ax.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"tickLblPos",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(pos),
		),
	)
	ax.AppendChild(elem)
}

// SetTitle sets the axis title.
func (ax *CategoryAxis) SetTitle(title *Title) {
	if existing := ax.GetElement("title", NamespaceChart); existing != nil {
		ax.RemoveChild(existing)
	}
	if title != nil {
		ax.AppendChild(title)
	}
}

// Clone creates a deep copy of this CategoryAxis.
func (ax *CategoryAxis) Clone() openxml.Element {
	clonedBase := ax.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &CategoryAxis{
		CompositeElementBase: cloned,
	}
}

// ValueAxis represents the c:valAx element.
type ValueAxis struct {
	*openxml.CompositeElementBase
}

// NewValueAxis creates a new value axis with the given IDs.
func NewValueAxis(
	axisID, crossAxisID uint32,
) *ValueAxis {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"valAx",
		PrefixChart,
	)
	ax := &ValueAxis{CompositeElementBase: elem}

	// Add axis ID
	axId := openxml.NewLeafElement(
		NamespaceChart,
		attrAxId,
		PrefixChart,
	)
	axId.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				axesBase10,
			),
		),
	)
	ax.AppendChild(axId)

	// Add scaling
	scaling := NewScaling()
	ax.AppendChild(scaling)

	// Add axis position (default left)
	axPos := openxml.NewLeafElement(
		NamespaceChart,
		attrAxPos,
		PrefixChart,
	)
	axPos.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(AxisPositionLeft),
		),
	)
	ax.AppendChild(axPos)

	// Add cross axis ID
	crossAx := openxml.NewLeafElement(
		NamespaceChart,
		attrCrossAx,
		PrefixChart,
	)
	crossAx.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(crossAxisID),
				axesBase10,
			),
		),
	)
	ax.AppendChild(crossAx)

	return ax
}

// AxisID returns the axis ID.
func (ax *ValueAxis) AxisID() uint32 {
	elem := ax.GetElement(
		attrAxId,
		NamespaceChart,
	)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		axesBase10,
		axesBitSize32,
	)

	return uint32(val)
}

// SetAxisPosition sets the axis position.
func (ax *ValueAxis) SetAxisPosition(
	pos AxisPositionValue,
) {
	elem := ax.GetElement(
		attrAxPos,
		NamespaceChart,
	)
	if elem == nil {
		return
	}
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(pos),
		),
	)
}

// SetMajorGridlines enables major gridlines.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ax *ValueAxis) SetMajorGridlines(
	enabled bool,
) {
	existing := ax.GetElement(
		"majorGridlines",
		NamespaceChart,
	)
	if existing != nil {
		ax.RemoveChild(existing)
	}
	if enabled {
		gridlines := openxml.NewCompositeElement(
			NamespaceChart,
			"majorGridlines",
			PrefixChart,
		)
		ax.AppendChild(gridlines)
	}
}

// SetMinorGridlines enables minor gridlines.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ax *ValueAxis) SetMinorGridlines(
	enabled bool,
) {
	existing := ax.GetElement(
		"minorGridlines",
		NamespaceChart,
	)
	if existing != nil {
		ax.RemoveChild(existing)
	}
	if enabled {
		gridlines := openxml.NewCompositeElement(
			NamespaceChart,
			"minorGridlines",
			PrefixChart,
		)
		ax.AppendChild(gridlines)
	}
}

// SetTitle sets the axis title.
func (ax *ValueAxis) SetTitle(title *Title) {
	if existing := ax.GetElement("title", NamespaceChart); existing != nil {
		ax.RemoveChild(existing)
	}
	if title != nil {
		ax.AppendChild(title)
	}
}

// SetNumberFormat sets the number format for the axis.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ax *ValueAxis) SetNumberFormat(
	formatCode string,
	sourceLinked bool,
) {
	if existing := ax.GetElement("numFmt", NamespaceChart); existing != nil {
		ax.RemoveChild(existing)
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
	ax.AppendChild(numFmt)
}

// Clone creates a deep copy of this ValueAxis.
func (ax *ValueAxis) Clone() openxml.Element {
	clonedBase := ax.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &ValueAxis{
		CompositeElementBase: cloned,
	}
}

// DateAxis represents the c:dateAx element.
type DateAxis struct {
	*openxml.CompositeElementBase
}

// NewDateAxis creates a new date axis with the given IDs.
func NewDateAxis(
	axisID, crossAxisID uint32,
) *DateAxis {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"dateAx",
		PrefixChart,
	)
	ax := &DateAxis{CompositeElementBase: elem}

	// Add axis ID
	axId := openxml.NewLeafElement(
		NamespaceChart,
		attrAxId,
		PrefixChart,
	)
	axId.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				axesBase10,
			),
		),
	)
	ax.AppendChild(axId)

	// Add scaling
	scaling := NewScaling()
	ax.AppendChild(scaling)

	// Add axis position (default bottom)
	axPos := openxml.NewLeafElement(
		NamespaceChart,
		attrAxPos,
		PrefixChart,
	)
	axPos.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(AxisPositionBottom),
		),
	)
	ax.AppendChild(axPos)

	// Add cross axis ID
	crossAx := openxml.NewLeafElement(
		NamespaceChart,
		attrCrossAx,
		PrefixChart,
	)
	crossAx.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(crossAxisID),
				axesBase10,
			),
		),
	)
	ax.AppendChild(crossAx)

	return ax
}

// AxisID returns the axis ID.
func (ax *DateAxis) AxisID() uint32 {
	elem := ax.GetElement(
		attrAxId,
		NamespaceChart,
	)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		axesBase10,
		axesBitSize32,
	)

	return uint32(val)
}

// Clone creates a deep copy of this DateAxis.
func (ax *DateAxis) Clone() openxml.Element {
	clonedBase := ax.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &DateAxis{CompositeElementBase: cloned}
}

// SeriesAxis represents the c:serAx element.
type SeriesAxis struct {
	*openxml.CompositeElementBase
}

// NewSeriesAxis creates a new series axis with the given IDs.
func NewSeriesAxis(
	axisID, crossAxisID uint32,
) *SeriesAxis {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"serAx",
		PrefixChart,
	)
	ax := &SeriesAxis{CompositeElementBase: elem}

	// Add axis ID
	axId := openxml.NewLeafElement(
		NamespaceChart,
		attrAxId,
		PrefixChart,
	)
	axId.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(axisID),
				axesBase10,
			),
		),
	)
	ax.AppendChild(axId)

	// Add scaling
	scaling := NewScaling()
	ax.AppendChild(scaling)

	// Add axis position
	axPos := openxml.NewLeafElement(
		NamespaceChart,
		attrAxPos,
		PrefixChart,
	)
	axPos.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(AxisPositionBottom),
		),
	)
	ax.AppendChild(axPos)

	// Add cross axis ID
	crossAx := openxml.NewLeafElement(
		NamespaceChart,
		attrCrossAx,
		PrefixChart,
	)
	crossAx.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatUint(
				uint64(crossAxisID),
				axesBase10,
			),
		),
	)
	ax.AppendChild(crossAx)

	return ax
}

// AxisID returns the axis ID.
func (ax *SeriesAxis) AxisID() uint32 {
	elem := ax.GetElement(
		attrAxId,
		NamespaceChart,
	)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(attrVal, "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		axesBase10,
		axesBitSize32,
	)

	return uint32(val)
}

// Clone creates a deep copy of this SeriesAxis.
func (ax *SeriesAxis) Clone() openxml.Element {
	clonedBase := ax.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &SeriesAxis{
		CompositeElementBase: cloned,
	}
}

// Scaling represents the c:scaling element for axis scaling.
type Scaling struct {
	*openxml.CompositeElementBase
}

// NewScaling creates a new scaling element.
func NewScaling() *Scaling {
	elem := openxml.NewCompositeElement(
		NamespaceChart,
		"scaling",
		PrefixChart,
	)
	s := &Scaling{CompositeElementBase: elem}

	// Add default orientation (minMax)
	orient := openxml.NewLeafElement(
		NamespaceChart,
		"orientation",
		PrefixChart,
	)
	orient.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(OrientationMinMax),
		),
	)
	s.AppendChild(orient)

	return s
}

// SetOrientation sets the axis orientation.
func (s *Scaling) SetOrientation(
	orientation OrientationValue,
) {
	elem := s.GetElement(
		"orientation",
		NamespaceChart,
	)
	if elem == nil {
		return
	}
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			string(orientation),
		),
	)
}

// SetMinimum sets the minimum value for the axis.
func (s *Scaling) SetMinimum(val float64) {
	if existing := s.GetElement("min", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	minElem := openxml.NewLeafElement(
		NamespaceChart,
		"min",
		PrefixChart,
	)
	minElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatFloat(
				val,
				'f',
				-1,
				axesBitSize64,
			),
		),
	)
	s.AppendChild(minElem)
}

// SetMaximum sets the maximum value for the axis.
func (s *Scaling) SetMaximum(val float64) {
	if existing := s.GetElement("max", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	maxElem := openxml.NewLeafElement(
		NamespaceChart,
		"max",
		PrefixChart,
	)
	maxElem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatFloat(
				val,
				'f',
				-1,
				axesBitSize64,
			),
		),
	)
	s.AppendChild(maxElem)
}

// SetLogBase sets the logarithmic base for the axis scaling.
// Typical values are 10 or 2. Range 2-1000.
func (s *Scaling) SetLogBase(val float64) {
	if existing := s.GetElement("logBase", NamespaceChart); existing != nil {
		s.RemoveChild(existing)
	}
	elem := openxml.NewLeafElement(
		NamespaceChart,
		"logBase",
		PrefixChart,
	)
	elem.SetAttribute(
		openxml.NewSimpleAttribute(
			attrVal,
			strconv.FormatFloat(
				val,
				'f',
				-1,
				axesBitSize64,
			),
		),
	)
	s.AppendChild(elem)
}

// Clone creates a deep copy of this Scaling.
func (s *Scaling) Clone() openxml.Element {
	clonedBase := s.CompositeElementBase.Clone()
	cloned, ok := clonedBase.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return &Scaling{CompositeElementBase: cloned}
}
