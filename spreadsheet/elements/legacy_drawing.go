package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// LegacyDrawing represents the legacy drawing element (x:legacyDrawing). This
// element references a VML drawing part containing legacy shapes, comments,
// etc.
type LegacyDrawing struct {
	*openxml.CompositeElementBase
}

// NewLegacyDrawing creates a new LegacyDrawing element.
func NewLegacyDrawing() *LegacyDrawing {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"legacyDrawing",
		PrefixDefault,
	)

	return &LegacyDrawing{
		CompositeElementBase: elem,
	}
}

// RelationshipId returns the relationship ID linking to the VML drawing part.
func (ld *LegacyDrawing) RelationshipId() string {
	attr, found := ld.GetAttribute(
		elemNameID,
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID linking to the VML drawing part.
func (ld *LegacyDrawing) SetRelationshipId(
	id string,
) {
	if id == "" {
		ld.RemoveAttribute(
			elemNameID,
			NamespaceRelationships,
		)

		return
	}
	ld.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			elemNameID,
			PrefixR,
			id,
		),
	)
}

// Clone creates a deep copy of this LegacyDrawing element.
func (ld *LegacyDrawing) Clone() openxml.Element {
	cloned := ld.CompositeElementBase.Clone()

	return &LegacyDrawing{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this LegacyDrawing element.
func (ld *LegacyDrawing) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ld.CompositeElementBase.CloneNode(
		deep,
	)

	return &LegacyDrawing{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// LegacyDrawingHF represents the legacy drawing for header/footer element
// (x:legacyDrawingHF).
type LegacyDrawingHF struct {
	*openxml.CompositeElementBase
}

// NewLegacyDrawingHF creates a new LegacyDrawingHF element.
func NewLegacyDrawingHF() *LegacyDrawingHF {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"legacyDrawingHF",
		PrefixDefault,
	)

	return &LegacyDrawingHF{
		CompositeElementBase: elem,
	}
}

// RelationshipId returns the relationship ID linking to the VML drawing part.
func (lh *LegacyDrawingHF) RelationshipId() string {
	attr, found := lh.GetAttribute(
		elemNameID,
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID linking to the VML drawing part.
func (lh *LegacyDrawingHF) SetRelationshipId(
	id string,
) {
	if id == "" {
		lh.RemoveAttribute(
			elemNameID,
			NamespaceRelationships,
		)

		return
	}
	lh.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			elemNameID,
			PrefixR,
			id,
		),
	)
}

// Clone creates a deep copy of this LegacyDrawingHF element.
func (lh *LegacyDrawingHF) Clone() openxml.Element {
	cloned := lh.CompositeElementBase.Clone()

	return &LegacyDrawingHF{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this LegacyDrawingHF element.
func (lh *LegacyDrawingHF) CloneNode(
	deep bool,
) openxml.Element {
	cloned := lh.CompositeElementBase.CloneNode(
		deep,
	)

	return &LegacyDrawingHF{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
