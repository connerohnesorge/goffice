# Change: Implement SmartArt Automatic Layout Engine

## Why

goffice currently has SmartArt element support from Phase 1 (schema generation and roundtrip), but lacks the automatic layout engine that is core to SmartArt functionality. SmartArt diagrams are designed to automatically position and size shapes based on the chosen layout algorithm, making it trivial to create professional diagrams like org charts, process flows, and hierarchies.

**Current State (After Phase 1)**:
- SmartArt elements can be read and preserved (roundtrip works)
- Diagram data model can be parsed (nodes, connections)
- No layout algorithms implemented
- Diagrams are static - adding/removing items doesn't trigger re-layout
- No support for layout constraints (hierarchy, flow, cycle, etc.)
- Manual positioning required for all SmartArt shapes
- Cannot programmatically create SmartArt with automatic positioning

**Impact Without This Change**:
- **PowerPoint**: Cannot create org charts, process diagrams, hierarchies programmatically
- **User Experience**: SmartArt diagrams don't auto-adjust when content changes
- **Office Compatibility**: SmartArt from Office applications loses layout intelligence on modification
- **PDF Rendering**: No positioned shapes (only placeholder boxes from Phase 1)

**Evidence**:
- presentation/elements has SmartArt types but no layout engine
- No algorithm implementations for hierarchy, list, cycle, pyramid, matrix layouts
- PDF rendering shows only bounding box placeholders (from Phase 1)
- No programmatic SmartArt creation API

## What Changes

**Phase 2 Scope: Core Layout Algorithms**

This proposal adds **5 fundamental layout algorithms** that cover 80% of SmartArt usage:

1. **Hierarchy Layout** (org chart, tree) - Vertical/horizontal tree layout with balanced positioning
2. **List Layout** (bullet points, process) - Linear vertical/horizontal flow with equal spacing
3. **Cycle Layout** (circular process) - Circular arrangement with directional arrows
4. **Pyramid Layout** (hierarchical levels) - Stacked triangle layout with proportional sizing
5. **Matrix Layout** (2x2 grid) - Grid-based positioning with cell alignment

**New Capabilities**:
- Automatic shape positioning based on layout type
- Automatic shape sizing based on text content
- Automatic connector routing between shapes
- Layout recalculation when data model changes
- Programmatic SmartArt creation API
- Full PDF rendering with positioned shapes

**Breaking Changes**: None. This is additive functionality building on Phase 1.

## Impact

**Affected specs**:
- `presentation-elements` - ADDED layout engine requirements
- `presentation-smartart` - NEW spec for SmartArt layout algorithms
- `pdf-drawing` - MODIFIED diagram rendering (replace placeholder with full rendering)

**New capabilities**:
- Create SmartArt diagrams programmatically
- Automatic layout application when data model changes
- Layout constraint system (spacing, sizing, alignment)
- 5 core layout algorithm implementations
- Full SmartArt PDF rendering

**Affected code**:
- `presentation/smartart/layout/` - NEW PACKAGE: Layout engine and algorithms
- `presentation/smartart/constraints/` - NEW PACKAGE: Layout constraints system
- `presentation/elements/` - MODIFIED: Add Layout() method to SmartArt types
- `presentation/` - MODIFIED: Add NewSmartArt() and SmartArtFromData() API
- `pdf/drawing/diagram_renderer.go` - MODIFIED: Replace placeholder with full rendering
- `drawingml/diagram/` - MODIFIED: Add layout metadata accessors

## Key Design Decisions

### 1. Five-Algorithm Core
**Decision**: Implement only 5 most common layout types in Phase 2

**Rationale**:
- 80/20 rule: Hierarchy, List, Cycle, Pyramid, Matrix cover ~80% of use cases
- Office has 50+ layout types, but most are variations
- Focused scope allows high-quality implementation
- Foundation for future layout additions
- Each algorithm requires significant effort (2-3 weeks per layout)

**Alternatives Considered**:
- All 50+ layouts: Too large, would delay delivery by months
- Only 1-2 layouts: Insufficient to demonstrate value
- 10 layouts: Still too broad for initial release

### 2. Layout Engine Architecture
**Decision**: Interface-based extensible design with algorithm registry

```go
type LayoutEngine interface {
    Layout(diagram *SmartArt, bounds Rectangle) error
    CalculateBounds(diagram *SmartArt) Rectangle
    SupportedTypes() []LayoutType
}
```

**Rationale**:
- Extensibility: Easy to add new layouts in future
- Testability: Each algorithm can be tested in isolation
- Separation of concerns: Algorithm logic separate from diagram data model
- Follows existing goffice patterns (e.g., function registry in formulas)

### 3. Constraint-Based Positioning
**Decision**: Separate constraint system for spacing, sizing, alignment

**Rationale**:
- Constraints are reusable across multiple layout types
- Allows fine-tuning without algorithm changes
- Matches Office SmartArt constraint model
- Enables future user-configurable layouts

### 4. Text Measurement Integration
**Decision**: Use existing PDF text layout engine for shape auto-sizing

**Rationale**:
- Avoid duplicating text measurement logic
- Consistent sizing between layout and PDF rendering
- Leverages existing font handling infrastructure
- Accurate sizing requires font metrics

### 5. PDF Rendering Integration
**Decision**: Replace Phase 1 placeholder renderer with full positioned rendering

**Rationale**:
- Layout engine provides positioned shapes
- Enables high-fidelity PDF output
- Completes SmartArt PDF rendering story
- No additional complexity (shapes already render)

## Implementation Scope

**Timeline: 8-10 weeks**

### Week 1-2: Layout Engine Foundation
- Create layout engine interfaces
- Implement constraint system (spacing, sizing, alignment)
- Add SmartArt.Layout() API
- Add layout type selection

### Week 3-4: Hierarchy Layout
- Tree layout algorithm (balanced positioning)
- Parent-child relationship handling
- Assistant node support (dotted lines)
- Horizontal and vertical variants

### Week 5: List Layout
- Vertical and horizontal flow
- Equal spacing calculation
- Connector placement
- Text wrapping support

### Week 6: Cycle Layout
- Circular positioning on perimeter
- Equal angular spacing
- Directional arrow placement
- Center text support

### Week 7: Pyramid & Matrix Layouts
- Pyramid: Stacked levels with proportional sizing
- Matrix: Grid positioning with cell alignment
- Merged cell support (matrix)

### Week 8: PDF Rendering & Testing
- Replace placeholder with full rendering
- Visual regression tests
- Roundtrip tests (layout → save → load → verify)
- Integration tests with all 5 layouts

### Week 9-10: API & Documentation
- Programmatic creation API (NewSmartArt, SmartArtFromData)
- Auto-layout toggle (enable/disable)
- Examples for each layout type
- Documentation and godoc

### Success Criteria

- [ ] 5 layout algorithms implemented (hierarchy, list, cycle, pyramid, matrix)
- [ ] Layout engine interface and constraint system complete
- [ ] SmartArt.Layout() API functional
- [ ] Auto-layout triggers when items added/removed
- [ ] Layout respects container bounds
- [ ] Text content auto-sizes shapes
- [ ] Connectors auto-route between shapes
- [ ] Round-trip preserves layout type and positioning
- [ ] PDF rendering shows fully laid-out diagrams
- [ ] Programmatic creation API (NewSmartArt, SmartArtFromData)
- [ ] All tests pass (unit, integration, visual regression)
- [ ] Documentation complete with examples for each layout

## Out of Scope

**Deferred to Future Proposals**:
- Advanced layout types (picture, relationship, venn, 3D)
- Custom user-defined layouts
- SmartArt styles and themes (colors, effects)
- Animation of layout transitions
- SmartArt quick styles
- SmartArt color schemes (separate from layout)
- Advanced text formatting within shapes
- Shape effects (shadows, reflections, glows)

**Not Planned**:
- 3D SmartArt layouts
- SmartArt animation sequences
- Live preview during editing

## Dependencies

**Existing Infrastructure**:
- SmartArt elements from Phase 1 (`drawingml/diagram/`)
- Diagram parts from Phase 1 (data, layout, style, colors)
- DrawingML shape positioning (`drawingml/`)
- PDF text layout engine (`pdf/layout/`)
- PDF font handling (`pdf/font/`)

**New Dependencies**: None. All required infrastructure exists.

## Risks & Mitigation

### Risk: Layout Algorithm Complexity
**Impact**: High (delays delivery)
**Probability**: Medium (algorithms are non-trivial)
**Mitigation**:
- Start with simplest layout (list) to validate approach
- Use Office reference documents for algorithm validation
- Visual regression tests to catch positioning errors
- Phased delivery: release layouts as completed

### Risk: Text Measurement Accuracy
**Impact**: Medium (shapes too small/large)
**Probability**: Low (existing PDF layout engine is mature)
**Mitigation**:
- Reuse proven PDF text measurement
- Test with various fonts and sizes
- Add padding for visual breathing room

### Risk: Connector Routing Complexity
**Impact**: Medium (overlapping connectors)
**Probability**: Medium (routing is hard in general)
**Mitigation**:
- Start with simple straight-line connectors
- Defer advanced routing to future work
- Document limitations clearly

### Risk: Office Compatibility
**Impact**: High (diagrams don't match Office output)
**Probability**: Medium (Office algorithms are undocumented)
**Mitigation**:
- Test with real Office documents
- Visual comparison with Office output
- Iterate based on fidelity testing
- Document known differences

## References

### Office Open XML Spec
- ECMA-376 Part 1, Section 21.4: Diagrams
- DiagramLayout element structure
- LayoutNode hierarchy
- Constraint and rule systems

### OpenXML-SDK References
- `DocumentFormat.OpenXml.Drawing.Diagrams` namespace
- DiagramLayout, LayoutNode, Constraint elements
- Layout algorithm metadata

### Existing goffice Code
- `pdf/layout/` - Text layout engine (for shape auto-sizing)
- `drawingml/diagram/` - Generated diagram elements (from Phase 1)
- `presentation/parts/diagram_part.go` - Diagram parts (from Phase 1)
- `pdf/drawing/` - Shape rendering

## Files to Create/Modify

### Layout Engine
1. `presentation/smartart/layout/engine.go` - NEW: LayoutEngine interface, algorithm registry
2. `presentation/smartart/layout/hierarchy.go` - NEW: Hierarchy layout algorithm
3. `presentation/smartart/layout/list.go` - NEW: List layout algorithm
4. `presentation/smartart/layout/cycle.go` - NEW: Cycle layout algorithm
5. `presentation/smartart/layout/pyramid.go` - NEW: Pyramid layout algorithm
6. `presentation/smartart/layout/matrix.go` - NEW: Matrix layout algorithm

### Constraints
7. `presentation/smartart/constraints/spacing.go` - NEW: Spacing constraints
8. `presentation/smartart/constraints/sizing.go` - NEW: Sizing constraints
9. `presentation/smartart/constraints/alignment.go` - NEW: Alignment constraints

### API
10. `presentation/smartart.go` - NEW: High-level SmartArt creation API
11. `presentation/elements/smartart.go` - MODIFY: Add Layout() method

### PDF Rendering
12. `pdf/drawing/diagram_renderer.go` - MODIFY: Replace placeholder with full rendering

### Testing
13. `presentation/smartart/layout/hierarchy_test.go` - NEW: Hierarchy layout tests
14. `presentation/smartart/layout/list_test.go` - NEW: List layout tests
15. `presentation/smartart/layout/cycle_test.go` - NEW: Cycle layout tests
16. `presentation/smartart/layout/pyramid_test.go` - NEW: Pyramid layout tests
17. `presentation/smartart/layout/matrix_test.go` - NEW: Matrix layout tests
18. `presentation/smartart/layout/integration_test.go` - NEW: Integration tests
19. `testdata/smartart/layouts/` - NEW: Visual regression test fixtures

### Documentation
20. `presentation/smartart/layout/doc.go` - NEW: Package documentation
21. `presentation/smartart/doc.go` - NEW: SmartArt API documentation
22. `examples/presentation/smartart/` - NEW: Example programs
23. `CHANGELOG.md` - MODIFY: Document SmartArt layout support

## Estimated Effort

**Phase 2 (This Proposal): 8-10 weeks**
- Layout engine foundation: 2 weeks
- Hierarchy layout: 2 weeks
- List layout: 1 week
- Cycle layout: 1 week
- Pyramid & Matrix layouts: 1 week
- PDF rendering integration: 1 week
- API & documentation: 1 week
- Testing & refinement: 1-2 weeks

**Total**: ~320-400 hours (1 full-time developer)

## Migration Path

**For Users**:
- After Phase 2: Programmatic SmartArt creation becomes possible
- Existing code from Phase 1 (roundtrip) continues to work
- New Layout() API is opt-in

**For Developers**:
- Phase 2 builds on Phase 1 foundation (no breaking changes)
- Layout engine extensible for future layout types
- API follows existing goffice patterns

## Future Enhancements (Not in Scope)

These will be addressed in separate proposals:
1. SmartArt styles and themes (Phase 3)
2. Advanced layout types (picture, relationship, venn, radial)
3. Custom layout definitions
4. SmartArt animation (presentation-specific)
5. SmartArt quick styles
6. SmartArt color schemes and effects
