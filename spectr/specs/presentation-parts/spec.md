# Presentation Parts Specification

## Requirements

### Requirement: PresentationPart

The system SHALL provide `PresentationPart` as the main presentation content part.

#### Scenario: Presentation part creation
- WHEN a PresentationPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml`
- AND relationship type is `http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument`

#### Scenario: Presentation part root element
- WHEN accessing `Presentation()` root element
- THEN the p:presentation element is returned

#### Scenario: Slide parts access
- WHEN accessing `SlideParts()`
- THEN all child slide parts are returned

#### Scenario: Add slide part
- WHEN `AddSlidePart()` is called
- THEN a new SlidePart is created
- AND it is linked to the presentation

### Requirement: SlidePart

The system SHALL provide `SlidePart` for individual slide content.

#### Scenario: Slide part creation
- WHEN a SlidePart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.slide+xml`
- AND URI follows pattern `/ppt/slides/slideN.xml`

#### Scenario: Slide part root element
- WHEN accessing `Slide()` root element
- THEN the p:sld element is returned

#### Scenario: Slide layout reference
- WHEN accessing `SlideLayoutPart()`
- THEN the linked SlideLayoutPart is returned

#### Scenario: Add note slide
- WHEN `AddNotesSlidePart()` is called
- THEN a NotesSlidePart is created
- AND it is linked to this slide

### Requirement: SlideLayoutPart

The system SHALL provide `SlideLayoutPart` for slide layout definitions.

#### Scenario: Slide layout part creation
- WHEN a SlideLayoutPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml`
- AND URI follows pattern `/ppt/slideLayouts/slideLayoutN.xml`

#### Scenario: Slide layout root element
- WHEN accessing `SlideLayout()` root element
- THEN the p:sldLayout element is returned

#### Scenario: Slide master reference
- WHEN accessing `SlideMasterPart()`
- THEN the parent SlideMasterPart is returned

### Requirement: SlideMasterPart

The system SHALL provide `SlideMasterPart` for master slide definitions.

#### Scenario: Slide master part creation
- WHEN a SlideMasterPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml`
- AND URI follows pattern `/ppt/slideMasters/slideMasterN.xml`

#### Scenario: Slide master root element
- WHEN accessing `SlideMaster()` root element
- THEN the p:sldMaster element is returned

#### Scenario: Slide layouts access
- WHEN accessing `SlideLayoutParts()`
- THEN all child SlideLayoutParts are returned

#### Scenario: Theme part access
- WHEN accessing `ThemePart()`
- THEN the linked ThemePart is returned

### Requirement: NotesSlidePart

The system SHALL provide `NotesSlidePart` for speaker notes.

#### Scenario: Notes slide part creation
- WHEN a NotesSlidePart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.notesSlide+xml`
- AND URI follows pattern `/ppt/notesSlides/notesSlideN.xml`

#### Scenario: Notes slide root element
- WHEN accessing `NotesSlide()` root element
- THEN the p:notes element is returned

### Requirement: NotesMasterPart

The system SHALL provide `NotesMasterPart` for notes page master.

#### Scenario: Notes master part creation
- WHEN a NotesMasterPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.notesMaster+xml`

#### Scenario: Notes master root element
- WHEN accessing `NotesMaster()` root element
- THEN the p:notesMaster element is returned

### Requirement: HandoutMasterPart

The system SHALL provide `HandoutMasterPart` for handout page master.

#### Scenario: Handout master part creation
- WHEN a HandoutMasterPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.handoutMaster+xml`

#### Scenario: Handout master root element
- WHEN accessing `HandoutMaster()` root element
- THEN the p:handoutMaster element is returned

### Requirement: ThemePart

The system SHALL provide `ThemePart` for theme definitions.

#### Scenario: Theme part creation
- WHEN a ThemePart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.theme+xml`
- AND URI follows pattern `/ppt/theme/themeN.xml`

#### Scenario: Theme root element
- WHEN accessing `Theme()` root element
- THEN the a:theme element is returned

### Requirement: TableStylesPart

The system SHALL provide `TableStylesPart` for table style definitions.

#### Scenario: Table styles part creation
- WHEN a TableStylesPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.tableStyles+xml`

#### Scenario: Table styles root element
- WHEN accessing `TableStyleList()` root element
- THEN the a:tblStyleLst element is returned

### Requirement: Comment Parts

The system SHALL provide comment-related parts.

#### Scenario: Comment authors part
- WHEN a CommentAuthorsPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.commentAuthors+xml`

#### Scenario: Slide comments part
- WHEN a SlideCommentsPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.presentationml.comments+xml`

### Requirement: ChartPart

The system SHALL provide `ChartPart` for embedded charts.

#### Scenario: Chart part creation
- WHEN a ChartPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.drawingml.chart+xml`
- AND URI follows pattern `/ppt/charts/chartN.xml`

#### Scenario: Chart root element
- WHEN accessing `ChartSpace()` root element
- THEN the c:chartSpace element is returned

#### Scenario: Chart style parts
- WHEN accessing `ChartColorStylePart()` or `ChartStylePart()`
- THEN the respective style parts are returned

### Requirement: Diagram Parts

The system SHALL provide parts for SmartArt diagrams.

#### Scenario: Diagram data part
- WHEN a DiagramDataPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.drawingml.diagramData+xml`

#### Scenario: Diagram colors part
- WHEN a DiagramColorsPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.drawingml.diagramColors+xml`

#### Scenario: Diagram layout definition part
- WHEN a DiagramLayoutDefinitionPart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.drawingml.diagramLayout+xml`

#### Scenario: Diagram style part
- WHEN a DiagramStylePart is created
- THEN content type is `application/vnd.openxmlformats-officedocument.drawingml.diagramStyle+xml`

### Requirement: ImagePart

The system SHALL provide `ImagePart` for embedded images.

#### Scenario: Image part creation
- WHEN an ImagePart is created with content type
- THEN appropriate extension is used
- AND image content can be fed

#### Scenario: Image types support
- WHEN creating ImagePart
- THEN PNG, JPEG, GIF, TIFF, BMP, EMF, WMF, SVG are supported
- AND each has correct content type

### Requirement: Media Parts

The system SHALL provide parts for embedded audio and video.

#### Scenario: Audio part
- WHEN embedding audio content
- THEN AudioReferenceRelationship links to MediaDataPart

#### Scenario: Video part
- WHEN embedding video content
- THEN VideoReferenceRelationship links to MediaDataPart

### Requirement: Font Part

The system SHALL provide `FontPart` for embedded fonts.

#### Scenario: Font part creation
- WHEN a FontPart is created
- THEN content type is appropriate for font format
- AND URI follows pattern `/ppt/fonts/fontN.*`

### Requirement: Embedded Object Parts

The system SHALL provide parts for embedded objects.

#### Scenario: Embedded package part
- WHEN an EmbeddedPackagePart is created
- THEN external documents can be embedded
- AND various formats are supported (.xlsx, .docx, .pdf, etc.)

#### Scenario: OLE object part
- WHEN an OleObjectPart is created
- THEN OLE objects can be embedded

### Requirement: Custom XML Parts

The system SHALL provide custom XML storage.

#### Scenario: Custom XML part
- WHEN a CustomXmlPart is created
- THEN arbitrary XML can be stored
- AND it is associated with CustomXmlPropertiesPart

### Requirement: VBA Parts

The system SHALL provide VBA project storage.

#### Scenario: VBA project part
- WHEN a VbaProjectPart is created
- THEN macro-enabled content can be stored
- AND it is only valid in macro-enabled document types

### Requirement: View Properties Part

The system SHALL provide `ViewPropertiesPart` for view settings.

#### Scenario: View properties part
- WHEN a ViewPropertiesPart is created
- THEN presentation view settings can be stored

### Requirement: Presentation Properties Part

The system SHALL provide `PresentationPropertiesPart` for presentation settings.

#### Scenario: Presentation properties part
- WHEN a PresentationPropertiesPart is created
- THEN slideshow and other settings can be stored

