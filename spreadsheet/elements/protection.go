package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Protection represents the protection element (x:protection) in cell
// formatting. This element controls whether a cell is locked or its formula
// is hidden.
type Protection struct {
	*openxml.LeafElementBase
}

// NewProtection creates a new Protection element.
func NewProtection() *Protection {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"protection",
		PrefixDefault,
	)

	return &Protection{LeafElementBase: elem}
}

// Locked returns whether the cell is locked.
// Default is true when the sheet is protected.
func (p *Protection) Locked() bool {
	attr, found := p.GetAttribute("locked", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetLocked sets whether the cell is locked.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *Protection) SetLocked(locked bool) {
	if locked {
		p.RemoveAttribute(
			"locked",
			"",
		) // true is default
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "locked", "", attrValueFalse),
		)
	}
}

// Hidden returns whether the cell's formula is hidden.
// Default is false.
func (p *Protection) Hidden() bool {
	attr, found := p.GetAttribute("hidden", "")
	if !found {
		return false // Default is false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHidden sets whether the cell's formula is hidden.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *Protection) SetHidden(hidden bool) {
	if hidden {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"hidden",
				"",
				attrValueTrue,
			),
		)
	} else {
		p.RemoveAttribute("hidden", "") // false is default
	}
}

// Clone creates a deep copy of this Protection element.
func (p *Protection) Clone() openxml.Element {
	cloned := p.LeafElementBase.Clone()

	return &Protection{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Protection element.
func (p *Protection) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.LeafElementBase.CloneNode(deep)

	return &Protection{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
