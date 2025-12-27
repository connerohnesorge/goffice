// Package testutil provides testing utilities for the goffice SDK.
package testutil

import (
	"archive/zip"
	"os"
	"path/filepath"
)

// dirPerm is the permission mode for directories.
const dirPerm = 0o755

// MinimalDocxFiles contains the minimal set of files required for a valid .docx
var MinimalDocxFiles = map[string]string{
	"[Content_Types].xml":          contentTypesXML,
	"_rels/.rels":                  relsXML,
	"word/document.xml":            documentXML,
	"word/_rels/document.xml.rels": documentRelsXML,
}

// XML namespace base.
const nsBase = "http://schemas.openxmlformats.org/"

// XML namespace constants.
var (
	nsContentTypes   = nsBase + "package/2006/content-types"
	nsRelationships  = nsBase + "package/2006/relationships"
	nsWordprocessing = nsBase + "wordprocessingml/2006/main"
	nsOfficeDocument = nsBase + "officeDocument/2006/relationships"
)

// Content type constants.
const (
	ctRelationships = "application/vnd.openxmlformats-package.relationships+xml"
	ctDocument      = "application/vnd.openxmlformats-officedocument." +
		"wordprocessingml.document.main+xml"
)

var contentTypesXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types xmlns="` + nsContentTypes + `">
  <Default Extension="rels" ContentType="` + ctRelationships + `"/>
  <Default Extension="xml" ContentType="application/xml"/>
  <Override PartName="/word/document.xml" ContentType="` + ctDocument + `"/>
</Types>`

var relsXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="` + nsRelationships + `">
  <Relationship Id="rId1" Type="` + nsOfficeDocument +
	`/officeDocument" Target="word/document.xml"/>
</Relationships>`

var documentXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document xmlns:w="` + nsWordprocessing + `">
  <w:body>
    <w:p>
      <w:r>
        <w:t>Hello, World!</w:t>
      </w:r>
    </w:p>
  </w:body>
</w:document>`

var documentRelsXML = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="` + nsRelationships + `">
</Relationships>`

// CreateMinimalDocx creates a minimal valid .docx file at the specified path.
func CreateMinimalDocx(path string) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return err
	}

	// Create the file
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	// Create ZIP writer
	zw := zip.NewWriter(f)

	// Write files in a specific order for consistency
	fileOrder := []string{
		"[Content_Types].xml",
		"_rels/.rels",
		"word/_rels/document.xml.rels",
		"word/document.xml",
	}

	for _, name := range fileOrder {
		content := MinimalDocxFiles[name]
		w, err := zw.Create(name)
		if err != nil {
			return err
		}
		if _, err := w.Write([]byte(content)); err != nil {
			return err
		}
	}

	return zw.Close()
}
