// Package features provides the feature collection infrastructure for Office Open XML.
// Features allow parts and elements to access shared services like relationships,
// namespaces, and document-level settings.
package features

import (
	"reflect"
	"sync"
)

// FeatureCollection is a hierarchical container for features.
// Features are stored by their interface type and can be inherited from parent collections.
// This allows Element features to inherit from Part features, which inherit from Package features.
//
// The collection is thread-safe for concurrent reads and exclusive writes.
type FeatureCollection struct {
	mu       sync.RWMutex
	parent   *FeatureCollection
	features map[reflect.Type]Feature
}

// NewFeatureCollection creates a new empty feature collection.
func NewFeatureCollection() *FeatureCollection {
	return &FeatureCollection{
		features: make(map[reflect.Type]Feature),
	}
}

// NewFeatureCollectionWithParent creates a new feature collection with the given parent.
// Features not found in this collection will be looked up in the parent chain.
func NewFeatureCollectionWithParent(
	parent *FeatureCollection,
) *FeatureCollection {
	return &FeatureCollection{
		parent:   parent,
		features: make(map[reflect.Type]Feature),
	}
}

// Parent returns the parent feature collection, or nil if this is a root collection.
func (fc *FeatureCollection) Parent() *FeatureCollection {
	fc.mu.RLock()
	defer fc.mu.RUnlock()

	return fc.parent
}

// SetParent sets the parent feature collection.
func (fc *FeatureCollection) SetParent(
	parent *FeatureCollection,
) {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	fc.parent = parent
}

// Set registers a feature in the collection.
// The feature is stored by its concrete type. If a feature of the same type
// already exists, it is replaced.
func (fc *FeatureCollection) Set(
	feature Feature,
) {
	if feature == nil {
		return
	}

	fc.mu.Lock()
	defer fc.mu.Unlock()

	t := reflect.TypeOf(feature)
	fc.features[t] = feature
}

// SetByType registers a feature under a specific interface type.
// This allows registering a feature that should be retrieved by its interface type.
func (fc *FeatureCollection) SetByType(
	featureType reflect.Type,
	feature Feature,
) {
	if feature == nil {
		return
	}

	fc.mu.Lock()
	defer fc.mu.Unlock()

	fc.features[featureType] = feature
}

// get returns a feature by type, searching the parent chain if not found locally.
// Returns nil if not found.
func (fc *FeatureCollection) get(
	t reflect.Type,
) Feature {
	fc.mu.RLock()
	feature, ok := fc.features[t]
	parent := fc.parent
	fc.mu.RUnlock()

	if ok {
		return feature
	}

	if parent != nil {
		return parent.get(t)
	}

	return nil
}

// getLocal returns a feature by type from this collection only (no parent chain).
func (fc *FeatureCollection) getLocal(
	t reflect.Type,
) Feature {
	fc.mu.RLock()
	defer fc.mu.RUnlock()

	return fc.features[t]
}

// Remove removes a feature from the collection by its type.
// Returns true if a feature was removed.
func (fc *FeatureCollection) Remove(
	featureType reflect.Type,
) bool {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	if _, ok := fc.features[featureType]; ok {
		delete(fc.features, featureType)

		return true
	}

	return false
}

// Has returns true if the collection (or parent chain) contains a feature of the given type.
func (fc *FeatureCollection) Has(
	featureType reflect.Type,
) bool {
	return fc.get(featureType) != nil
}

// HasLocal returns true if this collection directly contains a feature of the given type.
func (fc *FeatureCollection) HasLocal(
	featureType reflect.Type,
) bool {
	return fc.getLocal(featureType) != nil
}

// Get retrieves a feature by its type T.
// The type T should be an interface type (e.g., IPackageFeature).
// Returns nil if not found in this collection or any parent.
func Get[T Feature](fc *FeatureCollection) T {
	if fc == nil {
		var zero T

		return zero
	}

	t := reflect.TypeFor[T]()
	feature := fc.get(t)

	if feature == nil {
		var zero T

		return zero
	}

	if typed, ok := feature.(T); ok {
		return typed
	}

	var zero T

	return zero
}

// GetLocal retrieves a feature by its type T from this collection only (no parent chain).
func GetLocal[T Feature](
	fc *FeatureCollection,
) T {
	if fc == nil {
		var zero T

		return zero
	}

	t := reflect.TypeFor[T]()
	feature := fc.getLocal(t)

	if feature == nil {
		var zero T

		return zero
	}

	if typed, ok := feature.(T); ok {
		return typed
	}

	var zero T

	return zero
}

// GetRequired retrieves a feature by its type T, panicking if not found.
// Use this when a feature is known to be required.
func GetRequired[T Feature](
	fc *FeatureCollection,
) T {
	result := Get[T](fc)
	if reflect.ValueOf(result).IsNil() {
		t := reflect.TypeFor[T]()
		panic(
			"required feature not found: " + t.String(),
		)
	}

	return result
}
