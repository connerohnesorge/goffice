package elements

// CfRuleType represents the type of a conditional formatting rule.
type CfRuleType string

const (
	// CfRuleTypeExpression indicates an expression-based rule.
	CfRuleTypeExpression CfRuleType = "expression"
	// CfRuleTypeCellIs indicates a cell value comparison rule.
	CfRuleTypeCellIs CfRuleType = "cellIs"
	// CfRuleTypeColorScale indicates a color scale rule.
	CfRuleTypeColorScale CfRuleType = "colorScale"
	// CfRuleTypeDataBar indicates a data bar rule.
	CfRuleTypeDataBar CfRuleType = "dataBar"
	// CfRuleTypeIconSet indicates an icon set rule.
	CfRuleTypeIconSet CfRuleType = "iconSet"
	// CfRuleTypeTop10 indicates a top/bottom 10 rule.
	CfRuleTypeTop10 CfRuleType = "top10"
	// CfRuleTypeUniqueValues indicates a unique values rule.
	CfRuleTypeUniqueValues CfRuleType = "uniqueValues"
	// CfRuleTypeDuplicateValues indicates a duplicate values rule.
	CfRuleTypeDuplicateValues CfRuleType = "duplicateValues"
	// CfRuleTypeContainsText indicates a "contains text" rule.
	CfRuleTypeContainsText CfRuleType = "containsText"
	// CfRuleTypeNotContainsText indicates a "does not contain text" rule.
	CfRuleTypeNotContainsText CfRuleType = "notContainsText"
	// CfRuleTypeBeginsWith indicates a "begins with" rule.
	CfRuleTypeBeginsWith CfRuleType = "beginsWith"
	// CfRuleTypeEndsWith indicates an "ends with" rule.
	CfRuleTypeEndsWith CfRuleType = "endsWith"
	// CfRuleTypeContainsBlanks indicates a "contains blanks" rule.
	CfRuleTypeContainsBlanks CfRuleType = "containsBlanks"
	// CfRuleTypeNotContainsBlanks indicates a "does not contain blanks" rule.
	CfRuleTypeNotContainsBlanks CfRuleType = "notContainsBlanks"
	// CfRuleTypeContainsErrors indicates a "contains errors" rule.
	CfRuleTypeContainsErrors CfRuleType = "containsErrors"
	// CfRuleTypeNotContainsErrors indicates a "does not contain errors" rule.
	CfRuleTypeNotContainsErrors CfRuleType = "notContainsErrors"
	// CfRuleTypeTimePeriod indicates a time period rule.
	CfRuleTypeTimePeriod CfRuleType = "timePeriod"
	// CfRuleTypeAboveAverage indicates an above/below average rule.
	CfRuleTypeAboveAverage CfRuleType = "aboveAverage"
)

// CfRuleOperator represents the operator for a conditional formatting rule.
type CfRuleOperator string

const (
	// CfRuleOperatorLessThan indicates less than.
	CfRuleOperatorLessThan CfRuleOperator = "lessThan"
	// CfRuleOperatorLessThanOrEqual indicates less than or equal.
	CfRuleOperatorLessThanOrEqual CfRuleOperator = "lessThanOrEqual"
	// CfRuleOperatorEqual indicates equal.
	CfRuleOperatorEqual CfRuleOperator = "equal"
	// CfRuleOperatorNotEqual indicates not equal.
	CfRuleOperatorNotEqual CfRuleOperator = "notEqual"
	// CfRuleOperatorGreaterThanOrEqual indicates greater than or equal.
	CfRuleOperatorGreaterThanOrEqual CfRuleOperator = "greaterThanOrEqual"
	// CfRuleOperatorGreaterThan indicates greater than.
	CfRuleOperatorGreaterThan CfRuleOperator = "greaterThan"
	// CfRuleOperatorBetween indicates between (inclusive).
	CfRuleOperatorBetween CfRuleOperator = "between"
	// CfRuleOperatorNotBetween indicates not between.
	CfRuleOperatorNotBetween CfRuleOperator = "notBetween"
	// CfRuleOperatorContainsText indicates contains text.
	CfRuleOperatorContainsText CfRuleOperator = "containsText"
	// CfRuleOperatorNotContains indicates does not contain.
	CfRuleOperatorNotContains CfRuleOperator = "notContains"
	// CfRuleOperatorBeginsWith indicates begins with.
	CfRuleOperatorBeginsWith CfRuleOperator = "beginsWith"
	// CfRuleOperatorEndsWith indicates ends with.
	CfRuleOperatorEndsWith CfRuleOperator = "endsWith"
)

// TimePeriodType represents the time period for time-based conditional
// formatting.
type TimePeriodType string

const (
	// TimePeriodToday indicates today.
	TimePeriodToday TimePeriodType = "today"
	// TimePeriodYesterday indicates yesterday.
	TimePeriodYesterday TimePeriodType = "yesterday"
	// TimePeriodTomorrow indicates tomorrow.
	TimePeriodTomorrow TimePeriodType = "tomorrow"
	// TimePeriodLast7Days indicates the last 7 days.
	TimePeriodLast7Days TimePeriodType = "last7Days"
	// TimePeriodThisMonth indicates this month.
	TimePeriodThisMonth TimePeriodType = "thisMonth"
	// TimePeriodLastMonth indicates last month.
	TimePeriodLastMonth TimePeriodType = "lastMonth"
	// TimePeriodNextMonth indicates next month.
	TimePeriodNextMonth TimePeriodType = "nextMonth"
	// TimePeriodThisWeek indicates this week.
	TimePeriodThisWeek TimePeriodType = "thisWeek"
	// TimePeriodLastWeek indicates last week.
	TimePeriodLastWeek TimePeriodType = "lastWeek"
	// TimePeriodNextWeek indicates next week.
	TimePeriodNextWeek TimePeriodType = "nextWeek"
)
