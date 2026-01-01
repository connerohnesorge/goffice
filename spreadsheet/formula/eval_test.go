//nolint:all // formula package is scaffolding for future formula evaluation - WIP
package formula

import (
	"math"
	"testing"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

// mockResolver implements CellResolver for testing.
type mockResolver struct {
	cells  map[string]Value
	ranges map[string][][]Value
}

func newMockResolver() *mockResolver {
	return &mockResolver{
		cells:  make(map[string]Value),
		ranges: make(map[string][][]Value),
	}
}

func (m *mockResolver) SetCell(ref string, val Value) {
	m.cells[ref] = val
}

func (m *mockResolver) SetRange(ref string, val [][]Value) {
	m.ranges[ref] = val
}

func (m *mockResolver) ResolveCell(ref spreadsheet.CellRef) (Value, error) {
	val, ok := m.cells[ref.String()]
	if !ok {
		return NewNumberValue(0), nil
	}

	return val, nil
}

func (m *mockResolver) ResolveRange(rng spreadsheet.RangeRef) ([][]Value, error) {
	val, ok := m.ranges[rng.String()]
	if !ok {
		// Return empty range
		return [][]Value{}, nil
	}

	return val, nil
}

func TestValueConversion(t *testing.T) {
	t.Run("AsNumber", func(t *testing.T) {
		tests := []struct {
			name    string
			value   Value
			want    float64
			wantErr bool
		}{
			{
				name:  "number",
				value: NewNumberValue(42),
				want:  42,
			},
			{
				name:  "string number",
				value: NewStringValue("3.14"),
				want:  3.14,
			},
			{
				name:  "boolean true",
				value: NewBooleanValue(true),
				want:  1,
			},
			{
				name:  "boolean false",
				value: NewBooleanValue(false),
				want:  0,
			},
			{
				name:    "string non-number",
				value:   NewStringValue("hello"),
				wantErr: true,
			},
			{
				name:    "error",
				value:   NewErrorValue(ErrValue),
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := tt.value.AsNumber()
				if (err != nil) != tt.wantErr {
					t.Errorf("AsNumber() error = %v, wantErr %v", err, tt.wantErr)

					return
				}
				if !tt.wantErr && got != tt.want {
					t.Errorf("AsNumber() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("AsString", func(t *testing.T) {
		tests := []struct {
			name  string
			value Value
			want  string
		}{
			{
				name:  "string",
				value: NewStringValue("hello"),
				want:  "hello",
			},
			{
				name:  "number",
				value: NewNumberValue(42),
				want:  "42",
			},
			{
				name:  "boolean true",
				value: NewBooleanValue(true),
				want:  "TRUE",
			},
			{
				name:  "boolean false",
				value: NewBooleanValue(false),
				want:  "FALSE",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := tt.value.AsString(); got != tt.want {
					t.Errorf("AsString() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("AsBoolean", func(t *testing.T) {
		tests := []struct {
			name    string
			value   Value
			want    bool
			wantErr bool
		}{
			{
				name:  "boolean true",
				value: NewBooleanValue(true),
				want:  true,
			},
			{
				name:  "boolean false",
				value: NewBooleanValue(false),
				want:  false,
			},
			{
				name:  "number non-zero",
				value: NewNumberValue(42),
				want:  true,
			},
			{
				name:  "number zero",
				value: NewNumberValue(0),
				want:  false,
			},
			{
				name:  "string TRUE",
				value: NewStringValue("TRUE"),
				want:  true,
			},
			{
				name:  "string FALSE",
				value: NewStringValue("FALSE"),
				want:  false,
			},
			{
				name:    "string invalid",
				value:   NewStringValue("hello"),
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := tt.value.AsBoolean()
				if (err != nil) != tt.wantErr {
					t.Errorf("AsBoolean() error = %v, wantErr %v", err, tt.wantErr)

					return
				}
				if !tt.wantErr && got != tt.want {
					t.Errorf("AsBoolean() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func TestBinaryOperations(t *testing.T) {
	tests := []struct {
		name     string
		left     Value
		right    Value
		operator string
		want     Value
	}{
		// Arithmetic
		{
			name:     "addition",
			left:     NewNumberValue(1),
			right:    NewNumberValue(2),
			operator: "+",
			want:     NewNumberValue(3),
		},
		{
			name:     "subtraction",
			left:     NewNumberValue(5),
			right:    NewNumberValue(3),
			operator: "-",
			want:     NewNumberValue(2),
		},
		{
			name:     "multiplication",
			left:     NewNumberValue(2),
			right:    NewNumberValue(3),
			operator: "*",
			want:     NewNumberValue(6),
		},
		{
			name:     "division",
			left:     NewNumberValue(10),
			right:    NewNumberValue(2),
			operator: "/",
			want:     NewNumberValue(5),
		},
		{
			name:     "division by zero",
			left:     NewNumberValue(10),
			right:    NewNumberValue(0),
			operator: "/",
			want:     NewErrorValue(ErrDiv0),
		},
		{
			name:     "exponentiation",
			left:     NewNumberValue(2),
			right:    NewNumberValue(3),
			operator: "^",
			want:     NewNumberValue(8),
		},
		// String concatenation
		{
			name:     "concatenation",
			left:     NewStringValue("hello"),
			right:    NewStringValue("world"),
			operator: "&",
			want:     NewStringValue("helloworld"),
		},
		// Comparison
		{
			name:     "equal true",
			left:     NewNumberValue(1),
			right:    NewNumberValue(1),
			operator: "=",
			want:     NewBooleanValue(true),
		},
		{
			name:     "equal false",
			left:     NewNumberValue(1),
			right:    NewNumberValue(2),
			operator: "=",
			want:     NewBooleanValue(false),
		},
		{
			name:     "not equal",
			left:     NewNumberValue(1),
			right:    NewNumberValue(2),
			operator: "<>",
			want:     NewBooleanValue(true),
		},
		{
			name:     "less than",
			left:     NewNumberValue(1),
			right:    NewNumberValue(2),
			operator: "<",
			want:     NewBooleanValue(true),
		},
		{
			name:     "greater than",
			left:     NewNumberValue(2),
			right:    NewNumberValue(1),
			operator: ">",
			want:     NewBooleanValue(true),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := evaluateBinaryOp(tt.left, tt.right, tt.operator)
			if err != nil {
				t.Fatalf("evaluateBinaryOp() error = %v", err)
			}

			if got.Type != tt.want.Type {
				t.Errorf("Type = %v, want %v", got.Type, tt.want.Type)
			}

			// Compare values
			switch tt.want.Type {
			case ValueTypeNumber:
				gotVal, _ := got.AsNumber()
				wantVal, _ := tt.want.AsNumber()
				if math.Abs(gotVal-wantVal) > 1e-9 {
					t.Errorf("Value = %v, want %v", gotVal, wantVal)
				}
			case ValueTypeString:
				if got.AsString() != tt.want.AsString() {
					t.Errorf("Value = %v, want %v", got.AsString(), tt.want.AsString())
				}
			case ValueTypeBoolean:
				gotVal, _ := got.AsBoolean()
				wantVal, _ := tt.want.AsBoolean()
				if gotVal != wantVal {
					t.Errorf("Value = %v, want %v", gotVal, wantVal)
				}
			case ValueError:
				if got.Value != tt.want.Value {
					t.Errorf("Error = %v, want %v", got.Value, tt.want.Value)
				}
			}
		})
	}
}

func TestUnaryOperations(t *testing.T) {
	tests := []struct {
		name     string
		operand  Value
		operator string
		want     Value
	}{
		{
			name:     "negation",
			operand:  NewNumberValue(5),
			operator: "-",
			want:     NewNumberValue(-5),
		},
		{
			name:     "positive",
			operand:  NewNumberValue(5),
			operator: "+",
			want:     NewNumberValue(5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := evaluateUnaryOp(tt.operand, tt.operator)
			if err != nil {
				t.Fatalf("evaluateUnaryOp() error = %v", err)
			}

			gotVal, _ := got.AsNumber()
			wantVal, _ := tt.want.AsNumber()
			if math.Abs(gotVal-wantVal) > 1e-9 {
				t.Errorf("Value = %v, want %v", gotVal, wantVal)
			}
		})
	}
}

func TestEvaluateWithResolver(t *testing.T) {
	parser := NewParser()
	resolver := newMockResolver()
	ctx := NewEvalContext(nil, "Sheet1")
	ctx.Resolver = resolver

	// Set up test data
	resolver.SetCell("A1", NewNumberValue(10))
	resolver.SetCell("B1", NewNumberValue(20))
	resolver.SetRange("A1:A3", [][]Value{
		{NewNumberValue(1)},
		{NewNumberValue(2)},
		{NewNumberValue(3)},
	})

	tests := []struct {
		name    string
		formula string
		want    float64
	}{
		{
			name:    "cell reference",
			formula: "=A1",
			want:    10,
		},
		{
			name:    "cell addition",
			formula: "=A1+B1",
			want:    30,
		},
		{
			name:    "sum range",
			formula: "=SUM(A1:A3)",
			want:    6,
		},
		{
			name:    "complex expression",
			formula: "=A1*2+B1",
			want:    40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := parser.Parse(tt.formula)
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			result, err := expr.Evaluate(ctx)
			if err != nil {
				t.Fatalf("Evaluate() error = %v", err)
			}

			got, err := result.AsNumber()
			if err != nil {
				t.Fatalf("AsNumber() error = %v", err)
			}

			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Result = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorPropagation(t *testing.T) {
	parser := NewParser()
	ctx := NewEvalContext(nil, "Sheet1")

	tests := []struct {
		name      string
		formula   string
		wantError string
	}{
		{
			name:      "division by zero",
			formula:   "=10/0",
			wantError: ErrDiv0,
		},
		{
			name:      "invalid type conversion",
			formula:   `="hello"+5`,
			wantError: ErrValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := parser.Parse(tt.formula)
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			result, err := expr.Evaluate(ctx)
			if err != nil {
				t.Fatalf("Evaluate() error = %v", err)
			}

			if result.Type != ValueError {
				t.Errorf("Expected error value, got %v", result.Type)
			}

			if result.Value != tt.wantError {
				t.Errorf("Error = %v, want %v", result.Value, tt.wantError)
			}
		})
	}
}
