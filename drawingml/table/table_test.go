package table

import (
	"testing"

	"github.com/connerohnesorge/goffice/drawingml"
)

func TestNewTable(t *testing.T) {
	t.Run("Creates table with correct structure", func(t *testing.T) {
		tbl := NewTable(3, 4)

		if tbl == nil {
			t.Fatal("NewTable returned nil")
		}

		// Check row count
		if got := tbl.RowCount(); got != 3 {
			t.Errorf("RowCount() = %d, want 3", got)
		}

		// Check column count
		if got := tbl.ColumnCount(); got != 4 {
			t.Errorf("ColumnCount() = %d, want 4", got)
		}

		// Verify properties exist
		if props := tbl.Properties(); props == nil {
			t.Error("Properties() returned nil")
		}

		// Verify grid exists
		if grid := tbl.Grid(); grid == nil {
			t.Error("Grid() returned nil")
		}
	})

	t.Run("Each row has correct number of cells", func(t *testing.T) {
		tbl := NewTable(2, 3)
		rows := tbl.Rows()

		for i, row := range rows {
			if got := row.CellCount(); got != 3 {
				t.Errorf("Row %d has %d cells, want 3", i, got)
			}
		}
	})

	t.Run("Each cell has text body", func(t *testing.T) {
		tbl := NewTable(2, 2)

		for i, row := range tbl.Rows() {
			for j, cell := range row.Cells() {
				if tb := cell.TextBody(); tb == nil {
					t.Errorf("Cell[%d][%d].TextBody() is nil", i, j)
				}
			}
		}
	})
}

func TestTableProperties(t *testing.T) {
	t.Run("FirstRow", func(t *testing.T) {
		props := NewTableProperties()

		// Default is false
		if props.FirstRow() {
			t.Error("FirstRow() should be false by default")
		}

		// Set to true
		props.SetFirstRow(true)
		if !props.FirstRow() {
			t.Error("FirstRow() should be true after SetFirstRow(true)")
		}

		// Set back to false
		props.SetFirstRow(false)
		if props.FirstRow() {
			t.Error("FirstRow() should be false after SetFirstRow(false)")
		}
	})

	t.Run("BandRow", func(t *testing.T) {
		props := NewTableProperties()
		props.SetBandRow(true)

		if !props.BandRow() {
			t.Error("BandRow() should be true")
		}
	})

	t.Run("All properties", func(t *testing.T) {
		props := NewTableProperties()
		props.SetFirstRow(true)
		props.SetLastRow(true)
		props.SetBandRow(true)
		props.SetBandCol(true)
		props.SetFirstCol(true)
		props.SetLastCol(true)

		if !props.FirstRow() {
			t.Error("FirstRow should be true")
		}
		if !props.LastRow() {
			t.Error("LastRow should be true")
		}
		if !props.BandRow() {
			t.Error("BandRow should be true")
		}
		if !props.BandCol() {
			t.Error("BandCol should be true")
		}
		if !props.FirstCol() {
			t.Error("FirstCol should be true")
		}
		if !props.LastCol() {
			t.Error("LastCol should be true")
		}
	})
}

func TestTableRowOperations(t *testing.T) {
	t.Run("AppendRow", func(t *testing.T) {
		tbl := NewTable(2, 3)
		originalCount := tbl.RowCount()

		newRow := tbl.AppendRow()
		if newRow == nil {
			t.Fatal("AppendRow() returned nil")
		}

		if got := tbl.RowCount(); got != originalCount+1 {
			t.Errorf("RowCount() = %d, want %d", got, originalCount+1)
		}

		// New row should have correct number of columns
		if got := newRow.CellCount(); got != 3 {
			t.Errorf("New row has %d cells, want 3", got)
		}
	})

	t.Run("InsertRow", func(t *testing.T) {
		tbl := NewTable(3, 2)

		// Insert at beginning
		if err := tbl.InsertRow(0); err != nil {
			t.Fatalf("InsertRow(0) failed: %v", err)
		}
		if got := tbl.RowCount(); got != 4 {
			t.Errorf("After InsertRow(0), RowCount() = %d, want 4", got)
		}

		// Insert in middle
		if err := tbl.InsertRow(2); err != nil {
			t.Fatalf("InsertRow(2) failed: %v", err)
		}
		if got := tbl.RowCount(); got != 5 {
			t.Errorf("After InsertRow(2), RowCount() = %d, want 5", got)
		}

		// Insert at end
		if err := tbl.InsertRow(5); err != nil {
			t.Fatalf("InsertRow(5) failed: %v", err)
		}
		if got := tbl.RowCount(); got != 6 {
			t.Errorf("After InsertRow(5), RowCount() = %d, want 6", got)
		}
	})

	t.Run("DeleteRow", func(t *testing.T) {
		tbl := NewTable(4, 2)

		// Delete from middle
		if err := tbl.DeleteRow(1); err != nil {
			t.Fatalf("DeleteRow(1) failed: %v", err)
		}
		if got := tbl.RowCount(); got != 3 {
			t.Errorf("After DeleteRow(1), RowCount() = %d, want 3", got)
		}

		// Delete first row
		if err := tbl.DeleteRow(0); err != nil {
			t.Fatalf("DeleteRow(0) failed: %v", err)
		}
		if got := tbl.RowCount(); got != 2 {
			t.Errorf("After DeleteRow(0), RowCount() = %d, want 2", got)
		}

		// Out of range should error
		if err := tbl.DeleteRow(10); err == nil {
			t.Error("DeleteRow(10) should return error for out of range")
		}
	})
}

func TestTableCellOperations(t *testing.T) {
	t.Run("GetCell", func(t *testing.T) {
		tbl := NewTable(3, 3)

		// Get valid cell
		cell := tbl.GetCell(1, 2)
		if cell == nil {
			t.Error("GetCell(1, 2) returned nil")
		}

		// Out of range row
		if cell := tbl.GetCell(5, 0); cell != nil {
			t.Error("GetCell(5, 0) should return nil for out of range row")
		}

		// Out of range column
		if cell := tbl.GetCell(0, 5); cell != nil {
			t.Error("GetCell(0, 5) should return nil for out of range column")
		}
	})

	t.Run("SetText and GetText", func(t *testing.T) {
		tbl := NewTable(2, 2)
		cell := tbl.GetCell(0, 0)

		text := "Hello, World!"
		cell.SetText(text)

		if got := cell.GetText(); got != text {
			t.Errorf("GetText() = %q, want %q", got, text)
		}
	})

	t.Run("Cell properties margins", func(t *testing.T) {
		cell := NewTableCell()
		props := cell.Properties()

		// Set margins (in EMUs)
		leftMargin := int64(100000)
		rightMargin := int64(200000)
		topMargin := int64(150000)
		bottomMargin := int64(250000)

		props.SetMarginLeft(leftMargin)
		props.SetMarginRight(rightMargin)
		props.SetMarginTop(topMargin)
		props.SetMarginBottom(bottomMargin)

		if got := props.MarginLeft(); got != leftMargin {
			t.Errorf("MarginLeft() = %d, want %d", got, leftMargin)
		}
		if got := props.MarginRight(); got != rightMargin {
			t.Errorf("MarginRight() = %d, want %d", got, rightMargin)
		}
		if got := props.MarginTop(); got != topMargin {
			t.Errorf("MarginTop() = %d, want %d", got, topMargin)
		}
		if got := props.MarginBottom(); got != bottomMargin {
			t.Errorf("MarginBottom() = %d, want %d", got, bottomMargin)
		}
	})
}

func TestTableGrid(t *testing.T) {
	t.Run("Column count", func(t *testing.T) {
		grid := NewTableGrid(5)

		if got := grid.ColumnCount(); got != 5 {
			t.Errorf("ColumnCount() = %d, want 5", got)
		}
	})

	t.Run("Set column width", func(t *testing.T) {
		grid := NewTableGrid(3)
		newWidth := int64(2 * drawingml.EMUsPerInch) // 2 inches

		if err := grid.SetColumnWidth(1, newWidth); err != nil {
			t.Fatalf("SetColumnWidth(1, %d) failed: %v", newWidth, err)
		}

		col := grid.GetColumn(1)
		if col == nil {
			t.Fatal("GetColumn(1) returned nil")
		}

		if got := col.Width(); got != newWidth {
			t.Errorf("Column width = %d, want %d", got, newWidth)
		}
	})

	t.Run("Default column width", func(t *testing.T) {
		grid := NewTableGrid(1)
		col := grid.GetColumn(0)

		// Default width should be 1 inch
		expectedWidth := int64(drawingml.EMUsPerInch)
		if got := col.Width(); got != expectedWidth {
			t.Errorf("Default column width = %d, want %d", got, expectedWidth)
		}
	})
}

func TestTableRow(t *testing.T) {
	t.Run("Default height", func(t *testing.T) {
		row := NewTableRow(3)

		// Default height should be 0.25 inches
		expectedHeight := int64(drawingml.EMUsPerInch / 4)
		if got := row.Height(); got != expectedHeight {
			t.Errorf("Default row height = %d, want %d", got, expectedHeight)
		}
	})

	t.Run("Set height", func(t *testing.T) {
		row := NewTableRow(2)
		newHeight := int64(500000) // Custom height

		row.SetHeight(newHeight)
		if got := row.Height(); got != newHeight {
			t.Errorf("Row height = %d, want %d", got, newHeight)
		}
	})

	t.Run("Append cell", func(t *testing.T) {
		row := NewTableRow(2)
		originalCount := row.CellCount()

		newCell := row.AppendCell()
		if newCell == nil {
			t.Fatal("AppendCell() returned nil")
		}

		if got := row.CellCount(); got != originalCount+1 {
			t.Errorf("CellCount() = %d, want %d", got, originalCount+1)
		}
	})
}
