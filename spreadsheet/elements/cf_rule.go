package elements

//revive:disable:file-length-limit many cf rule properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// CfRule represents a conditional formatting rule element (x:cfRule).
type CfRule struct {
	*openxml.CompositeElementBase
}

// NewCfRule creates a new CfRule element.
func NewCfRule() *CfRule {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"cfRule",
		PrefixDefault,
	)

	return &CfRule{CompositeElementBase: elem}
}

// Type returns the type of the rule. Attribute: type.
func (r *CfRule) Type() CfRuleType {
	attr, found := r.GetAttribute("type", "")
	if !found {
		return ""
	}

	return CfRuleType(attr.Value())
}

// SetType sets the type of the rule. Attribute: type.
func (r *CfRule) SetType(ruleType CfRuleType) {
	if ruleType == "" {
		r.RemoveAttribute("type", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			string(ruleType),
		),
	)
}

// DxfId returns the differential formatting ID. Attribute: dxfId.
func (r *CfRule) DxfId() uint32 {
	attr, found := r.GetAttribute("dxfId", "")
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

// SetDxfId sets the differential formatting ID. Attribute: dxfId.
func (r *CfRule) SetDxfId(id uint32) {
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"dxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Priority returns the priority of the rule. Attribute: priority.
// Lower values have higher priority.
func (r *CfRule) Priority() int32 {
	attr, found := r.GetAttribute("priority", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return int32(val)
}

// SetPriority sets the priority of the rule. Attribute: priority.
func (r *CfRule) SetPriority(priority int32) {
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"priority",
			"",
			strconv.FormatInt(
				int64(priority),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// StopIfTrue returns whether to stop evaluating rules after this one matches.
// Attribute: stopIfTrue.
func (r *CfRule) StopIfTrue() bool {
	attr, found := r.GetAttribute(
		"stopIfTrue",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetStopIfTrue sets whether to stop evaluating rules after this one matches.
// Attribute: stopIfTrue.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *CfRule) SetStopIfTrue(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"stopIfTrue",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("stopIfTrue", "")
	}
}

// AboveAverage returns whether the rule applies to values above average.
// Attribute: aboveAverage. Default is true.
func (r *CfRule) AboveAverage() bool {
	attr, found := r.GetAttribute(
		"aboveAverage",
		"",
	)
	if !found {
		return true
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAboveAverage sets whether the rule applies to values above average.
// Attribute: aboveAverage.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *CfRule) SetAboveAverage(value bool) {
	if value {
		r.RemoveAttribute("aboveAverage", "")
	} else {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"aboveAverage",
				"",
				attrValueFalse,
			),
		)
	}
}

// Percent returns whether the rank value is a percentage. Attribute: percent.
func (r *CfRule) Percent() bool {
	attr, found := r.GetAttribute("percent", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPercent sets whether the rank value is a percentage. Attribute: percent.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *CfRule) SetPercent(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"percent",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("percent", "")
	}
}

// Bottom returns whether the rule applies to bottom values instead of top.
// Attribute: bottom.
func (r *CfRule) Bottom() bool {
	attr, found := r.GetAttribute("bottom", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetBottom sets whether the rule applies to bottom values instead of top.
// Attribute: bottom.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *CfRule) SetBottom(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"bottom",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("bottom", "")
	}
}

// Operator returns the comparison operator. Attribute: operator.
func (r *CfRule) Operator() CfRuleOperator {
	attr, found := r.GetAttribute("operator", "")
	if !found {
		return ""
	}

	return CfRuleOperator(attr.Value())
}

// SetOperator sets the comparison operator. Attribute: operator.
func (r *CfRule) SetOperator(op CfRuleOperator) {
	if op == "" {
		r.RemoveAttribute("operator", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"operator",
			"",
			string(op),
		),
	)
}

// Text returns the text value for text-based rules. Attribute: text.
func (r *CfRule) Text() string {
	attr, found := r.GetAttribute("text", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetText sets the text value for text-based rules. Attribute: text.
func (r *CfRule) SetText(text string) {
	if text == "" {
		r.RemoveAttribute("text", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"text",
			"",
			text,
		),
	)
}

// TimePeriod returns the time period for time-based rules.
// Attribute: timePeriod.
func (r *CfRule) TimePeriod() TimePeriodType {
	attr, found := r.GetAttribute(
		"timePeriod",
		"",
	)
	if !found {
		return ""
	}

	return TimePeriodType(attr.Value())
}

// SetTimePeriod sets the time period for time-based rules.
// Attribute: timePeriod.
func (r *CfRule) SetTimePeriod(
	period TimePeriodType,
) {
	if period == "" {
		r.RemoveAttribute("timePeriod", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"timePeriod",
			"",
			string(period),
		),
	)
}

// Rank returns the rank value for top/bottom rules. Attribute: rank.
func (r *CfRule) Rank() uint32 {
	attr, found := r.GetAttribute("rank", "")
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

// SetRank sets the rank value for top/bottom rules. Attribute: rank.
func (r *CfRule) SetRank(rank uint32) {
	if rank == 0 {
		r.RemoveAttribute("rank", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"rank",
			"",
			strconv.FormatUint(
				uint64(rank),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// StdDev returns the number of standard deviations for above/below
// average rules. Attribute: stdDev.
func (r *CfRule) StdDev() int32 {
	attr, found := r.GetAttribute("stdDev", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10, //nolint:revive // add-constant: base 10
		32, //nolint:revive // add-constant: bit size
	)

	return int32(val)
}

// SetStdDev sets the number of standard deviations for above/below
// average rules. Attribute: stdDev.
func (r *CfRule) SetStdDev(stdDev int32) {
	if stdDev == 0 {
		r.RemoveAttribute("stdDev", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"stdDev",
			"",
			strconv.FormatInt(
				int64(stdDev),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// EqualAverage returns whether the average itself is included.
// Attribute: equalAverage.
func (r *CfRule) EqualAverage() bool {
	attr, found := r.GetAttribute(
		"equalAverage",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetEqualAverage sets whether the average itself is included.
// Attribute: equalAverage.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *CfRule) SetEqualAverage(value bool) {
	if value {
		r.SetAttribute(
			openxml.NewAttribute(
				"",
				"equalAverage",
				"",
				attrValueTrue,
			),
		)
	} else {
		r.RemoveAttribute("equalAverage", "")
	}
}

// Formulas returns an iterator over all Formula elements in this rule.
func (r *CfRule) Formulas() iter.Seq[*CfFormula] {
	return func(yield func(*CfFormula) bool) {
		for child := range r.Children() {
			if child.LocalName() != "formula" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var formula *CfFormula
			if f, ok := child.(*CfFormula); ok {
				formula = f
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				formula = &CfFormula{LeafElementBase: leaf}
			}
			if formula != nil && !yield(formula) {
				return
			}
		}
	}
}

// FormulaCount returns the number of Formula elements.
func (r *CfRule) FormulaCount() int {
	count := 0
	for range r.Formulas() {
		count++
	}

	return count
}

// AddFormula adds a new Formula element to the rule.
func (r *CfRule) AddFormula(
	formula string,
) *CfFormula {
	f := NewCfFormulaWithText(formula)
	// Insert formulas before colorScale, dataBar, iconSet
	var insertBefore openxml.Element
	for child := range r.Children() {
		localName := child.LocalName()
		if (localName == "colorScale" ||
			localName == "dataBar" ||
			localName == "iconSet") &&
			child.NamespaceURI() == NamespaceSML {
			insertBefore = child

			break
		}
	}
	if insertBefore != nil {
		r.InsertBefore(f, insertBefore)
	} else {
		r.AppendChild(f)
	}

	return f
}

// ColorScale returns the ColorScale child element, or nil if not present.
func (r *CfRule) ColorScale() *ColorScale {
	elem := r.GetElement(
		"colorScale",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if cs, ok := elem.(*ColorScale); ok {
		return cs
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ColorScale{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateColorScale returns the ColorScale child element,
// creating it if needed.
func (r *CfRule) GetOrCreateColorScale() *ColorScale {
	cs := r.ColorScale()
	if cs != nil {
		return cs
	}
	cs = NewColorScale()
	r.AppendChild(cs)

	return cs
}

// DataBar returns the DataBar child element, or nil if not present.
func (r *CfRule) DataBar() *DataBar {
	elem := r.GetElement("dataBar", NamespaceSML)
	if elem == nil {
		return nil
	}
	if db, ok := elem.(*DataBar); ok {
		return db
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DataBar{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDataBar returns the DataBar child element, creating it if needed.
func (r *CfRule) GetOrCreateDataBar() *DataBar {
	db := r.DataBar()
	if db != nil {
		return db
	}
	db = NewDataBar()
	r.AppendChild(db)

	return db
}

// IconSet returns the IconSet child element, or nil if not present.
func (r *CfRule) IconSet() *IconSet {
	elem := r.GetElement("iconSet", NamespaceSML)
	if elem == nil {
		return nil
	}
	if is, ok := elem.(*IconSet); ok {
		return is
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &IconSet{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateIconSet returns the IconSet child element, creating it if needed.
func (r *CfRule) GetOrCreateIconSet() *IconSet {
	is := r.IconSet()
	if is != nil {
		return is
	}
	is = NewIconSet()
	r.AppendChild(is)

	return is
}

// Clone creates a deep copy of this CfRule element.
func (r *CfRule) Clone() openxml.Element {
	cloned := r.CompositeElementBase.Clone()

	return &CfRule{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CfRule element.
func (r *CfRule) CloneNode(
	deep bool,
) openxml.Element {
	cloned := r.CompositeElementBase.CloneNode(
		deep,
	)

	return &CfRule{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
