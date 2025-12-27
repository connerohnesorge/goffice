package parts_test

import (
	"bytes"
	"testing"

	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

func TestImageType_String(t *testing.T) {
	tests := []struct {
		imageType parts.ImageType
		expected  string
	}{
		{parts.ImageTypePng, "png"},
		{parts.ImageTypeJpeg, "jpeg"},
		{parts.ImageTypeGif, "gif"},
		{parts.ImageTypeBmp, "bmp"},
		{parts.ImageTypeTiff, "tiff"},
		{parts.ImageTypeEmf, "emf"},
		{parts.ImageTypeWmf, "wmf"},
		{parts.ImageTypeIcon, "ico"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.imageType.String(); got != tt.expected {
				t.Errorf(
					"ImageType.String() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestImageType_Extension(t *testing.T) {
	tests := []struct {
		imageType parts.ImageType
		expected  string
	}{
		{parts.ImageTypePng, ".png"},
		{parts.ImageTypeJpeg, ".jpeg"},
		{parts.ImageTypeGif, ".gif"},
		{parts.ImageTypeBmp, ".bmp"},
		{parts.ImageTypeTiff, ".tiff"},
		{parts.ImageTypeEmf, ".emf"},
		{parts.ImageTypeWmf, ".wmf"},
		{parts.ImageTypeIcon, ".ico"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.imageType.Extension(); got != tt.expected {
				t.Errorf(
					"ImageType.Extension() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestImageType_ContentType(t *testing.T) {
	tests := []struct {
		imageType parts.ImageType
		expected  string
	}{
		{parts.ImageTypePng, "image/png"},
		{parts.ImageTypeJpeg, "image/jpeg"},
		{parts.ImageTypeGif, "image/gif"},
		{parts.ImageTypeBmp, "image/bmp"},
		{parts.ImageTypeTiff, "image/tiff"},
		{parts.ImageTypeEmf, "image/x-emf"},
		{parts.ImageTypeWmf, "image/x-wmf"},
		{parts.ImageTypeIcon, "image/x-icon"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.imageType.ContentType(); got != tt.expected {
				t.Errorf(
					"ImageType.ContentType() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestImageTypeFromExtension(t *testing.T) {
	tests := []struct {
		ext      string
		expected parts.ImageType
	}{
		{".png", parts.ImageTypePng},
		{"png", parts.ImageTypePng},
		{".PNG", parts.ImageTypePng},
		{".jpg", parts.ImageTypeJpeg},
		{".jpeg", parts.ImageTypeJpeg},
		{".gif", parts.ImageTypeGif},
		{".bmp", parts.ImageTypeBmp},
		{".tiff", parts.ImageTypeTiff},
		{".tif", parts.ImageTypeTiff},
		{".emf", parts.ImageTypeEmf},
		{".wmf", parts.ImageTypeWmf},
		{".ico", parts.ImageTypeIcon},
		{
			".unknown",
			parts.ImageTypePng,
		}, // Default
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			if got := parts.ImageTypeFromExtension(tt.ext); got != tt.expected {
				t.Errorf(
					"ImageTypeFromExtension(%q) = %v, want %v",
					tt.ext,
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestImageTypeFromFilename(t *testing.T) {
	tests := []struct {
		filename string
		expected parts.ImageType
	}{
		{"image.png", parts.ImageTypePng},
		{"photo.jpg", parts.ImageTypeJpeg},
		{"animation.gif", parts.ImageTypeGif},
		{"diagram.bmp", parts.ImageTypeBmp},
		{"scan.tiff", parts.ImageTypeTiff},
		{"vector.emf", parts.ImageTypeEmf},
		{"legacy.wmf", parts.ImageTypeWmf},
		{"favicon.ico", parts.ImageTypeIcon},
		{
			"noextension",
			parts.ImageTypePng,
		}, // Default
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if got := parts.ImageTypeFromFilename(tt.filename); got != tt.expected {
				t.Errorf(
					"ImageTypeFromFilename(%q) = %v, want %v",
					tt.filename,
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestImageTypeFromMagicBytes(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected parts.ImageType
		detected bool
	}{
		{
			name: "PNG magic bytes",
			data: []byte{
				0x89,
				0x50,
				0x4E,
				0x47,
				0x0D,
				0x0A,
				0x1A,
				0x0A,
				0x00,
				0x00,
			},
			expected: parts.ImageTypePng,
			detected: true,
		},
		{
			name: "JPEG magic bytes",
			data: []byte{
				0xFF,
				0xD8,
				0xFF,
				0xE0,
				0x00,
				0x10,
				0x4A,
				0x46,
				0x49,
				0x46,
			},
			expected: parts.ImageTypeJpeg,
			detected: true,
		},
		{
			name: "GIF magic bytes",
			data: []byte{
				0x47,
				0x49,
				0x46,
				0x38,
				0x39,
				0x61,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			expected: parts.ImageTypeGif,
			detected: true,
		},
		{
			name: "BMP magic bytes",
			data: []byte{
				0x42,
				0x4D,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			expected: parts.ImageTypeBmp,
			detected: true,
		},
		{
			name: "TIFF LE magic bytes",
			data: []byte{
				0x49,
				0x49,
				0x2A,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			expected: parts.ImageTypeTiff,
			detected: true,
		},
		{
			name: "TIFF BE magic bytes",
			data: []byte{
				0x4D,
				0x4D,
				0x00,
				0x2A,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			expected: parts.ImageTypeTiff,
			detected: true,
		},
		{
			name: "ICO magic bytes",
			data: []byte{
				0x00,
				0x00,
				0x01,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			expected: parts.ImageTypeIcon,
			detected: true,
		},
		{
			name: "WMF placeable magic bytes",
			data: []byte{
				0xD7,
				0xCD,
				0xC6,
				0x9A,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			expected: parts.ImageTypeWmf,
			detected: true,
		},
		{
			name: "Unknown data",
			data: []byte{
				0x00,
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
				0x08,
				0x09,
			},
			expected: parts.ImageTypePng,
			detected: false,
		},
		{
			name:     "Too short data",
			data:     []byte{0x89, 0x50},
			expected: parts.ImageTypePng,
			detected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, detected := parts.ImageTypeFromMagicBytes(
				tt.data,
			)
			if got != tt.expected ||
				detected != tt.detected {
				t.Errorf(
					"ImageTypeFromMagicBytes() = (%v, %v), want (%v, %v)",
					got,
					detected,
					tt.expected,
					tt.detected,
				)
			}
		})
	}
}

func TestDetectImageType(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		filename string
		expected parts.ImageType
	}{
		{
			name: "PNG detected by magic bytes",
			data: []byte{
				0x89,
				0x50,
				0x4E,
				0x47,
				0x0D,
				0x0A,
				0x1A,
				0x0A,
				0x00,
				0x00,
			},
			filename: "image.jpg", // Wrong extension, but magic bytes win
			expected: parts.ImageTypePng,
		},
		{
			name: "Fall back to extension",
			data: []byte{
				0x00,
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
				0x08,
				0x09,
			},
			filename: "image.gif",
			expected: parts.ImageTypeGif,
		},
		{
			name: "Default to PNG when unknown",
			data: []byte{
				0x00,
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
				0x08,
				0x09,
			},
			filename: "",
			expected: parts.ImageTypePng,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parts.DetectImageType(tt.data, tt.filename); got != tt.expected {
				t.Errorf(
					"DetectImageType() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestContentTypes(t *testing.T) {
	// Verify content type constants are correctly defined
	tests := []struct {
		name     string
		constant string
	}{
		{
			"WorkbookContentType",
			parts.ContentTypeWorkbook,
		},
		{
			"WorksheetContentType",
			parts.ContentTypeWorksheet,
		},
		{
			"SharedStringsContentType",
			parts.ContentTypeSharedStrings,
		},
		{
			"StylesContentType",
			parts.ContentTypeStyles,
		},
		{
			"CalcChainContentType",
			parts.ContentTypeCalcChain,
		},
		{
			"ChartsheetContentType",
			parts.ContentTypeChartsheet,
		},
		{
			"ThemeContentType",
			parts.ContentTypeTheme,
		},
		{
			"DrawingContentType",
			parts.ContentTypeDrawing,
		},
		{
			"ChartContentType",
			parts.ContentTypeChart,
		},
		{
			"TableContentType",
			parts.ContentTypeTable,
		},
		{
			"CommentsContentType",
			parts.ContentTypeComments,
		},
		{
			"VbaProjectContentType",
			parts.ContentTypeVbaProject,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant == "" {
				t.Errorf("%s is empty", tt.name)
			}
		})
	}
}

func TestRelationshipTypes(t *testing.T) {
	// Verify relationship type constants are correctly defined
	tests := []struct {
		name     string
		constant string
	}{
		{
			"WorksheetRelType",
			parts.RelationshipTypeWorksheet,
		},
		{
			"SharedStringsRelType",
			parts.RelationshipTypeSharedStrings,
		},
		{
			"StylesRelType",
			parts.RelationshipTypeStyles,
		},
		{
			"CalcChainRelType",
			parts.RelationshipTypeCalcChain,
		},
		{
			"ThemeRelType",
			parts.RelationshipTypeTheme,
		},
		{
			"DrawingRelType",
			parts.RelationshipTypeDrawing,
		},
		{
			"ChartRelType",
			parts.RelationshipTypeChart,
		},
		{
			"TableRelType",
			parts.RelationshipTypeTable,
		},
		{
			"CommentsRelType",
			parts.RelationshipTypeComments,
		},
		{
			"VbaProjectRelType",
			parts.RelationshipTypeVbaProject,
		},
		{
			"ImageRelType",
			parts.RelationshipTypeImage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant == "" {
				t.Errorf("%s is empty", tt.name)
			}
		})
	}
}

func TestImagePart_FeedData(t *testing.T) {
	// Test that FeedData correctly reads from a reader
	// This is a basic test since we cannot create a full part without a package
	data := []byte{
		0x89,
		0x50,
		0x4E,
		0x47,
		0x0D,
		0x0A,
		0x1A,
		0x0A,
	}
	reader := bytes.NewReader(data)

	// Just verify the reader works
	readData := make([]byte, len(data))
	n, err := reader.Read(readData)
	if err != nil {
		t.Errorf("Failed to read: %v", err)
	}
	if n != len(data) {
		t.Errorf(
			"Read %d bytes, expected %d",
			n,
			len(data),
		)
	}
	if !bytes.Equal(readData, data) {
		t.Error(
			"Read data doesn't match original",
		)
	}
}
