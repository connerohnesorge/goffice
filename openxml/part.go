// Package openxml provides the core framework for Office Open XML document processing.
package openxml

import (
	"bytes"
	"io"
	"iter"
	"reflect"
	"sync"

	"github.com/connerohnesorge/goffice/openxml/features"
	"github.com/connerohnesorge/goffice/packaging"
)

// OpenXmlPartContainer is the interface for types that can contain parts.
// Both OpenXmlPackage and OpenXmlPart can be containers for child parts.
type OpenXmlPartContainer interface {
	// Parts returns an iterator over all child parts.
	Parts() iter.Seq[OpenXmlPart]

	// GetPartById returns the child part with the given relationship ID.
	GetPartById(id string) (OpenXmlPart, error)

	// GetPartsOfType returns an iterator over all child parts of type T.
	// This uses generics to filter parts by their concrete type.
	GetPartsOfType(
		contentType string,
	) iter.Seq[OpenXmlPart]

	// AddPart adds a part as a child with the given relationship ID.
	// If id is empty, a unique ID will be generated.
	AddPart(part OpenXmlPart, id string) error

	// DeletePart removes the child part with the given relationship ID.
	DeletePart(id string) error

	// Features returns the feature collection for this container.
	Features() *features.FeatureCollection

	// URI returns the URI of this container.
	URI() string

	// Package returns the underlying OPC Package.
	Package() *packaging.Package

	// GetPackagingPart returns the underlying packaging.Part by its URI.
	// This method is safe to call during part loading and doesn't acquire locks.
	GetPackagingPart(uri string) *packaging.Part
}

// OpenXmlPartData holds the actual part data for an OpenXmlPart.
// This extends the minimal OpenXmlPart interface defined in element.go.
type OpenXmlPartData struct {
	mu sync.RWMutex

	// Identity
	uri            string
	contentType    string
	relationshipID string

	// Underlying OPC part
	packagingPart *packaging.Part

	// Parent container (package or another part)
	container OpenXmlPartContainer

	// Features inherited from container
	features *features.FeatureCollection

	// Root element (lazy loaded)
	rootElement PartRootElement
	rootLoaded  bool
	rootFactory func() PartRootElement

	// Child parts (relationship ID -> part)
	childParts map[string]OpenXmlPart

	// ID generator for child relationships
	idGenerator *RelationshipIDGenerator

	// Track modifications
	isDirty bool
}

// NewOpenXmlPartData creates a new part data structure.
func NewOpenXmlPartData(
	uri, contentType string,
	packPart *packaging.Part,
	container OpenXmlPartContainer,
) *OpenXmlPartData {
	part := &OpenXmlPartData{
		uri:           uri,
		contentType:   contentType,
		packagingPart: packPart,
		container:     container,
		childParts: make(
			map[string]OpenXmlPart,
		),
		idGenerator: NewRelationshipIDGenerator(),
		isDirty:     false,
	}

	// Create features with parent from container
	if container != nil {
		part.features = features.NewFeatureCollectionWithParent(
			container.Features(),
		)
	} else {
		part.features = features.NewFeatureCollection()
	}

	return part
}

// URI returns the part URI.
func (p *OpenXmlPartData) URI() string {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.uri
}

// ContentType returns the content type.
func (p *OpenXmlPartData) ContentType() string {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.contentType
}

// RelationshipID returns the relationship ID within the parent container.
func (p *OpenXmlPartData) RelationshipID() string {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.relationshipID
}

// SetRelationshipID sets the relationship ID.
func (p *OpenXmlPartData) SetRelationshipID(
	id string,
) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.relationshipID = id
}

// Container returns the parent container.
func (p *OpenXmlPartData) Container() OpenXmlPartContainer {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.container
}

// Features returns the feature collection.
func (p *OpenXmlPartData) Features() *features.FeatureCollection {
	return p.features
}

// GetStream returns a reader for the part content.
func (p *OpenXmlPartData) GetStream() io.Reader {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.packagingPart != nil {
		return p.packagingPart.GetStream()
	}

	return bytes.NewReader(nil)
}

// SetData sets the part's raw content.
func (p *OpenXmlPartData) SetData(data []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.packagingPart != nil {
		p.packagingPart.SetData(data)
	}
	p.isDirty = true
}

// GetData returns the raw part content.
func (p *OpenXmlPartData) GetData() []byte {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.packagingPart != nil {
		return p.packagingPart.GetData()
	}

	return nil
}

// RootElement returns the root element of this part.
// The element is lazy-loaded on first access.
func (p *OpenXmlPartData) RootElement() PartRootElement {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.rootLoaded && p.rootFactory != nil {
		p.rootElement = p.rootFactory()
		if p.rootElement != nil {
			p.rootElement.SetPart(p)
		}
		p.rootLoaded = true
	}

	return p.rootElement
}

// SetRootElement sets the root element.
func (p *OpenXmlPartData) SetRootElement(
	root PartRootElement,
) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.rootElement = root
	p.rootLoaded = true
	if root != nil {
		root.SetPart(p)
	}
	p.isDirty = true
}

// SetRootFactory sets the factory function for lazy loading the root element.
func (p *OpenXmlPartData) SetRootFactory(
	factory func() PartRootElement,
) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.rootFactory = factory
	p.rootLoaded = false
}

// IsDirty returns true if the part has been modified.
func (p *OpenXmlPartData) IsDirty() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.isDirty
}

// MarkDirty marks the part as modified.
func (p *OpenXmlPartData) MarkDirty() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.isDirty = true
}

// ClearDirty clears the dirty flag.
func (p *OpenXmlPartData) ClearDirty() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.isDirty = false
}

// Save serializes the root element to the part stream.
func (p *OpenXmlPartData) Save() error {
	p.mu.Lock()
	root := p.rootElement
	p.mu.Unlock()

	if root == nil {
		return nil
	}

	if err := root.Save(); err != nil {
		return err
	}

	p.ClearDirty()

	return nil
}

// Reload reloads the root element from the part stream.
func (p *OpenXmlPartData) Reload() error {
	p.mu.Lock()
	root := p.rootElement
	p.mu.Unlock()

	if root == nil {
		return nil
	}

	return root.Reload()
}

// Package returns the underlying OPC package.
func (p *OpenXmlPartData) Package() *packaging.Package {
	p.mu.RLock()
	container := p.container
	p.mu.RUnlock()

	if container != nil {
		return container.Package()
	}

	return nil
}

// GetPackagingPart returns the underlying packaging.Part by its URI.
// This delegates to the container which holds the package.
func (p *OpenXmlPartData) GetPackagingPart(
	uri string,
) *packaging.Part {
	p.mu.RLock()
	container := p.container
	p.mu.RUnlock()

	if container != nil {
		return container.GetPackagingPart(uri)
	}

	return nil
}

// Parts returns an iterator over child parts.
func (p *OpenXmlPartData) Parts() iter.Seq[OpenXmlPart] {
	return func(yield func(OpenXmlPart) bool) {
		p.mu.RLock()
		defer p.mu.RUnlock()

		for _, part := range p.childParts {
			if !yield(part) {
				return
			}
		}
	}
}

// GetPartById returns a child part by relationship ID.
func (p *OpenXmlPartData) GetPartById(
	id string,
) (OpenXmlPart, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	part, ok := p.childParts[id]
	if !ok {
		return nil, ErrPartNotFound
	}

	return part, nil
}

// GetPartsOfType returns an iterator over child parts of a specific content type.
func (p *OpenXmlPartData) GetPartsOfType(
	contentType string,
) iter.Seq[OpenXmlPart] {
	return func(yield func(OpenXmlPart) bool) {
		p.mu.RLock()
		defer p.mu.RUnlock()

		for _, part := range p.childParts {
			if part.ContentType() == contentType {
				if !yield(part) {
					return
				}
			}
		}
	}
}

// AddPart adds a child part with the given relationship ID.
func (p *OpenXmlPartData) AddPart(
	part OpenXmlPart,
	id string,
) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if id == "" {
		id = p.idGenerator.Next()
	} else {
		p.idGenerator.Reserve(id)
	}

	// Set the relationship ID on the part
	if partData, ok := part.(*OpenXmlPartData); ok {
		partData.SetRelationshipID(id)
	}

	p.childParts[id] = part
	p.isDirty = true

	return nil
}

// DeletePart removes a child part by relationship ID.
func (p *OpenXmlPartData) DeletePart(
	id string,
) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.childParts[id]; !ok {
		return ErrPartNotFound
	}

	delete(p.childParts, id)
	p.isDirty = true

	return nil
}

// PackagingPart returns the underlying OPC part.
func (p *OpenXmlPartData) PackagingPart() *packaging.Part {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.packagingPart
}

// Part type registry

// PartTypeInfo holds metadata about a part type.
type PartTypeInfo struct {
	// ContentType is the MIME content type for this part type.
	ContentType string

	// RelationshipType is the relationship type used to identify this part.
	RelationshipType string

	// PartType is the reflect.Type of the concrete part implementation.
	PartType reflect.Type

	// RootElementType is the reflect.Type of the root element (if any).
	RootElementType reflect.Type

	// Factory creates a new instance of this part type.
	Factory func(uri string, container OpenXmlPartContainer) OpenXmlPart

	// DefaultURI is the default URI pattern for this part type.
	DefaultURI string

	// TargetPath is the default target path in relationships.
	TargetPath string

	// IsFixedContentType indicates if this part has a fixed content type.
	IsFixedContentType bool
}

// partTypeRegistry holds registered part types.
var partTypeRegistry = struct {
	mu             sync.RWMutex
	byContentType  map[string]*PartTypeInfo
	byRelationship map[string]*PartTypeInfo
}{
	byContentType: make(
		map[string]*PartTypeInfo,
	),
	byRelationship: make(
		map[string]*PartTypeInfo,
	),
}

// RegisterPartType registers a part type for automatic instantiation.
func RegisterPartType(info *PartTypeInfo) {
	partTypeRegistry.mu.Lock()
	defer partTypeRegistry.mu.Unlock()

	if info.ContentType != "" {
		partTypeRegistry.byContentType[info.ContentType] = info
	}
	if info.RelationshipType != "" {
		partTypeRegistry.byRelationship[info.RelationshipType] = info
	}
}

// GetPartTypeByContentType returns the part type info for a content type.
func GetPartTypeByContentType(
	contentType string,
) (*PartTypeInfo, bool) {
	partTypeRegistry.mu.RLock()
	defer partTypeRegistry.mu.RUnlock()

	info, ok := partTypeRegistry.byContentType[contentType]

	return info, ok
}

// GetPartTypeByRelationship returns the part type info for a relationship type.
func GetPartTypeByRelationship(
	relType string,
) (*PartTypeInfo, bool) {
	partTypeRegistry.mu.RLock()
	defer partTypeRegistry.mu.RUnlock()

	info, ok := partTypeRegistry.byRelationship[relType]

	return info, ok
}

// CreatePartByContentType creates a part instance based on content type.
// Returns a generic OpenXmlPartData if no registered type matches.
func CreatePartByContentType(
	contentType, uri string,
	packPart *packaging.Part,
	container OpenXmlPartContainer,
) OpenXmlPart {
	info, ok := GetPartTypeByContentType(
		contentType,
	)
	if ok && info.Factory != nil {
		return info.Factory(uri, container)
	}

	// Return generic part
	return NewOpenXmlPartData(
		uri,
		contentType,
		packPart,
		container,
	)
}

// CreatePartByRelationship creates a part instance based on relationship type.
// Returns a generic OpenXmlPartData if no registered type matches.
func CreatePartByRelationship(
	relType, uri string,
	packPart *packaging.Part,
	container OpenXmlPartContainer,
) OpenXmlPart {
	info, ok := GetPartTypeByRelationship(relType)
	if ok && info.Factory != nil {
		return info.Factory(uri, container)
	}

	// Return generic part
	contentType := ""
	if packPart != nil {
		contentType = packPart.ContentType()
	}

	return NewOpenXmlPartData(
		uri,
		contentType,
		packPart,
		container,
	)
}

// GetPartsOfType is a generic function to iterate over parts of a specific type.
// Usage: GetPartsOfType[*StylesPart](container)
func GetPartsOfType[T OpenXmlPart](
	container OpenXmlPartContainer,
) iter.Seq[T] {
	return func(yield func(T) bool) {
		for part := range container.Parts() {
			if typed, ok := part.(T); ok {
				if !yield(typed) {
					return
				}
			}
		}
	}
}

// FirstPartOfType returns the first child part of type T.
// Returns nil if no matching part is found.
func FirstPartOfType[T OpenXmlPart](
	container OpenXmlPartContainer,
) T {
	for part := range GetPartsOfType[T](container) {
		return part
	}
	var zero T

	return zero
}

// IFixedContentTypePart is implemented by parts that have a fixed content type.
type IFixedContentTypePart interface {
	OpenXmlPart
	// FixedContentType returns the fixed content type for this part.
	FixedContentType() string
}

// Ensure OpenXmlPartData implements OpenXmlPart.
var _ OpenXmlPart = (*OpenXmlPartData)(nil)

// Part errors
var (
	ErrPartNotFound = partError(
		"part not found",
	)
	ErrPartExists = partError(
		"part already exists",
	)
	ErrInvalidPartURI = partError(
		"invalid part URI",
	)
	ErrContentTypeMismatch = partError(
		"content type mismatch",
	)
)

type partError string

func (e partError) Error() string {
	return string(e)
}
