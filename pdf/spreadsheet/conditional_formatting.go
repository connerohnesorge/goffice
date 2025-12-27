// conditional_formatting.go handles the rendering of conditional formatting visualization.
// This includes cell highlighting, data bars, color scales, and icon sets.

package spreadsheet

import (
	"strconv"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// ConditionalFormattingResult contains the visual result of conditional formatting evaluation.
type ConditionalFormattingResult struct {
	Type    string // "highlight", "dataBar", "colorScale", "iconSet"
	BgColor *core.RGB
	FgColor *core.RGB
	DataBar *DataBarVisual
	Icon    *Icon
}

// DataBarVisual contains visual properties for a data bar.
type DataBarVisual struct {
	Value   float64
	MinVal  float64
	MaxVal  float64
	Color   core.RGB
	Percent float64
}

// Icon contains properties for an icon set icon.
type Icon struct {
	Type  string // "arrow", "trafficLight", "rating", etc.
	Index int
	Color core.RGB
}

// applyConditionalFormatting evaluates and applies conditional formatting to a cell.
func (r *SpreadsheetRenderer) applyConditionalFormatting(
	worksheet *elements.Worksheet,
	cellRef *CellRef,
	cellValue string,
	style *CellStyle,
) *ConditionalFormattingResult {
	// Get conditional formatting rules for the worksheet
	cfList := worksheet.ConditionalFormatting()
	if cfList == nil {
		return nil
	}

	// Iterate through conditional formatting rules
	for _, cf := range cfList {
		// Check if this rule applies to the cell
		sqref := cf.Sqref()
		if !r.cellInRange(cellRef, sqref) {
			continue
		}

		// Evaluate each rule
		for rule := range cf.Rules() {
			result := r.evaluateConditionalFormattingRule(
				rule,
				cellRef,
				cellValue,
				worksheet,
			)
			if result != nil {
				return result
			}
		}
	}

	return nil
}

// cellInRange checks if a cell is in a range reference.
func (r *SpreadsheetRenderer) cellInRange(
	cellRef *CellRef,
	sqref string,
) bool {
	// Parse the sequence of references (space-separated ranges)
	ranges := splitSpaces(sqref)

	for _, rangeStr := range ranges {
		cellRange := parseRange(rangeStr)
		if cellRange != nil {
			if cellRef.Row >= cellRange.StartRow &&
				cellRef.Row <= cellRange.EndRow &&
				cellRef.Col >= cellRange.StartCol &&
				cellRef.Col <= cellRange.EndCol {
				return true
			}
		}
	}

	return false
}

// splitSpaces splits a string by spaces.
func splitSpaces(s string) []string {
	var result []string
	current := ""

	for _, ch := range s {
		if ch == ' ' {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(ch)
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}

// evaluateConditionalFormattingRule evaluates a single conditional formatting rule.
func (r *SpreadsheetRenderer) evaluateConditionalFormattingRule(
	rule *elements.CfRule,
	cellRef *CellRef,
	cellValue string,
	worksheet *elements.Worksheet,
) *ConditionalFormattingResult {
	cfType := rule.Type()

	switch cfType {
	case elements.CfRuleTypeCellIs:
		return r.evaluateCellIsRule(
			rule,
			cellValue,
		)

	case elements.CfRuleTypeExpression:
		return r.evaluateExpressionRule(
			rule,
			cellRef,
			cellValue,
			worksheet,
		)

	case elements.CfRuleTypeDataBar:
		return r.evaluateDataBarRule(
			rule,
			cellRef,
			cellValue,
			worksheet,
		)

	case elements.CfRuleTypeColorScale:
		return r.evaluateColorScaleRule(
			rule,
			cellRef,
			cellValue,
			worksheet,
		)

	case elements.CfRuleTypeIconSet:
		return r.evaluateIconSetRule(
			rule,
			cellRef,
			cellValue,
			worksheet,
		)

	default:
		return nil
	}
}

// evaluateCellIsRule evaluates a "cellIs" type rule (e.g., greater than, less than).
func (r *SpreadsheetRenderer) evaluateCellIsRule(
	rule *elements.CfRule,
	cellValue string,
) *ConditionalFormattingResult {
	operator := rule.Operator()

	// Get the comparison value from the formula
	var compareValue string
	for formula := range rule.Formulas() {
		compareValue = formula.Formula()
		break
	}

	if compareValue == "" {
		return nil
	}

	// Parse cell value as number
	numValue, err := strconv.ParseFloat(
		cellValue,
		64,
	)
	if err != nil {
		return nil
	}

	// Parse compare value as number
	numCompare, err := strconv.ParseFloat(
		compareValue,
		64,
	)
	if err != nil {
		return nil
	}

	// Evaluate operator
	match := false
	switch operator {
	case elements.CfRuleOperatorGreaterThan:
		match = numValue > numCompare
	case elements.CfRuleOperatorLessThan:
		match = numValue < numCompare
	case elements.CfRuleOperatorGreaterThanOrEqual:
		match = numValue >= numCompare
	case elements.CfRuleOperatorLessThanOrEqual:
		match = numValue <= numCompare
	case elements.CfRuleOperatorEqual:
		match = numValue == numCompare
	case elements.CfRuleOperatorNotEqual:
		match = numValue != numCompare
	case elements.CfRuleOperatorBetween:
		// Need second formula for between
		idx := 0
		var compareValue2 string
		for formula := range rule.Formulas() {
			if idx == 1 {
				compareValue2 = formula.Formula()
				break
			}
			idx++
		}
		if compareValue2 != "" {
			numCompare2, err := strconv.ParseFloat(
				compareValue2,
				64,
			)
			if err == nil {
				match = numValue >= numCompare &&
					numValue <= numCompare2
			}
		}
	}

	if !match {
		return nil
	}

	// Get the DXF (differential formatting) for this rule
	dxfId := rule.DxfId()
	if dxfId == 0 {
		return nil
	}

	// Apply the highlighting
	// For simplicity, use a default highlight color
	return &ConditionalFormattingResult{
		Type: "highlight",
		BgColor: &core.RGB{
			R: 1.0,
			G: 1.0,
			B: 0.8,
		},
	}
}

// evaluateExpressionRule evaluates a formula-based rule.
func (r *SpreadsheetRenderer) evaluateExpressionRule(
	rule *elements.CfRule,
	cellRef *CellRef,
	cellValue string,
	worksheet *elements.Worksheet,
) *ConditionalFormattingResult {
	// Formula evaluation is complex - skip for now
	return nil
}

// evaluateDataBarRule evaluates a data bar rule.
func (r *SpreadsheetRenderer) evaluateDataBarRule(
	rule *elements.CfRule,
	cellRef *CellRef,
	cellValue string,
	worksheet *elements.Worksheet,
) *ConditionalFormattingResult {
	// Get data bar properties
	dataBar := rule.DataBar()
	if dataBar == nil {
		return nil
	}

	// Parse cell value as number
	numValue, err := strconv.ParseFloat(
		cellValue,
		64,
	)
	if err != nil {
		return nil
	}

	// Get min and max values from CFVOs
	minVal := 0.0
	maxVal := 100.0

	idx := 0
	for cfvo := range dataBar.Cfvos() {
		if idx == 0 {
			minVal = r.evaluateCfvo(
				cfvo,
				worksheet,
				cellRef,
			)
		} else if idx == 1 {
			maxVal = r.evaluateCfvo(
				cfvo,
				worksheet,
				cellRef,
			)
		}
		idx++
	}

	// Calculate percentage
	percent := 0.0
	if maxVal > minVal {
		percent = (numValue - minVal) / (maxVal - minVal)
		if percent < 0 {
			percent = 0
		}
		if percent > 1 {
			percent = 1
		}
	}

	// Get color
	color := core.RGB{
		R: 0.4,
		G: 0.6,
		B: 0.9,
	} // Default blue
	if barColor := dataBar.Color(); barColor != nil {
		color = r.parseColor(barColor)
	}

	return &ConditionalFormattingResult{
		Type: "dataBar",
		DataBar: &DataBarVisual{
			Value:   numValue,
			MinVal:  minVal,
			MaxVal:  maxVal,
			Color:   color,
			Percent: percent,
		},
	}
}

// evaluateColorScaleRule evaluates a color scale rule.
func (r *SpreadsheetRenderer) evaluateColorScaleRule(
	rule *elements.CfRule,
	cellRef *CellRef,
	cellValue string,
	worksheet *elements.Worksheet,
) *ConditionalFormattingResult {
	// Get color scale properties
	colorScale := rule.ColorScale()
	if colorScale == nil {
		return nil
	}

	// Parse cell value as number
	numValue, err := strconv.ParseFloat(
		cellValue,
		64,
	)
	if err != nil {
		return nil
	}

	// Get CFVOs and colors
	cfvos := []*elements.Cfvo{}
	colors := []*elements.Color{}

	for cfvo := range colorScale.Cfvos() {
		cfvos = append(cfvos, cfvo)
	}

	for color := range colorScale.Colors() {
		colors = append(colors, color)
	}

	if len(cfvos) < 2 || len(colors) < 2 {
		return nil
	}

	// Evaluate min and max
	minVal := r.evaluateCfvo(
		cfvos[0],
		worksheet,
		cellRef,
	)
	maxVal := r.evaluateCfvo(
		cfvos[len(cfvos)-1],
		worksheet,
		cellRef,
	)

	// Calculate position (0-1)
	position := 0.0
	if maxVal > minVal {
		position = (numValue - minVal) / (maxVal - minVal)
		if position < 0 {
			position = 0
		}
		if position > 1 {
			position = 1
		}
	}

	// Interpolate color
	// For 2-color scale: interpolate between colors[0] and colors[1]
	// For 3-color scale: interpolate based on position
	var color core.RGB
	if len(colors) == 2 {
		color = interpolateColor(
			r.parseColor(colors[0]),
			r.parseColor(colors[1]),
			position,
		)
	} else if len(colors) == 3 {
		if position < 0.5 {
			color = interpolateColor(
				r.parseColor(colors[0]),
				r.parseColor(colors[1]),
				position*2,
			)
		} else {
			color = interpolateColor(
				r.parseColor(colors[1]),
				r.parseColor(colors[2]),
				(position-0.5)*2,
			)
		}
	} else {
		color = r.parseColor(colors[0])
	}

	return &ConditionalFormattingResult{
		Type:    "highlight",
		BgColor: &color,
	}
}

// evaluateIconSetRule evaluates an icon set rule.
func (r *SpreadsheetRenderer) evaluateIconSetRule(
	rule *elements.CfRule,
	cellRef *CellRef,
	cellValue string,
	worksheet *elements.Worksheet,
) *ConditionalFormattingResult {
	// Icon sets are complex - simplified implementation
	iconSet := rule.IconSet()
	if iconSet == nil {
		return nil
	}

	// Parse cell value as number
	numValue, err := strconv.ParseFloat(
		cellValue,
		64,
	)
	if err != nil {
		return nil
	}

	// Get icon type (e.g., "3Arrows", "3TrafficLights", etc.)
	iconType := iconSet.IconSetType()

	// Determine which icon to show based on value
	// This is a simplified version - real implementation would evaluate CFVOs
	iconIndex := 0
	if numValue > 66 {
		iconIndex = 2 // High
	} else if numValue > 33 {
		iconIndex = 1 // Medium
	} else {
		iconIndex = 0 // Low
	}

	return &ConditionalFormattingResult{
		Type: "iconSet",
		Icon: &Icon{
			Type:  string(iconType),
			Index: iconIndex,
			Color: core.RGB{R: 0, G: 0, B: 0},
		},
	}
}

// evaluateCfvo evaluates a conditional formatting value object.
func (r *SpreadsheetRenderer) evaluateCfvo(
	cfvo *elements.Cfvo,
	worksheet *elements.Worksheet,
	cellRef *CellRef,
) float64 {
	cfvoType := cfvo.Type()

	switch cfvoType {
	case elements.CfvoTypeMin:
		// Find minimum value in range
		return 0.0 // Simplified

	case elements.CfvoTypeMax:
		// Find maximum value in range
		return 100.0 // Simplified

	case elements.CfvoTypeNum:
		// Direct number value
		val := cfvo.Val()
		num, _ := strconv.ParseFloat(val, 64)
		return num

	case elements.CfvoTypePercent:
		// Percentage of range
		val := cfvo.Val()
		percent, _ := strconv.ParseFloat(val, 64)
		return percent

	case elements.CfvoTypePercentile:
		// Percentile of range
		return 50.0 // Simplified

	case elements.CfvoTypeFormula:
		// Formula evaluation - skip for now
		return 0.0

	default:
		return 0.0
	}
}

// interpolateColor interpolates between two colors.
func interpolateColor(
	c1, c2 core.RGB,
	t float64,
) core.RGB {
	return core.RGB{
		R: c1.R + (c2.R-c1.R)*t,
		G: c1.G + (c2.G-c1.G)*t,
		B: c1.B + (c2.B-c1.B)*t,
	}
}

// renderConditionalFormatting renders conditional formatting visual elements.
func (r *SpreadsheetRenderer) renderConditionalFormatting(
	page *core.Page,
	x, y, width, height float64,
	result *ConditionalFormattingResult,
) {
	if result == nil {
		return
	}

	switch result.Type {
	case "highlight":
		// Render background color
		if result.BgColor != nil {
			drawFillRect(
				page,
				x,
				y,
				width,
				height,
				*result.BgColor,
				1.0,
			)
		}

	case "dataBar":
		// Render data bar
		if result.DataBar != nil {
			r.renderDataBar(
				page,
				x,
				y,
				width,
				height,
				result.DataBar,
			)
		}

	case "iconSet":
		// Render icon
		if result.Icon != nil {
			r.renderIcon(
				page,
				x,
				y,
				height,
				result.Icon,
			)
		}
	}
}

// renderDataBar renders a data bar within a cell.
func (r *SpreadsheetRenderer) renderDataBar(
	page *core.Page,
	x, y, width, height float64,
	dataBar *DataBarVisual,
) {
	// Use the helper function from drawing_helpers.go
	drawDataBar(
		page,
		x,
		y,
		width,
		height,
		dataBar.Percent,
		dataBar.Color,
	)
}

// renderIcon renders an icon set icon within a cell.
func (r *SpreadsheetRenderer) renderIcon(
	page *core.Page,
	x, y, height float64,
	icon *Icon,
) {
	// Determine color based on icon index
	var color core.RGB
	switch icon.Index {
	case 0:
		color = core.RGB{
			R: 1.0,
			G: 0,
			B: 0,
		} // Red
	case 1:
		color = core.RGB{
			R: 1.0,
			G: 1.0,
			B: 0,
		} // Yellow
	case 2:
		color = core.RGB{
			R: 0,
			G: 1.0,
			B: 0,
		} // Green
	default:
		color = core.RGB{
			R: 0.5,
			G: 0.5,
			B: 0.5,
		} // Gray
	}

	// Draw circle
	radius := height / 4
	centerX := x + radius + 2
	centerY := y + height/2

	drawCircle(
		page,
		centerX,
		centerY,
		radius,
		color,
	)
}
