// Package diagram provides types and functions for working with SmartArt diagrams
// in Office Open XML documents.
package diagram

import (
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// AddParagraph adds a new paragraph with the given text to the diagram TextBody.
func (tb *TextBody) AddParagraph(text string) *drawingml.TextParagraph {
	p := drawingml.NewTextParagraph()
	if text != "" {
		r := drawingml.NewTextRun(text)
		p.AppendChild(r)
	}
	tb.AppendChild(p)

	return p
}

// Paragraphs returns all paragraphs in the text body.
func (tb *TextBody) Paragraphs() []*drawingml.TextParagraph {
	var paragraphs []*drawingml.TextParagraph
	for child := range tb.Children() {
		if child.LocalName() != "p" ||
			child.NamespaceURI() != drawingml.NamespaceMain {
			continue
		}

		switch v := child.(type) {
		case *drawingml.TextParagraph:
			paragraphs = append(paragraphs, v)
		case *openxml.CompositeElementBase:
			paragraphs = append(paragraphs, &drawingml.TextParagraph{CompositeElementBase: v})
		}
	}

	return paragraphs
}
