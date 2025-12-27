package elements

//revive:disable:file-length-limit many formula properties

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// FormulaType represents the type of a cell formula.
type FormulaType string

const (
	// FormulaTypeNormal is a normal (non-array, non-shared) formula.
	FormulaTypeNormal FormulaType = "normal"
	// FormulaTypeArray indicates an array formula.
	FormulaTypeArray FormulaType = "array"
	// FormulaTypeDataTable indicates a data table formula.
	FormulaTypeDataTable FormulaType = "dataTable"
	// FormulaTypeShared indicates a shared formula.
	FormulaTypeShared FormulaType = "shared"
)

// CellFormula represents the cell formula element (x:f).
// The formula text is stored as the element's text content.
type CellFormula struct {
	*openxml.LeafElementBase
}

// NewCellFormula creates a new CellFormula element.
func NewCellFormula() *CellFormula {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"f",
		PrefixDefault,
	)

	return &CellFormula{LeafElementBase: elem}
}

// NewCellFormulaWithText creates a new CellFormula element with initial
// formula text.
func NewCellFormulaWithText(
	formula string,
) *CellFormula {
	elem := openxml.NewLeafElementWithText(
		NamespaceSML,
		"f",
		PrefixDefault,
		formula,
	)

	return &CellFormula{LeafElementBase: elem}
}

// Formula returns the formula text.
func (cf *CellFormula) Formula() string {
	return cf.InnerText()
}

// SetFormula sets the formula text.
func (cf *CellFormula) SetFormula(
	formula string,
) {
	cf.SetInnerText(formula)
}

// FormulaType returns the formula type. Attribute: t.
func (cf *CellFormula) FormulaType() FormulaType {
	attr, found := cf.GetAttribute("t", "")
	if !found {
		return FormulaTypeNormal
	}

	return FormulaType(attr.Value())
}

// SetFormulaType sets the formula type. Attribute: t.
func (cf *CellFormula) SetFormulaType(
	formulaType FormulaType,
) {
	if formulaType == "" ||
		formulaType == FormulaTypeNormal {
		cf.RemoveAttribute("t", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"t",
			"",
			string(formulaType),
		),
	)
}

// Ref returns the range for shared/array formulas. Attribute: ref.
func (cf *CellFormula) Ref() string {
	attr, found := cf.GetAttribute("ref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the range for shared/array formulas. Attribute: ref.
func (cf *CellFormula) SetRef(ref string) {
	if ref == "" {
		cf.RemoveAttribute("ref", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"ref",
			"",
			ref,
		),
	)
}

// SharedIndex returns the shared formula index. Attribute: si.
func (cf *CellFormula) SharedIndex() uint32 {
	attr, found := cf.GetAttribute("si", "")
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

// SetSharedIndex sets the shared formula index. Attribute: si.
func (cf *CellFormula) SetSharedIndex(
	index uint32,
) {
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"si",
			"",
			strconv.FormatUint(
				uint64(index),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// Bx returns whether the formula is a link to a 3D reference. Attribute: bx.
func (cf *CellFormula) Bx() bool {
	attr, found := cf.GetAttribute("bx", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetBx sets whether the formula is a link to a 3D reference. Attribute: bx.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CellFormula) SetBx(value bool) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"bx",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("bx", "")
	}
}

// Aca returns whether the formula should recalculate always. Attribute: aca.
func (cf *CellFormula) Aca() bool {
	attr, found := cf.GetAttribute("aca", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAca sets whether the formula should recalculate always. Attribute: aca.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CellFormula) SetAca(value bool) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"aca",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("aca", "")
	}
}

// Ca returns whether the formula needs recalculation. Attribute: ca.
func (cf *CellFormula) Ca() bool {
	attr, found := cf.GetAttribute("ca", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCa sets whether the formula needs recalculation. Attribute: ca.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CellFormula) SetCa(value bool) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"ca",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("ca", "")
	}
}

// Del1 returns whether the first input deleted (data table). Attribute: del1.
func (cf *CellFormula) Del1() bool {
	attr, found := cf.GetAttribute("del1", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDel1 sets whether the first input deleted (data table). Attribute: del1.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CellFormula) SetDel1(value bool) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"del1",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("del1", "")
	}
}

// Del2 returns whether the second input deleted (data table). Attribute: del2.
func (cf *CellFormula) Del2() bool {
	attr, found := cf.GetAttribute("del2", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDel2 sets whether the second input deleted (data table). Attribute: del2.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CellFormula) SetDel2(value bool) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"del2",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("del2", "")
	}
}

// Dt2D returns whether the data table is 2-dimensional. Attribute: dt2D.
func (cf *CellFormula) Dt2D() bool {
	attr, found := cf.GetAttribute("dt2D", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDt2D sets whether the data table is 2-dimensional. Attribute: dt2D.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CellFormula) SetDt2D(value bool) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"dt2D",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("dt2D", "")
	}
}

// Dtr returns whether the data table has row input. Attribute: dtr.
func (cf *CellFormula) Dtr() bool {
	attr, found := cf.GetAttribute("dtr", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDtr sets whether the data table has row input. Attribute: dtr.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *CellFormula) SetDtr(value bool) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"dtr",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("dtr", "")
	}
}

// R1 returns the first input cell reference for data table. Attribute: r1.
func (cf *CellFormula) R1() string {
	attr, found := cf.GetAttribute("r1", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetR1 sets the first input cell reference for data table. Attribute: r1.
func (cf *CellFormula) SetR1(ref string) {
	if ref == "" {
		cf.RemoveAttribute("r1", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"r1",
			"",
			ref,
		),
	)
}

// R2 returns the second input cell reference for data table. Attribute: r2.
func (cf *CellFormula) R2() string {
	attr, found := cf.GetAttribute("r2", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetR2 sets the second input cell reference for data table. Attribute: r2.
func (cf *CellFormula) SetR2(ref string) {
	if ref == "" {
		cf.RemoveAttribute("r2", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"r2",
			"",
			ref,
		),
	)
}

// Clone creates a deep copy of this CellFormula element.
func (cf *CellFormula) Clone() openxml.Element {
	cloned := cf.LeafElementBase.Clone()

	return &CellFormula{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CellFormula element.
func (cf *CellFormula) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cf.LeafElementBase.CloneNode(deep)

	return &CellFormula{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
