package validation

// FileFormatVersions represents the Office file format version for validation.
// This SDK only supports Office 2016 and later versions.
type FileFormatVersions int

const (
	// Office2016 represents Microsoft Office 2016 (ECMA-376 5th edition).
	Office2016 FileFormatVersions = 1 << iota
	// Office2019 represents Microsoft Office 2019.
	Office2019
	// Office2021 represents Microsoft Office 2021.
	Office2021
	// Microsoft365 represents the current Microsoft 365 (formerly Office 365).
	Microsoft365
)

// String returns the string representation of the version.
func (v FileFormatVersions) String() string {
	switch v {
	case Office2016:
		return "Office2016"
	case Office2019:
		return "Office2019"
	case Office2021:
		return "Office2021"
	case Microsoft365:
		return "Microsoft365"
	default:
		return "Unknown"
	}
}

// Description returns a longer description of the version.
func (v FileFormatVersions) Description() string {
	switch v {
	case Office2016:
		return "Microsoft Office 2016 (ECMA-376 5th Edition)"
	case Office2019:
		return "Microsoft Office 2019"
	case Office2021:
		return "Microsoft Office 2021"
	case Microsoft365:
		return "Microsoft 365 (current)"
	default:
		return "Unknown Office Version"
	}
}

// AtLeast returns true if this version is at least the given version.
func (v FileFormatVersions) AtLeast(
	other FileFormatVersions,
) bool {
	return v >= other
}

// AtMost returns true if this version is at most the given version.
func (v FileFormatVersions) AtMost(
	other FileFormatVersions,
) bool {
	return v <= other
}

// Includes returns true if this version includes features from the given version.
// Higher versions include features from lower versions.
func (v FileFormatVersions) Includes(
	other FileFormatVersions,
) bool {
	return v >= other
}

// VersionAvailability tracks the Office version availability of elements and attributes.
type VersionAvailability struct {
	// IntroducedIn is the first Office version where this feature was available.
	IntroducedIn FileFormatVersions
	// DeprecatedIn is the version where this feature was deprecated (0 if not deprecated).
	DeprecatedIn FileFormatVersions
	// RemovedIn is the version where this feature was removed (0 if not removed).
	RemovedIn FileFormatVersions
}

// NewVersionAvailability creates a new version availability starting from the given version.
func NewVersionAvailability(
	introducedIn FileFormatVersions,
) *VersionAvailability {
	return &VersionAvailability{
		IntroducedIn: introducedIn,
	}
}

// Since returns a VersionAvailability for features available since Office 2016 (all supported versions).
func Since2016() *VersionAvailability {
	return NewVersionAvailability(Office2016)
}

// Since2019 returns a VersionAvailability for features available since Office 2019.
func Since2019() *VersionAvailability {
	return NewVersionAvailability(Office2019)
}

// Since2021 returns a VersionAvailability for features available since Office 2021.
func Since2021() *VersionAvailability {
	return NewVersionAvailability(Office2021)
}

// Since365 returns a VersionAvailability for features available only in Microsoft 365.
func Since365() *VersionAvailability {
	return NewVersionAvailability(Microsoft365)
}

// IsAvailableIn returns true if the feature is available in the given version.
func (va *VersionAvailability) IsAvailableIn(
	version FileFormatVersions,
) bool {
	if va == nil {
		// If no availability info, assume it's available in all versions
		return true
	}

	// Check if introduced in this version or earlier
	if version < va.IntroducedIn {
		return false
	}

	// Check if removed in this version or earlier
	if va.RemovedIn != 0 &&
		version >= va.RemovedIn {
		return false
	}

	return true
}

// IsDeprecatedIn returns true if the feature is deprecated in the given version.
func (va *VersionAvailability) IsDeprecatedIn(
	version FileFormatVersions,
) bool {
	if va == nil || va.DeprecatedIn == 0 {
		return false
	}
	return version >= va.DeprecatedIn
}

// Deprecated marks this feature as deprecated starting from the given version.
func (va *VersionAvailability) Deprecated(
	version FileFormatVersions,
) *VersionAvailability {
	va.DeprecatedIn = version
	return va
}

// Removed marks this feature as removed starting from the given version.
func (va *VersionAvailability) Removed(
	version FileFormatVersions,
) *VersionAvailability {
	va.RemovedIn = version
	return va
}

// DefaultVersion is the default version used for validation when none is specified.
const DefaultVersion = Microsoft365

// AllVersions represents all supported Office versions.
var AllVersions = []FileFormatVersions{
	Office2016,
	Office2019,
	Office2021,
	Microsoft365,
}

// VersionedElement is an interface for elements that have version-specific availability.
type VersionedElement interface {
	// Availability returns the version availability for this element.
	Availability() *VersionAvailability
}

// VersionedAttribute represents an attribute with version availability information.
type VersionedAttribute interface {
	// AttributeAvailability returns the version availability for this attribute.
	AttributeAvailability() *VersionAvailability
}

// CheckElementVersion checks if an element is available in the given version.
// Returns a validation error if not available.
func CheckElementVersion(
	element interface{},
	version FileFormatVersions,
	path string,
) *ValidationError {
	if ve, ok := element.(VersionedElement); ok {
		avail := ve.Availability()
		if avail != nil &&
			!avail.IsAvailableIn(version) {
			return NewValidationError(
				Schema_ElementNotAvailable,
				"Element is not available in "+version.String(),
				path,
				element,
			)
		}
	}
	return nil
}
