package openxml

import (
	"testing"
)

// TestFileFormatVersionString tests the String() method for all versions.
func TestFileFormatVersionString(t *testing.T) {
	tests := []struct {
		version FileFormatVersion
		want    string
	}{
		{
			FileFormatVersionOffice2007,
			"Office2007",
		},
		{
			FileFormatVersionOffice2010,
			"Office2010",
		},
		{
			FileFormatVersionOffice2013,
			"Office2013",
		},
		{
			FileFormatVersionOffice2016,
			"Office2016",
		},
		{
			FileFormatVersionOffice2019,
			"Office2019",
		},
		{
			FileFormatVersionOffice2021,
			"Office2021",
		},
		{
			FileFormatVersionOffice2022,
			"Office2022",
		},
		{
			FileFormatVersionOffice2023,
			"Office2023",
		},
		{
			FileFormatVersionOffice2024,
			"Office2024",
		},
		{
			FileFormatVersionOffice2025,
			"Office2025",
		},
		{
			FileFormatVersionMicrosoft365,
			"Microsoft365",
		},
		{
			FileFormatVersion(999),
			"Unknown",
		}, // Invalid version
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := tt.version.String()
			if got != tt.want {
				t.Errorf(
					"FileFormatVersion(%d).String() = %q, want %q",
					tt.version,
					got,
					tt.want,
				)
			}
		})
	}
}

// TestFileFormatVersionDescription tests the Description() method for all versions.
func TestFileFormatVersionDescription(
	t *testing.T,
) {
	tests := []struct {
		version FileFormatVersion
		want    string
	}{
		{
			FileFormatVersionOffice2007,
			"Microsoft Office 2007 (ECMA-376 1st Edition)",
		},
		{
			FileFormatVersionOffice2010,
			"Microsoft Office 2010 (ECMA-376 2nd Edition)",
		},
		{
			FileFormatVersionOffice2013,
			"Microsoft Office 2013 (ECMA-376 3rd Edition)",
		},
		{
			FileFormatVersionOffice2016,
			"Microsoft Office 2016 (ECMA-376 5th Edition)",
		},
		{
			FileFormatVersionOffice2019,
			"Microsoft Office 2019",
		},
		{
			FileFormatVersionOffice2021,
			"Microsoft Office 2021",
		},
		{
			FileFormatVersionOffice2022,
			"Microsoft Office 2022",
		},
		{
			FileFormatVersionOffice2023,
			"Microsoft Office 2023",
		},
		{
			FileFormatVersionOffice2024,
			"Microsoft Office 2024",
		},
		{
			FileFormatVersionOffice2025,
			"Microsoft Office 2025",
		},
		{
			FileFormatVersionMicrosoft365,
			"Microsoft 365 (current)",
		},
		{
			FileFormatVersion(999),
			"Unknown Office Version",
		}, // Invalid version
	}

	for _, tt := range tests {
		t.Run(
			tt.version.String(),
			func(t *testing.T) {
				got := tt.version.Description()
				if got != tt.want {
					t.Errorf(
						"FileFormatVersion(%d).Description() = %q, want %q",
						tt.version,
						got,
						tt.want,
					)
				}
			},
		)
	}
}

// TestFileFormatVersionOrdering tests that versions can be compared.
func TestFileFormatVersionOrdering(t *testing.T) {
	// Test that versions are in chronological order
	versions := []FileFormatVersion{
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

	// Verify that each version is greater than the previous
	for i := 1; i < len(versions); i++ {
		if versions[i] <= versions[i-1] {
			t.Errorf(
				"Version ordering incorrect: %s (%d) should be > %s (%d)",
				versions[i].String(),
				versions[i],
				versions[i-1].String(),
				versions[i-1],
			)
		}
	}
}

// TestFileFormatVersionComparison tests version comparison operations.
func TestFileFormatVersionComparison(
	t *testing.T,
) {
	tests := []struct {
		name     string
		v1       FileFormatVersion
		v2       FileFormatVersion
		wantLess bool
		wantEq   bool
	}{
		{
			name:     "Office2007 < Office2010",
			v1:       FileFormatVersionOffice2007,
			v2:       FileFormatVersionOffice2010,
			wantLess: true,
			wantEq:   false,
		},
		{
			name:     "Office2010 < Office2013",
			v1:       FileFormatVersionOffice2010,
			v2:       FileFormatVersionOffice2013,
			wantLess: true,
			wantEq:   false,
		},
		{
			name:     "Office2013 < Office2016",
			v1:       FileFormatVersionOffice2013,
			v2:       FileFormatVersionOffice2016,
			wantLess: true,
			wantEq:   false,
		},
		{
			name:     "Office2016 < Office2019",
			v1:       FileFormatVersionOffice2016,
			v2:       FileFormatVersionOffice2019,
			wantLess: true,
			wantEq:   false,
		},
		{
			name:     "Office2019 < Office2021",
			v1:       FileFormatVersionOffice2019,
			v2:       FileFormatVersionOffice2021,
			wantLess: true,
			wantEq:   false,
		},
		{
			name:     "Office2024 < Office2025",
			v1:       FileFormatVersionOffice2024,
			v2:       FileFormatVersionOffice2025,
			wantLess: true,
			wantEq:   false,
		},
		{
			name:     "Office2025 < Microsoft365",
			v1:       FileFormatVersionOffice2025,
			v2:       FileFormatVersionMicrosoft365,
			wantLess: true,
			wantEq:   false,
		},
		{
			name:     "Office2013 == Office2013",
			v1:       FileFormatVersionOffice2013,
			v2:       FileFormatVersionOffice2013,
			wantLess: false,
			wantEq:   true,
		},
		{
			name:     "Office2016 > Office2010",
			v1:       FileFormatVersionOffice2016,
			v2:       FileFormatVersionOffice2010,
			wantLess: false,
			wantEq:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLess := tt.v1 < tt.v2
			if gotLess != tt.wantLess {
				t.Errorf(
					"%s < %s = %v, want %v",
					tt.v1.String(),
					tt.v2.String(),
					gotLess,
					tt.wantLess,
				)
			}

			gotEq := tt.v1 == tt.v2
			if gotEq != tt.wantEq {
				t.Errorf(
					"%s == %s = %v, want %v",
					tt.v1.String(),
					tt.v2.String(),
					gotEq,
					tt.wantEq,
				)
			}

			// Test >= for compatibility checks
			if !tt.wantLess {
				return
			}
			if tt.v2 < tt.v1 {
				t.Errorf(
					"Comparison not symmetric: %s < %s but %s < %s",
					tt.v1.String(),
					tt.v2.String(),
					tt.v2.String(),
					tt.v1.String(),
				)
			}
		})
	}
}

// TestAllFileFormatVersions tests that the AllFileFormatVersions slice is complete.
func TestAllFileFormatVersions(t *testing.T) {
	expectedVersions := []FileFormatVersion{
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

	if len(
		AllFileFormatVersions,
	) != len(
		expectedVersions,
	) {
		t.Errorf(
			"AllFileFormatVersions length = %d, want %d",
			len(AllFileFormatVersions),
			len(expectedVersions),
		)
	}

	for i, expected := range expectedVersions {
		if i >= len(AllFileFormatVersions) {
			t.Errorf(
				"AllFileFormatVersions missing version at index %d: %s",
				i,
				expected.String(),
			)

			continue
		}

		if AllFileFormatVersions[i] != expected {
			t.Errorf(
				"AllFileFormatVersions[%d] = %s, want %s",
				i,
				AllFileFormatVersions[i].String(),
				expected.String(),
			)
		}
	}

	// Verify chronological order
	for i := 1; i < len(AllFileFormatVersions); i++ {
		if AllFileFormatVersions[i] <= AllFileFormatVersions[i-1] {
			t.Errorf(
				"AllFileFormatVersions not in chronological order at index %d: %s <= %s",
				i,
				AllFileFormatVersions[i].String(),
				AllFileFormatVersions[i-1].String(),
			)
		}
	}
}

// TestFileFormatVersionUsageScenarios tests realistic usage scenarios.
func TestFileFormatVersionUsageScenarios(
	t *testing.T,
) {
	t.Run(
		"CheckOffice2010Feature",
		func(t *testing.T) {
			// Simulate checking if a feature requiring Office 2010 is supported
			requiredVersion := FileFormatVersionOffice2010
			targetVersion := FileFormatVersionOffice2013

			if targetVersion >= requiredVersion {
				// Feature is supported
				t.Logf(
					"Office 2013 supports Office 2010 features: %s >= %s",
					targetVersion.String(),
					requiredVersion.String(),
				)
			} else {
				t.Error("Office 2013 should support Office 2010 features")
			}
		},
	)

	t.Run(
		"CheckOffice2016FeatureInOffice2010",
		func(t *testing.T) {
			// Simulate checking if Office 2016 feature works in Office 2010
			requiredVersion := FileFormatVersionOffice2016
			targetVersion := FileFormatVersionOffice2010

			if targetVersion >= requiredVersion {
				t.Error(
					"Office 2010 should NOT support Office 2016 features",
				)
			} else {
				// Feature is not supported (expected)
				t.Logf("Office 2010 does not support Office 2016 features: %s < %s",
					targetVersion.String(), requiredVersion.String())
			}
		},
	)

	t.Run(
		"FindMinimumVersion",
		func(t *testing.T) {
			// Simulate finding minimum version for a document
			featuresUsed := []FileFormatVersion{
				FileFormatVersionOffice2007, // Basic Word features
				FileFormatVersionOffice2010, // Content Controls
				FileFormatVersionOffice2007, // More basic features
			}

			minVersion := FileFormatVersionOffice2007
			for _, feature := range featuresUsed {
				if feature > minVersion {
					minVersion = feature
				}
			}

			if minVersion != FileFormatVersionOffice2010 {
				t.Errorf(
					"Minimum version = %s, want Office2010",
					minVersion.String(),
				)
			}
		},
	)
}

// TestFileFormatVersionConstants tests that version constants are unique.
func TestFileFormatVersionConstants(
	t *testing.T,
) {
	versions := map[FileFormatVersion]bool{
		FileFormatVersionOffice2007:   true,
		FileFormatVersionOffice2010:   true,
		FileFormatVersionOffice2013:   true,
		FileFormatVersionOffice2016:   true,
		FileFormatVersionOffice2019:   true,
		FileFormatVersionOffice2021:   true,
		FileFormatVersionOffice2022:   true,
		FileFormatVersionOffice2023:   true,
		FileFormatVersionOffice2024:   true,
		FileFormatVersionOffice2025:   true,
		FileFormatVersionMicrosoft365: true,
	}

	// All versions should be unique (map should have 11 entries)
	if len(versions) != 11 {
		t.Errorf(
			"Expected 11 unique versions, got %d",
			len(versions),
		)
	}
}

// TestVersionStringRoundtrip verifies that version strings are meaningful.
func TestVersionStringRoundtrip(t *testing.T) {
	// Test that all version strings are non-empty and unique
	versionStrings := make(
		map[string]FileFormatVersion,
	)

	for _, version := range AllFileFormatVersions {
		str := version.String()
		if str == "" || str == "Unknown" {
			t.Errorf(
				"Version %d has invalid string representation: %q",
				version,
				str,
			)
		}

		if existing, exists := versionStrings[str]; exists {
			t.Errorf(
				"Duplicate version string %q for versions %d and %d",
				str,
				existing,
				version,
			)
		}

		versionStrings[str] = version
	}
}
