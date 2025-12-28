// Package elements provides manually-defined enumerations and type aliases for WordprocessingML.
//
//nolint:revive // line-length-limit, max-public-structs - manual enums file
package elements

import (
	"encoding/xml"
)

// SpaceProcessingModeValues for xml:space attribute
type SpaceProcessingModeValues string

const (
	SpaceProcessingModeDefault  SpaceProcessingModeValues = "default"
	SpaceProcessingModePreserve SpaceProcessingModeValues = "preserve"
)

func (e SpaceProcessingModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SpaceProcessingModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SpaceProcessingModeValues(attr.Value)

	return nil
}

// Type aliases for cleaner API
type (
	BreakType          = BreakValues
	LineSpacingRule    = LineSpacingRuleValues
	JustificationValue = JustificationValues
	TextAlignmentValue = TextAlignmentTypeValues
	TabAlignment       = TextTabAlignmentValues
	TabLeader          = TabStopLeaderCharValues
	BorderStyle        = BorderStyleValues
	ShadingPattern     = ShadingPatternValues
	UnderlineValue     = UnderlineValues
	HighlightColor     = HighlightColorValues
	VerticalAlignValue = VerticalAlignmentRunValues
	EmphasisMarkValue  = EmphasisMarkValues
	PageOrientation    = PageOrientationValues
)

// Break type constants
const (
	BreakLine   BreakType = BreakValuesTextwrapping
	BreakPage   BreakType = BreakValuesPage
	BreakColumn BreakType = BreakValuesColumn
)

// LineSpacingRule constants
const (
	LineSpacingRuleAuto    LineSpacingRule = LineSpacingRuleValuesAuto
	LineSpacingRuleExact   LineSpacingRule = LineSpacingRuleValuesExact
	LineSpacingRuleAtLeast LineSpacingRule = LineSpacingRuleValuesAtleast
)

// Justification constants
const (
	JustificationLeft           JustificationValue = JustificationValuesLeft
	JustificationStart          JustificationValue = JustificationValuesStart
	JustificationCenter         JustificationValue = JustificationValuesCenter
	JustificationRight          JustificationValue = JustificationValuesRight
	JustificationEnd            JustificationValue = JustificationValuesEnd
	JustificationBoth           JustificationValue = JustificationValuesBoth
	JustificationMediumKashida  JustificationValue = JustificationValuesMediumkashida
	JustificationDistribute     JustificationValue = JustificationValuesDistribute
	JustificationNumTab         JustificationValue = JustificationValuesNumtab
	JustificationHighKashida    JustificationValue = JustificationValuesHighkashida
	JustificationLowKashida     JustificationValue = JustificationValuesLowkashida
	JustificationThaiDistribute JustificationValue = JustificationValuesThaidistribute
)

// TabLeader constants
const (
	TabLeaderNone       TabLeader = TabStopLeaderCharValuesNone
	TabLeaderDot        TabLeader = TabStopLeaderCharValuesDot
	TabLeaderHyphen     TabLeader = TabStopLeaderCharValuesHyphen
	TabLeaderUnderscore TabLeader = TabStopLeaderCharValuesUnderscore
	TabLeaderHeavy      TabLeader = TabStopLeaderCharValuesHeavy
	TabLeaderMiddleDot  TabLeader = TabStopLeaderCharValuesMiddledot
)

// UnderlineValue constants
const (
	UnderlineNone   UnderlineValue = ""
	UnderlineSingle UnderlineValue = UnderlineValuesSingle
)

// TextAlignmentValue constants
const (
	TextAlignmentAuto TextAlignmentValue = TextAlignmentTypeValuesJust
)

// LineSpacingRule constants
const (
	LineSpacingAuto LineSpacingRule = LineSpacingRuleValuesAuto
)

// TabAlignment constants
const (
	TabAlignLeft TabAlignment = TextTabAlignmentValuesL
)

// BorderStyle constants
const (
	BorderNone   BorderStyle = BorderStyleValuesNone
	BorderSingle BorderStyle = "single" // Custom value for single border
	BorderDouble BorderStyle = BorderStyleValuesDouble
)

// ShadingPattern constants
const (
	ShadingClear ShadingPattern = ShadingPatternValuesClear
)

// HighlightColor constants
const (
	HighlightNone   HighlightColor = ""
	HighlightYellow HighlightColor = HighlightColorValuesYellow
)

// VerticalAlignValue constants
const (
	VerticalAlignBaseline VerticalAlignValue = VerticalAlignmentRunValuesBaseline
)

// EmphasisMarkValue constants
const (
	EmphasisNone EmphasisMarkValue = EmphasisMarkValuesNone
)

// PageOrientation constants
const (
	PageOrientationPortrait  PageOrientation = PageOrientationValuesPortrait
	PageOrientationLandscape PageOrientation = PageOrientationValuesLandscape
)

// HeaderFooter constants (note: need to add type alias)
type HeaderFooter = HeaderFooterValues

const (
	HeaderFooterDefault HeaderFooter = HeaderFooterValuesDefault
)
