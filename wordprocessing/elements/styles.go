package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// Styles represents the root element of the styles part (w:styles).
type Styles struct {
	*openxml.CompositeElementBase
}

// NewStyles creates a new Styles element.
func NewStyles() *Styles {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"styles",
		PrefixW,
	)

	return &Styles{CompositeElementBase: elem}
}

// DocDefaults returns the document defaults element, or nil if not present.
func (s *Styles) DocDefaults() *DocDefaults {
	elem := s.GetElement(
		"docDefaults",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if dd, ok := elem.(*DocDefaults); ok {
		return dd
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DocDefaults{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDocDefaults returns the document defaults, creating if needed.
func (s *Styles) GetOrCreateDocDefaults() *DocDefaults {
	dd := s.DocDefaults()
	if dd != nil {
		return dd
	}
	dd = NewDocDefaults()
	// DocDefaults should be first child
	if first := s.FirstChild(); first != nil {
		s.InsertBefore(dd, first)
	} else {
		s.AppendChild(dd)
	}

	return dd
}

// LatentStyles returns the latent styles element, or nil if not present.
func (s *Styles) LatentStyles() *LatentStyles {
	elem := s.GetElement(
		"latentStyles",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if ls, ok := elem.(*LatentStyles); ok {
		return ls
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &LatentStyles{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateLatentStyles returns the latent styles, creating if needed.
func (s *Styles) GetOrCreateLatentStyles() *LatentStyles {
	ls := s.LatentStyles()
	if ls != nil {
		return ls
	}
	ls = NewLatentStyles()
	// LatentStyles comes after DocDefaults
	dd := s.DocDefaults()
	if dd != nil {
		s.InsertAfter(ls, dd)
	} else if first := s.FirstChild(); first != nil {
		s.InsertBefore(ls, first)
	} else {
		s.AppendChild(ls)
	}

	return ls
}

// Styles returns an iterator over all Style elements.
func (s *Styles) Styles() iter.Seq[*Style] {
	return func(yield func(*Style) bool) {
		for child := range s.Children() {
			if child.LocalName() == "style" &&
				child.NamespaceURI() == NamespaceWML {
				var style *Style
				if st, ok := child.(*Style); ok {
					style = st
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					style = &Style{CompositeElementBase: comp}
				}
				if style != nil && !yield(style) {
					return
				}
			}
		}
	}
}

// GetStyleById returns the style with the given ID, or nil if not found.
func (s *Styles) GetStyleById(id string) *Style {
	for style := range s.Styles() {
		if style.StyleId() == id {
			return style
		}
	}

	return nil
}

// GetStyleByName returns the style with the given name, or nil if not found.
func (s *Styles) GetStyleByName(
	name string,
) *Style {
	for style := range s.Styles() {
		if style.StyleName() == name {
			return style
		}
	}

	return nil
}

// GetStylesByType returns an iterator over styles of the specified type.
func (s *Styles) GetStylesByType(
	styleType StyleType,
) iter.Seq[*Style] {
	return func(yield func(*Style) bool) {
		for style := range s.Styles() {
			if style.Type() == styleType {
				if !yield(style) {
					return
				}
			}
		}
	}
}

// AddStyle adds a style to the styles collection.
func (s *Styles) AddStyle(style *Style) {
	s.AppendChild(style)
}

// RemoveStyle removes a style from the styles collection.
func (s *Styles) RemoveStyle(style *Style) bool {
	return s.RemoveChild(style)
}

// Clone creates a deep copy of this Styles element.
func (s *Styles) Clone() openxml.Element {
	return &Styles{
		CompositeElementBase: s.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Styles element.
func (s *Styles) CloneNode(
	deep bool,
) openxml.Element {
	return &Styles{
		CompositeElementBase: s.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}
