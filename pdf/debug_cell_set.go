package main

import (
	"fmt"
	"log"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

func main() {
	// Create a new document
	doc, err := spreadsheet.Create(
		"../examples/ml_results/debug_test.xlsx",
		spreadsheet.DocTypeWorkbook,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer doc.Close()

	// Add a sheet
	sheet, err := doc.AddSheet("Debug Sheet")
	if err != nil {
		log.Fatal(err)
	}

	// Set a value
	fmt.Println("Setting A1 to 'Test Value'...")
	sheet.Cell("A1").SetString("Test Value")

	// Try to read it back immediately
	val := sheet.Cell("A1").GetString()
	fmt.Printf("Read back A1: '%s'\n", val)

	// Set a number
	fmt.Println("Setting B1 to 123.45...")
	sheet.Cell("B1").SetNumber(123.45)

	num := sheet.Cell("B1").GetNumber()
	fmt.Printf("Read back B1: %v\n", num)

	// Save
	if err := doc.SaveAs("../examples/ml_results/debug_test.xlsx"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nSaved. Now reopening...")
	doc.Close()

	// Reopen and check
	doc2, err := spreadsheet.Open(
		"../examples/ml_results/debug_test.xlsx",
		false,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer doc2.Close()

	for sheet2 := range doc2.Sheets() {
		val2 := sheet2.Cell("A1").GetString()
		num2 := sheet2.Cell("B1").GetNumber()
		fmt.Printf(
			"After reopen - A1: '%s', B1: %v\n",
			val2,
			num2,
		)
	}
}
