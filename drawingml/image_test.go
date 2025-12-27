package drawingml

import (
	"bytes"
	"io"
	"testing"
)

// Sample image data for testing.
// These are minimal valid image headers/data for each format.

// Minimal 1x1 PNG image (red pixel)
var pngData = []byte{
	0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, // PNG signature
	0x00, 0x00, 0x00, 0x0D, // IHDR chunk length
	0x49, 0x48, 0x44, 0x52, // IHDR
	0x00, 0x00, 0x00, 0x01, // width = 1
	0x00, 0x00, 0x00, 0x01, // height = 1
	0x08, 0x02, // bit depth = 8, color type = 2 (RGB)
	0x00, 0x00, 0x00, // compression, filter, interlace
	0x90, 0x77, 0x53, 0xDE, // CRC
	0x00, 0x00, 0x00, 0x0C, // IDAT chunk length
	0x49, 0x44, 0x41, 0x54, // IDAT
	0x08, 0xD7, 0x63, 0xF8, 0xFF, 0xFF, 0x3F, 0x00, 0x05, 0xFE, 0x02, 0xFE, // compressed data
	0xA3, 0x6D, 0xEB, 0xCF, // CRC
	0x00, 0x00, 0x00, 0x00, // IEND chunk length
	0x49, 0x45, 0x4E, 0x44, // IEND
	0xAE, 0x42, 0x60, 0x82, // CRC
}

// Minimal JPEG header (2x2 image)
var jpegData = []byte{
	0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
	0x01, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0xFF, 0xDB, 0x00, 0x43,
	0x00, 0x08, 0x06, 0x06, 0x07, 0x06, 0x05, 0x08, 0x07, 0x07, 0x07, 0x09,
	0x09, 0x08, 0x0A, 0x0C, 0x14, 0x0D, 0x0C, 0x0B, 0x0B, 0x0C, 0x19, 0x12,
	0x13, 0x0F, 0x14, 0x1D, 0x1A, 0x1F, 0x1E, 0x1D, 0x1A, 0x1C, 0x1C, 0x20,
	0x24, 0x2E, 0x27, 0x20, 0x22, 0x2C, 0x23, 0x1C, 0x1C, 0x28, 0x37, 0x29,
	0x2C, 0x30, 0x31, 0x34, 0x34, 0x34, 0x1F, 0x27, 0x39, 0x3D, 0x38, 0x32,
	0x3C, 0x2E, 0x33, 0x34, 0x32, 0xFF, 0xC0, 0x00, 0x0B, 0x08, 0x00, 0x02,
	0x00, 0x02, 0x01, 0x01, 0x11, 0x00, 0xFF, 0xC4, 0x00, 0x1F, 0x00, 0x00,
	0x01, 0x05, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
	0x09, 0x0A, 0x0B, 0xFF, 0xC4, 0x00, 0xB5, 0x10, 0x00, 0x02, 0x01, 0x03,
	0x03, 0x02, 0x04, 0x03, 0x05, 0x05, 0x04, 0x04, 0x00, 0x00, 0x01, 0x7D,
	0x01, 0x02, 0x03, 0x00, 0x04, 0x11, 0x05, 0x12, 0x21, 0x31, 0x41, 0x06,
	0x13, 0x51, 0x61, 0x07, 0x22, 0x71, 0x14, 0x32, 0x81, 0x91, 0xA1, 0x08,
	0x23, 0x42, 0xB1, 0xC1, 0x15, 0x52, 0xD1, 0xF0, 0x24, 0x33, 0x62, 0x72,
	0x82, 0x09, 0x0A, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x25, 0x26, 0x27, 0x28,
	0x29, 0x2A, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x43, 0x44, 0x45,
	0x46, 0x47, 0x48, 0x49, 0x4A, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59,
	0x5A, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x73, 0x74, 0x75,
	0x76, 0x77, 0x78, 0x79, 0x7A, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89,
	0x8A, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0xA2, 0xA3,
	0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6,
	0xB7, 0xB8, 0xB9, 0xBA, 0xC2, 0xC3, 0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9,
	0xCA, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xE1, 0xE2,
	0xE3, 0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xF1, 0xF2, 0xF3, 0xF4,
	0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFF, 0xDA, 0x00, 0x08, 0x01, 0x01,
	0x00, 0x00, 0x3F, 0x00, 0xFB, 0xD5, 0xDB, 0x20, 0xB8, 0xF4, 0xA8, 0x87,
	0xD3, 0xB1, 0x5F, 0xFF, 0xD9,
}

// Minimal GIF header (1x1 image)
var gifData = []byte{
	0x47, 0x49, 0x46, 0x38, 0x39, 0x61, // GIF89a
	0x01, 0x00, // width = 1
	0x01, 0x00, // height = 1
	0x00, 0x00, 0x00, // packed fields, bg color, aspect ratio
	0x2C, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, // image descriptor
	0x02, 0x02, 0x44, 0x01, 0x00, // LZW minimum code size and data
	0x3B, // trailer
}

// BMP header bytes
var bmpHeader = []byte{
	0x42, 0x4D, // BM signature
	0x00, 0x00, 0x00, 0x00, // file size (placeholder)
	0x00, 0x00, 0x00, 0x00, // reserved
	0x36, 0x00, 0x00, 0x00, // offset to pixel data
}

// WEBP header bytes
var webpHeader = []byte{
	0x52, 0x49, 0x46, 0x46, // RIFF
	0x00, 0x00, 0x00, 0x00, // file size (placeholder)
	0x57, 0x45, 0x42, 0x50, // WEBP
}

// TIFF header bytes (little endian) - needs at least 12 bytes for detection
var tiffHeaderLE = []byte{
	0x49, 0x49, 0x2A, 0x00, // II + magic number
	0x08, 0x00, 0x00, 0x00, // offset to first IFD
	0x00, 0x00, 0x00, 0x00, // padding to reach 12 bytes
}

// TIFF header bytes (big endian) - needs at least 12 bytes for detection
var tiffHeaderBE = []byte{
	0x4D, 0x4D, 0x00, 0x2A, // MM + magic number
	0x00, 0x00, 0x00, 0x08, // offset to first IFD
	0x00, 0x00, 0x00, 0x00, // padding to reach 12 bytes
}

// SVG header
var svgData = []byte(
	"<?xml version=\"1.0\" encoding=\"UTF-8\"?><svg></svg>",
)

// SVG without XML declaration
var svgDataNoXML = []byte(
	"<svg xmlns=\"http://www.w3.org/2000/svg\"></svg>",
)

func TestDetectImageType(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected string
	}{
		{"PNG", pngData, ContentTypePNG},
		{"JPEG", jpegData, ContentTypeJPEG},
		{"GIF", gifData, ContentTypeGIF},
		{"BMP", bmpHeader, ContentTypeBMP},
		{"WEBP", webpHeader, ContentTypeWEBP},
		{
			"TIFF LE",
			tiffHeaderLE,
			ContentTypeTIFF,
		},
		{
			"TIFF BE",
			tiffHeaderBE,
			ContentTypeTIFF,
		},
		{"SVG with XML", svgData, ContentTypeSVG},
		{
			"SVG without XML",
			svgDataNoXML,
			ContentTypeSVG,
		},
		{"Empty", make([]byte, 0), ""},
		{
			"Too short",
			[]byte{0x00, 0x01, 0x02},
			"",
		},
		{
			"Unknown",
			[]byte{
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
				0x0A,
				0x0B,
			},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DetectImageType(tt.data)
			if result != tt.expected {
				t.Errorf(
					"DetectImageType() = %q, want %q",
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestDetectImageTypeFromReader(t *testing.T) {
	tests := []struct {
		name         string
		data         []byte
		expectedType string
		expectErr    bool
	}{
		{"PNG", pngData, ContentTypePNG, false},
		{
			"JPEG",
			jpegData,
			ContentTypeJPEG,
			false,
		},
		{"GIF", gifData, ContentTypeGIF, false},
		{"Empty", make([]byte, 0), "", false},
		{"Nil reader", nil, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reader io.Reader
			if tt.data != nil {
				reader = bytes.NewReader(tt.data)
			}

			contentType, newReader, err := DetectImageTypeFromReader(
				reader,
			)
			if tt.expectErr {
				if err == nil {
					t.Error(
						"Expected error, got nil",
					)
				}

				return
			}
			if err != nil {
				t.Errorf(
					"Unexpected error: %v",
					err,
				)

				return
			}
			if contentType != tt.expectedType {
				t.Errorf(
					"contentType = %q, want %q",
					contentType,
					tt.expectedType,
				)
			}

			// Verify the new reader contains the original data
			if len(tt.data) == 0 {
				return
			}
			readBack, readErr := io.ReadAll(
				newReader,
			)
			if readErr != nil {
				t.Errorf(
					"Failed to read from new reader: %v",
					readErr,
				)

				return
			}
			if !bytes.Equal(
				readBack,
				tt.data,
			) {
				t.Error(
					"New reader does not contain original data",
				)
			}
		})
	}
}

func TestImageExtension(t *testing.T) {
	tests := []struct {
		contentType string
		expected    string
	}{
		{ContentTypeJPEG, ".jpg"},
		{ContentTypePNG, ".png"},
		{ContentTypeGIF, ".gif"},
		{ContentTypeBMP, ".bmp"},
		{ContentTypeWEBP, ".webp"},
		{ContentTypeTIFF, ".tiff"},
		{ContentTypeSVG, ".svg"},
		{ContentTypeEMF, ".emf"},
		{ContentTypeWMF, ".wmf"},
		{"unknown/type", ""},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.contentType, func(t *testing.T) {
			result := ImageExtension(
				tt.contentType,
			)
			if result != tt.expected {
				t.Errorf(
					"ImageExtension(%q) = %q, want %q",
					tt.contentType,
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestGetImageDimensions(t *testing.T) {
	// Test with PNG (1x1)
	t.Run("PNG 1x1", func(t *testing.T) {
		width, height, err := GetImageDimensions(
			pngData,
		)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)

			return
		}
		if width != 1 || height != 1 {
			t.Errorf(
				"Got dimensions %dx%d, want 1x1",
				width,
				height,
			)
		}
	})

	// Test with JPEG (2x2)
	t.Run("JPEG 2x2", func(t *testing.T) {
		width, height, err := GetImageDimensions(
			jpegData,
		)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)

			return
		}
		if width != 2 || height != 2 {
			t.Errorf(
				"Got dimensions %dx%d, want 2x2",
				width,
				height,
			)
		}
	})

	// Test with GIF (1x1)
	t.Run("GIF 1x1", func(t *testing.T) {
		width, height, err := GetImageDimensions(
			gifData,
		)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)

			return
		}
		if width != 1 || height != 1 {
			t.Errorf(
				"Got dimensions %dx%d, want 1x1",
				width,
				height,
			)
		}
	})

	// Test with empty data
	t.Run("Empty data", func(t *testing.T) {
		_, _, err := GetImageDimensions(
			make([]byte, 0),
		)
		if err != ErrEmptyImageData {
			t.Errorf(
				"Expected ErrEmptyImageData, got %v",
				err,
			)
		}
	})
}

func TestGetImageDimensionsFromReader(
	t *testing.T,
) {
	t.Run("PNG from reader", func(t *testing.T) {
		width, height, err := GetImageDimensionsFromReader(
			bytes.NewReader(pngData),
		)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)

			return
		}
		if width != 1 || height != 1 {
			t.Errorf(
				"Got dimensions %dx%d, want 1x1",
				width,
				height,
			)
		}
	})

	t.Run("Nil reader", func(t *testing.T) {
		_, _, err := GetImageDimensionsFromReader(
			nil,
		)
		if err != ErrEmptyImageData {
			t.Errorf(
				"Expected ErrEmptyImageData, got %v",
				err,
			)
		}
	})
}

func TestImageExtentFromPixels(t *testing.T) {
	tests := []struct {
		name       string
		widthPx    int
		heightPx   int
		dpi        float64
		wantWidth  EMU
		wantHeight EMU
	}{
		{
			name:     "96 DPI default",
			widthPx:  100,
			heightPx: 100,
			dpi:      96,
			wantWidth: EMU(
				952500,
			), // 100 * 914400 / 96 = 952500
			wantHeight: EMU(952500),
		},
		{
			name:     "72 DPI",
			widthPx:  100,
			heightPx: 200,
			dpi:      72,
			wantWidth: EMU(
				1270000,
			), // 100 * 914400 / 72 = 1270000
			wantHeight: EMU(
				2540000,
			), // 200 * 914400 / 72 = 2540000
		},
		{
			name:       "Zero DPI defaults to 96",
			widthPx:    100,
			heightPx:   100,
			dpi:        0,
			wantWidth:  EMU(952500),
			wantHeight: EMU(952500),
		},
		{
			name:       "Negative DPI defaults to 96",
			widthPx:    100,
			heightPx:   100,
			dpi:        -50,
			wantWidth:  EMU(952500),
			wantHeight: EMU(952500),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWidth, gotHeight := ImageExtentFromPixels(
				tt.widthPx,
				tt.heightPx,
				tt.dpi,
			)
			if gotWidth != tt.wantWidth {
				t.Errorf(
					"width = %d, want %d",
					gotWidth,
					tt.wantWidth,
				)
			}
			if gotHeight != tt.wantHeight {
				t.Errorf(
					"height = %d, want %d",
					gotHeight,
					tt.wantHeight,
				)
			}
		})
	}
}

func TestImageExtentFromPixelsDefault(
	t *testing.T,
) {
	widthEMU, heightEMU := ImageExtentFromPixelsDefault(
		100,
		200,
	)

	// At 96 DPI: EMU = pixels * EMUsPerInch / DPI = pixels * 914400 / 96 = pixels * 9525
	expectedWidth := EMU(100 * 9525)
	expectedHeight := EMU(200 * 9525)

	if widthEMU != expectedWidth {
		t.Errorf(
			"widthEMU = %d, want %d",
			widthEMU,
			expectedWidth,
		)
	}
	if heightEMU != expectedHeight {
		t.Errorf(
			"heightEMU = %d, want %d",
			heightEMU,
			expectedHeight,
		)
	}
}

func TestImageExtentFromData(t *testing.T) {
	t.Run("PNG data", func(t *testing.T) {
		widthEMU, heightEMU, err := ImageExtentFromData(
			pngData,
			96,
		)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)

			return
		}
		// 1 pixel at 96 DPI = 914400/96 = 9525 EMUs
		if widthEMU != 9525 || heightEMU != 9525 {
			t.Errorf(
				"Got %dx%d EMUs, want 9525x9525",
				widthEMU,
				heightEMU,
			)
		}
	})

	t.Run("Empty data", func(t *testing.T) {
		_, _, err := ImageExtentFromData(
			make([]byte, 0),
			96,
		)
		if err == nil {
			t.Error(
				"Expected error for empty data",
			)
		}
	})
}

func TestImageExtentFromDataDefault(
	t *testing.T,
) {
	widthEMU, heightEMU, err := ImageExtentFromDataDefault(
		pngData,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)

		return
	}
	// 1 pixel at 96 DPI = 914400/96 = 9525 EMUs
	if widthEMU != 9525 || heightEMU != 9525 {
		t.Errorf(
			"Got %dx%d EMUs, want 9525x9525",
			widthEMU,
			heightEMU,
		)
	}
}

func TestFitImageToMaxSize(t *testing.T) {
	tests := []struct {
		name         string
		widthEMU     EMU
		heightEMU    EMU
		maxWidthEMU  EMU
		maxHeightEMU EMU
		wantWidth    EMU
		wantHeight   EMU
	}{
		{
			name:         "Already fits",
			widthEMU:     100,
			heightEMU:    100,
			maxWidthEMU:  200,
			maxHeightEMU: 200,
			wantWidth:    100,
			wantHeight:   100,
		},
		{
			name:         "Width constrained",
			widthEMU:     400,
			heightEMU:    200,
			maxWidthEMU:  200,
			maxHeightEMU: 200,
			wantWidth:    200,
			wantHeight:   100,
		},
		{
			name:         "Height constrained",
			widthEMU:     200,
			heightEMU:    400,
			maxWidthEMU:  200,
			maxHeightEMU: 200,
			wantWidth:    100,
			wantHeight:   200,
		},
		{
			name:         "Both constrained - width is limiting",
			widthEMU:     600,
			heightEMU:    400,
			maxWidthEMU:  300,
			maxHeightEMU: 300,
			wantWidth:    300,
			wantHeight:   200,
		},
		{
			name:         "Both constrained - height is limiting",
			widthEMU:     400,
			heightEMU:    600,
			maxWidthEMU:  300,
			maxHeightEMU: 300,
			wantWidth:    200,
			wantHeight:   300,
		},
		{
			name:         "Zero width",
			widthEMU:     0,
			heightEMU:    100,
			maxWidthEMU:  200,
			maxHeightEMU: 200,
			wantWidth:    0,
			wantHeight:   100,
		},
		{
			name:         "Zero height",
			widthEMU:     100,
			heightEMU:    0,
			maxWidthEMU:  200,
			maxHeightEMU: 200,
			wantWidth:    100,
			wantHeight:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWidth, gotHeight := FitImageToMaxSize(
				tt.widthEMU,
				tt.heightEMU,
				tt.maxWidthEMU,
				tt.maxHeightEMU,
			)
			if gotWidth != tt.wantWidth {
				t.Errorf(
					"width = %d, want %d",
					gotWidth,
					tt.wantWidth,
				)
			}
			if gotHeight != tt.wantHeight {
				t.Errorf(
					"height = %d, want %d",
					gotHeight,
					tt.wantHeight,
				)
			}
		})
	}
}

func TestFitImageToWidth(t *testing.T) {
	tests := []struct {
		name           string
		widthEMU       EMU
		heightEMU      EMU
		targetWidthEMU EMU
		wantWidth      EMU
		wantHeight     EMU
	}{
		{
			name:           "Scale down",
			widthEMU:       400,
			heightEMU:      200,
			targetWidthEMU: 200,
			wantWidth:      200,
			wantHeight:     100,
		},
		{
			name:           "Scale up",
			widthEMU:       100,
			heightEMU:      50,
			targetWidthEMU: 200,
			wantWidth:      200,
			wantHeight:     100,
		},
		{
			name:           "Zero original width",
			widthEMU:       0,
			heightEMU:      100,
			targetWidthEMU: 200,
			wantWidth:      200,
			wantHeight:     100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWidth, gotHeight := FitImageToWidth(
				tt.widthEMU,
				tt.heightEMU,
				tt.targetWidthEMU,
			)
			if gotWidth != tt.wantWidth {
				t.Errorf(
					"width = %d, want %d",
					gotWidth,
					tt.wantWidth,
				)
			}
			if gotHeight != tt.wantHeight {
				t.Errorf(
					"height = %d, want %d",
					gotHeight,
					tt.wantHeight,
				)
			}
		})
	}
}

func TestFitImageToHeight(t *testing.T) {
	tests := []struct {
		name            string
		widthEMU        EMU
		heightEMU       EMU
		targetHeightEMU EMU
		wantWidth       EMU
		wantHeight      EMU
	}{
		{
			name:            "Scale down",
			widthEMU:        200,
			heightEMU:       400,
			targetHeightEMU: 200,
			wantWidth:       100,
			wantHeight:      200,
		},
		{
			name:            "Scale up",
			widthEMU:        50,
			heightEMU:       100,
			targetHeightEMU: 200,
			wantWidth:       100,
			wantHeight:      200,
		},
		{
			name:            "Zero original height",
			widthEMU:        100,
			heightEMU:       0,
			targetHeightEMU: 200,
			wantWidth:       100,
			wantHeight:      200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWidth, gotHeight := FitImageToHeight(
				tt.widthEMU,
				tt.heightEMU,
				tt.targetHeightEMU,
			)
			if gotWidth != tt.wantWidth {
				t.Errorf(
					"width = %d, want %d",
					gotWidth,
					tt.wantWidth,
				)
			}
			if gotHeight != tt.wantHeight {
				t.Errorf(
					"height = %d, want %d",
					gotHeight,
					tt.wantHeight,
				)
			}
		})
	}
}

func TestExtentFromImageData(t *testing.T) {
	extent, err := ExtentFromImageData(pngData)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)

		return
	}

	// 1x1 pixel at 96 DPI = 9525x9525 EMUs
	if extent.Cx != 9525 || extent.Cy != 9525 {
		t.Errorf(
			"Got extent %dx%d, want 9525x9525",
			extent.Cx,
			extent.Cy,
		)
	}
}

func TestExtentFromImageDataAtDPI(t *testing.T) {
	extent, err := ExtentFromImageDataAtDPI(
		pngData,
		72,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)

		return
	}

	// 1 pixel at 72 DPI = 914400/72 = 12700 EMUs
	if extent.Cx != 12700 || extent.Cy != 12700 {
		t.Errorf(
			"Got extent %dx%d, want 12700x12700",
			extent.Cx,
			extent.Cy,
		)
	}
}

func TestConstants(t *testing.T) {
	// Verify content type constants are correct
	tests := []struct {
		constant string
		value    string
	}{
		{ContentTypeJPEG, "image/jpeg"},
		{ContentTypePNG, "image/png"},
		{ContentTypeGIF, "image/gif"},
		{ContentTypeBMP, "image/bmp"},
		{ContentTypeWEBP, "image/webp"},
		{ContentTypeTIFF, "image/tiff"},
		{ContentTypeSVG, "image/svg+xml"},
		{ContentTypeEMF, "image/x-emf"},
		{ContentTypeWMF, "image/x-wmf"},
	}

	for _, tt := range tests {
		if tt.constant != tt.value {
			t.Errorf(
				"Constant value mismatch: got %q, want %q",
				tt.constant,
				tt.value,
			)
		}
	}

	// Verify default DPI
	if DefaultImageDPI != 96.0 {
		t.Errorf(
			"DefaultImageDPI = %f, want 96.0",
			DefaultImageDPI,
		)
	}
}

func TestErrorConstants(t *testing.T) {
	// Verify error constants are defined and have appropriate messages
	if ErrEmptyImageData == nil {
		t.Error("ErrEmptyImageData is nil")
	}
	if ErrUnknownImageFormat == nil {
		t.Error("ErrUnknownImageFormat is nil")
	}
	if ErrInvalidImageDimensions == nil {
		t.Error(
			"ErrInvalidImageDimensions is nil",
		)
	}

	// Verify error messages
	if ErrEmptyImageData.Error() != "image data is empty" {
		t.Errorf(
			"ErrEmptyImageData message = %q",
			ErrEmptyImageData.Error(),
		)
	}
	if ErrUnknownImageFormat.Error() != "unknown image format" {
		t.Errorf(
			"ErrUnknownImageFormat message = %q",
			ErrUnknownImageFormat.Error(),
		)
	}
	if ErrInvalidImageDimensions.Error() != "invalid image dimensions" {
		t.Errorf(
			"ErrInvalidImageDimensions message = %q",
			ErrInvalidImageDimensions.Error(),
		)
	}
}
