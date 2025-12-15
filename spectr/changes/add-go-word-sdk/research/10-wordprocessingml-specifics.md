# WordprocessingML-Specific Research

## Overview

This document covers WordprocessingML-specific architecture and elements that differ from PresentationML. Understanding these specifics is essential for implementing the Word SDK.

## Namespace

```
URI: http://schemas.openxmlformats.org/wordprocessingml/2006/main
Prefix: w
```

## Document Structure

### Root Element (w:document)

```xml
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:body>
    <!-- Document content -->
  </w:body>
</w:document>
```

### Body Element (w:body)

Contains all block-level content:

```xml
<w:body>
  <w:p><!-- Paragraph --></w:p>
  <w:tbl><!-- Table --></w:tbl>
  <w:sdt><!-- Structured Document Tag --></w:sdt>
  <w:sectPr><!-- Section Properties (last) --></w:sectPr>
</w:body>
```

## Block-Level Elements

### Paragraph (w:p)

```xml
<w:p>
  <w:pPr><!-- Paragraph Properties --></w:pPr>
  <w:r><!-- Run --></w:r>
  <w:hyperlink><!-- Hyperlink --></w:hyperlink>
  <w:bookmarkStart/>
  <w:bookmarkEnd/>
</w:p>
```

### Paragraph Properties (w:pPr)

```xml
<w:pPr>
  <w:pStyle w:val="Heading1"/>     <!-- Style reference -->
  <w:jc w:val="center"/>           <!-- Justification -->
  <w:spacing w:before="240" w:after="120" w:line="276" w:lineRule="auto"/>
  <w:ind w:left="720" w:right="0" w:firstLine="360"/>  <!-- Indentation -->
  <w:numPr>                        <!-- Numbering -->
    <w:ilvl w:val="0"/>
    <w:numId w:val="1"/>
  </w:numPr>
  <w:outlineLvl w:val="0"/>        <!-- Outline level -->
  <w:rPr><!-- Default run properties --></w:rPr>
</w:pPr>
```

### Justification Values (w:jc)

| Value | Description |
|-------|-------------|
| `left` | Left aligned |
| `center` | Center aligned |
| `right` | Right aligned |
| `both` | Justified |
| `distribute` | Distributed |

## Inline Elements

### Run (w:r)

```xml
<w:r>
  <w:rPr><!-- Run Properties --></w:rPr>
  <w:t>Text content</w:t>
  <w:br/>                          <!-- Line break -->
  <w:tab/>                         <!-- Tab character -->
  <w:sym w:font="Symbol" w:char="F0AE"/>  <!-- Symbol -->
  <w:drawing><!-- Inline/anchor drawing --></w:drawing>
</w:r>
```

### Run Properties (w:rPr)

```xml
<w:rPr>
  <w:rStyle w:val="Strong"/>       <!-- Character style -->
  <w:rFonts w:ascii="Arial" w:hAnsi="Arial"/>
  <w:b/>                           <!-- Bold -->
  <w:i/>                           <!-- Italic -->
  <w:u w:val="single"/>            <!-- Underline -->
  <w:strike/>                      <!-- Strikethrough -->
  <w:sz w:val="24"/>               <!-- Font size (half-points) -->
  <w:color w:val="FF0000"/>        <!-- Text color -->
  <w:highlight w:val="yellow"/>    <!-- Highlight color -->
  <w:vertAlign w:val="superscript"/> <!-- Vertical alignment -->
</w:rPr>
```

### Text (w:t)

```xml
<w:t xml:space="preserve">Text with spaces</w:t>
```

The `xml:space="preserve"` attribute preserves leading/trailing whitespace.

## Tables

### Table Structure (w:tbl)

```xml
<w:tbl>
  <w:tblPr><!-- Table Properties --></w:tblPr>
  <w:tblGrid>
    <w:gridCol w:w="2880"/>
    <w:gridCol w:w="2880"/>
  </w:tblGrid>
  <w:tr><!-- Table Row --></w:tr>
</w:tbl>
```

### Table Properties (w:tblPr)

```xml
<w:tblPr>
  <w:tblStyle w:val="TableGrid"/>
  <w:tblW w:w="5000" w:type="pct"/>  <!-- Width: 50% -->
  <w:jc w:val="center"/>              <!-- Table alignment -->
  <w:tblBorders>
    <w:top w:val="single" w:sz="4" w:space="0" w:color="000000"/>
    <w:left w:val="single" w:sz="4" w:space="0" w:color="000000"/>
    <w:bottom w:val="single" w:sz="4" w:space="0" w:color="000000"/>
    <w:right w:val="single" w:sz="4" w:space="0" w:color="000000"/>
    <w:insideH w:val="single" w:sz="4" w:space="0" w:color="000000"/>
    <w:insideV w:val="single" w:sz="4" w:space="0" w:color="000000"/>
  </w:tblBorders>
  <w:tblCellMar>
    <w:top w:w="0" w:type="dxa"/>
    <w:left w:w="108" w:type="dxa"/>
    <w:bottom w:w="0" w:type="dxa"/>
    <w:right w:w="108" w:type="dxa"/>
  </w:tblCellMar>
</w:tblPr>
```

### Table Row (w:tr)

```xml
<w:tr>
  <w:trPr>
    <w:trHeight w:val="400" w:hRule="atLeast"/>
    <w:jc w:val="center"/>
  </w:trPr>
  <w:tc><!-- Table Cell --></w:tc>
</w:tr>
```

### Table Cell (w:tc)

```xml
<w:tc>
  <w:tcPr>
    <w:tcW w:w="2880" w:type="dxa"/>
    <w:gridSpan w:val="2"/>           <!-- Column span -->
    <w:vMerge w:val="restart"/>       <!-- Row span start -->
    <w:vMerge/>                       <!-- Merged cell continuation -->
    <w:shd w:val="clear" w:color="auto" w:fill="FFFF00"/>
    <w:vAlign w:val="center"/>
  </w:tcPr>
  <w:p><!-- Cell must contain at least one paragraph --></w:p>
</w:tc>
```

## Sections

### Section Properties (w:sectPr)

```xml
<w:sectPr>
  <w:pgSz w:w="12240" w:h="15840" w:orient="portrait"/>
  <w:pgMar w:top="1440" w:right="1440" w:bottom="1440" w:left="1440"
           w:header="720" w:footer="720" w:gutter="0"/>
  <w:cols w:space="720" w:num="1"/>
  <w:headerReference w:type="default" r:id="rId1"/>
  <w:footerReference w:type="default" r:id="rId2"/>
  <w:type w:val="nextPage"/>          <!-- Section break type -->
</w:sectPr>
```

### Section Break Types

| Value | Description |
|-------|-------------|
| `nextPage` | Next page break |
| `continuous` | Continuous (no page break) |
| `evenPage` | Even page break |
| `oddPage` | Odd page break |

### Header/Footer Reference Types

| Value | Description |
|-------|-------------|
| `default` | Default header/footer |
| `first` | First page header/footer |
| `even` | Even page header/footer |

## Styles

### Style Definition (w:style)

```xml
<w:style w:type="paragraph" w:styleId="Heading1">
  <w:name w:val="heading 1"/>
  <w:basedOn w:val="Normal"/>
  <w:next w:val="Normal"/>
  <w:qFormat/>                        <!-- Quick format -->
  <w:uiPriority w:val="9"/>
  <w:pPr><!-- Paragraph properties --></w:pPr>
  <w:rPr><!-- Run properties --></w:rPr>
</w:style>
```

### Style Types

| Value | Description |
|-------|-------------|
| `paragraph` | Paragraph style |
| `character` | Character (run) style |
| `table` | Table style |
| `numbering` | Numbering style |

## Numbering

### Abstract Numbering (w:abstractNum)

```xml
<w:abstractNum w:abstractNumId="0">
  <w:lvl w:ilvl="0">
    <w:start w:val="1"/>
    <w:numFmt w:val="decimal"/>
    <w:lvlText w:val="%1."/>
    <w:lvlJc w:val="left"/>
    <w:pPr>
      <w:ind w:left="720" w:hanging="360"/>
    </w:pPr>
  </w:lvl>
</w:abstractNum>
```

### Number Formats (w:numFmt)

| Value | Description |
|-------|-------------|
| `decimal` | 1, 2, 3 |
| `upperRoman` | I, II, III |
| `lowerRoman` | i, ii, iii |
| `upperLetter` | A, B, C |
| `lowerLetter` | a, b, c |
| `bullet` | Bullet character |

## Measurement Units

| Unit | Description | Conversion |
|------|-------------|------------|
| Twips (dxa) | Twentieths of a point | 1440 twips = 1 inch |
| Half-points | Font size units | 24 half-points = 12pt |
| EMU | English Metric Units | 914400 EMU = 1 inch |
| Percent (pct) | Percentage * 50 | 5000 pct = 100% |

## Go Implementation Considerations

1. **Document structure**: Body -> Paragraphs/Tables -> Runs -> Text
2. **Property inheritance**: Document defaults -> Styles -> Direct formatting
3. **Table complexity**: Grid-based with merge support
4. **Section handling**: Properties apply to preceding content
5. **Style resolution**: basedOn chains, type-specific inheritance
6. **Numbering system**: Abstract -> Concrete -> Paragraph reference
7. **Unit conversions**: Support twips, EMUs, percentages
8. **Whitespace preservation**: Honor xml:space attribute
