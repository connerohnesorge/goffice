package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// FileFormatVersion represents an Office file format version.
type FileFormatVersion int

const (
	// Office2007 represents Microsoft Office 2007 (Office Open XML).
	Office2007 FileFormatVersion = iota
	// Office2010 represents Microsoft Office 2010.
	Office2010
	// Office2013 represents Microsoft Office 2013.
	Office2013
	// Office2016 represents Microsoft Office 2016.
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
	// Microsoft365 represents Microsoft 365 (cloud version).
	Microsoft365
)

// String returns the string representation of the file format version.
func (v FileFormatVersion) String() string {
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

// SchemaFileMetadata represents metadata about an OpenXML schema file.
type SchemaFileMetadata struct {
	// Path is the full file path to the schema JSON file.
	Path string
	// Namespace is the target XML namespace from the schema.
	Namespace string
	// NamespacePrefix is the derived namespace prefix (x14, x15, a14, etc.).
	NamespacePrefix string
	// Version is the Office version this schema targets.
	Version FileFormatVersion
	// TargetApp is the application: wordprocessing, spreadsheet, or presentation.
	TargetApp string
}

// parseSchemaFile extracts metadata from a schema JSON file.
// It determines version, target app, namespace, and prefix from the filename and content.
func parseSchemaFile(
	path string,
) (*SchemaFileMetadata, error) {
	filename := filepath.Base(path)

	// Extract version from filename
	version := extractVersionFromFilename(
		filename,
	)

	// Determine target app from filename patterns
	targetApp := determineTargetApp(filename)

	// Load the JSON schema to get the actual namespace
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read schema file %s: %w",
			path,
			err,
		)
	}

	var schema struct {
		TargetNamespace string `json:"TargetNamespace"`
	}
	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, fmt.Errorf(
			"failed to unmarshal schema %s: %w",
			path,
			err,
		)
	}

	// Derive namespace prefix from namespace URI and version
	prefix := deriveNamespacePrefix(
		schema.TargetNamespace,
		version,
		targetApp,
	)

	return &SchemaFileMetadata{
		Path:            path,
		Namespace:       schema.TargetNamespace,
		NamespacePrefix: prefix,
		Version:         version,
		TargetApp:       targetApp,
	}, nil
}

// extractVersionFromFilename extracts the Office version from a schema filename.
// Examples:
//   - "schemas_microsoft_com_office_excel_2010_spreadsheet.json" → Office2010
//   - "schemas_microsoft_com_office_drawing_2013_main_command.json" → Office2013
//   - "schemas_microsoft_com_office_spreadsheetml_2024_pivotAutoRefresh.json" → Office2024
func extractVersionFromFilename(
	filename string,
) FileFormatVersion {
	// Pattern to match year in filename (2006-2025)
	yearPattern := regexp.MustCompile(
		`_(20\d{2})_`,
	)
	matches := yearPattern.FindStringSubmatch(
		filename,
	)

	if len(matches) > 1 {
		if year, err := strconv.Atoi(matches[1]); err == nil {
			return yearToVersion(year)
		}
	}

	// Default to Office 2007 for schemas without year (main schemas)
	return Office2007
}

// determineTargetApp determines the target Office application from filename patterns.
func determineTargetApp(filename string) string {
	switch {
	case strings.Contains(filename, "word"):
		return "wordprocessing"
	case strings.Contains(filename, "spreadsheet") || strings.Contains(filename, "excel"):
		return "spreadsheet"
	case strings.Contains(filename, "powerpoint") || strings.Contains(filename, "presentation"):
		return "presentation"
	case strings.Contains(filename, "drawing"):
		return "drawing"
	default:
		return "unknown"
	}
}

// deriveNamespacePrefix derives the XML namespace prefix from the namespace URI and version.
// Examples:
//   - Excel 2010: "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main" → "x14"
//   - Excel 2013: "http://schemas.microsoft.com/office/spreadsheetml/2009/9/main" → "x15"
//   - Drawing 2010: "http://schemas.microsoft.com/office/drawing/2010/main" → "a14"
//   - Drawing 2013: "http://schemas.microsoft.com/office/drawing/2012/main" → "a15"
func deriveNamespacePrefix(
	namespace string,
	version FileFormatVersion,
	targetApp string,
) string {
	// Main namespaces have simple prefixes
	if namespace == "http://schemas.openxmlformats.org/wordprocessingml/2006/main" {
		return "w"
	}
	if namespace == "http://schemas.openxmlformats.org/spreadsheetml/2006/main" {
		return "x"
	}
	if namespace == "http://schemas.openxmlformats.org/presentationml/2006/main" {
		return "p"
	}
	if namespace == "http://schemas.openxmlformats.org/drawingml/2006/main" {
		return "a"
	}

	// Extension namespaces use base prefix + version suffix
	var basePrefix string
	switch {
	case strings.Contains(namespace, "/word/"):
		basePrefix = "w"
	case strings.Contains(namespace, "/spreadsheetml/"):
		basePrefix = "x"
	case strings.Contains(namespace, "/powerpoint/"):
		basePrefix = "p"
	case strings.Contains(namespace, "/drawing/"):
		basePrefix = "a"
	default:
		// For other Microsoft Office namespaces, try to infer from target app
		switch targetApp {
		case "wordprocessing":
			basePrefix = "w"
		case "spreadsheet":
			basePrefix = "x"
		case "presentation":
			basePrefix = "p"
		case "drawing":
			basePrefix = "a"
		default:
			basePrefix = "ns"
		}
	}

	// Add version suffix
	suffix := versionToSuffix(version)
	if suffix != "" {
		return basePrefix + suffix
	}

	return basePrefix
}

// yearToVersion converts a year to a FileFormatVersion.
//
//nolint:revive // Years are inherently numeric values, not magic numbers
func yearToVersion(year int) FileFormatVersion {
	switch year {
	case 2006, 2007:
		return Office2007
	case 2008, 2009, 2010:
		return Office2010
	case 2011, 2012, 2013:
		return Office2013
	case 2014, 2015, 2016:
		return Office2016
	case 2017, 2018, 2019:
		return Office2019
	case 2020, 2021:
		return Office2021
	case 2022:
		return Office2022
	case 2023:
		return Office2023
	case 2024:
		return Office2024
	case 2025:
		return Office2025
	default:
		return Office2007
	}
}

// versionToSuffix converts a FileFormatVersion to a namespace prefix suffix.
// Office 2007 has no suffix (main namespace).
// Office 2010 → "14", Office 2013 → "15", etc.
func versionToSuffix(
	version FileFormatVersion,
) string {
	switch version {
	case Office2007:
		return ""
	case Office2010:
		return "14"
	case Office2013:
		return "15"
	case Office2016:
		return "16"
	case Office2019:
		return "19"
	case Office2021:
		return "21"
	case Office2022:
		return "22"
	case Office2023:
		return "23"
	case Office2024:
		return "24"
	case Office2025:
		return "25"
	case Microsoft365:
		return "365"
	default:
		return ""
	}
}
