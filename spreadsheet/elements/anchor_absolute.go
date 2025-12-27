//nolint:revive // comments-density
package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// AbsoluteAnchor represents an absolute anchor element (xdr:absoluteAnchor).
// The drawing is at a fixed position and does not move or resize with cells.
type AbsoluteAnchor struct {
	*openxml.CompositeElementBase
}

// NewAbsoluteAnchor creates a new AbsoluteAnchor element.
func NewAbsoluteAnchor() *AbsoluteAnchor {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"absoluteAnchor",
		PrefixXDR,
	)

	return &AbsoluteAnchor{
		CompositeElementBase: elem,
	}
}

// Pos returns the Pos (position) element, or nil if not present.
func (aa *AbsoluteAnchor) Pos() *Position {
	elem := aa.GetElement(
		"pos",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if pos, ok := elem.(*Position); ok {
		return pos
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Position{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePos returns the Pos element, creating if needed.
func (aa *AbsoluteAnchor) GetOrCreatePos() *Position {
	pos := aa.Pos()
	if pos != nil {
		return pos
	}
	pos = NewPosition()
	if first := aa.FirstChild(); first != nil {
		aa.InsertBefore(pos, first)
	} else {
		aa.AppendChild(pos)
	}

	return pos
}

// Ext returns the Ext (extent) element, or nil if not present.
func (aa *AbsoluteAnchor) Ext() *Extent {
	elem := aa.GetElement(
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
func (aa *AbsoluteAnchor) GetOrCreateExt() *Extent {
	ext := aa.Ext()
	if ext != nil {
		return ext
	}
	ext = NewExtent()
	pos := aa.Pos()
	if pos != nil {
		aa.InsertAfter(ext, pos)
	} else {
		aa.AppendChild(ext)
	}

	return ext
}

// Shape returns the Shape element, or nil if not present.
func (aa *AbsoluteAnchor) Shape() *Shape {
	elem := aa.GetElement(
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
func (aa *AbsoluteAnchor) GetOrCreateShape() *Shape {
	sp := aa.Shape()
	if sp != nil {
		return sp
	}
	sp = NewShape()
	ext := aa.Ext()
	if ext != nil {
		aa.InsertAfter(sp, ext)
	} else {
		aa.AppendChild(sp)
	}

	return sp
}

// Picture returns the Picture element, or nil if not present.
func (aa *AbsoluteAnchor) Picture() *DrawingPicture {
	elem := aa.GetElement(
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
func (aa *AbsoluteAnchor) GetOrCreatePicture() *DrawingPicture {
	pic := aa.Picture()
	if pic != nil {
		return pic
	}
	pic = NewDrawingPicture()
	ext := aa.Ext()
	if ext != nil {
		aa.InsertAfter(pic, ext)
	} else {
		aa.AppendChild(pic)
	}

	return pic
}

// GraphicFrame returns the GraphicFrame element, or nil if not present.
func (aa *AbsoluteAnchor) GraphicFrame() *GraphicFrame {
	elem := aa.GetElement(
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
func (aa *AbsoluteAnchor) GetOrCreateGraphicFrame() *GraphicFrame {
	gf := aa.GraphicFrame()
	if gf != nil {
		return gf
	}
	gf = NewGraphicFrame()
	ext := aa.Ext()
	if ext != nil {
		aa.InsertAfter(gf, ext)
	} else {
		aa.AppendChild(gf)
	}

	return gf
}

// ClientData returns the ClientData element, or nil if not present.
func (aa *AbsoluteAnchor) ClientData() *ClientData {
	elem := aa.GetElement(
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
func (aa *AbsoluteAnchor) GetOrCreateClientData() *ClientData {
	cd := aa.ClientData()
	if cd != nil {
		return cd
	}
	cd = NewClientData()
	aa.AppendChild(cd)

	return cd
}

// Clone creates a deep copy of this AbsoluteAnchor element.
func (aa *AbsoluteAnchor) Clone() openxml.Element {
	cloned := aa.CompositeElementBase.Clone()

	return &AbsoluteAnchor{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this AbsoluteAnchor element.
func (aa *AbsoluteAnchor) CloneNode(
	deep bool,
) openxml.Element {
	cloned := aa.CompositeElementBase.CloneNode(
		deep,
	)

	return &AbsoluteAnchor{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
