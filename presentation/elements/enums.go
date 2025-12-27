// nolint
//
//revive:disable:file-length-limit,max-public-structs,line-length-limit,comments-density // Generated code
package elements

import (
	"encoding/xml"
)

// TextBoxTightWrapValues
type TextBoxTightWrapValues string

const (
	TextBoxTightWrapValuesNone             TextBoxTightWrapValues = "none"
	TextBoxTightWrapValuesAllLines         TextBoxTightWrapValues = "allLines"
	TextBoxTightWrapValuesFirstAndLastLine TextBoxTightWrapValues = "firstAndLastLine"
	TextBoxTightWrapValuesFirstLineOnly    TextBoxTightWrapValues = "firstLineOnly"
	TextBoxTightWrapValuesLastLineOnly     TextBoxTightWrapValues = "lastLineOnly"
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

// InfoTextValues
type InfoTextValues string

const (
	InfoTextValuesText     InfoTextValues = "text"
	InfoTextValuesAutoText InfoTextValues = "autoText"
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
	BorderValuesDotDash                BorderValues = "dotDash"
	BorderValuesDotDotDash             BorderValues = "dotDotDash"
	BorderValuesTriple                 BorderValues = "triple"
	BorderValuesThinThickSmallGap      BorderValues = "thinThickSmallGap"
	BorderValuesThickThinSmallGap      BorderValues = "thickThinSmallGap"
	BorderValuesThinThickThinSmallGap  BorderValues = "thinThickThinSmallGap"
	BorderValuesThinThickMediumGap     BorderValues = "thinThickMediumGap"
	BorderValuesThickThinMediumGap     BorderValues = "thickThinMediumGap"
	BorderValuesThinThickThinMediumGap BorderValues = "thinThickThinMediumGap"
	BorderValuesThinThickLargeGap      BorderValues = "thinThickLargeGap"
	BorderValuesThickThinLargeGap      BorderValues = "thickThinLargeGap"
	BorderValuesThinThickThinLargeGap  BorderValues = "thinThickThinLargeGap"
	BorderValuesWave                   BorderValues = "wave"
	BorderValuesDoubleWave             BorderValues = "doubleWave"
	BorderValuesDashSmallGap           BorderValues = "dashSmallGap"
	BorderValuesDashDotStroked         BorderValues = "dashDotStroked"
	BorderValuesThreeDEmboss           BorderValues = "threeDEmboss"
	BorderValuesThreeDEngrave          BorderValues = "threeDEngrave"
	BorderValuesOutset                 BorderValues = "outset"
	BorderValuesInset                  BorderValues = "inset"
	BorderValuesApples                 BorderValues = "apples"
	BorderValuesArchedScallops         BorderValues = "archedScallops"
	BorderValuesBabyPacifier           BorderValues = "babyPacifier"
	BorderValuesBabyRattle             BorderValues = "babyRattle"
	BorderValuesBalloons3Colors        BorderValues = "balloons3Colors"
	BorderValuesBalloonsHotAir         BorderValues = "balloonsHotAir"
	BorderValuesBasicBlackDashes       BorderValues = "basicBlackDashes"
	BorderValuesBasicBlackDots         BorderValues = "basicBlackDots"
	BorderValuesBasicBlackSquares      BorderValues = "basicBlackSquares"
	BorderValuesBasicThinLines         BorderValues = "basicThinLines"
	BorderValuesBasicWhiteDashes       BorderValues = "basicWhiteDashes"
	BorderValuesBasicWhiteDots         BorderValues = "basicWhiteDots"
	BorderValuesBasicWhiteSquares      BorderValues = "basicWhiteSquares"
	BorderValuesBasicWideInline        BorderValues = "basicWideInline"
	BorderValuesBasicWideMidline       BorderValues = "basicWideMidline"
	BorderValuesBasicWideOutline       BorderValues = "basicWideOutline"
	BorderValuesBats                   BorderValues = "bats"
	BorderValuesBirds                  BorderValues = "birds"
	BorderValuesBirdsFlight            BorderValues = "birdsFlight"
	BorderValuesCabins                 BorderValues = "cabins"
	BorderValuesCakeSlice              BorderValues = "cakeSlice"
	BorderValuesCandyCorn              BorderValues = "candyCorn"
	BorderValuesCelticKnotwork         BorderValues = "celticKnotwork"
	BorderValuesCertificateBanner      BorderValues = "certificateBanner"
	BorderValuesChainLink              BorderValues = "chainLink"
	BorderValuesChampagneBottle        BorderValues = "champagneBottle"
	BorderValuesCheckedBarBlack        BorderValues = "checkedBarBlack"
	BorderValuesCheckedBarColor        BorderValues = "checkedBarColor"
	BorderValuesCheckered              BorderValues = "checkered"
	BorderValuesChristmasTree          BorderValues = "christmasTree"
	BorderValuesCirclesLines           BorderValues = "circlesLines"
	BorderValuesCirclesRectangles      BorderValues = "circlesRectangles"
	BorderValuesClassicalWave          BorderValues = "classicalWave"
	BorderValuesClocks                 BorderValues = "clocks"
	BorderValuesCompass                BorderValues = "compass"
	BorderValuesConfetti               BorderValues = "confetti"
	BorderValuesConfettiGrays          BorderValues = "confettiGrays"
	BorderValuesConfettiOutline        BorderValues = "confettiOutline"
	BorderValuesConfettiStreamers      BorderValues = "confettiStreamers"
	BorderValuesConfettiWhite          BorderValues = "confettiWhite"
	BorderValuesCornerTriangles        BorderValues = "cornerTriangles"
	BorderValuesCouponCutoutDashes     BorderValues = "couponCutoutDashes"
	BorderValuesCouponCutoutDots       BorderValues = "couponCutoutDots"
	BorderValuesCrazyMaze              BorderValues = "crazyMaze"
	BorderValuesCreaturesButterfly     BorderValues = "creaturesButterfly"
	BorderValuesCreaturesFish          BorderValues = "creaturesFish"
	BorderValuesCreaturesInsects       BorderValues = "creaturesInsects"
	BorderValuesCreaturesLadyBug       BorderValues = "creaturesLadyBug"
	BorderValuesCrossStitch            BorderValues = "crossStitch"
	BorderValuesCup                    BorderValues = "cup"
	BorderValuesDecoArch               BorderValues = "decoArch"
	BorderValuesDecoArchColor          BorderValues = "decoArchColor"
	BorderValuesDecoBlocks             BorderValues = "decoBlocks"
	BorderValuesDiamondsGray           BorderValues = "diamondsGray"
	BorderValuesDoubleD                BorderValues = "doubleD"
	BorderValuesDoubleDiamonds         BorderValues = "doubleDiamonds"
	BorderValuesEarth1                 BorderValues = "earth1"
	BorderValuesEarth2                 BorderValues = "earth2"
	BorderValuesEclipsingSquares1      BorderValues = "eclipsingSquares1"
	BorderValuesEclipsingSquares2      BorderValues = "eclipsingSquares2"
	BorderValuesEggsBlack              BorderValues = "eggsBlack"
	BorderValuesFans                   BorderValues = "fans"
	BorderValuesFilm                   BorderValues = "film"
	BorderValuesFirecrackers           BorderValues = "firecrackers"
	BorderValuesFlowersBlockPrint      BorderValues = "flowersBlockPrint"
	BorderValuesFlowersDaisies         BorderValues = "flowersDaisies"
	BorderValuesFlowersModern1         BorderValues = "flowersModern1"
	BorderValuesFlowersModern2         BorderValues = "flowersModern2"
	BorderValuesFlowersPansy           BorderValues = "flowersPansy"
	BorderValuesFlowersRedRose         BorderValues = "flowersRedRose"
	BorderValuesFlowersRoses           BorderValues = "flowersRoses"
	BorderValuesFlowersTeacup          BorderValues = "flowersTeacup"
	BorderValuesFlowersTiny            BorderValues = "flowersTiny"
	BorderValuesGems                   BorderValues = "gems"
	BorderValuesGingerbreadMan         BorderValues = "gingerbreadMan"
	BorderValuesGradient               BorderValues = "gradient"
	BorderValuesHandmade1              BorderValues = "handmade1"
	BorderValuesHandmade2              BorderValues = "handmade2"
	BorderValuesHeartBalloon           BorderValues = "heartBalloon"
	BorderValuesHeartGray              BorderValues = "heartGray"
	BorderValuesHearts                 BorderValues = "hearts"
	BorderValuesHeebieJeebies          BorderValues = "heebieJeebies"
	BorderValuesHolly                  BorderValues = "holly"
	BorderValuesHouseFunky             BorderValues = "houseFunky"
	BorderValuesHypnotic               BorderValues = "hypnotic"
	BorderValuesIceCreamCones          BorderValues = "iceCreamCones"
	BorderValuesLightBulb              BorderValues = "lightBulb"
	BorderValuesLightning1             BorderValues = "lightning1"
	BorderValuesLightning2             BorderValues = "lightning2"
	BorderValuesMapPins                BorderValues = "mapPins"
	BorderValuesMapleLeaf              BorderValues = "mapleLeaf"
	BorderValuesMapleMuffins           BorderValues = "mapleMuffins"
	BorderValuesMarquee                BorderValues = "marquee"
	BorderValuesMarqueeToothed         BorderValues = "marqueeToothed"
	BorderValuesMoons                  BorderValues = "moons"
	BorderValuesMosaic                 BorderValues = "mosaic"
	BorderValuesMusicNotes             BorderValues = "musicNotes"
	BorderValuesNorthwest              BorderValues = "northwest"
	BorderValuesOvals                  BorderValues = "ovals"
	BorderValuesPackages               BorderValues = "packages"
	BorderValuesPalmsBlack             BorderValues = "palmsBlack"
	BorderValuesPalmsColor             BorderValues = "palmsColor"
	BorderValuesPaperClips             BorderValues = "paperClips"
	BorderValuesPapyrus                BorderValues = "papyrus"
	BorderValuesPartyFavor             BorderValues = "partyFavor"
	BorderValuesPartyGlass             BorderValues = "partyGlass"
	BorderValuesPencils                BorderValues = "pencils"
	BorderValuesPeople                 BorderValues = "people"
	BorderValuesPeopleWaving           BorderValues = "peopleWaving"
	BorderValuesPeopleHats             BorderValues = "peopleHats"
	BorderValuesPoinsettias            BorderValues = "poinsettias"
	BorderValuesPostageStamp           BorderValues = "postageStamp"
	BorderValuesPumpkin1               BorderValues = "pumpkin1"
	BorderValuesPushPinNote2           BorderValues = "pushPinNote2"
	BorderValuesPushPinNote1           BorderValues = "pushPinNote1"
	BorderValuesPyramids               BorderValues = "pyramids"
	BorderValuesPyramidsAbove          BorderValues = "pyramidsAbove"
	BorderValuesQuadrants              BorderValues = "quadrants"
	BorderValuesRings                  BorderValues = "rings"
	BorderValuesSafari                 BorderValues = "safari"
	BorderValuesSawtooth               BorderValues = "sawtooth"
	BorderValuesSawtoothGray           BorderValues = "sawtoothGray"
	BorderValuesScaredCat              BorderValues = "scaredCat"
	BorderValuesSeattle                BorderValues = "seattle"
	BorderValuesShadowedSquares        BorderValues = "shadowedSquares"
	BorderValuesSharksTeeth            BorderValues = "sharksTeeth"
	BorderValuesShorebirdTracks        BorderValues = "shorebirdTracks"
	BorderValuesSkyrocket              BorderValues = "skyrocket"
	BorderValuesSnowflakeFancy         BorderValues = "snowflakeFancy"
	BorderValuesSnowflakes             BorderValues = "snowflakes"
	BorderValuesSombrero               BorderValues = "sombrero"
	BorderValuesSouthwest              BorderValues = "southwest"
	BorderValuesStars                  BorderValues = "stars"
	BorderValuesStarsTop               BorderValues = "starsTop"
	BorderValuesStars3d                BorderValues = "stars3d"
	BorderValuesStarsBlack             BorderValues = "starsBlack"
	BorderValuesStarsShadowed          BorderValues = "starsShadowed"
	BorderValuesSun                    BorderValues = "sun"
	BorderValuesSwirligig              BorderValues = "swirligig"
	BorderValuesTornPaper              BorderValues = "tornPaper"
	BorderValuesTornPaperBlack         BorderValues = "tornPaperBlack"
	BorderValuesTrees                  BorderValues = "trees"
	BorderValuesTriangleParty          BorderValues = "triangleParty"
	BorderValuesTriangles              BorderValues = "triangles"
	BorderValuesTribal1                BorderValues = "tribal1"
	BorderValuesTribal2                BorderValues = "tribal2"
	BorderValuesTribal3                BorderValues = "tribal3"
	BorderValuesTribal4                BorderValues = "tribal4"
	BorderValuesTribal5                BorderValues = "tribal5"
	BorderValuesTribal6                BorderValues = "tribal6"
	BorderValuesTriangle1              BorderValues = "triangle1"
	BorderValuesTriangle2              BorderValues = "triangle2"
	BorderValuesTriangleCircle1        BorderValues = "triangleCircle1"
	BorderValuesTriangleCircle2        BorderValues = "triangleCircle2"
	BorderValuesShapes1                BorderValues = "shapes1"
	BorderValuesShapes2                BorderValues = "shapes2"
	BorderValuesTwistedLines1          BorderValues = "twistedLines1"
	BorderValuesTwistedLines2          BorderValues = "twistedLines2"
	BorderValuesVine                   BorderValues = "vine"
	BorderValuesWaveline               BorderValues = "waveline"
	BorderValuesWeavingAngles          BorderValues = "weavingAngles"
	BorderValuesWeavingBraid           BorderValues = "weavingBraid"
	BorderValuesWeavingRibbon          BorderValues = "weavingRibbon"
	BorderValuesWeavingStrips          BorderValues = "weavingStrips"
	BorderValuesWhiteFlowers           BorderValues = "whiteFlowers"
	BorderValuesWoodwork               BorderValues = "woodwork"
	BorderValuesXIllusions             BorderValues = "xIllusions"
	BorderValuesZanyTriangles          BorderValues = "zanyTriangles"
	BorderValuesZigZag                 BorderValues = "zigZag"
	BorderValuesZigZagStitch           BorderValues = "zigZagStitch"
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

// CryptProviderValues
type CryptProviderValues string

const (
	CryptProviderValuesRsaAES  CryptProviderValues = "rsaAES"
	CryptProviderValuesRsaFull CryptProviderValues = "rsaFull"
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

// DynamicFilterValues
type DynamicFilterValues string

const (
	DynamicFilterValuesNull         DynamicFilterValues = "null"
	DynamicFilterValuesAboveAverage DynamicFilterValues = "aboveAverage"
	DynamicFilterValuesBelowAverage DynamicFilterValues = "belowAverage"
	DynamicFilterValuesTomorrow     DynamicFilterValues = "tomorrow"
	DynamicFilterValuesToday        DynamicFilterValues = "today"
	DynamicFilterValuesYesterday    DynamicFilterValues = "yesterday"
	DynamicFilterValuesNextWeek     DynamicFilterValues = "nextWeek"
	DynamicFilterValuesThisWeek     DynamicFilterValues = "thisWeek"
	DynamicFilterValuesLastWeek     DynamicFilterValues = "lastWeek"
	DynamicFilterValuesNextMonth    DynamicFilterValues = "nextMonth"
	DynamicFilterValuesThisMonth    DynamicFilterValues = "thisMonth"
	DynamicFilterValuesLastMonth    DynamicFilterValues = "lastMonth"
	DynamicFilterValuesNextQuarter  DynamicFilterValues = "nextQuarter"
	DynamicFilterValuesThisQuarter  DynamicFilterValues = "thisQuarter"
	DynamicFilterValuesLastQuarter  DynamicFilterValues = "lastQuarter"
	DynamicFilterValuesNextYear     DynamicFilterValues = "nextYear"
	DynamicFilterValuesThisYear     DynamicFilterValues = "thisYear"
	DynamicFilterValuesLastYear     DynamicFilterValues = "lastYear"
	DynamicFilterValuesYearToDate   DynamicFilterValues = "yearToDate"
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

// QuestionType
type QuestionType string

const (
	QuestionTypeCheckBox            QuestionType = "checkBox"
	QuestionTypeChoice              QuestionType = "choice"
	QuestionTypeDate                QuestionType = "date"
	QuestionTypeTime                QuestionType = "time"
	QuestionTypeMultipleLinesOfText QuestionType = "multipleLinesOfText"
	QuestionTypeNumber              QuestionType = "number"
	QuestionTypeSingleLineOfText    QuestionType = "singleLineOfText"
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

// SortValues
type SortValues string

const (
	SortValuesNone              SortValues = "none"
	SortValuesAscending         SortValues = "ascending"
	SortValuesDescending        SortValues = "descending"
	SortValuesAscendingAlpha    SortValues = "ascendingAlpha"
	SortValuesDescendingAlpha   SortValues = "descendingAlpha"
	SortValuesAscendingNatural  SortValues = "ascendingNatural"
	SortValuesDescendingNatural SortValues = "descendingNatural"
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

// PivotFilterValues
type PivotFilterValues string

const (
	PivotFilterValuesUnknown                   PivotFilterValues = "unknown"
	PivotFilterValuesCount                     PivotFilterValues = "count"
	PivotFilterValuesPercent                   PivotFilterValues = "percent"
	PivotFilterValuesSum                       PivotFilterValues = "sum"
	PivotFilterValuesCaptionEqual              PivotFilterValues = "captionEqual"
	PivotFilterValuesCaptionNotEqual           PivotFilterValues = "captionNotEqual"
	PivotFilterValuesCaptionBeginsWith         PivotFilterValues = "captionBeginsWith"
	PivotFilterValuesCaptionNotBeginsWith      PivotFilterValues = "captionNotBeginsWith"
	PivotFilterValuesCaptionEndsWith           PivotFilterValues = "captionEndsWith"
	PivotFilterValuesCaptionNotEndsWith        PivotFilterValues = "captionNotEndsWith"
	PivotFilterValuesCaptionContains           PivotFilterValues = "captionContains"
	PivotFilterValuesCaptionNotContains        PivotFilterValues = "captionNotContains"
	PivotFilterValuesCaptionGreaterThan        PivotFilterValues = "captionGreaterThan"
	PivotFilterValuesCaptionGreaterThanOrEqual PivotFilterValues = "captionGreaterThanOrEqual"
	PivotFilterValuesCaptionLessThan           PivotFilterValues = "captionLessThan"
	PivotFilterValuesCaptionLessThanOrEqual    PivotFilterValues = "captionLessThanOrEqual"
	PivotFilterValuesCaptionBetween            PivotFilterValues = "captionBetween"
	PivotFilterValuesCaptionNotBetween         PivotFilterValues = "captionNotBetween"
	PivotFilterValuesValueEqual                PivotFilterValues = "valueEqual"
	PivotFilterValuesValueNotEqual             PivotFilterValues = "valueNotEqual"
	PivotFilterValuesValueGreaterThan          PivotFilterValues = "valueGreaterThan"
	PivotFilterValuesValueGreaterThanOrEqual   PivotFilterValues = "valueGreaterThanOrEqual"
	PivotFilterValuesValueLessThan             PivotFilterValues = "valueLessThan"
	PivotFilterValuesValueLessThanOrEqual      PivotFilterValues = "valueLessThanOrEqual"
	PivotFilterValuesValueBetween              PivotFilterValues = "valueBetween"
	PivotFilterValuesValueNotBetween           PivotFilterValues = "valueNotBetween"
	PivotFilterValuesDateEqual                 PivotFilterValues = "dateEqual"
	PivotFilterValuesDateNotEqual              PivotFilterValues = "dateNotEqual"
	PivotFilterValuesDateOlderThan             PivotFilterValues = "dateOlderThan"
	PivotFilterValuesDateOlderThanOrEqual      PivotFilterValues = "dateOlderThanOrEqual"
	PivotFilterValuesDateNewerThan             PivotFilterValues = "dateNewerThan"
	PivotFilterValuesDateNewerThanOrEqual      PivotFilterValues = "dateNewerThanOrEqual"
	PivotFilterValuesDateBetween               PivotFilterValues = "dateBetween"
	PivotFilterValuesDateNotBetween            PivotFilterValues = "dateNotBetween"
	PivotFilterValuesTomorrow                  PivotFilterValues = "tomorrow"
	PivotFilterValuesToday                     PivotFilterValues = "today"
	PivotFilterValuesYesterday                 PivotFilterValues = "yesterday"
	PivotFilterValuesNextWeek                  PivotFilterValues = "nextWeek"
	PivotFilterValuesThisWeek                  PivotFilterValues = "thisWeek"
	PivotFilterValuesLastWeek                  PivotFilterValues = "lastWeek"
	PivotFilterValuesNextMonth                 PivotFilterValues = "nextMonth"
	PivotFilterValuesThisMonth                 PivotFilterValues = "thisMonth"
	PivotFilterValuesLastMonth                 PivotFilterValues = "lastMonth"
	PivotFilterValuesNextQuarter               PivotFilterValues = "nextQuarter"
	PivotFilterValuesThisQuarter               PivotFilterValues = "thisQuarter"
	PivotFilterValuesLastQuarter               PivotFilterValues = "lastQuarter"
	PivotFilterValuesNextYear                  PivotFilterValues = "nextYear"
	PivotFilterValuesThisYear                  PivotFilterValues = "thisYear"
	PivotFilterValuesLastYear                  PivotFilterValues = "lastYear"
	PivotFilterValuesYearToDate                PivotFilterValues = "yearToDate"
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

// LineSpacingRuleValues
type LineSpacingRuleValues string

const (
	LineSpacingRuleValuesAuto    LineSpacingRuleValues = "auto"
	LineSpacingRuleValuesExact   LineSpacingRuleValues = "exact"
	LineSpacingRuleValuesAtLeast LineSpacingRuleValues = "atLeast"
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
	ColorSchemeIndexValuesFollowedHyperlink ColorSchemeIndexValues = "followedHyperlink"
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

// AnimationChartOnlyBuildValues
type AnimationChartOnlyBuildValues string

const (
	AnimationChartOnlyBuildValuesSeries     AnimationChartOnlyBuildValues = "series"
	AnimationChartOnlyBuildValuesCategory   AnimationChartOnlyBuildValues = "category"
	AnimationChartOnlyBuildValuesSeriesEl   AnimationChartOnlyBuildValues = "seriesEl"
	AnimationChartOnlyBuildValuesCategoryEl AnimationChartOnlyBuildValues = "categoryEl"
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
	CalendarValuesGregorianXlitEnglish CalendarValues = "gregorianXlitEnglish"
	CalendarValuesGregorianXlitFrench  CalendarValues = "gregorianXlitFrench"
	CalendarValuesGregorianUs          CalendarValues = "gregorianUs"
	CalendarValuesGregorianMeFrench    CalendarValues = "gregorianMeFrench"
	CalendarValuesGregorianArabic      CalendarValues = "gregorianArabic"
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

// OleUpdateValues
type OleUpdateValues string

const (
	OleUpdateValuesOLEUPDATEALWAYS OleUpdateValues = "OLEUPDATE_ALWAYS"
	OleUpdateValuesOLEUPDATEONCALL OleUpdateValues = "OLEUPDATE_ONCALL"
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

// IconSetTypeValues
type IconSetTypeValues string

const (
	IconSetTypeValuesIconSetTypeValues3Arrows         IconSetTypeValues = "3Arrows"
	IconSetTypeValuesIconSetTypeValues3ArrowsGray     IconSetTypeValues = "3ArrowsGray"
	IconSetTypeValuesIconSetTypeValues3Flags          IconSetTypeValues = "3Flags"
	IconSetTypeValuesIconSetTypeValues3TrafficLights1 IconSetTypeValues = "3TrafficLights1"
	IconSetTypeValuesIconSetTypeValues3TrafficLights2 IconSetTypeValues = "3TrafficLights2"
	IconSetTypeValuesIconSetTypeValues3Signs          IconSetTypeValues = "3Signs"
	IconSetTypeValuesIconSetTypeValues3Symbols        IconSetTypeValues = "3Symbols"
	IconSetTypeValuesIconSetTypeValues3Symbols2       IconSetTypeValues = "3Symbols2"
	IconSetTypeValuesIconSetTypeValues4Arrows         IconSetTypeValues = "4Arrows"
	IconSetTypeValuesIconSetTypeValues4ArrowsGray     IconSetTypeValues = "4ArrowsGray"
	IconSetTypeValuesIconSetTypeValues4RedToBlack     IconSetTypeValues = "4RedToBlack"
	IconSetTypeValuesIconSetTypeValues4Rating         IconSetTypeValues = "4Rating"
	IconSetTypeValuesIconSetTypeValues4TrafficLights  IconSetTypeValues = "4TrafficLights"
	IconSetTypeValuesIconSetTypeValues5Arrows         IconSetTypeValues = "5Arrows"
	IconSetTypeValuesIconSetTypeValues5ArrowsGray     IconSetTypeValues = "5ArrowsGray"
	IconSetTypeValuesIconSetTypeValues5Rating         IconSetTypeValues = "5Rating"
	IconSetTypeValuesIconSetTypeValues5Quarters       IconSetTypeValues = "5Quarters"
	IconSetTypeValuesIconSetTypeValues3Stars          IconSetTypeValues = "3Stars"
	IconSetTypeValuesIconSetTypeValues3Triangles      IconSetTypeValues = "3Triangles"
	IconSetTypeValuesIconSetTypeValues5Boxes          IconSetTypeValues = "5Boxes"
	IconSetTypeValuesNoIcons                          IconSetTypeValues = "NoIcons"
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

// ChartSubElementValues
type ChartSubElementValues string

const (
	ChartSubElementValuesGridLegend   ChartSubElementValues = "gridLegend"
	ChartSubElementValuesSeries       ChartSubElementValues = "series"
	ChartSubElementValuesCategory     ChartSubElementValues = "category"
	ChartSubElementValuesPtInSeries   ChartSubElementValues = "ptInSeries"
	ChartSubElementValuesPtInCategory ChartSubElementValues = "ptInCategory"
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

// SortMethodValues
type SortMethodValues string

const (
	SortMethodValuesStroke SortMethodValues = "stroke"
	SortMethodValuesPinYin SortMethodValues = "pinYin"
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

// DataConsolidateFunctionValues
type DataConsolidateFunctionValues string

const (
	DataConsolidateFunctionValuesAverage   DataConsolidateFunctionValues = "average"
	DataConsolidateFunctionValuesCount     DataConsolidateFunctionValues = "count"
	DataConsolidateFunctionValuesCountNums DataConsolidateFunctionValues = "countNums"
	DataConsolidateFunctionValuesMax       DataConsolidateFunctionValues = "max"
	DataConsolidateFunctionValuesMin       DataConsolidateFunctionValues = "min"
	DataConsolidateFunctionValuesProduct   DataConsolidateFunctionValues = "product"
	DataConsolidateFunctionValuesStdDev    DataConsolidateFunctionValues = "stdDev"
	DataConsolidateFunctionValuesStdDevp   DataConsolidateFunctionValues = "stdDevp"
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

// DataValidationImeModeValues
type DataValidationImeModeValues string

const (
	DataValidationImeModeValuesNoControl    DataValidationImeModeValues = "noControl"
	DataValidationImeModeValuesOff          DataValidationImeModeValues = "off"
	DataValidationImeModeValuesOn           DataValidationImeModeValues = "on"
	DataValidationImeModeValuesDisabled     DataValidationImeModeValues = "disabled"
	DataValidationImeModeValuesHiragana     DataValidationImeModeValues = "hiragana"
	DataValidationImeModeValuesFullKatakana DataValidationImeModeValues = "fullKatakana"
	DataValidationImeModeValuesHalfKatakana DataValidationImeModeValues = "halfKatakana"
	DataValidationImeModeValuesFullAlpha    DataValidationImeModeValues = "fullAlpha"
	DataValidationImeModeValuesHalfAlpha    DataValidationImeModeValues = "halfAlpha"
	DataValidationImeModeValuesFullHangul   DataValidationImeModeValues = "fullHangul"
	DataValidationImeModeValuesHalfHangul   DataValidationImeModeValues = "halfHangul"
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

// PatternValues
type PatternValues string

const (
	PatternValuesNone            PatternValues = "none"
	PatternValuesSolid           PatternValues = "solid"
	PatternValuesMediumGray      PatternValues = "mediumGray"
	PatternValuesDarkGray        PatternValues = "darkGray"
	PatternValuesLightGray       PatternValues = "lightGray"
	PatternValuesDarkHorizontal  PatternValues = "darkHorizontal"
	PatternValuesDarkVertical    PatternValues = "darkVertical"
	PatternValuesDarkDown        PatternValues = "darkDown"
	PatternValuesDarkUp          PatternValues = "darkUp"
	PatternValuesDarkGrid        PatternValues = "darkGrid"
	PatternValuesDarkTrellis     PatternValues = "darkTrellis"
	PatternValuesLightHorizontal PatternValues = "lightHorizontal"
	PatternValuesLightVertical   PatternValues = "lightVertical"
	PatternValuesLightDown       PatternValues = "lightDown"
	PatternValuesLightUp         PatternValues = "lightUp"
	PatternValuesLightGrid       PatternValues = "lightGrid"
	PatternValuesLightTrellis    PatternValues = "lightTrellis"
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

// CellValues
type CellValues string

const (
	CellValuesB         CellValues = "b"
	CellValuesN         CellValues = "n"
	CellValuesE         CellValues = "e"
	CellValuesS         CellValues = "s"
	CellValuesStr       CellValues = "str"
	CellValuesInlineStr CellValues = "inlineStr"
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

// AnimationBuildValues
type AnimationBuildValues string

const (
	AnimationBuildValuesAllAtOnce AnimationBuildValues = "allAtOnce"
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

// DataValidationOperatorValues
type DataValidationOperatorValues string

const (
	DataValidationOperatorValuesBetween            DataValidationOperatorValues = "between"
	DataValidationOperatorValuesNotBetween         DataValidationOperatorValues = "notBetween"
	DataValidationOperatorValuesEqual              DataValidationOperatorValues = "equal"
	DataValidationOperatorValuesNotEqual           DataValidationOperatorValues = "notEqual"
	DataValidationOperatorValuesLessThan           DataValidationOperatorValues = "lessThan"
	DataValidationOperatorValuesLessThanOrEqual    DataValidationOperatorValues = "lessThanOrEqual"
	DataValidationOperatorValuesGreaterThan        DataValidationOperatorValues = "greaterThan"
	DataValidationOperatorValuesGreaterThanOrEqual DataValidationOperatorValues = "greaterThanOrEqual"
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

// TextBoxFormFieldValues
type TextBoxFormFieldValues string

const (
	TextBoxFormFieldValuesRegular     TextBoxFormFieldValues = "regular"
	TextBoxFormFieldValuesNumber      TextBoxFormFieldValues = "number"
	TextBoxFormFieldValuesDate        TextBoxFormFieldValues = "date"
	TextBoxFormFieldValuesCurrentTime TextBoxFormFieldValues = "currentTime"
	TextBoxFormFieldValuesCurrentDate TextBoxFormFieldValues = "currentDate"
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

// TextUnderlineValues
type TextUnderlineValues string

const (
	TextUnderlineValuesNone            TextUnderlineValues = "none"
	TextUnderlineValuesWords           TextUnderlineValues = "words"
	TextUnderlineValuesSng             TextUnderlineValues = "sng"
	TextUnderlineValuesDbl             TextUnderlineValues = "dbl"
	TextUnderlineValuesHeavy           TextUnderlineValues = "heavy"
	TextUnderlineValuesDotted          TextUnderlineValues = "dotted"
	TextUnderlineValuesDottedHeavy     TextUnderlineValues = "dottedHeavy"
	TextUnderlineValuesDash            TextUnderlineValues = "dash"
	TextUnderlineValuesDashHeavy       TextUnderlineValues = "dashHeavy"
	TextUnderlineValuesDashLong        TextUnderlineValues = "dashLong"
	TextUnderlineValuesDashLongHeavy   TextUnderlineValues = "dashLongHeavy"
	TextUnderlineValuesDotDash         TextUnderlineValues = "dotDash"
	TextUnderlineValuesDotDashHeavy    TextUnderlineValues = "dotDashHeavy"
	TextUnderlineValuesDotDotDash      TextUnderlineValues = "dotDotDash"
	TextUnderlineValuesDotDotDashHeavy TextUnderlineValues = "dotDotDashHeavy"
	TextUnderlineValuesWavy            TextUnderlineValues = "wavy"
	TextUnderlineValuesWavyHeavy       TextUnderlineValues = "wavyHeavy"
	TextUnderlineValuesWavyDbl         TextUnderlineValues = "wavyDbl"
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

// TimePeriodValues
type TimePeriodValues string

const (
	TimePeriodValuesToday     TimePeriodValues = "today"
	TimePeriodValuesYesterday TimePeriodValues = "yesterday"
	TimePeriodValuesTomorrow  TimePeriodValues = "tomorrow"
	TimePeriodValuesLast7Days TimePeriodValues = "last7Days"
	TimePeriodValuesThisMonth TimePeriodValues = "thisMonth"
	TimePeriodValuesLastMonth TimePeriodValues = "lastMonth"
	TimePeriodValuesNextMonth TimePeriodValues = "nextMonth"
	TimePeriodValuesThisWeek  TimePeriodValues = "thisWeek"
	TimePeriodValuesLastWeek  TimePeriodValues = "lastWeek"
	TimePeriodValuesNextWeek  TimePeriodValues = "nextWeek"
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

// AnimationDiagramOnlyBuildValues
type AnimationDiagramOnlyBuildValues string

const (
	AnimationDiagramOnlyBuildValuesOne       AnimationDiagramOnlyBuildValues = "one"
	AnimationDiagramOnlyBuildValuesLvlOne    AnimationDiagramOnlyBuildValues = "lvlOne"
	AnimationDiagramOnlyBuildValuesLvlAtOnce AnimationDiagramOnlyBuildValues = "lvlAtOnce"
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

// ObjectUpdateMode
type ObjectUpdateMode string

const (
	ObjectUpdateModeAlways ObjectUpdateMode = "always"
	ObjectUpdateModeOnCall ObjectUpdateMode = "onCall"
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

// SortByValues
type SortByValues string

const (
	SortByValuesValue     SortByValues = "value"
	SortByValuesCellColor SortByValues = "cellColor"
	SortByValuesFontColor SortByValues = "fontColor"
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

// PivotTableAxisValues
type PivotTableAxisValues string

const (
	PivotTableAxisValuesAxisRow    PivotTableAxisValues = "axisRow"
	PivotTableAxisValuesAxisCol    PivotTableAxisValues = "axisCol"
	PivotTableAxisValuesAxisPage   PivotTableAxisValues = "axisPage"
	PivotTableAxisValuesAxisValues PivotTableAxisValues = "axisValues"
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

// HeightRuleValues
type HeightRuleValues string

const (
	HeightRuleValuesAuto    HeightRuleValues = "auto"
	HeightRuleValuesExact   HeightRuleValues = "exact"
	HeightRuleValuesAtLeast HeightRuleValues = "atLeast"
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

// MailMergeDocumentValues
type MailMergeDocumentValues string

const (
	MailMergeDocumentValuesCatalog       MailMergeDocumentValues = "catalog"
	MailMergeDocumentValuesEnvelopes     MailMergeDocumentValues = "envelopes"
	MailMergeDocumentValuesMailingLabels MailMergeDocumentValues = "mailingLabels"
	MailMergeDocumentValuesFormLetters   MailMergeDocumentValues = "formLetters"
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

// ThemeFontValues
type ThemeFontValues string

const (
	ThemeFontValuesMajorEastAsia ThemeFontValues = "majorEastAsia"
	ThemeFontValuesMajorBidi     ThemeFontValues = "majorBidi"
	ThemeFontValuesMajorAscii    ThemeFontValues = "majorAscii"
	ThemeFontValuesMajorHAnsi    ThemeFontValues = "majorHAnsi"
	ThemeFontValuesMinorEastAsia ThemeFontValues = "minorEastAsia"
	ThemeFontValuesMinorBidi     ThemeFontValues = "minorBidi"
	ThemeFontValuesMinorAscii    ThemeFontValues = "minorAscii"
	ThemeFontValuesMinorHAnsi    ThemeFontValues = "minorHAnsi"
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

// SystemColorValues
type SystemColorValues string

const (
	SystemColorValuesScrollBar                   SystemColorValues = "scrollBar"
	SystemColorValuesBackground                  SystemColorValues = "background"
	SystemColorValuesActiveCaption               SystemColorValues = "activeCaption"
	SystemColorValuesInactiveCaption             SystemColorValues = "inactiveCaption"
	SystemColorValuesMenu                        SystemColorValues = "menu"
	SystemColorValuesWindow                      SystemColorValues = "window"
	SystemColorValuesWindowFrame                 SystemColorValues = "windowFrame"
	SystemColorValuesMenuText                    SystemColorValues = "menuText"
	SystemColorValuesWindowText                  SystemColorValues = "windowText"
	SystemColorValuesCaptionText                 SystemColorValues = "captionText"
	SystemColorValuesActiveBorder                SystemColorValues = "activeBorder"
	SystemColorValuesInactiveBorder              SystemColorValues = "inactiveBorder"
	SystemColorValuesAppWorkspace                SystemColorValues = "appWorkspace"
	SystemColorValuesHighlight                   SystemColorValues = "highlight"
	SystemColorValuesHighlightText               SystemColorValues = "highlightText"
	SystemColorValuesBtnFace                     SystemColorValues = "btnFace"
	SystemColorValuesBtnShadow                   SystemColorValues = "btnShadow"
	SystemColorValuesGrayText                    SystemColorValues = "grayText"
	SystemColorValuesBtnText                     SystemColorValues = "btnText"
	SystemColorValuesInactiveCaptionText         SystemColorValues = "inactiveCaptionText"
	SystemColorValuesBtnHighlight                SystemColorValues = "btnHighlight"
	SystemColorValuesSystemColorValues3dDkShadow SystemColorValues = "3dDkShadow"
	SystemColorValuesSystemColorValues3dLight    SystemColorValues = "3dLight"
	SystemColorValuesInfoText                    SystemColorValues = "infoText"
	SystemColorValuesInfoBk                      SystemColorValues = "infoBk"
	SystemColorValuesHotLight                    SystemColorValues = "hotLight"
	SystemColorValuesGradientActiveCaption       SystemColorValues = "gradientActiveCaption"
	SystemColorValuesGradientInactiveCaption     SystemColorValues = "gradientInactiveCaption"
	SystemColorValuesMenuHighlight               SystemColorValues = "menuHighlight"
	SystemColorValuesMenuBar                     SystemColorValues = "menuBar"
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

// PresetCameraValues
type PresetCameraValues string

const (
	PresetCameraValuesLegacyObliqueTopLeft                PresetCameraValues = "legacyObliqueTopLeft"
	PresetCameraValuesLegacyObliqueTop                    PresetCameraValues = "legacyObliqueTop"
	PresetCameraValuesLegacyObliqueTopRight               PresetCameraValues = "legacyObliqueTopRight"
	PresetCameraValuesLegacyObliqueLeft                   PresetCameraValues = "legacyObliqueLeft"
	PresetCameraValuesLegacyObliqueFront                  PresetCameraValues = "legacyObliqueFront"
	PresetCameraValuesLegacyObliqueRight                  PresetCameraValues = "legacyObliqueRight"
	PresetCameraValuesLegacyObliqueBottomLeft             PresetCameraValues = "legacyObliqueBottomLeft"
	PresetCameraValuesLegacyObliqueBottom                 PresetCameraValues = "legacyObliqueBottom"
	PresetCameraValuesLegacyObliqueBottomRight            PresetCameraValues = "legacyObliqueBottomRight"
	PresetCameraValuesLegacyPerspectiveTopLeft            PresetCameraValues = "legacyPerspectiveTopLeft"
	PresetCameraValuesLegacyPerspectiveTop                PresetCameraValues = "legacyPerspectiveTop"
	PresetCameraValuesLegacyPerspectiveTopRight           PresetCameraValues = "legacyPerspectiveTopRight"
	PresetCameraValuesLegacyPerspectiveLeft               PresetCameraValues = "legacyPerspectiveLeft"
	PresetCameraValuesLegacyPerspectiveFront              PresetCameraValues = "legacyPerspectiveFront"
	PresetCameraValuesLegacyPerspectiveRight              PresetCameraValues = "legacyPerspectiveRight"
	PresetCameraValuesLegacyPerspectiveBottomLeft         PresetCameraValues = "legacyPerspectiveBottomLeft"
	PresetCameraValuesLegacyPerspectiveBottom             PresetCameraValues = "legacyPerspectiveBottom"
	PresetCameraValuesLegacyPerspectiveBottomRight        PresetCameraValues = "legacyPerspectiveBottomRight"
	PresetCameraValuesOrthographicFront                   PresetCameraValues = "orthographicFront"
	PresetCameraValuesIsometricTopUp                      PresetCameraValues = "isometricTopUp"
	PresetCameraValuesIsometricTopDown                    PresetCameraValues = "isometricTopDown"
	PresetCameraValuesIsometricBottomUp                   PresetCameraValues = "isometricBottomUp"
	PresetCameraValuesIsometricBottomDown                 PresetCameraValues = "isometricBottomDown"
	PresetCameraValuesIsometricLeftUp                     PresetCameraValues = "isometricLeftUp"
	PresetCameraValuesIsometricLeftDown                   PresetCameraValues = "isometricLeftDown"
	PresetCameraValuesIsometricRightUp                    PresetCameraValues = "isometricRightUp"
	PresetCameraValuesIsometricRightDown                  PresetCameraValues = "isometricRightDown"
	PresetCameraValuesIsometricOffAxis1Left               PresetCameraValues = "isometricOffAxis1Left"
	PresetCameraValuesIsometricOffAxis1Right              PresetCameraValues = "isometricOffAxis1Right"
	PresetCameraValuesIsometricOffAxis1Top                PresetCameraValues = "isometricOffAxis1Top"
	PresetCameraValuesIsometricOffAxis2Left               PresetCameraValues = "isometricOffAxis2Left"
	PresetCameraValuesIsometricOffAxis2Right              PresetCameraValues = "isometricOffAxis2Right"
	PresetCameraValuesIsometricOffAxis2Top                PresetCameraValues = "isometricOffAxis2Top"
	PresetCameraValuesIsometricOffAxis3Left               PresetCameraValues = "isometricOffAxis3Left"
	PresetCameraValuesIsometricOffAxis3Right              PresetCameraValues = "isometricOffAxis3Right"
	PresetCameraValuesIsometricOffAxis3Bottom             PresetCameraValues = "isometricOffAxis3Bottom"
	PresetCameraValuesIsometricOffAxis4Left               PresetCameraValues = "isometricOffAxis4Left"
	PresetCameraValuesIsometricOffAxis4Right              PresetCameraValues = "isometricOffAxis4Right"
	PresetCameraValuesIsometricOffAxis4Bottom             PresetCameraValues = "isometricOffAxis4Bottom"
	PresetCameraValuesObliqueTopLeft                      PresetCameraValues = "obliqueTopLeft"
	PresetCameraValuesObliqueTop                          PresetCameraValues = "obliqueTop"
	PresetCameraValuesObliqueTopRight                     PresetCameraValues = "obliqueTopRight"
	PresetCameraValuesObliqueLeft                         PresetCameraValues = "obliqueLeft"
	PresetCameraValuesObliqueRight                        PresetCameraValues = "obliqueRight"
	PresetCameraValuesObliqueBottomLeft                   PresetCameraValues = "obliqueBottomLeft"
	PresetCameraValuesObliqueBottom                       PresetCameraValues = "obliqueBottom"
	PresetCameraValuesObliqueBottomRight                  PresetCameraValues = "obliqueBottomRight"
	PresetCameraValuesPerspectiveFront                    PresetCameraValues = "perspectiveFront"
	PresetCameraValuesPerspectiveLeft                     PresetCameraValues = "perspectiveLeft"
	PresetCameraValuesPerspectiveRight                    PresetCameraValues = "perspectiveRight"
	PresetCameraValuesPerspectiveAbove                    PresetCameraValues = "perspectiveAbove"
	PresetCameraValuesPerspectiveBelow                    PresetCameraValues = "perspectiveBelow"
	PresetCameraValuesPerspectiveAboveLeftFacing          PresetCameraValues = "perspectiveAboveLeftFacing"
	PresetCameraValuesPerspectiveAboveRightFacing         PresetCameraValues = "perspectiveAboveRightFacing"
	PresetCameraValuesPerspectiveContrastingLeftFacing    PresetCameraValues = "perspectiveContrastingLeftFacing"
	PresetCameraValuesPerspectiveContrastingRightFacing   PresetCameraValues = "perspectiveContrastingRightFacing"
	PresetCameraValuesPerspectiveHeroicLeftFacing         PresetCameraValues = "perspectiveHeroicLeftFacing"
	PresetCameraValuesPerspectiveHeroicRightFacing        PresetCameraValues = "perspectiveHeroicRightFacing"
	PresetCameraValuesPerspectiveHeroicExtremeLeftFacing  PresetCameraValues = "perspectiveHeroicExtremeLeftFacing"
	PresetCameraValuesPerspectiveHeroicExtremeRightFacing PresetCameraValues = "perspectiveHeroicExtremeRightFacing"
	PresetCameraValuesPerspectiveRelaxed                  PresetCameraValues = "perspectiveRelaxed"
	PresetCameraValuesPerspectiveRelaxedModerately        PresetCameraValues = "perspectiveRelaxedModerately"
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

// BehaviorOverrideValues
type BehaviorOverrideValues string

const (
	BehaviorOverrideValuesNormal     BehaviorOverrideValues = "normal"
	BehaviorOverrideValuesChildStyle BehaviorOverrideValues = "childStyle"
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

// ExternalConnectionValues
type ExternalConnectionValues string

const (
	ExternalConnectionValuesGeneral ExternalConnectionValues = "general"
	ExternalConnectionValuesText    ExternalConnectionValues = "text"
	ExternalConnectionValuesMDY     ExternalConnectionValues = "MDY"
	ExternalConnectionValuesDMY     ExternalConnectionValues = "DMY"
	ExternalConnectionValuesYMD     ExternalConnectionValues = "YMD"
	ExternalConnectionValuesMYD     ExternalConnectionValues = "MYD"
	ExternalConnectionValuesDYM     ExternalConnectionValues = "DYM"
	ExternalConnectionValuesYDM     ExternalConnectionValues = "YDM"
	ExternalConnectionValuesSkip    ExternalConnectionValues = "skip"
	ExternalConnectionValuesEMD     ExternalConnectionValues = "EMD"
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

// FontTypeHintValues
type FontTypeHintValues string

const (
	FontTypeHintValuesDefault  FontTypeHintValues = "default"
	FontTypeHintValuesEastAsia FontTypeHintValues = "eastAsia"
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

// BevelPresetValues
type BevelPresetValues string

const (
	BevelPresetValuesRelaxedInset BevelPresetValues = "relaxedInset"
	BevelPresetValuesCircle       BevelPresetValues = "circle"
	BevelPresetValuesSlope        BevelPresetValues = "slope"
	BevelPresetValuesCross        BevelPresetValues = "cross"
	BevelPresetValuesAngle        BevelPresetValues = "angle"
	BevelPresetValuesSoftRound    BevelPresetValues = "softRound"
	BevelPresetValuesConvex       BevelPresetValues = "convex"
	BevelPresetValuesCoolSlant    BevelPresetValues = "coolSlant"
	BevelPresetValuesDivot        BevelPresetValues = "divot"
	BevelPresetValuesRiblet       BevelPresetValues = "riblet"
	BevelPresetValuesHardEdge     BevelPresetValues = "hardEdge"
	BevelPresetValuesArtDeco      BevelPresetValues = "artDeco"
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

// IconSetValues
type IconSetValues string

const (
	IconSetValuesIconSetValues3Arrows         IconSetValues = "3Arrows"
	IconSetValuesIconSetValues3ArrowsGray     IconSetValues = "3ArrowsGray"
	IconSetValuesIconSetValues3Flags          IconSetValues = "3Flags"
	IconSetValuesIconSetValues3TrafficLights1 IconSetValues = "3TrafficLights1"
	IconSetValuesIconSetValues3TrafficLights2 IconSetValues = "3TrafficLights2"
	IconSetValuesIconSetValues3Signs          IconSetValues = "3Signs"
	IconSetValuesIconSetValues3Symbols        IconSetValues = "3Symbols"
	IconSetValuesIconSetValues3Symbols2       IconSetValues = "3Symbols2"
	IconSetValuesIconSetValues4Arrows         IconSetValues = "4Arrows"
	IconSetValuesIconSetValues4ArrowsGray     IconSetValues = "4ArrowsGray"
	IconSetValuesIconSetValues4RedToBlack     IconSetValues = "4RedToBlack"
	IconSetValuesIconSetValues4Rating         IconSetValues = "4Rating"
	IconSetValuesIconSetValues4TrafficLights  IconSetValues = "4TrafficLights"
	IconSetValuesIconSetValues5Arrows         IconSetValues = "5Arrows"
	IconSetValuesIconSetValues5ArrowsGray     IconSetValues = "5ArrowsGray"
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

// XmlDataValues
type XmlDataValues string

const (
	XmlDataValuesString             XmlDataValues = "string"
	XmlDataValuesNormalizedString   XmlDataValues = "normalizedString"
	XmlDataValuesToken              XmlDataValues = "token"
	XmlDataValuesByte               XmlDataValues = "byte"
	XmlDataValuesUnsignedByte       XmlDataValues = "unsignedByte"
	XmlDataValuesBase64Binary       XmlDataValues = "base64Binary"
	XmlDataValuesHexBinary          XmlDataValues = "hexBinary"
	XmlDataValuesInteger            XmlDataValues = "integer"
	XmlDataValuesPositiveInteger    XmlDataValues = "positiveInteger"
	XmlDataValuesNegativeInteger    XmlDataValues = "negativeInteger"
	XmlDataValuesNonPositiveInteger XmlDataValues = "nonPositiveInteger"
	XmlDataValuesNonNegativeInteger XmlDataValues = "nonNegativeInteger"
	XmlDataValuesInt                XmlDataValues = "int"
	XmlDataValuesUnsignedInt        XmlDataValues = "unsignedInt"
	XmlDataValuesLong               XmlDataValues = "long"
	XmlDataValuesUnsignedLong       XmlDataValues = "unsignedLong"
	XmlDataValuesShort              XmlDataValues = "short"
	XmlDataValuesUnsignedShort      XmlDataValues = "unsignedShort"
	XmlDataValuesDecimal            XmlDataValues = "decimal"
	XmlDataValuesFloat              XmlDataValues = "float"
	XmlDataValuesDouble             XmlDataValues = "double"
	XmlDataValuesBoolean            XmlDataValues = "boolean"
	XmlDataValuesTime               XmlDataValues = "time"
	XmlDataValuesDateTime           XmlDataValues = "dateTime"
	XmlDataValuesDuration           XmlDataValues = "duration"
	XmlDataValuesDate               XmlDataValues = "date"
	XmlDataValuesGMonth             XmlDataValues = "gMonth"
	XmlDataValuesGYear              XmlDataValues = "gYear"
	XmlDataValuesGYearMonth         XmlDataValues = "gYearMonth"
	XmlDataValuesGDay               XmlDataValues = "gDay"
	XmlDataValuesGMonthDay          XmlDataValues = "gMonthDay"
	XmlDataValuesName               XmlDataValues = "Name"
	XmlDataValuesQName              XmlDataValues = "QName"
	XmlDataValuesNCName             XmlDataValues = "NCName"
	XmlDataValuesAnyURI             XmlDataValues = "anyURI"
	XmlDataValuesLanguage           XmlDataValues = "language"
	XmlDataValuesID                 XmlDataValues = "ID"
	XmlDataValuesIDREF              XmlDataValues = "IDREF"
	XmlDataValuesIDREFS             XmlDataValues = "IDREFS"
	XmlDataValuesENTITY             XmlDataValues = "ENTITY"
	XmlDataValuesENTITIES           XmlDataValues = "ENTITIES"
	XmlDataValuesNOTATION           XmlDataValues = "NOTATION"
	XmlDataValuesNMTOKEN            XmlDataValues = "NMTOKEN"
	XmlDataValuesNMTOKENS           XmlDataValues = "NMTOKENS"
	XmlDataValuesAnyType            XmlDataValues = "anyType"
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

// DiagramBuildValues
type DiagramBuildValues string

const (
	DiagramBuildValuesWhole         DiagramBuildValues = "whole"
	DiagramBuildValuesDepthByNode   DiagramBuildValues = "depthByNode"
	DiagramBuildValuesDepthByBranch DiagramBuildValues = "depthByBranch"
	DiagramBuildValuesBreadthByNode DiagramBuildValues = "breadthByNode"
	DiagramBuildValuesBreadthByLvl  DiagramBuildValues = "breadthByLvl"
	DiagramBuildValuesCw            DiagramBuildValues = "cw"
	DiagramBuildValuesCwIn          DiagramBuildValues = "cwIn"
	DiagramBuildValuesCwOut         DiagramBuildValues = "cwOut"
	DiagramBuildValuesCcw           DiagramBuildValues = "ccw"
	DiagramBuildValuesCcwIn         DiagramBuildValues = "ccwIn"
	DiagramBuildValuesCcwOut        DiagramBuildValues = "ccwOut"
	DiagramBuildValuesInByRing      DiagramBuildValues = "inByRing"
	DiagramBuildValuesOutByRing     DiagramBuildValues = "outByRing"
	DiagramBuildValuesUp            DiagramBuildValues = "up"
	DiagramBuildValuesDown          DiagramBuildValues = "down"
	DiagramBuildValuesAllAtOnce     DiagramBuildValues = "allAtOnce"
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

// MailMergeDestinationValues
type MailMergeDestinationValues string

const (
	MailMergeDestinationValuesNewDocument MailMergeDestinationValues = "newDocument"
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

// GrowShrinkValues
type GrowShrinkValues string

const (
	GrowShrinkValuesInsertDelete   GrowShrinkValues = "insertDelete"
	GrowShrinkValuesInsertClear    GrowShrinkValues = "insertClear"
	GrowShrinkValuesOverwriteClear GrowShrinkValues = "overwriteClear"
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

// DocGridValues
type DocGridValues string

const (
	DocGridValuesDefault       DocGridValues = "default"
	DocGridValuesLines         DocGridValues = "lines"
	DocGridValuesLinesAndChars DocGridValues = "linesAndChars"
	DocGridValuesSnapToChars   DocGridValues = "snapToChars"
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

// PaneValues
type PaneValues string

const (
	PaneValuesBottomRight PaneValues = "bottomRight"
	PaneValuesTopRight    PaneValues = "topRight"
	PaneValuesBottomLeft  PaneValues = "bottomLeft"
	PaneValuesTopLeft     PaneValues = "topLeft"
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

// ShadingPatternValues
type ShadingPatternValues string

const (
	ShadingPatternValuesNil                   ShadingPatternValues = "nil"
	ShadingPatternValuesClear                 ShadingPatternValues = "clear"
	ShadingPatternValuesSolid                 ShadingPatternValues = "solid"
	ShadingPatternValuesHorzStripe            ShadingPatternValues = "horzStripe"
	ShadingPatternValuesVertStripe            ShadingPatternValues = "vertStripe"
	ShadingPatternValuesReverseDiagStripe     ShadingPatternValues = "reverseDiagStripe"
	ShadingPatternValuesDiagStripe            ShadingPatternValues = "diagStripe"
	ShadingPatternValuesHorzCross             ShadingPatternValues = "horzCross"
	ShadingPatternValuesDiagCross             ShadingPatternValues = "diagCross"
	ShadingPatternValuesThinHorzStripe        ShadingPatternValues = "thinHorzStripe"
	ShadingPatternValuesThinVertStripe        ShadingPatternValues = "thinVertStripe"
	ShadingPatternValuesThinReverseDiagStripe ShadingPatternValues = "thinReverseDiagStripe"
	ShadingPatternValuesThinDiagStripe        ShadingPatternValues = "thinDiagStripe"
	ShadingPatternValuesThinHorzCross         ShadingPatternValues = "thinHorzCross"
	ShadingPatternValuesThinDiagCross         ShadingPatternValues = "thinDiagCross"
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

// DataViewAspectValues
type DataViewAspectValues string

const (
	DataViewAspectValuesDVASPECTCONTENT DataViewAspectValues = "DVASPECT_CONTENT"
	DataViewAspectValuesDVASPECTICON    DataViewAspectValues = "DVASPECT_ICON"
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

// DocumentProtectionValues
type DocumentProtectionValues string

const (
	DocumentProtectionValuesNone           DocumentProtectionValues = "none"
	DocumentProtectionValuesReadOnly       DocumentProtectionValues = "readOnly"
	DocumentProtectionValuesComments       DocumentProtectionValues = "comments"
	DocumentProtectionValuesTrackedChanges DocumentProtectionValues = "trackedChanges"
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

// ShapeTypeValues
type ShapeTypeValues string

const (
	ShapeTypeValuesLine                       ShapeTypeValues = "line"
	ShapeTypeValuesLineInv                    ShapeTypeValues = "lineInv"
	ShapeTypeValuesTriangle                   ShapeTypeValues = "triangle"
	ShapeTypeValuesRtTriangle                 ShapeTypeValues = "rtTriangle"
	ShapeTypeValuesRect                       ShapeTypeValues = "rect"
	ShapeTypeValuesDiamond                    ShapeTypeValues = "diamond"
	ShapeTypeValuesParallelogram              ShapeTypeValues = "parallelogram"
	ShapeTypeValuesTrapezoid                  ShapeTypeValues = "trapezoid"
	ShapeTypeValuesNonIsoscelesTrapezoid      ShapeTypeValues = "nonIsoscelesTrapezoid"
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
	ShapeTypeValuesRoundRect                  ShapeTypeValues = "roundRect"
	ShapeTypeValuesRound1Rect                 ShapeTypeValues = "round1Rect"
	ShapeTypeValuesRound2SameRect             ShapeTypeValues = "round2SameRect"
	ShapeTypeValuesRound2DiagRect             ShapeTypeValues = "round2DiagRect"
	ShapeTypeValuesSnipRoundRect              ShapeTypeValues = "snipRoundRect"
	ShapeTypeValuesSnip1Rect                  ShapeTypeValues = "snip1Rect"
	ShapeTypeValuesSnip2SameRect              ShapeTypeValues = "snip2SameRect"
	ShapeTypeValuesSnip2DiagRect              ShapeTypeValues = "snip2DiagRect"
	ShapeTypeValuesPlaque                     ShapeTypeValues = "plaque"
	ShapeTypeValuesEllipse                    ShapeTypeValues = "ellipse"
	ShapeTypeValuesTeardrop                   ShapeTypeValues = "teardrop"
	ShapeTypeValuesHomePlate                  ShapeTypeValues = "homePlate"
	ShapeTypeValuesChevron                    ShapeTypeValues = "chevron"
	ShapeTypeValuesPieWedge                   ShapeTypeValues = "pieWedge"
	ShapeTypeValuesPie                        ShapeTypeValues = "pie"
	ShapeTypeValuesBlockArc                   ShapeTypeValues = "blockArc"
	ShapeTypeValuesDonut                      ShapeTypeValues = "donut"
	ShapeTypeValuesNoSmoking                  ShapeTypeValues = "noSmoking"
	ShapeTypeValuesRightArrow                 ShapeTypeValues = "rightArrow"
	ShapeTypeValuesLeftArrow                  ShapeTypeValues = "leftArrow"
	ShapeTypeValuesUpArrow                    ShapeTypeValues = "upArrow"
	ShapeTypeValuesDownArrow                  ShapeTypeValues = "downArrow"
	ShapeTypeValuesStripedRightArrow          ShapeTypeValues = "stripedRightArrow"
	ShapeTypeValuesNotchedRightArrow          ShapeTypeValues = "notchedRightArrow"
	ShapeTypeValuesBentUpArrow                ShapeTypeValues = "bentUpArrow"
	ShapeTypeValuesLeftRightArrow             ShapeTypeValues = "leftRightArrow"
	ShapeTypeValuesUpDownArrow                ShapeTypeValues = "upDownArrow"
	ShapeTypeValuesLeftUpArrow                ShapeTypeValues = "leftUpArrow"
	ShapeTypeValuesLeftRightUpArrow           ShapeTypeValues = "leftRightUpArrow"
	ShapeTypeValuesQuadArrow                  ShapeTypeValues = "quadArrow"
	ShapeTypeValuesLeftArrowCallout           ShapeTypeValues = "leftArrowCallout"
	ShapeTypeValuesRightArrowCallout          ShapeTypeValues = "rightArrowCallout"
	ShapeTypeValuesUpArrowCallout             ShapeTypeValues = "upArrowCallout"
	ShapeTypeValuesDownArrowCallout           ShapeTypeValues = "downArrowCallout"
	ShapeTypeValuesLeftRightArrowCallout      ShapeTypeValues = "leftRightArrowCallout"
	ShapeTypeValuesUpDownArrowCallout         ShapeTypeValues = "upDownArrowCallout"
	ShapeTypeValuesQuadArrowCallout           ShapeTypeValues = "quadArrowCallout"
	ShapeTypeValuesBentArrow                  ShapeTypeValues = "bentArrow"
	ShapeTypeValuesUturnArrow                 ShapeTypeValues = "uturnArrow"
	ShapeTypeValuesCircularArrow              ShapeTypeValues = "circularArrow"
	ShapeTypeValuesLeftCircularArrow          ShapeTypeValues = "leftCircularArrow"
	ShapeTypeValuesLeftRightCircularArrow     ShapeTypeValues = "leftRightCircularArrow"
	ShapeTypeValuesCurvedRightArrow           ShapeTypeValues = "curvedRightArrow"
	ShapeTypeValuesCurvedLeftArrow            ShapeTypeValues = "curvedLeftArrow"
	ShapeTypeValuesCurvedUpArrow              ShapeTypeValues = "curvedUpArrow"
	ShapeTypeValuesCurvedDownArrow            ShapeTypeValues = "curvedDownArrow"
	ShapeTypeValuesSwooshArrow                ShapeTypeValues = "swooshArrow"
	ShapeTypeValuesCube                       ShapeTypeValues = "cube"
	ShapeTypeValuesCan                        ShapeTypeValues = "can"
	ShapeTypeValuesLightningBolt              ShapeTypeValues = "lightningBolt"
	ShapeTypeValuesHeart                      ShapeTypeValues = "heart"
	ShapeTypeValuesSun                        ShapeTypeValues = "sun"
	ShapeTypeValuesMoon                       ShapeTypeValues = "moon"
	ShapeTypeValuesSmileyFace                 ShapeTypeValues = "smileyFace"
	ShapeTypeValuesIrregularSeal1             ShapeTypeValues = "irregularSeal1"
	ShapeTypeValuesIrregularSeal2             ShapeTypeValues = "irregularSeal2"
	ShapeTypeValuesFoldedCorner               ShapeTypeValues = "foldedCorner"
	ShapeTypeValuesBevel                      ShapeTypeValues = "bevel"
	ShapeTypeValuesFrame                      ShapeTypeValues = "frame"
	ShapeTypeValuesHalfFrame                  ShapeTypeValues = "halfFrame"
	ShapeTypeValuesCorner                     ShapeTypeValues = "corner"
	ShapeTypeValuesDiagStripe                 ShapeTypeValues = "diagStripe"
	ShapeTypeValuesChord                      ShapeTypeValues = "chord"
	ShapeTypeValuesArc                        ShapeTypeValues = "arc"
	ShapeTypeValuesLeftBracket                ShapeTypeValues = "leftBracket"
	ShapeTypeValuesRightBracket               ShapeTypeValues = "rightBracket"
	ShapeTypeValuesLeftBrace                  ShapeTypeValues = "leftBrace"
	ShapeTypeValuesRightBrace                 ShapeTypeValues = "rightBrace"
	ShapeTypeValuesBracketPair                ShapeTypeValues = "bracketPair"
	ShapeTypeValuesBracePair                  ShapeTypeValues = "bracePair"
	ShapeTypeValuesStraightConnector1         ShapeTypeValues = "straightConnector1"
	ShapeTypeValuesBentConnector2             ShapeTypeValues = "bentConnector2"
	ShapeTypeValuesBentConnector3             ShapeTypeValues = "bentConnector3"
	ShapeTypeValuesBentConnector4             ShapeTypeValues = "bentConnector4"
	ShapeTypeValuesBentConnector5             ShapeTypeValues = "bentConnector5"
	ShapeTypeValuesCurvedConnector2           ShapeTypeValues = "curvedConnector2"
	ShapeTypeValuesCurvedConnector3           ShapeTypeValues = "curvedConnector3"
	ShapeTypeValuesCurvedConnector4           ShapeTypeValues = "curvedConnector4"
	ShapeTypeValuesCurvedConnector5           ShapeTypeValues = "curvedConnector5"
	ShapeTypeValuesCallout1                   ShapeTypeValues = "callout1"
	ShapeTypeValuesCallout2                   ShapeTypeValues = "callout2"
	ShapeTypeValuesCallout3                   ShapeTypeValues = "callout3"
	ShapeTypeValuesAccentCallout1             ShapeTypeValues = "accentCallout1"
	ShapeTypeValuesAccentCallout2             ShapeTypeValues = "accentCallout2"
	ShapeTypeValuesAccentCallout3             ShapeTypeValues = "accentCallout3"
	ShapeTypeValuesBorderCallout1             ShapeTypeValues = "borderCallout1"
	ShapeTypeValuesBorderCallout2             ShapeTypeValues = "borderCallout2"
	ShapeTypeValuesBorderCallout3             ShapeTypeValues = "borderCallout3"
	ShapeTypeValuesAccentBorderCallout1       ShapeTypeValues = "accentBorderCallout1"
	ShapeTypeValuesAccentBorderCallout2       ShapeTypeValues = "accentBorderCallout2"
	ShapeTypeValuesAccentBorderCallout3       ShapeTypeValues = "accentBorderCallout3"
	ShapeTypeValuesWedgeRectCallout           ShapeTypeValues = "wedgeRectCallout"
	ShapeTypeValuesWedgeRoundRectCallout      ShapeTypeValues = "wedgeRoundRectCallout"
	ShapeTypeValuesWedgeEllipseCallout        ShapeTypeValues = "wedgeEllipseCallout"
	ShapeTypeValuesCloudCallout               ShapeTypeValues = "cloudCallout"
	ShapeTypeValuesCloud                      ShapeTypeValues = "cloud"
	ShapeTypeValuesRibbon                     ShapeTypeValues = "ribbon"
	ShapeTypeValuesRibbon2                    ShapeTypeValues = "ribbon2"
	ShapeTypeValuesEllipseRibbon              ShapeTypeValues = "ellipseRibbon"
	ShapeTypeValuesEllipseRibbon2             ShapeTypeValues = "ellipseRibbon2"
	ShapeTypeValuesLeftRightRibbon            ShapeTypeValues = "leftRightRibbon"
	ShapeTypeValuesVerticalScroll             ShapeTypeValues = "verticalScroll"
	ShapeTypeValuesHorizontalScroll           ShapeTypeValues = "horizontalScroll"
	ShapeTypeValuesWave                       ShapeTypeValues = "wave"
	ShapeTypeValuesDoubleWave                 ShapeTypeValues = "doubleWave"
	ShapeTypeValuesPlus                       ShapeTypeValues = "plus"
	ShapeTypeValuesFlowChartProcess           ShapeTypeValues = "flowChartProcess"
	ShapeTypeValuesFlowChartDecision          ShapeTypeValues = "flowChartDecision"
	ShapeTypeValuesFlowChartInputOutput       ShapeTypeValues = "flowChartInputOutput"
	ShapeTypeValuesFlowChartPredefinedProcess ShapeTypeValues = "flowChartPredefinedProcess"
	ShapeTypeValuesFlowChartInternalStorage   ShapeTypeValues = "flowChartInternalStorage"
	ShapeTypeValuesFlowChartDocument          ShapeTypeValues = "flowChartDocument"
	ShapeTypeValuesFlowChartMultidocument     ShapeTypeValues = "flowChartMultidocument"
	ShapeTypeValuesFlowChartTerminator        ShapeTypeValues = "flowChartTerminator"
	ShapeTypeValuesFlowChartPreparation       ShapeTypeValues = "flowChartPreparation"
	ShapeTypeValuesFlowChartManualInput       ShapeTypeValues = "flowChartManualInput"
	ShapeTypeValuesFlowChartManualOperation   ShapeTypeValues = "flowChartManualOperation"
	ShapeTypeValuesFlowChartConnector         ShapeTypeValues = "flowChartConnector"
	ShapeTypeValuesFlowChartPunchedCard       ShapeTypeValues = "flowChartPunchedCard"
	ShapeTypeValuesFlowChartPunchedTape       ShapeTypeValues = "flowChartPunchedTape"
	ShapeTypeValuesFlowChartSummingJunction   ShapeTypeValues = "flowChartSummingJunction"
	ShapeTypeValuesFlowChartOr                ShapeTypeValues = "flowChartOr"
	ShapeTypeValuesFlowChartCollate           ShapeTypeValues = "flowChartCollate"
	ShapeTypeValuesFlowChartSort              ShapeTypeValues = "flowChartSort"
	ShapeTypeValuesFlowChartExtract           ShapeTypeValues = "flowChartExtract"
	ShapeTypeValuesFlowChartMerge             ShapeTypeValues = "flowChartMerge"
	ShapeTypeValuesFlowChartOfflineStorage    ShapeTypeValues = "flowChartOfflineStorage"
	ShapeTypeValuesFlowChartOnlineStorage     ShapeTypeValues = "flowChartOnlineStorage"
	ShapeTypeValuesFlowChartMagneticTape      ShapeTypeValues = "flowChartMagneticTape"
	ShapeTypeValuesFlowChartMagneticDisk      ShapeTypeValues = "flowChartMagneticDisk"
	ShapeTypeValuesFlowChartMagneticDrum      ShapeTypeValues = "flowChartMagneticDrum"
	ShapeTypeValuesFlowChartDisplay           ShapeTypeValues = "flowChartDisplay"
	ShapeTypeValuesFlowChartDelay             ShapeTypeValues = "flowChartDelay"
	ShapeTypeValuesFlowChartAlternateProcess  ShapeTypeValues = "flowChartAlternateProcess"
	ShapeTypeValuesFlowChartOffpageConnector  ShapeTypeValues = "flowChartOffpageConnector"
	ShapeTypeValuesActionButtonBlank          ShapeTypeValues = "actionButtonBlank"
	ShapeTypeValuesActionButtonHome           ShapeTypeValues = "actionButtonHome"
	ShapeTypeValuesActionButtonHelp           ShapeTypeValues = "actionButtonHelp"
	ShapeTypeValuesActionButtonInformation    ShapeTypeValues = "actionButtonInformation"
	ShapeTypeValuesActionButtonForwardNext    ShapeTypeValues = "actionButtonForwardNext"
	ShapeTypeValuesActionButtonBackPrevious   ShapeTypeValues = "actionButtonBackPrevious"
	ShapeTypeValuesActionButtonEnd            ShapeTypeValues = "actionButtonEnd"
	ShapeTypeValuesActionButtonBeginning      ShapeTypeValues = "actionButtonBeginning"
	ShapeTypeValuesActionButtonReturn         ShapeTypeValues = "actionButtonReturn"
	ShapeTypeValuesActionButtonDocument       ShapeTypeValues = "actionButtonDocument"
	ShapeTypeValuesActionButtonSound          ShapeTypeValues = "actionButtonSound"
	ShapeTypeValuesActionButtonMovie          ShapeTypeValues = "actionButtonMovie"
	ShapeTypeValuesGear6                      ShapeTypeValues = "gear6"
	ShapeTypeValuesGear9                      ShapeTypeValues = "gear9"
	ShapeTypeValuesFunnel                     ShapeTypeValues = "funnel"
	ShapeTypeValuesMathPlus                   ShapeTypeValues = "mathPlus"
	ShapeTypeValuesMathMinus                  ShapeTypeValues = "mathMinus"
	ShapeTypeValuesMathMultiply               ShapeTypeValues = "mathMultiply"
	ShapeTypeValuesMathDivide                 ShapeTypeValues = "mathDivide"
	ShapeTypeValuesMathEqual                  ShapeTypeValues = "mathEqual"
	ShapeTypeValuesMathNotEqual               ShapeTypeValues = "mathNotEqual"
	ShapeTypeValuesCornerTabs                 ShapeTypeValues = "cornerTabs"
	ShapeTypeValuesSquareTabs                 ShapeTypeValues = "squareTabs"
	ShapeTypeValuesPlaqueTabs                 ShapeTypeValues = "plaqueTabs"
	ShapeTypeValuesChartX                     ShapeTypeValues = "chartX"
	ShapeTypeValuesChartStar                  ShapeTypeValues = "chartStar"
	ShapeTypeValuesChartPlus                  ShapeTypeValues = "chartPlus"
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

// PhoneticValues
type PhoneticValues string

const (
	PhoneticValuesHalfwidthKatakana PhoneticValues = "halfwidthKatakana"
	PhoneticValuesFullwidthKatakana PhoneticValues = "fullwidthKatakana"
	PhoneticValuesHiragana          PhoneticValues = "Hiragana"
	PhoneticValuesNoConversion      PhoneticValues = "noConversion"
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

// TableValues
type TableValues string

const (
	TableValuesWorksheet  TableValues = "worksheet"
	TableValuesXml        TableValues = "xml"
	TableValuesQueryTable TableValues = "queryTable"
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

// KnownSemanticTypeValues
type KnownSemanticTypeValues string

const (
	KnownSemanticTypeValuesNone          KnownSemanticTypeValues = "none"
	KnownSemanticTypeValuesUnderline     KnownSemanticTypeValues = "underline"
	KnownSemanticTypeValuesStrikethrough KnownSemanticTypeValues = "strikethrough"
	KnownSemanticTypeValuesHighlight     KnownSemanticTypeValues = "highlight"
	KnownSemanticTypeValuesScratchOut    KnownSemanticTypeValues = "scratchOut"
	KnownSemanticTypeValuesVerticalRange KnownSemanticTypeValues = "verticalRange"
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

// PresetLineDashValues
type PresetLineDashValues string

const (
	PresetLineDashValuesSolid         PresetLineDashValues = "solid"
	PresetLineDashValuesDot           PresetLineDashValues = "dot"
	PresetLineDashValuesDash          PresetLineDashValues = "dash"
	PresetLineDashValuesLgDash        PresetLineDashValues = "lgDash"
	PresetLineDashValuesDashDot       PresetLineDashValues = "dashDot"
	PresetLineDashValuesLgDashDot     PresetLineDashValues = "lgDashDot"
	PresetLineDashValuesLgDashDotDot  PresetLineDashValues = "lgDashDotDot"
	PresetLineDashValuesSysDash       PresetLineDashValues = "sysDash"
	PresetLineDashValuesSysDot        PresetLineDashValues = "sysDot"
	PresetLineDashValuesSysDashDot    PresetLineDashValues = "sysDashDot"
	PresetLineDashValuesSysDashDotDot PresetLineDashValues = "sysDashDotDot"
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

// TextAutoNumberSchemeValues
type TextAutoNumberSchemeValues string

const (
	TextAutoNumberSchemeValuesAlphaLcParenBoth      TextAutoNumberSchemeValues = "alphaLcParenBoth"
	TextAutoNumberSchemeValuesAlphaUcParenBoth      TextAutoNumberSchemeValues = "alphaUcParenBoth"
	TextAutoNumberSchemeValuesAlphaLcParenR         TextAutoNumberSchemeValues = "alphaLcParenR"
	TextAutoNumberSchemeValuesAlphaUcParenR         TextAutoNumberSchemeValues = "alphaUcParenR"
	TextAutoNumberSchemeValuesAlphaLcPeriod         TextAutoNumberSchemeValues = "alphaLcPeriod"
	TextAutoNumberSchemeValuesAlphaUcPeriod         TextAutoNumberSchemeValues = "alphaUcPeriod"
	TextAutoNumberSchemeValuesArabicParenBoth       TextAutoNumberSchemeValues = "arabicParenBoth"
	TextAutoNumberSchemeValuesArabicParenR          TextAutoNumberSchemeValues = "arabicParenR"
	TextAutoNumberSchemeValuesArabicPeriod          TextAutoNumberSchemeValues = "arabicPeriod"
	TextAutoNumberSchemeValuesArabicPlain           TextAutoNumberSchemeValues = "arabicPlain"
	TextAutoNumberSchemeValuesRomanLcParenBoth      TextAutoNumberSchemeValues = "romanLcParenBoth"
	TextAutoNumberSchemeValuesRomanUcParenBoth      TextAutoNumberSchemeValues = "romanUcParenBoth"
	TextAutoNumberSchemeValuesRomanLcParenR         TextAutoNumberSchemeValues = "romanLcParenR"
	TextAutoNumberSchemeValuesRomanUcParenR         TextAutoNumberSchemeValues = "romanUcParenR"
	TextAutoNumberSchemeValuesRomanLcPeriod         TextAutoNumberSchemeValues = "romanLcPeriod"
	TextAutoNumberSchemeValuesRomanUcPeriod         TextAutoNumberSchemeValues = "romanUcPeriod"
	TextAutoNumberSchemeValuesCircleNumDbPlain      TextAutoNumberSchemeValues = "circleNumDbPlain"
	TextAutoNumberSchemeValuesCircleNumWdBlackPlain TextAutoNumberSchemeValues = "circleNumWdBlackPlain"
	TextAutoNumberSchemeValuesCircleNumWdWhitePlain TextAutoNumberSchemeValues = "circleNumWdWhitePlain"
	TextAutoNumberSchemeValuesArabicDbPeriod        TextAutoNumberSchemeValues = "arabicDbPeriod"
	TextAutoNumberSchemeValuesArabicDbPlain         TextAutoNumberSchemeValues = "arabicDbPlain"
	TextAutoNumberSchemeValuesEa1ChsPeriod          TextAutoNumberSchemeValues = "ea1ChsPeriod"
	TextAutoNumberSchemeValuesEa1ChsPlain           TextAutoNumberSchemeValues = "ea1ChsPlain"
	TextAutoNumberSchemeValuesEa1ChtPeriod          TextAutoNumberSchemeValues = "ea1ChtPeriod"
	TextAutoNumberSchemeValuesEa1ChtPlain           TextAutoNumberSchemeValues = "ea1ChtPlain"
	TextAutoNumberSchemeValuesEa1JpnChsDbPeriod     TextAutoNumberSchemeValues = "ea1JpnChsDbPeriod"
	TextAutoNumberSchemeValuesEa1JpnKorPlain        TextAutoNumberSchemeValues = "ea1JpnKorPlain"
	TextAutoNumberSchemeValuesEa1JpnKorPeriod       TextAutoNumberSchemeValues = "ea1JpnKorPeriod"
	TextAutoNumberSchemeValuesArabic1Minus          TextAutoNumberSchemeValues = "arabic1Minus"
	TextAutoNumberSchemeValuesArabic2Minus          TextAutoNumberSchemeValues = "arabic2Minus"
	TextAutoNumberSchemeValuesHebrew2Minus          TextAutoNumberSchemeValues = "hebrew2Minus"
	TextAutoNumberSchemeValuesThaiAlphaPeriod       TextAutoNumberSchemeValues = "thaiAlphaPeriod"
	TextAutoNumberSchemeValuesThaiAlphaParenR       TextAutoNumberSchemeValues = "thaiAlphaParenR"
	TextAutoNumberSchemeValuesThaiAlphaParenBoth    TextAutoNumberSchemeValues = "thaiAlphaParenBoth"
	TextAutoNumberSchemeValuesThaiNumPeriod         TextAutoNumberSchemeValues = "thaiNumPeriod"
	TextAutoNumberSchemeValuesThaiNumParenR         TextAutoNumberSchemeValues = "thaiNumParenR"
	TextAutoNumberSchemeValuesThaiNumParenBoth      TextAutoNumberSchemeValues = "thaiNumParenBoth"
	TextAutoNumberSchemeValuesHindiAlphaPeriod      TextAutoNumberSchemeValues = "hindiAlphaPeriod"
	TextAutoNumberSchemeValuesHindiNumPeriod        TextAutoNumberSchemeValues = "hindiNumPeriod"
	TextAutoNumberSchemeValuesHindiNumParenR        TextAutoNumberSchemeValues = "hindiNumParenR"
	TextAutoNumberSchemeValuesHindiAlpha1Period     TextAutoNumberSchemeValues = "hindiAlpha1Period"
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

// PhotoAlbumLayoutValues
type PhotoAlbumLayoutValues string

const (
	PhotoAlbumLayoutValuesFitToSlide                      PhotoAlbumLayoutValues = "fitToSlide"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues1pic      PhotoAlbumLayoutValues = "1pic"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues2pic      PhotoAlbumLayoutValues = "2pic"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues4pic      PhotoAlbumLayoutValues = "4pic"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues1picTitle PhotoAlbumLayoutValues = "1picTitle"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues2picTitle PhotoAlbumLayoutValues = "2picTitle"
	PhotoAlbumLayoutValuesPhotoAlbumLayoutValues4picTitle PhotoAlbumLayoutValues = "4picTitle"
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

// DocumentTypeValues
type DocumentTypeValues string

const (
	DocumentTypeValuesNotSpecified DocumentTypeValues = "notSpecified"
	DocumentTypeValuesLetter       DocumentTypeValues = "letter"
	DocumentTypeValuesEMail        DocumentTypeValues = "eMail"
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

// LightRigValues
type LightRigValues string

const (
	LightRigValuesLegacyFlat1   LightRigValues = "legacyFlat1"
	LightRigValuesLegacyFlat2   LightRigValues = "legacyFlat2"
	LightRigValuesLegacyFlat3   LightRigValues = "legacyFlat3"
	LightRigValuesLegacyFlat4   LightRigValues = "legacyFlat4"
	LightRigValuesLegacyNormal1 LightRigValues = "legacyNormal1"
	LightRigValuesLegacyNormal2 LightRigValues = "legacyNormal2"
	LightRigValuesLegacyNormal3 LightRigValues = "legacyNormal3"
	LightRigValuesLegacyNormal4 LightRigValues = "legacyNormal4"
	LightRigValuesLegacyHarsh1  LightRigValues = "legacyHarsh1"
	LightRigValuesLegacyHarsh2  LightRigValues = "legacyHarsh2"
	LightRigValuesLegacyHarsh3  LightRigValues = "legacyHarsh3"
	LightRigValuesLegacyHarsh4  LightRigValues = "legacyHarsh4"
	LightRigValuesThreePt       LightRigValues = "threePt"
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
	LightRigValuesTwoPt         LightRigValues = "twoPt"
	LightRigValuesGlow          LightRigValues = "glow"
	LightRigValuesBrightRoom    LightRigValues = "brightRoom"
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

// PrintErrorValues
type PrintErrorValues string

const (
	PrintErrorValuesDisplayed PrintErrorValues = "displayed"
	PrintErrorValuesBlank     PrintErrorValues = "blank"
	PrintErrorValuesDash      PrintErrorValues = "dash"
	PrintErrorValuesNA        PrintErrorValues = "NA"
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

// PaneStateValues
type PaneStateValues string

const (
	PaneStateValuesSplit       PaneStateValues = "split"
	PaneStateValuesFrozen      PaneStateValues = "frozen"
	PaneStateValuesFrozenSplit PaneStateValues = "frozenSplit"
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

// PivotAreaValues
type PivotAreaValues string

const (
	PivotAreaValuesNone     PivotAreaValues = "none"
	PivotAreaValuesNormal   PivotAreaValues = "normal"
	PivotAreaValuesData     PivotAreaValues = "data"
	PivotAreaValuesAll      PivotAreaValues = "all"
	PivotAreaValuesOrigin   PivotAreaValues = "origin"
	PivotAreaValuesButton   PivotAreaValues = "button"
	PivotAreaValuesTopRight PivotAreaValues = "topRight"
	PivotAreaValuesTopEnd   PivotAreaValues = "topEnd"
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
	HighlightColorValuesDarkBlue    HighlightColorValues = "darkBlue"
	HighlightColorValuesDarkCyan    HighlightColorValues = "darkCyan"
	HighlightColorValuesDarkGreen   HighlightColorValues = "darkGreen"
	HighlightColorValuesDarkMagenta HighlightColorValues = "darkMagenta"
	HighlightColorValuesDarkRed     HighlightColorValues = "darkRed"
	HighlightColorValuesDarkYellow  HighlightColorValues = "darkYellow"
	HighlightColorValuesDarkGray    HighlightColorValues = "darkGray"
	HighlightColorValuesLightGray   HighlightColorValues = "lightGray"
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

// PivotShowAsValues
type PivotShowAsValues string

const (
	PivotShowAsValuesPercentOfParent       PivotShowAsValues = "percentOfParent"
	PivotShowAsValuesPercentOfParentRow    PivotShowAsValues = "percentOfParentRow"
	PivotShowAsValuesPercentOfParentCol    PivotShowAsValues = "percentOfParentCol"
	PivotShowAsValuesPercentOfRunningTotal PivotShowAsValues = "percentOfRunningTotal"
	PivotShowAsValuesRankAscending         PivotShowAsValues = "rankAscending"
	PivotShowAsValuesRankDescending        PivotShowAsValues = "rankDescending"
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

// FormulaExpressionValues
type FormulaExpressionValues string

const (
	FormulaExpressionValuesRef          FormulaExpressionValues = "ref"
	FormulaExpressionValuesRefError     FormulaExpressionValues = "refError"
	FormulaExpressionValuesArea         FormulaExpressionValues = "area"
	FormulaExpressionValuesAreaError    FormulaExpressionValues = "areaError"
	FormulaExpressionValuesComputedArea FormulaExpressionValues = "computedArea"
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

// PresetZoomValues
type PresetZoomValues string

const (
	PresetZoomValuesNone     PresetZoomValues = "none"
	PresetZoomValuesFullPage PresetZoomValues = "fullPage"
	PresetZoomValuesBestFit  PresetZoomValues = "bestFit"
	PresetZoomValuesTextFit  PresetZoomValues = "textFit"
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

// BreakValues
type BreakValues string

const (
	BreakValuesPage         BreakValues = "page"
	BreakValuesColumn       BreakValues = "column"
	BreakValuesTextWrapping BreakValues = "textWrapping"
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

// DateFormatValues
type DateFormatValues string

const (
	DateFormatValuesText     DateFormatValues = "text"
	DateFormatValuesDate     DateFormatValues = "date"
	DateFormatValuesDateTime DateFormatValues = "dateTime"
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

// FootnotePositionValues
type FootnotePositionValues string

const (
	FootnotePositionValuesPageBottom  FootnotePositionValues = "pageBottom"
	FootnotePositionValuesBeneathText FootnotePositionValues = "beneathText"
	FootnotePositionValuesSectEnd     FootnotePositionValues = "sectEnd"
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

// DocPartValues
type DocPartValues string

const (
	DocPartValuesNone     DocPartValues = "none"
	DocPartValuesNormal   DocPartValues = "normal"
	DocPartValuesAutoExp  DocPartValues = "autoExp"
	DocPartValuesToolbar  DocPartValues = "toolbar"
	DocPartValuesSpeller  DocPartValues = "speller"
	DocPartValuesFormFld  DocPartValues = "formFld"
	DocPartValuesBbPlcHdr DocPartValues = "bbPlcHdr"
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

// OleChartBuildValues
type OleChartBuildValues string

const (
	OleChartBuildValuesAllAtOnce  OleChartBuildValues = "allAtOnce"
	OleChartBuildValuesSeries     OleChartBuildValues = "series"
	OleChartBuildValuesCategory   OleChartBuildValues = "category"
	OleChartBuildValuesSeriesEl   OleChartBuildValues = "seriesEl"
	OleChartBuildValuesCategoryEl OleChartBuildValues = "categoryEl"
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

// RowColumnActionValues
type RowColumnActionValues string

const (
	RowColumnActionValuesInsertRow RowColumnActionValues = "insertRow"
	RowColumnActionValuesDeleteRow RowColumnActionValues = "deleteRow"
	RowColumnActionValuesInsertCol RowColumnActionValues = "insertCol"
	RowColumnActionValuesDeleteCol RowColumnActionValues = "deleteCol"
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

// PageBorderDisplayValues
type PageBorderDisplayValues string

const (
	PageBorderDisplayValuesAllPages     PageBorderDisplayValues = "allPages"
	PageBorderDisplayValuesFirstPage    PageBorderDisplayValues = "firstPage"
	PageBorderDisplayValuesNotFirstPage PageBorderDisplayValues = "notFirstPage"
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

// TotalsRowFunctionValues
type TotalsRowFunctionValues string

const (
	TotalsRowFunctionValuesNone      TotalsRowFunctionValues = "none"
	TotalsRowFunctionValuesSum       TotalsRowFunctionValues = "sum"
	TotalsRowFunctionValuesMin       TotalsRowFunctionValues = "min"
	TotalsRowFunctionValuesMax       TotalsRowFunctionValues = "max"
	TotalsRowFunctionValuesAverage   TotalsRowFunctionValues = "average"
	TotalsRowFunctionValuesCount     TotalsRowFunctionValues = "count"
	TotalsRowFunctionValuesCountNums TotalsRowFunctionValues = "countNums"
	TotalsRowFunctionValuesStdDev    TotalsRowFunctionValues = "stdDev"
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

// SectionMarkValues
type SectionMarkValues string

const (
	SectionMarkValuesNextPage   SectionMarkValues = "nextPage"
	SectionMarkValuesNextColumn SectionMarkValues = "nextColumn"
	SectionMarkValuesContinuous SectionMarkValues = "continuous"
	SectionMarkValuesEvenPage   SectionMarkValues = "evenPage"
	SectionMarkValuesOddPage    SectionMarkValues = "oddPage"
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

// LockingValues
type LockingValues string

const (
	LockingValuesSdtLocked        LockingValues = "sdtLocked"
	LockingValuesContentLocked    LockingValues = "contentLocked"
	LockingValuesUnlocked         LockingValues = "unlocked"
	LockingValuesSdtContentLocked LockingValues = "sdtContentLocked"
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

// PivotEditValueTypeValues
type PivotEditValueTypeValues string

const (
	PivotEditValueTypeValuesNumber   PivotEditValueTypeValues = "number"
	PivotEditValueTypeValuesDateTime PivotEditValueTypeValues = "dateTime"
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

// StrictCharacterSet
type StrictCharacterSet string

const (
	StrictCharacterSetIso88591    StrictCharacterSet = "iso-8859-1"
	StrictCharacterSetMacintosh   StrictCharacterSet = "macintosh"
	StrictCharacterSetShiftJis    StrictCharacterSet = "shift_jis"
	StrictCharacterSetKsC56011987 StrictCharacterSet = "ks_c-5601-1987"
	StrictCharacterSetKSC56011992 StrictCharacterSet = "KS_C-5601-1992"
	StrictCharacterSetGBK         StrictCharacterSet = "GBK"
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

// ItemValues
type ItemValues string

const (
	ItemValuesData    ItemValues = "data"
	ItemValuesDefault ItemValues = "default"
	ItemValuesSum     ItemValues = "sum"
	ItemValuesCountA  ItemValues = "countA"
	ItemValuesAvg     ItemValues = "avg"
	ItemValuesMax     ItemValues = "max"
	ItemValuesMin     ItemValues = "min"
	ItemValuesProduct ItemValues = "product"
	ItemValuesCount   ItemValues = "count"
	ItemValuesStdDev  ItemValues = "stdDev"
	ItemValuesStdDevP ItemValues = "stdDevP"
	ItemValuesVar     ItemValues = "var"
	ItemValuesVarP    ItemValues = "varP"
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

// CalculateModeValues
type CalculateModeValues string

const (
	CalculateModeValuesManual      CalculateModeValues = "manual"
	CalculateModeValuesAuto        CalculateModeValues = "auto"
	CalculateModeValuesAutoNoTable CalculateModeValues = "autoNoTable"
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

// UnderlineValues
type UnderlineValues string

const (
	UnderlineValuesSingle          UnderlineValues = "single"
	UnderlineValuesWords           UnderlineValues = "words"
	UnderlineValuesDouble          UnderlineValues = "double"
	UnderlineValuesThick           UnderlineValues = "thick"
	UnderlineValuesDotted          UnderlineValues = "dotted"
	UnderlineValuesDottedHeavy     UnderlineValues = "dottedHeavy"
	UnderlineValuesDash            UnderlineValues = "dash"
	UnderlineValuesDashedHeavy     UnderlineValues = "dashedHeavy"
	UnderlineValuesDashLong        UnderlineValues = "dashLong"
	UnderlineValuesDashLongHeavy   UnderlineValues = "dashLongHeavy"
	UnderlineValuesDotDash         UnderlineValues = "dotDash"
	UnderlineValuesDashDotHeavy    UnderlineValues = "dashDotHeavy"
	UnderlineValuesDotDotDash      UnderlineValues = "dotDotDash"
	UnderlineValuesDashDotDotHeavy UnderlineValues = "dashDotDotHeavy"
	UnderlineValuesWave            UnderlineValues = "wave"
	UnderlineValuesWavyHeavy       UnderlineValues = "wavyHeavy"
	UnderlineValuesWavyDouble      UnderlineValues = "wavyDouble"
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

// RestartNumberValues
type RestartNumberValues string

const (
	RestartNumberValuesContinuous RestartNumberValues = "continuous"
	RestartNumberValuesEachSect   RestartNumberValues = "eachSect"
	RestartNumberValuesEachPage   RestartNumberValues = "eachPage"
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

// TimelineStyleType
type TimelineStyleType string

const (
	TimelineStyleTypeSelectionLabel         TimelineStyleType = "selectionLabel"
	TimelineStyleTypeTimeLevel              TimelineStyleType = "timeLevel"
	TimelineStyleTypePeriodLabel1           TimelineStyleType = "periodLabel1"
	TimelineStyleTypePeriodLabel2           TimelineStyleType = "periodLabel2"
	TimelineStyleTypeSelectedTimeBlock      TimelineStyleType = "selectedTimeBlock"
	TimelineStyleTypeUnselectedTimeBlock    TimelineStyleType = "unselectedTimeBlock"
	TimelineStyleTypeSelectedTimeBlockSpace TimelineStyleType = "selectedTimeBlockSpace"
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

// BlackWhiteModeValues
type BlackWhiteModeValues string

const (
	BlackWhiteModeValuesClr        BlackWhiteModeValues = "clr"
	BlackWhiteModeValuesAuto       BlackWhiteModeValues = "auto"
	BlackWhiteModeValuesGray       BlackWhiteModeValues = "gray"
	BlackWhiteModeValuesLtGray     BlackWhiteModeValues = "ltGray"
	BlackWhiteModeValuesInvGray    BlackWhiteModeValues = "invGray"
	BlackWhiteModeValuesGrayWhite  BlackWhiteModeValues = "grayWhite"
	BlackWhiteModeValuesBlackGray  BlackWhiteModeValues = "blackGray"
	BlackWhiteModeValuesBlackWhite BlackWhiteModeValues = "blackWhite"
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

// PresetColorValues
type PresetColorValues string

const (
	PresetColorValuesAliceBlue            PresetColorValues = "aliceBlue"
	PresetColorValuesAntiqueWhite         PresetColorValues = "antiqueWhite"
	PresetColorValuesAqua                 PresetColorValues = "aqua"
	PresetColorValuesAquamarine           PresetColorValues = "aquamarine"
	PresetColorValuesAzure                PresetColorValues = "azure"
	PresetColorValuesBeige                PresetColorValues = "beige"
	PresetColorValuesBisque               PresetColorValues = "bisque"
	PresetColorValuesBlack                PresetColorValues = "black"
	PresetColorValuesBlanchedAlmond       PresetColorValues = "blanchedAlmond"
	PresetColorValuesBlue                 PresetColorValues = "blue"
	PresetColorValuesBlueViolet           PresetColorValues = "blueViolet"
	PresetColorValuesBrown                PresetColorValues = "brown"
	PresetColorValuesBurlyWood            PresetColorValues = "burlyWood"
	PresetColorValuesCadetBlue            PresetColorValues = "cadetBlue"
	PresetColorValuesChartreuse           PresetColorValues = "chartreuse"
	PresetColorValuesChocolate            PresetColorValues = "chocolate"
	PresetColorValuesCoral                PresetColorValues = "coral"
	PresetColorValuesCornflowerBlue       PresetColorValues = "cornflowerBlue"
	PresetColorValuesCornsilk             PresetColorValues = "cornsilk"
	PresetColorValuesCrimson              PresetColorValues = "crimson"
	PresetColorValuesCyan                 PresetColorValues = "cyan"
	PresetColorValuesDkBlue               PresetColorValues = "dkBlue"
	PresetColorValuesDkCyan               PresetColorValues = "dkCyan"
	PresetColorValuesDkGoldenrod          PresetColorValues = "dkGoldenrod"
	PresetColorValuesDkGray               PresetColorValues = "dkGray"
	PresetColorValuesDkGreen              PresetColorValues = "dkGreen"
	PresetColorValuesDkKhaki              PresetColorValues = "dkKhaki"
	PresetColorValuesDkMagenta            PresetColorValues = "dkMagenta"
	PresetColorValuesDkOliveGreen         PresetColorValues = "dkOliveGreen"
	PresetColorValuesDkOrange             PresetColorValues = "dkOrange"
	PresetColorValuesDkOrchid             PresetColorValues = "dkOrchid"
	PresetColorValuesDkRed                PresetColorValues = "dkRed"
	PresetColorValuesDkSalmon             PresetColorValues = "dkSalmon"
	PresetColorValuesDkSeaGreen           PresetColorValues = "dkSeaGreen"
	PresetColorValuesDkSlateBlue          PresetColorValues = "dkSlateBlue"
	PresetColorValuesDkSlateGray          PresetColorValues = "dkSlateGray"
	PresetColorValuesDkTurquoise          PresetColorValues = "dkTurquoise"
	PresetColorValuesDkViolet             PresetColorValues = "dkViolet"
	PresetColorValuesDeepPink             PresetColorValues = "deepPink"
	PresetColorValuesDeepSkyBlue          PresetColorValues = "deepSkyBlue"
	PresetColorValuesDimGray              PresetColorValues = "dimGray"
	PresetColorValuesDodgerBlue           PresetColorValues = "dodgerBlue"
	PresetColorValuesFirebrick            PresetColorValues = "firebrick"
	PresetColorValuesFloralWhite          PresetColorValues = "floralWhite"
	PresetColorValuesForestGreen          PresetColorValues = "forestGreen"
	PresetColorValuesFuchsia              PresetColorValues = "fuchsia"
	PresetColorValuesGainsboro            PresetColorValues = "gainsboro"
	PresetColorValuesGhostWhite           PresetColorValues = "ghostWhite"
	PresetColorValuesGold                 PresetColorValues = "gold"
	PresetColorValuesGoldenrod            PresetColorValues = "goldenrod"
	PresetColorValuesGray                 PresetColorValues = "gray"
	PresetColorValuesGreen                PresetColorValues = "green"
	PresetColorValuesGreenYellow          PresetColorValues = "greenYellow"
	PresetColorValuesHoneydew             PresetColorValues = "honeydew"
	PresetColorValuesHotPink              PresetColorValues = "hotPink"
	PresetColorValuesIndianRed            PresetColorValues = "indianRed"
	PresetColorValuesIndigo               PresetColorValues = "indigo"
	PresetColorValuesIvory                PresetColorValues = "ivory"
	PresetColorValuesKhaki                PresetColorValues = "khaki"
	PresetColorValuesLavender             PresetColorValues = "lavender"
	PresetColorValuesLavenderBlush        PresetColorValues = "lavenderBlush"
	PresetColorValuesLawnGreen            PresetColorValues = "lawnGreen"
	PresetColorValuesLemonChiffon         PresetColorValues = "lemonChiffon"
	PresetColorValuesLtBlue               PresetColorValues = "ltBlue"
	PresetColorValuesLtCoral              PresetColorValues = "ltCoral"
	PresetColorValuesLtCyan               PresetColorValues = "ltCyan"
	PresetColorValuesLtGoldenrodYellow    PresetColorValues = "ltGoldenrodYellow"
	PresetColorValuesLtGray               PresetColorValues = "ltGray"
	PresetColorValuesLtGreen              PresetColorValues = "ltGreen"
	PresetColorValuesLtPink               PresetColorValues = "ltPink"
	PresetColorValuesLtSalmon             PresetColorValues = "ltSalmon"
	PresetColorValuesLtSeaGreen           PresetColorValues = "ltSeaGreen"
	PresetColorValuesLtSkyBlue            PresetColorValues = "ltSkyBlue"
	PresetColorValuesLtSlateGray          PresetColorValues = "ltSlateGray"
	PresetColorValuesLtSteelBlue          PresetColorValues = "ltSteelBlue"
	PresetColorValuesLtYellow             PresetColorValues = "ltYellow"
	PresetColorValuesLime                 PresetColorValues = "lime"
	PresetColorValuesLimeGreen            PresetColorValues = "limeGreen"
	PresetColorValuesLinen                PresetColorValues = "linen"
	PresetColorValuesMagenta              PresetColorValues = "magenta"
	PresetColorValuesMaroon               PresetColorValues = "maroon"
	PresetColorValuesMedAquamarine        PresetColorValues = "medAquamarine"
	PresetColorValuesMedBlue              PresetColorValues = "medBlue"
	PresetColorValuesMedOrchid            PresetColorValues = "medOrchid"
	PresetColorValuesMedPurple            PresetColorValues = "medPurple"
	PresetColorValuesMedSeaGreen          PresetColorValues = "medSeaGreen"
	PresetColorValuesMedSlateBlue         PresetColorValues = "medSlateBlue"
	PresetColorValuesMedSpringGreen       PresetColorValues = "medSpringGreen"
	PresetColorValuesMedTurquoise         PresetColorValues = "medTurquoise"
	PresetColorValuesMedVioletRed         PresetColorValues = "medVioletRed"
	PresetColorValuesMidnightBlue         PresetColorValues = "midnightBlue"
	PresetColorValuesMintCream            PresetColorValues = "mintCream"
	PresetColorValuesMistyRose            PresetColorValues = "mistyRose"
	PresetColorValuesMoccasin             PresetColorValues = "moccasin"
	PresetColorValuesNavajoWhite          PresetColorValues = "navajoWhite"
	PresetColorValuesNavy                 PresetColorValues = "navy"
	PresetColorValuesOldLace              PresetColorValues = "oldLace"
	PresetColorValuesOlive                PresetColorValues = "olive"
	PresetColorValuesOliveDrab            PresetColorValues = "oliveDrab"
	PresetColorValuesOrange               PresetColorValues = "orange"
	PresetColorValuesOrangeRed            PresetColorValues = "orangeRed"
	PresetColorValuesOrchid               PresetColorValues = "orchid"
	PresetColorValuesPaleGoldenrod        PresetColorValues = "paleGoldenrod"
	PresetColorValuesPaleGreen            PresetColorValues = "paleGreen"
	PresetColorValuesPaleTurquoise        PresetColorValues = "paleTurquoise"
	PresetColorValuesPaleVioletRed        PresetColorValues = "paleVioletRed"
	PresetColorValuesPapayaWhip           PresetColorValues = "papayaWhip"
	PresetColorValuesPeachPuff            PresetColorValues = "peachPuff"
	PresetColorValuesPeru                 PresetColorValues = "peru"
	PresetColorValuesPink                 PresetColorValues = "pink"
	PresetColorValuesPlum                 PresetColorValues = "plum"
	PresetColorValuesPowderBlue           PresetColorValues = "powderBlue"
	PresetColorValuesPurple               PresetColorValues = "purple"
	PresetColorValuesRed                  PresetColorValues = "red"
	PresetColorValuesRosyBrown            PresetColorValues = "rosyBrown"
	PresetColorValuesRoyalBlue            PresetColorValues = "royalBlue"
	PresetColorValuesSaddleBrown          PresetColorValues = "saddleBrown"
	PresetColorValuesSalmon               PresetColorValues = "salmon"
	PresetColorValuesSandyBrown           PresetColorValues = "sandyBrown"
	PresetColorValuesSeaGreen             PresetColorValues = "seaGreen"
	PresetColorValuesSeaShell             PresetColorValues = "seaShell"
	PresetColorValuesSienna               PresetColorValues = "sienna"
	PresetColorValuesSilver               PresetColorValues = "silver"
	PresetColorValuesSkyBlue              PresetColorValues = "skyBlue"
	PresetColorValuesSlateBlue            PresetColorValues = "slateBlue"
	PresetColorValuesSlateGray            PresetColorValues = "slateGray"
	PresetColorValuesSnow                 PresetColorValues = "snow"
	PresetColorValuesSpringGreen          PresetColorValues = "springGreen"
	PresetColorValuesSteelBlue            PresetColorValues = "steelBlue"
	PresetColorValuesTan                  PresetColorValues = "tan"
	PresetColorValuesTeal                 PresetColorValues = "teal"
	PresetColorValuesThistle              PresetColorValues = "thistle"
	PresetColorValuesTomato               PresetColorValues = "tomato"
	PresetColorValuesTurquoise            PresetColorValues = "turquoise"
	PresetColorValuesViolet               PresetColorValues = "violet"
	PresetColorValuesWheat                PresetColorValues = "wheat"
	PresetColorValuesWhite                PresetColorValues = "white"
	PresetColorValuesWhiteSmoke           PresetColorValues = "whiteSmoke"
	PresetColorValuesYellow               PresetColorValues = "yellow"
	PresetColorValuesYellowGreen          PresetColorValues = "yellowGreen"
	PresetColorValuesDarkBlue             PresetColorValues = "darkBlue"
	PresetColorValuesDarkCyan             PresetColorValues = "darkCyan"
	PresetColorValuesDarkGoldenrod        PresetColorValues = "darkGoldenrod"
	PresetColorValuesDarkGray             PresetColorValues = "darkGray"
	PresetColorValuesDarkGrey             PresetColorValues = "darkGrey"
	PresetColorValuesDarkGreen            PresetColorValues = "darkGreen"
	PresetColorValuesDarkKhaki            PresetColorValues = "darkKhaki"
	PresetColorValuesDarkMagenta          PresetColorValues = "darkMagenta"
	PresetColorValuesDarkOliveGreen       PresetColorValues = "darkOliveGreen"
	PresetColorValuesDarkOrange           PresetColorValues = "darkOrange"
	PresetColorValuesDarkOrchid           PresetColorValues = "darkOrchid"
	PresetColorValuesDarkRed              PresetColorValues = "darkRed"
	PresetColorValuesDarkSalmon           PresetColorValues = "darkSalmon"
	PresetColorValuesDarkSeaGreen         PresetColorValues = "darkSeaGreen"
	PresetColorValuesDarkSlateBlue        PresetColorValues = "darkSlateBlue"
	PresetColorValuesDarkSlateGray        PresetColorValues = "darkSlateGray"
	PresetColorValuesDarkSlateGrey        PresetColorValues = "darkSlateGrey"
	PresetColorValuesDarkTurquoise        PresetColorValues = "darkTurquoise"
	PresetColorValuesDarkViolet           PresetColorValues = "darkViolet"
	PresetColorValuesLightBlue            PresetColorValues = "lightBlue"
	PresetColorValuesLightCoral           PresetColorValues = "lightCoral"
	PresetColorValuesLightCyan            PresetColorValues = "lightCyan"
	PresetColorValuesLightGoldenrodYellow PresetColorValues = "lightGoldenrodYellow"
	PresetColorValuesLightGray            PresetColorValues = "lightGray"
	PresetColorValuesLightGrey            PresetColorValues = "lightGrey"
	PresetColorValuesLightGreen           PresetColorValues = "lightGreen"
	PresetColorValuesLightPink            PresetColorValues = "lightPink"
	PresetColorValuesLightSalmon          PresetColorValues = "lightSalmon"
	PresetColorValuesLightSeaGreen        PresetColorValues = "lightSeaGreen"
	PresetColorValuesLightSkyBlue         PresetColorValues = "lightSkyBlue"
	PresetColorValuesLightSlateGray       PresetColorValues = "lightSlateGray"
	PresetColorValuesLightSlateGrey       PresetColorValues = "lightSlateGrey"
	PresetColorValuesLightSteelBlue       PresetColorValues = "lightSteelBlue"
	PresetColorValuesLightYellow          PresetColorValues = "lightYellow"
	PresetColorValuesMediumAquamarine     PresetColorValues = "mediumAquamarine"
	PresetColorValuesMediumBlue           PresetColorValues = "mediumBlue"
	PresetColorValuesMediumOrchid         PresetColorValues = "mediumOrchid"
	PresetColorValuesMediumPurple         PresetColorValues = "mediumPurple"
	PresetColorValuesMediumSeaGreen       PresetColorValues = "mediumSeaGreen"
	PresetColorValuesMediumSlateBlue      PresetColorValues = "mediumSlateBlue"
	PresetColorValuesMediumSpringGreen    PresetColorValues = "mediumSpringGreen"
	PresetColorValuesMediumTurquoise      PresetColorValues = "mediumTurquoise"
	PresetColorValuesMediumVioletRed      PresetColorValues = "mediumVioletRed"
	PresetColorValuesDkGrey               PresetColorValues = "dkGrey"
	PresetColorValuesDimGrey              PresetColorValues = "dimGrey"
	PresetColorValuesDkSlateGrey          PresetColorValues = "dkSlateGrey"
	PresetColorValuesGrey                 PresetColorValues = "grey"
	PresetColorValuesLtGrey               PresetColorValues = "ltGrey"
	PresetColorValuesLtSlateGrey          PresetColorValues = "ltSlateGrey"
	PresetColorValuesSlateGrey            PresetColorValues = "slateGrey"
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

// FilterOperatorValues
type FilterOperatorValues string

const (
	FilterOperatorValuesEqual              FilterOperatorValues = "equal"
	FilterOperatorValuesLessThan           FilterOperatorValues = "lessThan"
	FilterOperatorValuesLessThanOrEqual    FilterOperatorValues = "lessThanOrEqual"
	FilterOperatorValuesNotEqual           FilterOperatorValues = "notEqual"
	FilterOperatorValuesGreaterThanOrEqual FilterOperatorValues = "greaterThanOrEqual"
	FilterOperatorValuesGreaterThan        FilterOperatorValues = "greaterThan"
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

// CharacterSpacingValues
type CharacterSpacingValues string

const (
	CharacterSpacingValuesDoNotCompress                      CharacterSpacingValues = "doNotCompress"
	CharacterSpacingValuesCompressPunctuation                CharacterSpacingValues = "compressPunctuation"
	CharacterSpacingValuesCompressPunctuationAndJapaneseKana CharacterSpacingValues = "compressPunctuationAndJapaneseKana"
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

// CryptAlgorithmValues
type CryptAlgorithmValues string

const (
	CryptAlgorithmValuesTypeAny CryptAlgorithmValues = "typeAny"
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

// SheetStateValues
type SheetStateValues string

const (
	SheetStateValuesVisible    SheetStateValues = "visible"
	SheetStateValuesHidden     SheetStateValues = "hidden"
	SheetStateValuesVeryHidden SheetStateValues = "veryHidden"
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

// TextDirectionValues
type TextDirectionValues string

const (
	TextDirectionValuesLrTb  TextDirectionValues = "lrTb"
	TextDirectionValuesTb    TextDirectionValues = "tb"
	TextDirectionValuesTbRl  TextDirectionValues = "tbRl"
	TextDirectionValuesRl    TextDirectionValues = "rl"
	TextDirectionValuesBtLr  TextDirectionValues = "btLr"
	TextDirectionValuesLr    TextDirectionValues = "lr"
	TextDirectionValuesLrTbV TextDirectionValues = "lrTbV"
	TextDirectionValuesTbV   TextDirectionValues = "tbV"
	TextDirectionValuesTbRlV TextDirectionValues = "tbRlV"
	TextDirectionValuesRlV   TextDirectionValues = "rlV"
	TextDirectionValuesTbLrV TextDirectionValues = "tbLrV"
	TextDirectionValuesLrV   TextDirectionValues = "lrV"
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

// PhoneticAlignmentValues
type PhoneticAlignmentValues string

const (
	PhoneticAlignmentValuesNoControl   PhoneticAlignmentValues = "noControl"
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

// ProofingErrorValues
type ProofingErrorValues string

const (
	ProofingErrorValuesSpellStart ProofingErrorValues = "spellStart"
	ProofingErrorValuesSpellEnd   ProofingErrorValues = "spellEnd"
	ProofingErrorValuesGramStart  ProofingErrorValues = "gramStart"
	ProofingErrorValuesGramEnd    ProofingErrorValues = "gramEnd"
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

// JustificationValues
type JustificationValues string

const (
	JustificationValuesLeft           JustificationValues = "left"
	JustificationValuesStart          JustificationValues = "start"
	JustificationValuesCenter         JustificationValues = "center"
	JustificationValuesRight          JustificationValues = "right"
	JustificationValuesEnd            JustificationValues = "end"
	JustificationValuesBoth           JustificationValues = "both"
	JustificationValuesMediumKashida  JustificationValues = "mediumKashida"
	JustificationValuesDistribute     JustificationValues = "distribute"
	JustificationValuesNumTab         JustificationValues = "numTab"
	JustificationValuesHighKashida    JustificationValues = "highKashida"
	JustificationValuesLowKashida     JustificationValues = "lowKashida"
	JustificationValuesThaiDistribute JustificationValues = "thaiDistribute"
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

// TextShapeValues
type TextShapeValues string

const (
	TextShapeValuesTextNoShape               TextShapeValues = "textNoShape"
	TextShapeValuesTextPlain                 TextShapeValues = "textPlain"
	TextShapeValuesTextStop                  TextShapeValues = "textStop"
	TextShapeValuesTextTriangle              TextShapeValues = "textTriangle"
	TextShapeValuesTextTriangleInverted      TextShapeValues = "textTriangleInverted"
	TextShapeValuesTextChevron               TextShapeValues = "textChevron"
	TextShapeValuesTextChevronInverted       TextShapeValues = "textChevronInverted"
	TextShapeValuesTextRingInside            TextShapeValues = "textRingInside"
	TextShapeValuesTextRingOutside           TextShapeValues = "textRingOutside"
	TextShapeValuesTextArchUp                TextShapeValues = "textArchUp"
	TextShapeValuesTextArchDown              TextShapeValues = "textArchDown"
	TextShapeValuesTextCircle                TextShapeValues = "textCircle"
	TextShapeValuesTextButton                TextShapeValues = "textButton"
	TextShapeValuesTextArchUpPour            TextShapeValues = "textArchUpPour"
	TextShapeValuesTextArchDownPour          TextShapeValues = "textArchDownPour"
	TextShapeValuesTextCirclePour            TextShapeValues = "textCirclePour"
	TextShapeValuesTextButtonPour            TextShapeValues = "textButtonPour"
	TextShapeValuesTextCurveUp               TextShapeValues = "textCurveUp"
	TextShapeValuesTextCurveDown             TextShapeValues = "textCurveDown"
	TextShapeValuesTextCanUp                 TextShapeValues = "textCanUp"
	TextShapeValuesTextCanDown               TextShapeValues = "textCanDown"
	TextShapeValuesTextWave1                 TextShapeValues = "textWave1"
	TextShapeValuesTextWave2                 TextShapeValues = "textWave2"
	TextShapeValuesTextDoubleWave1           TextShapeValues = "textDoubleWave1"
	TextShapeValuesTextWave4                 TextShapeValues = "textWave4"
	TextShapeValuesTextInflate               TextShapeValues = "textInflate"
	TextShapeValuesTextDeflate               TextShapeValues = "textDeflate"
	TextShapeValuesTextInflateBottom         TextShapeValues = "textInflateBottom"
	TextShapeValuesTextDeflateBottom         TextShapeValues = "textDeflateBottom"
	TextShapeValuesTextInflateTop            TextShapeValues = "textInflateTop"
	TextShapeValuesTextDeflateTop            TextShapeValues = "textDeflateTop"
	TextShapeValuesTextDeflateInflate        TextShapeValues = "textDeflateInflate"
	TextShapeValuesTextDeflateInflateDeflate TextShapeValues = "textDeflateInflateDeflate"
	TextShapeValuesTextFadeRight             TextShapeValues = "textFadeRight"
	TextShapeValuesTextFadeLeft              TextShapeValues = "textFadeLeft"
	TextShapeValuesTextFadeUp                TextShapeValues = "textFadeUp"
	TextShapeValuesTextFadeDown              TextShapeValues = "textFadeDown"
	TextShapeValuesTextSlantUp               TextShapeValues = "textSlantUp"
	TextShapeValuesTextSlantDown             TextShapeValues = "textSlantDown"
	TextShapeValuesTextCascadeUp             TextShapeValues = "textCascadeUp"
	TextShapeValuesTextCascadeDown           TextShapeValues = "textCascadeDown"
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

// TimeNodeRestartValues
type TimeNodeRestartValues string

const (
	TimeNodeRestartValuesAlways        TimeNodeRestartValues = "always"
	TimeNodeRestartValuesWhenNotActive TimeNodeRestartValues = "whenNotActive"
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

// PreviousActionValues
type PreviousActionValues string

const (
	PreviousActionValuesNone      PreviousActionValues = "none"
	PreviousActionValuesSkipTimed PreviousActionValues = "skipTimed"
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

// SlideLayoutValues
type SlideLayoutValues string

const (
	SlideLayoutValuesTitle                   SlideLayoutValues = "title"
	SlideLayoutValuesTx                      SlideLayoutValues = "tx"
	SlideLayoutValuesTwoColTx                SlideLayoutValues = "twoColTx"
	SlideLayoutValuesTbl                     SlideLayoutValues = "tbl"
	SlideLayoutValuesTxAndChart              SlideLayoutValues = "txAndChart"
	SlideLayoutValuesChartAndTx              SlideLayoutValues = "chartAndTx"
	SlideLayoutValuesDgm                     SlideLayoutValues = "dgm"
	SlideLayoutValuesChart                   SlideLayoutValues = "chart"
	SlideLayoutValuesTxAndClipArt            SlideLayoutValues = "txAndClipArt"
	SlideLayoutValuesClipArtAndTx            SlideLayoutValues = "clipArtAndTx"
	SlideLayoutValuesTitleOnly               SlideLayoutValues = "titleOnly"
	SlideLayoutValuesBlank                   SlideLayoutValues = "blank"
	SlideLayoutValuesTxAndObj                SlideLayoutValues = "txAndObj"
	SlideLayoutValuesObjAndTx                SlideLayoutValues = "objAndTx"
	SlideLayoutValuesObjOnly                 SlideLayoutValues = "objOnly"
	SlideLayoutValuesObj                     SlideLayoutValues = "obj"
	SlideLayoutValuesTxAndMedia              SlideLayoutValues = "txAndMedia"
	SlideLayoutValuesMediaAndTx              SlideLayoutValues = "mediaAndTx"
	SlideLayoutValuesObjOverTx               SlideLayoutValues = "objOverTx"
	SlideLayoutValuesTxOverObj               SlideLayoutValues = "txOverObj"
	SlideLayoutValuesTxAndTwoObj             SlideLayoutValues = "txAndTwoObj"
	SlideLayoutValuesTwoObjAndTx             SlideLayoutValues = "twoObjAndTx"
	SlideLayoutValuesTwoObjOverTx            SlideLayoutValues = "twoObjOverTx"
	SlideLayoutValuesFourObj                 SlideLayoutValues = "fourObj"
	SlideLayoutValuesVertTx                  SlideLayoutValues = "vertTx"
	SlideLayoutValuesClipArtAndVertTx        SlideLayoutValues = "clipArtAndVertTx"
	SlideLayoutValuesVertTitleAndTx          SlideLayoutValues = "vertTitleAndTx"
	SlideLayoutValuesVertTitleAndTxOverChart SlideLayoutValues = "vertTitleAndTxOverChart"
	SlideLayoutValuesTwoObj                  SlideLayoutValues = "twoObj"
	SlideLayoutValuesObjAndTwoObj            SlideLayoutValues = "objAndTwoObj"
	SlideLayoutValuesTwoObjAndObj            SlideLayoutValues = "twoObjAndObj"
	SlideLayoutValuesCust                    SlideLayoutValues = "cust"
	SlideLayoutValuesSecHead                 SlideLayoutValues = "secHead"
	SlideLayoutValuesTwoTxTwoObj             SlideLayoutValues = "twoTxTwoObj"
	SlideLayoutValuesObjTx                   SlideLayoutValues = "objTx"
	SlideLayoutValuesPicTx                   SlideLayoutValues = "picTx"
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

// ViewValues
type ViewValues string

const (
	ViewValuesNone        ViewValues = "none"
	ViewValuesPrint       ViewValues = "print"
	ViewValuesOutline     ViewValues = "outline"
	ViewValuesMasterPages ViewValues = "masterPages"
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

// CommentsValues
type CommentsValues string

const (
	CommentsValuesCommNone          CommentsValues = "commNone"
	CommentsValuesCommIndicator     CommentsValues = "commIndicator"
	CommentsValuesCommIndAndComment CommentsValues = "commIndAndComment"
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

// DocPartGalleryValues
type DocPartGalleryValues string

const (
	DocPartGalleryValuesPlaceholder       DocPartGalleryValues = "placeholder"
	DocPartGalleryValuesAny               DocPartGalleryValues = "any"
	DocPartGalleryValuesDefault           DocPartGalleryValues = "default"
	DocPartGalleryValuesDocParts          DocPartGalleryValues = "docParts"
	DocPartGalleryValuesCoverPg           DocPartGalleryValues = "coverPg"
	DocPartGalleryValuesEq                DocPartGalleryValues = "eq"
	DocPartGalleryValuesFtrs              DocPartGalleryValues = "ftrs"
	DocPartGalleryValuesHdrs              DocPartGalleryValues = "hdrs"
	DocPartGalleryValuesPgNum             DocPartGalleryValues = "pgNum"
	DocPartGalleryValuesTbls              DocPartGalleryValues = "tbls"
	DocPartGalleryValuesWatermarks        DocPartGalleryValues = "watermarks"
	DocPartGalleryValuesAutoTxt           DocPartGalleryValues = "autoTxt"
	DocPartGalleryValuesTxtBox            DocPartGalleryValues = "txtBox"
	DocPartGalleryValuesPgNumT            DocPartGalleryValues = "pgNumT"
	DocPartGalleryValuesPgNumB            DocPartGalleryValues = "pgNumB"
	DocPartGalleryValuesPgNumMargins      DocPartGalleryValues = "pgNumMargins"
	DocPartGalleryValuesTblOfContents     DocPartGalleryValues = "tblOfContents"
	DocPartGalleryValuesBib               DocPartGalleryValues = "bib"
	DocPartGalleryValuesCustQuickParts    DocPartGalleryValues = "custQuickParts"
	DocPartGalleryValuesCustCoverPg       DocPartGalleryValues = "custCoverPg"
	DocPartGalleryValuesCustEq            DocPartGalleryValues = "custEq"
	DocPartGalleryValuesCustFtrs          DocPartGalleryValues = "custFtrs"
	DocPartGalleryValuesCustHdrs          DocPartGalleryValues = "custHdrs"
	DocPartGalleryValuesCustPgNum         DocPartGalleryValues = "custPgNum"
	DocPartGalleryValuesCustTbls          DocPartGalleryValues = "custTbls"
	DocPartGalleryValuesCustWatermarks    DocPartGalleryValues = "custWatermarks"
	DocPartGalleryValuesCustAutoTxt       DocPartGalleryValues = "custAutoTxt"
	DocPartGalleryValuesCustTxtBox        DocPartGalleryValues = "custTxtBox"
	DocPartGalleryValuesCustPgNumT        DocPartGalleryValues = "custPgNumT"
	DocPartGalleryValuesCustPgNumB        DocPartGalleryValues = "custPgNumB"
	DocPartGalleryValuesCustPgNumMargins  DocPartGalleryValues = "custPgNumMargins"
	DocPartGalleryValuesCustTblOfContents DocPartGalleryValues = "custTblOfContents"
	DocPartGalleryValuesCustBib           DocPartGalleryValues = "custBib"
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

// SlicerStyleTypeValues
type SlicerStyleTypeValues string

const (
	SlicerStyleTypeValuesUnselectedItemWithData          SlicerStyleTypeValues = "unselectedItemWithData"
	SlicerStyleTypeValuesSelectedItemWithData            SlicerStyleTypeValues = "selectedItemWithData"
	SlicerStyleTypeValuesUnselectedItemWithNoData        SlicerStyleTypeValues = "unselectedItemWithNoData"
	SlicerStyleTypeValuesSelectedItemWithNoData          SlicerStyleTypeValues = "selectedItemWithNoData"
	SlicerStyleTypeValuesHoveredUnselectedItemWithData   SlicerStyleTypeValues = "hoveredUnselectedItemWithData"
	SlicerStyleTypeValuesHoveredSelectedItemWithData     SlicerStyleTypeValues = "hoveredSelectedItemWithData"
	SlicerStyleTypeValuesHoveredUnselectedItemWithNoData SlicerStyleTypeValues = "hoveredUnselectedItemWithNoData"
	SlicerStyleTypeValuesHoveredSelectedItemWithNoData   SlicerStyleTypeValues = "hoveredSelectedItemWithNoData"
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

// TextAlignmentTypeValues
type TextAlignmentTypeValues string

const (
	TextAlignmentTypeValuesL        TextAlignmentTypeValues = "l"
	TextAlignmentTypeValuesCtr      TextAlignmentTypeValues = "ctr"
	TextAlignmentTypeValuesR        TextAlignmentTypeValues = "r"
	TextAlignmentTypeValuesJust     TextAlignmentTypeValues = "just"
	TextAlignmentTypeValuesJustLow  TextAlignmentTypeValues = "justLow"
	TextAlignmentTypeValuesDist     TextAlignmentTypeValues = "dist"
	TextAlignmentTypeValuesThaiDist TextAlignmentTypeValues = "thaiDist"
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

// MailMergeOdsoFieldValues
type MailMergeOdsoFieldValues string

const (
	MailMergeOdsoFieldValuesNull     MailMergeOdsoFieldValues = "null"
	MailMergeOdsoFieldValuesDbColumn MailMergeOdsoFieldValues = "dbColumn"
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

// ReferenceModeValues
type ReferenceModeValues string

const (
	ReferenceModeValuesA1   ReferenceModeValues = "A1"
	ReferenceModeValuesR1C1 ReferenceModeValues = "R1C1"
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

// TargetScreenSize
type TargetScreenSize string

const (
	TargetScreenSizeTargetScreenSize544x376   TargetScreenSize = "544x376"
	TargetScreenSizeTargetScreenSize640x480   TargetScreenSize = "640x480"
	TargetScreenSizeTargetScreenSize720x512   TargetScreenSize = "720x512"
	TargetScreenSizeTargetScreenSize800x600   TargetScreenSize = "800x600"
	TargetScreenSizeTargetScreenSize1024x768  TargetScreenSize = "1024x768"
	TargetScreenSizeTargetScreenSize1152x882  TargetScreenSize = "1152x882"
	TargetScreenSizeTargetScreenSize1152x900  TargetScreenSize = "1152x900"
	TargetScreenSizeTargetScreenSize1280x1024 TargetScreenSize = "1280x1024"
	TargetScreenSizeTargetScreenSize1600x1200 TargetScreenSize = "1600x1200"
	TargetScreenSizeTargetScreenSize1800x1440 TargetScreenSize = "1800x1440"
	TargetScreenSizeTargetScreenSize1920x1200 TargetScreenSize = "1920x1200"
)

func (e TargetScreenSize) MarshalXMLAttr(
	name xml.Name,
) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: string(e),
	}, nil
}

func (e *TargetScreenSize) UnmarshalXMLAttr(
	attr xml.Attr,
) error {
	*e = TargetScreenSize(attr.Value)
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

// WebSourceValues
type WebSourceValues string

const (
	WebSourceValuesSheet      WebSourceValues = "sheet"
	WebSourceValuesPrintArea  WebSourceValues = "printArea"
	WebSourceValuesAutoFilter WebSourceValues = "autoFilter"
	WebSourceValuesRange      WebSourceValues = "range"
	WebSourceValuesChart      WebSourceValues = "chart"
	WebSourceValuesPivotTable WebSourceValues = "pivotTable"
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

// TextWrappingValues
type TextWrappingValues string

const (
	TextWrappingValuesAuto      TextWrappingValues = "auto"
	TextWrappingValuesNotBeside TextWrappingValues = "notBeside"
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
	WebColorValuesPresentationText   WebColorValues = "presentationText"
	WebColorValuesPresentationAccent WebColorValues = "presentationAccent"
	WebColorValuesWhiteTextOnBlack   WebColorValues = "whiteTextOnBlack"
	WebColorValuesBlackTextOnWhite   WebColorValues = "blackTextOnWhite"
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

// TriggerEventValues
type TriggerEventValues string

const (
	TriggerEventValuesNone            TriggerEventValues = "none"
	TriggerEventValuesOnBegin         TriggerEventValues = "onBegin"
	TriggerEventValuesOnEnd           TriggerEventValues = "onEnd"
	TriggerEventValuesBegin           TriggerEventValues = "begin"
	TriggerEventValuesEnd             TriggerEventValues = "end"
	TriggerEventValuesOnClick         TriggerEventValues = "onClick"
	TriggerEventValuesOnDblClick      TriggerEventValues = "onDblClick"
	TriggerEventValuesOnMouseOver     TriggerEventValues = "onMouseOver"
	TriggerEventValuesOnMouseOut      TriggerEventValues = "onMouseOut"
	TriggerEventValuesOnNext          TriggerEventValues = "onNext"
	TriggerEventValuesOnPrev          TriggerEventValues = "onPrev"
	TriggerEventValuesOnStopAudio     TriggerEventValues = "onStopAudio"
	TriggerEventValuesOnMediaBookmark TriggerEventValues = "onMediaBookmark"
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

// ShowDataAsValues
type ShowDataAsValues string

const (
	ShowDataAsValuesNormal         ShowDataAsValues = "normal"
	ShowDataAsValuesDifference     ShowDataAsValues = "difference"
	ShowDataAsValuesPercent        ShowDataAsValues = "percent"
	ShowDataAsValuesPercentDiff    ShowDataAsValues = "percentDiff"
	ShowDataAsValuesRunTotal       ShowDataAsValues = "runTotal"
	ShowDataAsValuesPercentOfRow   ShowDataAsValues = "percentOfRow"
	ShowDataAsValuesPercentOfCol   ShowDataAsValues = "percentOfCol"
	ShowDataAsValuesPercentOfTotal ShowDataAsValues = "percentOfTotal"
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

// SheetViewValues
type SheetViewValues string

const (
	SheetViewValuesNormal           SheetViewValues = "normal"
	SheetViewValuesPageBreakPreview SheetViewValues = "pageBreakPreview"
	SheetViewValuesPageLayout       SheetViewValues = "pageLayout"
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

// DataValidationValues
type DataValidationValues string

const (
	DataValidationValuesNone       DataValidationValues = "none"
	DataValidationValuesWhole      DataValidationValues = "whole"
	DataValidationValuesDecimal    DataValidationValues = "decimal"
	DataValidationValuesList       DataValidationValues = "list"
	DataValidationValuesDate       DataValidationValues = "date"
	DataValidationValuesTime       DataValidationValues = "time"
	DataValidationValuesTextLength DataValidationValues = "textLength"
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

// VolatileDependencyValues
type VolatileDependencyValues string

const (
	VolatileDependencyValuesRealTimeData  VolatileDependencyValues = "realTimeData"
	VolatileDependencyValuesOlapFunctions VolatileDependencyValues = "olapFunctions"
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

// TextEffectValues
type TextEffectValues string

const (
	TextEffectValuesBlinkBackground TextEffectValues = "blinkBackground"
	TextEffectValuesLights          TextEffectValues = "lights"
	TextEffectValuesAntsBlack       TextEffectValues = "antsBlack"
	TextEffectValuesAntsRed         TextEffectValues = "antsRed"
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

// CellFormulaValues
type CellFormulaValues string

const (
	CellFormulaValuesNormal    CellFormulaValues = "normal"
	CellFormulaValuesArray     CellFormulaValues = "array"
	CellFormulaValuesDataTable CellFormulaValues = "dataTable"
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

// DataBarDirectionValues
type DataBarDirectionValues string

const (
	DataBarDirectionValuesContext     DataBarDirectionValues = "context"
	DataBarDirectionValuesLeftToRight DataBarDirectionValues = "leftToRight"
	DataBarDirectionValuesRightToLeft DataBarDirectionValues = "rightToLeft"
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

// TimeNodeSyncValues
type TimeNodeSyncValues string

const (
	TimeNodeSyncValuesNone    TimeNodeSyncValues = "none"
	TimeNodeSyncValuesCanSlip TimeNodeSyncValues = "canSlip"
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

// ChartBuildStepValues
type ChartBuildStepValues string

const (
	ChartBuildStepValuesCategory     ChartBuildStepValues = "category"
	ChartBuildStepValuesPtInCategory ChartBuildStepValues = "ptInCategory"
	ChartBuildStepValuesSeries       ChartBuildStepValues = "series"
	ChartBuildStepValuesPtInSeries   ChartBuildStepValues = "ptInSeries"
	ChartBuildStepValuesAllPts       ChartBuildStepValues = "allPts"
	ChartBuildStepValuesGridLegend   ChartBuildStepValues = "gridLegend"
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

// WebScreenSizeValues
type WebScreenSizeValues string

const (
	WebScreenSizeValuesWebScreenSizeValues544x376   WebScreenSizeValues = "544x376"
	WebScreenSizeValuesWebScreenSizeValues640x480   WebScreenSizeValues = "640x480"
	WebScreenSizeValuesWebScreenSizeValues720x512   WebScreenSizeValues = "720x512"
	WebScreenSizeValuesWebScreenSizeValues800x600   WebScreenSizeValues = "800x600"
	WebScreenSizeValuesWebScreenSizeValues1024x768  WebScreenSizeValues = "1024x768"
	WebScreenSizeValuesWebScreenSizeValues1152x882  WebScreenSizeValues = "1152x882"
	WebScreenSizeValuesWebScreenSizeValues1152x900  WebScreenSizeValues = "1152x900"
	WebScreenSizeValuesWebScreenSizeValues1280x1024 WebScreenSizeValues = "1280x1024"
	WebScreenSizeValuesWebScreenSizeValues1600x1200 WebScreenSizeValues = "1600x1200"
	WebScreenSizeValuesWebScreenSizeValues1800x1400 WebScreenSizeValues = "1800x1400"
	WebScreenSizeValuesWebScreenSizeValues1920x1200 WebScreenSizeValues = "1920x1200"
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
	BorderStyleValuesMediumDashed     BorderStyleValues = "mediumDashed"
	BorderStyleValuesDashDot          BorderStyleValues = "dashDot"
	BorderStyleValuesMediumDashDot    BorderStyleValues = "mediumDashDot"
	BorderStyleValuesDashDotDot       BorderStyleValues = "dashDotDot"
	BorderStyleValuesMediumDashDotDot BorderStyleValues = "mediumDashDotDot"
	BorderStyleValuesSlantDashDot     BorderStyleValues = "slantDashDot"
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

// TimeNodeValues
type TimeNodeValues string

const (
	TimeNodeValuesClickEffect    TimeNodeValues = "clickEffect"
	TimeNodeValuesWithEffect     TimeNodeValues = "withEffect"
	TimeNodeValuesAfterEffect    TimeNodeValues = "afterEffect"
	TimeNodeValuesMainSeq        TimeNodeValues = "mainSeq"
	TimeNodeValuesInteractiveSeq TimeNodeValues = "interactiveSeq"
	TimeNodeValuesClickPar       TimeNodeValues = "clickPar"
	TimeNodeValuesWithGroup      TimeNodeValues = "withGroup"
	TimeNodeValuesAfterGroup     TimeNodeValues = "afterGroup"
	TimeNodeValuesTmRoot         TimeNodeValues = "tmRoot"
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

// TextVerticalValues
type TextVerticalValues string

const (
	TextVerticalValuesHorz           TextVerticalValues = "horz"
	TextVerticalValuesVert           TextVerticalValues = "vert"
	TextVerticalValuesVert270        TextVerticalValues = "vert270"
	TextVerticalValuesWordArtVert    TextVerticalValues = "wordArtVert"
	TextVerticalValuesEaVert         TextVerticalValues = "eaVert"
	TextVerticalValuesMongolianVert  TextVerticalValues = "mongolianVert"
	TextVerticalValuesWordArtVertRtl TextVerticalValues = "wordArtVertRtl"
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

// OleObjectFollowColorSchemeValues
type OleObjectFollowColorSchemeValues string

const (
	OleObjectFollowColorSchemeValuesNone              OleObjectFollowColorSchemeValues = "none"
	OleObjectFollowColorSchemeValuesFull              OleObjectFollowColorSchemeValues = "full"
	OleObjectFollowColorSchemeValuesTextAndBackground OleObjectFollowColorSchemeValues = "textAndBackground"
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

// NumberFormatValues
type NumberFormatValues string

const (
	NumberFormatValuesDecimal                      NumberFormatValues = "decimal"
	NumberFormatValuesUpperRoman                   NumberFormatValues = "upperRoman"
	NumberFormatValuesLowerRoman                   NumberFormatValues = "lowerRoman"
	NumberFormatValuesUpperLetter                  NumberFormatValues = "upperLetter"
	NumberFormatValuesLowerLetter                  NumberFormatValues = "lowerLetter"
	NumberFormatValuesOrdinal                      NumberFormatValues = "ordinal"
	NumberFormatValuesCardinalText                 NumberFormatValues = "cardinalText"
	NumberFormatValuesOrdinalText                  NumberFormatValues = "ordinalText"
	NumberFormatValuesHex                          NumberFormatValues = "hex"
	NumberFormatValuesChicago                      NumberFormatValues = "chicago"
	NumberFormatValuesIdeographDigital             NumberFormatValues = "ideographDigital"
	NumberFormatValuesJapaneseCounting             NumberFormatValues = "japaneseCounting"
	NumberFormatValuesAiueo                        NumberFormatValues = "aiueo"
	NumberFormatValuesIroha                        NumberFormatValues = "iroha"
	NumberFormatValuesDecimalFullWidth             NumberFormatValues = "decimalFullWidth"
	NumberFormatValuesDecimalHalfWidth             NumberFormatValues = "decimalHalfWidth"
	NumberFormatValuesJapaneseLegal                NumberFormatValues = "japaneseLegal"
	NumberFormatValuesJapaneseDigitalTenThousand   NumberFormatValues = "japaneseDigitalTenThousand"
	NumberFormatValuesDecimalEnclosedCircle        NumberFormatValues = "decimalEnclosedCircle"
	NumberFormatValuesDecimalFullWidth2            NumberFormatValues = "decimalFullWidth2"
	NumberFormatValuesAiueoFullWidth               NumberFormatValues = "aiueoFullWidth"
	NumberFormatValuesIrohaFullWidth               NumberFormatValues = "irohaFullWidth"
	NumberFormatValuesDecimalZero                  NumberFormatValues = "decimalZero"
	NumberFormatValuesBullet                       NumberFormatValues = "bullet"
	NumberFormatValuesGanada                       NumberFormatValues = "ganada"
	NumberFormatValuesChosung                      NumberFormatValues = "chosung"
	NumberFormatValuesDecimalEnclosedFullstop      NumberFormatValues = "decimalEnclosedFullstop"
	NumberFormatValuesDecimalEnclosedParen         NumberFormatValues = "decimalEnclosedParen"
	NumberFormatValuesDecimalEnclosedCircleChinese NumberFormatValues = "decimalEnclosedCircleChinese"
	NumberFormatValuesIdeographEnclosedCircle      NumberFormatValues = "ideographEnclosedCircle"
	NumberFormatValuesIdeographTraditional         NumberFormatValues = "ideographTraditional"
	NumberFormatValuesIdeographZodiac              NumberFormatValues = "ideographZodiac"
	NumberFormatValuesIdeographZodiacTraditional   NumberFormatValues = "ideographZodiacTraditional"
	NumberFormatValuesTaiwaneseCounting            NumberFormatValues = "taiwaneseCounting"
	NumberFormatValuesIdeographLegalTraditional    NumberFormatValues = "ideographLegalTraditional"
	NumberFormatValuesTaiwaneseCountingThousand    NumberFormatValues = "taiwaneseCountingThousand"
	NumberFormatValuesTaiwaneseDigital             NumberFormatValues = "taiwaneseDigital"
	NumberFormatValuesChineseCounting              NumberFormatValues = "chineseCounting"
	NumberFormatValuesChineseLegalSimplified       NumberFormatValues = "chineseLegalSimplified"
	NumberFormatValuesChineseCountingThousand      NumberFormatValues = "chineseCountingThousand"
	NumberFormatValuesKoreanDigital                NumberFormatValues = "koreanDigital"
	NumberFormatValuesKoreanCounting               NumberFormatValues = "koreanCounting"
	NumberFormatValuesKoreanLegal                  NumberFormatValues = "koreanLegal"
	NumberFormatValuesKoreanDigital2               NumberFormatValues = "koreanDigital2"
	NumberFormatValuesVietnameseCounting           NumberFormatValues = "vietnameseCounting"
	NumberFormatValuesRussianLower                 NumberFormatValues = "russianLower"
	NumberFormatValuesRussianUpper                 NumberFormatValues = "russianUpper"
	NumberFormatValuesNone                         NumberFormatValues = "none"
	NumberFormatValuesNumberInDash                 NumberFormatValues = "numberInDash"
	NumberFormatValuesHebrew1                      NumberFormatValues = "hebrew1"
	NumberFormatValuesHebrew2                      NumberFormatValues = "hebrew2"
	NumberFormatValuesArabicAlpha                  NumberFormatValues = "arabicAlpha"
	NumberFormatValuesArabicAbjad                  NumberFormatValues = "arabicAbjad"
	NumberFormatValuesHindiVowels                  NumberFormatValues = "hindiVowels"
	NumberFormatValuesHindiConsonants              NumberFormatValues = "hindiConsonants"
	NumberFormatValuesHindiNumbers                 NumberFormatValues = "hindiNumbers"
	NumberFormatValuesHindiCounting                NumberFormatValues = "hindiCounting"
	NumberFormatValuesThaiLetters                  NumberFormatValues = "thaiLetters"
	NumberFormatValuesThaiNumbers                  NumberFormatValues = "thaiNumbers"
	NumberFormatValuesThaiCounting                 NumberFormatValues = "thaiCounting"
	NumberFormatValuesBahtText                     NumberFormatValues = "bahtText"
	NumberFormatValuesDollarText                   NumberFormatValues = "dollarText"
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

// CompoundLineValues
type CompoundLineValues string

const (
	CompoundLineValuesSng       CompoundLineValues = "sng"
	CompoundLineValuesDbl       CompoundLineValues = "dbl"
	CompoundLineValuesThickThin CompoundLineValues = "thickThin"
	CompoundLineValuesThinThick CompoundLineValues = "thinThick"
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

// QualifierValues
type QualifierValues string

const (
	QualifierValuesDoubleQuote QualifierValues = "doubleQuote"
	QualifierValuesSingleQuote QualifierValues = "singleQuote"
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

// MailMergeDataValues
type MailMergeDataValues string

const (
	MailMergeDataValuesTextFile    MailMergeDataValues = "textFile"
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

// ChapterSeparatorValues
type ChapterSeparatorValues string

const (
	ChapterSeparatorValuesHyphen ChapterSeparatorValues = "hyphen"
	ChapterSeparatorValuesPeriod ChapterSeparatorValues = "period"
	ChapterSeparatorValuesColon  ChapterSeparatorValues = "colon"
	ChapterSeparatorValuesEmDash ChapterSeparatorValues = "emDash"
	ChapterSeparatorValuesEnDash ChapterSeparatorValues = "enDash"
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

// TableStyleOverrideValues
type TableStyleOverrideValues string

const (
	TableStyleOverrideValuesWholeTable TableStyleOverrideValues = "wholeTable"
	TableStyleOverrideValuesFirstRow   TableStyleOverrideValues = "firstRow"
	TableStyleOverrideValuesLastRow    TableStyleOverrideValues = "lastRow"
	TableStyleOverrideValuesFirstCol   TableStyleOverrideValues = "firstCol"
	TableStyleOverrideValuesLastCol    TableStyleOverrideValues = "lastCol"
	TableStyleOverrideValuesBand1Vert  TableStyleOverrideValues = "band1Vert"
	TableStyleOverrideValuesBand2Vert  TableStyleOverrideValues = "band2Vert"
	TableStyleOverrideValuesBand1Horz  TableStyleOverrideValues = "band1Horz"
	TableStyleOverrideValuesBand2Horz  TableStyleOverrideValues = "band2Horz"
	TableStyleOverrideValuesNeCell     TableStyleOverrideValues = "neCell"
	TableStyleOverrideValuesNwCell     TableStyleOverrideValues = "nwCell"
	TableStyleOverrideValuesSeCell     TableStyleOverrideValues = "seCell"
	TableStyleOverrideValuesSwCell     TableStyleOverrideValues = "swCell"
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

// AllocationMethodValues
type AllocationMethodValues string

const (
	AllocationMethodValuesEqualAllocation    AllocationMethodValues = "equalAllocation"
	AllocationMethodValuesEqualIncrement     AllocationMethodValues = "equalIncrement"
	AllocationMethodValuesWeightedAllocation AllocationMethodValues = "weightedAllocation"
	AllocationMethodValuesWeightedIncrement  AllocationMethodValues = "weightedIncrement"
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

// CompatSettingNameValues
type CompatSettingNameValues string

const (
	CompatSettingNameValuesCompatibilityMode                          CompatSettingNameValues = "compatibilityMode"
	CompatSettingNameValuesOverrideTableStyleFontSizeAndJustification CompatSettingNameValues = "overrideTableStyleFontSizeAndJustification"
	CompatSettingNameValuesEnableOpenTypeFeatures                     CompatSettingNameValues = "enableOpenTypeFeatures"
	CompatSettingNameValuesDoNotFlipMirrorIndents                     CompatSettingNameValues = "doNotFlipMirrorIndents"
	CompatSettingNameValuesDifferentiateMultirowTableHeaders          CompatSettingNameValues = "differentiateMultirowTableHeaders"
	CompatSettingNameValuesUseWord2013TrackBottomHyphenation          CompatSettingNameValues = "useWord2013TrackBottomHyphenation"
	CompatSettingNameValuesAllowHyphenationAtTrackBottom              CompatSettingNameValues = "allowHyphenationAtTrackBottom"
	CompatSettingNameValuesAllowTextAfterFloatingTableBreak           CompatSettingNameValues = "allowTextAfterFloatingTableBreak"
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

// TimeNodeMasterRelationValues
type TimeNodeMasterRelationValues string

const (
	TimeNodeMasterRelationValuesSameClick TimeNodeMasterRelationValues = "sameClick"
	TimeNodeMasterRelationValuesNextClick TimeNodeMasterRelationValues = "nextClick"
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

// EndnotePositionValues
type EndnotePositionValues string

const (
	EndnotePositionValuesSectEnd EndnotePositionValues = "sectEnd"
	EndnotePositionValuesDocEnd  EndnotePositionValues = "docEnd"
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

// MailMergeSourceValues
type MailMergeSourceValues string

const (
	MailMergeSourceValuesDatabase    MailMergeSourceValues = "database"
	MailMergeSourceValuesAddressBook MailMergeSourceValues = "addressBook"
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
	StylePaneSortMethodsValuesBasedOn                        StylePaneSortMethodsValues = "basedOn"
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

// PathFillModeValues
type PathFillModeValues string

const (
	PathFillModeValuesNone        PathFillModeValues = "none"
	PathFillModeValuesNorm        PathFillModeValues = "norm"
	PathFillModeValuesLighten     PathFillModeValues = "lighten"
	PathFillModeValuesLightenLess PathFillModeValues = "lightenLess"
	PathFillModeValuesDarken      PathFillModeValues = "darken"
	PathFillModeValuesDarkenLess  PathFillModeValues = "darkenLess"
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

// ParagraphBuildValues
type ParagraphBuildValues string

const (
	ParagraphBuildValuesAllAtOnce ParagraphBuildValues = "allAtOnce"
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
	PresetPatternValuesLtHorz     PresetPatternValues = "ltHorz"
	PresetPatternValuesLtVert     PresetPatternValues = "ltVert"
	PresetPatternValuesDkHorz     PresetPatternValues = "dkHorz"
	PresetPatternValuesDkVert     PresetPatternValues = "dkVert"
	PresetPatternValuesNarHorz    PresetPatternValues = "narHorz"
	PresetPatternValuesNarVert    PresetPatternValues = "narVert"
	PresetPatternValuesDashHorz   PresetPatternValues = "dashHorz"
	PresetPatternValuesDashVert   PresetPatternValues = "dashVert"
	PresetPatternValuesCross      PresetPatternValues = "cross"
	PresetPatternValuesDnDiag     PresetPatternValues = "dnDiag"
	PresetPatternValuesUpDiag     PresetPatternValues = "upDiag"
	PresetPatternValuesLtDnDiag   PresetPatternValues = "ltDnDiag"
	PresetPatternValuesLtUpDiag   PresetPatternValues = "ltUpDiag"
	PresetPatternValuesDkDnDiag   PresetPatternValues = "dkDnDiag"
	PresetPatternValuesDkUpDiag   PresetPatternValues = "dkUpDiag"
	PresetPatternValuesWdDnDiag   PresetPatternValues = "wdDnDiag"
	PresetPatternValuesWdUpDiag   PresetPatternValues = "wdUpDiag"
	PresetPatternValuesDashDnDiag PresetPatternValues = "dashDnDiag"
	PresetPatternValuesDashUpDiag PresetPatternValues = "dashUpDiag"
	PresetPatternValuesDiagCross  PresetPatternValues = "diagCross"
	PresetPatternValuesSmCheck    PresetPatternValues = "smCheck"
	PresetPatternValuesLgCheck    PresetPatternValues = "lgCheck"
	PresetPatternValuesSmGrid     PresetPatternValues = "smGrid"
	PresetPatternValuesLgGrid     PresetPatternValues = "lgGrid"
	PresetPatternValuesDotGrid    PresetPatternValues = "dotGrid"
	PresetPatternValuesSmConfetti PresetPatternValues = "smConfetti"
	PresetPatternValuesLgConfetti PresetPatternValues = "lgConfetti"
	PresetPatternValuesHorzBrick  PresetPatternValues = "horzBrick"
	PresetPatternValuesDiagBrick  PresetPatternValues = "diagBrick"
	PresetPatternValuesSolidDmnd  PresetPatternValues = "solidDmnd"
	PresetPatternValuesOpenDmnd   PresetPatternValues = "openDmnd"
	PresetPatternValuesDotDmnd    PresetPatternValues = "dotDmnd"
	PresetPatternValuesPlaid      PresetPatternValues = "plaid"
	PresetPatternValuesSphere     PresetPatternValues = "sphere"
	PresetPatternValuesWeave      PresetPatternValues = "weave"
	PresetPatternValuesDivot      PresetPatternValues = "divot"
	PresetPatternValuesShingle    PresetPatternValues = "shingle"
	PresetPatternValuesWave       PresetPatternValues = "wave"
	PresetPatternValuesTrellis    PresetPatternValues = "trellis"
	PresetPatternValuesZigZag     PresetPatternValues = "zigZag"
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

// PageOrderValues
type PageOrderValues string

const (
	PageOrderValuesDownThenOver PageOrderValues = "downThenOver"
	PageOrderValuesOverThenDown PageOrderValues = "overThenDown"
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

// MultiLevelValues
type MultiLevelValues string

const (
	MultiLevelValuesSingleLevel      MultiLevelValues = "singleLevel"
	MultiLevelValuesMultilevel       MultiLevelValues = "multilevel"
	MultiLevelValuesHybridMultilevel MultiLevelValues = "hybridMultilevel"
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

// PresetMaterialTypeValues
type PresetMaterialTypeValues string

const (
	PresetMaterialTypeValuesLegacyMatte       PresetMaterialTypeValues = "legacyMatte"
	PresetMaterialTypeValuesLegacyPlastic     PresetMaterialTypeValues = "legacyPlastic"
	PresetMaterialTypeValuesLegacyMetal       PresetMaterialTypeValues = "legacyMetal"
	PresetMaterialTypeValuesLegacyWireframe   PresetMaterialTypeValues = "legacyWireframe"
	PresetMaterialTypeValuesMatte             PresetMaterialTypeValues = "matte"
	PresetMaterialTypeValuesPlastic           PresetMaterialTypeValues = "plastic"
	PresetMaterialTypeValuesMetal             PresetMaterialTypeValues = "metal"
	PresetMaterialTypeValuesWarmMatte         PresetMaterialTypeValues = "warmMatte"
	PresetMaterialTypeValuesTranslucentPowder PresetMaterialTypeValues = "translucentPowder"
	PresetMaterialTypeValuesPowder            PresetMaterialTypeValues = "powder"
	PresetMaterialTypeValuesDkEdge            PresetMaterialTypeValues = "dkEdge"
	PresetMaterialTypeValuesSoftEdge          PresetMaterialTypeValues = "softEdge"
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

// TextStrikeValues
type TextStrikeValues string

const (
	TextStrikeValuesNoStrike  TextStrikeValues = "noStrike"
	TextStrikeValuesSngStrike TextStrikeValues = "sngStrike"
	TextStrikeValuesDblStrike TextStrikeValues = "dblStrike"
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

// PlaceholderValues
type PlaceholderValues string

const (
	PlaceholderValuesTitle    PlaceholderValues = "title"
	PlaceholderValuesBody     PlaceholderValues = "body"
	PlaceholderValuesCtrTitle PlaceholderValues = "ctrTitle"
	PlaceholderValuesSubTitle PlaceholderValues = "subTitle"
	PlaceholderValuesDt       PlaceholderValues = "dt"
	PlaceholderValuesSldNum   PlaceholderValues = "sldNum"
	PlaceholderValuesFtr      PlaceholderValues = "ftr"
	PlaceholderValuesHdr      PlaceholderValues = "hdr"
	PlaceholderValuesObj      PlaceholderValues = "obj"
	PlaceholderValuesChart    PlaceholderValues = "chart"
	PlaceholderValuesTbl      PlaceholderValues = "tbl"
	PlaceholderValuesClipArt  PlaceholderValues = "clipArt"
	PlaceholderValuesDgm      PlaceholderValues = "dgm"
	PlaceholderValuesMedia    PlaceholderValues = "media"
	PlaceholderValuesSldImg   PlaceholderValues = "sldImg"
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

// ConditionalFormattingOperatorValues
type ConditionalFormattingOperatorValues string

const (
	ConditionalFormattingOperatorValuesLessThan           ConditionalFormattingOperatorValues = "lessThan"
	ConditionalFormattingOperatorValuesLessThanOrEqual    ConditionalFormattingOperatorValues = "lessThanOrEqual"
	ConditionalFormattingOperatorValuesEqual              ConditionalFormattingOperatorValues = "equal"
	ConditionalFormattingOperatorValuesNotEqual           ConditionalFormattingOperatorValues = "notEqual"
	ConditionalFormattingOperatorValuesGreaterThanOrEqual ConditionalFormattingOperatorValues = "greaterThanOrEqual"
	ConditionalFormattingOperatorValuesGreaterThan        ConditionalFormattingOperatorValues = "greaterThan"
	ConditionalFormattingOperatorValuesBetween            ConditionalFormattingOperatorValues = "between"
	ConditionalFormattingOperatorValuesNotBetween         ConditionalFormattingOperatorValues = "notBetween"
	ConditionalFormattingOperatorValuesContainsText       ConditionalFormattingOperatorValues = "containsText"
	ConditionalFormattingOperatorValuesNotContains        ConditionalFormattingOperatorValues = "notContains"
	ConditionalFormattingOperatorValuesBeginsWith         ConditionalFormattingOperatorValues = "beginsWith"
	ConditionalFormattingOperatorValuesEndsWith           ConditionalFormattingOperatorValues = "endsWith"
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

// UpdateLinksBehaviorValues
type UpdateLinksBehaviorValues string

const (
	UpdateLinksBehaviorValuesUserSet UpdateLinksBehaviorValues = "userSet"
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

// ConditionalFormattingValueObjectTypeValues
type ConditionalFormattingValueObjectTypeValues string

const (
	ConditionalFormattingValueObjectTypeValuesNum        ConditionalFormattingValueObjectTypeValues = "num"
	ConditionalFormattingValueObjectTypeValuesPercent    ConditionalFormattingValueObjectTypeValues = "percent"
	ConditionalFormattingValueObjectTypeValuesMax        ConditionalFormattingValueObjectTypeValues = "max"
	ConditionalFormattingValueObjectTypeValuesMin        ConditionalFormattingValueObjectTypeValues = "min"
	ConditionalFormattingValueObjectTypeValuesFormula    ConditionalFormattingValueObjectTypeValues = "formula"
	ConditionalFormattingValueObjectTypeValuesPercentile ConditionalFormattingValueObjectTypeValues = "percentile"
	ConditionalFormattingValueObjectTypeValuesAutoMin    ConditionalFormattingValueObjectTypeValues = "autoMin"
	ConditionalFormattingValueObjectTypeValuesAutoMax    ConditionalFormattingValueObjectTypeValues = "autoMax"
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

// AbsolutePositionTabLeaderCharValues
type AbsolutePositionTabLeaderCharValues string

const (
	AbsolutePositionTabLeaderCharValuesNone       AbsolutePositionTabLeaderCharValues = "none"
	AbsolutePositionTabLeaderCharValuesDot        AbsolutePositionTabLeaderCharValues = "dot"
	AbsolutePositionTabLeaderCharValuesHyphen     AbsolutePositionTabLeaderCharValues = "hyphen"
	AbsolutePositionTabLeaderCharValuesUnderscore AbsolutePositionTabLeaderCharValues = "underscore"
	AbsolutePositionTabLeaderCharValuesMiddleDot  AbsolutePositionTabLeaderCharValues = "middleDot"
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

// PhotoAlbumFrameShapeValues
type PhotoAlbumFrameShapeValues string

const (
	PhotoAlbumFrameShapeValuesFrameStyle1 PhotoAlbumFrameShapeValues = "frameStyle1"
	PhotoAlbumFrameShapeValuesFrameStyle2 PhotoAlbumFrameShapeValues = "frameStyle2"
	PhotoAlbumFrameShapeValuesFrameStyle3 PhotoAlbumFrameShapeValues = "frameStyle3"
	PhotoAlbumFrameShapeValuesFrameStyle4 PhotoAlbumFrameShapeValues = "frameStyle4"
	PhotoAlbumFrameShapeValuesFrameStyle5 PhotoAlbumFrameShapeValues = "frameStyle5"
	PhotoAlbumFrameShapeValuesFrameStyle6 PhotoAlbumFrameShapeValues = "frameStyle6"
	PhotoAlbumFrameShapeValuesFrameStyle7 PhotoAlbumFrameShapeValues = "frameStyle7"
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

// KnownContextNodeTypeValues
type KnownContextNodeTypeValues string

const (
	KnownContextNodeTypeValuesRoot             KnownContextNodeTypeValues = "root"
	KnownContextNodeTypeValuesUnclassifiedInk  KnownContextNodeTypeValues = "unclassifiedInk"
	KnownContextNodeTypeValuesWritingRegion    KnownContextNodeTypeValues = "writingRegion"
	KnownContextNodeTypeValuesAnalysisHint     KnownContextNodeTypeValues = "analysisHint"
	KnownContextNodeTypeValuesObject           KnownContextNodeTypeValues = "object"
	KnownContextNodeTypeValuesInkDrawing       KnownContextNodeTypeValues = "inkDrawing"
	KnownContextNodeTypeValuesImage            KnownContextNodeTypeValues = "image"
	KnownContextNodeTypeValuesParagraph        KnownContextNodeTypeValues = "paragraph"
	KnownContextNodeTypeValuesLine             KnownContextNodeTypeValues = "line"
	KnownContextNodeTypeValuesInkBullet        KnownContextNodeTypeValues = "inkBullet"
	KnownContextNodeTypeValuesInkWord          KnownContextNodeTypeValues = "inkWord"
	KnownContextNodeTypeValuesTextWord         KnownContextNodeTypeValues = "textWord"
	KnownContextNodeTypeValuesCustomRecognizer KnownContextNodeTypeValues = "customRecognizer"
	KnownContextNodeTypeValuesMathRegion       KnownContextNodeTypeValues = "mathRegion"
	KnownContextNodeTypeValuesMathEquation     KnownContextNodeTypeValues = "mathEquation"
	KnownContextNodeTypeValuesMathStruct       KnownContextNodeTypeValues = "mathStruct"
	KnownContextNodeTypeValuesMathSymbol       KnownContextNodeTypeValues = "mathSymbol"
	KnownContextNodeTypeValuesMathIdentifier   KnownContextNodeTypeValues = "mathIdentifier"
	KnownContextNodeTypeValuesMathOperator     KnownContextNodeTypeValues = "mathOperator"
	KnownContextNodeTypeValuesMathNumber       KnownContextNodeTypeValues = "mathNumber"
	KnownContextNodeTypeValuesNonInkDrawing    KnownContextNodeTypeValues = "nonInkDrawing"
	KnownContextNodeTypeValuesGroupNode        KnownContextNodeTypeValues = "groupNode"
	KnownContextNodeTypeValuesMixedDrawing     KnownContextNodeTypeValues = "mixedDrawing"
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

// ObjectTypeValues
type ObjectTypeValues string

const (
	ObjectTypeValuesButton   ObjectTypeValues = "Button"
	ObjectTypeValuesCheckBox ObjectTypeValues = "CheckBox"
	ObjectTypeValuesDrop     ObjectTypeValues = "Drop"
	ObjectTypeValuesGBox     ObjectTypeValues = "GBox"
	ObjectTypeValuesLabel    ObjectTypeValues = "Label"
	ObjectTypeValuesList     ObjectTypeValues = "List"
	ObjectTypeValuesRadio    ObjectTypeValues = "Radio"
	ObjectTypeValuesScroll   ObjectTypeValues = "Scroll"
	ObjectTypeValuesSpin     ObjectTypeValues = "Spin"
	ObjectTypeValuesEditBox  ObjectTypeValues = "EditBox"
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

// CellCommentsValues
type CellCommentsValues string

const (
	CellCommentsValuesNone        CellCommentsValues = "none"
	CellCommentsValuesAsDisplayed CellCommentsValues = "asDisplayed"
	CellCommentsValuesAtEnd       CellCommentsValues = "atEnd"
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

// ConditionalFormatValues
type ConditionalFormatValues string

const (
	ConditionalFormatValuesExpression        ConditionalFormatValues = "expression"
	ConditionalFormatValuesCellIs            ConditionalFormatValues = "cellIs"
	ConditionalFormatValuesColorScale        ConditionalFormatValues = "colorScale"
	ConditionalFormatValuesDataBar           ConditionalFormatValues = "dataBar"
	ConditionalFormatValuesIconSet           ConditionalFormatValues = "iconSet"
	ConditionalFormatValuesTop10             ConditionalFormatValues = "top10"
	ConditionalFormatValuesUniqueValues      ConditionalFormatValues = "uniqueValues"
	ConditionalFormatValuesDuplicateValues   ConditionalFormatValues = "duplicateValues"
	ConditionalFormatValuesContainsText      ConditionalFormatValues = "containsText"
	ConditionalFormatValuesNotContainsText   ConditionalFormatValues = "notContainsText"
	ConditionalFormatValuesBeginsWith        ConditionalFormatValues = "beginsWith"
	ConditionalFormatValuesEndsWith          ConditionalFormatValues = "endsWith"
	ConditionalFormatValuesContainsBlanks    ConditionalFormatValues = "containsBlanks"
	ConditionalFormatValuesNotContainsBlanks ConditionalFormatValues = "notContainsBlanks"
	ConditionalFormatValuesContainsErrors    ConditionalFormatValues = "containsErrors"
	ConditionalFormatValuesNotContainsErrors ConditionalFormatValues = "notContainsErrors"
	ConditionalFormatValuesTimePeriod        ConditionalFormatValues = "timePeriod"
	ConditionalFormatValuesAboveAverage      ConditionalFormatValues = "aboveAverage"
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

// TargetScreenSizeValues
type TargetScreenSizeValues string

const (
	TargetScreenSizeValuesTargetScreenSizeValues544x376   TargetScreenSizeValues = "544x376"
	TargetScreenSizeValuesTargetScreenSizeValues640x480   TargetScreenSizeValues = "640x480"
	TargetScreenSizeValuesTargetScreenSizeValues720x512   TargetScreenSizeValues = "720x512"
	TargetScreenSizeValuesTargetScreenSizeValues800x600   TargetScreenSizeValues = "800x600"
	TargetScreenSizeValuesTargetScreenSizeValues1024x768  TargetScreenSizeValues = "1024x768"
	TargetScreenSizeValuesTargetScreenSizeValues1152x882  TargetScreenSizeValues = "1152x882"
	TargetScreenSizeValuesTargetScreenSizeValues1152x900  TargetScreenSizeValues = "1152x900"
	TargetScreenSizeValuesTargetScreenSizeValues1280x1024 TargetScreenSizeValues = "1280x1024"
	TargetScreenSizeValuesTargetScreenSizeValues1600x1200 TargetScreenSizeValues = "1600x1200"
	TargetScreenSizeValuesTargetScreenSizeValues1800x1440 TargetScreenSizeValues = "1800x1440"
	TargetScreenSizeValuesTargetScreenSizeValues1920x1200 TargetScreenSizeValues = "1920x1200"
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

// LineNumberRestartValues
type LineNumberRestartValues string

const (
	LineNumberRestartValuesNewPage    LineNumberRestartValues = "newPage"
	LineNumberRestartValuesNewSection LineNumberRestartValues = "newSection"
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

// SlicerCacheCrossFilterValues
type SlicerCacheCrossFilterValues string

const (
	SlicerCacheCrossFilterValuesNone                   SlicerCacheCrossFilterValues = "none"
	SlicerCacheCrossFilterValuesShowItemsWithDataAtTop SlicerCacheCrossFilterValues = "showItemsWithDataAtTop"
	SlicerCacheCrossFilterValuesShowItemsWithNoData    SlicerCacheCrossFilterValues = "showItemsWithNoData"
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

// TableStyleValues
type TableStyleValues string

const (
	TableStyleValuesWholeTable             TableStyleValues = "wholeTable"
	TableStyleValuesHeaderRow              TableStyleValues = "headerRow"
	TableStyleValuesTotalRow               TableStyleValues = "totalRow"
	TableStyleValuesFirstColumn            TableStyleValues = "firstColumn"
	TableStyleValuesLastColumn             TableStyleValues = "lastColumn"
	TableStyleValuesFirstRowStripe         TableStyleValues = "firstRowStripe"
	TableStyleValuesSecondRowStripe        TableStyleValues = "secondRowStripe"
	TableStyleValuesFirstColumnStripe      TableStyleValues = "firstColumnStripe"
	TableStyleValuesSecondColumnStripe     TableStyleValues = "secondColumnStripe"
	TableStyleValuesFirstHeaderCell        TableStyleValues = "firstHeaderCell"
	TableStyleValuesLastHeaderCell         TableStyleValues = "lastHeaderCell"
	TableStyleValuesFirstTotalCell         TableStyleValues = "firstTotalCell"
	TableStyleValuesLastTotalCell          TableStyleValues = "lastTotalCell"
	TableStyleValuesFirstSubtotalColumn    TableStyleValues = "firstSubtotalColumn"
	TableStyleValuesSecondSubtotalColumn   TableStyleValues = "secondSubtotalColumn"
	TableStyleValuesThirdSubtotalColumn    TableStyleValues = "thirdSubtotalColumn"
	TableStyleValuesFirstSubtotalRow       TableStyleValues = "firstSubtotalRow"
	TableStyleValuesSecondSubtotalRow      TableStyleValues = "secondSubtotalRow"
	TableStyleValuesThirdSubtotalRow       TableStyleValues = "thirdSubtotalRow"
	TableStyleValuesBlankRow               TableStyleValues = "blankRow"
	TableStyleValuesFirstColumnSubheading  TableStyleValues = "firstColumnSubheading"
	TableStyleValuesSecondColumnSubheading TableStyleValues = "secondColumnSubheading"
	TableStyleValuesThirdColumnSubheading  TableStyleValues = "thirdColumnSubheading"
	TableStyleValuesFirstRowSubheading     TableStyleValues = "firstRowSubheading"
	TableStyleValuesSecondRowSubheading    TableStyleValues = "secondRowSubheading"
	TableStyleValuesThirdRowSubheading     TableStyleValues = "thirdRowSubheading"
	TableStyleValuesPageFieldLabels        TableStyleValues = "pageFieldLabels"
	TableStyleValuesPageFieldValues        TableStyleValues = "pageFieldValues"
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

// EmphasisMarkValues
type EmphasisMarkValues string

const (
	EmphasisMarkValuesNone     EmphasisMarkValues = "none"
	EmphasisMarkValuesDot      EmphasisMarkValues = "dot"
	EmphasisMarkValuesComma    EmphasisMarkValues = "comma"
	EmphasisMarkValuesCircle   EmphasisMarkValues = "circle"
	EmphasisMarkValuesUnderDot EmphasisMarkValues = "underDot"
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

// TabStopLeaderCharValues
type TabStopLeaderCharValues string

const (
	TabStopLeaderCharValuesNone       TabStopLeaderCharValues = "none"
	TabStopLeaderCharValuesDot        TabStopLeaderCharValues = "dot"
	TabStopLeaderCharValuesHyphen     TabStopLeaderCharValues = "hyphen"
	TabStopLeaderCharValuesUnderscore TabStopLeaderCharValues = "underscore"
	TabStopLeaderCharValuesHeavy      TabStopLeaderCharValues = "heavy"
	TabStopLeaderCharValuesMiddleDot  TabStopLeaderCharValues = "middleDot"
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

// QuestionFormat
type QuestionFormat string

const (
	QuestionFormatGeneralDate   QuestionFormat = "generalDate"
	QuestionFormatLongDate      QuestionFormat = "longDate"
	QuestionFormatShortDate     QuestionFormat = "shortDate"
	QuestionFormatLongTime      QuestionFormat = "longTime"
	QuestionFormatShortTime     QuestionFormat = "shortTime"
	QuestionFormatGeneralNumber QuestionFormat = "generalNumber"
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

// FootnoteEndnoteValues
type FootnoteEndnoteValues string

const (
	FootnoteEndnoteValuesNormal                FootnoteEndnoteValues = "normal"
	FootnoteEndnoteValuesSeparator             FootnoteEndnoteValues = "separator"
	FootnoteEndnoteValuesContinuationSeparator FootnoteEndnoteValues = "continuationSeparator"
	FootnoteEndnoteValuesContinuationNotice    FootnoteEndnoteValues = "continuationNotice"
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
	ThemeColorValuesFollowedHyperlink ThemeColorValues = "followedHyperlink"
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
	SchemeColorValuesFolHlink SchemeColorValues = "folHlink"
	SchemeColorValuesPhClr    SchemeColorValues = "phClr"
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

// SlideSizeValues
type SlideSizeValues string

const (
	SlideSizeValuesScreen4x3           SlideSizeValues = "screen4x3"
	SlideSizeValuesLetter              SlideSizeValues = "letter"
	SlideSizeValuesA4                  SlideSizeValues = "A4"
	SlideSizeValuesSlideSizeValues35mm SlideSizeValues = "35mm"
	SlideSizeValuesOverhead            SlideSizeValues = "overhead"
	SlideSizeValuesBanner              SlideSizeValues = "banner"
	SlideSizeValuesCustom              SlideSizeValues = "custom"
	SlideSizeValuesLedger              SlideSizeValues = "ledger"
	SlideSizeValuesA3                  SlideSizeValues = "A3"
	SlideSizeValuesB4ISO               SlideSizeValues = "B4ISO"
	SlideSizeValuesB5ISO               SlideSizeValues = "B5ISO"
	SlideSizeValuesB4JIS               SlideSizeValues = "B4JIS"
	SlideSizeValuesB5JIS               SlideSizeValues = "B5JIS"
	SlideSizeValuesHagakiCard          SlideSizeValues = "hagakiCard"
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

// VisibilityValues
type VisibilityValues string

const (
	VisibilityValuesVisible    VisibilityValues = "visible"
	VisibilityValuesHidden     VisibilityValues = "hidden"
	VisibilityValuesVeryHidden VisibilityValues = "veryHidden"
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

// RubyAlignValues
type RubyAlignValues string

const (
	RubyAlignValuesCenter           RubyAlignValues = "center"
	RubyAlignValuesDistributeLetter RubyAlignValues = "distributeLetter"
	RubyAlignValuesDistributeSpace  RubyAlignValues = "distributeSpace"
	RubyAlignValuesLeft             RubyAlignValues = "left"
	RubyAlignValuesRight            RubyAlignValues = "right"
	RubyAlignValuesRightVertical    RubyAlignValues = "rightVertical"
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
