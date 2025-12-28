package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// WorkbookViewVisibility represents the visibility state of a workbook view.
type WorkbookViewVisibility string

const (
	// WorkbookViewVisible indicates the workbook view is visible.
	WorkbookViewVisible WorkbookViewVisibility = "visible"
	// WorkbookViewHidden indicates the workbook view is hidden.
	WorkbookViewHidden WorkbookViewVisibility = "hidden"
	// WorkbookViewVeryHidden indicates the workbook view is very hidden
	// (not accessible via UI).
	WorkbookViewVeryHidden WorkbookViewVisibility = "veryHidden"
)

// BookViews represents the book views container element (x:bookViews).
type BookViews struct {
	*openxml.CompositeElementBase
}

// NewBookViews creates a new BookViews element.
func NewBookViews() *BookViews {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"bookViews",
		PrefixDefault,
	)

	return &BookViews{CompositeElementBase: elem}
}

// WorkbookViews returns an iterator over all WorkbookView elements.
func (bv *BookViews) WorkbookViews() iter.Seq[*WorkbookView] {
	return func(yield func(*WorkbookView) bool) {
		for child := range bv.Children() {
			if child.LocalName() != "workbookView" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var wv *WorkbookView
			switch v := child.(type) {
			case *WorkbookView:
				wv = v
			case *openxml.CompositeElementBase:
				wv = &WorkbookView{CompositeElementBase: v}
			}
			if wv != nil && !yield(wv) {
				return
			}
		}
	}
}

// FirstWorkbookView returns the first workbook view, or nil if none exist.
func (bv *BookViews) FirstWorkbookView() *WorkbookView {
	for wv := range bv.WorkbookViews() {
		return wv
	}

	return nil
}

// GetOrCreateWorkbookView returns the first workbook view, creating one if
// none exist.
func (bv *BookViews) GetOrCreateWorkbookView() *WorkbookView {
	wv := bv.FirstWorkbookView()
	if wv != nil {
		return wv
	}
	wv = NewWorkbookView()
	bv.AppendChild(wv)

	return wv
}

// AddWorkbookView adds a new workbook view and returns it.
func (bv *BookViews) AddWorkbookView() *WorkbookView {
	wv := NewWorkbookView()
	bv.AppendChild(wv)

	return wv
}

// Clone creates a deep copy of this BookViews element.
func (bv *BookViews) Clone() openxml.Element {
	cloned := bv.CompositeElementBase.Clone()

	return &BookViews{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this BookViews element.
func (bv *BookViews) CloneNode(
	deep bool,
) openxml.Element {
	cloned := bv.CompositeElementBase.CloneNode(
		deep,
	)

	return &BookViews{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
