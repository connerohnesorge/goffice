package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// MergeCells represents the merge cells container element (x:mergeCells).
// It contains all merged cell ranges in a worksheet.
type MergeCells struct {
	*openxml.CompositeElementBase
}

// NewMergeCells creates a new MergeCells element.
func NewMergeCells() *MergeCells {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"mergeCells",
		PrefixDefault,
	)

	return &MergeCells{CompositeElementBase: elem}
}

// Count returns the count attribute value.
func (mc *MergeCells) Count() int {
	attr, found := mc.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCount sets the count attribute.
func (mc *MergeCells) SetCount(count int) {
	if count == 0 {
		mc.RemoveAttribute("count", "")

		return
	}
	mc.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.Itoa(count),
		),
	)
}

// updateCount updates the count attribute based on the number of merge cells.
func (mc *MergeCells) updateCount() {
	count := 0
	for range mc.GetMergeCells() {
		count++
	}
	mc.SetCount(count)
}

// GetMergeCells returns an iterator over all MergeCell elements.
func (mc *MergeCells) GetMergeCells() iter.Seq[*MergeCell] {
	return func(yield func(*MergeCell) bool) {
		for child := range mc.Children() {
			if child.LocalName() != "mergeCell" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var mergeCell *MergeCell
			switch v := child.(type) {
			case *MergeCell:
				mergeCell = v
			case *openxml.LeafElementBase:
				mergeCell = &MergeCell{LeafElementBase: v}
			}
			if mergeCell != nil &&
				!yield(mergeCell) {
				return
			}
		}
	}
}

// GetMergeCellByRef returns the merge cell with the given reference, or nil
// if not found.
func (mc *MergeCells) GetMergeCellByRef(
	ref string,
) *MergeCell {
	for mergeCell := range mc.GetMergeCells() {
		if mergeCell.Ref() == ref {
			return mergeCell
		}
	}

	return nil
}

// AddMergeCell adds a new merge cell with the given reference range
// (e.g., "A1:D1").
func (mc *MergeCells) AddMergeCell(
	ref string,
) *MergeCell {
	mergeCell := NewMergeCell()
	mergeCell.SetRef(ref)
	mc.AppendChild(mergeCell)
	mc.updateCount()

	return mergeCell
}

// RemoveMergeCell removes the merge cell with the given reference.
func (mc *MergeCells) RemoveMergeCell(
	ref string,
) bool {
	mergeCell := mc.GetMergeCellByRef(ref)
	if mergeCell == nil {
		return false
	}
	removed := mc.RemoveChild(mergeCell)
	if removed {
		mc.updateCount()
	}

	return removed
}

// HasMergeCell returns whether a merge cell exists with the given reference.
func (mc *MergeCells) HasMergeCell(
	ref string,
) bool {
	return mc.GetMergeCellByRef(ref) != nil
}

// Clone creates a deep copy of this MergeCells element.
func (mc *MergeCells) Clone() openxml.Element {
	cloned := mc.CompositeElementBase.Clone()

	return &MergeCells{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this MergeCells element.
func (mc *MergeCells) CloneNode(
	deep bool,
) openxml.Element {
	cloned := mc.CompositeElementBase.CloneNode(
		deep,
	)

	return &MergeCells{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// MergeCell represents a single merge cell element (x:mergeCell).
// It defines a range of cells that are merged.
type MergeCell struct {
	*openxml.LeafElementBase
}

// NewMergeCell creates a new MergeCell element.
func NewMergeCell() *MergeCell {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"mergeCell",
		PrefixDefault,
	)

	return &MergeCell{LeafElementBase: elem}
}

// NewMergeCellWithRef creates a new MergeCell element with the given reference.
func NewMergeCellWithRef(ref string) *MergeCell {
	mergeCell := NewMergeCell()
	mergeCell.SetRef(ref)

	return mergeCell
}

// Ref returns the cell reference range (e.g., "A1:D1"). Attribute: ref.
func (m *MergeCell) Ref() string {
	attr, found := m.GetAttribute("ref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the cell reference range (e.g., "A1:D1"). Attribute: ref.
func (m *MergeCell) SetRef(ref string) {
	m.SetAttribute(
		openxml.NewAttribute(
			"",
			"ref",
			"",
			ref,
		),
	)
}

// Clone creates a deep copy of this MergeCell element.
func (m *MergeCell) Clone() openxml.Element {
	cloned := m.LeafElementBase.Clone()

	return &MergeCell{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this MergeCell element.
func (m *MergeCell) CloneNode(
	deep bool,
) openxml.Element {
	cloned := m.LeafElementBase.CloneNode(deep)

	return &MergeCell{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
