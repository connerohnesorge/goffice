# Simple Types

Delta specification for OpenXML simple value types in Go.

## ADDED Requirements

### Requirement: OpenXmlSimpleType Base

The system SHALL provide a base interface for all simple types with common value semantics.

#### Scenario: Value presence check
- WHEN a simple type has a value set
- THEN `HasValue()` returns true
- AND when no value is set, `HasValue()` returns false

#### Scenario: Inner text representation
- WHEN `InnerText()` is called
- THEN the XML string representation is returned
- AND when no value, empty string is returned

#### Scenario: Clone operation
- WHEN `Clone()` is called
- THEN a new instance with the same value is returned

### Requirement: StringValue Type

The system SHALL provide `StringValue` for string attribute values with nil distinction.

#### Scenario: String value creation
- WHEN `NewStringValue("hello")` is called
- THEN `Value()` returns "hello"
- AND `HasValue()` returns true

#### Scenario: Empty vs nil string
- WHEN value is set to empty string ""
- THEN `HasValue()` returns true
- AND `Value()` returns ""
- WHEN value is not set
- THEN `HasValue()` returns false

#### Scenario: String implicit conversion
- WHEN assigning a StringValue to string variable
- THEN the conversion method extracts the value
- AND when assigning string to StringValue, it wraps

### Requirement: BooleanValue Type

The system SHALL provide `BooleanValue` supporting standard XML boolean formats.

#### Scenario: Boolean true values
- WHEN parsing "true", "1", or "True"
- THEN `Value()` returns true
- AND `HasValue()` returns true

#### Scenario: Boolean false values
- WHEN parsing "false", "0", or "False"
- THEN `Value()` returns false
- AND `HasValue()` returns true

#### Scenario: Boolean serialization
- WHEN serializing a true value
- THEN output is "true" (lowercase)
- AND when serializing false, output is "false"

### Requirement: Numeric Value Types

The system SHALL provide numeric value types for all OpenXML numeric types.

#### Scenario: Int32Value operations
- WHEN `NewInt32Value(42)` is called
- THEN `Value()` returns 42
- AND `InnerText()` returns "42"

#### Scenario: Int64Value operations
- WHEN `NewInt64Value(9223372036854775807)` is called
- THEN large values are supported
- AND serialization preserves precision

#### Scenario: UInt32Value operations
- WHEN `NewUInt32Value(4294967295)` is called
- THEN maximum uint32 is supported
- AND negative values are rejected

#### Scenario: DoubleValue operations
- WHEN `NewDoubleValue(3.14159)` is called
- THEN floating-point values are supported
- AND scientific notation is handled

#### Scenario: DecimalValue precision
- WHEN `NewDecimalValue("123.456789012345")` is called
- THEN full decimal precision is preserved
- AND rounding follows XML Schema rules

### Requirement: DateTimeValue Type

The system SHALL provide `DateTimeValue` with ISO 8601 date/time handling.

#### Scenario: DateTime parsing
- WHEN parsing "2024-01-15T10:30:00Z"
- THEN the value is correctly parsed
- AND timezone is preserved

#### Scenario: DateTime serialization
- WHEN serializing a DateTime value
- THEN ISO 8601 format is used
- AND UTC timezone is indicated with "Z"

### Requirement: HexBinaryValue Type

The system SHALL provide `HexBinaryValue` for hexadecimal binary data.

#### Scenario: Hex encoding
- WHEN `NewHexBinaryValue([]byte{0xDE, 0xAD, 0xBE, 0xEF})` is called
- THEN `InnerText()` returns "DEADBEEF"
- AND case is uppercase

#### Scenario: Hex decoding
- WHEN parsing "deadbeef" or "DEADBEEF"
- THEN `Value()` returns []byte{0xDE, 0xAD, 0xBE, 0xEF}
- AND case is ignored

#### Scenario: Hex validation
- WHEN parsing invalid hex "GHIJ"
- THEN an error is returned
- AND `HasValue()` returns false

### Requirement: Base64BinaryValue Type

The system SHALL provide `Base64BinaryValue` for Base64-encoded binary data.

#### Scenario: Base64 encoding
- WHEN `NewBase64BinaryValue([]byte("hello"))` is called
- THEN `InnerText()` returns "aGVsbG8="

#### Scenario: Base64 decoding
- WHEN parsing "aGVsbG8="
- THEN `Value()` returns []byte("hello")

### Requirement: EnumValue Generic Type

The system SHALL provide `EnumValue[T]` for strongly-typed enumeration values.

#### Scenario: Enum value creation
- WHEN `NewEnumValue[SlideLayoutType](SlideLayoutTypeTitle)` is called
- THEN `Value()` returns the enum constant
- AND `InnerText()` returns the XML string representation

#### Scenario: Enum parsing
- WHEN parsing "title" for SlideLayoutType
- THEN the corresponding enum value is returned
- AND invalid strings return error

#### Scenario: Enum serialization
- WHEN serializing an enum value
- THEN the XML string mapping is used
- AND not the Go constant name

### Requirement: ListValue Generic Type

The system SHALL provide `ListValue[T]` for space-separated list values.

#### Scenario: List creation
- WHEN `NewListValue[StringValue]([]StringValue{...})` is called
- THEN `Items()` returns the slice of values

#### Scenario: List parsing
- WHEN parsing "one two three"
- THEN three StringValue items are created
- AND whitespace is used as separator

#### Scenario: List serialization
- WHEN serializing a list of values
- THEN items are space-separated
- AND empty items are handled correctly

### Requirement: OnOffValue Type

The system SHALL provide `OnOffValue` for Office-style boolean (on/off/1/0/true/false).

#### Scenario: OnOff "on" value
- WHEN parsing "on" or "1" or "true"
- THEN `Value()` returns true

#### Scenario: OnOff "off" value
- WHEN parsing "off" or "0" or "false"
- THEN `Value()` returns false

#### Scenario: OnOff default serialization
- WHEN serializing OnOffValue
- THEN output is "1" for true, "0" for false by default

### Requirement: TrueFalseValue Type

The system SHALL provide `TrueFalseValue` for t/f/true/false format.

#### Scenario: TrueFalse parsing
- WHEN parsing "t" or "true"
- THEN `Value()` returns true
- WHEN parsing "f" or "false"
- THEN `Value()` returns false

#### Scenario: TrueFalse serialization
- WHEN serializing TrueFalseValue
- THEN output is "t" or "f" (short form)

### Requirement: TrueFalseBlankValue Type

The system SHALL provide `TrueFalseBlankValue` supporting blank/empty as valid value.

#### Scenario: Blank value handling
- WHEN parsing empty string ""
- THEN `HasValue()` returns true
- AND `IsBlank()` returns true

#### Scenario: TrueFalseBlank true/false
- WHEN parsing "true" or "false"
- THEN behaves like TrueFalseValue
- AND `IsBlank()` returns false
