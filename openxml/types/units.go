package types

import (
	"fmt"
	"strconv"
)

// Unit conversion constants
const (
	// TwipsPerPoint is the number of twips in one point.
	// 1 point = 20 twips
	TwipsPerPoint = 20

	// TwipsPerInch is the number of twips in one inch.
	// 1 inch = 72 points = 1440 twips
	TwipsPerInch = 1440

	// TwipsPerCm is the number of twips in one centimeter.
	// 1 inch = 2.54 cm, so 1 cm = 1440/2.54 = ~567.0 twips
	TwipsPerCm = 567

	// PointsPerInch is the number of points in one inch.
	PointsPerInch = 72

	// EmuPerInch is the number of EMUs in one inch.
	// 1 inch = 914400 EMU
	EmuPerInch = 914400

	// EmuPerPoint is the number of EMUs in one point.
	// 1 point = 914400/72 = 12700 EMU
	EmuPerPoint = 12700

	// EmuPerCm is the number of EMUs in one centimeter.
	// 1 cm = 914400/2.54 = 360000 EMU
	EmuPerCm = 360000

	// EmuPerTwip is the number of EMUs in one twip.
	// 1 twip = 914400/1440 = 635 EMU
	EmuPerTwip = 635

	// HalfPointsPerPoint is the number of half-points in one point.
	HalfPointsPerPoint = 2
)

// TwipsValue wraps an integer value representing twentieths of a point.
// 1 inch = 1440 twips, 1 point = 20 twips.
// It implements the SimpleValue interface.
type TwipsValue struct {
	value    int64
	hasValue bool
}

// NewTwipsValue creates a new TwipsValue with the given value in twips.
func NewTwipsValue(twips int64) *TwipsValue {
	return &TwipsValue{
		value:    twips,
		hasValue: true,
	}
}

// NewTwipsValueFromPoints creates a new TwipsValue from a point value.
func NewTwipsValueFromPoints(points float64) *TwipsValue {
	return &TwipsValue{
		value:    int64(points * TwipsPerPoint),
		hasValue: true,
	}
}

// NewTwipsValueFromInches creates a new TwipsValue from an inch value.
func NewTwipsValueFromInches(inches float64) *TwipsValue {
	return &TwipsValue{
		value:    int64(inches * TwipsPerInch),
		hasValue: true,
	}
}

// NewNilTwipsValue creates a new TwipsValue in the unset/nil state.
func NewNilTwipsValue() *TwipsValue {
	return &TwipsValue{
		hasValue: false,
	}
}

// Value returns the raw twips value.
func (tv *TwipsValue) Value() int64 {
	if !tv.hasValue {
		return 0
	}
	return tv.value
}

// SetValue sets the raw twips value.
func (tv *TwipsValue) SetValue(twips int64) {
	tv.value = twips
	tv.hasValue = true
}

// ToPoints converts the twips value to points.
func (tv *TwipsValue) ToPoints() float64 {
	if !tv.hasValue {
		return 0
	}
	return float64(tv.value) / TwipsPerPoint
}

// ToInches converts the twips value to inches.
func (tv *TwipsValue) ToInches() float64 {
	if !tv.hasValue {
		return 0
	}
	return float64(tv.value) / TwipsPerInch
}

// ToCentimeters converts the twips value to centimeters.
func (tv *TwipsValue) ToCentimeters() float64 {
	if !tv.hasValue {
		return 0
	}
	return float64(tv.value) / TwipsPerCm
}

// HasValue returns true if the value is set.
func (tv *TwipsValue) HasValue() bool {
	return tv.hasValue
}

// InnerText returns the string representation for XML serialization.
func (tv *TwipsValue) InnerText() string {
	if !tv.hasValue {
		return ""
	}
	return strconv.FormatInt(tv.value, 10)
}

// SetInnerText parses the value from a string.
func (tv *TwipsValue) SetInnerText(text string) error {
	if text == "" {
		tv.hasValue = false
		tv.value = 0
		return nil
	}
	v, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid twips value: %w", err)
	}
	tv.value = v
	tv.hasValue = true
	return nil
}

// SetNil clears the value.
func (tv *TwipsValue) SetNil() {
	tv.value = 0
	tv.hasValue = false
}

// HalfPointsValue wraps an integer value representing half-points.
// Value 24 = 12 points.
// It implements the SimpleValue interface.
type HalfPointsValue struct {
	value    int64
	hasValue bool
}

// NewHalfPointsValue creates a new HalfPointsValue with the given value.
func NewHalfPointsValue(halfPoints int64) *HalfPointsValue {
	return &HalfPointsValue{
		value:    halfPoints,
		hasValue: true,
	}
}

// NewHalfPointsValueFromPoints creates a new HalfPointsValue from a point value.
func NewHalfPointsValueFromPoints(points float64) *HalfPointsValue {
	return &HalfPointsValue{
		value:    int64(points * HalfPointsPerPoint),
		hasValue: true,
	}
}

// NewNilHalfPointsValue creates a new HalfPointsValue in the unset/nil state.
func NewNilHalfPointsValue() *HalfPointsValue {
	return &HalfPointsValue{
		hasValue: false,
	}
}

// Value returns the raw half-points value.
func (hv *HalfPointsValue) Value() int64 {
	if !hv.hasValue {
		return 0
	}
	return hv.value
}

// SetValue sets the raw half-points value.
func (hv *HalfPointsValue) SetValue(halfPoints int64) {
	hv.value = halfPoints
	hv.hasValue = true
}

// ToPoints converts the half-points value to points.
func (hv *HalfPointsValue) ToPoints() float64 {
	if !hv.hasValue {
		return 0
	}
	return float64(hv.value) / HalfPointsPerPoint
}

// ToInches converts the half-points value to inches.
func (hv *HalfPointsValue) ToInches() float64 {
	if !hv.hasValue {
		return 0
	}
	return float64(hv.value) / (HalfPointsPerPoint * PointsPerInch)
}

// HasValue returns true if the value is set.
func (hv *HalfPointsValue) HasValue() bool {
	return hv.hasValue
}

// InnerText returns the string representation for XML serialization.
func (hv *HalfPointsValue) InnerText() string {
	if !hv.hasValue {
		return ""
	}
	return strconv.FormatInt(hv.value, 10)
}

// SetInnerText parses the value from a string.
func (hv *HalfPointsValue) SetInnerText(text string) error {
	if text == "" {
		hv.hasValue = false
		hv.value = 0
		return nil
	}
	v, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid half-points value: %w", err)
	}
	hv.value = v
	hv.hasValue = true
	return nil
}

// SetNil clears the value.
func (hv *HalfPointsValue) SetNil() {
	hv.value = 0
	hv.hasValue = false
}

// EmuValue wraps an integer value representing English Metric Units.
// 1 inch = 914400 EMU.
// It implements the SimpleValue interface.
type EmuValue struct {
	value    int64
	hasValue bool
}

// NewEmuValue creates a new EmuValue with the given value in EMUs.
func NewEmuValue(emu int64) *EmuValue {
	return &EmuValue{
		value:    emu,
		hasValue: true,
	}
}

// NewEmuValueFromPoints creates a new EmuValue from a point value.
func NewEmuValueFromPoints(points float64) *EmuValue {
	return &EmuValue{
		value:    int64(points * EmuPerPoint),
		hasValue: true,
	}
}

// NewEmuValueFromInches creates a new EmuValue from an inch value.
func NewEmuValueFromInches(inches float64) *EmuValue {
	return &EmuValue{
		value:    int64(inches * EmuPerInch),
		hasValue: true,
	}
}

// NewEmuValueFromCentimeters creates a new EmuValue from a centimeter value.
func NewEmuValueFromCentimeters(cm float64) *EmuValue {
	return &EmuValue{
		value:    int64(cm * EmuPerCm),
		hasValue: true,
	}
}

// NewNilEmuValue creates a new EmuValue in the unset/nil state.
func NewNilEmuValue() *EmuValue {
	return &EmuValue{
		hasValue: false,
	}
}

// Value returns the raw EMU value.
func (ev *EmuValue) Value() int64 {
	if !ev.hasValue {
		return 0
	}
	return ev.value
}

// SetValue sets the raw EMU value.
func (ev *EmuValue) SetValue(emu int64) {
	ev.value = emu
	ev.hasValue = true
}

// ToPoints converts the EMU value to points.
func (ev *EmuValue) ToPoints() float64 {
	if !ev.hasValue {
		return 0
	}
	return float64(ev.value) / EmuPerPoint
}

// ToInches converts the EMU value to inches.
func (ev *EmuValue) ToInches() float64 {
	if !ev.hasValue {
		return 0
	}
	return float64(ev.value) / EmuPerInch
}

// ToCentimeters converts the EMU value to centimeters.
func (ev *EmuValue) ToCentimeters() float64 {
	if !ev.hasValue {
		return 0
	}
	return float64(ev.value) / EmuPerCm
}

// ToTwips converts the EMU value to twips.
func (ev *EmuValue) ToTwips() int64 {
	if !ev.hasValue {
		return 0
	}
	return ev.value / EmuPerTwip
}

// HasValue returns true if the value is set.
func (ev *EmuValue) HasValue() bool {
	return ev.hasValue
}

// InnerText returns the string representation for XML serialization.
func (ev *EmuValue) InnerText() string {
	if !ev.hasValue {
		return ""
	}
	return strconv.FormatInt(ev.value, 10)
}

// SetInnerText parses the value from a string.
func (ev *EmuValue) SetInnerText(text string) error {
	if text == "" {
		ev.hasValue = false
		ev.value = 0
		return nil
	}
	v, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid EMU value: %w", err)
	}
	ev.value = v
	ev.hasValue = true
	return nil
}

// SetNil clears the value.
func (ev *EmuValue) SetNil() {
	ev.value = 0
	ev.hasValue = false
}

// Ensure unit types implement SimpleValue and Resettable interfaces.
var (
	_ SimpleValue = (*TwipsValue)(nil)
	_ Resettable  = (*TwipsValue)(nil)
	_ SimpleValue = (*HalfPointsValue)(nil)
	_ Resettable  = (*HalfPointsValue)(nil)
	_ SimpleValue = (*EmuValue)(nil)
	_ Resettable  = (*EmuValue)(nil)
)
