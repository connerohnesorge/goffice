package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

const (
	// NamespaceWML is the WordprocessingML namespace URI.
	NamespaceWML = openxml.NamespaceWordprocessingML
	// PrefixW is the standard prefix for WordprocessingML elements.
	PrefixW = "w"
	// NamespaceXML is the XML namespace URI.
	NamespaceXML = "http://www.w3.org/XML/1998/namespace"
	// NamespaceW15 is the Word 2013 extension namespace (w15).
	NamespaceW15 = openxml.NamespaceWord2013
	// PrefixW15 is the standard prefix for Word 2013 extension elements.
	PrefixW15 = "w15"
)

// Text represents a text content element (w:t).
type Text struct {
	*openxml.LeafElementBase
}

// NewText creates a new Text element with the specified content.
func NewText(text string) *Text {
	elem := openxml.NewLeafElementWithText(
		NamespaceWML,
		"t",
		PrefixW,
		text,
	)
	t := &Text{LeafElementBase: elem}
	// Preserve whitespace if needed
	if needsSpacePreserve(text) {
		t.SetSpace("preserve")
	}

	return t
}

// needsSpacePreserve checks if the text needs xml:space="preserve".
func needsSpacePreserve(text string) bool {
	if text == "" {
		return false
	}
	// Check for leading or trailing whitespace
	if text[0] == ' ' || text[0] == '\t' ||
		text[len(text)-1] == ' ' ||
		text[len(text)-1] == '\t' {
		return true
	}
	// Check for multiple consecutive spaces
	for i := range len(text) - 1 {
		if text[i] == ' ' && text[i+1] == ' ' {
			return true
		}
	}

	return false
}

// SetText sets the text content.
func (t *Text) SetText(value string) {
	t.SetInnerText(value)
	if needsSpacePreserve(value) {
		t.SetSpace("preserve")
	} else {
		t.RemoveAttribute("space", NamespaceXML)
	}
}

// Space returns the xml:space attribute value.
func (t *Text) Space() string {
	attr, found := t.GetAttribute(
		"space",
		NamespaceXML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSpace sets the xml:space attribute.
func (t *Text) SetSpace(value string) {
	attr := openxml.NewAttribute(
		NamespaceXML,
		"space",
		"xml",
		value,
	)
	t.SetAttribute(attr)
}

// Clone creates a deep copy of this Text element.
func (t *Text) Clone() openxml.Element {
	return &Text{
		LeafElementBase: t.LeafElementBase.Clone().(*openxml.LeafElementBase),
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
