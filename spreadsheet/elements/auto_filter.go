package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// AutoFilter represents the autoFilter element (x:autoFilter).
// It defines a filtered range of data in a worksheet.
type AutoFilter struct {
	*openxml.CompositeElementBase
}

// NewAutoFilter creates a new AutoFilter element.
func NewAutoFilter() *AutoFilter {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"autoFilter",
		PrefixDefault,
	)

	return &AutoFilter{CompositeElementBase: elem}
}

// NewAutoFilterWithRef creates a new AutoFilter element with the given
// range reference.
func NewAutoFilterWithRef(
	ref string,
) *AutoFilter {
	af := NewAutoFilter()
	af.SetRef(ref)

	return af
}

// Ref returns the range reference for the auto filter. Attribute: ref.
// Example: "A1:D100"
func (af *AutoFilter) Ref() string {
	attr, found := af.GetAttribute("ref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the range reference for the auto filter. Attribute: ref.
// Example: "A1:D100"
func (af *AutoFilter) SetRef(ref string) {
	if ref == "" {
		af.RemoveAttribute("ref", "")

		return
	}
	af.SetAttribute(
		openxml.NewAttribute(
			"",
			"ref",
			"",
			ref,
		),
	)
}

// FilterColumns returns an iterator over all FilterColumn elements.
func (af *AutoFilter) FilterColumns() iter.Seq[*FilterColumn] {
	return func(yield func(*FilterColumn) bool) {
		for child := range af.Children() {
			if child.LocalName() != "filterColumn" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var fc *FilterColumn
			if f, ok := child.(*FilterColumn); ok {
				fc = f
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				fc = &FilterColumn{CompositeElementBase: comp}
			}
			if fc != nil && !yield(fc) {
				return
			}
		}
	}
}

// FilterColumnCount returns the number of filter columns.
func (af *AutoFilter) FilterColumnCount() int {
	count := 0
	for range af.FilterColumns() {
		count++
	}

	return count
}

// GetFilterColumnByColId returns the filter column with the given column ID,
// or nil if not found.
func (af *AutoFilter) GetFilterColumnByColId(
	colId uint32,
) *FilterColumn {
	for fc := range af.FilterColumns() {
		if fc.ColId() == colId {
			return fc
		}
	}

	return nil
}

// AddFilterColumn adds a new FilterColumn element for the given column ID.
func (af *AutoFilter) AddFilterColumn(
	colId uint32,
) *FilterColumn {
	fc := NewFilterColumn()
	fc.SetColId(colId)
	// Insert before sortState if it exists
	sortState := af.SortState()
	if sortState != nil {
		af.InsertBefore(fc, sortState)
	} else {
		af.AppendChild(fc)
	}

	return fc
}

// RemoveFilterColumn removes the filter column with the given column ID.
func (af *AutoFilter) RemoveFilterColumn(
	colId uint32,
) bool {
	fc := af.GetFilterColumnByColId(colId)
	if fc == nil {
		return false
	}

	return af.RemoveChild(fc)
}

// SortState returns the SortState child element, or nil if not present.
func (af *AutoFilter) SortState() *SortState {
	elem := af.GetElement(
		"sortState",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if ss, ok := elem.(*SortState); ok {
		return ss
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SortState{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSortState returns the SortState child element, creating it
// if needed.
func (af *AutoFilter) GetOrCreateSortState() *SortState {
	ss := af.SortState()
	if ss != nil {
		return ss
	}
	ss = NewSortState()
	af.AppendChild(ss)

	return ss
}

// RemoveSortState removes the SortState child element.
func (af *AutoFilter) RemoveSortState() bool {
	ss := af.SortState()
	if ss == nil {
		return false
	}

	return af.RemoveChild(ss)
}

// ClearFilters removes all filter columns.
func (af *AutoFilter) ClearFilters() {
	// Count children first for pre-allocation
	count := 0
	for range af.FilterColumns() {
		count++
	}
	toRemove := make([]openxml.Element, 0, count)
	for fc := range af.FilterColumns() {
		toRemove = append(toRemove, fc)
	}
	for _, fc := range toRemove {
		af.RemoveChild(fc)
	}
}

// Clone creates a deep copy of this AutoFilter element.
func (af *AutoFilter) Clone() openxml.Element {
	cloned := af.CompositeElementBase.Clone()

	return &AutoFilter{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this AutoFilter element.
func (af *AutoFilter) CloneNode(
	deep bool,
) openxml.Element {
	cloned := af.CompositeElementBase.CloneNode(
		deep,
	)

	return &AutoFilter{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
