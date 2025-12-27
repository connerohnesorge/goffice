package pdf

import (
	"bytes"
	"testing"
)

func TestFontEmbedModeString(t *testing.T) {
	tests := []struct {
		mode     FontEmbedMode
		expected string
	}{
		{EmbedSubset, "EmbedSubset"},
		{EmbedFull, "EmbedFull"},
		{NoEmbed, "NoEmbed"},
		{
			FontEmbedMode(99),
			"EmbedSubset",
		}, // Invalid defaults to EmbedSubset
	}

	for _, tt := range tests {
		if got := tt.mode.String(); got != tt.expected {
			t.Errorf(
				"FontEmbedMode(%d).String() = %q, want %q",
				tt.mode,
				got,
				tt.expected,
			)
		}
	}
}

func TestAComplianceString(t *testing.T) {
	tests := []struct {
		compliance ACompliance
		expected   string
	}{
		{PDFANone, "None"},
		{PDFA1b, "PDF/A-1b"},
		{PDFA2b, "PDF/A-2b"},
		{PDFA3b, "PDF/A-3b"},
		{
			ACompliance(99),
			"None",
		}, // Invalid defaults to None
	}

	for _, tt := range tests {
		if got := tt.compliance.String(); got != tt.expected {
			t.Errorf(
				"ACompliance(%d).String() = %q, want %q",
				tt.compliance,
				got,
				tt.expected,
			)
		}
	}
}

func TestDefaultRenderOptions(t *testing.T) {
	opts := DefaultRenderOptions()

	if opts.ImageDPI != 150 {
		t.Errorf(
			"ImageDPI = %d, want 150",
			opts.ImageDPI,
		)
	}
	if opts.ImageCacheLimit != 100*1024*1024 {
		t.Errorf(
			"ImageCacheLimit = %d, want %d",
			opts.ImageCacheLimit,
			100*1024*1024,
		)
	}
	if opts.FontEmbedding != EmbedSubset {
		t.Errorf(
			"FontEmbedding = %v, want EmbedSubset",
			opts.FontEmbedding,
		)
	}
	if opts.ACompliance != PDFANone {
		t.Errorf(
			"ACompliance = %v, want PDFANone",
			opts.ACompliance,
		)
	}
	if opts.TaggedPDF {
		t.Errorf(
			"TaggedPDF = %v, want false",
			opts.TaggedPDF,
		)
	}
	if opts.Creator != "goffice-pdf" {
		t.Errorf(
			"Creator = %q, want %q",
			opts.Creator,
			"goffice-pdf",
		)
	}
	if !opts.CompressContent {
		t.Errorf(
			"CompressContent = %v, want true",
			opts.CompressContent,
		)
	}
}

func TestRenderOptionsWithMethods(t *testing.T) {
	opts := DefaultRenderOptions()

	// Test WithImageDPI
	opts2 := opts.WithImageDPI(300)
	if opts2.ImageDPI != 300 {
		t.Errorf(
			"WithImageDPI: got %d, want 300",
			opts2.ImageDPI,
		)
	}
	if opts.ImageDPI != 150 {
		t.Error(
			"WithImageDPI modified original options",
		)
	}

	// Test WithFontEmbedding
	opts3 := opts.WithFontEmbedding(EmbedFull)
	if opts3.FontEmbedding != EmbedFull {
		t.Errorf(
			"WithFontEmbedding: got %v, want EmbedFull",
			opts3.FontEmbedding,
		)
	}

	// Test WithACompliance
	opts4 := opts.WithACompliance(PDFA2b)
	if opts4.ACompliance != PDFA2b {
		t.Errorf(
			"WithACompliance: got %v, want PDFA2b",
			opts4.ACompliance,
		)
	}

	// Test WithTaggedPDF
	opts5 := opts.WithTaggedPDF(true)
	if !opts5.TaggedPDF {
		t.Errorf(
			"WithTaggedPDF: got %v, want true",
			opts5.TaggedPDF,
		)
	}

	// Test WithImageCacheLimit
	opts6 := opts.WithImageCacheLimit(
		200 * 1024 * 1024,
	)
	if opts6.ImageCacheLimit != 200*1024*1024 {
		t.Errorf(
			"WithImageCacheLimit: got %d, want %d",
			opts6.ImageCacheLimit,
			200*1024*1024,
		)
	}

	// Test WithMetadata
	opts7 := opts.WithMetadata(
		"Title",
		"Author",
		"Subject",
		"Keywords",
	)
	if opts7.Title != "Title" ||
		opts7.Author != "Author" ||
		opts7.Subject != "Subject" ||
		opts7.Keywords != "Keywords" {
		t.Error(
			"WithMetadata did not set metadata correctly",
		)
	}

	// Test WithOnPageComplete
	callback := func(_ int, _ []byte) {}
	opts8 := opts.WithOnPageComplete(callback)
	if opts8.OnPageComplete == nil {
		t.Error(
			"WithOnPageComplete did not set callback",
		)
	}
}

func TestRenderWordNilDocument(t *testing.T) {
	var buf bytes.Buffer
	err := RenderWord(nil, &buf, nil)
	if err != ErrNilDocument {
		t.Errorf(
			"RenderWord(nil) = %v, want ErrNilDocument",
			err,
		)
	}
}

func TestRenderSpreadsheetNilDocument(
	t *testing.T,
) {
	var buf bytes.Buffer
	err := RenderSpreadsheet(nil, &buf, nil)
	if err != ErrNilDocument {
		t.Errorf(
			"RenderSpreadsheet(nil) = %v, want ErrNilDocument",
			err,
		)
	}
}

func TestRenderPresentationNilDocument(
	t *testing.T,
) {
	var buf bytes.Buffer
	err := RenderPresentation(nil, &buf, nil)
	if err != ErrNilDocument {
		t.Errorf(
			"RenderPresentation(nil) = %v, want ErrNilDocument",
			err,
		)
	}
}

func TestRenderWordToFileNilDocument(
	t *testing.T,
) {
	err := RenderWordToFile(
		nil,
		"output.pdf",
		nil,
	)
	if err != ErrNilDocument {
		t.Errorf(
			"RenderWordToFile(nil) = %v, want ErrNilDocument",
			err,
		)
	}
}

func TestRenderSpreadsheetToFileNilDocument(
	t *testing.T,
) {
	err := RenderSpreadsheetToFile(
		nil,
		"output.pdf",
		nil,
	)
	if err != ErrNilDocument {
		t.Errorf(
			"RenderSpreadsheetToFile(nil) = %v, want ErrNilDocument",
			err,
		)
	}
}

func TestRenderPresentationToFileNilDocument(
	t *testing.T,
) {
	err := RenderPresentationToFile(
		nil,
		"output.pdf",
		nil,
	)
	if err != ErrNilDocument {
		t.Errorf(
			"RenderPresentationToFile(nil) = %v, want ErrNilDocument",
			err,
		)
	}
}

func TestVersion(t *testing.T) {
	if Version != "0.1.0" {
		t.Errorf(
			"Version = %q, want %q",
			Version,
			"0.1.0",
		)
	}
}
