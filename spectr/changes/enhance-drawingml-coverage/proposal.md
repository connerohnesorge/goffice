# Change: Enhance DrawingML Coverage for Complete Shape and Chart Support

## Why
goffice implements basic DrawingML for shapes, text boxes, and charts but lacks comprehensive coverage of advanced drawing features present in Open-XML-SDK:
- Advanced shape properties (3D, shadows, reflections, glows)
- Fill patterns and gradients (linear, radial, path)
- Stroke properties (dashes, caps, joins, compound)
- Text box properties (wrapping, margins, rotation)
- Advanced text formatting within shapes (character spacing, kerning, ligatures)
- Shape connectors and connection points
- Callouts and line shapes
- WordprocessingDrawing integration (anchoring, wrapping)
- Picture effects (cropping, compression, borders)
- Chart axis formatting and grid lines
- Chart data labels and value labels
- Chart legend customization
- Line and shape effects

## What Changes
- Add 3D shape properties (extrusion, bevel, material, lighting)
- Add shadow effects (outer, inner, perspective)
- Add reflection and glow effects
- Add linear, radial, and path gradient fills
- Add pattern fills (polka dots, stripes, etc.)
- Add solid fill with color, transparency, and theme colors
- Add advanced stroke properties (dash styles, line caps, line joins)
- Add text box rotation and vertical text
- Add text wrapping modes (none, square, tight, through)
- Add shape connectors with adjustment handles
- Add picture cropping and compression
- Add chart axis formatting (number format, label position)
- Add chart grid lines (major, minor)
- Add chart data label formatting
- Add effect styles (theme effects)
- BREAKING: Shape and chart APIs extended with optional effect parameters

## Impact
- Affected specs: drawingml-core, drawingml-text, drawingml-charts, drawingml-anchors
- Affected code: drawingml/, spreadsheet/chart.go, presentation/elements/, wordprocessing/elements/
- New dependencies: None (uses stdlib only)
- Test coverage required: Roundtrip tests for all effect combinations
