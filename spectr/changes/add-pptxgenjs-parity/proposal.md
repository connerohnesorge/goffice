# Change: Achieve Full Functional Parity with pptxgenjs

## Why
Currently, goffice provides comprehensive PowerPoint support with slides, shapes, text, tables, charts, and SmartArt. However, there's a functional gap compared to pptxgenjs - the leading JavaScript PowerPoint generation library. pptxgenjs offers advanced features that goffice lacks, limiting our ability to serve users migrating from JavaScript environments or requiring rich presentation capabilities. Achieving parity will make goffice the definitive Go solution for PowerPoint generation.

## What Changes
This is a multi-faceted enhancement across several capabilities:

### Core Presentation Features
- **Media Support**: Add video and audio embedding capabilities
- **Speaker Notes**: Support adding presentation notes to slides
- **Master Slides**: Implement master slide and template system for consistent branding
- **RTL Text**: Full right-to-left language support (Arabic, Hebrew)
- **HTML to PowerPoint**: Convert HTML tables/content to presentation format

### Enhanced Shapes & Formatting
- **Rich Text Shapes**: Advanced text formatting with multiple formatting runs
- **Additional Shape Types**: Support for equation shapes, callouts, banners
- **Shape Effects**: Glow, shadow, reflection presets
- **Shape Styles**: Quick style application system

### Chart Enhancements
- **Chart Types**: Add missing chart types (bubble, radar, surface, stock, treemap, sunburst, box & whisker, waterfall)
- **Chart Formatting**: Enhanced chart styling and formatting options
- **Chart Data Labels**: Rich data label formatting and positioning
- **Chart Axes**: Multiple axis support, logarithmic scales

### Developer Experience
- **Fluent APIs**: Chainable APIs for easier document construction
- **Helper Methods**: Convenience methods for common operations
- **Better Defaults**: Sensible defaults reducing boilerplate code
- **Validation**: Enhanced validation with helpful error messages

## Impact
- **Affected specs:** presentation, drawingml, validation, pdf-rendering
- **Affected code:**
  - `presentation/document.go` - Core API enhancements
  - `presentation/parts/slide_part.go` - Slide content APIs
  - `presentation/elements/*.go` - New shape and formatting elements
  - `drawingml/chart*.go` - Chart enhancements
  - New packages for media and master slides
- **Breaking changes:** Minimal - mostly additive with some API deprecations
