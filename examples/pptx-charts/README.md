# PowerPoint Charts Example

This example demonstrates creating PowerPoint presentations with various chart types using the goffice library.

## Charts Demonstrated

The example creates a presentation with 6 slides, each showcasing a different chart type:

1. **Bar Chart (Clustered)** - Quarterly Sales Comparison
   - Shows sales data for three regions (North, South, East) across four quarters
   - Demonstrates bar chart with multiple series and category/value axes

2. **Line Chart** - Monthly Temperature Trends
   - Displays average temperatures for three cities throughout the year
   - Shows line chart with standard grouping and multiple series

3. **Pie Chart** - Market Share Distribution
   - Illustrates market share percentages for five companies
   - Demonstrates pie chart with category data and legend

4. **Area Chart (Stacked)** - Cumulative Revenue Growth
   - Visualizes revenue growth for three product lines over six months
   - Shows stacked area chart to represent cumulative values

5. **Scatter Chart** - Price vs Demand Analysis
   - Correlates price points with demand for two product lines
   - Demonstrates scatter chart with X and Y values

6. **Doughnut Chart** - Annual Budget Allocation
   - Breaks down budget distribution across six departments
   - Shows doughnut chart with customized hole size

## Usage

```bash
go run main.go
```

This will create a file named `charts-example.pptx` in the current directory.

## Chart API Features

The example showcases:

- Creating different chart types (Bar, Line, Pie, Area, Scatter, Doughnut)
- Adding multiple series to charts
- Setting chart titles and legends
- Configuring category and value axes
- Using Excel-style data references (e.g., "Sheet1!$A$2:$A$5")
- Positioning legends (right, bottom, top)
- Customizing chart properties (hole size for doughnut, scatter styles)

## Implementation Note

This example demonstrates the chart API structure. In the current implementation, the charts are created as ChartSpace objects with all the proper data structures, but the final integration into the presentation (creating chart parts and linking them via graphic frames) is handled with placeholder shapes.

A complete implementation would include:
- Creating ChartPart objects in the presentation
- Writing the ChartSpace XML to the chart part
- Creating GraphicFrame elements on slides
- Linking graphic frames to chart parts via relationship IDs

The example provides a solid foundation for understanding how to construct charts programmatically with the goffice library.

## Data References

The chart series reference Excel-style cell ranges like:
- `"Sheet1!$A$2:$A$5"` - Category labels or X values
- `"Sheet1!$B$2:$B$5"` - Numeric values or Y values

In a real presentation with embedded Excel data, these references would point to the actual spreadsheet data within the presentation package.

## Building

```bash
go build
./pptx-charts
```

The example requires the goffice library:
```bash
go get github.com/connerohnesorge/goffice
```
