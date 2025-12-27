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

// WorksheetPart represents a worksheet part (xl/worksheets/sheet1.xml, etc.).
type WorksheetPart struct {
	*openxml.OpenXmlPartData
}

// newWorksheetPart creates a new worksheet part.
func newWorksheetPart(
	workbookPart *WorkbookPart,
	uri string,
) (*WorksheetPart, error) {
	packPart, relID, err := workbookPart.addChildPart(
		uri,
		ContentTypeWorksheet,
		RelationshipTypeWorksheet,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeWorksheet,
		packPart,
		workbookPart,
	)
	partData.SetRelationshipID(relID)

	wsp := &WorksheetPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal worksheet content
	wsp.initializeContent()

	// Add to workbook part's child parts
	if err := workbookPart.AddPart(wsp, relID); err != nil {
		return nil, err
	}

	return wsp, nil
}

// initializeContent sets up minimal worksheet content.
func (wsp *WorksheetPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<worksheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
  <sheetData/>
</worksheet>`
	wsp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*WorksheetPart) FixedContentType() string {
	return ContentTypeWorksheet
}

// Worksheet returns the root Worksheet element.
func (wsp *WorksheetPart) Worksheet() *elements.Worksheet {
	if root := wsp.RootElement(); root != nil {
		if ws, ok := root.(*elements.Worksheet); ok {
			return ws
		}
	}

	return nil
}

// Counters for generating unique filenames
var (
	drawingCounter    uint64
	tableCounter      uint64
	pivotTableCounter uint64
	commentsCounter   uint64
	vmlDrawingCounter uint64
	queryTableCounter uint64
	slicerCounter     uint64
	timelineCounter   uint64
)

// AddDrawingsPart adds a drawings part to this worksheet.
func (wsp *WorksheetPart) AddDrawingsPart() (*DrawingsPart, error) {
	num := atomic.AddUint64(&drawingCounter, 1)
	uri := fmt.Sprintf(
		"/xl/drawings/drawing%d.xml",
		num,
	)

	return newDrawingsPart(wsp, uri)
}

// DrawingsPart returns the drawings part if present.
func (wsp *WorksheetPart) DrawingsPart() *DrawingsPart {
	for part := range wsp.Parts() {
		if dp, ok := part.(*DrawingsPart); ok {
			return dp
		}
	}

	return nil
}

// AddTableDefinitionPart adds a table definition part to this worksheet.
func (wsp *WorksheetPart) AddTableDefinitionPart() (*TableDefinitionPart, error) {
	num := atomic.AddUint64(&tableCounter, 1)
	uri := fmt.Sprintf(
		"/xl/tables/table%d.xml",
		num,
	)

	return newTableDefinitionPart(wsp, uri)
}

// TableDefinitionParts returns all table definition parts.
func (wsp *WorksheetPart) TableDefinitionParts() []*TableDefinitionPart {
	var tables []*TableDefinitionPart
	for part := range wsp.Parts() {
		if tp, ok := part.(*TableDefinitionPart); ok {
			tables = append(tables, tp)
		}
	}

	return tables
}

// AddPivotTablePart adds a pivot table part to this worksheet.
func (wsp *WorksheetPart) AddPivotTablePart() (*PivotTablePart, error) {
	num := atomic.AddUint64(&pivotTableCounter, 1)
	uri := fmt.Sprintf(
		"/xl/pivotTables/pivotTable%d.xml",
		num,
	)

	return newPivotTablePart(wsp, uri)
}

// PivotTableParts returns all pivot table parts.
func (wsp *WorksheetPart) PivotTableParts() []*PivotTablePart {
	var pivots []*PivotTablePart
	for part := range wsp.Parts() {
		if pp, ok := part.(*PivotTablePart); ok {
			pivots = append(pivots, pp)
		}
	}

	return pivots
}

// AddCommentsPart adds a comments part to this worksheet.
func (wsp *WorksheetPart) AddCommentsPart() (*WorksheetCommentsPart, error) {
	num := atomic.AddUint64(&commentsCounter, 1)
	uri := fmt.Sprintf("/xl/comments%d.xml", num)

	return newWorksheetCommentsPart(wsp, uri)
}

// CommentsPart returns the comments part if present.
func (wsp *WorksheetPart) CommentsPart() *WorksheetCommentsPart {
	for part := range wsp.Parts() {
		if cp, ok := part.(*WorksheetCommentsPart); ok {
			return cp
		}
	}

	return nil
}

// AddVmlDrawingPart adds a VML drawing part to this worksheet.
func (wsp *WorksheetPart) AddVmlDrawingPart() (*VmlDrawingPart, error) {
	num := atomic.AddUint64(&vmlDrawingCounter, 1)
	uri := fmt.Sprintf(
		"/xl/drawings/vmlDrawing%d.vml",
		num,
	)

	return newVmlDrawingPart(wsp, uri)
}

// VmlDrawingParts returns all VML drawing parts.
func (wsp *WorksheetPart) VmlDrawingParts() []*VmlDrawingPart {
	var vmls []*VmlDrawingPart
	for part := range wsp.Parts() {
		if vp, ok := part.(*VmlDrawingPart); ok {
			vmls = append(vmls, vp)
		}
	}

	return vmls
}

// AddQueryTablePart adds a query table part to this worksheet.
func (wsp *WorksheetPart) AddQueryTablePart() (*QueryTablePart, error) {
	num := atomic.AddUint64(&queryTableCounter, 1)
	uri := fmt.Sprintf(
		"/xl/queryTables/queryTable%d.xml",
		num,
	)

	return newQueryTablePart(wsp, uri)
}

// QueryTableParts returns all query table parts.
func (wsp *WorksheetPart) QueryTableParts() []*QueryTablePart {
	var queryTables []*QueryTablePart
	for part := range wsp.Parts() {
		if qp, ok := part.(*QueryTablePart); ok {
			queryTables = append(queryTables, qp)
		}
	}

	return queryTables
}

// AddSlicerPart adds a slicer part to this worksheet.
func (wsp *WorksheetPart) AddSlicerPart() (*SlicerPart, error) {
	num := atomic.AddUint64(&slicerCounter, 1)
	uri := fmt.Sprintf(
		"/xl/slicers/slicer%d.xml",
		num,
	)

	return newSlicerPart(wsp, uri)
}

// SlicerParts returns all slicer parts.
func (wsp *WorksheetPart) SlicerParts() []*SlicerPart {
	var slicers []*SlicerPart
	for part := range wsp.Parts() {
		if sp, ok := part.(*SlicerPart); ok {
			slicers = append(slicers, sp)
		}
	}

	return slicers
}

// AddTimelinePart adds a timeline part to this worksheet.
func (wsp *WorksheetPart) AddTimelinePart() (*TimeLinePart, error) {
	num := atomic.AddUint64(&timelineCounter, 1)
	uri := fmt.Sprintf(
		"/xl/timelines/timeline%d.xml",
		num,
	)

	return newTimeLinePart(wsp, uri)
}

// TimelineParts returns all timeline parts.
func (wsp *WorksheetPart) TimelineParts() []*TimeLinePart {
	var timelines []*TimeLinePart
	for part := range wsp.Parts() {
		if tp, ok := part.(*TimeLinePart); ok {
			timelines = append(timelines, tp)
		}
	}

	return timelines
}

// GetStream returns a reader for the part content.
func (wsp *WorksheetPart) GetStream() io.Reader {
	return wsp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (wsp *WorksheetPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		wsp,
		uri,
		contentType,
		relType,
	)
}

// Ensure WorksheetPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*WorksheetPart)(nil)

// WorksheetPartFactory creates a WorksheetPart from a URI and container.
func WorksheetPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeWorksheet,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewWorksheet()
		},
	)

	return &WorksheetPart{
		OpenXmlPartData: partData,
	}
}

// Register the WorksheetPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeWorksheet,
			RelationshipType:   RelationshipTypeWorksheet,
			Factory:            WorksheetPartFactory,
			DefaultURI:         "/xl/worksheets/sheet1.xml",
			IsFixedContentType: true,
		},
	)
}
