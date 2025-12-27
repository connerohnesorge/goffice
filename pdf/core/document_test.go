package core

import (
	"bytes"
	"errors"
	"testing"
	"time"
)

func TestNewDocument(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	if doc.IsClosed() {
		t.Error(
			"NewDocument() returned closed document",
		)
	}

	if doc.PageCount() != 0 {
		t.Errorf(
			"NewDocument().PageCount() = %d, want 0",
			doc.PageCount(),
		)
	}

	if doc.Version() != V17 {
		t.Errorf(
			"NewDocument().Version() = %v, want V17",
			doc.Version(),
		)
	}
}

func TestNewDocumentWithOptions(t *testing.T) {
	tests := []struct {
		name    string
		opts    *DocumentOptions
		wantVer Version
	}{
		{
			name:    "nil options uses defaults",
			opts:    nil,
			wantVer: V17,
		},
		{
			name: "V14 version",
			opts: &DocumentOptions{
				Version: V14,
			},
			wantVer: V14,
		},
		{
			name: "V20 version",
			opts: &DocumentOptions{
				Version: V20,
			},
			wantVer: V20,
		},
		{
			name: "with metadata",
			opts: &DocumentOptions{
				Version: V17,
				Metadata: Metadata{
					Title:  "Test Document",
					Author: "Test Author",
				},
			},
			wantVer: V17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := NewDocumentWithOptions(
				tt.opts,
			)
			if err != nil {
				t.Fatalf(
					"NewDocumentWithOptions() error = %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			if doc.Version() != tt.wantVer {
				t.Errorf(
					"Version() = %v, want %v",
					doc.Version(),
					tt.wantVer,
				)
			}

			if tt.opts != nil &&
				tt.opts.Metadata.Title != "" {
				meta := doc.Metadata()
				if meta.Title != tt.opts.Metadata.Title {
					t.Errorf(
						"Metadata().Title = %q, want %q",
						meta.Title,
						tt.opts.Metadata.Title,
					)
				}
			}
		})
	}
}

func TestDocument_AddPage(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Add first page
	page1, err := doc.AddPage(PageSizeA4)
	if err != nil {
		t.Fatalf("AddPage() error = %v", err)
	}

	if page1.Number() != 1 {
		t.Errorf(
			"First page Number() = %d, want 1",
			page1.Number(),
		)
	}

	if doc.PageCount() != 1 {
		t.Errorf(
			"PageCount() = %d, want 1",
			doc.PageCount(),
		)
	}

	// Add second page
	page2, err := doc.AddPage(PageSizeLetter)
	if err != nil {
		t.Fatalf("AddPage() error = %v", err)
	}

	if page2.Number() != 2 {
		t.Errorf(
			"Second page Number() = %d, want 2",
			page2.Number(),
		)
	}

	if doc.PageCount() != 2 {
		t.Errorf(
			"PageCount() = %d, want 2",
			doc.PageCount(),
		)
	}
}

func TestDocument_AddPage_InvalidSize(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	tests := []struct {
		name string
		size PageSize
	}{
		{
			"zero width",
			PageSize{Width: 0, Height: 100},
		},
		{
			"zero height",
			PageSize{Width: 100, Height: 0},
		},
		{
			"negative width",
			PageSize{Width: -100, Height: 100},
		},
		{
			"negative height",
			PageSize{Width: 100, Height: -100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := doc.AddPage(tt.size)
			if !errors.Is(
				err,
				ErrInvalidPageSize,
			) {
				t.Errorf(
					"AddPage(%v) error = %v, want ErrInvalidPageSize",
					tt.size,
					err,
				)
			}
		})
	}
}

func TestDocument_AddPageConvenience(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	pageA4, err := doc.AddPageA4()
	if err != nil {
		t.Fatalf("AddPageA4() error = %v", err)
	}

	if pageA4.Width() != PageSizeA4.Width ||
		pageA4.Height() != PageSizeA4.Height {
		t.Errorf(
			"AddPageA4() size = %vx%v, want %vx%v",
			pageA4.Width(),
			pageA4.Height(),
			PageSizeA4.Width,
			PageSizeA4.Height,
		)
	}

	pageLetter, err := doc.AddPageLetter()
	if err != nil {
		t.Fatalf(
			"AddPageLetter() error = %v",
			err,
		)
	}

	if pageLetter.Width() != PageSizeLetter.Width ||
		pageLetter.Height() != PageSizeLetter.Height {
		t.Errorf(
			"AddPageLetter() size = %vx%v, want %vx%v",
			pageLetter.Width(),
			pageLetter.Height(),
			PageSizeLetter.Width,
			PageSizeLetter.Height,
		)
	}
}

func TestDocument_GetPage(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	page1, _ := doc.AddPage(PageSizeA4)
	page2, _ := doc.AddPage(PageSizeLetter)

	// Get first page
	got1, err := doc.GetPage(0)
	if err != nil {
		t.Fatalf("GetPage(0) error = %v", err)
	}
	if got1 != page1 {
		t.Error("GetPage(0) returned wrong page")
	}

	// Get second page
	got2, err := doc.GetPage(1)
	if err != nil {
		t.Fatalf("GetPage(1) error = %v", err)
	}
	if got2 != page2 {
		t.Error("GetPage(1) returned wrong page")
	}

	// Invalid index
	_, err = doc.GetPage(-1)
	if !errors.Is(err, ErrInvalidPageIndex) {
		t.Errorf(
			"GetPage(-1) error = %v, want ErrInvalidPageIndex",
			err,
		)
	}

	_, err = doc.GetPage(2)
	if !errors.Is(err, ErrInvalidPageIndex) {
		t.Errorf(
			"GetPage(2) error = %v, want ErrInvalidPageIndex",
			err,
		)
	}
}

func TestDocument_Pages(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	page1, _ := doc.AddPage(PageSizeA4)
	page2, _ := doc.AddPage(PageSizeLetter)

	pages := doc.Pages()
	if len(pages) != 2 {
		t.Fatalf(
			"Pages() returned %d pages, want 2",
			len(pages),
		)
	}

	if pages[0] != page1 || pages[1] != page2 {
		t.Error("Pages() returned wrong pages")
	}
}

func TestDocument_Metadata(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Test SetTitle
	if err := doc.SetTitle("Test Title"); err != nil {
		t.Errorf("SetTitle() error = %v", err)
	}
	if doc.Metadata().Title != "Test Title" {
		t.Errorf(
			"Title = %q, want %q",
			doc.Metadata().Title,
			"Test Title",
		)
	}

	// Test SetAuthor
	if err := doc.SetAuthor("Test Author"); err != nil {
		t.Errorf("SetAuthor() error = %v", err)
	}
	if doc.Metadata().Author != "Test Author" {
		t.Errorf(
			"Author = %q, want %q",
			doc.Metadata().Author,
			"Test Author",
		)
	}

	// Test SetSubject
	if err := doc.SetSubject("Test Subject"); err != nil {
		t.Errorf("SetSubject() error = %v", err)
	}
	if doc.Metadata().Subject != "Test Subject" {
		t.Errorf(
			"Subject = %q, want %q",
			doc.Metadata().Subject,
			"Test Subject",
		)
	}

	// Test SetKeywords
	if err := doc.SetKeywords("test, keywords"); err != nil {
		t.Errorf("SetKeywords() error = %v", err)
	}
	if doc.Metadata().Keywords != "test, keywords" {
		t.Errorf(
			"Keywords = %q, want %q",
			doc.Metadata().Keywords,
			"test, keywords",
		)
	}

	// Test SetCreator
	if err := doc.SetCreator("Test Creator"); err != nil {
		t.Errorf("SetCreator() error = %v", err)
	}
	if doc.Metadata().Creator != "Test Creator" {
		t.Errorf(
			"Creator = %q, want %q",
			doc.Metadata().Creator,
			"Test Creator",
		)
	}

	// Test SetProducer
	if err := doc.SetProducer("Test Producer"); err != nil {
		t.Errorf("SetProducer() error = %v", err)
	}
	if doc.Metadata().Producer != "Test Producer" {
		t.Errorf(
			"Producer = %q, want %q",
			doc.Metadata().Producer,
			"Test Producer",
		)
	}

	// Test SetMetadata
	newMeta := Metadata{
		Title:    "New Title",
		Author:   "New Author",
		Subject:  "New Subject",
		Keywords: "new, keywords",
		Creator:  "New Creator",
		Producer: "New Producer",
		CreationDate: time.Date(
			2024,
			1,
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		),
	}
	if err := doc.SetMetadata(newMeta); err != nil {
		t.Errorf("SetMetadata() error = %v", err)
	}

	got := doc.Metadata()
	if got.Title != newMeta.Title {
		t.Errorf(
			"Metadata().Title = %q, want %q",
			got.Title,
			newMeta.Title,
		)
	}
	if got.Author != newMeta.Author {
		t.Errorf(
			"Metadata().Author = %q, want %q",
			got.Author,
			newMeta.Author,
		)
	}
}

func TestDocument_Close(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}

	// Add a page before closing
	_, _ = doc.AddPage(PageSizeA4)

	// First close should succeed
	if err := doc.Close(); err != nil {
		t.Errorf("Close() error = %v", err)
	}

	if !doc.IsClosed() {
		t.Error(
			"IsClosed() = false after Close()",
		)
	}

	// Second close should also succeed (idempotent)
	if err := doc.Close(); err != nil {
		t.Errorf("Second Close() error = %v", err)
	}

	// Operations on closed document should fail
	_, err = doc.AddPage(PageSizeA4)
	if !errors.Is(err, ErrDocumentClosed) {
		t.Errorf(
			"AddPage() on closed doc error = %v, want ErrDocumentClosed",
			err,
		)
	}

	_, err = doc.GetPage(0)
	if !errors.Is(err, ErrDocumentClosed) {
		t.Errorf(
			"GetPage() on closed doc error = %v, want ErrDocumentClosed",
			err,
		)
	}

	if err := doc.SetTitle("Test"); !errors.Is(
		err,
		ErrDocumentClosed,
	) {
		t.Errorf(
			"SetTitle() on closed doc error = %v, want ErrDocumentClosed",
			err,
		)
	}
}

func TestDocument_Write_NoPages(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	var buf bytes.Buffer
	err = doc.Write(&buf)
	if !errors.Is(err, ErrNoPages) {
		t.Errorf(
			"Write() with no pages error = %v, want ErrNoPages",
			err,
		)
	}
}

func TestDocument_Write_NilWriter(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	_, _ = doc.AddPage(PageSizeA4)

	err = doc.Write(nil)
	if !errors.Is(err, ErrNilWriter) {
		t.Errorf(
			"Write(nil) error = %v, want ErrNilWriter",
			err,
		)
	}
}

func TestDocument_Write(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Set metadata
	_ = doc.SetTitle("Test PDF")
	_ = doc.SetAuthor("Test Author")

	// Add pages
	_, err = doc.AddPage(PageSizeA4)
	if err != nil {
		t.Fatalf("AddPage() error = %v", err)
	}

	_, err = doc.AddPage(PageSizeLetter)
	if err != nil {
		t.Fatalf("AddPage() error = %v", err)
	}

	// Write to buffer
	var buf bytes.Buffer
	if err := doc.Write(&buf); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Verify PDF header
	data := buf.Bytes()
	if len(data) < 8 {
		t.Fatal(
			"Write() produced too little output",
		)
	}

	header := string(data[:8])
	if header != "%PDF-1.7" {
		t.Errorf(
			"PDF header = %q, want %%PDF-1.7",
			header,
		)
	}

	// Verify PDF has EOF marker
	if !bytes.Contains(data, []byte("%%EOF")) {
		t.Error("PDF output missing EOF marker")
	}
}

func TestDocument_Write_ClosedDocument(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}

	_, _ = doc.AddPage(PageSizeA4)
	_ = doc.Close()

	var buf bytes.Buffer
	err = doc.Write(&buf)
	if !errors.Is(err, ErrDocumentClosed) {
		t.Errorf(
			"Write() on closed doc error = %v, want ErrDocumentClosed",
			err,
		)
	}
}

func TestPageSize_Landscape(t *testing.T) {
	landscape := PageSizeA4.Landscape()

	if landscape.Width != PageSizeA4.Height ||
		landscape.Height != PageSizeA4.Width {
		t.Errorf(
			"Landscape() = %vx%v, want %vx%v",
			landscape.Width,
			landscape.Height,
			PageSizeA4.Height,
			PageSizeA4.Width,
		)
	}

	if !landscape.IsLandscape() {
		t.Error(
			"Landscape().IsLandscape() = false",
		)
	}

	if PageSizeA4.IsLandscape() {
		t.Error(
			"Portrait A4.IsLandscape() = true",
		)
	}
}

func TestPage_Content(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	page, _ := doc.AddPage(PageSizeA4)

	// Initially empty
	if len(page.Content()) != 0 {
		t.Errorf(
			"Initial Content() length = %d, want 0",
			len(page.Content()),
		)
	}

	// Write content
	testContent := []byte(
		"BT /F1 12 Tf 100 700 Td (Hello) Tj ET",
	)
	n, err := page.WriteContent(testContent)
	if err != nil {
		t.Fatalf("WriteContent() error = %v", err)
	}
	if n != len(testContent) {
		t.Errorf(
			"WriteContent() wrote %d bytes, want %d",
			n,
			len(testContent),
		)
	}

	// Verify content
	content := page.Content()
	if !bytes.Equal(content, testContent) {
		t.Errorf(
			"Content() = %q, want %q",
			content,
			testContent,
		)
	}
}

func TestPage_WriteContentString(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	page, _ := doc.AddPage(PageSizeA4)

	testString := "BT /F1 12 Tf 100 700 Td (Hello) Tj ET"
	n, err := page.WriteContentString(testString)
	if err != nil {
		t.Fatalf(
			"WriteContentString() error = %v",
			err,
		)
	}
	if n != len(testString) {
		t.Errorf(
			"WriteContentString() wrote %d bytes, want %d",
			n,
			len(testString),
		)
	}

	content := page.Content()
	if string(content) != testString {
		t.Errorf(
			"Content() = %q, want %q",
			string(content),
			testString,
		)
	}
}

func TestVersion_String(t *testing.T) {
	tests := []struct {
		version Version
		want    string
	}{
		{V14, "1.4"},
		{V15, "1.5"},
		{V16, "1.6"},
		{V17, "1.7"},
		{V20, "2.0"},
		{
			Version(99),
			"1.7",
		}, // Unknown version defaults to 1.7
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.version.String(); got != tt.want {
				t.Errorf(
					"Version(%d).String() = %q, want %q",
					tt.version,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	// Test DocumentError
	docErr := NewDocumentError(
		"test operation",
		errors.New("underlying error"),
	)
	if docErr.Error() != "core: test operation: underlying error" {
		t.Errorf(
			"DocumentError.Error() = %q",
			docErr.Error(),
		)
	}
	if docErr.Unwrap() == nil {
		t.Error("DocumentError.Unwrap() = nil")
	}

	// Test PageError
	pageErr := NewPageError(
		1,
		"test operation",
		errors.New("underlying error"),
	)
	if pageErr.Error() != "core: page 1: test operation: underlying error" {
		t.Errorf(
			"PageError.Error() = %q",
			pageErr.Error(),
		)
	}
	if pageErr.Unwrap() == nil {
		t.Error("PageError.Unwrap() = nil")
	}

	// Test WriteError
	writeErr := NewWriteError(
		"test operation",
		errors.New("underlying error"),
	)
	if writeErr.Error() != "core: write test operation: underlying error" {
		t.Errorf(
			"WriteError.Error() = %q",
			writeErr.Error(),
		)
	}
	if writeErr.Unwrap() == nil {
		t.Error(
			"WriteError.Unwrap() returned nil",
		)
	}

	// Test errors without underlying error
	docErr2 := NewDocumentError("test", nil)
	if docErr2.Error() != "core: test" {
		t.Errorf(
			"DocumentError without wrapped error: %q",
			docErr2.Error(),
		)
	}
}

func TestDefaultDocumentOptions(t *testing.T) {
	opts := DefaultDocumentOptions()

	if opts.Version != V17 {
		t.Errorf(
			"DefaultDocumentOptions().Version = %v, want V17",
			opts.Version,
		)
	}

	if !opts.CompressContent {
		t.Error(
			"DefaultDocumentOptions().CompressContent = false, want true",
		)
	}

	if opts.Metadata.Producer != "goffice-pdf" {
		t.Errorf(
			"DefaultDocumentOptions().Metadata.Producer = %q, want goffice-pdf",
			opts.Metadata.Producer,
		)
	}

	if opts.Metadata.CreationDate.IsZero() {
		t.Error(
			"DefaultDocumentOptions().Metadata.CreationDate is zero",
		)
	}
}

func TestDocument_AddPageWithOptions(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	t.Run(
		"with nil options uses defaults",
		func(t *testing.T) {
			page, err := doc.AddPageWithOptions(
				nil,
			)
			if err != nil {
				t.Fatalf(
					"AddPageWithOptions(nil) error = %v",
					err,
				)
			}
			if page.Size() != PageSizeA4 {
				t.Errorf(
					"Page size = %+v, want A4",
					page.Size(),
				)
			}
		},
	)

	t.Run(
		"with custom options",
		func(t *testing.T) {
			opts := NewPageOptions(
				PageSizeLetter,
			).
				WithMargins(NewMarginsUniform(72)).
				WithRotation(90).
				WithUserUnit(2.0)

			page, err := doc.AddPageWithOptions(
				opts,
			)
			if err != nil {
				t.Fatalf(
					"AddPageWithOptions() error = %v",
					err,
				)
			}
			if page.Size() != PageSizeLetter {
				t.Errorf(
					"Page size = %+v, want Letter",
					page.Size(),
				)
			}
			if page.Margins().Top != 72 {
				t.Errorf(
					"Margins.Top = %v, want 72",
					page.Margins().Top,
				)
			}
			if page.Rotation() != 90 {
				t.Errorf(
					"Rotation = %d, want 90",
					page.Rotation(),
				)
			}
			if page.UserUnit() != 2.0 {
				t.Errorf(
					"UserUnit = %v, want 2.0",
					page.UserUnit(),
				)
			}
		},
	)

	t.Run("with landscape", func(t *testing.T) {
		opts := NewPageOptions(
			PageSizeA4,
		).Landscape()
		page, err := doc.AddPageWithOptions(opts)
		if err != nil {
			t.Fatalf(
				"AddPageWithOptions() error = %v",
				err,
			)
		}
		if !page.IsLandscape() {
			t.Error("Page should be landscape")
		}
	})

	t.Run("with page boxes", func(t *testing.T) {
		boxes := NewPageBoxesFromSize(
			PageSizeLetter,
		).
			WithCropBox(NewRectangle(18, 18, 594, 774)).
			WithBleed(9)

		opts := NewPageOptions(PageSizeLetter).
			WithBoxes(boxes)

		page, err := doc.AddPageWithOptions(opts)
		if err != nil {
			t.Fatalf(
				"AddPageWithOptions() error = %v",
				err,
			)
		}
		if page.Boxes() == nil {
			t.Fatal("Page boxes should be set")
		}
		if page.Boxes().CropBox == nil {
			t.Error("CropBox should be set")
		}
		if page.Boxes().BleedBox == nil {
			t.Error("BleedBox should be set")
		}
	})

	t.Run(
		"with invalid size",
		func(t *testing.T) {
			opts := &PageOptions{
				Size: PageSize{
					Width:  0,
					Height: 100,
				},
			}
			_, err := doc.AddPageWithOptions(opts)
			if !errors.Is(
				err,
				ErrInvalidPageSize,
			) {
				t.Errorf(
					"Expected ErrInvalidPageSize, got %v",
					err,
				)
			}
		},
	)
}

func TestDocument_ConvenienceMethods(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	tests := []struct {
		name     string
		addFunc  func() (*Page, error)
		wantSize PageSize
	}{
		{"A3", doc.AddPageA3, PageSizeA3},
		{"A5", doc.AddPageA5, PageSizeA5},
		{
			"Legal",
			doc.AddPageLegal,
			PageSizeLegal,
		},
		{
			"Tabloid",
			doc.AddPageTabloid,
			PageSizeTabloid,
		},
		{
			"Executive",
			doc.AddPageExecutive,
			PageSizeExecutive,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			page, err := tt.addFunc()
			if err != nil {
				t.Fatalf(
					"AddPage%s() error = %v",
					tt.name,
					err,
				)
			}
			if page.Size().Width != tt.wantSize.Width ||
				page.Size().Height != tt.wantSize.Height {
				t.Errorf(
					"Page size = %+v, want %+v",
					page.Size(),
					tt.wantSize,
				)
			}
		})
	}
}

func TestPage_MarginsAndContentArea(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	margins := NewMargins(72, 54, 72, 54)
	opts := NewPageOptions(
		PageSizeLetter,
	).WithMargins(margins)

	page, err := doc.AddPageWithOptions(opts)
	if err != nil {
		t.Fatalf(
			"AddPageWithOptions() error = %v",
			err,
		)
	}

	// Check margins
	if page.Margins() != margins {
		t.Errorf(
			"Margins = %+v, want %+v",
			page.Margins(),
			margins,
		)
	}

	// Check content area
	area := page.ContentArea()
	if area.LLX != 54 {
		t.Errorf(
			"ContentArea.LLX = %v, want 54",
			area.LLX,
		)
	}
	if area.LLY != 72 {
		t.Errorf(
			"ContentArea.LLY = %v, want 72",
			area.LLY,
		)
	}

	// Check content dimensions
	expectedWidth := 612 - 54 - 54 // Letter width - left - right
	if page.ContentWidth() != float64(
		expectedWidth,
	) {
		t.Errorf(
			"ContentWidth = %v, want %v",
			page.ContentWidth(),
			expectedWidth,
		)
	}

	expectedHeight := 792 - 72 - 72 // Letter height - top - bottom
	if page.ContentHeight() != float64(
		expectedHeight,
	) {
		t.Errorf(
			"ContentHeight = %v, want %v",
			page.ContentHeight(),
			expectedHeight,
		)
	}
}

func TestPage_SettersAndGetters(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	page, err := doc.AddPageA4()
	if err != nil {
		t.Fatalf("AddPageA4() error = %v", err)
	}

	t.Run("SetMargins", func(t *testing.T) {
		margins := NewMarginsUniform(50)
		page.SetMargins(margins)
		if page.Margins() != margins {
			t.Errorf(
				"Margins = %+v, want %+v",
				page.Margins(),
				margins,
			)
		}
	})

	t.Run("SetRotation", func(t *testing.T) {
		page.SetRotation(90)
		if page.Rotation() != 90 {
			t.Errorf(
				"Rotation = %d, want 90",
				page.Rotation(),
			)
		}

		page.SetRotation(45) // Should round to 90
		if page.Rotation() != 90 {
			t.Errorf(
				"Rotation(45) = %d, want 90",
				page.Rotation(),
			)
		}

		page.SetRotation(
			-90,
		) // Should normalize to 270
		if page.Rotation() != 270 {
			t.Errorf(
				"Rotation(-90) = %d, want 270",
				page.Rotation(),
			)
		}
	})

	t.Run("SetUserUnit", func(t *testing.T) {
		page.SetUserUnit(2.0)
		if page.UserUnit() != 2.0 {
			t.Errorf(
				"UserUnit = %v, want 2.0",
				page.UserUnit(),
			)
		}

		page.SetUserUnit(
			0.5,
		) // Should clamp to 1.0
		if page.UserUnit() != 1.0 {
			t.Errorf(
				"UserUnit(0.5) = %v, want 1.0",
				page.UserUnit(),
			)
		}

		page.SetUserUnit(
			100000,
		) // Should clamp to 75000
		if page.UserUnit() != 75000 {
			t.Errorf(
				"UserUnit(100000) = %v, want 75000",
				page.UserUnit(),
			)
		}
	})

	t.Run("SetBoxes", func(t *testing.T) {
		boxes := NewPageBoxesFromSize(PageSizeA4)
		page.SetBoxes(&boxes)
		if page.Boxes() == nil {
			t.Error("Boxes should be set")
		}
	})
}

func TestPage_CoordinateSystem(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	page, err := doc.AddPageLetter()
	if err != nil {
		t.Fatalf(
			"AddPageLetter() error = %v",
			err,
		)
	}

	cs := page.CoordinateSystem()
	if cs == nil {
		t.Fatal(
			"CoordinateSystem should not be nil",
		)
	}

	if cs.PageHeight != page.Height() {
		t.Errorf(
			"CoordinateSystem.PageHeight = %v, want %v",
			cs.PageHeight,
			page.Height(),
		)
	}

	if cs.UserUnit != page.UserUnit() {
		t.Errorf(
			"CoordinateSystem.UserUnit = %v, want %v",
			cs.UserUnit,
			page.UserUnit(),
		)
	}
}

func TestPage_Orientation(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	t.Run("Portrait page", func(t *testing.T) {
		page, _ := doc.AddPageA4()
		if !page.IsPortrait() {
			t.Error("A4 should be portrait")
		}
		if page.IsLandscape() {
			t.Error("A4 should not be landscape")
		}
		if page.Orientation() != OrientationPortrait {
			t.Error(
				"A4 orientation should be portrait",
			)
		}
	})

	t.Run("Landscape page", func(t *testing.T) {
		opts := NewPageOptions(
			PageSizeA4,
		).Landscape()
		page, _ := doc.AddPageWithOptions(opts)
		if page.IsPortrait() {
			t.Error(
				"Landscape page should not be portrait",
			)
		}
		if !page.IsLandscape() {
			t.Error(
				"Landscape page should be landscape",
			)
		}
		if page.Orientation() != OrientationLandscape {
			t.Error(
				"Landscape page orientation should be landscape",
			)
		}
	})
}

func TestPage_MediaAndCropBox(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	t.Run(
		"without custom boxes",
		func(t *testing.T) {
			page, _ := doc.AddPageA4()
			mediaBox := page.MediaBox()
			if mediaBox.Width() != PageSizeA4.Width ||
				mediaBox.Height() != PageSizeA4.Height {
				t.Errorf(
					"MediaBox size = %vx%v, want %vx%v",
					mediaBox.Width(),
					mediaBox.Height(),
					PageSizeA4.Width,
					PageSizeA4.Height,
				)
			}
			cropBox := page.CropBox()
			if cropBox != mediaBox {
				t.Error(
					"CropBox should equal MediaBox when not set",
				)
			}
		},
	)

	t.Run(
		"with custom boxes",
		func(t *testing.T) {
			crop := NewRectangle(18, 18, 577, 824)
			boxes := NewPageBoxesFromSize(
				PageSizeA4,
			).WithCropBox(crop)
			opts := NewPageOptions(
				PageSizeA4,
			).WithBoxes(boxes)
			page, _ := doc.AddPageWithOptions(
				opts,
			)

			cropBox := page.CropBox()
			if cropBox != crop {
				t.Errorf(
					"CropBox = %+v, want %+v",
					cropBox,
					crop,
				)
			}
		},
	)
}

func TestDocument_Write_WithPageOptions(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Add page with various options
	boxes := NewPageBoxesFromSize(PageSizeLetter).
		WithCropBox(NewRectangle(18, 18, 594, 774)).
		WithTrimBox(NewRectangle(27, 27, 585, 765))

	opts := NewPageOptions(PageSizeLetter).
		WithMargins(DefaultMargins()).
		WithRotation(90).
		WithBoxes(boxes)

	page, err := doc.AddPageWithOptions(opts)
	if err != nil {
		t.Fatalf(
			"AddPageWithOptions() error = %v",
			err,
		)
	}

	// Verify page has the correct options set
	if page.Rotation() != 90 {
		t.Errorf(
			"Page rotation = %d, want 90",
			page.Rotation(),
		)
	}
	if page.Boxes() == nil {
		t.Error("Page boxes should be set")
	}
	if page.Boxes().CropBox == nil {
		t.Error("CropBox should be set")
	}
	if page.Boxes().TrimBox == nil {
		t.Error("TrimBox should be set")
	}

	// Write should succeed
	var buf bytes.Buffer
	if err := doc.Write(&buf); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Verify PDF was created with valid structure
	data := buf.Bytes()
	if len(data) < 100 {
		t.Error("PDF output too small")
	}

	// Check for valid PDF header
	if !bytes.HasPrefix(data, []byte("%PDF-")) {
		t.Error("PDF should start with %PDF-")
	}

	// Check for valid PDF EOF
	if !bytes.Contains(data, []byte("%%EOF")) {
		t.Error("PDF should contain EOF marker")
	}
}

func TestDocument_Write_WithUserUnit(
	t *testing.T,
) {
	// Test with PDF 1.6+ which supports UserUnit
	opts := &DocumentOptions{
		Version: V16,
	}
	doc, err := NewDocumentWithOptions(opts)
	if err != nil {
		t.Fatalf(
			"NewDocumentWithOptions() error = %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	pageOpts := NewPageOptions(
		PageSizeA4,
	).WithUserUnit(2.0)
	page, err := doc.AddPageWithOptions(pageOpts)
	if err != nil {
		t.Fatalf(
			"AddPageWithOptions() error = %v",
			err,
		)
	}

	// Verify UserUnit is set on the page
	if page.UserUnit() != 2.0 {
		t.Errorf(
			"Page UserUnit = %v, want 2.0",
			page.UserUnit(),
		)
	}

	var buf bytes.Buffer
	if err := doc.Write(&buf); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Verify PDF was created
	data := buf.Bytes()
	if len(data) < 100 {
		t.Error("PDF output too small")
	}

	// Check for valid PDF header (version may vary based on pdfcpu internals)
	if !bytes.HasPrefix(data, []byte("%PDF-1.")) {
		t.Errorf(
			"PDF should start with valid PDF header, got %s",
			string(data[:8]),
		)
	}
}
