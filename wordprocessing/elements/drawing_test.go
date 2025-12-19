package elements

import (
	"strings"
	"testing"
)

func TestNewInlineDrawing(t *testing.T) {
	width := int64(914400)  // 1 inch
	height := int64(914400) // 1 inch
	relId := "rId1"

	drawing := NewInlineDrawing(width, height, relId)

	if drawing == nil {
		t.Fatal("NewInlineDrawing returned nil")
	}

	inline := drawing.Inline()
	if inline == nil {
		t.Fatal("Drawing.Inline() returned nil")
	}

	if inline.Width() != width {
		t.Errorf("Expected width %d, got %d", width, inline.Width())
	}

	if inline.Height() != height {
		t.Errorf("Expected height %d, got %d", height, inline.Height())
	}

	graphic := inline.Graphic()
	if graphic == nil {
		t.Fatal("InlineDrawing.Graphic() returned nil")
	}

	graphicData := graphic.GraphicData()
	if graphicData == nil {
		t.Fatal("Graphic.GraphicData() returned nil")
	}

	pic := graphicData.Picture()
	if pic == nil {
		t.Fatal("GraphicData.Picture() returned nil")
	}

	blipFill := pic.BlipFill()
	if blipFill == nil {
		t.Fatal("Picture.BlipFill() returned nil")
	}

	blip := blipFill.Blip()
	if blip == nil {
		t.Fatal("BlipFill.Blip() returned nil")
	}

	if blip.Embed() != relId {
		t.Errorf("Expected embed %s, got %s", relId, blip.Embed())
	}
}

func TestNewAnchorDrawing(t *testing.T) {
	width := int64(914400)  // 1 inch
	height := int64(457200) // 0.5 inch
	relId := "rId2"

	drawing := NewAnchorDrawing(width, height, relId)

	if drawing == nil {
		t.Fatal("NewAnchorDrawing returned nil")
	}

	anchor := drawing.Anchor()
	if anchor == nil {
		t.Fatal("Drawing.Anchor() returned nil")
	}

	if anchor.Width() != width {
		t.Errorf("Expected width %d, got %d", width, anchor.Width())
	}

	if anchor.Height() != height {
		t.Errorf("Expected height %d, got %d", height, anchor.Height())
	}

	// Test position properties
	posH := anchor.PositionH()
	if posH == nil {
		t.Fatal("AnchorDrawing.PositionH() returned nil")
	}

	posV := anchor.PositionV()
	if posV == nil {
		t.Fatal("AnchorDrawing.PositionV() returned nil")
	}
}

func TestAnchorDrawingWrapTypes(t *testing.T) {
	drawing := NewAnchorDrawing(914400, 914400, "rId1")
	anchor := drawing.Anchor()

	// Test WrapNone (default)
	anchor.WrapNone()
	xml := anchor.OuterXml()
	if !strings.Contains(xml, "wrapNone") {
		t.Error("Expected wrapNone in XML")
	}

	// Test WrapSquare
	anchor.WrapSquare()
	xml = anchor.OuterXml()
	if !strings.Contains(xml, "wrapSquare") {
		t.Error("Expected wrapSquare in XML")
	}
	if strings.Contains(xml, "wrapNone") {
		t.Error("wrapNone should be removed when setting wrapSquare")
	}

	// Test WrapTight
	anchor.WrapTight()
	xml = anchor.OuterXml()
	if !strings.Contains(xml, "wrapTight") {
		t.Error("Expected wrapTight in XML")
	}

	// Test WrapThrough
	anchor.WrapThrough()
	xml = anchor.OuterXml()
	if !strings.Contains(xml, "wrapThrough") {
		t.Error("Expected wrapThrough in XML")
	}

	// Test WrapTopAndBottom
	anchor.WrapTopAndBottom()
	xml = anchor.OuterXml()
	if !strings.Contains(xml, "wrapTopAndBottom") {
		t.Error("Expected wrapTopAndBottom in XML")
	}
}

func TestAnchorDrawingProperties(t *testing.T) {
	drawing := NewAnchorDrawing(914400, 914400, "rId1")
	anchor := drawing.Anchor()

	// Test SetBehindDoc
	anchor.SetBehindDoc(true)
	if !anchor.BehindDoc() {
		t.Error("Expected BehindDoc to be true")
	}
	anchor.SetBehindDoc(false)
	if anchor.BehindDoc() {
		t.Error("Expected BehindDoc to be false")
	}

	// Test SetSimplePos
	anchor.SetSimplePos(true)
	if !anchor.SimplePos() {
		t.Error("Expected SimplePos to be true")
	}
	anchor.SetSimplePos(false)
	if anchor.SimplePos() {
		t.Error("Expected SimplePos to be false")
	}
}

func TestDocProperties(t *testing.T) {
	dp := NewDocProperties(1, "TestPicture")

	if dp.Id() != 1 {
		t.Errorf("Expected ID 1, got %d", dp.Id())
	}

	if dp.Name() != "TestPicture" {
		t.Errorf("Expected name TestPicture, got %s", dp.Name())
	}

	dp.SetDescription("A test description")
	if dp.Description() != "A test description" {
		t.Errorf("Expected description 'A test description', got %s", dp.Description())
	}

	dp.SetId(5)
	if dp.Id() != 5 {
		t.Errorf("Expected ID 5, got %d", dp.Id())
	}

	dp.SetName("NewName")
	if dp.Name() != "NewName" {
		t.Errorf("Expected name NewName, got %s", dp.Name())
	}
}

func TestExtent(t *testing.T) {
	width := int64(1000000)
	height := int64(500000)

	ext := NewExtent(width, height)

	if ext.Width() != width {
		t.Errorf("Expected width %d, got %d", width, ext.Width())
	}

	if ext.Height() != height {
		t.Errorf("Expected height %d, got %d", height, ext.Height())
	}

	ext.SetWidth(2000000)
	if ext.Width() != 2000000 {
		t.Errorf("Expected width 2000000, got %d", ext.Width())
	}

	ext.SetHeight(1000000)
	if ext.Height() != 1000000 {
		t.Errorf("Expected height 1000000, got %d", ext.Height())
	}
}

func TestPositionH(t *testing.T) {
	posH := NewPositionH(RelativeFromHPage, 914400)

	if posH.RelativeFrom() != RelativeFromHPage {
		t.Errorf("Expected relativeFrom %s, got %s", RelativeFromHPage, posH.RelativeFrom())
	}

	if posH.Offset() != 914400 {
		t.Errorf("Expected offset 914400, got %d", posH.Offset())
	}

	posH.SetRelativeFrom(RelativeFromHMargin)
	if posH.RelativeFrom() != RelativeFromHMargin {
		t.Errorf("Expected relativeFrom %s, got %s", RelativeFromHMargin, posH.RelativeFrom())
	}

	posH.SetOffset(457200)
	if posH.Offset() != 457200 {
		t.Errorf("Expected offset 457200, got %d", posH.Offset())
	}
}

func TestPositionV(t *testing.T) {
	posV := NewPositionV(RelativeFromVPage, 914400)

	if posV.RelativeFrom() != RelativeFromVPage {
		t.Errorf("Expected relativeFrom %s, got %s", RelativeFromVPage, posV.RelativeFrom())
	}

	if posV.Offset() != 914400 {
		t.Errorf("Expected offset 914400, got %d", posV.Offset())
	}

	posV.SetRelativeFrom(RelativeFromVMargin)
	if posV.RelativeFrom() != RelativeFromVMargin {
		t.Errorf("Expected relativeFrom %s, got %s", RelativeFromVMargin, posV.RelativeFrom())
	}

	posV.SetOffset(457200)
	if posV.Offset() != 457200 {
		t.Errorf("Expected offset 457200, got %d", posV.Offset())
	}
}

func TestBlip(t *testing.T) {
	blip := NewBlip("rId5")

	if blip.Embed() != "rId5" {
		t.Errorf("Expected embed rId5, got %s", blip.Embed())
	}

	blip.SetEmbed("rId10")
	if blip.Embed() != "rId10" {
		t.Errorf("Expected embed rId10, got %s", blip.Embed())
	}

	blip.SetCompressionState("print")
	if blip.CompressionState() != "print" {
		t.Errorf("Expected compression state print, got %s", blip.CompressionState())
	}
}

func TestTransform2D(t *testing.T) {
	xfrm := NewTransform2D(0, 0, 914400, 914400)

	// Test rotation
	xfrm.SetRotation(5400000) // 90 degrees
	if xfrm.Rotation() != 5400000 {
		t.Errorf("Expected rotation 5400000, got %d", xfrm.Rotation())
	}

	// Test flip
	xfrm.SetFlipH(true)
	if !xfrm.FlipH() {
		t.Error("Expected FlipH to be true")
	}

	xfrm.SetFlipV(true)
	if !xfrm.FlipV() {
		t.Error("Expected FlipV to be true")
	}
}

func TestPresetGeometry(t *testing.T) {
	pg := NewPresetGeometry("rect")

	if pg.Preset() != "rect" {
		t.Errorf("Expected preset rect, got %s", pg.Preset())
	}

	pg.SetPreset("ellipse")
	if pg.Preset() != "ellipse" {
		t.Errorf("Expected preset ellipse, got %s", pg.Preset())
	}
}

func TestEMUConstants(t *testing.T) {
	// Verify EMU constants are correct
	if EMUsPerInch != 914400 {
		t.Errorf("Expected EMUsPerInch to be 914400, got %d", EMUsPerInch)
	}

	if EMUsPerCm != 360000 {
		t.Errorf("Expected EMUsPerCm to be 360000, got %d", EMUsPerCm)
	}

	if EMUsPerPoint != 12700 {
		t.Errorf("Expected EMUsPerPoint to be 12700, got %d", EMUsPerPoint)
	}
}
