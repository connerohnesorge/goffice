# Presentation Routing Spec Delta

## ADDED Requirements

### Requirement: Router Interface

The system SHALL provide a Router interface for implementing connector routing algorithms.

#### Scenario: Router interface definition
- GIVEN the presentation/routing package
- WHEN the Router interface is defined
- THEN the interface has method Route(start, end Point, context *RoutingContext) (Path, error)
- AND the interface allows multiple routing algorithm implementations
- AND implementations can be swapped without changing caller code

#### Scenario: Router receives routing context
- GIVEN a Router implementation
- WHEN Route(start, end, context) is called
- THEN the Router has access to context.StartShape
- AND the Router has access to context.EndShape
- AND the Router has access to context.Obstacles (shapes to avoid)
- AND the Router has access to context.SlideWidth and SlideHeight
- AND the Router can use context information for intelligent routing

### Requirement: RoutingContext Data Structure

The system SHALL provide a RoutingContext struct containing information needed for routing algorithms.

#### Scenario: RoutingContext struct fields
- GIVEN a RoutingContext instance
- THEN the struct has field StartShape of type *Shape
- AND the struct has field EndShape of type *Shape
- AND the struct has field Obstacles of type []*Shape
- AND the struct has field SlideWidth of type EMU
- AND the struct has field SlideHeight of type EMU

#### Scenario: RoutingContext with obstacles
- GIVEN a RoutingContext with Obstacles containing 3 Shape elements
- WHEN a Router uses the context
- THEN the Router can access all 3 obstacle shapes
- AND the Router can calculate obstacle bounding boxes
- AND the Router can avoid overlapping obstacles during path calculation

#### Scenario: RoutingContext without obstacles
- GIVEN a RoutingContext with empty Obstacles slice
- WHEN a Router uses the context
- THEN the Router operates without obstacle avoidance
- AND the Router can generate direct paths without constraint

### Requirement: Straight Line Router

The system SHALL provide a StraightRouter that creates direct line paths between connection points.

#### Scenario: Route straight line
- GIVEN a StraightRouter instance
- AND start point at (100, 200) EMU
- AND end point at (500, 600) EMU
- WHEN Route(start, end, context) is called
- THEN a Path with 2 commands is returned
- AND command 1 is MoveTo(100, 200)
- AND command 2 is LineTo(500, 600)
- AND no error is returned

#### Scenario: Straight routing ignores obstacles
- GIVEN a StraightRouter instance
- AND a RoutingContext with obstacles between start and end points
- WHEN Route(start, end, context) is called
- THEN a direct line path from start to end is returned
- AND the path may pass through obstacles (straight router does not avoid)
- AND no error is returned

#### Scenario: Straight routing with coincident points
- GIVEN a StraightRouter instance
- AND start point at (100, 200)
- AND end point at (100, 200) (same as start)
- WHEN Route(start, end, context) is called
- THEN a Path with 1 command is returned
- AND command is MoveTo(100, 200)
- AND no LineTo command is added (zero-length line)
- AND no error is returned

#### Scenario: Straight routing performance
- GIVEN a StraightRouter instance
- WHEN Route is called
- THEN the algorithm completes in O(1) time
- AND no iterative calculations are performed
- AND memory allocation is minimal (2 path commands maximum)

### Requirement: Elbow Router (Orthogonal Routing)

The system SHALL provide an ElbowRouter that creates orthogonal paths with right-angle bends using A* pathfinding.

#### Scenario: Route elbow path without obstacles
- GIVEN an ElbowRouter instance
- AND start point at (100, 100)
- AND end point at (500, 300)
- AND RoutingContext with no obstacles
- WHEN Route(start, end, context) is called
- THEN a Path with orthogonal segments is returned
- AND all path segments are either horizontal or vertical
- AND the path connects start to end
- AND no error is returned

#### Scenario: Route elbow path avoiding obstacles
- GIVEN an ElbowRouter instance
- AND start point at (100, 100)
- AND end point at (500, 500)
- AND RoutingContext with obstacle shape at (250, 250) with size (100, 100)
- WHEN Route(start, end, context) is called
- THEN a Path is returned
- AND the path does not overlap the obstacle shape bounding box
- AND all path segments are orthogonal
- AND the path connects start to end
- AND no error is returned

#### Scenario: Elbow routing uses A* pathfinding
- GIVEN an ElbowRouter instance
- WHEN Route is called with obstacles
- THEN the algorithm uses A* graph search
- AND the algorithm uses Manhattan distance heuristic
- AND the algorithm guarantees shortest orthogonal path (given heuristic)
- AND the algorithm explores minimum nodes needed to find path

#### Scenario: Elbow router grid generation
- GIVEN an ElbowRouter instance
- AND RoutingContext with SlideWidth=9144000 and SlideHeight=6858000 (10x7.5 inches)
- WHEN Route is called
- THEN a grid is generated for pathfinding
- AND grid cell size is configurable (default 127000 EMU = 10 points)
- AND grid covers entire slide bounds
- AND grid cells are marked as free or blocked based on obstacles

#### Scenario: Elbow router obstacle marking
- GIVEN an ElbowRouter instance
- AND RoutingContext with obstacle at (100, 100) with extent (200, 200)
- AND grid with 10 EMU cell size
- WHEN the grid is generated
- THEN grid cells overlapping obstacle bounding box are marked as blocked
- AND grid cells outside obstacle are marked as free
- AND start and end cells are always marked as free (even if in obstacle)

#### Scenario: Elbow router Manhattan heuristic
- GIVEN an ElbowRouter implementing A* pathfinding
- AND grid cell A at (10, 20)
- AND grid cell B at (15, 25)
- WHEN the heuristic function h(A, B) is calculated
- THEN h(A, B) equals |15-10| + |25-20| = 5 + 5 = 10
- AND the heuristic is admissible (never overestimates actual cost)
- AND the heuristic is consistent (satisfies triangle inequality)

#### Scenario: Elbow router A* search
- GIVEN an ElbowRouter instance
- WHEN A* search is performed from start to end
- THEN the algorithm maintains openSet priority queue
- AND the algorithm maintains closedSet of explored nodes
- AND the algorithm maintains gScore (cost from start)
- AND the algorithm maintains fScore (gScore + heuristic)
- AND the algorithm explores node with minimum fScore first
- AND the algorithm terminates when end node is reached

#### Scenario: Elbow router path reconstruction
- GIVEN an ElbowRouter that completed A* search
- AND A* search reached end node
- WHEN path is reconstructed from cameFrom map
- THEN the path is traced backward from end to start
- AND the path is reversed to start → end order
- AND the path contains only cells explored by A*
- AND the path is converted from grid cells to EMU coordinates

#### Scenario: Elbow routing with no path available
- GIVEN an ElbowRouter instance
- AND start point completely surrounded by obstacles
- AND end point unreachable
- WHEN Route(start, end, context) is called
- THEN an error is returned
- AND error message indicates no path found
- AND returned Path is nil or empty

#### Scenario: Elbow router margin configuration
- GIVEN an ElbowRouter with Margin=10000 EMU
- AND RoutingContext with obstacle shape
- WHEN grid is generated
- THEN obstacle bounding box is expanded by Margin in all directions
- AND grid cells within expanded box are marked as blocked
- AND resulting path maintains at least 10000 EMU distance from obstacles

#### Scenario: Elbow routing performance with obstacles
- GIVEN an ElbowRouter instance
- AND RoutingContext with N obstacles
- WHEN Route is called
- THEN algorithm complexity is O(b^d) where b=branching factor (4 for grid), d=path depth
- AND obstacle marking is O(N * cells per obstacle)
- AND for typical slides (100x100 grid, 10 obstacles), routing completes in <100ms

### Requirement: Curved Router (Bezier Routing)

The system SHALL provide a CurvedRouter that creates smooth Bezier curve paths between connection points.

#### Scenario: Route Bezier curve
- GIVEN a CurvedRouter instance
- AND start point at (100, 100) with connection site angle 0° (pointing right)
- AND end point at (500, 500) with connection site angle 180° (pointing left)
- WHEN Route(start, end, context) is called
- THEN a Path with cubic Bezier curve is returned
- AND path has MoveTo(100, 100) command
- AND path has CubicBezierTo command with 2 control points
- AND Bezier curve endpoints match start and end points
- AND no error is returned

#### Scenario: Bezier control point calculation
- GIVEN a CurvedRouter instance
- AND start point at (100, 100) with angle 0° (pointing right)
- AND end point at (500, 500) with angle 180° (pointing left)
- AND distance between points is 400√2 ≈ 565.7 EMU
- WHEN control points are calculated
- THEN control point 1 is offset from start in direction of start angle
- AND control point 2 is offset from end in direction of end angle (reversed)
- AND control point distance is approximately 33% of total distance (≈ 187.9 EMU)
- AND control points create smooth S-curve or C-curve

#### Scenario: Bezier curve with same-direction angles
- GIVEN a CurvedRouter instance
- AND start point with angle 0° (pointing right)
- AND end point with angle 0° (pointing right)
- WHEN Route is called
- THEN control points are calculated along rightward direction
- AND resulting Bezier curve bends to connect points smoothly
- AND curve shape is C-curve (same-side control points)

#### Scenario: Bezier curve with opposite-direction angles
- GIVEN a CurvedRouter instance
- AND start point with angle 0° (pointing right)
- AND end point with angle 180° (pointing left)
- WHEN Route is called
- THEN control point 1 is to the right of start
- AND control point 2 is to the left of end
- AND resulting Bezier curve is S-curve (opposite-side control points)

#### Scenario: Curved routing ignores obstacles
- GIVEN a CurvedRouter instance
- AND RoutingContext with obstacles between start and end
- WHEN Route is called
- THEN Bezier curve is calculated without obstacle avoidance
- AND curve may pass through obstacles (curved router does not avoid)
- AND no error is returned

#### Scenario: Bezier curve smoothness
- GIVEN a CurvedRouter instance
- WHEN a Bezier curve is generated
- THEN the curve is C2 continuous (continuous second derivative)
- AND the curve has no cusps or sharp angles
- AND the curve appears visually smooth when rendered

#### Scenario: Bezier control point distance factor
- GIVEN a CurvedRouter instance with default configuration
- WHEN control point distance is calculated
- THEN distance factor is 0.33 (33% of total distance between points)
- AND factor is configurable in CurvedRouter struct
- AND factor values between 0.2 and 0.5 produce aesthetically pleasing curves

#### Scenario: Curved routing with coincident points
- GIVEN a CurvedRouter instance
- AND start point at (100, 100)
- AND end point at (100, 100) (same as start)
- WHEN Route is called
- THEN a Path with MoveTo(100, 100) is returned
- AND no Bezier curve is added (zero-length curve)
- AND no error is returned

#### Scenario: Bezier curve with perpendicular angles
- GIVEN a CurvedRouter instance
- AND start point with angle 0° (pointing right)
- AND end point with angle 90° (pointing down)
- WHEN Route is called
- THEN control point 1 extends rightward from start
- AND control point 2 extends upward from end (opposite of 90°)
- AND resulting curve smoothly transitions from horizontal to vertical direction

### Requirement: Path Data Structure

The system SHALL provide a Path struct representing connector paths as sequences of drawing commands.

#### Scenario: Path struct fields
- GIVEN a Path instance
- THEN the struct has field Commands of type []PathCommand
- AND PathCommand is an interface or union type
- AND PathCommand includes MoveTo, LineTo, CubicBezierTo command types

#### Scenario: Path with MoveTo and LineTo
- GIVEN a Path instance
- WHEN Commands contains [MoveTo(100, 100), LineTo(200, 200)]
- THEN the path represents a straight line from (100, 100) to (200, 200)
- AND the path can be rendered as a line segment

#### Scenario: Path with Bezier curve
- GIVEN a Path instance
- WHEN Commands contains [MoveTo(100, 100), CubicBezierTo(150, 100, 150, 200, 200, 200)]
- THEN the path represents a Bezier curve from (100, 100) to (200, 200)
- AND control point 1 is at (150, 100)
- AND control point 2 is at (150, 200)

#### Scenario: Path command sequence validation
- GIVEN a Path instance
- WHEN the first command is not MoveTo
- THEN the path is invalid
- AND rendering may fail or produce incorrect output
- AND path construction should ensure MoveTo is first command

#### Scenario: Empty path
- GIVEN a Path instance with empty Commands slice
- THEN the path represents no geometry
- AND rendering produces no output
- AND the path is valid but has no visual representation

### Requirement: Point Data Structure

The system SHALL provide a Point struct representing 2D coordinates in EMU.

#### Scenario: Point struct fields
- GIVEN a Point instance
- THEN the struct has field X of type EMU
- AND the struct has field Y of type EMU
- AND EMU is an integer type representing English Metric Units (914400 per inch)

#### Scenario: Point arithmetic
- GIVEN Point p1 at (100, 200)
- AND Point p2 at (300, 400)
- WHEN distance between points is calculated
- THEN distance equals sqrt((300-100)^2 + (400-200)^2) = sqrt(40000 + 40000) = 282.8 EMU

#### Scenario: Point from connection site
- GIVEN a Shape with connection site at index 0
- AND ConnectionSite has position (100, 200) in shape coordinates
- AND Shape has absolute offset (500, 600)
- WHEN Point is calculated for connection
- THEN Point X equals shape offset X + site X = 500 + 100 = 600
- AND Point Y equals shape offset Y + site Y = 600 + 200 = 800
- AND Point is in absolute slide coordinates

### Requirement: RouterType Enumeration

The system SHALL provide a RouterType enumeration identifying routing algorithm types.

#### Scenario: RouterType values
- GIVEN the RouterType enumeration
- THEN RouterType has value RouterTypeStraight
- AND RouterType has value RouterTypeElbow
- AND RouterType has value RouterTypeCurved

#### Scenario: RouterType selection
- GIVEN a RouterType value of RouterTypeElbow
- WHEN selecting a Router implementation
- THEN an ElbowRouter instance is created
- AND the ElbowRouter implements the Router interface
- AND Route calls use elbow (orthogonal) routing algorithm

#### Scenario: RouterType string representation
- GIVEN RouterType value RouterTypeStraight
- WHEN converted to string
- THEN the string is "straight" or "Straight"
- AND the string can be used for serialization or user display

### Requirement: Path Geometry Utilities

The system SHALL provide utility methods on Path for geometric calculations.

#### Scenario: Calculate path bounds
- GIVEN a Path with commands [MoveTo(100, 100), LineTo(500, 300), LineTo(200, 600)]
- WHEN Bounds() is called
- THEN minX equals 100, maxX equals 500
- AND minY equals 100, maxY equals 600
- AND bounding box encompasses all path points

#### Scenario: Calculate path length for straight segments
- GIVEN a Path with commands [MoveTo(0, 0), LineTo(300, 400)]
- WHEN Length() is called
- THEN length equals sqrt(300^2 + 400^2) = 500 EMU
- AND length is sum of all segment lengths

#### Scenario: Calculate path length for Bezier curve
- GIVEN a Path with cubic Bezier curve
- WHEN Length() is called
- THEN length is approximated using recursive subdivision
- AND subdivision continues until segments are nearly straight
- AND approximation error is less than 1 EMU

#### Scenario: Path bounds with Bezier curve
- GIVEN a Path with CubicBezierTo(cp1x, cp1y, cp2x, cp2y, x, y)
- WHEN Bounds() is called
- THEN bounding box encompasses start point, end point, and control points
- AND bounding box may be larger than actual curve bounds (conservative estimate)
- OR bounding box is calculated by sampling curve at multiple t values (accurate estimate)
