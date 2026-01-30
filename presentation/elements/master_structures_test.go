package elements_test

import (
	"testing"

	"github.com/connerohnesorge/goffice/presentation/elements"
)

func TestSlideMasterProperties(t *testing.T) {
	sm := elements.NewSlideMaster()

	// Test ColorMap
	cm := sm.ColorMap()
	if cm != nil {
		t.Errorf("Expected nil ColorMap initially, got %+v", cm)
	}
	newCm := elements.NewPresentationColorMap()
	sm.AppendChild(newCm)
	cm = sm.ColorMap()
	if cm == nil {
		t.Error("Expected non-nil ColorMap after adding")
	}

	// Test SlideLayoutIdList
	slil := sm.SlideLayoutIdList()
	if slil != nil {
		t.Errorf("Expected nil SlideLayoutIdList initially, got %+v", slil)
	}
	newSlil := elements.NewSlideLayoutIdList()
	sm.AppendChild(newSlil)
	slil = sm.SlideLayoutIdList()
	if slil == nil {
		t.Error("Expected non-nil SlideLayoutIdList after adding")
	}

	// Test Transition
	tr := sm.Transition()
	if tr != nil {
		t.Errorf("Expected nil Transition initially, got %+v", tr)
	}
	newTr := elements.NewSlideTransition()
	sm.AppendChild(newTr)
	tr = sm.Transition()
	if tr == nil {
		t.Error("Expected non-nil Transition after adding")
	}

	// Test Timing
	tm := sm.Timing()
	if tm != nil {
		t.Errorf("Expected nil Timing initially, got %+v", tm)
	}
	newTm := elements.NewSlideTiming()
	sm.AppendChild(newTm)
	tm = sm.Timing()
	if tm == nil {
		t.Error("Expected non-nil Timing after adding")
	}

	// Test HeaderFooter
	hf := sm.HeaderFooter()
	if hf != nil {
		t.Errorf("Expected nil HeaderFooter initially, got %+v", hf)
	}
	newHf := elements.NewExtHeaderFooter()
	sm.AppendChild(newHf)
	hf = sm.HeaderFooter()
	if hf == nil {
		t.Fatal("Expected non-nil HeaderFooter after adding")
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
	newTs := elements.NewTextStyles()
	sm.AppendChild(newTs)
	ts = sm.TextStyles()
	if ts == nil {
		t.Error("Expected non-nil TextStyles after adding")
	}

	// Test ExtensionList
	extLst := sm.ExtensionList()
	if extLst != nil {
		t.Errorf("Expected nil ExtensionList initially, got %+v", extLst)
	}
	newExtLst := elements.NewExtensionListModify()
	sm.AppendChild(newExtLst)
	extLst = sm.ExtensionList()
	if extLst == nil {
		t.Error("Expected non-nil ExtensionList after adding")
	}
}

func TestSlideLayoutProperties(t *testing.T) {
	sl := elements.NewSlideLayout()

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
	newCmo := elements.NewColorMapOverride()
	sl.AppendChild(newCmo)
	cmo = sl.ColorMapOverride()
	if cmo == nil {
		t.Error("Expected non-nil ColorMapOverride after adding")
	}

	// Test Transition
	tr := sl.Transition()
	if tr != nil {
		t.Errorf("Expected nil Transition initially, got %+v", tr)
	}
	newTr := elements.NewSlideTransition()
	sl.AppendChild(newTr)
	tr = sl.Transition()
	if tr == nil {
		t.Error("Expected non-nil Transition after adding")
	}

	// Test Timing
	tm := sl.Timing()
	if tm != nil {
		t.Errorf("Expected nil Timing initially, got %+v", tm)
	}
	newTm := elements.NewSlideTiming()
	sl.AppendChild(newTm)
	tm = sl.Timing()
	if tm == nil {
		t.Error("Expected non-nil Timing after adding")
	}

	// Test HeaderFooter
	hf := sl.HeaderFooter()
	if hf != nil {
		t.Errorf("Expected nil HeaderFooter initially, got %+v", hf)
	}
	newHf := elements.NewExtHeaderFooter()
	sl.AppendChild(newHf)
	hf = sl.HeaderFooter()
	if hf == nil {
		t.Fatal("Expected non-nil HeaderFooter after adding")
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
	newExtLst := elements.NewExtensionListModify()
	sl.AppendChild(newExtLst)
	extLst = sl.ExtensionList()
	if extLst == nil {
		t.Error("Expected non-nil ExtensionList after adding")
	}
}

func TestNotesSlideProperties(t *testing.T) {
	ns := elements.NewNotesSlide()

	// Test HeaderFooter
	hf := ns.HeaderFooter()
	if hf != nil {
		t.Errorf("Expected nil HeaderFooter initially, got %+v", hf)
	}
	newHf := elements.NewExtHeaderFooter()
	ns.AppendChild(newHf)
	hf = ns.HeaderFooter()
	if hf == nil {
		t.Fatal("Expected non-nil HeaderFooter after adding")
	}
	hf.SetDateTime(true)
	if !hf.DateTime.Value() {
		t.Errorf("Expected DateTime to be true, got false")
	}
}

func TestPlaceholderShapeProperties(t *testing.T) {
	ph := elements.NewPlaceholderShape()

	ph.SetOrientation(elements.DirectionValuesVert)
	if ph.Orientation.Value() != elements.DirectionValuesVert {
		t.Errorf("Expected orientation to be %s, got %s", elements.DirectionValuesVert, ph.Orientation.Value())
	}

	ph.SetSize(elements.PlaceholderSizeValuesFull)
	if ph.Size.Value() != elements.PlaceholderSizeValuesFull {
		t.Errorf("Expected size to be %s, got %s", elements.PlaceholderSizeValuesFull, ph.Size.Value())
	}

	ph.SetHasCustomPrompt(true)
	if !ph.HasCustomPrompt.Value() {
		t.Errorf("Expected hasCustomPrompt to be true, got false")
	}
}

func TestTextStylesParagraphProperties(t *testing.T) {
	ts := elements.NewTextStyles()

	// Test TitleStyle
	if ts.TitleStyle == nil {
		ts.TitleStyle = elements.NewTitleStyle()
	}
	ts.TitleStyle.GetOrCreateDefaultParagraphProperties().SetAlignment(elements.TextAlignCenter)
	if ts.TitleStyle.DefaultParagraphProperties().Alignment() != elements.TextAlignCenter {
		t.Errorf("Expected title style alignment to be %s, got %s", elements.TextAlignCenter, ts.TitleStyle.DefaultParagraphProperties().Alignment())
	}

	// Test BodyStyle
	if ts.BodyStyle == nil {
		ts.BodyStyle = elements.NewBodyStyle()
	}
	ts.BodyStyle.GetOrCreateLevel1ParagraphProperties().SetLeftMargin(1000)
	if ts.BodyStyle.Level1ParagraphProperties().LeftMargin() != 1000 {
		t.Errorf("Expected body style level 1 left margin to be 1000, got %d", ts.BodyStyle.Level1ParagraphProperties().LeftMargin())
	}
}
