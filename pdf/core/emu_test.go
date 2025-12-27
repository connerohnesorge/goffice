package core

import (
	"math"
	"testing"
)

// Tolerance for floating-point comparisons
const floatTolerance = 0.0001

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < floatTolerance
}

func TestEMUConstants(t *testing.T) {
	// Verify EMU per mm: 914400 / 25.4 = 36000
	expectedEMUPerMM := 36000.0
	if !floatEquals(EMUPerMM, expectedEMUPerMM) {
		t.Errorf(
			"EMUPerMM = %v, want %v",
			EMUPerMM,
			expectedEMUPerMM,
		)
	}

	// Verify EMU per cm: 914400 / 2.54 = 360000
	expectedEMUPerCM := 360000.0
	if !floatEquals(EMUPerCM, expectedEMUPerCM) {
		t.Errorf(
			"EMUPerCM = %v, want %v",
			EMUPerCM,
			expectedEMUPerCM,
		)
	}

	// Verify EMU per twip: 914400 / 1440 = 635
	expectedEMUPerTwip := 635.0
	if !floatEquals(
		EMUPerTwip,
		expectedEMUPerTwip,
	) {
		t.Errorf(
			"EMUPerTwip = %v, want %v",
			EMUPerTwip,
			expectedEMUPerTwip,
		)
	}
}

func TestEMUValue_Creation(t *testing.T) {
	tests := []struct {
		name  string
		value int64
	}{
		{"zero", 0},
		{"positive", 914400},
		{"negative", -914400},
		{"large", 9144000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emu := NewEMU(tt.value)
			if emu.Int64() != tt.value {
				t.Errorf(
					"NewEMU(%d).Int64() = %d, want %d",
					tt.value,
					emu.Int64(),
					tt.value,
				)
			}
			if emu.Float64() != float64(
				tt.value,
			) {
				t.Errorf(
					"NewEMU(%d).Float64() = %f, want %f",
					tt.value,
					emu.Float64(),
					float64(tt.value),
				)
			}
		})
	}
}

func TestEMUFromPoints(t *testing.T) {
	tests := []struct {
		points float64
		want   int64
	}{
		{72.0, 914400},   // 1 inch
		{36.0, 457200},   // 0.5 inch
		{144.0, 1828800}, // 2 inches
		{1.0, 12700},     // 1 point = 12700 EMU
	}

	for _, tt := range tests {
		got := EMUFromPoints(tt.points)
		if got.Int64() != tt.want {
			t.Errorf(
				"EMUFromPoints(%v) = %v, want %v",
				tt.points,
				got.Int64(),
				tt.want,
			)
		}
	}
}

func TestEMUFromInches(t *testing.T) {
	tests := []struct {
		inches float64
		want   int64
	}{
		{1.0, 914400},
		{0.5, 457200},
		{2.0, 1828800},
	}

	for _, tt := range tests {
		got := EMUFromInches(tt.inches)
		if got.Int64() != tt.want {
			t.Errorf(
				"EMUFromInches(%v) = %v, want %v",
				tt.inches,
				got.Int64(),
				tt.want,
			)
		}
	}
}

func TestEMUFromMM(t *testing.T) {
	tests := []struct {
		mm   float64
		want int64
	}{
		{25.4, 914400}, // 1 inch
		{1.0, 36000},   // 1 mm
		{10.0, 360000}, // 1 cm
	}

	for _, tt := range tests {
		got := EMUFromMM(tt.mm)
		if got.Int64() != tt.want {
			t.Errorf(
				"EMUFromMM(%v) = %v, want %v",
				tt.mm,
				got.Int64(),
				tt.want,
			)
		}
	}
}

func TestEMUFromCM(t *testing.T) {
	tests := []struct {
		cm   float64
		want int64
	}{
		{2.54, 914400},  // 1 inch
		{1.0, 360000},   // 1 cm
		{10.0, 3600000}, // 10 cm
	}

	for _, tt := range tests {
		got := EMUFromCM(tt.cm)
		if got.Int64() != tt.want {
			t.Errorf(
				"EMUFromCM(%v) = %v, want %v",
				tt.cm,
				got.Int64(),
				tt.want,
			)
		}
	}
}

func TestEMUFromTwips(t *testing.T) {
	tests := []struct {
		twips float64
		want  int64
	}{
		{1440.0, 914400},  // 1 inch
		{20.0, 12700},     // 1 point
		{2880.0, 1828800}, // 2 inches
	}

	for _, tt := range tests {
		got := EMUFromTwips(tt.twips)
		if got.Int64() != tt.want {
			t.Errorf(
				"EMUFromTwips(%v) = %v, want %v",
				tt.twips,
				got.Int64(),
				tt.want,
			)
		}
	}
}

func TestEMUValue_ToPoints(t *testing.T) {
	tests := []struct {
		emu  int64
		want float64
	}{
		{914400, 72.0},
		{457200, 36.0},
		{12700, 1.0},
	}

	for _, tt := range tests {
		got := NewEMU(tt.emu).ToPoints()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"EMU(%d).ToPoints() = %v, want %v",
				tt.emu,
				got,
				tt.want,
			)
		}
	}
}

func TestEMUValue_ToInches(t *testing.T) {
	tests := []struct {
		emu  int64
		want float64
	}{
		{914400, 1.0},
		{457200, 0.5},
		{1828800, 2.0},
	}

	for _, tt := range tests {
		got := NewEMU(tt.emu).ToInches()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"EMU(%d).ToInches() = %v, want %v",
				tt.emu,
				got,
				tt.want,
			)
		}
	}
}

func TestEMUValue_ToMM(t *testing.T) {
	tests := []struct {
		emu  int64
		want float64
	}{
		{914400, 25.4},
		{36000, 1.0},
		{360000, 10.0},
	}

	for _, tt := range tests {
		got := NewEMU(tt.emu).ToMM()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"EMU(%d).ToMM() = %v, want %v",
				tt.emu,
				got,
				tt.want,
			)
		}
	}
}

func TestEMUValue_ToCM(t *testing.T) {
	tests := []struct {
		emu  int64
		want float64
	}{
		{914400, 2.54},
		{360000, 1.0},
		{3600000, 10.0},
	}

	for _, tt := range tests {
		got := NewEMU(tt.emu).ToCM()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"EMU(%d).ToCM() = %v, want %v",
				tt.emu,
				got,
				tt.want,
			)
		}
	}
}

func TestEMUValue_ToTwips(t *testing.T) {
	tests := []struct {
		emu  int64
		want float64
	}{
		{914400, 1440.0},
		{12700, 20.0},
		{1828800, 2880.0},
	}

	for _, tt := range tests {
		got := NewEMU(tt.emu).ToTwips()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"EMU(%d).ToTwips() = %v, want %v",
				tt.emu,
				got,
				tt.want,
			)
		}
	}
}

func TestEMUValue_Arithmetic(t *testing.T) {
	a := NewEMU(914400)
	b := NewEMU(457200)

	// Test Add
	sum := a.Add(b)
	if sum.Int64() != 1371600 {
		t.Errorf(
			"EMU.Add: got %d, want 1371600",
			sum.Int64(),
		)
	}

	// Test Sub
	diff := a.Sub(b)
	if diff.Int64() != 457200 {
		t.Errorf(
			"EMU.Sub: got %d, want 457200",
			diff.Int64(),
		)
	}

	// Test Mul
	product := a.Mul(2.0)
	if product.Int64() != 1828800 {
		t.Errorf(
			"EMU.Mul: got %d, want 1828800",
			product.Int64(),
		)
	}

	// Test Div
	quotient := a.Div(2.0)
	if quotient.Int64() != 457200 {
		t.Errorf(
			"EMU.Div: got %d, want 457200",
			quotient.Int64(),
		)
	}

	// Test Div by zero
	zero := a.Div(0)
	if zero.Int64() != 0 {
		t.Errorf(
			"EMU.Div(0): got %d, want 0",
			zero.Int64(),
		)
	}
}

func TestEMUValue_IsPositive(t *testing.T) {
	tests := []struct {
		emu  int64
		want bool
	}{
		{914400, true},
		{0, false},
		{-914400, false},
	}

	for _, tt := range tests {
		got := NewEMU(tt.emu).IsPositive()
		if got != tt.want {
			t.Errorf(
				"EMU(%d).IsPositive() = %v, want %v",
				tt.emu,
				got,
				tt.want,
			)
		}
	}
}

func TestEMUValue_IsNonNegative(t *testing.T) {
	tests := []struct {
		emu  int64
		want bool
	}{
		{914400, true},
		{0, true},
		{-914400, false},
	}

	for _, tt := range tests {
		got := NewEMU(tt.emu).IsNonNegative()
		if got != tt.want {
			t.Errorf(
				"EMU(%d).IsNonNegative() = %v, want %v",
				tt.emu,
				got,
				tt.want,
			)
		}
	}
}

func TestParseSTCoordinate(t *testing.T) {
	tests := []struct {
		value int64
		want  int64
	}{
		{914400, 914400},
		{-914400, -914400},
		{0, 0},
	}

	for _, tt := range tests {
		got := ParseSTCoordinate(tt.value)
		if got.Int64() != tt.want {
			t.Errorf(
				"ParseSTCoordinate(%d) = %d, want %d",
				tt.value,
				got.Int64(),
				tt.want,
			)
		}
	}
}

func TestParseSTPositiveCoordinate(t *testing.T) {
	tests := []struct {
		value int64
		want  int64
	}{
		{914400, 914400},
		{-914400, 0}, // Negative becomes 0
		{0, 0},
	}

	for _, tt := range tests {
		got := ParseSTPositiveCoordinate(tt.value)
		if got.Int64() != tt.want {
			t.Errorf(
				"ParseSTPositiveCoordinate(%d) = %d, want %d",
				tt.value,
				got.Int64(),
				tt.want,
			)
		}
	}
}

func TestOOXMLAngle_Creation(t *testing.T) {
	angle := NewOOXMLAngle(5400000) // 90 degrees
	if angle.Int32() != 5400000 {
		t.Errorf(
			"NewOOXMLAngle(5400000).Int32() = %d, want 5400000",
			angle.Int32(),
		)
	}
}

func TestAngleFromDegrees(t *testing.T) {
	tests := []struct {
		degrees float64
		want    int32
	}{
		{90.0, 5400000},
		{180.0, 10800000},
		{45.0, 2700000},
		{0.0, 0},
		{-90.0, -5400000},
	}

	for _, tt := range tests {
		got := AngleFromDegrees(tt.degrees)
		if got.Int32() != tt.want {
			t.Errorf(
				"AngleFromDegrees(%v) = %d, want %d",
				tt.degrees,
				got.Int32(),
				tt.want,
			)
		}
	}
}

func TestAngleFromRadians(t *testing.T) {
	tests := []struct {
		radians float64
		want    float64 // expected degrees
	}{
		{math.Pi / 2, 90.0},
		{math.Pi, 180.0},
		{math.Pi / 4, 45.0},
		{0.0, 0.0},
	}

	for _, tt := range tests {
		got := AngleFromRadians(
			tt.radians,
		).ToDegrees()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"AngleFromRadians(%v).ToDegrees() = %v, want %v",
				tt.radians,
				got,
				tt.want,
			)
		}
	}
}

func TestOOXMLAngle_ToDegrees(t *testing.T) {
	tests := []struct {
		angle int32
		want  float64
	}{
		{5400000, 90.0},
		{10800000, 180.0},
		{21600000, 360.0},
		{0, 0.0},
		{-5400000, -90.0},
	}

	for _, tt := range tests {
		got := NewOOXMLAngle(tt.angle).ToDegrees()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"OOXMLAngle(%d).ToDegrees() = %v, want %v",
				tt.angle,
				got,
				tt.want,
			)
		}
	}
}

func TestOOXMLAngle_ToRadians(t *testing.T) {
	tests := []struct {
		angle int32
		want  float64
	}{
		{5400000, math.Pi / 2},
		{10800000, math.Pi},
		{21600000, 2 * math.Pi},
	}

	for _, tt := range tests {
		got := NewOOXMLAngle(tt.angle).ToRadians()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"OOXMLAngle(%d).ToRadians() = %v, want %v",
				tt.angle,
				got,
				tt.want,
			)
		}
	}
}

func TestOOXMLAngle_Normalize(t *testing.T) {
	tests := []struct {
		angle int32
		want  float64 // expected degrees after normalization
	}{
		{5400000, 90.0},
		{21600000, 0.0},   // 360 -> 0
		{27000000, 90.0},  // 450 -> 90
		{-5400000, 270.0}, // -90 -> 270
		{-21600000, 0.0},  // -360 -> 0
	}

	for _, tt := range tests {
		got := NewOOXMLAngle(
			tt.angle,
		).Normalize().
			ToDegrees()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"OOXMLAngle(%d).Normalize().ToDegrees() = %v, want %v",
				tt.angle,
				got,
				tt.want,
			)
		}
	}
}

func TestOOXMLAngle_Arithmetic(t *testing.T) {
	a := AngleFromDegrees(90) // 5400000
	b := AngleFromDegrees(45) // 2700000

	// Test Add
	sum := a.Add(b)
	if !floatEquals(sum.ToDegrees(), 135.0) {
		t.Errorf(
			"Angle.Add: got %v degrees, want 135",
			sum.ToDegrees(),
		)
	}

	// Test Sub
	diff := a.Sub(b)
	if !floatEquals(diff.ToDegrees(), 45.0) {
		t.Errorf(
			"Angle.Sub: got %v degrees, want 45",
			diff.ToDegrees(),
		)
	}

	// Test Negate
	neg := a.Negate()
	if !floatEquals(neg.ToDegrees(), -90.0) {
		t.Errorf(
			"Angle.Negate: got %v degrees, want -90",
			neg.ToDegrees(),
		)
	}
}

func TestOOXMLPercentage_Creation(t *testing.T) {
	pct := NewOOXMLPercentage(50000) // 50%
	if pct.Int32() != 50000 {
		t.Errorf(
			"NewOOXMLPercentage(50000).Int32() = %d, want 50000",
			pct.Int32(),
		)
	}
}

func TestPercentageFromFloat(t *testing.T) {
	tests := []struct {
		value float64
		want  int32
	}{
		{0.5, 50000},  // 50%
		{1.0, 100000}, // 100%
		{0.0, 0},      // 0%
		{0.25, 25000}, // 25%
		{1.5, 150000}, // 150%
	}

	for _, tt := range tests {
		got := PercentageFromFloat(tt.value)
		if got.Int32() != tt.want {
			t.Errorf(
				"PercentageFromFloat(%v) = %d, want %d",
				tt.value,
				got.Int32(),
				tt.want,
			)
		}
	}
}

func TestPercentageFromPercent(t *testing.T) {
	tests := []struct {
		percent float64
		want    int32
	}{
		{50.0, 50000},   // 50%
		{100.0, 100000}, // 100%
		{0.0, 0},        // 0%
		{25.0, 25000},   // 25%
	}

	for _, tt := range tests {
		got := PercentageFromPercent(tt.percent)
		if got.Int32() != tt.want {
			t.Errorf(
				"PercentageFromPercent(%v) = %d, want %d",
				tt.percent,
				got.Int32(),
				tt.want,
			)
		}
	}
}

func TestOOXMLPercentage_ToFloat(t *testing.T) {
	tests := []struct {
		pct  int32
		want float64
	}{
		{50000, 0.5},
		{100000, 1.0},
		{0, 0.0},
		{25000, 0.25},
	}

	for _, tt := range tests {
		got := NewOOXMLPercentage(
			tt.pct,
		).ToFloat()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"OOXMLPercentage(%d).ToFloat() = %v, want %v",
				tt.pct,
				got,
				tt.want,
			)
		}
	}
}

func TestOOXMLPercentage_ToPercent(t *testing.T) {
	tests := []struct {
		pct  int32
		want float64
	}{
		{50000, 50.0},
		{100000, 100.0},
		{0, 0.0},
		{25000, 25.0},
	}

	for _, tt := range tests {
		got := NewOOXMLPercentage(
			tt.pct,
		).ToPercent()
		if !floatEquals(got, tt.want) {
			t.Errorf(
				"OOXMLPercentage(%d).ToPercent() = %v, want %v",
				tt.pct,
				got,
				tt.want,
			)
		}
	}
}

func TestOOXMLPercentage_Clamp(t *testing.T) {
	tests := []struct {
		pct  int32
		want int32
	}{
		{50000, 50000},   // Within range
		{0, 0},           // At lower bound
		{100000, 100000}, // At upper bound
		{-10000, 0},      // Below range
		{150000, 100000}, // Above range
	}

	for _, tt := range tests {
		got := NewOOXMLPercentage(tt.pct).Clamp()
		if got.Int32() != tt.want {
			t.Errorf(
				"OOXMLPercentage(%d).Clamp() = %d, want %d",
				tt.pct,
				got.Int32(),
				tt.want,
			)
		}
	}
}

func TestOOXMLPercentage_Arithmetic(
	t *testing.T,
) {
	a := PercentageFromPercent(50) // 50%
	b := PercentageFromPercent(25) // 25%

	// Test Add
	sum := a.Add(b)
	if !floatEquals(sum.ToPercent(), 75.0) {
		t.Errorf(
			"Percentage.Add: got %v%%, want 75%%",
			sum.ToPercent(),
		)
	}

	// Test Sub
	diff := a.Sub(b)
	if !floatEquals(diff.ToPercent(), 25.0) {
		t.Errorf(
			"Percentage.Sub: got %v%%, want 25%%",
			diff.ToPercent(),
		)
	}

	// Test Mul
	product := a.Mul(b) // 50% * 25% = 12.5%
	if !floatEquals(product.ToPercent(), 12.5) {
		t.Errorf(
			"Percentage.Mul: got %v%%, want 12.5%%",
			product.ToPercent(),
		)
	}

	// Test Scale
	scaled := a.Scale(100.0) // 50% of 100 = 50
	if !floatEquals(scaled, 50.0) {
		t.Errorf(
			"Percentage.Scale: got %v, want 50",
			scaled,
		)
	}
}

func TestEMUConversionFunctions(t *testing.T) {
	// Test EMUToInches
	if !floatEquals(EMUToInches(914400), 1.0) {
		t.Errorf(
			"EMUToInches(914400) = %v, want 1.0",
			EMUToInches(914400),
		)
	}

	// Test InchesToEMU
	if !floatEquals(InchesToEMU(1.0), 914400) {
		t.Errorf(
			"InchesToEMU(1.0) = %v, want 914400",
			InchesToEMU(1.0),
		)
	}

	// Test EMUToMM
	if !floatEquals(EMUToMM(36000), 1.0) {
		t.Errorf(
			"EMUToMM(36000) = %v, want 1.0",
			EMUToMM(36000),
		)
	}

	// Test MMToEMU
	if !floatEquals(MMToEMU(1.0), 36000) {
		t.Errorf(
			"MMToEMU(1.0) = %v, want 36000",
			MMToEMU(1.0),
		)
	}

	// Test EMUToCM
	if !floatEquals(EMUToCM(360000), 1.0) {
		t.Errorf(
			"EMUToCM(360000) = %v, want 1.0",
			EMUToCM(360000),
		)
	}

	// Test CMToEMU
	if !floatEquals(CMToEMU(1.0), 360000) {
		t.Errorf(
			"CMToEMU(1.0) = %v, want 360000",
			CMToEMU(1.0),
		)
	}

	// Test TwipsToEMU
	if !floatEquals(TwipsToEMU(1440), 914400) {
		t.Errorf(
			"TwipsToEMU(1440) = %v, want 914400",
			TwipsToEMU(1440),
		)
	}

	// Test EMUToTwips
	if !floatEquals(EMUToTwips(914400), 1440) {
		t.Errorf(
			"EMUToTwips(914400) = %v, want 1440",
			EMUToTwips(914400),
		)
	}
}

func TestAngleConversionFunctions(t *testing.T) {
	// Test OOXMLAngleToDegrees
	if !floatEquals(
		OOXMLAngleToDegrees(5400000),
		90.0,
	) {
		t.Errorf(
			"OOXMLAngleToDegrees(5400000) = %v, want 90.0",
			OOXMLAngleToDegrees(5400000),
		)
	}

	// Test DegreesToOOXMLAngle
	if DegreesToOOXMLAngle(90.0) != 5400000 {
		t.Errorf(
			"DegreesToOOXMLAngle(90.0) = %d, want 5400000",
			DegreesToOOXMLAngle(90.0),
		)
	}

	// Test OOXMLAngleToRadians
	if !floatEquals(
		OOXMLAngleToRadians(5400000),
		math.Pi/2,
	) {
		t.Errorf(
			"OOXMLAngleToRadians(5400000) = %v, want %v",
			OOXMLAngleToRadians(5400000),
			math.Pi/2,
		)
	}

	// Test RadiansToOOXMLAngle
	if RadiansToOOXMLAngle(math.Pi/2) != 5400000 {
		t.Errorf(
			"RadiansToOOXMLAngle(pi/2) = %d, want 5400000",
			RadiansToOOXMLAngle(math.Pi/2),
		)
	}
}

func TestPercentageConversionFunctions(
	t *testing.T,
) {
	// Test OOXMLPercentToFloat
	if !floatEquals(
		OOXMLPercentToFloat(50000),
		0.5,
	) {
		t.Errorf(
			"OOXMLPercentToFloat(50000) = %v, want 0.5",
			OOXMLPercentToFloat(50000),
		)
	}

	// Test FloatToOOXMLPercent
	if FloatToOOXMLPercent(0.5) != 50000 {
		t.Errorf(
			"FloatToOOXMLPercent(0.5) = %d, want 50000",
			FloatToOOXMLPercent(0.5),
		)
	}
}

func TestEMURoundTrip(t *testing.T) {
	// Test that conversions round-trip correctly
	original := int64(914400) // 1 inch

	// EMU -> Points -> EMU
	points := NewEMU(original).ToPoints()
	backToEMU := EMUFromPoints(points)
	if backToEMU.Int64() != original {
		t.Errorf(
			"Points round-trip: got %d, want %d",
			backToEMU.Int64(),
			original,
		)
	}

	// EMU -> Inches -> EMU
	inches := NewEMU(original).ToInches()
	backToEMU = EMUFromInches(inches)
	if backToEMU.Int64() != original {
		t.Errorf(
			"Inches round-trip: got %d, want %d",
			backToEMU.Int64(),
			original,
		)
	}

	// EMU -> MM -> EMU
	mm := NewEMU(original).ToMM()
	backToEMU = EMUFromMM(mm)
	if backToEMU.Int64() != original {
		t.Errorf(
			"MM round-trip: got %d, want %d",
			backToEMU.Int64(),
			original,
		)
	}

	// EMU -> CM -> EMU
	cm := NewEMU(original).ToCM()
	backToEMU = EMUFromCM(cm)
	if backToEMU.Int64() != original {
		t.Errorf(
			"CM round-trip: got %d, want %d",
			backToEMU.Int64(),
			original,
		)
	}

	// EMU -> Twips -> EMU
	twips := NewEMU(original).ToTwips()
	backToEMU = EMUFromTwips(twips)
	if backToEMU.Int64() != original {
		t.Errorf(
			"Twips round-trip: got %d, want %d",
			backToEMU.Int64(),
			original,
		)
	}
}

func TestAngleRoundTrip(t *testing.T) {
	original := int32(5400000) // 90 degrees

	// OOXML -> Degrees -> OOXML
	degrees := NewOOXMLAngle(original).ToDegrees()
	backToOOXML := AngleFromDegrees(degrees)
	if backToOOXML.Int32() != original {
		t.Errorf(
			"Degrees round-trip: got %d, want %d",
			backToOOXML.Int32(),
			original,
		)
	}

	// OOXML -> Radians -> OOXML
	radians := NewOOXMLAngle(original).ToRadians()
	backToOOXML = AngleFromRadians(radians)
	if backToOOXML.Int32() != original {
		t.Errorf(
			"Radians round-trip: got %d, want %d",
			backToOOXML.Int32(),
			original,
		)
	}
}

func TestPercentageRoundTrip(t *testing.T) {
	original := int32(50000) // 50%

	// OOXML -> Float -> OOXML
	f := NewOOXMLPercentage(original).ToFloat()
	backToOOXML := PercentageFromFloat(f)
	if backToOOXML.Int32() != original {
		t.Errorf(
			"Float round-trip: got %d, want %d",
			backToOOXML.Int32(),
			original,
		)
	}

	// OOXML -> Percent -> OOXML
	pct := NewOOXMLPercentage(
		original,
	).ToPercent()
	backToOOXML = PercentageFromPercent(pct)
	if backToOOXML.Int32() != original {
		t.Errorf(
			"Percent round-trip: got %d, want %d",
			backToOOXML.Int32(),
			original,
		)
	}
}

// TestOOXMLAngleEdgeCases tests OOXML angle conversion edge cases.
func TestOOXMLAngleEdgeCases(t *testing.T) {
	t.Run("boundary angles", func(t *testing.T) {
		// 0 degrees
		angle0 := AngleFromDegrees(0)
		if angle0.Int32() != 0 {
			t.Errorf(
				"0 degrees = %d OOXML units, want 0",
				angle0.Int32(),
			)
		}

		// 360 degrees (full rotation)
		angle360 := AngleFromDegrees(360)
		if angle360.Int32() != 21600000 {
			t.Errorf(
				"360 degrees = %d OOXML units, want 21600000",
				angle360.Int32(),
			)
		}

		// 1 degree (smallest common unit)
		angle1 := AngleFromDegrees(1)
		if angle1.Int32() != 60000 {
			t.Errorf(
				"1 degree = %d OOXML units, want 60000",
				angle1.Int32(),
			)
		}
	})

	t.Run("negative angles", func(t *testing.T) {
		// -90 degrees
		angleNeg90 := AngleFromDegrees(-90)
		if !floatEquals(
			angleNeg90.ToDegrees(),
			-90,
		) {
			t.Errorf(
				"-90 degrees roundtrip = %v, want -90",
				angleNeg90.ToDegrees(),
			)
		}

		// -360 degrees
		angleNeg360 := AngleFromDegrees(-360)
		if !floatEquals(
			angleNeg360.ToDegrees(),
			-360,
		) {
			t.Errorf(
				"-360 degrees roundtrip = %v, want -360",
				angleNeg360.ToDegrees(),
			)
		}

		// Normalize negative angles
		normalized := angleNeg90.Normalize()
		if !floatEquals(
			normalized.ToDegrees(),
			270,
		) {
			t.Errorf(
				"-90 normalized = %v degrees, want 270",
				normalized.ToDegrees(),
			)
		}
	})

	t.Run(
		"angles greater than 360",
		func(t *testing.T) {
			// 450 degrees (360 + 90)
			angle450 := AngleFromDegrees(450)
			normalized := angle450.Normalize()
			if !floatEquals(
				normalized.ToDegrees(),
				90,
			) {
				t.Errorf(
					"450 normalized = %v degrees, want 90",
					normalized.ToDegrees(),
				)
			}

			// 720 degrees (2 full rotations)
			angle720 := AngleFromDegrees(720)
			normalized = angle720.Normalize()
			if !floatEquals(
				normalized.ToDegrees(),
				0,
			) {
				t.Errorf(
					"720 normalized = %v degrees, want 0",
					normalized.ToDegrees(),
				)
			}

			// 1080 degrees (3 full rotations)
			angle1080 := AngleFromDegrees(1080)
			normalized = angle1080.Normalize()
			if !floatEquals(
				normalized.ToDegrees(),
				0,
			) {
				t.Errorf(
					"1080 normalized = %v degrees, want 0",
					normalized.ToDegrees(),
				)
			}
		},
	)

	t.Run(
		"fractional angles",
		func(t *testing.T) {
			// 0.5 degrees
			angle05 := AngleFromDegrees(0.5)
			if angle05.Int32() != 30000 {
				t.Errorf(
					"0.5 degrees = %d OOXML units, want 30000",
					angle05.Int32(),
				)
			}

			// 45.5 degrees
			angle455 := AngleFromDegrees(45.5)
			expected := int32(45.5 * 60000)
			if angle455.Int32() != expected {
				t.Errorf(
					"45.5 degrees = %d OOXML units, want %d",
					angle455.Int32(),
					expected,
				)
			}
		},
	)

	t.Run(
		"radians conversion",
		func(t *testing.T) {
			// Test various radian values
			testCases := []struct {
				radians  float64
				expected float64 // degrees
			}{
				{0, 0},
				{
					math.Pi / 6,
					30,
				}, // 30 degrees
				{
					math.Pi / 4,
					45,
				}, // 45 degrees
				{
					math.Pi / 3,
					60,
				}, // 60 degrees
				{
					math.Pi / 2,
					90,
				}, // 90 degrees
				{
					2 * math.Pi / 3,
					120,
				}, // 120 degrees
				{
					math.Pi,
					180,
				}, // 180 degrees
				{
					3 * math.Pi / 2,
					270,
				}, // 270 degrees
				{
					2 * math.Pi,
					360,
				}, // 360 degrees
				{
					-math.Pi / 2,
					-90,
				}, // -90 degrees
			}

			for _, tc := range testCases {
				angle := AngleFromRadians(
					tc.radians,
				)
				degrees := angle.ToDegrees()
				if !floatEquals(
					degrees,
					tc.expected,
				) {
					t.Errorf(
						"AngleFromRadians(%v) = %v degrees, want %v",
						tc.radians,
						degrees,
						tc.expected,
					)
				}
			}
		},
	)
}

// TestOOXMLPercentageEdgeCases tests OOXML percentage conversion edge cases.
func TestOOXMLPercentageEdgeCases(t *testing.T) {
	t.Run(
		"boundary percentages",
		func(t *testing.T) {
			// 0%
			pct0 := PercentageFromFloat(0)
			if pct0.Int32() != 0 {
				t.Errorf(
					"0%% = %d OOXML units, want 0",
					pct0.Int32(),
				)
			}

			// 100%
			pct100 := PercentageFromFloat(1.0)
			if pct100.Int32() != 100000 {
				t.Errorf(
					"100%% = %d OOXML units, want 100000",
					pct100.Int32(),
				)
			}

			// 1%
			pct1 := PercentageFromFloat(0.01)
			if pct1.Int32() != 1000 {
				t.Errorf(
					"1%% = %d OOXML units, want 1000",
					pct1.Int32(),
				)
			}
		},
	)

	t.Run(
		"percentages greater than 100",
		func(t *testing.T) {
			// 150%
			pct150 := PercentageFromFloat(1.5)
			if pct150.Int32() != 150000 {
				t.Errorf(
					"150%% = %d OOXML units, want 150000",
					pct150.Int32(),
				)
			}
			if !floatEquals(
				pct150.ToFloat(),
				1.5,
			) {
				t.Errorf(
					"150%% roundtrip = %v, want 1.5",
					pct150.ToFloat(),
				)
			}

			// 200%
			pct200 := PercentageFromFloat(2.0)
			if !floatEquals(
				pct200.ToFloat(),
				2.0,
			) {
				t.Errorf(
					"200%% roundtrip = %v, want 2.0",
					pct200.ToFloat(),
				)
			}
		},
	)

	t.Run(
		"negative percentages",
		func(t *testing.T) {
			// -50%
			pctNeg50 := PercentageFromFloat(-0.5)
			if pctNeg50.Int32() != -50000 {
				t.Errorf(
					"-50%% = %d OOXML units, want -50000",
					pctNeg50.Int32(),
				)
			}
			if !floatEquals(
				pctNeg50.ToFloat(),
				-0.5,
			) {
				t.Errorf(
					"-50%% roundtrip = %v, want -0.5",
					pctNeg50.ToFloat(),
				)
			}
		},
	)

	t.Run("clamp behavior", func(t *testing.T) {
		// Values below 0 should clamp to 0
		pctNeg := NewOOXMLPercentage(-50000)
		clamped := pctNeg.Clamp()
		if clamped.Int32() != 0 {
			t.Errorf(
				"Clamp(-50%%) = %d, want 0",
				clamped.Int32(),
			)
		}

		// Values above 100% should clamp to 100%
		pctOver := NewOOXMLPercentage(150000)
		clamped = pctOver.Clamp()
		if clamped.Int32() != 100000 {
			t.Errorf(
				"Clamp(150%%) = %d, want 100000",
				clamped.Int32(),
			)
		}

		// Values in range should not change
		pctNormal := NewOOXMLPercentage(50000)
		clamped = pctNormal.Clamp()
		if clamped.Int32() != 50000 {
			t.Errorf(
				"Clamp(50%%) = %d, want 50000",
				clamped.Int32(),
			)
		}
	})

	t.Run(
		"very small percentages",
		func(t *testing.T) {
			// 0.1%
			pct01 := PercentageFromFloat(0.001)
			if pct01.Int32() != 100 {
				t.Errorf(
					"0.1%% = %d OOXML units, want 100",
					pct01.Int32(),
				)
			}

			// 0.01%
			pct001 := PercentageFromFloat(0.0001)
			if pct001.Int32() != 10 {
				t.Errorf(
					"0.01%% = %d OOXML units, want 10",
					pct001.Int32(),
				)
			}

			// 0.001%
			pct0001 := PercentageFromFloat(
				0.00001,
			)
			if pct0001.Int32() != 1 {
				t.Errorf(
					"0.001%% = %d OOXML units, want 1",
					pct0001.Int32(),
				)
			}
		},
	)

	t.Run("scale operation", func(t *testing.T) {
		// 50% of 200 = 100
		pct50 := PercentageFromPercent(50)
		scaled := pct50.Scale(200)
		if !floatEquals(scaled, 100) {
			t.Errorf(
				"50%% of 200 = %v, want 100",
				scaled,
			)
		}

		// 25% of 400 = 100
		pct25 := PercentageFromPercent(25)
		scaled = pct25.Scale(400)
		if !floatEquals(scaled, 100) {
			t.Errorf(
				"25%% of 400 = %v, want 100",
				scaled,
			)
		}

		// 150% of 100 = 150
		pct150 := PercentageFromPercent(150)
		scaled = pct150.Scale(100)
		if !floatEquals(scaled, 150) {
			t.Errorf(
				"150%% of 100 = %v, want 150",
				scaled,
			)
		}
	})
}

// TestEMUValueEdgeCases tests EMU value edge cases.
func TestEMUValueEdgeCases(t *testing.T) {
	t.Run(
		"very large EMU values",
		func(t *testing.T) {
			// Max int32 as EMU (about 168 inches)
			maxInt32 := int64(2147483647)
			emu := NewEMU(maxInt32)
			points := emu.ToPoints()
			expectedPoints := float64(
				maxInt32,
			) * PointsPerEMU

			if math.Abs(
				points-expectedPoints,
			) > 0.001 {
				t.Errorf(
					"Max int32 EMU -> points: got %v, want %v",
					points,
					expectedPoints,
				)
			}
		},
	)

	t.Run(
		"negative EMU values",
		func(t *testing.T) {
			// Negative EMU (for relative positioning)
			negEMU := NewEMU(-914400)
			points := negEMU.ToPoints()
			if !floatEquals(points, -72) {
				t.Errorf(
					"Negative EMU -> points: got %v, want -72",
					points,
				)
			}

			// IsPositive should return false
			if negEMU.IsPositive() {
				t.Error(
					"Negative EMU should not be positive",
				)
			}

			// IsNonNegative should return false
			if negEMU.IsNonNegative() {
				t.Error(
					"Negative EMU should not be non-negative",
				)
			}
		},
	)

	t.Run("zero EMU", func(t *testing.T) {
		zeroEMU := NewEMU(0)

		if zeroEMU.ToPoints() != 0 {
			t.Error(
				"Zero EMU -> points should be 0",
			)
		}
		if zeroEMU.ToInches() != 0 {
			t.Error(
				"Zero EMU -> inches should be 0",
			)
		}
		if zeroEMU.ToMM() != 0 {
			t.Error("Zero EMU -> mm should be 0")
		}
		if zeroEMU.ToCM() != 0 {
			t.Error("Zero EMU -> cm should be 0")
		}
		if zeroEMU.ToTwips() != 0 {
			t.Error(
				"Zero EMU -> twips should be 0",
			)
		}
		if zeroEMU.IsPositive() {
			t.Error(
				"Zero EMU should not be positive",
			)
		}
		if !zeroEMU.IsNonNegative() {
			t.Error(
				"Zero EMU should be non-negative",
			)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		emu := NewEMU(914400)
		result := emu.Div(0)
		if result.Int64() != 0 {
			t.Errorf(
				"EMU / 0 = %d, want 0",
				result.Int64(),
			)
		}
	})

	t.Run(
		"EMU arithmetic precision",
		func(t *testing.T) {
			// Test that multiplication doesn't lose precision
			emu := NewEMU(914400) // 1 inch
			scaled := emu.Mul(0.5)
			if scaled.Int64() != 457200 {
				t.Errorf(
					"EMU * 0.5 = %d, want 457200",
					scaled.Int64(),
				)
			}

			// Test fractional scaling
			scaled = emu.Mul(0.333333)
			expected := int64(
				math.Round(914400 * 0.333333),
			)
			if scaled.Int64() != expected {
				t.Errorf(
					"EMU * 0.333333 = %d, want %d",
					scaled.Int64(),
					expected,
				)
			}
		},
	)
}

// TestSTCoordinateTypes tests ST_Coordinate and ST_PositiveCoordinate types.
func TestSTCoordinateTypes(t *testing.T) {
	t.Run(
		"ST_Coordinate allows negative",
		func(t *testing.T) {
			coord := ParseSTCoordinate(-914400)
			if coord.Int64() != -914400 {
				t.Errorf(
					"ParseSTCoordinate(-914400) = %d, want -914400",
					coord.Int64(),
				)
			}
		},
	)

	t.Run(
		"ST_PositiveCoordinate clamps negative to zero",
		func(t *testing.T) {
			coord := ParseSTPositiveCoordinate(
				-914400,
			)
			if coord.Int64() != 0 {
				t.Errorf(
					"ParseSTPositiveCoordinate(-914400) = %d, want 0",
					coord.Int64(),
				)
			}

			// Positive values should pass through
			coord = ParseSTPositiveCoordinate(
				914400,
			)
			if coord.Int64() != 914400 {
				t.Errorf(
					"ParseSTPositiveCoordinate(914400) = %d, want 914400",
					coord.Int64(),
				)
			}

			// Zero should pass through
			coord = ParseSTPositiveCoordinate(0)
			if coord.Int64() != 0 {
				t.Errorf(
					"ParseSTPositiveCoordinate(0) = %d, want 0",
					coord.Int64(),
				)
			}
		},
	)
}

// TestRealWorldOOXMLValues tests with real-world OOXML values from actual documents.
func TestRealWorldOOXMLValues(t *testing.T) {
	t.Run(
		"PowerPoint slide dimensions",
		func(t *testing.T) {
			// PowerPoint default slide size (10" x 7.5")
			slideWidthEMU := int64(
				9144000,
			) // 10 inches
			slideHeightEMU := int64(
				6858000,
			) // 7.5 inches

			widthPt := NewEMU(
				slideWidthEMU,
			).ToPoints()
			heightPt := NewEMU(
				slideHeightEMU,
			).ToPoints()

			if !floatEquals(
				widthPt,
				720,
			) { // 10 * 72
				t.Errorf(
					"Slide width: EMU %d -> %v pt, want 720",
					slideWidthEMU,
					widthPt,
				)
			}
			if !floatEquals(
				heightPt,
				540,
			) { // 7.5 * 72
				t.Errorf(
					"Slide height: EMU %d -> %v pt, want 540",
					slideHeightEMU,
					heightPt,
				)
			}
		},
	)

	t.Run(
		"Word page margins",
		func(t *testing.T) {
			// Common Word margins (1" = 914400 EMU)
			marginEMU := int64(914400)
			marginPt := NewEMU(
				marginEMU,
			).ToPoints()

			if !floatEquals(marginPt, 72) {
				t.Errorf(
					"1 inch margin: EMU %d -> %v pt, want 72",
					marginEMU,
					marginPt,
				)
			}

			// Half-inch margin
			halfInchEMU := int64(457200)
			halfInchPt := NewEMU(
				halfInchEMU,
			).ToPoints()

			if !floatEquals(halfInchPt, 36) {
				t.Errorf(
					"0.5 inch margin: EMU %d -> %v pt, want 36",
					halfInchEMU,
					halfInchPt,
				)
			}
		},
	)

	t.Run(
		"Excel column width",
		func(t *testing.T) {
			// Excel default column width is about 64 pixels at 96 DPI
			// 64 pixels * (914400 / 96) EMU per pixel = 609600 EMU
			colWidthEMU := int64(609600)
			colWidthPt := NewEMU(
				colWidthEMU,
			).ToPoints()
			expectedPt := 64.0 * (72.0 / 96.0) // 48 points

			if !floatEquals(
				colWidthPt,
				expectedPt,
			) {
				t.Errorf(
					"Column width: EMU %d -> %v pt, want %v",
					colWidthEMU,
					colWidthPt,
					expectedPt,
				)
			}
		},
	)

	t.Run(
		"DrawingML rotation values",
		func(t *testing.T) {
			// Common DrawingML rotation values
			testCases := []struct {
				ooxmlAngle int32
				degrees    float64
				desc       string
			}{
				{0, 0, "no rotation"},
				{5400000, 90, "90 degrees"},
				{10800000, 180, "180 degrees"},
				{16200000, 270, "270 degrees"},
				{
					21600000,
					360,
					"360 degrees (full rotation)",
				},
				{2700000, 45, "45 degrees"},
				{
					-5400000,
					-90,
					"negative 90 degrees",
				},
			}

			for _, tc := range testCases {
				angle := NewOOXMLAngle(
					tc.ooxmlAngle,
				)
				degrees := angle.ToDegrees()
				if !floatEquals(
					degrees,
					tc.degrees,
				) {
					t.Errorf(
						"%s: OOXML %d -> %v degrees, want %v",
						tc.desc,
						tc.ooxmlAngle,
						degrees,
						tc.degrees,
					)
				}
			}
		},
	)

	t.Run(
		"DrawingML transparency values",
		func(t *testing.T) {
			// Common transparency values
			testCases := []struct {
				ooxmlPct int32
				float    float64
				desc     string
			}{
				{0, 0, "fully transparent"},
				{50000, 0.5, "50% opacity"},
				{100000, 1.0, "fully opaque"},
				{25000, 0.25, "25% opacity"},
				{75000, 0.75, "75% opacity"},
			}

			for _, tc := range testCases {
				pct := NewOOXMLPercentage(
					tc.ooxmlPct,
				)
				f := pct.ToFloat()
				if !floatEquals(f, tc.float) {
					t.Errorf(
						"%s: OOXML %d -> %v, want %v",
						tc.desc,
						tc.ooxmlPct,
						f,
						tc.float,
					)
				}
			}
		},
	)
}
