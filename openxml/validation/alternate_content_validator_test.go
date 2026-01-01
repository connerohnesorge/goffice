package validation

import (
	"testing"
)

//nolint:revive // test function with many subtests
func TestAlternateContentValidator(t *testing.T) {
	validator := NewAlternateContentValidator()

	t.Run(
		"accepts valid AlternateContent with Choice and Fallback",
		func(t *testing.T) {
			ac := newMockAlternateContent().
				withChild(newMockChoice("w14")).
				withChild(newMockFallback())

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			errors := validator.Validate(ctx, ac)

			if len(errors) > 0 {
				t.Errorf(
					"expected no errors, got %d: %v",
					len(errors),
					errors,
				)
			}
		},
	)

	t.Run(
		"rejects AlternateContent with Fallback before Choice",
		func(t *testing.T) {
			ac := newMockAlternateContent().
				withChild(newMockFallback()).
				withChild(newMockChoice("w14"))

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			errors := validator.Validate(ctx, ac)

			if len(errors) == 0 {
				t.Error(
					"expected error for Fallback before Choice",
				)
			}

			// Check for SchemaInvalidChildOrder error
			hasOrderError := false
			for _, err := range errors {
				if err.Code == SchemaInvalidChildOrder {
					hasOrderError = true
				}
			}

			if !hasOrderError {
				t.Errorf(
					"expected SchemaInvalidChildOrder error, got: %v",
					errors,
				)
			}
		},
	)

	t.Run(
		"rejects AlternateContent with multiple Fallbacks",
		func(t *testing.T) {
			ac := newMockAlternateContent().
				withChild(newMockChoice("w14")).
				withChild(newMockFallback()).
				withChild(newMockFallback())

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			errors := validator.Validate(ctx, ac)

			if len(errors) == 0 {
				t.Error(
					"expected error for multiple Fallbacks",
				)
			}

			// Check for SchemaTooManyElements error
			hasTooManyError := false
			for _, err := range errors {
				if err.Code == SchemaTooManyElements {
					hasTooManyError = true
				}
			}

			if !hasTooManyError {
				t.Errorf(
					"expected SchemaTooManyElements error, got: %v",
					errors,
				)
			}
		},
	)

	t.Run(
		"rejects Choice without Requires attribute",
		func(t *testing.T) {
			ac := newMockAlternateContent().
				withChild(newMockChoice("")).
				withChild(newMockFallback())

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			errors := validator.Validate(ctx, ac)

			if len(errors) == 0 {
				t.Error(
					"expected error for missing Requires attribute",
				)
			}

			// Check for SchemaMissingRequiredAttribute error
			hasMissingError := false
			for _, err := range errors {
				if err.Code == SchemaMissingRequiredAttribute {
					hasMissingError = true
				}
			}

			if !hasMissingError {
				t.Errorf(
					"expected SchemaMissingRequiredAttribute error, got: %v",
					errors,
				)
			}
		},
	)

	t.Run(
		"rejects Choice with invalid Requires value",
		func(t *testing.T) {
			ac := newMockAlternateContent().
				withChild(newMockChoice("invalid_prefix")).
				withChild(newMockFallback())

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			errors := validator.Validate(ctx, ac)

			if len(errors) == 0 {
				t.Error(
					"expected error for invalid Requires value",
				)
			}

			// Check for SchemaInvalidAttributeValue error
			hasInvalidError := false
			for _, err := range errors {
				if err.Code == SchemaInvalidAttributeValue {
					hasInvalidError = true
				}
			}

			if !hasInvalidError {
				t.Errorf(
					"expected SchemaInvalidAttributeValue error, got: %v",
					errors,
				)
			}
		},
	)

	t.Run(
		"accepts multiple Choice elements",
		func(t *testing.T) {
			ac := newMockAlternateContent().
				withChild(newMockChoice("w15")).
				withChild(newMockChoice("w14")).
				withChild(newMockFallback())

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			errors := validator.Validate(ctx, ac)

			if len(errors) > 0 {
				t.Errorf(
					"expected no errors with multiple Choices, got %d: %v",
					len(errors),
					errors,
				)
			}
		},
	)

	t.Run(
		"rejects AlternateContent with no children",
		func(t *testing.T) {
			ac := newMockAlternateContent()

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			errors := validator.Validate(ctx, ac)

			if len(errors) == 0 {
				t.Error(
					"expected error for empty AlternateContent",
				)
			}
		},
	)

	t.Run(
		"rejects AlternateContent with only Fallback (no Choice)",
		func(t *testing.T) {
			ac := newMockAlternateContent().
				withChild(newMockFallback())

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			errors := validator.Validate(ctx, ac)

			if len(errors) == 0 {
				t.Error(
					"expected error when Choice is missing",
				)
			}

			// Check for SchemaMissingRequiredElement error
			hasMissingError := false
			for _, err := range errors {
				if err.Code == SchemaMissingRequiredElement {
					hasMissingError = true
				}
			}

			if !hasMissingError {
				t.Errorf(
					"expected SchemaMissingRequiredElement error, got: %v",
					errors,
				)
			}
		},
	)

	t.Run(
		"rejects unexpected elements in AlternateContent",
		func(t *testing.T) {
			ac := newMockAlternateContent().
				withChild(newMockChoice("w14")).
				withChild(
					newMockElement("UnexpectedElement").
						withNS("http://unknown.namespace"),
				)

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			errors := validator.Validate(ctx, ac)

			if len(errors) == 0 {
				t.Error(
					"expected error for unexpected element",
				)
			}

			// Check for SchemaUnexpectedElement error
			hasUnexpectedError := false
			for _, err := range errors {
				if err.Code == SchemaUnexpectedElement {
					hasUnexpectedError = true
				}
			}

			if !hasUnexpectedError {
				t.Errorf(
					"expected SchemaUnexpectedElement error, got: %v",
					errors,
				)
			}
		},
	)
}

func TestValidRequiresValue(t *testing.T) {
	validator := NewAlternateContentValidator()

	tests := []struct {
		value string
		valid bool
	}{
		{"w14", true},
		{"w15", true},
		{"x14", true},
		{"x15", true},
		{"a14", true},
		{"p14", true},
		{"w22", true},
		{"x365", true},
		{"invalid", false},
		{"14w", false},
		{"w", false},
		{"14", false},
		{"", false},
		{"w_14", false},
		{"W14", false},
		{"w14x", false},
	}

	for _, tt := range tests {
		result := validator.isValidRequiresValue(tt.value)
		if result != tt.valid {
			t.Errorf(
				"isValidRequiresValue(%q) = %v, want %v",
				tt.value,
				result,
				tt.valid,
			)
		}
	}
}

// Helper functions for creating mock AlternateContent elements

func newMockAlternateContent() *mockElement {
	return newMockElement("AlternateContent").
		withNS("http://schemas.openxmlformats.org/markup-compatibility/2006")
}

func newMockChoice(requires string) *mockElement {
	elem := newMockElement("Choice").
		withNS("http://schemas.openxmlformats.org/markup-compatibility/2006")
	if requires != "" {
		elem.withAttr("Requires", requires)
	}

	return elem
}

func newMockFallback() *mockElement {
	return newMockElement("Fallback").
		withNS("http://schemas.openxmlformats.org/markup-compatibility/2006")
}
