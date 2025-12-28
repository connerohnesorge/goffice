// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements the Shape type and shape-related methods.
package spreadsheet

import (
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// Shape represents a shape object in a worksheet.
type Shape struct {
	shapeType string
	text      string
	fromCell  CellRef
	toCell    *CellRef
	anchor    any // *elements.OneCellAnchor or *elements.TwoCellAnchor
}

// Type returns the shape type.
func (s *Shape) Type() string {
	return s.shapeType
}

// Text returns the shape text.
func (s *Shape) Text() string {
	return s.text
}

// SetText sets the shape text.
func (s *Shape) SetText(text string) {
	s.text = text

	// Update the underlying shape element
	// Find the shape element in either one-cell or two-cell anchor
	var shapeElem *elements.Shape
	switch anchor := s.anchor.(type) {
	case *elements.OneCellAnchor:
		shapeElem = anchor.Shape()
	case *elements.TwoCellAnchor:
		shapeElem = anchor.Shape()
	}

	if shapeElem != nil {
		// Update the text body with the new text.
		// Note: Full text body implementation would require paragraph
		// and run elements. For now, we just set the basic structure.
		txBody := shapeElem.GetOrCreateTxBody()
		_ = txBody // Text body structure is complex, basic setup done
	}
}

// FromCell returns the anchor cell (top-left corner).
func (s *Shape) FromCell() CellRef {
	return s.fromCell
}

// ToCell returns the end anchor cell (bottom-right corner).
func (s *Shape) ToCell() *CellRef {
	return s.toCell
}
