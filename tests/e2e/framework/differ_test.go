package framework

import (
	"context"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

// createTestImage creates a simple test image with the given color
func createTestImage(
	width, height int,
	c color.Color,
) image.Image {
	img := image.NewRGBA(
		image.Rect(0, 0, width, height),
	)
	for y := range height {
		for x := range width {
			img.Set(x, y, c)
		}
	}

	return img
}

// createGradientImage creates a gradient test image
func createGradientImage(
	width, height int,
) image.Image {
	img := image.NewRGBA(
		image.Rect(0, 0, width, height),
	)
	for y := range height {
		for x := range width {
			// Create a gradient from black to white
			intensity := uint8(
				float64(x) / float64(width) * 255,
			)
			img.Set(
				x,
				y,
				color.RGBA{
					R: intensity,
					G: intensity,
					B: intensity,
					A: 255,
				},
			)
		}
	}

	return img
}

// createSlightlyDifferentImage creates an image that's mostly the same but with a small patch different
func createSlightlyDifferentImage(
	width, height int,
	baseColor, patchColor color.Color,
	patchSize int,
) image.Image {
	img := image.NewRGBA(
		image.Rect(0, 0, width, height),
	)
	for y := range height {
		for x := range width {
			// Create a small patch in the center with different color
			if x >= width/2-patchSize/2 &&
				x < width/2+patchSize/2 &&
				y >= height/2-patchSize/2 &&
				y < height/2+patchSize/2 {
				img.Set(x, y, patchColor)
			} else {
				img.Set(x, y, baseColor)
			}
		}
	}

	return img
}

// saveTestImage saves an image to a temporary file
func saveTestImage(
	t *testing.T,
	img image.Image,
	name string,
) string {
	t.Helper()

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, name)

	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("create test image: %v", err)
	}
	defer func() {
		_ = file.Close()
	}()

	if err := png.Encode(file, img); err != nil {
		t.Fatalf("encode test image: %v", err)
	}

	return path
}

func TestGoDiffer_IdenticalImages(t *testing.T) {
	differ := NewGoDiffer()
	ctx := context.Background()

	// Create two identical images
	img1 := createTestImage(
		100,
		100,
		color.RGBA{
			R: 128,
			G: 128,
			B: 128,
			A: 255,
		},
	)
	img2 := createTestImage(
		100,
		100,
		color.RGBA{
			R: 128,
			G: 128,
			B: 128,
			A: 255,
		},
	)

	img1Path := saveTestImage(t, img1, "img1.png")
	img2Path := saveTestImage(t, img2, "img2.png")
	diffPath := filepath.Join(
		filepath.Dir(img1Path),
		"diff.png",
	)

	config := DefaultTestConfig()
	config.DiffAlgorithm = DiffAlgorithmPixelPerfect
	config.DiffThreshold = 0.01 // 1% threshold

	result, err := differ.Compare(
		ctx,
		img1Path,
		img2Path,
		diffPath,
		config,
	)
	if err != nil {
		t.Fatalf("Compare failed: %v", err)
	}

	// Verify results
	if result.PixelDiff != 0 {
		t.Errorf(
			"Expected PixelDiff=0, got %d",
			result.PixelDiff,
		)
	}

	if result.DiffPercent != 0.0 {
		t.Errorf(
			"Expected DiffPercent=0.0, got %f",
			result.DiffPercent,
		)
	}

	if !result.Passed {
		t.Error(
			"Expected test to pass for identical images",
		)
	}

	if result.TotalPixels != 10000 {
		t.Errorf(
			"Expected TotalPixels=10000, got %d",
			result.TotalPixels,
		)
	}

	// Verify diff image was created
	if _, err := os.Stat(diffPath); os.IsNotExist(
		err,
	) {
		t.Error("Diff image was not created")
	}
}

func TestGoDiffer_DifferentImages(t *testing.T) {
	differ := NewGoDiffer()
	ctx := context.Background()

	// Create two completely different images
	img1 := createTestImage(
		100,
		100,
		color.RGBA{R: 0, G: 0, B: 0, A: 255},
	) // Black
	img2 := createTestImage(
		100,
		100,
		color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
	) // White

	img1Path := saveTestImage(t, img1, "img1.png")
	img2Path := saveTestImage(t, img2, "img2.png")
	diffPath := filepath.Join(
		filepath.Dir(img1Path),
		"diff.png",
	)

	config := DefaultTestConfig()
	config.DiffAlgorithm = DiffAlgorithmPixelPerfect
	config.DiffThreshold = 0.01 // 1% threshold

	result, err := differ.Compare(
		ctx,
		img1Path,
		img2Path,
		diffPath,
		config,
	)
	if err != nil {
		t.Fatalf("Compare failed: %v", err)
	}

	// Verify results
	if result.PixelDiff != 10000 {
		t.Errorf(
			"Expected PixelDiff=10000, got %d",
			result.PixelDiff,
		)
	}

	if result.DiffPercent != 100.0 {
		t.Errorf(
			"Expected DiffPercent=100.0, got %f",
			result.DiffPercent,
		)
	}

	if result.Passed {
		t.Error(
			"Expected test to fail for completely different images",
		)
	}

	// MSE should be non-zero
	if result.MSE == 0.0 {
		t.Error(
			"Expected non-zero MSE for different images",
		)
	}

	// PSNR should be calculated
	if result.PSNR == 0.0 {
		t.Error("Expected non-zero PSNR")
	}
}

func TestGoDiffer_SlightlyDifferentImages(
	t *testing.T,
) {
	differ := NewGoDiffer()
	ctx := context.Background()

	// Create images with a small 10x10 patch difference
	baseColor := color.RGBA{
		R: 128,
		G: 128,
		B: 128,
		A: 255,
	}
	patchColor := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	} // Red patch

	img1 := createSlightlyDifferentImage(
		100,
		100,
		baseColor,
		baseColor,
		10,
	)
	img2 := createSlightlyDifferentImage(
		100,
		100,
		baseColor,
		patchColor,
		10,
	)

	img1Path := saveTestImage(t, img1, "img1.png")
	img2Path := saveTestImage(t, img2, "img2.png")
	diffPath := filepath.Join(
		filepath.Dir(img1Path),
		"diff.png",
	)

	config := DefaultTestConfig()
	config.DiffAlgorithm = DiffAlgorithmPixelPerfect
	config.DiffThreshold = 0.01 // 1% threshold

	result, err := differ.Compare(
		ctx,
		img1Path,
		img2Path,
		diffPath,
		config,
	)
	if err != nil {
		t.Fatalf("Compare failed: %v", err)
	}

	// 10x10 = 100 pixels out of 10000 = 1%
	expectedDiff := 100
	if result.PixelDiff != expectedDiff {
		t.Errorf(
			"Expected PixelDiff=%d, got %d",
			expectedDiff,
			result.PixelDiff,
		)
	}

	expectedPercent := 1.0
	if result.DiffPercent != expectedPercent {
		t.Errorf(
			"Expected DiffPercent=%f, got %f",
			expectedPercent,
			result.DiffPercent,
		)
	}

	// With 1% threshold, this should pass (exactly at threshold, using <=)
	if !result.Passed {
		t.Error(
			"Expected test to pass with 1% difference and 1% threshold",
		)
	}

	// Test with higher threshold
	config.DiffThreshold = 0.02 // 2% threshold
	result, err = differ.Compare(
		ctx,
		img1Path,
		img2Path,
		diffPath,
		config,
	)
	if err != nil {
		t.Fatalf("Compare failed: %v", err)
	}

	if !result.Passed {
		t.Error(
			"Expected test to pass with 1% difference and 2% threshold",
		)
	}
}

func TestGoDiffer_DifferentDimensions(
	t *testing.T,
) {
	differ := NewGoDiffer()
	ctx := context.Background()

	img1 := createTestImage(
		100,
		100,
		color.RGBA{
			R: 128,
			G: 128,
			B: 128,
			A: 255,
		},
	)
	img2 := createTestImage(
		200,
		200,
		color.RGBA{
			R: 128,
			G: 128,
			B: 128,
			A: 255,
		},
	)

	img1Path := saveTestImage(t, img1, "img1.png")
	img2Path := saveTestImage(t, img2, "img2.png")
	diffPath := filepath.Join(
		filepath.Dir(img1Path),
		"diff.png",
	)

	config := DefaultTestConfig()

	_, err := differ.Compare(
		ctx,
		img1Path,
		img2Path,
		diffPath,
		config,
	)
	if err == nil {
		t.Fatal(
			"Expected error for images with different dimensions",
		)
	}

	expectedErrMsg := "image dimensions differ"
	if !contains(err.Error(), expectedErrMsg) {
		t.Errorf(
			"Expected error message to contain '%s', got: %v",
			expectedErrMsg,
			err,
		)
	}
}

func TestGoDiffer_AntiAliasingTolerance(
	t *testing.T,
) {
	differ := NewGoDiffer()
	ctx := context.Background()

	// Create images with very slight color differences (simulating anti-aliasing)
	img1 := createTestImage(
		100,
		100,
		color.RGBA{
			R: 128,
			G: 128,
			B: 128,
			A: 255,
		},
	)
	img2 := createTestImage(
		100,
		100,
		color.RGBA{
			R: 130,
			G: 130,
			B: 130,
			A: 255,
		},
	) // 2 units difference

	img1Path := saveTestImage(t, img1, "img1.png")
	img2Path := saveTestImage(t, img2, "img2.png")
	diffPath := filepath.Join(
		filepath.Dir(img1Path),
		"diff.png",
	)

	// Test without anti-aliasing tolerance
	config := DefaultTestConfig()
	config.IgnoreAntialiasing = false
	config.DiffThreshold = 0.01

	result, err := differ.Compare(
		ctx,
		img1Path,
		img2Path,
		diffPath,
		config,
	)
	if err != nil {
		t.Fatalf("Compare failed: %v", err)
	}

	if result.PixelDiff == 0 {
		t.Error(
			"Expected non-zero PixelDiff without anti-aliasing tolerance",
		)
	}

	// Test with anti-aliasing tolerance (should ignore 2-unit differences)
	config.IgnoreAntialiasing = true
	result, err = differ.Compare(
		ctx,
		img1Path,
		img2Path,
		diffPath,
		config,
	)
	if err != nil {
		t.Fatalf("Compare failed: %v", err)
	}

	if result.PixelDiff != 0 {
		t.Errorf(
			"Expected PixelDiff=0 with anti-aliasing tolerance, got %d",
			result.PixelDiff,
		)
	}
}

func TestGoDiffer_SSIMCalculation(t *testing.T) {
	differ := NewGoDiffer()
	ctx := context.Background()

	// Test with identical images
	img1 := createGradientImage(100, 100)
	img2 := createGradientImage(100, 100)

	img1Path := saveTestImage(t, img1, "img1.png")
	img2Path := saveTestImage(t, img2, "img2.png")
	diffPath := filepath.Join(
		filepath.Dir(img1Path),
		"diff.png",
	)

	config := DefaultTestConfig()
	config.DiffAlgorithm = DiffAlgorithmSSIM
	config.DiffThreshold = 0.01 // 1% tolerance

	result, err := differ.Compare(
		ctx,
		img1Path,
		img2Path,
		diffPath,
		config,
	)
	if err != nil {
		t.Fatalf("Compare failed: %v", err)
	}

	// SSIM should be close to 1.0 for identical images
	if result.SSIM < 0.99 {
		t.Errorf(
			"Expected SSIM close to 1.0 for identical images, got %f",
			result.SSIM,
		)
	}

	if !result.Passed {
		t.Error(
			"Expected test to pass for identical images with SSIM",
		)
	}

	// Test with different images
	img3 := createTestImage(
		100,
		100,
		color.RGBA{R: 0, G: 0, B: 0, A: 255},
	)
	img4 := createTestImage(
		100,
		100,
		color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
	)

	img3Path := saveTestImage(t, img3, "img3.png")
	img4Path := saveTestImage(t, img4, "img4.png")
	diffPath2 := filepath.Join(
		filepath.Dir(img3Path),
		"diff2.png",
	)

	result, err = differ.Compare(
		ctx,
		img3Path,
		img4Path,
		diffPath2,
		config,
	)
	if err != nil {
		t.Fatalf("Compare failed: %v", err)
	}

	// SSIM should be lower for very different images
	if result.SSIM > 0.5 {
		t.Errorf(
			"Expected SSIM < 0.5 for very different images, got %f",
			result.SSIM,
		)
	}

	if result.Passed {
		t.Error(
			"Expected test to fail for very different images with SSIM",
		)
	}
}

func TestCalculateSSIM(t *testing.T) {
	// Test SSIM calculation directly
	tests := []struct {
		name      string
		img1      image.Image
		img2      image.Image
		wantSSIM  float64 // Approximate expected SSIM
		tolerance float64
	}{
		{
			name: "Identical images",
			img1: createTestImage(
				50,
				50,
				color.RGBA{
					R: 128,
					G: 128,
					B: 128,
					A: 255,
				},
			),
			img2: createTestImage(
				50,
				50,
				color.RGBA{
					R: 128,
					G: 128,
					B: 128,
					A: 255,
				},
			),
			wantSSIM:  1.0,
			tolerance: 0.01,
		},
		{
			name: "Opposite colors",
			img1: createTestImage(
				50,
				50,
				color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 255,
				},
			),
			img2: createTestImage(
				50,
				50,
				color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			),
			wantSSIM:  -0.5, // Should be negative for opposite images
			tolerance: 1.0,  // Wide tolerance as exact value depends on implementation
		},
		{
			name: "Identical gradients",
			img1: createGradientImage(
				50,
				50,
			),
			img2: createGradientImage(
				50,
				50,
			),
			wantSSIM:  1.0,
			tolerance: 0.01,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ssim := calculateSSIM(
				tt.img1,
				tt.img2,
			)

			if ssim < tt.wantSSIM-tt.tolerance ||
				ssim > tt.wantSSIM+tt.tolerance {
				t.Errorf(
					"SSIM = %f, want %f ± %f",
					ssim,
					tt.wantSSIM,
					tt.tolerance,
				)
			}

			// SSIM should always be in range [-1, 1]
			if ssim < -1.0 || ssim > 1.0 {
				t.Errorf(
					"SSIM out of range: %f (should be in [-1, 1])",
					ssim,
				)
			}
		})
	}
}

func TestMapDiffAlgorithmToImageMagick(
	t *testing.T,
) {
	tests := []struct {
		algorithm DiffAlgorithm
		want      string
	}{
		{DiffAlgorithmSSIM, "SSIM"},
		{DiffAlgorithmMSE, "MSE"},
		{DiffAlgorithmPSNR, "PSNR"},
		{DiffAlgorithmPerceptual, "PHASH"},
		{DiffAlgorithmPixelPerfect, "AE"},
		{
			DiffAlgorithm("unknown"),
			"AE",
		}, // Default
	}

	for _, tt := range tests {
		t.Run(
			string(tt.algorithm),
			func(t *testing.T) {
				got := mapDiffAlgorithmToImageMagick(
					tt.algorithm,
				)
				if got != tt.want {
					t.Errorf(
						"mapDiffAlgorithmToImageMagick(%s) = %s, want %s",
						tt.algorithm,
						got,
						tt.want,
					)
				}
			},
		)
	}
}

func TestThresholdChecking(t *testing.T) {
	differ := NewGoDiffer()
	ctx := context.Background()

	tests := []struct {
		name         string
		threshold    float64
		algorithm    DiffAlgorithm
		createImages func(t *testing.T) (img1Path, img2Path string)
		wantPassed   bool
	}{
		{
			name:      "Pass with low threshold and identical images",
			threshold: 0.01,
			algorithm: DiffAlgorithmPixelPerfect,
			createImages: func(t *testing.T) (string, string) {
				img := createTestImage(
					50,
					50,
					color.RGBA{
						R: 128,
						G: 128,
						B: 128,
						A: 255,
					},
				)

				return saveTestImage(
						t,
						img,
						"img1.png",
					), saveTestImage(
						t,
						img,
						"img2.png",
					)
			},
			wantPassed: true,
		},
		{
			name:      "Fail with low threshold and different images",
			threshold: 0.01,
			algorithm: DiffAlgorithmPixelPerfect,
			createImages: func(t *testing.T) (string, string) {
				img1 := createTestImage(
					50,
					50,
					color.RGBA{
						R: 0,
						G: 0,
						B: 0,
						A: 255,
					},
				)
				img2 := createTestImage(
					50,
					50,
					color.RGBA{
						R: 255,
						G: 255,
						B: 255,
						A: 255,
					},
				)

				return saveTestImage(
						t,
						img1,
						"img1.png",
					), saveTestImage(
						t,
						img2,
						"img2.png",
					)
			},
			wantPassed: false,
		},
		{
			name:      "Pass with high threshold and slightly different images",
			threshold: 0.05,
			algorithm: DiffAlgorithmPixelPerfect,
			createImages: func(t *testing.T) (string, string) {
				baseColor := color.RGBA{
					R: 128,
					G: 128,
					B: 128,
					A: 255,
				}
				patchColor := color.RGBA{
					R: 255,
					G: 0,
					B: 0,
					A: 255,
				}
				img1 := createSlightlyDifferentImage(
					100,
					100,
					baseColor,
					baseColor,
					5,
				)
				img2 := createSlightlyDifferentImage(
					100,
					100,
					baseColor,
					patchColor,
					5,
				)

				return saveTestImage(
						t,
						img1,
						"img1.png",
					), saveTestImage(
						t,
						img2,
						"img2.png",
					)
			},
			wantPassed: true,
		},
		{
			name:      "SSIM threshold - pass",
			threshold: 0.01,
			algorithm: DiffAlgorithmSSIM,
			createImages: func(t *testing.T) (string, string) {
				img := createGradientImage(50, 50)

				return saveTestImage(
						t,
						img,
						"img1.png",
					), saveTestImage(
						t,
						img,
						"img2.png",
					)
			},
			wantPassed: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img1Path, img2Path := tt.createImages(
				t,
			)
			diffPath := filepath.Join(
				filepath.Dir(img1Path),
				"diff.png",
			)

			config := DefaultTestConfig()
			config.DiffThreshold = tt.threshold
			config.DiffAlgorithm = tt.algorithm

			result, err := differ.Compare(
				ctx,
				img1Path,
				img2Path,
				diffPath,
				config,
			)
			if err != nil {
				t.Fatalf(
					"Compare failed: %v",
					err,
				)
			}

			if result.Passed != tt.wantPassed {
				t.Errorf(
					"Passed = %v, want %v (DiffPercent=%f, SSIM=%f, Threshold=%f)",
					result.Passed,
					tt.wantPassed,
					result.DiffPercent,
					result.SSIM,
					result.Threshold,
				)
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr || len(substr) == 0 ||
			(len(s) > 0 && len(substr) > 0 && findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}

	return false
}
