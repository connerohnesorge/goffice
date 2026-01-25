// This file implements the high-level PivotTable API.
package spreadsheet

import (
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

const (
	// firstFieldIndex is the starting index for field collections.
	firstFieldIndex = int32(0)
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

	// Initialize the pivot table definition element
	ptDef := pivotPart.PivotTableDefinition()
	if ptDef != nil {
		// Set the name
		ptDef.SetName(pt.name)

		// Set up the location
		loc := ptDef.GetOrCreateLocation()
		loc.SetRef(
			destCell.String() + ":" + destCell.String(),
		)

		// Initialize pivot fields collection
		ptDef.GetOrCreatePivotFields()
	}

	return pt, nil
}

// Name returns the pivot table name.
func (p *PivotTable) Name() string {
	return p.name
}

// SetName sets the pivot table name.
func (p *PivotTable) SetName(name string) {
	p.name = name
	if ptDef := p.pivotPart.PivotTableDefinition(); ptDef != nil {
		ptDef.SetName(name)
	}
}

// SourceRange returns the source data range.
func (p *PivotTable) SourceRange() string {
	return p.sourceRange
}

// DestinationCell returns the destination cell.
func (p *PivotTable) DestinationCell() CellRef {
	return p.destCell
}
