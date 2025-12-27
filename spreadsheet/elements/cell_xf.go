package elements

//revive:disable:file-length-limit many cell xf properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Constants for cell format parsing and attributes.
const (
	parseBase10   = 10
	bitSize32     = 32
	attrNameCount = "count"
)

// CellStyleXfs represents the cell style formats container (x:cellStyleXfs).
// This contains base styles that cell formats (cellXfs) can reference.
type CellStyleXfs struct {
	*openxml.CompositeElementBase
}

// NewCellStyleXfs creates a new CellStyleXfs element.
func NewCellStyleXfs() *CellStyleXfs {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"cellStyleXfs",
		PrefixDefault,
	)

	return &CellStyleXfs{
		CompositeElementBase: elem,
	}
}

// Count returns the count attribute value.
func (c *CellStyleXfs) Count() uint32 {
	attr, found := c.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetCount sets the count attribute.
func (c *CellStyleXfs) SetCount(count uint32) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// Xfs returns an iterator over all Xf elements.
func (c *CellStyleXfs) Xfs() iter.Seq[*Xf] {
	return func(yield func(*Xf) bool) {
		for child := range c.Children() {
			if child.LocalName() != "xf" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var xf *Xf
			if x, ok := child.(*Xf); ok {
				xf = x
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				xf = &Xf{CompositeElementBase: comp}
			}
			if xf != nil && !yield(xf) {
				return
			}
		}
	}
}

// GetXf returns the Xf at the given index, or nil if out of range.
func (c *CellStyleXfs) GetXf(index uint32) *Xf {
	var i uint32
	for xf := range c.Xfs() {
		if i == index {
			return xf
		}
		i++
	}

	return nil
}

// AddXf adds a new Xf and returns it.
func (c *CellStyleXfs) AddXf() *Xf {
	xf := NewXf()
	c.AppendChild(xf)
	c.SetCount(c.Count() + 1)

	return xf
}

// ItemCount returns the actual number of Xf children.
func (c *CellStyleXfs) ItemCount() int {
	count := 0
	for range c.Xfs() {
		count++
	}

	return count
}

// Clone creates a deep copy of this CellStyleXfs element.
func (c *CellStyleXfs) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &CellStyleXfs{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CellStyleXfs element.
func (c *CellStyleXfs) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &CellStyleXfs{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CellXfs represents the cell formats container (x:cellXfs).
// This contains the actual cell formats referenced by cells.
type CellXfs struct {
	*openxml.CompositeElementBase
}

// NewCellXfs creates a new CellXfs element.
func NewCellXfs() *CellXfs {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"cellXfs",
		PrefixDefault,
	)

	return &CellXfs{CompositeElementBase: elem}
}

// Count returns the count attribute value.
func (c *CellXfs) Count() uint32 {
	attr, found := c.GetAttribute(
		attrNameCount,
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetCount sets the count attribute.
func (c *CellXfs) SetCount(count uint32) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			attrNameCount,
			"",
			strconv.FormatUint(
				uint64(count),
				parseBase10,
			),
		),
	)
}

// Xfs returns an iterator over all Xf elements.
func (c *CellXfs) Xfs() iter.Seq[*Xf] {
	return func(yield func(*Xf) bool) {
		for child := range c.Children() {
			if child.LocalName() != "xf" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var xf *Xf
			if x, ok := child.(*Xf); ok {
				xf = x
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				xf = &Xf{CompositeElementBase: comp}
			}
			if xf != nil && !yield(xf) {
				return
			}
		}
	}
}

// GetXf returns the Xf at the given index, or nil if out of range.
func (c *CellXfs) GetXf(index uint32) *Xf {
	var i uint32
	for xf := range c.Xfs() {
		if i == index {
			return xf
		}
		i++
	}

	return nil
}

// AddXf adds a new Xf and returns it.
func (c *CellXfs) AddXf() *Xf {
	xf := NewXf()
	c.AppendChild(xf)
	c.SetCount(c.Count() + 1)

	return xf
}

// ItemCount returns the actual number of Xf children.
func (c *CellXfs) ItemCount() int {
	count := 0
	for range c.Xfs() {
		count++
	}

	return count
}

// Clone creates a deep copy of this CellXfs element.
func (c *CellXfs) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &CellXfs{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CellXfs element.
func (c *CellXfs) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &CellXfs{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Xf represents a cell format element (x:xf).
// This element contains formatting information for cells.
type Xf struct {
	*openxml.CompositeElementBase
}

// NewXf creates a new Xf element.
func NewXf() *Xf {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"xf",
		PrefixDefault,
	)

	return &Xf{CompositeElementBase: elem}
}

// NumFmtId returns the number format ID.
func (x *Xf) NumFmtId() uint32 {
	attr, found := x.GetAttribute("numFmtId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetNumFmtId sets the number format ID.
func (x *Xf) SetNumFmtId(id uint32) {
	x.SetAttribute(
		openxml.NewAttribute(
			"",
			"numFmtId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// FontId returns the font ID.
func (x *Xf) FontId() uint32 {
	attr, found := x.GetAttribute("fontId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetFontId sets the font ID.
func (x *Xf) SetFontId(id uint32) {
	x.SetAttribute(
		openxml.NewAttribute(
			"",
			"fontId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// FillId returns the fill ID.
func (x *Xf) FillId() uint32 {
	attr, found := x.GetAttribute("fillId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetFillId sets the fill ID.
func (x *Xf) SetFillId(id uint32) {
	x.SetAttribute(
		openxml.NewAttribute(
			"",
			"fillId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// BorderId returns the border ID.
func (x *Xf) BorderId() uint32 {
	attr, found := x.GetAttribute("borderId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant: base 10
		32, //nolint:revive // add-constant: bit size
	)

	return uint32(val)
}

// SetBorderId sets the border ID.
func (x *Xf) SetBorderId(id uint32) {
	x.SetAttribute(
		openxml.NewAttribute(
			"",
			"borderId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// XfId returns the cell style xf ID (reference to cellStyleXfs).
func (x *Xf) XfId() uint32 {
	attr, found := x.GetAttribute("xfId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant: base 10
		32, //nolint:revive // add-constant: bit size
	)

	return uint32(val)
}

// SetXfId sets the cell style xf ID.
func (x *Xf) SetXfId(id uint32) {
	x.SetAttribute(
		openxml.NewAttribute(
			"",
			"xfId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// QuotePrefix returns whether text values start with a quote.
func (x *Xf) QuotePrefix() bool {
	attr, found := x.GetAttribute(
		"quotePrefix",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetQuotePrefix sets whether text values start with a quote.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (x *Xf) SetQuotePrefix(quote bool) {
	if quote {
		x.SetAttribute(
			openxml.NewAttribute(
				"",
				"quotePrefix",
				"",
				attrValueTrue,
			),
		)
	} else {
		x.RemoveAttribute("quotePrefix", "")
	}
}

// PivotButton returns whether this is a pivot table button format.
func (x *Xf) PivotButton() bool {
	attr, found := x.GetAttribute(
		"pivotButton",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPivotButton sets whether this is a pivot table button format.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (x *Xf) SetPivotButton(pivot bool) {
	if pivot {
		x.SetAttribute(
			openxml.NewAttribute(
				"",
				"pivotButton",
				"",
				attrValueTrue,
			),
		)
	} else {
		x.RemoveAttribute("pivotButton", "")
	}
}

// ApplyNumberFormat returns whether to apply the number format.
func (x *Xf) ApplyNumberFormat() bool {
	attr, found := x.GetAttribute(
		"applyNumberFormat",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyNumberFormat sets whether to apply the number format.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (x *Xf) SetApplyNumberFormat(apply bool) {
	if apply {
		x.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyNumberFormat",
				"",
				attrValueTrue,
			),
		)
	} else {
		x.RemoveAttribute("applyNumberFormat", "")
	}
}

// ApplyFont returns whether to apply the font.
func (x *Xf) ApplyFont() bool {
	attr, found := x.GetAttribute("applyFont", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyFont sets whether to apply the font.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (x *Xf) SetApplyFont(apply bool) {
	if apply {
		x.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyFont",
				"",
				attrValueTrue,
			),
		)
	} else {
		x.RemoveAttribute("applyFont", "")
	}
}

// ApplyFill returns whether to apply the fill.
func (x *Xf) ApplyFill() bool {
	attr, found := x.GetAttribute("applyFill", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyFill sets whether to apply the fill.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (x *Xf) SetApplyFill(apply bool) {
	if apply {
		x.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyFill",
				"",
				attrValueTrue,
			),
		)
	} else {
		x.RemoveAttribute("applyFill", "")
	}
}

// ApplyBorder returns whether to apply the border.
func (x *Xf) ApplyBorder() bool {
	attr, found := x.GetAttribute(
		"applyBorder",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyBorder sets whether to apply the border.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (x *Xf) SetApplyBorder(apply bool) {
	if apply {
		x.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyBorder",
				"",
				attrValueTrue,
			),
		)
	} else {
		x.RemoveAttribute("applyBorder", "")
	}
}

// ApplyAlignment returns whether to apply the alignment.
func (x *Xf) ApplyAlignment() bool {
	attr, found := x.GetAttribute(
		"applyAlignment",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyAlignment sets whether to apply the alignment.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (x *Xf) SetApplyAlignment(apply bool) {
	if apply {
		x.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyAlignment",
				"",
				attrValueTrue,
			),
		)
	} else {
		x.RemoveAttribute("applyAlignment", "")
	}
}

// ApplyProtection returns whether to apply the protection.
func (x *Xf) ApplyProtection() bool {
	attr, found := x.GetAttribute(
		"applyProtection",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyProtection sets whether to apply the protection.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (x *Xf) SetApplyProtection(apply bool) {
	if apply {
		x.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyProtection",
				"",
				attrValueTrue,
			),
		)
	} else {
		x.RemoveAttribute("applyProtection", "")
	}
}

// Alignment returns the alignment element, or nil if not present.
func (x *Xf) Alignment() *Alignment {
	elem := x.GetElement(
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
func (x *Xf) GetOrCreateAlignment() *Alignment {
	a := x.Alignment()
	if a != nil {
		return a
	}
	a = NewAlignment()
	// Alignment comes before protection
	p := x.Protection()
	if p != nil {
		x.InsertBefore(a, p)
	} else {
		x.AppendChild(a)
	}

	return a
}

// Protection returns the protection element, or nil if not present.
func (x *Xf) Protection() *Protection {
	elem := x.GetElement(
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
func (x *Xf) GetOrCreateProtection() *Protection {
	p := x.Protection()
	if p != nil {
		return p
	}
	p = NewProtection()
	x.AppendChild(p)

	return p
}

// Clone creates a deep copy of this Xf element.
func (x *Xf) Clone() openxml.Element {
	cloned := x.CompositeElementBase.Clone()

	return &Xf{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Xf element.
func (x *Xf) CloneNode(
	deep bool,
) openxml.Element {
	cloned := x.CompositeElementBase.CloneNode(
		deep,
	)

	return &Xf{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
