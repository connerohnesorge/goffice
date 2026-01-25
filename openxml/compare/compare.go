package compare

import (
	"fmt"

	"github.com/connerohnesorge/goffice/openxml"
)

// DiffType represents the type of difference between two elements.
type DiffType int

const (
	// NoDiff indicates no difference.
	NoDiff DiffType = iota
	// Added indicates an element was added.
	Added
	// Deleted indicates an element was deleted.
	Deleted
	// Modified indicates an element was modified.
	Modified
)

// Diff represents a difference between two elements.
type Diff struct {
	Type       DiffType
	Path       string
	Key        string // Key identifies the changed item (e.g., attribute name)
	Index      int    // Index identifies the changed item in a list (e.g., child index)
	Message    string
	OldValue   interface{}
	NewValue   interface{}
	ChildDiffs []Diff
}

// Comparator is the interface for comparing OpenXML elements.
type Comparator interface {
	Compare(a, b openxml.Element) []Diff
}

// ElementComparator is the default implementation of Comparator.
type ElementComparator struct{}

// NewElementComparator creates a new ElementComparator.
func NewElementComparator() *ElementComparator {
	return &ElementComparator{}
}

// Compare compares two OpenXML elements and returns a list of differences.
func (c *ElementComparator) Compare(a, b openxml.Element) []Diff {
	var diffs []Diff

	if a == nil && b == nil {
		return diffs
	}
	if a == nil {
		return []Diff{{Type: Added, Message: "Element added", NewValue: b}}
	}
	if b == nil {
		return []Diff{{Type: Deleted, Message: "Element deleted", OldValue: a}}
	}

	if !a.QName().Equals(b.QName()) {
		return []Diff{{
			Type:     Modified,
			Message:  "Element type mismatch",
			OldValue: a.QName().String(),
			NewValue: b.QName().String(),
		}}
	}

	attrDiffs := c.compareAttributes(a, b)
	diffs = append(diffs, attrDiffs...)

	leafA, isLeafA := a.(openxml.LeafElement)
	leafB, isLeafB := b.(openxml.LeafElement)

	if isLeafA && isLeafB {
		if leafA.InnerText() != leafB.InnerText() {
			diffs = append(diffs, Diff{
				Type:     Modified,
				Message:  "Text content changed",
				OldValue: leafA.InnerText(),
				NewValue: leafB.InnerText(),
			})
		}
	} else if isLeafA != isLeafB {
		diffs = append(diffs, Diff{
			Type:    Modified,
			Message: "Element structure changed (leaf vs composite)",
		})
	}

	compA, isCompA := a.(openxml.CompositeElement)
	compB, isCompB := b.(openxml.CompositeElement)

	if isCompA && isCompB {
		childDiffs := c.compareChildren(compA, compB)
		diffs = append(diffs, childDiffs...)
	}

	return diffs
}

func (c *ElementComparator) compareAttributes(a, b openxml.Element) []Diff {
	var diffs []Diff
	attrsA := a.Attributes()
	attrsB := b.Attributes()

	mapA := make(map[string]openxml.OpenXmlAttribute)
	for _, attr := range attrsA {
		mapA[attr.QName().String()] = attr
	}

	mapB := make(map[string]openxml.OpenXmlAttribute)
	for _, attr := range attrsB {
		mapB[attr.QName().String()] = attr
	}

	for key, attrA := range mapA {
		attrB, exists := mapB[key]
		if !exists {
			diffs = append(diffs, Diff{
				Type:     Deleted,
				Key:      attrA.LocalName(),
				Message:  "Attribute deleted: " + attrA.LocalName(),
				OldValue: attrA.Value(),
			})
		} else if attrA.Value() != attrB.Value() {
			diffs = append(diffs, Diff{
				Type:     Modified,
				Key:      attrA.LocalName(),
				Message:  "Attribute modified: " + attrA.LocalName(),
				OldValue: attrA.Value(),
				NewValue: attrB.Value(),
			})
		}
	}

	for key, attrB := range mapB {
		if _, exists := mapA[key]; !exists {
			diffs = append(diffs, Diff{
				Type:     Added,
				Key:      attrB.LocalName(),
				Message:  "Attribute added: " + attrB.LocalName(),
				NewValue: attrB.Value(),
			})
		}
	}

	return diffs
}

func (c *ElementComparator) compareChildren(a, b openxml.CompositeElement) []Diff {
	var diffs []Diff
	childrenA := openxml.ToSlice(a.Children())
	childrenB := openxml.ToSlice(b.Children())

	maxLen := len(childrenA)
	if len(childrenB) > maxLen {
		maxLen = len(childrenB)
	}

	for i := 0; i < maxLen; i++ {
		var childA, childB openxml.Element
		if i < len(childrenA) {
			childA = childrenA[i]
		}
		if i < len(childrenB) {
			childB = childrenB[i]
		}

		childDiffs := c.Compare(childA, childB)
		if len(childDiffs) > 0 {
			diffs = append(diffs, Diff{
				Type:       Modified,
				Index:      i,
				Message:    fmt.Sprintf("Child at index %d changed", i),
				ChildDiffs: childDiffs,
			})
		}

	}

	return diffs
}
