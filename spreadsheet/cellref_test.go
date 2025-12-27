package spreadsheet

import (
	"testing"
)

func TestColumnName(t *testing.T) {
	tests := []struct {
		col      int
		expected string
	}{
		{1, "A"},
		{2, "B"},
		{26, "Z"},
		{27, "AA"},
		{28, "AB"},
		{52, "AZ"},
		{53, "BA"},
		{702, "ZZ"},
		{703, "AAA"},
		{16384, "XFD"}, // Maximum Excel column
		{0, ""},        // Invalid
		{-1, ""},       // Invalid
	}

	for _, tt := range tests {
		result := ColumnName(tt.col)
		if result != tt.expected {
			t.Errorf(
				"ColumnName(%d) = %q, expected %q",
				tt.col,
				result,
				tt.expected,
			)
		}
	}
}

func TestColumnIndex(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{"A", 1},
		{"B", 2},
		{"Z", 26},
		{"AA", 27},
		{"AB", 28},
		{"AZ", 52},
		{"BA", 53},
		{"ZZ", 702},
		{"AAA", 703},
		{"XFD", 16384}, // Maximum Excel column
		{"a", 1},       // lowercase
		{"aa", 27},     // lowercase
		{"", 0},        // Invalid
		{"1", 0},       // Invalid
		{
			"A1",
			0,
		}, // Invalid (contains number)
	}

	for _, tt := range tests {
		result := ColumnIndex(tt.name)
		if result != tt.expected {
			t.Errorf(
				"ColumnIndex(%q) = %d, expected %d",
				tt.name,
				result,
				tt.expected,
			)
		}
	}
}

func TestColumnNameAndIndexRoundtrip(
	t *testing.T,
) {
	// Test that ColumnName and ColumnIndex are inverses
	for col := 1; col <= 16384; col++ {
		name := ColumnName(col)
		back := ColumnIndex(name)
		if back != col {
			t.Errorf(
				"Roundtrip failed: %d -> %q -> %d",
				col,
				name,
				back,
			)
		}
	}
}

func TestParseCellRef(t *testing.T) {
	tests := []struct {
		input   string
		wantCol int
		wantRow int
		absCol  bool
		absRow  bool
		wantErr bool
	}{
		{"A1", 1, 1, false, false, false},
		{"B2", 2, 2, false, false, false},
		{"Z26", 26, 26, false, false, false},
		{"AA100", 27, 100, false, false, false},
		{
			"XFD1048576",
			16384,
			1048576,
			false,
			false,
			false,
		}, // Max

		// Absolute references
		{"$A$1", 1, 1, true, true, false},
		{"$A1", 1, 1, true, false, false},
		{"A$1", 1, 1, false, true, false},
		{
			"$XFD$1048576",
			16384,
			1048576,
			true,
			true,
			false,
		},

		// Case insensitive
		{"a1", 1, 1, false, false, false},
		{"aa1", 27, 1, false, false, false},

		// Invalid references
		{"", 0, 0, false, false, true},
		{"1", 0, 0, false, false, true},
		{"A", 0, 0, false, false, true},
		{"$", 0, 0, false, false, true},
		{
			"A0",
			0,
			0,
			false,
			false,
			true,
		}, // Row 0 invalid
		{
			"XFE1",
			0,
			0,
			false,
			false,
			true,
		}, // Column too high
		{
			"A1048577",
			0,
			0,
			false,
			false,
			true,
		}, // Row too high
		{
			"A1B",
			0,
			0,
			false,
			false,
			true,
		}, // Invalid format
		{
			"$A$1$",
			0,
			0,
			false,
			false,
			true,
		}, // Invalid format
	}

	for _, tt := range tests {
		ref, err := ParseCellRef(tt.input)
		if tt.wantErr {
			if err == nil {
				t.Errorf(
					"ParseCellRef(%q) expected error, got nil",
					tt.input,
				)
			}

			continue
		}

		if err != nil {
			t.Errorf(
				"ParseCellRef(%q) unexpected error: %v",
				tt.input,
				err,
			)

			continue
		}

		if ref.Col != tt.wantCol {
			t.Errorf(
				"ParseCellRef(%q).Col = %d, want %d",
				tt.input,
				ref.Col,
				tt.wantCol,
			)
		}
		if ref.Row != tt.wantRow {
			t.Errorf(
				"ParseCellRef(%q).Row = %d, want %d",
				tt.input,
				ref.Row,
				tt.wantRow,
			)
		}
		if ref.AbsCol != tt.absCol {
			t.Errorf(
				"ParseCellRef(%q).AbsCol = %v, want %v",
				tt.input,
				ref.AbsCol,
				tt.absCol,
			)
		}
		if ref.AbsRow != tt.absRow {
			t.Errorf(
				"ParseCellRef(%q).AbsRow = %v, want %v",
				tt.input,
				ref.AbsRow,
				tt.absRow,
			)
		}
	}
}

func TestCellRefString(t *testing.T) {
	tests := []struct {
		ref      CellRef
		expected string
	}{
		{
			CellRef{
				Col:    1,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
			"A1",
		},
		{
			CellRef{
				Col:    2,
				Row:    10,
				AbsCol: false,
				AbsRow: false,
			},
			"B10",
		},
		{
			CellRef{
				Col:    27,
				Row:    100,
				AbsCol: false,
				AbsRow: false,
			},
			"AA100",
		},
		{
			CellRef{
				Col:    16384,
				Row:    1048576,
				AbsCol: false,
				AbsRow: false,
			},
			"XFD1048576",
		},

		// Absolute references
		{
			CellRef{
				Col:    1,
				Row:    1,
				AbsCol: true,
				AbsRow: true,
			},
			"$A$1",
		},
		{
			CellRef{
				Col:    1,
				Row:    1,
				AbsCol: true,
				AbsRow: false,
			},
			"$A1",
		},
		{
			CellRef{
				Col:    1,
				Row:    1,
				AbsCol: false,
				AbsRow: true,
			},
			"A$1",
		},
	}

	for _, tt := range tests {
		result := tt.ref.String()
		if result != tt.expected {
			t.Errorf(
				"CellRef{%d, %d, %v, %v}.String() = %q, want %q",
				tt.ref.Col,
				tt.ref.Row,
				tt.ref.AbsCol,
				tt.ref.AbsRow,
				result,
				tt.expected,
			)
		}
	}
}

func TestCellRefR1C1(t *testing.T) {
	tests := []struct {
		ref      CellRef
		expected string
	}{
		{CellRef{Col: 1, Row: 1}, "R1C1"},
		{CellRef{Col: 3, Row: 5}, "R5C3"},
		{
			CellRef{Col: 16384, Row: 1048576},
			"R1048576C16384",
		},
	}

	for _, tt := range tests {
		result := tt.ref.R1C1()
		if result != tt.expected {
			t.Errorf(
				"CellRef{%d, %d}.R1C1() = %q, want %q",
				tt.ref.Col,
				tt.ref.Row,
				result,
				tt.expected,
			)
		}
	}
}

func TestCellRefRoundtrip(t *testing.T) {
	testCases := []string{
		"A1",
		"Z26",
		"AA100",
		"XFD1048576",
		"$A$1",
		"$A1",
		"A$1",
		"$XFD$1048576",
	}

	for _, input := range testCases {
		ref, err := ParseCellRef(input)
		if err != nil {
			t.Errorf(
				"ParseCellRef(%q) failed: %v",
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

func TestCellRefIsAbsolute(t *testing.T) {
	tests := []struct {
		ref        CellRef
		isAbsolute bool
		isRelative bool
		isMixed    bool
	}{
		{
			CellRef{
				Col:    1,
				Row:    1,
				AbsCol: true,
				AbsRow: true,
			},
			true,
			false,
			false,
		},
		{
			CellRef{
				Col:    1,
				Row:    1,
				AbsCol: false,
				AbsRow: false,
			},
			false,
			true,
			false,
		},
		{
			CellRef{
				Col:    1,
				Row:    1,
				AbsCol: true,
				AbsRow: false,
			},
			false,
			false,
			true,
		},
		{
			CellRef{
				Col:    1,
				Row:    1,
				AbsCol: false,
				AbsRow: true,
			},
			false,
			false,
			true,
		},
	}

	for _, tt := range tests {
		if tt.ref.IsAbsolute() != tt.isAbsolute {
			t.Errorf(
				"CellRef{%v, %v}.IsAbsolute() = %v, want %v",
				tt.ref.AbsCol,
				tt.ref.AbsRow,
				tt.ref.IsAbsolute(),
				tt.isAbsolute,
			)
		}
		if tt.ref.IsRelative() != tt.isRelative {
			t.Errorf(
				"CellRef{%v, %v}.IsRelative() = %v, want %v",
				tt.ref.AbsCol,
				tt.ref.AbsRow,
				tt.ref.IsRelative(),
				tt.isRelative,
			)
		}
		if tt.ref.IsMixed() != tt.isMixed {
			t.Errorf(
				"CellRef{%v, %v}.IsMixed() = %v, want %v",
				tt.ref.AbsCol,
				tt.ref.AbsRow,
				tt.ref.IsMixed(),
				tt.isMixed,
			)
		}
	}
}

func TestCellRefIsValid(t *testing.T) {
	tests := []struct {
		ref   CellRef
		valid bool
	}{
		{CellRef{Col: 1, Row: 1}, true},
		{CellRef{Col: 16384, Row: 1048576}, true},
		{CellRef{Col: 0, Row: 1}, false},
		{CellRef{Col: 1, Row: 0}, false},
		{CellRef{Col: 16385, Row: 1}, false},
		{CellRef{Col: 1, Row: 1048577}, false},
	}

	for _, tt := range tests {
		if tt.ref.IsValid() != tt.valid {
			t.Errorf(
				"CellRef{%d, %d}.IsValid() = %v, want %v",
				tt.ref.Col,
				tt.ref.Row,
				tt.ref.IsValid(),
				tt.valid,
			)
		}
	}
}

func TestCellRefOffset(t *testing.T) {
	tests := []struct {
		ref       CellRef
		colOffset int
		rowOffset int
		expected  CellRef
		wantErr   bool
	}{
		{
			CellRef{Col: 5, Row: 5},
			2,
			3,
			CellRef{Col: 7, Row: 8},
			false,
		},
		{
			CellRef{Col: 5, Row: 5},
			-2,
			-3,
			CellRef{Col: 3, Row: 2},
			false,
		},
		{
			CellRef{Col: 5, Row: 5},
			0,
			0,
			CellRef{Col: 5, Row: 5},
			false,
		},

		// Out of range
		{
			CellRef{Col: 1, Row: 1},
			-1,
			0,
			CellRef{},
			true,
		},
		{
			CellRef{Col: 1, Row: 1},
			0,
			-1,
			CellRef{},
			true,
		},
		{
			CellRef{Col: 16384, Row: 1},
			1,
			0,
			CellRef{},
			true,
		},
		{
			CellRef{Col: 1, Row: 1048576},
			0,
			1,
			CellRef{},
			true,
		},
	}

	for _, tt := range tests {
		result, err := tt.ref.Offset(
			tt.colOffset,
			tt.rowOffset,
		)
		if tt.wantErr {
			if err == nil {
				t.Errorf(
					"CellRef{%d, %d}.Offset(%d, %d) expected error",
					tt.ref.Col,
					tt.ref.Row,
					tt.colOffset,
					tt.rowOffset,
				)
			}

			continue
		}

		if err != nil {
			t.Errorf(
				"CellRef{%d, %d}.Offset(%d, %d) unexpected error: %v",
				tt.ref.Col,
				tt.ref.Row,
				tt.colOffset,
				tt.rowOffset,
				err,
			)

			continue
		}

		if result.Col != tt.expected.Col ||
			result.Row != tt.expected.Row {
			t.Errorf(
				"CellRef{%d, %d}.Offset(%d, %d) = {%d, %d}, want {%d, %d}",
				tt.ref.Col,
				tt.ref.Row,
				tt.colOffset,
				tt.rowOffset,
				result.Col,
				result.Row,
				tt.expected.Col,
				tt.expected.Row,
			)
		}
	}
}

func TestNewCellRef(t *testing.T) {
	ref := NewCellRef(3, 5)
	if ref.Col != 3 || ref.Row != 5 {
		t.Errorf(
			"NewCellRef(3, 5) = {%d, %d}, want {3, 5}",
			ref.Col,
			ref.Row,
		)
	}
	if ref.AbsCol || ref.AbsRow {
		t.Error(
			"NewCellRef should create relative reference",
		)
	}
}

func TestNewAbsCellRef(t *testing.T) {
	ref := NewAbsCellRef(3, 5)
	if ref.Col != 3 || ref.Row != 5 {
		t.Errorf(
			"NewAbsCellRef(3, 5) = {%d, %d}, want {3, 5}",
			ref.Col,
			ref.Row,
		)
	}
	if !ref.AbsCol || !ref.AbsRow {
		t.Error(
			"NewAbsCellRef should create absolute reference",
		)
	}
}

func TestMustParseCellRef(t *testing.T) {
	// Valid reference should not panic
	ref := MustParseCellRef("A1")
	if ref.Col != 1 || ref.Row != 1 {
		t.Errorf(
			"MustParseCellRef(\"A1\") = %v, expected A1",
			ref,
		)
	}

	// Invalid reference should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error(
				"MustParseCellRef with invalid reference should panic",
			)
		}
	}()
	_ = MustParseCellRef("invalid")
}

func TestCellRefWithAbsolute(t *testing.T) {
	ref := CellRef{Col: 3, Row: 5}

	abs := ref.WithAbsolute(true, true)
	if !abs.AbsCol || !abs.AbsRow {
		t.Error(
			"WithAbsolute(true, true) should set both absolute",
		)
	}
	if abs.Col != 3 || abs.Row != 5 {
		t.Error(
			"WithAbsolute should preserve Col and Row",
		)
	}

	mixed := ref.WithAbsolute(true, false)
	if !mixed.AbsCol || mixed.AbsRow {
		t.Error(
			"WithAbsolute(true, false) should set only AbsCol",
		)
	}
}

func BenchmarkColumnName(b *testing.B) {
	for range b.N {
		ColumnName(16384)
	}
}

func BenchmarkColumnIndex(b *testing.B) {
	for range b.N {
		ColumnIndex("XFD")
	}
}

func BenchmarkParseCellRef(b *testing.B) {
	for range b.N {
		_, _ = ParseCellRef("$XFD$1048576")
	}
}

func BenchmarkCellRefString(b *testing.B) {
	ref := CellRef{
		Col:    16384,
		Row:    1048576,
		AbsCol: true,
		AbsRow: true,
	}
	for range b.N {
		_ = ref.String()
	}
}
