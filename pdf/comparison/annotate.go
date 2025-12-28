package comparison

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// AnnotateDiffImage creates a visual diff image highlighting pixel differences.
//
// The output image shows:
//   - Pixels exceeding tolerance are highlighted in red
//   - Pixels within tolerance are shown from the baseline
//
// Parameters:
//   - baselinePath: path to the baseline (reference) image
//   - generatedPath: path to the generated (test) image
//   - outputPath: path to save the annotated diff image
//   - config: comparison configuration (nil for defaults)
//
// Returns an error if the operation fails.
func AnnotateDiffImage(
	baselinePath, generatedPath, outputPath string,
	config *ComparisonConfig,
) error {
	if config == nil {
		config = DefaultComparisonConfig()
	}

	// Load baseline image
	baselineImg, err := loadPNG(baselinePath)
	if err != nil {
		return fmt.Errorf(
			"failed to load baseline image: %w",
			err,
		)
	}

	// Load generated image
	generatedImg, err := loadPNG(generatedPath)
	if err != nil {
		return fmt.Errorf(
			"failed to load generated image: %w",
			err,
		)
	}

	// Check dimensions match
	baseBounds := baselineImg.Bounds()
	genBounds := generatedImg.Bounds()

	if baseBounds.Dx() != genBounds.Dx() ||
		baseBounds.Dy() != genBounds.Dy() {
		return fmt.Errorf(
			"image dimensions do not match: baseline %dx%d, generated %dx%d",
			baseBounds.Dx(),
			baseBounds.Dy(),
			genBounds.Dx(),
			genBounds.Dy(),
		)
	}

	// Create output image
	diffImg := image.NewRGBA(baseBounds)

	// Highlight color for differences (bright red with some transparency effect)
	highlightColor := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}

	// Process each pixel
	for y := baseBounds.Min.Y; y < baseBounds.Max.Y; y++ {
		for x := baseBounds.Min.X; x < baseBounds.Max.X; x++ {
			baseColor := baselineImg.At(x, y)
			genColor := generatedImg.At(x, y)

			// Calculate color distance
			delta := colorDistance(
				baseColor,
				genColor,
			)

			if delta > config.PixelTolerance {
				// Highlight pixels exceeding tolerance in red
				diffImg.Set(x, y, highlightColor)
			} else {
				// Show baseline pixel for pixels within tolerance
				diffImg.Set(x, y, baseColor)
			}
		}
	}

	// Save the diff image
	if err := savePNG(diffImg, outputPath); err != nil {
		return fmt.Errorf(
			"failed to save diff image: %w",
			err,
		)
	}

	return nil
}

// CreateSideBySideDiff creates a side-by-side comparison image showing
// baseline, generated, and diff images.
//
// The output is a horizontal concatenation: [baseline | generated | diff]
func CreateSideBySideDiff(
	baselinePath, generatedPath, diffPath, outputPath string,
) error {
	// Load all three images
	baselineImg, err := loadPNG(baselinePath)
	if err != nil {
		return fmt.Errorf(
			"failed to load baseline image: %w",
			err,
		)
	}

	generatedImg, err := loadPNG(generatedPath)
	if err != nil {
		return fmt.Errorf(
			"failed to load generated image: %w",
			err,
		)
	}

	diffImg, err := loadPNG(diffPath)
	if err != nil {
		return fmt.Errorf(
			"failed to load diff image: %w",
			err,
		)
	}

	baseBounds := baselineImg.Bounds()
	width := baseBounds.Dx()
	height := baseBounds.Dy()

	// Create output image (3x width for side-by-side)
	sideBySide := image.NewRGBA(
		image.Rect(0, 0, width*3, height),
	)

	// Copy baseline to left third
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			sideBySide.Set(
				x,
				y,
				baselineImg.At(
					x+baseBounds.Min.X,
					y+baseBounds.Min.Y,
				),
			)
		}
	}

	// Copy generated to middle third
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			sideBySide.Set(
				x+width,
				y,
				generatedImg.At(
					x+baseBounds.Min.X,
					y+baseBounds.Min.Y,
				),
			)
		}
	}

	// Copy diff to right third
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			sideBySide.Set(
				x+width*2,
				y,
				diffImg.At(
					x+baseBounds.Min.X,
					y+baseBounds.Min.Y,
				),
			)
		}
	}

	// Save the side-by-side image
	if err := savePNG(sideBySide, outputPath); err != nil {
		return fmt.Errorf(
			"failed to save side-by-side diff: %w",
			err,
		)
	}

	return nil
}

// savePNG saves an image as a PNG file.
func savePNG(img image.Image, path string) error {
	// Ensure directory exists
	dir := path[:len(path)-len(path[len(path)-1:])]
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			dir = path[:i]
			break
		}
	}

	if dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf(
				"failed to create directory: %w",
				err,
			)
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf(
			"failed to create file: %w",
			err,
		)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf(
			"failed to encode PNG: %w",
			err,
		)
	}

	return nil
}
