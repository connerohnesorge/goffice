# Presentation Elements Spec Delta

## ADDED Requirements

### Requirement: Shape Connection Sites

The system SHALL provide connection site management for shapes, defining attachment points where connectors can connect.

#### Scenario: Get default connection sites for rectangle
- GIVEN a Shape element representing a rectangle
- WHEN ConnectionSites() is called
- THEN a slice of 4 ConnectionSite structs is returned
- AND ConnectionSite at index 0 has position at top center with angle 270° (pointing up)
- AND ConnectionSite at index 1 has position at right center with angle 0° (pointing right)
- AND ConnectionSite at index 2 has position at bottom center with angle 90° (pointing down)
- AND ConnectionSite at index 3 has position at left center with angle 180° (pointing left)
- AND all positions are in EMU relative to shape bounds

#### Scenario: Add custom connection site
- GIVEN a Shape element
- WHEN AddConnectionSite(x=100*EMUPerPoint, y=50*EMUPerPoint, angle=45*60000) is called
- THEN a ConnectionSite struct is returned
- AND the ConnectionSite has position (100*EMUPerPoint, 50*EMUPerPoint)
- AND the ConnectionSite has angle 45*60000 (45 degrees in 1/60000 units)
- AND subsequent ConnectionSites() call includes the custom site
- AND the custom site has index 4 (after 4 default sites)

#### Scenario: Get specific connection site by index
- GIVEN a Shape with default connection sites
- WHEN GetConnectionSite(index=2) is called
- THEN a ConnectionSite for bottom center is returned
- AND the ConnectionSite is non-nil
- AND the ConnectionSite has position at bottom center of shape bounds
- AND the ConnectionSite has angle 90*60000

#### Scenario: Get connection site with invalid index
- GIVEN a Shape with 4 default connection sites
- WHEN GetConnectionSite(index=10) is called
- THEN nil is returned
- AND no error or panic occurs

#### Scenario: ConnectionSite position calculation
- GIVEN a Shape with offset (100, 200) and extent (400, 300)
- AND a ConnectionSite at index 0 (top center)
- WHEN the ConnectionSite position is accessed
- THEN ConnectionSite.Position.X equals extent width / 2 = 200 (relative to shape origin)
- AND ConnectionSite.Position.Y equals 0 (relative to shape origin, top edge)
- AND position is in shape-relative coordinates (NOT absolute)
- AND absolute position can be calculated as: absolute = shape.Offset + site.Position

#### Scenario: ConnectionSite angle semantics
- GIVEN a ConnectionSite with angle 0
- THEN the angle represents 0 degrees (pointing right)
- AND angle is stored in 1/60000 degree units per ECMA-376
- AND angle 5400000 equals 90 degrees
- AND angle represents direction normal to shape edge (outward direction)

### Requirement: ConnectionShape Connection Management

The system SHALL provide APIs to manage start and end connections of ConnectionShape elements to Shape elements.

#### Scenario: Get start connection
- GIVEN a ConnectionShape with cNvCxnSpPr/stCxn element referencing shape ID "shape1" at connection site index 2
- WHEN shapeId, siteIndex, ok := StartConnection() is called
- THEN shapeId equals "shape1"
- AND siteIndex equals 2
- AND ok equals true

#### Scenario: Get start connection when not connected
- GIVEN a ConnectionShape with no stCxn element
- WHEN shapeId, siteIndex, ok := StartConnection() is called
- THEN shapeId equals ""
- AND siteIndex equals 0
- AND ok equals false

#### Scenario: Set start connection
- GIVEN a ConnectionShape element
- AND a Shape element with ID "targetShape"
- WHEN SetStartConnection(shapeId="targetShape", connectionSiteIndex=1) is called
- THEN cNvCxnSpPr/stCxn element is created or updated
- AND stCxn/@id attribute equals "targetShape"
- AND stCxn/@idx attribute equals 1
- AND subsequent StartConnection() returns ("targetShape", 1, true)

#### Scenario: Update existing start connection
- GIVEN a ConnectionShape with stCxn referencing shape "shape1" at index 0
- WHEN SetStartConnection(shapeId="shape2", connectionSiteIndex=3) is called
- THEN stCxn/@id attribute is updated to "shape2"
- AND stCxn/@idx attribute is updated to 3
- AND subsequent StartConnection() returns ("shape2", 3, true)

#### Scenario: Get end connection
- GIVEN a ConnectionShape with cNvCxnSpPr/endCxn element referencing shape ID "shape2" at connection site index 3
- WHEN shapeId, siteIndex, ok := EndConnection() is called
- THEN shapeId equals "shape2"
- AND siteIndex equals 3
- AND ok equals true

#### Scenario: Get end connection when not connected
- GIVEN a ConnectionShape with no endCxn element
- WHEN shapeId, siteIndex, ok := EndConnection() is called
- THEN shapeId equals ""
- AND siteIndex equals 0
- AND ok equals false

#### Scenario: Set end connection
- GIVEN a ConnectionShape element
- AND a Shape element with ID "targetShape"
- WHEN SetEndConnection(shapeId="targetShape", connectionSiteIndex=2) is called
- THEN cNvCxnSpPr/endCxn element is created or updated
- AND endCxn/@id attribute equals "targetShape"
- AND endCxn/@idx attribute equals 2
- AND subsequent EndConnection() returns ("targetShape", 2, true)

#### Scenario: Clear start connection
- GIVEN a ConnectionShape with stCxn element present
- WHEN ClearStartConnection() is called
- THEN stCxn element is removed from cNvCxnSpPr
- AND subsequent StartConnection() returns ("", 0, false)

#### Scenario: Clear end connection
- GIVEN a ConnectionShape with endCxn element present
- WHEN ClearEndConnection() is called
- THEN endCxn element is removed from cNvCxnSpPr
- AND subsequent EndConnection() returns ("", 0, false)

#### Scenario: Clear connection when already cleared
- GIVEN a ConnectionShape with no stCxn element
- WHEN ClearStartConnection() is called
- THEN no error or panic occurs
- AND cNvCxnSpPr element structure remains valid

#### Scenario: Connection persistence across save/load
- GIVEN a ConnectionShape with start connection to "shape1" index 0 and end connection to "shape2" index 2
- WHEN the presentation is saved to file
- AND the file is opened
- THEN the ConnectionShape's StartConnection() returns ("shape1", 0, true)
- AND the ConnectionShape's EndConnection() returns ("shape2", 2, true)
- AND connection references are preserved in XML

#### Scenario: ConnectionShape with one-sided connection
- GIVEN a ConnectionShape element
- WHEN SetStartConnection(shapeId="shape1", connectionSiteIndex=0) is called
- AND SetEndConnection is not called
- THEN StartConnection() returns ("shape1", 0, true)
- AND EndConnection() returns ("", 0, false)
- AND the ConnectionShape is valid (one-sided connectors are allowed)

### Requirement: ConnectionSite Data Structure

The system SHALL provide a ConnectionSite struct representing connection points on shapes.

#### Scenario: ConnectionSite struct fields
- GIVEN a ConnectionSite instance
- THEN the struct has field Position of type (X, Y EMU)
- AND the struct has field Angle of type int32 (in 1/60000 degree units)
- AND the struct has field Index of type int (0-based identifier)

#### Scenario: ConnectionSite position relative to shape
- GIVEN a Shape with offset (0, 0) and extent (914400, 914400) (1 inch square)
- AND a ConnectionSite at top center (index 0)
- WHEN ConnectionSite.Position is accessed
- THEN Position.X equals 914400 / 2 = 457200 (0.5 inches from left)
- AND Position.Y equals 0 (top edge)
- AND position is in shape-relative coordinates

#### Scenario: ConnectionSite angle conversion
- GIVEN a ConnectionSite with Angle = 5400000
- THEN the angle represents 90 degrees (5400000 / 60000 = 90)
- AND the angle points downward (90 degrees from right)
- AND angle is used by routing algorithms to calculate tangent directions

#### Scenario: Default connection sites for rectangle shape
- GIVEN a rectangular Shape with any offset and extent
- WHEN ConnectionSites() is called
- THEN index 0 site is at (width/2, 0) with angle 270° (top center, pointing up)
- AND index 1 site is at (width, height/2) with angle 0° (right center, pointing right)
- AND index 2 site is at (width/2, height) with angle 90° (bottom center, pointing down)
- AND index 3 site is at (0, height/2) with angle 180° (left center, pointing left)

### Requirement: cNvCxnSpPr XML Element Manipulation

The system SHALL manipulate cNvCxnSpPr (connection shape non-visual properties) element and its child elements per ECMA-376 §21.3.2.36.

#### Scenario: Access cNvCxnSpPr element
- GIVEN a ConnectionShape element
- WHEN the cNvCxnSpPr child element is accessed
- THEN a cNvCxnSpPr element is returned or created if not present
- AND the element conforms to CT_ConnectionShapeProperties schema

#### Scenario: Create stCxn element
- GIVEN a cNvCxnSpPr element with no stCxn child
- WHEN SetStartConnection(shapeId="shape1", connectionSiteIndex=0) is called
- THEN a stCxn child element is created
- AND stCxn element conforms to CT_Connection schema
- AND stCxn has @id attribute with value "shape1"
- AND stCxn has @idx attribute with value 0

#### Scenario: Update stCxn attributes
- GIVEN a cNvCxnSpPr element with existing stCxn element
- WHEN SetStartConnection(shapeId="newShape", connectionSiteIndex=2) is called
- THEN the existing stCxn element is updated (not replaced)
- AND stCxn/@id is set to "newShape"
- AND stCxn/@idx is set to 2

#### Scenario: Remove stCxn element
- GIVEN a cNvCxnSpPr element with stCxn child
- WHEN ClearStartConnection() is called
- THEN the stCxn child element is removed from cNvCxnSpPr
- AND cNvCxnSpPr has no stCxn child
- AND XML structure remains valid

#### Scenario: stCxn and endCxn independence
- GIVEN a cNvCxnSpPr element
- WHEN SetStartConnection is called
- AND SetEndConnection is called
- THEN both stCxn and endCxn child elements exist
- AND modifying stCxn does not affect endCxn
- AND clearing stCxn does not affect endCxn
- AND both connections can be set, cleared, or updated independently
