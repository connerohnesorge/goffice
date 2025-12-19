// Package packaging provides the OPC (Open Packaging Conventions) layer.
// OPC defines how Office Open XML documents are stored as ZIP-based packages
// containing XML parts, relationships, and content types.
package packaging

import (
	"archive/zip"
	"bytes"
	"io"
	"iter"
	"os"
	"strings"
	"sync"
)

// PackageCapability represents the access mode of a package.
type PackageCapability int

const (
	// Read indicates the package is open for reading only.
	Read PackageCapability = iota
	// Write indicates the package is open for writing only.
	Write
	// ReadWrite indicates the package is open for both reading and writing.
	ReadWrite
)

// String returns the string representation of the capability.
func (pc PackageCapability) String() string {
	switch pc {
	case Read:
		return "Read"
	case Write:
		return "Write"
	case ReadWrite:
		return "ReadWrite"
	default:
		return "Unknown"
	}
}

// Package represents an OPC package (typically a .docx, .xlsx, etc.).
// It provides access to parts, relationships, and content types.
type Package struct {
	mu sync.RWMutex

	// Package state
	path       string
	capability PackageCapability
	closed     bool

	// ZIP backing
	zipReader *zip.Reader
	zipData   []byte // loaded ZIP data for in-memory operations

	// Package contents
	parts           map[string]*Part          // normalized URI -> Part
	relationships   *Relationships            // package-level relationships
	partRels        map[string]*Relationships // part URI -> part relationships
	contentTypes    *ContentTypes
	coreProperties  *CoreProperties

	// Writer for streaming output
	writer io.Writer
}

// Create creates a new package at the given file path.
// The package will have ReadWrite capability.
func Create(path string) (*Package, error) {
	pkg := &Package{
		path:           path,
		capability:     ReadWrite,
		closed:         false,
		parts:          make(map[string]*Part),
		relationships:  NewRelationships("/"),
		partRels:       make(map[string]*Relationships),
		contentTypes:   NewContentTypes(),
		coreProperties: NewCoreProperties(),
	}

	// Initialize with standard defaults for Office documents
	pkg.contentTypes.SetDefault("xml", "application/xml")
	pkg.contentTypes.SetDefault("rels", "application/vnd.openxmlformats-package.relationships+xml")

	return pkg, nil
}

// CreateWriter creates a new package that writes to the given io.Writer.
// The package will have Write capability only.
func CreateWriter(w io.Writer) (*Package, error) {
	pkg := &Package{
		path:           "",
		capability:     Write,
		closed:         false,
		writer:         w,
		parts:          make(map[string]*Part),
		relationships:  NewRelationships("/"),
		partRels:       make(map[string]*Relationships),
		contentTypes:   NewContentTypes(),
		coreProperties: NewCoreProperties(),
	}

	// Initialize with standard defaults
	pkg.contentTypes.SetDefault("xml", "application/xml")
	pkg.contentTypes.SetDefault("rels", "application/vnd.openxmlformats-package.relationships+xml")

	return pkg, nil
}

// Open opens an existing package from a file path.
// If readOnly is true, the package is opened with Read capability.
// Otherwise, it is opened with ReadWrite capability.
func Open(path string, readOnly bool) (*Package, error) {
	// Read the file into memory
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	capability := ReadWrite
	if readOnly {
		capability = Read
	}

	pkg, err := openFromBytes(data, capability)
	if err != nil {
		return nil, err
	}

	pkg.path = path
	return pkg, nil
}

// OpenReader opens a package from an io.ReaderAt.
// The package will have Read capability only.
func OpenReader(r io.ReaderAt, size int64) (*Package, error) {
	// Read all data into memory
	data := make([]byte, size)
	if _, err := r.ReadAt(data, 0); err != nil && err != io.EOF {
		return nil, err
	}

	return openFromBytes(data, Read)
}

// openFromBytes opens a package from in-memory data.
func openFromBytes(data []byte, capability PackageCapability) (*Package, error) {
	reader := bytes.NewReader(data)
	zipReader, err := zip.NewReader(reader, int64(len(data)))
	if err != nil {
		return nil, err
	}

	pkg := &Package{
		capability:     capability,
		closed:         false,
		zipReader:      zipReader,
		zipData:        data,
		parts:          make(map[string]*Part),
		relationships:  NewRelationships("/"),
		partRels:       make(map[string]*Relationships),
		contentTypes:   NewContentTypes(),
		coreProperties: NewCoreProperties(),
	}

	// Load content types first
	if err := pkg.loadContentTypes(); err != nil {
		return nil, err
	}

	// Load all parts
	if err := pkg.loadParts(); err != nil {
		return nil, err
	}

	// Load package-level relationships
	if err := pkg.loadRelationships(); err != nil {
		return nil, err
	}

	// Load core properties if present
	pkg.loadCoreProperties()

	return pkg, nil
}

// loadContentTypes loads [Content_Types].xml from the package.
func (p *Package) loadContentTypes() error {
	for _, f := range p.zipReader.File {
		if strings.EqualFold(f.Name, "[Content_Types].xml") {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			return p.contentTypes.UnmarshalFromXML(rc)
		}
	}
	return ErrInvalidPackage
}

// loadParts loads all parts from the ZIP.
func (p *Package) loadParts() error {
	for _, f := range p.zipReader.File {
		name := f.Name

		// Skip [Content_Types].xml
		if strings.EqualFold(name, "[Content_Types].xml") {
			continue
		}

		// Normalize URI
		uri := NormalizeURI(name)

		// Skip relationship parts (they're loaded separately)
		if IsRelationshipURI(uri) {
			continue
		}

		// Get content type
		contentType, err := p.contentTypes.GetContentType(uri)
		if err != nil {
			// Use a default if not found
			contentType = "application/octet-stream"
		}

		// Create part
		part := newPart(uri, contentType, p)

		// Load data lazily - for now, load it all
		rc, err := f.Open()
		if err != nil {
			return err
		}
		data, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return err
		}
		part.setData(data)

		p.parts[uri] = part
	}

	// Load part-level relationships
	for _, f := range p.zipReader.File {
		uri := NormalizeURI(f.Name)
		if !IsRelationshipURI(uri) {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		sourceURI := PartURIFromRelationshipURI(uri)
		rels := NewRelationships(sourceURI)
		if err := rels.UnmarshalFromXML(rc); err != nil {
			rc.Close()
			return err
		}
		rc.Close()

		if sourceURI == "/" {
			p.relationships = rels
		} else {
			p.partRels[sourceURI] = rels
		}
	}

	return nil
}

// loadRelationships loads package-level relationships from _rels/.rels.
func (p *Package) loadRelationships() error {
	for _, f := range p.zipReader.File {
		if strings.EqualFold(f.Name, "_rels/.rels") {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			return p.relationships.UnmarshalFromXML(rc)
		}
	}
	// It's okay if there are no package-level relationships
	return nil
}

// loadCoreProperties loads core properties if present.
func (p *Package) loadCoreProperties() {
	part, ok := p.parts[NormalizeURI(CorePropertiesPartURI)]
	if !ok {
		return
	}

	r := part.GetStream()
	p.coreProperties.UnmarshalFromXML(r)
}

// Save saves the package to its original location.
// Returns ErrReadOnly if the package is read-only.
func (p *Package) Save() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrPackageClosed
	}

	if p.capability == Read {
		return ErrReadOnly
	}

	if p.path == "" && p.writer == nil {
		return ErrInvalidPackage
	}

	if p.path != "" {
		return p.saveToFile(p.path)
	}

	return p.saveToWriter(p.writer)
}

// SaveAs saves the package to a new file path.
// The original file is unchanged.
func (p *Package) SaveAs(path string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrPackageClosed
	}

	return p.saveToFile(path)
}

// saveToFile writes the package to a file.
func (p *Package) saveToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return p.saveToWriter(f)
}

// saveToWriter writes the package to an io.Writer.
func (p *Package) saveToWriter(w io.Writer) error {
	zw := zip.NewWriter(w)
	defer zw.Close()

	// Write [Content_Types].xml first
	ctData, err := p.contentTypes.MarshalToXML()
	if err != nil {
		return err
	}
	if err := writeZipFile(zw, "[Content_Types].xml", ctData); err != nil {
		return err
	}

	// Write package-level relationships
	if !p.relationships.IsEmpty() {
		relsData, err := p.relationships.MarshalToXML()
		if err != nil {
			return err
		}
		if err := writeZipFile(zw, "_rels/.rels", relsData); err != nil {
			return err
		}
	}

	// Write all parts
	for uri, part := range p.parts {
		// Skip the leading slash for ZIP paths
		zipPath := strings.TrimPrefix(uri, "/")
		data := part.GetData()
		if err := writeZipFile(zw, zipPath, data); err != nil {
			return err
		}
		part.clearModified()
	}

	// Write part-level relationships
	for sourceURI, rels := range p.partRels {
		if rels.IsEmpty() {
			continue
		}

		relsData, err := rels.MarshalToXML()
		if err != nil {
			return err
		}

		relPath := RelationshipPartURI(sourceURI)
		zipPath := strings.TrimPrefix(relPath, "/")
		if err := writeZipFile(zw, zipPath, relsData); err != nil {
			return err
		}
	}

	return nil
}

// writeZipFile writes a file to the ZIP archive.
func writeZipFile(zw *zip.Writer, name string, data []byte) error {
	w, err := zw.Create(name)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

// Close closes the package and releases all resources.
// Any unsaved changes are discarded.
func (p *Package) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil
	}

	p.closed = true
	p.zipReader = nil
	p.zipData = nil
	p.parts = nil
	p.relationships = nil
	p.partRels = nil
	p.contentTypes = nil
	p.coreProperties = nil
	p.writer = nil

	return nil
}

// Capability returns the package's access capability.
func (p *Package) Capability() PackageCapability {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.capability
}

// Path returns the file path of the package.
// Returns empty string if the package was created with a writer.
func (p *Package) Path() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.path
}

// IsClosed returns true if the package has been closed.
func (p *Package) IsClosed() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.closed
}

// CreatePart creates a new part with the given URI and content type.
// Returns ErrPartExists if a part with that URI already exists.
// Returns ErrReadOnly if the package is read-only.
func (p *Package) CreatePart(uri, contentType string) (*Part, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil, ErrPackageClosed
	}

	if p.capability == Read {
		return nil, ErrReadOnly
	}

	normalizedURI := NormalizeURI(uri)
	if !ValidateURI(normalizedURI) {
		return nil, ErrInvalidURI
	}

	if _, exists := p.parts[normalizedURI]; exists {
		return nil, ErrPartExists
	}

	part := newPart(normalizedURI, contentType, p)
	p.parts[normalizedURI] = part

	// Register content type
	p.contentTypes.SetOverride(normalizedURI, contentType)

	return part, nil
}

// Part returns the part with the given URI.
// Returns ErrPartNotFound if no such part exists.
func (p *Package) Part(uri string) (*Part, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return nil, ErrPackageClosed
	}

	normalizedURI := NormalizeURI(uri)
	part, ok := p.parts[normalizedURI]
	if !ok {
		return nil, ErrPartNotFound
	}

	return part, nil
}

// DeletePart removes the part with the given URI.
// Returns ErrPartNotFound if no such part exists.
// Returns ErrReadOnly if the package is read-only.
func (p *Package) DeletePart(uri string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrPackageClosed
	}

	if p.capability == Read {
		return ErrReadOnly
	}

	normalizedURI := NormalizeURI(uri)
	if _, exists := p.parts[normalizedURI]; !exists {
		return ErrPartNotFound
	}

	delete(p.parts, normalizedURI)
	p.contentTypes.RemoveOverride(normalizedURI)

	// Remove part relationships
	delete(p.partRels, normalizedURI)

	// Remove relationships to this part from package level
	p.removeRelationshipsTo(normalizedURI, p.relationships)

	// Remove relationships to this part from other parts
	for _, rels := range p.partRels {
		p.removeRelationshipsTo(normalizedURI, rels)
	}

	return nil
}

// removeRelationshipsTo removes all relationships targeting the given URI.
func (p *Package) removeRelationshipsTo(targetURI string, rels *Relationships) {
	var toDelete []string
	for rel := range rels.All() {
		if NormalizeURI(rel.Target()) == targetURI {
			toDelete = append(toDelete, rel.ID())
		}
	}
	for _, id := range toDelete {
		rels.Delete(id)
	}
}

// Parts returns an iterator over all parts in the package.
// Relationship parts are excluded.
func (p *Package) Parts() iter.Seq[*Part] {
	return func(yield func(*Part) bool) {
		p.mu.RLock()
		defer p.mu.RUnlock()

		if p.closed {
			return
		}

		for _, part := range p.parts {
			if !yield(part) {
				return
			}
		}
	}
}

// PartCount returns the number of parts in the package.
func (p *Package) PartCount() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return len(p.parts)
}

// ContentTypes returns the content types manager for this package.
func (p *Package) ContentTypes() *ContentTypes {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.contentTypes
}

// CreateRelationship creates a package-level relationship.
// If id is empty, an auto-generated ID will be used.
func (p *Package) CreateRelationship(target, relType, id string) (*Relationship, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil, ErrPackageClosed
	}

	if p.capability == Read {
		return nil, ErrReadOnly
	}

	return p.relationships.Create(target, relType, id)
}

// Relationship returns the package-level relationship with the given ID.
func (p *Package) Relationship(id string) (*Relationship, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return nil, ErrPackageClosed
	}

	return p.relationships.Get(id)
}

// RelationshipsByType returns an iterator over package-level relationships of the given type.
func (p *Package) RelationshipsByType(relType string) iter.Seq[*Relationship] {
	return func(yield func(*Relationship) bool) {
		p.mu.RLock()
		defer p.mu.RUnlock()

		if p.closed {
			return
		}

		for rel := range p.relationships.ByType(relType) {
			if !yield(rel) {
				return
			}
		}
	}
}

// DeleteRelationship deletes a package-level relationship.
func (p *Package) DeleteRelationship(id string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrPackageClosed
	}

	if p.capability == Read {
		return ErrReadOnly
	}

	return p.relationships.Delete(id)
}

// Relationships returns all package-level relationships.
func (p *Package) Relationships() *Relationships {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.relationships
}

// PartRelationships returns the relationships for a specific part.
// Returns nil if the part has no relationships.
func (p *Package) PartRelationships(partURI string) *Relationships {
	p.mu.RLock()
	defer p.mu.RUnlock()

	normalizedURI := NormalizeURI(partURI)
	return p.partRels[normalizedURI]
}

// CreatePartRelationship creates a relationship from a part to a target.
func (p *Package) CreatePartRelationship(sourceURI, target, relType, id string) (*Relationship, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil, ErrPackageClosed
	}

	if p.capability == Read {
		return nil, ErrReadOnly
	}

	normalizedURI := NormalizeURI(sourceURI)

	// Ensure relationships collection exists
	rels, ok := p.partRels[normalizedURI]
	if !ok {
		rels = NewRelationships(normalizedURI)
		p.partRels[normalizedURI] = rels
	}

	return rels.Create(target, relType, id)
}

// DeletePartRelationship deletes a relationship from a part.
func (p *Package) DeletePartRelationship(sourceURI, id string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrPackageClosed
	}

	if p.capability == Read {
		return ErrReadOnly
	}

	normalizedURI := NormalizeURI(sourceURI)
	rels, ok := p.partRels[normalizedURI]
	if !ok {
		return ErrRelationshipNotFound
	}

	return rels.Delete(id)
}

// CoreProperties returns the core document properties.
func (p *Package) CoreProperties() *CoreProperties {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.coreProperties
}

// EnsureCorePropertiesPart ensures the core properties part exists in the package.
// This should be called before saving if core properties have been modified.
func (p *Package) EnsureCorePropertiesPart() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrPackageClosed
	}

	if p.capability == Read {
		return ErrReadOnly
	}

	uri := NormalizeURI(CorePropertiesPartURI)

	// Create part if it doesn't exist
	if _, exists := p.parts[uri]; !exists {
		part := newPart(uri, CorePropertiesContentType, p)
		p.parts[uri] = part
		p.contentTypes.SetOverride(uri, CorePropertiesContentType)

		// Create relationship to core properties
		p.relationships.Create(CorePropertiesPartURI, CorePropertiesRelationshipType, "")
	}

	// Serialize core properties to part
	data, err := p.coreProperties.MarshalToXML()
	if err != nil {
		return err
	}

	p.parts[uri].SetData(data)
	return nil
}
