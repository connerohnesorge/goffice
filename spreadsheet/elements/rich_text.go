package elements

//revive:disable:file-length-limit many rich text properties
//revive:disable:max-public-structs many rich text types needed

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// VerticalAlignRun represents the vertical alignment of text in a run.
type VerticalAlignRun string

const (
	// VerticalAlignRunBaseline is baseline alignment (default).
	VerticalAlignRunBaseline VerticalAlignRun = "baseline"
	// VerticalAlignRunSuperscript is superscript alignment.
	VerticalAlignRunSuperscript VerticalAlignRun = "superscript"
	// VerticalAlignRunSubscript is subscript alignment.
	VerticalAlignRunSubscript VerticalAlignRun = "subscript"
)

// UnderlineStyle represents the underline style.
type UnderlineStyle string

const (
	// UnderlineStyleNone is no underline.
	UnderlineStyleNone UnderlineStyle = "none"
	// UnderlineStyleSingle is a single underline.
	UnderlineStyleSingle UnderlineStyle = "single"
	// UnderlineStyleDouble is a double underline.
	UnderlineStyleDouble UnderlineStyle = "double"
	// UnderlineStyleSingleAccounting is an accounting-style single underline.
	UnderlineStyleSingleAccounting UnderlineStyle = "singleAccounting"
	// UnderlineStyleDoubleAccounting is an accounting-style double underline.
	UnderlineStyleDoubleAccounting UnderlineStyle = "doubleAccounting"
)

// Text represents the text element (x:t).
// This element contains the actual text content of a rich text string.
type Text struct {
	*openxml.LeafElementBase
}

// NewText creates a new Text element.
func NewText() *Text {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"t",
		PrefixDefault,
	)

	return &Text{LeafElementBase: elem}
}

// NewTextWithContent creates a new Text element with initial content.
func NewTextWithContent(content string) *Text {
	elem := openxml.NewLeafElementWithText(
		NamespaceSML,
		"t",
		PrefixDefault,
		content,
	)

	return &Text{LeafElementBase: elem}
}

// Text returns the text content.
func (t *Text) Text() string {
	return t.InnerText()
}

// SetText sets the text content.
func (t *Text) SetText(text string) {
	t.SetInnerText(text)
}

// PreserveSpace returns whether whitespace should be preserved.
// This is indicated by the xml:space="preserve" attribute.
func (t *Text) PreserveSpace() bool {
	attr, found := t.GetAttribute(
		"space",
		"http://www.w3.org/XML/1998/namespace",
	)
	if !found {
		return false
	}

	return attr.Value() == "preserve"
}

// SetPreserveSpace sets whether whitespace should be preserved.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Text) SetPreserveSpace(preserve bool) {
	if preserve {
		t.SetAttribute(
			openxml.NewAttribute(
				"http://www.w3.org/XML/1998/namespace",
				"space",
				"xml",
				"preserve",
			),
		)
	} else {
		t.RemoveAttribute("space", "http://www.w3.org/XML/1998/namespace")
	}
}

// Clone creates a deep copy of this Text element.
func (t *Text) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &Text{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Text element.
func (t *Text) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.LeafElementBase.CloneNode(deep)

	return &Text{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// RichTextRun represents a rich text run element (x:r).
// It contains optional formatting properties and text content.
type RichTextRun struct {
	*openxml.CompositeElementBase
}

// NewRichTextRun creates a new RichTextRun element.
func NewRichTextRun() *RichTextRun {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"r",
		PrefixDefault,
	)

	return &RichTextRun{
		CompositeElementBase: elem,
	}
}

// RunProperties returns the run properties (x:rPr), or nil if not present.
func (rt *RichTextRun) RunProperties() *RunProperties {
	elem := rt.GetElement("rPr", NamespaceSML)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*RunProperties); ok {
		return rp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RunProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateRunProperties returns the run properties, creating if needed.
func (rt *RichTextRun) GetOrCreateRunProperties() *RunProperties {
	rp := rt.RunProperties()
	if rp != nil {
		return rp
	}
	rp = NewRunProperties()
	// Properties should come before text
	t := rt.Text()
	if t != nil {
		rt.InsertBefore(rp, t)
	} else {
		rt.AppendChild(rp)
	}

	return rp
}

// Text returns the text element (x:t), or nil if not present.
func (rt *RichTextRun) Text() *Text {
	elem := rt.GetElement("t", NamespaceSML)
	if elem == nil {
		return nil
	}
	if t, ok := elem.(*Text); ok {
		return t
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Text{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateText returns the text element, creating if needed.
func (rt *RichTextRun) GetOrCreateText() *Text {
	t := rt.Text()
	if t != nil {
		return t
	}
	t = NewText()
	rt.AppendChild(t)

	return t
}

// TextContent returns the text content as a string.
func (rt *RichTextRun) TextContent() string {
	t := rt.Text()
	if t == nil {
		return ""
	}

	return t.Text()
}

// SetTextContent sets the text content.
func (rt *RichTextRun) SetTextContent(
	text string,
) {
	t := rt.GetOrCreateText()
	t.SetText(text)
}

// Clone creates a deep copy of this RichTextRun element.
func (rt *RichTextRun) Clone() openxml.Element {
	cloned := rt.CompositeElementBase.Clone()

	return &RichTextRun{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this RichTextRun element.
func (rt *RichTextRun) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rt.CompositeElementBase.CloneNode(
		deep,
	)

	return &RichTextRun{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// RunProperties represents the run properties element (x:rPr).
// This contains font formatting for a rich text run.
type RunProperties struct {
	*openxml.CompositeElementBase
}

// NewRunProperties creates a new RunProperties element.
func NewRunProperties() *RunProperties {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"rPr",
		PrefixDefault,
	)

	return &RunProperties{
		CompositeElementBase: elem,
	}
}

// Bold returns the bold element, or nil if not present.
func (rp *RunProperties) Bold() *BoldProperty {
	elem := rp.GetElement("b", NamespaceSML)
	if elem == nil {
		return nil
	}
	if b, ok := elem.(*BoldProperty); ok {
		return b
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &BoldProperty{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// IsBold returns whether the run is bold.
func (rp *RunProperties) IsBold() bool {
	b := rp.Bold()
	if b == nil {
		return false
	}

	return b.Value()
}

// SetBold sets whether the run is bold.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (rp *RunProperties) SetBold(bold bool) {
	b := rp.Bold()
	if bold {
		if b == nil {
			b = NewBoldProperty()
			rp.AppendChild(b)
		}
		b.SetValue(true)
	} else if b != nil {
		rp.RemoveChild(b)
	}
}

// Italic returns the italic element, or nil if not present.
func (rp *RunProperties) Italic() *ItalicProperty {
	elem := rp.GetElement("i", NamespaceSML)
	if elem == nil {
		return nil
	}
	if i, ok := elem.(*ItalicProperty); ok {
		return i
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &ItalicProperty{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// IsItalic returns whether the run is italic.
func (rp *RunProperties) IsItalic() bool {
	i := rp.Italic()
	if i == nil {
		return false
	}

	return i.Value()
}

// SetItalic sets whether the run is italic.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (rp *RunProperties) SetItalic(italic bool) {
	i := rp.Italic()
	if italic {
		if i == nil {
			i = NewItalicProperty()
			rp.AppendChild(i)
		}
		i.SetValue(true)
	} else if i != nil {
		rp.RemoveChild(i)
	}
}

// Strike returns the strikethrough element, or nil if not present.
func (rp *RunProperties) Strike() *StrikeProperty {
	elem := rp.GetElement("strike", NamespaceSML)
	if elem == nil {
		return nil
	}
	if s, ok := elem.(*StrikeProperty); ok {
		return s
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &StrikeProperty{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// IsStrikethrough returns whether the run has strikethrough.
func (rp *RunProperties) IsStrikethrough() bool {
	s := rp.Strike()
	if s == nil {
		return false
	}

	return s.Value()
}

// SetStrikethrough sets whether the run has strikethrough.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (rp *RunProperties) SetStrikethrough(
	strike bool,
) {
	s := rp.Strike()
	if strike {
		if s == nil {
			s = NewStrikeProperty()
			rp.AppendChild(s)
		}
		s.SetValue(true)
	} else if s != nil {
		rp.RemoveChild(s)
	}
}

// Underline returns the underline element, or nil if not present.
func (rp *RunProperties) Underline() *UnderlineProperty {
	elem := rp.GetElement("u", NamespaceSML)
	if elem == nil {
		return nil
	}
	if u, ok := elem.(*UnderlineProperty); ok {
		return u
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &UnderlineProperty{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// UnderlineStyle returns the underline style, or empty string if none.
func (rp *RunProperties) UnderlineStyle() UnderlineStyle {
	u := rp.Underline()
	if u == nil {
		return UnderlineStyleNone
	}

	return u.Style()
}

// SetUnderline sets the underline style.
func (rp *RunProperties) SetUnderline(
	style UnderlineStyle,
) {
	u := rp.Underline()
	if style == UnderlineStyleNone ||
		style == "" {
		if u != nil {
			rp.RemoveChild(u)
		}

		return
	}
	if u == nil {
		u = NewUnderlineProperty()
		rp.AppendChild(u)
	}
	u.SetStyle(style)
}

// VertAlign returns the vertical alignment element, or nil if not present.
func (rp *RunProperties) VertAlign() *VertAlignProperty {
	elem := rp.GetElement(
		"vertAlign",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if v, ok := elem.(*VertAlignProperty); ok {
		return v
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &VertAlignProperty{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// VerticalAlignment returns the vertical alignment.
func (rp *RunProperties) VerticalAlignment() VerticalAlignRun {
	v := rp.VertAlign()
	if v == nil {
		return VerticalAlignRunBaseline
	}

	return v.Val()
}

// SetVerticalAlignment sets the vertical alignment.
func (rp *RunProperties) SetVerticalAlignment(
	align VerticalAlignRun,
) {
	v := rp.VertAlign()
	if align == VerticalAlignRunBaseline ||
		align == "" {
		if v != nil {
			rp.RemoveChild(v)
		}

		return
	}
	if v == nil {
		v = NewVertAlignProperty()
		rp.AppendChild(v)
	}
	v.SetVal(align)
}

// FontSize returns the font size element, or nil if not present.
func (rp *RunProperties) FontSize() *FontSizeProperty {
	elem := rp.GetElement("sz", NamespaceSML)
	if elem == nil {
		return nil
	}
	if sz, ok := elem.(*FontSizeProperty); ok {
		return sz
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontSizeProperty{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// Size returns the font size in points.
func (rp *RunProperties) Size() float64 {
	sz := rp.FontSize()
	if sz == nil {
		return 0
	}

	return sz.Val()
}

// SetSize sets the font size in points.
func (rp *RunProperties) SetSize(size float64) {
	sz := rp.FontSize()
	if size == 0 {
		if sz != nil {
			rp.RemoveChild(sz)
		}

		return
	}
	if sz == nil {
		sz = NewFontSizeProperty()
		rp.AppendChild(sz)
	}
	sz.SetVal(size)
}

// FontName returns the font name element, or nil if not present.
func (rp *RunProperties) FontName() *FontNameProperty {
	elem := rp.GetElement("rFont", NamespaceSML)
	if elem == nil {
		return nil
	}
	if fn, ok := elem.(*FontNameProperty); ok {
		return fn
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FontNameProperty{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// Font returns the font name.
func (rp *RunProperties) Font() string {
	fn := rp.FontName()
	if fn == nil {
		return ""
	}

	return fn.Val()
}

// SetFont sets the font name.
func (rp *RunProperties) SetFont(name string) {
	fn := rp.FontName()
	if name == "" {
		if fn != nil {
			rp.RemoveChild(fn)
		}

		return
	}
	if fn == nil {
		fn = NewFontNameProperty()
		rp.AppendChild(fn)
	}
	fn.SetVal(name)
}

// FontColor returns the color element, or nil if not present.
func (rp *RunProperties) FontColor() *ColorProperty {
	elem := rp.GetElement("color", NamespaceSML)
	if elem == nil {
		return nil
	}
	if c, ok := elem.(*ColorProperty); ok {
		return c
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &ColorProperty{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreateFontColor returns the font color, creating if needed.
func (rp *RunProperties) GetOrCreateFontColor() *ColorProperty {
	c := rp.FontColor()
	if c != nil {
		return c
	}
	c = NewColorProperty()
	rp.AppendChild(c)

	return c
}

// Clone creates a deep copy of this RunProperties element.
func (rp *RunProperties) Clone() openxml.Element {
	cloned := rp.CompositeElementBase.Clone()

	return &RunProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this RunProperties element.
func (rp *RunProperties) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rp.CompositeElementBase.CloneNode(
		deep,
	)

	return &RunProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// BoldProperty represents the bold property element (x:b).
type BoldProperty struct {
	*openxml.LeafElementBase
}

// NewBoldProperty creates a new BoldProperty element.
func NewBoldProperty() *BoldProperty {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"b",
		PrefixDefault,
	)

	return &BoldProperty{LeafElementBase: elem}
}

// Value returns whether bold is enabled.
func (b *BoldProperty) Value() bool {
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
func (b *BoldProperty) SetValue(value bool) {
	if value {
		b.RemoveAttribute("val", "")
	} else {
		b.SetAttribute(
			openxml.NewAttribute("", "val", "", attrValueFalse),
		)
	}
}

// Clone creates a deep copy of this element.
func (b *BoldProperty) Clone() openxml.Element {
	cloned := b.LeafElementBase.Clone()

	return &BoldProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (b *BoldProperty) CloneNode(
	deep bool,
) openxml.Element {
	cloned := b.LeafElementBase.CloneNode(deep)

	return &BoldProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// ItalicProperty represents the italic property element (x:i).
type ItalicProperty struct {
	*openxml.LeafElementBase
}

// NewItalicProperty creates a new ItalicProperty element.
func NewItalicProperty() *ItalicProperty {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"i",
		PrefixDefault,
	)

	return &ItalicProperty{LeafElementBase: elem}
}

// Value returns whether italic is enabled.
func (i *ItalicProperty) Value() bool {
	attr, found := i.GetAttribute(
		"val", //nolint:revive // add-constant
		"",
	)
	if !found {
		return true // Default is true when element present
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetValue sets whether italic is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (i *ItalicProperty) SetValue(value bool) {
	if value {
		i.RemoveAttribute("val", "")
	} else {
		i.SetAttribute(
			openxml.NewAttribute("", "val", "", attrValueFalse),
		)
	}
}

// Clone creates a deep copy of this element.
func (i *ItalicProperty) Clone() openxml.Element {
	cloned := i.LeafElementBase.Clone()

	return &ItalicProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (i *ItalicProperty) CloneNode(
	deep bool,
) openxml.Element {
	cloned := i.LeafElementBase.CloneNode(deep)

	return &ItalicProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// StrikeProperty represents the strikethrough property element (x:strike).
type StrikeProperty struct {
	*openxml.LeafElementBase
}

// NewStrikeProperty creates a new StrikeProperty element.
func NewStrikeProperty() *StrikeProperty {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"strike",
		PrefixDefault,
	)

	return &StrikeProperty{LeafElementBase: elem}
}

// Value returns whether strikethrough is enabled.
func (s *StrikeProperty) Value() bool {
	attr, found := s.GetAttribute("val", "")
	if !found {
		return true // Default is true when element present
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetValue sets whether strikethrough is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (s *StrikeProperty) SetValue(value bool) {
	if value {
		s.RemoveAttribute("val", "")
	} else {
		s.SetAttribute(
			openxml.NewAttribute("", "val", "", attrValueFalse),
		)
	}
}

// Clone creates a deep copy of this element.
func (s *StrikeProperty) Clone() openxml.Element {
	cloned := s.LeafElementBase.Clone()

	return &StrikeProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (s *StrikeProperty) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.LeafElementBase.CloneNode(deep)

	return &StrikeProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// UnderlineProperty represents the underline property element (x:u).
type UnderlineProperty struct {
	*openxml.LeafElementBase
}

// NewUnderlineProperty creates a new UnderlineProperty element.
func NewUnderlineProperty() *UnderlineProperty {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"u",
		PrefixDefault,
	)

	return &UnderlineProperty{
		LeafElementBase: elem,
	}
}

// Style returns the underline style.
func (u *UnderlineProperty) Style() UnderlineStyle {
	attr, found := u.GetAttribute("val", "")
	if !found {
		return UnderlineStyleSingle // Default when element present
	}

	return UnderlineStyle(attr.Value())
}

// SetStyle sets the underline style.
func (u *UnderlineProperty) SetStyle(
	style UnderlineStyle,
) {
	if style == UnderlineStyleSingle ||
		style == "" {
		u.RemoveAttribute("val", "")
	} else {
		u.SetAttribute(
			openxml.NewAttribute("", "val", "", string(style)),
		)
	}
}

// Clone creates a deep copy of this element.
func (u *UnderlineProperty) Clone() openxml.Element {
	cloned := u.LeafElementBase.Clone()

	return &UnderlineProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (u *UnderlineProperty) CloneNode(
	deep bool,
) openxml.Element {
	cloned := u.LeafElementBase.CloneNode(deep)

	return &UnderlineProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// VertAlignProperty represents the vertical alignment element (x:vertAlign).
type VertAlignProperty struct {
	*openxml.LeafElementBase
}

// NewVertAlignProperty creates a new VertAlignProperty element.
func NewVertAlignProperty() *VertAlignProperty {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"vertAlign",
		PrefixDefault,
	)

	return &VertAlignProperty{
		LeafElementBase: elem,
	}
}

// Val returns the vertical alignment value.
func (v *VertAlignProperty) Val() VerticalAlignRun {
	attr, found := v.GetAttribute("val", "")
	if !found {
		return VerticalAlignRunBaseline
	}

	return VerticalAlignRun(attr.Value())
}

// SetVal sets the vertical alignment value.
func (v *VertAlignProperty) SetVal(
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
func (v *VertAlignProperty) Clone() openxml.Element {
	cloned := v.LeafElementBase.Clone()

	return &VertAlignProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (v *VertAlignProperty) CloneNode(
	deep bool,
) openxml.Element {
	cloned := v.LeafElementBase.CloneNode(deep)

	return &VertAlignProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontSizeProperty represents the font size element (x:sz).
type FontSizeProperty struct {
	*openxml.LeafElementBase
}

// NewFontSizeProperty creates a new FontSizeProperty element.
func NewFontSizeProperty() *FontSizeProperty {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"sz",
		PrefixDefault,
	)

	return &FontSizeProperty{
		LeafElementBase: elem,
	}
}

// Val returns the font size in points.
func (sz *FontSizeProperty) Val() float64 {
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
func (sz *FontSizeProperty) SetVal(size float64) {
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
func (sz *FontSizeProperty) Clone() openxml.Element {
	cloned := sz.LeafElementBase.Clone()

	return &FontSizeProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (sz *FontSizeProperty) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sz.LeafElementBase.CloneNode(deep)

	return &FontSizeProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FontNameProperty represents the font name element (x:rFont).
type FontNameProperty struct {
	*openxml.LeafElementBase
}

// NewFontNameProperty creates a new FontNameProperty element.
func NewFontNameProperty() *FontNameProperty {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"rFont",
		PrefixDefault,
	)

	return &FontNameProperty{
		LeafElementBase: elem,
	}
}

// Val returns the font name.
func (fn *FontNameProperty) Val() string {
	attr, found := fn.GetAttribute("val", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetVal sets the font name.
func (fn *FontNameProperty) SetVal(name string) {
	fn.SetAttribute(
		openxml.NewAttribute("", "val", "", name),
	)
}

// Clone creates a deep copy of this element.
func (fn *FontNameProperty) Clone() openxml.Element {
	cloned := fn.LeafElementBase.Clone()

	return &FontNameProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (fn *FontNameProperty) CloneNode(
	deep bool,
) openxml.Element {
	cloned := fn.LeafElementBase.CloneNode(deep)

	return &FontNameProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// ColorProperty represents the color element (x:color).
type ColorProperty struct {
	*openxml.LeafElementBase
}

// NewColorProperty creates a new ColorProperty element.
func NewColorProperty() *ColorProperty {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"color",
		PrefixDefault,
	)

	return &ColorProperty{LeafElementBase: elem}
}

// RGB returns the RGB color value (e.g., "FF0000" for red).
func (c *ColorProperty) RGB() string {
	attr, found := c.GetAttribute("rgb", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRGB sets the RGB color value (e.g., "FF0000" for red).
func (c *ColorProperty) SetRGB(rgb string) {
	if rgb == "" {
		c.RemoveAttribute("rgb", "")
	} else {
		c.SetAttribute(
			openxml.NewAttribute("", "rgb", "", rgb),
		)
	}
}

// Theme returns the theme color index.
func (c *ColorProperty) Theme() int {
	attr, found := c.GetAttribute("theme", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetTheme sets the theme color index.
func (c *ColorProperty) SetTheme(theme int) {
	if theme == 0 {
		c.RemoveAttribute("theme", "")
	} else {
		c.SetAttribute(
			openxml.NewAttribute("", "theme", "", strconv.Itoa(theme)),
		)
	}
}

// Indexed returns the indexed color value.
func (c *ColorProperty) Indexed() int {
	attr, found := c.GetAttribute("indexed", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetIndexed sets the indexed color value.
func (c *ColorProperty) SetIndexed(indexed int) {
	if indexed == 0 {
		c.RemoveAttribute("indexed", "")
	} else {
		c.SetAttribute(
			openxml.NewAttribute("", "indexed", "", strconv.Itoa(indexed)),
		)
	}
}

// Tint returns the tint value (-1.0 to 1.0).
func (c *ColorProperty) Tint() float64 {
	attr, found := c.GetAttribute("tint", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetTint sets the tint value (-1.0 to 1.0).
func (c *ColorProperty) SetTint(tint float64) {
	if tint == 0 {
		c.RemoveAttribute("tint", "")
	} else {
		tintStr := strconv.FormatFloat(tint, 'f', -1, 64) //nolint:revive // add-constant
		c.SetAttribute(openxml.NewAttribute("", "tint", "", tintStr))
	}
}

// Clone creates a deep copy of this element.
func (c *ColorProperty) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &ColorProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (c *ColorProperty) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &ColorProperty{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
