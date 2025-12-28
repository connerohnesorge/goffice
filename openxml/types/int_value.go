package types

import (
	"fmt"
	"strconv"
)

// intParseBase is the base used for parsing integer strings.
const intParseBase = 10

// int32BitSize is the bit size for parsing int32 values.
const int32BitSize = 32

// Int32Value wraps a 32-bit signed integer value with optional nil/unset state.
// It implements the SimpleValue interface.
type Int32Value struct {
	value    int32
	hasValue bool
}

// NewInt32Value creates a new Int32Value with the given integer.
func NewInt32Value(v int32) *Int32Value {
	return &Int32Value{
		value:    v,
		hasValue: true,
	}
}

// NewNilInt32Value creates a new Int32Value in the unset/nil state.
func NewNilInt32Value() *Int32Value {
	return &Int32Value{
		hasValue: false,
	}
}

// Value returns the int32 value.
// Returns 0 if the value is not set.
func (iv *Int32Value) Value() int32 {
	if !iv.hasValue {
		return 0
	}

	return iv.value
}

// SetValue sets the int32 value.
func (iv *Int32Value) SetValue(v int32) {
	iv.value = v
	iv.hasValue = true
}

// HasValue returns true if the value is set.
func (iv *Int32Value) HasValue() bool {
	return iv.hasValue
}

// InnerText returns the string representation for XML serialization.
func (iv *Int32Value) InnerText() string {
	if !iv.hasValue {
		return ""
	}

	return strconv.FormatInt(
		int64(iv.value),
		intParseBase,
	)
}

// SetInnerText parses the value from a string.
// Returns an error if the string cannot be parsed as an int32.
func (iv *Int32Value) SetInnerText(
	text string,
) error {
	if text == "" {
		iv.hasValue = false
		iv.value = 0

		return nil
	}
	v, err := strconv.ParseInt(
		text,
		intParseBase,
		int32BitSize,
	)
	if err != nil {
		return fmt.Errorf(
			"invalid int32 value: %w",
			err,
		)
	}
	iv.value = int32(v)
	iv.hasValue = true

	return nil
}

// SetNil clears the value, making HasValue() return false.
func (iv *Int32Value) SetNil() {
	iv.value = 0
	iv.hasValue = false
}

// UInt32Value wraps a 32-bit unsigned integer value with optional nil/unset
// state. It implements the SimpleValue interface.
type UInt32Value struct {
	value    uint32
	hasValue bool
}

// NewUInt32Value creates a new UInt32Value with the given unsigned integer.
func NewUInt32Value(v uint32) *UInt32Value {
	return &UInt32Value{
		value:    v,
		hasValue: true,
	}
}

// NewNilUInt32Value creates a new UInt32Value in the unset/nil state.
func NewNilUInt32Value() *UInt32Value {
	return &UInt32Value{
		hasValue: false,
	}
}

// Value returns the uint32 value.
// Returns 0 if the value is not set.
func (uv *UInt32Value) Value() uint32 {
	if !uv.hasValue {
		return 0
	}

	return uv.value
}

// SetValue sets the uint32 value.
func (uv *UInt32Value) SetValue(v uint32) {
	uv.value = v
	uv.hasValue = true
}

// HasValue returns true if the value is set.
func (uv *UInt32Value) HasValue() bool {
	return uv.hasValue
}

// InnerText returns the string representation for XML serialization.
func (uv *UInt32Value) InnerText() string {
	if !uv.hasValue {
		return ""
	}

	return strconv.FormatUint(
		uint64(uv.value),
		intParseBase,
	)
}

// SetInnerText parses the value from a string.
// Returns an error if the string cannot be parsed as a uint32 or is negative.
func (uv *UInt32Value) SetInnerText(
	text string,
) error {
	if text == "" {
		uv.hasValue = false
		uv.value = 0

		return nil
	}
	// Check for negative values
	if text != "" && text[0] == '-' {
		return fmt.Errorf(
			"invalid uint32 value: negative values not allowed: %s",
			text,
		)
	}
	v, err := strconv.ParseUint(
		text,
		intParseBase,
		int32BitSize,
	)
	if err != nil {
		return fmt.Errorf(
			"invalid uint32 value: %w",
			err,
		)
	}
	uv.value = uint32(v)
	uv.hasValue = true

	return nil
}

// SetNil clears the value, making HasValue() return false.
func (uv *UInt32Value) SetNil() {
	uv.value = 0
	uv.hasValue = false
}

// Ensure Int32Value and UInt32Value implement SimpleValue and Resettable
// interfaces.
var (
	_ SimpleValue = (*Int32Value)(nil)
	_ Resettable  = (*Int32Value)(nil)
	_ SimpleValue = (*UInt32Value)(nil)
	_ Resettable  = (*UInt32Value)(nil)
)
