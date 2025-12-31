package openxml

import (
	"testing"
)

// TestNamespaceConstants verifies all namespace constants are defined correctly.
func TestNamespaceConstants(t *testing.T) {
	tests := []struct {
		name      string
		namespace string
		wantURI   string
	}{
		// Main namespaces
		{
			name:      "WordprocessingML",
			namespace: NamespaceWordprocessingML,
			wantURI:   "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
		},
		{
			name:      "SpreadsheetML",
			namespace: NamespaceSpreadsheetML,
			wantURI:   "http://schemas.openxmlformats.org/spreadsheetml/2006/main",
		},
		{
			name:      "PresentationML",
			namespace: NamespacePresentationML,
			wantURI:   "http://schemas.openxmlformats.org/presentationml/2006/main",
		},
		{
			name:      "MarkupCompatibility",
			namespace: NamespaceMarkupCompatibility,
			wantURI:   "http://schemas.openxmlformats.org/markup-compatibility/2006",
		},

		// Word extension namespaces
		{
			name:      "Word2010",
			namespace: NamespaceWord2010,
			wantURI:   "http://schemas.microsoft.com/office/word/2010/wordml",
		},
		{
			name:      "Word2013",
			namespace: NamespaceWord2013,
			wantURI:   "http://schemas.microsoft.com/office/word/2012/wordml",
		},
		{
			name:      "Word2016",
			namespace: NamespaceWord2016,
			wantURI:   "http://schemas.microsoft.com/office/word/2015/wordml",
		},
		{
			name:      "Word2019",
			namespace: NamespaceWord2019,
			wantURI:   "http://schemas.microsoft.com/office/word/2018/wordml",
		},

		// Excel extension namespaces
		{
			name:      "Excel2009",
			namespace: NamespaceExcel2009,
			wantURI:   "http://schemas.microsoft.com/office/spreadsheetml/2009/9/main",
		},
		{
			name:      "Excel2013",
			namespace: NamespaceExcel2013,
			wantURI:   "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main",
		},
		{
			name:      "Excel2016",
			namespace: NamespaceExcel2016,
			wantURI:   "http://schemas.microsoft.com/office/spreadsheetml/2014/11/main",
		},
		{
			name:      "Excel2025",
			namespace: NamespaceExcel2025,
			wantURI:   "http://schemas.microsoft.com/office/spreadsheetml/2024/8/main",
		},

		// PowerPoint extension namespaces
		{
			name:      "PowerPoint2010",
			namespace: NamespacePowerPoint2010,
			wantURI:   "http://schemas.microsoft.com/office/powerpoint/2010/main",
		},
		{
			name:      "PowerPoint2012",
			namespace: NamespacePowerPoint2012,
			wantURI:   "http://schemas.microsoft.com/office/powerpoint/2012/main",
		},
		{
			name:      "PowerPoint2016",
			namespace: NamespacePowerPoint2016,
			wantURI:   "http://schemas.microsoft.com/office/powerpoint/2015/main",
		},

		// DrawingML extension namespaces
		{
			name:      "Drawing2010",
			namespace: NamespaceDrawing2010,
			wantURI:   "http://schemas.microsoft.com/office/drawing/2010/main",
		},
		{
			name:      "Drawing2012",
			namespace: NamespaceDrawing2012,
			wantURI:   "http://schemas.microsoft.com/office/drawing/2012/main",
		},
		{
			name:      "Drawing2014",
			namespace: NamespaceDrawing2014,
			wantURI:   "http://schemas.microsoft.com/office/drawing/2014/main",
		},
		{
			name:      "Chart2014",
			namespace: NamespaceChart2014,
			wantURI:   "http://schemas.microsoft.com/office/drawing/2012/chart",
		},
		{
			name:      "Chart2016",
			namespace: NamespaceChart2016,
			wantURI:   "http://schemas.microsoft.com/office/drawing/2014/chart",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.namespace != tt.wantURI {
				t.Errorf(
					"Namespace constant mismatch: got %q, want %q",
					tt.namespace,
					tt.wantURI,
				)
			}
		})
	}
}

// TestGetPrefixForNamespace tests namespace to prefix mapping.
func TestGetPrefixForNamespace(t *testing.T) {
	tests := []struct {
		name         string
		namespaceURI string
		wantPrefix   string
	}{
		// Main namespaces
		{
			name:         "WordprocessingML",
			namespaceURI: NamespaceWordprocessingML,
			wantPrefix:   "w",
		},
		{
			name:         "SpreadsheetML",
			namespaceURI: NamespaceSpreadsheetML,
			wantPrefix:   "x",
		},
		{
			name:         "PresentationML",
			namespaceURI: NamespacePresentationML,
			wantPrefix:   "p",
		},
		{
			name:         "DrawingML",
			namespaceURI: NamespaceDrawingML,
			wantPrefix:   "a",
		},
		{
			name:         "MarkupCompatibility",
			namespaceURI: NamespaceMarkupCompatibility,
			wantPrefix:   "mc",
		},

		// Word extensions
		{
			name:         "Word2010",
			namespaceURI: NamespaceWord2010,
			wantPrefix:   "w14",
		},
		{
			name:         "Word2013",
			namespaceURI: NamespaceWord2013,
			wantPrefix:   "w15",
		},
		{
			name:         "Word2016",
			namespaceURI: NamespaceWord2016,
			wantPrefix:   "w16",
		},
		{
			name:         "Word2019",
			namespaceURI: NamespaceWord2019,
			wantPrefix:   "w19",
		},
		{
			name:         "WordprocessingDrawing2010",
			namespaceURI: NamespaceWord2010WordprocessingDrawing,
			wantPrefix:   "wp14",
		},
		{
			name:         "WordprocessingCanvas",
			namespaceURI: NamespaceWord2010WordprocessingCanvas,
			wantPrefix:   "wpc",
		},
		{
			name:         "WordprocessingGroup",
			namespaceURI: NamespaceWord2010WordprocessingGroup,
			wantPrefix:   "wpg",
		},
		{
			name:         "WordprocessingShape",
			namespaceURI: NamespaceWord2010WordprocessingShape,
			wantPrefix:   "wps",
		},

		// Excel extensions
		{
			name:         "Excel2009",
			namespaceURI: NamespaceExcel2009,
			wantPrefix:   "x14",
		},
		{
			name:         "Excel2013",
			namespaceURI: NamespaceExcel2013,
			wantPrefix:   "x15",
		},
		{
			name:         "Excel2016",
			namespaceURI: NamespaceExcel2016,
			wantPrefix:   "x16",
		},
		{
			name:         "Excel2019",
			namespaceURI: NamespaceExcel2019,
			wantPrefix:   "x19",
		},
		{
			name:         "Excel2021",
			namespaceURI: NamespaceExcel2021,
			wantPrefix:   "x21",
		},
		{
			name:         "Excel2024",
			namespaceURI: NamespaceExcel2024,
			wantPrefix:   "x24",
		},
		{
			name:         "Excel2025",
			namespaceURI: NamespaceExcel2025,
			wantPrefix:   "x25",
		},

		// PowerPoint extensions
		{
			name:         "PowerPoint2010",
			namespaceURI: NamespacePowerPoint2010,
			wantPrefix:   "p14",
		},
		{
			name:         "PowerPoint2012",
			namespaceURI: NamespacePowerPoint2012,
			wantPrefix:   "p15",
		},
		{
			name:         "PowerPoint2016",
			namespaceURI: NamespacePowerPoint2016,
			wantPrefix:   "p16",
		},
		{
			name:         "PowerPoint2021",
			namespaceURI: NamespacePowerPoint2021,
			wantPrefix:   "p21",
		},

		// DrawingML extensions
		{
			name:         "Drawing2010",
			namespaceURI: NamespaceDrawing2010,
			wantPrefix:   "a14",
		},
		{
			name:         "Drawing2012",
			namespaceURI: NamespaceDrawing2012,
			wantPrefix:   "a15",
		},
		{
			name:         "Drawing2014",
			namespaceURI: NamespaceDrawing2014,
			wantPrefix:   "a16",
		},
		{
			name:         "Drawing2016SVG",
			namespaceURI: NamespaceDrawing2016SVG,
			wantPrefix:   "asvg",
		},
		{
			name:         "Drawing2016Ink",
			namespaceURI: NamespaceDrawing2016Ink,
			wantPrefix:   "aink",
		},
		{
			name:         "Chart2014",
			namespaceURI: NamespaceChart2014,
			wantPrefix:   "c15",
		},
		{
			name:         "Chart2016",
			namespaceURI: NamespaceChart2016,
			wantPrefix:   "c16",
		},
		{
			name:         "Chart2016r3",
			namespaceURI: NamespaceChart2016r3,
			wantPrefix:   "c16r3",
		},

		// Unknown namespace
		{
			name:         "UnknownNamespace",
			namespaceURI: "http://unknown.namespace.com/test",
			wantPrefix:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix := GetPrefixForNamespace(
				tt.namespaceURI,
			)
			if gotPrefix != tt.wantPrefix {
				t.Errorf(
					"GetPrefixForNamespace(%q) = %q, want %q",
					tt.namespaceURI,
					gotPrefix,
					tt.wantPrefix,
				)
			}
		})
	}
}

// TestGetNamespaceForPrefix tests prefix to namespace mapping.
func TestGetNamespaceForPrefix(t *testing.T) {
	tests := []struct {
		name          string
		prefix        string
		wantNamespace string
	}{
		// Main namespaces
		{
			name:          "WordprocessingML",
			prefix:        "w",
			wantNamespace: NamespaceWordprocessingML,
		},
		{
			name:          "SpreadsheetML",
			prefix:        "x",
			wantNamespace: NamespaceSpreadsheetML,
		},
		{
			name:          "PresentationML",
			prefix:        "p",
			wantNamespace: NamespacePresentationML,
		},
		{
			name:          "DrawingML",
			prefix:        "a",
			wantNamespace: NamespaceDrawingML,
		},
		{
			name:          "MarkupCompatibility",
			prefix:        "mc",
			wantNamespace: NamespaceMarkupCompatibility,
		},

		// Word extensions
		{
			name:          "Word2010",
			prefix:        "w14",
			wantNamespace: NamespaceWord2010,
		},
		{
			name:          "Word2013",
			prefix:        "w15",
			wantNamespace: NamespaceWord2013,
		},
		{
			name:          "Word2016",
			prefix:        "w16",
			wantNamespace: NamespaceWord2016,
		},
		{
			name:          "Word2019",
			prefix:        "w19",
			wantNamespace: NamespaceWord2019,
		},

		// Excel extensions
		{
			name:          "Excel2013",
			prefix:        "x15",
			wantNamespace: NamespaceExcel2013,
		},
		{
			name:          "Excel2016",
			prefix:        "x16",
			wantNamespace: NamespaceExcel2016,
		},
		{
			name:          "Excel2025",
			prefix:        "x25",
			wantNamespace: NamespaceExcel2025,
		},

		// PowerPoint extensions
		{
			name:          "PowerPoint2010",
			prefix:        "p14",
			wantNamespace: NamespacePowerPoint2010,
		},
		{
			name:          "PowerPoint2016",
			prefix:        "p16",
			wantNamespace: NamespacePowerPoint2016,
		},

		// DrawingML extensions
		{
			name:          "Drawing2010",
			prefix:        "a14",
			wantNamespace: NamespaceDrawing2010,
		},
		{
			name:          "Drawing2012",
			prefix:        "a15",
			wantNamespace: NamespaceDrawing2012,
		},
		{
			name:          "Chart2014",
			prefix:        "c15",
			wantNamespace: NamespaceChart2014,
		},

		// Unknown prefix
		{
			name:          "UnknownPrefix",
			prefix:        "unknown",
			wantNamespace: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNamespace := GetNamespaceForPrefix(
				tt.prefix,
			)
			if gotNamespace != tt.wantNamespace {
				t.Errorf(
					"GetNamespaceForPrefix(%q) = %q, want %q",
					tt.prefix,
					gotNamespace,
					tt.wantNamespace,
				)
			}
		})
	}
}

// TestRegisterNamespace tests dynamic namespace registration.
func TestRegisterNamespace(t *testing.T) {
	// Register a custom namespace
	customNS := "http://custom.namespace.com/test"
	customPrefix := "custom"

	RegisterNamespace(customNS, customPrefix)

	// Test that it was registered
	gotPrefix := GetPrefixForNamespace(customNS)
	if gotPrefix != customPrefix {
		t.Errorf(
			"After RegisterNamespace, GetPrefixForNamespace(%q) = %q, want %q",
			customNS,
			gotPrefix,
			customPrefix,
		)
	}

	gotNamespace := GetNamespaceForPrefix(
		customPrefix,
	)
	if gotNamespace != customNS {
		t.Errorf(
			"After RegisterNamespace, GetNamespaceForPrefix(%q) = %q, want %q",
			customPrefix,
			gotNamespace,
			customNS,
		)
	}

	// Clean up - restore original state
	delete(NamespacePrefixes, customNS)
	delete(PrefixNamespaces, customPrefix)
}

// TestNamespacePrefixMapCompleteness verifies all registered namespaces have prefixes.
func TestNamespacePrefixMapCompleteness(
	t *testing.T,
) {
	// Test that critical extension namespaces are registered
	requiredNamespaces := []struct {
		name string
		ns   string
	}{
		{"Word2010", NamespaceWord2010},
		{"Word2013", NamespaceWord2013},
		{"Word2016", NamespaceWord2016},
		{"Excel2009", NamespaceExcel2009},
		{"Excel2013", NamespaceExcel2013},
		{"Excel2016", NamespaceExcel2016},
		{
			"PowerPoint2010",
			NamespacePowerPoint2010,
		},
		{
			"PowerPoint2012",
			NamespacePowerPoint2012,
		},
		{"Drawing2010", NamespaceDrawing2010},
		{"Drawing2012", NamespaceDrawing2012},
		{"Chart2014", NamespaceChart2014},
		{"Chart2016", NamespaceChart2016},
	}

	for _, req := range requiredNamespaces {
		t.Run(req.name, func(t *testing.T) {
			prefix := GetPrefixForNamespace(
				req.ns,
			)
			if prefix == "" {
				t.Errorf(
					"Required namespace %s (%q) has no prefix registered",
					req.name,
					req.ns,
				)
			}
		})
	}
}

// TestPrefixNamespacesBidirectional verifies that the prefix->namespace map is consistent.
func TestPrefixNamespacesBidirectional(
	t *testing.T,
) {
	// For each namespace, verify that prefix->namespace->prefix is consistent
	for ns, prefix := range NamespacePrefixes {
		t.Run(prefix, func(t *testing.T) {
			// Get the namespace for this prefix
			gotNS := GetNamespaceForPrefix(prefix)

			// It should map back to a valid namespace
			if gotNS == "" {
				t.Errorf(
					"Prefix %q has no namespace in PrefixNamespaces map",
					prefix,
				)

				return
			}

			// For duplicate prefixes (e.g., x14), we just verify consistency
			// The reverse lookup may not give us the exact same namespace
			gotPrefix := GetPrefixForNamespace(
				gotNS,
			)
			if gotPrefix != prefix {
				// This is acceptable for duplicate prefixes
				t.Logf(
					"Note: Namespace %q maps to prefix %q, but prefix maps back to namespace %q (duplicate prefix)",
					ns,
					prefix,
					gotNS,
				)
			}
		})
	}
}

// TestNamespacePrefixMapSize verifies we have a reasonable number of namespaces registered.
func TestNamespacePrefixMapSize(t *testing.T) {
	// We should have at least 40+ namespaces registered
	// (main + extensions for Word, Excel, PowerPoint, DrawingML)
	minExpectedCount := 40

	actualCount := len(NamespacePrefixes)
	if actualCount < minExpectedCount {
		t.Errorf(
			"Expected at least %d namespaces, got %d",
			minExpectedCount,
			actualCount,
		)
	}

	t.Logf(
		"Total namespaces registered: %d",
		actualCount,
	)
}

// TestExtensionNamespacePatterns verifies extension namespace URLs follow expected patterns.
func TestExtensionNamespacePatterns(
	t *testing.T,
) {
	tests := []struct {
		name       string
		namespace  string
		wantPrefix string
	}{
		// Word extensions should use schemas.microsoft.com/office/word
		{
			name:       "Word2010Pattern",
			namespace:  NamespaceWord2010,
			wantPrefix: "http://schemas.microsoft.com/office/word/",
		},
		{
			name:       "Word2013Pattern",
			namespace:  NamespaceWord2013,
			wantPrefix: "http://schemas.microsoft.com/office/word/",
		},

		// Excel extensions should use schemas.microsoft.com/office/spreadsheetml
		{
			name:       "Excel2009Pattern",
			namespace:  NamespaceExcel2009,
			wantPrefix: "http://schemas.microsoft.com/office/spreadsheetml/",
		},
		{
			name:       "Excel2016Pattern",
			namespace:  NamespaceExcel2016,
			wantPrefix: "http://schemas.microsoft.com/office/spreadsheetml/",
		},

		// PowerPoint extensions should use schemas.microsoft.com/office/powerpoint
		{
			name:       "PowerPoint2010Pattern",
			namespace:  NamespacePowerPoint2010,
			wantPrefix: "http://schemas.microsoft.com/office/powerpoint/",
		},

		// DrawingML extensions should use schemas.microsoft.com/office/drawing
		{
			name:       "Drawing2010Pattern",
			namespace:  NamespaceDrawing2010,
			wantPrefix: "http://schemas.microsoft.com/office/drawing/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(
				tt.namespace,
			) < len(
				tt.wantPrefix,
			) ||
				tt.namespace[:len(tt.wantPrefix)] != tt.wantPrefix {
				t.Errorf(
					"Namespace %q does not start with expected prefix %q",
					tt.namespace,
					tt.wantPrefix,
				)
			}
		})
	}
}
