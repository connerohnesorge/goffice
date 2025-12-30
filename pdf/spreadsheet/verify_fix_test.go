package spreadsheet_test

import (
	"testing"

	pdfspreadsheet "github.com/connerohnesorge/goffice-pdf/spreadsheet"
	"github.com/connerohnesorge/goffice/spreadsheet"
)

// TestVerifyTextInPDF verifies that the fix for text rendering works correctly
// by converting the ml_results_example.xlsx to PDF and checking it renders without errors
func TestVerifyTextInPDF(t *testing.T) {
	// Open the ml_results example file
	doc, err := spreadsheet.Open(
		"../../examples/ml_results/ml_results_example.xlsx",
		false,
	)
	if err != nil {
		t.Fatalf(
			"Failed to open Excel file: %v",
			err,
		)
	}
	defer doc.Close()

	// Create a renderer
	renderer, err := pdfspreadsheet.NewSpreadsheetRenderer(
		doc,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Use default options
	opts := pdfspreadsheet.DefaultRenderOptions()
	renderer.WithOptions(opts)

	// Render to PDF
	outputPath := "../../examples/ml_results/ml_results_verify_fix.pdf"
	if err := renderer.Render(outputPath); err != nil {
		t.Fatalf(
			"Failed to render to PDF: %v",
			err,
		)
	}

	t.Logf(
		"Successfully rendered PDF to: %s",
		outputPath,
	)
	t.Log("Manual verification required:")
	t.Log(
		"  1. Open the PDF and check that all titles are visible",
	)
	t.Log(
		"  2. Verify column headers are present",
	)
	t.Log(
		"  3. Confirm all text labels render correctly",
	)
}
