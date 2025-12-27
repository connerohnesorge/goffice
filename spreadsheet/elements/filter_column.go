package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// FilterColumn represents a filterColumn element (x:filterColumn).
// It defines filtering for a specific column within an AutoFilter.
type FilterColumn struct {
	*openxml.CompositeElementBase
}

// NewFilterColumn creates a new FilterColumn element.
func NewFilterColumn() *FilterColumn {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"filterColumn",
		PrefixDefault,
	)

	return &FilterColumn{
		CompositeElementBase: elem,
	}
}

// ColId returns the zero-based column ID. Attribute: colId.
func (fc *FilterColumn) ColId() uint32 {
	attr, found := fc.GetAttribute("colId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:mnd,revive // base 10
		32, //nolint:mnd,revive // 32-bit size
	)

	return uint32(val)
}

// SetColId sets the zero-based column ID. Attribute: colId.
func (fc *FilterColumn) SetColId(colId uint32) {
	fc.SetAttribute(
		openxml.NewAttribute(
			"",
			"colId",
			"",
			strconv.FormatUint(
				uint64(colId),
				10, //nolint:mnd,revive // add-constant
			),
		),
	)
}

// HiddenButton returns whether the filter button is hidden.
// Attribute: hiddenButton.
func (fc *FilterColumn) HiddenButton() bool {
	attr, found := fc.GetAttribute(
		"hiddenButton",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHiddenButton sets whether the filter button is hidden. Attribute: hiddenButton.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (fc *FilterColumn) SetHiddenButton(
	value bool,
) {
	if value {
		fc.SetAttribute(
			openxml.NewAttribute(
				"",
				"hiddenButton",
				"",
				attrValueTrue,
			),
		)
	} else {
		fc.RemoveAttribute("hiddenButton", "")
	}
}

// ShowButton returns whether the filter button is shown. Attribute: showButton.
// Default is true.
func (fc *FilterColumn) ShowButton() bool {
	attr, found := fc.GetAttribute(
		"showButton",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowButton sets whether the filter button is shown. Attribute: showButton.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (fc *FilterColumn) SetShowButton(
	value bool,
) {
	if value {
		fc.RemoveAttribute(
			"showButton",
			"",
		) // Default is true
	} else {
		fc.SetAttribute(
			openxml.NewAttribute(
				"",
				"showButton",
				"",
				attrValueFalse,
			),
		)
	}
}

// Filters returns the Filters child element, or nil if not present.
func (fc *FilterColumn) Filters() *Filters {
	elem := fc.GetElement("filters", NamespaceSML)
	if elem == nil {
		return nil
	}
	if f, ok := elem.(*Filters); ok {
		return f
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Filters{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateFilters returns the Filters child element, creating it if needed.
// This will remove any other filter type that may be present.
func (fc *FilterColumn) GetOrCreateFilters() *Filters {
	f := fc.Filters()
	if f != nil {
		return f
	}
	fc.clearFilterChildren()
	f = NewFilters()
	fc.AppendChild(f)

	return f
}

// Top10 returns the Top10 child element, or nil if not present.
func (fc *FilterColumn) Top10() *Top10 {
	elem := fc.GetElement("top10", NamespaceSML)
	if elem == nil {
		return nil
	}
	if t, ok := elem.(*Top10); ok {
		return t
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Top10{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateTop10 returns the Top10 child element, creating it if needed.
// This will remove any other filter type that may be present.
func (fc *FilterColumn) GetOrCreateTop10() *Top10 {
	t := fc.Top10()
	if t != nil {
		return t
	}
	fc.clearFilterChildren()
	t = NewTop10()
	fc.AppendChild(t)

	return t
}

// CustomFilters returns the CustomFilters child element, or nil if not present.
func (fc *FilterColumn) CustomFilters() *CustomFilters {
	elem := fc.GetElement(
		"customFilters",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if cf, ok := elem.(*CustomFilters); ok {
		return cf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CustomFilters{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCustomFilters returns the CustomFilters child element,
// creating it if needed. This will remove any other filter type
// that may be present.
func (fc *FilterColumn) GetOrCreateCustomFilters() *CustomFilters {
	cf := fc.CustomFilters()
	if cf != nil {
		return cf
	}
	fc.clearFilterChildren()
	cf = NewCustomFilters()
	fc.AppendChild(cf)

	return cf
}

// DynamicFilter returns the DynamicFilter child element, or nil if not present.
func (fc *FilterColumn) DynamicFilter() *DynamicFilter {
	elem := fc.GetElement(
		"dynamicFilter",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if df, ok := elem.(*DynamicFilter); ok {
		return df
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &DynamicFilter{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreateDynamicFilter returns the DynamicFilter child element,
// creating it if needed. This will remove any other filter type
// that may be present.
func (fc *FilterColumn) GetOrCreateDynamicFilter() *DynamicFilter {
	df := fc.DynamicFilter()
	if df != nil {
		return df
	}
	fc.clearFilterChildren()
	df = NewDynamicFilter()
	fc.AppendChild(df)

	return df
}

// ColorFilter returns the ColorFilter child element, or nil if not present.
func (fc *FilterColumn) ColorFilter() *ColorFilter {
	elem := fc.GetElement(
		"colorFilter",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if cf, ok := elem.(*ColorFilter); ok {
		return cf
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &ColorFilter{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateColorFilter returns the ColorFilter child element,
// creating it if needed. This will remove any other filter type
// that may be present.
func (fc *FilterColumn) GetOrCreateColorFilter() *ColorFilter {
	cf := fc.ColorFilter()
	if cf != nil {
		return cf
	}
	fc.clearFilterChildren()
	cf = NewColorFilter()
	fc.AppendChild(cf)

	return cf
}

// IconFilter returns the IconFilter child element, or nil if not present.
func (fc *FilterColumn) IconFilter() *IconFilter {
	elem := fc.GetElement(
		"iconFilter",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if i, ok := elem.(*IconFilter); ok {
		return i
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &IconFilter{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateIconFilter returns the IconFilter child element, creating it if
// needed. This will remove any other filter type that may be present.
func (fc *FilterColumn) GetOrCreateIconFilter() *IconFilter {
	i := fc.IconFilter()
	if i != nil {
		return i
	}
	fc.clearFilterChildren()
	i = NewIconFilter()
	fc.AppendChild(i)

	return i
}

// clearFilterChildren removes all filter-type children.
// Only one filter type is allowed per column.
func (fc *FilterColumn) clearFilterChildren() {
	// Collect elements to remove
	var toRemove []openxml.Element
	for child := range fc.Children() {
		if child.NamespaceURI() != NamespaceSML {
			continue
		}
		switch child.LocalName() {
		case "filters",
			"top10",
			"customFilters",
			"dynamicFilter",
			"colorFilter",
			"iconFilter":
			toRemove = append(toRemove, child)
		}
	}
	// Remove collected elements
	for _, elem := range toRemove {
		fc.RemoveChild(elem)
	}
}

// ClearFilter removes all filter settings from this column.
func (fc *FilterColumn) ClearFilter() {
	fc.clearFilterChildren()
}

// Clone creates a deep copy of this FilterColumn element.
func (fc *FilterColumn) Clone() openxml.Element {
	cloned := fc.CompositeElementBase.Clone()

	return &FilterColumn{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this FilterColumn element.
func (fc *FilterColumn) CloneNode(
	deep bool,
) openxml.Element {
	cloned := fc.CompositeElementBase.CloneNode(
		deep,
	)

	return &FilterColumn{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
