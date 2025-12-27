// Package features provides the feature collection infrastructure for Office
// Open XML. Features allow parts and elements to access shared services like
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
	Availability() OfficeVersion

	// Validators returns the schema validators for this element type.
	Validators() []Validator
}

// OfficeVersion represents which Office versions support an element.
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
