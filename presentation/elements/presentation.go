//nolint:revive // This file contains presentation element implementation.
package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Presentation represents the root presentation element (p:presentation).
type Presentation struct {
	*openxml.PartRootElementBase
}

// NewPresentation creates a new Presentation element.
func NewPresentation() *Presentation {
	elem := openxml.NewPartRootElement(
		NamespacePresentationML,
		"presentation",
		PrefixP,
	)
	p := &Presentation{PartRootElementBase: elem}

	// Add required child elements
	p.AppendChild(NewSlideMasterIdList())
	p.AppendChild(NewSlideIdList())
	p.AppendChild(
		NewSlideSize(
			Screen4x3Width,
			Screen4x3Height,
			SlideSizeScreen4x3,
		),
	)
	p.AppendChild(
		NewNotesSize(
			Screen4x3Height,
			Screen4x3Width,
		),
	)

	return p
}

// SlideMasterIdList returns the slide master ID list element.
func (p *Presentation) SlideMasterIdList() *SlideMasterIdList {
	elem := p.GetElement(
		"sldMasterIdLst",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if smil, ok := elem.(*SlideMasterIdList); ok {
		return smil
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &SlideMasterIdList{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSlideMasterIdList returns the slide master ID list, creating if needed.
func (p *Presentation) GetOrCreateSlideMasterIdList() *SlideMasterIdList {
	smil := p.SlideMasterIdList()
	if smil != nil {
		return smil
	}
	smil = NewSlideMasterIdList()
	// Insert as first child
	if first := p.FirstChild(); first != nil {
		p.InsertBefore(smil, first)
	} else {
		p.AppendChild(smil)
	}

	return smil
}

// NotesMasterIdList returns the notes master ID list element.
func (p *Presentation) NotesMasterIdList() *NotesMasterIdList {
	elem := p.GetElement(
		"notesMasterIdLst",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if nmil, ok := elem.(*NotesMasterIdList); ok {
		return nmil
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &NotesMasterIdList{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateNotesMasterIdList returns the notes master ID list, creating if needed.
func (p *Presentation) GetOrCreateNotesMasterIdList() *NotesMasterIdList {
	nmil := p.NotesMasterIdList()
	if nmil != nil {
		return nmil
	}
	nmil = NewNotesMasterIdList()
	// Insert after slide master ID list
	smil := p.SlideMasterIdList()
	if smil != nil {
		p.InsertAfter(nmil, smil)
	} else {
		if first := p.FirstChild(); first != nil {
			p.InsertBefore(nmil, first)
		} else {
			p.AppendChild(nmil)
		}
	}

	return nmil
}

// HandoutMasterIdList returns the handout master ID list element.
func (p *Presentation) HandoutMasterIdList() *HandoutMasterIdList {
	elem := p.GetElement(
		"handoutMasterIdLst",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if hmil, ok := elem.(*HandoutMasterIdList); ok {
		return hmil
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &HandoutMasterIdList{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SlideIdList returns the slide ID list element.
func (p *Presentation) SlideIdList() *SlideIdList {
	elem := p.GetElement(
		"sldIdLst",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if sil, ok := elem.(*SlideIdList); ok {
		return sil
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &SlideIdList{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSlideIdList returns the slide ID list, creating if needed.
func (p *Presentation) GetOrCreateSlideIdList() *SlideIdList {
	sil := p.SlideIdList()
	if sil != nil {
		return sil
	}
	sil = NewSlideIdList()
	// Insert after notes master ID list or slide master ID list
	if nmil := p.NotesMasterIdList(); nmil != nil {
		p.InsertAfter(sil, nmil)
	} else if smil := p.SlideMasterIdList(); smil != nil {
		p.InsertAfter(sil, smil)
	} else {
		if first := p.FirstChild(); first != nil {
			p.InsertBefore(sil, first)
		} else {
			p.AppendChild(sil)
		}
	}

	return sil
}

// SlideSize returns the slide size element.
func (p *Presentation) SlideSize() *SlideSize {
	elem := p.GetElement(
		"sldSz",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if ss, ok := elem.(*SlideSize); ok {
		return ss
	}
	if leaf := wrapLeafElement(elem); leaf != nil {
		return &SlideSize{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateSlideSize returns the slide size, creating with default 4:3 if needed.
func (p *Presentation) GetOrCreateSlideSize() *SlideSize {
	ss := p.SlideSize()
	if ss != nil {
		return ss
	}
	ss = NewSlideSize(
		Screen4x3Width,
		Screen4x3Height,
		SlideSizeScreen4x3,
	)
	p.AppendChild(ss)

	return ss
}

// NotesSize returns the notes size element.
func (p *Presentation) NotesSize() *NotesSize {
	elem := p.GetElement(
		"notesSz",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if ns, ok := elem.(*NotesSize); ok {
		return ns
	}
	if leaf := wrapLeafElement(elem); leaf != nil {
		return &NotesSize{LeafElementBase: leaf}
	}

	return nil
}

// AddSlideId adds a slide ID entry to the slide ID list.
func (p *Presentation) AddSlideId(
	id uint32,
	relId string,
) *SlideId {
	sil := p.GetOrCreateSlideIdList()

	return sil.AddSlideId(id, relId)
}

// AddSlideMasterId adds a slide master ID entry to the slide master ID list.
func (p *Presentation) AddSlideMasterId(
	id uint32,
	relId string,
) *SlideMasterId {
	smil := p.GetOrCreateSlideMasterIdList()

	return smil.AddSlideMasterId(id, relId)
}

// Clone creates a deep copy of this Presentation element.
func (p *Presentation) Clone() openxml.Element {
	cloned := p.PartRootElementBase.Clone()

	return &Presentation{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// ===========================================================================
// SlideMasterIdList (p:sldMasterIdLst)
// ===========================================================================

// SlideMasterIdList represents the slide master ID list element (p:sldMasterIdLst).
type SlideMasterIdList struct {
	*openxml.CompositeElementBase
}

// NewSlideMasterIdList creates a new SlideMasterIdList element.
func NewSlideMasterIdList() *SlideMasterIdList {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"sldMasterIdLst",
		PrefixP,
	)

	return &SlideMasterIdList{
		CompositeElementBase: elem,
	}
}

// SlideMasterIds returns all slide master ID entries.
func (smil *SlideMasterIdList) SlideMasterIds() []*SlideMasterId {
	var ids []*SlideMasterId
	for child := range smil.Children() {
		if child.LocalName() == "sldMasterId" &&
			child.NamespaceURI() == NamespacePresentationML {
			if smid, ok := child.(*SlideMasterId); ok {
				ids = append(ids, smid)
			} else if leaf := wrapLeafElement(child); leaf != nil {
				ids = append(ids, &SlideMasterId{LeafElementBase: leaf})
			}
		}
	}

	return ids
}

// AddSlideMasterId adds a slide master ID entry.
func (smil *SlideMasterIdList) AddSlideMasterId(
	id uint32,
	relId string,
) *SlideMasterId {
	smid := NewSlideMasterId(id, relId)
	smil.AppendChild(smid)

	return smid
}

// Clone creates a deep copy of this SlideMasterIdList element.
func (smil *SlideMasterIdList) Clone() openxml.Element {
	cloned := smil.CompositeElementBase.Clone()

	return &SlideMasterIdList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// SlideMasterId (p:sldMasterId)
// ===========================================================================

// SlideMasterId represents a slide master ID entry (p:sldMasterId).
type SlideMasterId struct {
	*openxml.LeafElementBase
}

// NewSlideMasterId creates a new SlideMasterId element.
func NewSlideMasterId(
	id uint32,
	relId string,
) *SlideMasterId {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"sldMasterId",
		PrefixP,
	)
	smid := &SlideMasterId{LeafElementBase: elem}
	smid.SetId(id)
	smid.SetRelId(relId)

	return smid
}

// Id returns the slide master ID.
func (smid *SlideMasterId) Id() uint32 {
	attr, found := smid.GetAttribute("id", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10,
		32,
	)

	return uint32(val)
}

// SetId sets the slide master ID.
func (smid *SlideMasterId) SetId(id uint32) {
	smid.SetAttribute(openxml.NewAttribute(
		"",
		"id",
		"",
		strconv.FormatUint(uint64(id), 10),
	))
}

// RelId returns the relationship ID.
func (smid *SlideMasterId) RelId() string {
	attr, found := smid.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelId sets the relationship ID.
func (smid *SlideMasterId) SetRelId(
	relId string,
) {
	smid.SetAttribute(openxml.NewAttribute(
		NamespaceRelationships,
		"id",
		PrefixR,
		relId,
	))
}

// Clone creates a deep copy of this SlideMasterId element.
func (smid *SlideMasterId) Clone() openxml.Element {
	cloned := smid.LeafElementBase.Clone()

	return &SlideMasterId{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// ===========================================================================
// NotesMasterIdList (p:notesMasterIdLst)
// ===========================================================================

// NotesMasterIdList represents the notes master ID list element (p:notesMasterIdLst).
type NotesMasterIdList struct {
	*openxml.CompositeElementBase
}

// NewNotesMasterIdList creates a new NotesMasterIdList element.
func NewNotesMasterIdList() *NotesMasterIdList {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"notesMasterIdLst",
		PrefixP,
	)

	return &NotesMasterIdList{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this NotesMasterIdList element.
func (nmil *NotesMasterIdList) Clone() openxml.Element {
	cloned := nmil.CompositeElementBase.Clone()

	return &NotesMasterIdList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// HandoutMasterIdList (p:handoutMasterIdLst)
// ===========================================================================

// HandoutMasterIdList represents the handout master ID list element (p:handoutMasterIdLst).
type HandoutMasterIdList struct {
	*openxml.CompositeElementBase
}

// NewHandoutMasterIdList creates a new HandoutMasterIdList element.
func NewHandoutMasterIdList() *HandoutMasterIdList {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"handoutMasterIdLst",
		PrefixP,
	)

	return &HandoutMasterIdList{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this HandoutMasterIdList element.
func (hmil *HandoutMasterIdList) Clone() openxml.Element {
	cloned := hmil.CompositeElementBase.Clone()

	return &HandoutMasterIdList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// SlideIdList (p:sldIdLst)
// ===========================================================================

// SlideIdList represents the slide ID list element (p:sldIdLst).
type SlideIdList struct {
	*openxml.CompositeElementBase
}

// NewSlideIdList creates a new SlideIdList element.
func NewSlideIdList() *SlideIdList {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"sldIdLst",
		PrefixP,
	)

	return &SlideIdList{
		CompositeElementBase: elem,
	}
}

// SlideIds returns all slide ID entries.
func (sil *SlideIdList) SlideIds() []*SlideId {
	var ids []*SlideId
	for child := range sil.Children() {
		if child.LocalName() == "sldId" &&
			child.NamespaceURI() == NamespacePresentationML {
			if sid, ok := child.(*SlideId); ok {
				ids = append(ids, sid)
			} else if leaf := wrapLeafElement(child); leaf != nil {
				ids = append(ids, &SlideId{LeafElementBase: leaf})
			}
		}
	}

	return ids
}

// AddSlideId adds a slide ID entry.
func (sil *SlideIdList) AddSlideId(
	id uint32,
	relId string,
) *SlideId {
	sid := NewSlideId(id, relId)
	sil.AppendChild(sid)

	return sid
}

// Count returns the number of slides.
func (sil *SlideIdList) Count() int {
	return len(sil.SlideIds())
}

// Clone creates a deep copy of this SlideIdList element.
func (sil *SlideIdList) Clone() openxml.Element {
	cloned := sil.CompositeElementBase.Clone()

	return &SlideIdList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// SlideId (p:sldId)
// ===========================================================================

// SlideId represents a slide ID entry (p:sldId).
type SlideId struct {
	*openxml.LeafElementBase
}

// NewSlideId creates a new SlideId element.
func NewSlideId(
	id uint32,
	relId string,
) *SlideId {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"sldId",
		PrefixP,
	)
	sid := &SlideId{LeafElementBase: elem}
	sid.SetId(id)
	sid.SetRelId(relId)

	return sid
}

// Id returns the slide ID.
func (sid *SlideId) Id() uint32 {
	attr, found := sid.GetAttribute("id", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10,
		32,
	)

	return uint32(val)
}

// SetId sets the slide ID.
func (sid *SlideId) SetId(id uint32) {
	sid.SetAttribute(openxml.NewAttribute(
		"",
		"id",
		"",
		strconv.FormatUint(uint64(id), 10),
	))
}

// RelId returns the relationship ID.
func (sid *SlideId) RelId() string {
	attr, found := sid.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelId sets the relationship ID.
func (sid *SlideId) SetRelId(relId string) {
	sid.SetAttribute(openxml.NewAttribute(
		NamespaceRelationships,
		"id",
		PrefixR,
		relId,
	))
}

// Clone creates a deep copy of this SlideId element.
func (sid *SlideId) Clone() openxml.Element {
	cloned := sid.LeafElementBase.Clone()

	return &SlideId{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// ===========================================================================
// SlideSize (p:sldSz)
// ===========================================================================

// SlideSize represents the slide size element (p:sldSz).
type SlideSize struct {
	*openxml.LeafElementBase
}

// NewSlideSize creates a new SlideSize element with the given dimensions.
func NewSlideSize(
	cx, cy int,
	sizeType SlideSizeType,
) *SlideSize {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"sldSz",
		PrefixP,
	)
	ss := &SlideSize{LeafElementBase: elem}
	ss.SetCx(cx)
	ss.SetCy(cy)
	if sizeType != "" {
		ss.SetType(sizeType)
	}

	return ss
}

// Cx returns the width in EMUs.
func (ss *SlideSize) Cx() int {
	attr, found := ss.GetAttribute("cx", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCx sets the width in EMUs.
func (ss *SlideSize) SetCx(cx int) {
	ss.SetAttribute(openxml.NewAttribute(
		"",
		"cx",
		"",
		strconv.Itoa(cx),
	))
}

// Cy returns the height in EMUs.
func (ss *SlideSize) Cy() int {
	attr, found := ss.GetAttribute("cy", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCy sets the height in EMUs.
func (ss *SlideSize) SetCy(cy int) {
	ss.SetAttribute(openxml.NewAttribute(
		"",
		"cy",
		"",
		strconv.Itoa(cy),
	))
}

// Type returns the slide size type.
func (ss *SlideSize) Type() SlideSizeType {
	attr, found := ss.GetAttribute("type", "")
	if !found {
		return ""
	}

	return SlideSizeType(attr.Value())
}

// SetType sets the slide size type.
func (ss *SlideSize) SetType(t SlideSizeType) {
	if t == "" {
		ss.RemoveAttribute("type", "")

		return
	}
	ss.SetAttribute(openxml.NewAttribute(
		"",
		"type",
		"",
		string(t),
	))
}

// Clone creates a deep copy of this SlideSize element.
func (ss *SlideSize) Clone() openxml.Element {
	cloned := ss.LeafElementBase.Clone()

	return &SlideSize{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// ===========================================================================
// NotesSize (p:notesSz)
// ===========================================================================

// NotesSize represents the notes size element (p:notesSz).
type NotesSize struct {
	*openxml.LeafElementBase
}

// NewNotesSize creates a new NotesSize element with the given dimensions.
func NewNotesSize(cx, cy int) *NotesSize {
	elem := openxml.NewLeafElement(
		NamespacePresentationML,
		"notesSz",
		PrefixP,
	)
	ns := &NotesSize{LeafElementBase: elem}
	ns.SetCx(cx)
	ns.SetCy(cy)

	return ns
}

// Cx returns the width in EMUs.
func (ns *NotesSize) Cx() int {
	attr, found := ns.GetAttribute("cx", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCx sets the width in EMUs.
func (ns *NotesSize) SetCx(cx int) {
	ns.SetAttribute(openxml.NewAttribute(
		"",
		"cx",
		"",
		strconv.Itoa(cx),
	))
}

// Cy returns the height in EMUs.
func (ns *NotesSize) Cy() int {
	attr, found := ns.GetAttribute("cy", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCy sets the height in EMUs.
func (ns *NotesSize) SetCy(cy int) {
	ns.SetAttribute(openxml.NewAttribute(
		"",
		"cy",
		"",
		strconv.Itoa(cy),
	))
}

// Clone creates a deep copy of this NotesSize element.
func (ns *NotesSize) Clone() openxml.Element {
	cloned := ns.LeafElementBase.Clone()

	return &NotesSize{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// ===========================================================================
// DefaultTextStyle (p:defaultTextStyle)
// ===========================================================================

// DefaultTextStyle represents the default text style element (p:defaultTextStyle).
type DefaultTextStyle struct {
	*openxml.CompositeElementBase
}

// NewDefaultTextStyle creates a new DefaultTextStyle element.
func NewDefaultTextStyle() *DefaultTextStyle {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"defaultTextStyle",
		PrefixP,
	)

	return &DefaultTextStyle{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this DefaultTextStyle element.
func (dts *DefaultTextStyle) Clone() openxml.Element {
	cloned := dts.CompositeElementBase.Clone()

	return &DefaultTextStyle{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
