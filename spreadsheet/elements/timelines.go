package elements

//revive:disable:file-length-limit many timeline types

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Timeline namespace constant.
const (
	// NamespaceTimelineX15 is the Excel 2013+ timeline namespace.
	//nolint:revive // line-length-limit
	NamespaceTimelineX15 = "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main"
)

// Timelines represents the timelines collection root element (x15:timelines).
// This element is the root of a timelines part.
type Timelines struct {
	*openxml.PartRootElementBase
}

// NewTimelines creates a new Timelines element.
func NewTimelines() *Timelines {
	elem := openxml.NewPartRootElement(
		NamespaceTimelineX15,
		"timelines",
		"x15",
	)

	return &Timelines{PartRootElementBase: elem}
}

// Timelines returns an iterator over all Timeline elements.
func (t *Timelines) Timelines() iter.Seq[*Timeline] {
	return func(yield func(*Timeline) bool) {
		for child := range t.Children() {
			if child.LocalName() != "timeline" {
				continue
			}
			var timeline *Timeline
			switch v := child.(type) {
			case *Timeline:
				timeline = v
			case *openxml.LeafElementBase:
				timeline = &Timeline{LeafElementBase: v}
			}
			if timeline != nil &&
				!yield(timeline) {
				return
			}
		}
	}
}

// TimelineCount returns the count of timeline elements.
func (t *Timelines) TimelineCount() int {
	count := 0
	for range t.Timelines() {
		count++
	}

	return count
}

// GetTimelineByName returns the timeline with the given name,
// or nil if not found.
func (t *Timelines) GetTimelineByName(
	name string,
) *Timeline {
	for timeline := range t.Timelines() {
		if timeline.Name() == name {
			return timeline
		}
	}

	return nil
}

// AddTimeline adds a new Timeline element.
func (t *Timelines) AddTimeline(
	name string,
) *Timeline {
	timeline := NewTimeline()
	timeline.SetName(name)
	t.AppendChild(timeline)

	return timeline
}

// RemoveTimeline removes a timeline from the collection.
func (t *Timelines) RemoveTimeline(
	timeline *Timeline,
) bool {
	return t.RemoveChild(timeline)
}

// Clone creates a deep copy of this Timelines element.
func (t *Timelines) Clone() openxml.Element {
	cloned := t.PartRootElementBase.Clone()

	return &Timelines{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this Timelines element.
func (t *Timelines) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.PartRootElementBase.CloneNode(
		deep,
	)

	return &Timelines{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// Timeline represents a single timeline element (x15:timeline).
type Timeline struct {
	*openxml.LeafElementBase
}

// NewTimeline creates a new Timeline element.
func NewTimeline() *Timeline {
	elem := openxml.NewLeafElement(
		NamespaceTimelineX15,
		"timeline",
		"x15",
	)

	return &Timeline{LeafElementBase: elem}
}

// Name returns the timeline name. Attribute: name.
func (t *Timeline) Name() string {
	attr, found := t.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the timeline name. Attribute: name.
func (t *Timeline) SetName(name string) {
	if name == "" {
		t.RemoveAttribute("name", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Cache returns the cache name. Attribute: cache.
func (t *Timeline) Cache() string {
	attr, found := t.GetAttribute("cache", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCache sets the cache name. Attribute: cache.
func (t *Timeline) SetCache(cache string) {
	if cache == "" {
		t.RemoveAttribute("cache", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"cache",
			"",
			cache,
		),
	)
}

// Caption returns the timeline caption. Attribute: caption.
func (t *Timeline) Caption() string {
	attr, found := t.GetAttribute("caption", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCaption sets the timeline caption. Attribute: caption.
func (t *Timeline) SetCaption(caption string) {
	if caption == "" {
		t.RemoveAttribute("caption", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"caption",
			"",
			caption,
		),
	)
}

// ShowHeader returns whether to show header. Attribute: showHeader.
func (t *Timeline) ShowHeader() bool {
	attr, found := t.GetAttribute(
		"showHeader",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowHeader sets whether to show header. Attribute: showHeader.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Timeline) SetShowHeader(value bool) {
	if value {
		t.RemoveAttribute(
			"showHeader",
			"",
		) // true is default
	} else {
		t.SetAttribute(
			openxml.NewAttribute("", "showHeader", "", attrValueFalse),
		)
	}
}

// ShowSelectionLabel returns whether to show selection label.
// Attribute: showSelectionLabel.
func (t *Timeline) ShowSelectionLabel() bool {
	attr, found := t.GetAttribute(
		"showSelectionLabel",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowSelectionLabel sets whether to show selection label.
// Attribute: showSelectionLabel.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Timeline) SetShowSelectionLabel(
	value bool,
) {
	if value {
		t.RemoveAttribute(
			"showSelectionLabel",
			"",
		) // true is default
	} else {
		t.SetAttribute(
			openxml.NewAttribute("", "showSelectionLabel", "", attrValueFalse),
		)
	}
}

// ShowTimeLevel returns whether to show time level. Attribute: showTimeLevel.
func (t *Timeline) ShowTimeLevel() bool {
	attr, found := t.GetAttribute(
		"showTimeLevel",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowTimeLevel sets whether to show time level. Attribute: showTimeLevel.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Timeline) SetShowTimeLevel(value bool) {
	if value {
		t.RemoveAttribute(
			"showTimeLevel",
			"",
		) // true is default
	} else {
		t.SetAttribute(
			openxml.NewAttribute("", "showTimeLevel", "", attrValueFalse),
		)
	}
}

// ShowHorizontalScrollbar returns whether to show horizontal scrollbar.
// Attribute: showHorizontalScrollbar.
func (t *Timeline) ShowHorizontalScrollbar() bool {
	attr, found := t.GetAttribute(
		"showHorizontalScrollbar",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowHorizontalScrollbar sets whether to show horizontal scrollbar.
// Attribute: showHorizontalScrollbar.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Timeline) SetShowHorizontalScrollbar(
	value bool,
) {
	if value {
		t.RemoveAttribute(
			"showHorizontalScrollbar",
			"",
		) // true is default
	} else {
		t.SetAttribute(
			openxml.NewAttribute("", "showHorizontalScrollbar", "", attrValueFalse),
		)
	}
}

// Level returns the time level. Attribute: level.
// Values: 0=Years, 1=Quarters, 2=Months, 3=Days
func (t *Timeline) Level() uint32 {
	attr, found := t.GetAttribute("level", "")
	if !found {
		return 2 // Default is 2 (Months)
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetLevel sets the time level. Attribute: level.
func (t *Timeline) SetLevel(level uint32) {
	if level == 2 {
		t.RemoveAttribute(
			"level",
			"",
		) // 2 is default

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"level",
			"",
			strconv.FormatUint(
				uint64(level),
				parseBase10,
			),
		),
	)
}

// SelectionLevel returns the selection level. Attribute: selectionLevel.
func (t *Timeline) SelectionLevel() uint32 {
	attr, found := t.GetAttribute(
		"selectionLevel",
		"",
	)
	if !found {
		return 2 // Default is 2 (Months)
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetSelectionLevel sets the selection level. Attribute: selectionLevel.
func (t *Timeline) SetSelectionLevel(
	level uint32,
) {
	if level == 2 {
		t.RemoveAttribute(
			"selectionLevel",
			"",
		) // 2 is default

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"selectionLevel",
			"",
			strconv.FormatUint(
				uint64(level),
				parseBase10,
			),
		),
	)
}

// ScrollPosition returns the scroll position. Attribute: scrollPosition.
func (t *Timeline) ScrollPosition() string {
	attr, found := t.GetAttribute(
		"scrollPosition",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetScrollPosition sets the scroll position (ISO 8601 date).
// Attribute: scrollPosition.
func (t *Timeline) SetScrollPosition(
	date string,
) {
	if date == "" {
		t.RemoveAttribute("scrollPosition", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"scrollPosition",
			"",
			date,
		),
	)
}

// Style returns the timeline style name. Attribute: style.
func (t *Timeline) Style() string {
	attr, found := t.GetAttribute("style", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetStyle sets the timeline style name. Attribute: style.
func (t *Timeline) SetStyle(style string) {
	if style == "" {
		t.RemoveAttribute("style", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"style",
			"",
			style,
		),
	)
}

// Clone creates a deep copy of this Timeline element.
func (t *Timeline) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &Timeline{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Timeline element.
func (t *Timeline) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.LeafElementBase.CloneNode(deep)

	return &Timeline{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// TimelineCacheDefinition represents the timeline cache definition root element
// (x15:timelineCacheDefinition).
type TimelineCacheDefinition struct {
	*openxml.PartRootElementBase
}

// NewTimelineCacheDefinition creates a new TimelineCacheDefinition element.
func NewTimelineCacheDefinition() *TimelineCacheDefinition {
	elem := openxml.NewPartRootElement(
		NamespaceTimelineX15,
		"timelineCacheDefinition",
		"x15",
	)

	return &TimelineCacheDefinition{
		PartRootElementBase: elem,
	}
}

// Name returns the cache name. Attribute: name.
func (t *TimelineCacheDefinition) Name() string {
	attr, found := t.GetAttribute(
		"name", //nolint:revive // add-constant
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the cache name. Attribute: name.
func (t *TimelineCacheDefinition) SetName(
	name string,
) {
	if name == "" {
		t.RemoveAttribute("name", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"name", //nolint:revive // add-constant: attribute name
			"",
			name,
		),
	)
}

// SourceName returns the source name. Attribute: sourceName.
func (t *TimelineCacheDefinition) SourceName() string {
	attr, found := t.GetAttribute(
		"sourceName",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSourceName sets the source name. Attribute: sourceName.
func (t *TimelineCacheDefinition) SetSourceName(
	name string,
) {
	if name == "" {
		t.RemoveAttribute("sourceName", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"sourceName",
			"",
			name,
		),
	)
}

// PivotTables returns the pivot tables element, or nil if not present.
func (t *TimelineCacheDefinition) PivotTables() *TimelineCachePivotTables {
	elem := t.GetElement(
		"pivotTables",
		NamespaceTimelineX15,
	)
	if elem == nil {
		return nil
	}
	if pt, ok := elem.(*TimelineCachePivotTables); ok {
		return pt
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TimelineCachePivotTables{
			CompositeElementBase: comp,
		}
	}

	return nil
}

//
//nolint:lll,revive // long function signature
func (t *TimelineCacheDefinition) GetOrCreatePivotTables() *TimelineCachePivotTables {
	pt := t.PivotTables()
	if pt != nil {
		return pt
	}
	pt = NewTimelineCachePivotTables()
	t.AppendChild(pt)

	return pt
}

// State returns the state element, or nil if not present.
func (t *TimelineCacheDefinition) State() *TimelineState {
	elem := t.GetElement(
		"state",
		NamespaceTimelineX15,
	)
	if elem == nil {
		return nil
	}
	if s, ok := elem.(*TimelineState); ok {
		return s
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &TimelineState{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreateState returns the state element, creating if needed.
func (t *TimelineCacheDefinition) GetOrCreateState() *TimelineState {
	s := t.State()
	if s != nil {
		return s
	}
	s = NewTimelineState()
	t.AppendChild(s)

	return s
}

// Clone creates a deep copy of this TimelineCacheDefinition element.
func (t *TimelineCacheDefinition) Clone() openxml.Element {
	cloned := t.PartRootElementBase.Clone()

	return &TimelineCacheDefinition{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this TimelineCacheDefinition element.
func (t *TimelineCacheDefinition) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.PartRootElementBase.CloneNode(
		deep,
	)

	return &TimelineCacheDefinition{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// TimelineCachePivotTables represents the pivot tables element
// (x15:pivotTables).
type TimelineCachePivotTables struct {
	*openxml.CompositeElementBase
}

// NewTimelineCachePivotTables creates a new TimelineCachePivotTables element.
func NewTimelineCachePivotTables() *TimelineCachePivotTables {
	elem := openxml.NewCompositeElement(
		NamespaceTimelineX15,
		"pivotTables",
		"x15", //nolint:revive // add-constant: namespace prefix
	)

	return &TimelineCachePivotTables{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this TimelineCachePivotTables element.
func (p *TimelineCachePivotTables) Clone() openxml.Element {
	cloned := p.CompositeElementBase.Clone()

	return &TimelineCachePivotTables{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this TimelineCachePivotTables element.
func (p *TimelineCachePivotTables) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.CompositeElementBase.CloneNode(
		deep,
	)

	return &TimelineCachePivotTables{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// TimelineState represents the state element (x15:state).
type TimelineState struct {
	*openxml.LeafElementBase
}

// NewTimelineState creates a new TimelineState element.
func NewTimelineState() *TimelineState {
	elem := openxml.NewLeafElement(
		NamespaceTimelineX15,
		"state",
		"x15",
	)

	return &TimelineState{LeafElementBase: elem}
}

// SingleRangeFilterState returns whether single range filter state.
// Attribute: singleRangeFilterState.
func (ts *TimelineState) SingleRangeFilterState() bool {
	attr, found := ts.GetAttribute(
		"singleRangeFilterState",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSingleRangeFilterState sets whether single range filter state.
// Attribute: singleRangeFilterState.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ts *TimelineState) SetSingleRangeFilterState(
	value bool,
) {
	if value {
		ts.SetAttribute(
			openxml.NewAttribute(
				"",
				"singleRangeFilterState",
				"",
				attrValueTrue,
			),
		)
	} else {
		ts.RemoveAttribute("singleRangeFilterState", "")
	}
}

// MinimalRefreshVersion returns the minimal refresh version.
// Attribute: minimalRefreshVersion.
func (ts *TimelineState) MinimalRefreshVersion() uint32 {
	attr, found := ts.GetAttribute(
		"minimalRefreshVersion",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetMinimalRefreshVersion sets the minimal refresh version.
// Attribute: minimalRefreshVersion.
func (ts *TimelineState) SetMinimalRefreshVersion(
	version uint32,
) {
	if version == 0 {
		ts.RemoveAttribute(
			"minimalRefreshVersion",
			"",
		)

		return
	}
	ts.SetAttribute(
		openxml.NewAttribute(
			"",
			"minimalRefreshVersion",
			"",
			strconv.FormatUint(
				uint64(version),
				parseBase10,
			),
		),
	)
}

// LastRefreshVersion returns the last refresh version.
// Attribute: lastRefreshVersion.
func (ts *TimelineState) LastRefreshVersion() uint32 {
	attr, found := ts.GetAttribute(
		"lastRefreshVersion",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetLastRefreshVersion sets the last refresh version.
// Attribute: lastRefreshVersion.
func (ts *TimelineState) SetLastRefreshVersion(
	version uint32,
) {
	if version == 0 {
		ts.RemoveAttribute(
			"lastRefreshVersion",
			"",
		)

		return
	}
	ts.SetAttribute(
		openxml.NewAttribute(
			"",
			"lastRefreshVersion",
			"",
			strconv.FormatUint(
				uint64(version),
				parseBase10,
			),
		),
	)
}

// PivotCacheId returns the pivot cache ID. Attribute: pivotCacheId.
func (ts *TimelineState) PivotCacheId() uint32 {
	attr, found := ts.GetAttribute(
		"pivotCacheId",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetPivotCacheId sets the pivot cache ID. Attribute: pivotCacheId.
func (ts *TimelineState) SetPivotCacheId(
	id uint32,
) {
	ts.SetAttribute(
		openxml.NewAttribute(
			"",
			"pivotCacheId",
			"",
			strconv.FormatUint(
				uint64(id),
				parseBase10,
			),
		),
	)
}

// FilterType returns the filter type. Attribute: filterType.
func (ts *TimelineState) FilterType() string {
	attr, found := ts.GetAttribute(
		"filterType",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetFilterType sets the filter type. Attribute: filterType.
func (ts *TimelineState) SetFilterType(
	filterType string,
) {
	if filterType == "" {
		ts.RemoveAttribute("filterType", "")

		return
	}
	ts.SetAttribute(
		openxml.NewAttribute(
			"",
			"filterType",
			"",
			filterType,
		),
	)
}

// Clone creates a deep copy of this TimelineState element.
func (ts *TimelineState) Clone() openxml.Element {
	cloned := ts.LeafElementBase.Clone()

	return &TimelineState{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this TimelineState element.
func (ts *TimelineState) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ts.LeafElementBase.CloneNode(deep)

	return &TimelineState{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
