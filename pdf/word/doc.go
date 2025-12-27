// Package word provides PDF rendering for WordprocessingML (DOCX) documents.
//
// This package converts parsed Word documents from the goffice/wordprocessing
// package into high-fidelity PDF output. It handles all Word-specific layout
// and formatting:
//
//   - Page layout (margins, orientation, headers, footers)
//   - Paragraph formatting (alignment, spacing, indentation)
//   - Character formatting (fonts, sizes, styles, colors)
//   - Tables (borders, cell merging, sizing)
//   - Lists (bullets, numbering, multi-level)
//   - Images and shapes (inline, floating, wrapped)
//   - Sections (columns, page breaks, section breaks)
//   - Table of contents and cross-references
//
// # Usage
//
//	import (
//		"github.com/connerohnesorge/goffice/wordprocessing"
//		"github.com/connerohnesorge/goffice-pdf/word"
//	)
//
//	doc, err := wordprocessing.Open("document.docx")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	renderer := word.NewRenderer()
//	err = renderer.Render(doc, "output.pdf")
//
// # Text Layout
//
// The renderer performs text layout including:
//   - Line breaking with proper Unicode support
//   - Hyphenation (optional)
//   - Justification with word and letter spacing
//   - Bidirectional text (RTL languages)
//
// # Fidelity
//
// The goal is visual fidelity to Microsoft Word's rendering. While perfect
// pixel-matching is not always possible, the output should be indistinguishable
// for typical documents.
package word
