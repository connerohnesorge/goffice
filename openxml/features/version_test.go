package features

import (
	"testing"
)

// TestFileFormatVersionString tests the String method.
func TestFileFormatVersionString(t *testing.T) {
	tests := []struct {
		version  FileFormatVersion
		expected string
	}{
		{Office2007Format, "Office 2007"},
		{Office2010Format, "Office 2010"},
		{Office2013Format, "Office 2013"},
		{Office2016Format, "Office 2016"},
		{Office2019Format, "Office 2019"},
		{Office2021Format, "Office 2021"},
		{Office2022Format, "Office 2022"},
		{Office2023Format, "Office 2023"},
		{Office2024Format, "Office 2024"},
		{Office2025Format, "Office 2025"},
		{Microsoft365Format, "Microsoft 365"},
		{FileFormatVersion(999), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			got := tt.version.String()
			if got != tt.expected {
				t.Errorf(
					"FileFormatVersion.String() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

// TestFileFormatVersionDescription tests the Description method.
func TestFileFormatVersionDescription(
	t *testing.T,
) {
	tests := []struct {
		version  FileFormatVersion
		expected string
	}{
		{
			Office2007Format,
			"Microsoft Office 2007 (ECMA-376 1st Edition)",
		},
		{
			Office2010Format,
			"Microsoft Office 2010 (ISO/IEC 29500:2008)",
		},
		{
			Office2013Format,
			"Microsoft Office 2013",
		},
		{
			Office2016Format,
			"Microsoft Office 2016",
		},
		{
			Office2019Format,
			"Microsoft Office 2019",
		},
		{
			Office2021Format,
			"Microsoft Office 2021",
		},
		{
			Office2022Format,
			"Microsoft Office 2022",
		},
		{
			Office2023Format,
			"Microsoft Office 2023",
		},
		{
			Office2024Format,
			"Microsoft Office 2024",
		},
		{
			Office2025Format,
			"Microsoft Office 2025",
		},
		{
			Microsoft365Format,
			"Microsoft 365 (continuously updated)",
		},
		{
			FileFormatVersion(999),
			"Unknown Office Version",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.version.String(),
			func(t *testing.T) {
				got := tt.version.Description()
				if got != tt.expected {
					t.Errorf(
						"FileFormatVersion.Description() = %v, want %v",
						got,
						tt.expected,
					)
				}
			},
		)
	}
}

// TestFileFormatVersionYear tests the Year method.
func TestFileFormatVersionYear(t *testing.T) {
	tests := []struct {
		version  FileFormatVersion
		expected int
	}{
		{Office2007Format, 2007},
		{Office2010Format, 2010},
		{Office2013Format, 2013},
		{Office2016Format, 2016},
		{Office2019Format, 2019},
		{Office2021Format, 2021},
		{Office2022Format, 2022},
		{Office2023Format, 2023},
		{Office2024Format, 2024},
		{Office2025Format, 2025},
		{
			Microsoft365Format,
			0,
		}, // Continuously updated
		{FileFormatVersion(999), 0},
	}

	for _, tt := range tests {
		t.Run(
			tt.version.String(),
			func(t *testing.T) {
				got := tt.version.Year()
				if got != tt.expected {
					t.Errorf(
						"FileFormatVersion.Year() = %v, want %v",
						got,
						tt.expected,
					)
				}
			},
		)
	}
}

// TestFileFormatVersionAtLeast tests the AtLeast method.
func TestFileFormatVersionAtLeast(t *testing.T) {
	tests := []struct {
		name     string
		version  FileFormatVersion
		other    FileFormatVersion
		expected bool
	}{
		{
			"Office2013 >= Office2010",
			Office2013Format,
			Office2010Format,
			true,
		},
		{
			"Office2013 >= Office2013",
			Office2013Format,
			Office2013Format,
			true,
		},
		{
			"Office2010 >= Office2013",
			Office2010Format,
			Office2013Format,
			false,
		},
		{
			"Office2007 >= Office2025",
			Office2007Format,
			Office2025Format,
			false,
		},
		{
			"Microsoft365 >= Office2007",
			Microsoft365Format,
			Office2007Format,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.version.AtLeast(tt.other)
			if got != tt.expected {
				t.Errorf(
					"FileFormatVersion.AtLeast() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

// TestFileFormatVersionAtMost tests the AtMost method.
func TestFileFormatVersionAtMost(t *testing.T) {
	tests := []struct {
		name     string
		version  FileFormatVersion
		other    FileFormatVersion
		expected bool
	}{
		{
			"Office2010 <= Office2013",
			Office2010Format,
			Office2013Format,
			true,
		},
		{
			"Office2013 <= Office2013",
			Office2013Format,
			Office2013Format,
			true,
		},
		{
			"Office2013 <= Office2010",
			Office2013Format,
			Office2010Format,
			false,
		},
		{
			"Office2007 <= Microsoft365",
			Office2007Format,
			Microsoft365Format,
			true,
		},
		{
			"Microsoft365 <= Office2007",
			Microsoft365Format,
			Office2007Format,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.version.AtMost(tt.other)
			if got != tt.expected {
				t.Errorf(
					"FileFormatVersion.AtMost() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

// TestFileFormatVersionIsNewerThan tests the IsNewerThan method.
func TestFileFormatVersionIsNewerThan(
	t *testing.T,
) {
	tests := []struct {
		name     string
		version  FileFormatVersion
		other    FileFormatVersion
		expected bool
	}{
		{
			"Office2013 > Office2010",
			Office2013Format,
			Office2010Format,
			true,
		},
		{
			"Office2013 > Office2013",
			Office2013Format,
			Office2013Format,
			false,
		},
		{
			"Office2010 > Office2013",
			Office2010Format,
			Office2013Format,
			false,
		},
		{
			"Microsoft365 > Office2025",
			Microsoft365Format,
			Office2025Format,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.version.IsNewerThan(
				tt.other,
			)
			if got != tt.expected {
				t.Errorf(
					"FileFormatVersion.IsNewerThan() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

// TestFileFormatVersionIsOlderThan tests the IsOlderThan method.
func TestFileFormatVersionIsOlderThan(
	t *testing.T,
) {
	tests := []struct {
		name     string
		version  FileFormatVersion
		other    FileFormatVersion
		expected bool
	}{
		{
			"Office2010 < Office2013",
			Office2010Format,
			Office2013Format,
			true,
		},
		{
			"Office2013 < Office2013",
			Office2013Format,
			Office2013Format,
			false,
		},
		{
			"Office2013 < Office2010",
			Office2013Format,
			Office2010Format,
			false,
		},
		{
			"Office2007 < Microsoft365",
			Office2007Format,
			Microsoft365Format,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.version.IsOlderThan(
				tt.other,
			)
			if got != tt.expected {
				t.Errorf(
					"FileFormatVersion.IsOlderThan() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

// TestAllFileFormatVersions tests that all versions are in chronological order.
func TestAllFileFormatVersions(t *testing.T) {
	if len(AllFileFormatVersions) != 11 {
		t.Errorf(
			"AllFileFormatVersions length = %d, want 11",
			len(AllFileFormatVersions),
		)
	}

	// Verify chronological ordering
	for i := 1; i < len(AllFileFormatVersions); i++ {
		prev := AllFileFormatVersions[i-1]
		curr := AllFileFormatVersions[i]
		if prev >= curr {
			t.Errorf(
				"AllFileFormatVersions not in chronological order: %v >= %v",
				prev.String(),
				curr.String(),
			)
		}
	}

	// Verify first and last versions
	if AllFileFormatVersions[0] != Office2007Format {
		t.Errorf(
			"First version = %v, want Office2007Format",
			AllFileFormatVersions[0].String(),
		)
	}
	if AllFileFormatVersions[len(AllFileFormatVersions)-1] !=
		Microsoft365Format {
		t.Errorf(
			"Last version = %v, want Microsoft365Format",
			AllFileFormatVersions[len(AllFileFormatVersions)-1].String(),
		)
	}
}

// TestFileFormatVersionOrdering tests that version constants are correctly ordered.
func TestFileFormatVersionOrdering(t *testing.T) {
	// Verify that versions are in ascending order
	if Office2007Format >= Office2010Format {
		t.Error(
			"Office2007Format should be < Office2010Format",
		)
	}
	if Office2010Format >= Office2013Format {
		t.Error(
			"Office2010Format should be < Office2013Format",
		)
	}
	if Office2013Format >= Office2016Format {
		t.Error(
			"Office2013Format should be < Office2016Format",
		)
	}
	if Office2016Format >= Office2019Format {
		t.Error(
			"Office2016Format should be < Office2019Format",
		)
	}
	if Office2019Format >= Office2021Format {
		t.Error(
			"Office2019Format should be < Office2021Format",
		)
	}
	if Office2021Format >= Office2022Format {
		t.Error(
			"Office2021Format should be < Office2022Format",
		)
	}
	if Office2022Format >= Office2023Format {
		t.Error(
			"Office2022Format should be < Office2023Format",
		)
	}
	if Office2023Format >= Office2024Format {
		t.Error(
			"Office2023Format should be < Office2024Format",
		)
	}
	if Office2024Format >= Office2025Format {
		t.Error(
			"Office2024Format should be < Office2025Format",
		)
	}
	if Office2025Format >= Microsoft365Format {
		t.Error(
			"Office2025Format should be < Microsoft365Format",
		)
	}
}
