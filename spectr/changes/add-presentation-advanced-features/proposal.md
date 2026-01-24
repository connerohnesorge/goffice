# Change: Add Advanced Presentation Features for Open-XML-SDK Parity

## Why
goffice currently implements core PowerPoint presentation functionality with basic slide, shape, and media support. However, it lacks many advanced features present in Open-XML-SDK:
- Master slides and layouts management
- Custom slide sizes and orientations
- Slide transitions with advanced timing
- Shape animations and animation effects
- Hyperlinks and action buttons
- SmartArt and diagram support
- Custom color schemes and themes
- Handout and notes page customization
- Speaker notes with rich formatting
- Slide numbering and custom footers
- Protected presentations (password, read-only)
- Embedded OLE objects
- Section management within presentations
- Outline view support
- Presenter view and display settings

## What Changes
- Add master slide creation and customization
- Add slide layout templates with inheritance
- Add custom slide dimensions and aspect ratios
- Add slide transition effects with duration and trigger options
- Add shape animation (entrance, exit, emphasis, motion path)
- Add animation timing and sequencing
- Add hyperlinks with actions (open file, go to URL, jump to slide, custom action)
- Add action buttons for interactivity
- Add SmartArt diagram support (organizational charts, process flows, etc.)
- Add color scheme customization
- Add theme management
- Add notes pages with rich text
- Add handout configuration
- Add slide numbers and headers/footers with placeholders
- Add presentation protection
- Add section support for slide organization
- BREAKING: Slide and shape APIs may be extended

## Impact
- Affected specs: presentation-elements, presentation-document, presentation-parts, drawingml-core
- Affected code: presentation/document.go, presentation/elements/, presentation/parts/
- New dependencies: None (uses stdlib only)
- Test coverage required: Comprehensive tests for all animation and master slide features
