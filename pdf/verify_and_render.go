package main

import (
	"fmt"
	"log"

	pdfspreadsheet "github.com/connerohnesorge/goffice-pdf/spreadsheet"
	"github.com/connerohnesorge/goffice/spreadsheet"
)

func main() {
	// Open the test file
	doc, err := spreadsheet.Open(
		"../examples/ml_results/test_text.xlsx",
		false,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer doc.Close()

	// Verify data is present
	fmt.Println("Verifying data in Excel file...")
	for sheet := range doc.Sheets() {
		fmt.Printf("\nSheet: %s\n", sheet.Name())

		// Check title cells
		a1 := sheet.Cell("A1").GetString()
		a2 := sheet.Cell("A2").GetString()
		a4 := sheet.Cell("A4").GetString()
		a5 := sheet.Cell("A5").GetString()

		fmt.Printf("  A1: %s\n", a1)
		fmt.Printf("  A2: %s\n", a2)
		fmt.Printf("  A4: %s\n", a4)
		fmt.Printf("  A5: %s\n", a5)
	}

	// Render to PDF
	fmt.Println("\nRendering to PDF...")
	renderer, err := pdfspreadsheet.NewSpreadsheetRenderer(
		doc,
	)
	if err != nil {
		log.Fatal(err)
	}

	opts := pdfspreadsheet.DefaultRenderOptions()
	renderer.WithOptions(opts)

	outputPath := "../examples/ml_results/test_text.pdf"
	if err := renderer.Render(outputPath); err != nil {
		log.Fatalf("Failed to render: %v", err)
	}

	fmt.Printf(
		"Successfully rendered to: %s\n",
		outputPath,
	)
}
