package comparison

import "fmt"

// DiffResult contains the results of comparing two images.
type DiffResult struct {
	// TotalPixels is the total number of pixels compared
	TotalPixels int

	// DiffPixels is the number of pixels that differ
	DiffPixels int

	// PixelsExceedingTolerance is the number of pixels with differences
	// exceeding the configured tolerance threshold
	PixelsExceedingTolerance int

	// MaxColorDelta is the maximum color distance found
	MaxColorDelta float64

	// AvgColorDelta is the average color distance across differing pixels
	AvgColorDelta float64
}

// DiffPercentage returns the percentage of pixels that differ.
func (r *DiffResult) DiffPercentage() float64 {
	if r.TotalPixels == 0 {
		return 0
	}
	return float64(
		r.DiffPixels,
	) / float64(
		r.TotalPixels,
	) * 100
}

// ExceedingTolerance returns the percentage of pixels exceeding tolerance.
func (r *DiffResult) ExceedingTolerance() float64 {
	if r.TotalPixels == 0 {
		return 0
	}
	return float64(
		r.PixelsExceedingTolerance,
	) / float64(
		r.TotalPixels,
	) * 100
}

// IsWithinTolerance returns true if the difference is within the image
// threshold (0.5% of pixels allowed to exceed tolerance).
func (r *DiffResult) IsWithinTolerance() bool {
	return r.ExceedingTolerance() <= imageDiffThreshold*100
}

// String returns a human-readable summary of the diff result.
func (r *DiffResult) String() string {
	return fmt.Sprintf(
		"Diff: %.2f%% pixels differ, %.2f%% exceed tolerance (max delta: %.2f, avg delta: %.2f)",
		r.DiffPercentage(),
		r.ExceedingTolerance(),
		r.MaxColorDelta,
		r.AvgColorDelta,
	)
}

// ComparisonConfig contains configuration for visual comparison.
type ComparisonConfig struct {
	// PixelTolerance is the maximum allowed color distance per pixel
	// (0-255 per channel, Euclidean distance in RGB space)
	PixelTolerance float64

	// ImageDiffThreshold is the maximum allowed percentage of pixels
	// exceeding tolerance (0.0-1.0, e.g., 0.005 = 0.5%)
	ImageDiffThreshold float64

	// DPI for PDF-to-PNG conversion (default: 150)
	DPI int

	// OutputDiffPath is the path to save annotated diff images
	OutputDiffPath string
}

// DefaultComparisonConfig returns the default comparison configuration.
func DefaultComparisonConfig() *ComparisonConfig {
	return &ComparisonConfig{
		PixelTolerance:     pixelTolerance,
		ImageDiffThreshold: imageDiffThreshold,
		DPI:                150,
	}
}

const (
	// pixelTolerance is the per-pixel color distance tolerance
	// in RGB Euclidean space (0-255 per channel).
	// A value of 2.0 means ~1% color variation is tolerated.
	pixelTolerance = 2.0

	// imageDiffThreshold is the per-image threshold for pixels
	// exceeding tolerance. 0.005 = 0.5% of pixels allowed to differ.
	imageDiffThreshold = 0.005
)
