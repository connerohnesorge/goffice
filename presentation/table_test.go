package presentation

import (
	"testing"

	"github.com/connerohnesorge/goffice/presentation/elements"
)

func TestNewTable(t *testing.T) {
	t.Run("Creates table with correct structure", func(t *testing.T) {
		slide := elements.NewSlide()
		table := NewTable(slide, 3, 4)

		if table == nil {
			t.Fatal("NewTable returned nil")
		}

		// Check row count
		if got := table.RowCount(); got != 3 {
			t.Errorf("RowCount() = %d, want 3", got)
		}

		// Check column count
		if got := table.ColumnCount(); got != 4 {
			t.Errorf("ColumnCount() = %d, want 4", got)
		}

		// Verify graphic frame exists
		if table.GraphicFrame() == nil {
			t.Error("GraphicFrame() returned nil")
		}

		// Verify underlying table exists
		if table.DrawingMLTable() == nil {
			t.Error("DrawingMLTable() returned nil")
		}
	})
}

func TestTableRowOperations(t *testing.T) {
	slide := elements.NewSlide()
	table := NewTable(slide, 2, 3)

	t.Run("AppendRow", func(t *testing.T) {
		originalCount := table.RowCount()
		table.AppendRow()

		if got := table.RowCount(); got != originalCount+1 {
			t.Errorf("After AppendRow(), RowCount() = %d, want %d", got, originalCount+1)
		}
	})

	t.Run("InsertRow", func(t *testing.T) {
		originalCount := table.RowCount()
		if err := table.InsertRow(1); err != nil {
			t.Fatalf("InsertRow(1) failed: %v", err)
		}

		if got := table.RowCount(); got != originalCount+1 {
			t.Errorf("After InsertRow(1), RowCount() = %d, want %d", got, originalCount+1)
		}
	})

	t.Run("DeleteRow", func(t *testing.T) {
		originalCount := table.RowCount()
		if err := table.DeleteRow(0); err != nil {
			t.Fatalf("DeleteRow(0) failed: %v", err)
		}

		if got := table.RowCount(); got != originalCount-1 {
			t.Errorf("After DeleteRow(0), RowCount() = %d, want %d", got, originalCount-1)
		}
	})
}

func TestTableCellOperations(t *testing.T) {
	slide := elements.NewSlide()
	table := NewTable(slide, 2, 2)

	t.Run("SetCellText and GetCellText", func(t *testing.T) {
		text := "Test Content"
		if err := table.SetCellText(0, 0, text); err != nil {
			t.Fatalf("SetCellText failed: %v", err)
		}

		got, err := table.GetCellText(0, 0)
		if err != nil {
			t.Fatalf("GetCellText failed: %v", err)
		}

		if got != text {
			t.Errorf("GetCellText() = %q, want %q", got, text)
		}
	})

	t.Run("GetCell", func(t *testing.T) {
		cell, err := table.GetCell(1, 1)
		if err != nil {
			t.Fatalf("GetCell(1, 1) failed: %v", err)
		}

		if cell == nil {
			t.Error("GetCell(1, 1) returned nil cell")
		}

		// Test setting text directly on cell
		cell.SetText("Direct text")
		if got := cell.GetText(); got != "Direct text" {
			t.Errorf("Cell.GetText() = %q, want %q", got, "Direct text")
		}
	})

	t.Run("Out of range cell", func(t *testing.T) {
		_, err := table.GetCell(10, 10)
		if err == nil {
			t.Error("GetCell(10, 10) should return error for out of range")
		}
	})
}

func TestTableColumnWidth(t *testing.T) {
	slide := elements.NewSlide()
	table := NewTable(slide, 2, 3)

	t.Run("SetColumnWidth", func(t *testing.T) {
		// Set column 1 to 100 points
		if err := table.SetColumnWidth(1, 100.0); err != nil {
			t.Fatalf("SetColumnWidth(1, 100.0) failed: %v", err)
		}

		// Verify it was set (we can check the underlying grid)
		grid := table.DrawingMLTable().Grid()
		if grid == nil {
			t.Fatal("Grid is nil")
		}

		col := grid.GetColumn(1)
		if col == nil {
			t.Fatal("Column 1 is nil")
		}

		// Width should be approximately 100 points in EMUs (100 * 12700)
		expectedEMU := int64(1270000)
		if got := col.Width(); got != expectedEMU {
			t.Errorf("Column width = %d EMU, want %d EMU", got, expectedEMU)
		}
	})

	t.Run("Out of range column", func(t *testing.T) {
		err := table.SetColumnWidth(10, 100.0)
		if err == nil {
			t.Error("SetColumnWidth(10, 100.0) should return error for out of range")
		}
	})
}

func TestTablePositionAndSize(t *testing.T) {
	slide := elements.NewSlide()
	table := NewTable(slide, 2, 2)

	t.Run("SetPosition", func(_ *testing.T) {
		table.SetPosition(100.0, 150.0)
		// Position is set in EMUs internally, we just verify no panic
	})

	t.Run("SetSize", func(_ *testing.T) {
		table.SetSize(400.0, 300.0)
		// Size is set in EMUs internally, we just verify no panic
	})
}

func TestTableProperties(t *testing.T) {
	slide := elements.NewSlide()
	table := NewTable(slide, 3, 3)

	t.Run("SetFirstRow", func(t *testing.T) {
		table.SetFirstRow(true)
		props := table.DrawingMLTable().Properties()
		if props == nil {
			t.Fatal("Properties is nil")
		}
		if !props.FirstRow() {
			t.Error("FirstRow should be true")
		}
	})

	t.Run("SetBandedRows", func(t *testing.T) {
		table.SetBandedRows(true)
		props := table.DrawingMLTable().Properties()
		if !props.BandRow() {
			t.Error("BandRow should be true")
		}
	})

	t.Run("SetBandedColumns", func(t *testing.T) {
		table.SetBandedColumns(true)
		props := table.DrawingMLTable().Properties()
		if !props.BandCol() {
			t.Error("BandCol should be true")
		}
	})
}

func TestTableCellTextBody(t *testing.T) {
	slide := elements.NewSlide()
	table := NewTable(slide, 2, 2)

	t.Run("Access TextBody for advanced formatting", func(t *testing.T) {
		cell, err := table.GetCell(0, 0)
		if err != nil {
			t.Fatalf("GetCell failed: %v", err)
		}

		tb := cell.TextBody()
		if tb == nil {
			t.Fatal("TextBody is nil")
		}

		// Add a paragraph with formatted text
		tb.ClearParagraphs()
		tb.AddParagraph("Formatted Text")

		// Verify we can access the paragraph
		paras := tb.Paragraphs()
		if len(paras) != 1 {
			t.Errorf("Expected 1 paragraph, got %d", len(paras))
		}
	})
}

func TestTableCellProperties(t *testing.T) {
	slide := elements.NewSlide()
	table := NewTable(slide, 2, 2)

	t.Run("Access cell properties", func(t *testing.T) {
		cell, err := table.GetCell(0, 0)
		if err != nil {
			t.Fatalf("GetCell failed: %v", err)
		}

		props := cell.Properties()
		if props == nil {
			t.Fatal("Properties is nil")
		}

		// Set margins
		props.SetMarginLeft(50000)
		props.SetMarginRight(60000)

		if got := props.MarginLeft(); got != 50000 {
			t.Errorf("MarginLeft = %d, want 50000", got)
		}
		if got := props.MarginRight(); got != 60000 {
			t.Errorf("MarginRight = %d, want 60000", got)
		}
	})
}
