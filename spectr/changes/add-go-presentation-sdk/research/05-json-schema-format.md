# JSON Schema Format Research

## Overview

The Open-XML-SDK uses JSON schema files to define all element types, attributes, and validation rules. These files are used by the code generator to produce strongly-typed classes.

## Schema File Structure

Located in `/data/schemas/*.json` (155 files):

```json
{
  "TargetNamespace": "http://schemas.microsoft.com/office/powerpoint/2022/08/main",
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
- `Version`: Office version (Office2016, Office2019, Office2021, Microsoft365 - Go SDK supports 2016+ only)

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
      "Name": "p228:CT_TaskHistory/p228:history",
      "InitialVersion": "Microsoft365",
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
| `{}` | Required, unbounded (1..∞) |
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
    "Prefix": "p228",
    "Uri": "http://schemas.microsoft.com/office/powerpoint/2022/08/main",
    "Version": "Microsoft365"
  }
]
```

### Presentation-specific Namespaces

| Prefix | Description |
|--------|-------------|
| `p` | Base presentation (2006) |
| `p14` | PowerPoint 2010 |
| `p15` | PowerPoint 2013 |
| `p16` | PowerPoint 2015/2016 |
| `p173`, `p184`, `p188`, etc. | Later versions |

## Parts Structure

`/data/parts/*.json`:

```json
{
  "Root": "p:CT_Presentation",
  "Name": "PresentationPart",
  "Base": "OpenXmlPart",
  "RelationshipType": "http://schemas.openxmlformats.org/.../officeDocument",
  "Target": "presentation",
  "RootElement": "presentation",
  "Paths": { "General": "ppt" },
  "Children": [
    {
      "MaxOccursGreatThanOne": true,
      "ApiName": "SlideParts",
      "Name": "SlidePart",
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
