# PDF Rendering Fidelity

This document describes known differences between goffice-pdf output and Microsoft Office "Save as PDF" output.

## Philosophy

goffice-pdf aims for **visual equivalence** rather than byte-for-byte identical PDF output. The goal is that rendered documents appear identical to human viewers, even if the underlying PDF structure differs.

## Known Differences from Microsoft Office

### Word Documents

#### Text Layout

- **Line breaking**: goffice-pdf implements Unicode UAX #14 line breaking with Word-specific rules. Minor differences may occur in edge cases involving complex scripts or rare Unicode characters.
  - **Impact**: Low - affects <1% of documents
  - **Workaround**: Explicitly set line breaks in source document

- **Kerning**: Font kerning tables may be interpreted slightly differently between goffice-pdf and Word.
  - **Impact**: Low - typically sub-pixel differences
  - **Workaround**: None needed for most use cases

- **Justification**: Justification algorithms may distribute whitespace slightly differently.
  - **Impact**: Low - visually equivalent in most cases
  - **Workaround**: Use left/right/center alignment for critical layouts

#### Fonts

- **Font fallback**: When a specified font is not available, goffice-pdf uses a configurable fallback chain. The default fallback may differ from Word's.
  - **Impact**: Medium - can affect appearance if fonts are missing
  - **Workaround**: Ensure required fonts are available or configure custom fallback chain

- **Font subsetting**: goffice-pdf embeds only used glyphs by default (subset embedding). This reduces file size but may affect PDF editing.
  - **Impact**: Low - viewing is identical
  - **Workaround**: Use `FontEmbedding: EmbedFull` option if editing PDFs

#### Tables

- **Auto-fit columns**: Auto-fit column width calculation may differ slightly.
  - **Impact**: Low - typically within 1-2 points
  - **Workaround**: Use explicit column widths

- **Cell vertical alignment**: Vertical centering in table cells may differ by 1-2 points.
  - **Impact**: Low - visually equivalent
  - **Workaround**: None needed

#### Page Layout

- **Widow/orphan control**: Implementation matches Word's algorithm but edge cases may differ.
  - **Impact**: Low - rare pagination differences
  - **Workaround**: Manually insert page breaks if needed

- **Keep with next**: Supported but complex scenarios (nested tables, etc.) may differ.
  - **Impact**: Low - affects complex documents
  - **Workaround**: Simplify document structure

#### Not Supported

The following features are **not supported** and will be omitted or simplified:

- **Track changes**: Revisions are rendered in their accepted state
- **Comments**: Comments are not rendered in the PDF
- **Form fields**: Form fields are rendered as static text
- **Macros/VBA**: Macros are not executed or included
- **Custom XML**: Custom XML parts are ignored
- **ActiveX controls**: ActiveX controls are not rendered
- **Embedded objects**: OLE objects are not rendered (placeholder shown)

### Excel Spreadsheets

#### Cell Rendering

- **Number formatting**: Standard number formats match Excel. Custom formats may differ in edge cases.
  - **Impact**: Low - standard formats are identical
  - **Workaround**: Test custom formats and adjust if needed

- **Text wrapping**: Text wrapping in cells may break at slightly different points.
  - **Impact**: Low - visually equivalent
  - **Workaround**: Explicitly set row heights

- **Auto-fit rows**: Auto-fit row height calculation may differ by 1-2 points.
  - **Impact**: Low - typically imperceptible
  - **Workaround**: Use explicit row heights

#### Charts

- **Chart rendering**: Charts are rendered using DrawingML. Minor differences in axis labels, gridlines may occur.
  - **Impact**: Low - visually equivalent
  - **Workaround**: None needed for most use cases

- **Data labels**: Data label positioning may differ slightly.
  - **Impact**: Low - labels remain readable
  - **Workaround**: Manually position labels if critical

#### Page Layout

- **Fit to page**: Scaling algorithm matches Excel but rounding differences may occur.
  - **Impact**: Low - typically <1% difference
  - **Workaround**: Use explicit scaling percentage

- **Print titles**: Repeat rows/columns are supported and match Excel.
  - **Impact**: None - identical behavior
  - **Workaround**: None needed

#### Not Supported

- **Formulas**: Formulas are evaluated but live recalculation is not possible in PDF
- **Macros/VBA**: Macros are not included
- **Data validation**: Validation rules are not active in PDF
- **Slicers**: Slicers are rendered in their current state (not interactive)
- **PivotTables**: PivotTables are rendered as static tables
- **Sparklines**: Sparklines are rendered as small charts
- **External data connections**: External data is not refreshed

### PowerPoint Presentations

#### Slide Rendering

- **Animations**: Animations are not rendered. Slides appear in their final state.
  - **Impact**: High - animated content shows end state only
  - **Workaround**: Create separate slides for animation states if needed

- **Transitions**: Slide transitions are not rendered.
  - **Impact**: High - no transition effects in PDF
  - **Workaround**: None - inherent limitation of PDF format

- **Master slides**: Master slide inheritance is supported and matches PowerPoint.
  - **Impact**: None - identical behavior
  - **Workaround**: None needed

#### Shapes and Objects

- **3D effects**: Complex 3D effects may be simplified or rendered as 2D.
  - **Impact**: Medium - visual approximation of 3D
  - **Workaround**: Use 2D effects for critical elements

- **SmartArt**: SmartArt is decomposed to primitive shapes. Layout matches PowerPoint.
  - **Impact**: Low - visual result is equivalent
  - **Workaround**: None needed

- **Video/Audio**: Media files are not embedded or playable in PDF.
  - **Impact**: High - media content not available
  - **Workaround**: Include placeholder or screenshot

#### Not Supported

- **Embedded videos**: Videos are not embedded (placeholder shown)
- **Audio clips**: Audio is not embedded
- **Animations**: No animation support
- **Hyperlinks to slides**: Hyperlinks within presentation may not work
- **Action buttons**: Actions are not interactive

## Testing Strategy

### Automated Testing

We maintain fidelity test suites in `pdf/fidelity_test.go`:

- **Word**: 10+ test documents covering common scenarios
- **Excel**: 10+ test spreadsheets with various features
- **PowerPoint**: 10+ presentations with different layouts

Each test generates PDF output that can be compared with Office output.

### Manual Comparison

For critical documents:

1. Generate PDF using goffice-pdf
2. Generate PDF using Microsoft Office "Save as PDF"
3. Open both PDFs side-by-side
4. Visually compare each page
5. Measure any differences (tolerate <2pt variations)

### Visual Diff Tools

Recommended tools for visual comparison:

- **DiffPDF**: Side-by-side PDF comparison
- **Adobe Acrobat**: Compare documents feature
- **PDFtk**: Extract and compare metadata
- **pdfimages**: Extract and compare embedded images
- **goffice-pdf/comparison**: Automated pixel-level comparison (see below)

### Automated Visual Comparison

The `pdf/comparison` package provides automated pixel-level visual comparison for PDF fidelity testing. This tool compares generated PDFs against baseline reference images to detect rendering differences.

#### How It Works

1. **PDF-to-PNG Conversion**: Converts both PDFs to PNG images using Ghostscript
2. **Pixel-Level Comparison**: Compares images pixel-by-pixel using Euclidean color distance in RGB space
3. **Tolerance Thresholds**: Applies two-level filtering to ignore minor rendering differences
4. **Diff Visualization**: Generates annotated images highlighting differences in red

#### Running Visual Comparison Tests

The comparison capability is integrated into the fidelity test suite:

```bash
# Install Ghostscript (required dependency)
# Ubuntu/Debian:
sudo apt-get install ghostscript

# macOS:
brew install ghostscript

# Run fidelity tests with visual comparison
go test -run=TestFidelity ./...
```

Tests will automatically:
- Generate PDF from your document
- Convert to PNG images
- Compare against baseline images
- Report differences and generate annotated diff images

#### Tolerance Thresholds

The comparison uses two-level filtering to handle expected rendering variations:

**Per-Pixel Tolerance**: `2.0` delta (Euclidean distance in RGB space)
- **Range**: 0-441 (sqrt(255² + 255² + 255²))
- **Purpose**: Ignore minor color variations from anti-aliasing, font hinting, and subpixel rendering
- **Example**: A pixel differing by (1,1,1) in RGB has delta ≈1.73, which is within tolerance

**Per-Image Threshold**: `0.5%` of pixels allowed to exceed per-pixel tolerance
- **Purpose**: Ignore isolated pixel differences while catching systematic rendering errors
- **Rationale**: Anti-aliasing edges, rounding differences, and platform-specific font rendering can affect <0.5% of pixels without visual impact
- **Example**: On a 1275x1650 page (2.1M pixels), up to ~10,500 pixels can exceed tolerance before failing

#### Why These Thresholds?

These values were chosen through empirical testing of Office documents:

1. **Anti-Aliasing**: Different PDF renderers (Office vs goffice-pdf) may use slightly different anti-aliasing algorithms, causing 1-2 pixel color variations along edges
2. **Font Rendering**: Subpixel font rendering and hinting can cause minor color differences in text, especially at smaller sizes
3. **Rounding**: Coordinate and color space conversions involve rounding that can differ by ±1 unit
4. **Platform Differences**: System-specific rendering (Linux vs Windows vs macOS) can cause isolated pixel variations

A threshold of 2.0 per-pixel and 0.5% per-image catches genuine rendering bugs while tolerating expected platform variations.

#### When to Update Baselines

Update baseline images when:

1. **Intentional Rendering Changes**: You've improved rendering quality or fixed a bug
2. **Font Changes**: System fonts or font embedding logic has changed
3. **Layout Improvements**: Text layout, line breaking, or justification algorithms updated
4. **Image Compression**: Image encoding or compression parameters changed

**How to update baselines:**

```bash
# Generate new baseline from current PDF output
go test -run=TestFidelity -update-baselines ./...

# Or manually:
# 1. Generate PDF from test document
go test -run=TestFidelityWord/basic_text -args -save-pdf

# 2. Convert PDF to PNG baseline
gs -q -dNOPAUSE -dBATCH -dSAFER -sDEVICE=png16m -r150 \
   -sOutputFile=testdata/fidelity/word/basic_text.docx.baseline.png \
   testdata/fidelity/word/basic_text.pdf

# 3. Commit the new baseline
git add testdata/fidelity/word/basic_text.docx.baseline.png
git commit -m "Update baseline for basic_text test"
```

**Important**: Always visually inspect the new baseline before committing! Compare side-by-side with Office output to ensure the change is acceptable.

#### Interpreting Results

**Test Pass**: Difference is within tolerance
```
Diff: 0.12% pixels differ, 0.03% exceed tolerance (max delta: 4.2, avg delta: 1.8)
PASS: Visual comparison within tolerance
```

**Test Fail**: Too many pixels exceed tolerance
```
Diff: 2.34% pixels differ, 1.87% exceed tolerance (max delta: 45.6, avg delta: 12.3)
FAIL: Visual comparison exceeds tolerance (0.5%)
Diff image saved to: testdata/fidelity/word/basic_text.diff.png
```

Check the diff image to see what changed. Red highlights show pixels exceeding tolerance.

#### Troubleshooting

**Ghostscript not found:**
```
Ghostscript not found in PATH. Please install Ghostscript to enable visual comparison.
```
Install Ghostscript using your package manager (see installation instructions above).

**Baseline image missing:**
```
Baseline image not found: testdata/fidelity/word/example.docx.baseline.png
```
Generate a baseline image for this test case (see "How to update baselines" above).

**Image dimensions don't match:**
```
Image dimensions do not match: baseline 1275x1650, generated 1280x1650
```
This indicates a page size or scaling difference. Check your PDF generation settings and ensure they match the baseline.

## Reporting Fidelity Issues

When reporting fidelity issues:

1. **Provide source document**: Share the Office document (anonymized if needed)
2. **Describe expected output**: What does Office produce?
3. **Describe actual output**: What does goffice-pdf produce?
4. **Include versions**: Office version, goffice-pdf version
5. **Attach screenshots**: Show the difference visually
6. **Measure impact**: How significant is the difference?

## Improvement Roadmap

Areas for future fidelity improvements:

- [ ] Font metric extraction accuracy
- [ ] Complex table layout edge cases
- [ ] Advanced DrawingML 3D effects
- [ ] Custom number format parsing
- [ ] Complex SmartArt layouts
- [ ] Embedded object rendering (OLE)
- [ ] Advanced chart types (waterfall, funnel)

## Conclusion

goffice-pdf provides high-fidelity PDF rendering for the majority of Office documents. Known differences are documented and typically have minimal visual impact. For the small percentage of edge cases, workarounds are available or planned for future releases.
