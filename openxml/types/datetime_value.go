package types

import (
	"fmt"
	"time"
)

// DateTimeValue wraps a time.Time value for XML date/time attributes.
// It uses ISO 8601 format for serialization.
// It implements the SimpleValue interface.
type DateTimeValue struct {
	value    time.Time
	hasValue bool
}

// Common ISO 8601 formats used in Office Open XML
var dateTimeFormats = []string{
	time.RFC3339,                    // "2006-01-02T15:04:05Z07:00"
	"2006-01-02T15:04:05Z",          // UTC with Z suffix
	"2006-01-02T15:04:05",           // No timezone
	"2006-01-02T15:04:05.000Z",      // With milliseconds and Z
	"2006-01-02T15:04:05.000",       // With milliseconds, no TZ
	"2006-01-02T15:04:05.000000Z",   // With microseconds and Z
	"2006-01-02T15:04:05.000000",    // With microseconds, no TZ
	"2006-01-02",                    // Date only
}

// NewDateTimeValue creates a new DateTimeValue with the given time.
func NewDateTimeValue(t time.Time) *DateTimeValue {
	return &DateTimeValue{
		value:    t,
		hasValue: true,
	}
}

// NewNilDateTimeValue creates a new DateTimeValue in the unset/nil state.
func NewNilDateTimeValue() *DateTimeValue {
	return &DateTimeValue{
		hasValue: false,
	}
}

// Value returns the time.Time value.
// Returns the zero time if the value is not set.
func (dv *DateTimeValue) Value() time.Time {
	if !dv.hasValue {
		return time.Time{}
	}
	return dv.value
}

// SetValue sets the time.Time value.
func (dv *DateTimeValue) SetValue(t time.Time) {
	dv.value = t
	dv.hasValue = true
}

// HasValue returns true if the value is set.
func (dv *DateTimeValue) HasValue() bool {
	return dv.hasValue
}

// InnerText returns the ISO 8601 formatted string for XML serialization.
// Uses RFC3339 format (e.g., "2006-01-02T15:04:05Z07:00").
func (dv *DateTimeValue) InnerText() string {
	if !dv.hasValue {
		return ""
	}
	return dv.value.Format(time.RFC3339)
}

// SetInnerText parses the value from an ISO 8601 formatted string.
// Supports various ISO 8601 variants.
// Returns an error if the string cannot be parsed.
func (dv *DateTimeValue) SetInnerText(text string) error {
	if text == "" {
		dv.hasValue = false
		dv.value = time.Time{}
		return nil
	}

	// Try each format until one works
	var parseErr error
	for _, format := range dateTimeFormats {
		t, err := time.Parse(format, text)
		if err == nil {
			dv.value = t
			dv.hasValue = true
			return nil
		}
		parseErr = err
	}

	return fmt.Errorf("invalid datetime value %q: %w", text, parseErr)
}

// SetNil clears the value, making HasValue() return false.
func (dv *DateTimeValue) SetNil() {
	dv.value = time.Time{}
	dv.hasValue = false
}

// IsZero returns true if the value is not set or is the zero time.
func (dv *DateTimeValue) IsZero() bool {
	return !dv.hasValue || dv.value.IsZero()
}

// Ensure DateTimeValue implements SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*DateTimeValue)(nil)
	_ Resettable  = (*DateTimeValue)(nil)
)
