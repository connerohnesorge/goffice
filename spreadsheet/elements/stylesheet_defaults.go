package elements

import "github.com/connerohnesorge/goffice/openxml"

// Stylesheet element ordering constants per ECMA-376 Part 1.
const (
	stylesheetOrderNumFmts = iota
	stylesheetOrderFonts
	stylesheetOrderFills
	stylesheetOrderBorders
	stylesheetOrderCellStyleXfs
	stylesheetOrderCellXfs
	stylesheetOrderCellStyles
	stylesheetOrderDxfs
	stylesheetOrderTableStyles
	stylesheetOrderColors
)

// Stylesheet default initialization constants.
const (
	// defaultFontSize is the default font size (11pt).
	defaultFontSize = 11
	// defaultFontName is the default font name.
	defaultFontName = "Calibri"
	// defaultFontFamily is the default font family (Swiss).
	defaultFontFamily = 2
)

// insertStylesheetElement inserts an element in the correct order.
// Order defined by ECMA-376 Part 1: numFmts, fonts, fills, borders,
// cellStyleXfs, cellXfs, cellStyles, dxfs, tableStyles, colors.
func (s *Stylesheet) insertStylesheetElement(
	elem openxml.Element,
	name string,
) {
	order := map[string]int{
		"numFmts":      stylesheetOrderNumFmts,
		"fonts":        stylesheetOrderFonts,
		"fills":        stylesheetOrderFills,
		"borders":      stylesheetOrderBorders,
		"cellStyleXfs": stylesheetOrderCellStyleXfs,
		"cellXfs":      stylesheetOrderCellXfs,
		"cellStyles":   stylesheetOrderCellStyles,
		"dxfs":         stylesheetOrderDxfs,
		"tableStyles":  stylesheetOrderTableStyles,
		"colors":       stylesheetOrderColors,
	}

	targetOrder := order[name]

	// Find the first element that should come after this one
	for child := range s.Children() {
		if childOrder, ok := order[child.LocalName()]; ok &&
			childOrder > targetOrder {
			s.InsertBefore(elem, child)

			return
		}
	}

	// No element found that should come after, so append
	s.AppendChild(elem)
}

// InitializeDefaults initializes the stylesheet with default values
// required for a valid Excel file.
//
//nolint:revive // function-length: initialization requires setting up all default styles
func (s *Stylesheet) InitializeDefaults() {
	// Add default font (required)
	fonts := s.GetOrCreateFonts()
	if fonts.ItemCount() == 0 {
		font := fonts.AddFont()
		font.SetFontSize(defaultFontSize)
		font.SetFontName(defaultFontName)
		font.SetFontFamily(defaultFontFamily)
		font.SetFontScheme(FontSchemeMinor)
	}

	// Add default fills (required - at least 2)
	fills := s.GetOrCreateFills()
	if fills.ItemCount() == 0 {
		// First fill: none
		fill1 := fills.AddFill()
		pf1 := fill1.GetOrCreatePatternFill()
		pf1.SetPatternType(PatternTypeNone)

		// Second fill: gray125
		fill2 := fills.AddFill()
		pf2 := fill2.GetOrCreatePatternFill()
		pf2.SetPatternType(PatternTypeGray125)
	}

	// Add default border (required)
	borders := s.GetOrCreateBorders()
	if borders.ItemCount() == 0 {
		borders.AddBorder()
	}

	// Add default cell style xf (required)
	cellStyleXfs := s.GetOrCreateCellStyleXfs()
	if cellStyleXfs.ItemCount() == 0 {
		cellStyleXfs.AddXf()
	}

	// Add default cell xf (required)
	cellXfs := s.GetOrCreateCellXfs()
	if cellXfs.ItemCount() == 0 {
		cellXfs.AddXf()
	}

	// Add default cell style (required)
	cellStyles := s.GetOrCreateCellStyles()
	if cellStyles.ItemCount() != 0 {
		return
	}
	style := cellStyles.AddCellStyle()
	style.SetName("Normal")
	style.SetXfId(0)
	style.SetBuiltinId(0)
}
