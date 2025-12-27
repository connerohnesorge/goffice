package parts

import (
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
)

// createTestPresentationPart creates a test presentation part for testing.
func createTestPresentationPart(
	t *testing.T,
) (*PresentationPart, func()) {
	t.Helper()
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.pptx")

	pkg, err := openxml.CreatePackage(path)
	if err != nil {
		t.Fatalf(
			"CreatePackage() error = %v",
			err,
		)
	}

	presPart, err := NewPresentationPart(
		"/ppt/presentation.xml",
		ContentTypePresentation,
		pkg,
	)
	if err != nil {
		_ = pkg.Close()
		t.Fatalf(
			"NewPresentationPart() error = %v",
			err,
		)
	}

	presPart.InitializeContent()
	pkg.SetMainPart(presPart)

	cleanup := func() {
		_ = pkg.Close()
	}

	return presPart, cleanup
}

// TestPresentationPart_AddSlidePart tests adding slide parts.
func TestPresentationPart_AddSlidePart(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Add first slide
	slide1, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}
	if slide1 == nil {
		t.Error("AddSlidePart() returned nil")
	}

	// Verify content type
	if slide1.FixedContentType() != ContentTypeSlide {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			slide1.FixedContentType(),
			ContentTypeSlide,
		)
	}

	// Verify URI pattern
	if slide1.URI() == "" {
		t.Error("URI() is empty")
	}

	// Add second slide
	slide2, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf(
			"AddSlidePart() second error = %v",
			err,
		)
	}
	if slide2 == nil {
		t.Fatal(
			"AddSlidePart() second returned nil",
		)
	}

	// URIs should be different
	if slide1.URI() == slide2.URI() {
		t.Error("Slide URIs should be different")
	}
}

// TestPresentationPart_SlideParts tests iterating slide parts.
func TestPresentationPart_SlideParts(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Initially no slides
	slides := presPart.SlideParts()
	if len(slides) != 0 {
		t.Errorf(
			"SlideParts() = %d, want 0",
			len(slides),
		)
	}

	// Add 3 slides
	for range 3 {
		_, err := presPart.AddSlidePart()
		if err != nil {
			t.Fatalf(
				"AddSlidePart() error = %v",
				err,
			)
		}
	}

	slides = presPart.SlideParts()
	if len(slides) != 3 {
		t.Errorf(
			"SlideParts() = %d, want 3",
			len(slides),
		)
	}

	// Verify each slide is valid
	for i, slide := range slides {
		if slide == nil {
			t.Errorf("SlideParts()[%d] is nil", i)
		}
	}
}

// TestSlidePart_AddNotesSlidePart tests adding notes to slides.
func TestSlidePart_AddNotesSlidePart(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Add a slide first
	slide, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	// Initially no notes
	if slide.NotesSlidePart() != nil {
		t.Error(
			"NotesSlidePart() should be nil initially",
		)
	}

	// Add notes slide
	notes, err := slide.AddNotesSlidePart()
	if err != nil {
		t.Fatalf(
			"AddNotesSlidePart() error = %v",
			err,
		)
	}
	if notes == nil {
		t.Error(
			"AddNotesSlidePart() returned nil",
		)
	}

	// Verify content type
	if notes.FixedContentType() != ContentTypeNotesSlide {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			notes.FixedContentType(),
			ContentTypeNotesSlide,
		)
	}

	// Now should have notes
	if slide.NotesSlidePart() == nil {
		t.Error(
			"NotesSlidePart() should not be nil after adding",
		)
	}
}

// TestSlideMasterPart_AddSlideLayoutPart tests adding layouts to masters.
func TestSlideMasterPart_AddSlideLayoutPart(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Add a slide master first
	master, err := presPart.AddSlideMasterPart()
	if err != nil {
		t.Fatalf(
			"AddSlideMasterPart() error = %v",
			err,
		)
	}
	if master == nil {
		t.Fatal(
			"AddSlideMasterPart() returned nil",
		)
	}

	// Verify master content type
	if master.FixedContentType() != ContentTypeSlideMaster {
		t.Errorf(
			"Master FixedContentType() = %v, want %v",
			master.FixedContentType(),
			ContentTypeSlideMaster,
		)
	}

	// Initially no layouts
	layouts := master.SlideLayoutParts()
	if len(layouts) != 0 {
		t.Errorf(
			"SlideLayoutParts() = %d, want 0",
			len(layouts),
		)
	}

	// Add a layout
	layout, err := master.AddSlideLayoutPart()
	if err != nil {
		t.Fatalf(
			"AddSlideLayoutPart() error = %v",
			err,
		)
	}
	if layout == nil {
		t.Error(
			"AddSlideLayoutPart() returned nil",
		)
	}

	// Verify layout content type
	if layout.FixedContentType() != ContentTypeSlideLayout {
		t.Errorf(
			"Layout FixedContentType() = %v, want %v",
			layout.FixedContentType(),
			ContentTypeSlideLayout,
		)
	}

	// Now should have 1 layout
	layouts = master.SlideLayoutParts()
	if len(layouts) != 1 {
		t.Errorf(
			"SlideLayoutParts() = %d, want 1",
			len(layouts),
		)
	}

	// Add another layout
	layout2, err := master.AddSlideLayoutPart()
	if err != nil {
		t.Fatalf(
			"AddSlideLayoutPart() second error = %v",
			err,
		)
	}
	if layout2 == nil {
		t.Error(
			"AddSlideLayoutPart() second returned nil",
		)
	}

	// Now should have 2 layouts
	layouts = master.SlideLayoutParts()
	if len(layouts) != 2 {
		t.Errorf(
			"SlideLayoutParts() = %d, want 2",
			len(layouts),
		)
	}
}

// TestImagePart_ImageType tests image type detection.
func TestImagePart_ImageType(t *testing.T) {
	tests := []struct {
		imageType ImageType
		wantStr   string
		wantExt   string
		wantCT    string
	}{
		{
			ImageTypePng,
			"png",
			".png",
			"image/png",
		},
		{
			ImageTypeJpeg,
			"jpeg",
			".jpeg",
			"image/jpeg",
		},
		{
			ImageTypeGif,
			"gif",
			".gif",
			"image/gif",
		},
		{
			ImageTypeBmp,
			"bmp",
			".bmp",
			"image/bmp",
		},
		{
			ImageTypeTiff,
			"tiff",
			".tiff",
			"image/tiff",
		},
		{
			ImageTypeEmf,
			"emf",
			".emf",
			"image/x-emf",
		},
		{
			ImageTypeWmf,
			"wmf",
			".wmf",
			"image/x-wmf",
		},
		{
			ImageTypeIcon,
			"ico",
			".ico",
			"image/x-icon",
		},
		{
			ImageTypeSvg,
			"svg",
			".svg",
			"image/svg+xml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.wantStr, func(t *testing.T) {
			if got := tt.imageType.String(); got != tt.wantStr {
				t.Errorf(
					"String() = %v, want %v",
					got,
					tt.wantStr,
				)
			}
			if got := tt.imageType.Extension(); got != tt.wantExt {
				t.Errorf(
					"Extension() = %v, want %v",
					got,
					tt.wantExt,
				)
			}
			if got := tt.imageType.ContentType(); got != tt.wantCT {
				t.Errorf(
					"ContentType() = %v, want %v",
					got,
					tt.wantCT,
				)
			}
		})
	}

	// Test unknown type defaults
	unknown := ImageType(999)
	if unknown.String() != "png" {
		t.Errorf(
			"Unknown.String() = %v, want png",
			unknown.String(),
		)
	}
	if unknown.Extension() != ".png" {
		t.Errorf(
			"Unknown.Extension() = %v, want .png",
			unknown.Extension(),
		)
	}
	if unknown.ContentType() != "image/png" {
		t.Errorf(
			"Unknown.ContentType() = %v, want image/png",
			unknown.ContentType(),
		)
	}
}

// TestImageTypeFromExtension tests extension to image type mapping.
func TestImageTypeFromExtension(t *testing.T) {
	tests := []struct {
		ext  string
		want ImageType
	}{
		{".png", ImageTypePng},
		{".PNG", ImageTypePng},
		{"png", ImageTypePng},
		{".jpg", ImageTypeJpeg},
		{".jpeg", ImageTypeJpeg},
		{".gif", ImageTypeGif},
		{".bmp", ImageTypeBmp},
		{".tif", ImageTypeTiff},
		{".tiff", ImageTypeTiff},
		{".emf", ImageTypeEmf},
		{".wmf", ImageTypeWmf},
		{".ico", ImageTypeIcon},
		{".svg", ImageTypeSvg},
		{".unknown", ImageTypePng}, // Default
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			if got := ImageTypeFromExtension(tt.ext); got != tt.want {
				t.Errorf(
					"ImageTypeFromExtension(%v) = %v, want %v",
					tt.ext,
					got,
					tt.want,
				)
			}
		})
	}
}

// TestImageTypeFromFilename tests filename to image type mapping.
func TestImageTypeFromFilename(t *testing.T) {
	tests := []struct {
		filename string
		want     ImageType
	}{
		{"image.png", ImageTypePng},
		{"photo.jpg", ImageTypeJpeg},
		{"icon.gif", ImageTypeGif},
		{"diagram.svg", ImageTypeSvg},
		{"/path/to/image.PNG", ImageTypePng},
		{"file.unknown", ImageTypePng}, // Default
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if got := ImageTypeFromFilename(tt.filename); got != tt.want {
				t.Errorf(
					"ImageTypeFromFilename(%v) = %v, want %v",
					tt.filename,
					got,
					tt.want,
				)
			}
		})
	}
}

// TestImageTypeFromMagicBytes tests magic byte detection.
func TestImageTypeFromMagicBytes(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		want     ImageType
		detected bool
	}{
		{
			name: "PNG",
			data: []byte{
				0x89,
				0x50,
				0x4E,
				0x47,
				0x0D,
				0x0A,
				0x1A,
				0x0A,
			},
			want:     ImageTypePng,
			detected: true,
		},
		{
			name: "JPEG",
			data: []byte{
				0xFF,
				0xD8,
				0xFF,
				0xE0,
				0x00,
				0x10,
				0x4A,
				0x46,
			},
			want:     ImageTypeJpeg,
			detected: true,
		},
		{
			name: "GIF",
			data: []byte{
				0x47,
				0x49,
				0x46,
				0x38,
				0x39,
				0x61,
				0x00,
				0x00,
			},
			want:     ImageTypeGif,
			detected: true,
		},
		{
			name: "BMP",
			data: []byte{
				0x42,
				0x4D,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			want:     ImageTypeBmp,
			detected: true,
		},
		{
			name: "TIFF LE",
			data: []byte{
				0x49,
				0x49,
				0x2A,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			want:     ImageTypeTiff,
			detected: true,
		},
		{
			name: "TIFF BE",
			data: []byte{
				0x4D,
				0x4D,
				0x00,
				0x2A,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			want:     ImageTypeTiff,
			detected: true,
		},
		{
			name: "ICO",
			data: []byte{
				0x00,
				0x00,
				0x01,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			want:     ImageTypeIcon,
			detected: true,
		},
		{
			name: "SVG XML",
			data: []byte(
				"<?xml version=\"1.0\"?>",
			),
			want:     ImageTypeSvg,
			detected: true,
		},
		{
			name: "SVG tag",
			data: []byte(
				"<svg xmlns=\"http://www.w3.org/2000/svg\">",
			),
			want:     ImageTypeSvg,
			detected: true,
		},
		{
			name: "Unknown",
			data: []byte{
				0x00,
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
			},
			want:     ImageTypePng,
			detected: false,
		},
		{
			name:     "Too short",
			data:     []byte{0x89, 0x50},
			want:     ImageTypePng,
			detected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, detected := ImageTypeFromMagicBytes(
				tt.data,
			)
			if detected != tt.detected {
				t.Errorf(
					"detected = %v, want %v",
					detected,
					tt.detected,
				)
			}
			if got != tt.want {
				t.Errorf(
					"ImageTypeFromMagicBytes() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

// TestDetectImageType tests combined detection.
func TestDetectImageType(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		filename string
		want     ImageType
	}{
		{
			name: "Magic bytes PNG",
			data: []byte{
				0x89,
				0x50,
				0x4E,
				0x47,
				0x0D,
				0x0A,
				0x1A,
				0x0A,
			},
			filename: "image.jpg", // Filename says JPEG but magic says PNG
			want:     ImageTypePng,
		},
		{
			name: "Fallback to extension",
			data: []byte{
				0x00,
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
			},
			filename: "image.gif",
			want:     ImageTypeGif,
		},
		{
			name: "Default when unknown",
			data: []byte{
				0x00,
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
			},
			filename: "",
			want:     ImageTypePng,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectImageType(tt.data, tt.filename); got != tt.want {
				t.Errorf(
					"DetectImageType() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

// TestSlidePart_AddImagePart tests adding images to slides.
func TestSlidePart_AddImagePart(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Add a slide first
	slide, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	// Initially no images
	images := slide.ImageParts()
	if len(images) != 0 {
		t.Errorf(
			"ImageParts() = %d, want 0",
			len(images),
		)
	}

	// Add a PNG image
	imgPart, err := slide.AddImagePart(
		ImageTypePng,
	)
	if err != nil {
		t.Fatalf("AddImagePart() error = %v", err)
	}
	if imgPart == nil {
		t.Error("AddImagePart() returned nil")
	}

	// Verify image type
	if imgPart.ImageType() != ImageTypePng {
		t.Errorf(
			"ImageType() = %v, want %v",
			imgPart.ImageType(),
			ImageTypePng,
		)
	}

	// Now should have 1 image
	images = slide.ImageParts()
	if len(images) != 1 {
		t.Errorf(
			"ImageParts() = %d, want 1",
			len(images),
		)
	}

	// Add JPEG image
	imgPart2, err := slide.AddImagePart(
		ImageTypeJpeg,
	)
	if err != nil {
		t.Fatalf(
			"AddImagePart() JPEG error = %v",
			err,
		)
	}
	if imgPart2.ImageType() != ImageTypeJpeg {
		t.Errorf(
			"ImageType() = %v, want %v",
			imgPart2.ImageType(),
			ImageTypeJpeg,
		)
	}

	// Now should have 2 images
	images = slide.ImageParts()
	if len(images) != 2 {
		t.Errorf(
			"ImageParts() = %d, want 2",
			len(images),
		)
	}
}

// TestSlidePart_SlideCommentsPart tests adding comments to slides.
func TestSlidePart_SlideCommentsPart(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Add a slide first
	slide, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	// Initially no comments
	if slide.SlideCommentsPart() != nil {
		t.Error(
			"SlideCommentsPart() should be nil initially",
		)
	}

	// Add comments
	comments, err := slide.AddSlideCommentsPart()
	if err != nil {
		t.Fatalf(
			"AddSlideCommentsPart() error = %v",
			err,
		)
	}
	if comments == nil {
		t.Error(
			"AddSlideCommentsPart() returned nil",
		)
	}

	// Now should have comments
	if slide.SlideCommentsPart() == nil {
		t.Error(
			"SlideCommentsPart() should not be nil after adding",
		)
	}
}

// TestPresentationPart_AddNotesMasterPart tests adding notes master.
func TestPresentationPart_AddNotesMasterPart(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Initially no notes master
	if presPart.NotesMasterPart() != nil {
		t.Error(
			"NotesMasterPart() should be nil initially",
		)
	}

	// Add notes master
	notesMaster, err := presPart.AddNotesMasterPart()
	if err != nil {
		t.Fatalf(
			"AddNotesMasterPart() error = %v",
			err,
		)
	}
	if notesMaster == nil {
		t.Error(
			"AddNotesMasterPart() returned nil",
		)
	}

	// Verify content type
	if notesMaster.FixedContentType() != ContentTypeNotesMaster {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			notesMaster.FixedContentType(),
			ContentTypeNotesMaster,
		)
	}

	// Now should have notes master
	if presPart.NotesMasterPart() == nil {
		t.Error(
			"NotesMasterPart() should not be nil after adding",
		)
	}
}

// TestPresentationPart_AddHandoutMasterPart tests adding handout master.
func TestPresentationPart_AddHandoutMasterPart(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Initially no handout master
	if presPart.HandoutMasterPart() != nil {
		t.Error(
			"HandoutMasterPart() should be nil initially",
		)
	}

	// Add handout master
	handoutMaster, err := presPart.AddHandoutMasterPart()
	if err != nil {
		t.Fatalf(
			"AddHandoutMasterPart() error = %v",
			err,
		)
	}
	if handoutMaster == nil {
		t.Error(
			"AddHandoutMasterPart() returned nil",
		)
	}

	// Verify content type
	if handoutMaster.FixedContentType() != ContentTypeHandoutMaster {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			handoutMaster.FixedContentType(),
			ContentTypeHandoutMaster,
		)
	}

	// Now should have handout master
	if presPart.HandoutMasterPart() == nil {
		t.Error(
			"HandoutMasterPart() should not be nil after adding",
		)
	}
}

// TestPresentationPart_AddThemePart tests adding theme.
func TestPresentationPart_AddThemePart(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Initially no theme
	if presPart.ThemePart() != nil {
		t.Error(
			"ThemePart() should be nil initially",
		)
	}

	// Add theme
	theme, err := presPart.AddThemePart()
	if err != nil {
		t.Fatalf("AddThemePart() error = %v", err)
	}
	if theme == nil {
		t.Error("AddThemePart() returned nil")
	}

	// Verify content type
	if theme.FixedContentType() != ContentTypeTheme {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			theme.FixedContentType(),
			ContentTypeTheme,
		)
	}

	// Now should have theme
	if presPart.ThemePart() == nil {
		t.Error(
			"ThemePart() should not be nil after adding",
		)
	}

	// Test ThemeParts
	themes := presPart.ThemeParts()
	if len(themes) != 1 {
		t.Errorf(
			"ThemeParts() = %d, want 1",
			len(themes),
		)
	}
}

// TestPresentationPart_AddCommentAuthorsPart tests adding comment authors.
func TestPresentationPart_AddCommentAuthorsPart(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Initially no comment authors
	if presPart.CommentAuthorsPart() != nil {
		t.Error(
			"CommentAuthorsPart() should be nil initially",
		)
	}

	// Add comment authors
	authors, err := presPart.AddCommentAuthorsPart()
	if err != nil {
		t.Fatalf(
			"AddCommentAuthorsPart() error = %v",
			err,
		)
	}
	if authors == nil {
		t.Error(
			"AddCommentAuthorsPart() returned nil",
		)
	}

	// Verify content type
	if authors.FixedContentType() != ContentTypeCommentAuthors {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			authors.FixedContentType(),
			ContentTypeCommentAuthors,
		)
	}

	// Now should have comment authors
	if presPart.CommentAuthorsPart() == nil {
		t.Error(
			"CommentAuthorsPart() should not be nil after adding",
		)
	}
}

// TestSlidePart_GetStream tests getting slide content stream.
func TestSlidePart_GetStream(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	slide, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	stream := slide.GetStream()
	if stream == nil {
		t.Error("GetStream() returned nil")
	}
}

// TestSlideMasterPart_AddThemePart tests adding theme to master.
func TestSlideMasterPart_AddThemePart(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	master, err := presPart.AddSlideMasterPart()
	if err != nil {
		t.Fatalf(
			"AddSlideMasterPart() error = %v",
			err,
		)
	}

	// Initially no theme
	if master.ThemePart() != nil {
		t.Error(
			"ThemePart() should be nil initially",
		)
	}

	// Add theme
	theme, err := master.AddThemePart()
	if err != nil {
		t.Fatalf("AddThemePart() error = %v", err)
	}
	if theme == nil {
		t.Error("AddThemePart() returned nil")
	}

	// Now should have theme
	if master.ThemePart() == nil {
		t.Error(
			"ThemePart() should not be nil after adding",
		)
	}
}

// TestNewPresentationPartFromData tests wrapping existing data.
func TestNewPresentationPartFromData(
	t *testing.T,
) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Get the data
	data := presPart.OpenXmlPartData

	// Wrap it
	wrapped := NewPresentationPartFromData(
		data,
		ContentTypePresentation,
	)
	if wrapped == nil {
		t.Error(
			"NewPresentationPartFromData() returned nil",
		)
	}

	if wrapped.FixedContentType() != ContentTypePresentation {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			wrapped.FixedContentType(),
			ContentTypePresentation,
		)
	}
}

// TestImagePart_FeedData tests feeding image data.
func TestImagePart_FeedData(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	slide, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	imgPart, err := slide.AddImagePart(
		ImageTypePng,
	)
	if err != nil {
		t.Fatalf("AddImagePart() error = %v", err)
	}

	// Feed some test data
	testData := []byte{
		0x89,
		0x50,
		0x4E,
		0x47,
		0x0D,
		0x0A,
		0x1A,
		0x0A,
	}
	imgPart.FeedDataBytes(testData)

	// Get the data back
	gotData, err := imgPart.GetData()
	if err != nil {
		t.Fatalf("GetData() error = %v", err)
	}

	if len(gotData) != len(testData) {
		t.Errorf(
			"GetData() length = %d, want %d",
			len(gotData),
			len(testData),
		)
	}
}
