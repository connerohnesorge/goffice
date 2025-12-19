// Package types provides simple value types for Office Open XML attributes.
// These types handle the conversion between XML string representations
// and Go native types (StringValue, Int32Value, BooleanValue, etc.).
//
// All value types implement the SimpleValue interface which provides:
//   - HasValue() bool - check if a value is set
//   - InnerText() string - get XML string representation
//   - SetInnerText(text string) error - parse from XML string
//
// Most types also implement Resettable which provides:
//   - SetNil() - clear the value to unset state
//
// # Available Types
//
// String and Numeric Types:
//   - StringValue - wraps string values
//   - Int32Value, UInt32Value - 32-bit integers
//   - Int64Value, UInt64Value - 64-bit integers
//   - DecimalValue - floating point (float64)
//
// Boolean Types:
//   - BooleanValue - standard true/false/1/0
//   - OnOffValue - Word-style on/off/true/false/1/0
//
// Binary Types:
//   - HexBinaryValue - hexadecimal encoded bytes
//   - Base64BinaryValue - base64 encoded bytes
//
// Date and Enum Types:
//   - DateTimeValue - ISO 8601 date/time (wraps time.Time)
//   - EnumValue[T] - generic enumeration type
//
// Measurement Units:
//   - TwipsValue - twentieths of a point (1440 twips = 1 inch)
//   - HalfPointsValue - half points (24 = 12 points)
//   - EmuValue - English Metric Units (914400 EMU = 1 inch)
//   - PercentageValue - 1000-based percentages (50000 = 50%)
package types
