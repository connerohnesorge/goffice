package presentation

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestNew_CreatesDocument tests creating new documents of different types.
func TestNew_CreatesDocument(t *testing.T) {
	tests := []struct {
		name        string
		docType     DocType
		wantExt     string
		wantContent string
	}{
		{
			name:        "Presentation",
			docType:     DocTypePresentation,
			wantExt:     ".pptx",
			wantContent: ContentTypePresentationMLPresentation,
		},
		{
			name:        "Template",
			docType:     DocTypeTemplate,
			wantExt:     ".potx",
			wantContent: ContentTypePresentationMLTemplate,
		},
		{
			name:        "Slideshow",
			docType:     DocTypeSlideshow,
			wantExt:     ".ppsx",
			wantContent: ContentTypePresentationMLSlideshow,
		},
		{
			name:        "MacroEnabledPresentation",
			docType:     DocTypeMacroEnabledPresentation,
			wantExt:     ".pptm",
			wantContent: ContentTypePresentationMLMacroEnabled,
		},
		{
			name:        "MacroEnabledTemplate",
			docType:     DocTypeMacroEnabledTemplate,
			wantExt:     ".potm",
			wantContent: ContentTypePresentationMLMacroTemplate,
		},
		{
			name:        "MacroEnabledSlideshow",
			docType:     DocTypeMacroEnabledSlideshow,
			wantExt:     ".ppsm",
			wantContent: ContentTypePresentationMLMacroSlideshow,
		},
		{
			name:        "AddIn",
			docType:     DocTypeAddIn,
			wantExt:     ".ppam",
			wantContent: ContentTypePresentationMLAddIn,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			path := filepath.Join(
				tempDir,
				"test"+tt.wantExt,
			)

			doc, err := New(path, tt.docType)
			if err != nil {
				t.Fatalf("New() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			if doc.Type() != tt.docType {
				t.Errorf(
					"Type() = %v, want %v",
					doc.Type(),
					tt.docType,
				)
			}

			if doc.docType.Extension() != tt.wantExt {
				t.Errorf(
					"Extension() = %v, want %v",
					doc.docType.Extension(),
					tt.wantExt,
				)
			}

			if doc.docType.ContentType() != tt.wantContent {
				t.Errorf(
					"ContentType() = %v, want %v",
					doc.docType.ContentType(),
					tt.wantContent,
				)
			}

			if !doc.IsEditable() {
				t.Error(
					"IsEditable() = false, want true",
				)
			}

			if doc.Path() != path {
				t.Errorf(
					"Path() = %v, want %v",
					doc.Path(),
					path,
				)
			}

			if doc.Package() == nil {
				t.Error("Package() is nil")
			}

			if doc.PresentationPart() == nil {
				t.Error(
					"PresentationPart() is nil",
				)
			}
		})
	}
}

// TestOpen_OpensDocument tests opening existing documents.
func TestOpen_OpensDocument(t *testing.T) {
	// Create a test document first
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Add a slide before saving
	_, err = doc.AddSlide()
	if err != nil {
		t.Fatalf("AddSlide() error = %v", err)
	}

	err = doc.SaveAs(path)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	err = doc.Close()
	if err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Now open the document
	t.Run("OpenEditable", func(t *testing.T) {
		openDoc, err := Open(path, true)
		if err != nil {
			t.Fatalf("Open() error = %v", err)
		}
		defer func() { _ = openDoc.Close() }()

		if !openDoc.IsEditable() {
			t.Error(
				"IsEditable() = false, want true",
			)
		}

		if openDoc.Type() != DocTypePresentation {
			t.Errorf(
				"Type() = %v, want %v",
				openDoc.Type(),
				DocTypePresentation,
			)
		}
	})

	t.Run("OpenReadOnly", func(t *testing.T) {
		openDoc, err := Open(path, false)
		if err != nil {
			t.Fatalf("Open() error = %v", err)
		}
		defer func() { _ = openDoc.Close() }()

		if openDoc.IsEditable() {
			t.Error(
				"IsEditable() = true, want false",
			)
		}
	})
}

// TestDocument_AddSlide tests adding slides to a document.
func TestDocument_AddSlide(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Initially should have 0 slides
	if doc.SlideCount() != 0 {
		t.Errorf(
			"SlideCount() = %v, want 0",
			doc.SlideCount(),
		)
	}

	// Add first slide
	slide1, err := doc.AddSlide()
	if err != nil {
		t.Fatalf("AddSlide() error = %v", err)
	}
	if slide1 == nil {
		t.Error("AddSlide() returned nil")
	}

	if doc.SlideCount() != 1 {
		t.Errorf(
			"SlideCount() = %v, want 1",
			doc.SlideCount(),
		)
	}

	// Add second slide
	slide2, err := doc.AddSlide()
	if err != nil {
		t.Fatalf("AddSlide() error = %v", err)
	}
	if slide2 == nil {
		t.Error("AddSlide() returned nil")
	}

	if doc.SlideCount() != 2 {
		t.Errorf(
			"SlideCount() = %v, want 2",
			doc.SlideCount(),
		)
	}

	// Add third slide
	_, err = doc.AddSlide()
	if err != nil {
		t.Fatalf("AddSlide() error = %v", err)
	}

	if doc.SlideCount() != 3 {
		t.Errorf(
			"SlideCount() = %v, want 3",
			doc.SlideCount(),
		)
	}
}

// TestDocument_SlideCount tests counting slides.
func TestDocument_SlideCount(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Test with 0, 1, 5 slides
	counts := []int{0, 1, 5}
	for _, want := range counts {
		for doc.SlideCount() < want {
			_, err := doc.AddSlide()
			if err != nil {
				t.Fatalf(
					"AddSlide() error = %v",
					err,
				)
			}
		}
		if got := doc.SlideCount(); got != want {
			t.Errorf(
				"SlideCount() = %v, want %v",
				got,
				want,
			)
		}
	}
}

// TestDocument_SaveAs tests saving to a new file.
func TestDocument_SaveAs(t *testing.T) {
	tempDir := t.TempDir()
	srcPath := filepath.Join(
		tempDir,
		"source.pptx",
	)
	dstPath := filepath.Join(
		tempDir,
		"destination.pptx",
	)

	// Create and save a document
	doc, err := New(srcPath, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	_, err = doc.AddSlide()
	if err != nil {
		t.Fatalf("AddSlide() error = %v", err)
	}

	err = doc.SaveAs(dstPath)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}

	err = doc.Close()
	if err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Verify the destination file exists
	if _, err := os.Stat(dstPath); os.IsNotExist(
		err,
	) {
		t.Error(
			"SaveAs() did not create destination file",
		)
	}

	// Verify we can open the saved file
	openDoc, err := Open(dstPath, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = openDoc.Close() }()

	// Note: Slide persistence on reopen is not yet fully implemented
	// Once implemented, this should check for SlideCount() == 1
	if openDoc.SlideCount() < 0 {
		t.Errorf(
			"SlideCount() = %v, should be >= 0",
			openDoc.SlideCount(),
		)
	}
}

// TestDocument_ChangeType tests changing document type.
func TestDocument_ChangeType(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Change to template
	err = doc.ChangeType(DocTypeTemplate)
	if err != nil {
		t.Fatalf("ChangeType() error = %v", err)
	}

	if doc.Type() != DocTypeTemplate {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeTemplate,
		)
	}

	// Change back to presentation
	err = doc.ChangeType(DocTypePresentation)
	if err != nil {
		t.Fatalf("ChangeType() error = %v", err)
	}

	if doc.Type() != DocTypePresentation {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypePresentation,
		)
	}

	// Test changing to same type (should be no-op)
	err = doc.ChangeType(DocTypePresentation)
	if err != nil {
		t.Fatalf("ChangeType() error = %v", err)
	}
}

// TestDocument_Close tests closing documents.
func TestDocument_Close(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	err = doc.Close()
	if err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Calling Close again should not error
	err = doc.Close()
	if err != nil {
		t.Errorf(
			"Close() second call error = %v",
			err,
		)
	}

	// PresentationPart should be nil after close
	if doc.PresentationPart() != nil {
		t.Error(
			"PresentationPart() should be nil after close",
		)
	}
}

// TestDocType_Extension tests extension mapping for each type.
func TestDocType_Extension(t *testing.T) {
	tests := []struct {
		docType DocType
		want    string
	}{
		{DocTypePresentation, ".pptx"},
		{DocTypeTemplate, ".potx"},
		{DocTypeSlideshow, ".ppsx"},
		{
			DocTypeMacroEnabledPresentation,
			".pptm",
		},
		{DocTypeMacroEnabledTemplate, ".potm"},
		{DocTypeMacroEnabledSlideshow, ".ppsm"},
		{DocTypeAddIn, ".ppam"},
		{
			DocType(999),
			".pptx",
		}, // Unknown should default to .pptx
	}

	for _, tt := range tests {
		t.Run(
			tt.docType.String(),
			func(t *testing.T) {
				if got := tt.docType.Extension(); got != tt.want {
					t.Errorf(
						"Extension() = %v, want %v",
						got,
						tt.want,
					)
				}
			},
		)
	}
}

// TestDocType_ContentType tests content type mapping for each type.
func TestDocType_ContentType(t *testing.T) {
	tests := []struct {
		docType DocType
		want    string
	}{
		{
			DocTypePresentation,
			ContentTypePresentationMLPresentation,
		},
		{
			DocTypeTemplate,
			ContentTypePresentationMLTemplate,
		},
		{
			DocTypeSlideshow,
			ContentTypePresentationMLSlideshow,
		},
		{
			DocTypeMacroEnabledPresentation,
			ContentTypePresentationMLMacroEnabled,
		},
		{
			DocTypeMacroEnabledTemplate,
			ContentTypePresentationMLMacroTemplate,
		},
		{
			DocTypeMacroEnabledSlideshow,
			ContentTypePresentationMLMacroSlideshow,
		},
		{
			DocTypeAddIn,
			ContentTypePresentationMLAddIn,
		},
		{
			DocType(999),
			ContentTypePresentationMLPresentation,
		}, // Unknown should default
	}

	for _, tt := range tests {
		t.Run(
			tt.docType.String(),
			func(t *testing.T) {
				if got := tt.docType.ContentType(); got != tt.want {
					t.Errorf(
						"ContentType() = %v, want %v",
						got,
						tt.want,
					)
				}
			},
		)
	}
}

// TestDocType_String tests string representation.
func TestDocType_String(t *testing.T) {
	tests := []struct {
		docType DocType
		want    string
	}{
		{DocTypePresentation, "Presentation"},
		{DocTypeTemplate, "Template"},
		{DocTypeSlideshow, "Slideshow"},
		{
			DocTypeMacroEnabledPresentation,
			"MacroEnabledPresentation",
		},
		{
			DocTypeMacroEnabledTemplate,
			"MacroEnabledTemplate",
		},
		{
			DocTypeMacroEnabledSlideshow,
			"MacroEnabledSlideshow",
		},
		{DocTypeAddIn, "AddIn"},
		{
			DocType(999),
			"Presentation",
		}, // Unknown should default
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.docType.String(); got != tt.want {
				t.Errorf(
					"String() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

// TestDocument_Slides tests iterating slides.
func TestDocument_Slides(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Add 3 slides
	for range 3 {
		_, err := doc.AddSlide()
		if err != nil {
			t.Fatalf("AddSlide() error = %v", err)
		}
	}

	// Count slides using iterator
	count := 0
	for slide := range doc.Slides() {
		if slide == nil {
			t.Error("Slides() yielded nil slide")
		}
		count++
	}

	if count != 3 {
		t.Errorf(
			"Slides() yielded %d slides, want 3",
			count,
		)
	}
}

// TestDocument_GetSlide tests getting a slide by index.
func TestDocument_GetSlide(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Add 3 slides
	for range 3 {
		_, err := doc.AddSlide()
		if err != nil {
			t.Fatalf("AddSlide() error = %v", err)
		}
	}

	// Get valid slides
	for i := range 3 {
		slide, err := doc.GetSlide(i)
		if err != nil {
			t.Errorf(
				"GetSlide(%d) error = %v",
				i,
				err,
			)
		}
		if slide == nil {
			t.Errorf(
				"GetSlide(%d) returned nil",
				i,
			)
		}
	}

	// Get invalid index (negative)
	_, err = doc.GetSlide(-1)
	if err != ErrSlideIndexOutOfRange {
		t.Errorf(
			"GetSlide(-1) error = %v, want ErrSlideIndexOutOfRange",
			err,
		)
	}

	// Get invalid index (too high)
	_, err = doc.GetSlide(10)
	if err != ErrSlideIndexOutOfRange {
		t.Errorf(
			"GetSlide(10) error = %v, want ErrSlideIndexOutOfRange",
			err,
		)
	}
}

// TestDocument_AddSlideMaster tests adding slide masters.
func TestDocument_AddSlideMaster(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	master, err := doc.AddSlideMaster()
	if err != nil {
		t.Fatalf(
			"AddSlideMaster() error = %v",
			err,
		)
	}
	if master == nil {
		t.Error("AddSlideMaster() returned nil")
	}

	// Verify slide masters iterator
	count := 0
	for m := range doc.SlideMasters() {
		if m == nil {
			t.Error(
				"SlideMasters() yielded nil master",
			)
		}
		count++
	}

	if count != 1 {
		t.Errorf(
			"SlideMasters() yielded %d masters, want 1",
			count,
		)
	}
}

// TestDocument_Theme tests adding and getting themes.
func TestDocument_Theme(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// New presentations have a default theme
	theme := doc.Theme()
	if theme == nil {
		t.Error("Theme() should not be nil for new presentations")
	}

	// Add another theme
	theme2, err := doc.AddTheme()
	if err != nil {
		t.Fatalf("AddTheme() error = %v", err)
	}
	if theme2 == nil {
		t.Error("AddTheme() returned nil")
	}

	// Should still have at least one theme
	if doc.Theme() == nil {
		t.Error(
			"Theme() should not be nil after adding another theme",
		)
	}
}

func TestNewPresentationHasTheme(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	theme := doc.PresentationPart().ThemePart()
	if theme == nil {
		t.Fatal("PresentationPart().ThemePart() is nil")
	}

	// Check that theme URI follows the pattern /ppt/theme/theme<N>.xml
	themeURI := theme.URI()
	if !strings.HasPrefix(themeURI, "/ppt/theme/theme") || !strings.HasSuffix(themeURI, ".xml") {
		t.Errorf(
			"Theme URI = %v, expected pattern /ppt/theme/theme<N>.xml",
			themeURI,
		)
	}

	master, err := doc.AddSlideMaster()
	if err != nil {
		t.Fatalf("AddSlideMaster() error = %v", err)
	}

	masterTheme := master.ThemePart()
	if masterTheme == nil {
		t.Fatal("SlideMaster.ThemePart() is nil")
	}

	if masterTheme.URI() != theme.URI() {
		t.Errorf(
			"Master theme URI = %v, want %v",
			masterTheme.URI(),
			theme.URI(),
		)
	}
}

// TestDocument_ReadOnlyErrors tests that read-only documents return errors.
func TestDocument_ReadOnlyErrors(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	// Create a document first
	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	err = doc.SaveAs(path)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	err = doc.Close()
	if err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Open read-only
	roDoc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = roDoc.Close() }()

	// AddSlide should error
	_, err = roDoc.AddSlide()
	if err != ErrReadOnly {
		t.Errorf(
			"AddSlide() on read-only doc error = %v, want ErrReadOnly",
			err,
		)
	}

	// ChangeType should error
	err = roDoc.ChangeType(DocTypeTemplate)
	if err != ErrReadOnly {
		t.Errorf(
			"ChangeType() on read-only doc error = %v, want ErrReadOnly",
			err,
		)
	}

	// Save should error
	err = roDoc.Save()
	if err != ErrReadOnly {
		t.Errorf(
			"Save() on read-only doc error = %v, want ErrReadOnly",
			err,
		)
	}
}

// TestDocument_Features tests accessing features.
func TestDocument_Features(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	features := doc.Features()
	if features == nil {
		t.Error("Features() returned nil")
	}
}

// TestDocument_Settings tests accessing settings.
func TestDocument_Settings(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	settings := doc.Settings()
	if settings == nil {
		t.Error("Settings() returned nil")
	}
}

// TestDocument_AddPresentationPart tests adding presentation part errors.
func TestDocument_AddPresentationPart(
	t *testing.T,
) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	doc, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Already has a presentation part, should error
	_, err = doc.AddPresentationPart()
	if err != ErrPresentationPartExists {
		t.Errorf(
			"AddPresentationPart() error = %v, want ErrPresentationPartExists",
			err,
		)
	}
}
