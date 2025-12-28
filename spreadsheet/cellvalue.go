// Package spreadsheet provides SpreadsheetML support for Excel documents.
//
//nolint:revive // file-length-limit: cell value logic is cohesive
package spreadsheet

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// CellValueType represents the logical type of a cell's value.
type CellValueType int

const (
	// CellValueEmpty represents an empty cell.
	CellValueEmpty CellValueType = iota

	// CellValueNumber represents a numeric value.
	CellValueNumber

	// CellValueString represents a text/string value.
	CellValueString

	// CellValueBoolean represents a boolean value (TRUE/FALSE).
	CellValueBoolean

	// CellValueError represents an error value (#DIV/0!, #N/A, etc.).
	CellValueError

	// CellValueDate represents a date/time value (stored as serial number).
	CellValueDate

	// CellValueFormula represents a cell with a formula.
	CellValueFormula
)

// String returns a string representation of the CellValueType.
func (t CellValueType) String() string {
	switch t {
	case CellValueEmpty:
		return "Empty"
	case CellValueNumber:
		return "Number"
	case CellValueString:
		return "String"
	case CellValueBoolean:
		return "Boolean"
	case CellValueError:
		return "Error"
	case CellValueDate:
		return "Date"
	case CellValueFormula:
		return "Formula"
	default:
		return "Unknown"
	}
}

// Excel error values.
const (
	ErrorNull    = "#NULL!"
	ErrorDiv0    = "#DIV/0!"
	ErrorValue   = "#VALUE!"
	ErrorRef     = "#REF!"
	ErrorName    = "#NAME?"
	ErrorNum     = "#NUM!"
	ErrorNA      = "#N/A"
	ErrorGetting = "#GETTING_DATA"
	ErrorSpill   = "#SPILL!"
	ErrorCalc    = "#CALC!"
)

// Bit size constants for parsing.
const (
	// bitSize64 is the bit size for 64-bit floating point parsing.
	bitSize64 = 64
)

const (
	booleanTrueString  = "TRUE"
	booleanFalseString = "FALSE"
)

// IsErrorValue returns true if the string is a valid Excel error value.
func IsErrorValue(s string) bool {
	switch s {
	case ErrorNull,
		ErrorDiv0,
		ErrorValue,
		ErrorRef,
		ErrorName,
		ErrorNum,
		ErrorNA,
		ErrorGetting,
		ErrorSpill,
		ErrorCalc:
		return true
	default:
		return false
	}
}

// GetCellValueType determines the logical type of a cell's value.
// This looks at the cell's data type attribute and formula presence.
func GetCellValueType(
	cell *elements.Cell,
) CellValueType {
	if cell == nil {
		return CellValueEmpty
	}

	// Check for formula first
	if cell.HasFormula() {
		return CellValueFormula
	}

	// Check the data type attribute
	switch cell.DataType() {
	case elements.CellTypeBoolean:
		return CellValueBoolean
	case elements.CellTypeError:
		return CellValueError
	case elements.CellTypeSharedString,
		elements.CellTypeInlineString,
		elements.CellTypeFormulaString:
		return CellValueString
	case elements.CellTypeDate:
		return CellValueDate
	case elements.CellTypeNumber:
		// Check if cell has any value
		if cell.CellValue() == nil ||
			cell.Value() == "" {
			return CellValueEmpty
		}

		return CellValueNumber
	default:
		// Default (no type attribute) is number
		if cell.CellValue() == nil ||
			cell.Value() == "" {
			return CellValueEmpty
		}

		return CellValueNumber
	}
}

// GetCellNumber returns the numeric value of a cell.
// Returns 0 and false if the cell doesn't contain a number.
func GetCellNumber(
	cell *elements.Cell,
) (float64, bool) {
	if cell == nil {
		return 0, false
	}

	// Must be a number type (or default)
	dt := cell.DataType()
	if dt != elements.CellTypeNumber && dt != "" {
		return 0, false
	}

	val := cell.Value()
	if val == "" {
		return 0, false
	}

	num, err := strconv.ParseFloat(val, bitSize64)
	if err != nil {
		return 0, false
	}

	return num, true
}

// GetCellString returns the string value of a cell.
// For shared string cells, it looks up the value in the shared string table.
// For inline strings, it extracts the text content.
// For other types, it returns the raw value.
func GetCellString(
	cell *elements.Cell,
	sst *parts.SharedStringTablePart,
) string {
	if cell == nil {
		return ""
	}

	//nolint:exhaustive // other cases handled by default
	switch cell.DataType() {
	case elements.CellTypeSharedString:
		// Look up in shared string table
		if sst == nil {
			return ""
		}
		indexStr := cell.Value()
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return ""
		}

		return sst.GetString(index)

	case elements.CellTypeInlineString:
		// Get from inline string element
		is := cell.InlineString()
		if is == nil {
			return ""
		}

		return is.PlainText()

	default:
		// Return raw value for other types
		return cell.Value()
	}
}

// GetCellBoolean returns the boolean value of a cell.
// Returns false and false if the cell doesn't contain a boolean.
func GetCellBoolean(
	cell *elements.Cell,
) (value, ok bool) {
	if cell == nil {
		return false, false
	}

	if cell.DataType() != elements.CellTypeBoolean {
		return false, false
	}

	val := cell.Value()
	switch val {
	case "1", "true", "TRUE":
		return true, true
	case "0", "false", "FALSE":
		return false, true
	default:
		return false, false
	}
}

// GetCellDate returns the date/time value of a cell.
// The date1904 parameter specifies whether to use the 1904 date system.
// Returns the zero time and false if the cell doesn't contain a valid date.
//
//nolint:revive // flag-parameter: date1904 is a standard Excel parameter
func GetCellDate(
	cell *elements.Cell,
	date1904 bool,
) (time.Time, bool) {
	if cell == nil {
		return time.Time{}, false
	}

	// Get the numeric value (date is stored as serial number)
	val := cell.Value()
	if val == "" {
		return time.Time{}, false
	}

	serial, err := strconv.ParseFloat(
		val,
		bitSize64,
	)
	if err != nil {
		return time.Time{}, false
	}

	if !IsValidExcelSerial(serial, date1904) {
		return time.Time{}, false
	}

	return DateFromExcelSerial(
		serial,
		date1904,
	), true
}

// GetCellError returns the error value of a cell.
// Returns an empty string and false if the cell doesn't contain an error.
func GetCellError(
	cell *elements.Cell,
) (string, bool) {
	if cell == nil {
		return "", false
	}

	if cell.DataType() != elements.CellTypeError {
		return "", false
	}

	val := cell.Value()
	if !IsErrorValue(val) {
		return "", false
	}

	return val, true
}

// SetCellNumber sets a cell to contain a numeric value.
func SetCellNumber(
	cell *elements.Cell,
	value float64,
) {
	cell.SetDataType(elements.CellTypeNumber)
	cell.SetValue(
		strconv.FormatFloat(
			value,
			'f',
			-1,
			bitSize64,
		),
	)
}

// SetCellString sets a cell to contain a string value using the shared
// string table. Returns the index of the string in the shared string table.
func SetCellString(
	cell *elements.Cell,
	value string,
	sst *parts.SharedStringTablePart,
) int {
	if sst == nil {
		// Fall back to inline string
		SetCellInlineString(cell, value)

		return -1
	}

	index := sst.AddString(value)
	cell.SetDataType(
		elements.CellTypeSharedString,
	)
	cell.SetValue(strconv.Itoa(index))

	return index
}

// SetCellInlineString sets a cell to contain an inline string value.
// This stores the string directly in the cell, not in the shared string table.
func SetCellInlineString(
	cell *elements.Cell,
	value string,
) {
	cell.SetDataType(
		elements.CellTypeInlineString,
	)
	// Remove any existing value element
	cell.RemoveValue()
	// Set inline string
	is := cell.GetOrCreateInlineString()
	is.SetPlainText(value)
}

// SetCellBoolean sets a cell to contain a boolean value.
//
//nolint:revive // flag-parameter: bool value is the actual data being set
func SetCellBoolean(
	cell *elements.Cell,
	value bool,
) {
	cell.SetDataType(elements.CellTypeBoolean)
	if value {
		cell.SetValue("1")
	} else {
		cell.SetValue("0")
	}
}

// SetCellError sets a cell to contain an error value.
func SetCellError(
	cell *elements.Cell,
	errorValue string,
) {
	cell.SetDataType(elements.CellTypeError)
	cell.SetValue(errorValue)
}

// SetCellDate sets a cell to contain a date value.
// The date is stored as an Excel serial number.
//
//nolint:revive // flag-parameter: date1904 is a standard Excel parameter
func SetCellDate(
	cell *elements.Cell,
	t time.Time,
	date1904 bool,
) {
	serial := ExcelDateSerial(t, date1904)
	// Dates are stored as numbers in Excel, but we can set the date type
	// attribute for applications that support it
	cell.SetDataType(elements.CellTypeNumber)
	cell.SetValue(
		strconv.FormatFloat(
			serial,
			'f',
			-1,
			bitSize64,
		),
	)
}

// FormatCellValue returns a formatted string representation of a cell's value.
// This is a simplified formatter; it doesn't apply full number format
// processing. The numFmt parameter is the number format code
// (e.g., "General", "#,##0.00", "yyyy-mm-dd").
//
//nolint:revive // confusing-results: return value names are clear
func FormatCellValue(
	cell *elements.Cell,
	sst *parts.SharedStringTablePart,
	numFmt string,
) string {
	if cell == nil {
		return ""
	}

	valueType := GetCellValueType(cell)

	switch valueType {
	case CellValueEmpty:
		return ""

	case CellValueString:
		return GetCellString(cell, sst)

	case CellValueBoolean:
		if val, ok := GetCellBoolean(cell); ok {
			if val {
				return booleanTrueString
			}

			return booleanFalseString
		}

		return ""

	case CellValueError:
		if val, ok := GetCellError(cell); ok {
			return val
		}

		return ""

	case CellValueNumber,
		CellValueDate,
		CellValueFormula:
		// For formulas, use the cached value
		val := cell.Value()
		if val == "" {
			return ""
		}

		// Try to parse as number for formatting
		num, err := strconv.ParseFloat(
			val,
			bitSize64,
		)
		if err != nil {
			// Not a number, return as-is
			return val
		}

		// Apply basic number formatting based on format code
		return formatNumber(num, numFmt)

	default:
		return cell.Value()
	}
}

// formatNumber applies basic number formatting.
// This is a simplified implementation - full format code parsing would
// be more complex.
func formatNumber(
	num float64,
	numFmt string,
) string {
	if numFmt == "" || numFmt == "General" {
		// General format - use default float formatting
		return strconv.FormatFloat(
			num,
			'f',
			-1,
			bitSize64,
		)
	}

	// Check for common date formats
	if isDateFormat(numFmt) {
		// Assume 1900 date system for simplicity
		t := DateFromExcelSerial(num, false)

		return formatDate(t, numFmt)
	}

	// Check for percentage
	if strings.Contains(numFmt, "%") {
		return strconv.FormatFloat(
			num*100,
			'f',
			2,
			bitSize64,
		) + "%"
	}

	// Check for decimal places
	decimalPlaces := countDecimalPlaces(numFmt)
	if decimalPlaces >= 0 {
		return strconv.FormatFloat(
			num,
			'f',
			decimalPlaces,
			bitSize64,
		)
	}

	// Default formatting
	return strconv.FormatFloat(
		num,
		'f',
		-1,
		bitSize64,
	)
}

// isDateFormat returns true if the format code appears to be a date format.
func isDateFormat(numFmt string) bool {
	lower := strings.ToLower(numFmt)
	// Check for common date/time format indicators
	datePatterns := []string{
		"yyyy",
		"yy",
		"mmmm",
		"mmm",
		"mm",
		"dd",
		"hh",
		"ss",
		"am/pm",
	}
	for _, pattern := range datePatterns {
		if strings.Contains(lower, pattern) {
			return true
		}
	}

	return false
}

// formatDate formats a time value using a simplified date format.
func formatDate(
	t time.Time,
	numFmt string,
) string {
	// Map common Excel format codes to Go format
	format := numFmt

	// Replace Excel format codes with Go equivalents
	replacements := map[string]string{
		"yyyy": "2006",
		"yy":   "06",
		"mmmm": "January",
		"mmm":  "Jan",
		"mm":   "01",
		"dd":   "02",
		"d":    "2",
		"hh":   "15",
		"h":    "3",
		"ss":   "05",
		"s":    "5",
	}

	for excel, golang := range replacements {
		format = strings.ReplaceAll(
			format,
			excel,
			golang,
		)
		format = strings.ReplaceAll(
			format,
			strings.ToUpper(excel),
			golang,
		)
	}

	return t.Format(format)
}

// countDecimalPlaces counts the decimal places in a format code.
// Returns -1 if the format doesn't specify decimal places.
func countDecimalPlaces(numFmt string) int {
	// Look for pattern like "0.00" or "#.##"
	re := regexp.MustCompile(`\.[0#]+`)
	match := re.FindString(numFmt)
	if match == "" {
		return -1
	}

	return len(
		match,
	) - 1 // Subtract 1 for the decimal point
}

// CellAddress combines a column name and row number into a cell reference
// string. For example: CellAddress("A", 1) returns "A1".
func CellAddress(col string, row int) string {
	return col + strconv.Itoa(row)
}

// CellAddressFromIndices creates a cell reference from column and row indices.
// Both indices are 1-based.
// For example: CellAddressFromIndices(1, 1) returns "A1".
func CellAddressFromIndices(col, row int) string {
	return ColumnName(col) + strconv.Itoa(row)
}

// SplitCellAddress splits a cell reference into column name and row number.
// Returns an error if the reference is invalid.
func SplitCellAddress(
	ref string,
) (col string, row int, err error) {
	cellRef, err := ParseCellRef(ref)
	if err != nil {
		return "", 0, err
	}

	return ColumnName(
		cellRef.Col,
	), cellRef.Row, nil
}
