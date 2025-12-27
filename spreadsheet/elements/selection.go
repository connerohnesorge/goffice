package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Selection represents the selection element (x:selection) for cell selection.
type Selection struct {
	*openxml.CompositeElementBase
}

// NewSelection creates a new Selection element.
func NewSelection() *Selection {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"selection",
		PrefixDefault,
	)

	return &Selection{CompositeElementBase: elem}
}

// Pane returns the pane this selection is in.
func (s *Selection) Pane() PanePosition {
	attr, found := s.GetAttribute("pane", "")
	if !found {
		return PanePositionTopLeft
	}

	return PanePosition(attr.Value())
}

// SetPane sets the pane this selection is in.
func (s *Selection) SetPane(pos PanePosition) {
	if pos == "" || pos == PanePositionTopLeft {
		s.RemoveAttribute("pane", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"pane",
			"",
			string(pos),
		),
	)
}

// ActiveCell returns the active cell reference (e.g., "A1").
func (s *Selection) ActiveCell() string {
	attr, found := s.GetAttribute(
		"activeCell",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetActiveCell sets the active cell reference.
func (s *Selection) SetActiveCell(cell string) {
	if cell == "" {
		s.RemoveAttribute("activeCell", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"activeCell",
			"",
			cell,
		),
	)
}

// ActiveCellId returns the 0-based index of the active cell in the selection.
func (s *Selection) ActiveCellId() int {
	attr, found := s.GetAttribute(
		"activeCellId",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetActiveCellId sets the index of the active cell in the selection.
func (s *Selection) SetActiveCellId(id int) {
	if id == 0 {
		s.RemoveAttribute("activeCellId", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"activeCellId",
			"",
			strconv.Itoa(id),
		),
	)
}

// Sqref returns the sequence of references for the selection.
// This is a space-separated list of cell references or ranges
// (e.g., "A1:B2 D4").
func (s *Selection) Sqref() string {
	attr, found := s.GetAttribute("sqref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSqref sets the sequence of references for the selection.
func (s *Selection) SetSqref(sqref string) {
	if sqref == "" {
		s.RemoveAttribute("sqref", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"sqref",
			"",
			sqref,
		),
	)
}

// SetSelection sets both the active cell and the selection range.
func (s *Selection) SetSelection(
	activeCell, sqref string,
) {
	s.SetActiveCell(activeCell)
	s.SetSqref(sqref)
}

// Clone creates a deep copy of this Selection element.
func (s *Selection) Clone() openxml.Element {
	cloned := s.CompositeElementBase.Clone()

	return &Selection{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Selection element.
func (s *Selection) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.CompositeElementBase.CloneNode(
		deep,
	)

	return &Selection{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
