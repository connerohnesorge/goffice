package drawingml

import "github.com/connerohnesorge/goffice/openxml"

// ShapeLocks represents shape locking properties (a:spLocks).
// This element specifies locking properties for a shape to prevent
// certain types of modifications by the user.
type ShapeLocks struct {
	*openxml.LeafElementBase
}

// NewShapeLocks creates a new ShapeLocks element.
func NewShapeLocks() *ShapeLocks {
	return &ShapeLocks{
		LeafElementBase: openxml.NewLeafElement(
			NamespaceMain,
			"spLocks",
			PrefixMain,
		),
	}
}

// NoGrouping returns whether grouping is locked.
func (s *ShapeLocks) NoGrouping() bool {
	return s.getBoolAttr("noGrp", false)
}

// SetNoGrouping sets whether grouping is locked.
func (s *ShapeLocks) SetNoGrouping(lock bool) {
	s.setBoolAttr("noGrp", lock, false)
}

// NoSelection returns whether selection is locked.
func (s *ShapeLocks) NoSelection() bool {
	return s.getBoolAttr("noSelect", false)
}

// SetNoSelection sets whether selection is locked.
func (s *ShapeLocks) SetNoSelection(lock bool) {
	s.setBoolAttr("noSelect", lock, false)
}

// NoRotation returns whether rotation is locked.
func (s *ShapeLocks) NoRotation() bool {
	return s.getBoolAttr("noRot", false)
}

// SetNoRotation sets whether rotation is locked.
func (s *ShapeLocks) SetNoRotation(lock bool) {
	s.setBoolAttr("noRot", lock, false)
}

// NoChangeAspect returns whether aspect ratio change is locked.
func (s *ShapeLocks) NoChangeAspect() bool {
	return s.getBoolAttr("noChangeAspect", false)
}

// SetNoChangeAspect sets whether aspect ratio change is locked.
func (s *ShapeLocks) SetNoChangeAspect(lock bool) {
	s.setBoolAttr("noChangeAspect", lock, false)
}

// NoMove returns whether moving is locked.
func (s *ShapeLocks) NoMove() bool {
	return s.getBoolAttr("noMove", false)
}

// SetNoMove sets whether moving is locked.
func (s *ShapeLocks) SetNoMove(lock bool) {
	s.setBoolAttr("noMove", lock, false)
}

// NoResize returns whether resizing is locked.
func (s *ShapeLocks) NoResize() bool {
	return s.getBoolAttr("noResize", false)
}

// SetNoResize sets whether resizing is locked.
func (s *ShapeLocks) SetNoResize(lock bool) {
	s.setBoolAttr("noResize", lock, false)
}

// NoEditPoints returns whether editing points is locked.
func (s *ShapeLocks) NoEditPoints() bool {
	return s.getBoolAttr("noEditPoints", false)
}

// SetNoEditPoints sets whether editing points is locked.
func (s *ShapeLocks) SetNoEditPoints(lock bool) {
	s.setBoolAttr("noEditPoints", lock, false)
}

// NoAdjustHandles returns whether adjusting handles is locked.
func (s *ShapeLocks) NoAdjustHandles() bool {
	return s.getBoolAttr("noAdjustHandles", false)
}

// SetNoAdjustHandles sets whether adjusting handles is locked.
func (s *ShapeLocks) SetNoAdjustHandles(lock bool) {
	s.setBoolAttr("noAdjustHandles", lock, false)
}

// NoChangeArrowheads returns whether changing arrowheads is locked.
func (s *ShapeLocks) NoChangeArrowheads() bool {
	return s.getBoolAttr("noChangeArrowheads", false)
}

// SetNoChangeArrowheads sets whether changing arrowheads is locked.
func (s *ShapeLocks) SetNoChangeArrowheads(lock bool) {
	s.setBoolAttr("noChangeArrowheads", lock, false)
}

// NoChangeShapeType returns whether changing shape type is locked.
func (s *ShapeLocks) NoChangeShapeType() bool {
	return s.getBoolAttr("noChangeShapeType", false)
}

// SetNoChangeShapeType sets whether changing shape type is locked.
func (s *ShapeLocks) SetNoChangeShapeType(lock bool) {
	s.setBoolAttr("noChangeShapeType", lock, false)
}

// NoTextEdit returns whether text editing is locked.
func (s *ShapeLocks) NoTextEdit() bool {
	return s.getBoolAttr("noTextEdit", false)
}

// SetNoTextEdit sets whether text editing is locked.
func (s *ShapeLocks) SetNoTextEdit(lock bool) {
	s.setBoolAttr("noTextEdit", lock, false)
}

// getBoolAttr retrieves a boolean attribute value.
func (s *ShapeLocks) getBoolAttr(name string, defaultVal bool) bool {
	attr, found := s.GetAttribute(name, "")
	if !found {
		return defaultVal
	}
	val := attr.Value()
	return val == "1" || val == "true"
}

// setBoolAttr sets a boolean attribute value.
func (s *ShapeLocks) setBoolAttr(name string, value, defaultVal bool) {
	if value == defaultVal {
		s.RemoveAttribute(name, "")
		return
	}
	var strVal string
	if value {
		strVal = "1"
	} else {
		strVal = "0"
	}
	s.SetAttribute(openxml.NewAttribute("", name, "", strVal))
}

// Clone creates a deep copy of this ShapeLocks element.
func (s *ShapeLocks) Clone() openxml.Element {
	return &ShapeLocks{
		LeafElementBase: s.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}
