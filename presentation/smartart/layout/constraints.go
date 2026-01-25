package layout

const (
	// EMUPerInch is the number of English Metric Units (EMU) in one inch
	EMUPerInch = 914400

	// Default spacing and sizing constants (in EMU)
	defaultSpacingDivisor = 4 // Divides EMUPerInch for default spacing (0.25 inch)
	defaultHeightDivisor  = 2 // Divides EMUPerInch for default min height (0.5 inch)
)

// Constraints control layout behavior
type Constraints struct {
	Spacing   int64     // Gap between shapes (EMU)
	Padding   int64     // Inner padding from bounds (EMU)
	MinWidth  int64     // Minimum shape width
	MinHeight int64     // Minimum shape height
	Alignment Alignment
}

// DefaultConstraints returns sensible defaults
func DefaultConstraints() Constraints {
	return Constraints{
		Spacing:   EMUPerInch / defaultSpacingDivisor, // 0.25 inch
		Padding:   EMUPerInch / defaultSpacingDivisor,
		MinWidth:  EMUPerInch,                         // 1 inch
		MinHeight: EMUPerInch / defaultHeightDivisor,  // 0.5 inch
		Alignment: AlignCenter,
	}
}
