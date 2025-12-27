//nolint:revive // comments-density
package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// OneCellAnchor represents a one-cell anchor element (xdr:oneCellAnchor).
// The drawing is anchored to one cell with a fixed extent.
type OneCellAnchor struct {
	*openxml.CompositeElementBase
}

// NewOneCellAnchor creates a new OneCellAnchor element.
func NewOneCellAnchor() *OneCellAnchor {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"oneCellAnchor",
		PrefixXDR,
	)

	return &OneCellAnchor{
		CompositeElementBase: elem,
	}
}

// From returns the From marker element, or nil if not present.
func (oc *OneCellAnchor) From() *FromMarker {
	elem := oc.GetElement(
		"from",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if fm, ok := elem.(*FromMarker); ok {
		return fm
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &FromMarker{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateFrom returns the From marker, creating if needed.
func (oc *OneCellAnchor) GetOrCreateFrom() *FromMarker {
	fm := oc.From()
	if fm != nil {
		return fm
	}
	fm = NewFromMarker()
	if first := oc.FirstChild(); first != nil {
		oc.InsertBefore(fm, first)
	} else {
		oc.AppendChild(fm)
	}

	return fm
}

// Ext returns the Ext (extent) element, or nil if not present.
func (oc *OneCellAnchor) Ext() *Extent {
	elem := oc.GetElement(
		"ext",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if ext, ok := elem.(*Extent); ok {
		return ext
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Extent{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateExt returns the Ext element, creating if needed.
func (oc *OneCellAnchor) GetOrCreateExt() *Extent {
	ext := oc.Ext()
	if ext != nil {
		return ext
	}
	ext = NewExtent()
	fm := oc.From()
	if fm != nil {
		oc.InsertAfter(ext, fm)
	} else {
		oc.AppendChild(ext)
	}

	return ext
}

// Shape returns the Shape element, or nil if not present.
func (oc *OneCellAnchor) Shape() *Shape {
	elem := oc.GetElement(
		"sp",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*Shape); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Shape{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateShape returns the Shape element, creating if needed.
func (oc *OneCellAnchor) GetOrCreateShape() *Shape {
	sp := oc.Shape()
	if sp != nil {
		return sp
	}
	sp = NewShape()
	ext := oc.Ext()
	if ext != nil {
		oc.InsertAfter(sp, ext)
	} else {
		oc.AppendChild(sp)
	}

	return sp
}

// Picture returns the Picture element, or nil if not present.
func (oc *OneCellAnchor) Picture() *DrawingPicture {
	elem := oc.GetElement(
		"pic",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if pic, ok := elem.(*DrawingPicture); ok {
		return pic
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DrawingPicture{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePicture returns the Picture element, creating if needed.
func (oc *OneCellAnchor) GetOrCreatePicture() *DrawingPicture {
	pic := oc.Picture()
	if pic != nil {
		return pic
	}
	pic = NewDrawingPicture()
	ext := oc.Ext()
	if ext != nil {
		oc.InsertAfter(pic, ext)
	} else {
		oc.AppendChild(pic)
	}

	return pic
}

// GraphicFrame returns the GraphicFrame element, or nil if not present.
func (oc *OneCellAnchor) GraphicFrame() *GraphicFrame {
	elem := oc.GetElement(
		"graphicFrame",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if gf, ok := elem.(*GraphicFrame); ok {
		return gf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &GraphicFrame{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateGraphicFrame returns the GraphicFrame element, creating if needed.
func (oc *OneCellAnchor) GetOrCreateGraphicFrame() *GraphicFrame {
	gf := oc.GraphicFrame()
	if gf != nil {
		return gf
	}
	gf = NewGraphicFrame()
	ext := oc.Ext()
	if ext != nil {
		oc.InsertAfter(gf, ext)
	} else {
		oc.AppendChild(gf)
	}

	return gf
}

// ClientData returns the ClientData element, or nil if not present.
func (oc *OneCellAnchor) ClientData() *ClientData {
	elem := oc.GetElement(
		"clientData",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cd, ok := elem.(*ClientData); ok {
		return cd
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ClientData{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateClientData returns the ClientData element, creating if needed.
func (oc *OneCellAnchor) GetOrCreateClientData() *ClientData {
	cd := oc.ClientData()
	if cd != nil {
		return cd
	}
	cd = NewClientData()
	oc.AppendChild(cd)

	return cd
}

// Clone creates a deep copy of this OneCellAnchor element.
func (oc *OneCellAnchor) Clone() openxml.Element {
	cloned := oc.CompositeElementBase.Clone()

	return &OneCellAnchor{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this OneCellAnchor element.
func (oc *OneCellAnchor) CloneNode(
	deep bool,
) openxml.Element {
	cloned := oc.CompositeElementBase.CloneNode(
		deep,
	)

	return &OneCellAnchor{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
