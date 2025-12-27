package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Dxfs represents the differential formats container element (x:dxfs).
// These formats are used for conditional formatting and other purposes
// where only partial formatting needs to be specified.
type Dxfs struct {
	*openxml.CompositeElementBase
}

// NewDxfs creates a new Dxfs element.
func NewDxfs() *Dxfs {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"dxfs",
		PrefixDefault,
	)

	return &Dxfs{CompositeElementBase: elem}
}

// Count returns the count attribute value.
func (d *Dxfs) Count() uint32 {
	attr, found := d.GetAttribute("count", "")
	if !found {
		return 0
	}
	//nolint:revive // add-constant: 10 and 32 are standard parse parameters
	val, _ := strconv.ParseUint(
		attr.Value(),
		10,
		32,
	)

	return uint32(val)
}

// SetCount sets the count attribute.
func (d *Dxfs) SetCount(count uint32) {
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Dxfs returns an iterator over all Dxf elements.
func (d *Dxfs) Dxfs() iter.Seq[*Dxf] {
	return func(yield func(*Dxf) bool) {
		for child := range d.Children() {
			if child.LocalName() != "dxf" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var dxf *Dxf
			if dx, ok := child.(*Dxf); ok {
				dxf = dx
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				dxf = &Dxf{CompositeElementBase: comp}
			}
			if dxf != nil && !yield(dxf) {
				return
			}
		}
	}
}

// GetDxf returns the Dxf at the given index, or nil if out of range.
func (d *Dxfs) GetDxf(index uint32) *Dxf {
	var i uint32
	for dxf := range d.Dxfs() {
		if i == index {
			return dxf
		}
		i++
	}

	return nil
}

// AddDxf adds a new Dxf and returns it.
func (d *Dxfs) AddDxf() *Dxf {
	dxf := NewDxf()
	d.AppendChild(dxf)
	d.SetCount(d.Count() + 1)

	return dxf
}

// ItemCount returns the actual number of Dxf children.
func (d *Dxfs) ItemCount() int {
	count := 0
	for range d.Dxfs() {
		count++
	}

	return count
}

// Clone creates a deep copy of this Dxfs element.
func (d *Dxfs) Clone() openxml.Element {
	cloned := d.CompositeElementBase.Clone()

	return &Dxfs{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Dxfs element.
func (d *Dxfs) CloneNode(
	deep bool,
) openxml.Element {
	cloned := d.CompositeElementBase.CloneNode(
		deep,
	)

	return &Dxfs{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Dxf represents a differential format element (x:dxf).
// This contains partial formatting that can be applied incrementally.
type Dxf struct {
	*openxml.CompositeElementBase
}

// NewDxf creates a new Dxf element.
func NewDxf() *Dxf {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"dxf",
		PrefixDefault,
	)

	return &Dxf{CompositeElementBase: elem}
}

// Font returns the font element, or nil if not present.
func (d *Dxf) Font() *Font {
	elem := d.GetElement("font", NamespaceSML)
	if elem == nil {
		return nil
	}
	if f, ok := elem.(*Font); ok {
		return f
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Font{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateFont returns the font element, creating if needed.
func (d *Dxf) GetOrCreateFont() *Font {
	f := d.Font()
	if f != nil {
		return f
	}
	f = NewFont()
	// Font should come first
	if first := d.FirstChild(); first != nil {
		d.InsertBefore(f, first)
	} else {
		d.AppendChild(f)
	}

	return f
}

// NumFmt returns the number format element, or nil if not present.
func (d *Dxf) NumFmt() *NumFmt {
	elem := d.GetElement("numFmt", NamespaceSML)
	if elem == nil {
		return nil
	}
	if n, ok := elem.(*NumFmt); ok {
		return n
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &NumFmt{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateNumFmt returns the number format element, creating if needed.
func (d *Dxf) GetOrCreateNumFmt() *NumFmt {
	n := d.NumFmt()
	if n != nil {
		return n
	}
	n = NewNumFmt()
	d.insertDxfElement(n, "numFmt")

	return n
}

// Fill returns the fill element, or nil if not present.
func (d *Dxf) Fill() *Fill {
	elem := d.GetElement("fill", NamespaceSML)
	if elem == nil {
		return nil
	}
	if f, ok := elem.(*Fill); ok {
		return f
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Fill{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateFill returns the fill element, creating if needed.
func (d *Dxf) GetOrCreateFill() *Fill {
	f := d.Fill()
	if f != nil {
		return f
	}
	f = NewFill()
	d.insertDxfElement(f, "fill")

	return f
}

// Border returns the border element, or nil if not present.
func (d *Dxf) Border() *Border {
	elem := d.GetElement("border", NamespaceSML)
	if elem == nil {
		return nil
	}
	if b, ok := elem.(*Border); ok {
		return b
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Border{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateBorder returns the border element, creating if needed.
func (d *Dxf) GetOrCreateBorder() *Border {
	b := d.Border()
	if b != nil {
		return b
	}
	b = NewBorder()
	d.insertDxfElement(b, "border")

	return b
}

// Alignment returns the alignment element, or nil if not present.
func (d *Dxf) Alignment() *Alignment {
	elem := d.GetElement(
		"alignment",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if a, ok := elem.(*Alignment); ok {
		return a
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Alignment{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateAlignment returns the alignment element, creating if needed.
func (d *Dxf) GetOrCreateAlignment() *Alignment {
	a := d.Alignment()
	if a != nil {
		return a
	}
	a = NewAlignment()
	d.insertDxfElement(a, "alignment")

	return a
}

// Protection returns the protection element, or nil if not present.
func (d *Dxf) Protection() *Protection {
	elem := d.GetElement(
		"protection",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if p, ok := elem.(*Protection); ok {
		return p
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Protection{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateProtection returns the protection element, creating if needed.
func (d *Dxf) GetOrCreateProtection() *Protection {
	p := d.Protection()
	if p != nil {
		return p
	}
	p = NewProtection()
	d.insertDxfElement(p, "protection")

	return p
}

// insertDxfElement inserts an element in the correct order.
// Order is: font, numFmt, fill, alignment, border, protection.
func (d *Dxf) insertDxfElement(
	elem openxml.Element,
	name string,
) {
	order := map[string]int{
		"font":       dxfOrderFont,
		"numFmt":     dxfOrderNumFmt,
		"fill":       dxfOrderFill,
		"alignment":  dxfOrderAlignment,
		"border":     dxfOrderBorder,
		"protection": dxfOrderProtection,
	}

	targetOrder := order[name]

	// Find the first element that should come after this one
	for child := range d.Children() {
		if childOrder, ok := order[child.LocalName()]; ok &&
			childOrder > targetOrder {
			d.InsertBefore(elem, child)

			return
		}
	}

	// No element found that should come after, so append
	d.AppendChild(elem)
}

// Clone creates a deep copy of this Dxf element.
func (d *Dxf) Clone() openxml.Element {
	cloned := d.CompositeElementBase.Clone()

	return &Dxf{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Dxf element.
func (d *Dxf) CloneNode(
	deep bool,
) openxml.Element {
	cloned := d.CompositeElementBase.CloneNode(
		deep,
	)

	return &Dxf{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
