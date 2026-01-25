package spreadsheet

import (
	"testing"
)

func TestWorkbookBuilder(t *testing.T) {
	wb := NewWorkbookBuilder()
	doc, err := wb.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	if doc == nil {
		t.Fatal("Build returned nil document")
	}
}

func TestSheetBuilder(t *testing.T) {
	wb := NewWorkbookBuilder()
	wb.AddSheet("Sheet1").
		AddRow().
		Cell(1).SetString("Hello").
		Row().
		Cell(2).SetNumber(42)

	doc, err := wb.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	sheet, ok := doc.SheetByName("Sheet1")
	if !ok {
		t.Fatal("Sheet1 not found")
	}

	cell1 := sheet.Cell("A1")
	if cell1.GetString() != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", cell1.GetString())
	}

	cell2 := sheet.Cell("B1")
	if cell2.GetNumber() != 42 {
		t.Errorf("Expected 42, got %f", cell2.GetNumber())
	}
}

func TestRangeBuilder(t *testing.T) {
	wb := NewWorkbookBuilder()
	wb.AddSheet("Sheet1").
		Range("A1:B2").
		SetValue("Test")

	doc, err := wb.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	sheet, _ := doc.SheetByName("Sheet1")
	if sheet.Cell("A1").GetString() != "Test" {
		t.Error("A1 should be 'Test'")
	}
	if sheet.Cell("B2").GetString() != "Test" {
		t.Error("B2 should be 'Test'")
	}
}
