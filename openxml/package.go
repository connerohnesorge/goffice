// document processing.
//
//nolint:revive // file-length-limit: large package with many types
package openxml

import (
	"io"
	"iter"
	"reflect"
	"sync"

	"github.com/connerohnesorge/goffice/openxml/features"
	"github.com/connerohnesorge/goffice/packaging"
)

// OpenXmlPackage wraps a packaging.Package with OpenXML-specific functionality.
// It provides access to typed parts, relationships, and package-level features.
type OpenXmlPackage struct {
	mu sync.RWMutex

	// Underlying OPC package
	pkg *packaging.Package

	// Package-level features
	featureCollection *features.FeatureCollection

	// Child parts (relationship ID -> OpenXmlPart)
	parts map[string]OpenXmlPart

	// Parts by URI for fast lookup
	partsByURI map[string]OpenXmlPart

	// ID generator for relationships
	idGenerator *RelationshipIDGenerator

	// Main document part (cached for quick access)
	mainPart OpenXmlPart

	// Track if any parts are modified
	isDirty bool
}

// NewOpenXmlPackage creates a new OpenXmlPackage wrapping the given
// OPC package.
func NewOpenXmlPackage(
	pkg *packaging.Package,
) *OpenXmlPackage {
	oxp := &OpenXmlPackage{
		pkg:               pkg,
		featureCollection: features.NewFeatureCollection(),
		parts: make(
			map[string]OpenXmlPart,
		),
		partsByURI: make(
			map[string]OpenXmlPart,
		),
		idGenerator: NewRelationshipIDGenerator(),
		isDirty:     false,
	}

	// Register package-level features
	oxp.registerFeatures()

	// Load parts from the underlying package
	oxp.loadParts()

	return oxp
}

// registerFeatures sets up the package-level feature implementations.
func (p *OpenXmlPackage) registerFeatures() {
	// Register IPackageFeature
	pkgFeature := &packageFeature{pkg: p}
	p.featureCollection.SetByType(
		reflect.TypeFor[features.IPackageFeature](),
		pkgFeature,
	)

	// Register IContentTypeFeature
	ctFeature := &contentTypeFeature{pkg: p}
	p.featureCollection.SetByType(
		reflect.TypeFor[features.IContentTypeFeature](),
		ctFeature,
	)
}

// loadParts loads OpenXmlParts from the underlying OPC package.
func (p *OpenXmlPackage) loadParts() {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Get package-level relationships
	rels := p.pkg.Relationships()
	if rels == nil {
		return
	}

	// Load parts based on relationships
	for rel := range rels.All() {
		p.idGenerator.Reserve(rel.ID())

		// Resolve target URI
		targetURI := packaging.ResolvePartURI(
			"/",
			rel.Target(),
		)

		// Get the underlying packaging part
		packPart, err := p.pkg.Part(targetURI)
		if err != nil {
			continue // Skip if part doesn't exist
		}

		// Create OpenXmlPart based on relationship type or content type
		var part OpenXmlPart
		if info, ok := GetPartTypeByRelationship(rel.Type()); ok &&
			info.Factory != nil {
			part = info.Factory(targetURI, p)
		} else {
			contentType := packPart.ContentType()
			part = NewOpenXmlPartData(
				targetURI, contentType, packPart, p,
			)
		}

		// Set relationship ID
		if partData, ok := part.(*OpenXmlPartData); ok {
			partData.SetRelationshipID(rel.ID())
		}

		p.parts[rel.ID()] = part
		p.partsByURI[targetURI] = part

		// Check if this is the main document part
		// Support both Microsoft namespace and PURL namespace variants
		if rel.Type() == RelationshipTypeOfficeDocument ||
			rel.Type() == RelationshipTypeDocument ||
			rel.Type() == RelationshipTypePURLOfficeDocument {
			p.mainPart = part
		}
	}

	// Now that all top-level parts are created, load their children recursively
	for _, part := range p.parts {
		if partData, ok := part.(*OpenXmlPartData); ok {
			partData.loadChildParts()
		} else {
			// Handle types that embed *OpenXmlPartData
			// Use reflection to access the embedded field
			v := reflect.ValueOf(part)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				// Look for an embedded *OpenXmlPartData field
				for i := range v.NumField() {
					field := v.Field(i)
					if field.Type() == reflect.TypeOf((*OpenXmlPartData)(nil)) {
						if partData, ok := field.Interface().(*OpenXmlPartData); ok && partData != nil {
							partData.loadChildParts()

							break
						}
					}
				}
			}
		}
	}
}

// Package returns the underlying OPC Package.
func (p *OpenXmlPackage) Package() *packaging.Package {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.pkg
}

// GetPackagingPart returns the underlying packaging.Part by its URI.
// This method doesn't acquire locks and is safe to call during part loading.
func (p *OpenXmlPackage) GetPackagingPart(
	uri string,
) *packaging.Part {
	// Don't acquire locks - this is called during initialization
	// The caller (loadParts) already holds the write lock
	if p.pkg == nil {
		return nil
	}
	part, err := p.pkg.Part(uri)
	if err != nil {
		return nil
	}

	return part
}

// URI returns "/" as this is the package root.
func (*OpenXmlPackage) URI() string {
	return "/"
}

// Features returns the package-level feature collection.
func (p *OpenXmlPackage) Features() *features.FeatureCollection {
	return p.featureCollection
}

// Parts returns an iterator over all child parts.
func (p *OpenXmlPackage) Parts() iter.Seq[OpenXmlPart] {
	return func(yield func(OpenXmlPart) bool) {
		p.mu.RLock()
		defer p.mu.RUnlock()

		for _, part := range p.parts {
			if !yield(part) {
				return
			}
		}
	}
}

// GetPartById returns a part by its relationship ID.
func (p *OpenXmlPackage) GetPartById(
	id string,
) (OpenXmlPart, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	part, ok := p.parts[id]
	if !ok {
		return nil, ErrPartNotFound
	}

	return part, nil
}

// GetPartByURI returns a part by its URI.
func (p *OpenXmlPackage) GetPartByURI(
	uri string,
) (OpenXmlPart, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	normalizedURI := packaging.NormalizeURI(uri)
	part, ok := p.partsByURI[normalizedURI]
	if !ok {
		return nil, ErrPartNotFound
	}

	return part, nil
}

// GetPartsOfType returns an iterator over parts of a specific content type.
func (p *OpenXmlPackage) GetPartsOfType(
	contentType string,
) iter.Seq[OpenXmlPart] {
	return func(yield func(OpenXmlPart) bool) {
		p.mu.RLock()
		defer p.mu.RUnlock()

		for _, part := range p.parts {
			if part.ContentType() != contentType {
				continue
			}
			if !yield(part) {
				return
			}
		}
	}
}

// AddPart adds a part to the package with the given relationship ID.
func (p *OpenXmlPackage) AddPart(
	part OpenXmlPart,
	id string,
) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if id == "" {
		id = p.idGenerator.Next() //nolint:revive // modifies-parameter: intentional ID generation
	} else {
		p.idGenerator.Reserve(id)
	}

	// Check if part already exists
	normalizedURI := packaging.NormalizeURI(
		part.URI(),
	)
	if _, exists := p.partsByURI[normalizedURI]; exists {
		return ErrPartExists
	}

	// Set relationship ID on part (works for both *OpenXmlPartData
	// and types that embed it like *WorkbookPart)
	if relPart, ok := part.(IRelationshipIDPart); ok {
		relPart.SetRelationshipID(id)
	}

	p.parts[id] = part
	p.partsByURI[normalizedURI] = part
	p.isDirty = true

	return nil
}

// AddNewPart creates and adds a new part with the given URI,
// content type, and relationship type.
func (p *OpenXmlPackage) AddNewPart(
	uri, contentType, relType string,
) (OpenXmlPart, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	normalizedURI := packaging.NormalizeURI(uri)

	// Check if part already exists
	if _, exists := p.partsByURI[normalizedURI]; exists {
		return nil, ErrPartExists
	}

	// Create the underlying OPC part
	packPart, err := p.pkg.CreatePart(
		uri,
		contentType,
	)
	if err != nil {
		return nil, err
	}

	// Create OpenXmlPart
	var part OpenXmlPart
	if info, ok := GetPartTypeByContentType(contentType); ok &&
		info.Factory != nil {
		part = info.Factory(normalizedURI, p)
	} else if info, ok := GetPartTypeByRelationship(relType); ok &&
		info.Factory != nil {
		part = info.Factory(normalizedURI, p)
	} else {
		part = NewOpenXmlPartData(normalizedURI, contentType, packPart, p)
	}

	// Generate relationship ID
	id := p.idGenerator.Next()

	// Set relationship ID on part
	if partData, ok := part.(*OpenXmlPartData); ok {
		partData.SetRelationshipID(id)
	}

	// Create package-level relationship
	target := normalizedURI
	if relType == "" {
		relType = RelationshipTypeDocument //nolint:revive // modifies-parameter: intentional default
	}
	_, err = p.pkg.CreateRelationship(
		target,
		relType,
		id,
	)
	if err != nil {
		return nil, err
	}

	p.parts[id] = part
	p.partsByURI[normalizedURI] = part
	p.isDirty = true

	return part, nil
}

// DeletePart removes a part by its relationship ID.
func (p *OpenXmlPackage) DeletePart(
	id string,
) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	part, ok := p.parts[id]
	if !ok {
		return ErrPartNotFound
	}

	// Remove from package
	if err := p.pkg.DeletePart(part.URI()); err != nil {
		return err
	}

	// Remove relationship - ignore error if relationship doesn't exist
	_ = p.pkg.DeleteRelationship(id)

	delete(p.parts, id)
	delete(p.partsByURI, part.URI())
	p.isDirty = true

	return nil
}

// MainPart returns the main document part.
func (p *OpenXmlPackage) MainPart() OpenXmlPart {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.mainPart
}

// SetMainPart sets the main document part.
func (p *OpenXmlPackage) SetMainPart(
	part OpenXmlPart,
) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.mainPart = part
}

// IsDirty returns true if any parts have been modified.
func (p *OpenXmlPackage) IsDirty() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.isDirty {
		return true
	}

	// Check if any parts are dirty (works for both *OpenXmlPartData
	// and types that embed it like *WorkbookPart)
	for _, part := range p.parts {
		if saveable, ok := part.(ISaveablePart); ok &&
			saveable.IsDirty() {
			return true
		}
	}

	return false
}

// Save saves all modified parts to the package.
func (p *OpenXmlPackage) Save() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Save all dirty parts
	for _, part := range p.parts {
		// Check if the part implements ISaveablePart (works for both
		// *OpenXmlPartData and types that embed it like *WorkbookPart)
		if saveable, ok := part.(ISaveablePart); ok {
			if saveable.IsDirty() {
				if err := saveable.Save(); err != nil {
					return err
				}
			}
		}
	}

	// Save the underlying package
	if err := p.pkg.Save(); err != nil {
		return err
	}

	p.isDirty = false

	return nil
}

// SaveAs saves the package to a new file path.
func (p *OpenXmlPackage) SaveAs(
	path string,
) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Save all dirty parts first
	for _, part := range p.parts {
		// Check if the part implements ISaveablePart (works for both
		// *OpenXmlPartData and types that embed it like *WorkbookPart)
		if saveable, ok := part.(ISaveablePart); ok {
			if saveable.IsDirty() {
				if err := saveable.Save(); err != nil {
					return err
				}
			}
		}
	}

	// Save to new location
	if err := p.pkg.SaveAs(path); err != nil {
		return err
	}

	p.isDirty = false

	return nil
}

// Close closes the package.
// For stream-based packages, all parts are saved before closing.
func (p *OpenXmlPackage) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.pkg != nil {
		// Save all dirty parts before closing so content is available
		// for stream-based packages
		for _, part := range p.parts {
			if saveable, ok := part.(ISaveablePart); ok {
				if saveable.IsDirty() {
					if err := saveable.Save(); err != nil {
						return err
					}
				}
			}
		}

		err := p.pkg.Close()
		p.pkg = nil
		p.parts = nil
		p.partsByURI = nil
		p.mainPart = nil

		return err
	}

	return nil
}

// Feature implementations

// packageFeature implements IPackageFeature.
type packageFeature struct {
	features.FeatureBase
	pkg *OpenXmlPackage
}

// Ensure packageFeature implements the Feature marker.
func (*packageFeature) featureMarker() {} //nolint:unused // interface implementation marker

// Package returns the underlying packaging.Package.
func (f *packageFeature) Package() any {
	if f.pkg != nil {
		return f.pkg.pkg
	}

	return nil
}

// Capabilities returns the package capabilities.
func (f *packageFeature) Capabilities() features.PackageCapabilities {
	if f.pkg == nil || f.pkg.pkg == nil {
		return features.PackageCapabilities{}
	}

	caps := f.pkg.pkg.Capability()

	return features.PackageCapabilities{
		CanRead: caps == packaging.Read ||
			caps == packaging.ReadWrite,
		CanWrite: caps == packaging.Write ||
			caps == packaging.ReadWrite,
		CanSave: caps == packaging.Write ||
			caps == packaging.ReadWrite,
	}
}

// contentTypeFeature implements IContentTypeFeature.
type contentTypeFeature struct {
	features.FeatureBase
	pkg *OpenXmlPackage
}

// Ensure contentTypeFeature implements the Feature marker.
func (*contentTypeFeature) featureMarker() {} //nolint:unused // interface implementation marker

// GetContentType returns the content type for a part URI.
func (f *contentTypeFeature) GetContentType(
	uri string,
) (string, error) {
	if f.pkg == nil || f.pkg.pkg == nil {
		return "", ErrPartNotFound
	}

	ct := f.pkg.pkg.ContentTypes()
	if ct == nil {
		return "", ErrPartNotFound
	}

	return ct.GetContentType(uri)
}

// SetContentType sets the content type for a part URI.
func (f *contentTypeFeature) SetContentType(
	uri, contentType string,
) error {
	if f.pkg == nil || f.pkg.pkg == nil {
		return ErrPartNotFound
	}

	ct := f.pkg.pkg.ContentTypes()
	if ct == nil {
		return ErrPartNotFound
	}

	ct.SetOverride(uri, contentType)

	return nil
}

// RemoveContentType removes the content type for a part URI.
func (f *contentTypeFeature) RemoveContentType(
	uri string,
) {
	if f.pkg == nil || f.pkg.pkg == nil {
		return
	}

	ct := f.pkg.pkg.ContentTypes()
	if ct != nil {
		ct.RemoveOverride(uri)
	}
}

// mainPartFeature implements IMainPartFeature.
type mainPartFeature struct {
	features.FeatureBase
	pkg              *OpenXmlPackage
	contentType      string
	relationshipType string
}

// Ensure mainPartFeature implements the Feature marker.
func (*mainPartFeature) featureMarker() {} //nolint:unused // interface implementation marker

// MainPart returns the main document part.
func (f *mainPartFeature) MainPart() any {
	if f.pkg != nil {
		return f.pkg.MainPart()
	}

	return nil
}

// ContentType returns the expected content type of the main part.
func (f *mainPartFeature) ContentType() string {
	return f.contentType
}

// RelationshipType returns the relationship type for the main part.
func (f *mainPartFeature) RelationshipType() string {
	return f.relationshipType
}

// SetMainPartInfo configures the main part feature with
// document-specific information.
func (p *OpenXmlPackage) SetMainPartInfo(
	contentType, relationshipType string,
) {
	feature := &mainPartFeature{
		pkg:              p,
		contentType:      contentType,
		relationshipType: relationshipType,
	}
	p.featureCollection.SetByType(
		reflect.TypeFor[features.IMainPartFeature](),
		feature,
	)
}

// OpenPackage opens an existing OpenXML package from a file.
func OpenPackage(
	path string,
	readOnly bool,
) (*OpenXmlPackage, error) {
	pkg, err := packaging.Open(path, readOnly)
	if err != nil {
		return nil, err
	}

	return NewOpenXmlPackage(pkg), nil
}

// CreatePackage creates a new OpenXML package at the given path.
func CreatePackage(
	path string,
) (*OpenXmlPackage, error) {
	pkg, err := packaging.Create(path)
	if err != nil {
		return nil, err
	}

	return NewOpenXmlPackage(pkg), nil
}

// OpenPackageFromReader opens an OpenXML package from an io.ReaderAt.
func OpenPackageFromReader(
	r io.ReaderAt,
	size int64,
) (*OpenXmlPackage, error) {
	pkg, err := packaging.OpenReader(r, size)
	if err != nil {
		return nil, err
	}

	return NewOpenXmlPackage(pkg), nil
}

// Ensure OpenXmlPackage implements OpenXmlPartContainer.
var _ OpenXmlPartContainer = (*OpenXmlPackage)(
	nil,
)
