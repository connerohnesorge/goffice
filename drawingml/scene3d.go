package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Scene3D represents the 3D scene properties (a:scene3d).
type Scene3D struct {
	*openxml.CompositeElementBase
}

// NewScene3D creates a new Scene3D element.
func NewScene3D() *Scene3D {
	return &Scene3D{
		CompositeElementBase: openxml.NewCompositeElement(
			NamespaceMain,
			"scene3d",
			PrefixMain,
		),
	}
}

// Camera returns the camera element.
func (s *Scene3D) Camera() *Camera {
	elem := s.GetElement("camera", NamespaceMain)
	if elem == nil {
		return nil
	}
	if c, ok := elem.(*Camera); ok {
		return c
	}
	return &Camera{CompositeElementBase: elem.(*openxml.CompositeElementBase)}
}

// SetCamera sets the camera element.
func (s *Scene3D) SetCamera(c *Camera) {
	if existing := s.GetElement("camera", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if c != nil {
		s.PrependChild(c)
	}
}

// LightRig returns the light rig element.
func (s *Scene3D) LightRig() *LightRig {
	elem := s.GetElement("lightRig", NamespaceMain)
	if elem == nil {
		return nil
	}
	if l, ok := elem.(*LightRig); ok {
		return l
	}
	return &LightRig{CompositeElementBase: elem.(*openxml.CompositeElementBase)}
}

// SetLightRig sets the light rig element.
func (s *Scene3D) SetLightRig(l *LightRig) {
	if existing := s.GetElement("lightRig", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if l != nil {
		s.AppendChild(l)
	}
}

// Backdrop returns the backdrop element.
func (s *Scene3D) Backdrop() *Backdrop {
	elem := s.GetElement("backdrop", NamespaceMain)
	if elem == nil {
		return nil
	}
	if b, ok := elem.(*Backdrop); ok {
		return b
	}
	return &Backdrop{CompositeElementBase: elem.(*openxml.CompositeElementBase)}
}

// SetBackdrop sets the backdrop element.
func (s *Scene3D) SetBackdrop(b *Backdrop) {
	if existing := s.GetElement("backdrop", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if b != nil {
		s.AppendChild(b)
	}
}

// Shape3D represents the 3D shape properties (a:sp3d).
type Shape3D struct {
	*openxml.CompositeElementBase
}

// NewShape3D creates a new Shape3D element.
func NewShape3D() *Shape3D {
	return &Shape3D{
		CompositeElementBase: openxml.NewCompositeElement(
			NamespaceMain,
			"sp3d",
			PrefixMain,
		),
	}
}

// BevelTop returns the top bevel.
func (s *Shape3D) BevelTop() *Bevel {
	elem := s.GetElement("bevelT", NamespaceMain)
	if elem == nil {
		return nil
	}
	if b, ok := elem.(*Bevel); ok {
		return b
	}
	return &Bevel{LeafElementBase: elem.(*openxml.LeafElementBase)}
}

// SetBevelTop sets the top bevel.
func (s *Shape3D) SetBevelTop(b *Bevel) {
	if existing := s.GetElement("bevelT", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if b != nil {
		s.PrependChild(b)
	}
}

// BevelBottom returns the bottom bevel.
func (s *Shape3D) BevelBottom() *Bevel {
	elem := s.GetElement("bevelB", NamespaceMain)
	if elem == nil {
		return nil
	}
	if b, ok := elem.(*Bevel); ok {
		return b
	}
	return &Bevel{LeafElementBase: elem.(*openxml.LeafElementBase)}
}

// SetBevelBottom sets the bottom bevel.
func (s *Shape3D) SetBevelBottom(b *Bevel) {
	if existing := s.GetElement("bevelB", NamespaceMain); existing != nil {
		s.RemoveChild(existing)
	}
	if b != nil {
		s.AppendChild(b)
	}
}

// ExtrusionColor returns the extrusion color container element.
func (s *Shape3D) ExtrusionColor() openxml.Element {
	return s.GetElement("extrusionClr", NamespaceMain)
}

// SetExtrusionColorRgb sets the extrusion color to an RGB color.
func (s *Shape3D) SetExtrusionColorRgb(hex string) {
	s.removeExtrusionColor()
	clr := NewRgbColor(hex)
	wrapper := openxml.NewCompositeElement(NamespaceMain, "extrusionClr", PrefixMain)
	wrapper.AppendChild(clr)
	s.AppendChild(wrapper)
}

// SetExtrusionColorScheme sets the extrusion color to a scheme color.
func (s *Shape3D) SetExtrusionColorScheme(color SchemeColorValue) {
	s.removeExtrusionColor()
	clr := NewSchemeColor(color)
	wrapper := openxml.NewCompositeElement(NamespaceMain, "extrusionClr", PrefixMain)
	wrapper.AppendChild(clr)
	s.AppendChild(wrapper)
}

func (s *Shape3D) removeExtrusionColor() {
	if elem := s.GetElement("extrusionClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
}

// ContourColor returns the contour color container element.
func (s *Shape3D) ContourColor() openxml.Element {
	return s.GetElement("contourClr", NamespaceMain)
}

// SetContourColorRgb sets the contour color to an RGB color.
func (s *Shape3D) SetContourColorRgb(hex string) {
	s.removeContourColor()
	clr := NewRgbColor(hex)
	wrapper := openxml.NewCompositeElement(NamespaceMain, "contourClr", PrefixMain)
	wrapper.AppendChild(clr)
	s.AppendChild(wrapper)
}

// SetContourColorScheme sets the contour color to a scheme color.
func (s *Shape3D) SetContourColorScheme(color SchemeColorValue) {
	s.removeContourColor()
	clr := NewSchemeColor(color)
	wrapper := openxml.NewCompositeElement(NamespaceMain, "contourClr", PrefixMain)
	wrapper.AppendChild(clr)
	s.AppendChild(wrapper)
}

func (s *Shape3D) removeContourColor() {
	if elem := s.GetElement("contourClr", NamespaceMain); elem != nil {
		s.RemoveChild(elem)
	}
}

// Z returns the z-coordinate in EMUs.
func (s *Shape3D) Z() EMU {
	attr, found := s.GetAttribute("z", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(attr.Value(), 10, 64)
	return EMU(val)
}

// SetZ sets the z-coordinate in EMUs.
func (s *Shape3D) SetZ(z EMU) {
	s.SetAttribute(openxml.NewAttribute("", "z", "", strconv.FormatInt(int64(z), 10)))
}

// ExtrusionHeight returns the extrusion height in EMUs.
func (s *Shape3D) ExtrusionHeight() EMU {
	attr, found := s.GetAttribute("extrusionH", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(attr.Value(), 10, 64)
	return EMU(val)
}

// SetExtrusionHeight sets the extrusion height in EMUs.
func (s *Shape3D) SetExtrusionHeight(h EMU) {
	s.SetAttribute(openxml.NewAttribute("", "extrusionH", "", strconv.FormatInt(int64(h), 10)))
}

// ContourWidth returns the contour width in EMUs.
func (s *Shape3D) ContourWidth() EMU {
	attr, found := s.GetAttribute("contourW", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(attr.Value(), 10, 64)
	return EMU(val)
}

// SetContourWidth sets the contour width in EMUs.
func (s *Shape3D) SetContourWidth(w EMU) {
	s.SetAttribute(openxml.NewAttribute("", "contourW", "", strconv.FormatInt(int64(w), 10)))
}

// PresetMaterial returns the preset material type.
func (s *Shape3D) PresetMaterial() string {
	attr, found := s.GetAttribute("prstMaterial", "")
	if !found {
		return "warmMatte"
	}
	return attr.Value()
}

// SetPresetMaterial sets the preset material type.
func (s *Shape3D) SetPresetMaterial(m string) {
	s.SetAttribute(openxml.NewAttribute("", "prstMaterial", "", m))
}

// Bevel represents a bevel (a:bevelT, a:bevelB).
type Bevel struct {
	*openxml.LeafElementBase
}

// NewBevel creates a new Bevel element.
func NewBevel(name string) *Bevel {
	return &Bevel{
		LeafElementBase: openxml.NewLeafElement(
			NamespaceMain,
			name, // bevelT or bevelB
			PrefixMain,
		),
	}
}

// Width returns the width in EMUs.
func (b *Bevel) Width() EMU {
	attr, found := b.GetAttribute("w", "")
	if !found {
		return 76200 // Default 6pt
	}
	val, _ := strconv.ParseInt(attr.Value(), 10, 64)
	return EMU(val)
}

// SetWidth sets the width in EMUs.
func (b *Bevel) SetWidth(w EMU) {
	b.SetAttribute(openxml.NewAttribute("", "w", "", strconv.FormatInt(int64(w), 10)))
}

// Height returns the height in EMUs.
func (b *Bevel) Height() EMU {
	attr, found := b.GetAttribute("h", "")
	if !found {
		return 76200 // Default 6pt
	}
	val, _ := strconv.ParseInt(attr.Value(), 10, 64)
	return EMU(val)
}

// SetHeight sets the height in EMUs.
func (b *Bevel) SetHeight(h EMU) {
	b.SetAttribute(openxml.NewAttribute("", "h", "", strconv.FormatInt(int64(h), 10)))
}

// Preset returns the preset bevel type.
func (b *Bevel) Preset() string {
	attr, found := b.GetAttribute("prst", "")
	if !found {
		return "circle"
	}
	return attr.Value()
}

// SetPreset sets the preset bevel type.
func (b *Bevel) SetPreset(p string) {
	b.SetAttribute(openxml.NewAttribute("", "prst", "", p))
}

// Camera represents a camera (a:camera).
type Camera struct {
	*openxml.CompositeElementBase
}

// NewCamera creates a new Camera element.
func NewCamera() *Camera {
	return &Camera{
		CompositeElementBase: openxml.NewCompositeElement(
			NamespaceMain,
			"camera",
			PrefixMain,
		),
	}
}

// Preset returns the preset camera type.
func (c *Camera) Preset() string {
	attr, found := c.GetAttribute("prst", "")
	if !found {
		return ""
	}
	return attr.Value()
}

// SetPreset sets the preset camera type.
func (c *Camera) SetPreset(p string) {
	c.SetAttribute(openxml.NewAttribute("", "prst", "", p))
}

// FieldOfView returns the field of view in 60000ths of a degree.
func (c *Camera) FieldOfView() int {
	attr, found := c.GetAttribute("fov", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetFieldOfView sets the field of view in 60000ths of a degree.
func (c *Camera) SetFieldOfView(fov int) {
	c.SetAttribute(openxml.NewAttribute("", "fov", "", strconv.Itoa(fov)))
}

// Zoom returns the zoom percentage.
func (c *Camera) Zoom() int {
	attr, found := c.GetAttribute("zoom", "")
	if !found {
		return 100000
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetZoom sets the zoom percentage.
func (c *Camera) SetZoom(z int) {
	c.SetAttribute(openxml.NewAttribute("", "zoom", "", strconv.Itoa(z)))
}

// Rotation returns the rotation element.
func (c *Camera) Rotation() *Rotation3D {
	elem := c.GetElement("rot", NamespaceMain)
	if elem == nil {
		return nil
	}
	if r, ok := elem.(*Rotation3D); ok {
		return r
	}
	return &Rotation3D{LeafElementBase: elem.(*openxml.LeafElementBase)}
}

// SetRotation sets the rotation element.
func (c *Camera) SetRotation(r *Rotation3D) {
	if existing := c.GetElement("rot", NamespaceMain); existing != nil {
		c.RemoveChild(existing)
	}
	if r != nil {
		c.AppendChild(r)
	}
}

// LightRig represents a light rig (a:lightRig).
type LightRig struct {
	*openxml.CompositeElementBase
}

// NewLightRig creates a new LightRig element.
func NewLightRig() *LightRig {
	return &LightRig{
		CompositeElementBase: openxml.NewCompositeElement(
			NamespaceMain,
			"lightRig",
			PrefixMain,
		),
	}
}

// Rig returns the rig preset type.
func (l *LightRig) Rig() string {
	attr, found := l.GetAttribute("rig", "")
	if !found {
		return "legacyFlat1"
	}
	return attr.Value()
}

// SetRig sets the rig preset type.
func (l *LightRig) SetRig(r string) {
	l.SetAttribute(openxml.NewAttribute("", "rig", "", r))
}

// Direction returns the light direction.
func (l *LightRig) Direction() string {
	attr, found := l.GetAttribute("dir", "")
	if !found {
		return "t"
	}
	return attr.Value()
}

// SetDirection sets the light direction.
func (l *LightRig) SetDirection(d string) {
	l.SetAttribute(openxml.NewAttribute("", "dir", "", d))
}

// Rotation returns the rotation element.
func (l *LightRig) Rotation() *Rotation3D {
	elem := l.GetElement("rot", NamespaceMain)
	if elem == nil {
		return nil
	}
	if r, ok := elem.(*Rotation3D); ok {
		return r
	}
	return &Rotation3D{LeafElementBase: elem.(*openxml.LeafElementBase)}
}

// SetRotation sets the rotation element.
func (l *LightRig) SetRotation(r *Rotation3D) {
	if existing := l.GetElement("rot", NamespaceMain); existing != nil {
		l.RemoveChild(existing)
	}
	if r != nil {
		l.AppendChild(r)
	}
}

// Rotation3D represents a 3D rotation (a:rot).
type Rotation3D struct {
	*openxml.LeafElementBase
}

// NewRotation3D creates a new Rotation3D element.
func NewRotation3D() *Rotation3D {
	return &Rotation3D{
		LeafElementBase: openxml.NewLeafElement(
			NamespaceMain,
			"rot",
			PrefixMain,
		),
	}
}

// Lat returns the latitude in 60000ths of a degree.
func (r *Rotation3D) Lat() int {
	attr, found := r.GetAttribute("lat", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetLat sets the latitude.
func (r *Rotation3D) SetLat(lat int) {
	r.SetAttribute(openxml.NewAttribute("", "lat", "", strconv.Itoa(lat)))
}

// Lon returns the longitude in 60000ths of a degree.
func (r *Rotation3D) Lon() int {
	attr, found := r.GetAttribute("lon", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetLon sets the longitude.
func (r *Rotation3D) SetLon(lon int) {
	r.SetAttribute(openxml.NewAttribute("", "lon", "", strconv.Itoa(lon)))
}

// Rev returns the revolution in 60000ths of a degree.
func (r *Rotation3D) Rev() int {
	attr, found := r.GetAttribute("rev", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetRev sets the revolution.
func (r *Rotation3D) SetRev(rev int) {
	r.SetAttribute(openxml.NewAttribute("", "rev", "", strconv.Itoa(rev)))
}

// Backdrop represents a backdrop (a:backdrop).
type Backdrop struct {
	*openxml.CompositeElementBase
}

// NewBackdrop creates a new Backdrop element.
func NewBackdrop() *Backdrop {
	return &Backdrop{
		CompositeElementBase: openxml.NewCompositeElement(
			NamespaceMain,
			"backdrop",
			PrefixMain,
		),
	}
}

// Anchor returns the anchor point coordinates (x, y, z).
// Implementation simplified for brevity.
