// Package drawingml provides shared DrawingML types for shapes, images,
// and effects.
//
//nolint:revive // This file contains many public types for OOXML geometry elements.
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// ShapeTypeValue represents preset shape type values.
type ShapeTypeValue string

// Preset shape type values defined in DrawingML.
// These represent the standard shapes available in Office applications.
const (
	// Basic Shapes
	ShapeTypeRectangle      ShapeTypeValue = "rect"
	ShapeTypeRoundRectangle ShapeTypeValue = "roundRect"
	ShapeTypeSnipRoundRect  ShapeTypeValue = "snipRoundRect"
	ShapeTypeSnip1Rect      ShapeTypeValue = "snip1Rect"
	ShapeTypeSnip2SameRect  ShapeTypeValue = "snip2SameRect"
	ShapeTypeSnip2DiagRect  ShapeTypeValue = "snip2DiagRect"
	ShapeTypeRound1Rect     ShapeTypeValue = "round1Rect"
	ShapeTypeRound2SameRect ShapeTypeValue = "round2SameRect"
	ShapeTypeRound2DiagRect ShapeTypeValue = "round2DiagRect"
	ShapeTypeEllipse        ShapeTypeValue = "ellipse"
	ShapeTypeTriangle       ShapeTypeValue = "triangle"
	ShapeTypeRightTriangle  ShapeTypeValue = "rtTriangle"
	ShapeTypeParallelogram  ShapeTypeValue = "parallelogram"
	ShapeTypeTrapezoid      ShapeTypeValue = "trapezoid"
	ShapeTypeDiamond        ShapeTypeValue = "diamond"
	ShapeTypePentagon       ShapeTypeValue = "pentagon"
	ShapeTypeHexagon        ShapeTypeValue = "hexagon"
	ShapeTypeHeptagon       ShapeTypeValue = "heptagon"
	ShapeTypeOctagon        ShapeTypeValue = "octagon"
	ShapeTypeDecagon        ShapeTypeValue = "decagon"
	ShapeTypeDodecagon      ShapeTypeValue = "dodecagon"
	ShapeTypePie            ShapeTypeValue = "pie"
	ShapeTypeChord          ShapeTypeValue = "chord"
	ShapeTypeTeardrop       ShapeTypeValue = "teardrop"
	ShapeTypeFrame          ShapeTypeValue = "frame"
	ShapeTypeHalfFrame      ShapeTypeValue = "halfFrame"
	ShapeTypeCorner         ShapeTypeValue = "corner"
	ShapeTypeDiagonalStripe ShapeTypeValue = "diagStripe"
	ShapeTypePlus           ShapeTypeValue = "plus"
	ShapeTypePlaque         ShapeTypeValue = "plaque"
	ShapeTypeCan            ShapeTypeValue = "can"
	ShapeTypeCube           ShapeTypeValue = "cube"
	ShapeTypeBevel          ShapeTypeValue = "bevel"
	ShapeTypeDonut          ShapeTypeValue = "donut"
	ShapeTypeNoSmoking      ShapeTypeValue = "noSmoking"
	ShapeTypeBlockArc       ShapeTypeValue = "blockArc"
	ShapeTypeFoldedCorner   ShapeTypeValue = "foldedCorner"
	ShapeTypeSmileyFace     ShapeTypeValue = "smileyFace"
	ShapeTypeHeart          ShapeTypeValue = "heart"
	ShapeTypeLightningBolt  ShapeTypeValue = "lightningBolt"
	ShapeTypeSun            ShapeTypeValue = "sun"
	ShapeTypeMoon           ShapeTypeValue = "moon"
	ShapeTypeCloud          ShapeTypeValue = "cloud"
	ShapeTypeArc            ShapeTypeValue = "arc"
	ShapeTypeBracketPair    ShapeTypeValue = "bracketPair"
	ShapeTypeBracePair      ShapeTypeValue = "bracePair"
	ShapeTypeLeftBracket    ShapeTypeValue = "leftBracket"
	ShapeTypeRightBracket   ShapeTypeValue = "rightBracket"
	ShapeTypeLeftBrace      ShapeTypeValue = "leftBrace"
	ShapeTypeRightBrace     ShapeTypeValue = "rightBrace"

	// Arrows
	ShapeTypeRightArrow        ShapeTypeValue = "rightArrow"
	ShapeTypeLeftArrow         ShapeTypeValue = "leftArrow"
	ShapeTypeUpArrow           ShapeTypeValue = "upArrow"
	ShapeTypeDownArrow         ShapeTypeValue = "downArrow"
	ShapeTypeLeftRightArrow    ShapeTypeValue = "leftRightArrow"
	ShapeTypeUpDownArrow       ShapeTypeValue = "upDownArrow"
	ShapeTypeQuadArrow         ShapeTypeValue = "quadArrow"
	ShapeTypeBentArrow         ShapeTypeValue = "bentArrow"
	ShapeTypeUturnArrow        ShapeTypeValue = "uturnArrow"
	ShapeTypeLeftUpArrow       ShapeTypeValue = "leftUpArrow"
	ShapeTypeBentUpArrow       ShapeTypeValue = "bentUpArrow"
	ShapeTypeCurvedRightArrow  ShapeTypeValue = "curvedRightArrow"
	ShapeTypeCurvedLeftArrow   ShapeTypeValue = "curvedLeftArrow"
	ShapeTypeCurvedUpArrow     ShapeTypeValue = "curvedUpArrow"
	ShapeTypeCurvedDownArrow   ShapeTypeValue = "curvedDownArrow"
	ShapeTypeStripedRightArrow ShapeTypeValue = "stripedRightArrow"
	ShapeTypeNotchedRightArrow ShapeTypeValue = "notchedRightArrow"
	ShapeTypeCircularArrow     ShapeTypeValue = "circularArrow"
	ShapeTypeLeftCircularArrow ShapeTypeValue = "leftCircularArrow"
	ShapeTypeSwooshArrow       ShapeTypeValue = "swooshArrow"

	// Callouts
	ShapeTypeWedgeRectCallout      ShapeTypeValue = "wedgeRectCallout"
	ShapeTypeWedgeRoundRectCallout ShapeTypeValue = "wedgeRoundRectCallout"
	ShapeTypeWedgeEllipseCallout   ShapeTypeValue = "wedgeEllipseCallout"
	ShapeTypeCloudCallout          ShapeTypeValue = "cloudCallout"
	ShapeTypeBorderCallout1        ShapeTypeValue = "borderCallout1"
	ShapeTypeBorderCallout2        ShapeTypeValue = "borderCallout2"
	ShapeTypeBorderCallout3        ShapeTypeValue = "borderCallout3"
	ShapeTypeAccentCallout1        ShapeTypeValue = "accentCallout1"
	ShapeTypeAccentCallout2        ShapeTypeValue = "accentCallout2"
	ShapeTypeAccentCallout3        ShapeTypeValue = "accentCallout3"
	ShapeTypeCallout1              ShapeTypeValue = "callout1"
	ShapeTypeCallout2              ShapeTypeValue = "callout2"
	ShapeTypeCallout3              ShapeTypeValue = "callout3"

	// Stars
	ShapeTypeStar4  ShapeTypeValue = "star4"
	ShapeTypeStar5  ShapeTypeValue = "star5"
	ShapeTypeStar6  ShapeTypeValue = "star6"
	ShapeTypeStar7  ShapeTypeValue = "star7"
	ShapeTypeStar8  ShapeTypeValue = "star8"
	ShapeTypeStar10 ShapeTypeValue = "star10"
	ShapeTypeStar12 ShapeTypeValue = "star12"
	ShapeTypeStar16 ShapeTypeValue = "star16"
	ShapeTypeStar24 ShapeTypeValue = "star24"
	ShapeTypeStar32 ShapeTypeValue = "star32"

	// Flowchart shapes
	ShapeTypeFlowChartProcess           ShapeTypeValue = "flowChartProcess"
	ShapeTypeFlowChartDecision          ShapeTypeValue = "flowChartDecision"
	ShapeTypeFlowChartInputOutput       ShapeTypeValue = "flowChartInputOutput"
	ShapeTypeFlowChartPredefinedProcess ShapeTypeValue = "flowChartPredefinedProcess"
	ShapeTypeFlowChartInternalStorage   ShapeTypeValue = "flowChartInternalStorage"
	ShapeTypeFlowChartDocument          ShapeTypeValue = "flowChartDocument"
	ShapeTypeFlowChartMultidocument     ShapeTypeValue = "flowChartMultidocument"
	ShapeTypeFlowChartTerminator        ShapeTypeValue = "flowChartTerminator"
	ShapeTypeFlowChartPreparation       ShapeTypeValue = "flowChartPreparation"
	ShapeTypeFlowChartManualInput       ShapeTypeValue = "flowChartManualInput"
	ShapeTypeFlowChartManualOperation   ShapeTypeValue = "flowChartManualOperation"
	ShapeTypeFlowChartConnector         ShapeTypeValue = "flowChartConnector"
	ShapeTypeFlowChartPunchedCard       ShapeTypeValue = "flowChartPunchedCard"
	ShapeTypeFlowChartPunchedTape       ShapeTypeValue = "flowChartPunchedTape"
	ShapeTypeFlowChartSummingJunction   ShapeTypeValue = "flowChartSummingJunction"
	ShapeTypeFlowChartOr                ShapeTypeValue = "flowChartOr"
	ShapeTypeFlowChartCollate           ShapeTypeValue = "flowChartCollate"
	ShapeTypeFlowChartSort              ShapeTypeValue = "flowChartSort"
	ShapeTypeFlowChartExtract           ShapeTypeValue = "flowChartExtract"
	ShapeTypeFlowChartMerge             ShapeTypeValue = "flowChartMerge"
	ShapeTypeFlowChartOfflineStorage    ShapeTypeValue = "flowChartOfflineStorage"
	ShapeTypeFlowChartOnlineStorage     ShapeTypeValue = "flowChartOnlineStorage"
	ShapeTypeFlowChartMagneticTape      ShapeTypeValue = "flowChartMagneticTape"
	ShapeTypeFlowChartMagneticDisk      ShapeTypeValue = "flowChartMagneticDisk"
	ShapeTypeFlowChartMagneticDrum      ShapeTypeValue = "flowChartMagneticDrum"
	ShapeTypeFlowChartDisplay           ShapeTypeValue = "flowChartDisplay"
	ShapeTypeFlowChartDelay             ShapeTypeValue = "flowChartDelay"
	ShapeTypeFlowChartAlternateProcess  ShapeTypeValue = "flowChartAlternateProcess"
	ShapeTypeFlowChartOffpageConnector  ShapeTypeValue = "flowChartOffpageConnector"

	// Lines and connectors
	ShapeTypeLine               ShapeTypeValue = "line"
	ShapeTypeStraightConnector1 ShapeTypeValue = "straightConnector1"
	ShapeTypeBentConnector2     ShapeTypeValue = "bentConnector2"
	ShapeTypeBentConnector3     ShapeTypeValue = "bentConnector3"
	ShapeTypeBentConnector4     ShapeTypeValue = "bentConnector4"
	ShapeTypeBentConnector5     ShapeTypeValue = "bentConnector5"
	ShapeTypeCurvedConnector2   ShapeTypeValue = "curvedConnector2"
	ShapeTypeCurvedConnector3   ShapeTypeValue = "curvedConnector3"
	ShapeTypeCurvedConnector4   ShapeTypeValue = "curvedConnector4"
	ShapeTypeCurvedConnector5   ShapeTypeValue = "curvedConnector5"

	// Action buttons
	ShapeTypeActionButtonBlank        ShapeTypeValue = "actionButtonBlank"
	ShapeTypeActionButtonHome         ShapeTypeValue = "actionButtonHome"
	ShapeTypeActionButtonHelp         ShapeTypeValue = "actionButtonHelp"
	ShapeTypeActionButtonInformation  ShapeTypeValue = "actionButtonInformation"
	ShapeTypeActionButtonForwardNext  ShapeTypeValue = "actionButtonForwardNext"
	ShapeTypeActionButtonBackPrevious ShapeTypeValue = "actionButtonBackPrevious"
	ShapeTypeActionButtonEnd          ShapeTypeValue = "actionButtonEnd"
	ShapeTypeActionButtonBeginning    ShapeTypeValue = "actionButtonBeginning"
	ShapeTypeActionButtonReturn       ShapeTypeValue = "actionButtonReturn"
	ShapeTypeActionButtonDocument     ShapeTypeValue = "actionButtonDocument"
	ShapeTypeActionButtonSound        ShapeTypeValue = "actionButtonSound"
	ShapeTypeActionButtonMovie        ShapeTypeValue = "actionButtonMovie"

	// Equation shapes
	ShapeTypeMathPlus     ShapeTypeValue = "mathPlus"
	ShapeTypeMathMinus    ShapeTypeValue = "mathMinus"
	ShapeTypeMathMultiply ShapeTypeValue = "mathMultiply"
	ShapeTypeMathDivide   ShapeTypeValue = "mathDivide"
	ShapeTypeMathEqual    ShapeTypeValue = "mathEqual"
	ShapeTypeMathNotEqual ShapeTypeValue = "mathNotEqual"

	// Text shapes
	ShapeTypeTextBox      ShapeTypeValue = "rect"
	ShapeTypeTextPlain    ShapeTypeValue = "textPlain"
	ShapeTypeTextStop     ShapeTypeValue = "textStop"
	ShapeTypeTextTriangle ShapeTypeValue = "textTriangle"
	ShapeTypeTextWave1    ShapeTypeValue = "textWave1"
	ShapeTypeTextWave2    ShapeTypeValue = "textWave2"
	ShapeTypeTextWave4    ShapeTypeValue = "textWave4"
	ShapeTypeTextInflate  ShapeTypeValue = "textInflate"
	ShapeTypeTextDeflate  ShapeTypeValue = "textDeflate"
	ShapeTypeTextArchUp   ShapeTypeValue = "textArchUp"
	ShapeTypeTextArchDown ShapeTypeValue = "textArchDown"
	ShapeTypeTextCircle   ShapeTypeValue = "textCircle"
	ShapeTypeTextButton   ShapeTypeValue = "textButton"
)

// PresetGeometry represents a preset geometry element (a:prstGeom).
// This specifies the use of a preset shape geometry.
type PresetGeometry struct {
	*openxml.CompositeElementBase
}

// NewPresetGeometry creates a new preset geometry with the specified shape type.
func NewPresetGeometry(
	shapeType ShapeTypeValue,
) *PresetGeometry {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"prstGeom",
		PrefixMain,
	)
	pg := &PresetGeometry{
		CompositeElementBase: elem,
	}
	pg.SetPreset(shapeType)

	return pg
}

// Preset returns the preset shape type.
func (p *PresetGeometry) Preset() ShapeTypeValue {
	attr, found := p.GetAttribute("prst", "")
	if !found {
		return ""
	}

	return ShapeTypeValue(attr.Value())
}

// SetPreset sets the preset shape type.
func (p *PresetGeometry) SetPreset(
	shapeType ShapeTypeValue,
) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"prst",
			"",
			string(shapeType),
		),
	)
}

// AddAdjustValue adds an adjust value to the geometry.
// Adjust values customize the shape's proportions.
func (p *PresetGeometry) AddAdjustValue(
	name string,
	formula string,
) {
	// Find or create the avLst element
	avLst := p.GetElement("avLst", NamespaceMain)
	if avLst == nil {
		avLst = openxml.NewCompositeElement(
			NamespaceMain,
			"avLst",
			PrefixMain,
		)
		p.AppendChild(avLst)
	}

	// Create the gd (guide) element
	gd := openxml.NewLeafElement(
		NamespaceMain,
		"gd",
		PrefixMain,
	)
	gd.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
	gd.SetAttribute(
		openxml.NewAttribute(
			"",
			"fmla",
			"",
			formula,
		),
	)

	if comp, ok := avLst.(*openxml.CompositeElementBase); ok {
		comp.AppendChild(gd)
	}
}

// SetAdjustValue sets an adjust value by name.
func (p *PresetGeometry) SetAdjustValue(
	name string,
	value int,
) {
	// Create a "val" formula with the value
	formula := "val " + strconv.Itoa(value)
	p.AddAdjustValue(name, formula)
}

// ClearAdjustValues removes all adjust values.
func (p *PresetGeometry) ClearAdjustValues() {
	if avLst := p.GetElement("avLst", NamespaceMain); avLst != nil {
		p.RemoveChild(avLst)
	}
}

// Clone creates a deep copy of this PresetGeometry element.
func (p *PresetGeometry) Clone() openxml.Element {
	return &PresetGeometry{
		CompositeElementBase: p.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ShapeGuide represents a shape guide (a:gd) for geometry calculations.
type ShapeGuide struct {
	*openxml.LeafElementBase
}

// NewShapeGuide creates a new shape guide.
func NewShapeGuide(
	name, formula string,
) *ShapeGuide {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		"gd",
		PrefixMain,
	)
	sg := &ShapeGuide{LeafElementBase: elem}
	sg.SetName(name)
	sg.SetFormula(formula)

	return sg
}

// Name returns the guide name.
func (g *ShapeGuide) Name() string {
	attr, found := g.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the guide name.
func (g *ShapeGuide) SetName(name string) {
	g.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// Formula returns the guide formula.
func (g *ShapeGuide) Formula() string {
	attr, found := g.GetAttribute("fmla", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetFormula sets the guide formula.
func (g *ShapeGuide) SetFormula(formula string) {
	g.SetAttribute(
		openxml.NewAttribute(
			"",
			"fmla",
			"",
			formula,
		),
	)
}

// Clone creates a deep copy of this ShapeGuide element.
func (g *ShapeGuide) Clone() openxml.Element {
	return &ShapeGuide{
		LeafElementBase: g.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// PathFillValue represents path fill mode values.
type PathFillValue string

// Path fill mode values.
const (
	PathFillNone        PathFillValue = "none"
	PathFillNorm        PathFillValue = "norm"
	PathFillLighten     PathFillValue = "lighten"
	PathFillLightenLess PathFillValue = "lightenLess"
	PathFillDarken      PathFillValue = "darken"
	PathFillDarkenLess  PathFillValue = "darkenLess"
)

// Path2D represents a 2D path element (a:path) within a custom geometry.
type Path2D struct {
	*openxml.CompositeElementBase
}

// NewPath2D creates a new 2D path element.
func NewPath2D() *Path2D {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"path",
		PrefixMain,
	)

	return &Path2D{CompositeElementBase: elem}
}

// NewPath2DWithSize creates a new 2D path with specified dimensions.
func NewPath2DWithSize(
	width, height int64,
) *Path2D {
	path := NewPath2D()
	path.SetWidth(width)
	path.SetHeight(height)

	return path
}

// Width returns the path width in EMUs.
func (p *Path2D) Width() int64 {
	attr, found := p.GetAttribute("w", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return val
}

// SetWidth sets the path width.
func (p *Path2D) SetWidth(width int64) {
	if width <= 0 {
		p.RemoveAttribute("w", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"w",
			"",
			strconv.FormatInt(width, 10),
		),
	)
}

// Height returns the path height in EMUs.
func (p *Path2D) Height() int64 {
	attr, found := p.GetAttribute("h", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		10,
		64,
	)

	return val
}

// SetHeight sets the path height.
func (p *Path2D) SetHeight(height int64) {
	if height <= 0 {
		p.RemoveAttribute("h", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"h",
			"",
			strconv.FormatInt(height, 10),
		),
	)
}

// Fill returns the path fill mode.
func (p *Path2D) Fill() PathFillValue {
	attr, found := p.GetAttribute("fill", "")
	if !found {
		return PathFillNorm
	}

	return PathFillValue(attr.Value())
}

// SetFill sets the path fill mode.
func (p *Path2D) SetFill(fill PathFillValue) {
	if fill == "" || fill == PathFillNorm {
		p.RemoveAttribute("fill", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"fill",
			"",
			string(fill),
		),
	)
}

// Stroke returns whether the path should be stroked.
func (p *Path2D) Stroke() bool {
	attr, found := p.GetAttribute("stroke", "")
	if !found {
		return true
	}

	return attr.Value() != "0" &&
		attr.Value() != "false"
}

// SetStroke sets whether the path should be stroked.
func (p *Path2D) SetStroke(stroke bool) {
	if stroke {
		p.RemoveAttribute("stroke", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"stroke",
			"",
			"0",
		),
	)
}

// ExtrusionOk returns whether the path allows extrusion.
func (p *Path2D) ExtrusionOk() bool {
	attr, found := p.GetAttribute(
		"extrusionOk",
		"",
	)
	if !found {
		return true
	}

	return attr.Value() != "0" &&
		attr.Value() != "false"
}

// SetExtrusionOk sets whether the path allows extrusion.
func (p *Path2D) SetExtrusionOk(ok bool) {
	if ok {
		p.RemoveAttribute("extrusionOk", "")

		return
	}
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"extrusionOk",
			"",
			"0",
		),
	)
}

// AddMoveTo adds a move-to command to the path.
func (p *Path2D) AddMoveTo(x, y int64) {
	moveTo := openxml.NewCompositeElement(
		NamespaceMain,
		"moveTo",
		PrefixMain,
	)
	pt := openxml.NewLeafElement(
		NamespaceMain,
		"pt",
		PrefixMain,
	)
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(x, 10),
		),
	)
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(y, 10),
		),
	)
	moveTo.AppendChild(pt)
	p.AppendChild(moveTo)
}

// AddLineTo adds a line-to command to the path.
func (p *Path2D) AddLineTo(x, y int64) {
	lnTo := openxml.NewCompositeElement(
		NamespaceMain,
		"lnTo",
		PrefixMain,
	)
	pt := openxml.NewLeafElement(
		NamespaceMain,
		"pt",
		PrefixMain,
	)
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(x, 10),
		),
	)
	pt.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(y, 10),
		),
	)
	lnTo.AppendChild(pt)
	p.AppendChild(lnTo)
}

// AddArcTo adds an arc-to command to the path.
func (p *Path2D) AddArcTo(
	wR, hR int64,
	stAng, swAng int,
) {
	arcTo := openxml.NewLeafElement(
		NamespaceMain,
		"arcTo",
		PrefixMain,
	)
	arcTo.SetAttribute(
		openxml.NewAttribute(
			"",
			"wR",
			"",
			strconv.FormatInt(wR, 10),
		),
	)
	arcTo.SetAttribute(
		openxml.NewAttribute(
			"",
			"hR",
			"",
			strconv.FormatInt(hR, 10),
		),
	)
	arcTo.SetAttribute(
		openxml.NewAttribute(
			"",
			"stAng",
			"",
			strconv.Itoa(stAng),
		),
	)
	arcTo.SetAttribute(
		openxml.NewAttribute(
			"",
			"swAng",
			"",
			strconv.Itoa(swAng),
		),
	)
	p.AppendChild(arcTo)
}

// AddQuadBezierTo adds a quadratic bezier curve to the path.
func (p *Path2D) AddQuadBezierTo(
	ctrlX, ctrlY, endX, endY int64,
) {
	quadBezTo := openxml.NewCompositeElement(
		NamespaceMain,
		"quadBezTo",
		PrefixMain,
	)

	ctrl := openxml.NewLeafElement(
		NamespaceMain,
		"pt",
		PrefixMain,
	)
	ctrl.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(ctrlX, 10),
		),
	)
	ctrl.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(ctrlY, 10),
		),
	)
	quadBezTo.AppendChild(ctrl)

	end := openxml.NewLeafElement(
		NamespaceMain,
		"pt",
		PrefixMain,
	)
	end.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(endX, 10),
		),
	)
	end.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(endY, 10),
		),
	)
	quadBezTo.AppendChild(end)

	p.AppendChild(quadBezTo)
}

// AddCubicBezierTo adds a cubic bezier curve to the path.
func (p *Path2D) AddCubicBezierTo(
	ctrl1X, ctrl1Y, ctrl2X, ctrl2Y, endX, endY int64,
) {
	cubicBezTo := openxml.NewCompositeElement(
		NamespaceMain,
		"cubicBezTo",
		PrefixMain,
	)

	ctrl1 := openxml.NewLeafElement(
		NamespaceMain,
		"pt",
		PrefixMain,
	)
	ctrl1.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(ctrl1X, 10),
		),
	)
	ctrl1.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(ctrl1Y, 10),
		),
	)
	cubicBezTo.AppendChild(ctrl1)

	ctrl2 := openxml.NewLeafElement(
		NamespaceMain,
		"pt",
		PrefixMain,
	)
	ctrl2.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(ctrl2X, 10),
		),
	)
	ctrl2.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(ctrl2Y, 10),
		),
	)
	cubicBezTo.AppendChild(ctrl2)

	end := openxml.NewLeafElement(
		NamespaceMain,
		"pt",
		PrefixMain,
	)
	end.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(endX, 10),
		),
	)
	end.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(endY, 10),
		),
	)
	cubicBezTo.AppendChild(end)

	p.AppendChild(cubicBezTo)
}

// AddClose adds a close command to the path.
func (p *Path2D) AddClose() {
	close := openxml.NewLeafElement(
		NamespaceMain,
		"close",
		PrefixMain,
	)
	p.AppendChild(close)
}

// Clone creates a deep copy of this Path2D element.
func (p *Path2D) Clone() openxml.Element {
	return &Path2D{
		CompositeElementBase: p.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// PathList represents a path list element (a:pathLst) containing multiple paths.
type PathList struct {
	*openxml.CompositeElementBase
}

// NewPathList creates a new path list element.
func NewPathList() *PathList {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"pathLst",
		PrefixMain,
	)

	return &PathList{CompositeElementBase: elem}
}

// AddPath adds a path to the list.
func (pl *PathList) AddPath(path *Path2D) {
	pl.AppendChild(path)
}

// Clone creates a deep copy of this PathList element.
func (pl *PathList) Clone() openxml.Element {
	return &PathList{
		CompositeElementBase: pl.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CustomGeometry represents a custom geometry element (a:custGeom).
// This allows defining arbitrary shape geometries using paths.
type CustomGeometry struct {
	*openxml.CompositeElementBase
}

// NewCustomGeometry creates a new custom geometry element.
func NewCustomGeometry() *CustomGeometry {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"custGeom",
		PrefixMain,
	)
	cg := &CustomGeometry{
		CompositeElementBase: elem,
	}

	// Add required pathLst element
	pathLst := NewPathList()
	cg.AppendChild(pathLst)

	return cg
}

// AddGuide adds a shape guide to the geometry.
func (c *CustomGeometry) AddGuide(
	name, formula string,
) {
	// Find or create the gdLst element
	gdLst := c.GetElement("gdLst", NamespaceMain)
	if gdLst == nil {
		gdLst = openxml.NewCompositeElement(
			NamespaceMain,
			"gdLst",
			PrefixMain,
		)
		c.PrependChild(gdLst)
	}

	gd := NewShapeGuide(name, formula)
	if comp, ok := gdLst.(*openxml.CompositeElementBase); ok {
		comp.AppendChild(gd)
	}
}

// AddAdjustValue adds an adjust value to the geometry.
func (c *CustomGeometry) AddAdjustValue(
	name, formula string,
) {
	// Find or create the avLst element
	avLst := c.GetElement("avLst", NamespaceMain)
	if avLst == nil {
		avLst = openxml.NewCompositeElement(
			NamespaceMain,
			"avLst",
			PrefixMain,
		)
		// Insert at beginning, before gdLst
		c.PrependChild(avLst)
	}

	gd := NewShapeGuide(name, formula)
	if comp, ok := avLst.(*openxml.CompositeElementBase); ok {
		comp.AppendChild(gd)
	}
}

// PathList returns the path list element.
func (c *CustomGeometry) PathList() *PathList {
	elem := c.GetElement("pathLst", NamespaceMain)
	if elem == nil {
		return nil
	}
	if pl, ok := elem.(*PathList); ok {
		return pl
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PathList{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// AddPath adds a path to the geometry's path list.
func (c *CustomGeometry) AddPath(path *Path2D) {
	pathLst := c.PathList()
	if pathLst == nil {
		pathLst = NewPathList()
		c.AppendChild(pathLst)
	}
	pathLst.AddPath(path)
}

// SetRectangle sets the text rectangle for the geometry.
func (c *CustomGeometry) SetRectangle(
	left, top, right, bottom string,
) {
	// Remove existing rect
	if rect := c.GetElement("rect", NamespaceMain); rect != nil {
		c.RemoveChild(rect)
	}

	rect := openxml.NewLeafElement(
		NamespaceMain,
		"rect",
		PrefixMain,
	)
	rect.SetAttribute(
		openxml.NewAttribute("", "l", "", left),
	)
	rect.SetAttribute(
		openxml.NewAttribute("", "t", "", top),
	)
	rect.SetAttribute(
		openxml.NewAttribute("", "r", "", right),
	)
	rect.SetAttribute(
		openxml.NewAttribute("", "b", "", bottom),
	)

	// Insert before pathLst
	pathLst := c.GetElement(
		"pathLst",
		NamespaceMain,
	)
	if pathLst != nil {
		c.InsertBefore(rect, pathLst)
	} else {
		c.AppendChild(rect)
	}
}

// Clone creates a deep copy of this CustomGeometry element.
func (c *CustomGeometry) Clone() openxml.Element {
	return &CustomGeometry{
		CompositeElementBase: c.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
