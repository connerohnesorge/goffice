package elements

import (
	"strings"
	"testing"
)

const (
	testRelationshipID = "rId2"
)

func TestNewWorksheetDrawing(t *testing.T) {
	wd := NewWorksheetDrawing()

	if wd == nil {
		t.Fatal(
			"expected non-nil WorksheetDrawing",
		)
	}
	if wd.LocalName() != "wsDr" {
		t.Errorf(
			"expected localName 'wsDr', got '%s'",
			wd.LocalName(),
		)
	}
	if wd.NamespaceURI() != NamespaceSpreadsheetDrawing {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceSpreadsheetDrawing,
			wd.NamespaceURI(),
		)
	}
}

func TestWorksheetDrawing_AddAnchors(
	t *testing.T,
) {
	wd := NewWorksheetDrawing()

	// Test TwoCellAnchor
	tca := wd.AddTwoCellAnchor()
	if tca == nil {
		t.Fatal("expected non-nil TwoCellAnchor")
	}
	if len(wd.TwoCellAnchors()) != 1 {
		t.Errorf(
			"expected 1 TwoCellAnchor, got %d",
			len(wd.TwoCellAnchors()),
		)
	}

	// Test OneCellAnchor
	oca := wd.AddOneCellAnchor()
	if oca == nil {
		t.Fatal("expected non-nil OneCellAnchor")
	}
	if len(wd.OneCellAnchors()) != 1 {
		t.Errorf(
			"expected 1 OneCellAnchor, got %d",
			len(wd.OneCellAnchors()),
		)
	}

	// Test AbsoluteAnchor
	aa := wd.AddAbsoluteAnchor()
	if aa == nil {
		t.Fatal("expected non-nil AbsoluteAnchor")
	}
	if len(wd.AbsoluteAnchors()) != 1 {
		t.Errorf(
			"expected 1 AbsoluteAnchor, got %d",
			len(wd.AbsoluteAnchors()),
		)
	}
}

func TestTwoCellAnchor(t *testing.T) {
	tca := NewTwoCellAnchor()

	// Test EditAs
	if tca.EditAs() != EditAsTwoCell {
		t.Errorf(
			"expected default EditAs 'twoCell', got '%s'",
			tca.EditAs(),
		)
	}
	tca.SetEditAs(EditAsOneCell)
	if tca.EditAs() != EditAsOneCell {
		t.Errorf(
			"expected EditAs 'oneCell', got '%s'",
			tca.EditAs(),
		)
	}
	tca.SetEditAs(EditAsAbsolute)
	if tca.EditAs() != EditAsAbsolute {
		t.Errorf(
			"expected EditAs 'absolute', got '%s'",
			tca.EditAs(),
		)
	}

	// Test From marker
	from := tca.GetOrCreateFrom()
	if from == nil {
		t.Fatal("expected non-nil From marker")
	}
	from.SetPosition(1, 100, 2, 200)
	if from.Col() != 1 {
		t.Errorf(
			"expected Col 1, got %d",
			from.Col(),
		)
	}
	if from.ColOff() != 100 {
		t.Errorf(
			"expected ColOff 100, got %d",
			from.ColOff(),
		)
	}
	if from.Row() != 2 {
		t.Errorf(
			"expected Row 2, got %d",
			from.Row(),
		)
	}
	if from.RowOff() != 200 {
		t.Errorf(
			"expected RowOff 200, got %d",
			from.RowOff(),
		)
	}

	// Test To marker
	to := tca.GetOrCreateTo()
	if to == nil {
		t.Fatal("expected non-nil To marker")
	}
	to.SetPosition(5, 500, 10, 1000)
	if to.Col() != 5 {
		t.Errorf(
			"expected Col 5, got %d",
			to.Col(),
		)
	}
	if to.Row() != 10 {
		t.Errorf(
			"expected Row 10, got %d",
			to.Row(),
		)
	}

	// Test ClientData
	cd := tca.GetOrCreateClientData()
	if cd == nil {
		t.Fatal("expected non-nil ClientData")
	}
}

func TestOneCellAnchor(t *testing.T) {
	oca := NewOneCellAnchor()

	// Test From
	from := oca.GetOrCreateFrom()
	if from == nil {
		t.Fatal("expected non-nil From marker")
	}

	// Test Ext
	ext := oca.GetOrCreateExt()
	if ext == nil {
		t.Fatal("expected non-nil Ext")
	}
	ext.SetSize(914400, 914400) // 1 inch x 1 inch
	if ext.Cx() != 914400 {
		t.Errorf(
			"expected Cx 914400, got %d",
			ext.Cx(),
		)
	}
	if ext.Cy() != 914400 {
		t.Errorf(
			"expected Cy 914400, got %d",
			ext.Cy(),
		)
	}
}

func TestAbsoluteAnchor(t *testing.T) {
	aa := NewAbsoluteAnchor()

	// Test Pos
	pos := aa.GetOrCreatePos()
	if pos == nil {
		t.Fatal("expected non-nil Pos")
	}
	pos.SetPosition(1000000, 2000000)
	if pos.X() != 1000000 {
		t.Errorf(
			"expected X 1000000, got %d",
			pos.X(),
		)
	}
	if pos.Y() != 2000000 {
		t.Errorf(
			"expected Y 2000000, got %d",
			pos.Y(),
		)
	}

	// Test Ext
	ext := aa.GetOrCreateExt()
	if ext == nil {
		t.Fatal("expected non-nil Ext")
	}
}

func TestFromMarker(t *testing.T) {
	fm := NewFromMarker()

	if fm.LocalName() != "from" {
		t.Errorf(
			"expected localName 'from', got '%s'",
			fm.LocalName(),
		)
	}

	// Set individual values
	fm.SetCol(3)
	fm.SetColOff(1000)
	fm.SetRow(5)
	fm.SetRowOff(2000)

	if fm.Col() != 3 {
		t.Errorf(
			"expected Col 3, got %d",
			fm.Col(),
		)
	}
	if fm.ColOff() != 1000 {
		t.Errorf(
			"expected ColOff 1000, got %d",
			fm.ColOff(),
		)
	}
	if fm.Row() != 5 {
		t.Errorf(
			"expected Row 5, got %d",
			fm.Row(),
		)
	}
	if fm.RowOff() != 2000 {
		t.Errorf(
			"expected RowOff 2000, got %d",
			fm.RowOff(),
		)
	}
}

func TestToMarker(t *testing.T) {
	tm := NewToMarker()

	if tm.LocalName() != "to" {
		t.Errorf(
			"expected localName 'to', got '%s'",
			tm.LocalName(),
		)
	}

	tm.SetPosition(10, 500, 20, 1000)
	if tm.Col() != 10 {
		t.Errorf(
			"expected Col 10, got %d",
			tm.Col(),
		)
	}
	if tm.ColOff() != 500 {
		t.Errorf(
			"expected ColOff 500, got %d",
			tm.ColOff(),
		)
	}
	if tm.Row() != 20 {
		t.Errorf(
			"expected Row 20, got %d",
			tm.Row(),
		)
	}
	if tm.RowOff() != 1000 {
		t.Errorf(
			"expected RowOff 1000, got %d",
			tm.RowOff(),
		)
	}
}

func TestPosition(t *testing.T) {
	pos := NewPosition()

	if pos.LocalName() != "pos" {
		t.Errorf(
			"expected localName 'pos', got '%s'",
			pos.LocalName(),
		)
	}

	pos.SetX(5000000)
	pos.SetY(3000000)

	if pos.X() != 5000000 {
		t.Errorf(
			"expected X 5000000, got %d",
			pos.X(),
		)
	}
	if pos.Y() != 3000000 {
		t.Errorf(
			"expected Y 3000000, got %d",
			pos.Y(),
		)
	}
}

func TestExtent(t *testing.T) {
	ext := NewExtent()

	if ext.LocalName() != "ext" {
		t.Errorf(
			"expected localName 'ext', got '%s'",
			ext.LocalName(),
		)
	}

	ext.SetCx(1000000)
	ext.SetCy(500000)

	if ext.Cx() != 1000000 {
		t.Errorf(
			"expected Cx 1000000, got %d",
			ext.Cx(),
		)
	}
	if ext.Cy() != 500000 {
		t.Errorf(
			"expected Cy 500000, got %d",
			ext.Cy(),
		)
	}
}

func TestEMUConversions(t *testing.T) {
	// Test inches
	emus := InchesToEMUs(1.0)
	if emus != 914400 {
		t.Errorf(
			"expected 914400 EMUs per inch, got %d",
			emus,
		)
	}
	inches := EMUsToInches(914400)
	if inches != 1.0 {
		t.Errorf(
			"expected 1.0 inches, got %f",
			inches,
		)
	}

	// Test centimeters
	emus = CmToEMUs(1.0)
	if emus != 360000 {
		t.Errorf(
			"expected 360000 EMUs per cm, got %d",
			emus,
		)
	}

	// Test points
	emus = PointsToEMUs(1.0)
	if emus != 12700 {
		t.Errorf(
			"expected 12700 EMUs per point, got %d",
			emus,
		)
	}

	// Test pixels
	emus = PixelsToEMUs(1)
	if emus != 9525 {
		t.Errorf(
			"expected 9525 EMUs per pixel, got %d",
			emus,
		)
	}
}

func TestShape(t *testing.T) {
	sp := NewShape()

	if sp.LocalName() != "sp" {
		t.Errorf(
			"expected localName 'sp', got '%s'",
			sp.LocalName(),
		)
	}

	// Test macro
	sp.SetMacro("TestMacro")
	if sp.Macro() != "TestMacro" {
		t.Errorf(
			"expected macro 'TestMacro', got '%s'",
			sp.Macro(),
		)
	}

	// Test textlink
	sp.SetTextLink("$A$1")
	if sp.TextLink() != "$A$1" {
		t.Errorf(
			"expected textlink '$A$1', got '%s'",
			sp.TextLink(),
		)
	}

	// Test FLocksText
	if !sp.FLocksText() {
		t.Error(
			"expected FLocksText to be true by default",
		)
	}
	sp.SetFLocksText(false)
	if sp.FLocksText() {
		t.Error("expected FLocksText to be false")
	}

	// Test FPublished
	if sp.FPublished() {
		t.Error(
			"expected FPublished to be false by default",
		)
	}
	sp.SetFPublished(true)
	if !sp.FPublished() {
		t.Error("expected FPublished to be true")
	}

	// Test NvSpPr
	nvSpPr := sp.GetOrCreateNvSpPr()
	if nvSpPr == nil {
		t.Fatal("expected non-nil NvSpPr")
	}

	// Test SpPr
	spPr := sp.GetOrCreateSpPr()
	if spPr == nil {
		t.Fatal("expected non-nil SpPr")
	}

	// Test Style
	style := sp.GetOrCreateStyle()
	if style == nil {
		t.Fatal("expected non-nil Style")
	}

	// Test TxBody
	txBody := sp.GetOrCreateTxBody()
	if txBody == nil {
		t.Fatal("expected non-nil TxBody")
	}
}

func TestNvSpPr(t *testing.T) {
	nvSpPr := NewNvSpPr()

	// Test CNvPr
	cNvPr := nvSpPr.GetOrCreateCNvPr()
	if cNvPr == nil {
		t.Fatal("expected non-nil CNvPr")
	}
	cNvPr.SetId(1)
	cNvPr.SetName("Shape 1")
	if cNvPr.Id() != 1 {
		t.Errorf(
			"expected Id 1, got %d",
			cNvPr.Id(),
		)
	}
	if cNvPr.Name() != "Shape 1" {
		t.Errorf(
			"expected name 'Shape 1', got '%s'",
			cNvPr.Name(),
		)
	}

	// Test CNvSpPr
	cNvSpPr := nvSpPr.GetOrCreateCNvSpPr()
	if cNvSpPr == nil {
		t.Fatal("expected non-nil CNvSpPr")
	}
}

func TestCNvPr(t *testing.T) {
	cNvPr := NewCNvPr()

	// Test Id
	cNvPr.SetId(5)
	if cNvPr.Id() != 5 {
		t.Errorf(
			"expected Id 5, got %d",
			cNvPr.Id(),
		)
	}

	// Test Name
	cNvPr.SetName("Test Shape")
	if cNvPr.Name() != "Test Shape" {
		t.Errorf(
			"expected name 'Test Shape', got '%s'",
			cNvPr.Name(),
		)
	}

	// Test Descr
	cNvPr.SetDescr("Description")
	if cNvPr.Descr() != "Description" {
		t.Errorf(
			"expected descr 'Description', got '%s'",
			cNvPr.Descr(),
		)
	}

	// Test Hidden
	if cNvPr.Hidden() {
		t.Error(
			"expected hidden to be false by default",
		)
	}
	cNvPr.SetHidden(true)
	if !cNvPr.Hidden() {
		t.Error("expected hidden to be true")
	}
}

func TestCNvSpPr(t *testing.T) {
	cNvSpPr := NewCNvSpPr()

	// Test TxBox
	if cNvSpPr.TxBox() {
		t.Error(
			"expected txBox to be false by default",
		)
	}
	cNvSpPr.SetTxBox(true)
	if !cNvSpPr.TxBox() {
		t.Error("expected txBox to be true")
	}
}

func TestDrawingPicture(t *testing.T) {
	pic := NewDrawingPicture()

	if pic.LocalName() != "pic" {
		t.Errorf(
			"expected localName 'pic', got '%s'",
			pic.LocalName(),
		)
	}

	// Test macro
	pic.SetMacro("PicMacro")
	if pic.Macro() != "PicMacro" {
		t.Errorf(
			"expected macro 'PicMacro', got '%s'",
			pic.Macro(),
		)
	}

	// Test NvPicPr
	nvPicPr := pic.GetOrCreateNvPicPr()
	if nvPicPr == nil {
		t.Fatal("expected non-nil NvPicPr")
	}

	// Test BlipFill
	blipFill := pic.GetOrCreateBlipFill()
	if blipFill == nil {
		t.Fatal("expected non-nil BlipFill")
	}

	// Test SpPr
	spPr := pic.GetOrCreateSpPr()
	if spPr == nil {
		t.Fatal("expected non-nil SpPr")
	}

	// Test convenience method
	pic.SetImageRelationship("rId1")
	blip := blipFill.Blip()
	if blip == nil {
		t.Fatal("expected non-nil Blip")
	}
	if blip.Embed() != relationIdOne {
		t.Errorf(
			"expected embed 'rId1', got '%s'",
			blip.Embed(),
		)
	}
}

func TestBlip(t *testing.T) {
	blip := NewBlip()

	if blip.LocalName() != "blip" {
		t.Errorf(
			"expected localName 'blip', got '%s'",
			blip.LocalName(),
		)
	}
	if blip.NamespaceURI() != NamespaceDrawingML {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceDrawingML,
			blip.NamespaceURI(),
		)
	}

	// Test Embed
	blip.SetEmbed(testRelationshipID)
	if blip.Embed() != testRelationshipID {
		t.Errorf(
			"expected embed 'rId2', got '%s'",
			blip.Embed(),
		)
	}

	// Test Link
	blip.SetLink("rId3")
	if blip.Link() != "rId3" {
		t.Errorf(
			"expected link 'rId3', got '%s'",
			blip.Link(),
		)
	}

	// Test Cstate
	blip.SetCstate("print")
	if blip.Cstate() != "print" {
		t.Errorf(
			"expected cstate 'print', got '%s'",
			blip.Cstate(),
		)
	}
}

func TestGraphicFrame(t *testing.T) {
	gf := NewGraphicFrame()

	if gf.LocalName() != "graphicFrame" {
		t.Errorf(
			"expected localName 'graphicFrame', got '%s'",
			gf.LocalName(),
		)
	}

	// Test macro
	gf.SetMacro("ChartMacro")
	if gf.Macro() != "ChartMacro" {
		t.Errorf(
			"expected macro 'ChartMacro', got '%s'",
			gf.Macro(),
		)
	}

	// Test NvGraphicFramePr
	nvGfPr := gf.GetOrCreateNvGraphicFramePr()
	if nvGfPr == nil {
		t.Fatal(
			"expected non-nil NvGraphicFramePr",
		)
	}

	// Test Xfrm
	xfrm := gf.GetOrCreateXfrm()
	if xfrm == nil {
		t.Fatal("expected non-nil Xfrm")
	}

	// Test Graphic
	g := gf.GetOrCreateGraphic()
	if g == nil {
		t.Fatal("expected non-nil Graphic")
	}
}

func TestGraphic(t *testing.T) {
	g := NewGraphic()

	if g.LocalName() != "graphic" {
		t.Errorf(
			"expected localName 'graphic', got '%s'",
			g.LocalName(),
		)
	}
	if g.NamespaceURI() != NamespaceDrawingML {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceDrawingML,
			g.NamespaceURI(),
		)
	}

	// Test GraphicData
	gd := g.GetOrCreateGraphicData()
	if gd == nil {
		t.Fatal("expected non-nil GraphicData")
	}
}

func TestGraphicData(t *testing.T) {
	gd := NewGraphicData()

	if gd.LocalName() != "graphicData" {
		t.Errorf(
			"expected localName 'graphicData', got '%s'",
			gd.LocalName(),
		)
	}

	// Test URI
	gd.SetURI("http://example.com/test")
	if gd.URI() != "http://example.com/test" {
		t.Errorf(
			"expected uri 'http://example.com/test', got '%s'",
			gd.URI(),
		)
	}

	// Test SetChartURI
	gd.SetChartURI()
	if gd.URI() != NamespaceChart {
		t.Errorf(
			"expected uri '%s', got '%s'",
			NamespaceChart,
			gd.URI(),
		)
	}
}

func TestClientData(t *testing.T) {
	cd := NewClientData()

	if cd.LocalName() != "clientData" {
		t.Errorf(
			"expected localName 'clientData', got '%s'",
			cd.LocalName(),
		)
	}

	// Test FLocksWithSheet (default true)
	if !cd.FLocksWithSheet() {
		t.Error(
			"expected FLocksWithSheet to be true by default",
		)
	}
	cd.SetFLocksWithSheet(false)
	if cd.FLocksWithSheet() {
		t.Error(
			"expected FLocksWithSheet to be false",
		)
	}

	// Test FPrintsWithSheet (default true)
	if !cd.FPrintsWithSheet() {
		t.Error(
			"expected FPrintsWithSheet to be true by default",
		)
	}
	cd.SetFPrintsWithSheet(false)
	if cd.FPrintsWithSheet() {
		t.Error(
			"expected FPrintsWithSheet to be false",
		)
	}
}

func TestWorksheetDrawing_XMLSerialization(
	t *testing.T,
) {
	wd := NewWorksheetDrawing()

	// Add a two-cell anchor with a shape
	tca := wd.AddTwoCellAnchor()
	tca.SetEditAs(EditAsOneCell)

	from := tca.GetOrCreateFrom()
	from.SetPosition(0, 0, 0, 0)

	to := tca.GetOrCreateTo()
	to.SetPosition(5, 500, 10, 1000)

	sp := tca.GetOrCreateShape()
	nvSpPr := sp.GetOrCreateNvSpPr()
	cNvPr := nvSpPr.GetOrCreateCNvPr()
	cNvPr.SetId(1)
	cNvPr.SetName("Rectangle 1")

	cd := tca.GetOrCreateClientData()
	cd.SetFLocksWithSheet(true)

	// Get XML output
	xml := wd.OuterXml()

	// Verify key elements are present
	if !strings.Contains(xml, "wsDr") {
		t.Error("expected 'wsDr' in XML")
	}
	if !strings.Contains(xml, "twoCellAnchor") {
		t.Error("expected 'twoCellAnchor' in XML")
	}
	if !strings.Contains(
		xml,
		"editAs=\"oneCell\"",
	) {
		t.Error(
			"expected 'editAs=\"oneCell\"' in XML",
		)
	}
	if !strings.Contains(xml, "from") {
		t.Error("expected 'from' in XML")
	}
	if !strings.Contains(xml, "to") {
		t.Error("expected 'to' in XML")
	}
	if !strings.Contains(xml, "sp") {
		t.Error("expected 'sp' in XML")
	}
	if !strings.Contains(xml, "clientData") {
		t.Error("expected 'clientData' in XML")
	}
}

func TestClone_DrawingElements(t *testing.T) {
	// Test WorksheetDrawing clone
	wd := NewWorksheetDrawing()
	tca := wd.AddTwoCellAnchor()
	tca.SetEditAs(EditAsAbsolute)

	cloned := wd.Clone()
	clone, ok := cloned.(*WorksheetDrawing)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *WorksheetDrawing",
			cloned,
		)
	}

	anchors := clone.TwoCellAnchors()
	if len(anchors) != 1 {
		t.Errorf(
			"expected 1 anchor in clone, got %d",
			len(anchors),
		)
	}

	// Modify original, verify clone is independent
	wd.AddTwoCellAnchor()
	if len(clone.TwoCellAnchors()) != 1 {
		t.Error(
			"expected clone to remain unchanged after modifying original",
		)
	}
}

func TestTwoCellAnchor_WithPicture(t *testing.T) {
	tca := NewTwoCellAnchor()

	from := tca.GetOrCreateFrom()
	from.SetPosition(0, 0, 0, 0)

	to := tca.GetOrCreateTo()
	to.SetPosition(3, 0, 5, 0)

	pic := tca.GetOrCreatePicture()
	if pic == nil {
		t.Fatal("expected non-nil Picture")
	}

	pic.SetPictureId(1, "Test Image")
	pic.SetImageRelationship("rId1")

	// Verify the picture was set up correctly
	nvPicPr := pic.NvPicPr()
	if nvPicPr == nil {
		t.Fatal("expected non-nil NvPicPr")
	}
	cNvPr := nvPicPr.CNvPr()
	if cNvPr == nil {
		t.Fatal("expected non-nil CNvPr")
	}
	if cNvPr.Id() != 1 {
		t.Errorf(
			"expected Id 1, got %d",
			cNvPr.Id(),
		)
	}
}

func TestTwoCellAnchor_WithGraphicFrame(
	t *testing.T,
) {
	tca := NewTwoCellAnchor()

	from := tca.GetOrCreateFrom()
	from.SetPosition(0, 0, 0, 0)

	to := tca.GetOrCreateTo()
	to.SetPosition(10, 0, 20, 0)

	gf := tca.GetOrCreateGraphicFrame()
	if gf == nil {
		t.Fatal("expected non-nil GraphicFrame")
	}

	gf.SetChartInfo(1, "Sales Chart")

	g := gf.GetOrCreateGraphic()
	gd := g.GetOrCreateGraphicData()
	gd.SetChartURI()

	// Verify
	nvGfPr := gf.NvGraphicFramePr()
	if nvGfPr == nil {
		t.Fatal(
			"expected non-nil NvGraphicFramePr",
		)
	}
	cNvPr := nvGfPr.CNvPr()
	if cNvPr == nil {
		t.Fatal("expected non-nil CNvPr")
	}
	if cNvPr.Id() != 1 {
		t.Errorf(
			"expected Id 1, got %d",
			cNvPr.Id(),
		)
	}
}

func TestConnectionShape(t *testing.T) {
	cs := NewConnectionShape()

	if cs.LocalName() != "cxnSp" {
		t.Errorf(
			"expected localName 'cxnSp', got '%s'",
			cs.LocalName(),
		)
	}

	// Test macro
	cs.SetMacro("ConnectorMacro")
	if cs.Macro() != "ConnectorMacro" {
		t.Errorf(
			"expected macro 'ConnectorMacro', got '%s'",
			cs.Macro(),
		)
	}

	// Test FPublished
	if cs.FPublished() {
		t.Error(
			"expected FPublished to be false by default",
		)
	}
	cs.SetFPublished(true)
	if !cs.FPublished() {
		t.Error("expected FPublished to be true")
	}

	// Test NvCxnSpPr
	nvCxnSpPr := cs.GetOrCreateNvCxnSpPr()
	if nvCxnSpPr == nil {
		t.Fatal("expected non-nil NvCxnSpPr")
	}

	// Test CNvPr in NvCxnSpPr
	cNvPr := nvCxnSpPr.GetOrCreateCNvPr()
	if cNvPr == nil {
		t.Fatal("expected non-nil CNvPr")
	}
	cNvPr.SetId(1)
	cNvPr.SetName("Connector 1")
	if cNvPr.Id() != 1 {
		t.Errorf(
			"expected Id 1, got %d",
			cNvPr.Id(),
		)
	}

	// Test CNvCxnSpPr in NvCxnSpPr
	cNvCxnSpPr := nvCxnSpPr.GetOrCreateCNvCxnSpPr()
	if cNvCxnSpPr == nil {
		t.Fatal("expected non-nil CNvCxnSpPr")
	}

	// Test SpPr
	spPr := cs.GetOrCreateSpPr()
	if spPr == nil {
		t.Fatal("expected non-nil SpPr")
	}
}

func TestGroupShape(t *testing.T) {
	gs := NewGroupShape()

	if gs.LocalName() != "grpSp" {
		t.Errorf(
			"expected localName 'grpSp', got '%s'",
			gs.LocalName(),
		)
	}

	// Test NvGrpSpPr
	nvGrpSpPr := gs.GetOrCreateNvGrpSpPr()
	if nvGrpSpPr == nil {
		t.Fatal("expected non-nil NvGrpSpPr")
	}

	// Test CNvPr in NvGrpSpPr
	cNvPr := nvGrpSpPr.GetOrCreateCNvPr()
	if cNvPr == nil {
		t.Fatal("expected non-nil CNvPr")
	}
	cNvPr.SetId(1)
	cNvPr.SetName("Group 1")

	// Test CNvGrpSpPr in NvGrpSpPr
	cNvGrpSpPr := nvGrpSpPr.GetOrCreateCNvGrpSpPr()
	if cNvGrpSpPr == nil {
		t.Fatal("expected non-nil CNvGrpSpPr")
	}

	// Test GrpSpPr
	grpSpPr := gs.GetOrCreateGrpSpPr()
	if grpSpPr == nil {
		t.Fatal("expected non-nil GrpSpPr")
	}

	// Test BWMode
	grpSpPr.SetBWMode("auto")
	if grpSpPr.BWMode() != "auto" {
		t.Errorf(
			"expected bwMode 'auto', got '%s'",
			grpSpPr.BWMode(),
		)
	}

	// Test adding child shapes
	sp := gs.AddShape()
	if sp == nil {
		t.Fatal("expected non-nil Shape")
	}

	pic := gs.AddPicture()
	if pic == nil {
		t.Fatal("expected non-nil Picture")
	}

	cxnSp := gs.AddConnectionShape()
	if cxnSp == nil {
		t.Fatal(
			"expected non-nil ConnectionShape",
		)
	}

	gf := gs.AddGraphicFrame()
	if gf == nil {
		t.Fatal("expected non-nil GraphicFrame")
	}

	// Test nested group shape
	nestedGs := gs.AddGroupShape()
	if nestedGs == nil {
		t.Fatal(
			"expected non-nil nested GroupShape",
		)
	}
}

func TestGroupShape_Clone(t *testing.T) {
	gs := NewGroupShape()

	nvGrpSpPr := gs.GetOrCreateNvGrpSpPr()
	cNvPr := nvGrpSpPr.GetOrCreateCNvPr()
	cNvPr.SetId(1)
	cNvPr.SetName("Group 1")

	gs.AddShape()
	gs.AddPicture()

	cloned := gs.Clone()
	clone, ok := cloned.(*GroupShape)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *GroupShape",
			cloned,
		)
	}

	cloneNvGrpSpPr := clone.NvGrpSpPr()
	if cloneNvGrpSpPr == nil {
		t.Fatal(
			"expected non-nil NvGrpSpPr in clone",
		)
	}

	cloneCNvPr := cloneNvGrpSpPr.CNvPr()
	if cloneCNvPr == nil {
		t.Fatal("expected non-nil CNvPr in clone")
	}
	if cloneCNvPr.Id() != 1 {
		t.Errorf(
			"expected Id 1 in clone, got %d",
			cloneCNvPr.Id(),
		)
	}
}
