package types

import "fmt"

// EnumStringer is the interface that enum types must implement
// to be usable with EnumValue.
type EnumStringer interface {
	~string
}

// EnumParser is an interface for types that can parse enum values from strings.
// This allows enum types to define their own parsing logic.
type EnumParser[T any] interface {
	// ParseEnum parses a string value into the enum type.
	// Returns the parsed value and true if successful, or zero value and false if not.
	ParseEnum(s string) (T, bool)
}

// EnumValue is a generic type for enumeration attributes.
// T must be a string-based type (typically a defined enum type).
type EnumValue[T EnumStringer] struct {
	value       T
	hasValue    bool
	validValues map[string]T // Optional mapping of valid string values to enum values
}

// NewEnumValue creates a new EnumValue with the given enum value.
func NewEnumValue[T EnumStringer](
	v T,
) *EnumValue[T] {
	return &EnumValue[T]{
		value:    v,
		hasValue: true,
	}
}

// NewEnumValueWithValidation creates a new EnumValue with validation support.
// The validValues map defines the set of valid string representations and their enum values.
func NewEnumValueWithValidation[T EnumStringer](
	v T,
	validValues map[string]T,
) *EnumValue[T] {
	return &EnumValue[T]{
		value:       v,
		hasValue:    true,
		validValues: validValues,
	}
}

// NewNilEnumValue creates a new EnumValue in the unset/nil state.
func NewNilEnumValue[T EnumStringer]() *EnumValue[T] {
	return &EnumValue[T]{
		hasValue: false,
	}
}

// NewNilEnumValueWithValidation creates a new EnumValue in the unset/nil state with validation support.
func NewNilEnumValueWithValidation[T EnumStringer](
	validValues map[string]T,
) *EnumValue[T] {
	return &EnumValue[T]{
		hasValue:    false,
		validValues: validValues,
	}
}

// Value returns the enum value.
// Returns the zero value if the value is not set.
func (ev *EnumValue[T]) Value() T {
	if !ev.hasValue {
		var zero T

		return zero
	}

	return ev.value
}

// SetValue sets the enum value.
func (ev *EnumValue[T]) SetValue(v T) {
	ev.value = v
	ev.hasValue = true
}

// HasValue returns true if the value is set.
func (ev *EnumValue[T]) HasValue() bool {
	return ev.hasValue
}

// SetValidValues sets the mapping of valid string values.
// This is used for validation during SetInnerText.
func (ev *EnumValue[T]) SetValidValues(
	validValues map[string]T,
) {
	ev.validValues = validValues
}

// InnerText returns the string representation for XML serialization.
func (ev *EnumValue[T]) InnerText() string {
	if !ev.hasValue {
		return ""
	}

	return string(ev.value)
}

// SetInnerText parses the value from a string.
// If validValues is set, validates that the string is a valid enum value.
// Returns an error if validation is enabled and the string is not valid.
func (ev *EnumValue[T]) SetInnerText(
	text string,
) error {
	if text == "" {
		ev.hasValue = false
		var zero T
		ev.value = zero

		return nil
	}

	// If we have a validation map, use it
	if ev.validValues != nil {
		if v, ok := ev.validValues[text]; ok {
			ev.value = v
			ev.hasValue = true

			return nil
		}

		return fmt.Errorf(
			"invalid enum value: %q",
			text,
		)
	}

	// Without validation, accept any string and convert it
	ev.value = T(text)
	ev.hasValue = true

	return nil
}

// SetNil clears the value, making HasValue() return false.
func (ev *EnumValue[T]) SetNil() {
	var zero T
	ev.value = zero
	ev.hasValue = false
}

// Ensure EnumValue implements SimpleValue and Resettable interfaces.
// Note: We can't use the interface check pattern with generics directly,
// but the methods are implemented correctly.
