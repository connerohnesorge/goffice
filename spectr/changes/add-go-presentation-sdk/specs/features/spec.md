# Spec: Features

## Overview

Implements the feature collection pattern for extensibility in the OpenXML SDK. Features provide a way to attach services and capabilities to elements, parts, and packages without tight coupling.

## API

### Types

```go
// IFeatureCollection provides access to a collection of features
type IFeatureCollection interface {
    // Get retrieves a feature by type, returns nil if not found
    Get(featureType reflect.Type) interface{}

    // Set registers a feature by type
    Set(featureType reflect.Type, feature interface{})

    // IsReadOnly returns true if the collection cannot be modified
    IsReadOnly() bool
}

// FeatureCollection is the default implementation of IFeatureCollection
type FeatureCollection struct {
    features map[reflect.Type]interface{}
    parent   IFeatureCollection
    readOnly bool
}
```

### Generic Helpers

```go
// GetFeature retrieves a typed feature from the collection
func GetFeature[T any](fc IFeatureCollection) T {
    result := fc.Get(reflect.TypeOf((*T)(nil)).Elem())
    if result != nil {
        return result.(T)
    }
    var zero T
    return zero
}

// SetFeature registers a typed feature in the collection
func SetFeature[T any](fc IFeatureCollection, feature T) {
    fc.Set(reflect.TypeOf((*T)(nil)).Elem(), feature)
}
```

### Common Features

```go
// IPartRootFeature provides access to the root element of a part
type IPartRootFeature interface {
    Root() OpenXmlElement
    SetRoot(element OpenXmlElement)
}

// IPackageEventsFeature provides package lifecycle events
type IPackageEventsFeature interface {
    OnOpening(handler func())
    OnClosing(handler func())
    OnSaving(handler func())
}

// IPartEventsFeature provides part lifecycle events
type IPartEventsFeature interface {
    OnLoading(handler func())
    OnLoaded(handler func())
    OnChanged(handler func())
}

// IElementEventsFeature provides element change notifications
type IElementEventsFeature interface {
    OnChildAdded(handler func(child OpenXmlElement))
    OnChildRemoved(handler func(child OpenXmlElement))
    OnAttributeChanged(handler func(attr OpenXmlAttribute))
}
```

## Implementation Notes

- Feature collections form a hierarchy: Element -> Part -> Package
- When a feature is not found locally, the parent collection is searched
- Features enable dependency injection without requiring constructor changes
- The read-only flag prevents modification of default/shared features
- Features are used extensively for validation, event handling, and customization
- Consider using sync.RWMutex for thread-safe access in concurrent scenarios
