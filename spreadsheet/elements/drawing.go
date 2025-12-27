package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// SpreadsheetDrawing namespace and prefix constants.
const (
	// NamespaceSpreadsheetDrawing is the SpreadsheetDrawing namespace.
	//nolint:revive // line-length-limit
	NamespaceSpreadsheetDrawing = "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"

	// PrefixXDR is the prefix for SpreadsheetDrawing elements.
	PrefixXDR = "xdr"

	// NamespaceDrawingML is the DrawingML namespace.
	NamespaceDrawingML = "http://schemas.openxmlformats.org/drawingml/2006/main"

	// PrefixA is the prefix for DrawingML elements.
	PrefixA = "a"
)

// WorksheetDrawing represents the root worksheet drawing element (xdr:wsDr).
// This is the root element of a drawing part and contains all anchored
// drawings.
type WorksheetDrawing struct {
	*openxml.PartRootElementBase
}

// NewWorksheetDrawing creates a new WorksheetDrawing element.
func NewWorksheetDrawing() *WorksheetDrawing {
	elem := openxml.NewPartRootElement(
		NamespaceSpreadsheetDrawing,
		"wsDr",
		PrefixXDR,
	)

	return &WorksheetDrawing{
		PartRootElementBase: elem,
	}
}

// TwoCellAnchors returns all TwoCellAnchor child elements.
func (wd *WorksheetDrawing) TwoCellAnchors() []*TwoCellAnchor {
	var result []*TwoCellAnchor
	for child := range wd.Children() {
		if child.LocalName() != "twoCellAnchor" ||
			child.NamespaceURI() != NamespaceSpreadsheetDrawing {
			continue
		}
		if tca, ok := child.(*TwoCellAnchor); ok {
			result = append(result, tca)
		} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
			result = append(result, &TwoCellAnchor{CompositeElementBase: comp})
		}
	}

	return result
}

// AddTwoCellAnchor adds a new TwoCellAnchor element.
func (wd *WorksheetDrawing) AddTwoCellAnchor() *TwoCellAnchor {
	tca := NewTwoCellAnchor()
	wd.AppendChild(tca)

	return tca
}

// OneCellAnchors returns all OneCellAnchor child elements.
func (wd *WorksheetDrawing) OneCellAnchors() []*OneCellAnchor {
	var result []*OneCellAnchor
	for child := range wd.Children() {
		if child.LocalName() != "oneCellAnchor" ||
			child.NamespaceURI() != NamespaceSpreadsheetDrawing {
			continue
		}
		if oca, ok := child.(*OneCellAnchor); ok {
			result = append(result, oca)
		} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
			result = append(result, &OneCellAnchor{CompositeElementBase: comp})
		}
	}

	return result
}

// AddOneCellAnchor adds a new OneCellAnchor element.
func (wd *WorksheetDrawing) AddOneCellAnchor() *OneCellAnchor {
	oca := NewOneCellAnchor()
	wd.AppendChild(oca)

	return oca
}

// AbsoluteAnchors returns all AbsoluteAnchor child elements.
func (wd *WorksheetDrawing) AbsoluteAnchors() []*AbsoluteAnchor {
	var result []*AbsoluteAnchor
	for child := range wd.Children() {
		if child.LocalName() != "absoluteAnchor" ||
			child.NamespaceURI() != NamespaceSpreadsheetDrawing {
			continue
		}
		if aa, ok := child.(*AbsoluteAnchor); ok {
			result = append(result, aa)
		} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
			result = append(result, &AbsoluteAnchor{CompositeElementBase: comp})
		}
	}

	return result
}

// AddAbsoluteAnchor adds a new AbsoluteAnchor element.
func (wd *WorksheetDrawing) AddAbsoluteAnchor() *AbsoluteAnchor {
	aa := NewAbsoluteAnchor()
	wd.AppendChild(aa)

	return aa
}

// Clone creates a deep copy of this WorksheetDrawing element.
func (wd *WorksheetDrawing) Clone() openxml.Element {
	cloned := wd.PartRootElementBase.Clone()

	return &WorksheetDrawing{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this WorksheetDrawing element.
func (wd *WorksheetDrawing) CloneNode(
	deep bool,
) openxml.Element {
	cloned := wd.PartRootElementBase.CloneNode(
		deep,
	)

	return &WorksheetDrawing{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}
