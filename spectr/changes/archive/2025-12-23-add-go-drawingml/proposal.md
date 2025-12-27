# Change: Add Go DrawingML Shared Package

## ID

add-go-drawingml

## Title

Go DrawingML Shared Package

## Summary

Shared package providing DrawingML types for shapes, images, effects, and charts used by Word, Presentation, and Spreadsheet SDKs. DrawingML is the common graphics layer in Office Open XML documents that handles all visual elements across document types.

## Why

DrawingML (Drawing Markup Language) is the shared graphics foundation used by Word, Excel, and PowerPoint documents. Rather than duplicating DrawingML types in each document SDK, this proposal creates a single shared package that all three SDKs can import.

**Problem Statement:**
- DrawingML types are used identically across Word, Presentation, and Spreadsheet documents
- Duplicating these types would lead to code maintenance burden and inconsistency
- Without a shared package, each SDK would need to define shapes, fills, effects, and text formatting independently
- The Open-XML-SDK for C# uses shared DrawingML types; Go implementation should follow the same pattern

**Opportunity:**
- Single source of truth for all DrawingML types
- Consistent API for shapes, effects, and text across all document types
- Easier maintenance and testing
- Matches the architecture of Microsoft's Open-XML-SDK

## What Changes

### New Capabilities

1. **Core DrawingML Types** (`drawingml/main/`)
   - Transform2D (a:xfrm) - position, size, rotation, flip
   - Preset Geometry (a:prstGeom) - 100+ built-in shapes
   - Custom Geometry (a:custGeom) - path-based custom shapes
   - Line Properties (a:ln) - width, style, compound, dash patterns

2. **Fill Types** (`drawingml/main/`)
   - Solid Fill (a:solidFill) - single color fills
   - Gradient Fill (a:gradFill) - linear, radial, path gradients
   - Pattern Fill (a:pattFill) - hatch and pattern fills
   - Blip Fill (a:blipFill) - image fills
   - Group Fill (a:grpFill) - inherit from parent group
   - No Fill (a:noFill) - transparent

3. **Effect System** (`drawingml/main/`)
   - Effect List (a:effectLst) - container for effects
   - Shadow effects - outer, inner shadows
   - Glow effects
   - Reflection effects
   - Soft Edge effects
   - Blur effects

4. **Text System** (`drawingml/main/`)
   - Text Body (a:txBody) - text container
   - Body Properties (a:bodyPr) - text box settings
   - Paragraph (a:p) - paragraph with properties
   - Run (a:r) - text run with formatting
   - Run Properties (a:rPr) - font, size, bold, italic, etc.
   - Text content (a:t)

5. **Color System** (`drawingml/main/`)
   - sRGB Color (a:srgbClr) - RGB hex colors
   - Scheme Color (a:schemeClr) - theme-based colors
   - System Color (a:sysClr) - system colors
   - HSL Color (a:hslClr) - HSL color model
   - Preset Color (a:prstClr) - named colors
   - Color transforms - tint, shade, saturation, luminance, alpha

6. **Blip/Image Support** (`drawingml/main/`)
   - Blip (a:blip) - image reference
   - Embedded images (r:embed)
   - Linked images (r:link)
   - Image effects and compression

7. **Charts** (`drawingml/chart/`)
   - Chart Space (c:chartSpace) - chart container
   - Chart (c:chart) - chart definition
   - Plot Area (c:plotArea) - plotting region
   - All chart types: Bar, Column, Line, Area, Pie, Doughnut, Scatter, Bubble, Radar, Surface, Stock
   - Axis types: Category, Value, Date, Series
   - Legend, Title, Data Labels

8. **Spreadsheet Drawing** (`drawingml/spreadsheet/`)
   - Worksheet Drawing (xdr:wsDr) - drawing container
   - Two Cell Anchor (xdr:twoCellAnchor) - cell range positioning
   - One Cell Anchor (xdr:oneCellAnchor) - single cell positioning
   - Absolute Anchor (xdr:absoluteAnchor) - absolute positioning
   - Shape, Picture, GraphicFrame, GroupShape, ConnectionShape

9. **Word Drawing** (`drawingml/word/`)
   - Inline (wp:inline) - inline with text
   - Anchor (wp:anchor) - floating positioning
   - Text wrapping modes - none, square, tight, through, top-and-bottom
   - Position controls - horizontal, vertical alignment

10. **Picture** (`drawingml/picture/`)
    - Picture element (pic:pic)
    - Non-visual properties (pic:nvPicPr)
    - Blip fill (pic:blipFill)
    - Shape properties (pic:spPr)

### BREAKING Changes
- None (new package)

### Non-Breaking Changes
- Introduces new `drawingml/` package directory with subpackages

## Impact

### Affected Specs
- `add-go-word-sdk` - Will import drawingml/ for image and shape support
- `add-go-presentation-sdk` - Will import drawingml/ for all visual elements
- `add-go-spreadsheet-sdk` - Will import drawingml/ for charts and drawings

### Affected Code
- New package: `github.com/connerohnesorge/goffice/drawingml/main` - Core DrawingML types (a: namespace)
- New package: `github.com/connerohnesorge/goffice/drawingml/chart` - Chart types (c: namespace)
- New package: `github.com/connerohnesorge/goffice/drawingml/spreadsheet` - Spreadsheet drawing (xdr: namespace)
- New package: `github.com/connerohnesorge/goffice/drawingml/word` - Word drawing (wp: namespace)
- New package: `github.com/connerohnesorge/goffice/drawingml/picture` - Picture types (pic: namespace)

### Dependencies
- `encoding/xml` - Standard library XML support
- `sync` - Standard library for sync.RWMutex
- No external dependencies

### Risks
- **Complexity**: DrawingML has many interconnected types; careful interface design needed
- **Namespace collision**: Different document types use DrawingML slightly differently
- **Performance**: Large documents with many shapes may need lazy loading
- **Schema versions**: Minor differences between Office versions

### Migration
- N/A - This is a new capability with no existing implementation to migrate

## Configuration

| Setting | Value | Rationale |
|---------|-------|-----------|
| **Go Version** | 1.25+ | Range-over-func for iterators, modern features |
| **Thread Safety** | sync.RWMutex | Concurrent reads, exclusive writes |
| **Code Generation** | Pre-generate and commit to repo | Users don't need generator |
| **Validation** | Strict by default | Reject invalid content, return errors |
| **EMU Units** | Native EMU storage with conversion helpers | 914400 EMUs = 1 inch |

## DrawingML Components

Based on exploration of ECMA-376 and Open-XML-SDK:

### Shapes
- **Preset Geometries** (a:prstGeom) - 100+ built-in shapes including:
  - Basic: rect, ellipse, roundRect, triangle, rtTriangle
  - Block arrows: rightArrow, leftArrow, upArrow, downArrow, leftRightArrow
  - Stars and banners: star4, star5, star6, irregularSeal1, ribbon2
  - Callouts: wedgeRoundRectCallout, cloudCallout
  - Flowchart: flowChartProcess, flowChartDecision, flowChartTerminator
  - And many more...
- **Custom Geometries** (a:custGeom) - path-based shapes with:
  - Path list (a:pathLst) with moveTo, lineTo, arcTo, cubicBezTo, close
  - Adjust values (a:avLst) for parameterized shapes
  - Guide definitions (a:gdLst) for calculations

### Fills
- **Solid Fill** (a:solidFill) - single color
- **Gradient Fill** (a:gradFill) - gradient stops with colors
  - Linear gradients (a:lin) with angle
  - Radial/path gradients (a:path)
  - Tile rectangle (a:tileRect)
- **Pattern Fill** (a:pattFill) - 50+ pattern presets
- **Blip Fill** (a:blipFill) - image fills with stretch/tile
- **Group Fill** (a:grpFill) - inherit parent fill
- **No Fill** (a:noFill) - transparent

### Lines
- **Line Properties** (a:ln) with:
  - Width (w) in EMUs
  - Compound type (cmpd) - single, double, thickThin, thinThick, triple
  - Pen alignment (algn) - center, inset
  - Cap type (cap) - flat, round, square
  - Dash style (prstDash) - solid, dash, dot, dashDot, etc.
  - Line joins - round, bevel, miter
  - Head/tail end types - arrow, diamond, oval, stealth, triangle

### Effects
- **Effect List** (a:effectLst) containing:
  - Outer Shadow (a:outerShdw) - blur, distance, direction, color
  - Inner Shadow (a:innerShdw)
  - Glow (a:glow) - radius, color
  - Soft Edge (a:softEdge) - radius
  - Reflection (a:reflection) - blur, direction, distance
  - Blur (a:blur)
  - Preset Shadow (a:prstShdw)
  - Fill Overlay (a:fillOverlay)

### Text
- **Text Body** (a:txBody) containing:
  - Body Properties (a:bodyPr) - wrap, anchor, rotation, insets, columns
  - List Style (a:lstStyle) - paragraph level defaults
  - Paragraphs (a:p)
- **Paragraph** (a:p) containing:
  - Paragraph Properties (a:pPr) - alignment, indentation, spacing, bullets
  - Runs (a:r), line breaks (a:br), fields (a:fld)
- **Run** (a:r) containing:
  - Run Properties (a:rPr) - font, size, bold, italic, underline, color
  - Text (a:t)

### Transforms
- **Transform 2D** (a:xfrm) with:
  - Offset (a:off) - x, y in EMUs
  - Extents (a:ext) - cx, cy in EMUs
  - Rotation (rot) - angle in 1/60000 degree
  - Flip horizontal/vertical (flipH, flipV)

### Blips (Images)
- **Blip** (a:blip) with:
  - Embedded reference (r:embed) - relationship ID to image part
  - Linked reference (r:link) - relationship ID to external image
  - Compression state (cstate) - none, print, screen, email
  - Effects - alpha modulation, tint, grayscale

### Charts
- **Chart Space** (c:chartSpace) - root element containing:
  - Chart (c:chart) - main chart definition
  - Print settings, external data, style, colors
- **Chart Types**:
  - Bar/Column charts (c:barChart, c:bar3DChart)
  - Line charts (c:lineChart, c:line3DChart)
  - Pie charts (c:pieChart, c:pie3DChart, c:doughnutChart)
  - Area charts (c:areaChart, c:area3DChart)
  - Scatter charts (c:scatterChart, c:bubbleChart)
  - Radar charts (c:radarChart)
  - Surface charts (c:surfaceChart, c:surface3DChart)
  - Stock charts (c:stockChart)
- **Axis Types**: Category, Value, Date, Series
- **Common Elements**: Legend, Title, Data Labels, Plot Area

### Colors
- **RGB Color** (a:srgbClr) - 6-digit hex (RRGGBB)
- **Scheme Color** (a:schemeClr) - theme reference:
  - dk1, lt1, dk2, lt2 (text/background)
  - accent1-6 (theme accents)
  - hlink, folHlink (hyperlinks)
- **System Color** (a:sysClr) - windowText, window, etc.
- **HSL Color** (a:hslClr) - hue, saturation, luminance
- **Preset Color** (a:prstClr) - 150+ named colors
- **Color Transforms**:
  - Alpha (a:alpha) - opacity
  - Tint (a:tint), Shade (a:shade)
  - Saturation modulation (a:satMod)
  - Luminance modulation (a:lumMod)
  - Hue offset (a:hueOff)

## Namespaces

| Prefix | URI | Description |
|--------|-----|-------------|
| `a` | `http://schemas.openxmlformats.org/drawingml/2006/main` | Core DrawingML |
| `c` | `http://schemas.openxmlformats.org/drawingml/2006/chart` | Charts |
| `xdr` | `http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing` | Spreadsheet drawings |
| `wp` | `http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing` | Word drawings |
| `pic` | `http://schemas.openxmlformats.org/drawingml/2006/picture` | Pictures |
| `dgm` | `http://schemas.openxmlformats.org/drawingml/2006/diagram` | Diagrams |
| `r` | `http://schemas.openxmlformats.org/officeDocument/2006/relationships` | Relationships |

## Validation

All 19+ validation constraint types apply to DrawingML elements:

1. **AttributeValueRangeConstraint** - min/max bounds (e.g., rotation 0-21599999)
2. **AttributeValuePatternConstraint** - regex validation (e.g., hex colors)
3. **AttributeValueSetConstraint** - enumeration values (e.g., preset shape names)
4. **ParentTypeConstraint** - allowed parent elements
5. **ChildElementConstraint** - required/allowed children
6. **UniqueValueConstraint** - uniqueness within scope
7. **RelationshipExistConstraint** - r:embed/r:link must exist
8. **RelationshipTypeConstraint** - relationship type validation
9. **ReferenceExistConstraint** - ID reference validation
10. **IndexRangeConstraint** - valid index bounds
11. **RootAttributeConstraint** - required root attributes
12. **AttributeAbsentConstraint** - mutually exclusive attributes
13. **AttributeCannotOmitConstraint** - conditionally required
14. **AttributeValueLengthConstraint** - string length limits
15. **UniqueAttributeValueConstraint** - unique across document
16. **PartContainerConstraint** - part must be in container
17. **DataPartConstraint** - data part validation
18. **PartTypeConstraint** - part content type validation
19. **Custom constraints** - extensible constraint system

## Unit Conversions

DrawingML uses English Metric Units (EMUs) for precise positioning:

| Unit | EMUs | Notes |
|------|------|-------|
| 1 inch | 914,400 | Base reference |
| 1 cm | 360,000 | Metric |
| 1 point | 12,700 | Typography |
| 1 pixel (96 DPI) | 9,525 | Screen |
| 1 degree rotation | 60,000 | Rotation in rot attribute |
| Full rotation | 21,600,000 | 360 degrees |

## Confirmed Design Decisions (ULTRATHINK Approved)

These decisions apply consistently across all goffice SDKs (Word, Presentation, Spreadsheet, DrawingML):

### Core Architecture

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Module Path** | `github.com/connerohnesorge/goffice` | Single unified module, simple imports, atomic versioning |
| **Go Version** | 1.25+ | Range-over-func for iterators, modern generics features |
| **Module Layout** | Single unified go.mod | Shared types, atomic versioning, simple dependency graph |
| **Code Generation** | JSON Schema → Go | Parse Open-XML-SDK's JSON schema files, generate Go structs with XML tags |
| **Office Version** | 2016+ only (ECMA-376 5th edition+) | Modern documents, reduced complexity, covers 95%+ of real-world files |
| **Thread Safety** | Per-Document RWMutex | Single sync.RWMutex per OpenXmlPackage, simple and sufficient |
| **Dependencies** | Pure stdlib only | archive/zip, encoding/xml, sync - no external dependencies |

### Element System

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Element Metadata** | Embedded struct fields | XMLName, Namespace, LocalName as struct fields, self-contained, works with encoding/xml |
| **Construction Pattern** | Functional options | `NewShape(WithPresetGeom("rect"), WithSolidFill("#FF0000"))` - idiomatic Go, extensible, self-documenting |
| **Child Access** | Generic functions | `First[T](el)`, `All[T](el)`, `OfType[T](el)` using Go 1.18+ generics with range-over-func iterators |
| **API Naming** | Go-idiomatic short names | `Xfrm` not `Transform2D`, `SolidFill` not `SolidFillProperties` |

### Validation System

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Validation Strategy** | Code-generated Validate() methods | Generate type-specific Validate() during code gen, zero reflection, compile-time type safety |
| **Error Handling** | Structured errors | Custom error types (ValidationError, ParseError) with path, element, constraint info, errors.Is/As compatible |
| **Validation Mode** | Strict by default | Reject invalid content, return errors for malformed documents |

### Extensibility

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Feature Collection** | Interface Registry pattern | Define Feature interface, register by interface type, supports hierarchy (Element→Part→Package→Global) with fallback chain |
| **Shared Package** | `drawingml/` at top level | Used by all three SDKs (Word, Presentation, Spreadsheet) for shapes, images, effects, and charts |

### XML Processing

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **XML Prefixes** | Fixed canonical | Always use canonical prefixes: a: for main DrawingML, c: for charts, wp: for Word, xdr: for Spreadsheet |
| **MC Handling** | Parse-time processing | Process AlternateContent/Choice/Fallback during XML parsing based on target FileFormatVersion |
| **Namespace Management** | Embedded in types | Each generated type knows its namespace URI and local name |

### Package Structure

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Layout** | Subpackages by namespace | `drawingml/{main,chart,spreadsheet,word,picture}` |
| **Phase 1 Scope** | Full DrawingML support | Complete implementation of all DrawingML types for Word, Presentation, and Spreadsheet SDKs |

## Estimated Scope

Based on ECMA-376 schema analysis:

- **Core DrawingML** (a:): ~400 element types
- **Charts** (c:): ~150 element types
- **Spreadsheet Drawing** (xdr:): ~30 element types
- **Word Drawing** (wp:): ~20 element types
- **Picture** (pic:): ~10 element types
- **Enumerations**: ~200 enum types
- **Total**: ~600+ types, estimated 30,000+ lines of generated code
