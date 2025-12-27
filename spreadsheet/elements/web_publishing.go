package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// TargetScreenSize represents the target screen size for web publishing.
type TargetScreenSize string

const (
	// TargetScreenSize544x376 is 544x376 resolution.
	TargetScreenSize544x376 TargetScreenSize = "544x376"
	// TargetScreenSize640x480 is 640x480 resolution.
	TargetScreenSize640x480 TargetScreenSize = "640x480"
	// TargetScreenSize720x512 is 720x512 resolution.
	TargetScreenSize720x512 TargetScreenSize = "720x512"
	// TargetScreenSize800x600 is 800x600 resolution.
	TargetScreenSize800x600 TargetScreenSize = "800x600"
	// TargetScreenSize1024x768 is 1024x768 resolution.
	TargetScreenSize1024x768 TargetScreenSize = "1024x768"
	// TargetScreenSize1152x882 is 1152x882 resolution.
	TargetScreenSize1152x882 TargetScreenSize = "1152x882"
	// TargetScreenSize1152x900 is 1152x900 resolution.
	TargetScreenSize1152x900 TargetScreenSize = "1152x900"
	// TargetScreenSize1280x1024 is 1280x1024 resolution.
	TargetScreenSize1280x1024 TargetScreenSize = "1280x1024"
	// TargetScreenSize1600x1200 is 1600x1200 resolution.
	TargetScreenSize1600x1200 TargetScreenSize = "1600x1200"
	// TargetScreenSize1800x1440 is 1800x1440 resolution.
	TargetScreenSize1800x1440 TargetScreenSize = "1800x1440"
	// TargetScreenSize1920x1200 is 1920x1200 resolution.
	TargetScreenSize1920x1200 TargetScreenSize = "1920x1200"
)

// WebPublishing represents the web publishing element (x:webPublishing).
type WebPublishing struct {
	*openxml.CompositeElementBase
}

// NewWebPublishing creates a new WebPublishing element.
func NewWebPublishing() *WebPublishing {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"webPublishing",
		PrefixDefault,
	)

	return &WebPublishing{
		CompositeElementBase: elem,
	}
}

// Css returns whether CSS is used for web publishing.
func (wp *WebPublishing) Css() bool {
	attr, found := wp.GetAttribute("css", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetCss sets whether CSS is used for web publishing.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wp *WebPublishing) SetCss(value bool) {
	if value {
		wp.RemoveAttribute("css", "")
	} else {
		wp.SetAttribute(openxml.NewAttribute("", "css", "", attrValueFalse))
	}
}

// Thicket returns whether files are organized in a thicket.
func (wp *WebPublishing) Thicket() bool {
	attr, found := wp.GetAttribute("thicket", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetThicket sets whether files are organized in a thicket.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wp *WebPublishing) SetThicket(value bool) {
	if value {
		wp.RemoveAttribute("thicket", "")
	} else {
		wp.SetAttribute(openxml.NewAttribute("", "thicket", "", attrValueFalse))
	}
}

// LongFileNames returns whether long file names are used.
func (wp *WebPublishing) LongFileNames() bool {
	attr, found := wp.GetAttribute(
		"longFileNames",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetLongFileNames sets whether long file names are used.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wp *WebPublishing) SetLongFileNames(
	value bool,
) {
	if value {
		wp.RemoveAttribute("longFileNames", "")
	} else {
		wp.SetAttribute(openxml.NewAttribute("", "longFileNames", "", attrValueFalse))
	}
}

// Vml returns whether VML is used for graphics.
func (wp *WebPublishing) Vml() bool {
	attr, found := wp.GetAttribute("vml", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetVml sets whether VML is used for graphics.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wp *WebPublishing) SetVml(value bool) {
	if value {
		wp.SetAttribute(
			openxml.NewAttribute(
				"",
				"vml",
				"",
				attrValueTrue,
			),
		)
	} else {
		wp.RemoveAttribute("vml", "")
	}
}

// AllowPng returns whether PNG images are allowed.
func (wp *WebPublishing) AllowPng() bool {
	attr, found := wp.GetAttribute("allowPng", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAllowPng sets whether PNG images are allowed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wp *WebPublishing) SetAllowPng(value bool) {
	if value {
		wp.SetAttribute(
			openxml.NewAttribute(
				"",
				"allowPng",
				"",
				attrValueTrue,
			),
		)
	} else {
		wp.RemoveAttribute("allowPng", "")
	}
}

// TargetScreenSz returns the target screen size.
func (wp *WebPublishing) TargetScreenSz() TargetScreenSize {
	attr, found := wp.GetAttribute(
		"targetScreenSz",
		"",
	)
	if !found {
		return TargetScreenSize800x600 // Default value
	}

	return TargetScreenSize(attr.Value())
}

// SetTargetScreenSz sets the target screen size.
func (wp *WebPublishing) SetTargetScreenSz(
	size TargetScreenSize,
) {
	if size == "" ||
		size == TargetScreenSize800x600 {
		wp.RemoveAttribute("targetScreenSz", "")

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			"targetScreenSz",
			"",
			string(size),
		),
	)
}

// Dpi returns the DPI for graphics.
func (wp *WebPublishing) Dpi() int {
	attr, found := wp.GetAttribute("dpi", "")
	if !found {
		return defaultDPI // Default value
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetDpi sets the DPI for graphics.
func (wp *WebPublishing) SetDpi(dpi int) {
	if dpi == defaultDPI {
		wp.RemoveAttribute("dpi", "")

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			"dpi",
			"",
			strconv.Itoa(dpi),
		),
	)
}

// CodePage returns the code page for web publishing.
func (wp *WebPublishing) CodePage() int {
	attr, found := wp.GetAttribute("codePage", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCodePage sets the code page for web publishing.
func (wp *WebPublishing) SetCodePage(
	codePage int,
) {
	if codePage == 0 {
		wp.RemoveAttribute("codePage", "")

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			"codePage",
			"",
			strconv.Itoa(codePage),
		),
	)
}

// CharacterSet returns the character set name.
func (wp *WebPublishing) CharacterSet() string {
	attr, found := wp.GetAttribute(
		"characterSet",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCharacterSet sets the character set name.
func (wp *WebPublishing) SetCharacterSet(
	charset string,
) {
	if charset == "" {
		wp.RemoveAttribute("characterSet", "")

		return
	}
	wp.SetAttribute(
		openxml.NewAttribute(
			"",
			"characterSet",
			"",
			charset,
		),
	)
}

// Clone creates a deep copy of this WebPublishing element.
func (wp *WebPublishing) Clone() openxml.Element {
	cloned := wp.CompositeElementBase.Clone()

	return &WebPublishing{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this WebPublishing element.
func (wp *WebPublishing) CloneNode(
	deep bool,
) openxml.Element {
	cloned := wp.CompositeElementBase.CloneNode(
		deep,
	)

	return &WebPublishing{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
