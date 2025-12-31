//nolint:revive // file-length-limit: revision tracking requires comprehensive accept/reject logic
package wordprocessing

import (
	"errors"
	"time"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// RevisionType represents the type of a tracked change.
type RevisionType int

const (
	// RevisionTypeInsert represents an insertion revision (w:ins).
	RevisionTypeInsert RevisionType = iota

	// RevisionTypeDelete represents a deletion revision (w:del).
	RevisionTypeDelete

	// RevisionTypeMoveFrom represents a move source revision (w:moveFrom).
	RevisionTypeMoveFrom

	// RevisionTypeMoveTo represents a move destination revision (w:moveTo).
	RevisionTypeMoveTo

	// RevisionTypeFormatChange represents a formatting change revision
	// (w:rPrChange or w:pPrChange).
	RevisionTypeFormatChange
)

// String returns the string representation of the revision type.
func (rt RevisionType) String() string {
	switch rt {
	case RevisionTypeInsert:
		return "Insert"
	case RevisionTypeDelete:
		return "Delete"
	case RevisionTypeMoveFrom:
		return "MoveFrom"
	case RevisionTypeMoveTo:
		return "MoveTo"
	case RevisionTypeFormatChange:
		return "FormatChange"
	default:
		return "Unknown"
	}
}

// Revision wraps a revision tracking element and provides unified access
// to its properties.
type Revision struct {
	// element is the underlying revision element (InsertedRun, DeletedRun, etc.).
	element any

	// typ is the type of this revision.
	typ RevisionType

	// parent is the parent OpenXML element.
	parent openxml.Element

	// document is the document this revision belongs to.
	document *Document
}

// newRevision creates a new Revision wrapper for the given element.
// It automatically detects the revision type based on the element type.
func newRevision(
	element any,
	parent openxml.Element,
	doc *Document,
) *Revision {
	rev := &Revision{
		element:  element,
		parent:   parent,
		document: doc,
	}

	// Detect revision type from element type
	switch element.(type) {
	case *elements.InsertedRun:
		rev.typ = RevisionTypeInsert
	case *elements.DeletedRun:
		rev.typ = RevisionTypeDelete
	case *elements.MoveFromRun:
		rev.typ = RevisionTypeMoveFrom
	case *elements.MoveToRun:
		rev.typ = RevisionTypeMoveTo
	case *elements.RunPropertiesChange, *elements.ParagraphPropertiesChange:
		rev.typ = RevisionTypeFormatChange
	}

	return rev
}

// Type returns the type of this revision.
func (r *Revision) Type() RevisionType {
	return r.typ
}

// Author returns the author who made this revision.
func (r *Revision) Author() string {
	switch elem := r.element.(type) {
	case *elements.InsertedRun:
		return elem.Author()
	case *elements.DeletedRun:
		return elem.Author()
	case *elements.MoveFromRun:
		return elem.Author()
	case *elements.MoveToRun:
		return elem.Author()
	case *elements.RunPropertiesChange:
		return elem.Author()
	case *elements.ParagraphPropertiesChange:
		return elem.Author()
	default:
		return ""
	}
}

// Date returns the date/time when this revision was made.
func (r *Revision) Date() time.Time {
	switch elem := r.element.(type) {
	case *elements.InsertedRun:
		return elem.Date()
	case *elements.DeletedRun:
		return elem.Date()
	case *elements.MoveFromRun:
		return elem.Date()
	case *elements.MoveToRun:
		return elem.Date()
	case *elements.RunPropertiesChange:
		return elem.Date()
	case *elements.ParagraphPropertiesChange:
		return elem.Date()
	default:
		return time.Time{}
	}
}

// Id returns the revision ID.
func (r *Revision) Id() int {
	switch elem := r.element.(type) {
	case *elements.InsertedRun:
		return elem.Id()
	case *elements.DeletedRun:
		return elem.Id()
	case *elements.MoveFromRun:
		return elem.Id()
	case *elements.MoveToRun:
		return elem.Id()
	case *elements.RunPropertiesChange:
		return elem.Id()
	case *elements.ParagraphPropertiesChange:
		return elem.Id()
	default:
		return 0
	}
}

// Content returns the text content of this revision.
// For insertions and moves, this returns the inserted/moved text.
// For deletions, this returns the deleted text.
// For format changes, this returns an empty string.
func (r *Revision) Content() string {
	switch elem := r.element.(type) {
	case *elements.InsertedRun:
		return elem.InnerText()
	case *elements.DeletedRun:
		return elem.InnerText()
	case *elements.MoveFromRun:
		// MoveFromRun doesn't have InnerText in the current implementation
		// We'll need to traverse children to get text
		return ""
	case *elements.MoveToRun:
		// MoveToRun doesn't have InnerText in the current implementation
		// We'll need to traverse children to get text
		return ""
	case *elements.RunPropertiesChange, *elements.ParagraphPropertiesChange:
		// Format changes don't have text content
		return ""
	default:
		return ""
	}
}

// Element returns the underlying revision element.
func (r *Revision) Element() any {
	return r.element
}

// Parent returns the parent OpenXML element.
func (r *Revision) Parent() openxml.Element {
	return r.parent
}

// Document returns the document this revision belongs to.
func (r *Revision) Document() *Document {
	return r.document
}

// Accept accepts this revision, applying the change to the document.
// The behavior depends on the revision type:
// - Insert: Replace the InsertedRun with its Run children
// - Delete: Remove the DeletedRun (content stays deleted)
// - MoveFrom/MoveTo: Returns error (requires move pair coordination)
// - FormatChange: Remove the change wrapper, keep current properties
func (r *Revision) Accept() error {
	switch r.typ {
	case RevisionTypeInsert:
		return r.acceptInsertedRun()
	case RevisionTypeDelete:
		return r.acceptDeletedRun()
	case RevisionTypeMoveFrom, RevisionTypeMoveTo:
		return r.acceptMoveRun()
	case RevisionTypeFormatChange:
		return r.acceptFormatChange()
	default:
		return nil
	}
}

// Reject rejects this revision, reverting the change.
// The behavior depends on the revision type:
// - Insert: Remove the InsertedRun and all content
// - Delete: Convert DeletedText to regular Text, restore content
// - MoveFrom/MoveTo: Returns error (requires move pair coordination)
// - FormatChange: Restore previous properties from the change wrapper
func (r *Revision) Reject() error {
	switch r.typ {
	case RevisionTypeInsert:
		return r.rejectInsertedRun()
	case RevisionTypeDelete:
		return r.rejectDeletedRun()
	case RevisionTypeMoveFrom, RevisionTypeMoveTo:
		return r.rejectMoveRun()
	case RevisionTypeFormatChange:
		return r.rejectFormatChange()
	default:
		return nil
	}
}

// acceptInsertedRun replaces the InsertedRun with its Run children.
func (r *Revision) acceptInsertedRun() error {
	ins, ok := r.element.(*elements.InsertedRun)
	if !ok {
		return nil
	}

	parent, ok := r.parent.(openxml.CompositeElement)
	if !ok || parent == nil {
		return nil
	}

	// Collect all runs before modifying the tree
	runs := make([]*elements.Run, 0, 2)
	for run := range ins.Runs() {
		runs = append(runs, run)
	}

	// Insert runs before the InsertedRun
	for _, run := range runs {
		parent.InsertBefore(run, ins)
	}

	// Remove the InsertedRun wrapper
	parent.RemoveChild(ins)

	return nil
}

// rejectInsertedRun removes the InsertedRun and all content.
func (r *Revision) rejectInsertedRun() error {
	ins, ok := r.element.(*elements.InsertedRun)
	if !ok {
		return nil
	}

	parent, ok := r.parent.(openxml.CompositeElement)
	if !ok || parent == nil {
		return nil
	}

	// Simply remove the entire InsertedRun (removes inserted content)
	parent.RemoveChild(ins)

	return nil
}

// acceptDeletedRun removes the DeletedRun (content stays deleted).
func (r *Revision) acceptDeletedRun() error {
	del, ok := r.element.(*elements.DeletedRun)
	if !ok {
		return nil
	}

	parent, ok := r.parent.(openxml.CompositeElement)
	if !ok || parent == nil {
		return nil
	}

	// Simply remove the DeletedRun (accepts the deletion)
	parent.RemoveChild(del)

	return nil
}

// rejectDeletedRun converts DeletedText to regular Text and restores content.
func (r *Revision) rejectDeletedRun() error {
	del, ok := r.element.(*elements.DeletedRun)
	if !ok {
		return nil
	}

	parent, ok := r.parent.(openxml.CompositeElement)
	if !ok || parent == nil {
		return nil
	}

	// Collect all runs before modifying the tree
	runs := make([]*elements.Run, 0, 2)
	for run := range del.Runs() {
		runs = append(runs, run)
	}

	// For each run, convert DeletedText to regular Text
	for _, run := range runs {
		r.convertDeletedTextsInRun(run)
		// Insert the run before the DeletedRun
		parent.InsertBefore(run, del)
	}

	// Remove the DeletedRun wrapper
	parent.RemoveChild(del)

	return nil
}

// convertDeletedTextsInRun converts all DeletedText elements to regular Text.
func (*Revision) convertDeletedTextsInRun(
	run *elements.Run,
) {
	// Collect deleted texts
	var deletedTexts []*elements.DeletedText
	for dt := range run.Children() {
		if dt.LocalName() != "delText" ||
			dt.NamespaceURI() != elements.NamespaceWML {
			continue
		}
		switch v := dt.(type) {
		case *elements.DeletedText:
			deletedTexts = append(deletedTexts, v)
		case *openxml.LeafElementBase:
			deletedTexts = append(deletedTexts, &elements.DeletedText{LeafElementBase: v})
		}
	}

	// Convert each DeletedText to regular Text
	for _, delText := range deletedTexts {
		textContent := delText.InnerText()
		// Create new Text element with same content
		newText := elements.NewText(textContent)
		// Insert before deleted text
		run.InsertBefore(newText, delText)
		// Remove deleted text
		run.RemoveChild(delText)
	}
}

// acceptMoveRun returns an error indicating move operations are not yet supported.
func (*Revision) acceptMoveRun() error {
	return ErrMoveCoordinationNotImplemented
}

// rejectMoveRun returns an error indicating move operations are not yet supported.
func (*Revision) rejectMoveRun() error {
	return ErrMoveCoordinationNotImplemented
}

// acceptFormatChange removes the change wrapper, keeping current properties.
func (r *Revision) acceptFormatChange() error {
	parent, ok := r.parent.(openxml.CompositeElement)
	if !ok || parent == nil {
		return nil
	}

	// Cast to appropriate type and remove the change wrapper
	switch change := r.element.(type) {
	case *elements.RunPropertiesChange:
		// The change wrapper is inside RunProperties, just remove it
		parent.RemoveChild(change)
	case *elements.ParagraphPropertiesChange:
		// The change wrapper is inside ParagraphProperties, just remove it
		parent.RemoveChild(change)
	}

	return nil
}

// rejectFormatChange extracts previous properties and restores them.
func (r *Revision) rejectFormatChange() error {
	parent, ok := r.parent.(openxml.CompositeElement)
	if !ok || parent == nil {
		return nil
	}

	switch change := r.element.(type) {
	case *elements.RunPropertiesChange:
		r.rejectRunPropertiesChange(parent, change)
	case *elements.ParagraphPropertiesChange:
		r.rejectParagraphPropertiesChange(parent, change)
	}

	return nil
}

// rejectRunPropertiesChange restores previous run properties.
func (*Revision) rejectRunPropertiesChange(
	parent openxml.CompositeElement,
	change *elements.RunPropertiesChange,
) {
	prevProps := change.PreviousRunProperties()
	if prevProps == nil {
		parent.RemoveChild(change)

		return
	}

	runProps, ok := parent.(*elements.RunProperties)
	if !ok {
		parent.RemoveChild(change)

		return
	}

	// Remove all current property children except the change element
	var toRemove []openxml.Element
	for child := range runProps.Children() {
		if child != change {
			toRemove = append(toRemove, child)
		}
	}
	for _, child := range toRemove {
		runProps.RemoveChild(child)
	}

	// Copy all children from previous properties
	for child := range prevProps.Children() {
		cloned := child.CloneNode(true)
		runProps.AppendChild(cloned)
	}

	// Remove the change wrapper
	parent.RemoveChild(change)
}

// rejectParagraphPropertiesChange restores previous paragraph properties.
func (*Revision) rejectParagraphPropertiesChange(
	parent openxml.CompositeElement,
	change *elements.ParagraphPropertiesChange,
) {
	prevProps := change.PreviousParagraphProperties()
	if prevProps == nil {
		parent.RemoveChild(change)

		return
	}

	paraProps, ok := parent.(*elements.ParagraphProperties)
	if !ok {
		parent.RemoveChild(change)

		return
	}

	// Remove all current property children except the change element
	var toRemove []openxml.Element
	for child := range paraProps.Children() {
		if child != change {
			toRemove = append(toRemove, child)
		}
	}
	for _, child := range toRemove {
		paraProps.RemoveChild(child)
	}

	// Copy all children from previous properties
	for child := range prevProps.Children() {
		cloned := child.CloneNode(true)
		paraProps.AppendChild(cloned)
	}

	// Remove the change wrapper
	parent.RemoveChild(change)
}

// ErrMoveCoordinationNotImplemented is returned when attempting to accept/reject
// move operations, which require coordinated handling of move pairs.
var ErrMoveCoordinationNotImplemented = errors.New(
	"move pair coordination not yet implemented",
)

// collectRevisionsFromElement recursively walks the element tree and collects
// all revisions. It detects InsertedRun, DeletedRun, MoveFromRun, MoveToRun,
// RunPropertiesChange, and ParagraphPropertiesChange elements and wraps them
// in Revision structs.
func collectRevisionsFromElement(
	elem openxml.Element,
	doc *Document,
) []*Revision {
	var revisions []*Revision

	// Check if this element itself is a revision
	switch v := elem.(type) {
	case *elements.InsertedRun:
		// Get parent for revision
		parent := elem.Parent()
		revisions = append(revisions, newRevision(v, parent, doc))
	case *elements.DeletedRun:
		parent := elem.Parent()
		revisions = append(revisions, newRevision(v, parent, doc))
	case *elements.MoveFromRun:
		parent := elem.Parent()
		revisions = append(revisions, newRevision(v, parent, doc))
	case *elements.MoveToRun:
		parent := elem.Parent()
		revisions = append(revisions, newRevision(v, parent, doc))
	case *elements.RunPropertiesChange:
		parent := elem.Parent()
		revisions = append(revisions, newRevision(v, parent, doc))
	case *elements.ParagraphPropertiesChange:
		parent := elem.Parent()
		revisions = append(revisions, newRevision(v, parent, doc))
	}

	// Recurse into composite elements
	if composite, ok := elem.(openxml.CompositeElement); ok {
		for child := range composite.Children() {
			childRevisions := collectRevisionsFromElement(
				child,
				doc,
			)
			revisions = append(
				revisions,
				childRevisions...)
		}
	}

	return revisions
}

// GetRevisions returns all revisions in the document.
// It traverses the entire document body, headers, and footers to collect
// all revision tracking elements.
func (d *Document) GetRevisions() []*Revision {
	var revisions []*Revision

	mainPart := d.MainPart()
	if mainPart == nil {
		return revisions
	}

	// Traverse the main document body
	doc := mainPart.Document()
	if doc != nil {
		body := doc.Body()
		if body != nil {
			bodyRevisions := collectRevisionsFromElement(
				body,
				d,
			)
			revisions = append(
				revisions,
				bodyRevisions...)
		}
	}

	// Traverse headers
	for _, headerPart := range mainPart.HeaderParts() {
		header := headerPart.Header()
		if header != nil {
			headerRevisions := collectRevisionsFromElement(
				header,
				d,
			)
			revisions = append(
				revisions,
				headerRevisions...)
		}
	}

	// Traverse footers
	for _, footerPart := range mainPart.FooterParts() {
		footer := footerPart.Footer()
		if footer != nil {
			footerRevisions := collectRevisionsFromElement(
				footer,
				d,
			)
			revisions = append(
				revisions,
				footerRevisions...)
		}
	}

	return revisions
}

// AcceptAllRevisions accepts all revisions in the document.
// This applies all tracked changes and removes the revision markup.
func (d *Document) AcceptAllRevisions() error {
	revisions := d.GetRevisions()

	// Accept revisions in reverse order to avoid issues with element removal
	for i := len(revisions) - 1; i >= 0; i-- {
		if err := revisions[i].Accept(); err != nil {
			return err
		}
	}

	return nil
}

// RejectAllRevisions rejects all revisions in the document.
// This reverts all tracked changes and removes the revision markup.
func (d *Document) RejectAllRevisions() error {
	revisions := d.GetRevisions()

	// Reject revisions in reverse order to avoid issues with element removal
	for i := len(revisions) - 1; i >= 0; i-- {
		if err := revisions[i].Reject(); err != nil {
			return err
		}
	}

	return nil
}

// AcceptRevisionsByAuthor accepts all revisions made by the specified author.
func (d *Document) AcceptRevisionsByAuthor(
	author string,
) error {
	revisions := d.GetRevisions()

	// Accept revisions in reverse order to avoid issues with element removal
	for i := len(revisions) - 1; i >= 0; i-- {
		if revisions[i].Author() == author {
			if err := revisions[i].Accept(); err != nil {
				return err
			}
		}
	}

	return nil
}

// RejectRevisionsByAuthor rejects all revisions made by the specified author.
func (d *Document) RejectRevisionsByAuthor(
	author string,
) error {
	revisions := d.GetRevisions()

	// Reject revisions in reverse order to avoid issues with element removal
	for i := len(revisions) - 1; i >= 0; i-- {
		if revisions[i].Author() == author {
			if err := revisions[i].Reject(); err != nil {
				return err
			}
		}
	}

	return nil
}

// RevisionCollection provides filtering and querying capabilities for revisions.
type RevisionCollection struct {
	revisions []*Revision
}

// NewRevisionCollection creates a new RevisionCollection from a slice of revisions.
func NewRevisionCollection(
	revisions []*Revision,
) *RevisionCollection {
	return &RevisionCollection{
		revisions: revisions,
	}
}

// ByAuthor filters revisions by author name.
func (rc *RevisionCollection) ByAuthor(
	author string,
) []*Revision {
	var filtered []*Revision
	for _, rev := range rc.revisions {
		if rev.Author() == author {
			filtered = append(filtered, rev)
		}
	}

	return filtered
}

// ByType filters revisions by type.
func (rc *RevisionCollection) ByType(
	typ RevisionType,
) []*Revision {
	var filtered []*Revision
	for _, rev := range rc.revisions {
		if rev.Type() == typ {
			filtered = append(filtered, rev)
		}
	}

	return filtered
}

// Count returns the number of revisions in the collection.
func (rc *RevisionCollection) Count() int {
	return len(rc.revisions)
}

// All returns all revisions in the collection.
func (rc *RevisionCollection) All() []*Revision {
	return rc.revisions
}
