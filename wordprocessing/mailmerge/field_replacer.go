// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// replaceFields replaces all merge fields with values from the getValue function.
// The getValue function is called for each field name and should return the replacement value.
func (m *MailMerge) replaceFields(
	fields []*MergeField,
	getValue func(string) (string, error),
) error {
	for _, field := range fields {
		value, err := getValue(field.FieldName)
		if err != nil {
			if m.options.StrictFields {
				return fmt.Errorf(
					"failed to get value for field %q: %w",
					field.FieldName,
					err,
				)
			}
			// In non-strict mode, use empty string for missing fields
			value = ""
		}

		// Apply format switches
		value = m.applyFormatSwitch(
			value,
			field.Switches,
		)

		// Replace the field based on its type
		var replaceErr error
		if field.FieldType == FieldTypeSimple {
			replaceErr = m.replaceSimpleField(
				field,
				value,
			)
		} else {
			replaceErr = m.replaceComplexField(field, value)
		}

		if replaceErr != nil {
			return fmt.Errorf(
				"failed to replace field %q: %w",
				field.FieldName,
				replaceErr,
			)
		}
	}

	return nil
}

// getFieldValue retrieves the value for a merge field from the data source.
// Returns the field value or an error if the field doesn't exist.
func (*MailMerge) getFieldValue(
	field *MergeField,
	ds DataSource,
) (string, error) {
	// Special handling for GREETINGLINE
	if field.FieldName == FieldNameGreetingLine {
		return generateGreetingLine(
			ds,
			field.Switches,
		), nil
	}

	// Get value from data source
	return ds.Get(field.FieldName)
}

// applyFormatSwitch applies formatting switches to the field value.
// Supports: \* Upper, \* Lower, \* FirstCap, \* Caps
func (*MailMerge) applyFormatSwitch(
	value string,
	switches map[string]string,
) string {
	if value == "" {
		return value
	}

	// Check for \* formatting switch
	formatSwitch, ok := switches["\\*"]
	if !ok {
		return value
	}

	formatSwitch = strings.ToUpper(
		strings.TrimSpace(formatSwitch),
	)

	switch formatSwitch {
	case "UPPER":
		return strings.ToUpper(value)
	case "LOWER":
		return strings.ToLower(value)
	case "FIRSTCAP":
		return firstCap(value)
	case "CAPS":
		return titleCase(value)
	default:
		// Unknown format switch, return as-is
		return value
	}
}

// firstCap capitalizes the first letter of the string.
func firstCap(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}

	return string(runes)
}

// titleCase converts the string to title case (capitalize first letter of each word).
func titleCase(s string) string {
	if s == "" {
		return s
	}

	words := strings.Fields(s)
	for i, word := range words {
		words[i] = firstCap(strings.ToLower(word))
	}

	return strings.Join(words, " ")
}

// replaceSimpleField replaces a SimpleField element with a Run containing the value.
//
//nolint:revive // function-length: field replacement requires element manipulation steps
func (m *MailMerge) replaceSimpleField(
	field *MergeField,
	value string,
) error {
	if field.Element == nil {
		return errors.New("field element is nil")
	}

	// Get the SimpleField element
	simpleField, ok := field.Element.(*elements.SimpleField)
	if !ok {
		return errors.New(
			"field element is not a SimpleField",
		)
	}

	// Get parent element
	parent := simpleField.Parent()
	if parent == nil {
		return errors.New("field has no parent")
	}

	parentComposite, ok := parent.(openxml.CompositeElement)
	if !ok {
		return errors.New(
			"field parent is not a composite element",
		)
	}

	// Create replacement Run with the value
	replacementRun := elements.NewRun(value)

	// Try to preserve formatting from the SimpleField's children
	// Look for Run elements with RunProperties
	for child := range simpleField.Children() {
		if run, ok := child.(*elements.Run); ok {
			if props := run.Properties(); props != nil {
				// Clone the properties to the replacement run
				replacementProps := props.CloneNode(true).(*elements.RunProperties)
				replacementRun.InsertBefore(
					replacementProps,
					replacementRun.FirstChild(),
				)

				break
			}
		}
	}

	// Find the position of SimpleField in parent's children
	var beforeElement openxml.Element
	foundField := false
	for child := range parentComposite.Children() {
		if foundField {
			beforeElement = child

			break
		}
		if child == simpleField {
			foundField = true
		}
	}

	// Remove the SimpleField
	parentComposite.RemoveChild(simpleField)

	// Insert the replacement Run at the same position
	if beforeElement != nil {
		parentComposite.InsertBefore(
			replacementRun,
			beforeElement,
		)
	} else {
		// It was the last child, append it
		parentComposite.AppendChild(replacementRun)
	}

	return nil
}

// replaceComplexField replaces a ComplexField with a Run containing the value.
//
//nolint:revive // function-length,early-return: complex field replacement requires state tracking
func (m *MailMerge) replaceComplexField(
	field *MergeField,
	value string,
) error {
	if field.Paragraph == nil {
		return errors.New(
			"field has no paragraph",
		)
	}

	if field.BeginRun == nil ||
		field.EndRun == nil {
		return errors.New(
			"field missing begin or end run",
		)
	}

	// Collect all runs between begin and end (inclusive)
	var runsToRemove []*elements.Run
	var resultRuns []*elements.Run
	inField := false
	inResult := false

	for child := range field.Paragraph.Children() {
		run, ok := child.(*elements.Run)
		if !ok {
			continue
		}

		if run == field.BeginRun {
			inField = true
			runsToRemove = append(
				runsToRemove,
				run,
			)

			continue
		}

		if inField {
			runsToRemove = append(
				runsToRemove,
				run,
			)

			if field.SeparateRun != nil &&
				run == field.SeparateRun {
				inResult = true

				continue
			}

			if inResult && run != field.EndRun {
				resultRuns = append(
					resultRuns,
					run,
				)
			}

			if run == field.EndRun {
				break
			}
		}
	}

	// Create replacement Run with the value
	replacementRun := elements.NewRun(value)

	// Preserve formatting from result runs (if any)
	if len(resultRuns) > 0 {
		firstResultRun := resultRuns[0]
		if props := firstResultRun.Properties(); props != nil {
			// Clone the properties to the replacement run
			replacementProps := props.CloneNode(true).(*elements.RunProperties)
			replacementRun.InsertBefore(
				replacementProps,
				replacementRun.FirstChild(),
			)
		}
	}

	// Find the position to insert the replacement
	// We want to insert it where the BeginRun was
	var beforeElement openxml.Element
	foundBegin := false
	for child := range field.Paragraph.Children() {
		if foundBegin {
			beforeElement = child

			break
		}
		if child == field.BeginRun {
			foundBegin = true
		}
	}

	// Remove all runs in the field range
	for _, run := range runsToRemove {
		field.Paragraph.RemoveChild(run)
	}

	// Insert replacement Run at the begin position
	if beforeElement != nil {
		field.Paragraph.InsertBefore(
			replacementRun,
			beforeElement,
		)
	} else {
		// It was the last child, append it
		field.Paragraph.AppendChild(replacementRun)
	}

	return nil
}

// validateFields validates that all merge fields exist in the data source.
// Skips validation for special fields like GREETINGLINE.
// Returns an error if StrictFields is true and any fields are missing.
func (m *MailMerge) validateFields(
	fields []*MergeField,
	ds DataSource,
) error {
	if !m.options.StrictFields {
		return nil
	}

	// Get available fields from data source
	availableFields := ds.Fields()
	fieldMap := make(map[string]bool)
	for _, field := range availableFields {
		fieldMap[strings.ToLower(field)] = true
	}

	// Check each merge field
	var missingFields []string
	for _, field := range fields {
		// Skip special fields
		if field.FieldName == FieldNameGreetingLine {
			continue
		}

		// Check if field exists (case-insensitive)
		fieldNameLower := strings.ToLower(
			field.FieldName,
		)
		if !fieldMap[fieldNameLower] {
			missingFields = append(
				missingFields,
				field.FieldName,
			)
		}
	}

	if len(missingFields) > 0 {
		return fmt.Errorf(
			"missing fields in data source: %v",
			missingFields,
		)
	}

	return nil
}
