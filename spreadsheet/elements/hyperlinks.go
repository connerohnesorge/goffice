package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// Hyperlinks represents the hyperlinks container element (x:hyperlinks).
// It contains all hyperlinks in a worksheet.
type Hyperlinks struct {
	*openxml.CompositeElementBase
}

// NewHyperlinks creates a new Hyperlinks element.
func NewHyperlinks() *Hyperlinks {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"hyperlinks",
		PrefixDefault,
	)

	return &Hyperlinks{CompositeElementBase: elem}
}

// GetHyperlinks returns an iterator over all Hyperlink elements.
func (h *Hyperlinks) GetHyperlinks() iter.Seq[*Hyperlink] {
	return func(yield func(*Hyperlink) bool) {
		for child := range h.Children() {
			if child.LocalName() != "hyperlink" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var hyperlink *Hyperlink
			if hl, ok := child.(*Hyperlink); ok {
				hyperlink = hl
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				hyperlink = &Hyperlink{LeafElementBase: leaf}
			}
			if hyperlink != nil &&
				!yield(hyperlink) {
				return
			}
		}
	}
}

// HyperlinkCount returns the number of hyperlinks.
func (h *Hyperlinks) HyperlinkCount() int {
	count := 0
	for range h.GetHyperlinks() {
		count++
	}

	return count
}

// GetHyperlinkByRef returns the hyperlink at the given cell reference,
// or nil if not found.
func (h *Hyperlinks) GetHyperlinkByRef(
	ref string,
) *Hyperlink {
	for hyperlink := range h.GetHyperlinks() {
		if hyperlink.Ref() == ref {
			return hyperlink
		}
	}

	return nil
}

// AddHyperlink adds a new hyperlink at the given cell reference.
// The relationshipId is the r:id that references the target in relationships.
func (h *Hyperlinks) AddHyperlink(
	ref, relationshipId string,
) *Hyperlink {
	hyperlink := NewHyperlink()
	hyperlink.SetRef(ref)
	if relationshipId != "" {
		hyperlink.SetRelationshipId(
			relationshipId,
		)
	}
	h.AppendChild(hyperlink)

	return hyperlink
}

// AddInternalHyperlink adds a new internal hyperlink (within workbook).
// Use this for links to other sheets or named ranges.
func (h *Hyperlinks) AddInternalHyperlink(
	ref, location string,
) *Hyperlink {
	hyperlink := NewHyperlink()
	hyperlink.SetRef(ref)
	hyperlink.SetLocation(location)
	h.AppendChild(hyperlink)

	return hyperlink
}

// RemoveHyperlink removes the hyperlink at the given cell reference.
func (h *Hyperlinks) RemoveHyperlink(
	ref string,
) bool {
	hyperlink := h.GetHyperlinkByRef(ref)
	if hyperlink == nil {
		return false
	}

	return h.RemoveChild(hyperlink)
}

// HasHyperlink returns whether a hyperlink exists at the given cell reference.
func (h *Hyperlinks) HasHyperlink(
	ref string,
) bool {
	return h.GetHyperlinkByRef(ref) != nil
}

// Clone creates a deep copy of this Hyperlinks element.
func (h *Hyperlinks) Clone() openxml.Element {
	cloned := h.CompositeElementBase.Clone()

	return &Hyperlinks{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Hyperlinks element.
func (h *Hyperlinks) CloneNode(
	deep bool,
) openxml.Element {
	cloned := h.CompositeElementBase.CloneNode(
		deep,
	)

	return &Hyperlinks{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Hyperlink represents a single hyperlink element (x:hyperlink).
type Hyperlink struct {
	*openxml.LeafElementBase
}

// NewHyperlink creates a new Hyperlink element.
func NewHyperlink() *Hyperlink {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"hyperlink",
		PrefixDefault,
	)

	return &Hyperlink{LeafElementBase: elem}
}

// Ref returns the cell reference for the hyperlink. Attribute: ref.
func (hl *Hyperlink) Ref() string {
	attr, found := hl.GetAttribute("ref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the cell reference for the hyperlink. Attribute: ref.
func (hl *Hyperlink) SetRef(ref string) {
	hl.SetAttribute(
		openxml.NewAttribute(
			"",
			"ref",
			"",
			ref,
		),
	)
}

// RelationshipId returns the relationship ID (r:id) linking to the target.
func (hl *Hyperlink) RelationshipId() string {
	attr, found := hl.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID (r:id) linking to the target.
func (hl *Hyperlink) SetRelationshipId(
	id string,
) {
	if id == "" {
		hl.RemoveAttribute(
			"id",
			NamespaceRelationships,
		)

		return
	}
	hl.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"id",
			PrefixR,
			id,
		),
	)
}

// Location returns the internal location (e.g., "Sheet2!A1").
// Attribute: location.
func (hl *Hyperlink) Location() string {
	attr, found := hl.GetAttribute("location", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetLocation sets the internal location. Attribute: location.
func (hl *Hyperlink) SetLocation(
	location string,
) {
	if location == "" {
		hl.RemoveAttribute("location", "")

		return
	}
	hl.SetAttribute(
		openxml.NewAttribute(
			"",
			"location",
			"",
			location,
		),
	)
}

// Display returns the display text for the hyperlink. Attribute: display.
func (hl *Hyperlink) Display() string {
	attr, found := hl.GetAttribute("display", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDisplay sets the display text for the hyperlink. Attribute: display.
func (hl *Hyperlink) SetDisplay(display string) {
	if display == "" {
		hl.RemoveAttribute("display", "")

		return
	}
	hl.SetAttribute(
		openxml.NewAttribute(
			"",
			"display",
			"",
			display,
		),
	)
}

// Tooltip returns the tooltip text for the hyperlink. Attribute: tooltip.
func (hl *Hyperlink) Tooltip() string {
	attr, found := hl.GetAttribute("tooltip", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTooltip sets the tooltip text for the hyperlink. Attribute: tooltip.
func (hl *Hyperlink) SetTooltip(tooltip string) {
	if tooltip == "" {
		hl.RemoveAttribute("tooltip", "")

		return
	}
	hl.SetAttribute(
		openxml.NewAttribute(
			"",
			"tooltip",
			"",
			tooltip,
		),
	)
}

// IsExternal returns whether the hyperlink is external (has relationship ID).
func (hl *Hyperlink) IsExternal() bool {
	return hl.RelationshipId() != ""
}

// IsInternal returns whether the hyperlink is internal (has location only).
func (hl *Hyperlink) IsInternal() bool {
	return hl.Location() != "" &&
		hl.RelationshipId() == ""
}

// Clone creates a deep copy of this Hyperlink element.
func (hl *Hyperlink) Clone() openxml.Element {
	cloned := hl.LeafElementBase.Clone()

	return &Hyperlink{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Hyperlink element.
func (hl *Hyperlink) CloneNode(
	deep bool,
) openxml.Element {
	cloned := hl.LeafElementBase.CloneNode(deep)

	return &Hyperlink{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
