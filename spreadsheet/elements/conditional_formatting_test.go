package elements

import (
	"strings"
	"testing"
)

func TestNewColor(t *testing.T) {
	color := NewColor()

	if color == nil {
		t.Fatal("expected non-nil Color")
	}
	if color.LocalName() != "color" {
		t.Errorf(
			"expected localName 'color', got '%s'",
			color.LocalName(),
		)
	}
	if color.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceSML,
			color.NamespaceURI(),
		)
	}
}

func TestColor_Attributes(t *testing.T) {
	color := NewColor()

	// Test Auto
	if color.Auto() {
		t.Error(
			"expected auto to be false initially",
		)
	}
	color.SetAuto(true)
	if !color.Auto() {
		t.Error("expected auto to be true")
	}

	// Test Indexed
	color.SetIndexed(5)
	if color.Indexed() != 5 {
		t.Errorf(
			"expected indexed 5, got %d",
			color.Indexed(),
		)
	}

	// Test RGB
	color.SetRGB("FF0000FF")
	if color.RGB() != "FF0000FF" {
		t.Errorf(
			"expected RGB 'FF0000FF', got '%s'",
			color.RGB(),
		)
	}

	// Test Theme
	color.SetTheme(3)
	if color.Theme() != 3 {
		t.Errorf(
			"expected theme 3, got %d",
			color.Theme(),
		)
	}

	// Test Tint
	color.SetTint(-0.5)
	if color.Tint() != -0.5 {
		t.Errorf(
			"expected tint -0.5, got %f",
			color.Tint(),
		)
	}
}

func TestNewCfvo(t *testing.T) {
	cfvo := NewCfvo()

	if cfvo == nil {
		t.Fatal("expected non-nil Cfvo")
	}
	if cfvo.LocalName() != "cfvo" {
		t.Errorf(
			"expected localName 'cfvo', got '%s'",
			cfvo.LocalName(),
		)
	}
}

func TestCfvo_Attributes(t *testing.T) {
	cfvo := NewCfvo()

	// Test Type
	cfvo.SetType(CfvoTypePercent)
	if cfvo.Type() != CfvoTypePercent {
		t.Errorf(
			"expected type 'percent', got '%s'",
			cfvo.Type(),
		)
	}

	// Test Val
	cfvo.SetVal("50")
	if cfvo.Val() != "50" {
		t.Errorf(
			"expected val '50', got '%s'",
			cfvo.Val(),
		)
	}

	// Test Gte (default is true)
	if !cfvo.Gte() {
		t.Error(
			"expected gte to be true by default",
		)
	}
	cfvo.SetGte(false)
	if cfvo.Gte() {
		t.Error("expected gte to be false")
	}
}

func TestNewCfFormula(t *testing.T) {
	formula := NewCfFormula()

	if formula == nil {
		t.Fatal("expected non-nil CfFormula")
	}
	if formula.LocalName() != "formula" {
		t.Errorf(
			"expected localName 'formula', got '%s'",
			formula.LocalName(),
		)
	}
}

func TestCfFormula_WithText(t *testing.T) {
	formula := NewCfFormulaWithText("$A$1>100")

	if formula.Formula() != "$A$1>100" {
		t.Errorf(
			"expected formula '$A$1>100', got '%s'",
			formula.Formula(),
		)
	}

	formula.SetFormula("$B$2<50")
	if formula.Formula() != "$B$2<50" {
		t.Errorf(
			"expected formula '$B$2<50', got '%s'",
			formula.Formula(),
		)
	}
}

func TestNewColorScale(t *testing.T) {
	cs := NewColorScale()

	if cs == nil {
		t.Fatal("expected non-nil ColorScale")
	}
	if cs.LocalName() != "colorScale" {
		t.Errorf(
			"expected localName 'colorScale', got '%s'",
			cs.LocalName(),
		)
	}
}

func TestColorScale_TwoColorScale(t *testing.T) {
	cs := NewColorScale()

	cs.SetTwoColorScale(
		CfvoTypeMin, "",
		"FFFF0000",
		CfvoTypeMax, "",
		"FF00FF00",
	)

	if cs.CfvoCount() != 2 {
		t.Errorf(
			"expected 2 cfvo elements, got %d",
			cs.CfvoCount(),
		)
	}
	if cs.ColorCount() != 2 {
		t.Errorf(
			"expected 2 color elements, got %d",
			cs.ColorCount(),
		)
	}
}

func TestColorScale_ThreeColorScale(
	t *testing.T,
) {
	cs := NewColorScale()

	cs.SetThreeColorScale(
		CfvoTypeMin, "", "FFFF0000",
		CfvoTypePercentile, "50", "FFFFFF00",
		CfvoTypeMax, "", "FF00FF00",
	)

	if cs.CfvoCount() != 3 {
		t.Errorf(
			"expected 3 cfvo elements, got %d",
			cs.CfvoCount(),
		)
	}
	if cs.ColorCount() != 3 {
		t.Errorf(
			"expected 3 color elements, got %d",
			cs.ColorCount(),
		)
	}
}

func TestNewDataBar(t *testing.T) {
	db := NewDataBar()

	if db == nil {
		t.Fatal("expected non-nil DataBar")
	}
	if db.LocalName() != "dataBar" {
		t.Errorf(
			"expected localName 'dataBar', got '%s'",
			db.LocalName(),
		)
	}
}

func TestDataBar_DefaultValues(t *testing.T) {
	db := NewDataBar()

	// Test defaults
	if db.MinLength() != 10 {
		t.Errorf(
			"expected minLength 10, got %d",
			db.MinLength(),
		)
	}
	if db.MaxLength() != 90 {
		t.Errorf(
			"expected maxLength 90, got %d",
			db.MaxLength(),
		)
	}
	if !db.ShowValue() {
		t.Error(
			"expected showValue to be true by default",
		)
	}
	if !db.Gradient() {
		t.Error(
			"expected gradient to be true by default",
		)
	}
}

func TestDataBar_SetDataBar(t *testing.T) {
	db := NewDataBar()

	db.SetDataBar(
		CfvoTypeMin, "",
		CfvoTypeMax, "",
		"FF638EC6",
	)

	// Check cfvo count
	count := 0
	for range db.Cfvos() {
		count++
	}
	if count != 2 {
		t.Errorf(
			"expected 2 cfvo elements, got %d",
			count,
		)
	}

	// Check color
	color := db.Color()
	if color == nil {
		t.Fatal("expected non-nil color")
	}
	if color.RGB() != "FF638EC6" {
		t.Errorf(
			"expected RGB 'FF638EC6', got '%s'",
			color.RGB(),
		)
	}
}

func TestDataBar_Attributes(t *testing.T) {
	db := NewDataBar()

	// Test MinLength
	db.SetMinLength(5)
	if db.MinLength() != 5 {
		t.Errorf(
			"expected minLength 5, got %d",
			db.MinLength(),
		)
	}

	// Test MaxLength
	db.SetMaxLength(95)
	if db.MaxLength() != 95 {
		t.Errorf(
			"expected maxLength 95, got %d",
			db.MaxLength(),
		)
	}

	// Test ShowValue
	db.SetShowValue(false)
	if db.ShowValue() {
		t.Error("expected showValue to be false")
	}

	// Test Gradient
	db.SetGradient(false)
	if db.Gradient() {
		t.Error("expected gradient to be false")
	}

	// Test Direction
	db.SetDirection(DataBarDirectionLTR)
	if db.Direction() != DataBarDirectionLTR {
		t.Errorf(
			"expected direction 'leftToRight', got '%s'",
			db.Direction(),
		)
	}

	// Test AxisPosition
	db.SetAxisPosition(DataBarAxisPositionMiddle)
	if db.AxisPosition() != DataBarAxisPositionMiddle {
		t.Errorf(
			"expected axisPosition 'middle', got '%s'",
			db.AxisPosition(),
		)
	}
}

func TestNewIconSet(t *testing.T) {
	is := NewIconSet()

	if is == nil {
		t.Fatal("expected non-nil IconSet")
	}
	if is.LocalName() != "iconSet" {
		t.Errorf(
			"expected localName 'iconSet', got '%s'",
			is.LocalName(),
		)
	}
}

func TestIconSet_DefaultValues(t *testing.T) {
	is := NewIconSet()

	// Default icon set type
	if is.IconSetType() != IconSet3TrafficLights1 {
		t.Errorf(
			"expected iconSet '3TrafficLights1', got '%s'",
			is.IconSetType(),
		)
	}
	if !is.ShowValue() {
		t.Error(
			"expected showValue to be true by default",
		)
	}
	if !is.Percent() {
		t.Error(
			"expected percent to be true by default",
		)
	}
	if is.Reverse() {
		t.Error(
			"expected reverse to be false by default",
		)
	}
}

func TestIconSet_Set3IconSet(t *testing.T) {
	is := NewIconSet()

	is.Set3IconSet(
		IconSet3Arrows,
		CfvoTypePercent, "33",
		CfvoTypePercent, "67",
	)

	if is.IconSetType() != IconSet3Arrows {
		t.Errorf(
			"expected iconSet '3Arrows', got '%s'",
			is.IconSetType(),
		)
	}
	if is.CfvoCount() != 3 {
		t.Errorf(
			"expected 3 cfvo elements, got %d",
			is.CfvoCount(),
		)
	}
}

func TestNewCfRule(t *testing.T) {
	rule := NewCfRule()

	if rule == nil {
		t.Fatal("expected non-nil CfRule")
	}
	if rule.LocalName() != "cfRule" {
		t.Errorf(
			"expected localName 'cfRule', got '%s'",
			rule.LocalName(),
		)
	}
}

func TestCfRule_TypeAndPriority(t *testing.T) {
	rule := NewCfRule()

	rule.SetType(CfRuleTypeCellIs)
	if rule.Type() != CfRuleTypeCellIs {
		t.Errorf(
			"expected type 'cellIs', got '%s'",
			rule.Type(),
		)
	}

	rule.SetPriority(1)
	if rule.Priority() != 1 {
		t.Errorf(
			"expected priority 1, got %d",
			rule.Priority(),
		)
	}
}

func TestCfRule_CellIsRule(t *testing.T) {
	rule := NewCfRule()

	rule.SetType(CfRuleTypeCellIs)
	rule.SetOperator(CfRuleOperatorGreaterThan)
	rule.SetDxfId(0)
	rule.SetPriority(1)
	rule.AddFormula("100")

	if rule.Operator() != CfRuleOperatorGreaterThan {
		t.Errorf(
			"expected operator 'greaterThan', got '%s'",
			rule.Operator(),
		)
	}
	if rule.FormulaCount() != 1 {
		t.Errorf(
			"expected 1 formula, got %d",
			rule.FormulaCount(),
		)
	}
}

func TestCfRule_Top10Rule(t *testing.T) {
	rule := NewCfRule()

	rule.SetType(CfRuleTypeTop10)
	rule.SetRank(10)
	rule.SetBottom(true)
	rule.SetPercent(true)

	if rule.Rank() != 10 {
		t.Errorf(
			"expected rank 10, got %d",
			rule.Rank(),
		)
	}
	if !rule.Bottom() {
		t.Error("expected bottom to be true")
	}
	if !rule.Percent() {
		t.Error("expected percent to be true")
	}
}

func TestCfRule_AboveAverageRule(t *testing.T) {
	rule := NewCfRule()

	rule.SetType(CfRuleTypeAboveAverage)
	rule.SetAboveAverage(false)
	rule.SetEqualAverage(true)
	rule.SetStdDev(2)

	if rule.AboveAverage() {
		t.Error(
			"expected aboveAverage to be false",
		)
	}
	if !rule.EqualAverage() {
		t.Error(
			"expected equalAverage to be true",
		)
	}
	if rule.StdDev() != 2 {
		t.Errorf(
			"expected stdDev 2, got %d",
			rule.StdDev(),
		)
	}
}

func TestCfRule_TextRule(t *testing.T) {
	rule := NewCfRule()

	rule.SetType(CfRuleTypeContainsText)
	rule.SetText("error")

	if rule.Text() != "error" {
		t.Errorf(
			"expected text 'error', got '%s'",
			rule.Text(),
		)
	}
}

func TestCfRule_TimePeriodRule(t *testing.T) {
	rule := NewCfRule()

	rule.SetType(CfRuleTypeTimePeriod)
	rule.SetTimePeriod(TimePeriodToday)

	if rule.TimePeriod() != TimePeriodToday {
		t.Errorf(
			"expected timePeriod 'today', got '%s'",
			rule.TimePeriod(),
		)
	}
}

func TestCfRule_ChildElements(t *testing.T) {
	rule := NewCfRule()

	// Test ColorScale
	cs := rule.GetOrCreateColorScale()
	if cs == nil {
		t.Fatal("expected non-nil ColorScale")
	}
	cs2 := rule.ColorScale()
	if cs2 == nil {
		t.Error("expected to find ColorScale")
	}

	// Reset
	rule.RemoveAllChildren()

	// Test DataBar
	db := rule.GetOrCreateDataBar()
	if db == nil {
		t.Fatal("expected non-nil DataBar")
	}
	db2 := rule.DataBar()
	if db2 == nil {
		t.Error("expected to find DataBar")
	}

	// Reset
	rule.RemoveAllChildren()

	// Test IconSet
	is := rule.GetOrCreateIconSet()
	if is == nil {
		t.Fatal("expected non-nil IconSet")
	}
	is2 := rule.IconSet()
	if is2 == nil {
		t.Error("expected to find IconSet")
	}
}

func TestNewConditionalFormatting(t *testing.T) {
	cf := NewConditionalFormatting()

	if cf == nil {
		t.Fatal(
			"expected non-nil ConditionalFormatting",
		)
	}
	if cf.LocalName() != "conditionalFormatting" {
		t.Errorf(
			"expected localName 'conditionalFormatting', got '%s'",
			cf.LocalName(),
		)
	}
}

func TestConditionalFormatting_Sqref(
	t *testing.T,
) {
	cf := NewConditionalFormatting()

	cf.SetSqref("A1:B10")
	if cf.Sqref() != "A1:B10" {
		t.Errorf(
			"expected sqref 'A1:B10', got '%s'",
			cf.Sqref(),
		)
	}

	// Multiple ranges
	cf.SetSqref("A1:B10 C1:D10")
	if cf.Sqref() != "A1:B10 C1:D10" {
		t.Errorf(
			"expected sqref 'A1:B10 C1:D10', got '%s'",
			cf.Sqref(),
		)
	}
}

func TestConditionalFormatting_Pivot(
	t *testing.T,
) {
	cf := NewConditionalFormatting()

	if cf.Pivot() {
		t.Error(
			"expected pivot to be false initially",
		)
	}
	cf.SetPivot(true)
	if !cf.Pivot() {
		t.Error("expected pivot to be true")
	}
}

func TestConditionalFormatting_AddRule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddRule()
	if rule == nil {
		t.Fatal("expected non-nil CfRule")
	}

	if cf.RuleCount() != 1 {
		t.Errorf(
			"expected 1 rule, got %d",
			cf.RuleCount(),
		)
	}
}

func TestConditionalFormatting_AddCellIsRule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddCellIsRule(
		CfRuleOperatorGreaterThan,
		"100",
		"",
		0,
		1,
	)

	if rule.Type() != CfRuleTypeCellIs {
		t.Errorf(
			"expected type 'cellIs', got '%s'",
			rule.Type(),
		)
	}
	if rule.Operator() != CfRuleOperatorGreaterThan {
		t.Errorf(
			"expected operator 'greaterThan', got '%s'",
			rule.Operator(),
		)
	}
}

func TestConditionalFormatting_AddExpressionRule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddExpressionRule(
		"$A1>$B1",
		0,
		1,
	)

	if rule.Type() != CfRuleTypeExpression {
		t.Errorf(
			"expected type 'expression', got '%s'",
			rule.Type(),
		)
	}
	if rule.FormulaCount() != 1 {
		t.Errorf(
			"expected 1 formula, got %d",
			rule.FormulaCount(),
		)
	}
}

func TestConditionalFormatting_AddColorScaleRule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddColorScaleRule(1)

	if rule.Type() != CfRuleTypeColorScale {
		t.Errorf(
			"expected type 'colorScale', got '%s'",
			rule.Type(),
		)
	}
	if rule.ColorScale() == nil {
		t.Error("expected non-nil ColorScale")
	}
}

func TestConditionalFormatting_AddDataBarRule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddDataBarRule(1)

	if rule.Type() != CfRuleTypeDataBar {
		t.Errorf(
			"expected type 'dataBar', got '%s'",
			rule.Type(),
		)
	}
	if rule.DataBar() == nil {
		t.Error("expected non-nil DataBar")
	}
}

func TestConditionalFormatting_AddIconSetRule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddIconSetRule(1)

	if rule.Type() != CfRuleTypeIconSet {
		t.Errorf(
			"expected type 'iconSet', got '%s'",
			rule.Type(),
		)
	}
	if rule.IconSet() == nil {
		t.Error("expected non-nil IconSet")
	}
}

func TestConditionalFormatting_AddTop10Rule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddTop10Rule(
		10,
		false,
		false,
		0,
		1,
	)

	if rule.Type() != CfRuleTypeTop10 {
		t.Errorf(
			"expected type 'top10', got '%s'",
			rule.Type(),
		)
	}
	if rule.Rank() != 10 {
		t.Errorf(
			"expected rank 10, got %d",
			rule.Rank(),
		)
	}
}

func TestConditionalFormatting_AddDuplicateValuesRule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddDuplicateValuesRule(0, 1)

	if rule.Type() != CfRuleTypeDuplicateValues {
		t.Errorf(
			"expected type 'duplicateValues', got '%s'",
			rule.Type(),
		)
	}
}

func TestConditionalFormatting_AddContainsTextRule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddContainsTextRule("error", 0, 1)

	if rule.Type() != CfRuleTypeContainsText {
		t.Errorf(
			"expected type 'containsText', got '%s'",
			rule.Type(),
		)
	}
	if rule.Text() != "error" {
		t.Errorf(
			"expected text 'error', got '%s'",
			rule.Text(),
		)
	}
}

func TestConditionalFormatting_AddAboveAverageRule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddAboveAverageRule(
		true,
		false,
		0,
		0,
		1,
	)

	if rule.Type() != CfRuleTypeAboveAverage {
		t.Errorf(
			"expected type 'aboveAverage', got '%s'",
			rule.Type(),
		)
	}
}

func TestConditionalFormatting_AddTimePeriodRule(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	rule := cf.AddTimePeriodRule(
		TimePeriodToday,
		0,
		1,
	)

	if rule.Type() != CfRuleTypeTimePeriod {
		t.Errorf(
			"expected type 'timePeriod', got '%s'",
			rule.Type(),
		)
	}
	if rule.TimePeriod() != TimePeriodToday {
		t.Errorf(
			"expected timePeriod 'today', got '%s'",
			rule.TimePeriod(),
		)
	}
}

func TestConditionalFormatting_XMLSerialization(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")

	// Add a cell value rule
	rule := cf.AddCellIsRule(
		CfRuleOperatorGreaterThan,
		"100",
		"",
		0,
		1,
	)
	rule.SetStopIfTrue(true)

	xml := cf.OuterXml()

	// Verify key elements are present
	if !strings.Contains(
		xml,
		"conditionalFormatting",
	) {
		t.Error(
			"expected 'conditionalFormatting' in XML",
		)
	}
	if !strings.Contains(
		xml,
		"sqref=\"A1:A10\"",
	) {
		t.Error(
			"expected 'sqref=\"A1:A10\"' in XML",
		)
	}
	if !strings.Contains(xml, "cfRule") {
		t.Error("expected 'cfRule' in XML")
	}
	if !strings.Contains(xml, "type=\"cellIs\"") {
		t.Error(
			"expected 'type=\"cellIs\"' in XML",
		)
	}
	if !strings.Contains(
		xml,
		"operator=\"greaterThan\"",
	) {
		t.Error(
			"expected 'operator=\"greaterThan\"' in XML",
		)
	}
	if !strings.Contains(xml, "formula") {
		t.Error("expected 'formula' in XML")
	}
}

func TestConditionalFormatting_ColorScaleXML(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("B1:B10")

	rule := cf.AddColorScaleRule(1)
	cs := rule.ColorScale()
	cs.SetTwoColorScale(
		CfvoTypeMin, "",
		"FFFF0000",
		CfvoTypeMax, "",
		"FF00FF00",
	)

	xml := cf.OuterXml()

	if !strings.Contains(xml, "colorScale") {
		t.Error("expected 'colorScale' in XML")
	}
	if !strings.Contains(xml, "cfvo") {
		t.Error("expected 'cfvo' in XML")
	}
	if !strings.Contains(xml, "color") {
		t.Error("expected 'color' in XML")
	}
	if !strings.Contains(xml, "type=\"min\"") {
		t.Error("expected 'type=\"min\"' in XML")
	}
	if !strings.Contains(xml, "type=\"max\"") {
		t.Error("expected 'type=\"max\"' in XML")
	}
}

func TestConditionalFormatting_Clone(
	t *testing.T,
) {
	cf := NewConditionalFormatting()
	cf.SetSqref("A1:A10")
	cf.AddCellIsRule(
		CfRuleOperatorGreaterThan,
		"100",
		"",
		0,
		1,
	)

	cloned := cf.Clone()
	clone, ok := cloned.(*ConditionalFormatting)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *ConditionalFormatting",
			cloned,
		)
	}
	if clone.Sqref() != "A1:A10" {
		t.Errorf(
			"expected sqref 'A1:A10', got '%s'",
			clone.Sqref(),
		)
	}
	if clone.RuleCount() != 1 {
		t.Errorf(
			"expected 1 rule in clone, got %d",
			clone.RuleCount(),
		)
	}

	// Modify original, verify clone is independent
	cf.SetSqref("B1:B10")
	if clone.Sqref() != "A1:A10" {
		t.Error(
			"expected clone to remain unchanged after modifying original",
		)
	}
}

func TestCfRuleTypes(t *testing.T) {
	types := []CfRuleType{
		CfRuleTypeExpression,
		CfRuleTypeCellIs,
		CfRuleTypeColorScale,
		CfRuleTypeDataBar,
		CfRuleTypeIconSet,
		CfRuleTypeTop10,
		CfRuleTypeUniqueValues,
		CfRuleTypeDuplicateValues,
		CfRuleTypeContainsText,
		CfRuleTypeNotContainsText,
		CfRuleTypeBeginsWith,
		CfRuleTypeEndsWith,
		CfRuleTypeContainsBlanks,
		CfRuleTypeNotContainsBlanks,
		CfRuleTypeContainsErrors,
		CfRuleTypeNotContainsErrors,
		CfRuleTypeTimePeriod,
		CfRuleTypeAboveAverage,
	}

	for _, ruleType := range types {
		rule := NewCfRule()
		rule.SetType(ruleType)
		if rule.Type() != ruleType {
			t.Errorf(
				"expected type '%s', got '%s'",
				ruleType,
				rule.Type(),
			)
		}
	}
}

func TestCfRuleOperators(t *testing.T) {
	operators := []CfRuleOperator{
		CfRuleOperatorLessThan,
		CfRuleOperatorLessThanOrEqual,
		CfRuleOperatorEqual,
		CfRuleOperatorNotEqual,
		CfRuleOperatorGreaterThanOrEqual,
		CfRuleOperatorGreaterThan,
		CfRuleOperatorBetween,
		CfRuleOperatorNotBetween,
		CfRuleOperatorContainsText,
		CfRuleOperatorNotContains,
		CfRuleOperatorBeginsWith,
		CfRuleOperatorEndsWith,
	}

	for _, op := range operators {
		rule := NewCfRule()
		rule.SetOperator(op)
		if rule.Operator() != op {
			t.Errorf(
				"expected operator '%s', got '%s'",
				op,
				rule.Operator(),
			)
		}
	}
}

func TestIconSetTypes(t *testing.T) {
	iconSets := []IconSetType{
		IconSet3Arrows,
		IconSet3ArrowsGray,
		IconSet3Flags,
		IconSet3TrafficLights1,
		IconSet3TrafficLights2,
		IconSet3Signs,
		IconSet3Symbols,
		IconSet3Symbols2,
		IconSet4Arrows,
		IconSet4ArrowsGray,
		IconSet4RedToBlack,
		IconSet4Rating,
		IconSet4TrafficLights,
		IconSet5Arrows,
		IconSet5ArrowsGray,
		IconSet5Rating,
		IconSet5Quarters,
	}

	for _, iconSetType := range iconSets {
		is := NewIconSet()
		is.SetIconSetType(iconSetType)
		if is.IconSetType() != iconSetType {
			t.Errorf(
				"expected iconSet '%s', got '%s'",
				iconSetType,
				is.IconSetType(),
			)
		}
	}
}

func TestCfvoTypes(t *testing.T) {
	cfvoTypes := []CfvoType{
		CfvoTypeNum,
		CfvoTypePercent,
		CfvoTypeMax,
		CfvoTypeMin,
		CfvoTypeFormula,
		CfvoTypePercentile,
		CfvoTypeAutoMin,
		CfvoTypeAutoMax,
	}

	for _, cfvoType := range cfvoTypes {
		cfvo := NewCfvo()
		cfvo.SetType(cfvoType)
		if cfvo.Type() != cfvoType {
			t.Errorf(
				"expected type '%s', got '%s'",
				cfvoType,
				cfvo.Type(),
			)
		}
	}
}

func TestTimePeriodTypes(t *testing.T) {
	periods := []TimePeriodType{
		TimePeriodToday,
		TimePeriodYesterday,
		TimePeriodTomorrow,
		TimePeriodLast7Days,
		TimePeriodThisMonth,
		TimePeriodLastMonth,
		TimePeriodNextMonth,
		TimePeriodThisWeek,
		TimePeriodLastWeek,
		TimePeriodNextWeek,
	}

	for _, period := range periods {
		rule := NewCfRule()
		rule.SetTimePeriod(period)
		if rule.TimePeriod() != period {
			t.Errorf(
				"expected timePeriod '%s', got '%s'",
				period,
				rule.TimePeriod(),
			)
		}
	}
}
