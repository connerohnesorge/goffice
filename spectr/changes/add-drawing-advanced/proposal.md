# Change: Advanced Drawing (DrawingML) Support Parity

## Why
DrawingML (Drawing Markup Language) support is basic. Complex shapes, 3D models, charts with data labels, SmartArt, and advanced shape manipulation are missing or incomplete.

## What Changes
- Advanced shape properties (effects, 3D, adjustments)
- Shape text wrapping with proper binding
- Picture cropping, rotation, and transformation
- Chart data binding and label positioning
- SmartArt shape hierarchies and styling
- Connector shapes with routing
- Callout shapes with properties
- Custom shape definition support
- Diagram markup language support
- Shape inheritance and grouping
- Z-order and layering control

## Impact
- Affected specs: drawingml-core, drawingml-charts, drawingml-anchors, drawingml-text
- Affected code: drawingml/*.go, pdf/drawing/*.go
- Breaking changes: None (additive only)
