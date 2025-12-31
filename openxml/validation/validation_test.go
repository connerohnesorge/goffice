package validation

import (
	"reflect"
	"testing"
)

// Mock element for testing
type mockElement struct {
	localName    string
	namespaceURI string
	attributes   map[string]string
	children     []any
	parent       any
}

func (e *mockElement) LocalName() string { return e.localName }

func (e *mockElement) NamespaceURI() string { return e.namespaceURI }

func (e *mockElement) Parent() any { return e.parent }

func (e *mockElement) Attributes() []mockAttribute { return mapToAttrs(e.attributes) }

func (e *mockElement) Children() []any { return e.children }

type mockAttribute struct {
	name  string
	value string
}

func (a mockAttribute) LocalName() string { return a.name }

func (a mockAttribute) Value() string { return a.value }

func mapToAttrs(
	m map[string]string,
) []mockAttribute {
	attrs := make([]mockAttribute, 0, len(m))
	for k, v := range m {
		attrs = append(
			attrs,
			mockAttribute{name: k, value: v},
		)
	}

	return attrs
}

func newMockElement(
	localName string,
) *mockElement {
	return &mockElement{
		localName:  localName,
		attributes: make(map[string]string),
		children:   make([]any, 0),
	}
}

func (e *mockElement) withNS(
	ns string,
) *mockElement {
	e.namespaceURI = ns

	return e
}

func (e *mockElement) withAttr(
	name, value string,
) *mockElement {
	e.attributes[name] = value

	return e
}

func (e *mockElement) withChild(
	child any,
) *mockElement {
	// Set parent based on child type
	if mockChild, ok := child.(*mockElement); ok {
		mockChild.parent = e
	}
	e.children = append(e.children, child)

	return e
}

func TestValidationError(t *testing.T) {
	t.Run(
		"creates error with correct fields",
		func(t *testing.T) {
			err := NewValidationError(
				SchemaMissingRequiredElement,
				"Element 'body' is required",
				"/w:document",
				nil,
			)

			if err.Code != SchemaMissingRequiredElement {
				t.Errorf(
					"expected code %s, got %s",
					SchemaMissingRequiredElement,
					err.Code,
				)
			}
			if err.Severity != SeverityError {
				t.Errorf(
					"expected severity Error, got %s",
					err.Severity,
				)
			}
			if err.Path != "/w:document" {
				t.Errorf(
					"expected path /w:document, got %s",
					err.Path,
				)
			}
		},
	)

	t.Run("creates warning", func(t *testing.T) {
		err := NewValidationWarning(
			SchemaValueNotInEnumeration,
			"Value not recommended",
			"/w:document/w:body",
			nil,
		)

		if err.Severity != SeverityWarning {
			t.Errorf(
				"expected severity Warning, got %s",
				err.Severity,
			)
		}
	})

	t.Run(
		"implements error interface",
		func(t *testing.T) {
			var err error = NewValidationError(
				SchemaMissingRequiredElement,
				"test error",
				"/path",
				nil,
			)

			if err.Error() == "" {
				t.Error(
					"expected non-empty error string",
				)
			}
		},
	)
}

func TestValidationErrors(t *testing.T) {
	t.Run(
		"HasErrors returns true for errors",
		func(t *testing.T) {
			errs := ValidationErrors{
				NewValidationWarning(
					SchemaValueNotInEnumeration,
					"warning",
					"/",
					nil,
				),
				NewValidationError(
					SchemaMissingRequiredElement,
					"error",
					"/",
					nil,
				),
			}

			if !errs.HasErrors() {
				t.Error(
					"expected HasErrors to return true",
				)
			}
		},
	)

	t.Run(
		"HasErrors returns false for warnings only",
		func(t *testing.T) {
			errs := ValidationErrors{
				NewValidationWarning(
					SchemaValueNotInEnumeration,
					"warning1",
					"/",
					nil,
				),
				NewValidationWarning(
					SchemaValueNotInEnumeration,
					"warning2",
					"/",
					nil,
				),
			}

			if errs.HasErrors() {
				t.Error(
					"expected HasErrors to return false",
				)
			}
		},
	)

	t.Run(
		"Errors filters to errors only",
		func(t *testing.T) {
			errs := ValidationErrors{
				NewValidationWarning(
					SchemaValueNotInEnumeration,
					"warning",
					"/",
					nil,
				),
				NewValidationError(
					SchemaMissingRequiredElement,
					"error",
					"/",
					nil,
				),
			}

			filtered := errs.Errors()
			if len(filtered) != 1 {
				t.Errorf(
					"expected 1 error, got %d",
					len(filtered),
				)
			}
		},
	)

	t.Run(
		"Warnings filters to warnings only",
		func(t *testing.T) {
			errs := ValidationErrors{
				NewValidationWarning(
					SchemaValueNotInEnumeration,
					"warning",
					"/",
					nil,
				),
				NewValidationError(
					SchemaMissingRequiredElement,
					"error",
					"/",
					nil,
				),
			}

			filtered := errs.Warnings()
			if len(filtered) != 1 {
				t.Errorf(
					"expected 1 warning, got %d",
					len(filtered),
				)
			}
		},
	)
}

func TestValidationContext(t *testing.T) {
	t.Run(
		"creates with default settings",
		func(t *testing.T) {
			ctx := NewValidationContext(
				nil,
				Office2016,
			)

			if ctx.Settings == nil {
				t.Error(
					"expected default settings",
				)
			}
			if ctx.Version != Office2016 {
				t.Errorf(
					"expected Office2016, got %v",
					ctx.Version,
				)
			}
		},
	)

	t.Run(
		"path stack operations",
		func(t *testing.T) {
			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)

			ctx.PushPath("w:document")
			ctx.PushPath("w:body")
			ctx.PushPath("w:p")

			path := ctx.CurrentPath()
			if path != "/w:document/w:body/w:p" {
				t.Errorf(
					"expected /w:document/w:body/w:p, got %s",
					path,
				)
			}

			ctx.PopPath()
			path = ctx.CurrentPath()
			if path != "/w:document/w:body" {
				t.Errorf(
					"expected /w:document/w:body, got %s",
					path,
				)
			}
		},
	)

	t.Run(
		"tracks duplicate IDs",
		func(t *testing.T) {
			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)

			elem1 := newMockElement("p")
			elem2 := newMockElement("p")

			// First occurrence should return nil
			existing := ctx.TrackID("id1", elem1)
			if existing != nil {
				t.Error(
					"expected nil for first occurrence",
				)
			}

			// Second occurrence should return the first element
			existing = ctx.TrackID("id1", elem2)
			if existing != elem1 {
				t.Error(
					"expected first element for duplicate",
				)
			}
		},
	)

	t.Run(
		"stops on max errors",
		func(t *testing.T) {
			settings := &ValidationSettings{
				MaxErrors:       2,
				ContinueOnError: true,
			}
			ctx := NewValidationContext(
				settings,
				Microsoft365,
			)

			ctx.AddError(
				NewValidationError(
					SchemaMissingRequiredElement,
					"error1",
					"/",
					nil,
				),
			)
			if ctx.ShouldStop() {
				t.Error(
					"should not stop after 1 error",
				)
			}

			ctx.AddError(
				NewValidationError(
					SchemaMissingRequiredElement,
					"error2",
					"/",
					nil,
				),
			)
			if !ctx.ShouldStop() {
				t.Error(
					"should stop after 2 errors",
				)
			}
		},
	)

	t.Run(
		"stops on first error when ContinueOnError is false",
		func(t *testing.T) {
			settings := &ValidationSettings{
				MaxErrors:       0,
				ContinueOnError: false,
			}
			ctx := NewValidationContext(
				settings,
				Microsoft365,
			)

			ctx.AddError(
				NewValidationError(
					SchemaMissingRequiredElement,
					"error1",
					"/",
					nil,
				),
			)
			if !ctx.ShouldStop() {
				t.Error(
					"should stop after first error",
				)
			}
		},
	)
}

func TestFileFormatVersions(t *testing.T) {
	t.Run(
		"version string representations",
		func(t *testing.T) {
			tests := []struct {
				version  FileFormatVersions
				expected string
			}{
				{Office2016, "Office2016"},
				{Office2019, "Office2019"},
				{Office2021, "Office2021"},
				{Microsoft365, "Microsoft365"},
			}

			for _, tt := range tests {
				if got := tt.version.String(); got != tt.expected {
					t.Errorf(
						"expected %s, got %s",
						tt.expected,
						got,
					)
				}
			}
		},
	)

	t.Run(
		"version comparison",
		func(t *testing.T) {
			if !Office2019.AtLeast(Office2016) {
				t.Error(
					"Office2019 should be at least Office2016",
				)
			}
			if Office2016.AtLeast(Office2019) {
				t.Error(
					"Office2016 should not be at least Office2019",
				)
			}
		},
	)
}

func TestVersionAvailability(t *testing.T) {
	t.Run(
		"available from introduced version",
		func(t *testing.T) {
			avail := NewVersionAvailability(
				Office2019,
			)

			if avail.IsAvailableIn(Office2016) {
				t.Error(
					"should not be available in Office2016",
				)
			}
			if !avail.IsAvailableIn(Office2019) {
				t.Error(
					"should be available in Office2019",
				)
			}
			if !avail.IsAvailableIn(
				Microsoft365,
			) {
				t.Error(
					"should be available in Microsoft365",
				)
			}
		},
	)

	t.Run(
		"not available after removal",
		func(t *testing.T) {
			avail := NewVersionAvailability(
				Office2016,
			).Removed(Office2021)

			if !avail.IsAvailableIn(Office2016) {
				t.Error(
					"should be available in Office2016",
				)
			}
			if !avail.IsAvailableIn(Office2019) {
				t.Error(
					"should be available in Office2019",
				)
			}
			if avail.IsAvailableIn(Office2021) {
				t.Error(
					"should not be available in Office2021",
				)
			}
		},
	)
}

func TestParticles(t *testing.T) {
	t.Run(
		"element particle matches",
		func(t *testing.T) {
			particle := NewElementParticle(
				"p",
				"http://schemas.openxmlformats.org/wordprocessingml/2006/main",
				1,
				1,
			)

			children := []ElementInfo{
				{
					LocalName:    "p",
					NamespaceURI: "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
				},
			}

			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)
			errors, consumed := particle.Validate(
				ctx,
				children,
				"/",
			)

			if len(errors) != 0 {
				t.Errorf(
					"expected no errors, got %d",
					len(errors),
				)
			}
			if consumed != 1 {
				t.Errorf(
					"expected 1 consumed, got %d",
					consumed,
				)
			}
		},
	)

	t.Run(
		"element particle missing required",
		func(t *testing.T) {
			particle := NewElementParticle(
				"p",
				"http://schemas.openxmlformats.org/wordprocessingml/2006/main",
				1,
				1,
			)

			children := make([]ElementInfo, 0)

			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)
			errors, consumed := particle.Validate(
				ctx,
				children,
				"/",
			)

			if len(errors) != 1 {
				t.Errorf(
					"expected 1 error, got %d",
					len(errors),
				)
			}
			if errors[0].Code != SchemaMissingRequiredElement {
				t.Errorf(
					"expected SchemaMissingRequiredElement, got %s",
					errors[0].Code,
				)
			}
			if consumed != 0 {
				t.Errorf(
					"expected 0 consumed, got %d",
					consumed,
				)
			}
		},
	)

	t.Run(
		"element particle too many",
		func(t *testing.T) {
			particle := NewElementParticle(
				"p",
				"",
				1,
				1,
			)

			children := []ElementInfo{
				{LocalName: "p"},
				{LocalName: "p"},
			}

			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)
			errors, consumed := particle.Validate(
				ctx,
				children,
				"/",
			)

			if len(errors) != 1 {
				t.Errorf(
					"expected 1 error, got %d",
					len(errors),
				)
			}
			if errors[0].Code != SchemaTooManyElements {
				t.Errorf(
					"expected SchemaTooManyElements, got %s",
					errors[0].Code,
				)
			}
			if consumed != 2 {
				t.Errorf(
					"expected 2 consumed, got %d",
					consumed,
				)
			}
		},
	)

	t.Run(
		"sequence particle in order",
		func(t *testing.T) {
			particle := NewSequenceParticle(1, 1,
				NewElementParticle("a", "", 1, 1),
				NewElementParticle("b", "", 1, 1),
				NewElementParticle("c", "", 1, 1),
			)

			children := []ElementInfo{
				{LocalName: "a"},
				{LocalName: "b"},
				{LocalName: "c"},
			}

			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)
			errors, consumed := particle.Validate(
				ctx,
				children,
				"/",
			)

			if len(errors) != 0 {
				t.Errorf(
					"expected no errors, got %d: %v",
					len(errors),
					errors,
				)
			}
			if consumed != 3 {
				t.Errorf(
					"expected 3 consumed, got %d",
					consumed,
				)
			}
		},
	)

	t.Run("choice particle", func(t *testing.T) {
		particle := NewChoiceParticle(1, 1,
			NewElementParticle("a", "", 1, 1),
			NewElementParticle("b", "", 1, 1),
		)

		children := []ElementInfo{
			{LocalName: "b"},
		}

		ctx := NewValidationContext(
			nil,
			Microsoft365,
		)
		errors, consumed := particle.Validate(
			ctx,
			children,
			"/",
		)

		if len(errors) != 0 {
			t.Errorf(
				"expected no errors, got %d",
				len(errors),
			)
		}
		if consumed != 1 {
			t.Errorf(
				"expected 1 consumed, got %d",
				consumed,
			)
		}
	})

	t.Run(
		"all particle any order",
		func(t *testing.T) {
			particle := NewAllParticle(1, 1,
				NewElementParticle("a", "", 1, 1),
				NewElementParticle("b", "", 1, 1),
			)

			// Reversed order should still work
			children := []ElementInfo{
				{LocalName: "b"},
				{LocalName: "a"},
			}

			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)
			errors, consumed := particle.Validate(
				ctx,
				children,
				"/",
			)

			if len(errors) != 0 {
				t.Errorf(
					"expected no errors, got %d: %v",
					len(errors),
					errors,
				)
			}
			if consumed != 2 {
				t.Errorf(
					"expected 2 consumed, got %d",
					consumed,
				)
			}
		},
	)

	t.Run(
		"empty particle no children",
		func(t *testing.T) {
			particle := NewEmptyParticle()

			children := make([]ElementInfo, 0)

			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)
			errors, consumed := particle.Validate(
				ctx,
				children,
				"/",
			)

			if len(errors) != 0 {
				t.Errorf(
					"expected no errors, got %d",
					len(errors),
				)
			}
			if consumed != 0 {
				t.Errorf(
					"expected 0 consumed, got %d",
					consumed,
				)
			}
		},
	)

	t.Run(
		"empty particle with children",
		func(t *testing.T) {
			particle := NewEmptyParticle()

			children := []ElementInfo{
				{LocalName: "unexpected"},
			}

			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)
			errors, _ := particle.Validate(
				ctx,
				children,
				"/",
			)

			if len(errors) != 1 {
				t.Errorf(
					"expected 1 error, got %d",
					len(errors),
				)
			}
			if errors[0].Code != SchemaUnexpectedElement {
				t.Errorf(
					"expected SchemaUnexpectedElement, got %s",
					errors[0].Code,
				)
			}
		},
	)
}

func TestConstraints(t *testing.T) {
	t.Run(
		"unique ID constraint",
		func(t *testing.T) {
			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)

			constraint := NewUniqueIDConstraint(
				func(e any) string {
					if me, ok := e.(*mockElement); ok {
						return me.attributes["id"]
					}

					return ""
				},
			)

			elem1 := newMockElement(
				"p",
			).withAttr("id", "unique1")
			elem2 := newMockElement(
				"p",
			).withAttr("id", "unique1")

			// First element should pass
			errs := constraint.Check(ctx, elem1)
			if len(errs) != 0 {
				t.Errorf(
					"expected no errors for first element, got %d",
					len(errs),
				)
			}

			// Second element with same ID should fail
			errs = constraint.Check(ctx, elem2)
			if len(errs) != 1 {
				t.Errorf(
					"expected 1 error for duplicate, got %d",
					len(errs),
				)
			}
			if len(errs) > 0 &&
				errs[0].Code != SemanticDuplicateID {
				t.Errorf(
					"expected SemanticDuplicateID, got %s",
					errs[0].Code,
				)
			}
		},
	)

	t.Run(
		"mutually exclusive constraint",
		func(t *testing.T) {
			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)

			constraint := NewMutuallyExclusiveConstraint(
				"attrA",
				"attrB",
			)

			// Single attribute is OK
			elem1 := newMockElement(
				"p",
			).withAttr("attrA", "value")
			errs := constraint.Check(ctx, elem1)
			if len(errs) != 0 {
				t.Errorf(
					"expected no errors for single attr, got %d",
					len(errs),
				)
			}

			// Both attributes should fail
			elem2 := newMockElement(
				"p",
			).withAttr("attrA", "value").
				withAttr("attrB", "value")
			errs = constraint.Check(ctx, elem2)
			if len(errs) != 1 {
				t.Errorf(
					"expected 1 error for both attrs, got %d",
					len(errs),
				)
			}
			if len(errs) > 0 &&
				errs[0].Code != SemanticMutuallyExclusive {
				t.Errorf(
					"expected SemanticMutuallyExclusive, got %s",
					errs[0].Code,
				)
			}
		},
	)

	t.Run(
		"register and retrieve constraints",
		func(t *testing.T) {
			ClearConstraints()

			elemType := reflect.TypeOf(
				&mockElement{},
			)
			constraint := NewUniqueIDConstraint(
				func(_ any) string { return "" },
			)

			RegisterConstraint(
				elemType,
				constraint,
			)
			constraints := GetConstraints(
				elemType,
			)

			if len(constraints) != 1 {
				t.Errorf(
					"expected 1 constraint, got %d",
					len(constraints),
				)
			}

			ClearConstraints()
		},
	)
}

func TestSchemaValidator(t *testing.T) {
	t.Run(
		"validates required attributes",
		func(t *testing.T) {
			validator := NewSchemaValidator(nil).
				WithAttribute(&AttributeSchema{
					LocalName: "required",
					Required:  true,
				})

			elem := newMockElement("test")

			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)
			errors := validator.Validate(
				ctx,
				elem,
			)

			if len(errors) != 1 {
				t.Errorf(
					"expected 1 error, got %d",
					len(errors),
				)
			}
			if len(errors) > 0 &&
				errors[0].Code != SchemaMissingRequiredAttribute {
				t.Errorf(
					"expected SchemaMissingRequiredAttribute, got %s",
					errors[0].Code,
				)
			}
		},
	)

	t.Run(
		"validates enumeration values",
		func(t *testing.T) {
			validator := NewSchemaValidator(nil).
				WithAttribute(&AttributeSchema{
					LocalName: "type",
					Values: []string{
						"a",
						"b",
						"c",
					},
				})

			elem := newMockElement(
				"test",
			).withAttr("type", "invalid")

			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)
			errors := validator.Validate(
				ctx,
				elem,
			)

			if len(errors) != 1 {
				t.Errorf(
					"expected 1 error, got %d",
					len(errors),
				)
			}
			if len(errors) > 0 &&
				errors[0].Code != SchemaValueNotInEnumeration {
				t.Errorf(
					"expected SchemaValueNotInEnumeration, got %s",
					errors[0].Code,
				)
			}
		},
	)
}

func TestValidateElement(t *testing.T) {
	t.Run(
		"validates element tree",
		func(t *testing.T) {
			root := newMockElement("document").
				withChild(newMockElement("body").
					withChild(newMockElement("paragraph")),
				)

			ctx := NewValidationContext(
				nil,
				Microsoft365,
			)
			errors := ValidateElement(root, ctx)

			// Without registered validators, should have no errors
			if len(errors) != 0 {
				t.Errorf(
					"expected 0 errors, got %d: %v",
					len(errors),
					errors,
				)
			}
		},
	)
}

func TestBuildXPath(t *testing.T) {
	t.Run(
		"builds path from element chain",
		func(t *testing.T) {
			doc := newMockElement(
				"document",
			).withNS("http://schemas.openxmlformats.org/wordprocessingml/2006/main")
			body := newMockElement(
				"body",
			).withNS("http://schemas.openxmlformats.org/wordprocessingml/2006/main")
			para := newMockElement(
				"p",
			).withNS("http://schemas.openxmlformats.org/wordprocessingml/2006/main")

			doc.withChild(body)
			body.withChild(para)

			path := BuildXPath(para)
			expected := "/w:document/w:body/w:p"
			if path != expected {
				t.Errorf(
					"expected %s, got %s",
					expected,
					path,
				)
			}
		},
	)
}
