//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package packaging provides the OPC (Open Packaging Conventions) layer
// for Office Open XML document handling.
//
// OPC defines how Office Open XML documents (such as .docx, .xlsx, .pptx) are
// stored as ZIP-based packages containing XML parts, relationships, and content
// types. This package implements the low-level package operations required by
// higher-level document APIs.
//
// # Package Structure
//
// An OPC package consists of:
//
//   - Parts: Individual files within the ZIP archive (XML documents, images, etc.)
//   - Relationships: Links between parts that define the document structure
//   - Content Types: MIME types that identify the format of each part
//   - Core Properties: Standard metadata like author, title, and dates
//
// # Creating Packages
//
// Use Create to create a new package at a file path:
//
//	pkg, err := packaging.Create("document.docx")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer pkg.Close()
//
// Use CreateWriter to create a package that writes to an io.Writer:
//
//	var buf bytes.Buffer
//	pkg, err := packaging.CreateWriter(&buf)
//
// # Opening Packages
//
// Use Open to open an existing package file:
//
//	pkg, err := packaging.Open("document.docx", true) // readOnly=true
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer pkg.Close()
//
// Use OpenReader to open from an io.ReaderAt:
//
//	pkg, err := packaging.OpenReader(reader, size)
//
// # Working with Parts
//
// Create parts to store document content:
//
//	part, err := pkg.CreatePart("/word/document.xml", contentType)
//	if err != nil {
//		log.Fatal(err)
//	}
//	part.SetData(xmlContent)
//
// Retrieve existing parts:
//
//	part, err := pkg.Part("/word/document.xml")
//	data := part.GetData()
//
// Iterate over all parts:
//
//	for part := range pkg.Parts() {
//		fmt.Println(part.URI(), part.ContentType())
//	}
//
// # Relationships
//
// Relationships define connections between parts. Package-level relationships
// link the package to its main parts:
//
//	rel, err := pkg.CreateRelationship(
//		"/word/document.xml",
//		"http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument",
//		"rId1",
//	)
//
// Part-level relationships link one part to related parts:
//
//	rel, err := pkg.CreatePartRelationship(
//		"/word/document.xml",                              // source
//		"styles.xml",                                       // target (relative)
//		"http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles",
//		"rId2",
//	)
//
// # Content Types
//
// The package automatically manages content types, but you can access them:
//
//	ct := pkg.ContentTypes()
//	ct.SetOverride("/word/document.xml", contentType)
//	ct.SetDefault("xml", "application/xml")
//
// # Core Properties
//
// Access standard document metadata:
//
//	props := pkg.CoreProperties()
//	props.SetTitle("My Document")
//	props.SetCreator("John Doe")
//	props.SetCreated(time.Now())
//
// # Saving Changes
//
// Save changes to the original file:
//
//	err := pkg.Save()
//
// Save to a new location:
//
//	err := pkg.SaveAs("copy.docx")
//
// # Thread Safety
//
// Package operations are protected by sync.RWMutex for safe concurrent access.
// Read operations can proceed concurrently, while write operations require
// exclusive access.
//
//nolint:revive // max-public-structs: OPC requires multiple public types
package packaging
