package spreadsheet

import (
	"testing"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

const (
	testKeep1 = "Keep1"
	testKeep2 = "Keep2"
)

// TestInsertRows tests the InsertRows method.
//
//nolint:revive // cyclomatic: comprehensive test coverage for row insertion
func TestInsertRows(t *testing.T) {
	t.Run(
		"InsertAtBeginning",
		func(t *testing.T) {
			// Create a new document and sheet
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Add some data
			_ = sheet.SetCellValue("A1", "Header")
			_ = sheet.SetCellValue("A2", "Data1")
			_ = sheet.SetCellValue("A3", "Data2")

			// Insert 2 rows at the beginning
			if err := sheet.InsertRows(1, 2); err != nil {
				t.Fatalf(
					"InsertRows failed: %v",
					err,
				)
			}

			// Verify rows shifted down
			if val := sheet.GetCellValue("A3"); val != "Header" {
				t.Errorf(
					"expected A3='Header', got %q",
					val,
				)
			}
			if val := sheet.GetCellValue("A4"); val != "Data1" {
				t.Errorf(
					"expected A4='Data1', got %q",
					val,
				)
			}
			if val := sheet.GetCellValue("A5"); val != "Data2" {
				t.Errorf(
					"expected A5='Data2', got %q",
					val,
				)
			}
		},
	)

	t.Run("InsertInMiddle", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Add some data
		_ = sheet.SetCellValue("A1", "Row1")
		_ = sheet.SetCellValue("A2", "Row2")
		_ = sheet.SetCellValue("A3", "Row3")

		// Insert 1 row at index 2
		if err := sheet.InsertRows(2, 1); err != nil {
			t.Fatalf("InsertRows failed: %v", err)
		}

		// Verify rows shifted down from index 2
		if val := sheet.GetCellValue("A1"); val != "Row1" {
			t.Errorf(
				"expected A1='Row1', got %q",
				val,
			)
		}
		if val := sheet.GetCellValue("A3"); val != "Row2" {
			t.Errorf(
				"expected A3='Row2', got %q",
				val,
			)
		}
		if val := sheet.GetCellValue("A4"); val != "Row3" {
			t.Errorf(
				"expected A4='Row3', got %q",
				val,
			)
		}
	})

	t.Run("FormulaUpdate", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Add data and formula
		_ = sheet.SetCellValue("A1", 10)
		_ = sheet.SetCellValue("A2", 20)
		_ = sheet.SetCellFormula(
			"A3",
			"=SUM(A1:A2)",
		)

		// Insert row at index 2
		if err := sheet.InsertRows(2, 1); err != nil {
			t.Fatalf("InsertRows failed: %v", err)
		}

		// Verify formula updated
		formula := sheet.GetCellFormula("A4")
		if formula != "=SUM(A1:A3)" {
			t.Errorf(
				"expected formula '=SUM(A1:A3)', got %q",
				formula,
			)
		}
	})

	t.Run(
		"InvalidParameters",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Test invalid index
			if err := sheet.InsertRows(0, 1); err == nil {
				t.Error(
					"expected error for index 0, got nil",
				)
			}

			// Test invalid count
			if err := sheet.InsertRows(1, 0); err == nil {
				t.Error(
					"expected error for count 0, got nil",
				)
			}

			// Test exceeding max rows
			if err := sheet.InsertRows(MaxRow, 2); err == nil {
				t.Error(
					"expected error for exceeding max rows, got nil",
				)
			}
		},
	)
}

// TestDeleteRows tests the DeleteRows method.
func TestDeleteRows(t *testing.T) {
	t.Run(
		"DeleteFromBeginning",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Add some data
			_ = sheet.SetCellValue(
				"A1",
				"Delete1",
			)
			_ = sheet.SetCellValue(
				"A2",
				"Delete2",
			)
			_ = sheet.SetCellValue(
				"A3",
				testKeep1,
			)
			_ = sheet.SetCellValue(
				"A4",
				testKeep2,
			)

			// Delete first 2 rows
			if err := sheet.DeleteRows(1, 2); err != nil {
				t.Fatalf(
					"DeleteRows failed: %v",
					err,
				)
			}

			// Verify rows shifted up
			if val := sheet.GetCellValue("A1"); val != testKeep1 {
				t.Errorf(
					"expected A1='Keep1', got %q",
					val,
				)
			}
			if val := sheet.GetCellValue("A2"); val != testKeep2 {
				t.Errorf(
					"expected A2='Keep2', got %q",
					val,
				)
			}
		},
	)

	t.Run("DeleteInMiddle", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Add some data
		_ = sheet.SetCellValue("A1", testKeep1)
		_ = sheet.SetCellValue("A2", "Delete1")
		_ = sheet.SetCellValue("A3", "Delete2")
		_ = sheet.SetCellValue("A4", testKeep2)

		// Delete rows 2-3
		if err := sheet.DeleteRows(2, 2); err != nil {
			t.Fatalf("DeleteRows failed: %v", err)
		}

		// Verify rows shifted up
		if val := sheet.GetCellValue("A1"); val != testKeep1 {
			t.Errorf(
				"expected A1='Keep1', got %q",
				val,
			)
		}
		if val := sheet.GetCellValue("A2"); val != testKeep2 {
			t.Errorf(
				"expected A2='Keep2', got %q",
				val,
			)
		}
	})

	t.Run(
		"FormulaReferencesDeleted",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Add data and formula referencing deleted rows
			_ = sheet.SetCellValue("A1", 10)
			_ = sheet.SetCellValue("A2", 20)
			_ = sheet.SetCellFormula(
				"A3",
				"=A2*2",
			)

			// Delete row 2
			if err := sheet.DeleteRows(2, 1); err != nil {
				t.Fatalf(
					"DeleteRows failed: %v",
					err,
				)
			}

			// Verify formula contains #REF!
			formula := sheet.GetCellFormula("A2")
			if formula != "=#REF! * 2" &&
				formula != "=#REF!*2" {
				t.Errorf(
					"expected formula with #REF!, got %q",
					formula,
				)
			}
		},
	)

	t.Run(
		"InvalidParameters",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Test invalid index
			if err := sheet.DeleteRows(0, 1); err == nil {
				t.Error(
					"expected error for index 0, got nil",
				)
			}

			// Test invalid count
			if err := sheet.DeleteRows(1, 0); err == nil {
				t.Error(
					"expected error for count 0, got nil",
				)
			}
		},
	)
}

// TestInsertColumns tests the InsertColumns method.
func TestInsertColumns(t *testing.T) {
	t.Run(
		"InsertAtBeginning",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Add some data
			_ = sheet.SetCellValue("A1", "ColA")
			_ = sheet.SetCellValue("B1", "ColB")
			_ = sheet.SetCellValue("C1", "ColC")

			// Insert 2 columns at the beginning
			if err := sheet.InsertColumns(1, 2); err != nil {
				t.Fatalf(
					"InsertColumns failed: %v",
					err,
				)
			}

			// Verify columns shifted right
			if val := sheet.GetCellValue("C1"); val != "ColA" {
				t.Errorf(
					"expected C1='ColA', got %q",
					val,
				)
			}
			if val := sheet.GetCellValue("D1"); val != "ColB" {
				t.Errorf(
					"expected D1='ColB', got %q",
					val,
				)
			}
			if val := sheet.GetCellValue("E1"); val != "ColC" {
				t.Errorf(
					"expected E1='ColC', got %q",
					val,
				)
			}
		},
	)

	t.Run("InsertInMiddle", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Add some data
		_ = sheet.SetCellValue("A1", "Col1")
		_ = sheet.SetCellValue("B1", "Col2")
		_ = sheet.SetCellValue("C1", "Col3")

		// Insert 1 column at index 2 (column B)
		if err := sheet.InsertColumns(2, 1); err != nil {
			t.Fatalf(
				"InsertColumns failed: %v",
				err,
			)
		}

		// Verify columns shifted right from index 2
		if val := sheet.GetCellValue("A1"); val != "Col1" {
			t.Errorf(
				"expected A1='Col1', got %q",
				val,
			)
		}
		if val := sheet.GetCellValue("C1"); val != "Col2" {
			t.Errorf(
				"expected C1='Col2', got %q",
				val,
			)
		}
		if val := sheet.GetCellValue("D1"); val != "Col3" {
			t.Errorf(
				"expected D1='Col3', got %q",
				val,
			)
		}
	})

	t.Run(
		"InvalidParameters",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Test invalid index
			if err := sheet.InsertColumns(0, 1); err == nil {
				t.Error(
					"expected error for index 0, got nil",
				)
			}

			// Test invalid count
			if err := sheet.InsertColumns(1, 0); err == nil {
				t.Error(
					"expected error for count 0, got nil",
				)
			}

			// Test exceeding max columns
			if err := sheet.InsertColumns(MaxColumn, 2); err == nil {
				t.Error(
					"expected error for exceeding max columns, got nil",
				)
			}
		},
	)
}

// TestDeleteColumns tests the DeleteColumns method.
func TestDeleteColumns(t *testing.T) {
	t.Run(
		"DeleteFromBeginning",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Add some data
			_ = sheet.SetCellValue(
				"A1",
				"DeleteA",
			)
			_ = sheet.SetCellValue(
				"B1",
				"DeleteB",
			)
			_ = sheet.SetCellValue("C1", "KeepC")
			_ = sheet.SetCellValue("D1", "KeepD")

			// Delete first 2 columns
			if err := sheet.DeleteColumns(1, 2); err != nil {
				t.Fatalf(
					"DeleteColumns failed: %v",
					err,
				)
			}

			// Verify columns shifted left
			if val := sheet.GetCellValue("A1"); val != "KeepC" {
				t.Errorf(
					"expected A1='KeepC', got %q",
					val,
				)
			}
			if val := sheet.GetCellValue("B1"); val != "KeepD" {
				t.Errorf(
					"expected B1='KeepD', got %q",
					val,
				)
			}
		},
	)

	t.Run("DeleteInMiddle", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Add some data
		_ = sheet.SetCellValue("A1", testKeep1)
		_ = sheet.SetCellValue("B1", "Delete1")
		_ = sheet.SetCellValue("C1", "Delete2")
		_ = sheet.SetCellValue("D1", testKeep2)

		// Delete columns 2-3 (B-C)
		if err := sheet.DeleteColumns(2, 2); err != nil {
			t.Fatalf(
				"DeleteColumns failed: %v",
				err,
			)
		}

		// Verify columns shifted left
		if val := sheet.GetCellValue("A1"); val != testKeep1 {
			t.Errorf(
				"expected A1='Keep1', got %q",
				val,
			)
		}
		if val := sheet.GetCellValue("B1"); val != testKeep2 {
			t.Errorf(
				"expected B1='Keep2', got %q",
				val,
			)
		}
	})

	t.Run(
		"InvalidParameters",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Test invalid index
			if err := sheet.DeleteColumns(0, 1); err == nil {
				t.Error(
					"expected error for index 0, got nil",
				)
			}

			// Test invalid count
			if err := sheet.DeleteColumns(1, 0); err == nil {
				t.Error(
					"expected error for count 0, got nil",
				)
			}
		},
	)
}

// TestMergedCellsUpdate tests that merged cells are updated correctly.
func TestMergedCellsUpdate(t *testing.T) {
	t.Run(
		"InsertRowsAboveMerge",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Create a merged cell range
			if err := sheet.MergeCells("A2:B3"); err != nil {
				t.Fatalf(
					"MergeCells failed: %v",
					err,
				)
			}

			// Insert a row at index 1
			if err := sheet.InsertRows(1, 1); err != nil {
				t.Fatalf(
					"InsertRows failed: %v",
					err,
				)
			}

			// Check that merged cells were updated (A2:B3 should become A3:B4)
			ws := sheet.Worksheet()
			mergeCells := ws.MergeCells()
			if mergeCells == nil {
				t.Fatal(
					"expected merge cells to exist",
				)
			}

			// Verify the merge cell range was updated
			found := false
			for mc := range mergeCells.GetMergeCells() {
				ref := mc.Ref()
				if ref == "A3:B4" {
					found = true

					break
				}
			}
			if !found {
				t.Error(
					"expected merged cell range A3:B4 not found",
				)
			}
		},
	)
}

// TestNamedRangeUpdate tests that named ranges are updated correctly.
func TestNamedRangeUpdate(t *testing.T) {
	t.Run(
		"InsertRowsAffectingNamedRange",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Add a named range pointing to Sheet1!A2:A5
			wp := doc.WorkbookPart()
			if wp == nil {
				t.Fatal("workbook part is nil")
			}
			wb, ok := wp.Workbook().(*elements.Workbook)
			if !ok {
				t.Fatal(
					"workbook is not *elements.Workbook",
				)
			}
			definedNames := wb.GetOrCreateDefinedNames()
			definedName := definedNames.AddDefinedName(
				"TestRange",
				"Sheet1!A2:A5",
			)
			if definedName == nil {
				t.Fatal(
					"failed to add defined name",
				)
			}

			// Insert 2 rows at index 2
			if err := sheet.InsertRows(2, 2); err != nil {
				t.Fatalf(
					"InsertRows failed: %v",
					err,
				)
			}

			// Verify the named range was updated
			updatedFormula := definedName.Formula()
			// The formula should be updated to reflect the inserted rows
			// Original: Sheet1!A2:A5
			// After inserting 2 rows at index 2:
			// - A2 and after shift down, so A2 becomes A4
			// - A5 becomes A7
			// Result: Sheet1!A4:A7
			expected := "Sheet1!A4:A7"
			if updatedFormula != expected {
				t.Errorf(
					"expected named range formula %q, got %q",
					expected,
					updatedFormula,
				)
			}
		},
	)
}

// TestGroupRows tests the GroupRows method.
//
//nolint:revive // cyclomatic: comprehensive test coverage for grouping
func TestGroupRows(t *testing.T) {
	t.Run("BasicGrouping", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Group rows 2-5
		if err := sheet.GroupRows(2, 5); err != nil {
			t.Fatalf("GroupRows failed: %v", err)
		}

		// Verify outline levels are set
		sheetData := sheet.SheetData()
		for i := uint32(2); i <= 5; i++ {
			row := sheetData.GetOrCreateRow(i)
			if level := row.OutlineLevel(); level != 1 {
				t.Errorf(
					"expected row %d outline level 1, got %d",
					i,
					level,
				)
			}
		}

		// Verify row 1 has no outline level
		row1 := sheetData.GetOrCreateRow(1)
		if level := row1.OutlineLevel(); level != 0 {
			t.Errorf(
				"expected row 1 outline level 0, got %d",
				level,
			)
		}
	})

	t.Run("NestedGrouping", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Group rows 2-5 (level 1)
		if err := sheet.GroupRows(2, 5); err != nil {
			t.Fatalf("GroupRows failed: %v", err)
		}

		// Group rows 3-4 again (level 2)
		if err := sheet.GroupRows(3, 4); err != nil {
			t.Fatalf("GroupRows failed: %v", err)
		}

		// Verify outline levels
		sheetData := sheet.SheetData()

		row2 := sheetData.GetOrCreateRow(2)
		if level := row2.OutlineLevel(); level != 1 {
			t.Errorf(
				"expected row 2 outline level 1, got %d",
				level,
			)
		}

		row3 := sheetData.GetOrCreateRow(3)
		if level := row3.OutlineLevel(); level != 2 {
			t.Errorf(
				"expected row 3 outline level 2, got %d",
				level,
			)
		}

		row4 := sheetData.GetOrCreateRow(4)
		if level := row4.OutlineLevel(); level != 2 {
			t.Errorf(
				"expected row 4 outline level 2, got %d",
				level,
			)
		}

		row5 := sheetData.GetOrCreateRow(5)
		if level := row5.OutlineLevel(); level != 1 {
			t.Errorf(
				"expected row 5 outline level 1, got %d",
				level,
			)
		}
	})

	t.Run("MaxLevel", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Group same rows 8 times (should cap at level 7)
		for i := range 8 {
			if err := sheet.GroupRows(2, 3); err != nil {
				t.Fatalf(
					"GroupRows failed on iteration %d: %v",
					i,
					err,
				)
			}
		}

		// Verify level is capped at 7
		sheetData := sheet.SheetData()
		row2 := sheetData.GetOrCreateRow(2)
		if level := row2.OutlineLevel(); level != 7 {
			t.Errorf(
				"expected row 2 outline level 7 (capped), got %d",
				level,
			)
		}
	})

	t.Run(
		"InvalidParameters",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Test start > end
			if err := sheet.GroupRows(5, 2); err == nil {
				t.Error(
					"expected error for start > end, got nil",
				)
			}

			// Test invalid start
			if err := sheet.GroupRows(0, 5); err == nil {
				t.Error(
					"expected error for start = 0, got nil",
				)
			}

			// Test invalid end
			if err := sheet.GroupRows(1, MaxRow+1); err == nil {
				t.Error(
					"expected error for end > MaxRow, got nil",
				)
			}
		},
	)
}

// TestUngroupRows tests the UngroupRows method.
//
//nolint:revive // cyclomatic: comprehensive test coverage for ungrouping
func TestUngroupRows(t *testing.T) {
	t.Run("BasicUngrouping", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Group rows 2-5, then ungroup
		if err := sheet.GroupRows(2, 5); err != nil {
			t.Fatalf("GroupRows failed: %v", err)
		}

		if err := sheet.UngroupRows(2, 5); err != nil {
			t.Fatalf(
				"UngroupRows failed: %v",
				err,
			)
		}

		// Verify outline levels are removed
		sheetData := sheet.SheetData()
		for i := uint32(2); i <= 5; i++ {
			row := sheetData.GetOrCreateRow(i)
			if level := row.OutlineLevel(); level != 0 {
				t.Errorf(
					"expected row %d outline level 0, got %d",
					i,
					level,
				)
			}
		}
	})

	t.Run(
		"PartialUngrouping",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Group rows 2-5 twice (level 2)
			if err := sheet.GroupRows(2, 5); err != nil {
				t.Fatalf(
					"GroupRows failed: %v",
					err,
				)
			}
			if err := sheet.GroupRows(2, 5); err != nil {
				t.Fatalf(
					"GroupRows failed: %v",
					err,
				)
			}

			// Ungroup once (should go from level 2 to level 1)
			if err := sheet.UngroupRows(2, 5); err != nil {
				t.Fatalf(
					"UngroupRows failed: %v",
					err,
				)
			}

			// Verify outline levels are decremented to 1
			sheetData := sheet.SheetData()
			for i := uint32(2); i <= 5; i++ {
				row := sheetData.GetOrCreateRow(i)
				if level := row.OutlineLevel(); level != 1 {
					t.Errorf(
						"expected row %d outline level 1, got %d",
						i,
						level,
					)
				}
			}
		},
	)

	t.Run(
		"UngroupAtLevelZero",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Ungroup rows that are not grouped (should be no-op)
			if err := sheet.UngroupRows(2, 5); err != nil {
				t.Fatalf(
					"UngroupRows failed: %v",
					err,
				)
			}

			// Verify outline levels remain 0
			sheetData := sheet.SheetData()
			for i := uint32(2); i <= 5; i++ {
				row := sheetData.GetOrCreateRow(i)
				if level := row.OutlineLevel(); level != 0 {
					t.Errorf(
						"expected row %d outline level 0, got %d",
						i,
						level,
					)
				}
			}
		},
	)
}

// TestGroupColumns tests the GroupColumns method.
//
//nolint:revive // cyclomatic: comprehensive test coverage for grouping
func TestGroupColumns(t *testing.T) {
	t.Run("BasicGrouping", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Group columns 2-5 (B-E)
		if err := sheet.GroupColumns(2, 5); err != nil {
			t.Fatalf(
				"GroupColumns failed: %v",
				err,
			)
		}

		// Verify outline levels are set
		ws := sheet.Worksheet()
		cols := ws.Cols()
		if cols == nil {
			t.Fatal(
				"expected Cols element to exist",
			)
		}

		for i := uint32(2); i <= 5; i++ {
			col := cols.GetColByIndex(int(i))
			if col == nil {
				t.Fatalf(
					"expected Col element for column %d to exist",
					i,
				)
			}
			if level := col.OutlineLevel(); level != 1 {
				t.Errorf(
					"expected column %d outline level 1, got %d",
					i,
					level,
				)
			}
		}
	})

	t.Run("NestedGrouping", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Group columns 2-5 (level 1)
		if err := sheet.GroupColumns(2, 5); err != nil {
			t.Fatalf(
				"GroupColumns failed: %v",
				err,
			)
		}

		// Group columns 3-4 again (level 2)
		if err := sheet.GroupColumns(3, 4); err != nil {
			t.Fatalf(
				"GroupColumns failed: %v",
				err,
			)
		}

		// Verify outline levels
		ws := sheet.Worksheet()
		cols := ws.Cols()

		col2 := cols.GetColByIndex(2)
		if level := col2.OutlineLevel(); level != 1 {
			t.Errorf(
				"expected column 2 outline level 1, got %d",
				level,
			)
		}

		col3 := cols.GetColByIndex(3)
		if level := col3.OutlineLevel(); level != 2 {
			t.Errorf(
				"expected column 3 outline level 2, got %d",
				level,
			)
		}

		col4 := cols.GetColByIndex(4)
		if level := col4.OutlineLevel(); level != 2 {
			t.Errorf(
				"expected column 4 outline level 2, got %d",
				level,
			)
		}

		col5 := cols.GetColByIndex(5)
		if level := col5.OutlineLevel(); level != 1 {
			t.Errorf(
				"expected column 5 outline level 1, got %d",
				level,
			)
		}
	})

	t.Run("MaxLevel", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Group same columns 8 times (should cap at level 7)
		for i := range 8 {
			if err := sheet.GroupColumns(2, 3); err != nil {
				t.Fatalf(
					"GroupColumns failed on iteration %d: %v",
					i,
					err,
				)
			}
		}

		// Verify level is capped at 7
		ws := sheet.Worksheet()
		cols := ws.Cols()
		col2 := cols.GetColByIndex(2)
		if level := col2.OutlineLevel(); level != 7 {
			t.Errorf(
				"expected column 2 outline level 7 (capped), got %d",
				level,
			)
		}
	})

	t.Run(
		"InvalidParameters",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Test start > end
			if err := sheet.GroupColumns(5, 2); err == nil {
				t.Error(
					"expected error for start > end, got nil",
				)
			}

			// Test invalid start
			if err := sheet.GroupColumns(0, 5); err == nil {
				t.Error(
					"expected error for start = 0, got nil",
				)
			}

			// Test invalid end
			if err := sheet.GroupColumns(1, MaxColumn+1); err == nil {
				t.Error(
					"expected error for end > MaxColumn, got nil",
				)
			}
		},
	)
}

// TestUngroupColumns tests the UngroupColumns method.
//
//nolint:revive // cyclomatic: comprehensive test coverage for ungrouping
func TestUngroupColumns(t *testing.T) {
	t.Run("BasicUngrouping", func(t *testing.T) {
		doc, err := Create(
			t.TempDir()+"/test.xlsx",
			DocTypeWorkbook,
		)
		if err != nil {
			t.Fatalf(
				"failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		sheet, err := doc.AddSheet("Sheet1")
		if err != nil {
			t.Fatalf(
				"failed to add sheet: %v",
				err,
			)
		}

		// Group columns 2-5, then ungroup
		if err := sheet.GroupColumns(2, 5); err != nil {
			t.Fatalf(
				"GroupColumns failed: %v",
				err,
			)
		}

		if err := sheet.UngroupColumns(2, 5); err != nil {
			t.Fatalf(
				"UngroupColumns failed: %v",
				err,
			)
		}

		// Verify outline levels are removed
		ws := sheet.Worksheet()
		cols := ws.Cols()
		for i := uint32(2); i <= 5; i++ {
			col := cols.GetColByIndex(int(i))
			if col == nil {
				continue // Col might be removed when level = 0
			}
			if level := col.OutlineLevel(); level != 0 {
				t.Errorf(
					"expected column %d outline level 0, got %d",
					i,
					level,
				)
			}
		}
	})

	t.Run(
		"PartialUngrouping",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Group columns 2-5 twice (level 2)
			if err := sheet.GroupColumns(2, 5); err != nil {
				t.Fatalf(
					"GroupColumns failed: %v",
					err,
				)
			}
			if err := sheet.GroupColumns(2, 5); err != nil {
				t.Fatalf(
					"GroupColumns failed: %v",
					err,
				)
			}

			// Ungroup once (should go from level 2 to level 1)
			if err := sheet.UngroupColumns(2, 5); err != nil {
				t.Fatalf(
					"UngroupColumns failed: %v",
					err,
				)
			}

			// Verify outline levels are decremented to 1
			ws := sheet.Worksheet()
			cols := ws.Cols()
			for i := uint32(2); i <= 5; i++ {
				col := cols.GetColByIndex(int(i))
				if col == nil {
					t.Fatalf(
						"expected Col element for column %d to exist",
						i,
					)
				}
				if level := col.OutlineLevel(); level != 1 {
					t.Errorf(
						"expected column %d outline level 1, got %d",
						i,
						level,
					)
				}
			}
		},
	)

	t.Run(
		"UngroupAtLevelZero",
		func(t *testing.T) {
			doc, err := Create(
				t.TempDir()+"/test.xlsx",
				DocTypeWorkbook,
			)
			if err != nil {
				t.Fatalf(
					"failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			sheet, err := doc.AddSheet("Sheet1")
			if err != nil {
				t.Fatalf(
					"failed to add sheet: %v",
					err,
				)
			}

			// Ungroup columns that are not grouped (should be no-op)
			if err := sheet.UngroupColumns(2, 5); err != nil {
				t.Fatalf(
					"UngroupColumns failed: %v",
					err,
				)
			}

			// Verify outline levels remain 0
			ws := sheet.Worksheet()
			cols := ws.Cols()
			for i := uint32(2); i <= 5; i++ {
				col := cols.GetColByIndex(int(i))
				if col == nil {
					continue // No Col element expected for ungrouped columns
				}
				if level := col.OutlineLevel(); level != 0 {
					t.Errorf(
						"expected column %d outline level 0, got %d",
						i,
						level,
					)
				}
			}
		},
	)
}
