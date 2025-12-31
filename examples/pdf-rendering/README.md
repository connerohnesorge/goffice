# DrawingML PDF Rendering Examples

This directory contains examples demonstrating the DrawingML to PDF rendering capabilities of the goffice library.

## Overview

These examples showcase how goffice renders Office Open XML (OOXML) DrawingML elements to PDF with high fidelity:

- **Word Images**: Inline images with various sizes
- **Excel Charts**: Data visualization and chart rendering (data foundation)
- **PowerPoint Shapes**: Vector shapes, fills, and text

## Examples

### 1. Word Images (`word-images/`)

Demonstrates rendering Word documents with embedded images to PDF.

**Features:**
- Multiple image sizes (small, medium, large)
- Image placement and scaling
- Mixed text and image content
- Professional layout

**Run:**
```bash
cd examples/pdf-rendering/word-images
go run main.go
```

**Output:**
- `word_images_example.docx` - Word document with images
- `word_images_example.pdf` - Rendered PDF

### 2. Excel Charts (`excel-charts/`)

Demonstrates rendering Excel workbooks with chart data to PDF.

**Features:**
- Multiple sheets (Sales Data, Product Performance, Summary)
- Quarterly sales tracking
- Product performance metrics
- Data foundation for chart rendering

**Run:**
```bash
cd examples/pdf-rendering/excel-charts
go run main.go
```

**Output:**
- `excel_charts_example.xlsx` - Excel workbook with data
- `excel_charts_example.pdf` - Rendered PDF

**Note:** Charts embedded in Excel files would be rendered using the DrawingML chart rendering system. This example demonstrates the data foundation.

### 3. PowerPoint Shapes (`powerpoint-shapes/`)

Demonstrates rendering PowerPoint presentations with shapes to PDF.

**Features:**
- Basic shapes (rectangles, ellipses, triangles, diamonds)
- Advanced shapes (lines, rounded rectangles)
- Shapes with text content
- Various fill colors
- Professional slide layouts

**Run:**
```bash
cd examples/pdf-rendering/powerpoint-shapes
go run main.go
```

**Output:**
- `powerpoint_shapes_example.pptx` - PowerPoint presentation with shapes
- `powerpoint_shapes_example.pdf` - Rendered PDF

## DrawingML Rendering System

The DrawingML rendering system in `pdf/drawing/` provides:

### Core Components

- **ShapeRenderer**: Renders DrawingML shapes with preset geometries
- **ChartRenderer**: Renders charts (bar, line, pie)
- **FillRenderer**: Handles solid fills, gradients, patterns
- **StrokeRenderer**: Renders shape outlines and borders
- **ImageRenderer**: Renders embedded images
- **TextInShapeRenderer**: Renders text within shape bounds

### Supported Shapes

**Basic Shapes:**
- Rectangle, RoundedRectangle
- Ellipse, Circle
- Triangle, Diamond

**Polygons:**
- Pentagon, Hexagon, Octagon

**Stars:**
- 4-point, 5-point, 6-point, 8-point, 10-point, 12-point stars

**Arrows:**
- Right, Left, Up, Down arrows

### Rendering Fidelity

The system aims for high fidelity with Microsoft Office:
- Accurate shape geometry
- Color matching (RGB, theme colors)
- Proper scaling and positioning
- Text layout within shapes
- Professional PDF output

## Usage Patterns

### Basic Workflow

1. **Create Office Document**
   ```go
   doc, _ := wordprocessing.New("doc.docx", wordprocessing.DocTypeDocument)
   // Add content with DrawingML elements
   doc.Save()
   doc.Close()
   ```

2. **Render to PDF**
   ```go
   doc, _ := wordprocessing.Open("doc.docx", false)
   defer doc.Close()

   fontCache := font.NewFontCache(10)
   engine := layout.NewTextLayoutEngine(fontCache)
   renderer, _ := pdfword.NewWordRenderer(doc, engine)

   renderer.Render("output.pdf")
   ```

### Adding Images to Word

```go
// Add image part
imagePart, _ := mainPart.AddImagePart(parts.ImageTypePng)
imagePart.FeedDataBytes(pngData)
relID := imagePart.RelationshipID()

// Create inline drawing
width := int64(3.0 * float64(elements.EMUsPerInch))
height := int64(2.0 * float64(elements.EMUsPerInch))
drawing := elements.NewInlineDrawing(width, height, relID)

// Add to paragraph
run.AppendChild(drawing)
```

### Adding Shapes to PowerPoint

```go
slide, _ := doc.AddSlide()
slideElem := slide.Slide()

shape := slideElem.AddShape()
shape.SetText("My Shape")
shape.SetPosition(914400, 914400)  // 1 inch from left/top (EMU)
shape.SetSize(2743200, 1371600)    // 3x1.5 inches (EMU)
shape.SetShapeType(elements.ShapeTypeRectangle)
shape.SetSolidFill("4472C4")       // Blue fill
```

## Coordinate System

### EMU (English Metric Units)

Office documents use EMU for measurements:
- 1 inch = 914,400 EMU
- 1 cm = 360,000 EMU
- 1 point = 12,700 EMU

**Common conversions:**
```go
// 1 inch from left
x := int64(914400)

// 2.5 inches from top
y := int64(2.5 * float64(elements.EMUsPerInch))

// 4 inches wide
width := int64(4 * 914400)
```

### PDF Coordinates

PDF uses points (1/72 inch) with bottom-left origin. The rendering system handles conversion automatically.

## Testing

Run the examples to verify DrawingML rendering:

```bash
# Test all examples
cd examples/pdf-rendering

cd word-images && go run main.go && cd ..
cd excel-charts && go run main.go && cd ..
cd powerpoint-shapes && go run main.go && cd ..

# Verify PDFs
find . -name "*.pdf" -exec ls -lh {} \;
```

**Verify output with:**
- Adobe Acrobat Reader
- Preview (macOS)
- LibreOffice Draw
- pdfinfo/pdftoppm tools

## Further Reading

- **[pdf/DRAWING_RENDERING.md](../../pdf/DRAWING_RENDERING.md)** - Developer guide for DrawingML rendering
- **[pdf/FIDELITY.md](../../pdf/FIDELITY.md)** - Rendering fidelity guidelines
- **[pdf/LIMITATIONS.md](../../pdf/LIMITATIONS.md)** - Known limitations
- **[CLAUDE.md](../../CLAUDE.md)** - Project development guide

## Requirements

- Go 1.25+
- goffice library
- System fonts for text rendering

## License

See the project LICENSE file.
