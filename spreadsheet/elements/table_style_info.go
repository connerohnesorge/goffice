package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// TableStyleInfo represents the table style info element (x:tableStyleInfo).
// It specifies the table style and which formatting elements are shown.
type TableStyleInfo struct {
	*openxml.LeafElementBase
}

// NewTableStyleInfo creates a new TableStyleInfo element.
func NewTableStyleInfo() *TableStyleInfo {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"tableStyleInfo",
		PrefixDefault,
	)

	return &TableStyleInfo{LeafElementBase: elem}
}

// NewTableStyleInfoWithDefaults creates a TableStyleInfo with common defaults.
// Uses TableStyleMedium2 as the default style.
func NewTableStyleInfoWithDefaults() *TableStyleInfo {
	ts := NewTableStyleInfo()
	ts.SetName(TableStyleMedium2)
	ts.SetShowRowStripes(true)

	return ts
}

// Name returns the table style name. Attribute: name.
// This should be either a built-in style (e.g., "TableStyleMedium2") or
// a custom table style defined in the stylesheet.
func (ts *TableStyleInfo) Name() string {
	attr, found := ts.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the table style name. Attribute: name.
func (ts *TableStyleInfo) SetName(name string) {
	if name == "" {
		ts.RemoveAttribute("name", "")

		return
	}
	ts.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// ShowFirstColumn returns whether special formatting is applied to the first
// column. Attribute: showFirstColumn.
func (ts *TableStyleInfo) ShowFirstColumn() bool {
	attr, found := ts.GetAttribute(
		"showFirstColumn",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowFirstColumn sets whether special formatting is applied to the first column.
// Attribute: showFirstColumn.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ts *TableStyleInfo) SetShowFirstColumn(
	value bool,
) {
	if value {
		ts.SetAttribute(
			openxml.NewAttribute(
				"",
				"showFirstColumn",
				"",
				attrValueOne,
			),
		)
	} else {
		ts.SetAttribute(
			openxml.NewAttribute("", "showFirstColumn", "", attrValueZero),
		)
	}
}

// ShowLastColumn returns whether special formatting is applied to the last
// column. Attribute: showLastColumn.
func (ts *TableStyleInfo) ShowLastColumn() bool {
	attr, found := ts.GetAttribute(
		"showLastColumn",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowLastColumn sets whether special formatting is applied to the last column.
// Attribute: showLastColumn.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ts *TableStyleInfo) SetShowLastColumn(
	value bool,
) {
	if value {
		ts.SetAttribute(
			openxml.NewAttribute(
				"",
				"showLastColumn",
				"",
				attrValueOne,
			),
		)
	} else {
		ts.SetAttribute(
			openxml.NewAttribute("", "showLastColumn", "", attrValueZero),
		)
	}
}

// ShowRowStripes returns whether row stripes are shown.
// Attribute: showRowStripes.
func (ts *TableStyleInfo) ShowRowStripes() bool {
	attr, found := ts.GetAttribute(
		"showRowStripes",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowRowStripes sets whether row stripes are shown.
// Attribute: showRowStripes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ts *TableStyleInfo) SetShowRowStripes(
	value bool,
) {
	if value {
		ts.SetAttribute(
			openxml.NewAttribute(
				"",
				"showRowStripes",
				"",
				attrValueOne,
			),
		)
	} else {
		ts.SetAttribute(
			openxml.NewAttribute("", "showRowStripes", "", attrValueZero),
		)
	}
}

// ShowColumnStripes returns whether column stripes are shown.
// Attribute: showColumnStripes.
func (ts *TableStyleInfo) ShowColumnStripes() bool {
	attr, found := ts.GetAttribute(
		"showColumnStripes",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowColumnStripes sets whether column stripes are shown.
// Attribute: showColumnStripes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ts *TableStyleInfo) SetShowColumnStripes(
	value bool,
) {
	if value {
		ts.SetAttribute(
			openxml.NewAttribute(
				"",
				"showColumnStripes",
				"",
				attrValueOne,
			),
		)
	} else {
		ts.SetAttribute(
			openxml.NewAttribute("", "showColumnStripes", "", attrValueZero),
		)
	}
}

// Clone creates a deep copy of this TableStyleInfo element.
func (ts *TableStyleInfo) Clone() openxml.Element {
	cloned := ts.LeafElementBase.Clone()

	return &TableStyleInfo{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this TableStyleInfo element.
func (ts *TableStyleInfo) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ts.LeafElementBase.CloneNode(deep)

	return &TableStyleInfo{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
