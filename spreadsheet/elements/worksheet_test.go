package elements

import (
	"strings"
	"testing"
)

func TestWorksheet_SheetPr(t *testing.T) {
	ws := NewWorksheet()

	// Initially nil
	if ws.SheetPr() != nil {
		t.Error("expected nil SheetPr initially")
	}

	// GetOrCreate should create it
	sp := ws.GetOrCreateSheetPr()
	if sp == nil {
		t.Fatal("expected non-nil SheetPr")
	}

	// Second call should return same element
	sp2 := ws.GetOrCreateSheetPr()
	if sp2 == nil {
		t.Fatal(
			"expected non-nil SheetPr on second call",
		)
	}
}

func TestWorksheet_Dimension(t *testing.T) {
	ws := NewWorksheet()

	dim := ws.GetOrCreateDimension()
	if dim == nil {
		t.Fatal("expected non-nil Dimension")
	}

	dim.SetRef("A1:D10")
	if dim.Ref() != "A1:D10" {
		t.Errorf(
			"expected ref 'A1:D10', got '%s'",
			dim.Ref(),
		)
	}
}

func TestWorksheet_SheetViews(t *testing.T) {
	ws := NewWorksheet()

	sv := ws.GetOrCreateSheetViews()
	if sv == nil {
		t.Fatal("expected non-nil SheetViews")
	}

	view := sv.GetOrCreateSheetView()
	if view == nil {
		t.Fatal("expected non-nil SheetView")
	}

	view.SetTabSelected(true)
	if !view.TabSelected() {
		t.Error("expected tabSelected to be true")
	}
}

func TestSheetPr(t *testing.T) {
	sp := NewSheetPr()

	// Test CodeName
	sp.SetCodeName("Sheet1")
	if sp.CodeName() != "Sheet1" {
		t.Errorf(
			"expected codeName 'Sheet1', got '%s'",
			sp.CodeName(),
		)
	}

	// Test FilterMode
	if sp.FilterMode() {
		t.Error(
			"expected filterMode to be false initially",
		)
	}
	sp.SetFilterMode(true)
	if !sp.FilterMode() {
		t.Error("expected filterMode to be true")
	}

	// Test Published (default true)
	if !sp.Published() {
		t.Error(
			"expected published to be true by default",
		)
	}
	sp.SetPublished(false)
	if sp.Published() {
		t.Error("expected published to be false")
	}
}

func TestSheetPr_TabColor(t *testing.T) {
	sp := NewSheetPr()

	tc := sp.GetOrCreateTabColor()
	if tc == nil {
		t.Fatal("expected non-nil TabColor")
	}

	tc.SetRGB("FF0000")
	if tc.RGB() != "FF0000" {
		t.Errorf(
			"expected RGB 'FF0000', got '%s'",
			tc.RGB(),
		)
	}

	tc.SetTheme("5")
	if tc.Theme() != "5" {
		t.Errorf(
			"expected theme '5', got '%s'",
			tc.Theme(),
		)
	}
}

func TestSheetView(t *testing.T) {
	view := NewSheetView()

	// Test ShowGridLines (default true)
	if !view.ShowGridLines() {
		t.Error(
			"expected showGridLines to be true by default",
		)
	}
	view.SetShowGridLines(false)
	if view.ShowGridLines() {
		t.Error(
			"expected showGridLines to be false",
		)
	}

	// Test ShowFormulas
	if view.ShowFormulas() {
		t.Error(
			"expected showFormulas to be false by default",
		)
	}
	view.SetShowFormulas(true)
	if !view.ShowFormulas() {
		t.Error(
			"expected showFormulas to be true",
		)
	}

	// Test ZoomScale
	if view.ZoomScale() != 100 {
		t.Errorf(
			"expected zoomScale 100, got %d",
			view.ZoomScale(),
		)
	}
	view.SetZoomScale(150)
	if view.ZoomScale() != 150 {
		t.Errorf(
			"expected zoomScale 150, got %d",
			view.ZoomScale(),
		)
	}

	// Test View type
	if view.View() != SheetViewNormal {
		t.Errorf(
			"expected view 'normal', got '%s'",
			view.View(),
		)
	}
	view.SetView(SheetViewPageLayout)
	if view.View() != SheetViewPageLayout {
		t.Errorf(
			"expected view 'pageLayout', got '%s'",
			view.View(),
		)
	}
}

func TestPane(t *testing.T) {
	pane := NewPane()

	// Test freeze panes
	pane.FreezePanes(2, 3)
	if pane.State() != PaneStateFrozen {
		t.Errorf(
			"expected state 'frozen', got '%s'",
			pane.State(),
		)
	}
	if pane.XSplit() != 3 {
		t.Errorf(
			"expected xSplit 3, got %f",
			pane.XSplit(),
		)
	}
	if pane.YSplit() != 2 {
		t.Errorf(
			"expected ySplit 2, got %f",
			pane.YSplit(),
		)
	}
	if pane.ActivePane() != PanePositionBottomRight {
		t.Errorf(
			"expected activePane 'bottomRight', got '%s'",
			pane.ActivePane(),
		)
	}

	// Test TopLeftCell
	pane.SetTopLeftCell("D3")
	if pane.TopLeftCell() != "D3" {
		t.Errorf(
			"expected topLeftCell 'D3', got '%s'",
			pane.TopLeftCell(),
		)
	}
}

func TestSelection(t *testing.T) {
	sel := NewSelection()

	sel.SetSelection("B2", "B2:E5")
	if sel.ActiveCell() != "B2" {
		t.Errorf(
			"expected activeCell 'B2', got '%s'",
			sel.ActiveCell(),
		)
	}
	if sel.Sqref() != "B2:E5" {
		t.Errorf(
			"expected sqref 'B2:E5', got '%s'",
			sel.Sqref(),
		)
	}

	sel.SetPane(PanePositionBottomLeft)
	if sel.Pane() != PanePositionBottomLeft {
		t.Errorf(
			"expected pane 'bottomLeft', got '%s'",
			sel.Pane(),
		)
	}
}

func TestSheetFormatPr(t *testing.T) {
	sfp := NewSheetFormatPr()

	// Test default values
	if sfp.BaseColWidth() != 8 {
		t.Errorf(
			"expected baseColWidth 8, got %d",
			sfp.BaseColWidth(),
		)
	}
	if sfp.DefaultRowHeight() != 15.0 {
		t.Errorf(
			"expected defaultRowHeight 15.0, got %f",
			sfp.DefaultRowHeight(),
		)
	}

	// Test setters
	sfp.SetDefaultColWidth(10.5)
	if sfp.DefaultColWidth() != 10.5 {
		t.Errorf(
			"expected defaultColWidth 10.5, got %f",
			sfp.DefaultColWidth(),
		)
	}

	sfp.SetDefaultRowHeight(20.0)
	if sfp.DefaultRowHeight() != 20.0 {
		t.Errorf(
			"expected defaultRowHeight 20.0, got %f",
			sfp.DefaultRowHeight(),
		)
	}

	sfp.SetOutlineLevelRow(3)
	if sfp.OutlineLevelRow() != 3 {
		t.Errorf(
			"expected outlineLevelRow 3, got %d",
			sfp.OutlineLevelRow(),
		)
	}
}

func TestCols(t *testing.T) {
	cols := NewCols()

	// Add column definition
	col := cols.AddCol(1, 5)
	if col == nil {
		t.Fatal("expected non-nil Col")
	}
	if col.Min() != 1 {
		t.Errorf(
			"expected min 1, got %d",
			col.Min(),
		)
	}
	if col.Max() != 5 {
		t.Errorf(
			"expected max 5, got %d",
			col.Max(),
		)
	}

	// Test width
	col.SetWidth(15.5)
	col.SetCustomWidth(true)
	if col.Width() != 15.5 {
		t.Errorf(
			"expected width 15.5, got %f",
			col.Width(),
		)
	}
	if !col.CustomWidth() {
		t.Error("expected customWidth to be true")
	}

	// Test hidden
	col.SetHidden(true)
	if !col.Hidden() {
		t.Error("expected hidden to be true")
	}

	// Test GetColByIndex
	found := cols.GetColByIndex(3)
	if found == nil {
		t.Fatal(
			"expected to find column at index 3",
		)
	}
}

func TestSheetProtection(t *testing.T) {
	sp := NewSheetProtection()

	// Test Sheet
	if sp.Sheet() {
		t.Error(
			"expected sheet to be false initially",
		)
	}
	sp.SetSheet(true)
	if !sp.Sheet() {
		t.Error("expected sheet to be true")
	}

	// Test FormatCells (default true)
	if !sp.FormatCells() {
		t.Error(
			"expected formatCells to be true by default",
		)
	}
	sp.SetFormatCells(false)
	if sp.FormatCells() {
		t.Error(
			"expected formatCells to be false",
		)
	}

	// Test secure protection
	sp.SetSecureProtection(
		"SHA-512",
		"hash",
		"salt",
		"100000",
	)
	if sp.AlgorithmName() != "SHA-512" {
		t.Errorf(
			"expected algorithm 'SHA-512', got '%s'",
			sp.AlgorithmName(),
		)
	}
	if sp.HashValue() != "hash" {
		t.Errorf(
			"expected hash 'hash', got '%s'",
			sp.HashValue(),
		)
	}

	// Test clear password
	sp.ClearPassword()
	if sp.AlgorithmName() != "" {
		t.Error(
			"expected algorithm to be empty after clear",
		)
	}
}

func TestPageMargins(t *testing.T) {
	pm := NewPageMargins()

	// Test defaults
	if pm.Left() != 0.7 {
		t.Errorf(
			"expected left 0.7, got %f",
			pm.Left(),
		)
	}
	if pm.Right() != 0.7 {
		t.Errorf(
			"expected right 0.7, got %f",
			pm.Right(),
		)
	}

	// Test setters
	pm.SetAllMargins(1.0, 1.0, 1.5, 1.5, 0.5, 0.5)
	if pm.Left() != 1.0 {
		t.Errorf(
			"expected left 1.0, got %f",
			pm.Left(),
		)
	}
	if pm.Top() != 1.5 {
		t.Errorf(
			"expected top 1.5, got %f",
			pm.Top(),
		)
	}
	if pm.Header() != 0.5 {
		t.Errorf(
			"expected header 0.5, got %f",
			pm.Header(),
		)
	}
}

func TestPageSetup(t *testing.T) {
	ps := NewPageSetup()

	// Test defaults
	if ps.PaperSize() != 1 {
		t.Errorf(
			"expected paperSize 1, got %d",
			ps.PaperSize(),
		)
	}
	if ps.Scale() != 100 {
		t.Errorf(
			"expected scale 100, got %d",
			ps.Scale(),
		)
	}

	// Test orientation
	if ps.Orientation() != OrientationDefault {
		t.Errorf(
			"expected orientation 'default', got '%s'",
			ps.Orientation(),
		)
	}
	ps.SetOrientation(OrientationLandscape)
	if ps.Orientation() != OrientationLandscape {
		t.Errorf(
			"expected orientation 'landscape', got '%s'",
			ps.Orientation(),
		)
	}

	// Test FitToPage
	ps.SetFitToWidth(2)
	ps.SetFitToHeight(3)
	if ps.FitToWidth() != 2 {
		t.Errorf(
			"expected fitToWidth 2, got %d",
			ps.FitToWidth(),
		)
	}
	if ps.FitToHeight() != 3 {
		t.Errorf(
			"expected fitToHeight 3, got %d",
			ps.FitToHeight(),
		)
	}
}

func TestHeaderFooter(t *testing.T) {
	hf := NewHeaderFooter()

	// Test header/footer
	hf.SetHeader("&LLeft&CCenter&RRight")
	if hf.OddHeader() != "&LLeft&CCenter&RRight" {
		t.Errorf(
			"expected header '&LLeft&CCenter&RRight', got '%s'",
			hf.OddHeader(),
		)
	}

	hf.SetFooter("&P of &N")
	if hf.OddFooter() != "&P of &N" {
		t.Errorf(
			"expected footer '&P of &N', got '%s'",
			hf.OddFooter(),
		)
	}

	// Test DifferentFirst
	hf.SetDifferentFirst(true)
	if !hf.DifferentFirst() {
		t.Error(
			"expected differentFirst to be true",
		)
	}

	hf.SetFirstHeader("First Page Header")
	if hf.FirstHeader() != "First Page Header" {
		t.Errorf(
			"expected firstHeader 'First Page Header', got '%s'",
			hf.FirstHeader(),
		)
	}
}

func TestPrintOptions(t *testing.T) {
	po := NewPrintOptions()

	// Test GridLines
	if po.GridLines() {
		t.Error(
			"expected gridLines to be false by default",
		)
	}
	po.SetGridLines(true)
	if !po.GridLines() {
		t.Error("expected gridLines to be true")
	}

	// Test Headings
	po.SetHeadings(true)
	if !po.Headings() {
		t.Error("expected headings to be true")
	}

	// Test centering
	po.SetHorizontalCentered(true)
	po.SetVerticalCentered(true)
	if !po.HorizontalCentered() {
		t.Error(
			"expected horizontalCentered to be true",
		)
	}
	if !po.VerticalCentered() {
		t.Error(
			"expected verticalCentered to be true",
		)
	}
}

func TestPageBreaks(t *testing.T) {
	rb := NewRowBreaks()

	// Add break
	brk := rb.AddBreak(10)
	if brk == nil {
		t.Fatal("expected non-nil Break")
	}
	if brk.Id() != 10 {
		t.Errorf(
			"expected id 10, got %d",
			brk.Id(),
		)
	}
	if !brk.Man() {
		t.Error("expected man to be true")
	}

	// Test counts
	if rb.Count() != 1 {
		t.Errorf(
			"expected count 1, got %d",
			rb.Count(),
		)
	}
	if rb.ManualBreakCount() != 1 {
		t.Errorf(
			"expected manualBreakCount 1, got %d",
			rb.ManualBreakCount(),
		)
	}

	// Test ColBreaks
	cb := NewColBreaks()
	cb.AddBreak(5)
	if cb.Count() != 1 {
		t.Errorf(
			"expected count 1, got %d",
			cb.Count(),
		)
	}
}

func TestDrawingRef(t *testing.T) {
	d := NewDrawing()
	d.SetRelationshipId("rId1")
	if d.RelationshipId() != "rId1" {
		t.Errorf(
			"expected relationshipId 'rId1', got '%s'",
			d.RelationshipId(),
		)
	}
}

func TestLegacyDrawing(t *testing.T) {
	ld := NewLegacyDrawing()
	ld.SetRelationshipId("rId2")
	if ld.RelationshipId() != "rId2" {
		t.Errorf(
			"expected relationshipId 'rId2', got '%s'",
			ld.RelationshipId(),
		)
	}
}

func TestPicture(t *testing.T) {
	p := NewPicture()
	p.SetRelationshipId("rId3")
	if p.RelationshipId() != "rId3" {
		t.Errorf(
			"expected relationshipId 'rId3', got '%s'",
			p.RelationshipId(),
		)
	}
}

func TestWorksheet_XMLSerialization(
	t *testing.T,
) {
	ws := NewWorksheet()

	// Add dimension
	dim := ws.GetOrCreateDimension()
	dim.SetRef("A1:D10")

	// Add sheet views
	sv := ws.GetOrCreateSheetViews()
	view := sv.GetOrCreateSheetView()
	view.SetTabSelected(true)
	view.SetWorkbookViewId(0)

	// Add sheet format properties
	sfp := ws.GetOrCreateSheetFormatPr()
	sfp.SetDefaultRowHeight(15)

	// Add columns
	cols := ws.GetOrCreateCols()
	col := cols.AddCol(1, 3)
	col.SetWidth(12.5)
	col.SetCustomWidth(true)

	// Add sheet data
	ws.GetOrCreateSheetData()

	// Add page margins
	pm := ws.GetOrCreatePageMargins()
	pm.SetDefaults()

	// Get XML output
	xml := ws.OuterXml()

	// Verify XML structure - should start with worksheet element with namespace
	if !strings.HasPrefix(xml, "\u003cx:worksheet") {
		t.Error("expected XML to start with \u003cx:worksheet")
	}
	if !strings.HasSuffix(xml, "\u003c/x:worksheet\u003e") {
		t.Error("expected XML to end with \u003c/x:worksheet\u003e")
	}

	// Verify specific elements are properly structured with namespace
	elements := []struct {
		element string
		content string
	}{
		{"x:dimension", `ref="A1:D10"`},
		{"x:sheetViews", ""},
		{"x:sheetFormatPr", `defaultRowHeight="15"`},
		{"x:cols", ""},
		{"x:sheetData", ""},
		{"x:pageMargins", ""},
	}

	for _, elem := range elements {
		if !strings.Contains(xml, "\u003c"+elem.element) {
			t.Errorf("expected '%s' element in XML", elem.element)
		}
		if elem.content != "" && !strings.Contains(xml, elem.content) {
			t.Errorf("expected '%s' to contain '%s'", elem.element, elem.content)
		}
	}

	// Verify sheet view specific attributes
	if !strings.Contains(xml, `tabSelected="true"`) {
		t.Error("expected sheet view to have tabSelected attribute")
	}
	if !strings.Contains(xml, `workbookViewId="0"`) {
		t.Error("expected sheet view to have workbookViewId attribute")
	}

	// Verify column specific attributes - customWidth is "true" not "1"
	if !strings.Contains(xml, `width="12.5"`) {
		t.Error("expected column to have width attribute")
	}
	if !strings.Contains(xml, `customWidth="true"`) {
		t.Error("expected column to have customWidth attribute")
	}
}

func TestWorksheet_Clone(t *testing.T) {
	ws := NewWorksheet()
	dim := ws.GetOrCreateDimension()
	dim.SetRef("A1:Z100")

	cloned := ws.Clone()
	clone, ok := cloned.(*Worksheet)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Worksheet",
			cloned,
		)
	}
	if clone.Dimension() == nil {
		t.Error(
			"expected clone to have Dimension",
		)
	}
	if clone.Dimension().Ref() != "A1:Z100" {
		t.Errorf(
			"expected ref 'A1:Z100', got '%s'",
			clone.Dimension().Ref(),
		)
	}

	// Modify original, verify clone is independent
	dim.SetRef("B2:Y50")
	if clone.Dimension().Ref() != "A1:Z100" {
		t.Error(
			"expected clone to remain unchanged after modifying original",
		)
	}
}
