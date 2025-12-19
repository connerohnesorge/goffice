package parts

import (
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func TestPackageExists(t *testing.T) {
	// Simple test to verify test infrastructure works for parts package
	t.Log("parts package test infrastructure is working")
}

func TestImageTypeString(t *testing.T) {
	tests := []struct {
		imageType ImageType
		expected  string
	}{
		{ImageTypePng, "png"},
		{ImageTypeJpeg, "jpeg"},
		{ImageTypeGif, "gif"},
		{ImageTypeBmp, "bmp"},
		{ImageTypeTiff, "tiff"},
		{ImageTypeEmf, "emf"},
		{ImageTypeWmf, "wmf"},
		{ImageTypeIcon, "ico"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.imageType.String(); got != tt.expected {
				t.Errorf("ImageType.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestImageTypeExtension(t *testing.T) {
	tests := []struct {
		imageType ImageType
		expected  string
	}{
		{ImageTypePng, ".png"},
		{ImageTypeJpeg, ".jpeg"},
		{ImageTypeGif, ".gif"},
		{ImageTypeBmp, ".bmp"},
		{ImageTypeTiff, ".tiff"},
		{ImageTypeEmf, ".emf"},
		{ImageTypeWmf, ".wmf"},
		{ImageTypeIcon, ".ico"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.imageType.Extension(); got != tt.expected {
				t.Errorf("ImageType.Extension() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestImageTypeContentType(t *testing.T) {
	tests := []struct {
		imageType ImageType
		expected  string
	}{
		{ImageTypePng, "image/png"},
		{ImageTypeJpeg, "image/jpeg"},
		{ImageTypeGif, "image/gif"},
		{ImageTypeBmp, "image/bmp"},
		{ImageTypeTiff, "image/tiff"},
		{ImageTypeEmf, "image/x-emf"},
		{ImageTypeWmf, "image/x-wmf"},
		{ImageTypeIcon, "image/x-icon"},
	}

	for _, tt := range tests {
		t.Run(tt.imageType.String(), func(t *testing.T) {
			if got := tt.imageType.ContentType(); got != tt.expected {
				t.Errorf("ImageType.ContentType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestImageTypeFromContentType(t *testing.T) {
	tests := []struct {
		contentType string
		expected    ImageType
	}{
		{"image/png", ImageTypePng},
		{"image/jpeg", ImageTypeJpeg},
		{"image/jpg", ImageTypeJpeg},
		{"image/gif", ImageTypeGif},
		{"image/bmp", ImageTypeBmp},
		{"image/tiff", ImageTypeTiff},
		{"image/x-emf", ImageTypeEmf},
		{"image/x-wmf", ImageTypeWmf},
		{"image/x-icon", ImageTypeIcon},
		{"unknown/type", ImageTypePng}, // default
	}

	for _, tt := range tests {
		t.Run(tt.contentType, func(t *testing.T) {
			if got := imageTypeFromContentType(tt.contentType); got != tt.expected {
				t.Errorf("imageTypeFromContentType(%q) = %v, want %v", tt.contentType, got, tt.expected)
			}
		})
	}
}

func TestImageTypeFromExtension(t *testing.T) {
	tests := []struct {
		extension string
		expected  ImageType
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
		{".unknown", ImageTypePng}, // Default to PNG
	}

	for _, tt := range tests {
		t.Run(tt.extension, func(t *testing.T) {
			if got := ImageTypeFromExtension(tt.extension); got != tt.expected {
				t.Errorf("ImageTypeFromExtension(%s) = %v, want %v", tt.extension, got, tt.expected)
			}
		})
	}
}

func TestImageTypeFromFilename(t *testing.T) {
	tests := []struct {
		filename string
		expected ImageType
	}{
		{"image.png", ImageTypePng},
		{"photo.jpg", ImageTypeJpeg},
		{"animation.gif", ImageTypeGif},
		{"/path/to/file.bmp", ImageTypeBmp},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if got := ImageTypeFromFilename(tt.filename); got != tt.expected {
				t.Errorf("ImageTypeFromFilename(%s) = %v, want %v", tt.filename, got, tt.expected)
			}
		})
	}
}

func TestImageTypeFromMagicBytes(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected ImageType
		detected bool
	}{
		{
			name:     "PNG",
			data:     []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A},
			expected: ImageTypePng,
			detected: true,
		},
		{
			name:     "JPEG",
			data:     []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46},
			expected: ImageTypeJpeg,
			detected: true,
		},
		{
			name:     "GIF",
			data:     []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x00, 0x00},
			expected: ImageTypeGif,
			detected: true,
		},
		{
			name:     "BMP",
			data:     []byte{0x42, 0x4D, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expected: ImageTypeBmp,
			detected: true,
		},
		{
			name:     "TIFF LE",
			data:     []byte{0x49, 0x49, 0x2A, 0x00, 0x00, 0x00, 0x00, 0x00},
			expected: ImageTypeTiff,
			detected: true,
		},
		{
			name:     "TIFF BE",
			data:     []byte{0x4D, 0x4D, 0x00, 0x2A, 0x00, 0x00, 0x00, 0x00},
			expected: ImageTypeTiff,
			detected: true,
		},
		{
			name:     "ICO",
			data:     []byte{0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00},
			expected: ImageTypeIcon,
			detected: true,
		},
		{
			name:     "WMF placeable",
			data:     []byte{0xD7, 0xCD, 0xC6, 0x9A, 0x00, 0x00, 0x00, 0x00},
			expected: ImageTypeWmf,
			detected: true,
		},
		{
			name:     "Unknown",
			data:     []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
			expected: ImageTypePng,
			detected: false,
		},
		{
			name:     "Too short",
			data:     []byte{0x89, 0x50, 0x4E},
			expected: ImageTypePng,
			detected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, detected := ImageTypeFromMagicBytes(tt.data)
			if detected != tt.detected {
				t.Errorf("ImageTypeFromMagicBytes() detected = %v, want %v", detected, tt.detected)
			}
			if got != tt.expected {
				t.Errorf("ImageTypeFromMagicBytes() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDetectImageType(t *testing.T) {
	// Test with valid PNG magic bytes
	pngData := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	got := DetectImageType(pngData, "image.jpg") // filename should be ignored
	if got != ImageTypePng {
		t.Errorf("DetectImageType with PNG magic bytes = %v, want %v", got, ImageTypePng)
	}

	// Test with unknown magic bytes, fallback to extension
	unknownData := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}
	got = DetectImageType(unknownData, "image.gif")
	if got != ImageTypeGif {
		t.Errorf("DetectImageType with fallback to extension = %v, want %v", got, ImageTypeGif)
	}

	// Test with unknown magic bytes and no filename
	got = DetectImageType(unknownData, "")
	if got != ImageTypePng {
		t.Errorf("DetectImageType default = %v, want %v", got, ImageTypePng)
	}
}

func TestContentTypeConstants(t *testing.T) {
	// Verify content type constants are correct
	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"ContentTypeDocument", ContentTypeDocument, "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"},
		{"ContentTypeTemplate", ContentTypeTemplate, "application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml"},
		{"ContentTypeMacroEnabled", ContentTypeMacroEnabled, "application/vnd.ms-word.document.macroEnabled.main+xml"},
		{"ContentTypeMacroTemplate", ContentTypeMacroTemplate, "application/vnd.ms-word.template.macroEnabled.main+xml"},
		{"ContentTypeStyles", ContentTypeStyles, "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"},
		{"ContentTypeNumbering", ContentTypeNumbering, "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml"},
		{"ContentTypeSettings", ContentTypeSettings, "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml"},
		{"ContentTypeFontTable", ContentTypeFontTable, "application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml"},
		{"ContentTypeWebSettings", ContentTypeWebSettings, "application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml"},
		{"ContentTypeHeader", ContentTypeHeader, "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml"},
		{"ContentTypeFooter", ContentTypeFooter, "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml"},
		{"ContentTypeFootnotes", ContentTypeFootnotes, "application/vnd.openxmlformats-officedocument.wordprocessingml.footnotes+xml"},
		{"ContentTypeEndnotes", ContentTypeEndnotes, "application/vnd.openxmlformats-officedocument.wordprocessingml.endnotes+xml"},
		{"ContentTypeComments", ContentTypeComments, "application/vnd.openxmlformats-officedocument.wordprocessingml.comments+xml"},
		{"ContentTypeTheme", ContentTypeTheme, "application/vnd.openxmlformats-officedocument.theme+xml"},
		{"ContentTypeVbaProject", ContentTypeVbaProject, "application/vnd.ms-office.vbaProject"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("%s = %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestRelationshipTypeConstants(t *testing.T) {
	// Verify relationship type constants are correct
	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"RelationshipTypeStyles", RelationshipTypeStyles, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"},
		{"RelationshipTypeNumbering", RelationshipTypeNumbering, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/numbering"},
		{"RelationshipTypeSettings", RelationshipTypeSettings, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings"},
		{"RelationshipTypeFontTable", RelationshipTypeFontTable, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable"},
		{"RelationshipTypeWebSettings", RelationshipTypeWebSettings, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/webSettings"},
		{"RelationshipTypeHeader", RelationshipTypeHeader, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/header"},
		{"RelationshipTypeFooter", RelationshipTypeFooter, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footer"},
		{"RelationshipTypeFootnotes", RelationshipTypeFootnotes, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footnotes"},
		{"RelationshipTypeEndnotes", RelationshipTypeEndnotes, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/endnotes"},
		{"RelationshipTypeComments", RelationshipTypeComments, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"},
		{"RelationshipTypeTheme", RelationshipTypeTheme, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme"},
		{"RelationshipTypeImage", RelationshipTypeImage, "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"},
		{"RelationshipTypeVbaProject", RelationshipTypeVbaProject, "http://schemas.microsoft.com/office/2006/relationships/vbaProject"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("%s = %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestPartErrors(t *testing.T) {
	// Test that error types have proper message
	if ErrNilPackage.Error() != "package is nil" {
		t.Errorf("ErrNilPackage.Error() = %v, want 'package is nil'", ErrNilPackage.Error())
	}
}

func TestImagePartImageType(t *testing.T) {
	// Test that ImagePart correctly tracks its image type
	ip := &ImagePart{
		imageType: ImageTypeJpeg,
	}

	if ip.ImageType() != ImageTypeJpeg {
		t.Errorf("ImageType() = %v, want %v", ip.ImageType(), ImageTypeJpeg)
	}
}

func TestMainPartFixedContentType(t *testing.T) {
	mp := &MainPart{
		contentType: ContentTypeDocument,
	}

	if mp.FixedContentType() != ContentTypeDocument {
		t.Errorf("FixedContentType() = %v, want %v", mp.FixedContentType(), ContentTypeDocument)
	}
}

func TestStylesPartFixedContentType(t *testing.T) {
	sp := &StylesPart{}

	if sp.FixedContentType() != ContentTypeStyles {
		t.Errorf("FixedContentType() = %v, want %v", sp.FixedContentType(), ContentTypeStyles)
	}
}

func TestNumberingPartFixedContentType(t *testing.T) {
	np := &NumberingPart{}

	if np.FixedContentType() != ContentTypeNumbering {
		t.Errorf("FixedContentType() = %v, want %v", np.FixedContentType(), ContentTypeNumbering)
	}
}

func TestSettingsPartFixedContentType(t *testing.T) {
	sp := &SettingsPart{}

	if sp.FixedContentType() != ContentTypeSettings {
		t.Errorf("FixedContentType() = %v, want %v", sp.FixedContentType(), ContentTypeSettings)
	}
}

func TestWebSettingsPartFixedContentType(t *testing.T) {
	wsp := &WebSettingsPart{}

	if wsp.FixedContentType() != ContentTypeWebSettings {
		t.Errorf("FixedContentType() = %v, want %v", wsp.FixedContentType(), ContentTypeWebSettings)
	}
}

func TestFontsPartFixedContentType(t *testing.T) {
	fp := &FontsPart{}

	if fp.FixedContentType() != ContentTypeFontTable {
		t.Errorf("FixedContentType() = %v, want %v", fp.FixedContentType(), ContentTypeFontTable)
	}
}

func TestHeaderPartFixedContentType(t *testing.T) {
	hp := &HeaderPart{}

	if hp.FixedContentType() != ContentTypeHeader {
		t.Errorf("FixedContentType() = %v, want %v", hp.FixedContentType(), ContentTypeHeader)
	}
}

func TestFooterPartFixedContentType(t *testing.T) {
	fp := &FooterPart{}

	if fp.FixedContentType() != ContentTypeFooter {
		t.Errorf("FixedContentType() = %v, want %v", fp.FixedContentType(), ContentTypeFooter)
	}
}

func TestFootnotesPartFixedContentType(t *testing.T) {
	fp := &FootnotesPart{}

	if fp.FixedContentType() != ContentTypeFootnotes {
		t.Errorf("FixedContentType() = %v, want %v", fp.FixedContentType(), ContentTypeFootnotes)
	}
}

func TestEndnotesPartFixedContentType(t *testing.T) {
	ep := &EndnotesPart{}

	if ep.FixedContentType() != ContentTypeEndnotes {
		t.Errorf("FixedContentType() = %v, want %v", ep.FixedContentType(), ContentTypeEndnotes)
	}
}

func TestThemePartFixedContentType(t *testing.T) {
	tp := &ThemePart{}

	if tp.FixedContentType() != ContentTypeTheme {
		t.Errorf("FixedContentType() = %v, want %v", tp.FixedContentType(), ContentTypeTheme)
	}
}

func TestVbaProjectPartFixedContentType(t *testing.T) {
	vp := &VbaProjectPart{}

	if vp.FixedContentType() != ContentTypeVbaProject {
		t.Errorf("FixedContentType() = %v, want %v", vp.FixedContentType(), ContentTypeVbaProject)
	}
}

func TestCommentsPartAddComment(t *testing.T) {
	// Create a minimal test setup with OpenXmlPartData
	cp := &CommentsPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData("/word/comments.xml", ContentTypeComments, nil, nil),
	}

	// Set root element to a new Comments element
	comments := elements.NewComments()
	cp.SetRootElement(comments)

	// Add a comment using the Comments element
	c := comments.AddComment("Test Author", "Test comment")
	if c.Id() != 1 {
		t.Errorf("AddComment() returned id = %v, want 1", c.Id())
	}

	// Add another comment
	c2 := comments.AddComment("Author 2", "Second comment")
	if c2.Id() != 2 {
		t.Errorf("AddComment() returned id = %v, want 2", c2.Id())
	}
}
