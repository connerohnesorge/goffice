package elements

//revive:disable:file-length-limit many sort state properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// SortMethod represents the method used for sorting.
type SortMethod string

const (
	// SortMethodStroke indicates stroke sorting.
	SortMethodStroke SortMethod = "stroke"
	// SortMethodPinYin indicates PinYin sorting.
	SortMethodPinYin SortMethod = "pinYin"
	// SortMethodNone indicates no special sorting method.
	SortMethodNone SortMethod = "none"
)

// SortBy represents what property to sort by.
type SortBy string

const (
	// SortByValue sorts by cell value.
	SortByValue SortBy = "value"
	// SortByCellColor sorts by cell background color.
	SortByCellColor SortBy = "cellColor"
	// SortByFontColor sorts by font color.
	SortByFontColor SortBy = "fontColor"
	// SortByIcon sorts by conditional formatting icon.
	SortByIcon SortBy = "icon"
)

// SortState represents the sortState element (x:sortState).
// It defines sorting state for an AutoFilter or table.
type SortState struct {
	*openxml.CompositeElementBase
}

// NewSortState creates a new SortState element.
func NewSortState() *SortState {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"sortState",
		PrefixDefault,
	)

	return &SortState{CompositeElementBase: elem}
}

// Ref returns the range reference for sorting. Attribute: ref.
func (ss *SortState) Ref() string {
	attr, found := ss.GetAttribute("ref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the range reference for sorting. Attribute: ref.
func (ss *SortState) SetRef(ref string) {
	if ref == "" {
		ss.RemoveAttribute("ref", "")

		return
	}
	ss.SetAttribute(
		openxml.NewAttribute(
			"",
			"ref",
			"",
			ref,
		),
	)
}

// CaseSensitive returns whether sorting is case-sensitive.
// Attribute: caseSensitive.
func (ss *SortState) CaseSensitive() bool {
	attr, found := ss.GetAttribute(
		"caseSensitive",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCaseSensitive sets whether sorting is case-sensitive. Attribute: caseSensitive.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ss *SortState) SetCaseSensitive(
	value bool,
) {
	if value {
		ss.SetAttribute(
			openxml.NewAttribute(
				"",
				"caseSensitive",
				"",
				attrValueTrue,
			),
		)
	} else {
		ss.RemoveAttribute("caseSensitive", "")
	}
}

// ColumnSort returns whether to sort by columns (true) or rows (false).
// Attribute: columnSort. Default is false (sort by rows).
func (ss *SortState) ColumnSort() bool {
	attr, found := ss.GetAttribute(
		"columnSort",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetColumnSort sets whether to sort by columns (true) or rows (false).
// Attribute: columnSort.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ss *SortState) SetColumnSort(value bool) {
	if value {
		ss.SetAttribute(
			openxml.NewAttribute(
				"",
				"columnSort",
				"",
				attrValueTrue,
			),
		)
	} else {
		ss.RemoveAttribute("columnSort", "")
	}
}

// SortMethod returns the sort method. Attribute: sortMethod.
func (ss *SortState) SortMethod() SortMethod {
	attr, found := ss.GetAttribute(
		"sortMethod",
		"",
	)
	if !found {
		return SortMethodNone
	}

	return SortMethod(attr.Value())
}

// SetSortMethod sets the sort method. Attribute: sortMethod.
func (ss *SortState) SetSortMethod(
	method SortMethod,
) {
	if method == "" || method == SortMethodNone {
		ss.RemoveAttribute("sortMethod", "")

		return
	}
	ss.SetAttribute(
		openxml.NewAttribute(
			"",
			"sortMethod",
			"",
			string(method),
		),
	)
}

// GetSortConditions returns an iterator over all SortCondition elements.
func (ss *SortState) GetSortConditions() iter.Seq[*SortCondition] {
	return func(yield func(*SortCondition) bool) {
		for child := range ss.Children() {
			if child.LocalName() != "sortCondition" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var sc *SortCondition
			switch v := child.(type) {
			case *SortCondition:
				sc = v
			case *openxml.LeafElementBase:
				sc = &SortCondition{LeafElementBase: v}
			}
			if sc != nil && !yield(sc) {
				return
			}
		}
	}
}

// SortConditionCount returns the number of SortCondition elements.
func (ss *SortState) SortConditionCount() int {
	count := 0
	for range ss.GetSortConditions() {
		count++
	}

	return count
}

// AddSortCondition adds a new SortCondition element with the given reference.
func (ss *SortState) AddSortCondition(
	ref string,
) *SortCondition {
	sc := NewSortCondition()
	sc.SetRef(ref)
	ss.AppendChild(sc)

	return sc
}

// AddDescendingSortCondition adds a new descending sort condition.
func (ss *SortState) AddDescendingSortCondition(
	ref string,
) *SortCondition {
	sc := ss.AddSortCondition(ref)
	sc.SetDescending(true)

	return sc
}

// ClearSortConditions removes all SortCondition children.
func (ss *SortState) ClearSortConditions() {
	count := 0
	for range ss.GetSortConditions() {
		count++
	}
	toRemove := make([]openxml.Element, 0, count)
	for sc := range ss.GetSortConditions() {
		toRemove = append(toRemove, sc)
	}
	for _, elem := range toRemove {
		ss.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this SortState element.
func (ss *SortState) Clone() openxml.Element {
	cloned := ss.CompositeElementBase.Clone()

	return &SortState{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SortState element.
func (ss *SortState) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ss.CompositeElementBase.CloneNode(
		deep,
	)

	return &SortState{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// SortCondition represents a single sort condition element (x:sortCondition).
type SortCondition struct {
	*openxml.LeafElementBase
}

// NewSortCondition creates a new SortCondition element.
func NewSortCondition() *SortCondition {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"sortCondition",
		PrefixDefault,
	)

	return &SortCondition{LeafElementBase: elem}
}

// Ref returns the range reference for this sort condition. Attribute: ref.
func (sc *SortCondition) Ref() string {
	attr, found := sc.GetAttribute(
		elemNameRef,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the range reference for this sort condition. Attribute: ref.
func (sc *SortCondition) SetRef(ref string) {
	if ref == "" {
		sc.RemoveAttribute(elemNameRef, "")

		return
	}
	sc.SetAttribute(
		openxml.NewAttribute(
			"",
			elemNameRef,
			"",
			ref,
		),
	)
}

// Descending returns whether to sort in descending order.
// Attribute: descending.
func (sc *SortCondition) Descending() bool {
	attr, found := sc.GetAttribute(
		"descending",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDescending sets whether to sort in descending order. Attribute: descending.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sc *SortCondition) SetDescending(
	value bool,
) {
	if value {
		sc.SetAttribute(
			openxml.NewAttribute(
				"",
				"descending",
				"",
				attrValueTrue,
			),
		)
	} else {
		sc.RemoveAttribute("descending", "")
	}
}

// SortBy returns what property to sort by. Attribute: sortBy.
func (sc *SortCondition) SortBy() SortBy {
	attr, found := sc.GetAttribute("sortBy", "")
	if !found {
		return SortByValue
	}

	return SortBy(attr.Value())
}

// SetSortBy sets what property to sort by. Attribute: sortBy.
func (sc *SortCondition) SetSortBy(
	sortBy SortBy,
) {
	if sortBy == "" || sortBy == SortByValue {
		sc.RemoveAttribute("sortBy", "")

		return
	}
	sc.SetAttribute(
		openxml.NewAttribute(
			"",
			"sortBy",
			"",
			string(sortBy),
		),
	)
}

// DxfId returns the differential formatting ID for color sorting.
// Attribute: dxfId.
func (sc *SortCondition) DxfId() uint32 {
	attr, found := sc.GetAttribute("dxfId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetDxfId sets the differential formatting ID for color sorting.
// Attribute: dxfId.
func (sc *SortCondition) SetDxfId(id uint32) {
	if id == 0 {
		sc.RemoveAttribute("dxfId", "")

		return
	}
	sc.SetAttribute(
		openxml.NewAttribute(
			"",
			"dxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10,
			),
		),
	)
}

// CustomList returns the custom list string for sorting. Attribute: customList.
func (sc *SortCondition) CustomList() string {
	attr, found := sc.GetAttribute(
		"customList",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCustomList sets the custom list string for sorting. Attribute: customList.
func (sc *SortCondition) SetCustomList(
	list string,
) {
	if list == "" {
		sc.RemoveAttribute("customList", "")

		return
	}
	sc.SetAttribute(
		openxml.NewAttribute(
			"",
			"customList",
			"",
			list,
		),
	)
}

// IconSet returns the icon set for icon sorting. Attribute: iconSet.
func (sc *SortCondition) IconSet() IconSetType {
	attr, found := sc.GetAttribute("iconSet", "")
	if !found {
		return ""
	}

	return IconSetType(attr.Value())
}

// SetIconSet sets the icon set for icon sorting. Attribute: iconSet.
func (sc *SortCondition) SetIconSet(
	iconSet IconSetType,
) {
	if iconSet == "" {
		sc.RemoveAttribute("iconSet", "")

		return
	}
	sc.SetAttribute(
		openxml.NewAttribute(
			"",
			"iconSet",
			"",
			string(iconSet),
		),
	)
}

// IconId returns the icon index for icon sorting. Attribute: iconId.
func (sc *SortCondition) IconId() uint32 {
	attr, found := sc.GetAttribute("iconId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetIconId sets the icon index for icon sorting. Attribute: iconId.
func (sc *SortCondition) SetIconId(id uint32) {
	if id == 0 {
		sc.RemoveAttribute("iconId", "")

		return
	}
	sc.SetAttribute(
		openxml.NewAttribute(
			"",
			"iconId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10,
			),
		),
	)
}

// Clone creates a deep copy of this SortCondition element.
func (sc *SortCondition) Clone() openxml.Element {
	cloned := sc.LeafElementBase.Clone()

	return &SortCondition{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this SortCondition element.
func (sc *SortCondition) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sc.LeafElementBase.CloneNode(deep)

	return &SortCondition{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
