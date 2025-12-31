# E2E Visual Test Case Library

This directory contains comprehensive test cases for the E2E visual testing framework.

## Overview

- **Total Test Cases:** 21
- **Categories:** Charts (10), Shapes (5), Text (4), Integration (2)
- **Format:** JSON files conforming to `framework.TestCase` structure

## Directory Structure

```
testcases/
├── charts/          # Chart rendering tests (10 cases)
├── shapes/          # Shape rendering tests (5 cases)
├── text/            # Text formatting tests (4 cases)
└── integration/     # Complex integration tests (2 cases)
```

## Chart Test Cases (10)

### 1. bar_basic.json
- **ID:** chart_bar_clustered_01
- **Description:** Basic clustered bar chart with 3 categories and 2 series
- **Tags:** bar, clustered, basic
- **Complexity:** Low

### 2. bar_stacked.json
- **ID:** chart_bar_stacked_01
- **Description:** Stacked bar chart with 4 categories, 3 series, legend, and data labels
- **Tags:** bar, stacked, legend, labels
- **Complexity:** Medium

### 3. bar_100_percent.json
- **ID:** chart_bar_percent_stacked_01
- **Description:** 100% stacked bar chart with 5 categories and 2 series
- **Tags:** bar, percent-stacked, percentage
- **Complexity:** Medium

### 4. line_basic.json
- **ID:** chart_line_smooth_01
- **Description:** Smooth line chart with 6 data points and 2 series
- **Tags:** line, smooth, basic
- **Complexity:** Low

### 5. line_markers.json
- **ID:** chart_line_markers_01
- **Description:** Line chart with different marker styles (circle, square, triangle)
- **Tags:** line, markers, symbols
- **Complexity:** Medium

### 6. pie_basic.json
- **ID:** chart_pie_01
- **Description:** Basic pie chart with 5 slices and percentage labels
- **Tags:** pie, basic, labels
- **Complexity:** Low

### 7. pie_exploded.json
- **ID:** chart_pie_exploded_01
- **Description:** Exploded pie chart with 6 slices
- **Tags:** pie, exploded
- **Complexity:** Medium

### 8. doughnut.json
- **ID:** chart_doughnut_01
- **Description:** Doughnut chart with 4 slices
- **Tags:** doughnut, pie, hollow
- **Complexity:** Low

### 9. scatter_xy.json
- **ID:** chart_scatter_01
- **Description:** Scatter plot with 10 data points and labeled axes
- **Tags:** scatter, xy, plot
- **Complexity:** Medium

### 10. area_stacked.json
- **ID:** chart_area_stacked_01
- **Description:** Stacked area chart with 5 data points and 3 series
- **Tags:** area, stacked, filled
- **Complexity:** Medium

## Shape Test Cases (5)

### 11. rectangles.json
- **ID:** shape_rectangles_01
- **Description:** Various rectangle shapes with different fills and outlines
- **Tags:** rectangle, rounded, basic
- **Complexity:** Low

### 12. circles.json
- **ID:** shape_circles_01
- **Description:** Circle and ellipse shapes with solid and gradient fills
- **Tags:** circle, ellipse, gradient
- **Complexity:** Medium

### 13. polygons.json
- **ID:** shape_polygons_01
- **Description:** Triangle, pentagon, and hexagon shapes
- **Tags:** polygon, triangle, pentagon, hexagon
- **Complexity:** Low

### 14. arrows.json
- **ID:** shape_arrows_01
- **Description:** Various arrow shapes with different orientations
- **Tags:** arrow, direction, block
- **Complexity:** Medium

### 15. effects.json
- **ID:** shape_effects_01
- **Description:** Shape effects including shadow, glow, and reflection
- **Tags:** effects, shadow, glow, reflection
- **Complexity:** High

## Text Test Cases (4)

### 16. formatting.json
- **ID:** text_formatting_01
- **Description:** Various text formatting (bold, italic, underline, font sizes, colors)
- **Tags:** text, formatting, font, styles
- **Complexity:** Medium

### 17. alignment.json
- **ID:** text_alignment_01
- **Description:** Text alignment options (left, center, right, justify)
- **Tags:** text, alignment, justify
- **Complexity:** Medium

### 18. multilevel.json
- **ID:** text_multilevel_01
- **Description:** Multi-level bulleted and numbered lists
- **Tags:** text, bullets, numbering, lists
- **Complexity:** Medium

### 19. fonts.json
- **ID:** text_fonts_01
- **Description:** Different font families (Arial, Times New Roman, Calibri)
- **Tags:** text, fonts, typography
- **Complexity:** Low

## Integration Test Cases (2)

### 20. dashboard.json
- **ID:** integration_dashboard_01
- **Description:** Full dashboard with multiple charts (bar, line, pie) and text elements
- **Tags:** dashboard, complex, charts, text, integration
- **Complexity:** High

### 21. mixed_content.json
- **ID:** integration_mixed_01
- **Description:** Mixed content with charts, shapes with text, callouts, and effects
- **Tags:** mixed, integration, complex, shapes, charts, text
- **Complexity:** High

## Test Configuration

All test cases use standard configuration:
- **DPI:** 300
- **Backend:** LibreOffice
- **Diff Threshold:** 0.02-0.03 (2-3% difference allowed)
- **Diff Algorithm:** SSIM (Structural Similarity Index)
- **Ignore Antialiasing:** true
- **Retry Count:** 2
- **Timeouts:** 60-90s generation, 120-180s rendering

## Usage

### Running Individual Test Cases

```bash
# Run a specific test case
go test ./tests/e2e -run TestVisualComparison -testcase charts/bar_basic.json

# Run all chart tests
go test ./tests/e2e -run TestVisualComparison -category chart

# Run by tag
go test ./tests/e2e -run TestVisualComparison -tag basic
```

### Adding New Test Cases

1. Create a JSON file in the appropriate category directory
2. Use the TestCase structure from `framework/testcase.go`
3. Validate JSON with `jq . your-test.json`
4. Run the test to generate baseline images

## EMU Units Reference

All positions and sizes use EMUs (English Metric Units):
- **1 inch = 914,400 EMUs**
- **1 cm = 360,000 EMUs**
- **1 pt = 12,700 EMUs**

Standard slide sizes:
- **16:9:** 9,144,000 × 5,143,500 EMUs (10" × 5.625")
- **4:3:** 9,144,000 × 6,858,000 EMUs (10" × 7.5")

## Color Format

Colors are specified as 6-digit hex strings (without # prefix):
- `"4472C4"` - Blue
- `"ED7D31"` - Orange
- `"70AD47"` - Green
- `"FFC000"` - Gold
- `"000000"` - Black
- `"FFFFFF"` - White

For transparency, use 8-digit hex with alpha channel:
- `"80000000"` - Semi-transparent black (50% opacity)

## Validation

All test cases have been validated:
```bash
# Validate all JSON files
for file in testcases/*/*.json; do jq empty "$file" && echo "✓ $file"; done
```

All 21 test cases are valid JSON and conform to the TestCase schema.
