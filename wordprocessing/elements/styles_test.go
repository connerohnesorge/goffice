package elements

import (
	"strings"
	"testing"
)

func TestStylesRoot(t *testing.T) {
	styles := NewStyles()
	if styles == nil {
		t.Fatal("NewStyles() returned nil")
	}
	if styles.LocalName() != "styles" {
		t.Errorf("Expected local name 'styles', got '%s'", styles.LocalName())
	}
}

func TestDocDefaults(t *testing.T) {
	styles := NewStyles()

	// Initially nil
	if styles.DocDefaults() != nil {
		t.Error("Expected DocDefaults() to be nil initially")
	}

	// GetOrCreate should create it
	dd := styles.GetOrCreateDocDefaults()
	if dd == nil {
		t.Fatal("GetOrCreateDocDefaults() returned nil")
	}

	// Should be first child
	if styles.DocDefaults() == nil {
		t.Error("DocDefaults() should not be nil after creation")
	}
}

func TestLatentStyles(t *testing.T) {
	styles := NewStyles()

	// Initially nil
	if styles.LatentStyles() != nil {
		t.Error("Expected LatentStyles() to be nil initially")
	}

	// GetOrCreate should create it
	ls := styles.GetOrCreateLatentStyles()
	if ls == nil {
		t.Fatal("GetOrCreateLatentStyles() returned nil")
	}

	// Set some defaults
	ls.SetDefSemiHidden(true)
	ls.SetDefUnhideWhenUsed(true)
	ls.SetDefUIPriority(99)

	if !ls.DefSemiHidden() {
		t.Error("DefSemiHidden() should be true")
	}
	if !ls.DefUnhideWhenUsed() {
		t.Error("DefUnhideWhenUsed() should be true")
	}
	if ls.DefUIPriority() != 99 {
		t.Errorf("Expected DefUIPriority() 99, got %d", ls.DefUIPriority())
	}
}

func TestStyleCreation(t *testing.T) {
	style := NewStyle("TestStyle", StyleTypeParagraph)
	if style == nil {
		t.Fatal("NewStyle() returned nil")
	}

	if style.StyleId() != "TestStyle" {
		t.Errorf("Expected StyleId 'TestStyle', got '%s'", style.StyleId())
	}
	if style.Type() != StyleTypeParagraph {
		t.Errorf("Expected Type 'paragraph', got '%s'", style.Type())
	}
}

func TestStyleProperties(t *testing.T) {
	style := NewParagraphStyle("MyStyle", "My Custom Style")

	if style.StyleId() != "MyStyle" {
		t.Errorf("Expected StyleId 'MyStyle', got '%s'", style.StyleId())
	}
	if style.StyleName() != "My Custom Style" {
		t.Errorf("Expected StyleName 'My Custom Style', got '%s'", style.StyleName())
	}

	// Test based on
	style.SetBasedOn(StyleIdNormal)
	if style.BasedOn() != StyleIdNormal {
		t.Errorf("Expected BasedOn '%s', got '%s'", StyleIdNormal, style.BasedOn())
	}

	// Test next style
	style.SetNextParagraphStyle(StyleIdNormal)
	if style.NextParagraphStyle() != StyleIdNormal {
		t.Errorf("Expected NextParagraphStyle '%s', got '%s'", StyleIdNormal, style.NextParagraphStyle())
	}

	// Test default flag
	style.SetDefault(true)
	if !style.Default() {
		t.Error("Expected Default() to be true")
	}

	// Test custom style flag
	style.SetCustomStyle(true)
	if !style.CustomStyle() {
		t.Error("Expected CustomStyle() to be true")
	}

	// Test UI priority
	style.SetUIPriority(10)
	if style.UIPriority() != 10 {
		t.Errorf("Expected UIPriority 10, got %d", style.UIPriority())
	}

	// Test quick format
	style.SetQuickFormat(true)
	if !style.QuickFormat() {
		t.Error("Expected QuickFormat() to be true")
	}

	// Test semi-hidden
	style.SetSemiHidden(true)
	if !style.SemiHidden() {
		t.Error("Expected SemiHidden() to be true")
	}

	// Test unhide when used
	style.SetUnhideWhenUsed(true)
	if !style.UnhideWhenUsed() {
		t.Error("Expected UnhideWhenUsed() to be true")
	}
}

func TestStylesCollection(t *testing.T) {
	styles := NewStyles()

	style1 := NewParagraphStyle("Style1", "Style One")
	style2 := NewParagraphStyle("Style2", "Style Two")
	style3 := NewCharacterStyle("Style3", "Style Three")

	styles.AddStyle(style1)
	styles.AddStyle(style2)
	styles.AddStyle(style3)

	// Test GetStyleById
	found := styles.GetStyleById("Style1")
	if found == nil {
		t.Error("GetStyleById('Style1') should not return nil")
	}
	if found.StyleName() != "Style One" {
		t.Errorf("Expected 'Style One', got '%s'", found.StyleName())
	}

	// Test GetStyleByName
	found = styles.GetStyleByName("Style Two")
	if found == nil {
		t.Error("GetStyleByName('Style Two') should not return nil")
	}
	if found.StyleId() != "Style2" {
		t.Errorf("Expected 'Style2', got '%s'", found.StyleId())
	}

	// Test GetStylesByType
	count := 0
	for range styles.GetStylesByType(StyleTypeParagraph) {
		count++
	}
	if count != 2 {
		t.Errorf("Expected 2 paragraph styles, got %d", count)
	}

	// Test iteration
	count = 0
	for range styles.Styles() {
		count++
	}
	if count != 3 {
		t.Errorf("Expected 3 styles, got %d", count)
	}
}

func TestHeadingStyleFactory(t *testing.T) {
	h1 := NewHeadingStyle(1)
	if h1 == nil {
		t.Fatal("NewHeadingStyle(1) returned nil")
	}

	if h1.StyleId() != "Heading1" {
		t.Errorf("Expected StyleId 'Heading1', got '%s'", h1.StyleId())
	}
	if h1.StyleName() != "Heading 1" {
		t.Errorf("Expected StyleName 'Heading 1', got '%s'", h1.StyleName())
	}
	if h1.BasedOn() != StyleIdNormal {
		t.Errorf("Expected BasedOn '%s', got '%s'", StyleIdNormal, h1.BasedOn())
	}
	if !h1.QuickFormat() {
		t.Error("Expected QuickFormat() to be true for heading style")
	}
}

func TestStyleParagraphProperties(t *testing.T) {
	style := NewParagraphStyle("Test", "Test Style")
	pp := style.GetOrCreateStyleParagraphProperties()
	if pp == nil {
		t.Fatal("GetOrCreateStyleParagraphProperties() returned nil")
	}

	pp.SetJustification(JustificationCenter)
	pp.SetKeepNext(true)
	pp.SetSpacingBefore(240)
	pp.SetSpacingAfter(120)

	// Verify XML output contains expected elements
	xml := style.OuterXml()
	if !strings.Contains(xml, "jc") {
		t.Error("Style XML should contain justification element")
	}
	if !strings.Contains(xml, "keepNext") {
		t.Error("Style XML should contain keepNext element")
	}
}

func TestStyleRunProperties(t *testing.T) {
	style := NewCharacterStyle("BoldStyle", "Bold Text")
	rp := style.GetOrCreateStyleRunProperties()
	if rp == nil {
		t.Fatal("GetOrCreateStyleRunProperties() returned nil")
	}

	rp.SetBold(true)
	rp.SetItalic(true)
	rp.SetFontSize(28) // 14pt
	rp.SetColor("FF0000")

	// Verify XML output
	xml := style.OuterXml()
	if !strings.Contains(xml, "<w:b") {
		t.Error("Style XML should contain bold element")
	}
	if !strings.Contains(xml, "<w:i") {
		t.Error("Style XML should contain italic element")
	}
	if !strings.Contains(xml, "<w:sz") {
		t.Error("Style XML should contain font size element")
	}
	if !strings.Contains(xml, "<w:color") {
		t.Error("Style XML should contain color element")
	}
}

func TestDocDefaultsRunProperties(t *testing.T) {
	dd := NewDocDefaults()
	rpd := dd.GetOrCreateRunPropertiesDefault()
	if rpd == nil {
		t.Fatal("GetOrCreateRunPropertiesDefault() returned nil")
	}

	rp := rpd.GetOrCreateRunProperties()
	if rp == nil {
		t.Fatal("GetOrCreateRunProperties() returned nil")
	}

	rp.SetFontSize(22) // 11pt
	rp.SetFont("Calibri")
	rp.SetLanguage("en-US")

	xml := dd.OuterXml()
	if !strings.Contains(xml, "rPrDefault") {
		t.Error("DocDefaults XML should contain rPrDefault")
	}
}

func TestLatentStyleException(t *testing.T) {
	ls := NewLatentStyles()
	ex := ls.AddException("Heading 1")

	ex.SetQFormat(true)
	ex.SetUIPriority(9)
	ex.SetSemiHidden(false)

	if ex.Name() != "Heading 1" {
		t.Errorf("Expected Name 'Heading 1', got '%s'", ex.Name())
	}
	if !ex.QFormat() {
		t.Error("Expected QFormat() to be true")
	}
	if ex.UIPriority() != 9 {
		t.Errorf("Expected UIPriority 9, got %d", ex.UIPriority())
	}

	// Find it
	found := ls.GetException("Heading 1")
	if found == nil {
		t.Error("GetException('Heading 1') should not return nil")
	}
}
