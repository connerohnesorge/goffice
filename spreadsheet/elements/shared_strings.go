package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// SharedStringTable represents the shared string table root element (x:sst).
// This element stores all unique strings used across the workbook for
// optimization. Cells reference strings by their index in this table.
type SharedStringTable struct {
	*openxml.PartRootElementBase

	// stringIndex maps string content to its index for O(1) lookups
	stringIndex map[string]int
}

// NewSharedStringTable creates a new SharedStringTable element.
func NewSharedStringTable() *SharedStringTable {
	elem := openxml.NewPartRootElement(
		NamespaceSML,
		"sst",
		PrefixDefault,
	)

	return &SharedStringTable{
		PartRootElementBase: elem,
		stringIndex: make(
			map[string]int,
		),
	}
}

// Count returns the total number of string references in the workbook.
// This may be greater than UniqueCount if strings are shared.
func (ss *SharedStringTable) Count() int {
	attr, found := ss.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCount sets the total number of string references.
func (ss *SharedStringTable) SetCount(
	count int,
) {
	ss.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.Itoa(count),
		),
	)
}

// UniqueCount returns the number of unique strings in the table.
func (ss *SharedStringTable) UniqueCount() int {
	attr, found := ss.GetAttribute(
		"uniqueCount",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetUniqueCount sets the number of unique strings.
func (ss *SharedStringTable) SetUniqueCount(
	count int,
) {
	ss.SetAttribute(
		openxml.NewAttribute(
			"",
			"uniqueCount",
			"",
			strconv.Itoa(count),
		),
	)
}

// Items returns an iterator over all SharedStringItem elements.
func (ss *SharedStringTable) Items() iter.Seq[*SharedStringItem] {
	return func(yield func(*SharedStringItem) bool) {
		for child := range ss.Children() {
			if child.LocalName() != "si" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var si *SharedStringItem
			if item, ok := child.(*SharedStringItem); ok {
				si = item
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				si = &SharedStringItem{CompositeElementBase: comp}
			}
			if si != nil && !yield(si) {
				return
			}
		}
	}
}

// ItemCount returns the count of unique string items.
func (ss *SharedStringTable) ItemCount() int {
	count := 0
	for range ss.Items() {
		count++
	}

	return count
}

// GetItem returns the SharedStringItem at the specified index,
// or nil if the index is out of range.
func (ss *SharedStringTable) GetItem(
	index int,
) *SharedStringItem {
	if index < 0 {
		return nil
	}
	i := 0
	for item := range ss.Items() {
		if i == index {
			return item
		}
		i++
	}

	return nil
}

// GetString returns the plain text content of the string at the
// specified index. Returns an empty string if the index is out of range.
func (ss *SharedStringTable) GetString(
	index int,
) string {
	item := ss.GetItem(index)
	if item == nil {
		return ""
	}

	return item.PlainText()
}

// AddString adds a string to the shared string table and returns its index.
// If the string already exists, returns the existing index (deduplication).
func (ss *SharedStringTable) AddString(
	text string,
) int {
	// Build index if not already built
	ss.ensureIndexBuilt()

	// Check for existing string
	if idx, exists := ss.stringIndex[text]; exists {
		return idx
	}

	// Create new string item
	si := NewSharedStringItem()
	si.SetPlainText(text)
	ss.AppendChild(si)

	// Update index
	idx := len(ss.stringIndex)
	ss.stringIndex[text] = idx

	// Update uniqueCount attribute
	ss.SetUniqueCount(idx + 1)

	return idx
}

// IndexOf returns the index of the specified string in the table.
// Returns -1 if the string is not found.
func (ss *SharedStringTable) IndexOf(
	text string,
) int {
	// Build index if not already built
	ss.ensureIndexBuilt()

	if idx, exists := ss.stringIndex[text]; exists {
		return idx
	}

	return -1
}

// ensureIndexBuilt builds the string index map if it hasn't been built yet.
func (ss *SharedStringTable) ensureIndexBuilt() {
	if ss.stringIndex == nil {
		ss.stringIndex = make(map[string]int)
	}

	// If the map is not empty, we're done
	if len(ss.stringIndex) != 0 {
		return
	}

	// Rebuild the index from existing items
	idx := 0
	for item := range ss.Items() {
		text := item.PlainText()
		if _, exists := ss.stringIndex[text]; !exists {
			ss.stringIndex[text] = idx
		}
		idx++
	}
}

// RebuildIndex rebuilds the internal string index from the current items.
// This should be called after loading a document or modifying items directly.
func (ss *SharedStringTable) RebuildIndex() {
	ss.stringIndex = make(map[string]int)
	idx := 0
	for item := range ss.Items() {
		text := item.PlainText()
		// Only store first occurrence for deduplication
		if _, exists := ss.stringIndex[text]; !exists {
			ss.stringIndex[text] = idx
		}
		idx++
	}
}

// Clear removes all string items and resets the index.
func (ss *SharedStringTable) Clear() {
	ss.RemoveAllChildren()
	ss.stringIndex = make(map[string]int)
	ss.SetCount(0)
	ss.SetUniqueCount(0)
}

// Clone creates a deep copy of this SharedStringTable element.
func (ss *SharedStringTable) Clone() openxml.Element {
	cloned := ss.PartRootElementBase.Clone()
	newSST := &SharedStringTable{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
		stringIndex: make(
			map[string]int,
		),
	}
	// Copy the string index
	for k, v := range ss.stringIndex {
		newSST.stringIndex[k] = v
	}

	return newSST
}

// CloneNode creates a copy of this SharedStringTable element.
//
//nolint:revive // flag-parameter: CloneNode interface requires deep parameter
func (ss *SharedStringTable) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ss.PartRootElementBase.CloneNode(
		deep,
	)
	newSST := &SharedStringTable{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
		stringIndex: make(
			map[string]int,
		),
	}
	if deep {
		// Copy the string index
		for k, v := range ss.stringIndex {
			newSST.stringIndex[k] = v
		}
	}

	return newSST
}
