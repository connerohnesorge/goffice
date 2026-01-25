// relationships, namespaces, and document-level settings.
package features //nolint:revive // max-public-structs: interfaces.go defines many public types by design

import (
	"io"
)

// Feature is a marker interface for all features that can be stored in a
// FeatureCollection. All specific feature interfaces embed this interface.
type Feature interface {
	// featureMarker is unexported to ensure only this package can
	// implement Feature.
	featureMarker()
}

// featureBase is embedded in all feature implementations (unexported).
type featureBase struct{}

func (featureBase) featureMarker() {}

// FeatureBase is the exported base type for feature implementations.
// Embed this in your feature implementation to satisfy the Feature interface.
type FeatureBase struct {
	featureBase
}

// IPackageFeature provides access to the underlying OPC Package.
type IPackageFeature interface {
	Feature

	// Package returns the underlying packaging.Package.
	// Returns nil if not associated with a package.
	Package() any

	// Capabilities returns the package access capabilities.
	Capabilities() PackageCapabilities
}

// PackageCapabilities describes what operations are allowed on the package.
type PackageCapabilities struct {
	CanRead  bool
	CanWrite bool
	CanSave  bool
}

// IContentTypeFeature provides content type resolution for parts.
type IContentTypeFeature interface {
	Feature

	// GetContentType returns the content type for the given part URI.
	GetContentType(uri string) (string, error)

	// SetContentType sets the content type for the given part URI.
	SetContentType(uri, contentType string) error

	// RemoveContentType removes the content type override for the given URI.
	RemoveContentType(uri string)
}

// INamespaceFeature provides namespace resolution services.
type INamespaceFeature interface {
	Feature

	// ResolveNamespace returns the namespace URI for the given prefix.
	ResolveNamespace(prefix string) (string, bool)

	// ResolvePrefix returns the preferred prefix for the given namespace URI.
	ResolvePrefix(
		namespaceURI string,
	) (string, bool)

	// RegisterNamespace registers a prefix-to-namespace mapping.
	RegisterNamespace(prefix, namespaceURI string)
}

// IPartRelationshipsFeature provides access to a part's child relationships.
type IPartRelationshipsFeature interface {
	Feature

	// CreateRelationship creates a new relationship from this part.
	CreateRelationship(
		target, relType, id string,
	) (any, error)

	// GetRelationship returns the relationship with the given ID.
	GetRelationship(id string) (any, error)

	// DeleteRelationship removes the relationship with the given ID.
	DeleteRelationship(id string) error

	// GetRelationshipsByType returns all relationships of the given type.
	GetRelationshipsByType(relType string) []any
}

// IMainPartFeature identifies and provides access to the main document part.
type IMainPartFeature interface {
	Feature

	// MainPart returns the main document part.
	MainPart() any

	// ContentType returns the expected content type of the main part.
	ContentType() string

	// RelationshipType returns the relationship type used to identify
	// the main part.
	RelationshipType() string
}

// IPartRootFeature provides access to a part's root element.
type IPartRootFeature interface {
	Feature

	// Root returns the root element of the part.
	Root() any

	// SetRoot sets the root element of the part.
	SetRoot(root any)
}

// IPartUriFeature provides part URI handling.
type IPartUriFeature interface {
	Feature

	// URI returns the part's URI.
	URI() string

	// GenerateChildUri generates a unique URI for a child part.
	GenerateChildUri(
		baseName, extension string,
	) string
}

// IElementMetadata provides schema metadata for elements.
type IElementMetadata interface {
	Feature

	// LocalName returns the element's local name.
	LocalName() string

	// NamespaceURI returns the element's namespace URI.
	NamespaceURI() string

	// Availability returns the Office versions that support this element.
	// This is a bitflag-based version set (deprecated in favor of AvailableInVersion).
	Availability() OfficeVersion

	// AvailableInVersion returns the first Office version where this element
	// was introduced. This provides more precise version tracking than the
	// bitflag-based Availability() method.
	AvailableInVersion() FileFormatVersion

	// Validators returns the schema validators for this element type.
	Validators() []Validator
}

// OfficeVersion represents which Office versions support an element.
// This is a bitflag-based version set (deprecated in favor of FileFormatVersion).
type OfficeVersion int

const (
	// Office2007 is the initial release (Office 2007 / ECMA-376 1st edition).
	Office2007 OfficeVersion = 1 << iota
	// Office2010 is Office 2010 (ISO/IEC 29500:2008).
	Office2010
	// Office2013 is Office 2013.
	Office2013
	// Office2016 is Office 2016.
	Office2016
	// Office2019 is Office 2019.
	Office2019
	// Office2021 is Office 2021.
	Office2021
	// OfficeAll represents all versions.
	OfficeAll = Office2007 | Office2010 | Office2013 |
		Office2016 | Office2019 | Office2021
)

// FileFormatVersion represents the Office version that introduced a feature.
// This is used to track which Office version first introduced an element,
// attribute, or other feature. It's a sequential enumeration for clear
// version ordering and comparison.
type FileFormatVersion int

const (
	// Office2007Format represents Microsoft Office 2007 (ECMA-376 1st edition).
	// This is the initial Open XML standard.
	Office2007Format FileFormatVersion = iota

	// Office2010Format represents Microsoft Office 2010 (ISO/IEC 29500:2008).
	// Introduces extensions like content controls, drawing canvas.
	Office2010Format

	// Office2013Format represents Microsoft Office 2013.
	// Introduces features like timeline slicers, webextensions.
	Office2013Format

	// Office2016Format represents Microsoft Office 2016.
	// Introduces modern animations, improved collaboration features.
	Office2016Format

	// Office2019Format represents Microsoft Office 2019.
	// Introduces new chart types, improved inking.
	Office2019Format

	// Office2021Format represents Microsoft Office 2021.
	// Introduces dynamic arrays (Excel), improved accessibility.
	Office2021Format

	// Office2022Format represents features specific to Office 2022.
	Office2022Format

	// Office2023Format represents features specific to Office 2023.
	Office2023Format

	// Office2024Format represents features specific to Office 2024.
	Office2024Format

	// Office2025Format represents features specific to Office 2025.
	Office2025Format

	// Microsoft365Format represents features exclusive to Microsoft 365
	// (formerly Office 365). This is a continuously updated version.
	Microsoft365Format
)

// String returns the human-readable name of the Office version.
func (v FileFormatVersion) String() string {
	switch v {
	case Office2007Format:
		return "Office 2007"
	case Office2010Format:
		return "Office 2010"
	case Office2013Format:
		return "Office 2013"
	case Office2016Format:
		return "Office 2016"
	case Office2019Format:
		return "Office 2019"
	case Office2021Format:
		return "Office 2021"
	case Office2022Format:
		return "Office 2022"
	case Office2023Format:
		return "Office 2023"
	case Office2024Format:
		return "Office 2024"
	case Office2025Format:
		return "Office 2025"
	case Microsoft365Format:
		return "Microsoft 365"
	default:
		return "Unknown"
	}
}

// Description returns a detailed description of the Office version.
func (v FileFormatVersion) Description() string {
	switch v {
	case Office2007Format:
		return "Microsoft Office 2007 (ECMA-376 1st Edition)"
	case Office2010Format:
		return "Microsoft Office 2010 (ISO/IEC 29500:2008)"
	case Office2013Format:
		return "Microsoft Office 2013"
	case Office2016Format:
		return "Microsoft Office 2016"
	case Office2019Format:
		return "Microsoft Office 2019"
	case Office2021Format:
		return "Microsoft Office 2021"
	case Office2022Format:
		return "Microsoft Office 2022"
	case Office2023Format:
		return "Microsoft Office 2023"
	case Office2024Format:
		return "Microsoft Office 2024"
	case Office2025Format:
		return "Microsoft Office 2025"
	case Microsoft365Format:
		return "Microsoft 365 (continuously updated)"
	default:
		return "Unknown Office Version"
	}
}

// Year returns the release year of the Office version (approximate).
// Returns 0 for Microsoft365Format as it's continuously updated.
//
//nolint:revive // Years are self-documenting literal values
func (v FileFormatVersion) Year() int {
	switch v {
	case Office2007Format:
		return 2007
	case Office2010Format:
		return 2010
	case Office2013Format:
		return 2013
	case Office2016Format:
		return 2016
	case Office2019Format:
		return 2019
	case Office2021Format:
		return 2021
	case Office2022Format:
		return 2022
	case Office2023Format:
		return 2023
	case Office2024Format:
		return 2024
	case Office2025Format:
		return 2025
	case Microsoft365Format:
		return 0 // Continuously updated, no fixed year
	default:
		return 0
	}
}

// AtLeast returns true if this version is at least the given version.
// Higher versions include features from lower versions.
func (v FileFormatVersion) AtLeast(
	other FileFormatVersion,
) bool {
	return v >= other
}

// AtMost returns true if this version is at most the given version.
func (v FileFormatVersion) AtMost(
	other FileFormatVersion,
) bool {
	return v <= other
}

// IsNewerThan returns true if this version is newer than the given version.
func (v FileFormatVersion) IsNewerThan(
	other FileFormatVersion,
) bool {
	return v > other
}

// IsOlderThan returns true if this version is older than the given version.
func (v FileFormatVersion) IsOlderThan(
	other FileFormatVersion,
) bool {
	return v < other
}

// AllFileFormatVersions returns all defined Office versions in chronological order.
var AllFileFormatVersions = []FileFormatVersion{
	Office2007Format,
	Office2010Format,
	Office2013Format,
	Office2016Format,
	Office2019Format,
	Office2021Format,
	Office2022Format,
	Office2023Format,
	Office2024Format,
	Office2025Format,
	Microsoft365Format,
}

// Validator is an interface for element validators.
type Validator interface {
	// Validate validates the element.
	Validate(element any) error
}

// IPartFeature provides access to the containing part.
type IPartFeature interface {
	Feature

	// Part returns the containing part.
	Part() any
}

// IStreamFeature provides access to a part's content stream.
type IStreamFeature interface {
	Feature

	// GetStream returns a reader for the part's content.
	GetStream() (io.Reader, error)

	// SetStream sets the part's content from a reader.
	SetStream(r io.Reader) error
}
