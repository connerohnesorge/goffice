package elements

//revive:disable:file-length-limit many font properties
//revive:disable:max-public-structs many font types

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// FontScheme represents the font scheme type.
type FontScheme string

const (
	// FontSchemeNone indicates no font scheme.
	FontSchemeNone FontScheme = "none"
	// FontSchemeMajor indicates the major (heading) font scheme.
	FontSchemeMajor FontScheme = "major"
	// FontSchemeMinor indicates the minor (body) font scheme.
	FontSchemeMinor FontScheme = "minor"
)

// Fonts represents the fonts container element (x:fonts).
type Fonts struct {
	*openxml.CompositeElementBase
}

// NewFonts creates a new Fonts element.
func NewFonts() *Fonts {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"fonts",
		PrefixDefault,
	)

	return &Fonts{CompositeElementBase: elem}
}

// Count returns the count attribute value.
func (f *Fonts) Count() uint32 {
	attr, found := f.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetCount sets the count attribute.
func (f *Fonts) SetCount(count uint32) {
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// Fonts returns an iterator over all Font elements.
func (f *Fonts) Fonts() iter.Seq[*Font] {
	return func(yield func(*Font) bool) {
		for child := range f.Children() {
			if child.LocalName() != "font" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var font *Font
			if fn, ok := child.(*Font); ok {
				font = fn
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				font = &Font{CompositeElementBase: comp}
			}
			if font != nil && !yield(font) {
				return
			}
		}
	}
}

// GetFont returns the font at the given index, or nil if out of range.
func (f *Fonts) GetFont(index uint32) *Font {
	var i uint32
	for font := range f.Fonts() {
		if i == index {
			return font
		}
		i++
	}

	return nil
}

// AddFont adds a new font and returns it.
func (f *Fonts) AddFont() *Font {
	font := NewFont()
	f.AppendChild(font)
	f.SetCount(f.Count() + 1)

	return font
}

// ItemCount returns the actual number of Font children.
func (f *Fonts) ItemCount() int {
	count := 0
	for range f.Fonts() {
		count++
	}

	return count
}

// Clone creates a deep copy of this Fonts element.
func (f *Fonts) Clone() openxml.Element {
	cloned := f.CompositeElementBase.Clone()

	return &Fonts{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Fonts element.
func (f *Fonts) CloneNode(
	deep bool,
) openxml.Element {
	cloned := f.CompositeElementBase.CloneNode(
		deep,
	)

	return &Fonts{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Font represents a font element (x:font) in the stylesheet.
type Font struct {
	*openxml.CompositeElementBase
}

// NewFont creates a new Font element.
func NewFont() *Font {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"font",
		PrefixDefault,
	)

	return &Font{CompositeElementBase: elem}
}

// Bold returns the bold element, or nil if not present.
func (f *Font) Bold() *FontBold {
	elem := f.GetElement("b", NamespaceSML)
	if elem == nil {
		return nil
	}
	if b, ok := elem.(*FontBold); ok {
		return b
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontBold{LeafElementBase: leaf}
	}

	return nil
}

// IsBold returns whether the font is bold.
func (f *Font) IsBold() bool {
	b := f.Bold()
	if b == nil {
		return false
	}

	return b.Value()
}

// SetBold sets whether the font is bold.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (f *Font) SetBold(bold bool) {
	b := f.Bold()
	if bold {
		if b == nil {
			b = NewFontBold()
			f.AppendChild(b)
		}
		b.SetValue(true)
	} else if b != nil {
		f.RemoveChild(b)
	}
}

// Italic returns the italic element, or nil if not present.
func (f *Font) Italic() *FontItalic {
	elem := f.GetElement("i", NamespaceSML)
	if elem == nil {
		return nil
	}
	if i, ok := elem.(*FontItalic); ok {
		return i
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontItalic{LeafElementBase: leaf}
	}

	return nil
}

// IsItalic returns whether the font is italic.
func (f *Font) IsItalic() bool {
	i := f.Italic()
	if i == nil {
		return false
	}

	return i.Value()
}

// SetItalic sets whether the font is italic.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (f *Font) SetItalic(italic bool) {
	i := f.Italic()
	if italic {
		if i == nil {
			i = NewFontItalic()
			f.AppendChild(i)
		}
		i.SetValue(true)
	} else if i != nil {
		f.RemoveChild(i)
	}
}

// Strike returns the strikethrough element, or nil if not present.
func (f *Font) Strike() *FontStrike {
	elem := f.GetElement("strike", NamespaceSML)
	if elem == nil {
		return nil
	}
	if s, ok := elem.(*FontStrike); ok {
		return s
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontStrike{LeafElementBase: leaf}
	}

	return nil
}

// IsStrikethrough returns whether the font has strikethrough.
func (f *Font) IsStrikethrough() bool {
	s := f.Strike()
	if s == nil {
		return false
	}

	return s.Value()
}

// SetStrikethrough sets whether the font has strikethrough.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (f *Font) SetStrikethrough(strike bool) {
	s := f.Strike()
	if strike {
		if s == nil {
			s = NewFontStrike()
			f.AppendChild(s)
		}
		s.SetValue(true)
	} else if s != nil {
		f.RemoveChild(s)
	}
}

// Underline returns the underline element, or nil if not present.
func (f *Font) Underline() *FontUnderline {
	elem := f.GetElement("u", NamespaceSML)
	if elem == nil {
		return nil
	}
	if u, ok := elem.(*FontUnderline); ok {
		return u
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontUnderline{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// UnderlineStyle returns the underline style.
func (f *Font) UnderlineStyle() UnderlineStyle {
	u := f.Underline()
	if u == nil {
		return UnderlineStyleNone
	}

	return u.Style()
}

// SetUnderline sets the underline style.
func (f *Font) SetUnderline(
	style UnderlineStyle,
) {
	u := f.Underline()
	if style == UnderlineStyleNone ||
		style == "" {
		if u != nil {
			f.RemoveChild(u)
		}

		return
	}
	if u == nil {
		u = NewFontUnderline()
		f.AppendChild(u)
	}
	u.SetStyle(style)
}

// Condense returns the condense element, or nil if not present.
func (f *Font) Condense() *FontCondense {
	elem := f.GetElement("condense", NamespaceSML)
	if elem == nil {
		return nil
	}
	if c, ok := elem.(*FontCondense); ok {
		return c
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontCondense{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// IsCondensed returns whether the font is condensed.
func (f *Font) IsCondensed() bool {
	c := f.Condense()
	if c == nil {
		return false
	}

	return c.Value()
}

// SetCondensed sets whether the font is condensed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (f *Font) SetCondensed(condense bool) {
	c := f.Condense()
	if condense {
		if c == nil {
			c = NewFontCondense()
			f.AppendChild(c)
		}
		c.SetValue(true)
	} else if c != nil {
		f.RemoveChild(c)
	}
}

// Extend returns the extend element, or nil if not present.
func (f *Font) Extend() *FontExtend {
	elem := f.GetElement("extend", NamespaceSML)
	if elem == nil {
		return nil
	}
	if e, ok := elem.(*FontExtend); ok {
		return e
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontExtend{LeafElementBase: leaf}
	}

	return nil
}

// IsExtended returns whether the font is extended.
func (f *Font) IsExtended() bool {
	e := f.Extend()
	if e == nil {
		return false
	}

	return e.Value()
}

// SetExtended sets whether the font is extended.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (f *Font) SetExtended(extend bool) {
	e := f.Extend()
	if extend {
		if e == nil {
			e = NewFontExtend()
			f.AppendChild(e)
		}
		e.SetValue(true)
	} else if e != nil {
		f.RemoveChild(e)
	}
}

// Outline returns the outline element, or nil if not present.
func (f *Font) Outline() *FontOutline {
	elem := f.GetElement("outline", NamespaceSML)
	if elem == nil {
		return nil
	}
	if o, ok := elem.(*FontOutline); ok {
		return o
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontOutline{LeafElementBase: leaf}
	}

	return nil
}

// IsOutline returns whether the font has outline effect.
func (f *Font) IsOutline() bool {
	o := f.Outline()
	if o == nil {
		return false
	}

	return o.Value()
}

// SetOutline sets whether the font has outline effect.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (f *Font) SetOutline(outline bool) {
	o := f.Outline()
	if outline {
		if o == nil {
			o = NewFontOutline()
			f.AppendChild(o)
		}
		o.SetValue(true)
	} else if o != nil {
		f.RemoveChild(o)
	}
}

// Shadow returns the shadow element, or nil if not present.
func (f *Font) Shadow() *FontShadow {
	elem := f.GetElement("shadow", NamespaceSML)
	if elem == nil {
		return nil
	}
	if s, ok := elem.(*FontShadow); ok {
		return s
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontShadow{LeafElementBase: leaf}
	}

	return nil
}

// IsShadow returns whether the font has shadow effect.
func (f *Font) IsShadow() bool {
	s := f.Shadow()
	if s == nil {
		return false
	}

	return s.Value()
}

// SetShadow sets whether the font has shadow effect.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (f *Font) SetShadow(shadow bool) {
	s := f.Shadow()
	if shadow {
		if s == nil {
			s = NewFontShadow()
			f.AppendChild(s)
		}
		s.SetValue(true)
	} else if s != nil {
		f.RemoveChild(s)
	}
}

// Size returns the size element, or nil if not present.
func (f *Font) Size() *FontSize {
	elem := f.GetElement("sz", NamespaceSML)
	if elem == nil {
		return nil
	}
	if sz, ok := elem.(*FontSize); ok {
		return sz
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontSize{LeafElementBase: leaf}
	}

	return nil
}

// FontSizeValue returns the font size in points.
func (f *Font) FontSizeValue() float64 {
	sz := f.Size()
	if sz == nil {
		return 0
	}

	return sz.Val()
}

// SetFontSize sets the font size in points.
func (f *Font) SetFontSize(size float64) {
	sz := f.Size()
	if size == 0 {
		if sz != nil {
			f.RemoveChild(sz)
		}

		return
	}
	if sz == nil {
		sz = NewFontSize()
		f.AppendChild(sz)
	}
	sz.SetVal(size)
}

// Color returns the color element, or nil if not present.
func (f *Font) Color() *Color {
	elem := f.GetElement("color", NamespaceSML)
	if elem == nil {
		return nil
	}
	if c, ok := elem.(*Color); ok {
		return c
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Color{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateColor returns the color element, creating if needed.
func (f *Font) GetOrCreateColor() *Color {
	c := f.Color()
	if c != nil {
		return c
	}
	c = NewColor()
	f.AppendChild(c)

	return c
}

// Name returns the font name element, or nil if not present.
func (f *Font) Name() *FontName {
	elem := f.GetElement("name", NamespaceSML)
	if elem == nil {
		return nil
	}
	if n, ok := elem.(*FontName); ok {
		return n
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontName{LeafElementBase: leaf}
	}

	return nil
}

// FontNameValue returns the font name.
func (f *Font) FontNameValue() string {
	n := f.Name()
	if n == nil {
		return ""
	}

	return n.Val()
}

// SetFontName sets the font name.
func (f *Font) SetFontName(name string) {
	n := f.Name()
	if name == "" {
		if n != nil {
			f.RemoveChild(n)
		}

		return
	}
	if n == nil {
		n = NewFontName()
		f.AppendChild(n)
	}
	n.SetVal(name)
}

// Family returns the font family element, or nil if not present.
func (f *Font) Family() *FontFamily {
	elem := f.GetElement("family", NamespaceSML)
	if elem == nil {
		return nil
	}
	if fam, ok := elem.(*FontFamily); ok {
		return fam
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontFamily{LeafElementBase: leaf}
	}

	return nil
}

// FontFamilyValue returns the font family number.
func (f *Font) FontFamilyValue() uint32 {
	fam := f.Family()
	if fam == nil {
		return 0
	}

	return fam.Val()
}

// SetFontFamily sets the font family number.
func (f *Font) SetFontFamily(family uint32) {
	fam := f.Family()
	if family == 0 {
		if fam != nil {
			f.RemoveChild(fam)
		}

		return
	}
	if fam == nil {
		fam = NewFontFamily()
		f.AppendChild(fam)
	}
	fam.SetVal(family)
}

// Charset returns the charset element, or nil if not present.
func (f *Font) Charset() *FontCharset {
	elem := f.GetElement("charset", NamespaceSML)
	if elem == nil {
		return nil
	}
	if cs, ok := elem.(*FontCharset); ok {
		return cs
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontCharset{LeafElementBase: leaf}
	}

	return nil
}

// CharsetValue returns the charset value.
func (f *Font) CharsetValue() uint32 {
	cs := f.Charset()
	if cs == nil {
		return 0
	}

	return cs.Val()
}

// SetCharset sets the charset value.
func (f *Font) SetCharset(charset uint32) {
	cs := f.Charset()
	if charset == 0 {
		if cs != nil {
			f.RemoveChild(cs)
		}

		return
	}
	if cs == nil {
		cs = NewFontCharset()
		f.AppendChild(cs)
	}
	cs.SetVal(charset)
}

// Scheme returns the scheme element, or nil if not present.
func (f *Font) Scheme() *FontSchemeElement {
	elem := f.GetElement("scheme", NamespaceSML)
	if elem == nil {
		return nil
	}
	if s, ok := elem.(*FontSchemeElement); ok {
		return s
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontSchemeElement{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// FontSchemeValue returns the font scheme.
func (f *Font) FontSchemeValue() FontScheme {
	s := f.Scheme()
	if s == nil {
		return FontSchemeNone
	}

	return s.Val()
}

// SetFontScheme sets the font scheme.
func (f *Font) SetFontScheme(scheme FontScheme) {
	s := f.Scheme()
	if scheme == FontSchemeNone || scheme == "" {
		if s != nil {
			f.RemoveChild(s)
		}

		return
	}
	if s == nil {
		s = NewFontSchemeElement()
		f.AppendChild(s)
	}
	s.SetVal(scheme)
}

// VertAlign returns the vertical alignment element, or nil if not present.
func (f *Font) VertAlign() *FontVertAlign {
	elem := f.GetElement(
		"vertAlign",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if v, ok := elem.(*FontVertAlign); ok {
		return v
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontVertAlign{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// VerticalAlignment returns the vertical alignment.
func (f *Font) VerticalAlignment() VerticalAlignRun {
	v := f.VertAlign()
	if v == nil {
		return VerticalAlignRunBaseline
	}

	return v.Val()
}

// SetVerticalAlignment sets the vertical alignment.
func (f *Font) SetVerticalAlignment(
	align VerticalAlignRun,
) {
	v := f.VertAlign()
	if align == VerticalAlignRunBaseline ||
		align == "" {
		if v != nil {
			f.RemoveChild(v)
		}

		return
	}
	if v == nil {
		v = NewFontVertAlign()
		f.AppendChild(v)
	}
	v.SetVal(align)
}

// Clone creates a deep copy of this Font element.
func (f *Font) Clone() openxml.Element {
	cloned := f.CompositeElementBase.Clone()

	return &Font{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Font element.
func (f *Font) CloneNode(
	deep bool,
) openxml.Element {
	cloned := f.CompositeElementBase.CloneNode(
		deep,
	)

	return &Font{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// FontBold represents the bold element (x:b) in a font.
type FontBold struct {
	*openxml.LeafElementBase
}

// NewFontBold creates a new FontBold element.
func NewFontBold() *FontBold {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"b",
		PrefixDefault,
	)

	return &FontBold{LeafElementBase: elem}
}

// Value returns whether bold is enabled.
func (b *FontBold) Value() bool {
	attr, found := b.GetAttribute("val", "")
	if !found {
		return true // Default is true when element present
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetValue sets whether bold is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (b *FontBold) SetValue(value bool) {
	if value {
		b.RemoveAttribute("val", "")
	} else {
		b.SetAttribute(openxml.NewAttribute("", "val", "", attrValueFalse))
	}
}

// Clone creates a deep copy of this element.
func (b *FontBold) Clone() openxml.Element {
	cloned := b.LeafElementBase.Clone()

	return &FontBold{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (b *FontBold) CloneNode(
	deep bool,
) openxml.Element {
	cloned := b.LeafElementBase.CloneNode(deep)

	return &FontBold{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontItalic represents the italic element (x:i) in a font.
type FontItalic struct {
	*openxml.LeafElementBase
}

// NewFontItalic creates a new FontItalic element.
func NewFontItalic() *FontItalic {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"i",
		PrefixDefault,
	)

	return &FontItalic{LeafElementBase: elem}
}

// Value returns whether italic is enabled.
func (i *FontItalic) Value() bool {
	attr, found := i.GetAttribute(
		"val", //nolint:revive // add-constant
		"",
	)
	if !found {
		return true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetValue sets whether italic is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (i *FontItalic) SetValue(value bool) {
	if value {
		i.RemoveAttribute("val", "")
	} else {
		i.SetAttribute(openxml.NewAttribute("", "val", "", attrValueFalse))
	}
}

// Clone creates a deep copy of this element.
func (i *FontItalic) Clone() openxml.Element {
	cloned := i.LeafElementBase.Clone()

	return &FontItalic{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (i *FontItalic) CloneNode(
	deep bool,
) openxml.Element {
	cloned := i.LeafElementBase.CloneNode(deep)

	return &FontItalic{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontStrike represents the strikethrough element (x:strike) in a font.
type FontStrike struct {
	*openxml.LeafElementBase
}

// NewFontStrike creates a new FontStrike element.
func NewFontStrike() *FontStrike {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"strike",
		PrefixDefault,
	)

	return &FontStrike{LeafElementBase: elem}
}

// Value returns whether strikethrough is enabled.
func (s *FontStrike) Value() bool {
	attr, found := s.GetAttribute("val", "")
	if !found {
		return true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetValue sets whether strikethrough is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (s *FontStrike) SetValue(value bool) {
	if value {
		s.RemoveAttribute("val", "")
	} else {
		s.SetAttribute(openxml.NewAttribute("", "val", "", attrValueFalse))
	}
}

// Clone creates a deep copy of this element.
func (s *FontStrike) Clone() openxml.Element {
	cloned := s.LeafElementBase.Clone()

	return &FontStrike{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (s *FontStrike) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.LeafElementBase.CloneNode(deep)

	return &FontStrike{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontUnderline represents the underline element (x:u) in a font.
type FontUnderline struct {
	*openxml.LeafElementBase
}

// NewFontUnderline creates a new FontUnderline element.
func NewFontUnderline() *FontUnderline {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"u",
		PrefixDefault,
	)

	return &FontUnderline{LeafElementBase: elem}
}

// Style returns the underline style.
func (u *FontUnderline) Style() UnderlineStyle {
	attr, found := u.GetAttribute("val", "")
	if !found {
		return UnderlineStyleSingle
	}

	return UnderlineStyle(attr.Value())
}

// SetStyle sets the underline style.
func (u *FontUnderline) SetStyle(
	style UnderlineStyle,
) {
	if style == UnderlineStyleSingle ||
		style == "" {
		u.RemoveAttribute("val", "")
	} else {
		u.SetAttribute(openxml.NewAttribute("", "val", "", string(style)))
	}
}

// Clone creates a deep copy of this element.
func (u *FontUnderline) Clone() openxml.Element {
	cloned := u.LeafElementBase.Clone()

	return &FontUnderline{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (u *FontUnderline) CloneNode(
	deep bool,
) openxml.Element {
	cloned := u.LeafElementBase.CloneNode(deep)

	return &FontUnderline{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontCondense represents the condense element (x:condense) in a font.
type FontCondense struct {
	*openxml.LeafElementBase
}

// NewFontCondense creates a new FontCondense element.
func NewFontCondense() *FontCondense {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"condense",
		PrefixDefault,
	)

	return &FontCondense{LeafElementBase: elem}
}

// Value returns whether condense is enabled.
func (c *FontCondense) Value() bool {
	attr, found := c.GetAttribute("val", "")
	if !found {
		return true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetValue sets whether condense is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *FontCondense) SetValue(value bool) {
	if value {
		c.RemoveAttribute("val", "")
	} else {
		c.SetAttribute(openxml.NewAttribute("", "val", "", attrValueFalse))
	}
}

// Clone creates a deep copy of this element.
func (c *FontCondense) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &FontCondense{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (c *FontCondense) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &FontCondense{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontExtend represents the extend element (x:extend) in a font.
type FontExtend struct {
	*openxml.LeafElementBase
}

// NewFontExtend creates a new FontExtend element.
func NewFontExtend() *FontExtend {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"extend",
		PrefixDefault,
	)

	return &FontExtend{LeafElementBase: elem}
}

// Value returns whether extend is enabled.
func (e *FontExtend) Value() bool {
	attr, found := e.GetAttribute("val", "")
	if !found {
		return true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetValue sets whether extend is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (e *FontExtend) SetValue(value bool) {
	if value {
		e.RemoveAttribute("val", "")
	} else {
		e.SetAttribute(openxml.NewAttribute("", "val", "", attrValueFalse))
	}
}

// Clone creates a deep copy of this element.
func (e *FontExtend) Clone() openxml.Element {
	cloned := e.LeafElementBase.Clone()

	return &FontExtend{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (e *FontExtend) CloneNode(
	deep bool,
) openxml.Element {
	cloned := e.LeafElementBase.CloneNode(deep)

	return &FontExtend{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontOutline represents the outline element (x:outline) in a font.
type FontOutline struct {
	*openxml.LeafElementBase
}

// NewFontOutline creates a new FontOutline element.
func NewFontOutline() *FontOutline {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"outline",
		PrefixDefault,
	)

	return &FontOutline{LeafElementBase: elem}
}

// Value returns whether outline is enabled.
func (o *FontOutline) Value() bool {
	attr, found := o.GetAttribute("val", "")
	if !found {
		return true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetValue sets whether outline is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (o *FontOutline) SetValue(value bool) {
	if value {
		o.RemoveAttribute("val", "")
	} else {
		o.SetAttribute(openxml.NewAttribute("", "val", "", attrValueFalse))
	}
}

// Clone creates a deep copy of this element.
func (o *FontOutline) Clone() openxml.Element {
	cloned := o.LeafElementBase.Clone()

	return &FontOutline{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (o *FontOutline) CloneNode(
	deep bool,
) openxml.Element {
	cloned := o.LeafElementBase.CloneNode(deep)

	return &FontOutline{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontShadow represents the shadow element (x:shadow) in a font.
type FontShadow struct {
	*openxml.LeafElementBase
}

// NewFontShadow creates a new FontShadow element.
func NewFontShadow() *FontShadow {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"shadow",
		PrefixDefault,
	)

	return &FontShadow{LeafElementBase: elem}
}

// Value returns whether shadow is enabled.
func (s *FontShadow) Value() bool {
	attr, found := s.GetAttribute("val", "")
	if !found {
		return true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetValue sets whether shadow is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (s *FontShadow) SetValue(value bool) {
	if value {
		s.RemoveAttribute("val", "")
	} else {
		s.SetAttribute(openxml.NewAttribute("", "val", "", attrValueFalse))
	}
}

// Clone creates a deep copy of this element.
func (s *FontShadow) Clone() openxml.Element {
	cloned := s.LeafElementBase.Clone()

	return &FontShadow{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (s *FontShadow) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.LeafElementBase.CloneNode(deep)

	return &FontShadow{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontSize represents the size element (x:sz) in a font.
type FontSize struct {
	*openxml.LeafElementBase
}

// NewFontSize creates a new FontSize element.
func NewFontSize() *FontSize {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"sz",
		PrefixDefault,
	)

	return &FontSize{LeafElementBase: elem}
}

// Val returns the font size in points.
func (sz *FontSize) Val() float64 {
	attr, found := sz.GetAttribute("val", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetVal sets the font size in points.
func (sz *FontSize) SetVal(size float64) {
	sz.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			strconv.FormatFloat(
				size,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// Clone creates a deep copy of this element.
func (sz *FontSize) Clone() openxml.Element {
	cloned := sz.LeafElementBase.Clone()

	return &FontSize{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (sz *FontSize) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sz.LeafElementBase.CloneNode(deep)

	return &FontSize{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontName represents the name element (x:name) in a font.
type FontName struct {
	*openxml.LeafElementBase
}

// NewFontName creates a new FontName element.
func NewFontName() *FontName {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"name",
		PrefixDefault,
	)

	return &FontName{LeafElementBase: elem}
}

// Val returns the font name.
func (n *FontName) Val() string {
	attr, found := n.GetAttribute("val", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetVal sets the font name.
func (n *FontName) SetVal(name string) {
	n.SetAttribute(
		openxml.NewAttribute("", "val", "", name),
	)
}

// Clone creates a deep copy of this element.
func (n *FontName) Clone() openxml.Element {
	cloned := n.LeafElementBase.Clone()

	return &FontName{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (n *FontName) CloneNode(
	deep bool,
) openxml.Element {
	cloned := n.LeafElementBase.CloneNode(deep)

	return &FontName{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontFamily represents the family element (x:family) in a font.
type FontFamily struct {
	*openxml.LeafElementBase
}

// NewFontFamily creates a new FontFamily element.
func NewFontFamily() *FontFamily {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"family",
		PrefixDefault,
	)

	return &FontFamily{LeafElementBase: elem}
}

// Val returns the font family number.
func (f *FontFamily) Val() uint32 {
	attr, found := f.GetAttribute("val", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetVal sets the font family number.
func (f *FontFamily) SetVal(family uint32) {
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			strconv.FormatUint(
				uint64(family),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Clone creates a deep copy of this element.
func (f *FontFamily) Clone() openxml.Element {
	cloned := f.LeafElementBase.Clone()

	return &FontFamily{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (f *FontFamily) CloneNode(
	deep bool,
) openxml.Element {
	cloned := f.LeafElementBase.CloneNode(deep)

	return &FontFamily{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontCharset represents the charset element (x:charset) in a font.
type FontCharset struct {
	*openxml.LeafElementBase
}

// NewFontCharset creates a new FontCharset element.
func NewFontCharset() *FontCharset {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"charset",
		PrefixDefault,
	)

	return &FontCharset{LeafElementBase: elem}
}

// Val returns the charset value.
func (c *FontCharset) Val() uint32 {
	attr, found := c.GetAttribute("val", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetVal sets the charset value.
func (c *FontCharset) SetVal(charset uint32) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			strconv.FormatUint(
				uint64(charset),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Clone creates a deep copy of this element.
func (c *FontCharset) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &FontCharset{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (c *FontCharset) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &FontCharset{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontSchemeElement represents the scheme element (x:scheme) in a font.
type FontSchemeElement struct {
	*openxml.LeafElementBase
}

// NewFontSchemeElement creates a new FontSchemeElement element.
func NewFontSchemeElement() *FontSchemeElement {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"scheme",
		PrefixDefault,
	)

	return &FontSchemeElement{
		LeafElementBase: elem,
	}
}

// Val returns the font scheme.
func (s *FontSchemeElement) Val() FontScheme {
	attr, found := s.GetAttribute("val", "")
	if !found {
		return FontSchemeNone
	}

	return FontScheme(attr.Value())
}

// SetVal sets the font scheme.
func (s *FontSchemeElement) SetVal(
	scheme FontScheme,
) {
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			string(scheme),
		),
	)
}

// Clone creates a deep copy of this element.
func (s *FontSchemeElement) Clone() openxml.Element {
	cloned := s.LeafElementBase.Clone()

	return &FontSchemeElement{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (s *FontSchemeElement) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.LeafElementBase.CloneNode(deep)

	return &FontSchemeElement{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontVertAlign represents the vertical alignment element (x:vertAlign)
// in a font.
type FontVertAlign struct {
	*openxml.LeafElementBase
}

// NewFontVertAlign creates a new FontVertAlign element.
func NewFontVertAlign() *FontVertAlign {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"vertAlign",
		PrefixDefault,
	)

	return &FontVertAlign{LeafElementBase: elem}
}

// Val returns the vertical alignment.
func (v *FontVertAlign) Val() VerticalAlignRun {
	attr, found := v.GetAttribute("val", "")
	if !found {
		return VerticalAlignRunBaseline
	}

	return VerticalAlignRun(attr.Value())
}

// SetVal sets the vertical alignment.
func (v *FontVertAlign) SetVal(
	align VerticalAlignRun,
) {
	v.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			string(align),
		),
	)
}

// Clone creates a deep copy of this element.
func (v *FontVertAlign) Clone() openxml.Element {
	cloned := v.LeafElementBase.Clone()

	return &FontVertAlign{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (v *FontVertAlign) CloneNode(
	deep bool,
) openxml.Element {
	cloned := v.LeafElementBase.CloneNode(deep)

	return &FontVertAlign{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
