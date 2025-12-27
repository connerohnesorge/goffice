// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements the high-level PivotTable API.
package spreadsheet

import (
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// AggregateFunction represents an aggregation function for data fields.
type AggregateFunction string

const (
	// AggregateSUM sums the values.
	AggregateSUM AggregateFunction = "sum"
	// AggregateCount counts the values.
	AggregateCount AggregateFunction = "count"
	// AggregateAverage averages the values.
	AggregateAverage AggregateFunction = "average"
	// AggregateMax returns the maximum value.
	AggregateMax AggregateFunction = "max"
	// AggregateMin returns the minimum value.
	AggregateMin AggregateFunction = "min"
	// AggregateProduct multiplies the values.
	AggregateProduct AggregateFunction = "product"
	// AggregateCountNums counts numeric values.
	AggregateCountNums AggregateFunction = "countNums"
	// AggregateStdDev calculates standard deviation.
	AggregateStdDev AggregateFunction = "stdDev"
	// AggregateStdDevP calculates population standard deviation.
	AggregateStdDevP AggregateFunction = "stdDevP"
	// AggregateVar calculates variance.
	AggregateVar AggregateFunction = "var"
	// AggregateVarP calculates population variance.
	AggregateVarP AggregateFunction = "varP"
)

// PivotTable represents a high-level wrapper around an Excel pivot table.
type PivotTable struct {
	sheet       *Sheet
	pivotPart   *parts.PivotTablePart
	sourceRange string
	destCell    CellRef
	rowFields   []string
	colFields   []string
	dataFields  []pivotDataField
	pageFields  []string
	name        string
}

// pivotDataField represents a data field with its aggregation function.
type pivotDataField struct {
	name      string
	aggregate AggregateFunction
}

// newPivotTable creates a new PivotTable wrapper.
func newPivotTable(
	sheet *Sheet,
	pivotPart *parts.PivotTablePart,
	sourceRange string,
	destCell CellRef,
) (*PivotTable, error) {
	pt := &PivotTable{
		sheet:       sheet,
		pivotPart:   pivotPart,
		sourceRange: sourceRange,
		destCell:    destCell,
		name:        "PivotTable1",
	}

	// TODO: Initialize the pivot table definition element
	// This would involve:
	// 1. Creating the PivotCacheDefinition
	// 2. Creating the PivotTableDefinition
	// 3. Setting up the location and cache reference

	return pt, nil
}

// Name returns the pivot table name.
func (p *PivotTable) Name() string {
	return p.name
}

// SetName sets the pivot table name.
func (p *PivotTable) SetName(name string) {
	p.name = name
	// TODO: Update the underlying element
}

// SourceRange returns the source data range.
func (p *PivotTable) SourceRange() string {
	return p.sourceRange
}

// DestinationCell returns the destination cell.
func (p *PivotTable) DestinationCell() CellRef {
	return p.destCell
}

// AddRowField adds a field to the row area.
func (p *PivotTable) AddRowField(
	name string,
) *PivotTable {
	p.rowFields = append(p.rowFields, name)
	// TODO: Update the underlying element

	return p
}

// AddColumnField adds a field to the column area.
func (p *PivotTable) AddColumnField(
	name string,
) *PivotTable {
	p.colFields = append(p.colFields, name)
	// TODO: Update the underlying element

	return p
}

// AddDataField adds a field to the data area with an aggregation function.
func (p *PivotTable) AddDataField(
	name string,
	aggregate AggregateFunction,
) *PivotTable {
	p.dataFields = append(
		p.dataFields,
		pivotDataField{
			name:      name,
			aggregate: aggregate,
		},
	)
	// TODO: Update the underlying element

	return p
}

// AddPageField adds a field to the page (filter) area.
func (p *PivotTable) AddPageField(
	name string,
) *PivotTable {
	p.pageFields = append(p.pageFields, name)
	// TODO: Update the underlying element

	return p
}

// RowFields returns the row field names.
func (p *PivotTable) RowFields() []string {
	result := make([]string, len(p.rowFields))
	copy(result, p.rowFields)

	return result
}

// ColumnFields returns the column field names.
func (p *PivotTable) ColumnFields() []string {
	result := make([]string, len(p.colFields))
	copy(result, p.colFields)

	return result
}

// PageFields returns the page (filter) field names.
func (p *PivotTable) PageFields() []string {
	result := make([]string, len(p.pageFields))
	copy(result, p.pageFields)

	return result
}

// SetStyle sets the pivot table style.
// Built-in styles include:
//   - PivotStyleLight1-28
//   - PivotStyleMedium1-28
//   - PivotStyleDark1-28
func (*PivotTable) SetStyle(styleName string) {
	// TODO: Update the underlying element
	_ = styleName
}

// SetShowRowHeaders sets whether to show row headers.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (*PivotTable) SetShowRowHeaders(
	show bool,
) {
	// TODO: Update the underlying element
	_ = show
}

// SetShowColumnHeaders sets whether to show column headers.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (*PivotTable) SetShowColumnHeaders(
	show bool,
) {
	// TODO: Update the underlying element
	_ = show
}

// SetShowRowStripes sets whether to show row stripes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (*PivotTable) SetShowRowStripes(
	show bool,
) {
	// TODO: Update the underlying element
	_ = show
}

// SetShowColumnStripes sets whether to show column stripes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (*PivotTable) SetShowColumnStripes(
	show bool,
) {
	// TODO: Update the underlying element
	_ = show
}

// Refresh refreshes the pivot table data.
func (*PivotTable) Refresh() error {
	// TODO: Implement refresh logic
	// This would recalculate the cache and update the pivot table

	return nil
}

// ClearAll clears all fields from the pivot table.
func (p *PivotTable) ClearAll() {
	p.rowFields = nil
	p.colFields = nil
	p.dataFields = nil
	p.pageFields = nil
	// TODO: Update the underlying element
}

// RemoveRowField removes a field from the row area.
func (p *PivotTable) RemoveRowField(
	name string,
) bool {
	for i, f := range p.rowFields {
		if f == name {
			p.rowFields = append(
				p.rowFields[:i],
				p.rowFields[i+1:]...,
			)
			// TODO: Update the underlying element

			return true
		}
	}

	return false
}

// RemoveColumnField removes a field from the column area.
func (p *PivotTable) RemoveColumnField(
	name string,
) bool {
	for i, f := range p.colFields {
		if f == name {
			p.colFields = append(
				p.colFields[:i],
				p.colFields[i+1:]...,
			)
			// TODO: Update the underlying element

			return true
		}
	}

	return false
}

// RemoveDataField removes a field from the data area.
func (p *PivotTable) RemoveDataField(
	name string,
) bool {
	for i, f := range p.dataFields {
		if f.name == name {
			p.dataFields = append(
				p.dataFields[:i],
				p.dataFields[i+1:]...,
			)
			// TODO: Update the underlying element

			return true
		}
	}

	return false
}

// RemovePageField removes a field from the page area.
func (p *PivotTable) RemovePageField(
	name string,
) bool {
	for i, f := range p.pageFields {
		if f == name {
			p.pageFields = append(
				p.pageFields[:i],
				p.pageFields[i+1:]...,
			)
			// TODO: Update the underlying element

			return true
		}
	}

	return false
}
