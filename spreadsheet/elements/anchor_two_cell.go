//nolint:revive // comments-density
package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// TwoCellAnchor represents a two-cell anchor element (xdr:twoCellAnchor).
// The drawing is anchored to two cells and resizes with cell dimensions.
type TwoCellAnchor struct {
	*openxml.CompositeElementBase
}

// NewTwoCellAnchor creates a new TwoCellAnchor element.
func NewTwoCellAnchor() *TwoCellAnchor {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"twoCellAnchor",
		PrefixXDR,
	)

	return &TwoCellAnchor{
		CompositeElementBase: elem,
	}
}

// EditAs returns the anchor behavior. Attribute: editAs.
func (tc *TwoCellAnchor) EditAs() EditAs {
	attr, found := tc.GetAttribute("editAs", "")
	if !found {
		return EditAsTwoCell // Default
	}

	return EditAs(attr.Value())
}

// SetEditAs sets the anchor behavior. Attribute: editAs.
func (tc *TwoCellAnchor) SetEditAs(
	value EditAs,
) {
	if value == EditAsTwoCell {
		tc.RemoveAttribute(
			"editAs",
			"",
		) // twoCell is default

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"editAs",
			"",
			string(value),
		),
	)
}

// From returns the From marker element, or nil if not present.
func (tc *TwoCellAnchor) From() *FromMarker {
	elem := tc.GetElement(
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
func (tc *TwoCellAnchor) GetOrCreateFrom() *FromMarker {
	fm := tc.From()
	if fm != nil {
		return fm
	}
	fm = NewFromMarker()
	if first := tc.FirstChild(); first != nil {
		tc.InsertBefore(fm, first)
	} else {
		tc.AppendChild(fm)
	}

	return fm
}

// To returns the To marker element, or nil if not present.
func (tc *TwoCellAnchor) To() *ToMarker {
	elem := tc.GetElement(
		"to",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if tm, ok := elem.(*ToMarker); ok {
		return tm
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ToMarker{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTo returns the To marker, creating if needed.
func (tc *TwoCellAnchor) GetOrCreateTo() *ToMarker {
	tm := tc.To()
	if tm != nil {
		return tm
	}
	tm = NewToMarker()
	fm := tc.From()
	if fm != nil {
		tc.InsertAfter(tm, fm)
	} else {
		tc.AppendChild(tm)
	}

	return tm
}

// Shape returns the Shape element, or nil if not present.
func (tc *TwoCellAnchor) Shape() *Shape {
	elem := tc.GetElement(
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
func (tc *TwoCellAnchor) GetOrCreateShape() *Shape {
	sp := tc.Shape()
	if sp != nil {
		return sp
	}
	sp = NewShape()
	tm := tc.To()
	if tm != nil {
		tc.InsertAfter(sp, tm)
	} else {
		tc.AppendChild(sp)
	}

	return sp
}

// Picture returns the Picture element, or nil if not present.
func (tc *TwoCellAnchor) Picture() *DrawingPicture {
	elem := tc.GetElement(
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
func (tc *TwoCellAnchor) GetOrCreatePicture() *DrawingPicture {
	pic := tc.Picture()
	if pic != nil {
		return pic
	}
	pic = NewDrawingPicture()
	tm := tc.To()
	if tm != nil {
		tc.InsertAfter(pic, tm)
	} else {
		tc.AppendChild(pic)
	}

	return pic
}

// GraphicFrame returns the GraphicFrame element, or nil if not present.
func (tc *TwoCellAnchor) GraphicFrame() *GraphicFrame {
	elem := tc.GetElement(
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
func (tc *TwoCellAnchor) GetOrCreateGraphicFrame() *GraphicFrame {
	gf := tc.GraphicFrame()
	if gf != nil {
		return gf
	}
	gf = NewGraphicFrame()
	tm := tc.To()
	if tm != nil {
		tc.InsertAfter(gf, tm)
	} else {
		tc.AppendChild(gf)
	}

	return gf
}

// ClientData returns the ClientData element, or nil if not present.
func (tc *TwoCellAnchor) ClientData() *ClientData {
	elem := tc.GetElement(
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
func (tc *TwoCellAnchor) GetOrCreateClientData() *ClientData {
	cd := tc.ClientData()
	if cd != nil {
		return cd
	}
	cd = NewClientData()
	tc.AppendChild(cd)

	return cd
}

// Clone creates a deep copy of this TwoCellAnchor element.
func (tc *TwoCellAnchor) Clone() openxml.Element {
	cloned := tc.CompositeElementBase.Clone()

	return &TwoCellAnchor{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this TwoCellAnchor element.
func (tc *TwoCellAnchor) CloneNode(
	deep bool,
) openxml.Element {
	cloned := tc.CompositeElementBase.CloneNode(
		deep,
	)

	return &TwoCellAnchor{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
