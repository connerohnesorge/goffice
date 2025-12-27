# FlatOPC Format Research

## Overview

FlatOPC is a single-XML representation of an OPC (Open Packaging Convention) document. Instead of a ZIP file with multiple parts, all content is embedded in one XML document.

## Use Cases

- Source control-friendly representation
- Network transmission as a single file
- Dynamic generation without ZIP I/O
- Streaming and text-based processing

## XML Structure

### Root Element

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<?mso-application progid="PowerPoint.Presentation"?>
<pkg:package xmlns:pkg="http://schemas.microsoft.com/office/2006/xmlPackage">
  <!-- Parts -->
</pkg:package>
```

- Namespace: `http://schemas.microsoft.com/office/2006/xmlPackage`
- Prefix: `pkg`
- Optional processing instruction with `progid`

### Part Elements

```xml
<pkg:part pkg:name="/{path}/{filename}"
          pkg:contentType="application/vnd.openxmlformats-{type}"
          pkg:padding="{bytes}"
          pkg:compression="store">
  <pkg:xmlData><!-- for XML parts --></pkg:xmlData>
  <!-- OR -->
  <pkg:binaryData><!-- base64 for binary parts --></pkg:binaryData>
</pkg:part>
```

| Attribute | Description |
|-----------|-------------|
| `pkg:name` | Part URI (e.g., `/ppt/presentation.xml`) |
| `pkg:contentType` | MIME type of the part |
| `pkg:padding` | Optional padding size in bytes |
| `pkg:compression` | For binary: `store` (no compression) |

### XML Parts

```xml
<pkg:part pkg:name="/ppt/presentation.xml"
          pkg:contentType="application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml">
  <pkg:xmlData>
    <p:presentation xmlns:p="...">
      <!-- Actual document content -->
    </p:presentation>
  </pkg:xmlData>
</pkg:part>
```

### Binary Parts

```xml
<pkg:part pkg:name="/ppt/media/image1.png"
          pkg:contentType="image/png"
          pkg:compression="store">
  <pkg:binaryData>
    iVBORw0KGgoAAAANSUhEUgAA...<!-- base64 encoded -->
  </pkg:binaryData>
</pkg:part>
```

### Relationships

```xml
<pkg:part pkg:name="/_rels/.rels"
          pkg:contentType="application/vnd.openxmlformats-package.relationships+xml"
          pkg:padding="512">
  <pkg:xmlData>
    <Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
      <Relationship Id="rId1" Type="...officeDocument" Target="ppt/presentation.xml"/>
    </Relationships>
  </pkg:xmlData>
</pkg:part>
```

## Conversion Process

### To FlatOPC

1. Save current package state
2. Identify AlternativeFormatInputParts (AltChunk) - always binary
3. Create root `pkg:package` element
4. For each package part:
   - If XML content type AND not AltChunk:
     - Load as XML
     - Create `pkg:part` with `pkg:xmlData`
   - Otherwise:
     - Read binary data
     - Base64 encode with line breaks
     - Create `pkg:part` with `pkg:binaryData`
5. Create XDocument with declaration

### From FlatOPC

1. Create empty package
2. For each `pkg:part` element:
   - Extract name and content type
   - If contains `pkg:xmlData`:
     - Create XML part
     - Write XML content
   - If contains `pkg:binaryData`:
     - Base64 decode
     - Write binary content
3. Reload package

## API Methods

### PresentationDocument

```csharp
// From XDocument
public static PresentationDocument FromFlatOpcDocument(XDocument document);
public static PresentationDocument FromFlatOpcDocument(XDocument document, Stream stream, bool isEditable);
public static PresentationDocument FromFlatOpcDocument(XDocument document, string path, bool isEditable);

// From String
public static PresentationDocument FromFlatOpcString(string text);
public static PresentationDocument FromFlatOpcString(string text, Stream stream, bool isEditable);
public static PresentationDocument FromFlatOpcString(string text, string path, bool isEditable);

// To FlatOPC
public static string ToFlatOpcString(this OpenXmlPackage package);
public static XDocument ToFlatOpcDocument(this OpenXmlPackage openXmlPackage);
```

## Binary Data Encoding

- Method: Base64 with line breaks
- Uses `Convert.ToBase64String` with `InsertLineBreaks`
- Prevents excessively long XML lines
- Decoded back to bytes on load

## Special Handling

### AlternativeFormatInputParts (AltChunk)

- Relationship type: `http://schemas.openxmlformats.org/.../aFChunk`
- Always treated as binary in FlatOPC
- Important for embedded documents

### Padding Attribute

- Typical values: 512 (root rels), 256 (part rels)
- Indicates original ZIP entry padding
- Can be ignored during recreation

## Go Implementation Considerations

1. **Support dual output formats**: ZIP-based OPC and FlatOPC XML
2. **Implement Base64 chunking** for readable binary data
3. **Part type detection** based on content type suffix
4. **Relationship parts** are XML but stored separately
5. **MSO application hint** via processing instruction
6. **Stream-based conversion** for large documents
