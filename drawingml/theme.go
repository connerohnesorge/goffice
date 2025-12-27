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
func (t *Theme) ThemeElements() openxml.Element {
	return t.GetElement(
		"themeElements",
		NamespaceMain,
	)
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
func (t *Theme) ExtraClrSchemeLst() openxml.Element {
	return t.GetElement(
		"extraClrSchemeLst",
		NamespaceMain,
	)
}
