// Package packaging provides the OPC (Open Packaging Conventions) layer.
package packaging

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"iter"
	"sync"
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

// Relationship represents an OPC relationship between a source and target.
type Relationship struct {
	id         string
	relType    string
	target     string
	targetMode TargetMode
}

// NewRelationship creates a new relationship with the given parameters.
func NewRelationship(
	id, relType, target string,
	targetMode TargetMode,
) *Relationship {
	return &Relationship{
		id:         id,
		relType:    relType,
		target:     target,
		targetMode: targetMode,
	}
}

// ID returns the relationship ID.
func (r *Relationship) ID() string {
	return r.id
}

// Type returns the relationship type URI.
func (r *Relationship) Type() string {
	return r.relType
}

// Target returns the target URI.
func (r *Relationship) Target() string {
	return r.target
}

// TargetMode returns the target mode (Internal or External).
func (r *Relationship) TargetMode() TargetMode {
	return r.targetMode
}

// Relationships manages a collection of relationships for a source part.
type Relationships struct {
	mu        sync.RWMutex
	rels      map[string]*Relationship // id -> relationship
	sourceURI string                   // URI of the source part (or "/" for package-level)
	nextID    int                      // for auto-generating IDs
}

// NewRelationships creates a new Relationships collection for the given source URI.
func NewRelationships(
	sourceURI string,
) *Relationships {
	return &Relationships{
		rels:      make(map[string]*Relationship),
		sourceURI: sourceURI,
		nextID:    1,
	}
}

// SourceURI returns the URI of the source part for these relationships.
func (rs *Relationships) SourceURI() string {
	return rs.sourceURI
}

// Create creates a new relationship with the given parameters.
// If id is empty, an auto-generated ID will be used.
// Returns ErrRelationshipExists if a relationship with the same ID already exists.
func (rs *Relationships) Create(
	target, relType, id string,
) (*Relationship, error) {
	return rs.CreateWithMode(
		target,
		relType,
		id,
		TargetModeInternal,
	)
}

// CreateWithMode creates a new relationship with the given parameters and target mode.
// If id is empty, an auto-generated ID will be used.
// Returns ErrRelationshipExists if a relationship with the same ID already exists.
func (rs *Relationships) CreateWithMode(
	target, relType, id string,
	mode TargetMode,
) (*Relationship, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	// Generate ID if not provided
	if id == "" {
		id = rs.generateID()
	}

	// Check for duplicate
	if _, exists := rs.rels[id]; exists {
		return nil, ErrRelationshipExists
	}

	rel := NewRelationship(
		id,
		relType,
		target,
		mode,
	)
	rs.rels[id] = rel

	return rel, nil
}

// generateID generates a unique relationship ID (must be called with lock held).
func (rs *Relationships) generateID() string {
	for {
		id := fmt.Sprintf("rId%d", rs.nextID)
		rs.nextID++
		if _, exists := rs.rels[id]; !exists {
			return id
		}
	}
}

// Get returns the relationship with the given ID.
// Returns ErrRelationshipNotFound if no such relationship exists.
func (rs *Relationships) Get(
	id string,
) (*Relationship, error) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	if rel, ok := rs.rels[id]; ok {
		return rel, nil
	}
	return nil, ErrRelationshipNotFound
}

// Delete removes the relationship with the given ID.
// Returns ErrRelationshipNotFound if no such relationship exists.
func (rs *Relationships) Delete(id string) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if _, exists := rs.rels[id]; !exists {
		return ErrRelationshipNotFound
	}

	delete(rs.rels, id)
	return nil
}

// ByType returns an iterator over all relationships of the given type.
func (rs *Relationships) ByType(
	relType string,
) iter.Seq[*Relationship] {
	return func(yield func(*Relationship) bool) {
		rs.mu.RLock()
		defer rs.mu.RUnlock()

		for _, rel := range rs.rels {
			if rel.relType == relType {
				if !yield(rel) {
					return
				}
			}
		}
	}
}

// All returns an iterator over all relationships.
func (rs *Relationships) All() iter.Seq[*Relationship] {
	return func(yield func(*Relationship) bool) {
		rs.mu.RLock()
		defer rs.mu.RUnlock()

		for _, rel := range rs.rels {
			if !yield(rel) {
				return
			}
		}
	}
}

// Count returns the number of relationships.
func (rs *Relationships) Count() int {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	return len(rs.rels)
}

// IsEmpty returns true if there are no relationships.
func (rs *Relationships) IsEmpty() bool {
	return rs.Count() == 0
}

// XML types for .rels file serialization

const relationshipsNamespace = "http://schemas.openxmlformats.org/package/2006/relationships"

type xmlRelationships struct {
	XMLName       xml.Name          `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationships"`
	Relationships []xmlRelationship `xml:"Relationship"`
}

type xmlRelationship struct {
	ID         string `xml:"Id,attr"`
	Type       string `xml:"Type,attr"`
	Target     string `xml:"Target,attr"`
	TargetMode string `xml:"TargetMode,attr,omitempty"`
}

// MarshalToXML serializes the Relationships to XML.
func (rs *Relationships) MarshalToXML() ([]byte, error) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	xmlRels := xmlRelationships{}

	for _, rel := range rs.rels {
		xmlRel := xmlRelationship{
			ID:     rel.id,
			Type:   rel.relType,
			Target: rel.target,
		}
		if rel.targetMode == TargetModeExternal {
			xmlRel.TargetMode = "External"
		}
		xmlRels.Relationships = append(
			xmlRels.Relationships,
			xmlRel,
		)
	}

	var buf bytes.Buffer
	buf.WriteString(xml.Header)

	encoder := xml.NewEncoder(&buf)
	encoder.Indent("", "  ")
	if err := encoder.Encode(xmlRels); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnmarshalFromXML deserializes the Relationships from XML.
func (rs *Relationships) UnmarshalFromXML(
	r io.Reader,
) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	var xmlRels xmlRelationships
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(&xmlRels); err != nil {
		return err
	}

	// Clear existing and repopulate
	rs.rels = make(map[string]*Relationship)
	rs.nextID = 1

	for _, xmlRel := range xmlRels.Relationships {
		mode := TargetModeInternal
		if xmlRel.TargetMode == "External" {
			mode = TargetModeExternal
		}

		rel := NewRelationship(
			xmlRel.ID,
			xmlRel.Type,
			xmlRel.Target,
			mode,
		)
		rs.rels[xmlRel.ID] = rel

		// Track highest ID for auto-generation
		var num int
		if _, err := fmt.Sscanf(xmlRel.ID, "rId%d", &num); err == nil {
			if num >= rs.nextID {
				rs.nextID = num + 1
			}
		}
	}

	return nil
}
