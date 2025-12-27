package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// SheetState represents the visibility state of a sheet.
type SheetState string

const (
	// SheetStateVisible indicates the sheet is visible (default).
	SheetStateVisible SheetState = "visible"
	// SheetStateHidden indicates the sheet is hidden but can be unhidden
	// via UI.
	SheetStateHidden SheetState = "hidden"
	// SheetStateVeryHidden indicates the sheet is hidden and cannot be
	// unhidden via UI.
	SheetStateVeryHidden SheetState = "veryHidden"
)

// Sheets represents the sheets container element (x:sheets).
type Sheets struct {
	*openxml.CompositeElementBase
}

// NewSheets creates a new Sheets element.
func NewSheets() *Sheets {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"sheets",
		PrefixDefault,
	)

	return &Sheets{CompositeElementBase: elem}
}

// Sheets returns an iterator over all Sheet elements.
func (s *Sheets) Sheets() iter.Seq[*Sheet] {
	return func(yield func(*Sheet) bool) {
		for child := range s.Children() {
			local := child.LocalName()
			if (local != "sheet" && local != "Sheet") ||
				(child.NamespaceURI() != NamespaceSML &&
					child.NamespaceURI() != "") {
				continue
			}
			var sheet *Sheet
			if sh, ok := child.(*Sheet); ok {
				sheet = sh
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				sheet = &Sheet{CompositeElementBase: comp}
			}
			if sheet != nil && !yield(sheet) {
				return
			}
		}
	}
}

// SheetCount returns the number of sheets.
func (s *Sheets) SheetCount() int {
	count := 0
	for range s.Sheets() {
		count++
	}

	return count
}

// GetSheetByName returns the sheet with the given name, or nil if not found.
func (s *Sheets) GetSheetByName(
	name string,
) *Sheet {
	for sheet := range s.Sheets() {
		if sheet.Name() == name {
			return sheet
		}
	}

	return nil
}

// GetSheetById returns the sheet with the given sheet ID, or nil if not found.
func (s *Sheets) GetSheetById(
	sheetId int,
) *Sheet {
	for sheet := range s.Sheets() {
		if sheet.SheetId() == sheetId {
			return sheet
		}
	}

	return nil
}

// GetSheetByIndex returns the sheet at the given zero-based index, or nil
// if out of range.
func (s *Sheets) GetSheetByIndex(
	index int,
) *Sheet {
	i := 0
	for sheet := range s.Sheets() {
		if i == index {
			return sheet
		}
		i++
	}

	return nil
}

// AddSheet adds a new sheet with the given name and returns it. The sheetId
// should be unique and is typically auto-generated. The relationshipId links
// to the worksheet part.
func (s *Sheets) AddSheet(
	name string,
	sheetId int,
	relationshipId string,
) *Sheet {
	sheet := NewSheet()
	sheet.SetName(name)
	sheet.SetSheetId(sheetId)
	sheet.SetRelationshipId(relationshipId)
	s.AppendChild(sheet)

	return sheet
}

// RemoveSheet removes a sheet from the collection.
func (s *Sheets) RemoveSheet(sheet *Sheet) bool {
	return s.RemoveChild(sheet)
}

// RemoveSheetByName removes the sheet with the given name.
func (s *Sheets) RemoveSheetByName(
	name string,
) bool {
	sheet := s.GetSheetByName(name)
	if sheet == nil {
		return false
	}

	return s.RemoveChild(sheet)
}

// NextSheetId returns the next available sheet ID.
func (s *Sheets) NextSheetId() int {
	maxId := 0
	for sheet := range s.Sheets() {
		if sheet.SheetId() > maxId {
			maxId = sheet.SheetId()
		}
	}

	return maxId + 1
}

// Clone creates a deep copy of this Sheets element.
func (s *Sheets) Clone() openxml.Element {
	cloned := s.CompositeElementBase.Clone()

	return &Sheets{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Sheets element.
func (s *Sheets) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.CompositeElementBase.CloneNode(
		deep,
	)

	return &Sheets{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Sheet represents a single sheet element (x:sheet).
// This is a reference to a worksheet part, not the worksheet content itself.
type Sheet struct {
	*openxml.CompositeElementBase
}

// NewSheet creates a new Sheet element.
func NewSheet() *Sheet {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"sheet",
		PrefixDefault,
	)

	return &Sheet{CompositeElementBase: elem}
}

// Name returns the name of the sheet.
func (sh *Sheet) Name() string {
	attr, found := sh.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the name of the sheet.
func (sh *Sheet) SetName(name string) {
	sh.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// SheetId returns the unique sheet ID within the workbook.
func (sh *Sheet) SheetId() int {
	attr, found := sh.GetAttribute("sheetId", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetSheetId sets the unique sheet ID.
func (sh *Sheet) SetSheetId(id int) {
	sh.SetAttribute(
		openxml.NewAttribute(
			"",
			"sheetId",
			"",
			strconv.Itoa(id),
		),
	)
}

// State returns the visibility state of the sheet.
func (sh *Sheet) State() SheetState {
	attr, found := sh.GetAttribute("state", "")
	if !found {
		return SheetStateVisible
	}

	return SheetState(attr.Value())
}

// SetState sets the visibility state of the sheet.
func (sh *Sheet) SetState(state SheetState) {
	if state == "" || state == SheetStateVisible {
		sh.RemoveAttribute("state", "")

		return
	}
	sh.SetAttribute(
		openxml.NewAttribute(
			"",
			"state",
			"",
			string(state),
		),
	)
}

// RelationshipId returns the relationship ID linking to the worksheet part.
// This is the r:id attribute.
func (sh *Sheet) RelationshipId() string {
	attr, found := sh.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID linking to the worksheet part.
func (sh *Sheet) SetRelationshipId(id string) {
	sh.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"id",
			PrefixR,
			id,
		),
	)
}

// IsVisible returns true if the sheet is visible.
func (sh *Sheet) IsVisible() bool {
	return sh.State() == SheetStateVisible ||
		sh.State() == ""
}

// IsHidden returns true if the sheet is hidden.
func (sh *Sheet) IsHidden() bool {
	return sh.State() == SheetStateHidden
}

// IsVeryHidden returns true if the sheet is very hidden.
func (sh *Sheet) IsVeryHidden() bool {
	return sh.State() == SheetStateVeryHidden
}

// Hide hides the sheet (can be unhidden via UI).
func (sh *Sheet) Hide() {
	sh.SetState(SheetStateHidden)
}

// VeryHide makes the sheet very hidden (cannot be unhidden via UI).
func (sh *Sheet) VeryHide() {
	sh.SetState(SheetStateVeryHidden)
}

// Show makes the sheet visible.
func (sh *Sheet) Show() {
	sh.SetState(SheetStateVisible)
}

// Clone creates a deep copy of this Sheet element.
func (sh *Sheet) Clone() openxml.Element {
	cloned := sh.CompositeElementBase.Clone()

	return &Sheet{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Sheet element.
func (sh *Sheet) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sh.CompositeElementBase.CloneNode(
		deep,
	)

	return &Sheet{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ItemCount returns the number of sheets.
// This is an alias for SheetCount for consistency with other containers.
func (s *Sheets) ItemCount() int {
	return s.SheetCount()
}

// SheetEntry is an alias for Sheet to avoid naming conflicts with high-level
// Sheet types. This represents a sheet reference entry in workbook.xml,
// not the actual worksheet content.
type SheetEntry = Sheet

// NewSheetEntry creates a new sheet entry element.
func NewSheetEntry() *SheetEntry {
	return NewSheet()
}
