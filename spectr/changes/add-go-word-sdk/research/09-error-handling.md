# Error Handling Patterns Research

## Overview

The Open-XML-SDK uses a combination of custom exceptions, standard .NET exceptions, and validation error objects for different error scenarios. This research is shared across all document types (Word, Excel, PowerPoint).

## Custom Exception Types

### OpenXmlPackageException

```csharp
public sealed class OpenXmlPackageException : Exception
```

- Package-level errors
- Invalid content types
- Document type changes
- Part relationship issues
- Serializable

### InvalidMCContentException

```csharp
public sealed class InvalidMCContentException : Exception
```

- Markup Compatibility content errors
- Invalid extensibility markup
- Unknown MC content
- Invalid MC attribute values

### NamespaceNotUnderstandException

```csharp
public sealed class NamespaceNotUnderstandException : Exception
```

- MC namespace prefix issues
- Current MC mode doesn't understand a namespace

## ExceptionMessages Resource

Located in `Resources/ExceptionMessages.resx` with 65+ localized messages.

### Package & Document Type Errors

| Key | Message |
|-----|---------|
| `CannotChangeDocumentType` | Cannot change document type |
| `CannotChangeDocumentTypeSerious` | Cannot change document type, may be corrupt |
| `DocumentTooBig` | Document exceeded size limit |
| `InvalidPackageType` | Specified package is not valid |

### Part Loading Errors

| Key | Message |
|-----|---------|
| `CannotLoadRootElement` | Cannot load root element from part |
| `InvalidContentTypePart` | Content type of part is invalid |
| `InvalidPartContentType` | Invalid part with unexpected content type |
| `Fmt_PartRootIsInvalid` | Root XML element is incorrect |

### XML & Element Errors

| Key | Message |
|-----|---------|
| `InvalidOuterXml` | XML has invalid content |
| `RootElementIsNull` | Root element is null |

### Markup Compatibility Errors

| Key | Message |
|-----|---------|
| `UnknowMCContent` | Invalid extensibility markup in element |
| `NsNotUnderStand` | MC mode doesn't understand namespace |

### Attribute Errors

| Key | Message |
|-----|---------|
| `TextIsInvalidEnumValue` | Invalid enumeration value |
| `TextIsInvalidOnOffValue` | Invalid (true/false/on/off/0/1 only) |
| `TextIsInvalidTrueFalseValue` | Invalid (true/false/t/f only) |

### Special Conditions

| Key | Message |
|-----|---------|
| `StrictReadOnly` | ISO 29500 Strict can't be opened with edit |
| `EncryptedPackageNotSupported` | Encrypted packages not supported |
| `MalformedUri` | Package contains malformed URI |

## SR.Format() Pattern

```csharp
throw new ExceptionType(SR.Format(ExceptionMessages.MessageKey, arg1, arg2));
```

Used ~457 times across 107 files for formatted error messages.

## Validation Errors vs Exceptions

### ValidationErrorInfo (Non-Exception)

```csharp
public class ValidationErrorInfo
{
    public string Id { get; set; }
    public ValidationErrorType ErrorType { get; set; }
    public string Description { get; set; }
    public OpenXmlElement Node { get; set; }
    public OpenXmlPart Part { get; set; }
    public string Path { get; set; }
    public OpenXmlElement RelatedNode { get; set; }
    public OpenXmlPart RelatedPart { get; set; }
}
```

### ValidationErrorType Enum

```csharp
public enum ValidationErrorType
{
    Schema,             // Structural/content model violations
    Semantic,           // Business logic constraints
    Package,            // Part structure/relationships
    MarkupCompatibility // MC attribute violations
}
```

### Key Differences

| Aspect | Exceptions | Validation Errors |
|--------|------------|-------------------|
| Behavior | Stop processing (fatal) | Collected, continue (non-fatal) |
| Use Case | Corruption, malformed data | Conformance checking |
| Collection | No | ValidationContext.Errors list |
| Throwing | Immediate | After validation completes |

## XML Parsing Errors

### XmlException Handling Pattern

```csharp
try
{
    XmlConvert.VerifyNCName(value);
}
catch (XmlException ex)
{
    // Convert to ValidationErrorInfo or ArgumentException
    throw new ArgumentException(SR.Format(...), ex);
}
```

Found in:
- `DocumentValidator.cs` - Converts to ValidationErrorInfo
- `QnameRestriction.cs` - QName parsing
- `StringValidator.cs` - IsValidNcName validation
- `OpenXmlPartContainer.cs` - XML ID verification

## Part Root Element Loading

```csharp
// In LoadFromPart()
if (!IsCorrectRootElement(reader))
{
    var message = SR.Format(
        ExceptionMessages.Fmt_PartRootIsInvalid,
        actualQName,
        expectedQName
    );
    throw new InvalidDataException(message);
}
```

## MC Context Error Handling

```csharp
bool OnMcContextError(string message)
{
    if (_noExceptionOnError)
    {
        return false; // Non-fatal
    }
    throw new InvalidMCContentException(message); // Fatal
}
```

Configurable: `_noExceptionOnError` flag controls behavior.

## Standard Exception Usage

| Exception | Usage |
|-----------|-------|
| `ArgumentException` | Invalid method arguments |
| `ArgumentNullException` | Null arguments |
| `ArgumentOutOfRangeException` | Out of range values |
| `InvalidOperationException` | Invalid state operations |
| `InvalidDataException` | Malformed part content |
| `KeyNotFoundException` | Missing attributes/parts |
| `NotSupportedException` | Unsupported operations |

## Go Implementation Considerations

1. **Define custom error types** for package, MC, and namespace errors
2. **Separate validation errors** from fatal exceptions
3. **Error collection** for non-fatal validation issues
4. **Formatted error messages** with context placeholders
5. **Wrap XML parsing errors** with additional context
6. **Configurable error handling** for MC processing
