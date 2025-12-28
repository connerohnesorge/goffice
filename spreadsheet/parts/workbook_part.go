//nolint:revive // line-length-limit: OOXML content types are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// WorkbookPart represents the main workbook part (xl/workbook.xml).
// This is the primary part containing the workbook definition and sheet references.
type WorkbookPart struct {
	*openxml.OpenXmlPartData

	// contentType is the specific content type for this workbook type
	contentType string
}

// NewWorkbookPart creates a new workbook part.
func NewWorkbookPart(
	uri, contentType string,
	container openxml.OpenXmlPartContainer,
) (*WorkbookPart, error) {
	// Create the underlying packaging part
	pkg := container.Package()
	if pkg == nil {
		return nil, ErrNilPackage
	}

	packPart, err := pkg.CreatePart(
		uri,
		contentType,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		contentType,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewWorkbook()
		},
	)

	wp := &WorkbookPart{
		OpenXmlPartData: partData,
		contentType:     contentType,
	}

	// Create package-level relationship
	_, err = pkg.CreateRelationship(
		uri,
		RelationshipTypeOfficeDocument,
		"",
	)
	if err != nil {
		return nil, err
	}

	return wp, nil
}

// NewWorkbookPartFromData wraps an existing OpenXmlPartData as a WorkbookPart.
func NewWorkbookPartFromData(
	data *openxml.OpenXmlPartData,
	contentType string,
) *WorkbookPart {
	return &WorkbookPart{
		OpenXmlPartData: data,
		contentType:     contentType,
	}
}

// FixedContentType returns the content type for this part.
func (wp *WorkbookPart) FixedContentType() string {
	return wp.contentType
}

// InitializeContent sets up minimal workbook content for a new workbook.
func (wp *WorkbookPart) InitializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<workbook xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
  <sheets/>
</workbook>`
	wp.SetData([]byte(content))
}

// Workbook returns the root Workbook element.
// This ensures the element is loaded from XML if it hasn't been already.
func (wp *WorkbookPart) Workbook() openxml.PartRootElement {
	root := wp.RootElement()
	if root != nil {
		// Ensure the element is loaded from XML by checking if it has been populated
		// If it's a composite element with no children and we have data, reload it
		if comp, ok := root.(openxml.CompositeElement); ok {
			hasChildren := false
			for range comp.Children() {
				hasChildren = true

				break
			}
			if !hasChildren &&
				wp.GetData() != nil {
				// Element hasn't been loaded yet, reload it
				_ = root.Reload()
			}
		}
	}

	return root
}

// Counter for generating unique worksheet filenames.
var worksheetCounter uint64

// AddWorksheetPart adds a new worksheet part to this workbook.
func (wp *WorkbookPart) AddWorksheetPart() (*WorksheetPart, error) {
	// Find an available worksheet URI that doesn't already exist
	pkg := wp.Package()
	if pkg == nil {
		return nil, ErrNilPackage
	}

	var uri string
	for range 1000 { // Safety limit
		num := atomic.AddUint64(
			&worksheetCounter,
			1,
		)
		uri = fmt.Sprintf(
			"/xl/worksheets/sheet%d.xml",
			num,
		)
		// Check if this part already exists
		if _, err := pkg.Part(uri); err != nil {
			// Part doesn't exist, we can use this URI
			break
		}
	}

	return newWorksheetPart(wp, uri)
}

// WorksheetParts returns all worksheet parts.
func (wp *WorkbookPart) WorksheetParts() []*WorksheetPart {
	var worksheets []*WorksheetPart
	for part := range wp.Parts() {
		if wsp, ok := part.(*WorksheetPart); ok {
			worksheets = append(worksheets, wsp)
		}
	}

	return worksheets
}

// AddSharedStringTablePart adds a shared string table part to this workbook.
func (wp *WorkbookPart) AddSharedStringTablePart() (*SharedStringTablePart, error) {
	return newSharedStringTablePart(wp)
}

// SharedStringTablePart returns the shared string table part if present.
func (wp *WorkbookPart) SharedStringTablePart() *SharedStringTablePart {
	for part := range wp.Parts() {
		if ssp, ok := part.(*SharedStringTablePart); ok {
			return ssp
		}
	}

	return nil
}

// AddStylesPart adds a workbook styles part to this workbook.
func (wp *WorkbookPart) AddStylesPart() (*WorkbookStylesPart, error) {
	return newWorkbookStylesPart(wp)
}

// StylesPart returns the workbook styles part if present.
func (wp *WorkbookPart) StylesPart() *WorkbookStylesPart {
	for part := range wp.Parts() {
		if sp, ok := part.(*WorkbookStylesPart); ok {
			return sp
		}
	}

	return nil
}

// AddCalculationChainPart adds a calculation chain part to this workbook.
func (wp *WorkbookPart) AddCalculationChainPart() (*CalculationChainPart, error) {
	return newCalculationChainPart(wp)
}

// CalculationChainPart returns the calculation chain part if present.
func (wp *WorkbookPart) CalculationChainPart() *CalculationChainPart {
	for part := range wp.Parts() {
		if ccp, ok := part.(*CalculationChainPart); ok {
			return ccp
		}
	}

	return nil
}

// AddConnectionsPart adds a connections part to this workbook.
func (wp *WorkbookPart) AddConnectionsPart() (*ConnectionsPart, error) {
	return newConnectionsPart(wp)
}

// ConnectionsPart returns the connections part if present.
func (wp *WorkbookPart) ConnectionsPart() *ConnectionsPart {
	for part := range wp.Parts() {
		if cp, ok := part.(*ConnectionsPart); ok {
			return cp
		}
	}

	return nil
}

// AddThemePart adds a theme part to this workbook.
func (wp *WorkbookPart) AddThemePart() (*ThemePart, error) {
	return newThemePart(wp)
}

// ThemePart returns the theme part if present.
func (wp *WorkbookPart) ThemePart() *ThemePart {
	for part := range wp.Parts() {
		if tp, ok := part.(*ThemePart); ok {
			return tp
		}
	}

	return nil
}

// Counter for generating unique chartsheet filenames.
var chartsheetCounter uint64

// AddChartsheetPart adds a new chartsheet part to this workbook.
func (wp *WorkbookPart) AddChartsheetPart() (*ChartsheetPart, error) {
	num := atomic.AddUint64(&chartsheetCounter, 1)
	uri := fmt.Sprintf(
		"/xl/chartsheets/sheet%d.xml",
		num,
	)

	return newChartsheetPart(wp, uri)
}

// ChartsheetParts returns all chartsheet parts.
func (wp *WorkbookPart) ChartsheetParts() []*ChartsheetPart {
	var chartsheets []*ChartsheetPart
	for part := range wp.Parts() {
		if csp, ok := part.(*ChartsheetPart); ok {
			chartsheets = append(chartsheets, csp)
		}
	}

	return chartsheets
}

// AddExternalWorkbookPart adds an external workbook reference part.
func (wp *WorkbookPart) AddExternalWorkbookPart() (*ExternalWorkbookPart, error) {
	return newExternalWorkbookPart(wp)
}

// ExternalWorkbookParts returns all external workbook parts.
func (wp *WorkbookPart) ExternalWorkbookParts() []*ExternalWorkbookPart {
	var externals []*ExternalWorkbookPart
	for part := range wp.Parts() {
		if ewp, ok := part.(*ExternalWorkbookPart); ok {
			externals = append(externals, ewp)
		}
	}

	return externals
}

// AddVbaProjectPart adds a VBA project part to this workbook.
func (wp *WorkbookPart) AddVbaProjectPart() (*VbaProjectPart, error) {
	return newVbaProjectPart(wp)
}

// VbaProjectPart returns the VBA project part if present.
func (wp *WorkbookPart) VbaProjectPart() *VbaProjectPart {
	for part := range wp.Parts() {
		if vp, ok := part.(*VbaProjectPart); ok {
			return vp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (wp *WorkbookPart) GetStream() io.Reader {
	return wp.OpenXmlPartData.GetStream()
}

// Ensure WorkbookPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*WorkbookPart)(nil)

// WorkbookPartFactory creates a WorkbookPart from a URI and container.
func WorkbookPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Get the packaging part from the container directly
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		packPart.ContentType(),
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewWorkbook()
		},
	)

	wp := NewWorkbookPartFromData(
		partData,
		packPart.ContentType(),
	)

	return wp
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (wp *WorkbookPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		wp,
		uri,
		contentType,
		relType,
	)
}

// Register the WorkbookPart type.
func init() {
	// Register for all Excel workbook content types
	contentTypes := []string{
		ContentTypeWorkbook,
		ContentTypeWorkbookTemplate,
		ContentTypeWorkbookMacroEnabled,
		ContentTypeMacroTemplate,
		ContentTypeAddIn,
	}

	for _, ct := range contentTypes {
		openxml.RegisterPartType(
			&openxml.PartTypeInfo{
				ContentType:        ct,
				RelationshipType:   RelationshipTypeOfficeDocument,
				Factory:            WorkbookPartFactory,
				DefaultURI:         "/xl/workbook.xml",
				IsFixedContentType: false, // Content type varies by workbook type
			},
		)

		// Also register PURL namespace variant for compatibility
		openxml.RegisterPartType(
			&openxml.PartTypeInfo{
				ContentType:        ct,
				RelationshipType:   RelationshipTypePURLOfficeDocument,
				Factory:            WorkbookPartFactory,
				DefaultURI:         "/xl/workbook.xml",
				IsFixedContentType: false, // Content type varies by workbook type
			},
		)
	}
}
