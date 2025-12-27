//go:build windows

// system_windows.go provides Windows-specific system font discovery.

package font

import (
	"os"
	"path/filepath"
)

// SystemFontPaths returns the directories where fonts are typically installed on Windows.
// This includes the system fonts directory and user-specific font directories.
func SystemFontPaths() []string {
	paths := []string{}

	// System fonts directory (typically C:\Windows\Fonts)
	winDir := os.Getenv("WINDIR")
	if winDir == "" {
		winDir = os.Getenv("SystemRoot")
	}
	if winDir == "" {
		winDir = `C:\Windows`
	}
	paths = append(
		paths,
		filepath.Join(winDir, "Fonts"),
	)

	// User fonts directory (Windows 10 1803+ feature)
	// Located at %LOCALAPPDATA%\Microsoft\Windows\Fonts
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData != "" {
		userFonts := filepath.Join(
			localAppData,
			"Microsoft",
			"Windows",
			"Fonts",
		)
		paths = append(paths, userFonts)
	}

	// User profile fonts (alternative location)
	userProfile := os.Getenv("USERPROFILE")
	if userProfile != "" {
		// Some applications install fonts here
		paths = append(
			paths,
			filepath.Join(
				userProfile,
				"AppData",
				"Local",
				"Microsoft",
				"Windows",
				"Fonts",
			),
		)
	}

	// Program Files fonts (some applications install fonts here)
	programFiles := os.Getenv("ProgramFiles")
	if programFiles != "" {
		// Common locations for application fonts
		paths = append(
			paths,
			filepath.Join(
				programFiles,
				"Common Files",
				"Microsoft Shared",
				"Fonts",
			),
		)
	}

	// Program Files (x86) fonts
	programFilesX86 := os.Getenv(
		"ProgramFiles(x86)",
	)
	if programFilesX86 != "" {
		paths = append(
			paths,
			filepath.Join(
				programFilesX86,
				"Common Files",
				"Microsoft Shared",
				"Fonts",
			),
		)
	}

	return paths
}
