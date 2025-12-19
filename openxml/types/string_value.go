package types

// StringValue wraps a string value with optional nil/unset state.
// It implements the SimpleValue interface.
type StringValue struct {
	value    string
	hasValue bool
}

// NewStringValue creates a new StringValue with the given string.
func NewStringValue(s string) *StringValue {
	return &StringValue{
		value:    s,
		hasValue: true,
	}
}

// NewNilStringValue creates a new StringValue in the unset/nil state.
func NewNilStringValue() *StringValue {
	return &StringValue{
		hasValue: false,
	}
}

// Value returns the string value.
// Returns an empty string if the value is not set.
func (sv *StringValue) Value() string {
	if !sv.hasValue {
		return ""
	}

	return sv.value
}

// SetValue sets the string value.
func (sv *StringValue) SetValue(s string) {
	sv.value = s
	sv.hasValue = true
}

// HasValue returns true if the value is set.
func (sv *StringValue) HasValue() bool {
	return sv.hasValue
}

// InnerText returns the string representation for XML serialization.
func (sv *StringValue) InnerText() string {
	if !sv.hasValue {
		return ""
	}

	return sv.value
}

// SetInnerText parses the value from a string.
// An empty string sets the value to empty (still considered "set").
// Use SetNil() to explicitly unset the value.
func (sv *StringValue) SetInnerText(
	text string,
) error {
	sv.value = text
	sv.hasValue = true

	return nil
}

// SetNil clears the value, making HasValue() return false.
func (sv *StringValue) SetNil() {
	sv.value = ""
	sv.hasValue = false
}

// Ensure StringValue implements SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*StringValue)(nil)
	_ Resettable  = (*StringValue)(nil)
)
