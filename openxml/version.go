package openxml

// FileFormatVersion represents the Office file format version.
// This enum includes all Office versions from 2007 onwards.
type FileFormatVersion int

const (
	// FileFormatVersionOffice2007 represents Microsoft Office 2007 (ECMA-376 1st edition).
	FileFormatVersionOffice2007 FileFormatVersion = iota
	// FileFormatVersionOffice2010 represents Microsoft Office 2010 (ECMA-376 2nd edition).
	FileFormatVersionOffice2010
	// FileFormatVersionOffice2013 represents Microsoft Office 2013 (ECMA-376 3rd edition).
	FileFormatVersionOffice2013
	// FileFormatVersionOffice2016 represents Microsoft Office 2016 (ECMA-376 5th edition).
	FileFormatVersionOffice2016
	// FileFormatVersionOffice2019 represents Microsoft Office 2019.
	FileFormatVersionOffice2019
	// FileFormatVersionOffice2021 represents Microsoft Office 2021.
	FileFormatVersionOffice2021
	// FileFormatVersionOffice2022 represents Microsoft Office 2022.
	FileFormatVersionOffice2022
	// FileFormatVersionOffice2023 represents Microsoft Office 2023.
	FileFormatVersionOffice2023
	// FileFormatVersionOffice2024 represents Microsoft Office 2024.
	FileFormatVersionOffice2024
	// FileFormatVersionOffice2025 represents Microsoft Office 2025.
	FileFormatVersionOffice2025
	// FileFormatVersionMicrosoft365 represents the current Microsoft 365 (formerly Office 365).
	FileFormatVersionMicrosoft365
)

// String returns the string representation of the version.
func (v FileFormatVersion) String() string {
	switch v {
	case FileFormatVersionOffice2007:
		return "Office2007"
	case FileFormatVersionOffice2010:
		return "Office2010"
	case FileFormatVersionOffice2013:
		return "Office2013"
	case FileFormatVersionOffice2016:
		return "Office2016"
	case FileFormatVersionOffice2019:
		return "Office2019"
	case FileFormatVersionOffice2021:
		return "Office2021"
	case FileFormatVersionOffice2022:
		return "Office2022"
	case FileFormatVersionOffice2023:
		return "Office2023"
	case FileFormatVersionOffice2024:
		return "Office2024"
	case FileFormatVersionOffice2025:
		return "Office2025"
	case FileFormatVersionMicrosoft365:
		return "Microsoft365"
	default:
		return "Unknown"
	}
}

// Description returns a longer description of the version.
func (v FileFormatVersion) Description() string {
	switch v {
	case FileFormatVersionOffice2007:
		return "Microsoft Office 2007 (ECMA-376 1st Edition)"
	case FileFormatVersionOffice2010:
		return "Microsoft Office 2010 (ECMA-376 2nd Edition)"
	case FileFormatVersionOffice2013:
		return "Microsoft Office 2013 (ECMA-376 3rd Edition)"
	case FileFormatVersionOffice2016:
		return "Microsoft Office 2016 (ECMA-376 5th Edition)"
	case FileFormatVersionOffice2019:
		return "Microsoft Office 2019"
	case FileFormatVersionOffice2021:
		return "Microsoft Office 2021"
	case FileFormatVersionOffice2022:
		return "Microsoft Office 2022"
	case FileFormatVersionOffice2023:
		return "Microsoft Office 2023"
	case FileFormatVersionOffice2024:
		return "Microsoft Office 2024"
	case FileFormatVersionOffice2025:
		return "Microsoft Office 2025"
	case FileFormatVersionMicrosoft365:
		return "Microsoft 365 (current)"
	default:
		return "Unknown Office Version"
	}
}

// AllFileFormatVersions returns all supported Office versions in chronological order.
var AllFileFormatVersions = []FileFormatVersion{
	FileFormatVersionOffice2007,
	FileFormatVersionOffice2010,
	FileFormatVersionOffice2013,
	FileFormatVersionOffice2016,
	FileFormatVersionOffice2019,
	FileFormatVersionOffice2021,
	FileFormatVersionOffice2022,
	FileFormatVersionOffice2023,
	FileFormatVersionOffice2024,
	FileFormatVersionOffice2025,
	FileFormatVersionMicrosoft365,
}
