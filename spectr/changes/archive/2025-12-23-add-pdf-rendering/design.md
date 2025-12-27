# PDF Rendering Design Document

## Context

goffice is a pure Go SDK for Office Open XML documents. Users have requested PDF rendering with exact fidelity matching Microsoft Office output. This requires a sophisticated text layout engine and graphics rendering pipeline.

**Constraints:**
- Separate module (`goffice-pdf`) to keep core SDK dependency-free
- Allow Go PDF libraries (no CGO)
- Must achieve exact fidelity with Microsoft Office rendering
- Support Word, Excel, and PowerPoint document types

**Stakeholders:**
- Document processing pipelines requiring PDF output
- Print preview implementations
- Archival systems requiring PDF/A output

## Goals / Non-Goals

### Goals
- Pixel-perfect PDF output matching Microsoft Word's "Save as PDF"
- Support all WordprocessingML, SpreadsheetML, and PresentationML elements
- Efficient rendering for large documents (streaming where possible)
- Font subsetting to minimize PDF file size
- Accessible PDF output (tagged PDF for screen readers)

### Non-Goals
- Interactive PDF features (forms, JavaScript)
- PDF editing (read-only output)
- PDF/X for print production (initially)
- Animation rendering for presentations (static snapshots only)

## Decisions

### Decision 1: Module Structure

**Choice:** Separate `goffice-pdf` module with its own go.mod

**Rationale:**
- Keeps core goffice dependency-free (pure stdlib)
- Users only import PDF module if needed
- Allows PDF-specific dependencies without bloating core
- Clear separation of concerns

**Structure:**
```
goffice-pdf/
├── go.mod                    # github.com/connerohnesorge/goffice-pdf
├── pdf.go                    # Main API entry points
├── word/                     # WordprocessingML renderer
│   ├── renderer.go           # Document → PDF conversion
│   ├── text_layout.go        # Text layout engine
│   ├── table_layout.go       # Table layout engine
│   └── page_layout.go        # Pagination algorithm
├── spreadsheet/              # SpreadsheetML renderer
│   ├── renderer.go
│   ├── cell_layout.go
│   └── chart_render.go
├── presentation/             # PresentationML renderer
│   ├── renderer.go
│   └── slide_render.go
├── drawing/                  # DrawingML shared rendering
│   ├── shape.go
│   ├── image.go
│   ├── effects.go
│   └── text_box.go
├── font/                     # Font handling
│   ├── metrics.go            # Font metrics extraction
│   ├── subset.go             # Font subsetting
│   └── fallback.go           # Fallback font mapping
├── layout/                   # Layout primitives
│   ├── box.go                # Box model
│   ├── line.go               # Line breaking
│   └── page.go               # Page model
└── internal/
    └── pdfgen/               # PDF generation abstraction
```

### Decision 2: PDF Generation Library

**Choice:** Build on `pdfcpu` for PDF primitives + custom layout engine

**Alternatives Considered:**
| Library | Pros | Cons |
|---------|------|------|
| `gofpdf` | Simple API, well-documented | Archived, limited features |
| `pdfcpu` | Active development, PDF/A support, manipulation | Lower-level API |
| `unipdf` | Full-featured | Commercial license |
| Custom | Full control | Massive effort |

**Rationale:**
- `pdfcpu` is actively maintained and MIT licensed
- Supports PDF/A for archival
- Lower-level API gives us control for exact layout
- Can add font subsetting and advanced features

### Decision 3: Text Layout Engine

**Choice:** Custom text layout engine implementing Word's algorithm

**Architecture:**
```go
// TextLayoutEngine handles Word-compatible text layout
type TextLayoutEngine struct {
    fontCache    *FontCache
    lineBreaker  *LineBreaker
    paraLayout   *ParagraphLayout
}

// Layout converts paragraph content to positioned glyphs
func (e *TextLayoutEngine) Layout(para *wordprocessing.Paragraph,
    width float64) []LayoutLine {
    // 1. Resolve fonts and get metrics
    // 2. Apply Unicode line breaking (UAX #14)
    // 3. Apply Word-specific adjustments
    // 4. Handle justification
    // 5. Return positioned glyph runs
}
```

**Key Algorithms:**
1. **Line Breaking**: Unicode UAX #14 + Word-specific rules
   - Soft hyphens, non-breaking spaces
   - CJK character breaking rules
   - Word's specific break opportunities

2. **Justification**: Word uses a specific algorithm
   - Inter-word spacing first
   - Inter-character spacing for CJK
   - Kashida insertion for Arabic

3. **Pagination**: Word's pagination algorithm
   - Widow/orphan control
   - Keep with next
   - Page break before
   - Section breaks

### Decision 4: Font Handling

**Choice:** TrueType/OpenType parsing with metrics extraction and subsetting

**Components:**
```go
// FontCache manages font metrics and data
type FontCache struct {
    fonts    map[string]*Font
    fallback *FallbackChain
}

// Font represents a parsed font with metrics
type Font struct {
    Family    string
    Style     FontStyle
    Metrics   FontMetrics
    GlyphData map[rune]GlyphMetrics
    // For embedding
    Data      []byte
}

// FontMetrics contains font-level measurements
type FontMetrics struct {
    UnitsPerEm   int
    Ascender     int
    Descender    int
    LineGap      int
    CapHeight    int
    XHeight      int
}
```

**Font Sources (priority order):**
1. Embedded fonts in document
2. System fonts (OS-specific paths)
3. Bundled fallback fonts (for consistent cross-platform output)

### Decision 5: Coordinate System

**Choice:** Use PDF points (1/72 inch) internally, convert from EMU

**Conversion:**
```go
const (
    EMUPerInch   = 914400
    PointsPerInch = 72
    EMUPerPoint  = EMUPerInch / PointsPerInch // 12700
)

func EMUToPoints(emu int64) float64 {
    return float64(emu) / float64(EMUPerPoint)
}
```

### Decision 6: Rendering Pipeline

**Choice:** Two-phase rendering (layout → paint)

**Phase 1 - Layout:**
```
Document → Layout Tree → Page Boxes
```
- Parse document structure
- Calculate all positions and sizes
- Paginate content
- Output: List of pages with positioned elements

**Phase 2 - Paint:**
```
Page Boxes → PDF Primitives → PDF Stream
```
- Iterate positioned elements
- Generate PDF drawing commands
- Embed fonts and images
- Output: PDF file

**Rationale:**
- Separation allows layout validation before PDF generation
- Can cache layout for preview scenarios
- Easier to test layout independently

### Decision 7: DrawingML Rendering

**Choice:** Shared rendering module for all document types

**Shape Rendering Pipeline:**
```
DrawingML Element → Path/Geometry → Fill/Stroke → Effects → PDF
```

**Supported Elements:**
| Element | Approach |
|---------|----------|
| Preset shapes | Geometry lookup table (190 shapes) |
| Custom shapes | Path parsing and conversion |
| Pictures | Image embedding with transforms |
| Charts | Chart-specific rendering |
| SmartArt | Decompose to primitives |
| Effects | Shadow, reflection, glow rendering |

### Decision 8: Streaming and Memory

**Choice:** Page-at-a-time streaming for large documents

**Strategy:**
- Layout computes one section at a time
- Pages written to PDF as completed
- Images loaded on demand and released after use
- Font data cached but glyphs subset at end

```go
// RenderOptions configures rendering behavior
type RenderOptions struct {
    // Stream pages as they're rendered (for progress)
    OnPageComplete func(pageNum int, data []byte)

    // Memory limit for image cache
    ImageCacheLimit int64

    // Font embedding mode
    FontEmbedding FontEmbedMode
}
```

## Risks / Trade-offs

| Risk | Impact | Mitigation |
|------|--------|------------|
| Exact fidelity impossible | High | Accept "visually equivalent" for edge cases, document known differences |
| Font availability | Medium | Bundle core fonts, clear error messages for missing fonts |
| Complex table layouts | High | Phased implementation, Word-generated test suite |
| Performance for large docs | Medium | Streaming, profiling, optimization passes |
| DrawingML complexity | High | Prioritize common shapes, defer complex effects |

## Migration Plan

Not applicable - this is a new module with no existing functionality to migrate.

## Open Questions

1. **Font licensing**: Can we bundle fallback fonts? Which ones are freely distributable?
2. **PDF/A compliance**: Which PDF/A level to target (1b, 2b, 3b)?
3. **Accessibility**: How much effort for tagged PDF / accessibility features?
4. **Testing strategy**: How to validate "exact fidelity" - pixel comparison vs structural?

## Implementation Phases

### Phase 1: Core Infrastructure (Foundation)
- PDF generation abstraction layer
- Font metrics extraction
- Basic text layout engine
- Simple paragraph rendering

### Phase 2: Word Document Rendering
- Full paragraph styling
- Table layout
- Lists and numbering
- Headers/footers
- Page layout and breaks

### Phase 3: DrawingML Rendering
- Preset shapes
- Images
- Text boxes
- Basic effects

### Phase 4: Spreadsheet Rendering
- Cell layout
- Cell formatting
- Simple charts
- Print area handling

### Phase 5: Presentation Rendering
- Slide rendering
- Shape composition
- Master slide inheritance
- Notes pages

### Phase 6: Polish and Optimization
- Font subsetting
- PDF/A output
- Performance optimization
- Accessibility (tagged PDF)
