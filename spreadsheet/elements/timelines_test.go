package elements

import (
	"testing"
)

func TestNewTimelines(t *testing.T) {
	timelines := NewTimelines()

	if timelines == nil {
		t.Fatal("NewTimelines returned nil")
	}

	if timelines.LocalName() != "timelines" {
		t.Errorf(
			"Expected LocalName 'timelines', got '%s'",
			timelines.LocalName(),
		)
	}

	if timelines.NamespaceURI() != NamespaceTimelineX15 {
		t.Errorf(
			"Expected NamespaceURI '%s', got '%s'",
			NamespaceTimelineX15,
			timelines.NamespaceURI(),
		)
	}
}

func TestTimelinesAddTimeline(t *testing.T) {
	timelines := NewTimelines()

	// Add a timeline with required name parameter
	timeline := timelines.AddTimeline(
		"TestTimeline",
	)
	if timeline == nil {
		t.Fatal("AddTimeline returned nil")
	}

	if timeline.LocalName() != "timeline" {
		t.Errorf(
			"Expected LocalName 'timeline', got '%s'",
			timeline.LocalName(),
		)
	}

	if timeline.Name() != "TestTimeline" {
		t.Errorf(
			"Expected Name 'TestTimeline', got '%s'",
			timeline.Name(),
		)
	}

	// Check timeline count
	if timelines.TimelineCount() != 1 {
		t.Errorf(
			"Expected TimelineCount 1, got %d",
			timelines.TimelineCount(),
		)
	}
}

func TestTimelinesIteration(t *testing.T) {
	timelines := NewTimelines()

	// Add timelines
	timelines.AddTimeline("Timeline1")
	timelines.AddTimeline("Timeline2")
	timelines.AddTimeline("Timeline3")

	// Iterate
	count := 0
	for timeline := range timelines.Timelines() {
		count++
		if timeline == nil {
			t.Error(
				"Iterator yielded nil timeline",
			)
		}
	}
	if count != 3 {
		t.Errorf(
			"Expected 3 timelines in iterator, got %d",
			count,
		)
	}
}

func TestTimelinesGetTimelineByName(
	t *testing.T,
) {
	timelines := NewTimelines()

	// Add timelines
	timelines.AddTimeline("First")
	timelines.AddTimeline("Second")

	// Get by name
	timeline := timelines.GetTimelineByName(
		"Second",
	)
	if timeline == nil {
		t.Fatal(
			"GetTimelineByName('Second') returned nil",
		)
	}
	if timeline.Name() != "Second" {
		t.Errorf(
			"Expected Name 'Second', got '%s'",
			timeline.Name(),
		)
	}

	// Non-existent name
	if timelines.GetTimelineByName(
		"NonExistent",
	) != nil {
		t.Error(
			"Expected nil for non-existent name",
		)
	}
}

func TestTimelineAttributes(t *testing.T) {
	timeline := NewTimeline()

	// Test Name
	timeline.SetName("DateTimeline")
	if timeline.Name() != "DateTimeline" {
		t.Errorf(
			"Expected Name 'DateTimeline', got '%s'",
			timeline.Name(),
		)
	}

	// Test Cache
	timeline.SetCache("TimelineCache1")
	if timeline.Cache() != "TimelineCache1" {
		t.Errorf(
			"Expected Cache 'TimelineCache1', got '%s'",
			timeline.Cache(),
		)
	}

	// Test Caption
	timeline.SetCaption("Date Range")
	if timeline.Caption() != "Date Range" {
		t.Errorf(
			"Expected Caption 'Date Range', got '%s'",
			timeline.Caption(),
		)
	}

	// Test Style
	timeline.SetStyle("TimelineStyleLight1")
	if timeline.Style() != "TimelineStyleLight1" {
		t.Errorf(
			"Expected Style 'TimelineStyleLight1', got '%s'",
			timeline.Style(),
		)
	}
}

func TestTimelineNumericAttributes(t *testing.T) {
	timeline := NewTimeline()

	// Test Level
	timeline.SetLevel(2)
	if timeline.Level() != 2 {
		t.Errorf(
			"Expected Level 2, got %d",
			timeline.Level(),
		)
	}

	// Test SelectionLevel
	timeline.SetSelectionLevel(1)
	if timeline.SelectionLevel() != 1 {
		t.Errorf(
			"Expected SelectionLevel 1, got %d",
			timeline.SelectionLevel(),
		)
	}

	// Test ScrollPosition
	timeline.SetScrollPosition(
		"2023-01-01T00:00:00",
	)
	if timeline.ScrollPosition() != "2023-01-01T00:00:00" {
		t.Errorf(
			"Expected ScrollPosition '2023-01-01T00:00:00', got '%s'",
			timeline.ScrollPosition(),
		)
	}
}

func TestTimelineBooleanAttributes(t *testing.T) {
	timeline := NewTimeline()

	// Test ShowHeader (default true)
	if !timeline.ShowHeader() {
		t.Error(
			"Expected ShowHeader true by default",
		)
	}
	timeline.SetShowHeader(false)
	if timeline.ShowHeader() {
		t.Error("Expected ShowHeader false")
	}

	// Test ShowSelectionLabel (default true)
	if !timeline.ShowSelectionLabel() {
		t.Error(
			"Expected ShowSelectionLabel true by default",
		)
	}
	timeline.SetShowSelectionLabel(false)
	if timeline.ShowSelectionLabel() {
		t.Error(
			"Expected ShowSelectionLabel false",
		)
	}

	// Test ShowTimeLevel (default true)
	if !timeline.ShowTimeLevel() {
		t.Error(
			"Expected ShowTimeLevel true by default",
		)
	}
	timeline.SetShowTimeLevel(false)
	if timeline.ShowTimeLevel() {
		t.Error("Expected ShowTimeLevel false")
	}

	// Test ShowHorizontalScrollbar (default true)
	if !timeline.ShowHorizontalScrollbar() {
		t.Error(
			"Expected ShowHorizontalScrollbar true by default",
		)
	}
	timeline.SetShowHorizontalScrollbar(false)
	if timeline.ShowHorizontalScrollbar() {
		t.Error(
			"Expected ShowHorizontalScrollbar false",
		)
	}
}

func TestTimelineCacheDefinition(t *testing.T) {
	tcd := NewTimelineCacheDefinition()

	if tcd == nil {
		t.Fatal(
			"NewTimelineCacheDefinition returned nil",
		)
	}

	if tcd.LocalName() != "timelineCacheDefinition" {
		t.Errorf(
			"Expected LocalName 'timelineCacheDefinition', got '%s'",
			tcd.LocalName(),
		)
	}

	// Test Name
	tcd.SetName("TimelineCache1")
	if tcd.Name() != "TimelineCache1" {
		t.Errorf(
			"Expected Name 'TimelineCache1', got '%s'",
			tcd.Name(),
		)
	}

	// Test SourceName
	tcd.SetSourceName("OrderDate")
	if tcd.SourceName() != "OrderDate" {
		t.Errorf(
			"Expected SourceName 'OrderDate', got '%s'",
			tcd.SourceName(),
		)
	}
}

func TestTimelineCacheDefinitionPivotTables(
	t *testing.T,
) {
	tcd := NewTimelineCacheDefinition()

	// Initially nil
	if tcd.PivotTables() != nil {
		t.Error(
			"Expected PivotTables to be nil initially",
		)
	}

	// Create pivot tables
	pt := tcd.GetOrCreatePivotTables()
	if pt == nil {
		t.Fatal(
			"GetOrCreatePivotTables returned nil",
		)
	}

	// Get existing
	pt2 := tcd.GetOrCreatePivotTables()
	if pt2 != pt {
		t.Error(
			"Expected GetOrCreatePivotTables to return existing element",
		)
	}
}

func TestTimelineCacheDefinitionState(
	t *testing.T,
) {
	tcd := NewTimelineCacheDefinition()

	// Initially nil
	if tcd.State() != nil {
		t.Error(
			"Expected State to be nil initially",
		)
	}

	// Create state
	state := tcd.GetOrCreateState()
	if state == nil {
		t.Fatal("GetOrCreateState returned nil")
	}

	// Get existing
	state2 := tcd.GetOrCreateState()
	if state2 != state {
		t.Error(
			"Expected GetOrCreateState to return existing element",
		)
	}
}

func TestTimelineCachePivotTables(t *testing.T) {
	tcpt := NewTimelineCachePivotTables()

	if tcpt == nil {
		t.Fatal(
			"NewTimelineCachePivotTables returned nil",
		)
	}

	if tcpt.LocalName() != "pivotTables" {
		t.Errorf(
			"Expected LocalName 'pivotTables', got '%s'",
			tcpt.LocalName(),
		)
	}
}

func TestTimelineState(t *testing.T) {
	ts := NewTimelineState()

	if ts == nil {
		t.Fatal("NewTimelineState returned nil")
	}

	if ts.LocalName() != "state" {
		t.Errorf(
			"Expected LocalName 'state', got '%s'",
			ts.LocalName(),
		)
	}

	// Test SingleRangeFilterState
	ts.SetSingleRangeFilterState(true)
	if !ts.SingleRangeFilterState() {
		t.Error(
			"Expected SingleRangeFilterState true",
		)
	}

	// Test MinimalRefreshVersion
	ts.SetMinimalRefreshVersion(5)
	if ts.MinimalRefreshVersion() != 5 {
		t.Errorf(
			"Expected MinimalRefreshVersion 5, got %d",
			ts.MinimalRefreshVersion(),
		)
	}

	// Test LastRefreshVersion
	ts.SetLastRefreshVersion(6)
	if ts.LastRefreshVersion() != 6 {
		t.Errorf(
			"Expected LastRefreshVersion 6, got %d",
			ts.LastRefreshVersion(),
		)
	}

	// Test PivotCacheId
	ts.SetPivotCacheId(1)
	if ts.PivotCacheId() != 1 {
		t.Errorf(
			"Expected PivotCacheId 1, got %d",
			ts.PivotCacheId(),
		)
	}

	// Test FilterType
	ts.SetFilterType("unknown")
	if ts.FilterType() != "unknown" {
		t.Errorf(
			"Expected FilterType 'unknown', got '%s'",
			ts.FilterType(),
		)
	}
}

func TestTimelinesClone(t *testing.T) {
	timelines := NewTimelines()
	tl := timelines.AddTimeline("TestTimeline")
	tl.SetCaption("Test")

	clonedResult := timelines.Clone()
	cloned, ok := clonedResult.(*Timelines)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Timelines",
			clonedResult,
		)
	}

	if cloned == timelines {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.TimelineCount() != 1 {
		t.Errorf(
			"Cloned TimelineCount should be 1, got %d",
			cloned.TimelineCount(),
		)
	}
}

func TestTimelineClone(t *testing.T) {
	timeline := NewTimeline()
	timeline.SetName("MyTimeline")
	timeline.SetCaption("My Caption")
	timeline.SetLevel(3)

	clonedResult := timeline.Clone()
	cloned, ok := clonedResult.(*Timeline)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Timeline",
			clonedResult,
		)
	}

	if cloned == timeline {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.Name() != "MyTimeline" {
		t.Errorf(
			"Cloned Name should be 'MyTimeline', got '%s'",
			cloned.Name(),
		)
	}

	if cloned.Caption() != "My Caption" {
		t.Errorf(
			"Cloned Caption should be 'My Caption', got '%s'",
			cloned.Caption(),
		)
	}

	if cloned.Level() != 3 {
		t.Errorf(
			"Cloned Level should be 3, got %d",
			cloned.Level(),
		)
	}
}

func TestTimelineCacheDefinitionClone(
	t *testing.T,
) {
	tcd := NewTimelineCacheDefinition()
	tcd.SetName("CacheName")
	tcd.SetSourceName("DateField")

	clonedResult := tcd.Clone()
	cloned, ok := clonedResult.(*TimelineCacheDefinition)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *TimelineCacheDefinition",
			clonedResult,
		)
	}

	if cloned == tcd {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.Name() != "CacheName" {
		t.Errorf(
			"Cloned Name should be 'CacheName', got '%s'",
			cloned.Name(),
		)
	}

	if cloned.SourceName() != "DateField" {
		t.Errorf(
			"Cloned SourceName should be 'DateField', got '%s'",
			cloned.SourceName(),
		)
	}
}

func TestTimelineCachePivotTablesClone(
	t *testing.T,
) {
	tcpt := NewTimelineCachePivotTables()

	clonedResult := tcpt.Clone()
	cloned, ok := clonedResult.(*TimelineCachePivotTables)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *TimelineCachePivotTables",
			clonedResult,
		)
	}

	if cloned == tcpt {
		t.Error(
			"Clone should return a new instance",
		)
	}
}

func TestTimelineStateClone(t *testing.T) {
	ts := NewTimelineState()
	ts.SetPivotCacheId(5)
	ts.SetFilterType("clear")

	clonedResult := ts.Clone()
	cloned, ok := clonedResult.(*TimelineState)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *TimelineState",
			clonedResult,
		)
	}

	if cloned == ts {
		t.Error(
			"Clone should return a new instance",
		)
	}

	if cloned.PivotCacheId() != 5 {
		t.Errorf(
			"Cloned PivotCacheId should be 5, got %d",
			cloned.PivotCacheId(),
		)
	}

	if cloned.FilterType() != "clear" {
		t.Errorf(
			"Cloned FilterType should be 'clear', got '%s'",
			cloned.FilterType(),
		)
	}
}
