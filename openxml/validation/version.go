package validation

// FileFormatVersions represents the Office file format version for validation.
// This type is kept for backward compatibility; prefer features.FileFormatVersion
// for new code.
type FileFormatVersions int

const (
	// Office2007 represents Microsoft Office 2007 (ECMA-376 1st edition).
	Office2007 FileFormatVersions = 1 << iota
	// Office2010 represents Microsoft Office 2010 (ISO/IEC 29500:2008).
	Office2010
	// Office2013 represents Microsoft Office 2013.
	Office2013
	// Office2016 represents Microsoft Office 2016 (ECMA-376 5th edition).
	Office2016
	// Office2019 represents Microsoft Office 2019.
	Office2019
	// Office2021 represents Microsoft Office 2021.
	Office2021
	// Office2022 represents Microsoft Office 2022.
	Office2022
	// Office2023 represents Microsoft Office 2023.
	Office2023
	// Office2024 represents Microsoft Office 2024.
	Office2024
	// Office2025 represents Microsoft Office 2025.
	Office2025
	// Microsoft365 represents the current Microsoft 365 (formerly Office 365).
	Microsoft365
)

// String returns the string representation of the version.
func (v FileFormatVersions) String() string {
	switch v {
	case Office2007:
		return "Office2007"
	case Office2010:
		return "Office2010"
	case Office2013:
		return "Office2013"
	case Office2016:
		return "Office2016"
	case Office2019:
		return "Office2019"
	case Office2021:
		return "Office2021"
	case Office2022:
		return "Office2022"
	case Office2023:
		return "Office2023"
	case Office2024:
		return "Office2024"
	case Office2025:
		return "Office2025"
	case Microsoft365:
		return "Microsoft365"
	default:
		return "Unknown"
	}
}

// Description returns a longer description of the version.
func (v FileFormatVersions) Description() string {
	switch v {
	case Office2007:
		return "Microsoft Office 2007 (ECMA-376 1st Edition)"
	case Office2010:
		return "Microsoft Office 2010 (ISO/IEC 29500:2008)"
	case Office2013:
		return "Microsoft Office 2013"
	case Office2016:
		return "Microsoft Office 2016 (ECMA-376 5th Edition)"
	case Office2019:
		return "Microsoft Office 2019"
	case Office2021:
		return "Microsoft Office 2021"
	case Office2022:
		return "Microsoft Office 2022"
	case Office2023:
		return "Microsoft Office 2023"
	case Office2024:
		return "Microsoft Office 2024"
	case Office2025:
		return "Microsoft Office 2025"
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

// Includes returns true if this version includes features from the given
// version. Higher versions include features from lower versions.
func (v FileFormatVersions) Includes(
	other FileFormatVersions,
) bool {
	return v >= other
}

// VersionAvailability tracks the Office version availability of elements
// and attributes.
type VersionAvailability struct {
	// IntroducedIn is the first Office version where this feature
	// was available.
	IntroducedIn FileFormatVersions
	// DeprecatedIn is the version where this feature was deprecated
	// (0 if not deprecated).
	DeprecatedIn FileFormatVersions
	// RemovedIn is the version where this feature was removed
	// (0 if not removed).
	RemovedIn FileFormatVersions
}

// NewVersionAvailability creates a new version availability starting from
// the given version.
func NewVersionAvailability(
	introducedIn FileFormatVersions,
) *VersionAvailability {
	return &VersionAvailability{
		IntroducedIn: introducedIn,
	}
}

// Since2007 returns a VersionAvailability for features available since
// Office 2007 (ECMA-376 1st edition).
func Since2007() *VersionAvailability {
	return NewVersionAvailability(Office2007)
}

// Since2010 returns a VersionAvailability for features available since
// Office 2010.
func Since2010() *VersionAvailability {
	return NewVersionAvailability(Office2010)
}

// Since2013 returns a VersionAvailability for features available since
// Office 2013.
func Since2013() *VersionAvailability {
	return NewVersionAvailability(Office2013)
}

// Since2016 returns a VersionAvailability for features available since
// Office 2016.
func Since2016() *VersionAvailability {
	return NewVersionAvailability(Office2016)
}

// Since2019 returns a VersionAvailability for features available since
// Office 2019.
func Since2019() *VersionAvailability {
	return NewVersionAvailability(Office2019)
}

// Since2021 returns a VersionAvailability for features available since
// Office 2021.
func Since2021() *VersionAvailability {
	return NewVersionAvailability(Office2021)
}

// Since2022 returns a VersionAvailability for features available since
// Office 2022.
func Since2022() *VersionAvailability {
	return NewVersionAvailability(Office2022)
}

// Since2023 returns a VersionAvailability for features available since
// Office 2023.
func Since2023() *VersionAvailability {
	return NewVersionAvailability(Office2023)
}

// Since2024 returns a VersionAvailability for features available since
// Office 2024.
func Since2024() *VersionAvailability {
	return NewVersionAvailability(Office2024)
}

// Since2025 returns a VersionAvailability for features available since
// Office 2025.
func Since2025() *VersionAvailability {
	return NewVersionAvailability(Office2025)
}

// Since365 returns a VersionAvailability for features available only in
// Microsoft 365.
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

// IsDeprecatedIn returns true if the feature is deprecated in the given
// version.
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

// DefaultVersion is the default version used for validation when none is
// specified.
const DefaultVersion = Microsoft365

// AllVersions represents all supported Office versions (2007-2025 and M365).
var AllVersions = []FileFormatVersions{
	Office2007,
	Office2010,
	Office2013,
	Office2016,
	Office2019,
	Office2021,
	Office2022,
	Office2023,
	Office2024,
	Office2025,
	Microsoft365,
}

// VersionedElement is an interface for elements that have version-specific
// availability.
type VersionedElement interface {
	// Availability returns the version availability for this element.
	Availability() *VersionAvailability
}

// VersionedAttribute represents an attribute with version availability
// information.
type VersionedAttribute interface {
	// AttributeAvailability returns the version availability for this
	// attribute.
	AttributeAvailability() *VersionAvailability
}

// CheckElementVersion checks if an element is available in the given version.
// Returns a validation error if not available.
func CheckElementVersion(
	element any,
	version FileFormatVersions,
	path string,
) *ValidationError {
	if ve, ok := element.(VersionedElement); ok {
		avail := ve.Availability()
		if avail != nil &&
			!avail.IsAvailableIn(version) {
			return NewValidationError(
				SchemaElementNotAvailable,
				"Element is not available in "+version.String(),
				path,
				element,
			)
		}
	}

	return nil
}
