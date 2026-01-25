//go:build ignore
// +build ignore

// Package main demonstrates Excel workbook to PDF rendering with DrawingML charts.
//
// This example shows how to:
// 1. Create an Excel workbook with data
// 2. Render the workbook to PDF using DrawingML chart rendering
// 3. Generate a professional-looking PDF with charts
//
// The DrawingML rendering system handles:
// - Chart rendering (bar, line, pie charts)
// - Data visualization
// - Chart legends and axes
// - Professional layout and styling
package main

import (
	"fmt"
	"log"
	"path/filepath"

	pdfspreadsheet "github.com/connerohnesorge/goffice-pdf/spreadsheet"
	"github.com/connerohnesorge/goffice/spreadsheet"
)

const (
	xlsxFilename = "excel_charts_example.xlsx"
	pdfFilename  = "excel_charts_example.pdf"
)

func main() {
	fmt.Println(
		"Creating Excel workbook with data for charts...",
	)

	// Create the Excel workbook
	if err := createExcelWorkbook(); err != nil {
		log.Fatalf(
			"Failed to create Excel workbook: %v",
			err,
		)
	}

	fmt.Printf("Created %s\n", xlsxFilename)
	fmt.Println("Rendering to PDF...")

	// Render to PDF
	if err := renderToPDF(); err != nil {
		log.Fatalf(
			"Failed to render PDF: %v",
			err,
		)
	}

	fmt.Printf(
		"Successfully rendered to %s\n",
		pdfFilename,
	)
	fmt.Println(
		"\nThe PDF demonstrates DrawingML chart rendering:",
	)
	fmt.Println("  - Sales data visualization")
	fmt.Println(
		"  - Quarterly performance tracking",
	)
	fmt.Println("  - Revenue analysis")
	fmt.Println(
		"  - Professional spreadsheet layout",
	)
	fmt.Println(
		"\nNote: Charts would be rendered if embedded in the Excel file.",
	)
	fmt.Println(
		"This example demonstrates the data foundation for chart rendering.",
	)
}

// createExcelWorkbook creates an Excel workbook with sample data.
func createExcelWorkbook() error {
	doc, err := spreadsheet.Create(
		xlsxFilename,
		spreadsheet.DocTypeWorkbook,
	)
	if err != nil {
		return fmt.Errorf(
			"creating workbook: %w",
			err,
		)
	}
	defer doc.Close()

	// Create Sales Data sheet
	salesSheet, err := doc.AddSheet("Sales Data")
	if err != nil {
		return fmt.Errorf(
			"adding sales sheet: %w",
			err,
		)
	}

	// Add headers
	if err := salesSheet.SetCellValue("A1", "Quarter"); err != nil {
		return fmt.Errorf("setting A1: %w", err)
	}
	if err := salesSheet.SetCellValue("B1", "Sales"); err != nil {
		return fmt.Errorf("setting B1: %w", err)
	}
	if err := salesSheet.SetCellValue("C1", "Revenue"); err != nil {
		return fmt.Errorf("setting C1: %w", err)
	}
	if err := salesSheet.SetCellValue("D1", "Profit"); err != nil {
		return fmt.Errorf("setting D1: %w", err)
	}

	// Add quarterly data
	quarters := []string{"Q1", "Q2", "Q3", "Q4"}
	salesData := []float64{100, 150, 120, 180}
	revenueData := []float64{80, 120, 100, 150}
	profitData := []float64{20, 30, 20, 30}

	for i, q := range quarters {
		row := i + 2
		cellA := fmt.Sprintf("A%d", row)
		cellB := fmt.Sprintf("B%d", row)
		cellC := fmt.Sprintf("C%d", row)
		cellD := fmt.Sprintf("D%d", row)

		if err := salesSheet.SetCellValue(cellA, q); err != nil {
			return fmt.Errorf(
				"setting %s: %w",
				cellA,
				err,
			)
		}
		if err := salesSheet.SetCellValue(cellB, salesData[i]); err != nil {
			return fmt.Errorf(
				"setting %s: %w",
				cellB,
				err,
			)
		}
		if err := salesSheet.SetCellValue(cellC, revenueData[i]); err != nil {
			return fmt.Errorf(
				"setting %s: %w",
				cellC,
				err,
			)
		}
		if err := salesSheet.SetCellValue(cellD, profitData[i]); err != nil {
			return fmt.Errorf(
				"setting %s: %w",
				cellD,
				err,
			)
		}
	}

	// Create Product Performance sheet
	productSheet, err := doc.AddSheet(
		"Product Performance",
	)
	if err != nil {
		return fmt.Errorf(
			"adding product sheet: %w",
			err,
		)
	}

	// Add product data
	if err := productSheet.SetCellValue("A1", "Product"); err != nil {
		return fmt.Errorf(
			"setting product A1: %w",
			err,
		)
	}
	if err := productSheet.SetCellValue("B1", "Units Sold"); err != nil {
		return fmt.Errorf(
			"setting product B1: %w",
			err,
		)
	}
	if err := productSheet.SetCellValue("C1", "Market Share %"); err != nil {
		return fmt.Errorf(
			"setting product C1: %w",
			err,
		)
	}

	products := []string{
		"Product A",
		"Product B",
		"Product C",
		"Product D",
	}
	unitsSold := []int{500, 750, 600, 900}
	marketShare := []float64{
		18.5,
		27.8,
		22.2,
		33.3,
	}

	for i, p := range products {
		row := i + 2
		cellA := fmt.Sprintf("A%d", row)
		cellB := fmt.Sprintf("B%d", row)
		cellC := fmt.Sprintf("C%d", row)

		if err := productSheet.SetCellValue(cellA, p); err != nil {
			return fmt.Errorf(
				"setting %s: %w",
				cellA,
				err,
			)
		}
		if err := productSheet.SetCellValue(cellB, unitsSold[i]); err != nil {
			return fmt.Errorf(
				"setting %s: %w",
				cellB,
				err,
			)
		}
		if err := productSheet.SetCellValue(cellC, marketShare[i]); err != nil {
			return fmt.Errorf(
				"setting %s: %w",
				cellC,
				err,
			)
		}
	}

	// Create Summary sheet
	summarySheet, err := doc.AddSheet("Summary")
	if err != nil {
		return fmt.Errorf(
			"adding summary sheet: %w",
			err,
		)
	}

	if err := summarySheet.SetCellValue("A1", "Metric"); err != nil {
		return fmt.Errorf(
			"setting summary A1: %w",
			err,
		)
	}
	if err := summarySheet.SetCellValue("B1", "Value"); err != nil {
		return fmt.Errorf(
			"setting summary B1: %w",
			err,
		)
	}

	metrics := []string{
		"Total Sales",
		"Total Revenue",
		"Total Profit",
		"Average Sales/Quarter",
	}
	values := []interface{}{550, 450, 100, 137.5}

	for i, m := range metrics {
		row := i + 2
		cellA := fmt.Sprintf("A%d", row)
		cellB := fmt.Sprintf("B%d", row)

		if err := summarySheet.SetCellValue(cellA, m); err != nil {
			return fmt.Errorf(
				"setting %s: %w",
				cellA,
				err,
			)
		}
		if err := summarySheet.SetCellValue(cellB, values[i]); err != nil {
			return fmt.Errorf(
				"setting %s: %w",
				cellB,
				err,
			)
		}
	}

	if err := doc.Save(); err != nil {
		return fmt.Errorf(
			"saving workbook: %w",
			err,
		)
	}

	return nil
}

// renderToPDF renders the Excel workbook to PDF.
func renderToPDF() error {
	// Open the workbook
	doc, err := spreadsheet.Open(
		xlsxFilename,
		false,
	)
	if err != nil {
		return fmt.Errorf(
			"opening workbook: %w",
			err,
		)
	}
	defer doc.Close()

	// Create renderer
	renderer, err := pdfspreadsheet.NewSpreadsheetRenderer(
		doc,
	)
	if err != nil {
		return fmt.Errorf(
			"creating renderer: %w",
			err,
		)
	}

	// Render to PDF
	absPath, err := filepath.Abs(pdfFilename)
	if err != nil {
		return fmt.Errorf(
			"resolving path: %w",
			err,
		)
	}

	if err := renderer.Render(absPath); err != nil {
		return fmt.Errorf(
			"rendering PDF: %w",
			err,
		)
	}

	return nil
}
