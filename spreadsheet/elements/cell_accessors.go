package elements

import (
	"strconv"
)

// SetNumberValue sets the cell to a numeric value.
func (c *Cell) SetNumberValue(value float64) {
	c.SetDataType(CellTypeNumber)
	cv := c.GetOrCreateCellValue()
	// 64-bit float size
	cv.SetValue(
		strconv.FormatFloat(
			value,
			'f',
			-1,
			64, //nolint:mnd,revive // add-constant: 64-bit float size
		),
	)
}

// SetBoolValue sets the cell to a boolean value.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Cell) SetBoolValue(value bool) {
	c.SetDataType(CellTypeBoolean)
	cv := c.GetOrCreateCellValue()
	if value {
		cv.SetValue("1")
	} else {
		cv.SetValue("0")
	}
}

// SetSharedStringIndex sets the cell to reference a shared string.
func (c *Cell) SetSharedStringIndex(index int) {
	c.SetDataType(CellTypeSharedString)
	cv := c.GetOrCreateCellValue()
	cv.SetValue(strconv.Itoa(index))
}

// SetNumber is an alias for SetNumberValue - sets the cell to a numeric value.
func (c *Cell) SetNumber(value float64) {
	c.SetNumberValue(value)
}

// SetBoolean is an alias for SetBoolValue - sets the cell to a boolean value.
func (c *Cell) SetBoolean(value bool) {
	c.SetBoolValue(value)
}

// SetInlineString sets the cell to contain an inline string.
// This stores the string directly in the cell instead of using
// the shared string table.
func (c *Cell) SetInlineString(text string) {
	c.SetDataType(CellTypeInlineString)
	// Remove any existing value element
	c.RemoveValue()
	// Set inline string content
	is := c.GetOrCreateInlineString()
	is.SetPlainText(text)
}

// SetError sets the cell to contain an error value
// (e.g., "#DIV/0!", "#N/A", "#VALUE!").
func (c *Cell) SetError(errorValue string) {
	c.SetDataType(CellTypeError)
	cv := c.GetOrCreateCellValue()
	cv.SetValue(errorValue)
}

// SetDateValue sets the cell to a date value stored as an Excel serial number.
// The serial number is calculated from the given time.
// The date1904 parameter specifies whether to use the 1904 date system
// (Mac Excel).
// Note: To display properly, you should also apply a date number format
// to the cell style.
func (c *Cell) SetDateValue(serial float64) {
	c.SetDataType(CellTypeNumber)
	cv := c.GetOrCreateCellValue()
	// 64-bit float size
	cv.SetValue(
		strconv.FormatFloat(
			serial,
			'f',
			-1,
			64, //nolint:mnd,revive // add-constant: 64-bit float size
		),
	)
}

// GetNumber returns the numeric value of the cell.
// Returns 0 and false if the cell doesn't contain a valid number.
func (c *Cell) GetNumber() (float64, bool) {
	val := c.Value()
	if val == "" {
		return 0, false
	}
	//nolint:revive // add-constant: 64-bit float size is standard
	num, err := strconv.ParseFloat(
		val,
		64,
	) //nolint:mnd
	if err != nil {
		return 0, false
	}

	return num, true
}

// GetBoolean returns the boolean value of the cell.
// Returns false and false if the cell doesn't contain a boolean.
//
//nolint:revive // enforce-repeated-arg-type-style: named returns for clarity
func (c *Cell) GetBoolean() (value bool, ok bool) {
	if c.DataType() != CellTypeBoolean {
		return false, false
	}
	val := c.Value()
	switch val {
	case "1", "true", "TRUE":
		return true, true
	case "0", "false", "FALSE":
		return false, true
	default:
		return false, false
	}
}

// GetInlineString returns the inline string content of the cell.
// Returns an empty string if the cell doesn't have an inline string.
func (c *Cell) GetInlineString() string {
	is := c.InlineString()
	if is == nil {
		return ""
	}

	return is.PlainText()
}

// GetSharedStringIndex returns the shared string index from the cell.
// Returns -1 and false if the cell doesn't contain a shared string reference.
func (c *Cell) GetSharedStringIndex() (int, bool) {
	if c.DataType() != CellTypeSharedString {
		return -1, false
	}
	val := c.Value()
	if val == "" {
		return -1, false
	}
	index, err := strconv.Atoi(val)
	if err != nil {
		return -1, false
	}

	return index, true
}

// GetError returns the error value of the cell.
// Returns an empty string and false if the cell doesn't contain an error.
func (c *Cell) GetError() (string, bool) {
	if c.DataType() != CellTypeError {
		return "", false
	}
	val := c.Value()

	return val, val != ""
}

// ValueType returns the type of value stored in the cell.
// This is useful for determining how to read the cell's value.
func (c *Cell) ValueType() CellType {
	return c.DataType()
}

// IsEmpty returns true if the cell has no value, formula, or inline string.
func (c *Cell) IsEmpty() bool {
	return c.CellValue() == nil &&
		c.CellFormula() == nil &&
		c.InlineString() == nil
}

// SetFormula sets the cell formula.
func (c *Cell) SetFormula(formula string) {
	cf := c.GetOrCreateCellFormula()
	cf.SetFormula(formula)
}

// HasFormula returns whether the cell has a formula.
func (c *Cell) HasFormula() bool {
	return c.CellFormula() != nil
}

// Formula returns the cell formula text, or empty string if no formula.
func (c *Cell) Formula() string {
	cf := c.CellFormula()
	if cf == nil {
		return ""
	}

	return cf.Formula()
}

// RemoveFormula removes the cell formula.
func (c *Cell) RemoveFormula() {
	cf := c.CellFormula()
	if cf != nil {
		c.RemoveChild(cf)
	}
}

// RemoveValue removes the cell value.
func (c *Cell) RemoveValue() {
	cv := c.CellValue()
	if cv != nil {
		c.RemoveChild(cv)
	}
}

// Clear removes both formula and value from the cell.
func (c *Cell) Clear() {
	c.RemoveFormula()
	c.RemoveValue()
	is := c.InlineString()
	if is != nil {
		c.RemoveChild(is)
	}
}
