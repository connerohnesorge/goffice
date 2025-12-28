package elements

//revive:disable:file-length-limit many validation properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// DataValidations represents the data validations container element
// (x:dataValidations). It contains all data validation rules for a worksheet.
type DataValidations struct {
	*openxml.CompositeElementBase
}

// NewDataValidations creates a new DataValidations element.
func NewDataValidations() *DataValidations {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"dataValidations",
		PrefixDefault,
	)

	return &DataValidations{
		CompositeElementBase: elem,
	}
}

// Count returns the count of data validations. Attribute: count.
func (dv *DataValidations) Count() int {
	attr, found := dv.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCount sets the count attribute.
func (dv *DataValidations) SetCount(count int) {
	if count == 0 {
		dv.RemoveAttribute("count", "")

		return
	}
	dv.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.Itoa(count),
		),
	)
}

// updateCount updates the count attribute based on child elements.
func (dv *DataValidations) updateCount() {
	count := 0
	for range dv.GetDataValidations() {
		count++
	}
	dv.SetCount(count)
}

// DisablePrompts returns whether input prompts are disabled.
// Attribute: disablePrompts.
func (dv *DataValidations) DisablePrompts() bool {
	attr, found := dv.GetAttribute(
		"disablePrompts",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDisablePrompts sets whether input prompts are disabled. Attribute: disablePrompts.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (dv *DataValidations) SetDisablePrompts(
	value bool,
) {
	if value {
		dv.SetAttribute(
			openxml.NewAttribute(
				"",
				"disablePrompts",
				"",
				attrValueTrue,
			),
		)
	} else {
		dv.RemoveAttribute("disablePrompts", "")
	}
}

// XWindow returns the x coordinate for input prompt. Attribute: xWindow.
func (dv *DataValidations) XWindow() int {
	attr, found := dv.GetAttribute("xWindow", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetXWindow sets the x coordinate for input prompt. Attribute: xWindow.
func (dv *DataValidations) SetXWindow(value int) {
	if value == 0 {
		dv.RemoveAttribute("xWindow", "")

		return
	}
	dv.SetAttribute(
		openxml.NewAttribute(
			"",
			"xWindow",
			"",
			strconv.Itoa(value),
		),
	)
}

// YWindow returns the y coordinate for input prompt. Attribute: yWindow.
func (dv *DataValidations) YWindow() int {
	attr, found := dv.GetAttribute("yWindow", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetYWindow sets the y coordinate for input prompt. Attribute: yWindow.
func (dv *DataValidations) SetYWindow(value int) {
	if value == 0 {
		dv.RemoveAttribute("yWindow", "")

		return
	}
	dv.SetAttribute(
		openxml.NewAttribute(
			"",
			"yWindow",
			"",
			strconv.Itoa(value),
		),
	)
}

// GetDataValidations returns an iterator over all DataValidation elements.
func (dv *DataValidations) GetDataValidations() iter.Seq[*DataValidation] {
	return func(yield func(*DataValidation) bool) {
		for child := range dv.Children() {
			if child.LocalName() != "dataValidation" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var validation *DataValidation
			switch v := child.(type) {
			case *DataValidation:
				validation = v
			case *openxml.CompositeElementBase:
				validation = &DataValidation{CompositeElementBase: v}
			}
			if validation != nil &&
				!yield(validation) {
				return
			}
		}
	}
}

// GetDataValidationBySqref returns the data validation for the given
// cell range, or nil if not found.
func (dv *DataValidations) GetDataValidationBySqref(
	sqref string,
) *DataValidation {
	for validation := range dv.GetDataValidations() {
		if validation.Sqref() == sqref {
			return validation
		}
	}

	return nil
}

// AddDataValidation adds a new DataValidation element.
func (dv *DataValidations) AddDataValidation(
	sqref string,
) *DataValidation {
	validation := NewDataValidation()
	validation.SetSqref(sqref)
	dv.AppendChild(validation)
	dv.updateCount()

	return validation
}

// AddListValidation adds a list validation to the specified cell range.
//
//nolint:revive // enforce-repeated-arg-type-style: API clarity
func (dv *DataValidations) AddListValidation(
	sqref string,
	formula string,
) *DataValidation {
	validation := dv.AddDataValidation(sqref)
	validation.SetType(ValidationTypeList)
	validation.SetAllowBlank(true)
	validation.SetShowDropDown(
		false,
	) // false means show dropdown
	validation.AddFormula1(formula)

	return validation
}

// AddWholeNumberValidation adds a whole number validation.
//
//nolint:revive // enforce-repeated-arg-type-style: API clarity
func (dv *DataValidations) AddWholeNumberValidation(
	sqref string,
	operator ValidationOperator,
	formula1 string,
	formula2 string,
) *DataValidation {
	validation := dv.AddDataValidation(sqref)
	validation.SetType(ValidationTypeWhole)
	validation.SetOperator(operator)
	validation.AddFormula1(formula1)
	if formula2 != "" {
		validation.AddFormula2(formula2)
	}

	return validation
}

// AddDecimalValidation adds a decimal number validation.
//
//nolint:revive // enforce-repeated-arg-type-style: API clarity
func (dv *DataValidations) AddDecimalValidation(
	sqref string,
	operator ValidationOperator,
	formula1 string,
	formula2 string,
) *DataValidation {
	validation := dv.AddDataValidation(sqref)
	validation.SetType(ValidationTypeDecimal)
	validation.SetOperator(operator)
	validation.AddFormula1(formula1)
	if formula2 != "" {
		validation.AddFormula2(formula2)
	}

	return validation
}

// AddDateValidation adds a date validation.
//
//nolint:revive // enforce-repeated-arg-type-style: API design for clarity
func (dv *DataValidations) AddDateValidation(
	sqref string,
	operator ValidationOperator,
	formula1 string,
	formula2 string,
) *DataValidation {
	validation := dv.AddDataValidation(sqref)
	validation.SetType(ValidationTypeDate)
	validation.SetOperator(operator)
	validation.AddFormula1(formula1)
	if formula2 != "" {
		validation.AddFormula2(formula2)
	}

	return validation
}

// AddTimeValidation adds a time validation.
//
//nolint:revive // enforce-repeated-arg-type-style: API design for clarity
func (dv *DataValidations) AddTimeValidation(
	sqref string,
	operator ValidationOperator,
	formula1 string,
	formula2 string,
) *DataValidation {
	validation := dv.AddDataValidation(sqref)
	validation.SetType(ValidationTypeTime)
	validation.SetOperator(operator)
	validation.AddFormula1(formula1)
	if formula2 != "" {
		validation.AddFormula2(formula2)
	}

	return validation
}

// AddTextLengthValidation adds a text length validation.
//
//nolint:revive // enforce-repeated-arg-type-style: API design for clarity
func (dv *DataValidations) AddTextLengthValidation(
	sqref string,
	operator ValidationOperator,
	formula1 string,
	formula2 string,
) *DataValidation {
	validation := dv.AddDataValidation(sqref)
	validation.SetType(ValidationTypeTextLength)
	validation.SetOperator(operator)
	validation.AddFormula1(formula1)
	if formula2 != "" {
		validation.AddFormula2(formula2)
	}

	return validation
}

// AddCustomValidation adds a custom formula validation.
//
//nolint:revive // enforce-repeated-arg-type-style: named params for clarity
func (dv *DataValidations) AddCustomValidation(
	sqref string,
	formula string,
) *DataValidation {
	validation := dv.AddDataValidation(sqref)
	validation.SetType(ValidationTypeCustom)
	validation.AddFormula1(formula)

	return validation
}

// RemoveDataValidation removes the data validation for the given cell range.
func (dv *DataValidations) RemoveDataValidation(
	sqref string,
) bool {
	validation := dv.GetDataValidationBySqref(
		sqref,
	)
	if validation == nil {
		return false
	}
	removed := dv.RemoveChild(validation)
	if removed {
		dv.updateCount()
	}

	return removed
}

// Clone creates a deep copy of this DataValidations element.
func (dv *DataValidations) Clone() openxml.Element {
	cloned := dv.CompositeElementBase.Clone()

	return &DataValidations{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DataValidations element.
func (dv *DataValidations) CloneNode(
	deep bool,
) openxml.Element {
	cloned := dv.CompositeElementBase.CloneNode(
		deep,
	)

	return &DataValidations{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// DataValidation represents a single data validation element
// (x:dataValidation).
//
//nolint:errname // DataValidation is not an error type
type DataValidation struct {
	*openxml.CompositeElementBase
}

// NewDataValidation creates a new DataValidation element.
func NewDataValidation() *DataValidation {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"dataValidation",
		PrefixDefault,
	)

	return &DataValidation{
		CompositeElementBase: elem,
	}
}

// Type returns the validation type. Attribute: type.
func (d *DataValidation) Type() ValidationType {
	attr, found := d.GetAttribute("type", "")
	if !found {
		return ValidationTypeNone
	}

	return ValidationType(attr.Value())
}

// SetType sets the validation type. Attribute: type.
func (d *DataValidation) SetType(
	validationType ValidationType,
) {
	if validationType == "" ||
		validationType == ValidationTypeNone {
		d.RemoveAttribute("type", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			string(validationType),
		),
	)
}

// Operator returns the validation operator. Attribute: operator.
func (d *DataValidation) Operator() ValidationOperator {
	attr, found := d.GetAttribute("operator", "")
	if !found {
		return ValidationOperatorBetween
	}

	return ValidationOperator(attr.Value())
}

// SetOperator sets the validation operator. Attribute: operator.
func (d *DataValidation) SetOperator(
	operator ValidationOperator,
) {
	if operator == "" ||
		operator == ValidationOperatorBetween {
		d.RemoveAttribute("operator", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"operator",
			"",
			string(operator),
		),
	)
}

// ErrorStyle returns the error alert style. Attribute: errorStyle.
func (d *DataValidation) ErrorStyle() ValidationErrorStyle {
	attr, found := d.GetAttribute(
		"errorStyle",
		"",
	)
	if !found {
		return ValidationErrorStyleStop
	}

	return ValidationErrorStyle(attr.Value())
}

// SetErrorStyle sets the error alert style. Attribute: errorStyle.
func (d *DataValidation) SetErrorStyle(
	style ValidationErrorStyle,
) {
	if style == "" ||
		style == ValidationErrorStyleStop {
		d.RemoveAttribute("errorStyle", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"errorStyle",
			"",
			string(style),
		),
	)
}

// ImeMode returns the IME mode. Attribute: imeMode.
func (d *DataValidation) ImeMode() ValidationImeMode {
	attr, found := d.GetAttribute("imeMode", "")
	if !found {
		return ValidationImeModeNoControl
	}

	return ValidationImeMode(attr.Value())
}

// SetImeMode sets the IME mode. Attribute: imeMode.
func (d *DataValidation) SetImeMode(
	mode ValidationImeMode,
) {
	if mode == "" ||
		mode == ValidationImeModeNoControl {
		d.RemoveAttribute("imeMode", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"imeMode",
			"",
			string(mode),
		),
	)
}

// AllowBlank returns whether blank values are allowed. Attribute: allowBlank.
func (d *DataValidation) AllowBlank() bool {
	attr, found := d.GetAttribute(
		"allowBlank",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAllowBlank sets whether blank values are allowed. Attribute: allowBlank.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataValidation) SetAllowBlank(
	value bool,
) {
	if value {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"allowBlank",
				"",
				attrValueTrue,
			),
		)
	} else {
		d.RemoveAttribute("allowBlank", "")
	}
}

// ShowDropDown returns whether to hide the dropdown (inverted logic).
// When true, the dropdown is hidden. Attribute: showDropDown.
func (d *DataValidation) ShowDropDown() bool {
	attr, found := d.GetAttribute(
		"showDropDown",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowDropDown sets whether to hide the dropdown (inverted logic).
// When true, the dropdown is hidden. Attribute: showDropDown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataValidation) SetShowDropDown(
	value bool,
) {
	if value {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"showDropDown",
				"",
				attrValueTrue,
			),
		)
	} else {
		d.RemoveAttribute("showDropDown", "")
	}
}

// ShowInputMessage returns whether to show input message.
// Attribute: showInputMessage.
func (d *DataValidation) ShowInputMessage() bool {
	attr, found := d.GetAttribute(
		"showInputMessage",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowInputMessage sets whether to show input message. Attribute: showInputMessage.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataValidation) SetShowInputMessage(
	value bool,
) {
	if value {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"showInputMessage",
				"",
				attrValueTrue,
			),
		)
	} else {
		d.RemoveAttribute("showInputMessage", "")
	}
}

// ShowErrorMessage returns whether to show error message.
// Attribute: showErrorMessage.
func (d *DataValidation) ShowErrorMessage() bool {
	attr, found := d.GetAttribute(
		"showErrorMessage",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowErrorMessage sets whether to show error message. Attribute: showErrorMessage.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DataValidation) SetShowErrorMessage(
	value bool,
) {
	if value {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"showErrorMessage",
				"",
				attrValueTrue,
			),
		)
	} else {
		d.RemoveAttribute("showErrorMessage", "")
	}
}

// ErrorTitle returns the error alert title. Attribute: errorTitle.
func (d *DataValidation) ErrorTitle() string {
	attr, found := d.GetAttribute(
		"errorTitle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetErrorTitle sets the error alert title. Attribute: errorTitle.
func (d *DataValidation) SetErrorTitle(
	title string,
) {
	if title == "" {
		d.RemoveAttribute("errorTitle", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"errorTitle",
			"",
			title,
		),
	)
}

// Error returns the error alert message. Attribute: error.
func (d *DataValidation) Error() string {
	attr, found := d.GetAttribute("error", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetError sets the error alert message. Attribute: error.
func (d *DataValidation) SetError(
	message string,
) {
	if message == "" {
		d.RemoveAttribute("error", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"error",
			"",
			message,
		),
	)
}

// PromptTitle returns the input prompt title. Attribute: promptTitle.
func (d *DataValidation) PromptTitle() string {
	attr, found := d.GetAttribute(
		"promptTitle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPromptTitle sets the input prompt title. Attribute: promptTitle.
func (d *DataValidation) SetPromptTitle(
	title string,
) {
	if title == "" {
		d.RemoveAttribute("promptTitle", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"promptTitle",
			"",
			title,
		),
	)
}

// Prompt returns the input prompt message. Attribute: prompt.
func (d *DataValidation) Prompt() string {
	attr, found := d.GetAttribute("prompt", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPrompt sets the input prompt message. Attribute: prompt.
func (d *DataValidation) SetPrompt(
	message string,
) {
	if message == "" {
		d.RemoveAttribute("prompt", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"prompt",
			"",
			message,
		),
	)
}

// Sqref returns the cell ranges for this validation. Attribute: sqref.
func (d *DataValidation) Sqref() string {
	attr, found := d.GetAttribute("sqref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSqref sets the cell ranges for this validation. Attribute: sqref.
func (d *DataValidation) SetSqref(sqref string) {
	if sqref == "" {
		d.RemoveAttribute("sqref", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"sqref",
			"",
			sqref,
		),
	)
}

// Formula1 returns the first formula child element, or nil if not present.
func (d *DataValidation) Formula1() *ValidationFormula {
	for child := range d.Children() {
		if child.LocalName() != "formula1" ||
			child.NamespaceURI() != NamespaceSML {
			continue
		}
		if f, ok := child.(*ValidationFormula); ok {
			return f
		}
		if leaf, ok := child.(*openxml.LeafElementBase); ok {
			return &ValidationFormula{
				LeafElementBase: leaf,
			}
		}
	}

	return nil
}

// Formula1Text returns the text of the first formula, or empty string.
func (d *DataValidation) Formula1Text() string {
	f := d.Formula1()
	if f == nil {
		return ""
	}

	return f.Formula()
}

// AddFormula1 adds or updates the first formula.
func (d *DataValidation) AddFormula1(
	formula string,
) *ValidationFormula {
	f := d.Formula1()
	if f != nil {
		f.SetFormula(formula)

		return f
	}
	f = NewValidationFormula1()
	f.SetFormula(formula)
	// Insert at the beginning of children
	d.PrependChild(f)

	return f
}

// Formula2 returns the second formula child element, or nil if not present.
func (d *DataValidation) Formula2() *ValidationFormula {
	for child := range d.Children() {
		if child.LocalName() != "formula2" ||
			child.NamespaceURI() != NamespaceSML {
			continue
		}
		if f, ok := child.(*ValidationFormula); ok {
			return f
		}
		if leaf, ok := child.(*openxml.LeafElementBase); ok {
			return &ValidationFormula{
				LeafElementBase: leaf,
			}
		}
	}

	return nil
}

// Formula2Text returns the text of the second formula, or empty string.
func (d *DataValidation) Formula2Text() string {
	f := d.Formula2()
	if f == nil {
		return ""
	}

	return f.Formula()
}

// AddFormula2 adds or updates the second formula.
func (d *DataValidation) AddFormula2(
	formula string,
) *ValidationFormula {
	f := d.Formula2()
	if f != nil {
		f.SetFormula(formula)

		return f
	}
	f = NewValidationFormula2()
	f.SetFormula(formula)
	// Insert after formula1 if it exists, or at beginning
	f1 := d.Formula1()
	if f1 != nil {
		d.InsertAfter(f, f1)
	} else {
		d.PrependChild(f)
	}

	return f
}

// SetInputMessage sets the input message title and text.
func (d *DataValidation) SetInputMessage(
	title, message string,
) {
	d.SetPromptTitle(title)
	d.SetPrompt(message)
	d.SetShowInputMessage(
		title != "" || message != "",
	)
}

// SetErrorAlert sets the error alert title, message, and style.
func (d *DataValidation) SetErrorAlert(
	title, message string,
	style ValidationErrorStyle,
) {
	d.SetErrorTitle(title)
	d.SetError(message)
	d.SetErrorStyle(style)
	d.SetShowErrorMessage(
		title != "" || message != "",
	)
}

// Clone creates a deep copy of this DataValidation element.
func (d *DataValidation) Clone() openxml.Element {
	cloned := d.CompositeElementBase.Clone()

	return &DataValidation{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DataValidation element.
func (d *DataValidation) CloneNode(
	deep bool,
) openxml.Element {
	cloned := d.CompositeElementBase.CloneNode(
		deep,
	)

	return &DataValidation{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ValidationFormula represents a formula element within data validation
// (x:formula1 or x:formula2).
type ValidationFormula struct {
	*openxml.LeafElementBase
}

// NewValidationFormula1 creates a new formula1 element.
func NewValidationFormula1() *ValidationFormula {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"formula1",
		PrefixDefault,
	)

	return &ValidationFormula{
		LeafElementBase: elem,
	}
}

// NewValidationFormula2 creates a new formula2 element.
func NewValidationFormula2() *ValidationFormula {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"formula2",
		PrefixDefault,
	)

	return &ValidationFormula{
		LeafElementBase: elem,
	}
}

// Formula returns the formula text.
func (f *ValidationFormula) Formula() string {
	return f.InnerText()
}

// SetFormula sets the formula text.
func (f *ValidationFormula) SetFormula(
	formula string,
) {
	f.SetInnerText(formula)
}

// Clone creates a deep copy of this ValidationFormula element.
func (f *ValidationFormula) Clone() openxml.Element {
	cloned := f.LeafElementBase.Clone()

	return &ValidationFormula{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this ValidationFormula element.
func (f *ValidationFormula) CloneNode(
	deep bool,
) openxml.Element {
	cloned := f.LeafElementBase.CloneNode(deep)

	return &ValidationFormula{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
