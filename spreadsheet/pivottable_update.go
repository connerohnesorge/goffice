// This file implements field update logic for PivotTable.
package spreadsheet

import (
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// updateFields updates the underlying XML elements based on
// the current field configuration.
func (p *PivotTable) updateFields() {
	ptDef := p.pivotPart.PivotTableDefinition()
	if ptDef == nil {
		return
	}

	// Build field map and all fields list
	fieldMap, allFields, axisInfo := p.buildFieldMap()

	// Update PivotFields collection with full metadata
	updatePivotFieldsCollection(
		ptDef,
		allFields,
		axisInfo,
	)

	// Update individual field areas
	p.updateRowFieldsArea(ptDef, fieldMap)
	p.updateColFieldsArea(ptDef, fieldMap)
	p.updatePageFieldsArea(ptDef, fieldMap)
	p.updateDataFieldsArea(ptDef, fieldMap)

	// Update location element
	p.updateLocation(ptDef, fieldMap)

	// Ensure cache relationship exists
	if err := p.ensureCacheRelationship(); err != nil {
		// Log error but don't fail the update
		// This is a non-critical operation
		_ = err
	}
}

// fieldAxisInfo contains axis information for a field.
type fieldAxisInfo struct {
	axis      string // "axisRow", "axisCol", "axisPage", "axisValues"
	dataField bool
}

// buildFieldMap creates a map of unique field names to their indices and axis info.
func (p *PivotTable) buildFieldMap() (fieldMap map[string]int32, allFields []string, axisInfo map[string]fieldAxisInfo) {
	fieldMap = make(map[string]int32)
	axisInfo = make(map[string]fieldAxisInfo)
	fieldIndex := firstFieldIndex
	allFields = make([]string, 0)

	addField := func(name string, axis string, isDataField bool) {
		if _, exists := fieldMap[name]; exists {
			// Update axis info if this field is appearing in multiple contexts
			if axis != "" {
				axisInfo[name] = fieldAxisInfo{
					axis:      axis,
					dataField: isDataField,
				}
			}

			return
		}
		fieldMap[name] = fieldIndex
		allFields = append(allFields, name)
		axisInfo[name] = fieldAxisInfo{
			axis:      axis,
			dataField: isDataField,
		}
		fieldIndex++
	}

	for _, name := range p.rowFields {
		addField(name, "axisRow", false)
	}
	for _, name := range p.colFields {
		addField(name, "axisCol", false)
	}
	for _, name := range p.pageFields {
		addField(name, "axisPage", false)
	}
	for _, df := range p.dataFields {
		addField(df.name, "axisValues", true)
	}

	return fieldMap, allFields, axisInfo
}

// updatePivotFieldsCollection updates the PivotFields collection.
// This now serializes ALL field metadata including axis, showAll, compact, etc.
func updatePivotFieldsCollection(
	ptDef *elements.PivotTableDefinition,
	allFields []string,
	axisInfo map[string]fieldAxisInfo,
) {
	pivotFields := ptDef.GetOrCreatePivotFields()
	pivotFields.SetCount(uint32(len(allFields)))

	// Clear existing pivot fields and recreate
	for child := range pivotFields.Children() {
		pivotFields.RemoveChild(child)
	}

	for _, fieldName := range allFields {
		field := pivotFields.AddField()
		field.SetName(fieldName)

		// Set axis type if field is assigned to an axis
		if info, ok := axisInfo[fieldName]; ok {
			if info.axis != "" {
				field.SetAxis(info.axis)
			}
			if info.dataField {
				field.SetDataField(true)
			}
		}

		// Set default attributes for field metadata
		// These can be overridden by users via the Field() API later
		// ShowAll defaults to true (per Excel spec)
		field.SetShowAll(true)

		// Compact defaults to true (per Excel spec)
		field.SetCompact(true)

		// HideNewItems defaults to false
		// (no need to set explicitly as false is default)

		// ShowDropDowns defaults to true
		field.SetShowDropDowns(true)

		// DefaultSubtotal defaults to true
		field.SetDefaultSubtotal(true)
	}
}

// updateRowFieldsArea updates the row fields area.
func (p *PivotTable) updateRowFieldsArea(
	ptDef *elements.PivotTableDefinition,
	fieldMap map[string]int32,
) {
	if len(p.rowFields) == 0 {
		if rf := ptDef.RowFields(); rf != nil {
			ptDef.RemoveChild(rf)
		}

		return
	}

	rowFields := ptDef.GetOrCreateRowFields()
	rowFields.SetCount(uint32(len(p.rowFields)))

	// Clear existing row fields
	for child := range rowFields.Children() {
		rowFields.RemoveChild(child)
	}

	for _, name := range p.rowFields {
		idx, ok := fieldMap[name]
		if !ok {
			continue
		}
		rowFields.AddField(idx)
	}
}

// updateColFieldsArea updates the column fields area.
func (p *PivotTable) updateColFieldsArea(
	ptDef *elements.PivotTableDefinition,
	fieldMap map[string]int32,
) {
	if len(p.colFields) == 0 {
		if cf := ptDef.ColFields(); cf != nil {
			ptDef.RemoveChild(cf)
		}

		return
	}

	colFields := ptDef.GetOrCreateColFields()
	colFields.SetCount(uint32(len(p.colFields)))

	// Clear existing col fields
	for child := range colFields.Children() {
		colFields.RemoveChild(child)
	}

	for _, name := range p.colFields {
		idx, ok := fieldMap[name]
		if !ok {
			continue
		}
		colFields.AddField(idx)
	}
}

// updatePageFieldsArea updates the page fields area.
func (p *PivotTable) updatePageFieldsArea(
	ptDef *elements.PivotTableDefinition,
	fieldMap map[string]int32,
) {
	if len(p.pageFields) == 0 {
		if pf := ptDef.PageFields(); pf != nil {
			ptDef.RemoveChild(pf)
		}

		return
	}

	pageFields := ptDef.GetOrCreatePageFields()
	pageFields.SetCount(uint32(len(p.pageFields)))

	// Clear existing page fields
	for child := range pageFields.Children() {
		pageFields.RemoveChild(child)
	}

	for _, name := range p.pageFields {
		idx, ok := fieldMap[name]
		if !ok {
			continue
		}
		pageFields.AddPageField(idx)
	}
}

// updateDataFieldsArea updates the data fields area.
func (p *PivotTable) updateDataFieldsArea(
	ptDef *elements.PivotTableDefinition,
	fieldMap map[string]int32,
) {
	if len(p.dataFields) == 0 {
		if df := ptDef.DataFields(); df != nil {
			ptDef.RemoveChild(df)
		}

		return
	}

	dataFields := ptDef.GetOrCreateDataFields()
	dataFields.SetCount(uint32(len(p.dataFields)))

	// Clear existing data fields
	for child := range dataFields.Children() {
		dataFields.RemoveChild(child)
	}

	for _, df := range p.dataFields {
		idx, ok := fieldMap[df.name]
		if !ok {
			continue
		}
		dataField := dataFields.AddDataField(
			uint32(idx),
		)
		dataField.SetName(df.name)
		dataField.SetSubtotal(
			string(df.aggregate),
		)
		// Set the field index (fld attribute) to reference the base field
		dataField.SetFld(uint32(idx))
	}
}

// updateLocation updates the Location element with the calculated range and attributes.
func (p *PivotTable) updateLocation(
	ptDef *elements.PivotTableDefinition,
	_ map[string]int32,
) {
	loc := ptDef.GetOrCreateLocation()

	// Calculate the pivot table range based on destination cell and field counts
	// The range should encompass the pivot table structure
	rowFieldCount := len(p.rowFields)
	colFieldCount := len(p.colFields)
	dataFieldCount := len(p.dataFields)

	// Default to at least 2x2 range
	endCol := p.destCell.Col + 1
	endRow := p.destCell.Row + 1

	// Expand range based on field counts
	if colFieldCount > 0 {
		endCol = p.destCell.Col + colFieldCount + 1
	}
	if rowFieldCount > 0 {
		endRow = p.destCell.Row + rowFieldCount + 2
	}
	if dataFieldCount > 1 {
		// Multiple data fields add extra columns
		endCol = p.destCell.Col + dataFieldCount
	}

	// Build range reference
	endCell := CellRef{
		Col: endCol,
		Row: endRow,
	}
	rangeRef := p.destCell.String() + ":" + endCell.String()
	loc.SetRef(rangeRef)

	// Set firstHeaderRow (typically 1 for standard pivot tables)
	loc.SetFirstHeaderRow(1)

	// Set firstDataRow (row where data starts, typically 1)
	loc.SetFirstDataRow(1)

	// Set firstDataCol (column where data starts, typically 1)
	loc.SetFirstDataCol(1)
}

// ensureCacheRelationship ensures that the cacheId attribute is set
// and the relationship exists between pivot table and cache definition.
func (p *PivotTable) ensureCacheRelationship() error {
	ptDef := p.pivotPart.PivotTableDefinition()
	if ptDef == nil {
		return nil
	}

	// Get or create cache definition part
	cachePart, err := p.getCacheDefinitionPart()
	if err != nil {
		return err
	}

	// Get the relationship ID from the cache part
	relID := cachePart.RelationshipID()
	if relID == "" {
		// Cache part doesn't have a relationship ID yet
		// This is handled by the part creation logic
		return nil
	}

	// Set the cacheId attribute on the pivot table definition
	// The cacheId is typically 0 for the first cache, incrementing for additional caches
	// For simplicity, we'll use 0 as the default
	ptDef.SetCacheId(0)

	return nil
}
