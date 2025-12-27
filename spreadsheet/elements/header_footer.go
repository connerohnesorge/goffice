package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// HeaderFooter represents the header/footer element (x:headerFooter).
type HeaderFooter struct {
	*openxml.CompositeElementBase
}

// NewHeaderFooter creates a new HeaderFooter element.
func NewHeaderFooter() *HeaderFooter {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"headerFooter",
		PrefixDefault,
	)

	return &HeaderFooter{
		CompositeElementBase: elem,
	}
}

// DifferentOddEven returns whether different headers/footers are used for
// odd and even pages.
func (hf *HeaderFooter) DifferentOddEven() bool {
	attr, found := hf.GetAttribute(
		"differentOddEven",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDifferentOddEven sets whether different headers/footers are used for
// odd and even pages.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (hf *HeaderFooter) SetDifferentOddEven(
	value bool,
) {
	if value {
		hf.SetAttribute(
			openxml.NewAttribute(
				"",
				"differentOddEven",
				"",
				attrValueTrue,
			),
		)
	} else {
		hf.RemoveAttribute("differentOddEven", "")
	}
}

// DifferentFirst returns whether a different header/footer is used for the
// first page.
func (hf *HeaderFooter) DifferentFirst() bool {
	attr, found := hf.GetAttribute(
		"differentFirst",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDifferentFirst sets whether a different header/footer is used for the
// first page.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (hf *HeaderFooter) SetDifferentFirst(
	value bool,
) {
	if value {
		hf.SetAttribute(
			openxml.NewAttribute(
				"",
				"differentFirst",
				"",
				attrValueTrue,
			),
		)
	} else {
		hf.RemoveAttribute("differentFirst", "")
	}
}

// ScaleWithDoc returns whether the header/footer scales with the document.
func (hf *HeaderFooter) ScaleWithDoc() bool {
	attr, found := hf.GetAttribute(
		"scaleWithDoc",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetScaleWithDoc sets whether the header/footer scales with the document.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (hf *HeaderFooter) SetScaleWithDoc(
	value bool,
) {
	if value {
		hf.RemoveAttribute("scaleWithDoc", "")
	} else {
		hf.SetAttribute(openxml.NewAttribute("", "scaleWithDoc", "", attrValueFalse))
	}
}

// AlignWithMargins returns whether the header/footer aligns with page margins.
func (hf *HeaderFooter) AlignWithMargins() bool {
	attr, found := hf.GetAttribute(
		"alignWithMargins",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetAlignWithMargins sets whether the header/footer aligns with page margins.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (hf *HeaderFooter) SetAlignWithMargins(
	value bool,
) {
	if value {
		hf.RemoveAttribute("alignWithMargins", "")
	} else {
		hf.SetAttribute(openxml.NewAttribute("", "alignWithMargins", "", attrValueFalse))
	}
}

// getOrCreateChild returns or creates a leaf child element with the given name.
func (hf *HeaderFooter) getOrCreateChild(
	name string,
) *openxml.LeafElementBase {
	elem := hf.GetElement(name, NamespaceSML)
	if elem != nil {
		if leaf, ok := elem.(*openxml.LeafElementBase); ok {
			return leaf
		}
	}
	leaf := openxml.NewLeafElement(
		NamespaceSML,
		name,
		PrefixDefault,
	)
	hf.AppendChild(leaf)

	return leaf
}

// getChildText returns the text content of a child element.
func (hf *HeaderFooter) getChildText(
	name string,
) string {
	elem := hf.GetElement(name, NamespaceSML)
	if elem == nil {
		return ""
	}
	if leaf, ok := elem.(openxml.LeafElement); ok {
		return leaf.InnerText()
	}

	return ""
}

// setChildText sets the text content of a child element.
func (hf *HeaderFooter) setChildText(
	name, text string,
) {
	if text == "" {
		// Remove the element if text is empty
		elem := hf.GetElement(name, NamespaceSML)
		if elem != nil {
			hf.RemoveChild(elem)
		}

		return
	}
	leaf := hf.getOrCreateChild(name)
	leaf.SetInnerText(text)
}

// OddHeader returns the odd page header string.
func (hf *HeaderFooter) OddHeader() string {
	return hf.getChildText("oddHeader")
}

// SetOddHeader sets the odd page header string.
func (hf *HeaderFooter) SetOddHeader(
	header string,
) {
	hf.setChildText("oddHeader", header)
}

// OddFooter returns the odd page footer string.
func (hf *HeaderFooter) OddFooter() string {
	return hf.getChildText("oddFooter")
}

// SetOddFooter sets the odd page footer string.
func (hf *HeaderFooter) SetOddFooter(
	footer string,
) {
	hf.setChildText("oddFooter", footer)
}

// EvenHeader returns the even page header string.
func (hf *HeaderFooter) EvenHeader() string {
	return hf.getChildText("evenHeader")
}

// SetEvenHeader sets the even page header string.
func (hf *HeaderFooter) SetEvenHeader(
	header string,
) {
	hf.setChildText("evenHeader", header)
}

// EvenFooter returns the even page footer string.
func (hf *HeaderFooter) EvenFooter() string {
	return hf.getChildText("evenFooter")
}

// SetEvenFooter sets the even page footer string.
func (hf *HeaderFooter) SetEvenFooter(
	footer string,
) {
	hf.setChildText("evenFooter", footer)
}

// FirstHeader returns the first page header string.
func (hf *HeaderFooter) FirstHeader() string {
	return hf.getChildText("firstHeader")
}

// SetFirstHeader sets the first page header string.
func (hf *HeaderFooter) SetFirstHeader(
	header string,
) {
	hf.setChildText("firstHeader", header)
}

// FirstFooter returns the first page footer string.
func (hf *HeaderFooter) FirstFooter() string {
	return hf.getChildText("firstFooter")
}

// SetFirstFooter sets the first page footer string.
func (hf *HeaderFooter) SetFirstFooter(
	footer string,
) {
	hf.setChildText("firstFooter", footer)
}

// SetHeader sets the header for all pages (odd header).
// This is the common case when not using different odd/even
// or first page headers.
func (hf *HeaderFooter) SetHeader(header string) {
	hf.SetOddHeader(header)
}

// SetFooter sets the footer for all pages (odd footer).
// This is the common case when not using different odd/even
// or first page footers.
func (hf *HeaderFooter) SetFooter(footer string) {
	hf.SetOddFooter(footer)
}

// Clone creates a deep copy of this HeaderFooter element.
func (hf *HeaderFooter) Clone() openxml.Element {
	cloned := hf.CompositeElementBase.Clone()

	return &HeaderFooter{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this HeaderFooter element.
func (hf *HeaderFooter) CloneNode(
	deep bool,
) openxml.Element {
	cloned := hf.CompositeElementBase.CloneNode(
		deep,
	)

	return &HeaderFooter{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
