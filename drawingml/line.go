// and effects.
//
//nolint:revive // This file contains many public types for OOXML line elements.
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// LineDashValue represents preset line dash styles.
type LineDashValue string

// Line dash values defined in DrawingML.
const (
	LineDashSolid            LineDashValue = "solid"
	LineDashDot              LineDashValue = "dot"
	LineDashDash             LineDashValue = "dash"
	LineDashLongDash         LineDashValue = "lgDash"
	LineDashDashDot          LineDashValue = "dashDot"
	LineDashLongDashDot      LineDashValue = "lgDashDot"
	LineDashLongDashDotDot   LineDashValue = "lgDashDotDot"
	LineDashSystemDash       LineDashValue = "sysDash"
	LineDashSystemDot        LineDashValue = "sysDot"
	LineDashSystemDashDot    LineDashValue = "sysDashDot"
	LineDashSystemDashDotDot LineDashValue = "sysDashDotDot"
)

// LineCapValue represents line cap (end) styles.
type LineCapValue string

// Line cap values.
const (
	LineCapFlat   LineCapValue = "flat"
	LineCapSquare LineCapValue = "sq"
	LineCapRound  LineCapValue = "rnd"
)

// CompoundLineValue represents compound line types.
type CompoundLineValue string

// Compound line values.
const (
	CompoundLineSingle    CompoundLineValue = "sng"
	CompoundLineDouble    CompoundLineValue = "dbl"
	CompoundLineThickThin CompoundLineValue = "thickThin"
	CompoundLineThinThick CompoundLineValue = "thinThick"
	CompoundLineTriple    CompoundLineValue = "tri"
)

// PenAlignmentValue represents pen alignment for lines.
type PenAlignmentValue string

// Pen alignment values.
const (
	PenAlignmentCenter PenAlignmentValue = "ctr"
	PenAlignmentInset  PenAlignmentValue = "in"
)

// LineEndValue represents line end (arrow) types.
type LineEndValue string

// Line end type values.
const (
	LineEndNone     LineEndValue = "none"
	LineEndTriangle LineEndValue = "triangle"
	LineEndStealth  LineEndValue = "stealth"
	LineEndDiamond  LineEndValue = "diamond"
	LineEndOval     LineEndValue = "oval"
	LineEndArrow    LineEndValue = "arrow"
)

// LineEndSizeValue represents line end size values.
type LineEndSizeValue string

// Line end size values.
const (
	LineEndSizeSmall  LineEndSizeValue = "sm"
	LineEndSizeMedium LineEndSizeValue = "med"
	LineEndSizeLarge  LineEndSizeValue = "lg"
)

// LineJoinValue represents line join types.
type LineJoinValue string

// Line join values.
const (
	LineJoinRound LineJoinValue = "round"
	LineJoinBevel LineJoinValue = "bevel"
	LineJoinMiter LineJoinValue = "miter"
)

// LineEndProperties represents line end properties (a:headEnd or a:tailEnd).
type LineEndProperties struct {
	*openxml.LeafElementBase
}

// NewHeadEnd creates a new line head end properties element.
func NewHeadEnd() *LineEndProperties {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		"headEnd",
		PrefixMain,
	)

	return &LineEndProperties{
		LeafElementBase: elem,
	}
}

// NewTailEnd creates a new line tail end properties element.
func NewTailEnd() *LineEndProperties {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		"tailEnd",
		PrefixMain,
	)

	return &LineEndProperties{
		LeafElementBase: elem,
	}
}

// NewHeadEndWithArrow creates a head end with an arrow.
func NewHeadEndWithArrow(
	endType LineEndValue,
	width, length LineEndSizeValue,
) *LineEndProperties {
	le := NewHeadEnd()
	le.SetType(endType)
	le.SetWidth(width)
	le.SetLength(length)

	return le
}

// NewTailEndWithArrow creates a tail end with an arrow.
func NewTailEndWithArrow(
	endType LineEndValue,
	width, length LineEndSizeValue,
) *LineEndProperties {
	le := NewTailEnd()
	le.SetType(endType)
	le.SetWidth(width)
	le.SetLength(length)

	return le
}

// Type returns the line end type.
func (e *LineEndProperties) Type() LineEndValue {
	attr, found := e.GetAttribute("type", "")
	if !found {
		return LineEndNone
	}

	return LineEndValue(attr.Value())
}

// SetType sets the line end type.
func (e *LineEndProperties) SetType(
	endType LineEndValue,
) {
	if endType == "" || endType == LineEndNone {
		e.RemoveAttribute("type", "")

		return
	}
	e.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			string(endType),
		),
	)
}

// Width returns the line end width.
func (e *LineEndProperties) Width() LineEndSizeValue {
	attr, found := e.GetAttribute("w", "")
	if !found {
		return LineEndSizeMedium
	}

	return LineEndSizeValue(attr.Value())
}

// SetWidth sets the line end width.
func (e *LineEndProperties) SetWidth(
	width LineEndSizeValue,
) {
	if width == "" || width == LineEndSizeMedium {
		e.RemoveAttribute("w", "")

		return
	}
	e.SetAttribute(
		openxml.NewAttribute(
			"",
			"w",
			"",
			string(width),
		),
	)
}

// Length returns the line end length.
func (e *LineEndProperties) Length() LineEndSizeValue {
	attr, found := e.GetAttribute("len", "")
	if !found {
		return LineEndSizeMedium
	}

	return LineEndSizeValue(attr.Value())
}

// SetLength sets the line end length.
func (e *LineEndProperties) SetLength(
	length LineEndSizeValue,
) {
	if length == "" ||
		length == LineEndSizeMedium {
		e.RemoveAttribute("len", "")

		return
	}
	e.SetAttribute(
		openxml.NewAttribute(
			"",
			"len",
			"",
			string(length),
		),
	)
}

// Clone creates a deep copy of this LineEndProperties element.
func (e *LineEndProperties) Clone() openxml.Element {
	return &LineEndProperties{
		LeafElementBase: e.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// LineProperties represents line/outline properties (a:ln).
// This element specifies the outline style for a shape.
type LineProperties struct {
	*openxml.CompositeElementBase
}

// NewLineProperties creates a new line properties element.
func NewLineProperties() *LineProperties {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"ln",
		PrefixMain,
	)

	return &LineProperties{
		CompositeElementBase: elem,
	}
}

// NewLinePropertiesWithWidth creates a new line properties with the specified width.
// width is in EMUs.
func NewLinePropertiesWithWidth(
	width EMU,
) *LineProperties {
	lp := NewLineProperties()
	lp.SetWidth(width)

	return lp
}

// Width returns the line width in EMUs.
func (l *LineProperties) Width() EMU {
	attr, found := l.GetAttribute("w", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetWidth sets the line width in EMUs.
// The maximum value is 20116800 EMUs (approximately 558 points).
func (l *LineProperties) SetWidth(width EMU) {
	if width <= 0 {
		l.RemoveAttribute("w", "")

		return
	}
	l.SetAttribute(
		openxml.NewAttribute(
			"",
			"w",
			"",
			strconv.FormatInt(width.Int64(), 10),
		),
	)
}

// WidthPoints returns the line width in points.
func (l *LineProperties) WidthPoints() float64 {
	return EmuToPoints(l.Width())
}

// SetWidthPoints sets the line width in points.
func (l *LineProperties) SetWidthPoints(
	points float64,
) {
	l.SetWidth(PointsToEmu(points))
}

// CapType returns the line cap type.
func (l *LineProperties) CapType() LineCapValue {
	attr, found := l.GetAttribute("cap", "")
	if !found {
		return ""
	}

	return LineCapValue(attr.Value())
}

// SetCapType sets the line cap type.
func (l *LineProperties) SetCapType(
	capType LineCapValue,
) {
	if capType == "" {
		l.RemoveAttribute("cap", "")

		return
	}
	l.SetAttribute(
		openxml.NewAttribute(
			"",
			"cap",
			"",
			string(capType),
		),
	)
}

// CompoundLineType returns the compound line type.
func (l *LineProperties) CompoundLineType() CompoundLineValue {
	attr, found := l.GetAttribute("cmpd", "")
	if !found {
		return CompoundLineSingle
	}

	return CompoundLineValue(attr.Value())
}

// SetCompoundLineType sets the compound line type.
func (l *LineProperties) SetCompoundLineType(
	compound CompoundLineValue,
) {
	if compound == "" ||
		compound == CompoundLineSingle {
		l.RemoveAttribute("cmpd", "")

		return
	}
	l.SetAttribute(
		openxml.NewAttribute(
			"",
			"cmpd",
			"",
			string(compound),
		),
	)
}

// Alignment returns the pen alignment.
func (l *LineProperties) Alignment() PenAlignmentValue {
	attr, found := l.GetAttribute("algn", "")
	if !found {
		return PenAlignmentCenter
	}

	return PenAlignmentValue(attr.Value())
}

// SetAlignment sets the pen alignment.
func (l *LineProperties) SetAlignment(
	alignment PenAlignmentValue,
) {
	if alignment == "" ||
		alignment == PenAlignmentCenter {
		l.RemoveAttribute("algn", "")

		return
	}
	l.SetAttribute(
		openxml.NewAttribute(
			"",
			"algn",
			"",
			string(alignment),
		),
	)
}

// SetNoFill sets the line to have no fill (invisible).
func (l *LineProperties) SetNoFill() {
	l.removeFill()
	noFill := NewNoFill()
	l.PrependChild(noFill)
}

// SetSolidFill sets the line to have a solid color fill.
func (l *LineProperties) SetSolidFill(
	hexColor string,
) {
	l.removeFill()
	solidFill := NewSolidFillWithRgb(hexColor)
	l.PrependChild(solidFill)
}

// SetSolidFillSchemeColor sets the line to have a solid scheme color fill.
func (l *LineProperties) SetSolidFillSchemeColor(
	color SchemeColorValue,
) {
	l.removeFill()
	solidFill := NewSolidFillWithSchemeColor(
		color,
	)
	l.PrependChild(solidFill)
}

// SetGradientFill sets the line to have a gradient fill.
func (l *LineProperties) SetGradientFill(
	fill *GradientFill,
) {
	l.removeFill()
	l.PrependChild(fill)
}

// SetPatternFill sets the line to have a pattern fill.
func (l *LineProperties) SetPatternFill(
	fill *PatternFill,
) {
	l.removeFill()
	l.PrependChild(fill)
}

// removeFill removes any existing fill element.
func (l *LineProperties) removeFill() {
	if noFill := l.GetElement("noFill", NamespaceMain); noFill != nil {
		l.RemoveChild(noFill)
	}
	if solidFill := l.GetElement("solidFill", NamespaceMain); solidFill != nil {
		l.RemoveChild(solidFill)
	}
	if gradFill := l.GetElement("gradFill", NamespaceMain); gradFill != nil {
		l.RemoveChild(gradFill)
	}
	if pattFill := l.GetElement("pattFill", NamespaceMain); pattFill != nil {
		l.RemoveChild(pattFill)
	}
}

// SolidFill returns the solid fill if set, or nil.
func (l *LineProperties) SolidFill() *SolidFill {
	elem := l.GetElement(
		"solidFill",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if sf, ok := elem.(*SolidFill); ok {
		return sf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SolidFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetPresetDash sets the line dash style to a preset value.
func (l *LineProperties) SetPresetDash(
	dash LineDashValue,
) {
	l.removeDash()
	if dash == "" || dash == LineDashSolid {
		return
	}
	prstDash := openxml.NewLeafElement(
		NamespaceMain,
		"prstDash",
		PrefixMain,
	)
	prstDash.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			string(dash),
		),
	)
	l.AppendChild(prstDash)
}

// PresetDash returns the preset dash style.
func (l *LineProperties) PresetDash() LineDashValue {
	prstDash := l.GetElement(
		"prstDash",
		NamespaceMain,
	)
	if prstDash == nil {
		return LineDashSolid
	}
	attr, found := prstDash.GetAttribute(
		"val",
		"",
	)
	if !found {
		return LineDashSolid
	}

	return LineDashValue(attr.Value())
}

// removeDash removes any existing dash element.
func (l *LineProperties) removeDash() {
	if prstDash := l.GetElement("prstDash", NamespaceMain); prstDash != nil {
		l.RemoveChild(prstDash)
	}
	if custDash := l.GetElement("custDash", NamespaceMain); custDash != nil {
		l.RemoveChild(custDash)
	}
}

// SetRoundJoin sets the line join to round.
func (l *LineProperties) SetRoundJoin() {
	l.removeJoin()
	round := openxml.NewLeafElement(
		NamespaceMain,
		"round",
		PrefixMain,
	)
	l.AppendChild(round)
}

// SetBevelJoin sets the line join to bevel.
func (l *LineProperties) SetBevelJoin() {
	l.removeJoin()
	bevel := openxml.NewLeafElement(
		NamespaceMain,
		"bevel",
		PrefixMain,
	)
	l.AppendChild(bevel)
}

// SetMiterJoin sets the line join to miter with the specified limit.
// limit is the miter limit in 1000ths of a percent.
func (l *LineProperties) SetMiterJoin(limit int) {
	l.removeJoin()
	miter := openxml.NewLeafElement(
		NamespaceMain,
		"miter",
		PrefixMain,
	)
	if limit > 0 {
		miter.SetAttribute(
			openxml.NewAttribute(
				"",
				"lim",
				"",
				strconv.Itoa(limit),
			),
		)
	}
	l.AppendChild(miter)
}

// LineJoin returns the current line join type.
func (l *LineProperties) LineJoin() LineJoinValue {
	if l.GetElement(
		"round",
		NamespaceMain,
	) != nil {
		return LineJoinRound
	}
	if l.GetElement(
		"bevel",
		NamespaceMain,
	) != nil {
		return LineJoinBevel
	}
	if l.GetElement(
		"miter",
		NamespaceMain,
	) != nil {
		return LineJoinMiter
	}

	return "" // No join specified
}

// removeJoin removes any existing join element.
func (l *LineProperties) removeJoin() {
	if round := l.GetElement("round", NamespaceMain); round != nil {
		l.RemoveChild(round)
	}
	if bevel := l.GetElement("bevel", NamespaceMain); bevel != nil {
		l.RemoveChild(bevel)
	}
	if miter := l.GetElement("miter", NamespaceMain); miter != nil {
		l.RemoveChild(miter)
	}
}

// SetHeadEnd sets the head (start) end of the line.
func (l *LineProperties) SetHeadEnd(
	head *LineEndProperties,
) {
	// Remove existing head end
	if existing := l.GetElement("headEnd", NamespaceMain); existing != nil {
		l.RemoveChild(existing)
	}
	if head != nil {
		l.AppendChild(head)
	}
}

// HeadEnd returns the head end properties, or nil if not set.
func (l *LineProperties) HeadEnd() *LineEndProperties {
	elem := l.GetElement("headEnd", NamespaceMain)
	if elem == nil {
		return nil
	}
	if he, ok := elem.(*LineEndProperties); ok {
		return he
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &LineEndProperties{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// SetTailEnd sets the tail (end) end of the line.
func (l *LineProperties) SetTailEnd(
	tail *LineEndProperties,
) {
	// Remove existing tail end
	if existing := l.GetElement("tailEnd", NamespaceMain); existing != nil {
		l.RemoveChild(existing)
	}
	if tail != nil {
		l.AppendChild(tail)
	}
}

// TailEnd returns the tail end properties, or nil if not set.
func (l *LineProperties) TailEnd() *LineEndProperties {
	elem := l.GetElement("tailEnd", NamespaceMain)
	if elem == nil {
		return nil
	}
	if te, ok := elem.(*LineEndProperties); ok {
		return te
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &LineEndProperties{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// SetArrow sets both head and tail to arrow style.
func (l *LineProperties) SetArrow(
	headType, tailType LineEndValue,
) {
	if headType != LineEndNone && headType != "" {
		head := NewHeadEndWithArrow(
			headType,
			LineEndSizeMedium,
			LineEndSizeMedium,
		)
		l.SetHeadEnd(head)
	}
	if tailType != LineEndNone && tailType != "" {
		tail := NewTailEndWithArrow(
			tailType,
			LineEndSizeMedium,
			LineEndSizeMedium,
		)
		l.SetTailEnd(tail)
	}
}

// Clone creates a deep copy of this LineProperties element.
func (l *LineProperties) Clone() openxml.Element {
	return &LineProperties{
		CompositeElementBase: l.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
