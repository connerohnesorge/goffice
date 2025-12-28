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

// SlicerPart represents a slicer part (xl/slicers/slicer1.xml, etc.).
// Slicers provide visual filtering controls for pivot tables and tables.
type SlicerPart struct {
	*openxml.OpenXmlPartData
}

// newSlicerPart creates a new slicer part.
func newSlicerPart(
	worksheetPart *WorksheetPart,
	uri string,
) (*SlicerPart, error) {
	packPart, relID, err := worksheetPart.addChildPart(
		uri,
		ContentTypeSlicer,
		RelationshipTypeSlicer,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlicer,
		packPart,
		worksheetPart,
	)
	partData.SetRelationshipID(relID)

	sp := &SlicerPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal slicer content
	sp.initializeContent()

	// Add to worksheet part's child parts
	if err := worksheetPart.AddPart(sp, relID); err != nil {
		return nil, err
	}

	return sp, nil
}

// initializeContent sets up minimal slicer content.
func (sp *SlicerPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<slicers xmlns="http://schemas.microsoft.com/office/spreadsheetml/2009/9/main" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" mc:Ignorable="x">
</slicers>`
	sp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*SlicerPart) FixedContentType() string {
	return ContentTypeSlicer
}

// Slicers returns the root Slicers element.
func (sp *SlicerPart) Slicers() *elements.Slicers {
	if root := sp.RootElement(); root != nil {
		if s, ok := root.(*elements.Slicers); ok {
			return s
		}
	}

	return nil
}

// Counter for generating unique slicer cache filenames.
var slicerCacheCounter uint64

// AddSlicerCachePart adds a slicer cache part.
func (sp *SlicerPart) AddSlicerCachePart() (*SlicerCachePart, error) {
	num := atomic.AddUint64(
		&slicerCacheCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/xl/slicerCaches/slicerCache%d.xml",
		num,
	)

	return newSlicerCachePart(sp, uri)
}

// SlicerCachePart returns the slicer cache part if present.
func (sp *SlicerPart) SlicerCachePart() *SlicerCachePart {
	for part := range sp.Parts() {
		if scp, ok := part.(*SlicerCachePart); ok {
			return scp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (sp *SlicerPart) GetStream() io.Reader {
	return sp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (sp *SlicerPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		sp,
		uri,
		contentType,
		relType,
	)
}

// Ensure SlicerPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*SlicerPart)(nil)

// SlicerCachePart represents a slicer cache part (xl/slicerCaches/slicerCache1.xml, etc.).
// The slicer cache stores the cached data and selection state for a slicer.
type SlicerCachePart struct {
	*openxml.OpenXmlPartData
}

// newSlicerCachePart creates a new slicer cache part.
func newSlicerCachePart(
	slicerPart *SlicerPart,
	uri string,
) (*SlicerCachePart, error) {
	packPart, relID, err := slicerPart.addChildPart(
		uri,
		ContentTypeSlicerCache,
		RelationshipTypeSlicerCache,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlicerCache,
		packPart,
		slicerPart,
	)
	partData.SetRelationshipID(relID)

	scp := &SlicerCachePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal slicer cache content
	scp.initializeContent()

	// Add to slicer part's child parts
	if err := slicerPart.AddPart(scp, relID); err != nil {
		return nil, err
	}

	return scp, nil
}

// initializeContent sets up minimal slicer cache content.
func (scp *SlicerCachePart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<slicerCacheDefinition xmlns="http://schemas.microsoft.com/office/spreadsheetml/2009/9/main" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" mc:Ignorable="x" name="Slicer_Field1">
</slicerCacheDefinition>`
	scp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*SlicerCachePart) FixedContentType() string {
	return ContentTypeSlicerCache
}

// SlicerCacheDefinition returns the root SlicerCacheDefinition element.
func (scp *SlicerCachePart) SlicerCacheDefinition() *elements.SlicerCacheDefinition {
	if root := scp.RootElement(); root != nil {
		if scd, ok := root.(*elements.SlicerCacheDefinition); ok {
			return scd
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (scp *SlicerCachePart) GetStream() io.Reader {
	return scp.OpenXmlPartData.GetStream()
}

// Ensure SlicerCachePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*SlicerCachePart)(
	nil,
)

// SlicerPartFactory creates a SlicerPart from a URI and container.
func SlicerPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	pkg := container.Package()
	if pkg == nil {
		return nil
	}

	packPart, err := pkg.Part(uri)
	if err != nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlicer,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewSlicers()
		},
	)

	return &SlicerPart{
		OpenXmlPartData: partData,
	}
}

// SlicerCachePartFactory creates a SlicerCachePart from a URI and container.
func SlicerCachePartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	pkg := container.Package()
	if pkg == nil {
		return nil
	}

	packPart, err := pkg.Part(uri)
	if err != nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlicerCache,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewSlicerCacheDefinition()
		},
	)

	return &SlicerCachePart{
		OpenXmlPartData: partData,
	}
}

// Register the slicer part types.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeSlicer,
			RelationshipType:   RelationshipTypeSlicer,
			Factory:            SlicerPartFactory,
			DefaultURI:         "/xl/slicers/slicer1.xml",
			IsFixedContentType: true,
		},
	)

	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeSlicerCache,
			RelationshipType:   RelationshipTypeSlicerCache,
			Factory:            SlicerCachePartFactory,
			DefaultURI:         "/xl/slicerCaches/slicerCache1.xml",
			IsFixedContentType: true,
		},
	)
}
