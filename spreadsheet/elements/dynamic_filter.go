package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// DynamicFilterType represents the type of dynamic filter.
type DynamicFilterType string

const (
	// DynamicFilterNull indicates null filter.
	DynamicFilterNull DynamicFilterType = "null"
	// DynamicFilterAboveAverage indicates above average filter.
	DynamicFilterAboveAverage DynamicFilterType = "aboveAverage"
	// DynamicFilterBelowAverage indicates below average filter.
	DynamicFilterBelowAverage DynamicFilterType = "belowAverage"
	// DynamicFilterTomorrow indicates tomorrow filter.
	DynamicFilterTomorrow DynamicFilterType = "tomorrow"
	// DynamicFilterToday indicates today filter.
	DynamicFilterToday DynamicFilterType = "today"
	// DynamicFilterYesterday indicates yesterday filter.
	DynamicFilterYesterday DynamicFilterType = "yesterday"
	// DynamicFilterNextWeek indicates next week filter.
	DynamicFilterNextWeek DynamicFilterType = "nextWeek"
	// DynamicFilterThisWeek indicates this week filter.
	DynamicFilterThisWeek DynamicFilterType = "thisWeek"
	// DynamicFilterLastWeek indicates last week filter.
	DynamicFilterLastWeek DynamicFilterType = "lastWeek"
	// DynamicFilterNextMonth indicates next month filter.
	DynamicFilterNextMonth DynamicFilterType = "nextMonth"
	// DynamicFilterThisMonth indicates this month filter.
	DynamicFilterThisMonth DynamicFilterType = "thisMonth"
	// DynamicFilterLastMonth indicates last month filter.
	DynamicFilterLastMonth DynamicFilterType = "lastMonth"
	// DynamicFilterNextQuarter indicates next quarter filter.
	DynamicFilterNextQuarter DynamicFilterType = "nextQuarter"
	// DynamicFilterThisQuarter indicates this quarter filter.
	DynamicFilterThisQuarter DynamicFilterType = "thisQuarter"
	// DynamicFilterLastQuarter indicates last quarter filter.
	DynamicFilterLastQuarter DynamicFilterType = "lastQuarter"
	// DynamicFilterNextYear indicates next year filter.
	DynamicFilterNextYear DynamicFilterType = "nextYear"
	// DynamicFilterThisYear indicates this year filter.
	DynamicFilterThisYear DynamicFilterType = "thisYear"
	// DynamicFilterLastYear indicates last year filter.
	DynamicFilterLastYear DynamicFilterType = "lastYear"
	// DynamicFilterYearToDate indicates year to date filter.
	DynamicFilterYearToDate DynamicFilterType = "yearToDate"
	// DynamicFilterQ1 indicates Q1 filter.
	DynamicFilterQ1 DynamicFilterType = "Q1"
	// DynamicFilterQ2 indicates Q2 filter.
	DynamicFilterQ2 DynamicFilterType = "Q2"
	// DynamicFilterQ3 indicates Q3 filter.
	DynamicFilterQ3 DynamicFilterType = "Q3"
	// DynamicFilterQ4 indicates Q4 filter.
	DynamicFilterQ4 DynamicFilterType = "Q4"
	// DynamicFilterM1 indicates January filter.
	DynamicFilterM1 DynamicFilterType = "M1"
	// DynamicFilterM2 indicates February filter.
	DynamicFilterM2 DynamicFilterType = "M2"
	// DynamicFilterM3 indicates March filter.
	DynamicFilterM3 DynamicFilterType = "M3"
	// DynamicFilterM4 indicates April filter.
	DynamicFilterM4 DynamicFilterType = "M4"
	// DynamicFilterM5 indicates May filter.
	DynamicFilterM5 DynamicFilterType = "M5"
	// DynamicFilterM6 indicates June filter.
	DynamicFilterM6 DynamicFilterType = "M6"
	// DynamicFilterM7 indicates July filter.
	DynamicFilterM7 DynamicFilterType = "M7"
	// DynamicFilterM8 indicates August filter.
	DynamicFilterM8 DynamicFilterType = "M8"
	// DynamicFilterM9 indicates September filter.
	DynamicFilterM9 DynamicFilterType = "M9"
	// DynamicFilterM10 indicates October filter.
	DynamicFilterM10 DynamicFilterType = "M10"
	// DynamicFilterM11 indicates November filter.
	DynamicFilterM11 DynamicFilterType = "M11"
	// DynamicFilterM12 indicates December filter.
	DynamicFilterM12 DynamicFilterType = "M12"
)

// DynamicFilter represents the dynamicFilter element (x:dynamicFilter).
// It defines a dynamic, date-based filter criteria.
type DynamicFilter struct {
	*openxml.LeafElementBase
}

// NewDynamicFilter creates a new DynamicFilter element.
func NewDynamicFilter() *DynamicFilter {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"dynamicFilter",
		PrefixDefault,
	)

	return &DynamicFilter{LeafElementBase: elem}
}

// NewDynamicFilterWithType creates a new DynamicFilter element with the given
// type.
func NewDynamicFilterWithType(
	filterType DynamicFilterType,
) *DynamicFilter {
	df := NewDynamicFilter()
	df.SetType(filterType)

	return df
}

// Type returns the dynamic filter type. Attribute: type.
func (df *DynamicFilter) Type() DynamicFilterType {
	attr, found := df.GetAttribute("type", "")
	if !found {
		return ""
	}

	return DynamicFilterType(attr.Value())
}

// SetType sets the dynamic filter type. Attribute: type.
func (df *DynamicFilter) SetType(
	filterType DynamicFilterType,
) {
	if filterType == "" {
		df.RemoveAttribute("type", "")

		return
	}
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			string(filterType),
		),
	)
}

// Val returns the value for the dynamic filter. Attribute: val.
// This is used as the boundary for average-based filters.
func (df *DynamicFilter) Val() float64 {
	attr, found := df.GetAttribute("val", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetVal sets the value for the dynamic filter. Attribute: val.
func (df *DynamicFilter) SetVal(value float64) {
	if value == 0 {
		df.RemoveAttribute("val", "")

		return
	}
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			strconv.FormatFloat(
				value,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// MaxVal returns the maximum value for the dynamic filter. Attribute: maxVal.
// This is used as the upper boundary for date range filters.
func (df *DynamicFilter) MaxVal() float64 {
	attr, found := df.GetAttribute("maxVal", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetMaxVal sets the maximum value for the dynamic filter. Attribute: maxVal.
func (df *DynamicFilter) SetMaxVal(
	value float64,
) {
	if value == 0 {
		df.RemoveAttribute("maxVal", "")

		return
	}
	df.SetAttribute(
		openxml.NewAttribute(
			"",
			"maxVal",
			"",
			strconv.FormatFloat(
				value,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// Clone creates a deep copy of this DynamicFilter element.
func (df *DynamicFilter) Clone() openxml.Element {
	cloned := df.LeafElementBase.Clone()

	return &DynamicFilter{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this DynamicFilter element.
func (df *DynamicFilter) CloneNode(
	deep bool,
) openxml.Element {
	cloned := df.LeafElementBase.CloneNode(deep)

	return &DynamicFilter{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
