package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// CfFormula represents a formula element (x:formula) used in conditional
// formatting rules. The formula text is stored as the element's text content.
type CfFormula struct {
	*openxml.LeafElementBase
}

// NewCfFormula creates a new CfFormula element.
func NewCfFormula() *CfFormula {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"formula",
		PrefixDefault,
	)

	return &CfFormula{LeafElementBase: elem}
}

// NewCfFormulaWithText creates a new CfFormula element with initial formula
// text.
func NewCfFormulaWithText(
	formula string,
) *CfFormula {
	elem := openxml.NewLeafElementWithText(
		NamespaceSML,
		"formula",
		PrefixDefault,
		formula,
	)

	return &CfFormula{LeafElementBase: elem}
}

// Formula returns the formula text.
func (f *CfFormula) Formula() string {
	return f.InnerText()
}

// SetFormula sets the formula text.
func (f *CfFormula) SetFormula(
	formula string,
) {
	f.SetInnerText(formula)
}

// Clone creates a deep copy of this CfFormula element.
func (f *CfFormula) Clone() openxml.Element {
	cloned := f.LeafElementBase.Clone()

	return &CfFormula{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CfFormula element.
func (f *CfFormula) CloneNode(
	deep bool,
) openxml.Element {
	cloned := f.LeafElementBase.CloneNode(deep)

	return &CfFormula{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
