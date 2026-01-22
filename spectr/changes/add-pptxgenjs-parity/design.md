# Design: Achieving pptxgenjs Parity

## Overview

This design document outlines the architectural approach for adding pptxgenjs-equivalent features to goffice while maintaining goffice's Go idiomatic patterns and high-quality implementation standards.

## Core Architectural Decisions

### 1. Media Support Architecture
**Decision**: Implement media embedding via the existing part system

**Rationale**:
- Leverage existing `packaging.Part` infrastructure
- Consistent with current image, chart, and diagram implementations
- Maintains proper OPC (Open Packaging Conventions) compliance
- Enables media relationships and content types

**Implementation**:
- Create `MediaPart` interface extending `packaging.Part`
- Implement `VideoPart` and `AudioPart` with format-specific handling
- Add media registry for format detection and validation
- Use streaming for large media files to reduce memory usage

### 2. Master Slides and Template System
**Decision**: Implement master slides as first-class parts with inheritance

**Rationale**:
- PowerPoint's master system is well-established
- Enables corporate template support
- Provides layout consistency across presentations
- Required for slide library and template features

**Implementation**:
- `SlideMasterPart` manages master slide definitions
- `SlideLayoutPart` for specific layout templates
- Inheritance chain: Slide → Layout → Master → Theme
- Override system for slide-specific customizations

### 3. RTL Text and Bidirectional Support
**Decision**: Integrate with `golang.org/x/text` for proper BiDi support

**Rationale**:
- PowerPoint compliance requires proper Unicode handling
- Existing bi-directional algorithm implementations available
- Ensures correct character reordering and shaping
- Required for Arabic, Hebrew, Persian language support

**Implementation**:
- Use `golang.org/x/text/unicode/bidi` for text processing
- Integrate Arabic shaping via `golang.org/x/text/language`
- Text runs marked with directionality metadata
- Paragraph-level direction controls

## Integration Points

### PDF Rendering
All new features must integrate with `pdf/` rendering system:
- Media: PDF attachment handling or fallbacks
- RTL: Proper text rendering in PDF
- Charts: Additional chart type renderers
- Animations: Progressive enhancement (static views)

### Validation System
Extend validation framework for new features:
- Media file format and codec validation
- RTL text compliance checking
- Chart data validation for new types
- Master slide consistency validation

## Phased Implementation

**Phase 1**: Core media, RTL, basic Fluent APIs (1-2 weeks)
**Phase 2**: Charts, shapes, rich text (2-3 weeks)
**Phase 3**: Masters, notes, HTML conversion (1-2 weeks)
**Phase 4**: Animations, transitions, comments (2-3 weeks)
**Phase 5**: Polish, documentation, examples (1 week)

**Total estimated timeline**: 8-12 weeks

## Performance Considerations

1. **Lazy Loading**: Media files loaded on demand
2. **Streaming**: Large videos/audio use io.Reader
3. **Caching**: Parsed HTML structures cached when appropriate
4. **Parallel Processing**: Slide generation can be parallelized
5. **Memory Management**: Proper cleanup of temporary resources
