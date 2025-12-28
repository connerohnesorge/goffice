package comparison

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// CompareImages compares two PNG images and returns detailed difference statistics.
//
// The comparison uses Euclidean color distance in RGB space to measure
// pixel-level differences. Two thresholds are applied:
//   - Per-pixel tolerance: ~2.0 delta (configurable)
//   - Per-image threshold: 0.5% of pixels allowed to exceed tolerance
//
// Parameters:
//   - baselinePath: path to the baseline (reference) image
//   - generatedPath: path to the generated (test) image
//   - config: comparison configuration (nil for defaults)
//
// Returns a DiffResult containing statistics and an error if comparison fails.
func CompareImages(
	baselinePath, generatedPath string,
	config *ComparisonConfig,
) (*DiffResult, error) {
	if config == nil {
		config = DefaultComparisonConfig()
	}

	// Load baseline image
	baselineImg, err := loadPNG(baselinePath)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to load baseline image: %w",
			err,
		)
	}

	// Load generated image
	generatedImg, err := loadPNG(generatedPath)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to load generated image: %w",
			err,
		)
	}

	// Check dimensions match
	baseBounds := baselineImg.Bounds()
	genBounds := generatedImg.Bounds()

	if baseBounds.Dx() != genBounds.Dx() ||
		baseBounds.Dy() != genBounds.Dy() {
		return nil, fmt.Errorf(
			"image dimensions do not match: baseline %dx%d, generated %dx%d",
			baseBounds.Dx(),
			baseBounds.Dy(),
			genBounds.Dx(),
			genBounds.Dy(),
		)
	}

	// Perform pixel-by-pixel comparison
	result := &DiffResult{
		TotalPixels: baseBounds.Dx() * baseBounds.Dy(),
	}

	var totalDelta float64
	var diffPixelCount int

	for y := baseBounds.Min.Y; y < baseBounds.Max.Y; y++ {
		for x := baseBounds.Min.X; x < baseBounds.Max.X; x++ {
			baseColor := baselineImg.At(x, y)
			genColor := generatedImg.At(x, y)

			// Calculate color distance
			delta := colorDistance(
				baseColor,
				genColor,
			)

			if delta > 0 {
				result.DiffPixels++
				totalDelta += delta
				diffPixelCount++

				// Track maximum delta
				if delta > result.MaxColorDelta {
					result.MaxColorDelta = delta
				}

				// Count pixels exceeding tolerance
				if delta > config.PixelTolerance {
					result.PixelsExceedingTolerance++
				}
			}
		}
	}

	// Calculate average delta across differing pixels
	if diffPixelCount > 0 {
		result.AvgColorDelta = totalDelta / float64(
			diffPixelCount,
		)
	}

	return result, nil
}

// colorDistance calculates the Euclidean distance between two colors in RGB space.
// Returns a value from 0 to ~441 (sqrt(255^2 + 255^2 + 255^2)).
func colorDistance(c1, c2 color.Color) float64 {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	// RGBA returns values in range [0, 65535], convert to [0, 255]
	r1, g1, b1 = r1>>8, g1>>8, b1>>8
	r2, g2, b2 = r2>>8, g2>>8, b2>>8

	// Calculate Euclidean distance
	dr := float64(int(r1) - int(r2))
	dg := float64(int(g1) - int(g2))
	db := float64(int(b1) - int(b2))

	return math.Sqrt(dr*dr + dg*dg + db*db)
}

// loadPNG loads a PNG image from the specified path.
func loadPNG(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to open image file: %w",
			err,
		)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to decode PNG image: %w",
			err,
		)
	}

	return img, nil
}
