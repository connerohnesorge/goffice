package elements

import (
	"strings"
	"testing"
)

func TestNewAutoFilter(t *testing.T) {
	af := NewAutoFilter()

	if af == nil {
		t.Fatal("expected non-nil AutoFilter")
	}
	if af.LocalName() != "autoFilter" {
		t.Errorf(
			"expected localName 'autoFilter', got '%s'",
			af.LocalName(),
		)
	}
	if af.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceSML,
			af.NamespaceURI(),
		)
	}
}

func TestNewAutoFilterWithRef(t *testing.T) {
	af := NewAutoFilterWithRef("A1:D100")

	if af.Ref() != "A1:D100" {
		t.Errorf(
			"expected ref 'A1:D100', got '%s'",
			af.Ref(),
		)
	}
}

func TestAutoFilter_Ref(t *testing.T) {
	af := NewAutoFilter()

	if af.Ref() != "" {
		t.Errorf(
			"expected empty ref, got '%s'",
			af.Ref(),
		)
	}

	af.SetRef("B2:E50")
	if af.Ref() != "B2:E50" {
		t.Errorf(
			"expected ref 'B2:E50', got '%s'",
			af.Ref(),
		)
	}
}

func TestAutoFilter_AddFilterColumn(
	t *testing.T,
) {
	af := NewAutoFilter()
	af.SetRef("A1:C10")

	fc := af.AddFilterColumn(0)
	if fc == nil {
		t.Fatal("expected non-nil FilterColumn")
	}
	if fc.ColId() != 0 {
		t.Errorf(
			"expected colId 0, got %d",
			fc.ColId(),
		)
	}

	fc2 := af.AddFilterColumn(1)
	if fc2.ColId() != 1 {
		t.Errorf(
			"expected colId 1, got %d",
			fc2.ColId(),
		)
	}

	if af.FilterColumnCount() != 2 {
		t.Errorf(
			"expected 2 filter columns, got %d",
			af.FilterColumnCount(),
		)
	}
}

func TestAutoFilter_GetFilterColumnByColId(
	t *testing.T,
) {
	af := NewAutoFilter()
	af.SetRef("A1:C10")
	af.AddFilterColumn(0)
	af.AddFilterColumn(2)

	found := af.GetFilterColumnByColId(2)
	if found == nil {
		t.Fatal("expected to find filter column")
	}
	if found.ColId() != 2 {
		t.Errorf(
			"expected colId 2, got %d",
			found.ColId(),
		)
	}

	notFound := af.GetFilterColumnByColId(1)
	if notFound != nil {
		t.Error(
			"expected nil for non-existent colId",
		)
	}
}

func TestAutoFilter_RemoveFilterColumn(
	t *testing.T,
) {
	af := NewAutoFilter()
	af.AddFilterColumn(0)
	af.AddFilterColumn(1)

	removed := af.RemoveFilterColumn(0)
	if !removed {
		t.Error(
			"expected RemoveFilterColumn to return true",
		)
	}
	if af.FilterColumnCount() != 1 {
		t.Errorf(
			"expected 1 filter column, got %d",
			af.FilterColumnCount(),
		)
	}

	removed = af.RemoveFilterColumn(0)
	if removed {
		t.Error(
			"expected RemoveFilterColumn to return false",
		)
	}
}

func TestAutoFilter_ClearFilters(t *testing.T) {
	af := NewAutoFilter()
	af.AddFilterColumn(0)
	af.AddFilterColumn(1)
	af.AddFilterColumn(2)

	af.ClearFilters()
	if af.FilterColumnCount() != 0 {
		t.Errorf(
			"expected 0 filter columns, got %d",
			af.FilterColumnCount(),
		)
	}
}

func TestAutoFilter_SortState(t *testing.T) {
	af := NewAutoFilter()
	af.SetRef("A1:C10")

	if af.SortState() != nil {
		t.Error(
			"expected nil SortState initially",
		)
	}

	ss := af.GetOrCreateSortState()
	if ss == nil {
		t.Fatal("expected non-nil SortState")
	}

	ss.SetRef("A1:C10")
	ss.AddSortCondition("A1:A10")

	// Should return the same element
	ss2 := af.GetOrCreateSortState()
	if ss2 != ss {
		t.Error(
			"expected same SortState instance",
		)
	}

	removed := af.RemoveSortState()
	if !removed {
		t.Error(
			"expected RemoveSortState to return true",
		)
	}
	if af.SortState() != nil {
		t.Error(
			"expected nil SortState after removal",
		)
	}
}

func TestNewFilterColumn(t *testing.T) {
	fc := NewFilterColumn()

	if fc == nil {
		t.Fatal("expected non-nil FilterColumn")
	}
	if fc.LocalName() != "filterColumn" {
		t.Errorf(
			"expected localName 'filterColumn', got '%s'",
			fc.LocalName(),
		)
	}
}

func TestFilterColumn_ColId(t *testing.T) {
	fc := NewFilterColumn()

	fc.SetColId(5)
	if fc.ColId() != 5 {
		t.Errorf(
			"expected colId 5, got %d",
			fc.ColId(),
		)
	}
}

func TestFilterColumn_HiddenButton(t *testing.T) {
	fc := NewFilterColumn()

	if fc.HiddenButton() {
		t.Error(
			"expected hiddenButton false initially",
		)
	}
	fc.SetHiddenButton(true)
	if !fc.HiddenButton() {
		t.Error("expected hiddenButton true")
	}
}

func TestFilterColumn_ShowButton(t *testing.T) {
	fc := NewFilterColumn()

	// Default is true
	if !fc.ShowButton() {
		t.Error(
			"expected showButton true by default",
		)
	}
	fc.SetShowButton(false)
	if fc.ShowButton() {
		t.Error("expected showButton false")
	}
}

func TestFilterColumn_Filters(t *testing.T) {
	fc := NewFilterColumn()

	if fc.Filters() != nil {
		t.Error("expected nil Filters initially")
	}

	f := fc.GetOrCreateFilters()
	if f == nil {
		t.Fatal("expected non-nil Filters")
	}

	// Should return same instance
	f2 := fc.GetOrCreateFilters()
	if f2 != f {
		t.Error("expected same Filters instance")
	}
}

func TestFilterColumn_Top10(t *testing.T) {
	fc := NewFilterColumn()

	if fc.Top10() != nil {
		t.Error("expected nil Top10 initially")
	}

	top := fc.GetOrCreateTop10()
	if top == nil {
		t.Fatal("expected non-nil Top10")
	}

	// Creating Top10 should clear Filters
	fc.GetOrCreateFilters()
	if fc.Top10() != nil {
		t.Error(
			"expected Top10 to be cleared when creating Filters",
		)
	}
}

func TestFilterColumn_CustomFilters(
	t *testing.T,
) {
	fc := NewFilterColumn()

	if fc.CustomFilters() != nil {
		t.Error(
			"expected nil CustomFilters initially",
		)
	}

	cf := fc.GetOrCreateCustomFilters()
	if cf == nil {
		t.Fatal("expected non-nil CustomFilters")
	}
}

func TestFilterColumn_DynamicFilter(
	t *testing.T,
) {
	fc := NewFilterColumn()

	if fc.DynamicFilter() != nil {
		t.Error(
			"expected nil DynamicFilter initially",
		)
	}

	df := fc.GetOrCreateDynamicFilter()
	if df == nil {
		t.Fatal("expected non-nil DynamicFilter")
	}
}

func TestFilterColumn_ColorFilter(t *testing.T) {
	fc := NewFilterColumn()

	if fc.ColorFilter() != nil {
		t.Error(
			"expected nil ColorFilter initially",
		)
	}

	cf := fc.GetOrCreateColorFilter()
	if cf == nil {
		t.Fatal("expected non-nil ColorFilter")
	}
}

func TestFilterColumn_IconFilter(t *testing.T) {
	fc := NewFilterColumn()

	if fc.IconFilter() != nil {
		t.Error(
			"expected nil IconFilter initially",
		)
	}

	i := fc.GetOrCreateIconFilter()
	if i == nil {
		t.Fatal("expected non-nil IconFilter")
	}
}

func TestFilterColumn_ClearFilter(t *testing.T) {
	fc := NewFilterColumn()

	fc.GetOrCreateFilters().AddFilter("Value1")
	fc.ClearFilter()

	if fc.Filters() != nil {
		t.Error(
			"expected nil Filters after ClearFilter",
		)
	}
}

func TestNewFilters(t *testing.T) {
	f := NewFilters()

	if f == nil {
		t.Fatal("expected non-nil Filters")
	}
	if f.LocalName() != "filters" {
		t.Errorf(
			"expected localName 'filters', got '%s'",
			f.LocalName(),
		)
	}
}

func TestFilters_Blank(t *testing.T) {
	f := NewFilters()

	if f.Blank() {
		t.Error("expected blank false initially")
	}
	f.SetBlank(true)
	if !f.Blank() {
		t.Error("expected blank true")
	}
}

func TestFilters_CalendarType(t *testing.T) {
	f := NewFilters()

	if f.CalendarType() != CalendarTypeNone {
		t.Errorf(
			"expected calendarType 'none', got '%s'",
			f.CalendarType(),
		)
	}
	f.SetCalendarType(CalendarTypeGregorian)
	if f.CalendarType() != CalendarTypeGregorian {
		t.Errorf(
			"expected calendarType 'gregorian', got '%s'",
			f.CalendarType(),
		)
	}
}

func TestFilters_AddFilter(t *testing.T) {
	f := NewFilters()

	filter := f.AddFilter("Value1")
	if filter == nil {
		t.Fatal("expected non-nil Filter")
	}
	if filter.Val() != "Value1" {
		t.Errorf(
			"expected val 'Value1', got '%s'",
			filter.Val(),
		)
	}

	f.AddFilter("Value2")
	if f.FilterCount() != 2 {
		t.Errorf(
			"expected 2 filters, got %d",
			f.FilterCount(),
		)
	}
}

func TestFilters_RemoveFilter(t *testing.T) {
	f := NewFilters()
	f.AddFilter("Value1")
	f.AddFilter("Value2")

	removed := f.RemoveFilter("Value1")
	if !removed {
		t.Error(
			"expected RemoveFilter to return true",
		)
	}
	if f.FilterCount() != 1 {
		t.Errorf(
			"expected 1 filter, got %d",
			f.FilterCount(),
		)
	}
}

func TestFilters_HasFilter(t *testing.T) {
	f := NewFilters()
	f.AddFilter("Value1")

	if !f.HasFilter("Value1") {
		t.Error(
			"expected HasFilter to return true",
		)
	}
	if f.HasFilter("Value2") {
		t.Error(
			"expected HasFilter to return false",
		)
	}
}

func TestFilters_AddDateGroupItem(t *testing.T) {
	f := NewFilters()

	dgi := f.AddDateGroupItem(
		DateTimeGroupingYear,
		2024,
	)
	if dgi == nil {
		t.Fatal("expected non-nil DateGroupItem")
	}
	if dgi.Year() != 2024 {
		t.Errorf(
			"expected year 2024, got %d",
			dgi.Year(),
		)
	}
	if dgi.DateTimeGrouping() != DateTimeGroupingYear {
		t.Errorf(
			"expected dateTimeGrouping 'year', got '%s'",
			dgi.DateTimeGrouping(),
		)
	}
}

func TestFilters_ClearAll(t *testing.T) {
	f := NewFilters()
	f.AddFilter("Value1")
	f.AddFilter("Value2")
	f.AddDateGroupItem(
		DateTimeGroupingMonth,
		2024,
	)

	f.ClearAll()
	if f.FilterCount() != 0 {
		t.Errorf(
			"expected 0 filters, got %d",
			f.FilterCount(),
		)
	}
	if f.DateGroupItemCount() != 0 {
		t.Errorf(
			"expected 0 dateGroupItems, got %d",
			f.DateGroupItemCount(),
		)
	}
}

func TestNewFilter(t *testing.T) {
	f := NewFilter()

	if f == nil {
		t.Fatal("expected non-nil Filter")
	}
	if f.LocalName() != "filter" {
		t.Errorf(
			"expected localName 'filter', got '%s'",
			f.LocalName(),
		)
	}
}

func TestFilter_Val(t *testing.T) {
	f := NewFilter()

	f.SetVal("TestValue")
	if f.Val() != "TestValue" {
		t.Errorf(
			"expected val 'TestValue', got '%s'",
			f.Val(),
		)
	}
}

func TestNewDateGroupItem(t *testing.T) {
	dgi := NewDateGroupItem()

	if dgi == nil {
		t.Fatal("expected non-nil DateGroupItem")
	}
	if dgi.LocalName() != "dateGroupItem" {
		t.Errorf(
			"expected localName 'dateGroupItem', got '%s'",
			dgi.LocalName(),
		)
	}
}

func TestDateGroupItem_Attributes(t *testing.T) {
	dgi := NewDateGroupItem()

	dgi.SetYear(2024)
	dgi.SetMonth(12)
	dgi.SetDay(25)
	dgi.SetHour(10)
	dgi.SetMinute(30)
	dgi.SetSecond(45)
	dgi.SetDateTimeGrouping(
		DateTimeGroupingSecond,
	)

	if dgi.Year() != 2024 {
		t.Errorf(
			"expected year 2024, got %d",
			dgi.Year(),
		)
	}
	if dgi.Month() != 12 {
		t.Errorf(
			"expected month 12, got %d",
			dgi.Month(),
		)
	}
	if dgi.Day() != 25 {
		t.Errorf(
			"expected day 25, got %d",
			dgi.Day(),
		)
	}
	if dgi.Hour() != 10 {
		t.Errorf(
			"expected hour 10, got %d",
			dgi.Hour(),
		)
	}
	if dgi.Minute() != 30 {
		t.Errorf(
			"expected minute 30, got %d",
			dgi.Minute(),
		)
	}
	if dgi.Second() != 45 {
		t.Errorf(
			"expected second 45, got %d",
			dgi.Second(),
		)
	}
	if dgi.DateTimeGrouping() != DateTimeGroupingSecond {
		t.Errorf(
			"expected dateTimeGrouping 'second', got '%s'",
			dgi.DateTimeGrouping(),
		)
	}
}

func TestNewCustomFilters(t *testing.T) {
	cf := NewCustomFilters()

	if cf == nil {
		t.Fatal("expected non-nil CustomFilters")
	}
	if cf.LocalName() != "customFilters" {
		t.Errorf(
			"expected localName 'customFilters', got '%s'",
			cf.LocalName(),
		)
	}
}

func TestCustomFilters_And(t *testing.T) {
	cf := NewCustomFilters()

	// Default is false (OR)
	if cf.And() {
		t.Error("expected and false by default")
	}
	cf.SetAnd(true)
	if !cf.And() {
		t.Error("expected and true")
	}
}

func TestCustomFilters_AddCustomFilter(
	t *testing.T,
) {
	cf := NewCustomFilters()

	filter := cf.AddCustomFilter(
		"100",
		FilterOperatorGreaterThan,
	)
	if filter == nil {
		t.Fatal("expected non-nil CustomFilter")
	}
	if filter.Val() != "100" {
		t.Errorf(
			"expected val '100', got '%s'",
			filter.Val(),
		)
	}
	if filter.Operator() != FilterOperatorGreaterThan {
		t.Errorf(
			"expected operator 'greaterThan', got '%s'",
			filter.Operator(),
		)
	}
}

func TestCustomFilters_SetSingleFilter(
	t *testing.T,
) {
	cf := NewCustomFilters()
	cf.AddCustomFilter(
		"50",
		FilterOperatorLessThan,
	)
	cf.AddCustomFilter(
		"100",
		FilterOperatorGreaterThan,
	)

	cf.SetSingleFilter("75", FilterOperatorEqual)
	if cf.CustomFilterCount() != 1 {
		t.Errorf(
			"expected 1 filter, got %d",
			cf.CustomFilterCount(),
		)
	}
}

func TestCustomFilters_SetDualFilters(
	t *testing.T,
) {
	cf := NewCustomFilters()

	cf.SetDualFilters(
		"10", FilterOperatorGreaterThanOrEqual,
		"100", FilterOperatorLessThanOrEqual,
		true,
	)

	if cf.CustomFilterCount() != 2 {
		t.Errorf(
			"expected 2 filters, got %d",
			cf.CustomFilterCount(),
		)
	}
	if !cf.And() {
		t.Error("expected and to be true")
	}
}

func TestNewCustomFilter(t *testing.T) {
	c := NewCustomFilter()

	if c == nil {
		t.Fatal("expected non-nil CustomFilter")
	}
	if c.LocalName() != "customFilter" {
		t.Errorf(
			"expected localName 'customFilter', got '%s'",
			c.LocalName(),
		)
	}
}

func TestCustomFilter_Operator(t *testing.T) {
	c := NewCustomFilter()

	// Default is equal
	if c.Operator() != FilterOperatorEqual {
		t.Errorf(
			"expected default operator 'equal', got '%s'",
			c.Operator(),
		)
	}

	c.SetOperator(FilterOperatorNotEqual)
	if c.Operator() != FilterOperatorNotEqual {
		t.Errorf(
			"expected operator 'notEqual', got '%s'",
			c.Operator(),
		)
	}
}

func TestNewDynamicFilter(t *testing.T) {
	df := NewDynamicFilter()

	if df == nil {
		t.Fatal("expected non-nil DynamicFilter")
	}
	if df.LocalName() != "dynamicFilter" {
		t.Errorf(
			"expected localName 'dynamicFilter', got '%s'",
			df.LocalName(),
		)
	}
}

func TestNewDynamicFilterWithType(t *testing.T) {
	df := NewDynamicFilterWithType(
		DynamicFilterToday,
	)

	if df.Type() != DynamicFilterToday {
		t.Errorf(
			"expected type 'today', got '%s'",
			df.Type(),
		)
	}
}

func TestDynamicFilter_Type(t *testing.T) {
	df := NewDynamicFilter()

	df.SetType(DynamicFilterAboveAverage)
	if df.Type() != DynamicFilterAboveAverage {
		t.Errorf(
			"expected type 'aboveAverage', got '%s'",
			df.Type(),
		)
	}
}

func TestDynamicFilter_Val(t *testing.T) {
	df := NewDynamicFilter()

	df.SetVal(50.5)
	if df.Val() != 50.5 {
		t.Errorf(
			"expected val 50.5, got %f",
			df.Val(),
		)
	}
}

func TestDynamicFilter_MaxVal(t *testing.T) {
	df := NewDynamicFilter()

	df.SetMaxVal(100.5)
	if df.MaxVal() != 100.5 {
		t.Errorf(
			"expected maxVal 100.5, got %f",
			df.MaxVal(),
		)
	}
}

func TestNewColorFilter(t *testing.T) {
	cf := NewColorFilter()

	if cf == nil {
		t.Fatal("expected non-nil ColorFilter")
	}
	if cf.LocalName() != "colorFilter" {
		t.Errorf(
			"expected localName 'colorFilter', got '%s'",
			cf.LocalName(),
		)
	}
}

func TestColorFilter_DxfId(t *testing.T) {
	cf := NewColorFilter()

	cf.SetDxfId(5)
	if cf.DxfId() != 5 {
		t.Errorf(
			"expected dxfId 5, got %d",
			cf.DxfId(),
		)
	}
}

func TestColorFilter_CellColor(t *testing.T) {
	cf := NewColorFilter()

	// Default is true (cell background)
	if !cf.CellColor() {
		t.Error(
			"expected cellColor true by default",
		)
	}
	cf.SetCellColor(false)
	if cf.CellColor() {
		t.Error("expected cellColor false")
	}
}

func TestNewIconFilter(t *testing.T) {
	i := NewIconFilter()

	if i == nil {
		t.Fatal("expected non-nil IconFilter")
	}
	if i.LocalName() != "iconFilter" {
		t.Errorf(
			"expected localName 'iconFilter', got '%s'",
			i.LocalName(),
		)
	}
}

func TestIconFilter_IconSet(t *testing.T) {
	i := NewIconFilter()

	i.SetIconSet(IconSet3Arrows)
	if i.IconSet() != IconSet3Arrows {
		t.Errorf(
			"expected iconSet '3Arrows', got '%s'",
			i.IconSet(),
		)
	}
}

func TestIconFilter_IconId(t *testing.T) {
	i := NewIconFilter()

	i.SetIconId(2)
	if i.IconId() != 2 {
		t.Errorf(
			"expected iconId 2, got %d",
			i.IconId(),
		)
	}
}

func TestNewTop10(t *testing.T) {
	top := NewTop10()

	if top == nil {
		t.Fatal("expected non-nil Top10")
	}
	if top.LocalName() != "top10" {
		t.Errorf(
			"expected localName 'top10', got '%s'",
			top.LocalName(),
		)
	}
}

func TestNewTop10WithParams(t *testing.T) {
	top := NewTop10WithParams(10, true, false)

	if top.Val() != 10 {
		t.Errorf(
			"expected val 10, got %f",
			top.Val(),
		)
	}
	if !top.Top() {
		t.Error("expected top true")
	}
	if top.Percent() {
		t.Error("expected percent false")
	}
}

func TestTop10_Top(t *testing.T) {
	top := NewTop10()

	// Default is true
	if !top.Top() {
		t.Error("expected top true by default")
	}
	top.SetTop(false)
	if top.Top() {
		t.Error("expected top false")
	}
}

func TestTop10_Percent(t *testing.T) {
	top := NewTop10()

	// Default is false
	if top.Percent() {
		t.Error(
			"expected percent false by default",
		)
	}
	top.SetPercent(true)
	if !top.Percent() {
		t.Error("expected percent true")
	}
}

func TestTop10_Val(t *testing.T) {
	top := NewTop10()

	top.SetVal(25)
	if top.Val() != 25 {
		t.Errorf(
			"expected val 25, got %f",
			top.Val(),
		)
	}
}

func TestTop10_FilterVal(t *testing.T) {
	top := NewTop10()

	top.SetFilterVal(150.5)
	if top.FilterVal() != 150.5 {
		t.Errorf(
			"expected filterVal 150.5, got %f",
			top.FilterVal(),
		)
	}
}

func TestNewSortState(t *testing.T) {
	ss := NewSortState()

	if ss == nil {
		t.Fatal("expected non-nil SortState")
	}
	if ss.LocalName() != "sortState" {
		t.Errorf(
			"expected localName 'sortState', got '%s'",
			ss.LocalName(),
		)
	}
}

func TestSortState_Ref(t *testing.T) {
	ss := NewSortState()

	ss.SetRef("A1:D100")
	if ss.Ref() != "A1:D100" {
		t.Errorf(
			"expected ref 'A1:D100', got '%s'",
			ss.Ref(),
		)
	}
}

func TestSortState_CaseSensitive(t *testing.T) {
	ss := NewSortState()

	if ss.CaseSensitive() {
		t.Error(
			"expected caseSensitive false initially",
		)
	}
	ss.SetCaseSensitive(true)
	if !ss.CaseSensitive() {
		t.Error("expected caseSensitive true")
	}
}

func TestSortState_ColumnSort(t *testing.T) {
	ss := NewSortState()

	if ss.ColumnSort() {
		t.Error(
			"expected columnSort false initially",
		)
	}
	ss.SetColumnSort(true)
	if !ss.ColumnSort() {
		t.Error("expected columnSort true")
	}
}

func TestSortState_SortMethod(t *testing.T) {
	ss := NewSortState()

	if ss.SortMethod() != SortMethodNone {
		t.Errorf(
			"expected sortMethod 'none', got '%s'",
			ss.SortMethod(),
		)
	}
	ss.SetSortMethod(SortMethodPinYin)
	if ss.SortMethod() != SortMethodPinYin {
		t.Errorf(
			"expected sortMethod 'pinYin', got '%s'",
			ss.SortMethod(),
		)
	}
}

func TestSortState_AddSortCondition(
	t *testing.T,
) {
	ss := NewSortState()

	sc := ss.AddSortCondition("A1:A100")
	if sc == nil {
		t.Fatal("expected non-nil SortCondition")
	}
	if sc.Ref() != "A1:A100" {
		t.Errorf(
			"expected ref 'A1:A100', got '%s'",
			sc.Ref(),
		)
	}
	if ss.SortConditionCount() != 1 {
		t.Errorf(
			"expected 1 sort condition, got %d",
			ss.SortConditionCount(),
		)
	}
}

func TestSortState_AddDescendingSortCondition(
	t *testing.T,
) {
	ss := NewSortState()

	sc := ss.AddDescendingSortCondition("B1:B100")
	if !sc.Descending() {
		t.Error("expected descending true")
	}
}

func TestSortState_ClearSortConditions(
	t *testing.T,
) {
	ss := NewSortState()
	ss.AddSortCondition("A1:A100")
	ss.AddSortCondition("B1:B100")

	ss.ClearSortConditions()
	if ss.SortConditionCount() != 0 {
		t.Errorf(
			"expected 0 sort conditions, got %d",
			ss.SortConditionCount(),
		)
	}
}

func TestNewSortCondition(t *testing.T) {
	sc := NewSortCondition()

	if sc == nil {
		t.Fatal("expected non-nil SortCondition")
	}
	if sc.LocalName() != "sortCondition" {
		t.Errorf(
			"expected localName 'sortCondition', got '%s'",
			sc.LocalName(),
		)
	}
}

func TestSortCondition_Descending(t *testing.T) {
	sc := NewSortCondition()

	if sc.Descending() {
		t.Error(
			"expected descending false initially",
		)
	}
	sc.SetDescending(true)
	if !sc.Descending() {
		t.Error("expected descending true")
	}
}

func TestSortCondition_SortBy(t *testing.T) {
	sc := NewSortCondition()

	// Default is value
	if sc.SortBy() != SortByValue {
		t.Errorf(
			"expected sortBy 'value', got '%s'",
			sc.SortBy(),
		)
	}
	sc.SetSortBy(SortByCellColor)
	if sc.SortBy() != SortByCellColor {
		t.Errorf(
			"expected sortBy 'cellColor', got '%s'",
			sc.SortBy(),
		)
	}
}

func TestSortCondition_DxfId(t *testing.T) {
	sc := NewSortCondition()

	sc.SetDxfId(3)
	if sc.DxfId() != 3 {
		t.Errorf(
			"expected dxfId 3, got %d",
			sc.DxfId(),
		)
	}
}

func TestSortCondition_CustomList(t *testing.T) {
	sc := NewSortCondition()

	sc.SetCustomList("Mon,Tue,Wed,Thu,Fri")
	if sc.CustomList() != "Mon,Tue,Wed,Thu,Fri" {
		t.Errorf(
			"expected customList 'Mon,Tue,Wed,Thu,Fri', got '%s'",
			sc.CustomList(),
		)
	}
}

func TestSortCondition_IconSet(t *testing.T) {
	sc := NewSortCondition()

	sc.SetIconSet(IconSet3Flags)
	if sc.IconSet() != IconSet3Flags {
		t.Errorf(
			"expected iconSet '3Flags', got '%s'",
			sc.IconSet(),
		)
	}
}

func TestSortCondition_IconId(t *testing.T) {
	sc := NewSortCondition()

	sc.SetIconId(1)
	if sc.IconId() != 1 {
		t.Errorf(
			"expected iconId 1, got %d",
			sc.IconId(),
		)
	}
}

func TestAutoFilter_XMLSerialization(
	t *testing.T,
) {
	af := NewAutoFilter()
	af.SetRef("A1:D100")

	fc := af.AddFilterColumn(0)
	f := fc.GetOrCreateFilters()
	f.AddFilter("Value1")
	f.AddFilter("Value2")

	xml := af.OuterXml()

	if !strings.Contains(xml, "autoFilter") {
		t.Error("expected 'autoFilter' in XML")
	}
	if !strings.Contains(xml, "ref=\"A1:D100\"") {
		t.Error(
			"expected 'ref=\"A1:D100\"' in XML",
		)
	}
	if !strings.Contains(xml, "filterColumn") {
		t.Error("expected 'filterColumn' in XML")
	}
	if !strings.Contains(xml, "filters") {
		t.Error("expected 'filters' in XML")
	}
	if !strings.Contains(xml, "filter") {
		t.Error("expected 'filter' in XML")
	}
}

func TestDynamicFilterTypes(t *testing.T) {
	types := []DynamicFilterType{
		DynamicFilterNull,
		DynamicFilterAboveAverage,
		DynamicFilterBelowAverage,
		DynamicFilterTomorrow,
		DynamicFilterToday,
		DynamicFilterYesterday,
		DynamicFilterNextWeek,
		DynamicFilterThisWeek,
		DynamicFilterLastWeek,
		DynamicFilterNextMonth,
		DynamicFilterThisMonth,
		DynamicFilterLastMonth,
		DynamicFilterNextQuarter,
		DynamicFilterThisQuarter,
		DynamicFilterLastQuarter,
		DynamicFilterNextYear,
		DynamicFilterThisYear,
		DynamicFilterLastYear,
		DynamicFilterYearToDate,
		DynamicFilterQ1,
		DynamicFilterQ2,
		DynamicFilterQ3,
		DynamicFilterQ4,
		DynamicFilterM1,
		DynamicFilterM2,
		DynamicFilterM3,
		DynamicFilterM4,
		DynamicFilterM5,
		DynamicFilterM6,
		DynamicFilterM7,
		DynamicFilterM8,
		DynamicFilterM9,
		DynamicFilterM10,
		DynamicFilterM11,
		DynamicFilterM12,
	}

	for _, filterType := range types {
		df := NewDynamicFilter()
		df.SetType(filterType)
		if df.Type() != filterType {
			t.Errorf(
				"expected type '%s', got '%s'",
				filterType,
				df.Type(),
			)
		}
	}
}

func TestFilterOperators(t *testing.T) {
	operators := []FilterOperator{
		FilterOperatorLessThan,
		FilterOperatorLessThanOrEqual,
		FilterOperatorEqual,
		FilterOperatorNotEqual,
		FilterOperatorGreaterThanOrEqual,
		FilterOperatorGreaterThan,
	}

	for _, op := range operators {
		c := NewCustomFilter()
		//nolint:revive // early-return: test structure pattern
		if op != FilterOperatorEqual {
			c.SetOperator(op)
			if c.Operator() != op {
				t.Errorf(
					"expected operator '%s', got '%s'",
					op,
					c.Operator(),
				)
			}
		}
	}
}

func TestCalendarTypes(t *testing.T) {
	types := []CalendarType{
		CalendarTypeNone,
		CalendarTypeGregorian,
		CalendarTypeGregorianUs,
		CalendarTypeJapan,
		CalendarTypeTaiwan,
		CalendarTypeKorea,
		CalendarTypeHijri,
		CalendarTypeThai,
		CalendarTypeHebrew,
		CalendarTypeGregorianMeFrench,
		CalendarTypeGregorianArabic,
		CalendarTypeGregorianXlitEnglish,
		CalendarTypeGregorianXlitFrench,
	}

	for _, ct := range types {
		f := NewFilters()
		//nolint:revive // early-return: test structure pattern
		if ct != CalendarTypeNone {
			f.SetCalendarType(ct)
			if f.CalendarType() != ct {
				t.Errorf(
					"expected calendarType '%s', got '%s'",
					ct,
					f.CalendarType(),
				)
			}
		}
	}
}

func TestDateTimeGroupings(t *testing.T) {
	groupings := []DateTimeGrouping{
		DateTimeGroupingYear,
		DateTimeGroupingMonth,
		DateTimeGroupingDay,
		DateTimeGroupingHour,
		DateTimeGroupingMinute,
		DateTimeGroupingSecond,
	}

	for _, g := range groupings {
		dgi := NewDateGroupItem()
		dgi.SetDateTimeGrouping(g)
		if dgi.DateTimeGrouping() != g {
			t.Errorf(
				"expected dateTimeGrouping '%s', got '%s'",
				g,
				dgi.DateTimeGrouping(),
			)
		}
	}
}

func TestSortByValues(t *testing.T) {
	values := []SortBy{
		SortByValue,
		SortByCellColor,
		SortByFontColor,
		SortByIcon,
	}

	for _, v := range values {
		sc := NewSortCondition()
		//nolint:revive // early-return: test structure pattern
		if v != SortByValue {
			sc.SetSortBy(v)
			if sc.SortBy() != v {
				t.Errorf(
					"expected sortBy '%s', got '%s'",
					v,
					sc.SortBy(),
				)
			}
		}
	}
}

func TestSortMethods(t *testing.T) {
	methods := []SortMethod{
		SortMethodNone,
		SortMethodStroke,
		SortMethodPinYin,
	}

	for _, m := range methods {
		ss := NewSortState()
		//nolint:revive // early-return: test structure pattern
		if m != SortMethodNone {
			ss.SetSortMethod(m)
			if ss.SortMethod() != m {
				t.Errorf(
					"expected sortMethod '%s', got '%s'",
					m,
					ss.SortMethod(),
				)
			}
		}
	}
}
