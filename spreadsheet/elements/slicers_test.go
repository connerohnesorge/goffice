package elements

import (
	"testing"
)

func TestNewSlicers(t *testing.T) {
	slicers := NewSlicers()

	if slicers == nil {
		t.Fatal("NewSlicers returned nil")
	}

	if slicers.LocalName() != "slicers" {
		t.Errorf(
			"Expected LocalName 'slicers', got '%s'",
			slicers.LocalName(),
		)
	}

	if slicers.NamespaceURI() != NamespaceSlicerX14 {
		t.Errorf(
			"Expected NamespaceURI '%s', got '%s'",
			NamespaceSlicerX14,
			slicers.NamespaceURI(),
		)
	}
}

func TestSlicersAddSlicer(t *testing.T) {
	slicers := NewSlicers()

	// Add a slicer with required name parameter
	slicer := slicers.AddSlicer("TestSlicer")
	if slicer == nil {
		t.Fatal("AddSlicer returned nil")
	}

	if slicer.LocalName() != "slicer" {
		t.Errorf(
			"Expected LocalName 'slicer', got '%s'",
			slicer.LocalName(),
		)
	}

	if slicer.Name() != "TestSlicer" {
		t.Errorf(
			"Expected Name 'TestSlicer', got '%s'",
			slicer.Name(),
		)
	}

	// Check slicer count
	if slicers.SlicerCount() != 1 {
		t.Errorf(
			"Expected SlicerCount 1, got %d",
			slicers.SlicerCount(),
		)
	}
}

func TestSlicersIteration(t *testing.T) {
	slicers := NewSlicers()

	// Add slicers
	slicers.AddSlicer("Slicer1")
	slicers.AddSlicer("Slicer2")
	slicers.AddSlicer("Slicer3")

	// Iterate
	count := 0
	for slicer := range slicers.Slicers() {
		count++
		if slicer == nil {
			t.Error("Iterator yielded nil slicer")
		}
	}
	if count != 3 {
		t.Errorf(
			"Expected 3 slicers in iterator, got %d",
			count,
		)
	}
}

func TestSlicersGetSlicerByName(t *testing.T) {
	slicers := NewSlicers()

	// Add slicers
	slicers.AddSlicer("First")
	slicers.AddSlicer("Second")

	// Get by name
	slicer := slicers.GetSlicerByName("Second")
	if slicer == nil {
		t.Fatal(
			"GetSlicerByName('Second') returned nil",
		)
	}
	if slicer.Name() != "Second" {
		t.Errorf(
			"Expected Name 'Second', got '%s'",
			slicer.Name(),
		)
	}

	// Non-existent name
	if slicers.GetSlicerByName(
		"NonExistent",
	) != nil {
		t.Error(
			"Expected nil for non-existent name",
		)
	}
}

func TestSlicerAttributes(t *testing.T) {
	slicer := NewSlicer()

	// Test Name
	slicer.SetName("CategorySlicer")
	if slicer.Name() != "CategorySlicer" {
		t.Errorf(
			"Expected Name 'CategorySlicer', got '%s'",
			slicer.Name(),
		)
	}

	// Test Cache
	slicer.SetCache("SlicerCache1")
	if slicer.Cache() != "SlicerCache1" {
		t.Errorf(
			"Expected Cache 'SlicerCache1', got '%s'",
			slicer.Cache(),
		)
	}

	// Test Caption
	slicer.SetCaption("Category")
	if slicer.Caption() != "Category" {
		t.Errorf(
			"Expected Caption 'Category', got '%s'",
			slicer.Caption(),
		)
	}

	// Test Style
	slicer.SetStyle("SlicerStyleLight1")
	if slicer.Style() != "SlicerStyleLight1" {
		t.Errorf(
			"Expected Style 'SlicerStyleLight1', got '%s'",
			slicer.Style(),
		)
	}

	// Test LockedPosition
	slicer.SetLockedPosition(true)
	if !slicer.LockedPosition() {
		t.Error("Expected LockedPosition true")
	}
	slicer.SetLockedPosition(false)
	if slicer.LockedPosition() {
		t.Error("Expected LockedPosition false")
	}
}

func TestSlicerNumericAttributes(t *testing.T) {
	slicer := NewSlicer()

	// Test RowHeight
	slicer.SetRowHeight(200000)
	if slicer.RowHeight() != 200000 {
		t.Errorf(
			"Expected RowHeight 200000, got %d",
			slicer.RowHeight(),
		)
	}

	// Test StartItem
	slicer.SetStartItem(5)
	if slicer.StartItem() != 5 {
		t.Errorf(
			"Expected StartItem 5, got %d",
			slicer.StartItem(),
		)
	}

	// Test ColumnCount (default 1)
	if slicer.ColumnCount() != 1 {
		t.Errorf(
			"Expected ColumnCount 1 by default, got %d",
			slicer.ColumnCount(),
		)
	}
	slicer.SetColumnCount(3)
	if slicer.ColumnCount() != 3 {
		t.Errorf(
			"Expected ColumnCount 3, got %d",
			slicer.ColumnCount(),
		)
	}

	// Test Level
	slicer.SetLevel(2)
	if slicer.Level() != 2 {
		t.Errorf(
			"Expected Level 2, got %d",
			slicer.Level(),
		)
	}
}

func TestSlicerBooleanAttributes(t *testing.T) {
	slicer := NewSlicer()

	// Test ShowCaption (default true)
	if !slicer.ShowCaption() {
		t.Error(
			"Expected ShowCaption true by default",
		)
	}
	slicer.SetShowCaption(false)
	if slicer.ShowCaption() {
		t.Error("Expected ShowCaption false")
	}

	// LockedPosition already tested above
}

func TestSlicerCacheDefinition(t *testing.T) {
	scd := NewSlicerCacheDefinition()

	if scd == nil {
		t.Fatal(
			"NewSlicerCacheDefinition returned nil",
		)
	}

	if scd.LocalName() != "slicerCacheDefinition" {
		t.Errorf(
			"Expected LocalName 'slicerCacheDefinition', got '%s'",
			scd.LocalName(),
		)
	}

	// Test Name
	scd.SetName("SlicerCache1")
	if scd.Name() != "SlicerCache1" {
		t.Errorf(
			"Expected Name 'SlicerCache1', got '%s'",
			scd.Name(),
		)
	}

	// Test SourceName
	scd.SetSourceName("Category")
	if scd.SourceName() != "Category" {
		t.Errorf(
			"Expected SourceName 'Category', got '%s'",
			scd.SourceName(),
		)
	}
}

func TestSlicerCacheDefinitionPivotTables(
	t *testing.T,
) {
	scd := NewSlicerCacheDefinition()

	// Initially nil
	if scd.PivotTables() != nil {
		t.Error(
			"Expected PivotTables to be nil initially",
		)
	}

	// Create pivot tables
	pt := scd.GetOrCreatePivotTables()
	if pt == nil {
		t.Fatal(
			"GetOrCreatePivotTables returned nil",
		)
	}

	// Get existing
	pt2 := scd.GetOrCreatePivotTables()
	if pt2 != pt {
		t.Error(
			"Expected GetOrCreatePivotTables to return existing element",
		)
	}
}

func TestSlicerCacheDefinitionData(t *testing.T) {
	scd := NewSlicerCacheDefinition()

	// Initially nil
	if scd.Data() != nil {
		t.Error(
			"Expected Data to be nil initially",
		)
	}

	// Create data
	data := scd.GetOrCreateData()
	if data == nil {
		t.Fatal("GetOrCreateData returned nil")
	}

	// Get existing
	data2 := scd.GetOrCreateData()
	if data2 != data {
		t.Error(
			"Expected GetOrCreateData to return existing element",
		)
	}
}

func TestSlicerCachePivotTables(t *testing.T) {
	scpt := NewSlicerCachePivotTables()

	if scpt == nil {
		t.Fatal(
			"NewSlicerCachePivotTables returned nil",
		)
	}

	if scpt.LocalName() != "pivotTables" {
		t.Errorf(
			"Expected LocalName 'pivotTables', got '%s'",
			scpt.LocalName(),
		)
	}
}

func TestSlicerCacheData(t *testing.T) {
	scData := NewSlicerCacheData()

	if scData == nil {
		t.Fatal("NewSlicerCacheData returned nil")
	}

	if scData.LocalName() != "data" {
		t.Errorf(
			"Expected LocalName 'data', got '%s'",
			scData.LocalName(),
		)
	}
}

func TestSlicersClone(t *testing.T) {
	slicers := NewSlicers()
	s := slicers.AddSlicer("TestSlicer")
	s.SetCaption("Test")

	clonedResult := slicers.Clone()
	cloned, ok := clonedResult.(*Slicers)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Slicers",
			clonedResult,
		)
	}

	if cloned == slicers {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.SlicerCount() != 1 {
		t.Errorf(
			"Cloned SlicerCount should be 1, got %d",
			cloned.SlicerCount(),
		)
	}
}

func TestSlicerClone(t *testing.T) {
	slicer := NewSlicer()
	slicer.SetName("MySlicer")
	slicer.SetCaption("My Caption")
	slicer.SetColumnCount(2)

	clonedResult := slicer.Clone()
	cloned, ok := clonedResult.(*Slicer)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Slicer",
			clonedResult,
		)
	}

	if cloned == slicer {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.Name() != "MySlicer" {
		t.Errorf(
			"Cloned Name should be 'MySlicer', got '%s'",
			cloned.Name(),
		)
	}

	if cloned.Caption() != "My Caption" {
		t.Errorf(
			"Cloned Caption should be 'My Caption', got '%s'",
			cloned.Caption(),
		)
	}

	if cloned.ColumnCount() != 2 {
		t.Errorf(
			"Cloned ColumnCount should be 2, got %d",
			cloned.ColumnCount(),
		)
	}
}

func TestSlicerCacheDefinitionClone(
	t *testing.T,
) {
	scd := NewSlicerCacheDefinition()
	scd.SetName("CacheName")
	scd.SetSourceName("SourceField")

	clonedResult := scd.Clone()
	cloned, ok := clonedResult.(*SlicerCacheDefinition)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *SlicerCacheDefinition",
			clonedResult,
		)
	}

	if cloned == scd {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.Name() != "CacheName" {
		t.Errorf(
			"Cloned Name should be 'CacheName', got '%s'",
			cloned.Name(),
		)
	}

	if cloned.SourceName() != "SourceField" {
		t.Errorf(
			"Cloned SourceName should be 'SourceField', got '%s'",
			cloned.SourceName(),
		)
	}
}

func TestSlicerCachePivotTablesClone(
	t *testing.T,
) {
	scpt := NewSlicerCachePivotTables()

	clonedResult := scpt.Clone()
	cloned, ok := clonedResult.(*SlicerCachePivotTables)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *SlicerCachePivotTables",
			clonedResult,
		)
	}

	if cloned == scpt {
		t.Error(
			"Clone should return a new instance",
		)
	}
}

func TestSlicerCacheDataClone(t *testing.T) {
	scData := NewSlicerCacheData()

	clonedResult := scData.Clone()
	cloned, ok := clonedResult.(*SlicerCacheData)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *SlicerCacheData",
			clonedResult,
		)
	}

	if cloned == scData {
		t.Error(
			"Clone should return a new instance",
		)
	}
}
