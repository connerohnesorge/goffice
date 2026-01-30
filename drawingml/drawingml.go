// Package drawingml provides shared DrawingML types for shapes, images, charts,
// diagrams, and effects.
//
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
