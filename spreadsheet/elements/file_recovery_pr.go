package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// FileRecoveryPr represents the file recovery properties element
// (x:fileRecoveryPr).
type FileRecoveryPr struct {
	*openxml.CompositeElementBase
}

// NewFileRecoveryPr creates a new FileRecoveryPr element.
func NewFileRecoveryPr() *FileRecoveryPr {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"fileRecoveryPr",
		PrefixDefault,
	)

	return &FileRecoveryPr{
		CompositeElementBase: elem,
	}
}

// AutoRecover returns whether auto-recover is enabled.
func (fr *FileRecoveryPr) AutoRecover() bool {
	attr, found := fr.GetAttribute(
		"autoRecover",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetAutoRecover sets whether auto-recover is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (fr *FileRecoveryPr) SetAutoRecover(
	value bool,
) {
	if value {
		fr.RemoveAttribute("autoRecover", "")
	} else {
		fr.SetAttribute(openxml.NewAttribute("", "autoRecover", "", attrValueFalse))
	}
}

// CrashSave returns whether this is a crash recovery save.
func (fr *FileRecoveryPr) CrashSave() bool {
	attr, found := fr.GetAttribute(
		"crashSave",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCrashSave sets whether this is a crash recovery save.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (fr *FileRecoveryPr) SetCrashSave(
	value bool,
) {
	if value {
		fr.SetAttribute(
			openxml.NewAttribute(
				"",
				"crashSave",
				"",
				attrValueTrue,
			),
		)
	} else {
		fr.RemoveAttribute("crashSave", "")
	}
}

// DataExtractLoad returns whether data was extracted during load.
func (fr *FileRecoveryPr) DataExtractLoad() bool {
	attr, found := fr.GetAttribute(
		"dataExtractLoad",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDataExtractLoad sets whether data was extracted during load.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (fr *FileRecoveryPr) SetDataExtractLoad(
	value bool,
) {
	if value {
		fr.SetAttribute(
			openxml.NewAttribute(
				"",
				"dataExtractLoad",
				"",
				attrValueTrue,
			),
		)
	} else {
		fr.RemoveAttribute("dataExtractLoad", "")
	}
}

// RepairLoad returns whether the file was repaired during load.
func (fr *FileRecoveryPr) RepairLoad() bool {
	attr, found := fr.GetAttribute(
		"repairLoad",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetRepairLoad sets whether the file was repaired during load.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (fr *FileRecoveryPr) SetRepairLoad(
	value bool,
) {
	if value {
		fr.SetAttribute(
			openxml.NewAttribute(
				"",
				"repairLoad",
				"",
				attrValueTrue,
			),
		)
	} else {
		fr.RemoveAttribute("repairLoad", "")
	}
}

// Clone creates a deep copy of this FileRecoveryPr element.
func (fr *FileRecoveryPr) Clone() openxml.Element {
	cloned := fr.CompositeElementBase.Clone()

	return &FileRecoveryPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this FileRecoveryPr element.
func (fr *FileRecoveryPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := fr.CompositeElementBase.CloneNode(
		deep,
	)

	return &FileRecoveryPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
