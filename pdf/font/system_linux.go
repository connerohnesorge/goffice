//go:build linux

// system_linux.go provides Linux-specific system font discovery.

package font

// SystemFontPaths returns the directories where fonts are typically installed on Linux.
// This includes system directories, user directories, and XDG-compliant locations.
func SystemFontPaths() []string {
	paths := []string{
		"/usr/share/fonts",
		"/usr/local/share/fonts",
	}

	// Add user-specific font directories
	// ~/.fonts is the legacy location
	paths = append(paths, expandHome("~/.fonts"))

	// ~/.local/share/fonts is the XDG-compliant location
	paths = append(
		paths,
		expandHome("~/.local/share/fonts"),
	)

	// Check XDG_DATA_HOME for custom data directory
	// Default is ~/.local/share
	xdgDataHome := expandHome("~/.local/share")
	if xdgDataHome != "" {
		xdgFonts := xdgDataHome + "/fonts"
		// Add if not already included
		found := false
		for _, p := range paths {
			if p == xdgFonts {
				found = true

				break
			}
		}
		if !found {
			paths = append(paths, xdgFonts)
		}
	}

	// Some distributions also use these directories
	additionalPaths := []string{
		"/usr/share/fonts/truetype",
		"/usr/share/fonts/opentype",
		"/usr/share/fonts/TTF",
		"/usr/share/fonts/OTF",
		"/usr/X11R6/lib/X11/fonts/TTFont",
	}

	paths = append(paths, additionalPaths...)

	return paths
}
