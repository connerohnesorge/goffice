package elements

//revive:disable:file-length-limit many calc properties

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// CalcMode represents the calculation mode for the workbook.
type CalcMode string

const (
	// CalcModeManual requires manual calculation trigger.
	CalcModeManual CalcMode = "manual"
	// CalcModeAuto calculates automatically when data changes.
	CalcModeAuto CalcMode = "auto"
	// CalcModeAutoNoTable calculates automatically except for tables.
	CalcModeAutoNoTable CalcMode = "autoNoTable"
)

// RefMode represents the reference style for formulas.
type RefMode string

const (
	// RefModeA1 uses A1-style references (default).
	RefModeA1 RefMode = "A1"
	// RefModeR1C1 uses R1C1-style references.
	RefModeR1C1 RefMode = "R1C1"
)

const (
	// defaultIterateCount is the default maximum number of iterations
	// for iterative calculation.
	defaultIterateCount = 100
	// defaultIterateDelta is the default maximum change between iterations.
	defaultIterateDelta = 0.001
)

// CalcPr represents the calculation properties element (x:calcPr).
type CalcPr struct {
	*openxml.CompositeElementBase
}

// NewCalcPr creates a new CalcPr element.
func NewCalcPr() *CalcPr {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"calcPr",
		PrefixDefault,
	)

	return &CalcPr{CompositeElementBase: elem}
}

// CalcId returns the calculation engine version ID.
func (cp *CalcPr) CalcId() int {
	attr, found := cp.GetAttribute("calcId", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCalcId sets the calculation engine version ID.
func (cp *CalcPr) SetCalcId(id int) {
	if id == 0 {
		cp.RemoveAttribute("calcId", "")

		return
	}
	cp.SetAttribute(
		openxml.NewAttribute(
			"",
			"calcId",
			"",
			strconv.Itoa(id),
		),
	)
}

// CalcMode returns the calculation mode.
func (cp *CalcPr) CalcMode() CalcMode {
	attr, found := cp.GetAttribute("calcMode", "")
	if !found {
		return CalcModeAuto
	}

	return CalcMode(attr.Value())
}

// SetCalcMode sets the calculation mode.
func (cp *CalcPr) SetCalcMode(mode CalcMode) {
	if mode == "" || mode == CalcModeAuto {
		cp.RemoveAttribute("calcMode", "")

		return
	}
	cp.SetAttribute(
		openxml.NewAttribute(
			"",
			"calcMode",
			"",
			string(mode),
		),
	)
}

// FullCalcOnLoad returns whether to perform full calculation on load.
func (cp *CalcPr) FullCalcOnLoad() bool {
	attr, found := cp.GetAttribute(
		"fullCalcOnLoad",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetFullCalcOnLoad sets whether to perform full calculation on load.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cp *CalcPr) SetFullCalcOnLoad(value bool) {
	if value {
		cp.SetAttribute(
			openxml.NewAttribute(
				"",
				"fullCalcOnLoad",
				"",
				attrValueTrue,
			),
		)
	} else {
		cp.RemoveAttribute("fullCalcOnLoad", "")
	}
}

// RefMode returns the reference mode (A1 or R1C1).
func (cp *CalcPr) RefMode() RefMode {
	attr, found := cp.GetAttribute("refMode", "")
	if !found {
		return RefModeA1
	}

	return RefMode(attr.Value())
}

// SetRefMode sets the reference mode.
func (cp *CalcPr) SetRefMode(mode RefMode) {
	if mode == "" || mode == RefModeA1 {
		cp.RemoveAttribute("refMode", "")

		return
	}
	cp.SetAttribute(
		openxml.NewAttribute(
			"",
			"refMode",
			"",
			string(mode),
		),
	)
}

// Iterate returns whether iterative calculation is enabled.
func (cp *CalcPr) Iterate() bool {
	attr, found := cp.GetAttribute("iterate", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetIterate sets whether iterative calculation is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cp *CalcPr) SetIterate(value bool) {
	if value {
		cp.SetAttribute(
			openxml.NewAttribute(
				"",
				"iterate",
				"",
				attrValueTrue,
			),
		)
	} else {
		cp.RemoveAttribute("iterate", "")
	}
}

// IterateCount returns the maximum number of iterations.
func (cp *CalcPr) IterateCount() int {
	attr, found := cp.GetAttribute(
		"iterateCount",
		"",
	)
	if !found {
		return defaultIterateCount
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetIterateCount sets the maximum number of iterations.
func (cp *CalcPr) SetIterateCount(count int) {
	if count == defaultIterateCount {
		cp.RemoveAttribute("iterateCount", "")

		return
	}
	attr := openxml.NewAttribute(
		"",
		"iterateCount",
		"",
		strconv.Itoa(count),
	)
	cp.SetAttribute(attr)
}

// IterateDelta returns the maximum change between iterations.
func (cp *CalcPr) IterateDelta() float64 {
	attr, found := cp.GetAttribute(
		"iterateDelta",
		"",
	)
	if !found {
		return defaultIterateDelta
	}
	//nolint:revive // 64 is the standard bit size for float64
	val, _ := strconv.ParseFloat(attr.Value(), 64)

	return val
}

// SetIterateDelta sets the maximum change between iterations.
func (cp *CalcPr) SetIterateDelta(delta float64) {
	if delta == defaultIterateDelta {
		cp.RemoveAttribute("iterateDelta", "")

		return
	}
	//nolint:revive // 64 is the standard bit size for float64
	floatStr := strconv.FormatFloat(
		delta,
		'g',
		-1,
		64,
	)
	attr := openxml.NewAttribute(
		"",
		"iterateDelta",
		"",
		floatStr,
	)
	cp.SetAttribute(attr)
}

// FullPrecision returns whether full precision is used for calculations.
func (cp *CalcPr) FullPrecision() bool {
	attr, found := cp.GetAttribute(
		"fullPrecision",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetFullPrecision sets whether full precision is used.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cp *CalcPr) SetFullPrecision(value bool) {
	if value {
		cp.RemoveAttribute("fullPrecision", "")
	} else {
		cp.SetAttribute(openxml.NewAttribute("", "fullPrecision", "", attrValueFalse))
	}
}

// CalcCompleted returns whether calculation is completed.
func (cp *CalcPr) CalcCompleted() bool {
	attr, found := cp.GetAttribute(
		"calcCompleted",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetCalcCompleted sets whether calculation is completed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cp *CalcPr) SetCalcCompleted(value bool) {
	if value {
		cp.RemoveAttribute("calcCompleted", "")
	} else {
		cp.SetAttribute(openxml.NewAttribute("", "calcCompleted", "", attrValueFalse))
	}
}

// CalcOnSave returns whether to calculate before saving.
func (cp *CalcPr) CalcOnSave() bool {
	attr, found := cp.GetAttribute(
		"calcOnSave",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetCalcOnSave sets whether to calculate before saving.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cp *CalcPr) SetCalcOnSave(value bool) {
	if value {
		cp.RemoveAttribute("calcOnSave", "")
	} else {
		cp.SetAttribute(openxml.NewAttribute("", "calcOnSave", "", attrValueFalse))
	}
}

// ConcurrentCalc returns whether concurrent calculation is enabled.
func (cp *CalcPr) ConcurrentCalc() bool {
	attr, found := cp.GetAttribute(
		"concurrentCalc",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetConcurrentCalc sets whether concurrent calculation is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cp *CalcPr) SetConcurrentCalc(value bool) {
	if value {
		cp.RemoveAttribute("concurrentCalc", "")
	} else {
		cp.SetAttribute(openxml.NewAttribute("", "concurrentCalc", "", attrValueFalse))
	}
}

// ConcurrentManualCount returns the number of manual calculation threads.
func (cp *CalcPr) ConcurrentManualCount() int {
	attr, found := cp.GetAttribute(
		"concurrentManualCount",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetConcurrentManualCount sets the number of manual calculation threads.
func (cp *CalcPr) SetConcurrentManualCount(
	count int,
) {
	if count == 0 {
		cp.RemoveAttribute(
			"concurrentManualCount",
			"",
		)

		return
	}
	attr := openxml.NewAttribute(
		"",
		"concurrentManualCount",
		"",
		strconv.Itoa(count),
	)
	cp.SetAttribute(attr)
}

// ForceFullCalc returns whether to force full calculation.
func (cp *CalcPr) ForceFullCalc() bool {
	attr, found := cp.GetAttribute(
		"forceFullCalc",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetForceFullCalc sets whether to force full calculation.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cp *CalcPr) SetForceFullCalc(value bool) {
	if value {
		cp.SetAttribute(
			openxml.NewAttribute(
				"",
				"forceFullCalc",
				"",
				attrValueTrue,
			),
		)
	} else {
		cp.RemoveAttribute("forceFullCalc", "")
	}
}

// Clone creates a deep copy of this CalcPr element.
func (cp *CalcPr) Clone() openxml.Element {
	cloned := cp.CompositeElementBase.Clone()

	return &CalcPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CalcPr element.
func (cp *CalcPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cp.CompositeElementBase.CloneNode(
		deep,
	)

	return &CalcPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
