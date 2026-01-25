//revive:disable:max-public-structs enum type definitions must be public

// and effects.
//
// This file contains chart enumeration types used across chart files.
package drawingml

// Common attribute name constants used across chart files.
const (
	// attrVal is the "val" attribute name used throughout charts.
	attrVal = "val"
	// attrZero is the "0" value used for boolean false attributes.
	attrZero = "0"
	// attrOne is the "1" value used for boolean true attributes.
	attrOne = "1"
)

// BarDirectionValue represents bar chart direction.
type BarDirectionValue string

// Bar direction values.
const (
	BarDirectionBar BarDirectionValue = "bar" // Horizontal bars
	BarDirectionCol BarDirectionValue = "col" // Vertical columns
)

// BarGroupingValue represents bar chart grouping style.
type BarGroupingValue string

// Bar grouping values.
const (
	BarGroupingClustered      BarGroupingValue = "clustered"
	BarGroupingStacked        BarGroupingValue = "stacked"
	BarGroupingPercentStacked BarGroupingValue = "percentStacked"
	BarGroupingStandard       BarGroupingValue = "standard"
)

// GroupingValue represents chart grouping for line/area charts.
type GroupingValue string

// Grouping values.
const (
	GroupingStandard       GroupingValue = "standard"
	GroupingStacked        GroupingValue = "stacked"
	GroupingPercentStacked GroupingValue = "percentStacked"
)

// ScatterStyleValue represents scatter chart style.
type ScatterStyleValue string

// Scatter style values.
const (
	ScatterStyleNone         ScatterStyleValue = "none"
	ScatterStyleLine         ScatterStyleValue = "line"
	ScatterStyleLineMarker   ScatterStyleValue = "lineMarker"
	ScatterStyleMarker       ScatterStyleValue = "marker"
	ScatterStyleSmooth       ScatterStyleValue = "smooth"
	ScatterStyleSmoothMarker ScatterStyleValue = "smoothMarker"
)

// MarkerStyleValue represents data marker styles.
type MarkerStyleValue string

// Marker style values.
const (
	MarkerStyleCircle   MarkerStyleValue = "circle"
	MarkerStyleDash     MarkerStyleValue = "dash"
	MarkerStyleDiamond  MarkerStyleValue = "diamond"
	MarkerStyleDot      MarkerStyleValue = "dot"
	MarkerStyleNone     MarkerStyleValue = "none"
	MarkerStylePlus     MarkerStyleValue = "plus"
	MarkerStyleSquare   MarkerStyleValue = "square"
	MarkerStyleStar     MarkerStyleValue = "star"
	MarkerStyleTriangle MarkerStyleValue = "triangle"
	MarkerStyleX        MarkerStyleValue = "x"
)

// LegendPositionValue represents legend position.
type LegendPositionValue string

// Legend position values.
const (
	LegendPositionBottom   LegendPositionValue = "b"
	LegendPositionTop      LegendPositionValue = "t"
	LegendPositionLeft     LegendPositionValue = "l"
	LegendPositionRight    LegendPositionValue = "r"
	LegendPositionTopRight LegendPositionValue = "tr"
)

// AxisPositionValue represents axis position.
type AxisPositionValue string

// Axis position values.
const (
	AxisPositionBottom AxisPositionValue = "b"
	AxisPositionTop    AxisPositionValue = "t"
	AxisPositionLeft   AxisPositionValue = "l"
	AxisPositionRight  AxisPositionValue = "r"
)

// TickMarkValue represents tick mark style.
type TickMarkValue string

// Tick mark values.
const (
	TickMarkNone  TickMarkValue = "none"
	TickMarkIn    TickMarkValue = "in"
	TickMarkOut   TickMarkValue = "out"
	TickMarkCross TickMarkValue = "cross"
)

// TickLabelPositionValue represents tick label position.
type TickLabelPositionValue string

// Tick label position values.
const (
	TickLabelPositionNone   TickLabelPositionValue = "none"
	TickLabelPositionHigh   TickLabelPositionValue = "high"
	TickLabelPositionLow    TickLabelPositionValue = "low"
	TickLabelPositionNextTo TickLabelPositionValue = "nextTo"
)

// DisplayBlanksAsValue represents how blank cells are displayed.
type DisplayBlanksAsValue string

// Display blanks as values.
const (
	DisplayBlanksAsSpan DisplayBlanksAsValue = "span"
	DisplayBlanksAsGap  DisplayBlanksAsValue = "gap"
	DisplayBlanksAsZero DisplayBlanksAsValue = "zero"
)

// CrossesValue represents axis crossing behavior.
type CrossesValue string

// Crosses values.
const (
	CrossesAutoZero CrossesValue = "autoZero"
	CrossesMax      CrossesValue = "max"
	CrossesMin      CrossesValue = "min"
)

// OrientationValue represents axis orientation.
type OrientationValue string

// Orientation values.
const (
	OrientationMaxMin OrientationValue = "maxMin"
	OrientationMinMax OrientationValue = "minMax"
)

// RadarStyleValue represents radar chart style.
type RadarStyleValue string

// Radar style values.
const (
	RadarStyleStandard RadarStyleValue = "standard"
	RadarStyleMarker   RadarStyleValue = "marker"
	RadarStyleFilled   RadarStyleValue = "filled"
)

// OfPieTypeValue represents pie-of-pie or bar-of-pie chart type.
type OfPieTypeValue string

// Of pie type values.
const (
	OfPieTypePie OfPieTypeValue = "pie"
	OfPieTypeBar OfPieTypeValue = "bar"
)

// DataLabelPositionValue represents data label position.
type DataLabelPositionValue string

// Data label position values.
const (
	DataLabelPositionBestFit DataLabelPositionValue = "bestFit"
	DataLabelPositionBottom  DataLabelPositionValue = "b"
	DataLabelPositionCenter  DataLabelPositionValue = "ctr"
	DataLabelPositionInBase  DataLabelPositionValue = "inBase"
	DataLabelPositionInEnd   DataLabelPositionValue = "inEnd"
	DataLabelPositionLeft    DataLabelPositionValue = "l"
	DataLabelPositionOutEnd  DataLabelPositionValue = "outEnd"
	DataLabelPositionRight   DataLabelPositionValue = "r"
	DataLabelPositionTop     DataLabelPositionValue = "t"
)
