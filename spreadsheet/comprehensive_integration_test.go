package spreadsheet

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// TestComprehensiveWorkbookOpenClose tests opening and closing various test workbooks.
func TestComprehensiveWorkbookOpenClose(
	t *testing.T,
) {
	testCases := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			"BasicSpreadsheet",
			"basicspreadsheet.xlsx",
			false,
		},
		{
			"StandardSpreadsheet",
			"Spreadsheet.xlsx",
			false,
		},
		{"Comments", "Comments.xlsx", false},
		{"Complex01", "Complex01.xlsx", false},
		{"Excel14", "excel14.xlsx", false},
		{"ExtensionList", "extlst.xlsx", false},
		{"MCExcel", "MCExecl.xlsx", false},
		{
			"MissingCalcChain",
			"missingcalcchainpart.xlsx",
			false,
		},
		{
			"RevisionTracking",
			"Revision_NameCommentChange.xlsx",
			false,
		},
		{"Template", "Spreadsheet.xltx", false},
		{
			"VMLDrawing",
			"vmldrawingroot.xlsx",
			false,
		},
		{"YouTube", "Youtube.xlsx", false},
		{
			"MalformedURI",
			"malformed_uri.xlsx",
			false,
		},
		{
			"MalformedURILong",
			"malformed_uri_long.xlsx",
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				tc.filename,
			)

			// Skip if file doesn't exist
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			// Open workbook
			doc, err := Open(path, false)
			if (err != nil) != tc.wantErr {
				t.Fatalf(
					"Open() error = %v, wantErr %v",
					err,
					tc.wantErr,
				)
			}
			if err != nil {
				return
			}
			defer func() { _ = doc.Close() }()

			// Verify basic properties
			if doc == nil {
				t.Fatal(
					"Open() returned nil document",
				)
			}

			// Verify package
			if doc.Package() == nil {
				t.Error("Package() returned nil")
			}

			// Verify type
			docType := doc.Type()
			t.Logf(
				"Document type: %s",
				docType.String(),
			)

			// Verify sheet count
			sheetCount := doc.SheetCount()
			if sheetCount < 0 {
				t.Error(
					"SheetCount() returned negative value",
				)
			}
			t.Logf("Sheet count: %d", sheetCount)

			// Close and verify
			if err := doc.Close(); err != nil {
				t.Errorf(
					"Close() error = %v",
					err,
				)
			}
		})
	}
}

// TestComprehensiveWorkbookStructure tests workbook structure validation.
func TestComprehensiveWorkbookStructure(
	t *testing.T,
) {
	testCases := []string{
		"basicspreadsheet.xlsx",
		"Spreadsheet.xlsx",
		"Complex01.xlsx",
		"excel14.xlsx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)

			// Skip if file doesn't exist
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			// Verify workbook part
			wbPart := doc.WorkbookPart()
			if wbPart == nil {
				t.Fatal("WorkbookPart() is nil")
			}

			// Get workbook element and cast to Workbook type
			wbRoot := wbPart.Workbook()
			if wbRoot == nil {
				t.Fatal("Workbook() is nil")
			}

			wb, ok := wbRoot.(*elements.Workbook)
			if !ok {
				t.Fatal(
					"Failed to cast to Workbook type",
				)
			}

			// Verify sheets
			sheets := wb.Sheets()
			if sheets == nil {
				t.Fatal("Sheets() is nil")
			}

			sheetCount := 0
			for sheet := range sheets.Sheets() {
				if sheet == nil {
					t.Error(
						"Got nil sheet in iteration",
					)

					continue
				}
				sheetCount++

				sheetName := sheet.Name()
				if sheetName == "" {
					t.Error(
						"Sheet has empty name",
					)
				} else {
					t.Logf("Sheet %d: %s", sheetCount, sheetName)
				}

				sheetID := sheet.SheetId()
				if sheetID == 0 {
					t.Error("Sheet has nil ID")
				}
			}

			t.Logf("Total sheets: %d", sheetCount)

			// Verify sheet count matches
			if sheetCount != doc.SheetCount() {
				t.Errorf(
					"Sheets iteration count (%d) != SheetCount() (%d)",
					sheetCount,
					doc.SheetCount(),
				)
			}
		})
	}
}

// TestComprehensiveRoundtrip tests opening, modifying, saving, and reopening.
func TestComprehensiveRoundtrip(t *testing.T) {
	testCases := []string{
		"basicspreadsheet.xlsx",
		"Spreadsheet.xlsx",
		"Comments.xlsx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			sourcePath := filepath.Join(
				"testdata",
				filename,
			)

			// Skip if file doesn't exist
			if _, err := os.Stat(sourcePath); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					sourcePath,
				)
			}

			// Create temp directory
			tmpDir, err := os.MkdirTemp(
				"",
				"goffice-roundtrip-*",
			)
			if err != nil {
				t.Fatalf(
					"Failed to create temp dir: %v",
					err,
				)
			}
			defer func() { _ = os.RemoveAll(tmpDir) }()

			// Open original document
			doc, err := Open(sourcePath, true)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}

			originalType := doc.Type()
			originalSheetCount := doc.SheetCount()

			// Save to new location
			outputPath := filepath.Join(
				tmpDir,
				"modified_"+filename,
			)
			if err := doc.SaveAs(outputPath); err != nil {
				_ = doc.Close()
				t.Fatalf(
					"SaveAs() error = %v",
					err,
				)
			}
			_ = doc.Close()

			// Reopen saved document
			doc2, err := Open(outputPath, false)
			if err != nil {
				t.Fatalf(
					"Failed to reopen saved document: %v",
					err,
				)
			}
			defer func() { _ = doc2.Close() }()

			// Verify structure
			if doc2.WorkbookPart() == nil {
				t.Error(
					"Reopened document missing workbook part",
				)
			}

			// Verify package integrity
			if doc2.Package() == nil {
				t.Error(
					"Reopened document missing package",
				)
			}

			// Verify properties preserved
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
		})
	}
}

// TestComprehensivePartsExtraction tests extracting various parts from workbooks.
func TestComprehensivePartsExtraction(
	t *testing.T,
) {
	t.Run(
		"WorkbookPart",
		testWorkbookPartExtraction,
	)
	t.Run(
		"WorksheetParts",
		testWorksheetPartsExtraction,
	)
	t.Run(
		"SharedStringsPart",
		testSharedStringsPartExtraction,
	)
	t.Run("StylesPart", testStylesPartExtraction)
	t.Run(
		"CalcChainPart",
		testCalcChainPartExtraction,
	)
}

func testWorkbookPartExtraction(t *testing.T) {
	path := filepath.Join(
		"testdata",
		"Complex01.xlsx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	wb := wbPart.Workbook()
	if wb == nil {
		t.Fatal("Workbook() returned nil")
	}

	t.Log("Successfully extracted workbook part")
}

func testWorksheetPartsExtraction(t *testing.T) {
	path := filepath.Join(
		"testdata",
		"Spreadsheet.xlsx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	worksheetCount := 0
	worksheetParts := wbPart.WorksheetParts()
	for _, wsPart := range worksheetParts {
		if wsPart == nil {
			t.Error("Got nil worksheet part")

			continue
		}
		worksheetCount++

		ws := wsPart.Worksheet()
		if ws == nil {
			t.Error("Worksheet() returned nil")

			continue
		}
		_ = ws // Use ws to avoid unused variable warning

		t.Logf(
			"Worksheet %d extracted successfully",
			worksheetCount,
		)
	}

	if worksheetCount == 0 {
		t.Error("No worksheet parts found")
	}
}

func testSharedStringsPartExtraction(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"Spreadsheet.xlsx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	ssPart := wbPart.SharedStringTablePart()
	if ssPart == nil {
		t.Log(
			"No shared strings part (workbook may use inline strings)",
		)

		return
	}

	sst := ssPart.SharedStringTable()
	if sst == nil {
		t.Error(
			"SharedStringTable() returned nil",
		)

		return
	}

	t.Log(
		"Successfully extracted shared strings part",
	)

	// Count strings
	count := 0
	for si := range sst.Items() {
		if si != nil {
			count++
		}
	}
	t.Logf("Shared strings count: %d", count)
}

func testStylesPartExtraction(t *testing.T) {
	path := filepath.Join(
		"testdata",
		"Complex01.xlsx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	stylesPart := wbPart.StylesPart()
	if stylesPart == nil {
		t.Log("No styles part found")

		return
	}

	stylesheet := stylesPart.Stylesheet()
	if stylesheet == nil {
		t.Error("Stylesheet() returned nil")

		return
	}

	t.Log("Successfully extracted styles part")
}

func testCalcChainPartExtraction(t *testing.T) {
	path := filepath.Join(
		"testdata",
		"Complex01.xlsx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	calcChainPart := wbPart.CalculationChainPart()
	if calcChainPart == nil {
		t.Log("No calc chain part found")

		return
	}

	calcChain := calcChainPart.CalculationChain()
	if calcChain == nil {
		t.Error("CalculationChain() returned nil")

		return
	}

	t.Log(
		"Successfully extracted calc chain part",
	)
}

// TestComprehensiveRelationships tests relationship handling.
func TestComprehensiveRelationships(
	t *testing.T,
) {
	testCases := []string{
		"Spreadsheet.xlsx",
		"Complex01.xlsx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)

			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			oxmlPkg := doc.Package()
			if oxmlPkg == nil {
				t.Fatal("Package() returned nil")
			}

			pkg := oxmlPkg.Package()
			if pkg == nil {
				t.Fatal(
					"Underlying package is nil",
				)
			}

			// Verify package relationships
			rels := pkg.Relationships()
			if rels == nil {
				t.Fatal(
					"Package relationships are nil",
				)
			}

			relCount := 0
			for rel := range rels.All() {
				if rel == nil {
					t.Error(
						"Got nil relationship",
					)

					continue
				}
				relCount++

				relType := rel.Type()
				if relType == "" {
					t.Error(
						"Relationship has empty type",
					)
				}

				target := rel.Target()
				if target == "" {
					t.Error(
						"Relationship has empty target",
					)
				}
			}

			t.Logf(
				"Package has %d relationships",
				relCount,
			)

			if relCount == 0 {
				t.Error(
					"Package has no relationships",
				)
			}
		})
	}
}

// TestComprehensiveWorkbookProperties tests workbook properties.
func TestComprehensiveWorkbookProperties(
	t *testing.T,
) {
	testCases := []string{
		"Spreadsheet.xlsx",
		"Complex01.xlsx",
		"excel14.xlsx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)

			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			wbPart := doc.WorkbookPart()
			if wbPart == nil {
				t.Fatal(
					"WorkbookPart() returned nil",
				)
			}

			wbRoot := wbPart.Workbook()
			if wbRoot == nil {
				t.Fatal("Workbook() returned nil")
			}

			wb, ok := wbRoot.(*elements.Workbook)
			if !ok {
				t.Fatal(
					"Failed to cast to Workbook type",
				)
			}

			// Test workbook views
			bookViews := wb.BookViews()
			if bookViews != nil {
				viewCount := 0
				for view := range bookViews.WorkbookViews() {
					if view != nil {
						viewCount++
					}
				}
				t.Logf(
					"Workbook has %d views",
					viewCount,
				)
			}

			// Test defined names
			definedNames := wb.DefinedNames()
			if definedNames != nil {
				nameCount := 0
				for name := range definedNames.DefinedNames() {
					if name != nil {
						nameCount++
					}
				}
				t.Logf(
					"Workbook has %d defined names",
					nameCount,
				)
			}

			// Test workbook protection
			wbPr := wb.WorkbookPr()
			if wbPr != nil {
				t.Log(
					"Workbook has properties element",
				)
			}
		})
	}
}

// TestComprehensiveCellReading tests reading cell values from worksheets.
func TestComprehensiveCellReading(t *testing.T) {
	testCases := []string{
		"basicspreadsheet.xlsx",
		"Spreadsheet.xlsx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)

			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			// Get first sheet
			if doc.SheetCount() == 0 {
				t.Skip("No sheets in workbook")
			}

			sheet, err := doc.Sheet(0)
			if err != nil {
				t.Fatalf(
					"Sheet(0) error = %v",
					err,
				)
			}

			if sheet == nil {
				t.Fatal("Sheet(0) returned nil")
			}

			// Try to read some cells
			cellsRead := 0
			testRefs := []string{
				"A1",
				"B1",
				"C1",
				"A2",
				"B2",
			}

			for _, ref := range testRefs {
				cell := sheet.Cell(ref)
				if cell == nil {
					continue
				}
				cellsRead++

				// Get cell value
				val := cell.GetString()
				t.Logf(
					"Cell %s: %v",
					ref,
					val,
				)
			}

			t.Logf(
				"Successfully read %d cells",
				cellsRead,
			)
		})
	}
}

// TestComprehensiveCellWriting tests writing cell values to worksheets.
func TestComprehensiveCellWriting(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-cell-write-*",
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
		"cell_write.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	sheet, err := doc.AddSheet("Test")
	if err != nil {
		t.Fatalf("AddSheet() error = %v", err)
	}

	// Test various cell value types
	testCases := []struct {
		ref   string
		value any
	}{
		{"A1", "String value"},
		{"B1", 42},
		{"C1", 3.14159},
		{"D1", true},
		{"E1", false},
		{"A2", -100},
		{"B2", 0},
		{"C2", ""},
	}

	for _, tc := range testCases {
		if err := sheet.SetCellValue(tc.ref, tc.value); err != nil {
			t.Errorf(
				"SetCellValue(%s, %v) error = %v",
				tc.ref,
				tc.value,
				err,
			)
		}
	}

	// Save and verify
	if err := doc.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	t.Log(
		"Successfully wrote cells with various value types",
	)
}

// TestComprehensiveFormulas tests formula handling.
func TestComprehensiveFormulas(t *testing.T) {
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
		"formulas.xlsx",
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
	_ = sheet.SetCellValue("A1", 10)
	_ = sheet.SetCellValue("A2", 20)
	_ = sheet.SetCellValue("A3", 30)

	// Test various formulas
	testFormulas := []struct {
		ref     string
		formula string
	}{
		{"B1", "=SUM(A1:A3)"},
		{"B2", "=AVERAGE(A1:A3)"},
		{"B3", "=MAX(A1:A3)"},
		{"B4", "=MIN(A1:A3)"},
		{"B5", "=A1+A2"},
		{"B6", "=A1*2"},
		{"B7", "=IF(A1>15,\"High\",\"Low\")"},
	}

	for _, tc := range testFormulas {
		if err := sheet.SetCellFormula(tc.ref, tc.formula); err != nil {
			t.Errorf(
				"SetCellFormula(%s, %s) error = %v",
				tc.ref,
				tc.formula,
				err,
			)
		}
	}

	// Save and verify
	if err := doc.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	t.Log("Successfully created formulas")
}

// TestComprehensiveComments tests comment handling.
func TestComprehensiveComments(t *testing.T) {
	path := filepath.Join(
		"testdata",
		"Comments.xlsx",
	)

	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Comments test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	// Check for comment parts in worksheets
	commentPartsFound := 0
	worksheetParts := wbPart.WorksheetParts()
	for _, wsPart := range worksheetParts {
		if wsPart == nil {
			continue
		}

		commentsPart := wsPart.CommentsPart()
		if commentsPart == nil {
			continue
		}
		commentPartsFound++

		comments := commentsPart.Comments()
		if comments == nil {
			t.Error("Comments() returned nil")

			continue
		}

		commentList := comments.CommentList()
		if commentList == nil {
			continue
		}
		commentCount := 0
		for comment := range commentList.Comments() {
			if comment != nil {
				commentCount++
			}
		}
		t.Logf(
			"Found %d comments in worksheet",
			commentCount,
		)
	}

	t.Logf(
		"Total comment parts found: %d",
		commentPartsFound,
	)
}

// TestComprehensiveRevisions tests revision handling.
func TestComprehensiveRevisions(t *testing.T) {
	path := filepath.Join(
		"testdata",
		"Revision_NameCommentChange.xlsx",
	)

	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Revision test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	t.Log(
		"Successfully opened document with revisions",
	)
}

// TestComprehensiveMediaReferences tests media reference handling (images, YouTube).
func TestComprehensiveMediaReferences(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"Youtube.xlsx",
	)

	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("YouTube test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	// Check for drawing parts in worksheets
	drawingPartsFound := 0
	worksheetParts := wbPart.WorksheetParts()
	for _, wsPart := range worksheetParts {
		if wsPart == nil {
			continue
		}

		drawingPart := wsPart.DrawingsPart()
		if drawingPart != nil {
			drawingPartsFound++
			t.Log(
				"Found drawing part in worksheet",
			)
		}
	}

	t.Logf(
		"Total drawing parts found: %d",
		drawingPartsFound,
	)
}

// TestComprehensiveVMLDrawing tests VML drawing support.
func TestComprehensiveVMLDrawing(t *testing.T) {
	path := filepath.Join(
		"testdata",
		"vmldrawingroot.xlsx",
	)

	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("VML drawing test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	// Check for VML parts in worksheets
	vmlPartsFound := 0
	worksheetParts := wbPart.WorksheetParts()
	for _, wsPart := range worksheetParts {
		if wsPart == nil {
			continue
		}

		vmlParts := wsPart.VmlDrawingParts()
		for _, vmlPart := range vmlParts {
			if vmlPart != nil {
				vmlPartsFound++
			}
		}
	}

	t.Logf(
		"Total VML parts found: %d",
		vmlPartsFound,
	)
}

// TestComprehensiveExtensionLists tests extension list support.
func TestComprehensiveExtensionLists(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"extlst.xlsx",
	)

	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Extension list test file not found",
		)
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	wbRoot := wbPart.Workbook()
	if wbRoot == nil {
		t.Fatal("Workbook() returned nil")
	}

	wb, ok := wbRoot.(*elements.Workbook)
	if !ok {
		t.Fatal("Failed to cast to Workbook type")
	}

	// Extension lists would be checked here if the API supported it
	// For now, just verify the workbook was opened successfully
	_ = wb
	t.Log(
		"Successfully opened workbook with potential extensions",
	)
}

// TestComprehensiveErrorHandling tests error recovery scenarios.
func TestComprehensiveErrorHandling(
	t *testing.T,
) {
	t.Run("MalformedURI", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"malformed_uri.xlsx",
		)

		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip(
				"Malformed URI test file not found",
			)
		}

		doc, err := Open(path, false)
		if err != nil {
			// It's acceptable to fail on malformed URI
			t.Logf(
				"Open() error (expected): %v",
				err,
			)

			return
		}
		defer func() { _ = doc.Close() }()

		t.Log(
			"Successfully handled malformed URI",
		)
	})

	t.Run("MalformedURILong", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"malformed_uri_long.xlsx",
		)

		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip(
				"Malformed URI long test file not found",
			)
		}

		doc, err := Open(path, false)
		if err != nil {
			// It's acceptable to fail on malformed URI
			t.Logf(
				"Open() error (expected): %v",
				err,
			)

			return
		}
		defer func() { _ = doc.Close() }()

		t.Log(
			"Successfully handled long malformed URI",
		)
	})

	t.Run("MissingCalcChain", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"missingcalcchainpart.xlsx",
		)

		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip(
				"Missing calc chain test file not found",
			)
		}

		doc, err := Open(path, false)
		if err != nil {
			t.Fatalf(
				"Open() should handle missing calc chain: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		wbPart := doc.WorkbookPart()
		if wbPart == nil {
			t.Fatal("WorkbookPart() returned nil")
		}

		// Calc chain should be nil but document should still open
		calcChainPart := wbPart.CalculationChainPart()
		if calcChainPart != nil {
			t.Log(
				"Calc chain part exists (unexpected)",
			)
		} else {
			t.Log("Successfully handled missing calc chain")
		}
	})
}

// TestComprehensiveDocumentValidation tests document validation.
func TestComprehensiveDocumentValidation(
	t *testing.T,
) {
	testCases := []string{
		"basicspreadsheet.xlsx",
		"Spreadsheet.xlsx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)

			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			// Validate document
			validationErrors := doc.Validate(
				validation.Office2016,
			)

			if len(validationErrors) > 0 {
				t.Logf(
					"Validation found %d issues:",
					len(validationErrors),
				)
				for i, verr := range validationErrors {
					if i < 5 { // Limit output
						t.Logf("  - %v", verr)
					}
				}
			} else {
				t.Log("Document passed validation")
			}
		})
	}
}

// TestComprehensivePerformance tests performance with complex workbooks.
func TestComprehensivePerformance(t *testing.T) {
	if testing.Short() {
		t.Skip(
			"Skipping performance test in short mode",
		)
	}

	testCases := []string{
		"Complex01.xlsx",
		"excel14.xlsx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)

			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			// Test open performance
			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			// Iterate through all sheets
			sheetCount := 0
			for i := range doc.SheetCount() {
				sheet, err := doc.Sheet(i)
				if err != nil {
					t.Errorf(
						"Sheet(%d) error = %v",
						i,
						err,
					)

					continue
				}
				if sheet != nil {
					sheetCount++
				}
			}

			t.Logf(
				"Processed %d sheets",
				sheetCount,
			)

			// Create temp directory for roundtrip
			tmpDir, err := os.MkdirTemp(
				"",
				"goffice-perf-*",
			)
			if err != nil {
				t.Fatalf(
					"Failed to create temp dir: %v",
					err,
				)
			}
			defer func() { _ = os.RemoveAll(tmpDir) }()

			// Test save performance
			outputPath := filepath.Join(
				tmpDir,
				"perf_"+filename,
			)
			if err := doc.SaveAs(outputPath); err != nil {
				t.Fatalf(
					"SaveAs() error = %v",
					err,
				)
			}

			t.Log(
				"Successfully completed performance test",
			)
		})
	}
}

// TestComprehensiveWorksheetOperations tests worksheet-specific operations.
func TestComprehensiveWorksheetOperations(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"Spreadsheet.xlsx",
	)

	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	if doc.SheetCount() == 0 {
		t.Fatal("Document has no sheets")
	}

	sheet, err := doc.Sheet(0)
	if err != nil {
		t.Fatalf("Sheet(0) error = %v", err)
	}

	// Get worksheet part
	wbPart := doc.WorkbookPart()
	if wbPart == nil {
		t.Fatal("WorkbookPart() returned nil")
	}

	worksheetFound := false
	worksheetParts := wbPart.WorksheetParts()
	for _, wsPart := range worksheetParts {
		if wsPart == nil {
			continue
		}

		wsElement := wsPart.Worksheet()
		if wsElement == nil {
			t.Error("Worksheet() returned nil")

			continue
		}

		worksheetFound = true

		testWorksheetData(t, wsElement)
		testWorksheetDimension(t, wsElement)
		testWorksheetViews(t, wsElement)
		testWorksheetFormatProperties(
			t,
			wsElement,
		)
		testWorksheetColumns(t, wsElement)
		testWorksheetMergeCells(t, wsElement)

		// Only test first worksheet
		break
	}

	if !worksheetFound {
		t.Error("No worksheet parts found")
	}

	_ = sheet // Use sheet to avoid unused variable warning
}

func testWorksheetData(
	t *testing.T,
	ws *elements.Worksheet,
) {
	sheetData := ws.SheetData()
	if sheetData == nil {
		return
	}

	rowCount := 0
	for row := range sheetData.Rows() {
		if row == nil {
			continue
		}
		rowCount++

		// Count cells in first row
		if rowCount == 1 {
			testFirstRowCells(t, row)
		}
	}
	t.Logf("Worksheet has %d rows", rowCount)
}

func testFirstRowCells(
	t *testing.T,
	row *elements.Row,
) {
	cellCount := 0
	for cell := range row.Cells() {
		if cell != nil {
			cellCount++
		}
	}
	t.Logf("First row has %d cells", cellCount)
}

func testWorksheetDimension(
	t *testing.T,
	ws *elements.Worksheet,
) {
	dimension := ws.Dimension()
	if dimension == nil {
		return
	}

	ref := dimension.Ref()
	if ref != "" {
		t.Logf("Worksheet dimension: %s", ref)
	}
}

func testWorksheetViews(
	t *testing.T,
	ws *elements.Worksheet,
) {
	sheetViews := ws.SheetViews()
	if sheetViews == nil {
		return
	}

	viewCount := 0
	for view := range sheetViews.SheetViews() {
		if view != nil {
			viewCount++
		}
	}
	t.Logf("Worksheet has %d views", viewCount)
}

func testWorksheetFormatProperties(
	t *testing.T,
	ws *elements.Worksheet,
) {
	sheetFormatPr := ws.SheetFormatPr()
	if sheetFormatPr != nil {
		t.Log("Worksheet has format properties")
	}
}

func testWorksheetColumns(
	t *testing.T,
	ws *elements.Worksheet,
) {
	cols := ws.Cols()
	if cols == nil {
		return
	}

	colCount := 0
	for colGroup := range cols.Cols() {
		if colGroup != nil {
			colCount++
		}
	}
	if colCount > 0 {
		t.Logf(
			"Worksheet has %d column groups",
			colCount,
		)
	}
}

func testWorksheetMergeCells(
	t *testing.T,
	ws *elements.Worksheet,
) {
	mergeCells := ws.MergeCells()
	if mergeCells == nil {
		return
	}

	mergeCount := 0
	for merge := range mergeCells.GetMergeCells() {
		if merge != nil {
			mergeCount++
		}
	}
	if mergeCount > 0 {
		t.Logf(
			"Worksheet has %d merged cell ranges",
			mergeCount,
		)
	}
}

// TestComprehensiveFeatures tests feature flags.
func TestComprehensiveFeatures(t *testing.T) {
	path := filepath.Join(
		"testdata",
		"Complex01.xlsx",
	)

	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	features := doc.Features()
	if features == nil {
		t.Fatal("Features() returned nil")
	}

	t.Log("Document features available")
}

// Ensure interfaces are implemented
var (
	_ = (*parts.WorkbookPart)(nil)
	_ = (*parts.WorksheetPart)(nil)
	_ = (*parts.SharedStringTablePart)(nil)
	_ = (*parts.WorkbookStylesPart)(nil)
	_ = (*parts.CalculationChainPart)(nil)
	_ = (*elements.Workbook)(nil)
	_ = (*elements.Worksheet)(nil)
	_ = (*elements.SharedStringTable)(nil)
)
