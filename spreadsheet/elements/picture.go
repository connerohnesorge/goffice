package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Picture represents the picture element (x:picture) for background images.
type Picture struct {
	*openxml.CompositeElementBase
}

// NewPicture creates a new Picture element.
func NewPicture() *Picture {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"picture",
		PrefixDefault,
	)

	return &Picture{CompositeElementBase: elem}
}

// RelationshipId returns the relationship ID linking to the image part.
func (p *Picture) RelationshipId() string {
	attr, found := p.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID linking to the image part.
func (p *Picture) SetRelationshipId(id string) {
	if id == "" {
		p.RemoveAttribute(
			"id",
			NamespaceRelationships,
		)

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"id",
			PrefixR,
			id,
		),
	)
}

// Clone creates a deep copy of this Picture element.
func (p *Picture) Clone() openxml.Element {
	cloned := p.CompositeElementBase.Clone()

	return &Picture{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Picture element.
func (p *Picture) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.CompositeElementBase.CloneNode(
		deep,
	)

	return &Picture{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
