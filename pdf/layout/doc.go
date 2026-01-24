//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package layout provides text layout algorithms for PDF rendering.
//
// This package implements the text layout engine required for rendering
// Office documents to PDF. It handles line breaking, text measurement,
// and paragraph layout according to Unicode standards.
//
// # Line Breaking
//
// The package implements Unicode UAX #14 (Unicode Line Breaking Algorithm)
// for determining where line breaks can occur in text. This ensures proper
// line breaking for all scripts including Latin, CJK, Arabic, and others.
//
// Example usage:
//
//	lb := layout.NewLineBreaker()
//	opportunities := lb.FindBreakOpportunities("Hello, world!")
//	for _, opp := range opportunities {
//	    fmt.Printf("Position %d: %s\n", opp.Position, opp.Type)
//	}
//
// # Text Measurement
//
// The package provides text measurement functions that work with font
// metrics to calculate the width of text runs. This is essential for
// determining where to break lines to fit within page margins.
//
// # Line Breaking Classes
//
// Characters are classified according to UAX #14 line breaking classes:
//
//   - Mandatory breaks: BK (Break), CR, LF, NL (Next Line)
//   - Break opportunities: SP (Space), ZW (Zero Width), etc.
//   - Non-break classes: GL (Glue), WJ (Word Joiner), etc.
//   - Alphabetic: AL (Ordinary Alphabetic)
//   - Numeric: NU (Numeric), etc.
//   - Punctuation: OP (Opening), CL (Closing), etc.
//   - CJK: ID (Ideographic), etc.
//
// # Integration
//
// This package is used by the word, spreadsheet, and presentation
// renderers to perform text layout before rendering to PDF.
package layout
