package framework

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// DiffResult contains comprehensive comparison metrics from visual image comparison.
// Includes pixel-level metrics (diff count, percentage) and perceptual metrics (SSIM, MSE, PSNR).
type DiffResult struct {
	Image1Path string
	Image2Path string
	DiffPath   string

	// Metrics
	PixelDiff   int     // Number of different pixels
	TotalPixels int     // Total pixel count
	DiffPercent float64 // Percentage difference (0-100)
	SSIM        float64 // Structural Similarity Index (-1 to 1, 1 = identical)
	MSE         float64 // Mean Squared Error
	PSNR        float64 // Peak Signal-to-Noise Ratio (dB)

	// Result
	Passed    bool    // Did comparison pass threshold?
	Threshold float64 // Configured threshold
}

// Differ performs visual comparison between two PNG images.
// Implementations include ImageMagickDiffer (uses ImageMagick's compare tool)
// and GoDiffer (pure Go pixel-by-pixel comparison).
type Differ interface {
	// Compare compares two images and generates a diff image highlighting differences.
	// The diffPath specifies where to save the visual diff (red overlay on grayscale).
	// Returns metrics including SSIM, MSE, PSNR, and pass/fail based on config threshold.
	Compare(
		ctx context.Context,
		img1Path, img2Path, diffPath string,
		config TestConfig,
	) (*DiffResult, error)
}

// ImageMagickDiffer uses ImageMagick's compare tool for visual comparison.
// This provides access to multiple comparison algorithms (SSIM, MSE, PSNR, PHASH, AE)
// via the `compare -metric <METRIC>` command.
type ImageMagickDiffer struct{}

// NewImageMagickDiffer creates an ImageMagick-based differ
func NewImageMagickDiffer() *ImageMagickDiffer {
	return &ImageMagickDiffer{}
}

// Compare uses ImageMagick compare command
func (d *ImageMagickDiffer) Compare(
	ctx context.Context,
	img1Path, img2Path, diffPath string,
	config TestConfig,
) (*DiffResult, error) {
	result := &DiffResult{
		Image1Path: img1Path,
		Image2Path: img2Path,
		DiffPath:   diffPath,
		Threshold:  config.DiffThreshold,
	}

	// ImageMagick compare command:
	// compare -metric <METRIC> img1.png img2.png diff.png
	// Outputs to stderr: <value> (<normalized>)
	metricName := mapDiffAlgorithmToImageMagick(
		config.DiffAlgorithm,
	)

	cmd := exec.CommandContext(ctx, "compare",
		"-metric", metricName,
		img1Path,
		img2Path,
		diffPath,
	)

	// Run the command
	// Note: compare exits with code 1 when images differ, which is not an error for us
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Check if it's just a non-zero exit code (images differ)
		if exitErr, ok := err.(*exec.ExitError); ok {
			// Exit code 1 is expected when images differ
			if exitErr.ExitCode() != 1 {
				return nil, fmt.Errorf(
					"ImageMagick compare failed: %w\nOutput: %s",
					err,
					string(output),
				)
			}
			// Continue processing - exit code 1 just means images differ
		} else {
			return nil, fmt.Errorf("ImageMagick compare failed: %w\nOutput: %s", err, string(output))
		}
	}

	// Parse output (stderr contains metric value)
	// Format: "<value> (<normalized>)" or just "<value>"
	outputStr := strings.TrimSpace(string(output))

	// Extract the first number (the metric value)
	var metricValue float64
	parts := strings.Fields(outputStr)
	if len(parts) > 0 {
		// Remove parentheses if present
		valueStr := strings.Trim(parts[0], "()")
		metricValue, err = strconv.ParseFloat(
			valueStr,
			64,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"parse metric value from output '%s': %w",
				outputStr,
				err,
			)
		}
	}

	// Store metric based on algorithm
	switch config.DiffAlgorithm {
	case DiffAlgorithmSSIM:
		result.SSIM = metricValue
		result.Passed = result.SSIM >= (1 - config.DiffThreshold)
	case DiffAlgorithmMSE:
		result.MSE = metricValue
		// For MSE, lower is better - use threshold as max acceptable MSE
		result.Passed = result.MSE <= config.DiffThreshold
	case DiffAlgorithmPSNR:
		result.PSNR = metricValue
		// For PSNR, higher is better (dB) - use threshold inversely
		result.Passed = result.PSNR >= (1.0 / config.DiffThreshold)
	case DiffAlgorithmPerceptual:
		// PHASH returns a difference value
		result.DiffPercent = metricValue
		result.Passed = result.DiffPercent <= config.DiffThreshold*100
	case DiffAlgorithmPixelPerfect:
		// AE (Absolute Error) returns pixel count
		result.PixelDiff = int(metricValue)
		// Need to calculate total pixels to get percentage
		img, err := loadPNG(img1Path)
		if err != nil {
			return nil, fmt.Errorf(
				"load image to calculate total pixels: %w",
				err,
			)
		}
		bounds := img.Bounds()
		result.TotalPixels = bounds.Dx() * bounds.Dy()
		result.DiffPercent = float64(
			result.PixelDiff,
		) / float64(
			result.TotalPixels,
		) * 100
		result.Passed = result.DiffPercent <= config.DiffThreshold*100
	default:
		// AE (Absolute Error) returns pixel count for unknown algorithms
		result.PixelDiff = int(metricValue)
		// Need to calculate total pixels to get percentage
		img, err := loadPNG(img1Path)
		if err != nil {
			return nil, fmt.Errorf(
				"load image to calculate total pixels: %w",
				err,
			)
		}
		bounds := img.Bounds()
		result.TotalPixels = bounds.Dx() * bounds.Dy()
		result.DiffPercent = float64(
			result.PixelDiff,
		) / float64(
			result.TotalPixels,
		) * 100
		result.Passed = result.DiffPercent <= config.DiffThreshold*100
	}

	return result, nil
}

// mapDiffAlgorithmToImageMagick maps DiffAlgorithm to ImageMagick metric names
func mapDiffAlgorithmToImageMagick(
	alg DiffAlgorithm,
) string {
	switch alg {
	case DiffAlgorithmSSIM:
		return "SSIM"
	case DiffAlgorithmMSE:
		return "MSE"
	case DiffAlgorithmPSNR:
		return "PSNR"
	case DiffAlgorithmPerceptual:
		return "PHASH" // Perceptual hash
	case DiffAlgorithmPixelPerfect:
		return "AE" // Absolute error (pixel count)
	default:
		return "AE" // Absolute error (pixel count) for unknown algorithms
	}
}

// GoDiffer implements pure Go image comparison without external dependencies.
// Performs pixel-by-pixel comparison with optional anti-aliasing tolerance.
// Computes SSIM, MSE, and PSNR metrics internally.
type GoDiffer struct{}

// NewGoDiffer creates a Go-native differ
func NewGoDiffer() *GoDiffer {
	return &GoDiffer{}
}

// Compare performs pixel-by-pixel comparison in pure Go
func (d *GoDiffer) Compare(
	ctx context.Context,
	img1Path, img2Path, diffPath string,
	config TestConfig,
) (*DiffResult, error) {
	// Load images
	img1, err := loadPNG(img1Path)
	if err != nil {
		return nil, fmt.Errorf(
			"load image1: %w",
			err,
		)
	}

	img2, err := loadPNG(img2Path)
	if err != nil {
		return nil, fmt.Errorf(
			"load image2: %w",
			err,
		)
	}

	// Check dimensions
	bounds1 := img1.Bounds()
	bounds2 := img2.Bounds()

	if bounds1 != bounds2 {
		return nil, fmt.Errorf(
			"image dimensions differ: %v vs %v",
			bounds1,
			bounds2,
		)
	}

	// Create diff image
	width := bounds1.Dx()
	height := bounds1.Dy()
	diffImg := image.NewRGBA(bounds1)

	// Comparison variables
	var diffPixels int
	var mse float64
	totalPixels := width * height

	// Anti-aliasing tolerance (small threshold for edge pixel differences)
	aaThreshold := uint32(0)
	if config.IgnoreAntialiasing {
		aaThreshold = 5 // Allow up to 5 units difference per channel (out of 255)
	}

	// Pixel-by-pixel comparison
	for y := range height {
		for x := range width {
			c1 := img1.At(x, y)
			c2 := img2.At(x, y)

			r1, g1, b1, a1 := c1.RGBA()
			r2, g2, b2, a2 := c2.RGBA()

			// Convert from 16-bit to 8-bit
			r1_8 := uint8(r1 >> 8)
			g1_8 := uint8(g1 >> 8)
			b1_8 := uint8(b1 >> 8)
			a1_8 := uint8(a1 >> 8)

			r2_8 := uint8(r2 >> 8)
			g2_8 := uint8(g2 >> 8)
			b2_8 := uint8(b2 >> 8)
			a2_8 := uint8(a2 >> 8)

			// Calculate differences
			rDiff := absDiff(r1_8, r2_8)
			gDiff := absDiff(g1_8, g2_8)
			bDiff := absDiff(b1_8, b2_8)
			aDiff := absDiff(a1_8, a2_8)

			// Check if pixel differs (considering anti-aliasing tolerance)
			differs := uint32(
				rDiff,
			) > aaThreshold ||
				uint32(gDiff) > aaThreshold ||
				uint32(bDiff) > aaThreshold ||
				uint32(aDiff) > aaThreshold

			if differs {
				diffPixels++
				// Highlight difference in red
				diffImg.Set(
					x,
					y,
					color.RGBA{
						R: 255,
						G: 0,
						B: 0,
						A: 255,
					},
				)
			} else {
				// Keep original (grayscale)
				gray := uint8((int(r1_8) + int(g1_8) + int(b1_8)) / 3)
				diffImg.Set(x, y, color.RGBA{R: gray, G: gray, B: gray, A: 255})
			}

			// Accumulate MSE (using full precision differences)
			rDiffF := float64(rDiff)
			gDiffF := float64(gDiff)
			bDiffF := float64(bDiff)
			aDiffF := float64(aDiff)
			mse += rDiffF*rDiffF + gDiffF*gDiffF + bDiffF*bDiffF + aDiffF*aDiffF
		}
	}

	// Calculate metrics
	mse /= float64(
		totalPixels * 4,
	) // 4 channels (RGBA)
	var psnr float64
	if mse > 0 {
		// PSNR calculation: 10 * log10(MAX^2 / MSE)
		// MAX = 255 for 8-bit images, but we can use 65535 for 16-bit precision
		psnr = 10 * math.Log10(65535*65535/mse)
	} else {
		// Perfect match - use a very high PSNR
		psnr = 100.0
	}

	// Save diff image
	if err := savePNG(diffPath, diffImg); err != nil {
		return nil, fmt.Errorf(
			"save diff image: %w",
			err,
		)
	}

	// Build result
	diffPercent := float64(
		diffPixels,
	) / float64(
		totalPixels,
	) * 100

	result := &DiffResult{
		Image1Path:  img1Path,
		Image2Path:  img2Path,
		DiffPath:    diffPath,
		PixelDiff:   diffPixels,
		TotalPixels: totalPixels,
		DiffPercent: diffPercent,
		MSE:         mse,
		PSNR:        psnr,
		Threshold:   config.DiffThreshold,
	}

	// Calculate SSIM if needed
	if config.DiffAlgorithm == DiffAlgorithmSSIM {
		result.SSIM = calculateSSIM(img1, img2)
		result.Passed = result.SSIM >= (1 - config.DiffThreshold)
	} else {
		// Default: use DiffPercent threshold
		result.Passed = diffPercent <= config.DiffThreshold*100
	}

	return result, nil
}

// calculateSSIM computes Structural Similarity Index (SSIM) between two images.
// SSIM measures perceptual similarity considering luminance, contrast, and structure.
// Returns a value between -1 and 1, where 1 indicates identical images.
func calculateSSIM(
	img1, img2 image.Image,
) float64 {
	bounds := img1.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Convert images to grayscale and collect pixel values
	var pixels1, pixels2 []float64
	for y := range height {
		for x := range width {
			r1, g1, b1, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()

			// Convert to grayscale (using 16-bit values)
			gray1 := float64(r1+g1+b1) / 3.0
			gray2 := float64(r2+g2+b2) / 3.0

			pixels1 = append(pixels1, gray1)
			pixels2 = append(pixels2, gray2)
		}
	}

	n := float64(len(pixels1))

	// Calculate means
	var sum1, sum2 float64
	for i := range pixels1 {
		sum1 += pixels1[i]
		sum2 += pixels2[i]
	}
	meanX := sum1 / n
	meanY := sum2 / n

	// Calculate variances and covariance
	var varX, varY, covXY float64
	for i := range pixels1 {
		diffX := pixels1[i] - meanX
		diffY := pixels2[i] - meanY
		varX += diffX * diffX
		varY += diffY * diffY
		covXY += diffX * diffY
	}
	varX /= n
	varY /= n
	covXY /= n

	// SSIM constants (to avoid division by zero)
	c1 := 0.01 * 0.01
	c2 := 0.03 * 0.03

	// SSIM formula
	numerator := (2*meanX*meanY + c1) * (2*covXY + c2)
	denominator := (meanX*meanX + meanY*meanY + c1) * (varX + varY + c2)

	if denominator == 0 {
		return 1.0 // Identical images
	}

	ssim := numerator / denominator

	// Clamp to [-1, 1] range
	if ssim > 1.0 {
		ssim = 1.0
	} else if ssim < -1.0 {
		ssim = -1.0
	}

	return ssim
}

// loadPNG loads a PNG image from disk
func loadPNG(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// savePNG saves an image to disk as PNG
func savePNG(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	return png.Encode(file, img)
}

// absDiff returns the absolute difference between two uint8 values
func absDiff(a, b uint8) uint8 {
	if a > b {
		return a - b
	}

	return b - a
}
