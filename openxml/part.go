// document processing.
//
//nolint:revive // file-length-limit: this file contains the core part types
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
	// This method is safe to call during part loading and doesn't
	// acquire locks.
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

	// Track if we're currently loading children to prevent infinite recursion
	loadingChildren bool
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

// loadChildParts loads child parts from the underlying packaging part relationships.
func (p *OpenXmlPartData) loadChildParts() {
	p.loadChildPartsRecursive(0, nil)
}

const maxPartRecursionDepth = 10

// loadChildPartsRecursive loads child parts with a recursion depth limit.
func (p *OpenXmlPartData) loadChildPartsRecursive(
	depth int,
	loadedURIs map[string]bool,
) {
	if depth > maxPartRecursionDepth {
		return
	}

	// Initialize loadedURIs map on first call
	if loadedURIs == nil {
		loadedURIs = make(map[string]bool)
	}

	// Check if we're already loading children to prevent infinite recursion
	p.mu.Lock()
	if p.loadingChildren {
		p.mu.Unlock()

		return
	}
	p.loadingChildren = true

	// Mark this URI as loaded to prevent circular loads
	normalizedURI := packaging.NormalizeURI(p.uri)
	alreadyLoaded := loadedURIs[normalizedURI]
	if !alreadyLoaded {
		loadedURIs[normalizedURI] = true
	}

	p.mu.Unlock()

	defer func() {
		p.mu.Lock()
		p.loadingChildren = false
		p.mu.Unlock()
	}()

	// If this URI was already loaded, skip its children
	if alreadyLoaded {
		return
	}

	p.mu.RLock()
	packagingPart := p.packagingPart
	p.mu.RUnlock()

	if packagingPart == nil {
		return
	}

	rels := packagingPart.Relationships()
	if rels == nil {
		return
	}

	for rel := range rels.All() {
		// Resolve target URI relative to this part
		targetURI := packaging.ResolvePartURI(
			p.uri,
			rel.Target(),
		)

		// Check if this URI has already been loaded to prevent circular loading
		normalizedTarget := packaging.NormalizeURI(targetURI)
		if loadedURIs[normalizedTarget] {
			continue // Skip already-loaded parts
		}

		// Get packaging part via container to avoid deadlock
		packPart := p.GetPackagingPart(targetURI)
		if packPart == nil {
			continue
		}

		// Create OpenXmlPart based on relationship type or content type
		var childPart OpenXmlPart
		if info, ok := GetPartTypeByRelationship(rel.Type()); ok &&
			info.Factory != nil {
			childPart = info.Factory(targetURI, p)
		} else if info, ok := GetPartTypeByContentType(packPart.ContentType()); ok &&
			info.Factory != nil {
			childPart = info.Factory(targetURI, p)
		} else {
			// Create a generic part without calling loadChildParts in constructor
			// to avoid infinite recursion. We'll load its children manually.
			childPart = &OpenXmlPartData{
				uri:           targetURI,
				contentType:   packPart.ContentType(),
				packagingPart: packPart,
				container:     p,
				childParts:    make(map[string]OpenXmlPart),
				idGenerator:   NewRelationshipIDGenerator(),
				features:      features.NewFeatureCollectionWithParent(p.features),
			}
		}

		// Set relationship ID
		switch cp := childPart.(type) {
		case *OpenXmlPartData:
			cp.SetRelationshipID(rel.ID())
		case IRelationshipIDPart:
			cp.SetRelationshipID(rel.ID())
		}

		p.childParts[rel.ID()] = childPart

		// Recursively load child parts
		// Get the underlying *OpenXmlPartData regardless of wrapping type
		var partData *OpenXmlPartData
		if cp, ok := childPart.(*OpenXmlPartData); ok {
			partData = cp
		} else {
			// Use reflection to access embedded *OpenXmlPartData field
			v := reflect.ValueOf(childPart)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				// Look for an embedded *OpenXmlPartData field
				for i := range v.NumField() {
					field := v.Field(i)
					if field.Type() == reflect.TypeOf((*OpenXmlPartData)(nil)) {
						if pd, ok := field.Interface().(*OpenXmlPartData); ok && pd != nil {
							partData = pd

							break
						}
					}
				}
			}
		}

		// Load child parts only once for the underlying partData
		if partData != nil {
			partData.loadChildPartsRecursive(
				depth + 1,
				loadedURIs,
			)
		}
	}
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
	if !p.rootLoaded && p.rootFactory != nil {
		p.rootElement = p.rootFactory()
		if p.rootElement != nil {
			p.rootElement.SetPart(p)
			// Unlock before calling Reload to avoid deadlock
			p.rootLoaded = true
			p.mu.Unlock()
			// Reload the element from the part's data if present
			_ = p.rootElement.Reload()

			return p.rootElement
		}
		p.rootLoaded = true
	}
	p.mu.Unlock()

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
	p.rootFactory = factory
	p.rootLoaded = false
	p.mu.Unlock()

	// Trigger lazy load immediately if we're not during initialization
	// or let it be lazy. For now, let's keep it lazy but ensure
	// it can be reloaded.
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

// Save serializes the root element to the part stream and recursively saves all child parts.
func (p *OpenXmlPartData) Save() error {
	p.mu.Lock()
	root := p.rootElement
	childParts := make(
		[]OpenXmlPart,
		0,
		len(p.childParts),
	)
	for _, child := range p.childParts {
		childParts = append(childParts, child)
	}
	p.mu.Unlock()

	if root != nil {
		if err := root.Save(); err != nil {
			return err
		}
	}

	// Recursively save all child parts
	for _, child := range childParts {
		if saveable, ok := child.(ISaveablePart); ok {
			// Always save child parts, not just dirty ones, to ensure
			// they are written to the package even on first save
			if err := saveable.Save(); err != nil {
				return err
			}
		}
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

// GetPartsOfType returns an iterator over child parts of a specific
// content type.
func (p *OpenXmlPartData) GetPartsOfType(
	contentType string,
) iter.Seq[OpenXmlPart] {
	return func(yield func(OpenXmlPart) bool) {
		p.mu.RLock()
		defer p.mu.RUnlock()

		for _, part := range p.childParts {
			if part.ContentType() != contentType {
				continue
			}
			if !yield(part) {
				return
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
		id = p.idGenerator.Next() //nolint:revive // modifies-parameter: intentional for generated ID
	} else {
		p.idGenerator.Reserve(id)
	}

	// Set the relationship ID on the part (works for both *OpenXmlPartData
	// and types that embed it like *WorksheetPart)
	if relPart, ok := part.(IRelationshipIDPart); ok {
		relPart.SetRelationshipID(id)
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

// GetPartsOfType is a generic function to iterate over parts of a
// specific type.
// Usage: GetPartsOfType[*StylesPart](container)
func GetPartsOfType[T OpenXmlPart](
	container OpenXmlPartContainer,
) iter.Seq[T] {
	return func(yield func(T) bool) {
		for part := range container.Parts() {
			typed, ok := part.(T)
			if !ok {
				continue
			}
			if !yield(typed) {
				return
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

// ISaveablePart is implemented by parts that can be saved.
type ISaveablePart interface {
	OpenXmlPart
	// IsDirty returns true if the part has been modified.
	IsDirty() bool
	// Save saves the part's content.
	Save() error
}

// IRelationshipIDPart is implemented by parts that have a relationship ID.
type IRelationshipIDPart interface {
	OpenXmlPart
	// RelationshipID returns the relationship ID for this part.
	RelationshipID() string
	// SetRelationshipID sets the relationship ID for this part.
	SetRelationshipID(id string)
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
