package spreadsheet

import (
	"testing"
)

// TestCellReferenceUpdaterInsertRows tests row insertion logic.
func TestCellReferenceUpdaterInsertRows(
	t *testing.T,
) {
	tests := []struct {
		name     string
		ref      CellRef
		index    uint32
		count    uint32
		expected CellRef
		wantErr  bool
	}{
		// Before insert point
		{
			name: "before insert",
			ref: CellRef{
				Col:    1,
				Row:    2,
				AbsCol: false,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    1,
				Row:    2,
				AbsCol: false,
				AbsRow: false,
			},
		},
		// At insert point
		{
			name: "at insert",
			ref: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: false,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    1,
				Row:    8,
				AbsCol: false,
				AbsRow: false,
			},
		},
		// After insert point
		{
			name: "after insert",
			ref: CellRef{
				Col:    1,
				Row:    10,
				AbsCol: false,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    1,
				Row:    13,
				AbsCol: false,
				AbsRow: false,
			},
		},
		// Absolute row not shifted
		{
			name: "absolute row",
			ref: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: false,
				AbsRow: true,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: false,
				AbsRow: true,
			},
		},
		// Absolute col shouldn't matter
		{
			name: "absolute col",
			ref: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: true,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    1,
				Row:    8,
				AbsCol: true,
				AbsRow: false,
			},
		},
		// Fully absolute
		{
			name: "fully absolute",
			ref: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: true,
				AbsRow: true,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: true,
				AbsRow: true,
			},
		},
		// Mixed absolute
		{
			name: "mixed $A5",
			ref: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: true,
				AbsRow: false,
			},
			index: 3,
			count: 2,
			expected: CellRef{
				Col:    1,
				Row:    7,
				AbsCol: true,
				AbsRow: false,
			},
		},
		{
			name: "mixed A$5",
			ref: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: false,
				AbsRow: true,
			},
			index: 3,
			count: 2,
			expected: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: false,
				AbsRow: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftInsertRows,
				Index: tt.index,
				Count: tt.count,
			}
			updater := NewCellReferenceUpdater(
				op,
				"Sheet1",
			)
			result, err := updater.UpdateReference(
				tt.ref,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"UpdateReference() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if tt.wantErr {
				return
			}
			if result.Col != tt.expected.Col ||
				result.Row != tt.expected.Row ||
				result.AbsCol != tt.expected.AbsCol ||
				result.AbsRow != tt.expected.AbsRow {
				t.Errorf(
					"UpdateReference() = %+v, want %+v",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestCellReferenceUpdaterDeleteRows tests row deletion logic.
func TestCellReferenceUpdaterDeleteRows(
	t *testing.T,
) {
	tests := []struct {
		name     string
		ref      CellRef
		index    uint32
		count    uint32
		expected CellRef
		wantErr  bool
	}{
		// Before delete range
		{
			name: "before delete",
			ref: CellRef{
				Col:    1,
				Row:    2,
				AbsCol: false,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    1,
				Row:    2,
				AbsCol: false,
				AbsRow: false,
			},
		},
		// In delete range
		{
			name: "in delete range start",
			ref: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: false,
				AbsRow: false,
			},
			index:   5,
			count:   3,
			wantErr: true,
		},
		{
			name: "in delete range middle",
			ref: CellRef{
				Col:    1,
				Row:    6,
				AbsCol: false,
				AbsRow: false,
			},
			index:   5,
			count:   3,
			wantErr: true,
		},
		{
			name: "in delete range end",
			ref: CellRef{
				Col:    1,
				Row:    7,
				AbsCol: false,
				AbsRow: false,
			},
			index:   5,
			count:   3,
			wantErr: true,
		},
		// After delete range
		{
			name: "after delete",
			ref: CellRef{
				Col:    1,
				Row:    10,
				AbsCol: false,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    1,
				Row:    7,
				AbsCol: false,
				AbsRow: false,
			},
		},
		// Absolute row not shifted
		{
			name: "absolute row after delete",
			ref: CellRef{
				Col:    1,
				Row:    10,
				AbsCol: false,
				AbsRow: true,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    1,
				Row:    10,
				AbsCol: false,
				AbsRow: true,
			},
		},
		// Absolute row in delete range - still gets deleted per Excel behavior
		{
			name: "absolute row in delete range",
			ref: CellRef{
				Col:    1,
				Row:    5,
				AbsCol: false,
				AbsRow: true,
			},
			index:   5,
			count:   3,
			wantErr: true, // Even absolute refs get deleted if they're in the deleted range
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftDeleteRows,
				Index: tt.index,
				Count: tt.count,
			}
			updater := NewCellReferenceUpdater(
				op,
				"Sheet1",
			)
			result, err := updater.UpdateReference(
				tt.ref,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"UpdateReference() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if tt.wantErr {
				return
			}
			if result.Col != tt.expected.Col ||
				result.Row != tt.expected.Row ||
				result.AbsCol != tt.expected.AbsCol ||
				result.AbsRow != tt.expected.AbsRow {
				t.Errorf(
					"UpdateReference() = %+v, want %+v",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestCellReferenceUpdaterInsertColumns tests column insertion logic.
func TestCellReferenceUpdaterInsertColumns(
	t *testing.T,
) {
	tests := []struct {
		name     string
		ref      CellRef
		index    uint32
		count    uint32
		expected CellRef
		wantErr  bool
	}{
		{
			name: "before insert",
			ref: CellRef{
				Col:    2,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    2,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
		},
		{
			name: "at insert",
			ref: CellRef{
				Col:    5,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    8,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
		},
		{
			name: "after insert",
			ref: CellRef{
				Col:    10,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    13,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
		},
		{
			name: "absolute col",
			ref: CellRef{
				Col:    5,
				Row:    1,
				AbsCol: true,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    5,
				Row:    1,
				AbsCol: true,
				AbsRow: false,
			},
		},
		{
			name: "mixed A$5",
			ref: CellRef{
				Col:    5,
				Row:    1,
				AbsCol: false,
				AbsRow: true,
			},
			index: 3,
			count: 2,
			expected: CellRef{
				Col:    7,
				Row:    1,
				AbsCol: false,
				AbsRow: true,
			},
		},
		{
			name: "mixed $A5",
			ref: CellRef{
				Col:    5,
				Row:    1,
				AbsCol: true,
				AbsRow: false,
			},
			index: 3,
			count: 2,
			expected: CellRef{
				Col:    5,
				Row:    1,
				AbsCol: true,
				AbsRow: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftInsertColumns,
				Index: tt.index,
				Count: tt.count,
			}
			updater := NewCellReferenceUpdater(
				op,
				"Sheet1",
			)
			result, err := updater.UpdateReference(
				tt.ref,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"UpdateReference() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if tt.wantErr {
				return
			}
			if result.Col != tt.expected.Col ||
				result.Row != tt.expected.Row ||
				result.AbsCol != tt.expected.AbsCol ||
				result.AbsRow != tt.expected.AbsRow {
				t.Errorf(
					"UpdateReference() = %+v, want %+v",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestCellReferenceUpdaterDeleteColumns tests column deletion logic.
func TestCellReferenceUpdaterDeleteColumns(
	t *testing.T,
) {
	tests := []struct {
		name     string
		ref      CellRef
		index    uint32
		count    uint32
		expected CellRef
		wantErr  bool
	}{
		{
			name: "before delete",
			ref: CellRef{
				Col:    2,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    2,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
		},
		{
			name: "in delete range",
			ref: CellRef{
				Col:    5,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
			index:   5,
			count:   3,
			wantErr: true,
		},
		{
			name: "after delete",
			ref: CellRef{
				Col:    10,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    7,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
		},
		{
			name: "absolute col after delete",
			ref: CellRef{
				Col:    10,
				Row:    1,
				AbsCol: true,
				AbsRow: false,
			},
			index: 5,
			count: 3,
			expected: CellRef{
				Col:    10,
				Row:    1,
				AbsCol: true,
				AbsRow: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftDeleteColumns,
				Index: tt.index,
				Count: tt.count,
			}
			updater := NewCellReferenceUpdater(
				op,
				"Sheet1",
			)
			result, err := updater.UpdateReference(
				tt.ref,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"UpdateReference() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if tt.wantErr {
				return
			}
			if result.Col != tt.expected.Col ||
				result.Row != tt.expected.Row ||
				result.AbsCol != tt.expected.AbsCol ||
				result.AbsRow != tt.expected.AbsRow {
				t.Errorf(
					"UpdateReference() = %+v, want %+v",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestCellReferenceUpdaterCrossSheet tests cross-sheet scenarios.
func TestCellReferenceUpdaterCrossSheet(
	t *testing.T,
) {
	tests := []struct {
		name         string
		ref          CellRef
		index        uint32
		count        uint32
		targetSheet  string
		currentSheet string
		expected     CellRef
		wantErr      bool
	}{
		{
			name: "same sheet updated",
			ref: CellRef{
				Col:   1,
				Row:   5,
				Sheet: "Sheet1",
			},
			index:        3,
			count:        2,
			targetSheet:  "Sheet1",
			currentSheet: "Sheet1",
			expected: CellRef{
				Col:   1,
				Row:   7,
				Sheet: "Sheet1",
			},
		},
		{
			name: "different sheet not updated",
			ref: CellRef{
				Col:   1,
				Row:   5,
				Sheet: "Sheet2",
			},
			index:        3,
			count:        2,
			targetSheet:  "Sheet1",
			currentSheet: "Sheet1",
			expected: CellRef{
				Col:   1,
				Row:   5,
				Sheet: "Sheet2",
			},
		},
		{
			name:         "no sheet name uses current",
			ref:          CellRef{Col: 1, Row: 5},
			index:        3,
			count:        2,
			targetSheet:  "Sheet1",
			currentSheet: "Sheet1",
			expected:     CellRef{Col: 1, Row: 7},
		},
		{
			name:         "no sheet name different current",
			ref:          CellRef{Col: 1, Row: 5},
			index:        3,
			count:        2,
			targetSheet:  "Sheet1",
			currentSheet: "Sheet2",
			expected:     CellRef{Col: 1, Row: 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:      ShiftInsertRows,
				Index:     tt.index,
				Count:     tt.count,
				Worksheet: tt.targetSheet,
			}
			updater := NewCellReferenceUpdater(
				op,
				tt.currentSheet,
			)
			result, err := updater.UpdateReference(
				tt.ref,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"UpdateReference() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if tt.wantErr {
				return
			}
			if result.Col != tt.expected.Col ||
				result.Row != tt.expected.Row ||
				result.Sheet != tt.expected.Sheet {
				t.Errorf(
					"UpdateReference() = %+v, want %+v",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestRangeReferenceUpdaterInsertRows tests range update for row insertion.
func TestRangeReferenceUpdaterInsertRows(
	t *testing.T,
) {
	tests := []struct {
		name     string
		rangeRef RangeRef
		index    uint32
		count    uint32
		expected RangeRef
		wantErr  bool
	}{
		{
			name: "range before insert",
			rangeRef: RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End:   CellRef{Col: 1, Row: 3},
			},
			index: 5,
			count: 2,
			expected: RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End:   CellRef{Col: 1, Row: 3},
			},
		},
		{
			name: "range after insert",
			rangeRef: RangeRef{
				Start: CellRef{Col: 1, Row: 10},
				End:   CellRef{Col: 1, Row: 15},
			},
			index: 5,
			count: 2,
			expected: RangeRef{
				Start: CellRef{Col: 1, Row: 12},
				End:   CellRef{Col: 1, Row: 17},
			},
		},
		{
			name: "range spanning insert",
			rangeRef: RangeRef{
				Start: CellRef{Col: 1, Row: 3},
				End:   CellRef{Col: 1, Row: 10},
			},
			index: 5,
			count: 2,
			expected: RangeRef{
				Start: CellRef{Col: 1, Row: 3},
				End:   CellRef{Col: 1, Row: 12},
			},
		},
		{
			name: "range absolute",
			rangeRef: RangeRef{
				Start: CellRef{
					Col:    1,
					Row:    3,
					AbsCol: true,
					AbsRow: true,
				},
				End: CellRef{
					Col:    1,
					Row:    10,
					AbsCol: true,
					AbsRow: true,
				},
			},
			index: 5,
			count: 2,
			expected: RangeRef{
				Start: CellRef{
					Col:    1,
					Row:    3,
					AbsCol: true,
					AbsRow: true,
				},
				End: CellRef{
					Col:    1,
					Row:    10,
					AbsCol: true,
					AbsRow: true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftInsertRows,
				Index: tt.index,
				Count: tt.count,
			}
			updater := NewCellReferenceUpdater(
				op,
				"Sheet1",
			)
			result, err := updater.UpdateRange(
				tt.rangeRef,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"UpdateRange() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if tt.wantErr {
				return
			}
			if result.Start.Row != tt.expected.Start.Row ||
				result.End.Row != tt.expected.End.Row {
				t.Errorf(
					"UpdateRange() = %+v, want %+v",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestRangeReferenceUpdaterDeleteRows tests range update for row deletion.
func TestRangeReferenceUpdaterDeleteRows(
	t *testing.T,
) {
	tests := []struct {
		name     string
		rangeRef RangeRef
		index    uint32
		count    uint32
		expected RangeRef
		wantErr  bool
	}{
		{
			name: "range before delete",
			rangeRef: RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End:   CellRef{Col: 1, Row: 3},
			},
			index: 5,
			count: 2,
			expected: RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End:   CellRef{Col: 1, Row: 3},
			},
		},
		{
			name: "range after delete",
			rangeRef: RangeRef{
				Start: CellRef{Col: 1, Row: 10},
				End:   CellRef{Col: 1, Row: 15},
			},
			index: 5,
			count: 2,
			expected: RangeRef{
				Start: CellRef{Col: 1, Row: 8},
				End:   CellRef{Col: 1, Row: 13},
			},
		},
		{
			name: "range fully deleted",
			rangeRef: RangeRef{
				Start: CellRef{Col: 1, Row: 5},
				End:   CellRef{Col: 1, Row: 6},
			},
			index:   5,
			count:   5,
			wantErr: true,
		},
		{
			name: "range partially deleted start",
			rangeRef: RangeRef{
				Start: CellRef{Col: 1, Row: 5},
				End:   CellRef{Col: 1, Row: 10},
			},
			index: 5,
			count: 3,
			expected: RangeRef{
				Start: CellRef{Col: 1, Row: 5},
				End:   CellRef{Col: 1, Row: 7},
			},
		},
		{
			name: "range partially deleted end",
			rangeRef: RangeRef{
				Start: CellRef{Col: 1, Row: 3},
				End:   CellRef{Col: 1, Row: 6},
			},
			index: 5,
			count: 3,
			expected: RangeRef{
				Start: CellRef{Col: 1, Row: 3},
				End:   CellRef{Col: 1, Row: 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftDeleteRows,
				Index: tt.index,
				Count: tt.count,
			}
			updater := NewCellReferenceUpdater(
				op,
				"Sheet1",
			)
			result, err := updater.UpdateRange(
				tt.rangeRef,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"UpdateRange() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if tt.wantErr {
				return
			}
			if result.Start.Row != tt.expected.Start.Row ||
				result.End.Row != tt.expected.End.Row {
				t.Errorf(
					"UpdateRange() = %+v, want %+v",
					result,
					tt.expected,
				)
			}
		})
	}
}
