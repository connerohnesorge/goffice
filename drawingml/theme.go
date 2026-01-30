package drawingml

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Theme represents the root theme element (a:theme).
// This element defines the color, font, and format schemes used in
// Office documents.
type Theme struct {
	*openxml.PartRootElementBase
}

// NewTheme creates a new Theme element with the given name.
func NewTheme(name string) *Theme {
	elem := openxml.NewPartRootElement(
		NamespaceMain,
		"theme",
		PrefixMain,
	)
	t := &Theme{PartRootElementBase: elem}

	// Set the name attribute if provided
	if name != "" {
		t.SetAttribute(
			openxml.NewAttribute(
				"",
				"name",
				"",
				name,
			),
		)
	}

	return t
}

// Name returns the name attribute of the theme.
func (t *Theme) Name() string {
	attr, found := t.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the name attribute of the theme.
func (t *Theme) SetName(name string) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// ThemeElements returns the theme elements child element (a:themeElements).
func (t *Theme) ThemeElements() *ThemeElements {
	elem := t.GetElement(
		"themeElements",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if te, ok := elem.(*ThemeElements); ok {
		return te
	}

	return nil
}

// GetOrCreateThemeElements returns or creates the theme elements.
func (t *Theme) GetOrCreateThemeElements() *ThemeElements {
	te := t.ThemeElements()
	if te != nil {
		return te
	}
	te = NewThemeElements()
	t.AppendChild(te)

	return te
}

// ObjectDefaults returns the object defaults child element (a:objectDefaults).
func (t *Theme) ObjectDefaults() openxml.Element {
	return t.GetElement(
		"objectDefaults",
		NamespaceMain,
	)
}

// ExtraClrSchemeLst returns the extra color scheme list child element
// (a:extraClrSchemeLst).
func (t *Theme) ExtraClrSchemeLst() *ExtraColorSchemeList {
	elem := t.GetElement(
		"extraClrSchemeLst",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if ecsl, ok := elem.(*ExtraColorSchemeList); ok {
		return ecsl
	}

	return nil
}

// GetOrCreateExtraClrSchemeLst returns or creates the extra color scheme list.
func (t *Theme) GetOrCreateExtraClrSchemeLst() *ExtraColorSchemeList {
	ecsl := t.ExtraClrSchemeLst()
	if ecsl != nil {
		return ecsl
	}
	ecsl = NewExtraColorSchemeList()
	t.AppendChild(ecsl)

	return ecsl
}
