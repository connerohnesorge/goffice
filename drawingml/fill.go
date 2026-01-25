// and effects.
//
//nolint:revive // This file contains many public types for OOXML fill elements.
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PatternFillValue represents pattern fill preset values.
type PatternFillValue string

// Pattern fill values defined in DrawingML.
const (
	PatternPct5              PatternFillValue = "pct5"
	PatternPct10             PatternFillValue = "pct10"
	PatternPct20             PatternFillValue = "pct20"
	PatternPct25             PatternFillValue = "pct25"
	PatternPct30             PatternFillValue = "pct30"
	PatternPct40             PatternFillValue = "pct40"
	PatternPct50             PatternFillValue = "pct50"
	PatternPct60             PatternFillValue = "pct60"
	PatternPct70             PatternFillValue = "pct70"
	PatternPct75             PatternFillValue = "pct75"
	PatternPct80             PatternFillValue = "pct80"
	PatternPct90             PatternFillValue = "pct90"
	PatternHorizontal        PatternFillValue = "horz"
	PatternVertical          PatternFillValue = "vert"
	PatternLightHorizontal   PatternFillValue = "ltHorz"
	PatternLightVertical     PatternFillValue = "ltVert"
	PatternDarkHorizontal    PatternFillValue = "dkHorz"
	PatternDarkVertical      PatternFillValue = "dkVert"
	PatternNarrowHorizontal  PatternFillValue = "narHorz"
	PatternNarrowVertical    PatternFillValue = "narVert"
	PatternDashedHorizontal  PatternFillValue = "dashHorz"
	PatternDashedVertical    PatternFillValue = "dashVert"
	PatternCross             PatternFillValue = "cross"
	PatternDownwardDiagonal  PatternFillValue = "dnDiag"
	PatternUpwardDiagonal    PatternFillValue = "upDiag"
	PatternLightDownwardDiag PatternFillValue = "ltDnDiag"
	PatternLightUpwardDiag   PatternFillValue = "ltUpDiag"
	PatternDarkDownwardDiag  PatternFillValue = "dkDnDiag"
	PatternDarkUpwardDiag    PatternFillValue = "dkUpDiag"
	PatternWideDnDiag        PatternFillValue = "wdDnDiag"
	PatternWideUpDiag        PatternFillValue = "wdUpDiag"
	PatternDashedDnDiag      PatternFillValue = "dashDnDiag"
	PatternDashedUpDiag      PatternFillValue = "dashUpDiag"
	PatternDiagonalCross     PatternFillValue = "diagCross"
	PatternSmallCheck        PatternFillValue = "smCheck"
	PatternLargeCheck        PatternFillValue = "lgCheck"
	PatternSmallGrid         PatternFillValue = "smGrid"
	PatternLargeGrid         PatternFillValue = "lgGrid"
	PatternDottedGrid        PatternFillValue = "dotGrid"
	PatternSmallConfetti     PatternFillValue = "smConfetti"
	PatternLargeConfetti     PatternFillValue = "lgConfetti"
	PatternHorzBrick         PatternFillValue = "horzBrick"
	PatternDiagBrick         PatternFillValue = "diagBrick"
	PatternSolidDiamond      PatternFillValue = "solidDmnd"
	PatternOpenDiamond       PatternFillValue = "openDmnd"
	PatternDottedDiamond     PatternFillValue = "dotDmnd"
	PatternPlaid             PatternFillValue = "plaid"
	PatternSphere            PatternFillValue = "sphere"
	PatternWeave             PatternFillValue = "weave"
	PatternDivot             PatternFillValue = "divot"
	PatternShingle           PatternFillValue = "shingle"
	PatternWave              PatternFillValue = "wave"
	PatternTrellis           PatternFillValue = "trellis"
	PatternZigZag            PatternFillValue = "zigZag"
)

// TileFlipValue represents how a tile fill should be flipped.
type TileFlipValue string

// Tile flip values for gradient and image fills.
const (
	TileFlipNone TileFlipValue = "none"
	TileFlipX    TileFlipValue = "x"
	TileFlipY    TileFlipValue = "y"
	TileFlipXY   TileFlipValue = "xy"
)

// RectAlignValue represents rectangle alignment for positioning.
type RectAlignValue string

// Rectangle alignment values.
const (
	RectAlignTopLeft     RectAlignValue = "tl"
	RectAlignTop         RectAlignValue = "t"
	RectAlignTopRight    RectAlignValue = "tr"
	RectAlignLeft        RectAlignValue = "l"
	RectAlignCenter      RectAlignValue = "ctr"
	RectAlignRight       RectAlignValue = "r"
	RectAlignBottomLeft  RectAlignValue = "bl"
	RectAlignBottom      RectAlignValue = "b"
	RectAlignBottomRight RectAlignValue = "br"
)

// PathShadeValue represents the path shape for gradient fills.
type PathShadeValue string

// Path shade values for gradient fills.
const (
	PathShapeShape  PathShadeValue = "shape"
	PathShapeCircle PathShadeValue = "circle"
	PathShapeRect   PathShadeValue = "rect"
)

// NoFill represents a no fill element (a:noFill).
// This indicates that no fill should be applied to a shape.
type NoFill struct {
	*openxml.LeafElementBase
}

// NewNoFill creates a new no fill element.
func NewNoFill() *NoFill {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		"noFill",
		PrefixMain,
	)

	return &NoFill{LeafElementBase: elem}
}

// Clone creates a deep copy of this NoFill element.
func (f *NoFill) Clone() openxml.Element {
	return &NoFill{
		LeafElementBase: f.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// SolidFill represents a solid color fill (a:solidFill).
type SolidFill struct {
	*openxml.CompositeElementBase
}

// NewSolidFill creates a new solid fill element without a color.
func NewSolidFill() *SolidFill {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"solidFill",
		PrefixMain,
	)

	return &SolidFill{CompositeElementBase: elem}
}

// NewSolidFillWithRgb creates a new solid fill with an RGB color.
func NewSolidFillWithRgb(
	hexColor string,
) *SolidFill {
	sf := NewSolidFill()
	sf.SetRgbColor(hexColor)

	return sf
}

// NewSolidFillWithSchemeColor creates a new solid fill with a scheme color.
func NewSolidFillWithSchemeColor(
	color SchemeColorValue,
) *SolidFill {
	sf := NewSolidFill()
	sf.SetSchemeColor(color)

	return sf
}

// SetRgbColor sets the fill color to an RGB color.
func (f *SolidFill) SetRgbColor(hexColor string) {
	// Remove any existing color
	f.RemoveAllChildren()
	// Add the RGB color
	rgb := NewRgbColor(hexColor)
	f.AppendChild(rgb)
}

// SetSchemeColor sets the fill color to a scheme color.
func (f *SolidFill) SetSchemeColor(
	color SchemeColorValue,
) {
	// Remove any existing color
	f.RemoveAllChildren()
	// Add the scheme color
	sc := NewSchemeColor(color)
	f.AppendChild(sc)
}

// SetPresetColor sets the fill color to a preset color.
func (f *SolidFill) SetPresetColor(
	color PresetColorValue,
) {
	// Remove any existing color
	f.RemoveAllChildren()
	// Add the preset color
	pc := NewPresetColor(color)
	f.AppendChild(pc)
}

// RgbColor returns the RGB color if set, or nil.
func (f *SolidFill) RgbColor() *RgbColor {
	elem := f.GetElement("srgbClr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if rgb, ok := elem.(*RgbColor); ok {
		return rgb
	}
	// Wrap the composite element
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RgbColor{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SchemeColor returns the scheme color if set, or nil.
func (f *SolidFill) SchemeColor() *SchemeColor {
	elem := f.GetElement(
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

// Clone creates a deep copy of this SolidFill element.
func (f *SolidFill) Clone() openxml.Element {
	return &SolidFill{
		CompositeElementBase: f.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// GradientStop represents a gradient stop (a:gs) in a gradient fill.
type GradientStop struct {
	*openxml.CompositeElementBase
}

// NewGradientStop creates a new gradient stop at the specified position.
// position is in 1000ths of a percent (0-100000).
func NewGradientStop(position int) *GradientStop {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"gs",
		PrefixMain,
	)
	gs := &GradientStop{
		CompositeElementBase: elem,
	}
	gs.SetPosition(position)

	return gs
}

// NewGradientStopWithRgb creates a new gradient stop with an RGB color.
func NewGradientStopWithRgb(
	position int,
	hexColor string,
) *GradientStop {
	gs := NewGradientStop(position)
	rgb := NewRgbColor(hexColor)
	gs.AppendChild(rgb)

	return gs
}

// NewGradientStopWithSchemeColor creates a new gradient stop with a scheme color.
func NewGradientStopWithSchemeColor(
	position int,
	color SchemeColorValue,
) *GradientStop {
	gs := NewGradientStop(position)
	sc := NewSchemeColor(color)
	gs.AppendChild(sc)

	return gs
}

// Position returns the position in 1000ths of a percent.
func (g *GradientStop) Position() int {
	attr, found := g.GetAttribute("pos", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetPosition sets the position in 1000ths of a percent.
func (g *GradientStop) SetPosition(position int) {
	g.SetAttribute(
		openxml.NewAttribute(
			"",
			"pos",
			"",
			strconv.Itoa(position),
		),
	)
}

// PositionPercent returns the position as a percentage (0-100).
func (g *GradientStop) PositionPercent() float64 {
	return UnitsToPercent(g.Position())
}

// Clone creates a deep copy of this GradientStop element.
func (g *GradientStop) Clone() openxml.Element {
	return &GradientStop{
		CompositeElementBase: g.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// GradientFill represents a gradient fill (a:gradFill).
type GradientFill struct {
	*openxml.CompositeElementBase
}

// NewGradientFill creates a new gradient fill element.
func NewGradientFill() *GradientFill {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"gradFill",
		PrefixMain,
	)

	return &GradientFill{
		CompositeElementBase: elem,
	}
}

// NewLinearGradientFill creates a linear gradient fill with the specified angle.
// angle is in 60000ths of a degree.
func NewLinearGradientFill(
	angle int,
) *GradientFill {
	gf := NewGradientFill()
	gf.SetLinear(angle, false)

	return gf
}

// AddStop adds a gradient stop to the fill.
func (g *GradientFill) AddStop(
	stop *GradientStop,
) {
	// Find or create the gradient stop list
	gsLst := g.GetElement("gsLst", NamespaceMain)
	if gsLst == nil {
		gsLst = openxml.NewCompositeElement(
			NamespaceMain,
			"gsLst",
			PrefixMain,
		)
		g.PrependChild(gsLst)
	}
	if comp, ok := gsLst.(*openxml.CompositeElementBase); ok {
		comp.AppendChild(stop)
	}
}

// AddRgbStop adds a gradient stop with an RGB color.
func (g *GradientFill) AddRgbStop(
	position int,
	hexColor string,
) {
	stop := NewGradientStopWithRgb(
		position,
		hexColor,
	)
	g.AddStop(stop)
}

// AddSchemeColorStop adds a gradient stop with a scheme color.
func (g *GradientFill) AddSchemeColorStop(
	position int,
	color SchemeColorValue,
) {
	stop := NewGradientStopWithSchemeColor(
		position,
		color,
	)
	g.AddStop(stop)
}

// SetLinear sets the gradient to a linear gradient with the specified angle.
// angle is in 60000ths of a degree (e.g., 5400000 for 90 degrees).
// scaled indicates whether the gradient is scaled with the shape.
func (g *GradientFill) SetLinear(
	angle int,
	scaled bool,
) {
	// Remove any existing path or tileRect
	g.removeGradientType()

	lin := openxml.NewLeafElement(
		NamespaceMain,
		"lin",
		PrefixMain,
	)
	lin.SetAttribute(
		openxml.NewAttribute(
			"",
			"ang",
			"",
			strconv.Itoa(angle),
		),
	)
	if scaled {
		lin.SetAttribute(
			openxml.NewAttribute(
				"",
				"scaled",
				"",
				"1",
			),
		)
	}
	g.AppendChild(lin)
}

// SetPath sets the gradient to a path gradient with the specified shape.
func (g *GradientFill) SetPath(
	pathType PathShadeValue,
) {
	g.removeGradientType()

	path := openxml.NewCompositeElement(
		NamespaceMain,
		"path",
		PrefixMain,
	)
	path.SetAttribute(
		openxml.NewAttribute(
			"",
			"path",
			"",
			string(pathType),
		),
	)
	g.AppendChild(path)
}

// removeGradientType removes existing gradient type elements (lin, path).
func (g *GradientFill) removeGradientType() {
	if lin := g.GetElement("lin", NamespaceMain); lin != nil {
		g.RemoveChild(lin)
	}
	if path := g.GetElement("path", NamespaceMain); path != nil {
		g.RemoveChild(path)
	}
}

// RotateWithShape returns whether the gradient rotates with the shape.
func (g *GradientFill) RotateWithShape() bool {
	attr, found := g.GetAttribute(
		"rotWithShape",
		"",
	)
	if !found {
		return true // Default is true
	}

	return attr.Value() == "1" ||
		attr.Value() == attrTrue
}

// SetRotateWithShape sets whether the gradient rotates with the shape.
func (g *GradientFill) SetRotateWithShape(
	rotate bool,
) {
	if rotate {
		g.RemoveAttribute("rotWithShape", "")

		return
	}
	g.SetAttribute(
		openxml.NewAttribute(
			"",
			"rotWithShape",
			"",
			"0",
		),
	)
}

// TileFlip returns the tile flip mode for the gradient.
func (g *GradientFill) TileFlip() TileFlipValue {
	attr, found := g.GetAttribute("flip", "")
	if !found {
		return TileFlipNone
	}

	return TileFlipValue(attr.Value())
}

// SetTileFlip sets the tile flip mode for the gradient.
func (g *GradientFill) SetTileFlip(
	flip TileFlipValue,
) {
	if flip == TileFlipNone || flip == "" {
		g.RemoveAttribute("flip", "")

		return
	}
	g.SetAttribute(
		openxml.NewAttribute(
			"",
			"flip",
			"",
			string(flip),
		),
	)
}

// Clone creates a deep copy of this GradientFill element.
func (g *GradientFill) Clone() openxml.Element {
	return &GradientFill{
		CompositeElementBase: g.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// PatternFill represents a pattern fill (a:pattFill).
type PatternFill struct {
	*openxml.CompositeElementBase
}

// NewPatternFill creates a new pattern fill with the specified preset pattern.
func NewPatternFill(
	preset PatternFillValue,
) *PatternFill {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"pattFill",
		PrefixMain,
	)
	pf := &PatternFill{CompositeElementBase: elem}
	pf.SetPreset(preset)

	return pf
}

// Preset returns the pattern fill preset value.
func (p *PatternFill) Preset() PatternFillValue {
	attr, found := p.GetAttribute("prst", "")
	if !found {
		return ""
	}

	return PatternFillValue(attr.Value())
}

// SetPreset sets the pattern fill preset value.
func (p *PatternFill) SetPreset(
	preset PatternFillValue,
) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"prst",
			"",
			string(preset),
		),
	)
}

// SetForegroundColor sets the foreground color of the pattern.
func (p *PatternFill) SetForegroundColor(
	hexColor string,
) {
	// Remove existing foreground color
	if fg := p.GetElement("fgClr", NamespaceMain); fg != nil {
		p.RemoveChild(fg)
	}

	fgClr := openxml.NewCompositeElement(
		NamespaceMain,
		"fgClr",
		PrefixMain,
	)
	rgb := NewRgbColor(hexColor)
	fgClr.AppendChild(rgb)
	p.AppendChild(fgClr)
}

// SetBackgroundColor sets the background color of the pattern.
func (p *PatternFill) SetBackgroundColor(
	hexColor string,
) {
	// Remove existing background color
	if bg := p.GetElement("bgClr", NamespaceMain); bg != nil {
		p.RemoveChild(bg)
	}

	bgClr := openxml.NewCompositeElement(
		NamespaceMain,
		"bgClr",
		PrefixMain,
	)
	rgb := NewRgbColor(hexColor)
	bgClr.AppendChild(rgb)
	p.AppendChild(bgClr)
}

// Clone creates a deep copy of this PatternFill element.
func (p *PatternFill) Clone() openxml.Element {
	return &PatternFill{
		CompositeElementBase: p.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// BlipFill represents an image/picture fill (a:blipFill).
type BlipFill struct {
	*openxml.CompositeElementBase
}

// NewBlipFill creates a new blip (image) fill element.
func NewBlipFill() *BlipFill {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"blipFill",
		PrefixMain,
	)

	return &BlipFill{CompositeElementBase: elem}
}

// NewBlipFillWithEmbed creates a new blip fill with an embedded relationship ID.
func NewBlipFillWithEmbed(
	embedId string,
) *BlipFill {
	bf := NewBlipFill()
	bf.SetEmbed(embedId)

	return bf
}

// SetEmbed sets the embedded relationship ID for the image.
func (b *BlipFill) SetEmbed(embedId string) {
	blip := b.Blip()
	blip.SetEmbed(embedId)
}

// Blip returns the Blip element, creating it if necessary.
func (b *BlipFill) Blip() *Blip {
	elem := b.GetElement("blip", NamespaceMain)
	if elem == nil {
		blip := NewBlip()
		b.PrependChild(blip)
		return blip
	}
	if blip, ok := elem.(*Blip); ok {
		return blip
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Blip{CompositeElementBase: comp}
	}
	// If it's a leaf element (legacy), we need to replace it with composite to support effects
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		comp := openxml.NewCompositeElement(NamespaceMain, "blip", PrefixMain)
		// Copy attributes
		for _, attr := range leaf.Attributes() {
			comp.SetAttribute(attr)
		}
		// Replace in parent
		b.RemoveChild(leaf)
		b.PrependChild(comp)
		return &Blip{CompositeElementBase: comp}
	}
	return nil
}

// Embed returns the embedded relationship ID.
func (b *BlipFill) Embed() string {
	blip := b.Blip()
	return blip.Embed()
}

// SetStretch sets the fill to stretch mode (fills entire shape).
func (b *BlipFill) SetStretch() {
	// Remove any existing tile or stretch
	if tile := b.GetElement("tile", NamespaceMain); tile != nil {
		b.RemoveChild(tile)
	}
	if stretch := b.GetElement("stretch", NamespaceMain); stretch != nil {
		b.RemoveChild(stretch)
	}

	stretch := openxml.NewCompositeElement(
		NamespaceMain,
		"stretch",
		PrefixMain,
	)
	fillRect := openxml.NewLeafElement(
		NamespaceMain,
		"fillRect",
		PrefixMain,
	)
	stretch.AppendChild(fillRect)
	b.AppendChild(stretch)
}

// SetTile sets the fill to tile mode.
func (b *BlipFill) SetTile(
	alignment RectAlignValue,
	flip TileFlipValue,
) {
	// Remove any existing tile or stretch
	if tile := b.GetElement("tile", NamespaceMain); tile != nil {
		b.RemoveChild(tile)
	}
	if stretch := b.GetElement("stretch", NamespaceMain); stretch != nil {
		b.RemoveChild(stretch)
	}

	tile := openxml.NewLeafElement(
		NamespaceMain,
		"tile",
		PrefixMain,
	)
	if alignment != "" {
		tile.SetAttribute(
			openxml.NewAttribute(
				"",
				"algn",
				"",
				string(alignment),
			),
		)
	}
	if flip != "" && flip != TileFlipNone {
		tile.SetAttribute(
			openxml.NewAttribute(
				"",
				"flip",
				"",
				string(flip),
			),
		)
	}
	b.AppendChild(tile)
}

// Dpi returns the DPI setting for the image.
func (b *BlipFill) Dpi() int {
	attr, found := b.GetAttribute("dpi", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetDpi sets the DPI for the image.
func (b *BlipFill) SetDpi(dpi int) {
	if dpi <= 0 {
		b.RemoveAttribute("dpi", "")

		return
	}
	b.SetAttribute(
		openxml.NewAttribute(
			"",
			"dpi",
			"",
			strconv.Itoa(dpi),
		),
	)
}

// RotateWithShape returns whether the fill rotates with the shape.
func (b *BlipFill) RotateWithShape() bool {
	attr, found := b.GetAttribute(
		"rotWithShape",
		"",
	)
	if !found {
		return true
	}

	return attr.Value() == "1" ||
		attr.Value() == attrTrue
}

// SetRotateWithShape sets whether the fill rotates with the shape.
func (b *BlipFill) SetRotateWithShape(
	rotate bool,
) {
	if rotate {
		b.RemoveAttribute("rotWithShape", "")

		return
	}
	b.SetAttribute(
		openxml.NewAttribute(
			"",
			"rotWithShape",
			"",
			"0",
		),
	)
}

// SourceRect represents the portion of the image to display (cropping).
// Values are percentages (0-100000 where 100000 = 100%).
type SourceRect struct {
	Left   int // Left crop percentage
	Top    int // Top crop percentage
	Right  int // Right crop percentage
	Bottom int // Bottom crop percentage
}

// SetSourceRect sets the source rectangle for cropping the image.
// All values are percentages where 100000 = 100%.
// For example, Left=10000 crops 10% from the left side.
func (b *BlipFill) SetSourceRect(rect SourceRect) {
	// Remove existing srcRect
	if existing := b.GetElement("srcRect", NamespaceMain); existing != nil {
		b.RemoveChild(existing)
	}

	// Create new srcRect element
	srcRect := openxml.NewLeafElement(
		NamespaceMain,
		"srcRect",
		PrefixMain,
	)

	// Set attributes only if non-zero
	if rect.Left != 0 {
		srcRect.SetAttribute(openxml.NewAttribute("", "l", "", strconv.Itoa(rect.Left)))
	}
	if rect.Top != 0 {
		srcRect.SetAttribute(openxml.NewAttribute("", "t", "", strconv.Itoa(rect.Top)))
	}
	if rect.Right != 0 {
		srcRect.SetAttribute(openxml.NewAttribute("", "r", "", strconv.Itoa(rect.Right)))
	}
	if rect.Bottom != 0 {
		srcRect.SetAttribute(openxml.NewAttribute("", "b", "", strconv.Itoa(rect.Bottom)))
	}

	// Insert after blip element if it exists
	if blip := b.GetElement("blip", NamespaceMain); blip != nil {
		b.InsertAfter(srcRect, blip)
	} else {
		b.PrependChild(srcRect)
	}
}

// SourceRect returns the source rectangle for cropping, or nil if not set.
func (b *BlipFill) SourceRect() *SourceRect {
	srcRect := b.GetElement("srcRect", NamespaceMain)
	if srcRect == nil {
		return nil
	}

	rect := &SourceRect{}
	if attr, found := srcRect.GetAttribute("l", ""); found {
		rect.Left, _ = strconv.Atoi(attr.Value())
	}
	if attr, found := srcRect.GetAttribute("t", ""); found {
		rect.Top, _ = strconv.Atoi(attr.Value())
	}
	if attr, found := srcRect.GetAttribute("r", ""); found {
		rect.Right, _ = strconv.Atoi(attr.Value())
	}
	if attr, found := srcRect.GetAttribute("b", ""); found {
		rect.Bottom, _ = strconv.Atoi(attr.Value())
	}

	return rect
}

// ClearSourceRect removes the source rectangle (removes cropping).
func (b *BlipFill) ClearSourceRect() {
	if srcRect := b.GetElement("srcRect", NamespaceMain); srcRect != nil {
		b.RemoveChild(srcRect)
	}
}

// Clone creates a deep copy of this BlipFill element.
func (b *BlipFill) Clone() openxml.Element {
	return &BlipFill{
		CompositeElementBase: b.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// GroupFill represents a group fill (a:grpFill).
// This indicates that the fill should be inherited from the group.
type GroupFill struct {
	*openxml.LeafElementBase
}

// NewGroupFill creates a new group fill element.
func NewGroupFill() *GroupFill {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		"grpFill",
		PrefixMain,
	)

	return &GroupFill{LeafElementBase: elem}
}

// Clone creates a deep copy of this GroupFill element.
func (f *GroupFill) Clone() openxml.Element {
	return &GroupFill{
		LeafElementBase: f.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}
