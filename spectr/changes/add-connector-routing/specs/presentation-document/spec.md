# Presentation Document Spec Delta

## ADDED Requirements

### Requirement: High-Level Shape Connection API

The system SHALL provide high-level methods on Slide to connect shapes with automatic routing.

#### Scenario: Connect two shapes with automatic routing
- GIVEN a Slide with two Shape elements (shape1 and shape2)
- AND shape1 is positioned at (100, 100)
- AND shape2 is positioned at (500, 500)
- WHEN connector := slide.ConnectShapes(shape1, shape2, RouterTypeElbow) is called
- THEN a ConnectionShape is created and added to the slide
- AND the ConnectionShape's start connection references shape1
- AND the ConnectionShape's end connection references shape2
- AND the ConnectionShape's path is routed using ElbowRouter
- AND the returned connector is non-nil

#### Scenario: Connect shapes with automatic connection site selection
- GIVEN a Slide with shape1 at (100, 100) and shape2 at (500, 500)
- WHEN slide.ConnectShapes(shape1, shape2, routerType) is called
- THEN the system selects nearest connection sites on both shapes
- AND start connection site is the site on shape1 closest to shape2
- AND end connection site is the site on shape2 closest to shape1
- AND the connector is attached to the selected sites

#### Scenario: Connect shapes with straight routing
- GIVEN a Slide with two shapes
- WHEN connector := slide.ConnectShapes(shape1, shape2, RouterTypeStraight) is called
- THEN a ConnectionShape with straight line path is created
- AND the path has 2 commands (MoveTo and LineTo)
- AND the path connects the two nearest connection sites

#### Scenario: Connect shapes with elbow routing
- GIVEN a Slide with two shapes
- AND other shapes on the slide as obstacles
- WHEN connector := slide.ConnectShapes(shape1, shape2, RouterTypeElbow) is called
- THEN a ConnectionShape with orthogonal path is created
- AND the path avoids overlapping obstacle shapes
- AND the path uses only horizontal and vertical segments

#### Scenario: Connect shapes with curved routing
- GIVEN a Slide with two shapes
- WHEN connector := slide.ConnectShapes(shape1, shape2, RouterTypeCurved) is called
- THEN a ConnectionShape with Bezier curve path is created
- AND the curve is smooth and visually appealing
- AND control points are calculated based on connection site angles

#### Scenario: Connect shapes returns connector reference
- GIVEN a Slide with two shapes
- WHEN connector := slide.ConnectShapes(shape1, shape2, routerType) is called
- THEN the returned connector is a valid *ConnectionShape
- AND the connector is accessible via slide.ShapeTree().ConnectionShapes()
- AND the connector can be further manipulated (styling, properties, etc.)

#### Scenario: ConnectShapes with nil shape
- GIVEN a Slide
- WHEN slide.ConnectShapes(nil, shape2, routerType) is called
- THEN nil is returned (no connector created)
- AND no error or panic occurs (gracefully handles nil input)
- AND the slide remains unchanged

#### Scenario: ConnectShapes creates connector in ShapeTree
- GIVEN a Slide with ShapeTree containing 5 shapes
- WHEN connector := slide.ConnectShapes(shape1, shape2, routerType) is called
- THEN the ShapeTree contains 6 elements (5 shapes + 1 connector)
- AND the connector is accessible via ConnectionShapes() method
- AND the connector is persisted when slide is saved

### Requirement: Connector Path Recalculation

The system SHALL provide methods to recalculate connector paths when connected shapes move.

#### Scenario: Update all connectors on slide
- GIVEN a Slide with 3 shapes and 2 connectors
- AND connector1 connects shape1 to shape2
- AND connector2 connects shape2 to shape3
- WHEN shape2 is moved to new position
- AND slide.UpdateAllConnectors() is called
- THEN connector1's path is recalculated to connect shape1 to new shape2 position
- AND connector2's path is recalculated to connect new shape2 position to shape3
- AND both connectors maintain their original RouterType

#### Scenario: UpdateAllConnectors finds all connectors
- GIVEN a Slide with mixed elements (shapes, pictures, connectors, groups)
- AND 5 ConnectionShape elements present
- WHEN slide.UpdateAllConnectors() is called
- THEN all 5 ConnectionShape elements are discovered
- AND each ConnectionShape's path is recalculated
- AND non-connector elements are not affected

#### Scenario: UpdateAllConnectors with no connectors
- GIVEN a Slide with shapes but no ConnectionShape elements
- WHEN slide.UpdateAllConnectors() is called
- THEN no errors occur
- AND the method completes successfully
- AND no elements are modified

#### Scenario: Connector recalculation uses original router type
- GIVEN a connector created with RouterTypeElbow
- WHEN the connected shapes move
- AND slide.UpdateAllConnectors() is called
- THEN the connector's path is recalculated using ElbowRouter
- AND the RouterType is not changed to a different algorithm

#### Scenario: Connector recalculation with deleted shape
- GIVEN a connector connecting shape1 to shape2
- AND shape2 is deleted from the slide
- WHEN slide.UpdateAllConnectors() is called
- THEN the connector's end connection reference is invalid
- AND the connector is skipped or marked as orphaned
- AND no error or panic occurs
- AND other valid connectors are updated normally

#### Scenario: Connector recalculation with disconnected connector
- GIVEN a ConnectionShape with no start or end connection
- WHEN slide.UpdateAllConnectors() is called
- THEN the disconnected connector is skipped
- AND no error or panic occurs
- AND connected connectors are updated normally

#### Scenario: UpdateAllConnectors performance
- GIVEN a Slide with 100 shapes and 50 connectors
- WHEN slide.UpdateAllConnectors() is called
- THEN all 50 connectors are recalculated
- AND the operation completes in reasonable time (<500ms for typical slide)
- AND memory allocations are minimized

### Requirement: Connected Shape Discovery

The system SHALL provide methods to discover connectors attached to shapes.

#### Scenario: Get connectors touching a shape
- GIVEN a Slide with shape1, shape2, shape3
- AND connector1 connects shape1 to shape2
- AND connector2 connects shape2 to shape3
- WHEN connectors := slide.GetConnectedShapes(shape2) is called
- THEN connectors contains connector1 and connector2
- AND connectors slice has length 2
- AND both connectors reference shape2 in either start or end connection

#### Scenario: Get connectors for shape with no connections
- GIVEN a Slide with shape1 that has no connectors attached
- WHEN connectors := slide.GetConnectedShapes(shape1) is called
- THEN connectors is an empty slice
- AND no error occurs

#### Scenario: Get connectors includes start and end connections
- GIVEN a Slide with shape1
- AND connector1 has start connection to shape1
- AND connector2 has end connection to shape1
- AND connector3 connects other shapes (not shape1)
- WHEN connectors := slide.GetConnectedShapes(shape1) is called
- THEN connectors contains connector1 and connector2
- AND connectors does not contain connector3
- AND both start and end connections are considered

#### Scenario: GetConnectedShapes with nil shape
- GIVEN a Slide
- WHEN connectors := slide.GetConnectedShapes(nil) is called
- THEN an empty slice is returned or panic occurs
- AND the method handles nil input gracefully

#### Scenario: GetConnectedShapes by shape ID
- GIVEN a Slide with shape1 having ID "shape1"
- AND connector with start connection to shape ID "shape1"
- WHEN connectors := slide.GetConnectedShapes(shape1) is called
- THEN the connector is found by matching shape ID
- AND the connector is included in returned slice

### Requirement: Routing Context Generation

The system SHALL generate RoutingContext from slide state for routing algorithms.

#### Scenario: Generate routing context for ConnectShapes
- GIVEN a Slide with SlideSize width=9144000, height=6858000
- AND 10 shapes on the slide
- WHEN slide.ConnectShapes(shape1, shape2, routerType) is called
- THEN a RoutingContext is created internally
- AND RoutingContext.StartShape is shape1
- AND RoutingContext.EndShape is shape2
- AND RoutingContext.Obstacles contains other shapes (excluding shape1 and shape2)
- AND RoutingContext.SlideWidth equals 9144000
- AND RoutingContext.SlideHeight equals 6858000

#### Scenario: Routing context excludes connected shapes from obstacles
- GIVEN a Slide with 5 shapes
- WHEN routing context is generated for connecting shape1 to shape2
- THEN Obstacles contains shape3, shape4, shape5
- AND Obstacles does not contain shape1 or shape2
- AND router can generate path without treating endpoints as obstacles

#### Scenario: Routing context includes all slide shapes as obstacles
- GIVEN a Slide with 20 shapes
- WHEN routing context is generated for connecting shape1 to shape2
- THEN Obstacles contains 18 shapes (20 - 2 connected shapes)
- AND all shapes are considered for obstacle avoidance (ElbowRouter)
- AND shapes include regular shapes, pictures, groups, etc.

#### Scenario: Routing context with no obstacles
- GIVEN a Slide with only 2 shapes
- WHEN routing context is generated for connecting the 2 shapes
- THEN Obstacles is an empty slice
- AND routers operate without obstacle avoidance

### Requirement: Connection Site Selection Algorithm

The system SHALL automatically select optimal connection sites when connecting shapes.

#### Scenario: Select nearest connection sites
- GIVEN shape1 with 4 default connection sites (top, right, bottom, left)
- AND shape2 with 4 default connection sites
- AND shape1 is positioned to the left of shape2
- WHEN nearest sites are selected for connecting shape1 to shape2
- THEN shape1's right connection site (index 1) is selected
- AND shape2's left connection site (index 3) is selected
- AND sites minimize distance between connection points

#### Scenario: Connection site distance calculation
- GIVEN shape1 at position (100, 100) with extent (200, 100)
- AND shape1's right connection site at (300, 150)
- AND shape2 at position (500, 200) with extent (200, 100)
- AND shape2's left connection site at (500, 250)
- WHEN distance between sites is calculated
- THEN distance equals sqrt((500-300)^2 + (250-150)^2) = sqrt(40000 + 10000) = 223.6 EMU

#### Scenario: Select sites for all combinations
- GIVEN shape1 with N connection sites
- AND shape2 with M connection sites
- WHEN optimal sites are selected
- THEN all N*M combinations are evaluated
- AND the combination with minimum distance is chosen
- AND algorithm complexity is O(N*M)

#### Scenario: Connection site selection with custom sites
- GIVEN shape1 with 4 default sites plus 2 custom sites
- AND shape2 with 4 default sites
- WHEN nearest sites are selected
- THEN all 6 sites on shape1 are considered
- AND custom site may be selected if it's nearest
- AND default and custom sites are treated equally

#### Scenario: Connection site selection for vertical alignment
- GIVEN shape1 directly above shape2 (same X position)
- WHEN nearest sites are selected
- THEN shape1's bottom connection site is selected
- AND shape2's top connection site is selected
- AND resulting connector is vertically aligned

#### Scenario: Connection site selection for diagonal alignment
- GIVEN shape1 at (100, 100) and shape2 at (500, 500) (diagonal)
- WHEN nearest sites are selected
- THEN sites are selected to minimize distance
- AND selected sites may not be directly facing each other
- AND algorithm handles diagonal cases correctly

### Requirement: Router Type Persistence

The system SHALL persist router type information for connectors to enable path recalculation.

#### Scenario: Store router type in ConnectionShape
- GIVEN a ConnectionShape created with RouterTypeElbow
- WHEN the connector is created via ConnectShapes
- THEN the RouterType is stored in the ConnectionShape
- AND subsequent calls to UpdateAllConnectors use the same RouterType

#### Scenario: Retrieve router type for recalculation
- GIVEN a ConnectionShape with stored RouterType of RouterTypeCurved
- WHEN slide.UpdateAllConnectors() is called
- THEN the system retrieves RouterTypeCurved from the connector
- AND uses CurvedRouter for path recalculation
- AND the recalculated path maintains curved routing

#### Scenario: Router type defaults to Straight if not specified
- GIVEN a ConnectionShape with no RouterType specified (legacy connector)
- WHEN slide.UpdateAllConnectors() is called
- THEN RouterTypeStraight is used as default
- AND the connector path is recalculated as straight line

#### Scenario: Router type serialization
- GIVEN a ConnectionShape with RouterTypeElbow
- WHEN the presentation is saved to file
- AND the file is opened
- THEN the ConnectionShape's RouterType is preserved
- AND UpdateAllConnectors uses ElbowRouter for the connector

### Requirement: Slide Bounds Access

The system SHALL provide access to slide dimensions for routing context generation.

#### Scenario: Access slide width and height
- GIVEN a Slide with slide size element specifying cx=9144000, cy=6858000
- WHEN slide width and height are accessed for routing context
- THEN SlideWidth equals 9144000 EMU (10 inches)
- AND SlideHeight equals 6858000 EMU (7.5 inches)
- AND values match slide size definition

#### Scenario: Default slide size if not specified
- GIVEN a Slide with no slide size element
- WHEN slide dimensions are accessed
- THEN default dimensions are used (typically 10" x 7.5" = 9144000 x 6858000 EMU)
- AND routing context uses default bounds

#### Scenario: Custom slide size
- GIVEN a Slide with custom size element specifying cx=12192000, cy=9144000 (13.33" x 10")
- WHEN routing context is generated
- THEN RoutingContext.SlideWidth equals 12192000
- AND RoutingContext.SlideHeight equals 9144000
- AND routers use correct slide bounds for calculations
