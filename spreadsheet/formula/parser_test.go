package formula

import (
	"testing"
)

func TestParserLiterals(t *testing.T) {
	tests := []struct {
		name     string
		formula  string
		expected string
	}{
		{
			name:     "number",
			formula:  "=42",
			expected: "42",
		},
		{
			name:     "decimal",
			formula:  "=3.14",
			expected: "3.14",
		},
		{
			name:     "string",
			formula:  `="hello"`,
			expected: `"hello"`,
		},
		{
			name:     "boolean true",
			formula:  "=TRUE",
			expected: "TRUE",
		},
		{
			name:     "boolean false",
			formula:  "=FALSE",
			expected: "FALSE",
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := parser.Parse(tt.formula)
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if got := expr.String(); got != tt.expected {
				t.Errorf("String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParserReferences(t *testing.T) {
	tests := []struct {
		name     string
		formula  string
		expected string
	}{
		{
			name:     "cell reference",
			formula:  "=A1",
			expected: "A1",
		},
		{
			name:     "absolute cell",
			formula:  "=$A$1",
			expected: "$A$1",
		},
		{
			name:     "range reference",
			formula:  "=A1:B2",
			expected: "A1:B2",
		},
		{
			name:     "sheet reference",
			formula:  "=Sheet1!A1",
			expected: "Sheet1!A1",
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := parser.Parse(tt.formula)
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if got := expr.String(); got != tt.expected {
				t.Errorf("String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParserBinaryOperators(t *testing.T) {
	tests := []struct {
		name     string
		formula  string
		expected string
	}{
		{
			name:     "addition",
			formula:  "=1+2",
			expected: "(1 + 2)",
		},
		{
			name:     "subtraction",
			formula:  "=5-3",
			expected: "(5 - 3)",
		},
		{
			name:     "multiplication",
			formula:  "=2*3",
			expected: "(2 * 3)",
		},
		{
			name:     "division",
			formula:  "=10/2",
			expected: "(10 / 2)",
		},
		{
			name:     "exponentiation",
			formula:  "=2^3",
			expected: "(2 ^ 3)",
		},
		{
			name:     "concatenation",
			formula:  `="hello"&"world"`,
			expected: `("hello" & "world")`,
		},
		{
			name:     "comparison equal",
			formula:  "=1=1",
			expected: "(1 = 1)",
		},
		{
			name:     "comparison not equal",
			formula:  "=1<>2",
			expected: "(1 <> 2)",
		},
		{
			name:     "comparison less than",
			formula:  "=1<2",
			expected: "(1 < 2)",
		},
		{
			name:     "comparison greater than",
			formula:  "=2>1",
			expected: "(2 > 1)",
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := parser.Parse(tt.formula)
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if got := expr.String(); got != tt.expected {
				t.Errorf("String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParserUnaryOperators(t *testing.T) {
	tests := []struct {
		name     string
		formula  string
		expected string
	}{
		{
			name:     "negation",
			formula:  "=-5",
			expected: "(-5)",
		},
		{
			name:     "positive",
			formula:  "=+5",
			expected: "(+5)",
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := parser.Parse(tt.formula)
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if got := expr.String(); got != tt.expected {
				t.Errorf("String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParserFunctions(t *testing.T) {
	tests := []struct {
		name     string
		formula  string
		expected string
	}{
		{
			name:     "no args",
			formula:  "=NOW()",
			expected: "NOW()",
		},
		{
			name:     "one arg",
			formula:  "=SUM(A1)",
			expected: "SUM(A1)",
		},
		{
			name:     "multiple args",
			formula:  "=SUM(A1,B1,C1)",
			expected: "SUM(A1, B1, C1)",
		},
		{
			name:     "nested functions",
			formula:  "=IF(SUM(A1:A10)>0,TRUE,FALSE)",
			expected: "IF((SUM(A1:A10) > 0), TRUE, FALSE)",
		},
		{
			name:     "function with range",
			formula:  "=SUM(A1:A10)",
			expected: "SUM(A1:A10)",
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := parser.Parse(tt.formula)
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if got := expr.String(); got != tt.expected {
				t.Errorf("String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParserPrecedence(t *testing.T) {
	tests := []struct {
		name     string
		formula  string
		expected string
	}{
		{
			name:     "addition before comparison",
			formula:  "=1+2=3",
			expected: "((1 + 2) = 3)",
		},
		{
			name:     "multiplication before addition",
			formula:  "=1+2*3",
			expected: "(1 + (2 * 3))",
		},
		{
			name:     "exponentiation before multiplication",
			formula:  "=2*3^4",
			expected: "(2 * (3 ^ 4))",
		},
		{
			name:     "parentheses override",
			formula:  "=(1+2)*3",
			expected: "((1 + 2) * 3)",
		},
		{
			name:     "right associative exponentiation",
			formula:  "=2^3^4",
			expected: "(2 ^ (3 ^ 4))",
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := parser.Parse(tt.formula)
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if got := expr.String(); got != tt.expected {
				t.Errorf("String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParserComplexFormulas(t *testing.T) {
	tests := []struct {
		name    string
		formula string
	}{
		{
			name:    "nested functions with operators",
			formula: "=IF(SUM(A1:A10)+B1>100,MAX(C1:C10),MIN(C1:C10))",
		},
		{
			name:    "mixed operators and functions",
			formula: "=SUM(A1:A10)*2+AVERAGE(B1:B10)/3",
		},
		{
			name:    "string concatenation with functions",
			formula: `=CONCATENATE("Total: ",SUM(A1:A10))`,
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parser.Parse(tt.formula)
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
		})
	}
}
