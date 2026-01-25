package presentation

import (
	"path/filepath"
	"testing"
)

// TestAddNotes tests adding speaker notes to a slide.
func TestAddNotes(t *testing.T) {
	pres, err := New(filepath.Join(t.TempDir(), "notes.pptx"), DocTypePresentation)
	if err != nil {
		t.Fatalf("Failed to create presentation: %v", err)
	}
	defer pres.Close()

	slide, _ := pres.AddSlide()

	notesPart, err := slide.AddNotesSlidePart()
	if err != nil {
		t.Fatalf("Failed to add notes slide part: %v", err)
	}

	notesText := "These are speaker notes."
	notesPart.SetNotes(notesText)

	// Verify it saves correctly
	if err := pres.Save(); err != nil {
		t.Fatalf("Failed to save: %v", err)
	}
}

// TestAddRichTextNotes tests adding rich text notes to a slide.
func TestAddRichTextNotes(t *testing.T) {
	pres, _ := New(filepath.Join(t.TempDir(), "rich_notes.pptx"), DocTypePresentation)
	slide, _ := pres.AddSlide()
	notesPart, _ := slide.AddNotesSlidePart()

	tb := notesPart.GetOrCreateTextBody()
	p := tb.AddParagraph("")
	run1 := p.AddRun("Bold red text")
	run1.SetBold(true)
	run1.SetColor("FF0000")

	run2 := p.AddRun(" and normal text.")
	_ = run2

	if err := pres.Save(); err != nil {
		t.Fatalf("Failed to save: %v", err)
	}
}
