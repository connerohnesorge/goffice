//nolint:revive // line-length-limit: documentation comments
package wordprocessing

// OpenSettings contains options for opening or creating documents.
type OpenSettings struct {
	// AutoSave determines if the document should be automatically saved
	// when Close() is called and there are unsaved changes.
	AutoSave bool

	// MaxCharactersInPart sets the maximum number of characters allowed
	// in a single part. A value of 0 means no limit.
	MaxCharactersInPart int64

	// MarkupCompatibilityProcessSettings controls how markup compatibility
	// elements are processed. Currently not implemented.
	MarkupCompatibilityProcessSettings MarkupCompatibilityProcessSettings
}

// MarkupCompatibilityProcessSettings controls markup compatibility processing.
type MarkupCompatibilityProcessSettings struct {
	// ProcessModeValues determines which mc:ProcessContent elements to process.
	ProcessModeValues ProcessMode

	// TargetFileFormatVersions specifies which Office versions to target.
	TargetFileFormatVersions FileFormatVersion
}

// ProcessMode specifies how to process markup compatibility content.
type ProcessMode int

const (
	// ProcessModeNoProcess does not process mc:ProcessContent.
	ProcessModeNoProcess ProcessMode = iota

	// ProcessModeProcessLoadedPartsOnly processes only loaded parts.
	ProcessModeProcessLoadedPartsOnly

	// ProcessModeProcessAllParts processes all parts.
	ProcessModeProcessAllParts
)

// FileFormatVersion represents the target Office version.
type FileFormatVersion int

const (
	// FileFormatVersionOffice2007 targets Office 2007 format.
	FileFormatVersionOffice2007 FileFormatVersion = iota

	// FileFormatVersionOffice2010 targets Office 2010 format.
	FileFormatVersionOffice2010

	// FileFormatVersionOffice2013 targets Office 2013 format.
	FileFormatVersionOffice2013

	// FileFormatVersionOffice2016 targets Office 2016 format.
	FileFormatVersionOffice2016

	// FileFormatVersionOffice2019 targets Office 2019 format.
	FileFormatVersionOffice2019

	// FileFormatVersionOffice2021 targets Office 2021 format.
	FileFormatVersionOffice2021

	// FileFormatVersionMicrosoft365 targets Microsoft 365 format.
	FileFormatVersionMicrosoft365
)

// DefaultOpenSettings returns the default settings for opening documents.
func DefaultOpenSettings() *OpenSettings {
	return &OpenSettings{
		AutoSave:            false,
		MaxCharactersInPart: 0, // No limit
		MarkupCompatibilityProcessSettings: MarkupCompatibilityProcessSettings{
			ProcessModeValues:        ProcessModeNoProcess,
			TargetFileFormatVersions: FileFormatVersionOffice2016,
		},
	}
}

// WithAutoSave returns a copy of the settings with AutoSave enabled.
func (s *OpenSettings) WithAutoSave(
	autoSave bool,
) *OpenSettings {
	settings := *s
	settings.AutoSave = autoSave

	return &settings
}

// WithMaxCharacters returns a copy of the settings with the specified max.
func (s *OpenSettings) WithMaxCharacters(
	maxChars int64,
) *OpenSettings {
	settings := *s
	settings.MaxCharactersInPart = maxChars

	return &settings
}

// WithTargetVersion returns a copy of the settings with the target version.
func (s *OpenSettings) WithTargetVersion(
	version FileFormatVersion,
) *OpenSettings {
	settings := *s
	settings.MarkupCompatibilityProcessSettings.TargetFileFormatVersions = version

	return &settings
}
