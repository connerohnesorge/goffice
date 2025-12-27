package elements

import "testing"

// TestCellTypeValidation tests CellType enum validation.
func TestCellTypeValidation(t *testing.T) {
	tests := []struct {
		value CellType
		valid bool
	}{
		{CellTypeBoolean, true},
		{CellTypeDate, true},
		{CellTypeError, true},
		{CellTypeInlineString, true},
		{CellTypeNumber, true},
		{CellTypeSharedString, true},
		{CellTypeFormulaString, true},
		{CellType("invalid"), false},
		{CellType(""), false},
	}

	for _, tc := range tests {
		t.Run(
			string(tc.value),
			func(t *testing.T) {
				result := IsValidCellType(
					tc.value,
				)
				if result != tc.valid {
					t.Errorf(
						"IsValidCellType(%q) = %v, want %v",
						tc.value,
						result,
						tc.valid,
					)
				}
			},
		)
	}
}

// TestFormulaTypeValidation tests FormulaType enum validation.
func TestFormulaTypeValidation(t *testing.T) {
	tests := []struct {
		value FormulaType
		valid bool
	}{
		{FormulaTypeNormal, true},
		{FormulaTypeArray, true},
		{FormulaTypeDataTable, true},
		{FormulaTypeShared, true},
		{FormulaType("invalid"), false},
	}

	for _, tc := range tests {
		t.Run(
			string(tc.value),
			func(t *testing.T) {
				result := IsValidFormulaType(
					tc.value,
				)
				if result != tc.valid {
					t.Errorf(
						"IsValidFormulaType(%q) = %v, want %v",
						tc.value,
						result,
						tc.valid,
					)
				}
			},
		)
	}
}

// TestBorderStyleValidation tests BorderStyle enum validation.
func TestBorderStyleValidation(t *testing.T) {
	validStyles := ValidBorderStyles()
	if len(validStyles) != 14 {
		t.Errorf(
			"ValidBorderStyles() returned %d values, want 14",
			len(validStyles),
		)
	}

	for _, style := range validStyles {
		if !IsValidBorderStyle(style) {
			t.Errorf(
				"IsValidBorderStyle(%q) = false, want true",
				style,
			)
		}
	}

	if IsValidBorderStyle(
		BorderStyle("invalid"),
	) {
		t.Error(
			"IsValidBorderStyle('invalid') = true, want false",
		)
	}
}

// TestPatternTypeValidation tests PatternType enum validation.
func TestPatternTypeValidation(t *testing.T) {
	validTypes := ValidPatternTypes()
	if len(validTypes) < 18 {
		t.Errorf(
			"ValidPatternTypes() returned %d values, want at least 18",
			len(validTypes),
		)
	}

	for _, pt := range validTypes {
		if !IsValidPatternType(pt) {
			t.Errorf(
				"IsValidPatternType(%q) = false, want true",
				pt,
			)
		}
	}

	if IsValidPatternType(
		PatternType("invalid"),
	) {
		t.Error(
			"IsValidPatternType('invalid') = true, want false",
		)
	}
}

// TestHorizontalAlignmentValidation tests HorizontalAlignment enum validation.
func TestHorizontalAlignmentValidation(
	t *testing.T,
) {
	validAligns := ValidHorizontalAlignments()
	if len(validAligns) != 8 {
		t.Errorf(
			"ValidHorizontalAlignments() returned %d values, want 8",
			len(validAligns),
		)
	}

	for _, align := range validAligns {
		if !IsValidHorizontalAlignment(align) {
			t.Errorf(
				"IsValidHorizontalAlignment(%q) = false, want true",
				align,
			)
		}
	}

	if IsValidHorizontalAlignment(
		HorizontalAlignment("invalid"),
	) {
		t.Error(
			"IsValidHorizontalAlignment('invalid') = true, want false",
		)
	}
}

// TestVerticalAlignmentValidation tests VerticalAlignment enum validation.
func TestVerticalAlignmentValidation(
	t *testing.T,
) {
	validAligns := ValidVerticalAlignments()
	if len(validAligns) != 5 {
		t.Errorf(
			"ValidVerticalAlignments() returned %d values, want 5",
			len(validAligns),
		)
	}

	for _, align := range validAligns {
		if !IsValidVerticalAlignment(align) {
			t.Errorf(
				"IsValidVerticalAlignment(%q) = false, want true",
				align,
			)
		}
	}

	if IsValidVerticalAlignment(
		VerticalAlignment("invalid"),
	) {
		t.Error(
			"IsValidVerticalAlignment('invalid') = true, want false",
		)
	}
}

// TestValidationTypeValidation tests ValidationType enum validation.
func TestValidationTypeValidation(t *testing.T) {
	validTypes := ValidValidationTypes()
	if len(validTypes) != 8 {
		t.Errorf(
			"ValidValidationTypes() returned %d values, want 8",
			len(validTypes),
		)
	}

	for _, vt := range validTypes {
		if !IsValidValidationType(vt) {
			t.Errorf(
				"IsValidValidationType(%q) = false, want true",
				vt,
			)
		}
	}

	if IsValidValidationType(
		ValidationType("invalid"),
	) {
		t.Error(
			"IsValidValidationType('invalid') = true, want false",
		)
	}
}

// TestCfRuleTypeValidation tests CfRuleType enum validation.
func TestCfRuleTypeValidation(t *testing.T) {
	validTypes := ValidCfRuleTypes()
	if len(validTypes) < 18 {
		t.Errorf(
			"ValidCfRuleTypes() returned %d values, want at least 18",
			len(validTypes),
		)
	}

	for _, crt := range validTypes {
		if !IsValidCfRuleType(crt) {
			t.Errorf(
				"IsValidCfRuleType(%q) = false, want true",
				crt,
			)
		}
	}

	if IsValidCfRuleType(CfRuleType("invalid")) {
		t.Error(
			"IsValidCfRuleType('invalid') = true, want false",
		)
	}
}

// TestChartTypeValidation tests ChartType enum validation.
func TestChartTypeValidation(t *testing.T) {
	validTypes := ValidChartTypes()
	if len(validTypes) < 35 {
		t.Errorf(
			"ValidChartTypes() returned %d values, want at least 35",
			len(validTypes),
		)
	}

	for _, ct := range validTypes {
		if !IsValidChartType(ct) {
			t.Errorf(
				"IsValidChartType(%q) = false, want true",
				ct,
			)
		}
	}

	if IsValidChartType(ChartType("invalid")) {
		t.Error(
			"IsValidChartType('invalid') = true, want false",
		)
	}
}

// Note: IconSetType validation is done via the icon_set.go definitions

// TestChartTypeConstants tests all chart type constants are defined.
func TestChartTypeConstants(t *testing.T) {
	// Test a sampling of chart types to ensure they're properly defined
	tests := []struct {
		chartType ChartType
		expected  string
	}{
		{ChartTypeArea, "area"},
		{ChartTypeBar, "bar"},
		{ChartTypeColumn, "column"},
		{ChartTypeLine, "line"},
		{ChartTypePie, "pie"},
		{ChartTypeDoughnut, "doughnut"},
		{ChartTypeScatter, "scatter"},
		{ChartTypeBubble, "bubble"},
		{ChartTypeRadar, "radar"},
		{ChartTypeStock, "stock"},
		{ChartTypeSurface, "surface"},
		{ChartTypeCombo, "combo"},
	}

	for _, tc := range tests {
		if string(tc.chartType) != tc.expected {
			t.Errorf(
				"ChartType constant = %q, want %q",
				tc.chartType,
				tc.expected,
			)
		}
	}
}

// TestSheetStateConstants tests sheet state constants.
func TestSheetStateConstants(t *testing.T) {
	tests := []struct {
		state    SheetState
		expected string
	}{
		{SheetStateVisible, "visible"},
		{SheetStateHidden, "hidden"},
		{SheetStateVeryHidden, "veryHidden"},
	}

	for _, tc := range tests {
		if string(tc.state) != tc.expected {
			t.Errorf(
				"SheetState constant = %q, want %q",
				tc.state,
				tc.expected,
			)
		}
	}
}

// TestPaneStateConstants tests pane state constants.
func TestPaneStateConstants(t *testing.T) {
	tests := []struct {
		state    PaneState
		expected string
	}{
		{PaneStateSplit, "split"},
		{PaneStateFrozen, "frozen"},
		{PaneStateFrozenSplit, "frozenSplit"},
	}

	for _, tc := range tests {
		if string(tc.state) != tc.expected {
			t.Errorf(
				"PaneState constant = %q, want %q",
				tc.state,
				tc.expected,
			)
		}
	}
}

// Note: ActivePane tests are in pane_test.go since that type is defined in pane.go

// Note: UnderlineStyle tests are in font_test.go since that type is defined in font.go

// TestPageOrientationConstants tests page orientation constants.
func TestPageOrientationConstants(t *testing.T) {
	tests := []struct {
		orient   PageOrientation
		expected string
	}{
		{PageOrientationDefault, "default"},
		{PageOrientationPortrait, "portrait"},
		{PageOrientationLandscape, "landscape"},
	}

	for _, tc := range tests {
		if string(tc.orient) != tc.expected {
			t.Errorf(
				"PageOrientation constant = %q, want %q",
				tc.orient,
				tc.expected,
			)
		}
	}
}

// Note: GradientType tests would go in gradient_fill_test.go

// TestCfvoTypeConstants tests CFVO type constants.
func TestCfvoTypeConstants(t *testing.T) {
	tests := []struct {
		cfvoType CfvoType
		expected string
	}{
		{CfvoTypeNum, "num"},
		{CfvoTypePercent, "percent"},
		{CfvoTypeMax, "max"},
		{CfvoTypeMin, "min"},
		{CfvoTypeFormula, "formula"},
		{CfvoTypePercentile, "percentile"},
	}

	for _, tc := range tests {
		if string(tc.cfvoType) != tc.expected {
			t.Errorf(
				"CfvoType constant = %q, want %q",
				tc.cfvoType,
				tc.expected,
			)
		}
	}
}

// TestChartGroupingConstants tests chart grouping constants.
func TestChartGroupingConstants(t *testing.T) {
	tests := []struct {
		grouping ChartGrouping
		expected string
	}{
		{ChartGroupingStandard, "standard"},
		{ChartGroupingClustered, "clustered"},
		{ChartGroupingStacked, "stacked"},
		{
			ChartGroupingPercentStacked,
			"percentStacked",
		},
	}

	for _, tc := range tests {
		if string(tc.grouping) != tc.expected {
			t.Errorf(
				"ChartGrouping constant = %q, want %q",
				tc.grouping,
				tc.expected,
			)
		}
	}
}

// TestChartLegendPositionConstants tests chart legend position constants.
func TestChartLegendPositionConstants(
	t *testing.T,
) {
	tests := []struct {
		pos      ChartLegendPosition
		expected string
	}{
		{ChartLegendPositionBottom, "b"},
		{ChartLegendPositionTop, "t"},
		{ChartLegendPositionLeft, "l"},
		{ChartLegendPositionRight, "r"},
		{ChartLegendPositionTopRight, "tr"},
	}

	for _, tc := range tests {
		if string(tc.pos) != tc.expected {
			t.Errorf(
				"ChartLegendPosition constant = %q, want %q",
				tc.pos,
				tc.expected,
			)
		}
	}
}

// TestDataBarDirectionConstants tests data bar direction constants.
func TestDataBarDirectionConstants(t *testing.T) {
	tests := []struct {
		dir      DataBarDirection
		expected string
	}{
		{DataBarDirectionContext, "context"},
		{DataBarDirectionLTR, "leftToRight"},
		{DataBarDirectionRTL, "rightToLeft"},
	}

	for _, tc := range tests {
		if string(tc.dir) != tc.expected {
			t.Errorf(
				"DataBarDirection constant = %q, want %q",
				tc.dir,
				tc.expected,
			)
		}
	}
}

// TestTimePeriodTypeConstants tests time period type constants.
func TestTimePeriodTypeConstants(t *testing.T) {
	tests := []struct {
		period   TimePeriodType
		expected string
	}{
		{TimePeriodToday, "today"},
		{TimePeriodYesterday, "yesterday"},
		{TimePeriodTomorrow, "tomorrow"},
		{TimePeriodLast7Days, "last7Days"},
		{TimePeriodThisMonth, "thisMonth"},
		{TimePeriodLastMonth, "lastMonth"},
		{TimePeriodNextMonth, "nextMonth"},
		{TimePeriodThisWeek, "thisWeek"},
		{TimePeriodLastWeek, "lastWeek"},
		{TimePeriodNextWeek, "nextWeek"},
	}

	for _, tc := range tests {
		if string(tc.period) != tc.expected {
			t.Errorf(
				"TimePeriodType constant = %q, want %q",
				tc.period,
				tc.expected,
			)
		}
	}
}
