package types

import (
	"fmt"
	"strconv"
	"strings"
)

// PercentageValue wraps an integer value representing a percentage.
// OOXML uses 1000-based percentages where 50000 = 50%.
// It implements the SimpleValue interface.
type PercentageValue struct {
	value    int64 // 1000-based: 50000 = 50%
	hasValue bool
}

// PercentageScale is the multiplier for percentage values in OOXML.
// 50000 = 50%, so the scale is 1000.
const PercentageScale = 1000

// NewPercentageValue creates a new PercentageValue with the given raw value.
// The value is in 1000-based format (50000 = 50%).
func NewPercentageValue(
	value int64,
) *PercentageValue {
	return &PercentageValue{
		value:    value,
		hasValue: true,
	}
}

// NewPercentageValueFromPercent creates a new PercentageValue from a percentage.
// For example, 50 creates a value representing 50%.
func NewPercentageValueFromPercent(
	percent float64,
) *PercentageValue {
	return &PercentageValue{
		value: int64(
			percent * PercentageScale,
		),
		hasValue: true,
	}
}

// NewPercentageValueFromFloat creates a new PercentageValue from a decimal.
// For example, 0.5 creates a value representing 50%.
func NewPercentageValueFromFloat(
	f float64,
) *PercentageValue {
	return &PercentageValue{
		value: int64(
			f * 100 * PercentageScale,
		),
		hasValue: true,
	}
}

// NewNilPercentageValue creates a new PercentageValue in the unset/nil state.
func NewNilPercentageValue() *PercentageValue {
	return &PercentageValue{
		hasValue: false,
	}
}

// Value returns the raw 1000-based value.
// Returns 0 if the value is not set.
func (pv *PercentageValue) Value() int64 {
	if !pv.hasValue {
		return 0
	}
	return pv.value
}

// SetValue sets the raw 1000-based value.
func (pv *PercentageValue) SetValue(value int64) {
	pv.value = value
	pv.hasValue = true
}

// ToFloat returns the percentage as a decimal (0.5 for 50%).
func (pv *PercentageValue) ToFloat() float64 {
	if !pv.hasValue {
		return 0
	}
	return float64(
		pv.value,
	) / (100 * PercentageScale)
}

// ToPercent returns the percentage as a number (50 for 50%).
func (pv *PercentageValue) ToPercent() float64 {
	if !pv.hasValue {
		return 0
	}
	return float64(pv.value) / PercentageScale
}

// HasValue returns true if the value is set.
func (pv *PercentageValue) HasValue() bool {
	return pv.hasValue
}

// InnerText returns the string representation for XML serialization.
// Returns the raw 1000-based integer value as a string.
func (pv *PercentageValue) InnerText() string {
	if !pv.hasValue {
		return ""
	}
	return strconv.FormatInt(pv.value, 10)
}

// SetInnerText parses the value from a string.
// Accepts either raw integer values (50000) or percentage strings (50%).
func (pv *PercentageValue) SetInnerText(
	text string,
) error {
	if text == "" {
		pv.hasValue = false
		pv.value = 0
		return nil
	}

	text = strings.TrimSpace(text)

	// Check if it's a percentage string (e.g., "50%")
	if strings.HasSuffix(text, "%") {
		percentStr := strings.TrimSuffix(
			text,
			"%",
		)
		percent, err := strconv.ParseFloat(
			percentStr,
			64,
		)
		if err != nil {
			return fmt.Errorf(
				"invalid percentage value: %w",
				err,
			)
		}
		pv.value = int64(
			percent * PercentageScale,
		)
		pv.hasValue = true
		return nil
	}

	// Otherwise, parse as raw integer
	v, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return fmt.Errorf(
			"invalid percentage value: %w",
			err,
		)
	}
	pv.value = v
	pv.hasValue = true
	return nil
}

// SetNil clears the value, making HasValue() return false.
func (pv *PercentageValue) SetNil() {
	pv.value = 0
	pv.hasValue = false
}

// Ensure PercentageValue implements SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*PercentageValue)(nil)
	_ Resettable  = (*PercentageValue)(nil)
)
