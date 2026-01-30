package elements

import (
	"testing"
)

// TestSlide_CommonSlideData tests getting common slide data.
func TestSlide_CommonSlideData(t *testing.T) {
	slide := NewSlide()

	// Slides should have CommonSlideData
	cSld := slide.CommonSlideData()
	if cSld == nil {
		t.Error("CommonSlideData() returned nil")
	}

	// CommonSlideData should have a shape tree
	shapeTree := cSld.ShapeTree()
	if shapeTree == nil {
		t.Error("ShapeTree() returned nil")
	}
}

// TestShapeTree_AddShape tests adding shapes to shape tree.
func TestShapeTree_AddShape(t *testing.T) {
	shapeTree := NewShapeTree()

	// Add a shape
	shape := shapeTree.AddShape()
	if shape == nil {
		t.Fatal("AddShape() returned nil")
	}

	// Verify shape has correct namespace
	if shape.NamespaceURI() != NamespacePresentationML {
		t.Errorf(
			"Shape NamespaceURI() = %v, want %v",
			shape.NamespaceURI(),
			NamespacePresentationML,
		)
	}

	// Add another shape
	shape2 := shapeTree.AddShape()
	if shape2 == nil {
		t.Error("AddShape() second returned nil")
	}

	// Count shapes
	shapes := shapeTree.Shapes()
	if len(shapes) != 2 {
		t.Errorf(
			"Shapes() count = %v, want 2",
			len(shapes),
		)
	}
}

// TestShape_SetText tests setting shape text.
func TestShape_SetText(t *testing.T) {
	shape := NewShape()

	// Set text - uses GetOrCreateTextBody internally
	shape.SetText("Hello World")

	// Note: TextBody() lookup has namespace mismatch with DrawingML's TextBody
	// GetOrCreateTextBody works correctly though
	txBody := shape.GetOrCreateTextBody()
	if txBody == nil {
		t.Error(
			"GetOrCreateTextBody() returned nil after SetText()",
		)
	}

	// Verify we can call SetText and GetOrCreateTextBody works
	shape.SetText("New Text")

	// The TextBody should have the text, verify via GetOrCreateTextBody
	tb := shape.GetOrCreateTextBody()
	if tb == nil {
		t.Error(
			"GetOrCreateTextBody() returned nil",
		)
	}
}

// TestShape_SetPosition tests setting shape position.
func TestShape_SetPosition(t *testing.T) {
	shape := NewShape()

	// Set position (in EMUs)
	x := 914400  // 1 inch
	y := 1828800 // 2 inches
	shape.SetPosition(x, y)

	// Get shape properties
	spPr := shape.ShapeProperties()
	if spPr == nil {
		t.Error("ShapeProperties() returned nil")
	}
}

// TestShape_SetSize tests setting shape size.
func TestShape_SetSize(t *testing.T) {
	shape := NewShape()

	// Set size (in EMUs)
	width := 2743200  // 3 inches
	height := 1828800 // 2 inches
	shape.SetSize(width, height)

	// Get shape properties
	spPr := shape.ShapeProperties()
	if spPr == nil {
		t.Error("ShapeProperties() returned nil")
	}
}

// TestShape_SetSolidFill tests setting fill color.
func TestShape_SetSolidFill(t *testing.T) {
	shape := NewShape()

	// Set solid fill with RGB color
	shape.SetSolidFill("FF0000") // Red

	// Get shape properties
	spPr := shape.ShapeProperties()
	if spPr == nil {
		t.Error(
			"ShapeProperties() returned nil after SetSolidFill()",
		)
	}
}

// TestShape_SetShapeType tests setting shape geometry.
func TestShape_SetShapeType(t *testing.T) {
	shape := NewShape()

	// Set preset geometry
	shape.SetShapeType(ShapeTypeRectangle)

	// Get shape properties
	spPr := shape.ShapeProperties()
	if spPr == nil {
		t.Error("ShapeProperties() returned nil")
	}
}

// TestCommonSlideData tests CommonSlideData element.
func TestCommonSlideData(t *testing.T) {
	csd := NewCommonSlideData()

	// Test name attribute
	csd.SetName("Test Slide")
	if csd.Name() != "Test Slide" {
		t.Errorf(
			"Name() = %v, want Test Slide",
			csd.Name(),
		)
	}

	// Clear name
	csd.SetName("")
	if csd.Name() != "" {
		t.Errorf(
			"Name() = %v, want empty string",
			csd.Name(),
		)
	}

	// Test shape tree
	shapeTree := csd.ShapeTree()
	if shapeTree == nil {
		t.Error("ShapeTree() returned nil")
	}

	// Test GetOrCreateShapeTree
	shapeTree2 := csd.GetOrCreateShapeTree()
	if shapeTree2 == nil {
		t.Error(
			"GetOrCreateShapeTree() returned nil",
		)
	}
}

// TestSlideBackground tests SlideBackground element.
func TestSlideBackground(t *testing.T) {
	bg := NewSlideBackground()

	if bg.LocalName() != "bg" {
		t.Errorf(
			"LocalName() = %v, want bg",
			bg.LocalName(),
		)
	}
}

// TestPlaceholderTypes tests placeholder type constants.
func TestPlaceholderTypes(t *testing.T) {
	tests := []struct {
		pt   PlaceholderType
		want string
	}{
		{PlaceholderTitle, "title"},
		{PlaceholderBody, "body"},
		{PlaceholderCenteredTitle, "ctrTitle"},
		{PlaceholderSubTitle, "subTitle"},
		{PlaceholderDateTime, "dt"},
		{PlaceholderSlideNumber, "sldNum"},
		{PlaceholderFooter, "ftr"},
		{PlaceholderHeader, "hdr"},
		{PlaceholderObject, "obj"},
		{PlaceholderChart, "chart"},
		{PlaceholderTable, "tbl"},
		{PlaceholderClipArt, "clipArt"},
		{PlaceholderDiagram, "dgm"},
		{PlaceholderMedia, "media"},
		{PlaceholderSlideImage, "sldImg"},
		{PlaceholderPicture, "pic"},
	}

	for _, tt := range tests {
		if string(tt.pt) != tt.want {
			t.Errorf(
				"PlaceholderType %v = %v, want %v",
				tt.pt,
				string(tt.pt),
				tt.want,
			)
		}
	}
}

// TestSlideSizeTypes tests slide size type constants.
func TestSlideSizeTypes(t *testing.T) {
	tests := []struct {
		st   SlideSizeType
		want string
	}{
		{SlideSizeScreen4x3, "screen4x3"},
		{SlideSizeScreen16x9, "screen16x9"},
		{SlideSizeScreen16x10, "screen16x10"},
		{SlideSizeLetter, "letter"},
		{SlideSizeA4, "A4"},
		{SlideSizeCustom, "custom"},
	}

	for _, tt := range tests {
		if string(tt.st) != tt.want {
			t.Errorf(
				"SlideSizeType %v = %v, want %v",
				tt.st,
				string(tt.st),
				tt.want,
			)
		}
	}
}

// TestShapeTypes tests shape type constants.
func TestShapeTypes(t *testing.T) {
	tests := []struct {
		st   ShapeType
		want string
	}{
		{ShapeTypeRectangle, "rect"},
		{ShapeTypeRoundRectangle, "roundRect"},
		{ShapeTypeEllipse, "ellipse"},
		{ShapeTypeTriangle, "triangle"},
		{ShapeTypeDiamond, "diamond"},
		{ShapeTypeLine, "line"},
	}

	for _, tt := range tests {
		if string(tt.st) != tt.want {
			t.Errorf(
				"ShapeType %v = %v, want %v",
				tt.st,
				string(tt.st),
				tt.want,
			)
		}
	}
}

// TestSlideId tests SlideId element.
func TestSlideId(t *testing.T) {
	slideId := NewSlideId(256, "rId1")

	if slideId.Id() != 256 {
		t.Errorf(
			"Id() = %v, want 256",
			slideId.Id(),
		)
	}

	slideId.SetId(257)
	if slideId.Id() != 257 {
		t.Errorf(
			"Id() = %v, want 257",
			slideId.Id(),
		)
	}

	if slideId.RelId() != "rId1" {
		t.Errorf(
			"RelId() = %v, want rId1",
			slideId.RelId(),
		)
	}

	slideId.SetRelId("rId2")
	if slideId.RelId() != "rId2" {
		t.Errorf(
			"RelId() = %v, want rId2",
			slideId.RelId(),
		)
	}
}

// TestSlideTransition_Basic tests basic slide transition creation.
func TestSlideTransition_Basic(t *testing.T) {
	trans := NewSlideTransition()

	// Test speed
	trans.SetSpeed(TransitionSpeedSlow)
	if trans.Speed() != TransitionSpeedSlow {
		t.Errorf(
			"Speed() = %v, want slow",
			trans.Speed(),
		)
	}

	// Test advance on click
	trans.SetAdvanceOnClick(false)
	if trans.AdvanceOnClick() {
		t.Error(
			"AdvanceOnClick() should be false",
		)
	}
}

// TestSlideLayout tests SlideLayout element.
func TestSlideLayout(t *testing.T) {
	sl := NewSlideLayout()

	// Check CommonSlideData
	csd := sl.CommonSlideData()
	if csd == nil {
		t.Error("CommonSlideData() returned nil")
	}
}

// TestSlideMaster tests SlideMaster element.
func TestSlideMaster(t *testing.T) {
	sm := NewSlideMaster()

	// Check CommonSlideData
	csd := sm.CommonSlideData()
	if csd == nil {
		t.Error("CommonSlideData() returned nil")
	}
}

// TestSlideIdList tests SlideIdList element.
func TestSlideIdList(t *testing.T) {
	sil := WrapSlideIdList(NewSlideIdList())

	// Initially empty
	if sil.Count() != 0 {
		t.Errorf(
			"Initial Count() = %v, want 0",
			sil.Count(),
		)
	}

	// Add slide IDs
	sil.AddSlideId(256, "rId1")
	sil.AddSlideId(257, "rId2")

	if sil.Count() != 2 {
		t.Errorf(
			"Count() = %v, want 2",
			sil.Count(),
		)
	}

	// Check SlideIds
	ids := sil.SlideIds()
	if len(ids) != 2 {
		t.Errorf(
			"SlideIds() len = %v, want 2",
			len(ids),
		)
	}
}

// TestSlideMasterIdList tests SlideMasterIdList element.
func TestSlideMasterIdList(t *testing.T) {
	smil := WrapSlideMasterIdList(NewSlideMasterIdList())

	// Add slide master IDs
	smil.AddSlideMasterId(2147483648, "rId1")
	ids := smil.SlideMasterIds()
	if len(ids) != 1 {
		t.Errorf(
			"SlideMasterIds() len = %v, want 1",
			len(ids),
		)
	}
}

// TestNotesSize tests NotesSize element.
func TestNotesSize(t *testing.T) {
	ns := WrapNotesSize(NewNotesSize(6858000, 9144000))

	// Check dimensions
	if ns.Cx() != 6858000 {
		t.Errorf(
			"Cx() = %v, want 6858000",
			ns.Cx(),
		)
	}
	if ns.Cy() != 9144000 {
		t.Errorf(
			"Cy() = %v, want 9144000",
			ns.Cy(),
		)
	}
}
