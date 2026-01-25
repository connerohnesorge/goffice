// This file implements chart series functionality.
package spreadsheet

// ChartSeries represents a data series in a chart.
type ChartSeries struct {
	chart       *Chart
	name        string
	valuesRange string
	catRange    string
	index       uint32
}

// newChartSeries creates a new ChartSeries.
func newChartSeries(
	chart *Chart,
	name, valuesRange string,
) *ChartSeries {
	// Get the next index
	index := uint32(len(chart.series))

	return &ChartSeries{
		chart:       chart,
		name:        name,
		valuesRange: valuesRange,
		index:       index,
	}
}

// Name returns the series name.
func (s *ChartSeries) Name() string {
	return s.name
}

// SetName sets the series name.
func (s *ChartSeries) SetName(name string) {
	s.name = name
	s.updateUnderlyingElement()
}

// ValuesRange returns the values range reference.
func (s *ChartSeries) ValuesRange() string {
	return s.valuesRange
}

// SetValuesRange sets the values range reference.
func (s *ChartSeries) SetValuesRange(
	rangeRef string,
) {
	s.valuesRange = rangeRef
	s.updateUnderlyingElement()
}

// CategoriesRange returns the categories range reference.
func (s *ChartSeries) CategoriesRange() string {
	return s.catRange
}

// SetCategoriesRange sets the categories range reference.
func (s *ChartSeries) SetCategoriesRange(
	rangeRef string,
) {
	s.catRange = rangeRef
	s.updateUnderlyingElement()
}

// updateUnderlyingElement updates the underlying XML series element.
// This is a placeholder - actual implementation would need to access
// the chart's underlying XML structure and update the appropriate series.
func (*ChartSeries) updateUnderlyingElement() {
	// This method would need to:
	// 1. Get the chart's underlying plot area
	// 2. Find the appropriate chart type element (barChart, lineChart, etc.)
	// 3. Add or update the series element with the current values
	//
	// For now, this is a placeholder that demonstrates the pattern.
	// Full implementation would require more complex XML manipulation.
}
