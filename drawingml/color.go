// and effects.
//
//nolint:revive // This file intentionally contains many public types for OOXML color enums.
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PresetColorValue represents preset color values from the DrawingML specification.
// These are predefined color names that can be used without specifying RGB values.
type PresetColorValue string

// Preset color values defined in DrawingML.
const (
	PresetColorBlack                PresetColorValue = "black"
	PresetColorWhite                PresetColorValue = "white"
	PresetColorRed                  PresetColorValue = "red"
	PresetColorGreen                PresetColorValue = "green"
	PresetColorBlue                 PresetColorValue = "blue"
	PresetColorYellow               PresetColorValue = "yellow"
	PresetColorCyan                 PresetColorValue = "cyan"
	PresetColorMagenta              PresetColorValue = "magenta"
	PresetColorDarkBlue             PresetColorValue = "dkBlue"
	PresetColorDarkCyan             PresetColorValue = "dkCyan"
	PresetColorDarkGreen            PresetColorValue = "dkGreen"
	PresetColorDarkMagenta          PresetColorValue = "dkMagenta"
	PresetColorDarkRed              PresetColorValue = "dkRed"
	PresetColorDarkYellow           PresetColorValue = "dkYellow"
	PresetColorDarkGray             PresetColorValue = "dkGray"
	PresetColorLightGray            PresetColorValue = "ltGray"
	PresetColorAliceBlue            PresetColorValue = "aliceBlue"
	PresetColorAntiqueWhite         PresetColorValue = "antiqueWhite"
	PresetColorAqua                 PresetColorValue = "aqua"
	PresetColorAquamarine           PresetColorValue = "aquamarine"
	PresetColorAzure                PresetColorValue = "azure"
	PresetColorBeige                PresetColorValue = "beige"
	PresetColorBisque               PresetColorValue = "bisque"
	PresetColorBlanchedAlmond       PresetColorValue = "blanchedAlmond"
	PresetColorBlueViolet           PresetColorValue = "blueViolet"
	PresetColorBrown                PresetColorValue = "brown"
	PresetColorBurlyWood            PresetColorValue = "burlyWood"
	PresetColorCadetBlue            PresetColorValue = "cadetBlue"
	PresetColorChartreuse           PresetColorValue = "chartreuse"
	PresetColorChocolate            PresetColorValue = "chocolate"
	PresetColorCoral                PresetColorValue = "coral"
	PresetColorCornflowerBlue       PresetColorValue = "cornflowerBlue"
	PresetColorCornsilk             PresetColorValue = "cornsilk"
	PresetColorCrimson              PresetColorValue = "crimson"
	PresetColorDarkGoldenrod        PresetColorValue = "dkGoldenrod"
	PresetColorDarkKhaki            PresetColorValue = "dkKhaki"
	PresetColorDarkOliveGreen       PresetColorValue = "dkOliveGreen"
	PresetColorDarkOrange           PresetColorValue = "dkOrange"
	PresetColorDarkOrchid           PresetColorValue = "dkOrchid"
	PresetColorDarkSalmon           PresetColorValue = "dkSalmon"
	PresetColorDarkSeaGreen         PresetColorValue = "dkSeaGreen"
	PresetColorDarkSlateBlue        PresetColorValue = "dkSlateBlue"
	PresetColorDarkSlateGray        PresetColorValue = "dkSlateGray"
	PresetColorDarkTurquoise        PresetColorValue = "dkTurquoise"
	PresetColorDarkViolet           PresetColorValue = "dkViolet"
	PresetColorDeepPink             PresetColorValue = "deepPink"
	PresetColorDeepSkyBlue          PresetColorValue = "deepSkyBlue"
	PresetColorDimGray              PresetColorValue = "dimGray"
	PresetColorDodgerBlue           PresetColorValue = "dodgerBlue"
	PresetColorFirebrick            PresetColorValue = "firebrick"
	PresetColorFloralWhite          PresetColorValue = "floralWhite"
	PresetColorForestGreen          PresetColorValue = "forestGreen"
	PresetColorFuchsia              PresetColorValue = "fuchsia"
	PresetColorGainsboro            PresetColorValue = "gainsboro"
	PresetColorGhostWhite           PresetColorValue = "ghostWhite"
	PresetColorGold                 PresetColorValue = "gold"
	PresetColorGoldenrod            PresetColorValue = "goldenrod"
	PresetColorGray                 PresetColorValue = "gray"
	PresetColorGreenYellow          PresetColorValue = "greenYellow"
	PresetColorHoneydew             PresetColorValue = "honeydew"
	PresetColorHotPink              PresetColorValue = "hotPink"
	PresetColorIndianRed            PresetColorValue = "indianRed"
	PresetColorIndigo               PresetColorValue = "indigo"
	PresetColorIvory                PresetColorValue = "ivory"
	PresetColorKhaki                PresetColorValue = "khaki"
	PresetColorLavender             PresetColorValue = "lavender"
	PresetColorLavenderBlush        PresetColorValue = "lavenderBlush"
	PresetColorLawnGreen            PresetColorValue = "lawnGreen"
	PresetColorLemonChiffon         PresetColorValue = "lemonChiffon"
	PresetColorLightBlue            PresetColorValue = "ltBlue"
	PresetColorLightCoral           PresetColorValue = "ltCoral"
	PresetColorLightCyan            PresetColorValue = "ltCyan"
	PresetColorLightGoldenrodYellow PresetColorValue = "ltGoldenrodYellow"
	PresetColorLightGreen           PresetColorValue = "ltGreen"
	PresetColorLightPink            PresetColorValue = "ltPink"
	PresetColorLightSalmon          PresetColorValue = "ltSalmon"
	PresetColorLightSeaGreen        PresetColorValue = "ltSeaGreen"
	PresetColorLightSkyBlue         PresetColorValue = "ltSkyBlue"
	PresetColorLightSlateGray       PresetColorValue = "ltSlateGray"
	PresetColorLightSteelBlue       PresetColorValue = "ltSteelBlue"
	PresetColorLightYellow          PresetColorValue = "ltYellow"
	PresetColorLime                 PresetColorValue = "lime"
	PresetColorLimeGreen            PresetColorValue = "limeGreen"
	PresetColorLinen                PresetColorValue = "linen"
	PresetColorMaroon               PresetColorValue = "maroon"
	PresetColorMediumAquamarine     PresetColorValue = "medAquamarine"
	PresetColorMediumBlue           PresetColorValue = "medBlue"
	PresetColorMediumOrchid         PresetColorValue = "medOrchid"
	PresetColorMediumPurple         PresetColorValue = "medPurple"
	PresetColorMediumSeaGreen       PresetColorValue = "medSeaGreen"
	PresetColorMediumSlateBlue      PresetColorValue = "medSlateBlue"
	PresetColorMediumSpringGreen    PresetColorValue = "medSpringGreen"
	PresetColorMediumTurquoise      PresetColorValue = "medTurquoise"
	PresetColorMediumVioletRed      PresetColorValue = "medVioletRed"
	PresetColorMidnightBlue         PresetColorValue = "midnightBlue"
	PresetColorMintCream            PresetColorValue = "mintCream"
	PresetColorMistyRose            PresetColorValue = "mistyRose"
	PresetColorMoccasin             PresetColorValue = "moccasin"
	PresetColorNavajoWhite          PresetColorValue = "navajoWhite"
	PresetColorNavy                 PresetColorValue = "navy"
	PresetColorOldLace              PresetColorValue = "oldLace"
	PresetColorOlive                PresetColorValue = "olive"
	PresetColorOliveDrab            PresetColorValue = "oliveDrab"
	PresetColorOrange               PresetColorValue = "orange"
	PresetColorOrangeRed            PresetColorValue = "orangeRed"
	PresetColorOrchid               PresetColorValue = "orchid"
	PresetColorPaleGoldenrod        PresetColorValue = "paleGoldenrod"
	PresetColorPaleGreen            PresetColorValue = "paleGreen"
	PresetColorPaleTurquoise        PresetColorValue = "paleTurquoise"
	PresetColorPaleVioletRed        PresetColorValue = "paleVioletRed"
	PresetColorPapayaWhip           PresetColorValue = "papayaWhip"
	PresetColorPeachPuff            PresetColorValue = "peachPuff"
	PresetColorPeru                 PresetColorValue = "peru"
	PresetColorPink                 PresetColorValue = "pink"
	PresetColorPlum                 PresetColorValue = "plum"
	PresetColorPowderBlue           PresetColorValue = "powderBlue"
	PresetColorPurple               PresetColorValue = "purple"
	PresetColorRosyBrown            PresetColorValue = "rosyBrown"
	PresetColorRoyalBlue            PresetColorValue = "royalBlue"
	PresetColorSaddleBrown          PresetColorValue = "saddleBrown"
	PresetColorSalmon               PresetColorValue = "salmon"
	PresetColorSandyBrown           PresetColorValue = "sandyBrown"
	PresetColorSeaGreen             PresetColorValue = "seaGreen"
	PresetColorSeaShell             PresetColorValue = "seaShell"
	PresetColorSienna               PresetColorValue = "sienna"
	PresetColorSilver               PresetColorValue = "silver"
	PresetColorSkyBlue              PresetColorValue = "skyBlue"
	PresetColorSlateBlue            PresetColorValue = "slateBlue"
	PresetColorSlateGray            PresetColorValue = "slateGray"
	PresetColorSnow                 PresetColorValue = "snow"
	PresetColorSpringGreen          PresetColorValue = "springGreen"
	PresetColorSteelBlue            PresetColorValue = "steelBlue"
	PresetColorTan                  PresetColorValue = "tan"
	PresetColorTeal                 PresetColorValue = "teal"
	PresetColorThistle              PresetColorValue = "thistle"
	PresetColorTomato               PresetColorValue = "tomato"
	PresetColorTurquoise            PresetColorValue = "turquoise"
	PresetColorViolet               PresetColorValue = "violet"
	PresetColorWheat                PresetColorValue = "wheat"
	PresetColorWhiteSmoke           PresetColorValue = "whiteSmoke"
	PresetColorYellowGreen          PresetColorValue = "yellowGreen"
)

// SchemeColorValue represents theme scheme color values.
// These reference colors defined in the document theme.
type SchemeColorValue string

// Scheme color values from document themes.
const (
	// SchemeColorDark1 is the first dark color from the theme.
	SchemeColorDark1 SchemeColorValue = "dk1"
	// SchemeColorLight1 is the first light color from the theme.
	SchemeColorLight1 SchemeColorValue = "lt1"
	// SchemeColorDark2 is the second dark color from the theme.
	SchemeColorDark2 SchemeColorValue = "dk2"
	// SchemeColorLight2 is the second light color from the theme.
	SchemeColorLight2 SchemeColorValue = "lt2"
	// SchemeColorAccent1 is the first accent color.
	SchemeColorAccent1 SchemeColorValue = "accent1"
	// SchemeColorAccent2 is the second accent color.
	SchemeColorAccent2 SchemeColorValue = "accent2"
	// SchemeColorAccent3 is the third accent color.
	SchemeColorAccent3 SchemeColorValue = "accent3"
	// SchemeColorAccent4 is the fourth accent color.
	SchemeColorAccent4 SchemeColorValue = "accent4"
	// SchemeColorAccent5 is the fifth accent color.
	SchemeColorAccent5 SchemeColorValue = "accent5"
	// SchemeColorAccent6 is the sixth accent color.
	SchemeColorAccent6 SchemeColorValue = "accent6"
	// SchemeColorHyperlink is the hyperlink color.
	SchemeColorHyperlink SchemeColorValue = "hlink"
	// SchemeColorFollowedHyperlink is the followed hyperlink color.
	SchemeColorFollowedHyperlink SchemeColorValue = "folHlink"
	// SchemeColorPhColor is the placeholder color.
	SchemeColorPhColor SchemeColorValue = "phClr"
	// SchemeColorBackground1 is the first background color.
	SchemeColorBackground1 SchemeColorValue = "bg1"
	// SchemeColorText1 is the first text color.
	SchemeColorText1 SchemeColorValue = "tx1"
	// SchemeColorBackground2 is the second background color.
	SchemeColorBackground2 SchemeColorValue = "bg2"
	// SchemeColorText2 is the second text color.
	SchemeColorText2 SchemeColorValue = "tx2"
)

// SystemColorValue represents Windows system color values.
// These reference colors from the operating system color scheme.
type SystemColorValue string

// System color values from the Windows color scheme.
const (
	SystemColorScrollBar               SystemColorValue = "scrollBar"
	SystemColorBackground              SystemColorValue = "background"
	SystemColorActiveCaption           SystemColorValue = "activeCaption"
	SystemColorInactiveCaption         SystemColorValue = "inactiveCaption"
	SystemColorMenu                    SystemColorValue = "menu"
	SystemColorWindow                  SystemColorValue = "window"
	SystemColorWindowFrame             SystemColorValue = "windowFrame"
	SystemColorMenuText                SystemColorValue = "menuText"
	SystemColorWindowText              SystemColorValue = "windowText"
	SystemColorCaptionText             SystemColorValue = "captionText"
	SystemColorActiveBorder            SystemColorValue = "activeBorder"
	SystemColorInactiveBorder          SystemColorValue = "inactiveBorder"
	SystemColorAppWorkspace            SystemColorValue = "appWorkspace"
	SystemColorHighlight               SystemColorValue = "highlight"
	SystemColorHighlightText           SystemColorValue = "highlightText"
	SystemColorBtnFace                 SystemColorValue = "btnFace"
	SystemColorBtnShadow               SystemColorValue = "btnShadow"
	SystemColorGrayText                SystemColorValue = "grayText"
	SystemColorBtnText                 SystemColorValue = "btnText"
	SystemColorInactiveCaptionText     SystemColorValue = "inactiveCaptionText"
	SystemColorBtnHighlight            SystemColorValue = "btnHighlight"
	SystemColor3DDarkShadow            SystemColorValue = "3dDkShadow"
	SystemColor3DLight                 SystemColorValue = "3dLight"
	SystemColorInfoText                SystemColorValue = "infoText"
	SystemColorInfoBk                  SystemColorValue = "infoBk"
	SystemColorHotLight                SystemColorValue = "hotLight"
	SystemColorGradientActiveCaption   SystemColorValue = "gradientActiveCaption"
	SystemColorGradientInactiveCaption SystemColorValue = "gradientInactiveCaption"
	SystemColorMenuHighlight           SystemColorValue = "menuHighlight"
	SystemColorMenuBar                 SystemColorValue = "menuBar"
)

// RgbColor represents an sRGB color (a:srgbClr) with a hexadecimal value.
// The color is specified as a 6-character hex string (RRGGBB format).
type RgbColor struct {
	*openxml.CompositeElementBase
}

// NewRgbColor creates a new RGB color element with the specified hex value.
// The hex value should be a 6-character string like "FF0000" for red.
func NewRgbColor(hexValue string) *RgbColor {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"srgbClr",
		PrefixMain,
	)
	rgb := &RgbColor{CompositeElementBase: elem}
	rgb.SetValue(hexValue)

	return rgb
}

// NewRgbColorFromRGB creates a new RGB color from individual R, G, B values (0-255).
func NewRgbColorFromRGB(r, g, b uint8) *RgbColor {
	hex := colorToHex(r, g, b)

	return NewRgbColor(hex)
}

// colorToHex converts RGB values to a hex string.
func colorToHex(r, g, b uint8) string {
	const hexDigits = "0123456789ABCDEF"

	return string([]byte{
		hexDigits[r>>4], hexDigits[r&0x0f],
		hexDigits[g>>4], hexDigits[g&0x0f],
		hexDigits[b>>4], hexDigits[b&0x0f],
	})
}

// Value returns the hex color value (RRGGBB format).
func (c *RgbColor) Value() string {
	attr, found := c.GetAttribute("val", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetValue sets the hex color value (RRGGBB format).
func (c *RgbColor) SetValue(hex string) {
	c.SetAttribute(
		openxml.NewAttribute("", "val", "", hex),
	)
}

// RGB returns the individual R, G, B components (0-255).
func (c *RgbColor) RGB() (r, g, b uint8) {
	hex := c.Value()
	if len(hex) < 6 {
		return 0, 0, 0
	}
	rVal, _ := strconv.ParseUint(hex[0:2], 16, 8)
	gVal, _ := strconv.ParseUint(hex[2:4], 16, 8)
	bVal, _ := strconv.ParseUint(hex[4:6], 16, 8)

	return uint8(rVal), uint8(gVal), uint8(bVal)
}

// AddTint adds a tint transformation to the color.
// The value is in 1000ths of a percent (0-100000).
func (c *RgbColor) AddTint(value int) {
	c.addTransformation("tint", value)
}

// AddShade adds a shade transformation to the color.
// The value is in 1000ths of a percent (0-100000).
func (c *RgbColor) AddShade(value int) {
	c.addTransformation("shade", value)
}

// AddAlpha adds an alpha (transparency) transformation to the color.
// The value is in 1000ths of a percent (0-100000), where 100000 is fully opaque.
func (c *RgbColor) AddAlpha(value int) {
	c.addTransformation("alpha", value)
}

// AddSaturationModulation adds a saturation modulation transformation.
// The value is in 1000ths of a percent.
func (c *RgbColor) AddSaturationModulation(
	value int,
) {
	c.addTransformation("satMod", value)
}

// AddLuminanceModulation adds a luminance modulation transformation.
// The value is in 1000ths of a percent.
func (c *RgbColor) AddLuminanceModulation(
	value int,
) {
	c.addTransformation("lumMod", value)
}

// AddLuminanceOffset adds a luminance offset transformation.
// The value is in 1000ths of a percent (-100000 to 100000).
func (c *RgbColor) AddLuminanceOffset(value int) {
	c.addTransformation("lumOff", value)
}

// addTransformation adds a color transformation child element.
func (c *RgbColor) addTransformation(
	name string,
	value int,
) {
	trans := openxml.NewLeafElement(
		NamespaceMain,
		name,
		PrefixMain,
	)
	trans.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			strconv.Itoa(value),
		),
	)
	c.AppendChild(trans)
}

// Clone creates a deep copy of this RgbColor element.
func (c *RgbColor) Clone() openxml.Element {
	return &RgbColor{
		CompositeElementBase: c.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// HslColor represents an HSL color (a:hslClr) with hue, saturation, and luminance values.
type HslColor struct {
	*openxml.CompositeElementBase
}

// NewHslColor creates a new HSL color element.
// hue is in 60000ths of a degree (0 to 21599999)
// sat and lum are in 1000ths of a percent (0 to 100000).
func NewHslColor(hue, sat, lum int) *HslColor {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"hslClr",
		PrefixMain,
	)
	hsl := &HslColor{CompositeElementBase: elem}
	hsl.SetHue(hue)
	hsl.SetSaturation(sat)
	hsl.SetLuminance(lum)

	return hsl
}

// NewHslColorFromDegrees creates a new HSL color from degree values.
// hue is in degrees (0-359), sat and lum are in percent (0-100).
func NewHslColorFromDegrees(
	hueDegrees float64,
	satPercent, lumPercent float64,
) *HslColor {
	hue := DegreesToAngleUnits(hueDegrees)
	sat := PercentToUnits(satPercent)
	lum := PercentToUnits(lumPercent)

	return NewHslColor(hue, sat, lum)
}

// Hue returns the hue value in 60000ths of a degree.
func (c *HslColor) Hue() int {
	attr, found := c.GetAttribute("hue", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetHue sets the hue value in 60000ths of a degree.
func (c *HslColor) SetHue(hue int) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"hue",
			"",
			strconv.Itoa(hue),
		),
	)
}

// Saturation returns the saturation value in 1000ths of a percent.
func (c *HslColor) Saturation() int {
	attr, found := c.GetAttribute("sat", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetSaturation sets the saturation value in 1000ths of a percent.
func (c *HslColor) SetSaturation(sat int) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"sat",
			"",
			strconv.Itoa(sat),
		),
	)
}

// Luminance returns the luminance value in 1000ths of a percent.
func (c *HslColor) Luminance() int {
	attr, found := c.GetAttribute("lum", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetLuminance sets the luminance value in 1000ths of a percent.
func (c *HslColor) SetLuminance(lum int) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"lum",
			"",
			strconv.Itoa(lum),
		),
	)
}

// HueDegrees returns the hue value in degrees (0-359).
func (c *HslColor) HueDegrees() float64 {
	return AngleUnitsToDegrees(c.Hue())
}

// SaturationPercent returns the saturation value as a percentage (0-100).
func (c *HslColor) SaturationPercent() float64 {
	return UnitsToPercent(c.Saturation())
}

// LuminancePercent returns the luminance value as a percentage (0-100).
func (c *HslColor) LuminancePercent() float64 {
	return UnitsToPercent(c.Luminance())
}

// Clone creates a deep copy of this HslColor element.
func (c *HslColor) Clone() openxml.Element {
	return &HslColor{
		CompositeElementBase: c.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// PresetColor represents a preset color (a:prstClr) with a named color value.
type PresetColor struct {
	*openxml.CompositeElementBase
}

// NewPresetColor creates a new preset color element with the specified color value.
func NewPresetColor(
	value PresetColorValue,
) *PresetColor {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"prstClr",
		PrefixMain,
	)
	pc := &PresetColor{CompositeElementBase: elem}
	pc.SetValue(value)

	return pc
}

// Value returns the preset color value.
func (c *PresetColor) Value() PresetColorValue {
	attr, found := c.GetAttribute("val", "")
	if !found {
		return ""
	}

	return PresetColorValue(attr.Value())
}

// SetValue sets the preset color value.
func (c *PresetColor) SetValue(
	value PresetColorValue,
) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			string(value),
		),
	)
}

// Clone creates a deep copy of this PresetColor element.
func (c *PresetColor) Clone() openxml.Element {
	return &PresetColor{
		CompositeElementBase: c.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// SchemeColor represents a theme scheme color (a:schemeClr) referencing
// a color from the document theme.
type SchemeColor struct {
	*openxml.CompositeElementBase
}

// NewSchemeColor creates a new scheme color element with the specified color value.
func NewSchemeColor(
	value SchemeColorValue,
) *SchemeColor {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"schemeClr",
		PrefixMain,
	)
	sc := &SchemeColor{CompositeElementBase: elem}
	sc.SetValue(value)

	return sc
}

// Value returns the scheme color value.
func (c *SchemeColor) Value() SchemeColorValue {
	attr, found := c.GetAttribute("val", "")
	if !found {
		return ""
	}

	return SchemeColorValue(attr.Value())
}

// SetValue sets the scheme color value.
func (c *SchemeColor) SetValue(
	value SchemeColorValue,
) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			string(value),
		),
	)
}

// AddTint adds a tint transformation to the color.
// The value is in 1000ths of a percent (0-100000).
func (c *SchemeColor) AddTint(value int) {
	c.addTransformation("tint", value)
}

// AddShade adds a shade transformation to the color.
// The value is in 1000ths of a percent (0-100000).
func (c *SchemeColor) AddShade(value int) {
	c.addTransformation("shade", value)
}

// AddAlpha adds an alpha (transparency) transformation to the color.
// The value is in 1000ths of a percent (0-100000), where 100000 is fully opaque.
func (c *SchemeColor) AddAlpha(value int) {
	c.addTransformation("alpha", value)
}

// AddSaturationModulation adds a saturation modulation transformation.
// The value is in 1000ths of a percent.
func (c *SchemeColor) AddSaturationModulation(
	value int,
) {
	c.addTransformation("satMod", value)
}

// AddLuminanceModulation adds a luminance modulation transformation.
// The value is in 1000ths of a percent.
func (c *SchemeColor) AddLuminanceModulation(
	value int,
) {
	c.addTransformation("lumMod", value)
}

// AddLuminanceOffset adds a luminance offset transformation.
// The value is in 1000ths of a percent (-100000 to 100000).
func (c *SchemeColor) AddLuminanceOffset(
	value int,
) {
	c.addTransformation("lumOff", value)
}

// addTransformation adds a color transformation child element.
func (c *SchemeColor) addTransformation(
	name string,
	value int,
) {
	trans := openxml.NewLeafElement(
		NamespaceMain,
		name,
		PrefixMain,
	)
	trans.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			strconv.Itoa(value),
		),
	)
	c.AppendChild(trans)
}

// Clone creates a deep copy of this SchemeColor element.
func (c *SchemeColor) Clone() openxml.Element {
	return &SchemeColor{
		CompositeElementBase: c.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// SystemColor represents a system color (a:sysClr) referencing a Windows
// system color from the operating system color scheme.
type SystemColor struct {
	*openxml.CompositeElementBase
}

// NewSystemColor creates a new system color element with the specified color value.
func NewSystemColor(
	value SystemColorValue,
) *SystemColor {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"sysClr",
		PrefixMain,
	)
	sc := &SystemColor{CompositeElementBase: elem}
	sc.SetValue(value)

	return sc
}

// NewSystemColorWithLastColor creates a new system color element with a fallback color.
// The lastColor is the last computed RGB value of this system color (RRGGBB hex).
func NewSystemColorWithLastColor(
	value SystemColorValue,
	lastColor string,
) *SystemColor {
	sc := NewSystemColor(value)
	sc.SetLastColor(lastColor)

	return sc
}

// Value returns the system color value.
func (c *SystemColor) Value() SystemColorValue {
	attr, found := c.GetAttribute("val", "")
	if !found {
		return ""
	}

	return SystemColorValue(attr.Value())
}

// SetValue sets the system color value.
func (c *SystemColor) SetValue(
	value SystemColorValue,
) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			string(value),
		),
	)
}

// LastColor returns the last computed RGB value of this system color.
// This is a 6-character hex string (RRGGBB format).
func (c *SystemColor) LastColor() string {
	attr, found := c.GetAttribute("lastClr", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetLastColor sets the last computed RGB value of this system color.
func (c *SystemColor) SetLastColor(hex string) {
	if hex == "" {
		c.RemoveAttribute("lastClr", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"lastClr",
			"",
			hex,
		),
	)
}

// AddTint adds a tint transformation to the color.
func (c *SystemColor) AddTint(value int) {
	c.addTransformation("tint", value)
}

// AddShade adds a shade transformation to the color.
func (c *SystemColor) AddShade(value int) {
	c.addTransformation("shade", value)
}

// AddAlpha adds an alpha (transparency) transformation to the color.
func (c *SystemColor) AddAlpha(value int) {
	c.addTransformation("alpha", value)
}

// addTransformation adds a color transformation child element.
func (c *SystemColor) addTransformation(
	name string,
	value int,
) {
	trans := openxml.NewLeafElement(
		NamespaceMain,
		name,
		PrefixMain,
	)
	trans.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			strconv.Itoa(value),
		),
	)
	c.AppendChild(trans)
}

// Clone creates a deep copy of this SystemColor element.
func (c *SystemColor) Clone() openxml.Element {
	return &SystemColor{
		CompositeElementBase: c.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ColorTransformationType represents the type of color transformation.
type ColorTransformationType string

// Color transformation types.
const (
	// TransformTint adds white to the color.
	TransformTint ColorTransformationType = "tint"
	// TransformShade adds black to the color.
	TransformShade ColorTransformationType = "shade"
	// TransformComplement returns the complement of the color.
	TransformComplement ColorTransformationType = "comp"
	// TransformInverse returns the inverse of the color.
	TransformInverse ColorTransformationType = "inv"
	// TransformGrayscale converts the color to grayscale.
	TransformGrayscale ColorTransformationType = "gray"
	// TransformAlpha sets the alpha (transparency) of the color.
	TransformAlpha ColorTransformationType = "alpha"
	// TransformAlphaOffset offsets the alpha by a percentage.
	TransformAlphaOffset ColorTransformationType = "alphaOff"
	// TransformAlphaModulation modulates the alpha by a percentage.
	TransformAlphaModulation ColorTransformationType = "alphaMod"
	// TransformHue sets the hue of the color.
	TransformHue ColorTransformationType = "hue"
	// TransformHueOffset offsets the hue by an angle.
	TransformHueOffset ColorTransformationType = "hueOff"
	// TransformHueModulation modulates the hue by a percentage.
	TransformHueModulation ColorTransformationType = "hueMod"
	// TransformSaturation sets the saturation of the color.
	TransformSaturation ColorTransformationType = "sat"
	// TransformSaturationOffset offsets the saturation.
	TransformSaturationOffset ColorTransformationType = "satOff"
	// TransformSaturationModulation modulates the saturation.
	TransformSaturationModulation ColorTransformationType = "satMod"
	// TransformLuminance sets the luminance of the color.
	TransformLuminance ColorTransformationType = "lum"
	// TransformLuminanceOffset offsets the luminance.
	TransformLuminanceOffset ColorTransformationType = "lumOff"
	// TransformLuminanceModulation modulates the luminance.
	TransformLuminanceModulation ColorTransformationType = "lumMod"
	// TransformRed sets the red component.
	TransformRed ColorTransformationType = "red"
	// TransformRedOffset offsets the red component.
	TransformRedOffset ColorTransformationType = "redOff"
	// TransformRedModulation modulates the red component.
	TransformRedModulation ColorTransformationType = "redMod"
	// TransformGreen sets the green component.
	TransformGreen ColorTransformationType = "green"
	// TransformGreenOffset offsets the green component.
	TransformGreenOffset ColorTransformationType = "greenOff"
	// TransformGreenModulation modulates the green component.
	TransformGreenModulation ColorTransformationType = "greenMod"
	// TransformBlue sets the blue component.
	TransformBlue ColorTransformationType = "blue"
	// TransformBlueOffset offsets the blue component.
	TransformBlueOffset ColorTransformationType = "blueOff"
	// TransformBlueModulation modulates the blue component.
	TransformBlueModulation ColorTransformationType = "blueMod"
	// TransformGamma applies gamma correction.
	TransformGamma ColorTransformationType = "gamma"
	// TransformInverseGamma applies inverse gamma correction.
	TransformInverseGamma ColorTransformationType = "invGamma"
)
