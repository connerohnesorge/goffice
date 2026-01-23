package elements

import (
	"testing"
)

func TestTableDeleteRow(t *testing.T) {
	table := NewTable(3, 2)

	table.SetCellText(0, 0, "Row0")
	table.SetCellText(1, 0, "Row1")
	table.SetCellText(2, 0, "Row2")

	// Delete middle row
	if !table.DeleteRow(1) {
		t.Fatal("DeleteRow returned false")
	}

	if table.RowCount() != 2 {
		t.Errorf(
			"Expected 2 rows, got %d",
			table.RowCount(),
		)
	}

	// Verify Row2 is now at index 1
	cell := table.GetCell(1, 0)
	if cell.InnerText() != "Row2" {
		t.Errorf(
			"Expected 'Row2', got '%s'",
			cell.InnerText(),
		)
	}
}

