//go:build darwin

// system_darwin.go provides macOS-specific system font discovery.

package font

// SystemFontPaths returns the directories where fonts are typically installed on macOS.
// This includes system fonts, library fonts, and user fonts.
func SystemFontPaths() []string {
	paths := []string{
		// System fonts - bundled with macOS
		"/System/Library/Fonts",

		// System fonts - supplemental
		"/System/Library/Fonts/Supplemental",

		// Library fonts - available to all users
		"/Library/Fonts",
	}

	// Add user-specific font directory
	// ~/Library/Fonts
	paths = append(
		paths,
		expandHome("~/Library/Fonts"),
	)

	// Network fonts (rarely used but included for completeness)
	paths = append(
		paths,
		"/Network/Library/Fonts",
	)

	// Some applications install fonts in their own directories
	// We include the common Application Support location
	paths = append(
		paths,
		expandHome(
			"~/Library/Application Support/Adobe/Fonts",
		),
	)
	paths = append(
		paths,
		"/Library/Application Support/Adobe/Fonts",
	)

	return paths
}
