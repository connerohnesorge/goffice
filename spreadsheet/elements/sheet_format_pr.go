package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Default sheet format values.
const (
	defaultBaseColWidth     = 8
	defaultDefaultRowHeight = 15.0
)

// SheetFormatPr represents the sheet format properties element.
// XML element: x:sheetFormatPr.
type SheetFormatPr struct {
	*openxml.CompositeElementBase
}

// NewSheetFormatPr creates a new SheetFormatPr element.
func NewSheetFormatPr() *SheetFormatPr {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"sheetFormatPr",
		PrefixDefault,
	)

	return &SheetFormatPr{
		CompositeElementBase: elem,
	}
}

// BaseColWidth returns the base column width in characters.
func (sf *SheetFormatPr) BaseColWidth() int {
	attr, found := sf.GetAttribute(
		"baseColWidth",
		"",
	)
	if !found {
		return defaultBaseColWidth
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetBaseColWidth sets the base column width in characters.
func (sf *SheetFormatPr) SetBaseColWidth(
	width int,
) {
	if width == defaultBaseColWidth {
		sf.RemoveAttribute("baseColWidth", "")

		return
	}
	sf.SetAttribute(
		openxml.NewAttribute(
			"",
			"baseColWidth",
			"",
			strconv.Itoa(width),
		),
	)
}

// DefaultColWidth returns the default column width.
func (sf *SheetFormatPr) DefaultColWidth() float64 {
	attr, found := sf.GetAttribute(
		"defaultColWidth",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetDefaultColWidth sets the default column width.
func (sf *SheetFormatPr) SetDefaultColWidth(
	width float64,
) {
	if width == 0 {
		sf.RemoveAttribute("defaultColWidth", "")

		return
	}
	sf.SetAttribute(
		openxml.NewAttribute(
			"",
			"defaultColWidth",
			"",
			strconv.FormatFloat(
				width,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// DefaultRowHeight returns the default row height in points.
func (sf *SheetFormatPr) DefaultRowHeight() float64 {
	attr, found := sf.GetAttribute(
		"defaultRowHeight",
		"",
	)
	if !found {
		return defaultDefaultRowHeight
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetDefaultRowHeight sets the default row height in points.
func (sf *SheetFormatPr) SetDefaultRowHeight(
	height float64,
) {
	sf.SetAttribute(
		openxml.NewAttribute(
			"",
			"defaultRowHeight",
			"",
			strconv.FormatFloat(
				height,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// CustomHeight returns whether the default row height was manually set.
func (sf *SheetFormatPr) CustomHeight() bool {
	attr, found := sf.GetAttribute(
		"customHeight",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCustomHeight sets whether the default row height was manually set.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sf *SheetFormatPr) SetCustomHeight(
	value bool,
) {
	if value {
		sf.SetAttribute(
			openxml.NewAttribute(
				"",
				"customHeight",
				"",
				attrValueTrue,
			),
		)
	} else {
		sf.RemoveAttribute("customHeight", "")
	}
}

// ZeroHeight returns whether rows are hidden by default.
func (sf *SheetFormatPr) ZeroHeight() bool {
	attr, found := sf.GetAttribute(
		"zeroHeight",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetZeroHeight sets whether rows are hidden by default.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sf *SheetFormatPr) SetZeroHeight(
	value bool,
) {
	if value {
		sf.SetAttribute(
			openxml.NewAttribute(
				"",
				"zeroHeight",
				"",
				attrValueTrue,
			),
		)
	} else {
		sf.RemoveAttribute("zeroHeight", "")
	}
}

// ThickTop returns whether thick top border is applied.
func (sf *SheetFormatPr) ThickTop() bool {
	attr, found := sf.GetAttribute(
		"thickTop",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetThickTop sets whether thick top border is applied.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sf *SheetFormatPr) SetThickTop(
	value bool,
) {
	if value {
		sf.SetAttribute(
			openxml.NewAttribute(
				"",
				"thickTop",
				"",
				attrValueTrue,
			),
		)
	} else {
		sf.RemoveAttribute("thickTop", "")
	}
}

// ThickBottom returns whether thick bottom border is applied.
func (sf *SheetFormatPr) ThickBottom() bool {
	attr, found := sf.GetAttribute(
		"thickBottom",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetThickBottom sets whether thick bottom border is applied.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sf *SheetFormatPr) SetThickBottom(
	value bool,
) {
	if value {
		sf.SetAttribute(
			openxml.NewAttribute(
				"",
				"thickBottom",
				"",
				attrValueTrue,
			),
		)
	} else {
		sf.RemoveAttribute("thickBottom", "")
	}
}

// OutlineLevelRow returns the maximum outline level for rows.
func (sf *SheetFormatPr) OutlineLevelRow() int {
	attr, found := sf.GetAttribute(
		"outlineLevelRow",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetOutlineLevelRow sets the maximum outline level for rows.
func (sf *SheetFormatPr) SetOutlineLevelRow(
	level int,
) {
	if level == 0 {
		sf.RemoveAttribute("outlineLevelRow", "")

		return
	}
	sf.SetAttribute(
		openxml.NewAttribute(
			"",
			"outlineLevelRow",
			"",
			strconv.Itoa(level),
		),
	)
}

// OutlineLevelCol returns the maximum outline level for columns.
func (sf *SheetFormatPr) OutlineLevelCol() int {
	attr, found := sf.GetAttribute(
		"outlineLevelCol",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetOutlineLevelCol sets the maximum outline level for columns.
func (sf *SheetFormatPr) SetOutlineLevelCol(
	level int,
) {
	if level == 0 {
		sf.RemoveAttribute("outlineLevelCol", "")

		return
	}
	sf.SetAttribute(
		openxml.NewAttribute(
			"",
			"outlineLevelCol",
			"",
			strconv.Itoa(level),
		),
	)
}

// Clone creates a deep copy of this SheetFormatPr element.
func (sf *SheetFormatPr) Clone() openxml.Element {
	cloned := sf.CompositeElementBase.Clone()

	return &SheetFormatPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SheetFormatPr element.
func (sf *SheetFormatPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sf.CompositeElementBase.CloneNode(
		deep,
	)

	return &SheetFormatPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
