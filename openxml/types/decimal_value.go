package types

import (
	"fmt"
	"strconv"
	"strings"
)

// DecimalValue wraps a float64 value for XML decimal attributes.
// It implements the SimpleValue interface.
type DecimalValue struct {
	value    float64
	hasValue bool
}

// NewDecimalValue creates a new DecimalValue with the given float64.
func NewDecimalValue(v float64) *DecimalValue {
	return &DecimalValue{
		value:    v,
		hasValue: true,
	}
}

// NewNilDecimalValue creates a new DecimalValue in the unset/nil state.
func NewNilDecimalValue() *DecimalValue {
	return &DecimalValue{
		hasValue: false,
	}
}

// Value returns the float64 value.
// Returns 0 if the value is not set.
func (dv *DecimalValue) Value() float64 {
	if !dv.hasValue {
		return 0
	}
	return dv.value
}

// SetValue sets the float64 value.
func (dv *DecimalValue) SetValue(v float64) {
	dv.value = v
	dv.hasValue = true
}

// HasValue returns true if the value is set.
func (dv *DecimalValue) HasValue() bool {
	return dv.hasValue
}

// InnerText returns the string representation for XML serialization.
// Uses minimal precision to represent the value accurately.
func (dv *DecimalValue) InnerText() string {
	if !dv.hasValue {
		return ""
	}
	// Use -1 precision to get minimal representation
	s := strconv.FormatFloat(dv.value, 'f', -1, 64)
	// Ensure there's no trailing zeros after decimal point for cleaner output
	// But preserve at least one decimal place if there is one
	if strings.Contains(s, ".") {
		s = strings.TrimRight(s, "0")
		s = strings.TrimRight(s, ".")
	}
	return s
}

// SetInnerText parses the value from a decimal string.
// Returns an error if the string cannot be parsed as a float64.
func (dv *DecimalValue) SetInnerText(text string) error {
	if text == "" {
		dv.hasValue = false
		dv.value = 0
		return nil
	}
	v, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
	if err != nil {
		return fmt.Errorf("invalid decimal value: %w", err)
	}
	dv.value = v
	dv.hasValue = true
	return nil
}

// SetNil clears the value, making HasValue() return false.
func (dv *DecimalValue) SetNil() {
	dv.value = 0
	dv.hasValue = false
}

// Ensure DecimalValue implements SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*DecimalValue)(nil)
	_ Resettable  = (*DecimalValue)(nil)
)
