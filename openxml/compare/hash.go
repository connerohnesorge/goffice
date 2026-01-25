package compare

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/connerohnesorge/goffice/openxml"
)

// Hash computes a SHA-256 hash of the element's XML representation.
func Hash(el openxml.Element) string {
	if el == nil {
		return ""
	}
	xmlContent := el.OuterXml()
	hash := sha256.Sum256([]byte(xmlContent))
	return hex.EncodeToString(hash[:])
}
