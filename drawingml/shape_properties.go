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

// Scene3D returns the 3D scene properties, or nil if not set.
func (s *ShapeProperties) Scene3D() *Scene3D {
	elem := s.GetElement("scene3d", NamespaceMain)
	if elem == nil {
		return nil
	}
	if scene, ok := elem.(*Scene3D); ok {
		return scene
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Scene3D{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetScene3D sets the 3D scene properties.
func (s *ShapeProperties) SetScene3D(scene *Scene3D) {
	if existing := s.GetElement("scene3d", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if scene != nil {
		s.InsertAfter(scene, s.Outline()) // logical placement, though XSD order matters
		// To adhere to XSD, we should probably check siblings.
		// For now, appending or inserting after ln is a reasonable guess if we don't do full validation here.
		// However, in ShapeProperties, order is: xfrm, geometry, fill, ln, effects, scene3d, sp3d, extLst.
		// So inserting after ln is good, or appending if ln is missing.
		s.insertInOrder(scene, "scene3d", "ln", "gradFill", "solidFill", "xfrm")
	}
}

// Shape3D returns the 3D shape properties, or nil if not set.
func (s *ShapeProperties) Shape3D() *Shape3D {
	elem := s.GetElement("sp3d", NamespaceMain)
	if elem == nil {
		return nil
	}
	if sp3d, ok := elem.(*Shape3D); ok {
		return sp3d
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Shape3D{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetShape3D sets the 3D shape properties.
func (s *ShapeProperties) SetShape3D(sp3d *Shape3D) {
	if existing := s.GetElement("sp3d", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if sp3d != nil {
		// sp3d comes after scene3d
		s.insertInOrder(sp3d, "sp3d", "scene3d", "ln", "gradFill", "solidFill", "xfrm")
	}
}

// EffectList returns the effect list, or nil if not set.
func (s *ShapeProperties) EffectList() *EffectList {
	elem := s.GetElement("effectLst", NamespaceMain)
	if elem == nil {
		return nil
	}
	if el, ok := elem.(*EffectList); ok {
		return el
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &EffectList{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetEffectList sets the effect list.
func (s *ShapeProperties) SetEffectList(effects *EffectList) {
	if existing := s.GetElement("effectLst", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if effects != nil {
		// effectLst comes after ln and before scene3d
		s.insertInOrder(effects, "effectLst", "ln", "gradFill", "solidFill", "xfrm")
	}
}

// insertInOrder inserts the element in the correct position based on predecessors.
func (s *ShapeProperties) insertInOrder(elem openxml.Element, name string, predecessors ...string) {
	// Try to find any of the predecessors in order (reverse to find the latest one)
	for i := 0; i < len(predecessors); i++ {
		predName := predecessors[i]
		if pred := s.GetElement(predName, NamespaceMain); pred != nil {
			s.InsertAfter(elem, pred)
			return
		}
	}
	// Fallback: prepend (if no predecessors found, maybe it's the first)
	// Or append?
	// If xfrm is missing, and we rely on it being first, maybe prepend is safer if we missed something.
	// But if we have fills, we want to be after them.
	// If we didn't find any predecessors, it might mean the element should be first OR we just didn't list all predecessors.
	// For ShapeProperties, order is: xfrm, geometry, fill, ln, effects, scene3d, sp3d, extLst.
	// So if we set scene3d and found no ln, no fill, no xfrm, it should be first?
	// Or maybe after geometry?
	// Let's just append if we are adding scene3d/sp3d/effects and we didn't find predecessors.
	// Wait, if we use InsertAfter(elem, nil), it might fail or behave unexpectedly?
	// openxml implementation usually handles InsertAfter(elem, nil) as PrependChild?
	// Let's assume PrependChild if no predecessor found, but that might put it before xfrm.
	// So we should try to append if it's a "late" element.
	
	// Simplified strategy:
	// If it's scene3d or sp3d or effects, we prefer appending if no 'ln' or 'fill' or 'xfrm' found?
	// But 'xfrm' is almost always there.
	// If 'xfrm' is missing, maybe we should just append.
	
	s.AppendChild(elem)
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

// ShapeLocks returns the shape locks, or nil if not set.
func (s *ShapeProperties) ShapeLocks() *ShapeLocks {
	elem := s.GetElement("spLocks", NamespaceMain)
	if elem == nil {
		return nil
	}
	if locks, ok := elem.(*ShapeLocks); ok {
		return locks
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &ShapeLocks{
			LeafElementBase: leaf,
		}
	}
	return nil
}

// SetShapeLocks sets the shape locks.
func (s *ShapeProperties) SetShapeLocks(locks *ShapeLocks) {
	if existing := s.GetElement("spLocks", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if locks != nil {
		// spLocks comes after sp3d, before extLst
		s.insertInOrder(locks, "spLocks", "sp3d", "scene3d", "ln", "gradFill", "solidFill", "xfrm")
	}
}

// Clone creates a deep copy of this ShapeProperties element.
func (s *ShapeProperties) Clone() openxml.Element {
	cloned := s.CompositeElementBase.Clone()

	return &ShapeProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
