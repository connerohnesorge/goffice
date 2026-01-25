// This file implements tests for PivotTable error handling and validation.
package spreadsheet

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// createTestDocument creates a temporary test document with one sheet.
func createTestDocument(
	t *testing.T,
) (*Document, *Sheet) {
	t.Helper()
	tmpFile := filepath.Join(
		t.TempDir(),
		"test.xlsx",
	)
	doc, err := Create(tmpFile, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	t.Cleanup(func() {
		_ = doc.Close()
		_ = os.Remove(tmpFile)
	})

	sheet, err := doc.AddSheet("Sheet1")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	return doc, sheet
}

// TestTryAddRowField_FieldNotFound tests that TryAddRowField returns
// ErrFieldNotFound when the field doesn't exist in the source data.
func TestTryAddRowField_FieldNotFound(
	t *testing.T,
) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range with headers: Name, Value
	sheet.Cell("A1").SetString("Name")
	sheet.Cell("B1").SetString("Value")
	sheet.Cell("A2").SetString("Item1")
	sheet.Cell("B2").SetNumber(100)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Try to add a non-existent field
	_, err = pt.TryAddRowField("NonExistentField")
	if err == nil {
		t.Error(
			"Expected error when adding non-existent field, got nil",
		)
	}

	if !errors.Is(err, ErrFieldNotFound) {
		t.Errorf(
			"Expected ErrFieldNotFound, got: %v",
			err,
		)
	}
}

// TestTryAddColumnField_FieldNotFound tests that TryAddColumnField returns
// ErrFieldNotFound when the field doesn't exist in the source data.
func TestTryAddColumnField_FieldNotFound(
	t *testing.T,
) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Category")
	sheet.Cell("B1").SetString("Amount")
	sheet.Cell("A2").SetString("A")
	sheet.Cell("B2").SetNumber(50)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Try to add a non-existent field
	_, err = pt.TryAddColumnField("InvalidField")
	if err == nil {
		t.Error(
			"Expected error when adding non-existent field, got nil",
		)
	}

	if !errors.Is(err, ErrFieldNotFound) {
		t.Errorf(
			"Expected ErrFieldNotFound, got: %v",
			err,
		)
	}
}

// TestTryAddDataField_FieldNotFound tests that TryAddDataField returns
// ErrFieldNotFound when the field doesn't exist in the source data.
func TestTryAddDataField_FieldNotFound(
	t *testing.T,
) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Product")
	sheet.Cell("B1").SetString("Sales")
	sheet.Cell("A2").SetString("Widget")
	sheet.Cell("B2").SetNumber(200)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Try to add a non-existent field
	_, err = pt.TryAddDataField(
		"NoSuchField",
		AggregateSUM,
	)
	if err == nil {
		t.Error(
			"Expected error when adding non-existent field, got nil",
		)
	}

	if !errors.Is(err, ErrFieldNotFound) {
		t.Errorf(
			"Expected ErrFieldNotFound, got: %v",
			err,
		)
	}
}

// TestTryAddPageField_FieldNotFound tests that TryAddPageField returns
// ErrFieldNotFound when the field doesn't exist in the source data.
func TestTryAddPageField_FieldNotFound(
	t *testing.T,
) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Region")
	sheet.Cell("B1").SetString("Revenue")
	sheet.Cell("A2").SetString("North")
	sheet.Cell("B2").SetNumber(1000)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Try to add a non-existent field
	_, err = pt.TryAddPageField("MissingField")
	if err == nil {
		t.Error(
			"Expected error when adding non-existent field, got nil",
		)
	}

	if !errors.Is(err, ErrFieldNotFound) {
		t.Errorf(
			"Expected ErrFieldNotFound, got: %v",
			err,
		)
	}
}

// TestTryAddRowField_DuplicateField tests that TryAddRowField returns
// ErrDuplicateField when the field is already on an axis.
func TestTryAddRowField_DuplicateField(
	t *testing.T,
) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Name")
	sheet.Cell("B1").SetString("Value")
	sheet.Cell("A2").SetString("Item1")
	sheet.Cell("B2").SetNumber(100)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Add field to row area first
	_, err = pt.TryAddRowField("Name")
	if err != nil {
		t.Fatalf(
			"Failed to add field first time: %v",
			err,
		)
	}

	// Try to add the same field again
	_, err = pt.TryAddRowField("Name")
	if err == nil {
		t.Error(
			"Expected error when adding duplicate field, got nil",
		)
	}

	if !errors.Is(err, ErrDuplicateField) {
		t.Errorf(
			"Expected ErrDuplicateField, got: %v",
			err,
		)
	}
}

// TestTryAddColumnField_DuplicateField tests that TryAddColumnField returns
// ErrDuplicateField when the field is already on another axis.
func TestTryAddColumnField_DuplicateField(
	t *testing.T,
) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Category")
	sheet.Cell("B1").SetString("Amount")
	sheet.Cell("A2").SetString("A")
	sheet.Cell("B2").SetNumber(50)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Add field to row area first
	_, err = pt.TryAddRowField("Category")
	if err != nil {
		t.Fatalf(
			"Failed to add field to row area: %v",
			err,
		)
	}

	// Try to add the same field to column area
	_, err = pt.TryAddColumnField("Category")
	if err == nil {
		t.Error(
			"Expected error when adding duplicate field, got nil",
		)
	}

	if !errors.Is(err, ErrDuplicateField) {
		t.Errorf(
			"Expected ErrDuplicateField, got: %v",
			err,
		)
	}
}

// TestTryAddDataField_DuplicateField tests that TryAddDataField returns
// ErrDuplicateField when the field is already on another axis.
func TestTryAddDataField_DuplicateField(
	t *testing.T,
) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Product")
	sheet.Cell("B1").SetString("Sales")
	sheet.Cell("A2").SetString("Widget")
	sheet.Cell("B2").SetNumber(200)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Add field to column area first
	_, err = pt.TryAddColumnField("Sales")
	if err != nil {
		t.Fatalf(
			"Failed to add field to column area: %v",
			err,
		)
	}

	// Try to add the same field to data area
	_, err = pt.TryAddDataField(
		"Sales",
		AggregateSUM,
	)
	if err == nil {
		t.Error(
			"Expected error when adding duplicate field, got nil",
		)
	}

	if !errors.Is(err, ErrDuplicateField) {
		t.Errorf(
			"Expected ErrDuplicateField, got: %v",
			err,
		)
	}
}

// TestTryAddPageField_DuplicateField tests that TryAddPageField returns
// ErrDuplicateField when the field is already on another axis.
func TestTryAddPageField_DuplicateField(
	t *testing.T,
) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Region")
	sheet.Cell("B1").SetString("Revenue")
	sheet.Cell("A2").SetString("North")
	sheet.Cell("B2").SetNumber(1000)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Add field to data area first
	_, err = pt.TryAddDataField(
		"Revenue",
		AggregateAverage,
	)
	if err != nil {
		t.Fatalf(
			"Failed to add field to data area: %v",
			err,
		)
	}

	// Try to add the same field to page area
	_, err = pt.TryAddPageField("Revenue")
	if err == nil {
		t.Error(
			"Expected error when adding duplicate field, got nil",
		)
	}

	if !errors.Is(err, ErrDuplicateField) {
		t.Errorf(
			"Expected ErrDuplicateField, got: %v",
			err,
		)
	}
}

// TestTryAddRowField_Success tests that TryAddRowField succeeds when the field is valid.
func TestTryAddRowField_Success(t *testing.T) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Name")
	sheet.Cell("B1").SetString("Value")
	sheet.Cell("A2").SetString("Item1")
	sheet.Cell("B2").SetNumber(100)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Try to add a valid field
	result, err := pt.TryAddRowField("Name")
	if err != nil {
		t.Errorf(
			"Expected no error, got: %v",
			err,
		)
	}

	if result == nil {
		t.Error("Expected non-nil result")
	}

	// Verify the field was added
	fields := pt.RowFields()
	if len(fields) != 1 {
		t.Errorf(
			"Expected 1 row field, got %d",
			len(fields),
		)
	}
	if len(fields) > 0 && fields[0] != "Name" {
		t.Errorf(
			"Expected 'Name', got '%s'",
			fields[0],
		)
	}
}

// TestTryAddColumnField_Success tests that TryAddColumnField succeeds when the field is valid.
func TestTryAddColumnField_Success(t *testing.T) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Category")
	sheet.Cell("B1").SetString("Amount")
	sheet.Cell("A2").SetString("A")
	sheet.Cell("B2").SetNumber(50)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Try to add a valid field
	result, err := pt.TryAddColumnField(
		"Category",
	)
	if err != nil {
		t.Errorf(
			"Expected no error, got: %v",
			err,
		)
	}

	if result == nil {
		t.Error("Expected non-nil result")
	}

	// Verify the field was added
	fields := pt.ColumnFields()
	if len(fields) != 1 {
		t.Errorf(
			"Expected 1 column field, got %d",
			len(fields),
		)
	}
	if len(fields) > 0 &&
		fields[0] != "Category" {
		t.Errorf(
			"Expected 'Category', got '%s'",
			fields[0],
		)
	}
}

// TestTryAddDataField_Success tests that TryAddDataField succeeds when the field is valid.
func TestTryAddDataField_Success(t *testing.T) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Product")
	sheet.Cell("B1").SetString("Sales")
	sheet.Cell("A2").SetString("Widget")
	sheet.Cell("B2").SetNumber(200)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Try to add a valid field
	result, err := pt.TryAddDataField(
		"Sales",
		AggregateSUM,
	)
	if err != nil {
		t.Errorf(
			"Expected no error, got: %v",
			err,
		)
	}

	if result == nil {
		t.Error("Expected non-nil result")
	}

	// Verify the field was added
	// Note: We can't directly access dataFields in tests, but we can verify no error occurred
}

// TestTryAddPageField_Success tests that TryAddPageField succeeds when the field is valid.
func TestTryAddPageField_Success(t *testing.T) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Region")
	sheet.Cell("B1").SetString("Revenue")
	sheet.Cell("A2").SetString("North")
	sheet.Cell("B2").SetNumber(1000)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Try to add a valid field
	result, err := pt.TryAddPageField("Region")
	if err != nil {
		t.Errorf(
			"Expected no error, got: %v",
			err,
		)
	}

	if result == nil {
		t.Error("Expected non-nil result")
	}

	// Verify the field was added
	fields := pt.PageFields()
	if len(fields) != 1 {
		t.Errorf(
			"Expected 1 page field, got %d",
			len(fields),
		)
	}
	if len(fields) > 0 && fields[0] != "Region" {
		t.Errorf(
			"Expected 'Region', got '%s'",
			fields[0],
		)
	}
}

// TestContainsField tests the containsField helper method.
func TestContainsField(t *testing.T) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("Field1")
	sheet.Cell("B1").SetString("Field2")
	sheet.Cell("C1").SetString("Field3")
	sheet.Cell("D1").SetString("Field4")
	sheet.Cell("A2").SetString("A")
	sheet.Cell("B2").SetString("B")
	sheet.Cell("C2").SetNumber(100)
	sheet.Cell("D2").SetNumber(200)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:D2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Add fields to different areas
	pt.AddRowField("Field1")
	pt.AddColumnField("Field2")
	pt.AddDataField("Field3", AggregateSUM)
	pt.AddPageField("Field4")

	// Test containsField
	tests := []struct {
		name     string
		field    string
		expected bool
	}{
		{"Field in row area", "Field1", true},
		{"Field in column area", "Field2", true},
		{"Field in data area", "Field3", true},
		{"Field in page area", "Field4", true},
		{
			"Field not assigned",
			"NonExistent",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pt.containsField(tt.field)
			if result != tt.expected {
				t.Errorf(
					"containsField(%q) = %v, expected %v",
					tt.field,
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestValidateFieldName tests the validateFieldName helper method.
func TestValidateFieldName(t *testing.T) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("ValidField1")
	sheet.Cell("B1").SetString("ValidField2")
	sheet.Cell("A2").SetString("Value1")
	sheet.Cell("B2").SetString("Value2")

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Test valid field names
	err = pt.validateFieldName("ValidField1")
	if err != nil {
		t.Errorf(
			"validateFieldName('ValidField1') returned error: %v",
			err,
		)
	}

	err = pt.validateFieldName("ValidField2")
	if err != nil {
		t.Errorf(
			"validateFieldName('ValidField2') returned error: %v",
			err,
		)
	}

	// Test invalid field name
	err = pt.validateFieldName("InvalidField")
	if err == nil {
		t.Error(
			"validateFieldName('InvalidField') should return error, got nil",
		)
	}
	if !errors.Is(err, ErrFieldNotFound) {
		t.Errorf(
			"Expected ErrFieldNotFound, got: %v",
			err,
		)
	}

	// Test empty field name
	err = pt.validateFieldName("")
	if err == nil {
		t.Error(
			"validateFieldName('') should return error, got nil",
		)
	}
	if !errors.Is(err, ErrFieldNotFound) {
		t.Errorf(
			"Expected ErrFieldNotFound, got: %v",
			err,
		)
	}
}

// TestValidateFieldName_WithCache tests validateFieldName with populated cache.
func TestValidateFieldName_WithCache(
	t *testing.T,
) {
	_, sheet := createTestDocument(t)

	// Create a simple source data range
	sheet.Cell("A1").SetString("CachedField1")
	sheet.Cell("B1").SetString("CachedField2")
	sheet.Cell("A2").SetString("Data1")
	sheet.Cell("B2").SetString("Data2")

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Populate the cache
	err = pt.Refresh()
	if err != nil {
		t.Fatalf(
			"Failed to refresh pivot table: %v",
			err,
		)
	}

	// Now test validation with cache populated
	err = pt.validateFieldName("CachedField1")
	if err != nil {
		t.Errorf(
			"validateFieldName('CachedField1') with cache returned error: %v",
			err,
		)
	}

	err = pt.validateFieldName("CachedField2")
	if err != nil {
		t.Errorf(
			"validateFieldName('CachedField2') with cache returned error: %v",
			err,
		)
	}

	// Test invalid field name with cache
	err = pt.validateFieldName("NotInCache")
	if err == nil {
		t.Error(
			"validateFieldName('NotInCache') should return error, got nil",
		)
	}
	if !errors.Is(err, ErrFieldNotFound) {
		t.Errorf(
			"Expected ErrFieldNotFound, got: %v",
			err,
		)
	}
}
