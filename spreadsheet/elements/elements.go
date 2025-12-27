// Package elements provides Excel document element types for SpreadsheetML.
//
//nolint:revive // max-public-structs: elements package defines many public types
package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// SpreadsheetML namespace and prefix constants.
const (
	// NamespaceSML is the main SpreadsheetML namespace.
	NamespaceSML = "http://schemas.openxmlformats.org/spreadsheetml/2006/main"

	// PrefixX is the canonical prefix for SpreadsheetML elements.
	PrefixX = "x"

	// PrefixDefault is the empty prefix when using default namespace.
	PrefixDefault = ""

	// NamespaceRelationships is the OPC relationships namespace.
	NamespaceRelationships = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"

	// PrefixR is the prefix for relationship references.
	PrefixR = "r"
)

// Common attribute value constants.
const (
	attrValueTrue  = "true"
	attrValueFalse = "false"
	attrValueOne   = "1"
	attrValueZero  = "0"
)

// Numeric parsing constants.
const (
	parseBase10 = 10 // Base 10 for parsing/formatting integers

	bitSize8  = 8  // 8-bit integer parsing
	bitSize16 = 16 // 16-bit integer parsing
	bitSize32 = 32 // 32-bit integer parsing
	bitSize64 = 64 // 64-bit integer/float parsing
)

// DXF element order constants for insertDxfElement.
const (
	dxfOrderFont       = 0
	dxfOrderNumFmt     = 1
	dxfOrderFill       = 2
	dxfOrderAlignment  = 3
	dxfOrderBorder     = 4
	dxfOrderProtection = 5
)

// Common element name constants.
const (
	elemNameDateGroupItem = "dateGroupItem"
	elemNameIconSet       = "iconSet"
	elemNameID            = "id"
	elemNameCol           = "col"
	elemNameColOff        = "colOff"
	elemNameRow           = "row"
	elemNameRowOff        = "rowOff"
	elemNameRef           = "ref"
)

// Common attribute name constants.
const (
	attrNamePassword      = "password"
	attrNameAlgorithmName = "algorithmName"
	attrNameHashValue     = "hashValue"
	attrNameSaltValue     = "saltValue"
	attrNameSpinCount     = "spinCount"
	attrNameCount         = "count"
	attrNameLocalSheetID  = "localSheetId"
	attrNameArray         = "array"
)

// Default DPI for web publishing.
const defaultDPI = 96

// NewNamespaceAttribute creates a new xmlns attribute for the SpreadsheetML namespace.
func NewNamespaceAttribute() openxml.OpenXmlAttribute {
	return openxml.NewAttribute(
		"http://www.w3.org/2000/xmlns/",
		"xmlns",
		"",
		NamespaceSML,
	)
}
