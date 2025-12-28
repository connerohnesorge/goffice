package elements

//revive:disable:file-length-limit many table style types

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// TableStyleElementType represents the type of table style element.
type TableStyleElementType string

const (
	// TableStyleElementWholeTable applies to the whole table.
	TableStyleElementWholeTable TableStyleElementType = "wholeTable"
	// TableStyleElementHeaderRow applies to the header row.
	TableStyleElementHeaderRow TableStyleElementType = "headerRow"
	// TableStyleElementTotalRow applies to the total row.
	TableStyleElementTotalRow TableStyleElementType = "totalRow"
	// TableStyleElementFirstColumn applies to the first column.
	TableStyleElementFirstColumn TableStyleElementType = "firstColumn"
	// TableStyleElementLastColumn applies to the last column.
	TableStyleElementLastColumn TableStyleElementType = "lastColumn"
	// TableStyleElementFirstRowStripe applies to the first row stripe.
	TableStyleElementFirstRowStripe TableStyleElementType = "firstRowStripe"
	// TableStyleElementSecondRowStripe applies to the second row stripe.
	TableStyleElementSecondRowStripe TableStyleElementType = "secondRowStripe"
	// TableStyleElementFirstColumnStripe applies to the first column stripe.
	//nolint:revive // line-length-limit
	TableStyleElementFirstColumnStripe TableStyleElementType = "firstColumnStripe"
	// TableStyleElementSecondColumnStripe applies to the second column stripe.
	//nolint:revive // line-length-limit
	TableStyleElementSecondColumnStripe TableStyleElementType = "secondColumnStripe"
	// TableStyleElementFirstHeaderCell applies to the first header cell.
	TableStyleElementFirstHeaderCell TableStyleElementType = "firstHeaderCell"
	// TableStyleElementLastHeaderCell applies to the last header cell.
	TableStyleElementLastHeaderCell TableStyleElementType = "lastHeaderCell"
	// TableStyleElementFirstTotalCell applies to the first total cell.
	TableStyleElementFirstTotalCell TableStyleElementType = "firstTotalCell"
	// TableStyleElementLastTotalCell applies to the last total cell.
	TableStyleElementLastTotalCell TableStyleElementType = "lastTotalCell"
	// TableStyleElementFirstSubtotalColumn applies to the first subtotal column.
	//nolint:revive // line-length-limit
	TableStyleElementFirstSubtotalColumn TableStyleElementType = "firstSubtotalColumn"
	// TableStyleElementSecondSubtotalColumn applies to the second subtotal
	// column.
	//nolint:revive // line-length-limit
	TableStyleElementSecondSubtotalColumn TableStyleElementType = "secondSubtotalColumn"
	// TableStyleElementThirdSubtotalColumn applies to the third subtotal column.
	//nolint:revive // line-length-limit
	TableStyleElementThirdSubtotalColumn TableStyleElementType = "thirdSubtotalColumn"
	// TableStyleElementFirstSubtotalRow applies to the first subtotal row.
	TableStyleElementFirstSubtotalRow TableStyleElementType = "firstSubtotalRow"
	// TableStyleElementSecondSubtotalRow applies to the second subtotal row.
	//nolint:revive // line-length-limit
	TableStyleElementSecondSubtotalRow TableStyleElementType = "secondSubtotalRow"
	// TableStyleElementThirdSubtotalRow applies to the third subtotal row.
	TableStyleElementThirdSubtotalRow TableStyleElementType = "thirdSubtotalRow"
	// TableStyleElementBlankRow applies to blank rows.
	TableStyleElementBlankRow TableStyleElementType = "blankRow"
	// TableStyleElementFirstColumnSubheading applies to first column subheading.
	//nolint:revive // line-length-limit
	TableStyleElementFirstColumnSubheading TableStyleElementType = "firstColumnSubheading"
	// TableStyleElementSecondColumnSubheading applies to second column
	// subheading.
	//nolint:revive // line-length-limit
	TableStyleElementSecondColumnSubheading TableStyleElementType = "secondColumnSubheading"
	// TableStyleElementThirdColumnSubheading applies to third column subheading.
	//nolint:revive // line-length-limit
	TableStyleElementThirdColumnSubheading TableStyleElementType = "thirdColumnSubheading"
	// TableStyleElementFirstRowSubheading applies to the first row subheading.
	//nolint:revive // line-length-limit
	TableStyleElementFirstRowSubheading TableStyleElementType = "firstRowSubheading"
	// TableStyleElementSecondRowSubheading applies to the second row subheading.
	//nolint:revive // line-length-limit
	TableStyleElementSecondRowSubheading TableStyleElementType = "secondRowSubheading"
	// TableStyleElementThirdRowSubheading applies to the third row subheading.
	//nolint:revive // line-length-limit
	TableStyleElementThirdRowSubheading TableStyleElementType = "thirdRowSubheading"
	// TableStyleElementPageFieldLabels applies to page field labels.
	TableStyleElementPageFieldLabels TableStyleElementType = "pageFieldLabels"
	// TableStyleElementPageFieldValues applies to page field values.
	TableStyleElementPageFieldValues TableStyleElementType = "pageFieldValues"
)

// Built-in table style names - Light styles (21 styles).
const (
	TableStyleLight1  = "TableStyleLight1"
	TableStyleLight2  = "TableStyleLight2"
	TableStyleLight3  = "TableStyleLight3"
	TableStyleLight4  = "TableStyleLight4"
	TableStyleLight5  = "TableStyleLight5"
	TableStyleLight6  = "TableStyleLight6"
	TableStyleLight7  = "TableStyleLight7"
	TableStyleLight8  = "TableStyleLight8"
	TableStyleLight9  = "TableStyleLight9"
	TableStyleLight10 = "TableStyleLight10"
	TableStyleLight11 = "TableStyleLight11"
	TableStyleLight12 = "TableStyleLight12"
	TableStyleLight13 = "TableStyleLight13"
	TableStyleLight14 = "TableStyleLight14"
	TableStyleLight15 = "TableStyleLight15"
	TableStyleLight16 = "TableStyleLight16"
	TableStyleLight17 = "TableStyleLight17"
	TableStyleLight18 = "TableStyleLight18"
	TableStyleLight19 = "TableStyleLight19"
	TableStyleLight20 = "TableStyleLight20"
	TableStyleLight21 = "TableStyleLight21"
)

// Built-in table style names - Medium styles (28 styles).
const (
	TableStyleMedium1  = "TableStyleMedium1"
	TableStyleMedium2  = "TableStyleMedium2"
	TableStyleMedium3  = "TableStyleMedium3"
	TableStyleMedium4  = "TableStyleMedium4"
	TableStyleMedium5  = "TableStyleMedium5"
	TableStyleMedium6  = "TableStyleMedium6"
	TableStyleMedium7  = "TableStyleMedium7"
	TableStyleMedium8  = "TableStyleMedium8"
	TableStyleMedium9  = "TableStyleMedium9"
	TableStyleMedium10 = "TableStyleMedium10"
	TableStyleMedium11 = "TableStyleMedium11"
	TableStyleMedium12 = "TableStyleMedium12"
	TableStyleMedium13 = "TableStyleMedium13"
	TableStyleMedium14 = "TableStyleMedium14"
	TableStyleMedium15 = "TableStyleMedium15"
	TableStyleMedium16 = "TableStyleMedium16"
	TableStyleMedium17 = "TableStyleMedium17"
	TableStyleMedium18 = "TableStyleMedium18"
	TableStyleMedium19 = "TableStyleMedium19"
	TableStyleMedium20 = "TableStyleMedium20"
	TableStyleMedium21 = "TableStyleMedium21"
	TableStyleMedium22 = "TableStyleMedium22"
	TableStyleMedium23 = "TableStyleMedium23"
	TableStyleMedium24 = "TableStyleMedium24"
	TableStyleMedium25 = "TableStyleMedium25"
	TableStyleMedium26 = "TableStyleMedium26"
	TableStyleMedium27 = "TableStyleMedium27"
	TableStyleMedium28 = "TableStyleMedium28"
)

// Built-in table style names - Dark styles (11 styles).
const (
	TableStyleDark1  = "TableStyleDark1"
	TableStyleDark2  = "TableStyleDark2"
	TableStyleDark3  = "TableStyleDark3"
	TableStyleDark4  = "TableStyleDark4"
	TableStyleDark5  = "TableStyleDark5"
	TableStyleDark6  = "TableStyleDark6"
	TableStyleDark7  = "TableStyleDark7"
	TableStyleDark8  = "TableStyleDark8"
	TableStyleDark9  = "TableStyleDark9"
	TableStyleDark10 = "TableStyleDark10"
	TableStyleDark11 = "TableStyleDark11"
)

// Built-in pivot table style names - Light styles (28 styles).
const (
	PivotStyleLight1  = "PivotStyleLight1"
	PivotStyleLight2  = "PivotStyleLight2"
	PivotStyleLight3  = "PivotStyleLight3"
	PivotStyleLight4  = "PivotStyleLight4"
	PivotStyleLight5  = "PivotStyleLight5"
	PivotStyleLight6  = "PivotStyleLight6"
	PivotStyleLight7  = "PivotStyleLight7"
	PivotStyleLight8  = "PivotStyleLight8"
	PivotStyleLight9  = "PivotStyleLight9"
	PivotStyleLight10 = "PivotStyleLight10"
	PivotStyleLight11 = "PivotStyleLight11"
	PivotStyleLight12 = "PivotStyleLight12"
	PivotStyleLight13 = "PivotStyleLight13"
	PivotStyleLight14 = "PivotStyleLight14"
	PivotStyleLight15 = "PivotStyleLight15"
	PivotStyleLight16 = "PivotStyleLight16"
	PivotStyleLight17 = "PivotStyleLight17"
	PivotStyleLight18 = "PivotStyleLight18"
	PivotStyleLight19 = "PivotStyleLight19"
	PivotStyleLight20 = "PivotStyleLight20"
	PivotStyleLight21 = "PivotStyleLight21"
	PivotStyleLight22 = "PivotStyleLight22"
	PivotStyleLight23 = "PivotStyleLight23"
	PivotStyleLight24 = "PivotStyleLight24"
	PivotStyleLight25 = "PivotStyleLight25"
	PivotStyleLight26 = "PivotStyleLight26"
	PivotStyleLight27 = "PivotStyleLight27"
	PivotStyleLight28 = "PivotStyleLight28"
)

// Built-in pivot table style names - Medium styles (28 styles).
const (
	PivotStyleMedium1  = "PivotStyleMedium1"
	PivotStyleMedium2  = "PivotStyleMedium2"
	PivotStyleMedium3  = "PivotStyleMedium3"
	PivotStyleMedium4  = "PivotStyleMedium4"
	PivotStyleMedium5  = "PivotStyleMedium5"
	PivotStyleMedium6  = "PivotStyleMedium6"
	PivotStyleMedium7  = "PivotStyleMedium7"
	PivotStyleMedium8  = "PivotStyleMedium8"
	PivotStyleMedium9  = "PivotStyleMedium9"
	PivotStyleMedium10 = "PivotStyleMedium10"
	PivotStyleMedium11 = "PivotStyleMedium11"
	PivotStyleMedium12 = "PivotStyleMedium12"
	PivotStyleMedium13 = "PivotStyleMedium13"
	PivotStyleMedium14 = "PivotStyleMedium14"
	PivotStyleMedium15 = "PivotStyleMedium15"
	PivotStyleMedium16 = "PivotStyleMedium16"
	PivotStyleMedium17 = "PivotStyleMedium17"
	PivotStyleMedium18 = "PivotStyleMedium18"
	PivotStyleMedium19 = "PivotStyleMedium19"
	PivotStyleMedium20 = "PivotStyleMedium20"
	PivotStyleMedium21 = "PivotStyleMedium21"
	PivotStyleMedium22 = "PivotStyleMedium22"
	PivotStyleMedium23 = "PivotStyleMedium23"
	PivotStyleMedium24 = "PivotStyleMedium24"
	PivotStyleMedium25 = "PivotStyleMedium25"
	PivotStyleMedium26 = "PivotStyleMedium26"
	PivotStyleMedium27 = "PivotStyleMedium27"
	PivotStyleMedium28 = "PivotStyleMedium28"
)

// Built-in pivot table style names - Dark styles (28 styles).
const (
	PivotStyleDark1  = "PivotStyleDark1"
	PivotStyleDark2  = "PivotStyleDark2"
	PivotStyleDark3  = "PivotStyleDark3"
	PivotStyleDark4  = "PivotStyleDark4"
	PivotStyleDark5  = "PivotStyleDark5"
	PivotStyleDark6  = "PivotStyleDark6"
	PivotStyleDark7  = "PivotStyleDark7"
	PivotStyleDark8  = "PivotStyleDark8"
	PivotStyleDark9  = "PivotStyleDark9"
	PivotStyleDark10 = "PivotStyleDark10"
	PivotStyleDark11 = "PivotStyleDark11"
	PivotStyleDark12 = "PivotStyleDark12"
	PivotStyleDark13 = "PivotStyleDark13"
	PivotStyleDark14 = "PivotStyleDark14"
	PivotStyleDark15 = "PivotStyleDark15"
	PivotStyleDark16 = "PivotStyleDark16"
	PivotStyleDark17 = "PivotStyleDark17"
	PivotStyleDark18 = "PivotStyleDark18"
	PivotStyleDark19 = "PivotStyleDark19"
	PivotStyleDark20 = "PivotStyleDark20"
	PivotStyleDark21 = "PivotStyleDark21"
	PivotStyleDark22 = "PivotStyleDark22"
	PivotStyleDark23 = "PivotStyleDark23"
	PivotStyleDark24 = "PivotStyleDark24"
	PivotStyleDark25 = "PivotStyleDark25"
	PivotStyleDark26 = "PivotStyleDark26"
	PivotStyleDark27 = "PivotStyleDark27"
	PivotStyleDark28 = "PivotStyleDark28"
)

// Default style constants.
const (
	DefaultTableStyle = TableStyleMedium2
	DefaultPivotStyle = PivotStyleLight16
)

// TableStyles represents the table styles container element (x:tableStyles).
type TableStyles struct {
	*openxml.CompositeElementBase
}

// NewTableStyles creates a new TableStyles element.
func NewTableStyles() *TableStyles {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"tableStyles",
		PrefixDefault,
	)

	return &TableStyles{
		CompositeElementBase: elem,
	}
}

// Count returns the count attribute value.
func (t *TableStyles) Count() uint32 {
	attr, found := t.GetAttribute(
		attrNameCount,
		"",
	)
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

// SetCount sets the count attribute.
func (t *TableStyles) SetCount(count uint32) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			attrNameCount,
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// DefaultTableStyle returns the default table style name.
func (t *TableStyles) DefaultTableStyle() string {
	attr, found := t.GetAttribute(
		"defaultTableStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDefaultTableStyle sets the default table style name.
func (t *TableStyles) SetDefaultTableStyle(
	name string,
) {
	if name == "" {
		t.RemoveAttribute("defaultTableStyle", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"defaultTableStyle",
			"",
			name,
		),
	)
}

// DefaultPivotStyle returns the default pivot table style name.
func (t *TableStyles) DefaultPivotStyle() string {
	attr, found := t.GetAttribute(
		"defaultPivotStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDefaultPivotStyle sets the default pivot table style name.
func (t *TableStyles) SetDefaultPivotStyle(
	name string,
) {
	if name == "" {
		t.RemoveAttribute("defaultPivotStyle", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"defaultPivotStyle",
			"",
			name,
		),
	)
}

// TableStyles returns an iterator over all TableStyle elements.
func (t *TableStyles) TableStyles() iter.Seq[*TableStyle] {
	return func(yield func(*TableStyle) bool) {
		for child := range t.Children() {
			if child.LocalName() != "tableStyle" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var style *TableStyle
			switch v := child.(type) {
			case *TableStyle:
				style = v
			case *openxml.CompositeElementBase:
				style = &TableStyle{CompositeElementBase: v}
			}
			if style != nil && !yield(style) {
				return
			}
		}
	}
}

// GetTableStyle returns the table style with the given name,
// or nil if not found.
func (t *TableStyles) GetTableStyle(
	name string,
) *TableStyle {
	for style := range t.TableStyles() {
		if style.Name() == name {
			return style
		}
	}

	return nil
}

// AddTableStyle adds a new table style and returns it.
func (t *TableStyles) AddTableStyle(
	name string,
) *TableStyle {
	style := NewTableStyle()
	style.SetName(name)
	t.AppendChild(style)
	t.SetCount(t.Count() + 1)

	return style
}

// ItemCount returns the actual number of TableStyle children.
func (t *TableStyles) ItemCount() int {
	count := 0
	for range t.TableStyles() {
		count++
	}

	return count
}

// Clone creates a deep copy of this TableStyles element.
func (t *TableStyles) Clone() openxml.Element {
	cloned := t.CompositeElementBase.Clone()

	return &TableStyles{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this TableStyles element.
func (t *TableStyles) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.CompositeElementBase.CloneNode(
		deep,
	)

	return &TableStyles{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// TableStyle represents a table style element (x:tableStyle).
type TableStyle struct {
	*openxml.CompositeElementBase
}

// NewTableStyle creates a new TableStyle element.
func NewTableStyle() *TableStyle {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"tableStyle",
		PrefixDefault,
	)

	return &TableStyle{CompositeElementBase: elem}
}

// Name returns the style name.
func (t *TableStyle) Name() string {
	attr, found := t.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the style name.
func (t *TableStyle) SetName(name string) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Pivot returns whether this is a pivot table style.
func (t *TableStyle) Pivot() bool {
	attr, found := t.GetAttribute("pivot", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPivot sets whether this is a pivot table style.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *TableStyle) SetPivot(pivot bool) {
	if pivot {
		t.SetAttribute(
			openxml.NewAttribute(
				"",
				"pivot",
				"",
				attrValueTrue,
			),
		)
	} else {
		t.RemoveAttribute("pivot", "")
	}
}

// Table returns whether this is a table style (not pivot).
func (t *TableStyle) Table() bool {
	attr, found := t.GetAttribute("table", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetTable sets whether this is a table style (not pivot).
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *TableStyle) SetTable(table bool) {
	if table {
		t.RemoveAttribute(
			"table",
			"",
		) // true is default
	} else {
		t.SetAttribute(
			openxml.NewAttribute("", "table", "", attrValueFalse),
		)
	}
}

// Count returns the count of table style elements.
func (t *TableStyle) Count() uint32 {
	attr, found := t.GetAttribute(
		attrNameCount,
		"",
	)
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

// SetCount sets the count of table style elements.
func (t *TableStyle) SetCount(count uint32) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			attrNameCount,
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// TableStyleElements returns an iterator over all TableStyleElement elements.
func (t *TableStyle) TableStyleElements() iter.Seq[*TableStyleElement] {
	return func(yield func(*TableStyleElement) bool) {
		for child := range t.Children() {
			if child.LocalName() != "tableStyleElement" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var elem *TableStyleElement
			switch v := child.(type) {
			case *TableStyleElement:
				elem = v
			case *openxml.LeafElementBase:
				elem = &TableStyleElement{LeafElementBase: v}
			}
			if elem != nil && !yield(elem) {
				return
			}
		}
	}
}

// AddTableStyleElement adds a new table style element and returns it.
func (t *TableStyle) AddTableStyleElement(
	elementType TableStyleElementType,
	dxfId uint32,
) *TableStyleElement {
	elem := NewTableStyleElement()
	elem.SetType(elementType)
	elem.SetDxfId(dxfId)
	t.AppendChild(elem)
	t.SetCount(t.Count() + 1)

	return elem
}

// Clone creates a deep copy of this TableStyle element.
func (t *TableStyle) Clone() openxml.Element {
	cloned := t.CompositeElementBase.Clone()

	return &TableStyle{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this TableStyle element.
func (t *TableStyle) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.CompositeElementBase.CloneNode(
		deep,
	)

	return &TableStyle{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
