package types

import (
	"fmt"
	"strconv"
)

// Int64Value wraps a 64-bit signed integer value with optional nil/unset state.
// It implements the SimpleValue interface.
type Int64Value struct {
	value    int64
	hasValue bool
}

// NewInt64Value creates a new Int64Value with the given integer.
func NewInt64Value(v int64) *Int64Value {
	return &Int64Value{
		value:    v,
		hasValue: true,
	}
}

// NewNilInt64Value creates a new Int64Value in the unset/nil state.
func NewNilInt64Value() *Int64Value {
	return &Int64Value{
		hasValue: false,
	}
}

// Value returns the int64 value.
// Returns 0 if the value is not set.
func (iv *Int64Value) Value() int64 {
	if !iv.hasValue {
		return 0
	}
	return iv.value
}

// SetValue sets the int64 value.
func (iv *Int64Value) SetValue(v int64) {
	iv.value = v
	iv.hasValue = true
}

// HasValue returns true if the value is set.
func (iv *Int64Value) HasValue() bool {
	return iv.hasValue
}

// InnerText returns the string representation for XML serialization.
func (iv *Int64Value) InnerText() string {
	if !iv.hasValue {
		return ""
	}
	return strconv.FormatInt(iv.value, 10)
}

// SetInnerText parses the value from a string.
// Returns an error if the string cannot be parsed as an int64.
func (iv *Int64Value) SetInnerText(
	text string,
) error {
	if text == "" {
		iv.hasValue = false
		iv.value = 0
		return nil
	}
	v, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return fmt.Errorf(
			"invalid int64 value: %w",
			err,
		)
	}
	iv.value = v
	iv.hasValue = true
	return nil
}

// SetNil clears the value, making HasValue() return false.
func (iv *Int64Value) SetNil() {
	iv.value = 0
	iv.hasValue = false
}

// UInt64Value wraps a 64-bit unsigned integer value with optional nil/unset state.
// It implements the SimpleValue interface.
type UInt64Value struct {
	value    uint64
	hasValue bool
}

// NewUInt64Value creates a new UInt64Value with the given unsigned integer.
func NewUInt64Value(v uint64) *UInt64Value {
	return &UInt64Value{
		value:    v,
		hasValue: true,
	}
}

// NewNilUInt64Value creates a new UInt64Value in the unset/nil state.
func NewNilUInt64Value() *UInt64Value {
	return &UInt64Value{
		hasValue: false,
	}
}

// Value returns the uint64 value.
// Returns 0 if the value is not set.
func (uv *UInt64Value) Value() uint64 {
	if !uv.hasValue {
		return 0
	}
	return uv.value
}

// SetValue sets the uint64 value.
func (uv *UInt64Value) SetValue(v uint64) {
	uv.value = v
	uv.hasValue = true
}

// HasValue returns true if the value is set.
func (uv *UInt64Value) HasValue() bool {
	return uv.hasValue
}

// InnerText returns the string representation for XML serialization.
func (uv *UInt64Value) InnerText() string {
	if !uv.hasValue {
		return ""
	}
	return strconv.FormatUint(uv.value, 10)
}

// SetInnerText parses the value from a string.
// Returns an error if the string cannot be parsed as a uint64 or is negative.
func (uv *UInt64Value) SetInnerText(
	text string,
) error {
	if text == "" {
		uv.hasValue = false
		uv.value = 0
		return nil
	}
	// Check for negative values
	if len(text) > 0 && text[0] == '-' {
		return fmt.Errorf(
			"invalid uint64 value: negative values not allowed: %s",
			text,
		)
	}
	v, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return fmt.Errorf(
			"invalid uint64 value: %w",
			err,
		)
	}
	uv.value = v
	uv.hasValue = true
	return nil
}

// SetNil clears the value, making HasValue() return false.
func (uv *UInt64Value) SetNil() {
	uv.value = 0
	uv.hasValue = false
}

// Ensure Int64Value and UInt64Value implement SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*Int64Value)(nil)
	_ Resettable  = (*Int64Value)(nil)
	_ SimpleValue = (*UInt64Value)(nil)
	_ Resettable  = (*UInt64Value)(nil)
)
