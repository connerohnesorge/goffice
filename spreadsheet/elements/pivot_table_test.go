package elements

import (
	"testing"
)

const (
	testRefA1D10 = "A1:D10"
)

func TestNewPivotTableDefinition(t *testing.T) {
	pt := NewPivotTableDefinition()

	if pt == nil {
		t.Fatal(
			"NewPivotTableDefinition returned nil",
		)
	}

	if pt.LocalName() != "pivotTableDefinition" {
		t.Errorf(
			"Expected LocalName 'pivotTableDefinition', got '%s'",
			pt.LocalName(),
		)
	}

	if pt.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"Expected NamespaceURI '%s', got '%s'",
			NamespaceSML,
			pt.NamespaceURI(),
		)
	}
}

func TestPivotTableDefinitionAttributes(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Test Name
	pt.SetName("SalesPivot")
	if pt.Name() != "SalesPivot" {
		t.Errorf(
			"Expected Name 'SalesPivot', got '%s'",
			pt.Name(),
		)
	}

	// Test CacheId
	pt.SetCacheId(42)
	if pt.CacheId() != 42 {
		t.Errorf(
			"Expected CacheId 42, got %d",
			pt.CacheId(),
		)
	}

	// Test DataOnRows
	pt.SetDataOnRows(true)
	if !pt.DataOnRows() {
		t.Error("Expected DataOnRows true")
	}
	pt.SetDataOnRows(false)
	if pt.DataOnRows() {
		t.Error("Expected DataOnRows false")
	}

	// Test DataPosition
	pt.SetDataPosition(5)
	pos, ok := pt.DataPosition()
	if !ok || pos != 5 {
		t.Errorf(
			"Expected DataPosition 5, got %d (ok=%v)",
			pos,
			ok,
		)
	}
	pt.ClearDataPosition()
	_, ok = pt.DataPosition()
	if ok {
		t.Error(
			"Expected DataPosition to be cleared",
		)
	}

	// Test DataCaption
	pt.SetDataCaption("Values")
	if pt.DataCaption() != "Values" {
		t.Errorf(
			"Expected DataCaption 'Values', got '%s'",
			pt.DataCaption(),
		)
	}

	// Test GrandTotalCaption
	pt.SetGrandTotalCaption("Total")
	if pt.GrandTotalCaption() != "Total" {
		t.Errorf(
			"Expected GrandTotalCaption 'Total', got '%s'",
			pt.GrandTotalCaption(),
		)
	}

	// Test ErrorCaption
	pt.SetErrorCaption("#ERR")
	if pt.ErrorCaption() != "#ERR" {
		t.Errorf(
			"Expected ErrorCaption '#ERR', got '%s'",
			pt.ErrorCaption(),
		)
	}

	// Test MissingCaption
	pt.SetMissingCaption("(blank)")
	if pt.MissingCaption() != "(blank)" {
		t.Errorf(
			"Expected MissingCaption '(blank)', got '%s'",
			pt.MissingCaption(),
		)
	}
}

func TestPivotTableDefinitionBooleanAttributes(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Test ShowError
	pt.SetShowError(true)
	if !pt.ShowError() {
		t.Error("Expected ShowError true")
	}

	// Test ShowMissing (default true)
	if !pt.ShowMissing() {
		t.Error(
			"Expected ShowMissing true by default",
		)
	}
	pt.SetShowMissing(false)
	if pt.ShowMissing() {
		t.Error("Expected ShowMissing false")
	}

	// Test RowGrandTotals (default true)
	if !pt.RowGrandTotals() {
		t.Error(
			"Expected RowGrandTotals true by default",
		)
	}
	pt.SetRowGrandTotals(false)
	if pt.RowGrandTotals() {
		t.Error("Expected RowGrandTotals false")
	}

	// Test ColGrandTotals (default true)
	if !pt.ColGrandTotals() {
		t.Error(
			"Expected ColGrandTotals true by default",
		)
	}
	pt.SetColGrandTotals(false)
	if pt.ColGrandTotals() {
		t.Error("Expected ColGrandTotals false")
	}

	// Test Compact (default true)
	if !pt.Compact() {
		t.Error(
			"Expected Compact true by default",
		)
	}
	pt.SetCompact(false)
	if pt.Compact() {
		t.Error("Expected Compact false")
	}

	// Test ShowHeaders (default true)
	if !pt.ShowHeaders() {
		t.Error(
			"Expected ShowHeaders true by default",
		)
	}
	pt.SetShowHeaders(false)
	if pt.ShowHeaders() {
		t.Error("Expected ShowHeaders false")
	}
}

func TestPivotTableDefinitionVersions(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Test UpdatedVersion
	pt.SetUpdatedVersion(5)
	if pt.UpdatedVersion() != 5 {
		t.Errorf(
			"Expected UpdatedVersion 5, got %d",
			pt.UpdatedVersion(),
		)
	}

	// Test MinRefreshableVersion
	pt.SetMinRefreshableVersion(3)
	if pt.MinRefreshableVersion() != 3 {
		t.Errorf(
			"Expected MinRefreshableVersion 3, got %d",
			pt.MinRefreshableVersion(),
		)
	}

	// Test CreatedVersion
	pt.SetCreatedVersion(4)
	if pt.CreatedVersion() != 4 {
		t.Errorf(
			"Expected CreatedVersion 4, got %d",
			pt.CreatedVersion(),
		)
	}
}

func TestPivotTableDefinitionStyles(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Test PivotTableStyle
	pt.SetPivotTableStyle("PivotStyleMedium9")
	if pt.PivotTableStyle() != "PivotStyleMedium9" {
		t.Errorf(
			"Expected PivotTableStyle 'PivotStyleMedium9', got '%s'",
			pt.PivotTableStyle(),
		)
	}

	// Test PageStyle
	pt.SetPageStyle("PageStyle1")
	if pt.PageStyle() != "PageStyle1" {
		t.Errorf(
			"Expected PageStyle 'PageStyle1', got '%s'",
			pt.PageStyle(),
		)
	}

	// Test VacatedStyle
	pt.SetVacatedStyle("VacatedStyle1")
	if pt.VacatedStyle() != "VacatedStyle1" {
		t.Errorf(
			"Expected VacatedStyle 'VacatedStyle1', got '%s'",
			pt.VacatedStyle(),
		)
	}

	// Test Tag
	pt.SetTag("MyPivotTag")
	if pt.Tag() != "MyPivotTag" {
		t.Errorf(
			"Expected Tag 'MyPivotTag', got '%s'",
			pt.Tag(),
		)
	}
}

func TestPivotTableDefinitionLocation(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Initially nil
	if pt.Location() != nil {
		t.Error(
			"Expected Location to be nil initially",
		)
	}

	// Create location
	loc := pt.GetOrCreateLocation()
	if loc == nil {
		t.Fatal(
			"GetOrCreateLocation returned nil",
		)
	}

	loc.SetRef("A3:D10")
	if loc.Ref() != "A3:D10" {
		t.Errorf(
			"Expected Ref 'A3:D10', got '%s'",
			loc.Ref(),
		)
	}

	loc.SetFirstHeaderRow(1)
	if loc.FirstHeaderRow() != 1 {
		t.Errorf(
			"Expected FirstHeaderRow 1, got %d",
			loc.FirstHeaderRow(),
		)
	}

	loc.SetFirstDataRow(2)
	if loc.FirstDataRow() != 2 {
		t.Errorf(
			"Expected FirstDataRow 2, got %d",
			loc.FirstDataRow(),
		)
	}

	loc.SetFirstDataCol(1)
	if loc.FirstDataCol() != 1 {
		t.Errorf(
			"Expected FirstDataCol 1, got %d",
			loc.FirstDataCol(),
		)
	}

	// Get existing location
	loc2 := pt.GetOrCreateLocation()
	if loc2 != loc {
		t.Error(
			"Expected GetOrCreateLocation to return existing element",
		)
	}
}

func TestPivotTableDefinitionPivotFields(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Initially nil
	if pt.PivotFields() != nil {
		t.Error(
			"Expected PivotFields to be nil initially",
		)
	}

	// Create pivot fields
	pf := pt.GetOrCreatePivotFields()
	if pf == nil {
		t.Fatal(
			"GetOrCreatePivotFields returned nil",
		)
	}

	// Add a pivot field using AddField
	field := pf.AddField()
	if field == nil {
		t.Fatal("AddField returned nil")
	}

	field.SetName("Category")
	if field.Name() != "Category" {
		t.Errorf(
			"Expected Name 'Category', got '%s'",
			field.Name(),
		)
	}

	// Get existing pivot fields
	pf2 := pt.GetOrCreatePivotFields()
	if pf2 != pf {
		t.Error(
			"Expected GetOrCreatePivotFields to return existing element",
		)
	}
}

func TestPivotTableDefinitionDataFields(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Initially nil
	if pt.DataFields() != nil {
		t.Error(
			"Expected DataFields to be nil initially",
		)
	}

	// Create data fields
	df := pt.GetOrCreateDataFields()
	if df == nil {
		t.Fatal(
			"GetOrCreateDataFields returned nil",
		)
	}

	// Add a data field with required fld parameter
	field := df.AddDataField(2)
	if field == nil {
		t.Fatal("AddDataField returned nil")
	}

	field.SetName("Sum of Sales")
	if field.Name() != "Sum of Sales" {
		t.Errorf(
			"Expected Name 'Sum of Sales', got '%s'",
			field.Name(),
		)
	}

	if field.Fld() != 2 {
		t.Errorf(
			"Expected Fld 2, got %d",
			field.Fld(),
		)
	}
}

func TestPivotTableDefinitionRowColFields(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Test RowFields
	rf := pt.GetOrCreateRowFields()
	if rf == nil {
		t.Fatal(
			"GetOrCreateRowFields returned nil",
		)
	}

	field := rf.AddField(0)
	if field.X() != 0 {
		t.Errorf(
			"Expected X 0, got %d",
			field.X(),
		)
	}

	// Test ColFields
	cf := pt.GetOrCreateColFields()
	if cf == nil {
		t.Fatal(
			"GetOrCreateColFields returned nil",
		)
	}

	colField := cf.AddField(1)
	if colField.X() != 1 {
		t.Errorf(
			"Expected X 1, got %d",
			colField.X(),
		)
	}
}

func TestPivotTableDefinitionPageFields(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Create page fields
	pf := pt.GetOrCreatePageFields()
	if pf == nil {
		t.Fatal(
			"GetOrCreatePageFields returned nil",
		)
	}

	// Add a page field with required fld parameter
	field := pf.AddPageField(3)
	if field == nil {
		t.Fatal("AddPageField returned nil")
	}

	if field.Fld() != 3 {
		t.Errorf(
			"Expected Fld 3, got %d",
			field.Fld(),
		)
	}

	field.SetHier(-1)
	if field.Hier() != -1 {
		t.Errorf(
			"Expected Hier -1, got %d",
			field.Hier(),
		)
	}

	field.SetName("Region")
	if field.Name() != "Region" {
		t.Errorf(
			"Expected Name 'Region', got '%s'",
			field.Name(),
		)
	}
}

func TestPivotTableDefinitionStyleInfo(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Initially nil
	if pt.PivotTableStyleInfo() != nil {
		t.Error(
			"Expected PivotTableStyleInfo to be nil initially",
		)
	}

	// Create style info
	psi := pt.GetOrCreatePivotTableStyleInfo()
	if psi == nil {
		t.Fatal(
			"GetOrCreatePivotTableStyleInfo returned nil",
		)
	}

	psi.SetName("PivotStyleMedium9")
	if psi.Name() != "PivotStyleMedium9" {
		t.Errorf(
			"Expected Name 'PivotStyleMedium9', got '%s'",
			psi.Name(),
		)
	}

	psi.SetShowRowHeaders(true)
	if !psi.ShowRowHeaders() {
		t.Error("Expected ShowRowHeaders true")
	}

	psi.SetShowColHeaders(true)
	if !psi.ShowColHeaders() {
		t.Error("Expected ShowColHeaders true")
	}

	psi.SetShowRowStripes(true)
	if !psi.ShowRowStripes() {
		t.Error("Expected ShowRowStripes true")
	}

	psi.SetShowColStripes(true)
	if !psi.ShowColStripes() {
		t.Error("Expected ShowColStripes true")
	}

	psi.SetShowLastColumn(true)
	if !psi.ShowLastColumn() {
		t.Error("Expected ShowLastColumn true")
	}
}

func TestPivotTableDefinitionRowColItems(
	t *testing.T,
) {
	pt := NewPivotTableDefinition()

	// Test RowItems
	ri := pt.GetOrCreateRowItems()
	if ri == nil {
		t.Fatal(
			"GetOrCreateRowItems returned nil",
		)
	}

	ri.SetCount(5)
	if ri.Count() != 5 {
		t.Errorf(
			"Expected Count 5, got %d",
			ri.Count(),
		)
	}

	item := ri.AddItem()
	if item == nil {
		t.Fatal("AddItem returned nil")
	}

	item.SetT("grand")
	if item.T() != "grand" {
		t.Errorf(
			"Expected T 'grand', got '%s'",
			item.T(),
		)
	}

	// Test ColItems
	ci := pt.GetOrCreateColItems()
	if ci == nil {
		t.Fatal(
			"GetOrCreateColItems returned nil",
		)
	}

	ci.SetCount(3)
	if ci.Count() != 3 {
		t.Errorf(
			"Expected Count 3, got %d",
			ci.Count(),
		)
	}
}

func TestLocation(t *testing.T) {
	loc := NewLocation()

	if loc == nil {
		t.Fatal("NewLocation returned nil")
	}

	if loc.LocalName() != "location" {
		t.Errorf(
			"Expected LocalName 'location', got '%s'",
			loc.LocalName(),
		)
	}

	// Test RowPageCount
	loc.SetRowPageCount(2)
	if loc.RowPageCount() != 2 {
		t.Errorf(
			"Expected RowPageCount 2, got %d",
			loc.RowPageCount(),
		)
	}

	// Test ColPageCount
	loc.SetColPageCount(1)
	if loc.ColPageCount() != 1 {
		t.Errorf(
			"Expected ColPageCount 1, got %d",
			loc.ColPageCount(),
		)
	}
}

func TestIElement(t *testing.T) {
	i := NewI()

	if i == nil {
		t.Fatal("NewI returned nil")
	}

	if i.LocalName() != "i" {
		t.Errorf(
			"Expected LocalName 'i', got '%s'",
			i.LocalName(),
		)
	}

	// Test T (default is "data")
	if i.T() != "data" {
		t.Errorf(
			"Expected T 'data', got '%s'",
			i.T(),
		)
	}

	i.SetT("grand")
	if i.T() != "grand" {
		t.Errorf(
			"Expected T 'grand', got '%s'",
			i.T(),
		)
	}

	// Test R
	i.SetR(3)
	if i.R() != 3 {
		t.Errorf("Expected R 3, got %d", i.R())
	}

	// Test I (index)
	i.SetI(5)
	if i.I() != 5 {
		t.Errorf("Expected I 5, got %d", i.I())
	}
}

func TestPivotTableDefinitionClone(t *testing.T) {
	pt := NewPivotTableDefinition()
	pt.SetName("TestPivot")
	pt.SetCacheId(1)
	pt.GetOrCreateLocation().SetRef("A3:D10")

	clonedResult := pt.Clone()
	cloned, ok := clonedResult.(*PivotTableDefinition)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *PivotTableDefinition",
			clonedResult,
		)
	}

	if cloned == pt {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.Name() != "TestPivot" {
		t.Errorf(
			"Cloned Name should be 'TestPivot', got '%s'",
			cloned.Name(),
		)
	}

	if cloned.CacheId() != 1 {
		t.Errorf(
			"Cloned CacheId should be 1, got %d",
			cloned.CacheId(),
		)
	}
}

func TestLocationClone(t *testing.T) {
	loc := NewLocation()
	loc.SetRef(testRefA1D10)
	loc.SetFirstDataRow(2)

	clonedResult := loc.Clone()
	cloned, ok := clonedResult.(*Location)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Location",
			clonedResult,
		)
	}

	if cloned == loc {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.Ref() != testRefA1D10 {
		t.Errorf(
			"Cloned Ref should be '%s', got '%s'",
			testRefA1D10,
			cloned.Ref(),
		)
	}
}

func TestPivotTableStyleInfoClone(t *testing.T) {
	psi := NewPivotTableStyleInfo()
	psi.SetName("TestStyle")
	psi.SetShowRowHeaders(true)

	clonedResult := psi.Clone()
	cloned, ok := clonedResult.(*PivotTableStyleInfo)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *PivotTableStyleInfo",
			clonedResult,
		)
	}

	if cloned == psi {
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

	if !cloned.ShowRowHeaders() {
		t.Error(
			"Cloned ShowRowHeaders should be true",
		)
	}
}
