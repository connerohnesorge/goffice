package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// ConditionalFormatting represents a conditional formatting element
// (x:conditionalFormatting). This element contains one or more formatting
// rules applied to a range of cells.
type ConditionalFormatting struct {
	*openxml.CompositeElementBase
}

// NewConditionalFormatting creates a new ConditionalFormatting element.
func NewConditionalFormatting() *ConditionalFormatting {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"conditionalFormatting",
		PrefixDefault,
	)

	return &ConditionalFormatting{
		CompositeElementBase: elem,
	}
}

// Sqref returns the cell ranges to which the formatting applies.
// Multiple ranges are space-separated (e.g., "A1:B10 C1:D10").
// Attribute: sqref.
func (cf *ConditionalFormatting) Sqref() string {
	attr, found := cf.GetAttribute("sqref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSqref sets the cell ranges to which the formatting applies.
// Multiple ranges are space-separated (e.g., "A1:B10 C1:D10").
// Attribute: sqref.
func (cf *ConditionalFormatting) SetSqref(
	sqref string,
) {
	if sqref == "" {
		cf.RemoveAttribute("sqref", "")

		return
	}
	cf.SetAttribute(
		openxml.NewAttribute(
			"",
			"sqref",
			"",
			sqref,
		),
	)
}

// Pivot returns whether this conditional formatting is associated with a
// PivotTable. Attribute: pivot.
func (cf *ConditionalFormatting) Pivot() bool {
	attr, found := cf.GetAttribute("pivot", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPivot sets whether this conditional formatting is associated with a PivotTable.
// Attribute: pivot.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *ConditionalFormatting) SetPivot(
	value bool,
) {
	if value {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"pivot",
				"",
				attrValueTrue,
			),
		)
	} else {
		cf.RemoveAttribute("pivot", "")
	}
}

// Rules returns an iterator over all CfRule elements
// in this conditional formatting.
func (cf *ConditionalFormatting) Rules() iter.Seq[*CfRule] {
	return func(yield func(*CfRule) bool) {
		for child := range cf.Children() {
			if child.LocalName() != "cfRule" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var rule *CfRule
			if r, ok := child.(*CfRule); ok {
				rule = r
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				rule = &CfRule{CompositeElementBase: comp}
			}
			if rule != nil && !yield(rule) {
				return
			}
		}
	}
}

// RuleCount returns the number of CfRule elements.
func (cf *ConditionalFormatting) RuleCount() int {
	count := 0
	for range cf.Rules() {
		count++
	}

	return count
}

// AddRule adds a new CfRule element to the conditional formatting.
func (cf *ConditionalFormatting) AddRule() *CfRule {
	rule := NewCfRule()
	cf.AppendChild(rule)

	return rule
}

// AddCellIsRule adds a new cell value comparison rule.
//
//nolint:revive // argument-limit
func (cf *ConditionalFormatting) AddCellIsRule(
	operator CfRuleOperator,
	formula1 string,
	formula2 string,
	dxfId uint32,
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeCellIs)
	rule.SetOperator(operator)
	rule.SetDxfId(dxfId)
	rule.SetPriority(priority)

	if formula1 != "" {
		rule.AddFormula(formula1)
	}
	if formula2 != "" {
		rule.AddFormula(formula2)
	}

	return rule
}

// AddExpressionRule adds a new expression-based rule.
func (cf *ConditionalFormatting) AddExpressionRule(
	formula string,
	dxfId uint32,
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeExpression)
	rule.SetDxfId(dxfId)
	rule.SetPriority(priority)
	rule.AddFormula(formula)

	return rule
}

// AddColorScaleRule adds a new color scale rule.
func (cf *ConditionalFormatting) AddColorScaleRule(
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeColorScale)
	rule.SetPriority(priority)
	rule.GetOrCreateColorScale()

	return rule
}

// AddDataBarRule adds a new data bar rule.
func (cf *ConditionalFormatting) AddDataBarRule(
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeDataBar)
	rule.SetPriority(priority)
	rule.GetOrCreateDataBar()

	return rule
}

// AddIconSetRule adds a new icon set rule.
func (cf *ConditionalFormatting) AddIconSetRule(
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeIconSet)
	rule.SetPriority(priority)
	rule.GetOrCreateIconSet()

	return rule
}

// AddTop10Rule adds a new top/bottom 10 rule.
//
//nolint:revive // argument-limit
func (cf *ConditionalFormatting) AddTop10Rule(
	rank uint32,
	bottom bool,
	percent bool,
	dxfId uint32,
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeTop10)
	rule.SetRank(rank)
	rule.SetBottom(bottom)
	rule.SetPercent(percent)
	rule.SetDxfId(dxfId)
	rule.SetPriority(priority)

	return rule
}

// AddDuplicateValuesRule adds a new duplicate values rule.
func (cf *ConditionalFormatting) AddDuplicateValuesRule(
	dxfId uint32,
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeDuplicateValues)
	rule.SetDxfId(dxfId)
	rule.SetPriority(priority)

	return rule
}

// AddUniqueValuesRule adds a new unique values rule.
func (cf *ConditionalFormatting) AddUniqueValuesRule(
	dxfId uint32,
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeUniqueValues)
	rule.SetDxfId(dxfId)
	rule.SetPriority(priority)

	return rule
}

// AddContainsTextRule adds a new "contains text" rule.
func (cf *ConditionalFormatting) AddContainsTextRule(
	text string,
	dxfId uint32,
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeContainsText)
	rule.SetOperator(CfRuleOperatorContainsText)
	rule.SetText(text)
	rule.SetDxfId(dxfId)
	rule.SetPriority(priority)

	return rule
}

// AddAboveAverageRule adds a new above/below average rule.
//
//nolint:revive // argument-limit, enforce-repeated-arg-type-style: API design
func (cf *ConditionalFormatting) AddAboveAverageRule(
	aboveAverage bool,
	equalAverage bool,
	stdDev int32,
	dxfId uint32,
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeAboveAverage)
	rule.SetAboveAverage(aboveAverage)
	rule.SetEqualAverage(equalAverage)
	if stdDev != 0 {
		rule.SetStdDev(stdDev)
	}
	rule.SetDxfId(dxfId)
	rule.SetPriority(priority)

	return rule
}

// AddTimePeriodRule adds a new time period rule.
func (cf *ConditionalFormatting) AddTimePeriodRule(
	timePeriod TimePeriodType,
	dxfId uint32,
	priority int32,
) *CfRule {
	rule := cf.AddRule()
	rule.SetType(CfRuleTypeTimePeriod)
	rule.SetTimePeriod(timePeriod)
	rule.SetDxfId(dxfId)
	rule.SetPriority(priority)

	return rule
}

// Clone creates a deep copy of this ConditionalFormatting element.
func (cf *ConditionalFormatting) Clone() openxml.Element {
	cloned := cf.CompositeElementBase.Clone()

	return &ConditionalFormatting{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ConditionalFormatting element.
func (cf *ConditionalFormatting) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cf.CompositeElementBase.CloneNode(
		deep,
	)

	return &ConditionalFormatting{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
