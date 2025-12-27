package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// RevisionsPassword returns the legacy revision password hash.
func (wp *WorkbookProtection) RevisionsPassword() string {
	attr, found := wp.GetAttribute(
		attrRevisionsPassword,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRevisionsPassword sets the legacy revision password hash.
func (wp *WorkbookProtection) SetRevisionsPassword(
	hash string,
) {
	if hash == "" {
		wp.RemoveAttribute(
			attrRevisionsPassword,
			"",
		)

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			attrRevisionsPassword,
			"",
			hash,
		),
	)
}

// RevisionsAlgorithmName returns the hash algorithm used for revision
// protection.
func (wp *WorkbookProtection) RevisionsAlgorithmName() string {
	attr, found := wp.GetAttribute(
		attrRevisionsAlgorithmName,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRevisionsAlgorithmName sets the hash algorithm for revision protection.
func (wp *WorkbookProtection) SetRevisionsAlgorithmName(
	algorithm string,
) {
	if algorithm == "" {
		wp.RemoveAttribute(
			attrRevisionsAlgorithmName,
			"",
		)

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			attrRevisionsAlgorithmName,
			"",
			algorithm,
		),
	)
}

// RevisionsHashValue returns the base64-encoded revision password hash.
func (wp *WorkbookProtection) RevisionsHashValue() string {
	attr, found := wp.GetAttribute(
		attrRevisionsHashValue,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRevisionsHashValue sets the base64-encoded revision password hash.
func (wp *WorkbookProtection) SetRevisionsHashValue(
	hash string,
) {
	if hash == "" {
		wp.RemoveAttribute(
			attrRevisionsHashValue,
			"",
		)

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			attrRevisionsHashValue,
			"",
			hash,
		),
	)
}

// RevisionsSaltValue returns the base64-encoded revision salt value.
func (wp *WorkbookProtection) RevisionsSaltValue() string {
	attr, found := wp.GetAttribute(
		attrRevisionsSaltValue,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRevisionsSaltValue sets the base64-encoded revision salt value.
func (wp *WorkbookProtection) SetRevisionsSaltValue(
	salt string,
) {
	if salt == "" {
		wp.RemoveAttribute(
			attrRevisionsSaltValue,
			"",
		)

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			attrRevisionsSaltValue,
			"",
			salt,
		),
	)
}

// RevisionsSpinCount returns the number of revision hash iterations.
func (wp *WorkbookProtection) RevisionsSpinCount() string {
	attr, found := wp.GetAttribute(
		attrRevisionsSpinCount,
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRevisionsSpinCount sets the number of revision hash iterations.
func (wp *WorkbookProtection) SetRevisionsSpinCount(
	count string,
) {
	if count == "" {
		wp.RemoveAttribute(
			attrRevisionsSpinCount,
			"",
		)

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			attrRevisionsSpinCount,
			"",
			count,
		),
	)
}

// ClearRevisionsPassword removes all revision password protection settings.
func (wp *WorkbookProtection) ClearRevisionsPassword() {
	wp.RemoveAttribute(attrRevisionsPassword, "")
	wp.RemoveAttribute(
		attrRevisionsAlgorithmName,
		"",
	)
	wp.RemoveAttribute(attrRevisionsHashValue, "")
	wp.RemoveAttribute(attrRevisionsSaltValue, "")
	wp.RemoveAttribute(attrRevisionsSpinCount, "")
}
