package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PhoneticType represents the type of phonetic text.
type PhoneticType string

const (
	// PhoneticTypeHalfwidthKatakana is half-width Katakana.
	PhoneticTypeHalfwidthKatakana PhoneticType = "halfwidthKatakana"
	// PhoneticTypeFullwidthKatakana is full-width Katakana.
	PhoneticTypeFullwidthKatakana PhoneticType = "fullwidthKatakana"
	// PhoneticTypeHiragana is Hiragana.
	PhoneticTypeHiragana PhoneticType = "Hiragana"
	// PhoneticTypeNoConversion means no phonetic conversion.
	PhoneticTypeNoConversion PhoneticType = "noConversion"
)

// PhoneticAlignment represents the alignment of phonetic text.
type PhoneticAlignment string

const (
	// PhoneticAlignmentNoControl means no alignment control.
	PhoneticAlignmentNoControl PhoneticAlignment = "noControl"
	// PhoneticAlignmentLeft is left alignment.
	PhoneticAlignmentLeft PhoneticAlignment = "left"
	// PhoneticAlignmentCenter is center alignment.
	PhoneticAlignmentCenter PhoneticAlignment = "center"
	// PhoneticAlignmentDistributed is distributed alignment.
	PhoneticAlignmentDistributed PhoneticAlignment = "distributed"
)

// PhoneticRun represents a phonetic run element (x:rPh).
// This element contains phonetic text that provides pronunciation hints,
// typically for East Asian languages (Japanese, Chinese, Korean).
type PhoneticRun struct {
	*openxml.CompositeElementBase
}

// NewPhoneticRun creates a new PhoneticRun element.
func NewPhoneticRun() *PhoneticRun {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"rPh",
		PrefixDefault,
	)

	return &PhoneticRun{
		CompositeElementBase: elem,
	}
}

// Sb returns the starting base text character index.
// This is a zero-based index into the base text.
func (rp *PhoneticRun) Sb() int {
	attr, found := rp.GetAttribute("sb", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetSb sets the starting base text character index.
func (rp *PhoneticRun) SetSb(startBase int) {
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"sb",
			"",
			strconv.Itoa(startBase),
		),
	)
}

// Eb returns the ending base text character index.
// This is a zero-based index into the base text (exclusive).
func (rp *PhoneticRun) Eb() int {
	attr, found := rp.GetAttribute("eb", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetEb sets the ending base text character index.
func (rp *PhoneticRun) SetEb(endBase int) {
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"eb",
			"",
			strconv.Itoa(endBase),
		),
	)
}

// Text returns the text element (x:t) containing the phonetic text,
// or nil if not present.
func (rp *PhoneticRun) Text() *Text {
	elem := rp.GetElement("t", NamespaceSML)
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

// GetOrCreateText returns the text element, creating it if needed.
func (rp *PhoneticRun) GetOrCreateText() *Text {
	t := rp.Text()
	if t != nil {
		return t
	}
	t = NewText()
	rp.AppendChild(t)

	return t
}

// TextContent returns the phonetic text content.
func (rp *PhoneticRun) TextContent() string {
	t := rp.Text()
	if t == nil {
		return ""
	}

	return t.Text()
}

// SetTextContent sets the phonetic text content.
func (rp *PhoneticRun) SetTextContent(
	text string,
) {
	t := rp.GetOrCreateText()
	t.SetText(text)
}

// Clone creates a deep copy of this PhoneticRun element.
func (rp *PhoneticRun) Clone() openxml.Element {
	cloned := rp.CompositeElementBase.Clone()

	return &PhoneticRun{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PhoneticRun element.
func (rp *PhoneticRun) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rp.CompositeElementBase.CloneNode(
		deep,
	)

	return &PhoneticRun{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// PhoneticPr represents the phonetic properties element (x:phoneticPr).
// This element specifies default settings for phonetic text in the string.
type PhoneticPr struct {
	*openxml.LeafElementBase
}

// NewPhoneticPr creates a new PhoneticPr element.
func NewPhoneticPr() *PhoneticPr {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"phoneticPr",
		PrefixDefault,
	)

	return &PhoneticPr{LeafElementBase: elem}
}

// FontId returns the font ID used for phonetic text.
func (pr *PhoneticPr) FontId() int {
	attr, found := pr.GetAttribute("fontId", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetFontId sets the font ID used for phonetic text.
func (pr *PhoneticPr) SetFontId(fontId int) {
	pr.SetAttribute(
		openxml.NewAttribute(
			"",
			"fontId",
			"",
			strconv.Itoa(fontId),
		),
	)
}

// Type returns the phonetic text type.
func (pr *PhoneticPr) Type() PhoneticType {
	attr, found := pr.GetAttribute("type", "")
	if !found {
		return PhoneticTypeFullwidthKatakana // Default
	}

	return PhoneticType(attr.Value())
}

// SetType sets the phonetic text type.
func (pr *PhoneticPr) SetType(t PhoneticType) {
	if t == PhoneticTypeFullwidthKatakana ||
		t == "" {
		pr.RemoveAttribute("type", "")
	} else {
		pr.SetAttribute(
			openxml.NewAttribute("", "type", "", string(t)),
		)
	}
}

// Alignment returns the phonetic text alignment.
func (pr *PhoneticPr) Alignment() PhoneticAlignment {
	attr, found := pr.GetAttribute(
		"alignment",
		"",
	)
	if !found {
		return PhoneticAlignmentLeft // Default
	}

	return PhoneticAlignment(attr.Value())
}

// SetAlignment sets the phonetic text alignment.
func (pr *PhoneticPr) SetAlignment(
	align PhoneticAlignment,
) {
	if align == PhoneticAlignmentLeft ||
		align == "" {
		pr.RemoveAttribute("alignment", "")
	} else {
		pr.SetAttribute(
			openxml.NewAttribute("", "alignment", "", string(align)),
		)
	}
}

// Clone creates a deep copy of this PhoneticPr element.
func (pr *PhoneticPr) Clone() openxml.Element {
	cloned := pr.LeafElementBase.Clone()

	return &PhoneticPr{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this PhoneticPr element.
func (pr *PhoneticPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pr.LeafElementBase.CloneNode(deep)

	return &PhoneticPr{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
