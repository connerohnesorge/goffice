# Bar Chart Rendering Fix

## Problem
Bar charts created with goffice were not rendering visible bars in LibreOffice Impress, while line charts rendered correctly.

## Root Cause
Bar chart series were missing `<c:spPr>` (shape properties) elements with fill colors. LibreOffice requires explicit fill colors for bar chart series, while line charts can use default styling.

## Solution
Modified `/home/connerohnesorge/Documents/001Repos/goffice/examples/pptx-charts/helpers.go` to add shape properties with solid fill colors to each bar chart series.

### Changes Made:

1. **Added Office color palette:**
   ```go
   var chartColors = []string{
       "4472C4", // Blue
       "ED7D31", // Orange
       "A5A5A5", // Gray
       "FFC000", // Yellow
       "5B9BD5", // Light Blue
       "70AD47", // Green
   }
   ```

2. **Updated `addBarChartSeries` function:**
   ```go
   // Add shape properties with fill color for bar charts
   // This is required for LibreOffice to render the bars correctly
   spPr := drawingml.NewChartShapeProperties()
   colorIndex := cfg.seriesIndex % uint32(len(chartColors))
   spPr.SetSolidFill(chartColors[colorIndex])
   series.SetShapeProperties(spPr)
   ```

## XML Comparison

### Before (No bars visible):
```xml
<c:ser>
  <c:idx val="0"/>
  <c:order val="0"/>
  <c:tx><c:v>North</c:v></c:tx>
  <c:cat>...</c:cat>
  <c:val>...</c:val>
</c:ser>
```

### After (Bars render correctly):
```xml
<c:ser>
  <c:idx val="0"/>
  <c:order val="0"/>
  <c:tx><c:v>North</c:v></c:tx>
  <c:spPr>
    <a:solidFill>
      <a:srgbClr val="4472C4"/>
    </a:solidFill>
  </c:spPr>
  <c:cat>...</c:cat>
  <c:val>...</c:val>
</c:ser>
```

## Testing
The fix has been applied and the example regenerated. The bar chart now includes proper shape properties with fill colors for each series.

## Files Modified
- `/home/connerohnesorge/Documents/001Repos/goffice/examples/pptx-charts/helpers.go`

## Status
✓ Fix implemented
✓ Code compiled successfully
✓ Example regenerated with correct XML structure
- Ready for visual testing in LibreOffice Impress
