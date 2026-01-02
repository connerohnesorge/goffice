# PDF RENDERING KNOWLEDGE BASE

## OVERVIEW
High-fidelity PDF rendering engine for Office documents.
Matches Microsoft Office's output visually, including fonts, layout, and graphics.
Opt-in module (separate from core XML handling).

## STRUCTURE
```
pdf/
├── pdf.go          # Main API (RenderWord, RenderSpreadsheet, etc.)
├── doc.go          # Detailed documentation and examples
├── font/           # Font discovery, embedding, and subsetting
├── layout/         # Text layout engine (line breaking, justification)
├── drawing/        # DrawingML renderer (shapes, charts)
├── word/           # Word-specific renderer
├── spreadsheet/    # Excel-specific renderer
└── presentation/   # PowerPoint-specific renderer
```

## WHERE TO LOOK
| Task | Location | Notes |
|------|----------|-------|
| **Entry Points** | `pdf.go` | `RenderWord`, `RenderSpreadsheet`, `RenderPresentation` |
| **Config** | `pdf.go` | `RenderOptions` (DPI, PDF/A, FontEmbedding) |
| **Text Layout** | `layout/` | Complex text handling, bidirectional text |
| **Fonts** | `font/` | System font loading and fallback chains |
| **Images** | `pdf.go` | `ImageDPI`, `ImageCacheLimit` settings |

## CONVENTIONS
*   **Rendering Options:** Use `RenderOptions` to control quality vs. file size.
*   **Font Embedding:** `EmbedSubset` is default (best balance). `EmbedFull` for archival.
*   **PDF/A:** Supports PDF/A-1b, 2b, 3b for archival compliance.
*   **Stateless:** Renderers are generally stateless per-document; `RenderOptions` controls behavior.
*   **Fidelity:** Prioritizes visual matching of MS Office over exact internal structure mapping.

## SUPPORT
*   **Word:** High maturity.
*   **Excel:** Supported (check current status for advanced features).
*   **PowerPoint:** Supported (animations rendered as static).
