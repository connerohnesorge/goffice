package spreadsheet

import (
	"math"
	"testing"
	"time"
)

func TestExcelDateSerial1900(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected float64
	}{
		{
			"Jan 1, 1900",
			time.Date(
				1900,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			1,
		},
		{
			"Jan 2, 1900",
			time.Date(
				1900,
				1,
				2,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			2,
		},
		{
			"Feb 28, 1900",
			time.Date(
				1900,
				2,
				28,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			59,
		},
		{
			"Mar 1, 1900 (after leap year bug)",
			time.Date(
				1900,
				3,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			61,
		},
		{
			"Jan 1, 1970 (Unix epoch)",
			time.Date(
				1970,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			25569,
		},
		{
			"Jan 1, 2000",
			time.Date(
				2000,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			36526,
		},
		{
			"Dec 31, 2099",
			time.Date(
				2099,
				12,
				31,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			73050,
		},
		// With time
		{
			"Jan 1, 1900 noon",
			time.Date(
				1900,
				1,
				1,
				12,
				0,
				0,
				0,
				time.UTC,
			),
			1.5,
		},
		{
			"Jan 1, 1900 6pm",
			time.Date(
				1900,
				1,
				1,
				18,
				0,
				0,
				0,
				time.UTC,
			),
			1.75,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExcelDateSerial(
				tt.date,
				false,
			)
			if math.Abs(
				result-tt.expected,
			) > 0.00001 {
				t.Errorf(
					"ExcelDateSerial(%v, false) = %v, want %v",
					tt.date,
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestExcelDateSerial1904(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected float64
	}{
		{
			"Jan 1, 1904 (epoch)",
			time.Date(
				1904,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			0,
		},
		{
			"Jan 2, 1904",
			time.Date(
				1904,
				1,
				2,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			1,
		},
		{
			"Jan 1, 1970",
			time.Date(
				1970,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			24107,
		},
		{
			"Jan 1, 2000",
			time.Date(
				2000,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			35064,
		},
		// With time
		{
			"Jan 1, 1904 noon",
			time.Date(
				1904,
				1,
				1,
				12,
				0,
				0,
				0,
				time.UTC,
			),
			0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExcelDateSerial(
				tt.date,
				true,
			)
			if math.Abs(
				result-tt.expected,
			) > 0.00001 {
				t.Errorf(
					"ExcelDateSerial(%v, true) = %v, want %v",
					tt.date,
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestDateFromExcelSerial1900(t *testing.T) {
	tests := []struct {
		name     string
		serial   float64
		expected time.Time
	}{
		{
			"Serial 1 (Jan 1, 1900)",
			1,
			time.Date(
				1900,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
		{
			"Serial 2 (Jan 2, 1900)",
			2,
			time.Date(
				1900,
				1,
				2,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
		{
			"Serial 59 (Feb 28, 1900)",
			59,
			time.Date(
				1900,
				2,
				28,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
		{
			"Serial 60 (Feb 29, 1900 - bug)",
			60,
			time.Date(
				1900,
				2,
				28,
				0,
				0,
				0,
				0,
				time.UTC,
			), // Maps to Feb 28
		},
		{
			"Serial 61 (Mar 1, 1900)",
			61,
			time.Date(
				1900,
				3,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
		{
			"Serial 25569 (Jan 1, 1970)",
			25569,
			time.Date(
				1970,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
		{
			"Serial 36526 (Jan 1, 2000)",
			36526,
			time.Date(
				2000,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
		// With time
		{
			"Serial 1.5 (Jan 1, 1900 noon)",
			1.5,
			time.Date(
				1900,
				1,
				1,
				12,
				0,
				0,
				0,
				time.UTC,
			),
		},
		{
			"Serial 1.75 (Jan 1, 1900 6pm)",
			1.75,
			time.Date(
				1900,
				1,
				1,
				18,
				0,
				0,
				0,
				time.UTC,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DateFromExcelSerial(
				tt.serial,
				false,
			)

			// Compare year, month, day, hour, minute
			if result.Year() != tt.expected.Year() ||
				result.Month() != tt.expected.Month() ||
				result.Day() != tt.expected.Day() ||
				result.Hour() != tt.expected.Hour() ||
				result.Minute() != tt.expected.Minute() {
				t.Errorf(
					"DateFromExcelSerial(%v, false) = %v, want %v",
					tt.serial,
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestDateFromExcelSerial1904(t *testing.T) {
	tests := []struct {
		name     string
		serial   float64
		expected time.Time
	}{
		{
			"Serial 0 (Jan 1, 1904)",
			0,
			time.Date(
				1904,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
		{
			"Serial 1 (Jan 2, 1904)",
			1,
			time.Date(
				1904,
				1,
				2,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
		{
			"Serial 24107 (Jan 1, 1970)",
			24107,
			time.Date(
				1970,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DateFromExcelSerial(
				tt.serial,
				true,
			)

			if result.Year() != tt.expected.Year() ||
				result.Month() != tt.expected.Month() ||
				result.Day() != tt.expected.Day() {
				t.Errorf(
					"DateFromExcelSerial(%v, true) = %v, want %v",
					tt.serial,
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestDateSerialRoundtrip1900(t *testing.T) {
	// Test dates after March 1, 1900 (after leap year bug)
	testDates := []time.Time{
		time.Date(
			1900,
			3,
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		),
		time.Date(
			1970,
			1,
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		),
		time.Date(
			2000,
			6,
			15,
			12,
			30,
			0,
			0,
			time.UTC,
		),
		time.Date(
			2024,
			12,
			25,
			18,
			45,
			0,
			0,
			time.UTC,
		),
	}

	for _, date := range testDates {
		serial := ExcelDateSerial(date, false)
		back := DateFromExcelSerial(serial, false)

		// Compare to second precision
		if date.Unix() != back.Unix() {
			t.Errorf(
				"Roundtrip failed for %v: serial=%v, back=%v",
				date,
				serial,
				back,
			)
		}
	}
}

func TestDateSerialRoundtrip1904(t *testing.T) {
	testDates := []time.Time{
		time.Date(
			1904,
			1,
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		),
		time.Date(
			1970,
			1,
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		),
		time.Date(
			2000,
			6,
			15,
			12,
			30,
			0,
			0,
			time.UTC,
		),
		time.Date(
			2024,
			12,
			25,
			18,
			45,
			0,
			0,
			time.UTC,
		),
	}

	for _, date := range testDates {
		serial := ExcelDateSerial(date, true)
		back := DateFromExcelSerial(serial, true)

		// Compare to second precision
		if date.Unix() != back.Unix() {
			t.Errorf(
				"Roundtrip failed for %v: serial=%v, back=%v",
				date,
				serial,
				back,
			)
		}
	}
}

func TestTimeToExcelTime(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected float64
	}{
		{
			"Midnight",
			time.Date(
				2000,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			0.0,
		},
		{
			"6 AM",
			time.Date(
				2000,
				1,
				1,
				6,
				0,
				0,
				0,
				time.UTC,
			),
			0.25,
		},
		{
			"Noon",
			time.Date(
				2000,
				1,
				1,
				12,
				0,
				0,
				0,
				time.UTC,
			),
			0.5,
		},
		{
			"6 PM",
			time.Date(
				2000,
				1,
				1,
				18,
				0,
				0,
				0,
				time.UTC,
			),
			0.75,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeToExcelTime(tt.time)
			if math.Abs(
				result-tt.expected,
			) > 0.00001 {
				t.Errorf(
					"TimeToExcelTime(%v) = %v, want %v",
					tt.time,
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestTimeFromExcelTime(t *testing.T) {
	tests := []struct {
		name    string
		frac    float64
		hours   int
		minutes int
		seconds int
	}{
		{"Midnight", 0.0, 0, 0, 0},
		{"6 AM", 0.25, 6, 0, 0},
		{"Noon", 0.5, 12, 0, 0},
		{"6 PM", 0.75, 18, 0, 0},
		{"3:30:15 PM", 0.6460069444, 15, 30, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, m, s := TimeFromExcelTime(tt.frac)
			if h != tt.hours || m != tt.minutes ||
				s != tt.seconds {
				t.Errorf(
					"TimeFromExcelTime(%v) = (%d, %d, %d), want (%d, %d, %d)",
					tt.frac,
					h,
					m,
					s,
					tt.hours,
					tt.minutes,
					tt.seconds,
				)
			}
		})
	}
}

func TestIsValidExcelSerial(t *testing.T) {
	tests := []struct {
		serial   float64
		date1904 bool
		valid    bool
	}{
		{1, false, true}, // 1900 system valid
		{
			0.5,
			false,
			false,
		}, // 1900 system, less than 1
		{
			0,
			false,
			false,
		}, // 1900 system, 0 is invalid
		{
			-1,
			false,
			false,
		}, // Negative is invalid

		{
			0,
			true,
			true,
		}, // 1904 system, 0 is valid
		{1, true, true},   // 1904 system valid
		{-1, true, false}, // Negative is invalid
	}

	for _, tt := range tests {
		result := IsValidExcelSerial(
			tt.serial,
			tt.date1904,
		)
		if result != tt.valid {
			t.Errorf(
				"IsValidExcelSerial(%v, %v) = %v, want %v",
				tt.serial,
				tt.date1904,
				result,
				tt.valid,
			)
		}
	}
}

func TestExcelSerialToUnix(t *testing.T) {
	// Jan 1, 1970 = Unix 0
	serial1970_1900 := ExcelDateSerial(
		time.Date(
			1970,
			1,
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		),
		false,
	)

	unix := ExcelSerialToUnix(
		serial1970_1900,
		false,
	)
	if unix != 0 {
		t.Errorf(
			"ExcelSerialToUnix(%v, false) = %v, want 0",
			serial1970_1900,
			unix,
		)
	}
}

func TestUnixToExcelSerial(t *testing.T) {
	// Unix 0 = Jan 1, 1970
	serial := UnixToExcelSerial(0, false)
	expected := float64(25569)

	if math.Abs(serial-expected) > 0.00001 {
		t.Errorf(
			"UnixToExcelSerial(0, false) = %v, want %v",
			serial,
			expected,
		)
	}
}

func TestDateOnly(t *testing.T) {
	tests := []struct {
		serial   float64
		expected int
	}{
		{36526.5, 36526},
		{36526.0, 36526},
		{36526.9999, 36526},
	}

	for _, tt := range tests {
		result := DateOnly(tt.serial)
		if result != tt.expected {
			t.Errorf(
				"DateOnly(%v) = %d, want %d",
				tt.serial,
				result,
				tt.expected,
			)
		}
	}
}

func TestTimeOnly(t *testing.T) {
	tests := []struct {
		serial   float64
		expected float64
	}{
		{36526.5, 0.5},
		{36526.0, 0.0},
		{36526.75, 0.75},
	}

	for _, tt := range tests {
		result := TimeOnly(tt.serial)
		if math.Abs(
			result-tt.expected,
		) > 0.00001 {
			t.Errorf(
				"TimeOnly(%v) = %v, want %v",
				tt.serial,
				result,
				tt.expected,
			)
		}
	}
}

func TestCombineDateAndTime(t *testing.T) {
	result := CombineDateAndTime(36526, 0.5)
	expected := 36526.5

	if math.Abs(result-expected) > 0.00001 {
		t.Errorf(
			"CombineDateAndTime(36526, 0.5) = %v, want %v",
			result,
			expected,
		)
	}
}

func TestConvertSerial1900To1904(t *testing.T) {
	// Same date in different systems should differ by ExcelSerialDiff1900To1904
	serial1900 := float64(
		36526,
	) // Jan 1, 2000 in 1900 system
	serial1904 := ConvertSerial1900To1904(
		serial1900,
	)
	expected := float64(
		35064,
	) // Jan 1, 2000 in 1904 system

	if math.Abs(serial1904-expected) > 0.00001 {
		t.Errorf(
			"ConvertSerial1900To1904(%v) = %v, want %v",
			serial1900,
			serial1904,
			expected,
		)
	}
}

func TestConvertSerial1904To1900(t *testing.T) {
	serial1904 := float64(
		35064,
	) // Jan 1, 2000 in 1904 system
	serial1900 := ConvertSerial1904To1900(
		serial1904,
	)
	expected := float64(
		36526,
	) // Jan 1, 2000 in 1900 system

	if math.Abs(serial1900-expected) > 0.00001 {
		t.Errorf(
			"ConvertSerial1904To1900(%v) = %v, want %v",
			serial1904,
			serial1900,
			expected,
		)
	}
}

func TestLeapYearBugHandling(t *testing.T) {
	// Test that we correctly handle the Excel 1900 leap year bug
	// Excel incorrectly treats 1900 as a leap year

	// Feb 28, 1900 should be serial 59
	feb28 := time.Date(
		1900,
		2,
		28,
		0,
		0,
		0,
		0,
		time.UTC,
	)
	serial := ExcelDateSerial(feb28, false)
	if serial != 59 {
		t.Errorf(
			"Feb 28, 1900 serial = %v, want 59",
			serial,
		)
	}

	// Mar 1, 1900 should be serial 61 (skipping 60 which is fake Feb 29)
	mar1 := time.Date(
		1900,
		3,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)
	serial = ExcelDateSerial(mar1, false)
	if serial != 61 {
		t.Errorf(
			"Mar 1, 1900 serial = %v, want 61",
			serial,
		)
	}

	// Serial 60 (Feb 29, 1900) should map to Feb 28, 1900
	back := DateFromExcelSerial(60, false)
	if back.Month() != 2 || back.Day() != 28 {
		t.Errorf(
			"Serial 60 should map to Feb 28, got %v",
			back,
		)
	}

	// Serial 61 should map to Mar 1, 1900
	back = DateFromExcelSerial(61, false)
	if back.Month() != 3 || back.Day() != 1 {
		t.Errorf(
			"Serial 61 should map to Mar 1, got %v",
			back,
		)
	}
}

