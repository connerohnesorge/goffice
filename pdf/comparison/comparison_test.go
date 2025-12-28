package comparison

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestCompareImages_Identical(t *testing.T) {
	// Create a simple test image
	img := createTestImage(
		100,
		100,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
	)

	// Save as two files
	tmpDir := t.TempDir()
	path1 := filepath.Join(tmpDir, "img1.png")
	path2 := filepath.Join(tmpDir, "img2.png")

	if err := saveTestImage(img, path1); err != nil {
		t.Fatalf(
			"Failed to save test image: %v",
			err,
		)
	}
	if err := saveTestImage(img, path2); err != nil {
		t.Fatalf(
			"Failed to save test image: %v",
			err,
		)
	}

	// Compare
	result, err := CompareImages(
		path1,
		path2,
		nil,
	)
	if err != nil {
		t.Fatalf("CompareImages failed: %v", err)
	}

	if result.DiffPixels != 0 {
		t.Errorf(
			"Expected 0 different pixels, got %d",
			result.DiffPixels,
		)
	}

	if result.PixelsExceedingTolerance != 0 {
		t.Errorf(
			"Expected 0 pixels exceeding tolerance, got %d",
			result.PixelsExceedingTolerance,
		)
	}

	if !result.IsWithinTolerance() {
		t.Error(
			"Expected comparison to be within tolerance",
		)
	}
}

func TestCompareImages_SlightDifference(
	t *testing.T,
) {
	tmpDir := t.TempDir()

	// Create two images with slight color difference
	img1 := createTestImage(
		100,
		100,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
	)
	img2 := createTestImage(
		100,
		100,
		color.RGBA{R: 254, G: 1, B: 1, A: 255},
	) // Slight difference

	path1 := filepath.Join(tmpDir, "img1.png")
	path2 := filepath.Join(tmpDir, "img2.png")

	if err := saveTestImage(img1, path1); err != nil {
		t.Fatalf(
			"Failed to save test image: %v",
			err,
		)
	}
	if err := saveTestImage(img2, path2); err != nil {
		t.Fatalf(
			"Failed to save test image: %v",
			err,
		)
	}

	// Compare
	result, err := CompareImages(
		path1,
		path2,
		nil,
	)
	if err != nil {
		t.Fatalf("CompareImages failed: %v", err)
	}

	if result.DiffPixels == 0 {
		t.Error("Expected some pixels to differ")
	}

	// Small color difference should be within tolerance
	if !result.IsWithinTolerance() {
		t.Errorf(
			"Expected slight differences to be within tolerance, got %s",
			result.String(),
		)
	}
}

func TestCompareImages_LargeDifference(
	t *testing.T,
) {
	tmpDir := t.TempDir()

	// Create two very different images
	img1 := createTestImage(
		100,
		100,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
	) // Red
	img2 := createTestImage(
		100,
		100,
		color.RGBA{R: 0, G: 0, B: 255, A: 255},
	) // Blue

	path1 := filepath.Join(tmpDir, "img1.png")
	path2 := filepath.Join(tmpDir, "img2.png")

	if err := saveTestImage(img1, path1); err != nil {
		t.Fatalf(
			"Failed to save test image: %v",
			err,
		)
	}
	if err := saveTestImage(img2, path2); err != nil {
		t.Fatalf(
			"Failed to save test image: %v",
			err,
		)
	}

	// Compare
	result, err := CompareImages(
		path1,
		path2,
		nil,
	)
	if err != nil {
		t.Fatalf("CompareImages failed: %v", err)
	}

	if result.DiffPixels != result.TotalPixels {
		t.Errorf(
			"Expected all pixels to differ, got %d/%d",
			result.DiffPixels,
			result.TotalPixels,
		)
	}

	if result.IsWithinTolerance() {
		t.Error(
			"Expected large differences to exceed tolerance",
		)
	}
}

func TestCompareImages_DimensionMismatch(
	t *testing.T,
) {
	tmpDir := t.TempDir()

	// Create images with different dimensions
	img1 := createTestImage(
		100,
		100,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
	)
	img2 := createTestImage(
		200,
		200,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
	)

	path1 := filepath.Join(tmpDir, "img1.png")
	path2 := filepath.Join(tmpDir, "img2.png")

	if err := saveTestImage(img1, path1); err != nil {
		t.Fatalf(
			"Failed to save test image: %v",
			err,
		)
	}
	if err := saveTestImage(img2, path2); err != nil {
		t.Fatalf(
			"Failed to save test image: %v",
			err,
		)
	}

	// Compare should fail
	_, err := CompareImages(path1, path2, nil)
	if err == nil {
		t.Error(
			"Expected error for dimension mismatch",
		)
	}
}

func TestAnnotateDiffImage(t *testing.T) {
	tmpDir := t.TempDir()

	// Create two different images
	img1 := createTestImage(
		50,
		50,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
	)
	img2 := createTestImage(
		50,
		50,
		color.RGBA{R: 0, G: 0, B: 255, A: 255},
	)

	path1 := filepath.Join(tmpDir, "img1.png")
	path2 := filepath.Join(tmpDir, "img2.png")
	diffPath := filepath.Join(tmpDir, "diff.png")

	if err := saveTestImage(img1, path1); err != nil {
		t.Fatalf(
			"Failed to save test image: %v",
			err,
		)
	}
	if err := saveTestImage(img2, path2); err != nil {
		t.Fatalf(
			"Failed to save test image: %v",
			err,
		)
	}

	// Create diff image
	if err := AnnotateDiffImage(path1, path2, diffPath, nil); err != nil {
		t.Fatalf(
			"AnnotateDiffImage failed: %v",
			err,
		)
	}

	// Verify diff image exists
	if _, err := os.Stat(diffPath); os.IsNotExist(
		err,
	) {
		t.Error("Diff image was not created")
	}

	// Verify it's a valid PNG
	_, err := loadPNG(diffPath)
	if err != nil {
		t.Errorf(
			"Diff image is not a valid PNG: %v",
			err,
		)
	}
}

func TestIsGhostscriptAvailable(t *testing.T) {
	// This test just verifies the function doesn't panic
	available := IsGhostscriptAvailable()
	t.Logf("Ghostscript available: %v", available)
}

func TestConvertPDFToPNG_ValidPDF(t *testing.T) {
	// Skip if Ghostscript is not available
	if !IsGhostscriptAvailable() {
		t.Skip(
			"Ghostscript not available, skipping test",
		)
	}

	tmpDir := t.TempDir()

	// Create a minimal valid PDF file
	pdfPath := filepath.Join(tmpDir, "test.pdf")
	pdfBytes := createMinimalPDF()
	if err := os.WriteFile(pdfPath, pdfBytes, 0644); err != nil {
		t.Fatalf(
			"Failed to create test PDF: %v",
			err,
		)
	}

	// Convert PDF to PNG
	outputPath := filepath.Join(
		tmpDir,
		"output.png",
	)
	resultPath, err := ConvertPDFToPNG(
		pdfPath,
		outputPath,
		150,
	)
	if err != nil {
		t.Fatalf(
			"ConvertPDFToPNG failed: %v",
			err,
		)
	}

	// Verify output path matches
	if resultPath != outputPath {
		t.Errorf(
			"Expected output path %s, got %s",
			outputPath,
			resultPath,
		)
	}

	// Verify output file exists
	stat, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf(
			"Output PNG file does not exist: %v",
			err,
		)
	}

	// Verify file size is greater than 0
	if stat.Size() == 0 {
		t.Error("Output PNG file is empty")
	}

	// Verify it's a valid PNG
	img, err := loadPNG(outputPath)
	if err != nil {
		t.Errorf(
			"Output is not a valid PNG: %v",
			err,
		)
	}

	// Verify image has reasonable dimensions
	if img != nil {
		bounds := img.Bounds()
		if bounds.Dx() == 0 || bounds.Dy() == 0 {
			t.Error("PNG has zero dimensions")
		}
		t.Logf(
			"Generated PNG dimensions: %dx%d",
			bounds.Dx(),
			bounds.Dy(),
		)
	}
}

func TestConvertPDFToPNG_OutputValidation(
	t *testing.T,
) {
	// Skip if Ghostscript is not available
	if !IsGhostscriptAvailable() {
		t.Skip(
			"Ghostscript not available, skipping test",
		)
	}

	tmpDir := t.TempDir()

	// Create a minimal valid PDF file
	pdfPath := filepath.Join(tmpDir, "test.pdf")
	pdfBytes := createMinimalPDF()
	if err := os.WriteFile(pdfPath, pdfBytes, 0644); err != nil {
		t.Fatalf(
			"Failed to create test PDF: %v",
			err,
		)
	}

	t.Run("ValidOutput", func(t *testing.T) {
		outputPath := filepath.Join(
			tmpDir,
			"valid_output.png",
		)
		_, err := ConvertPDFToPNG(
			pdfPath,
			outputPath,
			150,
		)
		if err != nil {
			t.Fatalf(
				"ConvertPDFToPNG failed: %v",
				err,
			)
		}

		// Verify file size > 0
		stat, err := os.Stat(outputPath)
		if err != nil {
			t.Fatalf(
				"Output file does not exist: %v",
				err,
			)
		}
		if stat.Size() == 0 {
			t.Error("Output file size is 0")
		}

		// Verify it's a valid PNG image
		img, err := loadPNG(outputPath)
		if err != nil {
			t.Errorf(
				"Failed to load PNG: %v",
				err,
			)
		}
		if img == nil {
			t.Error("Loaded image is nil")
		}
	})

	t.Run(
		"InvalidOutputPath",
		func(t *testing.T) {
			// Try to write to a directory that doesn't exist and can't be created
			// Use a path that would fail on permission issues
			invalidPath := "/invalid/nonexistent/path/output.png"
			_, err := ConvertPDFToPNG(
				pdfPath,
				invalidPath,
				150,
			)
			if err == nil {
				t.Error(
					"Expected error for invalid output path, got nil",
				)
			}
		},
	)

	t.Run("InvalidPDFPath", func(t *testing.T) {
		outputPath := filepath.Join(
			tmpDir,
			"output_from_invalid.png",
		)
		nonexistentPDF := filepath.Join(
			tmpDir,
			"nonexistent.pdf",
		)
		_, err := ConvertPDFToPNG(
			nonexistentPDF,
			outputPath,
			150,
		)
		if err == nil {
			t.Error(
				"Expected error for nonexistent PDF, got nil",
			)
		}
	})

	t.Run("CustomDPI", func(t *testing.T) {
		// Test with custom DPI
		outputPath := filepath.Join(
			tmpDir,
			"custom_dpi.png",
		)
		_, err := ConvertPDFToPNG(
			pdfPath,
			outputPath,
			300,
		)
		if err != nil {
			t.Fatalf(
				"ConvertPDFToPNG with custom DPI failed: %v",
				err,
			)
		}

		// Verify output exists
		if _, err := os.Stat(outputPath); err != nil {
			t.Errorf(
				"Output file with custom DPI does not exist: %v",
				err,
			)
		}
	})

	t.Run("DefaultDPI", func(t *testing.T) {
		// Test with default DPI (0 should use 150)
		outputPath := filepath.Join(
			tmpDir,
			"default_dpi.png",
		)
		_, err := ConvertPDFToPNG(
			pdfPath,
			outputPath,
			0,
		)
		if err != nil {
			t.Fatalf(
				"ConvertPDFToPNG with default DPI failed: %v",
				err,
			)
		}

		// Verify output exists
		if _, err := os.Stat(outputPath); err != nil {
			t.Errorf(
				"Output file with default DPI does not exist: %v",
				err,
			)
		}
	})
}

func TestConvertPDFBytesToPNG_ValidPDF(
	t *testing.T,
) {
	// Skip if Ghostscript is not available
	if !IsGhostscriptAvailable() {
		t.Skip(
			"Ghostscript not available, skipping test",
		)
	}

	tmpDir := t.TempDir()

	// Create a minimal valid PDF
	pdfBytes := createMinimalPDF()

	// Convert PDF bytes to PNG
	outputPath := filepath.Join(
		tmpDir,
		"output_from_bytes.png",
	)
	resultPath, err := ConvertPDFBytesToPNG(
		pdfBytes,
		outputPath,
		150,
	)
	if err != nil {
		t.Fatalf(
			"ConvertPDFBytesToPNG failed: %v",
			err,
		)
	}

	// Verify output path matches
	if resultPath != outputPath {
		t.Errorf(
			"Expected output path %s, got %s",
			outputPath,
			resultPath,
		)
	}

	// Verify output file exists
	stat, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf(
			"Output PNG file does not exist: %v",
			err,
		)
	}

	// Verify file size is greater than 0
	if stat.Size() == 0 {
		t.Error("Output PNG file is empty")
	}

	// Verify it's a valid PNG
	img, err := loadPNG(outputPath)
	if err != nil {
		t.Errorf(
			"Output is not a valid PNG: %v",
			err,
		)
	}

	// Verify image has reasonable dimensions
	if img != nil {
		bounds := img.Bounds()
		if bounds.Dx() == 0 || bounds.Dy() == 0 {
			t.Error("PNG has zero dimensions")
		}
		t.Logf(
			"Generated PNG dimensions: %dx%d",
			bounds.Dx(),
			bounds.Dy(),
		)
	}

	t.Run("EmptyPDFBytes", func(t *testing.T) {
		outputPath := filepath.Join(
			tmpDir,
			"output_from_empty.png",
		)
		_, err := ConvertPDFBytesToPNG(
			[]byte{},
			outputPath,
			150,
		)
		if err == nil {
			t.Error(
				"Expected error for empty PDF bytes, got nil",
			)
		}
	})

	t.Run("InvalidPDFBytes", func(t *testing.T) {
		outputPath := filepath.Join(
			tmpDir,
			"output_from_invalid.png",
		)
		invalidBytes := []byte(
			"This is not a valid PDF",
		)
		_, err := ConvertPDFBytesToPNG(
			invalidBytes,
			outputPath,
			150,
		)
		if err == nil {
			t.Error(
				"Expected error for invalid PDF bytes, got nil",
			)
		}
	})
}

// Helper functions for testing

func createTestImage(
	width, height int,
	c color.Color,
) *image.RGBA {
	img := image.NewRGBA(
		image.Rect(0, 0, width, height),
	)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, c)
		}
	}
	return img
}

func saveTestImage(
	img image.Image,
	path string,
) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// createMinimalPDF creates a minimal valid PDF file as a byte slice.
// This PDF contains a single blank page.
func createMinimalPDF() []byte {
	// Minimal PDF structure with a single blank page
	// This is a valid PDF 1.4 document with the bare minimum structure
	pdf := `%PDF-1.4
1 0 obj
<<
/Type /Catalog
/Pages 2 0 R
>>
endobj
2 0 obj
<<
/Type /Pages
/Kids [3 0 R]
/Count 1
>>
endobj
3 0 obj
<<
/Type /Page
/Parent 2 0 R
/Resources <<
/Font <<
/F1 <<
/Type /Font
/Subtype /Type1
/BaseFont /Helvetica
>>
>>
>>
/MediaBox [0 0 612 792]
/Contents 4 0 R
>>
endobj
4 0 obj
<<
/Length 44
>>
stream
BT
/F1 12 Tf
100 700 Td
(Test PDF) Tj
ET
endstream
endobj
xref
0 5
0000000000 65535 f
0000000009 00000 n
0000000058 00000 n
0000000115 00000 n
0000000317 00000 n
trailer
<<
/Size 5
/Root 1 0 R
>>
startxref
410
%%EOF`
	return []byte(pdf)
}
