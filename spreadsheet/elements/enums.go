// nolint
//
//revive:disable:file-length-limit,max-public-structs,line-length-limit,comments-density // Generated code
package elements

import (
	"encoding/xml"
)

// AnimateColorSpaceValues
type AnimateColorSpaceValues string

const (
	AnimateColorSpaceValuesRgb AnimateColorSpaceValues = "rgb"
	AnimateColorSpaceValuesHsl AnimateColorSpaceValues = "hsl"
)

func (e AnimateColorSpaceValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AnimateColorSpaceValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AnimateColorSpaceValues(attr.Value)
	return nil
}

// WebSourceValues
type WebSourceValues string

const (
	WebSourceValuesSheet      WebSourceValues = "sheet"
	WebSourceValuesPrintarea  WebSourceValues = "printArea"
	WebSourceValuesAutofilter WebSourceValues = "autoFilter"
	WebSourceValuesRange      WebSourceValues = "range"
	WebSourceValuesChart      WebSourceValues = "chart"
	WebSourceValuesPivottable WebSourceValues = "pivotTable"
	WebSourceValuesQuery      WebSourceValues = "query"
	WebSourceValuesLabel      WebSourceValues = "label"
)

func (e WebSourceValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *WebSourceValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = WebSourceValues(attr.Value)
	return nil
}

// MdxKPIPropertyValues
type MdxKPIPropertyValues string

const (
	MdxKPIPropertyValuesV MdxKPIPropertyValues = "v"
	MdxKPIPropertyValuesG MdxKPIPropertyValues = "g"
	MdxKPIPropertyValuesS MdxKPIPropertyValues = "s"
	MdxKPIPropertyValuesT MdxKPIPropertyValues = "t"
	MdxKPIPropertyValuesW MdxKPIPropertyValues = "w"
	MdxKPIPropertyValuesM MdxKPIPropertyValues = "m"
)

func (e MdxKPIPropertyValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MdxKPIPropertyValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MdxKPIPropertyValues(attr.Value)
	return nil
}

// PageBorderDisplayValues
type PageBorderDisplayValues string

const (
	PageBorderDisplayValuesAllpages     PageBorderDisplayValues = "allPages"
	PageBorderDisplayValuesFirstpage    PageBorderDisplayValues = "firstPage"
	PageBorderDisplayValuesNotfirstpage PageBorderDisplayValues = "notFirstPage"
)

func (e PageBorderDisplayValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PageBorderDisplayValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PageBorderDisplayValues(attr.Value)
	return nil
}

// BlipCompressionValues
type BlipCompressionValues string

const (
	BlipCompressionValuesEmail   BlipCompressionValues = "email"
	BlipCompressionValuesScreen  BlipCompressionValues = "screen"
	BlipCompressionValuesPrint   BlipCompressionValues = "print"
	BlipCompressionValuesHqprint BlipCompressionValues = "hqprint"
	BlipCompressionValuesNone    BlipCompressionValues = "none"
)

func (e BlipCompressionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BlipCompressionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BlipCompressionValues(attr.Value)
	return nil
}

// GrowShrinkValues
type GrowShrinkValues string

const (
	GrowShrinkValuesInsertdelete   GrowShrinkValues = "insertDelete"
	GrowShrinkValuesInsertclear    GrowShrinkValues = "insertClear"
	GrowShrinkValuesOverwriteclear GrowShrinkValues = "overwriteClear"
)

func (e GrowShrinkValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *GrowShrinkValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = GrowShrinkValues(attr.Value)
	return nil
}

// PatternValues
type PatternValues string

const (
	PatternValuesNone            PatternValues = "none"
	PatternValuesSolid           PatternValues = "solid"
	PatternValuesMediumgray      PatternValues = "mediumGray"
	PatternValuesDarkgray        PatternValues = "darkGray"
	PatternValuesLightgray       PatternValues = "lightGray"
	PatternValuesDarkhorizontal  PatternValues = "darkHorizontal"
	PatternValuesDarkvertical    PatternValues = "darkVertical"
	PatternValuesDarkdown        PatternValues = "darkDown"
	PatternValuesDarkup          PatternValues = "darkUp"
	PatternValuesDarkgrid        PatternValues = "darkGrid"
	PatternValuesDarktrellis     PatternValues = "darkTrellis"
	PatternValuesLighthorizontal PatternValues = "lightHorizontal"
	PatternValuesLightvertical   PatternValues = "lightVertical"
	PatternValuesLightdown       PatternValues = "lightDown"
	PatternValuesLightup         PatternValues = "lightUp"
	PatternValuesLightgrid       PatternValues = "lightGrid"
	PatternValuesLighttrellis    PatternValues = "lightTrellis"
	PatternValuesGray125         PatternValues = "gray125"
	PatternValuesGray0625        PatternValues = "gray0625"
)

func (e PatternValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PatternValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PatternValues(attr.Value)
	return nil
}

// ChapterSeparatorValues
type ChapterSeparatorValues string

const (
	ChapterSeparatorValuesHyphen ChapterSeparatorValues = "hyphen"
	ChapterSeparatorValuesPeriod ChapterSeparatorValues = "period"
	ChapterSeparatorValuesColon  ChapterSeparatorValues = "colon"
	ChapterSeparatorValuesEmdash ChapterSeparatorValues = "emDash"
	ChapterSeparatorValuesEndash ChapterSeparatorValues = "enDash"
)

func (e ChapterSeparatorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ChapterSeparatorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ChapterSeparatorValues(attr.Value)
	return nil
}

// TextFontAlignmentValues
type TextFontAlignmentValues string

const (
	TextFontAlignmentValuesAuto TextFontAlignmentValues = "auto"
	TextFontAlignmentValuesT    TextFontAlignmentValues = "t"
	TextFontAlignmentValuesCtr  TextFontAlignmentValues = "ctr"
	TextFontAlignmentValuesBase TextFontAlignmentValues = "base"
	TextFontAlignmentValuesB    TextFontAlignmentValues = "b"
)

func (e TextFontAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextFontAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextFontAlignmentValues(attr.Value)
	return nil
}

// LevelSuffixValues
type LevelSuffixValues string

const (
	LevelSuffixValuesTab     LevelSuffixValues = "tab"
	LevelSuffixValuesSpace   LevelSuffixValues = "space"
	LevelSuffixValuesNothing LevelSuffixValues = "nothing"
)

func (e LevelSuffixValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LevelSuffixValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LevelSuffixValues(attr.Value)
	return nil
}

// LineEndValues
type LineEndValues string

const (
	LineEndValuesNone     LineEndValues = "none"
	LineEndValuesTriangle LineEndValues = "triangle"
	LineEndValuesStealth  LineEndValues = "stealth"
	LineEndValuesDiamond  LineEndValues = "diamond"
	LineEndValuesOval     LineEndValues = "oval"
	LineEndValuesArrow    LineEndValues = "arrow"
)

func (e LineEndValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LineEndValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LineEndValues(attr.Value)
	return nil
}

// TransitionCornerDirectionValues
type TransitionCornerDirectionValues string

const (
	TransitionCornerDirectionValuesLu TransitionCornerDirectionValues = "lu"
	TransitionCornerDirectionValuesRu TransitionCornerDirectionValues = "ru"
	TransitionCornerDirectionValuesLd TransitionCornerDirectionValues = "ld"
	TransitionCornerDirectionValuesRd TransitionCornerDirectionValues = "rd"
)

func (e TransitionCornerDirectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TransitionCornerDirectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TransitionCornerDirectionValues(
		attr.Value,
	)
	return nil
}

// ConditionalFormatValueObjectValues
type ConditionalFormatValueObjectValues string

const (
	ConditionalFormatValueObjectValuesNum        ConditionalFormatValueObjectValues = "num"
	ConditionalFormatValueObjectValuesPercent    ConditionalFormatValueObjectValues = "percent"
	ConditionalFormatValueObjectValuesMax        ConditionalFormatValueObjectValues = "max"
	ConditionalFormatValueObjectValuesMin        ConditionalFormatValueObjectValues = "min"
	ConditionalFormatValueObjectValuesFormula    ConditionalFormatValueObjectValues = "formula"
	ConditionalFormatValueObjectValuesPercentile ConditionalFormatValueObjectValues = "percentile"
)

func (e ConditionalFormatValueObjectValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ConditionalFormatValueObjectValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ConditionalFormatValueObjectValues(
		attr.Value,
	)
	return nil
}

// BreakTextRestartLocationValues
type BreakTextRestartLocationValues string

const (
	BreakTextRestartLocationValuesNone  BreakTextRestartLocationValues = "none"
	BreakTextRestartLocationValuesLeft  BreakTextRestartLocationValues = "left"
	BreakTextRestartLocationValuesRight BreakTextRestartLocationValues = "right"
	BreakTextRestartLocationValuesAll   BreakTextRestartLocationValues = "all"
)

func (e BreakTextRestartLocationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BreakTextRestartLocationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BreakTextRestartLocationValues(
		attr.Value,
	)
	return nil
}

// MergedCellValues
type MergedCellValues string

const (
	MergedCellValuesContinue MergedCellValues = "continue"
	MergedCellValuesRestart  MergedCellValues = "restart"
)

func (e MergedCellValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MergedCellValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MergedCellValues(attr.Value)
	return nil
}

// DataBarDirectionValues
type DataBarDirectionValues string

const (
	DataBarDirectionValuesContext     DataBarDirectionValues = "context"
	DataBarDirectionValuesLefttoright DataBarDirectionValues = "leftToRight"
	DataBarDirectionValuesRighttoleft DataBarDirectionValues = "rightToLeft"
)

func (e DataBarDirectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DataBarDirectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DataBarDirectionValues(attr.Value)
	return nil
}

// ConditionalFormattingValueObjectTypeValues
type ConditionalFormattingValueObjectTypeValues string

const (
	ConditionalFormattingValueObjectTypeValuesNum        ConditionalFormattingValueObjectTypeValues = "num"
	ConditionalFormattingValueObjectTypeValuesPercent    ConditionalFormattingValueObjectTypeValues = "percent"
	ConditionalFormattingValueObjectTypeValuesMax        ConditionalFormattingValueObjectTypeValues = "max"
	ConditionalFormattingValueObjectTypeValuesMin        ConditionalFormattingValueObjectTypeValues = "min"
	ConditionalFormattingValueObjectTypeValuesFormula    ConditionalFormattingValueObjectTypeValues = "formula"
	ConditionalFormattingValueObjectTypeValuesPercentile ConditionalFormattingValueObjectTypeValues = "percentile"
	ConditionalFormattingValueObjectTypeValuesAutomin    ConditionalFormattingValueObjectTypeValues = "autoMin"
	ConditionalFormattingValueObjectTypeValuesAutomax    ConditionalFormattingValueObjectTypeValues = "autoMax"
)

func (e ConditionalFormattingValueObjectTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ConditionalFormattingValueObjectTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ConditionalFormattingValueObjectTypeValues(
		attr.Value,
	)
	return nil
}

// SelectionTypeValues
type SelectionTypeValues string

const (
	SelectionTypeValuesSingle   SelectionTypeValues = "single"
	SelectionTypeValuesMulti    SelectionTypeValues = "multi"
	SelectionTypeValuesExtended SelectionTypeValues = "extended"
)

func (e SelectionTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SelectionTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SelectionTypeValues(attr.Value)
	return nil
}

// PresetLineDashValues
type PresetLineDashValues string

const (
	PresetLineDashValuesSolid         PresetLineDashValues = "solid"
	PresetLineDashValuesDot           PresetLineDashValues = "dot"
	PresetLineDashValuesDash          PresetLineDashValues = "dash"
	PresetLineDashValuesLgdash        PresetLineDashValues = "lgDash"
	PresetLineDashValuesDashdot       PresetLineDashValues = "dashDot"
	PresetLineDashValuesLgdashdot     PresetLineDashValues = "lgDashDot"
	PresetLineDashValuesLgdashdotdot  PresetLineDashValues = "lgDashDotDot"
	PresetLineDashValuesSysdash       PresetLineDashValues = "sysDash"
	PresetLineDashValuesSysdot        PresetLineDashValues = "sysDot"
	PresetLineDashValuesSysdashdot    PresetLineDashValues = "sysDashDot"
	PresetLineDashValuesSysdashdotdot PresetLineDashValues = "sysDashDotDot"
)

func (e PresetLineDashValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PresetLineDashValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PresetLineDashValues(attr.Value)
	return nil
}

// CompoundLineValues
type CompoundLineValues string

const (
	CompoundLineValuesSng       CompoundLineValues = "sng"
	CompoundLineValuesDbl       CompoundLineValues = "dbl"
	CompoundLineValuesThickthin CompoundLineValues = "thickThin"
	CompoundLineValuesThinthick CompoundLineValues = "thinThick"
	CompoundLineValuesTri       CompoundLineValues = "tri"
)

func (e CompoundLineValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CompoundLineValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CompoundLineValues(attr.Value)
	return nil
}

// CommandValues
type CommandValues string

const (
	CommandValuesEvt  CommandValues = "evt"
	CommandValuesCall CommandValues = "call"
	CommandValuesVerb CommandValues = "verb"
)

func (e CommandValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CommandValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CommandValues(attr.Value)
	return nil
}

// ViewValues
type ViewValues string

const (
	ViewValuesNone        ViewValues = "none"
	ViewValuesPrint       ViewValues = "print"
	ViewValuesOutline     ViewValues = "outline"
	ViewValuesMasterpages ViewValues = "masterPages"
	ViewValuesNormal      ViewValues = "normal"
	ViewValuesWeb         ViewValues = "web"
)

func (e ViewValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ViewValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ViewValues(attr.Value)
	return nil
}

// SourceValues
type SourceValues string

const (
	SourceValuesWorksheet     SourceValues = "worksheet"
	SourceValuesExternal      SourceValues = "external"
	SourceValuesConsolidation SourceValues = "consolidation"
	SourceValuesScenario      SourceValues = "scenario"
)

func (e SourceValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SourceValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SourceValues(attr.Value)
	return nil
}

// DropStyleValues
type DropStyleValues string

const (
	DropStyleValuesCombo     DropStyleValues = "combo"
	DropStyleValuesComboedit DropStyleValues = "comboedit"
	DropStyleValuesSimple    DropStyleValues = "simple"
)

func (e DropStyleValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DropStyleValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DropStyleValues(attr.Value)
	return nil
}

// SlicerCacheCrossFilterValues
type SlicerCacheCrossFilterValues string

const (
	SlicerCacheCrossFilterValuesNone                   SlicerCacheCrossFilterValues = "none"
	SlicerCacheCrossFilterValuesShowitemswithdataattop SlicerCacheCrossFilterValues = "showItemsWithDataAtTop"
	SlicerCacheCrossFilterValuesShowitemswithnodata    SlicerCacheCrossFilterValues = "showItemsWithNoData"
)

func (e SlicerCacheCrossFilterValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SlicerCacheCrossFilterValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SlicerCacheCrossFilterValues(attr.Value)
	return nil
}

// BevelPresetValues
type BevelPresetValues string

const (
	BevelPresetValuesRelaxedinset BevelPresetValues = "relaxedInset"
	BevelPresetValuesCircle       BevelPresetValues = "circle"
	BevelPresetValuesSlope        BevelPresetValues = "slope"
	BevelPresetValuesCross        BevelPresetValues = "cross"
	BevelPresetValuesAngle        BevelPresetValues = "angle"
	BevelPresetValuesSoftround    BevelPresetValues = "softRound"
	BevelPresetValuesConvex       BevelPresetValues = "convex"
	BevelPresetValuesCoolslant    BevelPresetValues = "coolSlant"
	BevelPresetValuesDivot        BevelPresetValues = "divot"
	BevelPresetValuesRiblet       BevelPresetValues = "riblet"
	BevelPresetValuesHardedge     BevelPresetValues = "hardEdge"
	BevelPresetValuesArtdeco      BevelPresetValues = "artDeco"
)

func (e BevelPresetValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BevelPresetValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BevelPresetValues(attr.Value)
	return nil
}

// TimeNodePresetClassValues
type TimeNodePresetClassValues string

const (
	TimeNodePresetClassValuesEntr      TimeNodePresetClassValues = "entr"
	TimeNodePresetClassValuesExit      TimeNodePresetClassValues = "exit"
	TimeNodePresetClassValuesEmph      TimeNodePresetClassValues = "emph"
	TimeNodePresetClassValuesPath      TimeNodePresetClassValues = "path"
	TimeNodePresetClassValuesVerb      TimeNodePresetClassValues = "verb"
	TimeNodePresetClassValuesMediacall TimeNodePresetClassValues = "mediacall"
)

func (e TimeNodePresetClassValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TimeNodePresetClassValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TimeNodePresetClassValues(attr.Value)
	return nil
}

// TimeNodeRestartValues
type TimeNodeRestartValues string

const (
	TimeNodeRestartValuesAlways        TimeNodeRestartValues = "always"
	TimeNodeRestartValuesWhennotactive TimeNodeRestartValues = "whenNotActive"
	TimeNodeRestartValuesNever         TimeNodeRestartValues = "never"
)

func (e TimeNodeRestartValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TimeNodeRestartValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TimeNodeRestartValues(attr.Value)
	return nil
}

// RowColumnActionValues
type RowColumnActionValues string

const (
	RowColumnActionValuesInsertrow RowColumnActionValues = "insertRow"
	RowColumnActionValuesDeleterow RowColumnActionValues = "deleteRow"
	RowColumnActionValuesInsertcol RowColumnActionValues = "insertCol"
	RowColumnActionValuesDeletecol RowColumnActionValues = "deleteCol"
)

func (e RowColumnActionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *RowColumnActionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = RowColumnActionValues(attr.Value)
	return nil
}

// DataValidationImeModeValues
type DataValidationImeModeValues string

const (
	DataValidationImeModeValuesNocontrol    DataValidationImeModeValues = "noControl"
	DataValidationImeModeValuesOff          DataValidationImeModeValues = "off"
	DataValidationImeModeValuesOn           DataValidationImeModeValues = "on"
	DataValidationImeModeValuesDisabled     DataValidationImeModeValues = "disabled"
	DataValidationImeModeValuesHiragana     DataValidationImeModeValues = "hiragana"
	DataValidationImeModeValuesFullkatakana DataValidationImeModeValues = "fullKatakana"
	DataValidationImeModeValuesHalfkatakana DataValidationImeModeValues = "halfKatakana"
	DataValidationImeModeValuesFullalpha    DataValidationImeModeValues = "fullAlpha"
	DataValidationImeModeValuesHalfalpha    DataValidationImeModeValues = "halfAlpha"
	DataValidationImeModeValuesFullhangul   DataValidationImeModeValues = "fullHangul"
	DataValidationImeModeValuesHalfhangul   DataValidationImeModeValues = "halfHangul"
)

func (e DataValidationImeModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DataValidationImeModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DataValidationImeModeValues(attr.Value)
	return nil
}

// UnderlineValues
type UnderlineValues string

const (
	UnderlineValuesSingle          UnderlineValues = "single"
	UnderlineValuesWords           UnderlineValues = "words"
	UnderlineValuesDouble          UnderlineValues = "double"
	UnderlineValuesThick           UnderlineValues = "thick"
	UnderlineValuesDotted          UnderlineValues = "dotted"
	UnderlineValuesDottedheavy     UnderlineValues = "dottedHeavy"
	UnderlineValuesDash            UnderlineValues = "dash"
	UnderlineValuesDashedheavy     UnderlineValues = "dashedHeavy"
	UnderlineValuesDashlong        UnderlineValues = "dashLong"
	UnderlineValuesDashlongheavy   UnderlineValues = "dashLongHeavy"
	UnderlineValuesDotdash         UnderlineValues = "dotDash"
	UnderlineValuesDashdotheavy    UnderlineValues = "dashDotHeavy"
	UnderlineValuesDotdotdash      UnderlineValues = "dotDotDash"
	UnderlineValuesDashdotdotheavy UnderlineValues = "dashDotDotHeavy"
	UnderlineValuesWave            UnderlineValues = "wave"
	UnderlineValuesWavyheavy       UnderlineValues = "wavyHeavy"
	UnderlineValuesWavydouble      UnderlineValues = "wavyDouble"
	UnderlineValuesNone            UnderlineValues = "none"
)

func (e UnderlineValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *UnderlineValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = UnderlineValues(attr.Value)
	return nil
}

// AnimateColorDirectionValues
type AnimateColorDirectionValues string

const (
	AnimateColorDirectionValuesCw  AnimateColorDirectionValues = "cw"
	AnimateColorDirectionValuesCcw AnimateColorDirectionValues = "ccw"
)

func (e AnimateColorDirectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AnimateColorDirectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AnimateColorDirectionValues(attr.Value)
	return nil
}

// RuleValues
type RuleValues string

const (
	RuleValuesNone   RuleValues = "none"
	RuleValuesAll    RuleValues = "all"
	RuleValuesRow    RuleValues = "row"
	RuleValuesColumn RuleValues = "column"
)

func (e RuleValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *RuleValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = RuleValues(attr.Value)
	return nil
}

// PivotTableAxisValues
type PivotTableAxisValues string

const (
	PivotTableAxisValuesAxisrow    PivotTableAxisValues = "axisRow"
	PivotTableAxisValuesAxiscol    PivotTableAxisValues = "axisCol"
	PivotTableAxisValuesAxispage   PivotTableAxisValues = "axisPage"
	PivotTableAxisValuesAxisvalues PivotTableAxisValues = "axisValues"
)

func (e PivotTableAxisValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PivotTableAxisValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PivotTableAxisValues(attr.Value)
	return nil
}

// DocumentProtectionValues
type DocumentProtectionValues string

const (
	DocumentProtectionValuesNone           DocumentProtectionValues = "none"
	DocumentProtectionValuesReadonly       DocumentProtectionValues = "readOnly"
	DocumentProtectionValuesComments       DocumentProtectionValues = "comments"
	DocumentProtectionValuesTrackedchanges DocumentProtectionValues = "trackedChanges"
	DocumentProtectionValuesForms          DocumentProtectionValues = "forms"
)

func (e DocumentProtectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DocumentProtectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DocumentProtectionValues(attr.Value)
	return nil
}

// TextBoxTightWrapValues
type TextBoxTightWrapValues string

const (
	TextBoxTightWrapValuesNone             TextBoxTightWrapValues = "none"
	TextBoxTightWrapValuesAlllines         TextBoxTightWrapValues = "allLines"
	TextBoxTightWrapValuesFirstandlastline TextBoxTightWrapValues = "firstAndLastLine"
	TextBoxTightWrapValuesFirstlineonly    TextBoxTightWrapValues = "firstLineOnly"
	TextBoxTightWrapValuesLastlineonly     TextBoxTightWrapValues = "lastLineOnly"
)

func (e TextBoxTightWrapValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextBoxTightWrapValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextBoxTightWrapValues(attr.Value)
	return nil
}

// SectionMarkValues
type SectionMarkValues string

const (
	SectionMarkValuesNextpage   SectionMarkValues = "nextPage"
	SectionMarkValuesNextcolumn SectionMarkValues = "nextColumn"
	SectionMarkValuesContinuous SectionMarkValues = "continuous"
	SectionMarkValuesEvenpage   SectionMarkValues = "evenPage"
	SectionMarkValuesOddpage    SectionMarkValues = "oddPage"
)

func (e SectionMarkValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SectionMarkValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SectionMarkValues(attr.Value)
	return nil
}

// ThemeFontValues
type ThemeFontValues string

const (
	ThemeFontValuesMajoreastasia ThemeFontValues = "majorEastAsia"
	ThemeFontValuesMajorbidi     ThemeFontValues = "majorBidi"
	ThemeFontValuesMajorascii    ThemeFontValues = "majorAscii"
	ThemeFontValuesMajorhansi    ThemeFontValues = "majorHAnsi"
	ThemeFontValuesMinoreastasia ThemeFontValues = "minorEastAsia"
	ThemeFontValuesMinorbidi     ThemeFontValues = "minorBidi"
	ThemeFontValuesMinorascii    ThemeFontValues = "minorAscii"
	ThemeFontValuesMinorhansi    ThemeFontValues = "minorHAnsi"
)

func (e ThemeFontValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ThemeFontValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ThemeFontValues(attr.Value)
	return nil
}

// OrientationValues
type OrientationValues string

const (
	OrientationValuesDefault   OrientationValues = "default"
	OrientationValuesPortrait  OrientationValues = "portrait"
	OrientationValuesLandscape OrientationValues = "landscape"
)

func (e OrientationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *OrientationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = OrientationValues(attr.Value)
	return nil
}

// PaneStateValues
type PaneStateValues string

const (
	PaneStateValuesSplit       PaneStateValues = "split"
	PaneStateValuesFrozen      PaneStateValues = "frozen"
	PaneStateValuesFrozensplit PaneStateValues = "frozenSplit"
)

func (e PaneStateValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PaneStateValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PaneStateValues(attr.Value)
	return nil
}

// GradientValues
type GradientValues string

const (
	GradientValuesLinear GradientValues = "linear"
	GradientValuesPath   GradientValues = "path"
)

func (e GradientValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *GradientValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = GradientValues(attr.Value)
	return nil
}

// ObjectDisplayValues
type ObjectDisplayValues string

const (
	ObjectDisplayValuesAll          ObjectDisplayValues = "all"
	ObjectDisplayValuesPlaceholders ObjectDisplayValues = "placeholders"
	ObjectDisplayValuesNone         ObjectDisplayValues = "none"
)

func (e ObjectDisplayValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ObjectDisplayValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ObjectDisplayValues(attr.Value)
	return nil
}

// VerticalTextAlignmentValues
type VerticalTextAlignmentValues string

const (
	VerticalTextAlignmentValuesTop      VerticalTextAlignmentValues = "top"
	VerticalTextAlignmentValuesCenter   VerticalTextAlignmentValues = "center"
	VerticalTextAlignmentValuesBaseline VerticalTextAlignmentValues = "baseline"
	VerticalTextAlignmentValuesBottom   VerticalTextAlignmentValues = "bottom"
	VerticalTextAlignmentValuesAuto     VerticalTextAlignmentValues = "auto"
)

func (e VerticalTextAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *VerticalTextAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = VerticalTextAlignmentValues(attr.Value)
	return nil
}

// DisplacedByCustomXmlValues
type DisplacedByCustomXmlValues string

const (
	DisplacedByCustomXmlValuesNext DisplacedByCustomXmlValues = "next"
	DisplacedByCustomXmlValuesPrev DisplacedByCustomXmlValues = "prev"
)

func (e DisplacedByCustomXmlValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DisplacedByCustomXmlValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DisplacedByCustomXmlValues(attr.Value)
	return nil
}

// TabStopValues
type TabStopValues string

const (
	TabStopValuesClear   TabStopValues = "clear"
	TabStopValuesLeft    TabStopValues = "left"
	TabStopValuesStart   TabStopValues = "start"
	TabStopValuesCenter  TabStopValues = "center"
	TabStopValuesRight   TabStopValues = "right"
	TabStopValuesEnd     TabStopValues = "end"
	TabStopValuesDecimal TabStopValues = "decimal"
	TabStopValuesBar     TabStopValues = "bar"
	TabStopValuesNum     TabStopValues = "num"
)

func (e TabStopValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TabStopValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TabStopValues(attr.Value)
	return nil
}

// ClassificationOutcomeType
type ClassificationOutcomeType string

const (
	ClassificationOutcomeTypeNone      ClassificationOutcomeType = "none"
	ClassificationOutcomeTypeHdr       ClassificationOutcomeType = "hdr"
	ClassificationOutcomeTypeFtr       ClassificationOutcomeType = "ftr"
	ClassificationOutcomeTypeWatermark ClassificationOutcomeType = "watermark"
)

func (e ClassificationOutcomeType) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ClassificationOutcomeType) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ClassificationOutcomeType(attr.Value)
	return nil
}

// SparklineAxisMinMaxValues
type SparklineAxisMinMaxValues string

const (
	SparklineAxisMinMaxValuesIndividual SparklineAxisMinMaxValues = "individual"
	SparklineAxisMinMaxValuesGroup      SparklineAxisMinMaxValues = "group"
	SparklineAxisMinMaxValuesCustom     SparklineAxisMinMaxValues = "custom"
)

func (e SparklineAxisMinMaxValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SparklineAxisMinMaxValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SparklineAxisMinMaxValues(attr.Value)
	return nil
}

// ModelTimeGroupingContentType
type ModelTimeGroupingContentType string

const (
	ModelTimeGroupingContentTypeYears       ModelTimeGroupingContentType = "years"
	ModelTimeGroupingContentTypeQuarters    ModelTimeGroupingContentType = "quarters"
	ModelTimeGroupingContentTypeMonthsindex ModelTimeGroupingContentType = "monthsindex"
	ModelTimeGroupingContentTypeMonths      ModelTimeGroupingContentType = "months"
	ModelTimeGroupingContentTypeDaysindex   ModelTimeGroupingContentType = "daysindex"
	ModelTimeGroupingContentTypeDays        ModelTimeGroupingContentType = "days"
	ModelTimeGroupingContentTypeHours       ModelTimeGroupingContentType = "hours"
	ModelTimeGroupingContentTypeMinutes     ModelTimeGroupingContentType = "minutes"
	ModelTimeGroupingContentTypeSeconds     ModelTimeGroupingContentType = "seconds"
)

func (e ModelTimeGroupingContentType) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ModelTimeGroupingContentType) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ModelTimeGroupingContentType(attr.Value)
	return nil
}

// ShapeTypeValues
type ShapeTypeValues string

const (
	ShapeTypeValuesLine                       ShapeTypeValues = "line"
	ShapeTypeValuesLineinv                    ShapeTypeValues = "lineInv"
	ShapeTypeValuesTriangle                   ShapeTypeValues = "triangle"
	ShapeTypeValuesRttriangle                 ShapeTypeValues = "rtTriangle"
	ShapeTypeValuesRect                       ShapeTypeValues = "rect"
	ShapeTypeValuesDiamond                    ShapeTypeValues = "diamond"
	ShapeTypeValuesParallelogram              ShapeTypeValues = "parallelogram"
	ShapeTypeValuesTrapezoid                  ShapeTypeValues = "trapezoid"
	ShapeTypeValuesNonisoscelestrapezoid      ShapeTypeValues = "nonIsoscelesTrapezoid"
	ShapeTypeValuesPentagon                   ShapeTypeValues = "pentagon"
	ShapeTypeValuesHexagon                    ShapeTypeValues = "hexagon"
	ShapeTypeValuesHeptagon                   ShapeTypeValues = "heptagon"
	ShapeTypeValuesOctagon                    ShapeTypeValues = "octagon"
	ShapeTypeValuesDecagon                    ShapeTypeValues = "decagon"
	ShapeTypeValuesDodecagon                  ShapeTypeValues = "dodecagon"
	ShapeTypeValuesStar4                      ShapeTypeValues = "star4"
	ShapeTypeValuesStar5                      ShapeTypeValues = "star5"
	ShapeTypeValuesStar6                      ShapeTypeValues = "star6"
	ShapeTypeValuesStar7                      ShapeTypeValues = "star7"
	ShapeTypeValuesStar8                      ShapeTypeValues = "star8"
	ShapeTypeValuesStar10                     ShapeTypeValues = "star10"
	ShapeTypeValuesStar12                     ShapeTypeValues = "star12"
	ShapeTypeValuesStar16                     ShapeTypeValues = "star16"
	ShapeTypeValuesStar24                     ShapeTypeValues = "star24"
	ShapeTypeValuesStar32                     ShapeTypeValues = "star32"
	ShapeTypeValuesRoundrect                  ShapeTypeValues = "roundRect"
	ShapeTypeValuesRound1rect                 ShapeTypeValues = "round1Rect"
	ShapeTypeValuesRound2samerect             ShapeTypeValues = "round2SameRect"
	ShapeTypeValuesRound2diagrect             ShapeTypeValues = "round2DiagRect"
	ShapeTypeValuesSniproundrect              ShapeTypeValues = "snipRoundRect"
	ShapeTypeValuesSnip1rect                  ShapeTypeValues = "snip1Rect"
	ShapeTypeValuesSnip2samerect              ShapeTypeValues = "snip2SameRect"
	ShapeTypeValuesSnip2diagrect              ShapeTypeValues = "snip2DiagRect"
	ShapeTypeValuesPlaque                     ShapeTypeValues = "plaque"
	ShapeTypeValuesEllipse                    ShapeTypeValues = "ellipse"
	ShapeTypeValuesTeardrop                   ShapeTypeValues = "teardrop"
	ShapeTypeValuesHomeplate                  ShapeTypeValues = "homePlate"
	ShapeTypeValuesChevron                    ShapeTypeValues = "chevron"
	ShapeTypeValuesPiewedge                   ShapeTypeValues = "pieWedge"
	ShapeTypeValuesPie                        ShapeTypeValues = "pie"
	ShapeTypeValuesBlockarc                   ShapeTypeValues = "blockArc"
	ShapeTypeValuesDonut                      ShapeTypeValues = "donut"
	ShapeTypeValuesNosmoking                  ShapeTypeValues = "noSmoking"
	ShapeTypeValuesRightarrow                 ShapeTypeValues = "rightArrow"
	ShapeTypeValuesLeftarrow                  ShapeTypeValues = "leftArrow"
	ShapeTypeValuesUparrow                    ShapeTypeValues = "upArrow"
	ShapeTypeValuesDownarrow                  ShapeTypeValues = "downArrow"
	ShapeTypeValuesStripedrightarrow          ShapeTypeValues = "stripedRightArrow"
	ShapeTypeValuesNotchedrightarrow          ShapeTypeValues = "notchedRightArrow"
	ShapeTypeValuesBentuparrow                ShapeTypeValues = "bentUpArrow"
	ShapeTypeValuesLeftrightarrow             ShapeTypeValues = "leftRightArrow"
	ShapeTypeValuesUpdownarrow                ShapeTypeValues = "upDownArrow"
	ShapeTypeValuesLeftuparrow                ShapeTypeValues = "leftUpArrow"
	ShapeTypeValuesLeftrightuparrow           ShapeTypeValues = "leftRightUpArrow"
	ShapeTypeValuesQuadarrow                  ShapeTypeValues = "quadArrow"
	ShapeTypeValuesLeftarrowcallout           ShapeTypeValues = "leftArrowCallout"
	ShapeTypeValuesRightarrowcallout          ShapeTypeValues = "rightArrowCallout"
	ShapeTypeValuesUparrowcallout             ShapeTypeValues = "upArrowCallout"
	ShapeTypeValuesDownarrowcallout           ShapeTypeValues = "downArrowCallout"
	ShapeTypeValuesLeftrightarrowcallout      ShapeTypeValues = "leftRightArrowCallout"
	ShapeTypeValuesUpdownarrowcallout         ShapeTypeValues = "upDownArrowCallout"
	ShapeTypeValuesQuadarrowcallout           ShapeTypeValues = "quadArrowCallout"
	ShapeTypeValuesBentarrow                  ShapeTypeValues = "bentArrow"
	ShapeTypeValuesUturnarrow                 ShapeTypeValues = "uturnArrow"
	ShapeTypeValuesCirculararrow              ShapeTypeValues = "circularArrow"
	ShapeTypeValuesLeftcirculararrow          ShapeTypeValues = "leftCircularArrow"
	ShapeTypeValuesLeftrightcirculararrow     ShapeTypeValues = "leftRightCircularArrow"
	ShapeTypeValuesCurvedrightarrow           ShapeTypeValues = "curvedRightArrow"
	ShapeTypeValuesCurvedleftarrow            ShapeTypeValues = "curvedLeftArrow"
	ShapeTypeValuesCurveduparrow              ShapeTypeValues = "curvedUpArrow"
	ShapeTypeValuesCurveddownarrow            ShapeTypeValues = "curvedDownArrow"
	ShapeTypeValuesSwoosharrow                ShapeTypeValues = "swooshArrow"
	ShapeTypeValuesCube                       ShapeTypeValues = "cube"
	ShapeTypeValuesCan                        ShapeTypeValues = "can"
	ShapeTypeValuesLightningbolt              ShapeTypeValues = "lightningBolt"
	ShapeTypeValuesHeart                      ShapeTypeValues = "heart"
	ShapeTypeValuesSun                        ShapeTypeValues = "sun"
	ShapeTypeValuesMoon                       ShapeTypeValues = "moon"
	ShapeTypeValuesSmileyface                 ShapeTypeValues = "smileyFace"
	ShapeTypeValuesIrregularseal1             ShapeTypeValues = "irregularSeal1"
	ShapeTypeValuesIrregularseal2             ShapeTypeValues = "irregularSeal2"
	ShapeTypeValuesFoldedcorner               ShapeTypeValues = "foldedCorner"
	ShapeTypeValuesBevel                      ShapeTypeValues = "bevel"
	ShapeTypeValuesFrame                      ShapeTypeValues = "frame"
	ShapeTypeValuesHalfframe                  ShapeTypeValues = "halfFrame"
	ShapeTypeValuesCorner                     ShapeTypeValues = "corner"
	ShapeTypeValuesDiagstripe                 ShapeTypeValues = "diagStripe"
	ShapeTypeValuesChord                      ShapeTypeValues = "chord"
	ShapeTypeValuesArc                        ShapeTypeValues = "arc"
	ShapeTypeValuesLeftbracket                ShapeTypeValues = "leftBracket"
	ShapeTypeValuesRightbracket               ShapeTypeValues = "rightBracket"
	ShapeTypeValuesLeftbrace                  ShapeTypeValues = "leftBrace"
	ShapeTypeValuesRightbrace                 ShapeTypeValues = "rightBrace"
	ShapeTypeValuesBracketpair                ShapeTypeValues = "bracketPair"
	ShapeTypeValuesBracepair                  ShapeTypeValues = "bracePair"
	ShapeTypeValuesStraightconnector1         ShapeTypeValues = "straightConnector1"
	ShapeTypeValuesBentconnector2             ShapeTypeValues = "bentConnector2"
	ShapeTypeValuesBentconnector3             ShapeTypeValues = "bentConnector3"
	ShapeTypeValuesBentconnector4             ShapeTypeValues = "bentConnector4"
	ShapeTypeValuesBentconnector5             ShapeTypeValues = "bentConnector5"
	ShapeTypeValuesCurvedconnector2           ShapeTypeValues = "curvedConnector2"
	ShapeTypeValuesCurvedconnector3           ShapeTypeValues = "curvedConnector3"
	ShapeTypeValuesCurvedconnector4           ShapeTypeValues = "curvedConnector4"
	ShapeTypeValuesCurvedconnector5           ShapeTypeValues = "curvedConnector5"
	ShapeTypeValuesCallout1                   ShapeTypeValues = "callout1"
	ShapeTypeValuesCallout2                   ShapeTypeValues = "callout2"
	ShapeTypeValuesCallout3                   ShapeTypeValues = "callout3"
	ShapeTypeValuesAccentcallout1             ShapeTypeValues = "accentCallout1"
	ShapeTypeValuesAccentcallout2             ShapeTypeValues = "accentCallout2"
	ShapeTypeValuesAccentcallout3             ShapeTypeValues = "accentCallout3"
	ShapeTypeValuesBordercallout1             ShapeTypeValues = "borderCallout1"
	ShapeTypeValuesBordercallout2             ShapeTypeValues = "borderCallout2"
	ShapeTypeValuesBordercallout3             ShapeTypeValues = "borderCallout3"
	ShapeTypeValuesAccentbordercallout1       ShapeTypeValues = "accentBorderCallout1"
	ShapeTypeValuesAccentbordercallout2       ShapeTypeValues = "accentBorderCallout2"
	ShapeTypeValuesAccentbordercallout3       ShapeTypeValues = "accentBorderCallout3"
	ShapeTypeValuesWedgerectcallout           ShapeTypeValues = "wedgeRectCallout"
	ShapeTypeValuesWedgeroundrectcallout      ShapeTypeValues = "wedgeRoundRectCallout"
	ShapeTypeValuesWedgeellipsecallout        ShapeTypeValues = "wedgeEllipseCallout"
	ShapeTypeValuesCloudcallout               ShapeTypeValues = "cloudCallout"
	ShapeTypeValuesCloud                      ShapeTypeValues = "cloud"
	ShapeTypeValuesRibbon                     ShapeTypeValues = "ribbon"
	ShapeTypeValuesRibbon2                    ShapeTypeValues = "ribbon2"
	ShapeTypeValuesEllipseribbon              ShapeTypeValues = "ellipseRibbon"
	ShapeTypeValuesEllipseribbon2             ShapeTypeValues = "ellipseRibbon2"
	ShapeTypeValuesLeftrightribbon            ShapeTypeValues = "leftRightRibbon"
	ShapeTypeValuesVerticalscroll             ShapeTypeValues = "verticalScroll"
	ShapeTypeValuesHorizontalscroll           ShapeTypeValues = "horizontalScroll"
	ShapeTypeValuesWave                       ShapeTypeValues = "wave"
	ShapeTypeValuesDoublewave                 ShapeTypeValues = "doubleWave"
	ShapeTypeValuesPlus                       ShapeTypeValues = "plus"
	ShapeTypeValuesFlowchartprocess           ShapeTypeValues = "flowChartProcess"
	ShapeTypeValuesFlowchartdecision          ShapeTypeValues = "flowChartDecision"
	ShapeTypeValuesFlowchartinputoutput       ShapeTypeValues = "flowChartInputOutput"
	ShapeTypeValuesFlowchartpredefinedprocess ShapeTypeValues = "flowChartPredefinedProcess"
	ShapeTypeValuesFlowchartinternalstorage   ShapeTypeValues = "flowChartInternalStorage"
	ShapeTypeValuesFlowchartdocument          ShapeTypeValues = "flowChartDocument"
	ShapeTypeValuesFlowchartmultidocument     ShapeTypeValues = "flowChartMultidocument"
	ShapeTypeValuesFlowchartterminator        ShapeTypeValues = "flowChartTerminator"
	ShapeTypeValuesFlowchartpreparation       ShapeTypeValues = "flowChartPreparation"
	ShapeTypeValuesFlowchartmanualinput       ShapeTypeValues = "flowChartManualInput"
	ShapeTypeValuesFlowchartmanualoperation   ShapeTypeValues = "flowChartManualOperation"
	ShapeTypeValuesFlowchartconnector         ShapeTypeValues = "flowChartConnector"
	ShapeTypeValuesFlowchartpunchedcard       ShapeTypeValues = "flowChartPunchedCard"
	ShapeTypeValuesFlowchartpunchedtape       ShapeTypeValues = "flowChartPunchedTape"
	ShapeTypeValuesFlowchartsummingjunction   ShapeTypeValues = "flowChartSummingJunction"
	ShapeTypeValuesFlowchartor                ShapeTypeValues = "flowChartOr"
	ShapeTypeValuesFlowchartcollate           ShapeTypeValues = "flowChartCollate"
	ShapeTypeValuesFlowchartsort              ShapeTypeValues = "flowChartSort"
	ShapeTypeValuesFlowchartextract           ShapeTypeValues = "flowChartExtract"
	ShapeTypeValuesFlowchartmerge             ShapeTypeValues = "flowChartMerge"
	ShapeTypeValuesFlowchartofflinestorage    ShapeTypeValues = "flowChartOfflineStorage"
	ShapeTypeValuesFlowchartonlinestorage     ShapeTypeValues = "flowChartOnlineStorage"
	ShapeTypeValuesFlowchartmagnetictape      ShapeTypeValues = "flowChartMagneticTape"
	ShapeTypeValuesFlowchartmagneticdisk      ShapeTypeValues = "flowChartMagneticDisk"
	ShapeTypeValuesFlowchartmagneticdrum      ShapeTypeValues = "flowChartMagneticDrum"
	ShapeTypeValuesFlowchartdisplay           ShapeTypeValues = "flowChartDisplay"
	ShapeTypeValuesFlowchartdelay             ShapeTypeValues = "flowChartDelay"
	ShapeTypeValuesFlowchartalternateprocess  ShapeTypeValues = "flowChartAlternateProcess"
	ShapeTypeValuesFlowchartoffpageconnector  ShapeTypeValues = "flowChartOffpageConnector"
	ShapeTypeValuesActionbuttonblank          ShapeTypeValues = "actionButtonBlank"
	ShapeTypeValuesActionbuttonhome           ShapeTypeValues = "actionButtonHome"
	ShapeTypeValuesActionbuttonhelp           ShapeTypeValues = "actionButtonHelp"
	ShapeTypeValuesActionbuttoninformation    ShapeTypeValues = "actionButtonInformation"
	ShapeTypeValuesActionbuttonforwardnext    ShapeTypeValues = "actionButtonForwardNext"
	ShapeTypeValuesActionbuttonbackprevious   ShapeTypeValues = "actionButtonBackPrevious"
	ShapeTypeValuesActionbuttonend            ShapeTypeValues = "actionButtonEnd"
	ShapeTypeValuesActionbuttonbeginning      ShapeTypeValues = "actionButtonBeginning"
	ShapeTypeValuesActionbuttonreturn         ShapeTypeValues = "actionButtonReturn"
	ShapeTypeValuesActionbuttondocument       ShapeTypeValues = "actionButtonDocument"
	ShapeTypeValuesActionbuttonsound          ShapeTypeValues = "actionButtonSound"
	ShapeTypeValuesActionbuttonmovie          ShapeTypeValues = "actionButtonMovie"
	ShapeTypeValuesGear6                      ShapeTypeValues = "gear6"
	ShapeTypeValuesGear9                      ShapeTypeValues = "gear9"
	ShapeTypeValuesFunnel                     ShapeTypeValues = "funnel"
	ShapeTypeValuesMathplus                   ShapeTypeValues = "mathPlus"
	ShapeTypeValuesMathminus                  ShapeTypeValues = "mathMinus"
	ShapeTypeValuesMathmultiply               ShapeTypeValues = "mathMultiply"
	ShapeTypeValuesMathdivide                 ShapeTypeValues = "mathDivide"
	ShapeTypeValuesMathequal                  ShapeTypeValues = "mathEqual"
	ShapeTypeValuesMathnotequal               ShapeTypeValues = "mathNotEqual"
	ShapeTypeValuesCornertabs                 ShapeTypeValues = "cornerTabs"
	ShapeTypeValuesSquaretabs                 ShapeTypeValues = "squareTabs"
	ShapeTypeValuesPlaquetabs                 ShapeTypeValues = "plaqueTabs"
	ShapeTypeValuesChartx                     ShapeTypeValues = "chartX"
	ShapeTypeValuesChartstar                  ShapeTypeValues = "chartStar"
	ShapeTypeValuesChartplus                  ShapeTypeValues = "chartPlus"
)

func (e ShapeTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ShapeTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ShapeTypeValues(attr.Value)
	return nil
}

// PlaceholderValues
type PlaceholderValues string

const (
	PlaceholderValuesTitle    PlaceholderValues = "title"
	PlaceholderValuesBody     PlaceholderValues = "body"
	PlaceholderValuesCtrtitle PlaceholderValues = "ctrTitle"
	PlaceholderValuesSubtitle PlaceholderValues = "subTitle"
	PlaceholderValuesDt       PlaceholderValues = "dt"
	PlaceholderValuesSldnum   PlaceholderValues = "sldNum"
	PlaceholderValuesFtr      PlaceholderValues = "ftr"
	PlaceholderValuesHdr      PlaceholderValues = "hdr"
	PlaceholderValuesObj      PlaceholderValues = "obj"
	PlaceholderValuesChart    PlaceholderValues = "chart"
	PlaceholderValuesTbl      PlaceholderValues = "tbl"
	PlaceholderValuesClipart  PlaceholderValues = "clipArt"
	PlaceholderValuesDgm      PlaceholderValues = "dgm"
	PlaceholderValuesMedia    PlaceholderValues = "media"
	PlaceholderValuesSldimg   PlaceholderValues = "sldImg"
	PlaceholderValuesPic      PlaceholderValues = "pic"
)

func (e PlaceholderValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PlaceholderValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PlaceholderValues(attr.Value)
	return nil
}

// FileTypeValues
type FileTypeValues string

const (
	FileTypeValuesMac FileTypeValues = "mac"
	FileTypeValuesWin FileTypeValues = "win"
	FileTypeValuesDos FileTypeValues = "dos"
)

func (e FileTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FileTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FileTypeValues(attr.Value)
	return nil
}

// ConformanceClass
type ConformanceClass string

const (
	ConformanceClassStrict       ConformanceClass = "strict"
	ConformanceClassTransitional ConformanceClass = "transitional"
)

func (e ConformanceClass) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ConformanceClass) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ConformanceClass(attr.Value)
	return nil
}

// DocPartBehaviorValues
type DocPartBehaviorValues string

const (
	DocPartBehaviorValuesContent DocPartBehaviorValues = "content"
	DocPartBehaviorValuesP       DocPartBehaviorValues = "p"
	DocPartBehaviorValuesPg      DocPartBehaviorValues = "pg"
)

func (e DocPartBehaviorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DocPartBehaviorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DocPartBehaviorValues(attr.Value)
	return nil
}

// BehaviorTransformValues
type BehaviorTransformValues string

const (
	BehaviorTransformValuesPt  BehaviorTransformValues = "pt"
	BehaviorTransformValuesImg BehaviorTransformValues = "img"
)

func (e BehaviorTransformValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BehaviorTransformValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BehaviorTransformValues(attr.Value)
	return nil
}

// PaneValues
type PaneValues string

const (
	PaneValuesBottomright PaneValues = "bottomRight"
	PaneValuesTopright    PaneValues = "topRight"
	PaneValuesBottomleft  PaneValues = "bottomLeft"
	PaneValuesTopleft     PaneValues = "topLeft"
)

func (e PaneValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PaneValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PaneValues(attr.Value)
	return nil
}

// DocumentTypeValues
type DocumentTypeValues string

const (
	DocumentTypeValuesNotspecified DocumentTypeValues = "notSpecified"
	DocumentTypeValuesLetter       DocumentTypeValues = "letter"
	DocumentTypeValuesEmail        DocumentTypeValues = "eMail"
)

func (e DocumentTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DocumentTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DocumentTypeValues(attr.Value)
	return nil
}

// DocPartValues
type DocPartValues string

const (
	DocPartValuesNone     DocPartValues = "none"
	DocPartValuesNormal   DocPartValues = "normal"
	DocPartValuesAutoexp  DocPartValues = "autoExp"
	DocPartValuesToolbar  DocPartValues = "toolbar"
	DocPartValuesSpeller  DocPartValues = "speller"
	DocPartValuesFormfld  DocPartValues = "formFld"
	DocPartValuesBbplchdr DocPartValues = "bbPlcHdr"
)

func (e DocPartValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DocPartValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DocPartValues(attr.Value)
	return nil
}

// SlicerStyleTypeValues
type SlicerStyleTypeValues string

const (
	SlicerStyleTypeValuesUnselecteditemwithdata          SlicerStyleTypeValues = "unselectedItemWithData"
	SlicerStyleTypeValuesSelecteditemwithdata            SlicerStyleTypeValues = "selectedItemWithData"
	SlicerStyleTypeValuesUnselecteditemwithnodata        SlicerStyleTypeValues = "unselectedItemWithNoData"
	SlicerStyleTypeValuesSelecteditemwithnodata          SlicerStyleTypeValues = "selectedItemWithNoData"
	SlicerStyleTypeValuesHoveredunselecteditemwithdata   SlicerStyleTypeValues = "hoveredUnselectedItemWithData"
	SlicerStyleTypeValuesHoveredselecteditemwithdata     SlicerStyleTypeValues = "hoveredSelectedItemWithData"
	SlicerStyleTypeValuesHoveredunselecteditemwithnodata SlicerStyleTypeValues = "hoveredUnselectedItemWithNoData"
	SlicerStyleTypeValuesHoveredselecteditemwithnodata   SlicerStyleTypeValues = "hoveredSelectedItemWithNoData"
)

func (e SlicerStyleTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SlicerStyleTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SlicerStyleTypeValues(attr.Value)
	return nil
}

// TransitionSlideDirectionValues
type TransitionSlideDirectionValues string

const (
	TransitionSlideDirectionValuesL TransitionSlideDirectionValues = "l"
	TransitionSlideDirectionValuesU TransitionSlideDirectionValues = "u"
	TransitionSlideDirectionValuesR TransitionSlideDirectionValues = "r"
	TransitionSlideDirectionValuesD TransitionSlideDirectionValues = "d"
)

func (e TransitionSlideDirectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TransitionSlideDirectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TransitionSlideDirectionValues(
		attr.Value,
	)
	return nil
}

// DiagramBuildValues
type DiagramBuildValues string

const (
	DiagramBuildValuesWhole         DiagramBuildValues = "whole"
	DiagramBuildValuesDepthbynode   DiagramBuildValues = "depthByNode"
	DiagramBuildValuesDepthbybranch DiagramBuildValues = "depthByBranch"
	DiagramBuildValuesBreadthbynode DiagramBuildValues = "breadthByNode"
	DiagramBuildValuesBreadthbylvl  DiagramBuildValues = "breadthByLvl"
	DiagramBuildValuesCw            DiagramBuildValues = "cw"
	DiagramBuildValuesCwin          DiagramBuildValues = "cwIn"
	DiagramBuildValuesCwout         DiagramBuildValues = "cwOut"
	DiagramBuildValuesCcw           DiagramBuildValues = "ccw"
	DiagramBuildValuesCcwin         DiagramBuildValues = "ccwIn"
	DiagramBuildValuesCcwout        DiagramBuildValues = "ccwOut"
	DiagramBuildValuesInbyring      DiagramBuildValues = "inByRing"
	DiagramBuildValuesOutbyring     DiagramBuildValues = "outByRing"
	DiagramBuildValuesUp            DiagramBuildValues = "up"
	DiagramBuildValuesDown          DiagramBuildValues = "down"
	DiagramBuildValuesAllatonce     DiagramBuildValues = "allAtOnce"
	DiagramBuildValuesCust          DiagramBuildValues = "cust"
)

func (e DiagramBuildValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DiagramBuildValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DiagramBuildValues(attr.Value)
	return nil
}

// ConditionalFormattingOperatorValues
type ConditionalFormattingOperatorValues string

const (
	ConditionalFormattingOperatorValuesLessthan           ConditionalFormattingOperatorValues = "lessThan"
	ConditionalFormattingOperatorValuesLessthanorequal    ConditionalFormattingOperatorValues = "lessThanOrEqual"
	ConditionalFormattingOperatorValuesEqual              ConditionalFormattingOperatorValues = "equal"
	ConditionalFormattingOperatorValuesNotequal           ConditionalFormattingOperatorValues = "notEqual"
	ConditionalFormattingOperatorValuesGreaterthanorequal ConditionalFormattingOperatorValues = "greaterThanOrEqual"
	ConditionalFormattingOperatorValuesGreaterthan        ConditionalFormattingOperatorValues = "greaterThan"
	ConditionalFormattingOperatorValuesBetween            ConditionalFormattingOperatorValues = "between"
	ConditionalFormattingOperatorValuesNotbetween         ConditionalFormattingOperatorValues = "notBetween"
	ConditionalFormattingOperatorValuesContainstext       ConditionalFormattingOperatorValues = "containsText"
	ConditionalFormattingOperatorValuesNotcontains        ConditionalFormattingOperatorValues = "notContains"
	ConditionalFormattingOperatorValuesBeginswith         ConditionalFormattingOperatorValues = "beginsWith"
	ConditionalFormattingOperatorValuesEndswith           ConditionalFormattingOperatorValues = "endsWith"
)

func (e ConditionalFormattingOperatorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ConditionalFormattingOperatorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ConditionalFormattingOperatorValues(
		attr.Value,
	)
	return nil
}

// PageBorderOffsetValues
type PageBorderOffsetValues string

const (
	PageBorderOffsetValuesPage PageBorderOffsetValues = "page"
	PageBorderOffsetValuesText PageBorderOffsetValues = "text"
)

func (e PageBorderOffsetValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PageBorderOffsetValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PageBorderOffsetValues(attr.Value)
	return nil
}

// FootnotePositionValues
type FootnotePositionValues string

const (
	FootnotePositionValuesPagebottom  FootnotePositionValues = "pageBottom"
	FootnotePositionValuesBeneathtext FootnotePositionValues = "beneathText"
	FootnotePositionValuesSectend     FootnotePositionValues = "sectEnd"
)

func (e FootnotePositionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FootnotePositionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FootnotePositionValues(attr.Value)
	return nil
}

// RestartNumberValues
type RestartNumberValues string

const (
	RestartNumberValuesContinuous RestartNumberValues = "continuous"
	RestartNumberValuesEachsect   RestartNumberValues = "eachSect"
	RestartNumberValuesEachpage   RestartNumberValues = "eachPage"
)

func (e RestartNumberValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *RestartNumberValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = RestartNumberValues(attr.Value)
	return nil
}

// TransitionShredPatternValues
type TransitionShredPatternValues string

const (
	TransitionShredPatternValuesStrip     TransitionShredPatternValues = "strip"
	TransitionShredPatternValuesRectangle TransitionShredPatternValues = "rectangle"
)

func (e TransitionShredPatternValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TransitionShredPatternValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TransitionShredPatternValues(attr.Value)
	return nil
}

// TimelineStyleType
type TimelineStyleType string

const (
	TimelineStyleTypeSelectionlabel         TimelineStyleType = "selectionLabel"
	TimelineStyleTypeTimelevel              TimelineStyleType = "timeLevel"
	TimelineStyleTypePeriodlabel1           TimelineStyleType = "periodLabel1"
	TimelineStyleTypePeriodlabel2           TimelineStyleType = "periodLabel2"
	TimelineStyleTypeSelectedtimeblock      TimelineStyleType = "selectedTimeBlock"
	TimelineStyleTypeUnselectedtimeblock    TimelineStyleType = "unselectedTimeBlock"
	TimelineStyleTypeSelectedtimeblockspace TimelineStyleType = "selectedTimeBlockSpace"
)

func (e TimelineStyleType) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TimelineStyleType) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TimelineStyleType(attr.Value)
	return nil
}

// TextVerticalValues
type TextVerticalValues string

const (
	TextVerticalValuesHorz           TextVerticalValues = "horz"
	TextVerticalValuesVert           TextVerticalValues = "vert"
	TextVerticalValuesVert270        TextVerticalValues = "vert270"
	TextVerticalValuesWordartvert    TextVerticalValues = "wordArtVert"
	TextVerticalValuesEavert         TextVerticalValues = "eaVert"
	TextVerticalValuesMongolianvert  TextVerticalValues = "mongolianVert"
	TextVerticalValuesWordartvertrtl TextVerticalValues = "wordArtVertRtl"
)

func (e TextVerticalValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextVerticalValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextVerticalValues(attr.Value)
	return nil
}

// PhotoAlbumFrameShapeValues
type PhotoAlbumFrameShapeValues string

const (
	PhotoAlbumFrameShapeValuesFramestyle1 PhotoAlbumFrameShapeValues = "frameStyle1"
	PhotoAlbumFrameShapeValuesFramestyle2 PhotoAlbumFrameShapeValues = "frameStyle2"
	PhotoAlbumFrameShapeValuesFramestyle3 PhotoAlbumFrameShapeValues = "frameStyle3"
	PhotoAlbumFrameShapeValuesFramestyle4 PhotoAlbumFrameShapeValues = "frameStyle4"
	PhotoAlbumFrameShapeValuesFramestyle5 PhotoAlbumFrameShapeValues = "frameStyle5"
	PhotoAlbumFrameShapeValuesFramestyle6 PhotoAlbumFrameShapeValues = "frameStyle6"
	PhotoAlbumFrameShapeValuesFramestyle7 PhotoAlbumFrameShapeValues = "frameStyle7"
)

func (e PhotoAlbumFrameShapeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PhotoAlbumFrameShapeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PhotoAlbumFrameShapeValues(attr.Value)
	return nil
}

// VerticalAlignmentValues
type VerticalAlignmentValues string

const (
	VerticalAlignmentValuesInline  VerticalAlignmentValues = "inline"
	VerticalAlignmentValuesTop     VerticalAlignmentValues = "top"
	VerticalAlignmentValuesCenter  VerticalAlignmentValues = "center"
	VerticalAlignmentValuesBottom  VerticalAlignmentValues = "bottom"
	VerticalAlignmentValuesInside  VerticalAlignmentValues = "inside"
	VerticalAlignmentValuesOutside VerticalAlignmentValues = "outside"
)

func (e VerticalAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *VerticalAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = VerticalAlignmentValues(attr.Value)
	return nil
}

// TableStyleValues
type TableStyleValues string

const (
	TableStyleValuesWholetable             TableStyleValues = "wholeTable"
	TableStyleValuesHeaderrow              TableStyleValues = "headerRow"
	TableStyleValuesTotalrow               TableStyleValues = "totalRow"
	TableStyleValuesFirstcolumn            TableStyleValues = "firstColumn"
	TableStyleValuesLastcolumn             TableStyleValues = "lastColumn"
	TableStyleValuesFirstrowstripe         TableStyleValues = "firstRowStripe"
	TableStyleValuesSecondrowstripe        TableStyleValues = "secondRowStripe"
	TableStyleValuesFirstcolumnstripe      TableStyleValues = "firstColumnStripe"
	TableStyleValuesSecondcolumnstripe     TableStyleValues = "secondColumnStripe"
	TableStyleValuesFirstheadercell        TableStyleValues = "firstHeaderCell"
	TableStyleValuesLastheadercell         TableStyleValues = "lastHeaderCell"
	TableStyleValuesFirsttotalcell         TableStyleValues = "firstTotalCell"
	TableStyleValuesLasttotalcell          TableStyleValues = "lastTotalCell"
	TableStyleValuesFirstsubtotalcolumn    TableStyleValues = "firstSubtotalColumn"
	TableStyleValuesSecondsubtotalcolumn   TableStyleValues = "secondSubtotalColumn"
	TableStyleValuesThirdsubtotalcolumn    TableStyleValues = "thirdSubtotalColumn"
	TableStyleValuesFirstsubtotalrow       TableStyleValues = "firstSubtotalRow"
	TableStyleValuesSecondsubtotalrow      TableStyleValues = "secondSubtotalRow"
	TableStyleValuesThirdsubtotalrow       TableStyleValues = "thirdSubtotalRow"
	TableStyleValuesBlankrow               TableStyleValues = "blankRow"
	TableStyleValuesFirstcolumnsubheading  TableStyleValues = "firstColumnSubheading"
	TableStyleValuesSecondcolumnsubheading TableStyleValues = "secondColumnSubheading"
	TableStyleValuesThirdcolumnsubheading  TableStyleValues = "thirdColumnSubheading"
	TableStyleValuesFirstrowsubheading     TableStyleValues = "firstRowSubheading"
	TableStyleValuesSecondrowsubheading    TableStyleValues = "secondRowSubheading"
	TableStyleValuesThirdrowsubheading     TableStyleValues = "thirdRowSubheading"
	TableStyleValuesPagefieldlabels        TableStyleValues = "pageFieldLabels"
	TableStyleValuesPagefieldvalues        TableStyleValues = "pageFieldValues"
)

func (e TableStyleValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TableStyleValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TableStyleValues(attr.Value)
	return nil
}

// LineNumberRestartValues
type LineNumberRestartValues string

const (
	LineNumberRestartValuesNewpage    LineNumberRestartValues = "newPage"
	LineNumberRestartValuesNewsection LineNumberRestartValues = "newSection"
	LineNumberRestartValuesContinuous LineNumberRestartValues = "continuous"
)

func (e LineNumberRestartValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LineNumberRestartValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LineNumberRestartValues(attr.Value)
	return nil
}

// TextHorizontalAlignmentValues
type TextHorizontalAlignmentValues string

const (
	TextHorizontalAlignmentValuesLeft        TextHorizontalAlignmentValues = "left"
	TextHorizontalAlignmentValuesCenter      TextHorizontalAlignmentValues = "center"
	TextHorizontalAlignmentValuesRight       TextHorizontalAlignmentValues = "right"
	TextHorizontalAlignmentValuesJustify     TextHorizontalAlignmentValues = "justify"
	TextHorizontalAlignmentValuesDistributed TextHorizontalAlignmentValues = "distributed"
)

func (e TextHorizontalAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextHorizontalAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextHorizontalAlignmentValues(attr.Value)
	return nil
}

// DdeValues
type DdeValues string

const (
	DdeValuesNil DdeValues = "nil"
	DdeValuesB   DdeValues = "b"
	DdeValuesN   DdeValues = "n"
	DdeValuesE   DdeValues = "e"
	DdeValuesStr DdeValues = "str"
)

func (e DdeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DdeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DdeValues(attr.Value)
	return nil
}

// SheetStateValues
type SheetStateValues string

const (
	SheetStateValuesVisible    SheetStateValues = "visible"
	SheetStateValuesHidden     SheetStateValues = "hidden"
	SheetStateValuesVeryhidden SheetStateValues = "veryHidden"
)

func (e SheetStateValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SheetStateValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SheetStateValues(attr.Value)
	return nil
}

// ChartBuildStepValues
type ChartBuildStepValues string

const (
	ChartBuildStepValuesCategory     ChartBuildStepValues = "category"
	ChartBuildStepValuesPtincategory ChartBuildStepValues = "ptInCategory"
	ChartBuildStepValuesSeries       ChartBuildStepValues = "series"
	ChartBuildStepValuesPtinseries   ChartBuildStepValues = "ptInSeries"
	ChartBuildStepValuesAllpts       ChartBuildStepValues = "allPts"
	ChartBuildStepValuesGridlegend   ChartBuildStepValues = "gridLegend"
)

func (e ChartBuildStepValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ChartBuildStepValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ChartBuildStepValues(attr.Value)
	return nil
}

// DirectionValues
type DirectionValues string

const (
	DirectionValuesLtr DirectionValues = "ltr"
	DirectionValuesRtl DirectionValues = "rtl"
)

func (e DirectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DirectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DirectionValues(attr.Value)
	return nil
}

// SortByValues
type SortByValues string

const (
	SortByValuesValue     SortByValues = "value"
	SortByValuesCellcolor SortByValues = "cellColor"
	SortByValuesFontcolor SortByValues = "fontColor"
	SortByValuesIcon      SortByValues = "icon"
)

func (e SortByValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SortByValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SortByValues(attr.Value)
	return nil
}

// LineSpacingRuleValues
type LineSpacingRuleValues string

const (
	LineSpacingRuleValuesAuto    LineSpacingRuleValues = "auto"
	LineSpacingRuleValuesExact   LineSpacingRuleValues = "exact"
	LineSpacingRuleValuesAtleast LineSpacingRuleValues = "atLeast"
)

func (e LineSpacingRuleValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LineSpacingRuleValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LineSpacingRuleValues(attr.Value)
	return nil
}

// PageBorderZOrderValues
type PageBorderZOrderValues string

const (
	PageBorderZOrderValuesFront PageBorderZOrderValues = "front"
	PageBorderZOrderValuesBack  PageBorderZOrderValues = "back"
)

func (e PageBorderZOrderValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PageBorderZOrderValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PageBorderZOrderValues(attr.Value)
	return nil
}

// RangePermissionEditingGroupValues
type RangePermissionEditingGroupValues string

const (
	RangePermissionEditingGroupValuesNone           RangePermissionEditingGroupValues = "none"
	RangePermissionEditingGroupValuesEveryone       RangePermissionEditingGroupValues = "everyone"
	RangePermissionEditingGroupValuesAdministrators RangePermissionEditingGroupValues = "administrators"
	RangePermissionEditingGroupValuesContributors   RangePermissionEditingGroupValues = "contributors"
	RangePermissionEditingGroupValuesEditors        RangePermissionEditingGroupValues = "editors"
	RangePermissionEditingGroupValuesOwners         RangePermissionEditingGroupValues = "owners"
	RangePermissionEditingGroupValuesCurrent        RangePermissionEditingGroupValues = "current"
)

func (e RangePermissionEditingGroupValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *RangePermissionEditingGroupValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = RangePermissionEditingGroupValues(
		attr.Value,
	)
	return nil
}

// TextVerticalAlignmentValues
type TextVerticalAlignmentValues string

const (
	TextVerticalAlignmentValuesTop         TextVerticalAlignmentValues = "top"
	TextVerticalAlignmentValuesCenter      TextVerticalAlignmentValues = "center"
	TextVerticalAlignmentValuesBottom      TextVerticalAlignmentValues = "bottom"
	TextVerticalAlignmentValuesJustify     TextVerticalAlignmentValues = "justify"
	TextVerticalAlignmentValuesDistributed TextVerticalAlignmentValues = "distributed"
)

func (e TextVerticalAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextVerticalAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextVerticalAlignmentValues(attr.Value)
	return nil
}

// FormatActionValues
type FormatActionValues string

const (
	FormatActionValuesBlank      FormatActionValues = "blank"
	FormatActionValuesFormatting FormatActionValues = "formatting"
)

func (e FormatActionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FormatActionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FormatActionValues(attr.Value)
	return nil
}

// DataValidationValues
type DataValidationValues string

const (
	DataValidationValuesNone       DataValidationValues = "none"
	DataValidationValuesWhole      DataValidationValues = "whole"
	DataValidationValuesDecimal    DataValidationValues = "decimal"
	DataValidationValuesList       DataValidationValues = "list"
	DataValidationValuesDate       DataValidationValues = "date"
	DataValidationValuesTime       DataValidationValues = "time"
	DataValidationValuesTextlength DataValidationValues = "textLength"
	DataValidationValuesCustom     DataValidationValues = "custom"
)

func (e DataValidationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DataValidationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DataValidationValues(attr.Value)
	return nil
}

// CombineBracketValues
type CombineBracketValues string

const (
	CombineBracketValuesNone   CombineBracketValues = "none"
	CombineBracketValuesRound  CombineBracketValues = "round"
	CombineBracketValuesSquare CombineBracketValues = "square"
	CombineBracketValuesAngle  CombineBracketValues = "angle"
	CombineBracketValuesCurly  CombineBracketValues = "curly"
)

func (e CombineBracketValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CombineBracketValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CombineBracketValues(attr.Value)
	return nil
}

// TextWrappingValues
type TextWrappingValues string

const (
	TextWrappingValuesAuto      TextWrappingValues = "auto"
	TextWrappingValuesNotbeside TextWrappingValues = "notBeside"
	TextWrappingValuesAround    TextWrappingValues = "around"
	TextWrappingValuesTight     TextWrappingValues = "tight"
	TextWrappingValuesThrough   TextWrappingValues = "through"
	TextWrappingValuesNone      TextWrappingValues = "none"
)

func (e TextWrappingValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextWrappingValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextWrappingValues(attr.Value)
	return nil
}

// GroupByValues
type GroupByValues string

const (
	GroupByValuesRange    GroupByValues = "range"
	GroupByValuesSeconds  GroupByValues = "seconds"
	GroupByValuesMinutes  GroupByValues = "minutes"
	GroupByValuesHours    GroupByValues = "hours"
	GroupByValuesDays     GroupByValues = "days"
	GroupByValuesMonths   GroupByValues = "months"
	GroupByValuesQuarters GroupByValues = "quarters"
	GroupByValuesYears    GroupByValues = "years"
)

func (e GroupByValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *GroupByValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = GroupByValues(attr.Value)
	return nil
}

// TimePeriodValues
type TimePeriodValues string

const (
	TimePeriodValuesToday     TimePeriodValues = "today"
	TimePeriodValuesYesterday TimePeriodValues = "yesterday"
	TimePeriodValuesTomorrow  TimePeriodValues = "tomorrow"
	TimePeriodValuesLast7days TimePeriodValues = "last7Days"
	TimePeriodValuesThismonth TimePeriodValues = "thisMonth"
	TimePeriodValuesLastmonth TimePeriodValues = "lastMonth"
	TimePeriodValuesNextmonth TimePeriodValues = "nextMonth"
	TimePeriodValuesThisweek  TimePeriodValues = "thisWeek"
	TimePeriodValuesLastweek  TimePeriodValues = "lastWeek"
	TimePeriodValuesNextweek  TimePeriodValues = "nextWeek"
)

func (e TimePeriodValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TimePeriodValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TimePeriodValues(attr.Value)
	return nil
}

// StyleValues
type StyleValues string

const (
	StyleValuesParagraph StyleValues = "paragraph"
	StyleValuesCharacter StyleValues = "character"
	StyleValuesTable     StyleValues = "table"
	StyleValuesNumbering StyleValues = "numbering"
)

func (e StyleValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *StyleValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = StyleValues(attr.Value)
	return nil
}

// QuestionType
type QuestionType string

const (
	QuestionTypeCheckbox            QuestionType = "checkBox"
	QuestionTypeChoice              QuestionType = "choice"
	QuestionTypeDate                QuestionType = "date"
	QuestionTypeTime                QuestionType = "time"
	QuestionTypeMultiplelinesoftext QuestionType = "multipleLinesOfText"
	QuestionTypeNumber              QuestionType = "number"
	QuestionTypeSinglelineoftext    QuestionType = "singleLineOfText"
)

func (e QuestionType) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *QuestionType) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = QuestionType(attr.Value)
	return nil
}

// EmphasisMarkValues
type EmphasisMarkValues string

const (
	EmphasisMarkValuesNone     EmphasisMarkValues = "none"
	EmphasisMarkValuesDot      EmphasisMarkValues = "dot"
	EmphasisMarkValuesComma    EmphasisMarkValues = "comma"
	EmphasisMarkValuesCircle   EmphasisMarkValues = "circle"
	EmphasisMarkValuesUnderdot EmphasisMarkValues = "underDot"
)

func (e EmphasisMarkValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *EmphasisMarkValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = EmphasisMarkValues(attr.Value)
	return nil
}

// TableOverlapValues
type TableOverlapValues string

const (
	TableOverlapValuesNever   TableOverlapValues = "never"
	TableOverlapValuesOverlap TableOverlapValues = "overlap"
)

func (e TableOverlapValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TableOverlapValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TableOverlapValues(attr.Value)
	return nil
}

// FontFamilyValues
type FontFamilyValues string

const (
	FontFamilyValuesDecorative FontFamilyValues = "decorative"
	FontFamilyValuesModern     FontFamilyValues = "modern"
	FontFamilyValuesRoman      FontFamilyValues = "roman"
	FontFamilyValuesScript     FontFamilyValues = "script"
	FontFamilyValuesSwiss      FontFamilyValues = "swiss"
	FontFamilyValuesAuto       FontFamilyValues = "auto"
)

func (e FontFamilyValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FontFamilyValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FontFamilyValues(attr.Value)
	return nil
}

// TransitionSpeedValues
type TransitionSpeedValues string

const (
	TransitionSpeedValuesSlow TransitionSpeedValues = "slow"
	TransitionSpeedValuesMed  TransitionSpeedValues = "med"
	TransitionSpeedValuesFast TransitionSpeedValues = "fast"
)

func (e TransitionSpeedValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TransitionSpeedValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TransitionSpeedValues(attr.Value)
	return nil
}

// VerticalAlignmentRunValues
type VerticalAlignmentRunValues string

const (
	VerticalAlignmentRunValuesBaseline    VerticalAlignmentRunValues = "baseline"
	VerticalAlignmentRunValuesSuperscript VerticalAlignmentRunValues = "superscript"
	VerticalAlignmentRunValuesSubscript   VerticalAlignmentRunValues = "subscript"
)

func (e VerticalAlignmentRunValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *VerticalAlignmentRunValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = VerticalAlignmentRunValues(attr.Value)
	return nil
}

// CaptionPositionValues
type CaptionPositionValues string

const (
	CaptionPositionValuesAbove CaptionPositionValues = "above"
	CaptionPositionValuesBelow CaptionPositionValues = "below"
)

func (e CaptionPositionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CaptionPositionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CaptionPositionValues(attr.Value)
	return nil
}

// DisplayLocation
type DisplayLocation string

const (
	DisplayLocationMedia DisplayLocation = "media"
	DisplayLocationSlide DisplayLocation = "slide"
)

func (e DisplayLocation) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DisplayLocation) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DisplayLocation(attr.Value)
	return nil
}

// BehaviorOverrideValues
type BehaviorOverrideValues string

const (
	BehaviorOverrideValuesNormal     BehaviorOverrideValues = "normal"
	BehaviorOverrideValuesChildstyle BehaviorOverrideValues = "childStyle"
)

func (e BehaviorOverrideValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BehaviorOverrideValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BehaviorOverrideValues(attr.Value)
	return nil
}

// AnimateMotionPathEditModeValues
type AnimateMotionPathEditModeValues string

const (
	AnimateMotionPathEditModeValuesRelative AnimateMotionPathEditModeValues = "relative"
	AnimateMotionPathEditModeValuesFixed    AnimateMotionPathEditModeValues = "fixed"
)

func (e AnimateMotionPathEditModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AnimateMotionPathEditModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AnimateMotionPathEditModeValues(
		attr.Value,
	)
	return nil
}

// OlapSlicerCacheSortOrderValues
type OlapSlicerCacheSortOrderValues string

const (
	OlapSlicerCacheSortOrderValuesNatural    OlapSlicerCacheSortOrderValues = "natural"
	OlapSlicerCacheSortOrderValuesAscending  OlapSlicerCacheSortOrderValues = "ascending"
	OlapSlicerCacheSortOrderValuesDescending OlapSlicerCacheSortOrderValues = "descending"
)

func (e OlapSlicerCacheSortOrderValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *OlapSlicerCacheSortOrderValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = OlapSlicerCacheSortOrderValues(
		attr.Value,
	)
	return nil
}

// CalendarValues
type CalendarValues string

const (
	CalendarValuesGregorian            CalendarValues = "gregorian"
	CalendarValuesHijri                CalendarValues = "hijri"
	CalendarValuesUmalqura             CalendarValues = "umalqura"
	CalendarValuesHebrew               CalendarValues = "hebrew"
	CalendarValuesTaiwan               CalendarValues = "taiwan"
	CalendarValuesJapan                CalendarValues = "japan"
	CalendarValuesThai                 CalendarValues = "thai"
	CalendarValuesKorea                CalendarValues = "korea"
	CalendarValuesSaka                 CalendarValues = "saka"
	CalendarValuesGregorianxlitenglish CalendarValues = "gregorianXlitEnglish"
	CalendarValuesGregorianxlitfrench  CalendarValues = "gregorianXlitFrench"
	CalendarValuesGregorianus          CalendarValues = "gregorianUs"
	CalendarValuesGregorianmefrench    CalendarValues = "gregorianMeFrench"
	CalendarValuesGregorianarabic      CalendarValues = "gregorianArabic"
	CalendarValuesNone                 CalendarValues = "none"
)

func (e CalendarValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CalendarValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CalendarValues(attr.Value)
	return nil
}

// FontSchemeValues
type FontSchemeValues string

const (
	FontSchemeValuesNone  FontSchemeValues = "none"
	FontSchemeValuesMajor FontSchemeValues = "major"
	FontSchemeValuesMinor FontSchemeValues = "minor"
)

func (e FontSchemeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FontSchemeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FontSchemeValues(attr.Value)
	return nil
}

// MailMergeDataValues
type MailMergeDataValues string

const (
	MailMergeDataValuesTextfile    MailMergeDataValues = "textFile"
	MailMergeDataValuesDatabase    MailMergeDataValues = "database"
	MailMergeDataValuesSpreadsheet MailMergeDataValues = "spreadsheet"
	MailMergeDataValuesQuery       MailMergeDataValues = "query"
	MailMergeDataValuesOdbc        MailMergeDataValues = "odbc"
	MailMergeDataValuesNative      MailMergeDataValues = "native"
)

func (e MailMergeDataValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MailMergeDataValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MailMergeDataValues(attr.Value)
	return nil
}

// FootnoteEndnoteValues
type FootnoteEndnoteValues string

const (
	FootnoteEndnoteValuesNormal                FootnoteEndnoteValues = "normal"
	FootnoteEndnoteValuesSeparator             FootnoteEndnoteValues = "separator"
	FootnoteEndnoteValuesContinuationseparator FootnoteEndnoteValues = "continuationSeparator"
	FootnoteEndnoteValuesContinuationnotice    FootnoteEndnoteValues = "continuationNotice"
)

func (e FootnoteEndnoteValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FootnoteEndnoteValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FootnoteEndnoteValues(attr.Value)
	return nil
}

// PhotoAlbumLayoutValues
type PhotoAlbumLayoutValues string

const (
	PhotoAlbumLayoutValuesFittoslide                      PhotoAlbumLayoutValues = "fitToSlide"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues1Pic      PhotoAlbumLayoutValues = "1pic"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues2Pic      PhotoAlbumLayoutValues = "2pic"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues4Pic      PhotoAlbumLayoutValues = "4pic"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues1Pictitle PhotoAlbumLayoutValues = "1picTitle"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues2Pictitle PhotoAlbumLayoutValues = "2picTitle"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues4Pictitle PhotoAlbumLayoutValues = "4picTitle"
)

func (e PhotoAlbumLayoutValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PhotoAlbumLayoutValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PhotoAlbumLayoutValues(attr.Value)
	return nil
}

// OnOffOnlyValues
type OnOffOnlyValues string

const (
	OnOffOnlyValuesOn  OnOffOnlyValues = "on"
	OnOffOnlyValuesOff OnOffOnlyValues = "off"
)

func (e OnOffOnlyValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *OnOffOnlyValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = OnOffOnlyValues(attr.Value)
	return nil
}

// QualifierValues
type QualifierValues string

const (
	QualifierValuesDoublequote QualifierValues = "doubleQuote"
	QualifierValuesSinglequote QualifierValues = "singleQuote"
	QualifierValuesNone        QualifierValues = "none"
)

func (e QualifierValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *QualifierValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = QualifierValues(attr.Value)
	return nil
}

// SheetViewValues
type SheetViewValues string

const (
	SheetViewValuesNormal           SheetViewValues = "normal"
	SheetViewValuesPagebreakpreview SheetViewValues = "pageBreakPreview"
	SheetViewValuesPagelayout       SheetViewValues = "pageLayout"
)

func (e SheetViewValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SheetViewValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SheetViewValues(attr.Value)
	return nil
}

// CalculateModeValues
type CalculateModeValues string

const (
	CalculateModeValuesManual      CalculateModeValues = "manual"
	CalculateModeValuesAuto        CalculateModeValues = "auto"
	CalculateModeValuesAutonotable CalculateModeValues = "autoNoTable"
)

func (e CalculateModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CalculateModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CalculateModeValues(attr.Value)
	return nil
}

// PresetMaterialTypeValues
type PresetMaterialTypeValues string

const (
	PresetMaterialTypeValuesLegacymatte       PresetMaterialTypeValues = "legacyMatte"
	PresetMaterialTypeValuesLegacyplastic     PresetMaterialTypeValues = "legacyPlastic"
	PresetMaterialTypeValuesLegacymetal       PresetMaterialTypeValues = "legacyMetal"
	PresetMaterialTypeValuesLegacywireframe   PresetMaterialTypeValues = "legacyWireframe"
	PresetMaterialTypeValuesMatte             PresetMaterialTypeValues = "matte"
	PresetMaterialTypeValuesPlastic           PresetMaterialTypeValues = "plastic"
	PresetMaterialTypeValuesMetal             PresetMaterialTypeValues = "metal"
	PresetMaterialTypeValuesWarmmatte         PresetMaterialTypeValues = "warmMatte"
	PresetMaterialTypeValuesTranslucentpowder PresetMaterialTypeValues = "translucentPowder"
	PresetMaterialTypeValuesPowder            PresetMaterialTypeValues = "powder"
	PresetMaterialTypeValuesDkedge            PresetMaterialTypeValues = "dkEdge"
	PresetMaterialTypeValuesSoftedge          PresetMaterialTypeValues = "softEdge"
	PresetMaterialTypeValuesClear             PresetMaterialTypeValues = "clear"
	PresetMaterialTypeValuesFlat              PresetMaterialTypeValues = "flat"
	PresetMaterialTypeValuesSoftmetal         PresetMaterialTypeValues = "softmetal"
)

func (e PresetMaterialTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PresetMaterialTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PresetMaterialTypeValues(attr.Value)
	return nil
}

// CryptProviderValues
type CryptProviderValues string

const (
	CryptProviderValuesRsaaes  CryptProviderValues = "rsaAES"
	CryptProviderValuesRsafull CryptProviderValues = "rsaFull"
	CryptProviderValuesCustom  CryptProviderValues = "custom"
)

func (e CryptProviderValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CryptProviderValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CryptProviderValues(attr.Value)
	return nil
}

// HtmlPublishWebBrowserSupportValues
type HtmlPublishWebBrowserSupportValues string

const (
	HtmlPublishWebBrowserSupportValuesV4   HtmlPublishWebBrowserSupportValues = "v4"
	HtmlPublishWebBrowserSupportValuesV3   HtmlPublishWebBrowserSupportValues = "v3"
	HtmlPublishWebBrowserSupportValuesV3v4 HtmlPublishWebBrowserSupportValues = "v3v4"
)

func (e HtmlPublishWebBrowserSupportValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *HtmlPublishWebBrowserSupportValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = HtmlPublishWebBrowserSupportValues(
		attr.Value,
	)
	return nil
}

// WebColorValues
type WebColorValues string

const (
	WebColorValuesNone               WebColorValues = "none"
	WebColorValuesBrowser            WebColorValues = "browser"
	WebColorValuesPresentationtext   WebColorValues = "presentationText"
	WebColorValuesPresentationaccent WebColorValues = "presentationAccent"
	WebColorValuesWhitetextonblack   WebColorValues = "whiteTextOnBlack"
	WebColorValuesBlacktextonwhite   WebColorValues = "blackTextOnWhite"
)

func (e WebColorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *WebColorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = WebColorValues(attr.Value)
	return nil
}

// RevisionActionValues
type RevisionActionValues string

const (
	RevisionActionValuesAdd    RevisionActionValues = "add"
	RevisionActionValuesDelete RevisionActionValues = "delete"
)

func (e RevisionActionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *RevisionActionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = RevisionActionValues(attr.Value)
	return nil
}

// TextShapeValues
type TextShapeValues string

const (
	TextShapeValuesTextnoshape               TextShapeValues = "textNoShape"
	TextShapeValuesTextplain                 TextShapeValues = "textPlain"
	TextShapeValuesTextstop                  TextShapeValues = "textStop"
	TextShapeValuesTexttriangle              TextShapeValues = "textTriangle"
	TextShapeValuesTexttriangleinverted      TextShapeValues = "textTriangleInverted"
	TextShapeValuesTextchevron               TextShapeValues = "textChevron"
	TextShapeValuesTextchevroninverted       TextShapeValues = "textChevronInverted"
	TextShapeValuesTextringinside            TextShapeValues = "textRingInside"
	TextShapeValuesTextringoutside           TextShapeValues = "textRingOutside"
	TextShapeValuesTextarchup                TextShapeValues = "textArchUp"
	TextShapeValuesTextarchdown              TextShapeValues = "textArchDown"
	TextShapeValuesTextcircle                TextShapeValues = "textCircle"
	TextShapeValuesTextbutton                TextShapeValues = "textButton"
	TextShapeValuesTextarchuppour            TextShapeValues = "textArchUpPour"
	TextShapeValuesTextarchdownpour          TextShapeValues = "textArchDownPour"
	TextShapeValuesTextcirclepour            TextShapeValues = "textCirclePour"
	TextShapeValuesTextbuttonpour            TextShapeValues = "textButtonPour"
	TextShapeValuesTextcurveup               TextShapeValues = "textCurveUp"
	TextShapeValuesTextcurvedown             TextShapeValues = "textCurveDown"
	TextShapeValuesTextcanup                 TextShapeValues = "textCanUp"
	TextShapeValuesTextcandown               TextShapeValues = "textCanDown"
	TextShapeValuesTextwave1                 TextShapeValues = "textWave1"
	TextShapeValuesTextwave2                 TextShapeValues = "textWave2"
	TextShapeValuesTextdoublewave1           TextShapeValues = "textDoubleWave1"
	TextShapeValuesTextwave4                 TextShapeValues = "textWave4"
	TextShapeValuesTextinflate               TextShapeValues = "textInflate"
	TextShapeValuesTextdeflate               TextShapeValues = "textDeflate"
	TextShapeValuesTextinflatebottom         TextShapeValues = "textInflateBottom"
	TextShapeValuesTextdeflatebottom         TextShapeValues = "textDeflateBottom"
	TextShapeValuesTextinflatetop            TextShapeValues = "textInflateTop"
	TextShapeValuesTextdeflatetop            TextShapeValues = "textDeflateTop"
	TextShapeValuesTextdeflateinflate        TextShapeValues = "textDeflateInflate"
	TextShapeValuesTextdeflateinflatedeflate TextShapeValues = "textDeflateInflateDeflate"
	TextShapeValuesTextfaderight             TextShapeValues = "textFadeRight"
	TextShapeValuesTextfadeleft              TextShapeValues = "textFadeLeft"
	TextShapeValuesTextfadeup                TextShapeValues = "textFadeUp"
	TextShapeValuesTextfadedown              TextShapeValues = "textFadeDown"
	TextShapeValuesTextslantup               TextShapeValues = "textSlantUp"
	TextShapeValuesTextslantdown             TextShapeValues = "textSlantDown"
	TextShapeValuesTextcascadeup             TextShapeValues = "textCascadeUp"
	TextShapeValuesTextcascadedown           TextShapeValues = "textCascadeDown"
)

func (e TextShapeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextShapeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextShapeValues(attr.Value)
	return nil
}

// SlideSizeValues
type SlideSizeValues string

const (
	SlideSizeValuesScreen4x3           SlideSizeValues = "screen4x3"
	SlideSizeValuesLetter              SlideSizeValues = "letter"
	SlideSizeValuesA4                  SlideSizeValues = "A4"
	SlideSizeValuesSlideSizeValues35Mm SlideSizeValues = "35mm"
	SlideSizeValuesOverhead            SlideSizeValues = "overhead"
	SlideSizeValuesBanner              SlideSizeValues = "banner"
	SlideSizeValuesCustom              SlideSizeValues = "custom"
	SlideSizeValuesLedger              SlideSizeValues = "ledger"
	SlideSizeValuesA3                  SlideSizeValues = "A3"
	SlideSizeValuesB4iso               SlideSizeValues = "B4ISO"
	SlideSizeValuesB5iso               SlideSizeValues = "B5ISO"
	SlideSizeValuesB4jis               SlideSizeValues = "B4JIS"
	SlideSizeValuesB5jis               SlideSizeValues = "B5JIS"
	SlideSizeValuesHagakicard          SlideSizeValues = "hagakiCard"
	SlideSizeValuesScreen16x9          SlideSizeValues = "screen16x9"
	SlideSizeValuesScreen16x10         SlideSizeValues = "screen16x10"
)

func (e SlideSizeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SlideSizeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SlideSizeValues(attr.Value)
	return nil
}

// WebScreenSizeValues
type WebScreenSizeValues string

const (
	WebScreenSizeValuesWebScreenSizeValues544X376   WebScreenSizeValues = "544x376"
	WebScreenSizeValuesWebScreenSizeValues640X480   WebScreenSizeValues = "640x480"
	WebScreenSizeValuesWebScreenSizeValues720X512   WebScreenSizeValues = "720x512"
	WebScreenSizeValuesWebScreenSizeValues800X600   WebScreenSizeValues = "800x600"
	WebScreenSizeValuesWebScreenSizeValues1024X768  WebScreenSizeValues = "1024x768"
	WebScreenSizeValuesWebScreenSizeValues1152X882  WebScreenSizeValues = "1152x882"
	WebScreenSizeValuesWebScreenSizeValues1152X900  WebScreenSizeValues = "1152x900"
	WebScreenSizeValuesWebScreenSizeValues1280X1024 WebScreenSizeValues = "1280x1024"
	WebScreenSizeValuesWebScreenSizeValues1600X1200 WebScreenSizeValues = "1600x1200"
	WebScreenSizeValuesWebScreenSizeValues1800X1400 WebScreenSizeValues = "1800x1400"
	WebScreenSizeValuesWebScreenSizeValues1920X1200 WebScreenSizeValues = "1920x1200"
)

func (e WebScreenSizeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *WebScreenSizeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = WebScreenSizeValues(attr.Value)
	return nil
}

// HtmlFormattingValues
type HtmlFormattingValues string

const (
	HtmlFormattingValuesNone HtmlFormattingValues = "none"
	HtmlFormattingValuesRtf  HtmlFormattingValues = "rtf"
	HtmlFormattingValuesAll  HtmlFormattingValues = "all"
)

func (e HtmlFormattingValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *HtmlFormattingValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = HtmlFormattingValues(attr.Value)
	return nil
}

// ExternalConnectionValues
type ExternalConnectionValues string

const (
	ExternalConnectionValuesGeneral ExternalConnectionValues = "general"
	ExternalConnectionValuesText    ExternalConnectionValues = "text"
	ExternalConnectionValuesMdy     ExternalConnectionValues = "MDY"
	ExternalConnectionValuesDmy     ExternalConnectionValues = "DMY"
	ExternalConnectionValuesYmd     ExternalConnectionValues = "YMD"
	ExternalConnectionValuesMyd     ExternalConnectionValues = "MYD"
	ExternalConnectionValuesDym     ExternalConnectionValues = "DYM"
	ExternalConnectionValuesYdm     ExternalConnectionValues = "YDM"
	ExternalConnectionValuesSkip    ExternalConnectionValues = "skip"
	ExternalConnectionValuesEmd     ExternalConnectionValues = "EMD"
)

func (e ExternalConnectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ExternalConnectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ExternalConnectionValues(attr.Value)
	return nil
}

// HorizontalAlignmentValues
type HorizontalAlignmentValues string

const (
	HorizontalAlignmentValuesLeft    HorizontalAlignmentValues = "left"
	HorizontalAlignmentValuesCenter  HorizontalAlignmentValues = "center"
	HorizontalAlignmentValuesRight   HorizontalAlignmentValues = "right"
	HorizontalAlignmentValuesInside  HorizontalAlignmentValues = "inside"
	HorizontalAlignmentValuesOutside HorizontalAlignmentValues = "outside"
)

func (e HorizontalAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *HorizontalAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = HorizontalAlignmentValues(attr.Value)
	return nil
}

// FontTypeHintValues
type FontTypeHintValues string

const (
	FontTypeHintValuesDefault  FontTypeHintValues = "default"
	FontTypeHintValuesEastasia FontTypeHintValues = "eastAsia"
	FontTypeHintValuesCs       FontTypeHintValues = "cs"
)

func (e FontTypeHintValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FontTypeHintValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FontTypeHintValues(attr.Value)
	return nil
}

// LevelJustificationValues
type LevelJustificationValues string

const (
	LevelJustificationValuesLeft   LevelJustificationValues = "left"
	LevelJustificationValuesCenter LevelJustificationValues = "center"
	LevelJustificationValuesRight  LevelJustificationValues = "right"
)

func (e LevelJustificationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LevelJustificationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LevelJustificationValues(attr.Value)
	return nil
}

// TabularSlicerCacheSortOrderValues
type TabularSlicerCacheSortOrderValues string

const (
	TabularSlicerCacheSortOrderValuesAscending  TabularSlicerCacheSortOrderValues = "ascending"
	TabularSlicerCacheSortOrderValuesDescending TabularSlicerCacheSortOrderValues = "descending"
)

func (e TabularSlicerCacheSortOrderValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TabularSlicerCacheSortOrderValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TabularSlicerCacheSortOrderValues(
		attr.Value,
	)
	return nil
}

// SortValues
type SortValues string

const (
	SortValuesNone              SortValues = "none"
	SortValuesAscending         SortValues = "ascending"
	SortValuesDescending        SortValues = "descending"
	SortValuesAscendingalpha    SortValues = "ascendingAlpha"
	SortValuesDescendingalpha   SortValues = "descendingAlpha"
	SortValuesAscendingnatural  SortValues = "ascendingNatural"
	SortValuesDescendingnatural SortValues = "descendingNatural"
)

func (e SortValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SortValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SortValues(attr.Value)
	return nil
}

// DataViewAspectValues
type DataViewAspectValues string

const (
	DataViewAspectValuesDvaspectContent DataViewAspectValues = "DVASPECT_CONTENT"
	DataViewAspectValuesDvaspectIcon    DataViewAspectValues = "DVASPECT_ICON"
)

func (e DataViewAspectValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DataViewAspectValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DataViewAspectValues(attr.Value)
	return nil
}

// DocPartGalleryValues
type DocPartGalleryValues string

const (
	DocPartGalleryValuesPlaceholder       DocPartGalleryValues = "placeholder"
	DocPartGalleryValuesAny               DocPartGalleryValues = "any"
	DocPartGalleryValuesDefault           DocPartGalleryValues = "default"
	DocPartGalleryValuesDocparts          DocPartGalleryValues = "docParts"
	DocPartGalleryValuesCoverpg           DocPartGalleryValues = "coverPg"
	DocPartGalleryValuesEq                DocPartGalleryValues = "eq"
	DocPartGalleryValuesFtrs              DocPartGalleryValues = "ftrs"
	DocPartGalleryValuesHdrs              DocPartGalleryValues = "hdrs"
	DocPartGalleryValuesPgnum             DocPartGalleryValues = "pgNum"
	DocPartGalleryValuesTbls              DocPartGalleryValues = "tbls"
	DocPartGalleryValuesWatermarks        DocPartGalleryValues = "watermarks"
	DocPartGalleryValuesAutotxt           DocPartGalleryValues = "autoTxt"
	DocPartGalleryValuesTxtbox            DocPartGalleryValues = "txtBox"
	DocPartGalleryValuesPgnumt            DocPartGalleryValues = "pgNumT"
	DocPartGalleryValuesPgnumb            DocPartGalleryValues = "pgNumB"
	DocPartGalleryValuesPgnummargins      DocPartGalleryValues = "pgNumMargins"
	DocPartGalleryValuesTblofcontents     DocPartGalleryValues = "tblOfContents"
	DocPartGalleryValuesBib               DocPartGalleryValues = "bib"
	DocPartGalleryValuesCustquickparts    DocPartGalleryValues = "custQuickParts"
	DocPartGalleryValuesCustcoverpg       DocPartGalleryValues = "custCoverPg"
	DocPartGalleryValuesCusteq            DocPartGalleryValues = "custEq"
	DocPartGalleryValuesCustftrs          DocPartGalleryValues = "custFtrs"
	DocPartGalleryValuesCusthdrs          DocPartGalleryValues = "custHdrs"
	DocPartGalleryValuesCustpgnum         DocPartGalleryValues = "custPgNum"
	DocPartGalleryValuesCusttbls          DocPartGalleryValues = "custTbls"
	DocPartGalleryValuesCustwatermarks    DocPartGalleryValues = "custWatermarks"
	DocPartGalleryValuesCustautotxt       DocPartGalleryValues = "custAutoTxt"
	DocPartGalleryValuesCusttxtbox        DocPartGalleryValues = "custTxtBox"
	DocPartGalleryValuesCustpgnumt        DocPartGalleryValues = "custPgNumT"
	DocPartGalleryValuesCustpgnumb        DocPartGalleryValues = "custPgNumB"
	DocPartGalleryValuesCustpgnummargins  DocPartGalleryValues = "custPgNumMargins"
	DocPartGalleryValuesCusttblofcontents DocPartGalleryValues = "custTblOfContents"
	DocPartGalleryValuesCustbib           DocPartGalleryValues = "custBib"
	DocPartGalleryValuesCustom1           DocPartGalleryValues = "custom1"
	DocPartGalleryValuesCustom2           DocPartGalleryValues = "custom2"
	DocPartGalleryValuesCustom3           DocPartGalleryValues = "custom3"
	DocPartGalleryValuesCustom4           DocPartGalleryValues = "custom4"
	DocPartGalleryValuesCustom5           DocPartGalleryValues = "custom5"
)

func (e DocPartGalleryValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DocPartGalleryValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DocPartGalleryValues(attr.Value)
	return nil
}

// ObjectDrawAspect
type ObjectDrawAspect string

const (
	ObjectDrawAspectContent ObjectDrawAspect = "content"
	ObjectDrawAspectIcon    ObjectDrawAspect = "icon"
)

func (e ObjectDrawAspect) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ObjectDrawAspect) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ObjectDrawAspect(attr.Value)
	return nil
}

// AllocationMethodValues
type AllocationMethodValues string

const (
	AllocationMethodValuesEqualallocation    AllocationMethodValues = "equalAllocation"
	AllocationMethodValuesEqualincrement     AllocationMethodValues = "equalIncrement"
	AllocationMethodValuesWeightedallocation AllocationMethodValues = "weightedAllocation"
	AllocationMethodValuesWeightedincrement  AllocationMethodValues = "weightedIncrement"
)

func (e AllocationMethodValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AllocationMethodValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AllocationMethodValues(attr.Value)
	return nil
}

// QuestionFormat
type QuestionFormat string

const (
	QuestionFormatGeneraldate   QuestionFormat = "generalDate"
	QuestionFormatLongdate      QuestionFormat = "longDate"
	QuestionFormatShortdate     QuestionFormat = "shortDate"
	QuestionFormatLongtime      QuestionFormat = "longTime"
	QuestionFormatShorttime     QuestionFormat = "shortTime"
	QuestionFormatGeneralnumber QuestionFormat = "generalNumber"
	QuestionFormatStandard      QuestionFormat = "standard"
	QuestionFormatFixed         QuestionFormat = "fixed"
	QuestionFormatPercent       QuestionFormat = "percent"
	QuestionFormatCurrency      QuestionFormat = "currency"
)

func (e QuestionFormat) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *QuestionFormat) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = QuestionFormat(attr.Value)
	return nil
}

// SchemeColorValues
type SchemeColorValues string

const (
	SchemeColorValuesBg1      SchemeColorValues = "bg1"
	SchemeColorValuesTx1      SchemeColorValues = "tx1"
	SchemeColorValuesBg2      SchemeColorValues = "bg2"
	SchemeColorValuesTx2      SchemeColorValues = "tx2"
	SchemeColorValuesAccent1  SchemeColorValues = "accent1"
	SchemeColorValuesAccent2  SchemeColorValues = "accent2"
	SchemeColorValuesAccent3  SchemeColorValues = "accent3"
	SchemeColorValuesAccent4  SchemeColorValues = "accent4"
	SchemeColorValuesAccent5  SchemeColorValues = "accent5"
	SchemeColorValuesAccent6  SchemeColorValues = "accent6"
	SchemeColorValuesHlink    SchemeColorValues = "hlink"
	SchemeColorValuesFolhlink SchemeColorValues = "folHlink"
	SchemeColorValuesPhclr    SchemeColorValues = "phClr"
	SchemeColorValuesDk1      SchemeColorValues = "dk1"
	SchemeColorValuesLt1      SchemeColorValues = "lt1"
	SchemeColorValuesDk2      SchemeColorValues = "dk2"
	SchemeColorValuesLt2      SchemeColorValues = "lt2"
)

func (e SchemeColorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SchemeColorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SchemeColorValues(attr.Value)
	return nil
}

// CellValues
type CellValues string

const (
	CellValuesB         CellValues = "b"
	CellValuesN         CellValues = "n"
	CellValuesE         CellValues = "e"
	CellValuesS         CellValues = "s"
	CellValuesStr       CellValues = "str"
	CellValuesInlinestr CellValues = "inlineStr"
	CellValuesD         CellValues = "d"
)

func (e CellValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CellValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CellValues(attr.Value)
	return nil
}

// AutomaticColorValues
type AutomaticColorValues string

const (
	AutomaticColorValuesAuto AutomaticColorValues = "auto"
)

func (e AutomaticColorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AutomaticColorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AutomaticColorValues(attr.Value)
	return nil
}

// DocGridValues
type DocGridValues string

const (
	DocGridValuesDefault       DocGridValues = "default"
	DocGridValuesLines         DocGridValues = "lines"
	DocGridValuesLinesandchars DocGridValues = "linesAndChars"
	DocGridValuesSnaptochars   DocGridValues = "snapToChars"
)

func (e DocGridValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DocGridValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DocGridValues(attr.Value)
	return nil
}

// BreakValues
type BreakValues string

const (
	BreakValuesPage         BreakValues = "page"
	BreakValuesColumn       BreakValues = "column"
	BreakValuesTextwrapping BreakValues = "textWrapping"
)

func (e BreakValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BreakValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BreakValues(attr.Value)
	return nil
}

// AbsolutePositionTabPositioningBaseValues
type AbsolutePositionTabPositioningBaseValues string

const (
	AbsolutePositionTabPositioningBaseValuesMargin AbsolutePositionTabPositioningBaseValues = "margin"
	AbsolutePositionTabPositioningBaseValuesIndent AbsolutePositionTabPositioningBaseValues = "indent"
)

func (e AbsolutePositionTabPositioningBaseValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AbsolutePositionTabPositioningBaseValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AbsolutePositionTabPositioningBaseValues(
		attr.Value,
	)
	return nil
}

// UpdateLinksBehaviorValues
type UpdateLinksBehaviorValues string

const (
	UpdateLinksBehaviorValuesUserset UpdateLinksBehaviorValues = "userSet"
	UpdateLinksBehaviorValuesNever   UpdateLinksBehaviorValues = "never"
	UpdateLinksBehaviorValuesAlways  UpdateLinksBehaviorValues = "always"
)

func (e UpdateLinksBehaviorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *UpdateLinksBehaviorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = UpdateLinksBehaviorValues(attr.Value)
	return nil
}

// PivotShowAsValues
type PivotShowAsValues string

const (
	PivotShowAsValuesPercentofparent       PivotShowAsValues = "percentOfParent"
	PivotShowAsValuesPercentofparentrow    PivotShowAsValues = "percentOfParentRow"
	PivotShowAsValuesPercentofparentcol    PivotShowAsValues = "percentOfParentCol"
	PivotShowAsValuesPercentofrunningtotal PivotShowAsValues = "percentOfRunningTotal"
	PivotShowAsValuesRankascending         PivotShowAsValues = "rankAscending"
	PivotShowAsValuesRankdescending        PivotShowAsValues = "rankDescending"
)

func (e PivotShowAsValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PivotShowAsValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PivotShowAsValues(attr.Value)
	return nil
}

// TriggerRuntimeNodeValues
type TriggerRuntimeNodeValues string

const (
	TriggerRuntimeNodeValuesFirst TriggerRuntimeNodeValues = "first"
	TriggerRuntimeNodeValuesLast  TriggerRuntimeNodeValues = "last"
	TriggerRuntimeNodeValuesAll   TriggerRuntimeNodeValues = "all"
)

func (e TriggerRuntimeNodeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TriggerRuntimeNodeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TriggerRuntimeNodeValues(attr.Value)
	return nil
}

// FieldSortValues
type FieldSortValues string

const (
	FieldSortValuesManual     FieldSortValues = "manual"
	FieldSortValuesAscending  FieldSortValues = "ascending"
	FieldSortValuesDescending FieldSortValues = "descending"
)

func (e FieldSortValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FieldSortValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FieldSortValues(attr.Value)
	return nil
}

// FontPitchValues
type FontPitchValues string

const (
	FontPitchValuesFixed    FontPitchValues = "fixed"
	FontPitchValuesVariable FontPitchValues = "variable"
	FontPitchValuesDefault  FontPitchValues = "default"
)

func (e FontPitchValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FontPitchValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FontPitchValues(attr.Value)
	return nil
}

// StylePaneSortMethodsValues
type StylePaneSortMethodsValues string

const (
	StylePaneSortMethodsValuesStylePaneSortMethodsValues0000 StylePaneSortMethodsValues = "0000"
	StylePaneSortMethodsValuesName                           StylePaneSortMethodsValues = "name"
	StylePaneSortMethodsValuesStylePaneSortMethodsValues0001 StylePaneSortMethodsValues = "0001"
	StylePaneSortMethodsValuesPriority                       StylePaneSortMethodsValues = "priority"
	StylePaneSortMethodsValuesStylePaneSortMethodsValues0002 StylePaneSortMethodsValues = "0002"
	StylePaneSortMethodsValuesFont                           StylePaneSortMethodsValues = "font"
	StylePaneSortMethodsValuesStylePaneSortMethodsValues0003 StylePaneSortMethodsValues = "0003"
	StylePaneSortMethodsValuesBasedon                        StylePaneSortMethodsValues = "basedOn"
	StylePaneSortMethodsValuesStylePaneSortMethodsValues0004 StylePaneSortMethodsValues = "0004"
	StylePaneSortMethodsValuesType                           StylePaneSortMethodsValues = "type"
	StylePaneSortMethodsValuesStylePaneSortMethodsValues0005 StylePaneSortMethodsValues = "0005"
	StylePaneSortMethodsValuesDefault                        StylePaneSortMethodsValues = "default"
)

func (e StylePaneSortMethodsValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *StylePaneSortMethodsValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = StylePaneSortMethodsValues(attr.Value)
	return nil
}

// ObjectUpdateMode
type ObjectUpdateMode string

const (
	ObjectUpdateModeAlways ObjectUpdateMode = "always"
	ObjectUpdateModeOncall ObjectUpdateMode = "onCall"
)

func (e ObjectUpdateMode) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ObjectUpdateMode) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ObjectUpdateMode(attr.Value)
	return nil
}

// PresetColorValues
type PresetColorValues string

const (
	PresetColorValuesAliceblue            PresetColorValues = "aliceBlue"
	PresetColorValuesAntiquewhite         PresetColorValues = "antiqueWhite"
	PresetColorValuesAqua                 PresetColorValues = "aqua"
	PresetColorValuesAquamarine           PresetColorValues = "aquamarine"
	PresetColorValuesAzure                PresetColorValues = "azure"
	PresetColorValuesBeige                PresetColorValues = "beige"
	PresetColorValuesBisque               PresetColorValues = "bisque"
	PresetColorValuesBlack                PresetColorValues = "black"
	PresetColorValuesBlanchedalmond       PresetColorValues = "blanchedAlmond"
	PresetColorValuesBlue                 PresetColorValues = "blue"
	PresetColorValuesBlueviolet           PresetColorValues = "blueViolet"
	PresetColorValuesBrown                PresetColorValues = "brown"
	PresetColorValuesBurlywood            PresetColorValues = "burlyWood"
	PresetColorValuesCadetblue            PresetColorValues = "cadetBlue"
	PresetColorValuesChartreuse           PresetColorValues = "chartreuse"
	PresetColorValuesChocolate            PresetColorValues = "chocolate"
	PresetColorValuesCoral                PresetColorValues = "coral"
	PresetColorValuesCornflowerblue       PresetColorValues = "cornflowerBlue"
	PresetColorValuesCornsilk             PresetColorValues = "cornsilk"
	PresetColorValuesCrimson              PresetColorValues = "crimson"
	PresetColorValuesCyan                 PresetColorValues = "cyan"
	PresetColorValuesDkblue               PresetColorValues = "dkBlue"
	PresetColorValuesDkcyan               PresetColorValues = "dkCyan"
	PresetColorValuesDkgoldenrod          PresetColorValues = "dkGoldenrod"
	PresetColorValuesDkgray               PresetColorValues = "dkGray"
	PresetColorValuesDkgreen              PresetColorValues = "dkGreen"
	PresetColorValuesDkkhaki              PresetColorValues = "dkKhaki"
	PresetColorValuesDkmagenta            PresetColorValues = "dkMagenta"
	PresetColorValuesDkolivegreen         PresetColorValues = "dkOliveGreen"
	PresetColorValuesDkorange             PresetColorValues = "dkOrange"
	PresetColorValuesDkorchid             PresetColorValues = "dkOrchid"
	PresetColorValuesDkred                PresetColorValues = "dkRed"
	PresetColorValuesDksalmon             PresetColorValues = "dkSalmon"
	PresetColorValuesDkseagreen           PresetColorValues = "dkSeaGreen"
	PresetColorValuesDkslateblue          PresetColorValues = "dkSlateBlue"
	PresetColorValuesDkslategray          PresetColorValues = "dkSlateGray"
	PresetColorValuesDkturquoise          PresetColorValues = "dkTurquoise"
	PresetColorValuesDkviolet             PresetColorValues = "dkViolet"
	PresetColorValuesDeeppink             PresetColorValues = "deepPink"
	PresetColorValuesDeepskyblue          PresetColorValues = "deepSkyBlue"
	PresetColorValuesDimgray              PresetColorValues = "dimGray"
	PresetColorValuesDodgerblue           PresetColorValues = "dodgerBlue"
	PresetColorValuesFirebrick            PresetColorValues = "firebrick"
	PresetColorValuesFloralwhite          PresetColorValues = "floralWhite"
	PresetColorValuesForestgreen          PresetColorValues = "forestGreen"
	PresetColorValuesFuchsia              PresetColorValues = "fuchsia"
	PresetColorValuesGainsboro            PresetColorValues = "gainsboro"
	PresetColorValuesGhostwhite           PresetColorValues = "ghostWhite"
	PresetColorValuesGold                 PresetColorValues = "gold"
	PresetColorValuesGoldenrod            PresetColorValues = "goldenrod"
	PresetColorValuesGray                 PresetColorValues = "gray"
	PresetColorValuesGreen                PresetColorValues = "green"
	PresetColorValuesGreenyellow          PresetColorValues = "greenYellow"
	PresetColorValuesHoneydew             PresetColorValues = "honeydew"
	PresetColorValuesHotpink              PresetColorValues = "hotPink"
	PresetColorValuesIndianred            PresetColorValues = "indianRed"
	PresetColorValuesIndigo               PresetColorValues = "indigo"
	PresetColorValuesIvory                PresetColorValues = "ivory"
	PresetColorValuesKhaki                PresetColorValues = "khaki"
	PresetColorValuesLavender             PresetColorValues = "lavender"
	PresetColorValuesLavenderblush        PresetColorValues = "lavenderBlush"
	PresetColorValuesLawngreen            PresetColorValues = "lawnGreen"
	PresetColorValuesLemonchiffon         PresetColorValues = "lemonChiffon"
	PresetColorValuesLtblue               PresetColorValues = "ltBlue"
	PresetColorValuesLtcoral              PresetColorValues = "ltCoral"
	PresetColorValuesLtcyan               PresetColorValues = "ltCyan"
	PresetColorValuesLtgoldenrodyellow    PresetColorValues = "ltGoldenrodYellow"
	PresetColorValuesLtgray               PresetColorValues = "ltGray"
	PresetColorValuesLtgreen              PresetColorValues = "ltGreen"
	PresetColorValuesLtpink               PresetColorValues = "ltPink"
	PresetColorValuesLtsalmon             PresetColorValues = "ltSalmon"
	PresetColorValuesLtseagreen           PresetColorValues = "ltSeaGreen"
	PresetColorValuesLtskyblue            PresetColorValues = "ltSkyBlue"
	PresetColorValuesLtslategray          PresetColorValues = "ltSlateGray"
	PresetColorValuesLtsteelblue          PresetColorValues = "ltSteelBlue"
	PresetColorValuesLtyellow             PresetColorValues = "ltYellow"
	PresetColorValuesLime                 PresetColorValues = "lime"
	PresetColorValuesLimegreen            PresetColorValues = "limeGreen"
	PresetColorValuesLinen                PresetColorValues = "linen"
	PresetColorValuesMagenta              PresetColorValues = "magenta"
	PresetColorValuesMaroon               PresetColorValues = "maroon"
	PresetColorValuesMedaquamarine        PresetColorValues = "medAquamarine"
	PresetColorValuesMedblue              PresetColorValues = "medBlue"
	PresetColorValuesMedorchid            PresetColorValues = "medOrchid"
	PresetColorValuesMedpurple            PresetColorValues = "medPurple"
	PresetColorValuesMedseagreen          PresetColorValues = "medSeaGreen"
	PresetColorValuesMedslateblue         PresetColorValues = "medSlateBlue"
	PresetColorValuesMedspringgreen       PresetColorValues = "medSpringGreen"
	PresetColorValuesMedturquoise         PresetColorValues = "medTurquoise"
	PresetColorValuesMedvioletred         PresetColorValues = "medVioletRed"
	PresetColorValuesMidnightblue         PresetColorValues = "midnightBlue"
	PresetColorValuesMintcream            PresetColorValues = "mintCream"
	PresetColorValuesMistyrose            PresetColorValues = "mistyRose"
	PresetColorValuesMoccasin             PresetColorValues = "moccasin"
	PresetColorValuesNavajowhite          PresetColorValues = "navajoWhite"
	PresetColorValuesNavy                 PresetColorValues = "navy"
	PresetColorValuesOldlace              PresetColorValues = "oldLace"
	PresetColorValuesOlive                PresetColorValues = "olive"
	PresetColorValuesOlivedrab            PresetColorValues = "oliveDrab"
	PresetColorValuesOrange               PresetColorValues = "orange"
	PresetColorValuesOrangered            PresetColorValues = "orangeRed"
	PresetColorValuesOrchid               PresetColorValues = "orchid"
	PresetColorValuesPalegoldenrod        PresetColorValues = "paleGoldenrod"
	PresetColorValuesPalegreen            PresetColorValues = "paleGreen"
	PresetColorValuesPaleturquoise        PresetColorValues = "paleTurquoise"
	PresetColorValuesPalevioletred        PresetColorValues = "paleVioletRed"
	PresetColorValuesPapayawhip           PresetColorValues = "papayaWhip"
	PresetColorValuesPeachpuff            PresetColorValues = "peachPuff"
	PresetColorValuesPeru                 PresetColorValues = "peru"
	PresetColorValuesPink                 PresetColorValues = "pink"
	PresetColorValuesPlum                 PresetColorValues = "plum"
	PresetColorValuesPowderblue           PresetColorValues = "powderBlue"
	PresetColorValuesPurple               PresetColorValues = "purple"
	PresetColorValuesRed                  PresetColorValues = "red"
	PresetColorValuesRosybrown            PresetColorValues = "rosyBrown"
	PresetColorValuesRoyalblue            PresetColorValues = "royalBlue"
	PresetColorValuesSaddlebrown          PresetColorValues = "saddleBrown"
	PresetColorValuesSalmon               PresetColorValues = "salmon"
	PresetColorValuesSandybrown           PresetColorValues = "sandyBrown"
	PresetColorValuesSeagreen             PresetColorValues = "seaGreen"
	PresetColorValuesSeashell             PresetColorValues = "seaShell"
	PresetColorValuesSienna               PresetColorValues = "sienna"
	PresetColorValuesSilver               PresetColorValues = "silver"
	PresetColorValuesSkyblue              PresetColorValues = "skyBlue"
	PresetColorValuesSlateblue            PresetColorValues = "slateBlue"
	PresetColorValuesSlategray            PresetColorValues = "slateGray"
	PresetColorValuesSnow                 PresetColorValues = "snow"
	PresetColorValuesSpringgreen          PresetColorValues = "springGreen"
	PresetColorValuesSteelblue            PresetColorValues = "steelBlue"
	PresetColorValuesTan                  PresetColorValues = "tan"
	PresetColorValuesTeal                 PresetColorValues = "teal"
	PresetColorValuesThistle              PresetColorValues = "thistle"
	PresetColorValuesTomato               PresetColorValues = "tomato"
	PresetColorValuesTurquoise            PresetColorValues = "turquoise"
	PresetColorValuesViolet               PresetColorValues = "violet"
	PresetColorValuesWheat                PresetColorValues = "wheat"
	PresetColorValuesWhite                PresetColorValues = "white"
	PresetColorValuesWhitesmoke           PresetColorValues = "whiteSmoke"
	PresetColorValuesYellow               PresetColorValues = "yellow"
	PresetColorValuesYellowgreen          PresetColorValues = "yellowGreen"
	PresetColorValuesDarkblue             PresetColorValues = "darkBlue"
	PresetColorValuesDarkcyan             PresetColorValues = "darkCyan"
	PresetColorValuesDarkgoldenrod        PresetColorValues = "darkGoldenrod"
	PresetColorValuesDarkgray             PresetColorValues = "darkGray"
	PresetColorValuesDarkgrey             PresetColorValues = "darkGrey"
	PresetColorValuesDarkgreen            PresetColorValues = "darkGreen"
	PresetColorValuesDarkkhaki            PresetColorValues = "darkKhaki"
	PresetColorValuesDarkmagenta          PresetColorValues = "darkMagenta"
	PresetColorValuesDarkolivegreen       PresetColorValues = "darkOliveGreen"
	PresetColorValuesDarkorange           PresetColorValues = "darkOrange"
	PresetColorValuesDarkorchid           PresetColorValues = "darkOrchid"
	PresetColorValuesDarkred              PresetColorValues = "darkRed"
	PresetColorValuesDarksalmon           PresetColorValues = "darkSalmon"
	PresetColorValuesDarkseagreen         PresetColorValues = "darkSeaGreen"
	PresetColorValuesDarkslateblue        PresetColorValues = "darkSlateBlue"
	PresetColorValuesDarkslategray        PresetColorValues = "darkSlateGray"
	PresetColorValuesDarkslategrey        PresetColorValues = "darkSlateGrey"
	PresetColorValuesDarkturquoise        PresetColorValues = "darkTurquoise"
	PresetColorValuesDarkviolet           PresetColorValues = "darkViolet"
	PresetColorValuesLightblue            PresetColorValues = "lightBlue"
	PresetColorValuesLightcoral           PresetColorValues = "lightCoral"
	PresetColorValuesLightcyan            PresetColorValues = "lightCyan"
	PresetColorValuesLightgoldenrodyellow PresetColorValues = "lightGoldenrodYellow"
	PresetColorValuesLightgray            PresetColorValues = "lightGray"
	PresetColorValuesLightgrey            PresetColorValues = "lightGrey"
	PresetColorValuesLightgreen           PresetColorValues = "lightGreen"
	PresetColorValuesLightpink            PresetColorValues = "lightPink"
	PresetColorValuesLightsalmon          PresetColorValues = "lightSalmon"
	PresetColorValuesLightseagreen        PresetColorValues = "lightSeaGreen"
	PresetColorValuesLightskyblue         PresetColorValues = "lightSkyBlue"
	PresetColorValuesLightslategray       PresetColorValues = "lightSlateGray"
	PresetColorValuesLightslategrey       PresetColorValues = "lightSlateGrey"
	PresetColorValuesLightsteelblue       PresetColorValues = "lightSteelBlue"
	PresetColorValuesLightyellow          PresetColorValues = "lightYellow"
	PresetColorValuesMediumaquamarine     PresetColorValues = "mediumAquamarine"
	PresetColorValuesMediumblue           PresetColorValues = "mediumBlue"
	PresetColorValuesMediumorchid         PresetColorValues = "mediumOrchid"
	PresetColorValuesMediumpurple         PresetColorValues = "mediumPurple"
	PresetColorValuesMediumseagreen       PresetColorValues = "mediumSeaGreen"
	PresetColorValuesMediumslateblue      PresetColorValues = "mediumSlateBlue"
	PresetColorValuesMediumspringgreen    PresetColorValues = "mediumSpringGreen"
	PresetColorValuesMediumturquoise      PresetColorValues = "mediumTurquoise"
	PresetColorValuesMediumvioletred      PresetColorValues = "mediumVioletRed"
	PresetColorValuesDkgrey               PresetColorValues = "dkGrey"
	PresetColorValuesDimgrey              PresetColorValues = "dimGrey"
	PresetColorValuesDkslategrey          PresetColorValues = "dkSlateGrey"
	PresetColorValuesGrey                 PresetColorValues = "grey"
	PresetColorValuesLtgrey               PresetColorValues = "ltGrey"
	PresetColorValuesLtslategrey          PresetColorValues = "ltSlateGrey"
	PresetColorValuesSlategrey            PresetColorValues = "slateGrey"
)

func (e PresetColorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PresetColorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PresetColorValues(attr.Value)
	return nil
}

// PhoneticAlignmentValues
type PhoneticAlignmentValues string

const (
	PhoneticAlignmentValuesNocontrol   PhoneticAlignmentValues = "noControl"
	PhoneticAlignmentValuesLeft        PhoneticAlignmentValues = "left"
	PhoneticAlignmentValuesCenter      PhoneticAlignmentValues = "center"
	PhoneticAlignmentValuesDistributed PhoneticAlignmentValues = "distributed"
)

func (e PhoneticAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PhoneticAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PhoneticAlignmentValues(attr.Value)
	return nil
}

// HeaderFooterValues
type HeaderFooterValues string

const (
	HeaderFooterValuesEven    HeaderFooterValues = "even"
	HeaderFooterValuesDefault HeaderFooterValues = "default"
	HeaderFooterValuesFirst   HeaderFooterValues = "first"
)

func (e HeaderFooterValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *HeaderFooterValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = HeaderFooterValues(attr.Value)
	return nil
}

// PivotEditValueTypeValues
type PivotEditValueTypeValues string

const (
	PivotEditValueTypeValuesNumber   PivotEditValueTypeValues = "number"
	PivotEditValueTypeValuesDatetime PivotEditValueTypeValues = "dateTime"
	PivotEditValueTypeValuesString   PivotEditValueTypeValues = "string"
	PivotEditValueTypeValuesBoolean  PivotEditValueTypeValues = "boolean"
	PivotEditValueTypeValuesError    PivotEditValueTypeValues = "error"
)

func (e PivotEditValueTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PivotEditValueTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PivotEditValueTypeValues(attr.Value)
	return nil
}

// ThemeColorValues
type ThemeColorValues string

const (
	ThemeColorValuesDark1             ThemeColorValues = "dark1"
	ThemeColorValuesLight1            ThemeColorValues = "light1"
	ThemeColorValuesDark2             ThemeColorValues = "dark2"
	ThemeColorValuesLight2            ThemeColorValues = "light2"
	ThemeColorValuesAccent1           ThemeColorValues = "accent1"
	ThemeColorValuesAccent2           ThemeColorValues = "accent2"
	ThemeColorValuesAccent3           ThemeColorValues = "accent3"
	ThemeColorValuesAccent4           ThemeColorValues = "accent4"
	ThemeColorValuesAccent5           ThemeColorValues = "accent5"
	ThemeColorValuesAccent6           ThemeColorValues = "accent6"
	ThemeColorValuesHyperlink         ThemeColorValues = "hyperlink"
	ThemeColorValuesFollowedhyperlink ThemeColorValues = "followedHyperlink"
	ThemeColorValuesNone              ThemeColorValues = "none"
	ThemeColorValuesBackground1       ThemeColorValues = "background1"
	ThemeColorValuesText1             ThemeColorValues = "text1"
	ThemeColorValuesBackground2       ThemeColorValues = "background2"
	ThemeColorValuesText2             ThemeColorValues = "text2"
)

func (e ThemeColorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ThemeColorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ThemeColorValues(attr.Value)
	return nil
}

// MovingPeriodStep
type MovingPeriodStep string

const (
	MovingPeriodStepYear    MovingPeriodStep = "year"
	MovingPeriodStepQuarter MovingPeriodStep = "quarter"
	MovingPeriodStepMonth   MovingPeriodStep = "month"
	MovingPeriodStepWeek    MovingPeriodStep = "week"
	MovingPeriodStepDay     MovingPeriodStep = "day"
)

func (e MovingPeriodStep) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MovingPeriodStep) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MovingPeriodStep(attr.Value)
	return nil
}

// SystemColorValues
type SystemColorValues string

const (
	SystemColorValuesScrollbar                   SystemColorValues = "scrollBar"
	SystemColorValuesBackground                  SystemColorValues = "background"
	SystemColorValuesActivecaption               SystemColorValues = "activeCaption"
	SystemColorValuesInactivecaption             SystemColorValues = "inactiveCaption"
	SystemColorValuesMenu                        SystemColorValues = "menu"
	SystemColorValuesWindow                      SystemColorValues = "window"
	SystemColorValuesWindowframe                 SystemColorValues = "windowFrame"
	SystemColorValuesMenutext                    SystemColorValues = "menuText"
	SystemColorValuesWindowtext                  SystemColorValues = "windowText"
	SystemColorValuesCaptiontext                 SystemColorValues = "captionText"
	SystemColorValuesActiveborder                SystemColorValues = "activeBorder"
	SystemColorValuesInactiveborder              SystemColorValues = "inactiveBorder"
	SystemColorValuesAppworkspace                SystemColorValues = "appWorkspace"
	SystemColorValuesHighlight                   SystemColorValues = "highlight"
	SystemColorValuesHighlighttext               SystemColorValues = "highlightText"
	SystemColorValuesBtnface                     SystemColorValues = "btnFace"
	SystemColorValuesBtnshadow                   SystemColorValues = "btnShadow"
	SystemColorValuesGraytext                    SystemColorValues = "grayText"
	SystemColorValuesBtntext                     SystemColorValues = "btnText"
	SystemColorValuesInactivecaptiontext         SystemColorValues = "inactiveCaptionText"
	SystemColorValuesBtnhighlight                SystemColorValues = "btnHighlight"
	SystemColorValuesSystemColorValues3Ddkshadow SystemColorValues = "3dDkShadow"
	SystemColorValuesSystemColorValues3Dlight    SystemColorValues = "3dLight"
	SystemColorValuesInfotext                    SystemColorValues = "infoText"
	SystemColorValuesInfobk                      SystemColorValues = "infoBk"
	SystemColorValuesHotlight                    SystemColorValues = "hotLight"
	SystemColorValuesGradientactivecaption       SystemColorValues = "gradientActiveCaption"
	SystemColorValuesGradientinactivecaption     SystemColorValues = "gradientInactiveCaption"
	SystemColorValuesMenuhighlight               SystemColorValues = "menuHighlight"
	SystemColorValuesMenubar                     SystemColorValues = "menuBar"
)

func (e SystemColorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SystemColorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SystemColorValues(attr.Value)
	return nil
}

// MdxFunctionValues
type MdxFunctionValues string

const (
	MdxFunctionValuesM MdxFunctionValues = "m"
	MdxFunctionValuesV MdxFunctionValues = "v"
	MdxFunctionValuesS MdxFunctionValues = "s"
	MdxFunctionValuesC MdxFunctionValues = "c"
	MdxFunctionValuesR MdxFunctionValues = "r"
	MdxFunctionValuesP MdxFunctionValues = "p"
	MdxFunctionValuesK MdxFunctionValues = "k"
)

func (e MdxFunctionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MdxFunctionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MdxFunctionValues(attr.Value)
	return nil
}

// TotalsRowFunctionValues
type TotalsRowFunctionValues string

const (
	TotalsRowFunctionValuesNone      TotalsRowFunctionValues = "none"
	TotalsRowFunctionValuesSum       TotalsRowFunctionValues = "sum"
	TotalsRowFunctionValuesMin       TotalsRowFunctionValues = "min"
	TotalsRowFunctionValuesMax       TotalsRowFunctionValues = "max"
	TotalsRowFunctionValuesAverage   TotalsRowFunctionValues = "average"
	TotalsRowFunctionValuesCount     TotalsRowFunctionValues = "count"
	TotalsRowFunctionValuesCountnums TotalsRowFunctionValues = "countNums"
	TotalsRowFunctionValuesStddev    TotalsRowFunctionValues = "stdDev"
	TotalsRowFunctionValuesVar       TotalsRowFunctionValues = "var"
	TotalsRowFunctionValuesCustom    TotalsRowFunctionValues = "custom"
)

func (e TotalsRowFunctionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TotalsRowFunctionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TotalsRowFunctionValues(attr.Value)
	return nil
}

// MailMergeDestinationValues
type MailMergeDestinationValues string

const (
	MailMergeDestinationValuesNewdocument MailMergeDestinationValues = "newDocument"
	MailMergeDestinationValuesPrinter     MailMergeDestinationValues = "printer"
	MailMergeDestinationValuesEmail       MailMergeDestinationValues = "email"
	MailMergeDestinationValuesFax         MailMergeDestinationValues = "fax"
)

func (e MailMergeDestinationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MailMergeDestinationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MailMergeDestinationValues(attr.Value)
	return nil
}

// ObjectTypeValues
type ObjectTypeValues string

const (
	ObjectTypeValuesButton   ObjectTypeValues = "Button"
	ObjectTypeValuesCheckbox ObjectTypeValues = "CheckBox"
	ObjectTypeValuesDrop     ObjectTypeValues = "Drop"
	ObjectTypeValuesGbox     ObjectTypeValues = "GBox"
	ObjectTypeValuesLabel    ObjectTypeValues = "Label"
	ObjectTypeValuesList     ObjectTypeValues = "List"
	ObjectTypeValuesRadio    ObjectTypeValues = "Radio"
	ObjectTypeValuesScroll   ObjectTypeValues = "Scroll"
	ObjectTypeValuesSpin     ObjectTypeValues = "Spin"
	ObjectTypeValuesEditbox  ObjectTypeValues = "EditBox"
	ObjectTypeValuesDialog   ObjectTypeValues = "Dialog"
)

func (e ObjectTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ObjectTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ObjectTypeValues(attr.Value)
	return nil
}

// BooleanStyleValues
type BooleanStyleValues string

const (
	BooleanStyleValuesOn  BooleanStyleValues = "on"
	BooleanStyleValuesOff BooleanStyleValues = "off"
	BooleanStyleValuesDef BooleanStyleValues = "def"
)

func (e BooleanStyleValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BooleanStyleValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BooleanStyleValues(attr.Value)
	return nil
}

// TimeNodeSyncValues
type TimeNodeSyncValues string

const (
	TimeNodeSyncValuesNone    TimeNodeSyncValues = "none"
	TimeNodeSyncValuesCanslip TimeNodeSyncValues = "canSlip"
	TimeNodeSyncValuesLocked  TimeNodeSyncValues = "locked"
)

func (e TimeNodeSyncValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TimeNodeSyncValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TimeNodeSyncValues(attr.Value)
	return nil
}

// LockingValues
type LockingValues string

const (
	LockingValuesSdtlocked        LockingValues = "sdtLocked"
	LockingValuesContentlocked    LockingValues = "contentLocked"
	LockingValuesUnlocked         LockingValues = "unlocked"
	LockingValuesSdtcontentlocked LockingValues = "sdtContentLocked"
)

func (e LockingValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LockingValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LockingValues(attr.Value)
	return nil
}

// CryptAlgorithmValues
type CryptAlgorithmValues string

const (
	CryptAlgorithmValuesTypeany CryptAlgorithmValues = "typeAny"
	CryptAlgorithmValuesCustom  CryptAlgorithmValues = "custom"
)

func (e CryptAlgorithmValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CryptAlgorithmValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CryptAlgorithmValues(attr.Value)
	return nil
}

// SplitterBarStateValues
type SplitterBarStateValues string

const (
	SplitterBarStateValuesMinimized SplitterBarStateValues = "minimized"
	SplitterBarStateValuesRestored  SplitterBarStateValues = "restored"
	SplitterBarStateValuesMaximized SplitterBarStateValues = "maximized"
)

func (e SplitterBarStateValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SplitterBarStateValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SplitterBarStateValues(attr.Value)
	return nil
}

// CharacterSpacingValues
type CharacterSpacingValues string

const (
	CharacterSpacingValuesDonotcompress                      CharacterSpacingValues = "doNotCompress"
	CharacterSpacingValuesCompresspunctuation                CharacterSpacingValues = "compressPunctuation"
	CharacterSpacingValuesCompresspunctuationandjapanesekana CharacterSpacingValues = "compressPunctuationAndJapaneseKana"
)

func (e CharacterSpacingValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CharacterSpacingValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CharacterSpacingValues(attr.Value)
	return nil
}

// AnimationBuildValues
type AnimationBuildValues string

const (
	AnimationBuildValuesAllatonce AnimationBuildValues = "allAtOnce"
)

func (e AnimationBuildValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AnimationBuildValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AnimationBuildValues(attr.Value)
	return nil
}

// TextAnchoringTypeValues
type TextAnchoringTypeValues string

const (
	TextAnchoringTypeValuesT   TextAnchoringTypeValues = "t"
	TextAnchoringTypeValuesCtr TextAnchoringTypeValues = "ctr"
	TextAnchoringTypeValuesB   TextAnchoringTypeValues = "b"
)

func (e TextAnchoringTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextAnchoringTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextAnchoringTypeValues(attr.Value)
	return nil
}

// MdxSetOrderValues
type MdxSetOrderValues string

const (
	MdxSetOrderValuesU  MdxSetOrderValues = "u"
	MdxSetOrderValuesA  MdxSetOrderValues = "a"
	MdxSetOrderValuesD  MdxSetOrderValues = "d"
	MdxSetOrderValuesAa MdxSetOrderValues = "aa"
	MdxSetOrderValuesAd MdxSetOrderValues = "ad"
	MdxSetOrderValuesNa MdxSetOrderValues = "na"
	MdxSetOrderValuesNd MdxSetOrderValues = "nd"
)

func (e MdxSetOrderValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MdxSetOrderValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MdxSetOrderValues(attr.Value)
	return nil
}

// VerticalAnchorValues
type VerticalAnchorValues string

const (
	VerticalAnchorValuesText   VerticalAnchorValues = "text"
	VerticalAnchorValuesMargin VerticalAnchorValues = "margin"
	VerticalAnchorValuesPage   VerticalAnchorValues = "page"
)

func (e VerticalAnchorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *VerticalAnchorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = VerticalAnchorValues(attr.Value)
	return nil
}

// PageOrientationValues
type PageOrientationValues string

const (
	PageOrientationValuesPortrait  PageOrientationValues = "portrait"
	PageOrientationValuesLandscape PageOrientationValues = "landscape"
)

func (e PageOrientationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PageOrientationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PageOrientationValues(attr.Value)
	return nil
}

// TextCapsValues
type TextCapsValues string

const (
	TextCapsValuesNone  TextCapsValues = "none"
	TextCapsValuesSmall TextCapsValues = "small"
	TextCapsValuesAll   TextCapsValues = "all"
)

func (e TextCapsValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextCapsValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextCapsValues(attr.Value)
	return nil
}

// PlaceholderSizeValues
type PlaceholderSizeValues string

const (
	PlaceholderSizeValuesFull    PlaceholderSizeValues = "full"
	PlaceholderSizeValuesHalf    PlaceholderSizeValues = "half"
	PlaceholderSizeValuesQuarter PlaceholderSizeValues = "quarter"
)

func (e PlaceholderSizeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PlaceholderSizeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PlaceholderSizeValues(attr.Value)
	return nil
}

// SlideLayoutValues
type SlideLayoutValues string

const (
	SlideLayoutValuesTitle                   SlideLayoutValues = "title"
	SlideLayoutValuesTx                      SlideLayoutValues = "tx"
	SlideLayoutValuesTwocoltx                SlideLayoutValues = "twoColTx"
	SlideLayoutValuesTbl                     SlideLayoutValues = "tbl"
	SlideLayoutValuesTxandchart              SlideLayoutValues = "txAndChart"
	SlideLayoutValuesChartandtx              SlideLayoutValues = "chartAndTx"
	SlideLayoutValuesDgm                     SlideLayoutValues = "dgm"
	SlideLayoutValuesChart                   SlideLayoutValues = "chart"
	SlideLayoutValuesTxandclipart            SlideLayoutValues = "txAndClipArt"
	SlideLayoutValuesClipartandtx            SlideLayoutValues = "clipArtAndTx"
	SlideLayoutValuesTitleonly               SlideLayoutValues = "titleOnly"
	SlideLayoutValuesBlank                   SlideLayoutValues = "blank"
	SlideLayoutValuesTxandobj                SlideLayoutValues = "txAndObj"
	SlideLayoutValuesObjandtx                SlideLayoutValues = "objAndTx"
	SlideLayoutValuesObjonly                 SlideLayoutValues = "objOnly"
	SlideLayoutValuesObj                     SlideLayoutValues = "obj"
	SlideLayoutValuesTxandmedia              SlideLayoutValues = "txAndMedia"
	SlideLayoutValuesMediaandtx              SlideLayoutValues = "mediaAndTx"
	SlideLayoutValuesObjovertx               SlideLayoutValues = "objOverTx"
	SlideLayoutValuesTxoverobj               SlideLayoutValues = "txOverObj"
	SlideLayoutValuesTxandtwoobj             SlideLayoutValues = "txAndTwoObj"
	SlideLayoutValuesTwoobjandtx             SlideLayoutValues = "twoObjAndTx"
	SlideLayoutValuesTwoobjovertx            SlideLayoutValues = "twoObjOverTx"
	SlideLayoutValuesFourobj                 SlideLayoutValues = "fourObj"
	SlideLayoutValuesVerttx                  SlideLayoutValues = "vertTx"
	SlideLayoutValuesClipartandverttx        SlideLayoutValues = "clipArtAndVertTx"
	SlideLayoutValuesVerttitleandtx          SlideLayoutValues = "vertTitleAndTx"
	SlideLayoutValuesVerttitleandtxoverchart SlideLayoutValues = "vertTitleAndTxOverChart"
	SlideLayoutValuesTwoobj                  SlideLayoutValues = "twoObj"
	SlideLayoutValuesObjandtwoobj            SlideLayoutValues = "objAndTwoObj"
	SlideLayoutValuesTwoobjandobj            SlideLayoutValues = "twoObjAndObj"
	SlideLayoutValuesCust                    SlideLayoutValues = "cust"
	SlideLayoutValuesSechead                 SlideLayoutValues = "secHead"
	SlideLayoutValuesTwotxtwoobj             SlideLayoutValues = "twoTxTwoObj"
	SlideLayoutValuesObjtx                   SlideLayoutValues = "objTx"
	SlideLayoutValuesPictx                   SlideLayoutValues = "picTx"
)

func (e SlideLayoutValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SlideLayoutValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SlideLayoutValues(attr.Value)
	return nil
}

// DynamicFilterValues
type DynamicFilterValues string

const (
	DynamicFilterValuesNull         DynamicFilterValues = "null"
	DynamicFilterValuesAboveaverage DynamicFilterValues = "aboveAverage"
	DynamicFilterValuesBelowaverage DynamicFilterValues = "belowAverage"
	DynamicFilterValuesTomorrow     DynamicFilterValues = "tomorrow"
	DynamicFilterValuesToday        DynamicFilterValues = "today"
	DynamicFilterValuesYesterday    DynamicFilterValues = "yesterday"
	DynamicFilterValuesNextweek     DynamicFilterValues = "nextWeek"
	DynamicFilterValuesThisweek     DynamicFilterValues = "thisWeek"
	DynamicFilterValuesLastweek     DynamicFilterValues = "lastWeek"
	DynamicFilterValuesNextmonth    DynamicFilterValues = "nextMonth"
	DynamicFilterValuesThismonth    DynamicFilterValues = "thisMonth"
	DynamicFilterValuesLastmonth    DynamicFilterValues = "lastMonth"
	DynamicFilterValuesNextquarter  DynamicFilterValues = "nextQuarter"
	DynamicFilterValuesThisquarter  DynamicFilterValues = "thisQuarter"
	DynamicFilterValuesLastquarter  DynamicFilterValues = "lastQuarter"
	DynamicFilterValuesNextyear     DynamicFilterValues = "nextYear"
	DynamicFilterValuesThisyear     DynamicFilterValues = "thisYear"
	DynamicFilterValuesLastyear     DynamicFilterValues = "lastYear"
	DynamicFilterValuesYeartodate   DynamicFilterValues = "yearToDate"
	DynamicFilterValuesQ1           DynamicFilterValues = "Q1"
	DynamicFilterValuesQ2           DynamicFilterValues = "Q2"
	DynamicFilterValuesQ3           DynamicFilterValues = "Q3"
	DynamicFilterValuesQ4           DynamicFilterValues = "Q4"
	DynamicFilterValuesM1           DynamicFilterValues = "M1"
	DynamicFilterValuesM2           DynamicFilterValues = "M2"
	DynamicFilterValuesM3           DynamicFilterValues = "M3"
	DynamicFilterValuesM4           DynamicFilterValues = "M4"
	DynamicFilterValuesM5           DynamicFilterValues = "M5"
	DynamicFilterValuesM6           DynamicFilterValues = "M6"
	DynamicFilterValuesM7           DynamicFilterValues = "M7"
	DynamicFilterValuesM8           DynamicFilterValues = "M8"
	DynamicFilterValuesM9           DynamicFilterValues = "M9"
	DynamicFilterValuesM10          DynamicFilterValues = "M10"
	DynamicFilterValuesM11          DynamicFilterValues = "M11"
	DynamicFilterValuesM12          DynamicFilterValues = "M12"
)

func (e DynamicFilterValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DynamicFilterValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DynamicFilterValues(attr.Value)
	return nil
}

// PrintErrorValues
type PrintErrorValues string

const (
	PrintErrorValuesDisplayed PrintErrorValues = "displayed"
	PrintErrorValuesBlank     PrintErrorValues = "blank"
	PrintErrorValuesDash      PrintErrorValues = "dash"
	PrintErrorValuesNa        PrintErrorValues = "NA"
)

func (e PrintErrorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PrintErrorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PrintErrorValues(attr.Value)
	return nil
}

// InfoTextValues
type InfoTextValues string

const (
	InfoTextValuesText     InfoTextValues = "text"
	InfoTextValuesAutotext InfoTextValues = "autoText"
)

func (e InfoTextValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *InfoTextValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = InfoTextValues(attr.Value)
	return nil
}

// VerticalJustificationValues
type VerticalJustificationValues string

const (
	VerticalJustificationValuesTop    VerticalJustificationValues = "top"
	VerticalJustificationValuesCenter VerticalJustificationValues = "center"
	VerticalJustificationValuesBoth   VerticalJustificationValues = "both"
	VerticalJustificationValuesBottom VerticalJustificationValues = "bottom"
)

func (e VerticalJustificationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *VerticalJustificationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = VerticalJustificationValues(attr.Value)
	return nil
}

// BorderValues
type BorderValues string

const (
	BorderValuesNil                    BorderValues = "nil"
	BorderValuesNone                   BorderValues = "none"
	BorderValuesSingle                 BorderValues = "single"
	BorderValuesThick                  BorderValues = "thick"
	BorderValuesDouble                 BorderValues = "double"
	BorderValuesDotted                 BorderValues = "dotted"
	BorderValuesDashed                 BorderValues = "dashed"
	BorderValuesDotdash                BorderValues = "dotDash"
	BorderValuesDotdotdash             BorderValues = "dotDotDash"
	BorderValuesTriple                 BorderValues = "triple"
	BorderValuesThinthicksmallgap      BorderValues = "thinThickSmallGap"
	BorderValuesThickthinsmallgap      BorderValues = "thickThinSmallGap"
	BorderValuesThinthickthinsmallgap  BorderValues = "thinThickThinSmallGap"
	BorderValuesThinthickmediumgap     BorderValues = "thinThickMediumGap"
	BorderValuesThickthinmediumgap     BorderValues = "thickThinMediumGap"
	BorderValuesThinthickthinmediumgap BorderValues = "thinThickThinMediumGap"
	BorderValuesThinthicklargegap      BorderValues = "thinThickLargeGap"
	BorderValuesThickthinlargegap      BorderValues = "thickThinLargeGap"
	BorderValuesThinthickthinlargegap  BorderValues = "thinThickThinLargeGap"
	BorderValuesWave                   BorderValues = "wave"
	BorderValuesDoublewave             BorderValues = "doubleWave"
	BorderValuesDashsmallgap           BorderValues = "dashSmallGap"
	BorderValuesDashdotstroked         BorderValues = "dashDotStroked"
	BorderValuesThreedemboss           BorderValues = "threeDEmboss"
	BorderValuesThreedengrave          BorderValues = "threeDEngrave"
	BorderValuesOutset                 BorderValues = "outset"
	BorderValuesInset                  BorderValues = "inset"
	BorderValuesApples                 BorderValues = "apples"
	BorderValuesArchedscallops         BorderValues = "archedScallops"
	BorderValuesBabypacifier           BorderValues = "babyPacifier"
	BorderValuesBabyrattle             BorderValues = "babyRattle"
	BorderValuesBalloons3colors        BorderValues = "balloons3Colors"
	BorderValuesBalloonshotair         BorderValues = "balloonsHotAir"
	BorderValuesBasicblackdashes       BorderValues = "basicBlackDashes"
	BorderValuesBasicblackdots         BorderValues = "basicBlackDots"
	BorderValuesBasicblacksquares      BorderValues = "basicBlackSquares"
	BorderValuesBasicthinlines         BorderValues = "basicThinLines"
	BorderValuesBasicwhitedashes       BorderValues = "basicWhiteDashes"
	BorderValuesBasicwhitedots         BorderValues = "basicWhiteDots"
	BorderValuesBasicwhitesquares      BorderValues = "basicWhiteSquares"
	BorderValuesBasicwideinline        BorderValues = "basicWideInline"
	BorderValuesBasicwidemidline       BorderValues = "basicWideMidline"
	BorderValuesBasicwideoutline       BorderValues = "basicWideOutline"
	BorderValuesBats                   BorderValues = "bats"
	BorderValuesBirds                  BorderValues = "birds"
	BorderValuesBirdsflight            BorderValues = "birdsFlight"
	BorderValuesCabins                 BorderValues = "cabins"
	BorderValuesCakeslice              BorderValues = "cakeSlice"
	BorderValuesCandycorn              BorderValues = "candyCorn"
	BorderValuesCelticknotwork         BorderValues = "celticKnotwork"
	BorderValuesCertificatebanner      BorderValues = "certificateBanner"
	BorderValuesChainlink              BorderValues = "chainLink"
	BorderValuesChampagnebottle        BorderValues = "champagneBottle"
	BorderValuesCheckedbarblack        BorderValues = "checkedBarBlack"
	BorderValuesCheckedbarcolor        BorderValues = "checkedBarColor"
	BorderValuesCheckered              BorderValues = "checkered"
	BorderValuesChristmastree          BorderValues = "christmasTree"
	BorderValuesCircleslines           BorderValues = "circlesLines"
	BorderValuesCirclesrectangles      BorderValues = "circlesRectangles"
	BorderValuesClassicalwave          BorderValues = "classicalWave"
	BorderValuesClocks                 BorderValues = "clocks"
	BorderValuesCompass                BorderValues = "compass"
	BorderValuesConfetti               BorderValues = "confetti"
	BorderValuesConfettigrays          BorderValues = "confettiGrays"
	BorderValuesConfettioutline        BorderValues = "confettiOutline"
	BorderValuesConfettistreamers      BorderValues = "confettiStreamers"
	BorderValuesConfettiwhite          BorderValues = "confettiWhite"
	BorderValuesCornertriangles        BorderValues = "cornerTriangles"
	BorderValuesCouponcutoutdashes     BorderValues = "couponCutoutDashes"
	BorderValuesCouponcutoutdots       BorderValues = "couponCutoutDots"
	BorderValuesCrazymaze              BorderValues = "crazyMaze"
	BorderValuesCreaturesbutterfly     BorderValues = "creaturesButterfly"
	BorderValuesCreaturesfish          BorderValues = "creaturesFish"
	BorderValuesCreaturesinsects       BorderValues = "creaturesInsects"
	BorderValuesCreaturesladybug       BorderValues = "creaturesLadyBug"
	BorderValuesCrossstitch            BorderValues = "crossStitch"
	BorderValuesCup                    BorderValues = "cup"
	BorderValuesDecoarch               BorderValues = "decoArch"
	BorderValuesDecoarchcolor          BorderValues = "decoArchColor"
	BorderValuesDecoblocks             BorderValues = "decoBlocks"
	BorderValuesDiamondsgray           BorderValues = "diamondsGray"
	BorderValuesDoubled                BorderValues = "doubleD"
	BorderValuesDoublediamonds         BorderValues = "doubleDiamonds"
	BorderValuesEarth1                 BorderValues = "earth1"
	BorderValuesEarth2                 BorderValues = "earth2"
	BorderValuesEclipsingsquares1      BorderValues = "eclipsingSquares1"
	BorderValuesEclipsingsquares2      BorderValues = "eclipsingSquares2"
	BorderValuesEggsblack              BorderValues = "eggsBlack"
	BorderValuesFans                   BorderValues = "fans"
	BorderValuesFilm                   BorderValues = "film"
	BorderValuesFirecrackers           BorderValues = "firecrackers"
	BorderValuesFlowersblockprint      BorderValues = "flowersBlockPrint"
	BorderValuesFlowersdaisies         BorderValues = "flowersDaisies"
	BorderValuesFlowersmodern1         BorderValues = "flowersModern1"
	BorderValuesFlowersmodern2         BorderValues = "flowersModern2"
	BorderValuesFlowerspansy           BorderValues = "flowersPansy"
	BorderValuesFlowersredrose         BorderValues = "flowersRedRose"
	BorderValuesFlowersroses           BorderValues = "flowersRoses"
	BorderValuesFlowersteacup          BorderValues = "flowersTeacup"
	BorderValuesFlowerstiny            BorderValues = "flowersTiny"
	BorderValuesGems                   BorderValues = "gems"
	BorderValuesGingerbreadman         BorderValues = "gingerbreadMan"
	BorderValuesGradient               BorderValues = "gradient"
	BorderValuesHandmade1              BorderValues = "handmade1"
	BorderValuesHandmade2              BorderValues = "handmade2"
	BorderValuesHeartballoon           BorderValues = "heartBalloon"
	BorderValuesHeartgray              BorderValues = "heartGray"
	BorderValuesHearts                 BorderValues = "hearts"
	BorderValuesHeebiejeebies          BorderValues = "heebieJeebies"
	BorderValuesHolly                  BorderValues = "holly"
	BorderValuesHousefunky             BorderValues = "houseFunky"
	BorderValuesHypnotic               BorderValues = "hypnotic"
	BorderValuesIcecreamcones          BorderValues = "iceCreamCones"
	BorderValuesLightbulb              BorderValues = "lightBulb"
	BorderValuesLightning1             BorderValues = "lightning1"
	BorderValuesLightning2             BorderValues = "lightning2"
	BorderValuesMappins                BorderValues = "mapPins"
	BorderValuesMapleleaf              BorderValues = "mapleLeaf"
	BorderValuesMaplemuffins           BorderValues = "mapleMuffins"
	BorderValuesMarquee                BorderValues = "marquee"
	BorderValuesMarqueetoothed         BorderValues = "marqueeToothed"
	BorderValuesMoons                  BorderValues = "moons"
	BorderValuesMosaic                 BorderValues = "mosaic"
	BorderValuesMusicnotes             BorderValues = "musicNotes"
	BorderValuesNorthwest              BorderValues = "northwest"
	BorderValuesOvals                  BorderValues = "ovals"
	BorderValuesPackages               BorderValues = "packages"
	BorderValuesPalmsblack             BorderValues = "palmsBlack"
	BorderValuesPalmscolor             BorderValues = "palmsColor"
	BorderValuesPaperclips             BorderValues = "paperClips"
	BorderValuesPapyrus                BorderValues = "papyrus"
	BorderValuesPartyfavor             BorderValues = "partyFavor"
	BorderValuesPartyglass             BorderValues = "partyGlass"
	BorderValuesPencils                BorderValues = "pencils"
	BorderValuesPeople                 BorderValues = "people"
	BorderValuesPeoplewaving           BorderValues = "peopleWaving"
	BorderValuesPeoplehats             BorderValues = "peopleHats"
	BorderValuesPoinsettias            BorderValues = "poinsettias"
	BorderValuesPostagestamp           BorderValues = "postageStamp"
	BorderValuesPumpkin1               BorderValues = "pumpkin1"
	BorderValuesPushpinnote2           BorderValues = "pushPinNote2"
	BorderValuesPushpinnote1           BorderValues = "pushPinNote1"
	BorderValuesPyramids               BorderValues = "pyramids"
	BorderValuesPyramidsabove          BorderValues = "pyramidsAbove"
	BorderValuesQuadrants              BorderValues = "quadrants"
	BorderValuesRings                  BorderValues = "rings"
	BorderValuesSafari                 BorderValues = "safari"
	BorderValuesSawtooth               BorderValues = "sawtooth"
	BorderValuesSawtoothgray           BorderValues = "sawtoothGray"
	BorderValuesScaredcat              BorderValues = "scaredCat"
	BorderValuesSeattle                BorderValues = "seattle"
	BorderValuesShadowedsquares        BorderValues = "shadowedSquares"
	BorderValuesSharksteeth            BorderValues = "sharksTeeth"
	BorderValuesShorebirdtracks        BorderValues = "shorebirdTracks"
	BorderValuesSkyrocket              BorderValues = "skyrocket"
	BorderValuesSnowflakefancy         BorderValues = "snowflakeFancy"
	BorderValuesSnowflakes             BorderValues = "snowflakes"
	BorderValuesSombrero               BorderValues = "sombrero"
	BorderValuesSouthwest              BorderValues = "southwest"
	BorderValuesStars                  BorderValues = "stars"
	BorderValuesStarstop               BorderValues = "starsTop"
	BorderValuesStars3d                BorderValues = "stars3d"
	BorderValuesStarsblack             BorderValues = "starsBlack"
	BorderValuesStarsshadowed          BorderValues = "starsShadowed"
	BorderValuesSun                    BorderValues = "sun"
	BorderValuesSwirligig              BorderValues = "swirligig"
	BorderValuesTornpaper              BorderValues = "tornPaper"
	BorderValuesTornpaperblack         BorderValues = "tornPaperBlack"
	BorderValuesTrees                  BorderValues = "trees"
	BorderValuesTriangleparty          BorderValues = "triangleParty"
	BorderValuesTriangles              BorderValues = "triangles"
	BorderValuesTribal1                BorderValues = "tribal1"
	BorderValuesTribal2                BorderValues = "tribal2"
	BorderValuesTribal3                BorderValues = "tribal3"
	BorderValuesTribal4                BorderValues = "tribal4"
	BorderValuesTribal5                BorderValues = "tribal5"
	BorderValuesTribal6                BorderValues = "tribal6"
	BorderValuesTriangle1              BorderValues = "triangle1"
	BorderValuesTriangle2              BorderValues = "triangle2"
	BorderValuesTrianglecircle1        BorderValues = "triangleCircle1"
	BorderValuesTrianglecircle2        BorderValues = "triangleCircle2"
	BorderValuesShapes1                BorderValues = "shapes1"
	BorderValuesShapes2                BorderValues = "shapes2"
	BorderValuesTwistedlines1          BorderValues = "twistedLines1"
	BorderValuesTwistedlines2          BorderValues = "twistedLines2"
	BorderValuesVine                   BorderValues = "vine"
	BorderValuesWaveline               BorderValues = "waveline"
	BorderValuesWeavingangles          BorderValues = "weavingAngles"
	BorderValuesWeavingbraid           BorderValues = "weavingBraid"
	BorderValuesWeavingribbon          BorderValues = "weavingRibbon"
	BorderValuesWeavingstrips          BorderValues = "weavingStrips"
	BorderValuesWhiteflowers           BorderValues = "whiteFlowers"
	BorderValuesWoodwork               BorderValues = "woodwork"
	BorderValuesXillusions             BorderValues = "xIllusions"
	BorderValuesZanytriangles          BorderValues = "zanyTriangles"
	BorderValuesZigzag                 BorderValues = "zigZag"
	BorderValuesZigzagstitch           BorderValues = "zigZagStitch"
)

func (e BorderValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BorderValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BorderValues(attr.Value)
	return nil
}

// LineEndWidthValues
type LineEndWidthValues string

const (
	LineEndWidthValuesSm  LineEndWidthValues = "sm"
	LineEndWidthValuesMed LineEndWidthValues = "med"
	LineEndWidthValuesLg  LineEndWidthValues = "lg"
)

func (e LineEndWidthValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LineEndWidthValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LineEndWidthValues(attr.Value)
	return nil
}

// LineCapValues
type LineCapValues string

const (
	LineCapValuesRnd  LineCapValues = "rnd"
	LineCapValuesSq   LineCapValues = "sq"
	LineCapValuesFlat LineCapValues = "flat"
)

func (e LineCapValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LineCapValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LineCapValues(attr.Value)
	return nil
}

// XmlDataValues
type XmlDataValues string

const (
	XmlDataValuesString             XmlDataValues = "string"
	XmlDataValuesNormalizedstring   XmlDataValues = "normalizedString"
	XmlDataValuesToken              XmlDataValues = "token"
	XmlDataValuesByte               XmlDataValues = "byte"
	XmlDataValuesUnsignedbyte       XmlDataValues = "unsignedByte"
	XmlDataValuesBase64binary       XmlDataValues = "base64Binary"
	XmlDataValuesHexbinary          XmlDataValues = "hexBinary"
	XmlDataValuesInteger            XmlDataValues = "integer"
	XmlDataValuesPositiveinteger    XmlDataValues = "positiveInteger"
	XmlDataValuesNegativeinteger    XmlDataValues = "negativeInteger"
	XmlDataValuesNonpositiveinteger XmlDataValues = "nonPositiveInteger"
	XmlDataValuesNonnegativeinteger XmlDataValues = "nonNegativeInteger"
	XmlDataValuesInt                XmlDataValues = "int"
	XmlDataValuesUnsignedint        XmlDataValues = "unsignedInt"
	XmlDataValuesLong               XmlDataValues = "long"
	XmlDataValuesUnsignedlong       XmlDataValues = "unsignedLong"
	XmlDataValuesShort              XmlDataValues = "short"
	XmlDataValuesUnsignedshort      XmlDataValues = "unsignedShort"
	XmlDataValuesDecimal            XmlDataValues = "decimal"
	XmlDataValuesFloat              XmlDataValues = "float"
	XmlDataValuesDouble             XmlDataValues = "double"
	XmlDataValuesBoolean            XmlDataValues = "boolean"
	XmlDataValuesTime               XmlDataValues = "time"
	XmlDataValuesDatetime           XmlDataValues = "dateTime"
	XmlDataValuesDuration           XmlDataValues = "duration"
	XmlDataValuesDate               XmlDataValues = "date"
	XmlDataValuesGmonth             XmlDataValues = "gMonth"
	XmlDataValuesGyear              XmlDataValues = "gYear"
	XmlDataValuesGyearmonth         XmlDataValues = "gYearMonth"
	XmlDataValuesGday               XmlDataValues = "gDay"
	XmlDataValuesGmonthday          XmlDataValues = "gMonthDay"
	XmlDataValuesName               XmlDataValues = "Name"
	XmlDataValuesQname              XmlDataValues = "QName"
	XmlDataValuesNcname             XmlDataValues = "NCName"
	XmlDataValuesAnyuri             XmlDataValues = "anyURI"
	XmlDataValuesLanguage           XmlDataValues = "language"
	XmlDataValuesId                 XmlDataValues = "ID"
	XmlDataValuesIdref              XmlDataValues = "IDREF"
	XmlDataValuesIdrefs             XmlDataValues = "IDREFS"
	XmlDataValuesEntity             XmlDataValues = "ENTITY"
	XmlDataValuesEntities           XmlDataValues = "ENTITIES"
	XmlDataValuesNotation           XmlDataValues = "NOTATION"
	XmlDataValuesNmtoken            XmlDataValues = "NMTOKEN"
	XmlDataValuesNmtokens           XmlDataValues = "NMTOKENS"
	XmlDataValuesAnytype            XmlDataValues = "anyType"
)

func (e XmlDataValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *XmlDataValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = XmlDataValues(attr.Value)
	return nil
}

// MailMergeDocumentValues
type MailMergeDocumentValues string

const (
	MailMergeDocumentValuesCatalog       MailMergeDocumentValues = "catalog"
	MailMergeDocumentValuesEnvelopes     MailMergeDocumentValues = "envelopes"
	MailMergeDocumentValuesMailinglabels MailMergeDocumentValues = "mailingLabels"
	MailMergeDocumentValuesFormletters   MailMergeDocumentValues = "formLetters"
	MailMergeDocumentValuesEmail         MailMergeDocumentValues = "email"
	MailMergeDocumentValuesFax           MailMergeDocumentValues = "fax"
)

func (e MailMergeDocumentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MailMergeDocumentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MailMergeDocumentValues(attr.Value)
	return nil
}

// AbsolutePositionTabLeaderCharValues
type AbsolutePositionTabLeaderCharValues string

const (
	AbsolutePositionTabLeaderCharValuesNone       AbsolutePositionTabLeaderCharValues = "none"
	AbsolutePositionTabLeaderCharValuesDot        AbsolutePositionTabLeaderCharValues = "dot"
	AbsolutePositionTabLeaderCharValuesHyphen     AbsolutePositionTabLeaderCharValues = "hyphen"
	AbsolutePositionTabLeaderCharValuesUnderscore AbsolutePositionTabLeaderCharValues = "underscore"
	AbsolutePositionTabLeaderCharValuesMiddledot  AbsolutePositionTabLeaderCharValues = "middleDot"
)

func (e AbsolutePositionTabLeaderCharValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AbsolutePositionTabLeaderCharValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AbsolutePositionTabLeaderCharValues(
		attr.Value,
	)
	return nil
}

// RubyAlignValues
type RubyAlignValues string

const (
	RubyAlignValuesCenter           RubyAlignValues = "center"
	RubyAlignValuesDistributeletter RubyAlignValues = "distributeLetter"
	RubyAlignValuesDistributespace  RubyAlignValues = "distributeSpace"
	RubyAlignValuesLeft             RubyAlignValues = "left"
	RubyAlignValuesRight            RubyAlignValues = "right"
	RubyAlignValuesRightvertical    RubyAlignValues = "rightVertical"
)

func (e RubyAlignValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *RubyAlignValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = RubyAlignValues(attr.Value)
	return nil
}

// IndefiniteTimeDeclarationValues
type IndefiniteTimeDeclarationValues string

const (
	IndefiniteTimeDeclarationValuesIndefinite IndefiniteTimeDeclarationValues = "indefinite"
)

func (e IndefiniteTimeDeclarationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *IndefiniteTimeDeclarationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = IndefiniteTimeDeclarationValues(
		attr.Value,
	)
	return nil
}

// TimeNodeFillValues
type TimeNodeFillValues string

const (
	TimeNodeFillValuesRemove     TimeNodeFillValues = "remove"
	TimeNodeFillValuesFreeze     TimeNodeFillValues = "freeze"
	TimeNodeFillValuesHold       TimeNodeFillValues = "hold"
	TimeNodeFillValuesTransition TimeNodeFillValues = "transition"
)

func (e TimeNodeFillValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TimeNodeFillValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TimeNodeFillValues(attr.Value)
	return nil
}

// FilterOperatorValues
type FilterOperatorValues string

const (
	FilterOperatorValuesEqual              FilterOperatorValues = "equal"
	FilterOperatorValuesLessthan           FilterOperatorValues = "lessThan"
	FilterOperatorValuesLessthanorequal    FilterOperatorValues = "lessThanOrEqual"
	FilterOperatorValuesNotequal           FilterOperatorValues = "notEqual"
	FilterOperatorValuesGreaterthanorequal FilterOperatorValues = "greaterThanOrEqual"
	FilterOperatorValuesGreaterthan        FilterOperatorValues = "greaterThan"
)

func (e FilterOperatorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FilterOperatorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FilterOperatorValues(attr.Value)
	return nil
}

// VolatileValues
type VolatileValues string

const (
	VolatileValuesB VolatileValues = "b"
	VolatileValuesN VolatileValues = "n"
	VolatileValuesE VolatileValues = "e"
	VolatileValuesS VolatileValues = "s"
)

func (e VolatileValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *VolatileValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = VolatileValues(attr.Value)
	return nil
}

// HorizontalAnchorValues
type HorizontalAnchorValues string

const (
	HorizontalAnchorValuesText   HorizontalAnchorValues = "text"
	HorizontalAnchorValuesMargin HorizontalAnchorValues = "margin"
	HorizontalAnchorValuesPage   HorizontalAnchorValues = "page"
)

func (e HorizontalAnchorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *HorizontalAnchorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = HorizontalAnchorValues(attr.Value)
	return nil
}

// CompatSettingNameValues
type CompatSettingNameValues string

const (
	CompatSettingNameValuesCompatibilitymode                          CompatSettingNameValues = "compatibilityMode"
	CompatSettingNameValuesOverridetablestylefontsizeandjustification CompatSettingNameValues = "overrideTableStyleFontSizeAndJustification"
	CompatSettingNameValuesEnableopentypefeatures                     CompatSettingNameValues = "enableOpenTypeFeatures"
	CompatSettingNameValuesDonotflipmirrorindents                     CompatSettingNameValues = "doNotFlipMirrorIndents"
	CompatSettingNameValuesDifferentiatemultirowtableheaders          CompatSettingNameValues = "differentiateMultirowTableHeaders"
	CompatSettingNameValuesUseword2013trackbottomhyphenation          CompatSettingNameValues = "useWord2013TrackBottomHyphenation"
	CompatSettingNameValuesAllowhyphenationattrackbottom              CompatSettingNameValues = "allowHyphenationAtTrackBottom"
	CompatSettingNameValuesAllowtextafterfloatingtablebreak           CompatSettingNameValues = "allowTextAfterFloatingTableBreak"
)

func (e CompatSettingNameValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CompatSettingNameValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CompatSettingNameValues(attr.Value)
	return nil
}

// CredentialsMethodValues
type CredentialsMethodValues string

const (
	CredentialsMethodValuesIntegrated CredentialsMethodValues = "integrated"
	CredentialsMethodValuesNone       CredentialsMethodValues = "none"
	CredentialsMethodValuesStored     CredentialsMethodValues = "stored"
)

func (e CredentialsMethodValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CredentialsMethodValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CredentialsMethodValues(attr.Value)
	return nil
}

// HighlightColorValues
type HighlightColorValues string

const (
	HighlightColorValuesBlack       HighlightColorValues = "black"
	HighlightColorValuesBlue        HighlightColorValues = "blue"
	HighlightColorValuesCyan        HighlightColorValues = "cyan"
	HighlightColorValuesGreen       HighlightColorValues = "green"
	HighlightColorValuesMagenta     HighlightColorValues = "magenta"
	HighlightColorValuesRed         HighlightColorValues = "red"
	HighlightColorValuesYellow      HighlightColorValues = "yellow"
	HighlightColorValuesWhite       HighlightColorValues = "white"
	HighlightColorValuesDarkblue    HighlightColorValues = "darkBlue"
	HighlightColorValuesDarkcyan    HighlightColorValues = "darkCyan"
	HighlightColorValuesDarkgreen   HighlightColorValues = "darkGreen"
	HighlightColorValuesDarkmagenta HighlightColorValues = "darkMagenta"
	HighlightColorValuesDarkred     HighlightColorValues = "darkRed"
	HighlightColorValuesDarkyellow  HighlightColorValues = "darkYellow"
	HighlightColorValuesDarkgray    HighlightColorValues = "darkGray"
	HighlightColorValuesLightgray   HighlightColorValues = "lightGray"
	HighlightColorValuesNone        HighlightColorValues = "none"
)

func (e HighlightColorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *HighlightColorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = HighlightColorValues(attr.Value)
	return nil
}

// RectangleAlignmentValues
type RectangleAlignmentValues string

const (
	RectangleAlignmentValuesTl  RectangleAlignmentValues = "tl"
	RectangleAlignmentValuesT   RectangleAlignmentValues = "t"
	RectangleAlignmentValuesTr  RectangleAlignmentValues = "tr"
	RectangleAlignmentValuesL   RectangleAlignmentValues = "l"
	RectangleAlignmentValuesCtr RectangleAlignmentValues = "ctr"
	RectangleAlignmentValuesR   RectangleAlignmentValues = "r"
	RectangleAlignmentValuesBl  RectangleAlignmentValues = "bl"
	RectangleAlignmentValuesB   RectangleAlignmentValues = "b"
	RectangleAlignmentValuesBr  RectangleAlignmentValues = "br"
)

func (e RectangleAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *RectangleAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = RectangleAlignmentValues(attr.Value)
	return nil
}

// TileFlipValues
type TileFlipValues string

const (
	TileFlipValuesNone TileFlipValues = "none"
	TileFlipValuesX    TileFlipValues = "x"
	TileFlipValuesY    TileFlipValues = "y"
	TileFlipValuesXy   TileFlipValues = "xy"
)

func (e TileFlipValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TileFlipValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TileFlipValues(attr.Value)
	return nil
}

// TableRowAlignmentValues
type TableRowAlignmentValues string

const (
	TableRowAlignmentValuesLeft   TableRowAlignmentValues = "left"
	TableRowAlignmentValuesCenter TableRowAlignmentValues = "center"
	TableRowAlignmentValuesRight  TableRowAlignmentValues = "right"
)

func (e TableRowAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TableRowAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TableRowAlignmentValues(attr.Value)
	return nil
}

// PresetZoomValues
type PresetZoomValues string

const (
	PresetZoomValuesNone     PresetZoomValues = "none"
	PresetZoomValuesFullpage PresetZoomValues = "fullPage"
	PresetZoomValuesBestfit  PresetZoomValues = "bestFit"
	PresetZoomValuesTextfit  PresetZoomValues = "textFit"
)

func (e PresetZoomValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PresetZoomValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PresetZoomValues(attr.Value)
	return nil
}

// JustificationValues
type JustificationValues string

const (
	JustificationValuesLeft           JustificationValues = "left"
	JustificationValuesStart          JustificationValues = "start"
	JustificationValuesCenter         JustificationValues = "center"
	JustificationValuesRight          JustificationValues = "right"
	JustificationValuesEnd            JustificationValues = "end"
	JustificationValuesBoth           JustificationValues = "both"
	JustificationValuesMediumkashida  JustificationValues = "mediumKashida"
	JustificationValuesDistribute     JustificationValues = "distribute"
	JustificationValuesNumtab         JustificationValues = "numTab"
	JustificationValuesHighkashida    JustificationValues = "highKashida"
	JustificationValuesLowkashida     JustificationValues = "lowKashida"
	JustificationValuesThaidistribute JustificationValues = "thaiDistribute"
)

func (e JustificationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *JustificationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = JustificationValues(attr.Value)
	return nil
}

// LightRigDirectionValues
type LightRigDirectionValues string

const (
	LightRigDirectionValuesTl LightRigDirectionValues = "tl"
	LightRigDirectionValuesT  LightRigDirectionValues = "t"
	LightRigDirectionValuesTr LightRigDirectionValues = "tr"
	LightRigDirectionValuesL  LightRigDirectionValues = "l"
	LightRigDirectionValuesR  LightRigDirectionValues = "r"
	LightRigDirectionValuesBl LightRigDirectionValues = "bl"
	LightRigDirectionValuesB  LightRigDirectionValues = "b"
	LightRigDirectionValuesBr LightRigDirectionValues = "br"
)

func (e LightRigDirectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LightRigDirectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LightRigDirectionValues(attr.Value)
	return nil
}

// CryptAlgorithmClassValues
type CryptAlgorithmClassValues string

const (
	CryptAlgorithmClassValuesHash   CryptAlgorithmClassValues = "hash"
	CryptAlgorithmClassValuesCustom CryptAlgorithmClassValues = "custom"
)

func (e CryptAlgorithmClassValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CryptAlgorithmClassValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CryptAlgorithmClassValues(attr.Value)
	return nil
}

// PhoneticValues
type PhoneticValues string

const (
	PhoneticValuesHalfwidthkatakana PhoneticValues = "halfwidthKatakana"
	PhoneticValuesFullwidthkatakana PhoneticValues = "fullwidthKatakana"
	PhoneticValuesHiragana          PhoneticValues = "Hiragana"
	PhoneticValuesNoconversion      PhoneticValues = "noConversion"
)

func (e PhoneticValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PhoneticValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PhoneticValues(attr.Value)
	return nil
}

// TextStrikeValues
type TextStrikeValues string

const (
	TextStrikeValuesNostrike  TextStrikeValues = "noStrike"
	TextStrikeValuesSngstrike TextStrikeValues = "sngStrike"
	TextStrikeValuesDblstrike TextStrikeValues = "dblStrike"
)

func (e TextStrikeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextStrikeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextStrikeValues(attr.Value)
	return nil
}

// TextTabAlignmentValues
type TextTabAlignmentValues string

const (
	TextTabAlignmentValuesL   TextTabAlignmentValues = "l"
	TextTabAlignmentValuesCtr TextTabAlignmentValues = "ctr"
	TextTabAlignmentValuesR   TextTabAlignmentValues = "r"
	TextTabAlignmentValuesDec TextTabAlignmentValues = "dec"
)

func (e TextTabAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextTabAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextTabAlignmentValues(attr.Value)
	return nil
}

// ReferenceModeValues
type ReferenceModeValues string

const (
	ReferenceModeValuesA1   ReferenceModeValues = "A1"
	ReferenceModeValuesR1c1 ReferenceModeValues = "R1C1"
)

func (e ReferenceModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ReferenceModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ReferenceModeValues(attr.Value)
	return nil
}

// ProofingErrorValues
type ProofingErrorValues string

const (
	ProofingErrorValuesSpellstart ProofingErrorValues = "spellStart"
	ProofingErrorValuesSpellend   ProofingErrorValues = "spellEnd"
	ProofingErrorValuesGramstart  ProofingErrorValues = "gramStart"
	ProofingErrorValuesGramend    ProofingErrorValues = "gramEnd"
)

func (e ProofingErrorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ProofingErrorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ProofingErrorValues(attr.Value)
	return nil
}

// SparklineTypeValues
type SparklineTypeValues string

const (
	SparklineTypeValuesLine    SparklineTypeValues = "line"
	SparklineTypeValuesColumn  SparklineTypeValues = "column"
	SparklineTypeValuesStacked SparklineTypeValues = "stacked"
)

func (e SparklineTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SparklineTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SparklineTypeValues(attr.Value)
	return nil
}

// AnimationChartOnlyBuildValues
type AnimationChartOnlyBuildValues string

const (
	AnimationChartOnlyBuildValuesSeries     AnimationChartOnlyBuildValues = "series"
	AnimationChartOnlyBuildValuesCategory   AnimationChartOnlyBuildValues = "category"
	AnimationChartOnlyBuildValuesSeriesel   AnimationChartOnlyBuildValues = "seriesEl"
	AnimationChartOnlyBuildValuesCategoryel AnimationChartOnlyBuildValues = "categoryEl"
)

func (e AnimationChartOnlyBuildValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AnimationChartOnlyBuildValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AnimationChartOnlyBuildValues(attr.Value)
	return nil
}

// PresetPatternValues
type PresetPatternValues string

const (
	PresetPatternValuesPct5       PresetPatternValues = "pct5"
	PresetPatternValuesPct10      PresetPatternValues = "pct10"
	PresetPatternValuesPct20      PresetPatternValues = "pct20"
	PresetPatternValuesPct25      PresetPatternValues = "pct25"
	PresetPatternValuesPct30      PresetPatternValues = "pct30"
	PresetPatternValuesPct40      PresetPatternValues = "pct40"
	PresetPatternValuesPct50      PresetPatternValues = "pct50"
	PresetPatternValuesPct60      PresetPatternValues = "pct60"
	PresetPatternValuesPct70      PresetPatternValues = "pct70"
	PresetPatternValuesPct75      PresetPatternValues = "pct75"
	PresetPatternValuesPct80      PresetPatternValues = "pct80"
	PresetPatternValuesPct90      PresetPatternValues = "pct90"
	PresetPatternValuesHorz       PresetPatternValues = "horz"
	PresetPatternValuesVert       PresetPatternValues = "vert"
	PresetPatternValuesLthorz     PresetPatternValues = "ltHorz"
	PresetPatternValuesLtvert     PresetPatternValues = "ltVert"
	PresetPatternValuesDkhorz     PresetPatternValues = "dkHorz"
	PresetPatternValuesDkvert     PresetPatternValues = "dkVert"
	PresetPatternValuesNarhorz    PresetPatternValues = "narHorz"
	PresetPatternValuesNarvert    PresetPatternValues = "narVert"
	PresetPatternValuesDashhorz   PresetPatternValues = "dashHorz"
	PresetPatternValuesDashvert   PresetPatternValues = "dashVert"
	PresetPatternValuesCross      PresetPatternValues = "cross"
	PresetPatternValuesDndiag     PresetPatternValues = "dnDiag"
	PresetPatternValuesUpdiag     PresetPatternValues = "upDiag"
	PresetPatternValuesLtdndiag   PresetPatternValues = "ltDnDiag"
	PresetPatternValuesLtupdiag   PresetPatternValues = "ltUpDiag"
	PresetPatternValuesDkdndiag   PresetPatternValues = "dkDnDiag"
	PresetPatternValuesDkupdiag   PresetPatternValues = "dkUpDiag"
	PresetPatternValuesWddndiag   PresetPatternValues = "wdDnDiag"
	PresetPatternValuesWdupdiag   PresetPatternValues = "wdUpDiag"
	PresetPatternValuesDashdndiag PresetPatternValues = "dashDnDiag"
	PresetPatternValuesDashupdiag PresetPatternValues = "dashUpDiag"
	PresetPatternValuesDiagcross  PresetPatternValues = "diagCross"
	PresetPatternValuesSmcheck    PresetPatternValues = "smCheck"
	PresetPatternValuesLgcheck    PresetPatternValues = "lgCheck"
	PresetPatternValuesSmgrid     PresetPatternValues = "smGrid"
	PresetPatternValuesLggrid     PresetPatternValues = "lgGrid"
	PresetPatternValuesDotgrid    PresetPatternValues = "dotGrid"
	PresetPatternValuesSmconfetti PresetPatternValues = "smConfetti"
	PresetPatternValuesLgconfetti PresetPatternValues = "lgConfetti"
	PresetPatternValuesHorzbrick  PresetPatternValues = "horzBrick"
	PresetPatternValuesDiagbrick  PresetPatternValues = "diagBrick"
	PresetPatternValuesSoliddmnd  PresetPatternValues = "solidDmnd"
	PresetPatternValuesOpendmnd   PresetPatternValues = "openDmnd"
	PresetPatternValuesDotdmnd    PresetPatternValues = "dotDmnd"
	PresetPatternValuesPlaid      PresetPatternValues = "plaid"
	PresetPatternValuesSphere     PresetPatternValues = "sphere"
	PresetPatternValuesWeave      PresetPatternValues = "weave"
	PresetPatternValuesDivot      PresetPatternValues = "divot"
	PresetPatternValuesShingle    PresetPatternValues = "shingle"
	PresetPatternValuesWave       PresetPatternValues = "wave"
	PresetPatternValuesTrellis    PresetPatternValues = "trellis"
	PresetPatternValuesZigzag     PresetPatternValues = "zigZag"
)

func (e PresetPatternValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PresetPatternValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PresetPatternValues(attr.Value)
	return nil
}

// LineEndLengthValues
type LineEndLengthValues string

const (
	LineEndLengthValuesSm  LineEndLengthValues = "sm"
	LineEndLengthValuesMed LineEndLengthValues = "med"
	LineEndLengthValuesLg  LineEndLengthValues = "lg"
)

func (e LineEndLengthValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LineEndLengthValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LineEndLengthValues(attr.Value)
	return nil
}

// TableVerticalAlignmentValues
type TableVerticalAlignmentValues string

const (
	TableVerticalAlignmentValuesTop    TableVerticalAlignmentValues = "top"
	TableVerticalAlignmentValuesCenter TableVerticalAlignmentValues = "center"
	TableVerticalAlignmentValuesBottom TableVerticalAlignmentValues = "bottom"
)

func (e TableVerticalAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TableVerticalAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TableVerticalAlignmentValues(attr.Value)
	return nil
}

// PresetShadowValues
type PresetShadowValues string

const (
	PresetShadowValuesShdw1  PresetShadowValues = "shdw1"
	PresetShadowValuesShdw2  PresetShadowValues = "shdw2"
	PresetShadowValuesShdw3  PresetShadowValues = "shdw3"
	PresetShadowValuesShdw4  PresetShadowValues = "shdw4"
	PresetShadowValuesShdw5  PresetShadowValues = "shdw5"
	PresetShadowValuesShdw6  PresetShadowValues = "shdw6"
	PresetShadowValuesShdw7  PresetShadowValues = "shdw7"
	PresetShadowValuesShdw8  PresetShadowValues = "shdw8"
	PresetShadowValuesShdw9  PresetShadowValues = "shdw9"
	PresetShadowValuesShdw10 PresetShadowValues = "shdw10"
	PresetShadowValuesShdw11 PresetShadowValues = "shdw11"
	PresetShadowValuesShdw12 PresetShadowValues = "shdw12"
	PresetShadowValuesShdw13 PresetShadowValues = "shdw13"
	PresetShadowValuesShdw14 PresetShadowValues = "shdw14"
	PresetShadowValuesShdw15 PresetShadowValues = "shdw15"
	PresetShadowValuesShdw16 PresetShadowValues = "shdw16"
	PresetShadowValuesShdw17 PresetShadowValues = "shdw17"
	PresetShadowValuesShdw18 PresetShadowValues = "shdw18"
	PresetShadowValuesShdw19 PresetShadowValues = "shdw19"
	PresetShadowValuesShdw20 PresetShadowValues = "shdw20"
)

func (e PresetShadowValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PresetShadowValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PresetShadowValues(attr.Value)
	return nil
}

// TextHorizontalOverflowValues
type TextHorizontalOverflowValues string

const (
	TextHorizontalOverflowValuesOverflow TextHorizontalOverflowValues = "overflow"
	TextHorizontalOverflowValuesClip     TextHorizontalOverflowValues = "clip"
)

func (e TextHorizontalOverflowValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextHorizontalOverflowValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextHorizontalOverflowValues(attr.Value)
	return nil
}

// TransitionInOutDirectionValues
type TransitionInOutDirectionValues string

const (
	TransitionInOutDirectionValuesOut TransitionInOutDirectionValues = "out"
	TransitionInOutDirectionValuesIn  TransitionInOutDirectionValues = "in"
)

func (e TransitionInOutDirectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TransitionInOutDirectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TransitionInOutDirectionValues(
		attr.Value,
	)
	return nil
}

// AnimateMotionBehaviorOriginValues
type AnimateMotionBehaviorOriginValues string

const (
	AnimateMotionBehaviorOriginValuesParent AnimateMotionBehaviorOriginValues = "parent"
	AnimateMotionBehaviorOriginValuesLayout AnimateMotionBehaviorOriginValues = "layout"
)

func (e AnimateMotionBehaviorOriginValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AnimateMotionBehaviorOriginValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AnimateMotionBehaviorOriginValues(
		attr.Value,
	)
	return nil
}

// FieldCharValues
type FieldCharValues string

const (
	FieldCharValuesBegin    FieldCharValues = "begin"
	FieldCharValuesSeparate FieldCharValues = "separate"
	FieldCharValuesEnd      FieldCharValues = "end"
)

func (e FieldCharValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FieldCharValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FieldCharValues(attr.Value)
	return nil
}

// KnownSemanticTypeValues
type KnownSemanticTypeValues string

const (
	KnownSemanticTypeValuesNone          KnownSemanticTypeValues = "none"
	KnownSemanticTypeValuesUnderline     KnownSemanticTypeValues = "underline"
	KnownSemanticTypeValuesStrikethrough KnownSemanticTypeValues = "strikethrough"
	KnownSemanticTypeValuesHighlight     KnownSemanticTypeValues = "highlight"
	KnownSemanticTypeValuesScratchout    KnownSemanticTypeValues = "scratchOut"
	KnownSemanticTypeValuesVerticalrange KnownSemanticTypeValues = "verticalRange"
	KnownSemanticTypeValuesCallout       KnownSemanticTypeValues = "callout"
	KnownSemanticTypeValuesEnclosure     KnownSemanticTypeValues = "enclosure"
	KnownSemanticTypeValuesComment       KnownSemanticTypeValues = "comment"
	KnownSemanticTypeValuesContainer     KnownSemanticTypeValues = "container"
	KnownSemanticTypeValuesConnector     KnownSemanticTypeValues = "connector"
)

func (e KnownSemanticTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *KnownSemanticTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = KnownSemanticTypeValues(attr.Value)
	return nil
}

// SXVCellType
type SXVCellType string

const (
	SXVCellTypeB   SXVCellType = "b"
	SXVCellTypeN   SXVCellType = "n"
	SXVCellTypeE   SXVCellType = "e"
	SXVCellTypeStr SXVCellType = "str"
	SXVCellTypeD   SXVCellType = "d"
	SXVCellTypeBl  SXVCellType = "bl"
)

func (e SXVCellType) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SXVCellType) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SXVCellType(attr.Value)
	return nil
}

// TriggerEventValues
type TriggerEventValues string

const (
	TriggerEventValuesNone            TriggerEventValues = "none"
	TriggerEventValuesOnbegin         TriggerEventValues = "onBegin"
	TriggerEventValuesOnend           TriggerEventValues = "onEnd"
	TriggerEventValuesBegin           TriggerEventValues = "begin"
	TriggerEventValuesEnd             TriggerEventValues = "end"
	TriggerEventValuesOnclick         TriggerEventValues = "onClick"
	TriggerEventValuesOndblclick      TriggerEventValues = "onDblClick"
	TriggerEventValuesOnmouseover     TriggerEventValues = "onMouseOver"
	TriggerEventValuesOnmouseout      TriggerEventValues = "onMouseOut"
	TriggerEventValuesOnnext          TriggerEventValues = "onNext"
	TriggerEventValuesOnprev          TriggerEventValues = "onPrev"
	TriggerEventValuesOnstopaudio     TriggerEventValues = "onStopAudio"
	TriggerEventValuesOnmediabookmark TriggerEventValues = "onMediaBookmark"
)

func (e TriggerEventValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TriggerEventValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TriggerEventValues(attr.Value)
	return nil
}

// ParameterValues
type ParameterValues string

const (
	ParameterValuesPrompt ParameterValues = "prompt"
	ParameterValuesValue  ParameterValues = "value"
	ParameterValuesCell   ParameterValues = "cell"
)

func (e ParameterValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ParameterValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ParameterValues(attr.Value)
	return nil
}

// PivotFilterValues
type PivotFilterValues string

const (
	PivotFilterValuesUnknown                   PivotFilterValues = "unknown"
	PivotFilterValuesCount                     PivotFilterValues = "count"
	PivotFilterValuesPercent                   PivotFilterValues = "percent"
	PivotFilterValuesSum                       PivotFilterValues = "sum"
	PivotFilterValuesCaptionequal              PivotFilterValues = "captionEqual"
	PivotFilterValuesCaptionnotequal           PivotFilterValues = "captionNotEqual"
	PivotFilterValuesCaptionbeginswith         PivotFilterValues = "captionBeginsWith"
	PivotFilterValuesCaptionnotbeginswith      PivotFilterValues = "captionNotBeginsWith"
	PivotFilterValuesCaptionendswith           PivotFilterValues = "captionEndsWith"
	PivotFilterValuesCaptionnotendswith        PivotFilterValues = "captionNotEndsWith"
	PivotFilterValuesCaptioncontains           PivotFilterValues = "captionContains"
	PivotFilterValuesCaptionnotcontains        PivotFilterValues = "captionNotContains"
	PivotFilterValuesCaptiongreaterthan        PivotFilterValues = "captionGreaterThan"
	PivotFilterValuesCaptiongreaterthanorequal PivotFilterValues = "captionGreaterThanOrEqual"
	PivotFilterValuesCaptionlessthan           PivotFilterValues = "captionLessThan"
	PivotFilterValuesCaptionlessthanorequal    PivotFilterValues = "captionLessThanOrEqual"
	PivotFilterValuesCaptionbetween            PivotFilterValues = "captionBetween"
	PivotFilterValuesCaptionnotbetween         PivotFilterValues = "captionNotBetween"
	PivotFilterValuesValueequal                PivotFilterValues = "valueEqual"
	PivotFilterValuesValuenotequal             PivotFilterValues = "valueNotEqual"
	PivotFilterValuesValuegreaterthan          PivotFilterValues = "valueGreaterThan"
	PivotFilterValuesValuegreaterthanorequal   PivotFilterValues = "valueGreaterThanOrEqual"
	PivotFilterValuesValuelessthan             PivotFilterValues = "valueLessThan"
	PivotFilterValuesValuelessthanorequal      PivotFilterValues = "valueLessThanOrEqual"
	PivotFilterValuesValuebetween              PivotFilterValues = "valueBetween"
	PivotFilterValuesValuenotbetween           PivotFilterValues = "valueNotBetween"
	PivotFilterValuesDateequal                 PivotFilterValues = "dateEqual"
	PivotFilterValuesDatenotequal              PivotFilterValues = "dateNotEqual"
	PivotFilterValuesDateolderthan             PivotFilterValues = "dateOlderThan"
	PivotFilterValuesDateolderthanorequal      PivotFilterValues = "dateOlderThanOrEqual"
	PivotFilterValuesDatenewerthan             PivotFilterValues = "dateNewerThan"
	PivotFilterValuesDatenewerthanorequal      PivotFilterValues = "dateNewerThanOrEqual"
	PivotFilterValuesDatebetween               PivotFilterValues = "dateBetween"
	PivotFilterValuesDatenotbetween            PivotFilterValues = "dateNotBetween"
	PivotFilterValuesTomorrow                  PivotFilterValues = "tomorrow"
	PivotFilterValuesToday                     PivotFilterValues = "today"
	PivotFilterValuesYesterday                 PivotFilterValues = "yesterday"
	PivotFilterValuesNextweek                  PivotFilterValues = "nextWeek"
	PivotFilterValuesThisweek                  PivotFilterValues = "thisWeek"
	PivotFilterValuesLastweek                  PivotFilterValues = "lastWeek"
	PivotFilterValuesNextmonth                 PivotFilterValues = "nextMonth"
	PivotFilterValuesThismonth                 PivotFilterValues = "thisMonth"
	PivotFilterValuesLastmonth                 PivotFilterValues = "lastMonth"
	PivotFilterValuesNextquarter               PivotFilterValues = "nextQuarter"
	PivotFilterValuesThisquarter               PivotFilterValues = "thisQuarter"
	PivotFilterValuesLastquarter               PivotFilterValues = "lastQuarter"
	PivotFilterValuesNextyear                  PivotFilterValues = "nextYear"
	PivotFilterValuesThisyear                  PivotFilterValues = "thisYear"
	PivotFilterValuesLastyear                  PivotFilterValues = "lastYear"
	PivotFilterValuesYeartodate                PivotFilterValues = "yearToDate"
	PivotFilterValuesQ1                        PivotFilterValues = "Q1"
	PivotFilterValuesQ2                        PivotFilterValues = "Q2"
	PivotFilterValuesQ3                        PivotFilterValues = "Q3"
	PivotFilterValuesQ4                        PivotFilterValues = "Q4"
	PivotFilterValuesM1                        PivotFilterValues = "M1"
	PivotFilterValuesM2                        PivotFilterValues = "M2"
	PivotFilterValuesM3                        PivotFilterValues = "M3"
	PivotFilterValuesM4                        PivotFilterValues = "M4"
	PivotFilterValuesM5                        PivotFilterValues = "M5"
	PivotFilterValuesM6                        PivotFilterValues = "M6"
	PivotFilterValuesM7                        PivotFilterValues = "M7"
	PivotFilterValuesM8                        PivotFilterValues = "M8"
	PivotFilterValuesM9                        PivotFilterValues = "M9"
	PivotFilterValuesM10                       PivotFilterValues = "M10"
	PivotFilterValuesM11                       PivotFilterValues = "M11"
	PivotFilterValuesM12                       PivotFilterValues = "M12"
)

func (e PivotFilterValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PivotFilterValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PivotFilterValues(attr.Value)
	return nil
}

// DataConsolidateFunctionValues
type DataConsolidateFunctionValues string

const (
	DataConsolidateFunctionValuesAverage   DataConsolidateFunctionValues = "average"
	DataConsolidateFunctionValuesCount     DataConsolidateFunctionValues = "count"
	DataConsolidateFunctionValuesCountnums DataConsolidateFunctionValues = "countNums"
	DataConsolidateFunctionValuesMax       DataConsolidateFunctionValues = "max"
	DataConsolidateFunctionValuesMin       DataConsolidateFunctionValues = "min"
	DataConsolidateFunctionValuesProduct   DataConsolidateFunctionValues = "product"
	DataConsolidateFunctionValuesStddev    DataConsolidateFunctionValues = "stdDev"
	DataConsolidateFunctionValuesStddevp   DataConsolidateFunctionValues = "stdDevp"
	DataConsolidateFunctionValuesSum       DataConsolidateFunctionValues = "sum"
	DataConsolidateFunctionValuesVar       DataConsolidateFunctionValues = "var"
	DataConsolidateFunctionValuesVarp      DataConsolidateFunctionValues = "varp"
)

func (e DataConsolidateFunctionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DataConsolidateFunctionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DataConsolidateFunctionValues(attr.Value)
	return nil
}

// PathFillModeValues
type PathFillModeValues string

const (
	PathFillModeValuesNone        PathFillModeValues = "none"
	PathFillModeValuesNorm        PathFillModeValues = "norm"
	PathFillModeValuesLighten     PathFillModeValues = "lighten"
	PathFillModeValuesLightenless PathFillModeValues = "lightenLess"
	PathFillModeValuesDarken      PathFillModeValues = "darken"
	PathFillModeValuesDarkenless  PathFillModeValues = "darkenLess"
)

func (e PathFillModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PathFillModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PathFillModeValues(attr.Value)
	return nil
}

// TableValues
type TableValues string

const (
	TableValuesWorksheet  TableValues = "worksheet"
	TableValuesXml        TableValues = "xml"
	TableValuesQuerytable TableValues = "queryTable"
)

func (e TableValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TableValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TableValues(attr.Value)
	return nil
}

// HeightRuleValues
type HeightRuleValues string

const (
	HeightRuleValuesAuto    HeightRuleValues = "auto"
	HeightRuleValuesExact   HeightRuleValues = "exact"
	HeightRuleValuesAtleast HeightRuleValues = "atLeast"
)

func (e HeightRuleValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *HeightRuleValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = HeightRuleValues(attr.Value)
	return nil
}

// TableWidthValues
type TableWidthValues string

const (
	TableWidthValuesNil TableWidthValues = "nil"
	TableWidthValuesDxa TableWidthValues = "dxa"
)

func (e TableWidthValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TableWidthValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TableWidthValues(attr.Value)
	return nil
}

// CalculatedMemberNumberFormat
type CalculatedMemberNumberFormat string

const (
	CalculatedMemberNumberFormatDefault CalculatedMemberNumberFormat = "default"
	CalculatedMemberNumberFormatNumber  CalculatedMemberNumberFormat = "number"
	CalculatedMemberNumberFormatPercent CalculatedMemberNumberFormat = "percent"
)

func (e CalculatedMemberNumberFormat) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CalculatedMemberNumberFormat) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CalculatedMemberNumberFormat(attr.Value)
	return nil
}

// SurveyPosition
type SurveyPosition string

const (
	SurveyPositionAbsolute SurveyPosition = "absolute"
	SurveyPositionFixed    SurveyPosition = "fixed"
	SurveyPositionRelative SurveyPosition = "relative"
	SurveyPositionStatic   SurveyPosition = "static"
	SurveyPositionInherit  SurveyPosition = "inherit"
)

func (e SurveyPosition) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SurveyPosition) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SurveyPosition(attr.Value)
	return nil
}

// BlackWhiteModeValues
type BlackWhiteModeValues string

const (
	BlackWhiteModeValuesClr        BlackWhiteModeValues = "clr"
	BlackWhiteModeValuesAuto       BlackWhiteModeValues = "auto"
	BlackWhiteModeValuesGray       BlackWhiteModeValues = "gray"
	BlackWhiteModeValuesLtgray     BlackWhiteModeValues = "ltGray"
	BlackWhiteModeValuesInvgray    BlackWhiteModeValues = "invGray"
	BlackWhiteModeValuesGraywhite  BlackWhiteModeValues = "grayWhite"
	BlackWhiteModeValuesBlackgray  BlackWhiteModeValues = "blackGray"
	BlackWhiteModeValuesBlackwhite BlackWhiteModeValues = "blackWhite"
	BlackWhiteModeValuesBlack      BlackWhiteModeValues = "black"
	BlackWhiteModeValuesWhite      BlackWhiteModeValues = "white"
	BlackWhiteModeValuesHidden     BlackWhiteModeValues = "hidden"
)

func (e BlackWhiteModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BlackWhiteModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BlackWhiteModeValues(attr.Value)
	return nil
}

// TextAlignmentTypeValues
type TextAlignmentTypeValues string

const (
	TextAlignmentTypeValuesL        TextAlignmentTypeValues = "l"
	TextAlignmentTypeValuesCtr      TextAlignmentTypeValues = "ctr"
	TextAlignmentTypeValuesR        TextAlignmentTypeValues = "r"
	TextAlignmentTypeValuesJust     TextAlignmentTypeValues = "just"
	TextAlignmentTypeValuesJustlow  TextAlignmentTypeValues = "justLow"
	TextAlignmentTypeValuesDist     TextAlignmentTypeValues = "dist"
	TextAlignmentTypeValuesThaidist TextAlignmentTypeValues = "thaiDist"
)

func (e TextAlignmentTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextAlignmentTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextAlignmentTypeValues(attr.Value)
	return nil
}

// TimeNodeValues
type TimeNodeValues string

const (
	TimeNodeValuesClickeffect    TimeNodeValues = "clickEffect"
	TimeNodeValuesWitheffect     TimeNodeValues = "withEffect"
	TimeNodeValuesAftereffect    TimeNodeValues = "afterEffect"
	TimeNodeValuesMainseq        TimeNodeValues = "mainSeq"
	TimeNodeValuesInteractiveseq TimeNodeValues = "interactiveSeq"
	TimeNodeValuesClickpar       TimeNodeValues = "clickPar"
	TimeNodeValuesWithgroup      TimeNodeValues = "withGroup"
	TimeNodeValuesAftergroup     TimeNodeValues = "afterGroup"
	TimeNodeValuesTmroot         TimeNodeValues = "tmRoot"
)

func (e TimeNodeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TimeNodeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TimeNodeValues(attr.Value)
	return nil
}

// VerticalPositionValues
type VerticalPositionValues string

const (
	VerticalPositionValuesBaseline    VerticalPositionValues = "baseline"
	VerticalPositionValuesSuperscript VerticalPositionValues = "superscript"
	VerticalPositionValuesSubscript   VerticalPositionValues = "subscript"
)

func (e VerticalPositionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *VerticalPositionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = VerticalPositionValues(attr.Value)
	return nil
}

// MailMergeOdsoFieldValues
type MailMergeOdsoFieldValues string

const (
	MailMergeOdsoFieldValuesNull     MailMergeOdsoFieldValues = "null"
	MailMergeOdsoFieldValuesDbcolumn MailMergeOdsoFieldValues = "dbColumn"
)

func (e MailMergeOdsoFieldValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MailMergeOdsoFieldValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MailMergeOdsoFieldValues(attr.Value)
	return nil
}

// FrameScrollbarVisibilityValues
type FrameScrollbarVisibilityValues string

const (
	FrameScrollbarVisibilityValuesOn   FrameScrollbarVisibilityValues = "on"
	FrameScrollbarVisibilityValuesOff  FrameScrollbarVisibilityValues = "off"
	FrameScrollbarVisibilityValuesAuto FrameScrollbarVisibilityValues = "auto"
)

func (e FrameScrollbarVisibilityValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FrameScrollbarVisibilityValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FrameScrollbarVisibilityValues(
		attr.Value,
	)
	return nil
}

// TransitionCenterDirectionTypeValues
type TransitionCenterDirectionTypeValues string

const (
	TransitionCenterDirectionTypeValuesCenter TransitionCenterDirectionTypeValues = "center"
)

func (e TransitionCenterDirectionTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TransitionCenterDirectionTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TransitionCenterDirectionTypeValues(
		attr.Value,
	)
	return nil
}

// PenAlignmentValues
type PenAlignmentValues string

const (
	PenAlignmentValuesCtr PenAlignmentValues = "ctr"
	PenAlignmentValuesIn  PenAlignmentValues = "in"
)

func (e PenAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PenAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PenAlignmentValues(attr.Value)
	return nil
}

// TextUnderlineValues
type TextUnderlineValues string

const (
	TextUnderlineValuesNone            TextUnderlineValues = "none"
	TextUnderlineValuesWords           TextUnderlineValues = "words"
	TextUnderlineValuesSng             TextUnderlineValues = "sng"
	TextUnderlineValuesDbl             TextUnderlineValues = "dbl"
	TextUnderlineValuesHeavy           TextUnderlineValues = "heavy"
	TextUnderlineValuesDotted          TextUnderlineValues = "dotted"
	TextUnderlineValuesDottedheavy     TextUnderlineValues = "dottedHeavy"
	TextUnderlineValuesDash            TextUnderlineValues = "dash"
	TextUnderlineValuesDashheavy       TextUnderlineValues = "dashHeavy"
	TextUnderlineValuesDashlong        TextUnderlineValues = "dashLong"
	TextUnderlineValuesDashlongheavy   TextUnderlineValues = "dashLongHeavy"
	TextUnderlineValuesDotdash         TextUnderlineValues = "dotDash"
	TextUnderlineValuesDotdashheavy    TextUnderlineValues = "dotDashHeavy"
	TextUnderlineValuesDotdotdash      TextUnderlineValues = "dotDotDash"
	TextUnderlineValuesDotdotdashheavy TextUnderlineValues = "dotDotDashHeavy"
	TextUnderlineValuesWavy            TextUnderlineValues = "wavy"
	TextUnderlineValuesWavyheavy       TextUnderlineValues = "wavyHeavy"
	TextUnderlineValuesWavydbl         TextUnderlineValues = "wavyDbl"
)

func (e TextUnderlineValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextUnderlineValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextUnderlineValues(attr.Value)
	return nil
}

// AnimateBehaviorCalculateModeValues
type AnimateBehaviorCalculateModeValues string

const (
	AnimateBehaviorCalculateModeValuesDiscrete AnimateBehaviorCalculateModeValues = "discrete"
	AnimateBehaviorCalculateModeValuesLin      AnimateBehaviorCalculateModeValues = "lin"
	AnimateBehaviorCalculateModeValuesFmla     AnimateBehaviorCalculateModeValues = "fmla"
)

func (e AnimateBehaviorCalculateModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AnimateBehaviorCalculateModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AnimateBehaviorCalculateModeValues(
		attr.Value,
	)
	return nil
}

// ShowDataAsValues
type ShowDataAsValues string

const (
	ShowDataAsValuesNormal         ShowDataAsValues = "normal"
	ShowDataAsValuesDifference     ShowDataAsValues = "difference"
	ShowDataAsValuesPercent        ShowDataAsValues = "percent"
	ShowDataAsValuesPercentdiff    ShowDataAsValues = "percentDiff"
	ShowDataAsValuesRuntotal       ShowDataAsValues = "runTotal"
	ShowDataAsValuesPercentofrow   ShowDataAsValues = "percentOfRow"
	ShowDataAsValuesPercentofcol   ShowDataAsValues = "percentOfCol"
	ShowDataAsValuesPercentoftotal ShowDataAsValues = "percentOfTotal"
	ShowDataAsValuesIndex          ShowDataAsValues = "index"
)

func (e ShowDataAsValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ShowDataAsValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ShowDataAsValues(attr.Value)
	return nil
}

// DataValidationOperatorValues
type DataValidationOperatorValues string

const (
	DataValidationOperatorValuesBetween            DataValidationOperatorValues = "between"
	DataValidationOperatorValuesNotbetween         DataValidationOperatorValues = "notBetween"
	DataValidationOperatorValuesEqual              DataValidationOperatorValues = "equal"
	DataValidationOperatorValuesNotequal           DataValidationOperatorValues = "notEqual"
	DataValidationOperatorValuesLessthan           DataValidationOperatorValues = "lessThan"
	DataValidationOperatorValuesLessthanorequal    DataValidationOperatorValues = "lessThanOrEqual"
	DataValidationOperatorValuesGreaterthan        DataValidationOperatorValues = "greaterThan"
	DataValidationOperatorValuesGreaterthanorequal DataValidationOperatorValues = "greaterThanOrEqual"
)

func (e DataValidationOperatorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DataValidationOperatorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DataValidationOperatorValues(attr.Value)
	return nil
}

// TransitionLeftRightDirectionTypeValues
type TransitionLeftRightDirectionTypeValues string

const (
	TransitionLeftRightDirectionTypeValuesL TransitionLeftRightDirectionTypeValues = "l"
	TransitionLeftRightDirectionTypeValuesR TransitionLeftRightDirectionTypeValues = "r"
)

func (e TransitionLeftRightDirectionTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TransitionLeftRightDirectionTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TransitionLeftRightDirectionTypeValues(
		attr.Value,
	)
	return nil
}

// PresetCameraValues
type PresetCameraValues string

const (
	PresetCameraValuesLegacyobliquetopleft                PresetCameraValues = "legacyObliqueTopLeft"
	PresetCameraValuesLegacyobliquetop                    PresetCameraValues = "legacyObliqueTop"
	PresetCameraValuesLegacyobliquetopright               PresetCameraValues = "legacyObliqueTopRight"
	PresetCameraValuesLegacyobliqueleft                   PresetCameraValues = "legacyObliqueLeft"
	PresetCameraValuesLegacyobliquefront                  PresetCameraValues = "legacyObliqueFront"
	PresetCameraValuesLegacyobliqueright                  PresetCameraValues = "legacyObliqueRight"
	PresetCameraValuesLegacyobliquebottomleft             PresetCameraValues = "legacyObliqueBottomLeft"
	PresetCameraValuesLegacyobliquebottom                 PresetCameraValues = "legacyObliqueBottom"
	PresetCameraValuesLegacyobliquebottomright            PresetCameraValues = "legacyObliqueBottomRight"
	PresetCameraValuesLegacyperspectivetopleft            PresetCameraValues = "legacyPerspectiveTopLeft"
	PresetCameraValuesLegacyperspectivetop                PresetCameraValues = "legacyPerspectiveTop"
	PresetCameraValuesLegacyperspectivetopright           PresetCameraValues = "legacyPerspectiveTopRight"
	PresetCameraValuesLegacyperspectiveleft               PresetCameraValues = "legacyPerspectiveLeft"
	PresetCameraValuesLegacyperspectivefront              PresetCameraValues = "legacyPerspectiveFront"
	PresetCameraValuesLegacyperspectiveright              PresetCameraValues = "legacyPerspectiveRight"
	PresetCameraValuesLegacyperspectivebottomleft         PresetCameraValues = "legacyPerspectiveBottomLeft"
	PresetCameraValuesLegacyperspectivebottom             PresetCameraValues = "legacyPerspectiveBottom"
	PresetCameraValuesLegacyperspectivebottomright        PresetCameraValues = "legacyPerspectiveBottomRight"
	PresetCameraValuesOrthographicfront                   PresetCameraValues = "orthographicFront"
	PresetCameraValuesIsometrictopup                      PresetCameraValues = "isometricTopUp"
	PresetCameraValuesIsometrictopdown                    PresetCameraValues = "isometricTopDown"
	PresetCameraValuesIsometricbottomup                   PresetCameraValues = "isometricBottomUp"
	PresetCameraValuesIsometricbottomdown                 PresetCameraValues = "isometricBottomDown"
	PresetCameraValuesIsometricleftup                     PresetCameraValues = "isometricLeftUp"
	PresetCameraValuesIsometricleftdown                   PresetCameraValues = "isometricLeftDown"
	PresetCameraValuesIsometricrightup                    PresetCameraValues = "isometricRightUp"
	PresetCameraValuesIsometricrightdown                  PresetCameraValues = "isometricRightDown"
	PresetCameraValuesIsometricoffaxis1left               PresetCameraValues = "isometricOffAxis1Left"
	PresetCameraValuesIsometricoffaxis1right              PresetCameraValues = "isometricOffAxis1Right"
	PresetCameraValuesIsometricoffaxis1top                PresetCameraValues = "isometricOffAxis1Top"
	PresetCameraValuesIsometricoffaxis2left               PresetCameraValues = "isometricOffAxis2Left"
	PresetCameraValuesIsometricoffaxis2right              PresetCameraValues = "isometricOffAxis2Right"
	PresetCameraValuesIsometricoffaxis2top                PresetCameraValues = "isometricOffAxis2Top"
	PresetCameraValuesIsometricoffaxis3left               PresetCameraValues = "isometricOffAxis3Left"
	PresetCameraValuesIsometricoffaxis3right              PresetCameraValues = "isometricOffAxis3Right"
	PresetCameraValuesIsometricoffaxis3bottom             PresetCameraValues = "isometricOffAxis3Bottom"
	PresetCameraValuesIsometricoffaxis4left               PresetCameraValues = "isometricOffAxis4Left"
	PresetCameraValuesIsometricoffaxis4right              PresetCameraValues = "isometricOffAxis4Right"
	PresetCameraValuesIsometricoffaxis4bottom             PresetCameraValues = "isometricOffAxis4Bottom"
	PresetCameraValuesObliquetopleft                      PresetCameraValues = "obliqueTopLeft"
	PresetCameraValuesObliquetop                          PresetCameraValues = "obliqueTop"
	PresetCameraValuesObliquetopright                     PresetCameraValues = "obliqueTopRight"
	PresetCameraValuesObliqueleft                         PresetCameraValues = "obliqueLeft"
	PresetCameraValuesObliqueright                        PresetCameraValues = "obliqueRight"
	PresetCameraValuesObliquebottomleft                   PresetCameraValues = "obliqueBottomLeft"
	PresetCameraValuesObliquebottom                       PresetCameraValues = "obliqueBottom"
	PresetCameraValuesObliquebottomright                  PresetCameraValues = "obliqueBottomRight"
	PresetCameraValuesPerspectivefront                    PresetCameraValues = "perspectiveFront"
	PresetCameraValuesPerspectiveleft                     PresetCameraValues = "perspectiveLeft"
	PresetCameraValuesPerspectiveright                    PresetCameraValues = "perspectiveRight"
	PresetCameraValuesPerspectiveabove                    PresetCameraValues = "perspectiveAbove"
	PresetCameraValuesPerspectivebelow                    PresetCameraValues = "perspectiveBelow"
	PresetCameraValuesPerspectiveaboveleftfacing          PresetCameraValues = "perspectiveAboveLeftFacing"
	PresetCameraValuesPerspectiveaboverightfacing         PresetCameraValues = "perspectiveAboveRightFacing"
	PresetCameraValuesPerspectivecontrastingleftfacing    PresetCameraValues = "perspectiveContrastingLeftFacing"
	PresetCameraValuesPerspectivecontrastingrightfacing   PresetCameraValues = "perspectiveContrastingRightFacing"
	PresetCameraValuesPerspectiveheroicleftfacing         PresetCameraValues = "perspectiveHeroicLeftFacing"
	PresetCameraValuesPerspectiveheroicrightfacing        PresetCameraValues = "perspectiveHeroicRightFacing"
	PresetCameraValuesPerspectiveheroicextremeleftfacing  PresetCameraValues = "perspectiveHeroicExtremeLeftFacing"
	PresetCameraValuesPerspectiveheroicextremerightfacing PresetCameraValues = "perspectiveHeroicExtremeRightFacing"
	PresetCameraValuesPerspectiverelaxed                  PresetCameraValues = "perspectiveRelaxed"
	PresetCameraValuesPerspectiverelaxedmoderately        PresetCameraValues = "perspectiveRelaxedModerately"
)

func (e PresetCameraValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PresetCameraValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PresetCameraValues(attr.Value)
	return nil
}

// EffectContainerValues
type EffectContainerValues string

const (
	EffectContainerValuesSib  EffectContainerValues = "sib"
	EffectContainerValuesTree EffectContainerValues = "tree"
)

func (e EffectContainerValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *EffectContainerValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = EffectContainerValues(attr.Value)
	return nil
}

// TableStyleOverrideValues
type TableStyleOverrideValues string

const (
	TableStyleOverrideValuesWholetable TableStyleOverrideValues = "wholeTable"
	TableStyleOverrideValuesFirstrow   TableStyleOverrideValues = "firstRow"
	TableStyleOverrideValuesLastrow    TableStyleOverrideValues = "lastRow"
	TableStyleOverrideValuesFirstcol   TableStyleOverrideValues = "firstCol"
	TableStyleOverrideValuesLastcol    TableStyleOverrideValues = "lastCol"
	TableStyleOverrideValuesBand1vert  TableStyleOverrideValues = "band1Vert"
	TableStyleOverrideValuesBand2vert  TableStyleOverrideValues = "band2Vert"
	TableStyleOverrideValuesBand1horz  TableStyleOverrideValues = "band1Horz"
	TableStyleOverrideValuesBand2horz  TableStyleOverrideValues = "band2Horz"
	TableStyleOverrideValuesNecell     TableStyleOverrideValues = "neCell"
	TableStyleOverrideValuesNwcell     TableStyleOverrideValues = "nwCell"
	TableStyleOverrideValuesSecell     TableStyleOverrideValues = "seCell"
	TableStyleOverrideValuesSwcell     TableStyleOverrideValues = "swCell"
)

func (e TableStyleOverrideValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TableStyleOverrideValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TableStyleOverrideValues(attr.Value)
	return nil
}

// IconSetTypeValues
type IconSetTypeValues string

const (
	IconSetTypeValuesIconSetTypeValues3Arrows         IconSetTypeValues = "3Arrows"
	IconSetTypeValuesIconSetTypeValues3Arrowsgray     IconSetTypeValues = "3ArrowsGray"
	IconSetTypeValuesIconSetTypeValues3Flags          IconSetTypeValues = "3Flags"
	IconSetTypeValuesIconSetTypeValues3Trafficlights1 IconSetTypeValues = "3TrafficLights1"
	IconSetTypeValuesIconSetTypeValues3Trafficlights2 IconSetTypeValues = "3TrafficLights2"
	IconSetTypeValuesIconSetTypeValues3Signs          IconSetTypeValues = "3Signs"
	IconSetTypeValuesIconSetTypeValues3Symbols        IconSetTypeValues = "3Symbols"
	IconSetTypeValuesIconSetTypeValues3Symbols2       IconSetTypeValues = "3Symbols2"
	IconSetTypeValuesIconSetTypeValues4Arrows         IconSetTypeValues = "4Arrows"
	IconSetTypeValuesIconSetTypeValues4Arrowsgray     IconSetTypeValues = "4ArrowsGray"
	IconSetTypeValuesIconSetTypeValues4Redtoblack     IconSetTypeValues = "4RedToBlack"
	IconSetTypeValuesIconSetTypeValues4Rating         IconSetTypeValues = "4Rating"
	IconSetTypeValuesIconSetTypeValues4Trafficlights  IconSetTypeValues = "4TrafficLights"
	IconSetTypeValuesIconSetTypeValues5Arrows         IconSetTypeValues = "5Arrows"
	IconSetTypeValuesIconSetTypeValues5Arrowsgray     IconSetTypeValues = "5ArrowsGray"
	IconSetTypeValuesIconSetTypeValues5Rating         IconSetTypeValues = "5Rating"
	IconSetTypeValuesIconSetTypeValues5Quarters       IconSetTypeValues = "5Quarters"
	IconSetTypeValuesIconSetTypeValues3Stars          IconSetTypeValues = "3Stars"
	IconSetTypeValuesIconSetTypeValues3Triangles      IconSetTypeValues = "3Triangles"
	IconSetTypeValuesIconSetTypeValues5Boxes          IconSetTypeValues = "5Boxes"
	IconSetTypeValuesNoicons                          IconSetTypeValues = "NoIcons"
)

func (e IconSetTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *IconSetTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = IconSetTypeValues(attr.Value)
	return nil
}

// ConformanceClassValues
type ConformanceClassValues string

const (
	ConformanceClassValuesStrict       ConformanceClassValues = "strict"
	ConformanceClassValuesTransitional ConformanceClassValues = "transitional"
)

func (e ConformanceClassValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ConformanceClassValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ConformanceClassValues(attr.Value)
	return nil
}

// DateTimeGroupingValues
type DateTimeGroupingValues string

const (
	DateTimeGroupingValuesYear   DateTimeGroupingValues = "year"
	DateTimeGroupingValuesMonth  DateTimeGroupingValues = "month"
	DateTimeGroupingValuesDay    DateTimeGroupingValues = "day"
	DateTimeGroupingValuesHour   DateTimeGroupingValues = "hour"
	DateTimeGroupingValuesMinute DateTimeGroupingValues = "minute"
	DateTimeGroupingValuesSecond DateTimeGroupingValues = "second"
)

func (e DateTimeGroupingValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DateTimeGroupingValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DateTimeGroupingValues(attr.Value)
	return nil
}

// PivotAreaValues
type PivotAreaValues string

const (
	PivotAreaValuesNone     PivotAreaValues = "none"
	PivotAreaValuesNormal   PivotAreaValues = "normal"
	PivotAreaValuesData     PivotAreaValues = "data"
	PivotAreaValuesAll      PivotAreaValues = "all"
	PivotAreaValuesOrigin   PivotAreaValues = "origin"
	PivotAreaValuesButton   PivotAreaValues = "button"
	PivotAreaValuesTopright PivotAreaValues = "topRight"
	PivotAreaValuesTopend   PivotAreaValues = "topEnd"
)

func (e PivotAreaValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PivotAreaValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PivotAreaValues(attr.Value)
	return nil
}

// IterateValues
type IterateValues string

const (
	IterateValuesEl IterateValues = "el"
	IterateValuesWd IterateValues = "wd"
	IterateValuesLt IterateValues = "lt"
)

func (e IterateValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *IterateValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = IterateValues(attr.Value)
	return nil
}

// PrintColorModeValues
type PrintColorModeValues string

const (
	PrintColorModeValuesBw   PrintColorModeValues = "bw"
	PrintColorModeValuesGray PrintColorModeValues = "gray"
	PrintColorModeValuesClr  PrintColorModeValues = "clr"
)

func (e PrintColorModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PrintColorModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PrintColorModeValues(attr.Value)
	return nil
}

// TableLayoutValues
type TableLayoutValues string

const (
	TableLayoutValuesFixed   TableLayoutValues = "fixed"
	TableLayoutValuesAutofit TableLayoutValues = "autofit"
)

func (e TableLayoutValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TableLayoutValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TableLayoutValues(attr.Value)
	return nil
}

// EndnotePositionValues
type EndnotePositionValues string

const (
	EndnotePositionValuesSectend EndnotePositionValues = "sectEnd"
	EndnotePositionValuesDocend  EndnotePositionValues = "docEnd"
)

func (e EndnotePositionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *EndnotePositionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = EndnotePositionValues(attr.Value)
	return nil
}

// LinkDirectionValues
type LinkDirectionValues string

const (
	LinkDirectionValuesTo   LinkDirectionValues = "to"
	LinkDirectionValuesFrom LinkDirectionValues = "from"
	LinkDirectionValuesWith LinkDirectionValues = "with"
)

func (e LinkDirectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LinkDirectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LinkDirectionValues(attr.Value)
	return nil
}

// LightRigValues
type LightRigValues string

const (
	LightRigValuesLegacyflat1   LightRigValues = "legacyFlat1"
	LightRigValuesLegacyflat2   LightRigValues = "legacyFlat2"
	LightRigValuesLegacyflat3   LightRigValues = "legacyFlat3"
	LightRigValuesLegacyflat4   LightRigValues = "legacyFlat4"
	LightRigValuesLegacynormal1 LightRigValues = "legacyNormal1"
	LightRigValuesLegacynormal2 LightRigValues = "legacyNormal2"
	LightRigValuesLegacynormal3 LightRigValues = "legacyNormal3"
	LightRigValuesLegacynormal4 LightRigValues = "legacyNormal4"
	LightRigValuesLegacyharsh1  LightRigValues = "legacyHarsh1"
	LightRigValuesLegacyharsh2  LightRigValues = "legacyHarsh2"
	LightRigValuesLegacyharsh3  LightRigValues = "legacyHarsh3"
	LightRigValuesLegacyharsh4  LightRigValues = "legacyHarsh4"
	LightRigValuesThreept       LightRigValues = "threePt"
	LightRigValuesBalanced      LightRigValues = "balanced"
	LightRigValuesSoft          LightRigValues = "soft"
	LightRigValuesHarsh         LightRigValues = "harsh"
	LightRigValuesFlood         LightRigValues = "flood"
	LightRigValuesContrasting   LightRigValues = "contrasting"
	LightRigValuesMorning       LightRigValues = "morning"
	LightRigValuesSunrise       LightRigValues = "sunrise"
	LightRigValuesSunset        LightRigValues = "sunset"
	LightRigValuesChilly        LightRigValues = "chilly"
	LightRigValuesFreezing      LightRigValues = "freezing"
	LightRigValuesFlat          LightRigValues = "flat"
	LightRigValuesTwopt         LightRigValues = "twoPt"
	LightRigValuesGlow          LightRigValues = "glow"
	LightRigValuesBrightroom    LightRigValues = "brightRoom"
)

func (e LightRigValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *LightRigValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = LightRigValues(attr.Value)
	return nil
}

// ChartSubElementValues
type ChartSubElementValues string

const (
	ChartSubElementValuesGridlegend   ChartSubElementValues = "gridLegend"
	ChartSubElementValuesSeries       ChartSubElementValues = "series"
	ChartSubElementValuesCategory     ChartSubElementValues = "category"
	ChartSubElementValuesPtinseries   ChartSubElementValues = "ptInSeries"
	ChartSubElementValuesPtincategory ChartSubElementValues = "ptInCategory"
)

func (e ChartSubElementValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ChartSubElementValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ChartSubElementValues(attr.Value)
	return nil
}

// NextActionValues
type NextActionValues string

const (
	NextActionValuesNone NextActionValues = "none"
	NextActionValuesSeek NextActionValues = "seek"
)

func (e NextActionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *NextActionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = NextActionValues(attr.Value)
	return nil
}

// VisibilityValues
type VisibilityValues string

const (
	VisibilityValuesVisible    VisibilityValues = "visible"
	VisibilityValuesHidden     VisibilityValues = "hidden"
	VisibilityValuesVeryhidden VisibilityValues = "veryHidden"
)

func (e VisibilityValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *VisibilityValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = VisibilityValues(attr.Value)
	return nil
}

// CommentsValues
type CommentsValues string

const (
	CommentsValuesCommnone          CommentsValues = "commNone"
	CommentsValuesCommindicator     CommentsValues = "commIndicator"
	CommentsValuesCommindandcomment CommentsValues = "commIndAndComment"
)

func (e CommentsValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CommentsValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CommentsValues(attr.Value)
	return nil
}

// TargetScreenSizeValues
type TargetScreenSizeValues string

const (
	TargetScreenSizeValuesTargetScreenSizeValues544X376   TargetScreenSizeValues = "544x376"
	TargetScreenSizeValuesTargetScreenSizeValues640X480   TargetScreenSizeValues = "640x480"
	TargetScreenSizeValuesTargetScreenSizeValues720X512   TargetScreenSizeValues = "720x512"
	TargetScreenSizeValuesTargetScreenSizeValues800X600   TargetScreenSizeValues = "800x600"
	TargetScreenSizeValuesTargetScreenSizeValues1024X768  TargetScreenSizeValues = "1024x768"
	TargetScreenSizeValuesTargetScreenSizeValues1152X882  TargetScreenSizeValues = "1152x882"
	TargetScreenSizeValuesTargetScreenSizeValues1152X900  TargetScreenSizeValues = "1152x900"
	TargetScreenSizeValuesTargetScreenSizeValues1280X1024 TargetScreenSizeValues = "1280x1024"
	TargetScreenSizeValuesTargetScreenSizeValues1600X1200 TargetScreenSizeValues = "1600x1200"
	TargetScreenSizeValuesTargetScreenSizeValues1800X1440 TargetScreenSizeValues = "1800x1440"
	TargetScreenSizeValuesTargetScreenSizeValues1920X1200 TargetScreenSizeValues = "1920x1200"
)

func (e TargetScreenSizeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TargetScreenSizeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TargetScreenSizeValues(attr.Value)
	return nil
}

// TabStopLeaderCharValues
type TabStopLeaderCharValues string

const (
	TabStopLeaderCharValuesNone       TabStopLeaderCharValues = "none"
	TabStopLeaderCharValuesDot        TabStopLeaderCharValues = "dot"
	TabStopLeaderCharValuesHyphen     TabStopLeaderCharValues = "hyphen"
	TabStopLeaderCharValuesUnderscore TabStopLeaderCharValues = "underscore"
	TabStopLeaderCharValuesHeavy      TabStopLeaderCharValues = "heavy"
	TabStopLeaderCharValuesMiddledot  TabStopLeaderCharValues = "middleDot"
)

func (e TabStopLeaderCharValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TabStopLeaderCharValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TabStopLeaderCharValues(attr.Value)
	return nil
}

// OleUpdateValues
type OleUpdateValues string

const (
	OleUpdateValuesOleupdateAlways OleUpdateValues = "OLEUPDATE_ALWAYS"
	OleUpdateValuesOleupdateOncall OleUpdateValues = "OLEUPDATE_ONCALL"
)

func (e OleUpdateValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *OleUpdateValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = OleUpdateValues(attr.Value)
	return nil
}

// TextDirectionValues
type TextDirectionValues string

const (
	TextDirectionValuesLrtb  TextDirectionValues = "lrTb"
	TextDirectionValuesTb    TextDirectionValues = "tb"
	TextDirectionValuesTbrl  TextDirectionValues = "tbRl"
	TextDirectionValuesRl    TextDirectionValues = "rl"
	TextDirectionValuesBtlr  TextDirectionValues = "btLr"
	TextDirectionValuesLr    TextDirectionValues = "lr"
	TextDirectionValuesLrtbv TextDirectionValues = "lrTbV"
	TextDirectionValuesTbv   TextDirectionValues = "tbV"
	TextDirectionValuesTbrlv TextDirectionValues = "tbRlV"
	TextDirectionValuesRlv   TextDirectionValues = "rlV"
	TextDirectionValuesTblrv TextDirectionValues = "tbLrV"
	TextDirectionValuesLrv   TextDirectionValues = "lrV"
)

func (e TextDirectionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextDirectionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextDirectionValues(attr.Value)
	return nil
}

// ColorSchemeIndexValues
type ColorSchemeIndexValues string

const (
	ColorSchemeIndexValuesDark1             ColorSchemeIndexValues = "dark1"
	ColorSchemeIndexValuesLight1            ColorSchemeIndexValues = "light1"
	ColorSchemeIndexValuesDark2             ColorSchemeIndexValues = "dark2"
	ColorSchemeIndexValuesLight2            ColorSchemeIndexValues = "light2"
	ColorSchemeIndexValuesAccent1           ColorSchemeIndexValues = "accent1"
	ColorSchemeIndexValuesAccent2           ColorSchemeIndexValues = "accent2"
	ColorSchemeIndexValuesAccent3           ColorSchemeIndexValues = "accent3"
	ColorSchemeIndexValuesAccent4           ColorSchemeIndexValues = "accent4"
	ColorSchemeIndexValuesAccent5           ColorSchemeIndexValues = "accent5"
	ColorSchemeIndexValuesAccent6           ColorSchemeIndexValues = "accent6"
	ColorSchemeIndexValuesHyperlink         ColorSchemeIndexValues = "hyperlink"
	ColorSchemeIndexValuesFollowedhyperlink ColorSchemeIndexValues = "followedHyperlink"
)

func (e ColorSchemeIndexValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ColorSchemeIndexValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ColorSchemeIndexValues(attr.Value)
	return nil
}

// CellFormulaValues
type CellFormulaValues string

const (
	CellFormulaValuesNormal    CellFormulaValues = "normal"
	CellFormulaValuesArray     CellFormulaValues = "array"
	CellFormulaValuesDatatable CellFormulaValues = "dataTable"
	CellFormulaValuesShared    CellFormulaValues = "shared"
)

func (e CellFormulaValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CellFormulaValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CellFormulaValues(attr.Value)
	return nil
}

// ConditionalFormatValues
type ConditionalFormatValues string

const (
	ConditionalFormatValuesExpression        ConditionalFormatValues = "expression"
	ConditionalFormatValuesCellis            ConditionalFormatValues = "cellIs"
	ConditionalFormatValuesColorscale        ConditionalFormatValues = "colorScale"
	ConditionalFormatValuesDatabar           ConditionalFormatValues = "dataBar"
	ConditionalFormatValuesIconset           ConditionalFormatValues = "iconSet"
	ConditionalFormatValuesTop10             ConditionalFormatValues = "top10"
	ConditionalFormatValuesUniquevalues      ConditionalFormatValues = "uniqueValues"
	ConditionalFormatValuesDuplicatevalues   ConditionalFormatValues = "duplicateValues"
	ConditionalFormatValuesContainstext      ConditionalFormatValues = "containsText"
	ConditionalFormatValuesNotcontainstext   ConditionalFormatValues = "notContainsText"
	ConditionalFormatValuesBeginswith        ConditionalFormatValues = "beginsWith"
	ConditionalFormatValuesEndswith          ConditionalFormatValues = "endsWith"
	ConditionalFormatValuesContainsblanks    ConditionalFormatValues = "containsBlanks"
	ConditionalFormatValuesNotcontainsblanks ConditionalFormatValues = "notContainsBlanks"
	ConditionalFormatValuesContainserrors    ConditionalFormatValues = "containsErrors"
	ConditionalFormatValuesNotcontainserrors ConditionalFormatValues = "notContainsErrors"
	ConditionalFormatValuesTimeperiod        ConditionalFormatValues = "timePeriod"
	ConditionalFormatValuesAboveaverage      ConditionalFormatValues = "aboveAverage"
)

func (e ConditionalFormatValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ConditionalFormatValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ConditionalFormatValues(attr.Value)
	return nil
}

// TextVerticalOverflowValues
type TextVerticalOverflowValues string

const (
	TextVerticalOverflowValuesOverflow TextVerticalOverflowValues = "overflow"
	TextVerticalOverflowValuesEllipsis TextVerticalOverflowValues = "ellipsis"
	TextVerticalOverflowValuesClip     TextVerticalOverflowValues = "clip"
)

func (e TextVerticalOverflowValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextVerticalOverflowValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextVerticalOverflowValues(attr.Value)
	return nil
}

// OleChartBuildValues
type OleChartBuildValues string

const (
	OleChartBuildValuesAllatonce  OleChartBuildValues = "allAtOnce"
	OleChartBuildValuesSeries     OleChartBuildValues = "series"
	OleChartBuildValuesCategory   OleChartBuildValues = "category"
	OleChartBuildValuesSeriesel   OleChartBuildValues = "seriesEl"
	OleChartBuildValuesCategoryel OleChartBuildValues = "categoryEl"
)

func (e OleChartBuildValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *OleChartBuildValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = OleChartBuildValues(attr.Value)
	return nil
}

// PrintOutputValues
type PrintOutputValues string

const (
	PrintOutputValuesSlides    PrintOutputValues = "slides"
	PrintOutputValuesHandouts1 PrintOutputValues = "handouts1"
	PrintOutputValuesHandouts2 PrintOutputValues = "handouts2"
	PrintOutputValuesHandouts3 PrintOutputValues = "handouts3"
	PrintOutputValuesHandouts4 PrintOutputValues = "handouts4"
	PrintOutputValuesHandouts6 PrintOutputValues = "handouts6"
	PrintOutputValuesHandouts9 PrintOutputValues = "handouts9"
	PrintOutputValuesNotes     PrintOutputValues = "notes"
	PrintOutputValuesOutline   PrintOutputValues = "outline"
)

func (e PrintOutputValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PrintOutputValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PrintOutputValues(attr.Value)
	return nil
}

// DocumentConformance
type DocumentConformance string

const (
	DocumentConformanceTransitional DocumentConformance = "transitional"
	DocumentConformanceStrict       DocumentConformance = "strict"
)

func (e DocumentConformance) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DocumentConformance) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DocumentConformance(attr.Value)
	return nil
}

// CommentStatus
type CommentStatus string

const (
	CommentStatusActive   CommentStatus = "active"
	CommentStatusResolved CommentStatus = "resolved"
	CommentStatusClosed   CommentStatus = "closed"
)

func (e CommentStatus) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CommentStatus) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CommentStatus(attr.Value)
	return nil
}

// BehaviorAdditiveValues
type BehaviorAdditiveValues string

const (
	BehaviorAdditiveValuesBase BehaviorAdditiveValues = "base"
	BehaviorAdditiveValuesSum  BehaviorAdditiveValues = "sum"
	BehaviorAdditiveValuesRepl BehaviorAdditiveValues = "repl"
	BehaviorAdditiveValuesMult BehaviorAdditiveValues = "mult"
	BehaviorAdditiveValuesNone BehaviorAdditiveValues = "none"
)

func (e BehaviorAdditiveValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BehaviorAdditiveValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BehaviorAdditiveValues(attr.Value)
	return nil
}

// VerticalMergeRevisionValues
type VerticalMergeRevisionValues string

const (
	VerticalMergeRevisionValuesCont VerticalMergeRevisionValues = "cont"
	VerticalMergeRevisionValuesRest VerticalMergeRevisionValues = "rest"
)

func (e VerticalMergeRevisionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *VerticalMergeRevisionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = VerticalMergeRevisionValues(attr.Value)
	return nil
}

// StrictCharacterSet
type StrictCharacterSet string

const (
	StrictCharacterSetIso88591    StrictCharacterSet = "iso-8859-1"
	StrictCharacterSetMacintosh   StrictCharacterSet = "macintosh"
	StrictCharacterSetShiftJis    StrictCharacterSet = "shift_jis"
	StrictCharacterSetKsC56011987 StrictCharacterSet = "ks_c-5601-1987"
	StrictCharacterSetKsC56011992 StrictCharacterSet = "KS_C-5601-1992"
	StrictCharacterSetGbk         StrictCharacterSet = "GBK"
	StrictCharacterSetBig5        StrictCharacterSet = "Big5"
	StrictCharacterSetWindows1253 StrictCharacterSet = "windows-1253"
	StrictCharacterSetIso88599    StrictCharacterSet = "iso-8859-9"
	StrictCharacterSetWindows1258 StrictCharacterSet = "windows-1258"
	StrictCharacterSetWindows1255 StrictCharacterSet = "windows-1255"
	StrictCharacterSetWindows1256 StrictCharacterSet = "windows-1256"
	StrictCharacterSetWindows1257 StrictCharacterSet = "windows-1257"
	StrictCharacterSetWindows1251 StrictCharacterSet = "windows-1251"
	StrictCharacterSetWindows874  StrictCharacterSet = "windows-874"
	StrictCharacterSetWindows1250 StrictCharacterSet = "windows-1250"
)

func (e StrictCharacterSet) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *StrictCharacterSet) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = StrictCharacterSet(attr.Value)
	return nil
}

// KnownContextNodeTypeValues
type KnownContextNodeTypeValues string

const (
	KnownContextNodeTypeValuesRoot             KnownContextNodeTypeValues = "root"
	KnownContextNodeTypeValuesUnclassifiedink  KnownContextNodeTypeValues = "unclassifiedInk"
	KnownContextNodeTypeValuesWritingregion    KnownContextNodeTypeValues = "writingRegion"
	KnownContextNodeTypeValuesAnalysishint     KnownContextNodeTypeValues = "analysisHint"
	KnownContextNodeTypeValuesObject           KnownContextNodeTypeValues = "object"
	KnownContextNodeTypeValuesInkdrawing       KnownContextNodeTypeValues = "inkDrawing"
	KnownContextNodeTypeValuesImage            KnownContextNodeTypeValues = "image"
	KnownContextNodeTypeValuesParagraph        KnownContextNodeTypeValues = "paragraph"
	KnownContextNodeTypeValuesLine             KnownContextNodeTypeValues = "line"
	KnownContextNodeTypeValuesInkbullet        KnownContextNodeTypeValues = "inkBullet"
	KnownContextNodeTypeValuesInkword          KnownContextNodeTypeValues = "inkWord"
	KnownContextNodeTypeValuesTextword         KnownContextNodeTypeValues = "textWord"
	KnownContextNodeTypeValuesCustomrecognizer KnownContextNodeTypeValues = "customRecognizer"
	KnownContextNodeTypeValuesMathregion       KnownContextNodeTypeValues = "mathRegion"
	KnownContextNodeTypeValuesMathequation     KnownContextNodeTypeValues = "mathEquation"
	KnownContextNodeTypeValuesMathstruct       KnownContextNodeTypeValues = "mathStruct"
	KnownContextNodeTypeValuesMathsymbol       KnownContextNodeTypeValues = "mathSymbol"
	KnownContextNodeTypeValuesMathidentifier   KnownContextNodeTypeValues = "mathIdentifier"
	KnownContextNodeTypeValuesMathoperator     KnownContextNodeTypeValues = "mathOperator"
	KnownContextNodeTypeValuesMathnumber       KnownContextNodeTypeValues = "mathNumber"
	KnownContextNodeTypeValuesNoninkdrawing    KnownContextNodeTypeValues = "nonInkDrawing"
	KnownContextNodeTypeValuesGroupnode        KnownContextNodeTypeValues = "groupNode"
	KnownContextNodeTypeValuesMixeddrawing     KnownContextNodeTypeValues = "mixedDrawing"
)

func (e KnownContextNodeTypeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *KnownContextNodeTypeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = KnownContextNodeTypeValues(attr.Value)
	return nil
}

// TransitionPatternValues
type TransitionPatternValues string

const (
	TransitionPatternValuesDiamond TransitionPatternValues = "diamond"
	TransitionPatternValuesHexagon TransitionPatternValues = "hexagon"
)

func (e TransitionPatternValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TransitionPatternValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TransitionPatternValues(attr.Value)
	return nil
}

// DataValidationErrorStyleValues
type DataValidationErrorStyleValues string

const (
	DataValidationErrorStyleValuesStop        DataValidationErrorStyleValues = "stop"
	DataValidationErrorStyleValuesWarning     DataValidationErrorStyleValues = "warning"
	DataValidationErrorStyleValuesInformation DataValidationErrorStyleValues = "information"
)

func (e DataValidationErrorStyleValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DataValidationErrorStyleValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DataValidationErrorStyleValues(
		attr.Value,
	)
	return nil
}

// DropCapLocationValues
type DropCapLocationValues string

const (
	DropCapLocationValuesNone   DropCapLocationValues = "none"
	DropCapLocationValuesDrop   DropCapLocationValues = "drop"
	DropCapLocationValuesMargin DropCapLocationValues = "margin"
)

func (e DropCapLocationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DropCapLocationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DropCapLocationValues(attr.Value)
	return nil
}

// TextBoxFormFieldValues
type TextBoxFormFieldValues string

const (
	TextBoxFormFieldValuesRegular     TextBoxFormFieldValues = "regular"
	TextBoxFormFieldValuesNumber      TextBoxFormFieldValues = "number"
	TextBoxFormFieldValuesDate        TextBoxFormFieldValues = "date"
	TextBoxFormFieldValuesCurrenttime TextBoxFormFieldValues = "currentTime"
	TextBoxFormFieldValuesCurrentdate TextBoxFormFieldValues = "currentDate"
	TextBoxFormFieldValuesCalculated  TextBoxFormFieldValues = "calculated"
)

func (e TextBoxFormFieldValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextBoxFormFieldValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextBoxFormFieldValues(attr.Value)
	return nil
}

// MultiLevelValues
type MultiLevelValues string

const (
	MultiLevelValuesSinglelevel      MultiLevelValues = "singleLevel"
	MultiLevelValuesMultilevel       MultiLevelValues = "multilevel"
	MultiLevelValuesHybridmultilevel MultiLevelValues = "hybridMultilevel"
)

func (e MultiLevelValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MultiLevelValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MultiLevelValues(attr.Value)
	return nil
}

// DataBarAxisPositionValues
type DataBarAxisPositionValues string

const (
	DataBarAxisPositionValuesAutomatic DataBarAxisPositionValues = "automatic"
	DataBarAxisPositionValuesMiddle    DataBarAxisPositionValues = "middle"
	DataBarAxisPositionValuesNone      DataBarAxisPositionValues = "none"
)

func (e DataBarAxisPositionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DataBarAxisPositionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DataBarAxisPositionValues(attr.Value)
	return nil
}

// BehaviorAccumulateValues
type BehaviorAccumulateValues string

const (
	BehaviorAccumulateValuesNone   BehaviorAccumulateValues = "none"
	BehaviorAccumulateValuesAlways BehaviorAccumulateValues = "always"
)

func (e BehaviorAccumulateValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BehaviorAccumulateValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BehaviorAccumulateValues(attr.Value)
	return nil
}

// SortMethodValues
type SortMethodValues string

const (
	SortMethodValuesStroke SortMethodValues = "stroke"
	SortMethodValuesPinyin SortMethodValues = "pinYin"
	SortMethodValuesNone   SortMethodValues = "none"
)

func (e SortMethodValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *SortMethodValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = SortMethodValues(attr.Value)
	return nil
}

// ProofingStateValues
type ProofingStateValues string

const (
	ProofingStateValuesClean ProofingStateValues = "clean"
	ProofingStateValuesDirty ProofingStateValues = "dirty"
)

func (e ProofingStateValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ProofingStateValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ProofingStateValues(attr.Value)
	return nil
}

// TableWidthUnitValues
type TableWidthUnitValues string

const (
	TableWidthUnitValuesNil  TableWidthUnitValues = "nil"
	TableWidthUnitValuesPct  TableWidthUnitValues = "pct"
	TableWidthUnitValuesDxa  TableWidthUnitValues = "dxa"
	TableWidthUnitValuesAuto TableWidthUnitValues = "auto"
)

func (e TableWidthUnitValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TableWidthUnitValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TableWidthUnitValues(attr.Value)
	return nil
}

// FrameLayoutValues
type FrameLayoutValues string

const (
	FrameLayoutValuesRows FrameLayoutValues = "rows"
	FrameLayoutValuesCols FrameLayoutValues = "cols"
	FrameLayoutValuesNone FrameLayoutValues = "none"
)

func (e FrameLayoutValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FrameLayoutValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FrameLayoutValues(attr.Value)
	return nil
}

// TextEffectValues
type TextEffectValues string

const (
	TextEffectValuesBlinkbackground TextEffectValues = "blinkBackground"
	TextEffectValuesLights          TextEffectValues = "lights"
	TextEffectValuesAntsblack       TextEffectValues = "antsBlack"
	TextEffectValuesAntsred         TextEffectValues = "antsRed"
	TextEffectValuesShimmer         TextEffectValues = "shimmer"
	TextEffectValuesSparkle         TextEffectValues = "sparkle"
	TextEffectValuesNone            TextEffectValues = "none"
)

func (e TextEffectValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextEffectValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextEffectValues(attr.Value)
	return nil
}

// FontCollectionIndexValues
type FontCollectionIndexValues string

const (
	FontCollectionIndexValuesMajor FontCollectionIndexValues = "major"
	FontCollectionIndexValuesMinor FontCollectionIndexValues = "minor"
	FontCollectionIndexValuesNone  FontCollectionIndexValues = "none"
)

func (e FontCollectionIndexValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FontCollectionIndexValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FontCollectionIndexValues(attr.Value)
	return nil
}

// IconSetValues
type IconSetValues string

const (
	IconSetValuesIconSetValues3Arrows         IconSetValues = "3Arrows"
	IconSetValuesIconSetValues3Arrowsgray     IconSetValues = "3ArrowsGray"
	IconSetValuesIconSetValues3Flags          IconSetValues = "3Flags"
	IconSetValuesIconSetValues3Trafficlights1 IconSetValues = "3TrafficLights1"
	IconSetValuesIconSetValues3Trafficlights2 IconSetValues = "3TrafficLights2"
	IconSetValuesIconSetValues3Signs          IconSetValues = "3Signs"
	IconSetValuesIconSetValues3Symbols        IconSetValues = "3Symbols"
	IconSetValuesIconSetValues3Symbols2       IconSetValues = "3Symbols2"
	IconSetValuesIconSetValues4Arrows         IconSetValues = "4Arrows"
	IconSetValuesIconSetValues4Arrowsgray     IconSetValues = "4ArrowsGray"
	IconSetValuesIconSetValues4Redtoblack     IconSetValues = "4RedToBlack"
	IconSetValuesIconSetValues4Rating         IconSetValues = "4Rating"
	IconSetValuesIconSetValues4Trafficlights  IconSetValues = "4TrafficLights"
	IconSetValuesIconSetValues5Arrows         IconSetValues = "5Arrows"
	IconSetValuesIconSetValues5Arrowsgray     IconSetValues = "5ArrowsGray"
	IconSetValuesIconSetValues5Rating         IconSetValues = "5Rating"
	IconSetValuesIconSetValues5Quarters       IconSetValues = "5Quarters"
)

func (e IconSetValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *IconSetValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = IconSetValues(attr.Value)
	return nil
}

// ItemValues
type ItemValues string

const (
	ItemValuesData    ItemValues = "data"
	ItemValuesDefault ItemValues = "default"
	ItemValuesSum     ItemValues = "sum"
	ItemValuesCounta  ItemValues = "countA"
	ItemValuesAvg     ItemValues = "avg"
	ItemValuesMax     ItemValues = "max"
	ItemValuesMin     ItemValues = "min"
	ItemValuesProduct ItemValues = "product"
	ItemValuesCount   ItemValues = "count"
	ItemValuesStddev  ItemValues = "stdDev"
	ItemValuesStddevp ItemValues = "stdDevP"
	ItemValuesVar     ItemValues = "var"
	ItemValuesVarp    ItemValues = "varP"
	ItemValuesGrand   ItemValues = "grand"
	ItemValuesBlank   ItemValues = "blank"
)

func (e ItemValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ItemValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ItemValues(attr.Value)
	return nil
}

// FormulaExpressionValues
type FormulaExpressionValues string

const (
	FormulaExpressionValuesRef          FormulaExpressionValues = "ref"
	FormulaExpressionValuesReferror     FormulaExpressionValues = "refError"
	FormulaExpressionValuesArea         FormulaExpressionValues = "area"
	FormulaExpressionValuesAreaerror    FormulaExpressionValues = "areaError"
	FormulaExpressionValuesComputedarea FormulaExpressionValues = "computedArea"
)

func (e FormulaExpressionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *FormulaExpressionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = FormulaExpressionValues(attr.Value)
	return nil
}

// BorderStyleValues
type BorderStyleValues string

const (
	BorderStyleValuesNone             BorderStyleValues = "none"
	BorderStyleValuesThin             BorderStyleValues = "thin"
	BorderStyleValuesMedium           BorderStyleValues = "medium"
	BorderStyleValuesDashed           BorderStyleValues = "dashed"
	BorderStyleValuesDotted           BorderStyleValues = "dotted"
	BorderStyleValuesThick            BorderStyleValues = "thick"
	BorderStyleValuesDouble           BorderStyleValues = "double"
	BorderStyleValuesHair             BorderStyleValues = "hair"
	BorderStyleValuesMediumdashed     BorderStyleValues = "mediumDashed"
	BorderStyleValuesDashdot          BorderStyleValues = "dashDot"
	BorderStyleValuesMediumdashdot    BorderStyleValues = "mediumDashDot"
	BorderStyleValuesDashdotdot       BorderStyleValues = "dashDotDot"
	BorderStyleValuesMediumdashdotdot BorderStyleValues = "mediumDashDotDot"
	BorderStyleValuesSlantdashdot     BorderStyleValues = "slantDashDot"
)

func (e BorderStyleValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BorderStyleValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BorderStyleValues(attr.Value)
	return nil
}

// EditValidationValues
type EditValidationValues string

const (
	EditValidationValuesText      EditValidationValues = "text"
	EditValidationValuesInteger   EditValidationValues = "integer"
	EditValidationValuesNumber    EditValidationValues = "number"
	EditValidationValuesReference EditValidationValues = "reference"
	EditValidationValuesFormula   EditValidationValues = "formula"
)

func (e EditValidationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *EditValidationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = EditValidationValues(attr.Value)
	return nil
}

// DateFormatValues
type DateFormatValues string

const (
	DateFormatValuesText     DateFormatValues = "text"
	DateFormatValuesDate     DateFormatValues = "date"
	DateFormatValuesDatetime DateFormatValues = "dateTime"
)

func (e DateFormatValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DateFormatValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DateFormatValues(attr.Value)
	return nil
}

// MailMergeSourceValues
type MailMergeSourceValues string

const (
	MailMergeSourceValuesDatabase    MailMergeSourceValues = "database"
	MailMergeSourceValuesAddressbook MailMergeSourceValues = "addressBook"
	MailMergeSourceValuesDocument1   MailMergeSourceValues = "document1"
	MailMergeSourceValuesDocument2   MailMergeSourceValues = "document2"
	MailMergeSourceValuesText        MailMergeSourceValues = "text"
	MailMergeSourceValuesEmail       MailMergeSourceValues = "email"
	MailMergeSourceValuesNative      MailMergeSourceValues = "native"
	MailMergeSourceValuesLegacy      MailMergeSourceValues = "legacy"
	MailMergeSourceValuesMaster      MailMergeSourceValues = "master"
)

func (e MailMergeSourceValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *MailMergeSourceValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = MailMergeSourceValues(attr.Value)
	return nil
}

// DisplayBlanksAsValues
type DisplayBlanksAsValues string

const (
	DisplayBlanksAsValuesSpan DisplayBlanksAsValues = "span"
	DisplayBlanksAsValuesGap  DisplayBlanksAsValues = "gap"
	DisplayBlanksAsValuesZero DisplayBlanksAsValues = "zero"
)

func (e DisplayBlanksAsValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DisplayBlanksAsValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DisplayBlanksAsValues(attr.Value)
	return nil
}

// DiagramBuildStepValues
type DiagramBuildStepValues string

const (
	DiagramBuildStepValuesSp DiagramBuildStepValues = "sp"
	DiagramBuildStepValuesBg DiagramBuildStepValues = "bg"
)

func (e DiagramBuildStepValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *DiagramBuildStepValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = DiagramBuildStepValues(attr.Value)
	return nil
}

// PreviousActionValues
type PreviousActionValues string

const (
	PreviousActionValuesNone      PreviousActionValues = "none"
	PreviousActionValuesSkiptimed PreviousActionValues = "skipTimed"
)

func (e PreviousActionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PreviousActionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PreviousActionValues(attr.Value)
	return nil
}

// AnimateBehaviorValues
type AnimateBehaviorValues string

const (
	AnimateBehaviorValuesStr AnimateBehaviorValues = "str"
	AnimateBehaviorValuesNum AnimateBehaviorValues = "num"
	AnimateBehaviorValuesClr AnimateBehaviorValues = "clr"
)

func (e AnimateBehaviorValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AnimateBehaviorValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AnimateBehaviorValues(attr.Value)
	return nil
}

// TimeNodeMasterRelationValues
type TimeNodeMasterRelationValues string

const (
	TimeNodeMasterRelationValuesSameclick TimeNodeMasterRelationValues = "sameClick"
	TimeNodeMasterRelationValuesNextclick TimeNodeMasterRelationValues = "nextClick"
)

func (e TimeNodeMasterRelationValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TimeNodeMasterRelationValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TimeNodeMasterRelationValues(attr.Value)
	return nil
}

// ScopeValues
type ScopeValues string

const (
	ScopeValuesSelection ScopeValues = "selection"
	ScopeValuesData      ScopeValues = "data"
	ScopeValuesField     ScopeValues = "field"
)

func (e ScopeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ScopeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ScopeValues(attr.Value)
	return nil
}

// AbsolutePositionTabAlignmentValues
type AbsolutePositionTabAlignmentValues string

const (
	AbsolutePositionTabAlignmentValuesLeft   AbsolutePositionTabAlignmentValues = "left"
	AbsolutePositionTabAlignmentValuesCenter AbsolutePositionTabAlignmentValues = "center"
	AbsolutePositionTabAlignmentValuesRight  AbsolutePositionTabAlignmentValues = "right"
)

func (e AbsolutePositionTabAlignmentValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AbsolutePositionTabAlignmentValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AbsolutePositionTabAlignmentValues(
		attr.Value,
	)
	return nil
}

// ShadingPatternValues
type ShadingPatternValues string

const (
	ShadingPatternValuesNil                   ShadingPatternValues = "nil"
	ShadingPatternValuesClear                 ShadingPatternValues = "clear"
	ShadingPatternValuesSolid                 ShadingPatternValues = "solid"
	ShadingPatternValuesHorzstripe            ShadingPatternValues = "horzStripe"
	ShadingPatternValuesVertstripe            ShadingPatternValues = "vertStripe"
	ShadingPatternValuesReversediagstripe     ShadingPatternValues = "reverseDiagStripe"
	ShadingPatternValuesDiagstripe            ShadingPatternValues = "diagStripe"
	ShadingPatternValuesHorzcross             ShadingPatternValues = "horzCross"
	ShadingPatternValuesDiagcross             ShadingPatternValues = "diagCross"
	ShadingPatternValuesThinhorzstripe        ShadingPatternValues = "thinHorzStripe"
	ShadingPatternValuesThinvertstripe        ShadingPatternValues = "thinVertStripe"
	ShadingPatternValuesThinreversediagstripe ShadingPatternValues = "thinReverseDiagStripe"
	ShadingPatternValuesThindiagstripe        ShadingPatternValues = "thinDiagStripe"
	ShadingPatternValuesThinhorzcross         ShadingPatternValues = "thinHorzCross"
	ShadingPatternValuesThindiagcross         ShadingPatternValues = "thinDiagCross"
	ShadingPatternValuesPct5                  ShadingPatternValues = "pct5"
	ShadingPatternValuesPct10                 ShadingPatternValues = "pct10"
	ShadingPatternValuesPct12                 ShadingPatternValues = "pct12"
	ShadingPatternValuesPct15                 ShadingPatternValues = "pct15"
	ShadingPatternValuesPct20                 ShadingPatternValues = "pct20"
	ShadingPatternValuesPct25                 ShadingPatternValues = "pct25"
	ShadingPatternValuesPct30                 ShadingPatternValues = "pct30"
	ShadingPatternValuesPct35                 ShadingPatternValues = "pct35"
	ShadingPatternValuesPct37                 ShadingPatternValues = "pct37"
	ShadingPatternValuesPct40                 ShadingPatternValues = "pct40"
	ShadingPatternValuesPct45                 ShadingPatternValues = "pct45"
	ShadingPatternValuesPct50                 ShadingPatternValues = "pct50"
	ShadingPatternValuesPct55                 ShadingPatternValues = "pct55"
	ShadingPatternValuesPct60                 ShadingPatternValues = "pct60"
	ShadingPatternValuesPct62                 ShadingPatternValues = "pct62"
	ShadingPatternValuesPct65                 ShadingPatternValues = "pct65"
	ShadingPatternValuesPct70                 ShadingPatternValues = "pct70"
	ShadingPatternValuesPct75                 ShadingPatternValues = "pct75"
	ShadingPatternValuesPct80                 ShadingPatternValues = "pct80"
	ShadingPatternValuesPct85                 ShadingPatternValues = "pct85"
	ShadingPatternValuesPct87                 ShadingPatternValues = "pct87"
	ShadingPatternValuesPct90                 ShadingPatternValues = "pct90"
	ShadingPatternValuesPct95                 ShadingPatternValues = "pct95"
)

func (e ShadingPatternValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ShadingPatternValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ShadingPatternValues(attr.Value)
	return nil
}

// OleObjectFollowColorSchemeValues
type OleObjectFollowColorSchemeValues string

const (
	OleObjectFollowColorSchemeValuesNone              OleObjectFollowColorSchemeValues = "none"
	OleObjectFollowColorSchemeValuesFull              OleObjectFollowColorSchemeValues = "full"
	OleObjectFollowColorSchemeValuesTextandbackground OleObjectFollowColorSchemeValues = "textAndBackground"
)

func (e OleObjectFollowColorSchemeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *OleObjectFollowColorSchemeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = OleObjectFollowColorSchemeValues(
		attr.Value,
	)
	return nil
}

// CellCommentsValues
type CellCommentsValues string

const (
	CellCommentsValuesNone        CellCommentsValues = "none"
	CellCommentsValuesAsdisplayed CellCommentsValues = "asDisplayed"
	CellCommentsValuesAtend       CellCommentsValues = "atEnd"
)

func (e CellCommentsValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CellCommentsValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CellCommentsValues(attr.Value)
	return nil
}

// VolatileDependencyValues
type VolatileDependencyValues string

const (
	VolatileDependencyValuesRealtimedata  VolatileDependencyValues = "realTimeData"
	VolatileDependencyValuesOlapfunctions VolatileDependencyValues = "olapFunctions"
)

func (e VolatileDependencyValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *VolatileDependencyValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = VolatileDependencyValues(attr.Value)
	return nil
}

// CheckedValues
type CheckedValues string

const (
	CheckedValuesUnchecked CheckedValues = "Unchecked"
	CheckedValuesChecked   CheckedValues = "Checked"
	CheckedValuesMixed     CheckedValues = "Mixed"
)

func (e CheckedValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *CheckedValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = CheckedValues(attr.Value)
	return nil
}

// AnimationDiagramOnlyBuildValues
type AnimationDiagramOnlyBuildValues string

const (
	AnimationDiagramOnlyBuildValuesOne       AnimationDiagramOnlyBuildValues = "one"
	AnimationDiagramOnlyBuildValuesLvlone    AnimationDiagramOnlyBuildValues = "lvlOne"
	AnimationDiagramOnlyBuildValuesLvlatonce AnimationDiagramOnlyBuildValues = "lvlAtOnce"
)

func (e AnimationDiagramOnlyBuildValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AnimationDiagramOnlyBuildValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AnimationDiagramOnlyBuildValues(
		attr.Value,
	)
	return nil
}

// BlendModeValues
type BlendModeValues string

const (
	BlendModeValuesOver    BlendModeValues = "over"
	BlendModeValuesMult    BlendModeValues = "mult"
	BlendModeValuesScreen  BlendModeValues = "screen"
	BlendModeValuesDarken  BlendModeValues = "darken"
	BlendModeValuesLighten BlendModeValues = "lighten"
)

func (e BlendModeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *BlendModeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = BlendModeValues(attr.Value)
	return nil
}

// AnimateEffectTransitionValues
type AnimateEffectTransitionValues string

const (
	AnimateEffectTransitionValuesIn   AnimateEffectTransitionValues = "in"
	AnimateEffectTransitionValuesOut  AnimateEffectTransitionValues = "out"
	AnimateEffectTransitionValuesNone AnimateEffectTransitionValues = "none"
)

func (e AnimateEffectTransitionValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *AnimateEffectTransitionValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = AnimateEffectTransitionValues(attr.Value)
	return nil
}

// NumberFormatValues
type NumberFormatValues string

const (
	NumberFormatValuesDecimal                      NumberFormatValues = "decimal"
	NumberFormatValuesUpperroman                   NumberFormatValues = "upperRoman"
	NumberFormatValuesLowerroman                   NumberFormatValues = "lowerRoman"
	NumberFormatValuesUpperletter                  NumberFormatValues = "upperLetter"
	NumberFormatValuesLowerletter                  NumberFormatValues = "lowerLetter"
	NumberFormatValuesOrdinal                      NumberFormatValues = "ordinal"
	NumberFormatValuesCardinaltext                 NumberFormatValues = "cardinalText"
	NumberFormatValuesOrdinaltext                  NumberFormatValues = "ordinalText"
	NumberFormatValuesHex                          NumberFormatValues = "hex"
	NumberFormatValuesChicago                      NumberFormatValues = "chicago"
	NumberFormatValuesIdeographdigital             NumberFormatValues = "ideographDigital"
	NumberFormatValuesJapanesecounting             NumberFormatValues = "japaneseCounting"
	NumberFormatValuesAiueo                        NumberFormatValues = "aiueo"
	NumberFormatValuesIroha                        NumberFormatValues = "iroha"
	NumberFormatValuesDecimalfullwidth             NumberFormatValues = "decimalFullWidth"
	NumberFormatValuesDecimalhalfwidth             NumberFormatValues = "decimalHalfWidth"
	NumberFormatValuesJapaneselegal                NumberFormatValues = "japaneseLegal"
	NumberFormatValuesJapanesedigitaltenthousand   NumberFormatValues = "japaneseDigitalTenThousand"
	NumberFormatValuesDecimalenclosedcircle        NumberFormatValues = "decimalEnclosedCircle"
	NumberFormatValuesDecimalfullwidth2            NumberFormatValues = "decimalFullWidth2"
	NumberFormatValuesAiueofullwidth               NumberFormatValues = "aiueoFullWidth"
	NumberFormatValuesIrohafullwidth               NumberFormatValues = "irohaFullWidth"
	NumberFormatValuesDecimalzero                  NumberFormatValues = "decimalZero"
	NumberFormatValuesBullet                       NumberFormatValues = "bullet"
	NumberFormatValuesGanada                       NumberFormatValues = "ganada"
	NumberFormatValuesChosung                      NumberFormatValues = "chosung"
	NumberFormatValuesDecimalenclosedfullstop      NumberFormatValues = "decimalEnclosedFullstop"
	NumberFormatValuesDecimalenclosedparen         NumberFormatValues = "decimalEnclosedParen"
	NumberFormatValuesDecimalenclosedcirclechinese NumberFormatValues = "decimalEnclosedCircleChinese"
	NumberFormatValuesIdeographenclosedcircle      NumberFormatValues = "ideographEnclosedCircle"
	NumberFormatValuesIdeographtraditional         NumberFormatValues = "ideographTraditional"
	NumberFormatValuesIdeographzodiac              NumberFormatValues = "ideographZodiac"
	NumberFormatValuesIdeographzodiactraditional   NumberFormatValues = "ideographZodiacTraditional"
	NumberFormatValuesTaiwanesecounting            NumberFormatValues = "taiwaneseCounting"
	NumberFormatValuesIdeographlegaltraditional    NumberFormatValues = "ideographLegalTraditional"
	NumberFormatValuesTaiwanesecountingthousand    NumberFormatValues = "taiwaneseCountingThousand"
	NumberFormatValuesTaiwanesedigital             NumberFormatValues = "taiwaneseDigital"
	NumberFormatValuesChinesecounting              NumberFormatValues = "chineseCounting"
	NumberFormatValuesChineselegalsimplified       NumberFormatValues = "chineseLegalSimplified"
	NumberFormatValuesChinesecountingthousand      NumberFormatValues = "chineseCountingThousand"
	NumberFormatValuesKoreandigital                NumberFormatValues = "koreanDigital"
	NumberFormatValuesKoreancounting               NumberFormatValues = "koreanCounting"
	NumberFormatValuesKoreanlegal                  NumberFormatValues = "koreanLegal"
	NumberFormatValuesKoreandigital2               NumberFormatValues = "koreanDigital2"
	NumberFormatValuesVietnamesecounting           NumberFormatValues = "vietnameseCounting"
	NumberFormatValuesRussianlower                 NumberFormatValues = "russianLower"
	NumberFormatValuesRussianupper                 NumberFormatValues = "russianUpper"
	NumberFormatValuesNone                         NumberFormatValues = "none"
	NumberFormatValuesNumberindash                 NumberFormatValues = "numberInDash"
	NumberFormatValuesHebrew1                      NumberFormatValues = "hebrew1"
	NumberFormatValuesHebrew2                      NumberFormatValues = "hebrew2"
	NumberFormatValuesArabicalpha                  NumberFormatValues = "arabicAlpha"
	NumberFormatValuesArabicabjad                  NumberFormatValues = "arabicAbjad"
	NumberFormatValuesHindivowels                  NumberFormatValues = "hindiVowels"
	NumberFormatValuesHindiconsonants              NumberFormatValues = "hindiConsonants"
	NumberFormatValuesHindinumbers                 NumberFormatValues = "hindiNumbers"
	NumberFormatValuesHindicounting                NumberFormatValues = "hindiCounting"
	NumberFormatValuesThailetters                  NumberFormatValues = "thaiLetters"
	NumberFormatValuesThainumbers                  NumberFormatValues = "thaiNumbers"
	NumberFormatValuesThaicounting                 NumberFormatValues = "thaiCounting"
	NumberFormatValuesBahttext                     NumberFormatValues = "bahtText"
	NumberFormatValuesDollartext                   NumberFormatValues = "dollarText"
	NumberFormatValuesCustom                       NumberFormatValues = "custom"
)

func (e NumberFormatValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *NumberFormatValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = NumberFormatValues(attr.Value)
	return nil
}

// ParagraphBuildValues
type ParagraphBuildValues string

const (
	ParagraphBuildValuesAllatonce ParagraphBuildValues = "allAtOnce"
	ParagraphBuildValuesP         ParagraphBuildValues = "p"
	ParagraphBuildValuesCust      ParagraphBuildValues = "cust"
	ParagraphBuildValuesWhole     ParagraphBuildValues = "whole"
)

func (e ParagraphBuildValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *ParagraphBuildValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = ParagraphBuildValues(attr.Value)
	return nil
}

// PageOrderValues
type PageOrderValues string

const (
	PageOrderValuesDownthenover PageOrderValues = "downThenOver"
	PageOrderValuesOverthendown PageOrderValues = "overThenDown"
)

func (e PageOrderValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PageOrderValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PageOrderValues(attr.Value)
	return nil
}

// PathShadeValues
type PathShadeValues string

const (
	PathShadeValuesShape  PathShadeValues = "shape"
	PathShadeValuesCircle PathShadeValues = "circle"
	PathShadeValuesRect   PathShadeValues = "rect"
)

func (e PathShadeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *PathShadeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = PathShadeValues(attr.Value)
	return nil
}

// TextAutoNumberSchemeValues
type TextAutoNumberSchemeValues string

const (
	TextAutoNumberSchemeValuesAlphalcparenboth      TextAutoNumberSchemeValues = "alphaLcParenBoth"
	TextAutoNumberSchemeValuesAlphaucparenboth      TextAutoNumberSchemeValues = "alphaUcParenBoth"
	TextAutoNumberSchemeValuesAlphalcparenr         TextAutoNumberSchemeValues = "alphaLcParenR"
	TextAutoNumberSchemeValuesAlphaucparenr         TextAutoNumberSchemeValues = "alphaUcParenR"
	TextAutoNumberSchemeValuesAlphalcperiod         TextAutoNumberSchemeValues = "alphaLcPeriod"
	TextAutoNumberSchemeValuesAlphaucperiod         TextAutoNumberSchemeValues = "alphaUcPeriod"
	TextAutoNumberSchemeValuesArabicparenboth       TextAutoNumberSchemeValues = "arabicParenBoth"
	TextAutoNumberSchemeValuesArabicparenr          TextAutoNumberSchemeValues = "arabicParenR"
	TextAutoNumberSchemeValuesArabicperiod          TextAutoNumberSchemeValues = "arabicPeriod"
	TextAutoNumberSchemeValuesArabicplain           TextAutoNumberSchemeValues = "arabicPlain"
	TextAutoNumberSchemeValuesRomanlcparenboth      TextAutoNumberSchemeValues = "romanLcParenBoth"
	TextAutoNumberSchemeValuesRomanucparenboth      TextAutoNumberSchemeValues = "romanUcParenBoth"
	TextAutoNumberSchemeValuesRomanlcparenr         TextAutoNumberSchemeValues = "romanLcParenR"
	TextAutoNumberSchemeValuesRomanucparenr         TextAutoNumberSchemeValues = "romanUcParenR"
	TextAutoNumberSchemeValuesRomanlcperiod         TextAutoNumberSchemeValues = "romanLcPeriod"
	TextAutoNumberSchemeValuesRomanucperiod         TextAutoNumberSchemeValues = "romanUcPeriod"
	TextAutoNumberSchemeValuesCirclenumdbplain      TextAutoNumberSchemeValues = "circleNumDbPlain"
	TextAutoNumberSchemeValuesCirclenumwdblackplain TextAutoNumberSchemeValues = "circleNumWdBlackPlain"
	TextAutoNumberSchemeValuesCirclenumwdwhiteplain TextAutoNumberSchemeValues = "circleNumWdWhitePlain"
	TextAutoNumberSchemeValuesArabicdbperiod        TextAutoNumberSchemeValues = "arabicDbPeriod"
	TextAutoNumberSchemeValuesArabicdbplain         TextAutoNumberSchemeValues = "arabicDbPlain"
	TextAutoNumberSchemeValuesEa1chsperiod          TextAutoNumberSchemeValues = "ea1ChsPeriod"
	TextAutoNumberSchemeValuesEa1chsplain           TextAutoNumberSchemeValues = "ea1ChsPlain"
	TextAutoNumberSchemeValuesEa1chtperiod          TextAutoNumberSchemeValues = "ea1ChtPeriod"
	TextAutoNumberSchemeValuesEa1chtplain           TextAutoNumberSchemeValues = "ea1ChtPlain"
	TextAutoNumberSchemeValuesEa1jpnchsdbperiod     TextAutoNumberSchemeValues = "ea1JpnChsDbPeriod"
	TextAutoNumberSchemeValuesEa1jpnkorplain        TextAutoNumberSchemeValues = "ea1JpnKorPlain"
	TextAutoNumberSchemeValuesEa1jpnkorperiod       TextAutoNumberSchemeValues = "ea1JpnKorPeriod"
	TextAutoNumberSchemeValuesArabic1minus          TextAutoNumberSchemeValues = "arabic1Minus"
	TextAutoNumberSchemeValuesArabic2minus          TextAutoNumberSchemeValues = "arabic2Minus"
	TextAutoNumberSchemeValuesHebrew2minus          TextAutoNumberSchemeValues = "hebrew2Minus"
	TextAutoNumberSchemeValuesThaialphaperiod       TextAutoNumberSchemeValues = "thaiAlphaPeriod"
	TextAutoNumberSchemeValuesThaialphaparenr       TextAutoNumberSchemeValues = "thaiAlphaParenR"
	TextAutoNumberSchemeValuesThaialphaparenboth    TextAutoNumberSchemeValues = "thaiAlphaParenBoth"
	TextAutoNumberSchemeValuesThainumperiod         TextAutoNumberSchemeValues = "thaiNumPeriod"
	TextAutoNumberSchemeValuesThainumparenr         TextAutoNumberSchemeValues = "thaiNumParenR"
	TextAutoNumberSchemeValuesThainumparenboth      TextAutoNumberSchemeValues = "thaiNumParenBoth"
	TextAutoNumberSchemeValuesHindialphaperiod      TextAutoNumberSchemeValues = "hindiAlphaPeriod"
	TextAutoNumberSchemeValuesHindinumperiod        TextAutoNumberSchemeValues = "hindiNumPeriod"
	TextAutoNumberSchemeValuesHindinumparenr        TextAutoNumberSchemeValues = "hindiNumParenR"
	TextAutoNumberSchemeValuesHindialpha1period     TextAutoNumberSchemeValues = "hindiAlpha1Period"
)

func (e TextAutoNumberSchemeValues) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TextAutoNumberSchemeValues) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TextAutoNumberSchemeValues(attr.Value)
	return nil
}
