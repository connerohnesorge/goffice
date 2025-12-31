// Package spreadsheet provides SpreadsheetML support for Excel documents.
//
//nolint:revive // file-length-limit: cell reference logic is cohesive
package spreadsheet

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Cell reference constants.
const (
	// MaxColumn is the maximum column number (XFD = 16384).
	MaxColumn = 16384

	// MaxRow is the maximum row number.
	MaxRow = 1048576

	// MinColumn is the minimum column number.
	MinColumn = 1

	// MinRow is the minimum row number.
	MinRow = 1

	// alphabetSize is the number of letters in the English alphabet (A-Z).
	alphabetSize = 26
)

// CellRef represents an Excel cell reference in A1 or R1C1 notation.
// It supports both relative (A1) and absolute ($A$1) references,
// as well as cross-sheet references (Sheet2!A1) and R1C1 notation.
type CellRef struct {
	// Col is the 1-based column number (A=1, B=2, ..., XFD=16384).
	Col int

	// Row is the 1-based row number (1-1048576).
	Row int

	// AbsCol indicates whether the column reference is absolute ($A).
	AbsCol bool

	// AbsRow indicates whether the row reference is absolute ($1).
	AbsRow bool

	// Sheet is the sheet name for cross-sheet references (empty if same sheet).
	Sheet string

	// IsR1C1 indicates whether this reference uses R1C1 notation.
	IsR1C1 bool

	// R1C1RelCol is the R1C1 relative column offset (only if IsR1C1).
	R1C1RelCol int

	// R1C1RelRow is the R1C1 relative row offset (only if IsR1C1).
	R1C1RelRow int
}

// ErrInvalidCellRef is returned when a cell reference string is invalid.
var ErrInvalidCellRef = errors.New(
	"invalid cell reference",
)

// ErrColumnOutOfRange is returned when a column number is out of valid range.
var ErrColumnOutOfRange = errors.New(
	"column number out of range (1-16384)",
)

// ErrRowOutOfRange is returned when a row number is out of valid range.
var ErrRowOutOfRange = errors.New(
	"row number out of range (1-1048576)",
)

// ParseCellRef parses a cell reference string in A1 or R1C1 notation.
// It accepts references like "A1", "$A$1", "Sheet2!A1", "'Sheet Name'!A1",
// "R1C1", "R[-1]C[2]", etc.
//
//nolint:revive // function-length: parsing logic is cohesive
func ParseCellRef(s string) (CellRef, error) {
	if s == "" {
		return CellRef{}, ErrInvalidCellRef
	}

	var ref CellRef

	// Check for sheet name (cross-sheet reference)
	var bangIdx int
	if s[0] == '\'' {
		// Quoted sheet name
		endQuote := strings.Index(s[1:], "'")
		if endQuote == -1 {
			return CellRef{}, ErrInvalidCellRef
		}
		endQuote++ // Adjust for slice offset
		if endQuote+1 >= len(s) ||
			s[endQuote+1] != '!' {
			return CellRef{}, ErrInvalidCellRef
		}
		ref.Sheet = s[1:endQuote]
		bangIdx = endQuote + 1
	} else {
		// Unquoted sheet name or no sheet name
		bangIdx = strings.Index(s, "!")
		if bangIdx != -1 {
			ref.Sheet = s[:bangIdx]
		}
	}

	// Extract the cell reference part (after sheet name if present)
	cellPart := s
	if bangIdx != -1 {
		cellPart = s[bangIdx+1:]
	}

	// Check if this is R1C1 notation (must be "R" followed by digit or '[')
	if cellPart != "" &&
		(cellPart[0] == 'R' || cellPart[0] == 'r') {
		// Only treat as R1C1 if followed by digit or '['
		if len(cellPart) > 1 &&
			(cellPart[1] == '[' || (cellPart[1] >= '0' && cellPart[1] <= '9')) {
			return parseR1C1Ref(
				cellPart,
				ref.Sheet,
			)
		}
	}

	// Parse A1 notation
	return parseA1Ref(cellPart, ref.Sheet)
}

// parseA1Ref parses A1-style notation
//
//nolint:revive // function-length: parsing logic is cohesive
func parseA1Ref(
	s, sheet string,
) (CellRef, error) {
	var ref CellRef
	ref.Sheet = sheet
	pos := 0

	// Check for absolute column marker
	if pos < len(s) && s[pos] == '$' {
		ref.AbsCol = true
		pos++
	}

	// Extract column letters
	colStart := pos
	for pos < len(s) && isColLetter(s[pos]) {
		pos++
	}
	colEnd := pos

	if colStart == colEnd {
		return CellRef{}, ErrInvalidCellRef
	}

	colName := s[colStart:colEnd]
	ref.Col = ColumnIndex(colName)
	if ref.Col < MinColumn ||
		ref.Col > MaxColumn {
		return CellRef{}, ErrColumnOutOfRange
	}

	// Check for absolute row marker
	if pos < len(s) && s[pos] == '$' {
		ref.AbsRow = true
		pos++
	}

	// Extract row number
	if pos >= len(s) {
		return CellRef{}, ErrInvalidCellRef
	}

	rowStr := s[pos:]
	for _, c := range rowStr {
		if c < '0' || c > '9' {
			return CellRef{}, ErrInvalidCellRef
		}
	}

	row, err := strconv.Atoi(rowStr)
	if err != nil || row < MinRow ||
		row > MaxRow {
		if row < MinRow || row > MaxRow {
			return CellRef{}, ErrRowOutOfRange
		}

		return CellRef{}, ErrInvalidCellRef
	}
	ref.Row = row

	return ref, nil
}

// parseR1C1Ref parses R1C1-style notation
//
//nolint:revive // function-length: parsing logic is cohesive
func parseR1C1Ref(
	s, sheet string,
) (CellRef, error) {
	var ref CellRef
	ref.Sheet = sheet
	ref.IsR1C1 = true

	if s == "" || (s[0] != 'R' && s[0] != 'r') {
		return CellRef{}, ErrInvalidCellRef
	}

	pos := 1

	// Parse row part
	if pos < len(s) && s[pos] == '[' {
		// Relative row: R[-1]
		endBracket := strings.Index(s[pos:], "]")
		if endBracket == -1 {
			return CellRef{}, ErrInvalidCellRef
		}
		offsetStr := s[pos+1 : pos+endBracket]
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return CellRef{}, ErrInvalidCellRef
		}
		ref.R1C1RelRow = offset
		pos += endBracket + 1
	} else {
		// Absolute row: R5
		rowStart := pos
		for pos < len(s) && s[pos] >= '0' && s[pos] <= '9' {
			pos++
		}
		if rowStart == pos {
			return CellRef{}, ErrInvalidCellRef
		}
		row, err := strconv.Atoi(s[rowStart:pos])
		if err != nil || row < MinRow || row > MaxRow {
			return CellRef{}, ErrInvalidCellRef
		}
		ref.Row = row
		ref.AbsRow = true
	}

	// Check for column part
	if pos >= len(s) ||
		(s[pos] != 'C' && s[pos] != 'c') {
		return CellRef{}, ErrInvalidCellRef
	}
	pos++

	// Parse column part
	if pos < len(s) && s[pos] == '[' {
		// Relative column: C[-1]
		endBracket := strings.Index(s[pos:], "]")
		if endBracket == -1 {
			return CellRef{}, ErrInvalidCellRef
		}
		offsetStr := s[pos+1 : pos+endBracket]
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return CellRef{}, ErrInvalidCellRef
		}
		ref.R1C1RelCol = offset
		pos += endBracket + 1
	} else {
		// Absolute column: C10
		colStart := pos
		for pos < len(s) && s[pos] >= '0' && s[pos] <= '9' {
			pos++
		}
		if colStart == pos {
			return CellRef{}, ErrInvalidCellRef
		}
		col, err := strconv.Atoi(s[colStart:pos])
		if err != nil || col < MinColumn || col > MaxColumn {
			return CellRef{}, ErrInvalidCellRef
		}
		ref.Col = col
		ref.AbsCol = true
	}

	// Ensure we've consumed the entire string
	if pos != len(s) {
		return CellRef{}, ErrInvalidCellRef
	}

	return ref, nil
}

// MustParseCellRef parses an A1-style cell reference string.
// It panics if the reference is invalid.
// This is useful for tests and initialization code.
func MustParseCellRef(s string) CellRef {
	ref, err := ParseCellRef(s)
	if err != nil {
		panic(
			fmt.Sprintf(
				"invalid cell reference %q: %v",
				s,
				err,
			),
		)
	}

	return ref
}

// NewCellRef creates a new cell reference from column and row numbers.
// Column and row are 1-based.
func NewCellRef(col, row int) CellRef {
	return CellRef{Col: col, Row: row}
}

// NewAbsCellRef creates a new absolute cell reference from column and row
// numbers. Column and row are 1-based.
func NewAbsCellRef(col, row int) CellRef {
	return CellRef{
		Col:    col,
		Row:    row,
		AbsCol: true,
		AbsRow: true,
	}
}

// String returns the string representation of the cell reference.
// Returns A1-style notation by default, or R1C1 notation if IsR1C1 is true.
// For example, CellRef{Col: 1, Row: 1, AbsCol: true, AbsRow: true}
// returns "$A$1" or "Sheet2!$A$1" for cross-sheet references.
//
//nolint:revive // function-length: string building logic is cohesive
func (c CellRef) String() string {
	var sb strings.Builder

	// Add sheet name if present
	if c.Sheet != "" {
		// Quote sheet name if it contains special characters
		if needsQuoting(c.Sheet) {
			sb.WriteByte('\'')
			sb.WriteString(c.Sheet)
			sb.WriteByte('\'')
		} else {
			sb.WriteString(c.Sheet)
		}
		sb.WriteByte('!')
	}

	// Output R1C1 notation if requested
	if c.IsR1C1 {
		sb.WriteByte('R')
		if c.AbsRow {
			sb.WriteString(strconv.Itoa(c.Row))
		} else {
			sb.WriteByte('[')
			sb.WriteString(strconv.Itoa(c.R1C1RelRow))
			sb.WriteByte(']')
		}
		sb.WriteByte('C')
		if c.AbsCol {
			sb.WriteString(strconv.Itoa(c.Col))
		} else {
			sb.WriteByte('[')
			sb.WriteString(strconv.Itoa(c.R1C1RelCol))
			sb.WriteByte(']')
		}

		return sb.String()
	}

	// Output A1 notation
	if c.AbsCol {
		sb.WriteByte('$')
	}
	sb.WriteString(ColumnName(c.Col))

	if c.AbsRow {
		sb.WriteByte('$')
	}
	sb.WriteString(strconv.Itoa(c.Row))

	return sb.String()
}

// needsQuoting returns true if a sheet name needs to be quoted
func needsQuoting(name string) bool {
	// Sheet names need quoting if they contain spaces or special characters
	for _, r := range name {
		if r == ' ' || r == '\'' || r == '!' ||
			r == '(' ||
			r == ')' ||
			r == '[' ||
			r == ']' ||
			r == ':' {
			return true
		}
	}

	return false
}

// R1C1 returns the R1C1-style string representation of the cell reference.
// For example, CellRef{Col: 3, Row: 5} returns "R5C3".
// Absolute references don't affect R1C1 notation (Excel uses
// different syntax for relative R1C1).
func (c CellRef) R1C1() string {
	return fmt.Sprintf("R%dC%d", c.Row, c.Col)
}

// IsAbsolute returns true if both column and row references are absolute.
func (c CellRef) IsAbsolute() bool {
	return c.AbsCol && c.AbsRow
}

// IsRelative returns true if both column and row references are relative.
func (c CellRef) IsRelative() bool {
	return !c.AbsCol && !c.AbsRow
}

// IsMixed returns true if exactly one of column or row reference is absolute.
func (c CellRef) IsMixed() bool {
	return c.AbsCol != c.AbsRow
}

// IsValid returns true if the cell reference is within Excel's valid range.
func (c CellRef) IsValid() bool {
	return c.Col >= MinColumn &&
		c.Col <= MaxColumn &&
		c.Row >= MinRow &&
		c.Row <= MaxRow
}

// WithAbsolute returns a new CellRef with the specified absolute settings.
func (c CellRef) WithAbsolute(
	absCol, absRow bool,
) CellRef {
	return CellRef{
		Col:    c.Col,
		Row:    c.Row,
		AbsCol: absCol,
		AbsRow: absRow,
	}
}

// Offset returns a new CellRef offset by the given column and row amounts.
// Returns an error if the result would be out of range.
func (c CellRef) Offset(
	colOffset, rowOffset int,
) (CellRef, error) {
	newCol := c.Col + colOffset
	newRow := c.Row + rowOffset

	if newCol < MinColumn || newCol > MaxColumn {
		return CellRef{}, ErrColumnOutOfRange
	}
	if newRow < MinRow || newRow > MaxRow {
		return CellRef{}, ErrRowOutOfRange
	}

	return CellRef{
		Col:    newCol,
		Row:    newRow,
		AbsCol: c.AbsCol,
		AbsRow: c.AbsRow,
	}, nil
}

// ColumnName converts a 1-based column number to its A1-style letter
// representation. For example: 1 -> "A", 26 -> "Z", 27 -> "AA",
// 16384 -> "XFD".
//
//nolint:revive // modifies-parameter: col modification is for calculation only
func ColumnName(col int) string {
	if col < MinColumn {
		return ""
	}

	var result strings.Builder
	col-- // Convert to 0-based for calculation

	for col >= 0 {
		result.WriteByte(
			byte('A' + col%alphabetSize),
		)
		col = col/alphabetSize - 1
	}

	// Reverse the string
	s := result.String()
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// ColumnIndex converts an A1-style column letter to its 1-based column number.
// For example: "A" -> 1, "Z" -> 26, "AA" -> 27, "XFD" -> 16384.
// Returns 0 if the column name is invalid.
//
//nolint:revive // modifies-parameter: name is uppercased for normalization
func ColumnIndex(name string) int {
	if name == "" {
		return 0
	}

	name = strings.ToUpper(name)
	col := 0

	for _, c := range name {
		if c < 'A' || c > 'Z' {
			return 0
		}
		col = col*alphabetSize + int(c-'A'+1)
	}

	return col
}

// isColLetter returns true if the byte is a column letter (A-Z or a-z).
func isColLetter(b byte) bool {
	return (b >= 'A' && b <= 'Z') ||
		(b >= 'a' && b <= 'z')
}
