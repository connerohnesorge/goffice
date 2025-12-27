// Package spreadsheet provides SpreadsheetML support for Excel documents.
package spreadsheet

import (
	"math"
	"time"
)

// Excel date system epochs.
var (
	// Excel1900Epoch is the epoch for the 1900 date system.
	// Excel serial number 1 represents January 1, 1900.
	// We use Dec 31, 1899 as the epoch so that serial 1 = Jan 1, 1900.
	// Note: Excel incorrectly considers 1900 a leap year, so serial 60
	// is Feb 29, 1900 (which doesn't actually exist). We handle this
	// bug for compatibility.
	Excel1900Epoch = time.Date(
		1899,
		12,
		31,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	// Excel1904Epoch is the epoch for the 1904 date system (Jan 1, 1904).
	// This system was used by Excel for Mac in early versions.
	Excel1904Epoch = time.Date(
		1904,
		1,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)
)

// Excel date serial number constants.
const (
	// excel1900LeapYearBugSerial is the serial number for the
	// non-existent Feb 29, 1900. Excel incorrectly treats 1900
	// as a leap year.
	excel1900LeapYearBugSerial = 60

	// secondsPerDay is the number of seconds in a day.
	secondsPerDay = 24 * 60 * 60

	// nanosecondsPerDay is the number of nanoseconds in a day.
	nanosecondsPerDay = secondsPerDay * 1e9

	// secondsPerHour is the number of seconds in an hour.
	secondsPerHour = 3600

	// secondsPerMinute is the number of seconds in a minute.
	secondsPerMinute = 60

	// march1900Year is the year 1900 for Excel's leap year bug handling.
	march1900Year = 1900

	// march1900Month is the month of March (3) for Excel's leap year
	// bug handling.
	march1900Month = 3
)

// ExcelDateSerial converts a Go time.Time to an Excel date serial number.
// The date1904 parameter specifies whether to use the 1904 date system.
//
// In the 1900 date system:
//   - Serial 1 = January 1, 1900
//   - Serial 60 = February 29, 1900 (Excel's leap year bug - doesn't exist)
//   - Serial 61 = March 1, 1900
//
// In the 1904 date system:
//   - Serial 0 = January 1, 1904
//   - No leap year bug exists
//
// The fractional part represents the time of day:
//   - 0.0 = midnight (00:00:00)
//   - 0.5 = noon (12:00:00)
//   - 0.75 = 6:00 PM (18:00:00)
//
//nolint:revive // flag-parameter: bool param is standard for date system selection
func ExcelDateSerial(
	t time.Time,
	date1904 bool,
) float64 {
	// Use UTC for calculations
	utcTime := t.UTC()

	if date1904 {
		// 1904 date system - straightforward calculation
		duration := utcTime.Sub(Excel1904Epoch)

		return float64(
			duration,
		) / float64(
			nanosecondsPerDay,
		)
	}

	// 1900 date system
	// Calculate days since epoch (Dec 31, 1899)
	// This gives us serial 1 = Jan 1, 1900
	duration := utcTime.Sub(Excel1900Epoch)
	serial := float64(
		duration,
	) / float64(
		nanosecondsPerDay,
	)

	// Handle Excel's 1900 leap year bug:
	// Excel thinks Feb 29, 1900 exists (serial 60), but it doesn't.
	// For dates on or after Mar 1, 1900 (serial 60 without the bug),
	// we need to add 1 to account for the phantom Feb 29, 1900.
	// Feb 28, 1900 = serial 59
	// Mar 1, 1900 = serial 61 (not 60, because 60 is the fake Feb 29)
	mar1_1900 := time.Date(
		march1900Year,
		march1900Month,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)
	if !utcTime.Before(mar1_1900) {
		serial++
	}

	return serial
}

// DateFromExcelSerial converts an Excel date serial number to a Go time.Time.
// The date1904 parameter specifies whether to use the 1904 date system.
//
// This function handles Excel's leap year bug in the 1900 date system.
// Serial number 60 (Feb 29, 1900) is treated as Feb 28, 1900 since
// Feb 29, 1900 doesn't actually exist (1900 was not a leap year).
//
//nolint:revive // flag-parameter: bool param is standard for date system selection
func DateFromExcelSerial(
	serial float64,
	date1904 bool,
) time.Time {
	if date1904 {
		// 1904 date system - straightforward calculation
		duration := time.Duration(
			serial * float64(nanosecondsPerDay),
		)

		return Excel1904Epoch.Add(duration)
	}

	// 1900 date system with leap year bug handling
	adjustedSerial := serial

	// Handle the leap year bug:
	// Serial 60 is the phantom Feb 29, 1900 - map it to Feb 28
	// Serial >= 61 needs to be adjusted down by 1 to get the correct date
	if serial >= excel1900LeapYearBugSerial+1 {
		// After the phantom Feb 29, subtract 1
		adjustedSerial = serial - 1
	} else if serial >= excel1900LeapYearBugSerial {
		// Serial 60 (Feb 29, 1900) - map to Feb 28 (serial 59)
		adjustedSerial = excel1900LeapYearBugSerial - 1
	}

	duration := time.Duration(
		adjustedSerial * float64(
			nanosecondsPerDay,
		),
	)

	return Excel1900Epoch.Add(duration)
}

// TimeToExcelTime converts just the time portion of a time.Time to an Excel
// time value. The returned value is between 0.0 (midnight) and 1.0
// (just before midnight).
//
// Excel represents times as fractions of a day:
//   - 0.0 = 00:00:00
//   - 0.25 = 06:00:00
//   - 0.5 = 12:00:00
//   - 0.75 = 18:00:00
func TimeToExcelTime(t time.Time) float64 {
	utcTime := t.UTC()
	// Get the time of day as duration since midnight
	midnight := time.Date(
		utcTime.Year(),
		utcTime.Month(),
		utcTime.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)
	duration := utcTime.Sub(midnight)

	return float64(
		duration,
	) / float64(
		nanosecondsPerDay,
	)
}

// TimeFromExcelTime converts an Excel time fraction to hours, minutes, seconds.
// The fraction should be between 0.0 and 1.0.
func TimeFromExcelTime(
	fraction float64,
) (hours, minutes, seconds int) {
	// Ensure fraction is in valid range
	normalizedFraction := math.Mod(fraction, 1.0)
	if normalizedFraction < 0 {
		normalizedFraction += 1.0
	}

	totalSeconds := int(
		math.Round(
			normalizedFraction * secondsPerDay,
		),
	)
	hours = totalSeconds / secondsPerHour
	minutes = (totalSeconds % secondsPerHour) / secondsPerMinute
	seconds = totalSeconds % secondsPerMinute

	return hours, minutes, seconds
}

// IsValidExcelSerial returns true if the serial number represents a valid date.
// For the 1900 system, valid serials are >= 1 (Jan 1, 1900).
// For the 1904 system, valid serials are >= 0 (Jan 1, 1904).
// Negative serials (before the epoch) are not valid.
//
//nolint:revive // flag-parameter: bool param is standard for date system selection
func IsValidExcelSerial(
	serial float64,
	date1904 bool,
) bool {
	if date1904 {
		return serial >= 0
	}

	return serial >= 1
}

// ExcelSerialToUnix converts an Excel date serial to a Unix timestamp.
// This is useful for interoperability with other systems.
func ExcelSerialToUnix(
	serial float64,
	date1904 bool,
) int64 {
	t := DateFromExcelSerial(serial, date1904)

	return t.Unix()
}

// UnixToExcelSerial converts a Unix timestamp to an Excel date serial.
func UnixToExcelSerial(
	unix int64,
	date1904 bool,
) float64 {
	t := time.Unix(unix, 0).UTC()

	return ExcelDateSerial(t, date1904)
}

// DateOnly returns the integer part of an Excel serial (the date without time).
func DateOnly(serial float64) int {
	return int(math.Floor(serial))
}

// TimeOnly returns the fractional part of an Excel serial (the time
// without date).
func TimeOnly(serial float64) float64 {
	return serial - math.Floor(serial)
}

// CombineDateAndTime combines a date serial (integer) with a time fraction.
func CombineDateAndTime(
	dateSerial int,
	timeFraction float64,
) float64 {
	return float64(dateSerial) + timeFraction
}

// Excel serial number for specific dates (1900 system).
const (
	// ExcelSerial1900Jan1 is the serial for Jan 1, 1900.
	ExcelSerial1900Jan1 = 1

	// ExcelSerial1900Mar1 is the serial for Mar 1, 1900 (after the
	// leap year bug).
	ExcelSerial1900Mar1 = 61

	// ExcelSerial1904Jan1 is the serial for Jan 1, 1904 (in 1904 system).
	ExcelSerial1904Jan1 = 0

	// ExcelSerialDiff1900To1904 is the difference between date systems.
	// Adding this to a 1904 serial gives the equivalent 1900 serial.
	ExcelSerialDiff1900To1904 = 1462
)

// ConvertSerial1900To1904 converts a 1900 system serial to 1904 system.
func ConvertSerial1900To1904(
	serial float64,
) float64 {
	return serial - ExcelSerialDiff1900To1904
}

// ConvertSerial1904To1900 converts a 1904 system serial to 1900 system.
func ConvertSerial1904To1900(
	serial float64,
) float64 {
	return serial + ExcelSerialDiff1900To1904
}
