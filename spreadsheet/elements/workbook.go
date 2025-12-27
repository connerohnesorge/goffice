//nolint:revive // comments-density
package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Workbook represents the root workbook element (x:workbook).
// This is a PartRootElement since it's the root of the workbook.xml part.
type Workbook struct {
	*openxml.PartRootElementBase
}

// NewWorkbook creates a new Workbook element.
func NewWorkbook() *Workbook {
	elem := openxml.NewPartRootElement(
		NamespaceSML,
		"workbook",
		PrefixDefault,
	)

	return &Workbook{PartRootElementBase: elem}
}

// FileVersion returns the file version element, or nil if not present.
func (wb *Workbook) FileVersion() *FileVersion {
	elem := wb.GetElement(
		"fileVersion",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if fv, ok := elem.(*FileVersion); ok {
		return fv
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &FileVersion{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateFileVersion returns the file version element, creating if needed.
func (wb *Workbook) GetOrCreateFileVersion() *FileVersion {
	fv := wb.FileVersion()
	if fv != nil {
		return fv
	}
	fv = NewFileVersion()
	// FileVersion should be first
	if first := wb.FirstChild(); first != nil {
		wb.InsertBefore(fv, first)
	} else {
		wb.AppendChild(fv)
	}

	return fv
}

// WorkbookPr returns the workbook properties element, or nil if not present.
func (wb *Workbook) WorkbookPr() *WorkbookPr {
	elem := wb.GetElement(
		"workbookPr",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if wpr, ok := elem.(*WorkbookPr); ok {
		return wpr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &WorkbookPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateWorkbookPr returns the workbook properties, creating if needed.
func (wb *Workbook) GetOrCreateWorkbookPr() *WorkbookPr {
	wpr := wb.WorkbookPr()
	if wpr != nil {
		return wpr
	}
	wpr = NewWorkbookPr()
	// WorkbookPr comes after FileVersion
	fv := wb.FileVersion()
	if fv != nil {
		wb.InsertAfter(wpr, fv)
	} else if first := wb.FirstChild(); first != nil {
		wb.InsertBefore(wpr, first)
	} else {
		wb.AppendChild(wpr)
	}

	return wpr
}

// WorkbookProtection returns the workbook protection element, or nil if
// not present.
func (wb *Workbook) WorkbookProtection() *WorkbookProtection {
	elem := wb.GetElement(
		"workbookProtection",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if wp, ok := elem.(*WorkbookProtection); ok {
		return wp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &WorkbookProtection{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateWorkbookProtection returns the workbook protection, creating
// if needed.
func (wb *Workbook) GetOrCreateWorkbookProtection() *WorkbookProtection {
	wp := wb.WorkbookProtection()
	if wp != nil {
		return wp
	}
	wp = NewWorkbookProtection()
	wb.AppendChild(wp)

	return wp
}

// BookViews returns the book views element, or nil if not present.
func (wb *Workbook) BookViews() *BookViews {
	elem := wb.GetElement(
		"bookViews",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if bv, ok := elem.(*BookViews); ok {
		return bv
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &BookViews{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateBookViews returns the book views, creating if needed.
func (wb *Workbook) GetOrCreateBookViews() *BookViews {
	bv := wb.BookViews()
	if bv != nil {
		return bv
	}
	bv = NewBookViews()
	wb.AppendChild(bv)

	return bv
}

// Sheets returns the sheets element, or nil if not present.
func (wb *Workbook) Sheets() *Sheets {
	elem := wb.GetElement("sheets", NamespaceSML)
	if elem == nil {
		elem = wb.GetElement("sheets", "")
	}
	if elem == nil {
		return nil
	}
	if s, ok := elem.(*Sheets); ok {
		return s
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Sheets{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateSheets returns the sheets element, creating if needed.
func (wb *Workbook) GetOrCreateSheets() *Sheets {
	s := wb.Sheets()
	if s != nil {
		return s
	}
	s = NewSheets()
	wb.AppendChild(s)

	return s
}

// DefinedNames returns the defined names element, or nil if not present.
func (wb *Workbook) DefinedNames() *DefinedNames {
	elem := wb.GetElement(
		"definedNames",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if dn, ok := elem.(*DefinedNames); ok {
		return dn
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DefinedNames{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDefinedNames returns the defined names, creating if needed.
func (wb *Workbook) GetOrCreateDefinedNames() *DefinedNames {
	dn := wb.DefinedNames()
	if dn != nil {
		return dn
	}
	dn = NewDefinedNames()
	wb.AppendChild(dn)

	return dn
}

// CalcPr returns the calculation properties element, or nil if not present.
func (wb *Workbook) CalcPr() *CalcPr {
	elem := wb.GetElement("calcPr", NamespaceSML)
	if elem == nil {
		return nil
	}
	if cp, ok := elem.(*CalcPr); ok {
		return cp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CalcPr{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateCalcPr returns the calculation properties, creating if needed.
func (wb *Workbook) GetOrCreateCalcPr() *CalcPr {
	cp := wb.CalcPr()
	if cp != nil {
		return cp
	}
	cp = NewCalcPr()
	wb.AppendChild(cp)

	return cp
}

// Clone creates a deep copy of this Workbook element.
func (wb *Workbook) Clone() openxml.Element {
	cloned := wb.PartRootElementBase.Clone()

	return &Workbook{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this Workbook element.
func (wb *Workbook) CloneNode(
	deep bool,
) openxml.Element {
	cloned := wb.PartRootElementBase.CloneNode(
		deep,
	)

	return &Workbook{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}
