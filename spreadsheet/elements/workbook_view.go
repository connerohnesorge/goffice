package elements

// revive:disable:file-length-limit SpreadsheetML workbook view has many attrs

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

const (
	// defaultTabRatio is the default tab ratio value (600 = 60%).
	defaultTabRatio = 600
)

// WorkbookView represents a single workbook view element (x:workbookView).
type WorkbookView struct {
	*openxml.CompositeElementBase
}

// NewWorkbookView creates a new WorkbookView element.
func NewWorkbookView() *WorkbookView {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"workbookView",
		PrefixDefault,
	)

	return &WorkbookView{
		CompositeElementBase: elem,
	}
}

// Visibility returns the visibility state of the workbook view.
func (wv *WorkbookView) Visibility() WorkbookViewVisibility {
	attr, found := wv.GetAttribute(
		"visibility",
		"",
	)
	if !found {
		return WorkbookViewVisible
	}

	return WorkbookViewVisibility(attr.Value())
}

// SetVisibility sets the visibility state of the workbook view.
func (wv *WorkbookView) SetVisibility(
	visibility WorkbookViewVisibility,
) {
	if visibility == "" ||
		visibility == WorkbookViewVisible {
		wv.RemoveAttribute("visibility", "")

		return
	}
	wv.SetAttribute(
		openxml.NewAttribute(
			"",
			"visibility",
			"",
			string(visibility),
		),
	)
}

// Minimized returns whether the workbook window is minimized.
func (wv *WorkbookView) Minimized() bool {
	attr, found := wv.GetAttribute(
		"minimized",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetMinimized sets whether the workbook window is minimized.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wv *WorkbookView) SetMinimized(value bool) {
	if value {
		wv.SetAttribute(
			openxml.NewAttribute(
				"",
				"minimized",
				"",
				attrValueTrue,
			),
		)
	} else {
		wv.RemoveAttribute("minimized", "")
	}
}

// ShowHorizontalScroll returns whether the horizontal scroll bar is shown.
func (wv *WorkbookView) ShowHorizontalScroll() bool {
	attr, found := wv.GetAttribute(
		"showHorizontalScroll",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowHorizontalScroll sets whether the horizontal scroll bar is shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wv *WorkbookView) SetShowHorizontalScroll(
	value bool,
) {
	if value {
		wv.RemoveAttribute(
			"showHorizontalScroll",
			"",
		)
	} else {
		wv.SetAttribute(
			openxml.NewAttribute(
				"",
				"showHorizontalScroll",
				"",
				attrValueFalse,
			),
		)
	}
}

// ShowVerticalScroll returns whether the vertical scroll bar is shown.
func (wv *WorkbookView) ShowVerticalScroll() bool {
	attr, found := wv.GetAttribute(
		"showVerticalScroll",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowVerticalScroll sets whether the vertical scroll bar is shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wv *WorkbookView) SetShowVerticalScroll(
	value bool,
) {
	if value {
		wv.RemoveAttribute(
			"showVerticalScroll",
			"",
		)
	} else {
		wv.SetAttribute(
			openxml.NewAttribute(
				"",
				"showVerticalScroll",
				"",
				attrValueFalse,
			),
		)
	}
}

// ShowSheetTabs returns whether sheet tabs are shown.
func (wv *WorkbookView) ShowSheetTabs() bool {
	attr, found := wv.GetAttribute(
		"showSheetTabs",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowSheetTabs sets whether sheet tabs are shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wv *WorkbookView) SetShowSheetTabs(
	value bool,
) {
	if value {
		wv.RemoveAttribute("showSheetTabs", "")
	} else {
		wv.SetAttribute(
			openxml.NewAttribute(
				"",
				"showSheetTabs",
				"",
				attrValueFalse,
			),
		)
	}
}

// XWindow returns the X coordinate of the upper left corner of the window.
func (wv *WorkbookView) XWindow() int {
	attr, found := wv.GetAttribute("xWindow", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetXWindow sets the X coordinate of the upper left corner of the window.
func (wv *WorkbookView) SetXWindow(x int) {
	wv.SetAttribute(
		openxml.NewAttribute(
			"",
			"xWindow",
			"",
			strconv.Itoa(x),
		),
	)
}

// YWindow returns the Y coordinate of the upper left corner of the window.
func (wv *WorkbookView) YWindow() int {
	attr, found := wv.GetAttribute("yWindow", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetYWindow sets the Y coordinate of the upper left corner of the window.
func (wv *WorkbookView) SetYWindow(y int) {
	wv.SetAttribute(
		openxml.NewAttribute(
			"",
			"yWindow",
			"",
			strconv.Itoa(y),
		),
	)
}

// WindowWidth returns the width of the window.
func (wv *WorkbookView) WindowWidth() int {
	attr, found := wv.GetAttribute(
		"windowWidth",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetWindowWidth sets the width of the window.
func (wv *WorkbookView) SetWindowWidth(
	width int,
) {
	wv.SetAttribute(
		openxml.NewAttribute(
			"",
			"windowWidth",
			"",
			strconv.Itoa(width),
		),
	)
}

// WindowHeight returns the height of the window.
func (wv *WorkbookView) WindowHeight() int {
	attr, found := wv.GetAttribute(
		"windowHeight",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetWindowHeight sets the height of the window.
func (wv *WorkbookView) SetWindowHeight(
	height int,
) {
	wv.SetAttribute(
		openxml.NewAttribute(
			"",
			"windowHeight",
			"",
			strconv.Itoa(height),
		),
	)
}

// SetWindowBounds sets all window position and size values at once.
func (wv *WorkbookView) SetWindowBounds(
	x, y, width, height int,
) {
	wv.SetXWindow(x)
	wv.SetYWindow(y)
	wv.SetWindowWidth(width)
	wv.SetWindowHeight(height)
}

// TabRatio returns the ratio of the sheet tab bar width to the horizontal
// scroll bar width. Value is expressed as a percentage (0-1000), where 600
// means 60%.
func (wv *WorkbookView) TabRatio() int {
	attr, found := wv.GetAttribute("tabRatio", "")
	if !found {
		return defaultTabRatio
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetTabRatio sets the tab ratio (0-1000).
func (wv *WorkbookView) SetTabRatio(ratio int) {
	if ratio == defaultTabRatio {
		wv.RemoveAttribute("tabRatio", "")

		return
	}
	wv.SetAttribute(
		openxml.NewAttribute(
			"",
			"tabRatio",
			"",
			strconv.Itoa(ratio),
		),
	)
}

// FirstSheet returns the index of the first visible sheet.
func (wv *WorkbookView) FirstSheet() int {
	attr, found := wv.GetAttribute(
		"firstSheet",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetFirstSheet sets the index of the first visible sheet.
func (wv *WorkbookView) SetFirstSheet(index int) {
	if index == 0 {
		wv.RemoveAttribute("firstSheet", "")

		return
	}
	wv.SetAttribute(
		openxml.NewAttribute(
			"",
			"firstSheet",
			"",
			strconv.Itoa(index),
		),
	)
}

// ActiveTab returns the index of the active (selected) sheet.
func (wv *WorkbookView) ActiveTab() int {
	attr, found := wv.GetAttribute(
		"activeTab",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetActiveTab sets the index of the active sheet.
func (wv *WorkbookView) SetActiveTab(index int) {
	if index == 0 {
		wv.RemoveAttribute("activeTab", "")

		return
	}
	wv.SetAttribute(
		openxml.NewAttribute(
			"",
			"activeTab",
			"",
			strconv.Itoa(index),
		),
	)
}

// AutoFilterDateGrouping returns whether automatic date grouping is applied.
func (wv *WorkbookView) AutoFilterDateGrouping() bool {
	attr, found := wv.GetAttribute(
		"autoFilterDateGrouping",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetAutoFilterDateGrouping sets whether automatic date grouping is applied.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wv *WorkbookView) SetAutoFilterDateGrouping(
	value bool,
) {
	if value {
		wv.RemoveAttribute(
			"autoFilterDateGrouping",
			"",
		)
	} else {
		wv.SetAttribute(
			openxml.NewAttribute(
				"",
				"autoFilterDateGrouping",
				"",
				attrValueFalse,
			),
		)
	}
}

// Clone creates a deep copy of this WorkbookView element.
func (wv *WorkbookView) Clone() openxml.Element {
	cloned := wv.CompositeElementBase.Clone()

	return &WorkbookView{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this WorkbookView element.
func (wv *WorkbookView) CloneNode(
	deep bool,
) openxml.Element {
	cloned := wv.CompositeElementBase.CloneNode(
		deep,
	)

	return &WorkbookView{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
