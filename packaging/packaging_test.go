package packaging

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// Test Package struct and capability

func TestPackageCapability(t *testing.T) {
	tests := []struct {
		name string
		cap  PackageCapability
		want string
	}{
		{"Read", Read, "Read"},
		{"Write", Write, "Write"},
		{"ReadWrite", ReadWrite, "ReadWrite"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cap.String(); got != tt.want {
				t.Errorf("PackageCapability.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreatePackage(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.docx")

	pkg, err := Create(path)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer pkg.Close()

	if pkg.Capability() != ReadWrite {
		t.Errorf("Capability() = %v, want ReadWrite", pkg.Capability())
	}

	if pkg.Path() != path {
		t.Errorf("Path() = %v, want %v", pkg.Path(), path)
	}

	if pkg.IsClosed() {
		t.Error("IsClosed() = true, want false")
	}
}

func TestCreateWriter(t *testing.T) {
	var buf bytes.Buffer

	pkg, err := CreateWriter(&buf)
	if err != nil {
		t.Fatalf("CreateWriter() error = %v", err)
	}
	defer pkg.Close()

	if pkg.Capability() != Write {
		t.Errorf("Capability() = %v, want Write", pkg.Capability())
	}

	if pkg.Path() != "" {
		t.Errorf("Path() = %v, want empty", pkg.Path())
	}
}

func TestPackageCreateAndSave(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.docx")

	pkg, err := Create(path)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	// Create a part
	part, err := pkg.CreatePart("/word/document.xml", "application/xml")
	if err != nil {
		t.Fatalf("CreatePart() error = %v", err)
	}

	part.SetData([]byte("<document>Hello</document>"))

	// Save
	if err := pkg.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	pkg.Close()

	// Verify file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("Expected file to exist after save")
	}

	// Re-open and verify
	pkg2, err := Open(path, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer pkg2.Close()

	part2, err := pkg2.Part("/word/document.xml")
	if err != nil {
		t.Fatalf("Part() error = %v", err)
	}

	data := part2.GetData()
	if string(data) != "<document>Hello</document>" {
		t.Errorf("Part data = %q, want %q", string(data), "<document>Hello</document>")
	}
}

func TestPackageSaveAs(t *testing.T) {
	tmpDir := t.TempDir()
	path1 := filepath.Join(tmpDir, "test1.docx")
	path2 := filepath.Join(tmpDir, "test2.docx")

	pkg, err := Create(path1)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer pkg.Close()

	part, _ := pkg.CreatePart("/word/document.xml", "application/xml")
	part.SetData([]byte("<document>Test</document>"))

	if err := pkg.SaveAs(path2); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}

	// Verify both files exist (path1 shouldn't exist since we never called Save())
	if _, err := os.Stat(path2); os.IsNotExist(err) {
		t.Error("Expected path2 to exist after SaveAs")
	}
}

func TestPackageClose(t *testing.T) {
	pkg, err := Create(filepath.Join(t.TempDir(), "test.docx"))
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	if err := pkg.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	if !pkg.IsClosed() {
		t.Error("IsClosed() = false after Close()")
	}

	// Operations on closed package should fail
	_, err = pkg.CreatePart("/test.xml", "application/xml")
	if err != ErrPackageClosed {
		t.Errorf("CreatePart() on closed package error = %v, want ErrPackageClosed", err)
	}
}

func TestPackageReadOnly(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.docx")

	// Create a package first
	pkg, _ := Create(path)
	pkg.CreatePart("/word/document.xml", "application/xml")
	pkg.Save()
	pkg.Close()

	// Open read-only
	pkg2, err := Open(path, true)
	if err != nil {
		t.Fatalf("Open(readOnly=true) error = %v", err)
	}
	defer pkg2.Close()

	if pkg2.Capability() != Read {
		t.Errorf("Capability() = %v, want Read", pkg2.Capability())
	}

	// Write operations should fail
	_, err = pkg2.CreatePart("/test.xml", "application/xml")
	if err != ErrReadOnly {
		t.Errorf("CreatePart() on read-only package error = %v, want ErrReadOnly", err)
	}
}

// Test Parts

func TestCreatePart(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	part, err := pkg.CreatePart("/word/document.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml")
	if err != nil {
		t.Fatalf("CreatePart() error = %v", err)
	}

	if part.URI() != "/word/document.xml" {
		t.Errorf("URI() = %v, want /word/document.xml", part.URI())
	}

	if part.ContentType() != "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml" {
		t.Errorf("ContentType() = %v", part.ContentType())
	}
}

func TestCreatePartDuplicate(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	_, err := pkg.CreatePart("/word/document.xml", "application/xml")
	if err != nil {
		t.Fatalf("First CreatePart() error = %v", err)
	}

	_, err = pkg.CreatePart("/word/document.xml", "application/xml")
	if err != ErrPartExists {
		t.Errorf("Second CreatePart() error = %v, want ErrPartExists", err)
	}
}

func TestCreatePartInvalidURI(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	_, err := pkg.CreatePart("", "application/xml")
	if err != ErrInvalidURI {
		t.Errorf("CreatePart('') error = %v, want ErrInvalidURI", err)
	}
}

func TestGetPart(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	created, _ := pkg.CreatePart("/word/document.xml", "application/xml")
	created.SetData([]byte("test data"))

	retrieved, err := pkg.Part("/word/document.xml")
	if err != nil {
		t.Fatalf("Part() error = %v", err)
	}

	if string(retrieved.GetData()) != "test data" {
		t.Errorf("GetData() = %q, want %q", string(retrieved.GetData()), "test data")
	}
}

func TestGetPartNotFound(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	_, err := pkg.Part("/nonexistent.xml")
	if err != ErrPartNotFound {
		t.Errorf("Part() error = %v, want ErrPartNotFound", err)
	}
}

func TestDeletePart(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	pkg.CreatePart("/word/document.xml", "application/xml")

	if err := pkg.DeletePart("/word/document.xml"); err != nil {
		t.Fatalf("DeletePart() error = %v", err)
	}

	_, err := pkg.Part("/word/document.xml")
	if err != ErrPartNotFound {
		t.Errorf("Part() after delete error = %v, want ErrPartNotFound", err)
	}
}

func TestDeletePartNotFound(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	err := pkg.DeletePart("/nonexistent.xml")
	if err != ErrPartNotFound {
		t.Errorf("DeletePart() error = %v, want ErrPartNotFound", err)
	}
}

func TestPartsIterator(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	pkg.CreatePart("/word/document.xml", "application/xml")
	pkg.CreatePart("/word/styles.xml", "application/xml")
	pkg.CreatePart("/word/settings.xml", "application/xml")

	count := 0
	for range pkg.Parts() {
		count++
	}

	if count != 3 {
		t.Errorf("Parts() count = %d, want 3", count)
	}
}

func TestPartStream(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	part, _ := pkg.CreatePart("/word/document.xml", "application/xml")

	// Set via stream
	part.SetStream(strings.NewReader("stream data"))

	// Get via stream
	reader := part.GetStream()
	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}

	if string(data) != "stream data" {
		t.Errorf("Stream data = %q, want %q", string(data), "stream data")
	}
}

// Test URI functions

func TestNormalizeURI(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"/word/document.xml", "/word/document.xml"},
		{"word/document.xml", "/word/document.xml"},
		{"/word//document.xml", "/word/document.xml"},
		{"\\word\\document.xml", "/word/document.xml"},
		{"/word/../styles.xml", "/styles.xml"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := NormalizeURI(tt.input); got != tt.want {
				t.Errorf("NormalizeURI(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestValidateURI(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"/word/document.xml", true},
		{"/", false},
		{"", false},
		{"word/document.xml", false},
		{"/word/document.xml/", false},
		{"/word//document.xml", false},
		{"/word/document.xml?query", false},
		{"/word/document.xml#fragment", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ValidateURI(tt.input); got != tt.want {
				t.Errorf("ValidateURI(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestResolvePartURI(t *testing.T) {
	tests := []struct {
		base     string
		relative string
		want     string
	}{
		{"/word/document.xml", "media/image1.png", "/word/media/image1.png"},
		{"/word/document.xml", "../media/image1.png", "/media/image1.png"},
		{"/word/document.xml", "/absolute/path.xml", "/absolute/path.xml"},
	}

	for _, tt := range tests {
		t.Run(tt.relative, func(t *testing.T) {
			if got := ResolvePartURI(tt.base, tt.relative); got != tt.want {
				t.Errorf("ResolvePartURI(%q, %q) = %q, want %q", tt.base, tt.relative, got, tt.want)
			}
		})
	}
}

func TestRelationshipPartURI(t *testing.T) {
	tests := []struct {
		partURI string
		want    string
	}{
		{"/word/document.xml", "/word/_rels/document.xml.rels"},
		{"/", "/_rels/.rels"},
		{"", "/_rels/.rels"},
	}

	for _, tt := range tests {
		t.Run(tt.partURI, func(t *testing.T) {
			if got := RelationshipPartURI(tt.partURI); got != tt.want {
				t.Errorf("RelationshipPartURI(%q) = %q, want %q", tt.partURI, got, tt.want)
			}
		})
	}
}

// Test Content Types

func TestContentTypesDefaults(t *testing.T) {
	ct := NewContentTypes()

	// Check standard defaults
	if got := ct.GetDefault("xml"); got != "application/xml" {
		t.Errorf("GetDefault(xml) = %q, want application/xml", got)
	}

	if got := ct.GetDefault("rels"); got != "application/vnd.openxmlformats-package.relationships+xml" {
		t.Errorf("GetDefault(rels) = %q", got)
	}
}

func TestContentTypesSetDefault(t *testing.T) {
	ct := NewContentTypes()

	ct.SetDefault("png", "image/png")

	if got := ct.GetDefault("png"); got != "image/png" {
		t.Errorf("GetDefault(png) = %q, want image/png", got)
	}

	// Test with dot prefix
	ct.SetDefault(".jpg", "image/jpeg")
	if got := ct.GetDefault("jpg"); got != "image/jpeg" {
		t.Errorf("GetDefault(jpg) = %q, want image/jpeg", got)
	}
}

func TestContentTypesOverride(t *testing.T) {
	ct := NewContentTypes()

	ct.SetOverride("/word/document.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml")

	contentType, err := ct.GetContentType("/word/document.xml")
	if err != nil {
		t.Fatalf("GetContentType() error = %v", err)
	}

	if contentType != "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml" {
		t.Errorf("GetContentType() = %q", contentType)
	}
}

func TestContentTypesFallbackToDefault(t *testing.T) {
	ct := NewContentTypes()

	contentType, err := ct.GetContentType("/word/styles.xml")
	if err != nil {
		t.Fatalf("GetContentType() error = %v", err)
	}

	if contentType != "application/xml" {
		t.Errorf("GetContentType() = %q, want application/xml", contentType)
	}
}

func TestContentTypesNotFound(t *testing.T) {
	ct := NewContentTypes()

	_, err := ct.GetContentType("/image.unknown")
	if err != ErrContentTypeNotFound {
		t.Errorf("GetContentType() error = %v, want ErrContentTypeNotFound", err)
	}
}

func TestContentTypesSerialization(t *testing.T) {
	ct := NewContentTypes()
	ct.SetDefault("png", "image/png")
	ct.SetOverride("/word/document.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml")

	data, err := ct.MarshalToXML()
	if err != nil {
		t.Fatalf("MarshalToXML() error = %v", err)
	}

	ct2 := NewContentTypes()
	if err := ct2.UnmarshalFromXML(bytes.NewReader(data)); err != nil {
		t.Fatalf("UnmarshalFromXML() error = %v", err)
	}

	if got := ct2.GetDefault("png"); got != "image/png" {
		t.Errorf("After roundtrip, GetDefault(png) = %q, want image/png", got)
	}

	contentType, _ := ct2.GetContentType("/word/document.xml")
	if contentType != "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml" {
		t.Errorf("After roundtrip, override not preserved")
	}
}

// Test Relationships

func TestCreateRelationship(t *testing.T) {
	rels := NewRelationships("/")

	rel, err := rels.Create("/word/document.xml", "http://example.com/relationship", "rId1")
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	if rel.ID() != "rId1" {
		t.Errorf("ID() = %q, want rId1", rel.ID())
	}

	if rel.Target() != "/word/document.xml" {
		t.Errorf("Target() = %q", rel.Target())
	}

	if rel.Type() != "http://example.com/relationship" {
		t.Errorf("Type() = %q", rel.Type())
	}

	if rel.TargetMode() != TargetModeInternal {
		t.Errorf("TargetMode() = %v, want Internal", rel.TargetMode())
	}
}

func TestCreateRelationshipAutoID(t *testing.T) {
	rels := NewRelationships("/")

	rel1, _ := rels.Create("/target1", "type", "")
	rel2, _ := rels.Create("/target2", "type", "")

	if rel1.ID() == rel2.ID() {
		t.Error("Auto-generated IDs should be unique")
	}
}

func TestCreateRelationshipDuplicate(t *testing.T) {
	rels := NewRelationships("/")

	_, _ = rels.Create("/target", "type", "rId1")
	_, err := rels.Create("/target2", "type", "rId1")

	if err != ErrRelationshipExists {
		t.Errorf("Create() with duplicate ID error = %v, want ErrRelationshipExists", err)
	}
}

func TestGetRelationship(t *testing.T) {
	rels := NewRelationships("/")
	rels.Create("/target", "type", "rId1")

	rel, err := rels.Get("rId1")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if rel.Target() != "/target" {
		t.Errorf("Target() = %q, want /target", rel.Target())
	}
}

func TestGetRelationshipNotFound(t *testing.T) {
	rels := NewRelationships("/")

	_, err := rels.Get("nonexistent")
	if err != ErrRelationshipNotFound {
		t.Errorf("Get() error = %v, want ErrRelationshipNotFound", err)
	}
}

func TestDeleteRelationship(t *testing.T) {
	rels := NewRelationships("/")
	rels.Create("/target", "type", "rId1")

	if err := rels.Delete("rId1"); err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	_, err := rels.Get("rId1")
	if err != ErrRelationshipNotFound {
		t.Errorf("Get() after delete error = %v, want ErrRelationshipNotFound", err)
	}
}

func TestRelationshipsByType(t *testing.T) {
	rels := NewRelationships("/")
	rels.Create("/target1", "type-a", "rId1")
	rels.Create("/target2", "type-b", "rId2")
	rels.Create("/target3", "type-a", "rId3")

	count := 0
	for range rels.ByType("type-a") {
		count++
	}

	if count != 2 {
		t.Errorf("ByType(type-a) count = %d, want 2", count)
	}
}

func TestRelationshipsSerialization(t *testing.T) {
	rels := NewRelationships("/")
	rels.Create("/word/document.xml", "http://example.com/document", "rId1")
	rels.CreateWithMode("http://example.com", "http://example.com/hyperlink", "rId2", TargetModeExternal)

	data, err := rels.MarshalToXML()
	if err != nil {
		t.Fatalf("MarshalToXML() error = %v", err)
	}

	rels2 := NewRelationships("/")
	if err := rels2.UnmarshalFromXML(bytes.NewReader(data)); err != nil {
		t.Fatalf("UnmarshalFromXML() error = %v", err)
	}

	rel1, _ := rels2.Get("rId1")
	if rel1.Target() != "/word/document.xml" {
		t.Errorf("After roundtrip, rId1 Target() = %q", rel1.Target())
	}

	rel2, _ := rels2.Get("rId2")
	if rel2.TargetMode() != TargetModeExternal {
		t.Errorf("After roundtrip, rId2 TargetMode() = %v, want External", rel2.TargetMode())
	}
}

// Test Core Properties

func TestCorePropertiesGetSet(t *testing.T) {
	cp := NewCoreProperties()

	cp.SetTitle("Test Document")
	if cp.Title() != "Test Document" {
		t.Errorf("Title() = %q, want Test Document", cp.Title())
	}

	cp.SetCreator("Test Author")
	if cp.Creator() != "Test Author" {
		t.Errorf("Creator() = %q, want Test Author", cp.Creator())
	}

	cp.SetSubject("Test Subject")
	if cp.Subject() != "Test Subject" {
		t.Errorf("Subject() = %q", cp.Subject())
	}

	cp.SetDescription("Test Description")
	if cp.Description() != "Test Description" {
		t.Errorf("Description() = %q", cp.Description())
	}

	cp.SetKeywords("test, keywords")
	if cp.Keywords() != "test, keywords" {
		t.Errorf("Keywords() = %q", cp.Keywords())
	}
}

func TestCorePropertiesDates(t *testing.T) {
	cp := NewCoreProperties()

	now := time.Now()
	cp.SetCreated(now)
	cp.SetModified(now)

	if cp.Created() == nil || !cp.Created().Equal(now) {
		t.Error("Created() not set correctly")
	}

	if cp.Modified() == nil || !cp.Modified().Equal(now) {
		t.Error("Modified() not set correctly")
	}
}

func TestCorePropertiesSerialization(t *testing.T) {
	cp := NewCoreProperties()
	cp.SetTitle("Test Title")
	cp.SetCreator("Test Author")
	now := time.Now().Truncate(time.Second) // Truncate for comparison
	cp.SetCreated(now)

	data, err := cp.MarshalToXML()
	if err != nil {
		t.Fatalf("MarshalToXML() error = %v", err)
	}

	cp2 := NewCoreProperties()
	if err := cp2.UnmarshalFromXML(bytes.NewReader(data)); err != nil {
		t.Fatalf("UnmarshalFromXML() error = %v", err)
	}

	if cp2.Title() != "Test Title" {
		t.Errorf("After roundtrip, Title() = %q", cp2.Title())
	}

	if cp2.Creator() != "Test Author" {
		t.Errorf("After roundtrip, Creator() = %q", cp2.Creator())
	}
}

// Test Package Relationships

func TestPackageCreateRelationship(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	rel, err := pkg.CreateRelationship("/word/document.xml", "http://example.com/document", "rId1")
	if err != nil {
		t.Fatalf("CreateRelationship() error = %v", err)
	}

	if rel.ID() != "rId1" {
		t.Errorf("ID() = %q, want rId1", rel.ID())
	}
}

func TestPackageRelationshipsByType(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	pkg.CreateRelationship("/word/document.xml", "http://example.com/document", "rId1")
	pkg.CreateRelationship("/word/styles.xml", "http://example.com/styles", "rId2")
	pkg.CreateRelationship("/word/settings.xml", "http://example.com/document", "rId3")

	count := 0
	for range pkg.RelationshipsByType("http://example.com/document") {
		count++
	}

	if count != 2 {
		t.Errorf("RelationshipsByType() count = %d, want 2", count)
	}
}

func TestPartRelationships(t *testing.T) {
	pkg, _ := Create(filepath.Join(t.TempDir(), "test.docx"))
	defer pkg.Close()

	part, _ := pkg.CreatePart("/word/document.xml", "application/xml")

	rel, err := part.CreateRelationship("media/image1.png", "http://example.com/image", "rId1")
	if err != nil {
		t.Fatalf("Part.CreateRelationship() error = %v", err)
	}

	if rel.ID() != "rId1" {
		t.Errorf("ID() = %q, want rId1", rel.ID())
	}

	rels := pkg.PartRelationships("/word/document.xml")
	if rels == nil {
		t.Fatal("PartRelationships() returned nil")
	}

	if rels.Count() != 1 {
		t.Errorf("Count() = %d, want 1", rels.Count())
	}
}

// Test Package Save/Load roundtrip

func TestPackageRoundtrip(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.docx")

	// Create package with parts and relationships
	pkg, _ := Create(path)

	// Add parts
	docPart, _ := pkg.CreatePart("/word/document.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml")
	docPart.SetData([]byte("<document><body>Hello World</body></document>"))

	stylesPart, _ := pkg.CreatePart("/word/styles.xml", "application/xml")
	stylesPart.SetData([]byte("<styles/>"))

	// Add package relationship
	pkg.CreateRelationship("/word/document.xml", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument", "rId1")

	// Add part relationship
	docPart.CreateRelationship("styles.xml", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles", "rId1")

	// Set core properties
	pkg.CoreProperties().SetTitle("Test Document")
	pkg.CoreProperties().SetCreator("Test Author")
	pkg.EnsureCorePropertiesPart()

	// Save
	if err := pkg.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	pkg.Close()

	// Reopen
	pkg2, err := Open(path, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer pkg2.Close()

	// Verify parts
	docPart2, err := pkg2.Part("/word/document.xml")
	if err != nil {
		t.Fatalf("Part() error = %v", err)
	}
	if string(docPart2.GetData()) != "<document><body>Hello World</body></document>" {
		t.Error("Document part data not preserved")
	}

	// Verify relationships
	count := 0
	for range pkg2.RelationshipsByType("http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument") {
		count++
	}
	if count != 1 {
		t.Errorf("Package relationships count = %d, want 1", count)
	}

	// Verify part relationships
	docRels := pkg2.PartRelationships("/word/document.xml")
	if docRels == nil || docRels.Count() != 1 {
		t.Error("Part relationships not preserved")
	}

	// Verify core properties
	if pkg2.CoreProperties().Title() != "Test Document" {
		t.Errorf("CoreProperties.Title() = %q, want Test Document", pkg2.CoreProperties().Title())
	}
}

// Test TargetMode

func TestTargetModeString(t *testing.T) {
	tests := []struct {
		mode TargetMode
		want string
	}{
		{TargetModeInternal, "Internal"},
		{TargetModeExternal, "External"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.mode.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}
