// Package core provides PDF document creation and low-level PDF primitives.
package core

import (
	"math"
)

// OOXML-specific unit conversion constants.
const (
	// EMUPerMM is the number of EMU per millimeter.
	// 914400 EMU = 1 inch, 1 inch = 25.4 mm
	// Therefore: 914400 / 25.4 = 36000 EMU per mm
	EMUPerMM = EMUPerInch / 25.4

	// EMUPerCM is the number of EMU per centimeter.
	// 914400 EMU = 1 inch, 1 inch = 2.54 cm
	// Therefore: 914400 / 2.54 = 360000 EMU per cm
	EMUPerCM = EMUPerInch / 2.54

	// EMUPerTwip is the number of EMU per twip.
	// 914400 EMU = 1 inch, 1440 twips = 1 inch
	// Therefore: 914400 / 1440 = 635 EMU per twip
	EMUPerTwip = EMUPerInch / 1440.0

	// TwipsPerEMU is the number of twips per EMU.
	TwipsPerEMU = 1440.0 / EMUPerInch

	// OOXMLAngleDegree is the number of OOXML angle units in one degree.
	OOXMLAngleDegree = 60000.0

	// OOXMLAngleUnit represents 1/60000th of a degree (used in DrawingML).
	OOXMLAngleUnit = 1.0 / OOXMLAngleDegree

	// OOXMLPercentageFull is the number of OOXML percentage units in 100%.
	OOXMLPercentageFull = 100000.0

	// OOXMLPercentageUnit represents 1/100000th (used in DrawingML percentages).
	OOXMLPercentageUnit = 1.0 / OOXMLPercentageFull

	// OOXMLPercentagePoint is the number of OOXML percentage units in 1%.
	OOXMLPercentagePoint = 1000.0

	// DegreesInHalfCircle is the number of degrees in a half circle.
	DegreesInHalfCircle = 180.0

	// DegreesInFullCircle is the number of degrees in a full circle.
	DegreesInFullCircle = 360.0
)

// EMUValue represents a value in English Metric Units.
// EMU is the coordinate system used by OOXML documents where
// 914400 EMU = 1 inch.
type EMUValue int64

// NewEMU creates a new EMUValue from an int64.
func NewEMU(value int64) EMUValue {
	return EMUValue(value)
}

// EMUFromPoints converts PDF points to EMU.
func EMUFromPoints(points float64) EMUValue {
	return EMUValue(
		math.Round(points * EMUPerPoint),
	)
}

// EMUFromInches converts inches to EMU.
func EMUFromInches(inches float64) EMUValue {
	return EMUValue(
		math.Round(inches * EMUPerInch),
	)
}

// EMUFromMM converts millimeters to EMU.
func EMUFromMM(mm float64) EMUValue {
	return EMUValue(math.Round(mm * EMUPerMM))
}

// EMUFromCM converts centimeters to EMU.
func EMUFromCM(cm float64) EMUValue {
	return EMUValue(math.Round(cm * EMUPerCM))
}

// EMUFromTwips converts twips to EMU.
func EMUFromTwips(twips float64) EMUValue {
	return EMUValue(
		math.Round(twips * EMUPerTwip),
	)
}

// Int64 returns the EMU value as an int64.
func (e EMUValue) Int64() int64 {
	return int64(e)
}

// Float64 returns the EMU value as a float64.
func (e EMUValue) Float64() float64 {
	return float64(e)
}

// ToPoints converts EMU to PDF points.
func (e EMUValue) ToPoints() float64 {
	return float64(e) * PointsPerEMU
}

// ToInches converts EMU to inches.
func (e EMUValue) ToInches() float64 {
	return float64(e) / EMUPerInch
}

// ToMM converts EMU to millimeters.
func (e EMUValue) ToMM() float64 {
	return float64(e) / EMUPerMM
}

// ToCM converts EMU to centimeters.
func (e EMUValue) ToCM() float64 {
	return float64(e) / EMUPerCM
}

// ToTwips converts EMU to twips.
func (e EMUValue) ToTwips() float64 {
	return float64(e) / EMUPerTwip
}

// Add adds two EMU values.
func (e EMUValue) Add(other EMUValue) EMUValue {
	return e + other
}

// Sub subtracts an EMU value from this one.
func (e EMUValue) Sub(other EMUValue) EMUValue {
	return e - other
}

// Mul multiplies the EMU value by a scalar.
func (e EMUValue) Mul(scalar float64) EMUValue {
	return EMUValue(
		math.Round(float64(e) * scalar),
	)
}

// Div divides the EMU value by a scalar.
func (e EMUValue) Div(scalar float64) EMUValue {
	if scalar == 0 {
		return 0
	}

	return EMUValue(
		math.Round(float64(e) / scalar),
	)
}

// IsPositive returns true if the EMU value is greater than zero.
func (e EMUValue) IsPositive() bool {
	return e > 0
}

// IsNonNegative returns true if the EMU value is greater than or equal to zero.
func (e EMUValue) IsNonNegative() bool {
	return e >= 0
}

// STCoordinate represents an ST_Coordinate value from DrawingML.
// ST_Coordinate is a signed EMU value used for positioning.
type STCoordinate = EMUValue

// STPositiveCoordinate represents an ST_PositiveCoordinate value from DrawingML.
// ST_PositiveCoordinate is a non-negative EMU value used for dimensions.
type STPositiveCoordinate = EMUValue

// ParseSTCoordinate parses an ST_Coordinate value (EMU as int64).
func ParseSTCoordinate(value int64) STCoordinate {
	return STCoordinate(value)
}

// ParseSTPositiveCoordinate parses an ST_PositiveCoordinate value.
// Returns 0 if the value is negative.
func ParseSTPositiveCoordinate(
	value int64,
) STPositiveCoordinate {
	if value < 0 {
		return 0
	}

	return STPositiveCoordinate(value)
}

// OOXMLAngle represents an angle in OOXML format (60000ths of a degree).
// OOXML stores angles as integers where 60000 = 1 degree.
type OOXMLAngle int32

// NewOOXMLAngle creates a new OOXMLAngle from an int32.
func NewOOXMLAngle(value int32) OOXMLAngle {
	return OOXMLAngle(value)
}

// AngleFromDegrees creates an OOXMLAngle from degrees.
func AngleFromDegrees(
	degrees float64,
) OOXMLAngle {
	return OOXMLAngle(
		math.Round(degrees * OOXMLAngleDegree),
	)
}

// AngleFromRadians creates an OOXMLAngle from radians.
func AngleFromRadians(
	radians float64,
) OOXMLAngle {
	degrees := radians * DegreesInHalfCircle / math.Pi

	return AngleFromDegrees(degrees)
}

// Int32 returns the raw OOXML angle value.
func (a OOXMLAngle) Int32() int32 {
	return int32(a)
}

// ToDegrees converts the OOXML angle to degrees.
func (a OOXMLAngle) ToDegrees() float64 {
	return float64(a) / OOXMLAngleDegree
}

// ToRadians converts the OOXML angle to radians.
func (a OOXMLAngle) ToRadians() float64 {
	return a.ToDegrees() * math.Pi / DegreesInHalfCircle
}

// Normalize normalizes the angle to the range [0, DegreesInFullCircle) degrees.
func (a OOXMLAngle) Normalize() OOXMLAngle {
	degrees := a.ToDegrees()
	// Handle negative angles
	for degrees < 0 {
		degrees += DegreesInFullCircle
	}
	// Handle angles >= DegreesInFullCircle
	for degrees >= DegreesInFullCircle {
		degrees -= DegreesInFullCircle
	}

	return AngleFromDegrees(degrees)
}

// Add adds two OOXML angles.
func (a OOXMLAngle) Add(
	other OOXMLAngle,
) OOXMLAngle {
	return a + other
}

// Sub subtracts an OOXML angle from this one.
func (a OOXMLAngle) Sub(
	other OOXMLAngle,
) OOXMLAngle {
	return a - other
}

// Negate returns the negation of this angle.
func (a OOXMLAngle) Negate() OOXMLAngle {
	return -a
}

// OOXMLPercentage represents a percentage in OOXML format (OOXMLPercentageFullths).
// OOXML stores percentages as integers where OOXMLPercentageFull = 100% = 1.0.
type OOXMLPercentage int32

// NewOOXMLPercentage creates a new OOXMLPercentage from an int32.
func NewOOXMLPercentage(
	value int32,
) OOXMLPercentage {
	return OOXMLPercentage(value)
}

// PercentageFromFloat creates an OOXMLPercentage from a 0-1 float value.
// For example, 0.5 becomes 50000 (50%).
func PercentageFromFloat(
	value float64,
) OOXMLPercentage {
	return OOXMLPercentage(
		math.Round(value * OOXMLPercentageFull),
	)
}

// PercentageFromPercent creates an OOXMLPercentage from a 0-100 percent value.
// For example, 50 becomes 50000 (50%).
func PercentageFromPercent(
	percent float64,
) OOXMLPercentage {
	return OOXMLPercentage(
		math.Round(
			percent * OOXMLPercentagePoint,
		),
	)
}

// Int32 returns the raw OOXML percentage value.
func (p OOXMLPercentage) Int32() int32 {
	return int32(p)
}

// ToFloat converts the OOXML percentage to a 0-1 float value.
// For example, 50000 becomes 0.5.
func (p OOXMLPercentage) ToFloat() float64 {
	return float64(p) / OOXMLPercentageFull
}

// ToPercent converts the OOXML percentage to a 0-100 percent value.
// For example, 50000 becomes 50.0.
func (p OOXMLPercentage) ToPercent() float64 {
	return float64(p) / OOXMLPercentagePoint
}

// Clamp clamps the percentage to the range [0, OOXMLPercentageFull] (0-100%).
func (p OOXMLPercentage) Clamp() OOXMLPercentage {
	if p < 0 {
		return 0
	}
	if p > OOXMLPercentageFull {
		return OOXMLPercentageFull
	}

	return p
}

// Add adds two OOXML percentages.
func (p OOXMLPercentage) Add(
	other OOXMLPercentage,
) OOXMLPercentage {
	return p + other
}

// Sub subtracts an OOXML percentage from this one.
func (p OOXMLPercentage) Sub(
	other OOXMLPercentage,
) OOXMLPercentage {
	return p - other
}

// Mul multiplies two OOXML percentages.
// The result is the product of the two percentages as ratios.
// For example, 50000 (50%) * 50000 (50%) = 25000 (25%).
func (p OOXMLPercentage) Mul(
	other OOXMLPercentage,
) OOXMLPercentage {
	// (p/OOXMLPercentageFull) * (other/OOXMLPercentageFull) * OOXMLPercentageFull = p * other / OOXMLPercentageFull
	return OOXMLPercentage(
		int64(
			p,
		) * int64(
			other,
		) / OOXMLPercentageFull,
	)
}

// Scale scales a value by this percentage.
// For example, if p is 50000 (50%), Scale(100.0) returns 50.0.
func (p OOXMLPercentage) Scale(
	value float64,
) float64 {
	return value * p.ToFloat()
}

// Additional conversion utilities for DrawingML

// EMUToInches converts English Metric Units to inches.
func EMUToInches(emu float64) float64 {
	return emu / EMUPerInch
}

// InchesToEMU converts inches to English Metric Units.
func InchesToEMU(inches float64) float64 {
	return inches * EMUPerInch
}

// EMUToMM converts English Metric Units to millimeters.
func EMUToMM(emu float64) float64 {
	return emu / EMUPerMM
}

// MMToEMU converts millimeters to English Metric Units.
func MMToEMU(mm float64) float64 {
	return mm * EMUPerMM
}

// EMUToCM converts English Metric Units to centimeters.
func EMUToCM(emu float64) float64 {
	return emu / EMUPerCM
}

// CMToEMU converts centimeters to English Metric Units.
func CMToEMU(cm float64) float64 {
	return cm * EMUPerCM
}

// TwipsToEMU converts twips to English Metric Units.
func TwipsToEMU(twips float64) float64 {
	return twips * EMUPerTwip
}

// EMUToTwips converts English Metric Units to twips.
func EMUToTwips(emu float64) float64 {
	return emu / EMUPerTwip
}

// OOXMLAngleToDegrees converts an OOXML angle value (60000ths of a degree) to degrees.
func OOXMLAngleToDegrees(angle int32) float64 {
	return float64(angle) / OOXMLAngleDegree
}

// DegreesToOOXMLAngle converts degrees to an OOXML angle value (60000ths of a degree).
func DegreesToOOXMLAngle(degrees float64) int32 {
	return int32(
		math.Round(degrees * OOXMLAngleDegree),
	)
}

// OOXMLAngleToRadians converts an OOXML angle value to radians.
func OOXMLAngleToRadians(angle int32) float64 {
	return OOXMLAngleToDegrees(
		angle,
	) * math.Pi / DegreesInHalfCircle
}

// RadiansToOOXMLAngle converts radians to an OOXML angle value.
func RadiansToOOXMLAngle(radians float64) int32 {
	degrees := radians * DegreesInHalfCircle / math.Pi

	return DegreesToOOXMLAngle(degrees)
}

// OOXMLPercentToFloat converts an OOXML percentage (OOXMLPercentageFullths) to a 0-1 float.
func OOXMLPercentToFloat(percent int32) float64 {
	return float64(percent) / OOXMLPercentageFull
}

// FloatToOOXMLPercent converts a 0-1 float to an OOXML percentage (OOXMLPercentageFullths).
func FloatToOOXMLPercent(value float64) int32 {
	return int32(
		math.Round(value * OOXMLPercentageFull),
	)
}
