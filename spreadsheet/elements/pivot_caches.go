package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PivotCaches represents the pivot caches container element (x:pivotCaches).
type PivotCaches struct {
	*openxml.CompositeElementBase
}

// NewPivotCaches creates a new PivotCaches element.
func NewPivotCaches() *PivotCaches {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"pivotCaches",
		PrefixDefault,
	)

	return &PivotCaches{
		CompositeElementBase: elem,
	}
}

// PivotCaches returns an iterator over all PivotCache elements.
func (pc *PivotCaches) PivotCaches() iter.Seq[*PivotCache] {
	return func(yield func(*PivotCache) bool) {
		for child := range pc.Children() {
			if child.LocalName() != "pivotCache" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var cache *PivotCache
			switch v := child.(type) {
			case *PivotCache:
				cache = v
			case *openxml.CompositeElementBase:
				cache = &PivotCache{CompositeElementBase: v}
			}
			if cache != nil && !yield(cache) {
				return
			}
		}
	}
}

// Count returns the number of pivot caches.
func (pc *PivotCaches) Count() int {
	count := 0
	for range pc.PivotCaches() {
		count++
	}

	return count
}

// GetByCacheId returns the pivot cache with the given cache ID.
func (pc *PivotCaches) GetByCacheId(
	cacheId int,
) *PivotCache {
	for cache := range pc.PivotCaches() {
		if cache.CacheId() == cacheId {
			return cache
		}
	}

	return nil
}

// AddPivotCache adds a new pivot cache with the given cache ID and
// relationship ID.
func (pc *PivotCaches) AddPivotCache(
	cacheId int, relationshipId string,
) *PivotCache {
	cache := NewPivotCache()
	cache.SetCacheId(cacheId)
	cache.SetRelationshipId(relationshipId)
	pc.AppendChild(cache)

	return cache
}

// RemovePivotCache removes a pivot cache from the collection.
func (pc *PivotCaches) RemovePivotCache(
	cache *PivotCache,
) bool {
	return pc.RemoveChild(cache)
}

// NextCacheId returns the next available cache ID.
func (pc *PivotCaches) NextCacheId() int {
	maxId := 0
	for cache := range pc.PivotCaches() {
		if cache.CacheId() > maxId {
			maxId = cache.CacheId()
		}
	}

	return maxId + 1
}

// Clone creates a deep copy of this PivotCaches element.
func (pc *PivotCaches) Clone() openxml.Element {
	cloned := pc.CompositeElementBase.Clone()

	return &PivotCaches{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PivotCaches element.
func (pc *PivotCaches) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pc.CompositeElementBase.CloneNode(
		deep,
	)

	return &PivotCaches{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// PivotCache represents a single pivot cache element (x:pivotCache).
// This is a reference to a pivot cache definition part.
type PivotCache struct {
	*openxml.CompositeElementBase
}

// NewPivotCache creates a new PivotCache element.
func NewPivotCache() *PivotCache {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"pivotCache",
		PrefixDefault,
	)

	return &PivotCache{CompositeElementBase: elem}
}

// CacheId returns the unique cache ID.
func (c *PivotCache) CacheId() int {
	attr, found := c.GetAttribute("cacheId", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCacheId sets the unique cache ID.
func (c *PivotCache) SetCacheId(id int) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"cacheId",
			"",
			strconv.Itoa(id),
		),
	)
}

// RelationshipId returns the relationship ID linking to the pivot cache
// definition part.
func (c *PivotCache) RelationshipId() string {
	attr, found := c.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID.
func (c *PivotCache) SetRelationshipId(
	id string,
) {
	c.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"id",
			PrefixR,
			id,
		),
	)
}

// Clone creates a deep copy of this PivotCache element.
func (c *PivotCache) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &PivotCache{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PivotCache element.
func (c *PivotCache) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &PivotCache{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
