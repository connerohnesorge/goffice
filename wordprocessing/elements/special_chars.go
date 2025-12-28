//nolint:revive // file-length-limit: many specialized character types
package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Attribute name constants for special character elements.
const (
	attrNameFont        = "font"
	attrNameChar        = "char"
	attrNameFldCharType = "fldCharType"
)

// CarriageReturn represents a carriage return element (w:cr).
// It inserts a carriage return character in the run content.
type CarriageReturn struct {
	*openxml.LeafElementBase
}

// NewCarriageReturn creates a new CarriageReturn element.
func NewCarriageReturn() *CarriageReturn {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"cr",
		PrefixW,
	)

	return &CarriageReturn{LeafElementBase: elem}
}

// Clone creates a deep copy of this CarriageReturn element.
func (cr *CarriageReturn) Clone() openxml.Element {
	return &CarriageReturn{
		LeafElementBase: cr.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CarriageReturn element.
func (cr *CarriageReturn) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cr.LeafElementBase.CloneNode(deep)

	return &CarriageReturn{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// SoftHyphen represents a soft hyphen element (w:softHyphen).
// It marks an optional line break point where a hyphen should appear
// if the line breaks at that point.
type SoftHyphen struct {
	*openxml.LeafElementBase
}

// NewSoftHyphen creates a new SoftHyphen element.
func NewSoftHyphen() *SoftHyphen {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"softHyphen",
		PrefixW,
	)

	return &SoftHyphen{LeafElementBase: elem}
}

// Clone creates a deep copy of this SoftHyphen element.
func (sh *SoftHyphen) Clone() openxml.Element {
	return &SoftHyphen{
		LeafElementBase: sh.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this SoftHyphen element.
func (sh *SoftHyphen) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sh.LeafElementBase.CloneNode(deep)

	return &SoftHyphen{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// NoBreakHyphen represents a non-breaking hyphen element (w:noBreakHyphen).
// It inserts a hyphen that prevents line breaking at that point.
type NoBreakHyphen struct {
	*openxml.LeafElementBase
}

// NewNoBreakHyphen creates a new NoBreakHyphen element.
func NewNoBreakHyphen() *NoBreakHyphen {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"noBreakHyphen",
		PrefixW,
	)

	return &NoBreakHyphen{LeafElementBase: elem}
}

// Clone creates a deep copy of this NoBreakHyphen element.
func (n *NoBreakHyphen) Clone() openxml.Element {
	return &NoBreakHyphen{
		LeafElementBase: n.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this NoBreakHyphen element.
func (n *NoBreakHyphen) CloneNode(
	deep bool,
) openxml.Element {
	cloned := n.LeafElementBase.CloneNode(deep)

	return &NoBreakHyphen{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// LastRenderedPageBreak represents the last rendered page break position
// (w:lastRenderedPageBreak). This element marks where a page break was
// rendered the last time the document was saved.
// It is informational and used by applications for pagination hints.
type LastRenderedPageBreak struct {
	*openxml.LeafElementBase
}

// NewLastRenderedPageBreak creates a new LastRenderedPageBreak element.
func NewLastRenderedPageBreak() *LastRenderedPageBreak {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"lastRenderedPageBreak",
		PrefixW,
	)

	return &LastRenderedPageBreak{
		LeafElementBase: elem,
	}
}

// Clone creates a deep copy of this LastRenderedPageBreak element.
func (l *LastRenderedPageBreak) Clone() openxml.Element {
	return &LastRenderedPageBreak{
		LeafElementBase: l.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this LastRenderedPageBreak element.
func (l *LastRenderedPageBreak) CloneNode(
	deep bool,
) openxml.Element {
	cloned := l.LeafElementBase.CloneNode(deep)

	return &LastRenderedPageBreak{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// Symbol represents a symbol character element (w:sym).
// It allows insertion of characters from symbol fonts.
type Symbol struct {
	*openxml.LeafElementBase
}

// NewSymbol creates a new Symbol element with the specified font and
// character code. The font parameter specifies the symbol font name
// (e.g., "Wingdings", "Symbol"). The char parameter specifies the
// character code in hexadecimal (e.g., "F041").
func NewSymbol(font, char string) *Symbol {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"sym",
		PrefixW,
	)
	s := &Symbol{LeafElementBase: elem}
	if font != "" {
		s.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				attrNameFont,
				PrefixW,
				font,
			),
		)
	}
	if char != "" {
		s.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				attrNameChar,
				PrefixW,
				char,
			),
		)
	}

	return s
}

// Font returns the symbol font name.
func (s *Symbol) Font() string {
	attr, found := s.GetAttribute(
		attrNameFont,
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetFont sets the symbol font name.
func (s *Symbol) SetFont(font string) {
	if font == "" {
		s.RemoveAttribute(
			attrNameFont,
			NamespaceWML,
		)
	} else {
		s.SetAttribute(openxml.NewAttribute(
			NamespaceWML, attrNameFont, PrefixW, font,
		))
	}
}

// Char returns the character code in hexadecimal.
func (s *Symbol) Char() string {
	attr, found := s.GetAttribute(
		attrNameChar,
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetChar sets the character code in hexadecimal.
func (s *Symbol) SetChar(char string) {
	if char == "" {
		s.RemoveAttribute(
			attrNameChar,
			NamespaceWML,
		)
	} else {
		s.SetAttribute(openxml.NewAttribute(
			NamespaceWML, attrNameChar, PrefixW, char,
		))
	}
}

// Clone creates a deep copy of this Symbol element.
func (s *Symbol) Clone() openxml.Element {
	return &Symbol{
		LeafElementBase: s.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Symbol element.
func (s *Symbol) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.LeafElementBase.CloneNode(deep)

	return &Symbol{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// FieldCharType represents the type of field character.
type FieldCharType string

const (
	// FieldCharBegin marks the start of a complex field.
	FieldCharBegin FieldCharType = "begin"
	// FieldCharSeparate separates field code from field result.
	FieldCharSeparate FieldCharType = "separate"
	// FieldCharEnd marks the end of a complex field.
	FieldCharEnd FieldCharType = "end"
)

// FieldChar represents a field character element (w:fldChar).
// Field characters are used to delimit complex fields in WordprocessingML.
type FieldChar struct {
	*openxml.LeafElementBase
}

// NewFieldChar creates a new FieldChar element with the specified type.
func NewFieldChar(
	charType FieldCharType,
) *FieldChar {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"fldChar",
		PrefixW,
	)
	fc := &FieldChar{LeafElementBase: elem}
	if charType != "" {
		fc.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				attrNameFldCharType,
				PrefixW,
				string(charType),
			),
		)
	}

	return fc
}

// Type returns the field character type.
func (fc *FieldChar) Type() FieldCharType {
	attr, found := fc.GetAttribute(
		attrNameFldCharType,
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return FieldCharType(attr.Value())
}

// SetType sets the field character type.
func (fc *FieldChar) SetType(
	charType FieldCharType,
) {
	if charType == "" {
		fc.RemoveAttribute(
			attrNameFldCharType,
			NamespaceWML,
		)
	} else {
		fc.SetAttribute(openxml.NewAttribute(
			NamespaceWML, attrNameFldCharType, PrefixW, string(charType),
		))
	}
}

// Dirty returns whether the field result should be recalculated.
func (fc *FieldChar) Dirty() bool {
	attr, found := fc.GetAttribute(
		"dirty",
		NamespaceWML,
	)
	if !found {
		return false
	}

	return attr.Value() == attrValueTrue ||
		attr.Value() == "1"
}

// SetDirty sets whether the field result should be recalculated.
//
//nolint:revive // flag-parameter: API requires bool setter
func (fc *FieldChar) SetDirty(dirty bool) {
	if dirty {
		fc.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"dirty",
				PrefixW,
				"true",
			),
		)
	} else {
		fc.RemoveAttribute("dirty", NamespaceWML)
	}
}

// Clone creates a deep copy of this FieldChar element.
func (fc *FieldChar) Clone() openxml.Element {
	return &FieldChar{
		LeafElementBase: fc.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this FieldChar element.
func (fc *FieldChar) CloneNode(
	deep bool,
) openxml.Element {
	cloned := fc.LeafElementBase.CloneNode(deep)

	return &FieldChar{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// InstrText represents field instruction text (w:instrText).
// It contains the field code for a complex field.
type InstrText struct {
	*openxml.LeafElementBase
}

// NewInstrText creates a new InstrText element with the specified
// instruction text.
func NewInstrText(text string) *InstrText {
	elem := openxml.NewLeafElementWithText(
		NamespaceWML,
		"instrText",
		PrefixW,
		text,
	)
	it := &InstrText{LeafElementBase: elem}
	// Preserve whitespace if the text contains significant spaces
	if needsSpacePreserve(text) {
		it.SetSpace("preserve")
	}

	return it
}

// Text returns the instruction text content.
func (it *InstrText) Text() string {
	return it.InnerText()
}

// SetText sets the instruction text content.
func (it *InstrText) SetText(text string) {
	it.SetInnerText(text)
	if needsSpacePreserve(text) {
		it.SetSpace("preserve")
	} else {
		it.RemoveAttribute("space", NamespaceXML)
	}
}

// Space returns the xml:space attribute value.
func (it *InstrText) Space() string {
	attr, found := it.GetAttribute(
		"space",
		NamespaceXML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSpace sets the xml:space attribute.
func (it *InstrText) SetSpace(value string) {
	attr := openxml.NewAttribute(
		NamespaceXML,
		"space",
		"xml",
		value,
	)
	it.SetAttribute(attr)
}

// Clone creates a deep copy of this InstrText element.
func (it *InstrText) Clone() openxml.Element {
	return &InstrText{
		LeafElementBase: it.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this InstrText element.
func (it *InstrText) CloneNode(
	deep bool,
) openxml.Element {
	cloned := it.LeafElementBase.CloneNode(deep)

	return &InstrText{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// PositionalTabAlignment specifies the alignment of a positional tab.
type PositionalTabAlignment string

const (
	// PositionalTabAlignmentLeft aligns content to the left of the tab stop.
	PositionalTabAlignmentLeft PositionalTabAlignment = "left"
	// PositionalTabAlignmentCenter centers content at the tab stop.
	PositionalTabAlignmentCenter PositionalTabAlignment = "center"
	// PositionalTabAlignmentRight aligns content to the right of the tab stop.
	PositionalTabAlignmentRight PositionalTabAlignment = "right"
)

// PositionalTabRelativeTo specifies what the positional tab is relative to.
type PositionalTabRelativeTo string

const (
	// PositionalTabRelativeToMargin positions relative to the page margins.
	PositionalTabRelativeToMargin PositionalTabRelativeTo = "margin"
	// PositionalTabRelativeToIndent positions relative to the paragraph
	// indentation.
	PositionalTabRelativeToIndent PositionalTabRelativeTo = "indent"
)

// PositionalTabLeader specifies the leader character for a positional tab.
type PositionalTabLeader string

const (
	// PositionalTabLeaderNone uses no leader character.
	PositionalTabLeaderNone PositionalTabLeader = "none"
	// PositionalTabLeaderDot uses a dot leader.
	PositionalTabLeaderDot PositionalTabLeader = "dot"
	// PositionalTabLeaderHyphen uses a hyphen leader.
	PositionalTabLeaderHyphen PositionalTabLeader = "hyphen"
	// PositionalTabLeaderUnderscore uses an underscore leader.
	PositionalTabLeaderUnderscore PositionalTabLeader = "underscore"
	// PositionalTabLeaderMiddleDot uses a middle dot leader.
	PositionalTabLeaderMiddleDot PositionalTabLeader = "middleDot"
)

// PositionalTab represents a positional tab element (w:ptab).
// Positional tabs allow precise positioning of text within a document.
type PositionalTab struct {
	*openxml.LeafElementBase
}

// NewPositionalTab creates a new PositionalTab element.
func NewPositionalTab() *PositionalTab {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"ptab",
		PrefixW,
	)

	return &PositionalTab{LeafElementBase: elem}
}

// NewPositionalTabWithOptions creates a new PositionalTab element with
// specified options.
func NewPositionalTabWithOptions(
	alignment PositionalTabAlignment,
	relativeTo PositionalTabRelativeTo,
	leader PositionalTabLeader,
) *PositionalTab {
	pt := NewPositionalTab()
	if alignment != "" {
		pt.SetAlignment(alignment)
	}
	if relativeTo != "" {
		pt.SetRelativeTo(relativeTo)
	}
	if leader != "" {
		pt.SetLeader(leader)
	}

	return pt
}

// Alignment returns the positional tab alignment.
func (pt *PositionalTab) Alignment() PositionalTabAlignment {
	attr, found := pt.GetAttribute(
		"alignment",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return PositionalTabAlignment(attr.Value())
}

// SetAlignment sets the positional tab alignment.
func (pt *PositionalTab) SetAlignment(
	alignment PositionalTabAlignment,
) {
	if alignment == "" {
		pt.RemoveAttribute(
			"alignment",
			NamespaceWML,
		)
	} else {
		pt.SetAttribute(openxml.NewAttribute(
			NamespaceWML, "alignment", PrefixW, string(alignment),
		))
	}
}

// RelativeTo returns what the positional tab is relative to.
func (pt *PositionalTab) RelativeTo() PositionalTabRelativeTo {
	attr, found := pt.GetAttribute(
		"relativeTo",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return PositionalTabRelativeTo(attr.Value())
}

// SetRelativeTo sets what the positional tab is relative to.
func (pt *PositionalTab) SetRelativeTo(
	relativeTo PositionalTabRelativeTo,
) {
	if relativeTo == "" {
		pt.RemoveAttribute(
			"relativeTo",
			NamespaceWML,
		)
	} else {
		pt.SetAttribute(openxml.NewAttribute(
			NamespaceWML, "relativeTo", PrefixW, string(relativeTo),
		))
	}
}

// Leader returns the leader character.
func (pt *PositionalTab) Leader() PositionalTabLeader {
	attr, found := pt.GetAttribute(
		"leader",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return PositionalTabLeader(attr.Value())
}

// SetLeader sets the leader character.
func (pt *PositionalTab) SetLeader(
	leader PositionalTabLeader,
) {
	if leader == "" {
		pt.RemoveAttribute("leader", NamespaceWML)
	} else {
		pt.SetAttribute(openxml.NewAttribute(
			NamespaceWML, "leader", PrefixW, string(leader),
		))
	}
}

// Clone creates a deep copy of this PositionalTab element.
func (pt *PositionalTab) Clone() openxml.Element {
	return &PositionalTab{
		LeafElementBase: pt.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this PositionalTab element.
func (pt *PositionalTab) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pt.LeafElementBase.CloneNode(deep)

	return &PositionalTab{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// Separator represents a footnote/endnote separator element (w:separator).
// This element renders as a horizontal line in footnote/endnote areas.
type Separator struct {
	*openxml.LeafElementBase
}

// NewSeparator creates a new Separator element.
func NewSeparator() *Separator {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"separator",
		PrefixW,
	)

	return &Separator{LeafElementBase: elem}
}

// Clone creates a deep copy of this Separator element.
func (s *Separator) Clone() openxml.Element {
	return &Separator{
		LeafElementBase: s.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Separator element.
func (s *Separator) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.LeafElementBase.CloneNode(deep)

	return &Separator{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// ContinuationSeparator represents a continuation separator element
// (w:continuationSeparator). This element renders as a horizontal line
// that spans the width of the page when footnotes/endnotes continue from
// the previous page.
type ContinuationSeparator struct {
	*openxml.LeafElementBase
}

// NewContinuationSeparator creates a new ContinuationSeparator element.
func NewContinuationSeparator() *ContinuationSeparator {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"continuationSeparator",
		PrefixW,
	)

	return &ContinuationSeparator{
		LeafElementBase: elem,
	}
}

// Clone creates a deep copy of this ContinuationSeparator element.
func (cs *ContinuationSeparator) Clone() openxml.Element {
	return &ContinuationSeparator{
		LeafElementBase: cs.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this ContinuationSeparator element.
func (cs *ContinuationSeparator) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cs.LeafElementBase.CloneNode(deep)

	return &ContinuationSeparator{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
