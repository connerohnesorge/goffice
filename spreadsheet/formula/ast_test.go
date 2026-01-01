package formula

import (
	"testing"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

func TestLiteral(t *testing.T) {
	tests := []struct {
		name     string
		literal  *Literal
		expected string
	}{
		{
			name:     "number",
			literal:  &Literal{Val: NewNumberValue(42)},
			expected: "42",
		},
		{
			name:     "string",
			literal:  &Literal{Val: NewStringValue("hello")},
			expected: `"hello"`,
		},
		{
			name:     "boolean true",
			literal:  &Literal{Val: NewBooleanValue(true)},
			expected: "TRUE",
		},
		{
			name:     "boolean false",
			literal:  &Literal{Val: NewBooleanValue(false)},
			expected: "FALSE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.literal.String(); got != tt.expected {
				t.Errorf("String() = %v, want %v", got, tt.expected)
			}

			// Dependencies should be empty
			if deps := tt.literal.Dependencies(); len(deps) != 0 {
				t.Errorf("Dependencies() = %v, want empty slice", deps)
			}
		})
	}
}

func TestCellReference(t *testing.T) {
	ref := &CellReference{
		Ref: spreadsheet.MustParseCellRef("A1"),
	}

	if got := ref.String(); got != "A1" {
		t.Errorf("String() = %v, want A1", got)
	}

	deps := ref.Dependencies()
	if len(deps) != 1 {
		t.Fatalf("Dependencies() length = %v, want 1", len(deps))
	}
	if deps[0].String() != "A1" {
		t.Errorf("Dependencies()[0] = %v, want A1", deps[0].String())
	}
}

func TestRangeReference(t *testing.T) {
	ref := &RangeReference{
		Range: spreadsheet.MustParseRangeRef("A1:B2"),
	}

	if got := ref.String(); got != "A1:B2" {
		t.Errorf("String() = %v, want A1:B2", got)
	}

	deps := ref.Dependencies()
	if len(deps) != 4 {
		t.Errorf("Dependencies() length = %v, want 4", len(deps))
	}
}

func TestBinaryOp(t *testing.T) {
	op := &BinaryOp{
		Left:     &Literal{Val: NewNumberValue(1)},
		Right:    &Literal{Val: NewNumberValue(2)},
		Operator: "+",
	}

	if got := op.String(); got != "(1 + 2)" {
		t.Errorf("String() = %v, want (1 + 2)", got)
	}

	// Dependencies should be empty (both operands are literals)
	if deps := op.Dependencies(); len(deps) != 0 {
		t.Errorf("Dependencies() = %v, want empty slice", deps)
	}
}

func TestUnaryOp(t *testing.T) {
	op := &UnaryOp{
		Operand:  &Literal{Val: NewNumberValue(5)},
		Operator: "-",
	}

	if got := op.String(); got != "(-5)" {
		t.Errorf("String() = %v, want (-5)", got)
	}

	// Dependencies should be empty
	if deps := op.Dependencies(); len(deps) != 0 {
		t.Errorf("Dependencies() = %v, want empty slice", deps)
	}
}

func TestFunctionCall(t *testing.T) {
	fn := &FunctionCall{
		Name: "SUM",
		Args: []Expression{
			&Literal{Val: NewNumberValue(1)},
			&Literal{Val: NewNumberValue(2)},
		},
	}

	if got := fn.String(); got != "SUM(1, 2)" {
		t.Errorf("String() = %v, want SUM(1, 2)", got)
	}

	// Dependencies should be empty (all args are literals)
	if deps := fn.Dependencies(); len(deps) != 0 {
		t.Errorf("Dependencies() = %v, want empty slice", deps)
	}
}

func TestDependencies(t *testing.T) {
	// Create an expression with cell references
	expr := &BinaryOp{
		Left:     &CellReference{Ref: spreadsheet.MustParseCellRef("A1")},
		Right:    &CellReference{Ref: spreadsheet.MustParseCellRef("B1")},
		Operator: "+",
	}

	deps := expr.Dependencies()
	if len(deps) != 2 {
		t.Fatalf("Dependencies() length = %v, want 2", len(deps))
	}

	// Check both references are present
	refs := make(map[string]bool)
	for _, dep := range deps {
		refs[dep.String()] = true
	}

	if !refs["A1"] || !refs["B1"] {
		t.Errorf("Dependencies() = %v, want A1 and B1", deps)
	}
}
