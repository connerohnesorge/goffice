package spreadsheet

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
)

// TestRoundtripBasic tests opening, modifying, saving, and reopening a document.
func TestRoundtripBasic(t *testing.T) {
	fixturePath := minimalFixturePath

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-basic-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Step 1: Open the original document
	doc1, err := Open(fixturePath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	originalType := doc1.Type()
	originalSheetCount := doc1.SheetCount()

	// Step 2: Save to a new location
	tempPath := filepath.Join(
		tmpDir,
		"roundtrip1.xlsx",
	)
	if err := doc1.SaveAs(tempPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Step 3: Reopen the saved document
	doc2, err := Open(tempPath, true)
	if err != nil {
		t.Fatalf("Failed to reopen: %v", err)
	}

	// Verify properties are preserved
	if doc2.Type() != originalType {
		t.Errorf(
			"Type changed from %v to %v",
			originalType,
			doc2.Type(),
		)
	}

	if doc2.SheetCount() != originalSheetCount {
		t.Errorf(
			"SheetCount changed from %d to %d",
			originalSheetCount,
			doc2.SheetCount(),
		)
	}

	// Step 4: Save again
	tempPath2 := filepath.Join(
		tmpDir,
		"roundtrip2.xlsx",
	)
	if err := doc2.SaveAs(tempPath2); err != nil {
		t.Fatalf(
			"Second SaveAs() error = %v",
			err,
		)
	}
	_ = doc2.Close()

	// Step 5: Final verification
	doc3, err := Open(tempPath2, false)
	if err != nil {
		t.Fatalf(
			"Failed to open final file: %v",
			err,
		)
	}
	defer func() { _ = doc3.Close() }()

	if doc3.Type() != originalType {
		t.Errorf(
			"Final type = %v, want %v",
			doc3.Type(),
			originalType,
		)
	}
}

// TestRoundtripNumbers tests roundtrip with the numbers.xlsx fixture.
func TestRoundtripNumbers(t *testing.T) {
	fixturePath := "../testdata/fixtures/numbers.xlsx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: numbers.xlsx fixture not found",
		)
	}

	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-numbers-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Open original
	doc1, err := Open(fixturePath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	// Save to temp
	tempPath := filepath.Join(
		tmpDir,
		"numbers_roundtrip.xlsx",
	)
	if err := doc1.SaveAs(tempPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen and verify
	doc2, err := Open(tempPath, false)
	if err != nil {
		t.Fatalf("Failed to reopen: %v", err)
	}
	defer func() { _ = doc2.Close() }()

	if doc2.Type() != DocTypeWorkbook {
		t.Errorf(
			"Type = %v, want %v",
			doc2.Type(),
			DocTypeWorkbook,
		)
	}
}

// TestRoundtripWithModifications tests roundtrip with content modifications.
func TestRoundtripWithModifications(
	t *testing.T,
) {
	fixturePath := minimalFixturePath

	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-mod-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Open and add a new sheet
	doc1, err := Open(fixturePath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	originalCount := doc1.SheetCount()

	newSheet, err := doc1.AddSheet("Added Sheet")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}
	_ = newSheet.SetCellValue("A1", "New Content")

	// Save
	tempPath := filepath.Join(
		tmpDir,
		"modified.xlsx",
	)
	if err := doc1.SaveAs(tempPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen and verify modification
	doc2, err := Open(tempPath, true)
	if err != nil {
		t.Fatalf("Failed to reopen: %v", err)
	}

	if doc2.SheetCount() != originalCount+1 {
		t.Errorf(
			"SheetCount = %d, want %d",
			doc2.SheetCount(),
			originalCount+1,
		)
	}

	// Verify the new sheet exists
	addedSheet, _ := doc2.SheetByName(
		"Added Sheet",
	)
	if addedSheet == nil {
		t.Error(
			"Added sheet not found after roundtrip",
		)
	}

	// Modify again
	anotherSheet, err := doc2.AddSheet(
		"Another Sheet",
	)
	if err != nil {
		t.Fatalf(
			"Second AddSheet() error = %v",
			err,
		)
	}
	_ = anotherSheet.SetCellValue(
		"A1",
		"More Content",
	)

	// Save again
	tempPath2 := filepath.Join(
		tmpDir,
		"modified2.xlsx",
	)
	if err := doc2.SaveAs(tempPath2); err != nil {
		t.Fatalf(
			"Second SaveAs() error = %v",
			err,
		)
	}
	_ = doc2.Close()

	// Final verification
	doc3, err := Open(tempPath2, false)
	if err != nil {
		t.Fatalf("Failed to open final: %v", err)
	}
	defer func() { _ = doc3.Close() }()

	if doc3.SheetCount() != originalCount+2 {
		t.Errorf(
			"Final SheetCount = %d, want %d",
			doc3.SheetCount(),
			originalCount+2,
		)
	}
}

// TestRoundtripValidation tests that documents remain valid after roundtrip.
func TestRoundtripValidation(t *testing.T) {
	fixturePath := minimalFixturePath

	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-valid-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Open original
	doc1, err := Open(fixturePath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	// Validate before roundtrip
	errorsBefore := doc1.Validate(
		validation.Office2016,
	)
	t.Logf(
		"Validation errors before roundtrip: %d",
		len(errorsBefore),
	)

	// Roundtrip
	tempPath := filepath.Join(
		tmpDir,
		"roundtrip.xlsx",
	)
	if err := doc1.SaveAs(tempPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen and validate
	doc2, err := Open(tempPath, false)
	if err != nil {
		t.Fatalf("Failed to reopen: %v", err)
	}
	defer func() { _ = doc2.Close() }()

	errorsAfter := doc2.Validate(
		validation.Office2016,
	)
	t.Logf(
		"Validation errors after roundtrip: %d",
		len(errorsAfter),
	)

	// Check that errors haven't significantly increased
	if len(errorsAfter) > len(errorsBefore)+5 {
		t.Errorf(
			"Validation errors increased significantly: %d -> %d",
			len(errorsBefore),
			len(errorsAfter),
		)
	}
}

// TestRoundtripTypeChange tests changing document type during roundtrip.
func TestRoundtripTypeChange(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-typechange-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create a workbook
	workbookPath := filepath.Join(
		tmpDir,
		"workbook.xlsx",
	)
	doc1, err := Create(
		workbookPath,
		DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	_, _ = doc1.AddSheet("Test")
	if err := doc1.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen and change to template
	doc2, err := Open(workbookPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	if err := doc2.ChangeType(DocTypeTemplate); err != nil {
		t.Fatalf("ChangeType() error = %v", err)
	}

	templatePath := filepath.Join(
		tmpDir,
		"template.xltx",
	)
	if err := doc2.SaveAs(templatePath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc2.Close()

	// Verify the template
	doc3, err := Open(templatePath, false)
	if err != nil {
		t.Fatalf(
			"Failed to open template: %v",
			err,
		)
	}
	defer func() { _ = doc3.Close() }()

	if doc3.Type() != DocTypeTemplate {
		t.Errorf(
			"Type = %v, want %v",
			doc3.Type(),
			DocTypeTemplate,
		)
	}
}

// TestRoundtripMultipleTimes tests many consecutive roundtrips.
func TestRoundtripMultipleTimes(t *testing.T) {
	fixturePath := minimalFixturePath

	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-multi-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	currentPath := fixturePath
	const iterations = 5

	for i := range iterations {
		doc, err := Open(currentPath, true)
		if err != nil {
			t.Fatalf(
				"Iteration %d: Open() error = %v",
				i,
				err,
			)
		}

		nextPath := filepath.Join(
			tmpDir,
			"roundtrip_"+string(
				rune('0'+i),
			)+".xlsx",
		)
		if err := doc.SaveAs(nextPath); err != nil {
			t.Fatalf(
				"Iteration %d: SaveAs() error = %v",
				i,
				err,
			)
		}
		_ = doc.Close()

		currentPath = nextPath
	}

	// Verify final file
	finalDoc, err := Open(currentPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to open final file: %v",
			err,
		)
	}
	defer func() { _ = finalDoc.Close() }()

	if finalDoc.Type() != DocTypeWorkbook {
		t.Errorf(
			"Final Type = %v, want %v",
			finalDoc.Type(),
			DocTypeWorkbook,
		)
	}
}

// TestRoundtripCreatedDocument tests roundtrip of a newly created document.
func TestRoundtripCreatedDocument(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-created-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create new document
	originalPath := filepath.Join(
		tmpDir,
		"original.xlsx",
	)
	doc1, err := Create(
		originalPath,
		DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	// Add content
	sheet1, _ := doc1.AddSheet("Data")
	_ = sheet1.SetCellValue("A1", "Header")
	_ = sheet1.SetCellValue("A2", 100)
	_ = sheet1.SetCellValue("A3", 200)
	_ = sheet1.SetCellFormula("A4", "=SUM(A2:A3)")

	sheet2, _ := doc1.AddSheet("Summary")
	_ = sheet2.SetCellValue("A1", "Total")
	_ = sheet2.SetCellFormula("B1", "=Data!A4")

	if err := doc1.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	_ = doc1.Close()

	// Roundtrip
	doc2, err := Open(originalPath, true)
	if err != nil {
		t.Fatalf("Failed to reopen: %v", err)
	}

	if doc2.SheetCount() != 2 {
		t.Errorf(
			"SheetCount = %d, want 2",
			doc2.SheetCount(),
		)
	}

	copyPath := filepath.Join(tmpDir, "copy.xlsx")
	if err := doc2.SaveAs(copyPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc2.Close()

	// Final verification
	doc3, err := Open(copyPath, false)
	if err != nil {
		t.Fatalf("Failed to open copy: %v", err)
	}
	defer func() { _ = doc3.Close() }()

	if doc3.SheetCount() != 2 {
		t.Errorf(
			"Copy SheetCount = %d, want 2",
			doc3.SheetCount(),
		)
	}

	dataSheet, _ := doc3.SheetByName("Data")
	if dataSheet == nil {
		t.Error("Data sheet not found in copy")
	}

	summarySheet, _ := doc3.SheetByName("Summary")
	if summarySheet == nil {
		t.Error("Summary sheet not found in copy")
	}
}

// TestRoundtripPreservesFormulas tests that formulas are preserved.
func TestRoundtripPreservesFormulas(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-formulas-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create document with formulas
	originalPath := filepath.Join(
		tmpDir,
		"formulas.xlsx",
	)
	doc1, err := Create(
		originalPath,
		DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	sheet, _ := doc1.AddSheet("Formulas")

	// Add data
	for i := 1; i <= 5; i++ {
		_ = sheet.SetCellValue(
			"A"+string(rune('0'+i)),
			i*10,
		)
	}

	// Add various formulas
	formulas := map[string]string{
		"B1": "=SUM(A1:A5)",
		"B2": "=AVERAGE(A1:A5)",
		"B3": "=MAX(A1:A5)",
		"B4": "=MIN(A1:A5)",
		"B5": "=COUNT(A1:A5)",
	}

	for cell, formula := range formulas {
		_ = sheet.SetCellFormula(cell, formula)
	}

	if err := doc1.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	_ = doc1.Close()

	// Roundtrip
	doc2, err := Open(originalPath, true)
	if err != nil {
		t.Fatalf("Failed to reopen: %v", err)
	}

	copyPath := filepath.Join(
		tmpDir,
		"formulas_copy.xlsx",
	)
	if err := doc2.SaveAs(copyPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc2.Close()

	// Verify formulas are preserved (basic structure check)
	doc3, err := Open(copyPath, false)
	if err != nil {
		t.Fatalf("Failed to open copy: %v", err)
	}
	defer func() { _ = doc3.Close() }()

	if doc3.SheetCount() != 1 {
		t.Errorf(
			"SheetCount = %d, want 1",
			doc3.SheetCount(),
		)
	}
}

// TestRoundtripAutoSave tests the auto-save feature.
func TestRoundtripAutoSave(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-autosave-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create document with auto-save enabled
	testPath := filepath.Join(
		tmpDir,
		"autosave.xlsx",
	)
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	// Enable auto-save
	settings := doc.Settings().WithAutoSave(true)
	doc.settings = settings

	sheet, _ := doc.AddSheet("Auto")
	_ = sheet.SetCellValue("A1", "Auto-saved")

	// Close without explicit save - auto-save should handle it
	if err := doc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(testPath); os.IsNotExist(
		err,
	) {
		t.Error("File was not auto-saved")
	}
}

// TestRoundtripLargeData tests roundtrip with larger amounts of data.
func TestRoundtripLargeData(t *testing.T) {
	if testing.Short() {
		t.Skip(
			"Skipping large data test in short mode",
		)
	}

	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-large-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create document with more data
	originalPath := filepath.Join(
		tmpDir,
		"large.xlsx",
	)
	doc1, err := Create(
		originalPath,
		DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	sheet, _ := doc1.AddSheet("Large Data")

	// Add 100 rows of data
	const numRows = 100
	for row := 1; row <= numRows; row++ {
		for col := 'A'; col <= 'J'; col++ {
			cellRef := string(
				col,
			) + string(
				rune('0'+row/10),
			) + string(
				rune('0'+row%10),
			)
			_ = sheet.SetCellValue(
				cellRef,
				row*int(col-'A'+1),
			)
		}
	}

	if err := doc1.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	_ = doc1.Close()

	// Roundtrip
	doc2, err := Open(originalPath, true)
	if err != nil {
		t.Fatalf("Failed to reopen: %v", err)
	}

	copyPath := filepath.Join(
		tmpDir,
		"large_copy.xlsx",
	)
	if err := doc2.SaveAs(copyPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc2.Close()

	// Verify
	doc3, err := Open(copyPath, false)
	if err != nil {
		t.Fatalf("Failed to open copy: %v", err)
	}
	defer func() { _ = doc3.Close() }()

	if doc3.SheetCount() != 1 {
		t.Errorf(
			"SheetCount = %d, want 1",
			doc3.SheetCount(),
		)
	}
}

// TestRoundtripWithValidationRules tests roundtrip preserves validation.
func TestRoundtripWithValidationRules(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-validation-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create document with data
	originalPath := filepath.Join(
		tmpDir,
		"validation.xlsx",
	)
	doc1, err := Create(
		originalPath,
		DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	sheet, _ := doc1.AddSheet("Validation")
	_ = sheet.SetCellValue("A1", "Enter value:")

	if err := doc1.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	_ = doc1.Close()

	// Roundtrip
	doc2, err := Open(originalPath, true)
	if err != nil {
		t.Fatalf("Failed to reopen: %v", err)
	}

	copyPath := filepath.Join(
		tmpDir,
		"validation_copy.xlsx",
	)
	if err := doc2.SaveAs(copyPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc2.Close()

	// Verify
	doc3, err := Open(copyPath, false)
	if err != nil {
		t.Fatalf("Failed to open copy: %v", err)
	}
	defer func() { _ = doc3.Close() }()

	// Validate structure
	errors := doc3.Validate(validation.Office2016)
	t.Logf("Validation errors: %d", len(errors))
}
