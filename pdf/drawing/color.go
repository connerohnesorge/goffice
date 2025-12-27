package drawing

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Color represents an RGBA color with components in the range [0, 1].
// This is the internal representation used for PDF color operations.
type Color struct {
	R, G, B, A float64
}

// NewColor creates a new color with the specified RGBA values.
// All values should be in the range [0, 1].
func NewColor(r, g, b, a float64) Color {
	return Color{
		R: clamp01(r),
		G: clamp01(g),
		B: clamp01(b),
		A: clamp01(a),
	}
}

// NewRGB creates a new opaque color with the specified RGB values.
// Values should be in the range [0, 1].
func NewRGB(r, g, b float64) Color {
	return NewColor(r, g, b, 1.0)
}

// NewRGB8 creates a new opaque color from 8-bit RGB values (0-255).
func NewRGB8(r, g, b uint8) Color {
	return NewRGB(
		float64(r)/255.0,
		float64(g)/255.0,
		float64(b)/255.0,
	)
}

// NewRGBA8 creates a new color from 8-bit RGBA values (0-255).
func NewRGBA8(r, g, b, a uint8) Color {
	return NewColor(
		float64(r)/255.0,
		float64(g)/255.0,
		float64(b)/255.0,
		float64(a)/255.0,
	)
}

// NewGray creates a new opaque grayscale color.
// The value should be in the range [0, 1] where 0 is black and 1 is white.
func NewGray(gray float64) Color {
	return NewRGB(gray, gray, gray)
}

// NewGray8 creates a new opaque grayscale color from an 8-bit value (0-255).
func NewGray8(gray uint8) Color {
	g := float64(gray) / 255.0

	return NewRGB(g, g, g)
}

// ParseHex parses a hex color string (#RRGGBB or #RGB format).
// Returns black if the format is invalid.
func ParseHex(hex string) Color {
	color, _ := ParseHexError(hex)

	return color
}

// ParseHexError parses a hex color string and returns an error if invalid.
// Supported formats: #RRGGBB, RRGGBB, #RGB, RGB
func ParseHexError(hex string) (Color, error) {
	// Remove leading #
	s := strings.TrimPrefix(hex, "#")

	// Handle short format (#RGB -> #RRGGBB)
	if len(s) == 3 {
		s = string(
			[]byte{
				s[0],
				s[0],
				s[1],
				s[1],
				s[2],
				s[2],
			},
		)
	}

	if len(s) != 6 {
		return Color{}, fmt.Errorf(
			"invalid hex color format: %s",
			hex,
		)
	}

	r, err := strconv.ParseUint(s[0:2], 16, 8)
	if err != nil {
		return Color{}, fmt.Errorf(
			"invalid red component in hex color: %s",
			hex,
		)
	}

	g, err := strconv.ParseUint(s[2:4], 16, 8)
	if err != nil {
		return Color{}, fmt.Errorf(
			"invalid green component in hex color: %s",
			hex,
		)
	}

	b, err := strconv.ParseUint(s[4:6], 16, 8)
	if err != nil {
		return Color{}, fmt.Errorf(
			"invalid blue component in hex color: %s",
			hex,
		)
	}

	return NewRGB8(
		uint8(r),
		uint8(g),
		uint8(b),
	), nil
}

// ToHex returns the color as a hex string in #RRGGBB format.
func (c Color) ToHex() string {
	r := uint8(math.Round(c.R * 255))
	g := uint8(math.Round(c.G * 255))
	b := uint8(math.Round(c.B * 255))

	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// ToRGB8 returns the color as 8-bit RGB values (0-255).
func (c Color) ToRGB8() (r, g, b uint8) {
	return uint8(math.Round(c.R * 255)),
		uint8(math.Round(c.G * 255)),
		uint8(math.Round(c.B * 255))
}

// ToRGBA8 returns the color as 8-bit RGBA values (0-255).
func (c Color) ToRGBA8() (r, g, b, a uint8) {
	return uint8(math.Round(c.R * 255)),
		uint8(math.Round(c.G * 255)),
		uint8(math.Round(c.B * 255)),
		uint8(math.Round(c.A * 255))
}

// IsOpaque returns true if the color is fully opaque (alpha = 1).
func (c Color) IsOpaque() bool {
	return c.A >= 1.0
}

// IsTransparent returns true if the color is fully transparent (alpha = 0).
func (c Color) IsTransparent() bool {
	return c.A <= 0.0
}

// WithAlpha returns a copy of the color with the specified alpha value.
func (c Color) WithAlpha(a float64) Color {
	return NewColor(c.R, c.G, c.B, a)
}

// Blend blends this color with another color using alpha compositing.
// The amount parameter (0-1) controls the blend: 0 = this color, 1 = other color.
func (c Color) Blend(
	other Color,
	amount float64,
) Color {
	amount = clamp01(amount)

	return NewColor(
		c.R+(other.R-c.R)*amount,
		c.G+(other.G-c.G)*amount,
		c.B+(other.B-c.B)*amount,
		c.A+(other.A-c.A)*amount,
	)
}

// Lighten returns a lighter version of the color.
// The amount (0-1) specifies how much to lighten.
func (c Color) Lighten(amount float64) Color {
	return c.Blend(White, amount)
}

// Darken returns a darker version of the color.
// The amount (0-1) specifies how much to darken.
func (c Color) Darken(amount float64) Color {
	return c.Blend(Black, amount)
}

// CMYKColor represents a CMYK color with components in the range [0, 1].
type CMYKColor struct {
	C, M, Y, K float64
}

// NewCMYK creates a new CMYK color.
// All values should be in the range [0, 1].
func NewCMYK(c, m, y, k float64) CMYKColor {
	return CMYKColor{
		C: clamp01(c),
		M: clamp01(m),
		Y: clamp01(y),
		K: clamp01(k),
	}
}

// ToRGB converts a CMYK color to an RGB color.
func (c CMYKColor) ToRGB() Color {
	// CMYK to RGB conversion
	r := (1 - c.C) * (1 - c.K)
	g := (1 - c.M) * (1 - c.K)
	b := (1 - c.Y) * (1 - c.K)

	return NewRGB(r, g, b)
}

// RGBToCMYK converts an RGB color to a CMYK color.
func RGBToCMYK(color Color) CMYKColor {
	// Find K (black)
	k := 1 - math.Max(
		color.R,
		math.Max(color.G, color.B),
	)
	if k >= 1.0 {
		return NewCMYK(0, 0, 0, 1)
	}

	// Calculate CMY
	c := (1 - color.R - k) / (1 - k)
	m := (1 - color.G - k) / (1 - k)
	y := (1 - color.B - k) / (1 - k)

	return NewCMYK(c, m, y, k)
}

// HSLColor represents a color in HSL (Hue, Saturation, Lightness) space.
type HSLColor struct {
	H, S, L float64 // H: 0-360, S: 0-1, L: 0-1
}

// NewHSL creates a new HSL color.
// H is in degrees (0-360), S and L are in the range [0, 1].
func NewHSL(h, s, l float64) HSLColor {
	// Normalize hue to 0-360
	for h < 0 {
		h += 360
	}
	for h >= 360 {
		h -= 360
	}

	return HSLColor{
		H: h,
		S: clamp01(s),
		L: clamp01(l),
	}
}

// ToRGB converts an HSL color to an RGB color.
func (c HSLColor) ToRGB() Color {
	if c.S == 0 {
		// Achromatic (gray)
		return NewRGB(c.L, c.L, c.L)
	}

	var q float64
	if c.L < 0.5 {
		q = c.L * (1 + c.S)
	} else {
		q = c.L + c.S - c.L*c.S
	}
	p := 2*c.L - q

	hue := c.H / 360.0

	r := hueToRGB(p, q, hue+1.0/3.0)
	g := hueToRGB(p, q, hue)
	b := hueToRGB(p, q, hue-1.0/3.0)

	return NewRGB(r, g, b)
}

// hueToRGB is a helper function for HSL to RGB conversion.
func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 0.5 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}

	return p
}

// RGBToHSL converts an RGB color to an HSL color.
func RGBToHSL(color Color) HSLColor {
	maxC := math.Max(
		color.R,
		math.Max(color.G, color.B),
	)
	minC := math.Min(
		color.R,
		math.Min(color.G, color.B),
	)
	l := (maxC + minC) / 2

	if maxC == minC {
		return NewHSL(0, 0, l)
	}

	d := maxC - minC
	var s float64
	if l > 0.5 {
		s = d / (2 - maxC - minC)
	} else {
		s = d / (maxC + minC)
	}

	var h float64
	switch maxC {
	case color.R:
		h = (color.G - color.B) / d
		if color.G < color.B {
			h += 6
		}
	case color.G:
		h = (color.B-color.R)/d + 2
	case color.B:
		h = (color.R-color.G)/d + 4
	}
	h *= 60

	return NewHSL(h, s, l)
}

// PDF Color Operators

// SetFillRGB returns the PDF operator to set the fill color in RGB color space.
// This outputs the "rg" operator.
func (c Color) SetFillRGB() string {
	return fmt.Sprintf(
		"%s %s %s rg",
		formatFloat(
			c.R,
		),
		formatFloat(c.G),
		formatFloat(c.B),
	)
}

// SetStrokeRGB returns the PDF operator to set the stroke color in RGB color space.
// This outputs the "RG" operator.
func (c Color) SetStrokeRGB() string {
	return fmt.Sprintf(
		"%s %s %s RG",
		formatFloat(
			c.R,
		),
		formatFloat(c.G),
		formatFloat(c.B),
	)
}

// SetFillGray returns the PDF operator to set the fill color in grayscale.
// This outputs the "g" operator.
func SetFillGray(gray float64) string {
	return fmt.Sprintf(
		"%s g",
		formatFloat(clamp01(gray)),
	)
}

// SetStrokeGray returns the PDF operator to set the stroke color in grayscale.
// This outputs the "G" operator.
func SetStrokeGray(gray float64) string {
	return fmt.Sprintf(
		"%s G",
		formatFloat(clamp01(gray)),
	)
}

// SetFillCMYK returns the PDF operator to set the fill color in CMYK color space.
// This outputs the "k" operator.
func (c CMYKColor) SetFillCMYK() string {
	return fmt.Sprintf(
		"%s %s %s %s k",
		formatFloat(
			c.C,
		),
		formatFloat(c.M),
		formatFloat(c.Y),
		formatFloat(c.K),
	)
}

// SetStrokeCMYK returns the PDF operator to set the stroke color in CMYK color space.
// This outputs the "K" operator.
func (c CMYKColor) SetStrokeCMYK() string {
	return fmt.Sprintf(
		"%s %s %s %s K",
		formatFloat(
			c.C,
		),
		formatFloat(c.M),
		formatFloat(c.Y),
		formatFloat(c.K),
	)
}

// SetColorSpace returns the PDF operator to set the color space.
// Use "DeviceRGB", "DeviceCMYK", or "DeviceGray" for device color spaces.
func SetColorSpace(
	name string,
	forStroke bool,
) string {
	if forStroke {
		return fmt.Sprintf("/%s CS", name)
	}

	return fmt.Sprintf("/%s cs", name)
}

// SetColor returns the PDF operator to set the color in the current color space.
// This outputs the "sc" (fill) or "SC" (stroke) operator.
func SetColor(
	components []float64,
	forStroke bool,
) string {
	var sb strings.Builder
	for i, comp := range components {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(formatFloat(comp))
	}
	if forStroke {
		sb.WriteString(" SC")
	} else {
		sb.WriteString(" sc")
	}

	return sb.String()
}

// clamp01 clamps a value to the range [0, 1].
func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}

	return v
}

// Named colors - common web colors
var (
	// Basic colors
	Black = NewRGB(0, 0, 0)
	White = NewRGB(1, 1, 1)
	Red   = NewRGB(1, 0, 0)
	Green = NewRGB(
		0,
		0.5020,
		0,
	) // Web green (#008000)
	Blue        = NewRGB(0, 0, 1)
	Yellow      = NewRGB(1, 1, 0)
	Cyan        = NewRGB(0, 1, 1)
	Magenta     = NewRGB(1, 0, 1)
	Transparent = NewColor(0, 0, 0, 0)

	// Extended colors
	AliceBlue            = ParseHex("#F0F8FF")
	AntiqueWhite         = ParseHex("#FAEBD7")
	Aqua                 = ParseHex("#00FFFF")
	Aquamarine           = ParseHex("#7FFFD4")
	Azure                = ParseHex("#F0FFFF")
	Beige                = ParseHex("#F5F5DC")
	Bisque               = ParseHex("#FFE4C4")
	BlanchedAlmond       = ParseHex("#FFEBCD")
	BlueViolet           = ParseHex("#8A2BE2")
	Brown                = ParseHex("#A52A2A")
	BurlyWood            = ParseHex("#DEB887")
	CadetBlue            = ParseHex("#5F9EA0")
	Chartreuse           = ParseHex("#7FFF00")
	Chocolate            = ParseHex("#D2691E")
	Coral                = ParseHex("#FF7F50")
	CornflowerBlue       = ParseHex("#6495ED")
	Cornsilk             = ParseHex("#FFF8DC")
	Crimson              = ParseHex("#DC143C")
	DarkBlue             = ParseHex("#00008B")
	DarkCyan             = ParseHex("#008B8B")
	DarkGoldenrod        = ParseHex("#B8860B")
	DarkGray             = ParseHex("#A9A9A9")
	DarkGreen            = ParseHex("#006400")
	DarkKhaki            = ParseHex("#BDB76B")
	DarkMagenta          = ParseHex("#8B008B")
	DarkOliveGreen       = ParseHex("#556B2F")
	DarkOrange           = ParseHex("#FF8C00")
	DarkOrchid           = ParseHex("#9932CC")
	DarkRed              = ParseHex("#8B0000")
	DarkSalmon           = ParseHex("#E9967A")
	DarkSeaGreen         = ParseHex("#8FBC8F")
	DarkSlateBlue        = ParseHex("#483D8B")
	DarkSlateGray        = ParseHex("#2F4F4F")
	DarkTurquoise        = ParseHex("#00CED1")
	DarkViolet           = ParseHex("#9400D3")
	DeepPink             = ParseHex("#FF1493")
	DeepSkyBlue          = ParseHex("#00BFFF")
	DimGray              = ParseHex("#696969")
	DodgerBlue           = ParseHex("#1E90FF")
	Firebrick            = ParseHex("#B22222")
	FloralWhite          = ParseHex("#FFFAF0")
	ForestGreen          = ParseHex("#228B22")
	Fuchsia              = ParseHex("#FF00FF")
	Gainsboro            = ParseHex("#DCDCDC")
	GhostWhite           = ParseHex("#F8F8FF")
	Gold                 = ParseHex("#FFD700")
	Goldenrod            = ParseHex("#DAA520")
	Gray                 = ParseHex("#808080")
	GreenYellow          = ParseHex("#ADFF2F")
	Honeydew             = ParseHex("#F0FFF0")
	HotPink              = ParseHex("#FF69B4")
	IndianRed            = ParseHex("#CD5C5C")
	Indigo               = ParseHex("#4B0082")
	Ivory                = ParseHex("#FFFFF0")
	Khaki                = ParseHex("#F0E68C")
	Lavender             = ParseHex("#E6E6FA")
	LavenderBlush        = ParseHex("#FFF0F5")
	LawnGreen            = ParseHex("#7CFC00")
	LemonChiffon         = ParseHex("#FFFACD")
	LightBlue            = ParseHex("#ADD8E6")
	LightCoral           = ParseHex("#F08080")
	LightCyan            = ParseHex("#E0FFFF")
	LightGoldenrodYellow = ParseHex("#FAFAD2")
	LightGray            = ParseHex("#D3D3D3")
	LightGreen           = ParseHex("#90EE90")
	LightPink            = ParseHex("#FFB6C1")
	LightSalmon          = ParseHex("#FFA07A")
	LightSeaGreen        = ParseHex("#20B2AA")
	LightSkyBlue         = ParseHex("#87CEFA")
	LightSlateGray       = ParseHex("#778899")
	LightSteelBlue       = ParseHex("#B0C4DE")
	LightYellow          = ParseHex("#FFFFE0")
	Lime                 = ParseHex("#00FF00")
	LimeGreen            = ParseHex("#32CD32")
	Linen                = ParseHex("#FAF0E6")
	Maroon               = ParseHex("#800000")
	MediumAquamarine     = ParseHex("#66CDAA")
	MediumBlue           = ParseHex("#0000CD")
	MediumOrchid         = ParseHex("#BA55D3")
	MediumPurple         = ParseHex("#9370DB")
	MediumSeaGreen       = ParseHex("#3CB371")
	MediumSlateBlue      = ParseHex("#7B68EE")
	MediumSpringGreen    = ParseHex("#00FA9A")
	MediumTurquoise      = ParseHex("#48D1CC")
	MediumVioletRed      = ParseHex("#C71585")
	MidnightBlue         = ParseHex("#191970")
	MintCream            = ParseHex("#F5FFFA")
	MistyRose            = ParseHex("#FFE4E1")
	Moccasin             = ParseHex("#FFE4B5")
	NavajoWhite          = ParseHex("#FFDEAD")
	Navy                 = ParseHex("#000080")
	OldLace              = ParseHex("#FDF5E6")
	Olive                = ParseHex("#808000")
	OliveDrab            = ParseHex("#6B8E23")
	Orange               = ParseHex("#FFA500")
	OrangeRed            = ParseHex("#FF4500")
	Orchid               = ParseHex("#DA70D6")
	PaleGoldenrod        = ParseHex("#EEE8AA")
	PaleGreen            = ParseHex("#98FB98")
	PaleTurquoise        = ParseHex("#AFEEEE")
	PaleVioletRed        = ParseHex("#DB7093")
	PapayaWhip           = ParseHex("#FFEFD5")
	PeachPuff            = ParseHex("#FFDAB9")
	Peru                 = ParseHex("#CD853F")
	Pink                 = ParseHex("#FFC0CB")
	Plum                 = ParseHex("#DDA0DD")
	PowderBlue           = ParseHex("#B0E0E6")
	Purple               = ParseHex("#800080")
	RebeccaPurple        = ParseHex("#663399")
	RosyBrown            = ParseHex("#BC8F8F")
	RoyalBlue            = ParseHex("#4169E1")
	SaddleBrown          = ParseHex("#8B4513")
	Salmon               = ParseHex("#FA8072")
	SandyBrown           = ParseHex("#F4A460")
	SeaGreen             = ParseHex("#2E8B57")
	SeaShell             = ParseHex("#FFF5EE")
	Sienna               = ParseHex("#A0522D")
	Silver               = ParseHex("#C0C0C0")
	SkyBlue              = ParseHex("#87CEEB")
	SlateBlue            = ParseHex("#6A5ACD")
	SlateGray            = ParseHex("#708090")
	Snow                 = ParseHex("#FFFAFA")
	SpringGreen          = ParseHex("#00FF7F")
	SteelBlue            = ParseHex("#4682B4")
	Tan                  = ParseHex("#D2B48C")
	Teal                 = ParseHex("#008080")
	Thistle              = ParseHex("#D8BFD8")
	Tomato               = ParseHex("#FF6347")
	Turquoise            = ParseHex("#40E0D0")
	Violet               = ParseHex("#EE82EE")
	Wheat                = ParseHex("#F5DEB3")
	WhiteSmoke           = ParseHex("#F5F5F5")
	YellowGreen          = ParseHex("#9ACD32")
)

// namedColors maps color names to Color values.
var namedColors = map[string]Color{
	"black":                Black,
	"white":                White,
	"red":                  Red,
	"green":                Green,
	"blue":                 Blue,
	"yellow":               Yellow,
	"cyan":                 Cyan,
	"magenta":              Magenta,
	"transparent":          Transparent,
	"aliceblue":            AliceBlue,
	"antiquewhite":         AntiqueWhite,
	"aqua":                 Aqua,
	"aquamarine":           Aquamarine,
	"azure":                Azure,
	"beige":                Beige,
	"bisque":               Bisque,
	"blanchedalmond":       BlanchedAlmond,
	"blueviolet":           BlueViolet,
	"brown":                Brown,
	"burlywood":            BurlyWood,
	"cadetblue":            CadetBlue,
	"chartreuse":           Chartreuse,
	"chocolate":            Chocolate,
	"coral":                Coral,
	"cornflowerblue":       CornflowerBlue,
	"cornsilk":             Cornsilk,
	"crimson":              Crimson,
	"darkblue":             DarkBlue,
	"darkcyan":             DarkCyan,
	"darkgoldenrod":        DarkGoldenrod,
	"darkgray":             DarkGray,
	"darkgrey":             DarkGray,
	"darkgreen":            DarkGreen,
	"darkkhaki":            DarkKhaki,
	"darkmagenta":          DarkMagenta,
	"darkolivegreen":       DarkOliveGreen,
	"darkorange":           DarkOrange,
	"darkorchid":           DarkOrchid,
	"darkred":              DarkRed,
	"darksalmon":           DarkSalmon,
	"darkseagreen":         DarkSeaGreen,
	"darkslateblue":        DarkSlateBlue,
	"darkslategray":        DarkSlateGray,
	"darkslategrey":        DarkSlateGray,
	"darkturquoise":        DarkTurquoise,
	"darkviolet":           DarkViolet,
	"deeppink":             DeepPink,
	"deepskyblue":          DeepSkyBlue,
	"dimgray":              DimGray,
	"dimgrey":              DimGray,
	"dodgerblue":           DodgerBlue,
	"firebrick":            Firebrick,
	"floralwhite":          FloralWhite,
	"forestgreen":          ForestGreen,
	"fuchsia":              Fuchsia,
	"gainsboro":            Gainsboro,
	"ghostwhite":           GhostWhite,
	"gold":                 Gold,
	"goldenrod":            Goldenrod,
	"gray":                 Gray,
	"grey":                 Gray,
	"greenyellow":          GreenYellow,
	"honeydew":             Honeydew,
	"hotpink":              HotPink,
	"indianred":            IndianRed,
	"indigo":               Indigo,
	"ivory":                Ivory,
	"khaki":                Khaki,
	"lavender":             Lavender,
	"lavenderblush":        LavenderBlush,
	"lawngreen":            LawnGreen,
	"lemonchiffon":         LemonChiffon,
	"lightblue":            LightBlue,
	"lightcoral":           LightCoral,
	"lightcyan":            LightCyan,
	"lightgoldenrodyellow": LightGoldenrodYellow,
	"lightgray":            LightGray,
	"lightgrey":            LightGray,
	"lightgreen":           LightGreen,
	"lightpink":            LightPink,
	"lightsalmon":          LightSalmon,
	"lightseagreen":        LightSeaGreen,
	"lightskyblue":         LightSkyBlue,
	"lightslategray":       LightSlateGray,
	"lightslategrey":       LightSlateGray,
	"lightsteelblue":       LightSteelBlue,
	"lightyellow":          LightYellow,
	"lime":                 Lime,
	"limegreen":            LimeGreen,
	"linen":                Linen,
	"maroon":               Maroon,
	"mediumaquamarine":     MediumAquamarine,
	"mediumblue":           MediumBlue,
	"mediumorchid":         MediumOrchid,
	"mediumpurple":         MediumPurple,
	"mediumseagreen":       MediumSeaGreen,
	"mediumslateblue":      MediumSlateBlue,
	"mediumspringgreen":    MediumSpringGreen,
	"mediumturquoise":      MediumTurquoise,
	"mediumvioletred":      MediumVioletRed,
	"midnightblue":         MidnightBlue,
	"mintcream":            MintCream,
	"mistyrose":            MistyRose,
	"moccasin":             Moccasin,
	"navajowhite":          NavajoWhite,
	"navy":                 Navy,
	"oldlace":              OldLace,
	"olive":                Olive,
	"olivedrab":            OliveDrab,
	"orange":               Orange,
	"orangered":            OrangeRed,
	"orchid":               Orchid,
	"palegoldenrod":        PaleGoldenrod,
	"palegreen":            PaleGreen,
	"paleturquoise":        PaleTurquoise,
	"palevioletred":        PaleVioletRed,
	"papayawhip":           PapayaWhip,
	"peachpuff":            PeachPuff,
	"peru":                 Peru,
	"pink":                 Pink,
	"plum":                 Plum,
	"powderblue":           PowderBlue,
	"purple":               Purple,
	"rebeccapurple":        RebeccaPurple,
	"rosybrown":            RosyBrown,
	"royalblue":            RoyalBlue,
	"saddlebrown":          SaddleBrown,
	"salmon":               Salmon,
	"sandybrown":           SandyBrown,
	"seagreen":             SeaGreen,
	"seashell":             SeaShell,
	"sienna":               Sienna,
	"silver":               Silver,
	"skyblue":              SkyBlue,
	"slateblue":            SlateBlue,
	"slategray":            SlateGray,
	"slategrey":            SlateGray,
	"snow":                 Snow,
	"springgreen":          SpringGreen,
	"steelblue":            SteelBlue,
	"tan":                  Tan,
	"teal":                 Teal,
	"thistle":              Thistle,
	"tomato":               Tomato,
	"turquoise":            Turquoise,
	"violet":               Violet,
	"wheat":                Wheat,
	"whitesmoke":           WhiteSmoke,
	"yellowgreen":          YellowGreen,
}

// NamedColor returns a color by name.
// Returns black if the name is not found.
// Color names are case-insensitive.
func NamedColor(name string) Color {
	color, found := NamedColorLookup(name)
	if !found {
		return Black
	}

	return color
}

// NamedColorLookup returns a color by name and whether it was found.
// Color names are case-insensitive.
func NamedColorLookup(name string) (Color, bool) {
	c, found := namedColors[strings.ToLower(name)]

	return c, found
}

// ParseColor parses a color from various formats:
// - Named colors: "red", "blue", etc.
// - Hex colors: "#FF0000", "#F00", "FF0000"
// - RGB: "rgb(255, 0, 0)"
// Returns black if parsing fails.
func ParseColor(s string) Color {
	color, _ := ParseColorError(s)

	return color
}

// ParseColorError parses a color and returns an error if parsing fails.
func ParseColorError(s string) (Color, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return Black, errors.New(
			"empty color string",
		)
	}

	// Try named color first
	if c, found := NamedColorLookup(s); found {
		return c, nil
	}

	// Try hex color
	if strings.HasPrefix(s, "#") ||
		isHexString(s) {
		return ParseHexError(s)
	}

	// Try rgb() format
	if strings.HasPrefix(
		strings.ToLower(s),
		"rgb(",
	) {
		return parseRGBFunction(s)
	}

	// Try rgba() format
	if strings.HasPrefix(
		strings.ToLower(s),
		"rgba(",
	) {
		return parseRGBAFunction(s)
	}

	return Black, errors.New(
		"unrecognized color format",
	)
}

// isHexString checks if a string looks like a hex color without the # prefix.
func isHexString(s string) bool {
	if len(s) != 3 && len(s) != 6 {
		return false
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}

	return true
}

// parseRGBFunction parses an "rgb(r, g, b)" format string.
func parseRGBFunction(s string) (Color, error) {
	s = strings.TrimPrefix(
		strings.ToLower(s),
		"rgb(",
	)
	s = strings.TrimSuffix(s, ")")
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		return Black, fmt.Errorf(
			"invalid rgb() format: expected 3 components",
		)
	}

	r, err := parseColorComponent(
		strings.TrimSpace(parts[0]),
	)
	if err != nil {
		return Black, err
	}
	g, err := parseColorComponent(
		strings.TrimSpace(parts[1]),
	)
	if err != nil {
		return Black, err
	}
	b, err := parseColorComponent(
		strings.TrimSpace(parts[2]),
	)
	if err != nil {
		return Black, err
	}

	return NewRGB(r, g, b), nil
}

// parseRGBAFunction parses an "rgba(r, g, b, a)" format string.
func parseRGBAFunction(s string) (Color, error) {
	s = strings.TrimPrefix(
		strings.ToLower(s),
		"rgba(",
	)
	s = strings.TrimSuffix(s, ")")
	parts := strings.Split(s, ",")
	if len(parts) != 4 {
		return Black, fmt.Errorf(
			"invalid rgba() format: expected 4 components",
		)
	}

	r, err := parseColorComponent(
		strings.TrimSpace(parts[0]),
	)
	if err != nil {
		return Black, err
	}
	g, err := parseColorComponent(
		strings.TrimSpace(parts[1]),
	)
	if err != nil {
		return Black, err
	}
	b, err := parseColorComponent(
		strings.TrimSpace(parts[2]),
	)
	if err != nil {
		return Black, err
	}
	a, err := parseColorComponent(
		strings.TrimSpace(parts[3]),
	)
	if err != nil {
		return Black, err
	}

	return NewColor(r, g, b, a), nil
}

// parseColorComponent parses a color component which can be:
// - An integer (0-255)
// - A percentage (0%-100%)
// - A float (0-1)
func parseColorComponent(
	s string,
) (float64, error) {
	if strings.HasSuffix(s, "%") {
		// Percentage format
		s = strings.TrimSuffix(s, "%")
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, fmt.Errorf(
				"invalid percentage: %s",
				s,
			)
		}

		return v / 100.0, nil
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf(
			"invalid number: %s",
			s,
		)
	}

	// If the value is > 1, assume it's in 0-255 range
	if v > 1 {
		return v / 255.0, nil
	}

	return v, nil
}

// String returns a string representation of the color.
func (c Color) String() string {
	if c.A < 1.0 {
		return fmt.Sprintf(
			"rgba(%.0f, %.0f, %.0f, %.2f)",
			c.R*255,
			c.G*255,
			c.B*255,
			c.A,
		)
	}

	return c.ToHex()
}
