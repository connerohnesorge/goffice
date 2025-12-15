## ADDED Requirements

### Requirement: ChartSpace Root Element

The system SHALL provide a `ChartSpace` element as the root of ChartPart.

#### Scenario: Access chart components
- GIVEN a ChartSpace element
- WHEN components are accessed
- THEN Chart, PrintSettings, ExternalData, Style, ClrMapOvr are available

### Requirement: Chart Element

The system SHALL provide a `Chart` element containing chart definition.

#### Scenario: Chart structure
- GIVEN a Chart element
- WHEN components are accessed
- THEN Title, AutoTitleDeleted, PivotFmts, View3D, Floor, SideWall, BackWall, PlotArea, Legend, PlotVisOnly, DispBlanksAs are available

### Requirement: PlotArea Element

The system SHALL provide a `PlotArea` element for chart plot area.

#### Scenario: Plot area charts
- GIVEN a PlotArea element
- WHEN chart types are accessed
- THEN AreaChart, Area3DChart, BarChart, Bar3DChart, BubbleChart, DoughnutChart, LineChart, Line3DChart, OfPieChart, PieChart, Pie3DChart, RadarChart, ScatterChart, StockChart, SurfaceChart, Surface3DChart are available

#### Scenario: Plot area axes
- GIVEN a PlotArea element
- WHEN axes are accessed
- THEN CategoryAxis, ValueAxis, DateAxis, SeriesAxis are available

### Requirement: Chart Types

The system SHALL support all major chart types.

#### Scenario: Bar/Column chart
- GIVEN a BarChart element
- WHEN properties are accessed
- THEN BarDirection (bar/col), Grouping (clustered/stacked/percentStacked), VaryColors, Series, DataLabels, GapWidth, Overlap are available

#### Scenario: Line chart
- GIVEN a LineChart element
- WHEN properties are accessed
- THEN Grouping (standard/stacked/percentStacked), VaryColors, Series, DataLabels, DropLines, HiLowLines, UpDownBars, Marker, Smooth are available

#### Scenario: Pie chart
- GIVEN a PieChart or DoughnutChart element
- WHEN properties are accessed
- THEN VaryColors, Series, DataLabels, FirstSliceAngle, HoleSize (doughnut only) are available

#### Scenario: Scatter chart
- GIVEN a ScatterChart element
- WHEN properties are accessed
- THEN ScatterStyle (lineMarker/line/marker/smooth/smoothMarker), VaryColors, Series, DataLabels are available

#### Scenario: Area chart
- GIVEN an AreaChart element
- WHEN properties are accessed
- THEN Grouping, VaryColors, Series, DropLines, DataLabels are available

### Requirement: Chart Series

The system SHALL provide series definition for chart data.

#### Scenario: Series properties
- GIVEN a series element (e.g., BarChartSeries)
- WHEN properties are accessed
- THEN Index, Order, SeriesText, ShapeProperties, InvertIfNegative, PictureOptions, DataPoints, DataLabels, Trendline, ErrorBars, Categories, Values are available

#### Scenario: Series data references
- GIVEN a Categories or Values element
- WHEN data reference is accessed
- THEN NumericDataSource or StringDataSource with formula reference is available

#### Scenario: Series formula reference
- GIVEN a NumRef element (numeric reference)
- WHEN properties are accessed
- THEN Formula (e.g., "Sheet1!$B$2:$B$10"), NumCache with cached values are available

### Requirement: Chart Axes

The system SHALL provide axis definitions for charts.

#### Scenario: Category axis
- GIVEN a CategoryAxis element
- WHEN properties are accessed
- THEN AxisId, Scaling, Delete, AxisPosition, MajorGridlines, MinorGridlines, Title, NumberFormat, MajorTickMark, MinorTickMark, TickLabelPosition, ShapeProperties, TextProperties, CrossingAxis, Crosses, Auto, LabelAlignment, LabelOffset are available

#### Scenario: Value axis
- GIVEN a ValueAxis element
- WHEN properties are accessed
- THEN similar to CategoryAxis plus CrossBetween, MajorUnit, MinorUnit, DispUnits are available

#### Scenario: Date axis
- GIVEN a DateAxis element
- WHEN properties are accessed
- THEN similar to CategoryAxis plus BaseTimeUnit, MajorTimeUnit, MinorTimeUnit, MajorUnit, MinorUnit are available

### Requirement: Chart Title and Legend

The system SHALL provide title and legend elements.

#### Scenario: Chart title
- GIVEN a Title element
- WHEN properties are accessed
- THEN Text, Layout, Overlay, ShapeProperties, TextProperties are available

#### Scenario: Rich text title
- GIVEN a Title with Rich text
- WHEN Text element is accessed
- THEN Paragraph elements with Run and RunProperties are available

#### Scenario: Legend
- GIVEN a Legend element
- WHEN properties are accessed
- THEN LegendPosition, Layout, Overlay, ShapeProperties, TextProperties are available

### Requirement: Data Labels

The system SHALL provide data label configuration.

#### Scenario: Data labels properties
- GIVEN a DataLabels element
- WHEN properties are accessed
- THEN NumberFormat, ShapeProperties, TextProperties, ShowLegendKey, ShowValue, ShowCategoryName, ShowSeriesName, ShowPercent, ShowBubbleSize, Separator, ShowLeaderLines are available

#### Scenario: Individual data point label
- GIVEN a DataPoint in series
- WHEN DataLabel is accessed
- THEN individual point can have custom label settings

### Requirement: Chart Formatting

The system SHALL support chart formatting using DrawingML.

#### Scenario: Shape properties
- GIVEN a ShapeProperties element
- WHEN properties are accessed
- THEN Fill (SolidFill, GradientFill, PatternFill, NoFill), Outline, EffectList, Scene3D are available

#### Scenario: Text properties
- GIVEN a TextProperties element
- WHEN properties are accessed
- THEN BodyProperties, ListStyle, Paragraph with RunProperties are available

### Requirement: Trendlines

The system SHALL support trendlines on chart series.

#### Scenario: Trendline types
- GIVEN a Trendline element
- WHEN TrendlineType is set
- THEN values Exponential, Linear, Logarithmic, MovingAverage, Polynomial, Power are supported

#### Scenario: Trendline properties
- GIVEN a Trendline element
- WHEN properties are accessed
- THEN Name, ShapeProperties, Order (polynomial), Period (moving avg), Forward, Backward, Intercept, DispRSqr, DispEq, TrendlineLbl are available

### Requirement: Error Bars

The system SHALL support error bars on chart series.

#### Scenario: Error bar types
- GIVEN an ErrorBars element
- WHEN ErrorBarType is set
- THEN values Both, Minus, Plus are supported

#### Scenario: Error bar values
- GIVEN an ErrorBars element
- WHEN ErrorValueType is set
- THEN values Cust, FixedVal, Percentage, StdDev, StdErr are supported

### Requirement: Sparklines

The system SHALL support sparklines (mini-charts in cells).

#### Scenario: Sparkline group
- GIVEN a SparklineGroup element in worksheet extension
- WHEN properties are accessed
- THEN Type (line/column/stacked), SeriesColor, NegativeColor, AxisColor, MarkersColor, FirstMarkerColor, LastMarkerColor, HighMarkerColor, LowMarkerColor, DateAxis, DisplayEmptyCellsAs, Markers, High, Low, First, Last, Negative, DisplayXAxis, DisplayHidden, MinAxisType, MaxAxisType, RightToLeft, Sparklines are available

#### Scenario: Sparkline definition
- GIVEN a Sparkline element
- WHEN properties are accessed
- THEN DataFormula (data range), LocationFormula (cell location) are available

### Requirement: Chart Sheets

The system SHALL support chart-only sheets.

#### Scenario: Chartsheet root
- GIVEN a ChartsheetPart
- WHEN Chartsheet root element is accessed
- THEN SheetProperties, SheetViews, SheetProtection, CustomSheetViews, PageMargins, PageSetup, HeaderFooter, Drawing are available

#### Scenario: Chart in chartsheet
- GIVEN a Chartsheet
- WHEN Drawing is accessed
- THEN the embedded chart is referenced via DrawingsPart
