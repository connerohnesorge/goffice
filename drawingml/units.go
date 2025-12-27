package drawingml

import "math"

// EMU represents a value in English Metric Units.
// EMU is the fundamental unit of measurement in DrawingML.
// One EMU equals 1/914400 of an inch or 1/360000 of a centimeter.
type EMU int64

// EMU conversion constants.
// These constants define the relationship between EMUs and other common units.
const (
	// EMUsPerInch is the number of EMUs per inch.
	// This is the base conversion factor defined by the OOXML specification.
	EMUsPerInch EMU = 914400

	// EMUsPerPoint is the number of EMUs per typographic point.
	// One point equals 1/72 of an inch, so EMUsPerPoint = EMUsPerInch / 72.
	EMUsPerPoint EMU = 12700

	// EMUsPerCm is the number of EMUs per centimeter.
	// One centimeter equals 1/2.54 inches.
	EMUsPerCm EMU = 360000

	// EMUsPerMm is the number of EMUs per millimeter.
	// One millimeter equals 1/10 of a centimeter.
	EMUsPerMm EMU = 36000

	// EMUsPerPixel is the number of EMUs per pixel at 96 DPI.
	// This is the standard Windows screen resolution.
	// EMUsPerPixel = EMUsPerInch / 96.
	EMUsPerPixel EMU = 9525

	// EMUsPerTwip is the number of EMUs per twip.
	// One twip equals 1/20 of a point or 1/1440 of an inch.
	EMUsPerTwip EMU = 635

	// EMUsPerEighthPoint is the number of EMUs per eighth of a point.
	// Used for font sizes specified in eighth-points.
	EMUsPerEighthPoint EMU = 1588

	// EMUsPerHalfPoint is the number of EMUs per half-point.
	// Used for font sizes in half-points (common in Word).
	EMUsPerHalfPoint EMU = 6350
)

// defaultDPI is the standard screen resolution used for pixel conversions.
const defaultDPI = 96

// InchesToEmu converts inches to EMUs.
func InchesToEmu(inches float64) EMU {
	return EMU(
		math.Round(inches * float64(EMUsPerInch)),
	)
}

// EmuToInches converts EMUs to inches.
func EmuToInches(emu EMU) float64 {
	return float64(emu) / float64(EMUsPerInch)
}

// CmToEmu converts centimeters to EMUs.
func CmToEmu(cm float64) EMU {
	return EMU(
		math.Round(cm * float64(EMUsPerCm)),
	)
}

// EmuToCm converts EMUs to centimeters.
func EmuToCm(emu EMU) float64 {
	return float64(emu) / float64(EMUsPerCm)
}

// MmToEmu converts millimeters to EMUs.
func MmToEmu(mm float64) EMU {
	return EMU(
		math.Round(mm * float64(EMUsPerMm)),
	)
}

// EmuToMm converts EMUs to millimeters.
func EmuToMm(emu EMU) float64 {
	return float64(emu) / float64(EMUsPerMm)
}

// PointsToEmu converts typographic points to EMUs.
func PointsToEmu(points float64) EMU {
	return EMU(
		math.Round(
			points * float64(EMUsPerPoint),
		),
	)
}

// EmuToPoints converts EMUs to typographic points.
func EmuToPoints(emu EMU) float64 {
	return float64(emu) / float64(EMUsPerPoint)
}

// PixelsToEmu converts pixels (at 96 DPI) to EMUs.
func PixelsToEmu(pixels float64) EMU {
	return EMU(
		math.Round(
			pixels * float64(EMUsPerPixel),
		),
	)
}

// EmuToPixels converts EMUs to pixels (at 96 DPI).
func EmuToPixels(emu EMU) float64 {
	return float64(emu) / float64(EMUsPerPixel)
}

// PixelsToEmuAtDPI converts pixels at a specific DPI to EMUs.
func PixelsToEmuAtDPI(pixels, dpi float64) EMU {
	localDPI := dpi
	if localDPI <= 0 {
		localDPI = defaultDPI
	}
	inches := pixels / localDPI

	return InchesToEmu(inches)
}

// EmuToPixelsAtDPI converts EMUs to pixels at a specific DPI.
func EmuToPixelsAtDPI(
	emu EMU,
	dpi float64,
) float64 {
	localDPI := dpi
	if localDPI <= 0 {
		localDPI = defaultDPI
	}
	inches := EmuToInches(emu)

	return inches * localDPI
}

// TwipsToEmu converts twips to EMUs.
// One twip equals 1/20 of a point.
func TwipsToEmu(twips float64) EMU {
	return EMU(
		math.Round(twips * float64(EMUsPerTwip)),
	)
}

// EmuToTwips converts EMUs to twips.
func EmuToTwips(emu EMU) float64 {
	return float64(emu) / float64(EMUsPerTwip)
}

// HalfPointsToEmu converts half-points to EMUs.
// Half-points are commonly used for font sizes in WordprocessingML.
func HalfPointsToEmu(halfPoints float64) EMU {
	return EMU(
		math.Round(
			halfPoints * float64(
				EMUsPerHalfPoint,
			),
		),
	)
}

// EmuToHalfPoints converts EMUs to half-points.
func EmuToHalfPoints(emu EMU) float64 {
	return float64(
		emu,
	) / float64(
		EMUsPerHalfPoint,
	)
}

// Int64 returns the EMU value as an int64.
func (e EMU) Int64() int64 {
	return int64(e)
}

// Float64 returns the EMU value as a float64.
func (e EMU) Float64() float64 {
	return float64(e)
}

// ToInches converts this EMU value to inches.
func (e EMU) ToInches() float64 {
	return EmuToInches(e)
}

// ToCm converts this EMU value to centimeters.
func (e EMU) ToCm() float64 {
	return EmuToCm(e)
}

// ToMm converts this EMU value to millimeters.
func (e EMU) ToMm() float64 {
	return EmuToMm(e)
}

// ToPoints converts this EMU value to points.
func (e EMU) ToPoints() float64 {
	return EmuToPoints(e)
}

// ToPixels converts this EMU value to pixels at 96 DPI.
func (e EMU) ToPixels() float64 {
	return EmuToPixels(e)
}

// ToPixelsAtDPI converts this EMU value to pixels at a specific DPI.
func (e EMU) ToPixelsAtDPI(dpi float64) float64 {
	return EmuToPixelsAtDPI(e, dpi)
}

// ToTwips converts this EMU value to twips.
func (e EMU) ToTwips() float64 {
	return EmuToTwips(e)
}

// Angle constants for rotation values.
// DrawingML uses 60,000ths of a degree for rotation angles.
const (
	// DegreesPerFullRotation is 360 degrees.
	DegreesPerFullRotation = 360

	// AngleUnitsPerDegree is the number of angle units per degree.
	// DrawingML uses 60,000ths of a degree.
	AngleUnitsPerDegree = 60000

	// AngleUnitsPerFullRotation is the number of angle units
	// in a full rotation.
	AngleUnitsPerFullRotation = DegreesPerFullRotation * AngleUnitsPerDegree
)

// DegreesToAngleUnits converts degrees to DrawingML angle units
// (60,000ths of a degree).
func DegreesToAngleUnits(degrees float64) int {
	return int(
		math.Round(
			degrees * float64(
				AngleUnitsPerDegree,
			),
		),
	)
}

// AngleUnitsToDegrees converts DrawingML angle units to degrees.
func AngleUnitsToDegrees(units int) float64 {
	return float64(
		units,
	) / float64(
		AngleUnitsPerDegree,
	)
}

// Percentage constants.
// DrawingML uses 1000ths of a percent for percentage values.
const (
	// PercentageUnitsPerPercent is the number of units per percent.
	// DrawingML uses 1000ths of a percent.
	PercentageUnitsPerPercent = 1000

	// PercentageUnitsFor100Percent is the value for 100%.
	PercentageUnitsFor100Percent = 100 * PercentageUnitsPerPercent
)

// PercentToUnits converts a percentage (0-100) to DrawingML percentage units.
func PercentToUnits(percent float64) int {
	return int(
		math.Round(
			percent * float64(
				PercentageUnitsPerPercent,
			),
		),
	)
}

// UnitsToPercent converts DrawingML percentage units to a percentage (0-100).
func UnitsToPercent(units int) float64 {
	return float64(
		units,
	) / float64(
		PercentageUnitsPerPercent,
	)
}
