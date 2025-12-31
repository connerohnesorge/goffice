package main

import (
	"fmt"
	"log"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

func main() {
	// Create a new document
	doc, err := spreadsheet.Create(
		"../examples/ml_results/test_text.xlsx",
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

	// Save
	if err := doc.SaveAs("../examples/ml_results/test_text.xlsx"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(
		"Created test_text.xlsx successfully",
	)
}
