// Package spreadsheet provides SpreadsheetML support for Excel documents.
package spreadsheet

import (
	"errors"
)

// ShiftType represents the type of shift operation.
type ShiftType int

const (
	// ShiftInsertRows indicates rows are being inserted.
	ShiftInsertRows ShiftType = iota
	// ShiftDeleteRows indicates rows are being deleted.
	ShiftDeleteRows
	// ShiftInsertColumns indicates columns are being inserted.
	ShiftInsertColumns
	// ShiftDeleteColumns indicates columns are being deleted.
	ShiftDeleteColumns
)

// ShiftOperation describes a row/column insertion or deletion.
type ShiftOperation struct {
	// Type is the type of shift operation.
	Type ShiftType
	// Index is the 1-based starting row/column where the operation occurs.
	Index uint32
	// Count is the number of rows/columns being inserted or deleted.
	Count uint32
	// Worksheet is the target worksheet name (empty = current worksheet).
	Worksheet string
}

// ErrReferenceDeleted is returned when a reference points to deleted cells.
var ErrReferenceDeleted = errors.New(
	"reference points to deleted cells",
)

// CellReferenceUpdater updates cell references based on shift operations.
type CellReferenceUpdater struct {
	operation    ShiftOperation
	currentSheet string // Current worksheet for relative refs
}

// NewCellReferenceUpdater creates a new cell reference updater.
func NewCellReferenceUpdater(
	operation ShiftOperation,
	currentSheet string,
) *CellReferenceUpdater {
	return &CellReferenceUpdater{
		operation:    operation,
		currentSheet: currentSheet,
	}
}

// ShouldUpdate returns true if the reference should be updated.
// It checks whether the reference is on the affected worksheet and
// whether it's in the affected range.
func (u *CellReferenceUpdater) ShouldUpdate(
	ref CellRef,
) bool {
	// Check worksheet match
	refSheet := ref.Sheet
	if refSheet == "" {
		refSheet = u.currentSheet
	}
	targetSheet := u.operation.Worksheet
	if targetSheet == "" {
		targetSheet = u.currentSheet
	}
	if refSheet != targetSheet {
		return false
	}

	idx := int(u.operation.Index)

	switch u.operation.Type {
	case ShiftInsertRows:
		// Update if row >= index (relative refs only)
		return ref.Row >= idx && !ref.AbsRow
	case ShiftDeleteRows:
		// Update if row >= index (relative refs only)
		return ref.Row >= idx && !ref.AbsRow
	case ShiftInsertColumns:
		// Update if col >= index (relative refs only)
		return ref.Col >= idx && !ref.AbsCol
	case ShiftDeleteColumns:
		// Update if col >= index (relative refs only)
		return ref.Col >= idx && !ref.AbsCol
	}

	return false
}

// UpdateReference applies the shift to a cell reference.
// Returns ErrReferenceDeleted if the reference is in the deleted range.
//
//nolint:revive // cognitive-complexity: shift logic requires branches
func (u *CellReferenceUpdater) UpdateReference(
	ref CellRef,
) (CellRef, error) {
	// Check worksheet match
	refSheet := ref.Sheet
	if refSheet == "" {
		refSheet = u.currentSheet
	}
	targetSheet := u.operation.Worksheet
	if targetSheet == "" {
		targetSheet = u.currentSheet
	}
	if refSheet != targetSheet {
		// Different worksheet, no update needed
		return ref, nil
	}

	idx := int(u.operation.Index)
	count := int(u.operation.Count)

	switch u.operation.Type {
	case ShiftInsertRows:
		// Insert rows: shift references at or after index down
		if ref.Row >= idx {
			if ref.AbsRow {
				// Absolute reference: don't shift
				return ref, nil
			}
			// Shift down
			ref.Row += count
		}

		return ref, nil

	case ShiftDeleteRows:
		// Delete rows: check if in deleted range
		deleteEnd := idx + count - 1
		if ref.Row >= idx &&
			ref.Row <= deleteEnd {
			// Reference is in deleted range
			return ref, ErrReferenceDeleted
		}
		if ref.Row > deleteEnd {
			if ref.AbsRow {
				// Absolute reference: don't shift
				return ref, nil
			}
			// Shift up
			ref.Row -= count
		}

		return ref, nil

	case ShiftInsertColumns:
		// Insert columns: shift references at or after index right
		if ref.Col >= idx {
			if ref.AbsCol {
				// Absolute reference: don't shift
				return ref, nil
			}
			// Shift right
			ref.Col += count
		}

		return ref, nil

	case ShiftDeleteColumns:
		// Delete columns: check if in deleted range
		deleteEnd := idx + count - 1
		if ref.Col >= idx &&
			ref.Col <= deleteEnd {
			// Reference is in deleted range
			return ref, ErrReferenceDeleted
		}
		if ref.Col > deleteEnd {
			if ref.AbsCol {
				// Absolute reference: don't shift
				return ref, nil
			}
			// Shift left
			ref.Col -= count
		}

		return ref, nil
	}

	return ref, nil
}

// UpdateRange applies the shift to a range reference.
// Returns ErrReferenceDeleted if the entire range is deleted.
// If only part of the range is deleted, it adjusts the range boundaries.
//
//nolint:revive,gocritic // cognitive-complexity: range shift logic requires branches; hugeParam: RangeRef size is acceptable
func (u *CellReferenceUpdater) UpdateRange(
	ref RangeRef,
) (RangeRef, error) {
	// Update start cell
	start, errStart := u.UpdateReference(
		ref.Start,
	)

	// Update end cell
	end, errEnd := u.UpdateReference(ref.End)

	// Handle deletion cases
	if errStart == ErrReferenceDeleted &&
		errEnd == ErrReferenceDeleted {
		// Entire range deleted
		return ref, ErrReferenceDeleted
	}

	idx := int(u.operation.Index)
	count := int(u.operation.Count)

	// Check worksheet match for partial deletion handling
	refSheet := ref.Start.Sheet
	if refSheet == "" {
		refSheet = u.currentSheet
	}
	targetSheet := u.operation.Worksheet
	if targetSheet == "" {
		targetSheet = u.currentSheet
	}

	if refSheet == targetSheet {
		switch u.operation.Type {
		case ShiftDeleteRows:
			deleteEnd := idx + count - 1

			// Adjust range boundaries if partially deleted
			if errStart == ErrReferenceDeleted {
				// Start is deleted, move it to first row after deletion
				start = ref.Start
				start.Row = deleteEnd + 1 - count // Account for deletion
				if start.Row < idx {
					start.Row = idx
				}
			}
			if errEnd == ErrReferenceDeleted {
				// End is deleted, move it to last row before deletion
				end = ref.End
				end.Row = idx - 1
			}

		case ShiftDeleteColumns:
			deleteEnd := idx + count - 1

			// Adjust range boundaries if partially deleted
			if errStart == ErrReferenceDeleted {
				// Start is deleted, move it to first column after deletion
				start = ref.Start
				start.Col = deleteEnd + 1 - count // Account for deletion
				if start.Col < idx {
					start.Col = idx
				}
			}
			if errEnd == ErrReferenceDeleted {
				// End is deleted, move it to last column before deletion
				end = ref.End
				end.Col = idx - 1
			}

		case ShiftInsertRows, ShiftInsertColumns:
			// No boundary adjustment needed for insertions
		}
	}

	return RangeRef{Start: start, End: end}, nil
}
