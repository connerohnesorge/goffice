package main

import (
	"fmt"
	"log"

	pdfspreadsheet "github.com/connerohnesorge/goffice-pdf/spreadsheet"
	"github.com/connerohnesorge/goffice/spreadsheet"
)

func main() {
	// Create a simple document
	doc, err := spreadsheet.Create(
		"../examples/ml_results/simple_test.xlsx",
		spreadsheet.DocTypeWorkbook,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer doc.Close()

	sheet, err := doc.AddSheet("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	// Set just one cell with text
	sheet.Cell("A1").SetString("Hello World")

	// Verify it's set
	val := sheet.Cell("A1").GetString()
	fmt.Printf("Cell A1 value: '%s'\n", val)

	if val == "" {
		log.Fatal("Cell value is empty!")
	}

	// Render
	renderer, err := pdfspreadsheet.NewSpreadsheetRenderer(
		doc,
	)
	if err != nil {
		log.Fatal(err)
	}

	opts := pdfspreadsheet.DefaultRenderOptions()
	renderer.WithOptions(opts)

	outputPath := "../examples/ml_results/simple_test.pdf"
	if err := renderer.Render(outputPath); err != nil {
		log.Fatalf("Render failed: %v", err)
	}

	fmt.Printf("Rendered to: %s\n", outputPath)
}
