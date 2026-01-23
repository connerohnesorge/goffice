package elements

import (
	"testing"
)

func TestDocDefaults(t *testing.T) {
	styles := NewStyles()

	// Initially nil
	if styles.DocDefaults() != nil {
		t.Error(
			"Expected DocDefaults() to be nil initially",
		)
	}

	// GetOrCreate should create it
	dd := styles.GetOrCreateDocDefaults()
	if dd == nil {
		t.Fatal(
			"GetOrCreateDocDefaults() returned nil",
		)
	}

	// Should be first child
	if styles.DocDefaults() == nil {
		t.Error(
			"DocDefaults() should not be nil after creation",
		)
	}
}

func TestLatentStyles(t *testing.T) {
	styles := NewStyles()

	// Initially nil
	if styles.LatentStyles() != nil {
		t.Error(
			"Expected LatentStyles() to be nil initially",
		)
	}

	// GetOrCreate should create it
	ls := styles.GetOrCreateLatentStyles()
	if ls == nil {
		t.Fatal(
			"GetOrCreateLatentStyles() returned nil",
		)
	}

	// Set some defaults
	ls.SetDefSemiHidden(true)
	ls.SetDefUnhideWhenUsed(true)
	ls.SetDefUIPriority(99)

	if !ls.DefSemiHidden() {
		t.Error("DefSemiHidden() should be true")
	}
	if !ls.DefUnhideWhenUsed() {
		t.Error(
			"DefUnhideWhenUsed() should be true",
		)
	}
	if ls.DefUIPriority() != 99 {
		t.Errorf(
			"Expected DefUIPriority() 99, got %d",
			ls.DefUIPriority(),
		)
	}
}

func TestStylesCollection(t *testing.T) {
	styles := NewStyles()

	style1 := NewParagraphStyle(
		"Style1",
		"Style One",
	)
	style2 := NewParagraphStyle(
		"Style2",
		"Style Two",
	)
	style3 := NewCharacterStyle(
		"Style3",
		"Style Three",
	)

	styles.AddStyle(style1)
	styles.AddStyle(style2)
	styles.AddStyle(style3)

	// Test GetStyleById
	found := styles.GetStyleById("Style1")
	if found == nil {
		t.Error(
			"GetStyleById('Style1') should not return nil",
		)
	}
	if found.StyleName() != "Style One" {
		t.Errorf(
			"Expected 'Style One', got '%s'",
			found.StyleName(),
		)
	}

	// Test GetStyleByName
	found = styles.GetStyleByName("Style Two")
	if found == nil {
		t.Error(
			"GetStyleByName('Style Two') should not return nil",
		)
	}
	if found.StyleId() != "Style2" {
		t.Errorf(
			"Expected 'Style2', got '%s'",
			found.StyleId(),
		)
	}

	// Test GetStylesByType
	count := 0
	for range styles.GetStylesByType(StyleTypeParagraph) {
		count++
	}
	if count != 2 {
		t.Errorf(
			"Expected 2 paragraph styles, got %d",
			count,
		)
	}

	// Test iteration
	count = 0
	for range styles.Styles() {
		count++
	}
	if count != 3 {
		t.Errorf(
			"Expected 3 styles, got %d",
			count,
		)
	}
}

func TestLatentStyleException(t *testing.T) {
	ls := NewLatentStyles()
	ex := ls.AddException("Heading 1")

	ex.SetQFormat(true)
	ex.SetUIPriority(9)
	ex.SetSemiHidden(false)

	if ex.Name() != "Heading 1" {
		t.Errorf(
			"Expected Name 'Heading 1', got '%s'",
			ex.Name(),
		)
	}
	if !ex.QFormat() {
		t.Error("Expected QFormat() to be true")
	}
	if ex.UIPriority() != 9 {
		t.Errorf(
			"Expected UIPriority 9, got %d",
			ex.UIPriority(),
		)
	}

	// Find it
	found := ls.GetException("Heading 1")
	if found == nil {
		t.Error(
			"GetException('Heading 1') should not return nil",
		)
	}
}