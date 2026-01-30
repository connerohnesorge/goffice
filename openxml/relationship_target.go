package openxml

// RelationshipTarget represents the target of a relationship.
type RelationshipTarget interface {
	// URI returns the target URI.
	URI() string
	// IsExternal returns true if the target is external.
	IsExternal() bool
}

// InternalTarget represents a relationship to a part within the package.
type InternalTarget struct {
	Part OpenXmlPart
}

// NewInternalTarget creates a new internal target.
func NewInternalTarget(part OpenXmlPart) *InternalTarget {
	return &InternalTarget{Part: part}
}

// URI returns the part URI.
func (t *InternalTarget) URI() string {
	return t.Part.URI()
}

// IsExternal returns false for internal targets.
func (t *InternalTarget) IsExternal() bool {
	return false
}

// ExternalTarget represents a relationship to an external resource.
type ExternalTarget struct {
	URL string
}

// NewExternalTarget creates a new external target.
func NewExternalTarget(url string) *ExternalTarget {
	return &ExternalTarget{URL: url}
}

// URI returns the external URL.
func (t *ExternalTarget) URI() string {
	return t.URL
}

// IsExternal returns true for external targets.
func (t *ExternalTarget) IsExternal() bool {
	return true
}
