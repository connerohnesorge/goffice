package elements

//revive:disable:file-length-limit many break properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// RowBreaks represents the row breaks container element (x:rowBreaks).
type RowBreaks struct {
	*openxml.CompositeElementBase
}

// NewRowBreaks creates a new RowBreaks element.
func NewRowBreaks() *RowBreaks {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"rowBreaks",
		PrefixDefault,
	)

	return &RowBreaks{CompositeElementBase: elem}
}

// Count returns the total number of row breaks.
func (rb *RowBreaks) Count() int {
	attr, found := rb.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCount sets the total number of row breaks.
func (rb *RowBreaks) SetCount(count int) {
	if count == 0 {
		rb.RemoveAttribute("count", "")

		return
	}
	rb.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.Itoa(count),
		),
	)
}

// ManualBreakCount returns the number of manual breaks.
func (rb *RowBreaks) ManualBreakCount() int {
	attr, found := rb.GetAttribute(
		"manualBreakCount",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetManualBreakCount sets the number of manual breaks.
func (rb *RowBreaks) SetManualBreakCount(
	count int,
) {
	if count == 0 {
		rb.RemoveAttribute("manualBreakCount", "")

		return
	}
	rb.SetAttribute(
		openxml.NewAttribute(
			"",
			"manualBreakCount",
			"",
			strconv.Itoa(count),
		),
	)
}

// Breaks returns an iterator over all Break elements.
func (rb *RowBreaks) Breaks() iter.Seq[*Break] {
	return func(yield func(*Break) bool) {
		for child := range rb.Children() {
			if child.LocalName() != "brk" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var brk *Break
			if b, ok := child.(*Break); ok {
				brk = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				brk = &Break{CompositeElementBase: comp}
			}
			if brk != nil && !yield(brk) {
				return
			}
		}
	}
}

// AddBreak adds a new row break and returns it.
func (rb *RowBreaks) AddBreak(row int) *Break {
	brk := NewBreak()
	brk.SetId(row)
	brk.SetMan(true)
	rb.AppendChild(brk)

	// Update counts
	count := 0
	manCount := 0
	for b := range rb.Breaks() {
		count++
		if b.Man() {
			manCount++
		}
	}
	rb.SetCount(count)
	rb.SetManualBreakCount(manCount)

	return brk
}

// Clone creates a deep copy of this RowBreaks element.
func (rb *RowBreaks) Clone() openxml.Element {
	cloned := rb.CompositeElementBase.Clone()

	return &RowBreaks{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this RowBreaks element.
func (rb *RowBreaks) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rb.CompositeElementBase.CloneNode(
		deep,
	)

	return &RowBreaks{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ColBreaks represents the column breaks container element (x:colBreaks).
type ColBreaks struct {
	*openxml.CompositeElementBase
}

// NewColBreaks creates a new ColBreaks element.
func NewColBreaks() *ColBreaks {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"colBreaks",
		PrefixDefault,
	)

	return &ColBreaks{CompositeElementBase: elem}
}

// Count returns the total number of column breaks.
func (cb *ColBreaks) Count() int {
	attr, found := cb.GetAttribute(
		"count", //nolint:revive // add-constant
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCount sets the total number of column breaks.
func (cb *ColBreaks) SetCount(count int) {
	if count == 0 {
		cb.RemoveAttribute("count", "")

		return
	}
	cb.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.Itoa(count),
		),
	)
}

// ManualBreakCount returns the number of manual breaks.
func (cb *ColBreaks) ManualBreakCount() int {
	attr, found := cb.GetAttribute(
		"manualBreakCount", //nolint:revive // add-constant: attribute name
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetManualBreakCount sets the number of manual breaks.
func (cb *ColBreaks) SetManualBreakCount(
	count int,
) {
	if count == 0 {
		cb.RemoveAttribute("manualBreakCount", "")

		return
	}
	cb.SetAttribute(
		openxml.NewAttribute(
			"",
			"manualBreakCount",
			"",
			strconv.Itoa(count),
		),
	)
}

// Breaks returns an iterator over all Break elements.
func (cb *ColBreaks) Breaks() iter.Seq[*Break] {
	return func(yield func(*Break) bool) {
		for child := range cb.Children() {
			if child.LocalName() != "brk" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var brk *Break
			if b, ok := child.(*Break); ok {
				brk = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				brk = &Break{CompositeElementBase: comp}
			}
			if brk != nil && !yield(brk) {
				return
			}
		}
	}
}

// AddBreak adds a new column break and returns it.
func (cb *ColBreaks) AddBreak(col int) *Break {
	brk := NewBreak()
	brk.SetId(col)
	brk.SetMan(true)
	cb.AppendChild(brk)

	// Update counts
	count := 0
	manCount := 0
	for b := range cb.Breaks() {
		count++
		if b.Man() {
			manCount++
		}
	}
	cb.SetCount(count)
	cb.SetManualBreakCount(manCount)

	return brk
}

// Clone creates a deep copy of this ColBreaks element.
func (cb *ColBreaks) Clone() openxml.Element {
	cloned := cb.CompositeElementBase.Clone()

	return &ColBreaks{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ColBreaks element.
func (cb *ColBreaks) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cb.CompositeElementBase.CloneNode(
		deep,
	)

	return &ColBreaks{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Break represents a page break element (x:brk).
type Break struct {
	*openxml.CompositeElementBase
}

// NewBreak creates a new Break element.
func NewBreak() *Break {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"brk",
		PrefixDefault,
	)

	return &Break{CompositeElementBase: elem}
}

// Id returns the row or column index of the break (0-based).
func (b *Break) Id() int {
	attr, found := b.GetAttribute("id", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetId sets the row or column index of the break.
func (b *Break) SetId(id int) {
	b.SetAttribute(
		openxml.NewAttribute(
			"",
			"id",
			"",
			strconv.Itoa(id),
		),
	)
}

// Min returns the minimum row/column for the break.
func (b *Break) Min() int {
	attr, found := b.GetAttribute("min", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetMin sets the minimum row/column for the break.
func (b *Break) SetMin(minVal int) {
	if minVal == 0 {
		b.RemoveAttribute("min", "")

		return
	}
	b.SetAttribute(
		openxml.NewAttribute(
			"",
			"min",
			"",
			strconv.Itoa(minVal),
		),
	)
}

// Max returns the maximum row/column for the break.
func (b *Break) Max() int {
	attr, found := b.GetAttribute("max", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetMax sets the maximum row/column for the break.
func (b *Break) SetMax(maxVal int) {
	if maxVal == 0 {
		b.RemoveAttribute("max", "")

		return
	}
	b.SetAttribute(
		openxml.NewAttribute(
			"",
			"max",
			"",
			strconv.Itoa(maxVal),
		),
	)
}

// Man returns whether this is a manual break.
func (b *Break) Man() bool {
	attr, found := b.GetAttribute("man", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetMan sets whether this is a manual break.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (b *Break) SetMan(value bool) {
	if value {
		b.SetAttribute(
			openxml.NewAttribute(
				"",
				"man",
				"",
				attrValueTrue,
			),
		)
	} else {
		b.RemoveAttribute("man", "")
	}
}

// Pt returns whether this is a pivot table break.
func (b *Break) Pt() bool {
	attr, found := b.GetAttribute("pt", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPt sets whether this is a pivot table break.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (b *Break) SetPt(value bool) {
	if value {
		b.SetAttribute(
			openxml.NewAttribute(
				"",
				"pt",
				"",
				attrValueTrue,
			),
		)
	} else {
		b.RemoveAttribute("pt", "")
	}
}

// Clone creates a deep copy of this Break element.
func (b *Break) Clone() openxml.Element {
	cloned := b.CompositeElementBase.Clone()

	return &Break{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Break element.
func (b *Break) CloneNode(
	deep bool,
) openxml.Element {
	cloned := b.CompositeElementBase.CloneNode(
		deep,
	)

	return &Break{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
