# Font Requirements

This document describes font handling in goffice-pdf and system font requirements.

## Overview

goffice-pdf requires access to fonts to accurately render text in PDF documents. The renderer discovers and uses fonts from multiple sources in priority order:

1. **Embedded fonts** - Fonts embedded in the source Office document
2. **System fonts** - Fonts installed on the operating system
3. **Fallback fonts** - Built-in fallback chain for common fonts

## System Font Discovery

goffice-pdf automatically discovers fonts from OS-specific locations:

### Windows

```
C:\Windows\Fonts\
C:\Users\<username>\AppData\Local\Microsoft\Windows\Fonts\
```

### macOS

```
/System/Library/Fonts/
/Library/Fonts/
~/Library/Fonts/
```

### Linux

```
/usr/share/fonts/
/usr/local/share/fonts/
~/.fonts/
~/.local/share/fonts/
```

## Recommended Fonts

For best compatibility with Microsoft Office documents, install these font families:

### Core Fonts (Required)

These fonts are commonly used in Office documents:

- **Arial** - Sans-serif, used by default in many documents
- **Times New Roman** - Serif, common in academic documents
- **Calibri** - Default font in Word 2007+
- **Cambria** - Serif font in Word 2007+
- **Courier New** - Monospace font
- **Comic Sans MS** - Informal font
- **Georgia** - Web-optimized serif
- **Verdana** - Web-optimized sans-serif

### Office Fonts (Recommended)

Additional fonts included with Microsoft Office:

- **Calibri Light** - Light weight variant
- **Cambria Math** - Math symbols
- **Candara** - Humanist sans-serif
- **Consolas** - Monospace for code
- **Constantia** - Serif
- **Corbel** - Sans-serif

### CJK Fonts (For Asian Language Support)

- **Microsoft YaHei** - Simplified Chinese
- **SimSun** - Simplified Chinese
- **MS Gothic** - Japanese
- **Malgun Gothic** - Korean
- **Noto Sans CJK** - Open-source alternative

### Symbol and Special Fonts

- **Symbol** - Mathematical symbols
- **Wingdings** - Icon/symbol font
- **Webdings** - Additional symbols

## Installing Fonts

### Windows

1. Download font files (.ttf or .otf)
2. Right-click font file → Install
3. Or copy to `C:\Windows\Fonts\`

### macOS

1. Double-click font file
2. Click "Install Font"
3. Or copy to `/Library/Fonts/` or `~/Library/Fonts/`

### Linux (Debian/Ubuntu)

```bash
# Install Microsoft core fonts
sudo apt-get install ttf-mscorefonts-installer

# Or manually install fonts
sudo mkdir -p /usr/local/share/fonts/truetype/custom
sudo cp *.ttf /usr/local/share/fonts/truetype/custom/
sudo fc-cache -f -v
```

### Linux (RHEL/CentOS/Fedora)

```bash
# Install liberation fonts (metric-compatible with MS fonts)
sudo dnf install liberation-fonts

# Or manually install
sudo mkdir -p /usr/share/fonts/custom
sudo cp *.ttf /usr/share/fonts/custom/
sudo fc-cache -f -v
```

## Font Substitution

When a required font is not available, goffice-pdf uses a fallback chain:

```
Requested Font          →  Fallback
----------------------     -------------------------
Calibri                 →  Arial → Liberation Sans
Times New Roman         →  Liberation Serif
Courier New             →  Liberation Mono
Arial                   →  Liberation Sans → Helvetica
Cambria                 →  Times New Roman → Liberation Serif
Consolas                →  Courier New → Liberation Mono
```

### Configuring Custom Fallbacks

You can configure custom font fallback chains:

```go
opts := pdf.DefaultRenderOptions()
opts.FontFallbacks = map[string][]string{
    "Calibri": {"Arial", "Liberation Sans", "FreeSans"},
    "CustomFont": {"Arial", "Helvetica"},
}

// Or using the fluent API:
opts = pdf.DefaultRenderOptions().WithFontFallbacks(map[string][]string{
    "Calibri": {"Arial", "Liberation Sans", "FreeSans"},
    "CustomFont": {"Arial", "Helvetica"},
})
```

The fallback rules are processed in order. If "Calibri" is requested but not available, the renderer will try "Arial" first, then "Liberation Sans", and finally "FreeSans". Custom fallback rules take priority over the built-in defaults.

## Font Embedding Modes

### EmbedSubset (Default)

Embeds only the glyphs actually used in the document.

**Advantages:**
- Smallest file size
- Fast rendering
- Legal for most fonts

**Disadvantages:**
- PDF cannot be edited with new text
- Character coverage limited to used glyphs

**Use when:**
- Final output (no further editing)
- File size is important
- Distribution to others

### EmbedFull

Embeds complete font files.

**Advantages:**
- PDF can be edited with any characters
- Complete font features available

**Disadvantages:**
- Larger file size (2-10 MB per font)
- Slower rendering
- May violate font license

**Use when:**
- PDF will be edited later
- Need all font features
- Internal use only

### NoEmbed

References fonts by name without embedding.

**Advantages:**
- Smallest possible file size
- Fastest rendering

**Disadvantages:**
- Requires fonts on viewing system
- May not render correctly everywhere
- Not recommended for distribution

**Use when:**
- Controlled environment (known font installation)
- Temporary/preview output
- File size is critical

## Font Licensing

Be aware of font licensing when embedding fonts:

### Embeddable Fonts

These fonts allow embedding:
- **Liberation fonts** - GPL with exception
- **Noto fonts** - SIL Open Font License
- **DejaVu fonts** - Free license
- **GNU FreeFont** - GPL with exception

### Microsoft Fonts

Microsoft fonts typically allow:
- ✅ Print and preview embedding
- ✅ Subset embedding
- ❌ Full embedding for editing (check license)
- ❌ Redistribution of font files

Always check the specific font license before embedding.

## Font Metrics

goffice-pdf extracts metrics from TrueType and OpenType fonts:

- **Glyph widths** - For text layout
- **Kerning pairs** - For character spacing
- **Ascender/descender** - For line height
- **Cap height/x-height** - For vertical alignment
- **Bounding boxes** - For precise positioning

## Troubleshooting

### Font Not Found Error

```
Error: pdf: required font not found: Calibri
```

**Solution:**
1. Install the missing font
2. Configure a fallback font
3. Replace font in source document

### Incorrect Text Layout

**Symptoms:** Text overlaps, spacing is wrong, different line breaks

**Causes:**
- Different font metrics (substitute font)
- Missing kerning data
- Font version differences

**Solutions:**
1. Install exact font version
2. Use EmbedFull mode
3. Adjust fallback chain

### Missing Characters (Tofu)

**Symptoms:** � or □ characters appear

**Causes:**
- Font doesn't contain the required glyphs
- Character not in embedded subset

**Solutions:**
1. Use EmbedFull mode
2. Install fonts with broader Unicode coverage
3. Configure fallback fonts for specific scripts

### Slow Rendering

**Symptoms:** Font loading takes a long time

**Causes:**
- Large font files
- Many fonts being loaded
- Font cache not populated

**Solutions:**
1. Use EmbedSubset mode (default)
2. Limit font variety in documents
3. Pre-warm font cache

## Font Cache

goffice-pdf caches parsed fonts in memory:

**Cache Size:** LRU cache with configurable limit
**Default Size:** 50 fonts
**Cache Key:** Font family + style + weight

Clear cache between renders if memory is constrained.

## Testing Font Availability

To check which fonts are available:

```go
fonts := pdf.DiscoverFonts()
for _, font := range fonts {
    fmt.Printf("%s (%s)\n", font.Family, font.Style)
}
```

## Best Practices

1. **Document Creation**
   - Use standard fonts (Arial, Times New Roman, Calibri)
   - Avoid obscure or decorative fonts
   - Test with target font set

2. **Font Installation**
   - Install Microsoft core fonts package
   - Include CJK fonts if needed
   - Keep fonts updated

3. **Font Embedding**
   - Use EmbedSubset for distribution
   - Use EmbedFull for editable PDFs
   - Never use NoEmbed for distribution

4. **Cross-Platform**
   - Test on all target platforms
   - Use metric-compatible fonts
   - Configure appropriate fallbacks

5. **Performance**
   - Limit font variety
   - Use font caching
   - Subset fonts when possible

## Resources

- **Microsoft Typography:** https://docs.microsoft.com/en-us/typography/
- **Google Fonts:** https://fonts.google.com/
- **Liberation Fonts:** https://github.com/liberationfonts/liberation-fonts
- **Noto Fonts:** https://fonts.google.com/noto
- **Font Licensing Guide:** https://fonts.google.com/knowledge/glossary/licensing
