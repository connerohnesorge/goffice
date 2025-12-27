// Package drawingml provides shared DrawingML types for shapes, images,
// and effects.
package drawingml

import "github.com/connerohnesorge/goffice/openxml"

// BlackWhiteMode represents black and white mode values for rendering.
type BlackWhiteMode string

// Black and white mode values.
const (
	BWModeColor      BlackWhiteMode = "clr"
	BWModeAuto       BlackWhiteMode = "auto"
	BWModeGray       BlackWhiteMode = "gray"
	BWModeLightGray  BlackWhiteMode = "ltGray"
	BWModeInvGray    BlackWhiteMode = "invGray"
	BWModeGrayWhite  BlackWhiteMode = "grayWhite"
	BWModeBlackGray  BlackWhiteMode = "blackGray"
	BWModeBlackWhite BlackWhiteMode = "blackWhite"
	BWModeBlack      BlackWhiteMode = "black"
	BWModeWhite      BlackWhiteMode = "white"
	BWModeHidden     BlackWhiteMode = "hidden"
)

// ShapeProperties represents the shape properties element (a:spPr).
// This element specifies the visual properties of a shape, including
// transform, fill, line, and effects.
type ShapeProperties struct {
	*openxml.CompositeElementBase
}

// NewShapeProperties creates a new shape properties element.
func NewShapeProperties() *ShapeProperties {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"spPr",
		PrefixMain,
	)

	return &ShapeProperties{
		CompositeElementBase: elem,
	}
}

// BlackWhiteMode returns the black and white mode setting.
func (s *ShapeProperties) BlackWhiteMode() BlackWhiteMode {
	attr, found := s.GetAttribute("bwMode", "")
	if !found {
		return ""
	}

	return BlackWhiteMode(attr.Value())
}

// SetBlackWhiteMode sets the black and white mode setting.
func (s *ShapeProperties) SetBlackWhiteMode(
	mode BlackWhiteMode,
) {
	if mode == "" {
		s.RemoveAttribute("bwMode", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"bwMode",
			"",
			string(mode),
		),
	)
}

// Transform returns the 2D transform element, or nil if not set.
func (s *ShapeProperties) Transform() *Transform2D {
	elem := s.GetElement("xfrm", NamespaceMain)
	if elem == nil {
		return nil
	}
	if xfrm, ok := elem.(*Transform2D); ok {
		return xfrm
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Transform2D{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetTransform sets the 2D transform for the shape.
func (s *ShapeProperties) SetTransform(
	transform *Transform2D,
) {
	// Remove existing transform
	if existing := s.GetElement("xfrm", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if transform != nil {
		// Transform should be first child
		s.PrependChild(transform)
	}
}

// SetTransformValues sets the transform using offset and extent values.
func (s *ShapeProperties) SetTransformValues(
	offX, offY, extCx, extCy EMU,
) {
	transform := NewTransform2D(
		offX,
		offY,
		extCx,
		extCy,
	)
	s.SetTransform(transform)
}

// PresetGeometry returns the preset geometry, or nil if not set.
func (s *ShapeProperties) PresetGeometry() *PresetGeometry {
	elem := s.GetElement(
		"prstGeom",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if pg, ok := elem.(*PresetGeometry); ok {
		return pg
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PresetGeometry{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetPresetGeometry sets the preset geometry for the shape.
func (s *ShapeProperties) SetPresetGeometry(
	geometry *PresetGeometry,
) {
	s.removeGeometry()
	if geometry != nil {
		s.insertAfterTransform(geometry)
	}
}

// SetPresetShape sets the shape to a preset shape type.
func (s *ShapeProperties) SetPresetShape(
	shapeType ShapeTypeValue,
) {
	geometry := NewPresetGeometry(shapeType)
	s.SetPresetGeometry(geometry)
}

// CustomGeometry returns the custom geometry, or nil if not set.
func (s *ShapeProperties) CustomGeometry() *CustomGeometry {
	elem := s.GetElement(
		"custGeom",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if cg, ok := elem.(*CustomGeometry); ok {
		return cg
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CustomGeometry{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetCustomGeometry sets the custom geometry for the shape.
func (s *ShapeProperties) SetCustomGeometry(
	geometry *CustomGeometry,
) {
	s.removeGeometry()
	if geometry != nil {
		s.insertAfterTransform(geometry)
	}
}

// removeGeometry removes any existing geometry element.
func (s *ShapeProperties) removeGeometry() {
	if prstGeom := s.GetElement("prstGeom", NamespaceMain); prstGeom != nil {
		s.RemoveChild(prstGeom)
	}
	if custGeom := s.GetElement("custGeom", NamespaceMain); custGeom != nil {
		s.RemoveChild(custGeom)
	}
}

// insertAfterTransform inserts an element after the transform element.
func (s *ShapeProperties) insertAfterTransform(
	elem openxml.Element,
) {
	xfrm := s.GetElement("xfrm", NamespaceMain)
	if xfrm != nil {
		s.InsertAfter(elem, xfrm)
	} else {
		s.PrependChild(elem)
	}
}

// Outline returns the outline (line) properties, or nil if not set.
func (s *ShapeProperties) Outline() *LineProperties {
	elem := s.GetElement("ln", NamespaceMain)
	if elem == nil {
		return nil
	}
	if lp, ok := elem.(*LineProperties); ok {
		return lp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &LineProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetOutline sets the outline (line) properties.
func (s *ShapeProperties) SetOutline(
	outline *LineProperties,
) {
	// Remove existing outline
	if existing := s.GetElement("ln", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if outline != nil {
		s.insertAfterFill(outline)
	}
}

// SetOutlineColor sets a simple solid color outline.
func (s *ShapeProperties) SetOutlineColor(
	hexColor string,
	width EMU,
) {
	outline := NewLinePropertiesWithWidth(width)
	outline.SetSolidFill(hexColor)
	s.SetOutline(outline)
}

// SetOutlineSchemeColor sets a simple solid scheme color outline.
func (s *ShapeProperties) SetOutlineSchemeColor(
	color SchemeColorValue,
	width EMU,
) {
	outline := NewLinePropertiesWithWidth(width)
	outline.SetSolidFillSchemeColor(color)
	s.SetOutline(outline)
}

// SetNoOutline removes the outline from the shape.
func (s *ShapeProperties) SetNoOutline() {
	outline := NewLineProperties()
	outline.SetNoFill()
	s.SetOutline(outline)
}

// insertAfterFill inserts an element after fill elements.
func (s *ShapeProperties) insertAfterFill(
	elem openxml.Element,
) {
	// Check for fill elements in reverse order
	fillElements := []string{
		"grpFill",
		"blipFill",
		"pattFill",
		"gradFill",
		"solidFill",
		"noFill",
	}
	for _, name := range fillElements {
		if fill := s.GetElement(name, NamespaceMain); fill != nil {
			s.InsertAfter(elem, fill)

			return
		}
	}
	// No fill found, insert after geometry
	s.insertAfterGeometry(elem)
}

// Clone creates a deep copy of this ShapeProperties element.
func (s *ShapeProperties) Clone() openxml.Element {
	cloned := s.CompositeElementBase.Clone()

	return &ShapeProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
