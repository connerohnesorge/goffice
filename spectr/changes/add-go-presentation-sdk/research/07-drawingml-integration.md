# DrawingML Integration Research

## Overview

DrawingML (namespace prefix: `a`) is the shared graphics layer used by Word, Excel, and PowerPoint. PresentationML heavily depends on DrawingML for shapes, text, fills, and transforms.

## Namespace

```
URI: http://schemas.openxmlformats.org/drawingml/2006/main
Prefix: a
```

## Core DrawingML Elements in Presentations

### Transform2D (a:xfrm)

Defines position and size of shapes:

```xml
<a:xfrm rot="5400000" flipH="1" flipV="0">
  <a:off x="914400" y="914400"/>    <!-- Offset in EMUs -->
  <a:ext cx="1828800" cy="914400"/> <!-- Extents in EMUs -->
</a:xfrm>
```

| Attribute | Description |
|-----------|-------------|
| `rot` | Rotation in 1/60000 of a degree |
| `flipH` | Horizontal flip (0 or 1) |
| `flipV` | Vertical flip (0 or 1) |

| Child | Description |
|-------|-------------|
| `a:off` | Offset (x, y) in EMUs |
| `a:ext` | Extents (cx, cy) in EMUs |

EMU = English Metric Unit (914400 EMUs = 1 inch)

### Preset Geometry (a:prstGeom)

```xml
<a:prstGeom prst="rect">
  <a:avLst/>  <!-- Adjust values for shape parameters -->
</a:prstGeom>
```

Common preset shapes:
- `rect`, `ellipse`, `roundRect`, `triangle`
- `line`, `straightConnector1`
- `rightArrow`, `leftArrow`, `upArrow`, `downArrow`
- `star5`, `star6`, `heart`, `cloud`

### Custom Geometry (a:custGeom)

For complex paths not covered by presets:

```xml
<a:custGeom>
  <a:pathLst>
    <a:path w="21600" h="21600">
      <a:moveTo><a:pt x="0" y="0"/></a:moveTo>
      <a:lnTo><a:pt x="21600" y="0"/></a:lnTo>
      <a:close/>
    </a:path>
  </a:pathLst>
</a:custGeom>
```

## Fill Types

### Solid Fill (a:solidFill)

```xml
<a:solidFill>
  <a:srgbClr val="FF0000"/>  <!-- RGB color -->
</a:solidFill>
```

### Gradient Fill (a:gradFill)

```xml
<a:gradFill rotWithShape="1">
  <a:gsLst>
    <a:gs pos="0">
      <a:srgbClr val="FF0000"/>
    </a:gs>
    <a:gs pos="100000">
      <a:srgbClr val="0000FF"/>
    </a:gs>
  </a:gsLst>
  <a:lin ang="5400000"/>  <!-- Linear gradient -->
</a:gradFill>
```

### Blip Fill (a:blipFill)

For image fills:

```xml
<a:blipFill>
  <a:blip r:embed="rId1"/>
  <a:stretch>
    <a:fillRect/>
  </a:stretch>
</a:blipFill>
```

### Pattern Fill (a:pattFill)

```xml
<a:pattFill prst="pct50">
  <a:fgClr><a:srgbClr val="000000"/></a:fgClr>
  <a:bgClr><a:srgbClr val="FFFFFF"/></a:bgClr>
</a:pattFill>
```

### No Fill

```xml
<a:noFill/>
```

### Group Fill (a:grpFill)

Inherits fill from parent group shape.

## Text Body Structure

### TextBody (a:txBody)

```xml
<p:txBody>
  <a:bodyPr wrap="square" anchor="ctr"/>
  <a:lstStyle/>
  <a:p>
    <a:pPr algn="ctr"/>
    <a:r>
      <a:rPr lang="en-US" sz="1800" b="1" i="0"/>
      <a:t>Hello World</a:t>
    </a:r>
  </a:p>
</p:txBody>
```

### Body Properties (a:bodyPr)

| Attribute | Description |
|-----------|-------------|
| `wrap` | Text wrapping mode (none, square) |
| `anchor` | Vertical anchor (t, ctr, b) |
| `anchorCtr` | Center text horizontally |
| `rot` | Text rotation |
| `vert` | Vertical text mode |
| `lIns`, `tIns`, `rIns`, `bIns` | Inset margins |

### Paragraph (a:p)

Contains:
- `a:pPr` - Paragraph properties
- `a:r` - Text runs
- `a:br` - Line breaks
- `a:fld` - Fields

### Paragraph Properties (a:pPr)

| Attribute | Description |
|-----------|-------------|
| `algn` | Alignment (l, ctr, r, just) |
| `marL` | Left margin |
| `indent` | First line indent |
| `lvl` | Paragraph level (0-8) |

### Run (a:r)

Contains:
- `a:rPr` - Run properties
- `a:t` - Text content

### Run Properties (a:rPr)

| Attribute | Description |
|-----------|-------------|
| `lang` | Language code |
| `sz` | Font size in hundredths of a point |
| `b` | Bold (0 or 1) |
| `i` | Italic (0 or 1) |
| `u` | Underline style |
| `strike` | Strikethrough |

Child elements for colors, fonts, effects.

## Color Types

### sRGB Color

```xml
<a:srgbClr val="FF0000">
  <a:alpha val="50000"/>  <!-- 50% opacity -->
</a:srgbClr>
```

### Scheme Color

```xml
<a:schemeClr val="accent1">
  <a:shade val="75000"/>
</a:schemeClr>
```

Scheme values: dk1, lt1, dk2, lt2, accent1-6, hlink, folHlink

### System Color

```xml
<a:sysClr val="windowText"/>
```

## PresentationML Integration

### Shape (p:sp) uses DrawingML

```xml
<p:sp>
  <p:nvSpPr>...</p:nvSpPr>
  <p:spPr>
    <a:xfrm>...</a:xfrm>        <!-- Position/size -->
    <a:prstGeom prst="rect"/>   <!-- Geometry -->
    <a:solidFill>...</a:solidFill>  <!-- Fill -->
    <a:ln>...</a:ln>            <!-- Outline -->
  </p:spPr>
  <p:txBody>...</p:txBody>      <!-- Text (uses a: namespace) -->
</p:sp>
```

### Picture (p:pic) uses DrawingML

```xml
<p:pic>
  <p:nvPicPr>...</p:nvPicPr>
  <p:blipFill>
    <a:blip r:embed="rId1"/>
    <a:stretch>...</a:stretch>
  </p:blipFill>
  <p:spPr>
    <a:xfrm>...</a:xfrm>
  </p:spPr>
</p:pic>
```

## Go Implementation Considerations

1. **Separate DrawingML package** shared across document types
2. **EMU conversion utilities** (EMU ↔ inches, points, cm)
3. **Preset geometry constants** for all standard shapes
4. **Color type hierarchy** with scheme, RGB, system variants
5. **Text body parsing** with paragraph/run/text structure
6. **Fill type interface** with solid, gradient, blip, pattern implementations
