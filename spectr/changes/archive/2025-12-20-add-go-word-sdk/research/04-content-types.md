# Content Types System Research

## Overview

Content types in OpenXML identify the MIME type of each part in the package. The `[Content_Types].xml` file maps extensions and specific parts to their content types.

## [Content_Types].xml Structure

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
  <Default Extension="xml" ContentType="application/xml"/>
  <Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
  <Override PartName="/word/document.xml"
            ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"/>
</Types>
```

## Default vs Override Entries

### Default Entries

- Applied based on file extension
- Used for extensible parts (media, custom XML, etc.)
- Defined in `PartExtensionProvider.cs` with known MIME types

### Override Entries

- Applied to specific part URIs
- Used for fixed content type parts (document.xml, styles.xml, etc.)
- Each part type defines a `ContentTypeConstant` (immutable)

## PartExtensionProvider

```csharp
class PartExtensionProvider : Dictionary<string, string>, IPartExtensionFeature
{
    // Maps content types to file extensions
    // Pre-registers ~50+ known extensions on initialization

    public void Register(string contentType, string extension);
    public bool TryGetExtension(string contentType, out string extension);
}
```

### Pre-registered Extensions

| Category | Extensions |
|----------|------------|
| Images | `.bmp`, `.gif`, `.png`, `.jpg`, `.jpeg`, `.tif`, `.tiff`, `.svg`, `.emf`, `.wmf`, `.ico` |
| Audio | `.aiff`, `.midi`, `.mp3`, `.wav`, `.wma`, `.mpeg`, `.ogg` |
| Video | `.asx`, `.avi`, `.mp4`, `.mpg`, `.wmv`, `.wmx`, `.wvx`, `.mov` |

## Word Document Content Types

### Main Document Types

| Content Type | Extension | Description |
|--------------|-----------|-------------|
| `application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml` | .docx | Standard document |
| `application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml` | .dotx | Document template |
| `application/vnd.ms-word.document.macroEnabled.main+xml` | .docm | Macro-enabled document |
| `application/vnd.ms-word.template.macroEnabledTemplate.main+xml` | .dotm | Macro-enabled template |

### Word Part Types (Core Types)

| Part | Content Type |
|------|--------------|
| MainDocumentPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml` |
| StyleDefinitionsPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml` |
| NumberingDefinitionsPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml` |
| HeaderPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml` |
| FooterPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml` |
| FootnotesPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.footnotes+xml` |
| EndnotesPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.endnotes+xml` |
| CommentsPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.comments+xml` |
| DocumentSettingsPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml` |
| FontTablePart | `application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml` |
| WebSettingsPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml` |
| ThemePart | `application/vnd.openxmlformats-officedocument.theme+xml` |
| GlossaryDocumentPart | `application/vnd.openxmlformats-officedocument.wordprocessingml.document.glossary+xml` |

### Additional Word Part Types

| Part | Content Type |
|------|--------------|
| CustomXmlPart | `application/xml` or specific custom type |
| ImagePart | `image/png`, `image/jpeg`, etc. |
| EmbeddedPackagePart | `application/vnd.openxmlformats-officedocument.oleObject` |
| ChartPart | `application/vnd.openxmlformats-officedocument.drawingml.chart+xml` |
| DiagramPart | Various diagram content types |

## IContentTypeFeature Interface

```csharp
internal interface IContentTypeFeature
{
    bool IsConstant { get; }
}
```

- Indicates whether a part has a fixed/constant content type
- Fixed parts (document.xml, styles.xml): Content type never changes
- Flexible parts (CustomXmlPart, ImagePart): Content type is extensible

## Part Constraints Validation

```csharp
struct PartConstraintRule
{
    string RelationshipType;
    string ContentType;
    bool MinOccursIsNonZero;    // Part is required
    bool MaxOccursGreatThanOne; // Multiple instances allowed
    FileFormatVersions Version; // Office version support
}
```

## IPartFactoryFeature

Generated code that creates parts based on relationship type:

```csharp
// Word-specific relationship types -> part instance mappings via switch expression
OpenXmlPart? Create(string relationshipType) => relationshipType switch
{
    "http://schemas.openxmlformats.org/.../styles" => new StyleDefinitionsPart(),
    "http://schemas.openxmlformats.org/.../numbering" => new NumberingDefinitionsPart(),
    "http://schemas.openxmlformats.org/.../header" => new HeaderPart(),
    "http://schemas.openxmlformats.org/.../footer" => new FooterPart(),
    // ...
    _ => null
};
```

## Go Implementation Considerations

1. **Define all content types as constants**
2. **Implement part constraints** with relationship/content type/cardinality rules
3. **Create a content type registry** similar to PartExtensionProvider
4. **Generate part types** with fixed ContentType fields
5. **Support fixed vs flexible** content type parts
