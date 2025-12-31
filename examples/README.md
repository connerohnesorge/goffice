# goffice Examples

This directory contains example programs demonstrating the capabilities of the goffice library.

## ML Results Spreadsheet Example

The `ml_results/ml_results_example.go` program creates a realistic, complex ML research results spreadsheet in XLSX format.

See `ml_results/README.md` for full documentation.

### Quick Start

```bash
cd ml_results
go run ml_results_example.go
```

This creates `ml_results_example.xlsx` with 4 sheets, 10+ models, formulas, and 3 charts.

---

## ML Paper Example

The `ml_paper_example.go` program creates a realistic, complex ML research paper in DOCX format (ml_paper_example.docx).

### Features Demonstrated

This example showcases goffice's comprehensive wordprocessing capabilities:

- **Multi-page document structure** with title page, abstract, and content sections
- **Hierarchical heading styles** (H1 and H2) with proper spacing
- **Professional formatting**:
  - Centered and right-aligned text
  - Justified paragraphs
  - Custom font sizes, colors, and weights (bold, italic)
  - Proper spacing before/after paragraphs
  - Text indentation
  
- **Complex table creation and styling**:
  - Multiple tables with headers and data rows
  - Header row formatting with shading
  - Data cells with custom background colors
  - Properly calculated column widths
  
- **Document structure**:
  - Page breaks between sections
  - Consistent styling throughout
  - Mathematical notation (subscripts and superscripts in text)
  - Mixed formatting within paragraphs (bold/italic/color)
  
- **Realistic academic content**:
  - Title page with authors and affiliations
  - Abstract section
  - Introduction with subsections
  - Related work section
  - Methodology with architecture and training descriptions
  - Experiments section with datasets and results tables
  - Analysis and discussion
  - Conclusion
  - References section

### Running the Example

```bash
go run examples/ml_paper_example.go
```

This creates `ml_paper_example.docx` in the current directory.

### Validating Output

The generated DOCX file is fully compatible with:
- Microsoft Word (all versions supporting .docx)
- LibreOffice Writer
- Google Docs
- Other OOXML-compatible tools

To validate the document structure:

```bash
# Convert to PDF for verification
libreoffice --headless --convert-to pdf ml_paper_example.docx

# Inspect the ZIP structure and XML
unzip -l ml_paper_example.docx
```

### Code Quality

The example demonstrates:
- Proper resource management (doc.Close())
- Type-safe element construction
- Fluent-style API usage
- Clean separation of document structure from content
- Professional Go code style and organization

### Document Statistics

- **File size**: ~5 KB (compressed DOCX)
- **Pages**: 10 pages when rendered
- **Content sections**: 6 major sections + references
- **Tables**: 3 (with various styling)
- **Formatting variety**: Bold, italic, colors, font families, sizes
- **Structure elements**: Page breaks, indentation, alignment

### Future Enhancements

To extend this example with additional features:
- Add images/charts using the image insertion API
- Include footnotes and endnotes
- Add headers and footers with page numbers
- Create custom styles
- Add table of contents
- Include comments and tracked changes

---

## DrawingML PDF Rendering Examples

The `pdf-rendering/` directory contains examples demonstrating the DrawingML to PDF rendering capabilities.

See `pdf-rendering/README.md` for full documentation.

### Quick Start

```bash
cd pdf-rendering

# Word images to PDF
cd word-images && go run main.go && cd ..

# Excel charts to PDF
cd excel-charts && go run main.go && cd ..

# PowerPoint shapes to PDF
cd powerpoint-shapes && go run main.go && cd ..
```

### What's Demonstrated

- **Word Images**: Inline image rendering with various sizes
- **Excel Charts**: Data foundation and chart rendering
- **PowerPoint Shapes**: Vector shapes, fills, and text rendering

These examples showcase the high-fidelity DrawingML to PDF rendering system in the `pdf/drawing/` package.
