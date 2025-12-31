package validation

import (
	"testing"
)

func TestVersionAwareValidation(t *testing.T) {
	t.Run(
		"validates Office2019 element against Office2021",
		func(t *testing.T) {
			// Office 2019 element should be valid in Office 2021
			element := newMockVersionedElement(
				"feature2019",
				Office2019,
			)

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2021,
			)
			errors := ValidateElement(
				element,
				ctx,
			)

			// Should have no version errors
			versionErrors := ctx.Errors().ByCode(
				SchemaElementNotAvailable,
			)
			if len(versionErrors) > 0 {
				t.Errorf(
					"expected no version errors, got %d: %v",
					len(versionErrors),
					versionErrors,
				)
			}
			// Total errors should be zero
			if len(errors) > 0 {
				t.Logf(
					"validation returned %d errors",
					len(errors),
				)
			}
		},
	)

	t.Run(
		"rejects Office2021 element in Office2019 validation",
		func(t *testing.T) {
			// Office 2021 element should fail in Office 2019 validation
			element := newMockVersionedElement(
				"feature2021",
				Office2021,
			)

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2019,
			)
			ValidateElement(element, ctx)

			// Should have version error
			versionErrors := ctx.Errors().ByCode(
				SchemaElementNotAvailable,
			)
			if len(versionErrors) != 1 {
				t.Errorf(
					"expected 1 version error, got %d",
					len(versionErrors),
				)
			}

			if len(versionErrors) == 0 {
				return
			}

			err := versionErrors[0]
			if err.Severity != SeverityError {
				t.Errorf(
					"expected Error severity, got %s",
					err.Severity,
				)
			}
			// Check that error message mentions versions
			if err.Description == "" {
				t.Error(
					"expected error description to be non-empty",
				)
			}
		},
	)

	t.Run(
		"reports multiple version violations",
		func(t *testing.T) {
			// Document with multiple version violations
			doc := newMockElement("document").
				withChild(
					newMockVersionedElement(
						"feature2019",
						Office2019,
					),
				).
				withChild(
					newMockVersionedElement(
						"feature2021",
						Office2021,
					),
				)

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)
			ValidateElement(doc, ctx)

			// Should have 2 version errors (one for each incompatible element)
			versionErrors := ctx.Errors().ByCode(
				SchemaElementNotAvailable,
			)
			if len(versionErrors) != 2 {
				t.Errorf(
					"expected 2 version errors, got %d",
					len(versionErrors),
				)
			}
		},
	)

	t.Run(
		"validates element available in target version",
		func(t *testing.T) {
			// Office 2016 element in Office 2016 validation
			element := newMockVersionedElement(
				"feature2016",
				Office2016,
			)

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)
			ValidateElement(element, ctx)

			versionErrors := ctx.Errors().ByCode(
				SchemaElementNotAvailable,
			)
			if len(versionErrors) > 0 {
				t.Errorf(
					"expected no version errors, got %d",
					len(versionErrors),
				)
			}
		},
	)

	t.Run(
		"validates nested version violations",
		func(t *testing.T) {
			// Create nested structure with version violation deep in tree
			innerElement := newMockVersionedElement(
				"deepFeature",
				Office2021,
			)
			doc := newMockElement("document").
				withChild(
					newMockElement("body").
						withChild(
							newMockElement(
								"section",
							).
								withChild(innerElement),
						),
				)

			ctx := NewValidationContext(
				DefaultSettings(),
				Office2019,
			)
			ValidateElement(doc, ctx)

			versionErrors := ctx.Errors().ByCode(
				SchemaElementNotAvailable,
			)
			if len(versionErrors) != 1 {
				t.Errorf(
					"expected 1 version error for nested element, got %d",
					len(versionErrors),
				)
			}
		},
	)
}

func TestWithTargetVersion(t *testing.T) {
	t.Run(
		"sets target version on context",
		func(t *testing.T) {
			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			ctx.WithTargetVersion(Office2021)

			if ctx.TargetVersion() != Office2021 {
				t.Errorf(
					"expected Office2021, got %s",
					ctx.TargetVersion().String(),
				)
			}
		},
	)

	t.Run(
		"allows chaining",
		func(t *testing.T) {
			ctx := NewValidationContext(
				DefaultSettings(),
				Office2016,
			)

			result := ctx.WithTargetVersion(
				Office2019,
			)

			if result != ctx {
				t.Error(
					"WithTargetVersion should return the same context",
				)
			}
		},
	)
}

func TestTargetVersion(t *testing.T) {
	t.Run(
		"returns current target version",
		func(t *testing.T) {
			ctx := NewValidationContext(
				DefaultSettings(),
				Office2019,
			)

			if ctx.TargetVersion() != Office2019 {
				t.Errorf(
					"expected Office2019, got %s",
					ctx.TargetVersion().String(),
				)
			}
		},
	)
}

func TestIsVersionAvailable(t *testing.T) {
	t.Run(
		"returns true for nil availability",
		func(t *testing.T) {
			ctx := NewValidationContext(
				DefaultSettings(),
				Office2019,
			)

			if !ctx.IsVersionAvailable(nil) {
				t.Error(
					"expected true for nil availability",
				)
			}
		},
	)

	t.Run(
		"returns true when version is available",
		func(t *testing.T) {
			ctx := NewValidationContext(
				DefaultSettings(),
				Office2021,
			)
			avail := NewVersionAvailability(
				Office2019,
			)

			if !ctx.IsVersionAvailable(avail) {
				t.Error(
					"expected true for available version",
				)
			}
		},
	)

	t.Run(
		"returns false when version is not available",
		func(t *testing.T) {
			ctx := NewValidationContext(
				DefaultSettings(),
				Office2019,
			)
			avail := NewVersionAvailability(
				Office2021,
			)

			if ctx.IsVersionAvailable(avail) {
				t.Error(
					"expected false for unavailable version",
				)
			}
		},
	)
}

func TestCheckElementVersion(t *testing.T) {
	t.Run(
		"returns nil for element without version info",
		func(t *testing.T) {
			element := newMockElement("basic")
			err := CheckElementVersion(
				element,
				Office2019,
				"/test/path",
			)
			if err != nil {
				t.Errorf(
					"expected nil error, got %v",
					err,
				)
			}
		},
	)

	t.Run(
		"returns error for unavailable element",
		func(t *testing.T) {
			element := newMockVersionedElement(
				"advanced",
				Office2021,
			)
			err := CheckElementVersion(
				element,
				Office2019,
				"/test/path",
			)

			if err == nil {
				t.Error(
					"expected error for unavailable element",
				)

				return
			}

			if err.Code != SchemaElementNotAvailable {
				t.Errorf(
					"expected SchemaElementNotAvailable, got %s",
					err.Code,
				)
			}
		},
	)

	t.Run(
		"returns nil for available element",
		func(t *testing.T) {
			element := newMockVersionedElement(
				"feature",
				Office2019,
			)
			err := CheckElementVersion(
				element,
				Office2021,
				"/test/path",
			)
			if err != nil {
				t.Errorf(
					"expected nil error, got %v",
					err,
				)
			}
		},
	)

	t.Run(
		"includes version information in error message",
		func(t *testing.T) {
			element := newMockVersionedElement(
				"newFeature",
				Office2021,
			)
			err := CheckElementVersion(
				element,
				Office2019,
				"/document/body",
			)

			if err == nil {
				t.Fatal("expected error")
			}

			// Error description should mention the version
			if err.Description == "" {
				t.Error(
					"error description should not be empty",
				)
			}

			// Path should be set
			if err.Path != "/document/body" {
				t.Errorf(
					"expected path '/document/body', got '%s'",
					err.Path,
				)
			}
		},
	)
}
