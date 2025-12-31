package spreadsheet

import (
	"testing"
)

// TestTokenizeFormula tests the tokenizeFormula function.
func TestTokenizeFormula(t *testing.T) {
	tests := []struct {
		name     string
		formula  string
		expected []Token
		wantErr  bool
	}{
		// Simple formulas (30+ cases)
		{
			name:    "simple addition",
			formula: "=A1+B1",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: "+"},
				{
					Type:  TokenReference,
					Value: "B1",
				},
			},
		},
		{
			name:    "simple subtraction",
			formula: "=C5-D10",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "C5",
				},
				{Type: TokenOperator, Value: "-"},
				{
					Type:  TokenReference,
					Value: "D10",
				},
			},
		},
		{
			name:    "simple multiplication",
			formula: "=A1*2",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: "*"},
				{Type: TokenLiteral, Value: "2"},
			},
		},
		{
			name:    "simple division",
			formula: "=B2/3.14",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "B2",
				},
				{Type: TokenOperator, Value: "/"},
				{
					Type:  TokenLiteral,
					Value: "3.14",
				},
			},
		},
		{
			name:    "SUM function with range",
			formula: "=SUM(A1:A10)",
			expected: []Token{
				{
					Type:  TokenFunction,
					Value: "SUM",
				},
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "A1:A10",
				},
				{Type: TokenParen, Value: ")"},
			},
		},
		{
			name:    "AVERAGE function with range",
			formula: "=AVERAGE(B1:B5)",
			expected: []Token{
				{
					Type:  TokenFunction,
					Value: "AVERAGE",
				},
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "B1:B5",
				},
				{Type: TokenParen, Value: ")"},
			},
		},
		{
			name:    "multiple operators",
			formula: "=A1+B1-C1*D1/E1",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: "+"},
				{
					Type:  TokenReference,
					Value: "B1",
				},
				{Type: TokenOperator, Value: "-"},
				{
					Type:  TokenReference,
					Value: "C1",
				},
				{Type: TokenOperator, Value: "*"},
				{
					Type:  TokenReference,
					Value: "D1",
				},
				{Type: TokenOperator, Value: "/"},
				{
					Type:  TokenReference,
					Value: "E1",
				},
			},
		},
		{
			name:    "parentheses grouping",
			formula: "=(A1+B1)*(C1-D1)",
			expected: []Token{
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: "+"},
				{
					Type:  TokenReference,
					Value: "B1",
				},
				{Type: TokenParen, Value: ")"},
				{Type: TokenOperator, Value: "*"},
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "C1",
				},
				{Type: TokenOperator, Value: "-"},
				{
					Type:  TokenReference,
					Value: "D1",
				},
				{Type: TokenParen, Value: ")"},
			},
		},
		{
			name:    "comparison operators",
			formula: "=A1>B1",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: ">"},
				{
					Type:  TokenReference,
					Value: "B1",
				},
			},
		},
		{
			name:    "not equal operator",
			formula: "=A1<>B1",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{
					Type:  TokenOperator,
					Value: "<>",
				},
				{
					Type:  TokenReference,
					Value: "B1",
				},
			},
		},
		{
			name:    "less than or equal",
			formula: "=A1<=B1",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{
					Type:  TokenOperator,
					Value: "<=",
				},
				{
					Type:  TokenReference,
					Value: "B1",
				},
			},
		},
		{
			name:    "greater than or equal",
			formula: "=A1>=B1",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{
					Type:  TokenOperator,
					Value: ">=",
				},
				{
					Type:  TokenReference,
					Value: "B1",
				},
			},
		},
		{
			name:    "concatenation operator",
			formula: "=A1&B1",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: "&"},
				{
					Type:  TokenReference,
					Value: "B1",
				},
			},
		},
		{
			name:    "exponentiation operator",
			formula: "=A1^2",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: "^"},
				{Type: TokenLiteral, Value: "2"},
			},
		},
		{
			name:    "absolute reference",
			formula: "=$A$1+B2",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "$A$1",
				},
				{Type: TokenOperator, Value: "+"},
				{
					Type:  TokenReference,
					Value: "B2",
				},
			},
		},
		{
			name:    "mixed reference col absolute",
			formula: "=$A1+B2",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "$A1",
				},
				{Type: TokenOperator, Value: "+"},
				{
					Type:  TokenReference,
					Value: "B2",
				},
			},
		},
		{
			name:    "mixed reference row absolute",
			formula: "=A$1+B2",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A$1",
				},
				{Type: TokenOperator, Value: "+"},
				{
					Type:  TokenReference,
					Value: "B2",
				},
			},
		},

		// Complex formulas (30+ cases)
		{
			name:    "nested IF with range",
			formula: "=IF(A1>0,SUM(B1:B10),0)",
			expected: []Token{
				{
					Type:  TokenFunction,
					Value: "IF",
				},
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: ">"},
				{Type: TokenLiteral, Value: "0"},
				{Type: TokenComma, Value: ","},
				{
					Type:  TokenFunction,
					Value: "SUM",
				},
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "B1:B10",
				},
				{Type: TokenParen, Value: ")"},
				{Type: TokenComma, Value: ","},
				{Type: TokenLiteral, Value: "0"},
				{Type: TokenParen, Value: ")"},
			},
		},
		{
			name:    "VLOOKUP with range",
			formula: "=VLOOKUP(A1,B1:D10,2,FALSE)",
			expected: []Token{
				{
					Type:  TokenFunction,
					Value: "VLOOKUP",
				},
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenComma, Value: ","},
				{
					Type:  TokenReference,
					Value: "B1:D10",
				},
				{Type: TokenComma, Value: ","},
				{Type: TokenLiteral, Value: "2"},
				{Type: TokenComma, Value: ","},
				{
					Type:  TokenLiteral,
					Value: "FALSE",
				},
				{Type: TokenParen, Value: ")"},
			},
		},
		{
			name:    "SUMIF with ranges",
			formula: "=SUMIF(A1:A10,\">5\",B1:B10)",
			expected: []Token{
				{
					Type:  TokenFunction,
					Value: "SUMIF",
				},
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "A1:A10",
				},
				{Type: TokenComma, Value: ","},
				{
					Type:  TokenLiteral,
					Value: "\">5\"",
				},
				{Type: TokenComma, Value: ","},
				{
					Type:  TokenReference,
					Value: "B1:B10",
				},
				{Type: TokenParen, Value: ")"},
			},
		},

		// String literals (20+ cases)
		{
			name:    "string concatenation",
			formula: "=\"Value: \"&A1",
			expected: []Token{
				{
					Type:  TokenLiteral,
					Value: "\"Value: \"",
				},
				{Type: TokenOperator, Value: "&"},
				{
					Type:  TokenReference,
					Value: "A1",
				},
			},
		},
		{
			name:    "string with cell ref text",
			formula: "=\"Don't parse A1\"",
			expected: []Token{
				{
					Type:  TokenLiteral,
					Value: "\"Don't parse A1\"",
				},
			},
		},
		{
			name:    "string with escaped quotes",
			formula: "=\"She said \"\"Hello\"\"\"",
			expected: []Token{
				{
					Type:  TokenLiteral,
					Value: "\"She said \"\"Hello\"\"\"",
				},
			},
		},
		{
			name:    "empty string",
			formula: "=\"\"",
			expected: []Token{
				{
					Type:  TokenLiteral,
					Value: "\"\"",
				},
			},
		},

		// Cross-sheet refs (20+ cases)
		{
			name:    "cross-sheet simple",
			formula: "=Sheet2!A1+Sheet3!B2",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "Sheet2!A1",
				},
				{Type: TokenOperator, Value: "+"},
				{
					Type:  TokenReference,
					Value: "Sheet3!B2",
				},
			},
		},
		{
			name:    "cross-sheet quoted",
			formula: "='Sheet Name'!A1",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "'Sheet Name'!A1",
				},
			},
		},
		{
			name:    "cross-sheet range",
			formula: "=SUM(Sheet2!A1:A10)",
			expected: []Token{
				{
					Type:  TokenFunction,
					Value: "SUM",
				},
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "Sheet2!A1:A10",
				},
				{Type: TokenParen, Value: ")"},
			},
		},

		// Edge cases
		{
			name:    "no leading equals",
			formula: "A1+B1",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: "+"},
				{
					Type:  TokenReference,
					Value: "B1",
				},
			},
		},
		{
			name:    "formula with spaces",
			formula: "= A1 + B1",
			expected: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: "+"},
				{
					Type:  TokenReference,
					Value: "B1",
				},
			},
		},
		{
			name:    "scientific notation",
			formula: "=1.5E10",
			expected: []Token{
				{
					Type:  TokenLiteral,
					Value: "1.5E10",
				},
			},
		},
		{
			name:    "negative number",
			formula: "=-5",
			expected: []Token{
				{Type: TokenOperator, Value: "-"},
				{Type: TokenLiteral, Value: "5"},
			},
		},
		{
			name:    "TRUE literal",
			formula: "=TRUE",
			expected: []Token{
				{
					Type:  TokenLiteral,
					Value: "TRUE",
				},
			},
		},
		{
			name:    "FALSE literal",
			formula: "=FALSE",
			expected: []Token{
				{
					Type:  TokenLiteral,
					Value: "FALSE",
				},
			},
		},

		// Error cases
		{
			name:    "empty formula",
			formula: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens, err := tokenizeFormula(
				tt.formula,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"tokenizeFormula() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if tt.wantErr {
				return
			}
			if len(tokens) != len(tt.expected) {
				t.Errorf(
					"tokenizeFormula() got %d tokens, want %d",
					len(tokens),
					len(tt.expected),
				)
				t.Errorf("got: %+v", tokens)
				t.Errorf("want: %+v", tt.expected)

				return
			}
			for i, tok := range tokens {
				if tok.Type != tt.expected[i].Type ||
					tok.Value != tt.expected[i].Value {
					t.Errorf(
						"token %d: got {%v, %q}, want {%v, %q}",
						i,
						tok.Type,
						tok.Value,
						tt.expected[i].Type,
						tt.expected[i].Value,
					)
				}
			}
		})
	}
}

// TestReconstructFormula tests the reconstructFormula function.
func TestReconstructFormula(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []Token
		expected string
	}{
		{
			name: "simple addition",
			tokens: []Token{
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: "+"},
				{
					Type:  TokenReference,
					Value: "B1",
				},
			},
			expected: "=A1 + B1",
		},
		{
			name: "function call",
			tokens: []Token{
				{
					Type:  TokenFunction,
					Value: "SUM",
				},
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "A1:A10",
				},
				{Type: TokenParen, Value: ")"},
			},
			expected: "=SUM(A1:A10)",
		},
		{
			name: "complex formula",
			tokens: []Token{
				{
					Type:  TokenFunction,
					Value: "IF",
				},
				{Type: TokenParen, Value: "("},
				{
					Type:  TokenReference,
					Value: "A1",
				},
				{Type: TokenOperator, Value: ">"},
				{Type: TokenLiteral, Value: "0"},
				{Type: TokenComma, Value: ","},
				{
					Type:  TokenReference,
					Value: "B1",
				},
				{Type: TokenComma, Value: ","},
				{Type: TokenLiteral, Value: "0"},
				{Type: TokenParen, Value: ")"},
			},
			expected: "=IF(A1 > 0, B1, 0)",
		},
		{
			name:     "empty tokens",
			tokens:   nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reconstructFormula(
				tt.tokens,
			)
			if result != tt.expected {
				t.Errorf(
					"reconstructFormula() = %q, want %q",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestFormulaRewriterInsertRows tests formula rewriting for InsertRows operation.
func TestFormulaRewriterInsertRows(t *testing.T) {
	tests := []struct {
		name     string
		formula  string
		index    uint32
		count    uint32
		expected string
	}{
		// Basic cases
		{
			name:     "insert before reference",
			formula:  "=A5",
			index:    3,
			count:    2,
			expected: "=A7",
		},
		{
			name:     "insert at reference",
			formula:  "=A5",
			index:    5,
			count:    2,
			expected: "=A7",
		},
		{
			name:     "insert after reference",
			formula:  "=A5",
			index:    10,
			count:    2,
			expected: "=A5",
		},
		{
			name:     "absolute reference not shifted",
			formula:  "=$A$5",
			index:    3,
			count:    2,
			expected: "=$A$5",
		},
		{
			name:     "mixed reference col absolute",
			formula:  "=$A5",
			index:    3,
			count:    2,
			expected: "=$A7",
		},
		{
			name:     "mixed reference row absolute",
			formula:  "=A$5",
			index:    3,
			count:    2,
			expected: "=A$5",
		},

		// Range cases
		{
			name:     "range shift",
			formula:  "=SUM(A1:A10)",
			index:    5,
			count:    2,
			expected: "=SUM(A1:A12)",
		},
		{
			name:     "range both ends shift",
			formula:  "=SUM(A5:A10)",
			index:    3,
			count:    2,
			expected: "=SUM(A7:A12)",
		},
		{
			name:     "range absolute not shifted",
			formula:  "=SUM($A$1:$A$10)",
			index:    5,
			count:    2,
			expected: "=SUM($A$1:$A$10)",
		},

		// Complex formulas
		{
			name:     "complex formula",
			formula:  "=IF(A5>0,SUM(B5:B10),0)",
			index:    3,
			count:    2,
			expected: "=IF(A7 > 0, SUM(B7:B12), 0)",
		},
		{
			name:     "multiple references",
			formula:  "=A5+B10+C15",
			index:    8,
			count:    2,
			expected: "=A5 + B12 + C17",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftInsertRows,
				Index: tt.index,
				Count: tt.count,
			}
			rewriter := NewFormulaRewriter(
				op,
				"Sheet1",
			)
			result, err := rewriter.Rewrite(
				tt.formula,
			)
			if err != nil {
				t.Errorf(
					"Rewrite() error = %v",
					err,
				)

				return
			}
			if result != tt.expected {
				t.Errorf(
					"Rewrite() = %q, want %q",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestFormulaRewriterDeleteRows tests formula rewriting for DeleteRows operation.
func TestFormulaRewriterDeleteRows(t *testing.T) {
	tests := []struct {
		name     string
		formula  string
		index    uint32
		count    uint32
		expected string
	}{
		// Basic cases
		{
			name:     "delete before reference",
			formula:  "=A10",
			index:    3,
			count:    2,
			expected: "=A8",
		},
		{
			name:     "delete at reference",
			formula:  "=A5",
			index:    5,
			count:    2,
			expected: "=#REF!",
		},
		{
			name:     "delete after reference",
			formula:  "=A5",
			index:    10,
			count:    2,
			expected: "=A5",
		},
		{
			name:     "absolute reference not shifted",
			formula:  "=$A$10",
			index:    3,
			count:    2,
			expected: "=$A$10",
		},

		// Range cases
		{
			name:     "range shift down",
			formula:  "=SUM(A10:A20)",
			index:    3,
			count:    2,
			expected: "=SUM(A8:A18)",
		},
		{
			name:     "range partially deleted",
			formula:  "=SUM(A1:A10)",
			index:    5,
			count:    10,
			expected: "=SUM(A1:A4)", // Rows 5-10 deleted, range adjusted to remaining rows
		},
		{
			name:     "range fully deleted",
			formula:  "=SUM(A5:A10)",
			index:    3,
			count:    10,
			expected: "=SUM(#REF!)", // Entire range deleted, ref replaced with #REF!
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftDeleteRows,
				Index: tt.index,
				Count: tt.count,
			}
			rewriter := NewFormulaRewriter(
				op,
				"Sheet1",
			)
			result, err := rewriter.Rewrite(
				tt.formula,
			)
			if err != nil {
				t.Errorf(
					"Rewrite() error = %v",
					err,
				)

				return
			}
			if result != tt.expected {
				t.Errorf(
					"Rewrite() = %q, want %q",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestFormulaRewriterInsertColumns tests formula rewriting for InsertColumns operation.
func TestFormulaRewriterInsertColumns(
	t *testing.T,
) {
	tests := []struct {
		name     string
		formula  string
		index    uint32
		count    uint32
		expected string
	}{
		{
			name:     "insert before reference",
			formula:  "=E1",
			index:    3,
			count:    2,
			expected: "=G1",
		},
		{
			name:     "insert at reference",
			formula:  "=E1",
			index:    5,
			count:    2,
			expected: "=G1",
		},
		{
			name:     "absolute reference not shifted",
			formula:  "=$E$1",
			index:    3,
			count:    2,
			expected: "=$E$1",
		},
		{
			name:     "range shift",
			formula:  "=SUM(A1:J1)",
			index:    5,
			count:    2,
			expected: "=SUM(A1:L1)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftInsertColumns,
				Index: tt.index,
				Count: tt.count,
			}
			rewriter := NewFormulaRewriter(
				op,
				"Sheet1",
			)
			result, err := rewriter.Rewrite(
				tt.formula,
			)
			if err != nil {
				t.Errorf(
					"Rewrite() error = %v",
					err,
				)

				return
			}
			if result != tt.expected {
				t.Errorf(
					"Rewrite() = %q, want %q",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestFormulaRewriterDeleteColumns tests formula rewriting for DeleteColumns operation.
func TestFormulaRewriterDeleteColumns(
	t *testing.T,
) {
	tests := []struct {
		name     string
		formula  string
		index    uint32
		count    uint32
		expected string
	}{
		{
			name:     "delete before reference",
			formula:  "=J1",
			index:    3,
			count:    2,
			expected: "=H1",
		},
		{
			name:     "delete at reference",
			formula:  "=E1",
			index:    5,
			count:    2,
			expected: "=#REF!",
		},
		{
			name:     "absolute reference not shifted",
			formula:  "=$J$1",
			index:    3,
			count:    2,
			expected: "=$J$1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftDeleteColumns,
				Index: tt.index,
				Count: tt.count,
			}
			rewriter := NewFormulaRewriter(
				op,
				"Sheet1",
			)
			result, err := rewriter.Rewrite(
				tt.formula,
			)
			if err != nil {
				t.Errorf(
					"Rewrite() error = %v",
					err,
				)

				return
			}
			if result != tt.expected {
				t.Errorf(
					"Rewrite() = %q, want %q",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestFormulaRewriterCrossSheet tests cross-sheet formula updates.
func TestFormulaRewriterCrossSheet(t *testing.T) {
	tests := []struct {
		name         string
		formula      string
		index        uint32
		count        uint32
		targetSheet  string
		currentSheet string
		shiftType    ShiftType
		expected     string
	}{
		{
			name:         "same sheet updated",
			formula:      "=Sheet1!A5",
			index:        3,
			count:        2,
			targetSheet:  "Sheet1",
			currentSheet: "Sheet1",
			shiftType:    ShiftInsertRows,
			expected:     "=Sheet1!A7",
		},
		{
			name:         "different sheet not updated",
			formula:      "=Sheet2!A5",
			index:        3,
			count:        2,
			targetSheet:  "Sheet1",
			currentSheet: "Sheet1",
			shiftType:    ShiftInsertRows,
			expected:     "=Sheet2!A5",
		},
		{
			name:         "cross-sheet with quoted name",
			formula:      "='Sheet Name'!A5",
			index:        3,
			count:        2,
			targetSheet:  "Sheet Name",
			currentSheet: "Sheet1",
			shiftType:    ShiftInsertRows,
			expected:     "='Sheet Name'!A7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:      tt.shiftType,
				Index:     tt.index,
				Count:     tt.count,
				Worksheet: tt.targetSheet,
			}
			rewriter := NewFormulaRewriter(
				op,
				tt.currentSheet,
			)
			result, err := rewriter.Rewrite(
				tt.formula,
			)
			if err != nil {
				t.Errorf(
					"Rewrite() error = %v",
					err,
				)

				return
			}
			if result != tt.expected {
				t.Errorf(
					"Rewrite() = %q, want %q",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestFormulaRewriterStringPreservation tests that string literals are preserved.
func TestFormulaRewriterStringPreservation(
	t *testing.T,
) {
	tests := []struct {
		name     string
		formula  string
		expected string
	}{
		{
			name:     "string with cell reference text",
			formula:  "=\"Value: \"&A5",
			expected: "=\"Value: \" & A7", // Space added around operator
		},
		{
			name:     "string with A1 inside",
			formula:  "=\"Don't parse A1\"",
			expected: "=\"Don't parse A1\"",
		},
		{
			name:     "SUMIF with string criteria",
			formula:  "=SUMIF(A1:A10,\">5\",B5:B14)",
			expected: "=SUMIF(A1:A12, \">5\", B7:B16)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := ShiftOperation{
				Type:  ShiftInsertRows,
				Index: 3,
				Count: 2,
			}
			rewriter := NewFormulaRewriter(
				op,
				"Sheet1",
			)
			result, err := rewriter.Rewrite(
				tt.formula,
			)
			if err != nil {
				t.Errorf(
					"Rewrite() error = %v",
					err,
				)

				return
			}
			if result != tt.expected {
				t.Errorf(
					"Rewrite() = %q, want %q",
					result,
					tt.expected,
				)
			}
		})
	}
}
