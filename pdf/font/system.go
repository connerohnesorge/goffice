// system.go provides system font discovery functionality.
// This file contains common code used across all platforms.

package font

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// FontInfo contains metadata about a system font.
type FontInfo struct {
	Path   string     // Full path to the font file
	Family string     // Font family name
	Style  FontStyle  // Font style (Regular, Bold, Italic, etc.)
	Format FontFormat // Font format (TrueType, OpenType, etc.)
}

// FontRegistry is a registry of known fonts indexed by family name.
// It provides fast lookups for font files on the system.
type FontRegistry struct {
	fonts   map[string][]FontInfo // family (lowercase) -> variants
	indexed bool
	mu      sync.RWMutex
}

// ErrFontNotFound is returned when a requested font cannot be found.
var ErrFontNotFound = errors.New("font not found")

// defaultRegistry is the global default font registry.
var defaultRegistry = &FontRegistry{
	fonts: make(map[string][]FontInfo),
}

// defaultRegistryOnce ensures we only scan once for the default registry.
var defaultRegistryOnce sync.Once

// DefaultRegistry returns the default global font registry.
// On first access, it automatically scans system fonts.
func DefaultRegistry() *FontRegistry {
	defaultRegistryOnce.Do(func() {
		_ = defaultRegistry.Scan() // Ignore errors on initial scan
	})

	return defaultRegistry
}

// NewFontRegistry creates a new empty font registry.
func NewFontRegistry() *FontRegistry {
	return &FontRegistry{
		fonts: make(map[string][]FontInfo),
	}
}

// normalizeFamily normalizes a font family name for lookup.
func normalizeFamily(family string) string {
	return strings.ToLower(
		strings.TrimSpace(family),
	)
}

// Find searches for a font matching the given family and style.
// It returns the FontInfo if found, or an error if not found.
func (r *FontRegistry) Find(
	family string,
	style FontStyle,
) (FontInfo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := normalizeFamily(family)
	variants, ok := r.fonts[key]
	if !ok || len(variants) == 0 {
		return FontInfo{}, ErrFontNotFound
	}

	// First, try to find exact style match
	for _, info := range variants {
		if info.Style == style {
			return info, nil
		}
	}

	// If looking for BoldItalic, check if we have it
	if style == StyleBoldItalic {
		for _, info := range variants {
			if info.Style == StyleBoldItalic {
				return info, nil
			}
		}
	}

	// Fall back to regular style if available
	for _, info := range variants {
		if info.Style == StyleRegular {
			return info, nil
		}
	}

	// Return the first available variant
	return variants[0], nil
}

// Scan scans all system font directories and indexes the fonts.
func (r *FontRegistry) Scan() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Clear existing fonts
	r.fonts = make(map[string][]FontInfo)

	// Get system font paths
	paths := SystemFontPaths()

	// Scan each directory
	for _, dir := range paths {
		if err := r.scanDirectory(dir); err != nil {
			// Continue scanning other directories even if one fails
			continue
		}
	}

	r.indexed = true

	return nil
}

// scanDirectory scans a single directory for font files.
// Must be called with the lock held.
func (r *FontRegistry) scanDirectory(
	dir string,
) error {
	// Check if directory exists
	info, err := os.Stat(dir)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return nil
	}

	// Walk the directory
	return filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil // Skip files we can't access
			}
			if info.IsDir() {
				return nil // Continue into subdirectories
			}

			// Check if it's a font file
			ext := strings.ToLower(
				filepath.Ext(path),
			)
			if !isFontExtension(ext) {
				return nil
			}

			// Try to parse the font and extract metadata
			fontInfo, err := extractFontInfo(path)
			if err != nil {
				return nil // Skip fonts we can't parse
			}

			// Add to registry
			key := normalizeFamily(
				fontInfo.Family,
			)
			r.fonts[key] = append(
				r.fonts[key],
				fontInfo,
			)

			return nil
		},
	)
}

// isFontExtension returns true if the extension is a known font format.
func isFontExtension(ext string) bool {
	switch ext {
	case ".ttf", ".otf", ".ttc", ".otc":
		return true
	default:
		return false
	}
}

// extractFontInfo extracts metadata from a font file.
func extractFontInfo(
	path string,
) (FontInfo, error) {
	// Read the font file
	data, err := os.ReadFile(path)
	if err != nil {
		return FontInfo{}, err
	}

	// Detect format
	format := DetectFontFormat(data)
	if format == FormatUnknown {
		return FontInfo{}, errors.New(
			"unknown font format",
		)
	}

	// Parse the font
	font, err := ParseFont(data)
	if err != nil {
		return FontInfo{}, err
	}

	return FontInfo{
		Path:   path,
		Family: font.Family,
		Style:  font.Style,
		Format: format,
	}, nil
}

// IsIndexed returns true if the registry has been indexed.
func (r *FontRegistry) IsIndexed() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.indexed
}

// Families returns a list of all font families in the registry.
func (r *FontRegistry) Families() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	families := make([]string, 0, len(r.fonts))
	for family := range r.fonts {
		families = append(families, family)
	}

	return families
}

// Variants returns all variants of a font family.
func (r *FontRegistry) Variants(
	family string,
) []FontInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := normalizeFamily(family)
	variants := r.fonts[key]
	result := make([]FontInfo, len(variants))
	copy(result, variants)

	return result
}

// Add adds a font to the registry.
func (r *FontRegistry) Add(info FontInfo) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := normalizeFamily(info.Family)
	r.fonts[key] = append(r.fonts[key], info)
}

// Count returns the total number of fonts in the registry.
func (r *FontRegistry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	count := 0
	for _, variants := range r.fonts {
		count += len(variants)
	}

	return count
}

// FindSystemFont finds a system font matching the given family and style.
// This is a convenience function that uses the default registry.
func FindSystemFont(
	family string,
	style FontStyle,
) (string, error) {
	registry := DefaultRegistry()
	info, err := registry.Find(family, style)
	if err != nil {
		return "", err
	}

	return info.Path, nil
}

// ListSystemFonts returns a list of all available system fonts.
// This is a convenience function that uses the default registry.
func ListSystemFonts() ([]FontInfo, error) {
	registry := DefaultRegistry()

	registry.mu.RLock()
	defer registry.mu.RUnlock()

	var fonts []FontInfo
	for _, variants := range registry.fonts {
		fonts = append(fonts, variants...)
	}

	return fonts, nil
}

// IndexSystemFonts scans and indexes system fonts into the provided cache.
func IndexSystemFonts(cache *FontCache) error {
	if cache == nil {
		return errors.New("cache is nil")
	}

	// Scan system fonts
	registry := DefaultRegistry()

	registry.mu.RLock()
	defer registry.mu.RUnlock()

	// Load each font into the cache
	for _, variants := range registry.fonts {
		for _, info := range variants {
			// Load the font
			font, err := cache.LoadFile(info.Path)
			if err != nil {
				continue // Skip fonts that fail to load
			}

			// Ensure family and style are set
			if font.Family == "" {
				font.Family = info.Family
			}
			if font.Style == 0 &&
				info.Style != StyleRegular {
				font.Style = info.Style
			}
		}
	}

	return nil
}

// ScanDir scans a specific directory and adds fonts to the registry.
func (r *FontRegistry) ScanDir(
	dir string,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.scanDirectory(dir)
}

// Clear removes all fonts from the registry.
func (r *FontRegistry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.fonts = make(map[string][]FontInfo)
	r.indexed = false
}

// expandHome expands ~ to the user's home directory.
func expandHome(path string) string {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}

		return filepath.Join(home, path[1:])
	}

	return path
}
