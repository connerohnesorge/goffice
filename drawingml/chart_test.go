package drawingml

import (
	"strings"
	"testing"
)

func TestChartSpace(t *testing.T) {
	cs := NewChartSpace()

	if cs == nil {
		t.Fatal("NewChartSpace returned nil")
	}

	if cs.LocalName() != "chartSpace" {
		t.Errorf(
			"Expected localName 'chartSpace', got %q",
			cs.LocalName(),
		)
	}

	if cs.NamespaceURI() != NamespaceChart {
		t.Errorf(
			"Expected namespace %q, got %q",
			NamespaceChart,
			cs.NamespaceURI(),
		)
	}

	// Should have a chart child by default
	chart := cs.Chart()
	if chart == nil {
		t.Error(
			"Expected ChartSpace to have a Chart child",
		)
	}
}

func TestChartSpaceDate1904(t *testing.T) {
	cs := NewChartSpace()

	// Default should be false
	if cs.Date1904() {
		t.Error(
			"Expected Date1904 to be false by default",
		)
	}

	// Set to true
	cs.SetDate1904(true)
	if !cs.Date1904() {
		t.Error(
			"Expected Date1904 to be true after setting",
		)
	}

	// Set back to false
	cs.SetDate1904(false)
	if cs.Date1904() {
		t.Error(
			"Expected Date1904 to be false after unsetting",
		)
	}
}

func TestChartSpaceRoundedCorners(t *testing.T) {
	cs := NewChartSpace()

	cs.SetRoundedCorners(true)
	if !cs.RoundedCorners() {
		t.Error(
			"Expected RoundedCorners to be true",
		)
	}

	cs.SetRoundedCorners(false)
	if cs.RoundedCorners() {
		t.Error(
			"Expected RoundedCorners to be false",
		)
	}
}

func TestChart(t *testing.T) {
	c := NewChart()

	if c == nil {
		t.Fatal("NewChart returned nil")
	}

	if c.LocalName() != "chart" {
		t.Errorf(
			"Expected localName 'chart', got %q",
			c.LocalName(),
		)
	}

	// Should have a plot area by default
	plotArea := c.PlotArea()
	if plotArea == nil {
		t.Error(
			"Expected Chart to have a PlotArea child",
		)
	}
}

func TestChartTitle(t *testing.T) {
	c := NewChart()

	// No title by default
	if c.Title() != nil {
		t.Error("Expected no title by default")
	}

	// Set a title
	title := NewTitleWithText("Test Chart")
	c.SetTitle(title)

	if c.Title() == nil {
		t.Error("Expected title after setting")
	}
}

func TestChartDisplayBlanksAs(t *testing.T) {
	c := NewChart()

	c.SetDisplayBlanksAs(DisplayBlanksAsZero)
	if c.DisplayBlanksAs() != DisplayBlanksAsZero {
		t.Errorf(
			"Expected DisplayBlanksAsZero, got %q",
			c.DisplayBlanksAs(),
		)
	}

	c.SetDisplayBlanksAs(DisplayBlanksAsSpan)
	if c.DisplayBlanksAs() != DisplayBlanksAsSpan {
		t.Errorf(
			"Expected DisplayBlanksAsSpan, got %q",
			c.DisplayBlanksAs(),
		)
	}
}

func TestPlotArea(t *testing.T) {
	pa := NewPlotArea()

	if pa == nil {
		t.Fatal("NewPlotArea returned nil")
	}

	if pa.LocalName() != "plotArea" {
		t.Errorf(
			"Expected localName 'plotArea', got %q",
			pa.LocalName(),
		)
	}
}

func TestBarChart(t *testing.T) {
	bc := NewBarChart(
		BarDirectionCol,
		BarGroupingClustered,
	)

	if bc == nil {
		t.Fatal("NewBarChart returned nil")
	}

	if bc.LocalName() != "barChart" {
		t.Errorf(
			"Expected localName 'barChart', got %q",
			bc.LocalName(),
		)
	}

	if bc.Direction() != BarDirectionCol {
		t.Errorf(
			"Expected direction %q, got %q",
			BarDirectionCol,
			bc.Direction(),
		)
	}

	if bc.Grouping() != BarGroupingClustered {
		t.Errorf(
			"Expected grouping %q, got %q",
			BarGroupingClustered,
			bc.Grouping(),
		)
	}

	// Add a series
	ser := bc.AddSeries(0, 0)
	if ser == nil {
		t.Error("AddSeries returned nil")
	}
}

func TestBarChartHorizontal(t *testing.T) {
	bc := NewBarChart(
		BarDirectionBar,
		BarGroupingStacked,
	)

	if bc.Direction() != BarDirectionBar {
		t.Errorf(
			"Expected direction %q, got %q",
			BarDirectionBar,
			bc.Direction(),
		)
	}

	if bc.Grouping() != BarGroupingStacked {
		t.Errorf(
			"Expected grouping %q, got %q",
			BarGroupingStacked,
			bc.Grouping(),
		)
	}
}

func TestLineChart(t *testing.T) {
	lc := NewLineChart(GroupingStandard)

	if lc == nil {
		t.Fatal("NewLineChart returned nil")
	}

	if lc.LocalName() != "lineChart" {
		t.Errorf(
			"Expected localName 'lineChart', got %q",
			lc.LocalName(),
		)
	}

	if lc.Grouping() != GroupingStandard {
		t.Errorf(
			"Expected grouping %q, got %q",
			GroupingStandard,
			lc.Grouping(),
		)
	}
}

func TestPieChart(t *testing.T) {
	pc := NewPieChart()

	if pc == nil {
		t.Fatal("NewPieChart returned nil")
	}

	if pc.LocalName() != "pieChart" {
		t.Errorf(
			"Expected localName 'pieChart', got %q",
			pc.LocalName(),
		)
	}

	// Add a series
	ser := pc.AddSeries(0, 0)
	if ser == nil {
		t.Error("AddSeries returned nil")
	}
}

func TestScatterChart(t *testing.T) {
	sc := NewScatterChart(ScatterStyleLineMarker)

	if sc == nil {
		t.Fatal("NewScatterChart returned nil")
	}

	if sc.LocalName() != "scatterChart" {
		t.Errorf(
			"Expected localName 'scatterChart', got %q",
			sc.LocalName(),
		)
	}
}

func TestDoughnutChart(t *testing.T) {
	dc := NewDoughnutChart()

	if dc == nil {
		t.Fatal("NewDoughnutChart returned nil")
	}

	if dc.LocalName() != "doughnutChart" {
		t.Errorf(
			"Expected localName 'doughnutChart', got %q",
			dc.LocalName(),
		)
	}

	dc.SetHoleSize(50)
}

func TestRadarChart(t *testing.T) {
	rc := NewRadarChart(RadarStyleFilled)

	if rc == nil {
		t.Fatal("NewRadarChart returned nil")
	}

	if rc.LocalName() != "radarChart" {
		t.Errorf(
			"Expected localName 'radarChart', got %q",
			rc.LocalName(),
		)
	}
}

func TestAreaChart(t *testing.T) {
	ac := NewAreaChart(GroupingPercentStacked)

	if ac == nil {
		t.Fatal("NewAreaChart returned nil")
	}

	if ac.LocalName() != "areaChart" {
		t.Errorf(
			"Expected localName 'areaChart', got %q",
			ac.LocalName(),
		)
	}
}

func TestBubbleChart(t *testing.T) {
	bc := NewBubbleChart()

	if bc == nil {
		t.Fatal("NewBubbleChart returned nil")
	}

	if bc.LocalName() != "bubbleChart" {
		t.Errorf(
			"Expected localName 'bubbleChart', got %q",
			bc.LocalName(),
		)
	}
}

func TestStockChart(t *testing.T) {
	sc := NewStockChart()

	if sc == nil {
		t.Fatal("NewStockChart returned nil")
	}

	if sc.LocalName() != "stockChart" {
		t.Errorf(
			"Expected localName 'stockChart', got %q",
			sc.LocalName(),
		)
	}
}

func TestSurfaceChart(t *testing.T) {
	sc := NewSurfaceChart()

	if sc == nil {
		t.Fatal("NewSurfaceChart returned nil")
	}

	if sc.LocalName() != "surfaceChart" {
		t.Errorf(
			"Expected localName 'surfaceChart', got %q",
			sc.LocalName(),
		)
	}
}

func TestBarChartSeries(t *testing.T) {
	ser := NewBarChartSeries(1, 2)

	if ser == nil {
		t.Fatal("NewBarChartSeries returned nil")
	}

	if ser.Index() != 1 {
		t.Errorf(
			"Expected index 1, got %d",
			ser.Index(),
		)
	}

	if ser.Order() != 2 {
		t.Errorf(
			"Expected order 2, got %d",
			ser.Order(),
		)
	}

	// Set series text
	serText := NewSeriesTextWithValue("Sales")
	ser.SetSeriesText(serText)

	// Set category data
	cat := NewCategoryAxisData()
	cat.SetStringReference("Sheet1!$A$2:$A$5")
	ser.SetCategoryAxisData(cat)

	// Set values
	vals := NewValues()
	vals.SetNumberReference("Sheet1!$B$2:$B$5")
	ser.SetValues(vals)
}

func TestLineChartSeries(t *testing.T) {
	ser := NewLineChartSeries(0, 0)

	if ser == nil {
		t.Fatal("NewLineChartSeries returned nil")
	}

	if ser.Index() != 0 {
		t.Errorf(
			"Expected index 0, got %d",
			ser.Index(),
		)
	}

	// Set smooth
	ser.SetSmooth(true)

	// Set marker
	marker := NewMarkerWithStyle(
		MarkerStyleCircle,
	)
	marker.SetSize(7)
	ser.SetMarker(marker)
}

func TestScatterChartSeries(t *testing.T) {
	ser := NewScatterChartSeries(0, 0)

	if ser == nil {
		t.Fatal(
			"NewScatterChartSeries returned nil",
		)
	}

	// Set X values
	xVals := NewXValues()
	xVals.SetNumberReference("Sheet1!$A$2:$A$10")
	ser.SetXValues(xVals)

	// Set Y values
	yVals := NewYValues()
	yVals.SetNumberReference("Sheet1!$B$2:$B$10")
	ser.SetYValues(yVals)
}

func TestCategoryAxis(t *testing.T) {
	ax := NewCategoryAxis(1, 2)

	if ax == nil {
		t.Fatal("NewCategoryAxis returned nil")
	}

	if ax.LocalName() != "catAx" {
		t.Errorf(
			"Expected localName 'catAx', got %q",
			ax.LocalName(),
		)
	}

	if ax.AxisID() != 1 {
		t.Errorf(
			"Expected axis ID 1, got %d",
			ax.AxisID(),
		)
	}

	ax.SetAxisPosition(AxisPositionTop)
	ax.SetMajorTickMark(TickMarkOut)
	ax.SetMinorTickMark(TickMarkNone)
	ax.SetTickLabelPosition(
		TickLabelPositionNextTo,
	)
}

func TestValueAxis(t *testing.T) {
	ax := NewValueAxis(2, 1)

	if ax == nil {
		t.Fatal("NewValueAxis returned nil")
	}

	if ax.LocalName() != "valAx" {
		t.Errorf(
			"Expected localName 'valAx', got %q",
			ax.LocalName(),
		)
	}

	if ax.AxisID() != 2 {
		t.Errorf(
			"Expected axis ID 2, got %d",
			ax.AxisID(),
		)
	}

	ax.SetAxisPosition(AxisPositionRight)
	ax.SetMajorGridlines(true)
	ax.SetNumberFormat("0.00", true)
}

func TestDateAxis(t *testing.T) {
	ax := NewDateAxis(3, 4)

	if ax == nil {
		t.Fatal("NewDateAxis returned nil")
	}

	if ax.AxisID() != 3 {
		t.Errorf(
			"Expected axis ID 3, got %d",
			ax.AxisID(),
		)
	}
}

func TestSeriesAxis(t *testing.T) {
	ax := NewSeriesAxis(5, 6)

	if ax == nil {
		t.Fatal("NewSeriesAxis returned nil")
	}

	if ax.AxisID() != 5 {
		t.Errorf(
			"Expected axis ID 5, got %d",
			ax.AxisID(),
		)
	}
}

func TestScaling(t *testing.T) {
	s := NewScaling()

	if s == nil {
		t.Fatal("NewScaling returned nil")
	}

	s.SetOrientation(OrientationMaxMin)
	s.SetMinimum(0)
	s.SetMaximum(100)
}

func TestLegend(t *testing.T) {
	l := NewLegendWithPosition(
		LegendPositionBottom,
	)

	if l == nil {
		t.Fatal(
			"NewLegendWithPosition returned nil",
		)
	}

	if l.Position() != LegendPositionBottom {
		t.Errorf(
			"Expected position %q, got %q",
			LegendPositionBottom,
			l.Position(),
		)
	}

	l.SetPosition(LegendPositionRight)
	if l.Position() != LegendPositionRight {
		t.Errorf(
			"Expected position %q, got %q",
			LegendPositionRight,
			l.Position(),
		)
	}

	l.SetOverlay(false)
}

func TestDataLabels(t *testing.T) {
	dl := NewDataLabels()

	if dl == nil {
		t.Fatal("NewDataLabels returned nil")
	}

	dl.SetShowValue(true)
	dl.SetShowCategoryName(false)
	dl.SetShowSeriesName(true)
	dl.SetShowPercent(false)
	dl.SetShowLegendKey(false)
}

func TestMarker(t *testing.T) {
	m := NewMarkerWithStyle(MarkerStyleDiamond)

	if m == nil {
		t.Fatal("NewMarkerWithStyle returned nil")
	}

	m.SetStyle(MarkerStyleSquare)
	m.SetSize(10)
}

func TestStringReference(t *testing.T) {
	sr := NewStringReference("Sheet1!$A$1:$A$10")

	if sr == nil {
		t.Fatal("NewStringReference returned nil")
	}

	if sr.Formula() != "Sheet1!$A$1:$A$10" {
		t.Errorf(
			"Expected formula 'Sheet1!$A$1:$A$10', got %q",
			sr.Formula(),
		)
	}

	sr.SetFormula("Sheet2!$B$1:$B$5")
	if sr.Formula() != "Sheet2!$B$1:$B$5" {
		t.Errorf(
			"Expected formula 'Sheet2!$B$1:$B$5', got %q",
			sr.Formula(),
		)
	}
}

func TestNumberReference(t *testing.T) {
	nr := NewNumberReference("Sheet1!$B$1:$B$10")

	if nr == nil {
		t.Fatal("NewNumberReference returned nil")
	}

	if nr.Formula() != "Sheet1!$B$1:$B$10" {
		t.Errorf(
			"Expected formula 'Sheet1!$B$1:$B$10', got %q",
			nr.Formula(),
		)
	}
}

func TestSeriesText(t *testing.T) {
	// Test with literal value
	st := NewSeriesTextWithValue("My Series")
	if st == nil {
		t.Fatal(
			"NewSeriesTextWithValue returned nil",
		)
	}

	// Test with reference
	st2 := NewSeriesTextWithReference(
		"Sheet1!$A$1",
	)
	if st2 == nil {
		t.Fatal(
			"NewSeriesTextWithReference returned nil",
		)
	}
}

func TestChartShapeProperties(t *testing.T) {
	props := NewChartShapeProperties()

	if props == nil {
		t.Fatal(
			"NewChartShapeProperties returned nil",
		)
	}

	props.SetSolidFill("FF0000")
	props.SetOutline(12700, "000000")
}

func TestDataPoint(t *testing.T) {
	dp := NewDataPoint(5)

	if dp == nil {
		t.Fatal("NewDataPoint returned nil")
	}

	props := NewChartShapeProperties()
	props.SetSolidFill("00FF00")
	dp.SetShapeProperties(props)
}

func TestPlotAreaAddCharts(t *testing.T) {
	pa := NewPlotArea()

	// Add various chart types
	bc := pa.AddBarChart(
		BarDirectionCol,
		BarGroupingClustered,
	)
	if bc == nil {
		t.Error("AddBarChart returned nil")
	}

	lc := pa.AddLineChart(GroupingStandard)
	if lc == nil {
		t.Error("AddLineChart returned nil")
	}

	pc := pa.AddPieChart()
	if pc == nil {
		t.Error("AddPieChart returned nil")
	}

	ac := pa.AddAreaChart(GroupingStacked)
	if ac == nil {
		t.Error("AddAreaChart returned nil")
	}

	sc := pa.AddScatterChart(ScatterStyleMarker)
	if sc == nil {
		t.Error("AddScatterChart returned nil")
	}

	dc := pa.AddDoughnutChart()
	if dc == nil {
		t.Error("AddDoughnutChart returned nil")
	}

	rc := pa.AddRadarChart(RadarStyleStandard)
	if rc == nil {
		t.Error("AddRadarChart returned nil")
	}

	bbc := pa.AddBubbleChart()
	if bbc == nil {
		t.Error("AddBubbleChart returned nil")
	}

	stc := pa.AddStockChart()
	if stc == nil {
		t.Error("AddStockChart returned nil")
	}

	sfc := pa.AddSurfaceChart()
	if sfc == nil {
		t.Error("AddSurfaceChart returned nil")
	}
}

func TestPlotAreaAddAxes(t *testing.T) {
	pa := NewPlotArea()

	catAx := pa.AddCategoryAxis(1, 2)
	if catAx == nil {
		t.Error("AddCategoryAxis returned nil")
	}

	valAx := pa.AddValueAxis(2, 1)
	if valAx == nil {
		t.Error("AddValueAxis returned nil")
	}

	dateAx := pa.AddDateAxis(3, 4)
	if dateAx == nil {
		t.Error("AddDateAxis returned nil")
	}

	serAx := pa.AddSeriesAxis(5, 6)
	if serAx == nil {
		t.Error("AddSeriesAxis returned nil")
	}
}

func TestChartClone(t *testing.T) {
	cs := NewChartSpace()
	cs.SetRoundedCorners(true)

	chart := cs.Chart()
	pa := chart.PlotArea()
	bc := pa.AddBarChart(
		BarDirectionCol,
		BarGroupingClustered,
	)
	ser := bc.AddSeries(0, 0)
	_ = ser

	// Clone the chart space
	cloned := cs.Clone()
	clone, ok := cloned.(*ChartSpace)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *ChartSpace",
			cloned,
		)
	}

	if clone == cs {
		t.Error(
			"Clone should return a different instance",
		)
	}

	if !clone.RoundedCorners() {
		t.Error(
			"Cloned ChartSpace should have roundedCorners=true",
		)
	}
}

func TestChartXMLOutput(t *testing.T) {
	cs := NewChartSpace()
	chart := cs.Chart()

	// Set title
	title := NewTitleWithText("Sales Report")
	chart.SetTitle(title)

	// Add a bar chart
	pa := chart.PlotArea()
	bc := pa.AddBarChart(
		BarDirectionCol,
		BarGroupingClustered,
	)

	// Add a series
	ser := bc.AddSeries(0, 0)
	serText := NewSeriesTextWithValue(
		"2024 Sales",
	)
	ser.SetSeriesText(serText)

	// Add axis IDs
	bc.AddAxisID(1)
	bc.AddAxisID(2)

	// Add axes
	pa.AddCategoryAxis(1, 2)
	pa.AddValueAxis(2, 1)

	// Add legend
	legend := NewLegendWithPosition(
		LegendPositionRight,
	)
	chart.SetLegend(legend)

	// Get XML output
	xml := cs.OuterXml()

	// Verify it contains expected elements
	expectedElements := []string{
		"c:chartSpace",
		"c:chart",
		"c:plotArea",
		"c:barChart",
		"c:barDir",
		"c:grouping",
		"c:ser",
		"c:idx",
		"c:order",
		"c:tx",
		"c:catAx",
		"c:valAx",
		"c:legend",
		"c:title",
	}

	for _, elem := range expectedElements {
		if !strings.Contains(xml, elem) {
			t.Errorf(
				"Expected XML to contain %q",
				elem,
			)
		}
	}
}

func TestEnumValues(t *testing.T) {
	// Test that enum values are correct strings
	tests := []struct {
		val      string
		expected string
	}{
		{string(BarDirectionBar), "bar"},
		{string(BarDirectionCol), "col"},
		{
			string(BarGroupingClustered),
			"clustered",
		},
		{string(BarGroupingStacked), "stacked"},
		{string(GroupingStandard), "standard"},
		{
			string(GroupingPercentStacked),
			"percentStacked",
		},
		{
			string(ScatterStyleLineMarker),
			"lineMarker",
		},
		{string(MarkerStyleCircle), "circle"},
		{string(MarkerStyleNone), "none"},
		{string(LegendPositionBottom), "b"},
		{string(LegendPositionRight), "r"},
		{string(AxisPositionLeft), "l"},
		{string(AxisPositionBottom), "b"},
		{string(TickMarkNone), "none"},
		{string(TickMarkCross), "cross"},
		{
			string(TickLabelPositionNextTo),
			"nextTo",
		},
		{string(DisplayBlanksAsGap), "gap"},
		{string(DisplayBlanksAsZero), "zero"},
		{string(CrossesAutoZero), "autoZero"},
		{string(OrientationMinMax), "minMax"},
		{string(RadarStyleFilled), "filled"},
		{string(OfPieTypePie), "pie"},
	}

	for _, tt := range tests {
		if tt.val != tt.expected {
			t.Errorf(
				"Expected %q, got %q",
				tt.expected,
				tt.val,
			)
		}
	}
}
