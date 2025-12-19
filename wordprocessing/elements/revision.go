package elements

import (
	"iter"
	"strconv"
	"time"

	"github.com/connerohnesorge/goffice/openxml"
)

// InsertedRun represents the w:ins element for tracked insertions.
type InsertedRun struct {
	*openxml.CompositeElementBase
}

// NewInsertedRun creates a new InsertedRun element.
func NewInsertedRun(
	id int,
	author string,
	date time.Time,
) *InsertedRun {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"ins",
		PrefixW,
	)
	ins := &InsertedRun{
		CompositeElementBase: elem,
	}
	ins.SetId(id)
	ins.SetAuthor(author)
	ins.SetDate(date)

	return ins
}

// Id returns the revision ID.
func (ins *InsertedRun) Id() int {
	attr, found := ins.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetId sets the revision ID.
func (ins *InsertedRun) SetId(id int) {
	ins.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Author returns the author who made this change.
func (ins *InsertedRun) Author() string {
	attr, found := ins.GetAttribute(
		"author",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAuthor sets the author who made this change.
func (ins *InsertedRun) SetAuthor(author string) {
	ins.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"author",
			PrefixW,
			author,
		),
	)
}

// Date returns the date/time of this change.
func (ins *InsertedRun) Date() time.Time {
	attr, found := ins.GetAttribute(
		"date",
		NamespaceWML,
	)
	if !found {
		return time.Time{}
	}
	t, _ := time.Parse(time.RFC3339, attr.Value())

	return t
}

// SetDate sets the date/time of this change.
func (ins *InsertedRun) SetDate(date time.Time) {
	ins.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"date",
			PrefixW,
			date.Format(time.RFC3339),
		),
	)
}

// Runs returns an iterator over all Run elements in this insertion.
func (ins *InsertedRun) Runs() iter.Seq[*Run] {
	return func(yield func(*Run) bool) {
		for child := range ins.Children() {
			if child.LocalName() == "r" &&
				child.NamespaceURI() == NamespaceWML {
				var r *Run
				if run, ok := child.(*Run); ok {
					r = run
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					r = &Run{CompositeElementBase: comp}
				}
				if r != nil && !yield(r) {
					return
				}
			}
		}
	}
}

// AppendRun appends a new Run element with the given text.
func (ins *InsertedRun) AppendRun(
	text string,
) *Run {
	r := NewRun(text)
	ins.AppendChild(r)

	return r
}

// InnerText returns the concatenated text content of all runs.
func (ins *InsertedRun) InnerText() string {
	var text string
	for r := range ins.Runs() {
		text += r.InnerText()
	}

	return text
}

// Clone creates a deep copy of this InsertedRun element.
func (ins *InsertedRun) Clone() openxml.Element {
	return &InsertedRun{
		CompositeElementBase: ins.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this InsertedRun element.
func (ins *InsertedRun) CloneNode(
	deep bool,
) openxml.Element {
	return &InsertedRun{
		CompositeElementBase: ins.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// DeletedRun represents the w:del element for tracked deletions.
type DeletedRun struct {
	*openxml.CompositeElementBase
}

// NewDeletedRun creates a new DeletedRun element.
func NewDeletedRun(
	id int,
	author string,
	date time.Time,
) *DeletedRun {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"del",
		PrefixW,
	)
	del := &DeletedRun{CompositeElementBase: elem}
	del.SetId(id)
	del.SetAuthor(author)
	del.SetDate(date)

	return del
}

// Id returns the revision ID.
func (del *DeletedRun) Id() int {
	attr, found := del.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetId sets the revision ID.
func (del *DeletedRun) SetId(id int) {
	del.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Author returns the author who made this change.
func (del *DeletedRun) Author() string {
	attr, found := del.GetAttribute(
		"author",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAuthor sets the author who made this change.
func (del *DeletedRun) SetAuthor(author string) {
	del.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"author",
			PrefixW,
			author,
		),
	)
}

// Date returns the date/time of this change.
func (del *DeletedRun) Date() time.Time {
	attr, found := del.GetAttribute(
		"date",
		NamespaceWML,
	)
	if !found {
		return time.Time{}
	}
	t, _ := time.Parse(time.RFC3339, attr.Value())

	return t
}

// SetDate sets the date/time of this change.
func (del *DeletedRun) SetDate(date time.Time) {
	del.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"date",
			PrefixW,
			date.Format(time.RFC3339),
		),
	)
}

// Runs returns an iterator over all Run elements in this deletion.
func (del *DeletedRun) Runs() iter.Seq[*Run] {
	return func(yield func(*Run) bool) {
		for child := range del.Children() {
			if child.LocalName() == "r" &&
				child.NamespaceURI() == NamespaceWML {
				var r *Run
				if run, ok := child.(*Run); ok {
					r = run
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					r = &Run{CompositeElementBase: comp}
				}
				if r != nil && !yield(r) {
					return
				}
			}
		}
	}
}

// DeletedTexts returns an iterator over all DeletedText elements in this deletion.
func (del *DeletedRun) DeletedTexts() iter.Seq[*DeletedText] {
	return func(yield func(*DeletedText) bool) {
		// Look inside runs for delText elements
		for r := range del.Runs() {
			for child := range r.Children() {
				if child.LocalName() == "delText" &&
					child.NamespaceURI() == NamespaceWML {
					var dt *DeletedText
					if delText, ok := child.(*DeletedText); ok {
						dt = delText
					} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
						dt = &DeletedText{LeafElementBase: leaf}
					}
					if dt != nil && !yield(dt) {
						return
					}
				}
			}
		}
	}
}

// AppendDeletedRun appends a new Run with DeletedText.
func (del *DeletedRun) AppendDeletedRun(
	text string,
) *Run {
	r := NewRun("")
	// Remove the default text element
	t := r.Text()
	if t != nil {
		r.RemoveChild(t)
	}
	// Add deleted text
	delText := NewDeletedText(text)
	r.AppendChild(delText)
	del.AppendChild(r)

	return r
}

// InnerText returns the concatenated deleted text content.
func (del *DeletedRun) InnerText() string {
	var text string
	for dt := range del.DeletedTexts() {
		text += dt.InnerText()
	}

	return text
}

// Clone creates a deep copy of this DeletedRun element.
func (del *DeletedRun) Clone() openxml.Element {
	return &DeletedRun{
		CompositeElementBase: del.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DeletedRun element.
func (del *DeletedRun) CloneNode(
	deep bool,
) openxml.Element {
	return &DeletedRun{
		CompositeElementBase: del.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// DeletedText represents the w:delText element for deleted text content.
type DeletedText struct {
	*openxml.LeafElementBase
}

// NewDeletedText creates a new DeletedText element.
func NewDeletedText(text string) *DeletedText {
	elem := openxml.NewLeafElementWithText(
		NamespaceWML,
		"delText",
		PrefixW,
		text,
	)
	dt := &DeletedText{LeafElementBase: elem}
	// Preserve whitespace if needed
	if needsSpacePreserve(text) {
		dt.SetSpace("preserve")
	}

	return dt
}

// SetText sets the text content.
func (dt *DeletedText) SetText(value string) {
	dt.SetInnerText(value)
	if needsSpacePreserve(value) {
		dt.SetSpace("preserve")
	} else {
		dt.RemoveAttribute("space", NamespaceXML)
	}
}

// Space returns the xml:space attribute value.
func (dt *DeletedText) Space() string {
	attr, found := dt.GetAttribute(
		"space",
		NamespaceXML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSpace sets the xml:space attribute.
func (dt *DeletedText) SetSpace(value string) {
	attr := openxml.NewAttribute(
		NamespaceXML,
		"space",
		"xml",
		value,
	)
	dt.SetAttribute(attr)
}

// Clone creates a deep copy of this DeletedText element.
func (dt *DeletedText) Clone() openxml.Element {
	return &DeletedText{
		LeafElementBase: dt.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this DeletedText element.
func (dt *DeletedText) CloneNode(
	deep bool,
) openxml.Element {
	return &DeletedText{
		LeafElementBase: dt.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// MoveFromRun represents the w:moveFrom element for tracked move source.
type MoveFromRun struct {
	*openxml.CompositeElementBase
}

// NewMoveFromRun creates a new MoveFromRun element.
func NewMoveFromRun(
	id int,
	author string,
	date time.Time,
) *MoveFromRun {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"moveFrom",
		PrefixW,
	)
	mf := &MoveFromRun{CompositeElementBase: elem}
	mf.SetId(id)
	mf.SetAuthor(author)
	mf.SetDate(date)

	return mf
}

// Id returns the revision ID.
func (mf *MoveFromRun) Id() int {
	attr, found := mf.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetId sets the revision ID.
func (mf *MoveFromRun) SetId(id int) {
	mf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Author returns the author who made this change.
func (mf *MoveFromRun) Author() string {
	attr, found := mf.GetAttribute(
		"author",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAuthor sets the author who made this change.
func (mf *MoveFromRun) SetAuthor(author string) {
	mf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"author",
			PrefixW,
			author,
		),
	)
}

// Date returns the date/time of this change.
func (mf *MoveFromRun) Date() time.Time {
	attr, found := mf.GetAttribute(
		"date",
		NamespaceWML,
	)
	if !found {
		return time.Time{}
	}
	t, _ := time.Parse(time.RFC3339, attr.Value())

	return t
}

// SetDate sets the date/time of this change.
func (mf *MoveFromRun) SetDate(date time.Time) {
	mf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"date",
			PrefixW,
			date.Format(time.RFC3339),
		),
	)
}

// Clone creates a deep copy of this MoveFromRun element.
func (mf *MoveFromRun) Clone() openxml.Element {
	return &MoveFromRun{
		CompositeElementBase: mf.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// MoveToRun represents the w:moveTo element for tracked move destination.
type MoveToRun struct {
	*openxml.CompositeElementBase
}

// NewMoveToRun creates a new MoveToRun element.
func NewMoveToRun(
	id int,
	author string,
	date time.Time,
) *MoveToRun {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"moveTo",
		PrefixW,
	)
	mt := &MoveToRun{CompositeElementBase: elem}
	mt.SetId(id)
	mt.SetAuthor(author)
	mt.SetDate(date)

	return mt
}

// Id returns the revision ID.
func (mt *MoveToRun) Id() int {
	attr, found := mt.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetId sets the revision ID.
func (mt *MoveToRun) SetId(id int) {
	mt.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Author returns the author who made this change.
func (mt *MoveToRun) Author() string {
	attr, found := mt.GetAttribute(
		"author",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAuthor sets the author who made this change.
func (mt *MoveToRun) SetAuthor(author string) {
	mt.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"author",
			PrefixW,
			author,
		),
	)
}

// Date returns the date/time of this change.
func (mt *MoveToRun) Date() time.Time {
	attr, found := mt.GetAttribute(
		"date",
		NamespaceWML,
	)
	if !found {
		return time.Time{}
	}
	t, _ := time.Parse(time.RFC3339, attr.Value())

	return t
}

// SetDate sets the date/time of this change.
func (mt *MoveToRun) SetDate(date time.Time) {
	mt.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"date",
			PrefixW,
			date.Format(time.RFC3339),
		),
	)
}

// Clone creates a deep copy of this MoveToRun element.
func (mt *MoveToRun) Clone() openxml.Element {
	return &MoveToRun{
		CompositeElementBase: mt.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// RunPropertiesChange represents the w:rPrChange element for tracked run property changes.
type RunPropertiesChange struct {
	*openxml.CompositeElementBase
}

// NewRunPropertiesChange creates a new RunPropertiesChange element.
func NewRunPropertiesChange(
	id int,
	author string,
	date time.Time,
) *RunPropertiesChange {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"rPrChange",
		PrefixW,
	)
	rpc := &RunPropertiesChange{
		CompositeElementBase: elem,
	}
	rpc.SetId(id)
	rpc.SetAuthor(author)
	rpc.SetDate(date)

	return rpc
}

// Id returns the revision ID.
func (rpc *RunPropertiesChange) Id() int {
	attr, found := rpc.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetId sets the revision ID.
func (rpc *RunPropertiesChange) SetId(id int) {
	rpc.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Author returns the author who made this change.
func (rpc *RunPropertiesChange) Author() string {
	attr, found := rpc.GetAttribute(
		"author",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAuthor sets the author who made this change.
func (rpc *RunPropertiesChange) SetAuthor(
	author string,
) {
	rpc.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"author",
			PrefixW,
			author,
		),
	)
}

// Date returns the date/time of this change.
func (rpc *RunPropertiesChange) Date() time.Time {
	attr, found := rpc.GetAttribute(
		"date",
		NamespaceWML,
	)
	if !found {
		return time.Time{}
	}
	t, _ := time.Parse(time.RFC3339, attr.Value())

	return t
}

// SetDate sets the date/time of this change.
func (rpc *RunPropertiesChange) SetDate(
	date time.Time,
) {
	rpc.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"date",
			PrefixW,
			date.Format(time.RFC3339),
		),
	)
}

// PreviousRunProperties returns the previous run properties.
func (rpc *RunPropertiesChange) PreviousRunProperties() *RunProperties {
	elem := rpc.GetElement("rPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*RunProperties); ok {
		return rp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RunProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetPreviousRunProperties sets the previous run properties.
func (rpc *RunPropertiesChange) SetPreviousRunProperties(
	rp *RunProperties,
) {
	existing := rpc.GetElement(
		"rPr",
		NamespaceWML,
	)
	if existing != nil {
		rpc.RemoveChild(existing)
	}
	if rp != nil {
		rpc.AppendChild(rp)
	}
}

// Clone creates a deep copy of this RunPropertiesChange element.
func (rpc *RunPropertiesChange) Clone() openxml.Element {
	return &RunPropertiesChange{
		CompositeElementBase: rpc.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ParagraphPropertiesChange represents the w:pPrChange element for tracked paragraph property changes.
type ParagraphPropertiesChange struct {
	*openxml.CompositeElementBase
}

// NewParagraphPropertiesChange creates a new ParagraphPropertiesChange element.
func NewParagraphPropertiesChange(
	id int,
	author string,
	date time.Time,
) *ParagraphPropertiesChange {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"pPrChange",
		PrefixW,
	)
	ppc := &ParagraphPropertiesChange{
		CompositeElementBase: elem,
	}
	ppc.SetId(id)
	ppc.SetAuthor(author)
	ppc.SetDate(date)

	return ppc
}

// Id returns the revision ID.
func (ppc *ParagraphPropertiesChange) Id() int {
	attr, found := ppc.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetId sets the revision ID.
func (ppc *ParagraphPropertiesChange) SetId(
	id int,
) {
	ppc.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Author returns the author who made this change.
func (ppc *ParagraphPropertiesChange) Author() string {
	attr, found := ppc.GetAttribute(
		"author",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetAuthor sets the author who made this change.
func (ppc *ParagraphPropertiesChange) SetAuthor(
	author string,
) {
	ppc.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"author",
			PrefixW,
			author,
		),
	)
}

// Date returns the date/time of this change.
func (ppc *ParagraphPropertiesChange) Date() time.Time {
	attr, found := ppc.GetAttribute(
		"date",
		NamespaceWML,
	)
	if !found {
		return time.Time{}
	}
	t, _ := time.Parse(time.RFC3339, attr.Value())

	return t
}

// SetDate sets the date/time of this change.
func (ppc *ParagraphPropertiesChange) SetDate(
	date time.Time,
) {
	ppc.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"date",
			PrefixW,
			date.Format(time.RFC3339),
		),
	)
}

// PreviousParagraphProperties returns the previous paragraph properties.
func (ppc *ParagraphPropertiesChange) PreviousParagraphProperties() *ParagraphProperties {
	elem := ppc.GetElement("pPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if pp, ok := elem.(*ParagraphProperties); ok {
		return pp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ParagraphProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetPreviousParagraphProperties sets the previous paragraph properties.
func (ppc *ParagraphPropertiesChange) SetPreviousParagraphProperties(
	pp *ParagraphProperties,
) {
	existing := ppc.GetElement(
		"pPr",
		NamespaceWML,
	)
	if existing != nil {
		ppc.RemoveChild(existing)
	}
	if pp != nil {
		ppc.AppendChild(pp)
	}
}

// Clone creates a deep copy of this ParagraphPropertiesChange element.
func (ppc *ParagraphPropertiesChange) Clone() openxml.Element {
	return &ParagraphPropertiesChange{
		CompositeElementBase: ppc.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
