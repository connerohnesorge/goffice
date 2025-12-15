# Attribute System Research

## Overview

The Open-XML-SDK maintains a clear separation between schema-defined (fixed) attributes and unknown (extended) attributes.

## Fixed Attributes vs Extended Attributes

### Fixed Attributes

- Stored in `AttributeCollection` (array of `OpenXmlSimpleType[]`)
- Defined via `AttributeMetadata` objects in schema
- Lazy-initialized within `ElementState`
- Only parsed when accessed through `GetAttribute<T>()` or `ParsedState`
- Can be null (absence = attribute not set)

### Extended Attributes

- Stored in `List<OpenXmlAttribute>` within `MiscAttrContainer`
- Created on-demand (not in schema)
- Prefixed with namespace and local name
- Persist as `OpenXmlAttribute` struct containing: `QName`, `Prefix`, `Value`

## AttributeCollection Implementation

Located at: `Framework/Metadata/AttributeCollection.cs`

```csharp
readonly struct AttributeCollection
{
    private readonly OpenXmlSimpleType?[] _values;
    private readonly ReadOnlyArray<AttributeMetadata> _metadata;

    // Index-based lookup for fixed attributes
    // Linear search for metadata matching (O(n))
    // Returns AttributeEntry struct providing ref access to values
}
```

## GetAttribute/SetAttribute Mechanics

### GetAttribute Process

1. Checks `ParsedState.Attributes[qname]` for fixed attributes first
2. Falls back to `ExtendedAttributesField` for unknown attributes
3. Handles special Markup Compatibility (MC) attributes separately
4. Returns cloned `OpenXmlAttribute` (includes prefix resolution)

### SetAttribute Process

1. Tries `TrySetFixedAttribute()` for schema-defined attributes
2. If found: creates new attribute value object and sets `InnerText`
3. If not found: adds to `ExtendedAttributesField` list
4. Special handling for MC attributes via `LoadMCAttribute()`
5. Deduplication: removes old extended attribute before adding new one

## AttributeMetadata Structure

```csharp
abstract class AttributeMetadata
{
    string PropertyName { get; }
    OpenXmlQualifiedName QName { get; }
    ReadOnlyArray<IValidator> Validators { get; }
    Type Type { get; }
    OpenXmlSimpleType CreateNew();
}
```

- Builder pattern: `AttributeMetadata.Builder<TSimpleType>`
- Each attribute has associated validators
- QName consists of: namespace URI + local name

## Attribute Value Types

All attribute values stored as `OpenXmlSimpleType` subclasses:

| Type | Description |
|------|-------------|
| `StringValue` | Raw string storage |
| `EnumValue<T>` | Enumerated values |
| `Int32Value` | 32-bit integers |
| `Int64Value` | 64-bit integers |
| `UInt32Value` | Unsigned 32-bit integers |
| `DoubleValue` | Floating point |
| `DecimalValue` | Decimal numbers |
| `BooleanValue` | Boolean values |
| `DateTimeValue` | Date/time values |
| `HexBinaryValue` | Hexadecimal binary |
| `TrueFalseBlankValue` | Special boolean (true/false/blank) |
| `OnOffValue` | Special boolean (on/off/0/1) |

## OpenXmlAttribute Struct

```csharp
public struct OpenXmlAttribute
{
    public string LocalName { get; }
    public string NamespaceUri { get; }
    public string Prefix { get; }
    public string Value { get; set; }
}
```

## Namespace Handling

### Key Components

- `OpenXmlQualifiedName`: immutable struct with `Namespace` + `Name`
- `IOpenXmlNamespaceResolver`: provides prefix/URI resolution
- Prefix resolution: `LookupPrefix()` traverses element and ancestors

### Namespace Declarations

Stored as `List<KeyValuePair<string, string>>` where:
- Key = prefix
- Value = namespace URI

## Serialization

### Loading from XML

```csharp
LoadAttributes(XmlReader reader)
// For each attribute:
// 1. TrySetFixedAttribute() first
// 2. Unknown attributes → ExtendedAttributesField
// 3. Namespace declarations → NamespaceDeclField
// 4. MC attributes → MarkupCompatibilityAttributes
```

### Writing to XML

```csharp
WriteAttributesTo(XmlWriter writer)
// Order:
// 1. Namespace declarations
// 2. Fixed attributes
// 3. Extended attributes
// 4. MC attributes
```

## Go Implementation Considerations

1. **Separate fixed and extended attributes** into different storage
2. **Use generic value types** with implicit conversion methods
3. **Implement MiscAttrContainer** for lazy allocation
4. **Namespace lookup** must traverse parent chain
5. **Preserve attribute order** during serialization
