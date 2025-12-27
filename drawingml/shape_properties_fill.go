// Package drawingml provides shared DrawingML types for shapes, images,
// and effects.
package drawingml

import "github.com/connerohnesorge/goffice/openxml"

// NoFill returns true if no fill is set.
func (s *ShapeProperties) NoFill() bool {
	return s.GetElement(
		"noFill",
		NamespaceMain,
	) != nil
}

// SetNoFill sets the shape to have no fill.
func (s *ShapeProperties) SetNoFill() {
	s.removeFill()
	noFill := NewNoFill()
	s.insertAfterGeometry(noFill)
}

// SolidFill returns the solid fill, or nil if not set.
func (s *ShapeProperties) SolidFill() *SolidFill {
	elem := s.GetElement(
		"solidFill",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if sf, ok := elem.(*SolidFill); ok {
		return sf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SolidFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetSolidFill sets the shape to have a solid color fill.
func (s *ShapeProperties) SetSolidFill(
	fill *SolidFill,
) {
	s.removeFill()
	if fill != nil {
		s.insertAfterGeometry(fill)
	}
}

// SetSolidFillColor sets the shape to have a solid RGB color fill.
func (s *ShapeProperties) SetSolidFillColor(
	hexColor string,
) {
	fill := NewSolidFillWithRgb(hexColor)
	s.SetSolidFill(fill)
}

// SetSolidFillSchemeColor sets the shape to have a solid scheme color fill.
func (s *ShapeProperties) SetSolidFillSchemeColor(
	color SchemeColorValue,
) {
	fill := NewSolidFillWithSchemeColor(color)
	s.SetSolidFill(fill)
}

// GradientFill returns the gradient fill, or nil if not set.
func (s *ShapeProperties) GradientFill() *GradientFill {
	elem := s.GetElement(
		"gradFill",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if gf, ok := elem.(*GradientFill); ok {
		return gf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &GradientFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetGradientFill sets the shape to have a gradient fill.
func (s *ShapeProperties) SetGradientFill(
	fill *GradientFill,
) {
	s.removeFill()
	if fill != nil {
		s.insertAfterGeometry(fill)
	}
}

// PatternFill returns the pattern fill, or nil if not set.
func (s *ShapeProperties) PatternFill() *PatternFill {
	elem := s.GetElement(
		"pattFill",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if pf, ok := elem.(*PatternFill); ok {
		return pf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PatternFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetPatternFill sets the shape to have a pattern fill.
func (s *ShapeProperties) SetPatternFill(
	fill *PatternFill,
) {
	s.removeFill()
	if fill != nil {
		s.insertAfterGeometry(fill)
	}
}

// BlipFill returns the blip (image) fill, or nil if not set.
func (s *ShapeProperties) BlipFill() *BlipFill {
	elem := s.GetElement(
		"blipFill",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if bf, ok := elem.(*BlipFill); ok {
		return bf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &BlipFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetBlipFill sets the shape to have a blip (image) fill.
func (s *ShapeProperties) SetBlipFill(
	fill *BlipFill,
) {
	s.removeFill()
	if fill != nil {
		s.insertAfterGeometry(fill)
	}
}

// GroupFill returns true if group fill is set.
func (s *ShapeProperties) GroupFill() bool {
	return s.GetElement(
		"grpFill",
		NamespaceMain,
	) != nil
}

// SetGroupFill sets the shape to inherit fill from its group.
func (s *ShapeProperties) SetGroupFill() {
	s.removeFill()
	grpFill := NewGroupFill()
	s.insertAfterGeometry(grpFill)
}

// removeFill removes any existing fill element.
func (s *ShapeProperties) removeFill() {
	if noFill := s.GetElement("noFill", NamespaceMain); noFill != nil {
		s.RemoveChild(noFill)
	}
	if solidFill := s.GetElement("solidFill", NamespaceMain); solidFill != nil {
		s.RemoveChild(solidFill)
	}
	if gradFill := s.GetElement("gradFill", NamespaceMain); gradFill != nil {
		s.RemoveChild(gradFill)
	}
	if pattFill := s.GetElement("pattFill", NamespaceMain); pattFill != nil {
		s.RemoveChild(pattFill)
	}
	if blipFill := s.GetElement("blipFill", NamespaceMain); blipFill != nil {
		s.RemoveChild(blipFill)
	}
	if grpFill := s.GetElement("grpFill", NamespaceMain); grpFill != nil {
		s.RemoveChild(grpFill)
	}
}

// insertAfterGeometry inserts an element after the geometry elements.
func (s *ShapeProperties) insertAfterGeometry(
	elem openxml.Element,
) {
	// Check for geometry elements in order
	if custGeom := s.GetElement("custGeom", NamespaceMain); custGeom != nil {
		s.InsertAfter(elem, custGeom)

		return
	}
	if prstGeom := s.GetElement("prstGeom", NamespaceMain); prstGeom != nil {
		s.InsertAfter(elem, prstGeom)

		return
	}
	// Check for transform
	if xfrm := s.GetElement("xfrm", NamespaceMain); xfrm != nil {
		s.InsertAfter(elem, xfrm)

		return
	}
	s.PrependChild(elem)
}
