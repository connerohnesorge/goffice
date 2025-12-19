// Package validation provides the validation framework for Office Open XML documents.
// It supports schema validation, semantic validation, and custom validation rules
// to ensure documents conform to the OOXML specification.
package validation

import (
	"fmt"
	"strings"
)

// ValidationSeverity indicates the severity of a validation error.
type ValidationSeverity int

const (
	// SeverityError indicates a validation error that makes the document non-conformant.
	SeverityError ValidationSeverity = iota
	// SeverityWarning indicates a validation warning that may cause issues but is not strictly non-conformant.
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
	// Schema_MissingRequiredElement indicates a required child element is missing.
	Schema_MissingRequiredElement ErrorCode = "Schema_MissingRequiredElement"
	// Schema_UnexpectedElement indicates an unexpected element was found.
	Schema_UnexpectedElement ErrorCode = "Schema_UnexpectedElement"
	// Schema_InvalidChildOrder indicates child elements are in the wrong order.
	Schema_InvalidChildOrder ErrorCode = "Schema_InvalidChildOrder"
	// Schema_TooManyElements indicates too many occurrences of an element.
	Schema_TooManyElements ErrorCode = "Schema_TooManyElements"
	// Schema_MissingRequiredAttribute indicates a required attribute is missing.
	Schema_MissingRequiredAttribute ErrorCode = "Schema_MissingRequiredAttribute"
	// Schema_InvalidAttributeValue indicates an attribute has an invalid value.
	Schema_InvalidAttributeValue ErrorCode = "Schema_InvalidAttributeValue"
	// Schema_ValueNotInEnumeration indicates a value is not in the allowed enumeration.
	Schema_ValueNotInEnumeration ErrorCode = "Schema_ValueNotInEnumeration"
	// Schema_ValueOutOfRange indicates a value is outside the allowed range.
	Schema_ValueOutOfRange ErrorCode = "Schema_ValueOutOfRange"
	// Schema_ElementNotAvailable indicates an element is not available in the target version.
	Schema_ElementNotAvailable ErrorCode = "Schema_ElementNotAvailable"
	// Schema_AttributeNotAvailable indicates an attribute is not available in the target version.
	Schema_AttributeNotAvailable ErrorCode = "Schema_AttributeNotAvailable"
)

// Semantic validation error codes.
const (
	// Semantic_DuplicateID indicates duplicate unique identifiers were found.
	Semantic_DuplicateID ErrorCode = "Semantic_DuplicateID"
	// Semantic_RelationshipNotFound indicates a referenced relationship was not found.
	Semantic_RelationshipNotFound ErrorCode = "Semantic_RelationshipNotFound"
	// Semantic_InvalidParentType indicates an element has an invalid parent type.
	Semantic_InvalidParentType ErrorCode = "Semantic_InvalidParentType"
	// Semantic_MutuallyExclusiveAttributes indicates mutually exclusive attributes are present.
	Semantic_MutuallyExclusiveAttributes ErrorCode = "Semantic_MutuallyExclusiveAttributes"
	// Semantic_InvalidReference indicates an invalid reference was found.
	Semantic_InvalidReference ErrorCode = "Semantic_InvalidReference"
)

// ValidationError represents a single validation error or warning.
type ValidationError struct {
	// Code is the machine-readable error code.
	Code ErrorCode
	// Description is the human-readable error message.
	Description string
	// Path is the XPath-like path to the element with the error.
	Path string
	// Element is a reference to the actual element with the error (optional).
	// This uses interface{} to avoid circular imports with the openxml package.
	Element interface{}
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
	element interface{},
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
	element interface{},
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

// ValidationErrors is a collection of validation errors that implements the error interface.
type ValidationErrors []*ValidationError

// Error implements the error interface for the collection.
func (errs ValidationErrors) Error() string {
	if len(errs) == 0 {
		return "no validation errors"
	}
	if len(errs) == 1 {
		return errs[0].Error()
	}
	var sb strings.Builder
	sb.WriteString(
		fmt.Sprintf(
			"%d validation errors:\n",
			len(errs),
		),
	)
	for i, err := range errs {
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

// HasErrors returns true if the collection contains any errors (not just warnings).
func (errs ValidationErrors) HasErrors() bool {
	for _, e := range errs {
		if e.Severity == SeverityError {
			return true
		}
	}
	return false
}

// Errors returns only the errors (excluding warnings).
func (errs ValidationErrors) Errors() ValidationErrors {
	result := make(ValidationErrors, 0, len(errs))
	for _, e := range errs {
		if e.Severity == SeverityError {
			result = append(result, e)
		}
	}
	return result
}

// Warnings returns only the warnings (excluding errors).
func (errs ValidationErrors) Warnings() ValidationErrors {
	result := make(ValidationErrors, 0, len(errs))
	for _, e := range errs {
		if e.Severity == SeverityWarning {
			result = append(result, e)
		}
	}
	return result
}

// ByCode returns errors with the given error code.
func (errs ValidationErrors) ByCode(
	code ErrorCode,
) ValidationErrors {
	result := make(ValidationErrors, 0)
	for _, e := range errs {
		if e.Code == code {
			result = append(result, e)
		}
	}
	return result
}
