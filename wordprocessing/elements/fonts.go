//nolint:revive // file-length-limit: fonts.go contains all font-related types
package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// FontFamilyValue represents font family classification values.
type FontFamilyValue string

const (
	// FontFamilyAuto indicates automatic font family detection.
	FontFamilyAuto FontFamilyValue = "auto"
	// FontFamilyDecorative indicates decorative fonts.
	FontFamilyDecorative FontFamilyValue = "decorative"
	// FontFamilyModern indicates fixed-pitch fonts.
	FontFamilyModern FontFamilyValue = "modern"
	// FontFamilyRoman indicates proportional fonts with serifs.
	FontFamilyRoman FontFamilyValue = "roman"
	// FontFamilyScript indicates script/cursive fonts.
	FontFamilyScript FontFamilyValue = "script"
	// FontFamilySwiss indicates proportional fonts without serifs.
	FontFamilySwiss FontFamilyValue = "swiss"
)

// FontPitchValue represents font pitch values.
type FontPitchValue string

const (
	// FontPitchDefault indicates default pitch.
	FontPitchDefault FontPitchValue = "default"
	// FontPitchFixed indicates fixed-width characters.
	FontPitchFixed FontPitchValue = "fixed"
	// FontPitchVariable indicates variable-width characters.
	FontPitchVariable FontPitchValue = "variable"
)

// Fonts represents the root element for the font table part (w:fonts).
type Fonts struct {
	*openxml.PartRootElementBase
}

// NewFonts creates a new Fonts element.
func NewFonts() *Fonts {
	elem := openxml.NewPartRootElement(
		NamespaceWML,
		"fonts",
		PrefixW,
	)

	return &Fonts{PartRootElementBase: elem}
}

// Fonts returns an iterator over all Font elements.
func (f *Fonts) Fonts() iter.Seq[*Font] {
	return func(yield func(*Font) bool) {
		for child := range f.Children() {
			if child.LocalName() == "font" &&
				child.NamespaceURI() == NamespaceWML {
				var font *Font
				switch v := child.(type) {
				case *Font:
					font = v
				case *openxml.CompositeElementBase:
					font = &Font{CompositeElementBase: v}
				}
				if font != nil && !yield(font) {
					return
				}
			}
		}
	}
}

// GetFont returns the font with the given name, or nil if not found.
func (f *Fonts) GetFont(name string) *Font {
	for font := range f.Fonts() {
		if font.Name() == name {
			return font
		}
	}

	return nil
}

// AddFont adds a font to the fonts collection.
func (f *Fonts) AddFont(font *Font) {
	f.AppendChild(font)
}

// AddFontByName creates and adds a font with the given name.
func (f *Fonts) AddFontByName(name string) *Font {
	font := NewFont()
	font.SetName(name)
	f.AppendChild(font)

	return font
}

// RemoveFont removes a font from the fonts collection.
func (f *Fonts) RemoveFont(font *Font) bool {
	return f.RemoveChild(font)
}

// Clone creates a deep copy of this Fonts element.
func (f *Fonts) Clone() openxml.Element {
	return &Fonts{
		PartRootElementBase: f.PartRootElementBase.Clone().(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this Fonts element.
func (f *Fonts) CloneNode(
	deep bool,
) openxml.Element {
	return &Fonts{
		PartRootElementBase: f.PartRootElementBase.CloneNode(deep).(*openxml.PartRootElementBase),
	}
}

// Font represents a font definition element (w:font).
type Font struct {
	*openxml.CompositeElementBase
}

// NewFont creates a new Font element.
func NewFont() *Font {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"font",
		PrefixW,
	)

	return &Font{CompositeElementBase: elem}
}

// Name returns the font name.
func (fn *Font) Name() string {
	attr, found := fn.GetAttribute(
		"name",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the font name.
func (fn *Font) SetName(name string) {
	fn.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"name",
			PrefixW,
			name,
		),
	)
}

// Charset returns the font character set.
func (fn *Font) Charset() string {
	elem := fn.GetElement("charset", NamespaceWML)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCharset sets the font character set.
func (fn *Font) SetCharset(charset string) {
	if charset == "" {
		fn.removeElement("charset")

		return
	}
	elem := fn.getOrCreateElement("charset")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			charset,
		),
	)
}

// Family returns the font family.
func (fn *Font) Family() FontFamilyValue {
	elem := fn.GetElement("family", NamespaceWML)
	if elem == nil {
		return FontFamilyAuto
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return FontFamilyAuto
	}

	return FontFamilyValue(attr.Value())
}

// SetFamily sets the font family.
func (fn *Font) SetFamily(
	family FontFamilyValue,
) {
	if family == FontFamilyAuto {
		fn.removeElement("family")

		return
	}
	elem := fn.getOrCreateElement("family")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(family),
		),
	)
}

// Pitch returns the font pitch.
func (fn *Font) Pitch() FontPitchValue {
	elem := fn.GetElement("pitch", NamespaceWML)
	if elem == nil {
		return FontPitchDefault
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return FontPitchDefault
	}

	return FontPitchValue(attr.Value())
}

// SetPitch sets the font pitch.
func (fn *Font) SetPitch(pitch FontPitchValue) {
	if pitch == FontPitchDefault {
		fn.removeElement("pitch")

		return
	}
	elem := fn.getOrCreateElement("pitch")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(pitch),
		),
	)
}

// Panose1 returns the Panose-1 number for the font.
func (fn *Font) Panose1() string {
	elem := fn.GetElement("panose1", NamespaceWML)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPanose1 sets the Panose-1 number.
func (fn *Font) SetPanose1(panose string) {
	if panose == "" {
		fn.removeElement("panose1")

		return
	}
	elem := fn.getOrCreateElement("panose1")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			panose,
		),
	)
}

// AltName returns the alternate font name.
func (fn *Font) AltName() string {
	elem := fn.GetElement("altName", NamespaceWML)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAltName sets the alternate font name.
func (fn *Font) SetAltName(name string) {
	if name == "" {
		fn.removeElement("altName")

		return
	}
	elem := fn.getOrCreateElement("altName")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			name,
		),
	)
}

// EmbedRegular returns the embedded regular font, or nil if not present.
func (fn *Font) EmbedRegular() *EmbedFont {
	elem := fn.GetElement(
		"embedRegular",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if ef, ok := elem.(*EmbedFont); ok {
		return ef
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &EmbedFont{
			CompositeElementBase: comp,
			fontStyle:            "embedRegular",
		}
	}

	return nil
}

// GetOrCreateEmbedRegular returns the embedded regular font, creating if needed.
func (fn *Font) GetOrCreateEmbedRegular() *EmbedFont {
	ef := fn.EmbedRegular()
	if ef != nil {
		return ef
	}
	ef = NewEmbedFont("embedRegular")
	fn.AppendChild(ef)

	return ef
}

// EmbedBold returns the embedded bold font, or nil if not present.
func (fn *Font) EmbedBold() *EmbedFont {
	elem := fn.GetElement(
		"embedBold",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if ef, ok := elem.(*EmbedFont); ok {
		return ef
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &EmbedFont{
			CompositeElementBase: comp,
			fontStyle:            "embedBold",
		}
	}

	return nil
}

// GetOrCreateEmbedBold returns the embedded bold font, creating if needed.
func (fn *Font) GetOrCreateEmbedBold() *EmbedFont {
	ef := fn.EmbedBold()
	if ef != nil {
		return ef
	}
	ef = NewEmbedFont("embedBold")
	fn.AppendChild(ef)

	return ef
}

// EmbedItalic returns the embedded italic font, or nil if not present.
func (fn *Font) EmbedItalic() *EmbedFont {
	elem := fn.GetElement(
		"embedItalic",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if ef, ok := elem.(*EmbedFont); ok {
		return ef
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &EmbedFont{
			CompositeElementBase: comp,
			fontStyle:            "embedItalic",
		}
	}

	return nil
}

// GetOrCreateEmbedItalic returns the embedded italic font, creating if needed.
func (fn *Font) GetOrCreateEmbedItalic() *EmbedFont {
	ef := fn.EmbedItalic()
	if ef != nil {
		return ef
	}
	ef = NewEmbedFont("embedItalic")
	fn.AppendChild(ef)

	return ef
}

// EmbedBoldItalic returns the embedded bold italic font, or nil if not present.
func (fn *Font) EmbedBoldItalic() *EmbedFont {
	elem := fn.GetElement(
		"embedBoldItalic",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if ef, ok := elem.(*EmbedFont); ok {
		return ef
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &EmbedFont{
			CompositeElementBase: comp,
			fontStyle:            "embedBoldItalic",
		}
	}

	return nil
}

// GetOrCreateEmbedBoldItalic returns the embedded bold italic font, creating if needed.
func (fn *Font) GetOrCreateEmbedBoldItalic() *EmbedFont {
	ef := fn.EmbedBoldItalic()
	if ef != nil {
		return ef
	}
	ef = NewEmbedFont("embedBoldItalic")
	fn.AppendChild(ef)

	return ef
}

// Sig returns the font signature element, or nil if not present.
func (fn *Font) Sig() *FontSig {
	elem := fn.GetElement("sig", NamespaceWML)
	if elem == nil {
		return nil
	}
	if fs, ok := elem.(*FontSig); ok {
		return fs
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &FontSig{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSig returns the font signature, creating if needed.
func (fn *Font) GetOrCreateSig() *FontSig {
	fs := fn.Sig()
	if fs != nil {
		return fs
	}
	fs = NewFontSig()
	fn.AppendChild(fs)

	return fs
}

// Helper methods

func (fn *Font) getOrCreateElement(
	name string,
) openxml.Element {
	elem := fn.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	fn.AppendChild(newElem)

	return newElem
}

func (fn *Font) removeElement(name string) {
	elem := fn.GetElement(name, NamespaceWML)
	if elem != nil {
		fn.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this Font element.
func (fn *Font) Clone() openxml.Element {
	return &Font{
		CompositeElementBase: fn.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Font element.
func (fn *Font) CloneNode(
	deep bool,
) openxml.Element {
	return &Font{
		CompositeElementBase: fn.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// EmbedFont represents an embedded font reference (w:embedRegular, w:embedBold, etc.).
type EmbedFont struct {
	*openxml.CompositeElementBase
	fontStyle string
}

// NewEmbedFont creates a new EmbedFont element with the specified style.
func NewEmbedFont(style string) *EmbedFont {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		style,
		PrefixW,
	)

	return &EmbedFont{
		CompositeElementBase: elem,
		fontStyle:            style,
	}
}

// RelationshipId returns the relationship ID for the embedded font.
func (ef *EmbedFont) RelationshipId() string {
	// The r:id attribute is in the relationships namespace
	attr, found := ef.GetAttribute(
		"id",
		openxml.NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID for the embedded font.
func (ef *EmbedFont) SetRelationshipId(
	id string,
) {
	ef.SetAttribute(
		openxml.NewAttribute(
			openxml.NamespaceRelationships,
			"id",
			"r",
			id,
		),
	)
}

// FontKey returns the font key (GUID) for the embedded font.
func (ef *EmbedFont) FontKey() string {
	attr, found := ef.GetAttribute(
		"fontKey",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetFontKey sets the font key (GUID) for the embedded font.
func (ef *EmbedFont) SetFontKey(key string) {
	if key == "" {
		ef.RemoveAttribute(
			"fontKey",
			NamespaceWML,
		)

		return
	}
	ef.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"fontKey",
			PrefixW,
			key,
		),
	)
}

// SubsetFontLicensing returns the subset font licensing value.
func (ef *EmbedFont) SubsetFontLicensing() string {
	attr, found := ef.GetAttribute(
		"subsetted",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// Clone creates a deep copy of this EmbedFont element.
func (ef *EmbedFont) Clone() openxml.Element {
	return &EmbedFont{
		CompositeElementBase: ef.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
		fontStyle:            ef.fontStyle,
	}
}

// CloneNode creates a copy of this EmbedFont element.
func (ef *EmbedFont) CloneNode(
	deep bool,
) openxml.Element {
	return &EmbedFont{
		CompositeElementBase: ef.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
		fontStyle:            ef.fontStyle,
	}
}

// FontSig represents font signature information (w:sig).
type FontSig struct {
	*openxml.CompositeElementBase
}

// NewFontSig creates a new FontSig element.
func NewFontSig() *FontSig {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"sig",
		PrefixW,
	)

	return &FontSig{CompositeElementBase: elem}
}

// Usb0 returns the first Unicode subset bitfield.
func (fs *FontSig) Usb0() string {
	attr, found := fs.GetAttribute(
		"usb0",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetUsb0 sets the first Unicode subset bitfield.
func (fs *FontSig) SetUsb0(val string) {
	fs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"usb0",
			PrefixW,
			val,
		),
	)
}

// Usb1 returns the second Unicode subset bitfield.
func (fs *FontSig) Usb1() string {
	attr, found := fs.GetAttribute(
		"usb1",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetUsb1 sets the second Unicode subset bitfield.
func (fs *FontSig) SetUsb1(val string) {
	fs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"usb1",
			PrefixW,
			val,
		),
	)
}

// Usb2 returns the third Unicode subset bitfield.
func (fs *FontSig) Usb2() string {
	attr, found := fs.GetAttribute(
		"usb2",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetUsb2 sets the third Unicode subset bitfield.
func (fs *FontSig) SetUsb2(val string) {
	fs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"usb2",
			PrefixW,
			val,
		),
	)
}

// Usb3 returns the fourth Unicode subset bitfield.
func (fs *FontSig) Usb3() string {
	attr, found := fs.GetAttribute(
		"usb3",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetUsb3 sets the fourth Unicode subset bitfield.
func (fs *FontSig) SetUsb3(val string) {
	fs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"usb3",
			PrefixW,
			val,
		),
	)
}

// Csb0 returns the first code page bitfield.
func (fs *FontSig) Csb0() string {
	attr, found := fs.GetAttribute(
		"csb0",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCsb0 sets the first code page bitfield.
func (fs *FontSig) SetCsb0(val string) {
	fs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"csb0",
			PrefixW,
			val,
		),
	)
}

// Csb1 returns the second code page bitfield.
func (fs *FontSig) Csb1() string {
	attr, found := fs.GetAttribute(
		"csb1",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCsb1 sets the second code page bitfield.
func (fs *FontSig) SetCsb1(val string) {
	fs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"csb1",
			PrefixW,
			val,
		),
	)
}

// Clone creates a deep copy of this FontSig element.
func (fs *FontSig) Clone() openxml.Element {
	return &FontSig{
		CompositeElementBase: fs.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this FontSig element.
func (fs *FontSig) CloneNode(
	deep bool,
) openxml.Element {
	return &FontSig{
		CompositeElementBase: fs.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}
