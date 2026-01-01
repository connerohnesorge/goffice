package validation

import (
	"fmt"
)

// MarkupCompatibilityNamespace is the namespace URI for AlternateContent elements.
const MarkupCompatibilityNamespace = "http://schemas.openxmlformats.org/markup-compatibility/2006"

// AlternateContentValidator validates AlternateContent elements and their
// Choice/Fallback structure for proper version compatibility patterns.
type AlternateContentValidator struct {
	// NamespaceURI is the namespace of AlternateContent elements.
	NamespaceURI string
}

// NewAlternateContentValidator creates a new AlternateContent validator.
func NewAlternateContentValidator() *AlternateContentValidator {
	return &AlternateContentValidator{
		NamespaceURI: MarkupCompatibilityNamespace,
	}
}

// Validate validates an AlternateContent element.
// It checks that:
// - At least one Choice or Fallback element is present
// - All Choice elements have a valid Requires attribute
// - The structure is well-formed (Choices before Fallback)
func (v *AlternateContentValidator) Validate(
	ctx *ValidationContext,
	element any,
) []*ValidationError {
	var errors []*ValidationError

	// Check if this is an AlternateContent element
	if !v.isAlternateContent(element) {
		return errors
	}

	// Get children
	children := getChildElementsForValidator(element)
	if len(children) == 0 {
		errors = append(
			errors,
			NewValidationError(
				SchemaUnexpectedElement,
				"AlternateContent must have at least one Choice or Fallback element",
				ctx.CurrentPath(),
				element,
			),
		)

		return errors
	}

	// Track Fallback position
	var fallbackIndex = -1
	var choiceCount int

	// Validate structure: all Choices must come before Fallback
	for i, child := range children {
		childName := getElementLocalName(child)

		switch {
		case v.isChoice(child):
			choiceCount++
			if fallbackIndex >= 0 {
				errors = append(
					errors,
					NewValidationError(
						SchemaInvalidChildOrder,
						fmt.Sprintf(
							"Choice element at position %d cannot follow Fallback element at position %d",
							i,
							fallbackIndex,
						),
						ctx.CurrentPath(),
						element,
					),
				)
			}

			// Validate Choice element
			choiceErrors := v.validateChoice(
				ctx,
				child,
				i,
			)
			errors = append(errors, choiceErrors...)
		case v.isFallback(child):
			if fallbackIndex >= 0 {
				errors = append(
					errors,
					NewValidationError(
						SchemaTooManyElements,
						fmt.Sprintf(
							"Multiple Fallback elements found (positions %d and %d)",
							fallbackIndex,
							i,
						),
						ctx.CurrentPath(),
						element,
					),
				)
			}
			fallbackIndex = i
		case childName != "":
			// Unexpected element
			errors = append(
				errors,
				NewValidationError(
					SchemaUnexpectedElement,
					fmt.Sprintf(
						"Unexpected element '%s' in AlternateContent, only Choice and Fallback are allowed",
						childName,
					),
					ctx.CurrentPath(),
					element,
				),
			)
		}
	}

	// At least one Choice is required
	if choiceCount == 0 {
		errors = append(
			errors,
			NewValidationError(
				SchemaMissingRequiredElement,
				"AlternateContent must have at least one Choice element",
				ctx.CurrentPath(),
				element,
			),
		)
	}

	return errors
}

// validateChoice validates a single Choice element.
// Checks that it has a valid Requires attribute.
func (v *AlternateContentValidator) validateChoice(
	ctx *ValidationContext,
	choice any,
	index int,
) []*ValidationError {
	var errors []*ValidationError

	// Get Requires attribute
	requires := getAttributeValue(choice, "Requires")
	if requires == "" {
		errors = append(
			errors,
			NewValidationError(
				SchemaMissingRequiredAttribute,
				"Choice element is missing required 'Requires' attribute",
				fmt.Sprintf("%s/Choice[%d]", ctx.CurrentPath(), index),
				choice,
			),
		)

		return errors
	}

	// Validate Requires is a valid namespace prefix
	// Common valid prefixes: w14, w15, x14, x15, a14, p14, etc.
	if !v.isValidRequiresValue(requires) {
		errors = append(
			errors,
			NewValidationError(
				SchemaInvalidAttributeValue,
				fmt.Sprintf(
					"Choice 'Requires' attribute has invalid value '%s'. Expected a namespace prefix like 'w14', 'x15', etc.",
					requires,
				),
				fmt.Sprintf("%s/Choice[%d]", ctx.CurrentPath(), index),
				choice,
			),
		)
	}

	return errors
}

// isValidRequiresValue checks if a value is a valid namespace prefix for Choice/Requires.
// Valid prefixes follow the pattern [a-z][0-9]+ (e.g., "w14", "x15", "a14").
func (*AlternateContentValidator) isValidRequiresValue(value string) bool {
	if len(value) < 2 {
		return false
	}

	// First character should be a letter
	if value[0] < 'a' || value[0] > 'z' {
		return false
	}

	// Remaining characters should be digits
	for i := 1; i < len(value); i++ {
		if value[i] < '0' || value[i] > '9' {
			return false
		}
	}

	return true
}

// Helper methods

func (*AlternateContentValidator) isAlternateContent(element any) bool {
	name := getElementLocalName(element)
	ns := getElementNamespaceURI(element)

	return name == "AlternateContent" &&
		ns == MarkupCompatibilityNamespace
}

func (*AlternateContentValidator) isChoice(element any) bool {
	name := getElementLocalName(element)
	ns := getElementNamespaceURI(element)

	return name == "Choice" &&
		ns == MarkupCompatibilityNamespace
}

func (*AlternateContentValidator) isFallback(element any) bool {
	name := getElementLocalName(element)
	ns := getElementNamespaceURI(element)

	return name == "Fallback" &&
		ns == MarkupCompatibilityNamespace
}

// getElementLocalName extracts the local name from an element.
// This is a helper that works with any element type.
func getElementLocalName(element any) string {
	// Try getting from IElementMetadata feature first
	if localNamer, ok := element.(interface{ LocalName() string }); ok {
		return localNamer.LocalName()
	}

	// Try reflection as fallback
	if info := GetElementInfoFromInterface(element); info != nil {
		return info.LocalName
	}

	return ""
}

// getElementNamespaceURI extracts the namespace URI from an element.
func getElementNamespaceURI(element any) string {
	// Try getting from IElementMetadata feature first
	if nsUrier, ok := element.(interface{ NamespaceURI() string }); ok {
		return nsUrier.NamespaceURI()
	}

	// Try reflection as fallback
	if info := GetElementInfoFromInterface(element); info != nil {
		return info.NamespaceURI
	}

	return ""
}

// getAttributeValue extracts an attribute value from an element.
func getAttributeValue(element any, attrName string) string {
	attrs := GetElementAttributes(element)
	if value, ok := attrs[attrName]; ok {
		return value
	}

	return ""
}

// getChildElementsForValidator gets child elements for validation.
// Similar to GetChildElementInfos but returns raw element interfaces.
func getChildElementsForValidator(element any) []any {
	if composite, ok := element.(interface{ Children() []any }); ok {
		children := composite.Children()
		result := make([]any, 0, len(children))
		for _, child := range children {
			// Filter out text nodes and other non-element nodes
			if child != nil &&
				getElementLocalName(child) != "" {
				result = append(result, child)
			}
		}

		return result
	}

	return nil
}
