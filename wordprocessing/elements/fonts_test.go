package elements

import (
	"testing"
)


func TestFontsAddAndGetFont(t *testing.T) {
	f := NewFonts()

	// Initially no fonts
	count := 0
	for range f.Fonts() {
		count++
	}
	if count != 0 {
		t.Errorf(
			"Expected 0 fonts initially, got %d",
			count,
		)
	}

	// Add font by name
	font := f.AddFontByName(testFontArial)
	if font == nil {
		t.Fatal("Expected font to be created")
	}

	if font.Name() != testFontArial {
		t.Errorf(
			"Expected font name %q, got %q",
			testFontArial,
			font.Name(),
		)
	}

	// Get font by name
	found := f.GetFont(testFontArial)
	if found == nil {
		t.Fatalf(
			"Expected to find font %q",
			testFontArial,
		)
	}
	if found.Name() != testFontArial {
		t.Errorf(
			"Expected font name %q, got %q",
			testFontArial,
			found.Name(),
		)
	}

	// Font not found
	notFound := f.GetFont("NonExistent")
	if notFound != nil {
		t.Error(
			"Expected nil for non-existent font",
		)
	}
}

func TestFontProperties(t *testing.T) {
	font := NewFont()
	font.SetName("Calibri")

	if font.Name() != "Calibri" {
		t.Errorf(
			"Expected name 'Calibri', got %q",
			font.Name(),
		)
	}

	// Test charset
	font.SetCharset("00")
	if font.Charset() != "00" {
		t.Errorf(
			"Expected charset '00', got %q",
			font.Charset(),
		)
	}

	// Test family
	if font.Family() != FontFamilyAuto {
		t.Errorf(
			"Expected default family %q, got %q",
			FontFamilyAuto,
			font.Family(),
		)
	}
	font.SetFamily(FontFamilySwiss)
	if font.Family() != FontFamilySwiss {
		t.Errorf(
			"Expected family %q, got %q",
			FontFamilySwiss,
			font.Family(),
		)
	}

	// Test pitch
	if font.Pitch() != FontPitchDefault {
		t.Errorf(
			"Expected default pitch %q, got %q",
			FontPitchDefault,
			font.Pitch(),
		)
	}
	font.SetPitch(FontPitchVariable)
	if font.Pitch() != FontPitchVariable {
		t.Errorf(
			"Expected pitch %q, got %q",
			FontPitchVariable,
			font.Pitch(),
		)
	}

	// Test Panose-1
	font.SetPanose1("020B0604020202020204")
	if font.Panose1() != "020B0604020202020204" {
		t.Errorf(
			"Expected Panose-1 '020B0604020202020204', got %q",
			font.Panose1(),
		)
	}

	// Test alternate name
	font.SetAltName("Arial")
	if font.AltName() != "Arial" {
		t.Errorf(
			"Expected alt name 'Arial', got %q",
			font.AltName(),
		)
	}
}

func TestFontEmbedded(t *testing.T) {
	font := NewFont()
	font.SetName("TestFont")

	// Initially nil
	if font.EmbedRegular() != nil {
		t.Error(
			"Expected EmbedRegular to be nil initially",
		)
	}

	// Get or create
	ef := font.GetOrCreateEmbedRegular()
	if ef == nil {
		t.Fatal(
			"Expected EmbedRegular to be created",
		)
	}

	ef.SetRelationshipId("rId1")
	if ef.RelationshipId() != "rId1" {
		t.Errorf(
			"Expected relationship ID 'rId1', got %q",
			ef.RelationshipId(),
		)
	}

	ef.SetFontKey("{GUID-1234}")
	if ef.FontKey() != "{GUID-1234}" {
		t.Errorf(
			"Expected font key '{GUID-1234}', got %q",
			ef.FontKey(),
		)
	}

	// Test other embed styles
	bold := font.GetOrCreateEmbedBold()
	if bold == nil {
		t.Fatal(
			"Expected EmbedBold to be created",
		)
	}

	italic := font.GetOrCreateEmbedItalic()
	if italic == nil {
		t.Fatal(
			"Expected EmbedItalic to be created",
		)
	}

	boldItalic := font.GetOrCreateEmbedBoldItalic()
	if boldItalic == nil {
		t.Fatal(
			"Expected EmbedBoldItalic to be created",
		)
	}
}

func TestFontSig(t *testing.T) {
	font := NewFont()
	font.SetName("TestFont")

	sig := font.GetOrCreateSig()
	if sig == nil {
		t.Fatal("Expected FontSig to be created")
	}

	sig.SetUsb0("E0002AFF")
	if sig.Usb0() != "E0002AFF" {
		t.Errorf(
			"Expected usb0 'E0002AFF', got %q",
			sig.Usb0(),
		)
	}

	sig.SetUsb1("C0007843")
	if sig.Usb1() != "C0007843" {
		t.Errorf(
			"Expected usb1 'C0007843', got %q",
			sig.Usb1(),
		)
	}

	sig.SetCsb0("000001FF")
	if sig.Csb0() != "000001FF" {
		t.Errorf(
			"Expected csb0 '000001FF', got %q",
			sig.Csb0(),
		)
	}
}

func TestFontsFontsIterator(t *testing.T) {
	f := NewFonts()

	// Add multiple fonts
	f.AddFontByName("Arial")
	f.AddFontByName("Times New Roman")
	f.AddFontByName("Calibri")

	// Count fonts
	count := 0
	names := make([]string, 0)
	for font := range f.Fonts() {
		count++
		names = append(names, font.Name())
	}

	if count != 3 {
		t.Errorf(
			"Expected 3 fonts, got %d",
			count,
		)
	}

	// Verify all names are present
	expectedNames := []string{
		"Arial",
		"Times New Roman",
		"Calibri",
	}
	for _, expected := range expectedNames {
		found := false
		for _, name := range names {
			if name == expected {
				found = true

				break
			}
		}
		if !found {
			t.Errorf(
				"Expected to find font %q",
				expected,
			)
		}
	}
}

func TestFontsRemoveFont(t *testing.T) {
	f := NewFonts()

	font := f.AddFontByName("Arial")

	// Verify font exists
	if f.GetFont("Arial") == nil {
		t.Fatal("Expected font to exist")
	}

	// Remove font
	if !f.RemoveFont(font) {
		t.Error(
			"Expected RemoveFont to return true",
		)
	}

	// Verify font is removed
	if f.GetFont("Arial") != nil {
		t.Error("Expected font to be removed")
	}
}

func TestFontsClone(t *testing.T) {
	f := NewFonts()
	font := f.AddFontByName("Arial")
	font.SetFamily(FontFamilySwiss)

	cloneElem := f.Clone()
	clone, ok := cloneElem.(*Fonts)
	if !ok {
		t.Fatal("Expected Clone to return *Fonts")
	}

	// Verify clone has the font
	clonedFont := clone.GetFont("Arial")
	if clonedFont == nil {
		t.Fatal("Expected cloned font to exist")
	}

	if clonedFont.Family() != FontFamilySwiss {
		t.Errorf(
			"Expected cloned font family %q, got %q",
			FontFamilySwiss,
			clonedFont.Family(),
		)
	}

	// Modify original
	font.SetFamily(FontFamilyRoman)
	if clonedFont.Family() != FontFamilySwiss {
		t.Error(
			"Clone should be independent of original",
		)
	}
}

func TestFontClone(t *testing.T) {
	font := NewFont()
	font.SetName("TestFont")
	font.SetFamily(FontFamilyModern)
	font.SetPitch(FontPitchFixed)

	cloneElem := font.Clone()
	clone, ok := cloneElem.(*Font)
	if !ok {
		t.Fatal("Expected Clone to return *Font")
	}

	if clone.Name() != "TestFont" {
		t.Errorf(
			"Expected cloned name 'TestFont', got %q",
			clone.Name(),
		)
	}
	if clone.Family() != FontFamilyModern {
		t.Errorf(
			"Expected cloned family %q, got %q",
			FontFamilyModern,
			clone.Family(),
		)
	}
	if clone.Pitch() != FontPitchFixed {
		t.Errorf(
			"Expected cloned pitch %q, got %q",
			FontPitchFixed,
			clone.Pitch(),
		)
	}
}
