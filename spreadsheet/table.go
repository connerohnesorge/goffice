// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements the high-level Table API.
package spreadsheet

import (
	"sync/atomic"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// tableIDCounter is used to generate unique table IDs.
var tableIDCounter uint32

// nextTableID returns the next available table ID.
func nextTableID() uint32 {
	return atomic.AddUint32(&tableIDCounter, 1)
}

// Table represents a high-level wrapper around an Excel table.
type Table struct {
	sheet     *Sheet
	tablePart *parts.TableDefinitionPart
	table     *elements.Table
}

// newTable creates a new Table wrapper and initializes it.
func newTable(
	sheet *Sheet,
	tablePart *parts.TableDefinitionPart,
	ref, name string,
) (*Table, error) {
	t := &Table{
		sheet:     sheet,
		tablePart: tablePart,
	}

	// Create the table element
	tableID := nextTableID()
	t.table = elements.NewTableWithDefaults(
		tableID,
		name,
		name,
		ref,
	)

	// Auto-generate columns based on the range
	rangeRef, err := ParseRangeRef(ref)
	if err != nil {
		return nil, err
	}

	// Add table columns for each column in the range
	cols := t.table.GetOrCreateTableColumns()
	colCount := rangeRef.End.Col - rangeRef.Start.Col + 1
	for i := range colCount {
		colName := "Column" + string(rune('1'+i))
		cols.AddColumn(uint32(i+1), colName)
	}

	// Add default table style
	styleInfo := t.table.GetOrCreateTableStyleInfo()
	styleInfo.SetName("TableStyleMedium2")
	styleInfo.SetShowFirstColumn(false)
	styleInfo.SetShowLastColumn(false)
	styleInfo.SetShowRowStripes(true)
	styleInfo.SetShowColumnStripes(false)

	return t, nil
}

// newTableFromPart creates a Table wrapper from an existing table definition
// part.
func newTableFromPart(
	sheet *Sheet,
	tablePart *parts.TableDefinitionPart,
) *Table {
	t := &Table{
		sheet:     sheet,
		tablePart: tablePart,
	}

	// Try to get the table element from the part
	root := tablePart.Table()
	if root != nil {
		t.table = root
	}

	return t
}

// Name returns the table name.
func (t *Table) Name() string {
	if t.table == nil {
		return ""
	}

	return t.table.Name()
}

// DisplayName returns the display name of the table.
func (t *Table) DisplayName() string {
	if t.table == nil {
		return ""
	}

	return t.table.DisplayName()
}

// SetName sets both the name and display name of the table.
func (t *Table) SetName(name string) {
	if t.table == nil {
		return
	}
	t.table.SetName(name)
	t.table.SetDisplayName(name)
}

// Ref returns the table's cell range reference.
func (t *Table) Ref() string {
	if t.table == nil {
		return ""
	}

	return t.table.Ref()
}

// SetRef sets the table's cell range reference.
func (t *Table) SetRef(ref string) error {
	if t.table == nil {
		return nil
	}

	// Validate the range reference
	_, err := ParseRangeRef(ref)
	if err != nil {
		return err
	}

	t.table.SetRef(ref)

	return nil
}

// SetStyle sets the table style by name.
// Built-in table styles include:
//   - TableStyleLight1-21
//   - TableStyleMedium1-28
//   - TableStyleDark1-11
func (t *Table) SetStyle(styleName string) {
	if t.table == nil {
		return
	}

	styleInfo := t.table.GetOrCreateTableStyleInfo()
	styleInfo.SetName(styleName)
}

// SetShowHeaderRow sets whether to show the header row.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Table) SetShowHeaderRow(show bool) {
	if t.table == nil {
		return
	}

	if show {
		t.table.SetHeaderRowCount(1)
	} else {
		t.table.SetHeaderRowCount(0)
	}
}

// SetShowTotalsRow sets whether to show the totals row.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Table) SetShowTotalsRow(show bool) {
	if t.table == nil {
		return
	}

	if show {
		t.table.SetTotalsRowCount(1)
	} else {
		t.table.SetTotalsRowCount(0)
	}
	t.table.SetTotalsRowShown(show)
}

// SetShowRowStripes sets whether to show row stripes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Table) SetShowRowStripes(show bool) {
	if t.table == nil {
		return
	}

	styleInfo := t.table.GetOrCreateTableStyleInfo()
	styleInfo.SetShowRowStripes(show)
}

// SetShowColumnStripes sets whether to show column stripes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Table) SetShowColumnStripes(show bool) {
	if t.table == nil {
		return
	}

	styleInfo := t.table.GetOrCreateTableStyleInfo()
	styleInfo.SetShowColumnStripes(show)
}

// SetShowFirstColumn sets whether to highlight the first column.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Table) SetShowFirstColumn(show bool) {
	if t.table == nil {
		return
	}

	styleInfo := t.table.GetOrCreateTableStyleInfo()
	styleInfo.SetShowFirstColumn(show)
}

// SetShowLastColumn sets whether to highlight the last column.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Table) SetShowLastColumn(show bool) {
	if t.table == nil {
		return
	}

	styleInfo := t.table.GetOrCreateTableStyleInfo()
	styleInfo.SetShowLastColumn(show)
}

// Column returns a TableColumn by name.
func (t *Table) Column(name string) *TableColumn {
	if t.table == nil {
		return nil
	}

	col := t.table.GetColumnByName(name)
	if col == nil {
		return nil
	}

	return newTableColumn(t, col)
}

// ColumnByID returns a TableColumn by ID.
func (t *Table) ColumnByID(
	id uint32,
) *TableColumn {
	if t.table == nil {
		return nil
	}

	col := t.table.GetColumnById(id)
	if col == nil {
		return nil
	}

	return newTableColumn(t, col)
}

// ColumnCount returns the number of columns in the table.
func (t *Table) ColumnCount() int {
	if t.table == nil {
		return 0
	}

	return t.table.ColumnCount()
}

// AddColumn adds a new column to the table.
func (t *Table) AddColumn(
	name string,
) *TableColumn {
	if t.table == nil {
		return nil
	}

	nextID := uint32(t.table.ColumnCount() + 1)
	col := t.table.AddColumn(nextID, name)

	return newTableColumn(t, col)
}

// TableColumn represents a column in a table.
type TableColumn struct {
	table  *Table
	column *elements.TableColumn
}

// newTableColumn creates a new TableColumn wrapper.
func newTableColumn(
	table *Table,
	column *elements.TableColumn,
) *TableColumn {
	return &TableColumn{
		table:  table,
		column: column,
	}
}

// Name returns the column name.
func (tc *TableColumn) Name() string {
	return tc.column.Name()
}

// SetName sets the column name.
func (tc *TableColumn) SetName(name string) {
	tc.column.SetName(name)
}

// ID returns the column ID.
func (tc *TableColumn) ID() uint32 {
	return tc.column.Id()
}

// TotalsRowLabel returns the totals row label for this column.
func (tc *TableColumn) TotalsRowLabel() string {
	return tc.column.TotalsRowLabel()
}

// SetTotalsRowLabel sets the totals row label for this column.
func (tc *TableColumn) SetTotalsRowLabel(
	label string,
) {
	tc.column.SetTotalsRowLabel(label)
}

// TotalsRowFunction returns the totals row function for this column.
func (tc *TableColumn) TotalsRowFunction() string {
	return string(tc.column.TotalsRowFunction())
}

// SetTotalsRowFunction sets the totals row function for this column.
// Valid functions: sum, count, countNums, average, max, min,
// stdDev, var, custom
func (tc *TableColumn) SetTotalsRowFunction(
	function string,
) {
	tc.column.SetTotalsRowFunction(
		elements.TotalsRowFunction(function),
	)
}

// SetCalculatedColumnFormula sets a calculated column formula.
func (tc *TableColumn) SetCalculatedColumnFormula(
	formula string,
) {
	tc.column.SetCalculatedColumnFormula(formula)
}

// CalculatedColumnFormula returns the calculated column formula.
func (tc *TableColumn) CalculatedColumnFormula() string {
	ccf := tc.column.CalculatedColumnFormula()
	if ccf == nil {
		return ""
	}

	return ccf.Formula()
}
