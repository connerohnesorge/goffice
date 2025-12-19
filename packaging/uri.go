// Package packaging provides the OPC (Open Packaging Conventions) layer.
package packaging

import (
	"path"
	"strings"
)

// NormalizeURI normalizes a part URI according to OPC conventions.
// It ensures the URI starts with a forward slash and uses forward slashes throughout.
// The URI is also lowercased for case-insensitive comparison.
func NormalizeURI(uri string) string {
	// Replace backslashes with forward slashes
	uri = strings.ReplaceAll(uri, "\\", "/")

	// Ensure leading slash
	if !strings.HasPrefix(uri, "/") {
		uri = "/" + uri
	}

	// Clean the path (removes double slashes, resolves . and ..)
	uri = path.Clean(uri)

	// Ensure leading slash after clean (Clean removes it for root)
	if !strings.HasPrefix(uri, "/") {
		uri = "/" + uri
	}

	return uri
}

// ValidateURI validates that a string is a valid OPC part URI.
// A valid part URI:
// - Must not be empty
// - Must start with a forward slash
// - Must not end with a forward slash
// - Must not contain empty segments
// - Must not be only a forward slash
// - Must not contain query strings or fragments
func ValidateURI(uri string) bool {
	if uri == "" {
		return false
	}

	// Must start with /
	if !strings.HasPrefix(uri, "/") {
		return false
	}

	// Must not be just /
	if uri == "/" {
		return false
	}

	// Must not end with /
	if strings.HasSuffix(uri, "/") {
		return false
	}

	// Must not contain query string or fragment
	if strings.Contains(uri, "?") ||
		strings.Contains(uri, "#") {
		return false
	}

	// Check for empty segments (double slashes)
	segments := strings.Split(
		uri[1:],
		"/",
	) // Skip leading slash
	for _, seg := range segments {
		if seg == "" {
			return false
		}
	}

	return true
}

// ResolvePartURI resolves a relative URI against a base part URI.
// The base URI is the URI of the source part, and relative is the target reference.
// For example, ResolvePartURI("/word/document.xml", "media/image1.png") returns "/word/media/image1.png".
func ResolvePartURI(
	base, relative string,
) string {
	// If relative is already absolute, just normalize it
	if strings.HasPrefix(relative, "/") {
		return NormalizeURI(relative)
	}

	// Get the directory of the base URI
	baseDir := path.Dir(base)

	// Join with the relative path
	resolved := path.Join(baseDir, relative)

	// Ensure leading slash
	if !strings.HasPrefix(resolved, "/") {
		resolved = "/" + resolved
	}

	return resolved
}

// URIExtension returns the file extension from a part URI.
// The extension includes the leading dot, e.g., ".xml".
// Returns empty string if no extension is present.
func URIExtension(uri string) string {
	return strings.ToLower(path.Ext(uri))
}

// URIDirectory returns the directory portion of a part URI.
// For "/word/document.xml", this returns "/word".
func URIDirectory(uri string) string {
	dir := path.Dir(uri)
	if dir == "." {
		return "/"
	}
	return dir
}

// URIFilename returns the filename portion of a part URI.
// For "/word/document.xml", this returns "document.xml".
func URIFilename(uri string) string {
	return path.Base(uri)
}

// IsRelationshipURI checks if the URI is a relationship part URI.
// Relationship URIs are in _rels directories and end with .rels.
func IsRelationshipURI(uri string) bool {
	filename := path.Base(uri)
	dir := path.Dir(uri)
	return strings.HasSuffix(filename, ".rels") &&
		strings.HasSuffix(dir, "_rels")
}

// RelationshipPartURI returns the relationship part URI for a given part URI.
// For "/word/document.xml", this returns "/word/_rels/document.xml.rels".
// For package-level relationships, pass "/" to get "/_rels/.rels".
func RelationshipPartURI(partURI string) string {
	if partURI == "/" || partURI == "" {
		return "/_rels/.rels"
	}

	dir := path.Dir(partURI)
	filename := path.Base(partURI)

	return path.Join(
		dir,
		"_rels",
		filename+".rels",
	)
}

// PartURIFromRelationshipURI returns the source part URI for a relationship part.
// For "/word/_rels/document.xml.rels", this returns "/word/document.xml".
// For "/_rels/.rels", this returns "/".
func PartURIFromRelationshipURI(
	relURI string,
) string {
	if relURI == "/_rels/.rels" {
		return "/"
	}

	// Remove the .rels extension
	base := strings.TrimSuffix(relURI, ".rels")

	// Get directory and filename
	dir := path.Dir(base)
	filename := path.Base(base)

	// Remove _rels from directory
	parentDir := path.Dir(dir)

	return path.Join(parentDir, filename)
}
