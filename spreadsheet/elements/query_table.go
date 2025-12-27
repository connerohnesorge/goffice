package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// QueryTableRoot represents the queryTable root element (x:queryTable).
// It defines an Excel query table with external data queries for a
// worksheet.
// Note: This is the root element type, distinct from the auto-generated
// QueryTable type which is for the x15 namespace extension element.
type QueryTableRoot struct {
	*openxml.PartRootElementBase
}

// NewQueryTableRoot creates a new QueryTableRoot element.
func NewQueryTableRoot() *QueryTableRoot {
	elem := openxml.NewPartRootElement(
		NamespaceSML,
		"queryTable",
		PrefixDefault,
	)

	return &QueryTableRoot{
		PartRootElementBase: elem,
	}
}

// Name returns the query table name. Attribute: name.
func (qt *QueryTableRoot) Name() string {
	attr, found := qt.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the query table name. Attribute: name.
func (qt *QueryTableRoot) SetName(name string) {
	if name == "" {
		qt.RemoveAttribute("name", "")

		return
	}
	qt.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// QueryTableRefresh returns the QueryTableRefresh child element, or nil
// if not present.
func (qt *QueryTableRoot) QueryTableRefresh() *QueryTableRefresh {
	elem := qt.GetElement(
		"queryTableRefresh",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if qtr, ok := elem.(*QueryTableRefresh); ok {
		return qtr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &QueryTableRefresh{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateQueryTableRefresh returns the QueryTableRefresh child
// element, creating it if needed.
func (qt *QueryTableRoot) GetOrCreateQueryTableRefresh() *QueryTableRefresh {
	qtr := qt.QueryTableRefresh()
	if qtr != nil {
		return qtr
	}
	qtr = NewQueryTableRefresh()
	qt.AppendChild(qtr)

	return qtr
}

// RemoveQueryTableRefresh removes the QueryTableRefresh child element.
func (qt *QueryTableRoot) RemoveQueryTableRefresh() bool {
	qtr := qt.QueryTableRefresh()
	if qtr == nil {
		return false
	}

	return qt.RemoveChild(qtr)
}

// Clone creates a deep copy of this QueryTableRoot element.
func (qt *QueryTableRoot) Clone() openxml.Element {
	cloned := qt.PartRootElementBase.Clone()

	return &QueryTableRoot{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this QueryTableRoot element.
func (qt *QueryTableRoot) CloneNode(
	deep bool,
) openxml.Element {
	cloned := qt.PartRootElementBase.CloneNode(
		deep,
	)

	return &QueryTableRoot{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}
