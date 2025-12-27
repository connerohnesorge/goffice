package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// InlineString represents the inline string element (x:is).
// This is used for cells that contain string data directly rather than
// referencing the shared string table.
// It can contain either plain text (x:t) or rich text runs (x:r).
type InlineString struct {
	*openxml.CompositeElementBase
}

// NewInlineString creates a new InlineString element.
func NewInlineString() *InlineString {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"is",
		PrefixDefault,
	)

	return &InlineString{
		CompositeElementBase: elem,
	}
}

// Text returns the plain text element (x:t), or nil if not present.
func (is *InlineString) Text() *Text {
	elem := is.GetElement("t", NamespaceSML)
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

// GetOrCreateText returns the plain text element, creating it if needed.
// Note: Using plain text removes rich text runs.
func (is *InlineString) GetOrCreateText() *Text {
	t := is.Text()
	if t != nil {
		return t
	}
	t = NewText()
	is.AppendChild(t)

	return t
}

// SetPlainText sets the plain text content.
// This will remove any rich text runs and use a single text element.
func (is *InlineString) SetPlainText(
	text string,
) {
	// Remove all rich text runs
	for rtr := range is.RichTextRuns() {
		is.RemoveChild(rtr)
	}

	t := is.GetOrCreateText()
	t.SetText(text)
}

// PlainText returns the plain text content.
// If the inline string contains rich text runs, this concatenates all text.
func (is *InlineString) PlainText() string {
	// First check for simple text element
	t := is.Text()
	if t != nil {
		return t.Text()
	}

	// Concatenate text from rich text runs
	var result string
	for rtr := range is.RichTextRuns() {
		if text := rtr.Text(); text != nil {
			result += text.Text()
		}
	}

	return result
}

// RichTextRuns returns an iterator over all rich text run elements (x:r).
func (is *InlineString) RichTextRuns() iter.Seq[*RichTextRun] {
	return func(yield func(*RichTextRun) bool) {
		for child := range is.Children() {
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
func (is *InlineString) AddRichTextRun() *RichTextRun {
	// Remove plain text if present (switching to rich text)
	t := is.Text()
	if t != nil {
		is.RemoveChild(t)
	}

	rtr := NewRichTextRun()
	is.AppendChild(rtr)

	return rtr
}

// AddRichTextRunWithText adds a new rich text run with the given text.
func (is *InlineString) AddRichTextRunWithText(
	text string,
) *RichTextRun {
	rtr := is.AddRichTextRun()
	rtr.GetOrCreateText().SetText(text)

	return rtr
}

// HasRichText returns whether the inline string contains rich text runs.
func (is *InlineString) HasRichText() bool {
	for range is.RichTextRuns() {
		return true
	}

	return false
}

// Clone creates a deep copy of this InlineString element.
func (is *InlineString) Clone() openxml.Element {
	cloned := is.CompositeElementBase.Clone()

	return &InlineString{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this InlineString element.
func (is *InlineString) CloneNode(
	deep bool,
) openxml.Element {
	cloned := is.CompositeElementBase.CloneNode(
		deep,
	)

	return &InlineString{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
