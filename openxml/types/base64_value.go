package types

import (
	"encoding/base64"
	"fmt"
)

// Base64BinaryValue wraps a byte slice value represented as base64 in XML.
// It implements the SimpleValue interface.
type Base64BinaryValue struct {
	value    []byte
	hasValue bool
}

// NewBase64BinaryValue creates a new Base64BinaryValue with the given bytes.
func NewBase64BinaryValue(v []byte) *Base64BinaryValue {
	// Make a copy to avoid aliasing issues
	valueCopy := make([]byte, len(v))
	copy(valueCopy, v)
	return &Base64BinaryValue{
		value:    valueCopy,
		hasValue: true,
	}
}

// NewBase64BinaryValueFromString creates a new Base64BinaryValue by parsing a base64 string.
// Returns an error if the string is not valid base64.
func NewBase64BinaryValueFromString(s string) (*Base64BinaryValue, error) {
	bv := &Base64BinaryValue{}
	if err := bv.SetInnerText(s); err != nil {
		return nil, err
	}
	return bv, nil
}

// NewNilBase64BinaryValue creates a new Base64BinaryValue in the unset/nil state.
func NewNilBase64BinaryValue() *Base64BinaryValue {
	return &Base64BinaryValue{
		hasValue: false,
	}
}

// Value returns the byte slice value.
// Returns nil if the value is not set.
func (bv *Base64BinaryValue) Value() []byte {
	if !bv.hasValue {
		return nil
	}
	// Return a copy to prevent modification of internal state
	valueCopy := make([]byte, len(bv.value))
	copy(valueCopy, bv.value)
	return valueCopy
}

// SetValue sets the byte slice value.
func (bv *Base64BinaryValue) SetValue(v []byte) {
	// Make a copy to avoid aliasing issues
	bv.value = make([]byte, len(v))
	copy(bv.value, v)
	bv.hasValue = true
}

// HasValue returns true if the value is set.
func (bv *Base64BinaryValue) HasValue() bool {
	return bv.hasValue
}

// InnerText returns the base64 encoded string representation for XML serialization.
func (bv *Base64BinaryValue) InnerText() string {
	if !bv.hasValue {
		return ""
	}
	return base64.StdEncoding.EncodeToString(bv.value)
}

// SetInnerText parses the value from a base64 encoded string.
// Returns an error if the string is not valid base64.
func (bv *Base64BinaryValue) SetInnerText(text string) error {
	if text == "" {
		bv.hasValue = false
		bv.value = nil
		return nil
	}
	decoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return fmt.Errorf("invalid base64 value: %w", err)
	}
	bv.value = decoded
	bv.hasValue = true
	return nil
}

// SetNil clears the value, making HasValue() return false.
func (bv *Base64BinaryValue) SetNil() {
	bv.value = nil
	bv.hasValue = false
}

// Len returns the length of the byte slice value.
// Returns 0 if the value is not set.
func (bv *Base64BinaryValue) Len() int {
	if !bv.hasValue {
		return 0
	}
	return len(bv.value)
}

// Ensure Base64BinaryValue implements SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*Base64BinaryValue)(nil)
	_ Resettable  = (*Base64BinaryValue)(nil)
)
