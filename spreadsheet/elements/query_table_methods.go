package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// ConnectionId returns the connection ID. Attribute: connectionId.
func (qt *QueryTableRoot) ConnectionId() (uint32, bool) {
	attr, found := qt.GetAttribute(
		"connectionId",
		"",
	)
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val), true
}

// SetConnectionId sets the connection ID. Attribute: connectionId.
func (qt *QueryTableRoot) SetConnectionId(
	id uint32,
) {
	qt.SetAttribute(
		openxml.NewAttribute(
			"",
			"connectionId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10, //nolint:revive // add-constant
			),
		),
	)
}

// ClearConnectionId removes the connection ID attribute.
func (qt *QueryTableRoot) ClearConnectionId() {
	qt.RemoveAttribute("connectionId", "")
}

// AutoFormatId returns the auto format ID. Attribute: autoFormatId.
func (qt *QueryTableRoot) AutoFormatId() (uint32, bool) {
	attr, found := qt.GetAttribute(
		"autoFormatId",
		"",
	)
	if !found {
		return 0, false
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val), true
}

// SetAutoFormatId sets the auto format ID. Attribute: autoFormatId.
func (qt *QueryTableRoot) SetAutoFormatId(
	id uint32,
) {
	qt.SetAttribute(
		openxml.NewAttribute(
			"",
			"autoFormatId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10, //nolint:revive // add-constant
			),
		),
	)
}

// ClearAutoFormatId removes the auto format ID attribute.
func (qt *QueryTableRoot) ClearAutoFormatId() {
	qt.RemoveAttribute("autoFormatId", "")
}

// ApplyNumberFormats returns whether number formats are applied.
// Attribute: applyNumberFormats.
func (qt *QueryTableRoot) ApplyNumberFormats() bool {
	attr, found := qt.GetAttribute(
		"applyNumberFormats",
		"",
	)
	if !found {
		return true // Default
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyNumberFormats sets whether number formats are applied.
// Attribute: applyNumberFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (qt *QueryTableRoot) SetApplyNumberFormats(
	value bool,
) {
	if value {
		qt.RemoveAttribute(
			"applyNumberFormats",
			"",
		) // true is default
	} else {
		qt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyNumberFormats",
				"",
				attrValueZero,
			),
		)
	}
}

// ApplyBorderFormats returns whether border formats are applied.
// Attribute: applyBorderFormats.
func (qt *QueryTableRoot) ApplyBorderFormats() bool {
	attr, found := qt.GetAttribute(
		"applyBorderFormats",
		"",
	)
	if !found {
		return true // Default
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyBorderFormats sets whether border formats are applied.
// Attribute: applyBorderFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (qt *QueryTableRoot) SetApplyBorderFormats(
	value bool,
) {
	if value {
		qt.RemoveAttribute(
			"applyBorderFormats",
			"",
		) // true is default
	} else {
		qt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyBorderFormats",
				"",
				attrValueZero,
			),
		)
	}
}

// ApplyFontFormats returns whether font formats are applied.
// Attribute: applyFontFormats.
func (qt *QueryTableRoot) ApplyFontFormats() bool {
	attr, found := qt.GetAttribute(
		"applyFontFormats",
		"",
	)
	if !found {
		return true // Default
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyFontFormats sets whether font formats are applied.
// Attribute: applyFontFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (qt *QueryTableRoot) SetApplyFontFormats(
	value bool,
) {
	if value {
		qt.RemoveAttribute(
			"applyFontFormats",
			"",
		) // true is default
	} else {
		qt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyFontFormats",
				"",
				attrValueZero,
			),
		)
	}
}

// ApplyPatternFormats returns whether pattern formats are applied.
// Attribute: applyPatternFormats.
func (qt *QueryTableRoot) ApplyPatternFormats() bool {
	attr, found := qt.GetAttribute(
		"applyPatternFormats",
		"",
	)
	if !found {
		return true // Default
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyPatternFormats sets whether pattern formats are applied.
// Attribute: applyPatternFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (qt *QueryTableRoot) SetApplyPatternFormats(
	value bool,
) {
	if value {
		qt.RemoveAttribute(
			"applyPatternFormats",
			"",
		) // true is default
	} else {
		qt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyPatternFormats",
				"",
				attrValueZero,
			),
		)
	}
}

// ApplyAlignmentFormats returns whether alignment formats are applied.
// Attribute: applyAlignmentFormats.
func (qt *QueryTableRoot) ApplyAlignmentFormats() bool {
	attr, found := qt.GetAttribute(
		"applyAlignmentFormats",
		"",
	)
	if !found {
		return true // Default
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyAlignmentFormats sets whether alignment formats are applied.
// Attribute: applyAlignmentFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (qt *QueryTableRoot) SetApplyAlignmentFormats(
	value bool,
) {
	if value {
		qt.RemoveAttribute(
			"applyAlignmentFormats",
			"",
		) // true is default
	} else {
		qt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyAlignmentFormats",
				"",
				attrValueZero,
			),
		)
	}
}

// ApplyWidthHeightFormats returns whether width/height formats are
// applied. Attribute: applyWidthHeightFormats.
func (qt *QueryTableRoot) ApplyWidthHeightFormats() bool {
	attr, found := qt.GetAttribute(
		"applyWidthHeightFormats",
		"",
	)
	if !found {
		return true // Default
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyWidthHeightFormats sets whether width/height formats are
// applied. Attribute: applyWidthHeightFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (qt *QueryTableRoot) SetApplyWidthHeightFormats(
	value bool,
) {
	if value {
		qt.RemoveAttribute(
			"applyWidthHeightFormats",
			"",
		) // true is default
	} else {
		qt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyWidthHeightFormats",
				"",
				attrValueZero,
			),
		)
	}
}
