// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"strings"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// Field name constants
const (
	FieldNameGreetingLine = "GREETINGLINE"
	// Note: FieldNameSkipIf is defined in advanced_fields.go
)

// findMergeFields searches the document for all MERGEFIELD and GREETINGLINE fields.
// It searches in the main document body, headers, footers, and tables.
// Returns a slice of MergeField structs representing all found fields.
func (m *MailMerge) findMergeFields() []*MergeField {
	fields := make([]*MergeField, 0)

	if m.template == nil {
		return fields
	}

	// Get the main document part
	mainPart := m.template.MainPart()
	if mainPart == nil {
		return fields
	}

	// Get the document root element
	doc := mainPart.Document()
	if doc == nil {
		return fields
	}

	// Search in the document body
	body := doc.Body()
	if body != nil {
		fields = append(
			fields,
			findFieldsInElement(body)...)
	}

	// TODO: Search in headers and footers in later phases

	return fields
}

// findFieldsInElement recursively searches an element and its children for merge fields.
// It handles both SimpleField and ComplexField types.
func findFieldsInElement(
	elem openxml.Element,
) []*MergeField {
	return findFieldsInElementWithPara(elem, nil)
}

// findFieldsInElementWithPara is the internal implementation that tracks the containing paragraph.
//
//nolint:revive // modifies-parameter: intentional to track paragraph context during recursion
func findFieldsInElementWithPara(
	elem openxml.Element,
	currentPara *elements.Paragraph,
) []*MergeField {
	fields := make([]*MergeField, 0)

	if elem == nil {
		return fields
	}

	// Check if this element is a paragraph - update current paragraph context
	if elem.LocalName() == "p" &&
		elem.NamespaceURI() == elements.NamespaceWML {
		if para, ok := elem.(*elements.Paragraph); ok {
			currentPara = para
			// Look for complex fields in this paragraph
			complexFields := findComplexFieldsInParagraph(
				para,
			)
			fields = append(
				fields,
				complexFields...)
		}
	}

	// Check if this is a SimpleField
	if elem.LocalName() == "fldSimple" &&
		elem.NamespaceURI() == elements.NamespaceWML {
		if simpleField, ok := elem.(*elements.SimpleField); ok {
			if field := parseSimpleField(simpleField, currentPara); field != nil {
				fields = append(fields, field)
			}
		}
	}

	// Recursively search children if this is a composite element
	if composite, ok := elem.(openxml.CompositeElement); ok {
		for child := range composite.Children() {
			childFields := findFieldsInElementWithPara(
				child,
				currentPara,
			)
			fields = append(
				fields,
				childFields...)
		}
	}

	return fields
}

// parseSimpleField parses a SimpleField element and extracts merge field information.
// Returns nil if this is not a MERGEFIELD or GREETINGLINE.
func parseSimpleField(
	simpleField *elements.SimpleField,
	para *elements.Paragraph,
) *MergeField {
	if simpleField == nil {
		return nil
	}

	// Get the instruction attribute
	instrAttr := simpleField.Instruction
	if instrAttr == nil {
		return nil
	}

	instrText := instrAttr.Value()
	if instrText == "" {
		return nil
	}

	// Parse the field code
	fieldName, switches := parseMergeFieldCode(
		instrText,
	)
	if fieldName == "" {
		return nil
	}

	return &MergeField{
		FieldName: fieldName,
		FieldType: FieldTypeSimple,
		Element:   simpleField,
		FieldCode: instrText,
		Switches:  switches,
		Paragraph: para,
	}
}

// findComplexFieldsInParagraph finds all complex fields in a paragraph.
// Complex fields use fldChar markers (begin, separate, end) and instrText elements.
//
//nolint:revive // function-length,early-return,max-control-nesting: field parsing requires state machine logic
func findComplexFieldsInParagraph(
	para *elements.Paragraph,
) []*MergeField {
	fields := make([]*MergeField, 0)

	// Track complex field state
	var beginRun *elements.Run
	var separateRun *elements.Run
	var instrTextParts []string
	inField := false

	// Iterate through all runs in the paragraph
	for child := range para.Children() {
		run, ok := child.(*elements.Run)
		if !ok {
			continue
		}

		// Check for field characters in this run
		for runChild := range run.Children() {
			if runChild.LocalName() != "fldChar" ||
				runChild.NamespaceURI() != elements.NamespaceWML {
				// Check for instrText when inside a field
				if inField &&
					runChild.LocalName() == "instrText" &&
					runChild.NamespaceURI() == elements.NamespaceWML {
					if instrText, ok := runChild.(*elements.InstrText); ok {
						instrTextParts = append(
							instrTextParts,
							instrText.Text(),
						)
					}
				}

				continue
			}

			fieldChar, ok := runChild.(*elements.FieldChar)
			if !ok {
				continue
			}

			charType := fieldChar.Type()

			switch charType {
			case elements.FieldCharBegin:
				// Start of a new complex field
				beginRun = run
				separateRun = nil
				instrTextParts = make([]string, 0)
				inField = true

			case elements.FieldCharSeparate:
				// Separator between field code and result
				if inField {
					separateRun = run
				}

			case elements.FieldCharEnd:
				// End of complex field
				if inField && beginRun != nil {
					// Reconstruct the field code
					fieldCode := reconstructFieldCode(
						instrTextParts,
					)
					if fieldCode != "" {
						// Parse the field code
						fieldName, switches := parseMergeFieldCode(
							fieldCode,
						)
						if fieldName != "" {
							field := &MergeField{
								FieldName:   fieldName,
								FieldType:   FieldTypeComplex,
								Element:     beginRun,
								BeginRun:    beginRun,
								SeparateRun: separateRun,
								EndRun:      run,
								FieldCode:   fieldCode,
								Switches:    switches,
								Paragraph:   para,
							}
							fields = append(
								fields,
								field,
							)
						}
					}

					// Reset state
					beginRun = nil
					separateRun = nil
					instrTextParts = nil
					inField = false
				}
			}
		}
	}

	return fields
}

// reconstructFieldCode reassembles the field instruction from multiple instrText parts.
// Complex fields can have the instruction text split across multiple runs.
func reconstructFieldCode(parts []string) string {
	if len(parts) == 0 {
		return ""
	}

	// Join parts with spaces, then normalize whitespace
	combined := strings.Join(parts, " ")
	combined = strings.TrimSpace(combined)

	// Normalize multiple spaces to single space
	for strings.Contains(combined, "  ") {
		combined = strings.ReplaceAll(
			combined,
			"  ",
			" ",
		)
	}

	return combined
}

// parseMergeFieldCode parses a field instruction text and extracts the field name and switches.
// Supports MERGEFIELD and GREETINGLINE field types.
// Returns empty string for fieldName if this is not a recognized merge field.
//
// Examples:
//   - "MERGEFIELD FirstName \* MERGEFORMAT" -> ("FirstName", {"\*": "MERGEFORMAT"})
//   - "GREETINGLINE \f \"Dear \"" -> ("GREETINGLINE", {"\f": "Dear "})
//   - "PAGE" -> ("", {}) // Not a merge field
func parseMergeFieldCode(
	code string,
) (fieldName string, switches map[string]string) {
	switches = make(map[string]string)

	// Tokenize the field code
	tokens := tokenizeFieldCode(code)
	if len(tokens) == 0 {
		return "", switches
	}

	// First token should be the field type
	fieldType := strings.ToUpper(tokens[0])

	switch fieldType {
	case "MERGEFIELD":
		// MERGEFIELD FieldName [switches]
		if len(tokens) < 2 {
			return "", switches
		}
		fieldName = tokens[1]

		// Parse switches
		switches = parseFieldSwitches(tokens[2:])

	case FieldNameGreetingLine:
		// GREETINGLINE is a special merge field
		fieldName = FieldNameGreetingLine

		// Parse switches
		switches = parseFieldSwitches(tokens[1:])

	case FieldNameSkipIf:
		// SKIPIF is a conditional merge field
		fieldName = FieldNameSkipIf

		// For SKIPIF, we don't parse switches in the traditional way
		// The entire field code is needed for evaluation
		// But we still set fieldName so it's recognized
		switches = make(map[string]string)

	default:
		// Not a recognized merge field type
		return "", switches
	}

	return fieldName, switches
}

// tokenizeFieldCode splits a field code into tokens, respecting quoted strings.
// Quoted strings are kept as single tokens with quotes removed.
//
// Example: `MERGEFIELD Name \f "Dear " \* MERGEFORMAT`
// Returns: ["MERGEFIELD", "Name", "\f", "Dear ", "\*", "MERGEFORMAT"]
//
//nolint:revive // function-length,max-control-nesting: tokenizer requires state tracking
func tokenizeFieldCode(code string) []string {
	tokens := make([]string, 0)
	var current strings.Builder
	inQuotes := false
	escapeNext := false

	for i, ch := range code {
		if escapeNext {
			current.WriteRune(ch)
			escapeNext = false

			continue
		}

		switch ch {
		case '\\':
			// Check if this is a switch (backslash followed by letter/*/# etc.)
			if !inQuotes && i+1 < len(code) {
				nextCh := rune(code[i+1])
				if (nextCh >= 'a' && nextCh <= 'z') ||
					(nextCh >= 'A' && nextCh <= 'Z') ||
					nextCh == '*' ||
					nextCh == '#' {
					// This is a switch marker - save current token and start new
					if current.Len() > 0 {
						tokens = append(
							tokens,
							current.String(),
						)
						current.Reset()
					}
					current.WriteRune(ch)
				} else {
					escapeNext = true
				}
			} else {
				current.WriteRune(ch)
			}

		case '"':
			if inQuotes {
				// End of quoted string - save it (even if empty)
				tokens = append(
					tokens,
					current.String(),
				)
				current.Reset()
				inQuotes = false
			} else {
				// Start of quoted string - save any previous token
				if current.Len() > 0 {
					tokens = append(tokens, current.String())
					current.Reset()
				}
				inQuotes = true
			}

		case ' ', '\t', '\n', '\r':
			if inQuotes {
				current.WriteRune(ch)
			} else if current.Len() > 0 {
				// Whitespace outside quotes - token separator
				tokens = append(tokens, current.String())
				current.Reset()
			}

		default:
			current.WriteRune(ch)
		}
	}

	// Add final token
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

// parseFieldSwitches parses field switch tokens into a map.
// Switches start with backslash and may have values.
//
// Example tokens: ["\*", "MERGEFORMAT", "\f", "Dear "]
// Returns: {"\*": "MERGEFORMAT", "\f": "Dear "}
//
//nolint:revive // early-return,max-control-nesting: switch parsing logic is clearest as-is
func parseFieldSwitches(
	tokens []string,
) map[string]string {
	switches := make(map[string]string)

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Check if this is a switch (starts with \)
		if token != "" && token[0] == '\\' {
			switchName := token

			// Check if there's a value for this switch
			if i+1 < len(tokens) {
				nextToken := tokens[i+1]
				// If next token doesn't start with \, it's the value for this switch
				if nextToken == "" ||
					nextToken[0] != '\\' {
					switches[switchName] = nextToken
					i++ // Skip the value token

					continue
				}
			}

			// Switch with no value
			switches[switchName] = ""
		}
	}

	return switches
}
