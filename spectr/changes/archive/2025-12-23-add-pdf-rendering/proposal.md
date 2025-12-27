# Change: Add PDF Rendering for Office Documents

## Why

Users need to convert Office documents (Word, Excel, PowerPoint) to PDF format with exact fidelity matching Microsoft Office rendering. This is a fundamental requirement for document processing workflows, print previews, and archival purposes. Currently, goffice can read/write OOXML but cannot produce PDF output.

## What Changes

- **NEW**: Separate `goffice-pdf` module for PDF rendering
- **NEW**: Text layout engine with exact Microsoft Word fidelity
- **NEW**: WordprocessingML to PDF renderer (paragraphs, tables, lists, headers/footers)
- **NEW**: SpreadsheetML to PDF renderer (worksheets, cells, charts)
- **NEW**: PresentationML to PDF renderer (slides, shapes, animations as static)
- **NEW**: DrawingML rendering pipeline (shapes, images, effects, SmartArt)
- **NEW**: Font metrics system for precise text measurement
- **NEW**: Page layout algorithm matching Word's pagination

## Impact

- Affected specs: None (new capability in separate module)
- Affected code: New `pdf/` module alongside existing packages
- Dependencies: Will use Go PDF libraries (gofpdf, pdfcpu) - no CGO
- Breaking changes: None (additive only)

## Scope

| Document Type | Priority | Complexity |
|--------------|----------|------------|
| WordprocessingML (.docx) | P0 | Very High (text layout) |
| SpreadsheetML (.xlsx) | P1 | High (cell/chart layout) |
| PresentationML (.pptx) | P2 | High (slide composition) |
| DrawingML (shared) | P0 | High (shapes, images) |

## Key Challenges

1. **Text Layout Fidelity**: Microsoft's text layout algorithm is proprietary. Achieving exact match requires:
   - Font metrics extraction and caching
   - Line breaking algorithm (Unicode UAX #14 + Word-specific rules)
   - Paragraph pagination with widow/orphan control
   - Table cell content distribution

2. **Font Handling**: PDF requires embedded or referenced fonts
   - TrueType/OpenType font parsing for metrics
   - Font subsetting for embedded fonts
   - Fallback font mapping

3. **DrawingML Rendering**: Complex graphics pipeline
   - Shape geometry (preset shapes, custom paths)
   - Text in shapes with wrapping
   - Image placement and cropping
   - Effects (shadows, reflections, 3D)

## Success Criteria

- Word documents render with visually identical output to "Save as PDF" from Word
- Page breaks match Microsoft Word exactly
- All text formatting preserved (fonts, sizes, colors, effects)
- Tables render with correct cell sizes and borders
- Images and shapes positioned accurately
- Charts render with correct data visualization
