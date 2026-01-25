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
	BorderStyle        = BorderValues
	ShadingPattern     = ShadingPatternValues
	UnderlineValue     = UnderlineValues
	HighlightColor     = HighlightColorValues
	VerticalAlignValue = VerticalAlignmentValues
	EmphasisMarkValue  = EmphasisMarkValues
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
	UnderlineDouble UnderlineValue = UnderlineValuesDouble
	UnderlineThick  UnderlineValue = UnderlineValuesThick
	UnderlineDotted UnderlineValue = UnderlineValuesDotted
	UnderlineDash   UnderlineValue = UnderlineValuesDash
	UnderlineWave   UnderlineValue = UnderlineValuesWave
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
	BorderNone   BorderStyle = BorderValuesNone
	BorderSingle BorderStyle = BorderValuesSingle
	BorderDouble BorderStyle = BorderValuesDouble
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
// TODO: Re-enable these when VerticalAlignmentRunValues enum is regenerated
// const (
// 	VerticalAlignBaseline    VerticalAlignValue = VerticalAlignmentRunValuesBaseline
// 	VerticalAlignSuperscript VerticalAlignValue = VerticalAlignmentRunValuesSuperscript
// 	VerticalAlignSubscript   VerticalAlignValue = VerticalAlignmentRunValuesSubscript
// )

// EmphasisMarkValue constants
const (
	EmphasisNone EmphasisMarkValue = EmphasisMarkValuesNone
)

// HeaderFooter constants (note: need to add type alias)
type HeaderFooter = HeaderFooterValues

const (
	HeaderFooterDefault HeaderFooter = HeaderFooterValuesDefault
	HeaderFooterFirst   HeaderFooter = HeaderFooterValuesFirst
	HeaderFooterEven    HeaderFooter = HeaderFooterValuesEven
)
