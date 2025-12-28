// Package drawingml provides shared DrawingML types for shapes, images,
// and effects.
//
//nolint:revive // This file contains many public types for OOXML effect elements.
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PresetShadowValue represents preset shadow configuration values.
// These are predefined shadow configurations in DrawingML.
type PresetShadowValue string

// Preset shadow values defined in DrawingML (shdw1 through shdw20).
const (
	// PresetShadowTopLeftDropShadow is a drop shadow from top-left.
	PresetShadowTopLeftDropShadow PresetShadowValue = "shdw1"
	// PresetShadowTopRightDropShadow is a drop shadow from top-right.
	PresetShadowTopRightDropShadow PresetShadowValue = "shdw2"
	// PresetShadowBackLeftPerspective is a perspective shadow from back-left.
	PresetShadowBackLeftPerspective PresetShadowValue = "shdw3"
	// PresetShadowBackRightPerspective is a perspective shadow from back-right.
	PresetShadowBackRightPerspective PresetShadowValue = "shdw4"
	// PresetShadowBottomLeftDropShadow is a drop shadow from bottom-left.
	PresetShadowBottomLeftDropShadow PresetShadowValue = "shdw5"
	// PresetShadowBottomRightDropShadow is a drop shadow from bottom-right.
	PresetShadowBottomRightDropShadow PresetShadowValue = "shdw6"
	// PresetShadowFrontLeftPerspective is a perspective shadow from front-left.
	PresetShadowFrontLeftPerspective PresetShadowValue = "shdw7"
	// PresetShadowFrontRightPerspective is a perspective shadow from front-right.
	PresetShadowFrontRightPerspective PresetShadowValue = "shdw8"
	// PresetShadowTopLeftSmallDropShadow is a small drop shadow from top-left.
	PresetShadowTopLeftSmallDropShadow PresetShadowValue = "shdw9"
	// PresetShadowTopLeftLargeDropShadow is a large drop shadow from top-left.
	PresetShadowTopLeftLargeDropShadow PresetShadowValue = "shdw10"
	// PresetShadowBackLeftLongPerspective is a long perspective shadow from back-left.
	PresetShadowBackLeftLongPerspective PresetShadowValue = "shdw11"
	// PresetShadowBackRightLongPerspective is a long perspective shadow from back-right.
	PresetShadowBackRightLongPerspective PresetShadowValue = "shdw12"
	// PresetShadowTopLeftDoubleDropShadow is a double drop shadow from top-left.
	PresetShadowTopLeftDoubleDropShadow PresetShadowValue = "shdw13"
	// PresetShadowBottomRightSmallDropShadow is a small drop shadow from bottom-right.
	PresetShadowBottomRightSmallDropShadow PresetShadowValue = "shdw14"
	// PresetShadowFrontLeftLongPerspective is a long perspective shadow from front-left.
	PresetShadowFrontLeftLongPerspective PresetShadowValue = "shdw15"
	// PresetShadowFrontRightLongPerspective is a long perspective shadow from front-right.
	PresetShadowFrontRightLongPerspective PresetShadowValue = "shdw16"
	// PresetShadow3DOuterBox is a 3D outer box shadow.
	PresetShadow3DOuterBox PresetShadowValue = "shdw17"
	// PresetShadow3DInnerBox is a 3D inner box shadow.
	PresetShadow3DInnerBox PresetShadowValue = "shdw18"
	// PresetShadowBackCenterPerspective is a perspective shadow from back-center.
	PresetShadowBackCenterPerspective PresetShadowValue = "shdw19"
	// PresetShadowFrontBottom is a front bottom shadow.
	PresetShadowFrontBottom PresetShadowValue = "shdw20"
)

// BlendModeValue represents the blend mode for fill overlays and effects.
type BlendModeValue string

// Blend mode values for effects.
const (
	// BlendModeOverlay overlays the fill on top of effects.
	BlendModeOverlay BlendModeValue = "over"
	// BlendModeMultiply multiplies the fill with effects.
	BlendModeMultiply BlendModeValue = "mult"
	// BlendModeScreen uses screen blending mode.
	BlendModeScreen BlendModeValue = "screen"
	// BlendModeDarken uses darken blending mode.
	BlendModeDarken BlendModeValue = "darken"
	// BlendModeLighten uses lighten blending mode.
	BlendModeLighten BlendModeValue = "lighten"
)

// EffectList represents an effect list container (a:effectLst).
// Effects in the list are applied in order: blur, fillOverlay, glow, innerShdw,
// outerShdw, prstShdw, reflection, softEdge.
type EffectList struct {
	*openxml.CompositeElementBase
}

// NewEffectList creates a new effect list container.
func NewEffectList() *EffectList {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"effectLst",
		PrefixMain,
	)

	return &EffectList{CompositeElementBase: elem}
}

// Blur returns the blur effect, or nil if not set.
func (e *EffectList) Blur() *Blur {
	elem := e.GetElement("blur", NamespaceMain)
	if elem == nil {
		return nil
	}
	if blur, ok := elem.(*Blur); ok {
		return blur
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Blur{LeafElementBase: leaf}
	}

	return nil
}

// SetBlur sets the blur effect.
func (e *EffectList) SetBlur(blur *Blur) {
	if existing := e.GetElement("blur", NamespaceMain); existing != nil {
		e.RemoveChild(existing)
	}
	if blur != nil {
		e.PrependChild(blur)
	}
}

// Glow returns the glow effect, or nil if not set.
func (e *EffectList) Glow() *Glow {
	elem := e.GetElement("glow", NamespaceMain)
	if elem == nil {
		return nil
	}
	if glow, ok := elem.(*Glow); ok {
		return glow
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Glow{CompositeElementBase: comp}
	}

	return nil
}

// SetGlow sets the glow effect.
func (e *EffectList) SetGlow(glow *Glow) {
	if existing := e.GetElement("glow", NamespaceMain); existing != nil {
		e.RemoveChild(existing)
	}
	if glow != nil {
		e.AppendChild(glow)
	}
}

// InnerShadow returns the inner shadow effect, or nil if not set.
func (e *EffectList) InnerShadow() *InnerShadow {
	elem := e.GetElement(
		"innerShdw",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if inner, ok := elem.(*InnerShadow); ok {
		return inner
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &InnerShadow{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetInnerShadow sets the inner shadow effect.
func (e *EffectList) SetInnerShadow(
	shadow *InnerShadow,
) {
	if existing := e.GetElement("innerShdw", NamespaceMain); existing != nil {
		e.RemoveChild(existing)
	}
	if shadow != nil {
		e.AppendChild(shadow)
	}
}

// OuterShadow returns the outer shadow effect, or nil if not set.
func (e *EffectList) OuterShadow() *OuterShadow {
	elem := e.GetElement(
		"outerShdw",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if outer, ok := elem.(*OuterShadow); ok {
		return outer
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &OuterShadow{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetOuterShadow sets the outer shadow effect.
func (e *EffectList) SetOuterShadow(
	shadow *OuterShadow,
) {
	if existing := e.GetElement("outerShdw", NamespaceMain); existing != nil {
		e.RemoveChild(existing)
	}
	if shadow != nil {
		e.AppendChild(shadow)
	}
}

// PresetShadow returns the preset shadow effect, or nil if not set.
func (e *EffectList) PresetShadow() *PresetShadow {
	elem := e.GetElement(
		"prstShdw",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if preset, ok := elem.(*PresetShadow); ok {
		return preset
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PresetShadow{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetPresetShadow sets the preset shadow effect.
func (e *EffectList) SetPresetShadow(
	shadow *PresetShadow,
) {
	if existing := e.GetElement("prstShdw", NamespaceMain); existing != nil {
		e.RemoveChild(existing)
	}
	if shadow != nil {
		e.AppendChild(shadow)
	}
}

// Reflection returns the reflection effect, or nil if not set.
func (e *EffectList) Reflection() *Reflection {
	elem := e.GetElement(
		"reflection",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if refl, ok := elem.(*Reflection); ok {
		return refl
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Reflection{LeafElementBase: leaf}
	}

	return nil
}

// SetReflection sets the reflection effect.
func (e *EffectList) SetReflection(
	reflection *Reflection,
) {
	if existing := e.GetElement("reflection", NamespaceMain); existing != nil {
		e.RemoveChild(existing)
	}
	if reflection != nil {
		e.AppendChild(reflection)
	}
}

// SoftEdge returns the soft edge effect, or nil if not set.
func (e *EffectList) SoftEdge() *SoftEdge {
	elem := e.GetElement(
		"softEdge",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if soft, ok := elem.(*SoftEdge); ok {
		return soft
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &SoftEdge{LeafElementBase: leaf}
	}

	return nil
}

// SetSoftEdge sets the soft edge effect.
func (e *EffectList) SetSoftEdge(
	softEdge *SoftEdge,
) {
	if existing := e.GetElement("softEdge", NamespaceMain); existing != nil {
		e.RemoveChild(existing)
	}
	if softEdge != nil {
		e.AppendChild(softEdge)
	}
}

// FillOverlay returns the fill overlay effect, or nil if not set.
func (e *EffectList) FillOverlay() *FillOverlay {
	elem := e.GetElement(
		"fillOverlay",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if overlay, ok := elem.(*FillOverlay); ok {
		return overlay
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &FillOverlay{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetFillOverlay sets the fill overlay effect.
func (e *EffectList) SetFillOverlay(
	overlay *FillOverlay,
) {
	if existing := e.GetElement("fillOverlay", NamespaceMain); existing != nil {
		e.RemoveChild(existing)
	}
	if overlay != nil {
		e.AppendChild(overlay)
	}
}

// Clone creates a deep copy of this EffectList element.
func (e *EffectList) Clone() openxml.Element {
	return &EffectList{
		CompositeElementBase: e.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// EffectContainer represents an effect container for complex effect compositions (a:effectDag).
// This creates a directed acyclic graph (DAG) of effects that can be combined.
type EffectContainer struct {
	*openxml.CompositeElementBase
}

// NewEffectContainer creates a new effect container (DAG).
func NewEffectContainer() *EffectContainer {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"effectDag",
		PrefixMain,
	)

	return &EffectContainer{
		CompositeElementBase: elem,
	}
}

// Type returns the container type ("sib" or "tree").
func (c *EffectContainer) Type() string {
	attr, found := c.GetAttribute("type", "")
	if !found {
		return "sib"
	}

	return attr.Value()
}

// SetType sets the container type.
func (c *EffectContainer) SetType(
	containerType string,
) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			containerType,
		),
	)
}

// Name returns the container name.
func (c *EffectContainer) Name() string {
	attr, found := c.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the container name.
func (c *EffectContainer) SetName(name string) {
	if name == "" {
		c.RemoveAttribute("name", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Clone creates a deep copy of this EffectContainer element.
func (c *EffectContainer) Clone() openxml.Element {
	return &EffectContainer{
		CompositeElementBase: c.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// OuterShadow represents an outer (drop) shadow effect (a:outerShdw).
type OuterShadow struct {
	*openxml.CompositeElementBase
}

// NewOuterShadow creates a new outer shadow with default values.
func NewOuterShadow() *OuterShadow {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"outerShdw",
		PrefixMain,
	)

	return &OuterShadow{
		CompositeElementBase: elem,
	}
}

// NewDropShadow creates a typical drop shadow effect.
// blurRadius is in EMUs, distance is in EMUs, direction is in 60000ths of a degree
// (5400000 = 90 degrees = downward).
func NewDropShadow(
	blurRadius, distance EMU,
	direction int,
	colorHex string,
) *OuterShadow {
	shadow := NewOuterShadow()
	shadow.SetBlurRadius(blurRadius)
	shadow.SetDistance(distance)
	shadow.SetDirection(direction)
	shadow.SetRgbColor(colorHex)

	return shadow
}

// BlurRadius returns the blur radius in EMUs.
func (s *OuterShadow) BlurRadius() EMU {
	attr, found := s.GetAttribute("blurRad", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetBlurRadius sets the blur radius in EMUs.
func (s *OuterShadow) SetBlurRadius(radius EMU) {
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"blurRad",
			"",
			strconv.FormatInt(int64(radius), 10),
		),
	)
}

// Distance returns the shadow distance in EMUs.
func (s *OuterShadow) Distance() EMU {
	attr, found := s.GetAttribute("dist", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetDistance sets the shadow distance in EMUs.
func (s *OuterShadow) SetDistance(dist EMU) {
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"dist",
			"",
			strconv.FormatInt(int64(dist), 10),
		),
	)
}

// Direction returns the shadow direction in 60000ths of a degree.
// 0 = right, 5400000 = down, 10800000 = left, 16200000 = up.
func (s *OuterShadow) Direction() int {
	attr, found := s.GetAttribute("dir", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetDirection sets the shadow direction in 60000ths of a degree.
func (s *OuterShadow) SetDirection(dir int) {
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			strconv.Itoa(dir),
		),
	)
}

// HorizontalRatio returns the horizontal scaling factor (100000 = 100%).
func (s *OuterShadow) HorizontalRatio() int {
	attr, found := s.GetAttribute("sx", "")
	if !found {
		return 100000
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetHorizontalRatio sets the horizontal scaling factor.
func (s *OuterShadow) SetHorizontalRatio(
	ratio int,
) {
	if ratio == 100000 {
		s.RemoveAttribute("sx", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"sx",
			"",
			strconv.Itoa(ratio),
		),
	)
}

// VerticalRatio returns the vertical scaling factor (100000 = 100%).
func (s *OuterShadow) VerticalRatio() int {
	attr, found := s.GetAttribute("sy", "")
	if !found {
		return 100000
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetVerticalRatio sets the vertical scaling factor.
func (s *OuterShadow) SetVerticalRatio(
	ratio int,
) {
	if ratio == 100000 {
		s.RemoveAttribute("sy", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"sy",
			"",
			strconv.Itoa(ratio),
		),
	)
}

// HorizontalSkew returns the horizontal skew angle in 60000ths of a degree.
func (s *OuterShadow) HorizontalSkew() int {
	attr, found := s.GetAttribute("kx", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetHorizontalSkew sets the horizontal skew angle.
func (s *OuterShadow) SetHorizontalSkew(
	skew int,
) {
	if skew == 0 {
		s.RemoveAttribute("kx", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"kx",
			"",
			strconv.Itoa(skew),
		),
	)
}

// VerticalSkew returns the vertical skew angle in 60000ths of a degree.
func (s *OuterShadow) VerticalSkew() int {
	attr, found := s.GetAttribute("ky", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetVerticalSkew sets the vertical skew angle.
func (s *OuterShadow) SetVerticalSkew(skew int) {
	if skew == 0 {
		s.RemoveAttribute("ky", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"ky",
			"",
			strconv.Itoa(skew),
		),
	)
}

// Alignment returns the shadow alignment.
func (s *OuterShadow) Alignment() RectAlignValue {
	attr, found := s.GetAttribute("algn", "")
	if !found {
		return RectAlignBottom
	}

	return RectAlignValue(attr.Value())
}

// SetAlignment sets the shadow alignment.
func (s *OuterShadow) SetAlignment(
	align RectAlignValue,
) {
	if align == "" || align == RectAlignBottom {
		s.RemoveAttribute("algn", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"algn",
			"",
			string(align),
		),
	)
}

// RotateWithShape returns whether the shadow rotates with the shape.
func (s *OuterShadow) RotateWithShape() bool {
	attr, found := s.GetAttribute(
		"rotWithShape",
		"",
	)
	if !found {
		return true
	}

	return attr.Value() == "1" ||
		attr.Value() == attrTrue
}

// SetRotateWithShape sets whether the shadow rotates with the shape.
func (s *OuterShadow) SetRotateWithShape(
	rotate bool,
) {
	if rotate {
		s.RemoveAttribute("rotWithShape", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"rotWithShape",
			"",
			"0",
		),
	)
}

// removeColorElements removes all color child elements.
func (s *OuterShadow) removeColorElements() {
	if elem := s.GetElement("srgbClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("schemeClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("prstClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("sysClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("hslClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("scrgbClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
}

// SetRgbColor sets the shadow color to an RGB color.
func (s *OuterShadow) SetRgbColor(hex string) {
	s.removeColorElements()
	rgb := NewRgbColor(hex)
	s.AppendChild(rgb)
}

// SetSchemeColor sets the shadow color to a scheme color.
func (s *OuterShadow) SetSchemeColor(
	color SchemeColorValue,
) {
	s.removeColorElements()
	sc := NewSchemeColor(color)
	s.AppendChild(sc)
}

// RgbColor returns the RGB color if set, or nil.
func (s *OuterShadow) RgbColor() *RgbColor {
	elem := s.GetElement("srgbClr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if rgb, ok := elem.(*RgbColor); ok {
		return rgb
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RgbColor{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SchemeColor returns the scheme color if set, or nil.
func (s *OuterShadow) SchemeColor() *SchemeColor {
	elem := s.GetElement(
		"schemeClr",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if sc, ok := elem.(*SchemeColor); ok {
		return sc
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SchemeColor{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Clone creates a deep copy of this OuterShadow element.
func (s *OuterShadow) Clone() openxml.Element {
	return &OuterShadow{
		CompositeElementBase: s.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// InnerShadow represents an inner shadow effect (a:innerShdw).
type InnerShadow struct {
	*openxml.CompositeElementBase
}

// NewInnerShadow creates a new inner shadow effect.
func NewInnerShadow() *InnerShadow {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"innerShdw",
		PrefixMain,
	)

	return &InnerShadow{
		CompositeElementBase: elem,
	}
}

// NewInnerShadowWithParams creates an inner shadow with specific parameters.
func NewInnerShadowWithParams(
	blurRadius, distance EMU,
	direction int,
	colorHex string,
) *InnerShadow {
	shadow := NewInnerShadow()
	shadow.SetBlurRadius(blurRadius)
	shadow.SetDistance(distance)
	shadow.SetDirection(direction)
	shadow.SetRgbColor(colorHex)

	return shadow
}

// BlurRadius returns the blur radius in EMUs.
func (s *InnerShadow) BlurRadius() EMU {
	attr, found := s.GetAttribute("blurRad", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetBlurRadius sets the blur radius in EMUs.
func (s *InnerShadow) SetBlurRadius(radius EMU) {
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"blurRad",
			"",
			strconv.FormatInt(int64(radius), 10),
		),
	)
}

// Distance returns the shadow distance in EMUs.
func (s *InnerShadow) Distance() EMU {
	attr, found := s.GetAttribute("dist", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetDistance sets the shadow distance in EMUs.
func (s *InnerShadow) SetDistance(dist EMU) {
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"dist",
			"",
			strconv.FormatInt(int64(dist), 10),
		),
	)
}

// Direction returns the shadow direction in 60000ths of a degree.
func (s *InnerShadow) Direction() int {
	attr, found := s.GetAttribute("dir", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetDirection sets the shadow direction in 60000ths of a degree.
func (s *InnerShadow) SetDirection(dir int) {
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			strconv.Itoa(dir),
		),
	)
}

// removeColorElements removes all color child elements.
func (s *InnerShadow) removeColorElements() {
	if elem := s.GetElement("srgbClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("schemeClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("prstClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("sysClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("hslClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("scrgbClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
}

// SetRgbColor sets the shadow color to an RGB color.
func (s *InnerShadow) SetRgbColor(hex string) {
	s.removeColorElements()
	rgb := NewRgbColor(hex)
	s.AppendChild(rgb)
}

// SetSchemeColor sets the shadow color to a scheme color.
func (s *InnerShadow) SetSchemeColor(
	color SchemeColorValue,
) {
	s.removeColorElements()
	sc := NewSchemeColor(color)
	s.AppendChild(sc)
}

// RgbColor returns the RGB color if set, or nil.
func (s *InnerShadow) RgbColor() *RgbColor {
	elem := s.GetElement("srgbClr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if rgb, ok := elem.(*RgbColor); ok {
		return rgb
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RgbColor{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Clone creates a deep copy of this InnerShadow element.
func (s *InnerShadow) Clone() openxml.Element {
	return &InnerShadow{
		CompositeElementBase: s.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// PresetShadow represents a preset shadow effect (a:prstShdw).
type PresetShadow struct {
	*openxml.CompositeElementBase
}

// NewPresetShadow creates a new preset shadow with the specified preset value.
func NewPresetShadow(
	preset PresetShadowValue,
) *PresetShadow {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"prstShdw",
		PrefixMain,
	)
	ps := &PresetShadow{
		CompositeElementBase: elem,
	}
	ps.SetPreset(preset)

	return ps
}

// Preset returns the preset shadow value.
func (s *PresetShadow) Preset() PresetShadowValue {
	attr, found := s.GetAttribute("prst", "")
	if !found {
		return ""
	}

	return PresetShadowValue(attr.Value())
}

// SetPreset sets the preset shadow value.
func (s *PresetShadow) SetPreset(
	preset PresetShadowValue,
) {
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"prst",
			"",
			string(preset),
		),
	)
}

// Distance returns the shadow distance in EMUs.
func (s *PresetShadow) Distance() EMU {
	attr, found := s.GetAttribute("dist", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetDistance sets the shadow distance in EMUs.
func (s *PresetShadow) SetDistance(dist EMU) {
	if dist == 0 {
		s.RemoveAttribute("dist", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"dist",
			"",
			strconv.FormatInt(int64(dist), 10),
		),
	)
}

// Direction returns the shadow direction in 60000ths of a degree.
func (s *PresetShadow) Direction() int {
	attr, found := s.GetAttribute("dir", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetDirection sets the shadow direction in 60000ths of a degree.
func (s *PresetShadow) SetDirection(dir int) {
	if dir == 0 {
		s.RemoveAttribute("dir", "")

		return
	}
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			strconv.Itoa(dir),
		),
	)
}

// removeColorElements removes all color child elements.
func (s *PresetShadow) removeColorElements() {
	if elem := s.GetElement("srgbClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("schemeClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("prstClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("sysClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("hslClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
	if elem := s.GetElement("scrgbClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
}

// SetRgbColor sets the shadow color to an RGB color.
func (s *PresetShadow) SetRgbColor(hex string) {
	s.removeColorElements()
	rgb := NewRgbColor(hex)
	s.AppendChild(rgb)
}

// SetSchemeColor sets the shadow color to a scheme color.
func (s *PresetShadow) SetSchemeColor(
	color SchemeColorValue,
) {
	s.removeColorElements()
	sc := NewSchemeColor(color)
	s.AppendChild(sc)
}

// RgbColor returns the RGB color if set, or nil.
func (s *PresetShadow) RgbColor() *RgbColor {
	elem := s.GetElement("srgbClr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if rgb, ok := elem.(*RgbColor); ok {
		return rgb
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RgbColor{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SchemeColor returns the scheme color if set, or nil.
func (s *PresetShadow) SchemeColor() *SchemeColor {
	elem := s.GetElement(
		"schemeClr",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if sc, ok := elem.(*SchemeColor); ok {
		return sc
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SchemeColor{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Clone creates a deep copy of this PresetShadow element.
func (s *PresetShadow) Clone() openxml.Element {
	return &PresetShadow{
		CompositeElementBase: s.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// Glow represents a glow effect (a:glow).
type Glow struct {
	*openxml.CompositeElementBase
}

// NewGlow creates a new glow effect.
func NewGlow() *Glow {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"glow",
		PrefixMain,
	)

	return &Glow{CompositeElementBase: elem}
}

// NewGlowWithRadius creates a glow with the specified radius.
func NewGlowWithRadius(radius EMU) *Glow {
	glow := NewGlow()
	glow.SetRadius(radius)

	return glow
}

// NewGlowWithParams creates a glow with radius and color.
func NewGlowWithParams(
	radius EMU,
	colorHex string,
) *Glow {
	glow := NewGlow()
	glow.SetRadius(radius)
	glow.SetRgbColor(colorHex)

	return glow
}

// Radius returns the glow radius in EMUs.
func (g *Glow) Radius() EMU {
	attr, found := g.GetAttribute("rad", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetRadius sets the glow radius in EMUs.
func (g *Glow) SetRadius(radius EMU) {
	g.SetAttribute(
		openxml.NewAttribute(
			"",
			"rad",
			"",
			strconv.FormatInt(int64(radius), 10),
		),
	)
}

// removeColorElements removes all color child elements.
func (g *Glow) removeColorElements() {
	if elem := g.GetElement("srgbClr", NamespaceMain); elem != nil {
		g.RemoveChild(elem)
	}
	if elem := g.GetElement("schemeClr", NamespaceMain); elem != nil {
		g.RemoveChild(elem)
	}
	if elem := g.GetElement("prstClr", NamespaceMain); elem != nil {
		g.RemoveChild(elem)
	}
	if elem := g.GetElement("sysClr", NamespaceMain); elem != nil {
		g.RemoveChild(elem)
	}
	if elem := g.GetElement("hslClr", NamespaceMain); elem != nil {
		g.RemoveChild(elem)
	}
	if elem := g.GetElement("scrgbClr", NamespaceMain); elem != nil {
		g.RemoveChild(elem)
	}
}

// SetRgbColor sets the glow color to an RGB color.
func (g *Glow) SetRgbColor(hex string) {
	g.removeColorElements()
	rgb := NewRgbColor(hex)
	g.AppendChild(rgb)
}

// SetSchemeColor sets the glow color to a scheme color.
func (g *Glow) SetSchemeColor(
	color SchemeColorValue,
) {
	g.removeColorElements()
	sc := NewSchemeColor(color)
	g.AppendChild(sc)
}

// RgbColor returns the RGB color if set, or nil.
func (g *Glow) RgbColor() *RgbColor {
	elem := g.GetElement("srgbClr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if rgb, ok := elem.(*RgbColor); ok {
		return rgb
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RgbColor{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SchemeColor returns the scheme color if set, or nil.
func (g *Glow) SchemeColor() *SchemeColor {
	elem := g.GetElement(
		"schemeClr",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if sc, ok := elem.(*SchemeColor); ok {
		return sc
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SchemeColor{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Clone creates a deep copy of this Glow element.
func (g *Glow) Clone() openxml.Element {
	return &Glow{
		CompositeElementBase: g.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// SoftEdge represents a soft edge blur effect (a:softEdge).
type SoftEdge struct {
	*openxml.LeafElementBase
}

// NewSoftEdge creates a new soft edge effect with the specified radius.
func NewSoftEdge(radius EMU) *SoftEdge {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		"softEdge",
		PrefixMain,
	)
	se := &SoftEdge{LeafElementBase: elem}
	se.SetRadius(radius)

	return se
}

// Radius returns the soft edge radius in EMUs.
func (s *SoftEdge) Radius() EMU {
	attr, found := s.GetAttribute("rad", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetRadius sets the soft edge radius in EMUs.
func (s *SoftEdge) SetRadius(radius EMU) {
	s.SetAttribute(
		openxml.NewAttribute(
			"",
			"rad",
			"",
			strconv.FormatInt(int64(radius), 10),
		),
	)
}

// Clone creates a deep copy of this SoftEdge element.
func (s *SoftEdge) Clone() openxml.Element {
	return &SoftEdge{
		LeafElementBase: s.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// Reflection represents a reflection effect (a:reflection).
type Reflection struct {
	*openxml.LeafElementBase
}

// NewReflection creates a new reflection effect.
func NewReflection() *Reflection {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		"reflection",
		PrefixMain,
	)

	return &Reflection{LeafElementBase: elem}
}

// NewReflectionWithDefaults creates a reflection with typical default values.
func NewReflectionWithDefaults() *Reflection {
	refl := NewReflection()
	// Typical reflection defaults
	refl.SetStartOpacity(100000) // 100%
	refl.SetEndAlpha(0)          // 0%
	refl.SetStartPosition(0)
	refl.SetEndPosition(100000)
	refl.SetDirection(5400000) // downward

	return refl
}

// BlurRadius returns the blur radius in EMUs.
func (r *Reflection) BlurRadius() EMU {
	attr, found := r.GetAttribute("blurRad", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetBlurRadius sets the blur radius in EMUs.
func (r *Reflection) SetBlurRadius(radius EMU) {
	if radius == 0 {
		r.RemoveAttribute("blurRad", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"blurRad",
			"",
			strconv.FormatInt(int64(radius), 10),
		),
	)
}

// StartOpacity returns the start opacity (0-100000, where 100000 = 100%).
func (r *Reflection) StartOpacity() int {
	attr, found := r.GetAttribute("stA", "")
	if !found {
		return 100000
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetStartOpacity sets the start opacity.
func (r *Reflection) SetStartOpacity(
	opacity int,
) {
	if opacity == 100000 {
		r.RemoveAttribute("stA", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"stA",
			"",
			strconv.Itoa(opacity),
		),
	)
}

// EndAlpha returns the end alpha (0-100000, where 100000 = 100%).
func (r *Reflection) EndAlpha() int {
	attr, found := r.GetAttribute("endA", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetEndAlpha sets the end alpha.
func (r *Reflection) SetEndAlpha(alpha int) {
	if alpha == 0 {
		r.RemoveAttribute("endA", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"endA",
			"",
			strconv.Itoa(alpha),
		),
	)
}

// StartPosition returns the start position (0-100000).
func (r *Reflection) StartPosition() int {
	attr, found := r.GetAttribute("stPos", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetStartPosition sets the start position.
func (r *Reflection) SetStartPosition(pos int) {
	if pos == 0 {
		r.RemoveAttribute("stPos", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"stPos",
			"",
			strconv.Itoa(pos),
		),
	)
}

// EndPosition returns the end position (0-100000).
func (r *Reflection) EndPosition() int {
	attr, found := r.GetAttribute("endPos", "")
	if !found {
		return 100000
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetEndPosition sets the end position.
func (r *Reflection) SetEndPosition(pos int) {
	if pos == 100000 {
		r.RemoveAttribute("endPos", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"endPos",
			"",
			strconv.Itoa(pos),
		),
	)
}

// Distance returns the distance in EMUs.
func (r *Reflection) Distance() EMU {
	attr, found := r.GetAttribute("dist", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetDistance sets the distance in EMUs.
func (r *Reflection) SetDistance(dist EMU) {
	if dist == 0 {
		r.RemoveAttribute("dist", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"dist",
			"",
			strconv.FormatInt(int64(dist), 10),
		),
	)
}

// Direction returns the direction in 60000ths of a degree.
func (r *Reflection) Direction() int {
	attr, found := r.GetAttribute("dir", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetDirection sets the direction in 60000ths of a degree.
func (r *Reflection) SetDirection(dir int) {
	if dir == 0 {
		r.RemoveAttribute("dir", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"dir",
			"",
			strconv.Itoa(dir),
		),
	)
}

// FadeDirection returns the fade direction in 60000ths of a degree.
func (r *Reflection) FadeDirection() int {
	attr, found := r.GetAttribute("fadeDir", "")
	if !found {
		return 5400000
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetFadeDirection sets the fade direction in 60000ths of a degree.
func (r *Reflection) SetFadeDirection(dir int) {
	if dir == 5400000 {
		r.RemoveAttribute("fadeDir", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"fadeDir",
			"",
			strconv.Itoa(dir),
		),
	)
}

// HorizontalRatio returns the horizontal scaling factor (100000 = 100%).
func (r *Reflection) HorizontalRatio() int {
	attr, found := r.GetAttribute("sx", "")
	if !found {
		return 100000
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetHorizontalRatio sets the horizontal scaling factor.
func (r *Reflection) SetHorizontalRatio(
	ratio int,
) {
	if ratio == 100000 {
		r.RemoveAttribute("sx", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"sx",
			"",
			strconv.Itoa(ratio),
		),
	)
}

// VerticalRatio returns the vertical scaling factor (100000 = 100%).
func (r *Reflection) VerticalRatio() int {
	attr, found := r.GetAttribute("sy", "")
	if !found {
		return 100000
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetVerticalRatio sets the vertical scaling factor.
func (r *Reflection) SetVerticalRatio(ratio int) {
	if ratio == 100000 {
		r.RemoveAttribute("sy", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"sy",
			"",
			strconv.Itoa(ratio),
		),
	)
}

// HorizontalSkew returns the horizontal skew angle in 60000ths of a degree.
func (r *Reflection) HorizontalSkew() int {
	attr, found := r.GetAttribute("kx", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetHorizontalSkew sets the horizontal skew angle.
func (r *Reflection) SetHorizontalSkew(skew int) {
	if skew == 0 {
		r.RemoveAttribute("kx", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"kx",
			"",
			strconv.Itoa(skew),
		),
	)
}

// VerticalSkew returns the vertical skew angle in 60000ths of a degree.
func (r *Reflection) VerticalSkew() int {
	attr, found := r.GetAttribute("ky", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetVerticalSkew sets the vertical skew angle.
func (r *Reflection) SetVerticalSkew(skew int) {
	if skew == 0 {
		r.RemoveAttribute("ky", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"ky",
			"",
			strconv.Itoa(skew),
		),
	)
}

// Alignment returns the reflection alignment.
func (r *Reflection) Alignment() RectAlignValue {
	attr, found := r.GetAttribute("algn", "")
	if !found {
		return RectAlignBottom
	}

	return RectAlignValue(attr.Value())
}

// SetAlignment sets the reflection alignment.
func (r *Reflection) SetAlignment(
	align RectAlignValue,
) {
	if align == "" || align == RectAlignBottom {
		r.RemoveAttribute("algn", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"algn",
			"",
			string(align),
		),
	)
}

// RotateWithShape returns whether the reflection rotates with the shape.
func (r *Reflection) RotateWithShape() bool {
	attr, found := r.GetAttribute(
		"rotWithShape",
		"",
	)
	if !found {
		return true
	}

	return attr.Value() == "1" ||
		attr.Value() == attrTrue
}

// SetRotateWithShape sets whether the reflection rotates with the shape.
func (r *Reflection) SetRotateWithShape(
	rotate bool,
) {
	if rotate {
		r.RemoveAttribute("rotWithShape", "")

		return
	}
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"rotWithShape",
			"",
			"0",
		),
	)
}

// Clone creates a deep copy of this Reflection element.
func (r *Reflection) Clone() openxml.Element {
	return &Reflection{
		LeafElementBase: r.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// Blur represents a Gaussian blur effect (a:blur).
type Blur struct {
	*openxml.LeafElementBase
}

// NewBlur creates a new blur effect with the specified radius.
func NewBlur(radius EMU) *Blur {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		"blur",
		PrefixMain,
	)
	b := &Blur{LeafElementBase: elem}
	b.SetRadius(radius)

	return b
}

// Radius returns the blur radius in EMUs.
func (b *Blur) Radius() EMU {
	attr, found := b.GetAttribute("rad", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return EMU(val)
}

// SetRadius sets the blur radius in EMUs.
func (b *Blur) SetRadius(radius EMU) {
	b.SetAttribute(
		openxml.NewAttribute(
			"",
			"rad",
			"",
			strconv.FormatInt(int64(radius), 10),
		),
	)
}

// Grow returns whether the blur grows the bounds.
func (b *Blur) Grow() bool {
	attr, found := b.GetAttribute("grow", "")
	if !found {
		return true
	}

	return attr.Value() == "1" ||
		attr.Value() == attrTrue
}

// SetGrow sets whether the blur grows the bounds.
func (b *Blur) SetGrow(grow bool) {
	if grow {
		b.RemoveAttribute("grow", "")

		return
	}
	b.SetAttribute(
		openxml.NewAttribute("", "grow", "", "0"),
	)
}

// Clone creates a deep copy of this Blur element.
func (b *Blur) Clone() openxml.Element {
	return &Blur{
		LeafElementBase: b.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// FillOverlay represents a fill overlay effect (a:fillOverlay).
type FillOverlay struct {
	*openxml.CompositeElementBase
}

// NewFillOverlay creates a new fill overlay effect with the specified blend mode.
func NewFillOverlay(
	blend BlendModeValue,
) *FillOverlay {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"fillOverlay",
		PrefixMain,
	)
	fo := &FillOverlay{CompositeElementBase: elem}
	fo.SetBlend(blend)

	return fo
}

// Blend returns the blend mode.
func (f *FillOverlay) Blend() BlendModeValue {
	attr, found := f.GetAttribute("blend", "")
	if !found {
		return BlendModeOverlay
	}

	return BlendModeValue(attr.Value())
}

// SetBlend sets the blend mode.
func (f *FillOverlay) SetBlend(
	blend BlendModeValue,
) {
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"blend",
			"",
			string(blend),
		),
	)
}

// SetSolidFill sets the overlay fill to a solid color fill.
func (f *FillOverlay) SetSolidFill(
	fill *SolidFill,
) {
	f.removeFillElements()
	if fill != nil {
		f.AppendChild(fill)
	}
}

// SetGradientFill sets the overlay fill to a gradient fill.
func (f *FillOverlay) SetGradientFill(
	fill *GradientFill,
) {
	f.removeFillElements()
	if fill != nil {
		f.AppendChild(fill)
	}
}

// SetPatternFill sets the overlay fill to a pattern fill.
func (f *FillOverlay) SetPatternFill(
	fill *PatternFill,
) {
	f.removeFillElements()
	if fill != nil {
		f.AppendChild(fill)
	}
}

// SetNoFill sets the overlay to no fill.
func (f *FillOverlay) SetNoFill() {
	f.removeFillElements()
	noFill := NewNoFill()
	f.AppendChild(noFill)
}

// removeFillElements removes all fill child elements.
func (f *FillOverlay) removeFillElements() {
	if elem := f.GetElement("solidFill", NamespaceMain); elem != nil {
		f.RemoveChild(elem)
	}
	if elem := f.GetElement("gradFill", NamespaceMain); elem != nil {
		f.RemoveChild(elem)
	}
	if elem := f.GetElement("pattFill", NamespaceMain); elem != nil {
		f.RemoveChild(elem)
	}
	if elem := f.GetElement("blipFill", NamespaceMain); elem != nil {
		f.RemoveChild(elem)
	}
	if elem := f.GetElement("noFill", NamespaceMain); elem != nil {
		f.RemoveChild(elem)
	}
	if elem := f.GetElement("grpFill", NamespaceMain); elem != nil {
		f.RemoveChild(elem)
	}
}

// SolidFill returns the solid fill if set, or nil.
func (f *FillOverlay) SolidFill() *SolidFill {
	elem := f.GetElement(
		"solidFill",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if sf, ok := elem.(*SolidFill); ok {
		return sf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SolidFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Clone creates a deep copy of this FillOverlay element.
func (f *FillOverlay) Clone() openxml.Element {
	return &FillOverlay{
		CompositeElementBase: f.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
