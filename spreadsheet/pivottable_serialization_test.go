// This file tests PivotTable XML serialization.
package spreadsheet

import (
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// getTempFilePath creates a temporary file path for tests.
func getTempFilePath(
	t *testing.T,
	filename string,
) string {
	t.Helper()

	return filepath.Join(t.TempDir(), filename)
}

// TestPivotFieldMetadataSerialization tests that field metadata is serialized correctly.
func TestPivotFieldMetadataSerialization(
	t *testing.T,
) {
	doc, sheet := createTestDocument(t)

	// Create source data with headers
	sheet.Cell("A1").SetString("Region")
	sheet.Cell("B1").SetString("Product")
	sheet.Cell("C1").SetString("Sales")
	sheet.Cell("A2").SetString("East")
	sheet.Cell("B2").SetString("Widget")
	sheet.Cell("C2").SetNumber(100)
	sheet.Cell("A3").SetString("West")
	sheet.Cell("B3").SetString("Gadget")
	sheet.Cell("C3").SetNumber(200)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:C3",
		CellRef{Col: 0, Row: 5},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Add fields to different axes
	pt.AddRowField("Region")
	pt.AddColumnField("Product")
	pt.AddDataField("Sales", AggregateSUM)

	// Refresh to populate cache
	if err := pt.Refresh(); err != nil {
		t.Fatalf(
			"Failed to refresh pivot table: %v",
			err,
		)
	}

	// Get pivot table definition
	ptDef := pt.pivotPart.PivotTableDefinition()
	if ptDef == nil {
		t.Fatal("PivotTableDefinition is nil")
	}

	// Get pivot fields collection
	pivotFields := ptDef.PivotFields()
	if pivotFields == nil {
		t.Fatal("PivotFields collection is nil")
	}

	// Verify field count
	if pivotFields.Count() != 3 {
		t.Errorf(
			"Expected 3 fields, got %d",
			pivotFields.Count(),
		)
	}

	// Verify each field has metadata
	fieldIndex := 0
	for field := range pivotFields.PivotFields() {
		fieldName := field.Name()
		if fieldName == "" {
			t.Errorf(
				"Field %d has empty name",
				fieldIndex,
			)
		}

		// Verify ShowAll attribute is set
		if !field.ShowAll() {
			t.Errorf(
				"Field %s: ShowAll should be true by default",
				fieldName,
			)
		}

		// Verify Compact attribute is set
		if !field.Compact() {
			t.Errorf(
				"Field %s: Compact should be true by default",
				fieldName,
			)
		}

		// Verify ShowDropDowns attribute is set
		if !field.ShowDropDowns() {
			t.Errorf(
				"Field %s: ShowDropDowns should be true by default",
				fieldName,
			)
		}

		// Verify DefaultSubtotal attribute is set
		if !field.DefaultSubtotal() {
			t.Errorf(
				"Field %s: DefaultSubtotal should be true by default",
				fieldName,
			)
		}

		// Verify axis is set correctly
		axis := field.Axis()
		switch fieldName {
		case "Region":
			if axis != "axisRow" {
				t.Errorf(
					"Field %s: expected axis 'axisRow', got '%s'",
					fieldName,
					axis,
				)
			}
		case "Product":
			if axis != "axisCol" {
				t.Errorf(
					"Field %s: expected axis 'axisCol', got '%s'",
					fieldName,
					axis,
				)
			}
		case "Sales":
			if axis != "axisValues" {
				t.Errorf(
					"Field %s: expected axis 'axisValues', got '%s'",
					fieldName,
					axis,
				)
			}
			if !field.DataField() {
				t.Errorf(
					"Field %s: DataField should be true",
					fieldName,
				)
			}
		}

		fieldIndex++
	}

	// Save and verify the document can be written
	if err := doc.SaveAs(getTempFilePath(t, "pivot_field_metadata.xlsx")); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}
}

// TestPivotFieldIndicesSerialization tests that field indices are correct.
func TestPivotFieldIndicesSerialization(
	t *testing.T,
) {
	doc, sheet := createTestDocument(t)

	// Create source data
	sheet.Cell("A1").SetString("Category")
	sheet.Cell("B1").SetString("SubCategory")
	sheet.Cell("C1").SetString("Amount")
	sheet.Cell("A2").SetString("Food")
	sheet.Cell("B2").SetString("Fruit")
	sheet.Cell("C2").SetNumber(50)

	// Create pivot table
	pt, err := sheet.AddPivotTable(
		"A1:C2",
		CellRef{Col: 0, Row: 4},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	// Add fields
	pt.AddRowField("Category")
	pt.AddRowField("SubCategory")
	pt.AddDataField("Amount", AggregateSUM)

	// Refresh to populate cache
	if err := pt.Refresh(); err != nil {
		t.Fatalf(
			"Failed to refresh pivot table: %v",
			err,
		)
	}

	// Get pivot table definition
	ptDef := pt.pivotPart.PivotTableDefinition()
	if ptDef == nil {
		t.Fatal("PivotTableDefinition is nil")
	}

	// Verify RowFields indices
	rowFields := ptDef.RowFields()
	if rowFields == nil {
		t.Fatal("RowFields is nil")
	}

	if rowFields.Count() != 2 {
		t.Errorf(
			"Expected 2 row fields, got %d",
			rowFields.Count(),
		)
	}

	// Verify indices point to correct fields
	expectedIndices := []int32{
		0,
		1,
	} // Category=0, SubCategory=1
	actualIndices := make([]int32, 0)
	for field := range rowFields.Fields() {
		actualIndices = append(
			actualIndices,
			field.X(),
		)
	}

	for i, expected := range expectedIndices {
		if i >= len(actualIndices) {
			t.Errorf(
				"Missing row field at index %d",
				i,
			)

			continue
		}
		if actualIndices[i] != expected {
			t.Errorf(
				"Row field %d: expected index %d, got %d",
				i,
				expected,
				actualIndices[i],
			)
		}
	}

	// Verify DataFields
	dataFields := ptDef.DataFields()
	if dataFields == nil {
		t.Fatal("DataFields is nil")
	}

	if dataFields.Count() != 1 {
		t.Errorf(
			"Expected 1 data field, got %d",
			dataFields.Count(),
		)
	}

	// Verify data field has correct fld index and subtotal
	for dataField := range dataFields.DataFields() {
		fldIndex := dataField.Fld()
		if fldIndex != 2 { // Amount is the 3rd field (index 2)
			t.Errorf(
				"DataField: expected fld=2, got %d",
				fldIndex,
			)
		}

		subtotal := dataField.Subtotal()
		if subtotal != "sum" {
			t.Errorf(
				"DataField: expected subtotal='sum', got '%s'",
				subtotal,
			)
		}
	}

	// Save and verify
	if err := doc.SaveAs(getTempFilePath(t, "pivot_field_indices.xlsx")); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}
}

// TestPivotCacheRelationship tests that cache relationship exists.
func TestPivotCacheRelationship(t *testing.T) {
	doc, sheet := createTestDocument(t)

	// Create source data
	sheet.Cell("A1").SetString("Item")
	sheet.Cell("B1").SetString("Value")
	sheet.Cell("A2").SetString("A")
	sheet.Cell("B2").SetNumber(10)

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

	pt.AddDataField("Value", AggregateSUM)

	// Refresh to populate cache
	if err := pt.Refresh(); err != nil {
		t.Fatalf(
			"Failed to refresh pivot table: %v",
			err,
		)
	}

	// Verify cache definition part exists
	cachePart := pt.pivotPart.PivotTableCacheDefinitionPart()
	if cachePart == nil {
		t.Fatal(
			"PivotTableCacheDefinitionPart is nil",
		)
	}

	// Verify relationship ID exists
	relID := cachePart.RelationshipID()
	if relID == "" {
		t.Error(
			"Cache definition part has no relationship ID",
		)
	}

	// Verify cacheId is set on pivot table definition
	ptDef := pt.pivotPart.PivotTableDefinition()
	if ptDef == nil {
		t.Fatal("PivotTableDefinition is nil")
	}

	cacheId := ptDef.CacheId()
	// cacheId should be set (typically 0 for first cache)
	if cacheId > 100 {
		t.Errorf(
			"CacheId seems invalid: %d",
			cacheId,
		)
	}

	// Save and verify
	if err := doc.SaveAs(getTempFilePath(t, "pivot_cache_relationship.xlsx")); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}
}

// TestPivotLocationSerialization tests that Location element has correct attributes.
func TestPivotLocationSerialization(
	t *testing.T,
) {
	doc, sheet := createTestDocument(t)

	// Create source data
	sheet.Cell("A1").SetString("X")
	sheet.Cell("B1").SetString("Y")
	sheet.Cell("A2").SetString("A")
	sheet.Cell("B2").SetNumber(5)

	// Create pivot table at specific location
	destCell := CellRef{Col: 4, Row: 10} // E11
	pt, err := sheet.AddPivotTable(
		"A1:B2",
		destCell,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	pt.AddRowField("X")
	pt.AddDataField("Y", AggregateSUM)

	// Refresh to populate cache
	if err := pt.Refresh(); err != nil {
		t.Fatalf(
			"Failed to refresh pivot table: %v",
			err,
		)
	}

	// Get pivot table definition
	ptDef := pt.pivotPart.PivotTableDefinition()
	if ptDef == nil {
		t.Fatal("PivotTableDefinition is nil")
	}

	// Verify Location element
	loc := ptDef.Location()
	if loc == nil {
		t.Fatal("Location element is nil")
	}

	// Verify ref attribute is set and starts with destination cell
	ref := loc.Ref()
	if ref == "" {
		t.Error("Location ref is empty")
	}
	expectedStart := destCell.String()
	if len(ref) < len(expectedStart) ||
		ref[:len(expectedStart)] != expectedStart {
		t.Errorf(
			"Location ref should start with '%s', got '%s'",
			expectedStart,
			ref,
		)
	}

	// Verify firstHeaderRow is set
	firstHeaderRow := loc.FirstHeaderRow()
	if firstHeaderRow != 1 {
		t.Errorf(
			"Expected firstHeaderRow=1, got %d",
			firstHeaderRow,
		)
	}

	// Verify firstDataRow is set
	firstDataRow := loc.FirstDataRow()
	if firstDataRow != 1 {
		t.Errorf(
			"Expected firstDataRow=1, got %d",
			firstDataRow,
		)
	}

	// Verify firstDataCol is set
	firstDataCol := loc.FirstDataCol()
	if firstDataCol != 1 {
		t.Errorf(
			"Expected firstDataCol=1, got %d",
			firstDataCol,
		)
	}

	// Save and verify
	if err := doc.SaveAs(getTempFilePath(t, "pivot_location.xlsx")); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}
}

// TestPivotRoundtrip tests create -> save -> load -> verify structure.
//
// This test verifies complete roundtrip functionality for pivot tables:
// - XML namespace handling (xmlns:prefix attributes correctly skipped)
// - Text content preservation in leaf elements
// - SharedStrings and cell values roundtrip
// - PivotTablePart loading after document reopening
// - Recursive child part loading for embedded part types
//
//nolint:revive // cyclomatic: comprehensive roundtrip verification test
func TestPivotRoundtrip(t *testing.T) {
	// Create initial document
	doc1, sheet1 := createTestDocument(t)

	// Create source data
	sheet1.Cell("A1").SetString("Name")
	sheet1.Cell("B1").SetString("Score")
	sheet1.Cell("C1").SetString("Grade")
	sheet1.Cell("A2").SetString("Alice")
	sheet1.Cell("B2").SetNumber(95)
	sheet1.Cell("C2").SetString("A")
	sheet1.Cell("A3").SetString("Bob")
	sheet1.Cell("B3").SetNumber(85)
	sheet1.Cell("C3").SetString("B")

	// Create pivot table with multiple field types
	pt1, err := sheet1.AddPivotTable(
		"A1:C3",
		CellRef{Col: 0, Row: 5},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create pivot table: %v",
			err,
		)
	}

	pt1.AddRowField("Name")
	pt1.AddColumnField("Grade")
	pt1.AddDataField("Score", AggregateAverage)
	pt1.AddPageField("Grade")

	// Refresh to populate cache
	if err := pt1.Refresh(); err != nil {
		t.Fatalf(
			"Failed to refresh pivot table: %v",
			err,
		)
	}

	// Save to file
	tmpFile := getTempFilePath(
		t,
		"pivot_roundtrip.xlsx",
	)
	if err := doc1.SaveAs(tmpFile); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}

	// Close first document
	if err := doc1.Close(); err != nil {
		t.Fatalf(
			"Failed to close document: %v",
			err,
		)
	}

	// Open the saved document
	doc2, err := Open(tmpFile, true)
	if err != nil {
		t.Fatalf(
			"Failed to open saved document: %v",
			err,
		)
	}
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close reopened document: %v",
				err,
			)
		}
	}()

	// Get the sheet (first sheet)
	var sheet2 *Sheet
	for s := range doc2.Sheets() {
		sheet2 = s

		break
	}
	if sheet2 == nil {
		t.Fatal(
			"No sheets found in reopened document",
		)
	}

	// Verify source data is intact
	if sheet2.Cell("A1").GetString() != "Name" {
		t.Error(
			"Source data header A1 not preserved",
		)
	}
	if sheet2.Cell("B2").GetNumber() != 95 {
		t.Error(
			"Source data value B2 not preserved",
		)
	}

	// Verify pivot table part exists
	// Get worksheet part
	wsPart := sheet2.WorksheetPart()
	if wsPart == nil {
		t.Fatal("Worksheet part is nil")
	}

	// Find pivot table part
	var pivotPart *parts.PivotTablePart
	for part := range wsPart.Parts() {
		if pp, ok := part.(*parts.PivotTablePart); ok {
			pivotPart = pp

			break
		}
	}

	if pivotPart == nil {
		t.Fatal(
			"PivotTablePart not found after roundtrip",
		)
	}

	// Verify pivot table definition
	ptDef2 := pivotPart.PivotTableDefinition()
	if ptDef2 == nil {
		t.Fatal(
			"PivotTableDefinition is nil after roundtrip",
		)
	}

	// Verify fields collection
	pivotFields2 := ptDef2.PivotFields()
	if pivotFields2 == nil {
		t.Fatal(
			"PivotFields is nil after roundtrip",
		)
	}

	// Should have 3 unique fields: Name, Score, Grade
	if pivotFields2.Count() != 3 {
		t.Errorf(
			"Expected 3 fields after roundtrip, got %d",
			pivotFields2.Count(),
		)
	}

	// Verify row fields
	rowFields2 := ptDef2.RowFields()
	if rowFields2 == nil {
		t.Fatal(
			"RowFields is nil after roundtrip",
		)
	}
	if rowFields2.Count() != 1 {
		t.Errorf(
			"Expected 1 row field after roundtrip, got %d",
			rowFields2.Count(),
		)
	}

	// Verify column fields
	colFields2 := ptDef2.ColFields()
	if colFields2 == nil {
		t.Fatal(
			"ColFields is nil after roundtrip",
		)
	}
	if colFields2.Count() != 1 {
		t.Errorf(
			"Expected 1 column field after roundtrip, got %d",
			colFields2.Count(),
		)
	}

	// Verify data fields
	dataFields2 := ptDef2.DataFields()
	if dataFields2 == nil {
		t.Fatal(
			"DataFields is nil after roundtrip",
		)
	}
	if dataFields2.Count() != 1 {
		t.Errorf(
			"Expected 1 data field after roundtrip, got %d",
			dataFields2.Count(),
		)
	}

	// Verify page fields
	pageFields2 := ptDef2.PageFields()
	if pageFields2 == nil {
		t.Fatal(
			"PageFields is nil after roundtrip",
		)
	}
	if pageFields2.Count() != 1 {
		t.Errorf(
			"Expected 1 page field after roundtrip, got %d",
			pageFields2.Count(),
		)
	}

	// Verify cache relationship still exists
	cachePart2 := pivotPart.PivotTableCacheDefinitionPart()
	if cachePart2 == nil {
		t.Fatal(
			"PivotTableCacheDefinitionPart is nil after roundtrip",
		)
	}

	// Verify cache definition
	cacheDef2 := cachePart2.PivotCacheDefinition()
	if cacheDef2 == nil {
		t.Fatal(
			"PivotCacheDefinition is nil after roundtrip",
		)
	}

	// Verify cache fields
	cacheFields2 := cacheDef2.CacheFields()
	if cacheFields2 == nil {
		t.Fatal(
			"CacheFields is nil after roundtrip",
		)
	}
	if cacheFields2.Count() != 3 {
		t.Errorf(
			"Expected 3 cache fields after roundtrip, got %d",
			cacheFields2.Count(),
		)
	}

	// Save again to verify double-roundtrip
	tmpFile2 := getTempFilePath(
		t,
		"pivot_roundtrip2.xlsx",
	)
	if err := doc2.SaveAs(tmpFile2); err != nil {
		t.Fatalf(
			"Failed to save document second time: %v",
			err,
		)
	}
}
