// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// FieldType represents the type of merge field in the document.
type FieldType int

const (
	// FieldTypeSimple represents a simple field (w:fldSimple element).
	// Simple fields have the field code as an attribute and are easier to process.
	// Example: <w:fldSimple w:instr="MERGEFIELD FirstName \* MERGEFORMAT">
	//            <w:r><w:t>«FirstName»</w:t></w:r>
	//          </w:fldSimple>
	FieldTypeSimple FieldType = iota

	// FieldTypeComplex represents a complex field (w:fldChar markers).
	// Complex fields use begin/separate/end markers and can span multiple runs.
	// The field code is stored in w:instrText elements between begin and separate.
	// Example: <w:r><w:fldChar w:fldCharType="begin"/></w:r>
	//          <w:r><w:instrText>MERGEFIELD FirstName \* MERGEFORMAT</w:instrText></w:r>
	//          <w:r><w:fldChar w:fldCharType="separate"/></w:r>
	//          <w:r><w:t>«FirstName»</w:t></w:r>
	//          <w:r><w:fldChar w:fldCharType="end"/></w:r>
	FieldTypeComplex
)

// String returns a string representation of the FieldType.
func (ft FieldType) String() string {
	switch ft {
	case FieldTypeSimple:
		return "Simple"
	case FieldTypeComplex:
		return "Complex"
	default:
		return "Unknown"
	}
}

// MergeField represents a single merge field found in the document.
// It can be either a SimpleField or a ComplexField.
//
// For SimpleField (FieldTypeSimple):
//   - Element points to the w:fldSimple element
//   - FieldCode contains the w:instr attribute value
//   - BeginRun, SeparateRun, EndRun are nil
//
// For ComplexField (FieldTypeComplex):
//   - Element points to the run containing the begin fldChar
//   - BeginRun points to the run with fldCharType="begin"
//   - SeparateRun points to the run with fldCharType="separate" (may be nil)
//   - EndRun points to the run with fldCharType="end"
//   - FieldCode is reconstructed from w:instrText elements
type MergeField struct {
	// FieldName is the name of the merge field (e.g., "FirstName" from "MERGEFIELD FirstName")
	FieldName string

	// FieldType indicates whether this is a simple or complex field
	FieldType FieldType

	// Element is the primary element for this field:
	//   - For SimpleField: the w:fldSimple element
	//   - For ComplexField: the run containing the begin fldChar
	Element openxml.Element

	// BeginRun is the run containing the begin fldChar (ComplexField only)
	BeginRun *elements.Run

	// EndRun is the run containing the end fldChar (ComplexField only)
	EndRun *elements.Run

	// SeparateRun is the run containing the separate fldChar (ComplexField only, may be nil)
	SeparateRun *elements.Run

	// FieldCode is the complete field instruction text
	// Example: "MERGEFIELD FirstName \* MERGEFORMAT"
	FieldCode string

	// Switches contains parsed field switches (e.g., "\*" -> "MERGEFORMAT")
	// Common switches:
	//   - "\*" (formatting switch): preserves formatting
	//   - "\b" (text before): text to insert before non-empty field
	//   - "\f" (text after): text to insert after non-empty field
	Switches map[string]string

	// Paragraph is the containing paragraph for this field
	// This is used for navigation and context during replacement
	Paragraph *elements.Paragraph
}
