// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements the PivotField wrapper for the PivotTable API.
//
//nolint:revive // file-length-limit: comprehensive pivot field implementation
package spreadsheet

import (
	"fmt"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// AxisType represents the axis type for a pivot field.
type AxisType int

const (
	// AxisTypeNone indicates the field is not assigned to any axis.
	AxisTypeNone AxisType = iota
	// AxisTypeRow indicates the field is in the row area.
	AxisTypeRow
	// AxisTypeColumn indicates the field is in the column area.
	AxisTypeColumn
	// AxisTypePage indicates the field is in the page/filter area.
	AxisTypePage
	// AxisTypeData indicates the field is in the data/values area.
	AxisTypeData
)

// FieldDataType represents the data type of a pivot field.
type FieldDataType int

const (
	// FieldDataTypeString represents string data.
	FieldDataTypeString FieldDataType = iota
	// FieldDataTypeNumber represents numeric data.
	FieldDataTypeNumber
	// FieldDataTypeDate represents date data.
	FieldDataTypeDate
	// FieldDataTypeBoolean represents boolean data.
	FieldDataTypeBoolean
)

// PivotField represents a high-level wrapper around a pivot field.
type PivotField struct {
	pivot   *PivotTable
	element *elements.PivotField
	index   uint32
	name    string
}

// newPivotField creates a new PivotField wrapper.
func newPivotField(
	pivot *PivotTable,
	element *elements.PivotField,
	index uint32,
	name string,
) *PivotField {
	return &PivotField{
		pivot:   pivot,
		element: element,
		index:   index,
		name:    name,
	}
}

// Name returns the field name.
func (pf *PivotField) Name() string {
	return pf.name
}

// Index returns the field index.
func (pf *PivotField) Index() uint32 {
	return pf.index
}

// Axis determines the axis type from the parent PivotTable's field arrays.
func (pf *PivotField) Axis() AxisType {
	// Check if field is in row fields
	for _, f := range pf.pivot.rowFields {
		if f == pf.name {
			return AxisTypeRow
		}
	}

	// Check if field is in column fields
	for _, f := range pf.pivot.colFields {
		if f == pf.name {
			return AxisTypeColumn
		}
	}

	// Check if field is in page fields
	for _, f := range pf.pivot.pageFields {
		if f == pf.name {
			return AxisTypePage
		}
	}

	// Check if field is in data fields
	for _, f := range pf.pivot.dataFields {
		if f.name == pf.name {
			return AxisTypeData
		}
	}

	return AxisTypeNone
}

// ShowAll returns whether to show all items.
func (pf *PivotField) ShowAll() bool {
	if pf.element == nil {
		return true // Default is true
	}

	return pf.element.ShowAll()
}

// SetShowAll sets whether to show all items.
func (pf *PivotField) SetShowAll(value bool) {
	if pf.element == nil {
		return
	}
	pf.element.SetShowAll(value)
}

// Compact returns whether to use compact layout.
func (pf *PivotField) Compact() bool {
	if pf.element == nil {
		return true // Default is true
	}

	return pf.element.Compact()
}

// SetCompact sets whether to use compact layout.
func (pf *PivotField) SetCompact(value bool) {
	if pf.element == nil {
		return
	}
	pf.element.SetCompact(value)
}

// HideNewItems returns whether to hide new items.
func (pf *PivotField) HideNewItems() bool {
	if pf.element == nil {
		return false // Default is false
	}

	return pf.element.HideNewItems()
}

// SetHideNewItems sets whether to hide new items.
func (pf *PivotField) SetHideNewItems(
	value bool,
) {
	if pf.element == nil {
		return
	}
	pf.element.SetHideNewItems(value)
}

// ShowDropDowns returns whether to show filter dropdowns.
func (pf *PivotField) ShowDropDowns() bool {
	if pf.element == nil {
		return true // Default is true
	}

	return pf.element.ShowDropDowns()
}

// SetShowDropDowns sets whether to show filter dropdowns.
func (pf *PivotField) SetShowDropDowns(
	value bool,
) {
	if pf.element == nil {
		return
	}
	pf.element.SetShowDropDowns(value)
}

// DefaultSubtotal returns whether to show default subtotal.
func (pf *PivotField) DefaultSubtotal() bool {
	if pf.element == nil {
		return true // Default is true
	}

	return pf.element.DefaultSubtotal()
}

// SetDefaultSubtotal sets whether to show default subtotal.
func (pf *PivotField) SetDefaultSubtotal(
	value bool,
) {
	if pf.element == nil {
		return
	}
	pf.element.SetDefaultSubtotal(value)
}

// DataType infers the data type from cache data by examining the first non-nil value.
//
//nolint:revive // cognitive-complexity: function handles multiple data type detection paths
func (pf *PivotField) DataType() FieldDataType {
	// Get cache definition part
	cacheDefPart, err := pf.pivot.getCacheDefinitionPart()
	if err != nil {
		return FieldDataTypeString // Default to string
	}

	cacheDef := cacheDefPart.PivotCacheDefinition()
	if cacheDef == nil {
		return FieldDataTypeString
	}

	// Get the cache field by index
	cacheFields := cacheDef.CacheFields()
	if cacheFields == nil {
		return FieldDataTypeString
	}

	cacheField := cacheFields.GetField(
		int(pf.index),
	)
	if cacheField == nil {
		return FieldDataTypeString
	}

	// Get cache records part
	recordsPart, err := pf.pivot.getCacheRecordsPart()
	if err != nil {
		return FieldDataTypeString
	}

	cacheRecords := recordsPart.PivotCacheRecords()
	if cacheRecords == nil {
		return FieldDataTypeString
	}

	// Iterate through records to find first non-nil value for this field
	for record := range cacheRecords.Records() {
		// Get values in this record by iterating over children
		valueIndex := 0
		for child := range record.Children() {
			if uint32(valueIndex) == pf.index {
				// Found the value for this field
				// Determine type based on element type
				switch child.(type) {
				case *elements.RecordNumber:
					return FieldDataTypeNumber
				case *elements.RecordString:
					return FieldDataTypeString
				case *elements.RecordDateTime:
					return FieldDataTypeDate
				case *elements.RecordBoolean:
					return FieldDataTypeBoolean
				case *elements.RecordMissing:
					// Continue to next record
					continue
				default:
					// Try checking by local name as fallback
					switch child.LocalName() {
					case "n":
						return FieldDataTypeNumber
					case "s":
						return FieldDataTypeString
					case "d":
						return FieldDataTypeDate
					case "b":
						return FieldDataTypeBoolean
					}

					return FieldDataTypeString
				}
			}
			valueIndex++
		}
	}

	return FieldDataTypeString // Default to string
}

// ItemCount counts unique items in the cache for this field.
//
//nolint:revive // cognitive-complexity: function handles multiple data type conversions
func (pf *PivotField) ItemCount() int {
	// Get cache records part
	recordsPart, err := pf.pivot.getCacheRecordsPart()
	if err != nil {
		return 0
	}

	cacheRecords := recordsPart.PivotCacheRecords()
	if cacheRecords == nil {
		return 0
	}

	// Use a map to track unique values
	uniqueValues := make(map[string]bool)

	// Iterate through records to count unique values for this field
	for record := range cacheRecords.Records() {
		// Get values in this record by iterating over children
		valueIndex := 0
		for child := range record.Children() {
			if uint32(valueIndex) == pf.index {
				// Found the value for this field
				// Get a string representation for comparison
				var strValue string
				switch v := child.(type) {
				case *elements.RecordNumber:
					// Convert number to string
					num := v.V()
					strValue = fmt.Sprintf("%f", num)
				case *elements.RecordString:
					strValue = v.V()
				case *elements.RecordDateTime:
					strValue = v.V()
				case *elements.RecordBoolean:
					if v.V() {
						strValue = "true"
					} else {
						strValue = "false"
					}
				case *elements.RecordMissing:
					// Skip missing values
					continue
				default:
					// Try to get attribute value as fallback
					if attr, found := child.GetAttribute("v", ""); found {
						strValue = attr.Value()
					}
				}

				if strValue != "" {
					uniqueValues[strValue] = true
				}

				break
			}
			valueIndex++
		}
	}

	return len(uniqueValues)
}

// Field returns a specific field by name, or nil if not found.
func (p *PivotTable) Field(
	name string,
) *PivotField {
	if name == "" {
		return nil
	}

	// Check if pivot part exists
	if p.pivotPart == nil {
		return nil
	}

	// Get the pivot table definition
	ptDef := p.pivotPart.PivotTableDefinition()
	if ptDef == nil {
		return nil
	}

	// Get the pivot fields collection
	pivotFields := ptDef.PivotFields()
	if pivotFields == nil {
		return nil
	}

	// Find the field by name
	index := 0
	for field := range pivotFields.PivotFields() {
		fieldName := field.Name()
		if fieldName == "" {
			// Try to get name from cache definition
			headers, err := p.getCachedHeaders()
			if err == nil &&
				index < len(headers) {
				fieldName = headers[index]
			}
		}

		if fieldName == name {
			return newPivotField(
				p,
				field,
				uint32(index),
				name,
			)
		}
		index++
	}

	return nil
}

// Fields returns all fields in the pivot table.
func (p *PivotTable) Fields() []*PivotField {
	// Check if pivot part exists
	if p.pivotPart == nil {
		return nil
	}

	// Get the pivot table definition
	ptDef := p.pivotPart.PivotTableDefinition()
	if ptDef == nil {
		return nil
	}

	// Get the pivot fields collection
	pivotFields := ptDef.PivotFields()
	if pivotFields == nil {
		return nil
	}

	// Get headers from cache to use as field names
	headers, _ := p.getCachedHeaders()

	// Collect all fields
	var fields []*PivotField
	index := 0
	for field := range pivotFields.PivotFields() {
		fieldName := field.Name()
		if fieldName == "" &&
			index < len(headers) {
			// Use name from cache if not set on field
			fieldName = headers[index]
		}

		if fieldName != "" {
			fields = append(
				fields,
				newPivotField(
					p,
					field,
					uint32(index),
					fieldName,
				),
			)
		}
		index++
	}

	return fields
}
