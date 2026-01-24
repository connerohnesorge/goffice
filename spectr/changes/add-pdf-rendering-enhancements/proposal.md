# Add PDF Rendering Enhancements

## Overview
Enhance PDF rendering fidelity to achieve pixel-perfect matching with Microsoft Office PDF export. Includes advanced typography, complex layouts, embedded fonts, and PDF/A compliance.

## Motivation
Current goffice PDF rendering handles common cases but lacks advanced features for production-quality PDF. Users need PDF output matching Office quality for professional documents, archival, and compliance.

## Goals
- **Advanced typography**: OpenType features (ligatures, kerning), complex scripts (Arabic, Thai, Indic), vertical text, ruby text
- **Layout accuracy**: Precise table cell sizing, floating elements, multi-column layouts, section headers/footers
- **Font handling**: Font subsetting, embedding with licensing, fallback chains, system font detection
- **PDF/A compliance**: PDF/A-1b, PDF/A-2b, PDF/A-3b with metadata and color space requirements
- **PDF features**: Hyperlinks, bookmarks, document outline/TOC, form fields, digital signature placeholders, tagged PDF for accessibility

## Non-Goals
- PDF editing or manipulation (generation only)
- Full PDF/X support (focus on PDF/A)
- Interactive forms with JavaScript
- PDF encryption (future enhancement)

## Dependencies
- Depends on: pdf core, all document types
- Related: all document format features (better input = better PDF)

## Technical Approach

### OpenType Features
```go
type FontRenderer struct {
    font *truetype.Font
    features []OpenTypeFeature
}

func (r *FontRenderer) RenderText(text string) []Glyph {
    // Apply ligatures
    // Apply kerning
    // Apply contextual alternates
}
```

### PDF/A Compliance
```go
type PDFAOptions struct {
    Level PDFALevel // 1b, 2b, 3b
    EmbedAllFonts bool
    ColorProfile string
}

func RenderPDFA(doc *Document, opts PDFAOptions) error {
    // Validate compliance
    // Embed fonts
    // Include metadata
}
```

## Estimated Effort
20 weeks (1 engineer with PDF expertise)

## Priority
P0 - Production PDF requires Office quality
