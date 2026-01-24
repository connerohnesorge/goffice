//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package presentation provides PresentationML support for PowerPoint documents.
//
// This package implements the document-level API for creating, reading, and
// modifying .pptx files (and related formats like .potx, .ppsx, .pptm, .potm,
// .ppsm, .ppam).
//
// # Creating Documents
//
// Create a new presentation at a file path:
//
//	doc, err := presentation.New("presentation.pptx", presentation.DocTypePresentation)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer doc.Close()
//
// Create different document types:
//
//	doc, _ := presentation.New("template.potx", presentation.DocTypeTemplate)
//	doc, _ := presentation.New("show.ppsx", presentation.DocTypeSlideshow)
//	doc, _ := presentation.New("macro.pptm", presentation.DocTypeMacroEnabledPresentation)
//
// Write to an io.Writer:
//
//	var buf bytes.Buffer
//	doc, err := presentation.NewWriter(&buf, presentation.DocTypePresentation)
//
// # Opening Documents
//
// Open an existing presentation:
//
//	doc, err := presentation.Open("presentation.pptx", true) // editable=true
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer doc.Close()
//
// Open read-only:
//
//	doc, err := presentation.Open("presentation.pptx", false)
//
// Open from an io.ReaderAt:
//
//	doc, err := presentation.OpenReader(reader, size, true)
//
// # Accessing Document Content
//
// Access the main presentation part:
//
//	presPart := doc.PresentationPart()
//
// # Saving Documents
//
// Save to the original location:
//
//	err := doc.Save()
//
// Save to a new location:
//
//	err := doc.SaveAs("copy.pptx")
//
// # Document Types
//
// The package supports different PowerPoint document types:
//
//   - DocTypePresentation: Standard presentation (.pptx)
//   - DocTypeTemplate: Template (.potx)
//   - DocTypeSlideshow: Slideshow (.ppsx)
//   - DocTypeMacroEnabledPresentation: Macro-enabled presentation (.pptm)
//   - DocTypeMacroEnabledTemplate: Macro-enabled template (.potm)
//   - DocTypeMacroEnabledSlideshow: Macro-enabled slideshow (.ppsm)
//   - DocTypeAddIn: Add-in (.ppam)
//
// Change document type:
//
//	err := doc.ChangeType(presentation.DocTypeSlideshow)
//
// # Document Settings
//
// Configure document open settings:
//
//	settings := &presentation.OpenSettings{
//		AutoSave: true,
//		MaxCharactersInPart: 0, // unlimited
//	}
//	doc, err := presentation.OpenWithSettings("doc.pptx", true, settings)
//
// # Thread Safety
//
// Document operations are protected by sync.RWMutex. Read operations can
// proceed concurrently, while write operations require exclusive access.
//
// # Subpackages
//
// The presentation package is organized into subpackages:
//
//   - presentation/elements: PresentationML element types
//   - presentation/parts: Document part types (PresentationPart, SlidePart, etc.)
//
// See those packages for lower-level element manipulation.
//
//nolint:revive // line-length-limit: documentation examples contain long code samples
package presentation
