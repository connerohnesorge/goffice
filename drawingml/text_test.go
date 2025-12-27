package drawingml

import (
	"strings"
	"testing"
)

func TestNewTextBody(t *testing.T) {
	tb := NewTextBody()

	if tb == nil {
		t.Fatal("NewTextBody returned nil")
	}

	if tb.LocalName() != "txBody" {
		t.Errorf(
			"Expected localName 'txBody', got '%s'",
			tb.LocalName(),
		)
	}

	if tb.NamespaceURI() != NamespaceMain {
		t.Errorf(
			"Expected namespace '%s', got '%s'",
			NamespaceMain,
			tb.NamespaceURI(),
		)
	}

	// Should have body properties by default
	bp := tb.BodyProperties()
	if bp == nil {
		t.Error(
			"Expected default body properties",
		)
	}
}

func TestTextBodyAddParagraph(t *testing.T) {
	tb := NewTextBody()

	p := tb.AddParagraph("Hello World")
	if p == nil {
		t.Fatal("AddParagraph returned nil")
	}

	paragraphs := tb.Paragraphs()
	if len(paragraphs) != 1 {
		t.Errorf(
			"Expected 1 paragraph, got %d",
			len(paragraphs),
		)
	}

	text := p.GetText()
	if text != "Hello World" {
		t.Errorf(
			"Expected text 'Hello World', got '%s'",
			text,
		)
	}
}

func TestTextBodyMultipleParagraphs(
	t *testing.T,
) {
	tb := NewTextBody()

	tb.AddParagraph("First")
	tb.AddParagraph("Second")
	tb.AddParagraph("Third")

	paragraphs := tb.Paragraphs()
	if len(paragraphs) != 3 {
		t.Errorf(
			"Expected 3 paragraphs, got %d",
			len(paragraphs),
		)
	}

	expected := []string{
		"First",
		"Second",
		"Third",
	}
	for i, p := range paragraphs {
		if p.GetText() != expected[i] {
			t.Errorf(
				"Paragraph %d: expected '%s', got '%s'",
				i,
				expected[i],
				p.GetText(),
			)
		}
	}
}

func TestTextBodyClearParagraphs(t *testing.T) {
	tb := NewTextBody()
	tb.AddParagraph("One")
	tb.AddParagraph("Two")

	tb.ClearParagraphs()

	if len(tb.Paragraphs()) != 0 {
		t.Error(
			"Expected no paragraphs after clear",
		)
	}

	// Body properties should still exist
	if tb.BodyProperties() == nil {
		t.Error(
			"Body properties should remain after clearing paragraphs",
		)
	}
}

func TestTextBodyProperties(t *testing.T) {
	bp := NewTextBodyProperties()

	// Test rotation
	bp.SetRotationDegrees(45.0)
	if bp.RotationDegrees() != 45.0 {
		t.Errorf(
			"Expected rotation 45.0, got %f",
			bp.RotationDegrees(),
		)
	}

	// Test vertical orientation
	bp.SetVertical(TextVerticalRotate90)
	if bp.Vertical() != TextVerticalRotate90 {
		t.Errorf(
			"Expected vertical 'vert', got '%s'",
			bp.Vertical(),
		)
	}

	// Test wrap
	bp.SetWrap(TextWrapNone)
	if bp.Wrap() != TextWrapNone {
		t.Errorf(
			"Expected wrap 'none', got '%s'",
			bp.Wrap(),
		)
	}

	// Test anchor
	bp.SetAnchor(TextAnchorCenter)
	if bp.Anchor() != TextAnchorCenter {
		t.Errorf(
			"Expected anchor 'ctr', got '%s'",
			bp.Anchor(),
		)
	}

	// Test anchor center
	bp.SetAnchorCenter(true)
	if !bp.AnchorCenter() {
		t.Error(
			"Expected anchor center to be true",
		)
	}

	// Test insets
	bp.SetInsets(100, 200, 300, 400)
	if bp.LeftInset() != 100 {
		t.Errorf(
			"Expected left inset 100, got %d",
			bp.LeftInset(),
		)
	}
	if bp.TopInset() != 200 {
		t.Errorf(
			"Expected top inset 200, got %d",
			bp.TopInset(),
		)
	}
	if bp.RightInset() != 300 {
		t.Errorf(
			"Expected right inset 300, got %d",
			bp.RightInset(),
		)
	}
	if bp.BottomInset() != 400 {
		t.Errorf(
			"Expected bottom inset 400, got %d",
			bp.BottomInset(),
		)
	}

	// Test columns
	bp.SetColumnCount(3)
	if bp.ColumnCount() != 3 {
		t.Errorf(
			"Expected 3 columns, got %d",
			bp.ColumnCount(),
		)
	}

	bp.SetColumnSpacing(50000)
	if bp.ColumnSpacing() != 50000 {
		t.Errorf(
			"Expected column spacing 50000, got %d",
			bp.ColumnSpacing(),
		)
	}
}

func TestTextParagraphProperties(t *testing.T) {
	pp := NewTextParagraphProperties()

	// Test alignment
	pp.SetAlignment(TextAlignCenter)
	if pp.Alignment() != TextAlignCenter {
		t.Errorf(
			"Expected alignment 'ctr', got '%s'",
			pp.Alignment(),
		)
	}

	// Test level
	pp.SetLevel(2)
	if pp.Level() != 2 {
		t.Errorf(
			"Expected level 2, got %d",
			pp.Level(),
		)
	}

	// Test margins
	pp.SetLeftMargin(914400) // 1 inch
	if pp.LeftMargin() != 914400 {
		t.Errorf(
			"Expected left margin 914400, got %d",
			pp.LeftMargin(),
		)
	}

	pp.SetRightMargin(457200) // 0.5 inch
	if pp.RightMargin() != 457200 {
		t.Errorf(
			"Expected right margin 457200, got %d",
			pp.RightMargin(),
		)
	}

	// Test indent
	pp.SetIndent(
		-342900,
	) // -0.375 inch (hanging indent)
	if pp.Indent() != -342900 {
		t.Errorf(
			"Expected indent -342900, got %d",
			pp.Indent(),
		)
	}

	// Test RTL
	pp.SetRightToLeft(true)
	if !pp.RightToLeft() {
		t.Error("Expected RTL to be true")
	}

	// Test font alignment
	pp.SetFontAlignment(FontAlignBase)
	if pp.FontAlignment() != FontAlignBase {
		t.Errorf(
			"Expected font alignment 'base', got '%s'",
			pp.FontAlignment(),
		)
	}
}

func TestTextParagraphSpacing(t *testing.T) {
	pp := NewTextParagraphProperties()

	// Test line spacing percent
	pp.SetLineSpacingPercent(150000) // 150%
	ls := pp.LineSpacing()
	if ls == nil {
		t.Fatal("Line spacing should not be nil")
	}
	if !ls.IsPercent() {
		t.Error(
			"Expected line spacing to be percent",
		)
	}
	if ls.Percent() != 150000 {
		t.Errorf(
			"Expected percent 150000, got %d",
			ls.Percent(),
		)
	}

	// Test space before points
	pp.SetSpaceBeforePoints(1200) // 12 points
	sb := pp.SpaceBefore()
	if sb == nil {
		t.Fatal("Space before should not be nil")
	}
	if sb.Points() != 1200 {
		t.Errorf(
			"Expected points 1200, got %d",
			sb.Points(),
		)
	}

	// Test space after
	pp.SetSpaceAfterPercent(50000) // 50%
	sa := pp.SpaceAfter()
	if sa == nil {
		t.Fatal("Space after should not be nil")
	}
	if sa.Percent() != 50000 {
		t.Errorf(
			"Expected percent 50000, got %d",
			sa.Percent(),
		)
	}
}

func TestTextParagraphBullets(t *testing.T) {
	pp := NewTextParagraphProperties()

	// Test no bullet
	pp.SetNoBullet()
	xml := pp.OuterXml()
	if !strings.Contains(xml, "buNone") {
		t.Error("Expected buNone element")
	}

	// Test character bullet
	pp.SetCharacterBullet("-")
	xml = pp.OuterXml()
	if !strings.Contains(xml, "buChar") {
		t.Error("Expected buChar element")
	}
	if !strings.Contains(xml, `char="-"`) {
		t.Error(
			"Expected char attribute with '-'",
		)
	}

	// Test auto-numbered bullet
	pp.SetAutoNumberedBullet(
		AutoNumArabicPeriod,
		1,
	)
	xml = pp.OuterXml()
	if !strings.Contains(xml, "buAutoNum") {
		t.Error("Expected buAutoNum element")
	}
	if !strings.Contains(xml, "arabicPeriod") {
		t.Error("Expected arabicPeriod type")
	}

	// Test bullet font
	pp.SetBulletFont("Wingdings")
	xml = pp.OuterXml()
	if !strings.Contains(xml, "buFont") {
		t.Error("Expected buFont element")
	}
	if !strings.Contains(xml, "Wingdings") {
		t.Error("Expected Wingdings typeface")
	}

	// Test bullet color
	pp.SetBulletColor("FF0000")
	xml = pp.OuterXml()
	if !strings.Contains(xml, "buClr") {
		t.Error("Expected buClr element")
	}
}

func TestTextRun(t *testing.T) {
	r := NewTextRun("Hello")

	if r.Text() != "Hello" {
		t.Errorf(
			"Expected text 'Hello', got '%s'",
			r.Text(),
		)
	}

	r.SetText("World")
	if r.Text() != "World" {
		t.Errorf(
			"Expected text 'World', got '%s'",
			r.Text(),
		)
	}
}

func TestTextRunFormatting(t *testing.T) {
	r := NewTextRun("Formatted")

	// Bold
	r.SetBold(true)
	rp := r.Properties()
	if rp == nil {
		t.Fatal(
			"Run properties should not be nil",
		)
	}
	if !rp.Bold() {
		t.Error("Expected bold to be true")
	}

	// Italic
	r.SetItalic(true)
	if !rp.Italic() {
		t.Error("Expected italic to be true")
	}

	// Underline
	r.SetUnderline(UnderlineSingle)
	if rp.Underline() != UnderlineSingle {
		t.Errorf(
			"Expected underline 'sng', got '%s'",
			rp.Underline(),
		)
	}

	// Font size
	r.SetFontSizePoints(14.0)
	if rp.FontSizePoints() != 14.0 {
		t.Errorf(
			"Expected font size 14.0, got %f",
			rp.FontSizePoints(),
		)
	}

	// Color
	r.SetColor("FF0000")
	sf := rp.SolidFill()
	if sf == nil {
		t.Fatal("Solid fill should not be nil")
	}

	// Latin font
	r.SetLatinFont("Arial")
	lf := rp.LatinFont()
	if lf == nil {
		t.Fatal("Latin font should not be nil")
	}
	if lf.Typeface() != "Arial" {
		t.Errorf(
			"Expected typeface 'Arial', got '%s'",
			lf.Typeface(),
		)
	}
}

func TestTextCharacterProperties(t *testing.T) {
	rp := NewRunProperties()

	// Font size
	rp.SetFontSize(1200) // 12pt
	if rp.FontSize() != 1200 {
		t.Errorf(
			"Expected font size 1200, got %d",
			rp.FontSize(),
		)
	}

	// Strike
	rp.SetStrike(StrikeSingleStrike)
	if rp.Strike() != StrikeSingleStrike {
		t.Errorf(
			"Expected strike 'sngStrike', got '%s'",
			rp.Strike(),
		)
	}

	// Capitalization
	rp.SetCapitalization(CapAll)
	if rp.Capitalization() != CapAll {
		t.Errorf(
			"Expected cap 'all', got '%s'",
			rp.Capitalization(),
		)
	}

	// Kerning
	rp.SetKerning(1200)
	if rp.Kerning() != 1200 {
		t.Errorf(
			"Expected kerning 1200, got %d",
			rp.Kerning(),
		)
	}

	// Spacing
	rp.SetSpacing(300)
	if rp.Spacing() != 300 {
		t.Errorf(
			"Expected spacing 300, got %d",
			rp.Spacing(),
		)
	}

	// Baseline (superscript)
	rp.SetSuperscript()
	if rp.Baseline() != 30000 {
		t.Errorf(
			"Expected baseline 30000, got %d",
			rp.Baseline(),
		)
	}

	// Language
	rp.SetLanguage("en-US")
	if rp.Language() != "en-US" {
		t.Errorf(
			"Expected language 'en-US', got '%s'",
			rp.Language(),
		)
	}
}

func TestTextCharacterPropertiesFonts(
	t *testing.T,
) {
	rp := NewRunProperties()

	// Latin font
	rp.SetLatinFont("Times New Roman")
	lf := rp.LatinFont()
	if lf == nil ||
		lf.Typeface() != "Times New Roman" {
		t.Error(
			"Expected Latin font 'Times New Roman'",
		)
	}

	// East Asian font
	rp.SetEastAsianFont("MS Gothic")
	eaf := rp.EastAsianFont()
	if eaf == nil ||
		eaf.Typeface() != "MS Gothic" {
		t.Error(
			"Expected East Asian font 'MS Gothic'",
		)
	}

	// Complex script font
	rp.SetComplexScriptFont("Arial Unicode MS")
	csf := rp.ComplexScriptFont()
	if csf == nil ||
		csf.Typeface() != "Arial Unicode MS" {
		t.Error(
			"Expected Complex Script font 'Arial Unicode MS'",
		)
	}

	// Symbol font
	rp.SetSymbolFont("Symbol")
	sf := rp.SymbolFont()
	if sf == nil || sf.Typeface() != "Symbol" {
		t.Error("Expected Symbol font 'Symbol'")
	}
}

func TestTextFont(t *testing.T) {
	f := NewLatinFont("Calibri")

	if f.Typeface() != "Calibri" {
		t.Errorf(
			"Expected typeface 'Calibri', got '%s'",
			f.Typeface(),
		)
	}

	// Panose
	f.SetPanose("020F0502020204030204")
	if f.Panose() != "020F0502020204030204" {
		t.Error("Panose not set correctly")
	}

	// Pitch family
	f.SetPitchFamily(34)
	if f.PitchFamily() != 34 {
		t.Errorf(
			"Expected pitch family 34, got %d",
			f.PitchFamily(),
		)
	}

	// Character set
	f.SetCharacterSet(1)
	if f.CharacterSet() != 1 {
		t.Errorf(
			"Expected charset 1, got %d",
			f.CharacterSet(),
		)
	}
}

func TestTextField(t *testing.T) {
	f := NewTextField(
		"{12345678-1234-1234-1234-123456789012}",
		"slidenum",
	)

	if f.ID() != "{12345678-1234-1234-1234-123456789012}" {
		t.Error("Field ID not set correctly")
	}

	if f.FieldType() != "slidenum" {
		t.Errorf(
			"Expected field type 'slidenum', got '%s'",
			f.FieldType(),
		)
	}

	f.SetText("1")
	if f.Text() != "1" {
		t.Errorf(
			"Expected text '1', got '%s'",
			f.Text(),
		)
	}
}

func TestTextLineBreak(t *testing.T) {
	br := NewTextLineBreak()

	if br.LocalName() != "br" {
		t.Errorf(
			"Expected localName 'br', got '%s'",
			br.LocalName(),
		)
	}
}

func TestTabStop(t *testing.T) {
	ts := NewTabStop(
		914400,
		TabAlignCenter,
	) // 1 inch, center

	if ts.Position() != 914400 {
		t.Errorf(
			"Expected position 914400, got %d",
			ts.Position(),
		)
	}

	if ts.Alignment() != TabAlignCenter {
		t.Errorf(
			"Expected alignment 'ctr', got '%s'",
			ts.Alignment(),
		)
	}

	ts.SetAlignment(TabAlignRight)
	if ts.Alignment() != TabAlignRight {
		t.Errorf(
			"Expected alignment 'r', got '%s'",
			ts.Alignment(),
		)
	}
}

func TestTextBodyClone(t *testing.T) {
	tb := NewTextBody()
	p := tb.AddParagraph("Clone Test")
	p.SetAlignment(TextAlignCenter)

	cloned := tb.Clone()
	clone, ok := cloned.(*TextBody)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *TextBody",
			cloned,
		)
	}
	if clone == tb {
		t.Error(
			"Clone should be a different instance",
		)
	}

	paragraphs := clone.Paragraphs()
	if len(paragraphs) != 1 {
		t.Error("Clone should have 1 paragraph")
	}

	if paragraphs[0].GetText() != "Clone Test" {
		t.Error("Clone paragraph text mismatch")
	}
}

func TestTextParagraphClone(t *testing.T) {
	p := NewTextParagraph()
	p.AddRun("Test Run")
	p.SetAlignment(TextAlignRight)

	cloned := p.Clone()
	clone, ok := cloned.(*TextParagraph)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *TextParagraph",
			cloned,
		)
	}
	if clone == p {
		t.Error(
			"Clone should be a different instance",
		)
	}

	if clone.GetText() != "Test Run" {
		t.Error("Clone text mismatch")
	}
}

func TestTextRunClone(t *testing.T) {
	r := NewTextRun("Original")
	r.SetBold(true)
	r.SetFontSizePoints(16.0)

	cloned := r.Clone()
	clone, ok := cloned.(*TextRun)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *TextRun",
			cloned,
		)
	}
	if clone == r {
		t.Error(
			"Clone should be a different instance",
		)
	}

	if clone.Text() != "Original" {
		t.Error("Clone text mismatch")
	}

	rp := clone.Properties()
	if rp == nil {
		t.Fatal("Clone should have properties")
	}
	if !rp.Bold() {
		t.Error("Clone should be bold")
	}
}

func TestTextBodyXmlOutput(t *testing.T) {
	tb := NewTextBody()
	bp := tb.BodyProperties()
	bp.SetAnchor(TextAnchorCenter)
	bp.SetWrap(TextWrapNone)

	p := tb.AddParagraph("Hello World")
	pp := p.EnsureProperties()
	pp.SetAlignment(TextAlignCenter)

	r := p.Runs()[0]
	r.SetBold(true)
	r.SetFontSizePoints(24.0)

	xml := tb.OuterXml()

	// Check key elements are present
	if !strings.Contains(xml, "a:txBody") {
		t.Error("Expected a:txBody element")
	}
	if !strings.Contains(xml, "a:bodyPr") {
		t.Error("Expected a:bodyPr element")
	}
	if !strings.Contains(xml, "a:p") {
		t.Error("Expected a:p element")
	}
	if !strings.Contains(xml, "a:r") {
		t.Error("Expected a:r element")
	}
	if !strings.Contains(xml, "a:rPr") {
		t.Error("Expected a:rPr element")
	}
	if !strings.Contains(xml, "a:t") {
		t.Error("Expected a:t element")
	}
}

func TestParagraphAddRun(t *testing.T) {
	p := NewTextParagraph()

	r1 := p.AddRun("First ")
	r2 := p.AddRun("Second")

	runs := p.Runs()
	if len(runs) != 2 {
		t.Errorf(
			"Expected 2 runs, got %d",
			len(runs),
		)
	}

	if r1.Text() != "First " {
		t.Errorf(
			"Expected 'First ', got '%s'",
			r1.Text(),
		)
	}
	if r2.Text() != "Second" {
		t.Errorf(
			"Expected 'Second', got '%s'",
			r2.Text(),
		)
	}

	fullText := p.GetText()
	if fullText != "First Second" {
		t.Errorf(
			"Expected 'First Second', got '%s'",
			fullText,
		)
	}
}

func TestParagraphAddLineBreak(t *testing.T) {
	p := NewTextParagraph()
	p.AddRun("Line 1")
	p.AddLineBreak()
	p.AddRun("Line 2")

	xml := p.OuterXml()
	if !strings.Contains(xml, "a:br") {
		t.Error(
			"Expected a:br element for line break",
		)
	}
}

func TestEndParagraphRunProperties(t *testing.T) {
	p := NewTextParagraph()

	// Initially should be nil
	if p.EndParagraphRunProperties() != nil {
		t.Error(
			"Expected nil end paragraph run properties initially",
		)
	}

	// Create and set
	rp := p.EnsureEndParagraphRunProperties()
	if rp == nil {
		t.Fatal(
			"EnsureEndParagraphRunProperties returned nil",
		)
	}

	rp.SetFontSizePoints(12.0)

	// Should return same properties
	rp2 := p.EndParagraphRunProperties()
	if rp2 == nil {
		t.Error(
			"EndParagraphRunProperties returned nil after setting",
		)
	}
}

func TestDefaultRunProperties(t *testing.T) {
	pp := NewTextParagraphProperties()

	// Initially should be nil
	if pp.DefaultRunProperties() != nil {
		t.Error(
			"Expected nil default run properties initially",
		)
	}

	// Create and set
	rp := pp.EnsureDefaultRunProperties()
	if rp == nil {
		t.Fatal(
			"EnsureDefaultRunProperties returned nil",
		)
	}

	rp.SetBold(true)

	// Should return same properties
	rp2 := pp.DefaultRunProperties()
	if rp2 == nil {
		t.Error(
			"DefaultRunProperties returned nil after setting",
		)
	}
}

func TestTextFontCreators(t *testing.T) {
	latin := NewLatinFont("Calibri")
	if latin.LocalName() != "latin" ||
		latin.Typeface() != "Calibri" {
		t.Error("NewLatinFont failed")
	}

	ea := NewEastAsianFont("MS Mincho")
	if ea.LocalName() != "ea" ||
		ea.Typeface() != "MS Mincho" {
		t.Error("NewEastAsianFont failed")
	}

	cs := NewComplexScriptFont("Arial")
	if cs.LocalName() != "cs" ||
		cs.Typeface() != "Arial" {
		t.Error("NewComplexScriptFont failed")
	}

	sym := NewSymbolFont("Symbol")
	if sym.LocalName() != "sym" ||
		sym.Typeface() != "Symbol" {
		t.Error("NewSymbolFont failed")
	}

	bu := NewBulletFont("Wingdings")
	if bu.LocalName() != "buFont" ||
		bu.Typeface() != "Wingdings" {
		t.Error("NewBulletFont failed")
	}
}
