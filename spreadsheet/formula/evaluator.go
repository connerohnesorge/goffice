package formula

import (
	"fmt"
	"math"
	"strconv"
)

// ValueType represents the type of a formula result
type ValueType int

const (
	// ValueEmpty represents an empty value
	ValueEmpty ValueType = iota

	// ValueNumber represents a numeric value
	ValueNumber

	// ValueString represents a string value
	ValueString

	// ValueBoolean represents a boolean value
	ValueBoolean

	// ValueError represents an error value
	ValueError

	// ValueArray represents an array value
	ValueArray
)

// String returns the string representation of ValueType
func (vt ValueType) String() string {
	switch vt {
	case ValueEmpty:
		return "Empty"
	case ValueNumber:
		return "Number"
	case ValueString:
		return "String"
	case ValueBoolean:
		return "Boolean"
	case ValueError:
		return "Error"
	case ValueArray:
		return "Array"
	default:
		return "Unknown"
	}
}

// Value represents a formula evaluation result
type Value struct {
	Type  ValueType
	Number float64
	String string
	Bool  bool
	Error string
	Array []Value
}

// IsEmpty returns true if the value is empty
func (v Value) IsEmpty() bool {
	return v.Type == ValueEmpty
}

// IsNumber returns true if the value is a number
func (v Value) IsNumber() bool {
	return v.Type == ValueNumber
}

// IsString returns true if the value is a string
func (v Value) IsString() bool {
	return v.Type == ValueString
}

// IsBoolean returns true if the value is a boolean
func (v Value) IsBoolean() bool {
	return v.Type == ValueBoolean
}

// IsError returns true if the value is an error
func (v Value) IsError() bool {
	return v.Type == ValueError
}

// IsArray returns true if the value is an array
func (v Value) IsArray() bool {
	return v.Type == ValueArray
}

// AsNumber returns the numeric value, or 0 if not a number
func (v *Value) AsNumber() float64 {
	if v.Type == ValueNumber {
		return v.Number
	}
	return 0
}

// AsString returns the string value, or empty string if not a string
func (v *Value) AsString() string {
	if v.Type == ValueString {
		return v.String
	}
	return ""
}

// AsBoolean returns the boolean value, or false if not a boolean
func (v *Value) AsBoolean() bool {
	if v.Type == ValueBoolean {
		return v.Bool
	}
	return false
}

// AsError returns the error value, or empty string if not an error
func (v *Value) AsError() string {
	if v.Type == ValueError {
		return v.Error
	}
	return ""
}

// AsArray returns the array value, or nil if not an array
func (v *Value) AsArray() []Value {
	if v.Type == ValueArray {
		return v.Array
	}
	return nil
}

// NewEmptyValue creates an empty value
func NewEmptyValue() Value {
	return Value{Type: ValueEmpty}
}

// NewNumberValue creates a numeric value
func NewNumberValue(num float64) Value {
	return Value{
		Type:  ValueNumber,
		Number: num,
	}
}

// NewStringValue creates a string value
func NewStringValue(str string) Value {
	return Value{
		Type:  ValueString,
		String: str,
	}
}

// NewBooleanValue creates a boolean value
func NewBooleanValue(b bool) Value {
	return Value{
		Type:  ValueBoolean,
		Bool:  b,
	}
}

// NewErrorValue creates an error value
func NewErrorValue(err string) Value {
	return Value{
		Type:  ValueError,
		Error: err,
	}
}

// NewArrayValue creates an array value
func NewArrayValue(arr []Value) Value {
	return Value{
		Type:  ValueArray,
		Array: arr,
	}
}

// String returns a string representation of the value
func (v Value) String() string {
	switch v.Type {
	case ValueEmpty:
		return ""
	case ValueNumber:
		return fmt.Sprintf("%g", v.Number)
	case ValueString:
		return v.String
	case ValueBoolean:
		if v.Bool {
			return "TRUE"
		}
		return "FALSE"
	case ValueError:
		return v.Error
	case ValueArray:
		var strs []string
		for _, val := range v.Array {
			strs = append(strs, val.String())
		}
		return "{" + strings.Join(strs, ",") + "}"
	default:
		return "Unknown"
	}
}

// Equal returns true if two values are equal
func (v Value) Equal(other Value) bool {
	if v.Type != other.Type {
		return false
	}
	switch v.Type {
	case ValueEmpty:
		return other.Type == ValueEmpty
	case ValueNumber:
		return v.Number == other.Number
	case ValueString:
		return v.String == other.String
	case ValueBoolean:
		return v.Bool == other.Bool
	case ValueError:
		return v.Error == other.Error
	case ValueArray:
		if len(v.Array) != len(other.Array) {
			return false
		}
		for i := range v.Array {
			if !v.Array[i].Equal(other.Array[i]) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
}

// Value represents a formula evaluation result
type Value struct {
	Type   ValueType
	Number float64
	String string
	Bool   bool
	Error  string
	Array  []Value
}

// IsEmpty returns true if the value is empty
func (v Value) IsEmpty() bool {
	return v.Type == ValueEmpty
}

// IsNumber returns true if the value is a number
func (v Value) IsNumber() bool {
	return v.Type == ValueNumber
}

// IsString returns true if the value is a string
func (v Value) IsString() bool {
	return v.Type == ValueString
}

// IsBoolean returns true if the value is a boolean
func (v Value) IsBoolean() bool {
	return v.Type == ValueBoolean
}

// IsError returns true if the value is an error
func (v Value) IsError() bool {
	return v.Type == ValueError
}

// IsArray returns true if the value is an array
func (v Value) IsArray() bool {
	return v.Type == ValueArray
}

// AsNumber returns the numeric value, or 0 if not a number
func (v Value) AsNumber() float64 {
	if v.Type == ValueNumber {
		return v.Number
	}
	return 0
}

// AsString returns the string value, or empty string if not a string
func (v Value) AsString() string {
	if v.Type == ValueString {
		return v.String
	}
	return ""
}

// AsBoolean returns the boolean value, or false if not a boolean
func (v Value) AsBoolean() bool {
	if v.Type == ValueBoolean {
		return v.Bool
	}
	return false
}

// AsError returns the error value, or empty string if not an error
func (v Value) AsError() string {
	if v.Type == ValueError {
		return v.Error
	}
	return ""
}

// AsArray returns the array value, or nil if not an array
func (v Value) AsArray() []Value {
	if v.Type == ValueArray {
		return v.Array
	}
	return nil
}

// NewEmptyValue creates an empty value
func NewEmptyValue() Value {
	return Value{Type: ValueEmpty}
}

// NewNumberValue creates a numeric value
func NewNumberValue(num float64) Value {
	return Value{
		Type:   ValueNumber,
		Number: num,
	}
}

// NewStringValue creates a string value
func NewStringValue(str string) Value {
	return Value{
		Type:   ValueString,
		String: str,
	}
}

// NewBooleanValue creates a boolean value
func NewBooleanValue(b bool) Value {
	return Value{
		Type: ValueBoolean,
		Bool: b,
	}
}

// NewErrorValue creates an error value
func NewErrorValue(err string) Value {
	return Value{
		Type:  ValueError,
		Error: err,
	}
}

// NewArrayValue creates an array value
func NewArrayValue(arr []Value) Value {
	return Value{
		Type:  ValueArray,
		Array: arr,
	}
}

// String returns a string representation of the value
func (v Value) String() string {
	switch v.Type {
	case ValueEmpty:
		return ""
	case ValueNumber:
		return fmt.Sprintf("%g", v.Number)
	case ValueString:
		return v.String
	case ValueBoolean:
		if v.Bool {
			return "TRUE"
		}
		return "FALSE"
	case ValueError:
		return v.Error
	case ValueArray:
		var strs []string
		for _, val := range v.Array {
			strs = append(strs, val.String())
		}
		return "{" + strings.Join(strs, ",") + "}"
	default:
		return "Unknown"
	}
}

// Equal returns true if two values are equal
func (v Value) Equal(other Value) bool {
	if v.Type != other.Type {
		return false
	}
	switch v.Type {
	case ValueEmpty:
		return other.Type == ValueEmpty
	case ValueNumber:
		return v.Number == other.Number
	case ValueString:
		return v.String == other.String
	case ValueBoolean:
		return v.Bool == other.Bool
	case ValueError:
		return v.Error == other.Error
	case ValueArray:
		if len(v.Array) != len(other.Array) {
			return false
		}
		for i := range v.Array {
			if !v.Array[i].Equal(other.Array[i]) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
