package types

import (
	"bytes"
	"testing"
	"time"
)

func TestPackageExists(t *testing.T) {
	// Simple test to verify test infrastructure works for types package
	t.Log(
		"types package test infrastructure is working",
	)
}

// TestStringValue tests the StringValue type.
func TestStringValue(t *testing.T) {
	t.Run("NewStringValue", func(t *testing.T) {
		sv := NewStringValue("hello")
		if !sv.HasValue() {
			t.Error("HasValue should be true")
		}
		if sv.Value() != "hello" {
			t.Errorf(
				"Value = %q, want %q",
				sv.Value(),
				"hello",
			)
		}
		if sv.InnerText() != "hello" {
			t.Errorf(
				"InnerText = %q, want %q",
				sv.InnerText(),
				"hello",
			)
		}
	})

	t.Run("NilStringValue", func(t *testing.T) {
		sv := NewNilStringValue()
		if sv.HasValue() {
			t.Error("HasValue should be false")
		}
		if sv.Value() != "" {
			t.Errorf(
				"Value = %q, want empty",
				sv.Value(),
			)
		}
	})

	t.Run("SetValue", func(t *testing.T) {
		sv := NewNilStringValue()
		sv.SetValue("world")
		if !sv.HasValue() {
			t.Error(
				"HasValue should be true after SetValue",
			)
		}
		if sv.Value() != "world" {
			t.Errorf(
				"Value = %q, want %q",
				sv.Value(),
				"world",
			)
		}
	})

	t.Run("SetNil", func(t *testing.T) {
		sv := NewStringValue("test")
		sv.SetNil()
		if sv.HasValue() {
			t.Error(
				"HasValue should be false after SetNil",
			)
		}
	})

	t.Run("SetInnerText", func(t *testing.T) {
		sv := NewNilStringValue()
		if err := sv.SetInnerText("parsed"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if sv.Value() != "parsed" {
			t.Errorf(
				"Value = %q, want %q",
				sv.Value(),
				"parsed",
			)
		}
	})
}

// TestInt32Value tests the Int32Value type.
func TestInt32Value(t *testing.T) {
	t.Run("NewInt32Value", func(t *testing.T) {
		iv := NewInt32Value(42)
		if !iv.HasValue() {
			t.Error("HasValue should be true")
		}
		if iv.Value() != 42 {
			t.Errorf(
				"Value = %d, want %d",
				iv.Value(),
				42,
			)
		}
		if iv.InnerText() != "42" {
			t.Errorf(
				"InnerText = %q, want %q",
				iv.InnerText(),
				"42",
			)
		}
	})

	t.Run("ParseValid", func(t *testing.T) {
		iv := NewNilInt32Value()
		if err := iv.SetInnerText("42"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if iv.Value() != 42 {
			t.Errorf(
				"Value = %d, want %d",
				iv.Value(),
				42,
			)
		}
	})

	t.Run("ParseInvalid", func(t *testing.T) {
		iv := NewNilInt32Value()
		err := iv.SetInnerText("not-a-number")
		if err == nil {
			t.Error(
				"Expected error for invalid input",
			)
		}
	})

	t.Run("NegativeValue", func(t *testing.T) {
		iv := NewInt32Value(-100)
		if iv.Value() != -100 {
			t.Errorf(
				"Value = %d, want %d",
				iv.Value(),
				-100,
			)
		}
		if iv.InnerText() != "-100" {
			t.Errorf(
				"InnerText = %q, want %q",
				iv.InnerText(),
				"-100",
			)
		}
	})
}

// TestUInt32Value tests the UInt32Value type.
func TestUInt32Value(t *testing.T) {
	t.Run("NewUInt32Value", func(t *testing.T) {
		uv := NewUInt32Value(4294967295)
		if !uv.HasValue() {
			t.Error("HasValue should be true")
		}
		if uv.Value() != 4294967295 {
			t.Errorf(
				"Value = %d, want %d",
				uv.Value(),
				4294967295,
			)
		}
	})

	t.Run("RejectNegative", func(t *testing.T) {
		uv := NewNilUInt32Value()
		err := uv.SetInnerText("-1")
		if err == nil {
			t.Error(
				"Expected error for negative value",
			)
		}
	})
}

// TestInt64Value tests the Int64Value type.
func TestInt64Value(t *testing.T) {
	t.Run("LargeValue", func(t *testing.T) {
		iv := NewInt64Value(9223372036854775807)
		if iv.Value() != 9223372036854775807 {
			t.Errorf(
				"Value = %d, want max int64",
				iv.Value(),
			)
		}
	})
}

// TestBooleanValue tests the BooleanValue type.
func TestBooleanValue(t *testing.T) {
	t.Run("NewBooleanValue", func(t *testing.T) {
		bv := NewBooleanValue(true)
		if !bv.HasValue() {
			t.Error("HasValue should be true")
		}
		if !bv.Value() {
			t.Error("Value should be true")
		}
	})

	t.Run("ParseTrue", func(t *testing.T) {
		bv := NewNilBooleanValue()
		if err := bv.SetInnerText("true"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if !bv.Value() {
			t.Error("Value should be true")
		}
	})

	t.Run("Parse1", func(t *testing.T) {
		bv := NewNilBooleanValue()
		if err := bv.SetInnerText("1"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if !bv.Value() {
			t.Error("Value should be true")
		}
	})

	t.Run("ParseFalse", func(t *testing.T) {
		bv := NewNilBooleanValue()
		if err := bv.SetInnerText("false"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if bv.Value() {
			t.Error("Value should be false")
		}
	})

	t.Run("Parse0", func(t *testing.T) {
		bv := NewNilBooleanValue()
		if err := bv.SetInnerText("0"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if bv.Value() {
			t.Error("Value should be false")
		}
	})

	t.Run(
		"OutputFormatTrueFalse",
		func(t *testing.T) {
			bv := NewBooleanValue(true)
			if bv.InnerText() != "true" {
				t.Errorf(
					"InnerText = %q, want %q",
					bv.InnerText(),
					"true",
				)
			}
		},
	)

	t.Run(
		"OutputFormatOneZero",
		func(t *testing.T) {
			bv := NewBooleanValueWithFormat(
				true,
				BooleanFormatOneZero,
			)
			if bv.InnerText() != "1" {
				t.Errorf(
					"InnerText = %q, want %q",
					bv.InnerText(),
					"1",
				)
			}
		},
	)

	t.Run("InvalidInput", func(t *testing.T) {
		bv := NewNilBooleanValue()
		err := bv.SetInnerText("invalid")
		if err == nil {
			t.Error(
				"Expected error for invalid input",
			)
		}
	})
}

// TestOnOffValue tests the OnOffValue type.
func TestOnOffValue(t *testing.T) {
	t.Run("ParseOn", func(t *testing.T) {
		ov := NewNilOnOffValue()
		if err := ov.SetInnerText("on"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if !ov.Value() {
			t.Error("Value should be true")
		}
	})

	t.Run("ParseOff", func(t *testing.T) {
		ov := NewNilOnOffValue()
		if err := ov.SetInnerText("off"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if ov.Value() {
			t.Error("Value should be false")
		}
	})

	t.Run("ParseTrue", func(t *testing.T) {
		ov := NewNilOnOffValue()
		if err := ov.SetInnerText("true"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if !ov.Value() {
			t.Error("Value should be true")
		}
	})

	t.Run("OutputFormats", func(t *testing.T) {
		ov := NewOnOffValue(true)
		if ov.InnerText() != "on" {
			t.Errorf(
				"InnerText = %q, want %q",
				ov.InnerText(),
				"on",
			)
		}

		ov.SetOutputFormat(OnOffFormatTrueFalse)
		if ov.InnerText() != "true" {
			t.Errorf(
				"InnerText = %q, want %q",
				ov.InnerText(),
				"true",
			)
		}

		ov.SetOutputFormat(OnOffFormatOneZero)
		if ov.InnerText() != "1" {
			t.Errorf(
				"InnerText = %q, want %q",
				ov.InnerText(),
				"1",
			)
		}
	})
}

// TestEnumValue tests the EnumValue generic type.
func TestEnumValue(t *testing.T) {
	type Justification string
	const (
		justLeft   Justification = "left"
		justCenter Justification = "center"
		justRight  Justification = "right"
	)

	t.Run("NewEnumValue", func(t *testing.T) {
		ev := NewEnumValue(justCenter)
		if !ev.HasValue() {
			t.Error("HasValue should be true")
		}
		if ev.Value() != justCenter {
			t.Errorf(
				"Value = %q, want %q",
				ev.Value(),
				justCenter,
			)
		}
		if ev.InnerText() != "center" {
			t.Errorf(
				"InnerText = %q, want %q",
				ev.InnerText(),
				"center",
			)
		}
	})

	t.Run("WithValidation", func(t *testing.T) {
		validValues := map[string]Justification{
			"left":   justLeft,
			"center": justCenter,
			"right":  justRight,
		}
		ev := NewEnumValueWithValidation(
			justLeft,
			validValues,
		)

		if err := ev.SetInnerText("center"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if ev.Value() != justCenter {
			t.Errorf(
				"Value = %q, want %q",
				ev.Value(),
				justCenter,
			)
		}

		err := ev.SetInnerText("invalid")
		if err == nil {
			t.Error(
				"Expected error for invalid enum value",
			)
		}
	})
}

// TestHexBinaryValue tests the HexBinaryValue type.
func TestHexBinaryValue(t *testing.T) {
	t.Run(
		"NewHexBinaryValue",
		func(t *testing.T) {
			hv := NewHexBinaryValue(
				[]byte{0xFF, 0x00, 0xAB},
			)
			if !hv.HasValue() {
				t.Error("HasValue should be true")
			}
			if hv.InnerText() != "FF00AB" {
				t.Errorf(
					"InnerText = %q, want %q",
					hv.InnerText(),
					"FF00AB",
				)
			}
		},
	)

	t.Run("ParseHex", func(t *testing.T) {
		hv := NewNilHexBinaryValue()
		if err := hv.SetInnerText("FF00AB"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		expected := []byte{0xFF, 0x00, 0xAB}
		if !bytes.Equal(hv.Value(), expected) {
			t.Errorf(
				"Value = %v, want %v",
				hv.Value(),
				expected,
			)
		}
	})

	t.Run("ParseLowercase", func(t *testing.T) {
		hv := NewNilHexBinaryValue()
		if err := hv.SetInnerText("ff00ab"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if hv.Len() != 3 {
			t.Errorf(
				"Len = %d, want %d",
				hv.Len(),
				3,
			)
		}
	})

	t.Run("InvalidHex", func(t *testing.T) {
		hv := NewNilHexBinaryValue()
		err := hv.SetInnerText("not-hex")
		if err == nil {
			t.Error(
				"Expected error for invalid hex",
			)
		}
	})
}

// TestBase64BinaryValue tests the Base64BinaryValue type.
func TestBase64BinaryValue(t *testing.T) {
	t.Run(
		"NewBase64BinaryValue",
		func(t *testing.T) {
			bv := NewBase64BinaryValue(
				[]byte("Hello"),
			)
			if !bv.HasValue() {
				t.Error("HasValue should be true")
			}
			if bv.InnerText() != "SGVsbG8=" {
				t.Errorf(
					"InnerText = %q, want %q",
					bv.InnerText(),
					"SGVsbG8=",
				)
			}
		},
	)

	t.Run("ParseBase64", func(t *testing.T) {
		bv := NewNilBase64BinaryValue()
		if err := bv.SetInnerText("SGVsbG8="); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if string(bv.Value()) != "Hello" {
			t.Errorf(
				"Value = %q, want %q",
				string(bv.Value()),
				"Hello",
			)
		}
	})

	t.Run("InvalidBase64", func(t *testing.T) {
		bv := NewNilBase64BinaryValue()
		err := bv.SetInnerText(
			"not-valid-base64!!!",
		)
		if err == nil {
			t.Error(
				"Expected error for invalid base64",
			)
		}
	})
}

// TestDateTimeValue tests the DateTimeValue type.
func TestDateTimeValue(t *testing.T) {
	t.Run("NewDateTimeValue", func(t *testing.T) {
		tm := time.Date(
			2024,
			1,
			15,
			10,
			30,
			0,
			0,
			time.UTC,
		)
		dv := NewDateTimeValue(tm)
		if !dv.HasValue() {
			t.Error("HasValue should be true")
		}
	})

	t.Run("ParseISO8601", func(t *testing.T) {
		dv := NewNilDateTimeValue()
		if err := dv.SetInnerText("2024-01-15T10:30:00Z"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if dv.Value().Year() != 2024 {
			t.Errorf(
				"Year = %d, want 2024",
				dv.Value().Year(),
			)
		}
		if dv.Value().Month() != time.January {
			t.Errorf(
				"Month = %v, want January",
				dv.Value().Month(),
			)
		}
		if dv.Value().Day() != 15 {
			t.Errorf(
				"Day = %d, want 15",
				dv.Value().Day(),
			)
		}
	})

	t.Run("ParseDateOnly", func(t *testing.T) {
		dv := NewNilDateTimeValue()
		if err := dv.SetInnerText("2024-01-15"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if dv.Value().Year() != 2024 {
			t.Errorf(
				"Year = %d, want 2024",
				dv.Value().Year(),
			)
		}
	})

	t.Run("InvalidDateTime", func(t *testing.T) {
		dv := NewNilDateTimeValue()
		err := dv.SetInnerText("not-a-date")
		if err == nil {
			t.Error(
				"Expected error for invalid datetime",
			)
		}
	})
}

// TestDecimalValue tests the DecimalValue type.
func TestDecimalValue(t *testing.T) {
	t.Run("NewDecimalValue", func(t *testing.T) {
		dv := NewDecimalValue(3.14159)
		if !dv.HasValue() {
			t.Error("HasValue should be true")
		}
		if dv.Value() != 3.14159 {
			t.Errorf(
				"Value = %v, want 3.14159",
				dv.Value(),
			)
		}
	})

	t.Run("ParseDecimal", func(t *testing.T) {
		dv := NewNilDecimalValue()
		if err := dv.SetInnerText("3.14159"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if dv.Value() != 3.14159 {
			t.Errorf(
				"Value = %v, want 3.14159",
				dv.Value(),
			)
		}
	})

	t.Run("InvalidDecimal", func(t *testing.T) {
		dv := NewNilDecimalValue()
		err := dv.SetInnerText("not-a-number")
		if err == nil {
			t.Error(
				"Expected error for invalid decimal",
			)
		}
	})
}

// TestTwipsValue tests the TwipsValue type.
func TestTwipsValue(t *testing.T) {
	t.Run("NewTwipsValue", func(t *testing.T) {
		tv := NewTwipsValue(1440)
		if !tv.HasValue() {
			t.Error("HasValue should be true")
		}
		if tv.Value() != 1440 {
			t.Errorf(
				"Value = %d, want 1440",
				tv.Value(),
			)
		}
	})

	t.Run("ToPoints", func(t *testing.T) {
		tv := NewTwipsValue(
			20,
		) // 1 point = 20 twips
		if tv.ToPoints() != 1.0 {
			t.Errorf(
				"ToPoints = %v, want 1.0",
				tv.ToPoints(),
			)
		}
	})

	t.Run("ToInches", func(t *testing.T) {
		tv := NewTwipsValue(
			1440,
		) // 1 inch = 1440 twips
		if tv.ToInches() != 1.0 {
			t.Errorf(
				"ToInches = %v, want 1.0",
				tv.ToInches(),
			)
		}
	})

	t.Run("FromPoints", func(t *testing.T) {
		tv := NewTwipsValueFromPoints(
			72,
		) // 72 points = 1 inch = 1440 twips
		if tv.Value() != 1440 {
			t.Errorf(
				"Value = %d, want 1440",
				tv.Value(),
			)
		}
	})

	t.Run("FromInches", func(t *testing.T) {
		tv := NewTwipsValueFromInches(1)
		if tv.Value() != 1440 {
			t.Errorf(
				"Value = %d, want 1440",
				tv.Value(),
			)
		}
	})
}

// TestHalfPointsValue tests the HalfPointsValue type.
func TestHalfPointsValue(t *testing.T) {
	t.Run("ToPoints", func(t *testing.T) {
		hv := NewHalfPointsValue(
			24,
		) // 24 half-points = 12 points
		if hv.ToPoints() != 12.0 {
			t.Errorf(
				"ToPoints = %v, want 12.0",
				hv.ToPoints(),
			)
		}
	})

	t.Run("FromPoints", func(t *testing.T) {
		hv := NewHalfPointsValueFromPoints(
			12,
		) // 12 points = 24 half-points
		if hv.Value() != 24 {
			t.Errorf(
				"Value = %d, want 24",
				hv.Value(),
			)
		}
	})
}

// TestEmuValue tests the EmuValue type.
func TestEmuValue(t *testing.T) {
	t.Run("ToInches", func(t *testing.T) {
		ev := NewEmuValue(
			914400,
		) // 914400 EMU = 1 inch
		if ev.ToInches() != 1.0 {
			t.Errorf(
				"ToInches = %v, want 1.0",
				ev.ToInches(),
			)
		}
	})

	t.Run("ToPoints", func(t *testing.T) {
		ev := NewEmuValue(
			12700,
		) // 12700 EMU = 1 point
		if ev.ToPoints() != 1.0 {
			t.Errorf(
				"ToPoints = %v, want 1.0",
				ev.ToPoints(),
			)
		}
	})

	t.Run("FromInches", func(t *testing.T) {
		ev := NewEmuValueFromInches(1)
		if ev.Value() != 914400 {
			t.Errorf(
				"Value = %d, want 914400",
				ev.Value(),
			)
		}
	})

	t.Run("FromCentimeters", func(t *testing.T) {
		ev := NewEmuValueFromCentimeters(1)
		if ev.Value() != 360000 {
			t.Errorf(
				"Value = %d, want 360000",
				ev.Value(),
			)
		}
	})
}

// TestPercentageValue tests the PercentageValue type.
func TestPercentageValue(t *testing.T) {
	t.Run(
		"NewPercentageValue",
		func(t *testing.T) {
			pv := NewPercentageValue(
				50000,
			) // 50000 = 50%
			if !pv.HasValue() {
				t.Error("HasValue should be true")
			}
			if pv.Value() != 50000 {
				t.Errorf(
					"Value = %d, want 50000",
					pv.Value(),
				)
			}
		},
	)

	t.Run("ToFloat", func(t *testing.T) {
		pv := NewPercentageValue(50000) // 50%
		if pv.ToFloat() != 0.5 {
			t.Errorf(
				"ToFloat = %v, want 0.5",
				pv.ToFloat(),
			)
		}
	})

	t.Run("ToPercent", func(t *testing.T) {
		pv := NewPercentageValue(50000)
		if pv.ToPercent() != 50.0 {
			t.Errorf(
				"ToPercent = %v, want 50.0",
				pv.ToPercent(),
			)
		}
	})

	t.Run("FromPercent", func(t *testing.T) {
		pv := NewPercentageValueFromPercent(50)
		if pv.Value() != 50000 {
			t.Errorf(
				"Value = %d, want 50000",
				pv.Value(),
			)
		}
	})

	t.Run("FromFloat", func(t *testing.T) {
		pv := NewPercentageValueFromFloat(0.5)
		if pv.Value() != 50000 {
			t.Errorf(
				"Value = %d, want 50000",
				pv.Value(),
			)
		}
	})

	t.Run(
		"ParsePercentString",
		func(t *testing.T) {
			pv := NewNilPercentageValue()
			if err := pv.SetInnerText("50%"); err != nil {
				t.Errorf(
					"SetInnerText error: %v",
					err,
				)
			}
			if pv.Value() != 50000 {
				t.Errorf(
					"Value = %d, want 50000",
					pv.Value(),
				)
			}
		},
	)

	t.Run("ParseRawInteger", func(t *testing.T) {
		pv := NewNilPercentageValue()
		if err := pv.SetInnerText("50000"); err != nil {
			t.Errorf(
				"SetInnerText error: %v",
				err,
			)
		}
		if pv.Value() != 50000 {
			t.Errorf(
				"Value = %d, want 50000",
				pv.Value(),
			)
		}
	})
}

// TestSimpleValueInterface verifies all types implement SimpleValue.
func TestSimpleValueInterface(t *testing.T) {
	// This test ensures all types implement the SimpleValue interface
	// by using them through the interface
	values := []SimpleValue{
		NewStringValue("test"),
		NewInt32Value(42),
		NewUInt32Value(42),
		NewInt64Value(42),
		NewUInt64Value(42),
		NewBooleanValue(true),
		NewOnOffValue(true),
		NewHexBinaryValue([]byte{0xFF}),
		NewBase64BinaryValue([]byte("test")),
		NewDateTimeValue(time.Now()),
		NewDecimalValue(3.14),
		NewTwipsValue(1440),
		NewHalfPointsValue(24),
		NewEmuValue(914400),
		NewPercentageValue(50000),
	}

	for i, v := range values {
		if !v.HasValue() {
			t.Errorf(
				"values[%d].HasValue() = false, want true",
				i,
			)
		}
		if v.InnerText() == "" {
			t.Errorf(
				"values[%d].InnerText() is empty",
				i,
			)
		}
	}
}

// TestResettableInterface verifies all types implement Resettable.
func TestResettableInterface(t *testing.T) {
	resettables := []Resettable{
		NewStringValue("test"),
		NewInt32Value(42),
		NewUInt32Value(42),
		NewInt64Value(42),
		NewUInt64Value(42),
		NewBooleanValue(true),
		NewOnOffValue(true),
		NewHexBinaryValue([]byte{0xFF}),
		NewBase64BinaryValue([]byte("test")),
		NewDateTimeValue(time.Now()),
		NewDecimalValue(3.14),
		NewTwipsValue(1440),
		NewHalfPointsValue(24),
		NewEmuValue(914400),
		NewPercentageValue(50000),
	}

	for i, r := range resettables {
		r.SetNil()
		sv, ok := r.(SimpleValue)
		if !ok {
			t.Errorf(
				"resettables[%d] does not implement SimpleValue",
				i,
			)

			continue
		}
		if sv.HasValue() {
			t.Errorf(
				"resettables[%d].HasValue() = true after SetNil, want false",
				i,
			)
		}
	}
}
