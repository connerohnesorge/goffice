# Work-in-Progress Renderer Files

The following renderer files have been temporarily disabled (renamed to `.wip` extension) because they reference an incomplete API:

## Disabled Files

### Source Files
- `chart_renderer.go.wip` - Chart rendering functionality
- `effects_renderer.go.wip` - Visual effects (shadows, glows, etc.)
- `fill_renderer.go.wip` - Fill rendering (solid, gradient, pattern)
- `image_renderer.go.wip` - Image/picture rendering
- `shape_renderer.go.wip` - Shape rendering with geometry
- `stroke_renderer.go.wip` - Stroke/line rendering
- `text_in_shape.go.wip` - Text rendering within shapes
- `transform_renderer.go.wip` - Transformation rendering

### Test Files
- `chart_renderer_test.go.wip` - Tests for chart renderer
- `fill_renderer_test.go.wip` - Tests for fill renderer
- `shape_renderer_test.go.wip` - Tests for shape renderer
- `drawingml_integration_test.go.wip` - Integration tests for DrawingML rendering

## Issues

These files were attempting to use `core.RenderContext` (which should be `core.RenderingContext`) and more critically, they reference a `Page` field on the rendering context that doesn't exist.

The files also use methods like:
- `ctx.Page.SetFillColor()`
- `ctx.Page.DrawRectangle()`
- `ctx.Page.WriteContent()`
- `ctx.Page.SaveGraphicsState()`

However, `core.RenderingContext` doesn't have a `Page` field. This API needs to be designed and implemented before these renderer files can be completed.

## Fixes Applied

1. Changed `core.RenderContext` to `core.RenderingContext` throughout
2. Changed `*Path` to `*PathBuilder` to match the actual type
3. Changed `path.Fill()` to `ctx.Page.WriteContent(path.Fill())` (where appropriate)
4. Changed `path.Close()` to `path.ClosePath()` to match the correct method name

## To Re-enable

Once the core rendering API is properly defined with a Page abstraction, rename these files back to `.go` extension and they should work (modulo any additional API changes needed).
