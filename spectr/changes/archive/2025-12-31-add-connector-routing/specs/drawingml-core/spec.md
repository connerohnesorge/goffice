# DrawingML Core Spec Delta

## ADDED Requirements

### Requirement: Path Construction Methods

The system SHALL provide methods to construct paths by appending drawing commands.

#### Scenario: Add line segment to path
- GIVEN a Path instance with MoveTo(100, 100) command
- WHEN AddLine(x=200, y=300) is called
- THEN a LineTo(200, 300) command is appended to path
- AND the path has 2 commands total
- AND the line segment connects (100, 100) to (200, 300)

#### Scenario: Add multiple line segments
- GIVEN a Path instance with MoveTo(0, 0)
- WHEN AddLine(100, 0) is called
- AND AddLine(100, 100) is called
- AND AddLine(0, 100) is called
- THEN the path has 4 commands (1 MoveTo + 3 LineTo)
- AND the path forms a U-shape

#### Scenario: Add cubic Bezier curve to path
- GIVEN a Path instance with MoveTo(100, 100) command
- WHEN AddCubicBezier(cp1x=150, cp1y=100, cp2x=150, cp2y=200, x=200, y=200) is called
- THEN a CubicBezierTo command is appended to path
- AND the command has control point 1 at (150, 100)
- AND the command has control point 2 at (150, 200)
- AND the command has endpoint at (200, 200)

#### Scenario: Add line to empty path
- GIVEN an empty Path instance with no commands
- WHEN AddLine(100, 100) is called
- THEN a LineTo(100, 100) command is added
- AND the path is technically invalid (LineTo without MoveTo)
- OR a MoveTo(0, 0) is implicitly added first

#### Scenario: Construct complex path
- GIVEN an empty Path instance
- WHEN MoveTo(0, 0) is added
- AND AddLine(100, 0) is called
- AND AddCubicBezier(150, 0, 150, 100, 100, 100) is called
- AND AddLine(0, 100) is called
- THEN the path has 4 commands
- AND the path contains mix of straight and curved segments

### Requirement: Path Bounds Calculation

The system SHALL calculate bounding boxes for paths containing lines and curves.

#### Scenario: Calculate bounds for straight line path
- GIVEN a Path with [MoveTo(100, 200), LineTo(500, 600)]
- WHEN Bounds() is called
- THEN minX equals 100, maxX equals 500
- AND minY equals 200, maxY equals 600
- AND bounding box encompasses entire line segment

#### Scenario: Calculate bounds for path with multiple segments
- GIVEN a Path with [MoveTo(100, 100), LineTo(500, 200), LineTo(300, 600)]
- WHEN Bounds() is called
- THEN minX equals 100, maxX equals 500
- AND minY equals 100, maxY equals 600
- AND bounding box is tightest axis-aligned rectangle containing all segments

#### Scenario: Calculate bounds for Bezier curve (conservative)
- GIVEN a Path with CubicBezierTo(cp1x=200, cp1y=100, cp2x=300, cp2y=200, x=400, y=150)
- AND curve starts at (100, 100)
- WHEN Bounds() is called with conservative algorithm
- THEN minX equals min(100, 200, 300, 400) = 100
- AND maxX equals max(100, 200, 300, 400) = 400
- AND minY equals min(100, 100, 200, 150) = 100
- AND maxY equals max(100, 100, 200, 150) = 200
- AND bounding box includes start, end, and both control points

#### Scenario: Calculate bounds for Bezier curve (accurate)
- GIVEN a Path with cubic Bezier curve
- WHEN Bounds() is called with accurate algorithm
- THEN the curve is sampled at multiple t values (e.g., t = 0, 0.1, 0.2, ..., 1.0)
- AND min/max X and Y are calculated from sample points
- AND bounding box tightly fits the actual curve geometry
- AND accuracy improves with more samples

#### Scenario: Calculate bounds for empty path
- GIVEN an empty Path with no commands
- WHEN Bounds() is called
- THEN minX equals 0, maxX equals 0
- AND minY equals 0, maxY equals 0
- OR an error is returned indicating invalid path

#### Scenario: Calculate bounds for single point path
- GIVEN a Path with only MoveTo(100, 200)
- WHEN Bounds() is called
- THEN minX equals 100, maxX equals 100
- AND minY equals 200, maxY equals 200
- AND bounding box is a single point

#### Scenario: Bounds calculation performance
- GIVEN a Path with 100 segments (mix of lines and Bezier curves)
- WHEN Bounds() is called
- THEN calculation completes in O(N) time where N is number of commands
- AND memory allocation is O(1) (no allocations for temporary storage)

### Requirement: Path Length Calculation

The system SHALL calculate total path length for lines and curves.

#### Scenario: Calculate length for straight line
- GIVEN a Path with [MoveTo(0, 0), LineTo(300, 400)]
- WHEN Length() is called
- THEN length equals sqrt(300^2 + 400^2) = 500 EMU
- AND calculation is exact (no approximation needed)

#### Scenario: Calculate length for multiple straight segments
- GIVEN a Path with [MoveTo(0, 0), LineTo(100, 0), LineTo(100, 100), LineTo(0, 100)]
- WHEN Length() is called
- THEN length equals 100 + 100 + 100 = 300 EMU
- AND length is sum of individual segment lengths

#### Scenario: Calculate length for Bezier curve (recursive subdivision)
- GIVEN a Path with cubic Bezier curve from (0, 0) to (100, 100)
- AND control points at (50, 0) and (50, 100)
- WHEN Length() is called
- THEN the curve is subdivided recursively
- AND subdivision continues until segments are nearly straight (straightness threshold)
- AND length is sum of subdivided segment lengths
- AND approximation error is less than 1 EMU

#### Scenario: Bezier length subdivision algorithm
- GIVEN a cubic Bezier curve B(t) = (1-t)^3*P0 + 3(1-t)^2*t*P1 + 3(1-t)*t^2*P2 + t^3*P3
- WHEN Length() subdivision is performed
- THEN curve is split at t=0.5 into two sub-curves B1 and B2
- AND straightness is tested by comparing chord length to control polygon length
- AND if straight enough, chord length is used
- AND if not straight enough, B1 and B2 are recursively subdivided

#### Scenario: Bezier straightness test
- GIVEN a Bezier curve segment with start, end, and control points
- WHEN straightness is tested
- THEN chord length d_chord = distance(start, end)
- AND control polygon length d_polygon = distance(start, cp1) + distance(cp1, cp2) + distance(cp2, end)
- AND curve is "straight enough" if d_polygon / d_chord < 1.01 (1% tolerance)

#### Scenario: Calculate length for mixed path
- GIVEN a Path with [MoveTo(0, 0), LineTo(100, 0), CubicBezierTo(...), LineTo(300, 100)]
- WHEN Length() is called
- THEN line segment lengths are calculated exactly
- AND Bezier curve length is approximated via subdivision
- AND total length is sum of all segment lengths

#### Scenario: Length calculation for empty path
- GIVEN an empty Path with no commands
- WHEN Length() is called
- THEN length equals 0
- AND no error occurs

#### Scenario: Length calculation performance
- GIVEN a Path with 10 Bezier curves
- WHEN Length() is called
- THEN each curve is subdivided up to maximum depth (e.g., 10 levels)
- AND total subdivision creates at most 2^10 = 1024 segments per curve
- AND calculation completes in <10ms for typical paths

### Requirement: Path Command Types

The system SHALL provide distinct types for path drawing commands.

#### Scenario: MoveTo command structure
- GIVEN a MoveTo command instance
- THEN the command has field X of type EMU
- AND the command has field Y of type EMU
- AND the command represents moving drawing cursor without drawing

#### Scenario: LineTo command structure
- GIVEN a LineTo command instance
- THEN the command has field X of type EMU
- AND the command has field Y of type EMU
- AND the command represents drawing straight line from current position to (X, Y)

#### Scenario: CubicBezierTo command structure
- GIVEN a CubicBezierTo command instance
- THEN the command has field CP1X, CP1Y of type EMU (control point 1)
- AND the command has field CP2X, CP2Y of type EMU (control point 2)
- AND the command has field X, Y of type EMU (endpoint)
- AND the command represents cubic Bezier curve from current position to (X, Y)

#### Scenario: PathCommand interface or union
- GIVEN the PathCommand type
- THEN PathCommand is an interface implemented by MoveTo, LineTo, CubicBezierTo
- OR PathCommand is a union/variant type containing one of the command types
- AND Path.Commands slice can contain any PathCommand type

### Requirement: Bezier Curve Utilities

The system SHALL provide utilities for Bezier curve manipulation and calculation.

#### Scenario: Evaluate Bezier curve at parameter t
- GIVEN a cubic Bezier curve with P0=(0,0), P1=(100,0), P2=(100,100), P3=(200,100)
- WHEN Bezier.Evaluate(t=0.5) is called
- THEN the point at t=0.5 is calculated using Bezier formula
- AND result equals (1-0.5)^3*(0,0) + 3*(1-0.5)^2*0.5*(100,0) + 3*(1-0.5)*0.5^2*(100,100) + 0.5^3*(200,100)
- AND result is approximately (100, 37.5)

#### Scenario: Subdivide Bezier curve at t
- GIVEN a cubic Bezier curve B
- WHEN Subdivide(t=0.5) is called
- THEN two new Bezier curves B1 and B2 are returned
- AND B1 represents B(0..0.5)
- AND B2 represents B(0.5..1)
- AND B1.end == B2.start (curves connect at subdivision point)
- AND concatenating B1 and B2 equals original curve B

#### Scenario: Calculate Bezier tangent at t
- GIVEN a cubic Bezier curve B(t)
- WHEN Tangent(t=0) is called
- THEN tangent direction at start (t=0) is calculated
- AND tangent equals 3*(P1 - P0) (derivative of Bezier formula at t=0)
- AND tangent vector points in initial curve direction

#### Scenario: Calculate Bezier tangent at end
- GIVEN a cubic Bezier curve B(t)
- WHEN Tangent(t=1) is called
- THEN tangent direction at end (t=1) is calculated
- AND tangent equals 3*(P3 - P2) (derivative at t=1)
- AND tangent vector points in final curve direction

#### Scenario: Bezier curve control point calculation from constraints
- GIVEN a desired curve from start to end
- AND desired tangent directions at start and end
- AND desired control point distance factor (e.g., 0.33)
- WHEN control points are calculated
- THEN CP1 = start + factor * distance * normalize(start_tangent)
- AND CP2 = end - factor * distance * normalize(end_tangent)
- AND resulting Bezier curve has desired tangent directions

### Requirement: Path Transformation

The system SHALL provide methods to transform paths by matrices.

#### Scenario: Transform path by translation
- GIVEN a Path with [MoveTo(100, 100), LineTo(200, 200)]
- AND a translation matrix Translate(50, 75)
- WHEN Transform(matrix) is called
- THEN MoveTo is transformed to (150, 175)
- AND LineTo is transformed to (250, 275)
- AND path shape is preserved (only position changes)

#### Scenario: Transform path by rotation
- GIVEN a Path with horizontal line [MoveTo(0, 0), LineTo(100, 0)]
- AND a rotation matrix Rotate(90°)
- WHEN Transform(matrix) is called
- THEN the line is rotated to vertical [MoveTo(0, 0), LineTo(0, 100)]
- AND all points are transformed by rotation

#### Scenario: Transform Bezier curve control points
- GIVEN a Path with CubicBezierTo(cp1x, cp1y, cp2x, cp2y, x, y)
- AND a transformation matrix M
- WHEN Transform(M) is called
- THEN start point is transformed by M
- AND control point 1 is transformed by M
- AND control point 2 is transformed by M
- AND endpoint is transformed by M
- AND resulting curve has same shape in new coordinate system

#### Scenario: Transform empty path
- GIVEN an empty Path
- AND a transformation matrix
- WHEN Transform(matrix) is called
- THEN no error occurs
- AND path remains empty

### Requirement: Path Validation

The system SHALL provide validation for path command sequences.

#### Scenario: Validate path starts with MoveTo
- GIVEN a Path with [MoveTo(100, 100), LineTo(200, 200)]
- WHEN IsValid() is called
- THEN the method returns true
- AND the path is considered valid

#### Scenario: Validate path without initial MoveTo
- GIVEN a Path with [LineTo(100, 100), LineTo(200, 200)]
- WHEN IsValid() is called
- THEN the method returns false
- AND error indicates "path must start with MoveTo"

#### Scenario: Validate empty path
- GIVEN an empty Path with no commands
- WHEN IsValid() is called
- THEN the method returns false or true (implementation-dependent)
- AND behavior is documented

#### Scenario: Validate Bezier curve has control points
- GIVEN a Path with CubicBezierTo command
- WHEN IsValid() is called
- THEN the command is checked for valid control point coordinates
- AND if control points are NaN or infinite, path is invalid
- AND if control points are valid, path validation continues

### Requirement: Path Reversal

The system SHALL provide methods to reverse path direction.

#### Scenario: Reverse straight line path
- GIVEN a Path with [MoveTo(100, 100), LineTo(200, 200)]
- WHEN Reverse() is called
- THEN the path becomes [MoveTo(200, 200), LineTo(100, 100)]
- AND path traces same geometry in opposite direction

#### Scenario: Reverse multi-segment path
- GIVEN a Path with [MoveTo(0, 0), LineTo(100, 0), LineTo(100, 100)]
- WHEN Reverse() is called
- THEN the path becomes [MoveTo(100, 100), LineTo(100, 0), LineTo(0, 0)]
- AND path direction is reversed

#### Scenario: Reverse Bezier curve
- GIVEN a Path with [MoveTo(0, 0), CubicBezierTo(50, 0, 50, 100, 100, 100)]
- WHEN Reverse() is called
- THEN the path becomes [MoveTo(100, 100), CubicBezierTo(50, 100, 50, 0, 0, 0)]
- AND control points are reversed (CP2 becomes CP1, CP1 becomes CP2)
- AND curve traces same shape in opposite direction

#### Scenario: Reverse empty path
- GIVEN an empty Path
- WHEN Reverse() is called
- THEN the path remains empty
- AND no error occurs
