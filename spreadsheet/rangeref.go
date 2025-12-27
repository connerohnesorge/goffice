// Package spreadsheet provides SpreadsheetML support for Excel documents.
//
//nolint:revive // file-length-limit: range reference logic is cohesive
package spreadsheet

import (
	"errors"
	"fmt"
	"iter"
	"strings"
)

// ErrInvalidRangeRef is returned when a range reference string is invalid.
var ErrInvalidRangeRef = errors.New(
	"invalid range reference",
)

// RangeRef represents an Excel cell range reference in A1 notation.
// It can represent a single cell (A1) or a range of cells (A1:C10).
type RangeRef struct {
	// Start is the top-left cell of the range.
	Start CellRef

	// End is the bottom-right cell of the range.
	// For a single cell reference, End equals Start.
	End CellRef
}

// ParseRangeRef parses an A1-style range reference string.
// It accepts references like "A1", "A1:C10", "$A$1:$C$10", etc.
func ParseRangeRef(s string) (RangeRef, error) {
	if s == "" {
		return RangeRef{}, ErrInvalidRangeRef
	}

	// Check if this is a range (contains ':')
	colonIdx := strings.Index(s, ":")
	if colonIdx == -1 {
		// Single cell reference
		start, err := ParseCellRef(s)
		if err != nil {
			return RangeRef{}, fmt.Errorf(
				"%w: %v",
				ErrInvalidRangeRef,
				err,
			)
		}

		return RangeRef{
			Start: start,
			End:   start,
		}, nil
	}

	// Range reference
	startStr := s[:colonIdx]
	endStr := s[colonIdx+1:]

	start, err := ParseCellRef(startStr)
	if err != nil {
		return RangeRef{}, fmt.Errorf(
			"%w: invalid start cell: %v",
			ErrInvalidRangeRef,
			err,
		)
	}

	end, err := ParseCellRef(endStr)
	if err != nil {
		return RangeRef{}, fmt.Errorf(
			"%w: invalid end cell: %v",
			ErrInvalidRangeRef,
			err,
		)
	}

	// Normalize so Start is top-left and End is bottom-right
	r := RangeRef{Start: start, End: end}
	r = r.Normalize()

	return r, nil
}

// MustParseRangeRef parses an A1-style range reference string.
// It panics if the reference is invalid.
// This is useful for tests and initialization code.
func MustParseRangeRef(s string) RangeRef {
	ref, err := ParseRangeRef(s)
	if err != nil {
		panic(
			fmt.Sprintf(
				"invalid range reference %q: %v",
				s,
				err,
			),
		)
	}

	return ref
}

// NewRangeRef creates a new range reference from two cell references.
// The references are normalized so Start is top-left and End is bottom-right.
func NewRangeRef(start, end CellRef) RangeRef {
	r := RangeRef{Start: start, End: end}

	return r.Normalize()
}

// NewSingleCellRange creates a range reference for a single cell.
func NewSingleCellRange(cell CellRef) RangeRef {
	return RangeRef{Start: cell, End: cell}
}

// String returns the A1-style string representation of the range reference.
// For a single cell, returns just the cell reference (e.g., "A1").
// For a range, returns "start:end" (e.g., "A1:C10", "$A$1:$C$10").
func (r RangeRef) String() string {
	if r.IsSingleCell() {
		return r.Start.String()
	}

	return r.Start.String() + ":" + r.End.String()
}

// IsSingleCell returns true if the range represents a single cell.
func (r RangeRef) IsSingleCell() bool {
	return r.Start.Col == r.End.Col &&
		r.Start.Row == r.End.Row
}

// Normalize returns a new RangeRef where Start is the top-left cell
// and End is the bottom-right cell.
func (r RangeRef) Normalize() RangeRef {
	minCol := r.Start.Col
	maxCol := r.End.Col
	if minCol > maxCol {
		minCol, maxCol = maxCol, minCol
	}

	minRow := r.Start.Row
	maxRow := r.End.Row
	if minRow > maxRow {
		minRow, maxRow = maxRow, minRow
	}

	return RangeRef{
		Start: CellRef{
			Col:    minCol,
			Row:    minRow,
			AbsCol: r.Start.AbsCol,
			AbsRow: r.Start.AbsRow,
		},
		End: CellRef{
			Col:    maxCol,
			Row:    maxRow,
			AbsCol: r.End.AbsCol,
			AbsRow: r.End.AbsRow,
		},
	}
}

// Contains returns true if the given cell reference is within this range.
func (r RangeRef) Contains(ref CellRef) bool {
	return ref.Col >= r.Start.Col &&
		ref.Col <= r.End.Col &&
		ref.Row >= r.Start.Row &&
		ref.Row <= r.End.Row
}

// Intersects returns true if this range overlaps with another range.
func (r RangeRef) Intersects(
	other RangeRef,
) bool {
	return r.Start.Col <= other.End.Col &&
		r.End.Col >= other.Start.Col &&
		r.Start.Row <= other.End.Row &&
		r.End.Row >= other.Start.Row
}

// Intersection returns the intersection of two ranges,
// or an error if they don't intersect.
func (r RangeRef) Intersection(
	other RangeRef,
) (RangeRef, error) {
	if !r.Intersects(other) {
		return RangeRef{}, errors.New(
			"ranges do not intersect",
		)
	}

	return RangeRef{
		Start: CellRef{
			Col: max(
				r.Start.Col,
				other.Start.Col,
			),
			Row: max(
				r.Start.Row,
				other.Start.Row,
			),
		},
		End: CellRef{
			Col: min(r.End.Col, other.End.Col),
			Row: min(r.End.Row, other.End.Row),
		},
	}, nil
}

// Width returns the number of columns in the range.
func (r RangeRef) Width() int {
	return r.End.Col - r.Start.Col + 1
}

// Height returns the number of rows in the range.
func (r RangeRef) Height() int {
	return r.End.Row - r.Start.Row + 1
}

// Size returns the total number of cells in the range.
func (r RangeRef) Size() int {
	return r.Width() * r.Height()
}

// Cells returns an iterator over all cell references in the range.
// Cells are yielded row by row, from left to right, top to bottom.
func (r RangeRef) Cells() iter.Seq[CellRef] {
	return func(yield func(CellRef) bool) {
		for row := r.Start.Row; row <= r.End.Row; row++ {
			for col := r.Start.Col; col <= r.End.Col; col++ {
				cell := CellRef{
					Col: col,
					Row: row,
				}
				if !yield(cell) {
					return
				}
			}
		}
	}
}

// Rows returns an iterator over each row in the range.
// Each yielded value is itself an iterator over cells in that row.
func (r RangeRef) Rows() iter.Seq2[int, iter.Seq[CellRef]] {
	return func(yield func(int, iter.Seq[CellRef]) bool) {
		for row := r.Start.Row; row <= r.End.Row; row++ {
			rowCells := func(yield func(CellRef) bool) {
				for col := r.Start.Col; col <= r.End.Col; col++ {
					cell := CellRef{
						Col: col,
						Row: row,
					}
					if !yield(cell) {
						return
					}
				}
			}
			if !yield(row, rowCells) {
				return
			}
		}
	}
}

// Columns returns an iterator over each column in the range.
// Each yielded value is itself an iterator over cells in that column.
func (r RangeRef) Columns() iter.Seq2[int, iter.Seq[CellRef]] {
	return func(yield func(int, iter.Seq[CellRef]) bool) {
		for col := r.Start.Col; col <= r.End.Col; col++ {
			colCells := func(yield func(CellRef) bool) {
				for row := r.Start.Row; row <= r.End.Row; row++ {
					cell := CellRef{
						Col: col,
						Row: row,
					}
					if !yield(cell) {
						return
					}
				}
			}
			if !yield(col, colCells) {
				return
			}
		}
	}
}

// TopRow returns a RangeRef representing just the top row of this range.
func (r RangeRef) TopRow() RangeRef {
	return RangeRef{
		Start: r.Start,
		End: CellRef{
			Col:    r.End.Col,
			Row:    r.Start.Row,
			AbsCol: r.End.AbsCol,
			AbsRow: r.Start.AbsRow,
		},
	}
}

// BottomRow returns a RangeRef representing just the bottom row of this range.
func (r RangeRef) BottomRow() RangeRef {
	return RangeRef{
		Start: CellRef{
			Col:    r.Start.Col,
			Row:    r.End.Row,
			AbsCol: r.Start.AbsCol,
			AbsRow: r.End.AbsRow,
		},
		End: r.End,
	}
}

// LeftColumn returns a RangeRef representing
// just the left column of this range.
func (r RangeRef) LeftColumn() RangeRef {
	return RangeRef{
		Start: r.Start,
		End: CellRef{
			Col:    r.Start.Col,
			Row:    r.End.Row,
			AbsCol: r.Start.AbsCol,
			AbsRow: r.End.AbsRow,
		},
	}
}

// RightColumn returns a RangeRef representing just the right column
// of this range.
func (r RangeRef) RightColumn() RangeRef {
	return RangeRef{
		Start: CellRef{
			Col:    r.End.Col,
			Row:    r.Start.Row,
			AbsCol: r.End.AbsCol,
			AbsRow: r.Start.AbsRow,
		},
		End: r.End,
	}
}

// Offset returns a new RangeRef offset by the given column and row amounts.
// Returns an error if the result would be out of range.
func (r RangeRef) Offset(
	colOffset, rowOffset int,
) (RangeRef, error) {
	start, err := r.Start.Offset(
		colOffset,
		rowOffset,
	)
	if err != nil {
		return RangeRef{}, err
	}

	end, err := r.End.Offset(colOffset, rowOffset)
	if err != nil {
		return RangeRef{}, err
	}

	return RangeRef{Start: start, End: end}, nil
}

// Expand returns a new RangeRef expanded by the given amounts
// in each direction.
// Negative values shrink the range.
// Returns error if result would be out of range or have zero/negative size.
func (r RangeRef) Expand(
	left, right, top, bottom int,
) (RangeRef, error) {
	newStartCol := r.Start.Col - left
	newStartRow := r.Start.Row - top
	newEndCol := r.End.Col + right
	newEndRow := r.End.Row + bottom

	if newStartCol < MinColumn ||
		newEndCol > MaxColumn {
		return RangeRef{}, ErrColumnOutOfRange
	}
	if newStartRow < MinRow ||
		newEndRow > MaxRow {
		return RangeRef{}, ErrRowOutOfRange
	}
	if newStartCol > newEndCol ||
		newStartRow > newEndRow {
		return RangeRef{}, errors.New(
			"expansion would result in invalid range",
		)
	}

	return RangeRef{
		Start: CellRef{
			Col:    newStartCol,
			Row:    newStartRow,
			AbsCol: r.Start.AbsCol,
			AbsRow: r.Start.AbsRow,
		},
		End: CellRef{
			Col:    newEndCol,
			Row:    newEndRow,
			AbsCol: r.End.AbsCol,
			AbsRow: r.End.AbsRow,
		},
	}, nil
}

// IsValid returns true if the range is within Excel's valid bounds.
func (r RangeRef) IsValid() bool {
	return r.Start.IsValid() && r.End.IsValid()
}
