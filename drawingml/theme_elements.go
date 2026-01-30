package drawingml

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// ThemeElements represents the theme elements (a:themeElements).
type ThemeElements struct {
	*openxml.CompositeElementBase
}

// NewThemeElements creates a new theme elements container.
func NewThemeElements() *ThemeElements {
	return &ThemeElements{
		CompositeElementBase: openxml.NewCompositeElement(NamespaceMain, "themeElements", PrefixMain),
	}
}

// ColorScheme returns the color scheme element.
func (te *ThemeElements) ColorScheme() *ColorScheme {
	elem := te.GetElement("clrScheme", NamespaceMain)
	if elem == nil {
		return nil
	}
	if cs, ok := elem.(*ColorScheme); ok {
		return cs
	}

	return nil
}

// GetOrCreateColorScheme returns or creates the color scheme.
func (te *ThemeElements) GetOrCreateColorScheme(name string) *ColorScheme {
	cs := te.ColorScheme()
	if cs != nil {
		return cs
	}
	cs = NewColorScheme(name)
	te.AppendChild(cs)

	return cs
}

// FontScheme returns the font scheme element.
func (te *ThemeElements) FontScheme() *FontScheme {
	elem := te.GetElement("fontScheme", NamespaceMain)
	if elem == nil {
		return nil
	}
	if fs, ok := elem.(*FontScheme); ok {
		return fs
	}

	return nil
}

// GetOrCreateFontScheme returns or creates the font scheme.
func (te *ThemeElements) GetOrCreateFontScheme(name string) *FontScheme {
	fs := te.FontScheme()
	if fs != nil {
		return fs
	}
	fs = NewFontScheme(name)
	te.AppendChild(fs)

	return fs
}

// FormatScheme returns the format scheme element.
func (te *ThemeElements) FormatScheme() *FormatScheme {
	elem := te.GetElement("fmtScheme", NamespaceMain)
	if elem == nil {
		return nil
	}
	if fs, ok := elem.(*FormatScheme); ok {
		return fs
	}

	return nil
}

// GetOrCreateFormatScheme returns or creates the format scheme.
func (te *ThemeElements) GetOrCreateFormatScheme(name string) *FormatScheme {
	fs := te.FormatScheme()
	if fs != nil {
		return fs
	}
	fs = NewFormatScheme(name)
	te.AppendChild(fs)

	return fs
}

// ===========================================================================
// ColorScheme (a:clrScheme)
// ===========================================================================

// ColorScheme represents a color scheme (a:clrScheme).
type ColorScheme struct {
	*openxml.CompositeElementBase
}

// NewColorScheme creates a new color scheme with the given name.
func NewColorScheme(name string) *ColorScheme {
	elem := openxml.NewCompositeElement(NamespaceMain, "clrScheme", PrefixMain)
	cs := &ColorScheme{CompositeElementBase: elem}
	if name != "" {
		cs.SetAttribute(openxml.NewAttribute("", "name", "", name))
	}

	return cs
}

// SetColor sets a color for a specific scheme element.
// name is one of: dk1, lt1, dk2, lt2, accent1-6, hlink, folHlink.
func (cs *ColorScheme) SetColor(name string, color openxml.Element) {
	// Remove existing
	if existing := cs.GetElement(name, NamespaceMain); existing != nil {
		cs.RemoveChild(existing)
	}

	// Create container element (e.g. <a:dk1>)
	container := openxml.NewCompositeElement(NamespaceMain, name, PrefixMain)
	container.AppendChild(color)
	cs.AppendChild(container)
}

// SetSrgbColor sets a solid RGB color for a scheme element.
func (cs *ColorScheme) SetSrgbColor(name, hexColor string) {
	cs.SetColor(name, NewRgbColor(hexColor))
}

// SetSysColor sets a system color for a scheme element.
func (cs *ColorScheme) SetSysColor(name string, sysColor SystemColorValue, lastColor string) {
	cs.SetColor(name, NewSystemColorWithLastColor(sysColor, lastColor))
}

// ResolveColor resolves a scheme color name to its concrete hex color string.
func (cs *ColorScheme) ResolveColor(colorName SchemeColorValue) string {
	// Find the container element for the specific scheme color (e.g., <a:dk1>)
	schemeColorContainer := cs.GetElement(string(colorName), NamespaceMain)
	if schemeColorContainer == nil {
		return ""
	}

	// Need to cast schemeColorContainer to a CompositeElementBase to call GetElement on its children
	containerBase := wrapCompositeElement(schemeColorContainer)
	if containerBase == nil {
		return ""
	}

	// Prioritize SRGB color
	if srgbElem := containerBase.GetElement("srgbClr", NamespaceMain); srgbElem != nil {
		if rc, ok := srgbElem.(*RgbColor); ok {
			return rc.Value()
		}
		if comp := wrapCompositeElement(srgbElem); comp != nil {
			return (&RgbColor{CompositeElementBase: comp}).Value()
		}
	}

	// Try System Color
	if sysElem := containerBase.GetElement("sysClr", NamespaceMain); sysElem != nil {
		if sc, ok := sysElem.(*SystemColor); ok {
			return sc.LastColor()
		}
		if comp := wrapCompositeElement(sysElem); comp != nil {
			return (&SystemColor{CompositeElementBase: comp}).LastColor()
		}
	}

	// Try Preset Color
	if prstElem := containerBase.GetElement("prstClr", NamespaceMain); prstElem != nil {
		if pc, ok := prstElem.(*PresetColor); ok {
			return GetPresetColorHex(pc.Value())
		}
		if comp := wrapCompositeElement(prstElem); comp != nil {
			return GetPresetColorHex((&PresetColor{CompositeElementBase: comp}).Value())
		}
	}

	return ""
}

// ===========================================================================
// FontScheme (a:fontScheme)
// ===========================================================================

// FontScheme represents a font scheme (a:fontScheme).
type FontScheme struct {
	*openxml.CompositeElementBase
}

// NewFontScheme creates a new font scheme with the given name.
func NewFontScheme(name string) *FontScheme {
	elem := openxml.NewCompositeElement(NamespaceMain, "fontScheme", PrefixMain)
	fs := &FontScheme{CompositeElementBase: elem}
	if name != "" {
		fs.SetAttribute(openxml.NewAttribute("", "name", "", name))
	}

	return fs
}

// MajorFont returns the major font collection.
func (fs *FontScheme) MajorFont() *FontCollection {
	elem := fs.GetElement("majorFont", NamespaceMain)
	if elem == nil {
		return nil
	}
	if fc, ok := elem.(*FontCollection); ok {
		return fc
	}

	return nil
}

// GetOrCreateMajorFont returns or creates the major font collection.
func (fs *FontScheme) GetOrCreateMajorFont() *FontCollection {
	fc := fs.MajorFont()
	if fc != nil {
		return fc
	}
	fc = NewFontCollection("majorFont")
	fs.PrependChild(fc)

	return fc
}

// MinorFont returns the minor font collection.
func (fs *FontScheme) MinorFont() *FontCollection {
	elem := fs.GetElement("minorFont", NamespaceMain)
	if elem == nil {
		return nil
	}
	if fc, ok := elem.(*FontCollection); ok {
		return fc
	}

	return nil
}

// GetOrCreateMinorFont returns or creates the minor font collection.
func (fs *FontScheme) GetOrCreateMinorFont() *FontCollection {
	fc := fs.MinorFont()
	if fc != nil {
		return fc
	}
	fc = NewFontCollection("minorFont")
	// Insert after majorFont if present
	if major := fs.MajorFont(); major != nil {
		fs.InsertAfter(fc, major)
	} else {
		fs.PrependChild(fc)
	}

	return fc
}

// ===========================================================================
// FontCollection (a:majorFont, a:minorFont)
// ===========================================================================

// FontCollection represents a font collection (CT_FontCollection).
type FontCollection struct {
	*openxml.CompositeElementBase
}

// NewFontCollection creates a new font collection element.
func NewFontCollection(localName string) *FontCollection {
	elem := openxml.NewCompositeElement(NamespaceMain, localName, PrefixMain)

	return &FontCollection{CompositeElementBase: elem}
}

// SetLatinFont sets the Latin font.
func (fc *FontCollection) SetLatinFont(typeface string) {
	fc.setFont("latin", typeface)
}

// SetEaFont sets the East Asian font.
func (fc *FontCollection) SetEaFont(typeface string) {
	fc.setFont("ea", typeface)
}

// SetCsFont sets the Complex Script font.
func (fc *FontCollection) SetCsFont(typeface string) {
	fc.setFont("cs", typeface)
}

func (fc *FontCollection) setFont(name, typeface string) {
	if existing := fc.GetElement(name, NamespaceMain); existing != nil {
		fc.RemoveChild(existing)
	}
	font := NewTextFont(name)
	font.SetTypeface(typeface)
	fc.AppendChild(font)
}

// ===========================================================================
// FormatScheme (a:fmtScheme)
// ===========================================================================

// FormatScheme represents a format scheme (a:fmtScheme).
type FormatScheme struct {
	*openxml.CompositeElementBase
}

// NewFormatScheme creates a new format scheme element.
func NewFormatScheme(name string) *FormatScheme {
	elem := openxml.NewCompositeElement(NamespaceMain, "fmtScheme", PrefixMain)
	fs := &FormatScheme{CompositeElementBase: elem}
	if name != "" {
		fs.SetAttribute(openxml.NewAttribute("", "name", "", name))
	}

	return fs
}

// FillStyleList returns the fill style list.
func (fs *FormatScheme) FillStyleList() *FillStyleList {
	elem := fs.GetElement("fillStyleLst", NamespaceMain)
	if elem == nil {
		return nil
	}
	if fsl, ok := elem.(*FillStyleList); ok {
		return fsl
	}

	return nil
}

// GetOrCreateFillStyleList returns or creates the fill style list.
func (fs *FormatScheme) GetOrCreateFillStyleList() *FillStyleList {
	fsl := fs.FillStyleList()
	if fsl != nil {
		return fsl
	}
	fsl = NewFillStyleList()
	fs.AppendChild(fsl)

	return fsl
}

// LineStyleList returns the line style list.
func (fs *FormatScheme) LineStyleList() *LineStyleList {
	elem := fs.GetElement("lnStyleLst", NamespaceMain)
	if elem == nil {
		return nil
	}
	if lsl, ok := elem.(*LineStyleList); ok {
		return lsl
	}

	return nil
}

// GetOrCreateLineStyleList returns or creates the line style list.
func (fs *FormatScheme) GetOrCreateLineStyleList() *LineStyleList {
	lsl := fs.LineStyleList()
	if lsl != nil {
		return lsl
	}
	lsl = NewLineStyleList()
	fs.AppendChild(lsl)

	return lsl
}

// EffectStyleList returns the effect style list.
func (fs *FormatScheme) EffectStyleList() *EffectStyleList {
	elem := fs.GetElement("effectStyleLst", NamespaceMain)
	if elem == nil {
		return nil
	}
	if esl, ok := elem.(*EffectStyleList); ok {
		return esl
	}

	return nil
}

// GetOrCreateEffectStyleList returns or creates the effect style list.
func (fs *FormatScheme) GetOrCreateEffectStyleList() *EffectStyleList {
	esl := fs.EffectStyleList()
	if esl != nil {
		return esl
	}
	esl = NewEffectStyleList()
	fs.AppendChild(esl)

	return esl
}

// BackgroundFillStyleList returns the background fill style list.
func (fs *FormatScheme) BackgroundFillStyleList() *FillStyleList {
	elem := fs.GetElement("bgFillStyleLst", NamespaceMain)
	if elem == nil {
		return nil
	}
	if fsl, ok := elem.(*FillStyleList); ok {
		return fsl
	}

	return nil
}

// GetOrCreateBackgroundFillStyleList returns or creates the background fill style list.
func (fs *FormatScheme) GetOrCreateBackgroundFillStyleList() *FillStyleList {
	fsl := fs.BackgroundFillStyleList()
	if fsl != nil {
		return fsl
	}
	fsl = NewFillStyleList()
	// Override tag name to bgFillStyleLst
	fsl.CompositeElementBase = openxml.NewCompositeElement(NamespaceMain, "bgFillStyleLst", PrefixMain)
	fs.AppendChild(fsl)

	return fsl
}

// ===========================================================================
// Style Lists (a:fillStyleLst, a:lnStyleLst, a:effectStyleLst)
// ===========================================================================

// FillStyleList represents a list of fill styles (a:fillStyleLst).
type FillStyleList struct {
	*openxml.CompositeElementBase
}

// NewFillStyleList creates a new fill style list.
func NewFillStyleList() *FillStyleList {
	elem := openxml.NewCompositeElement(NamespaceMain, "fillStyleLst", PrefixMain)

	return &FillStyleList{CompositeElementBase: elem}
}

// AddSolidFill adds a solid color fill to the list.
func (fsl *FillStyleList) AddSolidFill(hexColor string) *SolidFill {
	sf := NewSolidFillWithRgb(hexColor)
	fsl.AppendChild(sf)

	return sf
}

// LineStyleList represents a list of line styles (a:lnStyleLst).
type LineStyleList struct {
	*openxml.CompositeElementBase
}

// NewLineStyleList creates a new line style list.
func NewLineStyleList() *LineStyleList {
	elem := openxml.NewCompositeElement(NamespaceMain, "lnStyleLst", PrefixMain)

	return &LineStyleList{CompositeElementBase: elem}
}

// AddLine adds a line properties element to the list.
func (lsl *LineStyleList) AddLine() *LineProperties {
	lp := NewLineProperties()
	lsl.AppendChild(lp)

	return lp
}

// EffectStyleList represents a list of effect styles (a:effectStyleLst).
type EffectStyleList struct {
	*openxml.CompositeElementBase
}

// NewEffectStyleList creates a new effect style list.
func NewEffectStyleList() *EffectStyleList {
	elem := openxml.NewCompositeElement(NamespaceMain, "effectStyleLst", PrefixMain)

	return &EffectStyleList{CompositeElementBase: elem}
}

// ===========================================================================
// ExtraColorSchemeList (a:extraClrSchemeLst)
// ===========================================================================

// ExtraColorSchemeList represents a list of extra color schemes (a:extraClrSchemeLst).
type ExtraColorSchemeList struct {
	*openxml.CompositeElementBase
}

// NewExtraColorSchemeList creates a new extra color scheme list.
func NewExtraColorSchemeList() *ExtraColorSchemeList {
	elem := openxml.NewCompositeElement(NamespaceMain, "extraClrSchemeLst", PrefixMain)

	return &ExtraColorSchemeList{CompositeElementBase: elem}
}

// ===========================================================================
// ExtraColorScheme (a:extraClrScheme)
// ===========================================================================

// ExtraColorScheme represents an extra color scheme (a:extraClrScheme).
type ExtraColorScheme struct {
	*openxml.CompositeElementBase
}

// NewExtraColorScheme creates a new extra color scheme.
func NewExtraColorScheme() *ExtraColorScheme {
	elem := openxml.NewCompositeElement(NamespaceMain, "extraClrScheme", PrefixMain)

	return &ExtraColorScheme{CompositeElementBase: elem}
}
