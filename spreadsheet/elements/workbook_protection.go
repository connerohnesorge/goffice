package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Workbook protection attribute name constants.
const (
	attrWorkbookPassword       = "workbookPassword"
	attrWorkbookAlgorithmName  = "workbookAlgorithmName"
	attrWorkbookHashValue      = "workbookHashValue"
	attrWorkbookSaltValue      = "workbookSaltValue"
	attrWorkbookSpinCount      = "workbookSpinCount"
	attrRevisionsPassword      = "revisionsPassword"
	attrRevisionsAlgorithmName = "revisionsAlgorithmName"
	attrRevisionsHashValue     = "revisionsHashValue"
	attrRevisionsSaltValue     = "revisionsSaltValue"
	attrRevisionsSpinCount     = "revisionsSpinCount"
)

// WorkbookProtection represents the workbook protection element
// (x:workbookProtection).
type WorkbookProtection struct {
	*openxml.CompositeElementBase
}

// NewWorkbookProtection creates a new WorkbookProtection element.
func NewWorkbookProtection() *WorkbookProtection {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"workbookProtection",
		PrefixDefault,
	)

	return &WorkbookProtection{
		CompositeElementBase: elem,
	}
}

// LockStructure returns whether the workbook structure is locked.
// When locked, sheets cannot be added, deleted, hidden, or renamed.
func (wp *WorkbookProtection) LockStructure() bool {
	attr, found := wp.GetAttribute(
		"lockStructure",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetLockStructure sets whether the workbook structure is locked.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wp *WorkbookProtection) SetLockStructure(
	value bool,
) {
	if value {
		wp.SetAttribute(
			openxml.NewAttribute(
				"",
				"lockStructure",
				"",
				attrValueTrue,
			),
		)
	} else {
		wp.RemoveAttribute("lockStructure", "")
	}
}

// LockWindows returns whether workbook windows are locked.
// When locked, windows cannot be moved or resized.
func (wp *WorkbookProtection) LockWindows() bool {
	attr, found := wp.GetAttribute(
		"lockWindows",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetLockWindows sets whether workbook windows are locked.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wp *WorkbookProtection) SetLockWindows(
	value bool,
) {
	if value {
		wp.SetAttribute(
			openxml.NewAttribute(
				"",
				"lockWindows",
				"",
				attrValueTrue,
			),
		)
	} else {
		wp.RemoveAttribute("lockWindows", "")
	}
}

// LockRevision returns whether revision history is locked.
func (wp *WorkbookProtection) LockRevision() bool {
	attr, found := wp.GetAttribute(
		"lockRevision",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetLockRevision sets whether revision history is locked.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wp *WorkbookProtection) SetLockRevision(
	value bool,
) {
	if value {
		wp.SetAttribute(
			openxml.NewAttribute(
				"",
				"lockRevision",
				"",
				attrValueTrue,
			),
		)
	} else {
		wp.RemoveAttribute("lockRevision", "")
	}
}

// WorkbookPassword returns the legacy password hash.
// Note: This is the legacy, less secure password storage format.
func (wp *WorkbookProtection) WorkbookPassword() string {
	attr, found := wp.GetAttribute(
		attrWorkbookPassword,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetWorkbookPassword sets the legacy password hash.
// Note: Consider using the more secure hash-based password methods.
func (wp *WorkbookProtection) SetWorkbookPassword(
	hash string,
) {
	if hash == "" {
		wp.RemoveAttribute(
			attrWorkbookPassword,
			"",
		)

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			attrWorkbookPassword,
			"",
			hash,
		),
	)
}

// WorkbookAlgorithmName returns the hash algorithm used for workbook
// protection. Common values: "SHA-512", "SHA-384", "SHA-256", "SHA-1", "MD5"
func (wp *WorkbookProtection) WorkbookAlgorithmName() string {
	attr, found := wp.GetAttribute(
		attrWorkbookAlgorithmName,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetWorkbookAlgorithmName sets the hash algorithm for workbook protection.
func (wp *WorkbookProtection) SetWorkbookAlgorithmName(
	algorithm string,
) {
	if algorithm == "" {
		wp.RemoveAttribute(
			attrWorkbookAlgorithmName,
			"",
		)

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			attrWorkbookAlgorithmName,
			"",
			algorithm,
		),
	)
}

// WorkbookHashValue returns the base64-encoded password hash value.
func (wp *WorkbookProtection) WorkbookHashValue() string {
	attr, found := wp.GetAttribute(
		attrWorkbookHashValue,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetWorkbookHashValue sets the base64-encoded password hash value.
func (wp *WorkbookProtection) SetWorkbookHashValue(
	hash string,
) {
	if hash == "" {
		wp.RemoveAttribute(
			attrWorkbookHashValue,
			"",
		)

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			attrWorkbookHashValue,
			"",
			hash,
		),
	)
}

// WorkbookSaltValue returns the base64-encoded salt value.
func (wp *WorkbookProtection) WorkbookSaltValue() string {
	attr, found := wp.GetAttribute(
		attrWorkbookSaltValue,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetWorkbookSaltValue sets the base64-encoded salt value.
func (wp *WorkbookProtection) SetWorkbookSaltValue(
	salt string,
) {
	if salt == "" {
		wp.RemoveAttribute(
			attrWorkbookSaltValue,
			"",
		)

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			attrWorkbookSaltValue,
			"",
			salt,
		),
	)
}

// WorkbookSpinCount returns the number of hash iterations.
func (wp *WorkbookProtection) WorkbookSpinCount() string {
	attr, found := wp.GetAttribute(
		attrWorkbookSpinCount,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetWorkbookSpinCount sets the number of hash iterations.
func (wp *WorkbookProtection) SetWorkbookSpinCount(
	count string,
) {
	if count == "" {
		wp.RemoveAttribute(
			attrWorkbookSpinCount,
			"",
		)

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			attrWorkbookSpinCount,
			"",
			count,
		),
	)
}

// SetSecureProtection sets up secure password protection using hash-based
// storage. algorithm should be one of: "SHA-512", "SHA-384", "SHA-256",
// "SHA-1", "MD5"
func (wp *WorkbookProtection) SetSecureProtection(
	algorithm, hashValue, saltValue, spinCount string,
) {
	wp.SetWorkbookAlgorithmName(algorithm)
	wp.SetWorkbookHashValue(hashValue)
	wp.SetWorkbookSaltValue(saltValue)
	wp.SetWorkbookSpinCount(spinCount)
}

// ClearWorkbookPassword removes all workbook password protection settings.
func (wp *WorkbookProtection) ClearWorkbookPassword() {
	wp.RemoveAttribute(attrWorkbookPassword, "")
	wp.RemoveAttribute(
		attrWorkbookAlgorithmName,
		"",
	)
	wp.RemoveAttribute(attrWorkbookHashValue, "")
	wp.RemoveAttribute(attrWorkbookSaltValue, "")
	wp.RemoveAttribute(attrWorkbookSpinCount, "")
}

// Clone creates a deep copy of this WorkbookProtection element.
func (wp *WorkbookProtection) Clone() openxml.Element {
	cloned := wp.CompositeElementBase.Clone()

	return &WorkbookProtection{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this WorkbookProtection element.
func (wp *WorkbookProtection) CloneNode(
	deep bool,
) openxml.Element {
	cloned := wp.CompositeElementBase.CloneNode(
		deep,
	)

	return &WorkbookProtection{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
