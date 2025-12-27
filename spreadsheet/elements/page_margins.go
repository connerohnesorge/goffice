package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Default page margin values (in inches).
const (
	defaultMarginLeft   = 0.7
	defaultMarginRight  = 0.7
	defaultMarginTop    = 0.75
	defaultMarginBottom = 0.75
	defaultMarginHeader = 0.3
	defaultMarginFooter = 0.3
)

// PageMargins represents the page margins element (x:pageMargins).
type PageMargins struct {
	*openxml.CompositeElementBase
}

// NewPageMargins creates a new PageMargins element.
func NewPageMargins() *PageMargins {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"pageMargins",
		PrefixDefault,
	)

	return &PageMargins{
		CompositeElementBase: elem,
	}
}

// Left returns the left margin in inches.
func (pm *PageMargins) Left() float64 {
	attr, found := pm.GetAttribute("left", "")
	if !found {
		return defaultMarginLeft
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetLeft sets the left margin in inches.
func (pm *PageMargins) SetLeft(margin float64) {
	pm.SetAttribute(
		openxml.NewAttribute(
			"",
			"left",
			"",
			strconv.FormatFloat(
				margin,
				'f',
				-1,
				64, //nolint:revive // add-constant: bit size 64
			),
		),
	)
}

// Right returns the right margin in inches.
func (pm *PageMargins) Right() float64 {
	attr, found := pm.GetAttribute("right", "")
	if !found {
		return defaultMarginRight
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetRight sets the right margin in inches.
func (pm *PageMargins) SetRight(margin float64) {
	pm.SetAttribute(
		openxml.NewAttribute(
			"",
			"right",
			"",
			strconv.FormatFloat(
				margin,
				'f',
				-1,
				64, //nolint:revive // add-constant: bit size 64
			),
		),
	)
}

// Top returns the top margin in inches.
func (pm *PageMargins) Top() float64 {
	attr, found := pm.GetAttribute("top", "")
	if !found {
		return defaultMarginTop
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetTop sets the top margin in inches.
func (pm *PageMargins) SetTop(margin float64) {
	pm.SetAttribute(
		openxml.NewAttribute(
			"",
			"top",
			"",
			strconv.FormatFloat(
				margin,
				'f',
				-1,
				64, //nolint:revive // add-constant: bit size 64
			),
		),
	)
}

// Bottom returns the bottom margin in inches.
func (pm *PageMargins) Bottom() float64 {
	attr, found := pm.GetAttribute("bottom", "")
	if !found {
		return defaultMarginBottom
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetBottom sets the bottom margin in inches.
func (pm *PageMargins) SetBottom(margin float64) {
	pm.SetAttribute(
		openxml.NewAttribute(
			"",
			"bottom",
			"",
			strconv.FormatFloat(
				margin,
				'f',
				-1,
				64, //nolint:revive // add-constant: bit size 64
			),
		),
	)
}

// Header returns the header margin in inches.
func (pm *PageMargins) Header() float64 {
	attr, found := pm.GetAttribute("header", "")
	if !found {
		return defaultMarginHeader
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetHeader sets the header margin in inches.
func (pm *PageMargins) SetHeader(margin float64) {
	pm.SetAttribute(
		openxml.NewAttribute(
			"",
			"header",
			"",
			strconv.FormatFloat(
				margin,
				'f',
				-1,
				64, //nolint:revive // add-constant: bit size 64
			),
		),
	)
}

// Footer returns the footer margin in inches.
func (pm *PageMargins) Footer() float64 {
	attr, found := pm.GetAttribute("footer", "")
	if !found {
		return defaultMarginFooter
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetFooter sets the footer margin in inches.
func (pm *PageMargins) SetFooter(margin float64) {
	pm.SetAttribute(
		openxml.NewAttribute(
			"",
			"footer",
			"",
			strconv.FormatFloat(
				margin,
				'f',
				-1,
				64, //nolint:revive // add-constant: bit size 64
			),
		),
	)
}

// SetAllMargins sets all page margins at once.
//
//nolint:revive // argument-limit
func (pm *PageMargins) SetAllMargins(
	left, right, top, bottom, header, footer float64,
) {
	pm.SetLeft(left)
	pm.SetRight(right)
	pm.SetTop(top)
	pm.SetBottom(bottom)
	pm.SetHeader(header)
	pm.SetFooter(footer)
}

// SetDefaults sets all margins to their default values.
func (pm *PageMargins) SetDefaults() {
	pm.SetAllMargins(
		defaultMarginLeft,
		defaultMarginRight,
		defaultMarginTop,
		defaultMarginBottom,
		defaultMarginHeader,
		defaultMarginFooter,
	)
}

// Clone creates a deep copy of this PageMargins element.
func (pm *PageMargins) Clone() openxml.Element {
	cloned := pm.CompositeElementBase.Clone()

	return &PageMargins{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PageMargins element.
func (pm *PageMargins) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pm.CompositeElementBase.CloneNode(
		deep,
	)

	return &PageMargins{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
