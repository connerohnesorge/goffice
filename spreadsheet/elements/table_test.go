package elements

import (
	"testing"
)

func TestNewTable(t *testing.T) {
	table := NewTable()

	if table == nil {
		t.Fatal("NewTable returned nil")
	}

	if table.LocalName() != "table" {
		t.Errorf(
			"Expected LocalName 'table', got '%s'",
			table.LocalName(),
		)
	}

	if table.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"Expected NamespaceURI '%s', got '%s'",
			NamespaceSML,
			table.NamespaceURI(),
		)
	}
}

func TestNewTableWithDefaults(t *testing.T) {
	table := NewTableWithDefaults(
		1,
		"Table1",
		"Table1",
		"A1:D10",
	)

	if table.Id() != 1 {
		t.Errorf(
			"Expected Id 1, got %d",
			table.Id(),
		)
	}

	if table.Name() != "Table1" {
		t.Errorf(
			"Expected Name 'Table1', got '%s'",
			table.Name(),
		)
	}

	if table.DisplayName() != "Table1" {
		t.Errorf(
			"Expected DisplayName 'Table1', got '%s'",
			table.DisplayName(),
		)
	}

	if table.Ref() != "A1:D10" {
		t.Errorf(
			"Expected Ref 'A1:D10', got '%s'",
			table.Ref(),
		)
	}

	if table.TotalsRowShown() {
		t.Error(
			"Expected TotalsRowShown to be false",
		)
	}
}

func TestTableAttributes(t *testing.T) {
	table := NewTable()

	// Test Id
	table.SetId(42)
	if table.Id() != 42 {
		t.Errorf(
			"Expected Id 42, got %d",
			table.Id(),
		)
	}

	// Test Name
	table.SetName("SalesData")
	if table.Name() != "SalesData" {
		t.Errorf(
			"Expected Name 'SalesData', got '%s'",
			table.Name(),
		)
	}

	// Test DisplayName
	table.SetDisplayName("Sales Data Table")
	if table.DisplayName() != "Sales Data Table" {
		t.Errorf(
			"Expected DisplayName 'Sales Data Table', got '%s'",
			table.DisplayName(),
		)
	}

	// Test Ref
	table.SetRef("B2:F50")
	if table.Ref() != "B2:F50" {
		t.Errorf(
			"Expected Ref 'B2:F50', got '%s'",
			table.Ref(),
		)
	}

	// Test Comment
	table.SetComment("This is a test table")
	if table.Comment() != "This is a test table" {
		t.Errorf(
			"Expected Comment, got '%s'",
			table.Comment(),
		)
	}

	// Test HeaderRowCount
	table.SetHeaderRowCount(2)
	if table.HeaderRowCount() != 2 {
		t.Errorf(
			"Expected HeaderRowCount 2, got %d",
			table.HeaderRowCount(),
		)
	}

	// Test TotalsRowCount
	table.SetTotalsRowCount(1)
	if table.TotalsRowCount() != 1 {
		t.Errorf(
			"Expected TotalsRowCount 1, got %d",
			table.TotalsRowCount(),
		)
	}

	// Test TotalsRowShown
	table.SetTotalsRowShown(true)
	if !table.TotalsRowShown() {
		t.Error("Expected TotalsRowShown true")
	}
	table.SetTotalsRowShown(false)
	if table.TotalsRowShown() {
		t.Error("Expected TotalsRowShown false")
	}

	// Test Published
	table.SetPublished(true)
	if !table.Published() {
		t.Error("Expected Published true")
	}

	// Test InsertRow
	table.SetInsertRow(true)
	if !table.InsertRow() {
		t.Error("Expected InsertRow true")
	}
}

func TestTableDxfIds(t *testing.T) {
	table := NewTable()

	// Test HeaderRowDxfId
	table.SetHeaderRowDxfId(5)
	id, ok := table.HeaderRowDxfId()
	if !ok || id != 5 {
		t.Errorf(
			"Expected HeaderRowDxfId 5, got %d (ok=%v)",
			id,
			ok,
		)
	}
	table.ClearHeaderRowDxfId()
	_, ok = table.HeaderRowDxfId()
	if ok {
		t.Error(
			"Expected HeaderRowDxfId to be cleared",
		)
	}

	// Test DataDxfId
	table.SetDataDxfId(10)
	id, ok = table.DataDxfId()
	if !ok || id != 10 {
		t.Errorf(
			"Expected DataDxfId 10, got %d (ok=%v)",
			id,
			ok,
		)
	}

	// Test TotalsRowDxfId
	table.SetTotalsRowDxfId(15)
	id, ok = table.TotalsRowDxfId()
	if !ok || id != 15 {
		t.Errorf(
			"Expected TotalsRowDxfId 15, got %d (ok=%v)",
			id,
			ok,
		)
	}

	// Test HeaderRowBorderDxfId
	table.SetHeaderRowBorderDxfId(20)
	id, ok = table.HeaderRowBorderDxfId()
	if !ok || id != 20 {
		t.Errorf(
			"Expected HeaderRowBorderDxfId 20, got %d (ok=%v)",
			id,
			ok,
		)
	}

	// Test TableBorderDxfId
	table.SetTableBorderDxfId(25)
	id, ok = table.TableBorderDxfId()
	if !ok || id != 25 {
		t.Errorf(
			"Expected TableBorderDxfId 25, got %d (ok=%v)",
			id,
			ok,
		)
	}

	// Test TotalsRowBorderDxfId
	table.SetTotalsRowBorderDxfId(30)
	id, ok = table.TotalsRowBorderDxfId()
	if !ok || id != 30 {
		t.Errorf(
			"Expected TotalsRowBorderDxfId 30, got %d (ok=%v)",
			id,
			ok,
		)
	}
}

func TestTableCellStyles(t *testing.T) {
	table := NewTable()

	table.SetHeaderRowCellStyle("Header Style")
	if table.HeaderRowCellStyle() != "Header Style" {
		t.Errorf(
			"Expected HeaderRowCellStyle 'Header Style', got '%s'",
			table.HeaderRowCellStyle(),
		)
	}

	table.SetDataCellStyle("Data Style")
	if table.DataCellStyle() != "Data Style" {
		t.Errorf(
			"Expected DataCellStyle 'Data Style', got '%s'",
			table.DataCellStyle(),
		)
	}

	table.SetTotalsRowCellStyle("Totals Style")
	if table.TotalsRowCellStyle() != "Totals Style" {
		t.Errorf(
			"Expected TotalsRowCellStyle 'Totals Style', got '%s'",
			table.TotalsRowCellStyle(),
		)
	}
}

func TestTableAutoFilter(t *testing.T) {
	table := NewTable()
	table.SetRef("A1:D10")

	// Initially nil
	if table.AutoFilter() != nil {
		t.Error(
			"Expected AutoFilter to be nil initially",
		)
	}

	// Create auto filter
	af := table.GetOrCreateAutoFilter()
	if af == nil {
		t.Fatal(
			"GetOrCreateAutoFilter returned nil",
		)
	}

	af.SetRef("A1:D10")
	if af.Ref() != "A1:D10" {
		t.Errorf(
			"Expected AutoFilter ref 'A1:D10', got '%s'",
			af.Ref(),
		)
	}

	// Get existing auto filter
	af2 := table.GetOrCreateAutoFilter()
	if af2 != af {
		t.Error(
			"Expected GetOrCreateAutoFilter to return existing element",
		)
	}

	// Remove auto filter
	if !table.RemoveAutoFilter() {
		t.Error(
			"Expected RemoveAutoFilter to return true",
		)
	}
	if table.AutoFilter() != nil {
		t.Error(
			"Expected AutoFilter to be nil after removal",
		)
	}
}

func TestTableSortState(t *testing.T) {
	table := NewTable()

	// Initially nil
	if table.SortState() != nil {
		t.Error(
			"Expected SortState to be nil initially",
		)
	}

	// Create sort state
	ss := table.GetOrCreateSortState()
	if ss == nil {
		t.Fatal(
			"GetOrCreateSortState returned nil",
		)
	}

	ss.SetRef("A2:D10")
	if ss.Ref() != "A2:D10" {
		t.Errorf(
			"Expected SortState ref 'A2:D10', got '%s'",
			ss.Ref(),
		)
	}

	// Get existing sort state
	ss2 := table.GetOrCreateSortState()
	if ss2 != ss {
		t.Error(
			"Expected GetOrCreateSortState to return existing element",
		)
	}

	// Remove sort state
	if !table.RemoveSortState() {
		t.Error(
			"Expected RemoveSortState to return true",
		)
	}
	if table.SortState() != nil {
		t.Error(
			"Expected SortState to be nil after removal",
		)
	}
}

func TestTableColumns(t *testing.T) {
	table := NewTable()

	// Initially nil
	if table.TableColumns() != nil {
		t.Error(
			"Expected TableColumns to be nil initially",
		)
	}

	// Add columns
	col1 := table.AddColumn(1, "Product")
	if col1 == nil {
		t.Fatal("AddColumn returned nil")
	}
	if col1.Id() != 1 {
		t.Errorf(
			"Expected column Id 1, got %d",
			col1.Id(),
		)
	}
	if col1.Name() != "Product" {
		t.Errorf(
			"Expected column Name 'Product', got '%s'",
			col1.Name(),
		)
	}

	_ = table.AddColumn(2, "Price")
	_ = table.AddColumn(3, "Quantity")

	// Check count
	if table.ColumnCount() != 3 {
		t.Errorf(
			"Expected ColumnCount 3, got %d",
			table.ColumnCount(),
		)
	}

	// Get by name
	found := table.GetColumnByName("Price")
	if found == nil || found.Id() != 2 {
		t.Error(
			"GetColumnByName failed to find 'Price'",
		)
	}

	// Get by id
	found = table.GetColumnById(3)
	if found == nil ||
		found.Name() != "Quantity" {
		t.Error(
			"GetColumnById failed to find column 3",
		)
	}

	// Iterate columns
	count := 0
	for col := range table.Columns() {
		count++
		if col == nil {
			t.Error("Iterator yielded nil column")
		}
	}
	if count != 3 {
		t.Errorf(
			"Expected 3 columns in iterator, got %d",
			count,
		)
	}
}

func TestTableStyleInfo(t *testing.T) {
	table := NewTable()

	// Initially nil
	if table.TableStyleInfo() != nil {
		t.Error(
			"Expected TableStyleInfo to be nil initially",
		)
	}

	// Create table style info
	tsi := table.GetOrCreateTableStyleInfo()
	if tsi == nil {
		t.Fatal(
			"GetOrCreateTableStyleInfo returned nil",
		)
	}

	tsi.SetName("TableStyleMedium9")
	if tsi.Name() != "TableStyleMedium9" {
		t.Errorf(
			"Expected Name 'TableStyleMedium9', got '%s'",
			tsi.Name(),
		)
	}

	tsi.SetShowRowStripes(true)
	if !tsi.ShowRowStripes() {
		t.Error("Expected ShowRowStripes true")
	}

	tsi.SetShowColumnStripes(true)
	if !tsi.ShowColumnStripes() {
		t.Error("Expected ShowColumnStripes true")
	}

	tsi.SetShowFirstColumn(true)
	if !tsi.ShowFirstColumn() {
		t.Error("Expected ShowFirstColumn true")
	}

	tsi.SetShowLastColumn(true)
	if !tsi.ShowLastColumn() {
		t.Error("Expected ShowLastColumn true")
	}

	// Remove table style info
	if !table.RemoveTableStyleInfo() {
		t.Error(
			"Expected RemoveTableStyleInfo to return true",
		)
	}
	if table.TableStyleInfo() != nil {
		t.Error(
			"Expected TableStyleInfo to be nil after removal",
		)
	}
}

func TestNewTableStyleInfoWithDefaults(
	t *testing.T,
) {
	tsi := NewTableStyleInfoWithDefaults()

	if tsi.Name() != TableStyleMedium2 {
		t.Errorf(
			"Expected default Name '%s', got '%s'",
			TableStyleMedium2,
			tsi.Name(),
		)
	}

	if !tsi.ShowRowStripes() {
		t.Error(
			"Expected ShowRowStripes true by default",
		)
	}
}

func TestTableColumn(t *testing.T) {
	col := NewTableColumn()

	// Test Id
	col.SetId(5)
	if col.Id() != 5 {
		t.Errorf(
			"Expected Id 5, got %d",
			col.Id(),
		)
	}

	// Test UniqueName
	col.SetUniqueName("unique_col")
	if col.UniqueName() != "unique_col" {
		t.Errorf(
			"Expected UniqueName 'unique_col', got '%s'",
			col.UniqueName(),
		)
	}

	// Test Name
	col.SetName("Column Name")
	if col.Name() != "Column Name" {
		t.Errorf(
			"Expected Name 'Column Name', got '%s'",
			col.Name(),
		)
	}

	// Test TotalsRowLabel
	col.SetTotalsRowLabel("Total")
	if col.TotalsRowLabel() != "Total" {
		t.Errorf(
			"Expected TotalsRowLabel 'Total', got '%s'",
			col.TotalsRowLabel(),
		)
	}

	// Test TotalsRowFunction
	col.SetTotalsRowFunction(TotalsRowFunctionSum)
	if col.TotalsRowFunction() != TotalsRowFunctionSum {
		t.Errorf(
			"Expected TotalsRowFunction 'sum', got '%s'",
			col.TotalsRowFunction(),
		)
	}

	// Test various functions
	testFunctions := []TotalsRowFunction{
		TotalsRowFunctionNone,
		TotalsRowFunctionSum,
		TotalsRowFunctionMin,
		TotalsRowFunctionMax,
		TotalsRowFunctionAverage,
		TotalsRowFunctionCount,
		TotalsRowFunctionCountNums,
		TotalsRowFunctionStdDev,
		TotalsRowFunctionVar,
		TotalsRowFunctionCustom,
	}
	for _, fn := range testFunctions {
		col.SetTotalsRowFunction(fn)
		got := col.TotalsRowFunction()
		if fn == TotalsRowFunctionNone &&
			got != TotalsRowFunctionNone {
			// None removes the attribute, so it should return None
			t.Errorf(
				"Expected function None, got '%s'",
				got,
			)
		} else if fn != TotalsRowFunctionNone && got != fn {
			t.Errorf("Expected function '%s', got '%s'", fn, got)
		}
	}
}

func TestTableColumnDxfIds(t *testing.T) {
	col := NewTableColumn()

	// Test HeaderRowDxfId
	col.SetHeaderRowDxfId(1)
	id, ok := col.HeaderRowDxfId()
	if !ok || id != 1 {
		t.Errorf(
			"Expected HeaderRowDxfId 1, got %d (ok=%v)",
			id,
			ok,
		)
	}
	col.ClearHeaderRowDxfId()
	_, ok = col.HeaderRowDxfId()
	if ok {
		t.Error(
			"Expected HeaderRowDxfId to be cleared",
		)
	}

	// Test DataDxfId
	col.SetDataDxfId(2)
	id, ok = col.DataDxfId()
	if !ok || id != 2 {
		t.Errorf(
			"Expected DataDxfId 2, got %d (ok=%v)",
			id,
			ok,
		)
	}

	// Test TotalsRowDxfId
	col.SetTotalsRowDxfId(3)
	id, ok = col.TotalsRowDxfId()
	if !ok || id != 3 {
		t.Errorf(
			"Expected TotalsRowDxfId 3, got %d (ok=%v)",
			id,
			ok,
		)
	}
}

func TestTableColumnCellStyles(t *testing.T) {
	col := NewTableColumn()

	col.SetHeaderRowCellStyle("Header")
	if col.HeaderRowCellStyle() != "Header" {
		t.Errorf(
			"Expected HeaderRowCellStyle 'Header', got '%s'",
			col.HeaderRowCellStyle(),
		)
	}

	col.SetDataCellStyle("Data")
	if col.DataCellStyle() != "Data" {
		t.Errorf(
			"Expected DataCellStyle 'Data', got '%s'",
			col.DataCellStyle(),
		)
	}

	col.SetTotalsRowCellStyle("Totals")
	if col.TotalsRowCellStyle() != "Totals" {
		t.Errorf(
			"Expected TotalsRowCellStyle 'Totals', got '%s'",
			col.TotalsRowCellStyle(),
		)
	}
}

func TestCalculatedColumnFormula(t *testing.T) {
	col := NewTableColumn()

	// Initially nil
	if col.CalculatedColumnFormula() != nil {
		t.Error(
			"Expected CalculatedColumnFormula to be nil initially",
		)
	}

	// Set formula
	col.SetCalculatedColumnFormula(
		"=[Price]*[Quantity]",
	)
	ccf := col.CalculatedColumnFormula()
	if ccf == nil {
		t.Fatal(
			"CalculatedColumnFormula returned nil after set",
		)
	}
	if ccf.Formula() != "=[Price]*[Quantity]" {
		t.Errorf(
			"Expected formula '=[Price]*[Quantity]', got '%s'",
			ccf.Formula(),
		)
	}

	// Test Array attribute
	ccf.SetArray(true)
	if !ccf.Array() {
		t.Error("Expected Array true")
	}

	// Clear formula
	col.SetCalculatedColumnFormula("")
	if col.CalculatedColumnFormula() != nil {
		t.Error(
			"Expected CalculatedColumnFormula to be nil after clear",
		)
	}
}

func TestTotalsRowFormula(t *testing.T) {
	col := NewTableColumn()

	// Initially nil
	if col.TotalsRowFormula() != nil {
		t.Error(
			"Expected TotalsRowFormula to be nil initially",
		)
	}

	// Set formula (should also set function to custom)
	col.SetTotalsRowFormula(
		"=SUMPRODUCT([Price],[Quantity])",
	)
	trf := col.TotalsRowFormula()
	if trf == nil {
		t.Fatal(
			"TotalsRowFormula returned nil after set",
		)
	}
	if trf.Formula() != "=SUMPRODUCT([Price],[Quantity])" {
		t.Errorf(
			"Expected formula, got '%s'",
			trf.Formula(),
		)
	}
	if col.TotalsRowFunction() != TotalsRowFunctionCustom {
		t.Errorf(
			"Expected function 'custom', got '%s'",
			col.TotalsRowFunction(),
		)
	}

	// Test Array attribute
	trf.SetArray(true)
	if !trf.Array() {
		t.Error("Expected Array true")
	}

	// Clear formula (should also clear custom function)
	col.SetTotalsRowFormula("")
	if col.TotalsRowFormula() != nil {
		t.Error(
			"Expected TotalsRowFormula to be nil after clear",
		)
	}
	if col.TotalsRowFunction() != TotalsRowFunctionNone {
		t.Errorf(
			"Expected function 'none' after clear, got '%s'",
			col.TotalsRowFunction(),
		)
	}
}

func TestTableColumnsContainer(t *testing.T) {
	tc := NewTableColumns()

	// Initial state
	if tc.Count() != 0 {
		t.Errorf(
			"Expected Count 0, got %d",
			tc.Count(),
		)
	}
	if tc.ItemCount() != 0 {
		t.Errorf(
			"Expected ItemCount 0, got %d",
			tc.ItemCount(),
		)
	}

	// Add columns
	col1 := tc.AddColumn(1, "A")
	col2 := tc.AddColumn(2, "B")
	tc.AddColumn(3, "C")

	if tc.Count() != 3 {
		t.Errorf(
			"Expected Count 3, got %d",
			tc.Count(),
		)
	}
	if tc.ItemCount() != 3 {
		t.Errorf(
			"Expected ItemCount 3, got %d",
			tc.ItemCount(),
		)
	}

	// Get by name
	if tc.GetColumnByName("B") != col2 {
		t.Error("GetColumnByName failed")
	}

	// Get by id
	if tc.GetColumnById(1) != col1 {
		t.Error("GetColumnById failed")
	}

	// Get by index
	if tc.GetColumnByIndex(0) != col1 {
		t.Error("GetColumnByIndex failed")
	}
	if tc.GetColumnByIndex(1) != col2 {
		t.Error(
			"GetColumnByIndex failed for index 1",
		)
	}
	if tc.GetColumnByIndex(99) != nil {
		t.Error(
			"GetColumnByIndex should return nil for out of range",
		)
	}

	// Remove column
	if !tc.RemoveColumn(2) {
		t.Error("RemoveColumn should return true")
	}
	if tc.Count() != 2 {
		t.Errorf(
			"Expected Count 2 after removal, got %d",
			tc.Count(),
		)
	}
	if tc.GetColumnByName("B") != nil {
		t.Error("Column B should be removed")
	}
}

func TestTableClone(t *testing.T) {
	table := NewTableWithDefaults(
		1,
		"Table1",
		"Table1",
		"A1:D10",
	)
	table.AddColumn(1, "Product")
	table.AddColumn(2, "Price")
	tsi := table.GetOrCreateTableStyleInfo()
	tsi.SetName("TableStyleMedium9")

	clonedResult := table.Clone()
	cloned, ok := clonedResult.(*Table)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Table",
			clonedResult,
		)
	}

	if cloned == table {
		t.Error(
			"Clone should return a new instance",
		)
	}
	if cloned.Id() != 1 {
		t.Errorf(
			"Cloned Id should be 1, got %d",
			cloned.Id(),
		)
	}
	if cloned.Name() != "Table1" {
		t.Errorf(
			"Cloned Name should be 'Table1', got '%s'",
			cloned.Name(),
		)
	}
	if cloned.ColumnCount() != 2 {
		t.Errorf(
			"Cloned ColumnCount should be 2, got %d",
			cloned.ColumnCount(),
		)
	}
}

func TestTableStyleInfoClone(t *testing.T) {
	tsi := NewTableStyleInfo()
	tsi.SetName("TestStyle")
	tsi.SetShowRowStripes(true)

	clonedResult := tsi.Clone()
	cloned, ok := clonedResult.(*TableStyleInfo)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *TableStyleInfo",
			clonedResult,
		)
	}

	if cloned == tsi {
		t.Error(
			"Clone should return a new instance",
		)
	}
	if cloned.Name() != "TestStyle" {
		t.Errorf(
			"Cloned Name should be 'TestStyle', got '%s'",
			cloned.Name(),
		)
	}
	if !cloned.ShowRowStripes() {
		t.Error(
			"Cloned ShowRowStripes should be true",
		)
	}
}
