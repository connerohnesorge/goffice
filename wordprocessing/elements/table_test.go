package elements

import (
	"strings"
	"testing"
)

func TestNewTable(t *testing.T) {
	table := NewTable(3, 4)

	if table == nil {
		t.Fatal("NewTable returned nil")
	}

	// Check row count
	if table.RowCount() != 3 {
		t.Errorf("Expected 3 rows, got %d", table.RowCount())
	}

	// Check each row has 4 cells
	for i := 0; i < 3; i++ {
		row := table.GetRow(i)
		if row == nil {
			t.Fatalf("Row %d is nil", i)
		}
		if row.CellCount() != 4 {
			t.Errorf("Row %d: expected 4 cells, got %d", i, row.CellCount())
		}
	}
}

func TestTableGetCell(t *testing.T) {
	table := NewTable(2, 3)

	cell := table.GetCell(0, 0)
	if cell == nil {
		t.Fatal("GetCell(0, 0) returned nil")
	}

	cell = table.GetCell(1, 2)
	if cell == nil {
		t.Fatal("GetCell(1, 2) returned nil")
	}

	// Out of bounds
	cell = table.GetCell(5, 5)
	if cell != nil {
		t.Error("GetCell(5, 5) should return nil for out of bounds")
	}
}

func TestTableSetCellText(t *testing.T) {
	table := NewTable(2, 2)

	table.SetCellText(0, 0, "Hello")
	table.SetCellText(1, 1, "World")

	cell00 := table.GetCell(0, 0)
	if cell00.InnerText() != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", cell00.InnerText())
	}

	cell11 := table.GetCell(1, 1)
	if cell11.InnerText() != "World" {
		t.Errorf("Expected 'World', got '%s'", cell11.InnerText())
	}
}

func TestTableAppendRow(t *testing.T) {
	table := NewTable(2, 3)

	initialCount := table.RowCount()
	table.AppendRow(3)

	if table.RowCount() != initialCount+1 {
		t.Errorf("Expected %d rows, got %d", initialCount+1, table.RowCount())
	}
}

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
		t.Errorf("Expected 2 rows, got %d", table.RowCount())
	}

	// Verify Row2 is now at index 1
	cell := table.GetCell(1, 0)
	if cell.InnerText() != "Row2" {
		t.Errorf("Expected 'Row2', got '%s'", cell.InnerText())
	}
}

func TestTableProperties(t *testing.T) {
	table := NewTable(2, 2)

	props := table.GetOrCreateTableProperties()
	if props == nil {
		t.Fatal("GetOrCreateTableProperties returned nil")
	}

	// Set style
	props.SetTableStyle("TableGrid")
	if props.TableStyle() != "TableGrid" {
		t.Errorf("Expected style 'TableGrid', got '%s'", props.TableStyle())
	}

	// Set width
	props.SetTableWidth(5000, TableWidthTypePct)
	tw := props.TableWidth()
	if tw == nil {
		t.Fatal("TableWidth is nil")
	}
	if tw.Width() != 5000 {
		t.Errorf("Expected width 5000, got %d", tw.Width())
	}
	if tw.Type() != TableWidthTypePct {
		t.Errorf("Expected type pct, got %s", tw.Type())
	}
}

func TestTableBorders(t *testing.T) {
	table := NewTable(2, 2)
	props := table.GetOrCreateTableProperties()
	borders := props.GetOrCreateTableBorders()

	borders.SetAllBorders(BorderSingle, 4, "000000")

	xml := borders.OuterXml()
	if !strings.Contains(xml, "top") {
		t.Error("Expected top border in XML")
	}
	if !strings.Contains(xml, "bottom") {
		t.Error("Expected bottom border in XML")
	}
	if !strings.Contains(xml, "left") {
		t.Error("Expected left border in XML")
	}
	if !strings.Contains(xml, "right") {
		t.Error("Expected right border in XML")
	}
	if !strings.Contains(xml, "insideH") {
		t.Error("Expected insideH border in XML")
	}
	if !strings.Contains(xml, "insideV") {
		t.Error("Expected insideV border in XML")
	}
}

func TestTableLook(t *testing.T) {
	table := NewTable(2, 2)
	props := table.GetOrCreateTableProperties()

	props.SetTableLook(true, false, true, false, true, false)

	tl := props.TableLook()
	if tl == nil {
		t.Fatal("TableLook is nil")
	}

	// Verify attributes are set
	xml := tl.OuterXml()
	if !strings.Contains(xml, "firstRow") {
		t.Error("Expected firstRow in XML")
	}
}

func TestTableRowProperties(t *testing.T) {
	table := NewTable(2, 2)
	row := table.GetRow(0)

	props := row.GetOrCreateTableRowProperties()
	if props == nil {
		t.Fatal("GetOrCreateTableRowProperties returned nil")
	}

	// Set height
	props.SetHeight(500, HeightRuleExact)
	xml := props.OuterXml()
	if !strings.Contains(xml, "trHeight") {
		t.Error("Expected trHeight in XML")
	}

	// Set header
	props.SetHeader(true)
	xml = props.OuterXml()
	if !strings.Contains(xml, "tblHeader") {
		t.Error("Expected tblHeader in XML")
	}
}

func TestTableCellProperties(t *testing.T) {
	table := NewTable(2, 2)
	cell := table.GetCell(0, 0)

	props := cell.GetOrCreateTableCellProperties()
	if props == nil {
		t.Fatal("GetOrCreateTableCellProperties returned nil")
	}

	// Set width
	props.SetWidth(2000, TableWidthTypeDxa)
	tcw := props.TableCellWidth()
	if tcw == nil {
		t.Fatal("TableCellWidth is nil")
	}
	if tcw.Width() != 2000 {
		t.Errorf("Expected width 2000, got %d", tcw.Width())
	}

	// Set shading
	props.SetShading("FF0000")
	shd := props.Shading()
	if shd == nil {
		t.Fatal("Shading is nil")
	}
	if shd.Fill() != "FF0000" {
		t.Errorf("Expected fill FF0000, got %s", shd.Fill())
	}

	// Set grid span
	props.SetGridSpan(2)
	if props.GridSpan() != 2 {
		t.Errorf("Expected gridSpan 2, got %d", props.GridSpan())
	}
}

func TestVerticalMerge(t *testing.T) {
	table := NewTable(3, 2)

	// Start merge at row 0
	cell0 := table.GetCell(0, 0)
	cell0.SetVerticalMerge(VerticalMergeRestart)

	// Continue merge at row 1
	cell1 := table.GetCell(1, 0)
	cell1.SetVerticalMerge(VerticalMergeContinue)

	// Verify
	props0 := cell0.TableCellProperties()
	vm0 := props0.VerticalMerge()
	if vm0 == nil {
		t.Fatal("VerticalMerge is nil for cell 0")
	}
	if vm0.Type() != VerticalMergeRestart {
		t.Errorf("Expected restart, got %s", vm0.Type())
	}

	props1 := cell1.TableCellProperties()
	vm1 := props1.VerticalMerge()
	if vm1 == nil {
		t.Fatal("VerticalMerge is nil for cell 1")
	}
	if vm1.Type() != VerticalMergeContinue {
		t.Errorf("Expected continue, got %s", vm1.Type())
	}
}

func TestTableCellBorders(t *testing.T) {
	cell := NewTableCell()
	props := cell.GetOrCreateTableCellProperties()
	borders := props.GetOrCreateTableCellBorders()

	borders.SetAllBorders(BorderDouble, 8, "FF0000")

	xml := borders.OuterXml()
	if !strings.Contains(xml, "top") {
		t.Error("Expected top border")
	}
	if !strings.Contains(xml, "double") {
		t.Error("Expected double border style")
	}
	if !strings.Contains(xml, "FF0000") {
		t.Error("Expected color FF0000")
	}
}

func TestTableFluentAPI(t *testing.T) {
	table := NewTable(2, 2)

	// Test chained methods
	table.SetStyle("TableGrid").
		SetWidth(5000, TableWidthTypePct).
		SetColumnWidth(0, 2000)

	if table.TableProperties().TableStyle() != "TableGrid" {
		t.Error("Style not set correctly")
	}
}

func TestTableCellParagraphs(t *testing.T) {
	cell := NewTableCell()

	// Cell should have one paragraph by default
	count := 0
	for range cell.Paragraphs() {
		count++
	}
	if count != 1 {
		t.Errorf("Expected 1 paragraph, got %d", count)
	}

	// Append another paragraph
	cell.AppendParagraph("Second paragraph")
	count = 0
	for range cell.Paragraphs() {
		count++
	}
	if count != 2 {
		t.Errorf("Expected 2 paragraphs, got %d", count)
	}
}

func TestTableGrid(t *testing.T) {
	table := NewTable(2, 3)

	grid := table.TableGrid()
	if grid == nil {
		t.Fatal("TableGrid is nil")
	}

	// Count grid columns
	count := 0
	for range grid.GridColumns() {
		count++
	}
	if count != 3 {
		t.Errorf("Expected 3 grid columns, got %d", count)
	}
}
