package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Tab represents a tab character element (w:tab).
type Tab struct {
	*openxml.CompositeElementBase
}

// NewTab creates a new Tab element.
func NewTab() *Tab {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tab",
		PrefixW,
	)

	return &Tab{CompositeElementBase: elem}
}

// Clone creates a deep copy of this Tab element.
func (t *Tab) Clone() openxml.Element {
	cloned := t.CompositeElementBase.Clone()

	return &Tab{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Tab element.
func (t *Tab) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.CompositeElementBase.CloneNode(
		deep,
	)

	return &Tab{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
