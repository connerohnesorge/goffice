package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Drawing represents the drawing element (x:drawing).
// This element references a DrawingsPart containing charts, shapes, etc.
type Drawing struct {
	*openxml.CompositeElementBase
}

// NewDrawing creates a new Drawing element.
func NewDrawing() *Drawing {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"drawing",
		PrefixDefault,
	)

	return &Drawing{CompositeElementBase: elem}
}

// RelationshipId returns the relationship ID linking to the drawings part.
func (d *Drawing) RelationshipId() string {
	attr, found := d.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID linking to the drawings part.
func (d *Drawing) SetRelationshipId(id string) {
	if id == "" {
		d.RemoveAttribute(
			"id",
			NamespaceRelationships,
		)

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"id",
			PrefixR,
			id,
		),
	)
}

// Clone creates a deep copy of this Drawing element.
func (d *Drawing) Clone() openxml.Element {
	cloned := d.CompositeElementBase.Clone()

	return &Drawing{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Drawing element.
func (d *Drawing) CloneNode(
	deep bool,
) openxml.Element {
	cloned := d.CompositeElementBase.CloneNode(
		deep,
	)

	return &Drawing{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
