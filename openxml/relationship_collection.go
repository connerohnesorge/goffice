package openxml

import (
	"errors"
	"sort"
	"sync"
)

// RelationshipCollection manages a collection of relationships.
type RelationshipCollection interface {
	// Add adds a new relationship to the collection.
	Add(relType RelationshipType, target RelationshipTarget) (OpenXmlRelationship, error)

	// Remove removes a relationship by ID.
	Remove(id string) error

	// GetByID returns a relationship by ID.
	GetByID(id string) (OpenXmlRelationship, error)

	// GetByType returns all relationships of the given type.
	GetByType(relType RelationshipType) []OpenXmlRelationship

	// All returns all relationships in the collection.
	All() []OpenXmlRelationship
}

type relationshipCollection struct {
	mu        sync.RWMutex
	rels      map[string]OpenXmlRelationship
	container OpenXmlPartContainer
	idGen     *RelationshipIDGenerator
}

// NewRelationshipCollection creates a new relationship collection.
func NewRelationshipCollection(container OpenXmlPartContainer) RelationshipCollection {
	return &relationshipCollection{
		rels:      make(map[string]OpenXmlRelationship),
		container: container,
		idGen:     NewRelationshipIDGenerator(),
	}
}

func (rc *relationshipCollection) Add(relType RelationshipType, target RelationshipTarget) (OpenXmlRelationship, error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	id := rc.idGen.Next()

	var rel OpenXmlRelationship

	if target.IsExternal() {
		rel = NewExternalRelationship(id, string(relType), target.URI(), rc.container)
	} else {
		internalTarget, ok := target.(*InternalTarget)
		if !ok {
			return nil, errors.New("internal target must be of type *InternalTarget")
		}
		rel = NewPartRelationship(id, string(relType), internalTarget.Part, rc.container)
	}

	rc.rels[id] = rel

	return rel, nil
}

func (rc *relationshipCollection) Remove(id string) error {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	if _, ok := rc.rels[id]; !ok {
		return errors.New("relationship not found")
	}

	delete(rc.rels, id)

	return nil
}

func (rc *relationshipCollection) GetByID(id string) (OpenXmlRelationship, error) {
	rc.mu.RLock()
	defer rc.mu.RUnlock()

	rel, ok := rc.rels[id]
	if !ok {
		return nil, errors.New("relationship not found")
	}

	return rel, nil
}

func (rc *relationshipCollection) GetByType(relType RelationshipType) []OpenXmlRelationship {
	rc.mu.RLock()
	defer rc.mu.RUnlock()

	var result []OpenXmlRelationship
	for _, rel := range rc.rels {
		if rel.Type() == string(relType) {
			result = append(result, rel)
		}
	}
	// Sort by ID for deterministic output
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID() < result[j].ID()
	})

	return result
}

func (rc *relationshipCollection) All() []OpenXmlRelationship {
	rc.mu.RLock()
	defer rc.mu.RUnlock()

	result := make([]OpenXmlRelationship, 0, len(rc.rels))
	for _, rel := range rc.rels {
		result = append(result, rel)
	}
	// Sort by ID for deterministic output
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID() < result[j].ID()
	})

	return result
}
