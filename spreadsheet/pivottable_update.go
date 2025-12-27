// Package spreadsheet provides SpreadsheetML support for Excel documents.
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
	fieldMap, allFields := p.buildFieldMap()

	// Update PivotFields collection
	updatePivotFieldsCollection(ptDef, allFields)

	// Update individual field areas
	p.updateRowFieldsArea(ptDef, fieldMap)
	p.updateColFieldsArea(ptDef, fieldMap)
	p.updatePageFieldsArea(ptDef, fieldMap)
	p.updateDataFieldsArea(ptDef, fieldMap)
}

// buildFieldMap creates a map of unique field names to their indices.
func (p *PivotTable) buildFieldMap() (map[string]int32, []string) {
	fieldMap := make(map[string]int32)
	fieldIndex := firstFieldIndex
	allFields := make([]string, 0)

	addField := func(name string) {
		if _, exists := fieldMap[name]; exists {
			return
		}
		fieldMap[name] = fieldIndex
		allFields = append(allFields, name)
		fieldIndex++
	}

	for _, name := range p.rowFields {
		addField(name)
	}
	for _, name := range p.colFields {
		addField(name)
	}
	for _, name := range p.pageFields {
		addField(name)
	}
	for _, df := range p.dataFields {
		addField(df.name)
	}

	return fieldMap, allFields
}

// updatePivotFieldsCollection updates the PivotFields collection.
func updatePivotFieldsCollection(
	ptDef *elements.PivotTableDefinition,
	allFields []string,
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
	}
}
