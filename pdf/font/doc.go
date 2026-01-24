//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package font provides font loading, metrics extraction, and subsetting for PDF generation.
//
// This package handles all font-related operations required for high-fidelity PDF rendering:
//   - Font file loading (TrueType, OpenType)
//   - Font metrics extraction (widths, heights, ascenders, descenders)
//   - Font subsetting for minimal file size
//   - Character-to-glyph mapping
//   - Font embedding into PDF documents
//
// # Supported Formats
//
// The package supports:
//   - TrueType fonts (.ttf)
//   - OpenType fonts with TrueType outlines (.otf with glyf table)
//   - OpenType fonts with CFF outlines (.otf with CFF table)
//   - TrueType Collection files (.ttc)
//
// # Font Resolution
//
// Fonts are resolved in the following order:
//  1. Embedded fonts from the Office document
//  2. System fonts matching the font name
//  3. Fallback to a default font (with warning)
//
// # Subsetting
//
// To minimize PDF file size, fonts are subsetted to include only the glyphs
// actually used in the document. This is critical for CJK fonts which can
// contain tens of thousands of glyphs.
package font
