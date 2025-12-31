// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"strings"
	"testing"
)

func TestGenerateGreetingLine_WithTitleAndLastName(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{
			"Title":     "Dr.",
			"FirstName": "John",
			"LastName":  "Smith",
		},
	})
	_ = ds.Open()
	ds.Next()

	greeting := generateGreetingLine(
		ds,
		make(map[string]string),
	)
	expected := "Dear Dr. Smith,"
	if greeting != expected {
		t.Errorf(
			"Expected %q, got %q",
			expected,
			greeting,
		)
	}
}

func TestGenerateGreetingLine_WithFirstAndLastName(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{
			"FirstName": "Jane",
			"LastName":  "Doe",
		},
	})
	_ = ds.Open()
	ds.Next()

	greeting := generateGreetingLine(
		ds,
		make(map[string]string),
	)
	expected := "Dear Jane Doe,"
	if greeting != expected {
		t.Errorf(
			"Expected %q, got %q",
			expected,
			greeting,
		)
	}
}

func TestGenerateGreetingLine_WithFirstNameOnly(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{
			"FirstName": "Alice",
		},
	})
	_ = ds.Open()
	ds.Next()

	greeting := generateGreetingLine(
		ds,
		make(map[string]string),
	)
	expected := "Dear Alice,"
	if greeting != expected {
		t.Errorf(
			"Expected %q, got %q",
			expected,
			greeting,
		)
	}
}

func TestGenerateGreetingLine_WithLastNameOnly(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{
			"LastName": "Johnson",
		},
	})
	_ = ds.Open()
	ds.Next()

	greeting := generateGreetingLine(
		ds,
		make(map[string]string),
	)
	expected := "Dear Johnson,"
	if greeting != expected {
		t.Errorf(
			"Expected %q, got %q",
			expected,
			greeting,
		)
	}
}

func TestGenerateGreetingLine_WithNoNameFields(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{
			"Email": "test@example.com",
		},
	})
	_ = ds.Open()
	ds.Next()

	greeting := generateGreetingLine(
		ds,
		make(map[string]string),
	)
	expected := "Dear Sir or Madam,"
	if greeting != expected {
		t.Errorf(
			"Expected %q, got %q",
			expected,
			greeting,
		)
	}
}

func TestGenerateGreetingLine_WithCustomFormat(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{
			"FirstName": "Bob",
			"LastName":  "Wilson",
		},
	})
	_ = ds.Open()
	ds.Next()

	switches := map[string]string{
		"\\f": "Hello ",
	}
	greeting := generateGreetingLine(ds, switches)
	expected := "Hello Bob Wilson,"
	if greeting != expected {
		t.Errorf(
			"Expected %q, got %q",
			expected,
			greeting,
		)
	}
}

func TestParseSkipIfCondition_Equality(
	t *testing.T,
) {
	fieldCode := "SKIPIF Status = Active"
	fieldName, operator, value, err := parseSkipIfCondition(
		fieldCode,
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if fieldName != "Status" {
		t.Errorf(
			"Expected fieldName='Status', got %q",
			fieldName,
		)
	}
	if operator != "=" {
		t.Errorf(
			"Expected operator='=', got %q",
			operator,
		)
	}
	if value != "Active" {
		t.Errorf(
			"Expected value='Active', got %q",
			value,
		)
	}
}

func TestParseSkipIfCondition_Inequality(
	t *testing.T,
) {
	fieldCode := "SKIPIF Country <> USA"
	fieldName, operator, value, err := parseSkipIfCondition(
		fieldCode,
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if fieldName != "Country" {
		t.Errorf(
			"Expected fieldName='Country', got %q",
			fieldName,
		)
	}
	if operator != "<>" {
		t.Errorf(
			"Expected operator='<>', got %q",
			operator,
		)
	}
	if value != testValueUSA {
		t.Errorf(
			"Expected value='%s', got %q",
			testValueUSA,
			value,
		)
	}
}

func TestParseSkipIfCondition_NumericComparison(
	t *testing.T,
) {
	testCases := []struct {
		fieldCode        string
		expectedField    string
		expectedOperator string
		expectedValue    string
	}{
		{"SKIPIF Age < 18", "Age", "<", "18"},
		{"SKIPIF Age > 65", "Age", ">", "65"},
		{"SKIPIF Age <= 21", "Age", "<=", "21"},
		{"SKIPIF Age >= 18", "Age", ">=", "18"},
	}

	for _, tc := range testCases {
		t.Run(tc.fieldCode, func(t *testing.T) {
			fieldName, operator, value, err := parseSkipIfCondition(
				tc.fieldCode,
			)
			if err != nil {
				t.Fatalf(
					"Unexpected error: %v",
					err,
				)
			}

			if fieldName != tc.expectedField {
				t.Errorf(
					"Expected fieldName=%q, got %q",
					tc.expectedField,
					fieldName,
				)
			}
			if operator != tc.expectedOperator {
				t.Errorf(
					"Expected operator=%q, got %q",
					tc.expectedOperator,
					operator,
				)
			}
			if value != tc.expectedValue {
				t.Errorf(
					"Expected value=%q, got %q",
					tc.expectedValue,
					value,
				)
			}
		})
	}
}

func TestEvaluateCondition_Equality(
	t *testing.T,
) {
	testCases := []struct {
		actual   string
		expected string
		result   bool
	}{
		{"Active", "Active", true},
		{
			"active",
			"ACTIVE",
			true,
		}, // Case-insensitive
		{"Active", "Inactive", false},
		{"", "", true},
	}

	for _, tc := range testCases {
		result, err := evaluateCondition(
			tc.actual,
			"=",
			tc.expected,
		)
		if err != nil {
			t.Errorf(
				"Unexpected error for %q = %q: %v",
				tc.actual,
				tc.expected,
				err,
			)
		}
		if result != tc.result {
			t.Errorf(
				"Expected %q = %q to be %v, got %v",
				tc.actual,
				tc.expected,
				tc.result,
				result,
			)
		}
	}
}

func TestEvaluateCondition_Inequality(
	t *testing.T,
) {
	testCases := []struct {
		actual   string
		expected string
		result   bool
	}{
		{"Active", "Inactive", true},
		{
			"active",
			"ACTIVE",
			false,
		}, // Case-insensitive
		{"USA", "Canada", true},
	}

	for _, tc := range testCases {
		result, err := evaluateCondition(
			tc.actual,
			"<>",
			tc.expected,
		)
		if err != nil {
			t.Errorf(
				"Unexpected error for %q <> %q: %v",
				tc.actual,
				tc.expected,
				err,
			)
		}
		if result != tc.result {
			t.Errorf(
				"Expected %q <> %q to be %v, got %v",
				tc.actual,
				tc.expected,
				tc.result,
				result,
			)
		}
	}
}

func TestEvaluateCondition_NumericComparison(
	t *testing.T,
) {
	testCases := []struct {
		actual   string
		operator string
		expected string
		result   bool
	}{
		{"20", "<", "18", false},
		{"15", "<", "18", true},
		{"20", ">", "18", true},
		{"15", ">", "18", false},
		{"18", "<=", "18", true},
		{"17", "<=", "18", true},
		{"19", "<=", "18", false},
		{"18", ">=", "18", true},
		{"19", ">=", "18", true},
		{"17", ">=", "18", false},
	}

	for _, tc := range testCases {
		result, err := evaluateCondition(
			tc.actual,
			tc.operator,
			tc.expected,
		)
		if err != nil {
			t.Errorf(
				"Unexpected error for %q %s %q: %v",
				tc.actual,
				tc.operator,
				tc.expected,
				err,
			)
		}
		if result != tc.result {
			t.Errorf(
				"Expected %q %s %q to be %v, got %v",
				tc.actual,
				tc.operator,
				tc.expected,
				tc.result,
				result,
			)
		}
	}
}

func TestEvaluateSkipIf_SimpleEquality(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{
			"Status": "Active",
		},
	})
	_ = ds.Open()
	ds.Next()

	// Should skip when Status = Active
	skip, err := evaluateSkipIf(
		ds,
		"SKIPIF Status = Active",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !skip {
		t.Error("Expected record to be skipped")
	}

	// Should not skip when Status != Inactive
	skip, err = evaluateSkipIf(
		ds,
		"SKIPIF Status = Inactive",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if skip {
		t.Error(
			"Expected record not to be skipped",
		)
	}
}

func TestEvaluateSkipIf_NumericComparison(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{
			"Age": "25",
		},
	})
	_ = ds.Open()
	ds.Next()

	testCases := []struct {
		fieldCode   string
		shouldSkip  bool
		description string
	}{
		{
			"SKIPIF Age < 18",
			false,
			"25 is not less than 18",
		},
		{
			"SKIPIF Age > 18",
			true,
			"25 is greater than 18",
		},
		{
			"SKIPIF Age <= 25",
			true,
			"25 is less than or equal to 25",
		},
		{
			"SKIPIF Age >= 25",
			true,
			"25 is greater than or equal to 25",
		},
		{
			"SKIPIF Age < 30",
			true,
			"25 is less than 30",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			skip, err := evaluateSkipIf(
				ds,
				tc.fieldCode,
			)
			if err != nil {
				t.Fatalf(
					"Unexpected error: %v",
					err,
				)
			}
			if skip != tc.shouldSkip {
				t.Errorf(
					"Expected shouldSkip=%v for %q, got %v",
					tc.shouldSkip,
					tc.fieldCode,
					skip,
				)
			}
		})
	}
}

func TestEvaluateSkipIf_MissingField(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{
			"Name": "John",
		},
	})
	_ = ds.Open()
	ds.Next()

	// Missing field should be treated as empty string
	// Comparing with empty string
	skip, err := evaluateSkipIf(
		ds,
		"SKIPIF Age = \"\"",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !skip {
		t.Error(
			"Expected record to be skipped (missing field treated as empty)",
		)
	}
}

func TestEvaluateSkipIf_CaseInsensitive(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{
			"Status": "active",
		},
	})
	_ = ds.Open()
	ds.Next()

	// Should skip with case-insensitive comparison
	skip, err := evaluateSkipIf(
		ds,
		"SKIPIF Status = ACTIVE",
	)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !skip {
		t.Error(
			"Expected record to be skipped (case-insensitive)",
		)
	}
}

func TestEvaluateCondition_StringComparison(
	t *testing.T,
) {
	// Test string comparison when values are not numeric
	testCases := []struct {
		actual   string
		operator string
		expected string
		result   bool
	}{
		{
			"apple",
			"<",
			"banana",
			true,
		}, // Lexicographic comparison
		{"banana", ">", "apple", true},
		{"apple", "<=", "apple", true},
		{"zebra", ">=", "apple", true},
	}

	for _, tc := range testCases {
		result, err := evaluateCondition(
			tc.actual,
			tc.operator,
			tc.expected,
		)
		if err != nil {
			t.Errorf(
				"Unexpected error for %q %s %q: %v",
				tc.actual,
				tc.operator,
				tc.expected,
				err,
			)
		}
		if result != tc.result {
			t.Errorf(
				"Expected %q %s %q to be %v, got %v",
				tc.actual,
				tc.operator,
				tc.expected,
				tc.result,
				result,
			)
		}
	}
}

func TestParseSkipIfCondition_InvalidSyntax(
	t *testing.T,
) {
	testCases := []string{
		"SKIPIF",          // Missing field name and operator
		"SKIPIF Status",   // Missing operator and value
		"SKIPIF Status =", // Missing value (should parse but value will be empty)
	}

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			_, _, _, err := parseSkipIfCondition(
				tc,
			)
			if err == nil &&
				!strings.Contains(
					tc,
					"Status =",
				) {
				t.Error(
					"Expected error for invalid SKIPIF syntax",
				)
			}
		})
	}
}
