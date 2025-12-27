package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Attribute name constants for hyperlink elements.
const (
	// attrNameID is already defined in section_properties.go
	attrNameAnchor = "anchor"
)

// Hyperlink represents a hyperlink element (w:hyperlink).
type Hyperlink struct {
	*openxml.CompositeElementBase
}

// NewHyperlink creates a new Hyperlink element with text and relationship ID.
func NewHyperlink(text, relId string) *Hyperlink {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"hyperlink",
		PrefixW,
	)
	h := &Hyperlink{CompositeElementBase: elem}
	if relId != "" {
		h.SetAttribute(
			openxml.NewAttribute(
				openxml.NamespaceRelationships,
				attrNameID,
				"r",
				relId,
			),
		)
	}
	if text != "" {
		r := NewRun(text)
		// Style the run as a hyperlink (blue, underlined)
		r.SetColor("0000FF")
		r.SetUnderline(UnderlineSingle)
		h.AppendChild(r)
	}

	return h
}

// NewInternalHyperlink creates a new Hyperlink element that links
// to a bookmark.
func NewInternalHyperlink(
	text, anchor string,
) *Hyperlink {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"hyperlink",
		PrefixW,
	)
	h := &Hyperlink{CompositeElementBase: elem}
	if anchor != "" {
		h.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				attrNameAnchor,
				PrefixW,
				anchor,
			),
		)
	}
	if text != "" {
		r := NewRun(text)
		r.SetColor("0000FF")
		r.SetUnderline(UnderlineSingle)
		h.AppendChild(r)
	}

	return h
}

// RelationshipId returns the relationship ID (for external links).
func (h *Hyperlink) RelationshipId() string {
	attr, found := h.GetAttribute(
		attrNameID,
		openxml.NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID.
func (h *Hyperlink) SetRelationshipId(
	relId string,
) {
	if relId == "" {
		h.RemoveAttribute(
			attrNameID,
			openxml.NamespaceRelationships,
		)
	} else {
		h.SetAttribute(openxml.NewAttribute(
			openxml.NamespaceRelationships, attrNameID, "r", relId,
		))
	}
}

// Anchor returns the anchor (bookmark name) for internal links.
func (h *Hyperlink) Anchor() string {
	attr, found := h.GetAttribute(
		attrNameAnchor,
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAnchor sets the anchor (bookmark name) for internal links.
func (h *Hyperlink) SetAnchor(anchor string) {
	if anchor == "" {
		h.RemoveAttribute(
			attrNameAnchor,
			NamespaceWML,
		)
	} else {
		h.SetAttribute(openxml.NewAttribute(
			NamespaceWML, attrNameAnchor, PrefixW, anchor,
		))
	}
}

// Tooltip returns the tooltip text.
func (h *Hyperlink) Tooltip() string {
	attr, found := h.GetAttribute(
		"tooltip",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTooltip sets the tooltip text.
func (h *Hyperlink) SetTooltip(tooltip string) {
	if tooltip == "" {
		h.RemoveAttribute("tooltip", NamespaceWML)
	} else {
		h.SetAttribute(openxml.NewAttribute(
			NamespaceWML, "tooltip", PrefixW, tooltip,
		))
	}
}

// History returns whether to add this link to history.
func (h *Hyperlink) History() bool {
	attr, found := h.GetAttribute(
		"history",
		NamespaceWML,
	)
	if !found {
		return true // default
	}
	val := attr.Value()

	return val != "false" && val != "0"
}

// SetHistory sets whether to add this link to history.
//
//nolint:revive // flag-parameter: bool param is appropriate for setter
func (h *Hyperlink) SetHistory(b bool) {
	if b {
		h.RemoveAttribute("history", NamespaceWML)
	} else {
		h.SetAttribute(openxml.NewAttribute(
			NamespaceWML, "history", PrefixW, "false",
		))
	}
}

// DocLocation returns the document location (for internal document links).
func (h *Hyperlink) DocLocation() string {
	attr, found := h.GetAttribute(
		"docLocation",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDocLocation sets the document location.
func (h *Hyperlink) SetDocLocation(loc string) {
	if loc == "" {
		h.RemoveAttribute(
			"docLocation",
			NamespaceWML,
		)
	} else {
		h.SetAttribute(openxml.NewAttribute(
			NamespaceWML, "docLocation", PrefixW, loc,
		))
	}
}

// IsExternal returns true if this is an external hyperlink.
func (h *Hyperlink) IsExternal() bool {
	return h.RelationshipId() != ""
}

// IsInternal returns true if this is an internal (anchor) hyperlink.
func (h *Hyperlink) IsInternal() bool {
	return h.Anchor() != ""
}

// AppendRun appends a run element to the hyperlink.
func (h *Hyperlink) AppendRun(text string) *Run {
	r := NewRun(text)
	h.AppendChild(r)

	return r
}

// Clone creates a deep copy of this Hyperlink element.
func (h *Hyperlink) Clone() openxml.Element {
	cloned := h.CompositeElementBase.Clone()

	return &Hyperlink{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Hyperlink element.
func (h *Hyperlink) CloneNode(
	deep bool, //nolint:revive // flag-parameter
) openxml.Element {
	cloned := h.CompositeElementBase.CloneNode(
		deep,
	)

	return &Hyperlink{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
