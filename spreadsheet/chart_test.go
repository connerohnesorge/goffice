package spreadsheet

import (
	"fmt"
	"testing"

	"github.com/connerohnesorge/goffice/drawingml"
)

func TestChartAxis_SetLogScale(t *testing.T) {
	doc := createTestWorkbook(t)
	defer func() { _ = doc.Close() }()

	sheet, _ := doc.AddSheet("TestData")
	addSampleData(sheet)

	chart, err := sheet.AddChart(ChartTypeBar, "A1:B5", MustParseCellRef("B6"))
	if err != nil {
		t.Fatalf("Failed to create chart: %v", err)
	}

	// Test SetLogScale method
	chart.ValueAxis().SetLogScale(10.0)

	// Verify the scaling was set (by checking it doesn't panic)
	if chart.ValueAxis() == nil {
		t.Error("ValueAxis should not be nil")
	}
}

func TestChartAxis_SetOrientation(t *testing.T) {
	doc := createTestWorkbook(t)
	defer func() { _ = doc.Close() }()

	sheet, _ := doc.AddSheet("TestData")
	addSampleData(sheet)

	chart, err := sheet.AddChart(ChartTypeLine, "A1:B5", MustParseCellRef("B6"))
	if err != nil {
		t.Fatalf("Failed to create chart: %v", err)
	}

	// Test SetOrientation method
	chart.ValueAxis().SetOrientation(drawingml.OrientationMaxMin)

	if chart.ValueAxis() == nil {
		t.Error("ValueAxis should not be nil")
	}
}

func TestChartAxis_SetReverseOrder(t *testing.T) {
	doc := createTestWorkbook(t)
	defer func() { _ = doc.Close() }()

	sheet, _ := doc.AddSheet("TestData")
	addSampleData(sheet)

	chart, err := sheet.AddChart(ChartTypeBar, "A1:B5", MustParseCellRef("B6"))
	if err != nil {
		t.Fatalf("Failed to create chart: %v", err)
	}

	// Test SetReverseOrder method
	chart.ValueAxis().SetReverseOrder(true)

	if chart.ValueAxis() == nil {
		t.Error("ValueAxis should not be nil")
	}
}

func TestChartAxis_SetAxisPosition(t *testing.T) {
	doc := createTestWorkbook(t)
	defer func() { _ = doc.Close() }()

	sheet, _ := doc.AddSheet("TestData")
	addSampleData(sheet)

	chart, err := sheet.AddChart(ChartTypePie, "A1:B5", MustParseCellRef("B6"))
	if err != nil {
		t.Fatalf("Failed to create chart: %v", err)
	}

	// Test SetAxisPosition method
	chart.CategoryAxis().SetAxisPosition(drawingml.AxisPositionBottom)

	if chart.CategoryAxis() == nil {
		t.Error("CategoryAxis should not be nil")
	}
}

func TestChart_CreationWithAxisFormatting(t *testing.T) {
	testCases := []struct {
		chartType ChartType
		name      string
	}{
		{ChartTypeBar, "Bar Chart"},
		{ChartTypeLine, "Line Chart"},
		{ChartTypePie, "Pie Chart"},
		{ChartTypeScatter, "Scatter Chart"},
		{ChartTypeArea, "Area Chart"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			doc := createTestWorkbook(t)
			defer func() { _ = doc.Close() }()

			sheet, _ := doc.AddSheet("TestData")
			addSampleData(sheet)

			chart, err := sheet.AddChart(tc.chartType, "A1:B5", MustParseCellRef("B6"))
			if err != nil {
				t.Fatalf("Failed to create %s: %v", tc.name, err)
			}

			if chart == nil {
				t.Fatalf("%s should not be nil", tc.name)
			}

			// Test that axis methods don't panic when called
			if chart.ValueAxis() != nil {
				chart.ValueAxis().SetTitle("Test Value Axis")
				chart.ValueAxis().SetLogScale(10.0)
				chart.ValueAxis().SetOrientation(drawingml.OrientationMaxMin)
				chart.ValueAxis().SetReverseOrder(false)
				chart.ValueAxis().SetAxisPosition(drawingml.AxisPositionLeft)
			}

			if chart.CategoryAxis() != nil {
				chart.CategoryAxis().SetTitle("Test Category Axis")
				chart.CategoryAxis().SetAxisPosition(drawingml.AxisPositionBottom)
			}

			// If we get here without panic, the test passes
		})
	}
}

func TestChart_BasicAxisFormatting(t *testing.T) {
	doc := createTestWorkbook(t)
	defer func() { _ = doc.Close() }()

	sheet, _ := doc.AddSheet("TestData")
	addSampleData(sheet)

	chart, err := sheet.AddChart(ChartTypeBar, "A1:B5", MustParseCellRef("B6"))
	if err != nil {
		t.Fatalf("Failed to create chart: %v", err)
	}

	// Test that axis methods work without panicking
	if chart.ValueAxis() != nil {
		chart.ValueAxis().SetTitle("Test Value Axis")
		chart.ValueAxis().SetLogScale(10.0)
		chart.ValueAxis().SetOrientation(drawingml.OrientationMaxMin)
		chart.ValueAxis().SetReverseOrder(false)
		chart.ValueAxis().SetAxisPosition(drawingml.AxisPositionLeft)
	}

	if chart.CategoryAxis() != nil {
		chart.CategoryAxis().SetTitle("Test Category Axis")
		chart.CategoryAxis().SetAxisPosition(drawingml.AxisPositionBottom)
	}

	// If we get here without panic, the test passes
}

// Helper functions for testing

func createTestWorkbook(t *testing.T) *Document {
	doc, err := Create("test.xlsx", DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Failed to create workbook: %v", err)
	}

	return doc
}

func addSampleData(sheet *Sheet) {
	// Add headers
	_ = sheet.SetCellValue("A1", "Category")
	_ = sheet.SetCellValue("B1", "Value")

	// Add data rows
	data := []float64{25, 30, 45, 20, 35}
	for i, value := range data {
		row := i + 2
		_ = sheet.SetCellValue(RowCol(row, 1), fmt.Sprintf("Item %d", i+1))
		_ = sheet.SetCellValue(RowCol(row, 2), value)
	}
}

func RowCol(row, col int) string {
	return fmt.Sprintf("%s%d", string(rune('A'+col-1)), row)
}
