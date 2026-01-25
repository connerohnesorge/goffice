package drawing

import (
	"fmt"
	"math"

	"github.com/connerohnesorge/goffice-pdf/core"
)

// ChartRenderer handles rendering of DrawingML charts.
type ChartRenderer struct {
	ctx *core.RenderingContext
}

// NewChartRenderer creates a new chart renderer.
func NewChartRenderer(
	ctx *core.RenderingContext,
) *ChartRenderer {
	return &ChartRenderer{ctx: ctx}
}

// ChartData represents data for chart rendering.
type ChartData struct {
	Categories []string
	Series     []ChartSeries
}

// ChartSeries represents a data series in a chart.
type ChartSeries struct {
	Name   string
	Values []float64
	Color  Color
}

// RenderBarChart renders a bar/column chart.
func (r *ChartRenderer) RenderBarChart(
	x, y, width, height float64,
	data ChartData,
	horizontal bool,
) error {
	if len(data.Series) == 0 {
		return fmt.Errorf("no data series")
	}

	// Calculate chart area (leave room for axes and labels)
	chartX := x + 40
	chartY := y + 20
	chartWidth := width - 60
	chartHeight := height - 60

	// Draw axes
	r.drawAxes(
		chartX,
		chartY,
		chartWidth,
		chartHeight,
	)

	// Calculate bar dimensions
	categoryCount := len(data.Categories)
	if categoryCount == 0 {
		categoryCount = len(data.Series[0].Values)
	}

	seriesCount := len(data.Series)
	barSpacing := 2.0

	var barWidth, barGroupWidth float64
	if horizontal {
		barGroupWidth = chartHeight / float64(
			categoryCount,
		)
		barWidth = (barGroupWidth - barSpacing*float64(seriesCount+1)) / float64(
			seriesCount,
		)
	} else {
		barGroupWidth = chartWidth / float64(categoryCount)
		barWidth = (barGroupWidth - barSpacing*float64(seriesCount+1)) / float64(seriesCount)
	}

	// Find max value for scaling
	maxValue := 0.0
	for _, series := range data.Series {
		for _, val := range series.Values {
			if val > maxValue {
				maxValue = val
			}
		}
	}

	// Render bars for each series
	for seriesIdx, series := range data.Series {
		r.ctx.Page.SetFillColor(
			series.Color.R,
			series.Color.G,
			series.Color.B,
		)

		for catIdx, val := range series.Values {
			if catIdx >= categoryCount {
				break
			}

			var barX, barY, barW, barH float64

			if horizontal {
				// Horizontal bar chart
				barLength := (val / maxValue) * chartWidth
				barY = chartY + chartHeight - float64(
					catIdx+1,
				)*barGroupWidth + barSpacing
				barY += float64(
					seriesIdx,
				) * (barWidth + barSpacing)
				barX = chartX
				barW = barLength
				barH = barWidth
			} else {
				// Vertical column chart
				barHeight := (val / maxValue) * chartHeight
				barX = chartX + float64(catIdx)*barGroupWidth + barSpacing
				barX += float64(seriesIdx) * (barWidth + barSpacing)
				barY = chartY + chartHeight - barHeight
				barW = barWidth
				barH = barHeight
			}

			r.ctx.Page.DrawRectangle(
				barX,
				barY,
				barW,
				barH,
				true,
				true,
			)
		}
	}

	// Render legend
	r.renderLegend(x+width-150, y+20, data.Series)

	return nil
}

// RenderLineChart renders a line chart.
func (r *ChartRenderer) RenderLineChart(
	x, y, width, height float64,
	data ChartData,
) error {
	if len(data.Series) == 0 {
		return fmt.Errorf("no data series")
	}

	// Calculate chart area
	chartX := x + 40
	chartY := y + 20
	chartWidth := width - 60
	chartHeight := height - 60

	// Draw axes
	r.drawAxes(
		chartX,
		chartY,
		chartWidth,
		chartHeight,
	)

	// Find value range
	minValue := 0.0
	maxValue := 0.0
	for _, series := range data.Series {
		for _, val := range series.Values {
			if val > maxValue {
				maxValue = val
			}
			if val < minValue {
				minValue = val
			}
		}
	}

	valueRange := maxValue - minValue
	if valueRange == 0 {
		valueRange = 1
	}

	// Render each series
	for _, series := range data.Series {
		if len(series.Values) == 0 {
			continue
		}

		r.ctx.Page.SetStrokeColor(
			series.Color.R,
			series.Color.G,
			series.Color.B,
		)
		r.ctx.Page.SetLineWidth(2)

		// Calculate point spacing
		pointSpacing := chartWidth / float64(
			len(series.Values)-1,
		)

		// Draw line segments
		for i := range len(series.Values) - 1 {
			val1 := series.Values[i]
			val2 := series.Values[i+1]

			x1 := chartX + float64(i)*pointSpacing
			y1 := chartY + chartHeight - ((val1-minValue)/valueRange)*chartHeight
			x2 := chartX + float64(
				i+1,
			)*pointSpacing
			y2 := chartY + chartHeight - ((val2-minValue)/valueRange)*chartHeight

			path := NewPathBuilder()
			path.MoveTo(x1, y1)
			path.LineTo(x2, y2)
			r.ctx.Page.WriteContent(path.Stroke())

			// Draw marker at data point
			r.ctx.Page.DrawCircle(
				x1,
				y1,
				3,
				true,
				true,
			)
		}

		// Draw last marker
		lastIdx := len(series.Values) - 1
		lastX := chartX + float64(
			lastIdx,
		)*pointSpacing
		lastY := chartY + chartHeight - ((series.Values[lastIdx]-minValue)/valueRange)*chartHeight
		r.ctx.Page.DrawCircle(
			lastX,
			lastY,
			3,
			true,
			true,
		)
	}

	// Render legend
	r.renderLegend(x+width-150, y+20, data.Series)

	return nil
}

// RenderPieChart renders a pie chart.
func (r *ChartRenderer) RenderPieChart(
	x, y, width, height float64,
	data ChartData,
) error {
	if len(data.Series) == 0 ||
		len(data.Series[0].Values) == 0 {
		return fmt.Errorf("no data")
	}

	// Use first series for pie chart
	series := data.Series[0]
	values := series.Values

	// Calculate total
	total := 0.0
	for _, val := range values {
		total += val
	}

	if total == 0 {
		return fmt.Errorf("total is zero")
	}

	// Calculate pie center and radius
	cx := x + width/2
	cy := y + height/2
	radius := math.Min(width, height) / 2 * 0.8

	// Draw pie slices
	startAngle := -90.0 // Start at top

	colors := generateChartColors(len(values))

	for i, val := range values {
		sweepAngle := (val / total) * 360.0

		// Draw pie slice
		r.drawPieSlice(
			cx,
			cy,
			radius,
			startAngle,
			sweepAngle,
			colors[i],
		)

		startAngle += sweepAngle
	}

	// Render legend with values
	legendSeries := make(
		[]ChartSeries,
		len(values),
	)
	for i, val := range values {
		name := ""
		if i < len(data.Categories) {
			name = data.Categories[i]
		} else {
			name = fmt.Sprintf("Item %d", i+1)
		}
		legendSeries[i] = ChartSeries{
			Name:   name,
			Values: []float64{val},
			Color:  colors[i],
		}
	}
	r.renderLegend(
		x+width-150,
		y+20,
		legendSeries,
	)

	return nil
}

// drawPieSlice draws a single pie slice.
func (r *ChartRenderer) drawPieSlice(
	cx, cy, radius float64,
	startAngle, sweepAngle float64,
	color Color,
) {
	// Convert angles to radians
	startRad := startAngle * math.Pi / 180.0
	endRad := (startAngle + sweepAngle) * math.Pi / 180.0

	// Create path
	path := NewPathBuilder()

	// Move to center
	path.MoveTo(cx, cy)

	// Line to start of arc
	startX := cx + radius*math.Cos(startRad)
	startY := cy + radius*math.Sin(startRad)
	path.LineTo(startX, startY)

	// Draw arc (approximate with line segments)
	steps := int(
		math.Abs(sweepAngle) / 5,
	) // 5 degrees per step
	if steps < 2 {
		steps = 2
	}

	for i := 1; i <= steps; i++ {
		angle := startRad + (endRad-startRad)*float64(
			i,
		)/float64(
			steps,
		)
		px := cx + radius*math.Cos(angle)
		py := cy + radius*math.Sin(angle)
		path.LineTo(px, py)
	}

	// Close path back to center
	path.ClosePath()

	// Fill with color
	r.ctx.Page.SetFillColor(
		color.R,
		color.G,
		color.B,
	)
	r.ctx.Page.WriteContent(path.Fill())

	// Stroke outline
	r.ctx.Page.SetStrokeColor(
		1,
		1,
		1,
	) // White outline
	r.ctx.Page.SetLineWidth(1)
	r.ctx.Page.WriteContent(path.Stroke())
}

// drawAxes draws the X and Y axes.
func (r *ChartRenderer) drawAxes(
	x, y, width, height float64,
) {
	r.ctx.Page.SetStrokeColor(0, 0, 0)
	r.ctx.Page.SetLineWidth(1)

	// Y axis
	path := NewPathBuilder()
	path.MoveTo(x, y)
	path.LineTo(x, y+height)
	r.ctx.Page.WriteContent(path.Stroke())

	// X axis
	path = NewPathBuilder()
	path.MoveTo(x, y+height)
	path.LineTo(x+width, y+height)
	r.ctx.Page.WriteContent(path.Stroke())
}

// renderLegend renders the chart legend.
func (r *ChartRenderer) renderLegend(
	x, y float64,
	series []ChartSeries,
) {
	// Draw legend background
	legendWidth := 140.0
	legendHeight := float64(len(series))*20 + 10
	r.ctx.Page.SetFillColor(1, 1, 1)
	r.ctx.Page.SetStrokeColor(0, 0, 0)
	r.ctx.Page.DrawRectangle(
		x,
		y,
		legendWidth,
		legendHeight,
		true,
		true,
	)

	// Draw legend items
	currentY := y + 10
	for _, s := range series {
		// Draw color box
		r.ctx.Page.SetFillColor(
			s.Color.R,
			s.Color.G,
			s.Color.B,
		)
		r.ctx.Page.DrawRectangle(
			x+5,
			currentY,
			10,
			10,
			true,
			true,
		)

		// Draw series name
		r.ctx.Page.SetFillColor(0, 0, 0)
		r.ctx.Page.SetFont("Helvetica", 10)
		r.ctx.Page.DrawText(
			x+20,
			currentY,
			s.Name,
		)

		currentY += 20
	}
}

// generateChartColors generates a set of distinct colors for chart elements.
func generateChartColors(count int) []Color {
	// Predefined color palette
	baseColors := []Color{
		NewRGB(0.2, 0.4, 0.8), // Blue
		NewRGB(0.8, 0.3, 0.3), // Red
		NewRGB(0.3, 0.7, 0.3), // Green
		NewRGB(0.9, 0.6, 0.2), // Orange
		NewRGB(0.6, 0.3, 0.8), // Purple
		NewRGB(0.3, 0.8, 0.8), // Cyan
		NewRGB(0.9, 0.7, 0.3), // Yellow
		NewRGB(0.8, 0.4, 0.6), // Pink
	}

	colors := make([]Color, count)
	for i := range count {
		colors[i] = baseColors[i%len(baseColors)]
	}

	return colors
}

// RenderChartTitle renders the chart title.
func (r *ChartRenderer) RenderChartTitle(
	x, y, width float64,
	title string,
) {
	// Center the title
	r.ctx.Page.SetFont("Helvetica-Bold", 14)
	r.ctx.Page.SetFillColor(0, 0, 0)
	// Calculate approximate text width
	textWidth := float64(
		len(title),
	) * 7 // Rough estimate
	titleX := x + (width-textWidth)/2
	r.ctx.Page.DrawText(titleX, y, title)
}
