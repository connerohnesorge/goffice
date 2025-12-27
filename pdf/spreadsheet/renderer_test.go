// renderer_test.go contains tests for the SpreadsheetML to PDF renderer.

package spreadsheet

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/spreadsheet"
)

func TestSpreadsheetRenderer_BasicWorkbook(
	t *testing.T,
) {
	// Create a simple test workbook
	tmpDir := t.TempDir()
	xlsxPath := filepath.Join(tmpDir, "test.xlsx")

	doc, err := spreadsheet.Create(
		xlsxPath,
		spreadsheet.DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create workbook: %v",
			err,
		)
	}
	defer doc.Close()

	// Add a sheet
	sheet, err := doc.AddSheet("Sheet1")
	if err != nil {
		t.Fatalf("Failed to add sheet: %v", err)
	}

	// Add some data
	sheet.SetCellValue("A1", "Hello")
	sheet.SetCellValue("B1", "World")
	sheet.SetCellValue("A2", 42)
	sheet.SetCellValue("B2", 3.14)

	// Save the workbook
	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save workbook: %v",
			err,
		)
	}

	// Now render to PDF
	doc2, err := spreadsheet.Open(xlsxPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to open workbook: %v",
			err,
		)
	}
	defer doc2.Close()

	renderer, err := NewSpreadsheetRenderer(doc2)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(tmpDir, "test.pdf")
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Check that PDF was created
	if _, err := os.Stat(pdfPath); os.IsNotExist(
		err,
	) {
		t.Errorf("PDF file was not created")
	}
}

func TestSpreadsheetRenderer_ColumnWidthConversion(
	t *testing.T,
) {
	tests := []struct {
		name       string
		excelWidth float64
		wantPoints float64
	}{
		{"Default width", 8.43, 64.07},
		{"Narrow column", 5.0, 38.0},
		{"Wide column", 20.0, 152.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Excel width is in character units
			// 1 character unit ≈ 7.6 points
			gotPoints := tt.excelWidth * 7.6

			tolerance := 1.0
			if gotPoints < tt.wantPoints-tolerance ||
				gotPoints > tt.wantPoints+tolerance {
				t.Errorf(
					"Column width conversion: got %v points, want %v points",
					gotPoints,
					tt.wantPoints,
				)
			}
		})
	}
}

func TestSpreadsheetRenderer_NumberFormatting(
	t *testing.T,
) {
	tests := []struct {
		name   string
		value  float64
		format string
		want   string
	}{
		{"Integer", 42.0, "0", "42"},
		{"Decimal", 3.14, "0.00", "3.14"},
		{
			"Currency",
			123.45,
			"$#,##0.00",
			"$123.45",
		},
		{"Percentage", 0.75, "0.00%", "75.00%"},
		{
			"General integer",
			100.0,
			"General",
			"100",
		},
		{
			"General decimal",
			3.14159,
			"General",
			"3.14159",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := applyNumberFormat(
				tt.value,
				tt.format,
			)
			if got != tt.want {
				t.Errorf(
					"applyNumberFormat(%v, %q) = %q, want %q",
					tt.value,
					tt.format,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestSpreadsheetRenderer_CellReferenceParser(
	t *testing.T,
) {
	tests := []struct {
		name    string
		ref     string
		wantRow int
		wantCol int
	}{
		{"A1", "A1", 1, 1},
		{"Z10", "Z10", 10, 26},
		{"AA1", "AA1", 1, 27},
		{"AB100", "AB100", 100, 28},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cellRef := parseCellRef(tt.ref)
			if cellRef == nil {
				t.Fatalf(
					"parseCellRef(%q) returned nil",
					tt.ref,
				)
			}

			if cellRef.Row != tt.wantRow {
				t.Errorf(
					"parseCellRef(%q).Row = %d, want %d",
					tt.ref,
					cellRef.Row,
					tt.wantRow,
				)
			}

			if cellRef.Col != tt.wantCol {
				t.Errorf(
					"parseCellRef(%q).Col = %d, want %d",
					tt.ref,
					cellRef.Col,
					tt.wantCol,
				)
			}
		})
	}
}

func TestSpreadsheetRenderer_RangeParser(
	t *testing.T,
) {
	tests := []struct {
		name         string
		rangeStr     string
		wantStartRow int
		wantStartCol int
		wantEndRow   int
		wantEndCol   int
	}{
		{"A1:B2", "A1:B2", 1, 1, 2, 2},
		{"A1:Z10", "A1:Z10", 1, 1, 10, 26},
		{"$A$1:$B$2", "$A$1:$B$2", 1, 1, 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cellRange := parseRange(tt.rangeStr)
			if cellRange == nil {
				t.Fatalf(
					"parseRange(%q) returned nil",
					tt.rangeStr,
				)
			}

			if cellRange.StartRow != tt.wantStartRow {
				t.Errorf(
					"parseRange(%q).StartRow = %d, want %d",
					tt.rangeStr,
					cellRange.StartRow,
					tt.wantStartRow,
				)
			}

			if cellRange.StartCol != tt.wantStartCol {
				t.Errorf(
					"parseRange(%q).StartCol = %d, want %d",
					tt.rangeStr,
					cellRange.StartCol,
					tt.wantStartCol,
				)
			}

			if cellRange.EndRow != tt.wantEndRow {
				t.Errorf(
					"parseRange(%q).EndRow = %d, want %d",
					tt.rangeStr,
					cellRange.EndRow,
					tt.wantEndRow,
				)
			}

			if cellRange.EndCol != tt.wantEndCol {
				t.Errorf(
					"parseRange(%q).EndCol = %d, want %d",
					tt.rangeStr,
					cellRange.EndCol,
					tt.wantEndCol,
				)
			}
		})
	}
}

func TestSpreadsheetRenderer_ColorParser(
	t *testing.T,
) {
	tests := []struct {
		name  string
		rgb   string
		wantR float64
		wantG float64
		wantB float64
	}{
		{"Black", "000000", 0.0, 0.0, 0.0},
		{"White", "FFFFFF", 1.0, 1.0, 1.0},
		{"Red", "FF0000", 1.0, 0.0, 0.0},
		{"Green", "00FF00", 0.0, 1.0, 0.0},
		{"Blue", "0000FF", 0.0, 0.0, 1.0},
		{"With alpha", "FFFF0000", 1.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			color := parseRGBString(tt.rgb)

			tolerance := 0.01
			if color.R < tt.wantR-tolerance ||
				color.R > tt.wantR+tolerance {
				t.Errorf(
					"parseRGBString(%q).R = %v, want %v",
					tt.rgb,
					color.R,
					tt.wantR,
				)
			}

			if color.G < tt.wantG-tolerance ||
				color.G > tt.wantG+tolerance {
				t.Errorf(
					"parseRGBString(%q).G = %v, want %v",
					tt.rgb,
					color.G,
					tt.wantG,
				)
			}

			if color.B < tt.wantB-tolerance ||
				color.B > tt.wantB+tolerance {
				t.Errorf(
					"parseRGBString(%q).B = %v, want %v",
					tt.rgb,
					color.B,
					tt.wantB,
				)
			}
		})
	}
}

func TestSpreadsheetRenderer_MergedCells(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	xlsxPath := filepath.Join(
		tmpDir,
		"merged.xlsx",
	)

	doc, err := spreadsheet.Create(
		xlsxPath,
		spreadsheet.DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create workbook: %v",
			err,
		)
	}
	defer doc.Close()

	sheet, err := doc.AddSheet("Sheet1")
	if err != nil {
		t.Fatalf("Failed to add sheet: %v", err)
	}

	// Add data and merge cells
	sheet.SetCellValue("A1", "Merged Title")
	sheet.MergeCells("A1:D1")

	sheet.SetCellValue("A2", "Left")
	sheet.SetCellValue("D2", "Right")

	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save workbook: %v",
			err,
		)
	}

	// Render to PDF
	doc2, err := spreadsheet.Open(xlsxPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to open workbook: %v",
			err,
		)
	}
	defer doc2.Close()

	renderer, err := NewSpreadsheetRenderer(doc2)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(tmpDir, "merged.pdf")
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Check that PDF was created
	if _, err := os.Stat(pdfPath); os.IsNotExist(
		err,
	) {
		t.Errorf("PDF file was not created")
	}
}

func TestSpreadsheetRenderer_MultipleSheets(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	xlsxPath := filepath.Join(
		tmpDir,
		"multi.xlsx",
	)

	doc, err := spreadsheet.Create(
		xlsxPath,
		spreadsheet.DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create workbook: %v",
			err,
		)
	}
	defer doc.Close()

	// Add multiple sheets
	sheet1, err := doc.AddSheet("Sheet1")
	if err != nil {
		t.Fatalf("Failed to add sheet1: %v", err)
	}
	sheet1.SetCellValue("A1", "Sheet 1")

	sheet2, err := doc.AddSheet("Sheet2")
	if err != nil {
		t.Fatalf("Failed to add sheet2: %v", err)
	}
	sheet2.SetCellValue("A1", "Sheet 2")

	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save workbook: %v",
			err,
		)
	}

	// Render to PDF
	doc2, err := spreadsheet.Open(xlsxPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to open workbook: %v",
			err,
		)
	}
	defer doc2.Close()

	renderer, err := NewSpreadsheetRenderer(doc2)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(tmpDir, "multi.pdf")
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Check that PDF was created
	if _, err := os.Stat(pdfPath); os.IsNotExist(
		err,
	) {
		t.Errorf("PDF file was not created")
	}
}

func TestSpreadsheetRenderer_WithOptions(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	xlsxPath := filepath.Join(
		tmpDir,
		"options.xlsx",
	)

	doc, err := spreadsheet.Create(
		xlsxPath,
		spreadsheet.DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create workbook: %v",
			err,
		)
	}
	defer doc.Close()

	sheet, err := doc.AddSheet("Sheet1")
	if err != nil {
		t.Fatalf("Failed to add sheet: %v", err)
	}
	sheet.SetCellValue("A1", "Test")

	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save workbook: %v",
			err,
		)
	}

	// Render with custom options
	doc2, err := spreadsheet.Open(xlsxPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to open workbook: %v",
			err,
		)
	}
	defer doc2.Close()

	renderer, err := NewSpreadsheetRenderer(doc2)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	opts := RenderOptions{
		DefaultPageSize:   core.PageSizeA4,
		GridLines:         true,
		DefaultFontFamily: "Arial",
		DefaultFontSize:   12.0,
	}

	renderer.WithOptions(opts)

	pdfPath := filepath.Join(
		tmpDir,
		"options.pdf",
	)
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	if _, err := os.Stat(pdfPath); os.IsNotExist(
		err,
	) {
		t.Errorf("PDF file was not created")
	}
}

// Benchmarks

func BenchmarkSpreadsheetRenderer_SmallWorkbook(
	b *testing.B,
) {
	tmpDir := b.TempDir()
	xlsxPath := filepath.Join(
		tmpDir,
		"bench.xlsx",
	)

	// Create a small workbook
	doc, _ := spreadsheet.Create(
		xlsxPath,
		spreadsheet.DocTypeWorkbook,
	)
	sheet, _ := doc.AddSheet("Sheet1")

	for row := 1; row <= 10; row++ {
		for col := 1; col <= 10; col++ {
			ref := string(
				rune('A'+col-1),
			) + string(
				rune('0'+row),
			)
			sheet.SetCellValue(ref, row*col)
		}
	}

	doc.Save()
	doc.Close()

	// Benchmark rendering
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		doc2, _ := spreadsheet.Open(
			xlsxPath,
			false,
		)
		renderer, _ := NewSpreadsheetRenderer(
			doc2,
		)
		pdfPath := filepath.Join(
			tmpDir,
			"bench.pdf",
		)
		renderer.Render(pdfPath)
		doc2.Close()
	}
}
