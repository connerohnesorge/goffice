package elements

//revive:disable:file-length-limit many alignment properties

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Parse int constants.
const (
	parseIntBase10   = 10
	parseIntBitSize8 = 8
)

// HorizontalAlignment represents the horizontal alignment type.
type HorizontalAlignment string

const (
	// HorizontalAlignmentGeneral indicates general alignment (default).
	HorizontalAlignmentGeneral HorizontalAlignment = "general"
	// HorizontalAlignmentLeft indicates left alignment.
	HorizontalAlignmentLeft HorizontalAlignment = "left"
	// HorizontalAlignmentCenter indicates center alignment.
	HorizontalAlignmentCenter HorizontalAlignment = "center"
	// HorizontalAlignmentRight indicates right alignment.
	HorizontalAlignmentRight HorizontalAlignment = "right"
	// HorizontalAlignmentFill indicates fill alignment.
	HorizontalAlignmentFill HorizontalAlignment = "fill"
	// HorizontalAlignmentJustify indicates justified alignment.
	HorizontalAlignmentJustify HorizontalAlignment = "justify"
	// HorizontalAlignmentCenterContinuous indicates center-continuous
	// alignment.
	HorizontalAlignmentCenterContinuous HorizontalAlignment = "centerContinuous"
	// HorizontalAlignmentDistributed indicates distributed alignment.
	HorizontalAlignmentDistributed HorizontalAlignment = "distributed"
)

// VerticalAlignment represents the vertical alignment type.
type VerticalAlignment string

const (
	// VerticalAlignmentTop indicates top alignment.
	VerticalAlignmentTop VerticalAlignment = "top"
	// VerticalAlignmentCenter indicates center alignment.
	VerticalAlignmentCenter VerticalAlignment = "center"
	// VerticalAlignmentBottom indicates bottom alignment (default).
	VerticalAlignmentBottom VerticalAlignment = "bottom"
	// VerticalAlignmentJustify indicates justified alignment.
	VerticalAlignmentJustify VerticalAlignment = "justify"
	// VerticalAlignmentDistributed indicates distributed alignment.
	VerticalAlignmentDistributed VerticalAlignment = "distributed"
)

// ReadingOrder represents the reading order direction.
type ReadingOrder uint32

const (
	// ReadingOrderContextDependent indicates context-dependent reading order.
	ReadingOrderContextDependent ReadingOrder = 0
	// ReadingOrderLeftToRight indicates left-to-right reading order.
	ReadingOrderLeftToRight ReadingOrder = 1
	// ReadingOrderRightToLeft indicates right-to-left reading order.
	ReadingOrderRightToLeft ReadingOrder = 2
)

// Alignment represents the alignment element (x:alignment) in cell formatting.
type Alignment struct {
	*openxml.LeafElementBase
}

// NewAlignment creates a new Alignment element.
func NewAlignment() *Alignment {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"alignment",
		PrefixDefault,
	)

	return &Alignment{LeafElementBase: elem}
}

// Horizontal returns the horizontal alignment.
func (a *Alignment) Horizontal() HorizontalAlignment {
	attr, found := a.GetAttribute(
		"horizontal",
		"",
	)
	if !found {
		return HorizontalAlignmentGeneral
	}

	return HorizontalAlignment(attr.Value())
}

// SetHorizontal sets the horizontal alignment.
func (a *Alignment) SetHorizontal(
	align HorizontalAlignment,
) {
	if align == "" ||
		align == HorizontalAlignmentGeneral {
		a.RemoveAttribute("horizontal", "")

		return
	}
	a.SetAttribute(
		openxml.NewAttribute(
			"",
			"horizontal",
			"",
			string(align),
		),
	)
}

// Vertical returns the vertical alignment.
func (a *Alignment) Vertical() VerticalAlignment {
	attr, found := a.GetAttribute("vertical", "")
	if !found {
		return VerticalAlignmentBottom
	}

	return VerticalAlignment(attr.Value())
}

// SetVertical sets the vertical alignment.
func (a *Alignment) SetVertical(
	align VerticalAlignment,
) {
	if align == "" ||
		align == VerticalAlignmentBottom {
		a.RemoveAttribute("vertical", "")

		return
	}
	a.SetAttribute(
		openxml.NewAttribute(
			"",
			"vertical",
			"",
			string(align),
		),
	)
}

// TextRotation returns the text rotation angle (0-180 or 255 for vertical).
// 0-90 is up (0-90 degrees), 91-180 is down (degrees = value - 90).
// 255 indicates vertical text.
func (a *Alignment) TextRotation() uint32 {
	attr, found := a.GetAttribute(
		"textRotation",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseIntBase10,
		parseIntBitSize8,
	)

	return uint32(val)
}

// SetTextRotation sets the text rotation angle (0-180 or 255 for vertical).
func (a *Alignment) SetTextRotation(
	rotation uint32,
) {
	if rotation == 0 {
		a.RemoveAttribute("textRotation", "")

		return
	}
	a.SetAttribute(
		openxml.NewAttribute(
			"",
			"textRotation",
			"",
			strconv.FormatUint(
				uint64(rotation),
				parseIntBase10,
			),
		),
	)
}

// WrapText returns whether text wrapping is enabled.
func (a *Alignment) WrapText() bool {
	attr, found := a.GetAttribute("wrapText", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetWrapText sets whether text wrapping is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (a *Alignment) SetWrapText(wrap bool) {
	if wrap {
		a.SetAttribute(
			openxml.NewAttribute(
				"",
				"wrapText",
				"",
				attrValueTrue,
			),
		)
	} else {
		a.RemoveAttribute("wrapText", "")
	}
}

// Indent returns the indent level (number of spaces).
func (a *Alignment) Indent() uint32 {
	attr, found := a.GetAttribute("indent", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseIntBase10,
		parseIntBitSize8,
	)

	return uint32(val)
}

// SetIndent sets the indent level (number of spaces).
func (a *Alignment) SetIndent(indent uint32) {
	if indent == 0 {
		a.RemoveAttribute("indent", "")

		return
	}
	a.SetAttribute(
		openxml.NewAttribute(
			"",
			"indent",
			"",
			strconv.FormatUint(
				uint64(indent),
				parseIntBase10,
			),
		),
	)
}

// RelativeIndent returns the relative indent adjustment.
func (a *Alignment) RelativeIndent() int32 {
	attr, found := a.GetAttribute(
		"relativeIndent",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseIntBase10,
		parseIntBitSize8,
	)

	return int32(val)
}

// SetRelativeIndent sets the relative indent adjustment.
func (a *Alignment) SetRelativeIndent(
	indent int32,
) {
	if indent == 0 {
		a.RemoveAttribute("relativeIndent", "")

		return
	}
	a.SetAttribute(
		openxml.NewAttribute(
			"",
			"relativeIndent",
			"",
			strconv.FormatInt(
				int64(indent),
				parseIntBase10,
			),
		),
	)
}

// JustifyLastLine returns whether the last line is justified.
func (a *Alignment) JustifyLastLine() bool {
	attr, found := a.GetAttribute(
		"justifyLastLine",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetJustifyLastLine sets whether the last line is justified.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (a *Alignment) SetJustifyLastLine(
	justify bool,
) {
	if justify {
		a.SetAttribute(
			openxml.NewAttribute(
				"",
				"justifyLastLine",
				"",
				attrValueTrue,
			),
		)
	} else {
		a.RemoveAttribute("justifyLastLine", "")
	}
}

// ShrinkToFit returns whether text should shrink to fit the cell width.
func (a *Alignment) ShrinkToFit() bool {
	attr, found := a.GetAttribute(
		"shrinkToFit",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShrinkToFit sets whether text should shrink to fit the cell width.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (a *Alignment) SetShrinkToFit(shrink bool) {
	if shrink {
		a.SetAttribute(
			openxml.NewAttribute(
				"",
				"shrinkToFit",
				"",
				attrValueTrue,
			),
		)
	} else {
		a.RemoveAttribute("shrinkToFit", "")
	}
}

// ReadingOrder returns the reading order direction.
func (a *Alignment) ReadingOrder() ReadingOrder {
	attr, found := a.GetAttribute(
		"readingOrder",
		"",
	)
	if !found {
		return ReadingOrderContextDependent
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseIntBase10,
		parseIntBitSize8,
	)

	return ReadingOrder(val)
}

// SetReadingOrder sets the reading order direction.
func (a *Alignment) SetReadingOrder(
	order ReadingOrder,
) {
	if order == ReadingOrderContextDependent {
		a.RemoveAttribute("readingOrder", "")

		return
	}
	a.SetAttribute(
		openxml.NewAttribute(
			"",
			"readingOrder",
			"",
			strconv.FormatUint(
				uint64(order),
				parseIntBase10,
			),
		),
	)
}

// Clone creates a deep copy of this Alignment element.
func (a *Alignment) Clone() openxml.Element {
	cloned := a.LeafElementBase.Clone()

	return &Alignment{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Alignment element.
func (a *Alignment) CloneNode(
	deep bool,
) openxml.Element {
	cloned := a.LeafElementBase.CloneNode(deep)

	return &Alignment{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
