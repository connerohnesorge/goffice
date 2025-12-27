package drawing

import (
	"strconv"
	"strings"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// DrawingMLColorType represents the type of DrawingML color specification.
type DrawingMLColorType int

const (
	// ColorTypeUnknown indicates an unknown or unrecognized color type.
	ColorTypeUnknown DrawingMLColorType = iota
	// ColorTypeSRGB indicates an sRGB color (a:srgbClr).
	ColorTypeSRGB
	// ColorTypeScheme indicates a theme scheme color (a:schemeClr).
	ColorTypeScheme
	// ColorTypeHSL indicates an HSL color (a:hslClr).
	ColorTypeHSL
	// ColorTypeSystem indicates a system color (a:sysClr).
	ColorTypeSystem
	// ColorTypePreset indicates a preset color (a:prstClr).
	ColorTypePreset
	// ColorTypeSCRGB indicates a scRGB color (a:scrgbClr).
	ColorTypeSCRGB
)

// DrawingMLColor represents a color from DrawingML elements that can be
// resolved to an RGBA Color for PDF rendering.
type DrawingMLColor struct {
	// Type is the DrawingML color type.
	Type DrawingMLColorType

	// Value is the raw value from the element (hex for sRGB, name for scheme/preset/system).
	Value string

	// HSL values for HSL color type (in DrawingML units).
	Hue        int // 0-21600000 (60000ths of a degree)
	Saturation int // 0-100000 (1000ths of a percent)
	Luminance  int // 0-100000 (1000ths of a percent)

	// SCRGB values for scRGB color type (in percentage units).
	SCR int // 0-100000 (percentage for red)
	SCG int // 0-100000 (percentage for green)
	SCB int // 0-100000 (percentage for blue)

	// Transformations are color modifications to apply.
	Transformations []ColorTransformation

	// LastColor is the fallback color for system colors (hex value).
	LastColor string
}

// ColorTransformation represents a DrawingML color transformation.
type ColorTransformation struct {
	// Type is the transformation type (tint, shade, satMod, lumMod, etc.).
	Type string
	// Value is the transformation value in DrawingML units.
	// For most transformations, this is in 1000ths of a percent (0-100000).
	// For hue offsets, this is in 60000ths of a degree.
	Value int
}

// ParseDrawingMLColor extracts color information from a DrawingML color element.
// Supports srgbClr, schemeClr, hslClr, sysClr, prstClr, and scrgbClr elements.
func ParseDrawingMLColor(
	elem openxml.Element,
) *DrawingMLColor {
	if elem == nil {
		return nil
	}

	localName := elem.QName().LocalName()

	color := &DrawingMLColor{}

	switch localName {
	case "srgbClr":
		color.Type = ColorTypeSRGB
		if attr, found := elem.GetAttribute("val", ""); found {
			color.Value = attr.Value()
		}

	case "schemeClr":
		color.Type = ColorTypeScheme
		if attr, found := elem.GetAttribute("val", ""); found {
			color.Value = attr.Value()
		}

	case "hslClr":
		color.Type = ColorTypeHSL
		if attr, found := elem.GetAttribute("hue", ""); found {
			color.Hue, _ = strconv.Atoi(
				attr.Value(),
			)
		}
		if attr, found := elem.GetAttribute("sat", ""); found {
			color.Saturation, _ = strconv.Atoi(
				attr.Value(),
			)
		}
		if attr, found := elem.GetAttribute("lum", ""); found {
			color.Luminance, _ = strconv.Atoi(
				attr.Value(),
			)
		}

	case "sysClr":
		color.Type = ColorTypeSystem
		if attr, found := elem.GetAttribute("val", ""); found {
			color.Value = attr.Value()
		}
		if attr, found := elem.GetAttribute("lastClr", ""); found {
			color.LastColor = attr.Value()
		}

	case "prstClr":
		color.Type = ColorTypePreset
		if attr, found := elem.GetAttribute("val", ""); found {
			color.Value = attr.Value()
		}

	case "scrgbClr":
		color.Type = ColorTypeSCRGB
		if attr, found := elem.GetAttribute("r", ""); found {
			color.SCR, _ = strconv.Atoi(
				attr.Value(),
			)
		}
		if attr, found := elem.GetAttribute("g", ""); found {
			color.SCG, _ = strconv.Atoi(
				attr.Value(),
			)
		}
		if attr, found := elem.GetAttribute("b", ""); found {
			color.SCB, _ = strconv.Atoi(
				attr.Value(),
			)
		}

	default:
		color.Type = ColorTypeUnknown

		return color
	}

	// Parse child transformations.
	color.Transformations = parseColorTransformations(
		elem,
	)

	return color
}

// parseColorTransformations extracts color transformation elements from a color element.
func parseColorTransformations(
	elem openxml.Element,
) []ColorTransformation {
	composite, ok := elem.(openxml.CompositeElement)
	if !ok {
		return nil
	}

	var transformations []ColorTransformation
	for child := range composite.Children() {
		localName := child.QName().LocalName()
		trans := ColorTransformation{
			Type: localName,
		}

		// Get the value attribute.
		if attr, found := child.GetAttribute("val", ""); found {
			trans.Value, _ = strconv.Atoi(
				attr.Value(),
			)
		}

		// Only add known transformation types.
		if isKnownTransformation(localName) {
			transformations = append(
				transformations,
				trans,
			)
		}
	}

	return transformations
}

// isKnownTransformation returns true if the transformation type is recognized.
func isKnownTransformation(name string) bool {
	switch name {
	case "tint", "shade", "comp", "inv", "gray",
		"alpha", "alphaOff", "alphaMod",
		"hue", "hueOff", "hueMod",
		"sat", "satOff", "satMod",
		"lum", "lumOff", "lumMod",
		"red", "redOff", "redMod",
		"green", "greenOff", "greenMod",
		"blue", "blueOff", "blueMod",
		"gamma", "invGamma":
		return true
	}

	return false
}

// Resolve converts a DrawingML color to an RGBA Color.
// The theme parameter is used to resolve scheme colors; pass nil for default colors.
func (c *DrawingMLColor) Resolve(
	theme *ThemePalette,
) Color {
	if c == nil {
		return Black
	}

	var baseColor Color

	switch c.Type {
	case ColorTypeSRGB:
		baseColor = parseSRGBHex(c.Value)

	case ColorTypeScheme:
		baseColor = resolveSchemeColor(
			c.Value,
			theme,
		)

	case ColorTypeHSL:
		baseColor = resolveHSLColor(
			c.Hue,
			c.Saturation,
			c.Luminance,
		)

	case ColorTypeSystem:
		baseColor = resolveSystemColor(
			c.Value,
			c.LastColor,
		)

	case ColorTypePreset:
		baseColor = resolvePresetColor(c.Value)

	case ColorTypeSCRGB:
		baseColor = resolveSCRGBColor(
			c.SCR,
			c.SCG,
			c.SCB,
		)

	case ColorTypeUnknown:
		return Black

	default:
		return Black
	}

	// Apply transformations in order.
	return applyTransformations(
		baseColor,
		c.Transformations,
	)
}

// parseSRGBHex parses a 6-character hex color string (RRGGBB).
func parseSRGBHex(hex string) Color {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return Black
	}

	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return Black
	}
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return Black
	}
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return Black
	}

	return NewRGB8(uint8(r), uint8(g), uint8(b))
}

// resolveSchemeColor converts a scheme color name to an RGBA color using the theme palette.
func resolveSchemeColor(
	name string,
	theme *ThemePalette,
) Color {
	if theme != nil {
		if color, ok := theme.GetColor(drawingml.SchemeColorValue(name)); ok {
			return color
		}
	}
	// Fall back to default theme colors.
	return DefaultThemePalette.GetColorOrDefault(
		drawingml.SchemeColorValue(name),
	)
}

// resolveHSLColor converts DrawingML HSL values to an RGBA color.
// hue is in 60000ths of a degree (0-21600000).
// sat and lum are in 1000ths of a percent (0-100000).
func resolveHSLColor(hue, sat, lum int) Color {
	// Convert from DrawingML units to standard HSL.
	h := float64(
		hue,
	) / 60000.0 // Convert to degrees (0-360)
	s := float64(
		sat,
	) / 100000.0 // Convert to 0-1 range
	l := float64(
		lum,
	) / 100000.0 // Convert to 0-1 range

	hsl := NewHSL(h, s, l)

	return hsl.ToRGB()
}

// resolveSystemColor converts a system color name to an RGBA color.
// Uses the lastColor fallback if available, otherwise returns a default.
func resolveSystemColor(
	name, lastColor string,
) Color {
	// If lastColor is provided, use it as the resolved color.
	if lastColor != "" {
		return parseSRGBHex(lastColor)
	}

	// Look up system color default.
	if color, ok := SystemColors[name]; ok {
		return color
	}

	return Black
}

// resolvePresetColor converts a preset color name to an RGBA color.
func resolvePresetColor(name string) Color {
	if color, ok := PresetColors[name]; ok {
		return color
	}

	return Black
}

// resolveSCRGBColor converts scRGB percentage values to an RGBA color.
// The values are in percentage units (0-100000 = 0-100%).
func resolveSCRGBColor(r, g, b int) Color {
	// scRGB values are in percentage units (1000ths of a percent).
	// 100000 = 100% = 1.0.
	rf := float64(r) / 100000.0
	gf := float64(g) / 100000.0
	bf := float64(b) / 100000.0

	// scRGB can have values > 1.0 for HDR, but we clamp for standard rendering.
	return NewRGB(rf, gf, bf)
}

// PresetColors maps DrawingML preset color names to RGBA colors.
// These are the standard predefined colors from the OOXML specification.
var PresetColors = map[string]Color{
	"black":     Black,
	"white":     White,
	"red":       Red,
	"green":     Green,
	"blue":      Blue,
	"yellow":    Yellow,
	"cyan":      Cyan,
	"magenta":   Magenta,
	"dkBlue":    DarkBlue,
	"dkCyan":    DarkCyan,
	"dkGreen":   DarkGreen,
	"dkMagenta": DarkMagenta,
	"dkRed":     DarkRed,
	"dkYellow": ParseHex(
		"#808000",
	), // Dark yellow/olive
	"dkGray":            DarkGray,
	"ltGray":            LightGray,
	"aliceBlue":         AliceBlue,
	"antiqueWhite":      AntiqueWhite,
	"aqua":              Aqua,
	"aquamarine":        Aquamarine,
	"azure":             Azure,
	"beige":             Beige,
	"bisque":            Bisque,
	"blanchedAlmond":    BlanchedAlmond,
	"blueViolet":        BlueViolet,
	"brown":             Brown,
	"burlyWood":         BurlyWood,
	"cadetBlue":         CadetBlue,
	"chartreuse":        Chartreuse,
	"chocolate":         Chocolate,
	"coral":             Coral,
	"cornflowerBlue":    CornflowerBlue,
	"cornsilk":          Cornsilk,
	"crimson":           Crimson,
	"dkGoldenrod":       DarkGoldenrod,
	"dkKhaki":           DarkKhaki,
	"dkOliveGreen":      DarkOliveGreen,
	"dkOrange":          DarkOrange,
	"dkOrchid":          DarkOrchid,
	"dkSalmon":          DarkSalmon,
	"dkSeaGreen":        DarkSeaGreen,
	"dkSlateBlue":       DarkSlateBlue,
	"dkSlateGray":       DarkSlateGray,
	"dkTurquoise":       DarkTurquoise,
	"dkViolet":          DarkViolet,
	"deepPink":          DeepPink,
	"deepSkyBlue":       DeepSkyBlue,
	"dimGray":           DimGray,
	"dodgerBlue":        DodgerBlue,
	"firebrick":         Firebrick,
	"floralWhite":       FloralWhite,
	"forestGreen":       ForestGreen,
	"fuchsia":           Fuchsia,
	"gainsboro":         Gainsboro,
	"ghostWhite":        GhostWhite,
	"gold":              Gold,
	"goldenrod":         Goldenrod,
	"gray":              Gray,
	"greenYellow":       GreenYellow,
	"honeydew":          Honeydew,
	"hotPink":           HotPink,
	"indianRed":         IndianRed,
	"indigo":            Indigo,
	"ivory":             Ivory,
	"khaki":             Khaki,
	"lavender":          Lavender,
	"lavenderBlush":     LavenderBlush,
	"lawnGreen":         LawnGreen,
	"lemonChiffon":      LemonChiffon,
	"ltBlue":            LightBlue,
	"ltCoral":           LightCoral,
	"ltCyan":            LightCyan,
	"ltGoldenrodYellow": LightGoldenrodYellow,
	"ltGreen":           LightGreen,
	"ltPink":            LightPink,
	"ltSalmon":          LightSalmon,
	"ltSeaGreen":        LightSeaGreen,
	"ltSkyBlue":         LightSkyBlue,
	"ltSlateGray":       LightSlateGray,
	"ltSteelBlue":       LightSteelBlue,
	"ltYellow":          LightYellow,
	"lime":              Lime,
	"limeGreen":         LimeGreen,
	"linen":             Linen,
	"maroon":            Maroon,
	"medAquamarine":     MediumAquamarine,
	"medBlue":           MediumBlue,
	"medOrchid":         MediumOrchid,
	"medPurple":         MediumPurple,
	"medSeaGreen":       MediumSeaGreen,
	"medSlateBlue":      MediumSlateBlue,
	"medSpringGreen":    MediumSpringGreen,
	"medTurquoise":      MediumTurquoise,
	"medVioletRed":      MediumVioletRed,
	"midnightBlue":      MidnightBlue,
	"mintCream":         MintCream,
	"mistyRose":         MistyRose,
	"moccasin":          Moccasin,
	"navajoWhite":       NavajoWhite,
	"navy":              Navy,
	"oldLace":           OldLace,
	"olive":             Olive,
	"oliveDrab":         OliveDrab,
	"orange":            Orange,
	"orangeRed":         OrangeRed,
	"orchid":            Orchid,
	"paleGoldenrod":     PaleGoldenrod,
	"paleGreen":         PaleGreen,
	"paleTurquoise":     PaleTurquoise,
	"paleVioletRed":     PaleVioletRed,
	"papayaWhip":        PapayaWhip,
	"peachPuff":         PeachPuff,
	"peru":              Peru,
	"pink":              Pink,
	"plum":              Plum,
	"powderBlue":        PowderBlue,
	"purple":            Purple,
	"rosyBrown":         RosyBrown,
	"royalBlue":         RoyalBlue,
	"saddleBrown":       SaddleBrown,
	"salmon":            Salmon,
	"sandyBrown":        SandyBrown,
	"seaGreen":          SeaGreen,
	"seaShell":          SeaShell,
	"sienna":            Sienna,
	"silver":            Silver,
	"skyBlue":           SkyBlue,
	"slateBlue":         SlateBlue,
	"slateGray":         SlateGray,
	"snow":              Snow,
	"springGreen":       SpringGreen,
	"steelBlue":         SteelBlue,
	"tan":               Tan,
	"teal":              Teal,
	"thistle":           Thistle,
	"tomato":            Tomato,
	"turquoise":         Turquoise,
	"violet":            Violet,
	"wheat":             Wheat,
	"whiteSmoke":        WhiteSmoke,
	"yellowGreen":       YellowGreen,
}

// SystemColors maps Windows system color names to default RGBA colors.
// These are fallback values for when the actual system colors are unavailable.
var SystemColors = map[string]Color{
	"scrollBar": ParseHex(
		"#C8C8C8",
	),
	"background": ParseHex(
		"#000080",
	), // Desktop background (navy)
	"activeCaption": ParseHex(
		"#99B4D1",
	),
	"inactiveCaption": ParseHex(
		"#BFCDDB",
	),
	"menu": ParseHex(
		"#F0F0F0",
	),
	"window": White,
	"windowFrame": ParseHex(
		"#646464",
	),
	"menuText":    Black,
	"windowText":  Black,
	"captionText": Black,
	"activeBorder": ParseHex(
		"#B4B4B4",
	),
	"inactiveBorder": ParseHex(
		"#F4F7FC",
	),
	"appWorkspace": ParseHex(
		"#ABABAB",
	),
	"highlight": ParseHex(
		"#0078D7",
	),
	"highlightText": White,
	"btnFace": ParseHex(
		"#F0F0F0",
	),
	"btnShadow": ParseHex(
		"#A0A0A0",
	),
	"grayText": ParseHex(
		"#6D6D6D",
	),
	"btnText": Black,
	"inactiveCaptionText": ParseHex(
		"#434E54",
	),
	"btnHighlight": White,
	"3dDkShadow": ParseHex(
		"#696969",
	),
	"3dLight": ParseHex(
		"#E3E3E3",
	),
	"infoText": Black,
	"infoBk": ParseHex(
		"#FFFFE1",
	),
	"hotLight": ParseHex(
		"#0066CC",
	),
	"gradientActiveCaption": ParseHex(
		"#B9D1EA",
	),
	"gradientInactiveCaption": ParseHex(
		"#D7E4F2",
	),
	"menuHighlight": ParseHex(
		"#0078D7",
	),
	"menuBar": ParseHex(
		"#F0F0F0",
	),
}

// FromRgbColor converts a DrawingML RgbColor element to a DrawingMLColor.
func FromRgbColor(
	rgb *drawingml.RgbColor,
) *DrawingMLColor {
	if rgb == nil {
		return nil
	}

	color := &DrawingMLColor{
		Type:  ColorTypeSRGB,
		Value: rgb.Value(),
	}

	// Parse transformations from child elements.
	color.Transformations = parseColorTransformations(
		rgb,
	)

	return color
}

// FromSchemeColor converts a DrawingML SchemeColor element to a DrawingMLColor.
func FromSchemeColor(
	sc *drawingml.SchemeColor,
) *DrawingMLColor {
	if sc == nil {
		return nil
	}

	color := &DrawingMLColor{
		Type:  ColorTypeScheme,
		Value: string(sc.Value()),
	}

	// Parse transformations from child elements.
	color.Transformations = parseColorTransformations(
		sc,
	)

	return color
}

// FromHslColor converts a DrawingML HslColor element to a DrawingMLColor.
func FromHslColor(
	hsl *drawingml.HslColor,
) *DrawingMLColor {
	if hsl == nil {
		return nil
	}

	return &DrawingMLColor{
		Type:       ColorTypeHSL,
		Hue:        hsl.Hue(),
		Saturation: hsl.Saturation(),
		Luminance:  hsl.Luminance(),
	}
}

// FromSystemColor converts a DrawingML SystemColor element to a DrawingMLColor.
func FromSystemColor(
	sys *drawingml.SystemColor,
) *DrawingMLColor {
	if sys == nil {
		return nil
	}

	color := &DrawingMLColor{
		Type:      ColorTypeSystem,
		Value:     string(sys.Value()),
		LastColor: sys.LastColor(),
	}

	// Parse transformations from child elements.
	color.Transformations = parseColorTransformations(
		sys,
	)

	return color
}

// FromPresetColor converts a DrawingML PresetColor element to a DrawingMLColor.
func FromPresetColor(
	preset *drawingml.PresetColor,
) *DrawingMLColor {
	if preset == nil {
		return nil
	}

	return &DrawingMLColor{
		Type:  ColorTypePreset,
		Value: string(preset.Value()),
	}
}
