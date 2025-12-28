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
