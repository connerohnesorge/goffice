package elements

//revive:disable:file-length-limit many table column properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// TotalsRowFunction represents the function used in a table totals row.
type TotalsRowFunction string

const (
	// TotalsRowFunctionNone indicates no function.
	TotalsRowFunctionNone TotalsRowFunction = "none"
	// TotalsRowFunctionSum indicates a SUM function.
	TotalsRowFunctionSum TotalsRowFunction = "sum"
	// TotalsRowFunctionMin indicates a MIN function.
	TotalsRowFunctionMin TotalsRowFunction = "min"
	// TotalsRowFunctionMax indicates a MAX function.
	TotalsRowFunctionMax TotalsRowFunction = "max"
	// TotalsRowFunctionAverage indicates an AVERAGE function.
	TotalsRowFunctionAverage TotalsRowFunction = "average"
	// TotalsRowFunctionCount indicates a COUNT function.
	TotalsRowFunctionCount TotalsRowFunction = "count"
	// TotalsRowFunctionCountNums indicates a COUNTNUMS function.
	TotalsRowFunctionCountNums TotalsRowFunction = "countNums"
	// TotalsRowFunctionStdDev indicates a STDDEV function.
	TotalsRowFunctionStdDev TotalsRowFunction = "stdDev"
	// TotalsRowFunctionVar indicates a VAR function.
	TotalsRowFunctionVar TotalsRowFunction = "var"
	// TotalsRowFunctionCustom indicates a custom formula.
	TotalsRowFunctionCustom TotalsRowFunction = "custom"
)

// TableColumns represents the container element for table columns
// (x:tableColumns).
type TableColumns struct {
	*openxml.CompositeElementBase
}

// NewTableColumns creates a new TableColumns element.
func NewTableColumns() *TableColumns {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"tableColumns",
		PrefixDefault,
	)

	return &TableColumns{
		CompositeElementBase: elem,
	}
}

// Count returns the count attribute value.
func (tc *TableColumns) Count() uint32 {
	attr, found := tc.GetAttribute(
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
func (tc *TableColumns) SetCount(count uint32) {
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Columns returns an iterator over all TableColumn elements.
func (tc *TableColumns) Columns() iter.Seq[*TableColumn] {
	return func(yield func(*TableColumn) bool) {
		for child := range tc.Children() {
			if child.LocalName() != "tableColumn" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var col *TableColumn
			if c, ok := child.(*TableColumn); ok {
				col = c
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				col = &TableColumn{CompositeElementBase: comp}
			}
			if col != nil && !yield(col) {
				return
			}
		}
	}
}

// ItemCount returns the actual number of TableColumn children.
func (tc *TableColumns) ItemCount() int {
	count := 0
	for range tc.Columns() {
		count++
	}

	return count
}

// AddColumn adds a new TableColumn and returns it.
func (tc *TableColumns) AddColumn(
	id uint32,
	name string,
) *TableColumn {
	col := NewTableColumn()
	col.SetId(id)
	col.SetName(name)
	tc.AppendChild(col)
	tc.SetCount(uint32(tc.ItemCount()))

	return col
}

// GetColumnByName returns the column with the given name, or nil if not found.
func (tc *TableColumns) GetColumnByName(
	name string,
) *TableColumn {
	for col := range tc.Columns() {
		if col.Name() == name {
			return col
		}
	}

	return nil
}

// GetColumnById returns the column with the given ID, or nil if not found.
func (tc *TableColumns) GetColumnById(
	id uint32,
) *TableColumn {
	for col := range tc.Columns() {
		if col.Id() == id {
			return col
		}
	}

	return nil
}

// GetColumnByIndex returns the column at the given index (0-based),
// or nil if out of range.
func (tc *TableColumns) GetColumnByIndex(
	index int,
) *TableColumn {
	i := 0
	for col := range tc.Columns() {
		if i == index {
			return col
		}
		i++
	}

	return nil
}

// RemoveColumn removes the column with the given ID.
func (tc *TableColumns) RemoveColumn(
	id uint32,
) bool {
	col := tc.GetColumnById(id)
	if col == nil {
		return false
	}
	removed := tc.RemoveChild(col)
	if removed {
		tc.SetCount(uint32(tc.ItemCount()))
	}

	return removed
}

// Clone creates a deep copy of this TableColumns element.
func (tc *TableColumns) Clone() openxml.Element {
	cloned := tc.CompositeElementBase.Clone()

	return &TableColumns{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this TableColumns element.
func (tc *TableColumns) CloneNode(
	deep bool,
) openxml.Element {
	cloned := tc.CompositeElementBase.CloneNode(
		deep,
	)

	return &TableColumns{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// TableColumn represents a table column element (x:tableColumn).
type TableColumn struct {
	*openxml.CompositeElementBase
}

// NewTableColumn creates a new TableColumn element.
func NewTableColumn() *TableColumn {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"tableColumn",
		PrefixDefault,
	)

	return &TableColumn{
		CompositeElementBase: elem,
	}
}

// Id returns the column ID. Attribute: id.
// This is a unique identifier within the table.
func (tc *TableColumn) Id() uint32 {
	attr, found := tc.GetAttribute("id", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetId sets the column ID. Attribute: id.
func (tc *TableColumn) SetId(id uint32) {
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"id",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// UniqueName returns the unique name for the column. Attribute: uniqueName.
// This is used for calculated column references.
func (tc *TableColumn) UniqueName() string {
	attr, found := tc.GetAttribute(
		"uniqueName",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetUniqueName sets the unique name for the column. Attribute: uniqueName.
func (tc *TableColumn) SetUniqueName(
	name string,
) {
	if name == "" {
		tc.RemoveAttribute("uniqueName", "")

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"uniqueName",
			"",
			name,
		),
	)
}

// Name returns the display name of the column. Attribute: name.
// This appears in the column header.
func (tc *TableColumn) Name() string {
	attr, found := tc.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the display name of the column. Attribute: name.
func (tc *TableColumn) SetName(name string) {
	if name == "" {
		tc.RemoveAttribute("name", "")

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// TotalsRowFunction returns the function for the totals row.
// Attribute: totalsRowFunction.
func (tc *TableColumn) TotalsRowFunction() TotalsRowFunction {
	attr, found := tc.GetAttribute(
		"totalsRowFunction",
		"",
	)
	if !found {
		return TotalsRowFunctionNone
	}

	return TotalsRowFunction(attr.Value())
}

// SetTotalsRowFunction sets the function for the totals row.
// Attribute: totalsRowFunction.
func (tc *TableColumn) SetTotalsRowFunction(
	fn TotalsRowFunction,
) {
	if fn == "" || fn == TotalsRowFunctionNone {
		tc.RemoveAttribute(
			"totalsRowFunction",
			"",
		)

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"totalsRowFunction",
			"",
			string(fn),
		),
	)
}

// TotalsRowLabel returns the label for the totals row.
// Attribute: totalsRowLabel. This is displayed instead of a function result.
func (tc *TableColumn) TotalsRowLabel() string {
	attr, found := tc.GetAttribute(
		"totalsRowLabel",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTotalsRowLabel sets the label for the totals row.
// Attribute: totalsRowLabel.
func (tc *TableColumn) SetTotalsRowLabel(
	label string,
) {
	if label == "" {
		tc.RemoveAttribute("totalsRowLabel", "")

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"totalsRowLabel",
			"",
			label,
		),
	)
}

// QueryTableFieldId returns the query table field ID.
// Attribute: queryTableFieldId.
func (tc *TableColumn) QueryTableFieldId() (uint32, bool) {
	attr, found := tc.GetAttribute(
		"queryTableFieldId",
		"",
	)
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val), true
}

// SetQueryTableFieldId sets the query table field ID.
// Attribute: queryTableFieldId.
func (tc *TableColumn) SetQueryTableFieldId(
	id uint32,
) {
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"queryTableFieldId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// ClearQueryTableFieldId removes the query table field ID attribute.
func (tc *TableColumn) ClearQueryTableFieldId() {
	tc.RemoveAttribute("queryTableFieldId", "")
}

// HeaderRowDxfId returns the differential formatting ID for the header cell.
// Attribute: headerRowDxfId.
func (tc *TableColumn) HeaderRowDxfId() (uint32, bool) {
	attr, found := tc.GetAttribute(
		"headerRowDxfId",
		"",
	)
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val), true
}

// SetHeaderRowDxfId sets the differential formatting ID for the header cell.
// Attribute: headerRowDxfId.
func (tc *TableColumn) SetHeaderRowDxfId(
	id uint32,
) {
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"headerRowDxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// ClearHeaderRowDxfId removes the header row DxfId attribute.
func (tc *TableColumn) ClearHeaderRowDxfId() {
	tc.RemoveAttribute("headerRowDxfId", "")
}

// DataDxfId returns the differential formatting ID for the data cells.
// Attribute: dataDxfId.
func (tc *TableColumn) DataDxfId() (uint32, bool) {
	attr, found := tc.GetAttribute(
		"dataDxfId",
		"",
	)
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val), true
}

// SetDataDxfId sets the differential formatting ID for the data cells.
// Attribute: dataDxfId.
func (tc *TableColumn) SetDataDxfId(id uint32) {
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"dataDxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// ClearDataDxfId removes the data DxfId attribute.
func (tc *TableColumn) ClearDataDxfId() {
	tc.RemoveAttribute("dataDxfId", "")
}

// TotalsRowDxfId returns the differential formatting ID for the totals cell.
// Attribute: totalsRowDxfId.
func (tc *TableColumn) TotalsRowDxfId() (uint32, bool) {
	attr, found := tc.GetAttribute(
		"totalsRowDxfId",
		"",
	)
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant: base 10
		32, //nolint:revive // add-constant: bit size
	)

	return uint32(val), true
}

// SetTotalsRowDxfId sets the differential formatting ID for the totals cell.
// Attribute: totalsRowDxfId.
func (tc *TableColumn) SetTotalsRowDxfId(
	id uint32,
) {
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"totalsRowDxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// ClearTotalsRowDxfId removes the totals row DxfId attribute.
func (tc *TableColumn) ClearTotalsRowDxfId() {
	tc.RemoveAttribute("totalsRowDxfId", "")
}

// HeaderRowCellStyle returns the named cell style for the header cell.
// Attribute: headerRowCellStyle.
func (tc *TableColumn) HeaderRowCellStyle() string {
	attr, found := tc.GetAttribute(
		"headerRowCellStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetHeaderRowCellStyle sets the named cell style for the header cell.
// Attribute: headerRowCellStyle.
func (tc *TableColumn) SetHeaderRowCellStyle(
	style string,
) {
	if style == "" {
		tc.RemoveAttribute(
			"headerRowCellStyle",
			"",
		)

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"headerRowCellStyle",
			"",
			style,
		),
	)
}

// DataCellStyle returns the named cell style for the data cells.
// Attribute: dataCellStyle.
func (tc *TableColumn) DataCellStyle() string {
	attr, found := tc.GetAttribute(
		"dataCellStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDataCellStyle sets the named cell style for the data cells.
// Attribute: dataCellStyle.
func (tc *TableColumn) SetDataCellStyle(
	style string,
) {
	if style == "" {
		tc.RemoveAttribute("dataCellStyle", "")

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"dataCellStyle",
			"",
			style,
		),
	)
}

// TotalsRowCellStyle returns the named cell style for the totals cell.
// Attribute: totalsRowCellStyle.
func (tc *TableColumn) TotalsRowCellStyle() string {
	attr, found := tc.GetAttribute(
		"totalsRowCellStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTotalsRowCellStyle sets the named cell style for the totals cell.
// Attribute: totalsRowCellStyle.
func (tc *TableColumn) SetTotalsRowCellStyle(
	style string,
) {
	if style == "" {
		tc.RemoveAttribute(
			"totalsRowCellStyle",
			"",
		)

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"totalsRowCellStyle",
			"",
			style,
		),
	)
}

// CalculatedColumnFormula returns the CalculatedColumnFormula child element,
// or nil if not present.
func (tc *TableColumn) CalculatedColumnFormula() *CalculatedColumnFormula {
	elem := tc.GetElement(
		"calculatedColumnFormula",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if ccf, ok := elem.(*CalculatedColumnFormula); ok {
		return ccf
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &CalculatedColumnFormula{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreateCalculatedColumnFormula returns the CalculatedColumnFormula
// child, creating it if needed.
func (tc *TableColumn) GetOrCreateCalculatedColumnFormula() *CalculatedColumnFormula { //nolint:revive // line-length-limit
	ccf := tc.CalculatedColumnFormula()
	if ccf != nil {
		return ccf
	}
	ccf = NewCalculatedColumnFormula()
	// Insert before totalsRowFormula if present
	trf := tc.TotalsRowFormula()
	if trf != nil {
		tc.InsertBefore(ccf, trf)
	} else {
		tc.AppendChild(ccf)
	}

	return ccf
}

// SetCalculatedColumnFormula sets the calculated column formula.
func (tc *TableColumn) SetCalculatedColumnFormula(
	formula string,
) {
	if formula == "" {
		ccf := tc.CalculatedColumnFormula()
		if ccf != nil {
			tc.RemoveChild(ccf)
		}

		return
	}
	ccf := tc.GetOrCreateCalculatedColumnFormula()
	ccf.SetFormula(formula)
}

// TotalsRowFormula returns the TotalsRowFormula child element,
// or nil if not present.
func (tc *TableColumn) TotalsRowFormula() *TotalsRowFormula {
	elem := tc.GetElement(
		"totalsRowFormula",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if trf, ok := elem.(*TotalsRowFormula); ok {
		return trf
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &TotalsRowFormula{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreateTotalsRowFormula returns the TotalsRowFormula child,
// creating it if needed.
func (tc *TableColumn) GetOrCreateTotalsRowFormula() *TotalsRowFormula {
	trf := tc.TotalsRowFormula()
	if trf != nil {
		return trf
	}
	trf = NewTotalsRowFormula()
	tc.AppendChild(trf)

	return trf
}

// SetTotalsRowFormula sets the totals row formula and function to custom.
func (tc *TableColumn) SetTotalsRowFormula(
	formula string,
) {
	if formula == "" {
		trf := tc.TotalsRowFormula()
		if trf != nil {
			tc.RemoveChild(trf)
		}
		// Clear the custom function if formula is removed
		if tc.TotalsRowFunction() == TotalsRowFunctionCustom {
			tc.SetTotalsRowFunction(
				TotalsRowFunctionNone,
			)
		}

		return
	}
	tc.SetTotalsRowFunction(
		TotalsRowFunctionCustom,
	)
	trf := tc.GetOrCreateTotalsRowFormula()
	trf.SetFormula(formula)
}

// Clone creates a deep copy of this TableColumn element.
func (tc *TableColumn) Clone() openxml.Element {
	cloned := tc.CompositeElementBase.Clone()

	return &TableColumn{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this TableColumn element.
func (tc *TableColumn) CloneNode(
	deep bool,
) openxml.Element {
	cloned := tc.CompositeElementBase.CloneNode(
		deep,
	)

	return &TableColumn{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CalculatedColumnFormula represents a calculated column formula element
// (x:calculatedColumnFormula). This contains the formula for a calculated
// column.
type CalculatedColumnFormula struct {
	*openxml.LeafElementBase
}

// NewCalculatedColumnFormula creates a new CalculatedColumnFormula element.
func NewCalculatedColumnFormula() *CalculatedColumnFormula {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"calculatedColumnFormula",
		PrefixDefault,
	)

	return &CalculatedColumnFormula{
		LeafElementBase: elem,
	}
}

// Formula returns the formula text.
func (cc *CalculatedColumnFormula) Formula() string {
	return cc.InnerText()
}

// SetFormula sets the formula text.
func (cc *CalculatedColumnFormula) SetFormula(
	formula string,
) {
	cc.SetInnerText(formula)
}

// Array returns whether this is an array formula. Attribute: array.
func (cc *CalculatedColumnFormula) Array() bool {
	attr, found := cc.GetAttribute("array", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetArray sets whether this is an array formula. Attribute: array.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cc *CalculatedColumnFormula) SetArray(
	value bool,
) {
	if value {
		cc.SetAttribute(
			openxml.NewAttribute(
				"",
				"array",
				"",
				attrValueTrue,
			),
		)
	} else {
		cc.RemoveAttribute("array", "")
	}
}

// Clone creates a deep copy of this CalculatedColumnFormula element.
func (cc *CalculatedColumnFormula) Clone() openxml.Element {
	cloned := cc.LeafElementBase.Clone()

	return &CalculatedColumnFormula{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CalculatedColumnFormula element.
func (cc *CalculatedColumnFormula) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cc.LeafElementBase.CloneNode(deep)

	return &CalculatedColumnFormula{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TotalsRowFormula represents a totals row formula element
// (x:totalsRowFormula). This contains a custom formula for the totals row.
type TotalsRowFormula struct {
	*openxml.LeafElementBase
}

// NewTotalsRowFormula creates a new TotalsRowFormula element.
func NewTotalsRowFormula() *TotalsRowFormula {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"totalsRowFormula",
		PrefixDefault,
	)

	return &TotalsRowFormula{
		LeafElementBase: elem,
	}
}

// Formula returns the formula text.
func (tr *TotalsRowFormula) Formula() string {
	return tr.InnerText()
}

// SetFormula sets the formula text.
func (tr *TotalsRowFormula) SetFormula(
	formula string,
) {
	tr.SetInnerText(formula)
}

// Array returns whether this is an array formula. Attribute: array.
func (tr *TotalsRowFormula) Array() bool {
	attr, found := tr.GetAttribute(
		attrNameArray,
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetArray sets whether this is an array formula. Attribute: array.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (tr *TotalsRowFormula) SetArray(
	value bool,
) {
	if value {
		tr.SetAttribute(
			openxml.NewAttribute(
				"",
				attrNameArray,
				"",
				attrValueTrue,
			),
		)
	} else {
		tr.RemoveAttribute(attrNameArray, "")
	}
}

// Clone creates a deep copy of this TotalsRowFormula element.
func (tr *TotalsRowFormula) Clone() openxml.Element {
	cloned := tr.LeafElementBase.Clone()

	return &TotalsRowFormula{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this TotalsRowFormula element.
func (tr *TotalsRowFormula) CloneNode(
	deep bool,
) openxml.Element {
	cloned := tr.LeafElementBase.CloneNode(deep)

	return &TotalsRowFormula{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
