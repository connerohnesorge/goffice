package elements

//revive:disable:file-length-limit many filter types

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// CalendarType represents the calendar system for date grouping.
type CalendarType string

const (
	// CalendarTypeNone indicates no calendar type.
	CalendarTypeNone CalendarType = "none"
	// CalendarTypeGregorian indicates Gregorian calendar.
	CalendarTypeGregorian CalendarType = "gregorian"
	// CalendarTypeGregorianUs indicates US Gregorian calendar.
	CalendarTypeGregorianUs CalendarType = "gregorianUs"
	// CalendarTypeJapan indicates Japanese calendar.
	CalendarTypeJapan CalendarType = "japan"
	// CalendarTypeTaiwan indicates Taiwan calendar.
	CalendarTypeTaiwan CalendarType = "taiwan"
	// CalendarTypeKorea indicates Korean calendar.
	CalendarTypeKorea CalendarType = "korea"
	// CalendarTypeHijri indicates Hijri calendar.
	CalendarTypeHijri CalendarType = "hijri"
	// CalendarTypeThai indicates Thai calendar.
	CalendarTypeThai CalendarType = "thai"
	// CalendarTypeHebrew indicates Hebrew calendar.
	CalendarTypeHebrew CalendarType = "hebrew"
	// CalendarTypeGregorianMeFrench indicates Middle East French calendar.
	CalendarTypeGregorianMeFrench CalendarType = "gregorianMeFrench"
	// CalendarTypeGregorianArabic indicates Arabic calendar.
	CalendarTypeGregorianArabic CalendarType = "gregorianArabic"
	// CalendarTypeGregorianXlitEnglish indicates English transliterated
	// calendar.
	CalendarTypeGregorianXlitEnglish CalendarType = "gregorianXlitEnglish"
	// CalendarTypeGregorianXlitFrench indicates French transliterated calendar.
	CalendarTypeGregorianXlitFrench CalendarType = "gregorianXlitFrench"
)

// DateTimeGrouping represents the date/time grouping granularity.
type DateTimeGrouping string

const (
	// DateTimeGroupingYear groups by year.
	DateTimeGroupingYear DateTimeGrouping = "year"
	// DateTimeGroupingMonth groups by month.
	DateTimeGroupingMonth DateTimeGrouping = "month"
	// DateTimeGroupingDay groups by day.
	DateTimeGroupingDay DateTimeGrouping = "day"
	// DateTimeGroupingHour groups by hour.
	DateTimeGroupingHour DateTimeGrouping = "hour"
	// DateTimeGroupingMinute groups by minute.
	DateTimeGroupingMinute DateTimeGrouping = "minute"
	// DateTimeGroupingSecond groups by second.
	DateTimeGroupingSecond DateTimeGrouping = "second"
)

// Filters represents the filters container element (x:filters).
// It contains filter values or date group items for a column.
type Filters struct {
	*openxml.CompositeElementBase
}

// NewFilters creates a new Filters element.
func NewFilters() *Filters {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"filters",
		PrefixDefault,
	)

	return &Filters{CompositeElementBase: elem}
}

// Blank returns whether to filter blank values. Attribute: blank.
func (f *Filters) Blank() bool {
	attr, found := f.GetAttribute("blank", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetBlank sets whether to filter blank values. Attribute: blank.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (f *Filters) SetBlank(value bool) {
	if value {
		f.SetAttribute(
			openxml.NewAttribute(
				"",
				"blank",
				"",
				attrValueTrue,
			),
		)
	} else {
		f.RemoveAttribute("blank", "")
	}
}

// CalendarType returns the calendar type for date filtering.
// Attribute: calendarType.
func (f *Filters) CalendarType() CalendarType {
	attr, found := f.GetAttribute(
		"calendarType",
		"",
	)
	if !found {
		return CalendarTypeNone
	}

	return CalendarType(attr.Value())
}

// SetCalendarType sets the calendar type for date filtering.
// Attribute: calendarType.
func (f *Filters) SetCalendarType(
	calType CalendarType,
) {
	if calType == "" ||
		calType == CalendarTypeNone {
		f.RemoveAttribute("calendarType", "")

		return
	}
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"calendarType",
			"",
			string(calType),
		),
	)
}

// GetFilters returns an iterator over all Filter elements.
func (f *Filters) GetFilters() iter.Seq[*Filter] {
	return func(yield func(*Filter) bool) {
		for child := range f.Children() {
			if child.LocalName() != "filter" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var filter *Filter
			if fl, ok := child.(*Filter); ok {
				filter = fl
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				filter = &Filter{LeafElementBase: leaf}
			}
			if filter != nil && !yield(filter) {
				return
			}
		}
	}
}

// FilterCount returns the number of Filter elements.
func (f *Filters) FilterCount() int {
	count := 0
	for range f.GetFilters() {
		count++
	}

	return count
}

// AddFilter adds a new Filter element with the given value.
func (f *Filters) AddFilter(
	value string,
) *Filter {
	filter := NewFilter()
	filter.SetVal(value)
	// Insert filters before dateGroupItem elements
	var insertBefore openxml.Element
	for child := range f.Children() {
		if child.LocalName() == "dateGroupItem" &&
			child.NamespaceURI() == NamespaceSML {
			insertBefore = child

			break
		}
	}
	if insertBefore != nil {
		f.InsertBefore(filter, insertBefore)
	} else {
		f.AppendChild(filter)
	}

	return filter
}

// RemoveFilter removes the filter with the given value.
func (f *Filters) RemoveFilter(
	value string,
) bool {
	for filter := range f.GetFilters() {
		if filter.Val() == value {
			return f.RemoveChild(filter)
		}
	}

	return false
}

// HasFilter returns whether a filter with the given value exists.
func (f *Filters) HasFilter(value string) bool {
	for filter := range f.GetFilters() {
		if filter.Val() == value {
			return true
		}
	}

	return false
}

// GetDateGroupItems returns an iterator over all DateGroupItem elements.
func (f *Filters) GetDateGroupItems() iter.Seq[*DateGroupItem] {
	return func(yield func(*DateGroupItem) bool) {
		for child := range f.Children() {
			if child.LocalName() != "dateGroupItem" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var dgi *DateGroupItem
			if d, ok := child.(*DateGroupItem); ok {
				dgi = d
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				dgi = &DateGroupItem{LeafElementBase: leaf}
			}
			if dgi != nil && !yield(dgi) {
				return
			}
		}
	}
}

// DateGroupItemCount returns the number of DateGroupItem elements.
func (f *Filters) DateGroupItemCount() int {
	count := 0
	for range f.GetDateGroupItems() {
		count++
	}

	return count
}

// AddDateGroupItem adds a new DateGroupItem element.
func (f *Filters) AddDateGroupItem(
	dateTimeGrouping DateTimeGrouping,
	year uint16,
) *DateGroupItem {
	dgi := NewDateGroupItem()
	dgi.SetDateTimeGrouping(dateTimeGrouping)
	dgi.SetYear(year)
	f.AppendChild(dgi)

	return dgi
}

// ClearAll removes all filter and dateGroupItem children.
func (f *Filters) ClearAll() {
	var toRemove []openxml.Element
	for child := range f.Children() {
		localName := child.LocalName()
		if (localName == "filter" || localName == "dateGroupItem") &&
			child.NamespaceURI() == NamespaceSML {
			toRemove = append(toRemove, child)
		}
	}
	for _, elem := range toRemove {
		f.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this Filters element.
func (f *Filters) Clone() openxml.Element {
	cloned := f.CompositeElementBase.Clone()

	return &Filters{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Filters element.
func (f *Filters) CloneNode(
	deep bool,
) openxml.Element {
	cloned := f.CompositeElementBase.CloneNode(
		deep,
	)

	return &Filters{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Filter represents a single filter value element (x:filter).
type Filter struct {
	*openxml.LeafElementBase
}

// NewFilter creates a new Filter element.
func NewFilter() *Filter {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"filter",
		PrefixDefault,
	)

	return &Filter{LeafElementBase: elem}
}

// Val returns the filter value. Attribute: val.
func (f *Filter) Val() string {
	attr, found := f.GetAttribute("val", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetVal sets the filter value. Attribute: val.
func (f *Filter) SetVal(value string) {
	if value == "" {
		f.RemoveAttribute("val", "")

		return
	}
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			value,
		),
	)
}

// Clone creates a deep copy of this Filter element.
func (f *Filter) Clone() openxml.Element {
	cloned := f.LeafElementBase.Clone()

	return &Filter{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Filter element.
func (f *Filter) CloneNode(
	deep bool,
) openxml.Element {
	cloned := f.LeafElementBase.CloneNode(deep)

	return &Filter{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// DateGroupItem represents a date grouping item element (x:dateGroupItem).
type DateGroupItem struct {
	*openxml.LeafElementBase
}

// NewDateGroupItem creates a new DateGroupItem element.
func NewDateGroupItem() *DateGroupItem {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		elemNameDateGroupItem,
		PrefixDefault,
	)

	return &DateGroupItem{LeafElementBase: elem}
}

// Year returns the year value. Attribute: year.
func (d *DateGroupItem) Year() uint16 {
	attr, found := d.GetAttribute("year", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize16,
	)

	return uint16(val)
}

// SetYear sets the year value. Attribute: year.
func (d *DateGroupItem) SetYear(year uint16) {
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"year",
			"",
			strconv.FormatUint(
				uint64(year),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Month returns the month value (1-12). Attribute: month.
func (d *DateGroupItem) Month() uint8 {
	attr, found := d.GetAttribute("month", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize8,
	)

	return uint8(val)
}

// SetMonth sets the month value (1-12). Attribute: month.
func (d *DateGroupItem) SetMonth(month uint8) {
	if month == 0 {
		d.RemoveAttribute("month", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"month",
			"",
			strconv.FormatUint(
				uint64(month),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Day returns the day value (1-31). Attribute: day.
func (d *DateGroupItem) Day() uint8 {
	attr, found := d.GetAttribute("day", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize8,
	)

	return uint8(val)
}

// SetDay sets the day value (1-31). Attribute: day.
func (d *DateGroupItem) SetDay(day uint8) {
	if day == 0 {
		d.RemoveAttribute("day", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"day",
			"",
			strconv.FormatUint(
				uint64(day),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Hour returns the hour value (0-23). Attribute: hour.
func (d *DateGroupItem) Hour() uint8 {
	attr, found := d.GetAttribute("hour", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize8,
	)

	return uint8(val)
}

// SetHour sets the hour value (0-23). Attribute: hour.
func (d *DateGroupItem) SetHour(hour uint8) {
	if hour == 0 {
		d.RemoveAttribute("hour", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"hour",
			"",
			strconv.FormatUint(
				uint64(hour),
				10, //nolint:revive // add-constant: base 10
			),
		),
	)
}

// Minute returns the minute value (0-59). Attribute: minute.
func (d *DateGroupItem) Minute() uint8 {
	attr, found := d.GetAttribute("minute", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		8,  //nolint:revive // add-constant
	)

	return uint8(val)
}

// SetMinute sets the minute value (0-59). Attribute: minute.
func (d *DateGroupItem) SetMinute(minute uint8) {
	if minute == 0 {
		d.RemoveAttribute("minute", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"minute",
			"",
			strconv.FormatUint(
				uint64(minute),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Second returns the second value (0-59). Attribute: second.
func (d *DateGroupItem) Second() uint8 {
	attr, found := d.GetAttribute("second", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		8,  //nolint:revive // add-constant
	)

	return uint8(val)
}

// SetSecond sets the second value (0-59). Attribute: second.
func (d *DateGroupItem) SetSecond(second uint8) {
	if second == 0 {
		d.RemoveAttribute("second", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"second",
			"",
			strconv.FormatUint(
				uint64(second),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// DateTimeGrouping returns the grouping granularity.
// Attribute: dateTimeGrouping.
func (d *DateGroupItem) DateTimeGrouping() DateTimeGrouping {
	attr, found := d.GetAttribute(
		"dateTimeGrouping",
		"",
	)
	if !found {
		return DateTimeGroupingYear
	}

	return DateTimeGrouping(attr.Value())
}

// SetDateTimeGrouping sets the grouping granularity.
// Attribute: dateTimeGrouping.
func (d *DateGroupItem) SetDateTimeGrouping(
	grouping DateTimeGrouping,
) {
	if grouping == "" {
		d.RemoveAttribute("dateTimeGrouping", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"dateTimeGrouping",
			"",
			string(grouping),
		),
	)
}

// Clone creates a deep copy of this DateGroupItem element.
func (d *DateGroupItem) Clone() openxml.Element {
	cloned := d.LeafElementBase.Clone()

	return &DateGroupItem{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this DateGroupItem element.
func (d *DateGroupItem) CloneNode(
	deep bool,
) openxml.Element {
	cloned := d.LeafElementBase.CloneNode(deep)

	return &DateGroupItem{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
