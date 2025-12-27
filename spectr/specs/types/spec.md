# Types Specification

## Requirements

### Requirement: Simple Value Base Interface
The system SHALL provide a base interface for all typed attribute values.

#### Scenario: Check if value is set
- GIVEN a SimpleValue instance
- WHEN HasValue() is called
- THEN true if value is set, false if nil/unset

#### Scenario: Get inner text representation
- GIVEN a SimpleValue with a value
- WHEN InnerText() is called
- THEN the string representation for XML is returned

#### Scenario: Set from inner text
- GIVEN a SimpleValue
- WHEN SetInnerText(text) is called
- THEN the value is parsed from the string

#### Scenario: Unset value
- GIVEN a SimpleValue with a value
- WHEN SetInnerText("") is called with empty or SetNil()
- THEN the value becomes unset (HasValue() returns false)

### Requirement: StringValue Type
The system SHALL provide a typed wrapper for string attributes.

#### Scenario: Create string value
- GIVEN a string "hello"
- WHEN NewStringValue("hello") is called
- THEN a StringValue with that value is created

#### Scenario: Get string value
- GIVEN a StringValue with value "hello"
- WHEN Value() is called
- THEN "hello" is returned

#### Scenario: Set string value
- GIVEN a StringValue
- WHEN SetValue("world") is called
- THEN the value is updated to "world"

#### Scenario: Nil string value
- GIVEN an unset StringValue
- WHEN Value() is called
- THEN empty string is returned and HasValue() is false

### Requirement: Int32Value Type
The system SHALL provide a typed wrapper for 32-bit integer attributes.

#### Scenario: Create int32 value
- GIVEN an integer 42
- WHEN NewInt32Value(42) is called
- THEN an Int32Value with that value is created

#### Scenario: Parse int32 from text
- GIVEN an Int32Value
- WHEN SetInnerText("42") is called
- THEN the value is set to 42

#### Scenario: Invalid int32 text
- GIVEN an Int32Value
- WHEN SetInnerText("not-a-number") is called
- THEN an error is returned

#### Scenario: Int32 to text
- GIVEN an Int32Value with value 42
- WHEN InnerText() is called
- THEN "42" is returned

### Requirement: UInt32Value Type
The system SHALL provide a typed wrapper for unsigned 32-bit integer attributes.

#### Scenario: Create uint32 value
- GIVEN an unsigned integer 4294967295
- WHEN NewUInt32Value(4294967295) is called
- THEN a UInt32Value with that value is created

#### Scenario: Negative value rejected
- GIVEN a UInt32Value
- WHEN SetInnerText("-1") is called
- THEN an error is returned

### Requirement: Int64Value Type
The system SHALL provide a typed wrapper for 64-bit integer attributes.

#### Scenario: Create int64 value
- GIVEN a large integer
- WHEN NewInt64Value(9223372036854775807) is called
- THEN an Int64Value with that value is created

### Requirement: BooleanValue Type
The system SHALL provide a typed wrapper for boolean attributes.

#### Scenario: Create boolean value
- GIVEN a boolean true
- WHEN NewBooleanValue(true) is called
- THEN a BooleanValue with value true is created

#### Scenario: Parse "true" string
- GIVEN a BooleanValue
- WHEN SetInnerText("true") is called
- THEN the value is set to true

#### Scenario: Parse "1" as true
- GIVEN a BooleanValue
- WHEN SetInnerText("1") is called
- THEN the value is set to true

#### Scenario: Parse "false" string
- GIVEN a BooleanValue
- WHEN SetInnerText("false") is called
- THEN the value is set to false

#### Scenario: Parse "0" as false
- GIVEN a BooleanValue
- WHEN SetInnerText("0") is called
- THEN the value is set to false

#### Scenario: Boolean to text
- GIVEN a BooleanValue with value true
- WHEN InnerText() is called
- THEN "true" or "1" is returned (configurable)

### Requirement: OnOffValue Type
The system SHALL provide a typed wrapper for on/off boolean attributes (Word-style).

#### Scenario: Parse "on" as true
- GIVEN an OnOffValue
- WHEN SetInnerText("on") is called
- THEN the value is true

#### Scenario: Parse "off" as false
- GIVEN an OnOffValue
- WHEN SetInnerText("off") is called
- THEN the value is false

#### Scenario: Default value handling
- GIVEN an OnOffValue with no value
- WHEN used in context where default is "on"
- THEN the default applies

### Requirement: EnumValue Generic Type
The system SHALL provide a generic typed wrapper for enumeration attributes.

#### Scenario: Create enum value
- GIVEN an enum type JustificationValues
- WHEN NewEnumValue(JustificationValues.Center) is called
- THEN an EnumValue[JustificationValues] is created

#### Scenario: Parse enum from text
- GIVEN an EnumValue[JustificationValues]
- WHEN SetInnerText("center") is called
- THEN the value is set to JustificationValues.Center

#### Scenario: Invalid enum text
- GIVEN an EnumValue[JustificationValues]
- WHEN SetInnerText("invalid") is called
- THEN an error is returned

#### Scenario: Enum to text
- GIVEN an EnumValue with value Center
- WHEN InnerText() is called
- THEN "center" is returned

### Requirement: HexBinaryValue Type
The system SHALL provide a typed wrapper for hexadecimal binary attributes.

#### Scenario: Create hex value
- GIVEN bytes []byte{0xFF, 0x00, 0xAB}
- WHEN NewHexBinaryValue(bytes) is called
- THEN a HexBinaryValue is created

#### Scenario: Parse hex string
- GIVEN a HexBinaryValue
- WHEN SetInnerText("FF00AB") is called
- THEN the value is set to []byte{0xFF, 0x00, 0xAB}

#### Scenario: Hex to text
- GIVEN a HexBinaryValue with bytes
- WHEN InnerText() is called
- THEN uppercase hex string is returned

### Requirement: Base64BinaryValue Type
The system SHALL provide a typed wrapper for base64 binary attributes.

#### Scenario: Create base64 value
- GIVEN bytes to encode
- WHEN NewBase64BinaryValue(bytes) is called
- THEN a Base64BinaryValue is created

#### Scenario: Parse base64 string
- GIVEN a Base64BinaryValue
- WHEN SetInnerText("SGVsbG8=") is called
- THEN the value is decoded to "Hello" bytes

### Requirement: DateTimeValue Type
The system SHALL provide a typed wrapper for date/time attributes.

#### Scenario: Create datetime value
- GIVEN a time.Time value
- WHEN NewDateTimeValue(t) is called
- THEN a DateTimeValue is created

#### Scenario: Parse ISO 8601 datetime
- GIVEN a DateTimeValue
- WHEN SetInnerText("2024-01-15T10:30:00Z") is called
- THEN the time.Time value is parsed correctly

#### Scenario: DateTime to text
- GIVEN a DateTimeValue
- WHEN InnerText() is called
- THEN ISO 8601 formatted string is returned

### Requirement: DecimalValue Type
The system SHALL provide a typed wrapper for decimal number attributes.

#### Scenario: Create decimal value
- GIVEN a float64 value 3.14159
- WHEN NewDecimalValue(3.14159) is called
- THEN a DecimalValue is created

#### Scenario: Parse decimal string
- GIVEN a DecimalValue
- WHEN SetInnerText("3.14159") is called
- THEN the value is parsed correctly

### Requirement: Twips and Points Values
The system SHALL provide unit-aware value types for measurements.

#### Scenario: TwipsValue (twentieths of a point)
- GIVEN a TwipsValue
- WHEN SetValue(1440) is called (1 inch = 1440 twips)
- THEN the value represents 1 inch

#### Scenario: HalfPointsValue
- GIVEN a HalfPointsValue with value 24
- WHEN ToPoints() is called
- THEN 12.0 points is returned

#### Scenario: EmuValue (English Metric Units)
- GIVEN an EmuValue
- WHEN SetValue(914400) is called
- THEN the value represents 1 inch (914400 EMU = 1 inch)

### Requirement: Percentage Values
The system SHALL provide percentage value types.

#### Scenario: PercentageValue
- GIVEN a PercentageValue
- WHEN SetInnerText("50%") or SetInnerText("50000") is called
- THEN 50% is represented (50000 = 50% in OOXML)

#### Scenario: Convert to float
- GIVEN a PercentageValue representing 50%
- WHEN ToFloat() is called
- THEN 0.5 is returned

