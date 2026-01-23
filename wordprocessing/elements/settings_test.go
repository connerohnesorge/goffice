package elements

import (
	"testing"
)

func TestSettingsDocumentProtection(
	t *testing.T,
) {
	s := NewSettings()

	dp := s.GetOrCreateDocumentProtection()
	if dp == nil {
		t.Fatal(
			"Expected DocumentProtection to be created",
		)
	}

	// Set protection type
	dp.SetEdit(DocumentProtectionReadOnly)
	if dp.Edit() != DocumentProtectionReadOnly {
		t.Errorf(
			"Expected edit type %q, got %q",
			DocumentProtectionReadOnly,
			dp.Edit(),
		)
	}

	// Test enforcement
	if dp.Enforcement() {
		t.Error(
			"Expected enforcement to be false initially",
		)
	}
	dp.SetEnforcement(true)
	if !dp.Enforcement() {
		t.Error("Expected enforcement to be true")
	}

	// Test formatting restriction
	dp.SetFormatting(true)
	if !dp.Formatting() {
		t.Error(
			"Expected formatting restriction to be true",
		)
	}
}

func TestSettingsTrackRevisions(t *testing.T) {
	s := NewSettings()

	// Initially false
	if s.TrackRevisions() {
		t.Error(
			"Expected TrackRevisions to be false initially",
		)
	}

	s.SetTrackRevisions(true)
	if !s.TrackRevisions() {
		t.Error(
			"Expected TrackRevisions to be true",
		)
	}

	s.SetTrackRevisions(false)
	if s.TrackRevisions() {
		t.Error(
			"Expected TrackRevisions to be false after disabling",
		)
	}
}

func TestSettingsOnOffElements(t *testing.T) {
	s := NewSettings()

	tests := []struct {
		name   string
		getter func() bool
		setter func(bool)
	}{
		{
			"MirrorMargins",
			s.MirrorMargins,
			s.SetMirrorMargins,
		},
		{
			"EvenAndOddHeaders",
			s.EvenAndOddHeaders,
			s.SetEvenAndOddHeaders,
		},
		{
			"DisplayBackgroundShape",
			s.DisplayBackgroundShape,
			s.SetDisplayBackgroundShape,
		},
		{
			"HideSpellingErrors",
			s.HideSpellingErrors,
			s.SetHideSpellingErrors,
		},
		{
			"HideGrammaticalErrors",
			s.HideGrammaticalErrors,
			s.SetHideGrammaticalErrors,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.getter() {
				t.Errorf(
					"Expected %s to be false initially",
					tc.name,
				)
			}
			tc.setter(true)
			if !tc.getter() {
				t.Errorf(
					"Expected %s to be true",
					tc.name,
				)
			}
			tc.setter(false)
			if tc.getter() {
				t.Errorf(
					"Expected %s to be false after disabling",
					tc.name,
				)
			}
		})
	}
}

func TestSettingsCompatibility(t *testing.T) {
	s := NewSettings()

	c := s.GetOrCreateCompatibility()
	if c == nil {
		t.Fatal(
			"Expected Compatibility to be created",
		)
	}

	// Test UseFELayout
	if c.UseFELayout() {
		t.Error(
			"Expected UseFELayout to be false initially",
		)
	}
	c.SetUseFELayout(true)
	if !c.UseFELayout() {
		t.Error("Expected UseFELayout to be true")
	}

	// Test compatibility setting
	cs := c.AddCompatSetting(
		"testSetting",
		"http://example.com",
		"testValue",
	)
	if cs.Name() != "testSetting" {
		t.Errorf(
			"Expected name 'testSetting', got %q",
			cs.Name(),
		)
	}
	if cs.Uri() != "http://example.com" {
		t.Errorf(
			"Expected uri 'http://example.com', got %q",
			cs.Uri(),
		)
	}
	if cs.Val() != "testValue" {
		t.Errorf(
			"Expected val 'testValue', got %q",
			cs.Val(),
		)
	}
}

func TestSettingsDocumentVariables(t *testing.T) {
	s := NewSettings()

	dv := s.GetOrCreateDocumentVariables()
	if dv == nil {
		t.Fatal(
			"Expected DocumentVariables to be created",
		)
	}

	// Set variable
	dv.SetVariable("myVar", "myValue")
	if dv.GetVariable("myVar") != "myValue" {
		t.Errorf(
			"Expected variable value 'myValue', got %q",
			dv.GetVariable("myVar"),
		)
	}

	// Update variable
	dv.SetVariable("myVar", "newValue")
	if dv.GetVariable("myVar") != "newValue" {
		t.Errorf(
			"Expected variable value 'newValue', got %q",
			dv.GetVariable("myVar"),
		)
	}

	// Add another variable
	dv.SetVariable("anotherVar", "anotherValue")

	// Count variables
	count := 0
	for range dv.Variables() {
		count++
	}
	if count != 2 {
		t.Errorf(
			"Expected 2 variables, got %d",
			count,
		)
	}

	// Remove variable
	dv.RemoveVariable("myVar")
	if dv.GetVariable("myVar") != "" {
		t.Error("Expected variable to be removed")
	}
}

func TestSettingsProofState(t *testing.T) {
	s := NewSettings()

	ps := s.GetOrCreateProofState()
	if ps == nil {
		t.Fatal(
			"Expected ProofState to be created",
		)
	}

	ps.SetSpelling(ProofStateClean)
	if ps.Spelling() != ProofStateClean {
		t.Errorf(
			"Expected spelling state %q, got %q",
			ProofStateClean,
			ps.Spelling(),
		)
	}

	ps.SetGrammar(ProofStateDirty)
	if ps.Grammar() != ProofStateDirty {
		t.Errorf(
			"Expected grammar state %q, got %q",
			ProofStateDirty,
			ps.Grammar(),
		)
	}
}

func TestSettingsRevisionView(t *testing.T) {
	s := NewSettings()

	rv := s.GetOrCreateRevisionView()
	if rv == nil {
		t.Fatal(
			"Expected RevisionView to be created",
		)
	}

	// Default values should be true
	if !rv.Markup() {
		t.Error(
			"Expected Markup to be true by default",
		)
	}
	if !rv.Comments() {
		t.Error(
			"Expected Comments to be true by default",
		)
	}
	if !rv.InsertionsAndDeletions() {
		t.Error(
			"Expected InsertionsAndDeletions to be true by default",
		)
	}
	if !rv.Formatting() {
		t.Error(
			"Expected Formatting to be true by default",
		)
	}

	// Test setting to false
	rv.SetMarkup(false)
	if rv.Markup() {
		t.Error("Expected Markup to be false")
	}
}

func TestSettingsWriteProtection(t *testing.T) {
	s := NewSettings()

	wp := s.GetOrCreateWriteProtection()
	if wp == nil {
		t.Fatal(
			"Expected WriteProtection to be created",
		)
	}

	if wp.Recommended() {
		t.Error(
			"Expected Recommended to be false initially",
		)
	}

	wp.SetRecommended(true)
	if !wp.Recommended() {
		t.Error("Expected Recommended to be true")
	}
}

func TestSettingsRsidRoot(t *testing.T) {
	s := NewSettings()

	if s.RsidRoot() != "" {
		t.Error(
			"Expected RsidRoot to be empty initially",
		)
	}

	s.SetRsidRoot("00A12345")
	if s.RsidRoot() != "00A12345" {
		t.Errorf(
			"Expected RsidRoot '00A12345', got %q",
			s.RsidRoot(),
		)
	}
}

func TestSettingsClone(t *testing.T) {
	s := NewSettings()
	s.SetZoom(150)
	s.SetTrackRevisions(true)

	clone, ok := s.Clone().(*Settings)
	if !ok {
		t.Fatal("Clone did not return *Settings")
	}
	if clone.Zoom().Percent() != 150 {
		t.Errorf(
			"Expected cloned zoom 150, got %d",
			clone.Zoom().Percent(),
		)
	}
	if !clone.TrackRevisions() {
		t.Error(
			"Expected cloned TrackRevisions to be true",
		)
	}

	// Modify original
	s.SetZoom(200)
	if clone.Zoom().Percent() != 150 {
		t.Error(
			"Clone should be independent of original",
		)
	}
}
