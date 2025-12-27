package drawing

import (
	"github.com/connerohnesorge/goffice/drawingml"
)

// ThemePalette represents a document's theme color palette.
// It provides colors for the 12 standard theme color slots plus
// additional color mappings.
type ThemePalette struct {
	// Core theme colors (12 slots).
	Dark1             Color // dk1 - typically black or dark text
	Light1            Color // lt1 - typically white or light background
	Dark2             Color // dk2 - secondary dark color
	Light2            Color // lt2 - secondary light color
	Accent1           Color // accent1 - first accent color
	Accent2           Color // accent2 - second accent color
	Accent3           Color // accent3 - third accent color
	Accent4           Color // accent4 - fourth accent color
	Accent5           Color // accent5 - fifth accent color
	Accent6           Color // accent6 - sixth accent color
	Hyperlink         Color // hlink - hyperlink color
	FollowedHyperlink Color // folHlink - followed hyperlink color

	// Additional color aliases.
	aliases map[drawingml.SchemeColorValue]drawingml.SchemeColorValue
}

// NewThemePalette creates a new theme palette with default Office colors.
func NewThemePalette() *ThemePalette {
	return &ThemePalette{
		Dark1:  Black,
		Light1: White,
		Dark2: ParseHex(
			"#44546A",
		), // Office dark gray-blue
		Light2: ParseHex(
			"#E7E6E6",
		), // Office light gray
		Accent1: ParseHex(
			"#4472C4",
		), // Office blue
		Accent2: ParseHex(
			"#ED7D31",
		), // Office orange
		Accent3: ParseHex(
			"#A5A5A5",
		), // Office gray
		Accent4: ParseHex(
			"#FFC000",
		), // Office gold
		Accent5: ParseHex(
			"#5B9BD5",
		), // Office light blue
		Accent6: ParseHex(
			"#70AD47",
		), // Office green
		Hyperlink: ParseHex(
			"#0563C1",
		), // Office hyperlink blue
		FollowedHyperlink: ParseHex(
			"#954F72",
		), // Office followed link purple
		aliases: defaultColorAliases(),
	}
}

// NewThemePaletteFromColors creates a theme palette from individual color values.
func NewThemePaletteFromColors(
	dk1, lt1, dk2, lt2,
	accent1, accent2, accent3, accent4, accent5, accent6,
	hlink, folHlink Color,
) *ThemePalette {
	return &ThemePalette{
		Dark1:             dk1,
		Light1:            lt1,
		Dark2:             dk2,
		Light2:            lt2,
		Accent1:           accent1,
		Accent2:           accent2,
		Accent3:           accent3,
		Accent4:           accent4,
		Accent5:           accent5,
		Accent6:           accent6,
		Hyperlink:         hlink,
		FollowedHyperlink: folHlink,
		aliases:           defaultColorAliases(),
	}
}

// defaultColorAliases returns the default color alias mappings.
func defaultColorAliases() map[drawingml.SchemeColorValue]drawingml.SchemeColorValue {
	return map[drawingml.SchemeColorValue]drawingml.SchemeColorValue{
		// Background and text aliases.
		drawingml.SchemeColorBackground1: drawingml.SchemeColorLight1,
		drawingml.SchemeColorText1:       drawingml.SchemeColorDark1,
		drawingml.SchemeColorBackground2: drawingml.SchemeColorLight2,
		drawingml.SchemeColorText2:       drawingml.SchemeColorDark2,
	}
}

// GetColor returns the color for a scheme color value.
// Returns false if the color is not found in the palette.
func (p *ThemePalette) GetColor(
	value drawingml.SchemeColorValue,
) (Color, bool) {
	if p == nil {
		return Black, false
	}

	// Resolve aliases first.
	resolved := value
	if alias, ok := p.aliases[value]; ok {
		resolved = alias
	}

	switch resolved {
	case drawingml.SchemeColorDark1:
		return p.Dark1, true
	case drawingml.SchemeColorLight1:
		return p.Light1, true
	case drawingml.SchemeColorDark2:
		return p.Dark2, true
	case drawingml.SchemeColorLight2:
		return p.Light2, true
	case drawingml.SchemeColorAccent1:
		return p.Accent1, true
	case drawingml.SchemeColorAccent2:
		return p.Accent2, true
	case drawingml.SchemeColorAccent3:
		return p.Accent3, true
	case drawingml.SchemeColorAccent4:
		return p.Accent4, true
	case drawingml.SchemeColorAccent5:
		return p.Accent5, true
	case drawingml.SchemeColorAccent6:
		return p.Accent6, true
	case drawingml.SchemeColorHyperlink:
		return p.Hyperlink, true
	case drawingml.SchemeColorFollowedHyperlink:
		return p.FollowedHyperlink, true
	case drawingml.SchemeColorPhColor:
		// Placeholder color - typically transparent or inherited.
		return Transparent, true
	case drawingml.SchemeColorBackground1:
		return p.Light1, true
	case drawingml.SchemeColorText1:
		return p.Dark1, true
	case drawingml.SchemeColorBackground2:
		return p.Light2, true
	case drawingml.SchemeColorText2:
		return p.Dark2, true
	default:
		return Black, false
	}
}

// GetColorOrDefault returns the color for a scheme color value,
// or black if not found.
func (p *ThemePalette) GetColorOrDefault(
	value drawingml.SchemeColorValue,
) Color {
	color, ok := p.GetColor(value)
	if !ok {
		return Black
	}

	return color
}

// SetColor sets the color for a scheme color value.
func (p *ThemePalette) SetColor(
	value drawingml.SchemeColorValue,
	color Color,
) {
	if p == nil {
		return
	}

	switch value {
	case drawingml.SchemeColorDark1:
		p.Dark1 = color
	case drawingml.SchemeColorLight1:
		p.Light1 = color
	case drawingml.SchemeColorDark2:
		p.Dark2 = color
	case drawingml.SchemeColorLight2:
		p.Light2 = color
	case drawingml.SchemeColorAccent1:
		p.Accent1 = color
	case drawingml.SchemeColorAccent2:
		p.Accent2 = color
	case drawingml.SchemeColorAccent3:
		p.Accent3 = color
	case drawingml.SchemeColorAccent4:
		p.Accent4 = color
	case drawingml.SchemeColorAccent5:
		p.Accent5 = color
	case drawingml.SchemeColorAccent6:
		p.Accent6 = color
	case drawingml.SchemeColorHyperlink:
		p.Hyperlink = color
	case drawingml.SchemeColorFollowedHyperlink:
		p.FollowedHyperlink = color
	case drawingml.SchemeColorPhColor:
		// Cannot set placeholder color
	case drawingml.SchemeColorBackground1:
		p.Light1 = color
	case drawingml.SchemeColorText1:
		p.Dark1 = color
	case drawingml.SchemeColorBackground2:
		p.Light2 = color
	case drawingml.SchemeColorText2:
		p.Dark2 = color
	}
}

// Clone creates a deep copy of the theme palette.
func (p *ThemePalette) Clone() *ThemePalette {
	if p == nil {
		return nil
	}

	clone := &ThemePalette{
		Dark1:             p.Dark1,
		Light1:            p.Light1,
		Dark2:             p.Dark2,
		Light2:            p.Light2,
		Accent1:           p.Accent1,
		Accent2:           p.Accent2,
		Accent3:           p.Accent3,
		Accent4:           p.Accent4,
		Accent5:           p.Accent5,
		Accent6:           p.Accent6,
		Hyperlink:         p.Hyperlink,
		FollowedHyperlink: p.FollowedHyperlink,
		aliases: make(
			map[drawingml.SchemeColorValue]drawingml.SchemeColorValue,
		),
	}

	for k, v := range p.aliases {
		clone.aliases[k] = v
	}

	return clone
}

// DefaultThemePalette is the default Office theme color palette.
var DefaultThemePalette = NewThemePalette()

// OfficeThemes provides predefined theme color palettes for common Office themes.
var OfficeThemes = map[string]*ThemePalette{
	"Office": NewThemePalette(),
	"Grayscale": NewThemePaletteFromColors(
		Black,               // dk1
		White,               // lt1
		ParseHex("#636363"), // dk2
		ParseHex("#D0CECE"), // lt2
		ParseHex("#A6A6A6"), // accent1
		ParseHex("#595959"), // accent2
		ParseHex("#7F7F7F"), // accent3
		ParseHex("#BFBFBF"), // accent4
		ParseHex("#404040"), // accent5
		ParseHex("#D8D8D8"), // accent6
		ParseHex("#0563C1"), // hlink
		ParseHex("#954F72"), // folHlink
	),
	"Blue": NewThemePaletteFromColors(
		ParseHex("#1E3264"), // dk1
		White,               // lt1
		ParseHex("#2A5599"), // dk2
		ParseHex("#DEEAF6"), // lt2
		ParseHex("#2F5496"), // accent1
		ParseHex("#ED7D31"), // accent2
		ParseHex("#A5A5A5"), // accent3
		ParseHex("#FFC000"), // accent4
		ParseHex("#5B9BD5"), // accent5
		ParseHex("#70AD47"), // accent6
		ParseHex("#2E75B5"), // hlink
		ParseHex("#954F72"), // folHlink
	),
	"Green": NewThemePaletteFromColors(
		ParseHex("#385623"), // dk1
		White,               // lt1
		ParseHex("#548235"), // dk2
		ParseHex("#E2EFD9"), // lt2
		ParseHex("#70AD47"), // accent1
		ParseHex("#5B9BD5"), // accent2
		ParseHex("#A5A5A5"), // accent3
		ParseHex("#FFC000"), // accent4
		ParseHex("#4472C4"), // accent5
		ParseHex("#ED7D31"), // accent6
		ParseHex("#548235"), // hlink
		ParseHex("#954F72"), // folHlink
	),
	"Red": NewThemePaletteFromColors(
		ParseHex("#833C0C"), // dk1
		White,               // lt1
		ParseHex("#C55A11"), // dk2
		ParseHex("#FBE4D5"), // lt2
		ParseHex("#C00000"), // accent1
		ParseHex("#5B9BD5"), // accent2
		ParseHex("#A5A5A5"), // accent3
		ParseHex("#FFC000"), // accent4
		ParseHex("#4472C4"), // accent5
		ParseHex("#70AD47"), // accent6
		ParseHex("#C00000"), // hlink
		ParseHex("#954F72"), // folHlink
	),
	"Purple": NewThemePaletteFromColors(
		ParseHex("#482D8B"), // dk1
		White,               // lt1
		ParseHex("#6B4DA3"), // dk2
		ParseHex("#E6E0EC"), // lt2
		ParseHex("#7030A0"), // accent1
		ParseHex("#ED7D31"), // accent2
		ParseHex("#A5A5A5"), // accent3
		ParseHex("#FFC000"), // accent4
		ParseHex("#5B9BD5"), // accent5
		ParseHex("#70AD47"), // accent6
		ParseHex("#7030A0"), // hlink
		ParseHex("#954F72"), // folHlink
	),
}

// ThemePaletteFromHexColors creates a theme palette from hex color strings.
// This is a convenience function for creating palettes from color definitions.
func ThemePaletteFromHexColors(
	dk1, lt1, dk2, lt2,
	accent1, accent2, accent3, accent4, accent5, accent6,
	hlink, folHlink string,
) *ThemePalette {
	return NewThemePaletteFromColors(
		ParseHex(dk1),
		ParseHex(lt1),
		ParseHex(dk2),
		ParseHex(lt2),
		ParseHex(accent1),
		ParseHex(accent2),
		ParseHex(accent3),
		ParseHex(accent4),
		ParseHex(accent5),
		ParseHex(accent6),
		ParseHex(hlink),
		ParseHex(folHlink),
	)
}

// AccentColors returns all accent colors as a slice.
func (p *ThemePalette) AccentColors() []Color {
	return []Color{
		p.Accent1,
		p.Accent2,
		p.Accent3,
		p.Accent4,
		p.Accent5,
		p.Accent6,
	}
}

// AccentColor returns an accent color by index (1-6).
// Returns Accent1 for invalid indices.
func (p *ThemePalette) AccentColor(
	index int,
) Color {
	switch index {
	case 1:
		return p.Accent1
	case 2:
		return p.Accent2
	case 3:
		return p.Accent3
	case 4:
		return p.Accent4
	case 5:
		return p.Accent5
	case 6:
		return p.Accent6
	default:
		return p.Accent1
	}
}
