package elements

import (
	"iter"
	"strings"

	"github.com/connerohnesorge/goffice/openxml"
)

// SharedStringItem represents a string item element (x:si) in the shared
// string table. It can contain either simple text (x:t) or rich text runs
// (x:r). It may also contain phonetic run (x:rPh) and phonetic properties
// (x:phoneticPr).
type SharedStringItem struct {
	*openxml.CompositeElementBase
}

// NewSharedStringItem creates a new SharedStringItem element.
func NewSharedStringItem() *SharedStringItem {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"si",
		PrefixDefault,
	)

	return &SharedStringItem{
		CompositeElementBase: elem,
	}
}

// Text returns the simple text element (x:t), or nil if not present.
// If the string item contains rich text runs instead of simple text,
// this returns nil.
func (si *SharedStringItem) Text() *Text {
	elem := si.GetElement("t", NamespaceSML)
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
// Note: This removes any rich text runs, converting to simple text.
func (si *SharedStringItem) GetOrCreateText() *Text {
	t := si.Text()
	if t != nil {
		return t
	}
	t = NewText()
	si.AppendChild(t)

	return t
}

// PlainText returns the plain text content of this string item.
// If the item contains simple text (x:t), returns that text.
// If the item contains rich text runs (x:r), concatenates all text from runs.
func (si *SharedStringItem) PlainText() string {
	// First check for simple text element
	t := si.Text()
	if t != nil {
		return t.Text()
	}

	// Concatenate text from rich text runs
	var sb strings.Builder
	for rtr := range si.RichTextRuns() {
		if text := rtr.Text(); text != nil {
			sb.WriteString(text.Text())
		}
	}

	return sb.String()
}

// SetPlainText sets the content as simple text.
// This removes any rich text runs and uses a single text element.
func (si *SharedStringItem) SetPlainText(
	text string,
) {
	// Remove all rich text runs
	for rtr := range si.RichTextRuns() {
		si.RemoveChild(rtr)
	}

	// Remove phonetic runs and properties
	for rph := range si.PhoneticRuns() {
		si.RemoveChild(rph)
	}
	if pr := si.PhoneticPr(); pr != nil {
		si.RemoveChild(pr)
	}

	t := si.GetOrCreateText()
	t.SetText(text)
}

// RichTextRuns returns an iterator over all rich text run elements (x:r).
func (si *SharedStringItem) RichTextRuns() iter.Seq[*RichTextRun] {
	return func(yield func(*RichTextRun) bool) {
		for child := range si.Children() {
			if child.LocalName() != "r" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var rtr *RichTextRun
			if r, ok := child.(*RichTextRun); ok {
				rtr = r
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				rtr = &RichTextRun{CompositeElementBase: comp}
			}
			if rtr != nil && !yield(rtr) {
				return
			}
		}
	}
}

// AddRichTextRun adds a new rich text run and returns it.
// If the item currently contains simple text, it removes it.
func (si *SharedStringItem) AddRichTextRun() *RichTextRun {
	// Remove plain text if present (switching to rich text)
	t := si.Text()
	if t != nil {
		si.RemoveChild(t)
	}

	rtr := NewRichTextRun()
	// Insert before phonetic elements if present
	pr := si.PhoneticPr()
	if pr != nil {
		si.InsertBefore(rtr, pr)
	} else {
		// Check for phonetic runs and insert before first one
		for rph := range si.PhoneticRuns() {
			si.InsertBefore(rtr, rph)

			return rtr
		}
		si.AppendChild(rtr)
	}

	return rtr
}

// AddRichTextRunWithText adds a new rich text run with the given text.
func (si *SharedStringItem) AddRichTextRunWithText(
	text string,
) *RichTextRun {
	rtr := si.AddRichTextRun()
	rtr.GetOrCreateText().SetText(text)

	return rtr
}

// IsRichText returns whether the string item contains rich text runs.
func (si *SharedStringItem) IsRichText() bool {
	for range si.RichTextRuns() {
		return true
	}

	return false
}

// PhoneticRuns returns an iterator over all phonetic run elements (x:rPh).
func (si *SharedStringItem) PhoneticRuns() iter.Seq[*PhoneticRun] {
	return func(yield func(*PhoneticRun) bool) {
		for child := range si.Children() {
			if child.LocalName() != "rPh" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var rph *PhoneticRun
			if r, ok := child.(*PhoneticRun); ok {
				rph = r
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				rph = &PhoneticRun{CompositeElementBase: comp}
			}
			if rph != nil && !yield(rph) {
				return
			}
		}
	}
}

// AddPhoneticRun adds a new phonetic run and returns it.
func (si *SharedStringItem) AddPhoneticRun(
	startBase, endBase int,
) *PhoneticRun {
	rph := NewPhoneticRun()
	rph.SetSb(startBase)
	rph.SetEb(endBase)

	// Insert before phoneticPr if present
	pr := si.PhoneticPr()
	if pr != nil {
		si.InsertBefore(rph, pr)
	} else {
		si.AppendChild(rph)
	}

	return rph
}

// PhoneticPr returns the phonetic properties element (x:phoneticPr),
// or nil if not present.
func (si *SharedStringItem) PhoneticPr() *PhoneticPr {
	elem := si.GetElement(
		"phoneticPr",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if pr, ok := elem.(*PhoneticPr); ok {
		return pr
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &PhoneticPr{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreatePhoneticPr returns the phonetic properties, creating if needed.
func (si *SharedStringItem) GetOrCreatePhoneticPr() *PhoneticPr {
	pr := si.PhoneticPr()
	if pr != nil {
		return pr
	}
	pr = NewPhoneticPr()
	si.AppendChild(pr)

	return pr
}

// Clone creates a deep copy of this SharedStringItem element.
func (si *SharedStringItem) Clone() openxml.Element {
	cloned := si.CompositeElementBase.Clone()

	return &SharedStringItem{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SharedStringItem element.
func (si *SharedStringItem) CloneNode(
	deep bool,
) openxml.Element {
	cloned := si.CompositeElementBase.CloneNode(
		deep,
	)

	return &SharedStringItem{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
