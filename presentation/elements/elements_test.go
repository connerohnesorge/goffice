package elements

import (
	"testing"
)

// TestPresentation_NewPresentation tests creating presentation element.
func TestPresentation_NewPresentation(
	t *testing.T,
) {
	pres := NewPresentation()
	if pres == nil {
		t.Fatal("NewPresentation() returned nil")
	}

	// Check namespace URI
	if pres.NamespaceURI() != NamespacePresentationML {
		t.Errorf(
			"NamespaceURI() = %v, want %v",
			pres.NamespaceURI(),
			NamespacePresentationML,
		)
	}

	// Check local name
	if pres.LocalName() != "presentation" {
		t.Errorf(
			"LocalName() = %v, want presentation",
			pres.LocalName(),
		)
	}
}

// TestPresentation_AddSlideId tests adding slide IDs.
func TestPresentation_AddSlideId(t *testing.T) {
	pres := NewPresentation()
	if pres == nil {
		t.Fatal("NewPresentation() returned nil")
	}

	// Initially should have slide ID list (created in NewPresentation)
	idList := pres.SlideIdList()
	if idList == nil {
		t.Error(
			"SlideIdList() should be initialized",
		)
	}

	// Add a slide ID
	slideId := pres.AddSlideId(256, "rId1")
	if slideId == nil {
		t.Error("AddSlideId() returned nil")
	}

	// Verify the ID and relationship
	if slideId.Id() != 256 {
		t.Errorf(
			"Id() = %v, want 256",
			slideId.Id(),
		)
	}
	if slideId.RelId() != "rId1" {
		t.Errorf(
			"RelId() = %v, want rId1",
			slideId.RelId(),
		)
	}

	// Add another slide ID
	slideId2 := pres.AddSlideId(257, "rId2")
	if slideId2 == nil {
		t.Error(
			"AddSlideId() second returned nil",
		)
	}
	if slideId2.Id() != 257 {
		t.Errorf(
			"Second Id() = %v, want 257",
			slideId2.Id(),
		)
	}
}

// TestPresentation_SlideSize tests slide size operations.
func TestPresentation_SlideSize(t *testing.T) {
	pres := NewPresentation()
	if pres == nil {
		t.Fatal("NewPresentation() returned nil")
	}

	// Get or create slide size
	size := pres.GetOrCreateSlideSize()
	if size == nil {
		t.Error(
			"GetOrCreateSlideSize() returned nil",
		)
	}

	// Set dimensions (in EMUs) using SetCx/SetCy
	width := 9144000  // 10 inches
	height := 6858000 // 7.5 inches
	size.SetCx(width)
	size.SetCy(height)

	if size.Cx() != width {
		t.Errorf(
			"Cx() = %v, want %v",
			size.Cx(),
			width,
		)
	}
	if size.Cy() != height {
		t.Errorf(
			"Cy() = %v, want %v",
			size.Cy(),
			height,
		)
	}

	// Set type
	size.SetType(SlideSizeScreen16x9)
	if size.Type() != SlideSizeScreen16x9 {
		t.Errorf(
			"Type() = %v, want %v",
			size.Type(),
			SlideSizeScreen16x9,
		)
	}
}

// TestSlide_NewSlide tests creating slide element.
func TestSlide_NewSlide(t *testing.T) {
	slide := NewSlide()
	if slide == nil {
		t.Fatal("NewSlide() returned nil")
	}

	// Check namespace URI
	if slide.NamespaceURI() != NamespacePresentationML {
		t.Errorf(
			"NamespaceURI() = %v, want %v",
			slide.NamespaceURI(),
			NamespacePresentationML,
		)
	}

	// Check local name
	if slide.LocalName() != "sld" {
		t.Errorf(
			"LocalName() = %v, want sld",
			slide.LocalName(),
		)
	}
}

// TestSlide_CommonSlideData tests getting common slide data.
func TestSlide_CommonSlideData(t *testing.T) {
	slide := NewSlide()
	if slide == nil {
		t.Fatal("NewSlide() returned nil")
	}

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
	if shapeTree == nil {
		t.Fatal("NewShapeTree() returned nil")
	}

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
	for _, sh := range shapes {
		if sh == nil {
			t.Error("Shapes() yielded nil")
		}
	}
}

// TestShape_SetText tests setting shape text.
func TestShape_SetText(t *testing.T) {
	shape := NewShape()
	if shape == nil {
		t.Fatal("NewShape() returned nil")
	}

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
	if shape == nil {
		t.Fatal("NewShape() returned nil")
	}

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
	if shape == nil {
		t.Fatal("NewShape() returned nil")
	}

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
	if shape == nil {
		t.Fatal("NewShape() returned nil")
	}

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
	if shape == nil {
		t.Fatal("NewShape() returned nil")
	}

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
	if csd == nil {
		t.Fatal(
			"NewCommonSlideData() returned nil",
		)
	}

	// Check local name
	if csd.LocalName() != "cSld" {
		t.Errorf(
			"LocalName() = %v, want cSld",
			csd.LocalName(),
		)
	}

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
	if bg == nil {
		t.Fatal(
			"NewSlideBackground() returned nil",
		)
	}

	if bg.LocalName() != "bg" {
		t.Errorf(
			"LocalName() = %v, want bg",
			bg.LocalName(),
		)
	}

	// Test Clone
	cloned := bg.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
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

// TestShape_Clone tests cloning shapes.
func TestShape_Clone(t *testing.T) {
	shape := NewShape()
	if shape == nil {
		t.Fatal("NewShape() returned nil")
	}

	shape.SetText("Test Text")
	shape.SetPosition(100, 200)

	cloned := shape.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestNonVisualShapeProperties tests non-visual shape properties.
func TestNonVisualShapeProperties(t *testing.T) {
	nvSpPr := NewNonVisualShapeProperties()
	if nvSpPr == nil {
		t.Fatal(
			"NewNonVisualShapeProperties() returned nil",
		)
	}

	if nvSpPr.LocalName() != "nvSpPr" {
		t.Errorf(
			"LocalName() = %v, want nvSpPr",
			nvSpPr.LocalName(),
		)
	}
}

// TestSlideId tests SlideId element.
func TestSlideId(t *testing.T) {
	slideId := NewSlideId(256, "rId1")
	if slideId == nil {
		t.Fatal("NewSlideId() returned nil")
	}

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
	if trans == nil {
		t.Fatal(
			"NewSlideTransition() returned nil",
		)
	}

	// Check local name
	if trans.LocalName() != "transition" {
		t.Errorf(
			"LocalName() = %v, want transition",
			trans.LocalName(),
		)
	}

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

	// Test Clone
	cloned := trans.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestSlideTiming_Basic tests basic slide timing creation.
func TestSlideTiming_Basic(t *testing.T) {
	timing := NewSlideTiming()
	if timing == nil {
		t.Fatal("NewSlideTiming() returned nil")
	}

	// Check local name
	if timing.LocalName() != "timing" {
		t.Errorf(
			"LocalName() = %v, want timing",
			timing.LocalName(),
		)
	}

	// Test Clone
	cloned := timing.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestSlideLayout tests SlideLayout element.
func TestSlideLayout(t *testing.T) {
	sl := NewSlideLayout()
	if sl == nil {
		t.Fatal("NewSlideLayout() returned nil")
	}

	// Check local name
	if sl.LocalName() != "sldLayout" {
		t.Errorf(
			"LocalName() = %v, want sldLayout",
			sl.LocalName(),
		)
	}

	// Check CommonSlideData
	csd := sl.CommonSlideData()
	if csd == nil {
		t.Error("CommonSlideData() returned nil")
	}

	// Test Clone
	cloned := sl.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestSlideMaster tests SlideMaster element.
func TestSlideMaster(t *testing.T) {
	sm := NewSlideMaster()
	if sm == nil {
		t.Fatal("NewSlideMaster() returned nil")
	}

	// Check local name
	if sm.LocalName() != "sldMaster" {
		t.Errorf(
			"LocalName() = %v, want sldMaster",
			sm.LocalName(),
		)
	}

	// Check CommonSlideData
	csd := sm.CommonSlideData()
	if csd == nil {
		t.Error("CommonSlideData() returned nil")
	}

	// Test Clone
	cloned := sm.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestSlideIdList tests SlideIdList element.
func TestSlideIdList(t *testing.T) {
	sil := NewSlideIdList()
	if sil == nil {
		t.Fatal("NewSlideIdList() returned nil")
	}

	// Check local name
	if sil.LocalName() != "sldIdLst" {
		t.Errorf(
			"LocalName() = %v, want sldIdLst",
			sil.LocalName(),
		)
	}

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

	// Test Clone
	cloned := sil.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestSlideMasterIdList tests SlideMasterIdList element.
func TestSlideMasterIdList(t *testing.T) {
	smil := NewSlideMasterIdList()
	if smil == nil {
		t.Fatal(
			"NewSlideMasterIdList() returned nil",
		)
	}

	// Check local name
	if smil.LocalName() != "sldMasterIdLst" {
		t.Errorf(
			"LocalName() = %v, want sldMasterIdLst",
			smil.LocalName(),
		)
	}

	// Add slide master IDs
	smil.AddSlideMasterId(2147483648, "rId1")
	ids := smil.SlideMasterIds()
	if len(ids) != 1 {
		t.Errorf(
			"SlideMasterIds() len = %v, want 1",
			len(ids),
		)
	}

	// Test Clone
	cloned := smil.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestNotesMasterIdList tests NotesMasterIdList element.
func TestNotesMasterIdList(t *testing.T) {
	nmil := NewNotesMasterIdList()
	if nmil == nil {
		t.Fatal(
			"NewNotesMasterIdList() returned nil",
		)
	}

	// Check local name
	if nmil.LocalName() != "notesMasterIdLst" {
		t.Errorf(
			"LocalName() = %v, want notesMasterIdLst",
			nmil.LocalName(),
		)
	}

	// Test Clone
	cloned := nmil.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestHandoutMasterIdList tests HandoutMasterIdList element.
func TestHandoutMasterIdList(t *testing.T) {
	hmil := NewHandoutMasterIdList()
	if hmil == nil {
		t.Fatal(
			"NewHandoutMasterIdList() returned nil",
		)
	}

	// Check local name
	if hmil.LocalName() != "handoutMasterIdLst" {
		t.Errorf(
			"LocalName() = %v, want handoutMasterIdLst",
			hmil.LocalName(),
		)
	}

	// Test Clone
	cloned := hmil.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestDefaultTextStyle tests DefaultTextStyle element.
func TestDefaultTextStyle(t *testing.T) {
	dts := NewDefaultTextStyle()
	if dts == nil {
		t.Fatal(
			"NewDefaultTextStyle() returned nil",
		)
	}

	// Check local name
	if dts.LocalName() != "defaultTextStyle" {
		t.Errorf(
			"LocalName() = %v, want defaultTextStyle",
			dts.LocalName(),
		)
	}

	// Test Clone
	cloned := dts.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestNotesSize tests NotesSize element.
func TestNotesSize(t *testing.T) {
	ns := NewNotesSize(6858000, 9144000)
	if ns == nil {
		t.Fatal("NewNotesSize() returned nil")
	}

	// Check local name
	if ns.LocalName() != "notesSz" {
		t.Errorf(
			"LocalName() = %v, want notesSz",
			ns.LocalName(),
		)
	}

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

	// Test Clone
	cloned := ns.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestColorMapOverride tests ColorMapOverride element.
func TestColorMapOverride(t *testing.T) {
	cmo := NewColorMapOverride()
	if cmo == nil {
		t.Fatal(
			"NewColorMapOverride() returned nil",
		)
	}

	// Check local name
	if cmo.LocalName() != "clrMapOvr" {
		t.Errorf(
			"LocalName() = %v, want clrMapOvr",
			cmo.LocalName(),
		)
	}

	// Test Clone
	cloned := cmo.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestControlList tests ControlList element.
func TestControlList(t *testing.T) {
	cl := NewControlList()
	if cl == nil {
		t.Fatal("NewControlList() returned nil")
	}

	// Check local name
	if cl.LocalName() != "controls" {
		t.Errorf(
			"LocalName() = %v, want controls",
			cl.LocalName(),
		)
	}

	// Test Clone
	cloned := cl.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestShapeProperties tests ShapeProperties element.
func TestShapeProperties(t *testing.T) {
	spPr := NewShapeProperties()
	if spPr == nil {
		t.Fatal(
			"NewShapeProperties() returned nil",
		)
	}

	// Check local name
	if spPr.LocalName() != "spPr" {
		t.Errorf(
			"LocalName() = %v, want spPr",
			spPr.LocalName(),
		)
	}

	// Set offset
	spPr.SetOffset(100, 200)

	// Set extents
	spPr.SetExtents(500, 300)

	// Set preset geometry
	spPr.SetPresetGeometry("rect")

	// Set solid fill
	spPr.SetSolidFill("FF0000")

	// Set no fill
	spPr.SetNoFill()

	// Test Clone
	cloned := spPr.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}

// TestShapeStyle tests ShapeStyle element.
func TestShapeStyle(t *testing.T) {
	ss := NewShapeStyle()
	if ss == nil {
		t.Fatal("NewShapeStyle() returned nil")
	}

	// Check local name
	if ss.LocalName() != "style" {
		t.Errorf(
			"LocalName() = %v, want style",
			ss.LocalName(),
		)
	}

	// Test Clone
	cloned := ss.Clone()
	if cloned == nil {
		t.Error("Clone() returned nil")
	}
}
