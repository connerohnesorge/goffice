// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements style and formatting for PivotTable.
package spreadsheet

// SetStyle sets the pivot table style.
// Built-in styles include:
//   - PivotStyleLight1-28
//   - PivotStyleMedium1-28
//   - PivotStyleDark1-28
func (p *PivotTable) SetStyle(styleName string) {
	if ptDef := p.pivotPart.PivotTableDefinition(); ptDef != nil {
		styleInfo := ptDef.GetOrCreatePivotTableStyleInfo()
		styleInfo.SetName(styleName)
	}
}

// SetShowRowHeaders sets whether to show row headers.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotTable) SetShowRowHeaders(
	show bool,
) {
	if ptDef := p.pivotPart.PivotTableDefinition(); ptDef != nil {
		styleInfo := ptDef.GetOrCreatePivotTableStyleInfo()
		styleInfo.SetShowRowHeaders(show)
	}
}

// SetShowColumnHeaders sets whether to show column headers.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotTable) SetShowColumnHeaders(
	show bool,
) {
	if ptDef := p.pivotPart.PivotTableDefinition(); ptDef != nil {
		styleInfo := ptDef.GetOrCreatePivotTableStyleInfo()
		styleInfo.SetShowColHeaders(show)
	}
}

// SetShowRowStripes sets whether to show row stripes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotTable) SetShowRowStripes(
	show bool,
) {
	if ptDef := p.pivotPart.PivotTableDefinition(); ptDef != nil {
		styleInfo := ptDef.GetOrCreatePivotTableStyleInfo()
		styleInfo.SetShowRowStripes(show)
	}
}

// SetShowColumnStripes sets whether to show column stripes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotTable) SetShowColumnStripes(
	show bool,
) {
	if ptDef := p.pivotPart.PivotTableDefinition(); ptDef != nil {
		styleInfo := ptDef.GetOrCreatePivotTableStyleInfo()
		styleInfo.SetShowColStripes(show)
	}
}

// Refresh refreshes the pivot table data.
func (p *PivotTable) Refresh() error {
	// Mark the cache as needing refresh
	cacheDef := p.pivotPart.PivotTableCacheDefinitionPart()
	if cacheDef != nil {
		if pcd := cacheDef.PivotCacheDefinition(); pcd != nil {
			// Set the cache as invalid to force a refresh
			pcd.SetInvalid(true)
		}
	}

	return nil
}

// ClearAll clears all fields from the pivot table.
func (p *PivotTable) ClearAll() {
	p.rowFields = nil
	p.colFields = nil
	p.dataFields = nil
	p.pageFields = nil
	p.updateFields()
}
