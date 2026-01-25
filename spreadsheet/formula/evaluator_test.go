package formula

import (
	"testing"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

func TestEvaluateCellFormula(t *testing.T) {
	// Test literal values
	tests := []struct {
		name     string
		cell     interface{}
		expected Value
	}{
		{
			name:     "string literal",
			cell:     "hello",
			expected: NewStringValue("hello"),
		},
		{
			name:     "number literal",
			cell:     42.0,
			expected: NewNumberValue(42),
		},
		{
			name:     "boolean literal",
			cell:     true,
			expected: NewBooleanValue(true),
		},
		{
			name:     "nil value",
			cell:     nil,
			expected: NewEmptyValue(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := EvaluateCellFormula(test.cell, nil)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if !result.Equal(test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
				return
			}
		})
	}
}

func TestEvaluateFormula(t *testing.T) {
	// Test simple formulas
	testCases := []struct {
		name     string
		formula  string
		expected Value
	}{
		{
			name:     "simple addition",
			formula:  "=1+2",
			expected: NewNumberValue(3),
		},
		{
			name:     "simple subtraction",
			formula:  "=5-3",
			expected: NewNumberValue(2),
		},
		{
			name:     "simple multiplication",
			formula:  "=3*4",
			expected: NewNumberValue(12),
		},
		{
			name:     "simple division",
			formula:  "=12/4",
			expected: NewNumberValue(3),
		},
		{
			name:     "error - division by zero",
			formula:  "=1/0",
			expected: NewErrorValue("#DIV/0!"),
		},
		{
			name:     "function call - SUM",
			formula:  "=SUM(1,2,3)",
			expected: NewNumberValue(6),
		},
		{
			name:     "function call - AVERAGE",
			formula:  "=AVERAGE(2,4,6)",
			expected: NewNumberValue(4),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := EvaluateFormula(test.formula, nil)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if !result.Equal(test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
				return
			}
		})
	}
}
