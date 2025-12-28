package types

import (
	"fmt"
	"strconv"
)

// byteBitSize is the bit size for parsing byte values.
const byteBitSize = 8

// ByteValue wraps an unsigned 8-bit integer value with optional
// nil/unset state. It implements the SimpleValue interface.
type ByteValue struct {
	value    byte
	hasValue bool
}

// NewByteValue creates a new ByteValue with the given byte.
func NewByteValue(v byte) *ByteValue {
	return &ByteValue{
		value:    v,
		hasValue: true,
	}
}

// NewNilByteValue creates a new ByteValue in the unset/nil state.
func NewNilByteValue() *ByteValue {
	return &ByteValue{
		hasValue: false,
	}
}

// Value returns the byte value.
// Returns 0 if the value is not set.
func (bv *ByteValue) Value() byte {
	if !bv.hasValue {
		return 0
	}

	return bv.value
}

// SetValue sets the byte value.
func (bv *ByteValue) SetValue(v byte) {
	bv.value = v
	bv.hasValue = true
}

// HasValue returns true if the value is set.
func (bv *ByteValue) HasValue() bool {
	return bv.hasValue
}

// InnerText returns the string representation for XML serialization.
func (bv *ByteValue) InnerText() string {
	if !bv.hasValue {
		return ""
	}

	return strconv.FormatUint(
		uint64(bv.value),
		intParseBase,
	)
}

// SetInnerText parses the value from a string.
// Returns an error if the string cannot be parsed as a byte or is negative.
func (bv *ByteValue) SetInnerText(
	text string,
) error {
	if text == "" {
		bv.hasValue = false
		bv.value = 0

		return nil
	}
	// Check for negative values
	if text[0] == '-' {
		return fmt.Errorf(
			"invalid byte value: negative values not allowed: %s",
			text,
		)
	}
	v, err := strconv.ParseUint(
		text,
		intParseBase,
		byteBitSize,
	)
	if err != nil {
		return fmt.Errorf(
			"invalid byte value: %w",
			err,
		)
	}
	bv.value = byte(v)
	bv.hasValue = true

	return nil
}

// SetNil clears the value, making HasValue() return false.
func (bv *ByteValue) SetNil() {
	bv.value = 0
	bv.hasValue = false
}

// Ensure ByteValue implements SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*ByteValue)(nil)
	_ Resettable  = (*ByteValue)(nil)
)
