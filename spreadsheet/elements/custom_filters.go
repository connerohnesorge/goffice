package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// FilterOperator represents the comparison operator for custom filters.
type FilterOperator string

const (
	// FilterOperatorLessThan indicates less than.
	FilterOperatorLessThan FilterOperator = "lessThan"
	// FilterOperatorLessThanOrEqual indicates less than or equal.
	FilterOperatorLessThanOrEqual FilterOperator = "lessThanOrEqual"
	// FilterOperatorEqual indicates equal (default).
	FilterOperatorEqual FilterOperator = "equal"
	// FilterOperatorNotEqual indicates not equal.
	FilterOperatorNotEqual FilterOperator = "notEqual"
	// FilterOperatorGreaterThanOrEqual indicates greater than or equal.
	FilterOperatorGreaterThanOrEqual FilterOperator = "greaterThanOrEqual"
	// FilterOperatorGreaterThan indicates greater than.
	FilterOperatorGreaterThan FilterOperator = "greaterThan"
)

// CustomFilters represents the customFilters container element
// (x:customFilters). It contains one or two custom filter criteria for a
// column.
type CustomFilters struct {
	*openxml.CompositeElementBase
}

// NewCustomFilters creates a new CustomFilters element.
func NewCustomFilters() *CustomFilters {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"customFilters",
		PrefixDefault,
	)

	return &CustomFilters{
		CompositeElementBase: elem,
	}
}

// And returns whether to use AND logic (true) or OR logic (false) between
// criteria. Attribute: and. Default is false (OR logic).
func (cf *CustomFilters) And() bool {
	attr, found := cf.GetAttribute("and", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAnd sets whether to use AND logic (true) or OR logic (false) between criteria.
// Attribute: and.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CustomFilters) SetAnd(value bool) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"and",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("and", "")
	}
}

// GetCustomFilters returns an iterator over all CustomFilter elements.
func (cf *CustomFilters) GetCustomFilters() iter.Seq[*CustomFilter] {
	return func(yield func(*CustomFilter) bool) {
		for child := range cf.Children() {
			if child.LocalName() != "customFilter" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var filter *CustomFilter
			switch v := child.(type) {
			case *CustomFilter:
				filter = v
			case *openxml.LeafElementBase:
				filter = &CustomFilter{LeafElementBase: v}
			}
			if filter != nil && !yield(filter) {
				return
			}
		}
	}
}

// CustomFilterCount returns the number of CustomFilter elements.
func (cf *CustomFilters) CustomFilterCount() int {
	count := 0
	for range cf.GetCustomFilters() {
		count++
	}

	return count
}

// AddCustomFilter adds a new CustomFilter element with the given
// value and operator.
func (cf *CustomFilters) AddCustomFilter(
	value string,
	operator FilterOperator,
) *CustomFilter {
	filter := NewCustomFilter()
	filter.SetVal(value)
	filter.SetOperator(operator)
	cf.AppendChild(filter)

	return filter
}

// SetSingleFilter sets a single filter criterion, removing any existing
// filters.
func (cf *CustomFilters) SetSingleFilter(
	value string,
	operator FilterOperator,
) *CustomFilter {
	cf.ClearFilters()

	return cf.AddCustomFilter(value, operator)
}

// SetDualFilters sets two filter criteria with AND or OR logic.
//
//nolint:revive // argument-limit
func (cf *CustomFilters) SetDualFilters(
	value1 string,
	operator1 FilterOperator,
	value2 string,
	operator2 FilterOperator,
	andLogic bool,
) {
	cf.ClearFilters()
	cf.SetAnd(andLogic)
	cf.AddCustomFilter(value1, operator1)
	cf.AddCustomFilter(value2, operator2)
}

// ClearFilters removes all CustomFilter children.
func (cf *CustomFilters) ClearFilters() {
	// Count children first for pre-allocation
	count := 0
	for range cf.GetCustomFilters() {
		count++
	}
	toRemove := make([]openxml.Element, 0, count)
	for filter := range cf.GetCustomFilters() {
		toRemove = append(toRemove, filter)
	}
	for _, elem := range toRemove {
		cf.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this CustomFilters element.
func (cf *CustomFilters) Clone() openxml.Element {
	cloned := cf.CompositeElementBase.Clone()

	return &CustomFilters{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CustomFilters element.
func (cf *CustomFilters) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cf.CompositeElementBase.CloneNode(
		deep,
	)

	return &CustomFilters{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CustomFilter represents a single custom filter criterion element
// (x:customFilter).
type CustomFilter struct {
	*openxml.LeafElementBase
}

// NewCustomFilter creates a new CustomFilter element.
func NewCustomFilter() *CustomFilter {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"customFilter",
		PrefixDefault,
	)

	return &CustomFilter{LeafElementBase: elem}
}

// Val returns the filter value. Attribute: val.
// Supports wildcards: * (any characters), ? (single character).
func (c *CustomFilter) Val() string {
	attr, found := c.GetAttribute("val", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetVal sets the filter value. Attribute: val.
// Supports wildcards: * (any characters), ? (single character).
func (c *CustomFilter) SetVal(value string) {
	if value == "" {
		c.RemoveAttribute("val", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			value,
		),
	)
}

// Operator returns the comparison operator. Attribute: operator.
// Default is "equal".
func (c *CustomFilter) Operator() FilterOperator {
	attr, found := c.GetAttribute("operator", "")
	if !found {
		return FilterOperatorEqual
	}

	return FilterOperator(attr.Value())
}

// SetOperator sets the comparison operator. Attribute: operator.
func (c *CustomFilter) SetOperator(
	operator FilterOperator,
) {
	if operator == "" ||
		operator == FilterOperatorEqual {
		c.RemoveAttribute("operator", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"operator",
			"",
			string(operator),
		),
	)
}

// Clone creates a deep copy of this CustomFilter element.
func (c *CustomFilter) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &CustomFilter{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CustomFilter element.
func (c *CustomFilter) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &CustomFilter{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
