# Presentation Elements Spec Delta

## ADDED Requirements

### Requirement: SmartArt Layout Application
The system SHALL provide methods to apply automatic layout algorithms to SmartArt diagrams.

#### Scenario: Apply layout to SmartArt diagram
- GIVEN a SmartArt diagram with data model populated
- WHEN Layout() is called on the diagram
- THEN shapes are automatically positioned based on the layout algorithm
- AND connectors are routed between shapes

#### Scenario: Layout respects container bounds
- GIVEN a SmartArt diagram with bounds 400x300 points
- WHEN Layout() is called
- THEN all shapes are positioned within the bounds
- AND shapes are scaled if necessary to fit

#### Scenario: Layout with empty diagram
- GIVEN a SmartArt diagram with no data points
- WHEN Layout() is called
- THEN no error occurs
- AND no shapes are positioned

### Requirement: SmartArt Layout Type Selection
The system SHALL allow selection and changing of layout algorithm types.

#### Scenario: Set layout type to hierarchy
- GIVEN a SmartArt diagram
- WHEN SetLayoutType(LayoutHierarchy) is called
- THEN layout engine is set to hierarchy algorithm
- AND subsequent Layout() calls use hierarchy positioning

#### Scenario: Change layout type dynamically
- GIVEN a SmartArt diagram with list layout applied
- WHEN SetLayoutType(LayoutCycle) is called
- THEN layout engine changes to cycle algorithm
- AND if auto-layout is enabled, layout is immediately reapplied

#### Scenario: Set invalid layout type
- GIVEN a SmartArt diagram
- WHEN SetLayoutType(999) is called with invalid type
- THEN an error is returned
- AND layout engine remains unchanged

### Requirement: SmartArt Auto-Layout
The system SHALL support automatic re-layout when diagram data changes.

#### Scenario: Auto-layout enabled
- GIVEN a SmartArt diagram with AutoLayout(true)
- WHEN a data point is added
- THEN layout is automatically recalculated
- AND shapes are repositioned

#### Scenario: Auto-layout disabled
- GIVEN a SmartArt diagram with AutoLayout(false)
- WHEN a data point is added
- THEN layout is NOT automatically recalculated
- AND shapes remain in previous positions

#### Scenario: Auto-layout toggle
- GIVEN a SmartArt diagram with auto-layout enabled
- WHEN AutoLayout(false) is called
- THEN auto-layout is disabled
- AND future data changes do not trigger layout

### Requirement: SmartArt Layout Engine Access
The system SHALL provide access to the current layout engine for inspection and configuration.

#### Scenario: Get current layout engine
- GIVEN a SmartArt diagram with hierarchy layout
- WHEN GetLayoutEngine() is called
- THEN the hierarchy layout engine is returned

#### Scenario: Query supported layout types
- GIVEN a layout engine
- WHEN SupportedTypes() is called
- THEN list of supported layout type constants is returned

### Requirement: SmartArt Programmatic Creation
The system SHALL provide high-level API for creating SmartArt diagrams programmatically.

#### Scenario: Create new SmartArt with layout type
- WHEN NewSmartArt(LayoutHierarchy) is called
- THEN a new SmartArt diagram is created
- AND layout engine is set to hierarchy
- AND auto-layout is enabled by default

#### Scenario: Create SmartArt from data structure
- GIVEN a SmartArtData structure with 5 items and 4 connections
- WHEN SmartArtFromData(data, LayoutList) is called
- THEN a SmartArt diagram is created with data populated
- AND list layout is applied
- AND shapes are automatically positioned

#### Scenario: SmartArt creation with invalid layout
- WHEN NewSmartArt(999) is called with invalid type
- THEN an error is returned
- AND no SmartArt diagram is created

### Requirement: SmartArt Data Manipulation
The system SHALL provide methods to add and remove data points and connections.

#### Scenario: Add data point
- GIVEN a SmartArt diagram
- WHEN AddPoint("Manager") is called
- THEN a new data point is added to diagram
- AND point ID is returned
- AND if auto-layout enabled, layout is recalculated

#### Scenario: Add connection between points
- GIVEN a SmartArt diagram with points "A" and "B"
- WHEN AddConnection("A", "B", ConnectionParentChild) is called
- THEN a parent-child connection is created
- AND if auto-layout enabled, connectors are routed

#### Scenario: Remove data point
- GIVEN a SmartArt diagram with point "X"
- WHEN RemovePoint("X") is called
- THEN point is removed from diagram
- AND all connections to/from point are removed
- AND if auto-layout enabled, layout is recalculated

#### Scenario: Remove connection
- GIVEN a SmartArt diagram with connection "conn1"
- WHEN RemoveConnection("conn1") is called
- THEN connection is removed
- AND connector is removed from rendering

### Requirement: SmartArt Layout Caching
The system SHALL cache layout results to avoid redundant calculations.

#### Scenario: Layout cached after application
- GIVEN a SmartArt diagram with layout applied
- WHEN Layout() is called again without data changes
- THEN layout calculation is skipped
- AND cached positions are used

#### Scenario: Layout invalidated on data change
- GIVEN a SmartArt diagram with cached layout
- WHEN a data point is added
- THEN layout cache is invalidated
- AND next Layout() call recalculates positions

#### Scenario: Layout version tracking
- GIVEN a SmartArt diagram with data version 1
- WHEN data is modified
- THEN data version increments to 2
- AND layout version becomes stale
