// and effects.
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Attribute name constants for shape style elements.
const (
	attrIdx = "idx"
)

// ShapeStyle represents the shape style element (a:style).
// This element contains references to style matrix entries.
type ShapeStyle struct {
	*openxml.CompositeElementBase
}

// NewShapeStyle creates a new shape style element.
func NewShapeStyle() *ShapeStyle {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"style",
		PrefixMain,
	)

	return &ShapeStyle{CompositeElementBase: elem}
}

// SetLineReference sets the line style reference.
func (s *ShapeStyle) SetLineReference(
	idx int,
	color SchemeColorValue,
) {
	s.setStyleReference("lnRef", idx, color)
}

// SetFillReference sets the fill style reference.
func (s *ShapeStyle) SetFillReference(
	idx int,
	color SchemeColorValue,
) {
	s.setStyleReference("fillRef", idx, color)
}

// SetEffectReference sets the effect style reference.
func (s *ShapeStyle) SetEffectReference(
	idx int,
	color SchemeColorValue,
) {
	s.setStyleReference("effectRef", idx, color)
}

// SetFontReference sets the font style reference.
func (s *ShapeStyle) SetFontReference(
	idx string,
	color SchemeColorValue,
) {
	// Remove existing
	if existing := s.GetElement("fontRef", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}

	fontRef := openxml.NewCompositeElement(
		NamespaceMain,
		"fontRef",
		PrefixMain,
	)
	fontRef.SetAttribute(
		openxml.NewAttribute(
			"",
			attrIdx,
			"",
			idx,
		),
	)

	if color != "" {
		sc := NewSchemeColor(color)
		fontRef.AppendChild(sc)
	}

	s.AppendChild(fontRef)
}

// setStyleReference sets a style matrix reference element.
func (s *ShapeStyle) setStyleReference(
	name string,
	idx int,
	color SchemeColorValue,
) {
	// Remove existing
	if existing := s.GetElement(name, NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}

	ref := openxml.NewCompositeElement(
		NamespaceMain,
		name,
		PrefixMain,
	)
	ref.SetAttribute(
		openxml.NewAttribute(
			"",
			attrIdx,
			"",
			strconv.Itoa(idx),
		),
	)

	if color != "" {
		sc := NewSchemeColor(color)
		ref.AppendChild(sc)
	}

	s.AppendChild(ref)
}

// Clone creates a deep copy of this ShapeStyle element.
func (s *ShapeStyle) Clone() openxml.Element {
	cloned := s.CompositeElementBase.Clone()

	return &ShapeStyle{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// StyleMatrixReference represents a reference to a style matrix entry.
type StyleMatrixReference struct {
	*openxml.CompositeElementBase
}

// NewStyleMatrixReference creates a new style matrix reference with the
// given index.
func NewStyleMatrixReference(
	localName string,
	idx int,
) *StyleMatrixReference {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		localName,
		PrefixMain,
	)
	ref := &StyleMatrixReference{
		CompositeElementBase: elem,
	}
	ref.SetIndex(idx)

	return ref
}

// Index returns the style matrix index.
func (r *StyleMatrixReference) Index() int {
	attr, found := r.GetAttribute(attrIdx, "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetIndex sets the style matrix index.
func (r *StyleMatrixReference) SetIndex(idx int) {
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			attrIdx,
			"",
			strconv.Itoa(idx),
		),
	)
}

// SetSchemeColor sets the scheme color for the reference.
func (r *StyleMatrixReference) SetSchemeColor(
	color SchemeColorValue,
) {
	// Remove existing color children
	r.RemoveAllChildren()

	if color != "" {
		sc := NewSchemeColor(color)
		r.AppendChild(sc)
	}
}

// Clone creates a deep copy of this StyleMatrixReference element.
func (r *StyleMatrixReference) Clone() openxml.Element {
	cloned := r.CompositeElementBase.Clone()

	return &StyleMatrixReference{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
