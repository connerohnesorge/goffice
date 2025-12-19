package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Break represents a break element (w:br).
type Break struct {
	*openxml.CompositeElementBase
}

// NewBreak creates a new Break element with the specified type.
// If breakType is empty or BreakLine, a line break is created.
func NewBreak(breakType BreakType) *Break {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"br",
		PrefixW,
	)
	br := &Break{CompositeElementBase: elem}
	if breakType != "" && breakType != BreakLine {
		br.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"type",
				PrefixW,
				string(breakType),
			),
		)
	}

	return br
}

// NewLineBreak creates a new line break element.
func NewLineBreak() *Break {
	return NewBreak(BreakLine)
}

// NewPageBreak creates a new page break element.
func NewPageBreak() *Break {
	return NewBreak(BreakPage)
}

// NewColumnBreak creates a new column break element.
func NewColumnBreak() *Break {
	return NewBreak(BreakColumn)
}

// Type returns the break type.
func (br *Break) Type() BreakType {
	attr, found := br.GetAttribute(
		"type",
		NamespaceWML,
	)
	if !found {
		return BreakLine // default
	}

	return BreakType(attr.Value())
}

// SetType sets the break type.
func (br *Break) SetType(t BreakType) {
	if t == "" || t == BreakLine {
		br.RemoveAttribute("type", NamespaceWML)
	} else {
		br.SetAttribute(openxml.NewAttribute(NamespaceWML, "type", PrefixW, string(t)))
	}
}

// Clear returns the clear attribute value for text wrapping breaks.
func (br *Break) Clear() string {
	attr, found := br.GetAttribute(
		"clear",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetClear sets the clear attribute for text wrapping breaks.
func (br *Break) SetClear(value string) {
	if value == "" {
		br.RemoveAttribute("clear", NamespaceWML)
	} else {
		br.SetAttribute(openxml.NewAttribute(NamespaceWML, "clear", PrefixW, value))
	}
}

// Clone creates a deep copy of this Break element.
func (br *Break) Clone() openxml.Element {
	return &Break{
		CompositeElementBase: br.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Break element.
func (br *Break) CloneNode(
	deep bool,
) openxml.Element {
	return &Break{
		CompositeElementBase: br.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}
