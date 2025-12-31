package openxml

import (
	"strings"
	"testing"
)

func TestAlternateContent_Creation(t *testing.T) {
	ac := NewAlternateContent()

	if ac == nil {
		t.Fatal(
			"NewAlternateContent() returned nil",
		)
	}

	if ac.LocalName() != "AlternateContent" {
		t.Errorf(
			"LocalName() = %q, want %q",
			ac.LocalName(),
			"AlternateContent",
		)
	}

	if ac.NamespaceURI() != NamespaceMarkupCompatibility {
		t.Errorf(
			"NamespaceURI() = %q, want %q",
			ac.NamespaceURI(),
			NamespaceMarkupCompatibility,
		)
	}
}

func TestChoice_Creation(t *testing.T) {
	c := NewChoice()

	if c == nil {
		t.Fatal("NewChoice() returned nil")
	}

	if c.LocalName() != "Choice" {
		t.Errorf(
			"LocalName() = %q, want %q",
			c.LocalName(),
			"Choice",
		)
	}

	if c.NamespaceURI() != NamespaceMarkupCompatibility {
		t.Errorf(
			"NamespaceURI() = %q, want %q",
			c.NamespaceURI(),
			NamespaceMarkupCompatibility,
		)
	}
}

func TestFallback_Creation(t *testing.T) {
	fb := NewFallback()

	if fb == nil {
		t.Fatal("NewFallback() returned nil")
	}

	if fb.LocalName() != "Fallback" {
		t.Errorf(
			"LocalName() = %q, want %q",
			fb.LocalName(),
			"Fallback",
		)
	}

	if fb.NamespaceURI() != NamespaceMarkupCompatibility {
		t.Errorf(
			"NamespaceURI() = %q, want %q",
			fb.NamespaceURI(),
			NamespaceMarkupCompatibility,
		)
	}
}

func TestChoice_Requires(t *testing.T) {
	c := NewChoice()

	// Initially should be empty
	if requires := c.Requires(); requires != "" {
		t.Errorf(
			"Requires() = %q, want empty string",
			requires,
		)
	}

	// Set and get
	c.SetRequires("w14")
	if requires := c.Requires(); requires != "w14" {
		t.Errorf(
			"After SetRequires('w14'), Requires() = %q, want %q",
			requires,
			"w14",
		)
	}
}

func TestAlternateContent_ChoicesAndFallback(
	t *testing.T,
) {
	ac := NewAlternateContent()

	// Initially no children
	if choices := ac.Choices(); len(
		choices,
	) != 0 {
		t.Errorf(
			"Choices() = %d elements, want 0",
			len(choices),
		)
	}
	if fb := ac.Fallback(); fb != nil {
		t.Error(
			"Fallback() returned non-nil, want nil",
		)
	}

	// Add a Choice
	c1 := NewChoice()
	c1.SetRequires("w14")
	ac.AppendChild(c1)

	choices := ac.Choices()
	if len(choices) != 1 {
		t.Errorf(
			"After adding Choice, Choices() = %d elements, want 1",
			len(choices),
		)
	}
	if choices[0].Requires() != "w14" {
		t.Errorf(
			"Choice.Requires() = %q, want %q",
			choices[0].Requires(),
			"w14",
		)
	}

	// Add another Choice
	c2 := NewChoice()
	c2.SetRequires("w15")
	ac.AppendChild(c2)

	choices = ac.Choices()
	if len(choices) != 2 {
		t.Errorf(
			"After adding second Choice, Choices() = %d elements, want 2",
			len(choices),
		)
	}

	// Add a Fallback
	fb := NewFallback()
	ac.AppendChild(fb)

	if retrievedFb := ac.Fallback(); retrievedFb == nil {
		t.Error(
			"After adding Fallback, Fallback() returned nil",
		)
	}
}

func TestPrefixToVersion(t *testing.T) {
	tests := []struct {
		prefix  string
		want    FileFormatVersion
		wantStr string
	}{
		{
			"w14",
			FileFormatVersionOffice2010,
			"Office2010",
		},
		{
			"x14",
			FileFormatVersionOffice2010,
			"Office2010",
		},
		{
			"a14",
			FileFormatVersionOffice2010,
			"Office2010",
		},
		{
			"w15",
			FileFormatVersionOffice2013,
			"Office2013",
		},
		{
			"x15",
			FileFormatVersionOffice2013,
			"Office2013",
		},
		{
			"w16",
			FileFormatVersionOffice2016,
			"Office2016",
		},
		{
			"w19",
			FileFormatVersionOffice2019,
			"Office2019",
		},
		{
			"w21",
			FileFormatVersionOffice2021,
			"Office2021",
		},
		{
			"w22",
			FileFormatVersionOffice2022,
			"Office2022",
		},
		{
			"w23",
			FileFormatVersionOffice2023,
			"Office2023",
		},
		{
			"w24",
			FileFormatVersionOffice2024,
			"Office2024",
		},
		{
			"w25",
			FileFormatVersionOffice2025,
			"Office2025",
		},
		{
			"unknown",
			FileFormatVersionOffice2007,
			"Office2007",
		},
		{
			"",
			FileFormatVersionOffice2007,
			"Office2007",
		},
		{
			"w",
			FileFormatVersionOffice2007,
			"Office2007",
		},
	}

	for _, tt := range tests {
		t.Run(tt.prefix, func(t *testing.T) {
			got := prefixToVersion(tt.prefix)
			if got != tt.want {
				t.Errorf(
					"prefixToVersion(%q) = %v (%s), want %v (%s)",
					tt.prefix,
					got,
					got.String(),
					tt.want,
					tt.wantStr,
				)
			}
		})
	}
}

func TestChoice_IsSatisfiedBy(t *testing.T) {
	tests := []struct {
		name          string
		requires      string
		targetVersion FileFormatVersion
		want          bool
	}{
		{
			name:          "Office2010 requirement satisfied by Office2010",
			requires:      "w14",
			targetVersion: FileFormatVersionOffice2010,
			want:          true,
		},
		{
			name:          "Office2010 requirement satisfied by Office2013",
			requires:      "w14",
			targetVersion: FileFormatVersionOffice2013,
			want:          true,
		},
		{
			name:          "Office2013 requirement NOT satisfied by Office2010",
			requires:      "w15",
			targetVersion: FileFormatVersionOffice2010,
			want:          false,
		},
		{
			name:          "Office2013 requirement satisfied by Office2013",
			requires:      "w15",
			targetVersion: FileFormatVersionOffice2013,
			want:          true,
		},
		{
			name:          "Empty requires not satisfied",
			requires:      "",
			targetVersion: FileFormatVersionOffice2013,
			want:          false,
		},
		{
			name:          "Unknown prefix not satisfied",
			requires:      "unknown",
			targetVersion: FileFormatVersionOffice2013,
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewChoice()
			c.SetRequires(tt.requires)

			got := c.IsSatisfiedBy(
				tt.targetVersion,
			)
			if got != tt.want {
				t.Errorf(
					"IsSatisfiedBy(%s) = %v, want %v",
					tt.targetVersion.String(),
					got,
					tt.want,
				)
			}
		})
	}
}

func TestAlternateContent_SelectContent(
	t *testing.T,
) {
	tests := []struct {
		name          string
		setupChoices  []string // Requires values for each Choice
		targetVersion FileFormatVersion
		wantChoice    bool // true if should select a Choice, false if should select Fallback
		wantIndex     int  // which Choice index (if wantChoice is true)
	}{
		{
			name:          "Select Office2010 Choice for Office2010",
			setupChoices:  []string{"w14"},
			targetVersion: FileFormatVersionOffice2010,
			wantChoice:    true,
			wantIndex:     0,
		},
		{
			name:          "Select Office2010 Choice for Office2013",
			setupChoices:  []string{"w14"},
			targetVersion: FileFormatVersionOffice2013,
			wantChoice:    true,
			wantIndex:     0,
		},
		{
			name:          "Select Fallback when Office2013 Choice not satisfied",
			setupChoices:  []string{"w15"},
			targetVersion: FileFormatVersionOffice2010,
			wantChoice:    false,
		},
		{
			name:          "Select first satisfied Choice",
			setupChoices:  []string{"w14", "w15"},
			targetVersion: FileFormatVersionOffice2013,
			wantChoice:    true,
			wantIndex:     0,
		},
		{
			name:          "Select second Choice when first not satisfied",
			setupChoices:  []string{"w15", "w14"},
			targetVersion: FileFormatVersionOffice2010,
			wantChoice:    true, // w14 (Office2010) should satisfy Office2010 target
			wantIndex:     1,    // Second choice (index 1)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := NewAlternateContent()

			// Add Choices
			for _, requires := range tt.setupChoices {
				c := NewChoice()
				c.SetRequires(requires)
				// Add some child to identify it
				child := NewCompositeElement(
					"test",
					"child",
					"",
				)
				c.AppendChild(child)
				ac.AppendChild(c)
			}

			// Add Fallback
			fb := NewFallback()
			fbChild := NewCompositeElement(
				"test",
				"fallback-child",
				"",
			)
			fb.AppendChild(fbChild)
			ac.AppendChild(fb)

			// Select content
			selected := ac.SelectContent(
				tt.targetVersion,
			)

			if tt.wantChoice {
				choice, ok := selected.(*Choice)
				if !ok {
					t.Fatalf(
						"SelectContent() returned %T, want *Choice",
						selected,
					)
				}
				// Verify it's the right Choice by checking the Requires value
				expectedRequires := tt.setupChoices[tt.wantIndex]
				if choice.Requires() != expectedRequires {
					t.Errorf(
						"Selected Choice has Requires=%q, want %q",
						choice.Requires(),
						expectedRequires,
					)
				}
			} else {
				fallback, ok := selected.(*Fallback)
				if !ok {
					t.Fatalf("SelectContent() returned %T, want *Fallback", selected)
				}
				if fallback == nil {
					t.Error("SelectContent() returned nil Fallback")
				}
			}
		})
	}
}

func TestAlternateContent_WriteXML(t *testing.T) {
	ac := NewAlternateContent()

	// Add a Choice with w14 requirement
	c := NewChoice()
	c.SetRequires("w14")
	innerChoice := NewCompositeElement(
		"http://example.com",
		"modern",
		"ex",
	)
	c.AppendChild(innerChoice)
	ac.AppendChild(c)

	// Add a Fallback
	fb := NewFallback()
	innerFallback := NewCompositeElement(
		"http://example.com",
		"legacy",
		"ex",
	)
	fb.AppendChild(innerFallback)
	ac.AppendChild(fb)

	// Write to XML
	xml := ac.OuterXml()

	// Check that it contains the expected elements
	if !strings.Contains(
		xml,
		"AlternateContent",
	) {
		t.Error(
			"OuterXml() missing AlternateContent element",
		)
	}
	if !strings.Contains(xml, "Choice") {
		t.Error(
			"OuterXml() missing Choice element",
		)
	}
	if !strings.Contains(xml, "Fallback") {
		t.Error(
			"OuterXml() missing Fallback element",
		)
	}
	if !strings.Contains(
		xml,
		"Requires=\"w14\"",
	) {
		t.Error(
			"OuterXml() missing Requires attribute",
		)
	}
}

func TestFileFormatVersion_String(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := tt.version.String()
			if got != tt.want {
				t.Errorf(
					"String() = %q, want %q",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestFileFormatVersion_Description(
	t *testing.T,
) {
	// Test a few key versions
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
			FileFormatVersionMicrosoft365,
			"Microsoft 365 (current)",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.version.String(),
			func(t *testing.T) {
				got := tt.version.Description()
				if got != tt.want {
					t.Errorf(
						"Description() = %q, want %q",
						got,
						tt.want,
					)
				}
			},
		)
	}
}

func TestFileFormatVersion_Comparison(
	t *testing.T,
) {
	// Test that version comparison works as expected
	if FileFormatVersionOffice2007 >= FileFormatVersionOffice2010 {
		t.Error(
			"Office2007 should be < Office2010",
		)
	}

	if FileFormatVersionOffice2013 >= FileFormatVersionOffice2016 {
		t.Error(
			"Office2013 should be < Office2016",
		)
	}

	if FileFormatVersionOffice2010 >= FileFormatVersionOffice2013 {
		t.Error(
			"Office2010 should be < Office2013",
		)
	}

	if FileFormatVersionOffice2024 <= FileFormatVersionOffice2021 {
		t.Error(
			"Office2024 should be > Office2021",
		)
	}
}
