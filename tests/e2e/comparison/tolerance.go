// of OOXML documents across different implementations.
package comparison

import "github.com/connerohnesorge/goffice/tests/e2e/framework"

const (
	// DefaultPixelTolerance is the default pixel color difference tolerance
	DefaultPixelTolerance = 2.0
	// DefaultDiffThreshold is the default percentage threshold for visual differences
	DefaultDiffThreshold = 0.005
)

// ToleranceConfig is re-exported from framework for convenience
// This allows comparison package users to access tolerance configuration
// without importing the framework package directly
type ToleranceConfig = framework.ToleranceConfig

// DefaultTolerance returns a default tolerance configuration
func DefaultTolerance() ToleranceConfig {
	return ToleranceConfig{
		XMLAttributeOrderSensitive: false,
		VisualPixelTolerance:       DefaultPixelTolerance,
		VisualDiffThreshold:        DefaultDiffThreshold,
	}
}

// StrictTolerance returns a strict tolerance configuration
// where all differences are considered failures
func StrictTolerance() ToleranceConfig {
	return ToleranceConfig{
		XMLAttributeOrderSensitive: true,
		VisualPixelTolerance:       0.0,
		VisualDiffThreshold:        0.0,
	}
}
