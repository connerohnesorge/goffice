# JSON Schema Format Research

## Overview

The Open-XML-SDK uses JSON schema files to define all element types, attributes, and validation rules. These files are used by the code generator to produce strongly-typed classes. This research is shared across all document types (Word, Excel, PowerPoint).

## Schema File Structure

Located in `/data/schemas/*.json` (155 files):

```json
{
  "TargetNamespace": "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
  "Types": [
    // Array of SchemaType objects
  ]
}
```

## SchemaType Object Fields

### Basic Properties

- `Name`: QName format "prefix:ClassName/elementName"
- `ClassName`: C# class name
- `Summary`: Documentation string
- `Version`: Office version (Office2010, Office2013, Microsoft365, etc.)

### Type Classification

- `BaseClass`: Parent class (OpenXmlCompositeElement, OpenXmlLeafElement)
- `CompositeType`: Content model (`OneSequence`, `OneChoice`, `OneAll`)
- `IsLeafElement`: Boolean, true if element has no children
- `IsAbstract`: Boolean, true for abstract base types

### Structural Properties

- `Attributes`: Array of attribute definitions
- `Children`: Array of child element definitions
- `Particle`: Content model describing child ordering

## Particle Structure

```json
{
  "Kind": "Sequence",
  "Items": [
    {
      "Name": "w:CT_Body/w:body",
      "InitialVersion": "Office2007",
      "Occurs": [{ "Min": 1, "Max": 1 }]
    },
    {
      "Kind": "Choice",
      "Occurs": [{ "Max": 1 }],
      "Items": [/* nested particles */]
    }
  ]
}
```

### Particle.Kind Values

| Kind | Description |
|------|-------------|
| `Sequence` | Children must appear in specified order |
| `Choice` | Only one of the children can appear |
| `Group` | Logical grouping of particles |
| `All` | Children can appear in any order |

### Occurs Array Format

| Format | Meaning |
|--------|---------|
| `{}` | Required, unbounded (1..infinity) |
| `{"Max": 1}` | Optional, max 1 (0..1) |
| `{"Min": 0, "Max": 1}` | Explicitly optional |
| `{"Min": 1}` | Required, unbounded |
| `{"Min": 2, "Max": 5}` | 2 to 5 occurrences |

## Attribute Object Structure

```json
{
  "QName": ":attributeName",
  "PropertyName": "PropertyName",
  "Type": "StringValue",
  "PropertyComments": "documentation",
  "Version": "Microsoft365",
  "Validators": [
    {
      "Name": "RequiredValidator"
    },
    {
      "Name": "StringValidator",
      "Arguments": [
        {"Type": "Boolean", "Name": "IsToken", "Value": "True"},
        {"Type": "String", "Name": "Pattern", "Value": "\\{[0-9A-F]{8}...\\}"}
      ]
    }
  ]
}
```

## Validator Types

| Validator | Arguments | Purpose |
|-----------|-----------|---------|
| `RequiredValidator` | None | Attribute is required |
| `StringValidator` | IsToken, Pattern, MaxLength, MinLength | String value validation |
| `NumberValidator` | MinInclusive, MaxInclusive | Numeric range validation |
| `EnumValidator` | (implied by Type) | Enumerated value validation |
| `OfficeVersionValidator` | Value (Version) | Minimum Office version |

## Attribute Value Types

| Type | XML Schema Equivalent |
|------|----------------------|
| `StringValue` | xs:string |
| `BooleanValue` | xs:boolean |
| `Int32Value` | xs:int |
| `Int64Value` | xs:long |
| `UInt32Value` | xs:unsignedInt |
| `DoubleValue` | xs:double |
| `DecimalValue` | xs:decimal |
| `DateTimeValue` | xs:dateTime |
| `HexBinaryValue` | xs:hexBinary |
| `EnumValue<ClassName>` | enumerated type |

## Namespace Mapping

`/data/namespaces.json`:

```json
[
  {
    "Prefix": "w",
    "Uri": "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
    "Version": "Office2007"
  }
]
```

### WordprocessingML-specific Namespaces

| Prefix | Description |
|--------|-------------|
| `w` | Base wordprocessingML (2006) |
| `w14` | Word 2010 |
| `w15` | Word 2013 |
| `w16` | Word 2016 |
| `w16cex`, `w16cid`, `w16sdtdh`, etc. | Later Word versions |
| `wp` | DrawingML WordprocessingML Drawing |
| `wpc` | Drawing Canvas |
| `wpg` | Drawing Group |
| `wps` | Drawing Shape |

## Parts Structure

`/data/parts/*.json`:

```json
{
  "Root": "w:CT_Document",
  "Name": "MainDocumentPart",
  "Base": "OpenXmlPart",
  "RelationshipType": "http://schemas.openxmlformats.org/.../officeDocument",
  "Target": "document",
  "RootElement": "document",
  "Paths": { "General": "word" },
  "Children": [
    {
      "MaxOccursGreatThanOne": true,
      "ApiName": "HeaderParts",
      "Name": "HeaderPart",
      "HasFixedContent": true
    }
  ]
}
```

## Schematrons

`/data/schematrons.json`:

```json
{
  "Context": "w:control",
  "Test": "document(rels)//r:Relationship[@Id = current()/@w:id]/@Type = '...'",
  "App": "All"
}
```

XPath-based validation rules for complex constraints.

## Go Implementation Considerations

1. **Parse JSON schemas** to generate Go types
2. **Map particle types** to content model validation
3. **Implement validators** for each type
4. **Generate namespace constants** from namespaces.json
5. **Create part type definitions** from parts/*.json
