// Package openxml provides the core framework for Office Open XML document processing.
package openxml

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/packaging"
)

// TargetMode specifies how a relationship target should be interpreted.
type TargetMode int

const (
	// TargetModeInternal indicates the target is a part within the package.
	TargetModeInternal TargetMode = iota
	// TargetModeExternal indicates the target is an external resource.
	TargetModeExternal
)

// String returns the string representation of the TargetMode.
func (tm TargetMode) String() string {
	switch tm {
	case TargetModeExternal:
		return "External"
	default:
		return "Internal"
	}
}

// OpenXmlRelationship is the interface for all relationship types.
type OpenXmlRelationship interface {
	// ID returns the relationship ID (e.g., "rId1").
	ID() string

	// Type returns the relationship type URI.
	Type() string

	// Target returns the target URI.
	Target() string

	// TargetMode returns whether the target is internal or external.
	TargetMode() TargetMode

	// Container returns the part or package that owns this relationship.
	Container() OpenXmlPartContainer
}

// baseRelationship provides the common implementation for relationships.
type baseRelationship struct {
	id         string
	relType    string
	target     string
	targetMode TargetMode
	container  OpenXmlPartContainer
}

// ID returns the relationship ID.
func (r *baseRelationship) ID() string {
	return r.id
}

// Type returns the relationship type URI.
func (r *baseRelationship) Type() string {
	return r.relType
}

// Target returns the target URI.
func (r *baseRelationship) Target() string {
	return r.target
}

// TargetMode returns the target mode.
func (r *baseRelationship) TargetMode() TargetMode {
	return r.targetMode
}

// Container returns the owning container.
func (r *baseRelationship) Container() OpenXmlPartContainer {
	return r.container
}

// PartRelationship represents an internal relationship to another part.
type PartRelationship struct {
	baseRelationship
	targetPart OpenXmlPart
}

// NewPartRelationship creates a new relationship to an internal part.
func NewPartRelationship(
	id, relType string,
	target OpenXmlPart,
	container OpenXmlPartContainer,
) *PartRelationship {
	return &PartRelationship{
		baseRelationship: baseRelationship{
			id:         id,
			relType:    relType,
			target:     target.URI(),
			targetMode: TargetModeInternal,
			container:  container,
		},
		targetPart: target,
	}
}

// TargetPart returns the target part.
func (r *PartRelationship) TargetPart() OpenXmlPart {
	return r.targetPart
}

// ExternalRelationship represents an external URI relationship.
type ExternalRelationship struct {
	baseRelationship
}

// NewExternalRelationship creates a new relationship to an external URI.
func NewExternalRelationship(
	id, relType, targetURI string,
	container OpenXmlPartContainer,
) *ExternalRelationship {
	return &ExternalRelationship{
		baseRelationship: baseRelationship{
			id:         id,
			relType:    relType,
			target:     targetURI,
			targetMode: TargetModeExternal,
			container:  container,
		},
	}
}

// HyperlinkRelationship represents a hyperlink relationship.
type HyperlinkRelationship struct {
	baseRelationship
	isExternal bool
}

// NewHyperlinkRelationship creates a new hyperlink relationship.
func NewHyperlinkRelationship(
	id, targetURI string,
	isExternal bool,
	container OpenXmlPartContainer,
) *HyperlinkRelationship {
	mode := TargetModeInternal
	if isExternal {
		mode = TargetModeExternal
	}
	return &HyperlinkRelationship{
		baseRelationship: baseRelationship{
			id:         id,
			relType:    RelationshipTypeHyperlink,
			target:     targetURI,
			targetMode: mode,
			container:  container,
		},
		isExternal: isExternal,
	}
}

// IsExternal returns true if the hyperlink points to an external resource.
func (r *HyperlinkRelationship) IsExternal() bool {
	return r.isExternal
}

// DataPartReferenceRelationship represents a relationship to a data part.
type DataPartReferenceRelationship struct {
	baseRelationship
}

// NewDataPartReferenceRelationship creates a new data part reference relationship.
func NewDataPartReferenceRelationship(
	id, relType, target string,
	container OpenXmlPartContainer,
) *DataPartReferenceRelationship {
	return &DataPartReferenceRelationship{
		baseRelationship: baseRelationship{
			id:         id,
			relType:    relType,
			target:     target,
			targetMode: TargetModeInternal,
			container:  container,
		},
	}
}

// Standard relationship type constants.
const (
	// Core package relationships
	RelationshipTypePackageStart = "http://schemas.openxmlformats.org/package/2006/relationships/"

	// Thumbnail relationship type.
	RelationshipTypeThumbnail = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/thumbnail"

	// Digital signature relationship types.
	RelationshipTypeDigitalSignatureOrigin       = "http://schemas.openxmlformats.org/package/2006/relationships/digital-signature/origin"
	RelationshipTypeDigitalSignatureCertificate  = "http://schemas.openxmlformats.org/package/2006/relationships/digital-signature/certificate"
	RelationshipTypeDigitalSignatureSignature    = "http://schemas.openxmlformats.org/package/2006/relationships/digital-signature/signature"
	RelationshipTypeDigitalSignatureXpsSignature = "http://schemas.microsoft.com/xps/2005/06/signature-definitions"

	// Additional Word relationships
	RelationshipTypeCustomXml             = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/customXml"
	RelationshipTypeCustomXmlProps        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/customXmlProps"
	RelationshipTypeGlossaryDocument      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/glossaryDocument"
	RelationshipTypeDocument              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	RelationshipTypeAlternativeFormat     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/aFChunk"
	RelationshipTypeFramesetRelationship  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/frame"
	RelationshipTypeControl               = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/control"
	RelationshipTypeOleObject             = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/oleObject"
	RelationshipTypePackage               = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/package"
	RelationshipTypeEmbeddedPackage       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/embeddedPackage"
	RelationshipTypeVideo                 = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/video"
	RelationshipTypeAudio                 = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/audio"
	RelationshipTypeFont                  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/font"
	RelationshipTypeMailMergeHeaderSource = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/mailMergeSource"
	RelationshipTypeMailMergeDataSource   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/mailMergeSource"
	RelationshipTypeDiagramColors         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramColors"
	RelationshipTypeDiagramData           = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramData"
	RelationshipTypeDiagramLayout         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramLayout"
	RelationshipTypeDiagramQuickStyle     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramQuickStyle"
	RelationshipTypeDiagramDrawing        = "http://schemas.microsoft.com/office/2007/relationships/diagramDrawing"
	RelationshipTypeChart                 = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chart"
	RelationshipTypeChartSheet            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chartsheet"
	RelationshipTypeVmlDrawing            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/vmlDrawing"
	RelationshipTypeDrawing               = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/drawing"
	RelationshipTypeCustomization         = "http://schemas.microsoft.com/office/2006/relationships/ui/extensibility"
	RelationshipTypeRibbonExtensibility   = "http://schemas.microsoft.com/office/2007/relationships/ui/extensibility"
	RelationshipTypeQuickAccessToolbar    = "http://schemas.microsoft.com/office/2006/relationships/ui/userCustomization"
	RelationshipTypeAttachedToolbars      = "http://schemas.microsoft.com/office/2011/relationships/attachedToolbars"
	RelationshipTypePrinterSettings       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/printerSettings"
	RelationshipTypeExternalLink          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/externalLink"
	RelationshipTypeExternalLinkPath      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/externalLinkPath"
	RelationshipTypeTableStyle            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/tableStyles"
	RelationshipTypePivotTable            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/pivotTable"
	RelationshipTypePivotCacheDefinition  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/pivotCacheDefinition"
	RelationshipTypePivotCacheRecords     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/pivotCacheRecords"
	RelationshipTypeQueryTable            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/queryTable"
	RelationshipTypeConnectionTable       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/connections"
	RelationshipTypeSharedStrings         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/sharedStrings"
	RelationshipTypeCalculationChain      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/calcChain"
	RelationshipTypeWorksheet             = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet"
	RelationshipTypeSlide                 = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slide"
	RelationshipTypeSlideMaster           = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster"
	RelationshipTypeSlideLayout           = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideLayout"
	RelationshipTypeNotesMaster           = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesMaster"
	RelationshipTypeNotesSlide            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesSlide"
	RelationshipTypeHandoutMaster         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/handoutMaster"
	RelationshipTypePresProps             = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/presProps"
	RelationshipTypeViewProps             = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/viewProps"
	RelationshipTypeTableStyles           = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/tableStyles"
	RelationshipTypeCommentsExtended      = "http://schemas.microsoft.com/office/2011/relationships/commentsExtended"
	RelationshipTypeCommentsIDs           = "http://schemas.microsoft.com/office/2016/09/relationships/commentsIds"
	RelationshipTypePeople                = "http://schemas.microsoft.com/office/2011/relationships/people"
	RelationshipTypeTimeline              = "http://schemas.microsoft.com/office/2011/relationships/timeline"
	RelationshipTypeSlicer                = "http://schemas.microsoft.com/office/2011/relationships/slicer"
	RelationshipTypeSlicerCache           = "http://schemas.microsoft.com/office/2011/relationships/slicerCache"
	RelationshipTypeModel                 = "http://schemas.microsoft.com/office/2011/relationships/model"
	RelationshipTypeModel3D               = "http://schemas.microsoft.com/office/2017/06/relationships/model3d"
)

// RelationshipIDGenerator generates unique relationship IDs.
type RelationshipIDGenerator struct {
	mu     sync.Mutex
	nextID uint64
	used   map[string]bool
}

// NewRelationshipIDGenerator creates a new ID generator.
func NewRelationshipIDGenerator() *RelationshipIDGenerator {
	return &RelationshipIDGenerator{
		nextID: 1,
		used:   make(map[string]bool),
	}
}

// Reserve marks an ID as used (for loading existing relationships).
func (g *RelationshipIDGenerator) Reserve(
	id string,
) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.used[id] = true

	// Extract number if it follows rIdN pattern
	var num uint64
	if _, err := fmt.Sscanf(id, "rId%d", &num); err == nil {
		if num >= g.nextID {
			g.nextID = num + 1
		}
	}
}

// Next generates the next unique relationship ID.
func (g *RelationshipIDGenerator) Next() string {
	g.mu.Lock()
	defer g.mu.Unlock()

	for {
		id := fmt.Sprintf("rId%d", g.nextID)
		g.nextID++
		if !g.used[id] {
			g.used[id] = true
			return id
		}
	}
}

// globalIDCounter is used for generating unique IDs across packages.
var globalIDCounter uint64

// GenerateUniqueID generates a globally unique relationship ID.
func GenerateUniqueID() string {
	n := atomic.AddUint64(&globalIDCounter, 1)
	return fmt.Sprintf("rId%d", n)
}

// ResolveTargetURI resolves a relative target URI against a source part URI.
// This is used to get the absolute URI of a relationship target.
func ResolveTargetURI(
	sourceURI, relativeTarget string,
) string {
	return packaging.ResolvePartURI(
		sourceURI,
		relativeTarget,
	)
}

// RelativeTargetURI computes the relative URI from source to target.
// This is the inverse of ResolveTargetURI.
func RelativeTargetURI(
	sourceURI, targetURI string,
) string {
	// Get directories
	sourceDir := packaging.URIDirectory(sourceURI)
	targetDir := packaging.URIDirectory(targetURI)
	targetFile := packaging.URIFilename(targetURI)

	// If same directory, just return the filename
	if sourceDir == targetDir {
		return targetFile
	}

	// If target is in a subdirectory of source's directory
	if len(targetDir) > len(sourceDir) &&
		targetDir[:len(sourceDir)] == sourceDir {
		return targetDir[len(sourceDir)+1:] + "/" + targetFile
	}

	// Otherwise, use absolute path
	return targetURI
}

// RelationshipTypeInfo provides metadata about a relationship type.
type RelationshipTypeInfo struct {
	Type           string
	DefaultPartURI string
	ContentType    string
}

// knownRelationshipTypes maps relationship types to their default information.
var knownRelationshipTypes = map[string]RelationshipTypeInfo{
	RelationshipTypeOfficeDocument: {
		Type:           RelationshipTypeOfficeDocument,
		DefaultPartURI: "/word/document.xml",
		ContentType:    ContentTypeWordprocessingMLDocument,
	},
	RelationshipTypeStyles: {
		Type:           RelationshipTypeStyles,
		DefaultPartURI: "/word/styles.xml",
		ContentType:    ContentTypeStyles,
	},
	RelationshipTypeNumbering: {
		Type:           RelationshipTypeNumbering,
		DefaultPartURI: "/word/numbering.xml",
		ContentType:    ContentTypeNumbering,
	},
	RelationshipTypeSettings: {
		Type:           RelationshipTypeSettings,
		DefaultPartURI: "/word/settings.xml",
		ContentType:    ContentTypeSettings,
	},
	RelationshipTypeFontTable: {
		Type:           RelationshipTypeFontTable,
		DefaultPartURI: "/word/fontTable.xml",
		ContentType:    ContentTypeFontTable,
	},
	RelationshipTypeWebSettings: {
		Type:           RelationshipTypeWebSettings,
		DefaultPartURI: "/word/webSettings.xml",
		ContentType:    ContentTypeWebSettings,
	},
	RelationshipTypeTheme: {
		Type:           RelationshipTypeTheme,
		DefaultPartURI: "/word/theme/theme1.xml",
		ContentType:    ContentTypeTheme,
	},
	RelationshipTypeCoreProperties: {
		Type:           RelationshipTypeCoreProperties,
		DefaultPartURI: "/docProps/core.xml",
		ContentType:    ContentTypeCoreProperties,
	},
	RelationshipTypeExtendedProperties: {
		Type:           RelationshipTypeExtendedProperties,
		DefaultPartURI: "/docProps/app.xml",
		ContentType:    ContentTypeExtendedProperties,
	},
}

// GetRelationshipTypeInfo returns metadata about a relationship type.
func GetRelationshipTypeInfo(
	relType string,
) (RelationshipTypeInfo, bool) {
	info, ok := knownRelationshipTypes[relType]
	return info, ok
}

// RegisterRelationshipType registers a new relationship type with its metadata.
func RegisterRelationshipType(
	info RelationshipTypeInfo,
) {
	knownRelationshipTypes[info.Type] = info
}
