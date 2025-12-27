package elements

//revive:disable:file-length-limit many table properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Table represents the table root element (x:table).
// It defines an Excel table with columns, filters, and styles.
type Table struct {
	*openxml.PartRootElementBase
}

// NewTable creates a new Table element.
func NewTable() *Table {
	elem := openxml.NewPartRootElement(
		NamespaceSML,
		"table",
		PrefixDefault,
	)

	return &Table{PartRootElementBase: elem}
}

// NewTableWithDefaults creates a new Table element with sensible defaults.
func NewTableWithDefaults(
	id uint32,
	name, displayName, ref string,
) *Table {
	t := NewTable()
	t.SetId(id)
	t.SetName(name)
	t.SetDisplayName(displayName)
	t.SetRef(ref)
	t.SetTotalsRowShown(false)

	return t
}

// Id returns the table ID. Attribute: id.
// This is a unique identifier for the table within the workbook.
func (t *Table) Id() uint32 {
	attr, found := t.GetAttribute("id", "")
	if !found {
		return 0
	}
	//nolint:revive // add-constant: 10 and 32 are standard parse parameters
	val, _ := strconv.ParseUint(
		attr.Value(),
		10,
		32,
	)

	return uint32(val)
}

// SetId sets the table ID. Attribute: id.
func (t *Table) SetId(id uint32) {
	t.SetAttribute(
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

// Name returns the table name. Attribute: name.
// This is the internal name used in formulas.
func (t *Table) Name() string {
	attr, found := t.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the table name. Attribute: name.
func (t *Table) SetName(name string) {
	if name == "" {
		t.RemoveAttribute("name", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// DisplayName returns the table display name. Attribute: displayName.
// This is the name shown to users in the UI.
func (t *Table) DisplayName() string {
	attr, found := t.GetAttribute(
		"displayName",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDisplayName sets the table display name. Attribute: displayName.
func (t *Table) SetDisplayName(
	displayName string,
) {
	if displayName == "" {
		t.RemoveAttribute("displayName", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"displayName",
			"",
			displayName,
		),
	)
}

// Ref returns the cell range reference for the table. Attribute: ref.
// Example: "A1:D10"
func (t *Table) Ref() string {
	attr, found := t.GetAttribute("ref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the cell range reference for the table. Attribute: ref.
func (t *Table) SetRef(ref string) {
	if ref == "" {
		t.RemoveAttribute("ref", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"ref",
			"",
			ref,
		),
	)
}

// Comment returns the table comment. Attribute: comment.
func (t *Table) Comment() string {
	attr, found := t.GetAttribute("comment", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetComment sets the table comment. Attribute: comment.
func (t *Table) SetComment(comment string) {
	if comment == "" {
		t.RemoveAttribute("comment", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"comment",
			"",
			comment,
		),
	)
}

// HeaderRowCount returns the number of header rows. Attribute: headerRowCount.
// Default is 1.
func (t *Table) HeaderRowCount() uint32 {
	attr, found := t.GetAttribute(
		"headerRowCount",
		"",
	)
	if !found {
		return 1 // Default
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetHeaderRowCount sets the number of header rows. Attribute: headerRowCount.
func (t *Table) SetHeaderRowCount(count uint32) {
	if count == 1 {
		t.RemoveAttribute(
			"headerRowCount",
			"",
		) // 1 is default

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"headerRowCount",
			"",
			strconv.FormatUint(
				uint64(count),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// InsertRow returns whether an insert row is shown. Attribute: insertRow.
func (t *Table) InsertRow() bool {
	attr, found := t.GetAttribute("insertRow", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetInsertRow sets whether an insert row is shown. Attribute: insertRow.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Table) SetInsertRow(value bool) {
	if value {
		t.SetAttribute(
			openxml.NewAttribute(
				"",
				"insertRow",
				"",
				attrValueTrue,
			),
		)
	} else {
		t.RemoveAttribute("insertRow", "")
	}
}

// InsertRowShift returns whether rows are shifted when inserting.
// Attribute: insertRowShift.
func (t *Table) InsertRowShift() bool {
	attr, found := t.GetAttribute(
		"insertRowShift",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetInsertRowShift sets whether rows are shifted when inserting.
// Attribute: insertRowShift.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Table) SetInsertRowShift(value bool) {
	if value {
		t.SetAttribute(
			openxml.NewAttribute(
				"",
				"insertRowShift",
				"",
				attrValueTrue,
			),
		)
	} else {
		t.RemoveAttribute("insertRowShift", "")
	}
}

// TotalsRowCount returns the number of total rows. Attribute: totalsRowCount.
// Default is 0.
func (t *Table) TotalsRowCount() uint32 {
	attr, found := t.GetAttribute(
		"totalsRowCount",
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

// SetTotalsRowCount sets the number of total rows. Attribute: totalsRowCount.
func (t *Table) SetTotalsRowCount(count uint32) {
	if count == 0 {
		t.RemoveAttribute("totalsRowCount", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"totalsRowCount",
			"",
			strconv.FormatUint(
				uint64(count),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// TotalsRowShown returns whether the totals row is shown.
// Attribute: totalsRowShown. Default is true.
func (t *Table) TotalsRowShown() bool {
	attr, found := t.GetAttribute(
		"totalsRowShown",
		"",
	)
	if !found {
		return true // Default
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetTotalsRowShown sets whether the totals row is shown. Attribute: totalsRowShown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Table) SetTotalsRowShown(value bool) {
	if value {
		t.RemoveAttribute(
			"totalsRowShown",
			"",
		) // true is default
	} else {
		t.SetAttribute(
			openxml.NewAttribute("", "totalsRowShown", "", attrValueZero),
		)
	}
}

// Published returns whether the table is published. Attribute: published.
func (t *Table) Published() bool {
	attr, found := t.GetAttribute("published", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPublished sets whether the table is published. Attribute: published.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Table) SetPublished(value bool) {
	if value {
		t.SetAttribute(
			openxml.NewAttribute(
				"",
				"published",
				"",
				attrValueTrue,
			),
		)
	} else {
		t.RemoveAttribute("published", "")
	}
}

// HeaderRowDxfId returns the differential formatting ID for the header row.
// Attribute: headerRowDxfId.
func (t *Table) HeaderRowDxfId() (uint32, bool) {
	attr, found := t.GetAttribute(
		"headerRowDxfId",
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

// SetHeaderRowDxfId sets the differential formatting ID for the header row.
// Attribute: headerRowDxfId.
func (t *Table) SetHeaderRowDxfId(id uint32) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"headerRowDxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10, //nolint:revive // add-constant
			),
		),
	)
}

// ClearHeaderRowDxfId removes the header row DxfId attribute.
func (t *Table) ClearHeaderRowDxfId() {
	t.RemoveAttribute("headerRowDxfId", "")
}

// DataDxfId returns the differential formatting ID for the data area.
// Attribute: dataDxfId.
func (t *Table) DataDxfId() (uint32, bool) {
	attr, found := t.GetAttribute("dataDxfId", "")
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

// SetDataDxfId sets the differential formatting ID for the data area.
// Attribute: dataDxfId.
func (t *Table) SetDataDxfId(id uint32) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dataDxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10, //nolint:revive // add-constant
			),
		),
	)
}

// ClearDataDxfId removes the data DxfId attribute.
func (t *Table) ClearDataDxfId() {
	t.RemoveAttribute("dataDxfId", "")
}

// TotalsRowDxfId returns the differential formatting ID for the totals row.
// Attribute: totalsRowDxfId.
func (t *Table) TotalsRowDxfId() (uint32, bool) {
	attr, found := t.GetAttribute(
		"totalsRowDxfId",
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

// SetTotalsRowDxfId sets the differential formatting ID for the totals row.
// Attribute: totalsRowDxfId.
func (t *Table) SetTotalsRowDxfId(id uint32) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"totalsRowDxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10, //nolint:revive // add-constant
			),
		),
	)
}

// ClearTotalsRowDxfId removes the totals row DxfId attribute.
func (t *Table) ClearTotalsRowDxfId() {
	t.RemoveAttribute("totalsRowDxfId", "")
}

// HeaderRowBorderDxfId returns the differential formatting ID for the header
// row border. Attribute: headerRowBorderDxfId.
func (t *Table) HeaderRowBorderDxfId() (uint32, bool) {
	attr, found := t.GetAttribute(
		"headerRowBorderDxfId",
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

// SetHeaderRowBorderDxfId sets the differential formatting ID for the header
// row border. Attribute: headerRowBorderDxfId.
func (t *Table) SetHeaderRowBorderDxfId(
	id uint32,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"headerRowBorderDxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// ClearHeaderRowBorderDxfId removes the header row border DxfId attribute.
func (t *Table) ClearHeaderRowBorderDxfId() {
	t.RemoveAttribute("headerRowBorderDxfId", "")
}

// TableBorderDxfId returns the differential formatting ID for the table border.
// Attribute: tableBorderDxfId.
func (t *Table) TableBorderDxfId() (uint32, bool) {
	attr, found := t.GetAttribute(
		"tableBorderDxfId",
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

// SetTableBorderDxfId sets the differential formatting ID for the table border.
// Attribute: tableBorderDxfId.
func (t *Table) SetTableBorderDxfId(id uint32) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"tableBorderDxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// ClearTableBorderDxfId removes the table border DxfId attribute.
func (t *Table) ClearTableBorderDxfId() {
	t.RemoveAttribute("tableBorderDxfId", "")
}

// TotalsRowBorderDxfId returns the differential formatting ID for the totals
// row border. Attribute: totalsRowBorderDxfId.
func (t *Table) TotalsRowBorderDxfId() (uint32, bool) {
	attr, found := t.GetAttribute(
		"totalsRowBorderDxfId",
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

// SetTotalsRowBorderDxfId sets the differential formatting ID
// for the totals row border. Attribute: totalsRowBorderDxfId.
func (t *Table) SetTotalsRowBorderDxfId(
	id uint32,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"totalsRowBorderDxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// ClearTotalsRowBorderDxfId removes the totals row border DxfId attribute.
func (t *Table) ClearTotalsRowBorderDxfId() {
	t.RemoveAttribute("totalsRowBorderDxfId", "")
}

// HeaderRowCellStyle returns the named cell style for the header row.
// Attribute: headerRowCellStyle.
func (t *Table) HeaderRowCellStyle() string {
	attr, found := t.GetAttribute(
		"headerRowCellStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetHeaderRowCellStyle sets the named cell style for the header row.
// Attribute: headerRowCellStyle.
func (t *Table) SetHeaderRowCellStyle(
	style string,
) {
	if style == "" {
		t.RemoveAttribute(
			"headerRowCellStyle",
			"",
		)

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"headerRowCellStyle",
			"",
			style,
		),
	)
}

// DataCellStyle returns the named cell style for the data area.
// Attribute: dataCellStyle.
func (t *Table) DataCellStyle() string {
	attr, found := t.GetAttribute(
		"dataCellStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDataCellStyle sets the named cell style for the data area.
// Attribute: dataCellStyle.
func (t *Table) SetDataCellStyle(style string) {
	if style == "" {
		t.RemoveAttribute("dataCellStyle", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dataCellStyle",
			"",
			style,
		),
	)
}

// TotalsRowCellStyle returns the named cell style for the totals row.
// Attribute: totalsRowCellStyle.
func (t *Table) TotalsRowCellStyle() string {
	attr, found := t.GetAttribute(
		"totalsRowCellStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTotalsRowCellStyle sets the named cell style for the totals row.
// Attribute: totalsRowCellStyle.
func (t *Table) SetTotalsRowCellStyle(
	style string,
) {
	if style == "" {
		t.RemoveAttribute(
			"totalsRowCellStyle",
			"",
		)

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"totalsRowCellStyle",
			"",
			style,
		),
	)
}

// ConnectionId returns the connection ID for external data.
// Attribute: connectionId.
func (t *Table) ConnectionId() (uint32, bool) {
	attr, found := t.GetAttribute(
		"connectionId",
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

// SetConnectionId sets the connection ID for external data.
// Attribute: connectionId.
func (t *Table) SetConnectionId(id uint32) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"connectionId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// ClearConnectionId removes the connection ID attribute.
func (t *Table) ClearConnectionId() {
	t.RemoveAttribute("connectionId", "")
}

// AutoFilter returns the AutoFilter child element, or nil if not present.
func (t *Table) AutoFilter() *AutoFilter {
	elem := t.GetElement(
		"autoFilter",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if af, ok := elem.(*AutoFilter); ok {
		return af
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &AutoFilter{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateAutoFilter returns the AutoFilter child element,
// creating it if needed.
func (t *Table) GetOrCreateAutoFilter() *AutoFilter {
	af := t.AutoFilter()
	if af != nil {
		return af
	}
	af = NewAutoFilter()
	// AutoFilter should come before tableColumns
	tc := t.TableColumns()
	if tc != nil {
		t.InsertBefore(af, tc)
	} else {
		t.AppendChild(af)
	}

	return af
}

// RemoveAutoFilter removes the AutoFilter child element.
func (t *Table) RemoveAutoFilter() bool {
	af := t.AutoFilter()
	if af == nil {
		return false
	}

	return t.RemoveChild(af)
}

// SortState returns the SortState child element, or nil if not present.
func (t *Table) SortState() *SortState {
	elem := t.GetElement(
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

// GetOrCreateSortState returns the SortState child element,
// creating it if needed.
func (t *Table) GetOrCreateSortState() *SortState {
	ss := t.SortState()
	if ss != nil {
		return ss
	}
	ss = NewSortState()
	// SortState should come after autoFilter but before tableColumns
	af := t.AutoFilter()
	tc := t.TableColumns()

	switch {
	case tc != nil:
		t.InsertBefore(ss, tc)
	case af != nil:
		t.InsertAfter(ss, af)
	default:
		t.AppendChild(ss)
	}

	return ss
}

// RemoveSortState removes the SortState child element.
func (t *Table) RemoveSortState() bool {
	ss := t.SortState()
	if ss == nil {
		return false
	}

	return t.RemoveChild(ss)
}

// TableColumns returns the TableColumns child element, or nil if not present.
func (t *Table) TableColumns() *TableColumns {
	elem := t.GetElement(
		"tableColumns",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if tc, ok := elem.(*TableColumns); ok {
		return tc
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableColumns{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTableColumns returns the TableColumns child element,
// creating it if needed.
func (t *Table) GetOrCreateTableColumns() *TableColumns {
	tc := t.TableColumns()
	if tc != nil {
		return tc
	}
	tc = NewTableColumns()
	// TableColumns should come before tableStyleInfo
	tsi := t.TableStyleInfo()
	if tsi != nil {
		t.InsertBefore(tc, tsi)
	} else {
		t.AppendChild(tc)
	}

	return tc
}

// TableStyleInfo returns the TableStyleInfo child element, nil if not present.
func (t *Table) TableStyleInfo() *TableStyleInfo {
	elem := t.GetElement(
		"tableStyleInfo",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if tsi, ok := elem.(*TableStyleInfo); ok {
		return tsi
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &TableStyleInfo{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreateTableStyleInfo returns the TableStyleInfo child element,
// creating it if needed.
func (t *Table) GetOrCreateTableStyleInfo() *TableStyleInfo {
	tsi := t.TableStyleInfo()
	if tsi != nil {
		return tsi
	}
	tsi = NewTableStyleInfo()
	// TableStyleInfo should be last
	t.AppendChild(tsi)

	return tsi
}

// RemoveTableStyleInfo removes the TableStyleInfo child element.
func (t *Table) RemoveTableStyleInfo() bool {
	tsi := t.TableStyleInfo()
	if tsi == nil {
		return false
	}

	return t.RemoveChild(tsi)
}

// AddColumn adds a new column to the table.
func (t *Table) AddColumn(
	id uint32,
	name string,
) *TableColumn {
	tc := t.GetOrCreateTableColumns()

	return tc.AddColumn(id, name)
}

// GetColumnByName returns the column with the given name, or nil if not found.
func (t *Table) GetColumnByName(
	name string,
) *TableColumn {
	tc := t.TableColumns()
	if tc == nil {
		return nil
	}

	return tc.GetColumnByName(name)
}

// GetColumnById returns the column with the given ID, or nil if not found.
func (t *Table) GetColumnById(
	id uint32,
) *TableColumn {
	tc := t.TableColumns()
	if tc == nil {
		return nil
	}

	return tc.GetColumnById(id)
}

// ColumnCount returns the number of columns in the table.
func (t *Table) ColumnCount() int {
	tc := t.TableColumns()
	if tc == nil {
		return 0
	}

	return tc.ItemCount()
}

// Columns returns an iterator over all table columns.
func (t *Table) Columns() iter.Seq[*TableColumn] {
	tc := t.TableColumns()
	if tc == nil {
		return func(_ func(*TableColumn) bool) {}
	}

	return tc.Columns()
}

// Clone creates a deep copy of this Table element.
func (t *Table) Clone() openxml.Element {
	cloned := t.PartRootElementBase.Clone()

	return &Table{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this Table element.
func (t *Table) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.PartRootElementBase.CloneNode(
		deep,
	)

	return &Table{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}
