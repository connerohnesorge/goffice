package elements

import (
	"testing"
)

func TestWebSettings(t *testing.T) {
	ws := NewWebSettings()

	if ws.LocalName() != "webSettings" {
		t.Errorf(
			"Expected LocalName 'webSettings', got %q",
			ws.LocalName(),
		)
	}

	if ws.NamespaceURI() != NamespaceWML {
		t.Errorf(
			"Expected NamespaceURI %q, got %q",
			NamespaceWML,
			ws.NamespaceURI(),
		)
	}
}

func TestWebSettingsOptimizeForBrowser(
	t *testing.T,
) {
	ws := NewWebSettings()

	if ws.OptimizeForBrowser() {
		t.Error(
			"Expected OptimizeForBrowser to be false initially",
		)
	}

	ws.SetOptimizeForBrowser(true)
	if !ws.OptimizeForBrowser() {
		t.Error(
			"Expected OptimizeForBrowser to be true",
		)
	}

	ws.SetOptimizeForBrowser(false)
	if ws.OptimizeForBrowser() {
		t.Error(
			"Expected OptimizeForBrowser to be false after disabling",
		)
	}
}

func TestWebSettingsAllowPNG(t *testing.T) {
	ws := NewWebSettings()

	if ws.AllowPNG() {
		t.Error(
			"Expected AllowPNG to be false initially",
		)
	}

	ws.SetAllowPNG(true)
	if !ws.AllowPNG() {
		t.Error("Expected AllowPNG to be true")
	}
}

func TestWebSettingsTargetScreenSize(
	t *testing.T,
) {
	ws := NewWebSettings()

	if ws.TargetScreenSize() != "" {
		t.Error(
			"Expected TargetScreenSize to be empty initially",
		)
	}

	ws.SetTargetScreenSize(
		TargetScreenSize1024x768,
	)
	if ws.TargetScreenSize() != TargetScreenSize1024x768 {
		t.Errorf(
			"Expected TargetScreenSize %q, got %q",
			TargetScreenSize1024x768,
			ws.TargetScreenSize(),
		)
	}

	ws.SetTargetScreenSize("")
	if ws.TargetScreenSize() != "" {
		t.Error(
			"Expected TargetScreenSize to be empty after clearing",
		)
	}
}

func TestWebSettingsEncoding(t *testing.T) {
	ws := NewWebSettings()

	if ws.Encoding() != "" {
		t.Error(
			"Expected Encoding to be empty initially",
		)
	}

	ws.SetEncoding("utf-8")
	if ws.Encoding() != "utf-8" {
		t.Errorf(
			"Expected Encoding 'utf-8', got %q",
			ws.Encoding(),
		)
	}

	ws.SetEncoding("")
	if ws.Encoding() != "" {
		t.Error(
			"Expected Encoding to be empty after clearing",
		)
	}
}

func TestWebSettingsPixelsPerInch(t *testing.T) {
	ws := NewWebSettings()

	// Default is 96
	if ws.PixelsPerInch() != 96 {
		t.Errorf(
			"Expected default PixelsPerInch 96, got %d",
			ws.PixelsPerInch(),
		)
	}

	ws.SetPixelsPerInch(72)
	if ws.PixelsPerInch() != 72 {
		t.Errorf(
			"Expected PixelsPerInch 72, got %d",
			ws.PixelsPerInch(),
		)
	}
}

func TestWebSettingsOnOffElements(t *testing.T) {
	ws := NewWebSettings()

	tests := []struct {
		name   string
		getter func() bool
		setter func(bool)
	}{
		{
			"DoNotUseLongFileNames",
			ws.DoNotUseLongFileNames,
			ws.SetDoNotUseLongFileNames,
		},
		{
			"RelyOnVML",
			ws.RelyOnVML,
			ws.SetRelyOnVML,
		},
		{
			"DoNotRelyOnCSS",
			ws.DoNotRelyOnCSS,
			ws.SetDoNotRelyOnCSS,
		},
		{
			"DoNotSaveAsSingleFile",
			ws.DoNotSaveAsSingleFile,
			ws.SetDoNotSaveAsSingleFile,
		},
		{
			"DoNotOrganizeInFolder",
			ws.DoNotOrganizeInFolder,
			ws.SetDoNotOrganizeInFolder,
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

func TestWebSettingsClone(t *testing.T) {
	ws := NewWebSettings()
	ws.SetOptimizeForBrowser(true)
	ws.SetEncoding("utf-8")

	cloned := ws.Clone()
	clone, ok := cloned.(*WebSettings)
	if !ok {
		t.Fatal(
			"Expected Clone to return *WebSettings",
		)
	}
	if !clone.OptimizeForBrowser() {
		t.Error(
			"Expected cloned OptimizeForBrowser to be true",
		)
	}
	if clone.Encoding() != "utf-8" {
		t.Errorf(
			"Expected cloned Encoding 'utf-8', got %q",
			clone.Encoding(),
		)
	}

	// Modify original
	ws.SetOptimizeForBrowser(false)
	if !clone.OptimizeForBrowser() {
		t.Error(
			"Clone should be independent of original",
		)
	}
}
