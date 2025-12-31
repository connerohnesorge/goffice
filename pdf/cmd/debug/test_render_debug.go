package main

import (
	"fmt"
	"log"

	pdfspreadsheet "github.com/connerohnesorge/goffice-pdf/spreadsheet"
	"github.com/connerohnesorge/goffice/spreadsheet"
)

func main() {
	// Open the Excel file
	doc, err := spreadsheet.Open(
		"../examples/ml_results/ml_results_example.xlsx",
		false,
	)
	if err != nil {
		log.Fatalf(
			"Failed to open Excel file: %v",
			err,
		)
	}
	defer doc.Close()

	// Check what's in the document
	wb := doc.WorkbookPart()
	if wb == nil {
		log.Fatal("No workbook part!")
	}

	// Try to get sheets using the iterator
	sheetCount := 0
	for sheet := range doc.Sheets() {
		sheetCount++
		fmt.Printf(
			"Sheet %d: name=%s\n",
			sheetCount,
			sheet.Name(),
		)

		// Try to get some cells
		cellA1 := sheet.Cell("A1")
		if cellA1 != nil {
			val := cellA1.GetString()
			fmt.Printf("  A1 value: %v\n", val)
		}

		cellA2 := sheet.Cell("A2")
		if cellA2 != nil {
			val := cellA2.GetString()
			fmt.Printf("  A2 value: %v\n", val)
		}

		cellA4 := sheet.Cell("A4")
		if cellA4 != nil {
			val := cellA4.GetString()
			fmt.Printf("  A4 value: %v\n", val)
		}
	}

	fmt.Printf(
		"\nTotal sheets found: %d\n",
		sheetCount,
	)

	// Now try rendering
	renderer, err := pdfspreadsheet.NewSpreadsheetRenderer(
		doc,
	)
	if err != nil {
		log.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	opts := pdfspreadsheet.DefaultRenderOptions()
	renderer.WithOptions(opts)

	outputPath := "../examples/ml_results/ml_results_debug.pdf"
	if err := renderer.Render(outputPath); err != nil {
		log.Fatalf("Failed to render: %v", err)
	}

	fmt.Printf("\nRendered to: %s\n", outputPath)
}
