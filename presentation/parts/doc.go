//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package parts provides PowerPoint presentation part types.
//
// Parts are the individual XML files within a .pptx package. Each part
// represents a distinct aspect of the presentation such as the main presentation,
// slides, slide layouts, slide masters, themes, and more.
//
// # Part Types
//
// The package provides the following part types:
//
// Presentation Content:
//   - PresentationPart: Main presentation content (ppt/presentation.xml)
//   - SlidePart: Individual slide content (ppt/slides/slide1.xml, etc.)
//   - SlideLayoutPart: Slide layout definitions (ppt/slideLayouts/slideLayout1.xml, etc.)
//   - SlideMasterPart: Slide master definitions (ppt/slideMasters/slideMaster1.xml, etc.)
//
// Notes:
//   - NotesSlidePart: Speaker notes for slides (ppt/notesSlides/notesSlide1.xml, etc.)
//   - NotesMasterPart: Notes master definition (ppt/notesMasters/notesMaster1.xml)
//   - HandoutMasterPart: Handout master definition (ppt/handoutMasters/handoutMaster1.xml)
//
// Themes and Styling:
//   - ThemePart: Theme definitions (ppt/theme/theme1.xml)
//
// Media:
//   - ImagePart: Embedded images (ppt/media/image1.png, etc.)
//
// Comments:
//   - CommentAuthorsPart: Comment authors (ppt/commentAuthors.xml)
//   - SlideCommentsPart: Slide comments (ppt/comments/comment1.xml, etc.)
//
// # PresentationPart
//
// The PresentationPart is the central part of any PowerPoint presentation:
//
//	presPart := doc.PresentationPart()
//	presentation := presPart.Presentation()
//
// Add supporting parts through PresentationPart:
//
//	slidePart, err := presPart.AddSlidePart()
//	themePart, err := presPart.AddThemePart()
//
// # Slides and Layouts
//
// Add slides and configure layouts:
//
//	slidePart, err := presPart.AddSlidePart()
//	slide := slidePart.Slide()
//
//	layoutPart, err := masterPart.AddSlideLayoutPart()
//	layout := layoutPart.SlideLayout()
//
// # Slide Masters
//
// Slide masters define the base formatting:
//
//	masterPart, err := presPart.AddSlideMasterPart()
//	master := masterPart.SlideMaster()
//
// # Speaker Notes
//
// Add speaker notes to slides:
//
//	notesPart, err := slidePart.AddNotesSlidePart()
//	notes := notesPart.NotesSlide()
//
// # Part URIs
//
// Standard part URIs follow OPC conventions:
//
//   - /ppt/presentation.xml - Main presentation
//   - /ppt/slides/slide1.xml - First slide
//   - /ppt/slideLayouts/slideLayout1.xml - First layout
//   - /ppt/slideMasters/slideMaster1.xml - First master
//   - /ppt/notesSlides/notesSlide1.xml - First notes slide
//   - /ppt/notesMasters/notesMaster1.xml - Notes master
//   - /ppt/handoutMasters/handoutMaster1.xml - Handout master
//   - /ppt/theme/theme1.xml - First theme
//   - /ppt/media/image1.png - First image
//
// # Content Types
//
// Each part type has a specific content type:
//
//	parts.ContentTypePresentation   // Main presentation
//	parts.ContentTypeSlide          // Slides
//	parts.ContentTypeSlideLayout    // Slide layouts
//	parts.ContentTypeSlideMaster    // Slide masters
//
// # Relationship Types
//
// Parts are connected via relationships:
//
//	parts.RelationshipTypeSlide         // Slides
//	parts.RelationshipTypeSlideLayout   // Slide layouts
//	parts.RelationshipTypeSlideMaster   // Slide masters
//	parts.RelationshipTypeTheme         // Themes
//
// # Part Registration
//
// All part types are automatically registered with the openxml package's
// part type registry via init() functions. This enables automatic part
// instantiation when opening existing presentations.
//
// # Thread Safety
//
// Part operations should be performed through the parent document's
// thread-safe API. Direct part manipulation is not thread-safe.
//
//nolint:revive // line-length-limit: documentation examples contain long method chains
package parts
