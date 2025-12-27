# OpenSettings and Configuration Research

## Overview

The Open-XML-SDK provides several configuration options for controlling document loading, saving, and validation behavior. This research is shared across all document types (Word, Excel, PowerPoint).

## OpenSettings Class

```csharp
public sealed class OpenSettings
{
    public bool AutoSave { get; set; } = true;
    public MarkupCompatibilityProcessSettings MarkupCompatibilityProcessSettings { get; set; }
    public long MaxCharactersInPart { get; set; } = 0;
    public CompatibilityLevel CompatibilityLevel { get; set; }
}
```

### AutoSave Property

- Default: `true`
- Controls auto-save on dispose
- When true, changes are saved when document is disposed
- When false, requires explicit Save() call

### MaxCharactersInPart Property

- Default: `0` (unlimited)
- Security limit on maximum characters in a part
- Prevents denial-of-service via large documents

## MarkupCompatibilityProcessSettings

```csharp
public class MarkupCompatibilityProcessSettings
{
    public MarkupCompatibilityProcessMode ProcessMode { get; set; }
    public FileFormatVersions TargetFileFormatVersions { get; set; }
}
```

### ProcessMode Values

| Value | Description |
|-------|-------------|
| `NoProcess` (0) | Do not process MC tags |
| `ProcessLoadedPartsOnly` | Process only loaded parts |
| `ProcessAllParts` | Process all parts in package |

## CompatibilityLevel Enum

```csharp
public enum CompatibilityLevel
{
    Default,
    Version_2_20,
    Version_3_0
}
```

Controls API compatibility behavior.

## FileFormatVersions Enum (Flags)

**C# SDK Original (for reference):**
```csharp
[Flags]
public enum FileFormatVersions
{
    None = 0,
    Office2007 = 0x1,
    Office2010 = 0x2,
    Office2013 = 0x4,
    Office2016 = 0x8,
    Office2019 = 0x10,
    Office2021 = 0x20,
    Microsoft365 = 0x40000000
}
```

**Go Implementation (Office 2016+ only):**
```go
// FileFormatVersion represents supported Office versions.
// This SDK targets Office 2016+ only (ECMA-376 5th edition and later).
type FileFormatVersion int

const (
    Office2016    FileFormatVersion = iota
    Office2019
    Office2021
    Microsoft365
)
```

**Note:** The Go SDK intentionally excludes Office 2007, Office 2010, and Office 2013 to reduce complexity and focus on modern document formats. This aligns with the design decision to support Office 2016+ only.

Used for:
- Targeting specific Office versions (2016+)
- Validation against version-specific schemas
- MC namespace understanding

## Document Open/Create Methods

### WordprocessingDocument.Open Overloads

```csharp
// From file path
public static WordprocessingDocument Open(string path, bool isEditable);
public static WordprocessingDocument Open(string path, bool isEditable, OpenSettings openSettings);

// From stream
public static WordprocessingDocument Open(Stream stream, bool isEditable);
public static WordprocessingDocument Open(Stream stream, bool isEditable, OpenSettings openSettings);

// From package
public static WordprocessingDocument Open(Package package);
public static WordprocessingDocument Open(Package package, OpenSettings openSettings);
```

### WordprocessingDocument.Create Overloads

```csharp
public static WordprocessingDocument Create(string path, WordprocessingDocumentType type);
public static WordprocessingDocument Create(string path, WordprocessingDocumentType type, bool autoSave);
public static WordprocessingDocument Create(Stream stream, WordprocessingDocumentType type);
public static WordprocessingDocument Create(Stream stream, WordprocessingDocumentType type, bool autoSave);
public static WordprocessingDocument Create(Package package, WordprocessingDocumentType type);
```

### WordprocessingDocumentType Enum

```csharp
public enum WordprocessingDocumentType
{
    Document,       // .docx
    Template,       // .dotx
    MacroEnabledDocument,    // .docm
    MacroEnabledTemplate     // .dotm
}
```

## PackageCapabilities (Feature Flags)

```csharp
[Flags]
public enum PackageCapabilities
{
    None = 0,
    Save = 0x1,           // Package can be saved
    Reload = 0x2,         // Package can be reloaded
    Cached = 0x4,         // Returns same part/relationship
    LargePartStreams = 0x8,  // Handles large streams
    MalformedUri = 0x10   // Handles malformed URIs
}
```

## OpenXmlPartWriterSettings

```csharp
public class OpenXmlPartWriterSettings
{
#if FEATURE_ASYNC_SAX_XML
    public bool Async { get; set; }
#endif
    public bool CloseOutput { get; set; }
    public Encoding Encoding { get; set; } = Encoding.UTF8;
}
```

## ValidationSettings

```csharp
public class ValidationSettings
{
    public FileFormatVersions FileFormat { get; set; }
    public int MaxNumberOfErrors { get; set; }
}
```

## Feature Interfaces

### ISaveFeature

```csharp
internal interface ISaveFeature
{
    void Save();
}
```

### IPackageFeature

```csharp
internal interface IPackageFeature
{
    IPackage Package { get; }
    PackageCapabilities Capabilities { get; }
}
```

### ILockFeature

```csharp
internal interface ILockFeature
{
    void AcquireLock();
    void ReleaseLock();
}
```

## Go Implementation Considerations

1. **Create OpenSettings struct** with all configuration fields
2. **Implement FileFormatVersions** as a flags type
3. **Support multiple Open/Create method signatures** with functional options
4. **Add MaxCharactersInPart security limit**
5. **Implement AutoSave behavior** in Close/Dispose
6. **Package capabilities** for feature detection
