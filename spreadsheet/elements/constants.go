// Package elements provides common constants used across spreadsheet elements.
package elements

import "github.com/connerohnesorge/goffice/openxml"

// Common element and attribute names.
const (
	elemNameRef           = "ref"
	elemNameDateGroupItem = "dateGroupItem"
	elemNameFilter        = "filter"
	elemNameIconSet       = "iconSet"
	elemNameID            = "id"
	elemNameCol           = "col"
	elemNameColOff        = "colOff"
	elemNameRow           = "row"
	elemNameRowOff        = "rowOff"
	attrNameLocalSheetID  = "localSheetId"
	attrNamePassword      = "password"
	attrNameAlgorithmName = "algorithmName"
	attrNameHashValue     = "hashValue"
	attrNameSaltValue     = "saltValue"
	attrNameSpinCount     = "spinCount"
	attrNameArray         = "array"
	attrValueWorksheet    = "worksheet"
	PrefixR               = "r"
)

// Common namespace constants.
const (
	NamespaceRelationships = openxml.NamespaceRelationships
)

// DXF (Differential Format) element order constants.
const (
	dxfOrderFont       = 0
	dxfOrderNumFmt     = 1
	dxfOrderFill       = 2
	dxfOrderAlignment  = 3
	dxfOrderBorder     = 4
	dxfOrderProtection = 5
)

// Bit size constants for parsing various integer types.
const (
	bitSize8  = 8
	bitSize16 = 16
	bitSize64 = 64
)
