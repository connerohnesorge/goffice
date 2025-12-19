package elements

import (
	"strings"
	"testing"
)

func TestNumberingRoot(t *testing.T) {
	numbering := NewNumbering()
	if numbering == nil {
		t.Fatal("NewNumbering() returned nil")
	}
	if numbering.LocalName() != "numbering" {
		t.Errorf("Expected local name 'numbering', got '%s'", numbering.LocalName())
	}
}

func TestAbstractNumCreation(t *testing.T) {
	an := NewAbstractNum()
	if an == nil {
		t.Fatal("NewAbstractNum() returned nil")
	}

	an.SetAbstractNumId(1)
	if an.AbstractNumId() != 1 {
		t.Errorf("Expected AbstractNumId 1, got %d", an.AbstractNumId())
	}

	an.SetMultiLevelType(MultiLevelHybridMultilevel)
	if an.MultiLevelType() != MultiLevelHybridMultilevel {
		t.Errorf("Expected MultiLevelType 'hybridMultilevel', got '%s'", an.MultiLevelType())
	}
}

func TestLevelCreation(t *testing.T) {
	lvl := NewLevel(0)
	if lvl == nil {
		t.Fatal("NewLevel() returned nil")
	}

	if lvl.LevelIndex() != 0 {
		t.Errorf("Expected LevelIndex 0, got %d", lvl.LevelIndex())
	}

	lvl.SetStart(1)
	if lvl.Start() != 1 {
		t.Errorf("Expected Start 1, got %d", lvl.Start())
	}

	lvl.SetNumberFormat(NumberFormatDecimal)
	if lvl.NumberFormat() != NumberFormatDecimal {
		t.Errorf("Expected NumberFormat 'decimal', got '%s'", lvl.NumberFormat())
	}

	lvl.SetLevelText("%1.")
	if lvl.LevelText() != "%1." {
		t.Errorf("Expected LevelText '%%1.', got '%s'", lvl.LevelText())
	}

	lvl.SetLevelJustification(JustificationLeft)
	if lvl.LevelJustification() != JustificationLeft {
		t.Errorf("Expected LevelJustification 'left', got '%s'", lvl.LevelJustification())
	}

	lvl.SetLevelSuffix(LevelSuffixSpace)
	if lvl.LevelSuffix() != LevelSuffixSpace {
		t.Errorf("Expected LevelSuffix 'space', got '%s'", lvl.LevelSuffix())
	}

	lvl.SetLevelRestart(0)
	if lvl.LevelRestart() != 0 {
		t.Errorf("Expected LevelRestart 0, got %d", lvl.LevelRestart())
	}

	lvl.SetLegalNumbering(true)
	if !lvl.IsLegalNumbering() {
		t.Error("Expected IsLegalNumbering() to be true")
	}
}

func TestLevelIndentation(t *testing.T) {
	lvl := NewLevel(0)
	lvl.SetIndentation(720, 360)

	pp := lvl.ParagraphProperties()
	if pp == nil {
		t.Fatal("ParagraphProperties() should not be nil after SetIndentation")
	}

	if pp.Left() != 720 {
		t.Errorf("Expected Left 720, got %d", pp.Left())
	}
	if pp.Hanging() != 360 {
		t.Errorf("Expected Hanging 360, got %d", pp.Hanging())
	}
}

func TestNumberingRunProperties(t *testing.T) {
	lvl := NewLevel(0)
	rp := lvl.GetOrCreateNumberingRunProperties()
	if rp == nil {
		t.Fatal("GetOrCreateNumberingRunProperties() returned nil")
	}

	rp.SetFont("Symbol")
	rp.SetFontSize(24)
	rp.SetBold(true)
	rp.SetColor("FF0000")

	xml := lvl.OuterXml()
	if !strings.Contains(xml, "rFonts") {
		t.Error("Level XML should contain rFonts element")
	}
	if !strings.Contains(xml, "<w:b") {
		t.Error("Level XML should contain bold element")
	}
}

func TestNumberingInstanceCreation(t *testing.T) {
	ni := NewNumberingInstance(0)
	if ni == nil {
		t.Fatal("NewNumberingInstance() returned nil")
	}

	ni.SetNumId(1)
	if ni.NumId() != 1 {
		t.Errorf("Expected NumId 1, got %d", ni.NumId())
	}

	if ni.AbstractNumId() != 0 {
		t.Errorf("Expected AbstractNumId 0, got %d", ni.AbstractNumId())
	}
}

func TestLevelOverride(t *testing.T) {
	ni := NewNumberingInstance(0)
	lo := ni.AddLevelOverride(0)

	if lo == nil {
		t.Fatal("AddLevelOverride() returned nil")
	}

	if lo.LevelIndex() != 0 {
		t.Errorf("Expected LevelIndex 0, got %d", lo.LevelIndex())
	}

	lo.SetStartOverride(5)
	if lo.StartOverride() != 5 {
		t.Errorf("Expected StartOverride 5, got %d", lo.StartOverride())
	}

	// Test level override with complete level
	newLvl := NewLevel(0)
	newLvl.SetNumberFormat(NumberFormatUpperRoman)
	lo.SetLevel(newLvl)

	if lo.Level() == nil {
		t.Error("Level() should not be nil after SetLevel")
	}
}

func TestNumberingAddAbstractNum(t *testing.T) {
	numbering := NewNumbering()

	an1 := NewAbstractNum()
	id1 := numbering.AddAbstractNum(an1)
	if id1 != 0 {
		t.Errorf("Expected first AbstractNum ID 0, got %d", id1)
	}

	an2 := NewAbstractNum()
	id2 := numbering.AddAbstractNum(an2)
	if id2 != 1 {
		t.Errorf("Expected second AbstractNum ID 1, got %d", id2)
	}

	// Verify we can find them
	found := numbering.GetAbstractNum(0)
	if found == nil {
		t.Error("GetAbstractNum(0) should not return nil")
	}

	found = numbering.GetAbstractNum(1)
	if found == nil {
		t.Error("GetAbstractNum(1) should not return nil")
	}
}

func TestNumberingAddNumInstance(t *testing.T) {
	numbering := NewNumbering()

	// Add abstract num first
	an := NewAbstractNum()
	anId := numbering.AddAbstractNum(an)

	// Add numbering instances
	ni1 := NewNumberingInstance(anId)
	id1 := numbering.AddNumInstance(ni1)
	if id1 != 1 {
		t.Errorf("Expected first NumInstance ID 1, got %d", id1)
	}

	ni2 := NewNumberingInstance(anId)
	id2 := numbering.AddNumInstance(ni2)
	if id2 != 2 {
		t.Errorf("Expected second NumInstance ID 2, got %d", id2)
	}

	// Verify we can find them
	found := numbering.GetNumInstance(1)
	if found == nil {
		t.Error("GetNumInstance(1) should not return nil")
	}
}

func TestCreateNumberingInstance(t *testing.T) {
	numbering := NewNumbering()

	an := NewAbstractNum()
	anId := numbering.AddAbstractNum(an)

	ni := numbering.CreateNumberingInstance(anId)
	if ni == nil {
		t.Fatal("CreateNumberingInstance() returned nil")
	}

	if ni.AbstractNumId() != anId {
		t.Errorf("Expected AbstractNumId %d, got %d", anId, ni.AbstractNumId())
	}
}

func TestBulletListFactory(t *testing.T) {
	bulletList := NewStandardBulletList()
	if bulletList == nil {
		t.Fatal("NewStandardBulletList() returned nil")
	}

	// Check that all 9 levels are created
	count := 0
	for range bulletList.Levels() {
		count++
	}
	if count != 9 {
		t.Errorf("Expected 9 levels, got %d", count)
	}

	// Check first level
	lvl0 := bulletList.GetLevel(0)
	if lvl0 == nil {
		t.Fatal("GetLevel(0) should not return nil")
	}
	if lvl0.NumberFormat() != NumberFormatBullet {
		t.Errorf("Expected bullet format, got '%s'", lvl0.NumberFormat())
	}
}

func TestDecimalListFactory(t *testing.T) {
	decimalList := NewDecimalList()
	if decimalList == nil {
		t.Fatal("NewDecimalList() returned nil")
	}

	lvl0 := decimalList.GetLevel(0)
	if lvl0 == nil {
		t.Fatal("GetLevel(0) should not return nil")
	}
	if lvl0.NumberFormat() != NumberFormatDecimal {
		t.Errorf("Expected decimal format, got '%s'", lvl0.NumberFormat())
	}
	if lvl0.LevelText() != "%1." {
		t.Errorf("Expected level text '%%1.', got '%s'", lvl0.LevelText())
	}
}

func TestAlphabeticListFactory(t *testing.T) {
	lowerList := NewAlphabeticList(true)
	if lowerList == nil {
		t.Fatal("NewAlphabeticList(true) returned nil")
	}

	lvl0 := lowerList.GetLevel(0)
	if lvl0.NumberFormat() != NumberFormatLowerLetter {
		t.Errorf("Expected lowercase letter format, got '%s'", lvl0.NumberFormat())
	}

	upperList := NewAlphabeticList(false)
	lvl0 = upperList.GetLevel(0)
	if lvl0.NumberFormat() != NumberFormatUpperLetter {
		t.Errorf("Expected uppercase letter format, got '%s'", lvl0.NumberFormat())
	}
}

func TestRomanNumeralListFactory(t *testing.T) {
	lowerList := NewRomanNumeralList(true)
	if lowerList == nil {
		t.Fatal("NewRomanNumeralList(true) returned nil")
	}

	lvl0 := lowerList.GetLevel(0)
	if lvl0.NumberFormat() != NumberFormatLowerRoman {
		t.Errorf("Expected lowercase roman format, got '%s'", lvl0.NumberFormat())
	}

	upperList := NewRomanNumeralList(false)
	lvl0 = upperList.GetLevel(0)
	if lvl0.NumberFormat() != NumberFormatUpperRoman {
		t.Errorf("Expected uppercase roman format, got '%s'", lvl0.NumberFormat())
	}
}

func TestOutlineListFactory(t *testing.T) {
	outlineList := NewOutlineList()
	if outlineList == nil {
		t.Fatal("NewOutlineList() returned nil")
	}

	if outlineList.MultiLevelType() != MultiLevelMultilevel {
		t.Errorf("Expected multilevel type, got '%s'", outlineList.MultiLevelType())
	}

	// Check level text patterns
	lvl0 := outlineList.GetLevel(0)
	if lvl0.LevelText() != "%1" {
		t.Errorf("Expected level 0 text '%%1', got '%s'", lvl0.LevelText())
	}

	lvl1 := outlineList.GetLevel(1)
	if lvl1.LevelText() != "%1.%2" {
		t.Errorf("Expected level 1 text '%%1.%%2', got '%s'", lvl1.LevelText())
	}

	lvl2 := outlineList.GetLevel(2)
	if lvl2.LevelText() != "%1.%2.%3" {
		t.Errorf("Expected level 2 text '%%1.%%2.%%3', got '%s'", lvl2.LevelText())
	}
}

func TestParagraphNumberingHelpers(t *testing.T) {
	p := NewParagraph("Test paragraph")

	// Initially not numbered
	if p.IsNumbered() {
		t.Error("New paragraph should not be numbered")
	}

	// Apply numbering
	p.ApplyNumbering(1, 0)
	if !p.IsNumbered() {
		t.Error("Paragraph should be numbered after ApplyNumbering")
	}
	if p.NumberingId() != 1 {
		t.Errorf("Expected NumberingId 1, got %d", p.NumberingId())
	}
	if p.NumberingLevel() != 0 {
		t.Errorf("Expected NumberingLevel 0, got %d", p.NumberingLevel())
	}

	// Change level
	p.SetNumberingLevel(2)
	if p.NumberingLevel() != 2 {
		t.Errorf("Expected NumberingLevel 2, got %d", p.NumberingLevel())
	}

	// Remove numbering
	p.RemoveNumbering()
	if p.IsNumbered() {
		t.Error("Paragraph should not be numbered after RemoveNumbering")
	}
}

func TestNumberingXMLOutput(t *testing.T) {
	numbering := NewNumbering()

	// Create a simple decimal list
	an := NewDecimalList()
	anId := numbering.AddAbstractNum(an)

	ni := numbering.CreateNumberingInstance(anId)

	xml := numbering.OuterXml()

	// Verify structure
	if !strings.Contains(xml, "<w:numbering") {
		t.Error("XML should contain numbering root element")
	}
	if !strings.Contains(xml, "<w:abstractNum") {
		t.Error("XML should contain abstractNum element")
	}
	if !strings.Contains(xml, "<w:num") {
		t.Error("XML should contain num element")
	}
	if !strings.Contains(xml, "<w:lvl") {
		t.Error("XML should contain lvl element")
	}
	if !strings.Contains(xml, "numFmt") {
		t.Error("XML should contain numFmt element")
	}

	_ = ni // silence unused variable warning
}

func TestAbstractNumIterators(t *testing.T) {
	numbering := NewNumbering()

	an1 := NewAbstractNum()
	an2 := NewAbstractNum()
	numbering.AddAbstractNum(an1)
	numbering.AddAbstractNum(an2)

	count := 0
	for range numbering.AbstractNums() {
		count++
	}
	if count != 2 {
		t.Errorf("Expected 2 abstract nums, got %d", count)
	}

	ni1 := NewNumberingInstance(0)
	numbering.AddNumInstance(ni1)

	count = 0
	for range numbering.NumInstances() {
		count++
	}
	if count != 1 {
		t.Errorf("Expected 1 num instance, got %d", count)
	}
}
