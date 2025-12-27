# OpenXmlElement System Research

## Overview

The OpenXmlElement system forms the foundation of the Open-XML-SDK. This document details the internal architecture based on exploration of the C# SDK.

## Element Hierarchy

```
OpenXmlElement (abstract)
  ├── OpenXmlCompositeElement (abstract) - elements that can contain children
  ├── OpenXmlLeafElement (abstract) - elements without children
  ├── OpenXmlUnknownElement - unknown/extension elements
  ├── OpenXmlMiscNode - comments, processing instructions, text nodes
  └── AlternateContent - markup compatibility wrapper
```

## Internal State Management

### Element State Fields

- `_rawOuterXml: string` - Raw unparsed XML content (lazy loading)
- `_state: ElementState` - Fixed attributes and metadata cache
- `Next: OpenXmlElement?` - Linked list pointer (circular for children)
- `_lastChild: OpenXmlElement?` (in CompositeElement) - Last child in circular linked list
- `Parent: OpenXmlElement?` - Parent element reference
- `MiscAttrContainer` - Container for extended attributes, MC attributes, and namespace declarations
- `_features: IFeatureCollection?` - Feature collection (lazy initialization)

### XmlParsed Property

```csharp
internal bool XmlParsed => string.IsNullOrEmpty(_rawOuterXml);
```
- Returns `true` if element has been parsed
- Returns `false` if element is in lazy-loaded state

## Child Element Management (Circular Linked List)

### Data Structure

OpenXmlCompositeElement uses a **circular singly-linked list** to manage children:

- Single `Next` pointer per element
- Single `_lastChild` pointer in OpenXmlCompositeElement
- **Circular**: When children exist, `_lastChild.Next` points to `FirstChild`
- **Empty case**: When no children, `_lastChild` is null

```
Empty element: _lastChild = null
Single child: _lastChild.Next = _lastChild (self-pointing circle)
Multiple children: _lastChild.Next points to FirstChild
```

### Child Manipulation Methods

| Method | Implementation |
|--------|----------------|
| **AppendChild** | O(1) - Direct pointer update to _lastChild |
| **PrependChild** | Delegates to InsertBefore(newChild, FirstChild) |
| **InsertBefore** | O(n) - Finds predecessor, updates pointers |
| **InsertAfter** | O(1) if at end, O(n) otherwise |
| **RemoveChild** | O(1) if known predecessor, O(n) for arbitrary |

### Navigation Methods

| Method | Complexity | Notes |
|--------|-----------|-------|
| FirstChild | O(1) | Direct `_lastChild.Next` |
| LastChild | O(1) | Direct field |
| NextSibling | O(1) | Direct `.Next` pointer |
| PreviousSibling | O(n) | Linear search from first |
| ChildElements.Count | O(n) | Iterates all children |
| ChildElements[i] | O(n) | Linear search |

## Attribute Management

### Storage Categories

1. **Fixed Attributes** (schema-defined): Stored in `ElementState.Attributes` collection
2. **Extended Attributes** (custom): Stored in `ExtendedAttributesField` list
3. **MC Attributes** (markup compatibility): Stored in `McAttributes`
4. **Namespace Declarations**: Stored in `NamespaceDeclField`

### MiscAttrContainer

```csharp
internal class MiscAttrContainer
{
    List<OpenXmlAttribute>? ExtendedAttributesField { get; set; }
    MarkupCompatibilityAttributes? McAttributes { get; set; }
    List<KeyValuePair<string, string>>? NsMappings { get; set; }
}
```

## Lazy Loading

### Dual State Model

- **Full Mode**: Complete parsing on load
- **Lazy Mode**: Parsing deferred until property access

```csharp
internal void MakeSureParsed()
{
    if (XmlParsed) return;
    ParseXml();
    RawOuterXml = string.Empty; // Mark as parsed
}
```

## Features & Metadata

### Feature Collection Pattern

```csharp
public IFeatureCollection Features
{
    get
    {
        if (_features is null)
        {
            _features = CreateFeatures();
        }
        return _features;
    }
}
```

### Metadata Access

```csharp
internal IElementMetadata Metadata => Features.GetRequired<IElementMetadata>();
```

## XML Writing

```csharp
public virtual void WriteTo(XmlWriter xmlWriter)
{
    if (XmlParsed)
    {
        // Parsed state: reconstruct XML
        WriteStartElement(prefix, LocalName, NamespaceUri);
        WriteAttributesTo(xmlWriter);
        WriteContentTo(xmlWriter);
        WriteEndElement();
    }
    else
    {
        // Lazy state: write raw XML directly
        xmlWriter.WriteRaw(RawOuterXml);
    }
}
```

## Go Implementation Considerations

1. **Use struct embedding** instead of inheritance
2. **Implement circular linked list** for children with `lastChild` and `Next` pointers
3. **Separate attribute storage** into fixed (schema) and extended (unknown)
4. **Support lazy parsing** with raw XML storage
5. **Use interface-based feature collection** for extensibility
