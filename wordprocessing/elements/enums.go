// Package elements provides Word document element types.
package elements

// JustificationValue represents paragraph alignment values.
type JustificationValue string

const (
	// JustificationLeft aligns text to the left margin.
	JustificationLeft JustificationValue = "left"
	// JustificationCenter centers text between margins.
	JustificationCenter JustificationValue = "center"
	// JustificationRight aligns text to the right margin.
	JustificationRight JustificationValue = "right"
	// JustificationBoth justifies text between both margins.
	JustificationBoth JustificationValue = "both"
	// JustificationDistribute distributes text evenly.
	JustificationDistribute JustificationValue = "distribute"
)

// UnderlineValue represents underline style values.
type UnderlineValue string

const (
	// UnderlineNone removes underline.
	UnderlineNone UnderlineValue = "none"
	// UnderlineSingle applies single underline.
	UnderlineSingle UnderlineValue = "single"
	// UnderlineDouble applies double underline.
	UnderlineDouble UnderlineValue = "double"
	// UnderlineDash applies dashed underline.
	UnderlineDash UnderlineValue = "dash"
	// UnderlineDotDash applies dot-dash underline.
	UnderlineDotDash UnderlineValue = "dotDash"
	// UnderlineDotDotDash applies dot-dot-dash underline.
	UnderlineDotDotDash UnderlineValue = "dotDotDash"
	// UnderlineDotted applies dotted underline.
	UnderlineDotted UnderlineValue = "dotted"
	// UnderlineThick applies thick underline.
	UnderlineThick UnderlineValue = "thick"
	// UnderlineWave applies wavy underline.
	UnderlineWave UnderlineValue = "wave"
	// UnderlineWords applies underline to words only.
	UnderlineWords UnderlineValue = "words"
	// UnderlineWavyDouble applies double wavy underline.
	UnderlineWavyDouble UnderlineValue = "wavyDouble"
	// UnderlineWavyHeavy applies heavy wavy underline.
	UnderlineWavyHeavy UnderlineValue = "wavyHeavy"
	// UnderlineDashLong applies long dash underline.
	UnderlineDashLong UnderlineValue = "dashLong"
	// UnderlineDashDotHeavy applies heavy dash-dot underline.
	UnderlineDashDotHeavy UnderlineValue = "dashDotHeavy"
	// UnderlineDashDotDotHeavy applies heavy dash-dot-dot underline.
	UnderlineDashDotDotHeavy UnderlineValue = "dashDotDotHeavy"
	// UnderlineDashLongHeavy applies heavy long dash underline.
	UnderlineDashLongHeavy UnderlineValue = "dashLongHeavy"
	// UnderlineDottedHeavy applies heavy dotted underline.
	UnderlineDottedHeavy UnderlineValue = "dottedHeavy"
)

// HighlightColor represents text highlight color values.
type HighlightColor string

const (
	// HighlightNone removes highlight.
	HighlightNone HighlightColor = "none"
	// HighlightBlack applies black highlight.
	HighlightBlack HighlightColor = "black"
	// HighlightBlue applies blue highlight.
	HighlightBlue HighlightColor = "blue"
	// HighlightCyan applies cyan highlight.
	HighlightCyan HighlightColor = "cyan"
	// HighlightDarkBlue applies dark blue highlight.
	HighlightDarkBlue HighlightColor = "darkBlue"
	// HighlightDarkCyan applies dark cyan highlight.
	HighlightDarkCyan HighlightColor = "darkCyan"
	// HighlightDarkGray applies dark gray highlight.
	HighlightDarkGray HighlightColor = "darkGray"
	// HighlightDarkGreen applies dark green highlight.
	HighlightDarkGreen HighlightColor = "darkGreen"
	// HighlightDarkMagenta applies dark magenta highlight.
	HighlightDarkMagenta HighlightColor = "darkMagenta"
	// HighlightDarkRed applies dark red highlight.
	HighlightDarkRed HighlightColor = "darkRed"
	// HighlightDarkYellow applies dark yellow highlight.
	HighlightDarkYellow HighlightColor = "darkYellow"
	// HighlightGreen applies green highlight.
	HighlightGreen HighlightColor = "green"
	// HighlightLightGray applies light gray highlight.
	HighlightLightGray HighlightColor = "lightGray"
	// HighlightMagenta applies magenta highlight.
	HighlightMagenta HighlightColor = "magenta"
	// HighlightRed applies red highlight.
	HighlightRed HighlightColor = "red"
	// HighlightWhite applies white highlight.
	HighlightWhite HighlightColor = "white"
	// HighlightYellow applies yellow highlight.
	HighlightYellow HighlightColor = "yellow"
)

// BreakType represents break type values.
type BreakType string

const (
	// BreakLine represents a line break (default).
	BreakLine BreakType = ""
	// BreakPage represents a page break.
	BreakPage BreakType = "page"
	// BreakColumn represents a column break.
	BreakColumn BreakType = "column"
	// BreakTextWrapping represents a text wrapping break.
	BreakTextWrapping BreakType = "textWrapping"
)

// LineSpacingRule represents line spacing rule values.
type LineSpacingRule string

const (
	// LineSpacingAuto uses automatic line spacing.
	LineSpacingAuto LineSpacingRule = "auto"
	// LineSpacingExact uses exact line spacing.
	LineSpacingExact LineSpacingRule = "exact"
	// LineSpacingAtLeast uses at least the specified spacing.
	LineSpacingAtLeast LineSpacingRule = "atLeast"
)

// TabAlignment represents tab stop alignment values.
type TabAlignment string

const (
	// TabAlignLeft aligns tab stop to the left.
	TabAlignLeft TabAlignment = "left"
	// TabAlignCenter centers tab stop.
	TabAlignCenter TabAlignment = "center"
	// TabAlignRight aligns tab stop to the right.
	TabAlignRight TabAlignment = "right"
	// TabAlignDecimal aligns on decimal point.
	TabAlignDecimal TabAlignment = "decimal"
	// TabAlignBar creates a vertical bar.
	TabAlignBar TabAlignment = "bar"
	// TabAlignClear clears a tab stop.
	TabAlignClear TabAlignment = "clear"
)

// TabLeader represents tab leader character values.
type TabLeader string

const (
	// TabLeaderNone uses no leader character.
	TabLeaderNone TabLeader = "none"
	// TabLeaderDot uses dot leader.
	TabLeaderDot TabLeader = "dot"
	// TabLeaderHyphen uses hyphen leader.
	TabLeaderHyphen TabLeader = "hyphen"
	// TabLeaderUnderscore uses underscore leader.
	TabLeaderUnderscore TabLeader = "underscore"
	// TabLeaderHeavy uses heavy underscore leader.
	TabLeaderHeavy TabLeader = "heavy"
	// TabLeaderMiddleDot uses middle dot leader.
	TabLeaderMiddleDot TabLeader = "middleDot"
)

// BorderStyle represents border style values.
type BorderStyle string

const (
	// BorderNil removes the border.
	BorderNil BorderStyle = "nil"
	// BorderNone indicates no border.
	BorderNone BorderStyle = "none"
	// BorderSingle applies single line border.
	BorderSingle BorderStyle = "single"
	// BorderDouble applies double line border.
	BorderDouble BorderStyle = "double"
	// BorderDotted applies dotted border.
	BorderDotted BorderStyle = "dotted"
	// BorderDashed applies dashed border.
	BorderDashed BorderStyle = "dashed"
	// BorderDotDash applies dot-dash border.
	BorderDotDash BorderStyle = "dotDash"
	// BorderDotDotDash applies dot-dot-dash border.
	BorderDotDotDash BorderStyle = "dotDotDash"
	// BorderTriple applies triple line border.
	BorderTriple BorderStyle = "triple"
	// BorderThick applies thick border.
	BorderThick BorderStyle = "thick"
	// BorderThickThinSmallGap applies thick-thin border with small gap.
	BorderThickThinSmallGap BorderStyle = "thickThinSmallGap"
	// BorderThinThickSmallGap applies thin-thick border with small gap.
	BorderThinThickSmallGap BorderStyle = "thinThickSmallGap"
	// BorderThinThickThinSmallGap applies thin-thick-thin border with small gap.
	BorderThinThickThinSmallGap BorderStyle = "thinThickThinSmallGap"
	// BorderThickThinMediumGap applies thick-thin border with medium gap.
	BorderThickThinMediumGap BorderStyle = "thickThinMediumGap"
	// BorderThinThickMediumGap applies thin-thick border with medium gap.
	BorderThinThickMediumGap BorderStyle = "thinThickMediumGap"
	// BorderThinThickThinMediumGap applies thin-thick-thin border with medium gap.
	BorderThinThickThinMediumGap BorderStyle = "thinThickThinMediumGap"
	// BorderThickThinLargeGap applies thick-thin border with large gap.
	BorderThickThinLargeGap BorderStyle = "thickThinLargeGap"
	// BorderThinThickLargeGap applies thin-thick border with large gap.
	BorderThinThickLargeGap BorderStyle = "thinThickLargeGap"
	// BorderThinThickThinLargeGap applies thin-thick-thin border with large gap.
	BorderThinThickThinLargeGap BorderStyle = "thinThickThinLargeGap"
	// BorderWave applies wave border.
	BorderWave BorderStyle = "wave"
	// BorderDoubleWave applies double wave border.
	BorderDoubleWave BorderStyle = "doubleWave"
	// BorderDashSmallGap applies dashed border with small gap.
	BorderDashSmallGap BorderStyle = "dashSmallGap"
	// BorderDashDotStroked applies dash-dot-stroked border.
	BorderDashDotStroked BorderStyle = "dashDotStroked"
	// BorderThreeDEmboss applies 3D emboss border.
	BorderThreeDEmboss BorderStyle = "threeDEmboss"
	// BorderThreeDEngrave applies 3D engrave border.
	BorderThreeDEngrave BorderStyle = "threeDEngrave"
	// BorderOutset applies outset border.
	BorderOutset BorderStyle = "outset"
	// BorderInset applies inset border.
	BorderInset BorderStyle = "inset"
)

// VerticalAlignValue represents vertical text alignment values.
type VerticalAlignValue string

const (
	// VerticalAlignBaseline places text at baseline.
	VerticalAlignBaseline VerticalAlignValue = "baseline"
	// VerticalAlignSubscript places text as subscript.
	VerticalAlignSubscript VerticalAlignValue = "subscript"
	// VerticalAlignSuperscript places text as superscript.
	VerticalAlignSuperscript VerticalAlignValue = "superscript"
)

// ShadingPattern represents shading pattern values.
type ShadingPattern string

const (
	// ShadingNil removes shading.
	ShadingNil ShadingPattern = "nil"
	// ShadingClear uses no shading pattern (clear).
	ShadingClear ShadingPattern = "clear"
	// ShadingSolid uses solid fill.
	ShadingSolid ShadingPattern = "solid"
	// ShadingHorzStripe uses horizontal stripe pattern.
	ShadingHorzStripe ShadingPattern = "horzStripe"
	// ShadingVertStripe uses vertical stripe pattern.
	ShadingVertStripe ShadingPattern = "vertStripe"
	// ShadingReverseDiagStripe uses reverse diagonal stripe pattern.
	ShadingReverseDiagStripe ShadingPattern = "reverseDiagStripe"
	// ShadingDiagStripe uses diagonal stripe pattern.
	ShadingDiagStripe ShadingPattern = "diagStripe"
	// ShadingHorzCross uses horizontal cross pattern.
	ShadingHorzCross ShadingPattern = "horzCross"
	// ShadingDiagCross uses diagonal cross pattern.
	ShadingDiagCross ShadingPattern = "diagCross"
	// ShadingThinHorzStripe uses thin horizontal stripe pattern.
	ShadingThinHorzStripe ShadingPattern = "thinHorzStripe"
	// ShadingThinVertStripe uses thin vertical stripe pattern.
	ShadingThinVertStripe ShadingPattern = "thinVertStripe"
	// ShadingThinReverseDiagStripe uses thin reverse diagonal stripe pattern.
	ShadingThinReverseDiagStripe ShadingPattern = "thinReverseDiagStripe"
	// ShadingThinDiagStripe uses thin diagonal stripe pattern.
	ShadingThinDiagStripe ShadingPattern = "thinDiagStripe"
	// ShadingThinHorzCross uses thin horizontal cross pattern.
	ShadingThinHorzCross ShadingPattern = "thinHorzCross"
	// ShadingThinDiagCross uses thin diagonal cross pattern.
	ShadingThinDiagCross ShadingPattern = "thinDiagCross"
	// ShadingPct5 uses 5% pattern.
	ShadingPct5 ShadingPattern = "pct5"
	// ShadingPct10 uses 10% pattern.
	ShadingPct10 ShadingPattern = "pct10"
	// ShadingPct12 uses 12.5% pattern.
	ShadingPct12 ShadingPattern = "pct12"
	// ShadingPct15 uses 15% pattern.
	ShadingPct15 ShadingPattern = "pct15"
	// ShadingPct20 uses 20% pattern.
	ShadingPct20 ShadingPattern = "pct20"
	// ShadingPct25 uses 25% pattern.
	ShadingPct25 ShadingPattern = "pct25"
	// ShadingPct30 uses 30% pattern.
	ShadingPct30 ShadingPattern = "pct30"
	// ShadingPct35 uses 35% pattern.
	ShadingPct35 ShadingPattern = "pct35"
	// ShadingPct37 uses 37.5% pattern.
	ShadingPct37 ShadingPattern = "pct37"
	// ShadingPct40 uses 40% pattern.
	ShadingPct40 ShadingPattern = "pct40"
	// ShadingPct45 uses 45% pattern.
	ShadingPct45 ShadingPattern = "pct45"
	// ShadingPct50 uses 50% pattern.
	ShadingPct50 ShadingPattern = "pct50"
	// ShadingPct55 uses 55% pattern.
	ShadingPct55 ShadingPattern = "pct55"
	// ShadingPct60 uses 60% pattern.
	ShadingPct60 ShadingPattern = "pct60"
	// ShadingPct62 uses 62.5% pattern.
	ShadingPct62 ShadingPattern = "pct62"
	// ShadingPct65 uses 65% pattern.
	ShadingPct65 ShadingPattern = "pct65"
	// ShadingPct70 uses 70% pattern.
	ShadingPct70 ShadingPattern = "pct70"
	// ShadingPct75 uses 75% pattern.
	ShadingPct75 ShadingPattern = "pct75"
	// ShadingPct80 uses 80% pattern.
	ShadingPct80 ShadingPattern = "pct80"
	// ShadingPct85 uses 85% pattern.
	ShadingPct85 ShadingPattern = "pct85"
	// ShadingPct87 uses 87.5% pattern.
	ShadingPct87 ShadingPattern = "pct87"
	// ShadingPct90 uses 90% pattern.
	ShadingPct90 ShadingPattern = "pct90"
	// ShadingPct95 uses 95% pattern.
	ShadingPct95 ShadingPattern = "pct95"
)

// HeaderFooterType represents header/footer type values.
type HeaderFooterType string

const (
	// HeaderFooterDefault is the default header/footer.
	HeaderFooterDefault HeaderFooterType = "default"
	// HeaderFooterFirst is the first page header/footer.
	HeaderFooterFirst HeaderFooterType = "first"
	// HeaderFooterEven is the even page header/footer.
	HeaderFooterEven HeaderFooterType = "even"
)

// PageOrientation represents page orientation values.
type PageOrientation string

const (
	// PageOrientationPortrait is portrait orientation.
	PageOrientationPortrait PageOrientation = "portrait"
	// PageOrientationLandscape is landscape orientation.
	PageOrientationLandscape PageOrientation = "landscape"
)

// TextAlignmentValue represents vertical text alignment within a line (w:textAlignment).
type TextAlignmentValue string

const (
	// TextAlignmentAuto uses automatic text alignment.
	TextAlignmentAuto TextAlignmentValue = "auto"
	// TextAlignmentTop aligns text to the top of the line.
	TextAlignmentTop TextAlignmentValue = "top"
	// TextAlignmentCenter centers text vertically within the line.
	TextAlignmentCenter TextAlignmentValue = "center"
	// TextAlignmentBaseline aligns text to the baseline.
	TextAlignmentBaseline TextAlignmentValue = "baseline"
	// TextAlignmentBottom aligns text to the bottom of the line.
	TextAlignmentBottom TextAlignmentValue = "bottom"
)

// EmphasisMarkValue represents emphasis mark type values (w:em).
type EmphasisMarkValue string

const (
	// EmphasisNone removes emphasis mark.
	EmphasisNone EmphasisMarkValue = "none"
	// EmphasisDot applies dot emphasis mark.
	EmphasisDot EmphasisMarkValue = "dot"
	// EmphasisComma applies comma emphasis mark.
	EmphasisComma EmphasisMarkValue = "comma"
	// EmphasisCircle applies circle emphasis mark.
	EmphasisCircle EmphasisMarkValue = "circle"
	// EmphasisUnderDot applies under dot emphasis mark.
	EmphasisUnderDot EmphasisMarkValue = "underDot"
)
