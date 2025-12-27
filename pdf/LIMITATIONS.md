# Limitations and Known Issues

This document describes current limitations, known issues, and future improvements for goffice-pdf.

## General Limitations

### PDF Format Limitations

These are inherent limitations of the PDF format itself:

- **No interactivity** - PDF is a static format
  - Form fields become static text
  - Hyperlinks work, but no scripting
  - No animation or video playback
  - No live data connections

- **No recalculation** - Formulas are not live
  - Excel formulas show last calculated values
  - Cannot recalculate on opening
  - Conditional formatting is static

- **No editing** - PDF is primarily for viewing
  - Text editing limited (depends on font embedding)
  - Cannot restructure content
  - Cannot add/remove pages easily

## Word Document Limitations

### Not Supported

These Word features are **not supported** and will be omitted or simplified:

#### Content Features

- ❌ **Track changes** - Rendered in accepted state
- ❌ **Comments** - Not rendered
- ❌ **Form fields** - Rendered as static text
- ❌ **Macros/VBA** - Not executed or included
- ❌ **Custom XML** - Ignored
- ❌ **ActiveX controls** - Not rendered
- ❌ **Embedded objects (OLE)** - Placeholder only
- ❌ **Content controls** - Rendered as plain text
- ❌ **Bibliography** - Rendered as text, not updateable

#### Advanced Typography

- ❌ **OpenType features** - Limited support
  - Ligatures may not render correctly
  - Contextual alternates not supported
  - Stylistic sets not supported
- ❌ **Complex scripts** - Partial support
  - Arabic shaping: Basic support
  - Indic scripts: Limited
  - Right-to-left: Supported via Unicode Bidi
- ❌ **Font fallback for emoji** - May show as □

#### Advanced Layout

- ❌ **Text boxes with advanced wrapping** - Simplified
- ❌ **Shapes with complex text flow** - Approximated
- ❌ **Absolutely positioned frames** - May shift slightly
- ❌ **Multi-column text balancing** - May differ from Word

### Partially Supported

These features are supported but may have minor differences:

#### Tables

- ⚠️ **Auto-fit table layout** - May differ by 1-2 points
- ⚠️ **Nested tables** - Supported but complex nesting may differ
- ⚠️ **Table formulas** - Rendered as values, not live
- ⚠️ **Diagonal borders** - Not supported (omitted)

#### Lists

- ⚠️ **Custom bullet characters** - May use fallback
- ⚠️ **Picture bullets** - Rendered if available
- ⚠️ **Restart numbering** - Supported but edge cases may differ

#### Page Layout

- ⚠️ **Widow/orphan control** - Supported, minor differences possible
- ⚠️ **Keep with next** - Supported, complex cases may differ
- ⚠️ **Text wrapping around shapes** - Approximation
- ⚠️ **Mirror margins** - Supported for odd/even pages

#### Fields

- ⚠️ **PAGE field** - Supported
- ⚠️ **NUMPAGES field** - Supported
- ⚠️ **DATE field** - Rendered as current date
- ⚠️ **Complex fields** - May not update correctly
- ❌ **IF fields** - Evaluated at render time only
- ❌ **DOCPROPERTY fields** - Static values only

## Excel Spreadsheet Limitations

### Not Supported

#### Interactive Features

- ❌ **Formulas** - Shown as values (no recalculation)
- ❌ **Data validation** - Not enforced
- ❌ **Slicers** - Rendered in current state (not interactive)
- ❌ **PivotTables** - Rendered as static tables
- ❌ **External data connections** - Not refreshed
- ❌ **Macros/VBA** - Not included

#### Advanced Features

- ❌ **Sparklines** - Rendered as small charts (basic)
- ❌ **Power Query** - Not executed
- ❌ **Power Pivot** - Data models not included
- ❌ **Cube formulas** - Shown as values
- ❌ **Array formulas** - Shown as values
- ❌ **Dynamic arrays** - Shown as values

#### Charts

- ❌ **3D chart rotation** - Rendered at default angle
- ❌ **Chart animations** - Not rendered (Excel doesn't have these)
- ❌ **Interactive chart elements** - Static
- ⚠️ **Combination charts** - Supported but complex ones may differ
- ⚠️ **Waterfall charts** - Limited support
- ⚠️ **Funnel charts** - Limited support
- ⚠️ **Stock charts** - Basic support

### Partially Supported

- ⚠️ **Conditional formatting** - Most rules supported
  - Data bars: Supported
  - Color scales: Supported
  - Icon sets: Supported
  - Formula-based: Evaluated at render time

- ⚠️ **Number formatting** - Standard formats supported
  - Custom formats: Mostly supported
  - Complex format strings: May differ slightly
  - Locale-specific: Uses render-time locale

- ⚠️ **Print settings**
  - Fit to page: Supported
  - Page breaks: Supported
  - Print titles: Supported
  - Print area: Supported
  - Header/footer: Supported

## PowerPoint Presentation Limitations

### Not Supported

#### Animation and Transitions

- ❌ **Animations** - Slides rendered in final state
- ❌ **Transitions** - No transition effects
- ❌ **Slide timings** - Not applicable
- ❌ **Action buttons** - Not interactive
- ❌ **Hyperlinks to slides** - May not work correctly

#### Media

- ❌ **Videos** - Placeholder shown or omitted
- ❌ **Audio** - Not embedded
- ❌ **Animated GIFs** - Shown as first frame
- ❌ **Flash content** - Not rendered

#### Advanced Features

- ❌ **Slide zoom** - Not interactive
- ❌ **Morph transition** - Not rendered
- ❌ **3D models** - Rendered as placeholder or 2D
- ❌ **Embedded fonts** - May use system fonts instead

### Partially Supported

- ⚠️ **SmartArt** - Decomposed to shapes
  - Layout matches PowerPoint
  - Complex animations not rendered
  - 3D effects simplified

- ⚠️ **Master slides** - Fully supported
  - Inheritance works correctly
  - Placeholders rendered correctly
  - Some complex overrides may differ

- ⚠️ **Notes pages** - Supported
  - Layout matches PowerPoint
  - Note text rendered correctly

- ⚠️ **Handout layouts** - Basic support
  - Standard layouts supported
  - Custom layouts may differ

## DrawingML (Shapes) Limitations

### Not Supported

- ❌ **3D effects** - Rendered as 2D approximation
  - Rotation: Not rendered
  - Extrusion: Not rendered
  - Lighting: Ignored
  - Materials: Ignored

- ❌ **Advanced artistic effects**
  - Blur: Not rendered
  - Brightness/Contrast: Basic support
  - Artistic filters: Not rendered

- ❌ **Complex shape operations**
  - Boolean operations (union, subtract): Limited
  - Custom geometric formulas: Limited

### Partially Supported

- ⚠️ **Effects**
  - Shadow: Supported
  - Reflection: Supported
  - Glow: Supported (as outline)
  - Soft edges: Approximated

- ⚠️ **Gradients**
  - Linear: Fully supported
  - Radial: Fully supported
  - Path: Limited support
  - Complex multi-stop: Simplified

- ⚠️ **Text in shapes**
  - Standard layouts: Supported
  - Vertical text: Supported
  - Text on path: Not supported
  - 3D text: Rendered as 2D

## Performance Limitations

### Large Documents

Potential issues with very large documents:

- **Memory usage** - Large documents consume significant memory
  - Mitigation: Streaming output, page-at-a-time rendering
  - Recommendation: Limit to <1000 pages per render

- **Rendering time** - Complex documents take time
  - Typical: 1-5 seconds per page
  - Complex graphics: Up to 30 seconds per page
  - Recommendation: Use progress callbacks

- **File size** - PDF output can be large
  - Images: Major contributor to file size
  - Fonts: Each font adds 200KB-2MB
  - Mitigation: Image compression, font subsetting

### Memory Recommendations

Based on document complexity:

| Document Type | Recommended RAM |
|---------------|----------------|
| Simple text (Word) | 100 MB |
| Text + images (Word) | 500 MB |
| Complex tables (Word) | 200 MB |
| Large spreadsheet (Excel) | 500 MB |
| Many charts (Excel) | 1 GB |
| Presentation (PowerPoint) | 300 MB per 50 slides |

## Platform Limitations

### Font Availability

Different fonts available per platform:

- **Windows** - Best font coverage
- **macOS** - Good coverage, some Windows fonts missing
- **Linux** - Requires manual font installation
- **Docker** - Minimal fonts by default (install explicitly)

### OS-Specific Issues

- **Linux** - Font discovery may be slow on first run
- **Docker** - No GUI fonts installed by default
- **Windows** - Font names case-sensitive
- **macOS** - Some fonts in protected locations

## Known Issues

### High Priority

1. **Text layout edge cases**
   - Complex CJK text may break differently than Word
   - Status: Being investigated
   - Workaround: Use simpler layouts

2. **Table auto-fit accuracy**
   - Auto-fit columns may differ by 1-5 points
   - Status: Algorithm refinement in progress
   - Workaround: Use fixed column widths

3. **Font kerning precision**
   - Kerning may differ slightly from Word
   - Status: Improving font metrics extraction
   - Workaround: Accept minor differences (<1 point)

### Medium Priority

4. **Complex gradient rendering**
   - Multi-stop radial gradients may differ
   - Status: Planned improvement
   - Workaround: Use simpler gradients

5. **SmartArt complex layouts**
   - Some SmartArt layouts not perfect
   - Status: Expanding SmartArt support
   - Workaround: Convert to shapes in PowerPoint

6. **Chart axis label positioning**
   - Labels may shift slightly
   - Status: Refining chart renderer
   - Workaround: Accept minor differences

### Low Priority

7. **Obscure number formats**
   - Custom Excel number formats with complex conditions
   - Status: Expanding format parser
   - Workaround: Simplify number formats

8. **Rare Unicode characters**
   - May not render if font doesn't support
   - Status: Improving font fallback
   - Workaround: Use standard characters

## Future Improvements

### Planned Features

- [ ] PDF/UA (Universal Accessibility) compliance
- [ ] OCR layer for images containing text
- [ ] Improved SmartArt rendering
- [ ] More chart types (waterfall, funnel, etc.)
- [ ] Better 3D shape rendering
- [ ] Advanced OpenType features
- [ ] Form field preservation (non-interactive)
- [ ] Annotation preservation
- [ ] Digital signatures

### Performance Improvements

- [ ] Parallel page rendering
- [ ] Incremental font loading
- [ ] On-demand image decoding
- [ ] Memory-mapped font files
- [ ] Compiled geometry cache

### Fidelity Improvements

- [ ] Better text layout matching
- [ ] Improved kerning accuracy
- [ ] Complex table layout edge cases
- [ ] Advanced gradient patterns
- [ ] Font shaping for complex scripts

## Reporting Issues

When reporting limitations or issues:

1. **Check this document first** - Issue may be known
2. **Provide test document** - Minimal reproducible example
3. **Include output** - Both goffice-pdf and Office PDF
4. **Describe expectation** - What should happen?
5. **Measure impact** - How visible is the difference?
6. **Check workarounds** - Are alternatives available?

## Workarounds and Best Practices

### For Best Results

1. **Document creation**
   - Use standard fonts
   - Avoid complex layouts
   - Test early and often
   - Keep it simple

2. **Font handling**
   - Install required fonts
   - Use standard fonts when possible
   - Configure fallbacks

3. **Images**
   - Use standard formats (JPEG, PNG)
   - Optimize before inserting
   - Avoid huge images

4. **Tables**
   - Use fixed widths when possible
   - Avoid deeply nested tables
   - Simplify complex layouts

5. **Charts**
   - Use standard chart types
   - Simplify complex charts
   - Test output

## Conclusion

While goffice-pdf has some limitations, it provides high-fidelity PDF rendering for the majority of Office documents. Known limitations are documented, and workarounds are available for most cases. The project is actively developed, and limitations are being addressed in priority order based on user feedback.
