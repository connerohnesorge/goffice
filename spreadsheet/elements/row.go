package elements

//revive:disable:file-length-limit many row properties and methods

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Row represents a row element (x:row) in a worksheet.
type Row struct {
	*openxml.CompositeElementBase
}

// NewRow creates a new Row element.
func NewRow() *Row {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		elemNameRow,
		PrefixDefault,
	)

	return &Row{CompositeElementBase: elem}
}

// RowIndex returns the row index (1-based). Attribute: r.
func (r *Row) RowIndex() uint32 {
	attr, found := r.GetAttribute("r", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetRowIndex sets the row index (1-based). Attribute: r.
func (r *Row) SetRowIndex(index uint32) {
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"r",
			"",
			strconv.FormatUint(
				uint64(index),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Spans returns the spans attribute value (e.g., "1:10").
func (r *Row) Spans() string {
	attr, found := r.GetAttribute("spans", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSpans sets the spans attribute.
func (r *Row) SetSpans(spans string) {
	if spans == "" {
		r.RemoveAttribute("spans", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"spans",
			"",
			spans,
		),
	)
}

// StyleIndex returns the style index for the row. Attribute: s.
func (r *Row) StyleIndex() uint32 {
	attr, found := r.GetAttribute("s", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetStyleIndex sets the style index for the row. Attribute: s.
func (r *Row) SetStyleIndex(index uint32) {
	if index == 0 {
		r.RemoveAttribute("s", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"s",
			"",
			strconv.FormatUint(
				uint64(index),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// CustomFormat returns whether a custom format is applied.
// Attribute: customFormat.
func (r *Row) CustomFormat() bool {
	attr, found := r.GetAttribute(
		"customFormat",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCustomFormat sets whether a custom format is applied. Attribute: customFormat.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *Row) SetCustomFormat(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"customFormat",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("customFormat", "")
	}
}

// Height returns the row height in points. Attribute: ht.
func (r *Row) Height() float64 {
	attr, found := r.GetAttribute("ht", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetHeight sets the row height in points. Attribute: ht.
func (r *Row) SetHeight(height float64) {
	if height == 0 {
		r.RemoveAttribute("ht", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"ht",
			"",
			strconv.FormatFloat(
				height,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// Hidden returns whether the row is hidden. Attribute: hidden.
func (r *Row) Hidden() bool {
	attr, found := r.GetAttribute("hidden", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHidden sets whether the row is hidden. Attribute: hidden.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *Row) SetHidden(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"hidden",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("hidden", "")
	}
}

// CustomHeight returns whether a custom height was set.
// Attribute: customHeight.
func (r *Row) CustomHeight() bool {
	attr, found := r.GetAttribute(
		"customHeight",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCustomHeight sets whether a custom height was set. Attribute: customHeight.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *Row) SetCustomHeight(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"customHeight",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("customHeight", "")
	}
}

// OutlineLevel returns the outline level of the row (0-7).
// Attribute: outlineLevel.
func (r *Row) OutlineLevel() int {
	attr, found := r.GetAttribute(
		"outlineLevel",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetOutlineLevel sets the outline level of the row (0-7).
// Attribute: outlineLevel.
func (r *Row) SetOutlineLevel(level int) {
	if level == 0 {
		r.RemoveAttribute("outlineLevel", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"outlineLevel",
			"",
			strconv.Itoa(level),
		),
	)
}

// Collapsed returns whether the row outline is collapsed. Attribute: collapsed.
func (r *Row) Collapsed() bool {
	attr, found := r.GetAttribute("collapsed", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCollapsed sets whether the row outline is collapsed. Attribute: collapsed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *Row) SetCollapsed(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"collapsed",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("collapsed", "")
	}
}

// ThickTop returns whether the row has a thick top border. Attribute: thickTop.
func (r *Row) ThickTop() bool {
	attr, found := r.GetAttribute("thickTop", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetThickTop sets whether the row has a thick top border. Attribute: thickTop.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *Row) SetThickTop(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"thickTop",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("thickTop", "")
	}
}

// ThickBot returns whether the row has a thick bottom border.
// Attribute: thickBot.
func (r *Row) ThickBot() bool {
	attr, found := r.GetAttribute("thickBot", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetThickBot sets whether the row has a thick bottom border. Attribute: thickBot.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *Row) SetThickBot(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"thickBot",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("thickBot", "")
	}
}

// Ph returns whether phonetic information should be displayed. Attribute: ph.
func (r *Row) Ph() bool {
	attr, found := r.GetAttribute("ph", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPh sets whether phonetic information should be displayed. Attribute: ph.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *Row) SetPh(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"ph",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("ph", "")
	}
}

// Cells returns an iterator over all Cell elements in this row.
func (r *Row) Cells() iter.Seq[*Cell] {
	return func(yield func(*Cell) bool) {
		for child := range r.Children() {
			if child.LocalName() != "c" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var cell *Cell
			switch v := child.(type) {
			case *Cell:
				cell = v
			case *openxml.CompositeElementBase:
				cell = &Cell{CompositeElementBase: v}
			}
			if cell != nil && !yield(cell) {
				return
			}
		}
	}
}

// CellCount returns the number of cells in the row.
func (r *Row) CellCount() int {
	count := 0
	for range r.Cells() {
		count++
	}

	return count
}

// GetCell returns the cell at the given reference (e.g., "A1"), or nil if not
// found.
func (r *Row) GetCell(ref string) *Cell {
	for cell := range r.Cells() {
		if cell.Reference() == ref {
			return cell
		}
	}

	return nil
}

// GetOrCreateCell returns the cell at the given reference, creating it if it
// does not exist.
func (r *Row) GetOrCreateCell(ref string) *Cell {
	cell := r.GetCell(ref)
	if cell != nil {
		return cell
	}

	return r.AddCell(ref)
}

// AddCell adds a new cell at the given reference (e.g., "A1"). The cell is
// inserted in the correct position to maintain sorted order by column.
func (r *Row) AddCell(ref string) *Cell {
	cell := CreateCell()
	cell.SetReference(ref)

	// Find the correct position to insert the cell (sorted by column)
	newCol := columnFromReference(ref)
	var insertBefore openxml.Element
	for child := range r.Children() {
		if child.LocalName() != "c" ||
			child.NamespaceURI() != NamespaceSML {
			continue
		}
		var existingCell *Cell
		switch v := child.(type) {
		case *Cell:
			existingCell = v
		case *openxml.CompositeElementBase:
			existingCell = &Cell{CompositeElementBase: v}
		}
		if existingCell == nil {
			continue
		}
		existingCol := columnFromReference(
			existingCell.Reference(),
		)
		if compareColumns(
			existingCol,
			newCol,
		) > 0 {
			insertBefore = child

			break
		}
	}

	if insertBefore != nil {
		r.InsertBefore(cell, insertBefore)
	} else {
		r.AppendChild(cell)
	}

	return cell
}

// RemoveCell removes the cell at the given reference.
func (r *Row) RemoveCell(ref string) bool {
	cell := r.GetCell(ref)
	if cell == nil {
		return false
	}

	return r.RemoveChild(cell)
}

// Clone creates a deep copy of this Row element.
func (r *Row) Clone() openxml.Element {
	cloned := r.CompositeElementBase.Clone()

	return &Row{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Row element.
func (r *Row) CloneNode(
	deep bool,
) openxml.Element {
	cloned := r.CompositeElementBase.CloneNode(
		deep,
	)

	return &Row{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// columnFromReference extracts the column letters from a cell reference.
// For example, "AA15" returns "AA".
func columnFromReference(ref string) string {
	for i, c := range ref {
		if c >= '0' && c <= '9' {
			return ref[:i]
		}
	}

	return ref
}

// compareColumns compares two column letters.
// Returns negative if col1 < col2, zero if equal, positive if col1 > col2.
func compareColumns(col1, col2 string) int {
	// First compare by length (shorter column letters come first: A < AA)
	if len(col1) != len(col2) {
		return len(col1) - len(col2)
	}

	// Same length, compare lexicographically
	if col1 < col2 {
		return -1
	}
	if col1 > col2 {
		return 1
	}

	return 0
}
