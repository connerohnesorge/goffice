package types

import (
	"fmt"
	"strings"
)

// OnOffOutputFormat specifies how on/off values are serialized.
type OnOffOutputFormat int

const (
	// OnOffFormatOnOff serializes values as "on"/"off".
	OnOffFormatOnOff OnOffOutputFormat = iota
	// OnOffFormatTrueFalse serializes values as "true"/"false".
	OnOffFormatTrueFalse
	// OnOffFormatOneZero serializes values as "1"/"0".
	OnOffFormatOneZero
)

// OnOffValue wraps a Word-style on/off boolean value with optional nil/unset state.
// It implements the SimpleValue interface.
// This type accepts "on", "off", "true", "false", "1", "0" on input.
type OnOffValue struct {
	value        bool
	hasValue     bool
	outputFormat OnOffOutputFormat
}

// NewOnOffValue creates a new OnOffValue with the given boolean.
func NewOnOffValue(v bool) *OnOffValue {
	return &OnOffValue{
		value:        v,
		hasValue:     true,
		outputFormat: OnOffFormatOnOff,
	}
}

// NewOnOffValueWithFormat creates a new OnOffValue with the given boolean and output format.
func NewOnOffValueWithFormat(
	v bool,
	format OnOffOutputFormat,
) *OnOffValue {
	return &OnOffValue{
		value:        v,
		hasValue:     true,
		outputFormat: format,
	}
}

// NewNilOnOffValue creates a new OnOffValue in the unset/nil state.
func NewNilOnOffValue() *OnOffValue {
	return &OnOffValue{
		hasValue:     false,
		outputFormat: OnOffFormatOnOff,
	}
}

// Value returns the boolean value.
// Returns false if the value is not set.
func (ov *OnOffValue) Value() bool {
	if !ov.hasValue {
		return false
	}

	return ov.value
}

// SetValue sets the boolean value.
func (ov *OnOffValue) SetValue(v bool) {
	ov.value = v
	ov.hasValue = true
}

// HasValue returns true if the value is set.
func (ov *OnOffValue) HasValue() bool {
	return ov.hasValue
}

// SetOutputFormat sets the format used for serialization.
func (ov *OnOffValue) SetOutputFormat(
	format OnOffOutputFormat,
) {
	ov.outputFormat = format
}

// OutputFormat returns the current output format.
func (ov *OnOffValue) OutputFormat() OnOffOutputFormat {
	return ov.outputFormat
}

// InnerText returns the string representation for XML serialization.
func (ov *OnOffValue) InnerText() string {
	if !ov.hasValue {
		return ""
	}
	switch ov.outputFormat {
	case OnOffFormatTrueFalse:
		if ov.value {
			return "true"
		}

		return "false"
	case OnOffFormatOneZero:
		if ov.value {
			return "1"
		}

		return "0"
	default: // OnOffFormatOnOff
		if ov.value {
			return "on"
		}

		return "off"
	}
}

// SetInnerText parses the value from a string.
// Accepts "on", "off", "true", "false", "1", "0" (case-insensitive for text values).
// Returns an error if the string cannot be parsed.
func (ov *OnOffValue) SetInnerText(
	text string,
) error {
	if text == "" {
		ov.hasValue = false
		ov.value = false

		return nil
	}
	lower := strings.ToLower(
		strings.TrimSpace(text),
	)
	switch lower {
	case "on", "true", "1":
		ov.value = true
		ov.hasValue = true

		return nil
	case "off", "false", "0":
		ov.value = false
		ov.hasValue = true

		return nil
	default:
		return fmt.Errorf(
			"invalid on/off value: %q (expected on, off, true, false, 1, or 0)",
			text,
		)
	}
}

// SetNil clears the value, making HasValue() return false.
func (ov *OnOffValue) SetNil() {
	ov.value = false
	ov.hasValue = false
}

// Ensure OnOffValue implements SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*OnOffValue)(nil)
	_ Resettable  = (*OnOffValue)(nil)
)
