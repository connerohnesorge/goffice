package elements

import (
	"iter"
	"strings"

	"github.com/connerohnesorge/goffice/openxml"
)

// Run represents an inline content element (w:r).
type Run struct {
	*openxml.CompositeElementBase
}

// NewRun creates a new Run element with optional text content.
func NewRun(text string) *Run {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"r",
		PrefixW,
	)
	r := &Run{CompositeElementBase: elem}
	if text != "" {
		r.AppendChild(NewText(text))
	}

	return r
}

// Properties returns the run properties element, or nil if not present.
func (r *Run) Properties() *RunProperties {
	elem := r.GetElement("rPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*RunProperties); ok {
		return rp
	}
	// Wrap existing element
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RunProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateProperties returns the run properties element,
// creating it if needed.
func (r *Run) GetOrCreateProperties() *RunProperties { //nolint:ireturn
	props := r.Properties()
	if props != nil {
		return props
	}
	props = NewRunProperties()
	// Properties should be first child
	if first := r.FirstChild(); first != nil {
		r.InsertBefore(props, first)
	} else {
		r.AppendChild(props)
	}

	return props
}

// Text returns the first Text element, or nil if not present.
func (r *Run) Text() *Text {
	elem := r.GetElement("t", NamespaceWML)
	if elem == nil {
		return nil
	}
	if t, ok := elem.(*Text); ok {
		return t
	}
	// Wrap existing element
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Text{LeafElementBase: leaf}
	}

	return nil
}

// Texts returns an iterator over all Text elements in this run.
func (r *Run) Texts() iter.Seq[*Text] {
	return func(yield func(*Text) bool) {
		for child := range r.Children() {
			if child.LocalName() != "t" ||
				child.NamespaceURI() != NamespaceWML {
				continue
			}
			var t *Text
			if text, ok := child.(*Text); ok {
				t = text
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				t = &Text{LeafElementBase: leaf}
			}
			if t != nil && !yield(t) {
				return
			}
		}
	}
}

// InnerText returns the concatenated text content of all Text elements.
func (r *Run) InnerText() string {
	var sb strings.Builder
	for t := range r.Texts() {
		sb.WriteString(t.InnerText())
	}

	return sb.String()
}

// SetText sets the text content, replacing any existing text elements.
func (r *Run) SetText(value string) {
	// Remove existing text elements
	var textsToRemove []openxml.Element
	for child := range r.Children() {
		if child.LocalName() == "t" &&
			child.NamespaceURI() == NamespaceWML {
			textsToRemove = append(
				textsToRemove,
				child,
			)
		}
	}
	for _, t := range textsToRemove {
		r.RemoveChild(t)
	}
	// Add new text element
	if value != "" {
		r.AppendChild(NewText(value))
	}
}

// AppendText appends a new Text element with the given content.
func (r *Run) AppendText(value string) *Text {
	t := NewText(value)
	r.AppendChild(t)

	return t
}

// AppendBreak appends a Break element to this run.
func (r *Run) AppendBreak(
	breakType BreakType,
) *Break {
	br := NewBreak(breakType)
	r.AppendChild(br)

	return br
}

// AppendTab appends a Tab element to this run.
func (r *Run) AppendTab() *Tab {
	tab := NewTab()
	r.AppendChild(tab)

	return tab
}

// Clone creates a deep copy of this Run element.
func (r *Run) Clone() openxml.Element {
	cloned := r.CompositeElementBase.Clone()

	return &Run{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Run element.
func (r *Run) CloneNode(
	deep bool,
) openxml.Element {
	cloned := r.CompositeElementBase.CloneNode(
		deep,
	)

	return &Run{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Convenience methods for run properties

// SetBold sets bold formatting on this run.
func (r *Run) SetBold(b bool) *Run {
	r.GetOrCreateProperties().SetBold(b)

	return r
}

// SetItalic sets italic formatting on this run.
func (r *Run) SetItalic(b bool) *Run {
	r.GetOrCreateProperties().SetItalic(b)

	return r
}

// SetUnderline sets underline formatting on this run.
func (r *Run) SetUnderline(
	u UnderlineValue,
) *Run {
	r.GetOrCreateProperties().SetUnderline(u)

	return r
}

// SetFontSize sets the font size in half-points.
func (r *Run) SetFontSize(halfPoints int) *Run {
	r.GetOrCreateProperties().
		SetFontSize(halfPoints)

	return r
}

// SetFont sets the font name.
func (r *Run) SetFont(fontName string) *Run {
	r.GetOrCreateProperties().SetFont(fontName)

	return r
}

// SetColor sets the text color.
func (r *Run) SetColor(hex string) *Run {
	r.GetOrCreateProperties().SetColor(hex)

	return r
}

// SetStrike sets strikethrough formatting on this run.
func (r *Run) SetStrike(b bool) *Run {
	r.GetOrCreateProperties().SetStrike(b)

	return r
}

// SetHighlight sets the highlight color.
func (r *Run) SetHighlight(
	color HighlightColor,
) *Run {
	r.GetOrCreateProperties().SetHighlight(color)

	return r
}

// SetVerticalTextAlignment sets vertical text alignment
// (subscript, superscript, baseline).
func (r *Run) SetVerticalTextAlignment(
	v VerticalAlignValue,
) *Run {
	r.GetOrCreateProperties().
		SetVerticalTextAlignment(v)

	return r
}

// AppendFootnoteReference appends a footnote reference to this run.
func (r *Run) AppendFootnoteReference(
	id int,
) *FootnoteReference {
	fr := NewFootnoteReference(id)
	r.AppendChild(fr)

	return fr
}

// AppendEndnoteReference appends an endnote reference to this run.
func (r *Run) AppendEndnoteReference(
	id int,
) *EndnoteReference {
	er := NewEndnoteReference(id)
	r.AppendChild(er)

	return er
}

// AppendCommentReference appends a comment reference to this run.
func (r *Run) AppendCommentReference(
	id int,
) *CommentReference {
	cr := NewCommentReference(id)
	r.AppendChild(cr)

	return cr
}

// AppendCarriageReturn appends a carriage return element to this run.
func (r *Run) AppendCarriageReturn() *CarriageReturn {
	cr := NewCarriageReturn()
	r.AppendChild(cr)

	return cr
}

// AppendSoftHyphen appends a soft hyphen element to this run.
func (r *Run) AppendSoftHyphen() *SoftHyphen {
	sh := NewSoftHyphen()
	r.AppendChild(sh)

	return sh
}

// AppendNoBreakHyphen appends a non-breaking hyphen element to this run.
func (r *Run) AppendNoBreakHyphen() *NoBreakHyphen {
	nbh := NewNoBreakHyphen()
	r.AppendChild(nbh)

	return nbh
}

// AppendSymbol appends a symbol element to this run.
func (r *Run) AppendSymbol(
	font, char string,
) *Symbol {
	s := NewSymbol(font, char)
	r.AppendChild(s)

	return s
}

// AppendFieldChar appends a field character element to this run.
func (r *Run) AppendFieldChar(
	charType FieldCharType,
) *FieldChar {
	fc := NewFieldChar(charType)
	r.AppendChild(fc)

	return fc
}

// AppendInstrText appends field instruction text to this run.
func (r *Run) AppendInstrText(
	text string,
) *InstrText {
	it := NewInstrText(text)
	r.AppendChild(it)

	return it
}

// AppendPositionalTab appends a positional tab element to this run.
func (r *Run) AppendPositionalTab() *PositionalTab {
	pt := NewPositionalTab()
	r.AppendChild(pt)

	return pt
}

// AppendSeparator appends a separator element to this run.
func (r *Run) AppendSeparator() *Separator {
	s := NewSeparator()
	r.AppendChild(s)

	return s
}

// AppendContinuationSeparator appends a continuation separator
// element to this run.
func (r *Run) AppendContinuationSeparator() *ContinuationSeparator {
	cs := NewContinuationSeparator()
	r.AppendChild(cs)

	return cs
}

// AppendLastRenderedPageBreak appends a last rendered page break
// element to this run.
func (r *Run) AppendLastRenderedPageBreak() *LastRenderedPageBreak {
	lrpb := NewLastRenderedPageBreak()
	r.AppendChild(lrpb)

	return lrpb
}
