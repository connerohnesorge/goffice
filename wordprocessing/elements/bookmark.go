package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// BookmarkStart represents a bookmark start element (w:bookmarkStart).
type BookmarkStart struct {
	*openxml.CompositeElementBase
}

// NewBookmarkStart creates a new BookmarkStart element.
func NewBookmarkStart(id int, name string) *BookmarkStart {
	elem := openxml.NewCompositeElement(NamespaceWML, "bookmarkStart", PrefixW)
	bs := &BookmarkStart{CompositeElementBase: elem}
	bs.SetAttribute(openxml.NewAttribute(NamespaceWML, "id", PrefixW, strconv.Itoa(id)))
	bs.SetAttribute(openxml.NewAttribute(NamespaceWML, "name", PrefixW, name))
	return bs
}

// Id returns the bookmark ID.
func (bs *BookmarkStart) Id() int {
	attr, found := bs.GetAttribute("id", NamespaceWML)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetId sets the bookmark ID.
func (bs *BookmarkStart) SetId(id int) {
	bs.SetAttribute(openxml.NewAttribute(NamespaceWML, "id", PrefixW, strconv.Itoa(id)))
}

// Name returns the bookmark name.
func (bs *BookmarkStart) Name() string {
	attr, found := bs.GetAttribute("name", NamespaceWML)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetName sets the bookmark name.
func (bs *BookmarkStart) SetName(name string) {
	bs.SetAttribute(openxml.NewAttribute(NamespaceWML, "name", PrefixW, name))
}

// Clone creates a deep copy of this BookmarkStart element.
func (bs *BookmarkStart) Clone() openxml.Element {
	return &BookmarkStart{
		CompositeElementBase: bs.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this BookmarkStart element.
func (bs *BookmarkStart) CloneNode(deep bool) openxml.Element {
	return &BookmarkStart{
		CompositeElementBase: bs.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// BookmarkEnd represents a bookmark end element (w:bookmarkEnd).
type BookmarkEnd struct {
	*openxml.CompositeElementBase
}

// NewBookmarkEnd creates a new BookmarkEnd element.
func NewBookmarkEnd(id int) *BookmarkEnd {
	elem := openxml.NewCompositeElement(NamespaceWML, "bookmarkEnd", PrefixW)
	be := &BookmarkEnd{CompositeElementBase: elem}
	be.SetAttribute(openxml.NewAttribute(NamespaceWML, "id", PrefixW, strconv.Itoa(id)))
	return be
}

// Id returns the bookmark ID.
func (be *BookmarkEnd) Id() int {
	attr, found := be.GetAttribute("id", NamespaceWML)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetId sets the bookmark ID.
func (be *BookmarkEnd) SetId(id int) {
	be.SetAttribute(openxml.NewAttribute(NamespaceWML, "id", PrefixW, strconv.Itoa(id)))
}

// Clone creates a deep copy of this BookmarkEnd element.
func (be *BookmarkEnd) Clone() openxml.Element {
	return &BookmarkEnd{
		CompositeElementBase: be.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this BookmarkEnd element.
func (be *BookmarkEnd) CloneNode(deep bool) openxml.Element {
	return &BookmarkEnd{
		CompositeElementBase: be.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// CreateBookmarkPair creates a matching pair of bookmark start and end elements.
func CreateBookmarkPair(id int, name string) (*BookmarkStart, *BookmarkEnd) {
	return NewBookmarkStart(id, name), NewBookmarkEnd(id)
}
