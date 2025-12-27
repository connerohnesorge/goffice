package elements

//revive:disable:file-length-limit many pivot table types

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PivotTableDefinition represents the pivot table definition root element
// (x:pivotTableDefinition).
// This element is the root of a pivot table definition part and contains
// the structure and configuration of a pivot table.
type PivotTableDefinition struct {
	*openxml.CompositeElementBase
}

// NewPivotTableDefinition creates a new PivotTableDefinition element.
func NewPivotTableDefinition() *PivotTableDefinition {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"pivotTableDefinition",
		PrefixDefault,
	)

	return &PivotTableDefinition{
		CompositeElementBase: elem,
	}
}

// Name returns the name of the pivot table. Attribute: name.
func (pt *PivotTableDefinition) Name() string {
	attr, found := pt.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the name of the pivot table. Attribute: name.
func (pt *PivotTableDefinition) SetName(
	name string,
) {
	if name == "" {
		pt.RemoveAttribute("name", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// CacheId returns the cache ID. Attribute: cacheId.
func (pt *PivotTableDefinition) CacheId() uint32 {
	attr, found := pt.GetAttribute("cacheId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetCacheId sets the cache ID. Attribute: cacheId.
func (pt *PivotTableDefinition) SetCacheId(
	id uint32,
) {
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"cacheId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10,
			),
		),
	)
}

// DataOnRows returns whether data fields are on rows. Attribute: dataOnRows.
func (pt *PivotTableDefinition) DataOnRows() bool {
	attr, found := pt.GetAttribute(
		"dataOnRows",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDataOnRows sets whether data fields are on rows. Attribute: dataOnRows.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetDataOnRows(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"dataOnRows",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("dataOnRows", "")
	}
}

// DataPosition returns the data field position. Attribute: dataPosition.
func (pt *PivotTableDefinition) DataPosition() (uint32, bool) {
	attr, found := pt.GetAttribute(
		"dataPosition",
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

// SetDataPosition sets the data field position. Attribute: dataPosition.
func (pt *PivotTableDefinition) SetDataPosition(
	pos uint32,
) {
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"dataPosition",
			"",
			strconv.FormatUint(
				uint64(pos),
				parseBase10,
			),
		),
	)
}

// ClearDataPosition removes the data position attribute.
func (pt *PivotTableDefinition) ClearDataPosition() {
	pt.RemoveAttribute("dataPosition", "")
}

// AutoFormatId returns the auto format ID. Attribute: autoFormatId.
func (pt *PivotTableDefinition) AutoFormatId() (uint32, bool) {
	attr, found := pt.GetAttribute(
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
func (pt *PivotTableDefinition) SetAutoFormatId(
	id uint32,
) {
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"autoFormatId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10,
			),
		),
	)
}

// ApplyNumberFormats returns whether to apply number formats.
// Attribute: applyNumberFormats.
func (pt *PivotTableDefinition) ApplyNumberFormats() bool {
	attr, found := pt.GetAttribute(
		"applyNumberFormats",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyNumberFormats sets whether to apply number formats.
// Attribute: applyNumberFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetApplyNumberFormats(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyNumberFormats",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("applyNumberFormats", "")
	}
}

// ApplyBorderFormats returns whether to apply border formats.
// Attribute: applyBorderFormats.
func (pt *PivotTableDefinition) ApplyBorderFormats() bool {
	attr, found := pt.GetAttribute(
		"applyBorderFormats",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyBorderFormats sets whether to apply border formats.
// Attribute: applyBorderFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetApplyBorderFormats(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyBorderFormats",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("applyBorderFormats", "")
	}
}

// ApplyFontFormats returns whether to apply font formats.
// Attribute: applyFontFormats.
func (pt *PivotTableDefinition) ApplyFontFormats() bool {
	attr, found := pt.GetAttribute(
		"applyFontFormats",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyFontFormats sets whether to apply font formats.
// Attribute: applyFontFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetApplyFontFormats(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyFontFormats",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("applyFontFormats", "")
	}
}

// ApplyPatternFormats returns whether to apply pattern formats.
// Attribute: applyPatternFormats.
func (pt *PivotTableDefinition) ApplyPatternFormats() bool {
	attr, found := pt.GetAttribute(
		"applyPatternFormats",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyPatternFormats sets whether to apply pattern formats.
// Attribute: applyPatternFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetApplyPatternFormats(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyPatternFormats",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("applyPatternFormats", "")
	}
}

// ApplyAlignmentFormats returns whether to apply alignment formats.
// Attribute: applyAlignmentFormats.
func (pt *PivotTableDefinition) ApplyAlignmentFormats() bool {
	attr, found := pt.GetAttribute(
		"applyAlignmentFormats",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyAlignmentFormats sets whether to apply alignment formats.
// Attribute: applyAlignmentFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetApplyAlignmentFormats(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyAlignmentFormats",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("applyAlignmentFormats", "")
	}
}

// ApplyWidthHeightFormats returns whether to apply width/height formats.
// Attribute: applyWidthHeightFormats.
func (pt *PivotTableDefinition) ApplyWidthHeightFormats() bool {
	attr, found := pt.GetAttribute(
		"applyWidthHeightFormats",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyWidthHeightFormats sets whether to apply width/height formats.
// Attribute: applyWidthHeightFormats.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetApplyWidthHeightFormats(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyWidthHeightFormats",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("applyWidthHeightFormats", "")
	}
}

// DataCaption returns the data caption. Attribute: dataCaption.
func (pt *PivotTableDefinition) DataCaption() string {
	attr, found := pt.GetAttribute(
		"dataCaption",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDataCaption sets the data caption. Attribute: dataCaption.
func (pt *PivotTableDefinition) SetDataCaption(
	caption string,
) {
	if caption == "" {
		pt.RemoveAttribute("dataCaption", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"dataCaption",
			"",
			caption,
		),
	)
}

// GrandTotalCaption returns the grand total caption.
// Attribute: grandTotalCaption.
func (pt *PivotTableDefinition) GrandTotalCaption() string {
	attr, found := pt.GetAttribute(
		"grandTotalCaption",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetGrandTotalCaption sets the grand total caption.
// Attribute: grandTotalCaption.
func (pt *PivotTableDefinition) SetGrandTotalCaption(
	caption string,
) {
	if caption == "" {
		pt.RemoveAttribute(
			"grandTotalCaption",
			"",
		)

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"grandTotalCaption",
			"",
			caption,
		),
	)
}

// ErrorCaption returns the error caption. Attribute: errorCaption.
func (pt *PivotTableDefinition) ErrorCaption() string {
	attr, found := pt.GetAttribute(
		"errorCaption",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetErrorCaption sets the error caption. Attribute: errorCaption.
func (pt *PivotTableDefinition) SetErrorCaption(
	caption string,
) {
	if caption == "" {
		pt.RemoveAttribute("errorCaption", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"errorCaption",
			"",
			caption,
		),
	)
}

// ShowError returns whether to show error values. Attribute: showError.
func (pt *PivotTableDefinition) ShowError() bool {
	attr, found := pt.GetAttribute(
		"showError",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowError sets whether to show error values. Attribute: showError.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowError(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"showError",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("showError", "")
	}
}

// MissingCaption returns the missing value caption. Attribute: missingCaption.
func (pt *PivotTableDefinition) MissingCaption() string {
	attr, found := pt.GetAttribute(
		"missingCaption",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetMissingCaption sets the missing value caption. Attribute: missingCaption.
func (pt *PivotTableDefinition) SetMissingCaption(
	caption string,
) {
	if caption == "" {
		pt.RemoveAttribute("missingCaption", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"missingCaption",
			"",
			caption,
		),
	)
}

// ShowMissing returns whether to show missing values. Attribute: showMissing.
func (pt *PivotTableDefinition) ShowMissing() bool {
	attr, found := pt.GetAttribute(
		"showMissing",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowMissing sets whether to show missing values. Attribute: showMissing.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowMissing(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"showMissing",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "showMissing", "", attrValueFalse),
		)
	}
}

// PageStyle returns the page field style. Attribute: pageStyle.
func (pt *PivotTableDefinition) PageStyle() string {
	attr, found := pt.GetAttribute(
		"pageStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPageStyle sets the page field style. Attribute: pageStyle.
func (pt *PivotTableDefinition) SetPageStyle(
	style string,
) {
	if style == "" {
		pt.RemoveAttribute("pageStyle", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"pageStyle",
			"",
			style,
		),
	)
}

// PivotTableStyle returns the pivot table style. Attribute: pivotTableStyle.
func (pt *PivotTableDefinition) PivotTableStyle() string {
	attr, found := pt.GetAttribute(
		"pivotTableStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPivotTableStyle sets the pivot table style. Attribute: pivotTableStyle.
func (pt *PivotTableDefinition) SetPivotTableStyle(
	style string,
) {
	if style == "" {
		pt.RemoveAttribute("pivotTableStyle", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"pivotTableStyle",
			"",
			style,
		),
	)
}

// VacatedStyle returns the vacated style. Attribute: vacatedStyle.
func (pt *PivotTableDefinition) VacatedStyle() string {
	attr, found := pt.GetAttribute(
		"vacatedStyle",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetVacatedStyle sets the vacated style. Attribute: vacatedStyle.
func (pt *PivotTableDefinition) SetVacatedStyle(
	style string,
) {
	if style == "" {
		pt.RemoveAttribute("vacatedStyle", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"vacatedStyle",
			"",
			style,
		),
	)
}

// Tag returns the pivot table tag. Attribute: tag.
func (pt *PivotTableDefinition) Tag() string {
	attr, found := pt.GetAttribute("tag", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTag sets the pivot table tag. Attribute: tag.
func (pt *PivotTableDefinition) SetTag(
	tag string,
) {
	if tag == "" {
		pt.RemoveAttribute("tag", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute("", "tag", "", tag),
	)
}

// UpdatedVersion returns the updated version. Attribute: updatedVersion.
func (pt *PivotTableDefinition) UpdatedVersion() uint8 {
	attr, found := pt.GetAttribute(
		"updatedVersion",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize8,
	)

	return uint8(val)
}

// SetUpdatedVersion sets the updated version. Attribute: updatedVersion.
func (pt *PivotTableDefinition) SetUpdatedVersion(
	version uint8,
) {
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"updatedVersion",
			"",
			strconv.FormatUint(
				uint64(version),
				parseBase10,
			),
		),
	)
}

// MinRefreshableVersion returns the minimum refreshable version.
// Attribute: minRefreshableVersion.
func (pt *PivotTableDefinition) MinRefreshableVersion() uint8 {
	attr, found := pt.GetAttribute(
		"minRefreshableVersion",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize8,
	)

	return uint8(val)
}

// SetMinRefreshableVersion sets the minimum refreshable version.
// Attribute: minRefreshableVersion.
func (pt *PivotTableDefinition) SetMinRefreshableVersion(
	version uint8,
) {
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"minRefreshableVersion",
			"",
			strconv.FormatUint(
				uint64(version),
				parseBase10,
			),
		),
	)
}

// AsteriskTotals returns whether to show asterisk for totals.
// Attribute: asteriskTotals.
func (pt *PivotTableDefinition) AsteriskTotals() bool {
	attr, found := pt.GetAttribute(
		"asteriskTotals",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAsteriskTotals sets whether to show asterisk for totals.
// Attribute: asteriskTotals.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetAsteriskTotals(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"asteriskTotals",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("asteriskTotals", "")
	}
}

// ShowItems returns whether to show items. Attribute: showItems.
func (pt *PivotTableDefinition) ShowItems() bool {
	attr, found := pt.GetAttribute(
		"showItems",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowItems sets whether to show items. Attribute: showItems.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowItems(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"showItems",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "showItems", "", attrValueFalse),
		)
	}
}

// EditData returns whether data can be edited. Attribute: editData.
func (pt *PivotTableDefinition) EditData() bool {
	attr, found := pt.GetAttribute("editData", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetEditData sets whether data can be edited. Attribute: editData.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetEditData(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"editData",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("editData", "")
	}
}

// DisableFieldList returns whether to disable field list.
// Attribute: disableFieldList.
func (pt *PivotTableDefinition) DisableFieldList() bool {
	attr, found := pt.GetAttribute(
		"disableFieldList",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDisableFieldList sets whether to disable field list.
// Attribute: disableFieldList.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetDisableFieldList(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"disableFieldList",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("disableFieldList", "")
	}
}

// ShowCalcMbrs returns whether to show calculated members.
// Attribute: showCalcMbrs.
func (pt *PivotTableDefinition) ShowCalcMbrs() bool {
	attr, found := pt.GetAttribute(
		"showCalcMbrs",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowCalcMbrs sets whether to show calculated members.
// Attribute: showCalcMbrs.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowCalcMbrs(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"showCalcMbrs",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "showCalcMbrs", "", attrValueFalse),
		)
	}
}

// VisualTotals returns whether to show visual totals.
// Attribute: visualTotals.
func (pt *PivotTableDefinition) VisualTotals() bool {
	attr, found := pt.GetAttribute(
		"visualTotals",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetVisualTotals sets whether to show visual totals.
// Attribute: visualTotals.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetVisualTotals(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"visualTotals",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "visualTotals", "", attrValueFalse),
		)
	}
}

// ShowMultipleLabel returns whether to show multiple labels.
// Attribute: showMultipleLabel.
func (pt *PivotTableDefinition) ShowMultipleLabel() bool {
	attr, found := pt.GetAttribute(
		"showMultipleLabel",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowMultipleLabel sets whether to show multiple labels.
// Attribute: showMultipleLabel.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowMultipleLabel(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"showMultipleLabel",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "showMultipleLabel", "", attrValueFalse),
		)
	}
}

// ShowDataDropDown returns whether to show data drop down.
// Attribute: showDataDropDown.
func (pt *PivotTableDefinition) ShowDataDropDown() bool {
	attr, found := pt.GetAttribute(
		"showDataDropDown",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowDataDropDown sets whether to show data drop down.
// Attribute: showDataDropDown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowDataDropDown(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"showDataDropDown",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "showDataDropDown", "", attrValueFalse),
		)
	}
}

// ShowDrill returns whether to show drill indicators.
// Attribute: showDrill.
func (pt *PivotTableDefinition) ShowDrill() bool {
	attr, found := pt.GetAttribute(
		"showDrill",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowDrill sets whether to show drill indicators.
// Attribute: showDrill.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowDrill(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"showDrill",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "showDrill", "", attrValueFalse),
		)
	}
}

// PrintDrill returns whether to print drill indicators.
// Attribute: printDrill.
func (pt *PivotTableDefinition) PrintDrill() bool {
	attr, found := pt.GetAttribute(
		"printDrill",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPrintDrill sets whether to print drill indicators.
// Attribute: printDrill.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetPrintDrill(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"printDrill",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("printDrill", "")
	}
}

// ShowMemberPropertyTips returns whether to show member property tips.
// Attribute: showMemberPropertyTips.
func (pt *PivotTableDefinition) ShowMemberPropertyTips() bool {
	attr, found := pt.GetAttribute(
		"showMemberPropertyTips",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowMemberPropertyTips sets whether to show member property tips.
// Attribute: showMemberPropertyTips.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowMemberPropertyTips(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"showMemberPropertyTips",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "showMemberPropertyTips", "", attrValueFalse),
		)
	}
}

// ShowDataTips returns whether to show data tips.
// Attribute: showDataTips.
func (pt *PivotTableDefinition) ShowDataTips() bool {
	attr, found := pt.GetAttribute(
		"showDataTips",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowDataTips sets whether to show data tips.
// Attribute: showDataTips.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowDataTips(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"showDataTips",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "showDataTips", "", attrValueFalse),
		)
	}
}

// EnableWizard returns whether to enable wizard.
// Attribute: enableWizard.
func (pt *PivotTableDefinition) EnableWizard() bool {
	attr, found := pt.GetAttribute(
		"enableWizard",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetEnableWizard sets whether to enable wizard.
// Attribute: enableWizard.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetEnableWizard(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"enableWizard",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "enableWizard", "", attrValueFalse),
		)
	}
}

// EnableDrill returns whether to enable drilling.
// Attribute: enableDrill.
func (pt *PivotTableDefinition) EnableDrill() bool {
	attr, found := pt.GetAttribute(
		"enableDrill",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetEnableDrill sets whether to enable drilling.
// Attribute: enableDrill.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetEnableDrill(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"enableDrill",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "enableDrill", "", attrValueFalse),
		)
	}
}

// EnableFieldProperties returns whether to enable field properties.
// Attribute: enableFieldProperties.
func (pt *PivotTableDefinition) EnableFieldProperties() bool {
	attr, found := pt.GetAttribute(
		"enableFieldProperties",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetEnableFieldProperties sets whether to enable field properties.
// Attribute: enableFieldProperties.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetEnableFieldProperties(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"enableFieldProperties",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "enableFieldProperties", "", attrValueFalse),
		)
	}
}

// PreserveFormatting returns whether to preserve formatting.
// Attribute: preserveFormatting.
func (pt *PivotTableDefinition) PreserveFormatting() bool {
	attr, found := pt.GetAttribute(
		"preserveFormatting",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetPreserveFormatting sets whether to preserve formatting.
// Attribute: preserveFormatting.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetPreserveFormatting(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"preserveFormatting",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "preserveFormatting", "", attrValueFalse),
		)
	}
}

// UseAutoFormatting returns whether to use auto formatting.
// Attribute: useAutoFormatting.
func (pt *PivotTableDefinition) UseAutoFormatting() bool {
	attr, found := pt.GetAttribute(
		"useAutoFormatting",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetUseAutoFormatting sets whether to use auto formatting.
// Attribute: useAutoFormatting.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetUseAutoFormatting(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"useAutoFormatting",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("useAutoFormatting", "")
	}
}

// PageWrap returns the page wrap count. Attribute: pageWrap.
func (pt *PivotTableDefinition) PageWrap() uint32 {
	attr, found := pt.GetAttribute("pageWrap", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetPageWrap sets the page wrap count. Attribute: pageWrap.
func (pt *PivotTableDefinition) SetPageWrap(
	count uint32,
) {
	if count == 0 {
		pt.RemoveAttribute("pageWrap", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"pageWrap",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// PageOverThenDown returns whether pages flow over then down.
// Attribute: pageOverThenDown.
func (pt *PivotTableDefinition) PageOverThenDown() bool {
	attr, found := pt.GetAttribute(
		"pageOverThenDown",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPageOverThenDown sets whether pages flow over then down.
// Attribute: pageOverThenDown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetPageOverThenDown(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"pageOverThenDown",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("pageOverThenDown", "")
	}
}

// SubtotalHiddenItems returns whether to subtotal hidden items.
// Attribute: subtotalHiddenItems.
func (pt *PivotTableDefinition) SubtotalHiddenItems() bool {
	attr, found := pt.GetAttribute(
		"subtotalHiddenItems",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSubtotalHiddenItems sets whether to subtotal hidden items.
// Attribute: subtotalHiddenItems.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetSubtotalHiddenItems(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"subtotalHiddenItems",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("subtotalHiddenItems", "")
	}
}

// RowGrandTotals returns whether to show row grand totals.
// Attribute: rowGrandTotals.
func (pt *PivotTableDefinition) RowGrandTotals() bool {
	attr, found := pt.GetAttribute(
		"rowGrandTotals",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetRowGrandTotals sets whether to show row grand totals.
// Attribute: rowGrandTotals.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetRowGrandTotals(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"rowGrandTotals",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "rowGrandTotals", "", attrValueFalse),
		)
	}
}

// ColGrandTotals returns whether to show column grand totals.
// Attribute: colGrandTotals.
func (pt *PivotTableDefinition) ColGrandTotals() bool {
	attr, found := pt.GetAttribute(
		"colGrandTotals",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetColGrandTotals sets whether to show column grand totals.
// Attribute: colGrandTotals.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetColGrandTotals(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"colGrandTotals",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "colGrandTotals", "", attrValueFalse),
		)
	}
}

// FieldPrintTitles returns whether to print field titles.
// Attribute: fieldPrintTitles.
func (pt *PivotTableDefinition) FieldPrintTitles() bool {
	attr, found := pt.GetAttribute(
		"fieldPrintTitles",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetFieldPrintTitles sets whether to print field titles.
// Attribute: fieldPrintTitles.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetFieldPrintTitles(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"fieldPrintTitles",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("fieldPrintTitles", "")
	}
}

// ItemPrintTitles returns whether to print item titles.
// Attribute: itemPrintTitles.
func (pt *PivotTableDefinition) ItemPrintTitles() bool {
	attr, found := pt.GetAttribute(
		"itemPrintTitles",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetItemPrintTitles sets whether to print item titles.
// Attribute: itemPrintTitles.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetItemPrintTitles(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"itemPrintTitles",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("itemPrintTitles", "")
	}
}

// MergeItem returns whether to merge items. Attribute: mergeItem.
func (pt *PivotTableDefinition) MergeItem() bool {
	attr, found := pt.GetAttribute(
		"mergeItem",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetMergeItem sets whether to merge items. Attribute: mergeItem.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetMergeItem(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"mergeItem",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("mergeItem", "")
	}
}

// ShowDropZones returns whether to show drop zones.
// Attribute: showDropZones.
func (pt *PivotTableDefinition) ShowDropZones() bool {
	attr, found := pt.GetAttribute(
		"showDropZones",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowDropZones sets whether to show drop zones.
// Attribute: showDropZones.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowDropZones(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"showDropZones",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "showDropZones", "", attrValueFalse),
		)
	}
}

// CreatedVersion returns the created version. Attribute: createdVersion.
func (pt *PivotTableDefinition) CreatedVersion() uint8 {
	attr, found := pt.GetAttribute(
		"createdVersion",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize8,
	)

	return uint8(val)
}

// SetCreatedVersion sets the created version. Attribute: createdVersion.
func (pt *PivotTableDefinition) SetCreatedVersion(
	version uint8,
) {
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"createdVersion",
			"",
			strconv.FormatUint(
				uint64(version),
				parseBase10,
			),
		),
	)
}

// Indent returns the indent level. Attribute: indent.
func (pt *PivotTableDefinition) Indent() uint32 {
	attr, found := pt.GetAttribute("indent", "")
	if !found {
		return 1 // Default is 1
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetIndent sets the indent level. Attribute: indent.
func (pt *PivotTableDefinition) SetIndent(
	indent uint32,
) {
	if indent == 1 {
		pt.RemoveAttribute(
			"indent",
			"",
		) // 1 is default

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"indent",
			"",
			strconv.FormatUint(
				uint64(indent),
				parseBase10,
			),
		),
	)
}

// ShowEmptyRow returns whether to show empty rows.
// Attribute: showEmptyRow.
func (pt *PivotTableDefinition) ShowEmptyRow() bool {
	attr, found := pt.GetAttribute(
		"showEmptyRow",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowEmptyRow sets whether to show empty rows.
// Attribute: showEmptyRow.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowEmptyRow(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"showEmptyRow",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("showEmptyRow", "")
	}
}

// ShowEmptyCol returns whether to show empty columns.
// Attribute: showEmptyCol.
func (pt *PivotTableDefinition) ShowEmptyCol() bool {
	attr, found := pt.GetAttribute(
		"showEmptyCol",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowEmptyCol sets whether to show empty columns.
// Attribute: showEmptyCol.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowEmptyCol(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"showEmptyCol",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("showEmptyCol", "")
	}
}

// ShowHeaders returns whether to show headers.
// Attribute: showHeaders.
func (pt *PivotTableDefinition) ShowHeaders() bool {
	attr, found := pt.GetAttribute(
		"showHeaders",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowHeaders sets whether to show headers.
// Attribute: showHeaders.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetShowHeaders(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"showHeaders",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "showHeaders", "", attrValueFalse),
		)
	}
}

// Compact returns whether to use compact layout.
// Attribute: compact.
func (pt *PivotTableDefinition) Compact() bool {
	attr, found := pt.GetAttribute("compact", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetCompact sets whether to use compact layout.
// Attribute: compact.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetCompact(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"compact",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "compact", "", attrValueFalse),
		)
	}
}

// Outline returns whether to use outline layout.
// Attribute: outline.
func (pt *PivotTableDefinition) Outline() bool {
	attr, found := pt.GetAttribute("outline", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetOutline sets whether to use outline layout.
// Attribute: outline.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetOutline(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"outline",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("outline", "")
	}
}

// OutlineData returns whether to outline data.
// Attribute: outlineData.
func (pt *PivotTableDefinition) OutlineData() bool {
	attr, found := pt.GetAttribute(
		"outlineData",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetOutlineData sets whether to outline data.
// Attribute: outlineData.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetOutlineData(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"outlineData",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("outlineData", "")
	}
}

// CompactData returns whether to use compact data.
// Attribute: compactData.
func (pt *PivotTableDefinition) CompactData() bool {
	attr, found := pt.GetAttribute(
		"compactData",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetCompactData sets whether to use compact data.
// Attribute: compactData.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetCompactData(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"compactData",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "compactData", "", attrValueFalse),
		)
	}
}

// Published returns whether the pivot table is published.
// Attribute: published.
func (pt *PivotTableDefinition) Published() bool {
	attr, found := pt.GetAttribute(
		"published",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPublished sets whether the pivot table is published.
// Attribute: published.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetPublished(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"published",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("published", "")
	}
}

// GridDropZones returns whether to show grid drop zones.
// Attribute: gridDropZones.
func (pt *PivotTableDefinition) GridDropZones() bool {
	attr, found := pt.GetAttribute(
		"gridDropZones",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetGridDropZones sets whether to show grid drop zones.
// Attribute: gridDropZones.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetGridDropZones(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"gridDropZones",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("gridDropZones", "")
	}
}

// Immersive returns whether to use immersive mode.
// Attribute: immersive.
func (pt *PivotTableDefinition) Immersive() bool {
	attr, found := pt.GetAttribute(
		"immersive",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetImmersive sets whether to use immersive mode.
// Attribute: immersive.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetImmersive(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"immersive",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "immersive", "", attrValueFalse),
		)
	}
}

// MultipleFieldFilters returns whether to allow multiple field filters.
// Attribute: multipleFieldFilters.
func (pt *PivotTableDefinition) MultipleFieldFilters() bool {
	attr, found := pt.GetAttribute(
		"multipleFieldFilters",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetMultipleFieldFilters sets whether to allow multiple field filters.
// Attribute: multipleFieldFilters.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetMultipleFieldFilters(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"multipleFieldFilters",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "multipleFieldFilters", "", attrValueFalse),
		)
	}
}

// ChartFormat returns the chart format. Attribute: chartFormat.
func (pt *PivotTableDefinition) ChartFormat() uint32 {
	attr, found := pt.GetAttribute(
		"chartFormat",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetChartFormat sets the chart format. Attribute: chartFormat.
func (pt *PivotTableDefinition) SetChartFormat(
	format uint32,
) {
	if format == 0 {
		pt.RemoveAttribute("chartFormat", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"chartFormat",
			"",
			strconv.FormatUint(
				uint64(format),
				parseBase10,
			),
		),
	)
}

// RowHeaderCaption returns the row header caption.
// Attribute: rowHeaderCaption.
func (pt *PivotTableDefinition) RowHeaderCaption() string {
	attr, found := pt.GetAttribute(
		"rowHeaderCaption",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRowHeaderCaption sets the row header caption.
// Attribute: rowHeaderCaption.
func (pt *PivotTableDefinition) SetRowHeaderCaption(
	caption string,
) {
	if caption == "" {
		pt.RemoveAttribute("rowHeaderCaption", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"rowHeaderCaption",
			"",
			caption,
		),
	)
}

// ColHeaderCaption returns the column header caption.
// Attribute: colHeaderCaption.
func (pt *PivotTableDefinition) ColHeaderCaption() string {
	attr, found := pt.GetAttribute(
		"colHeaderCaption",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetColHeaderCaption sets the column header caption.
// Attribute: colHeaderCaption.
func (pt *PivotTableDefinition) SetColHeaderCaption(
	caption string,
) {
	if caption == "" {
		pt.RemoveAttribute("colHeaderCaption", "")

		return
	}
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"colHeaderCaption",
			"",
			caption,
		),
	)
}

// FieldListSortAscending returns whether to sort field list ascending.
// Attribute: fieldListSortAscending.
func (pt *PivotTableDefinition) FieldListSortAscending() bool {
	attr, found := pt.GetAttribute(
		"fieldListSortAscending",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetFieldListSortAscending sets whether to sort field list ascending.
// Attribute: fieldListSortAscending.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetFieldListSortAscending(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"fieldListSortAscending",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("fieldListSortAscending", "")
	}
}

// MdxSubqueries returns whether to allow MDX subqueries.
// Attribute: mdxSubqueries.
func (pt *PivotTableDefinition) MdxSubqueries() bool {
	attr, found := pt.GetAttribute(
		"mdxSubqueries",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetMdxSubqueries sets whether to allow MDX subqueries.
// Attribute: mdxSubqueries.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetMdxSubqueries(
	value bool,
) {
	if value {
		pt.SetAttribute(
			openxml.NewAttribute(
				"",
				"mdxSubqueries",
				"",
				attrValueTrue,
			),
		)
	} else {
		pt.RemoveAttribute("mdxSubqueries", "")
	}
}

// CustomListSort returns whether to use custom list sort.
// Attribute: customListSort.
func (pt *PivotTableDefinition) CustomListSort() bool {
	attr, found := pt.GetAttribute(
		"customListSort",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetCustomListSort sets whether to use custom list sort.
// Attribute: customListSort.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (pt *PivotTableDefinition) SetCustomListSort(
	value bool,
) {
	if value {
		pt.RemoveAttribute(
			"customListSort",
			"",
		) // true is default
	} else {
		pt.SetAttribute(
			openxml.NewAttribute("", "customListSort", "", attrValueFalse),
		)
	}
}

// Location returns the Location child element, or nil if not present.
func (pt *PivotTableDefinition) Location() *Location {
	elem := pt.GetElement(
		"location",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if loc, ok := elem.(*Location); ok {
		return loc
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Location{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateLocation returns the Location child element, creating if needed.
func (pt *PivotTableDefinition) GetOrCreateLocation() *Location {
	loc := pt.Location()
	if loc != nil {
		return loc
	}
	loc = NewLocation()
	// Location should be first
	if first := pt.FirstChild(); first != nil {
		pt.InsertBefore(loc, first)
	} else {
		pt.AppendChild(loc)
	}

	return loc
}

// PivotFields returns the PivotFields child element, or nil if not present.
func (pt *PivotTableDefinition) PivotFields() *PivotFields {
	elem := pt.GetElement(
		"pivotFields",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if pf, ok := elem.(*PivotFields); ok {
		return pf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PivotFields{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePivotFields returns the PivotFields child element,
// creating it if needed.
func (pt *PivotTableDefinition) GetOrCreatePivotFields() *PivotFields {
	pf := pt.PivotFields()
	if pf != nil {
		return pf
	}
	pf = NewPivotFields()
	// PivotFields should come after location
	loc := pt.Location()
	if loc != nil {
		pt.InsertAfter(pf, loc)
	} else {
		pt.AppendChild(pf)
	}

	return pf
}

// RowFields returns the RowFields child element, or nil if not present.
func (pt *PivotTableDefinition) RowFields() *RowFields {
	elem := pt.GetElement(
		"rowFields",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if rf, ok := elem.(*RowFields); ok {
		return rf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RowFields{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateRowFields returns the RowFields child element,
// creating it if needed.
func (pt *PivotTableDefinition) GetOrCreateRowFields() *RowFields {
	rf := pt.RowFields()
	if rf != nil {
		return rf
	}
	rf = NewRowFields()
	pt.AppendChild(rf)

	return rf
}

// RowItems returns the RowItems child element, or nil if not present.
func (pt *PivotTableDefinition) RowItems() *RowItems {
	elem := pt.GetElement(
		"rowItems",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if ri, ok := elem.(*RowItems); ok {
		return ri
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RowItems{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateRowItems returns the RowItems child element,
// creating it if needed.
func (pt *PivotTableDefinition) GetOrCreateRowItems() *RowItems {
	ri := pt.RowItems()
	if ri != nil {
		return ri
	}
	ri = NewRowItems()
	pt.AppendChild(ri)

	return ri
}

// ColFields returns the ColFields child element, or nil if not present.
func (pt *PivotTableDefinition) ColFields() *ColFields {
	elem := pt.GetElement(
		"colFields",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if cf, ok := elem.(*ColFields); ok {
		return cf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ColFields{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateColFields returns the ColFields child element,
// creating it if needed.
func (pt *PivotTableDefinition) GetOrCreateColFields() *ColFields {
	cf := pt.ColFields()
	if cf != nil {
		return cf
	}
	cf = NewColFields()
	pt.AppendChild(cf)

	return cf
}

// ColItems returns the ColItems child element, or nil if not present.
func (pt *PivotTableDefinition) ColItems() *ColItems {
	elem := pt.GetElement(
		"colItems",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if ci, ok := elem.(*ColItems); ok {
		return ci
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ColItems{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateColItems returns the ColItems child element,
// creating it if needed.
func (pt *PivotTableDefinition) GetOrCreateColItems() *ColItems {
	ci := pt.ColItems()
	if ci != nil {
		return ci
	}
	ci = NewColItems()
	pt.AppendChild(ci)

	return ci
}

// PageFields returns the PageFields child element, or nil if not present.
func (pt *PivotTableDefinition) PageFields() *PageFields {
	elem := pt.GetElement(
		"pageFields",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if pf, ok := elem.(*PageFields); ok {
		return pf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PageFields{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePageFields returns the PageFields child element,
// creating it if needed.
func (pt *PivotTableDefinition) GetOrCreatePageFields() *PageFields {
	pf := pt.PageFields()
	if pf != nil {
		return pf
	}
	pf = NewPageFields()
	pt.AppendChild(pf)

	return pf
}

// DataFields returns the DataFields child element, or nil if not present.
func (pt *PivotTableDefinition) DataFields() *DataFields {
	elem := pt.GetElement(
		"dataFields",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if df, ok := elem.(*DataFields); ok {
		return df
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DataFields{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDataFields returns the DataFields child element,
// creating it if needed.
func (pt *PivotTableDefinition) GetOrCreateDataFields() *DataFields {
	df := pt.DataFields()
	if df != nil {
		return df
	}
	df = NewDataFields()
	pt.AppendChild(df)

	return df
}

// PivotTableStyleInfo returns the PivotTableStyleInfo child element,
// or nil if not present.
func (pt *PivotTableDefinition) PivotTableStyleInfo() *PivotTableStyleInfo {
	elem := pt.GetElement(
		"pivotTableStyleInfo",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if psi, ok := elem.(*PivotTableStyleInfo); ok {
		return psi
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &PivotTableStyleInfo{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreatePivotTableStyleInfo returns the PivotTableStyleInfo child element,
// creating it if needed.
func (pt *PivotTableDefinition) GetOrCreatePivotTableStyleInfo() *PivotTableStyleInfo { //nolint:revive // line-length-limit
	psi := pt.PivotTableStyleInfo()
	if psi != nil {
		return psi
	}
	psi = NewPivotTableStyleInfo()
	pt.AppendChild(psi)

	return psi
}

// Clone creates a deep copy of this PivotTableDefinition element.
func (pt *PivotTableDefinition) Clone() openxml.Element {
	cloned := pt.CompositeElementBase.Clone()

	return &PivotTableDefinition{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PivotTableDefinition element.
func (pt *PivotTableDefinition) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pt.CompositeElementBase.CloneNode(
		deep,
	)

	return &PivotTableDefinition{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Location represents the location element (x:location).
// It specifies the cell range that the pivot table occupies.
type Location struct {
	*openxml.LeafElementBase
}

// NewLocation creates a new Location element.
func NewLocation() *Location {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"location",
		PrefixDefault,
	)

	return &Location{LeafElementBase: elem}
}

// Ref returns the cell range reference. Attribute: ref.
func (l *Location) Ref() string {
	attr, found := l.GetAttribute("ref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the cell range reference. Attribute: ref.
func (l *Location) SetRef(ref string) {
	if ref == "" {
		l.RemoveAttribute("ref", "")

		return
	}
	l.SetAttribute(
		openxml.NewAttribute("", "ref", "", ref),
	)
}

// FirstHeaderRow returns the first header row. Attribute: firstHeaderRow.
func (l *Location) FirstHeaderRow() uint32 {
	attr, found := l.GetAttribute(
		"firstHeaderRow",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetFirstHeaderRow sets the first header row. Attribute: firstHeaderRow.
func (l *Location) SetFirstHeaderRow(row uint32) {
	l.SetAttribute(
		openxml.NewAttribute(
			"",
			"firstHeaderRow",
			"",
			strconv.FormatUint(
				uint64(row),
				parseBase10,
			),
		),
	)
}

// FirstDataRow returns the first data row. Attribute: firstDataRow.
func (l *Location) FirstDataRow() uint32 {
	attr, found := l.GetAttribute(
		"firstDataRow",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetFirstDataRow sets the first data row. Attribute: firstDataRow.
func (l *Location) SetFirstDataRow(row uint32) {
	l.SetAttribute(
		openxml.NewAttribute(
			"",
			"firstDataRow",
			"",
			strconv.FormatUint(
				uint64(row),
				parseBase10,
			),
		),
	)
}

// FirstDataCol returns the first data column. Attribute: firstDataCol.
func (l *Location) FirstDataCol() uint32 {
	attr, found := l.GetAttribute(
		"firstDataCol",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetFirstDataCol sets the first data column. Attribute: firstDataCol.
func (l *Location) SetFirstDataCol(col uint32) {
	l.SetAttribute(
		openxml.NewAttribute(
			"",
			"firstDataCol",
			"",
			strconv.FormatUint(
				uint64(col),
				parseBase10,
			),
		),
	)
}

// RowPageCount returns the row page count. Attribute: rowPageCount.
func (l *Location) RowPageCount() uint32 {
	attr, found := l.GetAttribute(
		"rowPageCount",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetRowPageCount sets the row page count. Attribute: rowPageCount.
func (l *Location) SetRowPageCount(count uint32) {
	if count == 0 {
		l.RemoveAttribute("rowPageCount", "")

		return
	}
	l.SetAttribute(
		openxml.NewAttribute(
			"",
			"rowPageCount",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// ColPageCount returns the column page count. Attribute: colPageCount.
func (l *Location) ColPageCount() uint32 {
	attr, found := l.GetAttribute(
		"colPageCount",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetColPageCount sets the column page count. Attribute: colPageCount.
func (l *Location) SetColPageCount(count uint32) {
	if count == 0 {
		l.RemoveAttribute("colPageCount", "")

		return
	}
	l.SetAttribute(
		openxml.NewAttribute(
			"",
			"colPageCount",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// Clone creates a deep copy of this Location element.
func (l *Location) Clone() openxml.Element {
	cloned := l.LeafElementBase.Clone()

	return &Location{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Location element.
func (l *Location) CloneNode(
	deep bool,
) openxml.Element {
	cloned := l.LeafElementBase.CloneNode(deep)

	return &Location{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// PivotTableStyleInfo represents the pivot table style info element
// (x:pivotTableStyleInfo).
// It specifies the pivot table style and which formatting elements are shown.
type PivotTableStyleInfo struct {
	*openxml.LeafElementBase
}

// NewPivotTableStyleInfo creates a new PivotTableStyleInfo element.
func NewPivotTableStyleInfo() *PivotTableStyleInfo {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"pivotTableStyleInfo",
		PrefixDefault,
	)

	return &PivotTableStyleInfo{
		LeafElementBase: elem,
	}
}

// Name returns the pivot table style name. Attribute: name.
func (p *PivotTableStyleInfo) Name() string {
	attr, found := p.GetAttribute(
		"name", //nolint:revive // add-constant
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the pivot table style name. Attribute: name.
func (p *PivotTableStyleInfo) SetName(
	name string,
) {
	if name == "" {
		p.RemoveAttribute("name", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// ShowRowHeaders returns whether to show row headers.
// Attribute: showRowHeaders.
func (p *PivotTableStyleInfo) ShowRowHeaders() bool {
	attr, found := p.GetAttribute(
		"showRowHeaders",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowRowHeaders sets whether to show row headers.
// Attribute: showRowHeaders.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotTableStyleInfo) SetShowRowHeaders(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"showRowHeaders",
				"",
				attrValueOne,
			),
		)
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "showRowHeaders", "", attrValueZero),
		)
	}
}

// ShowColHeaders returns whether to show column headers.
// Attribute: showColHeaders.
func (p *PivotTableStyleInfo) ShowColHeaders() bool {
	attr, found := p.GetAttribute(
		"showColHeaders",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowColHeaders sets whether to show column headers.
// Attribute: showColHeaders.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotTableStyleInfo) SetShowColHeaders(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"showColHeaders",
				"",
				attrValueOne,
			),
		)
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "showColHeaders", "", attrValueZero),
		)
	}
}

// ShowRowStripes returns whether to show row stripes.
// Attribute: showRowStripes.
func (p *PivotTableStyleInfo) ShowRowStripes() bool {
	attr, found := p.GetAttribute(
		"showRowStripes",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowRowStripes sets whether to show row stripes.
// Attribute: showRowStripes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotTableStyleInfo) SetShowRowStripes(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"showRowStripes",
				"",
				attrValueOne,
			),
		)
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "showRowStripes", "", attrValueZero),
		)
	}
}

// ShowColStripes returns whether to show column stripes.
// Attribute: showColStripes.
func (p *PivotTableStyleInfo) ShowColStripes() bool {
	attr, found := p.GetAttribute(
		"showColStripes",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowColStripes sets whether to show column stripes.
// Attribute: showColStripes.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotTableStyleInfo) SetShowColStripes(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"showColStripes",
				"",
				attrValueOne,
			),
		)
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "showColStripes", "", attrValueZero),
		)
	}
}

// ShowLastColumn returns whether to show last column formatting.
// Attribute: showLastColumn.
func (p *PivotTableStyleInfo) ShowLastColumn() bool {
	attr, found := p.GetAttribute(
		"showLastColumn",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowLastColumn sets whether to show last column formatting.
// Attribute: showLastColumn.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PivotTableStyleInfo) SetShowLastColumn(
	value bool,
) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"showLastColumn",
				"",
				attrValueOne,
			),
		)
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "showLastColumn", "", attrValueZero),
		)
	}
}

// Clone creates a deep copy of this PivotTableStyleInfo element.
func (p *PivotTableStyleInfo) Clone() openxml.Element {
	cloned := p.LeafElementBase.Clone()

	return &PivotTableStyleInfo{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this PivotTableStyleInfo element.
func (p *PivotTableStyleInfo) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.LeafElementBase.CloneNode(deep)

	return &PivotTableStyleInfo{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// RowItems represents the row items collection element (x:rowItems).
type RowItems struct {
	*openxml.CompositeElementBase
}

// NewRowItems creates a new RowItems element.
func NewRowItems() *RowItems {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"rowItems",
		PrefixDefault,
	)

	return &RowItems{CompositeElementBase: elem}
}

// Count returns the count of row items. Attribute: count.
func (ri *RowItems) Count() uint32 {
	attr, found := ri.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetCount sets the count of row items. Attribute: count.
func (ri *RowItems) SetCount(count uint32) {
	ri.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// Items returns an iterator over all I elements.
func (ri *RowItems) Items() iter.Seq[*I] {
	return func(yield func(*I) bool) {
		for child := range ri.Children() {
			if child.LocalName() != "i" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var item *I
			if i, ok := child.(*I); ok {
				item = i
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				item = &I{CompositeElementBase: comp}
			}
			if item != nil && !yield(item) {
				return
			}
		}
	}
}

// AddItem adds a new I element.
func (ri *RowItems) AddItem() *I {
	item := NewI()
	ri.AppendChild(item)

	return item
}

// Clone creates a deep copy of this RowItems element.
func (ri *RowItems) Clone() openxml.Element {
	cloned := ri.CompositeElementBase.Clone()

	return &RowItems{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this RowItems element.
func (ri *RowItems) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ri.CompositeElementBase.CloneNode(
		deep,
	)

	return &RowItems{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ColItems represents the column items collection element (x:colItems).
type ColItems struct {
	*openxml.CompositeElementBase
}

// NewColItems creates a new ColItems element.
func NewColItems() *ColItems {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"colItems",
		PrefixDefault,
	)

	return &ColItems{CompositeElementBase: elem}
}

// Count returns the count of column items. Attribute: count.
func (ci *ColItems) Count() uint32 {
	attr, found := ci.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetCount sets the count of column items. Attribute: count.
func (ci *ColItems) SetCount(count uint32) {
	ci.SetAttribute(
		openxml.NewAttribute(
			"",
			"count", //nolint:revive // add-constant: attribute name
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// Items returns an iterator over all I elements.
func (ci *ColItems) Items() iter.Seq[*I] {
	return func(yield func(*I) bool) {
		for child := range ci.Children() {
			if child.LocalName() != "i" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var item *I
			if i, ok := child.(*I); ok {
				item = i
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				item = &I{CompositeElementBase: comp}
			}
			if item != nil && !yield(item) {
				return
			}
		}
	}
}

// AddItem adds a new I element.
func (ci *ColItems) AddItem() *I {
	item := NewI()
	ci.AppendChild(item)

	return item
}

// Clone creates a deep copy of this ColItems element.
func (ci *ColItems) Clone() openxml.Element {
	cloned := ci.CompositeElementBase.Clone()

	return &ColItems{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ColItems element.
func (ci *ColItems) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ci.CompositeElementBase.CloneNode(
		deep,
	)

	return &ColItems{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// I represents a row/column item element (x:i).
type I struct {
	*openxml.CompositeElementBase
}

// NewI creates a new I element.
func NewI() *I {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"i",
		PrefixDefault,
	)

	return &I{CompositeElementBase: elem}
}

// T returns the item type. Attribute: t.
func (i *I) T() string {
	attr, found := i.GetAttribute("t", "")
	if !found {
		return "data" // Default is data
	}

	return attr.Value()
}

// SetT sets the item type. Attribute: t.
func (i *I) SetT(t string) {
	if t == "" || t == "data" {
		i.RemoveAttribute(
			"t",
			"",
		) // data is default

		return
	}
	i.SetAttribute(
		openxml.NewAttribute("", "t", "", t),
	)
}

// R returns the repeat count. Attribute: r.
func (i *I) R() uint32 {
	attr, found := i.GetAttribute("r", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetR sets the repeat count. Attribute: r.
func (i *I) SetR(r uint32) {
	if r == 0 {
		i.RemoveAttribute("r", "")

		return
	}
	i.SetAttribute(
		openxml.NewAttribute(
			"",
			"r",
			"",
			strconv.FormatUint(
				uint64(r),
				parseBase10,
			),
		),
	)
}

// I returns the index. Attribute: i.
func (i *I) I() uint32 {
	attr, found := i.GetAttribute(
		"i", //nolint:revive // add-constant
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetI sets the index. Attribute: i.
func (i *I) SetI(idx uint32) {
	if idx == 0 {
		i.RemoveAttribute("i", "")

		return
	}
	i.SetAttribute(
		openxml.NewAttribute(
			"",
			"i",
			"",
			strconv.FormatUint(
				uint64(idx),
				parseBase10,
			),
		),
	)
}

// Clone creates a deep copy of this I element.
func (i *I) Clone() openxml.Element {
	cloned := i.CompositeElementBase.Clone()

	return &I{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this I element.
func (i *I) CloneNode(deep bool) openxml.Element {
	cloned := i.CompositeElementBase.CloneNode(
		deep,
	)

	return &I{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
