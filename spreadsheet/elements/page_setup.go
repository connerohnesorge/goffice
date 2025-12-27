package elements

//revive:disable:file-length-limit SpreadsheetML page setup has many attributes

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Orientation represents the page orientation.
type Orientation string

const (
	// OrientationDefault is the default orientation.
	OrientationDefault Orientation = "default"
	// OrientationPortrait is portrait orientation.
	OrientationPortrait Orientation = "portrait"
	// OrientationLandscape is landscape orientation.
	OrientationLandscape Orientation = "landscape"
)

// PageOrder represents the page print order.
type PageOrder string

const (
	// PageOrderDownThenOver prints down then over.
	PageOrderDownThenOver PageOrder = "downThenOver"
	// PageOrderOverThenDown prints over then down.
	PageOrderOverThenDown PageOrder = "overThenDown"
)

// CellComments represents how cell comments are printed.
type CellComments string

const (
	// CellCommentsNone does not print comments.
	CellCommentsNone CellComments = "none"
	// CellCommentsAsDisplayed prints comments as displayed.
	CellCommentsAsDisplayed CellComments = "asDisplayed"
	// CellCommentsAtEnd prints comments at end.
	CellCommentsAtEnd CellComments = "atEnd"
)

// PrintErrors represents how errors are printed.
type PrintErrors string

const (
	// PrintErrorsDisplayed prints errors as displayed.
	PrintErrorsDisplayed PrintErrors = "displayed"
	// PrintErrorsBlank prints errors as blank.
	PrintErrorsBlank PrintErrors = "blank"
	// PrintErrorsDash prints errors as dash.
	PrintErrorsDash PrintErrors = "dash"
	// PrintErrorsNA prints errors as #N/A.
	PrintErrorsNA PrintErrors = "NA"
)

// PageSetup represents the page setup element (x:pageSetup).
type PageSetup struct {
	*openxml.CompositeElementBase
}

// NewPageSetup creates a new PageSetup element.
func NewPageSetup() *PageSetup {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"pageSetup",
		PrefixDefault,
	)

	return &PageSetup{CompositeElementBase: elem}
}

// PaperSize returns the paper size index.
// Common values: 1=Letter, 9=A4, 5=Legal, 8=A3, 11=A5
func (ps *PageSetup) PaperSize() int {
	attr, found := ps.GetAttribute(
		"paperSize",
		"",
	)
	if !found {
		return 1 // Default is Letter
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetPaperSize sets the paper size index.
func (ps *PageSetup) SetPaperSize(size int) {
	if size == 1 {
		ps.RemoveAttribute("paperSize", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"paperSize",
			"",
			strconv.Itoa(size),
		),
	)
}

// Scale returns the print scale percentage (10-400).
func (ps *PageSetup) Scale() int {
	attr, found := ps.GetAttribute("scale", "")
	if !found {
		return 100 //nolint:revive // add-constant: default scale 100%
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetScale sets the print scale percentage.
func (ps *PageSetup) SetScale(scale int) {
	if scale == 100 { //nolint:revive // add-constant: default scale 100%
		ps.RemoveAttribute("scale", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"scale",
			"",
			strconv.Itoa(scale),
		),
	)
}

// FirstPageNumber returns the first page number.
func (ps *PageSetup) FirstPageNumber() int {
	attr, found := ps.GetAttribute(
		"firstPageNumber",
		"",
	)
	if !found {
		return 1
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetFirstPageNumber sets the first page number.
func (ps *PageSetup) SetFirstPageNumber(num int) {
	if num == 1 {
		ps.RemoveAttribute("firstPageNumber", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"firstPageNumber",
			"",
			strconv.Itoa(num),
		),
	)
}

// FitToWidth returns the number of horizontal pages to fit to.
func (ps *PageSetup) FitToWidth() int {
	attr, found := ps.GetAttribute(
		"fitToWidth",
		"",
	)
	if !found {
		return 1
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetFitToWidth sets the number of horizontal pages to fit to.
func (ps *PageSetup) SetFitToWidth(pages int) {
	if pages == 1 {
		ps.RemoveAttribute("fitToWidth", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"fitToWidth",
			"",
			strconv.Itoa(pages),
		),
	)
}

// FitToHeight returns the number of vertical pages to fit to.
func (ps *PageSetup) FitToHeight() int {
	attr, found := ps.GetAttribute(
		"fitToHeight",
		"",
	)
	if !found {
		return 1
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetFitToHeight sets the number of vertical pages to fit to.
func (ps *PageSetup) SetFitToHeight(pages int) {
	if pages == 1 {
		ps.RemoveAttribute("fitToHeight", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"fitToHeight",
			"",
			strconv.Itoa(pages),
		),
	)
}

// PageOrder returns the page print order.
func (ps *PageSetup) PageOrder() PageOrder {
	attr, found := ps.GetAttribute(
		"pageOrder",
		"",
	)
	if !found {
		return PageOrderDownThenOver
	}

	return PageOrder(attr.Value())
}

// SetPageOrder sets the page print order.
func (ps *PageSetup) SetPageOrder(
	order PageOrder,
) {
	if order == "" ||
		order == PageOrderDownThenOver {
		ps.RemoveAttribute("pageOrder", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"pageOrder",
			"",
			string(order),
		),
	)
}

// Orientation returns the page orientation.
func (ps *PageSetup) Orientation() Orientation {
	attr, found := ps.GetAttribute(
		"orientation",
		"",
	)
	if !found {
		return OrientationDefault
	}

	return Orientation(attr.Value())
}

// SetOrientation sets the page orientation.
func (ps *PageSetup) SetOrientation(
	orientation Orientation,
) {
	if orientation == "" ||
		orientation == OrientationDefault {
		ps.RemoveAttribute("orientation", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"orientation",
			"",
			string(orientation),
		),
	)
}

// UsePrinterDefaults returns whether printer defaults are used.
func (ps *PageSetup) UsePrinterDefaults() bool {
	attr, found := ps.GetAttribute(
		"usePrinterDefaults",
		"",
	)
	if !found {
		return true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetUsePrinterDefaults sets whether printer defaults are used.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ps *PageSetup) SetUsePrinterDefaults(
	value bool,
) {
	if value {
		ps.RemoveAttribute(
			"usePrinterDefaults",
			"",
		)
	} else {
		ps.SetAttribute(openxml.NewAttribute("", "usePrinterDefaults", "", attrValueFalse))
	}
}

// BlackAndWhite returns whether to print in black and white.
func (ps *PageSetup) BlackAndWhite() bool {
	attr, found := ps.GetAttribute(
		"blackAndWhite",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetBlackAndWhite sets whether to print in black and white.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ps *PageSetup) SetBlackAndWhite(
	value bool,
) {
	if value {
		ps.SetAttribute(
			openxml.NewAttribute(
				"",
				"blackAndWhite",
				"",
				attrValueTrue,
			),
		)
	} else {
		ps.RemoveAttribute("blackAndWhite", "")
	}
}

// Draft returns whether to print in draft mode.
func (ps *PageSetup) Draft() bool {
	attr, found := ps.GetAttribute("draft", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDraft sets whether to print in draft mode.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ps *PageSetup) SetDraft(value bool) {
	if value {
		ps.SetAttribute(
			openxml.NewAttribute(
				"",
				"draft",
				"",
				attrValueTrue,
			),
		)
	} else {
		ps.RemoveAttribute("draft", "")
	}
}

// CellComments returns how cell comments are printed.
func (ps *PageSetup) CellComments() CellComments {
	attr, found := ps.GetAttribute(
		"cellComments",
		"",
	)
	if !found {
		return CellCommentsNone
	}

	return CellComments(attr.Value())
}

// SetCellComments sets how cell comments are printed.
func (ps *PageSetup) SetCellComments(
	comments CellComments,
) {
	if comments == "" ||
		comments == CellCommentsNone {
		ps.RemoveAttribute("cellComments", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"cellComments",
			"",
			string(comments),
		),
	)
}

// UseFirstPageNumber returns whether to use the custom first page number.
func (ps *PageSetup) UseFirstPageNumber() bool {
	attr, found := ps.GetAttribute(
		"useFirstPageNumber",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetUseFirstPageNumber sets whether to use the custom first page number.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (ps *PageSetup) SetUseFirstPageNumber(
	value bool,
) {
	if value {
		ps.SetAttribute(
			openxml.NewAttribute(
				"",
				"useFirstPageNumber",
				"",
				attrValueTrue,
			),
		)
	} else {
		ps.RemoveAttribute("useFirstPageNumber", "")
	}
}

// Errors returns how errors are printed.
func (ps *PageSetup) Errors() PrintErrors {
	attr, found := ps.GetAttribute("errors", "")
	if !found {
		return PrintErrorsDisplayed
	}

	return PrintErrors(attr.Value())
}

// SetErrors sets how errors are printed.
func (ps *PageSetup) SetErrors(
	errors PrintErrors,
) {
	if errors == "" ||
		errors == PrintErrorsDisplayed {
		ps.RemoveAttribute("errors", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"errors",
			"",
			string(errors),
		),
	)
}

// HorizontalDpi returns the horizontal print resolution.
func (ps *PageSetup) HorizontalDpi() int {
	attr, found := ps.GetAttribute(
		"horizontalDpi",
		"",
	)
	if !found {
		return 600 //nolint:revive // add-constant: default DPI
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetHorizontalDpi sets the horizontal print resolution.
func (ps *PageSetup) SetHorizontalDpi(dpi int) {
	if dpi == 600 { //nolint:revive // add-constant: default DPI
		ps.RemoveAttribute("horizontalDpi", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"horizontalDpi",
			"",
			strconv.Itoa(dpi),
		),
	)
}

// VerticalDpi returns the vertical print resolution.
func (ps *PageSetup) VerticalDpi() int {
	attr, found := ps.GetAttribute(
		"verticalDpi",
		"",
	)
	if !found {
		return 600 //nolint:revive // add-constant: default DPI
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetVerticalDpi sets the vertical print resolution.
func (ps *PageSetup) SetVerticalDpi(dpi int) {
	if dpi == 600 { //nolint:revive // add-constant: default DPI
		ps.RemoveAttribute("verticalDpi", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"verticalDpi",
			"",
			strconv.Itoa(dpi),
		),
	)
}

// Copies returns the number of copies to print.
func (ps *PageSetup) Copies() int {
	attr, found := ps.GetAttribute("copies", "")
	if !found {
		return 1
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCopies sets the number of copies to print.
func (ps *PageSetup) SetCopies(copies int) {
	if copies == 1 {
		ps.RemoveAttribute("copies", "")

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			"",
			"copies",
			"",
			strconv.Itoa(copies),
		),
	)
}

// RelationshipId returns the relationship ID for printer settings.
func (ps *PageSetup) RelationshipId() string {
	attr, found := ps.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelationshipId sets the relationship ID for printer settings.
func (ps *PageSetup) SetRelationshipId(
	id string,
) {
	if id == "" {
		ps.RemoveAttribute(
			"id",
			NamespaceRelationships,
		)

		return
	}
	ps.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"id",
			PrefixR,
			id,
		),
	)
}

// Clone creates a deep copy of this PageSetup element.
func (ps *PageSetup) Clone() openxml.Element {
	cloned := ps.CompositeElementBase.Clone()

	return &PageSetup{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PageSetup element.
func (ps *PageSetup) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ps.CompositeElementBase.CloneNode(
		deep,
	)

	return &PageSetup{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
