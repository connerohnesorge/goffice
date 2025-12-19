package types

import (
	"fmt"
	"strings"
)

// BooleanOutputFormat specifies how boolean values are serialized.
type BooleanOutputFormat int

const (
	// BooleanFormatTrueFalse serializes booleans as "true"/"false".
	BooleanFormatTrueFalse BooleanOutputFormat = iota
	// BooleanFormatOneZero serializes booleans as "1"/"0".
	BooleanFormatOneZero
)

// BooleanValue wraps a boolean value with optional nil/unset state.
// It implements the SimpleValue interface.
type BooleanValue struct {
	value        bool
	hasValue     bool
	outputFormat BooleanOutputFormat
}

// NewBooleanValue creates a new BooleanValue with the given boolean.
func NewBooleanValue(v bool) *BooleanValue {
	return &BooleanValue{
		value:        v,
		hasValue:     true,
		outputFormat: BooleanFormatTrueFalse,
	}
}

// NewBooleanValueWithFormat creates a new BooleanValue with the given boolean and output format.
func NewBooleanValueWithFormat(v bool, format BooleanOutputFormat) *BooleanValue {
	return &BooleanValue{
		value:        v,
		hasValue:     true,
		outputFormat: format,
	}
}

// NewNilBooleanValue creates a new BooleanValue in the unset/nil state.
func NewNilBooleanValue() *BooleanValue {
	return &BooleanValue{
		hasValue:     false,
		outputFormat: BooleanFormatTrueFalse,
	}
}

// Value returns the boolean value.
// Returns false if the value is not set.
func (bv *BooleanValue) Value() bool {
	if !bv.hasValue {
		return false
	}
	return bv.value
}

// SetValue sets the boolean value.
func (bv *BooleanValue) SetValue(v bool) {
	bv.value = v
	bv.hasValue = true
}

// HasValue returns true if the value is set.
func (bv *BooleanValue) HasValue() bool {
	return bv.hasValue
}

// SetOutputFormat sets the format used for serialization.
func (bv *BooleanValue) SetOutputFormat(format BooleanOutputFormat) {
	bv.outputFormat = format
}

// OutputFormat returns the current output format.
func (bv *BooleanValue) OutputFormat() BooleanOutputFormat {
	return bv.outputFormat
}

// InnerText returns the string representation for XML serialization.
func (bv *BooleanValue) InnerText() string {
	if !bv.hasValue {
		return ""
	}
	switch bv.outputFormat {
	case BooleanFormatOneZero:
		if bv.value {
			return "1"
		}
		return "0"
	default:
		if bv.value {
			return "true"
		}
		return "false"
	}
}

// SetInnerText parses the value from a string.
// Accepts "true", "false", "1", "0" (case-insensitive for true/false).
// Returns an error if the string cannot be parsed as a boolean.
func (bv *BooleanValue) SetInnerText(text string) error {
	if text == "" {
		bv.hasValue = false
		bv.value = false
		return nil
	}
	lower := strings.ToLower(strings.TrimSpace(text))
	switch lower {
	case "true", "1":
		bv.value = true
		bv.hasValue = true
		return nil
	case "false", "0":
		bv.value = false
		bv.hasValue = true
		return nil
	default:
		return fmt.Errorf("invalid boolean value: %q (expected true, false, 1, or 0)", text)
	}
}

// SetNil clears the value, making HasValue() return false.
func (bv *BooleanValue) SetNil() {
	bv.value = false
	bv.hasValue = false
}

// Ensure BooleanValue implements SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*BooleanValue)(nil)
	_ Resettable  = (*BooleanValue)(nil)
)
