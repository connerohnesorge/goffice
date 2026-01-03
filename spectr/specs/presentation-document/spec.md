# Presentation Document Specification

## Requirements

### Requirement: PresentationDocument Type

The system SHALL provide `PresentationDocument` for creating and managing PowerPoint presentations.

#### Scenario: Document type property
- WHEN accessing `DocumentType()` on a presentation
- THEN the current PresentationDocumentType is returned
- AND the type reflects the content type

### Requirement: PresentationDocumentType Enumeration

The system SHALL support all PowerPoint document types.

#### Scenario: Presentation type (.pptx)
- WHEN `DocumentType()` is Presentation
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml`
- AND file extension is .pptx

#### Scenario: Template type (.potx)
- WHEN `DocumentType()` is Template
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.template.main+xml`
- AND file extension is .potx

#### Scenario: Slideshow type (.ppsx)
- WHEN `DocumentType()` is Slideshow
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml`
- AND file extension is .ppsx

#### Scenario: MacroEnabledPresentation type (.pptm)
- WHEN `DocumentType()` is MacroEnabledPresentation
- THEN content type is `application/vnd.ms-powerpoint.presentation.macroEnabled.main+xml`
- AND file extension is .pptm

#### Scenario: MacroEnabledTemplate type (.potm)
- WHEN `DocumentType()` is MacroEnabledTemplate
- THEN content type is `application/vnd.ms-powerpoint.template.macroEnabled.main+xml`
- AND file extension is .potm

#### Scenario: MacroEnabledSlideshow type (.ppsm)
- WHEN `DocumentType()` is MacroEnabledSlideshow
- THEN content type is `application/vnd.ms-powerpoint.slideshow.macroEnabled.main+xml`
- AND file extension is .ppsm

#### Scenario: AddIn type (.ppam)
- WHEN `DocumentType()` is AddIn
- THEN content type is `application/vnd.ms-powerpoint.addin.macroEnabled.main+xml`
- AND file extension is .ppam

### Requirement: Document Creation

The system SHALL provide multiple methods for creating new presentations.

#### Scenario: Create from file path
- WHEN `PresentationDocument.Create(path, documentType)` is called
- THEN a new file is created at the path
- AND the document type determines content type
- AND auto-save defaults to true

#### Scenario: Create from file path with auto-save option
- WHEN `PresentationDocument.Create(path, documentType, autoSave)` is called
- THEN auto-save behavior is set as specified

#### Scenario: Create from stream
- WHEN `PresentationDocument.Create(stream, documentType)` is called
- THEN the package is created in the stream
- AND the stream must be writable

#### Scenario: Create from template
- WHEN `PresentationDocument.CreateFromTemplate(templatePath)` is called
- THEN the template is loaded
- AND a new editable document is created
- AND changes don't affect the template

### Requirement: Document Opening

The system SHALL provide multiple methods for opening existing presentations.

#### Scenario: Open from file path read-only
- WHEN `PresentationDocument.Open(path, false)` is called
- THEN the document is opened read-only
- AND modifications are not allowed

#### Scenario: Open from file path editable
- WHEN `PresentationDocument.Open(path, true)` is called
- THEN the document is opened for editing
- AND changes can be saved

#### Scenario: Open from stream
- WHEN `PresentationDocument.Open(stream, isEditable)` is called
- THEN the document is loaded from stream
- AND editability depends on stream and flag

#### Scenario: Open with settings
- WHEN `PresentationDocument.Open(path, isEditable, openSettings)` is called
- THEN advanced settings are applied
- AND markup compatibility is processed per settings

### Requirement: Document Type Conversion

The system SHALL support converting between document types.

#### Scenario: Change document type
- WHEN `ChangeDocumentType(newType)` is called
- THEN the document type is changed
- AND the PresentationPart content type is updated
- AND the file can be saved with new extension

#### Scenario: Change to macro-enabled
- WHEN changing from Presentation to MacroEnabledPresentation
- THEN VBA content can be added
- AND content type changes appropriately

#### Scenario: Change from macro-enabled
- WHEN changing from MacroEnabledPresentation to Presentation
- THEN VBA content is preserved but not executable
- AND content type changes appropriately

### Requirement: Part Management

The system SHALL provide methods for adding and accessing standard parts.

#### Scenario: Add presentation part
- WHEN `AddPresentationPart()` is called
- THEN a new PresentationPart is created
- AND it becomes the main part of the document

#### Scenario: Access presentation part
- WHEN accessing `PresentationPart()` property
- THEN the main presentation part is returned
- AND nil if not yet created

#### Scenario: Add core file properties part
- WHEN `AddCoreFilePropertiesPart()` is called
- THEN core properties part is created
- AND document metadata can be set

#### Scenario: Add extended file properties part
- WHEN `AddExtendedFilePropertiesPart()` is called
- THEN extended properties part is created

#### Scenario: Add custom file properties part
- WHEN `AddCustomFilePropertiesPart()` is called
- THEN custom properties part is created

#### Scenario: Add digital signature origin part
- WHEN `AddDigitalSignatureOriginPart()` is called
- THEN signature origin part is created

#### Scenario: Add ribbon customization parts
- WHEN `AddRibbonExtensibilityPart()` is called
- THEN ribbon customization part is created
- WHEN `AddQuickAccessToolbarCustomizationsPart()` is called
- THEN QAT customization part is created

#### Scenario: Add web extensions taskpane part
- WHEN `AddWebExTaskpanesPart()` is called
- THEN web extensions taskpane part is created

### Requirement: FlatOPC Support

The system SHALL support FlatOPC (single-file XML) format.

#### Scenario: Convert to FlatOPC string
- WHEN `ToFlatOpcString()` is called
- THEN the entire package is serialized to XML string
- AND parts are base64 encoded inline

#### Scenario: Convert to FlatOPC document
- WHEN `ToFlatOpcDocument()` is called
- THEN an XML document representation is returned

#### Scenario: Create from FlatOPC string
- WHEN `PresentationDocument.FromFlatOpcString(xml)` is called
- THEN a presentation is created from the XML
- AND all parts are restored

#### Scenario: Create from FlatOPC document
- WHEN `PresentationDocument.FromFlatOpcDocument(xdoc)` is called
- THEN a presentation is created from the XML document

### Requirement: Package Features

The system SHALL provide presentation-specific package features.

#### Scenario: Application type feature
- WHEN accessing application type feature
- THEN it returns ApplicationType.PowerPoint

#### Scenario: Programmatic identifier feature
- WHEN accessing programmatic ID feature
- THEN it returns "PowerPoint.Show"

#### Scenario: Main part feature
- WHEN accessing main part feature
- THEN it provides relationship type and content type info

### Requirement: Default Theme Generation
New Presentation documents MUST contain a `theme1.xml` part to ensure valid visual styling.

#### Scenario: New presentation has theme
- WHEN a new `presentation.Document` is created via `New()`
- THEN document package MUST contain a `theme` part (e.g., `/ppt/theme/theme1.xml`)
- AND `SlideMaster` MUST reference this theme via a relationship
