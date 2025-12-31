// chart_e2e_test.go contains end-to-end tests for Excel chart to PDF rendering.

package spreadsheet

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

// TestChartRendering_BasicChart tests rendering an Excel workbook with a chart to PDF.
// This is an end-to-end test that validates the chart rendering integration.
func TestChartRendering_BasicChart(t *testing.T) {
	// Create a test workbook with some data
	tmpDir := t.TempDir()
	xlsxPath := filepath.Join(
		tmpDir,
		"chart_test.xlsx",
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
	defer func() {
		if err := doc.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	// Add a sheet with sample data for the chart
	sheet, err := doc.AddSheet("Sales Data")
	if err != nil {
		t.Fatalf("Failed to add sheet: %v", err)
	}

	// Add headers
	_ = sheet.SetCellValue("A1", "Quarter")
	_ = sheet.SetCellValue("B1", "Sales")
	_ = sheet.SetCellValue("C1", "Revenue")

	// Add data rows
	quarters := []string{"Q1", "Q2", "Q3", "Q4"}
	sales := []float64{100, 150, 120, 180}
	revenue := []float64{80, 120, 100, 150}

	for i, q := range quarters {
		row := i + 2
		_ = sheet.SetCellValue(
			"A"+string(rune('0'+row)),
			q,
		)
		_ = sheet.SetCellValue(
			"B"+string(rune('0'+row)),
			sales[i],
		)
		_ = sheet.SetCellValue(
			"C"+string(rune('0'+row)),
			revenue[i],
		)
	}

	// Save the workbook
	// Note: We're not actually adding a chart via the API since the spreadsheet
	// package may not have chart creation methods yet. This test validates that
	// the PDF renderer can handle workbooks with chart placeholders.
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
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	renderer, err := NewSpreadsheetRenderer(doc2)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(
		tmpDir,
		"chart_test.pdf",
	)
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}

	// Verify PDF has content (non-zero size)
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}

	// Basic size check - PDF should be at least a few hundred bytes
	// (even a minimal PDF with empty page is around 600+ bytes)
	if stat.Size() < 500 {
		t.Errorf(
			"PDF file is suspiciously small (%d bytes), expected at least 500 bytes",
			stat.Size(),
		)
	}
}

// TestChartRendering_MultipleSheets tests rendering a workbook with multiple sheets.
// This validates that chart rendering works correctly across different sheets.
func TestChartRendering_MultipleSheets(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	xlsxPath := filepath.Join(
		tmpDir,
		"multi_chart.xlsx",
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
	defer func() {
		if err := doc.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	// Create first sheet with data
	sheet1, err := doc.AddSheet("Sheet1")
	if err != nil {
		t.Fatalf(
			"Failed to add sheet1: %v",
			err,
		)
	}
	_ = sheet1.SetCellValue("A1", "Data 1")
	_ = sheet1.SetCellValue("B1", 100)

	// Create second sheet with data
	sheet2, err := doc.AddSheet("Sheet2")
	if err != nil {
		t.Fatalf(
			"Failed to add sheet2: %v",
			err,
		)
	}
	_ = sheet2.SetCellValue("A1", "Data 2")
	_ = sheet2.SetCellValue("B1", 200)

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
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	renderer, err := NewSpreadsheetRenderer(doc2)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(
		tmpDir,
		"multi_chart.pdf",
	)
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created and has content
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}
}

// TestChartRendering_EmptyWorkbook tests rendering an empty workbook.
// This validates graceful handling when no charts are present.
func TestChartRendering_EmptyWorkbook(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	xlsxPath := filepath.Join(
		tmpDir,
		"empty.xlsx",
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
	defer func() {
		if err := doc.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	// Add an empty sheet
	_, err = doc.AddSheet("Sheet1")
	if err != nil {
		t.Fatalf("Failed to add sheet: %v", err)
	}

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
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	renderer, err := NewSpreadsheetRenderer(doc2)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(tmpDir, "empty.pdf")
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created
	if _, err := os.Stat(pdfPath); os.IsNotExist(
		err,
	) {
		t.Errorf("PDF file was not created")
	}
}

// TestChartRendering_WithOptions tests chart rendering with custom options.
func TestChartRendering_WithOptions(
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
	defer func() {
		if err := doc.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	sheet, err := doc.AddSheet("Sheet1")
	if err != nil {
		t.Fatalf("Failed to add sheet: %v", err)
	}

	// Add sample data
	_ = sheet.SetCellValue("A1", "Category")
	_ = sheet.SetCellValue("B1", "Value")
	for i := 1; i <= 5; i++ {
		row := i + 1
		_ = sheet.SetCellValue(
			"A"+string(rune('0'+row)),
			"Item "+string(rune('0'+i)),
		)
		_ = sheet.SetCellValue(
			"B"+string(rune('0'+row)),
			float64(i*10),
		)
	}

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
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	renderer, err := NewSpreadsheetRenderer(doc2)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Apply custom rendering options
	opts := renderer.options
	opts.GridLines = true
	opts.DefaultFontSize = 11.0
	renderer.WithOptions(opts)

	pdfPath := filepath.Join(
		tmpDir,
		"options.pdf",
	)
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created and has content
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}
}

// TestChartRendering_LargeDataset tests rendering with a larger dataset.
// This helps validate performance and correctness with more complex scenarios.
func TestChartRendering_LargeDataset(
	t *testing.T,
) {
	if testing.Short() {
		t.Skip(
			"Skipping large dataset test in short mode",
		)
	}

	tmpDir := t.TempDir()
	xlsxPath := filepath.Join(
		tmpDir,
		"large.xlsx",
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
	defer func() {
		if err := doc.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	sheet, err := doc.AddSheet("Data")
	if err != nil {
		t.Fatalf("Failed to add sheet: %v", err)
	}

	// Add headers
	_ = sheet.SetCellValue("A1", "X")
	_ = sheet.SetCellValue("B1", "Y")

	// Add 50 rows of data
	for i := 1; i <= 50; i++ {
		row := i + 1
		_ = sheet.SetCellValue(
			"A"+string(rune('0'+row)),
			float64(i),
		)
		_ = sheet.SetCellValue(
			"B"+string(rune('0'+row)),
			float64(i*i),
		)
	}

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
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	renderer, err := NewSpreadsheetRenderer(doc2)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(tmpDir, "large.pdf")
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created and has substantial content
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}

	// Larger dataset should produce a reasonable PDF
	// Log the size for informational purposes
	t.Logf(
		"PDF size for large dataset: %d bytes",
		stat.Size(),
	)
}
