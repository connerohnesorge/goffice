package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Top10 represents the top10 filter element (x:top10).
// It filters to show only the top or bottom N items.
type Top10 struct {
	*openxml.LeafElementBase
}

// NewTop10 creates a new Top10 element.
func NewTop10() *Top10 {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"top10",
		PrefixDefault,
	)

	return &Top10{LeafElementBase: elem}
}

// NewTop10WithParams creates a new Top10 element with common parameters.
//
//nolint:revive // enforce-repeated-arg-type-style: separate params for clarity
func NewTop10WithParams(
	val float64,
	top bool,
	percent bool,
) *Top10 {
	t := NewTop10()
	t.SetVal(val)
	t.SetTop(top)
	t.SetPercent(percent)

	return t
}

// Top returns whether to show top (true) or bottom (false) items.
// Attribute: top. Default is true.
func (t *Top10) Top() bool {
	attr, found := t.GetAttribute("top", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetTop sets whether to show top (true) or bottom (false) items.
// Attribute: top.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Top10) SetTop(value bool) {
	if value {
		t.RemoveAttribute(
			"top",
			"",
		) // Default is true
	} else {
		t.SetAttribute(
			openxml.NewAttribute(
				"",
				"top",
				"",
				attrValueFalse,
			),
		)
	}
}

// Percent returns whether Val represents a percentage (true) or count (false).
// Attribute: percent. Default is false.
func (t *Top10) Percent() bool {
	attr, found := t.GetAttribute("percent", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPercent sets whether Val represents a percentage (true) or count (false).
// Attribute: percent.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Top10) SetPercent(value bool) {
	if value {
		t.SetAttribute(
			openxml.NewAttribute(
				"",
				"percent",
				"",
				attrValueTrue,
			),
		)
	} else {
		t.RemoveAttribute("percent", "")
	}
}

// Val returns the number of items or percentage to show. Attribute: val.
func (t *Top10) Val() float64 {
	attr, found := t.GetAttribute("val", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		bitSize64,
	)

	return val
}

// SetVal sets the number of items or percentage to show. Attribute: val.
func (t *Top10) SetVal(value float64) {
	if value == 0 {
		t.RemoveAttribute("val", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			strconv.FormatFloat(
				value,
				'f',
				-1,
				bitSize64,
			),
		),
	)
}

// FilterVal returns the actual filter value computed by the application.
// Attribute: filterVal. This is used when the source data changes.
func (t *Top10) FilterVal() float64 {
	attr, found := t.GetAttribute("filterVal", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		bitSize64,
	)

	return val
}

// SetFilterVal sets the actual filter value computed by the application.
// Attribute: filterVal.
func (t *Top10) SetFilterVal(value float64) {
	if value == 0 {
		t.RemoveAttribute("filterVal", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"filterVal",
			"",
			strconv.FormatFloat(
				value,
				'f',
				-1,
				bitSize64,
			),
		),
	)
}

// Clone creates a deep copy of this Top10 element.
func (t *Top10) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &Top10{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Top10 element.
func (t *Top10) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.LeafElementBase.CloneNode(deep)

	return &Top10{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
