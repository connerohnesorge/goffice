# Drawingml Core Specification

## Requirements

### Requirement: Transform2D Element

The system SHALL provide a `Transform2D` element (a:xfrm) for positioning and sizing shapes.

#### Scenario: Basic transform properties
- GIVEN a Transform2D element
- WHEN properties are accessed
- THEN Offset (a:off with x, y in EMUs) is available
- AND Extents (a:ext with cx, cy in EMUs) is available
- AND Rotation (rot in 1/60000 degree units) is available
- AND FlipH and FlipV boolean properties are available

#### Scenario: Rotation conversion
- GIVEN a rotation value of 5400000
- WHEN converted to degrees
- THEN the result is 90 degrees
- AND full rotation is 21600000 (360 degrees)

#### Scenario: EMU to points conversion
- GIVEN dimensions in EMUs
- WHEN converted to points
- THEN 12700 EMUs equals 1 point
- AND 914400 EMUs equals 72 points (1 inch)

### Requirement: Preset Geometry Element

The system SHALL provide a `PresetGeometry` element (a:prstGeom) for built-in shapes.

#### Scenario: Preset shape selection
- GIVEN a PresetGeometry element
- WHEN the prst attribute is set
- THEN values like "rect", "ellipse", "roundRect", "triangle" are supported
- AND over 100 preset shape types are available

#### Scenario: Adjust value list
- GIVEN a PresetGeometry with adjustable parameters
- WHEN AdjustValueList (a:avLst) is accessed
- THEN shape-specific adjustment values can be modified
- AND adjustments control things like corner radius, arrow width

#### Scenario: Common preset shapes
- WHEN preset geometry shapes are enumerated
- THEN the following categories are available:
  - Basic: rect, ellipse, roundRect, triangle, rtTriangle
  - Block arrows: rightArrow, leftArrow, upArrow, downArrow, leftRightArrow
  - Stars and banners: star4, star5, star6, irregularSeal1, ribbon2
  - Callouts: wedgeRoundRectCallout, cloudCallout
  - Flowchart: flowChartProcess, flowChartDecision, flowChartTerminator

### Requirement: Custom Geometry Element

The system SHALL provide a `CustomGeometry` element (a:custGeom) for path-based shapes.

#### Scenario: Custom geometry structure
- GIVEN a CustomGeometry element
- WHEN components are accessed
- THEN PathList (a:pathLst), AdjustValueList (a:avLst), GuideList (a:gdLst), AdjustHandleList (a:ahLst), ConnectionSiteList (a:cxnLst), ShapeTextRectangle (a:rect) are available

#### Scenario: Path commands
- GIVEN a Path element in PathList
- WHEN path commands are accessed
- THEN MoveTo (a:moveTo), LineTo (a:lnTo), ArcTo (a:arcTo), QuadBezTo (a:quadBezTo), CubicBezTo (a:cubicBezTo), Close (a:close) are available

#### Scenario: Guide formulas
- GIVEN a GuideList element
- WHEN Guide elements are accessed
- THEN name, formula (fmla) for calculated values are available
- AND formulas support operations like +, -, *, /, abs, at2, cat2, cos, max, min, mod, pin, sat2, sin, sqrt, tan, val

### Requirement: Solid Fill Element

The system SHALL provide a `SolidFill` element (a:solidFill) for single-color fills.

#### Scenario: RGB color fill
- GIVEN a SolidFill with sRGB color
- WHEN color is accessed
- THEN SrgbClr (a:srgbClr) with val attribute contains 6-digit hex color
- AND color transforms can be applied

#### Scenario: Scheme color fill
- GIVEN a SolidFill with scheme color
- WHEN color is accessed
- THEN SchemeClr (a:schemeClr) with val like "accent1", "dk1", "lt1" is available
- AND the color resolves from document theme

#### Scenario: Color transforms
- GIVEN any color element
- WHEN transforms are applied
- THEN Alpha (a:alpha), Tint (a:tint), Shade (a:shade), SatMod (a:satMod), LumMod (a:lumMod), HueOff (a:hueOff) are available
- AND transform values are in 1/1000 percent (e.g., 50000 = 50%)

### Requirement: Gradient Fill Element

The system SHALL provide a `GradientFill` element (a:gradFill) for gradient fills.

#### Scenario: Linear gradient
- GIVEN a GradientFill with linear type
- WHEN properties are accessed
- THEN GradientStopList (a:gsLst) contains color stops with position
- AND Linear (a:lin) contains angle in 1/60000 degrees
- AND scaled attribute controls scaling behavior

#### Scenario: Path gradient
- GIVEN a GradientFill with path type
- WHEN properties are accessed
- THEN Path (a:path) with path type (circle, rect, shape) is available
- AND FillToRect (a:fillToRect) defines gradient focus

#### Scenario: Gradient stops
- GIVEN a GradientStopList element
- WHEN stops are accessed
- THEN each GradientStop (a:gs) has pos (position 0-100000) and a color element

### Requirement: Pattern Fill Element

The system SHALL provide a `PatternFill` element (a:pattFill) for pattern fills.

#### Scenario: Pattern properties
- GIVEN a PatternFill element
- WHEN properties are accessed
- THEN prst attribute contains pattern type (50+ presets)
- AND ForegroundColor (a:fgClr) and BackgroundColor (a:bgClr) are available

#### Scenario: Common pattern types
- WHEN pattern types are enumerated
- THEN values like pct5, pct10, pct20, pct25, pct30, pct40, pct50, pct60, pct70, pct75, pct80, pct90, horz, vert, ltHorz, ltVert, dkHorz, dkVert, narHorz, narVert, dashHorz, dashVert, cross, dnDiag, upDiag, ltDnDiag, ltUpDiag, dkDnDiag, dkUpDiag, wdDnDiag, wdUpDiag, dashDnDiag, dashUpDiag, diagCross, smCheck, lgCheck, smGrid, lgGrid, dotGrid, smConfetti, lgConfetti, horzBrick, diagBrick, solidDmnd, openDmnd, dotDmnd, plaid, sphere, weave, divot, shingle, wave, trellis, zigZag are available

### Requirement: Blip Fill Element

The system SHALL provide a `BlipFill` element (a:blipFill) for image fills.

#### Scenario: Blip reference
- GIVEN a BlipFill element
- WHEN Blip (a:blip) is accessed
- THEN r:embed contains relationship ID to embedded image
- OR r:link contains relationship ID to linked image
- AND cstate attribute controls compression (none, print, screen, email)

#### Scenario: Fill modes
- GIVEN a BlipFill element
- WHEN fill mode is accessed
- THEN Stretch (a:stretch) with FillRect (a:fillRect) for stretch-to-fill
- OR Tile (a:tile) for tiled pattern with tx, ty, sx, sy, flip, algn

#### Scenario: Source rectangle
- GIVEN a BlipFill element
- WHEN SourceRect (a:srcRect) is accessed
- THEN l, t, r, b attributes define crop region (in 1/1000 percent)

### Requirement: No Fill and Group Fill Elements

The system SHALL provide `NoFill` (a:noFill) and `GroupFill` (a:grpFill) elements.

#### Scenario: No fill
- GIVEN a NoFill element
- WHEN applied to a shape
- THEN the shape has no fill (transparent)

#### Scenario: Group fill
- GIVEN a GroupFill element
- WHEN applied to a shape in a group
- THEN the shape inherits its fill from the parent group

### Requirement: Line Properties Element

The system SHALL provide a `LineProperties` element (a:ln) for shape outlines.

#### Scenario: Line basic properties
- GIVEN a LineProperties element
- WHEN properties are accessed
- THEN w (width in EMUs), cap (flat/round/square), cmpd (single/double/thickThin/thinThick/triple), algn (center/inset) are available

#### Scenario: Line fill
- GIVEN a LineProperties element
- WHEN fill is accessed
- THEN SolidFill, GradientFill, PatternFill, or NoFill is available

#### Scenario: Dash style
- GIVEN a LineProperties element
- WHEN dash style is accessed
- THEN PresetDash (a:prstDash) with val like solid, dash, dot, dashDot, lgDash, lgDashDot, lgDashDotDot, sysDash, sysDot, sysDashDot, sysDashDotDot is available
- OR CustomDash (a:custDash) with dash/space pattern is available

#### Scenario: Line join
- GIVEN a LineProperties element
- WHEN join is accessed
- THEN Round (a:round), Bevel (a:bevel), or Miter (a:miter with lim) is available

#### Scenario: Head and tail end
- GIVEN a LineProperties element
- WHEN end types are accessed
- THEN HeadEnd (a:headEnd) and TailEnd (a:tailEnd) are available
- AND type (none/triangle/stealth/diamond/oval/arrow), w (width), len (length) are available

### Requirement: Effect List Element

The system SHALL provide an `EffectList` element (a:effectLst) for visual effects.

#### Scenario: Outer shadow
- GIVEN an EffectList element
- WHEN OuterShadow (a:outerShdw) is accessed
- THEN blurRad (blur radius), dist (distance), dir (direction in 1/60000 degrees), sx, sy (scale), kx, ky (skew), algn (alignment), rotWithShape are available
- AND a color element defines shadow color

#### Scenario: Inner shadow
- GIVEN an EffectList element
- WHEN InnerShadow (a:innerShdw) is accessed
- THEN blurRad, dist, dir are available
- AND a color element defines shadow color

#### Scenario: Glow effect
- GIVEN an EffectList element
- WHEN Glow (a:glow) is accessed
- THEN rad (radius in EMUs) is available
- AND a color element defines glow color

#### Scenario: Soft edge effect
- GIVEN an EffectList element
- WHEN SoftEdge (a:softEdge) is accessed
- THEN rad (radius in EMUs) is available

#### Scenario: Reflection effect
- GIVEN an EffectList element
- WHEN Reflection (a:reflection) is accessed
- THEN blurRad, stA (start alpha), endA (end alpha), dist, dir, fadeDir, sx, sy, kx, ky, algn, rotWithShape are available

#### Scenario: Blur effect
- GIVEN an EffectList element
- WHEN Blur (a:blur) is accessed
- THEN rad (radius), grow (expand bounds) are available

### Requirement: Shape Properties Element

The system SHALL provide a `ShapeProperties` element (a:spPr) as the main container for shape visual properties.

#### Scenario: Shape properties structure
- GIVEN a ShapeProperties element
- WHEN children are accessed
- THEN Transform2D (a:xfrm), Geometry (PresetGeometry or CustomGeometry), Fill (SolidFill, GradientFill, etc.), Outline (a:ln), EffectList or EffectDag, Scene3D, Shape3D are available

#### Scenario: Black-white mode
- GIVEN a ShapeProperties element
- WHEN bwMode attribute is accessed
- THEN values clr, auto, gray, ltGray, invGray, grayWhite, blackGray, blackWhite, black, white, hidden are available

### Requirement: Shape Style Element

The system SHALL provide a `ShapeStyle` element (a:style) for theme-based styling.

#### Scenario: Style references
- GIVEN a ShapeStyle element
- WHEN references are accessed
- THEN LineReference (a:lnRef), FillReference (a:fillRef), EffectReference (a:effectRef), FontReference (a:fontRef) are available
- AND each reference has idx (index into theme) and optional color override

### Requirement: Color Types

The system SHALL support all DrawingML color types.

#### Scenario: sRGB color
- GIVEN an SrgbClr element (a:srgbClr)
- WHEN val attribute is accessed
- THEN a 6-digit hex color (RRGGBB) is available

#### Scenario: Scheme color
- GIVEN a SchemeClr element (a:schemeClr)
- WHEN val attribute is accessed
- THEN values dk1, lt1, dk2, lt2, accent1-6, hlink, folHlink are available

#### Scenario: System color
- GIVEN a SysClr element (a:sysClr)
- WHEN val attribute is accessed
- THEN values windowText, window, highlightText, highlight, buttonFace, etc. are available
- AND lastClr attribute stores last resolved color

#### Scenario: HSL color
- GIVEN an HslClr element (a:hslClr)
- WHEN properties are accessed
- THEN hue (0-21600000 = 0-360 degrees), sat (0-100000), lum (0-100000) are available

#### Scenario: Preset color
- GIVEN a PrstClr element (a:prstClr)
- WHEN val attribute is accessed
- THEN 150+ named colors like aliceBlue, antiqueWhite, aqua, black, blue, etc. are available

