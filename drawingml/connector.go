package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Connection represents a connection to a shape (a:stCxn or a:endCxn).
type Connection struct {
	*openxml.LeafElementBase
}

// NewConnection creates a new connection element.
// name is either "stCxn" or "endCxn".
func NewConnection(name, id string, idx int) *Connection {
	c := &Connection{
		LeafElementBase: openxml.NewLeafElement(NamespaceMain, name, PrefixMain),
	}
	c.SetId(id)
	c.SetIndex(idx)

	return c
}

// Id returns the shape ID.
func (c *Connection) Id() string {
	attr, found := c.GetAttribute("id", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetId sets the shape ID.
func (c *Connection) SetId(id string) {
	c.SetAttribute(openxml.NewAttribute("", "id", "", id))
}

// Index returns the connection site index.
func (c *Connection) Index() int {
	attr, found := c.GetAttribute("idx", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetIndex sets the connection site index.
func (c *Connection) SetIndex(idx int) {
	c.SetAttribute(openxml.NewAttribute("", "idx", "", strconv.Itoa(idx)))
}

// ConnectorLocks represents connector locking properties (a:cxnSpLocks).
type ConnectorLocks struct {
	*openxml.LeafElementBase
}

// NewConnectorLocks creates a new ConnectorLocks element.
func NewConnectorLocks() *ConnectorLocks {
	return &ConnectorLocks{
		LeafElementBase: openxml.NewLeafElement(
			NamespaceMain,
			"cxnSpLocks",
			PrefixMain,
		),
	}
}

// NoGrouping returns whether grouping is locked.
func (s *ConnectorLocks) NoGrouping() bool {
	return s.getBoolAttr("noGrp", false)
}

// SetNoGrouping sets whether grouping is locked.
func (s *ConnectorLocks) SetNoGrouping(lock bool) {
	s.setBoolAttr("noGrp", lock, false)
}

// NoSelection returns whether selection is locked.
func (s *ConnectorLocks) NoSelection() bool {
	return s.getBoolAttr("noSelect", false)
}

// SetNoSelection sets whether selection is locked.
func (s *ConnectorLocks) SetNoSelection(lock bool) {
	s.setBoolAttr("noSelect", lock, false)
}

// NoRotation returns whether rotation is locked.
func (s *ConnectorLocks) NoRotation() bool {
	return s.getBoolAttr("noRot", false)
}

// SetNoRotation sets whether rotation is locked.
func (s *ConnectorLocks) SetNoRotation(lock bool) {
	s.setBoolAttr("noRot", lock, false)
}

// NoChangeAspect returns whether aspect ratio change is locked.
func (s *ConnectorLocks) NoChangeAspect() bool {
	return s.getBoolAttr("noChangeAspect", false)
}

// SetNoChangeAspect sets whether aspect ratio change is locked.
func (s *ConnectorLocks) SetNoChangeAspect(lock bool) {
	s.setBoolAttr("noChangeAspect", lock, false)
}

// NoMove returns whether moving is locked.
func (s *ConnectorLocks) NoMove() bool {
	return s.getBoolAttr("noMove", false)
}

// SetNoMove sets whether moving is locked.
func (s *ConnectorLocks) SetNoMove(lock bool) {
	s.setBoolAttr("noMove", lock, false)
}

// NoResize returns whether resizing is locked.
func (s *ConnectorLocks) NoResize() bool {
	return s.getBoolAttr("noResize", false)
}

// SetNoResize sets whether resizing is locked.
func (s *ConnectorLocks) SetNoResize(lock bool) {
	s.setBoolAttr("noResize", lock, false)
}

// NoEditPoints returns whether editing points is locked.
func (s *ConnectorLocks) NoEditPoints() bool {
	return s.getBoolAttr("noEditPoints", false)
}

// SetNoEditPoints sets whether editing points is locked.
func (s *ConnectorLocks) SetNoEditPoints(lock bool) {
	s.setBoolAttr("noEditPoints", lock, false)
}

// NoAdjustHandles returns whether adjusting handles is locked.
func (s *ConnectorLocks) NoAdjustHandles() bool {
	return s.getBoolAttr("noAdjustHandles", false)
}

// SetNoAdjustHandles sets whether adjusting handles is locked.
func (s *ConnectorLocks) SetNoAdjustHandles(lock bool) {
	s.setBoolAttr("noAdjustHandles", lock, false)
}

// NoChangeArrowheads returns whether changing arrowheads is locked.
func (s *ConnectorLocks) NoChangeArrowheads() bool {
	return s.getBoolAttr("noChangeArrowheads", false)
}

// SetNoChangeArrowheads sets whether changing arrowheads is locked.
func (s *ConnectorLocks) SetNoChangeArrowheads(lock bool) {
	s.setBoolAttr("noChangeArrowheads", lock, false)
}

// NoChangeShapeType returns whether changing shape type is locked.
func (s *ConnectorLocks) NoChangeShapeType() bool {
	return s.getBoolAttr("noChangeShapeType", false)
}

// SetNoChangeShapeType sets whether changing shape type is locked.
func (s *ConnectorLocks) SetNoChangeShapeType(lock bool) {
	s.setBoolAttr("noChangeShapeType", lock, false)
}

// getBoolAttr retrieves a boolean attribute value.
func (s *ConnectorLocks) getBoolAttr(name string, defaultVal bool) bool {
	attr, found := s.GetAttribute(name, "")
	if !found {
		return defaultVal
	}
	val := attr.Value()

	return val == "1" || val == "true"
}

// setBoolAttr sets a boolean attribute value.
func (s *ConnectorLocks) setBoolAttr(name string, value, defaultVal bool) {
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

// Clone creates a deep copy of this ConnectorLocks element.
func (s *ConnectorLocks) Clone() openxml.Element {
	return &ConnectorLocks{
		LeafElementBase: s.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}
