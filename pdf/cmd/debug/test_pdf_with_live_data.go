package main

import (
	"fmt"
	"log"

	pdfspreadsheet "github.com/connerohnesorge/goffice-pdf/spreadsheet"
	"github.com/connerohnesorge/goffice/spreadsheet"
)

func main() {
	// Create a new document
	doc, err := spreadsheet.Create(
		"../examples/ml_results/live_test.xlsx",
		spreadsheet.DocTypeWorkbook,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer doc.Close()

	// Add a sheet with various text content
	sheet, err := doc.AddSheet("Test Sheet")
	if err != nil {
		log.Fatal(err)
	}

	// Add title
	sheet.Cell("A1").
		SetString("Deep Learning Model Performance Comparison")
	sheet.Cell("A2").
		SetString("Neural Network Architecture Benchmark")

	// Add headers
	sheet.Cell("A4").SetString("Model Name")
	sheet.Cell("B4").SetString("Parameters (M)")
	sheet.Cell("C4").
		SetString("Training Time (hrs)")
	sheet.Cell("D4").SetString("Accuracy %")

	// Add data rows
	sheet.Cell("A5").SetString("ResNet-50")
	sheet.Cell("B5").SetNumber(25.6)
	sheet.Cell("C5").SetNumber(12.5)
	sheet.Cell("D5").SetNumber(76.15)

	sheet.Cell("A6").SetString("VGG-16")
	sheet.Cell("B6").SetNumber(138.4)
	sheet.Cell("C6").SetNumber(18.2)
	sheet.Cell("D6").SetNumber(71.3)

	sheet.Cell("A7").SetString("EfficientNet-B0")
	sheet.Cell("B7").SetNumber(5.3)
	sheet.Cell("C7").SetNumber(8.1)
	sheet.Cell("D7").SetNumber(77.1)

	// Verify data is in memory
	fmt.Println("Data in memory:")
	fmt.Printf(
		"  A1: %s\n",
		sheet.Cell("A1").GetString(),
	)
	fmt.Printf(
		"  A2: %s\n",
		sheet.Cell("A2").GetString(),
	)
	fmt.Printf(
		"  A4: %s\n",
		sheet.Cell("A4").GetString(),
	)
	fmt.Printf(
		"  A5: %s\n",
		sheet.Cell("A5").GetString(),
	)
	fmt.Printf(
		"  B5: %v\n",
		sheet.Cell("B5").GetNumber(),
	)

	// Render to PDF while data is still in memory
	fmt.Println("\nRendering to PDF...")
	renderer, err := pdfspreadsheet.NewSpreadsheetRenderer(
		doc,
	)
	if err != nil {
		log.Fatal(err)
	}

	opts := pdfspreadsheet.DefaultRenderOptions()
	renderer.WithOptions(opts)

	outputPath := "../examples/ml_results/live_test.pdf"
	if err := renderer.Render(outputPath); err != nil {
		log.Fatalf("Failed to render: %v", err)
	}

	fmt.Printf(
		"Successfully rendered to: %s\n",
		outputPath,
	)

	// Save the Excel file too
	if err := doc.SaveAs("../examples/ml_results/live_test.xlsx"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved Excel file")
}
