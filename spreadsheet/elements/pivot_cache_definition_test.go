package elements

import (
	"testing"
)

const (
	testCategoryName = "Category"
	testSheet1Name   = "Sheet1"
)

func TestNewPivotCacheDefinition(t *testing.T) {
	pcd := NewPivotCacheDefinition()

	if pcd == nil {
		t.Fatal(
			"NewPivotCacheDefinition returned nil",
		)
	}

	if pcd.LocalName() != "pivotCacheDefinition" {
		t.Errorf(
			"Expected LocalName 'pivotCacheDefinition', got '%s'",
			pcd.LocalName(),
		)
	}

	if pcd.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"Expected NamespaceURI '%s', got '%s'",
			NamespaceSML,
			pcd.NamespaceURI(),
		)
	}
}

func TestPivotCacheDefinitionAttributes(
	t *testing.T,
) {
	pcd := NewPivotCacheDefinition()

	// Test RelationshipId
	pcd.SetRelationshipId("rId1")
	if pcd.RelationshipId() != "rId1" {
		t.Errorf(
			"Expected RelationshipId 'rId1', got '%s'",
			pcd.RelationshipId(),
		)
	}

	// Test Invalid
	pcd.SetInvalid(true)
	if !pcd.Invalid() {
		t.Error("Expected Invalid true")
	}
	pcd.SetInvalid(false)
	if pcd.Invalid() {
		t.Error("Expected Invalid false")
	}

	// Test SaveData
	if !pcd.SaveData() { // Default is true
		t.Error(
			"Expected SaveData true by default",
		)
	}
	pcd.SetSaveData(false)
	if pcd.SaveData() {
		t.Error("Expected SaveData false")
	}

	// Test RefreshOnLoad
	pcd.SetRefreshOnLoad(true)
	if !pcd.RefreshOnLoad() {
		t.Error("Expected RefreshOnLoad true")
	}

	// Test OptimizeMemory
	pcd.SetOptimizeMemory(true)
	if !pcd.OptimizeMemory() {
		t.Error("Expected OptimizeMemory true")
	}

	// Test EnableRefresh (default true)
	if !pcd.EnableRefresh() {
		t.Error(
			"Expected EnableRefresh true by default",
		)
	}
	pcd.SetEnableRefresh(false)
	if pcd.EnableRefresh() {
		t.Error("Expected EnableRefresh false")
	}

	// Test UpgradeOnRefresh
	pcd.SetUpgradeOnRefresh(true)
	if !pcd.UpgradeOnRefresh() {
		t.Error("Expected UpgradeOnRefresh true")
	}
}

func TestPivotCacheDefinitionNumericAttributes(
	t *testing.T,
) {
	pcd := NewPivotCacheDefinition()

	// Test RecordCount
	pcd.SetRecordCount(100)
	if pcd.RecordCount() != 100 {
		t.Errorf(
			"Expected RecordCount 100, got %d",
			pcd.RecordCount(),
		)
	}

	// Test CreatedVersion
	pcd.SetCreatedVersion(5)
	if pcd.CreatedVersion() != 5 {
		t.Errorf(
			"Expected CreatedVersion 5, got %d",
			pcd.CreatedVersion(),
		)
	}

	// Test RefreshedVersion
	pcd.SetRefreshedVersion(6)
	if pcd.RefreshedVersion() != 6 {
		t.Errorf(
			"Expected RefreshedVersion 6, got %d",
			pcd.RefreshedVersion(),
		)
	}

	// Test MinRefreshableVersion
	pcd.SetMinRefreshableVersion(3)
	if pcd.MinRefreshableVersion() != 3 {
		t.Errorf(
			"Expected MinRefreshableVersion 3, got %d",
			pcd.MinRefreshableVersion(),
		)
	}
}

func TestPivotCacheDefinitionCacheSource(
	t *testing.T,
) {
	pcd := NewPivotCacheDefinition()

	// Initially nil
	if pcd.CacheSource() != nil {
		t.Error(
			"Expected CacheSource to be nil initially",
		)
	}

	// Create cache source
	cs := pcd.GetOrCreateCacheSource()
	if cs == nil {
		t.Fatal(
			"GetOrCreateCacheSource returned nil",
		)
	}

	cs.SetType("worksheet")
	if cs.Type() != "worksheet" {
		t.Errorf(
			"Expected Type 'worksheet', got '%s'",
			cs.Type(),
		)
	}

	// Get existing cache source
	cs2 := pcd.GetOrCreateCacheSource()
	if cs2 != cs {
		t.Error(
			"Expected GetOrCreateCacheSource to return existing element",
		)
	}
}

func TestPivotCacheDefinitionCacheFields(
	t *testing.T,
) {
	pcd := NewPivotCacheDefinition()

	// Initially nil
	if pcd.CacheFields() != nil {
		t.Error(
			"Expected CacheFields to be nil initially",
		)
	}

	// Create cache fields
	cfs := pcd.GetOrCreateCacheFields()
	if cfs == nil {
		t.Fatal(
			"GetOrCreateCacheFields returned nil",
		)
	}

	// Add a cache field using AddField
	cf := cfs.AddField(testCategoryName)
	if cf == nil {
		t.Fatal("AddField returned nil")
	}

	if cf.Name() != testCategoryName {
		t.Errorf(
			"Expected Name '%s', got '%s'",
			testCategoryName,
			cf.Name(),
		)
	}

	// Get existing cache fields
	cfs2 := pcd.GetOrCreateCacheFields()
	if cfs2 != cfs {
		t.Error(
			"Expected GetOrCreateCacheFields to return existing element",
		)
	}
}

func TestCacheSource(t *testing.T) {
	cs := NewCacheSource()

	if cs == nil {
		t.Fatal("NewCacheSource returned nil")
	}

	if cs.LocalName() != "cacheSource" {
		t.Errorf(
			"Expected LocalName 'cacheSource', got '%s'",
			cs.LocalName(),
		)
	}

	// Test Type (default is worksheet)
	if cs.Type() != "worksheet" {
		t.Errorf(
			"Expected default Type 'worksheet', got '%s'",
			cs.Type(),
		)
	}

	cs.SetType("external")
	if cs.Type() != "external" {
		t.Errorf(
			"Expected Type 'external', got '%s'",
			cs.Type(),
		)
	}

	// Test ConnectionId
	cs.SetConnectionId(5)
	cid, ok := cs.ConnectionId()
	if !ok || cid != 5 {
		t.Errorf(
			"Expected ConnectionId 5, got %d (ok=%v)",
			cid,
			ok,
		)
	}
}

func TestCacheSourceWorksheetSource(
	t *testing.T,
) {
	cs := NewCacheSource()
	cs.SetType("worksheet")

	// Initially nil
	if cs.WorksheetSource() != nil {
		t.Error(
			"Expected WorksheetSource to be nil initially",
		)
	}

	// Create worksheet source
	ws := cs.GetOrCreateWorksheetSource()
	if ws == nil {
		t.Fatal(
			"GetOrCreateWorksheetSource returned nil",
		)
	}

	ws.SetRef("A1:D100")
	if ws.Ref() != "A1:D100" {
		t.Errorf(
			"Expected Ref 'A1:D100', got '%s'",
			ws.Ref(),
		)
	}

	ws.SetSheet(testSheet1Name)
	if ws.Sheet() != testSheet1Name {
		t.Errorf(
			"Expected Sheet '%s', got '%s'",
			testSheet1Name,
			ws.Sheet(),
		)
	}

	ws.SetName("SalesData")
	if ws.Name() != "SalesData" {
		t.Errorf(
			"Expected Name 'SalesData', got '%s'",
			ws.Name(),
		)
	}

	ws.SetRelationshipId("rId2")
	if ws.RelationshipId() != "rId2" {
		t.Errorf(
			"Expected RelationshipId 'rId2', got '%s'",
			ws.RelationshipId(),
		)
	}
}

func TestCacheField(t *testing.T) {
	cf := NewCacheField()

	if cf == nil {
		t.Fatal("NewCacheField returned nil")
	}

	if cf.LocalName() != "cacheField" {
		t.Errorf(
			"Expected LocalName 'cacheField', got '%s'",
			cf.LocalName(),
		)
	}

	// Test Name
	cf.SetName("ProductName")
	if cf.Name() != "ProductName" {
		t.Errorf(
			"Expected Name 'ProductName', got '%s'",
			cf.Name(),
		)
	}

	// Test Caption
	cf.SetCaption("Product Name")
	if cf.Caption() != "Product Name" {
		t.Errorf(
			"Expected Caption 'Product Name', got '%s'",
			cf.Caption(),
		)
	}

	// Test NumFmtId
	cf.SetNumFmtId(0)
	id, ok := cf.NumFmtId()
	if !ok || id != 0 {
		t.Errorf(
			"Expected NumFmtId 0, got %d (ok=%v)",
			id,
			ok,
		)
	}

	// Test PropertyName
	cf.SetPropertyName("prop1")
	if cf.PropertyName() != "prop1" {
		t.Errorf(
			"Expected PropertyName 'prop1', got '%s'",
			cf.PropertyName(),
		)
	}

	// Test Formula
	cf.SetFormula("=A1*B1")
	if cf.Formula() != "=A1*B1" {
		t.Errorf(
			"Expected Formula '=A1*B1', got '%s'",
			cf.Formula(),
		)
	}
}

func TestCacheFieldBooleanAttributes(
	t *testing.T,
) {
	cf := NewCacheField()

	// Test ServerField
	cf.SetServerField(true)
	if !cf.ServerField() {
		t.Error("Expected ServerField true")
	}

	// Test UniqueList (default true)
	if !cf.UniqueList() {
		t.Error(
			"Expected UniqueList true by default",
		)
	}
	cf.SetUniqueList(false)
	if cf.UniqueList() {
		t.Error("Expected UniqueList false")
	}

	// Test DatabaseField (default true)
	if !cf.DatabaseField() {
		t.Error(
			"Expected DatabaseField true by default",
		)
	}
	cf.SetDatabaseField(false)
	if cf.DatabaseField() {
		t.Error("Expected DatabaseField false")
	}

	// Test MemberPropertyField
	cf.SetMemberPropertyField(true)
	if !cf.MemberPropertyField() {
		t.Error(
			"Expected MemberPropertyField true",
		)
	}
}

func TestCacheFieldSharedItems(t *testing.T) {
	cf := NewCacheField()

	// Initially nil
	if cf.SharedItems() != nil {
		t.Error(
			"Expected SharedItems to be nil initially",
		)
	}

	// Create shared items
	si := cf.GetOrCreateSharedItems()
	if si == nil {
		t.Fatal(
			"GetOrCreateSharedItems returned nil",
		)
	}

	// Add a string item
	strItem := si.AddString("Category1")
	if strItem == nil {
		t.Fatal("AddString returned nil")
	}
	if strItem.V() != "Category1" {
		t.Errorf(
			"Expected V 'Category1', got '%s'",
			strItem.V(),
		)
	}

	// Add a number item
	numItem := si.AddNumber(123.45)
	if numItem == nil {
		t.Fatal("AddNumber returned nil")
	}
	if numItem.V() != 123.45 {
		t.Errorf(
			"Expected V 123.45, got %f",
			numItem.V(),
		)
	}

	// Add a boolean item
	boolItem := si.AddBoolean(true)
	if boolItem == nil {
		t.Fatal("AddBoolean returned nil")
	}
	if !boolItem.V() {
		t.Error("Expected V true")
	}

	// Add a missing item
	missingItem := si.AddMissing()
	if missingItem == nil {
		t.Fatal("AddMissing returned nil")
	}

	// Add an error item
	errItem := si.AddError("#N/A")
	if errItem == nil {
		t.Fatal("AddError returned nil")
	}
	if errItem.V() != "#N/A" {
		t.Errorf(
			"Expected V '#N/A', got '%s'",
			errItem.V(),
		)
	}

	// Add a datetime item
	dtItem := si.AddDateTime(
		"2023-12-20T10:30:00",
	)
	if dtItem == nil {
		t.Fatal("AddDateTime returned nil")
	}
	if dtItem.V() != "2023-12-20T10:30:00" {
		t.Errorf(
			"Expected V '2023-12-20T10:30:00', got '%s'",
			dtItem.V(),
		)
	}

	// Set count manually (it's not auto-incremented)
	si.SetCount(6)
	if si.Count() != 6 {
		t.Errorf(
			"Expected Count 6, got %d",
			si.Count(),
		)
	}
}

func TestSharedItemsAttributes(t *testing.T) {
	si := NewSharedItems()

	if si == nil {
		t.Fatal("NewSharedItems returned nil")
	}

	// Test ContainsSemiMixedTypes (default true)
	if !si.ContainsSemiMixedTypes() {
		t.Error(
			"Expected ContainsSemiMixedTypes true by default",
		)
	}
	si.SetContainsSemiMixedTypes(false)
	if si.ContainsSemiMixedTypes() {
		t.Error(
			"Expected ContainsSemiMixedTypes false",
		)
	}

	// Test ContainsNonDate (default true)
	if !si.ContainsNonDate() {
		t.Error(
			"Expected ContainsNonDate true by default",
		)
	}

	// Test ContainsDate
	si.SetContainsDate(true)
	if !si.ContainsDate() {
		t.Error("Expected ContainsDate true")
	}

	// Test ContainsString (default true)
	if !si.ContainsString() {
		t.Error(
			"Expected ContainsString true by default",
		)
	}

	// Test ContainsBlank
	si.SetContainsBlank(true)
	if !si.ContainsBlank() {
		t.Error("Expected ContainsBlank true")
	}

	// Test ContainsMixedTypes
	si.SetContainsMixedTypes(true)
	if !si.ContainsMixedTypes() {
		t.Error(
			"Expected ContainsMixedTypes true",
		)
	}

	// Test ContainsNumber
	si.SetContainsNumber(true)
	if !si.ContainsNumber() {
		t.Error("Expected ContainsNumber true")
	}

	// Test ContainsInteger
	si.SetContainsInteger(true)
	if !si.ContainsInteger() {
		t.Error("Expected ContainsInteger true")
	}

	// Test MinValue
	si.SetMinValue(10.5)
	minVal, ok := si.MinValue()
	if !ok || minVal != 10.5 {
		t.Errorf(
			"Expected MinValue 10.5, got %f (ok=%v)",
			minVal,
			ok,
		)
	}

	// Test MaxValue
	si.SetMaxValue(99.9)
	maxVal, ok := si.MaxValue()
	if !ok || maxVal != 99.9 {
		t.Errorf(
			"Expected MaxValue 99.9, got %f (ok=%v)",
			maxVal,
			ok,
		)
	}

	// Test LongText
	si.SetLongText(true)
	if !si.LongText() {
		t.Error("Expected LongText true")
	}
}

func TestCacheItemElements(t *testing.T) {
	// Test CacheItemString
	strItem := NewCacheItemString()
	if strItem.LocalName() != "s" {
		t.Errorf(
			"Expected LocalName 's', got '%s'",
			strItem.LocalName(),
		)
	}
	strItem.SetV("test")
	if strItem.V() != "test" {
		t.Errorf(
			"Expected V 'test', got '%s'",
			strItem.V(),
		)
	}

	// Test CacheItemNumber
	numItem := NewCacheItemNumber()
	if numItem.LocalName() != "n" {
		t.Errorf(
			"Expected LocalName 'n', got '%s'",
			numItem.LocalName(),
		)
	}
	numItem.SetV(42.5)
	if numItem.V() != 42.5 {
		t.Errorf(
			"Expected V 42.5, got %f",
			numItem.V(),
		)
	}

	// Test CacheItemBoolean
	boolItem := NewCacheItemBoolean()
	if boolItem.LocalName() != "b" {
		t.Errorf(
			"Expected LocalName 'b', got '%s'",
			boolItem.LocalName(),
		)
	}
	boolItem.SetV(true)
	if !boolItem.V() {
		t.Error("Expected V true")
	}

	// Test CacheItemMissing
	missingItem := NewCacheItemMissing()
	if missingItem.LocalName() != "m" {
		t.Errorf(
			"Expected LocalName 'm', got '%s'",
			missingItem.LocalName(),
		)
	}

	// Test CacheItemError
	errItem := NewCacheItemError()
	if errItem.LocalName() != "e" {
		t.Errorf(
			"Expected LocalName 'e', got '%s'",
			errItem.LocalName(),
		)
	}
	errItem.SetV("#REF!")
	if errItem.V() != "#REF!" {
		t.Errorf(
			"Expected V '#REF!', got '%s'",
			errItem.V(),
		)
	}

	// Test CacheItemDateTime
	dtItem := NewCacheItemDateTime()
	if dtItem.LocalName() != "d" {
		t.Errorf(
			"Expected LocalName 'd', got '%s'",
			dtItem.LocalName(),
		)
	}
	dtItem.SetV("2023-01-01T00:00:00")
	if dtItem.V() != "2023-01-01T00:00:00" {
		t.Errorf(
			"Expected V '2023-01-01T00:00:00', got '%s'",
			dtItem.V(),
		)
	}
}

func TestPivotCacheDefinitionClone(t *testing.T) {
	pcd := NewPivotCacheDefinition()
	pcd.SetRelationshipId("rId1")
	pcd.SetRecordCount(100)
	pcd.GetOrCreateCacheSource().
		SetType("worksheet")

	clonedResult := pcd.Clone()
	cloned, ok := clonedResult.(*PivotCacheDefinition)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *PivotCacheDefinition",
			clonedResult,
		)
	}

	if cloned == pcd {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.RelationshipId() != "rId1" {
		t.Errorf(
			"Cloned RelationshipId should be 'rId1', got '%s'",
			cloned.RelationshipId(),
		)
	}
}

func TestCacheFieldsIteration(t *testing.T) {
	cfs := NewCacheFields()

	// Add fields using AddField
	cfs.AddField("Field1")
	cfs.AddField("Field2")
	cfs.AddField("Field3")

	// Verify count using FieldCount
	if cfs.FieldCount() != 3 {
		t.Errorf(
			"Expected FieldCount 3, got %d",
			cfs.FieldCount(),
		)
	}

	// Iterate using CacheFields()
	count := 0
	for cf := range cfs.CacheFields() {
		count++
		if cf == nil {
			t.Error("Iterator yielded nil field")
		}
	}
	if count != 3 {
		t.Errorf(
			"Expected 3 fields in iterator, got %d",
			count,
		)
	}

	// Get by index
	cf := cfs.GetField(1)
	if cf == nil {
		t.Fatal("GetField(1) returned nil")
	}
	if cf.Name() != "Field2" {
		t.Errorf(
			"Expected Name 'Field2', got '%s'",
			cf.Name(),
		)
	}

	// Out of range
	if cfs.GetField(99) != nil {
		t.Error(
			"Expected nil for out of range index",
		)
	}
}
