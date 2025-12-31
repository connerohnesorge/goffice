package spreadsheet

import (
	"testing"
)

func TestParseRangeRef(t *testing.T) {
	tests := []struct {
		input    string
		startCol int
		startRow int
		endCol   int
		endRow   int
		wantErr  bool
	}{
		// Single cell (range equals single cell)
		{"A1", 1, 1, 1, 1, false},
		{
			"XFD1048576",
			16384,
			1048576,
			16384,
			1048576,
			false,
		},

		// Ranges
		{"A1:C10", 1, 1, 3, 10, false},
		{"A1:A10", 1, 1, 1, 10, false},
		{"A1:Z1", 1, 1, 26, 1, false},
		{"AA1:AZ100", 27, 1, 52, 100, false},

		// Absolute references
		{"$A$1:$C$10", 1, 1, 3, 10, false},
		{"$A1:C$10", 1, 1, 3, 10, false},

		// Reversed ranges (should normalize)
		{"C10:A1", 1, 1, 3, 10, false},
		{"Z1:A10", 1, 1, 26, 10, false},

		// Invalid
		{"", 0, 0, 0, 0, true},
		{":", 0, 0, 0, 0, true},
		{"A1:", 0, 0, 0, 0, true},
		{":C10", 0, 0, 0, 0, true},
		{"A1:B", 0, 0, 0, 0, true},
		{"A:C10", 0, 0, 0, 0, true},
	}

	for _, tt := range tests {
		ref, err := ParseRangeRef(tt.input)
		if tt.wantErr {
			if err == nil {
				t.Errorf(
					"ParseRangeRef(%q) expected error, got nil",
					tt.input,
				)
			}

			continue
		}

		if err != nil {
			t.Errorf(
				"ParseRangeRef(%q) unexpected error: %v",
				tt.input,
				err,
			)

			continue
		}

		if ref.Start.Col != tt.startCol {
			t.Errorf(
				"ParseRangeRef(%q).Start.Col = %d, want %d",
				tt.input,
				ref.Start.Col,
				tt.startCol,
			)
		}
		if ref.Start.Row != tt.startRow {
			t.Errorf(
				"ParseRangeRef(%q).Start.Row = %d, want %d",
				tt.input,
				ref.Start.Row,
				tt.startRow,
			)
		}
		if ref.End.Col != tt.endCol {
			t.Errorf(
				"ParseRangeRef(%q).End.Col = %d, want %d",
				tt.input,
				ref.End.Col,
				tt.endCol,
			)
		}
		if ref.End.Row != tt.endRow {
			t.Errorf(
				"ParseRangeRef(%q).End.Row = %d, want %d",
				tt.input,
				ref.End.Row,
				tt.endRow,
			)
		}
	}
}

func TestRangeRefString(t *testing.T) {
	tests := []struct {
		ref      RangeRef
		expected string
	}{
		// Single cell
		{
			RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End:   CellRef{Col: 1, Row: 1},
			},
			"A1",
		},

		// Ranges
		{
			RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End:   CellRef{Col: 3, Row: 10},
			},
			"A1:C10",
		},

		// Absolute
		{
			RangeRef{
				Start: CellRef{
					Col:    1,
					Row:    1,
					AbsCol: true,
					AbsRow: true,
				},
				End: CellRef{
					Col:    3,
					Row:    10,
					AbsCol: true,
					AbsRow: true,
				},
			},
			"$A$1:$C$10",
		},
	}

	for _, tt := range tests {
		result := tt.ref.String()
		if result != tt.expected {
			t.Errorf(
				"RangeRef{%v, %v}.String() = %q, want %q",
				tt.ref.Start,
				tt.ref.End,
				result,
				tt.expected,
			)
		}
	}
}

func TestRangeRefRoundtrip(t *testing.T) {
	testCases := []string{
		"A1",
		"A1:C10",
		"$A$1:$C$10",
		"AA1:ZZ100",
		"XFD1:XFD1048576",
	}

	for _, input := range testCases {
		ref, err := ParseRangeRef(input)
		if err != nil {
			t.Errorf(
				"ParseRangeRef(%q) failed: %v",
				input,
				err,
			)

			continue
		}

		result := ref.String()
		if result != input {
			t.Errorf(
				"Roundtrip failed: %q -> %v -> %q",
				input,
				ref,
				result,
			)
		}
	}
}

func TestRangeRefIsSingleCell(t *testing.T) {
	tests := []struct {
		ref      RangeRef
		expected bool
	}{
		{
			RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End:   CellRef{Col: 1, Row: 1},
			},
			true,
		},
		{
			RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End:   CellRef{Col: 2, Row: 1},
			},
			false,
		},
		{
			RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End:   CellRef{Col: 1, Row: 2},
			},
			false,
		},
	}

	for _, tt := range tests {
		result := tt.ref.IsSingleCell()
		if result != tt.expected {
			t.Errorf(
				"RangeRef{%v, %v}.IsSingleCell() = %v, want %v",
				tt.ref.Start,
				tt.ref.End,
				result,
				tt.expected,
			)
		}
	}
}

func TestRangeRefContains(t *testing.T) {
	r := MustParseRangeRef("B2:D4")

	tests := []struct {
		cell     CellRef
		expected bool
	}{
		{
			CellRef{Col: 2, Row: 2},
			true,
		}, // Top-left
		{
			CellRef{Col: 4, Row: 4},
			true,
		}, // Bottom-right
		{
			CellRef{Col: 3, Row: 3},
			true,
		}, // Middle
		{
			CellRef{Col: 1, Row: 2},
			false,
		}, // Left of range
		{
			CellRef{Col: 5, Row: 2},
			false,
		}, // Right of range
		{
			CellRef{Col: 2, Row: 1},
			false,
		}, // Above range
		{
			CellRef{Col: 2, Row: 5},
			false,
		}, // Below range
	}

	for _, tt := range tests {
		result := r.Contains(tt.cell)
		if result != tt.expected {
			t.Errorf(
				"B2:D4.Contains(%s) = %v, want %v",
				tt.cell.String(),
				result,
				tt.expected,
			)
		}
	}
}

func TestRangeRefIntersects(t *testing.T) {
	r1 := MustParseRangeRef("B2:D4")

	tests := []struct {
		other    string
		expected bool
	}{
		{"B2:D4", true},  // Same range
		{"C3:C3", true},  // Single cell inside
		{"A1:B2", true},  // Overlaps corner
		{"D4:E5", true},  // Overlaps other corner
		{"A1:A1", false}, // Before
		{"E5:F6", false}, // After
		{"A1:A4", false}, // Left of range
		{"E1:E4", false}, // Right of range
		{"A1:D1", false}, // Above range
		{"A5:D5", false}, // Below range
	}

	for _, tt := range tests {
		other := MustParseRangeRef(tt.other)
		result := r1.Intersects(other)
		if result != tt.expected {
			t.Errorf(
				"B2:D4.Intersects(%s) = %v, want %v",
				tt.other,
				result,
				tt.expected,
			)
		}
	}
}

func TestRangeRefIntersection(t *testing.T) {
	tests := []struct {
		r1       string
		r2       string
		expected string
		wantErr  bool
	}{
		{"A1:C3", "B2:D4", "B2:C3", false},
		{"A1:D4", "B2:C3", "B2:C3", false},
		{"A1:A1", "A1:A1", "A1", false},
		{
			"A1:B2",
			"C3:D4",
			"",
			true,
		}, // No intersection
	}

	for _, tt := range tests {
		r1 := MustParseRangeRef(tt.r1)
		r2 := MustParseRangeRef(tt.r2)

		result, err := r1.Intersection(r2)
		if tt.wantErr {
			if err == nil {
				t.Errorf(
					"%s.Intersection(%s) expected error",
					tt.r1,
					tt.r2,
				)
			}

			continue
		}

		if err != nil {
			t.Errorf(
				"%s.Intersection(%s) unexpected error: %v",
				tt.r1,
				tt.r2,
				err,
			)

			continue
		}

		if result.String() != tt.expected {
			t.Errorf(
				"%s.Intersection(%s) = %s, want %s",
				tt.r1,
				tt.r2,
				result.String(),
				tt.expected,
			)
		}
	}
}

func TestRangeRefWidthHeightSize(t *testing.T) {
	tests := []struct {
		ref    string
		width  int
		height int
		size   int
	}{
		{"A1", 1, 1, 1},
		{"A1:A10", 1, 10, 10},
		{"A1:J1", 10, 1, 10},
		{"A1:J10", 10, 10, 100},
		{"B2:D4", 3, 3, 9},
	}

	for _, tt := range tests {
		ref := MustParseRangeRef(tt.ref)

		if ref.Width() != tt.width {
			t.Errorf(
				"%s.Width() = %d, want %d",
				tt.ref,
				ref.Width(),
				tt.width,
			)
		}
		if ref.Height() != tt.height {
			t.Errorf(
				"%s.Height() = %d, want %d",
				tt.ref,
				ref.Height(),
				tt.height,
			)
		}
		if ref.Size() != tt.size {
			t.Errorf(
				"%s.Size() = %d, want %d",
				tt.ref,
				ref.Size(),
				tt.size,
			)
		}
	}
}

func TestRangeRefCells(t *testing.T) {
	ref := MustParseRangeRef("B2:C3")

	expected := []string{"B2", "C2", "B3", "C3"}
	result := make([]string, 0, ref.Size())

	for cell := range ref.Cells() {
		result = append(result, cell.String())
	}

	if len(result) != len(expected) {
		t.Errorf(
			"Cells() returned %d cells, want %d",
			len(result),
			len(expected),
		)
	}

	for i, cell := range result {
		if cell != expected[i] {
			t.Errorf(
				"Cells()[%d] = %s, want %s",
				i,
				cell,
				expected[i],
			)
		}
	}
}

func TestRangeRefRows(t *testing.T) {
	ref := MustParseRangeRef("B2:C3")

	expectedRows := []int{2, 3}
	expectedCellsPerRow := [][]string{
		{"B2", "C2"},
		{"B3", "C3"},
	}

	rowIdx := 0
	for row, cells := range ref.Rows() {
		if row != expectedRows[rowIdx] {
			t.Errorf(
				"Row %d: got row number %d, want %d",
				rowIdx,
				row,
				expectedRows[rowIdx],
			)
		}

		cellIdx := 0
		for cell := range cells {
			expected := expectedCellsPerRow[rowIdx][cellIdx]
			if cell.String() != expected {
				t.Errorf(
					"Row %d, cell %d: got %s, want %s",
					rowIdx,
					cellIdx,
					cell.String(),
					expected,
				)
			}
			cellIdx++
		}
		rowIdx++
	}
}

func TestRangeRefColumns(t *testing.T) {
	ref := MustParseRangeRef("B2:C3")

	expectedCols := []int{2, 3}
	expectedCellsPerCol := [][]string{
		{"B2", "B3"},
		{"C2", "C3"},
	}

	colIdx := 0
	for col, cells := range ref.Columns() {
		if col != expectedCols[colIdx] {
			t.Errorf(
				"Column %d: got column number %d, want %d",
				colIdx,
				col,
				expectedCols[colIdx],
			)
		}

		cellIdx := 0
		for cell := range cells {
			expected := expectedCellsPerCol[colIdx][cellIdx]
			if cell.String() != expected {
				t.Errorf(
					"Column %d, cell %d: got %s, want %s",
					colIdx,
					cellIdx,
					cell.String(),
					expected,
				)
			}
			cellIdx++
		}
		colIdx++
	}
}

func TestRangeRefTopBottomLeftRight(
	t *testing.T,
) {
	ref := MustParseRangeRef("B2:D4")

	tests := []struct {
		name     string
		fn       func(RangeRef) RangeRef
		expected string
	}{
		{"TopRow", RangeRef.TopRow, "B2:D2"},
		{
			"BottomRow",
			RangeRef.BottomRow,
			"B4:D4",
		},
		{
			"LeftColumn",
			RangeRef.LeftColumn,
			"B2:B4",
		},
		{
			"RightColumn",
			RangeRef.RightColumn,
			"D2:D4",
		},
	}

	for _, tt := range tests {
		result := tt.fn(ref)
		if result.String() != tt.expected {
			t.Errorf(
				"%s(B2:D4) = %s, want %s",
				tt.name,
				result.String(),
				tt.expected,
			)
		}
	}
}

func TestRangeRefOffset(t *testing.T) {
	tests := []struct {
		ref       string
		colOffset int
		rowOffset int
		expected  string
		wantErr   bool
	}{
		{"A1:C3", 2, 3, "C4:E6", false},
		{"B2:D4", -1, -1, "A1:C3", false},
		{"A1:C3", 0, 0, "A1:C3", false},

		// Out of range
		{"A1:C3", -1, 0, "", true},
		{"A1:C3", 0, -1, "", true},
		{"XFD1:XFD1", 1, 0, "", true},
	}

	for _, tt := range tests {
		ref := MustParseRangeRef(tt.ref)
		result, err := ref.Offset(
			tt.colOffset,
			tt.rowOffset,
		)

		if tt.wantErr {
			if err == nil {
				t.Errorf(
					"%s.Offset(%d, %d) expected error",
					tt.ref,
					tt.colOffset,
					tt.rowOffset,
				)
			}

			continue
		}

		if err != nil {
			t.Errorf(
				"%s.Offset(%d, %d) unexpected error: %v",
				tt.ref,
				tt.colOffset,
				tt.rowOffset,
				err,
			)

			continue
		}

		if result.String() != tt.expected {
			t.Errorf(
				"%s.Offset(%d, %d) = %s, want %s",
				tt.ref,
				tt.colOffset,
				tt.rowOffset,
				result.String(),
				tt.expected,
			)
		}
	}
}

func TestRangeRefExpand(t *testing.T) {
	tests := []struct {
		ref      string
		left     int
		right    int
		top      int
		bottom   int
		expected string
		wantErr  bool
	}{
		{"B2:C3", 1, 1, 1, 1, "A1:D4", false},
		{"B2:C3", 0, 0, 0, 0, "B2:C3", false},
		{
			"B2:C3",
			-1,
			-1,
			-1,
			-1,
			"",
			true,
		}, // Would result in empty range
		{
			"A1:B2",
			1,
			0,
			0,
			0,
			"",
			true,
		}, // Would go before column A
		{
			"A1:B2",
			0,
			0,
			1,
			0,
			"",
			true,
		}, // Would go before row 1
	}

	for _, tt := range tests {
		ref := MustParseRangeRef(tt.ref)
		result, err := ref.Expand(
			tt.left,
			tt.right,
			tt.top,
			tt.bottom,
		)

		if tt.wantErr {
			if err == nil {
				t.Errorf(
					"%s.Expand(%d, %d, %d, %d) expected error",
					tt.ref,
					tt.left,
					tt.right,
					tt.top,
					tt.bottom,
				)
			}

			continue
		}

		if err != nil {
			t.Errorf(
				"%s.Expand(%d, %d, %d, %d) unexpected error: %v",
				tt.ref,
				tt.left,
				tt.right,
				tt.top,
				tt.bottom,
				err,
			)

			continue
		}

		if result.String() != tt.expected {
			t.Errorf(
				"%s.Expand(%d, %d, %d, %d) = %s, want %s",
				tt.ref,
				tt.left,
				tt.right,
				tt.top,
				tt.bottom,
				result.String(),
				tt.expected,
			)
		}
	}
}

func TestRangeRefNormalize(t *testing.T) {
	tests := []struct {
		start    CellRef
		end      CellRef
		expected string
	}{
		// Already normalized
		{
			CellRef{Col: 1, Row: 1},
			CellRef{Col: 3, Row: 3},
			"A1:C3",
		},
		// Reversed
		{
			CellRef{Col: 3, Row: 3},
			CellRef{Col: 1, Row: 1},
			"A1:C3",
		},
		// Mixed
		{
			CellRef{Col: 3, Row: 1},
			CellRef{Col: 1, Row: 3},
			"A1:C3",
		},
		{
			CellRef{Col: 1, Row: 3},
			CellRef{Col: 3, Row: 1},
			"A1:C3",
		},
	}

	for _, tt := range tests {
		ref := RangeRef{
			Start: tt.start,
			End:   tt.end,
		}
		normalized := ref.Normalize()
		if normalized.String() != tt.expected {
			t.Errorf(
				"Normalize({%s, %s}) = %s, want %s",
				tt.start.String(),
				tt.end.String(),
				normalized.String(),
				tt.expected,
			)
		}
	}
}

func TestMustParseRangeRef(t *testing.T) {
	// Valid reference should not panic
	ref := MustParseRangeRef("A1:C3")
	if ref.String() != "A1:C3" {
		t.Errorf(
			"MustParseRangeRef(\"A1:C3\") = %s",
			ref.String(),
		)
	}

	// Invalid reference should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error(
				"MustParseRangeRef with invalid reference should panic",
			)
		}
	}()
	_ = MustParseRangeRef("invalid:range")
}

func TestNewRangeRef(t *testing.T) {
	start := CellRef{Col: 3, Row: 3}
	end := CellRef{Col: 1, Row: 1}

	ref := NewRangeRef(start, end)

	// Should be normalized
	if ref.Start.Col != 1 || ref.Start.Row != 1 {
		t.Errorf(
			"NewRangeRef start = {%d, %d}, want {1, 1}",
			ref.Start.Col,
			ref.Start.Row,
		)
	}
	if ref.End.Col != 3 || ref.End.Row != 3 {
		t.Errorf(
			"NewRangeRef end = {%d, %d}, want {3, 3}",
			ref.End.Col,
			ref.End.Row,
		)
	}
}

func TestNewSingleCellRange(t *testing.T) {
	cell := CellRef{Col: 5, Row: 10}
	ref := NewSingleCellRange(cell)

	if !ref.IsSingleCell() {
		t.Error(
			"NewSingleCellRange should create single cell range",
		)
	}
	if ref.Start.Col != 5 || ref.Start.Row != 10 {
		t.Errorf(
			"NewSingleCellRange = {%d, %d}",
			ref.Start.Col,
			ref.Start.Row,
		)
	}
}

func TestRangeRefIsValid(t *testing.T) {
	tests := []struct {
		ref   RangeRef
		valid bool
	}{
		{
			RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End: CellRef{
					Col: 16384,
					Row: 1048576,
				},
			},
			true,
		},
		{
			RangeRef{
				Start: CellRef{Col: 0, Row: 1},
				End:   CellRef{Col: 1, Row: 1},
			},
			false,
		},
		{
			RangeRef{
				Start: CellRef{Col: 1, Row: 1},
				End: CellRef{
					Col: 16385,
					Row: 1,
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		result := tt.ref.IsValid()
		if result != tt.valid {
			t.Errorf(
				"RangeRef{%v, %v}.IsValid() = %v, want %v",
				tt.ref.Start,
				tt.ref.End,
				result,
				tt.valid,
			)
		}
	}
}

func BenchmarkParseRangeRef(b *testing.B) {
	for range b.N {
		_, _ = ParseRangeRef("$A$1:$XFD$1048576")
	}
}

func BenchmarkRangeRefCells(b *testing.B) {
	ref := MustParseRangeRef("A1:Z100")
	b.ResetTimer()

	for range b.N {
		for cell := range ref.Cells() {
			_ = cell // consume iterator
		}
	}
}

// TestParseCrossSheetRange tests parsing of cross-sheet range references
func TestParseCrossSheetRange(t *testing.T) {
	tests := []struct {
		input     string
		wantSheet string
		startCol  int
		startRow  int
		endCol    int
		endRow    int
		wantErr   bool
	}{
		// Unquoted sheet names
		{
			"Sheet2!A1:C3",
			"Sheet2",
			1,
			1,
			3,
			3,
			false,
		},
		{
			"Data!A1:Z100",
			"Data",
			1,
			1,
			26,
			100,
			false,
		},
		{
			"Summary!B2:D4",
			"Summary",
			2,
			2,
			4,
			4,
			false,
		},

		// Quoted sheet names
		{
			"'My Sheet'!A1:C3",
			"My Sheet",
			1,
			1,
			3,
			3,
			false,
		},
		{
			"'Sheet (2024)'!A1:D10",
			"Sheet (2024)",
			1,
			1,
			4,
			10,
			false,
		},
		{
			"'Data:Summary'!B2:E5",
			"Data:Summary",
			2,
			2,
			5,
			5,
			false,
		},

		// Single cell with sheet
		{
			"Sheet2!A1",
			"Sheet2",
			1,
			1,
			1,
			1,
			false,
		},
		{
			"'My Sheet'!B2",
			"My Sheet",
			2,
			2,
			2,
			2,
			false,
		},

		// Invalid
		{"'Unclosed!A1:C3", "", 0, 0, 0, 0, true},
		{"Sheet!:C3", "", 0, 0, 0, 0, true},
	}

	for _, tt := range tests {
		ref, err := ParseRangeRef(tt.input)
		if tt.wantErr {
			if err == nil {
				t.Errorf(
					"ParseRangeRef(%q) expected error, got nil",
					tt.input,
				)
			}

			continue
		}

		if err != nil {
			t.Errorf(
				"ParseRangeRef(%q) unexpected error: %v",
				tt.input,
				err,
			)

			continue
		}

		if ref.Start.Sheet != tt.wantSheet {
			t.Errorf(
				"ParseRangeRef(%q).Start.Sheet = %q, want %q",
				tt.input,
				ref.Start.Sheet,
				tt.wantSheet,
			)
		}
		if ref.End.Sheet != tt.wantSheet {
			t.Errorf(
				"ParseRangeRef(%q).End.Sheet = %q, want %q",
				tt.input,
				ref.End.Sheet,
				tt.wantSheet,
			)
		}

		if ref.Start.Col != tt.startCol {
			t.Errorf(
				"ParseRangeRef(%q).Start.Col = %d, want %d",
				tt.input,
				ref.Start.Col,
				tt.startCol,
			)
		}
		if ref.Start.Row != tt.startRow {
			t.Errorf(
				"ParseRangeRef(%q).Start.Row = %d, want %d",
				tt.input,
				ref.Start.Row,
				tt.startRow,
			)
		}
		if ref.End.Col != tt.endCol {
			t.Errorf(
				"ParseRangeRef(%q).End.Col = %d, want %d",
				tt.input,
				ref.End.Col,
				tt.endCol,
			)
		}
		if ref.End.Row != tt.endRow {
			t.Errorf(
				"ParseRangeRef(%q).End.Row = %d, want %d",
				tt.input,
				ref.End.Row,
				tt.endRow,
			)
		}
	}
}

// TestCrossSheetRangeRoundtrip tests roundtrip of cross-sheet range references
func TestCrossSheetRangeRoundtrip(t *testing.T) {
	testCases := []string{
		"Sheet2!A1:C3",
		"Sheet2!$A$1:$C$3",
		"'My Sheet'!A1:D10",
		"'Sheet (2024)'!B2:E5",
		"'Data:Summary'!AA1:ZZ100",
	}

	for _, input := range testCases {
		ref, err := ParseRangeRef(input)
		if err != nil {
			t.Errorf(
				"ParseRangeRef(%q) failed: %v",
				input,
				err,
			)

			continue
		}

		result := ref.String()
		if result != input {
			t.Errorf(
				"Roundtrip failed: %q -> %q",
				input,
				result,
			)
		}
	}
}
