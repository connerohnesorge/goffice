package elements

//revive:disable:file-length-limit many workbook properties

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// WorkbookPr represents the workbook properties element (x:workbookPr).
type WorkbookPr struct {
	*openxml.CompositeElementBase
}

// NewWorkbookPr creates a new WorkbookPr element.
func NewWorkbookPr() *WorkbookPr {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"workbookPr",
		PrefixDefault,
	)

	return &WorkbookPr{CompositeElementBase: elem}
}

// Date1904 returns whether the workbook uses the 1904 date system.
// In the 1904 date system, the first date is January 1, 1904.
func (w *WorkbookPr) Date1904() bool {
	attr, found := w.GetAttribute("date1904", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDate1904 sets whether the workbook uses the 1904 date system.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetDate1904(value bool) {
	if value {
		w.SetAttribute(
			openxml.NewAttribute(
				"",
				"date1904",
				"",
				attrValueTrue,
			),
		)
	} else {
		w.RemoveAttribute("date1904", "")
	}
}

// FilterPrivacy returns whether the workbook has its personal
// information filtered.
func (w *WorkbookPr) FilterPrivacy() bool {
	attr, found := w.GetAttribute(
		"filterPrivacy",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetFilterPrivacy sets whether personal information should be filtered.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetFilterPrivacy(
	value bool,
) {
	if value {
		w.SetAttribute(
			openxml.NewAttribute(
				"",
				"filterPrivacy",
				"",
				attrValueTrue,
			),
		)
	} else {
		w.RemoveAttribute("filterPrivacy", "")
	}
}

// DefaultThemeVersion returns the default theme version.
func (w *WorkbookPr) DefaultThemeVersion() int {
	attr, found := w.GetAttribute(
		"defaultThemeVersion",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetDefaultThemeVersion sets the default theme version.
func (w *WorkbookPr) SetDefaultThemeVersion(
	version int,
) {
	if version == 0 {
		w.RemoveAttribute(
			"defaultThemeVersion",
			"",
		)

		return
	}
	verStr := strconv.Itoa(version)
	attr := openxml.NewAttribute(
		"",
		"defaultThemeVersion",
		"",
		verStr,
	)
	w.SetAttribute(attr)
}

// ShowObjects returns the show objects setting.
// Possible values: "all", "placeholders", "none"
func (w *WorkbookPr) ShowObjects() string {
	attr, found := w.GetAttribute(
		"showObjects",
		"",
	)
	if !found {
		return "all" // Default value
	}

	return attr.Value()
}

// SetShowObjects sets the show objects setting.
func (w *WorkbookPr) SetShowObjects(
	value string,
) {
	if value == "" || value == "all" {
		w.RemoveAttribute("showObjects", "")

		return
	}
	w.SetAttribute(
		openxml.NewAttribute(
			"",
			"showObjects",
			"",
			value,
		),
	)
}

// ShowBorderUnselectedTables returns whether borders are shown
// for unselected tables.
func (w *WorkbookPr) ShowBorderUnselectedTables() bool {
	attr, found := w.GetAttribute(
		"showBorderUnselectedTables",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowBorderUnselectedTables sets whether borders are shown for unselected tables.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetShowBorderUnselectedTables(
	value bool,
) {
	if value {
		w.RemoveAttribute(
			"showBorderUnselectedTables",
			"",
		)
	} else {
		w.SetAttribute(openxml.NewAttribute("", "showBorderUnselectedTables", "", attrValueFalse))
	}
}

// PromptedSolutions returns whether prompted solutions are shown.
func (w *WorkbookPr) PromptedSolutions() bool {
	attr, found := w.GetAttribute(
		"promptedSolutions",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPromptedSolutions sets whether prompted solutions are shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetPromptedSolutions(
	value bool,
) {
	if value {
		w.SetAttribute(
			openxml.NewAttribute(
				"",
				"promptedSolutions",
				"",
				attrValueTrue,
			),
		)
	} else {
		w.RemoveAttribute("promptedSolutions", "")
	}
}

// ShowInkAnnotation returns whether ink annotations are shown.
func (w *WorkbookPr) ShowInkAnnotation() bool {
	attr, found := w.GetAttribute(
		"showInkAnnotation",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowInkAnnotation sets whether ink annotations are shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetShowInkAnnotation(
	value bool,
) {
	if value {
		w.RemoveAttribute("showInkAnnotation", "")
	} else {
		w.SetAttribute(openxml.NewAttribute("", "showInkAnnotation", "", attrValueFalse))
	}
}

// BackupFile returns whether a backup file should be created.
func (w *WorkbookPr) BackupFile() bool {
	attr, found := w.GetAttribute(
		"backupFile",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetBackupFile sets whether a backup file should be created.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetBackupFile(value bool) {
	if value {
		w.SetAttribute(
			openxml.NewAttribute(
				"",
				"backupFile",
				"",
				attrValueTrue,
			),
		)
	} else {
		w.RemoveAttribute("backupFile", "")
	}
}

// SaveExternalLinkValues returns whether external link values are saved.
func (w *WorkbookPr) SaveExternalLinkValues() bool {
	attr, found := w.GetAttribute(
		"saveExternalLinkValues",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetSaveExternalLinkValues sets whether external link values are saved.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetSaveExternalLinkValues(
	value bool,
) {
	if value {
		w.RemoveAttribute(
			"saveExternalLinkValues",
			"",
		)
	} else {
		w.SetAttribute(openxml.NewAttribute("", "saveExternalLinkValues", "", attrValueFalse))
	}
}

// UpdateLinks returns the update links behavior.
// Possible values: "userSet", "never", "always"
func (w *WorkbookPr) UpdateLinks() string {
	attr, found := w.GetAttribute(
		"updateLinks",
		"",
	)
	if !found {
		return "userSet" // Default value
	}

	return attr.Value()
}

// SetUpdateLinks sets the update links behavior.
func (w *WorkbookPr) SetUpdateLinks(
	value string,
) {
	if value == "" || value == "userSet" {
		w.RemoveAttribute("updateLinks", "")

		return
	}
	w.SetAttribute(
		openxml.NewAttribute(
			"",
			"updateLinks",
			"",
			value,
		),
	)
}

// CodeName returns the VBA code name of the workbook.
func (w *WorkbookPr) CodeName() string {
	attr, found := w.GetAttribute("codeName", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCodeName sets the VBA code name.
func (w *WorkbookPr) SetCodeName(name string) {
	if name == "" {
		w.RemoveAttribute("codeName", "")

		return
	}
	w.SetAttribute(
		openxml.NewAttribute(
			"",
			"codeName",
			"",
			name,
		),
	)
}

// HidePivotFieldList returns whether the PivotTable field list is hidden.
func (w *WorkbookPr) HidePivotFieldList() bool {
	attr, found := w.GetAttribute(
		"hidePivotFieldList",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHidePivotFieldList sets whether the PivotTable field list is hidden.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetHidePivotFieldList(
	value bool,
) {
	if value {
		w.SetAttribute(
			openxml.NewAttribute(
				"",
				"hidePivotFieldList",
				"",
				attrValueTrue,
			),
		)
	} else {
		w.RemoveAttribute("hidePivotFieldList", "")
	}
}

// ShowPivotChartFilter returns whether PivotChart filter buttons are shown.
func (w *WorkbookPr) ShowPivotChartFilter() bool {
	attr, found := w.GetAttribute(
		"showPivotChartFilter",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetShowPivotChartFilter sets whether PivotChart filter buttons are shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetShowPivotChartFilter(
	value bool,
) {
	if value {
		w.SetAttribute(
			openxml.NewAttribute(
				"",
				"showPivotChartFilter",
				"",
				attrValueTrue,
			),
		)
	} else {
		w.RemoveAttribute("showPivotChartFilter", "")
	}
}

// AllowRefreshQuery returns whether refresh queries are allowed.
func (w *WorkbookPr) AllowRefreshQuery() bool {
	attr, found := w.GetAttribute(
		"allowRefreshQuery",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAllowRefreshQuery sets whether refresh queries are allowed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetAllowRefreshQuery(
	value bool,
) {
	if value {
		w.SetAttribute(
			openxml.NewAttribute(
				"",
				"allowRefreshQuery",
				"",
				attrValueTrue,
			),
		)
	} else {
		w.RemoveAttribute("allowRefreshQuery", "")
	}
}

// PublishItems returns whether items should be published.
func (w *WorkbookPr) PublishItems() bool {
	attr, found := w.GetAttribute(
		"publishItems",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPublishItems sets whether items should be published.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetPublishItems(value bool) {
	if value {
		w.SetAttribute(
			openxml.NewAttribute(
				"",
				"publishItems",
				"",
				attrValueTrue,
			),
		)
	} else {
		w.RemoveAttribute("publishItems", "")
	}
}

// CheckCompatibility returns whether to check compatibility on save.
func (w *WorkbookPr) CheckCompatibility() bool {
	attr, found := w.GetAttribute(
		"checkCompatibility",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCheckCompatibility sets whether to check compatibility on save.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetCheckCompatibility(
	value bool,
) {
	if value {
		w.SetAttribute(
			openxml.NewAttribute(
				"",
				"checkCompatibility",
				"",
				attrValueTrue,
			),
		)
	} else {
		w.RemoveAttribute("checkCompatibility", "")
	}
}

// AutoCompressPictures returns whether pictures are automatically compressed.
func (w *WorkbookPr) AutoCompressPictures() bool {
	attr, found := w.GetAttribute(
		"autoCompressPictures",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetAutoCompressPictures sets whether pictures are automatically compressed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetAutoCompressPictures(
	value bool,
) {
	if value {
		w.RemoveAttribute(
			"autoCompressPictures",
			"",
		)
	} else {
		w.SetAttribute(openxml.NewAttribute("", "autoCompressPictures", "", attrValueFalse))
	}
}

// RefreshAllConnections returns whether all connections should be
// refreshed on open.
func (w *WorkbookPr) RefreshAllConnections() bool {
	attr, found := w.GetAttribute(
		"refreshAllConnections",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetRefreshAllConnections sets whether all connections should be refreshed on open.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (w *WorkbookPr) SetRefreshAllConnections(
	value bool,
) {
	if value {
		w.SetAttribute(
			openxml.NewAttribute(
				"",
				"refreshAllConnections",
				"",
				attrValueTrue,
			),
		)
	} else {
		w.RemoveAttribute("refreshAllConnections", "")
	}
}

// Clone creates a deep copy of this WorkbookPr element.
func (w *WorkbookPr) Clone() openxml.Element {
	cloned := w.CompositeElementBase.Clone()

	return &WorkbookPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this WorkbookPr element.
func (w *WorkbookPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := w.CompositeElementBase.CloneNode(
		deep,
	)

	return &WorkbookPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
