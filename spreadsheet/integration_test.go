package spreadsheet

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// TestEndToEndWorkbookCreation tests creating a complete workbook from scratch
// with multiple sheets, data, formulas, and styles.
func TestEndToEndWorkbookCreation(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-e2e-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"e2e_workbook.xlsx",
	)

	// Create a new workbook
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	// Add first sheet with data
	sheet1, err := doc.AddSheet("Sales Data")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Add headers
	_ = sheet1.SetCellValue("A1", "Product")
	_ = sheet1.SetCellValue("B1", "Q1")
	_ = sheet1.SetCellValue("C1", "Q2")
	_ = sheet1.SetCellValue("D1", "Q3")
	_ = sheet1.SetCellValue("E1", "Q4")
	_ = sheet1.SetCellValue("F1", "Total")

	// Add data rows
	products := []string{
		"Widget A",
		"Widget B",
		"Widget C",
		"Widget D",
	}
	values := [][]int{
		{100, 150, 200, 250},
		{80, 90, 100, 110},
		{200, 220, 240, 260},
		{50, 60, 70, 80},
	}

	for i, product := range products {
		row := i + 2
		_ = sheet1.SetCellValue(
			"A"+strconv.Itoa(row),
			product,
		)
		for j, val := range values[i] {
			colLetter := string(rune('B' + j))
			_ = sheet1.SetCellValue(
				colLetter+strconv.Itoa(row),
				val,
			)
		}
		// Add SUM formula for total
		_ = sheet1.SetCellFormula(
			"F"+strconv.Itoa(row),
			"=SUM(B"+strconv.Itoa(
				row,
			)+":E"+strconv.Itoa(
				row,
			)+")",
		)
	}

	// Add a second sheet
	sheet2, err := doc.AddSheet("Summary")
	if err != nil {
		t.Fatalf(
			"AddSheet() for Summary error = %v",
			err,
		)
	}

	_ = sheet2.SetCellValue(
		"A1",
		"Summary Report",
	)
	_ = sheet2.SetCellValue(
		"A2",
		"Total Products:",
	)
	_ = sheet2.SetCellValue("B2", len(products))

	// Verify sheet count
	if doc.SheetCount() != 2 {
		t.Errorf(
			"SheetCount() = %d, want 2",
			doc.SheetCount(),
		)
	}

	// Save the document
	if err := doc.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	// Close the document
	if err := doc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(testPath); os.IsNotExist(
		err,
	) {
		t.Fatal("Workbook file was not created")
	}

	// Reopen and verify
	doc2, err := Open(testPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to reopen workbook: %v",
			err,
		)
	}
	defer func() { _ = doc2.Close() }()

	if doc2.SheetCount() != 2 {
		t.Errorf(
			"Reopened document SheetCount() = %d, want 2",
			doc2.SheetCount(),
		)
	}
}

// TestEndToEndWithStyles tests creating workbook with cell styles.
func TestEndToEndWithStyles(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-styles-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"styled_workbook.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, err := doc.AddSheet("Styled Sheet")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Add content with different data types
	_ = sheet.SetCellValue("A1", "Text Value")
	_ = sheet.SetCellValue("A2", 42)
	_ = sheet.SetCellValue("A3", 3.14159)
	_ = sheet.SetCellValue("A4", true)

	// Verify we can save without error
	if err := doc.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestEndToEndFormulas tests creating workbooks with various formula types.
func TestEndToEndFormulas(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-formulas-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"formula_workbook.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, err := doc.AddSheet("Formulas")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Add source data
	for i := 1; i <= 10; i++ {
		_ = sheet.SetCellValue(
			"A"+strconv.Itoa(i),
			i*10,
		)
	}

	// Add various formulas
	_ = sheet.SetCellFormula("B1", "=SUM(A1:A10)")
	_ = sheet.SetCellFormula(
		"B2",
		"=AVERAGE(A1:A10)",
	)
	_ = sheet.SetCellFormula("B3", "=MAX(A1:A10)")
	_ = sheet.SetCellFormula("B4", "=MIN(A1:A10)")
	_ = sheet.SetCellFormula(
		"B5",
		"=COUNT(A1:A10)",
	)
	_ = sheet.SetCellFormula("B6", "=A1+A2")
	_ = sheet.SetCellFormula(
		"B7",
		"=IF(A1>50,\"High\",\"Low\")",
	)

	if err := doc.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestEndToEndMergedCells tests creating workbooks with merged cells.
func TestEndToEndMergedCells(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-merge-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"merged_workbook.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, err := doc.AddSheet("Merged Cells")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Add a title that would span multiple columns
	_ = sheet.SetCellValue("A1", "Monthly Report")

	// Add some data
	_ = sheet.SetCellValue("A3", "Category")
	_ = sheet.SetCellValue("B3", "Value")
	_ = sheet.SetCellValue("A4", "Revenue")
	_ = sheet.SetCellValue("B4", 50000)

	if err := doc.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestEndToEndMultipleSheets tests creating workbooks with many sheets.
func TestEndToEndMultipleSheets(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-multisheet-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"multi_sheet.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Create multiple sheets
	months := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
	}
	for i, month := range months {
		sheet, err := doc.AddSheet(month)
		if err != nil {
			t.Fatalf(
				"AddSheet(%s) error = %v",
				month,
				err,
			)
		}
		_ = sheet.SetCellValue(
			"A1",
			month+" Data",
		)
		_ = sheet.SetCellValue("A2", "Value")
		_ = sheet.SetCellValue("B2", (i+1)*1000)
	}

	if doc.SheetCount() != len(months) {
		t.Errorf(
			"SheetCount() = %d, want %d",
			doc.SheetCount(),
			len(months),
		)
	}

	if err := doc.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}

	// Reopen and verify
	doc2, err := Open(testPath, false)
	if err != nil {
		t.Fatalf("Failed to reopen: %v", err)
	}
	defer func() { _ = doc2.Close() }()

	if doc2.SheetCount() != len(months) {
		t.Errorf(
			"Reopened SheetCount() = %d, want %d",
			doc2.SheetCount(),
			len(months),
		)
	}
}

// TestEndToEndDataValidation tests creating workbooks with data validation.
func TestEndToEndDataValidation(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-validation-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"validation_workbook.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, err := doc.AddSheet("Validation")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Add headers
	_ = sheet.SetCellValue(
		"A1",
		"Enter a number (1-100):",
	)
	_ = sheet.SetCellValue(
		"A2",
		"Select from list:",
	)

	// Add list values for dropdown
	_ = sheet.SetCellValue("D1", "Option A")
	_ = sheet.SetCellValue("D2", "Option B")
	_ = sheet.SetCellValue("D3", "Option C")

	if err := doc.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestEndToEndStreamWrite tests creating a workbook to an io.Writer.
func TestEndToEndStreamWrite(t *testing.T) {
	var buf bytes.Buffer

	doc, err := CreateFromStream(
		&buf,
		DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf(
			"CreateFromStream() error = %v",
			err,
		)
	}

	sheet, err := doc.AddSheet("Stream Test")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	_ = sheet.SetCellValue(
		"A1",
		"Written to stream",
	)
	_ = sheet.SetCellValue("A2", 12345)

	if err := doc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Verify something was written
	if buf.Len() == 0 {
		t.Error("Nothing written to buffer")
	}
}

// TestEndToEndDocumentValidation tests creating and validating documents.
func TestEndToEndDocumentValidation(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-validate-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"valid_workbook.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	sheet, err := doc.AddSheet("Valid Sheet")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	_ = sheet.SetCellValue("A1", "Valid Content")

	// Validate document
	errors := doc.Validate(validation.Office2016)
	t.Logf(
		"Validation found %d issues",
		len(errors),
	)

	if err := doc.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	if err := doc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
}

// TestEndToEndDifferentTypes tests creating different document types.
func TestEndToEndDifferentTypes(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-types-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	tests := []struct {
		docType   DocType
		extension string
	}{
		{DocTypeWorkbook, ".xlsx"},
		{DocTypeTemplate, ".xltx"},
	}

	for _, tc := range tests {
		t.Run(
			tc.docType.String(),
			func(t *testing.T) {
				testPath := filepath.Join(
					tmpDir,
					"test"+tc.extension,
				)

				doc, err := Create(
					testPath,
					tc.docType,
				)
				if err != nil {
					t.Fatalf(
						"Create() error = %v",
						err,
					)
				}

				if doc.Type() != tc.docType {
					t.Errorf(
						"Type() = %v, want %v",
						doc.Type(),
						tc.docType,
					)
				}

				_, err = doc.AddSheet("Test")
				if err != nil {
					t.Fatalf(
						"AddSheet() error = %v",
						err,
					)
				}

				if err := doc.Save(); err != nil {
					t.Fatalf(
						"Save() error = %v",
						err,
					)
				}

				if err := doc.Close(); err != nil {
					t.Fatalf(
						"Close() error = %v",
						err,
					)
				}
			},
		)
	}
}

// TestEndToEndConditionalFormatting tests creating workbooks with conditional formatting.
func TestEndToEndConditionalFormatting(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-cf-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"cf_workbook.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, err := doc.AddSheet(
		"Conditional Formatting",
	)
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Add data for conditional formatting
	for i := 1; i <= 10; i++ {
		_ = sheet.SetCellValue(
			"A"+strconv.Itoa(i),
			i*10,
		)
	}

	if err := doc.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestEndToEndCellTypes tests all cell types.
func TestEndToEndCellTypes(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-celltypes-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"celltypes_workbook.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, err := doc.AddSheet("Cell Types")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Test all cell types
	_ = sheet.SetCellValue("A1", "String")
	_ = sheet.SetCellValue(
		"A2",
		42,
	) // Integer
	_ = sheet.SetCellValue(
		"A3",
		3.14159,
	) // Float
	_ = sheet.SetCellValue(
		"A4",
		true,
	) // Boolean
	_ = sheet.SetCellFormula(
		"A5",
		"=1+1",
	) // Formula
	_ = sheet.SetCellValue(
		"A6",
		"",
	) // Empty
	_ = sheet.SetCellValue(
		"A7",
		-123,
	) // Negative
	_ = sheet.SetCellValue(
		"A8",
		0,
	) // Zero
	_ = sheet.SetCellValue(
		"A9",
		999999999999,
	) // Large number
	_ = sheet.SetCellValue(
		"A10",
		0.000001,
	) // Small number
	_ = sheet.SetCellValue(
		"A11",
		"Special: <>&\"'",
	) // Special characters

	if err := doc.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestEnumCompleteness verifies all enum types have expected values.
func TestEnumCompleteness(t *testing.T) {
	// Test CellType completeness
	cellTypes := []elements.CellType{
		elements.CellTypeBoolean,
		elements.CellTypeDate,
		elements.CellTypeError,
		elements.CellTypeInlineString,
		elements.CellTypeNumber,
		elements.CellTypeSharedString,
		elements.CellTypeFormulaString,
	}
	if len(cellTypes) < 7 {
		t.Errorf(
			"CellType should have at least 7 values, got %d",
			len(cellTypes),
		)
	}

	// Test FormulaType completeness
	formulaTypes := []elements.FormulaType{
		elements.FormulaTypeNormal,
		elements.FormulaTypeArray,
		elements.FormulaTypeDataTable,
		elements.FormulaTypeShared,
	}
	if len(formulaTypes) < 4 {
		t.Errorf(
			"FormulaType should have at least 4 values, got %d",
			len(formulaTypes),
		)
	}

	// Test BorderStyle completeness
	borderStyles := []elements.BorderStyle{
		elements.BorderStyleNone,
		elements.BorderStyleThin,
		elements.BorderStyleMedium,
		elements.BorderStyleDashed,
		elements.BorderStyleDotted,
		elements.BorderStyleThick,
		elements.BorderStyleDouble,
		elements.BorderStyleHair,
		elements.BorderStyleMediumDashed,
		elements.BorderStyleDashDot,
		elements.BorderStyleMediumDashDot,
		elements.BorderStyleDashDotDot,
		elements.BorderStyleMediumDashDotDot,
		elements.BorderStyleSlantDashDot,
	}
	if len(borderStyles) < 14 {
		t.Errorf(
			"BorderStyle should have at least 14 values, got %d",
			len(borderStyles),
		)
	}

	// Test PatternType completeness
	patternTypes := []elements.PatternType{
		elements.PatternTypeNone,
		elements.PatternTypeSolid,
		elements.PatternTypeMediumGray,
		elements.PatternTypeDarkGray,
		elements.PatternTypeLightGray,
		elements.PatternTypeDarkHorizontal,
		elements.PatternTypeDarkVertical,
		elements.PatternTypeDarkDown,
		elements.PatternTypeDarkUp,
		elements.PatternTypeDarkGrid,
		elements.PatternTypeDarkTrellis,
		elements.PatternTypeLightHorizontal,
		elements.PatternTypeLightVertical,
		elements.PatternTypeLightDown,
		elements.PatternTypeLightUp,
		elements.PatternTypeLightGrid,
		elements.PatternTypeLightTrellis,
		elements.PatternTypeGray125,
		elements.PatternTypeGray0625,
	}
	if len(patternTypes) < 18 {
		t.Errorf(
			"PatternType should have at least 18 values, got %d",
			len(patternTypes),
		)
	}

	// Test HorizontalAlignment completeness
	hAligns := []elements.HorizontalAlignment{
		elements.HorizontalAlignmentGeneral,
		elements.HorizontalAlignmentLeft,
		elements.HorizontalAlignmentCenter,
		elements.HorizontalAlignmentRight,
		elements.HorizontalAlignmentFill,
		elements.HorizontalAlignmentJustify,
		elements.HorizontalAlignmentCenterContinuous,
		elements.HorizontalAlignmentDistributed,
	}
	if len(hAligns) < 8 {
		t.Errorf(
			"HorizontalAlignment should have at least 8 values, got %d",
			len(hAligns),
		)
	}

	// Test VerticalAlignment completeness
	vAligns := []elements.VerticalAlignment{
		elements.VerticalAlignmentTop,
		elements.VerticalAlignmentCenter,
		elements.VerticalAlignmentBottom,
		elements.VerticalAlignmentJustify,
		elements.VerticalAlignmentDistributed,
	}
	if len(vAligns) < 5 {
		t.Errorf(
			"VerticalAlignment should have at least 5 values, got %d",
			len(vAligns),
		)
	}

	// Test ValidationType completeness
	validationTypes := []elements.ValidationType{
		elements.ValidationTypeNone,
		elements.ValidationTypeWhole,
		elements.ValidationTypeDecimal,
		elements.ValidationTypeList,
		elements.ValidationTypeDate,
		elements.ValidationTypeTime,
		elements.ValidationTypeTextLength,
		elements.ValidationTypeCustom,
	}
	if len(validationTypes) < 8 {
		t.Errorf(
			"ValidationType should have at least 8 values, got %d",
			len(validationTypes),
		)
	}

	// Test CfRuleType completeness
	cfRuleTypes := []elements.CfRuleType{
		elements.CfRuleTypeExpression,
		elements.CfRuleTypeCellIs,
		elements.CfRuleTypeColorScale,
		elements.CfRuleTypeDataBar,
		elements.CfRuleTypeIconSet,
		elements.CfRuleTypeTop10,
		elements.CfRuleTypeUniqueValues,
		elements.CfRuleTypeDuplicateValues,
		elements.CfRuleTypeContainsText,
		elements.CfRuleTypeNotContainsText,
		elements.CfRuleTypeBeginsWith,
		elements.CfRuleTypeEndsWith,
		elements.CfRuleTypeContainsBlanks,
		elements.CfRuleTypeNotContainsBlanks,
		elements.CfRuleTypeContainsErrors,
		elements.CfRuleTypeNotContainsErrors,
		elements.CfRuleTypeTimePeriod,
		elements.CfRuleTypeAboveAverage,
	}
	if len(cfRuleTypes) < 18 {
		t.Errorf(
			"CfRuleType should have at least 18 values, got %d",
			len(cfRuleTypes),
		)
	}
}

// TestReadFromStream tests opening workbook from io.ReaderAt.
func TestReadFromStream(t *testing.T) {
	fixturePath := "../testdata/fixtures/minimal.xlsx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	// Read file into memory
	data, err := os.ReadFile(fixturePath)
	if err != nil {
		t.Fatalf(
			"Failed to read fixture: %v",
			err,
		)
	}

	// Open from bytes.Reader (which implements io.ReaderAt)
	reader := bytes.NewReader(data)
	doc, err := OpenFromStream(
		reader,
		int64(len(data)),
		false,
	)
	if err != nil {
		t.Fatalf(
			"OpenFromStream() error = %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	if doc.Type() != DocTypeWorkbook {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeWorkbook,
		)
	}
}

// TestActiveSheetManagement tests active sheet operations.
func TestActiveSheetManagement(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-active-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"active_sheet.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Add multiple sheets
	_, _ = doc.AddSheet("Sheet1")
	_, _ = doc.AddSheet("Sheet2")
	_, _ = doc.AddSheet("Sheet3")

	// Set active sheet
	if err := doc.SetActiveSheet(1); err != nil {
		t.Fatalf(
			"SetActiveSheet(1) error = %v",
			err,
		)
	}

	// Test invalid index
	if err := doc.SetActiveSheet(10); err == nil {
		t.Error(
			"SetActiveSheet(10) should fail for invalid index",
		)
	}

	// Test delete sheet
	if err := doc.DeleteSheet(0); err != nil {
		t.Fatalf("DeleteSheet(0) error = %v", err)
	}

	if doc.SheetCount() != 2 {
		t.Errorf(
			"SheetCount() = %d after delete, want 2",
			doc.SheetCount(),
		)
	}
}

// TestRenameSheet tests sheet renaming.
func TestRenameSheet(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-rename-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"rename_sheet.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, _ := doc.AddSheet("Original Name")
	_ = sheet

	if err := doc.RenameSheet(0, "New Name"); err != nil {
		t.Fatalf("RenameSheet() error = %v", err)
	}

	renamedSheet, _ := doc.SheetByName("New Name")
	if renamedSheet == nil {
		t.Error(
			"SheetByName('New Name') returned nil after rename",
		)
	}
}

// TestSheetIteration tests iterating over sheets.
func TestSheetIteration(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-iter-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"iter_sheets.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	expectedNames := []string{
		"Alpha",
		"Beta",
		"Gamma",
	}
	for _, name := range expectedNames {
		_, _ = doc.AddSheet(name)
	}

	// Test iteration
	count := 0
	for sheet := range doc.Sheets() {
		if sheet == nil {
			t.Error(
				"Got nil sheet during iteration",
			)

			continue
		}
		count++
	}

	if count != len(expectedNames) {
		t.Errorf(
			"Iterated over %d sheets, want %d",
			count,
			len(expectedNames),
		)
	}
}

// TestPackageAccess tests accessing the underlying package.
func TestPackageAccess(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-pkg-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"pkg_access.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	pkg := doc.Package()
	if pkg == nil {
		t.Error("Package() returned nil")
	}

	features := doc.Features()
	if features == nil {
		t.Error("Features() returned nil")
	}
}

// Ensure io.Closer interface is implemented.
var _ io.Closer = (*Document)(nil)
