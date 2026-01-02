# DRAWINGML KNOWLEDGE BASE

## OVERVIEW
Implementation of the shared DrawingML specification. Used by Word, Excel, and PowerPoint for visuals.
Includes Charts, Shapes, Pictures, and Transformations.

## STRUCTURE
```
drawingml/
├── diagram/      # SmartArt and diagrams
├── table/        # DrawingML tables (different from Word tables)
├── chart.go      # High-level chart wrappers (ChartSpace, Chart, PlotArea)
├── color.go      # Color models (RGB, Scheme, System)
├── geometry.go   # Shape paths and presets
└── transform.go  # 2D transforms (rotation, offset, scale)
```

## WHERE TO LOOK
| Task | Location | Notes |
|------|----------|-------|
| **Chart Space** | `chart.go` | `ChartSpace` is the root element for chart parts |
| **Chart Data** | `chart.go` | `Chart` contains `Title`, `PlotArea`, `Legend` |
| **Plotting** | `chart.go` | `PlotArea` contains specific charts (Bar, Line, Pie) |
| **Series** | `chart_series.go` | Data series management |
| **Colors** | `color.go` | Handling HSL/RGB/Scheme colors |
| **Position** | `transform.go` | Offsets and Extents (EMU units) |

## CONVENTIONS
*   **Hierarchy:** `ChartSpace` -> `Chart` -> `PlotArea` -> `[Type]Chart` -> `Series`.
*   **EMUs:** Coordinates are in English Metric Units (914400 per inch). Use `units.go` helpers.
*   **Shared:** Changes here affect all three document types.
*   **Properties:** Shapes use `SpPr` (Shape Properties) for fill, line, and effects.