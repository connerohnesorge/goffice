package elements

//revive:disable:file-length-limit many sheet view properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// SheetViewType represents the type of view for a worksheet.
type SheetViewType string

const (
	// SheetViewNormal indicates normal view.
	SheetViewNormal SheetViewType = "normal"
	// SheetViewPageBreakPreview indicates page break preview.
	SheetViewPageBreakPreview SheetViewType = "pageBreakPreview"
	// SheetViewPageLayout indicates page layout view.
	SheetViewPageLayout SheetViewType = "pageLayout"
)

// Default zoom scale constants.
const (
	defaultZoomScale = 100
)

// SheetViews represents the sheet views container element (x:sheetViews).
type SheetViews struct {
	*openxml.CompositeElementBase
}

// NewSheetViews creates a new SheetViews element.
func NewSheetViews() *SheetViews {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"sheetViews",
		PrefixDefault,
	)

	return &SheetViews{CompositeElementBase: elem}
}

// SheetViews returns an iterator over all SheetView elements.
func (sv *SheetViews) SheetViews() iter.Seq[*SheetView] {
	return func(yield func(*SheetView) bool) {
		for child := range sv.Children() {
			if child.LocalName() != "sheetView" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var view *SheetView
			switch sv := child.(type) {
			case *SheetView:
				view = sv
			case *openxml.CompositeElementBase:
				view = &SheetView{CompositeElementBase: sv}
			}
			if view != nil && !yield(view) {
				return
			}
		}
	}
}

// FirstSheetView returns the first sheet view, or nil if none exist.
func (sv *SheetViews) FirstSheetView() *SheetView {
	for view := range sv.SheetViews() {
		return view
	}

	return nil
}

// GetOrCreateSheetView returns the first sheet view,
// creating one if none exist.
func (sv *SheetViews) GetOrCreateSheetView() *SheetView {
	view := sv.FirstSheetView()
	if view != nil {
		return view
	}
	view = NewSheetView()
	sv.AppendChild(view)

	return view
}

// AddSheetView adds a new sheet view and returns it.
func (sv *SheetViews) AddSheetView() *SheetView {
	view := NewSheetView()
	sv.AppendChild(view)

	return view
}

// Clone creates a deep copy of this SheetViews element.
func (sv *SheetViews) Clone() openxml.Element {
	cloned := sv.CompositeElementBase.Clone()

	return &SheetViews{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SheetViews element.
func (sv *SheetViews) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sv.CompositeElementBase.CloneNode(
		deep,
	)

	return &SheetViews{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// SheetView represents a single sheet view element (x:sheetView).
type SheetView struct {
	*openxml.CompositeElementBase
}

// NewSheetView creates a new SheetView element.
func NewSheetView() *SheetView {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"sheetView",
		PrefixDefault,
	)

	return &SheetView{CompositeElementBase: elem}
}

// WindowProtection returns whether the window is protected.
func (sv *SheetView) WindowProtection() bool {
	attr, found := sv.GetAttribute(
		"windowProtection",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetWindowProtection sets whether the window is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetWindowProtection(
	value bool,
) {
	if value {
		sv.SetAttribute(
			openxml.NewAttribute(
				"",
				"windowProtection",
				"",
				attrValueTrue,
			),
		)
	} else {
		sv.RemoveAttribute("windowProtection", "")
	}
}

// ShowFormulas returns whether formulas are shown instead of values.
func (sv *SheetView) ShowFormulas() bool {
	attr, found := sv.GetAttribute(
		"showFormulas",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowFormulas sets whether formulas are shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetShowFormulas(value bool) {
	if value {
		sv.SetAttribute(
			openxml.NewAttribute(
				"",
				"showFormulas",
				"",
				attrValueTrue,
			),
		)
	} else {
		sv.RemoveAttribute("showFormulas", "")
	}
}

// ShowGridLines returns whether grid lines are shown.
func (sv *SheetView) ShowGridLines() bool {
	attr, found := sv.GetAttribute(
		"showGridLines",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowGridLines sets whether grid lines are shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetShowGridLines(
	value bool,
) {
	if value {
		sv.RemoveAttribute("showGridLines", "")
	} else {
		sv.SetAttribute(openxml.NewAttribute("", "showGridLines", "", attrValueFalse))
	}
}

// ShowRowColHeaders returns whether row and column headers are shown.
func (sv *SheetView) ShowRowColHeaders() bool {
	attr, found := sv.GetAttribute(
		"showRowColHeaders",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowRowColHeaders sets whether row and column headers are shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetShowRowColHeaders(
	value bool,
) {
	if value {
		sv.RemoveAttribute(
			"showRowColHeaders",
			"",
		)
	} else {
		sv.SetAttribute(openxml.NewAttribute("", "showRowColHeaders", "", attrValueFalse))
	}
}

// ShowZeros returns whether zero values are shown.
func (sv *SheetView) ShowZeros() bool {
	attr, found := sv.GetAttribute(
		"showZeros",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowZeros sets whether zero values are shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetShowZeros(value bool) {
	if value {
		sv.RemoveAttribute("showZeros", "")
	} else {
		sv.SetAttribute(openxml.NewAttribute("", "showZeros", "", attrValueFalse))
	}
}

// RightToLeft returns whether the sheet is displayed right-to-left.
func (sv *SheetView) RightToLeft() bool {
	attr, found := sv.GetAttribute(
		"rightToLeft",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetRightToLeft sets whether the sheet is displayed right-to-left.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetRightToLeft(value bool) {
	if value {
		sv.SetAttribute(
			openxml.NewAttribute(
				"",
				"rightToLeft",
				"",
				attrValueTrue,
			),
		)
	} else {
		sv.RemoveAttribute("rightToLeft", "")
	}
}

// TabSelected returns whether this sheet's tab is selected.
func (sv *SheetView) TabSelected() bool {
	attr, found := sv.GetAttribute(
		"tabSelected",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetTabSelected sets whether this sheet's tab is selected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetTabSelected(value bool) {
	if value {
		sv.SetAttribute(
			openxml.NewAttribute(
				"",
				"tabSelected",
				"",
				attrValueTrue,
			),
		)
	} else {
		sv.RemoveAttribute("tabSelected", "")
	}
}

// ShowRuler returns whether the ruler is shown in page layout view.
func (sv *SheetView) ShowRuler() bool {
	attr, found := sv.GetAttribute(
		"showRuler",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowRuler sets whether the ruler is shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetShowRuler(value bool) {
	if value {
		sv.RemoveAttribute("showRuler", "")
	} else {
		sv.SetAttribute(openxml.NewAttribute("", "showRuler", "", attrValueFalse))
	}
}

// ShowOutlineSymbols returns whether outline symbols are shown.
func (sv *SheetView) ShowOutlineSymbols() bool {
	attr, found := sv.GetAttribute(
		"showOutlineSymbols",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowOutlineSymbols sets whether outline symbols are shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetShowOutlineSymbols(
	value bool,
) {
	if value {
		sv.RemoveAttribute(
			"showOutlineSymbols",
			"",
		)
	} else {
		sv.SetAttribute(openxml.NewAttribute("", "showOutlineSymbols", "", attrValueFalse))
	}
}

// DefaultGridColor returns whether the default grid color is used.
func (sv *SheetView) DefaultGridColor() bool {
	attr, found := sv.GetAttribute(
		"defaultGridColor",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetDefaultGridColor sets whether the default grid color is used.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetDefaultGridColor(
	value bool,
) {
	if value {
		sv.RemoveAttribute("defaultGridColor", "")
	} else {
		sv.SetAttribute(openxml.NewAttribute("", "defaultGridColor", "", attrValueFalse))
	}
}

// ShowWhiteSpace returns whether white space is shown in page layout view.
func (sv *SheetView) ShowWhiteSpace() bool {
	attr, found := sv.GetAttribute(
		"showWhiteSpace",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowWhiteSpace sets whether white space is shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sv *SheetView) SetShowWhiteSpace(
	value bool,
) {
	if value {
		sv.RemoveAttribute("showWhiteSpace", "")
	} else {
		sv.SetAttribute(openxml.NewAttribute("", "showWhiteSpace", "", attrValueFalse))
	}
}

// View returns the view type for this sheet view.
func (sv *SheetView) View() SheetViewType {
	attr, found := sv.GetAttribute("view", "")
	if !found {
		return SheetViewNormal
	}

	return SheetViewType(attr.Value())
}

// SetView sets the view type.
func (sv *SheetView) SetView(
	viewType SheetViewType,
) {
	if viewType == "" ||
		viewType == SheetViewNormal {
		sv.RemoveAttribute("view", "")

		return
	}
	sv.SetAttribute(
		openxml.NewAttribute(
			"",
			"view",
			"",
			string(viewType),
		),
	)
}

// TopLeftCell returns the top-left cell visible in the bottom-right pane.
func (sv *SheetView) TopLeftCell() string {
	attr, found := sv.GetAttribute(
		"topLeftCell",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTopLeftCell sets the top-left cell visible in the bottom-right pane.
func (sv *SheetView) SetTopLeftCell(cell string) {
	if cell == "" {
		sv.RemoveAttribute("topLeftCell", "")

		return
	}
	sv.SetAttribute(
		openxml.NewAttribute(
			"",
			"topLeftCell",
			"",
			cell,
		),
	)
}

// ColorId returns the grid line color index.
func (sv *SheetView) ColorId() int {
	attr, found := sv.GetAttribute("colorId", "")
	if !found {
		return 64 //nolint:revive // add-constant: Default value
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetColorId sets the grid line color index.
func (sv *SheetView) SetColorId(id int) {
	if id == 64 { //nolint:revive // add-constant: default value
		sv.RemoveAttribute("colorId", "")

		return
	}
	sv.SetAttribute(
		openxml.NewAttribute(
			"",
			"colorId",
			"",
			strconv.Itoa(id),
		),
	)
}

// ZoomScale returns the zoom scale percentage (10-400).
func (sv *SheetView) ZoomScale() int {
	attr, found := sv.GetAttribute(
		"zoomScale",
		"",
	)
	if !found {
		return defaultZoomScale
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetZoomScale sets the zoom scale percentage.
func (sv *SheetView) SetZoomScale(scale int) {
	if scale == defaultZoomScale {
		sv.RemoveAttribute("zoomScale", "")

		return
	}
	sv.SetAttribute(
		openxml.NewAttribute(
			"",
			"zoomScale",
			"",
			strconv.Itoa(scale),
		),
	)
}

// ZoomScaleNormal returns the zoom scale for normal view.
func (sv *SheetView) ZoomScaleNormal() int {
	attr, found := sv.GetAttribute(
		"zoomScaleNormal",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetZoomScaleNormal sets the zoom scale for normal view.
func (sv *SheetView) SetZoomScaleNormal(
	scale int,
) {
	if scale == 0 {
		sv.RemoveAttribute("zoomScaleNormal", "")

		return
	}
	sv.SetAttribute(
		openxml.NewAttribute(
			"",
			"zoomScaleNormal",
			"",
			strconv.Itoa(scale),
		),
	)
}

// ZoomScaleSheetLayoutView returns the zoom scale for sheet layout view.
func (sv *SheetView) ZoomScaleSheetLayoutView() int {
	attr, found := sv.GetAttribute(
		"zoomScaleSheetLayoutView",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetZoomScaleSheetLayoutView sets the zoom scale for sheet layout view.
func (sv *SheetView) SetZoomScaleSheetLayoutView(
	scale int,
) {
	if scale == 0 {
		sv.RemoveAttribute(
			"zoomScaleSheetLayoutView",
			"",
		)

		return
	}
	sv.SetAttribute(
		openxml.NewAttribute(
			"",
			"zoomScaleSheetLayoutView",
			"",
			strconv.Itoa(scale),
		),
	)
}

// ZoomScalePageLayoutView returns the zoom scale for page layout view.
func (sv *SheetView) ZoomScalePageLayoutView() int {
	attr, found := sv.GetAttribute(
		"zoomScalePageLayoutView",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetZoomScalePageLayoutView sets the zoom scale for page layout view.
func (sv *SheetView) SetZoomScalePageLayoutView(
	scale int,
) {
	if scale == 0 {
		sv.RemoveAttribute(
			"zoomScalePageLayoutView",
			"",
		)

		return
	}
	sv.SetAttribute(
		openxml.NewAttribute(
			"",
			"zoomScalePageLayoutView",
			"",
			strconv.Itoa(scale),
		),
	)
}

// WorkbookViewId returns the workbook view index (0-based).
func (sv *SheetView) WorkbookViewId() int {
	attr, found := sv.GetAttribute(
		"workbookViewId",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetWorkbookViewId sets the workbook view index.
func (sv *SheetView) SetWorkbookViewId(id int) {
	sv.SetAttribute(
		openxml.NewAttribute(
			"",
			"workbookViewId",
			"",
			strconv.Itoa(id),
		),
	)
}

// Pane returns the pane element, or nil if not present.
func (sv *SheetView) Pane() *Pane {
	elem := sv.GetElement("pane", NamespaceSML)
	if elem == nil {
		return nil
	}
	if p, ok := elem.(*Pane); ok {
		return p
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Pane{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreatePane returns the pane element, creating if needed.
func (sv *SheetView) GetOrCreatePane() *Pane {
	p := sv.Pane()
	if p != nil {
		return p
	}
	p = NewPane()
	// Pane should be first child
	if first := sv.FirstChild(); first != nil {
		sv.InsertBefore(p, first)
	} else {
		sv.AppendChild(p)
	}

	return p
}

// Selection returns the selection element, or nil if not present.
func (sv *SheetView) Selection() *Selection {
	elem := sv.GetElement(
		"selection",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if s, ok := elem.(*Selection); ok {
		return s
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Selection{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSelection returns the selection element, creating if needed.
func (sv *SheetView) GetOrCreateSelection() *Selection {
	s := sv.Selection()
	if s != nil {
		return s
	}
	s = NewSelection()
	sv.AppendChild(s)

	return s
}

// Clone creates a deep copy of this SheetView element.
func (sv *SheetView) Clone() openxml.Element {
	cloned := sv.CompositeElementBase.Clone()

	return &SheetView{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SheetView element.
func (sv *SheetView) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sv.CompositeElementBase.CloneNode(
		deep,
	)

	return &SheetView{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
