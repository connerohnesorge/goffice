# DrawingML Charts

Delta specification for DrawingML chart types and components.

## Summary

This spec defines the DrawingML chart system that provides charting capabilities for Spreadsheet, Presentation, and Word documents. The chart system includes all major chart types, axes, legends, data labels, and formatting options.

## Motivation

Charts are a fundamental visualization component in Office documents. The DrawingML chart namespace (c:) provides:

- Multiple chart types (bar, line, pie, scatter, area, etc.)
- 2D and 3D chart variants
- Axis configuration (category, value, date, series)
- Data series with formula references
- Legend and title formatting
- Data labels and trendlines
- Chart formatting using DrawingML shape properties

All Office applications share the same chart format, making it efficient to implement once and share across SDKs.

## ADDED Requirements

### Requirement: ChartSpace Root Element

The system SHALL provide a `ChartSpace` element (c:chartSpace) as the root of chart parts.

#### Scenario: ChartSpace structure
- GIVEN a ChartSpace element
- WHEN children are accessed
- THEN Date1904 (c:date1904), Language (c:lang), RoundedCorners (c:roundedCorners), Style (c:style), ColorMapOverride (c:clrMapOvr), Chart (c:chart), ShapeProperties (c:spPr), TextProperties (c:txPr), ExternalData (c:externalData), PrintSettings (c:printSettings), UserShapes (c:userShapes) are available

#### Scenario: Chart style
- GIVEN a ChartSpace element
- WHEN Style (c:style) is accessed
- THEN val attribute contains style number (1-48 for built-in styles)

### Requirement: Chart Element

The system SHALL provide a `Chart` element (c:chart) containing the main chart definition.

#### Scenario: Chart structure
- GIVEN a Chart element
- WHEN children are accessed
- THEN Title (c:title), AutoTitleDeleted (c:autoTitleDeleted), PivotFormats (c:pivotFmts), View3D (c:view3D), Floor (c:floor), SideWall (c:sideWall), BackWall (c:backWall), PlotArea (c:plotArea), Legend (c:legend), PlotVisibleOnly (c:plotVisOnly), DisplayBlanksAs (c:dispBlanksAs), ShowDataLabelsOverMax (c:showDLblsOverMax) are available

#### Scenario: Display blanks as
- GIVEN a Chart element
- WHEN DisplayBlanksAs is accessed
- THEN values gap, zero, span are available

### Requirement: PlotArea Element

The system SHALL provide a `PlotArea` element (c:plotArea) for the chart plotting region.

#### Scenario: PlotArea charts
- GIVEN a PlotArea element
- WHEN chart type elements are accessed
- THEN AreaChart (c:areaChart), Area3DChart (c:area3DChart), BarChart (c:barChart), Bar3DChart (c:bar3DChart), BubbleChart (c:bubbleChart), DoughnutChart (c:doughnutChart), LineChart (c:lineChart), Line3DChart (c:line3DChart), OfPieChart (c:ofPieChart), PieChart (c:pieChart), Pie3DChart (c:pie3DChart), RadarChart (c:radarChart), ScatterChart (c:scatterChart), StockChart (c:stockChart), SurfaceChart (c:surfaceChart), Surface3DChart (c:surface3DChart) are available

#### Scenario: PlotArea axes
- GIVEN a PlotArea element
- WHEN axis elements are accessed
- THEN CategoryAxis (c:catAx), ValueAxis (c:valAx), DateAxis (c:dateAx), SeriesAxis (c:serAx) are available

#### Scenario: Layout
- GIVEN a PlotArea element
- WHEN Layout (c:layout) is accessed
- THEN ManualLayout (c:manualLayout) with x, y, w, h, xMode, yMode, wMode, hMode, target is available

### Requirement: Bar Chart Element

The system SHALL provide `BarChart` (c:barChart) and `Bar3DChart` (c:bar3DChart) elements.

#### Scenario: Bar chart properties
- GIVEN a BarChart element
- WHEN properties are accessed
- THEN BarDirection (c:barDir with val bar/col), Grouping (c:grouping with val clustered/stacked/percentStacked/standard), VaryColors (c:varyColors), Series (c:ser), DataLabels (c:dLbls), GapWidth (c:gapWidth), Overlap (c:overlap), SeriesLines (c:serLines), AxisId (c:axId) are available

#### Scenario: Bar direction
- GIVEN a BarChart element
- WHEN BarDirection val is "bar"
- THEN horizontal bars are rendered
- WHEN BarDirection val is "col"
- THEN vertical columns are rendered

#### Scenario: Bar 3D properties
- GIVEN a Bar3DChart element
- WHEN additional properties are accessed
- THEN Shape (c:shape with val box/cone/coneToMax/cylinder/pyramid/pyramidToMax), GapDepth (c:gapDepth) are available

### Requirement: Line Chart Element

The system SHALL provide `LineChart` (c:lineChart) and `Line3DChart` (c:line3DChart) elements.

#### Scenario: Line chart properties
- GIVEN a LineChart element
- WHEN properties are accessed
- THEN Grouping (c:grouping with val standard/stacked/percentStacked), VaryColors (c:varyColors), Series (c:ser), DataLabels (c:dLbls), DropLines (c:dropLines), HighLowLines (c:hiLowLines), UpDownBars (c:upDownBars), Marker (c:marker), Smooth (c:smooth), AxisId (c:axId) are available

#### Scenario: Line markers
- GIVEN a LineChart with Marker enabled
- WHEN series markers are accessed
- THEN Marker (c:marker) with Symbol (c:symbol), Size (c:size), ShapeProperties is available
- AND symbol values circle, dash, diamond, dot, none, picture, plus, square, star, triangle, x, auto are available

### Requirement: Pie Chart Elements

The system SHALL provide `PieChart` (c:pieChart), `Pie3DChart` (c:pie3DChart), and `DoughnutChart` (c:doughnutChart) elements.

#### Scenario: Pie chart properties
- GIVEN a PieChart element
- WHEN properties are accessed
- THEN VaryColors (c:varyColors), Series (c:ser), DataLabels (c:dLbls), FirstSliceAngle (c:firstSliceAng) are available

#### Scenario: Doughnut chart properties
- GIVEN a DoughnutChart element
- WHEN properties are accessed
- THEN same as PieChart plus HoleSize (c:holeSize with val 10-90 percent) is available

#### Scenario: Of-pie chart
- GIVEN an OfPieChart element (c:ofPieChart)
- WHEN properties are accessed
- THEN OfPieType (c:ofPieType with val pie/bar), Series, GapWidth, SplitType, SplitPos, CustSplit, SecondPieSize are available

### Requirement: Area Chart Elements

The system SHALL provide `AreaChart` (c:areaChart) and `Area3DChart` (c:area3DChart) elements.

#### Scenario: Area chart properties
- GIVEN an AreaChart element
- WHEN properties are accessed
- THEN Grouping (c:grouping), VaryColors (c:varyColors), Series (c:ser), DropLines (c:dropLines), DataLabels (c:dLbls), AxisId (c:axId) are available

### Requirement: Scatter Chart Element

The system SHALL provide `ScatterChart` (c:scatterChart) and `BubbleChart` (c:bubbleChart) elements.

#### Scenario: Scatter chart properties
- GIVEN a ScatterChart element
- WHEN properties are accessed
- THEN ScatterStyle (c:scatterStyle with val line/lineMarker/marker/none/smooth/smoothMarker), VaryColors (c:varyColors), Series (c:ser), DataLabels (c:dLbls), AxisId (c:axId) are available

#### Scenario: Bubble chart properties
- GIVEN a BubbleChart element
- WHEN properties are accessed
- THEN VaryColors, Series, DataLabels, Bubble3D (c:bubble3D), BubbleScale (c:bubbleScale), ShowNegBubbles (c:showNegBubbles), SizeRepresents (c:sizeRepresents with val area/w) are available

#### Scenario: Scatter series data
- GIVEN a ScatterChartSeries element
- WHEN data is accessed
- THEN XValues (c:xVal) and YValues (c:yVal) are available
- AND each contains NumericDataSource or StringDataSource

### Requirement: Other Chart Types

The system SHALL provide additional chart type elements.

#### Scenario: Radar chart
- GIVEN a RadarChart element (c:radarChart)
- WHEN properties are accessed
- THEN RadarStyle (c:radarStyle with val standard/marker/filled), VaryColors, Series, DataLabels, AxisId are available

#### Scenario: Stock chart
- GIVEN a StockChart element (c:stockChart)
- WHEN properties are accessed
- THEN Series (3-5 for high-low-close, open-high-low-close, volume-high-low-close, volume-open-high-low-close), DropLines, HighLowLines, UpDownBars, AxisId are available

#### Scenario: Surface chart
- GIVEN a SurfaceChart element (c:surfaceChart)
- WHEN properties are accessed
- THEN Wireframe (c:wireframe), Series, BandFormats (c:bandFmts), AxisId are available

### Requirement: Chart Series

The system SHALL provide series elements for chart data.

#### Scenario: Series common properties
- GIVEN a series element (e.g., BarChartSeries)
- WHEN properties are accessed
- THEN Index (c:idx), Order (c:order), SeriesText (c:tx), ShapeProperties (c:spPr), InvertIfNegative (c:invertIfNegative), PictureOptions (c:pictureOptions), DataPoints (c:dPt), DataLabels (c:dLbls), Trendline (c:trendline), ErrorBars (c:errBars) are available

#### Scenario: Category reference
- GIVEN a series element
- WHEN Categories (c:cat) is accessed
- THEN StringReference (c:strRef) or NumericReference (c:numRef) or MultiLevelStringReference (c:multiLvlStrRef) or StringLiteral (c:strLit) or NumericLiteral (c:numLit) is available

#### Scenario: Value reference
- GIVEN a series element
- WHEN Values (c:val) is accessed
- THEN NumericReference (c:numRef) or NumericLiteral (c:numLit) is available

#### Scenario: Formula reference
- GIVEN a NumericReference element
- WHEN properties are accessed
- THEN Formula (c:f) contains cell range reference (e.g., "Sheet1!$B$2:$B$10")
- AND NumericCache (c:numCache) contains cached values

#### Scenario: String reference
- GIVEN a StringReference element
- WHEN properties are accessed
- THEN Formula (c:f) contains cell range reference
- AND StringCache (c:strCache) contains cached values

### Requirement: Chart Axes

The system SHALL provide axis elements for charts.

#### Scenario: Category axis properties
- GIVEN a CategoryAxis element (c:catAx)
- WHEN properties are accessed
- THEN AxisId (c:axId), Scaling (c:scaling), Delete (c:delete), AxisPosition (c:axPos with val b/l/r/t), MajorGridlines (c:majorGridlines), MinorGridlines (c:minorGridlines), Title (c:title), NumberFormat (c:numFmt), MajorTickMark (c:majorTickMark), MinorTickMark (c:minorTickMark), TickLabelPosition (c:tickLblPos), ShapeProperties (c:spPr), TextProperties (c:txPr), CrossingAxis (c:crossAx), Crosses (c:crosses with val autoZero/max/min), CrossesAt (c:crossesAt), Auto (c:auto), LabelAlignment (c:lblAlgn), LabelOffset (c:lblOffset), TickLabelSkip (c:tickLblSkip), TickMarkSkip (c:tickMarkSkip), NoMultiLevelLabels (c:noMultiLvlLbl) are available

#### Scenario: Value axis properties
- GIVEN a ValueAxis element (c:valAx)
- WHEN properties are accessed
- THEN same as CategoryAxis plus CrossBetween (c:crossBetween with val between/midCat), MajorUnit (c:majorUnit), MinorUnit (c:minorUnit), DisplayUnits (c:dispUnits) are available

#### Scenario: Date axis properties
- GIVEN a DateAxis element (c:dateAx)
- WHEN properties are accessed
- THEN same as CategoryAxis plus BaseTimeUnit (c:baseTimeUnit), MajorTimeUnit (c:majorTimeUnit), MinorTimeUnit (c:minorTimeUnit), MajorUnit (c:majorUnit), MinorUnit (c:minorUnit) are available
- AND time unit values days, months, years are available

#### Scenario: Scaling properties
- GIVEN a Scaling element (c:scaling)
- WHEN properties are accessed
- THEN LogBase (c:logBase), Orientation (c:orientation with val maxMin/minMax), Minimum (c:min), Maximum (c:max) are available

### Requirement: Chart Title

The system SHALL provide a `Title` element (c:title) for chart titles.

#### Scenario: Title structure
- GIVEN a Title element
- WHEN properties are accessed
- THEN Text (c:tx), Layout (c:layout), Overlay (c:overlay), ShapeProperties (c:spPr), TextProperties (c:txPr) are available

#### Scenario: Rich text title
- GIVEN a Title with rich text
- WHEN Text element contains Rich (c:rich)
- THEN DrawingML TextBody (a:txBody) with paragraphs and runs is available

#### Scenario: String reference title
- GIVEN a Title with cell reference
- WHEN Text element contains StringReference (c:strRef)
- THEN formula references a cell for dynamic title

### Requirement: Chart Legend

The system SHALL provide a `Legend` element (c:legend) for chart legends.

#### Scenario: Legend properties
- GIVEN a Legend element
- WHEN properties are accessed
- THEN LegendPosition (c:legendPos with val b/l/r/t/tr), LegendEntry (c:legendEntry), Layout (c:layout), Overlay (c:overlay), ShapeProperties (c:spPr), TextProperties (c:txPr) are available

#### Scenario: Legend entry customization
- GIVEN a LegendEntry element
- WHEN properties are accessed
- THEN Index (c:idx), Delete (c:delete), TextProperties (c:txPr) are available

### Requirement: Data Labels

The system SHALL provide a `DataLabels` element (c:dLbls) for chart data labels.

#### Scenario: Data labels properties
- GIVEN a DataLabels element
- WHEN properties are accessed
- THEN NumberFormat (c:numFmt), ShapeProperties (c:spPr), TextProperties (c:txPr), DataLabelPosition (c:dLblPos), ShowLegendKey (c:showLegendKey), ShowValue (c:showVal), ShowCategoryName (c:showCatName), ShowSeriesName (c:showSerName), ShowPercent (c:showPercent), ShowBubbleSize (c:showBubbleSize), Separator (c:separator), ShowLeaderLines (c:showLeaderLines) are available

#### Scenario: Data label position
- GIVEN a DataLabelPosition element
- WHEN val attribute is accessed
- THEN values bestFit, b, ctr, inBase, inEnd, l, outEnd, r, t are available

#### Scenario: Individual data label
- GIVEN a DataLabel element (c:dLbl)
- WHEN properties are accessed
- THEN Index (c:idx), Delete (c:delete), Layout, plus data label properties are available

### Requirement: Trendlines

The system SHALL provide a `Trendline` element (c:trendline) for chart trendlines.

#### Scenario: Trendline types
- GIVEN a Trendline element
- WHEN TrendlineType (c:trendlineType) is accessed
- THEN values exp (exponential), linear, log (logarithmic), movingAvg (moving average), poly (polynomial), power are available

#### Scenario: Trendline properties
- GIVEN a Trendline element
- WHEN properties are accessed
- THEN Name (c:name), ShapeProperties (c:spPr), TrendlineType (c:trendlineType), Order (c:order for polynomial), Period (c:period for moving avg), Forward (c:forward), Backward (c:backward), Intercept (c:intercept), DisplayRSquared (c:dispRSqr), DisplayEquation (c:dispEq), TrendlineLabel (c:trendlineLbl) are available

### Requirement: Error Bars

The system SHALL provide an `ErrorBars` element (c:errBars) for chart error bars.

#### Scenario: Error bar direction
- GIVEN an ErrorBars element
- WHEN ErrorDirection (c:errDir) is accessed
- THEN values x, y are available

#### Scenario: Error bar type
- GIVEN an ErrorBars element
- WHEN ErrorBarType (c:errBarType) is accessed
- THEN values both, minus, plus are available

#### Scenario: Error value type
- GIVEN an ErrorBars element
- WHEN ErrorValueType (c:errValType) is accessed
- THEN values cust (custom), fixedVal, percentage, stdDev, stdErr are available

#### Scenario: Custom error values
- GIVEN an ErrorBars with custom values
- WHEN Plus (c:plus) and Minus (c:minus) are accessed
- THEN NumericDataSource with formula or literal values is available

### Requirement: 3D View Settings

The system SHALL provide a `View3D` element (c:view3D) for 3D chart settings.

#### Scenario: 3D view properties
- GIVEN a View3D element
- WHEN properties are accessed
- THEN RotateX (c:rotX -90 to 90), RotateY (c:rotY 0 to 360), RightAngleAxes (c:rAngAx), Perspective (c:perspective 0-240), HeightPercent (c:hPercent), DepthPercent (c:depthPercent) are available

### Requirement: Chart Formatting

The system SHALL support chart formatting using DrawingML.

#### Scenario: Shape properties
- GIVEN a ShapeProperties element in chart
- WHEN properties are accessed
- THEN DrawingML Fill (SolidFill, GradientFill, etc.), Outline (a:ln), EffectList, Scene3D are available

#### Scenario: Text properties
- GIVEN a TextProperties element in chart
- WHEN properties are accessed
- THEN DrawingML BodyProperties (a:bodyPr), ListStyle (a:lstStyle), Paragraph with RunProperties are available

## Design

### Package Structure

```
drawingml/
  chart/              # Chart types (c: namespace)
    chartspace.go     # ChartSpace root element
    chart.go          # Chart element
    plotarea.go       # PlotArea element
    barchart.go       # BarChart, Bar3DChart
    linechart.go      # LineChart, Line3DChart
    piechart.go       # PieChart, Pie3DChart, DoughnutChart
    areachart.go      # AreaChart, Area3DChart
    scatterchart.go   # ScatterChart, BubbleChart
    otherchart.go     # Radar, Stock, Surface charts
    series.go         # Series elements and data references
    axis.go           # CategoryAxis, ValueAxis, DateAxis, SeriesAxis
    title.go          # Title element
    legend.go         # Legend element
    datalabel.go      # DataLabels, DataLabel
    trendline.go      # Trendline element
    errorbar.go       # ErrorBars element
    view3d.go         # View3D element
```

### Type Naming Convention

| XML Element | Go Type |
|-------------|---------|
| c:chartSpace | ChartSpace |
| c:chart | Chart |
| c:plotArea | PlotArea |
| c:barChart | BarChart |
| c:lineChart | LineChart |
| c:pieChart | PieChart |
| c:scatterChart | ScatterChart |
| c:ser | Ser |
| c:catAx | CatAx |
| c:valAx | ValAx |
| c:title | Title |
| c:legend | Legend |
| c:dLbls | DLbls |

## API

### Creating a Chart

```go
// Create chart space
chartSpace := chart.NewChartSpace()

// Create chart with bar chart
ch := chart.NewChart()
plotArea := chart.NewPlotArea()

barChart := chart.NewBarChart(
    chart.WithBarDirection("col"),      // column chart
    chart.WithGrouping("clustered"),
)

// Add series
series := chart.NewSer(
    chart.WithSeriesIndex(0),
    chart.WithSeriesOrder(0),
    chart.WithSeriesText("Sales"),
    chart.WithCategories(chart.NewStrRef("Sheet1!$A$2:$A$6")),
    chart.WithValues(chart.NewNumRef("Sheet1!$B$2:$B$6")),
)
barChart.AddSeries(series)

plotArea.AddChart(barChart)
ch.SetPlotArea(plotArea)
chartSpace.SetChart(ch)
```

### Chart Axes

```go
// Create category axis
catAx := chart.NewCatAx(
    chart.WithAxisId(1),
    chart.WithAxisPosition("b"),  // bottom
    chart.WithCrossingAxis(2),
)

// Create value axis
valAx := chart.NewValAx(
    chart.WithAxisId(2),
    chart.WithAxisPosition("l"),  // left
    chart.WithCrossingAxis(1),
    chart.WithMajorGridlines(),
)

plotArea.AddAxis(catAx)
plotArea.AddAxis(valAx)
```

### Chart Title and Legend

```go
// Add title
title := chart.NewTitle(
    chart.WithRichText("Quarterly Sales"),
)
ch.SetTitle(title)

// Add legend
legend := chart.NewLegend(
    chart.WithLegendPosition("r"),  // right
)
ch.SetLegend(legend)
```

### Data Labels

```go
// Configure data labels
dLbls := chart.NewDLbls(
    chart.WithShowValue(true),
    chart.WithShowPercent(false),
    chart.WithDataLabelPosition("outEnd"),
)
barChart.SetDataLabels(dLbls)
```

## Testing

### Unit Tests

- ChartSpace creation and serialization
- All chart type elements
- Series data references (formula and cache)
- Axis configuration
- Title and legend properties
- Data label options

### Integration Tests

- Complete chart round-trip (create, serialize, parse)
- Chart with multiple series
- Chart with multiple axes
- 3D chart view settings
- Trendline calculations

### Validation Tests

- Invalid chart type combinations
- Out-of-range axis values
- Invalid trendline types for chart types
- Missing required axis references

## Dependencies

- `encoding/xml` - Standard library XML support
- `drawingml/main` - Core DrawingML types for shape properties
- No external dependencies

## Migration

N/A - New capability with no existing implementation.
