//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// TimeLinePart represents a timeline part (xl/timelines/timeline1.xml, etc.).
// Timelines provide date-based filtering controls for pivot tables.
type TimeLinePart struct {
	*openxml.OpenXmlPartData
}

// newTimeLinePart creates a new timeline part.
func newTimeLinePart(
	worksheetPart *WorksheetPart,
	uri string,
) (*TimeLinePart, error) {
	packPart, relID, err := worksheetPart.addChildPart(
		uri,
		ContentTypeTimeline,
		RelationshipTypeTimeline,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTimeline,
		packPart,
		worksheetPart,
	)
	partData.SetRelationshipID(relID)

	tp := &TimeLinePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal timeline content
	tp.initializeContent()

	// Add to worksheet part's child parts
	if err := worksheetPart.AddPart(tp, relID); err != nil {
		return nil, err
	}

	return tp, nil
}

// initializeContent sets up minimal timeline content.
func (tp *TimeLinePart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<timelines xmlns="http://schemas.microsoft.com/office/spreadsheetml/2010/11/main" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" mc:Ignorable="x">
</timelines>`
	tp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*TimeLinePart) FixedContentType() string {
	return ContentTypeTimeline
}

// Timelines returns the root Timelines element.
func (tp *TimeLinePart) Timelines() *elements.Timelines {
	if root := tp.RootElement(); root != nil {
		if t, ok := root.(*elements.Timelines); ok {
			return t
		}
	}

	return nil
}

// Counter for generating unique timeline cache filenames.
var timelineCacheCounter uint64

// AddTimeLineCachePart adds a timeline cache part.
func (tp *TimeLinePart) AddTimeLineCachePart() (*TimeLineCachePart, error) {
	num := atomic.AddUint64(
		&timelineCacheCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/xl/timelineCaches/timelineCache%d.xml",
		num,
	)

	return newTimeLineCachePart(tp, uri)
}

// TimeLineCachePart returns the timeline cache part if present.
func (tp *TimeLinePart) TimeLineCachePart() *TimeLineCachePart {
	for part := range tp.Parts() {
		if tcp, ok := part.(*TimeLineCachePart); ok {
			return tcp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (tp *TimeLinePart) GetStream() io.Reader {
	return tp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (tp *TimeLinePart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		tp,
		uri,
		contentType,
		relType,
	)
}

// Ensure TimeLinePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*TimeLinePart)(nil)

// TimeLineCachePart represents a timeline cache part (xl/timelineCaches/timelineCache1.xml, etc.).
// The timeline cache stores the cached data and selection state for a timeline.
type TimeLineCachePart struct {
	*openxml.OpenXmlPartData
}

// newTimeLineCachePart creates a new timeline cache part.
func newTimeLineCachePart(
	timelinePart *TimeLinePart,
	uri string,
) (*TimeLineCachePart, error) {
	packPart, relID, err := timelinePart.addChildPart(
		uri,
		ContentTypeTimelineCache,
		RelationshipTypeTimelineCache,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTimelineCache,
		packPart,
		timelinePart,
	)
	partData.SetRelationshipID(relID)

	tcp := &TimeLineCachePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal timeline cache content
	tcp.initializeContent()

	// Add to timeline part's child parts
	if err := timelinePart.AddPart(tcp, relID); err != nil {
		return nil, err
	}

	return tcp, nil
}

// initializeContent sets up minimal timeline cache content.
func (tcp *TimeLineCachePart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<timelineCacheDefinition xmlns="http://schemas.microsoft.com/office/spreadsheetml/2010/11/main" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" mc:Ignorable="x" name="Timeline_Date">
</timelineCacheDefinition>`
	tcp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*TimeLineCachePart) FixedContentType() string {
	return ContentTypeTimelineCache
}

// TimelineCacheDefinition returns the root TimelineCacheDefinition element.
func (tcp *TimeLineCachePart) TimelineCacheDefinition() *elements.TimelineCacheDefinition {
	if root := tcp.RootElement(); root != nil {
		if tcd, ok := root.(*elements.TimelineCacheDefinition); ok {
			return tcd
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (tcp *TimeLineCachePart) GetStream() io.Reader {
	return tcp.OpenXmlPartData.GetStream()
}

// Ensure TimeLineCachePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*TimeLineCachePart)(
	nil,
)

// TimeLinePartFactory creates a TimeLinePart from a URI and container.
func TimeLinePartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Use GetPackagingPart instead of Package() to avoid deadlock during loading
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTimeline,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewTimelines()
		},
	)

	return &TimeLinePart{
		OpenXmlPartData: partData,
	}
}

// TimeLineCachePartFactory creates a TimeLineCachePart from a URI and container.
func TimeLineCachePartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Use GetPackagingPart instead of Package() to avoid deadlock during loading
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeTimelineCache,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewTimelineCacheDefinition()
		},
	)

	return &TimeLineCachePart{
		OpenXmlPartData: partData,
	}
}

// Register the timeline part types.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeTimeline,
			RelationshipType:   RelationshipTypeTimeline,
			Factory:            TimeLinePartFactory,
			DefaultURI:         "/xl/timelines/timeline1.xml",
			IsFixedContentType: true,
		},
	)

	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeTimelineCache,
			RelationshipType:   RelationshipTypeTimelineCache,
			Factory:            TimeLineCachePartFactory,
			DefaultURI:         "/xl/timelineCaches/timelineCache1.xml",
			IsFixedContentType: true,
		},
	)
}
