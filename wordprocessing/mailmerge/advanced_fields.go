// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	commaSuffix           = ","
	defaultGreeting       = "Dear Sir or Madam,"
	defaultGreetingPrefix = "Dear "
	minSkipIfTokens       = 4
	floatBitSize          = 64
)

// generateGreetingLine generates a greeting line based on data source fields.
// It reads common name fields (Title, FirstName, LastName) and formats them
// according to the switches provided.
//
// Supported switches:
//   - \f "format string" - custom format prefix (e.g., "Dear ")
//
// Default behavior:
//   - If Title and LastName exist: "Dear [Title] [LastName],"
//   - If only FirstName and LastName exist: "Dear [FirstName] [LastName],"
//   - If no name fields: "Dear Sir or Madam,"
//
//nolint:revive // max-control-nesting: greeting line logic requires conditional structure
func generateGreetingLine(
	ds DataSource,
	switches map[string]string,
) string {
	// Get name components from data source
	title, _ := ds.Get("Title")
	firstName, _ := ds.Get("FirstName")
	lastName, _ := ds.Get("LastName")

	// Get custom format prefix from \f switch
	formatPrefix := defaultGreetingPrefix
	if customFormat, ok := switches["\\f"]; ok &&
		customFormat != "" {
		formatPrefix = customFormat
	}

	// Build greeting based on available fields
	var greeting string

	switch {
	case title != "" && lastName != "":
		// Use Title + LastName (most formal)
		greeting = formatPrefix + title + " " + lastName + commaSuffix
	case firstName != "" && lastName != "":
		// Use FirstName + LastName
		greeting = formatPrefix + firstName + " " + lastName + commaSuffix
	case firstName != "":
		// Use FirstName only
		greeting = formatPrefix + firstName + commaSuffix
	case lastName != "":
		// Use LastName only
		greeting = formatPrefix + lastName + commaSuffix
	default:
		// No name fields available, use generic greeting
		greeting = defaultGreeting
	}

	return greeting
}

// FieldNameSkipIf is the field name constant for SKIPIF conditional fields.
//
// SKIPIF fields allow conditional skipping of records during mail merge based on
// field value comparisons. When a SKIPIF condition evaluates to true, the current
// record is excluded from the merge output.
//
// This is only supported in ExecuteToDocuments(). In Execute() (single merged document),
// SKIPIF fields are ignored since all records are combined into one document.
const FieldNameSkipIf = "SKIPIF"

// evaluateSkipIf evaluates a SKIPIF field condition and returns true if the record should be skipped.
//
// SKIPIF syntax: "SKIPIF fieldname operator value"
//
// Supported operators:
//   - = (equal)
//   - <> (not equal)
//   - < (less than)
//   - > (greater than)
//   - <= (less than or equal)
//   - >= (greater than or equal)
//
// Examples:
//   - "SKIPIF Status = Active" - skip if Status equals "Active"
//   - "SKIPIF Age < 18" - skip if Age is less than 18
//   - "SKIPIF Country <> USA" - skip if Country is not "USA"
func evaluateSkipIf(
	ds DataSource,
	fieldCode string,
) (bool, error) {
	// Parse the SKIPIF condition
	fieldName, operator, expectedValue, err := parseSkipIfCondition(
		fieldCode,
	)
	if err != nil {
		return false, err
	}

	// Get the actual value from data source
	actualValue, err := ds.Get(fieldName)
	if err != nil {
		// If field doesn't exist, treat as empty string
		actualValue = ""
	}

	// Evaluate the condition
	return evaluateCondition(
		actualValue,
		operator,
		expectedValue,
	)
}

// parseSkipIfCondition parses a SKIPIF field code into its components.
//
// Examples:
//   - "SKIPIF Status = Active" -> ("Status", "=", "Active", nil)
//   - "SKIPIF Age >= 18" -> ("Age", ">=", "18", nil)
//   - "SKIPIF Name <> \"John Doe\"" -> ("Name", "<>", "John Doe", nil)
//
//nolint:revive // function-result-limit: 4 return values needed for meaningful error handling
func parseSkipIfCondition(
	code string,
) (fieldName, operator, value string, err error) {
	// Tokenize the field code
	tokens := tokenizeFieldCode(code)
	if len(tokens) < minSkipIfTokens {
		return "", "", "", errors.New(
			"invalid SKIPIF syntax: expected at least 4 tokens",
		)
	}

	// First token should be "SKIPIF"
	if !strings.EqualFold(
		tokens[0],
		FieldNameSkipIf,
	) {
		return "", "", "", fmt.Errorf(
			"expected SKIPIF, got %q",
			tokens[0],
		)
	}

	// Second token is the field name
	fieldName = tokens[1]

	// Third token is the operator
	operator = tokens[2]

	// Remaining tokens form the value (may be multiple tokens if quoted)
	// Join all remaining tokens with spaces
	value = strings.Join(tokens[3:], " ")

	return fieldName, operator, value, nil
}

// evaluateCondition evaluates a comparison between actual and expected values.
func evaluateCondition(
	actual, operator, expected string,
) (bool, error) {
	switch operator {
	case "=", "==":
		// Case-insensitive string equality
		return strings.EqualFold(
			actual,
			expected,
		), nil

	case "<>", "!=":
		// Case-insensitive string inequality
		return !strings.EqualFold(
			actual,
			expected,
		), nil

	case "<", ">", "<=", ">=":
		// Numeric comparison
		return evaluateNumericCondition(
			actual,
			operator,
			expected,
		)

	default:
		return false, fmt.Errorf(
			"unsupported operator: %q",
			operator,
		)
	}
}

// evaluateNumericCondition evaluates a numeric comparison.
// If either value cannot be parsed as a number, it falls back to string comparison.
func evaluateNumericCondition(
	actual, operator, expected string,
) (bool, error) {
	// Try to parse as numbers
	actualNum, actualErr := strconv.ParseFloat(
		strings.TrimSpace(actual),
		floatBitSize,
	)
	expectedNum, expectedErr := strconv.ParseFloat(
		strings.TrimSpace(expected),
		floatBitSize,
	)

	// If both parse successfully, do numeric comparison
	if actualErr == nil && expectedErr == nil {
		switch operator {
		case "<":
			return actualNum < expectedNum, nil
		case ">":
			return actualNum > expectedNum, nil
		case "<=":
			return actualNum <= expectedNum, nil
		case ">=":
			return actualNum >= expectedNum, nil
		default:
			return false, fmt.Errorf(
				"unsupported numeric operator: %q",
				operator,
			)
		}
	}

	// Fall back to string comparison (lexicographic)
	switch operator {
	case "<":
		return actual < expected, nil
	case ">":
		return actual > expected, nil
	case "<=":
		return actual <= expected, nil
	case ">=":
		return actual >= expected, nil
	default:
		return false, fmt.Errorf(
			"unsupported comparison operator: %q",
			operator,
		)
	}
}
