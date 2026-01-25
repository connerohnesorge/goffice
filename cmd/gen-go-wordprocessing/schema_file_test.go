package main

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	testSchemasDir = "../../Open-XML-SDK/data/schemas"
)

// TestFileFormatVersionString tests the String() method of FileFormatVersion.
func TestFileFormatVersionString(t *testing.T) {
	tests := []struct {
		version  FileFormatVersion
		expected string
	}{
		{Office2007, "Office2007"},
		{Office2010, "Office2010"},
		{Office2013, "Office2013"},
		{Office2016, "Office2016"},
		{Office2019, "Office2019"},
		{Office2021, "Office2021"},
		{Office2022, "Office2022"},
		{Office2023, "Office2023"},
		{Office2024, "Office2024"},
		{Office2025, "Office2025"},
		{Microsoft365, "Microsoft365"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.version.String()
			if result != tt.expected {
				t.Errorf(
					"version.String() = %q, want %q",
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestYearToVersion tests the yearToVersion function.
func TestYearToVersion(t *testing.T) {
	tests := []struct {
		year     int
		expected FileFormatVersion
	}{
		{2006, Office2007},
		{2007, Office2007},
		{2008, Office2010},
		{2009, Office2010},
		{2010, Office2010},
		{2011, Office2013},
		{2012, Office2013},
		{2013, Office2013},
		{2014, Office2016},
		{2015, Office2016},
		{2016, Office2016},
		{2017, Office2019},
		{2018, Office2019},
		{2019, Office2019},
		{2020, Office2021},
		{2021, Office2021},
		{2022, Office2022},
		{2023, Office2023},
		{2024, Office2024},
		{2025, Office2025},
	}

	for _, tt := range tests {
		t.Run(
			string(rune(tt.year)),
			func(t *testing.T) {
				result := yearToVersion(tt.year)
				if result != tt.expected {
					t.Errorf(
						"yearToVersion(%d) = %v, want %v",
						tt.year,
						result,
						tt.expected,
					)
				}
			},
		)
	}
}

// TestVersionToSuffix tests the versionToSuffix function.
func TestVersionToSuffix(t *testing.T) {
	tests := []struct {
		version  FileFormatVersion
		expected string
	}{
		{Office2007, ""},
		{Office2010, "14"},
		{Office2013, "15"},
		{Office2016, "16"},
		{Office2019, "19"},
		{Office2021, "21"},
		{Office2022, "22"},
		{Office2023, "23"},
		{Office2024, "24"},
		{Office2025, "25"},
		{Microsoft365, "365"},
	}

	for _, tt := range tests {
		t.Run(
			tt.version.String(),
			func(t *testing.T) {
				result := versionToSuffix(
					tt.version,
				)
				if result != tt.expected {
					t.Errorf(
						"versionToSuffix(%v) = %q, want %q",
						tt.version,
						result,
						tt.expected,
					)
				}
			},
		)
	}
}

// TestPrefixToVersion tests the prefixToVersion function.
func TestPrefixToVersion(t *testing.T) {
	tests := []struct {
		prefix   string
		expected FileFormatVersion
	}{
		{"w", Office2007},
		{"x", Office2007},
		{"p", Office2007},
		{"a", Office2007},
		{"w14", Office2010},
		{"x14", Office2010},
		{"a14", Office2010},
		{"w15", Office2013},
		{"x15", Office2013},
		{"a15", Office2013},
		{"w16", Office2016},
		{"x16", Office2016},
		{"a16", Office2016},
		{"w19", Office2019},
		{"w21", Office2021},
		{"w22", Office2022},
		{"w23", Office2023},
		{"w24", Office2024},
		{"w25", Office2025},
		{"w365", Microsoft365},
	}

	for _, tt := range tests {
		t.Run(tt.prefix, func(t *testing.T) {
			result := prefixToVersion(tt.prefix)
			if result != tt.expected {
				t.Errorf(
					"prefixToVersion(%q) = %v, want %v",
					tt.prefix,
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestExtractVersionFromFilename tests the extractVersionFromFilename function.
func TestExtractVersionFromFilename(
	t *testing.T,
) {
	tests := []struct {
		filename string
		expected FileFormatVersion
	}{
		{
			"schemas_microsoft_com_office_word_2010_wordml.json",
			Office2010,
		},
		{
			"schemas_microsoft_com_office_word_2010_wordprocessingCanvas.json",
			Office2010,
		},
		{
			"schemas_microsoft_com_office_drawing_2010_main.json",
			Office2010,
		},
		{
			"schemas_microsoft_com_office_drawing_2012_main.json",
			Office2013,
		},
		{
			"schemas_microsoft_com_office_drawing_2013_main_command.json",
			Office2013,
		},
		{
			"schemas_microsoft_com_office_drawing_2014_main.json",
			Office2016,
		},
		{
			"schemas_microsoft_com_office_spreadsheetml_2010_11_main.json",
			Office2010,
		},
		{
			"schemas_microsoft_com_office_spreadsheetml_2015_02_main.json",
			Office2016,
		},
		{
			"schemas_microsoft_com_office_spreadsheetml_2024_pivotAutoRefresh.json",
			Office2024,
		},
		{
			"schemas_microsoft_com_office_powerpoint_2010_main.json",
			Office2010,
		},
		{
			"schemas_microsoft_com_office_powerpoint_2022_03_main.json",
			Office2022,
		},
		{
			"schemas_microsoft_com_office_powerpoint_2023_02_main.json",
			Office2023,
		},
		// Schemas without year default to Office2007
		{"main.json", Office2007},
		{
			"schemas-openxmlformats-org_wordprocessingml_main.json",
			Office2007,
		},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := extractVersionFromFilename(
				tt.filename,
			)
			if result != tt.expected {
				t.Errorf(
					"extractVersionFromFilename(%q) = %v, want %v",
					tt.filename,
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestDetermineTargetApp tests the determineTargetApp function.
func TestDetermineTargetApp(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{
			"schemas_microsoft_com_office_word_2010_wordml.json",
			"wordprocessing",
		},
		{
			"schemas_microsoft_com_office_word_2010_wordprocessingCanvas.json",
			"wordprocessing",
		},
		{
			"schemas_microsoft_com_office_excel_2006_main.json",
			"spreadsheet",
		},
		{
			"schemas_microsoft_com_office_spreadsheetml_2010_11_main.json",
			"spreadsheet",
		},
		{
			"schemas_microsoft_com_office_powerpoint_2010_main.json",
			"presentation",
		},
		{
			"schemas_microsoft_com_office_drawing_2010_main.json",
			"drawing",
		},
		{
			"schemas_microsoft_com_office_drawing_2012_main.json",
			"drawing",
		},
		{
			"schemas_microsoft_com_office_2006_metadata_properties.json",
			"unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := determineTargetApp(
				tt.filename,
			)
			if result != tt.expected {
				t.Errorf(
					"determineTargetApp(%q) = %q, want %q",
					tt.filename,
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestDeriveNamespacePrefix tests the deriveNamespacePrefix function.
func TestDeriveNamespacePrefix(t *testing.T) {
	tests := []struct {
		namespace string
		version   FileFormatVersion
		targetApp string
		expected  string
	}{
		// Main namespaces (Office 2007)
		{
			"http://schemas.openxmlformats.org/wordprocessingml/2006/main",
			Office2007,
			"wordprocessing",
			"w",
		},
		{
			"http://schemas.openxmlformats.org/spreadsheetml/2006/main",
			Office2007,
			"spreadsheet",
			"x",
		},
		{
			"http://schemas.openxmlformats.org/presentationml/2006/main",
			Office2007,
			"presentation",
			"p",
		},
		{
			"http://schemas.openxmlformats.org/drawingml/2006/main",
			Office2007,
			"drawing",
			"a",
		},
		// Word extension namespaces
		{
			"http://schemas.microsoft.com/office/word/2010/wordml",
			Office2010,
			"wordprocessing",
			"w14",
		},
		{
			"http://schemas.microsoft.com/office/word/2012/wordml",
			Office2013,
			"wordprocessing",
			"w15",
		},
		// Excel/Spreadsheet extension namespaces
		{
			"http://schemas.microsoft.com/office/spreadsheetml/2010/11/main",
			Office2010,
			"spreadsheet",
			"x14",
		},
		{
			"http://schemas.microsoft.com/office/spreadsheetml/2015/02/main",
			Office2016,
			"spreadsheet",
			"x16",
		},
		// PowerPoint extension namespaces
		{
			"http://schemas.microsoft.com/office/powerpoint/2010/main",
			Office2010,
			"presentation",
			"p14",
		},
		{
			"http://schemas.microsoft.com/office/powerpoint/2012/main",
			Office2013,
			"presentation",
			"p15",
		},
		// Drawing extension namespaces
		{
			"http://schemas.microsoft.com/office/drawing/2010/main",
			Office2010,
			"drawing",
			"a14",
		},
		{
			"http://schemas.microsoft.com/office/drawing/2012/main",
			Office2013,
			"drawing",
			"a15",
		},
		{
			"http://schemas.microsoft.com/office/drawing/2014/main",
			Office2016,
			"drawing",
			"a16",
		},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := deriveNamespacePrefix(
				tt.namespace,
				tt.version,
				tt.targetApp,
			)
			if result != tt.expected {
				t.Errorf(
					"deriveNamespacePrefix(%q, %v, %q) = %q, want %q",
					tt.namespace,
					tt.version,
					tt.targetApp,
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestParseSchemaFile tests the parseSchemaFile function with real schema files.
func TestParseSchemaFile(t *testing.T) {
	// This test requires actual schema files to exist
	schemasDir := testSchemasDir

	// Skip if schemas directory doesn't exist
	if _, err := os.Stat(schemasDir); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Schemas directory not found, skipping integration test",
		)
	}

	tests := []struct {
		filename        string
		expectedVersion FileFormatVersion
		expectedApp     string
	}{
		{
			"schemas_microsoft_com_office_word_2010_wordml.json",
			Office2010,
			"wordprocessing",
		},
		{
			"schemas_microsoft_com_office_drawing_2010_main.json",
			Office2010,
			"drawing",
		},
		{
			"schemas_microsoft_com_office_spreadsheetml_2010_11_main.json",
			Office2010,
			"spreadsheet",
		},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			path := filepath.Join(
				schemasDir,
				tt.filename,
			)

			// Skip if file doesn't exist
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Schema file %s not found",
					tt.filename,
				)
			}

			metadata, err := parseSchemaFile(path)
			if err != nil {
				t.Fatalf(
					"parseSchemaFile(%q) error: %v",
					path,
					err,
				)
			}

			if metadata.Version != tt.expectedVersion {
				t.Errorf(
					"parseSchemaFile(%q).Version = %v, want %v",
					path,
					metadata.Version,
					tt.expectedVersion,
				)
			}

			if metadata.TargetApp != tt.expectedApp {
				t.Errorf(
					"parseSchemaFile(%q).TargetApp = %q, want %q",
					path,
					metadata.TargetApp,
					tt.expectedApp,
				)
			}

			if metadata.Namespace == "" {
				t.Errorf(
					"parseSchemaFile(%q).Namespace is empty",
					path,
				)
			}

			if metadata.NamespacePrefix == "" {
				t.Errorf(
					"parseSchemaFile(%q).NamespacePrefix is empty",
					path,
				)
			}

			if metadata.Path != path {
				t.Errorf(
					"parseSchemaFile(%q).Path = %q, want %q",
					path,
					metadata.Path,
					path,
				)
			}
		})
	}
}

// TestParseSchemaFileNamespaces tests that parseSchemaFile correctly extracts namespaces.
func TestParseSchemaFileNamespaces(t *testing.T) {
	schemasDir := testSchemasDir

	// Skip if schemas directory doesn't exist
	if _, err := os.Stat(schemasDir); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Schemas directory not found, skipping integration test",
		)
	}

	// Test a specific schema file
	path := filepath.Join(
		schemasDir,
		"schemas_microsoft_com_office_word_2010_wordml.json",
	)

	// Skip if file doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test schema file not found")
	}

	metadata, err := parseSchemaFile(path)
	if err != nil {
		t.Fatalf(
			"parseSchemaFile() error: %v",
			err,
		)
	}

	expectedNamespace := "http://schemas.microsoft.com/office/word/2010/wordml"
	if metadata.Namespace != expectedNamespace {
		t.Errorf(
			"parseSchemaFile().Namespace = %q, want %q",
			metadata.Namespace,
			expectedNamespace,
		)
	}

	expectedPrefix := "w14"
	if metadata.NamespacePrefix != expectedPrefix {
		t.Errorf(
			"parseSchemaFile().NamespacePrefix = %q, want %q",
			metadata.NamespacePrefix,
			expectedPrefix,
		)
	}
}
