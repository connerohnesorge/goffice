# Content Types System Research

## Overview

Content types in OpenXML identify the MIME type of each part in the package. The `[Content_Types].xml` file maps extensions and specific parts to their content types.

## [Content_Types].xml Structure

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
  <Default Extension="xml" ContentType="application/xml"/>
  <Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
  <Override PartName="/ppt/presentation.xml"
            ContentType="application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml"/>
</Types>
```

## Default vs Override Entries

### Default Entries

- Applied based on file extension
- Used for extensible parts (media, custom XML, etc.)
- Defined in `PartExtensionProvider.cs` with known MIME types

### Override Entries

- Applied to specific part URIs
- Used for fixed content type parts (presentation.xml, slide.xml, etc.)
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

## Presentation Content Types

### Main Document Types

| Content Type | Extension |
|--------------|-----------|
| `application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml` | .pptx |
| `application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml` | .ppsx |
| `application/vnd.openxmlformats-officedocument.presentationml.template.main+xml` | .potx |

### Part Types (14 core types)

| Part | Content Type |
|------|--------------|
| SlidePart | `application/vnd.openxmlformats-officedocument.presentationml.slide+xml` |
| SlideLayoutPart | `application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml` |
| SlideMasterPart | `application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml` |
| HandoutMasterPart | `application/vnd.openxmlformats-officedocument.presentationml.handoutMaster+xml` |
| NotesMasterPart | `application/vnd.openxmlformats-officedocument.presentationml.notesMaster+xml` |
| NotesSlidePart | `application/vnd.openxmlformats-officedocument.presentationml.notesSlide+xml` |
| SlideCommentsPart | `application/vnd.openxmlformats-officedocument.presentationml.comments+xml` |
| CommentAuthorsPart | `application/vnd.openxmlformats-officedocument.presentationml.commentAuthors+xml` |
| PresentationPropertiesPart | `application/vnd.openxmlformats-officedocument.presentationml.presProps+xml` |
| ViewPropertiesPart | `application/vnd.openxmlformats-officedocument.presentationml.viewProps+xml` |
| TableStylesPart | `application/vnd.openxmlformats-officedocument.presentationml.tableStyles+xml` |
| ThemePart | `application/vnd.openxmlformats-officedocument.theme+xml` |

## IContentTypeFeature Interface

```csharp
internal interface IContentTypeFeature
{
    bool IsConstant { get; }
}
```

- Indicates whether a part has a fixed/constant content type
- Fixed parts (presentation.xml, slide.xml): Content type never changes
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
// 90+ relationship types → part instance mappings via switch expression
OpenXmlPart? Create(string relationshipType) => relationshipType switch
{
    "http://schemas.openxmlformats.org/.../slide" => new SlidePart(),
    "http://schemas.openxmlformats.org/.../slideLayout" => new SlideLayoutPart(),
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
