package elements

//revive:disable:file-length-limit many protection properties

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// SheetProtection represents the sheet protection element (x:sheetProtection).
type SheetProtection struct {
	*openxml.CompositeElementBase
}

// NewSheetProtection creates a new SheetProtection element.
func NewSheetProtection() *SheetProtection {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"sheetProtection",
		PrefixDefault,
	)

	return &SheetProtection{
		CompositeElementBase: elem,
	}
}

// Password returns the legacy password hash.
// Note: This is the legacy, less secure password storage format.
func (sp *SheetProtection) Password() string {
	attr, found := sp.GetAttribute("password", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPassword sets the legacy password hash.
func (sp *SheetProtection) SetPassword(
	hash string,
) {
	if hash == "" {
		sp.RemoveAttribute("password", "")

		return
	}
	sp.SetAttribute(
		openxml.NewAttribute(
			"",
			"password",
			"",
			hash,
		),
	)
}

// AlgorithmName returns the hash algorithm used for protection.
// Common values: "SHA-512", "SHA-384", "SHA-256", "SHA-1", "MD5"
func (sp *SheetProtection) AlgorithmName() string {
	attr, found := sp.GetAttribute(
		"algorithmName",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAlgorithmName sets the hash algorithm for protection.
func (sp *SheetProtection) SetAlgorithmName(
	algorithm string,
) {
	if algorithm == "" {
		sp.RemoveAttribute("algorithmName", "")

		return
	}
	sp.SetAttribute(
		openxml.NewAttribute(
			"",
			"algorithmName",
			"",
			algorithm,
		),
	)
}

// HashValue returns the base64-encoded password hash value.
func (sp *SheetProtection) HashValue() string {
	attr, found := sp.GetAttribute(
		"hashValue",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetHashValue sets the base64-encoded password hash value.
func (sp *SheetProtection) SetHashValue(
	hash string,
) {
	if hash == "" {
		sp.RemoveAttribute("hashValue", "")

		return
	}
	sp.SetAttribute(
		openxml.NewAttribute(
			"",
			"hashValue",
			"",
			hash,
		),
	)
}

// SaltValue returns the base64-encoded salt value.
func (sp *SheetProtection) SaltValue() string {
	attr, found := sp.GetAttribute(
		"saltValue",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSaltValue sets the base64-encoded salt value.
func (sp *SheetProtection) SetSaltValue(
	salt string,
) {
	if salt == "" {
		sp.RemoveAttribute("saltValue", "")

		return
	}
	sp.SetAttribute(
		openxml.NewAttribute(
			"",
			"saltValue",
			"",
			salt,
		),
	)
}

// SpinCount returns the number of hash iterations.
func (sp *SheetProtection) SpinCount() string {
	attr, found := sp.GetAttribute(
		"spinCount",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSpinCount sets the number of hash iterations.
func (sp *SheetProtection) SetSpinCount(
	count string,
) {
	if count == "" {
		sp.RemoveAttribute("spinCount", "")

		return
	}
	sp.SetAttribute(
		openxml.NewAttribute(
			"",
			"spinCount",
			"",
			count,
		),
	)
}

// Sheet returns whether the sheet is protected.
func (sp *SheetProtection) Sheet() bool {
	attr, found := sp.GetAttribute("sheet", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSheet sets whether the sheet is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetSheet(value bool) {
	if value {
		sp.SetAttribute(
			openxml.NewAttribute(
				"",
				"sheet",
				"",
				attrValueTrue,
			),
		)
	} else {
		sp.RemoveAttribute("sheet", "")
	}
}

// Objects returns whether objects are protected.
func (sp *SheetProtection) Objects() bool {
	attr, found := sp.GetAttribute("objects", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetObjects sets whether objects are protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetObjects(
	value bool,
) {
	if value {
		sp.SetAttribute(
			openxml.NewAttribute(
				"",
				"objects",
				"",
				attrValueTrue,
			),
		)
	} else {
		sp.RemoveAttribute("objects", "")
	}
}

// Scenarios returns whether scenarios are protected.
func (sp *SheetProtection) Scenarios() bool {
	attr, found := sp.GetAttribute(
		"scenarios",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetScenarios sets whether scenarios are protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetScenarios(
	value bool,
) {
	if value {
		sp.SetAttribute(
			openxml.NewAttribute(
				"",
				"scenarios",
				"",
				attrValueTrue,
			),
		)
	} else {
		sp.RemoveAttribute("scenarios", "")
	}
}

// FormatCells returns whether cell formatting is protected.
func (sp *SheetProtection) FormatCells() bool {
	attr, found := sp.GetAttribute(
		"formatCells",
		"",
	)
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetFormatCells sets whether cell formatting is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetFormatCells(
	value bool,
) {
	if value {
		sp.RemoveAttribute("formatCells", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "formatCells", "", attrValueFalse))
	}
}

// FormatColumns returns whether column formatting is protected.
func (sp *SheetProtection) FormatColumns() bool {
	attr, found := sp.GetAttribute(
		"formatColumns",
		"",
	)
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetFormatColumns sets whether column formatting is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetFormatColumns(
	value bool,
) {
	if value {
		sp.RemoveAttribute("formatColumns", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "formatColumns", "", attrValueFalse))
	}
}

// FormatRows returns whether row formatting is protected.
func (sp *SheetProtection) FormatRows() bool {
	attr, found := sp.GetAttribute(
		"formatRows",
		"",
	)
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetFormatRows sets whether row formatting is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetFormatRows(
	value bool,
) {
	if value {
		sp.RemoveAttribute("formatRows", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "formatRows", "", attrValueFalse))
	}
}

// InsertColumns returns whether column insertion is protected.
func (sp *SheetProtection) InsertColumns() bool {
	attr, found := sp.GetAttribute(
		"insertColumns",
		"",
	)
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetInsertColumns sets whether column insertion is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetInsertColumns(
	value bool,
) {
	if value {
		sp.RemoveAttribute("insertColumns", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "insertColumns", "", attrValueFalse))
	}
}

// InsertRows returns whether row insertion is protected.
func (sp *SheetProtection) InsertRows() bool {
	attr, found := sp.GetAttribute(
		"insertRows",
		"",
	)
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetInsertRows sets whether row insertion is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetInsertRows(
	value bool,
) {
	if value {
		sp.RemoveAttribute("insertRows", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "insertRows", "", attrValueFalse))
	}
}

// InsertHyperlinks returns whether hyperlink insertion is protected.
func (sp *SheetProtection) InsertHyperlinks() bool {
	attr, found := sp.GetAttribute(
		"insertHyperlinks",
		"",
	)
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetInsertHyperlinks sets whether hyperlink insertion is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetInsertHyperlinks(
	value bool,
) {
	if value {
		sp.RemoveAttribute("insertHyperlinks", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "insertHyperlinks", "", attrValueFalse))
	}
}

// DeleteColumns returns whether column deletion is protected.
func (sp *SheetProtection) DeleteColumns() bool {
	attr, found := sp.GetAttribute(
		"deleteColumns",
		"",
	)
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetDeleteColumns sets whether column deletion is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetDeleteColumns(
	value bool,
) {
	if value {
		sp.RemoveAttribute("deleteColumns", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "deleteColumns", "", attrValueFalse))
	}
}

// DeleteRows returns whether row deletion is protected.
func (sp *SheetProtection) DeleteRows() bool {
	attr, found := sp.GetAttribute(
		"deleteRows",
		"",
	)
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetDeleteRows sets whether row deletion is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetDeleteRows(
	value bool,
) {
	if value {
		sp.RemoveAttribute("deleteRows", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "deleteRows", "", attrValueFalse))
	}
}

// SelectLockedCells returns whether selecting locked cells is protected.
func (sp *SheetProtection) SelectLockedCells() bool {
	attr, found := sp.GetAttribute(
		"selectLockedCells",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSelectLockedCells sets whether selecting locked cells is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetSelectLockedCells(
	value bool,
) {
	if value {
		sp.SetAttribute(
			openxml.NewAttribute(
				"",
				"selectLockedCells",
				"",
				attrValueTrue,
			),
		)
	} else {
		sp.RemoveAttribute("selectLockedCells", "")
	}
}

// Sort returns whether sorting is protected.
func (sp *SheetProtection) Sort() bool {
	attr, found := sp.GetAttribute("sort", "")
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetSort sets whether sorting is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetSort(value bool) {
	if value {
		sp.RemoveAttribute("sort", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "sort", "", attrValueFalse))
	}
}

// AutoFilter returns whether auto filter is protected.
func (sp *SheetProtection) AutoFilter() bool {
	attr, found := sp.GetAttribute(
		"autoFilter",
		"",
	)
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetAutoFilter sets whether auto filter is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetAutoFilter(
	value bool,
) {
	if value {
		sp.RemoveAttribute("autoFilter", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "autoFilter", "", attrValueFalse))
	}
}

// PivotTables returns whether pivot tables are protected.
func (sp *SheetProtection) PivotTables() bool {
	attr, found := sp.GetAttribute(
		"pivotTables",
		"",
	)
	if !found {
		return true // Default is true (protected)
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetPivotTables sets whether pivot tables are protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetPivotTables(
	value bool,
) {
	if value {
		sp.RemoveAttribute("pivotTables", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "pivotTables", "", attrValueFalse))
	}
}

// SelectUnlockedCells returns whether selecting unlocked cells is protected.
func (sp *SheetProtection) SelectUnlockedCells() bool {
	attr, found := sp.GetAttribute(
		"selectUnlockedCells",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSelectUnlockedCells sets whether selecting unlocked cells is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetProtection) SetSelectUnlockedCells(
	value bool,
) {
	if value {
		sp.SetAttribute(
			openxml.NewAttribute(
				"",
				"selectUnlockedCells",
				"",
				attrValueTrue,
			),
		)
	} else {
		sp.RemoveAttribute("selectUnlockedCells", "")
	}
}

// SetSecureProtection sets up secure password protection using hash-based
// storage.
func (sp *SheetProtection) SetSecureProtection(
	algorithm, hashValue, saltValue, spinCount string,
) {
	sp.SetAlgorithmName(algorithm)
	sp.SetHashValue(hashValue)
	sp.SetSaltValue(saltValue)
	sp.SetSpinCount(spinCount)
}

// ClearPassword removes all password protection settings.
func (sp *SheetProtection) ClearPassword() {
	sp.RemoveAttribute(attrNamePassword, "")
	sp.RemoveAttribute(attrNameAlgorithmName, "")
	sp.RemoveAttribute(attrNameHashValue, "")
	sp.RemoveAttribute(attrNameSaltValue, "")
	sp.RemoveAttribute(attrNameSpinCount, "")
}

// Clone creates a deep copy of this SheetProtection element.
func (sp *SheetProtection) Clone() openxml.Element {
	cloned := sp.CompositeElementBase.Clone()

	return &SheetProtection{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SheetProtection element.
func (sp *SheetProtection) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sp.CompositeElementBase.CloneNode(
		deep,
	)

	return &SheetProtection{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
