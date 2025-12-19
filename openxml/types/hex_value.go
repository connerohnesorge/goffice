package types

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// HexBinaryValue wraps a byte slice value represented as hexadecimal in XML.
// It implements the SimpleValue interface.
type HexBinaryValue struct {
	value    []byte
	hasValue bool
}

// NewHexBinaryValue creates a new HexBinaryValue with the given bytes.
func NewHexBinaryValue(v []byte) *HexBinaryValue {
	// Make a copy to avoid aliasing issues
	valueCopy := make([]byte, len(v))
	copy(valueCopy, v)
	return &HexBinaryValue{
		value:    valueCopy,
		hasValue: true,
	}
}

// NewHexBinaryValueFromString creates a new HexBinaryValue by parsing a hex string.
// Returns an error if the string is not valid hexadecimal.
func NewHexBinaryValueFromString(
	s string,
) (*HexBinaryValue, error) {
	hv := &HexBinaryValue{}
	if err := hv.SetInnerText(s); err != nil {
		return nil, err
	}
	return hv, nil
}

// NewNilHexBinaryValue creates a new HexBinaryValue in the unset/nil state.
func NewNilHexBinaryValue() *HexBinaryValue {
	return &HexBinaryValue{
		hasValue: false,
	}
}

// Value returns the byte slice value.
// Returns nil if the value is not set.
func (hv *HexBinaryValue) Value() []byte {
	if !hv.hasValue {
		return nil
	}
	// Return a copy to prevent modification of internal state
	valueCopy := make([]byte, len(hv.value))
	copy(valueCopy, hv.value)
	return valueCopy
}

// SetValue sets the byte slice value.
func (hv *HexBinaryValue) SetValue(v []byte) {
	// Make a copy to avoid aliasing issues
	hv.value = make([]byte, len(v))
	copy(hv.value, v)
	hv.hasValue = true
}

// HasValue returns true if the value is set.
func (hv *HexBinaryValue) HasValue() bool {
	return hv.hasValue
}

// InnerText returns the uppercase hexadecimal string representation for XML serialization.
func (hv *HexBinaryValue) InnerText() string {
	if !hv.hasValue {
		return ""
	}
	return strings.ToUpper(
		hex.EncodeToString(hv.value),
	)
}

// SetInnerText parses the value from a hexadecimal string.
// Accepts both uppercase and lowercase hex digits.
// Returns an error if the string is not valid hexadecimal.
func (hv *HexBinaryValue) SetInnerText(
	text string,
) error {
	if text == "" {
		hv.hasValue = false
		hv.value = nil
		return nil
	}
	decoded, err := hex.DecodeString(text)
	if err != nil {
		return fmt.Errorf(
			"invalid hexadecimal value: %w",
			err,
		)
	}
	hv.value = decoded
	hv.hasValue = true
	return nil
}

// SetNil clears the value, making HasValue() return false.
func (hv *HexBinaryValue) SetNil() {
	hv.value = nil
	hv.hasValue = false
}

// Len returns the length of the byte slice value.
// Returns 0 if the value is not set.
func (hv *HexBinaryValue) Len() int {
	if !hv.hasValue {
		return 0
	}
	return len(hv.value)
}

// Ensure HexBinaryValue implements SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*HexBinaryValue)(nil)
	_ Resettable  = (*HexBinaryValue)(nil)
)
