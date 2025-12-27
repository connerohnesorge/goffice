package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// PrintOptions represents the print options element (x:printOptions).
type PrintOptions struct {
	*openxml.CompositeElementBase
}

// NewPrintOptions creates a new PrintOptions element.
func NewPrintOptions() *PrintOptions {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"printOptions",
		PrefixDefault,
	)

	return &PrintOptions{
		CompositeElementBase: elem,
	}
}

// HorizontalCentered returns whether the page is centered horizontally
// when printed.
func (po *PrintOptions) HorizontalCentered() bool {
	attr, found := po.GetAttribute(
		"horizontalCentered",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHorizontalCentered sets whether the page is centered horizontally.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (po *PrintOptions) SetHorizontalCentered(
	value bool,
) {
	if value {
		po.SetAttribute(
			openxml.NewAttribute(
				"",
				"horizontalCentered",
				"",
				attrValueTrue,
			),
		)
	} else {
		po.RemoveAttribute("horizontalCentered", "")
	}
}

// VerticalCentered returns whether the page is centered vertically
// when printed.
func (po *PrintOptions) VerticalCentered() bool {
	attr, found := po.GetAttribute(
		"verticalCentered",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetVerticalCentered sets whether the page is centered vertically.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (po *PrintOptions) SetVerticalCentered(
	value bool,
) {
	if value {
		po.SetAttribute(
			openxml.NewAttribute(
				"",
				"verticalCentered",
				"",
				attrValueTrue,
			),
		)
	} else {
		po.RemoveAttribute("verticalCentered", "")
	}
}

// Headings returns whether row and column headings are printed.
func (po *PrintOptions) Headings() bool {
	attr, found := po.GetAttribute("headings", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHeadings sets whether row and column headings are printed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (po *PrintOptions) SetHeadings(value bool) {
	if value {
		po.SetAttribute(
			openxml.NewAttribute(
				"",
				"headings",
				"",
				attrValueTrue,
			),
		)
	} else {
		po.RemoveAttribute("headings", "")
	}
}

// GridLines returns whether grid lines are printed.
func (po *PrintOptions) GridLines() bool {
	attr, found := po.GetAttribute(
		"gridLines",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetGridLines sets whether grid lines are printed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (po *PrintOptions) SetGridLines(value bool) {
	if value {
		po.SetAttribute(
			openxml.NewAttribute(
				"",
				"gridLines",
				"",
				attrValueTrue,
			),
		)
	} else {
		po.RemoveAttribute("gridLines", "")
	}
}

// GridLinesSet returns whether the gridLines attribute is set.
func (po *PrintOptions) GridLinesSet() bool {
	attr, found := po.GetAttribute(
		"gridLinesSet",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetGridLinesSet sets whether the gridLines attribute is set.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (po *PrintOptions) SetGridLinesSet(
	value bool,
) {
	if value {
		po.RemoveAttribute("gridLinesSet", "")
	} else {
		po.SetAttribute(openxml.NewAttribute("", "gridLinesSet", "", attrValueFalse))
	}
}

// Clone creates a deep copy of this PrintOptions element.
func (po *PrintOptions) Clone() openxml.Element {
	cloned := po.CompositeElementBase.Clone()

	return &PrintOptions{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PrintOptions element.
func (po *PrintOptions) CloneNode(
	deep bool,
) openxml.Element {
	cloned := po.CompositeElementBase.CloneNode(
		deep,
	)

	return &PrintOptions{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
