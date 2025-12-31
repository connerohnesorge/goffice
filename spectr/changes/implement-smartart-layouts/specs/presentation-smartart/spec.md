# Presentation SmartArt Spec Delta

## ADDED Requirements

### Requirement: Hierarchy Layout Algorithm
The system SHALL provide automatic tree layout for hierarchical SmartArt diagrams using the Reingold-Tilford algorithm.

#### Scenario: Simple org chart layout
- GIVEN a SmartArt diagram with 1 root and 2 children
- WHEN hierarchy layout is applied
- THEN root is positioned at top center
- AND children are evenly spaced below parent
- AND children are horizontally aligned

#### Scenario: Multi-level hierarchy
- GIVEN a SmartArt diagram with 3 levels (root, 3 managers, 6 employees)
- WHEN hierarchy layout is applied
- THEN levels are vertically spaced
- AND each level is centered horizontally
- AND tree is balanced (minimal width)

#### Scenario: Unbalanced tree layout
- GIVEN a tree with one branch having 5 levels and another having 2 levels
- WHEN hierarchy layout is applied
- THEN deep branch is positioned on one side
- AND shallow branch is positioned on other side
- AND branches do not overlap

#### Scenario: Assistant node positioning
- GIVEN a hierarchy with an assistant node (dotted line subordinate)
- WHEN layout is applied
- THEN assistant is positioned adjacent to manager
- AND assistant connector uses dotted line style

#### Scenario: Horizontal hierarchy orientation
- GIVEN a hierarchy diagram with horizontal orientation
- WHEN layout is applied
- THEN root is positioned at left center
- AND children extend to the right
- AND levels are horizontally spaced

#### Scenario: Vertical spacing configuration
- GIVEN a hierarchy layout with vertical spacing = 40 points
- WHEN layout is applied
- THEN vertical gap between levels is 40 points

#### Scenario: Horizontal spacing configuration
- GIVEN a hierarchy layout with horizontal spacing = 30 points
- WHEN layout is applied
- THEN horizontal gap between siblings is at least 30 points

### Requirement: List Layout Algorithm
The system SHALL provide automatic linear flow layout for list-style SmartArt diagrams.

#### Scenario: Vertical list layout
- GIVEN a SmartArt diagram with 5 items in vertical list
- WHEN layout is applied
- THEN items are stacked top to bottom
- AND items are evenly spaced vertically
- AND items are centered horizontally

#### Scenario: Horizontal list layout
- GIVEN a SmartArt diagram with 5 items in horizontal list
- WHEN layout is applied
- THEN items are arranged left to right
- AND items are evenly spaced horizontally
- AND items are centered vertically

#### Scenario: List with left alignment
- GIVEN a vertical list with left alignment
- WHEN layout is applied
- THEN all items are aligned to left edge of bounds

#### Scenario: List with right alignment
- GIVEN a vertical list with right alignment
- WHEN layout is applied
- THEN all items are aligned to right edge of bounds

#### Scenario: List connectors
- GIVEN a list with 4 items
- WHEN layout is applied
- THEN 3 connectors are created between consecutive items
- AND connectors are straight lines from item to next

#### Scenario: List auto-sizing
- GIVEN a list with container bounds 200 points wide
- AND items with varying text lengths
- WHEN layout is applied
- THEN all items fit within 200 point width

### Requirement: Cycle Layout Algorithm
The system SHALL provide automatic circular arrangement for cycle-style SmartArt diagrams.

#### Scenario: Circular positioning
- GIVEN a cycle diagram with 6 items
- WHEN layout is applied
- THEN items are positioned on circle perimeter
- AND items are evenly spaced angularly (60 degrees apart)

#### Scenario: Circle radius calculation
- GIVEN a cycle diagram with bounds 400x400
- WHEN layout is applied
- THEN circle radius is calculated to fit all items
- AND circle is centered in bounds
- AND radius is at most 80% of bounds radius

#### Scenario: Clockwise direction
- GIVEN a cycle with clockwise direction
- WHEN layout is applied
- THEN items are ordered clockwise starting from top
- AND arrows point clockwise

#### Scenario: Counter-clockwise direction
- GIVEN a cycle with counter-clockwise direction
- WHEN layout is applied
- THEN items are ordered counter-clockwise
- AND arrows point counter-clockwise

#### Scenario: Curved connector arrows
- GIVEN a cycle with 5 items
- WHEN layout is applied
- THEN 5 curved connectors are created
- AND connectors form circular flow
- AND arrows point to next item in sequence

#### Scenario: Center text in cycle
- GIVEN a cycle with centerText = true
- WHEN layout is applied
- THEN center text is positioned at circle center
- AND center text is visible

### Requirement: Pyramid Layout Algorithm
The system SHALL provide automatic stacked-level layout for pyramid-style SmartArt diagrams.

#### Scenario: Normal pyramid (apex top)
- GIVEN a pyramid diagram with 4 levels
- WHEN normal orientation layout is applied
- THEN top level is narrowest
- AND bottom level is widest
- AND width increases linearly from top to bottom

#### Scenario: Inverted pyramid (apex bottom)
- GIVEN a pyramid diagram with 4 levels
- WHEN inverted orientation layout is applied
- THEN top level is widest
- AND bottom level is narrowest
- AND width decreases linearly from top to bottom

#### Scenario: Level proportional sizing
- GIVEN a 3-level pyramid
- WHEN layout is applied
- THEN level 1 width is 33% of container
- AND level 2 width is 67% of container
- AND level 3 width is 100% of container

#### Scenario: Multiple items per level
- GIVEN a pyramid level with 3 items
- WHEN layout is applied
- THEN items are distributed evenly across level width
- AND items are centered vertically within level

#### Scenario: Level height distribution
- GIVEN a pyramid with 5 levels and container height 500
- WHEN layout is applied
- THEN each level has height of 100 points (500 / 5)

### Requirement: Matrix Layout Algorithm
The system SHALL provide automatic grid-based layout for matrix-style SmartArt diagrams.

#### Scenario: 2x2 matrix
- GIVEN a matrix diagram with 4 items and 2 rows, 2 columns
- WHEN layout is applied
- THEN items are arranged in 2x2 grid
- AND each cell is 50% of container width and height

#### Scenario: 3x3 matrix
- GIVEN a matrix diagram with 9 items and 3 rows, 3 columns
- WHEN layout is applied
- THEN items are arranged in 3x3 grid
- AND cells are evenly sized

#### Scenario: Auto-determined grid size
- GIVEN a matrix with 6 items and rows = 0, columns = 0
- WHEN layout is applied
- THEN grid size is auto-determined (e.g., 2x3 or 3x2)
- AND grid is approximately square

#### Scenario: Cell centering
- GIVEN a matrix with items smaller than cells
- WHEN layout is applied
- THEN items are centered within their cells
- AND items do not overlap cell boundaries

#### Scenario: Partial grid filling
- GIVEN a 3x3 matrix with only 7 items
- WHEN layout is applied
- THEN first 7 cells are filled
- AND last 2 cells are empty

### Requirement: Layout Constraint System
The system SHALL provide a constraint system for configuring layout spacing, sizing, and alignment.

#### Scenario: Horizontal spacing constraint
- GIVEN a layout with horizontal spacing = 25 points
- WHEN layout is applied
- THEN minimum horizontal gap between shapes is 25 points

#### Scenario: Vertical spacing constraint
- GIVEN a layout with vertical spacing = 30 points
- WHEN layout is applied
- THEN minimum vertical gap between shapes is 30 points

#### Scenario: Minimum shape size constraint
- GIVEN sizing constraint with minWidth = 60, minHeight = 40
- WHEN shapes are auto-sized
- THEN no shape is smaller than 60x40 points

#### Scenario: Maximum shape size constraint
- GIVEN sizing constraint with maxWidth = 150, maxHeight = 100
- WHEN shapes are auto-sized
- THEN no shape exceeds 150x100 points

#### Scenario: Aspect ratio constraint
- GIVEN sizing constraint with aspectRatio = 2.0 (width/height)
- WHEN shapes are auto-sized
- THEN all shapes maintain 2:1 aspect ratio

#### Scenario: Padding constraint
- GIVEN spacing constraint with padding = 8 points
- WHEN text is measured for shape sizing
- THEN 8 points padding is added on all sides

### Requirement: Text Measurement for Auto-Sizing
The system SHALL measure text content and automatically size shapes accordingly.

#### Scenario: Auto-size shape based on text
- GIVEN a data point with text "Project Manager"
- AND font "Arial" size 12 points
- WHEN shape is auto-sized
- THEN shape width accommodates text plus padding
- AND shape height accommodates font size plus padding

#### Scenario: Auto-size with long text
- GIVEN a data point with very long text (50 characters)
- AND maxWidth constraint of 200 points
- WHEN shape is auto-sized
- THEN shape width is capped at 200 points
- AND text wraps within shape

#### Scenario: Auto-size with font variation
- GIVEN data points with different font sizes (10pt, 14pt, 18pt)
- WHEN shapes are auto-sized
- THEN each shape size reflects its font size
- AND larger fonts result in larger shapes

#### Scenario: Text measurement caching
- GIVEN 100 shapes with same text and font
- WHEN auto-sizing is performed
- THEN text is measured once and cached
- AND cached measurement is reused for all 100 shapes

### Requirement: Layout Engine Registry
The system SHALL maintain a registry of available layout engines for dynamic selection.

#### Scenario: Register layout engine
- GIVEN a custom layout engine implementation
- WHEN Register(LayoutCustom, engine) is called
- THEN engine is added to registry

#### Scenario: Get layout engine by type
- GIVEN registry with hierarchy layout registered
- WHEN GetEngine(LayoutHierarchy) is called
- THEN hierarchy layout engine is returned

#### Scenario: Get unregistered layout engine
- WHEN GetEngine(9999) is called for unregistered type
- THEN an error is returned
- AND error message indicates type not found

#### Scenario: List supported layout types
- GIVEN registry with 5 layout engines registered
- WHEN SupportedTypes() is called
- THEN list of 5 layout type constants is returned

### Requirement: Connector Routing
The system SHALL automatically route connectors between shapes based on layout type.

#### Scenario: Straight-line connectors
- GIVEN a list layout with items A, B, C
- WHEN connectors are routed
- THEN connector from A to B is straight line
- AND connector from B to C is straight line

#### Scenario: Curved connectors
- GIVEN a cycle layout with items in circular arrangement
- WHEN connectors are routed
- THEN connectors follow circular path (bezier curves)
- AND connectors do not cross through center

#### Scenario: Connector endpoints
- GIVEN shapes positioned by layout
- WHEN connectors are routed
- THEN connector starts at center-bottom of source shape
- AND connector ends at center-top of destination shape

#### Scenario: Dotted line connectors
- GIVEN a hierarchy with assistant node
- WHEN connectors are routed
- THEN assistant connector uses dotted line style
- AND regular connectors use solid line style

#### Scenario: Arrow heads
- GIVEN connectors with directional flow
- WHEN connectors are routed
- THEN arrow heads are added at destination end
- AND arrow heads point towards destination shape

### Requirement: Layout Performance
The system SHALL optimize layout calculations for reasonable performance with typical diagrams.

#### Scenario: Layout 100-node hierarchy
- GIVEN a hierarchy diagram with 100 nodes
- WHEN Layout() is called
- THEN layout completes in under 100 milliseconds

#### Scenario: Layout 50-item list
- GIVEN a list diagram with 50 items
- WHEN Layout() is called
- THEN layout completes in under 50 milliseconds

#### Scenario: Layout caching performance
- GIVEN a diagram with layout already applied
- WHEN Layout() is called again without data changes
- THEN layout completes in under 1 millisecond (cached)

### Requirement: Layout Error Handling
The system SHALL handle invalid or edge-case diagram structures gracefully.

#### Scenario: Layout empty diagram
- GIVEN a SmartArt diagram with zero data points
- WHEN Layout() is called
- THEN no error occurs
- AND no shapes are positioned

#### Scenario: Layout with circular dependencies
- GIVEN a hierarchy diagram with circular reference (A → B → A)
- WHEN Layout() is called
- THEN an error is returned
- AND error indicates circular dependency

#### Scenario: Layout with disconnected nodes
- GIVEN a diagram with nodes A, B, C where C is orphaned
- WHEN Layout() is called
- THEN connected nodes (A, B) are laid out
- AND orphaned node (C) is positioned separately

#### Scenario: Layout exceeding bounds
- GIVEN a large diagram that doesn't fit in small bounds
- WHEN Layout() is called
- THEN shapes are scaled down to fit
- AND aspect ratios are preserved

#### Scenario: Layout with invalid connector
- GIVEN a diagram with connector referencing non-existent node
- WHEN Layout() is applied
- THEN invalid connector is ignored
- AND valid connectors are routed normally
