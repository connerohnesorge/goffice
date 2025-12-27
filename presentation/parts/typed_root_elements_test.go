package parts

import (
	"testing"

	_ "github.com/connerohnesorge/goffice/presentation/elements"
)

// TestHandoutMasterPartTypedElement tests that HandoutMasterPart.HandoutMaster() returns proper typed element.
func TestHandoutMasterPartTypedElement(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	hmp, err := presPart.AddHandoutMasterPart()
	if err != nil {
		t.Fatalf(
			"Failed to create handout master part: %v",
			err,
		)
	}

	hm := hmp.HandoutMaster()
	if hm == nil {
		t.Fatal("HandoutMaster() returned nil")
	}

	// Type assertion - no need to check ok since we're asserting the type
	_ = hm
}

// TestNotesSlidePartTypedElement tests that NotesSlidePart.NotesSlide() returns proper typed element.
func TestNotesSlidePartTypedElement(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	slide, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf(
			"Failed to create slide part: %v",
			err,
		)
	}

	nsp, err := slide.AddNotesSlidePart()
	if err != nil {
		t.Fatalf(
			"Failed to create notes slide part: %v",
			err,
		)
	}

	ns := nsp.NotesSlide()
	if ns == nil {
		t.Fatal("NotesSlide() returned nil")
	}

	_ = ns
}

// TestSlideMasterPartTypedElement tests that SlideMasterPart.SlideMaster() returns proper typed element.
func TestSlideMasterPartTypedElement(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	smp, err := presPart.AddSlideMasterPart()
	if err != nil {
		t.Fatalf(
			"Failed to create slide master part: %v",
			err,
		)
	}

	sm := smp.SlideMaster()
	if sm == nil {
		t.Fatal("SlideMaster() returned nil")
	}

	_ = sm
}

// TestCommentAuthorsPartTypedElement tests that CommentAuthorsPart.CommentAuthors() returns proper typed element.
func TestCommentAuthorsPartTypedElement(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	commentAuthorsPart, err := presPart.AddCommentAuthorsPart()
	if err != nil {
		t.Fatalf(
			"Failed to create comment authors part: %v",
			err,
		)
	}

	ca := commentAuthorsPart.CommentAuthors()
	if ca == nil {
		t.Fatal("CommentAuthors() returned nil")
	}

	_ = ca
}

// TestSlideCommentsPartTypedElement tests that SlideCommentsPart.Comments() returns proper typed element.
func TestSlideCommentsPartTypedElement(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	slide, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf(
			"Failed to create slide part: %v",
			err,
		)
	}

	scp, err := slide.AddSlideCommentsPart()
	if err != nil {
		t.Fatalf(
			"Failed to create slide comments part: %v",
			err,
		)
	}

	cl := scp.Comments()
	if cl == nil {
		t.Fatal("Comments() returned nil")
	}

	_ = cl
}

// TestSlideLayoutPartTypedElement tests that SlideLayoutPart.SlideLayout() returns proper typed element.
func TestSlideLayoutPartTypedElement(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	smp, err := presPart.AddSlideMasterPart()
	if err != nil {
		t.Fatalf(
			"Failed to create slide master part: %v",
			err,
		)
	}

	slp, err := smp.AddSlideLayoutPart()
	if err != nil {
		t.Fatalf(
			"Failed to create slide layout part: %v",
			err,
		)
	}

	sl := slp.SlideLayout()
	if sl == nil {
		t.Fatal("SlideLayout() returned nil")
	}

	_ = sl
}

// TestNotesMasterPartTypedElement tests that NotesMasterPart.NotesMaster() returns proper typed element.
func TestNotesMasterPartTypedElement(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	nmp, err := presPart.AddNotesMasterPart()
	if err != nil {
		t.Fatalf(
			"Failed to create notes master part: %v",
			err,
		)
	}

	nm := nmp.NotesMaster()
	if nm == nil {
		t.Fatal("NotesMaster() returned nil")
	}

	_ = nm
}

// TestSlidePartTypedElement tests that SlidePart.Slide() returns proper typed element (already implemented).
func TestSlidePartTypedElement(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	slide, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf(
			"Failed to create slide part: %v",
			err,
		)
	}

	s := slide.Slide()
	if s == nil {
		t.Fatal("Slide() returned nil")
	}

	_ = s
}
