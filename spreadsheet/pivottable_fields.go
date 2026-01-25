// This file implements field management for PivotTable.
package spreadsheet

// AddRowField adds a field to the row area.
func (p *PivotTable) AddRowField(
	name string,
) *PivotTable {
	p.rowFields = append(p.rowFields, name)
	p.updateFields()

	return p
}

// AddColumnField adds a field to the column area.
func (p *PivotTable) AddColumnField(
	name string,
) *PivotTable {
	p.colFields = append(p.colFields, name)
	p.updateFields()

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
	p.updateFields()

	return p
}

// AddPageField adds a field to the page (filter) area.
func (p *PivotTable) AddPageField(
	name string,
) *PivotTable {
	p.pageFields = append(p.pageFields, name)
	p.updateFields()

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

// RemoveRowField removes a field from the row area.
func (p *PivotTable) RemoveRowField(
	name string,
) bool {
	for i, f := range p.rowFields {
		if f != name {
			continue
		}
		p.rowFields = append(
			p.rowFields[:i],
			p.rowFields[i+1:]...,
		)
		p.updateFields()

		return true
	}

	return false
}

// RemoveColumnField removes a field from the column area.
func (p *PivotTable) RemoveColumnField(
	name string,
) bool {
	for i, f := range p.colFields {
		if f != name {
			continue
		}
		p.colFields = append(
			p.colFields[:i],
			p.colFields[i+1:]...,
		)
		p.updateFields()

		return true
	}

	return false
}

// RemoveDataField removes a field from the data area.
func (p *PivotTable) RemoveDataField(
	name string,
) bool {
	for i, f := range p.dataFields {
		if f.name != name {
			continue
		}
		p.dataFields = append(
			p.dataFields[:i],
			p.dataFields[i+1:]...,
		)
		p.updateFields()

		return true
	}

	return false
}

// RemovePageField removes a field from the page area.
func (p *PivotTable) RemovePageField(
	name string,
) bool {
	for i, f := range p.pageFields {
		if f != name {
			continue
		}
		p.pageFields = append(
			p.pageFields[:i],
			p.pageFields[i+1:]...,
		)
		p.updateFields()

		return true
	}

	return false
}

// TryAddRowField adds a field to the row area after validation.
// Returns an error if the field doesn't exist in the source data or is already assigned to an axis.
func (p *PivotTable) TryAddRowField(
	name string,
) (*PivotTable, error) {
	// Validate that the field exists in the source data
	if err := p.validateFieldName(name); err != nil {
		return nil, err
	}

	// Check if the field is already assigned to an axis
	if p.containsField(name) {
		return nil, ErrDuplicateField
	}

	// Field is valid and not a duplicate, add it
	p.AddRowField(name)

	return p, nil
}

// TryAddColumnField adds a field to the column area after validation.
// Returns an error if the field doesn't exist in the source data or is already assigned to an axis.
func (p *PivotTable) TryAddColumnField(
	name string,
) (*PivotTable, error) {
	// Validate that the field exists in the source data
	if err := p.validateFieldName(name); err != nil {
		return nil, err
	}

	// Check if the field is already assigned to an axis
	if p.containsField(name) {
		return nil, ErrDuplicateField
	}

	// Field is valid and not a duplicate, add it
	p.AddColumnField(name)

	return p, nil
}

// TryAddDataField adds a field to the data area with an aggregation function after validation.
// Returns an error if the field doesn't exist in the source data or is already assigned to an axis.
func (p *PivotTable) TryAddDataField(
	name string,
	aggregate AggregateFunction,
) (*PivotTable, error) {
	// Validate that the field exists in the source data
	if err := p.validateFieldName(name); err != nil {
		return nil, err
	}

	// Check if the field is already assigned to an axis
	if p.containsField(name) {
		return nil, ErrDuplicateField
	}

	// Field is valid and not a duplicate, add it
	p.AddDataField(name, aggregate)

	return p, nil
}

// TryAddPageField adds a field to the page (filter) area after validation.
// Returns an error if the field doesn't exist in the source data or is already assigned to an axis.
func (p *PivotTable) TryAddPageField(
	name string,
) (*PivotTable, error) {
	// Validate that the field exists in the source data
	if err := p.validateFieldName(name); err != nil {
		return nil, err
	}

	// Check if the field is already assigned to an axis
	if p.containsField(name) {
		return nil, ErrDuplicateField
	}

	// Field is valid and not a duplicate, add it
	p.AddPageField(name)

	return p, nil
}
