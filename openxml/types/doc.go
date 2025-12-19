// Package types provides simple value types for Office Open XML attributes.
//
// These types handle the conversion between XML string representations and
// Go native types, providing type-safe attribute handling throughout the SDK.
//
// # SimpleValue Interface
//
// All value types implement the SimpleValue interface:
//
//	type SimpleValue interface {
//		HasValue() bool           // Check if a value is set
//		InnerText() string        // Get XML string representation
//		SetInnerText(string) error // Parse from XML string
//	}
//
// Most types also implement Resettable:
//
//	type Resettable interface {
//		SetNil() // Clear to unset state
//	}
//
// # String and Numeric Types
//
// StringValue wraps string values:
//
//	sv := types.NewStringValue("hello")
//	text := sv.InnerText() // "hello"
//	sv.SetNil()
//	sv.HasValue() // false
//
// Int32Value and Int64Value handle integers:
//
//	iv := types.NewInt32Value(42)
//	text := iv.InnerText() // "42"
//	iv.SetValue(100)
//	val := iv.Value() // 100
//
// DecimalValue handles floating-point:
//
//	dv := types.NewDecimalValue(3.14)
//	text := dv.InnerText() // "3.14"
//
// # Boolean Types
//
// BooleanValue handles standard XML boolean values:
//
//	bv := types.NewBooleanValue(true)
//	text := bv.InnerText() // "true"
//	bv.SetInnerText("1")   // Also accepts "1" and "0"
//	bv.Value() // true
//
// OnOffValue handles Word-style on/off values:
//
//	ov := types.NewOnOffValue(true)
//	text := ov.InnerText() // "on"
//	ov.SetInnerText("off")
//	ov.SetInnerText("true")  // Also accepts "true"/"false"
//	ov.SetInnerText("1")     // Also accepts "1"/"0"
//
// # Binary Types
//
// HexBinaryValue handles hex-encoded binary data:
//
//	hv := types.NewHexBinaryValue([]byte{0xFF, 0x00})
//	text := hv.InnerText() // "FF00"
//	hv.SetInnerText("0A1B2C")
//	data := hv.Value() // []byte{0x0A, 0x1B, 0x2C}
//
// Base64BinaryValue handles base64-encoded binary data:
//
//	bv := types.NewBase64BinaryValue([]byte("hello"))
//	text := bv.InnerText() // "aGVsbG8="
//
// # Date/Time Types
//
// DateTimeValue wraps time.Time with ISO 8601 formatting:
//
//	dt := types.NewDateTimeValue(time.Now())
//	text := dt.InnerText() // "2024-01-15T10:30:00Z"
//
// # Enumeration Types
//
// EnumValue[T] provides type-safe enumeration handling:
//
//	type Alignment string
//	const (
//		AlignLeft   Alignment = "left"
//		AlignCenter Alignment = "center"
//		AlignRight  Alignment = "right"
//	)
//
//	ev := types.NewEnumValue(AlignCenter)
//	text := ev.InnerText() // "center"
//	ev.SetInnerText("right")
//	ev.Value() // AlignRight
//
// # Measurement Units
//
// The package provides unit-aware measurement types commonly used in OOXML:
//
// TwipsValue (1 inch = 1440 twips):
//
//	tv := types.NewTwipsValue(720) // 0.5 inch
//	inches := tv.Inches() // 0.5
//	tv.SetInches(1.0)
//	tv.Value() // 1440
//
// HalfPointsValue (font sizes, 1 point = 2 half-points):
//
//	hp := types.NewHalfPointsValue(24) // 12pt
//	points := hp.Points() // 12
//	hp.SetPoints(14)
//	hp.Value() // 28
//
// EmuValue (1 inch = 914400 EMU):
//
//	ev := types.NewEmuValue(914400) // 1 inch
//	inches := ev.Inches() // 1.0
//	cm := ev.Centimeters() // 2.54
//	ev.SetInches(2.0)
//
// PercentageValue (1000-based, 50000 = 50%):
//
//	pv := types.NewPercentageValue(50000) // 50%
//	pct := pv.Percent() // 50.0
//	pv.SetPercent(75.5)
//	pv.Value() // 75500
//
// # Conversion Utilities
//
// The package provides utility functions for common conversions:
//
//	twips := types.InchesToTwips(1.0)    // 1440
//	inches := types.TwipsToInches(720)   // 0.5
//	emus := types.InchesToEmus(1.0)      // 914400
//	halfPts := types.PointsToHalfPoints(12) // 24
package types
