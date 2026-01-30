package elements_test

import (
	"testing"

	. "github.com/connerohnesorge/goffice/presentation/elements"
)

func TestSlideMasterProperties(t *testing.T) {
	sm := NewSlideMaster()

	// Test ColorMap
	cm := sm.ColorMap()
	if cm != nil {
		t.Errorf("Expected nil ColorMap initially, got %+v", cm)
	}
	newCm := NewPresentationColorMap()
	sm.AppendChild(newCm)
	cm = sm.ColorMap()
	if cm == nil {
		t.Errorf("Expected non-nil ColorMap after adding")
	}

	// Test SlideLayoutIdList
	slil := sm.SlideLayoutIdList()
	if slil != nil {
		t.Errorf("Expected nil SlideLayoutIdList initially, got %+v", slil)
	}
	newSlil := NewSlideLayoutIdList()
	sm.AppendChild(newSlil)
	slil = sm.SlideLayoutIdList()
	if slil == nil {
		t.Errorf("Expected non-nil SlideLayoutIdList after adding")
	}

	// Test Transition
	tr := sm.Transition()
	if tr != nil {
		t.Errorf("Expected nil Transition initially, got %+v", tr)
	}
	newTr := NewSlideTransition()
	sm.AppendChild(newTr)
	tr = sm.Transition()
	if tr == nil {
		t.Errorf("Expected non-nil Transition after adding")
	}

	// Test Timing
	tm := sm.Timing()
	if tm != nil {
		t.Errorf("Expected nil Timing initially, got %+v", tm)
	}
	newTm := NewSlideTiming()
	sm.AppendChild(newTm)
	tm = sm.Timing()
	if tm == nil {
		t.Errorf("Expected non-nil Timing after adding")
	}

	// Test HeaderFooter
	hf := sm.HeaderFooter()
	if hf != nil {
		t.Errorf("Expected nil HeaderFooter initially, got %+v", hf)
	}
	newHf := NewExtHeaderFooter()
	sm.AppendChild(newHf)
	hf = sm.HeaderFooter()
	if hf == nil {
		t.Fatalf("Expected non-nil HeaderFooter after adding")
	}
	hf.SetSlideNumber(true)
	if !hf.SlideNumber.Value() {
		t.Errorf("Expected SlideNumber to be true, got false")
	}

	// Test TextStyles
	ts := sm.TextStyles()
	if ts != nil {
		t.Errorf("Expected nil TextStyles initially, got %+v", ts)
	}
	newTs := NewTextStyles()
	sm.AppendChild(newTs)
	ts = sm.TextStyles()
	if ts == nil {
		t.Errorf("Expected non-nil TextStyles after adding")
	}

	// Test ExtensionList
	extLst := sm.ExtensionList()
	if extLst != nil {
		t.Errorf("Expected nil ExtensionList initially, got %+v", extLst)
	}
	newExtLst := NewExtensionListModify()
	sm.AppendChild(newExtLst)
	extLst = sm.ExtensionList()
	if extLst == nil {
		t.Errorf("Expected non-nil ExtensionList after adding")
	}
}

func TestSlideLayoutProperties(t *testing.T) {
	sl := NewSlideLayout()

	// Test Name/SetName
	sl.SetName("My Custom Layout")
	if sl.Name() != "My Custom Layout" {
		t.Errorf("Expected layout name to be 'My Custom Layout', got '%s'", sl.Name())
	}

	// Test ColorMapOverride
	cmo := sl.ColorMapOverride()
	if cmo != nil {
		t.Errorf("Expected nil ColorMapOverride initially, got %+v", cmo)
	}
	newCmo := NewColorMapOverride()
	sl.AppendChild(newCmo)
	cmo = sl.ColorMapOverride()
	if cmo == nil {
		t.Errorf("Expected non-nil ColorMapOverride after adding")
	}

	// Test Transition
	tr := sl.Transition()
	if tr != nil {
		t.Errorf("Expected nil Transition initially, got %+v", tr)
	}
	newTr := NewSlideTransition()
	sl.AppendChild(newTr)
	tr = sl.Transition()
	if tr == nil {
		t.Errorf("Expected non-nil Transition after adding")
	}

	// Test Timing
	tm := sl.Timing()
	if tm != nil {
		t.Errorf("Expected nil Timing initially, got %+v", tm)
	}
	newTm := NewSlideTiming()
	sl.AppendChild(newTm)
	tm = sl.Timing()
	if tm == nil {
		t.Errorf("Expected non-nil Timing after adding")
	}

	// Test HeaderFooter
	hf := sl.HeaderFooter()
	if hf != nil {
		t.Errorf("Expected nil HeaderFooter initially, got %+v", hf)
	}
	newHf := NewExtHeaderFooter()
	sl.AppendChild(newHf)
	hf = sl.HeaderFooter()
	if hf == nil {
		t.Fatalf("Expected non-nil HeaderFooter after adding")
	}
	hf.SetFooter(true)
	if !hf.Footer.Value() {
		t.Errorf("Expected Footer to be true, got false")
	}

	// Test ExtensionList
	extLst := sl.ExtensionList()
	if extLst != nil {
		t.Errorf("Expected nil ExtensionList initially, got %+v", extLst)
	}
	newExtLst := NewExtensionListModify()
	sl.AppendChild(newExtLst)
	extLst = sl.ExtensionList()
	if extLst == nil {
		t.Errorf("Expected non-nil ExtensionList after adding")
	}
}

func TestNotesSlideProperties(t *testing.T) {
	ns := NewNotesSlide()

	// Test HeaderFooter
	hf := ns.HeaderFooter()
	if hf != nil {
		t.Errorf("Expected nil HeaderFooter initially, got %+v", hf)
	}
	newHf := NewExtHeaderFooter()
	ns.AppendChild(newHf)
	hf = ns.HeaderFooter()
	if hf == nil {
		t.Fatalf("Expected non-nil HeaderFooter after adding")
	}
	hf.SetDateTime(true)
	if !hf.DateTime.Value() {
		t.Errorf("Expected DateTime to be true, got false")
	}
}

func TestPlaceholderShapeProperties(t *testing.T) {
	ph := NewPlaceholderShape()

	ph.SetOrientation(DirectionValuesVert)
	if ph.Orientation.Value() != DirectionValuesVert {
		t.Errorf("Expected orientation to be %s, got %s", DirectionValuesVert, ph.Orientation.Value())
	}

	ph.SetSize(PlaceholderSizeValuesFull)
	if ph.Size.Value() != PlaceholderSizeValuesFull {
		t.Errorf("Expected size to be %s, got %s", PlaceholderSizeValuesFull, ph.Size.Value())
	}

	ph.SetHasCustomPrompt(true)
	if !ph.HasCustomPrompt.Value() {
		t.Errorf("Expected hasCustomPrompt to be true, got false")
	}
}

func TestTextStylesParagraphProperties(t *testing.T) {
	ts := NewTextStyles()

	// Test TitleStyle
	if ts.TitleStyle == nil {
		ts.TitleStyle = NewTitleStyle()
	}
	ts.TitleStyle.GetOrCreateDefaultParagraphProperties().SetAlignment(TextAlignCenter)
	if ts.TitleStyle.DefaultParagraphProperties().Alignment() != TextAlignCenter {
		t.Errorf("Expected title style alignment to be %s, got %s", TextAlignCenter, ts.TitleStyle.DefaultParagraphProperties().Alignment())
	}

	// Test BodyStyle
	if ts.BodyStyle == nil {
		ts.BodyStyle = NewBodyStyle()
	}
	ts.BodyStyle.GetOrCreateLevel1ParagraphProperties().SetLeftMargin(1000)
	if ts.BodyStyle.Level1ParagraphProperties().LeftMargin() != 1000 {
		t.Errorf("Expected body style level 1 left margin to be 1000, got %d", ts.BodyStyle.Level1ParagraphProperties().LeftMargin())
	}
}
