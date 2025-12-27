package types

import (
	"fmt"
	"strconv"
	"strings"
)

// floatBitSize is the bit size for parsing float64 values.
// Already defined in decimal_value.go

// DoubleValue wraps a float64 value for XML double attributes.
// It implements the SimpleValue interface.
type DoubleValue struct {
	value    float64
	hasValue bool
}

// NewDoubleValue creates a new DoubleValue with the given float64.
func NewDoubleValue(v float64) *DoubleValue {
	return &DoubleValue{
		value:    v,
		hasValue: true,
	}
}

// NewNilDoubleValue creates a new DoubleValue in the unset/nil state.
func NewNilDoubleValue() *DoubleValue {
	return &DoubleValue{
		hasValue: false,
	}
}

// Value returns the float64 value.
// Returns 0 if the value is not set.
func (dv *DoubleValue) Value() float64 {
	if !dv.hasValue {
		return 0
	}

	return dv.value
}

// SetValue sets the float64 value.
func (dv *DoubleValue) SetValue(v float64) {
	dv.value = v
	dv.hasValue = true
}

// HasValue returns true if the value is set.
func (dv *DoubleValue) HasValue() bool {
	return dv.hasValue
}

// InnerText returns the string representation for XML serialization.
// Uses minimal precision to represent the value accurately.
func (dv *DoubleValue) InnerText() string {
	if !dv.hasValue {
		return ""
	}
	// Use -1 precision to get minimal representation
	s := strconv.FormatFloat(
		dv.value,
		'f',
		-1,
		floatBitSize,
	)
	// Ensure there's no trailing zeros after decimal point for cleaner output
	// But preserve at least one decimal place if there is one
	if strings.Contains(s, ".") {
		s = strings.TrimRight(s, "0")
		s = strings.TrimRight(s, ".")
	}

	return s
}

// SetInnerText parses the value from a double string.
// Returns an error if the string cannot be parsed as a float64.
func (dv *DoubleValue) SetInnerText(
	text string,
) error {
	if text == "" {
		dv.hasValue = false
		dv.value = 0

		return nil
	}
	v, err := strconv.ParseFloat(
		strings.TrimSpace(text),
		floatBitSize,
	)
	if err != nil {
		return fmt.Errorf(
			"invalid double value: %w",
			err,
		)
	}
	dv.value = v
	dv.hasValue = true

	return nil
}

// SetNil clears the value, making HasValue() return false.
func (dv *DoubleValue) SetNil() {
	dv.value = 0
	dv.hasValue = false
}

// Ensure DoubleValue implements SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*DoubleValue)(nil)
	_ Resettable  = (*DoubleValue)(nil)
)
