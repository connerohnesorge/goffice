package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// TargetScreenSizeValue represents target screen size values for web settings.
type TargetScreenSizeValue string

const (
	// TargetScreenSize544x376 represents 544x376 screen size.
	TargetScreenSize544x376 TargetScreenSizeValue = "544x376"
	// TargetScreenSize640x480 represents 640x480 screen size.
	TargetScreenSize640x480 TargetScreenSizeValue = "640x480"
	// TargetScreenSize720x512 represents 720x512 screen size.
	TargetScreenSize720x512 TargetScreenSizeValue = "720x512"
	// TargetScreenSize800x600 represents 800x600 screen size.
	TargetScreenSize800x600 TargetScreenSizeValue = "800x600"
	// TargetScreenSize1024x768 represents 1024x768 screen size.
	TargetScreenSize1024x768 TargetScreenSizeValue = "1024x768"
	// TargetScreenSize1152x882 represents 1152x882 screen size.
	TargetScreenSize1152x882 TargetScreenSizeValue = "1152x882"
	// TargetScreenSize1152x900 represents 1152x900 screen size.
	TargetScreenSize1152x900 TargetScreenSizeValue = "1152x900"
	// TargetScreenSize1280x1024 represents 1280x1024 screen size.
	TargetScreenSize1280x1024 TargetScreenSizeValue = "1280x1024"
	// TargetScreenSize1600x1200 represents 1600x1200 screen size.
	TargetScreenSize1600x1200 TargetScreenSizeValue = "1600x1200"
	// TargetScreenSize1800x1440 represents 1800x1440 screen size.
	TargetScreenSize1800x1440 TargetScreenSizeValue = "1800x1440"
	// TargetScreenSize1920x1200 represents 1920x1200 screen size.
	TargetScreenSize1920x1200 TargetScreenSizeValue = "1920x1200"
)

// WebSettings represents the root element for web settings part (w:webSettings).
type WebSettings struct {
	*openxml.CompositeElementBase
}

// NewWebSettings creates a new WebSettings element.
func NewWebSettings() *WebSettings {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"webSettings",
		PrefixW,
	)
	return &WebSettings{
		CompositeElementBase: elem,
	}
}

// OptimizeForBrowser returns whether to optimize for browser.
func (ws *WebSettings) OptimizeForBrowser() bool {
	return ws.hasOnOffElement(
		"optimizeForBrowser",
	)
}

// SetOptimizeForBrowser sets whether to optimize for browser.
func (ws *WebSettings) SetOptimizeForBrowser(
	b bool,
) {
	ws.setOnOffElement("optimizeForBrowser", b)
}

// AllowPNG returns whether PNG images are allowed.
func (ws *WebSettings) AllowPNG() bool {
	return ws.hasOnOffElement("allowPNG")
}

// SetAllowPNG sets whether PNG images are allowed.
func (ws *WebSettings) SetAllowPNG(b bool) {
	ws.setOnOffElement("allowPNG", b)
}

// TargetScreenSize returns the target screen size.
func (ws *WebSettings) TargetScreenSize() TargetScreenSizeValue {
	elem := ws.GetElement(
		"targetScreenSz",
		NamespaceWML,
	)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}
	return TargetScreenSizeValue(attr.Value())
}

// SetTargetScreenSize sets the target screen size.
func (ws *WebSettings) SetTargetScreenSize(
	size TargetScreenSizeValue,
) {
	if size == "" {
		ws.removeElement("targetScreenSz")
		return
	}
	elem := ws.getOrCreateElement(
		"targetScreenSz",
	)
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(size),
		),
	)
}

// Encoding returns the web page encoding.
func (ws *WebSettings) Encoding() string {
	elem := ws.GetElement(
		"encoding",
		NamespaceWML,
	)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetEncoding sets the web page encoding.
func (ws *WebSettings) SetEncoding(enc string) {
	if enc == "" {
		ws.removeElement("encoding")
		return
	}
	elem := ws.getOrCreateElement("encoding")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			enc,
		),
	)
}

// DoNotUseLongFileNames returns whether to use short file names.
func (ws *WebSettings) DoNotUseLongFileNames() bool {
	return ws.hasOnOffElement(
		"doNotUseLongFileNames",
	)
}

// SetDoNotUseLongFileNames sets whether to use short file names.
func (ws *WebSettings) SetDoNotUseLongFileNames(
	b bool,
) {
	ws.setOnOffElement("doNotUseLongFileNames", b)
}

// RelyOnVML returns whether to rely on VML for graphics.
func (ws *WebSettings) RelyOnVML() bool {
	return ws.hasOnOffElement("relyOnVML")
}

// SetRelyOnVML sets whether to rely on VML for graphics.
func (ws *WebSettings) SetRelyOnVML(b bool) {
	ws.setOnOffElement("relyOnVML", b)
}

// DoNotRelyOnCSS returns whether to avoid relying on CSS.
func (ws *WebSettings) DoNotRelyOnCSS() bool {
	return ws.hasOnOffElement("doNotRelyOnCSS")
}

// SetDoNotRelyOnCSS sets whether to avoid relying on CSS.
func (ws *WebSettings) SetDoNotRelyOnCSS(b bool) {
	ws.setOnOffElement("doNotRelyOnCSS", b)
}

// DoNotSaveAsSingleFile returns whether to avoid saving as single file.
func (ws *WebSettings) DoNotSaveAsSingleFile() bool {
	return ws.hasOnOffElement(
		"doNotSaveAsSingleFile",
	)
}

// SetDoNotSaveAsSingleFile sets whether to avoid saving as single file.
func (ws *WebSettings) SetDoNotSaveAsSingleFile(
	b bool,
) {
	ws.setOnOffElement("doNotSaveAsSingleFile", b)
}

// DoNotOrganizeInFolder returns whether to not organize in folder.
func (ws *WebSettings) DoNotOrganizeInFolder() bool {
	return ws.hasOnOffElement(
		"doNotOrganizeInFolder",
	)
}

// SetDoNotOrganizeInFolder sets whether to not organize in folder.
func (ws *WebSettings) SetDoNotOrganizeInFolder(
	b bool,
) {
	ws.setOnOffElement("doNotOrganizeInFolder", b)
}

// PixelsPerInch returns the pixels per inch setting.
func (ws *WebSettings) PixelsPerInch() int {
	elem := ws.GetElement(
		"pixelsPerInch",
		NamespaceWML,
	)
	if elem == nil {
		return 96 // Default
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return 96
	}
	// Parse as integer
	val := 0
	for _, c := range attr.Value() {
		if c >= '0' && c <= '9' {
			val = val*10 + int(c-'0')
		} else {
			break
		}
	}
	if val == 0 {
		return 96
	}
	return val
}

// SetPixelsPerInch sets the pixels per inch setting.
func (ws *WebSettings) SetPixelsPerInch(ppi int) {
	if ppi <= 0 {
		ws.removeElement("pixelsPerInch")
		return
	}
	elem := ws.getOrCreateElement("pixelsPerInch")
	// Convert int to string manually
	s := ""
	if ppi == 0 {
		s = "0"
	} else {
		digits := make([]byte, 0, 10)
		for ppi > 0 {
			digits = append(digits, byte('0'+ppi%10))
			ppi /= 10
		}
		// Reverse
		for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
			digits[i], digits[j] = digits[j], digits[i]
		}
		s = string(digits)
	}
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			s,
		),
	)
}

// Divs returns the Divs element containing div definitions, or nil if not present.
func (ws *WebSettings) Divs() *Divs {
	elem := ws.GetElement("divs", NamespaceWML)
	if elem == nil {
		return nil
	}
	if d, ok := elem.(*Divs); ok {
		return d
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Divs{CompositeElementBase: comp}
	}
	return nil
}

// GetOrCreateDivs returns the Divs element, creating if needed.
func (ws *WebSettings) GetOrCreateDivs() *Divs {
	d := ws.Divs()
	if d != nil {
		return d
	}
	d = NewDivs()
	ws.AppendChild(d)
	return d
}

// Helper methods

func (ws *WebSettings) hasOnOffElement(
	name string,
) bool {
	elem := ws.GetElement(name, NamespaceWML)
	if elem == nil {
		return false
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if found {
		val := attr.Value()
		return val != "false" && val != "0" &&
			val != "off"
	}
	return true
}

func (ws *WebSettings) setOnOffElement(
	name string,
	value bool,
) {
	if value {
		ws.getOrCreateElement(name)
	} else {
		ws.removeElement(name)
	}
}

func (ws *WebSettings) getOrCreateElement(
	name string,
) openxml.Element {
	elem := ws.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	ws.AppendChild(newElem)
	return newElem
}

func (ws *WebSettings) removeElement(
	name string,
) {
	elem := ws.GetElement(name, NamespaceWML)
	if elem != nil {
		ws.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this WebSettings element.
func (ws *WebSettings) Clone() openxml.Element {
	return &WebSettings{
		CompositeElementBase: ws.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this WebSettings element.
func (ws *WebSettings) CloneNode(
	deep bool,
) openxml.Element {
	return &WebSettings{
		CompositeElementBase: ws.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// Divs represents the divs container element (w:divs).
type Divs struct {
	*openxml.CompositeElementBase
}

// NewDivs creates a new Divs element.
func NewDivs() *Divs {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"divs",
		PrefixW,
	)
	return &Divs{CompositeElementBase: elem}
}

// Clone creates a deep copy of this Divs element.
func (d *Divs) Clone() openxml.Element {
	return &Divs{
		CompositeElementBase: d.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Divs element.
func (d *Divs) CloneNode(
	deep bool,
) openxml.Element {
	return &Divs{
		CompositeElementBase: d.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}
