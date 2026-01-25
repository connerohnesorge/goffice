// and effects.
// DrawingML (Drawing Markup Language) is the common drawing specification used
// across Office Open XML documents for charts, diagrams, and graphical content.
package drawingml

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// wrapCompositeElement attempts to wrap an element as a CompositeElementBase.
func wrapCompositeElement(
	elem openxml.Element,
) *openxml.CompositeElementBase {
	if elem == nil {
		return nil
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return comp
	}
	return nil
}

// wrapLeafElement attempts to wrap an element as a LeafElementBase.
func wrapLeafElement(
	elem openxml.Element,
) *openxml.LeafElementBase {
	if elem == nil {
		return nil
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return leaf
	}
	return nil
}
