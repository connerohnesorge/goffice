# Change: Enhance PDF Rendering Fidelity for Complete Office Format Support

## Why
goffice implements PDF rendering for Word, Excel, and PowerPoint documents but lacks comprehensive support for complex formatting and advanced features present in Open-XML-SDK:
- Advanced font metrics and kerning
- Complex text rendering (right-to-left, bidirectional, ligatures)
- Shape effects (shadows, glows, reflections, 3D)
- Advanced fill and stroke rendering
- Gradient and pattern fills
- Chart rendering with all advanced options
- Table cell merging and complex layouts
- Header/footer rendering with different sections
- Form field rendering
- Comment rendering in PDF
- Watermarks
- Linked PDF pages with named destinations
- Text rotation and vertical text
- OLE object rendering

## What Changes
- Add advanced font rendering with kerning and OpenType features
- Add bidirectional text support (Arabic, Hebrew, etc.)
- Add shape effect rendering (shadows, glows, reflections)
- Add gradient fill rendering (linear, radial, path)
- Add pattern fill rendering
- Add improved chart rendering with proper formatting
- Add table merging and complex layout support
- Add header/footer section-specific rendering
- Add form field rendering with proper widgets
- Add comment/note rendering as annotations
- Add watermark rendering
- Add PDF link annotations for hyperlinks
- Add text rotation and vertical text support
- Add proper font subsetting for embedded fonts
- Add PDF compression and optimization
- BREAKING: RenderOptions API may be extended

## Impact
- Affected specs: pdf-core, pdf-word, pdf-spreadsheet, pdf-presentation, pdf-drawing
- Affected code: pdf/, spreadsheet/pdf*, wordprocessing/pdf*, presentation/pdf*
- New dependencies: Possibly golang.org/x/text for advanced text rendering
- Test coverage required: Visual comparison tests for rendered PDF outputs
