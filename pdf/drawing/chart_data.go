package drawing

import (
	"fmt"
	"strconv"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// ChartKind identifies the supported chart types.
type ChartKind int

const (
	// ChartKindUnknown represents an unsupported chart type.
	ChartKindUnknown ChartKind = iota
	// ChartKindBar represents a bar/column chart.
	ChartKindBar
	// ChartKindLine represents a line chart.
	ChartKindLine
	// ChartKindPie represents a pie chart.
	ChartKindPie
)

type seriesData struct {
	name       string
	values     []float64
	categories []string
}

// ExtractChartData parses a chart space into renderable chart data.
func ExtractChartData(
	chartSpace *drawingml.ChartSpace,
) (ChartKind, ChartData, bool, error) {
	if chartSpace == nil {
		return ChartKindUnknown, ChartData{}, false, fmt.Errorf(
			"chart space is nil",
		)
	}
	chart := chartSpace.Chart()
	if chart == nil {
		return ChartKindUnknown, ChartData{}, false, fmt.Errorf(
			"chart is nil",
		)
	}
	plotArea := chart.PlotArea()
	if plotArea == nil {
		return ChartKindUnknown, ChartData{}, false, fmt.Errorf(
			"plot area is nil",
		)
	}

	if barElem := plotArea.GetElement(
		"barChart",
		drawingml.NamespaceChart,
	); barElem != nil {
		return parseBarChart(barElem)
	}

	if lineElem := plotArea.GetElement(
		"lineChart",
		drawingml.NamespaceChart,
	); lineElem != nil {
		return parseLineChart(lineElem)
	}

	if pieElem := plotArea.GetElement(
		"pieChart",
		drawingml.NamespaceChart,
	); pieElem != nil {
		return parsePieChart(pieElem)
	}

	return ChartKindUnknown, ChartData{}, false, fmt.Errorf(
		"unsupported chart type",
	)
}

func parseBarChart(
	elem openxml.Element,
) (ChartKind, ChartData, bool, error) {
	barChart := wrapBarChart(elem)
	if barChart == nil {
		return ChartKindUnknown, ChartData{}, false, fmt.Errorf(
			"bar chart is invalid",
		)
	}

	series := extractSeriesFromComposite(barChart)
	data := buildChartData(series)
	horizontal := barChart.Direction() == drawingml.BarDirectionBar

	return ChartKindBar, data, horizontal, nil
}

func parseLineChart(
	elem openxml.Element,
) (ChartKind, ChartData, bool, error) {
	lineChart := wrapLineChart(elem)
	if lineChart == nil {
		return ChartKindUnknown, ChartData{}, false, fmt.Errorf(
			"line chart is invalid",
		)
	}

	series := extractSeriesFromComposite(lineChart)
	data := buildChartData(series)

	return ChartKindLine, data, false, nil
}

func parsePieChart(
	elem openxml.Element,
) (ChartKind, ChartData, bool, error) {
	pieChart := wrapPieChart(elem)
	if pieChart == nil {
		return ChartKindUnknown, ChartData{}, false, fmt.Errorf(
			"pie chart is invalid",
		)
	}

	series := extractSeriesFromComposite(pieChart)
	data := buildChartData(series)

	return ChartKindPie, data, false, nil
}

func wrapBarChart(
	elem openxml.Element,
) *drawingml.BarChart {
	if elem == nil {
		return nil
	}
	if bc, ok := elem.(*drawingml.BarChart); ok {
		return bc
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &drawingml.BarChart{
			CompositeElementBase: comp,
		}
	}

	return nil
}

func wrapLineChart(
	elem openxml.Element,
) *drawingml.LineChart {
	if elem == nil {
		return nil
	}
	if lc, ok := elem.(*drawingml.LineChart); ok {
		return lc
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &drawingml.LineChart{
			CompositeElementBase: comp,
		}
	}

	return nil
}

func wrapPieChart(
	elem openxml.Element,
) *drawingml.PieChart {
	if elem == nil {
		return nil
	}
	if pc, ok := elem.(*drawingml.PieChart); ok {
		return pc
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &drawingml.PieChart{
			CompositeElementBase: comp,
		}
	}

	return nil
}

func extractSeriesFromComposite(
	comp openxml.CompositeElement,
) []seriesData {
	if comp == nil {
		return nil
	}

	series := make([]seriesData, 0)
	for child := range comp.Children() {
		if child.LocalName() != "ser" ||
			child.NamespaceURI() != drawingml.NamespaceChart {
			continue
		}
		series = append(series, parseSeries(child))
	}

	return series
}

func parseSeries(elem openxml.Element) seriesData {
	comp, ok := elem.(openxml.CompositeElement)
	if !ok {
		return seriesData{}
	}

	name := seriesName(comp)
	categories := extractCategories(comp)
	values := extractValues(comp)

	return seriesData{
		name:       name,
		values:     values,
		categories: categories,
	}
}

func seriesName(
	series openxml.CompositeElement,
) string {
	tx := findChildElement(series, "tx")
	if tx == nil {
		return ""
	}

	if val := childText(tx, "v"); val != "" {
		return val
	}

	if strRef := findChildElement(tx, "strRef"); strRef != nil {
		values := readStringCache(strRef)
		if len(values) > 0 {
			return values[0]
		}
	}

	return ""
}

func extractCategories(
	series openxml.CompositeElement,
) []string {
	cat := findChildElement(series, "cat")
	if cat == nil {
		return nil
	}

	if strRef := findChildElement(cat, "strRef"); strRef != nil {
		return readStringCache(strRef)
	}

	if numRef := findChildElement(cat, "numRef"); numRef != nil {
		nums := readNumberCache(numRef)
		cats := make([]string, len(nums))
		for i, val := range nums {
			cats[i] = formatNumber(val)
		}
		return cats
	}

	if multi := findChildElement(cat, "multiLvlStrRef"); multi != nil {
		return readStringCache(multi)
	}

	return nil
}

func extractValues(
	series openxml.CompositeElement,
) []float64 {
	val := findChildElement(series, "val")
	if val == nil {
		return nil
	}

	if numRef := findChildElement(val, "numRef"); numRef != nil {
		return readNumberCache(numRef)
	}

	return nil
}

func findChildElement(
	parent openxml.CompositeElement,
	localName string,
) openxml.CompositeElement {
	if parent == nil {
		return nil
	}
	for child := range parent.Children() {
		if child.LocalName() != localName ||
			child.NamespaceURI() != drawingml.NamespaceChart {
			continue
		}
		if comp, ok := child.(openxml.CompositeElement); ok {
			return comp
		}
		if base, ok := child.(*openxml.CompositeElementBase); ok {
			return base
		}
	}

	return nil
}

func childText(
	parent openxml.CompositeElement,
	localName string,
) string {
	if parent == nil {
		return ""
	}
	for child := range parent.Children() {
		if child.LocalName() != localName ||
			child.NamespaceURI() != drawingml.NamespaceChart {
			continue
		}
		switch node := child.(type) {
		case *openxml.LeafElementBase:
			return node.InnerText()
		case openxml.CompositeElement:
			for nested := range node.Children() {
				if leaf, ok := nested.(*openxml.LeafElementBase); ok {
					return leaf.InnerText()
				}
			}
		}
	}

	return ""
}

func readStringCache(
	parent openxml.CompositeElement,
) []string {
	cache := findChildElement(parent, "strCache")
	if cache == nil {
		return nil
	}

	return readStringPoints(cache)
}

func readNumberCache(
	parent openxml.CompositeElement,
) []float64 {
	cache := findChildElement(parent, "numCache")
	if cache == nil {
		return nil
	}

	points := readStringPoints(cache)
	values := make([]float64, 0, len(points))
	for _, point := range points {
		val, err := strconv.ParseFloat(point, 64)
		if err != nil {
			values = append(values, 0)
			continue
		}
		values = append(values, val)
	}

	return values
}

func readStringPoints(
	cache openxml.CompositeElement,
) []string {
	pointCount := -1
	points := map[int]string{}

	for child := range cache.Children() {
		if child.LocalName() == "ptCount" &&
			child.NamespaceURI() == drawingml.NamespaceChart {
			if attr, found := child.GetAttribute("val", ""); found {
				if count, err := strconv.Atoi(attr.Value()); err == nil {
					pointCount = count
				}
			}
		}

		if child.LocalName() != "pt" ||
			child.NamespaceURI() != drawingml.NamespaceChart {
			continue
		}

		idx := len(points)
		if attr, found := child.GetAttribute("idx", ""); found {
			if parsed, err := strconv.Atoi(attr.Value()); err == nil {
				idx = parsed
			}
		}

		val := ""
		if comp, ok := child.(openxml.CompositeElement); ok {
			val = childText(comp, "v")
		}

		points[idx] = val
	}

	if len(points) == 0 {
		return nil
	}

	maxIndex := -1
	for idx := range points {
		if idx > maxIndex {
			maxIndex = idx
		}
	}

	size := maxIndex + 1
	if pointCount > size {
		size = pointCount
	}

	values := make([]string, size)
	for idx, val := range points {
		if idx < 0 || idx >= size {
			continue
		}
		values[idx] = val
	}

	return values
}

func buildChartData(
	series []seriesData,
) ChartData {
	chartData := ChartData{
		Categories: nil,
		Series:     nil,
	}

	if len(series) == 0 {
		return chartData
	}

	categories := series[0].categories
	if len(categories) == 0 {
		categories = defaultCategories(series)
	}

	colors := chartSeriesColors(len(series))
	chartSeries := make([]ChartSeries, len(series))

	for i, s := range series {
		name := s.name
		if name == "" {
			name = fmt.Sprintf("Series %d", i+1)
		}
		chartSeries[i] = ChartSeries{
			Name:   name,
			Values: s.values,
			Color:  colors[i],
		}
	}

	chartData.Categories = categories
	chartData.Series = chartSeries

	return chartData
}

func chartSeriesColors(
	count int,
) []Color {
	base := []Color{
		NewRGB(0.2, 0.4, 0.8),
		NewRGB(0.8, 0.3, 0.3),
		NewRGB(0.3, 0.7, 0.3),
		NewRGB(0.9, 0.6, 0.2),
		NewRGB(0.6, 0.3, 0.8),
		NewRGB(0.3, 0.8, 0.8),
		NewRGB(0.9, 0.7, 0.3),
		NewRGB(0.8, 0.4, 0.6),
	}

	colors := make([]Color, count)
	for i := range count {
		colors[i] = base[i%len(base)]
	}

	return colors
}

func defaultCategories(
	series []seriesData,
) []string {
	maxLen := 0
	for _, s := range series {
		if len(s.values) > maxLen {
			maxLen = len(s.values)
		}
	}

	if maxLen == 0 {
		return nil
	}

	categories := make([]string, maxLen)
	for i := range categories {
		categories[i] = fmt.Sprintf("Category %d", i+1)
	}

	return categories
}

func formatNumber(value float64) string {
	if value == float64(int64(value)) {
		return fmt.Sprintf("%.0f", value)
	}

	return fmt.Sprintf("%.2f", value)
}
