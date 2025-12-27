// Package elements provides Excel document element types for SpreadsheetML.
// This file contains enum type definitions and validation functions.
package elements

//revive:disable:file-length-limit many enum types and validators

// ChartType represents the type of chart in SpreadsheetML.
type ChartType string

const (
	// ChartTypeArea indicates an area chart.
	ChartTypeArea ChartType = "area"
	// ChartTypeAreaStacked indicates a stacked area chart.
	ChartTypeAreaStacked ChartType = "areaStacked"
	// ChartTypeAreaPercentStacked indicates a 100% stacked area chart.
	ChartTypeAreaPercentStacked ChartType = "areaPercentStacked"
	// ChartTypeArea3D indicates a 3D area chart.
	ChartTypeArea3D ChartType = "area3D"
	// ChartTypeBar indicates a bar chart.
	ChartTypeBar ChartType = "bar"
	// ChartTypeBarStacked indicates a stacked bar chart.
	ChartTypeBarStacked ChartType = "barStacked"
	// ChartTypeBarPercentStacked indicates a 100% stacked bar chart.
	ChartTypeBarPercentStacked ChartType = "barPercentStacked"
	// ChartTypeBar3D indicates a 3D bar chart.
	ChartTypeBar3D ChartType = "bar3D"
	// ChartTypeColumn indicates a column chart.
	ChartTypeColumn ChartType = "column"
	// ChartTypeColumnStacked indicates a stacked column chart.
	ChartTypeColumnStacked ChartType = "columnStacked"
	// ChartTypeColumnPercentStacked indicates a 100% stacked column chart.
	ChartTypeColumnPercentStacked ChartType = "columnPercentStacked"
	// ChartTypeColumn3D indicates a 3D column chart.
	ChartTypeColumn3D ChartType = "column3D"
	// ChartTypeLine indicates a line chart.
	ChartTypeLine ChartType = "line"
	// ChartTypeLineStacked indicates a stacked line chart.
	ChartTypeLineStacked ChartType = "lineStacked"
	// ChartTypeLinePercentStacked indicates a 100% stacked line chart.
	ChartTypeLinePercentStacked ChartType = "linePercentStacked"
	// ChartTypeLine3D indicates a 3D line chart.
	ChartTypeLine3D ChartType = "line3D"
	// ChartTypeLineMarker indicates a line chart with markers.
	ChartTypeLineMarker ChartType = "lineMarker"
	// ChartTypePie indicates a pie chart.
	ChartTypePie ChartType = "pie"
	// ChartTypePie3D indicates a 3D pie chart.
	ChartTypePie3D ChartType = "pie3D"
	// ChartTypePieExploded indicates an exploded pie chart.
	ChartTypePieExploded ChartType = "pieExploded"
	// ChartTypeDoughnut indicates a doughnut chart.
	ChartTypeDoughnut ChartType = "doughnut"
	// ChartTypeDoughnutExploded indicates an exploded doughnut chart.
	ChartTypeDoughnutExploded ChartType = "doughnutExploded"
	// ChartTypeScatter indicates a scatter chart.
	ChartTypeScatter ChartType = "scatter"
	// ChartTypeScatterLine indicates a scatter chart with lines.
	ChartTypeScatterLine ChartType = "scatterLine"
	// ChartTypeScatterLineSmooth indicates a scatter chart with smooth lines.
	ChartTypeScatterLineSmooth ChartType = "scatterLineSmooth"
	// ChartTypeScatterMarker indicates a scatter chart with markers only.
	ChartTypeScatterMarker ChartType = "scatterMarker"
	// ChartTypeBubble indicates a bubble chart.
	ChartTypeBubble ChartType = "bubble"
	// ChartTypeBubble3D indicates a 3D bubble chart.
	ChartTypeBubble3D ChartType = "bubble3D"
	// ChartTypeRadar indicates a radar chart.
	ChartTypeRadar ChartType = "radar"
	// ChartTypeRadarFilled indicates a filled radar chart.
	ChartTypeRadarFilled ChartType = "radarFilled"
	// ChartTypeStock indicates a stock chart.
	ChartTypeStock ChartType = "stock"
	// ChartTypeSurface indicates a surface chart.
	ChartTypeSurface ChartType = "surface"
	// ChartTypeSurface3D indicates a 3D surface chart.
	ChartTypeSurface3D ChartType = "surface3D"
	// ChartTypeSurfaceWireframe indicates a wireframe surface chart.
	ChartTypeSurfaceWireframe ChartType = "surfaceWireframe"
	// ChartTypeSurfaceContour indicates a contour surface chart.
	ChartTypeSurfaceContour ChartType = "surfaceContour"
	// ChartTypeCombo indicates a combo chart.
	ChartTypeCombo ChartType = "combo"
)

// ChartGrouping represents how chart data series are grouped.
type ChartGrouping string

const (
	// ChartGroupingStandard indicates standard grouping.
	ChartGroupingStandard ChartGrouping = "standard"
	// ChartGroupingClustered indicates clustered grouping.
	ChartGroupingClustered ChartGrouping = "clustered"
	// ChartGroupingStacked indicates stacked grouping.
	ChartGroupingStacked ChartGrouping = "stacked"
	// ChartGroupingPercentStacked indicates percent stacked grouping.
	ChartGroupingPercentStacked ChartGrouping = "percentStacked"
)

// ChartLegendPosition represents the position of a chart legend.
type ChartLegendPosition string

const (
	// ChartLegendPositionBottom indicates legend at the bottom.
	ChartLegendPositionBottom ChartLegendPosition = "b"
	// ChartLegendPositionTop indicates legend at the top.
	ChartLegendPositionTop ChartLegendPosition = "t"
	// ChartLegendPositionLeft indicates legend at the left.
	ChartLegendPositionLeft ChartLegendPosition = "l"
	// ChartLegendPositionRight indicates legend at the right.
	ChartLegendPositionRight ChartLegendPosition = "r"
	// ChartLegendPositionTopRight indicates legend at top-right.
	ChartLegendPositionTopRight ChartLegendPosition = "tr"
)

// Note: SheetState is defined in sheets.go

// PaperSize represents standard paper sizes.
type PaperSize int

const (
	// PaperSizeLetter indicates Letter (8.5 x 11 inches).
	PaperSizeLetter PaperSize = 1
	// PaperSizeLetterSmall indicates Letter Small.
	PaperSizeLetterSmall PaperSize = 2
	// PaperSizeTabloid indicates Tabloid (11 x 17 inches).
	PaperSizeTabloid PaperSize = 3
	// PaperSizeLedger indicates Ledger (17 x 11 inches).
	PaperSizeLedger PaperSize = 4
	// PaperSizeLegal indicates Legal (8.5 x 14 inches).
	PaperSizeLegal PaperSize = 5
	// PaperSizeStatement indicates Statement.
	PaperSizeStatement PaperSize = 6
	// PaperSizeExecutive indicates Executive.
	PaperSizeExecutive PaperSize = 7
	// PaperSizeA3 indicates A3 (297 x 420 mm).
	PaperSizeA3 PaperSize = 8
	// PaperSizeA4 indicates A4 (210 x 297 mm).
	PaperSizeA4 PaperSize = 9
	// PaperSizeA4Small indicates A4 Small.
	PaperSizeA4Small PaperSize = 10
	// PaperSizeA5 indicates A5 (148 x 210 mm).
	PaperSizeA5 PaperSize = 11
	// PaperSizeB4 indicates B4 (250 x 353 mm).
	PaperSizeB4 PaperSize = 12
	// PaperSizeB5 indicates B5 (182 x 257 mm).
	PaperSizeB5 PaperSize = 13
	// PaperSizeFolio indicates Folio.
	PaperSizeFolio PaperSize = 14
	// PaperSizeQuarto indicates Quarto.
	PaperSizeQuarto PaperSize = 15
	// PaperSize10x14 indicates 10 x 14 inches.
	PaperSize10x14 PaperSize = 16
	// PaperSize11x17 indicates 11 x 17 inches.
	PaperSize11x17 PaperSize = 17
)

// PageOrientation represents page orientation for printing.
type PageOrientation string

const (
	// PageOrientationDefault indicates default orientation.
	PageOrientationDefault PageOrientation = "default"
	// PageOrientationPortrait indicates portrait orientation.
	PageOrientationPortrait PageOrientation = "portrait"
	// PageOrientationLandscape indicates landscape orientation.
	PageOrientationLandscape PageOrientation = "landscape"
)

// Note: PrintErrors, CellComments, PageOrder are defined in page_setup.go

// Note: FontScheme is defined in font.go

// Note: UnderlineStyle is defined in font.go as the font's underline element
// handles this.

// VerticalTextAlignment represents vertical text alignment in a run.
type VerticalTextAlignment string

const (
	// VerticalTextAlignmentBaseline indicates baseline alignment.
	VerticalTextAlignmentBaseline VerticalTextAlignment = "baseline"
	// VerticalTextAlignmentSuperscript indicates superscript.
	VerticalTextAlignmentSuperscript VerticalTextAlignment = "superscript"
	// VerticalTextAlignmentSubscript indicates subscript.
	VerticalTextAlignmentSubscript VerticalTextAlignment = "subscript"
)

// Note: GradientType is defined in gradient_fill.go

// Note: IconSetType is defined in icon_set.go

// Note: CfvoType is defined in cfvo.go

// Note: PaneState and ActivePane are defined in pane.go

// Note: DataBarDirection is defined in data_bar.go

// enum validation functions

// ValidCellTypes returns all valid CellType values.
func ValidCellTypes() []CellType {
	return []CellType{
		CellTypeBoolean,
		CellTypeDate,
		CellTypeError,
		CellTypeInlineString,
		CellTypeNumber,
		CellTypeSharedString,
		CellTypeFormulaString,
	}
}

// IsValidCellType returns true if the given CellType is valid.
func IsValidCellType(ct CellType) bool {
	for _, valid := range ValidCellTypes() {
		if ct == valid {
			return true
		}
	}

	return false
}

// ValidFormulaTypes returns all valid FormulaType values.
func ValidFormulaTypes() []FormulaType {
	return []FormulaType{
		FormulaTypeNormal,
		FormulaTypeArray,
		FormulaTypeDataTable,
		FormulaTypeShared,
	}
}

// IsValidFormulaType returns true if the given FormulaType is valid.
func IsValidFormulaType(ft FormulaType) bool {
	for _, valid := range ValidFormulaTypes() {
		if ft == valid {
			return true
		}
	}

	return false
}

// ValidBorderStyles returns all valid BorderStyle values.
func ValidBorderStyles() []BorderStyle {
	return []BorderStyle{
		BorderStyleNone,
		BorderStyleThin,
		BorderStyleMedium,
		BorderStyleDashed,
		BorderStyleDotted,
		BorderStyleThick,
		BorderStyleDouble,
		BorderStyleHair,
		BorderStyleMediumDashed,
		BorderStyleDashDot,
		BorderStyleMediumDashDot,
		BorderStyleDashDotDot,
		BorderStyleMediumDashDotDot,
		BorderStyleSlantDashDot,
	}
}

// IsValidBorderStyle returns true if the given BorderStyle is valid.
func IsValidBorderStyle(bs BorderStyle) bool {
	for _, valid := range ValidBorderStyles() {
		if bs == valid {
			return true
		}
	}

	return false
}

// ValidPatternTypes returns all valid PatternType values.
func ValidPatternTypes() []PatternType {
	return []PatternType{
		PatternTypeNone,
		PatternTypeSolid,
		PatternTypeMediumGray,
		PatternTypeDarkGray,
		PatternTypeLightGray,
		PatternTypeDarkHorizontal,
		PatternTypeDarkVertical,
		PatternTypeDarkDown,
		PatternTypeDarkUp,
		PatternTypeDarkGrid,
		PatternTypeDarkTrellis,
		PatternTypeLightHorizontal,
		PatternTypeLightVertical,
		PatternTypeLightDown,
		PatternTypeLightUp,
		PatternTypeLightGrid,
		PatternTypeLightTrellis,
		PatternTypeGray125,
		PatternTypeGray0625,
	}
}

// IsValidPatternType returns true if the given PatternType is valid.
func IsValidPatternType(pt PatternType) bool {
	for _, valid := range ValidPatternTypes() {
		if pt == valid {
			return true
		}
	}

	return false
}

// ValidHorizontalAlignments returns all valid HorizontalAlignment values.
func ValidHorizontalAlignments() []HorizontalAlignment {
	return []HorizontalAlignment{
		HorizontalAlignmentGeneral,
		HorizontalAlignmentLeft,
		HorizontalAlignmentCenter,
		HorizontalAlignmentRight,
		HorizontalAlignmentFill,
		HorizontalAlignmentJustify,
		HorizontalAlignmentCenterContinuous,
		HorizontalAlignmentDistributed,
	}
}

// IsValidHorizontalAlignment returns true if the given HorizontalAlignment is
// valid.
func IsValidHorizontalAlignment(
	ha HorizontalAlignment,
) bool {
	for _, valid := range ValidHorizontalAlignments() {
		if ha == valid {
			return true
		}
	}

	return false
}

// ValidVerticalAlignments returns all valid VerticalAlignment values.
func ValidVerticalAlignments() []VerticalAlignment {
	return []VerticalAlignment{
		VerticalAlignmentTop,
		VerticalAlignmentCenter,
		VerticalAlignmentBottom,
		VerticalAlignmentJustify,
		VerticalAlignmentDistributed,
	}
}

// IsValidVerticalAlignment returns true if the given VerticalAlignment is
// valid.
func IsValidVerticalAlignment(
	va VerticalAlignment,
) bool {
	for _, valid := range ValidVerticalAlignments() {
		if va == valid {
			return true
		}
	}

	return false
}

// ValidValidationTypes returns all valid ValidationType values.
func ValidValidationTypes() []ValidationType {
	return []ValidationType{
		ValidationTypeNone,
		ValidationTypeWhole,
		ValidationTypeDecimal,
		ValidationTypeList,
		ValidationTypeDate,
		ValidationTypeTime,
		ValidationTypeTextLength,
		ValidationTypeCustom,
	}
}

// IsValidValidationType returns true if the given ValidationType is valid.
func IsValidValidationType(
	vt ValidationType,
) bool {
	for _, valid := range ValidValidationTypes() {
		if vt == valid {
			return true
		}
	}

	return false
}

// ValidCfRuleTypes returns all valid CfRuleType values.
func ValidCfRuleTypes() []CfRuleType {
	return []CfRuleType{
		CfRuleTypeExpression,
		CfRuleTypeCellIs,
		CfRuleTypeColorScale,
		CfRuleTypeDataBar,
		CfRuleTypeIconSet,
		CfRuleTypeTop10,
		CfRuleTypeUniqueValues,
		CfRuleTypeDuplicateValues,
		CfRuleTypeContainsText,
		CfRuleTypeNotContainsText,
		CfRuleTypeBeginsWith,
		CfRuleTypeEndsWith,
		CfRuleTypeContainsBlanks,
		CfRuleTypeNotContainsBlanks,
		CfRuleTypeContainsErrors,
		CfRuleTypeNotContainsErrors,
		CfRuleTypeTimePeriod,
		CfRuleTypeAboveAverage,
	}
}

// IsValidCfRuleType returns true if the given CfRuleType is valid.
func IsValidCfRuleType(crt CfRuleType) bool {
	for _, valid := range ValidCfRuleTypes() {
		if crt == valid {
			return true
		}
	}

	return false
}

// ValidChartTypes returns all valid ChartType values.
func ValidChartTypes() []ChartType {
	return []ChartType{
		ChartTypeArea,
		ChartTypeAreaStacked,
		ChartTypeAreaPercentStacked,
		ChartTypeArea3D,
		ChartTypeBar,
		ChartTypeBarStacked,
		ChartTypeBarPercentStacked,
		ChartTypeBar3D,
		ChartTypeColumn,
		ChartTypeColumnStacked,
		ChartTypeColumnPercentStacked,
		ChartTypeColumn3D,
		ChartTypeLine,
		ChartTypeLineStacked,
		ChartTypeLinePercentStacked,
		ChartTypeLine3D,
		ChartTypeLineMarker,
		ChartTypePie,
		ChartTypePie3D,
		ChartTypePieExploded,
		ChartTypeDoughnut,
		ChartTypeDoughnutExploded,
		ChartTypeScatter,
		ChartTypeScatterLine,
		ChartTypeScatterLineSmooth,
		ChartTypeScatterMarker,
		ChartTypeBubble,
		ChartTypeBubble3D,
		ChartTypeRadar,
		ChartTypeRadarFilled,
		ChartTypeStock,
		ChartTypeSurface,
		ChartTypeSurface3D,
		ChartTypeSurfaceWireframe,
		ChartTypeSurfaceContour,
		ChartTypeCombo,
	}
}

// IsValidChartType returns true if the given ChartType is valid.
func IsValidChartType(ct ChartType) bool {
	for _, valid := range ValidChartTypes() {
		if ct == valid {
			return true
		}
	}

	return false
}

// ValidIconSetTypes returns all valid IconSetType values.
func ValidIconSetTypes() []IconSetType {
	return []IconSetType{
		IconSet3Arrows,
		IconSet3ArrowsGray,
		IconSet3Flags,
		IconSet3TrafficLights1,
		IconSet3TrafficLights2,
		IconSet3Signs,
		IconSet3Symbols,
		IconSet3Symbols2,
		IconSet4Arrows,
		IconSet4ArrowsGray,
		IconSet4RedToBlack,
		IconSet4Rating,
		IconSet4TrafficLights,
		IconSet5Arrows,
		IconSet5ArrowsGray,
		IconSet5Rating,
		IconSet5Quarters,
	}
}

// IsValidIconSetType returns true if the given IconSetType is valid.
func IsValidIconSetType(ist IconSetType) bool {
	for _, valid := range ValidIconSetTypes() {
		if ist == valid {
			return true
		}
	}

	return false
}
