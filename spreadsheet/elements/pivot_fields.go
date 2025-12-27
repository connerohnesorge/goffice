package elements

//revive:disable:file-length-limit many pivot field types
//revive:disable:max-public-structs many pivot field types

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PivotFields represents the pivot fields collection element (x:pivotFields).
type PivotFields struct {
	*openxml.CompositeElementBase
}

// NewPivotFields creates a new PivotFields element.
func NewPivotFields() *PivotFields {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"pivotFields",
		PrefixDefault,
	)

	return &PivotFields{
		CompositeElementBase: elem,
	}
}

// Count returns the count of pivot fields. Attribute: count.
func (pf *PivotFields) Count() uint32 {
	attr, found := pf.GetAttribute("count", "")
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

// SetCount sets the count of pivot fields. Attribute: count.
func (pf *PivotFields) SetCount(count uint32) {
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// PivotFields returns an iterator over all PivotField elements.
func (pf *PivotFields) PivotFields() iter.Seq[*PivotField] {
	return func(yield func(*PivotField) bool) {
		for child := range pf.Children() {
			if child.LocalName() != "pivotField" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var field *PivotField
			if f, ok := child.(*PivotField); ok {
				field = f
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				field = &PivotField{CompositeElementBase: comp}
			}
			if field != nil && !yield(field) {
				return
			}
		}
	}
}

// FieldCount returns the count of pivot field elements.
func (pf *PivotFields) FieldCount() int {
	count := 0
	for range pf.PivotFields() {
		count++
	}

	return count
}

// GetField returns the pivot field at the given index, or nil if out of range.
func (pf *PivotFields) GetField(
	index int,
) *PivotField {
	if index < 0 {
		return nil
	}
	i := 0
	for field := range pf.PivotFields() {
		if i == index {
			return field
		}
		i++
	}

	return nil
}

// AddField adds a new PivotField element.
func (pf *PivotFields) AddField() *PivotField {
	field := NewPivotField()
	pf.AppendChild(field)

	return field
}

// Clone creates a deep copy of this PivotFields element.
func (pf *PivotFields) Clone() openxml.Element {
	cloned := pf.CompositeElementBase.Clone()

	return &PivotFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PivotFields element.
func (pf *PivotFields) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pf.CompositeElementBase.CloneNode(
		deep,
	)

	return &PivotFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// PivotField represents a single pivot field element (x:pivotField).
type PivotField struct {
	*openxml.CompositeElementBase
}

// NewPivotField creates a new PivotField element.
func NewPivotField() *PivotField {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"pivotField",
		PrefixDefault,
	)

	return &PivotField{CompositeElementBase: elem}
}

// Name returns the field name. Attribute: name.
func (pf *PivotField) Name() string {
	attr, found := pf.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the field name. Attribute: name.
func (pf *PivotField) SetName(name string) {
	if name == "" {
		pf.RemoveAttribute("name", "")

		return
	}
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Axis returns the axis type. Attribute: axis.
func (pf *PivotField) Axis() string {
	attr, found := pf.GetAttribute("axis", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAxis sets the axis type. Attribute: axis.
// Valid values: axisRow, axisCol, axisPage, axisValues
func (pf *PivotField) SetAxis(axis string) {
	if axis == "" {
		pf.RemoveAttribute("axis", "")

		return
	}
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"axis",
			"",
			axis,
		),
	)
}

// DataField returns whether this is a data field. Attribute: dataField.
func (pf *PivotField) DataField() bool {
	attr, found := pf.GetAttribute(
		"dataField",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDataField sets whether this is a data field. Attribute: dataField.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetDataField(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"dataField",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("dataField", "")
	}
}

// SubtotalCaption returns the subtotal caption. Attribute: subtotalCaption.
func (pf *PivotField) SubtotalCaption() string {
	attr, found := pf.GetAttribute(
		"subtotalCaption",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSubtotalCaption sets the subtotal caption. Attribute: subtotalCaption.
func (pf *PivotField) SetSubtotalCaption(
	caption string,
) {
	if caption == "" {
		pf.RemoveAttribute("subtotalCaption", "")

		return
	}
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"subtotalCaption",
			"",
			caption,
		),
	)
}

// ShowDropDowns returns whether to show drop downs. Attribute: showDropDowns.
func (pf *PivotField) ShowDropDowns() bool {
	attr, found := pf.GetAttribute(
		"showDropDowns",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowDropDowns sets whether to show drop downs. Attribute: showDropDowns.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetShowDropDowns(
	value bool,
) {
	if value {
		pf.RemoveAttribute(
			"showDropDowns",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "showDropDowns", "", attrValueFalse),
		)
	}
}

// HiddenLevel returns whether the level is hidden. Attribute: hiddenLevel.
func (pf *PivotField) HiddenLevel() bool {
	attr, found := pf.GetAttribute(
		"hiddenLevel",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHiddenLevel sets whether the level is hidden. Attribute: hiddenLevel.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetHiddenLevel(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"hiddenLevel",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("hiddenLevel", "")
	}
}

// UniqueMemberProperty returns the unique member property.
// Attribute: uniqueMemberProperty.
func (pf *PivotField) UniqueMemberProperty() string {
	attr, found := pf.GetAttribute(
		"uniqueMemberProperty",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetUniqueMemberProperty sets the unique member property.
// Attribute: uniqueMemberProperty.
func (pf *PivotField) SetUniqueMemberProperty(
	prop string,
) {
	if prop == "" {
		pf.RemoveAttribute(
			"uniqueMemberProperty",
			"",
		)

		return
	}
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"uniqueMemberProperty",
			"",
			prop,
		),
	)
}

// Compact returns whether to use compact layout. Attribute: compact.
func (pf *PivotField) Compact() bool {
	attr, found := pf.GetAttribute("compact", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetCompact sets whether to use compact layout. Attribute: compact.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetCompact(value bool) {
	if value {
		pf.RemoveAttribute(
			"compact",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "compact", "", attrValueFalse),
		)
	}
}

// AllDrilled returns whether all items are drilled. Attribute: allDrilled.
func (pf *PivotField) AllDrilled() bool {
	attr, found := pf.GetAttribute(
		"allDrilled",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAllDrilled sets whether all items are drilled. Attribute: allDrilled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetAllDrilled(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"allDrilled",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("allDrilled", "")
	}
}

// NumFmtId returns the number format ID. Attribute: numFmtId.
func (pf *PivotField) NumFmtId() (uint32, bool) {
	attr, found := pf.GetAttribute("numFmtId", "")
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val), true
}

// SetNumFmtId sets the number format ID. Attribute: numFmtId.
func (pf *PivotField) SetNumFmtId(id uint32) {
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"numFmtId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10,
			),
		),
	)
}

// ClearNumFmtId removes the number format ID attribute.
func (pf *PivotField) ClearNumFmtId() {
	pf.RemoveAttribute("numFmtId", "")
}

// Outline returns whether to use outline layout. Attribute: outline.
func (pf *PivotField) Outline() bool {
	attr, found := pf.GetAttribute("outline", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetOutline sets whether to use outline layout. Attribute: outline.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetOutline(value bool) {
	if value {
		pf.RemoveAttribute(
			"outline",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "outline", "", attrValueFalse),
		)
	}
}

// SubtotalTop returns whether subtotals are at the top. Attribute: subtotalTop.
func (pf *PivotField) SubtotalTop() bool {
	attr, found := pf.GetAttribute(
		"subtotalTop",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetSubtotalTop sets whether subtotals are at the top. Attribute: subtotalTop.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetSubtotalTop(value bool) {
	if value {
		pf.RemoveAttribute(
			"subtotalTop",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "subtotalTop", "", attrValueFalse),
		)
	}
}

// DragToRow returns whether field can be dragged to rows.
// Attribute: dragToRow.
func (pf *PivotField) DragToRow() bool {
	attr, found := pf.GetAttribute(
		"dragToRow",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetDragToRow sets whether field can be dragged to rows.
// Attribute: dragToRow.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetDragToRow(value bool) {
	if value {
		pf.RemoveAttribute(
			"dragToRow",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "dragToRow", "", attrValueFalse),
		)
	}
}

// DragToCol returns whether field can be dragged to columns.
// Attribute: dragToCol.
func (pf *PivotField) DragToCol() bool {
	attr, found := pf.GetAttribute(
		"dragToCol",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetDragToCol sets whether field can be dragged to columns.
// Attribute: dragToCol.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetDragToCol(value bool) {
	if value {
		pf.RemoveAttribute(
			"dragToCol",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "dragToCol", "", attrValueFalse),
		)
	}
}

// MultipleItemSelectionAllowed returns whether multiple item selection is
// allowed. Attribute: multipleItemSelectionAllowed.
func (pf *PivotField) MultipleItemSelectionAllowed() bool {
	attr, found := pf.GetAttribute(
		"multipleItemSelectionAllowed",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetMultipleItemSelectionAllowed sets whether multiple item selection is allowed.
// Attribute: multipleItemSelectionAllowed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetMultipleItemSelectionAllowed(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"multipleItemSelectionAllowed",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("multipleItemSelectionAllowed", "")
	}
}

// DragToPage returns whether field can be dragged to page.
// Attribute: dragToPage.
func (pf *PivotField) DragToPage() bool {
	attr, found := pf.GetAttribute(
		"dragToPage",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetDragToPage sets whether field can be dragged to page.
// Attribute: dragToPage.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetDragToPage(value bool) {
	if value {
		pf.RemoveAttribute(
			"dragToPage",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "dragToPage", "", attrValueFalse),
		)
	}
}

// DragToData returns whether field can be dragged to data.
// Attribute: dragToData.
func (pf *PivotField) DragToData() bool {
	attr, found := pf.GetAttribute(
		"dragToData",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetDragToData sets whether field can be dragged to data.
// Attribute: dragToData.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetDragToData(value bool) {
	if value {
		pf.RemoveAttribute(
			"dragToData",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "dragToData", "", attrValueFalse),
		)
	}
}

// DragOff returns whether field can be dragged off.
// Attribute: dragOff.
func (pf *PivotField) DragOff() bool {
	attr, found := pf.GetAttribute("dragOff", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetDragOff sets whether field can be dragged off.
// Attribute: dragOff.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetDragOff(value bool) {
	if value {
		pf.RemoveAttribute(
			"dragOff",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "dragOff", "", attrValueFalse),
		)
	}
}

// ShowAll returns whether to show all items. Attribute: showAll.
func (pf *PivotField) ShowAll() bool {
	attr, found := pf.GetAttribute("showAll", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowAll sets whether to show all items. Attribute: showAll.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetShowAll(value bool) {
	if value {
		pf.RemoveAttribute(
			"showAll",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "showAll", "", attrValueFalse),
		)
	}
}

// InsertBlankRow returns whether to insert blank row.
// Attribute: insertBlankRow.
func (pf *PivotField) InsertBlankRow() bool {
	attr, found := pf.GetAttribute(
		"insertBlankRow",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetInsertBlankRow sets whether to insert blank row. Attribute: insertBlankRow.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetInsertBlankRow(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"insertBlankRow",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("insertBlankRow", "")
	}
}

// ServerField returns whether this is a server field. Attribute: serverField.
func (pf *PivotField) ServerField() bool {
	attr, found := pf.GetAttribute(
		"serverField",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetServerField sets whether this is a server field. Attribute: serverField.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetServerField(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"serverField",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("serverField", "")
	}
}

// InsertPageBreak returns whether to insert page break.
// Attribute: insertPageBreak.
func (pf *PivotField) InsertPageBreak() bool {
	attr, found := pf.GetAttribute(
		"insertPageBreak",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetInsertPageBreak sets whether to insert page break.
// Attribute: insertPageBreak.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetInsertPageBreak(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"insertPageBreak",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("insertPageBreak", "")
	}
}

// AutoShow returns whether to use auto show. Attribute: autoShow.
func (pf *PivotField) AutoShow() bool {
	attr, found := pf.GetAttribute("autoShow", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAutoShow sets whether to use auto show. Attribute: autoShow.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetAutoShow(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"autoShow",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("autoShow", "")
	}
}

// TopAutoShow returns whether to show top items. Attribute: topAutoShow.
func (pf *PivotField) TopAutoShow() bool {
	attr, found := pf.GetAttribute(
		"topAutoShow",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetTopAutoShow sets whether to show top items. Attribute: topAutoShow.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetTopAutoShow(value bool) {
	if value {
		pf.RemoveAttribute(
			"topAutoShow",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "topAutoShow", "", attrValueFalse),
		)
	}
}

// HideNewItems returns whether to hide new items. Attribute: hideNewItems.
func (pf *PivotField) HideNewItems() bool {
	attr, found := pf.GetAttribute(
		"hideNewItems",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHideNewItems sets whether to hide new items. Attribute: hideNewItems.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetHideNewItems(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"hideNewItems",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("hideNewItems", "")
	}
}

// MeasureFilter returns whether this is a measure filter.
// Attribute: measureFilter.
func (pf *PivotField) MeasureFilter() bool {
	attr, found := pf.GetAttribute(
		"measureFilter",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetMeasureFilter sets whether this is a measure filter.
// Attribute: measureFilter.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetMeasureFilter(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"measureFilter",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("measureFilter", "")
	}
}

// IncludeNewItemsInFilter returns whether to include new items in filter.
// Attribute: includeNewItemsInFilter.
func (pf *PivotField) IncludeNewItemsInFilter() bool {
	attr, found := pf.GetAttribute(
		"includeNewItemsInFilter",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetIncludeNewItemsInFilter sets whether to include new items in filter.
// Attribute: includeNewItemsInFilter.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetIncludeNewItemsInFilter(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"includeNewItemsInFilter",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("includeNewItemsInFilter", "")
	}
}

// ItemPageCount returns the item page count. Attribute: itemPageCount.
func (pf *PivotField) ItemPageCount() uint32 {
	attr, found := pf.GetAttribute(
		"itemPageCount",
		"",
	)
	if !found {
		return 10 //nolint:revive // add-constant - Default is 10
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetItemPageCount sets the item page count. Attribute: itemPageCount.
func (pf *PivotField) SetItemPageCount(
	count uint32,
) {
	if count == 10 { //nolint:revive // add-constant - 10 is default
		pf.RemoveAttribute(
			"itemPageCount",
			"",
		)

		return
	}
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"itemPageCount",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// SortType returns the sort type. Attribute: sortType.
func (pf *PivotField) SortType() string {
	attr, found := pf.GetAttribute("sortType", "")
	if !found {
		return "manual" // Default is manual
	}

	return attr.Value()
}

// SetSortType sets the sort type. Attribute: sortType.
// Valid values: manual, ascending, descending
func (pf *PivotField) SetSortType(
	sortType string,
) {
	if sortType == "" || sortType == "manual" {
		pf.RemoveAttribute(
			"sortType",
			"",
		) // manual is default

		return
	}
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"sortType",
			"",
			sortType,
		),
	)
}

// DataSourceSort returns whether to use data source sort.
// Attribute: dataSourceSort.
func (pf *PivotField) DataSourceSort() bool {
	attr, found := pf.GetAttribute(
		"dataSourceSort",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDataSourceSort sets whether to use data source sort.
// Attribute: dataSourceSort.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetDataSourceSort(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"dataSourceSort",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("dataSourceSort", "")
	}
}

// NonAutoSortDefault returns whether non-auto sort is the default.
// Attribute: nonAutoSortDefault.
func (pf *PivotField) NonAutoSortDefault() bool {
	attr, found := pf.GetAttribute(
		"nonAutoSortDefault",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetNonAutoSortDefault sets whether non-auto sort is the default.
// Attribute: nonAutoSortDefault.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetNonAutoSortDefault(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"nonAutoSortDefault",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("nonAutoSortDefault", "")
	}
}

// RankBy returns the rank by field index. Attribute: rankBy.
func (pf *PivotField) RankBy() (uint32, bool) {
	attr, found := pf.GetAttribute("rankBy", "")
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val), true
}

// SetRankBy sets the rank by field index. Attribute: rankBy.
func (pf *PivotField) SetRankBy(index uint32) {
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"rankBy",
			"",
			strconv.FormatUint(
				uint64(index),
				parseBase10,
			),
		),
	)
}

// ClearRankBy removes the rank by attribute.
func (pf *PivotField) ClearRankBy() {
	pf.RemoveAttribute("rankBy", "")
}

// DefaultSubtotal returns whether to show default subtotal.
// Attribute: defaultSubtotal.
func (pf *PivotField) DefaultSubtotal() bool {
	attr, found := pf.GetAttribute(
		"defaultSubtotal",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetDefaultSubtotal sets whether to show default subtotal.
// Attribute: defaultSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetDefaultSubtotal(
	value bool,
) {
	if value {
		pf.RemoveAttribute(
			"defaultSubtotal",
			"",
		) // true is default
	} else {
		pf.SetAttribute(
			openxml.NewAttribute("", "defaultSubtotal", "", attrValueFalse),
		)
	}
}

// SumSubtotal returns whether to show sum subtotal. Attribute: sumSubtotal.
func (pf *PivotField) SumSubtotal() bool {
	attr, found := pf.GetAttribute(
		"sumSubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSumSubtotal sets whether to show sum subtotal. Attribute: sumSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetSumSubtotal(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"sumSubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("sumSubtotal", "")
	}
}

// CountASubtotal returns whether to show count all subtotal.
// Attribute: countASubtotal.
func (pf *PivotField) CountASubtotal() bool {
	attr, found := pf.GetAttribute(
		"countASubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCountASubtotal sets whether to show count all subtotal.
// Attribute: countASubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetCountASubtotal(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"countASubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("countASubtotal", "")
	}
}

// AvgSubtotal returns whether to show average subtotal. Attribute: avgSubtotal.
func (pf *PivotField) AvgSubtotal() bool {
	attr, found := pf.GetAttribute(
		"avgSubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAvgSubtotal sets whether to show average subtotal. Attribute: avgSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetAvgSubtotal(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"avgSubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("avgSubtotal", "")
	}
}

// MaxSubtotal returns whether to show max subtotal. Attribute: maxSubtotal.
func (pf *PivotField) MaxSubtotal() bool {
	attr, found := pf.GetAttribute(
		"maxSubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetMaxSubtotal sets whether to show max subtotal. Attribute: maxSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetMaxSubtotal(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"maxSubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("maxSubtotal", "")
	}
}

// MinSubtotal returns whether to show min subtotal. Attribute: minSubtotal.
func (pf *PivotField) MinSubtotal() bool {
	attr, found := pf.GetAttribute(
		"minSubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetMinSubtotal sets whether to show min subtotal. Attribute: minSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetMinSubtotal(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"minSubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("minSubtotal", "")
	}
}

// ProductSubtotal returns whether to show product subtotal.
// Attribute: productSubtotal.
func (pf *PivotField) ProductSubtotal() bool {
	attr, found := pf.GetAttribute(
		"productSubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetProductSubtotal sets whether to show product subtotal.
// Attribute: productSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetProductSubtotal(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"productSubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("productSubtotal", "")
	}
}

// CountSubtotal returns whether to show count subtotal.
// Attribute: countSubtotal.
func (pf *PivotField) CountSubtotal() bool {
	attr, found := pf.GetAttribute(
		"countSubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCountSubtotal sets whether to show count subtotal.
// Attribute: countSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetCountSubtotal(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"countSubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("countSubtotal", "")
	}
}

// StdDevSubtotal returns whether to show standard deviation subtotal.
// Attribute: stdDevSubtotal.
func (pf *PivotField) StdDevSubtotal() bool {
	attr, found := pf.GetAttribute(
		"stdDevSubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetStdDevSubtotal sets whether to show standard deviation subtotal.
// Attribute: stdDevSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetStdDevSubtotal(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"stdDevSubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("stdDevSubtotal", "")
	}
}

// StdDevPSubtotal returns whether to show population standard
// deviation subtotal. Attribute: stdDevPSubtotal.
func (pf *PivotField) StdDevPSubtotal() bool {
	attr, found := pf.GetAttribute(
		"stdDevPSubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetStdDevPSubtotal sets whether to show population standard deviation subtotal.
// Attribute: stdDevPSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetStdDevPSubtotal(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"stdDevPSubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("stdDevPSubtotal", "")
	}
}

// VarSubtotal returns whether to show variance subtotal.
// Attribute: varSubtotal.
func (pf *PivotField) VarSubtotal() bool {
	attr, found := pf.GetAttribute(
		"varSubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetVarSubtotal sets whether to show variance subtotal.
// Attribute: varSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetVarSubtotal(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"varSubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("varSubtotal", "")
	}
}

// VarPSubtotal returns whether to show population variance subtotal.
// Attribute: varPSubtotal.
func (pf *PivotField) VarPSubtotal() bool {
	attr, found := pf.GetAttribute(
		"varPSubtotal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetVarPSubtotal sets whether to show population variance subtotal.
// Attribute: varPSubtotal.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetVarPSubtotal(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"varPSubtotal",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("varPSubtotal", "")
	}
}

// ShowPropCell returns whether to show property in cell.
// Attribute: showPropCell.
func (pf *PivotField) ShowPropCell() bool {
	attr, found := pf.GetAttribute(
		"showPropCell",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowPropCell sets whether to show property in cell.
// Attribute: showPropCell.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetShowPropCell(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"showPropCell",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("showPropCell", "")
	}
}

// ShowPropTip returns whether to show property in tooltip.
// Attribute: showPropTip.
func (pf *PivotField) ShowPropTip() bool {
	attr, found := pf.GetAttribute(
		"showPropTip",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowPropTip sets whether to show property in tooltip.
// Attribute: showPropTip.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetShowPropTip(value bool) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"showPropTip",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("showPropTip", "")
	}
}

// ShowPropAsCaption returns whether to show property as caption.
// Attribute: showPropAsCaption.
func (pf *PivotField) ShowPropAsCaption() bool {
	attr, found := pf.GetAttribute(
		"showPropAsCaption",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowPropAsCaption sets whether to show property as caption.
// Attribute: showPropAsCaption.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetShowPropAsCaption(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"showPropAsCaption",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("showPropAsCaption", "")
	}
}

// DefaultAttributeDrillState returns whether to default attribute drill state.
// Attribute: defaultAttributeDrillState.
func (pf *PivotField) DefaultAttributeDrillState() bool {
	attr, found := pf.GetAttribute(
		"defaultAttributeDrillState",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDefaultAttributeDrillState sets whether to default attribute drill state.
// Attribute: defaultAttributeDrillState.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pf *PivotField) SetDefaultAttributeDrillState(
	value bool,
) {
	if value {
		pf.SetAttribute(
			openxml.NewAttribute(
				"",
				"defaultAttributeDrillState",
				"",
				attrValueTrue,
			),
		)
	} else {
		pf.RemoveAttribute("defaultAttributeDrillState", "")
	}
}

// Items returns the Items child element, or nil if not present.
func (pf *PivotField) Items() *Items {
	elem := pf.GetElement("items", NamespaceSML)
	if elem == nil {
		return nil
	}
	if items, ok := elem.(*Items); ok {
		return items
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Items{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateItems returns the Items child element, creating it if needed.
func (pf *PivotField) GetOrCreateItems() *Items {
	items := pf.Items()
	if items != nil {
		return items
	}
	items = NewItems()
	pf.AppendChild(items)

	return items
}

// Clone creates a deep copy of this PivotField element.
func (pf *PivotField) Clone() openxml.Element {
	cloned := pf.CompositeElementBase.Clone()

	return &PivotField{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PivotField element.
func (pf *PivotField) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pf.CompositeElementBase.CloneNode(
		deep,
	)

	return &PivotField{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Items represents the items collection element (x:items).
type Items struct {
	*openxml.CompositeElementBase
}

// NewItems creates a new Items element.
func NewItems() *Items {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"items",
		PrefixDefault,
	)

	return &Items{CompositeElementBase: elem}
}

// Count returns the count of items. Attribute: count.
func (i *Items) Count() uint32 {
	attr, found := i.GetAttribute("count", "")
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

// SetCount sets the count of items. Attribute: count.
func (i *Items) SetCount(count uint32) {
	i.SetAttribute(
		openxml.NewAttribute(
			"",
			"count", //nolint:revive // add-constant: attribute name
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// Items returns an iterator over all Item elements.
func (i *Items) Items() iter.Seq[*Item] {
	return func(yield func(*Item) bool) {
		for child := range i.Children() {
			if child.LocalName() != "item" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var item *Item
			if it, ok := child.(*Item); ok {
				item = it
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				item = &Item{LeafElementBase: leaf}
			}
			if item != nil && !yield(item) {
				return
			}
		}
	}
}

// AddItem adds a new Item element.
func (i *Items) AddItem() *Item {
	item := NewItem()
	i.AppendChild(item)

	return item
}

// Clone creates a deep copy of this Items element.
func (i *Items) Clone() openxml.Element {
	cloned := i.CompositeElementBase.Clone()

	return &Items{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Items element.
func (i *Items) CloneNode(
	deep bool,
) openxml.Element {
	cloned := i.CompositeElementBase.CloneNode(
		deep,
	)

	return &Items{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Item represents a single pivot field item element (x:item).
type Item struct {
	*openxml.LeafElementBase
}

// NewItem creates a new Item element.
func NewItem() *Item {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"item",
		PrefixDefault,
	)

	return &Item{LeafElementBase: elem}
}

// N returns the item name. Attribute: n.
func (i *Item) N() string {
	attr, found := i.GetAttribute("n", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetN sets the item name. Attribute: n.
func (i *Item) SetN(name string) {
	if name == "" {
		i.RemoveAttribute("n", "")

		return
	}
	i.SetAttribute(
		openxml.NewAttribute("", "n", "", name),
	)
}

// T returns the item type. Attribute: t.
func (i *Item) T() string {
	attr, found := i.GetAttribute("t", "")
	if !found {
		return "data" // Default is data
	}

	return attr.Value()
}

// SetT sets the item type. Attribute: t.
// Valid values: data, default, sum, countA, avg, max, min, product, count,
// stdDev, stdDevP, var, varP, grand, blank
func (i *Item) SetT(t string) {
	if t == "" || t == "data" {
		i.RemoveAttribute(
			"t",
			"",
		) // data is default

		return
	}
	i.SetAttribute(
		openxml.NewAttribute("", "t", "", t),
	)
}

// H returns whether item is hidden. Attribute: h.
func (i *Item) H() bool {
	attr, found := i.GetAttribute("h", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetH sets whether item is hidden. Attribute: h.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (i *Item) SetH(value bool) {
	if value {
		i.SetAttribute(
			openxml.NewAttribute(
				"",
				"h",
				"",
				attrValueTrue,
			),
		)
	} else {
		i.RemoveAttribute("h", "")
	}
}

// S returns whether item detail is hidden. Attribute: s.
func (i *Item) S() bool {
	attr, found := i.GetAttribute("s", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetS sets whether item detail is hidden. Attribute: s.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (i *Item) SetS(value bool) {
	if value {
		i.SetAttribute(
			openxml.NewAttribute(
				"",
				"s",
				"",
				attrValueTrue,
			),
		)
	} else {
		i.RemoveAttribute("s", "")
	}
}

// Sd returns whether show details. Attribute: sd.
func (i *Item) Sd() bool {
	attr, found := i.GetAttribute("sd", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetSd sets whether show details. Attribute: sd.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (i *Item) SetSd(value bool) {
	if value {
		i.RemoveAttribute(
			"sd",
			"",
		) // true is default
	} else {
		i.SetAttribute(
			openxml.NewAttribute("", "sd", "", attrValueFalse),
		)
	}
}

// F returns whether item has calculated member formula. Attribute: f.
func (i *Item) F() bool {
	attr, found := i.GetAttribute("f", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetF sets whether item has calculated member formula. Attribute: f.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (i *Item) SetF(value bool) {
	if value {
		i.SetAttribute(
			openxml.NewAttribute(
				"",
				"f",
				"",
				attrValueTrue,
			),
		)
	} else {
		i.RemoveAttribute("f", "")
	}
}

// M returns whether item is missing. Attribute: m.
func (i *Item) M() bool {
	attr, found := i.GetAttribute("m", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetM sets whether item is missing. Attribute: m.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (i *Item) SetM(value bool) {
	if value {
		i.SetAttribute(
			openxml.NewAttribute(
				"",
				"m",
				"",
				attrValueTrue,
			),
		)
	} else {
		i.RemoveAttribute("m", "")
	}
}

// C returns whether item has a child. Attribute: c.
func (i *Item) C() bool {
	attr, found := i.GetAttribute("c", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetC sets whether item has a child. Attribute: c.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (i *Item) SetC(value bool) {
	if value {
		i.SetAttribute(
			openxml.NewAttribute(
				"",
				"c",
				"",
				attrValueTrue,
			),
		)
	} else {
		i.RemoveAttribute("c", "")
	}
}

// X returns the shared item index. Attribute: x.
func (i *Item) X() (uint32, bool) {
	attr, found := i.GetAttribute("x", "")
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val), true
}

// SetX sets the shared item index. Attribute: x.
func (i *Item) SetX(index uint32) {
	i.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatUint(
				uint64(index),
				parseBase10,
			),
		),
	)
}

// ClearX removes the shared item index attribute.
func (i *Item) ClearX() {
	i.RemoveAttribute("x", "")
}

// D returns whether item is expanded. Attribute: d.
func (i *Item) D() bool {
	attr, found := i.GetAttribute("d", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetD sets whether item is expanded. Attribute: d.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (i *Item) SetD(value bool) {
	if value {
		i.SetAttribute(
			openxml.NewAttribute(
				"",
				"d",
				"",
				attrValueTrue,
			),
		)
	} else {
		i.RemoveAttribute("d", "")
	}
}

// E returns whether item is drilled. Attribute: e.
func (i *Item) E() bool {
	attr, found := i.GetAttribute("e", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetE sets whether item is drilled. Attribute: e.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (i *Item) SetE(value bool) {
	if value {
		i.RemoveAttribute(
			"e",
			"",
		) // true is default
	} else {
		i.SetAttribute(
			openxml.NewAttribute("", "e", "", attrValueFalse),
		)
	}
}

// Clone creates a deep copy of this Item element.
func (i *Item) Clone() openxml.Element {
	cloned := i.LeafElementBase.Clone()

	return &Item{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Item element.
func (i *Item) CloneNode(
	deep bool,
) openxml.Element {
	cloned := i.LeafElementBase.CloneNode(deep)

	return &Item{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// RowFields represents the row fields collection element (x:rowFields).
type RowFields struct {
	*openxml.CompositeElementBase
}

// NewRowFields creates a new RowFields element.
func NewRowFields() *RowFields {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"rowFields",
		PrefixDefault,
	)

	return &RowFields{CompositeElementBase: elem}
}

// Count returns the count of row fields. Attribute: count.
func (rf *RowFields) Count() uint32 {
	attr, found := rf.GetAttribute("count", "")
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

// SetCount sets the count of row fields. Attribute: count.
func (rf *RowFields) SetCount(count uint32) {
	rf.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// Fields returns an iterator over all Field elements.
func (rf *RowFields) Fields() iter.Seq[*Field] {
	return func(yield func(*Field) bool) {
		for child := range rf.Children() {
			if child.LocalName() != "field" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var field *Field
			if f, ok := child.(*Field); ok {
				field = f
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				field = &Field{LeafElementBase: leaf}
			}
			if field != nil && !yield(field) {
				return
			}
		}
	}
}

// AddField adds a new Field element.
func (rf *RowFields) AddField(x int32) *Field {
	field := NewField()
	field.SetX(x)
	rf.AppendChild(field)

	return field
}

// Clone creates a deep copy of this RowFields element.
func (rf *RowFields) Clone() openxml.Element {
	cloned := rf.CompositeElementBase.Clone()

	return &RowFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this RowFields element.
func (rf *RowFields) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rf.CompositeElementBase.CloneNode(
		deep,
	)

	return &RowFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ColFields represents the column fields collection element (x:colFields).
type ColFields struct {
	*openxml.CompositeElementBase
}

// NewColFields creates a new ColFields element.
func NewColFields() *ColFields {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"colFields",
		PrefixDefault,
	)

	return &ColFields{CompositeElementBase: elem}
}

// Count returns the count of column fields. Attribute: count.
func (cf *ColFields) Count() uint32 {
	attr, found := cf.GetAttribute("count", "")
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

// SetCount sets the count of column fields. Attribute: count.
func (cf *ColFields) SetCount(count uint32) {
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// Fields returns an iterator over all Field elements.
func (cf *ColFields) Fields() iter.Seq[*Field] {
	return func(yield func(*Field) bool) {
		for child := range cf.Children() {
			if child.LocalName() != "field" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var field *Field
			if f, ok := child.(*Field); ok {
				field = f
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				field = &Field{LeafElementBase: leaf}
			}
			if field != nil && !yield(field) {
				return
			}
		}
	}
}

// AddField adds a new Field element.
func (cf *ColFields) AddField(x int32) *Field {
	field := NewField()
	field.SetX(x)
	cf.AppendChild(field)

	return field
}

// Clone creates a deep copy of this ColFields element.
func (cf *ColFields) Clone() openxml.Element {
	cloned := cf.CompositeElementBase.Clone()

	return &ColFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ColFields element.
func (cf *ColFields) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cf.CompositeElementBase.CloneNode(
		deep,
	)

	return &ColFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Field represents a field reference element (x:field).
type Field struct {
	*openxml.LeafElementBase
}

// NewField creates a new Field element.
func NewField() *Field {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"field",
		PrefixDefault,
	)

	return &Field{LeafElementBase: elem}
}

// X returns the field index. Attribute: x.
// -2 means "data" field, otherwise refers to pivot field index.
func (f *Field) X() int32 {
	attr, found := f.GetAttribute(
		"x", //nolint:revive // add-constant
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return int32(val)
}

// SetX sets the field index. Attribute: x.
func (f *Field) SetX(x int32) {
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(
				int64(x),
				parseBase10,
			),
		),
	)
}

// Clone creates a deep copy of this Field element.
func (f *Field) Clone() openxml.Element {
	cloned := f.LeafElementBase.Clone()

	return &Field{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Field element.
func (f *Field) CloneNode(
	deep bool,
) openxml.Element {
	cloned := f.LeafElementBase.CloneNode(deep)

	return &Field{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// PageFields represents the page fields collection element (x:pageFields).
type PageFields struct {
	*openxml.CompositeElementBase
}

// NewPageFields creates a new PageFields element.
func NewPageFields() *PageFields {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"pageFields",
		PrefixDefault,
	)

	return &PageFields{CompositeElementBase: elem}
}

// Count returns the count of page fields. Attribute: count.
func (pf *PageFields) Count() uint32 {
	attr, found := pf.GetAttribute("count", "")
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

// SetCount sets the count of page fields. Attribute: count.
func (pf *PageFields) SetCount(count uint32) {
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// PageFields returns an iterator over all PageField elements.
func (pf *PageFields) PageFields() iter.Seq[*PageField] {
	return func(yield func(*PageField) bool) {
		for child := range pf.Children() {
			if child.LocalName() != "pageField" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var field *PageField
			if f, ok := child.(*PageField); ok {
				field = f
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				field = &PageField{LeafElementBase: leaf}
			}
			if field != nil && !yield(field) {
				return
			}
		}
	}
}

// AddPageField adds a new PageField element.
func (pf *PageFields) AddPageField(
	fld int32,
) *PageField {
	field := NewPageField()
	field.SetFld(fld)
	pf.AppendChild(field)

	return field
}

// Clone creates a deep copy of this PageFields element.
func (pf *PageFields) Clone() openxml.Element {
	cloned := pf.CompositeElementBase.Clone()

	return &PageFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PageFields element.
func (pf *PageFields) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pf.CompositeElementBase.CloneNode(
		deep,
	)

	return &PageFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// PageField represents a single page field element (x:pageField).
type PageField struct {
	*openxml.LeafElementBase
}

// NewPageField creates a new PageField element.
func NewPageField() *PageField {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"pageField",
		PrefixDefault,
	)

	return &PageField{LeafElementBase: elem}
}

// Fld returns the field index. Attribute: fld.
func (pf *PageField) Fld() int32 {
	attr, found := pf.GetAttribute("fld", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return int32(val)
}

// SetFld sets the field index. Attribute: fld.
func (pf *PageField) SetFld(fld int32) {
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"fld",
			"",
			strconv.FormatInt(
				int64(fld),
				parseBase10,
			),
		),
	)
}

// Item returns the item index. Attribute: item.
func (pf *PageField) Item() (uint32, bool) {
	attr, found := pf.GetAttribute("item", "")
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val), true
}

// SetItem sets the item index. Attribute: item.
func (pf *PageField) SetItem(item uint32) {
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"item", //nolint:revive // add-constant: attribute name
			"",
			strconv.FormatUint(
				uint64(item),
				parseBase10,
			),
		),
	)
}

// ClearItem removes the item attribute.
func (pf *PageField) ClearItem() {
	pf.RemoveAttribute("item", "")
}

// Hier returns the hierarchy index. Attribute: hier.
func (pf *PageField) Hier() int32 {
	attr, found := pf.GetAttribute("hier", "")
	if !found {
		return -1 // Default is -1
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return int32(val)
}

// SetHier sets the hierarchy index. Attribute: hier.
func (pf *PageField) SetHier(hier int32) {
	if hier == -1 {
		pf.RemoveAttribute(
			"hier",
			"",
		) // -1 is default

		return
	}
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"hier",
			"",
			strconv.FormatInt(
				int64(hier),
				parseBase10,
			),
		),
	)
}

// Name returns the page field name. Attribute: name.
func (pf *PageField) Name() string {
	attr, found := pf.GetAttribute(
		"name", //nolint:revive // add-constant
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the page field name. Attribute: name.
func (pf *PageField) SetName(name string) {
	if name == "" {
		pf.RemoveAttribute("name", "")

		return
	}
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Cap returns the page field caption. Attribute: cap.
func (pf *PageField) Cap() string {
	attr, found := pf.GetAttribute("cap", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCap sets the page field caption. Attribute: cap.
func (pf *PageField) SetCap(caption string) {
	if caption == "" {
		pf.RemoveAttribute("cap", "")

		return
	}
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"cap",
			"",
			caption,
		),
	)
}

// Clone creates a deep copy of this PageField element.
func (pf *PageField) Clone() openxml.Element {
	cloned := pf.LeafElementBase.Clone()

	return &PageField{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this PageField element.
func (pf *PageField) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pf.LeafElementBase.CloneNode(deep)

	return &PageField{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// DataFields represents the data fields collection element (x:dataFields).
type DataFields struct {
	*openxml.CompositeElementBase
}

// NewDataFields creates a new DataFields element.
func NewDataFields() *DataFields {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"dataFields",
		PrefixDefault,
	)

	return &DataFields{CompositeElementBase: elem}
}

// Count returns the count of data fields. Attribute: count.
func (df *DataFields) Count() uint32 {
	attr, found := df.GetAttribute("count", "")
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

// SetCount sets the count of data fields. Attribute: count.
func (df *DataFields) SetCount(count uint32) {
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// DataFields returns an iterator over all DataField elements.
func (df *DataFields) DataFields() iter.Seq[*DataField] {
	return func(yield func(*DataField) bool) {
		for child := range df.Children() {
			if child.LocalName() != "dataField" || //nolint:revive // add-constant: element name
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var field *DataField
			if f, ok := child.(*DataField); ok {
				field = f
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				field = &DataField{LeafElementBase: leaf}
			}
			if field != nil && !yield(field) {
				return
			}
		}
	}
}

// AddDataField adds a new DataField element.
func (df *DataFields) AddDataField(
	fld uint32,
) *DataField {
	field := NewDataField()
	field.SetFld(fld)
	df.AppendChild(field)

	return field
}

// Clone creates a deep copy of this DataFields element.
func (df *DataFields) Clone() openxml.Element {
	cloned := df.CompositeElementBase.Clone()

	return &DataFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DataFields element.
func (df *DataFields) CloneNode(
	deep bool,
) openxml.Element {
	cloned := df.CompositeElementBase.CloneNode(
		deep,
	)

	return &DataFields{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// DataField represents a single data field element (x:dataField).
type DataField struct {
	*openxml.LeafElementBase
}

// NewDataField creates a new DataField element.
func NewDataField() *DataField {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"dataField",
		PrefixDefault,
	)

	return &DataField{LeafElementBase: elem}
}

// Name returns the data field name. Attribute: name.
func (df *DataField) Name() string {
	attr, found := df.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the data field name. Attribute: name.
func (df *DataField) SetName(name string) {
	if name == "" {
		df.RemoveAttribute("name", "")

		return
	}
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Fld returns the field index. Attribute: fld.
func (df *DataField) Fld() uint32 {
	attr, found := df.GetAttribute("fld", "")
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

// SetFld sets the field index. Attribute: fld.
func (df *DataField) SetFld(fld uint32) {
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"fld", //nolint:revive // add-constant: attribute name
			"",
			strconv.FormatUint(
				uint64(fld),
				parseBase10,
			),
		),
	)
}

// Subtotal returns the subtotal function. Attribute: subtotal.
func (df *DataField) Subtotal() string {
	attr, found := df.GetAttribute("subtotal", "")
	if !found {
		return "sum" // Default is sum
	}

	return attr.Value()
}

// SetSubtotal sets the subtotal function. Attribute: subtotal.
// Valid values: sum, count, average, max, min, product, countNums, stdDev,
// stdDevp, var, varp
func (df *DataField) SetSubtotal(
	subtotal string,
) {
	if subtotal == "" || subtotal == "sum" {
		df.RemoveAttribute(
			"subtotal",
			"",
		) // sum is default

		return
	}
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"subtotal",
			"",
			subtotal,
		),
	)
}

// ShowDataAs returns how data is shown. Attribute: showDataAs.
func (df *DataField) ShowDataAs() string {
	attr, found := df.GetAttribute(
		"showDataAs",
		"",
	)
	if !found {
		return "normal" // Default is normal
	}

	return attr.Value()
}

// SetShowDataAs sets how data is shown. Attribute: showDataAs.
// Valid values: normal, difference, percent, percentDiff, runTotal,
// percentOfRow, percentOfCol, percentOfTotal, index
func (df *DataField) SetShowDataAs(
	showDataAs string,
) {
	if showDataAs == "" ||
		showDataAs == "normal" {
		df.RemoveAttribute(
			"showDataAs",
			"",
		) // normal is default

		return
	}
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"showDataAs",
			"",
			showDataAs,
		),
	)
}

// BaseField returns the base field index. Attribute: baseField.
func (df *DataField) BaseField() int32 {
	attr, found := df.GetAttribute(
		"baseField",
		"",
	)
	if !found {
		return -1 // Default is -1
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return int32(val)
}

// SetBaseField sets the base field index. Attribute: baseField.
func (df *DataField) SetBaseField(field int32) {
	if field == -1 {
		df.RemoveAttribute(
			"baseField",
			"",
		) // -1 is default

		return
	}
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"baseField",
			"",
			strconv.FormatInt(
				int64(field),
				parseBase10,
			),
		),
	)
}

// BaseItem returns the base item index. Attribute: baseItem.
func (df *DataField) BaseItem() uint32 {
	attr, found := df.GetAttribute("baseItem", "")
	if !found {
		return 1048832 //nolint:revive // add-constant: default value per OOXML spec
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetBaseItem sets the base item index. Attribute: baseItem.
func (df *DataField) SetBaseItem(item uint32) {
	if item == 1048832 { //nolint:revive // add-constant: default value per OOXML spec
		df.RemoveAttribute(
			"baseItem",
			"",
		) // 1048832 is default

		return
	}
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"baseItem",
			"",
			strconv.FormatUint(
				uint64(item),
				parseBase10,
			),
		),
	)
}

// NumFmtId returns the number format ID. Attribute: numFmtId.
func (df *DataField) NumFmtId() (uint32, bool) {
	attr, found := df.GetAttribute(
		"numFmtId", //nolint:revive // add-constant
		"",
	)
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val), true
}

// SetNumFmtId sets the number format ID. Attribute: numFmtId.
func (df *DataField) SetNumFmtId(id uint32) {
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"numFmtId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10,
			),
		),
	)
}

// ClearNumFmtId removes the number format ID attribute.
func (df *DataField) ClearNumFmtId() {
	df.RemoveAttribute("numFmtId", "")
}

// Clone creates a deep copy of this DataField element.
func (df *DataField) Clone() openxml.Element {
	cloned := df.LeafElementBase.Clone()

	return &DataField{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this DataField element.
func (df *DataField) CloneNode(
	deep bool,
) openxml.Element {
	cloned := df.LeafElementBase.CloneNode(deep)

	return &DataField{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
