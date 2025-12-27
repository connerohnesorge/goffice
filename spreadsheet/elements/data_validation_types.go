package elements

// ValidationType represents the type of data validation.
type ValidationType string

const (
	// ValidationTypeNone indicates no validation.
	ValidationTypeNone ValidationType = "none"
	// ValidationTypeWhole indicates whole number validation.
	ValidationTypeWhole ValidationType = "whole"
	// ValidationTypeDecimal indicates decimal number validation.
	ValidationTypeDecimal ValidationType = "decimal"
	// ValidationTypeList indicates list validation.
	ValidationTypeList ValidationType = "list"
	// ValidationTypeDate indicates date validation.
	ValidationTypeDate ValidationType = "date"
	// ValidationTypeTime indicates time validation.
	ValidationTypeTime ValidationType = "time"
	// ValidationTypeTextLength indicates text length validation.
	ValidationTypeTextLength ValidationType = "textLength"
	// ValidationTypeCustom indicates custom formula validation.
	ValidationTypeCustom ValidationType = "custom"
)

// ValidationOperator represents the comparison operator for data validation.
type ValidationOperator string

const (
	// ValidationOperatorBetween indicates value must be between two values.
	ValidationOperatorBetween ValidationOperator = "between"
	// ValidationOperatorNotBetween indicates value must not be between two
	// values.
	ValidationOperatorNotBetween ValidationOperator = "notBetween"
	// ValidationOperatorEqual indicates value must equal.
	ValidationOperatorEqual ValidationOperator = "equal"
	// ValidationOperatorNotEqual indicates value must not equal.
	ValidationOperatorNotEqual ValidationOperator = "notEqual"
	// ValidationOperatorLessThan indicates value must be less than.
	ValidationOperatorLessThan ValidationOperator = "lessThan"
	// ValidationOperatorLessThanOrEqual indicates value must be
	// less than or equal.
	ValidationOperatorLessThanOrEqual ValidationOperator = "lessThanOrEqual"
	// ValidationOperatorGreaterThan indicates value must be greater than.
	ValidationOperatorGreaterThan ValidationOperator = "greaterThan"
	// ValidationOperatorGreaterThanOrEqual indicates value must be >=.
	//nolint:revive // line-length-limit: long identifier
	ValidationOperatorGreaterThanOrEqual ValidationOperator = "greaterThanOrEqual"
)

// ValidationErrorStyle represents the error alert style for data validation.
type ValidationErrorStyle string

const (
	// ValidationErrorStyleStop prevents invalid data entry.
	ValidationErrorStyleStop ValidationErrorStyle = "stop"
	// ValidationErrorStyleWarning allows invalid data with warning.
	ValidationErrorStyleWarning ValidationErrorStyle = "warning"
	// ValidationErrorStyleInformation shows information message.
	ValidationErrorStyleInformation ValidationErrorStyle = "information"
)

// ValidationImeMode represents the IME mode for data validation.
type ValidationImeMode string

const (
	// ValidationImeModeNoControl indicates no IME control.
	ValidationImeModeNoControl ValidationImeMode = "noControl"
	// ValidationImeModeOff indicates IME off.
	ValidationImeModeOff ValidationImeMode = "off"
	// ValidationImeModeOn indicates IME on.
	ValidationImeModeOn ValidationImeMode = "on"
	// ValidationImeModeDisabled indicates IME disabled.
	ValidationImeModeDisabled ValidationImeMode = "disabled"
	// ValidationImeModeHiragana indicates Hiragana mode.
	ValidationImeModeHiragana ValidationImeMode = "hiragana"
	// ValidationImeModeFullKatakana indicates full Katakana mode.
	ValidationImeModeFullKatakana ValidationImeMode = "fullKatakana"
	// ValidationImeModeHalfKatakana indicates half Katakana mode.
	ValidationImeModeHalfKatakana ValidationImeMode = "halfKatakana"
	// ValidationImeModeFullAlpha indicates full-width alpha mode.
	ValidationImeModeFullAlpha ValidationImeMode = "fullAlpha"
	// ValidationImeModeHalfAlpha indicates half-width alpha mode.
	ValidationImeModeHalfAlpha ValidationImeMode = "halfAlpha"
	// ValidationImeModeFullHangul indicates full Hangul mode.
	ValidationImeModeFullHangul ValidationImeMode = "fullHangul"
	// ValidationImeModeHalfHangul indicates half Hangul mode.
	ValidationImeModeHalfHangul ValidationImeMode = "halfHangul"
)
