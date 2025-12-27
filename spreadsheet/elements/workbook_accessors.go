//nolint:revive // comments-density
package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// PivotCaches returns the pivot caches element, or nil if not present.
func (wb *Workbook) PivotCaches() *PivotCaches {
	elem := wb.GetElement(
		"pivotCaches",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if pc, ok := elem.(*PivotCaches); ok {
		return pc
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PivotCaches{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePivotCaches returns the pivot caches, creating if needed.
func (wb *Workbook) GetOrCreatePivotCaches() *PivotCaches {
	pc := wb.PivotCaches()
	if pc != nil {
		return pc
	}
	pc = NewPivotCaches()
	wb.AppendChild(pc)

	return pc
}

// FileRecoveryPr returns the file recovery properties element, or nil if
// not present.
func (wb *Workbook) FileRecoveryPr() *FileRecoveryPr {
	elem := wb.GetElement(
		"fileRecoveryPr",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if fr, ok := elem.(*FileRecoveryPr); ok {
		return fr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &FileRecoveryPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateFileRecoveryPr returns the file recovery properties,
// creating if needed.
func (wb *Workbook) GetOrCreateFileRecoveryPr() *FileRecoveryPr {
	fr := wb.FileRecoveryPr()
	if fr != nil {
		return fr
	}
	fr = NewFileRecoveryPr()
	wb.AppendChild(fr)

	return fr
}

// WebPublishing returns the web publishing element, or nil if not present.
func (wb *Workbook) WebPublishing() *WebPublishing {
	elem := wb.GetElement(
		"webPublishing",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if wp, ok := elem.(*WebPublishing); ok {
		return wp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &WebPublishing{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateWebPublishing returns the web publishing element,
// creating if needed.
func (wb *Workbook) GetOrCreateWebPublishing() *WebPublishing {
	wp := wb.WebPublishing()
	if wp != nil {
		return wp
	}
	wp = NewWebPublishing()
	wb.AppendChild(wp)

	return wp
}
