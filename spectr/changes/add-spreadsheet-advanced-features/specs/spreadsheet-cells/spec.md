## ADDED Requirements

### Requirement: Data Validation
The system SHALL support multiple types of data validation on cell ranges.

#### Scenario: Apply list validation
- GIVEN a cell or range
- WHEN validation.SetListValidation([]string{"Option1", "Option2", "Option3"}) is called
- THEN dropdown validation is applied
- AND cells show dropdown arrows with options

#### Scenario: Apply list validation from range
- GIVEN a cell range for validation
- WHEN validation.SetListValidationRange("$Sheet1.$A$1:$A$10") is called
- THEN validation references external range
- AND dropdown shows items from referenced range

#### Scenario: Apply whole number validation
- GIVEN a cell range
- WHEN validation.SetWholeNumberValidation(Operator.Between, 1, 100) is called
- THEN validation accepts whole numbers 1-100
- AND other operators (GreaterThan, LessThan, Equal, NotEqual) are supported

#### Scenario: Apply decimal validation
- GIVEN a cell range
- WHEN validation.SetDecimalValidation(Operator.LessThanOrEqual, 99.99) is called
- THEN decimal numbers are validated
- AND comparison operators work with decimal precision

#### Scenario: Apply date validation
- GIVEN a cell range
- WHEN validation.SetDateValidation(Operator.Between, date1, date2) is called
- THEN date range validation is applied
- AND dates outside range are rejected

#### Scenario: Apply custom formula validation
- GIVEN a cell range
- WHEN validation.SetCustomValidation("MOD(A1,2)=0") is called
- THEN custom formula is evaluated
- AND cell is valid if formula returns true

#### Scenario: Set validation error message
- GIVEN a validation rule
- WHEN validation.SetErrorMessage(ErrorStyle.Stop, "Invalid Input", "Please enter value between 1 and 100") are called
- THEN error alert is shown on invalid entry
- AND error style options (Stop, Warning, Information) are supported

#### Scenario: Set validation input message
- GIVEN a validation rule
- WHEN validation.SetInputMessage("Enter a value", "Please enter a number from 1 to 100") are called
- THEN tooltip appears when cell is selected
- AND message guides user input

### Requirement: Conditional Formatting
The system SHALL support multiple conditional formatting rule types.

#### Scenario: Apply 3-color scale
- GIVEN a range of numeric cells
- WHEN formatting.SetColorScale3(minColor=Color.Blue, midColor=Color.White, maxColor=Color.Red) is called
- THEN colors are applied based on cell values
- AND low values are blue, mid values white, high values red

#### Scenario: Apply data bars
- GIVEN a range of numeric cells
- WHEN formatting.SetDataBar(Color.Blue, gradient=true) is called
- THEN colored bars appear in cells
- AND bar length is proportional to value

#### Scenario: Apply icon set (5-icon)
- GIVEN a range of cells
- WHEN formatting.SetIconSet(IconSet.FiveQuarters) is called
- THEN icons appear based on value ranges
- AND 5-quarter, 4-traffic, 3-arrow sets are supported

#### Scenario: Apply formula-based conditional formatting
- GIVEN a range
- WHEN formatting.SetFormulaRule("$A1>100", Color.Red) is called
- THEN rule applies to cells matching formula
- AND cell references are evaluated relatively

#### Scenario: Apply top/bottom conditional formatting
- GIVEN a range
- WHEN formatting.SetTopBottomRule(RuleType.Top10Percent, Color.Green) is called
- THEN cells in top 10% are formatted
- AND bottom 10% option is also supported

#### Scenario: Apply duplicate value formatting
- GIVEN a range
- WHEN formatting.SetDuplicateValue(Color.Yellow) is called
- THEN duplicate values are highlighted
- AND unique values are not formatted

#### Scenario: Apply above/below average formatting
- GIVEN a range with numeric data
- WHEN formatting.SetAboveAverageRule(Color.Orange) is called
- THEN values above average are formatted
- AND below average option is supported

#### Scenario: Apply text contains formatting
- GIVEN a range with text
- WHEN formatting.SetTextContainsRule("error", Color.Red) is called
- THEN cells containing "error" text are formatted
- AND case sensitivity is configurable

### Requirement: Filtering and Sorting
The system SHALL support auto-filter with multiple filtering options.

#### Scenario: Apply auto-filter to range
- GIVEN a data range with headers
- WHEN sheet.ApplyAutoFilter(firstRow=1, lastRow=100, firstCol='A', lastCol='D') is called
- THEN filter dropdown arrows appear in header row
- AND filter can be applied to each column

#### Scenario: Filter column by single value
- GIVEN a filtered range
- WHEN filter.FilterColumn('A', "Value1") is called
- THEN only rows with "Value1" in column A are shown
- AND other rows are hidden

#### Scenario: Filter column by multiple values
- GIVEN a filtered range
- WHEN filter.FilterColumn('B', []string{"Option1", "Option2"}) is called
- THEN only rows with Option1 or Option2 are shown
- AND checkboxes in filter dialog reflect selections

#### Scenario: Filter by custom date
- GIVEN a date column with filter
- WHEN filter.FilterByDate('C', DateFilter.ThisWeek) is called
- THEN rows from this week are shown
- AND predefined date ranges (Today, Yesterday, This Month, etc.) are supported

#### Scenario: Filter by numeric criteria
- GIVEN a numeric column
- WHEN filter.FilterByNumeric('D', NumericFilter.Top10) is called
- THEN top 10 values are shown
- AND average, above average options work

#### Scenario: Sort range primary key
- GIVEN a data range
- WHEN sort.AddSortKey(1, 'A', SortOrder.Ascending) is called
- THEN range is sorted by column A ascending
- AND secondary and tertiary keys can be added

#### Scenario: Sort with multiple keys
- GIVEN sort configuration
- WHEN sort.AddSortKey(2, 'B', SortOrder.Descending) is called after first key
- THEN range is sorted by multiple columns
- AND sort stability is maintained

### Requirement: Sparklines
The system SHALL support inline chart sparklines.

#### Scenario: Create line sparkline
- GIVEN cells with data range and sparkline cell
- WHEN sheet.AddLineSparkline(dataRange="$A$1:$A$10", sparklineCell="$B$1") is called
- THEN sparkline chart appears in cell
- AND chart shows line trend

#### Scenario: Create column sparkline
- GIVEN data range and sparkline cell
- WHEN sheet.AddColumnSparkline(dataRange, sparklineCell) is called
- THEN column chart appears in cell
- AND columns represent individual values

#### Scenario: Create win/loss sparkline
- GIVEN binary data (+1, 0, -1 or positive/negative values)
- WHEN sheet.AddWinLossSparkline(dataRange, sparklineCell) is called
- THEN up/down pattern is displayed
- AND colors indicate wins/losses

#### Scenario: Style sparkline
- GIVEN a sparkline
- WHEN sparkline.SetColor(Color.Blue) and SetLineWeight(1.5) are called
- THEN sparkline appearance is customized
- AND style persists in roundtrip

### Requirement: Named Ranges
The system SHALL support workbook and sheet-level named ranges.

#### Scenario: Create named range at workbook level
- GIVEN a cell range and name
- WHEN document.AddNamedRange("SalesData", "Sheet1!$A$1:$D$100") is called
- THEN named range is created
- AND name is accessible from any sheet

#### Scenario: Create named range at sheet level
- GIVEN a cell range and sheet
- WHEN sheet.AddNamedRange("LocalData", "$A$1:$A$50") is called
- THEN named range is created
- AND name is scoped to this sheet only

#### Scenario: Resolve named range reference
- GIVEN a named range name
- WHEN document.GetNamedRange("SalesData") is called
- THEN the range address and sheet reference are returned
- AND range can be used in formulas

#### Scenario: Create named formula
- GIVEN a formula expression
- WHEN document.AddNamedFormula("MyFormula", "Sheet1.A1:A100") is called
- THEN formula is stored as named definition
- AND formula can be referenced by name

### Requirement: Advanced Chart Types
The system SHALL support additional chart types beyond basic charts.

#### Scenario: Create waterfall chart
- GIVEN data with values and totals
- WHEN sheet.AddWaterfallChart(dataRange, chartName) is called
- THEN waterfall chart is created
- AND columns show increases, decreases, and totals

#### Scenario: Create funnel chart
- GIVEN sequential data
- WHEN sheet.AddFunnelChart(dataRange, chartName) is called
- THEN funnel chart is created
- AND segments decrease in size

#### Scenario: Create sunburst chart
- GIVEN hierarchical data
- WHEN sheet.AddSunburstChart(dataRange, chartName) is called
- THEN sunburst chart is created
- AND hierarchy is represented in rings

#### Scenario: Create treemap chart
- GIVEN hierarchical data with values
- WHEN sheet.AddTreemapChart(dataRange, chartName) is called
- THEN treemap chart is created
- AND rectangles represent hierarchical values

#### Scenario: Create combo chart with 2 Y-axes
- GIVEN multiple data series with different scales
- WHEN sheet.AddComboChart(series1Range, series2Range, chartName) is called
- THEN combo chart has 2 Y-axes
- AND series can have different chart types (column, line)

### Requirement: Advanced Chart Formatting
The system SHALL support detailed chart customization.

#### Scenario: Add axis labels
- GIVEN a chart
- WHEN chart.SetAxisLabels(AxisType.XAxis, LabelPosition.Low) is called
- THEN axis labels are displayed
- AND label position options (Low, High, NextToAxis) are supported

#### Scenario: Add trend line
- GIVEN a chart data series
- WHEN series.AddTrendLine(TrendLineType.Linear, forecast=2) is called
- THEN trend line is drawn
- AND forecast periods extend line forward

#### Scenario: Add error bars
- GIVEN a chart data series
- WHEN series.AddErrorBars(ErrorBarType.StandardDeviation, value=1) is called
- THEN error bars are displayed
- AND bar types (Fixed, Percentage, StandardDeviation) are supported

#### Scenario: Configure chart legend
- GIVEN a chart
- WHEN chart.SetLegend(LegendPosition.Right, includeInLayout=true) is called
- THEN legend is positioned and styled
- AND legend entries match data series

### Requirement: Slicers and Timelines
The system SHALL support pivot table slicers and timelines.

#### Scenario: Create slicer for pivot table
- GIVEN a pivot table
- WHEN pivotTable.AddSlicer(fieldName="Region", slicerName="RegionSlicer") is called
- THEN slicer is created with buttons
- AND clicking buttons filters pivot table

#### Scenario: Create timeline for pivot table
- GIVEN a pivot table with date field
- WHEN pivotTable.AddTimeline(dateFieldName="Date", timelineName="DateTimeline") is called
- THEN timeline is created with date range
- AND timeline filtering updates pivot table

### Requirement: Cell Comments
The system SHALL support cell comments with threading.

#### Scenario: Add comment to cell
- GIVEN a cell
- WHEN cell.AddComment("Author Name", "Comment text") is called
- THEN comment is attached to cell
- AND comment marker appears in cell corner

#### Scenario: Reply to comment
- GIVEN a cell comment
- WHEN comment.Reply("Reply Author", "Reply text") is called
- THEN reply comment is threaded
- AND reply appears under original comment

#### Scenario: Access comment metadata
- GIVEN a comment
- WHEN comment.Author(), comment.Date(), comment.Resolved() are called
- THEN metadata fields are returned
- AND resolved status can be toggled

### Requirement: Merged Cells
The system SHALL provide utilities for detecting and managing merged cells.

#### Scenario: Detect merged cell
- GIVEN a range with merged cells
- WHEN sheet.GetMergedCells().IsCellMerged("A1") is called
- THEN true is returned if cell is in merge
- AND merge range can be retrieved

#### Scenario: Get merge range
- GIVEN a merged cell
- WHEN sheet.GetMergeRange("A1:A5") is called
- THEN merge range address is returned
- AND merge start and end can be accessed
