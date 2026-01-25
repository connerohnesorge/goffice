// This file implements the high-level Cell API for working with cell values.
package spreadsheet

import (
	"strconv"
	"time"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// bitSize64 is the bit size for 64-bit floating point parsing.
const bitSize64ForCell = 64

// Cell represents a high-level wrapper around a cell with convenient methods.
type Cell struct {
	sheet    *Sheet
	elemCell *elements.Cell
	row      uint32
	col      uint32
}

// newCell creates a new Cell wrapper.
func newCell(
	sheet *Sheet,
	elemCell *elements.Cell,
	row, col uint32,
) *Cell {
	return &Cell{
		sheet:    sheet,
		elemCell: elemCell,
		row:      row,
		col:      col,
	}
}

// Reference returns the A1-style cell reference.
func (c *Cell) Reference() string {
	return c.elemCell.Reference()
}

// Row returns the 1-based row index.
func (c *Cell) Row() uint32 {
	return c.row
}

// Column returns the 1-based column index.
func (c *Cell) Column() uint32 {
	return c.col
}

// SetNumber sets the cell value to a number.
func (c *Cell) SetNumber(value float64) {
	c.elemCell.SetDataType(
		elements.CellTypeNumber,
	)
	cv := c.elemCell.GetOrCreateCellValue()
	cv.SetValue(
		strconv.FormatFloat(
			value,
			'f',
			-1,
			bitSize64ForCell,
		),
	)
}

// SetString sets the cell value to a string using shared strings.
// This is more efficient for repeated strings.
func (c *Cell) SetString(value string) {
	// Get or create the shared strings table
	workbookPart := c.getWorkbookPart()
	if workbookPart == nil {
		// Fall back to inline string
		c.SetInlineString(value)

		return
	}

	ssp := workbookPart.SharedStringTablePart()
	if ssp == nil {
		var err error
		ssp, err = workbookPart.AddSharedStringTablePart()
		if err != nil {
			// Fall back to inline string
			c.SetInlineString(value)

			return
		}
	}

	// Add the string and get its index
	index := ssp.AddString(value)

	// Set the cell to reference the shared string
	c.elemCell.SetDataType(
		elements.CellTypeSharedString,
	)
	cv := c.elemCell.GetOrCreateCellValue()
	cv.SetValue(strconv.Itoa(index))
}

// SetInlineString sets the cell value to an inline string.
// This stores the string directly in the cell, not in the shared strings table.
//
//nolint:revive // flag-parameter: value is a string, not a bool flag
func (c *Cell) SetInlineString(value string) {
	c.elemCell.SetDataType(
		elements.CellTypeInlineString,
	)

	// Remove any existing value element
	if cv := c.elemCell.CellValue(); cv != nil {
		c.elemCell.RemoveChild(cv)
	}

	// Create inline string
	is := c.elemCell.GetOrCreateInlineString()
	is.SetPlainText(value)
}

// SetBoolean sets the cell value to a boolean.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Cell) SetBoolean(value bool) {
	c.elemCell.SetDataType(
		elements.CellTypeBoolean,
	)
	cv := c.elemCell.GetOrCreateCellValue()
	if value {
		cv.SetValue("1")
	} else {
		cv.SetValue("0")
	}
}

// SetDate sets the cell value to a date.
// The value is stored as an Excel serial date number.
// Note: You should also set a date format style for proper display.
func (c *Cell) SetDate(value time.Time) {
	c.elemCell.SetDataType(
		elements.CellTypeNumber,
	)
	// Use 1900 date system (most common)
	serial := ExcelDateSerial(value, false)
	cv := c.elemCell.GetOrCreateCellValue()
	cv.SetValue(
		strconv.FormatFloat(
			serial,
			'f',
			-1,
			bitSize64ForCell,
		),
	)
}

// SetFormula sets a formula for the cell.
func (c *Cell) SetFormula(formula string) {
	cf := c.elemCell.GetOrCreateCellFormula()
	cf.SetFormula(formula)

	// Clear any cached value since formula needs recalculation
	if cv := c.elemCell.CellValue(); cv != nil {
		cv.SetValue("")
	}
}

// SetStyle sets the cell style.
func (c *Cell) SetStyle(style *Style) {
	if style == nil {
		c.elemCell.SetStyleIndex(0)

		return
	}

	// Get the style index from the styles part
	styleIndex := style.index
	c.elemCell.SetStyleIndex(styleIndex)
}

// GetNumber returns the cell value as a number.
// Returns 0 if the cell doesn't contain a number or is empty.
//
//nolint:revive // optimize-operands-order: clearer with value check first
func (c *Cell) GetNumber() float64 {
	value := c.elemCell.Value()
	if value == "" {
		return 0
	}

	num, err := strconv.ParseFloat(
		value,
		bitSize64ForCell,
	)
	if err != nil {
		return 0
	}

	return num
}

// GetString returns the cell value as a string.
func (c *Cell) GetString() string {
	switch c.elemCell.DataType() {
	case elements.CellTypeSharedString:
		// Look up the shared string
		workbookPart := c.getWorkbookPart()
		if workbookPart == nil {
			return ""
		}
		ssp := workbookPart.SharedStringTablePart()
		if ssp == nil {
			return ""
		}
		index, err := strconv.Atoi(
			c.elemCell.Value(),
		)
		if err != nil {
			return ""
		}

		return ssp.GetString(index)

	case elements.CellTypeInlineString:
		is := c.elemCell.InlineString()
		if is == nil {
			return ""
		}

		return is.PlainText()

	case elements.CellTypeBoolean:
		if c.elemCell.Value() == "1" {
			return booleanTrueString
		}

		return booleanFalseString

	case elements.CellTypeDate:
		// Date values are stored as strings in ISO 8601 format
		return c.elemCell.Value()

	case elements.CellTypeError:
		// Error values like #DIV/0!, #N/A, etc.
		return c.elemCell.Value()

	case elements.CellTypeNumber:
		// Numeric values
		return c.elemCell.Value()

	case elements.CellTypeFormulaString:
		// Formula string results
		return c.elemCell.Value()

	default:
		return ""
	}
}

// GetBoolean returns the cell value as a boolean.
// Returns false if the cell doesn't contain a boolean or is empty.
func (c *Cell) GetBoolean() bool {
	if c.elemCell.DataType() != elements.CellTypeBoolean {
		return false
	}

	return c.elemCell.Value() == "1"
}

// GetDate returns the cell value as a time.Time.
// Returns the zero time if the cell doesn't contain a valid date serial.
func (c *Cell) GetDate() time.Time {
	num := c.GetNumber()
	if num == 0 {
		return time.Time{}
	}

	// Use 1900 date system (most common)
	return DateFromExcelSerial(num, false)
}

// GetFormula returns the cell's formula, or an empty string if none.
func (c *Cell) GetFormula() string {
	cf := c.elemCell.CellFormula()
	if cf == nil {
		return ""
	}

	return cf.Formula()
}

// IsEmpty returns true if the cell has no value.
//
//nolint:revive // optimize-operands-order: clearer with value check first
func (c *Cell) IsEmpty() bool {
	cv := c.elemCell.CellValue()
	is := c.elemCell.InlineString()
	cf := c.elemCell.CellFormula()

	return (cv == nil || cv.Value() == "") &&
		is == nil &&
		cf == nil
}

// DataType returns the cell's data type.
func (c *Cell) DataType() elements.CellType {
	return c.elemCell.DataType()
}

// StyleIndex returns the cell's style index.
func (c *Cell) StyleIndex() uint32 {
	return c.elemCell.StyleIndex()
}

// Clear removes the cell's value, formula, and inline string.
func (c *Cell) Clear() {
	if cv := c.elemCell.CellValue(); cv != nil {
		c.elemCell.RemoveChild(cv)
	}
	if is := c.elemCell.InlineString(); is != nil {
		c.elemCell.RemoveChild(is)
	}
	if cf := c.elemCell.CellFormula(); cf != nil {
		c.elemCell.RemoveChild(cf)
	}
	c.elemCell.SetDataType(
		elements.CellTypeNumber,
	)
}

// getWorkbookPart returns the workbook part for this cell's document.
func (c *Cell) getWorkbookPart() *parts.WorkbookPart {
	if c.sheet == nil || c.sheet.doc == nil {
		return nil
	}

	return c.sheet.doc.WorkbookPart()
}
