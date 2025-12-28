package elements

import (
	"strings"
	"testing"
)

func TestNewDataValidations(t *testing.T) {
	dv := NewDataValidations()

	if dv == nil {
		t.Fatal(
			"expected non-nil DataValidations",
		)
	}
	if dv.LocalName() != "dataValidations" {
		t.Errorf(
			"expected localName 'dataValidations', got '%s'",
			dv.LocalName(),
		)
	}
	if dv.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceSML,
			dv.NamespaceURI(),
		)
	}
}

func TestDataValidations_Count(t *testing.T) {
	dv := NewDataValidations()

	if dv.Count() != 0 {
		t.Errorf(
			"expected count 0, got %d",
			dv.Count(),
		)
	}

	dv.SetCount(5)
	if dv.Count() != 5 {
		t.Errorf(
			"expected count 5, got %d",
			dv.Count(),
		)
	}
}

func TestDataValidations_DisablePrompts(
	t *testing.T,
) {
	dv := NewDataValidations()

	if dv.DisablePrompts() {
		t.Error(
			"expected disablePrompts to be false initially",
		)
	}

	dv.SetDisablePrompts(true)
	if !dv.DisablePrompts() {
		t.Error(
			"expected disablePrompts to be true",
		)
	}
}

func TestDataValidations_WindowCoordinates(
	t *testing.T,
) {
	dv := NewDataValidations()

	dv.SetXWindow(100)
	dv.SetYWindow(200)

	if dv.XWindow() != 100 {
		t.Errorf(
			"expected xWindow 100, got %d",
			dv.XWindow(),
		)
	}
	if dv.YWindow() != 200 {
		t.Errorf(
			"expected yWindow 200, got %d",
			dv.YWindow(),
		)
	}
}

func TestDataValidations_AddDataValidation(
	t *testing.T,
) {
	dv := NewDataValidations()

	validation := dv.AddDataValidation("A1:A10")
	if validation == nil {
		t.Fatal("expected non-nil DataValidation")
	}
	if validation.Sqref() != testRangeA1A10 {
		t.Errorf(
			"expected sqref 'A1:A10', got '%s'",
			validation.Sqref(),
		)
	}
	if dv.Count() != 1 {
		t.Errorf(
			"expected count 1, got %d",
			dv.Count(),
		)
	}
}

func TestDataValidations_AddListValidation(
	t *testing.T,
) {
	dv := NewDataValidations()

	validation := dv.AddListValidation(
		"B1:B10",
		"\"Option1,Option2,Option3\"",
	)

	if validation.Type() != ValidationTypeList {
		t.Errorf(
			"expected type 'list', got '%s'",
			validation.Type(),
		)
	}
	if !validation.AllowBlank() {
		t.Error("expected allowBlank to be true")
	}
	if validation.Formula1Text() != "\"Option1,Option2,Option3\"" {
		t.Errorf(
			"expected formula1 '\"Option1,Option2,Option3\"', got '%s'",
			validation.Formula1Text(),
		)
	}
}

func TestDataValidations_AddWholeNumberValidation(
	t *testing.T,
) {
	dv := NewDataValidations()

	validation := dv.AddWholeNumberValidation(
		"C1:C10",
		ValidationOperatorBetween,
		"1",
		"100",
	)

	if validation.Type() != ValidationTypeWhole {
		t.Errorf(
			"expected type 'whole', got '%s'",
			validation.Type(),
		)
	}
	if validation.Operator() != ValidationOperatorBetween {
		t.Errorf(
			"expected operator 'between', got '%s'",
			validation.Operator(),
		)
	}
	if validation.Formula1Text() != "1" {
		t.Errorf(
			"expected formula1 '1', got '%s'",
			validation.Formula1Text(),
		)
	}
	if validation.Formula2Text() != "100" {
		t.Errorf(
			"expected formula2 '100', got '%s'",
			validation.Formula2Text(),
		)
	}
}

func TestDataValidations_AddDecimalValidation(
	t *testing.T,
) {
	dv := NewDataValidations()

	validation := dv.AddDecimalValidation(
		"D1:D10",
		ValidationOperatorGreaterThan,
		"0.5",
		"",
	)

	if validation.Type() != ValidationTypeDecimal {
		t.Errorf(
			"expected type 'decimal', got '%s'",
			validation.Type(),
		)
	}
	if validation.Operator() != ValidationOperatorGreaterThan {
		t.Errorf(
			"expected operator 'greaterThan', got '%s'",
			validation.Operator(),
		)
	}
}

func TestDataValidations_AddDateValidation(
	t *testing.T,
) {
	dv := NewDataValidations()

	validation := dv.AddDateValidation(
		"E1:E10",
		ValidationOperatorLessThan,
		"TODAY()",
		"",
	)

	if validation.Type() != ValidationTypeDate {
		t.Errorf(
			"expected type 'date', got '%s'",
			validation.Type(),
		)
	}
}

func TestDataValidations_AddTimeValidation(
	t *testing.T,
) {
	dv := NewDataValidations()

	validation := dv.AddTimeValidation(
		"F1:F10",
		ValidationOperatorBetween,
		"0.333333",
		"0.75",
	)

	if validation.Type() != ValidationTypeTime {
		t.Errorf(
			"expected type 'time', got '%s'",
			validation.Type(),
		)
	}
}

func TestDataValidations_AddTextLengthValidation(
	t *testing.T,
) {
	dv := NewDataValidations()

	validation := dv.AddTextLengthValidation(
		"G1:G10",
		ValidationOperatorLessThanOrEqual,
		"50",
		"",
	)

	if validation.Type() != ValidationTypeTextLength {
		t.Errorf(
			"expected type 'textLength', got '%s'",
			validation.Type(),
		)
	}
}

func TestDataValidations_AddCustomValidation(
	t *testing.T,
) {
	dv := NewDataValidations()

	validation := dv.AddCustomValidation(
		"H1:H10",
		"AND(A1>0,A1<100)",
	)

	if validation.Type() != ValidationTypeCustom {
		t.Errorf(
			"expected type 'custom', got '%s'",
			validation.Type(),
		)
	}
	if validation.Formula1Text() != "AND(A1>0,A1<100)" {
		t.Errorf(
			"expected formula 'AND(A1>0,A1<100)', got '%s'",
			validation.Formula1Text(),
		)
	}
}

func TestDataValidations_RemoveDataValidation(
	t *testing.T,
) {
	dv := NewDataValidations()
	_ = dv.AddDataValidation("A1:A10")
	_ = dv.AddDataValidation("B1:B10")

	if dv.Count() != 2 {
		t.Errorf(
			"expected count 2, got %d",
			dv.Count(),
		)
	}

	removed := dv.RemoveDataValidation("A1:A10")
	if !removed {
		t.Error(
			"expected RemoveDataValidation to return true",
		)
	}
	if dv.Count() != 1 {
		t.Errorf(
			"expected count 1, got %d",
			dv.Count(),
		)
	}

	removed = dv.RemoveDataValidation("A1:A10")
	if removed {
		t.Error(
			"expected RemoveDataValidation to return false for non-existent",
		)
	}
}

func TestDataValidations_GetDataValidationBySqref(
	t *testing.T,
) {
	dv := NewDataValidations()
	_ = dv.AddDataValidation("A1:A10")
	_ = dv.AddDataValidation("B1:B10")

	found := dv.GetDataValidationBySqref("B1:B10")
	if found == nil {
		t.Fatal("expected to find validation")
	}
	if found.Sqref() != "B1:B10" {
		t.Errorf(
			"expected sqref 'B1:B10', got '%s'",
			found.Sqref(),
		)
	}

	notFound := dv.GetDataValidationBySqref(
		"C1:C10",
	)
	if notFound != nil {
		t.Error(
			"expected nil for non-existent sqref",
		)
	}
}

func TestNewDataValidation(t *testing.T) {
	d := NewDataValidation()

	if d == nil {
		t.Fatal("expected non-nil DataValidation")
	}
	if d.LocalName() != "dataValidation" {
		t.Errorf(
			"expected localName 'dataValidation', got '%s'",
			d.LocalName(),
		)
	}
}

func TestDataValidation_Type(t *testing.T) {
	d := NewDataValidation()

	// Default type is none
	if d.Type() != ValidationTypeNone {
		t.Errorf(
			"expected default type 'none', got '%s'",
			d.Type(),
		)
	}

	d.SetType(ValidationTypeList)
	if d.Type() != ValidationTypeList {
		t.Errorf(
			"expected type 'list', got '%s'",
			d.Type(),
		)
	}
}

func TestDataValidation_Operator(t *testing.T) {
	d := NewDataValidation()

	// Default operator is between
	if d.Operator() != ValidationOperatorBetween {
		t.Errorf(
			"expected default operator 'between', got '%s'",
			d.Operator(),
		)
	}

	d.SetOperator(ValidationOperatorGreaterThan)
	if d.Operator() != ValidationOperatorGreaterThan {
		t.Errorf(
			"expected operator 'greaterThan', got '%s'",
			d.Operator(),
		)
	}
}

func TestDataValidation_ErrorStyle(t *testing.T) {
	d := NewDataValidation()

	// Default is stop
	if d.ErrorStyle() != ValidationErrorStyleStop {
		t.Errorf(
			"expected default errorStyle 'stop', got '%s'",
			d.ErrorStyle(),
		)
	}

	d.SetErrorStyle(ValidationErrorStyleWarning)
	if d.ErrorStyle() != ValidationErrorStyleWarning {
		t.Errorf(
			"expected errorStyle 'warning', got '%s'",
			d.ErrorStyle(),
		)
	}

	d.SetErrorStyle(
		ValidationErrorStyleInformation,
	)
	if d.ErrorStyle() != ValidationErrorStyleInformation {
		t.Errorf(
			"expected errorStyle 'information', got '%s'",
			d.ErrorStyle(),
		)
	}
}

func TestDataValidation_ImeMode(t *testing.T) {
	d := NewDataValidation()

	// Default is noControl
	if d.ImeMode() != ValidationImeModeNoControl {
		t.Errorf(
			"expected default imeMode 'noControl', got '%s'",
			d.ImeMode(),
		)
	}

	d.SetImeMode(ValidationImeModeHiragana)
	if d.ImeMode() != ValidationImeModeHiragana {
		t.Errorf(
			"expected imeMode 'hiragana', got '%s'",
			d.ImeMode(),
		)
	}
}

func TestDataValidation_BooleanAttributes(
	t *testing.T,
) {
	d := NewDataValidation()

	// AllowBlank
	if d.AllowBlank() {
		t.Error(
			"expected allowBlank false initially",
		)
	}
	d.SetAllowBlank(true)
	if !d.AllowBlank() {
		t.Error("expected allowBlank true")
	}

	// ShowDropDown
	if d.ShowDropDown() {
		t.Error(
			"expected showDropDown false initially",
		)
	}
	d.SetShowDropDown(true)
	if !d.ShowDropDown() {
		t.Error("expected showDropDown true")
	}

	// ShowInputMessage
	if d.ShowInputMessage() {
		t.Error(
			"expected showInputMessage false initially",
		)
	}
	d.SetShowInputMessage(true)
	if !d.ShowInputMessage() {
		t.Error("expected showInputMessage true")
	}

	// ShowErrorMessage
	if d.ShowErrorMessage() {
		t.Error(
			"expected showErrorMessage false initially",
		)
	}
	d.SetShowErrorMessage(true)
	if !d.ShowErrorMessage() {
		t.Error("expected showErrorMessage true")
	}
}

func TestDataValidation_ErrorAndPromptMessages(
	t *testing.T,
) {
	d := NewDataValidation()

	d.SetErrorTitle("Error")
	d.SetError("Invalid value")
	d.SetPromptTitle("Help")
	d.SetPrompt("Enter a valid value")

	if d.ErrorTitle() != "Error" {
		t.Errorf(
			"expected errorTitle 'Error', got '%s'",
			d.ErrorTitle(),
		)
	}
	if d.Error() != "Invalid value" {
		t.Errorf(
			"expected error 'Invalid value', got '%s'",
			d.Error(),
		)
	}
	if d.PromptTitle() != "Help" {
		t.Errorf(
			"expected promptTitle 'Help', got '%s'",
			d.PromptTitle(),
		)
	}
	if d.Prompt() != "Enter a valid value" {
		t.Errorf(
			"expected prompt 'Enter a valid value', got '%s'",
			d.Prompt(),
		)
	}
}

func TestDataValidation_SetInputMessage(
	t *testing.T,
) {
	d := NewDataValidation()

	d.SetInputMessage("Title", "Message")

	if d.PromptTitle() != "Title" {
		t.Errorf(
			"expected promptTitle 'Title', got '%s'",
			d.PromptTitle(),
		)
	}
	if d.Prompt() != "Message" {
		t.Errorf(
			"expected prompt 'Message', got '%s'",
			d.Prompt(),
		)
	}
	if !d.ShowInputMessage() {
		t.Error(
			"expected showInputMessage to be true",
		)
	}
}

func TestDataValidation_SetErrorAlert(
	t *testing.T,
) {
	d := NewDataValidation()

	d.SetErrorAlert(
		"Error Title",
		"Error Message",
		ValidationErrorStyleWarning,
	)

	if d.ErrorTitle() != "Error Title" {
		t.Errorf(
			"expected errorTitle 'Error Title', got '%s'",
			d.ErrorTitle(),
		)
	}
	if d.Error() != "Error Message" {
		t.Errorf(
			"expected error 'Error Message', got '%s'",
			d.Error(),
		)
	}
	if d.ErrorStyle() != ValidationErrorStyleWarning {
		t.Errorf(
			"expected errorStyle 'warning', got '%s'",
			d.ErrorStyle(),
		)
	}
	if !d.ShowErrorMessage() {
		t.Error(
			"expected showErrorMessage to be true",
		)
	}
}

func TestDataValidation_XMLSerialization(
	t *testing.T,
) {
	dv := NewDataValidations()
	dv.SetCount(1)

	v := dv.AddDataValidation("A1:A10")
	v.SetType(ValidationTypeWhole)
	v.SetOperator(ValidationOperatorBetween)
	v.SetAllowBlank(true)
	v.SetShowInputMessage(true)
	v.SetShowErrorMessage(true)
	v.SetPromptTitle("Input")
	v.SetPrompt(
		"Enter a number between 1 and 100",
	)
	v.SetErrorTitle("Error")
	v.SetError("Value must be between 1 and 100")
	v.AddFormula1("1")
	v.AddFormula2("100")

	xml := dv.OuterXml()

	// Verify key elements are present
	if !strings.Contains(xml, "dataValidations") {
		t.Error(
			"expected 'dataValidations' in XML",
		)
	}
	if !strings.Contains(xml, "dataValidation") {
		t.Error(
			"expected 'dataValidation' in XML",
		)
	}
	if !strings.Contains(xml, "type=\"whole\"") {
		t.Error(
			"expected 'type=\"whole\"' in XML",
		)
	}
	if !strings.Contains(
		xml,
		"sqref=\"A1:A10\"",
	) {
		t.Error(
			"expected 'sqref=\"A1:A10\"' in XML",
		)
	}
	if !strings.Contains(xml, "formula1") {
		t.Error("expected 'formula1' in XML")
	}
	if !strings.Contains(xml, "formula2") {
		t.Error("expected 'formula2' in XML")
	}
}

func TestDataValidation_Clone(t *testing.T) {
	d := NewDataValidation()
	d.SetSqref("A1:A10")
	d.SetType(ValidationTypeList)
	d.AddFormula1("\"Yes,No\"")

	cloned := d.Clone()
	clone, ok := cloned.(*DataValidation)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *DataValidation",
			cloned,
		)
	}

	if clone.Sqref() != "A1:A10" {
		t.Errorf(
			"expected sqref 'A1:A10', got '%s'",
			clone.Sqref(),
		)
	}
	if clone.Type() != ValidationTypeList {
		t.Errorf(
			"expected type 'list', got '%s'",
			clone.Type(),
		)
	}

	// Modify original
	d.SetSqref("B1:B10")
	if clone.Sqref() != "A1:A10" {
		t.Error(
			"expected clone to be independent",
		)
	}
}

func TestValidationTypes(t *testing.T) {
	types := []ValidationType{
		ValidationTypeNone,
		ValidationTypeWhole,
		ValidationTypeDecimal,
		ValidationTypeList,
		ValidationTypeDate,
		ValidationTypeTime,
		ValidationTypeTextLength,
		ValidationTypeCustom,
	}

	for _, vt := range types {
		d := NewDataValidation()
		if vt == ValidationTypeNone {
			continue
		}
		d.SetType(vt)
		if d.Type() != vt {
			t.Errorf(
				"expected type '%s', got '%s'",
				vt,
				d.Type(),
			)
		}
	}
}

func TestValidationOperators(t *testing.T) {
	operators := []ValidationOperator{
		ValidationOperatorBetween,
		ValidationOperatorNotBetween,
		ValidationOperatorEqual,
		ValidationOperatorNotEqual,
		ValidationOperatorLessThan,
		ValidationOperatorLessThanOrEqual,
		ValidationOperatorGreaterThan,
		ValidationOperatorGreaterThanOrEqual,
	}

	for _, op := range operators {
		d := NewDataValidation()
		if op == ValidationOperatorBetween {
			continue
		}
		d.SetOperator(op)
		if d.Operator() != op {
			t.Errorf(
				"expected operator '%s', got '%s'",
				op,
				d.Operator(),
			)
		}
	}
}

func TestValidationErrorStyles(t *testing.T) {
	styles := []ValidationErrorStyle{
		ValidationErrorStyleStop,
		ValidationErrorStyleWarning,
		ValidationErrorStyleInformation,
	}

	for _, style := range styles {
		d := NewDataValidation()
		if style == ValidationErrorStyleStop {
			continue
		}
		d.SetErrorStyle(style)
		if d.ErrorStyle() != style {
			t.Errorf(
				"expected errorStyle '%s', got '%s'",
				style,
				d.ErrorStyle(),
			)
		}
	}
}

func TestValidationImeModes(t *testing.T) {
	modes := []ValidationImeMode{
		ValidationImeModeNoControl,
		ValidationImeModeOff,
		ValidationImeModeOn,
		ValidationImeModeDisabled,
		ValidationImeModeHiragana,
		ValidationImeModeFullKatakana,
		ValidationImeModeHalfKatakana,
		ValidationImeModeFullAlpha,
		ValidationImeModeHalfAlpha,
		ValidationImeModeFullHangul,
		ValidationImeModeHalfHangul,
	}

	for _, mode := range modes {
		d := NewDataValidation()
		//nolint:revive // early-return: test structure pattern
		if mode != ValidationImeModeNoControl {
			d.SetImeMode(mode)
			if d.ImeMode() != mode {
				t.Errorf(
					"expected imeMode '%s', got '%s'",
					mode,
					d.ImeMode(),
				)
			}
		}
	}
}
