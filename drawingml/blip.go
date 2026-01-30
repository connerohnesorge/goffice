package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// BlipCompression represents the compression state of a BLIP.
type BlipCompression string

// Blip compression constants.
const (
	BlipCompressionEmail   BlipCompression = "email"
	BlipCompressionScreen  BlipCompression = "screen"
	BlipCompressionPrint   BlipCompression = "print"
	BlipCompressionHQPrint BlipCompression = "hqprint"
	BlipCompressionNone    BlipCompression = "none"
)

// Relationship namespace constants.
const (
	// relNamespace is the namespace for relationship attributes.
	relNamespace = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"
)

// Attribute name constants.
const (
	// attrThresh is the threshold attribute name.
	attrThresh = "thresh"
)

// Magic number constants.
const (
	// fullAlpha is the value representing 100% alpha (100000).
	fullAlpha = 100000
)

// Blip represents a Binary Large Image or Picture (a:blip).
type Blip struct {
	*openxml.CompositeElementBase
}

// NewBlip creates a new BLIP element.
func NewBlip() *Blip {
	return &Blip{
		CompositeElementBase: openxml.NewCompositeElement(
			NamespaceMain,
			"blip",
			PrefixMain,
		),
	}
}

// CompressionState returns the compression state.
func (b *Blip) CompressionState() BlipCompression {
	attr, found := b.GetAttribute("cstate", "")
	if !found {
		return BlipCompressionNone
	}

	return BlipCompression(attr.Value())
}

// SetCompressionState sets the compression state.
func (b *Blip) SetCompressionState(state BlipCompression) {
	if state == BlipCompressionNone || state == "" {
		b.RemoveAttribute("cstate", "")

		return
	}
	b.SetAttribute(openxml.NewAttribute("", "cstate", "", string(state)))
}

// Embed returns the embedded relationship ID.
func (b *Blip) Embed() string {
	attr, found := b.GetAttribute("embed", relNamespace)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetEmbed sets the embedded relationship ID.
func (b *Blip) SetEmbed(embedId string) {
	b.SetAttribute(openxml.NewAttribute(
		relNamespace,
		"embed",
		"r",
		embedId,
	))
}

// Link returns the linked relationship ID.
func (b *Blip) Link() string {
	attr, found := b.GetAttribute("link", relNamespace)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetLink sets the linked relationship ID.
func (b *Blip) SetLink(linkId string) {
	b.SetAttribute(openxml.NewAttribute(
		relNamespace,
		"link",
		"r",
		linkId,
	))
}

// Clone creates a deep copy of this Blip element.
func (b *Blip) Clone() openxml.Element {
	return &Blip{
		CompositeElementBase: b.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// LuminanceEffect represents a luminance effect (a:lum).
type LuminanceEffect struct {
	*openxml.LeafElementBase
}

// NewLuminanceEffect creates a new luminance effect.
func NewLuminanceEffect(bright, contrast int) *LuminanceEffect {
	lum := &LuminanceEffect{
		LeafElementBase: openxml.NewLeafElement(NamespaceMain, "lum", PrefixMain),
	}
	lum.SetBrightness(bright)
	lum.SetContrast(contrast)

	return lum
}

// Brightness returns the brightness adjustment (0-100000).
func (l *LuminanceEffect) Brightness() int {
	attr, found := l.GetAttribute("bright", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetBrightness sets the brightness adjustment.
func (l *LuminanceEffect) SetBrightness(bright int) {
	if bright == 0 {
		l.RemoveAttribute("bright", "")

		return
	}
	l.SetAttribute(openxml.NewAttribute("", "bright", "", strconv.Itoa(bright)))
}

// Contrast returns the contrast adjustment (0-100000).
func (l *LuminanceEffect) Contrast() int {
	attr, found := l.GetAttribute("contrast", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetContrast sets the contrast adjustment.
func (l *LuminanceEffect) SetContrast(contrast int) {
	if contrast == 0 {
		l.RemoveAttribute("contrast", "")

		return
	}
	l.SetAttribute(openxml.NewAttribute("", "contrast", "", strconv.Itoa(contrast)))
}

// AddLuminanceEffect adds a luminance effect to the Blip.
func (b *Blip) AddLuminanceEffect(bright, contrast int) {
	lum := NewLuminanceEffect(bright, contrast)
	b.AppendChild(lum)
}

// BiLevelEffect represents a bi-level (threshold) effect (a:biLevel).
// Makes image black and white based on threshold.
type BiLevelEffect struct {
	*openxml.LeafElementBase
}

// NewBiLevelEffect creates a new bi-level effect.
func NewBiLevelEffect(thresh int) *BiLevelEffect {
	bl := &BiLevelEffect{
		LeafElementBase: openxml.NewLeafElement(NamespaceMain, "biLevel", PrefixMain),
	}
	bl.SetThreshold(thresh)

	return bl
}

// Threshold returns the threshold value.
func (b *BiLevelEffect) Threshold() int {
	attr, found := b.GetAttribute(attrThresh, "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetThreshold sets the threshold value.
func (b *BiLevelEffect) SetThreshold(thresh int) {
	b.SetAttribute(openxml.NewAttribute("", attrThresh, "", strconv.Itoa(thresh)))
}

// AddBiLevelEffect adds a bi-level effect to the Blip.
func (b *Blip) AddBiLevelEffect(thresh int) {
	bl := NewBiLevelEffect(thresh)
	b.AppendChild(bl)
}

// HSLEffect represents a Hue/Saturation/Luminance effect (a:hsl).
type HSLEffect struct {
	*openxml.LeafElementBase
}

// NewHSLEffect creates a new HSL effect.
func NewHSLEffect(hue, sat, lum int) *HSLEffect {
	h := &HSLEffect{
		LeafElementBase: openxml.NewLeafElement(NamespaceMain, "hsl", PrefixMain),
	}
	h.SetHue(hue)
	h.SetSaturation(sat)
	h.SetLuminance(lum)

	return h
}

// Hue returns the hue adjustment (0-36000000).
func (h *HSLEffect) Hue() int {
	attr, found := h.GetAttribute("hue", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetHue sets the hue adjustment.
func (h *HSLEffect) SetHue(hue int) {
	h.SetAttribute(openxml.NewAttribute("", "hue", "", strconv.Itoa(hue)))
}

// Saturation returns the saturation adjustment (0-100000).
func (h *HSLEffect) Saturation() int {
	attr, found := h.GetAttribute("sat", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetSaturation sets the saturation adjustment.
func (h *HSLEffect) SetSaturation(sat int) {
	h.SetAttribute(openxml.NewAttribute("", "sat", "", strconv.Itoa(sat)))
}

// Luminance returns the luminance adjustment (0-100000).
func (h *HSLEffect) Luminance() int {
	attr, found := h.GetAttribute("lum", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetLuminance sets the luminance adjustment.
func (h *HSLEffect) SetLuminance(lum int) {
	h.SetAttribute(openxml.NewAttribute("", "lum", "", strconv.Itoa(lum)))
}

// AddHSLEffect adds an HSL effect to the Blip.
func (b *Blip) AddHSLEffect(hue, sat, lum int) {
	hsl := NewHSLEffect(hue, sat, lum)
	b.AppendChild(hsl)
}

// GrayscaleEffect represents a grayscale effect (a:grayscl).
type GrayscaleEffect struct {
	*openxml.LeafElementBase
}

// NewGrayscaleEffect creates a new grayscale effect.
func NewGrayscaleEffect() *GrayscaleEffect {
	return &GrayscaleEffect{
		LeafElementBase: openxml.NewLeafElement(NamespaceMain, "grayscl", PrefixMain),
	}
}

// AddGrayscaleEffect adds a grayscale effect to the Blip.
func (b *Blip) AddGrayscaleEffect() {
	gs := NewGrayscaleEffect()
	b.AppendChild(gs)
}

// AlphaBiLevelEffect represents an alpha bi-level effect (a:alphaBiLevel).
type AlphaBiLevelEffect struct {
	*openxml.LeafElementBase
}

// NewAlphaBiLevelEffect creates a new alpha bi-level effect.
func NewAlphaBiLevelEffect(thresh int) *AlphaBiLevelEffect {
	a := &AlphaBiLevelEffect{
		LeafElementBase: openxml.NewLeafElement(NamespaceMain, "alphaBiLevel", PrefixMain),
	}
	a.SetThreshold(thresh)

	return a
}

// Threshold returns the threshold value.
func (a *AlphaBiLevelEffect) Threshold() int {
	attr, found := a.GetAttribute(attrThresh, "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetThreshold sets the threshold value.
func (a *AlphaBiLevelEffect) SetThreshold(thresh int) {
	a.SetAttribute(openxml.NewAttribute("", attrThresh, "", strconv.Itoa(thresh)))
}

// AddAlphaBiLevelEffect adds an alpha bi-level effect to the Blip.
func (b *Blip) AddAlphaBiLevelEffect(thresh int) {
	a := NewAlphaBiLevelEffect(thresh)
	b.AppendChild(a)
}

// AlphaModulationEffect represents an alpha modulation effect (a:alphaMod).
type AlphaModulationEffect struct {
	*openxml.CompositeElementBase
}

// NewAlphaModulationEffect creates a new alpha modulation effect.
func NewAlphaModulationEffect() *AlphaModulationEffect {
	return &AlphaModulationEffect{
		CompositeElementBase: openxml.NewCompositeElement(NamespaceMain, "alphaMod", PrefixMain),
	}
}

// AddAlphaModulationEffect adds an alpha modulation effect to the Blip.
func (b *Blip) AddAlphaModulationEffect(container *openxml.CompositeElementBase) {
	// alphaMod usually contains an effect container (a:cont) which defines the alpha mask?
	// The spec says alphaMod "Modulates the alpha of the input with the alpha of the mask".
	// It has a "cont" child.
	// For simplicity, we create the element. User must populate children if needed.
	am := NewAlphaModulationEffect()
	if container != nil {
		am.AppendChild(container)
	}
	b.AppendChild(am)
}

// AlphaFixedEffect represents a fixed alpha effect (a:alphaModFix).
type AlphaFixedEffect struct {
	*openxml.LeafElementBase
}

// NewAlphaFixedEffect creates a new fixed alpha effect.
func NewAlphaFixedEffect(amt int) *AlphaFixedEffect {
	a := &AlphaFixedEffect{
		LeafElementBase: openxml.NewLeafElement(NamespaceMain, "alphaModFix", PrefixMain),
	}
	a.SetAmount(amt)

	return a
}

// Amount returns the alpha amount (0-100000).
func (a *AlphaFixedEffect) Amount() int {
	attr, found := a.GetAttribute("amt", "")
	if !found {
		return fullAlpha
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetAmount sets the alpha amount.
func (a *AlphaFixedEffect) SetAmount(amt int) {
	a.SetAttribute(openxml.NewAttribute("", "amt", "", strconv.Itoa(amt)))
}

// AddAlphaFixedEffect adds a fixed alpha effect to the Blip.
func (b *Blip) AddAlphaFixedEffect(amt int) {
	a := NewAlphaFixedEffect(amt)
	b.AppendChild(a)
}
