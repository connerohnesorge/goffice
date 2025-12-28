package elements

//revive:disable:file-length-limit many defined name properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// DefinedNames represents the defined names container element (x:definedNames).
type DefinedNames struct {
	*openxml.CompositeElementBase
}

// NewDefinedNames creates a new DefinedNames element.
func NewDefinedNames() *DefinedNames {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"definedNames",
		PrefixDefault,
	)

	return &DefinedNames{
		CompositeElementBase: elem,
	}
}

// DefinedNames returns an iterator over all DefinedName elements.
func (dn *DefinedNames) DefinedNames() iter.Seq[*DefinedName] {
	return func(yield func(*DefinedName) bool) {
		for child := range dn.Children() {
			if child.LocalName() != "definedName" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var name *DefinedName
			switch v := child.(type) {
			case *DefinedName:
				name = v
			case *openxml.LeafElementBase:
				name = &DefinedName{LeafElementBase: v}
			}
			if name != nil && !yield(name) {
				return
			}
		}
	}
}

// Count returns the number of defined names.
func (dn *DefinedNames) Count() int {
	count := 0
	for range dn.DefinedNames() {
		count++
	}

	return count
}

// GetByName returns the defined name with the given name, or nil if not found.
// If localSheetId is -1, it looks for a workbook-level name.
// Otherwise, it looks for a sheet-level name with that local sheet ID.
//
//nolint:revive // early-return: nested conditionals for clarity
func (dn *DefinedNames) GetByName(
	name string,
	localSheetId int,
) *DefinedName {
	for defName := range dn.DefinedNames() {
		if defName.Name() == name {
			if localSheetId == -1 {
				// Looking for workbook-level name
				if !defName.HasLocalSheetId() {
					return defName
				}
			} else {
				// Looking for sheet-level name
				if defName.LocalSheetId() == localSheetId {
					return defName
				}
			}
		}
	}

	return nil
}

// GetWorkbookLevelName returns a workbook-level defined name.
func (dn *DefinedNames) GetWorkbookLevelName(
	name string,
) *DefinedName {
	return dn.GetByName(name, -1)
}

// GetSheetLevelName returns a sheet-level defined name.
func (dn *DefinedNames) GetSheetLevelName(
	name string, localSheetId int,
) *DefinedName {
	return dn.GetByName(name, localSheetId)
}

// AddDefinedName adds a new defined name with the given name and formula.
func (dn *DefinedNames) AddDefinedName(
	name, formula string,
) *DefinedName {
	defName := NewDefinedName()
	defName.SetName(name)
	defName.SetFormula(formula)
	dn.AppendChild(defName)

	return defName
}

// AddSheetLevelName adds a sheet-level defined name.
func (dn *DefinedNames) AddSheetLevelName(
	name string, localSheetId int, formula string,
) *DefinedName {
	defName := NewDefinedName()
	defName.SetName(name)
	defName.SetLocalSheetId(localSheetId)
	defName.SetFormula(formula)
	dn.AppendChild(defName)

	return defName
}

// RemoveDefinedName removes a defined name from the collection.
func (dn *DefinedNames) RemoveDefinedName(
	defName *DefinedName,
) bool {
	return dn.RemoveChild(defName)
}

// Clone creates a deep copy of this DefinedNames element.
func (dn *DefinedNames) Clone() openxml.Element {
	cloned := dn.CompositeElementBase.Clone()

	return &DefinedNames{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DefinedNames element.
func (dn *DefinedNames) CloneNode(
	deep bool,
) openxml.Element {
	cloned := dn.CompositeElementBase.CloneNode(
		deep,
	)

	return &DefinedNames{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// DefinedName represents a single defined name element (x:definedName).
// The text content of this element is the formula/range reference.
type DefinedName struct {
	*openxml.LeafElementBase
}

// NewDefinedName creates a new DefinedName element.
func NewDefinedName() *DefinedName {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"definedName",
		PrefixDefault,
	)

	return &DefinedName{LeafElementBase: elem}
}

// Name returns the name of the defined name.
func (d *DefinedName) Name() string {
	attr, found := d.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the name of the defined name.
func (d *DefinedName) SetName(name string) {
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Comment returns the comment associated with this defined name.
func (d *DefinedName) Comment() string {
	attr, found := d.GetAttribute("comment", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetComment sets the comment for this defined name.
func (d *DefinedName) SetComment(comment string) {
	if comment == "" {
		d.RemoveAttribute("comment", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"comment",
			"",
			comment,
		),
	)
}

// CustomMenu returns the custom menu text.
func (d *DefinedName) CustomMenu() string {
	attr, found := d.GetAttribute(
		"customMenu",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCustomMenu sets the custom menu text.
func (d *DefinedName) SetCustomMenu(text string) {
	if text == "" {
		d.RemoveAttribute("customMenu", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"customMenu",
			"",
			text,
		),
	)
}

// Description returns the description of the defined name.
func (d *DefinedName) Description() string {
	attr, found := d.GetAttribute(
		"description",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDescription sets the description.
func (d *DefinedName) SetDescription(
	desc string,
) {
	if desc == "" {
		d.RemoveAttribute("description", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"description",
			"",
			desc,
		),
	)
}

// Help returns the help text for the defined name.
func (d *DefinedName) Help() string {
	attr, found := d.GetAttribute("help", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetHelp sets the help text.
func (d *DefinedName) SetHelp(help string) {
	if help == "" {
		d.RemoveAttribute("help", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"help",
			"",
			help,
		),
	)
}

// StatusBar returns the status bar text.
func (d *DefinedName) StatusBar() string {
	attr, found := d.GetAttribute("statusBar", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetStatusBar sets the status bar text.
func (d *DefinedName) SetStatusBar(text string) {
	if text == "" {
		d.RemoveAttribute("statusBar", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"statusBar",
			"",
			text,
		),
	)
}

// LocalSheetId returns the local sheet ID that scopes this name.
// Returns -1 if not set (workbook-level name).
func (d *DefinedName) LocalSheetId() int {
	attr, found := d.GetAttribute(
		attrNameLocalSheetID,
		"",
	)
	if !found {
		return -1
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return -1
	}

	return val
}

// HasLocalSheetId returns true if this is a sheet-level name.
func (d *DefinedName) HasLocalSheetId() bool {
	_, found := d.GetAttribute(
		attrNameLocalSheetID,
		"",
	)

	return found
}

// SetLocalSheetId sets the local sheet ID to scope this name to a specific
// sheet. Use -1 to remove the scope and make it a workbook-level name.
func (d *DefinedName) SetLocalSheetId(id int) {
	if id < 0 {
		d.RemoveAttribute(
			attrNameLocalSheetID,
			"",
		)

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			attrNameLocalSheetID,
			"",
			strconv.Itoa(id),
		),
	)
}

// Hidden returns whether the defined name is hidden from the UI.
func (d *DefinedName) Hidden() bool {
	attr, found := d.GetAttribute("hidden", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHidden sets whether the defined name is hidden.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DefinedName) SetHidden(hidden bool) {
	if hidden {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"hidden",
				"",
				attrValueTrue,
			),
		)
	} else {
		d.RemoveAttribute("hidden", "")
	}
}

// Function returns whether this is a user-defined function.
func (d *DefinedName) Function() bool {
	attr, found := d.GetAttribute("function", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetFunction sets whether this is a user-defined function.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DefinedName) SetFunction(
	isFunction bool,
) {
	if isFunction {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"function",
				"",
				attrValueTrue,
			),
		)
	} else {
		d.RemoveAttribute("function", "")
	}
}

// VbProcedure returns whether this is a VBA procedure.
func (d *DefinedName) VbProcedure() bool {
	attr, found := d.GetAttribute(
		"vbProcedure",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetVbProcedure sets whether this is a VBA procedure.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DefinedName) SetVbProcedure(
	isVbProc bool,
) {
	if isVbProc {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"vbProcedure",
				"",
				attrValueTrue,
			),
		)
	} else {
		d.RemoveAttribute("vbProcedure", "")
	}
}

// Xlm returns whether this is an XLM (Excel 4.0) macro.
func (d *DefinedName) Xlm() bool {
	attr, found := d.GetAttribute("xlm", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetXlm sets whether this is an XLM macro.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DefinedName) SetXlm(isXlm bool) {
	if isXlm {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"xlm",
				"",
				attrValueTrue,
			),
		)
	} else {
		d.RemoveAttribute("xlm", "")
	}
}

// FunctionGroupId returns the function group ID.
func (d *DefinedName) FunctionGroupId() int {
	attr, found := d.GetAttribute(
		"functionGroupId",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetFunctionGroupId sets the function group ID.
func (d *DefinedName) SetFunctionGroupId(id int) {
	if id == 0 {
		d.RemoveAttribute("functionGroupId", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"functionGroupId",
			"",
			strconv.Itoa(id),
		),
	)
}

// ShortcutKey returns the keyboard shortcut for this name.
func (d *DefinedName) ShortcutKey() string {
	attr, found := d.GetAttribute(
		"shortcutKey",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetShortcutKey sets the keyboard shortcut.
func (d *DefinedName) SetShortcutKey(key string) {
	if key == "" {
		d.RemoveAttribute("shortcutKey", "")

		return
	}
	d.SetAttribute(
		openxml.NewAttribute(
			"",
			"shortcutKey",
			"",
			key,
		),
	)
}

// PublishToServer returns whether to publish to server.
func (d *DefinedName) PublishToServer() bool {
	attr, found := d.GetAttribute(
		"publishToServer",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPublishToServer sets whether to publish to server.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DefinedName) SetPublishToServer(
	publish bool,
) {
	if publish {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"publishToServer",
				"",
				attrValueTrue,
			),
		)
	} else {
		d.RemoveAttribute("publishToServer", "")
	}
}

// WorkbookParameter returns whether this is a workbook parameter.
func (d *DefinedName) WorkbookParameter() bool {
	attr, found := d.GetAttribute(
		"workbookParameter",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetWorkbookParameter sets whether this is a workbook parameter.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (d *DefinedName) SetWorkbookParameter(
	isParam bool,
) {
	if isParam {
		d.SetAttribute(
			openxml.NewAttribute(
				"",
				"workbookParameter",
				"",
				attrValueTrue,
			),
		)
	} else {
		d.RemoveAttribute("workbookParameter", "")
	}
}

// Formula returns the formula/range reference (text content of the element).
func (d *DefinedName) Formula() string {
	return d.InnerText()
}

// SetFormula sets the formula/range reference.
func (d *DefinedName) SetFormula(formula string) {
	d.SetInnerText(formula)
}

// Clone creates a deep copy of this DefinedName element.
func (d *DefinedName) Clone() openxml.Element {
	cloned := d.LeafElementBase.Clone()

	return &DefinedName{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this DefinedName element.
func (d *DefinedName) CloneNode(
	deep bool,
) openxml.Element {
	cloned := d.LeafElementBase.CloneNode(deep)

	return &DefinedName{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
