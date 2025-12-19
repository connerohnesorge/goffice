// Package packaging provides the OPC (Open Packaging Conventions) layer.
package packaging

import (
	"bytes"
	"io"
	"sync"
)

// Part represents a part within an OPC package.
// A part is identified by its URI and has associated content type and data.
type Part struct {
	mu          sync.RWMutex
	uri         string
	contentType string
	data        []byte
	modified    bool
	pkg         *Package // parent package reference
}

// newPart creates a new Part with the given URI and content type.
func newPart(
	uri, contentType string,
	pkg *Package,
) *Part {
	return &Part{
		uri:         NormalizeURI(uri),
		contentType: contentType,
		data:        nil,
		modified:    false,
		pkg:         pkg,
	}
}

// URI returns the normalized URI of the part.
func (p *Part) URI() string {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.uri
}

// ContentType returns the content type of the part.
func (p *Part) ContentType() string {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.contentType
}

// GetStream returns a reader for the part's content.
// The reader provides access to the current content of the part.
func (p *Part) GetStream() io.Reader {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return bytes.NewReader(p.data)
}

// SetStream sets the part's content from the given reader.
// The content is read entirely into memory.
func (p *Part) SetStream(r io.Reader) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	p.data = data
	p.modified = true

	return nil
}

// GetData returns a copy of the part's data.
func (p *Part) GetData() []byte {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.data == nil {
		return nil
	}

	result := make([]byte, len(p.data))
	copy(result, p.data)

	return result
}

// SetData sets the part's content directly.
func (p *Part) SetData(data []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.data = make([]byte, len(data))
	copy(p.data, data)
	p.modified = true
}

// Size returns the size of the part's content in bytes.
func (p *Part) Size() int64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return int64(len(p.data))
}

// IsModified returns true if the part has been modified since creation or loading.
func (p *Part) IsModified() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.modified
}

// clearModified clears the modified flag (used after saving).
func (p *Part) clearModified() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.modified = false
}

// setData sets the data without marking as modified (for loading).
func (p *Part) setData(data []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.data = data
}

// Package returns the parent package.
func (p *Part) Package() *Package {
	return p.pkg
}

// Relationships returns the relationships for this part.
// Returns nil if the part has no relationships.
func (p *Part) Relationships() *Relationships {
	if p.pkg == nil {
		return nil
	}

	return p.pkg.PartRelationships(p.uri)
}

// CreateRelationship creates a relationship from this part to a target.
func (p *Part) CreateRelationship(
	target, relType, id string,
) (*Relationship, error) {
	if p.pkg == nil {
		return nil, ErrPackageClosed
	}

	return p.pkg.CreatePartRelationship(
		p.uri,
		target,
		relType,
		id,
	)
}

// DeleteRelationship deletes a relationship from this part.
func (p *Part) DeleteRelationship(
	id string,
) error {
	if p.pkg == nil {
		return ErrPackageClosed
	}

	return p.pkg.DeletePartRelationship(p.uri, id)
}
