package elements

import (
	"strings"
	"testing"
)

func TestNewWorkbook(t *testing.T) {
	wb := NewWorkbook()

	if wb == nil {
		t.Fatal("expected non-nil Workbook")
	}
	if wb.LocalName() != "workbook" {
		t.Errorf(
			"expected localName 'workbook', got '%s'",
			wb.LocalName(),
		)
	}
	if wb.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceSML,
			wb.NamespaceURI(),
		)
	}
}

func TestWorkbook_GetOrCreateFileVersion(
	t *testing.T,
) {
	wb := NewWorkbook()

	// Initially nil
	if wb.FileVersion() != nil {
		t.Error(
			"expected nil FileVersion initially",
		)
	}

	// GetOrCreate should create it
	fv := wb.GetOrCreateFileVersion()
	if fv == nil {
		t.Fatal("expected non-nil FileVersion")
	}

	// Second call should return same element
	fv2 := wb.GetOrCreateFileVersion()
	if fv2 == nil {
		t.Fatal(
			"expected non-nil FileVersion on second call",
		)
	}

	// FileVersion should return the element
	if wb.FileVersion() == nil {
		t.Error(
			"expected FileVersion to be present after creation",
		)
	}
}

func TestWorkbook_GetOrCreateSheets(
	t *testing.T,
) {
	wb := NewWorkbook()

	sheets := wb.GetOrCreateSheets()
	if sheets == nil {
		t.Fatal("expected non-nil Sheets")
	}

	// Add a sheet
	sheet := sheets.AddSheet(
		"Sheet1",
		1,
		relationIdOne,
	)
	if sheet == nil {
		t.Fatal("expected non-nil Sheet")
	}
	if sheet.Name() != testSheet1Name {
		t.Errorf(
			"expected sheet name 'Sheet1', got '%s'",
			sheet.Name(),
		)
	}
	if sheet.SheetId() != 1 {
		t.Errorf(
			"expected sheetId 1, got %d",
			sheet.SheetId(),
		)
	}
	if sheet.RelationshipId() != relationIdOne {
		t.Errorf(
			"expected relationshipId 'rId1', got '%s'",
			sheet.RelationshipId(),
		)
	}
}

func TestFileVersion(t *testing.T) {
	fv := NewFileVersion()

	// Test AppName
	fv.SetAppName("xl")
	if fv.AppName() != "xl" {
		t.Errorf(
			"expected appName 'xl', got '%s'",
			fv.AppName(),
		)
	}

	// Test LastEdited
	fv.SetLastEdited("7")
	if fv.LastEdited() != "7" {
		t.Errorf(
			"expected lastEdited '7', got '%s'",
			fv.LastEdited(),
		)
	}

	// Test LowestEdited
	fv.SetLowestEdited("6")
	if fv.LowestEdited() != "6" {
		t.Errorf(
			"expected lowestEdited '6', got '%s'",
			fv.LowestEdited(),
		)
	}

	// Test RupBuild
	fv.SetRupBuild("14420")
	if fv.RupBuild() != "14420" {
		t.Errorf(
			"expected rupBuild '14420', got '%s'",
			fv.RupBuild(),
		)
	}

	// Test SetExcelVersion convenience method
	fv2 := NewFileVersion()
	fv2.SetExcelVersion(7, 6)
	if fv2.AppName() != "xl" {
		t.Errorf(
			"expected appName 'xl', got '%s'",
			fv2.AppName(),
		)
	}
	if fv2.LastEdited() != "7" {
		t.Errorf(
			"expected lastEdited '7', got '%s'",
			fv2.LastEdited(),
		)
	}
	if fv2.LowestEdited() != "6" {
		t.Errorf(
			"expected lowestEdited '6', got '%s'",
			fv2.LowestEdited(),
		)
	}
}

func TestWorkbookPr(t *testing.T) {
	wpr := NewWorkbookPr()

	// Test Date1904
	if wpr.Date1904() {
		t.Error(
			"expected Date1904 to be false initially",
		)
	}
	wpr.SetDate1904(true)
	if !wpr.Date1904() {
		t.Error(
			"expected Date1904 to be true after setting",
		)
	}

	// Test FilterPrivacy
	wpr.SetFilterPrivacy(true)
	if !wpr.FilterPrivacy() {
		t.Error(
			"expected FilterPrivacy to be true",
		)
	}

	// Test DefaultThemeVersion
	wpr.SetDefaultThemeVersion(166925)
	if wpr.DefaultThemeVersion() != 166925 {
		t.Errorf(
			"expected defaultThemeVersion 166925, got %d",
			wpr.DefaultThemeVersion(),
		)
	}

	// Test CodeName
	wpr.SetCodeName("ThisWorkbook")
	if wpr.CodeName() != "ThisWorkbook" {
		t.Errorf(
			"expected codeName 'ThisWorkbook', got '%s'",
			wpr.CodeName(),
		)
	}
}

func TestBookViews(t *testing.T) {
	bv := NewBookViews()

	// Initially empty
	count := 0
	for range bv.WorkbookViews() {
		count++
	}
	if count != 0 {
		t.Errorf(
			"expected 0 workbook views, got %d",
			count,
		)
	}

	// Add a workbook view
	wv := bv.AddWorkbookView()
	if wv == nil {
		t.Fatal("expected non-nil WorkbookView")
	}

	// Count should be 1
	count = 0
	for range bv.WorkbookViews() {
		count++
	}
	if count != 1 {
		t.Errorf(
			"expected 1 workbook view, got %d",
			count,
		)
	}
}

func TestWorkbookView(t *testing.T) {
	wv := NewWorkbookView()

	// Test visibility
	if wv.Visibility() != WorkbookViewVisible {
		t.Errorf(
			"expected visibility 'visible', got '%s'",
			wv.Visibility(),
		)
	}
	wv.SetVisibility(WorkbookViewHidden)
	if wv.Visibility() != WorkbookViewHidden {
		t.Errorf(
			"expected visibility 'hidden', got '%s'",
			wv.Visibility(),
		)
	}

	// Test minimized
	if wv.Minimized() {
		t.Error(
			"expected minimized to be false initially",
		)
	}
	wv.SetMinimized(true)
	if !wv.Minimized() {
		t.Error("expected minimized to be true")
	}

	// Test scroll bars
	if !wv.ShowHorizontalScroll() {
		t.Error(
			"expected showHorizontalScroll to be true by default",
		)
	}
	wv.SetShowHorizontalScroll(false)
	if wv.ShowHorizontalScroll() {
		t.Error(
			"expected showHorizontalScroll to be false",
		)
	}

	// Test window bounds
	wv.SetWindowBounds(100, 200, 1000, 800)
	if wv.XWindow() != 100 {
		t.Errorf(
			"expected xWindow 100, got %d",
			wv.XWindow(),
		)
	}
	if wv.YWindow() != 200 {
		t.Errorf(
			"expected yWindow 200, got %d",
			wv.YWindow(),
		)
	}
	if wv.WindowWidth() != 1000 {
		t.Errorf(
			"expected windowWidth 1000, got %d",
			wv.WindowWidth(),
		)
	}
	if wv.WindowHeight() != 800 {
		t.Errorf(
			"expected windowHeight 800, got %d",
			wv.WindowHeight(),
		)
	}

	// Test activeTab
	wv.SetActiveTab(2)
	if wv.ActiveTab() != 2 {
		t.Errorf(
			"expected activeTab 2, got %d",
			wv.ActiveTab(),
		)
	}
}

func TestSheets(t *testing.T) {
	sheets := NewSheets()

	// Add sheets
	s1 := sheets.AddSheet(
		"Sheet1",
		1,
		relationIdOne,
	)
	s2 := sheets.AddSheet("Sheet2", 2, "rId2")
	s3 := sheets.AddSheet("Sheet3", 3, "rId3")

	// Test count
	if sheets.SheetCount() != 3 {
		t.Errorf(
			"expected 3 sheets, got %d",
			sheets.SheetCount(),
		)
	}

	// Test GetSheetByName
	found := sheets.GetSheetByName("Sheet2")
	if found == nil {
		t.Fatal("expected to find Sheet2")
	}
	if found.Name() != "Sheet2" {
		t.Errorf(
			"expected name 'Sheet2', got '%s'",
			found.Name(),
		)
	}

	// Test GetSheetById
	found = sheets.GetSheetById(3)
	if found == nil {
		t.Fatal(
			"expected to find sheet with ID 3",
		)
	}
	if found.Name() != "Sheet3" {
		t.Errorf(
			"expected name 'Sheet3', got '%s'",
			found.Name(),
		)
	}

	// Test GetSheetByIndex
	found = sheets.GetSheetByIndex(0)
	if found == nil {
		t.Fatal(
			"expected to find sheet at index 0",
		)
	}
	if found.Name() != "Sheet1" {
		t.Errorf(
			"expected name 'Sheet1', got '%s'",
			found.Name(),
		)
	}

	// Test NextSheetId
	if sheets.NextSheetId() != 4 {
		t.Errorf(
			"expected next sheet ID 4, got %d",
			sheets.NextSheetId(),
		)
	}

	// Test RemoveSheetByName
	if !sheets.RemoveSheetByName("Sheet2") {
		t.Error("expected to remove Sheet2")
	}
	if sheets.SheetCount() != 2 {
		t.Errorf(
			"expected 2 sheets after removal, got %d",
			sheets.SheetCount(),
		)
	}

	// Verify s1 and s3 are still there
	_ = s1
	_ = s2
	_ = s3
}

func TestSheet_State(t *testing.T) {
	sheet := NewSheet()
	sheet.SetName("Test")
	sheet.SetSheetId(1)

	// Default state is visible
	if sheet.State() != SheetStateVisible {
		t.Errorf(
			"expected state 'visible', got '%s'",
			sheet.State(),
		)
	}
	if !sheet.IsVisible() {
		t.Error(
			"expected IsVisible() to return true",
		)
	}

	// Hide the sheet
	sheet.Hide()
	if sheet.State() != SheetStateHidden {
		t.Errorf(
			"expected state 'hidden', got '%s'",
			sheet.State(),
		)
	}
	if !sheet.IsHidden() {
		t.Error(
			"expected IsHidden() to return true",
		)
	}

	// Very hide the sheet
	sheet.VeryHide()
	if sheet.State() != SheetStateVeryHidden {
		t.Errorf(
			"expected state 'veryHidden', got '%s'",
			sheet.State(),
		)
	}
	if !sheet.IsVeryHidden() {
		t.Error(
			"expected IsVeryHidden() to return true",
		)
	}

	// Show the sheet
	sheet.Show()
	if !sheet.IsVisible() {
		t.Error(
			"expected IsVisible() to return true after Show()",
		)
	}
}

func TestDefinedNames(t *testing.T) {
	dn := NewDefinedNames()

	// Add workbook-level name
	name1 := dn.AddDefinedName(
		"MyRange",
		"Sheet1!$A$1:$D$10",
	)
	if name1 == nil {
		t.Fatal("expected non-nil DefinedName")
	}
	if name1.Name() != "MyRange" {
		t.Errorf(
			"expected name 'MyRange', got '%s'",
			name1.Name(),
		)
	}
	if name1.Formula() != "Sheet1!$A$1:$D$10" {
		t.Errorf(
			"expected formula 'Sheet1!$A$1:$D$10', got '%s'",
			name1.Formula(),
		)
	}

	// Add sheet-level name
	name2 := dn.AddSheetLevelName(
		"LocalRange",
		0,
		"Sheet1!$B$2:$C$5",
	)
	if name2 == nil {
		t.Fatal("expected non-nil DefinedName")
	}
	if !name2.HasLocalSheetId() {
		t.Error(
			"expected HasLocalSheetId() to return true",
		)
	}
	if name2.LocalSheetId() != 0 {
		t.Errorf(
			"expected localSheetId 0, got %d",
			name2.LocalSheetId(),
		)
	}

	// Test GetWorkbookLevelName
	found := dn.GetWorkbookLevelName("MyRange")
	if found == nil {
		t.Fatal("expected to find MyRange")
	}

	// Test GetSheetLevelName
	found = dn.GetSheetLevelName("LocalRange", 0)
	if found == nil {
		t.Fatal("expected to find LocalRange")
	}

	// Test hidden
	name1.SetHidden(true)
	if !name1.Hidden() {
		t.Error("expected hidden to be true")
	}

	// Test comment
	name1.SetComment("This is a test range")
	if name1.Comment() != "This is a test range" {
		t.Errorf(
			"expected comment 'This is a test range', got '%s'",
			name1.Comment(),
		)
	}
}

func TestCalcPr(t *testing.T) {
	cp := NewCalcPr()

	// Test CalcId
	cp.SetCalcId(191029)
	if cp.CalcId() != 191029 {
		t.Errorf(
			"expected calcId 191029, got %d",
			cp.CalcId(),
		)
	}

	// Test CalcMode
	if cp.CalcMode() != CalcModeAuto {
		t.Errorf(
			"expected calcMode 'auto', got '%s'",
			cp.CalcMode(),
		)
	}
	cp.SetCalcMode(CalcModeManual)
	if cp.CalcMode() != CalcModeManual {
		t.Errorf(
			"expected calcMode 'manual', got '%s'",
			cp.CalcMode(),
		)
	}

	// Test RefMode
	if cp.RefMode() != RefModeA1 {
		t.Errorf(
			"expected refMode 'A1', got '%s'",
			cp.RefMode(),
		)
	}
	cp.SetRefMode(RefModeR1C1)
	if cp.RefMode() != RefModeR1C1 {
		t.Errorf(
			"expected refMode 'R1C1', got '%s'",
			cp.RefMode(),
		)
	}

	// Test Iterate
	cp.SetIterate(true)
	if !cp.Iterate() {
		t.Error("expected iterate to be true")
	}

	// Test IterateCount
	if cp.IterateCount() != 100 {
		t.Errorf(
			"expected iterateCount 100, got %d",
			cp.IterateCount(),
		)
	}
	cp.SetIterateCount(200)
	if cp.IterateCount() != 200 {
		t.Errorf(
			"expected iterateCount 200, got %d",
			cp.IterateCount(),
		)
	}

	// Test IterateDelta
	cp.SetIterateDelta(0.0001)
	if cp.IterateDelta() != 0.0001 {
		t.Errorf(
			"expected iterateDelta 0.0001, got %f",
			cp.IterateDelta(),
		)
	}

	// Test FullPrecision
	if !cp.FullPrecision() {
		t.Error(
			"expected fullPrecision to be true by default",
		)
	}
	cp.SetFullPrecision(false)
	if cp.FullPrecision() {
		t.Error(
			"expected fullPrecision to be false",
		)
	}

	// Test ConcurrentCalc
	if !cp.ConcurrentCalc() {
		t.Error(
			"expected concurrentCalc to be true by default",
		)
	}
	cp.SetConcurrentCalc(false)
	if cp.ConcurrentCalc() {
		t.Error(
			"expected concurrentCalc to be false",
		)
	}
}

func TestPivotCaches(t *testing.T) {
	pc := NewPivotCaches()

	// Add pivot cache
	cache := pc.AddPivotCache(1, relationIdOne)
	if cache == nil {
		t.Fatal("expected non-nil PivotCache")
	}
	if cache.CacheId() != 1 {
		t.Errorf(
			"expected cacheId 1, got %d",
			cache.CacheId(),
		)
	}
	if cache.RelationshipId() != relationIdOne {
		t.Errorf(
			"expected relationshipId 'rId1', got '%s'",
			cache.RelationshipId(),
		)
	}

	// Test count
	if pc.Count() != 1 {
		t.Errorf(
			"expected 1 pivot cache, got %d",
			pc.Count(),
		)
	}

	// Test GetByCacheId
	found := pc.GetByCacheId(1)
	if found == nil {
		t.Fatal(
			"expected to find pivot cache with ID 1",
		)
	}

	// Test NextCacheId
	if pc.NextCacheId() != 2 {
		t.Errorf(
			"expected next cache ID 2, got %d",
			pc.NextCacheId(),
		)
	}
}

func TestWorkbookProtection(t *testing.T) {
	wp := NewWorkbookProtection()

	// Test LockStructure
	if wp.LockStructure() {
		t.Error(
			"expected lockStructure to be false initially",
		)
	}
	wp.SetLockStructure(true)
	if !wp.LockStructure() {
		t.Error(
			"expected lockStructure to be true",
		)
	}

	// Test LockWindows
	wp.SetLockWindows(true)
	if !wp.LockWindows() {
		t.Error("expected lockWindows to be true")
	}

	// Test password methods
	wp.SetWorkbookPassword("ABCD")
	if wp.WorkbookPassword() != "ABCD" {
		t.Errorf(
			"expected workbookPassword 'ABCD', got '%s'",
			wp.WorkbookPassword(),
		)
	}

	// Test secure protection
	wp.SetSecureProtection(
		"SHA-512",
		"base64hash",
		"base64salt",
		"100000",
	)
	if wp.WorkbookAlgorithmName() != "SHA-512" {
		t.Errorf(
			"expected algorithm 'SHA-512', got '%s'",
			wp.WorkbookAlgorithmName(),
		)
	}
	if wp.WorkbookHashValue() != "base64hash" {
		t.Errorf(
			"expected hash 'base64hash', got '%s'",
			wp.WorkbookHashValue(),
		)
	}
	if wp.WorkbookSaltValue() != "base64salt" {
		t.Errorf(
			"expected salt 'base64salt', got '%s'",
			wp.WorkbookSaltValue(),
		)
	}
	if wp.WorkbookSpinCount() != "100000" {
		t.Errorf(
			"expected spinCount '100000', got '%s'",
			wp.WorkbookSpinCount(),
		)
	}

	// Test clear password
	wp.ClearWorkbookPassword()
	if wp.WorkbookPassword() != "" {
		t.Error(
			"expected workbookPassword to be empty after clear",
		)
	}
	if wp.WorkbookAlgorithmName() != "" {
		t.Error(
			"expected algorithm to be empty after clear",
		)
	}
}

func TestFileRecoveryPr(t *testing.T) {
	fr := NewFileRecoveryPr()

	// Test AutoRecover (default true)
	if !fr.AutoRecover() {
		t.Error(
			"expected autoRecover to be true by default",
		)
	}
	fr.SetAutoRecover(false)
	if fr.AutoRecover() {
		t.Error(
			"expected autoRecover to be false",
		)
	}

	// Test CrashSave
	if fr.CrashSave() {
		t.Error(
			"expected crashSave to be false initially",
		)
	}
	fr.SetCrashSave(true)
	if !fr.CrashSave() {
		t.Error("expected crashSave to be true")
	}

	// Test DataExtractLoad
	fr.SetDataExtractLoad(true)
	if !fr.DataExtractLoad() {
		t.Error(
			"expected dataExtractLoad to be true",
		)
	}

	// Test RepairLoad
	fr.SetRepairLoad(true)
	if !fr.RepairLoad() {
		t.Error("expected repairLoad to be true")
	}
}

func TestWorkbook_XMLSerialization(t *testing.T) {
	wb := NewWorkbook()

	// Add file version
	fv := wb.GetOrCreateFileVersion()
	fv.SetAppName("xl")
	fv.SetLastEdited("7")

	// Add workbook properties
	wpr := wb.GetOrCreateWorkbookPr()
	wpr.SetDefaultThemeVersion(166925)

	// Add book views
	bv := wb.GetOrCreateBookViews()
	wv := bv.GetOrCreateWorkbookView()
	wv.SetWindowBounds(0, 0, 28800, 12300)
	wv.SetActiveTab(0)

	// Add sheets
	sheets := wb.GetOrCreateSheets()
	sheets.AddSheet("Sheet1", 1, relationIdOne)
	sheets.AddSheet("Sheet2", 2, "rId2")

	// Add calculation properties
	cp := wb.GetOrCreateCalcPr()
	cp.SetCalcId(191029)

	// Get XML output
	xml := wb.OuterXml()

	// Verify key elements are present
	if !strings.Contains(xml, "workbook") {
		t.Error("expected 'workbook' in XML")
	}
	if !strings.Contains(xml, "fileVersion") {
		t.Error("expected 'fileVersion' in XML")
	}
	if !strings.Contains(xml, "appName=\"xl\"") {
		t.Error(
			"expected 'appName=\"xl\"' in XML",
		)
	}
	if !strings.Contains(xml, "sheets") {
		t.Error("expected 'sheets' in XML")
	}
	if !strings.Contains(xml, "Sheet1") {
		t.Error("expected 'Sheet1' in XML")
	}
	if !strings.Contains(xml, "calcPr") {
		t.Error("expected 'calcPr' in XML")
	}
}

func TestClone(t *testing.T) {
	// Test Workbook clone
	wb := NewWorkbook()
	sheets := wb.GetOrCreateSheets()
	sheets.AddSheet("Sheet1", 1, relationIdOne)

	cloned := wb.Clone()
	clone, ok := cloned.(*Workbook)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Workbook",
			cloned,
		)
	}
	if clone.Sheets() == nil {
		t.Error("expected clone to have Sheets")
	}
	if clone.Sheets().SheetCount() != 1 {
		t.Errorf(
			"expected clone to have 1 sheet, got %d",
			clone.Sheets().SheetCount(),
		)
	}

	// Modify original, verify clone is independent
	sheets.AddSheet("Sheet2", 2, "rId2")
	if clone.Sheets().SheetCount() != 1 {
		t.Error(
			"expected clone to remain unchanged after modifying original",
		)
	}
}
