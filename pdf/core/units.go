// Package core provides PDF document creation and low-level PDF primitives.
package core

import (
	"math"
	"strconv"
)

// Unit conversion constants.
// PDF uses points as its native unit, where 72 points = 1 inch.
const (
	// PointsPerInch is the number of PDF points per inch.
	PointsPerInch = 72.0

	// PointsPerMM is the number of PDF points per millimeter.
	PointsPerMM = PointsPerInch / 25.4

	// PointsPerCM is the number of PDF points per centimeter.
	PointsPerCM = PointsPerInch / 2.54

	// PointsPerEMU is the number of PDF points per EMU (English Metric Unit).
	// 1 inch = 914400 EMU, 1 inch = 72 points
	// Therefore: 1 point = 914400/72 = 12700 EMU
	PointsPerEMU = 72.0 / 914400.0

	// EMUPerInch is the number of EMU per inch (OOXML standard).
	EMUPerInch = 914400.0

	// EMUPerPoint is the number of EMU per PDF point.
	EMUPerPoint = EMUPerInch / PointsPerInch

	// PointsPerTwip is the number of PDF points per twip.
	// 1 inch = 1440 twips, 1 inch = 72 points
	// Therefore: 1 point = 1440/72 = 20 twips, so 1 twip = 1/20 point
	PointsPerTwip = 72.0 / 1440.0

	// TwipsPerPoint is the number of twips per PDF point.
	TwipsPerPoint = 1440.0 / 72.0
)

// Unit represents a unit of measurement for PDF dimensions.
type Unit int

const (
	// UnitPoint represents PDF points (1/72 inch).
	UnitPoint Unit = iota
	// UnitInch represents inches.
	UnitInch
	// UnitMM represents millimeters.
	UnitMM
	// UnitCM represents centimeters.
	UnitCM
	// UnitEMU represents English Metric Units (used in OOXML).
	UnitEMU
	// UnitTwip represents twips (1/20 of a point, used in OOXML).
	UnitTwip
)

// String returns the string representation of the unit.
func (u Unit) String() string {
	switch u {
	case UnitPoint:
		return "pt"
	case UnitInch:
		return "in"
	case UnitMM:
		return "mm"
	case UnitCM:
		return "cm"
	case UnitEMU:
		return "emu"
	case UnitTwip:
		return "twip"
	default:
		return "pt"
	}
}

// ToPoints converts a value from this unit to PDF points.
func (u Unit) ToPoints(value float64) float64 {
	switch u {
	case UnitPoint:
		return value
	case UnitInch:
		return InchesToPoints(value)
	case UnitMM:
		return MMToPoints(value)
	case UnitCM:
		return CMToPoints(value)
	case UnitEMU:
		return EMUToPoints(value)
	case UnitTwip:
		return TwipsToPoints(value)
	default:
		return value
	}
}

// FromPoints converts a value from PDF points to this unit.
func (u Unit) FromPoints(points float64) float64 {
	switch u {
	case UnitPoint:
		return points
	case UnitInch:
		return PointsToInches(points)
	case UnitMM:
		return PointsToMM(points)
	case UnitCM:
		return PointsToCM(points)
	case UnitEMU:
		return PointsToEMU(points)
	case UnitTwip:
		return PointsToTwips(points)
	default:
		return points
	}
}

// InchesToPoints converts inches to PDF points.
func InchesToPoints(inches float64) float64 {
	return inches * PointsPerInch
}

// PointsToInches converts PDF points to inches.
func PointsToInches(points float64) float64 {
	return points / PointsPerInch
}

// MMToPoints converts millimeters to PDF points.
func MMToPoints(mm float64) float64 {
	return mm * PointsPerMM
}

// PointsToMM converts PDF points to millimeters.
func PointsToMM(points float64) float64 {
	return points / PointsPerMM
}

// CMToPoints converts centimeters to PDF points.
func CMToPoints(cm float64) float64 {
	return cm * PointsPerCM
}

// PointsToCM converts PDF points to centimeters.
func PointsToCM(points float64) float64 {
	return points / PointsPerCM
}

// EMUToPoints converts English Metric Units to PDF points.
// EMU is used in OOXML documents (914400 EMU = 1 inch).
func EMUToPoints(emu float64) float64 {
	return emu * PointsPerEMU
}

// PointsToEMU converts PDF points to English Metric Units.
func PointsToEMU(points float64) float64 {
	return points * EMUPerPoint
}

// TwipsToPoints converts twips to PDF points.
// Twips are used in OOXML documents (1440 twips = 1 inch, 20 twips = 1 point).
func TwipsToPoints(twips float64) float64 {
	return twips * PointsPerTwip
}

// PointsToTwips converts PDF points to twips.
func PointsToTwips(points float64) float64 {
	return points * TwipsPerPoint
}

// Constants for float formatting.
const (
	base10            = 10   // Base 10 for math and parsing
	bitSize64         = 64   // Bit size for float64
	truncatePrecision = 1000 // Truncate to 3 decimal places
)

// RoundToDecimalPlaces rounds a float64 to the specified number of decimal places.
// This is useful for ensuring consistent precision in PDF output.
func RoundToDecimalPlaces(
	value float64,
	places int,
) float64 {
	multiplier := math.Pow(
		base10,
		float64(places),
	)

	return math.Round(
		value*multiplier,
	) / multiplier
}

// Dimension represents a value with an associated unit.
type Dimension struct {
	Value float64
	Unit  Unit
}

// NewDimension creates a new Dimension with the given value and unit.
func NewDimension(
	value float64,
	unit Unit,
) Dimension {
	return Dimension{Value: value, Unit: unit}
}

// Points creates a Dimension in PDF points.
func Points(value float64) Dimension {
	return Dimension{
		Value: value,
		Unit:  UnitPoint,
	}
}

// Inches creates a Dimension in inches.
func Inches(value float64) Dimension {
	return Dimension{Value: value, Unit: UnitInch}
}

// MM creates a Dimension in millimeters.
func MM(value float64) Dimension {
	return Dimension{Value: value, Unit: UnitMM}
}

// CM creates a Dimension in centimeters.
func CM(value float64) Dimension {
	return Dimension{Value: value, Unit: UnitCM}
}

// EMU creates a Dimension in English Metric Units.
func EMU(value float64) Dimension {
	return Dimension{Value: value, Unit: UnitEMU}
}

// Twips creates a Dimension in twips.
func Twips(value float64) Dimension {
	return Dimension{Value: value, Unit: UnitTwip}
}

// ToPoints converts the dimension to PDF points.
func (d Dimension) ToPoints() float64 {
	return d.Unit.ToPoints(d.Value)
}

// Convert converts the dimension to another unit.
func (d Dimension) Convert(to Unit) Dimension {
	points := d.ToPoints()

	return Dimension{
		Value: to.FromPoints(points),
		Unit:  to,
	}
}

// Add adds two dimensions and returns the result in points.
func (d Dimension) Add(
	other Dimension,
) Dimension {
	return Points(d.ToPoints() + other.ToPoints())
}

// Sub subtracts another dimension and returns the result in points.
func (d Dimension) Sub(
	other Dimension,
) Dimension {
	return Points(d.ToPoints() - other.ToPoints())
}

// Mul multiplies the dimension by a scalar.
func (d Dimension) Mul(scalar float64) Dimension {
	return Dimension{
		Value: d.Value * scalar,
		Unit:  d.Unit,
	}
}

// String returns a string representation of the dimension.
func (d Dimension) String() string {
	return FormatFloat(d.Value) + d.Unit.String()
}

// FormatFloat formats a float64 for display, removing trailing zeros.
func FormatFloat(f float64) string {
	s := math.Trunc(
		f*truncatePrecision,
	) / truncatePrecision // Truncate to 3 decimal places
	if s == math.Trunc(s) {
		return strconv.FormatFloat(
			s,
			'f',
			0,
			bitSize64,
		)
	}

	return strconv.FormatFloat(
		s,
		'f',
		-1,
		bitSize64,
	)
}
