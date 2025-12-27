// Package validation provides the validation framework for Office Open XML
// documents. It supports schema validation, semantic validation, and custom
// validation rules to ensure documents conform to the OOXML specification.
package validation

import (
	"fmt"
	"strings"
)

// ValidationSeverity indicates the severity of a validation error.
type ValidationSeverity int

const (
	// SeverityError indicates a validation error that makes the document
	// non-conformant.
	SeverityError ValidationSeverity = iota
	// SeverityWarning indicates a validation warning that may cause issues
	// but is not strictly non-conformant.
	SeverityWarning
)

// String returns the string representation of the severity.
func (s ValidationSeverity) String() string {
	switch s {
	case SeverityError:
		return "Error"
	case SeverityWarning:
		return "Warning"
	default:
		return "Unknown"
	}
}

// ErrorCode represents a machine-readable error code for validation errors.
type ErrorCode string

// Schema validation error codes.
const (
	// SchemaMissingRequiredElement indicates a required child element
	// is missing.
	SchemaMissingRequiredElement ErrorCode = "SchemaMissingRequiredElement"
	// SchemaUnexpectedElement indicates an unexpected element was found.
	SchemaUnexpectedElement ErrorCode = "SchemaUnexpectedElement"
	// SchemaInvalidChildOrder indicates child elements are in the wrong order.
	SchemaInvalidChildOrder ErrorCode = "SchemaInvalidChildOrder"
	// SchemaTooManyElements indicates too many occurrences of an element.
	SchemaTooManyElements ErrorCode = "SchemaTooManyElements"
	// SchemaMissingRequiredAttribute indicates a required attribute is missing.
	SchemaMissingRequiredAttribute ErrorCode = "SchemaMissingRequiredAttribute"
	// SchemaInvalidAttributeValue indicates an attribute has an invalid value.
	SchemaInvalidAttributeValue ErrorCode = "SchemaInvalidAttributeValue"
	// SchemaValueNotInEnumeration indicates a value is not in the allowed
	// enumeration.
	SchemaValueNotInEnumeration ErrorCode = "SchemaValueNotInEnumeration"
	// SchemaValueOutOfRange indicates a value is outside the allowed range.
	SchemaValueOutOfRange ErrorCode = "SchemaValueOutOfRange"
	// SchemaElementNotAvailable indicates an element is not available
	// in the target version.
	SchemaElementNotAvailable ErrorCode = "SchemaElementNotAvailable"
	// SchemaAttributeNotAvailable indicates an attribute is not available
	// in the target version.
	SchemaAttributeNotAvailable ErrorCode = "SchemaAttributeNotAvailable"
)

// Semantic validation error codes.
const (
	// SemanticDuplicateID indicates duplicate unique identifiers were found.
	SemanticDuplicateID ErrorCode = "SemanticDuplicateID"
	// SemanticRelationshipNotFound indicates a referenced relationship
	// was not found.
	SemanticRelationshipNotFound ErrorCode = "SemanticRelationshipNotFound"
	// SemanticInvalidParentType indicates an element has an invalid
	// parent type.
	SemanticInvalidParentType ErrorCode = "SemanticInvalidParentType"
	// SemanticMutuallyExclusive indicates mutually exclusive attributes
	// are present.
	SemanticMutuallyExclusive ErrorCode = "SemanticMutuallyExclusive"
	// SemanticInvalidReference indicates an invalid reference was found.
	SemanticInvalidReference ErrorCode = "SemanticInvalidReference"
)

// ValidationError represents a single validation error or warning.
type ValidationError struct {
	// Code is the machine-readable error code.
	Code ErrorCode
	// Description is the human-readable error message.
	Description string
	// Path is the XPath-like path to the element with the error.
	Path string
	// Element is a reference to the actual element with the error
	// (optional).
	// This uses any to avoid circular imports with the openxml package.
	Element any
	// Severity indicates whether this is an error or warning.
	Severity ValidationSeverity
	// RelatedInfo provides additional context about the error.
	RelatedInfo string
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	var sb strings.Builder
	sb.WriteString(e.Severity.String())
	sb.WriteString(": ")
	sb.WriteString(string(e.Code))
	sb.WriteString(" - ")
	sb.WriteString(e.Description)
	if e.Path != "" {
		sb.WriteString(" at ")
		sb.WriteString(e.Path)
	}

	return sb.String()
}

// NewValidationError creates a new validation error.
func NewValidationError(
	code ErrorCode,
	description, path string,
	element any,
) *ValidationError {
	return &ValidationError{
		Code:        code,
		Description: description,
		Path:        path,
		Element:     element,
		Severity:    SeverityError,
	}
}

// NewValidationWarning creates a new validation warning.
func NewValidationWarning(
	code ErrorCode,
	description, path string,
	element any,
) *ValidationError {
	return &ValidationError{
		Code:        code,
		Description: description,
		Path:        path,
		Element:     element,
		Severity:    SeverityWarning,
	}
}

// WithRelatedInfo adds related information to the error.
func (e *ValidationError) WithRelatedInfo(
	info string,
) *ValidationError {
	e.RelatedInfo = info

	return e
}

// ValidationErrors is a collection of validation errors that implements
// the error interface.
type ValidationErrors []*ValidationError

// Error implements the error interface for the collection.
func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return "no validation errors"
	}
	if len(ve) == 1 {
		return ve[0].Error()
	}
	var sb strings.Builder
	sb.WriteString(
		fmt.Sprintf(
			"%d validation errors:\n",
			len(ve),
		),
	)
	for i, err := range ve {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(
			fmt.Sprintf(
				"  %d. %s",
				i+1,
				err.Error(),
			),
		)
	}

	return sb.String()
}

// HasErrors returns true if the collection contains any errors
// (not just warnings).
func (ve ValidationErrors) HasErrors() bool {
	for _, e := range ve {
		if e.Severity == SeverityError {
			return true
		}
	}

	return false
}

// Errors returns only the errors (excluding warnings).
func (ve ValidationErrors) Errors() ValidationErrors {
	result := make(ValidationErrors, 0, len(ve))
	for _, e := range ve {
		if e.Severity == SeverityError {
			result = append(result, e)
		}
	}

	return result
}

// Warnings returns only the warnings (excluding errors).
func (ve ValidationErrors) Warnings() ValidationErrors {
	result := make(ValidationErrors, 0, len(ve))
	for _, e := range ve {
		if e.Severity == SeverityWarning {
			result = append(result, e)
		}
	}

	return result
}

// ByCode returns errors with the given error code.
func (ve ValidationErrors) ByCode(
	code ErrorCode,
) ValidationErrors {
	result := make(ValidationErrors, 0)
	for _, e := range ve {
		if e.Code == code {
			result = append(result, e)
		}
	}

	return result
}
