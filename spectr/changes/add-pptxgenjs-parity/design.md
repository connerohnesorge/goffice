# Design: Achieving pptxgenjs Parity

## Overview

This design document outlines the architectural approach for adding pptxgenjs-equivalent features to goffice. The goal is comprehensive feature parity while maintaining goffice's Go idiomatic patterns and high-quality implementation standards.

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

**Trade-offs**:
- Requires additional part management complexity
- Increases package size with media files
- Needs format validation and potentially transcoding guidance

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

**Trade-offs**:
- Complex inheritance management required
- Need careful handling of override precedence
- Memory overhead for multiple masters

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

**Trade-offs**:
- External dependency (already used by goffice)
- Complexity in mixed LTR/RTL text handling
- Performance impact for large RTL documents

### 4. HTML Conversion Module

**Decision**: Implement HTML parsing and conversion as separate module

**Rationale**:
- Clean separation from core presentation logic
- Enables potential reuse in other contexts (Word, Excel)
- Follows single responsibility principle
- Easier testing and maintenance

**Implementation**:
- `html2pptx` package with conversion functions
- Uses Go's `html` package for parsing
- CSS style extraction and mapping
- Modular mapping functions for each HTML element type

**Trade-offs**:
- Additional package to maintain
- Limited CSS support (focus on inline styles)
- May need updates for new HTML elements

### 5. Chart Type Extensions

**Decision**: Extend existing chart framework with new chart types

**Rationale**:
- Current chart implementation provides good foundation
- Consistent API across all chart types
- Leverages existing chart parts and relationships
- Maintains compatibility with existing code

**Implementation**:
- Extend `chart_types.go` with new chart classes
- Add type-specific rendering in `pdf/drawing/`
- Update validation for new chart configurations
- Maintain XML schema compliance for each type

**Trade-offs**:
- Risk of chart code complexity increase
- Need for extensive testing of each type
- XML generation differences between chart types

### 6. Fluent API Design

**Decision**: Add fluent interfaces alongside existing APIs (non-breaking)

**Rationale**:
- Maintains backward compatibility
- Offers pptxgenjs-like developer experience
- Enables both imperative and fluent coding styles
- Reduces boilerplate for common patterns

**Implementation**:
- Fluent methods return receiver/self
- Builder pattern for complex objects (charts, SmartArt)
- Extension methods/higher-order functions for collections
- Helper methods for common slide patterns

**Trade-offs**:
- API surface area increases significantly
- Potential confusion between API styles
- Need for comprehensive documentation

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

### Theme Integration
Coordinate with existing theme system:
- Master slides respect theme colors 🎨
- Chart formatting uses theme palettes
- RTL text respects theme fonts
- Media players respect theme colors

## Performance Considerations

1. **Lazy Loading**: Media files loaded on demand
2. **Streaming**: Large videos/audio use io.Reader
3. **Caching**: Parsed HTML structures cached when appropriate
4. **Parallel Processing**: Slide generation can be parallelized
5. **Memory Management**: Proper cleanup of temporary resources

## Compatibility and Migration

### Backward Compatibility
- All new features are additive
- Existing APIs remain unchanged
- Graceful degradation for unsupported features
- Clear feature detection capabilities

### Migration from pptxgenjs
- Provide wrapper functions matching pptxgenjs names
- Create compatibility layer examples
- Document equivalent patterns and APIs
- Highlight Go-specific improvements

## Testing Strategy

1. **Unit Tests**: Per-feature testing with mocks
2. **Integration Tests**: Round-trip tests (create → save → open → verify)
3. **E2E Tests**: Visual comparison with pptxgenjs output
4. **Cross-Platform**: Test on Windows, macOS, Linux
5. **Real Files**: Test with actual PowerPoint installations

## Phased Implementation

**Phase 1**: Core media, RTL, basic Fluent APIs
**Phase 2**: Charts, shapes, rich text
**Phase 3**: Masters, notes, HTML conversion
**Phase 4**: Animations, transitions, comments
**Phase 5**: Polish, documentation, examples

## Risk Mitigation

1. **Feature Creep**: Strict scope adherence to pptxgenjs parity
2. **Breaking Changes**: Maintain backward compatibility
3. **Performance**: Benchmark and optimize early
4. **Complexity**: Refactor before adding features
5. **Quality**: Comprehensive test coverage required

## Success Metrics

- Feature parity checklist completion
- Performance benchmarks within 10% of baseline
- 90%+ test coverage for new code
- Zero breaking changes in existing APIs
- Positive feedback from beta testers
