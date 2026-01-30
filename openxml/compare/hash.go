// Package compare provides utilities for comparing and merging OpenXML elements.
package compare

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/connerohnesorge/goffice/openxml"
)

// Hash computes a SHA-256 hash of the element's XML representation.
// Returns an empty string if the element is nil.
// The hash is computed from the element's OuterXml() output, providing
// a stable identifier for the element's current state.
func Hash(el openxml.Element) string {
	if el == nil {
		return ""
	}
	xmlContent := el.OuterXml()
	hash := sha256.Sum256([]byte(xmlContent))

	return hex.EncodeToString(hash[:])
}
