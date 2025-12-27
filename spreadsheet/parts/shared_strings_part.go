//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// SharedStringTablePart represents the shared strings part (xl/sharedStrings.xml).
// This part stores all unique strings used across the workbook for optimization.
type SharedStringTablePart struct {
	*openxml.OpenXmlPartData

	// sst is the cached shared string table element
	sst *elements.SharedStringTable
}

// newSharedStringTablePart creates a new shared string table part.
func newSharedStringTablePart(
	workbookPart *WorkbookPart,
) (*SharedStringTablePart, error) {
	uri := "/xl/sharedStrings.xml"

	packPart, relID, err := workbookPart.addChildPart(
		uri,
		ContentTypeSharedStrings,
		RelationshipTypeSharedStrings,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSharedStrings,
		packPart,
		workbookPart,
	)
	partData.SetRelationshipID(relID)

	ssp := &SharedStringTablePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal shared string table content
	ssp.initializeContent()

	// Add to workbook part's child parts
	if err := workbookPart.AddPart(ssp, relID); err != nil {
		return nil, err
	}

	return ssp, nil
}

// initializeContent sets up minimal shared string table content.
func (ssp *SharedStringTablePart) initializeContent() {
	// Create the shared string table element
	ssp.sst = elements.NewSharedStringTable()
	ssp.sst.SetCount(0)
	ssp.sst.SetUniqueCount(0)

	// Set the namespace attribute
	ssp.sst.SetAttribute(
		openxml.NewAttribute(
			"http://www.w3.org/2000/xmlns/",
			"xmlns",
			"",
			elements.NamespaceSML,
		),
	)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*SharedStringTablePart) FixedContentType() string {
	return ContentTypeSharedStrings
}

// SharedStringTable returns the root SharedStringTable element.
func (ssp *SharedStringTablePart) SharedStringTable() *elements.SharedStringTable {
	if ssp.sst == nil {
		// Create a new shared string table
		ssp.sst = elements.NewSharedStringTable()
	}

	return ssp.sst
}

// GetString returns the string at the specified index.
// Returns an empty string if the index is out of range.
func (ssp *SharedStringTablePart) GetString(
	index int,
) string {
	return ssp.SharedStringTable().
		GetString(index)
}

// AddString adds a string to the shared string table and returns its index.
// If the string already exists, returns the existing index (deduplication).
func (ssp *SharedStringTablePart) AddString(
	text string,
) int {
	return ssp.SharedStringTable().AddString(text)
}

// IndexOf returns the index of the specified string in the table.
// Returns -1 if the string is not found.
func (ssp *SharedStringTablePart) IndexOf(
	text string,
) int {
	return ssp.SharedStringTable().IndexOf(text)
}

// ItemCount returns the number of unique strings in the table.
func (ssp *SharedStringTablePart) ItemCount() int {
	return ssp.SharedStringTable().ItemCount()
}

// GetItem returns the SharedStringItem at the specified index.
func (ssp *SharedStringTablePart) GetItem(
	index int,
) *elements.SharedStringItem {
	return ssp.SharedStringTable().GetItem(index)
}

// GetStream returns a reader for the part content.
func (ssp *SharedStringTablePart) GetStream() io.Reader {
	return ssp.OpenXmlPartData.GetStream()
}

// Ensure SharedStringTablePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*SharedStringTablePart)(
	nil,
)

// SharedStringTablePartFactory creates a SharedStringTablePart from a URI and container.
func SharedStringTablePartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSharedStrings,
		packPart,
		container,
	)

	return &SharedStringTablePart{
		OpenXmlPartData: partData,
	}
}

// Register the SharedStringTablePart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeSharedStrings,
			RelationshipType:   RelationshipTypeSharedStrings,
			Factory:            SharedStringTablePartFactory,
			DefaultURI:         "/xl/sharedStrings.xml",
			IsFixedContentType: true,
		},
	)
}
