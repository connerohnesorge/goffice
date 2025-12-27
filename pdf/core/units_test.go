package core

import (
	"math"
	"testing"
)

func TestPointsPerInch(t *testing.T) {
	if PointsPerInch != 72.0 {
		t.Errorf(
			"PointsPerInch = %v, want 72.0",
			PointsPerInch,
		)
	}
}

func TestInchesToPoints(t *testing.T) {
	tests := []struct {
		inches float64
		want   float64
	}{
		{1.0, 72.0},
		{0.5, 36.0},
		{2.0, 144.0},
		{8.5, 612.0},  // Letter width
		{11.0, 792.0}, // Letter height
	}

	for _, tt := range tests {
		got := InchesToPoints(tt.inches)
		if got != tt.want {
			t.Errorf(
				"InchesToPoints(%v) = %v, want %v",
				tt.inches,
				got,
				tt.want,
			)
		}
	}
}

func TestPointsToInches(t *testing.T) {
	tests := []struct {
		points float64
		want   float64
	}{
		{72.0, 1.0},
		{36.0, 0.5},
		{144.0, 2.0},
		{612.0, 8.5},
		{792.0, 11.0},
	}

	for _, tt := range tests {
		got := PointsToInches(tt.points)
		if got != tt.want {
			t.Errorf(
				"PointsToInches(%v) = %v, want %v",
				tt.points,
				got,
				tt.want,
			)
		}
	}
}

func TestMMToPoints(t *testing.T) {
	// 25.4mm = 1 inch = 72 points
	mm := 25.4
	want := 72.0
	got := MMToPoints(mm)
	if math.Abs(got-want) > 0.001 {
		t.Errorf(
			"MMToPoints(%v) = %v, want %v",
			mm,
			got,
			want,
		)
	}
}

func TestPointsToMM(t *testing.T) {
	// 72 points = 1 inch = 25.4mm
	points := 72.0
	want := 25.4
	got := PointsToMM(points)
	if math.Abs(got-want) > 0.001 {
		t.Errorf(
			"PointsToMM(%v) = %v, want %v",
			points,
			got,
			want,
		)
	}
}

func TestCMToPoints(t *testing.T) {
	// 2.54cm = 1 inch = 72 points
	cm := 2.54
	want := 72.0
	got := CMToPoints(cm)
	if math.Abs(got-want) > 0.001 {
		t.Errorf(
			"CMToPoints(%v) = %v, want %v",
			cm,
			got,
			want,
		)
	}
}

func TestPointsToCM(t *testing.T) {
	// 72 points = 1 inch = 2.54cm
	points := 72.0
	want := 2.54
	got := PointsToCM(points)
	if math.Abs(got-want) > 0.001 {
		t.Errorf(
			"PointsToCM(%v) = %v, want %v",
			points,
			got,
			want,
		)
	}
}

func TestEMUToPoints(t *testing.T) {
	// 914400 EMU = 1 inch = 72 points
	emu := 914400.0
	want := 72.0
	got := EMUToPoints(emu)
	if math.Abs(got-want) > 0.001 {
		t.Errorf(
			"EMUToPoints(%v) = %v, want %v",
			emu,
			got,
			want,
		)
	}
}

func TestPointsToEMU(t *testing.T) {
	// 72 points = 1 inch = 914400 EMU
	points := 72.0
	want := 914400.0
	got := PointsToEMU(points)
	if math.Abs(got-want) > 0.001 {
		t.Errorf(
			"PointsToEMU(%v) = %v, want %v",
			points,
			got,
			want,
		)
	}
}

func TestTwipsToPoints(t *testing.T) {
	// 20 twips = 1 point
	// 1440 twips = 1 inch = 72 points
	tests := []struct {
		twips float64
		want  float64
	}{
		{20.0, 1.0},
		{1440.0, 72.0},
		{2880.0, 144.0},
	}

	for _, tt := range tests {
		got := TwipsToPoints(tt.twips)
		if math.Abs(got-tt.want) > 0.001 {
			t.Errorf(
				"TwipsToPoints(%v) = %v, want %v",
				tt.twips,
				got,
				tt.want,
			)
		}
	}
}

func TestPointsToTwips(t *testing.T) {
	tests := []struct {
		points float64
		want   float64
	}{
		{1.0, 20.0},
		{72.0, 1440.0},
		{144.0, 2880.0},
	}

	for _, tt := range tests {
		got := PointsToTwips(tt.points)
		if math.Abs(got-tt.want) > 0.001 {
			t.Errorf(
				"PointsToTwips(%v) = %v, want %v",
				tt.points,
				got,
				tt.want,
			)
		}
	}
}

func TestUnit_String(t *testing.T) {
	tests := []struct {
		unit Unit
		want string
	}{
		{UnitPoint, "pt"},
		{UnitInch, "in"},
		{UnitMM, "mm"},
		{UnitCM, "cm"},
		{UnitEMU, "emu"},
		{UnitTwip, "twip"},
		{
			Unit(99),
			"pt",
		}, // Unknown defaults to pt
	}

	for _, tt := range tests {
		if got := tt.unit.String(); got != tt.want {
			t.Errorf(
				"Unit(%d).String() = %q, want %q",
				tt.unit,
				got,
				tt.want,
			)
		}
	}
}

func TestUnit_ToPoints(t *testing.T) {
	tests := []struct {
		unit  Unit
		value float64
		want  float64
	}{
		{UnitPoint, 72.0, 72.0},
		{UnitInch, 1.0, 72.0},
		{UnitMM, 25.4, 72.0},
		{UnitCM, 2.54, 72.0},
		{UnitEMU, 914400.0, 72.0},
		{UnitTwip, 1440.0, 72.0},
	}

	for _, tt := range tests {
		got := tt.unit.ToPoints(tt.value)
		if math.Abs(got-tt.want) > 0.001 {
			t.Errorf(
				"%s.ToPoints(%v) = %v, want %v",
				tt.unit,
				tt.value,
				got,
				tt.want,
			)
		}
	}
}

func TestUnit_FromPoints(t *testing.T) {
	points := 72.0
	tests := []struct {
		unit Unit
		want float64
	}{
		{UnitPoint, 72.0},
		{UnitInch, 1.0},
		{UnitMM, 25.4},
		{UnitCM, 2.54},
		{UnitEMU, 914400.0},
		{UnitTwip, 1440.0},
	}

	for _, tt := range tests {
		got := tt.unit.FromPoints(points)
		if math.Abs(got-tt.want) > 0.001 {
			t.Errorf(
				"%s.FromPoints(%v) = %v, want %v",
				tt.unit,
				points,
				got,
				tt.want,
			)
		}
	}
}

func TestDimension_ToPoints(t *testing.T) {
	tests := []struct {
		dim  Dimension
		want float64
	}{
		{Points(72), 72.0},
		{Inches(1), 72.0},
		{MM(25.4), 72.0},
		{CM(2.54), 72.0},
		{EMU(914400), 72.0},
		{Twips(1440), 72.0},
	}

	for _, tt := range tests {
		got := tt.dim.ToPoints()
		if math.Abs(got-tt.want) > 0.001 {
			t.Errorf(
				"%v.ToPoints() = %v, want %v",
				tt.dim,
				got,
				tt.want,
			)
		}
	}
}

func TestDimension_Convert(t *testing.T) {
	// Convert 1 inch to various units
	oneInch := Inches(1)

	tests := []struct {
		toUnit Unit
		want   float64
	}{
		{UnitPoint, 72.0},
		{UnitInch, 1.0},
		{UnitMM, 25.4},
		{UnitCM, 2.54},
	}

	for _, tt := range tests {
		got := oneInch.Convert(tt.toUnit)
		if got.Unit != tt.toUnit {
			t.Errorf(
				"Convert(%s).Unit = %s, want %s",
				tt.toUnit,
				got.Unit,
				tt.toUnit,
			)
		}
		if math.Abs(got.Value-tt.want) > 0.001 {
			t.Errorf(
				"Convert(%s).Value = %v, want %v",
				tt.toUnit,
				got.Value,
				tt.want,
			)
		}
	}
}

func TestDimension_Add(t *testing.T) {
	a := Inches(1)  // 72 points
	b := Points(36) // 36 points
	result := a.Add(b)

	if result.Unit != UnitPoint {
		t.Errorf(
			"Add result unit = %s, want pt",
			result.Unit,
		)
	}
	if result.Value != 108 {
		t.Errorf(
			"Add result value = %v, want 108",
			result.Value,
		)
	}
}

func TestDimension_Sub(t *testing.T) {
	a := Inches(1)  // 72 points
	b := Points(36) // 36 points
	result := a.Sub(b)

	if result.Unit != UnitPoint {
		t.Errorf(
			"Sub result unit = %s, want pt",
			result.Unit,
		)
	}
	if result.Value != 36 {
		t.Errorf(
			"Sub result value = %v, want 36",
			result.Value,
		)
	}
}

func TestDimension_Mul(t *testing.T) {
	d := Inches(1)
	result := d.Mul(2)

	if result.Unit != UnitInch {
		t.Errorf(
			"Mul result unit = %s, want in",
			result.Unit,
		)
	}
	if result.Value != 2 {
		t.Errorf(
			"Mul result value = %v, want 2",
			result.Value,
		)
	}
}

func TestDimension_String(t *testing.T) {
	tests := []struct {
		dim  Dimension
		want string
	}{
		{Points(72), "72pt"},
		{Inches(1), "1in"},
		{Inches(1.5), "1.5in"},
	}

	for _, tt := range tests {
		if got := tt.dim.String(); got != tt.want {
			t.Errorf(
				"%v.String() = %q, want %q",
				tt.dim,
				got,
				tt.want,
			)
		}
	}
}

func TestRoundToDecimalPlaces(t *testing.T) {
	tests := []struct {
		value  float64
		places int
		want   float64
	}{
		{3.14159, 2, 3.14},
		{3.145, 2, 3.15},
		{3.144, 2, 3.14},
		{100.0, 2, 100.0},
		{0.001, 2, 0.0},
		{0.005, 2, 0.01},
	}

	for _, tt := range tests {
		got := RoundToDecimalPlaces(
			tt.value,
			tt.places,
		)
		if got != tt.want {
			t.Errorf(
				"RoundToDecimalPlaces(%v, %d) = %v, want %v",
				tt.value,
				tt.places,
				got,
				tt.want,
			)
		}
	}
}

func TestNewDimension(t *testing.T) {
	d := NewDimension(100, UnitMM)
	if d.Value != 100 {
		t.Errorf(
			"NewDimension Value = %v, want 100",
			d.Value,
		)
	}
	if d.Unit != UnitMM {
		t.Errorf(
			"NewDimension Unit = %s, want mm",
			d.Unit,
		)
	}
}

func TestConversionRoundTrip(t *testing.T) {
	// Test that conversions round-trip correctly
	original := 72.0 // 1 inch in points

	// Points -> Inches -> Points
	inches := PointsToInches(original)
	backToPoints := InchesToPoints(inches)
	if backToPoints != original {
		t.Errorf(
			"Inches round-trip: got %v, want %v",
			backToPoints,
			original,
		)
	}

	// Points -> MM -> Points
	mm := PointsToMM(original)
	backToPoints = MMToPoints(mm)
	if math.Abs(backToPoints-original) > 0.0001 {
		t.Errorf(
			"MM round-trip: got %v, want %v",
			backToPoints,
			original,
		)
	}

	// Points -> CM -> Points
	cm := PointsToCM(original)
	backToPoints = CMToPoints(cm)
	if math.Abs(backToPoints-original) > 0.0001 {
		t.Errorf(
			"CM round-trip: got %v, want %v",
			backToPoints,
			original,
		)
	}

	// Points -> EMU -> Points
	emu := PointsToEMU(original)
	backToPoints = EMUToPoints(emu)
	if math.Abs(backToPoints-original) > 0.0001 {
		t.Errorf(
			"EMU round-trip: got %v, want %v",
			backToPoints,
			original,
		)
	}

	// Points -> Twips -> Points
	twips := PointsToTwips(original)
	backToPoints = TwipsToPoints(twips)
	if math.Abs(backToPoints-original) > 0.0001 {
		t.Errorf(
			"Twips round-trip: got %v, want %v",
			backToPoints,
			original,
		)
	}
}

// TestEMUToPointsRoundTripAccuracy tests EMU to points conversion with various edge cases.
func TestEMUToPointsRoundTripAccuracy(
	t *testing.T,
) {
	testCases := []struct {
		name     string
		emuValue float64
		desc     string
	}{
		{"zero", 0, "zero value"},
		{"one_point", 12700, "1 point in EMU"},
		{"one_inch", 914400, "1 inch in EMU"},
		{"half_point", 6350, "0.5 points in EMU"},
		{
			"large_value",
			914400 * 100,
			"100 inches in EMU",
		},
		{
			"very_large",
			914400 * 1000,
			"1000 inches in EMU",
		},
		{
			"small_fraction",
			1270,
			"0.1 points in EMU",
		},
		{
			"tiny_fraction",
			127,
			"0.01 points in EMU",
		},
		{
			"max_int32_range",
			2147483647,
			"max int32 as EMU",
		},
		{
			"negative_value",
			-914400,
			"negative 1 inch",
		},
		{
			"negative_small",
			-12700,
			"negative 1 point",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// EMU -> Points -> EMU
			points := EMUToPoints(tc.emuValue)
			backToEMU := PointsToEMU(points)

			// Allow small tolerance for floating point
			tolerance := math.Abs(
				tc.emuValue,
			) * 1e-10
			if tolerance < 1e-6 {
				tolerance = 1e-6
			}

			if math.Abs(
				backToEMU-tc.emuValue,
			) > tolerance {
				t.Errorf(
					"%s: EMU round-trip failed: original=%v, after round-trip=%v, diff=%v",
					tc.desc,
					tc.emuValue,
					backToEMU,
					math.Abs(
						backToEMU-tc.emuValue,
					),
				)
			}
		})
	}
}

// TestPointsToEMURoundTripAccuracy tests points to EMU conversion with various edge cases.
func TestPointsToEMURoundTripAccuracy(
	t *testing.T,
) {
	testCases := []struct {
		name        string
		pointsValue float64
		desc        string
	}{
		{"zero", 0, "zero value"},
		{"one_point", 1.0, "1 point"},
		{"one_inch", 72.0, "1 inch (72 points)"},
		{"half_point", 0.5, "0.5 points"},
		{
			"letter_width",
			612.0,
			"letter page width",
		},
		{
			"letter_height",
			792.0,
			"letter page height",
		},
		{"a4_width", 595.276, "A4 page width"},
		{"a4_height", 841.89, "A4 page height"},
		{"large_value", 7200.0, "100 inches"},
		{"very_large", 72000.0, "1000 inches"},
		{"small_fraction", 0.1, "0.1 points"},
		{"tiny_fraction", 0.01, "0.01 points"},
		{
			"negative_value",
			-72.0,
			"negative 1 inch",
		},
		{
			"negative_small",
			-1.0,
			"negative 1 point",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Points -> EMU -> Points
			emu := PointsToEMU(tc.pointsValue)
			backToPoints := EMUToPoints(emu)

			// Allow small tolerance for floating point
			tolerance := math.Abs(
				tc.pointsValue,
			) * 1e-10
			if tolerance < 1e-10 {
				tolerance = 1e-10
			}

			if math.Abs(
				backToPoints-tc.pointsValue,
			) > tolerance {
				t.Errorf(
					"%s: Points round-trip failed: original=%v, after round-trip=%v, diff=%v",
					tc.desc,
					tc.pointsValue,
					backToPoints,
					math.Abs(
						backToPoints-tc.pointsValue,
					),
				)
			}
		})
	}
}

// TestUnitConversionEdgeCases tests edge cases in unit conversions.
func TestUnitConversionEdgeCases(t *testing.T) {
	t.Run("zero values", func(t *testing.T) {
		if InchesToPoints(0) != 0 {
			t.Error(
				"InchesToPoints(0) should be 0",
			)
		}
		if PointsToInches(0) != 0 {
			t.Error(
				"PointsToInches(0) should be 0",
			)
		}
		if MMToPoints(0) != 0 {
			t.Error("MMToPoints(0) should be 0")
		}
		if PointsToMM(0) != 0 {
			t.Error("PointsToMM(0) should be 0")
		}
		if EMUToPoints(0) != 0 {
			t.Error("EMUToPoints(0) should be 0")
		}
		if PointsToEMU(0) != 0 {
			t.Error("PointsToEMU(0) should be 0")
		}
		if TwipsToPoints(0) != 0 {
			t.Error(
				"TwipsToPoints(0) should be 0",
			)
		}
		if PointsToTwips(0) != 0 {
			t.Error(
				"PointsToTwips(0) should be 0",
			)
		}
	})

	t.Run("negative values", func(t *testing.T) {
		// Negative values should convert correctly (preserving sign)
		if InchesToPoints(-1) != -72 {
			t.Errorf(
				"InchesToPoints(-1) = %v, want -72",
				InchesToPoints(-1),
			)
		}
		if PointsToInches(-72) != -1 {
			t.Errorf(
				"PointsToInches(-72) = %v, want -1",
				PointsToInches(-72),
			)
		}
		negEMU := EMUToPoints(-914400)
		if math.Abs(negEMU-(-72)) > 0.001 {
			t.Errorf(
				"EMUToPoints(-914400) = %v, want -72",
				negEMU,
			)
		}
	})

	t.Run(
		"very large values",
		func(t *testing.T) {
			// Test with very large values (within float64 precision)
			largeInches := 1e6 // 1 million inches
			largePoints := InchesToPoints(
				largeInches,
			)
			if math.Abs(
				largePoints-72e6,
			) > 0.001 {
				t.Errorf(
					"InchesToPoints(1e6) = %v, want 72e6",
					largePoints,
				)
			}
			backToInches := PointsToInches(
				largePoints,
			)
			if math.Abs(
				backToInches-largeInches,
			) > 1e-6 {
				t.Errorf(
					"Round-trip large inches failed: got %v, want %v",
					backToInches,
					largeInches,
				)
			}
		},
	)

	t.Run(
		"very small values",
		func(t *testing.T) {
			// Test with very small values
			tinyInches := 1e-6 // 1 millionth of an inch
			tinyPoints := InchesToPoints(
				tinyInches,
			)
			if math.Abs(
				tinyPoints-72e-6,
			) > 1e-12 {
				t.Errorf(
					"InchesToPoints(1e-6) = %v, want 72e-6",
					tinyPoints,
				)
			}
		},
	)
}

// TestDimensionOperationsEdgeCases tests edge cases in Dimension operations.
func TestDimensionOperationsEdgeCases(
	t *testing.T,
) {
	t.Run(
		"mixed unit operations",
		func(t *testing.T) {
			// Adding dimensions in different units
			d1 := Inches(1) // 72 points
			d2 := MM(
				25.4,
			) // 72 points (1 inch)
			d3 := EMU(
				914400,
			) // 72 points (1 inch)
			d4 := Twips(
				1440,
			) // 72 points (1 inch)

			sum := d1.Add(d2).Add(d3).Add(d4)
			expected := 72.0 * 4

			if math.Abs(
				sum.ToPoints()-expected,
			) > 0.01 {
				t.Errorf(
					"Sum of 4 inches in different units = %v points, want %v",
					sum.ToPoints(),
					expected,
				)
			}
		},
	)

	t.Run(
		"subtraction resulting in negative",
		func(t *testing.T) {
			d1 := Points(50)
			d2 := Points(100)
			result := d1.Sub(d2)

			if result.Value != -50 {
				t.Errorf(
					"50pt - 100pt = %v, want -50",
					result.Value,
				)
			}
		},
	)

	t.Run(
		"multiplication by zero",
		func(t *testing.T) {
			d := Inches(1)
			result := d.Mul(0)

			if result.Value != 0 {
				t.Errorf(
					"1in * 0 = %v, want 0",
					result.Value,
				)
			}
		},
	)

	t.Run(
		"multiplication by negative",
		func(t *testing.T) {
			d := Inches(1)
			result := d.Mul(-2)

			if result.Value != -2 {
				t.Errorf(
					"1in * -2 = %v, want -2",
					result.Value,
				)
			}
		},
	)
}

// TestUnitEnumEdgeCases tests edge cases in Unit enum.
func TestUnitEnumEdgeCases(t *testing.T) {
	t.Run(
		"unknown unit defaults",
		func(t *testing.T) {
			unknownUnit := Unit(999)

			// String should default to "pt"
			if unknownUnit.String() != "pt" {
				t.Errorf(
					"Unknown unit String() = %q, want %q",
					unknownUnit.String(),
					"pt",
				)
			}

			// ToPoints should return value unchanged
			if unknownUnit.ToPoints(100) != 100 {
				t.Errorf(
					"Unknown unit ToPoints(100) = %v, want 100",
					unknownUnit.ToPoints(100),
				)
			}

			// FromPoints should return value unchanged
			if unknownUnit.FromPoints(
				100,
			) != 100 {
				t.Errorf(
					"Unknown unit FromPoints(100) = %v, want 100",
					unknownUnit.FromPoints(100),
				)
			}
		},
	)
}

// TestKnownConversions tests conversions with known real-world values.
func TestKnownConversions(t *testing.T) {
	// Test with standard page sizes and their EMU equivalents
	t.Run("US Letter page", func(t *testing.T) {
		// US Letter: 8.5 x 11 inches
		letterWidthPt := 612.0  // 8.5 * 72
		letterHeightPt := 792.0 // 11 * 72

		letterWidthEMU := 7772400.0   // 8.5 * 914400
		letterHeightEMU := 10058400.0 // 11 * 914400

		// Convert EMU to points and verify
		widthPt := EMUToPoints(letterWidthEMU)
		heightPt := EMUToPoints(letterHeightEMU)

		if math.Abs(
			widthPt-letterWidthPt,
		) > 0.001 {
			t.Errorf(
				"Letter width: EMU %v -> %v pt, want %v pt",
				letterWidthEMU,
				widthPt,
				letterWidthPt,
			)
		}
		if math.Abs(
			heightPt-letterHeightPt,
		) > 0.001 {
			t.Errorf(
				"Letter height: EMU %v -> %v pt, want %v pt",
				letterHeightEMU,
				heightPt,
				letterHeightPt,
			)
		}
	})

	t.Run("A4 page", func(t *testing.T) {
		// A4: 210 x 297 mm
		a4WidthMM := 210.0
		a4HeightMM := 297.0

		// Convert to points
		a4WidthPt := MMToPoints(a4WidthMM)
		a4HeightPt := MMToPoints(a4HeightMM)

		// Expected values (approximately)
		expectedWidthPt := 595.28  // 210mm in points
		expectedHeightPt := 841.89 // 297mm in points

		if math.Abs(
			a4WidthPt-expectedWidthPt,
		) > 0.1 {
			t.Errorf(
				"A4 width: %v mm -> %v pt, want ~%v pt",
				a4WidthMM,
				a4WidthPt,
				expectedWidthPt,
			)
		}
		if math.Abs(
			a4HeightPt-expectedHeightPt,
		) > 0.1 {
			t.Errorf(
				"A4 height: %v mm -> %v pt, want ~%v pt",
				a4HeightMM,
				a4HeightPt,
				expectedHeightPt,
			)
		}
	})

	t.Run(
		"standard line height in twips",
		func(t *testing.T) {
			// Word uses 240 twips for single line spacing (12pt)
			lineHeightTwips := 240.0
			lineHeightPt := TwipsToPoints(
				lineHeightTwips,
			)

			if math.Abs(
				lineHeightPt-12.0,
			) > 0.001 {
				t.Errorf(
					"240 twips -> %v pt, want 12 pt",
					lineHeightPt,
				)
			}
		},
	)
}
