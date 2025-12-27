package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Dimension represents the dimension element (x:dimension).
// This element specifies the used range of the worksheet.
type Dimension struct {
	*openxml.CompositeElementBase
}

// NewDimension creates a new Dimension element.
func NewDimension() *Dimension {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"dimension",
		PrefixDefault,
	)

	return &Dimension{CompositeElementBase: elem}
}

// Ref returns the reference string (e.g., "A1:Z100").
func (d *Dimension) Ref() string {
	attr, found := d.GetAttribute("ref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the reference string for the used range.
func (d *Dimension) SetRef(ref string) {
	if ref == "" {
		d.RemoveAttribute("ref", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute("", "ref", "", ref),
	)
}

// Clone creates a deep copy of this Dimension element.
func (d *Dimension) Clone() openxml.Element {
	cloned := d.CompositeElementBase.Clone()

	return &Dimension{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Dimension element.
func (d *Dimension) CloneNode(
	deep bool,
) openxml.Element {
	cloned := d.CompositeElementBase.CloneNode(
		deep,
	)

	return &Dimension{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
