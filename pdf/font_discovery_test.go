package pdf

import (
	"testing"
)

func TestDiscoverFonts(t *testing.T) {
	fonts := DiscoverFonts()

	// Check that we get some fonts (most systems have fonts)
	if len(fonts) == 0 {
		t.Skip(
			"No system fonts found - this might be expected in a minimal environment",
		)
	}

	// Check the structure of the first font entry
	if fonts[0].Family == "" {
		t.Error(
			"First font has empty Family field",
		)
	}

	if fonts[0].Style == "" {
		t.Error(
			"First font has empty Style field",
		)
	}

	t.Logf("Discovered %d fonts", len(fonts))
	t.Logf(
		"Sample font: %s (%s)",
		fonts[0].Family,
		fonts[0].Style,
	)
}
