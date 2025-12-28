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

// PivotTablePart represents a pivot table part (xl/pivotTables/pivotTable1.xml, etc.).
type PivotTablePart struct {
	*openxml.OpenXmlPartData
}

// newPivotTablePart creates a new pivot table part.
func newPivotTablePart(
	worksheetPart *WorksheetPart,
	uri string,
) (*PivotTablePart, error) {
	packPart, relID, err := worksheetPart.addChildPart(
		uri,
		ContentTypePivotTable,
		RelationshipTypePivotTable,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypePivotTable,
		packPart,
		worksheetPart,
	)
	partData.SetRelationshipID(relID)

	pp := &PivotTablePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal pivot table content
	pp.initializeContent()

	// Add to worksheet part's child parts
	if err := worksheetPart.AddPart(pp, relID); err != nil {
		return nil, err
	}

	return pp, nil
}

// initializeContent sets up minimal pivot table content.
func (pp *PivotTablePart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<pivotTableDefinition xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" name="PivotTable1" cacheId="0" applyNumberFormats="0" applyBorderFormats="0" applyFontFormats="0" applyPatternFormats="0" applyAlignmentFormats="0" applyWidthHeightFormats="1" dataCaption="Values" updatedVersion="6" minRefreshableVersion="3" useAutoFormatting="1" itemPrintTitles="1" createdVersion="6" indent="0" outline="1" outlineData="1" multipleFieldFilters="0">
  <location ref="A3:B4" firstHeaderRow="1" firstDataRow="1" firstDataCol="1"/>
  <pivotFields count="0"/>
</pivotTableDefinition>`
	pp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*PivotTablePart) FixedContentType() string {
	return ContentTypePivotTable
}

// PivotTableDefinition returns the root PivotTableDefinition element.
func (pp *PivotTablePart) PivotTableDefinition() *elements.PivotTableDefinition {
	if root := pp.RootElement(); root != nil {
		if ptd, ok := root.(*elements.PivotTableDefinition); ok {
			return ptd
		}
	}

	return nil
}

// Counter for generating unique pivot cache filenames.
var pivotCacheCounter uint64

// AddPivotTableCacheDefinitionPart adds a pivot cache definition part.
func (pp *PivotTablePart) AddPivotTableCacheDefinitionPart() (*PivotTableCacheDefinitionPart, error) {
	num := atomic.AddUint64(&pivotCacheCounter, 1)
	uri := fmt.Sprintf(
		"/xl/pivotCache/pivotCacheDefinition%d.xml",
		num,
	)

	return newPivotTableCacheDefinitionPart(
		pp,
		uri,
	)
}

// PivotTableCacheDefinitionPart returns the pivot cache definition part if present.
func (pp *PivotTablePart) PivotTableCacheDefinitionPart() *PivotTableCacheDefinitionPart {
	for part := range pp.Parts() {
		if pcp, ok := part.(*PivotTableCacheDefinitionPart); ok {
			return pcp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (pp *PivotTablePart) GetStream() io.Reader {
	return pp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (pp *PivotTablePart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		pp,
		uri,
		contentType,
		relType,
	)
}

// Ensure PivotTablePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*PivotTablePart)(nil)

// PivotTableCacheDefinitionPart represents a pivot cache definition part.
type PivotTableCacheDefinitionPart struct {
	*openxml.OpenXmlPartData
}

// newPivotTableCacheDefinitionPart creates a new pivot cache definition part.
func newPivotTableCacheDefinitionPart(
	pivotTablePart *PivotTablePart,
	uri string,
) (*PivotTableCacheDefinitionPart, error) {
	packPart, relID, err := pivotTablePart.addChildPart(
		uri,
		ContentTypePivotCacheDefinition,
		RelationshipTypePivotCacheDefinition,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypePivotCacheDefinition,
		packPart,
		pivotTablePart,
	)
	partData.SetRelationshipID(relID)

	pcp := &PivotTableCacheDefinitionPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal pivot cache definition content
	pcp.initializeContent()

	// Add to pivot table part's child parts
	if err := pivotTablePart.AddPart(pcp, relID); err != nil {
		return nil, err
	}

	return pcp, nil
}

// initializeContent sets up minimal pivot cache definition content.
func (pcp *PivotTableCacheDefinitionPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<pivotCacheDefinition xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" refreshedBy="User" refreshedDate="0" createdVersion="6" refreshedVersion="6" minRefreshableVersion="3" recordCount="0">
  <cacheSource type="worksheet">
    <worksheetSource ref="A1:A1"/>
  </cacheSource>
  <cacheFields count="0"/>
</pivotCacheDefinition>`
	pcp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*PivotTableCacheDefinitionPart) FixedContentType() string {
	return ContentTypePivotCacheDefinition
}

// PivotCacheDefinition returns the root PivotCacheDefinition element.
func (pcp *PivotTableCacheDefinitionPart) PivotCacheDefinition() *elements.PivotCacheDefinition {
	if root := pcp.RootElement(); root != nil {
		if pcd, ok := root.(*elements.PivotCacheDefinition); ok {
			return pcd
		}
	}

	return nil
}

// Counter for generating unique pivot cache records filenames.
var pivotCacheRecordsCounter uint64

// AddPivotTableCacheRecordsPart adds a pivot cache records part.
func (pcp *PivotTableCacheDefinitionPart) AddPivotTableCacheRecordsPart() (*PivotTableCacheRecordsPart, error) {
	num := atomic.AddUint64(
		&pivotCacheRecordsCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/xl/pivotCache/pivotCacheRecords%d.xml",
		num,
	)

	return newPivotTableCacheRecordsPart(pcp, uri)
}

// PivotTableCacheRecordsPart returns the pivot cache records part if present.
func (pcp *PivotTableCacheDefinitionPart) PivotTableCacheRecordsPart() *PivotTableCacheRecordsPart {
	for part := range pcp.Parts() {
		if prp, ok := part.(*PivotTableCacheRecordsPart); ok {
			return prp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (pcp *PivotTableCacheDefinitionPart) GetStream() io.Reader {
	return pcp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (pcp *PivotTableCacheDefinitionPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		pcp,
		uri,
		contentType,
		relType,
	)
}

// Ensure PivotTableCacheDefinitionPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*PivotTableCacheDefinitionPart)(
	nil,
)

// PivotTableCacheRecordsPart represents a pivot cache records part.
type PivotTableCacheRecordsPart struct {
	*openxml.OpenXmlPartData
}

// newPivotTableCacheRecordsPart creates a new pivot cache records part.
func newPivotTableCacheRecordsPart(
	cachePart *PivotTableCacheDefinitionPart,
	uri string,
) (*PivotTableCacheRecordsPart, error) {
	packPart, relID, err := cachePart.addChildPart(
		uri,
		ContentTypePivotCacheRecords,
		RelationshipTypePivotCacheRecords,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypePivotCacheRecords,
		packPart,
		cachePart,
	)
	partData.SetRelationshipID(relID)

	prp := &PivotTableCacheRecordsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal pivot cache records content
	prp.initializeContent()

	// Add to cache part's child parts
	if err := cachePart.AddPart(prp, relID); err != nil {
		return nil, err
	}

	return prp, nil
}

// initializeContent sets up minimal pivot cache records content.
func (prp *PivotTableCacheRecordsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<pivotCacheRecords xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" count="0">
</pivotCacheRecords>`
	prp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*PivotTableCacheRecordsPart) FixedContentType() string {
	return ContentTypePivotCacheRecords
}

// PivotCacheRecords returns the root PivotCacheRecords element.
func (prp *PivotTableCacheRecordsPart) PivotCacheRecords() *elements.PivotCacheRecords {
	if root := prp.RootElement(); root != nil {
		if pcr, ok := root.(*elements.PivotCacheRecords); ok {
			return pcr
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (prp *PivotTableCacheRecordsPart) GetStream() io.Reader {
	return prp.OpenXmlPartData.GetStream()
}

// Ensure PivotTableCacheRecordsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*PivotTableCacheRecordsPart)(
	nil,
)

// PivotTablePartFactory creates a PivotTablePart from a URI and container.
func PivotTablePartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Use GetPackagingPart which is safe during initialization (no locks)
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypePivotTable,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewPivotTableDefinition()
		},
	)

	return &PivotTablePart{
		OpenXmlPartData: partData,
	}
}

// PivotTableCacheDefinitionPartFactory creates a PivotTableCacheDefinitionPart from a URI and container.
func PivotTableCacheDefinitionPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Use GetPackagingPart which is safe during initialization (no locks)
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypePivotCacheDefinition,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewPivotCacheDefinition()
		},
	)

	return &PivotTableCacheDefinitionPart{
		OpenXmlPartData: partData,
	}
}

// PivotTableCacheRecordsPartFactory creates a PivotTableCacheRecordsPart from a URI and container.
func PivotTableCacheRecordsPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Use GetPackagingPart which is safe during initialization (no locks)
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypePivotCacheRecords,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewPivotCacheRecords()
		},
	)

	return &PivotTableCacheRecordsPart{
		OpenXmlPartData: partData,
	}
}

// Register the pivot table part types.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypePivotTable,
			RelationshipType:   RelationshipTypePivotTable,
			Factory:            PivotTablePartFactory,
			DefaultURI:         "/xl/pivotTables/pivotTable1.xml",
			IsFixedContentType: true,
		},
	)

	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypePivotCacheDefinition,
			RelationshipType:   RelationshipTypePivotCacheDefinition,
			Factory:            PivotTableCacheDefinitionPartFactory,
			DefaultURI:         "/xl/pivotCache/pivotCacheDefinition1.xml",
			IsFixedContentType: true,
		},
	)

	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypePivotCacheRecords,
			RelationshipType:   RelationshipTypePivotCacheRecords,
			Factory:            PivotTableCacheRecordsPartFactory,
			DefaultURI:         "/xl/pivotCache/pivotCacheRecords1.xml",
			IsFixedContentType: true,
		},
	)
}
