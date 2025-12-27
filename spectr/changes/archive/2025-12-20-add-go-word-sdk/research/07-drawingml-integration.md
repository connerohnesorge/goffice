# DrawingML Integration in WordprocessingML Research

## Overview

DrawingML (namespace prefix: `a`) is the shared graphics layer used by Word, Excel, and PowerPoint. WordprocessingML uses DrawingML for inline and floating images, shapes, charts, and diagrams, but with a different integration pattern than PresentationML.

## Namespace

```
DrawingML Main: http://schemas.openxmlformats.org/drawingml/2006/main (prefix: a)
WordprocessingML Drawing: http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing (prefix: wp)
Word Processing Shape: http://schemas.microsoft.com/office/word/2010/wordprocessingShape (prefix: wps)
Word Processing Group: http://schemas.microsoft.com/office/word/2010/wordprocessingGroup (prefix: wpg)
Word Processing Canvas: http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas (prefix: wpc)
```

## Drawing Integration in Word

### Two Drawing Types

1. **Inline Drawings** (`wp:inline`) - Flow with text, no wrapping
2. **Anchor Drawings** (`wp:anchor`) - Positioned relative to page/paragraph, support text wrapping

### Inline Drawing Structure

```xml
<w:drawing>
  <wp:inline distT="0" distB="0" distL="0" distR="0">
    <wp:extent cx="914400" cy="914400"/>    <!-- Size in EMUs -->
    <wp:effectExtent l="0" t="0" r="0" b="0"/>
    <wp:docPr id="1" name="Picture 1"/>     <!-- Document properties -->
    <wp:cNvGraphicFramePr>
      <a:graphicFrameLocks noChangeAspect="1"/>
    </wp:cNvGraphicFramePr>
    <a:graphic>
      <a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/picture">
        <pic:pic><!-- Picture content --></pic:pic>
      </a:graphicData>
    </a:graphic>
  </wp:inline>
</w:drawing>
```

### Anchor Drawing Structure

```xml
<w:drawing>
  <wp:anchor distT="0" distB="0" distL="114300" distR="114300"
             simplePos="0" relativeHeight="251658240"
             behindDoc="0" locked="0" layoutInCell="1"
             allowOverlap="1">
    <wp:simplePos x="0" y="0"/>
    <wp:positionH relativeFrom="column">
      <wp:posOffset>0</wp:posOffset>
    </wp:positionH>
    <wp:positionV relativeFrom="paragraph">
      <wp:posOffset>0</wp:posOffset>
    </wp:positionV>
    <wp:extent cx="1828800" cy="914400"/>
    <wp:effectExtent l="0" t="0" r="0" b="0"/>
    <wp:wrapSquare wrapText="bothSides"/>  <!-- Text wrapping -->
    <wp:docPr id="1" name="Shape 1"/>
    <a:graphic>
      <!-- Graphic content -->
    </a:graphic>
  </wp:anchor>
</w:drawing>
```

## Text Wrapping Types

| Element | Description |
|---------|-------------|
| `wp:wrapNone` | No text wrapping (floats over text) |
| `wp:wrapSquare` | Square wrapping around bounding box |
| `wp:wrapTight` | Tight wrapping following contours |
| `wp:wrapThrough` | Text flows through transparent areas |
| `wp:wrapTopAndBottom` | Text above and below only |

### wrapSquare Attributes

```xml
<wp:wrapSquare wrapText="bothSides" distT="0" distB="0" distL="114300" distR="114300"/>
```

| Attribute | Values |
|-----------|--------|
| `wrapText` | `bothSides`, `left`, `right`, `largest` |
| `distT/B/L/R` | Distance from text in EMUs |

## Position Reference Points

### Horizontal Position (relativeFrom)

| Value | Description |
|-------|-------------|
| `margin` | Relative to page margins |
| `page` | Relative to page edge |
| `column` | Relative to column |
| `character` | Relative to character position |
| `leftMargin` | Relative to left margin |
| `rightMargin` | Relative to right margin |
| `insideMargin` | Relative to inside margin (for facing pages) |
| `outsideMargin` | Relative to outside margin |

### Vertical Position (relativeFrom)

| Value | Description |
|-------|-------------|
| `margin` | Relative to page margins |
| `page` | Relative to page edge |
| `paragraph` | Relative to paragraph |
| `line` | Relative to line |
| `topMargin` | Relative to top margin |
| `bottomMargin` | Relative to bottom margin |
| `insideMargin` | Relative to inside margin |
| `outsideMargin` | Relative to outside margin |

## Picture Element (pic:pic)

```xml
<pic:pic xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture">
  <pic:nvPicPr>
    <pic:cNvPr id="0" name="image.png"/>
    <pic:cNvPicPr/>
  </pic:nvPicPr>
  <pic:blipFill>
    <a:blip r:embed="rId1">
      <a:extLst>
        <a:ext uri="{28A0092B-C50C-407E-A947-70E740481C1C}">
          <a14:useLocalDpi val="0"/>
        </a:ext>
      </a:extLst>
    </a:blip>
    <a:stretch>
      <a:fillRect/>
    </a:stretch>
  </pic:blipFill>
  <pic:spPr>
    <a:xfrm>
      <a:off x="0" y="0"/>
      <a:ext cx="914400" cy="914400"/>
    </a:xfrm>
    <a:prstGeom prst="rect">
      <a:avLst/>
    </a:prstGeom>
  </pic:spPr>
</pic:pic>
```

## Word Processing Shapes (wps:wsp)

```xml
<wps:wsp>
  <wps:cNvSpPr/>
  <wps:spPr>
    <a:xfrm>
      <a:off x="0" y="0"/>
      <a:ext cx="914400" cy="457200"/>
    </a:xfrm>
    <a:prstGeom prst="rect">
      <a:avLst/>
    </a:prstGeom>
    <a:solidFill>
      <a:srgbClr val="4F81BD"/>
    </a:solidFill>
    <a:ln w="9525">
      <a:solidFill>
        <a:srgbClr val="000000"/>
      </a:solidFill>
    </a:ln>
  </wps:spPr>
  <wps:txbx>
    <w:txbxContent>
      <w:p>
        <w:r><w:t>Text in shape</w:t></w:r>
      </w:p>
    </w:txbxContent>
  </wps:txbx>
  <wps:bodyPr rot="0" vert="horz" wrap="square" anchor="t"/>
</wps:wsp>
```

## Text Box Content

Word shapes can contain rich text via `wps:txbx`:

```xml
<wps:txbx>
  <w:txbxContent>
    <!-- Full WordprocessingML content -->
    <w:p>
      <w:pPr>...</w:pPr>
      <w:r>
        <w:rPr>...</w:rPr>
        <w:t>Formatted text</w:t>
      </w:r>
    </w:p>
  </w:txbxContent>
</wps:txbx>
```

## EMU (English Metric Units) Conversions

| Unit | EMUs |
|------|------|
| 1 inch | 914400 |
| 1 cm | 360000 |
| 1 point | 12700 |
| 1 pixel (96 DPI) | 9525 |
| 1 twip | 635 |

## Go Implementation Considerations

1. **Separate WordprocessingDrawing package** (`wp` namespace)
2. **Support both inline and anchor drawings** with different positioning models
3. **EMU conversion utilities** (EMU <-> inches, points, cm, twips)
4. **Text wrapping implementations** for all wrapping types
5. **Position reference handling** for horizontal and vertical positioning
6. **Picture element support** with blip references to image parts
7. **Shape support** via `wps:wsp` with text box content
8. **Reuse DrawingML core** (`a` namespace) from shared package
