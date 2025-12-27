//nolint:revive // comments-density
package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Stylesheet represents the stylesheet root element (x:styleSheet).
// This is the root element for styles.xml in an Excel package.
type Stylesheet struct {
	*openxml.PartRootElementBase
}

// NewStylesheet creates a new Stylesheet element.
func NewStylesheet() *Stylesheet {
	elem := openxml.NewPartRootElement(
		NamespaceSML,
		"styleSheet",
		PrefixDefault,
	)

	return &Stylesheet{PartRootElementBase: elem}
}

// NumFmts returns the number formats element, or nil if not present.
func (s *Stylesheet) NumFmts() *NumFmts {
	elem := s.GetElement("numFmts", NamespaceSML)
	if elem == nil {
		return nil
	}
	if nf, ok := elem.(*NumFmts); ok {
		return nf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &NumFmts{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateNumFmts returns the number formats, creating if needed.
func (s *Stylesheet) GetOrCreateNumFmts() *NumFmts {
	nf := s.NumFmts()
	if nf != nil {
		return nf
	}
	nf = NewNumFmts()
	s.insertStylesheetElement(nf, "numFmts")

	return nf
}

// Fonts returns the fonts element, or nil if not present.
func (s *Stylesheet) Fonts() *Fonts {
	elem := s.GetElement("fonts", NamespaceSML)
	if elem == nil {
		return nil
	}
	if f, ok := elem.(*Fonts); ok {
		return f
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Fonts{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateFonts returns the fonts, creating if needed.
func (s *Stylesheet) GetOrCreateFonts() *Fonts {
	f := s.Fonts()
	if f != nil {
		return f
	}
	f = NewFonts()
	s.insertStylesheetElement(f, "fonts")

	return f
}

// Fills returns the fills element, or nil if not present.
func (s *Stylesheet) Fills() *Fills {
	elem := s.GetElement("fills", NamespaceSML)
	if elem == nil {
		return nil
	}
	if f, ok := elem.(*Fills); ok {
		return f
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Fills{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateFills returns the fills, creating if needed.
func (s *Stylesheet) GetOrCreateFills() *Fills {
	f := s.Fills()
	if f != nil {
		return f
	}
	f = NewFills()
	s.insertStylesheetElement(f, "fills")

	return f
}

// Borders returns the borders element, or nil if not present.
func (s *Stylesheet) Borders() *Borders {
	elem := s.GetElement("borders", NamespaceSML)
	if elem == nil {
		return nil
	}
	if b, ok := elem.(*Borders); ok {
		return b
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Borders{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateBorders returns the borders, creating if needed.
func (s *Stylesheet) GetOrCreateBorders() *Borders {
	b := s.Borders()
	if b != nil {
		return b
	}
	b = NewBorders()
	s.insertStylesheetElement(b, "borders")

	return b
}

// CellStyleXfs returns the cell style formats element, or nil if not present.
func (s *Stylesheet) CellStyleXfs() *CellStyleXfs {
	elem := s.GetElement(
		"cellStyleXfs",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if csxf, ok := elem.(*CellStyleXfs); ok {
		return csxf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CellStyleXfs{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCellStyleXfs returns the cell style formats, creating if needed.
func (s *Stylesheet) GetOrCreateCellStyleXfs() *CellStyleXfs {
	csxf := s.CellStyleXfs()
	if csxf != nil {
		return csxf
	}
	csxf = NewCellStyleXfs()
	s.insertStylesheetElement(
		csxf,
		"cellStyleXfs",
	)

	return csxf
}

// CellXfs returns the cell formats element, or nil if not present.
func (s *Stylesheet) CellXfs() *CellXfs {
	elem := s.GetElement("cellXfs", NamespaceSML)
	if elem == nil {
		return nil
	}
	if cxf, ok := elem.(*CellXfs); ok {
		return cxf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CellXfs{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCellXfs returns the cell formats, creating if needed.
func (s *Stylesheet) GetOrCreateCellXfs() *CellXfs {
	cxf := s.CellXfs()
	if cxf != nil {
		return cxf
	}
	cxf = NewCellXfs()
	s.insertStylesheetElement(cxf, "cellXfs")

	return cxf
}

// CellStyles returns the cell styles element, or nil if not present.
func (s *Stylesheet) CellStyles() *CellStyles {
	elem := s.GetElement(
		"cellStyles",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if cs, ok := elem.(*CellStyles); ok {
		return cs
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CellStyles{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCellStyles returns the cell styles, creating if needed.
func (s *Stylesheet) GetOrCreateCellStyles() *CellStyles {
	cs := s.CellStyles()
	if cs != nil {
		return cs
	}
	cs = NewCellStyles()
	s.insertStylesheetElement(cs, "cellStyles")

	return cs
}

// Dxfs returns the differential formats element, or nil if not present.
func (s *Stylesheet) Dxfs() *Dxfs {
	elem := s.GetElement("dxfs", NamespaceSML)
	if elem == nil {
		return nil
	}
	if d, ok := elem.(*Dxfs); ok {
		return d
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Dxfs{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateDxfs returns the differential formats, creating if needed.
func (s *Stylesheet) GetOrCreateDxfs() *Dxfs {
	d := s.Dxfs()
	if d != nil {
		return d
	}
	d = NewDxfs()
	s.insertStylesheetElement(d, "dxfs")

	return d
}

// TableStyles returns the table styles element, or nil if not present.
func (s *Stylesheet) TableStyles() *TableStyles {
	elem := s.GetElement(
		"tableStyles",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if ts, ok := elem.(*TableStyles); ok {
		return ts
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableStyles{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTableStyles returns the table styles, creating if needed.
func (s *Stylesheet) GetOrCreateTableStyles() *TableStyles {
	ts := s.TableStyles()
	if ts != nil {
		return ts
	}
	ts = NewTableStyles()
	s.insertStylesheetElement(ts, "tableStyles")

	return ts
}

// Colors returns the colors element, or nil if not present.
func (s *Stylesheet) Colors() *Colors {
	elem := s.GetElement("colors", NamespaceSML)
	if elem == nil {
		return nil
	}
	if c, ok := elem.(*Colors); ok {
		return c
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Colors{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateColors returns the colors, creating if needed.
func (s *Stylesheet) GetOrCreateColors() *Colors {
	c := s.Colors()
	if c != nil {
		return c
	}
	c = NewColors()
	s.insertStylesheetElement(c, "colors")

	return c
}

// Clone creates a deep copy of this Stylesheet element.
func (s *Stylesheet) Clone() openxml.Element {
	cloned := s.PartRootElementBase.Clone()

	return &Stylesheet{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this Stylesheet element.
func (s *Stylesheet) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.PartRootElementBase.CloneNode(
		deep,
	)

	return &Stylesheet{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}
