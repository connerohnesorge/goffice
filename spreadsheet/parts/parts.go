//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"github.com/connerohnesorge/goffice/packaging"
)

// Content types for SpreadsheetML parts.
const (
	// Workbook content types
	ContentTypeWorkbook             = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"
	ContentTypeWorkbookTemplate     = "application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml"
	ContentTypeWorkbookMacroEnabled = "application/vnd.ms-excel.sheet.macroEnabled.main+xml"
	ContentTypeMacroTemplate        = "application/vnd.ms-excel.template.macroEnabled.main+xml"
	ContentTypeAddIn                = "application/vnd.ms-excel.addin.macroEnabled.main+xml"

	// Worksheet content types
	ContentTypeWorksheet   = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"
	ContentTypeChartsheet  = "application/vnd.openxmlformats-officedocument.spreadsheetml.chartsheet+xml"
	ContentTypeDialogsheet = "application/vnd.openxmlformats-officedocument.spreadsheetml.dialogsheet+xml"
	ContentTypeMacrosheet  = "application/vnd.ms-excel.macrosheet+xml"

	// Data parts
	ContentTypeSharedStrings = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"
	ContentTypeStyles        = "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"
	ContentTypeCalcChain     = "application/vnd.openxmlformats-officedocument.spreadsheetml.calcChain+xml"
	ContentTypeConnections   = "application/vnd.openxmlformats-officedocument.spreadsheetml.connections+xml"

	// Table parts
	ContentTypeTable      = "application/vnd.openxmlformats-officedocument.spreadsheetml.table+xml"
	ContentTypeQueryTable = "application/vnd.openxmlformats-officedocument.spreadsheetml.queryTable+xml"

	// Pivot table parts
	ContentTypePivotTable           = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotTable+xml"
	ContentTypePivotCacheDefinition = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotCacheDefinition+xml"
	ContentTypePivotCacheRecords    = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotCacheRecords+xml"

	// Drawing parts
	ContentTypeDrawing = "application/vnd.openxmlformats-officedocument.drawing+xml"
	ContentTypeChart   = "application/vnd.openxmlformats-officedocument.drawingml.chart+xml"
	ContentTypeVml     = "application/vnd.openxmlformats-officedocument.vmlDrawing"

	// Comment parts
	ContentTypeComments = "application/vnd.openxmlformats-officedocument.spreadsheetml.comments+xml"

	// Theme parts
	ContentTypeTheme = "application/vnd.openxmlformats-officedocument.theme+xml"

	// VBA parts
	ContentTypeVbaProject = "application/vnd.ms-office.vbaProject"

	// External references
	ContentTypeExternalLink = "application/vnd.openxmlformats-officedocument.spreadsheetml.externalLink+xml"

	// Slicer parts
	ContentTypeSlicer      = "application/vnd.ms-excel.slicer+xml"
	ContentTypeSlicerCache = "application/vnd.ms-excel.slicerCache+xml"

	// Timeline parts
	ContentTypeTimeline      = "application/vnd.ms-excel.timeline+xml"
	ContentTypeTimelineCache = "application/vnd.ms-excel.timelineCache+xml"
)

// Relationship types for SpreadsheetML parts.
const (
	// Core relationships
	RelationshipTypeOfficeDocument = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	RelationshipTypeWorksheet      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet"
	RelationshipTypeChartsheet     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chartsheet"
	RelationshipTypeDialogsheet    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/dialogsheet"
	RelationshipTypeMacrosheet     = "http://schemas.microsoft.com/office/2006/relationships/xlMacrosheet"

	// PURL namespace variant for officeDocument (used by some tools)
	RelationshipTypePURLOfficeDocument = "http://purl.oclc.org/ooxml/officeDocument/relationships/officeDocument"

	// Data relationships
	RelationshipTypeSharedStrings = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/sharedStrings"
	RelationshipTypeStyles        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"
	RelationshipTypeCalcChain     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/calcChain"
	RelationshipTypeConnections   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/connections"

	// Table relationships
	RelationshipTypeTable      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/table"
	RelationshipTypeQueryTable = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/queryTable"

	// Pivot table relationships
	RelationshipTypePivotTable           = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/pivotTable"
	RelationshipTypePivotCacheDefinition = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/pivotCacheDefinition"
	RelationshipTypePivotCacheRecords    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/pivotCacheRecords"

	// Drawing relationships
	RelationshipTypeDrawing = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/drawing"
	RelationshipTypeChart   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chart"
	RelationshipTypeVml     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/vmlDrawing"
	RelationshipTypeImage   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"

	// Comment relationships
	RelationshipTypeComments = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"

	// Theme relationships
	RelationshipTypeTheme = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme"

	// VBA relationships
	RelationshipTypeVbaProject = "http://schemas.microsoft.com/office/2006/relationships/vbaProject"

	// External references
	RelationshipTypeExternalLink     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/externalLink"
	RelationshipTypeExternalLinkPath = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/externalLinkPath"

	// Slicer relationships
	RelationshipTypeSlicer      = "http://schemas.microsoft.com/office/2007/relationships/slicer"
	RelationshipTypeSlicerCache = "http://schemas.microsoft.com/office/2007/relationships/slicerCache"

	// Timeline relationships
	RelationshipTypeTimeline      = "http://schemas.microsoft.com/office/2013/relationships/timeline"
	RelationshipTypeTimelineCache = "http://schemas.microsoft.com/office/2013/relationships/timelineCache"

	// Hyperlink
	RelationshipTypeHyperlink = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"
)

// ErrNilPackage is returned when the package is nil.
var ErrNilPackage = partError("package is nil")

type partError string

func (e partError) Error() string {
	return string(e)
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func addChildPart(
	parent partWithPackage,
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	pkg := parent.Package()
	if pkg == nil {
		return nil, "", ErrNilPackage
	}

	// Create the underlying packaging part
	packPart, err := pkg.CreatePart(
		uri,
		contentType,
	)
	if err != nil {
		return nil, "", err
	}

	// Create relationship from parent to this part
	rel, err := pkg.CreatePartRelationship(
		parent.URI(),
		uri,
		relType,
		"",
	)
	if err != nil {
		return nil, "", err
	}

	return packPart, rel.ID(), nil
}

// partWithPackage is an interface for parts that can access the package.
type partWithPackage interface {
	Package() *packaging.Package
	URI() string
}
