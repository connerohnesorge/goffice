package font

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestSystemFontPaths(t *testing.T) {
	paths := SystemFontPaths()

	if len(paths) == 0 {
		t.Error(
			"SystemFontPaths() returned empty slice",
		)
	}

	t.Logf(
		"Found %d font paths for %s:",
		len(paths),
		runtime.GOOS,
	)
	for _, path := range paths {
		t.Logf("  %s", path)
	}
}

func TestSystemFontPathsContainsExpected(
	t *testing.T,
) {
	paths := SystemFontPaths()

	// Check for expected paths based on OS
	switch runtime.GOOS {
	case "linux":
		expectedPaths := []string{
			"/usr/share/fonts",
			"/usr/local/share/fonts",
		}
		for _, expected := range expectedPaths {
			found := false
			for _, path := range paths {
				if path == expected {
					found = true

					break
				}
			}
			if !found {
				t.Errorf(
					"Expected path %s not found in SystemFontPaths()",
					expected,
				)
			}
		}

	case "darwin":
		expectedPaths := []string{
			"/System/Library/Fonts",
			"/Library/Fonts",
		}
		for _, expected := range expectedPaths {
			found := false
			for _, path := range paths {
				if path == expected {
					found = true

					break
				}
			}
			if !found {
				t.Errorf(
					"Expected path %s not found in SystemFontPaths()",
					expected,
				)
			}
		}

	case "windows":
		// Check that Windows fonts path is included
		found := false
		for _, path := range paths {
			if filepath.Base(path) == "Fonts" {
				found = true

				break
			}
		}
		if !found {
			t.Error(
				"Expected Fonts directory not found in SystemFontPaths()",
			)
		}
	}
}

func TestFontRegistry(t *testing.T) {
	registry := NewFontRegistry()

	if registry.IsIndexed() {
		t.Error(
			"New registry should not be indexed",
		)
	}

	if registry.Count() != 0 {
		t.Error(
			"New registry should have count 0",
		)
	}
}

func TestFontRegistryAdd(t *testing.T) {
	registry := NewFontRegistry()

	info := FontInfo{
		Path:   "/path/to/font.ttf",
		Family: "Test Font",
		Style:  StyleRegular,
		Format: FormatTrueType,
	}

	registry.Add(info)

	if registry.Count() != 1 {
		t.Errorf(
			"Expected count 1, got %d",
			registry.Count(),
		)
	}

	families := registry.Families()
	if len(families) != 1 {
		t.Errorf(
			"Expected 1 family, got %d",
			len(families),
		)
	}

	if families[0] != "test font" { // normalized to lowercase
		t.Errorf(
			"Expected 'test font', got '%s'",
			families[0],
		)
	}
}

func TestFontRegistryFind(t *testing.T) {
	registry := NewFontRegistry()

	// Add multiple variants
	variants := []FontInfo{
		{
			Path:   "/path/font-regular.ttf",
			Family: "Test Font",
			Style:  StyleRegular,
			Format: FormatTrueType,
		},
		{
			Path:   "/path/font-bold.ttf",
			Family: "Test Font",
			Style:  StyleBold,
			Format: FormatTrueType,
		},
		{
			Path:   "/path/font-italic.ttf",
			Family: "Test Font",
			Style:  StyleItalic,
			Format: FormatTrueType,
		},
		{
			Path:   "/path/font-bolditalic.ttf",
			Family: "Test Font",
			Style:  StyleBoldItalic,
			Format: FormatTrueType,
		},
	}

	for _, v := range variants {
		registry.Add(v)
	}

	// Test finding each style
	tests := []struct {
		style    FontStyle
		wantPath string
	}{
		{StyleRegular, "/path/font-regular.ttf"},
		{StyleBold, "/path/font-bold.ttf"},
		{StyleItalic, "/path/font-italic.ttf"},
		{
			StyleBoldItalic,
			"/path/font-bolditalic.ttf",
		},
	}

	for _, tt := range tests {
		info, err := registry.Find(
			"Test Font",
			tt.style,
		)
		if err != nil {
			t.Errorf(
				"Find(%s) returned error: %v",
				tt.style,
				err,
			)

			continue
		}
		if info.Path != tt.wantPath {
			t.Errorf(
				"Find(%s) = %s, want %s",
				tt.style,
				info.Path,
				tt.wantPath,
			)
		}
	}
}

func TestFontRegistryFindCaseInsensitive(
	t *testing.T,
) {
	registry := NewFontRegistry()

	registry.Add(FontInfo{
		Path:   "/path/arial.ttf",
		Family: "Arial",
		Style:  StyleRegular,
		Format: FormatTrueType,
	})

	// Test case variations
	cases := []string{
		"Arial",
		"arial",
		"ARIAL",
		"ArIaL",
		"  Arial  ",
	}
	for _, name := range cases {
		info, err := registry.Find(
			name,
			StyleRegular,
		)
		if err != nil {
			t.Errorf(
				"Find(%q) returned error: %v",
				name,
				err,
			)

			continue
		}
		if info.Family != "Arial" {
			t.Errorf(
				"Find(%q) = %s, want Arial",
				name,
				info.Family,
			)
		}
	}
}

func TestFontRegistryFindNotFound(t *testing.T) {
	registry := NewFontRegistry()

	_, err := registry.Find(
		"NonexistentFont",
		StyleRegular,
	)
	if err != ErrFontNotFound {
		t.Errorf(
			"Expected ErrFontNotFound, got %v",
			err,
		)
	}
}

func TestFontRegistryFindFallbackToRegular(
	t *testing.T,
) {
	registry := NewFontRegistry()

	// Add only regular style
	registry.Add(FontInfo{
		Path:   "/path/font-regular.ttf",
		Family: "Test Font",
		Style:  StyleRegular,
		Format: FormatTrueType,
	})

	// Request bold, should fall back to regular
	info, err := registry.Find(
		"Test Font",
		StyleBold,
	)
	if err != nil {
		t.Errorf("Find() returned error: %v", err)
	}
	if info.Style != StyleRegular {
		t.Errorf(
			"Expected fallback to Regular, got %s",
			info.Style,
		)
	}
}

func TestFontRegistryVariants(t *testing.T) {
	registry := NewFontRegistry()

	variants := []FontInfo{
		{
			Path:   "/path/font-regular.ttf",
			Family: "Test Font",
			Style:  StyleRegular,
			Format: FormatTrueType,
		},
		{
			Path:   "/path/font-bold.ttf",
			Family: "Test Font",
			Style:  StyleBold,
			Format: FormatTrueType,
		},
	}

	for _, v := range variants {
		registry.Add(v)
	}

	result := registry.Variants("Test Font")
	if len(result) != 2 {
		t.Errorf(
			"Expected 2 variants, got %d",
			len(result),
		)
	}
}

func TestFontRegistryClear(t *testing.T) {
	registry := NewFontRegistry()

	registry.Add(FontInfo{
		Path:   "/path/font.ttf",
		Family: "Test Font",
		Style:  StyleRegular,
		Format: FormatTrueType,
	})

	registry.Clear()

	if registry.Count() != 0 {
		t.Error("Expected count 0 after clear")
	}
	if registry.IsIndexed() {
		t.Error(
			"Expected IsIndexed to be false after clear",
		)
	}
}

func TestFontRegistryScan(t *testing.T) {
	registry := NewFontRegistry()

	// Scan will attempt to read system fonts
	// This test just verifies it doesn't crash
	err := registry.Scan()
	if err != nil {
		// Scanning may fail on systems without fonts, that's okay
		t.Logf(
			"Scan returned error (may be expected): %v",
			err,
		)
	}

	if !registry.IsIndexed() {
		t.Error(
			"Expected IsIndexed to be true after Scan",
		)
	}

	t.Logf(
		"Found %d system fonts in %d families",
		registry.Count(),
		len(registry.Families()),
	)

	// Log some font families if found
	families := registry.Families()
	if len(families) > 0 {
		t.Logf("Sample families:")
		count := 5
		if len(families) < count {
			count = len(families)
		}
		for i := range count {
			variants := registry.Variants(
				families[i],
			)
			t.Logf(
				"  %s (%d variants)",
				families[i],
				len(variants),
			)
		}
	}
}

func TestFontRegistryScanDir(t *testing.T) {
	registry := NewFontRegistry()

	// Try scanning a real font directory
	paths := SystemFontPaths()
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			err := registry.ScanDir(path)
			if err != nil {
				t.Logf(
					"ScanDir(%s) error: %v",
					path,
					err,
				)
			} else {
				t.Logf("ScanDir(%s) found %d fonts", path, registry.Count())
			}
			// Just test one directory
			break
		}
	}
}

func TestExpandHome(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Could not get home directory")
	}

	tests := []struct {
		input    string
		expected string
	}{
		{"~", home},
		{"~/fonts", filepath.Join(home, "fonts")},
		{
			"~/.local/share/fonts",
			filepath.Join(
				home,
				".local/share/fonts",
			),
		},
		{"/usr/share/fonts", "/usr/share/fonts"},
		{"relative/path", "relative/path"},
	}

	for _, tt := range tests {
		result := expandHome(tt.input)
		if result != tt.expected {
			t.Errorf(
				"expandHome(%q) = %q, want %q",
				tt.input,
				result,
				tt.expected,
			)
		}
	}
}

func TestIsFontExtension(t *testing.T) {
	tests := []struct {
		ext  string
		want bool
	}{
		{".ttf", true},
		{".otf", true},
		{".ttc", true},
		{".otc", true},
		{
			".TTF",
			false,
		}, // Case sensitive, use lowercase
		{".txt", false},
		{".woff", false},
		{".woff2", false},
		{"", false},
	}

	for _, tt := range tests {
		got := isFontExtension(tt.ext)
		if got != tt.want {
			t.Errorf(
				"isFontExtension(%q) = %v, want %v",
				tt.ext,
				got,
				tt.want,
			)
		}
	}
}

func TestFindSystemFont(t *testing.T) {
	// This test depends on system fonts being available
	// It may not find fonts in CI environments

	// Try to find a common font
	commonFonts := []string{
		"DejaVu Sans",
		"Liberation Sans",
		"Arial",
		"Helvetica",
	}

	found := false
	for _, font := range commonFonts {
		path, err := FindSystemFont(
			font,
			StyleRegular,
		)
		if err == nil {
			t.Logf("Found %s at %s", font, path)
			found = true

			break
		}
	}

	if !found {
		t.Log(
			"No common fonts found (may be expected in CI environment)",
		)
	}
}

func TestListSystemFonts(t *testing.T) {
	fonts, err := ListSystemFonts()
	if err != nil {
		t.Errorf(
			"ListSystemFonts() error: %v",
			err,
		)
	}

	t.Logf(
		"ListSystemFonts() returned %d fonts",
		len(fonts),
	)

	// Log first few fonts
	count := 5
	if len(fonts) < count {
		count = len(fonts)
	}
	for i := range count {
		t.Logf(
			"  %s (%s) - %s",
			fonts[i].Family,
			fonts[i].Style,
			fonts[i].Path,
		)
	}
}

func TestIndexSystemFonts(t *testing.T) {
	cache := NewFontCache(10)

	err := IndexSystemFonts(cache)
	if err != nil {
		t.Errorf(
			"IndexSystemFonts() error: %v",
			err,
		)
	}

	t.Logf(
		"Cache size after indexing: %d",
		cache.Size(),
	)
}

func TestIndexSystemFontsNilCache(t *testing.T) {
	err := IndexSystemFonts(nil)
	if err == nil {
		t.Error("Expected error for nil cache")
	}
}

func TestDefaultRegistry(t *testing.T) {
	// Get default registry
	registry1 := DefaultRegistry()
	if registry1 == nil {
		t.Fatal("DefaultRegistry() returned nil")
	}

	// Get it again, should be same instance
	registry2 := DefaultRegistry()
	if registry1 != registry2 {
		t.Error(
			"DefaultRegistry() should return same instance",
		)
	}
}

func TestNormalizeFamily(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Arial", "arial"},
		{"ARIAL", "arial"},
		{"Arial Black", "arial black"},
		{"  Arial  ", "arial"},
		{"", ""},
	}

	for _, tt := range tests {
		result := normalizeFamily(tt.input)
		if result != tt.expected {
			t.Errorf(
				"normalizeFamily(%q) = %q, want %q",
				tt.input,
				result,
				tt.expected,
			)
		}
	}
}

func TestFontInfoFormat(t *testing.T) {
	info := FontInfo{
		Path:   "/path/to/font.ttf",
		Family: "Test Font",
		Style:  StyleBold,
		Format: FormatTrueType,
	}

	if info.Path != "/path/to/font.ttf" {
		t.Errorf(
			"Path = %s, want /path/to/font.ttf",
			info.Path,
		)
	}
	if info.Family != "Test Font" {
		t.Errorf(
			"Family = %s, want Test Font",
			info.Family,
		)
	}
	if info.Style != StyleBold {
		t.Errorf(
			"Style = %s, want Bold",
			info.Style,
		)
	}
	if info.Format != FormatTrueType {
		t.Errorf(
			"Format = %s, want TrueType",
			info.Format,
		)
	}
}
